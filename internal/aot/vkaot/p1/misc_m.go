package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___memrchr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v9 = l2
	goto L1
L1:
	;
	if v9 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v15
L3:
	;
	v14 = v9 + int32(-1)
	v15 = l0 + v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v16 != l1&int32(255) {
		v9 = v14
		goto L1
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L2
}
func F___month_to_secs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1052])))
	if l1 != 0 {
		v11 = v8 + int32(86400)
	} else {
		v11 = v8
	}
	if int32(1) < l0 {
		v14 = v11
	} else {
		v14 = v8
	}
	return v14
}
func F_main(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int64
	_ = v53
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v70 int32
	_ = v70
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v93 int64
	_ = v93
	var v96 int64
	_ = v96
	var v107 int64
	_ = v107
	var v110 int64
	_ = v110
	var v123 int64
	_ = v123
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v165 int64
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
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
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v338 int32
	_ = v338
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
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
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v748 int32
	_ = v748
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int64
	_ = v1102
	var v1109 int32
	_ = v1109
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1148 int32
	_ = v1148
	var v1154 int32
	_ = v1154
	var v1155 int64
	_ = v1155
	var v1160 int64
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1176 int32
	_ = v1176
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1201 int64
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1220 int32
	_ = v1220
	var v1224 int32
	_ = v1224
	var v1229 int32
	_ = v1229
	var v1236 int32
	_ = v1236
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1356 int32
	_ = v1356
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1564 int64
	_ = v1564
	var v1570 int32
	_ = v1570
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1630 int32
	_ = v1630
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1670 int32
	_ = v1670
	var v1676 int32
	_ = v1676
	v13 = m.G0
	v15 = v13 - int32(192)
	m.G0 = v15
	v17 = int32(9116396)
	F___lock(m, v17)
	mBase = m.M
	F_do_tzset(m)
	mBase = m.M
	F___unlock(m, v17)
	mBase = m.M
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[803])) = int32(1034)
	goto L2
L2:
	;
	v25 = int32(0)
	v29 = F___gettimeofday(m, v15+int32(176), v25)
	mBase = m.M
	v31 = F___time(m, v25)
	mBase = m.M
	v32 = F___syscall_getpid(m)
	mBase = m.M
	goto L3
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v15)+184))
	*(*int64)(unsafe.Add(mBase, _consts[343])) = base.I64_extend_i32_u(v32 ^ base.I32_wrap_i64(v31) ^ v35 + int32(-1))
	goto L4
L4:
	;
	v43 = F___time(m, int32(0))
	mBase = m.M
	v44 = F___syscall_getpid(m)
	mBase = m.M
	goto L5
L5:
	;
	v48 = int32(9116960)
	F___lock(m, v48)
	mBase = m.M
	F___srandom(m, v35^(v44^base.I32_wrap_i64(v43)))
	mBase = m.M
	F___unlock(m, v48)
	mBase = m.M
	goto L6
L6:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v15)+176))
	v58 = F___syscall_getpid(m)
	mBase = m.M
	goto L7
L7:
	;
	v60 = v53*int64(1000000) + base.I64_extend_i32_s(v35) ^ base.I64_extend_i32_s(v58)
	*(*int64)(unsafe.Add(mBase, _consts[467])) = v60
	v66 = v60
	v67 = int64(1)
	goto L10
L8:
	;
	F_crcspeed64native_init(m, int32(513), int32(_a555))
	mBase = m.M
	goto L13
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[466])) = int32(312)
	goto L8
L10:
	;
	v70 = int32(3)
	v74 = int64(62)
	v77 = int64(6364136223846793005)
	v79 = (int64(base.Ui64(v66)>>(uint(v74)%64))^v66)*v77 + v67
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v67)<<(uint(v70)%32))+uint32(_consts[467]))) = v79
	v82 = v67 + int64(1)
	v93 = (int64(base.Ui64(v79)>>(uint(v74)%64))^v79)*v77 + v82
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v82)<<(uint(v70)%32))+uint32(_consts[467]))) = v93
	v96 = v67 + int64(2)
	v107 = (int64(base.Ui64(v93)>>(uint(v74)%64))^v93)*v77 + v96
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v96)<<(uint(v70)%32))+uint32(_consts[467]))) = v107
	v110 = v67 + int64(3)
	if v110 == int64(312) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v123 = (int64(base.Ui64(v107)>>(uint(int64(62))%64))^v107)*int64(6364136223846793005) + v110
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v110)<<(uint(int32(3))%32))+uint32(_consts[467]))) = v123
	v66 = v123
	v67 = v67 + int64(4)
	goto L10
L13:
	;
	v135 = F___syscall_umask(m, int32(511))
	mBase = m.M
	v136 = F___syscall_ret(m, v135)
	mBase = m.M
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[253])) = v136
	v138 = F___syscall_umask(m, v136)
	mBase = m.M
	v139 = F___syscall_ret(m, v138)
	mBase = m.M
	goto L15
L15:
	;
	F_getRandomBytes(m, v15+int32(160), int32(16))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	v149 = int32(0)
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(160))))
	*(*int64)(unsafe.Add(mBase, _consts[804])) = v150
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(168))))
	*(*int64)(unsafe.Add(mBase, _consts[805])) = v155
	goto L18
L18:
	;
	v159 = int32(0)
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(160))))
	*(*int64)(unsafe.Add(mBase, _consts[341])) = v160
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(168))))
	*(*int64)(unsafe.Add(mBase, _consts[342])) = v165
	goto L19
L19:
	;
	v167 = int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v170 = F_strlen(m, v168)
	mBase = m.M
	v173 = F___memrchr(m, v168, int32(47), v170+v167)
	mBase = m.M
	goto L21
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[250])) = v298
	F_initServerConfig(m)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L16
	} else {
		goto L73
	}
L21:
	;
	if v173 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v174 = v173
	goto L24
L23:
	;
	v174 = v168
	goto L24
L24:
	;
	v175 = int32(_a2295)
	v178 = int32(*(*int8)(unsafe.Add(mBase, _consts[806])))
	if v178 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v203 != 0 {
		v298 = v167
		goto L20
	} else {
		goto L41
	}
L26:
	;
	v179 = int32(0)
	v180 = F_strchr(m, v174, v178)
	mBase = m.M
	if v180 == v179 {
		v200 = v179
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v203 = v174
	goto L25
L28:
	;
	v203 = v200
	goto L25
L29:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, _consts[807])))
	if v183 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
	if v184 == int32(0) {
		v200 = v179
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v203 = v180
	goto L25
L32:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, _consts[808])))
	if v187 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+2)))
	if v189 == int32(0) {
		v200 = v179
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v188 = F_twobyte_strstr(m, v180, v175)
	mBase = m.M
	v203 = v188
	goto L25
L35:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, _consts[809])))
	if v192 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+3)))
	if v194 == int32(0) {
		v200 = v179
		goto L28
	} else {
		goto L38
	}
L37:
	;
	v193 = F_threebyte_strstr(m, v180, v175)
	mBase = m.M
	v203 = v193
	goto L25
L38:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, _consts[810])))
	if v197 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v199 = F_twoway_strstr(m, v180, v175)
	mBase = m.M
	v200 = v199
	goto L28
L40:
	;
	v198 = F_fourbyte_strstr(m, v180, v175)
	mBase = m.M
	v203 = v198
	goto L25
L41:
	;
	v205 = int32(_a2296)
	v208 = int32(*(*int8)(unsafe.Add(mBase, _consts[811])))
	if v208 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v233 != 0 {
		v298 = int32(1)
		goto L20
	} else {
		goto L58
	}
L43:
	;
	v209 = int32(0)
	v210 = F_strchr(m, v174, v208)
	mBase = m.M
	if v210 == v209 {
		v230 = v209
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v233 = v174
	goto L42
L45:
	;
	v233 = v230
	goto L42
L46:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, _consts[812])))
	if v213 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
	if v214 == int32(0) {
		v230 = v209
		goto L45
	} else {
		goto L49
	}
L48:
	;
	v233 = v210
	goto L42
L49:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, _consts[813])))
	if v217 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+2)))
	if v219 == int32(0) {
		v230 = v209
		goto L45
	} else {
		goto L52
	}
L51:
	;
	v218 = F_twobyte_strstr(m, v210, v205)
	mBase = m.M
	v233 = v218
	goto L42
L52:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, _consts[814])))
	if v222 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+3)))
	if v224 == int32(0) {
		v230 = v209
		goto L45
	} else {
		goto L55
	}
L54:
	;
	v223 = F_threebyte_strstr(m, v210, v205)
	mBase = m.M
	v233 = v223
	goto L42
L55:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, _consts[815])))
	if v227 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v229 = F_twoway_strstr(m, v210, v205)
	mBase = m.M
	v230 = v229
	goto L45
L57:
	;
	v228 = F_fourbyte_strstr(m, v210, v205)
	mBase = m.M
	v233 = v228
	goto L42
L58:
	;
	v234 = int32(1)
	v235 = int32(0)
	if l0 <= v234 {
		v298 = v235
		goto L20
	} else {
		goto L59
	}
L59:
	;
	v243 = v234
	goto L61
L60:
	;
	v298 = int32(1)
	goto L20
L61:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1+v243<<(uint(int32(2))%32))))
	v254 = int32(_a2297)
	v257 = int32(*(*uint8)(unsafe.Add(mBase, _consts[816])))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	if v258 == int32(0) {
		v281 = v257
		v282 = v258
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v282-v281&int32(255) == int32(0) {
		goto L60
	} else {
		goto L71
	}
L64:
	;
	goto L63
L65:
	;
	if v258 != v257&int32(255) {
		v281 = v257
		v282 = v258
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v264 = v253
	v265 = v254
	goto L67
L67:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+1)))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+1)))
	if v269 == int32(0) {
		v281 = v268
		v282 = v269
		goto L64
	} else {
		goto L69
	}
L68:
	;
	v281 = v268
	v282 = v269
	goto L64
L69:
	;
	v272 = int32(1)
	if v269 == v268&int32(255) {
		v264 = v264 + v272
		v265 = v265 + v272
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v289 = v243 + int32(1)
	if v289 == l0 {
		v298 = v235
		goto L20
	} else {
		goto L72
	}
L72:
	;
	v243 = v289
	goto L61
L73:
	;
	v309 = F___syscall_getpid(m)
	mBase = m.M
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, _consts[154])) = v309
	F_ACLInit(m)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	F_moduleInitModulesSystem(m)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L16
	} else {
		goto L76
	}
L76:
	;
	v315 = F_connTypeInitialize(m)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L16
	} else {
		goto L77
	}
L77:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v319 = F_getAbsolutePath(m, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L16
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, _consts[749])) = v319
	v324 = l0 << (uint(int32(2)) % 32)
	v327 = F_valkey_malloc(m, v324+int32(4))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L16
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, _consts[531])) = v327
	*(*int32)(unsafe.Add(mBase, uint32(v327+v324))) = int32(0)
	if l0 < int32(1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v372 = int32(0)
	v373 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	if v373 == v372 {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v338 = v25
	goto L82
L82:
	;
	v348 = v338 << (uint(int32(2)) % 32)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l1+v348)))
	v351 = F_zstrdup(m, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L16
	} else {
		goto L84
	}
L83:
	;
	goto L80
L84:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _consts[531]))
	*(*int32)(unsafe.Add(mBase, uint32(v354+v348))) = v351
	v358 = v338 + int32(1)
	if v358 != l0 {
		v338 = v358
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v384 = int32(_a2298)
	v387 = int32(*(*int8)(unsafe.Add(mBase, _consts[817])))
	if v387 != 0 {
		goto L93
	} else {
		goto L94
	}
L87:
	;
	v376 = int32(_a20)
	*(*int32)(unsafe.Add(mBase, _consts[489])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(26379)
	goto L88
L88:
	;
	F_initSentinel(m)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L16
	} else {
		goto L89
	}
L89:
	;
	goto L86
L90:
	;
	v451 = int32(_a2299)
	v454 = int32(*(*int8)(unsafe.Add(mBase, _consts[818])))
	if v454 != 0 {
		goto L131
	} else {
		goto L132
	}
L91:
	;
	v418 = int32(_a2300)
	v421 = int32(*(*int8)(unsafe.Add(mBase, _consts[819])))
	if v421 != 0 {
		goto L111
	} else {
		goto L112
	}
L92:
	;
	if v412 == int32(0) {
		goto L91
	} else {
		goto L108
	}
L93:
	;
	v388 = int32(0)
	v389 = F_strchr(m, v174, v387)
	mBase = m.M
	if v389 == v388 {
		v409 = v388
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v412 = v174
	goto L92
L95:
	;
	v412 = v409
	goto L92
L96:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, _consts[820])))
	if v392 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+1)))
	if v393 == int32(0) {
		v409 = v388
		goto L95
	} else {
		goto L99
	}
L98:
	;
	v412 = v389
	goto L92
L99:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, _consts[821])))
	if v396 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+2)))
	if v398 == int32(0) {
		v409 = v388
		goto L95
	} else {
		goto L102
	}
L101:
	;
	v397 = F_twobyte_strstr(m, v389, v384)
	mBase = m.M
	v412 = v397
	goto L92
L102:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, _consts[822])))
	if v401 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+3)))
	if v403 == int32(0) {
		v409 = v388
		goto L95
	} else {
		goto L105
	}
L104:
	;
	v402 = F_threebyte_strstr(m, v389, v384)
	mBase = m.M
	v412 = v402
	goto L92
L105:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, _consts[823])))
	if v406 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v408 = F_twoway_strstr(m, v389, v384)
	mBase = m.M
	v409 = v408
	goto L95
L107:
	;
	v407 = F_fourbyte_strstr(m, v389, v384)
	mBase = m.M
	v412 = v407
	goto L92
L108:
	;
	v416 = F_redis_check_rdb_main(m, l0, l1, int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L16
	} else {
		goto L109
	}
L109:
	;
	goto L90
L110:
	;
	if v446 == int32(0) {
		goto L90
	} else {
		goto L126
	}
L111:
	;
	v422 = int32(0)
	v423 = F_strchr(m, v174, v421)
	mBase = m.M
	if v423 == v422 {
		v443 = v422
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v446 = v174
	goto L110
L113:
	;
	v446 = v443
	goto L110
L114:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, _consts[824])))
	if v426 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+1)))
	if v427 == int32(0) {
		v443 = v422
		goto L113
	} else {
		goto L117
	}
L116:
	;
	v446 = v423
	goto L110
L117:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, _consts[825])))
	if v430 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+2)))
	if v432 == int32(0) {
		v443 = v422
		goto L113
	} else {
		goto L120
	}
L119:
	;
	v431 = F_twobyte_strstr(m, v423, v418)
	mBase = m.M
	v446 = v431
	goto L110
L120:
	;
	v435 = int32(*(*uint8)(unsafe.Add(mBase, _consts[826])))
	if v435 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+3)))
	if v437 == int32(0) {
		v443 = v422
		goto L113
	} else {
		goto L123
	}
L122:
	;
	v436 = F_threebyte_strstr(m, v423, v418)
	mBase = m.M
	v446 = v436
	goto L110
L123:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, _consts[827])))
	if v440 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v442 = F_twoway_strstr(m, v423, v418)
	mBase = m.M
	v443 = v442
	goto L113
L125:
	;
	v441 = F_fourbyte_strstr(m, v423, v418)
	mBase = m.M
	v446 = v441
	goto L110
L126:
	;
	v449 = F_redis_check_aof_main(m, l0, l1)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L16
	} else {
		goto L127
	}
L127:
	;
	goto L90
L128:
	;
	if l0 < int32(2) {
		goto L172
	} else {
		goto L173
	}
L129:
	;
	v485 = int32(_a2301)
	v488 = int32(*(*int8)(unsafe.Add(mBase, _consts[828])))
	if v488 != 0 {
		goto L149
	} else {
		goto L150
	}
L130:
	;
	if v479 == int32(0) {
		goto L129
	} else {
		goto L146
	}
L131:
	;
	v455 = int32(0)
	v456 = F_strchr(m, v174, v454)
	mBase = m.M
	if v456 == v455 {
		v476 = v455
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v479 = v174
	goto L130
L133:
	;
	v479 = v476
	goto L130
L134:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, _consts[829])))
	if v459 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456)+1)))
	if v460 == int32(0) {
		v476 = v455
		goto L133
	} else {
		goto L137
	}
L136:
	;
	v479 = v456
	goto L130
L137:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, _consts[830])))
	if v463 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456)+2)))
	if v465 == int32(0) {
		v476 = v455
		goto L133
	} else {
		goto L140
	}
L139:
	;
	v464 = F_twobyte_strstr(m, v456, v451)
	mBase = m.M
	v479 = v464
	goto L130
L140:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, _consts[831])))
	if v468 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456)+3)))
	if v470 == int32(0) {
		v476 = v455
		goto L133
	} else {
		goto L143
	}
L142:
	;
	v469 = F_threebyte_strstr(m, v456, v451)
	mBase = m.M
	v479 = v469
	goto L130
L143:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, _consts[832])))
	if v473 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v475 = F_twoway_strstr(m, v456, v451)
	mBase = m.M
	v476 = v475
	goto L133
L145:
	;
	v474 = F_fourbyte_strstr(m, v456, v451)
	mBase = m.M
	v479 = v474
	goto L130
L146:
	;
	v483 = F_redis_check_rdb_main(m, l0, l1, int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L16
	} else {
		goto L147
	}
L147:
	;
	goto L128
L148:
	;
	if v513 == int32(0) {
		goto L128
	} else {
		goto L164
	}
L149:
	;
	v489 = int32(0)
	v490 = F_strchr(m, v174, v488)
	mBase = m.M
	if v490 == v489 {
		v510 = v489
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v513 = v174
	goto L148
L151:
	;
	v513 = v510
	goto L148
L152:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, _consts[833])))
	if v493 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+1)))
	if v494 == int32(0) {
		v510 = v489
		goto L151
	} else {
		goto L155
	}
L154:
	;
	v513 = v490
	goto L148
L155:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, _consts[834])))
	if v497 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+2)))
	if v499 == int32(0) {
		v510 = v489
		goto L151
	} else {
		goto L158
	}
L157:
	;
	v498 = F_twobyte_strstr(m, v490, v485)
	mBase = m.M
	v513 = v498
	goto L148
L158:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, _consts[835])))
	if v502 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+3)))
	if v504 == int32(0) {
		v510 = v489
		goto L151
	} else {
		goto L161
	}
L160:
	;
	v503 = F_threebyte_strstr(m, v490, v485)
	mBase = m.M
	v513 = v503
	goto L148
L161:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, _consts[836])))
	if v507 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v509 = F_twoway_strstr(m, v490, v485)
	mBase = m.M
	v510 = v509
	goto L151
L163:
	;
	v508 = F_fourbyte_strstr(m, v490, v485)
	mBase = m.M
	v513 = v508
	goto L148
L164:
	;
	v516 = F_redis_check_aof_main(m, l0, l1)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L16
	} else {
		goto L165
	}
L165:
	;
	goto L128
L166:
	;
	F__serverAssert(m, int32(_a2302), int32(_a2157), int32(7749))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L16
	} else {
		goto L498
	}
L167:
	;
	F__serverPanic_1(m, int32(_a2157), int32(7728), int32(_a2303), int32(0))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L16
	} else {
		goto L497
	}
L168:
	;
	v1660 = F_syscheck(m)
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L16
	} else {
		goto L496
	}
L169:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1609 = v1605
	goto L481
