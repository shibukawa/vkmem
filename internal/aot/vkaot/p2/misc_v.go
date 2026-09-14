package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_validateProcTitleTemplate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	v2 = int32(0)
	v7 = F_sdstemplate(m, l0, int32(1033), int32(_a139))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v135
L2:
	;
	return int32(0)
L3:
	;
	if v7 == int32(0) {
		v135 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v13 = int32(_a66)
	v19 = v7 + int32(-1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	switch v20 & int32(7) {
	case 0:
		goto L11
	case 1:
		goto L10
	case 2:
		goto L9
	case 3:
		goto L8
	case 4:
		goto L7
	default:
		v37 = int32(0)
		goto L6
	}
L5:
	;
	if v7 == int32(0) {
		v135 = v2
		goto L1
	} else {
		goto L32
	}
L6:
	;
	v40 = v7 + v37 + int32(-1)
	if base.Ui32(v40) < base.Ui32(v7) {
		v58 = v7
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-17))))
	v37 = v36
	goto L6
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-9))))
	v37 = v33
	goto L6
L9:
	;
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(-5)))))
	v37 = v30
	goto L6
L10:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-3)))))
	v37 = v27
	goto L6
L11:
	;
	v37 = int32(base.Ui32(v20) >> (uint(int32(3)) % 32))
	goto L6
L12:
	;
	if base.Ui32(v40) <= base.Ui32(v58) {
		v74 = v40
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v46 = v7
	goto L14
L14:
	;
	v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(v46))))
	v48 = F_strchr(m, v13, v47)
	mBase = m.M
	if v48 == int32(0) {
		v58 = v46
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v58 = v52
	goto L12
L16:
	;
	v52 = v46 + int32(1)
	if base.Ui32(v52) <= base.Ui32(v40) {
		v46 = v52
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v79 = v74 - v58 + int32(1)
	if v7 == v58 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v62 = v40
	goto L20
L20:
	;
	v65 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62))))
	v66 = F_strchr(m, v13, v65)
	mBase = m.M
	if v66 == int32(0) {
		v74 = v62
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v74 = v58
	goto L18
L22:
	;
	v70 = v62 + int32(-1)
	if base.Ui32(v58) < base.Ui32(v70) {
		v62 = v70
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v83 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7+v79))) = uint8(v83)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	switch v85 & int32(7) {
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
		goto L26
	}
L25:
	;
	v81 = F_memmove(m, v7, v58, v79)
	mBase = m.M
	goto L24
L26:
	;
	goto L5
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(-17)))) = base.I64_extend_i32_u(v79)
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7+int32(-9)))) = v79
	goto L5
L29:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(-5)))) = uint16(v79)
	goto L5
L30:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-3)))) = uint8(v79)
	goto L5
L31:
	;
	v89 = v79 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v89)
	goto L5
L32:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-1)))))
	switch v108 & int32(7) {
	case 0:
		goto L40
	case 1:
		goto L39
	case 2:
		goto L38
	case 3:
		goto L37
	case 4:
		goto L36
	default:
		goto L34
	}
L33:
	;
	F_sdsfree(m, v7)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L42
	}
L34:
	;
	v130 = int32(0)
	goto L33
L35:
	;
	if v125 != 0 {
		v130 = int32(1)
		goto L33
	} else {
		goto L41
	}
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-17))))
	v125 = v124
	goto L35
L37:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-9))))
	v125 = v121
	goto L35
L38:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(-5)))))
	v125 = v118
	goto L35
L39:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-3)))))
	v125 = v115
	goto L35
L40:
	;
	v125 = int32(base.Ui32(v108) >> (uint(int32(3)) % 32))
	goto L35
L41:
	;
	goto L34
L42:
	;
	v135 = v130
	goto L1
}
func F_valkeyvFormatCommand(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v119 int64
	_ = v119
	var v133 int32
	_ = v133
	var v137 int64
	_ = v137
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1095 int32
	_ = v1095
	var v1102 int32
	_ = v1102
	var v1115 int32
	_ = v1115
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1142 int32
	_ = v1142
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1166 int32
	_ = v1166
	var v1177 int64
	_ = v1177
	var v1191 int32
	_ = v1191
	var v1195 int64
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1215 int32
	_ = v1215
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1237 int32
	_ = v1237
	var v1248 int64
	_ = v1248
	var v1262 int32
	_ = v1262
	var v1266 int64
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1304 int32
	_ = v1304
	var v1312 int32
	_ = v1312
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1419 int32
	_ = v1419
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1438 int32
	_ = v1438
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1473 int32
	_ = v1473
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1531 int32
	_ = v1531
	var v1545 int32
	_ = v1545
	v15 = m.G0
	v17 = v15 - int32(64)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = l2
	v20 = int32(-1)
	if l0 == int32(0) {
		v1531 = v20
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v1545 = m.G3
	m.Env.X__assert_fail(m, v1545+int32(_a1909), v1545+int32(_a1910), int32(542), v1545+int32(_a1911))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	m.G0 = v17 + int32(64)
	return v1531
L3:
	;
	v23 = F_sdsempty(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v23 == int32(0) {
		v1531 = v20
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v29 = int32(0)
	v34 = l1
	v35 = v29
	v38 = v23
	v39 = v29
	v40 = v29
	v41 = v29
	goto L13
L7:
	;
	if v1454 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L8:
	;
	v1452 = int32(-1)
	v1453 = v1438
	v1454 = v39
	v1455 = v40
	goto L7
L9:
	;
	if base.Ui32(v1227) < base.Ui32(int32(10)) {
		v1269 = int32(1)
		goto L297
	} else {
		goto L298
	}
L10:
	;
	F_sdsfree(m, v38)
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L4
	} else {
		goto L296
	}
L11:
	;
	v1124 = int32(1)
	v1126 = v40 + v1124
	v1129 = m.G4
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+8))
	v1131 = m.T0[v1130].(func(*base.Module, int32, int32) int32)(m, v39, v1126<<(uint(int32(2))%32))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L4
	} else {
		goto L277
	}
L12:
	;
	if v35 == int32(0) {
		goto L10
	} else {
		goto L276
	}
L13:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	switch v47 + int32(-32) {
	case 0:
		goto L22
	case 1, 2, 3, 4:
		goto L20
	case 5:
		goto L23
	default:
		goto L21
	}
L15:
	;
	if v1095 == int32(0) {
		v1438 = v38
		goto L8
	} else {
		goto L273
	}
L16:
	;
	v1084 = F_sdscatlen(m, v38, v177, v233)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L4
	} else {
		goto L272
	}
L17:
	;
	v250 = v50
	v260 = v34 + int32(1)
	goto L80
L18:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v234 + int32(4)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if v238 != 0 {
		goto L74
	} else {
		goto L75
	}
L19:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v173 + int32(4)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	if v177 != 0 {
		goto L55
	} else {
		goto L56
	}
L20:
	;
	v165 = int32(1)
	v167 = F_sdscatlen(m, v38, v34, v165)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L53
	}
L21:
	;
	if v47 == int32(0) {
		goto L12
	} else {
		goto L52
	}
L22:
	;
	if v35 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	switch v50 + int32(-98) {
	case 0:
		goto L18
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16:
		goto L17
	case 17:
		goto L19
	default:
		goto L24
	}
L24:
	;
	if v50 == int32(0) {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	if v50 != int32(37) {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	v57 = m.G3
	v60 = F_sdscat(m, v38, v57+int32(_a1912))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v1087 = v34
	v1095 = v60
	goto L15
L28:
	;
	v65 = int32(-1)
	v67 = v40 + int32(1)
	v70 = m.G4
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v72 = m.T0[v71].(func(*base.Module, int32, int32) int32)(m, v39, v67<<(uint(int32(2))%32))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	v34 = v34 + int32(1)
	v35 = int32(0)
	goto L13
L30:
	;
	if v72 == int32(0) {
		v1452 = v65
		v1453 = v38
		v1454 = v39
		v1455 = v40
		goto L7
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72+v40<<(uint(int32(2))%32)))) = v38
	v80 = int32(1)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+int32(-1)))))
	switch v84 & int32(7) {
	case 0:
		goto L38
	case 1:
		goto L37
	case 2:
		goto L36
	case 3:
		goto L35
	case 4:
		goto L34
	default:
		v140 = v80
		v150 = int32(0)
		goto L32
	}
L32:
	;
	v152 = F_sdsempty(m)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L50
	}
L33:
	;
	if base.Ui32(v101) < base.Ui32(int32(10)) {
		v140 = v80
		v150 = v101
		goto L32
	} else {
		goto L39
	}
L34:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(-17))))
	v101 = v100
	goto L33
L35:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(-9))))
	v101 = v97
	goto L33
L36:
	;
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38+int32(-5)))))
	v101 = v94
	goto L33
L37:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+int32(-3)))))
	v101 = v91
	goto L33
L38:
	;
	v101 = int32(base.Ui32(v84) >> (uint(int32(3)) % 32))
	goto L33
L39:
	;
	v108 = int32(1)
	v119 = base.I64_extend_i32_u(v101)
	goto L40
L40:
	;
	if base.Ui64(int64(99)) < base.Ui64(v119) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v140 = v133
	v150 = v101
	goto L32
L42:
	;
	if base.Ui64(int64(999)) < base.Ui64(v119) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v140 = v108 + int32(1)
	v150 = v101
	goto L32
L44:
	;
	if base.Ui64(int64(9999)) < base.Ui64(v119) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v140 = v108 + int32(2)
	v150 = v101
	goto L32
L46:
	;
	v133 = v108 + int32(4)
	v137 = base.I64_div_u_s(v119, int64(10000))
	if base.Ui64(int64(99999)) < base.Ui64(v119) {
		v108 = v133
		v119 = v137
		goto L40
	} else {
		goto L48
	}
L47:
	;
	v140 = v108 + int32(3)
	v150 = v101
	goto L32
L48:
	;
	goto L41
L49:
	;
	v34 = v34 + int32(1)
	v35 = int32(0)
	v38 = v152
	v39 = v72
	v40 = v67
	v41 = v41 + v150 + v140 + int32(5)
	goto L13
L50:
	;
	if v152 != 0 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v1452 = v65
	v1453 = int32(0)
	v1454 = v72
	v1455 = v67
	goto L7
L52:
	;
	goto L20
L53:
	;
	if v167 == int32(0) {
		v1438 = v38
		goto L8
	} else {
		goto L54
	}
L54:
	;
	v34 = v34 + int32(1)
	v35 = v165
	v38 = v167
	goto L13
L55:
	;
	if v177&int32(3) == int32(0) {
		v200 = v177
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v1452 = int32(-2)
	v1453 = v38
	v1454 = v39
	v1455 = v40
	goto L7
L57:
	;
	if v233 != 0 {
		goto L16
	} else {
		goto L73
	}
L58:
	;
	v233 = v225 - v177
	goto L57
L59:
	;
	v204 = v200
	goto L67
L60:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if v186 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v189 = v177
	goto L63
L62:
	;
	v233 = v177 - v177
	goto L57
L63:
	;
	v193 = v189 + int32(1)
	if v193&int32(3) == int32(0) {
		v200 = v193
		goto L59
	} else {
		goto L65
	}
L65:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	if v198 != 0 {
		v189 = v193
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v225 = v193
	goto L58
L67:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v213 = int32(-2139062144)
	if (int32(16843008)-v210|v210)&v213 == v213 {
		v204 = v204 + int32(4)
		goto L67
	} else {
		goto L69
	}
L68:
	;
	v219 = v204
	goto L70
L69:
	;
	goto L68
L70:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	if v223 != 0 {
		v219 = v219 + int32(1)
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v225 = v219
	goto L58
L72:
	;
	goto L71
L73:
	;
	v1087 = v34
	v1095 = v38
	goto L15
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v234 + int32(8)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if v243 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v1452 = int32(-2)
	v1453 = v38
	v1454 = v39
	v1455 = v40
	goto L7
L76:
	;
	v244 = F_sdscatlen(m, v38, v238, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L78
	}
L77:
	;
	v1087 = v34
	v1095 = v38
	goto L15
L78:
	;
	v1087 = v34
	v1095 = v244
	goto L15
L79:
	;
	v292 = v281
	goto L87
L80:
	;
	if base.Ui32(v250&int32(255)) <= base.Ui32(int32(63)) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v281 = v279
	goto L79
L82:
	;
	if base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v250)&int64(255))%64)&int64(325494096527361) == int64(0)) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v281 = v260
	goto L79
L84:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
	v279 = v260 + int32(1)
	if v277 != 0 {
		v250 = v277
		v260 = v279
		goto L80
	} else {
		goto L86
	}
L85:
	;
	v281 = v260
	goto L79
L86:
	;
	goto L81
L87:
	;
	v299 = int32(*(*int8)(unsafe.Add(mBase, uint32(v292))))
	if base.Ui32(v299+int32(-48)) < base.Ui32(int32(10)) {
		v292 = v292 + int32(1)
		goto L87
	} else {
		goto L89
	}
L88:
	;
	if v299 != int32(46) {
		v329 = v292
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L88
L90:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v341
	v343 = int32(*(*int8)(unsafe.Add(mBase, uint32(v329))))
	v345 = v343 & int32(255)
	if v345 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L91:
	;
	v308 = v292
	goto L92
L92:
	;
	v320 = int32(*(*int8)(unsafe.Add(mBase, uint32(v308)+1)))
	v322 = v308 + int32(1)
	if base.Ui32(v320+int32(-48)) < base.Ui32(int32(10)) {
		v308 = v322
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v329 = v322
	goto L90
L94:
	;
	goto L93
L95:
	;
	v1452 = int32(-2)
	v1453 = v38
	v1454 = v39
	v1455 = v40
	goto L7
L96:
	;
	v348 = m.G3
	v350 = v348 + int32(_a1913)
	v351 = int32(7)
	if v350&int32(3) == int32(0) {
		v381 = v350
		v383 = v351
		v384 = int32(1)
		goto L102
	} else {
		goto L103
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v1057
	v1061 = v1056 - v34 + int32(1)
	if base.Ui32(v1061) <= base.Ui32(int32(13)) {
		goto L266
	} else {
		goto L267
	}
L98:
	;
	v460 = m.G3
	v462 = v460 + int32(_a1914)
	v463 = int32(9)
	if v462&int32(3) == int32(0) {
		v493 = v462
		v495 = v463
		v496 = int32(1)
		goto L129
	} else {
		goto L130
	}
L99:
	;
	if v454 == int32(0) {
		goto L98
	} else {
		goto L124
	}
L100:
	;
	v454 = int32(0)
	goto L99
L101:
	;
	v432 = v425
	v434 = v427
	goto L119
L102:
	;
	if v384 == int32(0) {
		goto L100
	} else {
		goto L110
	}
L103:
	;
	goto L104
L104:
	;
	v364 = v350
	v366 = v351
	goto L105
L105:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364))))
	if v369 == v343&int32(255) {
		v425 = v364
		v427 = v366
		goto L101
	} else {
		goto L107
	}