L170:
	;
	F_usage(m)
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L16
	} else {
		goto L479
	}
L171:
	;
	v1593 = F_getVersion(m)
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L16
	} else {
		goto L476
	}
L172:
	;
	v1092 = int32(0)
	v1093 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	if v1093 == v1092 {
		goto L344
	} else {
		goto L345
	}
L173:
	;
	v520 = F_sdsempty(m)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L16
	} else {
		goto L174
	}
L174:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	if v523 != int32(45) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v532 = int32(_a2304)
	v535 = int32(*(*uint8)(unsafe.Add(mBase, _consts[837])))
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	if v536 == int32(0) {
		v559 = v535
		v560 = v536
		goto L180
	} else {
		goto L181
	}
L176:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522)+1)))
	if v526 != int32(118) {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522)+2)))
	if v529 == int32(0) {
		goto L171
	} else {
		goto L178
	}
L178:
	;
	goto L175
L179:
	;
	if v560-v559&int32(255) == int32(0) {
		goto L171
	} else {
		goto L187
	}
L180:
	;
	goto L179
L181:
	;
	if v536 != v535&int32(255) {
		v559 = v535
		v560 = v536
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v542 = v522
	v543 = v532
	goto L183
L183:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543)+1)))
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+1)))
	if v547 == int32(0) {
		v559 = v546
		v560 = v547
		goto L180
	} else {
		goto L185
	}
L184:
	;
	v559 = v546
	v560 = v547
	goto L180
L185:
	;
	v550 = int32(1)
	if v547 == v546&int32(255) {
		v542 = v542 + v550
		v543 = v543 + v550
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v566 = int32(_a2305)
	v569 = int32(*(*uint8)(unsafe.Add(mBase, _consts[838])))
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	if v570 == int32(0) {
		v593 = v569
		v594 = v570
		goto L189
	} else {
		goto L190
	}
L188:
	;
	if v594-v593&int32(255) == int32(0) {
		goto L170
	} else {
		goto L196
	}
L189:
	;
	goto L188
L190:
	;
	if v570 != v569&int32(255) {
		v593 = v569
		v594 = v570
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v576 = v522
	v577 = v566
	goto L192
L192:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577)+1)))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+1)))
	if v581 == int32(0) {
		v593 = v580
		v594 = v581
		goto L189
	} else {
		goto L194
	}
L193:
	;
	v593 = v580
	v594 = v581
	goto L189
L194:
	;
	v584 = int32(1)
	if v581 == v580&int32(255) {
		v576 = v576 + v584
		v577 = v577 + v584
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	if v523 != int32(45) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v608 = int32(_a2306)
	v611 = int32(*(*uint8)(unsafe.Add(mBase, _consts[839])))
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	if v612 == int32(0) {
		v635 = v611
		v636 = v612
		goto L203
	} else {
		goto L204
	}
L198:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522)+1)))
	if v602 != int32(104) {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522)+2)))
	if v605 == int32(0) {
		goto L170
	} else {
		goto L200
	}
L200:
	;
	goto L197
L201:
	;
	v656 = int32(_a2307)
	v659 = int32(*(*uint8)(unsafe.Add(mBase, _consts[840])))
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	if v660 == int32(0) {
		v683 = v659
		v684 = v660
		goto L215
	} else {
		goto L216
	}
L202:
	;
	if v636-v635&int32(255) != 0 {
		goto L201
	} else {
		goto L210
	}
L203:
	;
	goto L202
L204:
	;
	if v612 != v611&int32(255) {
		v635 = v611
		v636 = v612
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v618 = v522
	v619 = v608
	goto L206
L206:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+1)))
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618)+1)))
	if v623 == int32(0) {
		v635 = v622
		v636 = v623
		goto L203
	} else {
		goto L208
	}
L207:
	;
	v635 = v622
	v636 = v623
	goto L203
L208:
	;
	v626 = int32(1)
	if v623 == v622&int32(255) {
		v618 = v618 + v626
		v619 = v619 + v626
		goto L206
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	if l0 == int32(3) {
		goto L169
	} else {
		goto L211
	}
L211:
	;
	v646 = *(*int32)(unsafe.Add(mBase, _consts[245]))
	v647 = F_fwrite(m, int32(_a2308), int32(58), int32(1), v646)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L16
	} else {
		goto L212
	}
L212:
	;
	v652 = F_fwrite(m, int32(_a2309), int32(45), int32(1), v646)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L16
	} else {
		goto L213
	}
L213:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	if v684-v683&int32(255) == int32(0) {
		goto L168
	} else {
		goto L222
	}
L215:
	;
	goto L214
L216:
	;
	if v660 != v659&int32(255) {
		v683 = v659
		v684 = v660
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v666 = v522
	v667 = v656
	goto L218
L218:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+1)))
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+1)))
	if v671 == int32(0) {
		v683 = v670
		v684 = v671
		goto L215
	} else {
		goto L220
	}
L219:
	;
	v683 = v670
	v684 = v671
	goto L215
L220:
	;
	v674 = int32(1)
	if v671 == v670&int32(255) {
		v666 = v666 + v674
		v667 = v667 + v674
		goto L218
	} else {
		goto L221
	}
L221:
	;
	goto L219
L222:
	;
	if v523 == int32(45) {
		v710 = int32(1)
		goto L223
	} else {
		goto L224
	}
L223:
	;
	if base.Ui32(v710) < base.Ui32(l0) {
		goto L229
	} else {
		goto L230
	}
L224:
	;
	v694 = F_getAbsolutePath(m, v522)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L16
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, _consts[690])) = v694
	v698 = *(*int32)(unsafe.Add(mBase, _consts[531]))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v698)+4))
	F_valkey_free(m, v699)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L16
	} else {
		goto L226
	}
L226:
	;
	v703 = *(*int32)(unsafe.Add(mBase, _consts[690]))
	v704 = F_zstrdup(m, v703)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L16
	} else {
		goto L227
	}
L227:
	;
	v707 = *(*int32)(unsafe.Add(mBase, _consts[531]))
	*(*int32)(unsafe.Add(mBase, uint32(v707)+4)) = v704
	v710 = int32(2)
	goto L223
L228:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, _consts[690]))
	F_loadServerConfig(m, v1068, base.I32_extend8_s(v1063), v1061)
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L16
	} else {
		goto L339
	}
L229:
	;
	v714 = l0 + int32(-1)
	v720 = v710
	v723 = v520
	v725 = int32(0)
	v727 = int32(1)
	goto L231
L230:
	;
	v1061 = v520
	v1063 = int32(0)
	goto L228
L231:
	;
	v731 = l1 + v720<<(uint(int32(2))%32)
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732))))
	if v733 != int32(45) {
		goto L234
	} else {
		goto L235
	}
L232:
	;
	v1061 = v1047
	v1063 = v1049
	goto L228
L233:
	;
	v1053 = v720 + int32(1)
	if v1053 != l0 {
		v720 = v1053
		v723 = v1047
		v725 = v1049
		v727 = v1050
		goto L231
	} else {
		goto L338
	}
L234:
	;
	if v732&int32(3) == int32(0) {
		v1007 = v732
		goto L322
	} else {
		goto L323
	}
L235:
	;
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732)+1)))
	if v736 != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	if v727 == int32(0) {
		goto L234
	} else {
		goto L241
	}
L237:
	;
	if v720 != int32(1) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	if v720 != v714 {
		goto L234
	} else {
		goto L240
	}
L239:
	;
	v1047 = v723
	v1049 = int32(1)
	v1050 = v727
	goto L233
L240:
	;
	v1047 = v723
	v1049 = int32(1)
	v1050 = v727
	goto L233
L241:
	;
	if v736 != int32(45) {
		goto L234
	} else {
		goto L242
	}
L242:
	;
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723+int32(-1)))))
	switch v748 & int32(7) {
	case 0:
		goto L249
	case 1:
		goto L248
	case 2:
		goto L247
	case 3:
		goto L246
	case 4:
		goto L245
	default:
		v772 = v732
		v773 = v723
		goto L243
	}
L243:
	;
	v777 = F_sdscat(m, v773, v772+int32(2))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L16
	} else {
		goto L252
	}
L244:
	;
	if v765 == int32(0) {
		v772 = v732
		v773 = v723
		goto L243
	} else {
		goto L250
	}
L245:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v723+int32(-17))))
	v765 = v764
	goto L244
L246:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v723+int32(-9))))
	v765 = v761
	goto L244
L247:
	;
	v758 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v723+int32(-5)))))
	v765 = v758
	goto L244
L248:
	;
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723+int32(-3)))))
	v765 = v755
	goto L244
L249:
	;
	v765 = int32(base.Ui32(v748) >> (uint(int32(3)) % 32))
	goto L244
L250:
	;
	v769 = F_sdscat(m, v723, int32(_a26))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L16
	} else {
		goto L251
	}
L251:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v772 = v771
	v773 = v769
	goto L243
L252:
	;
	v780 = F_sdscat(m, v777, int32(_a6))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L16
	} else {
		goto L253
	}
L253:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v786 = F_sdssplitargs(m, v783, v15+int32(156))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L16
	} else {
		goto L254
	}
L254:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v15)+156))
	if v788 != int32(1) {
		v977 = v780
		v979 = int32(1)
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v15)+156))
	F_sdsfreesplitres(m, v786, v981)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L16
	} else {
		goto L319
	}
L256:
	;
	if v720 == v714 {
		goto L261
	} else {
		goto L262
	}
L257:
	;
	v975 = F_sdscat(m, v780, v972)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L16
	} else {
		goto L318
	}
L258:
	;
	v972 = int32(_a320)
	v973 = v969
	goto L257
L259:
	;
	v928 = int32(0)
	v929 = int32(_a2297)
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842))))
	if v932 != 0 {
		goto L307
	} else {
		goto L308
	}
L260:
	;
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793)+1)))
	if v884 != int32(45) {
		v977 = v780
		v979 = v792
		goto L255
	} else {
		goto L291
	}
L261:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v843 = int32(_a2310)
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842))))
	if v846 != 0 {
		goto L280
	} else {
		goto L281
	}
L262:
	;
	v792 = int32(0)
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v731)+4))
	v794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793))))
	if v794 != int32(45) {
		v977 = v780
		v979 = v792
		goto L255
	} else {
		goto L263
	}
L263:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793)+1)))
	if v797 != int32(45) {
		goto L260
	} else {
		goto L264
	}
L264:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v801 = int32(_a2310)
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800))))
	if v804 != 0 {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	if v836-v838 != 0 {
		goto L260
	} else {
		goto L277
	}
L266:
	;
	v836 = F_tolower(m, v832)
	mBase = m.M
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v833))))
	v838 = F_tolower(m, v837)
	mBase = m.M
	goto L265
L267:
	;
	v806 = v800
	v807 = v801
	v808 = v804
	goto L270
L268:
	;
	v832 = int32(0)
	v833 = v801
	goto L266
L269:
	;
	v832 = v829 & int32(255)
	v833 = v828
	goto L266
L270:
	;
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807))))
	if v810 == int32(0) {
		v828 = v807
		v829 = v808
		goto L269
	} else {
		goto L272
	}
L271:
	;
	v828 = v822
	v829 = int32(0)
	goto L269
L272:
	;
	v814 = v808 & int32(255)
	if v814 == v810 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v821 = int32(1)
	v822 = v807 + v821
	v823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v806)+1)))
	if v823 != 0 {
		v806 = v806 + v821
		v807 = v822
		v808 = v823
		goto L270
	} else {
		goto L276
	}
L274:
	;
	v816 = F_tolower(m, v814)
	mBase = m.M
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807))))
	v818 = F_tolower(m, v817)
	mBase = m.M
	if v816 == v818 {
		goto L273
	} else {
		goto L275
	}
L275:
	;
	v820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v806))))
	v828 = v807
	v829 = v820
	goto L269
L276:
	;
	goto L271
L277:
	;
	v972 = int32(_a542)
	v973 = int32(1)
	goto L257
L278:
	;
	if v878-v880 != 0 {
		goto L259
	} else {
		goto L290
	}
L279:
	;
	v878 = F_tolower(m, v874)
	mBase = m.M
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v875))))
	v880 = F_tolower(m, v879)
	mBase = m.M
	goto L278
L280:
	;
	v848 = v842
	v849 = v843
	v850 = v846
	goto L283
L281:
	;
	v874 = int32(0)
	v875 = v843
	goto L279
L282:
	;
	v874 = v871 & int32(255)
	v875 = v870
	goto L279
L283:
	;
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849))))
	if v852 == int32(0) {
		v870 = v849
		v871 = v850
		goto L282
	} else {
		goto L285
	}
L284:
	;
	v870 = v864
	v871 = int32(0)
	goto L282
L285:
	;
	v856 = v850 & int32(255)
	if v856 == v852 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v863 = int32(1)
	v864 = v849 + v863
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848)+1)))
	if v865 != 0 {
		v848 = v848 + v863
		v849 = v864
		v850 = v865
		goto L283
	} else {
		goto L289
	}
L287:
	;
	v858 = F_tolower(m, v856)
	mBase = m.M
	v859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849))))
	v860 = F_tolower(m, v859)
	mBase = m.M
	if v858 == v860 {
		goto L286
	} else {
		goto L288
	}
L288:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848))))
	v870 = v849
	v871 = v862
	goto L282
L289:
	;
	goto L284
L290:
	;
	v972 = int32(_a542)
	v973 = int32(0)
	goto L257
L291:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v888 = int32(_a2297)
	v891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887))))
	if v891 != 0 {
		goto L294
	} else {
		goto L295
	}
L292:
	;
	if v923-v925 != 0 {
		v977 = v780
		v979 = v792
		goto L255
	} else {
		goto L304
	}
L293:
	;
	v923 = F_tolower(m, v919)
	mBase = m.M
	v924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920))))
	v925 = F_tolower(m, v924)
	mBase = m.M
	goto L292
L294:
	;
	v893 = v887
	v894 = v888
	v895 = v891
	goto L297
L295:
	;
	v919 = int32(0)
	v920 = v888
	goto L293
L296:
	;
	v919 = v916 & int32(255)
	v920 = v915
	goto L293
L297:
	;
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v894))))
	if v897 == int32(0) {
		v915 = v894
		v916 = v895
		goto L296
	} else {
		goto L299
	}
L298:
	;
	v915 = v909
	v916 = int32(0)
	goto L296
L299:
	;
	v901 = v895 & int32(255)
	if v901 == v897 {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v908 = int32(1)
	v909 = v894 + v908
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893)+1)))
	if v910 != 0 {
		v893 = v893 + v908
		v894 = v909
		v895 = v910
		goto L297
	} else {
		goto L303
	}
L301:
	;
	v903 = F_tolower(m, v901)
	mBase = m.M
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v894))))
	v905 = F_tolower(m, v904)
	mBase = m.M
	if v903 == v905 {
		goto L300
	} else {
		goto L302
	}
L302:
	;
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893))))
	v915 = v894
	v916 = v907
	goto L296
L303:
	;
	goto L298
L304:
	;
	v969 = int32(1)
	goto L258
L305:
	;
	if v964-v966 != 0 {
		v977 = v780
		v979 = v928
		goto L255
	} else {
		goto L317
	}
L306:
	;
	v964 = F_tolower(m, v960)
	mBase = m.M
	v965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v961))))
	v966 = F_tolower(m, v965)
	mBase = m.M
	goto L305
L307:
	;
	v934 = v842
	v935 = v929
	v936 = v932
	goto L310
L308:
	;
	v960 = int32(0)
	v961 = v929
	goto L306
L309:
	;
	v960 = v957 & int32(255)
	v961 = v956
	goto L306
L310:
	;
	v938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935))))
	if v938 == int32(0) {
		v956 = v935
		v957 = v936
		goto L309
	} else {
		goto L312
	}
L311:
	;
	v956 = v950
	v957 = int32(0)
	goto L309
L312:
	;
	v942 = v936 & int32(255)
	if v942 == v938 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v949 = int32(1)
	v950 = v935 + v949
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v934)+1)))
	if v951 != 0 {
		v934 = v934 + v949
		v935 = v950
		v936 = v951
		goto L310
	} else {
		goto L316
	}
L314:
	;
	v944 = F_tolower(m, v942)
	mBase = m.M
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935))))
	v946 = F_tolower(m, v945)
	mBase = m.M
	if v944 == v946 {
		goto L313
	} else {
		goto L315
	}
L315:
	;
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v934))))
	v956 = v935
	v957 = v948
	goto L309
L316:
	;
	goto L311
L317:
	;
	v969 = v928
	goto L258
L318:
	;
	v977 = v975
	v979 = v973
	goto L255
L319:
	;
	v1047 = v977
	v1049 = v725
	v1050 = v979
	goto L233
L320:
	;
	v1041 = F_sdscatrepr(m, v723, v732, v1040)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L16
	} else {
		goto L336
	}
L321:
	;
	v1040 = v1032 - v732
	goto L320
L322:
	;
	v1011 = v1007
	goto L330
L323:
	;
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732))))
	if v993 != 0 {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v996 = v732
	goto L326
L325:
	;
	v1040 = v732 - v732
	goto L320
L326:
	;
	v1000 = v996 + int32(1)
	if v1000&int32(3) == int32(0) {
		v1007 = v1000
		goto L322
	} else {
		goto L328
	}
L328:
	;
	v1005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1000))))
	if v1005 != 0 {
		v996 = v1000
		goto L326
	} else {
		goto L329
	}
L329:
	;
	v1032 = v1000
	goto L321
L330:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	v1020 = int32(-2139062144)
	if (int32(16843008)-v1017|v1017)&v1020 == v1020 {
		v1011 = v1011 + int32(4)
		goto L330
	} else {
		goto L332
	}
L331:
	;
	v1026 = v1011
	goto L333
L332:
	;
	goto L331
L333:
	;
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1026))))
	if v1030 != 0 {
		v1026 = v1026 + int32(1)
		goto L333
	} else {
		goto L335
	}
L334:
	;
	v1032 = v1026
	goto L321
L335:
	;
	goto L334
L336:
	;
	v1044 = F_sdscat(m, v1041, int32(_a6))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L16
	} else {
		goto L337
	}
L337:
	;
	v1047 = v1044
	v1049 = v725
	v1050 = int32(1)
	goto L233
L338:
	;
	goto L232
L339:
	;
	v1072 = int32(0)
	v1073 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	if v1073 == v1072 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	F_sdsfree(m, v1061)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L16
	} else {
		goto L343
	}
L341:
	;
	F_loadSentinelConfigFromQueue(m)
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L16
	} else {
		goto L342
	}
L342:
	;
	goto L340
L343:
	;
	goto L172
L344:
	;
	v1098 = int32(0)
	v1099 = *(*int32)(unsafe.Add(mBase, _consts[841]))
	if v1099 == v1098 {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	F_sentinelCheckConfigFile(m)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L16
	} else {
		goto L346
	}
L346:
	;
	goto L344
L347:
	;
	v1164 = int32(0)
	v1166 = *(*int32)(unsafe.Add(mBase, _consts[611]))
	v1167 = F_serverIsSupervised(m, v1166)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L16
	} else {
		goto L360
	}
L348:
	;
	v1102 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+168)) = v1102
	*(*int64)(unsafe.Add(mBase, uint32(v15)+160)) = v1102
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099+int32(-1)))))
	switch v1109 & int32(7) {
	case 0:
		goto L354
	case 1:
		goto L353
	case 2:
		goto L352
	case 3:
		goto L351
	case 4:
		goto L350
	default:
		v1126 = int32(0)
		goto L349
	}
L349:
	;
	v1131 = m.G0
	v1133 = v1131 - int32(144)
	m.G0 = v1133
	v1136 = v1133 + int32(32)
	F_sha256_init(m, v1136)
	mBase = m.M
	F_sha256_update(m, v1136, v1099, v1126)
	mBase = m.M
	F_sha256_final(m, v1136, v1133)
	mBase = m.M
	goto L356
L350:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1099+int32(-17))))
	v1126 = v1125
	goto L349
L351:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1099+int32(-9))))
	v1126 = v1122
	goto L349
L352:
	;
	v1119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1099+int32(-5)))))
	v1126 = v1119
	goto L349
L353:
	;
	v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099+int32(-3)))))
	v1126 = v1116
	goto L349
L354:
	;
	v1126 = int32(base.Ui32(v1109) >> (uint(int32(3)) % 32))
	goto L349
L355:
	;
	v1154 = int32(0)
	v1155 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(160))))
	*(*int64)(unsafe.Add(mBase, _consts[341])) = v1155
	v1160 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(168))))
	*(*int64)(unsafe.Add(mBase, _consts[342])) = v1160
	goto L359
L356:
	;
	goto L358
L358:
	;
	v1148 = F___memcpy(m, v15+int32(160), v1133, int32(16))
	mBase = m.M
	m.G0 = v1133 + int32(144)
	goto L355
L359:
	;
	goto L347
L360:
	;
	*(*int32)(unsafe.Add(mBase, _consts[842])) = v1167
	v1170 = int32(0)
	v1171 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	v1176 = base.B2i32(v1171 != v1170) & base.B2i32(v1167 == v1170)
	if v1176 != int32(1) {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v1184 {
		goto L364
	} else {
		goto L365
	}
L362:
	;
	F_daemonize(m)
	mBase = m.M
	v1181 = F___syscall_getpid(m)
	mBase = m.M
	goto L363
L363:
	;
	*(*int32)(unsafe.Add(mBase, _consts[154])) = v1181
	goto L361
L364:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if l0 != int32(1) {
		goto L374
	} else {
		goto L375
	}
L365:
	;
	F__serverLog(m, int32(2), int32(_a2311), int32(0))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L16
	} else {
		goto L366
	}
L366:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v1193 {
		goto L364
	} else {
		goto L367
	}
L367:
	;
	goto L368
L368:
	;
	goto L369
L369:
	;
	v1201 = F_strtox_2(m, int32(_a1908), int32(0), int32(10), int64(2147483648))
	mBase = m.M
	goto L370
L370:
	;
	v1205 = F___syscall_getpid(m)
	mBase = m.M
	goto L371
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(144)))) = v1205
	*(*int32)(unsafe.Add(mBase, uint32(v15)+140)) = base.B2i32(int32(0) < base.I32_wrap_i64(v1201))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = int32(_a2283)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = int32(_a469)
	F__serverLog(m, int32(2), int32(_a2312), v15+int32(128))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L16
	} else {
		goto L372
	}
L372:
	;
	goto L364
L373:
	;
	F_initServer(m)
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L16
	} else {
		goto L380
	}
L374:
	;
	if int32(2) < v1224 {
		goto L373
	} else {
		goto L378
	}
L375:
	;
	if int32(3) < v1224 {
		goto L373
	} else {
		goto L376
	}
L376:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v1229
	F__serverLog(m, int32(3), int32(_a2313), v15+int32(112))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L16
	} else {
		goto L377
	}
L377:
	;
	goto L373
L378:
	;
	F__serverLog(m, int32(2), int32(_a2314), int32(0))
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L16
	} else {
		goto L379
	}
L379:
	;
	goto L373
L380:
	;
	v1246 = int32(0)
	v1247 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	if v1176|base.B2i32(v1247 != v1246) != int32(1) {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	F_serverAsciiArt(m)
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L16
	} else {
		goto L384
	}
L382:
	;
	F_createPidFile(m)
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L16
	} else {
		goto L383
	}
L383:
	;
	goto L381
L384:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	if v1258 < int32(129) {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1274 = int32(0)
	v1275 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v1275 == v1274 {
		goto L389
	} else {
		goto L390
	}
L386:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v1262 {
		goto L385
	} else {
		goto L387
	}
L387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v1258
	F__serverLog(m, int32(3), int32(_a2315), v15+int32(96))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L16
	} else {
		goto L388
	}
L388:
	;
	goto L385
L389:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	if v1281 != 0 {
		goto L392
	} else {
		goto L393
	}
L390:
	;
	F_clusterInit(m)
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L16
	} else {
		goto L391
	}
L391:
	;
	goto L389
L392:
	;
	F_ACLLoadUsersAtStartup(m)
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L16
	} else {
		goto L396
	}
L393:
	;
	goto L394
L394:
	;
	F_moduleLoadFromQueue(m)
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L16
	} else {
		goto L395
	}
L395:
	;
	goto L392
L396:
	;
	F_initListeners(m)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L16
	} else {
		goto L397
	}
L397:
	;
	v1288 = int32(0)
	v1289 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v1289 == v1288 {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	if v1295 != 0 {
		goto L401
	} else {
		goto L402
	}
L399:
	;
	F_clusterInitLast(m)
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L16
	} else {
		goto L400
	}
L400:
	;
	goto L398
L401:
	;
	v1305 = int32(0)
	v1307 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	*(*int32)(unsafe.Add(mBase, _consts[844])) = v1307
	F_bioInit(m)
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L16
	} else {
		goto L407
	}
L402:
	;
	v1297 = F_scriptingEngineManagerFind(m, int32(_a2316))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L16
	} else {
		goto L403
	}
L403:
	;
	if v1297 != 0 {
		goto L401
	} else {
		goto L404
	}
L404:
	;
	v1300 = int32(0)
	v1303 = F_moduleLoadStatic(m, int32(_a2316), v1300, v1300, v1300)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L16
	} else {
		goto L405
	}
L405:
	;
	if v1303 != 0 {
		goto L167
	} else {
		goto L406
	}
L406:
	;
	goto L401
L407:
	;
	F_initIOThreads(m, int32(1))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L16
	} else {
		goto L408
	}
L408:
	;
	goto L409
L409:
	;
	v1316 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[845])) = v1316
	v1319 = F_getMemoryOverheadData(m)
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L16
	} else {
		goto L410
	}
L410:
	;
	v1321 = int32(0)
	v1331 = *(*int32)(unsafe.Add(mBase, _consts[315]))
	if v1331 < int32(261) {
		goto L414
	} else {
		goto L415
	}
L411:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1319)+52))
	*(*int32)(unsafe.Add(mBase, _consts[845])) = v1408 - v1415
	F_freeMemoryOverheadData(m, v1319)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L16
	} else {
		goto L427
	}
L412:
	;
	goto L411
L413:
	;
	v1342 = v1340 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v1340) {
		goto L418
	} else {
		goto L419
	}
L414:
	;
	if v1331 < int32(1) {
		v1408 = v1321
		goto L412
	} else {
		goto L416
	}
L415:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	v1339 = v1335
	v1340 = int32(260)
	goto L413
L416:
	;
	v1339 = v1321
	v1340 = v1331
	goto L413
L417:
	;
	if v1342 == int32(0) {
		v1408 = v1381
		goto L412
	} else {
		goto L423
	}
L418:
	;
	v1349 = int32(0)
	v1351 = v1339
	v1352 = v1349
	v1356 = v1349
	goto L420
L419:
	;
	v1381 = v1339
	v1382 = int32(0)
	goto L417
L420:
	;
	v1359 = v1352 << (uint(int32(2)) % 32)
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v1359)+uint32(_consts[317])))
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1359)+uint32(_consts[318])))
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1359)+uint32(_consts[319])))
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v1359)+uint32(_consts[320])))
	v1375 = v1362 + (v1365 + (v1368 + (v1371 + v1351)))
	v1376 = int32(4)
	v1377 = v1352 + v1376
	v1379 = v1356 + v1376
	if v1379 != v1340&int32(2147483644) {
		v1351 = v1375
		v1352 = v1377
		v1356 = v1379
		goto L420
	} else {
		goto L422
	}
L421:
	;
	v1381 = v1375
	v1382 = v1377
	goto L417
L422:
	;
	goto L421
L423:
	;
	v1390 = v1381
	v1391 = v1382
	v1393 = int32(0)
	goto L424
L424:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1391<<(uint(int32(2))%32))+uint32(_consts[320])))
	v1402 = v1401 + v1390
	v1403 = int32(1)
	v1406 = v1393 + v1403
	if v1406 != v1342 {
		v1390 = v1402
		v1391 = v1391 + v1403
		v1393 = v1406
		goto L424
	} else {
		goto L426
	}
L425:
	;
	v1408 = v1402
	goto L412
L426:
	;
	goto L425
L427:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	if v1421 != 0 {
		goto L429
	} else {
		goto L430
	}
L428:
	;
	v1564 = *(*int64)(unsafe.Add(mBase, _consts[321]))
	if base.Ui64(v1564+int64(-1048576)) < base.Ui64(int64(-1048575)) {
		goto L470
	} else {
		goto L471
	}
L429:
	;
	F_sentinelIsRunning(m)
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L16
	} else {
		goto L469
	}
L430:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v1423 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	F_aofLoadManifestFromDisk(m)
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L16
	} else {
		goto L434
	}
L432:
	;
	F__serverLog(m, int32(2), int32(_a2317), int32(0))
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L16
	} else {
		goto L433
	}
L433:
	;
	goto L431
L434:
	;
	F_loadDataFromDisk(m)
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L16
	} else {
		goto L435
	}
L435:
	;
	F_aofOpenIfNeededOnServerStart(m)
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L16
	} else {
		goto L436
	}
L436:
	;
	v1437 = F_aofDelHistoryFiles(m)
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L16
	} else {
		goto L437
	}
L437:
	;
	v1439 = int32(0)
	v1440 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v1440 == v1439 {
		goto L438
	} else {
		goto L439
	}
L438:
	;
	v1445 = int32(0)
	v1446 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v1448 = *(*int32)(unsafe.Add(mBase, _consts[846]))
	if v1448 == v1445 {
		v1474 = v1446
		goto L442
	} else {
		goto L443
	}
L439:
	;
	v1443 = F_verifyClusterConfigWithData(m)
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L16
	} else {
		goto L440
	}
L440:
	;
	if v1443 != 0 {
		goto L166
	} else {
		goto L441
	}
L441:
	;
	goto L438
L442:
	;
	v1476 = int32(0)
	v1477 = *(*int32)(unsafe.Add(mBase, _consts[847]))
	if v1477 == v1476 {
		v1503 = v1474
		goto L449
	} else {
		goto L450
	}
L443:
	;
	if int32(2) < v1446 {
		v1474 = v1446
		goto L442
	} else {
		goto L444
	}
L444:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1448)))
	v1455 = m.T0[v1454].(func(*base.Module) int32)(m)
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L16
	} else {
		goto L446
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v1464
	F__serverLog(m, int32(2), int32(_a2318), v15+int32(80))
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L16
	} else {
		goto L448
	}
L446:
	;
	if base.Ui32(int32(3)) < base.Ui32(v1455) {
		v1464 = int32(_a265)
		goto L445
	} else {
		goto L447
	}
L447:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1455<<(uint(int32(2))%32))+uint32(_consts[769])))
	v1464 = v1463
	goto L445
L448:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v1474 = v1473
	goto L442
L449:
	;
	v1505 = int32(0)
	v1506 = *(*int32)(unsafe.Add(mBase, _consts[848]))
	if v1506 == v1505 {
		v1532 = v1503
		goto L456
	} else {
		goto L457
	}
L450:
	;
	if int32(2) < v1474 {
		v1503 = v1474
		goto L449
	} else {
		goto L451
	}
L451:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1477)))
	v1484 = m.T0[v1483].(func(*base.Module) int32)(m)
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L16
	} else {
		goto L453
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v1493
	F__serverLog(m, int32(2), int32(_a2318), v15+int32(64))
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L16
	} else {
		goto L455
	}
L453:
	;
	if base.Ui32(int32(3)) < base.Ui32(v1484) {
		v1493 = int32(_a265)
		goto L452
	} else {
		goto L454
	}
L454:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1484<<(uint(int32(2))%32))+uint32(_consts[769])))
	v1493 = v1492
	goto L452
L455:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v1503 = v1502
	goto L449
L456:
	;
	v1534 = int32(0)
	v1535 = *(*int32)(unsafe.Add(mBase, _consts[849]))
	if v1535 == v1534 {
		goto L428
	} else {
		goto L463
	}
L457:
	;
	if int32(2) < v1503 {
		v1532 = v1503
		goto L456
	} else {
		goto L458
	}
L458:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1506)))
	v1513 = m.T0[v1512].(func(*base.Module) int32)(m)
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L16
	} else {
		goto L460
	}
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v1522
	F__serverLog(m, int32(2), int32(_a2318), v15+int32(48))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L16
	} else {
		goto L462
	}
L460:
	;
	if base.Ui32(int32(3)) < base.Ui32(v1513) {
		v1522 = int32(_a265)
		goto L459
	} else {
		goto L461
	}
L461:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1513<<(uint(int32(2))%32))+uint32(_consts[769])))
	v1522 = v1521
	goto L459
L462:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v1532 = v1531
	goto L456
L463:
	;
	if int32(2) < v1532 {
		goto L428
	} else {
		goto L464
	}
L464:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v1535)))
	v1542 = m.T0[v1541].(func(*base.Module) int32)(m)
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L16
	} else {
		goto L466
	}
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v1551
	F__serverLog(m, int32(2), int32(_a2318), v15+int32(32))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L16
	} else {
		goto L468
	}
L466:
	;
	if base.Ui32(int32(3)) < base.Ui32(v1542) {
		v1551 = int32(_a265)
		goto L465
	} else {
		goto L467
	}
L467:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1542<<(uint(int32(2))%32))+uint32(_consts[769])))
	v1551 = v1550
	goto L465
L468:
	;
	goto L428
L469:
	;
	goto L428
L470:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	F_aeMain(m, v1581)
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L16
	} else {
		goto L474
	}
L471:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v1570 {
		goto L470
	} else {
		goto L472
	}
L472:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v1564
	F__serverLog(m, int32(3), int32(_a2319), v15+int32(16))
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L16
	} else {
		goto L473
	}
L473:
	;
	goto L470
L474:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	F_aeDeleteEventLoop(m, v1585)
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L16
	} else {
		goto L475
	}
L475:
	;
	m.G0 = v15 + int32(192)
	return int32(0)
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v1593
	v1597 = F_iprintf(m, int32(_a2320), v15)
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L16
	} else {
		goto L477
	}
L477:
	;
	F_sdsfree(m, v1593)
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L16
	} else {
		goto L478
	}
L478:
	;
	m.Env.Exit(m, int32(0))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L479:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L480:
	;
	F_memtest(m, v1654, int32(50))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L16
	} else {
		goto L495
	}
L481:
	;
	v1614 = v1609 + int32(1)
	v1615 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1609))))
	v1616 = F___isspace_1(m, v1615)
	mBase = m.M
	if v1616 != 0 {
		v1609 = v1614
		goto L481
	} else {
		goto L483
	}
L482:
	;
	v1617 = int32(1)
	switch v1615&int32(255) + int32(-43) {
	case 0:
		v1623 = v1617
		goto L485
	default:
		v1625 = v1609
		v1626 = v1615
		v1627 = v1617
		goto L484
	case 2:
		goto L486
	}
L483:
	;
	goto L482
L484:
	;
	v1630 = v1626 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1630) {
		v1648 = int32(0)
		goto L487
	} else {
		goto L488
	}
L485:
	;
	v1624 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1614))))
	v1625 = v1614
	v1626 = v1624
	v1627 = v1623
	goto L484
L486:
	;
	v1623 = int32(0)
	goto L485
L487:
	;
	if v1627 != 0 {
		goto L492
	} else {
		goto L493
	}
L488:
	;
	v1634 = int32(0)
	v1635 = v1625
	v1636 = v1630
	goto L489
L489:
	;
	v1638 = int32(10)
	v1640 = v1634*v1638 - v1636
	v1641 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1635)+1)))
	v1645 = v1641 + int32(-48)
	if base.Ui32(v1645) < base.Ui32(v1638) {
		v1634 = v1640
		v1635 = v1635 + int32(1)
		v1636 = v1645
		goto L489
	} else {
		goto L491
	}
L490:
	;
	v1648 = v1640
	goto L487
L491:
	;
	goto L490
L492:
	;
	v1654 = int32(0) - v1648
	goto L494
L493:
	;
	v1654 = v1648
	goto L494
L494:
	;
	goto L480
L495:
	;
	m.Env.Exit(m, int32(0))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L496:
	;
	m.Env.Exit(m, base.B2i32(v1660 == int32(0)))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L497:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L498:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_makePath(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = F_sdsempty(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		v16 = F_sdscatfmt(m, v9, int32(_a2469), v7)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v16
		}
	}
}
func F_markroot(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
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
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
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
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+44)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v4)+36)) = int64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+112))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)))
	if v10&int32(3) == v2 {
		v83 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+80))
	if v84 < int32(4) {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)))
	v19 = v9
	v20 = v17
	goto L11
L3:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v4)+112))
	v83 = v82
	goto L1
L4:
	;
	goto L3
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+36)) = v19
	goto L4
L6:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v73
	goto L5
L7:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v71
	goto L5
L8:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v69
	goto L5
L9:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v67
	goto L5
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v47 < int32(4) {
		v58 = v46
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v23 = v20 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)) = uint8(v23)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+4)))
	if v25 == int32(7) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v31 = v23 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)) = uint8(v31)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v33 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	switch v25 + int32(-5) {
	case 0:
		goto L8
	case 1:
		goto L9
	default:
		goto L4
	case 3:
		goto L7
	case 4:
		goto L6
	case 5:
		goto L10
	}
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+5)))
	if v43&int32(3) != 0 {
		v19 = v42
		v20 = v43
		goto L11
	} else {
		goto L18
	}
L16:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+5)))
	if v36&int32(3) == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_reallymarkobject(m, v4, v33)
	mBase = m.M
	goto L15
L18:
	;
	goto L4
L19:
	;
	if v58 != v19+int32(16) {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+5)))
	if v51&int32(3) == int32(0) {
		v58 = v46
		goto L19
	} else {
		goto L21
	}
L21:
	;
	F_reallymarkobject(m, v4, v50)
	mBase = m.M
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v58 = v57
	goto L19
L22:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)))
	v65 = v63 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)) = uint8(v65)
	goto L3
L23:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+104))
	if v162 < int32(4) {
		goto L46
	} else {
		goto L47
	}
L24:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)+72))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+5)))
	if v88&int32(3) == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+5)))
	v97 = v87
	v98 = v95
	goto L34
L26:
	;
	goto L23
L27:
	;
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+36)) = v97
	goto L27
L29:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+68)) = v151
	goto L28
L30:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+108)) = v149
	goto L28
L31:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+32)) = v147
	goto L28
L32:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v145
	goto L28
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	if v125 < int32(4) {
		v136 = v124
		goto L42
	} else {
		goto L43
	}
L34:
	;
	v101 = v98 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)) = uint8(v101)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v103 == int32(7) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v109 = v101 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)) = uint8(v109)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	if v111 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	switch v103 + int32(-5) {
	case 0:
		goto L31
	case 1:
		goto L32
	default:
		goto L27
	case 3:
		goto L30
	case 4:
		goto L29
	case 5:
		goto L33
	}