L106:
	;
	v381 = v376
	v383 = v372
	v384 = v374
	goto L102
L107:
	;
	v372 = v366 + int32(-1)
	v373 = int32(0)
	v374 = base.B2i32(v372 != v373)
	v376 = v364 + int32(1)
	if v376&int32(3) == v373 {
		v381 = v376
		v383 = v372
		v384 = v374
		goto L102
	} else {
		goto L108
	}
L108:
	;
	if v372 != 0 {
		v364 = v376
		v366 = v372
		goto L105
	} else {
		goto L109
	}
L109:
	;
	goto L106
L110:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	if v388 == v343&int32(255) {
		v418 = v381
		v420 = v383
		goto L111
	} else {
		goto L112
	}
L111:
	;
	if v420 == int32(0) {
		goto L100
	} else {
		goto L118
	}
L112:
	;
	if base.Ui32(v383) < base.Ui32(int32(4)) {
		v418 = v381
		v420 = v383
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v398 = v381
	v400 = v383
	goto L114
L114:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	v405 = v404 ^ v343&int32(255)*int32(16843009)
	v408 = int32(-2139062144)
	if (int32(16843008)-v405|v405)&v408 != v408 {
		v425 = v398
		v427 = v400
		goto L101
	} else {
		goto L116
	}
L115:
	;
	v418 = v413
	v420 = v415
	goto L111
L116:
	;
	v413 = v398 + int32(4)
	v415 = v400 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v415) {
		v398 = v413
		v400 = v415
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v425 = v418
	v427 = v420
	goto L101
L119:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	if v437 != v343&int32(255) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	goto L100
L121:
	;
	v442 = v434 + int32(-1)
	if v442 != 0 {
		v432 = v432 + int32(1)
		v434 = v442
		goto L119
	} else {
		goto L123
	}
L122:
	;
	v454 = v432
	goto L99
L123:
	;
	goto L120
L124:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	v1056 = v329
	v1057 = v457 + int32(4)
	goto L97
L125:
	;
	switch v345 + int32(-104) {
	case 0:
		goto L153
	default:
		goto L95
	case 4:
		goto L152
	}
L126:
	;
	if v566 == int32(0) {
		goto L125
	} else {
		goto L151
	}
L127:
	;
	v566 = int32(0)
	goto L126
L128:
	;
	v544 = v537
	v546 = v539
	goto L146
L129:
	;
	if v496 == int32(0) {
		goto L127
	} else {
		goto L137
	}
L130:
	;
	goto L131
L131:
	;
	v476 = v462
	v478 = v463
	goto L132
L132:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
	if v481 == v343&int32(255) {
		v537 = v476
		v539 = v478
		goto L128
	} else {
		goto L134
	}
L133:
	;
	v493 = v488
	v495 = v484
	v496 = v486
	goto L129
L134:
	;
	v484 = v478 + int32(-1)
	v485 = int32(0)
	v486 = base.B2i32(v484 != v485)
	v488 = v476 + int32(1)
	if v488&int32(3) == v485 {
		v493 = v488
		v495 = v484
		v496 = v486
		goto L129
	} else {
		goto L135
	}
L135:
	;
	if v484 != 0 {
		v476 = v488
		v478 = v484
		goto L132
	} else {
		goto L136
	}
L136:
	;
	goto L133
L137:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
	if v500 == v343&int32(255) {
		v530 = v493
		v532 = v495
		goto L138
	} else {
		goto L139
	}
L138:
	;
	if v532 == int32(0) {
		goto L127
	} else {
		goto L145
	}
L139:
	;
	if base.Ui32(v495) < base.Ui32(int32(4)) {
		v530 = v493
		v532 = v495
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v510 = v493
	v512 = v495
	goto L141
L141:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v510)))
	v517 = v516 ^ v343&int32(255)*int32(16843009)
	v520 = int32(-2139062144)
	if (int32(16843008)-v517|v517)&v520 != v520 {
		v537 = v510
		v539 = v512
		goto L128
	} else {
		goto L143
	}
L142:
	;
	v530 = v525
	v532 = v527
	goto L138
L143:
	;
	v525 = v510 + int32(4)
	v527 = v512 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v527) {
		v510 = v525
		v512 = v527
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v537 = v530
	v539 = v532
	goto L128
L146:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544))))
	if v549 != v343&int32(255) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L127
L148:
	;
	v554 = v546 + int32(-1)
	if v554 != 0 {
		v544 = v544 + int32(1)
		v546 = v554
		goto L146
	} else {
		goto L150
	}
L149:
	;
	v566 = v544
	goto L126
L150:
	;
	goto L147
L151:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	v1056 = v329
	v1057 = (v569+int32(7))&int32(-8) + int32(8)
	goto L97
L152:
	;
	v815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+1)))
	if v815 == int32(108) {
		goto L210
	} else {
		goto L211
	}
L153:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+1)))
	if v578 == int32(104) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v698 = int32(*(*int8)(unsafe.Add(mBase, uint32(v329)+2)))
	if v698 == int32(0) {
		goto L95
	} else {
		goto L183
	}
L155:
	;
	if v578 == int32(0) {
		goto L95
	} else {
		goto L156
	}
L156:
	;
	v583 = m.G3
	v585 = v583 + int32(_a1913)
	v586 = base.I32_extend8_s(v578)
	v587 = int32(7)
	if v585&int32(3) == int32(0) {
		v617 = v585
		v619 = v587
		v620 = int32(1)
		goto L160
	} else {
		goto L161
	}
L157:
	;
	if v690 == int32(0) {
		goto L95
	} else {
		goto L182
	}
L158:
	;
	v690 = int32(0)
	goto L157
L159:
	;
	v668 = v661
	v670 = v663
	goto L177
L160:
	;
	if v620 == int32(0) {
		goto L158
	} else {
		goto L168
	}
L161:
	;
	goto L162
L162:
	;
	v600 = v585
	v602 = v587
	goto L163
L163:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600))))
	if v605 == v586&int32(255) {
		v661 = v600
		v663 = v602
		goto L159
	} else {
		goto L165
	}
L164:
	;
	v617 = v612
	v619 = v608
	v620 = v610
	goto L160
L165:
	;
	v608 = v602 + int32(-1)
	v609 = int32(0)
	v610 = base.B2i32(v608 != v609)
	v612 = v600 + int32(1)
	if v612&int32(3) == v609 {
		v617 = v612
		v619 = v608
		v620 = v610
		goto L160
	} else {
		goto L166
	}
L166:
	;
	if v608 != 0 {
		v600 = v612
		v602 = v608
		goto L163
	} else {
		goto L167
	}
L167:
	;
	goto L164
L168:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617))))
	if v624 == v586&int32(255) {
		v654 = v617
		v656 = v619
		goto L169
	} else {
		goto L170
	}
L169:
	;
	if v656 == int32(0) {
		goto L158
	} else {
		goto L176
	}
L170:
	;
	if base.Ui32(v619) < base.Ui32(int32(4)) {
		v654 = v617
		v656 = v619
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v634 = v617
	v636 = v619
	goto L172
L172:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v634)))
	v641 = v640 ^ v586&int32(255)*int32(16843009)
	v644 = int32(-2139062144)
	if (int32(16843008)-v641|v641)&v644 != v644 {
		v661 = v634
		v663 = v636
		goto L159
	} else {
		goto L174
	}
L173:
	;
	v654 = v649
	v656 = v651
	goto L169
L174:
	;
	v649 = v634 + int32(4)
	v651 = v636 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v651) {
		v634 = v649
		v636 = v651
		goto L172
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	v661 = v654
	v663 = v656
	goto L159
L177:
	;
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668))))
	if v673 != v586&int32(255) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	goto L158
L179:
	;
	v678 = v670 + int32(-1)
	if v678 != 0 {
		v668 = v668 + int32(1)
		v670 = v678
		goto L177
	} else {
		goto L181
	}
L180:
	;
	v690 = v668
	goto L157
L181:
	;
	goto L178
L182:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	v1056 = v329 + int32(1)
	v1057 = v695 + int32(4)
	goto L97
L183:
	;
	v701 = m.G3
	v703 = v701 + int32(_a1913)
	v704 = int32(7)
	if v703&int32(3) == int32(0) {
		v734 = v703
		v736 = v704
		v737 = int32(1)
		goto L187
	} else {
		goto L188
	}
L184:
	;
	if v807 == int32(0) {
		goto L95
	} else {
		goto L209
	}
L185:
	;
	v807 = int32(0)
	goto L184
L186:
	;
	v785 = v778
	v787 = v780
	goto L204
L187:
	;
	if v737 == int32(0) {
		goto L185
	} else {
		goto L195
	}
L188:
	;
	goto L189
L189:
	;
	v717 = v703
	v719 = v704
	goto L190
L190:
	;
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717))))
	if v722 == v698&int32(255) {
		v778 = v717
		v780 = v719
		goto L186
	} else {
		goto L192
	}
L191:
	;
	v734 = v729
	v736 = v725
	v737 = v727
	goto L187
L192:
	;
	v725 = v719 + int32(-1)
	v726 = int32(0)
	v727 = base.B2i32(v725 != v726)
	v729 = v717 + int32(1)
	if v729&int32(3) == v726 {
		v734 = v729
		v736 = v725
		v737 = v727
		goto L187
	} else {
		goto L193
	}
L193:
	;
	if v725 != 0 {
		v717 = v729
		v719 = v725
		goto L190
	} else {
		goto L194
	}
L194:
	;
	goto L191
L195:
	;
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734))))
	if v741 == v698&int32(255) {
		v771 = v734
		v773 = v736
		goto L196
	} else {
		goto L197
	}
L196:
	;
	if v773 == int32(0) {
		goto L185
	} else {
		goto L203
	}
L197:
	;
	if base.Ui32(v736) < base.Ui32(int32(4)) {
		v771 = v734
		v773 = v736
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v751 = v734
	v753 = v736
	goto L199
L199:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v751)))
	v758 = v757 ^ v698&int32(255)*int32(16843009)
	v761 = int32(-2139062144)
	if (int32(16843008)-v758|v758)&v761 != v761 {
		v778 = v751
		v780 = v753
		goto L186
	} else {
		goto L201
	}
L200:
	;
	v771 = v766
	v773 = v768
	goto L196
L201:
	;
	v766 = v751 + int32(4)
	v768 = v753 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v768) {
		v751 = v766
		v753 = v768
		goto L199
	} else {
		goto L202
	}
L202:
	;
	goto L200
L203:
	;
	v778 = v771
	v780 = v773
	goto L186
L204:
	;
	v790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v785))))
	if v790 != v698&int32(255) {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	goto L185
L206:
	;
	v795 = v787 + int32(-1)
	if v795 != 0 {
		v785 = v785 + int32(1)
		v787 = v795
		goto L204
	} else {
		goto L208
	}
L207:
	;
	v807 = v785
	goto L184
L208:
	;
	goto L205
L209:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	v1056 = v329 + int32(2)
	v1057 = v812 + int32(4)
	goto L97
L210:
	;
	v935 = int32(*(*int8)(unsafe.Add(mBase, uint32(v329)+2)))
	if v935 == int32(0) {
		goto L95
	} else {
		goto L239
	}
L211:
	;
	if v815 == int32(0) {
		goto L95
	} else {
		goto L212
	}
L212:
	;
	v820 = m.G3
	v822 = v820 + int32(_a1913)
	v823 = base.I32_extend8_s(v815)
	v824 = int32(7)
	if v822&int32(3) == int32(0) {
		v854 = v822
		v856 = v824
		v857 = int32(1)
		goto L216
	} else {
		goto L217
	}
L213:
	;
	if v927 == int32(0) {
		goto L95
	} else {
		goto L238
	}
L214:
	;
	v927 = int32(0)
	goto L213
L215:
	;
	v905 = v898
	v907 = v900
	goto L233
L216:
	;
	if v857 == int32(0) {
		goto L214
	} else {
		goto L224
	}
L217:
	;
	goto L218
L218:
	;
	v837 = v822
	v839 = v824
	goto L219
L219:
	;
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837))))
	if v842 == v823&int32(255) {
		v898 = v837
		v900 = v839
		goto L215
	} else {
		goto L221
	}
L220:
	;
	v854 = v849
	v856 = v845
	v857 = v847
	goto L216
L221:
	;
	v845 = v839 + int32(-1)
	v846 = int32(0)
	v847 = base.B2i32(v845 != v846)
	v849 = v837 + int32(1)
	if v849&int32(3) == v846 {
		v854 = v849
		v856 = v845
		v857 = v847
		goto L216
	} else {
		goto L222
	}
L222:
	;
	if v845 != 0 {
		v837 = v849
		v839 = v845
		goto L219
	} else {
		goto L223
	}
L223:
	;
	goto L220
L224:
	;
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v854))))
	if v861 == v823&int32(255) {
		v891 = v854
		v893 = v856
		goto L225
	} else {
		goto L226
	}
L225:
	;
	if v893 == int32(0) {
		goto L214
	} else {
		goto L232
	}
L226:
	;
	if base.Ui32(v856) < base.Ui32(int32(4)) {
		v891 = v854
		v893 = v856
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v871 = v854
	v873 = v856
	goto L228
L228:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v871)))
	v878 = v877 ^ v823&int32(255)*int32(16843009)
	v881 = int32(-2139062144)
	if (int32(16843008)-v878|v878)&v881 != v881 {
		v898 = v871
		v900 = v873
		goto L215
	} else {
		goto L230
	}
L229:
	;
	v891 = v886
	v893 = v888
	goto L225
L230:
	;
	v886 = v871 + int32(4)
	v888 = v873 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v888) {
		v871 = v886
		v873 = v888
		goto L228
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	v898 = v891
	v900 = v893
	goto L215
L233:
	;
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v905))))
	if v910 != v823&int32(255) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	goto L214
L235:
	;
	v915 = v907 + int32(-1)
	if v915 != 0 {
		v905 = v905 + int32(1)
		v907 = v915
		goto L233
	} else {
		goto L237
	}
L236:
	;
	v927 = v905
	goto L213
L237:
	;
	goto L234
L238:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	v1056 = v329 + int32(1)
	v1057 = v932 + int32(4)
	goto L97
L239:
	;
	v938 = m.G3
	v940 = v938 + int32(_a1913)
	v941 = int32(7)
	if v940&int32(3) == int32(0) {
		v971 = v940
		v973 = v941
		v974 = int32(1)
		goto L243
	} else {
		goto L244
	}
L240:
	;
	if v1044 == int32(0) {
		goto L95
	} else {
		goto L265
	}
L241:
	;
	v1044 = int32(0)
	goto L240
L242:
	;
	v1022 = v1015
	v1024 = v1017
	goto L260
L243:
	;
	if v974 == int32(0) {
		goto L241
	} else {
		goto L251
	}
L244:
	;
	goto L245
L245:
	;
	v954 = v940
	v956 = v941
	goto L246
L246:
	;
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v954))))
	if v959 == v935&int32(255) {
		v1015 = v954
		v1017 = v956
		goto L242
	} else {
		goto L248
	}
L247:
	;
	v971 = v966
	v973 = v962
	v974 = v964
	goto L243
L248:
	;
	v962 = v956 + int32(-1)
	v963 = int32(0)
	v964 = base.B2i32(v962 != v963)
	v966 = v954 + int32(1)
	if v966&int32(3) == v963 {
		v971 = v966
		v973 = v962
		v974 = v964
		goto L243
	} else {
		goto L249
	}
L249:
	;
	if v962 != 0 {
		v954 = v966
		v956 = v962
		goto L246
	} else {
		goto L250
	}
L250:
	;
	goto L247
L251:
	;
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971))))
	if v978 == v935&int32(255) {
		v1008 = v971
		v1010 = v973
		goto L252
	} else {
		goto L253
	}
L252:
	;
	if v1010 == int32(0) {
		goto L241
	} else {
		goto L259
	}
L253:
	;
	if base.Ui32(v973) < base.Ui32(int32(4)) {
		v1008 = v971
		v1010 = v973
		goto L252
	} else {
		goto L254
	}
L254:
	;
	v988 = v971
	v990 = v973
	goto L255
L255:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v988)))
	v995 = v994 ^ v935&int32(255)*int32(16843009)
	v998 = int32(-2139062144)
	if (int32(16843008)-v995|v995)&v998 != v998 {
		v1015 = v988
		v1017 = v990
		goto L242
	} else {
		goto L257
	}
L256:
	;
	v1008 = v1003
	v1010 = v1005
	goto L252
L257:
	;
	v1003 = v988 + int32(4)
	v1005 = v990 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v1005) {
		v988 = v1003
		v990 = v1005
		goto L255
	} else {
		goto L258
	}
L258:
	;
	goto L256
L259:
	;
	v1015 = v1008
	v1017 = v1010
	goto L242
L260:
	;
	v1027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1022))))
	if v1027 != v935&int32(255) {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	goto L241
L262:
	;
	v1032 = v1024 + int32(-1)
	if v1032 != 0 {
		v1022 = v1022 + int32(1)
		v1024 = v1032
		goto L260
	} else {
		goto L264
	}
L263:
	;
	v1044 = v1022
	goto L240
L264:
	;
	goto L261
L265:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	v1056 = v329 + int32(2)
	v1057 = (v1049+int32(7))&int32(-8) + int32(8)
	goto L97
L266:
	;
	if v1061 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	v1087 = v34
	v1095 = v38
	goto L15
L268:
	;
	v1071 = v17 + int32(32)
	v1073 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1071+v1061))) = uint8(v1073)
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v1080 = F_sdscatvprintf(m, v38, v1071, v1079)
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L4
	} else {
		goto L271
	}
L269:
	;
	goto L268
L270:
	;
	v1068 = F__emscripten_memcpy_bulkmem(m, v17+int32(32), v34, v1061)
	mBase = m.M
	goto L269
L271:
	;
	v1087 = v1056 + int32(-1)
	v1095 = v1080
	goto L15
L272:
	;
	v1087 = v34
	v1095 = v1084
	goto L15
L273:
	;
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1087)+1)))
	if v1102 != 0 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v34 = v1087 + int32(2)
	v35 = int32(1)
	v38 = v1095
	goto L13
L275:
	;
	v1115 = v1095
	goto L11
L276:
	;
	v1115 = v38
	goto L11
L277:
	;
	if v1131 == int32(0) {
		v1438 = v1115
		goto L8
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1131+v40<<(uint(int32(2))%32)))) = v1115
	v1142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1115+int32(-1)))))
	switch v1142 & int32(7) {
	case 0:
		goto L285
	case 1:
		goto L284
	case 2:
		goto L283
	case 3:
		goto L282
	case 4:
		goto L281
	default:
		v1197 = int32(0)
		v1198 = v1124
		goto L279
	}
L279:
	;
	v1224 = v41 + v1197 + v1198 + int32(5)
	v1226 = v1131
	v1227 = v1126
	goto L9
L280:
	;
	if base.Ui32(v1159) < base.Ui32(int32(10)) {
		v1197 = v1159
		v1198 = v1124
		goto L279
	} else {
		goto L286
	}
L281:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1115+int32(-17))))
	v1159 = v1158
	goto L280
L282:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1115+int32(-9))))
	v1159 = v1155
	goto L280
L283:
	;
	v1152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1115+int32(-5)))))
	v1159 = v1152
	goto L280
L284:
	;
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1115+int32(-3)))))
	v1159 = v1149
	goto L280
L285:
	;
	v1159 = int32(base.Ui32(v1142) >> (uint(int32(3)) % 32))
	goto L280
L286:
	;
	v1166 = int32(1)
	v1177 = base.I64_extend_i32_u(v1159)
	goto L287
L287:
	;
	if base.Ui64(int64(99)) < base.Ui64(v1177) {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	v1197 = v1159
	v1198 = v1191
	goto L279
L289:
	;
	if base.Ui64(int64(999)) < base.Ui64(v1177) {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	v1197 = v1159
	v1198 = v1166 + int32(1)
	goto L279
L291:
	;
	if base.Ui64(int64(9999)) < base.Ui64(v1177) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	v1197 = v1159
	v1198 = v1166 + int32(2)
	goto L279
L293:
	;
	v1191 = v1166 + int32(4)
	v1195 = base.I64_div_u_s(v1177, int64(10000))
	if base.Ui64(int64(99999)) < base.Ui64(v1177) {
		v1166 = v1191
		v1177 = v1195
		goto L287
	} else {
		goto L295
	}
L294:
	;
	v1197 = v1159
	v1198 = v1166 + int32(3)
	goto L279
L295:
	;
	goto L288
L296:
	;
	v1224 = v41
	v1226 = v39
	v1227 = v40
	goto L9
L297:
	;
	v1281 = v1224 + v1269
	v1284 = m.G4
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	v1286 = m.T0[v1285].(func(*base.Module, int32) int32)(m, v1281+int32(4))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L4
	} else {
		goto L309
	}
L298:
	;
	v1237 = int32(1)
	v1248 = base.I64_extend_i32_s(v1227)
	goto L299
L299:
	;
	if base.Ui64(int64(99)) < base.Ui64(v1248) {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	v1269 = v1262
	goto L297
L301:
	;
	if base.Ui64(int64(999)) < base.Ui64(v1248) {
		goto L303
	} else {
		goto L304
	}
L302:
	;
	v1269 = v1237 + int32(1)
	goto L297
L303:
	;
	if base.Ui64(int64(9999)) < base.Ui64(v1248) {
		goto L305
	} else {
		goto L306
	}
L304:
	;
	v1269 = v1237 + int32(2)
	goto L297
L305:
	;
	v1262 = v1237 + int32(4)
	v1266 = base.I64_div_u_s(v1248, int64(10000))
	if base.Ui64(int64(99999)) < base.Ui64(v1248) {
		v1237 = v1262
		v1248 = v1266
		goto L299
	} else {
		goto L307
	}
L306:
	;
	v1269 = v1237 + int32(3)
	goto L297
L307:
	;
	goto L300
L308:
	;
	v1291 = v1281 + int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v1227
	v1293 = m.G3
	v1298 = F_siprintf(m, v1286, v1293+int32(_a1915), v17+int32(16))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L4
	} else {
		goto L311
	}
L309:
	;
	if v1286 != 0 {
		goto L308
	} else {
		goto L310
	}
L310:
	;
	v1452 = int32(-1)
	v1453 = int32(0)
	v1454 = v1226
	v1455 = v1227
	goto L7
L311:
	;
	if v1227 < int32(1) {
		v1419 = v1298
		goto L312
	} else {
		goto L313
	}
L312:
	;
	if v1419 != v1291 {
		goto L1
	} else {
		goto L340
	}
L313:
	;
	v1304 = int32(0)
	v1312 = v1298
	goto L314
L314:
	;
	v1317 = int32(0)
	v1321 = v1226 + v1304<<(uint(int32(2))%32)
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1321)))
	v1325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322+int32(-1)))))
	switch v1325 & int32(7) {
	case 0:
		goto L321
	case 1:
		goto L320
	case 2:
		goto L319
	case 3:
		goto L318
	case 4:
		goto L317
	default:
		v1342 = v1317
		goto L316
	}
L315:
	;
	v1419 = v1406
	goto L312
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v1342
	v1345 = m.G3
	v1348 = F_siprintf(m, v1286+v1312, v1345+int32(_a1916), v17)
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L4
	} else {
		goto L322
	}
L317:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1322+int32(-17))))
	v1342 = v1341
	goto L316
L318:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1322+int32(-9))))
	v1342 = v1338
	goto L316
L319:
	;
	v1335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1322+int32(-5)))))
	v1342 = v1335
	goto L316
L320:
	;
	v1332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322+int32(-3)))))
	v1342 = v1332
	goto L316
L321:
	;
	v1342 = int32(base.Ui32(v1325) >> (uint(int32(3)) % 32))
	goto L316
L322:
	;
	v1350 = v1348 + v1312
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v1321)))
	v1355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1352+int32(-1)))))
	switch v1355 & int32(7) {
	case 0:
		goto L328
	case 1:
		goto L327
	case 2:
		goto L326
	case 3:
		goto L325
	case 4:
		goto L324
	default:
		v1372 = v1317
		goto L323
	}
L323:
	;
	if v1372 == int32(0) {
		goto L330
	} else {
		goto L331
	}
L324:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v1352+int32(-17))))
	v1372 = v1371
	goto L323
L325:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1352+int32(-9))))
	v1372 = v1368
	goto L323
L326:
	;
	v1365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1352+int32(-5)))))
	v1372 = v1365
	goto L323
L327:
	;
	v1362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1352+int32(-3)))))
	v1372 = v1362
	goto L323
L328:
	;
	v1372 = int32(base.Ui32(v1355) >> (uint(int32(3)) % 32))
	goto L323
L329:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1321)))
	v1381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1378+int32(-1)))))
	switch v1381 & int32(7) {
	case 0:
		goto L337
	case 1:
		goto L336
	case 2:
		goto L335
	case 3:
		goto L334
	case 4:
		goto L333
	default:
		v1398 = int32(0)
		goto L332
	}
L330:
	;
	goto L329
L331:
	;
	v1375 = F__emscripten_memcpy_bulkmem(m, v1286+v1350, v1352, v1372)
	mBase = m.M
	goto L330
L332:
	;
	F_sdsfree(m, v1378)
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L4
	} else {
		goto L338
	}
L333:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1378+int32(-17))))
	v1398 = v1397
	goto L332
L334:
	;
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1378+int32(-9))))
	v1398 = v1394
	goto L332
L335:
	;
	v1391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1378+int32(-5)))))
	v1398 = v1391
	goto L332
L336:
	;
	v1388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1378+int32(-3)))))
	v1398 = v1388
	goto L332
L337:
	;
	v1398 = int32(base.Ui32(v1381) >> (uint(int32(3)) % 32))
	goto L332
L338:
	;
	v1401 = v1398 + v1350
	v1403 = int32(2573)
	*(*uint16)(unsafe.Add(mBase, uint32(v1286+v1401))) = uint16(v1403)
	v1406 = v1401 + int32(2)
	v1408 = v1304 + int32(1)
	if v1408 != v1227 {
		v1304 = v1408
		v1312 = v1406
		goto L314
	} else {
		goto L339
	}
L339:
	;
	goto L315
L340:
	;
	v1426 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1286+v1291))) = uint8(v1426)
	v1428 = m.G4
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1428)+16))
	m.T0[v1429].(func(*base.Module, int32))(m, v1226)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L4
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1286
	v1531 = v1291
	goto L2
L342:
	;
	F_sdsfree(m, v1453)
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L4
	} else {
		goto L351
	}
L343:
	;
	if v1455 == int32(0) {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1502 = m.G4
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1502)+16))
	m.T0[v1503].(func(*base.Module, int32))(m, v1454)
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L4
	} else {
		goto L350
	}
L345:
	;
	v1473 = v1455
	goto L346
L346:
	;
	v1481 = v1473 + int32(-1)
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1454+v1481<<(uint(int32(2))%32))))
	F_sdsfree(m, v1485)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L4
	} else {
		goto L348
	}
L347:
	;
	goto L344
L348:
	;
	if v1481 != 0 {
		v1473 = v1481
		goto L346
	} else {
		goto L349
	}
L349:
	;
	goto L347
L350:
	;
	goto L342
L351:
	;
	v1523 = m.G4
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v1523)+16))
	m.T0[v1524].(func(*base.Module, int32))(m, int32(0))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L4
	} else {
		goto L352
	}
L352:
	;
	v1531 = v1452
	goto L2
}
func F_vdprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v5 = m.G0
	v6 = int32(144)
	v7 = v5 - v6
	m.G0 = v7
	v12 = F__emscripten_memset_bulkmem(m, v7, base.I32_extend8_s(int32(0)), v6)
	mBase = m.M
	v13 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(1376)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v13
	v21 = F_vfprintf(m, v12, l1, l2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		m.G0 = v12 + int32(144)
		return v21
	}
}
func F_vectorCleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v2 == int32(0) {
		return
	} else {
		F_valkey_free(m, v2)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_vectorLen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	return v2
}
func F_verifyDumpPayload(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(-1)
	if base.Ui32(l1) < base.Ui32(int32(10)) {
		v65 = v13
	} else {
		v16 = l0 + l1
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-10)))))
		if l2 == int32(0) {
		} else {
			*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v19)
		}
		v24 = v19 & int32(65535)
		v25 = int32(0)
		if v24 < int32(1) {
			v46 = v25
		} else {
			if v25&base.B2i32(v24 < int32(80)) != 0 {
				v46 = v25
			} else {
				if base.B2i32(int32(79) < v24)&v25 != 0 {
					v46 = v25
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, _consts[127]))
					if v38 != 0 {
						v46 = int32(1)
					} else {
						if base.Ui32(int32(80)) < base.Ui32(v24) {
							v46 = v25
						} else {
							if base.Ui32(v24+int32(-12)) < base.Ui32(int32(68)) {
								v46 = v25
							} else {
								v46 = int32(1)
							}
						}
					}
				}
			}
		}
		if v46 == int32(0) {
			v65 = v13
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, _consts[128]))
			if v51 != 0 {
				v65 = int32(0)
			} else {
				v53 = int32(-8)
				v56 = F_crc64(m, int64(0), l0, base.I64_extend_i32_u(l1+v53))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v56
				v62 = *(*int64)(unsafe.Add(mBase, uint32(v16+v53)))
				if v56 != v62 {
					v64 = int32(-1)
				} else {
					v64 = int32(0)
				}
				v65 = v64
			}
		}
	}
	m.G0 = v11 + int32(16)
	return v65
}
func F_verifyGossipSectionNodeIds(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	if l1 == v3 {
		v122 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v122
L2:
	;
	v20 = v3
	v21 = int32(0)
	goto L3
L3:
	;
	v27 = l0 + v21*int32(104)
	goto L8
L4:
	;
	v122 = v111
	goto L1
L5:
	;
	v117 = v21 + int32(1)
	if v117 != l1 {
		v20 = v111
		v21 = v117
		goto L3
	} else {
		goto L25
	}
L6:
	;
	if int32(0)-v51 == int32(0) {
		v111 = v20
		goto L5
	} else {
		goto L14
	}
L7:
	;
	goto L6
L8:
	;
	v36 = int32(0)
	goto L10
L9:
	;
	goto L7
L10:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v36))))
	v41 = int32(255)
	v51 = base.B2i32(base.Ui32((v38+int32(-123))&v41) < base.Ui32(int32(230))) & base.B2i32(base.Ui32((v38+int32(-58))&v41) < base.Ui32(int32(246)))
	if v51 != 0 {
		goto L9
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v53 = v36 + int32(1)
	if v53 != int32(40) {
		v36 = v53
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v66 = F_valkey_malloc(m, int32(193))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	v76 = int32(0)
	goto L17
L17:
	;
	v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0+v76))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v79
	v88 = F_snprintf(m, v66+v76<<(uint(int32(2))%32), int32(5), int32(_a271), v11+int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L15
	} else {
		goto L19
	}