L38:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+5)))
	if v121&int32(3) != 0 {
		v97 = v120
		v98 = v121
		goto L34
	} else {
		goto L41
	}
L39:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+5)))
	if v114&int32(3) == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	F_reallymarkobject(m, v4, v111)
	mBase = m.M
	goto L38
L41:
	;
	goto L27
L42:
	;
	if v136 != v97+int32(16) {
		goto L27
	} else {
		goto L45
	}
L43:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+5)))
	if v129&int32(3) == int32(0) {
		v136 = v124
		goto L42
	} else {
		goto L44
	}
L44:
	;
	F_reallymarkobject(m, v4, v128)
	mBase = m.M
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v136 = v135
	goto L42
L45:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)))
	v143 = v141 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)) = uint8(v143)
	goto L26
L46:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v4)+152))
	if v240 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L47:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161)+96))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
	if v166&int32(3) == int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
	v175 = v165
	v176 = v173
	goto L57
L49:
	;
	goto L46
L50:
	;
	goto L49
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+36)) = v175
	goto L50
L52:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+68)) = v229
	goto L51
L53:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+108)) = v227
	goto L51
L54:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+32)) = v225
	goto L51
L55:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+8)) = v223
	goto L51
L56:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v175)+8))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	if v203 < int32(4) {
		v214 = v202
		goto L65
	} else {
		goto L66
	}
L57:
	;
	v179 = v176 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v175)+5)) = uint8(v179)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
	if v181 == int32(7) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v187 = v179 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v175)+5)) = uint8(v187)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v175)+8))
	if v189 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	switch v181 + int32(-5) {
	case 0:
		goto L54
	case 1:
		goto L55
	default:
		goto L50
	case 3:
		goto L53
	case 4:
		goto L52
	case 5:
		goto L56
	}
L61:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+5)))
	if v199&int32(3) != 0 {
		v175 = v198
		v176 = v199
		goto L57
	} else {
		goto L64
	}
L62:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+5)))
	if v192&int32(3) == int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	F_reallymarkobject(m, v4, v189)
	mBase = m.M
	goto L61
L64:
	;
	goto L50
L65:
	;
	if v214 != v175+int32(16) {
		goto L50
	} else {
		goto L68
	}
L66:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+5)))
	if v207&int32(3) == int32(0) {
		v214 = v202
		goto L65
	} else {
		goto L67
	}
L67:
	;
	F_reallymarkobject(m, v4, v206)
	mBase = m.M
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v175)+8))
	v214 = v213
	goto L65
L68:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+5)))
	v221 = v219 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v175)+5)) = uint8(v221)
	goto L49
L69:
	;
	v321 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4)+21)) = uint8(v321)
	return
L70:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v4)+156))
	if v249 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+5)))
	if v243&int32(3) == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	F_reallymarkobject(m, v4, v240)
	mBase = m.M
	goto L70
L73:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v4)+160))
	if v258 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+5)))
	if v252&int32(3) == int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	F_reallymarkobject(m, v4, v249)
	mBase = m.M
	goto L73
L76:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v4)+164))
	if v267 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+5)))
	if v261&int32(3) == int32(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	F_reallymarkobject(m, v4, v258)
	mBase = m.M
	goto L76
L79:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v4)+168))
	if v276 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+5)))
	if v270&int32(3) == int32(0) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	F_reallymarkobject(m, v4, v267)
	mBase = m.M
	goto L79
L82:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v4)+172))
	if v285 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+5)))
	if v279&int32(3) == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	F_reallymarkobject(m, v4, v276)
	mBase = m.M
	goto L82
L85:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v4)+176))
	if v294 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+5)))
	if v288&int32(3) == int32(0) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	F_reallymarkobject(m, v4, v285)
	mBase = m.M
	goto L85
L88:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v4)+180))
	if v303 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+5)))
	if v297&int32(3) == int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	F_reallymarkobject(m, v4, v294)
	mBase = m.M
	goto L88
L91:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v4)+184))
	if v312 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+5)))
	if v306&int32(3) == int32(0) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	F_reallymarkobject(m, v4, v303)
	mBase = m.M
	goto L91
L94:
	;
	goto L69
L95:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+5)))
	if v315&int32(3) == int32(0) {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	F_reallymarkobject(m, v4, v312)
	mBase = m.M
	goto L94
}
func F_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
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
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var __phi219 int32
	_ = __phi219
	var v226 int32
	_ = v226
	var __phi226 int32
	_ = __phi226
	var v227 int32
	_ = v227
	var __phi227 int32
	_ = __phi227
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
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
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v508 int32
	_ = v508
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v569 int32
	_ = v569
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v596 int32
	_ = v596
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v619 int32
	_ = v619
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v656 int32
	_ = v656
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v716 int32
	_ = v716
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v739 int32
	_ = v739
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v796 int32
	_ = v796
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v819 int32
	_ = v819
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v920 int32
	_ = v920
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v946 int32
	_ = v946
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v969 int32
	_ = v969
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v998 int32
	_ = v998
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1024 int32
	_ = v1024
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1047 int32
	_ = v1047
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1091 int32
	_ = v1091
	var v1116 int32
	_ = v1116
	var v1126 int32
	_ = v1126
	var v1147 int32
	_ = v1147
	var v1158 int32
	_ = v1158
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1195 int32
	_ = v1195
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1221 int32
	_ = v1221
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1244 int32
	_ = v1244
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1273 int32
	_ = v1273
	var v1279 int32
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1299 int32
	_ = v1299
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1322 int32
	_ = v1322
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1366 int32
	_ = v1366
	var v1391 int32
	_ = v1391
	var v1401 int32
	_ = v1401
	var v1422 int32
	_ = v1422
	var v1433 int32
	_ = v1433
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1504 int32
	_ = v1504
	var v1516 int32
	_ = v1516
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1563 int32
	_ = v1563
	var v1566 int32
	_ = v1566
	var v1575 int32
	_ = v1575
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1598 int32
	_ = v1598
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1638 int32
	_ = v1638
	var v1642 int32
	_ = v1642
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1655 int32
	_ = v1655
	var v1664 int32
	_ = v1664
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1687 int32
	_ = v1687
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1702 int32
	_ = v1702
	var v1705 int32
	_ = v1705
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1715 int32
	_ = v1715
	var v1727 int32
	_ = v1727
	v21 = l0 + int32(16)
	v23 = l0 + int32(20)
	v25 = l1
	v26 = l2
	goto L1
L1:
	;
	v44 = v25 + int32(1)
	v49 = v26
	goto L4
L2:
	;
	return v1727
L3:
	;
	goto L2
L4:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v66 == int32(37) {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	v1727 = v1521
	goto L3
L6:
	;
	v1521 = int32(0)
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v25 == v1523 {
		v1526 = v1521
		goto L412
	} else {
		goto L413
	}
L7:
	;
	v1447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+3)))
	v1449 = base.B2i32(v1447 == int32(94))
	if v1447 == int32(94) {
		goto L394
	} else {
		goto L395
	}
L8:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	if v666 == int32(63) {
		goto L180
	} else {
		goto L181
	}
L9:
	;
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	v552 = base.B2i32(v550 == int32(94))
	if v550 == int32(94) {
		goto L145
	} else {
		goto L146
	}
L10:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v25 == v547 {
		goto L142
	} else {
		goto L143
	}
L11:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v486 = base.B2i32(base.Ui32(v485) <= base.Ui32(v25))
	if v486 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L12:
	;
	v406 = v49 + int32(1)
	if v66 == int32(91) {
		goto L100
	} else {
		goto L101
	}
L13:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if v400 == int32(0) {
		goto L10
	} else {
		goto L99
	}
L14:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if v189 == int32(102) {
		goto L43
	} else {
		goto L44
	}
L15:
	;
	switch v66 + int32(-36) {
	case 0:
		goto L13
	case 1, 2, 3:
		goto L12
	case 4:
		goto L17
	case 5:
		goto L16
	default:
		goto L18
	}
L16:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v148 = v139
	goto L35
L17:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if v72 != int32(41) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v66 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	return v25
L20:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v109 < int32(32) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v77 < int32(32) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v92 = v21 + v77<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v25
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v77 + int32(1)
	v99 = F_match(m, l0, v25, v49+int32(2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L24
	} else {
		goto L26
	}
L23:
	;
	v80 = m.G3
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v85 = F_luaL_error(m, v81, v80+int32(_a2727), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	goto L22
L26:
	;
	if v99 != 0 {
		v1727 = v99
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v101 + int32(-1)
	return int32(0)
L28:
	;
	v122 = v21 + v109<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v25
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v109 + int32(1)
	v129 = F_match(m, l0, v25, v49+int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L24
	} else {
		goto L31
	}
L29:
	;
	v112 = m.G3
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v117 = F_luaL_error(m, v113, v112+int32(_a2727), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	if v129 != 0 {
		v1727 = v129
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v131 + int32(-1)
	return int32(0)
L33:
	;
	v179 = v21 + v176<<(uint(int32(3))%32)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = v25 - v180
	v183 = F_match(m, l0, v25, v49+int32(1))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L24
	} else {
		goto L40
	}
L34:
	;
	v169 = m.G3
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v174 = F_luaL_error(m, v170, v169+int32(_a2728), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L24
	} else {
		goto L39
	}
L35:
	;
	if v148 < int32(1) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v161 = int32(-1)
	v162 = v148 + v161
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v23+v162<<(uint(int32(3))%32))))
	if v166 != v161 {
		v148 = v162
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v176 = v162
	goto L33
L39:
	;
	v176 = v174
	goto L33
L40:
	;
	if v183 != 0 {
		v1727 = v183
		goto L3
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = int32(-1)
	return int32(0)
L42:
	;
	if base.Ui32((v189+int32(-48))&int32(255)) < base.Ui32(int32(10)) {
		goto L69
	} else {
		goto L70
	}
L43:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+2)))
	if v253 != int32(91) {
		goto L60
	} else {
		goto L61
	}
L44:
	;
	if v189 != int32(98) {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+2)))
	if v194 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v207 = int32(0)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v210 = v206 & int32(255)
	if v208 != v210 {
		v1727 = v207
		goto L3
	} else {
		goto L51
	}
L47:
	;
	v198 = m.G3
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v203 = F_luaL_error(m, v199, v198+int32(_a2729), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L24
	} else {
		goto L50
	}
L48:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+3)))
	if v197 != 0 {
		v206 = v194
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+2)))
	v206 = v205
	goto L46
L51:
	;
	v212 = int32(1)
	v214 = v25 + v212
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v215) <= base.Ui32(v214) {
		v1727 = v207
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+3)))
	__phi219 = v25
	__phi226 = v214
	__phi227 = v212
	v219 = __phi219
	v226 = __phi226
	v227 = __phi227
	goto L53
L53:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)))
	if v237 != v217&int32(255) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v251 = v226 + int32(1)
	if v251 != v215 {
		__phi219 = v226
		__phi226 = v251
		__phi227 = v249
		v219 = __phi219
		v226 = __phi226
		v227 = __phi227
		goto L53
	} else {
		goto L59
	}
L56:
	;
	v249 = v227 + base.B2i32(v237 == v210)
	goto L55
L57:
	;
	v242 = v227 + int32(-1)
	if v242 != 0 {
		v249 = v242
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v25 = v219 + int32(2)
	v26 = v49 + int32(4)
	goto L1
L59:
	;
	v1727 = v207
	goto L3
L60:
	;
	v258 = m.G3
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v263 = F_luaL_error(m, v259, v258+int32(_a2730), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L24
	} else {
		goto L62
	}
L61:
	;
	v1444 = v49 + int32(3)
	goto L7
L62:
	;
	v266 = v49 + int32(3)
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+2)))
	if v267 == int32(91) {
		v1444 = v266
		goto L7
	} else {
		goto L63
	}
L63:
	;
	if v267 == int32(37) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	if v272 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v1504 = v266
	v1516 = v266
	goto L6
L66:
	;
	v1504 = v49 + int32(4)
	v1516 = v266
	goto L6
L67:
	;
	v273 = m.G3
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v278 = F_luaL_error(m, v274, v273+int32(_a2731), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L24
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	if base.Ui32(v189) < base.Ui32(int32(49)) {
		goto L75
	} else {
		goto L76
	}
L70:
	;
	if v189 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v468 = v49 + int32(2)
	v480 = v49 + int32(1)
	goto L11
L72:
	;
	v289 = m.G3
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v294 = F_luaL_error(m, v290, v289+int32(_a2731), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L24
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v322 = int32(0)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v327 = v21 + v321<<(uint(int32(3))%32)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if base.Ui32(v323-v25) < base.Ui32(v328) {
		v1727 = v322
		goto L3
	} else {
		goto L80
	}
L75:
	;
	v314 = m.G3
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v319 = F_luaL_error(m, v315, v314+int32(_a2732), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L24
	} else {
		goto L79
	}
L76:
	;
	v304 = v189 + int32(-49)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v305 <= v304 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v23+v304<<(uint(int32(3))%32))))
	if v310 != int32(-1) {
		v321 = v304
		goto L74
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	v321 = v319
	goto L74
L80:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	if base.Ui32(v328) < base.Ui32(int32(4)) {
		v354 = v330
		v355 = v25
		v356 = v328
		goto L84
	} else {
		goto L85
	}
L81:
	;
	if v394 != 0 {
		v1727 = v322
		goto L3
	} else {
		goto L97
	}
L82:
	;
	v394 = int32(0)
	goto L81
L83:
	;
	v366 = v361
	v367 = v362
	v368 = v363
	goto L93
L84:
	;
	if v356 == int32(0) {
		goto L82
	} else {
		goto L91
	}
L85:
	;
	if (v25|v330)&int32(3) != 0 {
		v361 = v330
		v362 = v25
		v363 = v328
		goto L83
	} else {
		goto L86
	}
L86:
	;
	v338 = v330
	v339 = v25
	v340 = v328
	goto L87
L87:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	if v343 != v344 {
		v361 = v338
		v362 = v339
		v363 = v340
		goto L83
	} else {
		goto L89
	}
L88:
	;
	v354 = v349
	v355 = v347
	v356 = v351
	goto L84
L89:
	;
	v346 = int32(4)
	v347 = v339 + v346
	v349 = v338 + v346
	v351 = v340 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v351) {
		v338 = v349
		v339 = v347
		v340 = v351
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v361 = v354
	v362 = v355
	v363 = v356
	goto L83
L92:
	;
	v394 = v371 - v372
	goto L81
L93:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367))))
	if v371 != v372 {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	v374 = int32(1)
	v379 = v368 + int32(-1)
	if v379 == int32(0) {
		goto L82
	} else {
		goto L96
	}
L96:
	;
	v366 = v366 + v374
	v367 = v367 + v374
	v368 = v379
	goto L93
L97:
	;
	if v25 == int32(0) {
		v1727 = v322
		goto L3
	} else {
		goto L98
	}
L98:
	;
	v25 = v25 + v328
	v26 = v49 + int32(2)
	goto L1
L99:
	;
	v404 = v49 + int32(1)
	v468 = v404
	v480 = v404
	goto L11
L100:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	v413 = base.B2i32(v411 == int32(94))
	if v411 == int32(94) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v468 = v406
	v480 = v406
	goto L11
L102:
	;
	v414 = v49 + int32(2)
	goto L104
L103:
	;
	v414 = v406
	goto L104
L104:
	;
	if v411 == int32(94) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v417 = int32(2)
	goto L107
L106:
	;
	v417 = int32(1)
	goto L107
L107:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v417))))
	v428 = v414
	v429 = v419
	goto L108
L108:
	;
	if v429&int32(255) != 0 {
		v449 = v429
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v468 = v460 + int32(1)
	v480 = v406
	goto L11
L110:
	;
	v451 = v428 + int32(1)
	if v449&int32(255) == int32(37) {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	v441 = m.G3
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v446 = F_luaL_error(m, v442, v441+int32(_a2733), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L24
	} else {
		goto L112
	}
L112:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428))))
	v449 = v448
	goto L110
L113:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	if v461 != int32(93) {
		v428 = v460
		v429 = v461
		goto L108
	} else {
		goto L119
	}
L114:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428)+1)))
	if v458 != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v460 = v451
	goto L113
L116:
	;
	v459 = v428 + int32(2)
	goto L118
L117:
	;
	v459 = v451
	goto L118
L118:
	;
	v460 = v459
	goto L113
L119:
	;
	goto L109
L120:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	switch v492 + int32(-37) {
	case 0:
		goto L124
	case 1, 2, 3, 4, 5, 6, 7, 8:
		goto L122
	case 9:
		v656 = int32(1)
		goto L8
	default:
		goto L123
	}
L121:
	;
	v656 = int32(0)
	goto L8
L122:
	;
	v656 = base.B2i32(v490 == v492)
	goto L8
L123:
	;
	if v492 == int32(91) {
		goto L9
	} else {
		goto L141
	}
L124:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	v496 = F_tolower(m, v495)
	mBase = m.M
	switch v496 + int32(-97) {
	case 0:
		goto L127
	default:
		goto L128
	case 2:
		goto L137
	case 3:
		goto L136
	case 11:
		goto L135
	case 15:
		goto L134
	case 18:
		goto L133
	case 20:
		goto L132
	case 22:
		goto L131
	case 23:
		goto L130
	case 25:
		goto L129
	}
L125:
	;
	v656 = base.B2i32(v540 != int32(0))
	goto L8
L126:
	;
	if base.Ui32(v495+int32(-97)) < base.Ui32(int32(26)) {
		goto L138
	} else {
		goto L139
	}
L127:
	;
	v531 = base.B2i32(base.Ui32(v490|int32(32)+int32(-97)) < base.Ui32(int32(26)))
	goto L126
L128:
	;
	v540 = base.B2i32(v495 == v490)
	goto L125
L129:
	;
	v531 = base.B2i32(v490 == int32(0))
	goto L126
L130:
	;
	v521 = F_isxdigit(m, v490)
	mBase = m.M
	v531 = v521
	goto L126
L131:
	;
	v520 = F_isalnum(m, v490)
	mBase = m.M
	v531 = v520
	goto L126
L132:
	;
	v531 = base.B2i32(base.Ui32(v490+int32(-65)) < base.Ui32(int32(26)))
	goto L126
L133:
	;
	v531 = base.B2i32(v490 == int32(32)) | base.B2i32(base.Ui32(v490+int32(-9)) < base.Ui32(int32(5)))
	goto L126
L134:
	;
	v508 = F_ispunct(m, v490)
	mBase = m.M
	v531 = v508
	goto L126
L135:
	;
	v531 = base.B2i32(base.Ui32(v490+int32(-97)) < base.Ui32(int32(26)))
	goto L126
L136:
	;
	v531 = base.B2i32(base.Ui32(v490+int32(-48)) < base.Ui32(int32(10)))
	goto L126
L137:
	;
	v499 = F_iscntrl(m, v490)
	mBase = m.M
	v531 = v499
	goto L126
L138:
	;
	v538 = v531
	goto L140
L139:
	;
	v538 = base.B2i32(v531 == int32(0))
	goto L140
L140:
	;
	v540 = v538
	goto L125
L141:
	;
	goto L122
L142:
	;
	v549 = v25
	goto L144
L143:
	;
	v549 = int32(0)
	goto L144
L144:
	;
	v1727 = v549
	goto L3
L145:
	;
	v553 = v480
	goto L147