L18:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v95 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v91 = v76 + int32(1)
	if v91 != int32(48) {
		v76 = v91
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	F_valkey_free(m, v66)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L15
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v27
	F__serverLog(m, int32(3), int32(_a272), v11)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v111 = v20 + int32(1)
	goto L5
L25:
	;
	goto L4
}
func F_vfprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F___vfprintf_internal(m, l0, l1, l2, int32(1387), int32(1388))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_vfscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v213 int32
	_ = v213
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v288 int32
	_ = v288
	var v290 int64
	_ = v290
	var v293 int32
	_ = v293
	var v299 int64
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v546 int32
	_ = v546
	var v547 int64
	_ = v547
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int64
	_ = v554
	var v556 int32
	_ = v556
	var v568 int64
	_ = v568
	var v570 int32
	_ = v570
	var v582 int64
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int64
	_ = v611
	var v614 int32
	_ = v614
	var v622 int32
	_ = v622
	var v636 int32
	_ = v636
	var v637 int64
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v647 int64
	_ = v647
	var v648 int64
	_ = v648
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v792 int32
	_ = v792
	var v804 int32
	_ = v804
	var v814 int32
	_ = v814
	var v830 int32
	_ = v830
	var v833 int64
	_ = v833
	var v834 int32
	_ = v834
	var v835 int64
	_ = v835
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v856 float32
	_ = v856
	var v858 float64
	_ = v858
	var v867 int32
	_ = v867
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v926 int32
	_ = v926
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1006 int32
	_ = v1006
	var v1013 int32
	_ = v1013
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1076 int32
	_ = v1076
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1106 int32
	_ = v1106
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1217 int32
	_ = v1217
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1350 int32
	_ = v1350
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1420 int32
	_ = v1420
	var v1421 int64
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1428 int64
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1432 int64
	_ = v1432
	var v1453 int32
	_ = v1453
	var v1460 int32
	_ = v1460
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1477 int64
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1489 int32
	_ = v1489
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1495 int64
	_ = v1495
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1511 int32
	_ = v1511
	var v1515 int32
	_ = v1515
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1540 int32
	_ = v1540
	var v1547 int32
	_ = v1547
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1622 int32
	_ = v1622
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1633 int32
	_ = v1633
	var __phi1633 int32
	_ = __phi1633
	var v1634 int32
	_ = v1634
	var __phi1634 int32
	_ = __phi1634
	var v1640 int32
	_ = v1640
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1665 int32
	_ = v1665
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1689 int32
	_ = v1689
	var v1696 int32
	_ = v1696
	var v1701 int32
	_ = v1701
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1716 int32
	_ = v1716
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1784 int32
	_ = v1784
	var v1789 int32
	_ = v1789
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1800 int32
	_ = v1800
	var __phi1800 int32
	_ = __phi1800
	var v1801 int32
	_ = v1801
	var __phi1801 int32
	_ = __phi1801
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1824 int32
	_ = v1824
	var v1833 int32
	_ = v1833
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1848 int32
	_ = v1848
	var v1855 int32
	_ = v1855
	var v1860 int32
	_ = v1860
	var v1882 int32
	_ = v1882
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1898 int32
	_ = v1898
	var v1902 int32
	_ = v1902
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1919 int32
	_ = v1919
	var v1922 int32
	_ = v1922
	var v1929 int32
	_ = v1929
	var v1934 int32
	_ = v1934
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1960 int32
	_ = v1960
	var v1963 int32
	_ = v1963
	var v1967 int32
	_ = v1967
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1984 int32
	_ = v1984
	var v1986 int32
	_ = v1986
	var v1989 int32
	_ = v1989
	var v1993 int32
	_ = v1993
	var v1999 int32
	_ = v1999
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2013 int32
	_ = v2013
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2019 int32
	_ = v2019
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2053 int32
	_ = v2053
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2075 int32
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2079 int32
	_ = v2079
	var v2086 int32
	_ = v2086
	var v2088 int32
	_ = v2088
	var v2091 int32
	_ = v2091
	var v2096 int32
	_ = v2096
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2107 int32
	_ = v2107
	var __phi2107 int32
	_ = __phi2107
	var v2108 int32
	_ = v2108
	var __phi2108 int32
	_ = __phi2108
	var v2114 int32
	_ = v2114
	var v2117 int32
	_ = v2117
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2139 int32
	_ = v2139
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2153 int32
	_ = v2153
	var v2156 int32
	_ = v2156
	var v2158 int32
	_ = v2158
	var v2163 int32
	_ = v2163
	var v2170 int32
	_ = v2170
	var v2175 int32
	_ = v2175
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2190 int32
	_ = v2190
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2211 int32
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2220 int32
	_ = v2220
	var v2222 int32
	_ = v2222
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2240 int32
	_ = v2240
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2253 int32
	_ = v2253
	var v2255 int32
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2263 int32
	_ = v2263
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2274 int32
	_ = v2274
	var __phi2274 int32
	_ = __phi2274
	var v2275 int32
	_ = v2275
	var __phi2275 int32
	_ = __phi2275
	var v2281 int32
	_ = v2281
	var v2284 int32
	_ = v2284
	var v2298 int32
	_ = v2298
	var v2307 int32
	_ = v2307
	var v2309 int32
	_ = v2309
	var v2312 int32
	_ = v2312
	var v2315 int32
	_ = v2315
	var v2317 int32
	_ = v2317
	var v2322 int32
	_ = v2322
	var v2329 int32
	_ = v2329
	var v2334 int32
	_ = v2334
	var v2356 int32
	_ = v2356
	var v2368 int32
	_ = v2368
	var v2370 int32
	_ = v2370
	var v2372 int32
	_ = v2372
	var v2376 int32
	_ = v2376
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2393 int32
	_ = v2393
	var v2396 int32
	_ = v2396
	var v2403 int32
	_ = v2403
	var v2408 int32
	_ = v2408
	var v2412 int32
	_ = v2412
	var v2414 int32
	_ = v2414
	var v2429 int32
	_ = v2429
	var v2431 int32
	_ = v2431
	var v2434 int32
	_ = v2434
	var v2437 int32
	_ = v2437
	var v2441 int32
	_ = v2441
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2458 int32
	_ = v2458
	var v2460 int32
	_ = v2460
	var v2463 int32
	_ = v2463
	var v2467 int32
	_ = v2467
	var v2473 int32
	_ = v2473
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2487 int32
	_ = v2487
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2493 int32
	_ = v2493
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2542 int32
	_ = v2542
	v4 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(304)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v4 <= v27 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L4
L3:
	;
	v34 = int32(1)
	goto L1
L4:
	;
	v34 = int32(0)
	goto L1
L5:
	;
	if v2540 != 0 {
		goto L561
	} else {
		goto L562
	}
L6:
	;
	v2539 = v2516
	v2540 = v2517
	v2542 = int32(-1)
	goto L5
L7:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v36 = F___toread(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v40 == int32(0) {
		v2516 = v25
		v2517 = v34
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	v50 = l1
	v51 = l2
	v54 = v43
	v55 = int32(0)
	v57 = int64(0)
	v62 = v4
	v63 = v4
	goto L17
L13:
	;
	v2539 = v25
	v2540 = v34
	v2542 = int32(0)
	goto L5
L14:
	;
	if v1552 == int32(0) {
		v2539 = v25
		v2540 = v34
		v2542 = v1547
		goto L5
	} else {
		goto L360
	}
L15:
	;
	if v55 != 0 {
		goto L357
	} else {
		goto L358
	}
L16:
	;
	v1515 = int32(0)
	v1528 = int32(1)
	v1530 = v1515
	v1531 = v1515
	goto L15
L17:
	;
	v72 = v54 & int32(255)
	goto L21
L19:
	;
	v1511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1492)+1)))
	if v1511 != 0 {
		v50 = v1492 + int32(1)
		v51 = v1489
		v54 = v1511
		v55 = v1493
		v57 = v1495
		v62 = v1500
		v63 = v1501
		goto L17
	} else {
		goto L356
	}
L20:
	;
	if v72 != int32(37) {
		goto L44
	} else {
		goto L45
	}
L21:
	;
	if base.B2i32(v72 == int32(32))|base.B2i32(base.Ui32(v72+int32(-9)) < base.Ui32(int32(5))) == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v83 = v50
	goto L23
L23:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	goto L25
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = base.I64_extend_i32_s(v118 - v119)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L28
L25:
	;
	if base.B2i32(v106 == int32(32))|base.B2i32(base.Ui32(v106+int32(-9)) < base.Ui32(int32(5))) != 0 {
		v83 = v83 + int32(1)
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	goto L31
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v123
	goto L27
L31:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v155 == v156 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v173 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v173 < int64(0) {
		v179 = v172
		goto L39
	} else {
		goto L40
	}
L33:
	;
	goto L37
L34:
	;
	v162 = F___shgetc(m, l0)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v155 + int32(1)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	v164 = v161
	goto L33
L36:
	;
	v164 = v162
	goto L33
L37:
	;
	if base.B2i32(v164 == int32(32))|base.B2i32(base.Ui32(v164+int32(-9)) < base.Ui32(int32(5))) != 0 {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	goto L32
L39:
	;
	v180 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1489 = v51
	v1492 = v83
	v1493 = v55
	v1495 = v180 + v57 + base.I64_extend_i32_s(v179-v182)
	v1500 = v62
	v1501 = v63
	goto L19
L40:
	;
	v177 = v172 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v177
	v179 = v177
	goto L39
L41:
	;
	v344 = int32(0)
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if base.Ui32(int32(9)) < base.Ui32((v346+int32(-48))&int32(255)) {
		v392 = v346
		v396 = v341
		v401 = v344
		goto L79
	} else {
		goto L80
	}
L42:
	;
	v310 = v188 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v310) {
		goto L72
	} else {
		goto L73
	}
L43:
	;
	v340 = v51
	v341 = v50 + int32(2)
	v342 = int32(0)
	goto L41
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = base.I64_extend_i32_s(v198 - v199)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L49
L45:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v188 == int32(42) {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	if v188 != int32(37) {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v213 != int32(37) {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v203
	goto L48
L52:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	if v271 == v288 {
		goto L66
	} else {
		goto L67
	}
L53:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v257 == v258 {
		goto L63
	} else {
		goto L64
	}
L54:
	;
	goto L55
L55:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v238 == v239 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v267 = v50 + int32(1)
	v271 = v247
	goto L52
L57:
	;
	goto L61
L58:
	;
	v245 = F___shgetc(m, l0)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L9
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v238 + int32(1)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	v247 = v244
	goto L57
L60:
	;
	v247 = v245
	goto L57
L61:
	;
	if base.B2i32(v247 == int32(32))|base.B2i32(base.Ui32(v247+int32(-9)) < base.Ui32(int32(5))) != 0 {
		goto L55
	} else {
		goto L62
	}
L62:
	;
	goto L56
L63:
	;
	v264 = F___shgetc(m, l0)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L9
	} else {
		goto L65
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v257 + int32(1)
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
	v267 = v50
	v271 = v263
	goto L52
L65:
	;
	v267 = v50
	v271 = v264
	goto L52
L66:
	;
	v299 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1489 = v51
	v1492 = v267
	v1493 = v55
	v1495 = v299 + v57 + base.I64_extend_i32_s(v301-v302)
	v1500 = v62
	v1501 = v63
	goto L19
L67:
	;
	v290 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v290 < int64(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if int32(-1) < v271 {
		v2539 = v25
		v2540 = v34
		v2542 = v55
		goto L5
	} else {
		goto L70
	}
L69:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v293 + int32(-1)
	goto L68
L70:
	;
	if v55 != 0 {
		v2539 = v25
		v2540 = v34
		v2542 = v55
		goto L5
	} else {
		goto L71
	}
L71:
	;
	v2516 = v25
	v2517 = v34
	goto L6
L72:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v340 = v51 + int32(4)
	v341 = v50 + int32(1)
	v342 = v337
	goto L41
L73:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+2)))
	if v313 != int32(36) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v319 = m.G0
	v321 = v319 - int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v321)+12)) = v51
	if base.Ui32(int32(1)) < base.Ui32(v310) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v340 = v51
	v341 = v50 + int32(3)
	v342 = v334
	goto L41
L76:
	;
	v330 = v51 + v310<<(uint(int32(2))%32) + int32(-4)
	goto L78
L77:
	;
	v330 = v51
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321)+8)) = v330 + int32(4)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	goto L75
L79:
	;
	if v392&int32(255) == int32(109) {
		goto L85
	} else {
		goto L86
	}
L80:
	;
	v354 = v346
	v358 = v341
	v363 = v344
	goto L81
L81:
	;
	v375 = int32(10)
	v377 = int32(255)
	v380 = int32(-48)
	v381 = v363*v375 + v354&v377 + v380
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+1)))
	v384 = v358 + int32(1)
	if base.Ui32((v382+v380)&v377) < base.Ui32(v375) {
		v354 = v382
		v358 = v384
		v363 = v381
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v392 = v382
	v396 = v384
	v401 = v381
	goto L79
L83:
	;
	goto L82
L84:
	;
	v430 = v426 + int32(1)
	switch v424&int32(255) + int32(-65) {
	case 0, 2, 4, 5, 6, 18, 23, 26, 32, 34, 35, 36, 37, 38, 40, 45, 46, 47, 50, 52, 55:
		goto L88
	default:
		v1528 = v425
		v1530 = v427
		v1531 = v428
		goto L15
	case 11:
		goto L89
	case 39:
		goto L92
	case 41:
		v458 = v430
		v459 = int32(3)
		goto L87
	case 43:
		goto L91
	case 51, 57:
		goto L90
	}
L85:
	;
	v419 = int32(0)
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+1)))
	v424 = v422
	v425 = base.B2i32(v342 != v419)
	v426 = v396 + int32(1)
	v427 = v419
	v428 = v419
	goto L84
L86:
	;
	v424 = v392
	v425 = v344
	v426 = v396
	v427 = v62
	v428 = v63
	goto L84
L87:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	v465 = base.B2i32(v461&int32(47) == int32(3))
	if v461&int32(47) == int32(3) {
		goto L105
	} else {
		goto L106
	}
L88:
	;
	v458 = v426
	v459 = int32(0)
	goto L87
L89:
	;
	v458 = v430
	v459 = int32(2)
	goto L87
L90:
	;
	v458 = v430
	v459 = int32(1)
	goto L87
L91:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426)+1)))
	v449 = base.B2i32(v447 == int32(108))
	if v447 == int32(108) {
		goto L99
	} else {
		goto L100
	}
L92:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426)+1)))
	v440 = base.B2i32(v438 == int32(104))
	if v438 == int32(104) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v441 = v426 + int32(2)
	goto L95
L94:
	;
	v441 = v430
	goto L95
L95:
	;
	if v438 == int32(104) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v444 = int32(-2)
	goto L98
L97:
	;
	v444 = int32(-1)
	goto L98
L98:
	;
	v458 = v441
	v459 = v444
	goto L87
L99:
	;
	v450 = v426 + int32(2)
	goto L101
L100:
	;
	v450 = v430
	goto L101
L101:
	;
	if v447 == int32(108) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v453 = int32(3)
	goto L104
L103:
	;
	v453 = int32(1)
	goto L104
L104:
	;
	v458 = v450
	v459 = v453
	goto L87
L105:
	;
	v466 = int32(1)
	goto L107
L106:
	;
	v466 = v459
	goto L107
L107:
	;
	if v461&int32(47) == int32(3) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v582 = base.I64_extend_i32_s(v570)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v582
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = base.I64_extend_i32_s(v586 - v587)
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v582 == int64(0) {
		v599 = v591
		goto L142
	} else {
		goto L143
	}
L109:
	;
	v469 = v461 | int32(32)
	goto L111
L110:
	;
	v469 = v461
	goto L111
L111:
	;
	if v469 == int32(91) {
		v568 = v57
		v570 = v401
		goto L108
	} else {
		goto L112
	}
L112:
	;
	if v469 == int32(110) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = base.I64_extend_i32_s(v492 - v493)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L128
L114:
	;
	if v342 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L115:
	;
	if v469 != int32(99) {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v476 = int32(1)
	if v476 < v401 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v479 = v401
	goto L119
L118:
	;
	v479 = v476
	goto L119
L119:
	;
	v568 = v57
	v570 = v479
	goto L108
L120:
	;
	v1489 = v340
	v1492 = v458
	v1493 = v55
	v1495 = v57
	v1500 = v427
	v1501 = v428
	goto L19
L121:
	;
	goto L120
L122:
	;
	switch v466 + int32(2) {
	case 0:
		goto L126
	case 1:
		goto L125
	case 2, 3:
		goto L124
	default:
		goto L121
	case 5:
		goto L123
	}
L123:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v342))) = v57
	goto L121
L124:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v342))) = uint32(v57)
	goto L120
L125:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v342))) = uint16(v57)
	goto L120
L126:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v342))) = uint8(v57)
	goto L120
L127:
	;
	goto L131
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v497
	goto L127
L131:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v529 == v530 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v547 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v547 < int64(0) {
		v553 = v546
		goto L139
	} else {
		goto L140
	}
L133:
	;
	goto L137
L134:
	;
	v536 = F___shgetc(m, l0)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L9
	} else {
		goto L136
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v529 + int32(1)
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529))))
	v538 = v535
	goto L133
L136:
	;
	v538 = v536
	goto L133
L137:
	;
	if base.B2i32(v538 == int32(32))|base.B2i32(base.Ui32(v538+int32(-9)) < base.Ui32(int32(5))) != 0 {
		goto L131
	} else {
		goto L138
	}
L138:
	;
	goto L132
L139:
	;
	v554 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v568 = v554 + v57 + base.I64_extend_i32_s(v553-v556)
	v570 = v401
	goto L108
L140:
	;
	v551 = v546 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v551
	v553 = v551
	goto L139
L141:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v601 == v602 {
		goto L146
	} else {
		goto L147
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v599
	goto L141
L143:
	;
	if base.I64_extend_i32_s(v591-v587) <= v582 {
		v599 = v591
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v599 = v587 + base.I32_wrap_i64(v582)
	goto L142
L145:
	;
	v611 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v611 < int64(0) {
		goto L150
	} else {
		goto L151
	}
L146:
	;
	v607 = F___shgetc(m, l0)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L9
	} else {
		goto L148
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v601 + int32(1)
	goto L145
L148:
	;
	if v607 < int32(0) {
		v1528 = v425
		v1530 = v427
		v1531 = v428
		goto L15
	} else {
		goto L149
	}
L149:
	;
	goto L145
L150:
	;
	switch v469 + int32(-88) {
	case 0, 24, 32:
		v830 = int32(16)
		goto L157
	case 1, 2, 4, 5, 6, 7, 8, 10, 16, 18, 19, 20, 21, 22, 25, 26, 28, 30, 31:
		v1460 = v458
		v1468 = v427
		v1469 = v428
		goto L152
	case 3, 11, 27:
		goto L161
	case 9, 13, 14, 15:
		goto L162
	case 12, 29:
		goto L159
	case 17:
		goto L158
	case 23:
		goto L160
	default:
		goto L163
	}
L151:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v614 + int32(-1)
	goto L150
L152:
	;
	v1477 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1489 = v340
	v1492 = v1460
	v1493 = v55 + base.B2i32(v342 != int32(0))
	v1495 = v1477 + v568 + base.I64_extend_i32_s(v1479-v1480)
	v1500 = v1468
	v1501 = v1469
	goto L19
L153:
	;
	v888 = base.B2i32(v469 != int32(99))
	if v469 != int32(99) {
		goto L213
	} else {
		goto L214
	}
L154:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v342))) = v648
	*(*int64)(unsafe.Add(mBase, uint32(v342)+8)) = v647
	v1460 = v458
	v1468 = v427
	v1469 = v428
	goto L152
L155:
	;
	v858 = F___trunctfdf2(m, v648, v647)
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(v342))) = v858
	v1460 = v458
	v1468 = v427
	v1469 = v428
	goto L152
L156:
	;
	v856 = F___trunctfsf2(m, v648, v647)
	mBase = m.M
	*(*float32)(unsafe.Add(mBase, uint32(v342))) = v856
	v1460 = v458
	v1468 = v427
	v1469 = v428
	goto L152
L157:
	;
	v833 = F___intscan(m, l0, v830, int32(0), int64(-1))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L9
	} else {
		goto L201
	}
L158:
	;
	v830 = int32(0)
	goto L157
L159:
	;
	v830 = int32(10)
	goto L157
L160:
	;
	v830 = int32(8)
	goto L157
L161:
	;
	if v469|int32(16) != int32(115) {
		goto L169
	} else {
		goto L170
	}
L162:
	;
	F___floatscan(m, v25+int32(8), l0, v466, int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L9
	} else {
		goto L166
	}
L163:
	;
	v622 = v469 + int32(-65)
	if base.Ui32(int32(6)) < base.Ui32(v622) {
		v1460 = v458
		v1468 = v427
		v1469 = v428
		goto L152
	} else {
		goto L164
	}
L164:
	;
	if int32(1)<<(uint(v622)%32)&int32(113) == int32(0) {
		v1460 = v458
		v1468 = v427
		v1469 = v428
		goto L152
	} else {
		goto L165
	}
L165:
	;
	goto L162
L166:
	;
	v637 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v637 == int64(0)-base.I64_extend_i32_s(v639-v640) {
		v1547 = v55
		v1552 = v425
		v1554 = v427
		v1555 = v428
		goto L14
	} else {
		goto L167
	}
L167:
	;
	if v342 == int32(0) {
		v1460 = v458
		v1468 = v427
		v1469 = v428
		goto L152
	} else {
		goto L168
	}
L168:
	;
	v647 = *(*int64)(unsafe.Add(mBase, uint32(v25+int32(16))))
	v648 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
	switch v466 {
	case 0:
		goto L156
	case 1:
		goto L155
	case 2:
		goto L154
	default:
		v1460 = v458
		v1468 = v427
		v1469 = v428
		goto L152
	}
L169:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458)+1)))
	v673 = base.B2i32(v671 == int32(94))
	v676 = F__emscripten_memset_bulkmem(m, v25+int32(32), base.I32_extend8_s(v673), int32(257))
	mBase = m.M
	goto L173
L170:
	;
	v658 = F__emscripten_memset_bulkmem(m, v25+int32(32), base.I32_extend8_s(int32(-1)), int32(257))
	mBase = m.M
	goto L171
L171:
	;
	v659 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+32)) = uint8(v659)
	if v469 != int32(115) {
		v867 = v458
		goto L153
	} else {
		goto L172
	}
L172:
	;
	v663 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+65)) = uint8(v663)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+46)) = uint8(v663)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+42)) = v663
	v867 = v458
	goto L153
L173:
	;
	v677 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+32)) = uint8(v677)
	if v671 == int32(94) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v683 = v458 + int32(2)
	goto L176
L175:
	;
	v683 = v458 + int32(1)
	goto L176
L176:
	;
	if v671 == int32(94) {
		goto L181
	} else {
		goto L182
	}
L177:
	;
	v711 = v704
	goto L186
L178:
	;
	v704 = v683 + int32(1)
	v705 = v701
	goto L177
L179:
	;
	v699 = base.B2i32(v671 != int32(94))
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+126)) = uint8(v699)
	v701 = v699
	goto L178
L180:
	;
	v696 = base.B2i32(v671 != int32(94))
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+78)) = uint8(v696)
	v701 = v696
	goto L178
L181:
	;
	v686 = int32(2)
	goto L183
L182:
	;
	v686 = int32(1)
	goto L183
L183:
	;
	v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458+v686))))
	if v688 == int32(45) {
		goto L180
	} else {
		goto L184
	}
L184:
	;
	if v688 == int32(93) {
		goto L179
	} else {
		goto L185
	}
L185:
	;
	v704 = v683
	v705 = base.B2i32(v671 != int32(94))
	goto L177
L186:
	;
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711))))
	if v728 == int32(45) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v814+(v25+int32(32)))+1)) = uint8(v705)
	v711 = v804 + int32(1)
	goto L186
L189:
	;
	v735 = int32(45)
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711)+1)))
	if v736 == int32(0) {
		v804 = v711
		v814 = v735
		goto L188
	} else {
		goto L193
	}
L190:
	;
	if v728 == int32(0) {
		v1528 = v425
		v1530 = v427
		v1531 = v428
		goto L15
	} else {
		goto L191
	}
L191:
	;
	if v728 == int32(93) {
		v867 = v711
		goto L153
	} else {
		goto L192
	}
L192:
	;
	v804 = v711
	v814 = v728
	goto L188
L193:
	;
	if v736 == int32(93) {
		v804 = v711
		v814 = v735
		goto L188
	} else {
		goto L194
	}
L194:
	;
	v742 = v711 + int32(1)
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711+int32(-1)))))
	if base.Ui32(v745) < base.Ui32(v736) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v804 = v742
	v814 = v792
	goto L188
L196:
	;
	v748 = v745
	goto L198
L197:
	;
	v792 = v736
	goto L195
L198:
	;
	v772 = v748 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(32)+v772))) = uint8(v705)
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742))))
	if base.Ui32(v772) < base.Ui32(v775) {
		v748 = v772
		goto L198
	} else {
		goto L200
	}
L199:
	;
	v792 = v775
	goto L195
L200:
	;
	goto L199
L201:
	;
	v835 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v835 == int64(0)-base.I64_extend_i32_s(v837-v838) {
		v1547 = v55
		v1552 = v425
		v1554 = v427
		v1555 = v428
		goto L14
	} else {
		goto L202
	}
L202:
	;
	if v469 != int32(112) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	if v342 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	if v342 == int32(0) {
		goto L203
	} else {
		goto L205
	}
L205:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v342))) = uint32(v833)
	v1460 = v458
	v1468 = v427
	v1469 = v428
	goto L152
L206:
	;
	v1460 = v458
	v1468 = v427
	v1469 = v428
	goto L152
L207:
	;
	goto L206
L208:
	;
	switch v466 + int32(2) {
	case 0:
		goto L212
	case 1:
		goto L211
	case 2, 3:
		goto L210
	default:
		goto L207
	case 5:
		goto L209
	}
L209:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v342))) = v833
	goto L207
L210:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v342))) = uint32(v833)
	goto L206
L211:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v342))) = uint16(v833)
	goto L206
L212:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v342))) = uint8(v833)
	goto L206
L213:
	;
	v889 = int32(31)
	goto L215
L214:
	;
	v889 = v570 + int32(1)
	goto L215
L215:
	;
	if v466 != int32(1) {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1421 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v1421 < int64(0) {
		v1427 = v1420
		goto L345
	} else {
		goto L346
	}
L217:
	;
	if v425 == int32(0) {
		goto L295
	} else {
		goto L296
	}
L218:
	;
	if v425 == int32(0) {
		v899 = v342
		goto L219
	} else {
		goto L220
	}
L219:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+296)) = int64(0)
	v904 = int32(0)
	v913 = v899
	v915 = v889
	goto L224
L220:
	;
	v896 = F_emscripten_builtin_malloc(m, v889<<(uint(int32(2))%32))
	mBase = m.M
	if v896 == int32(0) {
		goto L16
	} else {
		goto L221
	}
L221:
	;
	v899 = v896
	goto L219
L222:
	;
	v1528 = v425
	v1530 = v1186
	v1531 = v913
	goto L15
L223:
	;
	v1177 = int32(0)
	v1179 = v25 + int32(296)
	if v1179 != 0 {
		goto L292
	} else {
		goto L293
	}
L224:
	;
	v926 = v904
	goto L226
L225:
	;
	v1528 = int32(1)
	v1530 = int32(0)
	v1531 = v913
	goto L15
L226:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v947 == v948 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	v1125 = int32(1)
	v1128 = v915<<(uint(v1125)%32) | v1125
	v1130 = v1128 << (uint(int32(2)) % 32)
	if v913 != 0 {
		goto L273
	} else {
		goto L274
	}
L228:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v956+(v25+int32(32))+int32(1)))))
	if v962 == int32(0) {
		goto L223
	} else {
		goto L232
	}
L229:
	;
	v954 = F___shgetc(m, l0)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L9
	} else {
		goto L231
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v947 + int32(1)
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v947))))
	v956 = v953
	goto L228
L231:
	;
	v956 = v954
	goto L228