L146:
	;
	v553 = v49
	goto L147
L147:
	;
	v555 = v553 + int32(1)
	v557 = v468 + int32(-1)
	if base.Ui32(v557) <= base.Ui32(v555) {
		v656 = v552
		goto L8
	} else {
		goto L148
	}
L148:
	;
	v560 = base.B2i32(v550 != int32(94))
	v569 = v553
	v577 = v555
	goto L149
L149:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+2)))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+1)))
	if v581 != int32(37) {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	v656 = v552
	goto L8
L151:
	;
	v645 = v643 + int32(1)
	if base.Ui32(v645) < base.Ui32(v557) {
		v569 = v643
		v577 = v645
		goto L149
	} else {
		goto L178
	}
L152:
	;
	v643 = v569 + int32(2)
	goto L151
L153:
	;
	if v580 != int32(45) {
		goto L172
	} else {
		goto L173
	}
L154:
	;
	v584 = F_tolower(m, v580)
	mBase = m.M
	switch v584 + int32(-97) {
	case 0:
		goto L157
	default:
		goto L158
	case 2:
		goto L167
	case 3:
		goto L166
	case 11:
		goto L165
	case 15:
		goto L164
	case 18:
		goto L163
	case 20:
		goto L162
	case 22:
		goto L161
	case 23:
		goto L160
	case 25:
		goto L159
	}
L155:
	;
	if v628 == int32(0) {
		goto L152
	} else {
		goto L171
	}
L156:
	;
	if base.Ui32(v580+int32(-97)) < base.Ui32(int32(26)) {
		goto L168
	} else {
		goto L169
	}
L157:
	;
	v619 = base.B2i32(base.Ui32(v490|int32(32)+int32(-97)) < base.Ui32(int32(26)))
	goto L156
L158:
	;
	v628 = base.B2i32(v580 == v490)
	goto L155
L159:
	;
	v619 = base.B2i32(v490 == int32(0))
	goto L156
L160:
	;
	v609 = F_isxdigit(m, v490)
	mBase = m.M
	v619 = v609
	goto L156
L161:
	;
	v608 = F_isalnum(m, v490)
	mBase = m.M
	v619 = v608
	goto L156
L162:
	;
	v619 = base.B2i32(base.Ui32(v490+int32(-65)) < base.Ui32(int32(26)))
	goto L156
L163:
	;
	v619 = base.B2i32(v490 == int32(32)) | base.B2i32(base.Ui32(v490+int32(-9)) < base.Ui32(int32(5)))
	goto L156
L164:
	;
	v596 = F_ispunct(m, v490)
	mBase = m.M
	v619 = v596
	goto L156
L165:
	;
	v619 = base.B2i32(base.Ui32(v490+int32(-97)) < base.Ui32(int32(26)))
	goto L156
L166:
	;
	v619 = base.B2i32(base.Ui32(v490+int32(-48)) < base.Ui32(int32(10)))
	goto L156
L167:
	;
	v587 = F_iscntrl(m, v490)
	mBase = m.M
	v619 = v587
	goto L156
L168:
	;
	v626 = v619
	goto L170
L169:
	;
	v626 = base.B2i32(v619 == int32(0))
	goto L170
L170:
	;
	v628 = v626
	goto L155
L171:
	;
	v656 = v560
	goto L8
L172:
	;
	if v490 != v581 {
		v643 = v577
		goto L151
	} else {
		goto L177
	}
L173:
	;
	v634 = v569 + int32(3)
	if base.Ui32(v557) <= base.Ui32(v634) {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	if base.Ui32(v490) < base.Ui32(v581) {
		v643 = v634
		goto L151
	} else {
		goto L175
	}
L175:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
	if base.Ui32(v637) < base.Ui32(v490) {
		v643 = v634
		goto L151
	} else {
		goto L176
	}
L176:
	;
	v656 = v560
	goto L8
L177:
	;
	v656 = v560
	goto L8
L178:
	;
	goto L150
L179:
	;
	if base.Ui32(v485) <= base.Ui32(v25) {
		v1401 = int32(0)
		goto L326
	} else {
		goto L327
	}
L180:
	;
	if v656 == int32(0) {
		goto L322
	} else {
		goto L323
	}
L181:
	;
	switch v666 + int32(-42) {
	case 0:
		goto L179
	case 1:
		goto L182
	default:
		goto L183
	case 3:
		goto L184
	}
L182:
	;
	if v656 != 0 {
		goto L252
	} else {
		goto L253
	}
L183:
	;
	if v656 != 0 {
		goto L250
	} else {
		goto L251
	}
L184:
	;
	v672 = v468 + int32(1)
	v673 = F_match(m, l0, v25, v672)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L24
	} else {
		goto L185
	}
L185:
	;
	if v673 != 0 {
		v1727 = v673
		goto L3
	} else {
		goto L186
	}
L186:
	;
	v676 = v468 + int32(-1)
	v678 = v25
	goto L187
L187:
	;
	v696 = int32(0)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v697) <= base.Ui32(v678) {
		v1727 = v696
		goto L3
	} else {
		goto L189
	}
L189:
	;
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678))))
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	switch v700 + int32(-37) {
	case 0:
		goto L195
	case 1, 2, 3, 4, 5, 6, 7, 8:
		goto L193
	case 9:
		goto L190
	default:
		goto L194
	}
L190:
	;
	v888 = v678 + int32(1)
	v889 = F_match(m, l0, v888, v672)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L24
	} else {
		goto L248
	}
L191:
	;
	if v849 == int32(0) {
		v1727 = v696
		goto L3
	} else {
		goto L247
	}
L192:
	;
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	v754 = base.B2i32(v752 == int32(94))
	if v752 == int32(94) {
		goto L213
	} else {
		goto L214
	}
L193:
	;
	v849 = base.B2i32(v699 == v700)
	goto L191
L194:
	;
	if v700 == int32(91) {
		goto L192
	} else {
		goto L212
	}
L195:
	;
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	v704 = F_tolower(m, v703)
	mBase = m.M
	switch v704 + int32(-97) {
	case 0:
		goto L198
	default:
		goto L199
	case 2:
		goto L208
	case 3:
		goto L207
	case 11:
		goto L206
	case 15:
		goto L205
	case 18:
		goto L204
	case 20:
		goto L203
	case 22:
		goto L202
	case 23:
		goto L201
	case 25:
		goto L200
	}
L196:
	;
	v849 = v748
	goto L191
L197:
	;
	if base.Ui32(v703+int32(-97)) < base.Ui32(int32(26)) {
		goto L209
	} else {
		goto L210
	}
L198:
	;
	v739 = base.B2i32(base.Ui32(v699|int32(32)+int32(-97)) < base.Ui32(int32(26)))
	goto L197
L199:
	;
	v748 = base.B2i32(v703 == v699)
	goto L196
L200:
	;
	v739 = base.B2i32(v699 == int32(0))
	goto L197
L201:
	;
	v729 = F_isxdigit(m, v699)
	mBase = m.M
	v739 = v729
	goto L197
L202:
	;
	v728 = F_isalnum(m, v699)
	mBase = m.M
	v739 = v728
	goto L197
L203:
	;
	v739 = base.B2i32(base.Ui32(v699+int32(-65)) < base.Ui32(int32(26)))
	goto L197
L204:
	;
	v739 = base.B2i32(v699 == int32(32)) | base.B2i32(base.Ui32(v699+int32(-9)) < base.Ui32(int32(5)))
	goto L197
L205:
	;
	v716 = F_ispunct(m, v699)
	mBase = m.M
	v739 = v716
	goto L197
L206:
	;
	v739 = base.B2i32(base.Ui32(v699+int32(-97)) < base.Ui32(int32(26)))
	goto L197
L207:
	;
	v739 = base.B2i32(base.Ui32(v699+int32(-48)) < base.Ui32(int32(10)))
	goto L197
L208:
	;
	v707 = F_iscntrl(m, v699)
	mBase = m.M
	v739 = v707
	goto L197
L209:
	;
	v746 = v739
	goto L211
L210:
	;
	v746 = base.B2i32(v739 == int32(0))
	goto L211
L211:
	;
	v748 = v746
	goto L196
L212:
	;
	goto L193
L213:
	;
	v755 = v480
	goto L215
L214:
	;
	v755 = v49
	goto L215
L215:
	;
	v757 = v755 + int32(1)
	if base.Ui32(v676) <= base.Ui32(v757) {
		v849 = v754
		goto L191
	} else {
		goto L216
	}
L216:
	;
	v760 = base.B2i32(v752 != int32(94))
	v769 = v755
	v773 = v757
	goto L217
L217:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+2)))
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+1)))
	if v781 != int32(37) {
		goto L221
	} else {
		goto L222
	}
L218:
	;
	v849 = v754
	goto L191
L219:
	;
	v845 = v843 + int32(1)
	if base.Ui32(v845) < base.Ui32(v676) {
		v769 = v843
		v773 = v845
		goto L217
	} else {
		goto L246
	}
L220:
	;
	v843 = v769 + int32(2)
	goto L219
L221:
	;
	if v780 != int32(45) {
		goto L240
	} else {
		goto L241
	}
L222:
	;
	v784 = F_tolower(m, v780)
	mBase = m.M
	switch v784 + int32(-97) {
	case 0:
		goto L225
	default:
		goto L226
	case 2:
		goto L235
	case 3:
		goto L234
	case 11:
		goto L233
	case 15:
		goto L232
	case 18:
		goto L231
	case 20:
		goto L230
	case 22:
		goto L229
	case 23:
		goto L228
	case 25:
		goto L227
	}
L223:
	;
	if v828 == int32(0) {
		goto L220
	} else {
		goto L239
	}
L224:
	;
	if base.Ui32(v780+int32(-97)) < base.Ui32(int32(26)) {
		goto L236
	} else {
		goto L237
	}
L225:
	;
	v819 = base.B2i32(base.Ui32(v699|int32(32)+int32(-97)) < base.Ui32(int32(26)))
	goto L224
L226:
	;
	v828 = base.B2i32(v780 == v699)
	goto L223
L227:
	;
	v819 = base.B2i32(v699 == int32(0))
	goto L224
L228:
	;
	v809 = F_isxdigit(m, v699)
	mBase = m.M
	v819 = v809
	goto L224
L229:
	;
	v808 = F_isalnum(m, v699)
	mBase = m.M
	v819 = v808
	goto L224
L230:
	;
	v819 = base.B2i32(base.Ui32(v699+int32(-65)) < base.Ui32(int32(26)))
	goto L224
L231:
	;
	v819 = base.B2i32(v699 == int32(32)) | base.B2i32(base.Ui32(v699+int32(-9)) < base.Ui32(int32(5)))
	goto L224
L232:
	;
	v796 = F_ispunct(m, v699)
	mBase = m.M
	v819 = v796
	goto L224
L233:
	;
	v819 = base.B2i32(base.Ui32(v699+int32(-97)) < base.Ui32(int32(26)))
	goto L224
L234:
	;
	v819 = base.B2i32(base.Ui32(v699+int32(-48)) < base.Ui32(int32(10)))
	goto L224
L235:
	;
	v787 = F_iscntrl(m, v699)
	mBase = m.M
	v819 = v787
	goto L224
L236:
	;
	v826 = v819
	goto L238
L237:
	;
	v826 = base.B2i32(v819 == int32(0))
	goto L238
L238:
	;
	v828 = v826
	goto L223
L239:
	;
	v849 = v760
	goto L191
L240:
	;
	if v699 != v781 {
		v843 = v773
		goto L219
	} else {
		goto L245
	}
L241:
	;
	v834 = v769 + int32(3)
	if base.Ui32(v676) <= base.Ui32(v834) {
		goto L240
	} else {
		goto L242
	}
L242:
	;
	if base.Ui32(v699) < base.Ui32(v781) {
		v843 = v834
		goto L219
	} else {
		goto L243
	}
L243:
	;
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834))))
	if base.Ui32(v837) < base.Ui32(v699) {
		v843 = v834
		goto L219
	} else {
		goto L244
	}
L244:
	;
	v849 = v760
	goto L191
L245:
	;
	v849 = v760
	goto L191
L246:
	;
	goto L218
L247:
	;
	goto L190
L248:
	;
	if v889 == int32(0) {
		v678 = v888
		goto L187
	} else {
		goto L249
	}
L249:
	;
	v1727 = v889
	goto L3
L250:
	;
	v25 = v44
	v26 = v468
	goto L1
L251:
	;
	return int32(0)
L252:
	;
	v899 = v25 + int32(1)
	if base.Ui32(v485) <= base.Ui32(v899) {
		v1126 = int32(0)
		goto L254
	} else {
		goto L255
	}
L253:
	;
	return int32(0)
L254:
	;
	v1147 = v1126
	goto L317
L255:
	;
	v901 = int32(-1)
	v902 = v468 + v901
	v905 = v25 ^ v901 + v485
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v920 = int32(0)
	goto L256
L256:
	;
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v899+v920))))
	switch v907 + int32(-37) {
	case 0:
		goto L263
	case 1, 2, 3, 4, 5, 6, 7, 8:
		goto L261
	case 9:
		goto L258
	default:
		goto L262
	}
L257:
	;
	v1126 = v905
	goto L254
L258:
	;
	v1116 = v920 + int32(1)
	if v1116 != v905 {
		v920 = v1116
		goto L256
	} else {
		goto L316
	}
L259:
	;
	if v1091 == int32(0) {
		v1126 = v920
		goto L254
	} else {
		goto L315
	}
L260:
	;
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	v982 = base.B2i32(v980 == int32(94))
	if v980 == int32(94) {
		goto L281
	} else {
		goto L282
	}
L261:
	;
	v1091 = base.B2i32(v932 == v907)
	goto L259
L262:
	;
	if v907 == int32(91) {
		goto L260
	} else {
		goto L280
	}
L263:
	;
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	v934 = F_tolower(m, v933)
	mBase = m.M
	switch v934 + int32(-97) {
	case 0:
		goto L266
	default:
		goto L267
	case 2:
		goto L276
	case 3:
		goto L275
	case 11:
		goto L274
	case 15:
		goto L273
	case 18:
		goto L272
	case 20:
		goto L271
	case 22:
		goto L270
	case 23:
		goto L269
	case 25:
		goto L268
	}
L264:
	;
	v1091 = v978
	goto L259
L265:
	;
	if base.Ui32(v933+int32(-97)) < base.Ui32(int32(26)) {
		goto L277
	} else {
		goto L278
	}
L266:
	;
	v969 = base.B2i32(base.Ui32(v932|int32(32)+int32(-97)) < base.Ui32(int32(26)))
	goto L265
L267:
	;
	v978 = base.B2i32(v933 == v932)
	goto L264
L268:
	;
	v969 = base.B2i32(v932 == int32(0))
	goto L265
L269:
	;
	v959 = F_isxdigit(m, v932)
	mBase = m.M
	v969 = v959
	goto L265
L270:
	;
	v958 = F_isalnum(m, v932)
	mBase = m.M
	v969 = v958
	goto L265
L271:
	;
	v969 = base.B2i32(base.Ui32(v932+int32(-65)) < base.Ui32(int32(26)))
	goto L265
L272:
	;
	v969 = base.B2i32(v932 == int32(32)) | base.B2i32(base.Ui32(v932+int32(-9)) < base.Ui32(int32(5)))
	goto L265
L273:
	;
	v946 = F_ispunct(m, v932)
	mBase = m.M
	v969 = v946
	goto L265
L274:
	;
	v969 = base.B2i32(base.Ui32(v932+int32(-97)) < base.Ui32(int32(26)))
	goto L265
L275:
	;
	v969 = base.B2i32(base.Ui32(v932+int32(-48)) < base.Ui32(int32(10)))
	goto L265
L276:
	;
	v937 = F_iscntrl(m, v932)
	mBase = m.M
	v969 = v937
	goto L265
L277:
	;
	v976 = v969
	goto L279
L278:
	;
	v976 = base.B2i32(v969 == int32(0))
	goto L279
L279:
	;
	v978 = v976
	goto L264
L280:
	;
	goto L261
L281:
	;
	v983 = v480
	goto L283
L282:
	;
	v983 = v49
	goto L283
L283:
	;
	v985 = v983 + int32(1)
	if base.Ui32(v902) <= base.Ui32(v985) {
		v1091 = v982
		goto L259
	} else {
		goto L284
	}
L284:
	;
	v988 = base.B2i32(v980 != int32(94))
	v990 = v985
	v998 = v983
	goto L285
L285:
	;
	v1008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998)+2)))
	v1009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998)+1)))
	if v1009 != int32(37) {
		goto L289
	} else {
		goto L290
	}
L286:
	;
	v1091 = v982
	goto L259
L287:
	;
	v1073 = v1071 + int32(1)
	if base.Ui32(v1073) < base.Ui32(v902) {
		v990 = v1073
		v998 = v1071
		goto L285
	} else {
		goto L314
	}
L288:
	;
	v1071 = v998 + int32(2)
	goto L287
L289:
	;
	if v1008 != int32(45) {
		goto L308
	} else {
		goto L309
	}
L290:
	;
	v1012 = F_tolower(m, v1008)
	mBase = m.M
	switch v1012 + int32(-97) {
	case 0:
		goto L293
	default:
		goto L294
	case 2:
		goto L303
	case 3:
		goto L302
	case 11:
		goto L301
	case 15:
		goto L300
	case 18:
		goto L299
	case 20:
		goto L298
	case 22:
		goto L297
	case 23:
		goto L296
	case 25:
		goto L295
	}
L291:
	;
	if v1056 == int32(0) {
		goto L288
	} else {
		goto L307
	}
L292:
	;
	if base.Ui32(v1008+int32(-97)) < base.Ui32(int32(26)) {
		goto L304
	} else {
		goto L305
	}
L293:
	;
	v1047 = base.B2i32(base.Ui32(v932|int32(32)+int32(-97)) < base.Ui32(int32(26)))
	goto L292
L294:
	;
	v1056 = base.B2i32(v1008 == v932)
	goto L291
L295:
	;
	v1047 = base.B2i32(v932 == int32(0))
	goto L292
L296:
	;
	v1037 = F_isxdigit(m, v932)
	mBase = m.M
	v1047 = v1037
	goto L292
L297:
	;
	v1036 = F_isalnum(m, v932)
	mBase = m.M
	v1047 = v1036
	goto L292
L298:
	;
	v1047 = base.B2i32(base.Ui32(v932+int32(-65)) < base.Ui32(int32(26)))
	goto L292
L299:
	;
	v1047 = base.B2i32(v932 == int32(32)) | base.B2i32(base.Ui32(v932+int32(-9)) < base.Ui32(int32(5)))
	goto L292
L300:
	;
	v1024 = F_ispunct(m, v932)
	mBase = m.M
	v1047 = v1024
	goto L292
L301:
	;
	v1047 = base.B2i32(base.Ui32(v932+int32(-97)) < base.Ui32(int32(26)))
	goto L292
L302:
	;
	v1047 = base.B2i32(base.Ui32(v932+int32(-48)) < base.Ui32(int32(10)))
	goto L292
L303:
	;
	v1015 = F_iscntrl(m, v932)
	mBase = m.M
	v1047 = v1015
	goto L292
L304:
	;
	v1054 = v1047
	goto L306
L305:
	;
	v1054 = base.B2i32(v1047 == int32(0))
	goto L306
L306:
	;
	v1056 = v1054
	goto L291
L307:
	;
	v1091 = v988
	goto L259