L232:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+27)) = uint8(v956)
	v967 = v25 + int32(28)
	v969 = v25 + int32(27)
	v970 = int32(1)
	v972 = v25 + int32(296)
	if v972 != 0 {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	if v1106 == int32(-2) {
		goto L226
	} else {
		goto L265
	}
L234:
	;
	v978 = v972
	goto L236
L235:
	;
	v978 = int32(9128420)
	goto L236
L236:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v978)))
	if v969 != 0 {
		goto L240
	} else {
		goto L241
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v978))) = v1091
	v1106 = int32(-2)
	goto L233
L238:
	;
	v1106 = v1085
	goto L233
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v978))) = int32(0)
	v1076 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1076))) = int32(25)
	v1085 = int32(-1)
	goto L238
L240:
	;
	goto L243
L241:
	;
	if v979 != 0 {
		goto L239
	} else {
		goto L242
	}
L242:
	;
	v1106 = int32(0)
	goto L233
L243:
	;
	if v979 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	v1023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v969))))
	v1025 = int32(base.Ui32(v1023) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(7)) < base.Ui32(v1025+int32(-16)|(v979>>(uint(int32(26))%32)+v1025)) {
		goto L239
	} else {
		goto L256
	}
L245:
	;
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v969))))
	v987 = base.I32_extend8_s(v986)
	if v987 < int32(0) {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	goto L244
L247:
	;
	v995 = F___get_tp(m)
	mBase = m.M
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v995)+96))
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v996)))
	if v997 != 0 {
		goto L251
	} else {
		goto L252
	}
L248:
	;
	if v967 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1106 = base.B2i32(v987 != int32(0))
	goto L233
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v967))) = v986
	goto L249
L251:
	;
	v1006 = v986 + int32(-194)
	if base.Ui32(int32(50)) < base.Ui32(v1006) {
		goto L239
	} else {
		goto L254
	}
L252:
	;
	if v967 == int32(0) {
		v1085 = int32(1)
		goto L238
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v967))) = v987 & int32(57343)
	v1106 = int32(1)
	goto L233
L254:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v1006<<(uint(int32(2))%32))+uint32(_consts[1046])))
	v1091 = v1013
	goto L237
L256:
	;
	v1035 = v969
	v1037 = v979
	v1039 = v970
	v1040 = v1023
	goto L257
L257:
	;
	v1043 = v1039 + int32(-1)
	v1050 = v1040&int32(255) + int32(-128) | v1037<<(uint(int32(6))%32)
	if v1050 < int32(0) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	goto L239
L259:
	;
	if v1043 == int32(0) {
		v1091 = v1050
		goto L237
	} else {
		goto L263
	}
L260:
	;
	v1053 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v978))) = v1053
	if v967 == v1053 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1106 = v970 - v1043
	goto L233
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v967))) = v1050
	goto L261
L263:
	;
	v1062 = v1035 + int32(1)
	v1063 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1062))))
	if v1063 < int32(-64) {
		v1035 = v1062
		v1037 = v1050
		v1039 = v1043
		v1040 = v1063
		goto L257
	} else {
		goto L264
	}
L264:
	;
	goto L258
L265:
	;
	if v1106 != int32(-1) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	if v913 == int32(0) {
		v1121 = v926
		goto L268
	} else {
		goto L269
	}
L267:
	;
	v1186 = int32(0)
	goto L222
L268:
	;
	if v425 == int32(0) {
		v926 = v1121
		goto L226
	} else {
		goto L270
	}
L269:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v913+v926<<(uint(int32(2))%32)))) = v1117
	v1121 = v926 + int32(1)
	goto L268
L270:
	;
	if v1121 != v915 {
		v926 = v1121
		goto L226
	} else {
		goto L271
	}
L271:
	;
	goto L227
L272:
	;
	if v1174 != 0 {
		v904 = v1121
		v913 = v1174
		v915 = v1128
		goto L224
	} else {
		goto L290
	}
L273:
	;
	if base.Ui32(v1130) < base.Ui32(int32(-64)) {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	v1133 = F_emscripten_builtin_malloc(m, v1130)
	mBase = m.M
	v1174 = v1133
	goto L272
L275:
	;
	v1140 = int32(-8)
	v1143 = int32(11)
	if base.Ui32(v1130) < base.Ui32(v1143) {
		goto L278
	} else {
		goto L279
	}
L276:
	;
	v1136 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1136))) = int32(48)
	v1174 = int32(0)
	goto L272
L277:
	;
	v1155 = F_emscripten_builtin_malloc(m, v1130)
	mBase = m.M
	if v1155 != 0 {
		goto L282
	} else {
		goto L283
	}
L278:
	;
	v1149 = int32(16)
	goto L280
L279:
	;
	v1149 = (v1130 + v1143) & v1140
	goto L280
L280:
	;
	v1150 = F_try_realloc_chunk(m, v913+v1140, v1149)
	mBase = m.M
	if v1150 == int32(0) {
		goto L277
	} else {
		goto L281
	}
L281:
	;
	v1174 = v1150 + int32(8)
	goto L272
L282:
	;
	v1157 = int32(-4)
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v913+v1157)))
	if v1161&int32(3) != 0 {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	v1174 = int32(0)
	goto L272
L284:
	;
	v1164 = v1157
	goto L286
L285:
	;
	v1164 = int32(-8)
	goto L286
L286:
	;
	v1167 = v1164 + v1161&int32(-8)
	if base.Ui32(v1167) < base.Ui32(v1130) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1169 = v1167
	goto L289
L288:
	;
	v1169 = v1130
	goto L289
L289:
	;
	v1170 = F___memcpy(m, v1155, v913, v1169)
	mBase = m.M
	F_emscripten_builtin_free(m, v913)
	mBase = m.M
	v1174 = v1155
	goto L272
L290:
	;
	goto L225
L291:
	;
	if v1184 != 0 {
		v1399 = v926
		v1411 = v1177
		v1412 = v913
		v1413 = v913
		goto L216
	} else {
		goto L294
	}
L292:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1179)))
	v1184 = base.B2i32(v1181 == int32(0))
	goto L291
L293:
	;
	v1184 = int32(1)
	goto L291
L294:
	;
	v1186 = v1177
	goto L222
L295:
	;
	v1310 = int32(0)
	if v342 == v1310 {
		goto L328
	} else {
		goto L329
	}
L296:
	;
	v1190 = int32(0)
	v1191 = F_emscripten_builtin_malloc(m, v889)
	mBase = m.M
	if v1191 == v1190 {
		goto L16
	} else {
		goto L297
	}
L297:
	;
	v1195 = v1190
	v1204 = v1191
	v1206 = v889
	goto L298
L298:
	;
	v1217 = v1195
	goto L300
L299:
	;
	v1528 = int32(1)
	v1530 = v1204
	v1531 = int32(0)
	goto L15
L300:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1238 == v1239 {
		goto L303
	} else {
		goto L304
	}
L301:
	;
	v1260 = int32(1)
	v1263 = v1206<<(uint(v1260)%32) | v1260
	if v1204 != 0 {
		goto L310
	} else {
		goto L311
	}
L302:
	;
	v1253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1247+(v25+int32(32))+int32(1)))))
	if v1253 != 0 {
		goto L306
	} else {
		goto L307
	}
L303:
	;
	v1245 = F___shgetc(m, l0)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L9
	} else {
		goto L305
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1238 + int32(1)
	v1244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1238))))
	v1247 = v1244
	goto L302
L305:
	;
	v1247 = v1245
	goto L302
L306:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1204+v1217))) = uint8(v1247)
	v1258 = v1217 + int32(1)
	if v1258 != v1206 {
		v1217 = v1258
		goto L300
	} else {
		goto L308
	}
L307:
	;
	v1399 = v1217
	v1411 = v1204
	v1412 = int32(0)
	v1413 = v1204
	goto L216
L308:
	;
	goto L301
L309:
	;
	if v1307 != 0 {
		v1195 = v1258
		v1204 = v1307
		v1206 = v1263
		goto L298
	} else {
		goto L327
	}
L310:
	;
	if base.Ui32(v1263) < base.Ui32(int32(-64)) {
		goto L312
	} else {
		goto L313
	}
L311:
	;
	v1266 = F_emscripten_builtin_malloc(m, v1263)
	mBase = m.M
	v1307 = v1266
	goto L309
L312:
	;
	v1273 = int32(-8)
	v1276 = int32(11)
	if base.Ui32(v1263) < base.Ui32(v1276) {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	v1269 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1269))) = int32(48)
	v1307 = int32(0)
	goto L309
L314:
	;
	v1288 = F_emscripten_builtin_malloc(m, v1263)
	mBase = m.M
	if v1288 != 0 {
		goto L319
	} else {
		goto L320
	}
L315:
	;
	v1282 = int32(16)
	goto L317
L316:
	;
	v1282 = (v1263 + v1276) & v1273
	goto L317
L317:
	;
	v1283 = F_try_realloc_chunk(m, v1204+v1273, v1282)
	mBase = m.M
	if v1283 == int32(0) {
		goto L314
	} else {
		goto L318
	}
L318:
	;
	v1307 = v1283 + int32(8)
	goto L309
L319:
	;
	v1290 = int32(-4)
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1204+v1290)))
	if v1294&int32(3) != 0 {
		goto L321
	} else {
		goto L322
	}
L320:
	;
	v1307 = int32(0)
	goto L309
L321:
	;
	v1297 = v1290
	goto L323
L322:
	;
	v1297 = int32(-8)
	goto L323
L323:
	;
	v1300 = v1297 + v1294&int32(-8)
	if base.Ui32(v1300) < base.Ui32(v1263) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1302 = v1300
	goto L326
L325:
	;
	v1302 = v1263
	goto L326
L326:
	;
	v1303 = F___memcpy(m, v1288, v1204, v1302)
	mBase = m.M
	F_emscripten_builtin_free(m, v1204)
	mBase = m.M
	v1307 = v1288
	goto L309
L327:
	;
	goto L299
L328:
	;
	goto L338
L329:
	;
	v1314 = v1310
	goto L330
L330:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1335 == v1336 {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	v1350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1344+(v25+int32(32))+int32(1)))))
	if v1350 != 0 {
		goto L336
	} else {
		goto L337
	}
L333:
	;
	v1342 = F___shgetc(m, l0)
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L9
	} else {
		goto L335
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1335 + int32(1)
	v1341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335))))
	v1344 = v1341
	goto L332
L335:
	;
	v1344 = v1342
	goto L332
L336:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v342+v1314))) = uint8(v1344)
	v1314 = v1314 + int32(1)
	goto L330
L337:
	;
	v1399 = v1314
	v1411 = v342
	v1412 = int32(0)
	v1413 = v342
	goto L216
L338:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1378 == v1379 {
		goto L341
	} else {
		goto L342
	}
L339:
	;
	v1394 = int32(0)
	v1399 = v1394
	v1411 = v1394
	v1412 = v1394
	v1413 = v1394
	goto L216
L340:
	;
	v1393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1387+(v25+int32(32))+int32(1)))))
	if v1393 != 0 {
		goto L338
	} else {
		goto L344
	}
L341:
	;
	v1385 = F___shgetc(m, l0)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L9
	} else {
		goto L343
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1378 + int32(1)
	v1384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1378))))
	v1387 = v1384
	goto L340
L343:
	;
	v1387 = v1385
	goto L340
L344:
	;
	goto L339
L345:
	;
	v1428 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1432 = v1428 + base.I64_extend_i32_s(v1427-v1429)
	if v1432 == int64(0) {
		v1547 = v55
		v1552 = v425
		v1554 = v1411
		v1555 = v1412
		goto L14
	} else {
		goto L347
	}
L346:
	;
	v1425 = v1420 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1425
	v1427 = v1425
	goto L345
L347:
	;
	if v888|base.B2i32(v1432 == v582) == int32(0) {
		v1547 = v55
		v1552 = v425
		v1554 = v1411
		v1555 = v1412
		goto L14
	} else {
		goto L348
	}
L348:
	;
	if v425 == int32(0) {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	if v469 == int32(99) {
		v1460 = v867
		v1468 = v1411
		v1469 = v1412
		goto L152
	} else {
		goto L351
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v342))) = v1413
	goto L349
L351:
	;
	if v1412 == int32(0) {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	if v1411 != 0 {
		goto L354
	} else {
		goto L355
	}
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1412+v1399<<(uint(int32(2))%32)))) = int32(0)
	goto L352
L354:
	;
	v1453 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1411+v1399))) = uint8(v1453)
	v1460 = v867
	v1468 = v1411
	v1469 = v1412
	goto L152
L355:
	;
	v1460 = v867
	v1468 = int32(0)
	v1469 = v1412
	goto L152
L356:
	;
	v2539 = v25
	v2540 = v34
	v2542 = v1493
	goto L5
L357:
	;
	v1540 = v55
	goto L359
L358:
	;
	v1540 = int32(-1)
	goto L359
L359:
	;
	v1547 = v1540
	v1552 = v1528
	v1554 = v1530
	v1555 = v1531
	goto L14
L360:
	;
	if v1554 == int32(0) {
		goto L362
	} else {
		goto L363
	}
L361:
	;
	if v1555 == int32(0) {
		goto L462
	} else {
		goto L463
	}
L362:
	;
	goto L361
L363:
	;
	v1575 = int32(-8)
	v1576 = v1554 + v1575
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1554+int32(-4))))
	v1581 = v1579 & v1575
	v1582 = v1576 + v1581
	if v1579&int32(1) != 0 {
		v1706 = v1581
		v1707 = v1576
		goto L364
	} else {
		goto L365
	}
L364:
	;
	if base.Ui32(v1582) <= base.Ui32(v1707) {
		goto L362
	} else {
		goto L399
	}
L365:
	;
	if v1579&int32(2) == int32(0) {
		goto L362
	} else {
		goto L366
	}
L366:
	;
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1576)))
	v1590 = v1576 - v1589
	v1592 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	if base.Ui32(v1590) < base.Ui32(v1592) {
		goto L362
	} else {
		goto L367
	}
L367:
	;
	v1594 = v1589 + v1581
	v1596 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v1590 == v1596 {
		goto L371
	} else {
		goto L372
	}
L368:
	;
	if v1612 == int32(0) {
		v1706 = v1594
		v1707 = v1590
		goto L364
	} else {
		goto L387
	}
L369:
	;
	v1665 = int32(0)
	goto L368
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1601)+12)) = v1598
	*(*int32)(unsafe.Add(mBase, uint32(v1598)+8)) = v1601
	v1706 = v1594
	v1707 = v1590
	goto L364
L371:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+4))
	v1647 = int32(3)
	if v1646&v1647 != v1647 {
		v1706 = v1594
		v1707 = v1590
		goto L364
	} else {
		goto L386
	}
L372:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v1590)+12))
	if base.Ui32(int32(255)) < base.Ui32(v1589) {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1590)+24))
	if v1598 == v1590 {
		goto L376
	} else {
		goto L377
	}
L374:
	;
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v1590)+8))
	if v1598 != v1601 {
		goto L370
	} else {
		goto L375
	}
L375:
	;
	v1603 = int32(0)
	v1605 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v1605 & base.I32_rotl(int32(-2), int32(base.Ui32(v1589)>>(uint(int32(3))%32)))
	v1706 = v1594
	v1707 = v1590
	goto L364
L376:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1590)+20))
	if v1617 == int32(0) {
		goto L379
	} else {
		goto L380
	}
L377:
	;
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1590)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1614)+12)) = v1598
	*(*int32)(unsafe.Add(mBase, uint32(v1598)+8)) = v1614
	v1665 = v1598
	goto L368
L378:
	;
	__phi1633 = v1627
	__phi1634 = v1628
	v1633 = __phi1633
	v1634 = __phi1634
	goto L382
L379:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1590)+16))
	if v1622 == int32(0) {
		goto L369
	} else {
		goto L381
	}
L380:
	;
	v1627 = v1617
	v1628 = v1590 + int32(20)
	goto L378
L381:
	;
	v1627 = v1622
	v1628 = v1590 + int32(16)
	goto L378
L382:
	;
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1633)+20))
	if v1640 != 0 {
		__phi1633 = v1640
		__phi1634 = v1633 + int32(20)
		v1633 = __phi1633
		v1634 = __phi1634
		goto L382
	} else {
		goto L384
	}
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1634))) = int32(0)
	v1665 = v1633
	goto L368
L384:
	;
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1633)+16))
	if v1643 != 0 {
		__phi1633 = v1643
		__phi1634 = v1633 + int32(16)
		v1633 = __phi1633
		v1634 = __phi1634
		goto L382
	} else {
		goto L385
	}
L385:
	;
	goto L383
L386:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v1594
	*(*int32)(unsafe.Add(mBase, uint32(v1582)+4)) = v1646 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v1590)+4)) = v1594 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1582))) = v1594
	goto L361
L387:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1590)+28))
	v1676 = v1674 << (uint(int32(2)) % 32)
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1676)+uint32(_consts[519])))
	if v1590 != v1679 {
		goto L389
	} else {
		goto L390
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1665)+24)) = v1612
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v1590)+16))
	if v1696 == int32(0) {
		goto L396
	} else {
		goto L397
	}
L389:
	;
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1612)+16))
	if v1689 != v1590 {
		goto L393
	} else {
		goto L394
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1676)+uint32(_consts[519]))) = v1665
	if v1665 != 0 {
		goto L388
	} else {
		goto L391
	}
L391:
	;
	v1682 = int32(0)
	v1684 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v1684 & base.I32_rotl(int32(-2), v1674)
	v1706 = v1594
	v1707 = v1590
	goto L364
L392:
	;
	if v1665 == int32(0) {
		v1706 = v1594
		v1707 = v1590
		goto L364
	} else {
		goto L395
	}
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1612)+20)) = v1665
	goto L392
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1612)+16)) = v1665
	goto L392
L395:
	;
	goto L388
L396:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v1590)+20))
	if v1701 == int32(0) {
		v1706 = v1594
		v1707 = v1590
		goto L364
	} else {
		goto L398
	}
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1665)+16)) = v1696
	*(*int32)(unsafe.Add(mBase, uint32(v1696)+24)) = v1665
	goto L396
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1665)+20)) = v1701
	*(*int32)(unsafe.Add(mBase, uint32(v1701)+24)) = v1665
	v1706 = v1594
	v1707 = v1590
	goto L364
L399:
	;
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+4))
	if v1716&int32(1) == int32(0) {
		goto L362
	} else {
		goto L400
	}
L400:
	;
	if v1716&int32(2) != 0 {
		goto L405
	} else {
		goto L406
	}
L401:
	;
	if base.Ui32(int32(255)) < base.Ui32(v1882) {
		goto L439
	} else {
		goto L440
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1707)+4)) = v1762 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1707+v1762))) = v1762
	if v1707 != v1746 {
		v1882 = v1762
		goto L401
	} else {
		goto L438
	}
L403:
	;
	if v1779 == int32(0) {
		goto L402
	} else {
		goto L426
	}
L404:
	;
	v1824 = int32(0)
	goto L403
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1582)+4)) = v1716 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v1707)+4)) = v1706 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1707+v1706))) = v1706
	v1882 = v1706
	goto L401
L406:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	if v1582 != v1724 {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v1582 != v1746 {
		goto L410
	} else {
		goto L411
	}
L408:
	;
	v1726 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v1707
	v1730 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v1731 = v1730 + v1706
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v1731
	*(*int32)(unsafe.Add(mBase, uint32(v1707)+4)) = v1731 | int32(1)
	v1737 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v1707 != v1737 {
		goto L362
	} else {
		goto L409
	}
L409:
	;
	v1739 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v1739
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v1739
	goto L361
L410:
	;
	v1762 = v1716&int32(-8) + v1706
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+12))
	if base.Ui32(int32(255)) < base.Ui32(v1716) {
		goto L412
	} else {
		goto L413
	}
L411:
	;
	v1748 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v1707
	v1752 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	v1753 = v1752 + v1706
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v1753
	*(*int32)(unsafe.Add(mBase, uint32(v1707)+4)) = v1753 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1707+v1753))) = v1753
	goto L361
L412:
	;
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+24))
	if v1763 == v1582 {
		goto L416
	} else {
		goto L417
	}
L413:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+8))
	if v1763 != v1766 {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1766)+12)) = v1763
	*(*int32)(unsafe.Add(mBase, uint32(v1763)+8)) = v1766
	goto L402
L415:
	;
	v1768 = int32(0)
	v1770 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v1770 & base.I32_rotl(int32(-2), int32(base.Ui32(v1716)>>(uint(int32(3))%32)))
	goto L402
L416:
	;
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+20))
	if v1784 == int32(0) {
		goto L419
	} else {
		goto L420
	}
L417:
	;
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+12)) = v1763
	*(*int32)(unsafe.Add(mBase, uint32(v1763)+8)) = v1781
	v1824 = v1763
	goto L403
L418:
	;
	__phi1800 = v1794
	__phi1801 = v1795
	v1800 = __phi1800
	v1801 = __phi1801
	goto L422
L419:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+16))
	if v1789 == int32(0) {
		goto L404
	} else {
		goto L421
	}
L420:
	;
	v1794 = v1784
	v1795 = v1582 + int32(20)
	goto L418
L421:
	;
	v1794 = v1789
	v1795 = v1582 + int32(16)
	goto L418
L422:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+20))
	if v1807 != 0 {
		__phi1800 = v1807
		__phi1801 = v1800 + int32(20)
		v1800 = __phi1800
		v1801 = __phi1801
		goto L422
	} else {
		goto L424
	}
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1801))) = int32(0)
	v1824 = v1800
	goto L403
L424:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+16))
	if v1810 != 0 {
		__phi1800 = v1810
		__phi1801 = v1800 + int32(16)
		v1800 = __phi1800
		v1801 = __phi1801
		goto L422
	} else {
		goto L425
	}
L425:
	;
	goto L423
L426:
	;
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+28))
	v1835 = v1833 << (uint(int32(2)) % 32)
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1835)+uint32(_consts[519])))
	if v1582 != v1838 {
		goto L428
	} else {
		goto L429
	}
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1824)+24)) = v1779
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+16))
	if v1855 == int32(0) {
		goto L435
	} else {
		goto L436
	}
L428:
	;
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1779)+16))
	if v1848 != v1582 {
		goto L432
	} else {
		goto L433
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1835)+uint32(_consts[519]))) = v1824
	if v1824 != 0 {
		goto L427
	} else {
		goto L430
	}
L430:
	;
	v1841 = int32(0)
	v1843 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v1843 & base.I32_rotl(int32(-2), v1833)
	goto L402
L431:
	;
	if v1824 == int32(0) {
		goto L402
	} else {
		goto L434
	}
L432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1779)+20)) = v1824
	goto L431
L433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1779)+16)) = v1824
	goto L431
L434:
	;
	goto L427
L435:
	;
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+20))
	if v1860 == int32(0) {
		goto L402
	} else {
		goto L437
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1824)+16)) = v1855
	*(*int32)(unsafe.Add(mBase, uint32(v1855)+24)) = v1824
	goto L435
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1824)+20)) = v1860
	*(*int32)(unsafe.Add(mBase, uint32(v1860)+24)) = v1824
	goto L402
L438:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v1762
	goto L361
L439:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v1882) {
		v1929 = int32(31)
		goto L444
	} else {
		goto L445
	}
L440:
	;
	v1894 = v1882 & int32(-8)
	v1896 = v1894 + int32(9128464)
	v1898 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v1902 = int32(1) << (uint(int32(base.Ui32(v1882)>>(uint(int32(3))%32))) % 32)
	if v1898&v1902 != 0 {
		goto L442
	} else {
		goto L443
	}
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+uint32(_consts[523]))) = v1707
	*(*int32)(unsafe.Add(mBase, uint32(v1908)+12)) = v1707
	*(*int32)(unsafe.Add(mBase, uint32(v1707)+12)) = v1896
	*(*int32)(unsafe.Add(mBase, uint32(v1707)+8)) = v1908
	goto L361
L442:
	;
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1894)+uint32(_consts[523])))
	v1908 = v1907
	goto L441
L443:
	;
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v1898 | v1902
	v1908 = v1896
	goto L441
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1707)+28)) = v1929
	*(*int64)(unsafe.Add(mBase, uint32(v1707)+16)) = int64(0)
	v1934 = v1929 << (uint(int32(2)) % 32)
	v1938 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	v1940 = int32(1) << (uint(v1929) % 32)
	if v1938&v1940 != 0 {
		goto L449
	} else {
		goto L450
	}
L445:
	;
	v1919 = base.I32_clz(int32(base.Ui32(v1882) >> (uint(int32(8)) % 32)))
	v1922 = int32(1)
	v1929 = int32(base.Ui32(v1882)>>(uint(int32(38)-v1919)%32))&v1922 - v1919<<(uint(v1922)%32) + int32(62)
	goto L444
L446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1707+v2001))) = v2004
	*(*int32)(unsafe.Add(mBase, uint32(v1707)+12)) = v2003
	*(*int32)(unsafe.Add(mBase, uint32(v1707+v1999))) = v2002
	v2013 = int32(0)
	v2015 = *(*int32)(unsafe.Add(mBase, _consts[524]))
	v2016 = int32(-1)
	v2017 = v2015 + v2016
	if v2017 != 0 {
		goto L458
	} else {
		goto L459
	}
L447:
	;
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v1963)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1993)+12)) = v1707
	*(*int32)(unsafe.Add(mBase, uint32(v1963)+8)) = v1707
	v1999 = int32(24)
	v2001 = int32(8)
	v2002 = int32(0)
	v2003 = v1963
	v2004 = v1993
	goto L446
L448:
	;
	v1999 = v1984
	v2001 = v1986
	v2002 = v1707
	v2003 = v1707
	v2004 = v1989
	goto L446
L449:
	;
	if v1929 == int32(31) {
		goto L451
	} else {
		goto L452
	}
L450:
	;
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v1938 | v1940
	*(*int32)(unsafe.Add(mBase, uint32(v1934)+uint32(_consts[519]))) = v1707
	v1984 = int32(8)
	v1986 = int32(24)
	v1989 = v1934 + int32(9128728)
	goto L448
L451:
	;
	v1955 = int32(0)
	goto L453
L452:
	;
	v1955 = int32(25) - int32(base.Ui32(v1929)>>(uint(int32(1))%32))
	goto L453
L453:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1934)+uint32(_consts[519])))
	v1960 = v1882 << (uint(v1955) % 32)
	v1963 = v1957
	goto L454
L454:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v1963)+4))
	if v1967&int32(-8) == v1882 {
		goto L447
	} else {
		goto L456
	}
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1977+int32(16)))) = v1707
	v1984 = int32(8)
	v1986 = int32(24)
	v1989 = v1963
	goto L448
L456:
	;
	v1977 = v1963 + int32(base.Ui32(v1960)>>(uint(int32(29))%32))&int32(4)
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v1977)+16))
	if v1978 != 0 {
		v1960 = v1960 << (uint(int32(1)) % 32)
		v1963 = v1978
		goto L454
	} else {
		goto L457
	}
L457:
	;
	goto L455
L458:
	;
	v2019 = v2017
	goto L460
L459:
	;
	v2019 = v2016
	goto L460
L460:
	;
	*(*int32)(unsafe.Add(mBase, _consts[524])) = v2019
	goto L362
L461:
	;
	v2539 = v25
	v2540 = v34
	v2542 = v1547
	goto L5
L462:
	;
	goto L461
L463:
	;
	v2049 = int32(-8)
	v2050 = v1555 + v2049
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v1555+int32(-4))))
	v2055 = v2053 & v2049
	v2056 = v2050 + v2055
	if v2053&int32(1) != 0 {
		v2180 = v2055
		v2181 = v2050
		goto L464
	} else {
		goto L465
	}
L464:
	;
	if base.Ui32(v2056) <= base.Ui32(v2181) {
		goto L462
	} else {
		goto L499
	}
L465:
	;
	if v2053&int32(2) == int32(0) {
		goto L462
	} else {
		goto L466
	}
L466:
	;
	v2063 = *(*int32)(unsafe.Add(mBase, uint32(v2050)))
	v2064 = v2050 - v2063
	v2066 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	if base.Ui32(v2064) < base.Ui32(v2066) {
		goto L462
	} else {
		goto L467
	}
L467:
	;
	v2068 = v2063 + v2055
	v2070 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v2064 == v2070 {
		goto L471
	} else {
		goto L472
	}
L468:
	;
	if v2086 == int32(0) {
		v2180 = v2068
		v2181 = v2064
		goto L464
	} else {
		goto L487
	}
L469:
	;
	v2139 = int32(0)
	goto L468
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2075)+12)) = v2072
	*(*int32)(unsafe.Add(mBase, uint32(v2072)+8)) = v2075
	v2180 = v2068
	v2181 = v2064
	goto L464
L471:
	;
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v2056)+4))
	v2121 = int32(3)
	if v2120&v2121 != v2121 {
		v2180 = v2068
		v2181 = v2064
		goto L464
	} else {
		goto L486
	}
L472:
	;
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v2064)+12))
	if base.Ui32(int32(255)) < base.Ui32(v2063) {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v2064)+24))
	if v2072 == v2064 {
		goto L476
	} else {
		goto L477
	}
L474:
	;
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v2064)+8))
	if v2072 != v2075 {
		goto L470
	} else {
		goto L475
	}
L475:
	;
	v2077 = int32(0)
	v2079 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v2079 & base.I32_rotl(int32(-2), int32(base.Ui32(v2063)>>(uint(int32(3))%32)))
	v2180 = v2068
	v2181 = v2064
	goto L464
L476:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v2064)+20))
	if v2091 == int32(0) {
		goto L479
	} else {
		goto L480
	}
L477:
	;
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(v2064)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2088)+12)) = v2072
	*(*int32)(unsafe.Add(mBase, uint32(v2072)+8)) = v2088
	v2139 = v2072
	goto L468