L308:
	;
	if v932 != v1009 {
		v1071 = v990
		goto L287
	} else {
		goto L313
	}
L309:
	;
	v1062 = v998 + int32(3)
	if base.Ui32(v902) <= base.Ui32(v1062) {
		goto L308
	} else {
		goto L310
	}
L310:
	;
	if base.Ui32(v932) < base.Ui32(v1009) {
		v1071 = v1062
		goto L287
	} else {
		goto L311
	}
L311:
	;
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062))))
	if base.Ui32(v1065) < base.Ui32(v932) {
		v1071 = v1062
		goto L287
	} else {
		goto L312
	}
L312:
	;
	v1091 = v988
	goto L259
L313:
	;
	v1091 = v988
	goto L259
L314:
	;
	goto L286
L315:
	;
	goto L258
L316:
	;
	goto L257
L317:
	;
	v1158 = int32(0)
	if v1147 < v1158 {
		v1727 = v1158
		goto L3
	} else {
		goto L319
	}
L319:
	;
	v1164 = F_match(m, l0, v899+v1147, v468+int32(1))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L24
	} else {
		goto L320
	}
L320:
	;
	if v1164 == int32(0) {
		v1147 = v1147 + int32(-1)
		goto L317
	} else {
		goto L321
	}
L321:
	;
	v1727 = v1164
	goto L3
L322:
	;
	v49 = v468 + int32(1)
	goto L4
L323:
	;
	v1172 = F_match(m, l0, v44, v468+int32(1))
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L24
	} else {
		goto L324
	}
L324:
	;
	if v1172 != 0 {
		v1727 = v1172
		goto L3
	} else {
		goto L325
	}
L325:
	;
	goto L322
L326:
	;
	v1422 = v1401
	goto L389
L327:
	;
	v1178 = v485 - v25
	v1180 = v468 + int32(-1)
	v1182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v1195 = int32(0)
	goto L328
L328:
	;
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v1195))))
	switch v1182 + int32(-37) {
	case 0:
		goto L335
	case 1, 2, 3, 4, 5, 6, 7, 8:
		goto L333
	case 9:
		goto L330
	default:
		goto L334
	}
L329:
	;
	v1401 = v1178
	goto L326
L330:
	;
	v1391 = v1195 + int32(1)
	if v1391 != v1178 {
		v1195 = v1391
		goto L328
	} else {
		goto L388
	}
L331:
	;
	if v1366 == int32(0) {
		v1401 = v1195
		goto L326
	} else {
		goto L387
	}
L332:
	;
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	v1257 = base.B2i32(v1255 == int32(94))
	if v1255 == int32(94) {
		goto L353
	} else {
		goto L354
	}
L333:
	;
	v1366 = base.B2i32(v1207 == v1182)
	goto L331
L334:
	;
	if v1182 == int32(91) {
		goto L332
	} else {
		goto L352
	}
L335:
	;
	v1208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	v1209 = F_tolower(m, v1208)
	mBase = m.M
	switch v1209 + int32(-97) {
	case 0:
		goto L338
	default:
		goto L339
	case 2:
		goto L348
	case 3:
		goto L347
	case 11:
		goto L346
	case 15:
		goto L345
	case 18:
		goto L344
	case 20:
		goto L343
	case 22:
		goto L342
	case 23:
		goto L341
	case 25:
		goto L340
	}
L336:
	;
	v1366 = v1253
	goto L331
L337:
	;
	if base.Ui32(v1208+int32(-97)) < base.Ui32(int32(26)) {
		goto L349
	} else {
		goto L350
	}
L338:
	;
	v1244 = base.B2i32(base.Ui32(v1207|int32(32)+int32(-97)) < base.Ui32(int32(26)))
	goto L337
L339:
	;
	v1253 = base.B2i32(v1208 == v1207)
	goto L336
L340:
	;
	v1244 = base.B2i32(v1207 == int32(0))
	goto L337
L341:
	;
	v1234 = F_isxdigit(m, v1207)
	mBase = m.M
	v1244 = v1234
	goto L337
L342:
	;
	v1233 = F_isalnum(m, v1207)
	mBase = m.M
	v1244 = v1233
	goto L337
L343:
	;
	v1244 = base.B2i32(base.Ui32(v1207+int32(-65)) < base.Ui32(int32(26)))
	goto L337
L344:
	;
	v1244 = base.B2i32(v1207 == int32(32)) | base.B2i32(base.Ui32(v1207+int32(-9)) < base.Ui32(int32(5)))
	goto L337
L345:
	;
	v1221 = F_ispunct(m, v1207)
	mBase = m.M
	v1244 = v1221
	goto L337
L346:
	;
	v1244 = base.B2i32(base.Ui32(v1207+int32(-97)) < base.Ui32(int32(26)))
	goto L337
L347:
	;
	v1244 = base.B2i32(base.Ui32(v1207+int32(-48)) < base.Ui32(int32(10)))
	goto L337
L348:
	;
	v1212 = F_iscntrl(m, v1207)
	mBase = m.M
	v1244 = v1212
	goto L337
L349:
	;
	v1251 = v1244
	goto L351
L350:
	;
	v1251 = base.B2i32(v1244 == int32(0))
	goto L351
L351:
	;
	v1253 = v1251
	goto L336
L352:
	;
	goto L333
L353:
	;
	v1258 = v480
	goto L355
L354:
	;
	v1258 = v49
	goto L355
L355:
	;
	v1260 = v1258 + int32(1)
	if base.Ui32(v1180) <= base.Ui32(v1260) {
		v1366 = v1257
		goto L331
	} else {
		goto L356
	}
L356:
	;
	v1263 = base.B2i32(v1255 != int32(94))
	v1273 = v1258
	v1279 = v1260
	goto L357
L357:
	;
	v1283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1273)+2)))
	v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1273)+1)))
	if v1284 != int32(37) {
		goto L361
	} else {
		goto L362
	}
L358:
	;
	v1366 = v1257
	goto L331
L359:
	;
	v1348 = v1346 + int32(1)
	if base.Ui32(v1348) < base.Ui32(v1180) {
		v1273 = v1346
		v1279 = v1348
		goto L357
	} else {
		goto L386
	}
L360:
	;
	v1346 = v1273 + int32(2)
	goto L359
L361:
	;
	if v1283 != int32(45) {
		goto L380
	} else {
		goto L381
	}
L362:
	;
	v1287 = F_tolower(m, v1283)
	mBase = m.M
	switch v1287 + int32(-97) {
	case 0:
		goto L365
	default:
		goto L366
	case 2:
		goto L375
	case 3:
		goto L374
	case 11:
		goto L373
	case 15:
		goto L372
	case 18:
		goto L371
	case 20:
		goto L370
	case 22:
		goto L369
	case 23:
		goto L368
	case 25:
		goto L367
	}
L363:
	;
	if v1331 == int32(0) {
		goto L360
	} else {
		goto L379
	}
L364:
	;
	if base.Ui32(v1283+int32(-97)) < base.Ui32(int32(26)) {
		goto L376
	} else {
		goto L377
	}
L365:
	;
	v1322 = base.B2i32(base.Ui32(v1207|int32(32)+int32(-97)) < base.Ui32(int32(26)))
	goto L364
L366:
	;
	v1331 = base.B2i32(v1283 == v1207)
	goto L363
L367:
	;
	v1322 = base.B2i32(v1207 == int32(0))
	goto L364
L368:
	;
	v1312 = F_isxdigit(m, v1207)
	mBase = m.M
	v1322 = v1312
	goto L364
L369:
	;
	v1311 = F_isalnum(m, v1207)
	mBase = m.M
	v1322 = v1311
	goto L364
L370:
	;
	v1322 = base.B2i32(base.Ui32(v1207+int32(-65)) < base.Ui32(int32(26)))
	goto L364
L371:
	;
	v1322 = base.B2i32(v1207 == int32(32)) | base.B2i32(base.Ui32(v1207+int32(-9)) < base.Ui32(int32(5)))
	goto L364
L372:
	;
	v1299 = F_ispunct(m, v1207)
	mBase = m.M
	v1322 = v1299
	goto L364
L373:
	;
	v1322 = base.B2i32(base.Ui32(v1207+int32(-97)) < base.Ui32(int32(26)))
	goto L364
L374:
	;
	v1322 = base.B2i32(base.Ui32(v1207+int32(-48)) < base.Ui32(int32(10)))
	goto L364
L375:
	;
	v1290 = F_iscntrl(m, v1207)
	mBase = m.M
	v1322 = v1290
	goto L364
L376:
	;
	v1329 = v1322
	goto L378
L377:
	;
	v1329 = base.B2i32(v1322 == int32(0))
	goto L378
L378:
	;
	v1331 = v1329
	goto L363
L379:
	;
	v1366 = v1263
	goto L331
L380:
	;
	if v1207 != v1284 {
		v1346 = v1279
		goto L359
	} else {
		goto L385
	}
L381:
	;
	v1337 = v1273 + int32(3)
	if base.Ui32(v1180) <= base.Ui32(v1337) {
		goto L380
	} else {
		goto L382
	}
L382:
	;
	if base.Ui32(v1207) < base.Ui32(v1284) {
		v1346 = v1337
		goto L359
	} else {
		goto L383
	}
L383:
	;
	v1340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1337))))
	if base.Ui32(v1340) < base.Ui32(v1207) {
		v1346 = v1337
		goto L359
	} else {
		goto L384
	}
L384:
	;
	v1366 = v1263
	goto L331
L385:
	;
	v1366 = v1263
	goto L331
L386:
	;
	goto L358
L387:
	;
	goto L330
L388:
	;
	goto L329
L389:
	;
	v1433 = int32(0)
	if v1422 < v1433 {
		v1727 = v1433
		goto L3
	} else {
		goto L391
	}
L391:
	;
	v1439 = F_match(m, l0, v25+v1422, v468+int32(1))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L24
	} else {
		goto L392
	}
L392:
	;
	if v1439 == int32(0) {
		v1422 = v1422 + int32(-1)
		goto L389
	} else {
		goto L393
	}
L393:
	;
	v1727 = v1439
	goto L3
L394:
	;
	v1450 = v49 + int32(4)
	goto L396
L395:
	;
	v1450 = v1444
	goto L396
L396:
	;
	if v1447 == int32(94) {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v1453 = int32(4)
	goto L399
L398:
	;
	v1453 = int32(3)
	goto L399
L399:
	;
	v1455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v1453))))
	v1464 = v1450
	v1465 = v1455
	goto L400
L400:
	;
	if v1465&int32(255) != 0 {
		v1485 = v1465
		goto L402
	} else {
		goto L403
	}
L401:
	;
	v1504 = v1496 + int32(1)
	v1516 = v1444
	goto L6
L402:
	;
	v1487 = v1464 + int32(1)
	if v1485&int32(255) == int32(37) {
		goto L406
	} else {
		goto L407
	}
L403:
	;
	v1477 = m.G3
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1482 = F_luaL_error(m, v1478, v1477+int32(_a2733), int32(0))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L24
	} else {
		goto L404
	}
L404:
	;
	v1484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1464))))
	v1485 = v1484
	goto L402
L405:
	;
	v1497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1496))))
	if v1497 != int32(93) {
		v1464 = v1496
		v1465 = v1497
		goto L400
	} else {
		goto L411
	}
L406:
	;
	v1494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1464)+1)))
	if v1494 != 0 {
		goto L408
	} else {
		goto L409
	}
L407:
	;
	v1496 = v1487
	goto L405
L408:
	;
	v1495 = v1464 + int32(2)
	goto L410
L409:
	;
	v1495 = v1487
	goto L410
L410:
	;
	v1496 = v1495
	goto L405
L411:
	;
	goto L401
L412:
	;
	v1529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1516))))
	v1531 = base.B2i32(v1529 == int32(94))
	if v1529 == int32(94) {
		goto L414
	} else {
		goto L415
	}
L413:
	;
	v1525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(-1)))))
	v1526 = v1525
	goto L412
L414:
	;
	v1532 = v1516
	goto L416
L415:
	;
	v1532 = v49 + int32(2)
	goto L416
L416:
	;
	v1534 = v1532 + int32(1)
	v1536 = v1504 + int32(-1)
	if base.Ui32(v1536) <= base.Ui32(v1534) {
		v1727 = v1521
		goto L3
	} else {
		goto L417
	}
L417:
	;
	v1539 = base.B2i32(v1529 != int32(94))
	v1548 = v1532
	v1552 = v1534
	goto L419
L418:
	;
	if v1626 != 0 {
		v1727 = v1521
		goto L3
	} else {
		goto L449
	}
L419:
	;
	v1559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1548)+2)))
	v1560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1548)+1)))
	if v1560 != int32(37) {
		goto L423
	} else {
		goto L424
	}
L420:
	;
	v1626 = v1531
	goto L418
L421:
	;
	v1624 = v1622 + int32(1)
	if base.Ui32(v1624) < base.Ui32(v1536) {
		v1548 = v1622
		v1552 = v1624
		goto L419
	} else {
		goto L448
	}
L422:
	;
	v1622 = v1548 + int32(2)
	goto L421
L423:
	;
	if v1559 != int32(45) {
		goto L442
	} else {
		goto L443
	}
L424:
	;
	v1563 = F_tolower(m, v1559)
	mBase = m.M
	switch v1563 + int32(-97) {
	case 0:
		goto L427
	default:
		goto L428
	case 2:
		goto L437
	case 3:
		goto L436
	case 11:
		goto L435
	case 15:
		goto L434
	case 18:
		goto L433
	case 20:
		goto L432
	case 22:
		goto L431
	case 23:
		goto L430
	case 25:
		goto L429
	}
L425:
	;
	if v1607 == int32(0) {
		goto L422
	} else {
		goto L441
	}
L426:
	;
	if base.Ui32(v1559+int32(-97)) < base.Ui32(int32(26)) {
		goto L438
	} else {
		goto L439
	}
L427:
	;
	v1598 = base.B2i32(base.Ui32(v1526|int32(32)+int32(-97)) < base.Ui32(int32(26)))
	goto L426
L428:
	;
	v1607 = base.B2i32(v1559 == v1526)
	goto L425
L429:
	;
	v1598 = base.B2i32(v1526 == int32(0))
	goto L426
L430:
	;
	v1588 = F_isxdigit(m, v1526)
	mBase = m.M
	v1598 = v1588
	goto L426
L431:
	;
	v1587 = F_isalnum(m, v1526)
	mBase = m.M
	v1598 = v1587
	goto L426
L432:
	;
	v1598 = base.B2i32(base.Ui32(v1526+int32(-65)) < base.Ui32(int32(26)))
	goto L426
L433:
	;
	v1598 = base.B2i32(v1526 == int32(32)) | base.B2i32(base.Ui32(v1526+int32(-9)) < base.Ui32(int32(5)))
	goto L426
L434:
	;
	v1575 = F_ispunct(m, v1526)
	mBase = m.M
	v1598 = v1575
	goto L426
L435:
	;
	v1598 = base.B2i32(base.Ui32(v1526+int32(-97)) < base.Ui32(int32(26)))
	goto L426
L436:
	;
	v1598 = base.B2i32(base.Ui32(v1526+int32(-48)) < base.Ui32(int32(10)))
	goto L426
L437:
	;
	v1566 = F_iscntrl(m, v1526)
	mBase = m.M
	v1598 = v1566
	goto L426
L438:
	;
	v1605 = v1598
	goto L440
L439:
	;
	v1605 = base.B2i32(v1598 == int32(0))
	goto L440
L440:
	;
	v1607 = v1605
	goto L425
L441:
	;
	v1626 = v1539
	goto L418
L442:
	;
	if v1526 != v1560 {
		v1622 = v1552
		goto L421
	} else {
		goto L447
	}
L443:
	;
	v1613 = v1548 + int32(3)
	if base.Ui32(v1536) <= base.Ui32(v1613) {
		goto L442
	} else {
		goto L444
	}
L444:
	;
	if base.Ui32(v1526) < base.Ui32(v1560) {
		v1622 = v1613
		goto L421
	} else {
		goto L445
	}
L445:
	;
	v1616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1613))))
	if base.Ui32(v1616) < base.Ui32(v1526) {
		v1622 = v1613
		goto L421
	} else {
		goto L446
	}
L446:
	;
	v1626 = v1539
	goto L418
L447:
	;
	v1626 = v1539
	goto L418
L448:
	;
	goto L420
L449:
	;
	v1628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v1638 = v1532
	v1642 = v1534
	goto L451
L450:
	;
	if v1715 != 0 {
		v49 = v1504
		goto L4
	} else {
		goto L480
	}
L451:
	;
	v1648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1638)+2)))
	v1649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1638)+1)))
	if v1649 != int32(37) {
		goto L454
	} else {
		goto L455
	}
L452:
	;
	v1715 = v1531
	goto L450
L453:
	;
	v1711 = v1709 + int32(1)
	if base.Ui32(v1711) < base.Ui32(v1536) {
		v1638 = v1709
		v1642 = v1711
		goto L451
	} else {
		goto L479
	}
L454:
	;
	if v1648 != int32(45) {
		goto L473
	} else {
		goto L474
	}
L455:
	;
	v1652 = F_tolower(m, v1648)
	mBase = m.M
	switch v1652 + int32(-97) {
	case 0:
		goto L458
	default:
		goto L459
	case 2:
		goto L468
	case 3:
		goto L467
	case 11:
		goto L466
	case 15:
		goto L465
	case 18:
		goto L464
	case 20:
		goto L463
	case 22:
		goto L462
	case 23:
		goto L461
	case 25:
		goto L460
	}
L456:
	;
	if v1696 != 0 {
		v1715 = v1539
		goto L450
	} else {
		goto L472
	}
L457:
	;
	if base.Ui32(v1648+int32(-97)) < base.Ui32(int32(26)) {
		goto L469
	} else {
		goto L470
	}
L458:
	;
	v1687 = base.B2i32(base.Ui32(v1628|int32(32)+int32(-97)) < base.Ui32(int32(26)))
	goto L457
L459:
	;
	v1696 = base.B2i32(v1648 == v1628)
	goto L456
L460:
	;
	v1687 = base.B2i32(v1628 == int32(0))
	goto L457
L461:
	;
	v1677 = F_isxdigit(m, v1628)
	mBase = m.M
	v1687 = v1677
	goto L457
L462:
	;
	v1676 = F_isalnum(m, v1628)
	mBase = m.M
	v1687 = v1676
	goto L457
L463:
	;
	v1687 = base.B2i32(base.Ui32(v1628+int32(-65)) < base.Ui32(int32(26)))
	goto L457
L464:
	;
	v1687 = base.B2i32(v1628 == int32(32)) | base.B2i32(base.Ui32(v1628+int32(-9)) < base.Ui32(int32(5)))
	goto L457
L465:
	;
	v1664 = F_ispunct(m, v1628)
	mBase = m.M
	v1687 = v1664
	goto L457
L466:
	;
	v1687 = base.B2i32(base.Ui32(v1628+int32(-97)) < base.Ui32(int32(26)))
	goto L457
L467:
	;
	v1687 = base.B2i32(base.Ui32(v1628+int32(-48)) < base.Ui32(int32(10)))
	goto L457
L468:
	;
	v1655 = F_iscntrl(m, v1628)
	mBase = m.M
	v1687 = v1655
	goto L457
L469:
	;
	v1694 = v1687
	goto L471
L470:
	;
	v1694 = base.B2i32(v1687 == int32(0))
	goto L471
L471:
	;
	v1696 = v1694
	goto L456