L478:
	;
	__phi2107 = v2101
	__phi2108 = v2102
	v2107 = __phi2107
	v2108 = __phi2108
	goto L482
L479:
	;
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v2064)+16))
	if v2096 == int32(0) {
		goto L469
	} else {
		goto L481
	}
L480:
	;
	v2101 = v2091
	v2102 = v2064 + int32(20)
	goto L478
L481:
	;
	v2101 = v2096
	v2102 = v2064 + int32(16)
	goto L478
L482:
	;
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v2107)+20))
	if v2114 != 0 {
		__phi2107 = v2114
		__phi2108 = v2107 + int32(20)
		v2107 = __phi2107
		v2108 = __phi2108
		goto L482
	} else {
		goto L484
	}
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2108))) = int32(0)
	v2139 = v2107
	goto L468
L484:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v2107)+16))
	if v2117 != 0 {
		__phi2107 = v2117
		__phi2108 = v2107 + int32(16)
		v2107 = __phi2107
		v2108 = __phi2108
		goto L482
	} else {
		goto L485
	}
L485:
	;
	goto L483
L486:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v2068
	*(*int32)(unsafe.Add(mBase, uint32(v2056)+4)) = v2120 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v2064)+4)) = v2068 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2056))) = v2068
	goto L461
L487:
	;
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(v2064)+28))
	v2150 = v2148 << (uint(int32(2)) % 32)
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v2150)+uint32(_consts[519])))
	if v2064 != v2153 {
		goto L489
	} else {
		goto L490
	}
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2139)+24)) = v2086
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v2064)+16))
	if v2170 == int32(0) {
		goto L496
	} else {
		goto L497
	}
L489:
	;
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v2086)+16))
	if v2163 != v2064 {
		goto L493
	} else {
		goto L494
	}
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2150)+uint32(_consts[519]))) = v2139
	if v2139 != 0 {
		goto L488
	} else {
		goto L491
	}
L491:
	;
	v2156 = int32(0)
	v2158 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v2158 & base.I32_rotl(int32(-2), v2148)
	v2180 = v2068
	v2181 = v2064
	goto L464
L492:
	;
	if v2139 == int32(0) {
		v2180 = v2068
		v2181 = v2064
		goto L464
	} else {
		goto L495
	}
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2086)+20)) = v2139
	goto L492
L494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2086)+16)) = v2139
	goto L492
L495:
	;
	goto L488
L496:
	;
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v2064)+20))
	if v2175 == int32(0) {
		v2180 = v2068
		v2181 = v2064
		goto L464
	} else {
		goto L498
	}
L497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2139)+16)) = v2170
	*(*int32)(unsafe.Add(mBase, uint32(v2170)+24)) = v2139
	goto L496
L498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2139)+20)) = v2175
	*(*int32)(unsafe.Add(mBase, uint32(v2175)+24)) = v2139
	v2180 = v2068
	v2181 = v2064
	goto L464
L499:
	;
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v2056)+4))
	if v2190&int32(1) == int32(0) {
		goto L462
	} else {
		goto L500
	}
L500:
	;
	if v2190&int32(2) != 0 {
		goto L505
	} else {
		goto L506
	}
L501:
	;
	if base.Ui32(int32(255)) < base.Ui32(v2356) {
		goto L539
	} else {
		goto L540
	}
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+4)) = v2236 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2181+v2236))) = v2236
	if v2181 != v2220 {
		v2356 = v2236
		goto L501
	} else {
		goto L538
	}
L503:
	;
	if v2253 == int32(0) {
		goto L502
	} else {
		goto L526
	}
L504:
	;
	v2298 = int32(0)
	goto L503
L505:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2056)+4)) = v2190 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+4)) = v2180 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2181+v2180))) = v2180
	v2356 = v2180
	goto L501
L506:
	;
	v2198 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	if v2056 != v2198 {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v2220 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v2056 != v2220 {
		goto L510
	} else {
		goto L511
	}
L508:
	;
	v2200 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v2181
	v2204 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v2205 = v2204 + v2180
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v2205
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+4)) = v2205 | int32(1)
	v2211 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v2181 != v2211 {
		goto L462
	} else {
		goto L509
	}
L509:
	;
	v2213 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v2213
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v2213
	goto L461
L510:
	;
	v2236 = v2190&int32(-8) + v2180
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v2056)+12))
	if base.Ui32(int32(255)) < base.Ui32(v2190) {
		goto L512
	} else {
		goto L513
	}
L511:
	;
	v2222 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v2181
	v2226 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	v2227 = v2226 + v2180
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v2227
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+4)) = v2227 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2181+v2227))) = v2227
	goto L461
L512:
	;
	v2253 = *(*int32)(unsafe.Add(mBase, uint32(v2056)+24))
	if v2237 == v2056 {
		goto L516
	} else {
		goto L517
	}
L513:
	;
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v2056)+8))
	if v2237 != v2240 {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2240)+12)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v2237)+8)) = v2240
	goto L502
L515:
	;
	v2242 = int32(0)
	v2244 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v2244 & base.I32_rotl(int32(-2), int32(base.Ui32(v2190)>>(uint(int32(3))%32)))
	goto L502
L516:
	;
	v2258 = *(*int32)(unsafe.Add(mBase, uint32(v2056)+20))
	if v2258 == int32(0) {
		goto L519
	} else {
		goto L520
	}
L517:
	;
	v2255 = *(*int32)(unsafe.Add(mBase, uint32(v2056)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2255)+12)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v2237)+8)) = v2255
	v2298 = v2237
	goto L503
L518:
	;
	__phi2274 = v2268
	__phi2275 = v2269
	v2274 = __phi2274
	v2275 = __phi2275
	goto L522
L519:
	;
	v2263 = *(*int32)(unsafe.Add(mBase, uint32(v2056)+16))
	if v2263 == int32(0) {
		goto L504
	} else {
		goto L521
	}
L520:
	;
	v2268 = v2258
	v2269 = v2056 + int32(20)
	goto L518
L521:
	;
	v2268 = v2263
	v2269 = v2056 + int32(16)
	goto L518
L522:
	;
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+20))
	if v2281 != 0 {
		__phi2274 = v2281
		__phi2275 = v2274 + int32(20)
		v2274 = __phi2274
		v2275 = __phi2275
		goto L522
	} else {
		goto L524
	}
L523:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2275))) = int32(0)
	v2298 = v2274
	goto L503
L524:
	;
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+16))
	if v2284 != 0 {
		__phi2274 = v2284
		__phi2275 = v2274 + int32(16)
		v2274 = __phi2274
		v2275 = __phi2275
		goto L522
	} else {
		goto L525
	}
L525:
	;
	goto L523
L526:
	;
	v2307 = *(*int32)(unsafe.Add(mBase, uint32(v2056)+28))
	v2309 = v2307 << (uint(int32(2)) % 32)
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(v2309)+uint32(_consts[519])))
	if v2056 != v2312 {
		goto L528
	} else {
		goto L529
	}
L527:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2298)+24)) = v2253
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(v2056)+16))
	if v2329 == int32(0) {
		goto L535
	} else {
		goto L536
	}
L528:
	;
	v2322 = *(*int32)(unsafe.Add(mBase, uint32(v2253)+16))
	if v2322 != v2056 {
		goto L532
	} else {
		goto L533
	}
L529:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2309)+uint32(_consts[519]))) = v2298
	if v2298 != 0 {
		goto L527
	} else {
		goto L530
	}
L530:
	;
	v2315 = int32(0)
	v2317 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v2317 & base.I32_rotl(int32(-2), v2307)
	goto L502
L531:
	;
	if v2298 == int32(0) {
		goto L502
	} else {
		goto L534
	}
L532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2253)+20)) = v2298
	goto L531
L533:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2253)+16)) = v2298
	goto L531
L534:
	;
	goto L527
L535:
	;
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(v2056)+20))
	if v2334 == int32(0) {
		goto L502
	} else {
		goto L537
	}
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2298)+16)) = v2329
	*(*int32)(unsafe.Add(mBase, uint32(v2329)+24)) = v2298
	goto L535
L537:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2298)+20)) = v2334
	*(*int32)(unsafe.Add(mBase, uint32(v2334)+24)) = v2298
	goto L502
L538:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v2236
	goto L461
L539:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v2356) {
		v2403 = int32(31)
		goto L544
	} else {
		goto L545
	}
L540:
	;
	v2368 = v2356 & int32(-8)
	v2370 = v2368 + int32(9128464)
	v2372 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v2376 = int32(1) << (uint(int32(base.Ui32(v2356)>>(uint(int32(3))%32))) % 32)
	if v2372&v2376 != 0 {
		goto L542
	} else {
		goto L543
	}
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2368)+uint32(_consts[523]))) = v2181
	*(*int32)(unsafe.Add(mBase, uint32(v2382)+12)) = v2181
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+12)) = v2370
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+8)) = v2382
	goto L461
L542:
	;
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v2368)+uint32(_consts[523])))
	v2382 = v2381
	goto L541
L543:
	;
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v2372 | v2376
	v2382 = v2370
	goto L541
L544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+28)) = v2403
	*(*int64)(unsafe.Add(mBase, uint32(v2181)+16)) = int64(0)
	v2408 = v2403 << (uint(int32(2)) % 32)
	v2412 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	v2414 = int32(1) << (uint(v2403) % 32)
	if v2412&v2414 != 0 {
		goto L549
	} else {
		goto L550
	}
L545:
	;
	v2393 = base.I32_clz(int32(base.Ui32(v2356) >> (uint(int32(8)) % 32)))
	v2396 = int32(1)
	v2403 = int32(base.Ui32(v2356)>>(uint(int32(38)-v2393)%32))&v2396 - v2393<<(uint(v2396)%32) + int32(62)
	goto L544
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2181+v2475))) = v2478
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+12)) = v2477
	*(*int32)(unsafe.Add(mBase, uint32(v2181+v2473))) = v2476
	v2487 = int32(0)
	v2489 = *(*int32)(unsafe.Add(mBase, _consts[524]))
	v2490 = int32(-1)
	v2491 = v2489 + v2490
	if v2491 != 0 {
		goto L558
	} else {
		goto L559
	}
L547:
	;
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(v2437)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2467)+12)) = v2181
	*(*int32)(unsafe.Add(mBase, uint32(v2437)+8)) = v2181
	v2473 = int32(24)
	v2475 = int32(8)
	v2476 = int32(0)
	v2477 = v2437
	v2478 = v2467
	goto L546
L548:
	;
	v2473 = v2458
	v2475 = v2460
	v2476 = v2181
	v2477 = v2181
	v2478 = v2463
	goto L546
L549:
	;
	if v2403 == int32(31) {
		goto L551
	} else {
		goto L552
	}
L550:
	;
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v2412 | v2414
	*(*int32)(unsafe.Add(mBase, uint32(v2408)+uint32(_consts[519]))) = v2181
	v2458 = int32(8)
	v2460 = int32(24)
	v2463 = v2408 + int32(9128728)
	goto L548
L551:
	;
	v2429 = int32(0)
	goto L553
L552:
	;
	v2429 = int32(25) - int32(base.Ui32(v2403)>>(uint(int32(1))%32))
	goto L553
L553:
	;
	v2431 = *(*int32)(unsafe.Add(mBase, uint32(v2408)+uint32(_consts[519])))
	v2434 = v2356 << (uint(v2429) % 32)
	v2437 = v2431
	goto L554
L554:
	;
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(v2437)+4))
	if v2441&int32(-8) == v2356 {
		goto L547
	} else {
		goto L556
	}
L555:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2451+int32(16)))) = v2181
	v2458 = int32(8)
	v2460 = int32(24)
	v2463 = v2437
	goto L548
L556:
	;
	v2451 = v2437 + int32(base.Ui32(v2434)>>(uint(int32(29))%32))&int32(4)
	v2452 = *(*int32)(unsafe.Add(mBase, uint32(v2451)+16))
	if v2452 != 0 {
		v2434 = v2434 << (uint(int32(1)) % 32)
		v2437 = v2452
		goto L554
	} else {
		goto L557
	}
L557:
	;
	goto L555
L558:
	;
	v2493 = v2491
	goto L560
L559:
	;
	v2493 = v2490
	goto L560
L560:
	;
	*(*int32)(unsafe.Add(mBase, _consts[524])) = v2493
	goto L462
L561:
	;
	m.G0 = v2539 + int32(304)
	return v2542
L562:
	;
	goto L563
L563:
	;
	goto L561
}
func F_vkmem_dlerror(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	*(*int32)(unsafe.Add(mBase, _consts[2])) = v1
	return v3
}
func F_vsetBucketRemoveExpired_HASHTABLE(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	v8 = m.G0
	v10 = v8 - int32(128)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 + int32(1) {
	case 0:
		goto L2
	case 1:
		goto L4
	default:
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a1872), int32(_a1861), int32(1523))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L35
	}
L2:
	;
	F__serverAssert(m, int32(_a1873), int32(_a1861), int32(802))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L34
	}
L3:
	;
	if v12&int32(7) != int32(4) {
		goto L2
	} else {
		goto L7
	}
L4:
	;
	F__serverAssert(m, int32(_a1864), int32(_a1861), int32(781))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L7:
	;
	v28 = v10 + int32(80)
	v30 = v12 & int32(-8)
	v31 = int32(1)
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+14)) = uint8(v32)
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v32
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+15)) = uint8(v31)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = int32(-1)
	if v30 == v32 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v49 = int32(0)
	if l2 == v49 {
		v83 = v49
		goto L12
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
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v30)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30)+40)) = v28
	goto L9
L12:
	;
	F_hashtableCleanupIterator(m, v10+int32(80))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L22
	}
L13:
	;
	v57 = v49
	goto L14
L14:
	;
	v63 = F_hashtableNext(m, v10+int32(80), v10+int32(76))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	v83 = l2
	goto L12
L16:
	;
	if v63 == int32(0) {
		v83 = v57
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
	v68 = F_hashtableDelete(m, v30, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	if v68 == int32(0) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
	v73 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v72, l3)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v76 = v57 + int32(1)
	if v76 != l2 {
		v57 = v76
		goto L14
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	goto L26
L23:
	;
	m.G0 = v10 + int32(128)
	return v83
L24:
	;
	v97 = v10 + int32(16)
	v98 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+14)) = uint8(v98)
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v97)+24)) = v98
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+15)) = uint8(v98)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = int32(-1)
	if v30 == v98 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	F_hashtableRelease(m, v30)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L27
	}
L26:
	;
	switch v89 + v90 {
	case 0:
		goto L25
	case 1:
		goto L24
	default:
		goto L23
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(-1)
	goto L23
L28:
	;
	v120 = F_hashtableNext(m, v10+int32(16), v10+int32(12))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L32
	}
L29:
	;
	goto L28
L30:
	;
	goto L29
L32:
	;
	F_hashtableRelease(m, v30)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v124
	goto L23
L34:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