L472:
	;
	v1709 = v1638 + int32(2)
	goto L453
L473:
	;
	if v1628 == v1649 {
		v1715 = v1539
		goto L450
	} else {
		goto L478
	}
L474:
	;
	v1702 = v1638 + int32(3)
	if base.Ui32(v1536) <= base.Ui32(v1702) {
		goto L473
	} else {
		goto L475
	}
L475:
	;
	if base.Ui32(v1628) < base.Ui32(v1649) {
		v1709 = v1702
		goto L453
	} else {
		goto L476
	}
L476:
	;
	v1705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1702))))
	if base.Ui32(v1705) < base.Ui32(v1628) {
		v1709 = v1702
		goto L453
	} else {
		goto L477
	}
L477:
	;
	v1715 = v1539
	goto L450
L478:
	;
	v1709 = v1642
	goto L453
L479:
	;
	goto L452
L480:
	;
	goto L5
}
func F_maxlenCommandHandler(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v106 int64
	_ = v106
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v128 int64
	_ = v128
	var v133 int64
	_ = v133
	var v137 int32
	_ = v137
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v149 int64
	_ = v149
	var v173 int64
	_ = v173
	var v192 int32
	_ = v192
	var v201 int64
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	switch l1 + int32(-1) {
	case 0:
		goto L4
	case 1:
		goto L3
	default:
		goto L2
	}
L1:
	;
	F_scriptingEngineDebuggerFlushLogs(m)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L5
	} else {
		goto L61
	}
L2:
	;
	v247 = F_sdsnew(m, int32(_a2040))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L5
	} else {
		goto L58
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = F_objectGetVal(m, v31)
	mBase = m.M
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = F_objectGetVal(m, v34)
	mBase = m.M
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+int32(-1)))))
	switch v38 & int32(7) {
	case 0:
		goto L15
	case 1:
		goto L14
	case 2:
		goto L13
	case 3:
		goto L12
	case 4:
		goto L11
	default:
		v55 = int32(0)
		goto L10
	}
L4:
	;
	v14 = F_sdsempty(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[660]))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v19
	v23 = F_sdscatfmt(m, v14, int32(_a2041), v10)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v25 = F_createObject(m, v18, v23)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	v29 = F_listAddNodeTail(m, v28, v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L1
L10:
	;
	v57 = v10 + int32(24)
	v58 = int32(0)
	if base.Ui32(v55+int32(-21)) < base.Ui32(int32(-20)) {
		v192 = v58
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v35+int32(-17))))
	v55 = v54
	goto L10
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v35+int32(-9))))
	v55 = v51
	goto L10
L13:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(-5)))))
	v55 = v48
	goto L10
L14:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+int32(-3)))))
	v55 = v45
	goto L10
L15:
	;
	v55 = int32(base.Ui32(v38) >> (uint(int32(3)) % 32))
	goto L10
L16:
	;
	v239 = F_createObject(m, int32(0), v236)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L5
	} else {
		goto L56
	}
L17:
	;
	v234 = F_sdsnew(m, int32(_a2042))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L5
	} else {
		goto L55
	}
L18:
	;
	if v192 == int32(0) {
		goto L17
	} else {
		goto L45
	}
L19:
	;
	goto L18
L20:
	;
	v70 = int32(1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v55 != v70 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v192 = int32(1)
	goto L19
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57))) = v173
	goto L21
L23:
	;
	if v71&int32(255) == int32(45) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v75 = v71 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v75&int32(255)) {
		v192 = v58
		goto L19
	} else {
		goto L25
	}
L25:
	;
	if v57 == int32(0) {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v173 = base.I64_extend_i32_u(v75) & int64(255)
	goto L22
L27:
	;
	if base.Ui32(int32(8)) < base.Ui32((v94+int32(-49))&int32(255)) {
		v192 = v58
		goto L19
	} else {
		goto L30
	}
L28:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	v93 = int32(2)
	v94 = v91
	v95 = v32 + int32(1)
	goto L27
L29:
	;
	v93 = v70
	v94 = v71
	v95 = v32
	goto L27
L30:
	;
	v106 = base.I64_extend_i32_u(v94+int32(-48)) & int64(255)
	if base.Ui32(v55) <= base.Ui32(v93) {
		v149 = v106
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v71&int32(255) != int32(45) {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	v112 = v93
	v114 = v106
	v116 = v95
	goto L33
L33:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if base.Ui32((v118+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v192 = v58
		goto L19
	} else {
		goto L35
	}
L34:
	;
	v149 = v139
	goto L31
L35:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v114) {
		v192 = v58
		goto L19
	} else {
		goto L36
	}
L36:
	;
	v128 = v114 * int64(10)
	v133 = base.I64_extend_i32_u(v118+int32(-48)) & int64(255)
	if base.Ui64(v133^int64(-1)) < base.Ui64(v128) {
		v192 = v58
		goto L19
	} else {
		goto L37
	}
L37:
	;
	v137 = int32(1)
	v139 = v128 + v133
	v141 = v112 + v137
	if v141 != v55 {
		v112 = v141
		v114 = v139
		v116 = v116 + v137
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	if v149 < int64(0) {
		v192 = v58
		goto L19
	} else {
		goto L43
	}
L40:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v149) {
		v192 = v58
		goto L19
	} else {
		goto L41
	}
L41:
	;
	if v57 == int32(0) {
		goto L21
	} else {
		goto L42
	}
L42:
	;
	v173 = int64(0) - v149
	goto L22
L43:
	;
	if v57 == int32(0) {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	v173 = v149
	goto L22
L45:
	;
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
	if v201 < int64(0) {
		goto L17
	} else {
		goto L46
	}
L46:
	;
	v204 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[661])) = int32(1)
	v208 = int32(60)
	v209 = base.I32_wrap_i64(v201)
	if base.Ui32(v209+int32(-1)) < base.Ui32(v208) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v214 = v208
	goto L49
L48:
	;
	v214 = v209
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, _consts[660])) = v214
	v216 = F_sdsempty(m)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	if v201 != int64(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _consts[660]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v225
	v230 = F_sdscatfmt(m, v216, int32(_a2043), v10+int32(16))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L54
	}
L52:
	;
	v222 = F_sdscatfmt(m, v216, int32(_a2044), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	v236 = v222
	goto L16
L54:
	;
	v236 = v230
	goto L16
L55:
	;
	v236 = v234
	goto L16
L56:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	v243 = F_listAddNodeTail(m, v242, v239)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	goto L1
L58:
	;
	v249 = F_createObject(m, int32(0), v247)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	v253 = F_listAddNodeTail(m, v252, v249)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	goto L1
L61:
	;
	m.G0 = v10 + int32(32)
	return int32(1)
}
func F_memmapchars(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v57 int32
	_ = v57
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	v18 = int32(0)
	goto L3
L3:
	;
	if l4 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v57 = v18 + int32(1)
	if v57 != l1 {
		v18 = v57
		goto L3
	} else {
		goto L12
	}
L6:
	;
	v25 = l0 + v18
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v35 = int32(0)
	goto L7
L7:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v35))))
	if v26&int32(255) != v39 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	v45 = v35 + int32(1)
	if v45 != l4 {
		v35 = v45
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v35))))
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v42)
	goto L5
L11:
	;
	goto L8
L12:
	;
	goto L4
}
func F_mempbrk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v67 int32
	_ = v67
	if l1 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v67
L2:
	;
	v67 = int32(0)
	goto L1
L3:
	;
	v16 = int32(0)
	goto L4
L4:
	;
	if l3 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L2
L6:
	;
	v50 = v16 + int32(1)
	if v50 != l1 {
		v16 = v50
		goto L4
	} else {
		goto L12
	}
L7:
	;
	v23 = l0 + v16
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v32 = int32(0)
	goto L8
L8:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v32))))
	if v24&int32(255) == v36 {
		v67 = v23
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v39 = v32 + int32(1)
	if v39 != l3 {
		v32 = v39
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	goto L5
}
func F_mgetCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_addReplyArrayLen(m, l0, v4+int32(-1))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 < int32(2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v14 = int32(1)
	goto L5
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v14<<(uint(int32(2))%32))))
	v22 = F_lookupKeyRead(m, v16, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	goto L3
L7:
	;
	v36 = v14 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v36 < v37 {
		v14 = v36
		goto L5
	} else {
		goto L16
	}
L8:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v26&int32(15) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	if v22 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	F_addReplyBulk(m, l0, v22)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L7
L15:
	;
	goto L7
L16:
	;
	goto L6
}
func F_migrateCloseTimedoutSockets(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
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
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v190 int64
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	v6 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	v7 = F_dictGetSafeIterator(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v7)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L2
	} else {
		goto L67
	}
L2:
	;
	return
L3:
	;
	v16 = v7 + int32(20)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v112 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v23 = v16
	v24 = v20
	goto L8
L6:
	;
	v20 = int32(1)
	goto L5
L7:
	;
	v20 = int32(0)
	goto L5
L8:
	;
	switch v24 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v24 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v104
	if v104 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v28 != int32(-1) {
		v67 = v28
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v68 = int32(1)
	v69 = v67 + v68
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v69
	v71 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v75+int32(26)))))
	if v79 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v32 != 0 {
		v67 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v34 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	if v61 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v41 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v33)+16)))
	v42 = int64(*(*int8)(unsafe.Add(mBase, uint32(v33)+27)))
	v43 = int64(*(*int32)(unsafe.Add(mBase, uint32(v33)+8)))
	v44 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v33)+12)))
	v45 = int64(*(*int8)(unsafe.Add(mBase, uint32(v33)+26)))
	v46 = int64(*(*int32)(unsafe.Add(mBase, uint32(v33)+4)))
	v47 = F_wangHash64(m, v46)
	mBase = m.M
	v49 = F_wangHash64(m, v45+v47)
	mBase = m.M
	v51 = F_wangHash64(m, v44+v49)
	mBase = m.M
	v53 = F_wangHash64(m, v43+v51)
	mBase = m.M
	v55 = F_wangHash64(m, v42+v53)
	mBase = m.M
	v57 = F_wangHash64(m, v41+v55)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v60 = v59
	goto L17
L19:
	;
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+24)))
	v39 = v37 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+24)) = uint16(v39)
	v60 = v33
	goto L17
L20:
	;
	v67 = v61 + int32(-1)
	goto L14
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v67 = v64
	goto L14
L22:
	;
	v94 = int32(2)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v74+v92<<(uint(v94)%32)+int32(4))))
	v23 = v99 + v93<<(uint(v94)%32)
	v24 = int32(1)
	goto L8
L23:
	;
	v83 = v71
	goto L25
L24:
	;
	v83 = v68 << (uint(v79) % 32)
	goto L25
L25:
	;
	if v69 < v83 {
		v92 = v75
		v93 = v69
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v75 != 0 {
		v112 = v71
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	if v85 == int32(-1) {
		v112 = v71
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967296)
	v92 = int32(1)
	v93 = int32(0)
	goto L22
L29:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v108
	v112 = v104
	goto L11
L30:
	;
	v119 = v112
	goto L31
L31:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
	goto L33
L32:
	;
	goto L1
L33:
	;
	v124 = *(*int64)(unsafe.Add(mBase, _consts[109]))
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v122)+8))
	if v124-v125 < int64(11) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v149 = v7 + int32(20)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v150 != 0 {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+52))
	m.T0[v131].(func(*base.Module, int32))(m, v129)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	F_valkey_free(m, v122)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	goto L38
L38:
	;
	v139 = F_dictDelete(m, v137, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	goto L34
L40:
	;
	if v245 != 0 {
		v119 = v245
		goto L31
	} else {
		goto L66
	}
L41:
	;
	v156 = v149
	v157 = v153
	goto L44
L42:
	;
	v153 = int32(1)
	goto L41
L43:
	;
	v153 = int32(0)
	goto L41
L44:
	;
	switch v157 {
	case 0:
		goto L49
	default:
		goto L48
	}
L46:
	;
	v157 = int32(0)
	goto L44
L47:
	;
	goto L40
L48:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v237
	if v237 == int32(0) {
		goto L46
	} else {
		goto L65
	}
L49:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v161 != int32(-1) {
		v200 = v161
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v201 = int32(1)
	v202 = v200 + v201
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v202
	v204 = int32(0)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v208+int32(26)))))
	if v212 == int32(255) {
		goto L59
	} else {
		goto L60
	}
L51:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v165 != 0 {
		v200 = int32(-1)
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v167 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+20))
	if v194 != int32(-1) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v174 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v166)+16)))
	v175 = int64(*(*int8)(unsafe.Add(mBase, uint32(v166)+27)))
	v176 = int64(*(*int32)(unsafe.Add(mBase, uint32(v166)+8)))
	v177 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v166)+12)))
	v178 = int64(*(*int8)(unsafe.Add(mBase, uint32(v166)+26)))
	v179 = int64(*(*int32)(unsafe.Add(mBase, uint32(v166)+4)))
	v180 = F_wangHash64(m, v179)
	mBase = m.M
	v182 = F_wangHash64(m, v178+v180)
	mBase = m.M
	v184 = F_wangHash64(m, v177+v182)
	mBase = m.M
	v186 = F_wangHash64(m, v176+v184)
	mBase = m.M
	v188 = F_wangHash64(m, v175+v186)
	mBase = m.M
	v190 = F_wangHash64(m, v174+v188)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v193 = v192
	goto L53
L55:
	;
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+24)))
	v172 = v170 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v166)+24)) = uint16(v172)
	v193 = v166
	goto L53
L56:
	;
	v200 = v194 + int32(-1)
	goto L50
L57:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v200 = v197
	goto L50
L58:
	;
	v227 = int32(2)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v207+v225<<(uint(v227)%32)+int32(4))))
	v156 = v232 + v226<<(uint(v227)%32)
	v157 = int32(1)
	goto L44
L59:
	;
	v216 = v204
	goto L61
L60:
	;
	v216 = v201 << (uint(v212) % 32)
	goto L61
L61:
	;
	if v202 < v216 {
		v225 = v208
		v226 = v202
		goto L58
	} else {
		goto L62
	}
L62:
	;
	if v208 != 0 {
		v245 = v204
		goto L47
	} else {
		goto L63
	}
L63:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v207)+20))
	if v218 == int32(-1) {
		v245 = v204
		goto L47
	} else {
		goto L64
	}
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967296)
	v225 = int32(1)
	v226 = int32(0)
	goto L58
L65:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v241
	v245 = v237
	goto L47
L66:
	;
	goto L32
L67:
	;
	return
}
func F_migrateGetSocket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
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
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
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
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int64
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	v11 = F_sdsempty(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_objectGetVal(m, l1)
		mBase = m.M
		v16 = int32(0)
		v18 = F_objectGetVal(m, l1)
		mBase = m.M
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(-1)))))
		switch v21 & int32(7) {
		case 0:
			v38 = int32(base.Ui32(v21) >> (uint(int32(3)) % 32))
		case 1:
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(-3)))))
			v38 = v28
		case 2:
			v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18+int32(-5)))))
			v38 = v31
		case 3:
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(-9))))
			v38 = v34
		case 4:
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(-17))))
			v38 = v37
		default:
			v38 = v16
		}
		v39 = F_sdscatlen(m, v11, v15, v38)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			v43 = F_sdscatlen(m, v39, int32(_a206), int32(1))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v45 = F_objectGetVal(m, l2)
				mBase = m.M
				v46 = F_objectGetVal(m, l2)
				mBase = m.M
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+int32(-1)))))
				switch v49 & int32(7) {
				case 0:
					v66 = int32(base.Ui32(v49) >> (uint(int32(3)) % 32))
				case 1:
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+int32(-3)))))
					v66 = v56
				case 2:
					v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46+int32(-5)))))
					v66 = v59
				case 3:
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(-9))))
					v66 = v62
				case 4:
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(-17))))
					v66 = v65
				default:
					v66 = v16
				}
				v67 = F_sdscatlen(m, v43, v45, v66)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, _consts[108]))
					v71 = F_dictFetchValue(m, v70, v67)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						if v71 == int32(0) {
							v82 = *(*int32)(unsafe.Add(mBase, _consts[108]))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
							if v83+v84 != int32(64) {
								v107 = *(*int32)(unsafe.Add(mBase, _consts[107]))
								if v107 == int32(0) {
									v112 = F_connectionTypeTcp(m)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										v114 = v112
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+40))
										v116 = m.T0[v115].(func(*base.Module) int32)(m)
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int32(0)
										} else {
											v118 = F_objectGetVal(m, l1)
											mBase = m.M
											v119 = F_objectGetVal(m, l2)
											mBase = m.M
											v123 = v119
											for {
												v128 = v123 + int32(1)
												v129 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123))))
												v130 = F___isspace_1(m, v129)
												mBase = m.M
												if v130 != 0 {
													v123 = v128
													continue
												} else {
													break
												}
												break
											}
											v131 = int32(1)
											switch v129&int32(255) + int32(-43) {
											case 0:
												v137 = v131
												v138 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
												v139 = v128
												v140 = v138
												v141 = v137
											default:
												v139 = v123
												v140 = v129
												v141 = v131
											case 2:
												v137 = int32(0)
												v138 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
												v139 = v128
												v140 = v138
												v141 = v137
											}
											v144 = v140 + int32(-48)
											if base.Ui32(int32(9)) < base.Ui32(v144) {
												v162 = int32(0)
											} else {
												v148 = int32(0)
												v149 = v139
												v150 = v144
												for {
													v152 = int32(10)
													v154 = v148*v152 - v150
													v155 = int32(*(*int8)(unsafe.Add(mBase, uint32(v149)+1)))
													v159 = v155 + int32(-48)
													if base.Ui32(v159) < base.Ui32(v152) {
														v148 = v154
														v149 = v149 + int32(1)
														v150 = v159
														continue
													} else {
														break
													}
													break
												}
												v162 = v154
											}
											if v141 != 0 {
												v168 = int32(0) - v162
											} else {
												v168 = v162
											}
											v170 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
											v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+60))
											v172 = m.T0[v171].(func(*base.Module, int32, int32, int32, int64) int32)(m, v116, v118, v168, base.I64_extend_i32_s(l3))
											mBase = m.M
											v173 = m.ExcPending
											if v173 != 0 {
												return int32(0)
											} else {
												if v172 == int32(0) {
													v188 = F_valkey_malloc(m, int32(16))
													mBase = m.M
													v189 = m.ExcPending
													if v189 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = int32(-1)
														*(*int32)(unsafe.Add(mBase, uint32(v188))) = v116
														v193 = int32(_a20)
														v194 = *(*int64)(unsafe.Add(mBase, _consts[109]))
														*(*int64)(unsafe.Add(mBase, uint32(v188)+8)) = v194
														v197 = *(*int32)(unsafe.Add(mBase, _consts[108]))
														v198 = F_dictAdd(m, v197, v67, v188)
														mBase = m.M
														v199 = m.ExcPending
														if v199 != 0 {
															return int32(0)
														} else {
															return v188
														}
													}
												} else {
													F_addReplyError(m, l0, int32(_a207))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return int32(0)
													} else {
														v179 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
														v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+52))
														m.T0[v180].(func(*base.Module, int32))(m, v116)
														mBase = m.M
														v182 = m.ExcPending
														if v182 != 0 {
															return int32(0)
														} else {
															F_sdsfree(m, v67)
															mBase = m.M
															v184 = m.ExcPending
															if v184 != 0 {
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
								} else {
									v110 = F_connectionTypeTls(m)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										v114 = v110
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+40))
										v116 = m.T0[v115].(func(*base.Module) int32)(m)
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int32(0)
										} else {
											v118 = F_objectGetVal(m, l1)
											mBase = m.M
											v119 = F_objectGetVal(m, l2)
											mBase = m.M
											v123 = v119
											for {
												v128 = v123 + int32(1)
												v129 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123))))
												v130 = F___isspace_1(m, v129)
												mBase = m.M
												if v130 != 0 {
													v123 = v128
													continue
												} else {
													break
												}
												break
											}
											v131 = int32(1)
											switch v129&int32(255) + int32(-43) {
											case 0:
												v137 = v131
												v138 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
												v139 = v128
												v140 = v138
												v141 = v137
											default:
												v139 = v123
												v140 = v129
												v141 = v131
											case 2:
												v137 = int32(0)
												v138 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
												v139 = v128
												v140 = v138
												v141 = v137
											}
											v144 = v140 + int32(-48)
											if base.Ui32(int32(9)) < base.Ui32(v144) {
												v162 = int32(0)
											} else {
												v148 = int32(0)
												v149 = v139
												v150 = v144
												for {
													v152 = int32(10)
													v154 = v148*v152 - v150
													v155 = int32(*(*int8)(unsafe.Add(mBase, uint32(v149)+1)))
													v159 = v155 + int32(-48)
													if base.Ui32(v159) < base.Ui32(v152) {
														v148 = v154
														v149 = v149 + int32(1)
														v150 = v159
														continue
													} else {
														break
													}
													break
												}
												v162 = v154
											}
											if v141 != 0 {
												v168 = int32(0) - v162
											} else {
												v168 = v162
											}
											v170 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
											v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+60))
											v172 = m.T0[v171].(func(*base.Module, int32, int32, int32, int64) int32)(m, v116, v118, v168, base.I64_extend_i32_s(l3))
											mBase = m.M
											v173 = m.ExcPending
											if v173 != 0 {
												return int32(0)
											} else {
												if v172 == int32(0) {
													v188 = F_valkey_malloc(m, int32(16))
													mBase = m.M
													v189 = m.ExcPending
													if v189 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = int32(-1)
														*(*int32)(unsafe.Add(mBase, uint32(v188))) = v116
														v193 = int32(_a20)
														v194 = *(*int64)(unsafe.Add(mBase, _consts[109]))
														*(*int64)(unsafe.Add(mBase, uint32(v188)+8)) = v194
														v197 = *(*int32)(unsafe.Add(mBase, _consts[108]))
														v198 = F_dictAdd(m, v197, v67, v188)
														mBase = m.M
														v199 = m.ExcPending
														if v199 != 0 {
															return int32(0)
														} else {
															return v188
														}
													}
												} else {
													F_addReplyError(m, l0, int32(_a207))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return int32(0)
													} else {
														v179 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
														v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+52))
														m.T0[v180].(func(*base.Module, int32))(m, v116)
														mBase = m.M
														v182 = m.ExcPending
														if v182 != 0 {
															return int32(0)
														} else {
															F_sdsfree(m, v67)
															mBase = m.M
															v184 = m.ExcPending
															if v184 != 0 {
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
							} else {
								v88 = F_dictGetRandomKey(m, v82)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+52))
									m.T0[v93].(func(*base.Module, int32))(m, v91)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										F_valkey_free(m, v90)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, _consts[108]))
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
											v101 = F_dictDelete(m, v99, v100)
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return int32(0)
											} else {
												v107 = *(*int32)(unsafe.Add(mBase, _consts[107]))
												if v107 == int32(0) {
													v112 = F_connectionTypeTcp(m)
													mBase = m.M
													v113 = m.ExcPending
													if v113 != 0 {
														return int32(0)
													} else {
														v114 = v112
														v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+40))
														v116 = m.T0[v115].(func(*base.Module) int32)(m)
														mBase = m.M
														v117 = m.ExcPending
														if v117 != 0 {
															return int32(0)
														} else {
															v118 = F_objectGetVal(m, l1)
															mBase = m.M
															v119 = F_objectGetVal(m, l2)
															mBase = m.M
															v123 = v119
															for {
																v128 = v123 + int32(1)
																v129 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123))))
																v130 = F___isspace_1(m, v129)
																mBase = m.M
																if v130 != 0 {
																	v123 = v128
																	continue
																} else {
																	break
																}
																break
															}
															v131 = int32(1)
															switch v129&int32(255) + int32(-43) {
															case 0:
																v137 = v131
																v138 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
																v139 = v128
																v140 = v138
																v141 = v137
															default:
																v139 = v123
																v140 = v129
																v141 = v131
															case 2:
																v137 = int32(0)
																v138 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
																v139 = v128
																v140 = v138
																v141 = v137
															}
															v144 = v140 + int32(-48)
															if base.Ui32(int32(9)) < base.Ui32(v144) {
																v162 = int32(0)
															} else {
																v148 = int32(0)
																v149 = v139
																v150 = v144
																for {
																	v152 = int32(10)
																	v154 = v148*v152 - v150
																	v155 = int32(*(*int8)(unsafe.Add(mBase, uint32(v149)+1)))
																	v159 = v155 + int32(-48)
																	if base.Ui32(v159) < base.Ui32(v152) {
																		v148 = v154
																		v149 = v149 + int32(1)
																		v150 = v159
																		continue
																	} else {
																		break
																	}
																	break
																}
																v162 = v154
															}
															if v141 != 0 {
																v168 = int32(0) - v162
															} else {
																v168 = v162
															}
															v170 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
															v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+60))
															v172 = m.T0[v171].(func(*base.Module, int32, int32, int32, int64) int32)(m, v116, v118, v168, base.I64_extend_i32_s(l3))
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return int32(0)
															} else {
																if v172 == int32(0) {
																	v188 = F_valkey_malloc(m, int32(16))
																	mBase = m.M
																	v189 = m.ExcPending
																	if v189 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = int32(-1)
																		*(*int32)(unsafe.Add(mBase, uint32(v188))) = v116
																		v193 = int32(_a20)
																		v194 = *(*int64)(unsafe.Add(mBase, _consts[109]))
																		*(*int64)(unsafe.Add(mBase, uint32(v188)+8)) = v194
																		v197 = *(*int32)(unsafe.Add(mBase, _consts[108]))
																		v198 = F_dictAdd(m, v197, v67, v188)
																		mBase = m.M
																		v199 = m.ExcPending
																		if v199 != 0 {
																			return int32(0)
																		} else {
																			return v188
																		}
																	}
																} else {
																	F_addReplyError(m, l0, int32(_a207))
																	mBase = m.M
																	v178 = m.ExcPending
																	if v178 != 0 {
																		return int32(0)
																	} else {
																		v179 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
																		v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+52))
																		m.T0[v180].(func(*base.Module, int32))(m, v116)
																		mBase = m.M
																		v182 = m.ExcPending
																		if v182 != 0 {
																			return int32(0)
																		} else {
																			F_sdsfree(m, v67)
																			mBase = m.M
																			v184 = m.ExcPending
																			if v184 != 0 {
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
												} else {
													v110 = F_connectionTypeTls(m)
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int32(0)
													} else {
														v114 = v110
														v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+40))
														v116 = m.T0[v115].(func(*base.Module) int32)(m)
														mBase = m.M
														v117 = m.ExcPending
														if v117 != 0 {
															return int32(0)
														} else {
															v118 = F_objectGetVal(m, l1)
															mBase = m.M
															v119 = F_objectGetVal(m, l2)
															mBase = m.M
															v123 = v119
															for {
																v128 = v123 + int32(1)
																v129 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123))))
																v130 = F___isspace_1(m, v129)
																mBase = m.M
																if v130 != 0 {
																	v123 = v128
																	continue
																} else {
																	break
																}
																break
															}
															v131 = int32(1)
															switch v129&int32(255) + int32(-43) {
															case 0:
																v137 = v131
																v138 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
																v139 = v128
																v140 = v138
																v141 = v137
															default:
																v139 = v123
																v140 = v129
																v141 = v131
															case 2:
																v137 = int32(0)
																v138 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
																v139 = v128
																v140 = v138
																v141 = v137
															}
															v144 = v140 + int32(-48)
															if base.Ui32(int32(9)) < base.Ui32(v144) {
																v162 = int32(0)
															} else {
																v148 = int32(0)
																v149 = v139
																v150 = v144
																for {
																	v152 = int32(10)
																	v154 = v148*v152 - v150
																	v155 = int32(*(*int8)(unsafe.Add(mBase, uint32(v149)+1)))
																	v159 = v155 + int32(-48)
																	if base.Ui32(v159) < base.Ui32(v152) {
																		v148 = v154
																		v149 = v149 + int32(1)
																		v150 = v159
																		continue
																	} else {
																		break
																	}
																	break
																}
																v162 = v154
															}
															if v141 != 0 {
																v168 = int32(0) - v162
															} else {
																v168 = v162
															}
															v170 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
															v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+60))
															v172 = m.T0[v171].(func(*base.Module, int32, int32, int32, int64) int32)(m, v116, v118, v168, base.I64_extend_i32_s(l3))
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return int32(0)
															} else {
																if v172 == int32(0) {
																	v188 = F_valkey_malloc(m, int32(16))
																	mBase = m.M
																	v189 = m.ExcPending
																	if v189 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = int32(-1)
																		*(*int32)(unsafe.Add(mBase, uint32(v188))) = v116
																		v193 = int32(_a20)
																		v194 = *(*int64)(unsafe.Add(mBase, _consts[109]))
																		*(*int64)(unsafe.Add(mBase, uint32(v188)+8)) = v194
																		v197 = *(*int32)(unsafe.Add(mBase, _consts[108]))
																		v198 = F_dictAdd(m, v197, v67, v188)
																		mBase = m.M
																		v199 = m.ExcPending
																		if v199 != 0 {
																			return int32(0)
																		} else {
																			return v188
																		}
																	}
																} else {
																	F_addReplyError(m, l0, int32(_a207))
																	mBase = m.M
																	v178 = m.ExcPending
																	if v178 != 0 {
																		return int32(0)
																	} else {
																		v179 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
																		v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+52))
																		m.T0[v180].(func(*base.Module, int32))(m, v116)
																		mBase = m.M
																		v182 = m.ExcPending
																		if v182 != 0 {
																			return int32(0)
																		} else {
																			F_sdsfree(m, v67)
																			mBase = m.M
																			v184 = m.ExcPending
																			if v184 != 0 {
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
							}
						} else {
							F_sdsfree(m, v67)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								v78 = *(*int64)(unsafe.Add(mBase, _consts[109]))
								*(*int64)(unsafe.Add(mBase, uint32(v71)+8)) = v78
								return v71
							}
						}
					}
				}
			}
		}
	}
}
func F_monitorActiveDefrag(m *base.Module) {
	return
}
func F_monitorCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
	if v2&int32(16) == int32(0) {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
		if v10&int32(2) != 0 {
			return
		} else {
			F_initClientReplicationData(m, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v15 | int32(6)
				v20 = *(*int32)(unsafe.Add(mBase, _consts[336]))
				v21 = F_listAddNodeTail(m, v20, l0)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, _consts[27]))
					F_addReply(m, l0, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		F_addReplyError(m, l0, int32(_a2282))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	}
}
func F_monotonicGetType(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	return base.B2i32(v2 != int32(948))
}
func F_monotonicInit(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v23 int64
	_ = v23
	var v36 int32
	_ = v36
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	if v7 != 0 {
		m.G0 = v4 + int32(16)
		return int32(_a1623)
	} else {
		v9 = F___clock_gettime(m, int32(1), v4)
		mBase = m.M
		if v9 != 0 {
			F__serverAssert(m, int32(_a1624), int32(_a1625), int32(172))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v10 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[34])) = int32(948)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[460]))
			*(*int32)(unsafe.Add(mBase, _consts[461])) = v15
			v19 = *(*int64)(unsafe.Add(mBase, _consts[462]))
			*(*int64)(unsafe.Add(mBase, _consts[463])) = v19
			v23 = *(*int64)(unsafe.Add(mBase, _consts[464]))
			*(*int64)(unsafe.Add(mBase, _consts[465])) = v23
			m.G0 = v4 + int32(16)
			return int32(_a1623)
		}
	}
}
func F_move_next(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
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
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int64
	_ = v65
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v83 int64
	_ = v83
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = v12 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	if v17 <= v14 {
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
		if v19 == int32(0) {
			v32 = v14
		} else {
			v22 = int32(0)
			v25 = v14 - v19
			if v25 < v17 {
				v27 = v22
			} else {
				v27 = v22 - v17
			}
			if v25 < int32(0) {
				v30 = v17
			} else {
				v30 = v27
			}
			v32 = v30 + v25
		}
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
		v37 = *(*int64)(unsafe.Add(mBase, uint32(v33+v32<<(uint(int32(3))%32))))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v37
		v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v37 + v39
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
		v43 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
		v48 = int32(0)
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
		v50 = v14 >> (uint(v49) % 32)
		if v48 < v50 {
			v53 = v44
		} else {
			v53 = v48
		}
		v56 = int32(1)
		if v56 < v50 {
			v59 = v50
		} else {
			v59 = v56
		}
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
		v65 = base.I64_extend_i32_s((v44+int32(-1))&v14+v53) << (uint(base.I64_extend_i32_u(v59+v60+int32(-1))) % 64)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v65
		v72 = int32(63) - (v49 + base.I32_wrap_i64(base.I64_clz(v65|v43)))
		v73 = base.I64_extend_i32_u(v72)
		v74 = v65 >> (uint(v73) % 64)
		v76 = base.I64_extend32_s(v74) << (uint(v73) % 64)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v76
		v78 = int64(1)
		v83 = v78 << (uint(base.I64_extend_i32_u(v72+base.B2i32(v42 <= base.I32_wrap_i64(v74)))) % 64)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v83>>(uint(v78)%64) + v76
		*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v76 + v83 + int64(-1)
	}
	return base.B2i32(v14 < v17)
}
func F_msetGenericCommand(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int64
	_ = v142
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v13&int32(1) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	if l1 != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	F_addReplyErrorArity(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	goto L1
L6:
	;
	v140 = int32(_a20)
	v142 = *(*int64)(unsafe.Add(mBase, _consts[180]))
	v146 = base.I32_div_s(v135+int32(-1), int32(2))
	*(*int64)(unsafe.Add(mBase, _consts[180])) = v142 + base.I64_extend_i32_s(v146)
	if l1 != 0 {
		goto L35
	} else {
		goto L36
	}
L7:
	;
	if v50 < int32(2) {
		v135 = v50
		goto L6
	} else {
		goto L18
	}
L8:
	;
	if v13 < int32(2) {
		v135 = v13
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v50 = v13
	v51 = int32(0)
	goto L7
L10:
	;
	v27 = int32(1)
	goto L12
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	F_addReply(m, l0, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L17
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v27<<(uint(int32(2))%32))))
	v36 = F_lookupKeyWrite(m, v30, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L14
	}
L13:
	;
	v50 = v40
	v51 = int32(8)
	goto L7
L14:
	;
	if v36 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v39 = v27 + int32(2)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v39 < v40 {
		v27 = v39
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	goto L1
L18:
	;
	v62 = v51
	v63 = int32(1)
	goto L19
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v67 = int32(1)
	v68 = v63 + v67
	v70 = v68 << (uint(int32(2)) % 32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66+v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v72
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v74&v67 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v135 = v130
	goto L6
L21:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115+v63<<(uint(int32(2))%32))))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+28))
	F_notifyKeyspaceEvent(m, int32(8), int32(_a2432), v121, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L30
	}
L22:
	;
	v95 = F_tryObjectEncoding(m, v72)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L27
	}
L23:
	;
	F_incrRefCount(m, v72)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v63<<(uint(int32(2))%32))))
	F_setKey(m, l0, v81, v86, v11+int32(12), v62)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_rewriteClientCommandArgument(m, l0, v68, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v115 = v94
	goto L21
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v95
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v99+v63<<(uint(int32(2))%32))))
	F_setKey(m, l0, v98, v103, v11+int32(12), v62)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_incrRefCount(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v111+v70))) = v113
	v115 = v111
	goto L21
L30:
	;
	if l1 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v127 = int32(16)
	goto L33
L32:
	;
	v127 = v62
	goto L33
L33:
	;
	v129 = v63 + int32(2)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v129 < v130 {
		v62 = v127
		v63 = v129
		goto L19
	} else {
		goto L34
	}
L34:
	;
	goto L20
L35:
	;
	v154 = int32(_a2434)
	goto L37
L36:
	;
	v154 = int32(_a388)
	goto L37
L37:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	F_addReply(m, l0, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	goto L1
}
func F_multiStateMemOverhead(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v13 = v9*int32(40) + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
	if v14 == int32(0) {
		v52 = v6
		v53 = v13
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return int32(0)
L3:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	return v56*int32(20) + v53
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	v21 = v18<<(uint(int32(2))%32) + v13
	if v18 < int32(1) {
		v52 = v6
		v53 = v21
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v26 = int32(0)
	v27 = v21
	v28 = v18
	goto L6
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v26<<(uint(int32(2))%32))))
	if v35 == int32(0) {
		v45 = v27
		v46 = v28
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v52 = v50
	v53 = v45
	goto L3
L8:
	;
	v48 = v26 + int32(1)
	if v48 < v46 {
		v26 = v48
		v27 = v45
		v28 = v46
		goto L6
	} else {
		goto L12
	}
L9:
	;
	v38 = F_hashtableMemUsage(m, v35)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	v45 = v38 + v27
	v46 = v44
	goto L8
L12:
	;
	goto L7
}
func F_mustObeyClient(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v3 != int64(-1) {
		v8 = int32(1)
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
		if v9&v8 != 0 {
			v19 = v8
			return v19
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
			if v12 != 0 {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v19 = base.B2i32(v15 == int32(1))
				return v19
			} else {
				return int32(0)
			}
		}
	} else {
		return int32(1)
	}
}
func F_myselfIsBestRankedReplica(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[147])))
	if v7 != int64(0) {
		v60 = int32(0)
		return v60
	} else {
		v10 = int32(0)
		v11 = *(*int32)(unsafe.Add(mBase, _consts[119]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
		if v12&int32(16) == v10 {
			v47 = *(*int64)(unsafe.Add(mBase, uint32(v11)+2248))
			v48 = v47
		} else {
			if v12&int32(2) == int32(0) {
				v46 = *(*int64)(unsafe.Add(mBase, _consts[47]))
				v48 = v46
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _consts[166]))
				if v25 == int32(0) {
					v39 = int64(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, _consts[167]))
					if v29 != 0 {
						v36 = v29
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+104))
						v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+48))
						v39 = v38
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, _consts[168]))
						if v32 == int32(0) {
							v39 = int64(0)
						} else {
							v36 = v32
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+104))
							v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+48))
							v39 = v38
						}
					}
				}
				v41 = int64(0)
				if v41 < v39 {
					v44 = v39
				} else {
					v44 = v41
				}
				v48 = v44
			}
		}
		v49 = int32(0)
		if v48 == int64(0) {
			v60 = v49
			return v60
		} else {
			v53 = *(*int32)(unsafe.Add(mBase, _consts[111]))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[128])))
			if v54 != 0 {
				v60 = v49
				return v60
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[184])))
				if v55 != 0 {
					v60 = v49
					return v60
				} else {
					v56 = F_clusterAllReplicasThinkPrimaryIsFail(m)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						v60 = v56
						return v60
					}
				}
			}
		}
	}
}
