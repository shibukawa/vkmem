package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F__pqsort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
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
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
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
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1221 int32
	_ = v1221
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1277 int32
	_ = v1277
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1354 int32
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1419 int32
	_ = v1419
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1450 int32
	_ = v1450
	var v1493 int32
	_ = v1493
	var v1502 int32
	_ = v1502
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1598 int32
	_ = v1598
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1625 int32
	_ = v1625
	var v1629 int32
	_ = v1629
	var v1631 int32
	_ = v1631
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1651 int32
	_ = v1651
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1663 int32
	_ = v1663
	var v1681 int32
	_ = v1681
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1719 int32
	_ = v1719
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1746 int32
	_ = v1746
	var v1750 int32
	_ = v1750
	var v1752 int32
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1772 int32
	_ = v1772
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1796 int32
	_ = v1796
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1817 int32
	_ = v1817
	var v1834 int32
	_ = v1834
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1861 int32
	_ = v1861
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1892 int32
	_ = v1892
	var v1921 int32
	_ = v1921
	var v1924 int32
	_ = v1924
	var v1926 int32
	_ = v1926
	var v1944 int32
	_ = v1944
	var v1948 int32
	_ = v1948
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1982 int32
	_ = v1982
	var v1986 int32
	_ = v1986
	var v1988 int32
	_ = v1988
	var v2009 int32
	_ = v2009
	var v2013 int32
	_ = v2013
	var v2015 int32
	_ = v2015
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2040 int32
	_ = v2040
	var v2066 int32
	_ = v2066
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2073 int32
	_ = v2073
	var v2076 int32
	_ = v2076
	var v2078 int32
	_ = v2078
	var v2096 int32
	_ = v2096
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2115 int32
	_ = v2115
	var v2117 int32
	_ = v2117
	var v2134 int32
	_ = v2134
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2161 int32
	_ = v2161
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2187 int32
	_ = v2187
	var v2192 int32
	_ = v2192
	var v2194 int32
	_ = v2194
	var v2197 int32
	_ = v2197
	var v2199 int32
	_ = v2199
	var v2217 int32
	_ = v2217
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2255 int32
	_ = v2255
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2282 int32
	_ = v2282
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2308 int32
	_ = v2308
	var v2313 int32
	_ = v2313
	var v2344 int32
	_ = v2344
	var v2347 int32
	_ = v2347
	var v2349 int32
	_ = v2349
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2357 int32
	_ = v2357
	v26 = int32(3)
	v27 = l2 & v26
	v28 = int32(-1)
	v29 = l2 + v28
	v33 = int32(0) - l2
	v35 = int32(base.Ui32(l2) >> (uint(int32(2)) % 32))
	v37 = v35 & v26
	v39 = v35 + v28
	v40 = l0
	v41 = l1
	goto L2
L1:
	;
	return
L2:
	;
	v68 = (l2 | v40) & int32(3)
	if v68 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	goto L1
L4:
	;
	v69 = int32(2)
	goto L6
L5:
	;
	v69 = base.B2i32(l2 != int32(4))
	goto L6
L6:
	;
	if base.Ui32(int32(6)) < base.Ui32(v41) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v428 = v40 + int32(base.Ui32(v41)>>(uint(int32(1))%32))*l2
	if v41 == int32(7) {
		v541 = v428
		goto L44
	} else {
		goto L45
	}
L8:
	;
	v72 = v41 * l2
	if base.Ui32(v72) <= base.Ui32(l2) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v75 = int32(3)
	v76 = v35 & v75
	v78 = l2 & v75
	v91 = v40 + l2
	goto L10
L10:
	;
	if base.Ui32(v91) <= base.Ui32(v40) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v423 = v91 + l2
	if base.Ui32(v423) < base.Ui32(v40+v72) {
		v91 = v423
		goto L10
	} else {
		goto L43
	}
L13:
	;
	v124 = v91
	goto L14
L14:
	;
	v131 = v124 + v33
	v132 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v131, v124)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L12
L16:
	;
	return
L17:
	;
	if v132 < int32(1) {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	if v69 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if base.Ui32(v40) < base.Ui32(v131) {
		v124 = v131
		goto L14
	} else {
		goto L42
	}
L20:
	;
	if v68 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v136
	goto L19
L22:
	;
	v257 = int32(0)
	if v76 == v257 {
		v313 = v124
		v317 = v131
		v319 = v35
		goto L33
	} else {
		goto L34
	}
L23:
	;
	v142 = int32(0)
	if v78 == v142 {
		v198 = v124
		v202 = v131
		v204 = l2
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if base.Ui32(v29) < base.Ui32(int32(3)) {
		goto L19
	} else {
		goto L29
	}
L25:
	;
	v160 = v124
	v164 = v131
	v165 = v142
	v166 = l2
	goto L26
L26:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	*(*uint8)(unsafe.Add(mBase, uint32(v160))) = uint8(v171)
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v170)
	v175 = v166 + int32(-1)
	v176 = int32(1)
	v177 = v164 + v176
	v179 = v160 + v176
	v181 = v165 + v176
	if v181 != v78 {
		v160 = v179
		v164 = v177
		v165 = v181
		v166 = v175
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v198 = v179
	v202 = v177
	v204 = v175
	goto L24
L28:
	;
	goto L27
L29:
	;
	v225 = v198
	v229 = v202
	v231 = v204
	goto L30
L30:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	*(*uint8)(unsafe.Add(mBase, uint32(v225))) = uint8(v236)
	*(*uint8)(unsafe.Add(mBase, uint32(v229))) = uint8(v235)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)) = uint8(v240)
	*(*uint8)(unsafe.Add(mBase, uint32(v229)+1)) = uint8(v239)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+2)))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v225)+2)) = uint8(v244)
	*(*uint8)(unsafe.Add(mBase, uint32(v229)+2)) = uint8(v243)
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+3)))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v225)+3)) = uint8(v248)
	*(*uint8)(unsafe.Add(mBase, uint32(v229)+3)) = uint8(v247)
	v251 = int32(4)
	v256 = v231 + int32(-4)
	if v256 != 0 {
		v225 = v225 + v251
		v229 = v229 + v251
		v231 = v256
		goto L30
	} else {
		goto L32
	}
L32:
	;
	goto L19
L33:
	;
	if base.Ui32(v39) < base.Ui32(int32(3)) {
		goto L19
	} else {
		goto L38
	}
L34:
	;
	v275 = v124
	v279 = v131
	v280 = v257
	v281 = v35
	goto L35
L35:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	*(*int32)(unsafe.Add(mBase, uint32(v275))) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v279))) = v285
	v290 = v281 + int32(-1)
	v291 = int32(4)
	v292 = v279 + v291
	v294 = v275 + v291
	v296 = v280 + int32(1)
	if v296 != v76 {
		v275 = v294
		v279 = v292
		v280 = v296
		v281 = v290
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v313 = v294
	v317 = v292
	v319 = v290
	goto L33
L37:
	;
	goto L36
L38:
	;
	v340 = v313
	v344 = v317
	v346 = v319
	goto L39
L39:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	*(*int32)(unsafe.Add(mBase, uint32(v340))) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = v350
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v340)+4))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v340)+4)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v344)+4)) = v354
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v340)+8))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v344)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v340)+8)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v344)+8)) = v358
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v340)+12))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v340)+12)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v344)+12)) = v362
	v366 = int32(16)
	v371 = v346 + int32(-4)
	if v371 != 0 {
		v340 = v340 + v366
		v344 = v344 + v366
		v346 = v371
		goto L39
	} else {
		goto L41
	}
L40:
	;
	goto L19
L41:
	;
	goto L40
L42:
	;
	goto L15
L43:
	;
	goto L1
L44:
	;
	if v69 != 0 {
		goto L110
	} else {
		goto L111
	}
L45:
	;
	v434 = v40 + (v41+int32(-1))*l2
	if base.Ui32(int32(41)) <= base.Ui32(v41) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v521 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v518, v512)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L16
	} else {
		goto L95
	}
L47:
	;
	v439 = int32(base.Ui32(v41)>>(uint(int32(3))%32)) * l2
	v440 = v40 + v439
	v441 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v40, v440)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L16
	} else {
		goto L49
	}
L48:
	;
	v512 = v428
	v516 = v434
	v518 = v40
	goto L46
L49:
	;
	v444 = v439 << (uint(int32(1)) % 32)
	v445 = v40 + v444
	v446 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v440, v445)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L16
	} else {
		goto L50
	}
L50:
	;
	if int32(-1) < v441 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v467 = v428 - v439
	v468 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v467, v428)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L16
	} else {
		goto L64
	}
L52:
	;
	if int32(0) < v446 {
		v464 = v440
		goto L51
	} else {
		goto L59
	}
L53:
	;
	if v446 < int32(0) {
		v464 = v440
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v452 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v40, v445)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L16
	} else {
		goto L55
	}
L55:
	;
	if v452 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v456 = v445
	goto L58
L57:
	;
	v456 = v40
	goto L58
L58:
	;
	v464 = v456
	goto L51
L59:
	;
	v459 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v40, v445)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L16
	} else {
		goto L60
	}
L60:
	;
	if v459 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v463 = v40
	goto L63
L62:
	;
	v463 = v445
	goto L63
L63:
	;
	v464 = v463
	goto L51
L64:
	;
	v470 = v428 + v439
	v471 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v428, v470)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L16
	} else {
		goto L65
	}
L65:
	;
	if int32(-1) < v468 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v490 = v434 - v444
	v491 = v434 + (int32(0) - v439)
	v492 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v490, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L16
	} else {
		goto L79
	}
L67:
	;
	if int32(0) < v471 {
		v489 = v428
		goto L66
	} else {
		goto L74
	}
L68:
	;
	if v471 < int32(0) {
		v489 = v428
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v477 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v467, v470)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L16
	} else {
		goto L70
	}
L70:
	;
	if v477 < int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v481 = v470
	goto L73
L72:
	;
	v481 = v467
	goto L73
L73:
	;
	v489 = v481
	goto L66
L74:
	;
	v484 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v467, v470)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	if v484 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v488 = v467
	goto L78
L77:
	;
	v488 = v470
	goto L78
L78:
	;
	v489 = v488
	goto L66
L79:
	;
	v494 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v491, v434)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	if int32(-1) < v492 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if v494 <= int32(0) {
		goto L89
	} else {
		goto L90
	}
L82:
	;
	if int32(0) <= v494 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v500 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v490, v434)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L16
	} else {
		goto L85
	}
L84:
	;
	v512 = v489
	v516 = v491
	v518 = v464
	goto L46
L85:
	;
	if v500 < int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v504 = v434
	goto L88
L87:
	;
	v504 = v490
	goto L88
L88:
	;
	v512 = v489
	v516 = v504
	v518 = v464
	goto L46
L89:
	;
	v507 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v490, v434)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L16
	} else {
		goto L91
	}
L90:
	;
	v512 = v489
	v516 = v491
	v518 = v464
	goto L46
L91:
	;
	if v507 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v511 = v490
	goto L94
L93:
	;
	v511 = v434
	goto L94
L94:
	;
	v512 = v489
	v516 = v511
	v518 = v464
	goto L46
L95:
	;
	v523 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v512, v516)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L16
	} else {
		goto L96
	}
L96:
	;
	if int32(-1) < v521 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	if int32(0) < v523 {
		v541 = v512
		goto L44
	} else {
		goto L104
	}
L98:
	;
	if v523 < int32(0) {
		v541 = v512
		goto L44
	} else {
		goto L99
	}
L99:
	;
	v529 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v518, v516)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L16
	} else {
		goto L100
	}
L100:
	;
	if v529 < int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v533 = v516
	goto L103
L102:
	;
	v533 = v518
	goto L103
L103:
	;
	v541 = v533
	goto L44
L104:
	;
	v536 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v518, v516)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L16
	} else {
		goto L105
	}
L105:
	;
	if v536 < int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v540 = v518
	goto L108
L107:
	;
	v540 = v516
	goto L108
L108:
	;
	v541 = v540
	goto L44
L109:
	;
	v814 = v40 + (v41+int32(-1))*l2
	v815 = v40 + l2
	v832 = v814
	v834 = v814
	v836 = v815
	v839 = v815
	goto L133
L110:
	;
	if v68 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v541)))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v551
	*(*int32)(unsafe.Add(mBase, uint32(v541))) = v550
	goto L109
L112:
	;
	v671 = int32(0)
	if v37 == v671 {
		v727 = v541
		v731 = v40
		v733 = v35
		goto L123
	} else {
		goto L124
	}
L113:
	;
	v556 = int32(0)
	if v27 == v556 {
		v612 = v541
		v616 = v40
		v618 = l2
		goto L114
	} else {
		goto L115
	}
L114:
	;
	if base.Ui32(v29) < base.Ui32(int32(3)) {
		goto L109
	} else {
		goto L119
	}
L115:
	;
	v574 = v541
	v578 = v40
	v579 = v556
	v580 = l2
	goto L116
L116:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578))))
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	*(*uint8)(unsafe.Add(mBase, uint32(v578))) = uint8(v585)
	*(*uint8)(unsafe.Add(mBase, uint32(v574))) = uint8(v584)
	v589 = v580 + int32(-1)
	v590 = int32(1)
	v591 = v574 + v590
	v593 = v578 + v590
	v595 = v579 + v590
	if v595 != v27 {
		v574 = v591
		v578 = v593
		v579 = v595
		v580 = v589
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v612 = v591
	v616 = v593
	v618 = v589
	goto L114
L118:
	;
	goto L117
L119:
	;
	v639 = v612
	v643 = v616
	v645 = v618
	goto L120
L120:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v643))))
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639))))
	*(*uint8)(unsafe.Add(mBase, uint32(v643))) = uint8(v650)
	*(*uint8)(unsafe.Add(mBase, uint32(v639))) = uint8(v649)
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v643)+1)))
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v643)+1)) = uint8(v654)
	*(*uint8)(unsafe.Add(mBase, uint32(v639)+1)) = uint8(v653)
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v643)+2)))
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v643)+2)) = uint8(v658)
	*(*uint8)(unsafe.Add(mBase, uint32(v639)+2)) = uint8(v657)
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v643)+3)))
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v643)+3)) = uint8(v662)
	*(*uint8)(unsafe.Add(mBase, uint32(v639)+3)) = uint8(v661)
	v665 = int32(4)
	v670 = v645 + int32(-4)
	if v670 != 0 {
		v639 = v639 + v665
		v643 = v643 + v665
		v645 = v670
		goto L120
	} else {
		goto L122
	}
L122:
	;
	goto L109
L123:
	;
	if base.Ui32(v39) < base.Ui32(int32(3)) {
		goto L109
	} else {
		goto L128
	}
L124:
	;
	v689 = v541
	v693 = v40
	v694 = v671
	v695 = v35
	goto L125
L125:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v693)))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v689)))
	*(*int32)(unsafe.Add(mBase, uint32(v693))) = v700
	*(*int32)(unsafe.Add(mBase, uint32(v689))) = v699
	v704 = v695 + int32(-1)
	v705 = int32(4)
	v706 = v689 + v705
	v708 = v693 + v705
	v710 = v694 + int32(1)
	if v710 != v37 {
		v689 = v706
		v693 = v708
		v694 = v710
		v695 = v704
		goto L125
	} else {
		goto L127
	}
L126:
	;
	v727 = v706
	v731 = v708
	v733 = v704
	goto L123
L127:
	;
	goto L126
L128:
	;
	v754 = v727
	v758 = v731
	v760 = v733
	goto L129
L129:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v754)))
	*(*int32)(unsafe.Add(mBase, uint32(v758))) = v765
	*(*int32)(unsafe.Add(mBase, uint32(v754))) = v764
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v758)+4))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v754)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v758)+4)) = v769
	*(*int32)(unsafe.Add(mBase, uint32(v754)+4)) = v768
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v758)+8))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v754)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v758)+8)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v754)+8)) = v772
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v758)+12))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v754)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v758)+12)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v754)+12)) = v776
	v780 = int32(16)
	v785 = v760 + int32(-4)
	if v785 != 0 {
		v754 = v754 + v780
		v758 = v758 + v780
		v760 = v785
		goto L129
	} else {
		goto L131
	}
L130:
	;
	goto L109
L131:
	;
	goto L130
L132:
	;
	v2066 = v1520 - v1522
	v2068 = v1530 - (l2 + v1520)
	if base.Ui32(v2066) < base.Ui32(v2068) {
		goto L248
	} else {
		goto L249
	}
L133:
	;
	if base.Ui32(v834) < base.Ui32(v836) {
		v1180 = v836
		v1183 = v839
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v1921 = int32(base.Ui32(v1534) >> (uint(int32(2)) % 32))
	v1924 = int32(0)
	v1926 = v1921 & int32(3)
	if v1926 == v1924 {
		v1982 = v1537
		v1986 = v40
		v1988 = v1921
		goto L238
	} else {
		goto L239
	}
L135:
	;
	if base.Ui32(v834) < base.Ui32(v1180) {
		v1520 = v832
		v1522 = v834
		goto L169
	} else {
		goto L170
	}
L136:
	;
	v862 = v836
	v865 = v839
	goto L137
L137:
	;
	v867 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v862, v40)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L16
	} else {
		goto L139
	}
L138:
	;
	v1180 = v1158
	v1183 = v1156
	goto L135
L139:
	;
	if int32(0) < v867 {
		v1180 = v862
		v1183 = v865
		goto L135
	} else {
		goto L140
	}
L140:
	;
	if v867 != 0 {
		v1156 = v865
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v1158 = v862 + l2
	if base.Ui32(v1158) <= base.Ui32(v834) {
		v862 = v1158
		v865 = v1156
		goto L137
	} else {
		goto L166
	}
L142:
	;
	if v69 != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v1156 = v865 + l2
	goto L141
L144:
	;
	if v68 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v865)))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v862)))
	*(*int32)(unsafe.Add(mBase, uint32(v865))) = v872
	*(*int32)(unsafe.Add(mBase, uint32(v862))) = v871
	goto L143
L146:
	;
	v992 = int32(0)
	if v37 == v992 {
		v1048 = v862
		v1052 = v865
		v1054 = v35
		goto L157
	} else {
		goto L158
	}
L147:
	;
	v877 = int32(0)
	if v27 == v877 {
		v933 = v865
		v937 = v862
		v939 = l2
		goto L148
	} else {
		goto L149
	}
L148:
	;
	if base.Ui32(v29) < base.Ui32(int32(3)) {
		goto L143
	} else {
		goto L153
	}
L149:
	;
	v895 = v865
	v899 = v862
	v901 = l2
	v902 = v877
	goto L150
L150:
	;
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895))))
	v906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v899))))
	*(*uint8)(unsafe.Add(mBase, uint32(v895))) = uint8(v906)
	*(*uint8)(unsafe.Add(mBase, uint32(v899))) = uint8(v905)
	v910 = v901 + int32(-1)
	v911 = int32(1)
	v912 = v899 + v911
	v914 = v895 + v911
	v916 = v902 + v911
	if v916 != v27 {
		v895 = v914
		v899 = v912
		v901 = v910
		v902 = v916
		goto L150
	} else {
		goto L152
	}
L151:
	;
	v933 = v914
	v937 = v912
	v939 = v910
	goto L148
L152:
	;
	goto L151
L153:
	;
	v960 = v933
	v964 = v937
	v966 = v939
	goto L154
L154:
	;
	v970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v960))))
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964))))
	*(*uint8)(unsafe.Add(mBase, uint32(v960))) = uint8(v971)
	*(*uint8)(unsafe.Add(mBase, uint32(v964))) = uint8(v970)
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v960)+1)))
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v960)+1)) = uint8(v975)
	*(*uint8)(unsafe.Add(mBase, uint32(v964)+1)) = uint8(v974)
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v960)+2)))
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v960)+2)) = uint8(v979)
	*(*uint8)(unsafe.Add(mBase, uint32(v964)+2)) = uint8(v978)
	v982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v960)+3)))
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v960)+3)) = uint8(v983)
	*(*uint8)(unsafe.Add(mBase, uint32(v964)+3)) = uint8(v982)
	v986 = int32(4)
	v991 = v966 + int32(-4)
	if v991 != 0 {
		v960 = v960 + v986
		v964 = v964 + v986
		v966 = v991
		goto L154
	} else {
		goto L156
	}
L156:
	;
	goto L143
L157:
	;
	if base.Ui32(v39) < base.Ui32(int32(3)) {
		goto L143
	} else {
		goto L162
	}
L158:
	;
	v1010 = v862
	v1014 = v865
	v1016 = v35
	v1017 = v992
	goto L159
L159:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1010)))
	*(*int32)(unsafe.Add(mBase, uint32(v1014))) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v1010))) = v1020
	v1025 = v1016 + int32(-1)
	v1026 = int32(4)
	v1027 = v1010 + v1026
	v1029 = v1014 + v1026
	v1031 = v1017 + int32(1)
	if v1031 != v37 {
		v1010 = v1027
		v1014 = v1029
		v1016 = v1025
		v1017 = v1031
		goto L159
	} else {
		goto L161
	}
L160:
	;
	v1048 = v1027
	v1052 = v1029
	v1054 = v1025
	goto L157
L161:
	;
	goto L160
L162:
	;
	v1075 = v1048
	v1079 = v1052
	v1081 = v1054
	goto L163
L163:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1079)))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1075)))
	*(*int32)(unsafe.Add(mBase, uint32(v1079))) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v1075))) = v1085
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+4))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+4)) = v1090
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+4)) = v1089
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+8))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+8)) = v1094
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+8)) = v1093
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+12))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+12)) = v1098
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+12)) = v1097
	v1101 = int32(16)
	v1106 = v1081 + int32(-4)
	if v1106 != 0 {
		v1075 = v1075 + v1101
		v1079 = v1079 + v1101
		v1081 = v1106
		goto L163
	} else {
		goto L165
	}
L164:
	;
	goto L143
L165:
	;
	goto L164
L166:
	;
	goto L138
L167:
	;
	goto L134
L168:
	;
	if v69 != 0 {
		goto L216
	} else {
		goto L217
	}
L169:
	;
	v1530 = v40 + v41*l2
	v1531 = v1183 - v40
	v1532 = v1180 - v1183
	if v1531 < v1532 {
		goto L201
	} else {
		goto L202
	}
L170:
	;
	v1202 = v832
	v1204 = v834
	goto L171
L171:
	;
	v1211 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, v1204, v40)
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L16
	} else {
		goto L173
	}
L172:
	;
	v1520 = v1493
	v1522 = v1502
	goto L169
L173:
	;
	if v1211 < int32(0) {
		goto L168
	} else {
		goto L174
	}
L174:
	;
	if v1211 != 0 {
		v1493 = v1202
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v1502 = v1204 + v33
	if base.Ui32(v1180) <= base.Ui32(v1502) {
		v1202 = v1493
		v1204 = v1502
		goto L171
	} else {
		goto L200
	}
L176:
	;
	if v69 != 0 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v1493 = v1202 + v33
	goto L175
L178:
	;
	if v68 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1204)))
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1202)))
	*(*int32)(unsafe.Add(mBase, uint32(v1204))) = v1216
	*(*int32)(unsafe.Add(mBase, uint32(v1202))) = v1215
	goto L177
L180:
	;
	v1336 = int32(0)
	if v37 == v1336 {
		v1392 = v1202
		v1396 = v1204
		v1398 = v35
		goto L191
	} else {
		goto L192
	}
L181:
	;
	v1221 = int32(0)
	if v27 == v1221 {
		v1277 = v1204
		v1281 = v1202
		v1283 = l2
		goto L182
	} else {
		goto L183
	}
L182:
	;
	if base.Ui32(v29) < base.Ui32(int32(3)) {
		goto L177
	} else {
		goto L187
	}
L183:
	;
	v1239 = v1204
	v1243 = v1202
	v1245 = l2
	v1246 = v1221
	goto L184
L184:
	;
	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239))))
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1239))) = uint8(v1250)
	*(*uint8)(unsafe.Add(mBase, uint32(v1243))) = uint8(v1249)
	v1254 = v1245 + int32(-1)
	v1255 = int32(1)
	v1256 = v1243 + v1255
	v1258 = v1239 + v1255
	v1260 = v1246 + v1255
	if v1260 != v27 {
		v1239 = v1258
		v1243 = v1256
		v1245 = v1254
		v1246 = v1260
		goto L184
	} else {
		goto L186
	}
L185:
	;
	v1277 = v1258
	v1281 = v1256
	v1283 = v1254
	goto L182
L186:
	;
	goto L185
L187:
	;
	v1304 = v1277
	v1308 = v1281
	v1310 = v1283
	goto L188
L188:
	;
	v1314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1304))))
	v1315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1308))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1304))) = uint8(v1315)
	*(*uint8)(unsafe.Add(mBase, uint32(v1308))) = uint8(v1314)
	v1318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1304)+1)))
	v1319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1308)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1304)+1)) = uint8(v1319)
	*(*uint8)(unsafe.Add(mBase, uint32(v1308)+1)) = uint8(v1318)
	v1322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1304)+2)))
	v1323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1308)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1304)+2)) = uint8(v1323)
	*(*uint8)(unsafe.Add(mBase, uint32(v1308)+2)) = uint8(v1322)
	v1326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1304)+3)))
	v1327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1308)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1304)+3)) = uint8(v1327)
	*(*uint8)(unsafe.Add(mBase, uint32(v1308)+3)) = uint8(v1326)
	v1330 = int32(4)
	v1335 = v1310 + int32(-4)
	if v1335 != 0 {
		v1304 = v1304 + v1330
		v1308 = v1308 + v1330
		v1310 = v1335
		goto L188
	} else {
		goto L190
	}
L190:
	;
	goto L177
L191:
	;
	if base.Ui32(v39) < base.Ui32(int32(3)) {
		goto L177
	} else {
		goto L196
	}
L192:
	;
	v1354 = v1202
	v1358 = v1204
	v1360 = v35
	v1361 = v1336
	goto L193
L193:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1358)))
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1354)))
	*(*int32)(unsafe.Add(mBase, uint32(v1358))) = v1365
	*(*int32)(unsafe.Add(mBase, uint32(v1354))) = v1364
	v1369 = v1360 + int32(-1)
	v1370 = int32(4)
	v1371 = v1354 + v1370
	v1373 = v1358 + v1370
	v1375 = v1361 + int32(1)
	if v1375 != v37 {
		v1354 = v1371
		v1358 = v1373
		v1360 = v1369
		v1361 = v1375
		goto L193
	} else {
		goto L195
	}
L194:
	;
	v1392 = v1371
	v1396 = v1373
	v1398 = v1369
	goto L191
L195:
	;
	goto L194
L196:
	;
	v1419 = v1392
	v1423 = v1396
	v1425 = v1398
	goto L197
L197:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1423)))
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1419)))
	*(*int32)(unsafe.Add(mBase, uint32(v1423))) = v1430
	*(*int32)(unsafe.Add(mBase, uint32(v1419))) = v1429
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v1423)+4))
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1423)+4)) = v1434
	*(*int32)(unsafe.Add(mBase, uint32(v1419)+4)) = v1433
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v1423)+8))
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1423)+8)) = v1438
	*(*int32)(unsafe.Add(mBase, uint32(v1419)+8)) = v1437
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1423)+12))
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1423)+12)) = v1442
	*(*int32)(unsafe.Add(mBase, uint32(v1419)+12)) = v1441
	v1445 = int32(16)
	v1450 = v1425 + int32(-4)
	if v1450 != 0 {
		v1419 = v1419 + v1445
		v1423 = v1423 + v1445
		v1425 = v1450
		goto L197
	} else {
		goto L199
	}
L198:
	;
	goto L177
L199:
	;
	goto L198
L200:
	;
	goto L172
L201:
	;
	v1534 = v1531
	goto L203
L202:
	;
	v1534 = v1532
	goto L203
L203:
	;
	if v1534 == int32(0) {
		goto L132
	} else {
		goto L204
	}
L204:
	;
	v1537 = v1180 - v1534
	if v68 == int32(0) {
		goto L167
	} else {
		goto L205
	}
L205:
	;
	v1540 = int32(0)
	v1542 = v1534 & int32(3)
	if v1542 == v1540 {
		v1598 = v1537
		v1602 = v40
		v1604 = v1534
		goto L206
	} else {
		goto L207
	}
L206:
	;
	if base.Ui32(v1534) < base.Ui32(int32(4)) {
		goto L132
	} else {
		goto L211
	}
L207:
	;
	v1560 = v1537
	v1564 = v40
	v1566 = v1534
	v1567 = v1540
	goto L208
L208:
	;
	v1570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1564))))
	v1571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1560))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1564))) = uint8(v1571)
	*(*uint8)(unsafe.Add(mBase, uint32(v1560))) = uint8(v1570)
	v1575 = v1566 + int32(-1)
	v1576 = int32(1)
	v1577 = v1560 + v1576
	v1579 = v1564 + v1576
	v1581 = v1567 + v1576
	if v1581 != v1542 {
		v1560 = v1577
		v1564 = v1579
		v1566 = v1575
		v1567 = v1581
		goto L208
	} else {
		goto L210
	}
L209:
	;
	v1598 = v1577
	v1602 = v1579
	v1604 = v1575
	goto L206
L210:
	;
	goto L209
L211:
	;
	v1625 = v1598
	v1629 = v1602
	v1631 = v1604
	goto L212
L212:
	;
	v1635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1629))))
	v1636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1625))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1629))) = uint8(v1636)
	*(*uint8)(unsafe.Add(mBase, uint32(v1625))) = uint8(v1635)
	v1639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1629)+1)))
	v1640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1625)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1629)+1)) = uint8(v1640)
	*(*uint8)(unsafe.Add(mBase, uint32(v1625)+1)) = uint8(v1639)
	v1643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1629)+2)))
	v1644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1625)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1629)+2)) = uint8(v1644)
	*(*uint8)(unsafe.Add(mBase, uint32(v1625)+2)) = uint8(v1643)
	v1647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1629)+3)))
	v1648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1625)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1629)+3)) = uint8(v1648)
	*(*uint8)(unsafe.Add(mBase, uint32(v1625)+3)) = uint8(v1647)
	v1651 = int32(4)
	v1656 = v1631 + int32(-4)
	if v1656 != 0 {
		v1625 = v1625 + v1651
		v1629 = v1629 + v1651
		v1631 = v1656
		goto L212
	} else {
		goto L214
	}
L214:
	;
	goto L132
L215:
	;
	v832 = v1202
	v834 = v1204 + v33
	v836 = v1180 + l2
	v839 = v1183
	goto L133
L216:
	;
	if v68 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1180)))
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v1204)))
	*(*int32)(unsafe.Add(mBase, uint32(v1180))) = v1658
	*(*int32)(unsafe.Add(mBase, uint32(v1204))) = v1657
	goto L215
L218:
	;
	v1778 = int32(0)
	if v37 == v1778 {
		v1834 = v1204
		v1838 = v1180
		v1840 = v35
		goto L229
	} else {
		goto L230
	}
L219:
	;
	v1663 = int32(0)
	if v27 == v1663 {
		v1719 = v1180
		v1723 = v1204
		v1725 = l2
		goto L220
	} else {
		goto L221
	}
L220:
	;
	if base.Ui32(v29) < base.Ui32(int32(3)) {
		goto L215
	} else {
		goto L225
	}
L221:
	;
	v1681 = v1180
	v1685 = v1204
	v1687 = l2
	v1688 = v1663
	goto L222
L222:
	;
	v1691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1681))))
	v1692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1685))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1681))) = uint8(v1692)
	*(*uint8)(unsafe.Add(mBase, uint32(v1685))) = uint8(v1691)
	v1696 = v1687 + int32(-1)
	v1697 = int32(1)
	v1698 = v1685 + v1697
	v1700 = v1681 + v1697
	v1702 = v1688 + v1697
	if v1702 != v27 {
		v1681 = v1700
		v1685 = v1698
		v1687 = v1696
		v1688 = v1702
		goto L222
	} else {
		goto L224
	}
L223:
	;
	v1719 = v1700
	v1723 = v1698
	v1725 = v1696
	goto L220
L224:
	;
	goto L223
L225:
	;
	v1746 = v1719
	v1750 = v1723
	v1752 = v1725
	goto L226
L226:
	;
	v1756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1746))))
	v1757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1750))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1746))) = uint8(v1757)
	*(*uint8)(unsafe.Add(mBase, uint32(v1750))) = uint8(v1756)
	v1760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1746)+1)))
	v1761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1750)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1746)+1)) = uint8(v1761)
	*(*uint8)(unsafe.Add(mBase, uint32(v1750)+1)) = uint8(v1760)
	v1764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1746)+2)))
	v1765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1750)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1746)+2)) = uint8(v1765)
	*(*uint8)(unsafe.Add(mBase, uint32(v1750)+2)) = uint8(v1764)
	v1768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1746)+3)))
	v1769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1750)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1746)+3)) = uint8(v1769)
	*(*uint8)(unsafe.Add(mBase, uint32(v1750)+3)) = uint8(v1768)
	v1772 = int32(4)
	v1777 = v1752 + int32(-4)
	if v1777 != 0 {
		v1746 = v1746 + v1772
		v1750 = v1750 + v1772
		v1752 = v1777
		goto L226
	} else {
		goto L228
	}
L228:
	;
	goto L215
L229:
	;
	if base.Ui32(v39) < base.Ui32(int32(3)) {
		goto L215
	} else {
		goto L234
	}
L230:
	;
	v1796 = v1204
	v1800 = v1180
	v1802 = v35
	v1803 = v1778
	goto L231
L231:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1800)))
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1796)))
	*(*int32)(unsafe.Add(mBase, uint32(v1800))) = v1807
	*(*int32)(unsafe.Add(mBase, uint32(v1796))) = v1806
	v1811 = v1802 + int32(-1)
	v1812 = int32(4)
	v1813 = v1796 + v1812
	v1815 = v1800 + v1812
	v1817 = v1803 + int32(1)
	if v1817 != v37 {
		v1796 = v1813
		v1800 = v1815
		v1802 = v1811
		v1803 = v1817
		goto L231
	} else {
		goto L233
	}
L232:
	;
	v1834 = v1813
	v1838 = v1815
	v1840 = v1811
	goto L229
L233:
	;
	goto L232
L234:
	;
	v1861 = v1834
	v1865 = v1838
	v1867 = v1840
	goto L235
L235:
	;
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1865)))
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1861)))
	*(*int32)(unsafe.Add(mBase, uint32(v1865))) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v1861))) = v1871
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+4))
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1861)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1865)+4)) = v1876
	*(*int32)(unsafe.Add(mBase, uint32(v1861)+4)) = v1875
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+8))
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1861)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1865)+8)) = v1880
	*(*int32)(unsafe.Add(mBase, uint32(v1861)+8)) = v1879
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+12))
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1861)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1865)+12)) = v1884
	*(*int32)(unsafe.Add(mBase, uint32(v1861)+12)) = v1883
	v1887 = int32(16)
	v1892 = v1867 + int32(-4)
	if v1892 != 0 {
		v1861 = v1861 + v1887
		v1865 = v1865 + v1887
		v1867 = v1892
		goto L235
	} else {
		goto L237
	}
L236:
	;
	goto L215
L237:
	;
	goto L236
L238:
	;
	if base.Ui32(v1921+int32(-1)) < base.Ui32(int32(3)) {
		goto L132
	} else {
		goto L243
	}
L239:
	;
	v1944 = v1537
	v1948 = v40
	v1950 = v1921
	v1951 = v1924
	goto L240
L240:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v1948)))
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(v1944)))
	*(*int32)(unsafe.Add(mBase, uint32(v1948))) = v1955
	*(*int32)(unsafe.Add(mBase, uint32(v1944))) = v1954
	v1959 = v1950 + int32(-1)
	v1960 = int32(4)
	v1961 = v1944 + v1960
	v1963 = v1948 + v1960
	v1965 = v1951 + int32(1)
	if v1965 != v1926 {
		v1944 = v1961
		v1948 = v1963
		v1950 = v1959
		v1951 = v1965
		goto L240
	} else {
		goto L242
	}
L241:
	;
	v1982 = v1961
	v1986 = v1963
	v1988 = v1959
	goto L238
L242:
	;
	goto L241
L243:
	;
	v2009 = v1982
	v2013 = v1986
	v2015 = v1988
	goto L244
L244:
	;
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(v2013)))
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v2009)))
	*(*int32)(unsafe.Add(mBase, uint32(v2013))) = v2020
	*(*int32)(unsafe.Add(mBase, uint32(v2009))) = v2019
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v2013)+4))
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v2009)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2013)+4)) = v2024
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+4)) = v2023
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v2013)+8))
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v2009)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2013)+8)) = v2028
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+8)) = v2027
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v2013)+12))
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v2009)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2013)+12)) = v2032
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+12)) = v2031
	v2035 = int32(16)
	v2040 = v2015 + int32(-4)
	if v2040 != 0 {
		v2009 = v2009 + v2035
		v2013 = v2013 + v2035
		v2015 = v2040
		goto L244
	} else {
		goto L246
	}
L245:
	;
	goto L132
L246:
	;
	goto L245
L247:
	;
	if base.Ui32(v1532) <= base.Ui32(l2) {
		goto L272
	} else {
		goto L273
	}
L248:
	;
	v2070 = v2066
	goto L250
L249:
	;
	v2070 = v2068
	goto L250
L250:
	;
	if v2070 == int32(0) {
		goto L247
	} else {
		goto L251
	}
L251:
	;
	v2073 = v1530 - v2070
	if v68 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v2194 = int32(base.Ui32(v2070) >> (uint(int32(2)) % 32))
	v2197 = int32(0)
	v2199 = v2194 & int32(3)
	if v2199 == v2197 {
		v2255 = v2073
		v2259 = v2194
		v2260 = v1180
		goto L263
	} else {
		goto L264
	}
L253:
	;
	v2076 = int32(0)
	v2078 = v2070 & int32(3)
	if v2078 == v2076 {
		v2134 = v2073
		v2138 = v2070
		v2139 = v1180
		goto L254
	} else {
		goto L255
	}
L254:
	;
	if base.Ui32(v2070) < base.Ui32(int32(4)) {
		goto L247
	} else {
		goto L259
	}
L255:
	;
	v2096 = v2073
	v2100 = v2070
	v2101 = v1180
	v2102 = v2076
	goto L256
L256:
	;
	v2106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2101))))
	v2107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2096))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2101))) = uint8(v2107)
	*(*uint8)(unsafe.Add(mBase, uint32(v2096))) = uint8(v2106)
	v2111 = v2100 + int32(-1)
	v2112 = int32(1)
	v2113 = v2096 + v2112
	v2115 = v2101 + v2112
	v2117 = v2102 + v2112
	if v2117 != v2078 {
		v2096 = v2113
		v2100 = v2111
		v2101 = v2115
		v2102 = v2117
		goto L256
	} else {
		goto L258
	}
L257:
	;
	v2134 = v2113
	v2138 = v2111
	v2139 = v2115
	goto L254
L258:
	;
	goto L257
L259:
	;
	v2161 = v2134
	v2165 = v2138
	v2166 = v2139
	goto L260
L260:
	;
	v2171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2166))))
	v2172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2161))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2166))) = uint8(v2172)
	*(*uint8)(unsafe.Add(mBase, uint32(v2161))) = uint8(v2171)
	v2175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2166)+1)))
	v2176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2161)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2166)+1)) = uint8(v2176)
	*(*uint8)(unsafe.Add(mBase, uint32(v2161)+1)) = uint8(v2175)
	v2179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2166)+2)))
	v2180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2161)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2166)+2)) = uint8(v2180)
	*(*uint8)(unsafe.Add(mBase, uint32(v2161)+2)) = uint8(v2179)
	v2183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2166)+3)))
	v2184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2161)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2166)+3)) = uint8(v2184)
	*(*uint8)(unsafe.Add(mBase, uint32(v2161)+3)) = uint8(v2183)
	v2187 = int32(4)
	v2192 = v2165 + int32(-4)
	if v2192 != 0 {
		v2161 = v2161 + v2187
		v2165 = v2192
		v2166 = v2166 + v2187
		goto L260
	} else {
		goto L262
	}
L262:
	;
	goto L247
L263:
	;
	if base.Ui32(v2194+int32(-1)) < base.Ui32(int32(3)) {
		goto L247
	} else {
		goto L268
	}
L264:
	;
	v2217 = v2073
	v2221 = v2194
	v2222 = v1180
	v2223 = v2197
	goto L265
L265:
	;
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v2222)))
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v2217)))
	*(*int32)(unsafe.Add(mBase, uint32(v2222))) = v2228
	*(*int32)(unsafe.Add(mBase, uint32(v2217))) = v2227
	v2232 = v2221 + int32(-1)
	v2233 = int32(4)
	v2234 = v2217 + v2233
	v2236 = v2222 + v2233
	v2238 = v2223 + int32(1)
	if v2238 != v2199 {
		v2217 = v2234
		v2221 = v2232
		v2222 = v2236
		v2223 = v2238
		goto L265
	} else {
		goto L267
	}
L266:
	;
	v2255 = v2234
	v2259 = v2232
	v2260 = v2236
	goto L263
L267:
	;
	goto L266
L268:
	;
	v2282 = v2255
	v2286 = v2259
	v2287 = v2260
	goto L269
L269:
	;
	v2292 = *(*int32)(unsafe.Add(mBase, uint32(v2287)))
	v2293 = *(*int32)(unsafe.Add(mBase, uint32(v2282)))
	*(*int32)(unsafe.Add(mBase, uint32(v2287))) = v2293
	*(*int32)(unsafe.Add(mBase, uint32(v2282))) = v2292
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2287)+4))
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v2282)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2287)+4)) = v2297
	*(*int32)(unsafe.Add(mBase, uint32(v2282)+4)) = v2296
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v2287)+8))
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v2282)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2287)+8)) = v2301
	*(*int32)(unsafe.Add(mBase, uint32(v2282)+8)) = v2300
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v2287)+12))
	v2305 = *(*int32)(unsafe.Add(mBase, uint32(v2282)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2287)+12)) = v2305
	*(*int32)(unsafe.Add(mBase, uint32(v2282)+12)) = v2304
	v2308 = int32(16)
	v2313 = v2286 + int32(-4)
	if v2313 != 0 {
		v2282 = v2282 + v2308
		v2286 = v2313
		v2287 = v2287 + v2308
		goto L269
	} else {
		goto L271
	}
L270:
	;
	goto L247
L271:
	;
	goto L270
L272:
	;
	if base.Ui32(v2066) <= base.Ui32(l2) {
		goto L1
	} else {
		goto L281
	}
L273:
	;
	if base.Ui32(v40) <= base.Ui32(l4) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v2344 = v40 + v1532 + int32(-1)
	if base.Ui32(l4) <= base.Ui32(v2344) {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	if base.Ui32(l5) < base.Ui32(v40) {
		goto L272
	} else {
		goto L276
	}
L276:
	;
	goto L274
L277:
	;
	v2347 = base.I32_div_u_s(v1532, l2)
	F__pqsort(m, v40, v2347, l2, l3, l4, l5)
	mBase = m.M
	v2349 = m.ExcPending
	if v2349 != 0 {
		goto L16
	} else {
		goto L280
	}
L278:
	;
	if base.Ui32(v2344) < base.Ui32(l5) {
		goto L272
	} else {
		goto L279
	}
L279:
	;
	goto L277
L280:
	;
	goto L272
L281:
	;
	v2352 = v1530 - v2066
	if base.Ui32(v2352) <= base.Ui32(l4) {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v2355 = base.I32_div_u_s(v2066, l2)
	v2357 = v1530 + int32(-1)
	if base.Ui32(l4) <= base.Ui32(v2357) {
		v40 = v2352
		v41 = v2355
		goto L2
	} else {
		goto L285
	}
L283:
	;
	if base.Ui32(l5) < base.Ui32(v2352) {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	goto L282
L285:
	;
	if base.Ui32(l5) <= base.Ui32(v2357) {
		v40 = v2352
		v41 = v2355
		goto L2
	} else {
		goto L286
	}
L286:
	;
	goto L3
}
func F_parseInputBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
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
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	if v5 != 0 {
		F__serverAssert(m, int32(_a806), int32(_a774), int32(4020))
		mBase = m.M
		v63 = m.ExcPending
		if v63 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v6 == int32(0) {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			switch v24 {
			case 0:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v26))))
				if v28 != int32(42) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(1)
					F_parseInlineBuffer(m, l0)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						if v6 == int32(0) {
							return
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
							if v43 == int32(0) {
								return
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
								if v46 == int32(0) {
									return
								} else {
									v50 = base.B2i32(v41 != int32(0))
									if v42 != 0 {
										v53 = v50 | int32(2)
									} else {
										v53 = v50
									}
									m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(2)
					F_parseMultibulkBuffer(m, l0)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						if v6 == int32(0) {
							return
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
							if v43 == int32(0) {
								return
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
								if v46 == int32(0) {
									return
								} else {
									v50 = base.B2i32(v41 != int32(0))
									if v42 != 0 {
										v53 = v50 | int32(2)
									} else {
										v53 = v50
									}
									m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			case 1:
				F_parseInlineBuffer(m, l0)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					if v6 == int32(0) {
						return
					} else {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						if v43 == int32(0) {
							return
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
							if v46 == int32(0) {
								return
							} else {
								v50 = base.B2i32(v41 != int32(0))
								if v42 != 0 {
									v53 = v50 | int32(2)
								} else {
									v53 = v50
								}
								m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			case 2:
				F_parseMultibulkBuffer(m, l0)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					if v6 == int32(0) {
						return
					} else {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						if v43 == int32(0) {
							return
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
							if v46 == int32(0) {
								return
							} else {
								v50 = base.B2i32(v41 != int32(0))
								if v42 != 0 {
									v53 = v50 | int32(2)
								} else {
									v53 = v50
								}
								m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			default:
				F__serverPanic_1(m, int32(_a774), int32(4041), int32(_a807), int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			if v11 == int32(0) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				switch v24 {
				case 0:
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v26))))
					if v28 != int32(42) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(1)
						F_parseInlineBuffer(m, l0)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							if v6 == int32(0) {
								return
							} else {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
								if v43 == int32(0) {
									return
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
									if v46 == int32(0) {
										return
									} else {
										v50 = base.B2i32(v41 != int32(0))
										if v42 != 0 {
											v53 = v50 | int32(2)
										} else {
											v53 = v50
										}
										m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(2)
						F_parseMultibulkBuffer(m, l0)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							if v6 == int32(0) {
								return
							} else {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
								if v43 == int32(0) {
									return
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
									if v46 == int32(0) {
										return
									} else {
										v50 = base.B2i32(v41 != int32(0))
										if v42 != 0 {
											v53 = v50 | int32(2)
										} else {
											v53 = v50
										}
										m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					}
				case 1:
					F_parseInlineBuffer(m, l0)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						if v6 == int32(0) {
							return
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
							if v43 == int32(0) {
								return
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
								if v46 == int32(0) {
									return
								} else {
									v50 = base.B2i32(v41 != int32(0))
									if v42 != 0 {
										v53 = v50 | int32(2)
									} else {
										v53 = v50
									}
									m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				case 2:
					F_parseMultibulkBuffer(m, l0)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						if v6 == int32(0) {
							return
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
							if v43 == int32(0) {
								return
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
								if v46 == int32(0) {
									return
								} else {
									v50 = base.B2i32(v41 != int32(0))
									if v42 != 0 {
										v53 = v50 | int32(2)
									} else {
										v53 = v50
									}
									m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				default:
					F__serverPanic_1(m, int32(_a774), int32(4041), int32(_a807), int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+112))
				if v14 == int32(0) {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					switch v24 {
					case 0:
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v26))))
						if v28 != int32(42) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(1)
							F_parseInlineBuffer(m, l0)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								if v6 == int32(0) {
									return
								} else {
									v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
									v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
									if v43 == int32(0) {
										return
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
										if v46 == int32(0) {
											return
										} else {
											v50 = base.B2i32(v41 != int32(0))
											if v42 != 0 {
												v53 = v50 | int32(2)
											} else {
												v53 = v50
											}
											m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(2)
							F_parseMultibulkBuffer(m, l0)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								if v6 == int32(0) {
									return
								} else {
									v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
									v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
									if v43 == int32(0) {
										return
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
										if v46 == int32(0) {
											return
										} else {
											v50 = base.B2i32(v41 != int32(0))
											if v42 != 0 {
												v53 = v50 | int32(2)
											} else {
												v53 = v50
											}
											m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						}
					case 1:
						F_parseInlineBuffer(m, l0)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							if v6 == int32(0) {
								return
							} else {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
								if v43 == int32(0) {
									return
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
									if v46 == int32(0) {
										return
									} else {
										v50 = base.B2i32(v41 != int32(0))
										if v42 != 0 {
											v53 = v50 | int32(2)
										} else {
											v53 = v50
										}
										m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					case 2:
						F_parseMultibulkBuffer(m, l0)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							if v6 == int32(0) {
								return
							} else {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
								if v43 == int32(0) {
									return
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
									if v46 == int32(0) {
										return
									} else {
										v50 = base.B2i32(v41 != int32(0))
										if v42 != 0 {
											v53 = v50 | int32(2)
										} else {
											v53 = v50
										}
										m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					default:
						F__serverPanic_1(m, int32(_a774), int32(4041), int32(_a807), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				} else {
					if v10 != 0 {
						v19 = int32(3)
					} else {
						v19 = int32(1)
					}
					m.T0[v14].(func(*base.Module, int32, int32))(m, v6, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						switch v24 {
						case 0:
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v26))))
							if v28 != int32(42) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(1)
								F_parseInlineBuffer(m, l0)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									if v6 == int32(0) {
										return
									} else {
										v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
										v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
										if v43 == int32(0) {
											return
										} else {
											v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
											if v46 == int32(0) {
												return
											} else {
												v50 = base.B2i32(v41 != int32(0))
												if v42 != 0 {
													v53 = v50 | int32(2)
												} else {
													v53 = v50
												}
												m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
												mBase = m.M
												v55 = m.ExcPending
												if v55 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(2)
								F_parseMultibulkBuffer(m, l0)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									if v6 == int32(0) {
										return
									} else {
										v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
										v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
										if v43 == int32(0) {
											return
										} else {
											v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
											if v46 == int32(0) {
												return
											} else {
												v50 = base.B2i32(v41 != int32(0))
												if v42 != 0 {
													v53 = v50 | int32(2)
												} else {
													v53 = v50
												}
												m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
												mBase = m.M
												v55 = m.ExcPending
												if v55 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								}
							}
						case 1:
							F_parseInlineBuffer(m, l0)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								if v6 == int32(0) {
									return
								} else {
									v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
									v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
									if v43 == int32(0) {
										return
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
										if v46 == int32(0) {
											return
										} else {
											v50 = base.B2i32(v41 != int32(0))
											if v42 != 0 {
												v53 = v50 | int32(2)
											} else {
												v53 = v50
											}
											m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						case 2:
							F_parseMultibulkBuffer(m, l0)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								if v6 == int32(0) {
									return
								} else {
									v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
									v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
									if v43 == int32(0) {
										return
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+112))
										if v46 == int32(0) {
											return
										} else {
											v50 = base.B2i32(v41 != int32(0))
											if v42 != 0 {
												v53 = v50 | int32(2)
											} else {
												v53 = v50
											}
											m.T0[v46].(func(*base.Module, int32, int32))(m, v6, v53)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						default:
							F__serverPanic_1(m, int32(_a774), int32(4041), int32(_a807), int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
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
				}
			}
		}
	}
}
func F_parseScanCursorOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v78 int64
	_ = v78
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v100 int64
	_ = v100
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v121 int64
	_ = v121
	var v132 int64
	_ = v132
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int64
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int64
	_ = v178
	var v185 int32
	_ = v185
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	v4 = int32(0)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v11 & int32(7) {
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
		v28 = v4
		goto L1
	}
L1:
	;
	v36 = m.G0
	v38 = v36 - int32(16)
	m.G0 = v38
	if base.Ui32(v28+int32(-21)) < base.Ui32(int32(-20)) {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v28 = v27
	goto L1
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v28 = v24
	goto L1
L4:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v28 = v21
	goto L1
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v28 = v18
	goto L1
L6:
	;
	v28 = int32(base.Ui32(v11) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	return v203
L8:
	;
	if v185 != 0 {
		v203 = v4
		goto L7
	} else {
		goto L35
	}
L9:
	;
	m.G0 = v38 + int32(16)
	goto L8
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v178
	v185 = int32(1)
	goto L9
L11:
	;
	v149 = int32(0)
	v150 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v149
	v158 = F_strtoull(m, l1, v38+int32(12), int32(10))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if v160 == int32(28) {
		v185 = v149
		goto L9
	} else {
		goto L32
	}
L12:
	;
	v44 = int32(1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v28 != v44 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v45&int32(255) != int32(45) {
		v65 = v44
		v66 = v45
		v67 = l1
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v49 = v45 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v49&int32(255)) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v178 = base.I64_extend_i32_u(v49) & int64(255)
	goto L10
L16:
	;
	if base.Ui32(int32(8)) < base.Ui32((v66+int32(-49))&int32(255)) {
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v65 = int32(2)
	v66 = v63
	v67 = l1 + int32(1)
	goto L16
L18:
	;
	v78 = base.I64_extend_i32_u(v66+int32(-48)) & int64(255)
	if base.Ui32(v28) <= base.Ui32(v65) {
		v121 = v78
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v45&int32(255) != int32(45) {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v84 = v65
	v86 = v78
	v88 = v67
	goto L21
L21:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	if base.Ui32((v90+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L11
	} else {
		goto L23
	}
L22:
	;
	v121 = v111
	goto L19
L23:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v86) {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v100 = v86 * int64(10)
	v105 = base.I64_extend_i32_u(v90+int32(-48)) & int64(255)
	if base.Ui64(v105^int64(-1)) < base.Ui64(v100) {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v109 = int32(1)
	v111 = v100 + v105
	v113 = v84 + v109
	if v113 != v28 {
		v84 = v113
		v86 = v111
		v88 = v88 + v109
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	if int64(0) <= v121 {
		v178 = v121
		goto L10
	} else {
		goto L31
	}
L28:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v121) {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v132 = int64(-1)
	if v132 < v121+v132 {
		v185 = int32(0)
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v178 = int64(0)
	goto L10
L31:
	;
	goto L11
L32:
	;
	if v160 == int32(68) {
		v185 = v149
		goto L9
	} else {
		goto L33
	}
L33:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v165 == int32(0) {
		v185 = v149
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	v185 = base.B2i32(v169 == int32(0))
	goto L9
L35:
	;
	F_addReplyError(m, l0, int32(_a487))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	return int32(0)
L37:
	;
	v203 = int32(-1)
	goto L7
}
func F_pat_next(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v326
L2:
	;
	v19 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v22 + int32(-91) {
	case 0:
		goto L14
	case 1:
		goto L15
	default:
		goto L16
	}
L3:
	;
	v16 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v16
	v326 = v16
	goto L1
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v15 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v315
	v326 = v317
	goto L1
L7:
	;
	v315 = int32(1)
	v317 = int32(91)
	goto L6
L8:
	;
	if v289 == l1 {
		goto L7
	} else {
		goto L94
	}
L9:
	;
	v289 = v68
	goto L8
L10:
	;
	v326 = int32(-4)
	goto L1
L11:
	;
	v326 = v280 & int32(255)
	goto L1
L12:
	;
	if int32(-1) < base.I32_extend8_s(v154) {
		v280 = v154
		goto L11
	} else {
		goto L60
	}
L13:
	;
	if v22 == int32(63) {
		goto L10
	} else {
		goto L59
	}
L14:
	;
	if l1 == int32(1) {
		v54 = v19
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v28 = int32(92)
	if l3&int32(2) != 0 {
		v280 = v28
		goto L11
	} else {
		goto L18
	}
L16:
	;
	if v22 != int32(42) {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v326 = int32(-5)
	goto L1
L18:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v31&int32(255) == int32(0) {
		v280 = v28
		goto L11
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(2)
	v38 = int32(1)
	v39 = l0 + v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	v152 = v39
	v153 = v38
	v154 = v40
	goto L12
L20:
	;
	if base.Ui32(l1) <= base.Ui32(v54) {
		v61 = v54
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v44 = int32(2)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v47 == int32(33) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v50 = v44
	goto L24
L23:
	;
	v50 = int32(1)
	goto L24
L24:
	;
	if v47 == int32(94) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v53 = v44
	goto L27
L26:
	;
	v53 = v50
	goto L27
L27:
	;
	v54 = v53
	goto L20
L28:
	;
	if base.Ui32(l1) <= base.Ui32(v61) {
		v289 = v61
		goto L8
	} else {
		goto L30
	}
L29:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v54))))
	v61 = v54 + base.B2i32(v57 == int32(93))
	goto L28
L30:
	;
	v68 = v61
	goto L31
L31:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v68))))
	if v72 == int32(0) {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	if v72 == int32(93) {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v78 = v68 + int32(1)
	if base.Ui32(l1) <= base.Ui32(v78) {
		v145 = v78
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if base.Ui32(v145) < base.Ui32(l1) {
		v68 = v145
		goto L31
	} else {
		goto L58
	}
L36:
	;
	if v72 != int32(91) {
		v145 = v78
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v78))))
	if v83 == int32(0) {
		v145 = v78
		goto L35
	} else {
		goto L38
	}
L38:
	;
	switch v83 + int32(-58) {
	case 0, 3:
		goto L39
	case 1, 2:
		v145 = v78
		goto L35
	default:
		goto L40
	}
L39:
	;
	v91 = v68 + int32(2)
	if base.Ui32(l1) <= base.Ui32(v91) {
		v98 = v91
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if v83 != int32(46) {
		v145 = v78
		goto L35
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	if base.Ui32(l1) <= base.Ui32(v98) {
		v130 = v98
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v91))))
	if v96 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v97 = v68 + int32(3)
	goto L46
L45:
	;
	v97 = v91
	goto L46
L46:
	;
	v98 = v97
	goto L42
L47:
	;
	if v130 == l1 {
		goto L7
	} else {
		goto L56
	}
L48:
	;
	v105 = v98
	goto L49
L49:
	;
	v108 = l0 + v105
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v109 == int32(0) {
		v130 = v105
		goto L47
	} else {
		goto L51
	}
L50:
	;
	v315 = v120
	v317 = int32(91)
	goto L6
L51:
	;
	if v109 != int32(93) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v120 = int32(1)
	v122 = v105 + v120
	if v122 != l1 {
		v105 = v122
		goto L49
	} else {
		goto L55
	}
L53:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+int32(-1)))))
	if v116&int32(255) == v83 {
		v130 = v105
		goto L47
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	goto L50
L56:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v130))))
	if v135 == int32(0) {
		v289 = v130
		goto L8
	} else {
		goto L57
	}
L57:
	;
	v145 = v130 + int32(1)
	goto L35
L58:
	;
	v289 = v145
	goto L8
L59:
	;
	v152 = l0
	v153 = int32(0)
	v154 = v22
	goto L12
L60:
	;
	v159 = v11 + int32(12)
	if v152 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v270 = int32(0)
	v273 = base.B2i32(v268 < v270)
	if v268 < v270 {
		goto L88
	} else {
		goto L89
	}
L62:
	;
	if l1 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v268 = int32(0)
	goto L61
L64:
	;
	v268 = v260
	goto L61
L65:
	;
	v256 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = int32(25)
	v260 = int32(-1)
	goto L64
L66:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	v166 = base.I32_extend8_s(v165)
	if v166 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v174 = F___get_tp(m)
	mBase = m.M
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+96))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	if v176 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	if v159 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v268 = base.B2i32(v166 != int32(0))
	goto L61
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v165
	goto L69
L71:
	;
	v185 = v165 + int32(-194)
	if base.Ui32(int32(50)) < base.Ui32(v185) {
		goto L65
	} else {
		goto L74
	}
L72:
	;
	if v159 == int32(0) {
		v260 = int32(1)
		goto L64
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v166 & int32(57343)
	v268 = int32(1)
	goto L61
L74:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v185<<(uint(int32(2))%32))+uint32(_consts[1236])))
	if base.Ui32(int32(3)) < base.Ui32(l1) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+1)))
	v204 = int32(base.Ui32(v202) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(7)) < base.Ui32(v204+int32(-16)|(v204+v192>>(uint(int32(26))%32))) {
		goto L65
	} else {
		goto L78
	}
L76:
	;
	if v192<<(uint(l1*int32(6)+int32(-6))%32) < int32(0) {
		goto L65
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v217 = v202 + int32(-128) | v192<<(uint(int32(6))%32)
	if v217 < int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+2)))
	v227 = v225 + int32(-128)
	if base.Ui32(int32(63)) < base.Ui32(v227) {
		goto L65
	} else {
		goto L82
	}
L80:
	;
	if v159 == int32(0) {
		v260 = int32(2)
		goto L64
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v217
	v268 = int32(2)
	goto L61
L82:
	;
	v231 = v217 << (uint(int32(6)) % 32)
	v232 = v227 | v231
	if v231 < int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+3)))
	v242 = v240 + int32(-128)
	if base.Ui32(int32(63)) < base.Ui32(v242) {
		goto L65
	} else {
		goto L86
	}
L84:
	;
	if v159 == int32(0) {
		v260 = int32(3)
		goto L64
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v232
	v268 = int32(3)
	goto L61
L86:
	;
	if v159 == int32(0) {
		v260 = int32(4)
		goto L64
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v242 | v232<<(uint(int32(6))%32)
	v268 = int32(4)
	goto L61
L88:
	;
	v274 = v270
	goto L90
L89:
	;
	v274 = v268 + v153
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v274
	if v268 < v270 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v277 = int32(-2)
	goto L93
L92:
	;
	v277 = v269
	goto L93
L93:
	;
	v326 = v277
	goto L1
L94:
	;
	v293 = int32(1)
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v289))))
	if v297 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v298 = v289 + v293
	goto L97
L96:
	;
	v298 = v293
	goto L97
L97:
	;
	if v297 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v301 = int32(-3)
	goto L100
L99:
	;
	v301 = int32(91)
	goto L100
L100:
	;
	v315 = v298
	v317 = v301
	goto L6
}
func F_performEvictions(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
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
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int64
	_ = v68
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v89 int64
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
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
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v190 float64
	_ = v190
	var v192 float64
	_ = v192
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int64
	_ = v205
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int64
	_ = v222
	var v231 int32
	_ = v231
	var v234 int64
	_ = v234
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int64
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int64
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v362 int32
	_ = v362
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v394 int32
	_ = v394
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
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
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v680 int64
	_ = v680
	var v682 int32
	_ = v682
	var v686 int64
	_ = v686
	var v690 int64
	_ = v690
	var v693 int32
	_ = v693
	var v701 int64
	_ = v701
	var v705 int32
	_ = v705
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v799 int64
	_ = v799
	var v802 int64
	_ = v802
	var v803 int64
	_ = v803
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int64
	_ = v810
	var v813 int64
	_ = v813
	var v815 int64
	_ = v815
	var v818 int64
	_ = v818
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v921 int32
	_ = v921
	var v923 int64
	_ = v923
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v969 int32
	_ = v969
	var v974 int64
	_ = v974
	var v983 int64
	_ = v983
	var v987 int32
	_ = v987
	var v988 int64
	_ = v988
	var v990 int32
	_ = v990
	var v995 int64
	_ = v995
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1027 int64
	_ = v1027
	var v1049 int32
	_ = v1049
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1061 int64
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1075 int64
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1080 int64
	_ = v1080
	var v1081 int64
	_ = v1081
	var v1108 int32
	_ = v1108
	var v1113 int32
	_ = v1113
	var v1119 int32
	_ = v1119
	var v1125 int32
	_ = v1125
	var v1152 int64
	_ = v1152
	var v1158 int64
	_ = v1158
	var v1159 int64
	_ = v1159
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1196 int32
	_ = v1196
	var v1197 int64
	_ = v1197
	var v1210 int32
	_ = v1210
	var v1215 int64
	_ = v1215
	var v1224 int64
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1229 int64
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1236 int64
	_ = v1236
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1268 int64
	_ = v1268
	var v1290 int32
	_ = v1290
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1334 int64
	_ = v1334
	var v1337 int64
	_ = v1337
	var v1339 int64
	_ = v1339
	var v1342 int64
	_ = v1342
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1374 int64
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1383 int64
	_ = v1383
	var v1401 int64
	_ = v1401
	var v1403 int64
	_ = v1403
	var v1406 int64
	_ = v1406
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1441 int32
	_ = v1441
	var v1444 int32
	_ = v1444
	var v1466 int64
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1471 int64
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1497 int64
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1504 int64
	_ = v1504
	var v1507 int64
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1519 int64
	_ = v1519
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	v1 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(16)
	m.G0 = v27
	v30 = F_scriptIsTimedout(m)
	mBase = m.M
	v32 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	goto L2
L1:
	;
	m.G0 = v1537 + int32(16)
	return v1538
L2:
	;
	if v30|v32 != v1 {
		v1537 = v27
		v1538 = v1
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	if v37 != 0 {
		v1537 = v27
		v1538 = v1
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	if v39 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v45 = F_isPausedActionsWithUpdate(m, int32(8))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[263]))
	if v43 != 0 {
		v1537 = v27
		v1538 = v1
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	return int32(0)
L9:
	;
	if v45 != 0 {
		v1537 = v27
		v1538 = v1
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	v52 = int32(0)
	v54 = v27 + int32(12)
	v57 = v27 + int32(8)
	v63 = F_zmalloc_used_memory(m)
	mBase = m.M
	if v54 == v52 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, _consts[264])) = v1519
	v1537 = v1511
	v1538 = v1512
	goto L1
L12:
	;
	v1497 = *(*int64)(unsafe.Add(mBase, _consts[264]))
	if v1497 == int64(0) {
		v1537 = v1472
		v1538 = v1473
		goto L1
	} else {
		goto L295
	}
L13:
	;
	if v150 == int32(0) {
		v1472 = v27
		v1473 = v52
		goto L12
	} else {
		goto L42
	}
L14:
	;
	v68 = *(*int64)(unsafe.Add(mBase, _consts[265]))
	if v68 != int64(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v63
	goto L14
L16:
	;
	v150 = v143
	goto L13
L17:
	;
	v77 = base.I64_extend_i32_u(v63)
	goto L21
L18:
	;
	v143 = int32(0)
	goto L16
L20:
	;
	v81 = int32(_a69)
	v82 = *(*int64)(unsafe.Add(mBase, _consts[266]))
	v84 = *(*int32)(unsafe.Add(mBase, _consts[267]))
	if base.I64_extend_i32_u(v84) <= v82 {
		v99 = int32(0)
		goto L23
	} else {
		goto L24
	}
L21:
	;
	if base.Ui64(v68) < base.Ui64(v77) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v150 = int32(0)
	goto L13
L23:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v101 == int32(0) {
		v108 = v99
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v89 = base.I64_div_s(v82, int64(16384))
	v96 = v84 - base.I32_wrap_i64(v82+v89*int64(44)) + int32(-44)
	if base.Ui32(v84) < base.Ui32(v96) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v98 = int32(0)
	goto L27
L26:
	;
	v98 = v96
	goto L27
L27:
	;
	v99 = v98
	goto L23
L28:
	;
	v109 = F_clusterIsAnySlotExporting(m)
	mBase = m.M
	if v109 == int32(0) {
		v114 = v108
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[268]))
	v106 = F_sdsAllocSize(m, v105)
	mBase = m.M
	v108 = v106 + v99
	goto L28
L30:
	;
	v115 = int32(0)
	v117 = v63 - v114
	if base.Ui32(v63) < base.Ui32(v117) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v112 = F_clusterGetTotalSlotExportBufferMemory(m)
	mBase = m.M
	v114 = v112 + v108
	goto L30
L32:
	;
	v119 = v115
	goto L34
L33:
	;
	v119 = v117
	goto L34
L34:
	;
	v121 = *(*int64)(unsafe.Add(mBase, _consts[265]))
	goto L35
L35:
	;
	if base.Ui64(v77) <= base.Ui64(v121) {
		v143 = v115
		goto L16
	} else {
		goto L37
	}
L37:
	;
	if base.Ui64(base.I64_extend_i32_u(v119)) <= base.Ui64(v121) {
		v143 = v115
		goto L16
	} else {
		goto L38
	}
L38:
	;
	goto L39
L39:
	;
	v134 = int32(-1)
	if v57 == int32(0) {
		v143 = v134
		goto L16
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v119 - base.I32_wrap_i64(v121)
	v143 = v134
	goto L16
L42:
	;
	v153 = int32(2)
	v155 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	if v155 == int32(1792) {
		v1441 = v27
		v1444 = v153
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v1466 = *(*int64)(unsafe.Add(mBase, _consts[264]))
	if v1466 == int64(0) {
		goto L293
	} else {
		goto L294
	}
L44:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v159 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _consts[269]))
	if v174 <= int32(-1) {
		goto L57
	} else {
		goto L58
	}
L46:
	;
	if v168 == int32(0) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	v164 = F_getMyClusterNode(m)
	mBase = m.M
	v165 = F_clusterNodeIsPrimary(m, v164)
	mBase = m.M
	v168 = base.B2i32(v165 != int32(0))
	goto L46
L48:
	;
	v160 = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	v168 = base.B2i32(v161 == v160)
	goto L46
L49:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _consts[218]))
	if v172 != 0 {
		v1441 = v27
		v1444 = v153
		goto L43
	} else {
		goto L50
	}
L50:
	;
	goto L45
L51:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1416+int32(-1)) {
		v1472 = v1413
		v1473 = v1414
		goto L12
	} else {
		goto L292
	}
L52:
	;
	if v1383 == int64(0) {
		v1413 = v1375
		v1414 = v1376
		v1416 = v1378
		goto L51
	} else {
		goto L288
	}
L53:
	;
	v1374 = *(*int64)(unsafe.Add(mBase, _consts[270]))
	v1375 = v27
	v1376 = v52
	v1378 = v1352
	v1383 = v1374
	goto L52
L54:
	;
	v1152 = *(*int64)(unsafe.Add(mBase, _consts[270]))
	if base.B2i32(v1152 == int64(0)) == int32(0) {
		goto L238
	} else {
		goto L239
	}
L55:
	;
	F__serverAssert(m, int32(_a565), int32(_a566), int32(435))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L8
	} else {
		goto L236
	}
L56:
	;
	F__serverAssert(m, int32(_a567), int32(_a566), int32(365))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L8
	} else {
		goto L235
	}
L57:
	;
	F__serverAssert(m, int32(_a568), int32(_a566), int32(364))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L8
	} else {
		goto L234
	}
L58:
	;
	if base.Ui32(int32(101)) <= base.Ui32(v174) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	if base.Ui32(int32(10)) < base.Ui32(v174) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v205 = *(*int64)(unsafe.Add(mBase, _consts[270]))
	if base.B2i32(v205 == int64(0)) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L61:
	;
	if v174 == int32(100) {
		v202 = int32(-1)
		goto L60
	} else {
		goto L63
	}
L62:
	;
	v202 = v174 * int32(50)
	goto L60
L63:
	;
	v190 = F_pow(m, float64(1.15), base.F64_add(base.F64_convert_i32_u(v174), float64(-10)))
	mBase = m.M
	v192 = base.F64_mul(v190, float64(500))
	if base.F64_lt(v192, float64(4.294967296e+09))&base.F64_ge(v192, float64(0)) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v202 = int32(0)
	goto L60
L65:
	;
	v200 = base.I32_trunc_f64_u(v192)
	v202 = v200
	goto L60
L66:
	;
	v213 = int32(0)
	v215 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	v216 = m.T0[v215].(func(*base.Module) int64)(m)
	mBase = m.M
	v218 = *(*int32)(unsafe.Add(mBase, _consts[272]))
	if v218 != 0 {
		goto L55
	} else {
		goto L69
	}
L67:
	;
	v211 = F_ustime(m)
	mBase = m.M
	v212 = v211
	goto L66
L68:
	;
	v212 = int64(0)
	goto L66
L69:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v219 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, _consts[262])))
	v1352 = v1108
	goto L53
L71:
	;
	v222 = base.I64_extend_i32_u(v202)
	v231 = v213
	v234 = int64(0)
	goto L72
L72:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	if v249 == int32(512) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	goto L70
L74:
	;
	v637 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629+int32(-1)))))
	switch v644 & int32(7) {
	case 0:
		goto L139
	case 1:
		goto L138
	case 2:
		goto L137
	case 3:
		goto L136
	case 4:
		goto L135
	default:
		v661 = int32(0)
		goto L134
	}
L75:
	;
	if v249 == int32(1540) {
		goto L116
	} else {
		goto L117
	}
L76:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _consts[261]))
	goto L79
L77:
	;
	if v249&int32(3) == int32(0) {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v282 = int32(0)
	v285 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v285 < int32(1) {
		goto L54
	} else {
		goto L81
	}
L81:
	;
	v300 = v282
	v301 = v282
	goto L82
L82:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v313+v301<<(uint(int32(2))%32))))
	if v317 == int32(0) {
		v394 = v300
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v394 == int32(0) {
		goto L54
	} else {
		goto L100
	}
L84:
	;
	v407 = v301 + int32(1)
	v409 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v407 < v409 {
		v300 = v394
		v301 = v407
		goto L82
	} else {
		goto L99
	}
L85:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v317+(v321^int32(-1))&int32(4))))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+12))
	if v328 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v339 = base.I32_wrap_i64(v338)
	if v339 == int32(0) {
		v394 = v300
		goto L84
	} else {
		goto L91
	}
L87:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v327)+8))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	if v333 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v327)+40))
	v338 = v331
	goto L86
L89:
	;
	v335 = F_hashtableSize(m, v333)
	mBase = m.M
	v338 = base.I64_extend_i32_u(v335)
	goto L86
L90:
	;
	v338 = int64(0)
	goto L86
L91:
	;
	v342 = v300 + v339
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v327)+32))
	goto L92
L92:
	;
	v348 = v344
	v362 = int32(0)
	goto L93
L93:
	;
	if v348 == int32(0) {
		v394 = v342
		goto L84
	} else {
		goto L95
	}
L94:
	;
	v394 = v342
	goto L84
L95:
	;
	v371 = F_evictionPoolPopulate(m, v317, v327, v257)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	v373 = v371 + v362
	v375 = *(*int32)(unsafe.Add(mBase, _consts[273]))
	if base.Ui32(v375) <= base.Ui32(v373) {
		v394 = v342
		goto L84
	} else {
		goto L97
	}
L97:
	;
	if base.Ui32(v375*int32(10)) <= base.Ui32(v339) {
		v348 = v348 + int32(-1)
		v362 = v373
		goto L93
	} else {
		goto L98
	}
L98:
	;
	goto L94
L99:
	;
	goto L83
L100:
	;
	v417 = int32(15)
	goto L101
L101:
	;
	v440 = v257 + v417*int32(24)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)+8))
	if v441 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L79
L103:
	;
	if v417 != 0 {
		v417 = v417 + int32(-1)
		goto L101
	} else {
		goto L114
	}
L104:
	;
	v444 = int32(_a69)
	v445 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v440)+16))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v445+v446<<(uint(int32(2))%32))))
	v452 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	v455 = int32(4)
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v450+(v452^int32(-1))&v455)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(0)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v440)+20))
	v464 = F_kvstoreHashtableFind(m, v458, v461, v441, v27+v455)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v440)+8))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v440)+12))
	if v466 == v467 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v440))) = int64(0)
	v473 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v440)+8)) = v473
	if v464 == v473 {
		goto L103
	} else {
		goto L109
	}
L107:
	;
	F_sdsfree(m, v466)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v478 = int32(0)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v477)+4))
	if v481&int32(2) == v478 {
		v501 = v478
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v440)+20))
	if v501 == int32(0) {
		goto L79
	} else {
		goto L113
	}
L111:
	;
	goto L110
L112:
	;
	v495 = v477 + (v481&int32(4) ^ int32(12)) + v481<<(uint(int32(3))%32)&int32(8)
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495))))
	v501 = v495 + v496 + int32(1)
	goto L111
L113:
	;
	v629 = v501
	v631 = v446
	v632 = v502
	goto L74
L114:
	;
	goto L102
L116:
	;
	v520 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v520 < int32(1) {
		goto L54
	} else {
		goto L119
	}
L117:
	;
	if v249 != int32(768) {
		goto L54
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v526 = v520
	v540 = int32(0)
	goto L121
L120:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v586 = int32(0)
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v585)+4))
	if v589&int32(2) == v586 {
		v609 = v586
		goto L131
	} else {
		goto L132
	}
L121:
	;
	v547 = int32(0)
	v549 = *(*int32)(unsafe.Add(mBase, _consts[274]))
	v551 = v549 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[274])) = v551
	v554 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v555 = base.I32_rem_u_s(v551, v526)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v554+v555<<(uint(int32(2))%32))))
	if v559 == v547 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v581 = v540 + int32(1)
	v583 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v581 < v583 {
		v526 = v583
		v540 = v581
		goto L121
	} else {
		goto L129
	}
L124:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v559+base.B2i32(v563 != int32(1540))<<(uint(int32(2))%32))))
	v570 = F_kvstoreGetFairRandomHashtableIndex(m, v569)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L8
	} else {
		goto L125
	}
L125:
	;
	if v570 == int32(-1) {
		goto L123
	} else {
		goto L126
	}
L126:
	;
	v576 = F_kvstoreHashtableRandomEntry(m, v569, v570, v27+int32(4))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L8
	} else {
		goto L127
	}
L127:
	;
	if v576 != 0 {
		goto L120
	} else {
		goto L128
	}
L128:
	;
	goto L123
L129:
	;
	goto L54
L130:
	;
	if v609 == int32(0) {
		goto L54
	} else {
		goto L133
	}
L131:
	;
	goto L130
L132:
	;
	v603 = v585 + (v589&int32(4) ^ int32(12)) + v589<<(uint(int32(3))%32)&int32(8)
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603))))
	v609 = v603 + v604 + int32(1)
	goto L131
L133:
	;
	v629 = v609
	v631 = v555
	v632 = v570
	goto L74
L134:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v637+v631<<(uint(int32(2))%32))))
	v664 = F_createStringObject_1(m, v629, v661)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L8
	} else {
		goto L140
	}
L135:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v629+int32(-17))))
	v661 = v660
	goto L134
L136:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v629+int32(-9))))
	v661 = v657
	goto L134
L137:
	;
	v654 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v629+int32(-5)))))
	v661 = v654
	goto L134
L138:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629+int32(-3)))))
	v661 = v651
	goto L134
L139:
	;
	v661 = int32(base.Ui32(v644) >> (uint(int32(3)) % 32))
	goto L134
L140:
	;
	v668 = int32(0)
	v672 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	*(*int32)(unsafe.Add(mBase, _consts[275])) = v672 + int32(1)
	goto L143
L141:
	;
	v705 = int32(0)
	v714 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	if v714 < int32(261) {
		goto L150
	} else {
		goto L151
	}
L142:
	;
	goto L141
L143:
	;
	if v672 != 0 {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	goto L146
L145:
	;
	v682 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v680
	v686 = base.I64_div_s(v680, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[32])) = v686
	v690 = base.I64_div_s(v680, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[37])) = v690
	v693 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	F_lrulfu_updateClockAndPolicy(m, v686, int32(base.Ui32(v693&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v701 = *(*int64)(unsafe.Add(mBase, _consts[32]))
	*(*int64)(unsafe.Add(mBase, _consts[78])) = v701
	goto L142
L146:
	;
	v680 = F_ustime(m)
	mBase = m.M
	goto L145
L147:
	;
	v799 = *(*int64)(unsafe.Add(mBase, _consts[270]))
	if v799 == int64(0) {
		v803 = int64(0)
		goto L163
	} else {
		goto L164
	}
L148:
	;
	goto L147
L149:
	;
	v725 = v723 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v723) {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	if v714 < int32(1) {
		v791 = v705
		goto L148
	} else {
		goto L152
	}
L151:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v722 = v718
	v723 = int32(260)
	goto L149
L152:
	;
	v722 = v705
	v723 = v714
	goto L149
L153:
	;
	if v725 == int32(0) {
		v791 = v764
		goto L148
	} else {
		goto L159
	}
L154:
	;
	v732 = int32(0)
	v734 = v722
	v735 = v732
	v739 = v732
	goto L156
L155:
	;
	v764 = v722
	v765 = int32(0)
	goto L153
L156:
	;
	v742 = v735 << (uint(int32(2)) % 32)
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v742)+uint32(_consts[279])))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v742)+uint32(_consts[280])))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v742)+uint32(_consts[281])))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v742)+uint32(_consts[282])))
	v758 = v745 + (v748 + (v751 + (v754 + v734)))
	v759 = int32(4)
	v760 = v735 + v759
	v762 = v739 + v759
	if v762 != v723&int32(2147483644) {
		v734 = v758
		v735 = v760
		v739 = v762
		goto L156
	} else {
		goto L158
	}
L157:
	;
	v764 = v758
	v765 = v760
	goto L153
L158:
	;
	goto L157
L159:
	;
	v773 = v764
	v774 = v765
	v776 = int32(0)
	goto L160
L160:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v774<<(uint(int32(2))%32))+uint32(_consts[282])))
	v785 = v784 + v773
	v786 = int32(1)
	v789 = v776 + v786
	if v789 != v725 {
		v773 = v785
		v774 = v774 + v786
		v776 = v789
		goto L160
	} else {
		goto L162
	}
L161:
	;
	v791 = v785
	goto L148
L162:
	;
	goto L161
L163:
	;
	v805 = *(*int32)(unsafe.Add(mBase, _consts[283]))
	v807 = F_dbGenericDelete(m, v662, v664, v805, int32(4))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L8
	} else {
		goto L165
	}
L164:
	;
	v802 = F_ustime(m)
	mBase = m.M
	v803 = v802
	goto L163
L165:
	;
	v810 = *(*int64)(unsafe.Add(mBase, _consts[270]))
	if v810 == int64(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v828 = int32(0)
	v837 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	if v837 < int32(261) {
		goto L174
	} else {
		goto L175
	}
L167:
	;
	v813 = F_ustime(m)
	mBase = m.M
	v815 = *(*int64)(unsafe.Add(mBase, _consts[270]))
	if v815 == int64(0) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v818 = v813 - v803
	if v818 < v815*int64(1000) {
		goto L166
	} else {
		goto L169
	}
L169:
	;
	F_latencyAddSample(m, int32(_a569), v818)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L8
	} else {
		goto L170
	}
L170:
	;
	goto L166
L171:
	;
	v921 = int32(_a69)
	v923 = *(*int64)(unsafe.Add(mBase, _consts[284]))
	*(*int64)(unsafe.Add(mBase, _consts[284])) = v923 + int64(1)
	F_signalModifiedKey(m, int32(0), v662, v664)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L8
	} else {
		goto L187
	}
L172:
	;
	goto L171
L173:
	;
	v848 = v846 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v846) {
		goto L178
	} else {
		goto L179
	}
L174:
	;
	if v837 < int32(1) {
		v914 = v828
		goto L172
	} else {
		goto L176
	}
L175:
	;
	v841 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v845 = v841
	v846 = int32(260)
	goto L173
L176:
	;
	v845 = v828
	v846 = v837
	goto L173
L177:
	;
	if v848 == int32(0) {
		v914 = v887
		goto L172
	} else {
		goto L183
	}
L178:
	;
	v855 = int32(0)
	v857 = v845
	v858 = v855
	v862 = v855
	goto L180
L179:
	;
	v887 = v845
	v888 = int32(0)
	goto L177
L180:
	;
	v865 = v858 << (uint(int32(2)) % 32)
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v865)+uint32(_consts[279])))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v865)+uint32(_consts[280])))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v865)+uint32(_consts[281])))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v865)+uint32(_consts[282])))
	v881 = v868 + (v871 + (v874 + (v877 + v857)))
	v882 = int32(4)
	v883 = v858 + v882
	v885 = v862 + v882
	if v885 != v846&int32(2147483644) {
		v857 = v881
		v858 = v883
		v862 = v885
		goto L180
	} else {
		goto L182
	}
L181:
	;
	v887 = v881
	v888 = v883
	goto L177
L182:
	;
	goto L181
L183:
	;
	v896 = v887
	v897 = v888
	v899 = int32(0)
	goto L184
L184:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v897<<(uint(int32(2))%32))+uint32(_consts[282])))
	v908 = v907 + v896
	v909 = int32(1)
	v912 = v899 + v909
	if v912 != v848 {
		v896 = v908
		v897 = v897 + v909
		v899 = v912
		goto L184
	} else {
		goto L186
	}
L185:
	;
	v914 = v908
	goto L172
L186:
	;
	goto L185
L187:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v662)+28))
	F_notifyKeyspaceEvent(m, int32(512), int32(_a570), v664, v932)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L8
	} else {
		goto L188
	}
L188:
	;
	v936 = *(*int32)(unsafe.Add(mBase, _consts[283]))
	F_propagateDeletion(m, v662, v664, v936, v632)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L8
	} else {
		goto L189
	}
L189:
	;
	v939 = int32(0)
	v941 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	*(*int32)(unsafe.Add(mBase, _consts[275])) = v941 + int32(-1)
	goto L190
L190:
	;
	F_postExecutionUnitOperations(m)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L8
	} else {
		goto L191
	}
L191:
	;
	F_decrRefCount(m, v664)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L8
	} else {
		goto L192
	}
L192:
	;
	v950 = v231 + int32(1)
	if v950&int32(15) != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v1080 = base.I64_extend_i32_u(v791) - base.I64_extend_i32_u(v914) + v234
	v1081 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v27)+8)))
	if v1080 < v1081 {
		v231 = v950
		v234 = v1080
		goto L72
	} else {
		goto L233
	}
L194:
	;
	if v51 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v958 = *(*int32)(unsafe.Add(mBase, _consts[283]))
	if v958 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	F_flushReplicasOutputBuffers(m)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L8
	} else {
		goto L197
	}
L197:
	;
	goto L195
L198:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	v1061 = m.T0[v1060].(func(*base.Module) int64)(m)
	mBase = m.M
	if base.Ui64(v1061-v216) <= base.Ui64(v222) {
		goto L193
	} else {
		goto L230
	}
L199:
	;
	v969 = F_zmalloc_used_memory(m)
	mBase = m.M
	goto L201
L200:
	;
	if v1056 == int32(0) {
		goto L70
	} else {
		goto L229
	}
L201:
	;
	v974 = *(*int64)(unsafe.Add(mBase, _consts[265]))
	if v974 != int64(0) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	v1056 = v1049
	goto L200
L204:
	;
	v983 = base.I64_extend_i32_u(v969)
	goto L208
L205:
	;
	v1049 = int32(0)
	goto L203
L207:
	;
	v987 = int32(_a69)
	v988 = *(*int64)(unsafe.Add(mBase, _consts[266]))
	v990 = *(*int32)(unsafe.Add(mBase, _consts[267]))
	if base.I64_extend_i32_u(v990) <= v988 {
		v1005 = int32(0)
		goto L210
	} else {
		goto L211
	}
L208:
	;
	if base.Ui64(v974) < base.Ui64(v983) {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v1056 = int32(0)
	goto L200
L210:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v1007 == int32(0) {
		v1014 = v1005
		goto L215
	} else {
		goto L216
	}
L211:
	;
	v995 = base.I64_div_s(v988, int64(16384))
	v1002 = v990 - base.I32_wrap_i64(v988+v995*int64(44)) + int32(-44)
	if base.Ui32(v990) < base.Ui32(v1002) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1004 = int32(0)
	goto L214
L213:
	;
	v1004 = v1002
	goto L214
L214:
	;
	v1005 = v1004
	goto L210
L215:
	;
	v1015 = F_clusterIsAnySlotExporting(m)
	mBase = m.M
	if v1015 == int32(0) {
		v1020 = v1014
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, _consts[268]))
	v1012 = F_sdsAllocSize(m, v1011)
	mBase = m.M
	v1014 = v1012 + v1005
	goto L215
L217:
	;
	v1021 = int32(0)
	v1023 = v969 - v1020
	if base.Ui32(v969) < base.Ui32(v1023) {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v1018 = F_clusterGetTotalSlotExportBufferMemory(m)
	mBase = m.M
	v1020 = v1018 + v1014
	goto L217
L219:
	;
	v1025 = v1021
	goto L221
L220:
	;
	v1025 = v1023
	goto L221
L221:
	;
	v1027 = *(*int64)(unsafe.Add(mBase, _consts[265]))
	goto L222
L222:
	;
	if base.Ui64(v983) <= base.Ui64(v1027) {
		v1049 = v1021
		goto L203
	} else {
		goto L224
	}
L224:
	;
	if base.Ui64(base.I64_extend_i32_u(v1025)) <= base.Ui64(v1027) {
		v1049 = v1021
		goto L203
	} else {
		goto L225
	}
L225:
	;
	goto L226
L226:
	;
	v1049 = int32(-1)
	goto L203
L229:
	;
	goto L198
L230:
	;
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, _consts[262])))
	if v1065 != 0 {
		goto L70
	} else {
		goto L231
	}
L231:
	;
	v1066 = int32(0)
	v1067 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[262])) = uint8(v1067)
	v1070 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v1075 = F_aeCreateTimeEvent(m, v1070, int64(0), int32(524), v1066, v1066)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L8
	} else {
		goto L232
	}
L232:
	;
	goto L70
L233:
	;
	goto L73
L234:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, _consts[285]))
	goto L242
L238:
	;
	v1158 = F_ustime(m)
	mBase = m.M
	v1159 = v1158
	goto L237
L239:
	;
	v1159 = int64(0)
	goto L237
L240:
	;
	v1334 = *(*int64)(unsafe.Add(mBase, _consts[270]))
	if v1334 == int64(0) {
		v1413 = v27
		v1414 = v52
		v1416 = v1312
		goto L51
	} else {
		goto L284
	}
L241:
	;
	v1167 = int32(1000)
	if base.Ui32(v202) < base.Ui32(v1167) {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	if v1165 != 0 {
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v1312 = int32(2)
	goto L240
L244:
	;
	v1170 = v202
	goto L246
L245:
	;
	v1170 = v1167
	goto L246
L246:
	;
	goto L247
L247:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	v1197 = m.T0[v1196].(func(*base.Module) int64)(m)
	mBase = m.M
	if base.Ui64(v1197-v216) < base.Ui64(v222) {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v1312 = int32(2)
	goto L240
L249:
	;
	v1210 = F_zmalloc_used_memory(m)
	mBase = m.M
	goto L252
L250:
	;
	v1312 = int32(2)
	goto L240
L251:
	;
	if v1297 == int32(0) {
		v1312 = int32(0)
		goto L240
	} else {
		goto L280
	}
L252:
	;
	v1215 = *(*int64)(unsafe.Add(mBase, _consts[265]))
	if v1215 != int64(0) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1297 = v1290
	goto L251
L255:
	;
	v1224 = base.I64_extend_i32_u(v1210)
	goto L259
L256:
	;
	v1290 = int32(0)
	goto L254
L258:
	;
	v1228 = int32(_a69)
	v1229 = *(*int64)(unsafe.Add(mBase, _consts[266]))
	v1231 = *(*int32)(unsafe.Add(mBase, _consts[267]))
	if base.I64_extend_i32_u(v1231) <= v1229 {
		v1246 = int32(0)
		goto L261
	} else {
		goto L262
	}
L259:
	;
	if base.Ui64(v1215) < base.Ui64(v1224) {
		goto L258
	} else {
		goto L260
	}
L260:
	;
	v1297 = int32(0)
	goto L251
L261:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v1248 == int32(0) {
		v1255 = v1246
		goto L266
	} else {
		goto L267
	}
L262:
	;
	v1236 = base.I64_div_s(v1229, int64(16384))
	v1243 = v1231 - base.I32_wrap_i64(v1229+v1236*int64(44)) + int32(-44)
	if base.Ui32(v1231) < base.Ui32(v1243) {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1245 = int32(0)
	goto L265
L264:
	;
	v1245 = v1243
	goto L265
L265:
	;
	v1246 = v1245
	goto L261
L266:
	;
	v1256 = F_clusterIsAnySlotExporting(m)
	mBase = m.M
	if v1256 == int32(0) {
		v1261 = v1255
		goto L268
	} else {
		goto L269
	}
L267:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, _consts[268]))
	v1253 = F_sdsAllocSize(m, v1252)
	mBase = m.M
	v1255 = v1253 + v1246
	goto L266
L268:
	;
	v1262 = int32(0)
	v1264 = v1210 - v1261
	if base.Ui32(v1210) < base.Ui32(v1264) {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	v1259 = F_clusterGetTotalSlotExportBufferMemory(m)
	mBase = m.M
	v1261 = v1259 + v1255
	goto L268
L270:
	;
	v1266 = v1262
	goto L272
L271:
	;
	v1266 = v1264
	goto L272
L272:
	;
	v1268 = *(*int64)(unsafe.Add(mBase, _consts[265]))
	goto L273
L273:
	;
	if base.Ui64(v1224) <= base.Ui64(v1268) {
		v1290 = v1262
		goto L254
	} else {
		goto L275
	}
L275:
	;
	if base.Ui64(base.I64_extend_i32_u(v1266)) <= base.Ui64(v1268) {
		v1290 = v1262
		goto L254
	} else {
		goto L276
	}
L276:
	;
	goto L277
L277:
	;
	v1290 = int32(-1)
	goto L254
L280:
	;
	v1300 = F_usleep(m, v1170)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L8
	} else {
		goto L281
	}
L281:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, _consts[285]))
	goto L282
L282:
	;
	if v1308 != 0 {
		goto L247
	} else {
		goto L283
	}
L283:
	;
	goto L248
L284:
	;
	v1337 = F_ustime(m)
	mBase = m.M
	v1339 = *(*int64)(unsafe.Add(mBase, _consts[270]))
	if v1339 == int64(0) {
		v1375 = v27
		v1376 = v52
		v1378 = v1312
		v1383 = v1339
		goto L52
	} else {
		goto L285
	}
L285:
	;
	v1342 = v1337 - v1159
	if v1342 < v1339*int64(1000) {
		v1375 = v27
		v1376 = v52
		v1378 = v1312
		v1383 = v1339
		goto L52
	} else {
		goto L286
	}
L286:
	;
	F_latencyAddSample(m, int32(_a571), v1342)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L8
	} else {
		goto L287
	}
L287:
	;
	v1352 = v1312
	goto L53
L288:
	;
	v1401 = F_ustime(m)
	mBase = m.M
	v1403 = *(*int64)(unsafe.Add(mBase, _consts[270]))
	if v1403 == int64(0) {
		v1413 = v1375
		v1414 = v1376
		v1416 = v1378
		goto L51
	} else {
		goto L289
	}
L289:
	;
	v1406 = v1401 - v212
	if v1406 < v1403*int64(1000) {
		v1413 = v1375
		v1414 = v1376
		v1416 = v1378
		goto L51
	} else {
		goto L290
	}
L290:
	;
	F_latencyAddSample(m, int32(_a572), v1406)
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L8
	} else {
		goto L291
	}
L291:
	;
	v1413 = v1375
	v1414 = v1376
	v1416 = v1378
	goto L51
L292:
	;
	v1441 = v1413
	v1444 = v1416
	goto L43
L293:
	;
	v1470 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	v1471 = m.T0[v1470].(func(*base.Module) int64)(m)
	mBase = m.M
	v1511 = v1441
	v1512 = v1444
	v1519 = v1471
	goto L11
L294:
	;
	v1537 = v1441
	v1538 = v1444
	goto L1
L295:
	;
	v1500 = int32(0)
	v1501 = int32(_a69)
	v1503 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	v1504 = m.T0[v1503].(func(*base.Module) int64)(m)
	mBase = m.M
	v1507 = *(*int64)(unsafe.Add(mBase, _consts[286]))
	*(*int64)(unsafe.Add(mBase, _consts[286])) = v1504 - v1497 + v1507
	v1511 = v1472
	v1512 = v1500
	v1519 = int64(0)
	goto L11
}
func F_performInterfaceSet(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v12&int32(8) == int32(0) {
		v47 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v47
		v52 = v9 + int32(12)
		v53 = v47
		v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v56 = m.T0[v55].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v52, v53, l2)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
			if v58&int32(8) == int32(0) {
				m.G0 = v9 + int32(16)
				return v56
			} else {
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				F_sdsfreesplitres(m, v52, v63)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(16)
					return v56
				}
			}
		}
	} else {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
		switch v20 & int32(7) {
		case 0:
			v37 = int32(base.Ui32(v20) >> (uint(int32(3)) % 32))
		case 1:
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
			v37 = v27
		case 2:
			v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
			v37 = v30
		case 3:
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
			v37 = v33
		case 4:
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
			v37 = v36
		default:
			v37 = int32(0)
		}
		v42 = F_sdssplitlen(m, l1, v37, int32(_a10), int32(1), v9+int32(8))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			v52 = v42
			v53 = v46
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v56 = m.T0[v55].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v52, v53, l2)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
				if v58&int32(8) == int32(0) {
					m.G0 = v9 + int32(16)
					return v56
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					F_sdsfreesplitres(m, v52, v63)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v56
					}
				}
			}
		}
	}
}
func F_pfaddCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int64
	_ = v145
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = F_lookupKeyWrite(m, v14, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return
L2:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v54 < int32(3) {
		v116 = v52
		goto L13
	} else {
		goto L14
	}
L3:
	;
	v42 = F_isHLLObjectOrReply(m, l0, v17)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L10
	}
L4:
	;
	return
L5:
	;
	if v17 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v21 = F_sdsnewlen(m, int32(0), int32(18))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v23 = int32(65407)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)) = uint16(v23)
	v27 = F_createObject(m, int32(0), v21)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v29 = F_objectGetVal(m, v27)
	mBase = m.M
	v30 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)) = uint8(v30)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(1280072008)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	F_dbAdd(m, v35, v37, v12+int32(12))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v52 = int32(1)
	goto L2
L10:
	;
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v47 = F_dbUnshareStringValue(m, v44, v46, v17)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v47
	v52 = int32(0)
	goto L2
L13:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v123 = F_objectGetVal(m, v122)
	mBase = m.M
	if v116 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L14:
	;
	v60 = int32(2)
	v61 = v52
	goto L15
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v70 = v60 << (uint(int32(2)) % 32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v70)))
	v73 = F_objectGetVal(m, v72)
	mBase = m.M
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v75+v70)))
	v78 = F_objectGetVal(m, v77)
	mBase = m.M
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+int32(-1)))))
	switch v81 & int32(7) {
	case 0:
		goto L22
	case 1:
		goto L21
	case 2:
		goto L20
	case 3:
		goto L19
	case 4:
		goto L18
	default:
		v98 = int32(0)
		goto L17
	}
L16:
	;
	v116 = v108
	goto L13
L17:
	;
	v99 = F_hllAdd(m, v67, v73, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L26
	}
L18:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v78+int32(-17))))
	v98 = v97
	goto L17
L19:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v78+int32(-9))))
	v98 = v94
	goto L17
L20:
	;
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78+int32(-5)))))
	v98 = v91
	goto L17
L21:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+int32(-3)))))
	v98 = v88
	goto L17
L22:
	;
	v98 = int32(base.Ui32(v81) >> (uint(int32(3)) % 32))
	goto L17
L23:
	;
	v110 = v60 + int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v110 < v111 {
		v60 = v110
		v61 = v108
		goto L15
	} else {
		goto L28
	}
L24:
	;
	v108 = v61 + int32(1)
	goto L23
L25:
	;
	F_addReplyError(m, l0, int32(_a641))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L27
	}
L26:
	;
	switch v99 + int32(1) {
	case 0:
		goto L25
	default:
		v108 = v61
		goto L23
	case 2:
		goto L24
	}
L27:
	;
	goto L1
L28:
	;
	goto L16
L29:
	;
	if v116 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+15)))
	v128 = v126 | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+15)) = uint8(v128)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	F_signalModifiedKey(m, l0, v130, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
	F_notifyKeyspaceEvent(m, int32(8), int32(_a642), v138, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v143 = int32(_a69)
	v145 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v145 + base.I64_extend_i32_s(v116)
	goto L29
L33:
	;
	v152 = int32(16)
	goto L35
L34:
	;
	v152 = int32(12)
	goto L35
L35:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)+uint32(_consts[77])))
	F_addReply(m, l0, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	goto L1
}
func F_pfdebugCommand(m *base.Module, l0 int32) {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
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
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
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
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int64
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
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
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int64
	_ = v640
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = F_objectGetVal(m, v14)
	mBase = m.M
	v16 = int32(_a643)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v11 + int32(80)
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v15
	F_addReplyErrorFormat(m, l0, int32(_a644), v11)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L45
	} else {
		goto L199
	}
L3:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	v155 = F_lookupKeyWrite(m, v152, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L45
	} else {
		goto L49
	}
L4:
	;
	if v51-v53 != 0 {
		goto L3
	} else {
		goto L16
	}
L5:
	;
	v51 = F_tolower(m, v47)
	mBase = m.M
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v53 = F_tolower(m, v52)
	mBase = m.M
	goto L4
L6:
	;
	v21 = v15
	v22 = v16
	v23 = v19
	goto L9
L7:
	;
	v47 = int32(0)
	v48 = v16
	goto L5
L8:
	;
	v47 = v44 & int32(255)
	v48 = v43
	goto L5
L9:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v25 == int32(0) {
		v43 = v22
		v44 = v23
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v43 = v37
	v44 = int32(0)
	goto L8
L11:
	;
	v29 = v23 & int32(255)
	if v29 == v25 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v36 = int32(1)
	v37 = v22 + v36
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v38 != 0 {
		v21 = v21 + v36
		v22 = v37
		v23 = v38
		goto L9
	} else {
		goto L15
	}
L13:
	;
	v31 = F_tolower(m, v29)
	mBase = m.M
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v33 = F_tolower(m, v32)
	mBase = m.M
	if v31 == v33 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v43 = v22
	v44 = v35
	goto L8
L15:
	;
	goto L10
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v55 != int32(3) {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v60 = F_objectGetVal(m, v59)
	mBase = m.M
	v61 = int32(_a25)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v64 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	F_addReplyStatus(m, l0, int32(_a645))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L45
	} else {
		goto L47
	}
L19:
	;
	if v96-v98 == int32(0) {
		goto L18
	} else {
		goto L31
	}
L20:
	;
	v96 = F_tolower(m, v92)
	mBase = m.M
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	v98 = F_tolower(m, v97)
	mBase = m.M
	goto L19
L21:
	;
	v66 = v60
	v67 = v61
	v68 = v64
	goto L24
L22:
	;
	v92 = int32(0)
	v93 = v61
	goto L20
L23:
	;
	v92 = v89 & int32(255)
	v93 = v88
	goto L20
L24:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v70 == int32(0) {
		v88 = v67
		v89 = v68
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v88 = v82
	v89 = int32(0)
	goto L23
L26:
	;
	v74 = v68 & int32(255)
	if v74 == v70 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v81 = int32(1)
	v82 = v67 + v81
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v83 != 0 {
		v66 = v66 + v81
		v67 = v82
		v68 = v83
		goto L24
	} else {
		goto L30
	}
L28:
	;
	v76 = F_tolower(m, v74)
	mBase = m.M
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v78 = F_tolower(m, v77)
	mBase = m.M
	if v76 == v78 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v88 = v67
	v89 = v80
	goto L23
L30:
	;
	goto L25
L31:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	v104 = F_objectGetVal(m, v103)
	mBase = m.M
	v105 = int32(_a26)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v108 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if v140-v142 == int32(0) {
		goto L18
	} else {
		goto L44
	}
L33:
	;
	v140 = F_tolower(m, v136)
	mBase = m.M
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	v142 = F_tolower(m, v141)
	mBase = m.M
	goto L32
L34:
	;
	v110 = v104
	v111 = v105
	v112 = v108
	goto L37
L35:
	;
	v136 = int32(0)
	v137 = v105
	goto L33
L36:
	;
	v136 = v133 & int32(255)
	v137 = v132
	goto L33
L37:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v114 == int32(0) {
		v132 = v111
		v133 = v112
		goto L36
	} else {
		goto L39
	}
L38:
	;
	v132 = v126
	v133 = int32(0)
	goto L36
L39:
	;
	v118 = v112 & int32(255)
	if v118 == v114 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v125 = int32(1)
	v126 = v111 + v125
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	if v127 != 0 {
		v110 = v110 + v125
		v111 = v126
		v112 = v127
		goto L37
	} else {
		goto L43
	}
L41:
	;
	v120 = F_tolower(m, v118)
	mBase = m.M
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v122 = F_tolower(m, v121)
	mBase = m.M
	if v120 == v122 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v132 = v111
	v133 = v124
	goto L36
L43:
	;
	goto L38
L44:
	;
	F_addReplyError(m, l0, int32(_a646))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	return
L46:
	;
	goto L18
L47:
	;
	goto L1
L48:
	;
	v160 = F_isHLLObjectOrReply(m, l0, v155)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L45
	} else {
		goto L52
	}
L49:
	;
	if v155 != 0 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	F_addReplyError(m, l0, int32(_a647))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	goto L1
L52:
	;
	if v160 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	v165 = F_dbUnshareStringValue(m, v162, v164, v155)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L45
	} else {
		goto L54
	}
L54:
	;
	v167 = F_objectGetVal(m, v165)
	mBase = m.M
	v168 = int32(_a648)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v171 != 0 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v263 = int32(_a649)
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v266 != 0 {
		goto L84
	} else {
		goto L85
	}
L56:
	;
	if v203-v205 != 0 {
		goto L55
	} else {
		goto L68
	}
L57:
	;
	v203 = F_tolower(m, v199)
	mBase = m.M
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	v205 = F_tolower(m, v204)
	mBase = m.M
	goto L56
L58:
	;
	v173 = v15
	v174 = v168
	v175 = v171
	goto L61
L59:
	;
	v199 = int32(0)
	v200 = v168
	goto L57
L60:
	;
	v199 = v196 & int32(255)
	v200 = v195
	goto L57
L61:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v177 == int32(0) {
		v195 = v174
		v196 = v175
		goto L60
	} else {
		goto L63
	}
L62:
	;
	v195 = v189
	v196 = int32(0)
	goto L60
L63:
	;
	v181 = v175 & int32(255)
	if v181 == v177 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v188 = int32(1)
	v189 = v174 + v188
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
	if v190 != 0 {
		v173 = v173 + v188
		v174 = v189
		v175 = v190
		goto L61
	} else {
		goto L67
	}
L65:
	;
	v183 = F_tolower(m, v181)
	mBase = m.M
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v185 = F_tolower(m, v184)
	mBase = m.M
	if v183 == v185 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	v195 = v174
	v196 = v187
	goto L60
L67:
	;
	goto L62
L68:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v207 != int32(3) {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
	if v210 != int32(1) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v226 = F_objectGetVal(m, v165)
	mBase = m.M
	F_addReplyArrayLen(m, l0, int32(16384))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L45
	} else {
		goto L76
	}
L71:
	;
	v213 = F_hllSparseToDense(m, v165)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L45
	} else {
		goto L73
	}
L72:
	;
	v220 = int32(_a69)
	v222 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v222 + int64(1)
	goto L70
L73:
	;
	if v213 != int32(-1) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	F_addReplyError(m, l0, int32(_a641))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L45
	} else {
		goto L75
	}
L75:
	;
	goto L1
L76:
	;
	v235 = int32(0)
	goto L77
L77:
	;
	v241 = int32(6)
	v242 = v235 * v241
	v245 = v226 + int32(16) + int32(base.Ui32(v242)>>(uint(int32(3))%32))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+1)))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(int32(base.Ui32(v246<<(uint(int32(8))%32)|v249)>>(uint(v242&v241)%32))&int32(63)))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L45
	} else {
		goto L79
	}
L79:
	;
	v260 = v235 + int32(1)
	if v260 != int32(16384) {
		v235 = v260
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L1
L81:
	;
	v536 = int32(_a650)
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v539 != 0 {
		goto L160
	} else {
		goto L161
	}
L82:
	;
	if v298-v300 != 0 {
		goto L81
	} else {
		goto L94
	}
L83:
	;
	v298 = F_tolower(m, v294)
	mBase = m.M
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	v300 = F_tolower(m, v299)
	mBase = m.M
	goto L82
L84:
	;
	v268 = v15
	v269 = v263
	v270 = v266
	goto L87
L85:
	;
	v294 = int32(0)
	v295 = v263
	goto L83
L86:
	;
	v294 = v291 & int32(255)
	v295 = v290
	goto L83
L87:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	if v272 == int32(0) {
		v290 = v269
		v291 = v270
		goto L86
	} else {
		goto L89
	}
L88:
	;
	v290 = v284
	v291 = int32(0)
	goto L86
L89:
	;
	v276 = v270 & int32(255)
	if v276 == v272 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v283 = int32(1)
	v284 = v269 + v283
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)))
	if v285 != 0 {
		v268 = v268 + v283
		v269 = v284
		v270 = v285
		goto L87
	} else {
		goto L93
	}
L91:
	;
	v278 = F_tolower(m, v276)
	mBase = m.M
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	v280 = F_tolower(m, v279)
	mBase = m.M
	if v278 == v280 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	v290 = v269
	v291 = v282
	goto L86
L93:
	;
	goto L88
L94:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v302 != int32(3) {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	v305 = F_objectGetVal(m, v165)
	mBase = m.M
	v306 = F_objectGetVal(m, v165)
	mBase = m.M
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306+int32(-1)))))
	switch v312 & int32(7) {
	case 0:
		goto L102
	case 1:
		goto L101
	case 2:
		goto L100
	case 3:
		goto L99
	case 4:
		goto L98
	default:
		v329 = int32(0)
		goto L97
	}
L96:
	;
	v332 = F_sdsempty(m)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L45
	} else {
		goto L103
	}
L97:
	;
	v331 = v329
	goto L96
L98:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v306+int32(-17))))
	v329 = v328
	goto L97
L99:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v306+int32(-9))))
	v331 = v325
	goto L96
L100:
	;
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v306+int32(-5)))))
	v331 = v322
	goto L96
L101:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306+int32(-3)))))
	v331 = v319
	goto L96
L102:
	;
	v331 = int32(base.Ui32(v312) >> (uint(int32(3)) % 32))
	goto L96
L103:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
	if v334 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if base.Ui32(v331) < base.Ui32(int32(17)) {
		v413 = v332
		goto L108
	} else {
		goto L109
	}
L105:
	;
	F_sdsfree(m, v332)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L45
	} else {
		goto L106
	}
L106:
	;
	F_addReplyError(m, l0, int32(_a651))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L45
	} else {
		goto L107
	}
L107:
	;
	goto L1
L108:
	;
	v416 = int32(_a10)
	v422 = v413 + int32(-1)
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422))))
	switch v423 & int32(7) {
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
		v440 = int32(0)
		goto L122
	}
L109:
	;
	v349 = v305 + int32(16)
	v352 = v332
	goto L110
L110:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	v357 = v355 & int32(192)
	if v357 == int32(64) {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	v413 = v406
	goto L108
L112:
	;
	if base.Ui32(v405) < base.Ui32(v305+v331) {
		v349 = v405
		v352 = v406
		goto L110
	} else {
		goto L120
	}
L113:
	;
	v388 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v355&int32(3) + v388
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(base.Ui32(v355)>>(uint(int32(2))%32))&int32(31) + v388
	v403 = F_sdscatprintf(m, v352, int32(_a652), v11+int32(16))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L45
	} else {
		goto L119
	}
L114:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v370 | v355<<(uint(int32(8))%32)&int32(16128) + int32(1)
	v384 = F_sdscatprintf(m, v352, int32(_a653), v11+int32(48))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L45
	} else {
		goto L118
	}
L115:
	;
	if v357 != 0 {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v360 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v355 + v360
	v368 = F_sdscatprintf(m, v352, int32(_a654), v11+int32(32))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L45
	} else {
		goto L117
	}
L117:
	;
	v405 = v349 + v360
	v406 = v368
	goto L112
L118:
	;
	v405 = v349 + int32(2)
	v406 = v384
	goto L112
L119:
	;
	v405 = v349 + v388
	v406 = v403
	goto L112
L120:
	;
	goto L111
L121:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413+int32(-1)))))
	switch v512 & int32(7) {
	case 0:
		goto L154
	case 1:
		goto L153
	case 2:
		goto L152
	case 3:
		goto L151
	case 4:
		goto L150
	default:
		v529 = int32(0)
		goto L149
	}
L122:
	;
	v443 = v413 + v440 + int32(-1)
	if base.Ui32(v443) < base.Ui32(v413) {
		v461 = v413
		goto L128
	} else {
		goto L129
	}
L123:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v413+int32(-17))))
	v440 = v439
	goto L122
L124:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v413+int32(-9))))
	v440 = v436
	goto L122
L125:
	;
	v433 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v413+int32(-5)))))
	v440 = v433
	goto L122
L126:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413+int32(-3)))))
	v440 = v430
	goto L122
L127:
	;
	v440 = int32(base.Ui32(v423) >> (uint(int32(3)) % 32))
	goto L122
L128:
	;
	if base.Ui32(v443) <= base.Ui32(v461) {
		v477 = v443
		goto L134
	} else {
		goto L135
	}
L129:
	;
	v449 = v413
	goto L130
L130:
	;
	v450 = int32(*(*int8)(unsafe.Add(mBase, uint32(v449))))
	v451 = F_strchr(m, v416, v450)
	mBase = m.M
	if v451 == int32(0) {
		v461 = v449
		goto L128
	} else {
		goto L132
	}
L131:
	;
	v461 = v455
	goto L128
L132:
	;
	v455 = v449 + int32(1)
	if base.Ui32(v455) <= base.Ui32(v443) {
		v449 = v455
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	v482 = v477 - v461 + int32(1)
	if v413 == v461 {
		goto L140
	} else {
		goto L141
	}
L135:
	;
	v465 = v443
	goto L136
L136:
	;
	v468 = int32(*(*int8)(unsafe.Add(mBase, uint32(v465))))
	v469 = F_strchr(m, v416, v468)
	mBase = m.M
	if v469 == int32(0) {
		v477 = v465
		goto L134
	} else {
		goto L138
	}
L137:
	;
	v477 = v461
	goto L134
L138:
	;
	v473 = v465 + int32(-1)
	if base.Ui32(v461) < base.Ui32(v473) {
		v465 = v473
		goto L136
	} else {
		goto L139
	}
L139:
	;
	goto L137
L140:
	;
	v486 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v413+v482))) = uint8(v486)
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422))))
	switch v488 & int32(7) {
	case 0:
		goto L147
	case 1:
		goto L146
	case 2:
		goto L145
	case 3:
		goto L144
	case 4:
		goto L143
	default:
		goto L142
	}
L141:
	;
	v484 = F_memmove(m, v413, v461, v482)
	mBase = m.M
	goto L140
L142:
	;
	goto L121
L143:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v413+int32(-17)))) = base.I64_extend_i32_u(v482)
	goto L142
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413+int32(-9)))) = v482
	goto L121
L145:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v413+int32(-5)))) = uint16(v482)
	goto L121
L146:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v413+int32(-3)))) = uint8(v482)
	goto L121
L147:
	;
	v492 = v482 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v422))) = uint8(v492)
	goto L121
L148:
	;
	F_addReplyBulkCBuffer(m, l0, v413, v531)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L45
	} else {
		goto L155
	}
L149:
	;
	v531 = v529
	goto L148
L150:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v413+int32(-17))))
	v529 = v528
	goto L149
L151:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v413+int32(-9))))
	v531 = v525
	goto L148
L152:
	;
	v522 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v413+int32(-5)))))
	v531 = v522
	goto L148
L153:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413+int32(-3)))))
	v531 = v519
	goto L148
L154:
	;
	v531 = int32(base.Ui32(v512) >> (uint(int32(3)) % 32))
	goto L148
L155:
	;
	F_sdsfree(m, v413)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L45
	} else {
		goto L156
	}
L156:
	;
	goto L1
L157:
	;
	v586 = int32(_a655)
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v589 != 0 {
		goto L176
	} else {
		goto L177
	}
L158:
	;
	if v571-v573 != 0 {
		goto L157
	} else {
		goto L170
	}
L159:
	;
	v571 = F_tolower(m, v567)
	mBase = m.M
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568))))
	v573 = F_tolower(m, v572)
	mBase = m.M
	goto L158
L160:
	;
	v541 = v15
	v542 = v536
	v543 = v539
	goto L163
L161:
	;
	v567 = int32(0)
	v568 = v536
	goto L159
L162:
	;
	v567 = v564 & int32(255)
	v568 = v563
	goto L159
L163:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542))))
	if v545 == int32(0) {
		v563 = v542
		v564 = v543
		goto L162
	} else {
		goto L165
	}
L164:
	;
	v563 = v557
	v564 = int32(0)
	goto L162
L165:
	;
	v549 = v543 & int32(255)
	if v549 == v545 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v556 = int32(1)
	v557 = v542 + v556
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+1)))
	if v558 != 0 {
		v541 = v541 + v556
		v542 = v557
		v543 = v558
		goto L163
	} else {
		goto L169
	}
L167:
	;
	v551 = F_tolower(m, v549)
	mBase = m.M
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542))))
	v553 = F_tolower(m, v552)
	mBase = m.M
	if v551 == v553 {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541))))
	v563 = v542
	v564 = v555
	goto L162
L169:
	;
	goto L164
L170:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v575 != int32(3) {
		goto L2
	} else {
		goto L171
	}
L171:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v578<<(uint(int32(2))%32))+uint32(_consts[297])))
	F_addReplyStatus(m, l0, v583)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L45
	} else {
		goto L172
	}
L172:
	;
	goto L1
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v15
	F_addReplyErrorFormat(m, l0, int32(_a656), v11+int32(64))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L45
	} else {
		goto L198
	}
L174:
	;
	if v621-v623 != 0 {
		goto L173
	} else {
		goto L186
	}
L175:
	;
	v621 = F_tolower(m, v617)
	mBase = m.M
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618))))
	v623 = F_tolower(m, v622)
	mBase = m.M
	goto L174
L176:
	;
	v591 = v15
	v592 = v586
	v593 = v589
	goto L179
L177:
	;
	v617 = int32(0)
	v618 = v586
	goto L175
L178:
	;
	v617 = v614 & int32(255)
	v618 = v613
	goto L175
L179:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592))))
	if v595 == int32(0) {
		v613 = v592
		v614 = v593
		goto L178
	} else {
		goto L181
	}
L180:
	;
	v613 = v607
	v614 = int32(0)
	goto L178
L181:
	;
	v599 = v593 & int32(255)
	if v599 == v595 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v606 = int32(1)
	v607 = v592 + v606
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591)+1)))
	if v608 != 0 {
		v591 = v591 + v606
		v592 = v607
		v593 = v608
		goto L179
	} else {
		goto L185
	}
L183:
	;
	v601 = F_tolower(m, v599)
	mBase = m.M
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592))))
	v603 = F_tolower(m, v602)
	mBase = m.M
	if v601 == v603 {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591))))
	v613 = v592
	v614 = v605
	goto L178
L185:
	;
	goto L180
L186:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v625 != int32(3) {
		goto L2
	} else {
		goto L187
	}
L187:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
	if v628 != int32(1) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	if v628 == int32(1) {
		goto L194
	} else {
		goto L195
	}
L189:
	;
	v631 = F_hllSparseToDense(m, v165)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L45
	} else {
		goto L191
	}
L190:
	;
	v638 = int32(_a69)
	v640 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v640 + int64(1)
	goto L188
L191:
	;
	if v631 != int32(-1) {
		goto L190
	} else {
		goto L192
	}
L192:
	;
	F_addReplyError(m, l0, int32(_a641))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L45
	} else {
		goto L193
	}
L193:
	;
	goto L1
L194:
	;
	v649 = int32(16)
	goto L196
L195:
	;
	v649 = int32(12)
	goto L196
L196:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v649)+uint32(_consts[77])))
	F_addReply(m, l0, v651)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L45
	} else {
		goto L197
	}
L197:
	;
	goto L1
L198:
	;
	goto L1
L199:
	;
	goto L1
}
func F_pfmergeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
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
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
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
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
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
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int64
	_ = v266
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16400)
	m.G0 = v12
	v20 = F__emscripten_memset_bulkmem(m, v12+int32(16), base.I32_extend8_s(v2), int32(16384))
	mBase = m.M
	goto L1
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v22 < int32(2) {
		v67 = int32(1)
		goto L6
	} else {
		goto L7
	}
L2:
	;
	m.G0 = v12 + int32(16400)
	return
L3:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v246 = F_objectGetVal(m, v245)
	mBase = m.M
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+15)))
	v249 = v247 | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v246)+15)) = uint8(v249)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	F_signalModifiedKey(m, l0, v251, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L11
	} else {
		goto L47
	}
L4:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v182 = F_hllSparseToDense(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L11
	} else {
		goto L41
	}
L5:
	;
	F_addReplyError(m, l0, int32(_a641))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L11
	} else {
		goto L39
	}
L6:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v76 = F_lookupKeyWrite(m, v73, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L11
	} else {
		goto L23
	}
L7:
	;
	v28 = v2
	v29 = int32(1)
	goto L8
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v29<<(uint(int32(2))%32))))
	v41 = F_lookupKeyRead(m, v35, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v67 = base.B2i32(v56 == int32(0))
	goto L6
L10:
	;
	v59 = v29 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v59 < v60 {
		v28 = v56
		v29 = v59
		goto L8
	} else {
		goto L20
	}
L11:
	;
	return
L12:
	;
	if v41 == int32(0) {
		v56 = v28
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v45 = F_isHLLObjectOrReply(m, l0, v41)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	if v45 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v47 = F_objectGetVal(m, v41)
	mBase = m.M
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+4)))
	v51 = F_hllMerge(m, v12+int32(16), v41)
	mBase = m.M
	if v51 == int32(-1) {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	if v48 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v55 = v28
	goto L19
L18:
	;
	v55 = int32(1)
	goto L19
L19:
	;
	v56 = v55
	goto L10
L20:
	;
	goto L9
L21:
	;
	if v67 == int32(0) {
		goto L4
	} else {
		goto L29
	}
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v103 = F_dbUnshareStringValue(m, v100, v102, v76)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L11
	} else {
		goto L28
	}
L23:
	;
	if v76 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v80 = F_sdsnewlen(m, int32(0), int32(18))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v82 = int32(65407)
	*(*uint16)(unsafe.Add(mBase, uint32(v80)+16)) = uint16(v82)
	v85 = F_createObject(m, int32(0), v80)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v87 = F_objectGetVal(m, v85)
	mBase = m.M
	v88 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+4)) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = int32(1280072008)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v85
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	F_dbAdd(m, v93, v95, v12+int32(12))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	goto L21
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v103
	goto L21
L29:
	;
	v114 = int32(0)
	goto L30
L30:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(16)+v114))))
	if v123 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v175 = v114 + int32(1)
	if v175 != int32(16384) {
		v114 = v175
		goto L30
	} else {
		goto L38
	}
L33:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v127 = F_objectGetVal(m, v126)
	mBase = m.M
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
	switch v128 {
	case 0:
		goto L35
	case 1:
		goto L34
	default:
		goto L32
	}
L34:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v167 = F_hllSparseSet(m, v166, v114, v123)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L11
	} else {
		goto L37
	}
L35:
	;
	v129 = int32(6)
	v130 = v114 * v129
	v133 = v127 + int32(base.Ui32(v130)>>(uint(int32(3))%32))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+int32(17)))))
	v139 = v130 & v129
	v140 = int32(8) - v139
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+16)))
	if base.Ui32(v123) <= base.Ui32((v136<<(uint(v140)%32)|int32(base.Ui32(v142)>>(uint(v139)%32)))&int32(63)) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v156 = v136&(int32(-64)>>(uint(v140)%32)) | int32(base.Ui32(v123)>>(uint(v140)%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v133+int32(17)))) = uint8(v156)
	v164 = v142&(int32(63)<<(uint(v139)%32)^int32(-1)) | v123<<(uint(v139)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v133+int32(16)))) = uint8(v164)
	goto L32
L37:
	;
	goto L32
L38:
	;
	goto L3
L39:
	;
	goto L2
L40:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v190 = F_objectGetVal(m, v189)
	mBase = m.M
	v197 = int32(0)
	goto L44
L41:
	;
	if v182 != int32(-1) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	F_addReplyError(m, l0, int32(_a641))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	goto L2
L44:
	;
	v203 = int32(6)
	v204 = v197 * v203
	v207 = v190 + int32(16) + int32(base.Ui32(v204)>>(uint(int32(3))%32))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	v211 = v204 & v203
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(16)+v197))))
	v221 = v208&(int32(63)<<(uint(v211)%32)^int32(-1)) | v219<<(uint(v211)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v221)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
	v226 = int32(8) - v211
	v230 = v223&(int32(-64)>>(uint(v226)%32)) | int32(base.Ui32(v219)>>(uint(v226)%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)) = uint8(v230)
	v233 = v197 + int32(1)
	if v233 != int32(16384) {
		v197 = v233
		goto L44
	} else {
		goto L46
	}
L45:
	;
	goto L3
L46:
	;
	goto L45
L47:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+28))
	F_notifyKeyspaceEvent(m, int32(8), int32(_a642), v259, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	v264 = int32(_a69)
	v266 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v266 + int64(1)
	v271 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	goto L2
}
func F_pingCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v3 < int32(3) {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+202)))
		if v8&int32(4) == int32(0) {
			if v3 != int32(1) {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
				F_addReplyBulk(m, l0, v42)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					return
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, _consts[693]))
				F_addReply(m, l0, v38)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
			if v13 != int32(2) {
				if v3 != int32(1) {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
					F_addReplyBulk(m, l0, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						return
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, _consts[693]))
					F_addReply(m, l0, v38)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _consts[846]))
				F_addReply(m, l0, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_addReplyBulkCBuffer(m, l0, int32(_a1359), int32(4))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v24 != int32(1) {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
							F_addReplyBulk(m, l0, v32)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								return
							}
						} else {
							F_addReplyBulkCBuffer(m, l0, int32(_a188), int32(0))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_addReplyErrorArity(m, l0)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	}
}
func F_pop_arg_long_double(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v16 float64
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = (v4 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v8 + int32(16)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(8))))
	v16 = F___trunctfdf2(m, v12, v15)
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v16
	return
}
func F_populateCommandLegacyRangeSpec(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v32 int64
	_ = v32
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int64
	_ = v51
	var v57 int64
	_ = v57
	var v63 int64
	_ = v63
	var v69 int64
	_ = v69
	var v75 int64
	_ = v75
	var v81 int64
	_ = v81
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v167 int64
	_ = v167
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v206 int32
	_ = v206
	v2 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(192)))) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(184)))) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(176)))) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(168)))) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(160)))) = v2
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if v32&int64(2097152) == v2 {
		v40 = v32
	} else {
		v38 = v32 | int64(33554432)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v38
		v40 = v38
	}
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	switch v41 {
	case 0:
		return
	case 1:
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
		if v43 != int32(2) {
			v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v98 = int32(0)
			v101 = v40
			v103 = v98
			v106 = int32(2147483647)
			v107 = v98
			for {
				v111 = v96 + v103*int32(48)
				v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
				if v112 != int32(2) {
					v148 = v101 | int64(33554432)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v148
					v150 = v148
					v152 = v106
					v153 = v107
				} else {
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
					if v115 != int32(2) {
						v148 = v101 | int64(33554432)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v148
						v150 = v148
						v152 = v106
						v153 = v107
					} else {
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
						if v118 != int32(1) {
							v148 = v101 | int64(33554432)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v148
							v150 = v148
							v152 = v106
							v153 = v107
						} else {
							if v107 == int32(0) {
								v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+9)))
								if v127&int32(2) == int32(0) {
									v135 = v101
								} else {
									v133 = v101 | int64(33554432)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v133
									v135 = v133
								}
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
								v137 = int32(0)
								v138 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
								if v136 < v137 {
									v141 = v137
								} else {
									v141 = v138
								}
								v142 = v136 + v141
								if base.Ui32(v142) < base.Ui32(v107) {
									v144 = v107
								} else {
									v144 = v142
								}
								if v106 < v138 {
									v146 = v106
								} else {
									v146 = v138
								}
								v150 = v135
								v152 = v146
								v153 = v144
							} else {
								v123 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
								if v107 != v123+int32(-1) {
									v148 = v101 | int64(33554432)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v148
									v150 = v148
									v152 = v106
									v153 = v107
								} else {
									v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+9)))
									if v127&int32(2) == int32(0) {
										v135 = v101
									} else {
										v133 = v101 | int64(33554432)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v133
										v135 = v133
									}
									v136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
									v137 = int32(0)
									v138 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
									if v136 < v137 {
										v141 = v137
									} else {
										v141 = v138
									}
									v142 = v136 + v141
									if base.Ui32(v142) < base.Ui32(v107) {
										v144 = v107
									} else {
										v144 = v142
									}
									if v106 < v138 {
										v146 = v106
									} else {
										v146 = v138
									}
									v150 = v135
									v152 = v146
									v153 = v144
								}
							}
						}
					}
				}
				v156 = v103 + int32(1)
				if v156 != v41 {
					v101 = v150
					v103 = v156
					v106 = v152
					v107 = v153
					continue
				} else {
					break
				}
				break
			}
			if v152 == int32(2147483647) {
				v167 = v150
				*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v167 | int64(33554432)
				return
			} else {
				if v152 != 0 {
					if v153 == int32(0) {
						F__serverAssert(m, int32(_a1356), int32(_a1240), int32(3288))
						mBase = m.M
						v206 = m.ExcPending
						if v206 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(1)
						v182 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v182
						*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v152
						*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v182
						v187 = int32(0)
						if v153 < v187 {
							v190 = v187
						} else {
							v190 = v152
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v153 - v190
						return
					}
				} else {
					F__serverAssert(m, int32(_a1357), int32(_a1240), int32(3287))
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
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
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
			if v46 != int32(2) {
				v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v98 = int32(0)
				v101 = v40
				v103 = v98
				v106 = int32(2147483647)
				v107 = v98
				for {
					v111 = v96 + v103*int32(48)
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
					if v112 != int32(2) {
						v148 = v101 | int64(33554432)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v148
						v150 = v148
						v152 = v106
						v153 = v107
					} else {
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
						if v115 != int32(2) {
							v148 = v101 | int64(33554432)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v148
							v150 = v148
							v152 = v106
							v153 = v107
						} else {
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
							if v118 != int32(1) {
								v148 = v101 | int64(33554432)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v148
								v150 = v148
								v152 = v106
								v153 = v107
							} else {
								if v107 == int32(0) {
									v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+9)))
									if v127&int32(2) == int32(0) {
										v135 = v101
									} else {
										v133 = v101 | int64(33554432)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v133
										v135 = v133
									}
									v136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
									v137 = int32(0)
									v138 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
									if v136 < v137 {
										v141 = v137
									} else {
										v141 = v138
									}
									v142 = v136 + v141
									if base.Ui32(v142) < base.Ui32(v107) {
										v144 = v107
									} else {
										v144 = v142
									}
									if v106 < v138 {
										v146 = v106
									} else {
										v146 = v138
									}
									v150 = v135
									v152 = v146
									v153 = v144
								} else {
									v123 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
									if v107 != v123+int32(-1) {
										v148 = v101 | int64(33554432)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v148
										v150 = v148
										v152 = v106
										v153 = v107
									} else {
										v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+9)))
										if v127&int32(2) == int32(0) {
											v135 = v101
										} else {
											v133 = v101 | int64(33554432)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v133
											v135 = v133
										}
										v136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
										v137 = int32(0)
										v138 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
										if v136 < v137 {
											v141 = v137
										} else {
											v141 = v138
										}
										v142 = v136 + v141
										if base.Ui32(v142) < base.Ui32(v107) {
											v144 = v107
										} else {
											v144 = v142
										}
										if v106 < v138 {
											v146 = v106
										} else {
											v146 = v138
										}
										v150 = v135
										v152 = v146
										v153 = v144
									}
								}
							}
						}
					}
					v156 = v103 + int32(1)
					if v156 != v41 {
						v101 = v150
						v103 = v156
						v106 = v152
						v107 = v153
						continue
					} else {
						break
					}
					break
				}
				if v152 == int32(2147483647) {
					v167 = v150
					*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v167 | int64(33554432)
					return
				} else {
					if v152 != 0 {
						if v153 == int32(0) {
							F__serverAssert(m, int32(_a1356), int32(_a1240), int32(3288))
							mBase = m.M
							v206 = m.ExcPending
							if v206 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(1)
							v182 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v182
							*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v152
							*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v182
							v187 = int32(0)
							if v153 < v187 {
								v190 = v187
							} else {
								v190 = v152
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v153 - v190
							return
						}
					} else {
						F__serverAssert(m, int32(_a1357), int32(_a1240), int32(3287))
						mBase = m.M
						v164 = m.ExcPending
						if v164 != 0 {
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
			} else {
				v51 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
				*(*int64)(unsafe.Add(mBase, uint32(l0+int32(152)))) = v51
				v57 = *(*int64)(unsafe.Add(mBase, uint32(v42+int32(40))))
				*(*int64)(unsafe.Add(mBase, uint32(l0+int32(192)))) = v57
				v63 = *(*int64)(unsafe.Add(mBase, uint32(v42+int32(32))))
				*(*int64)(unsafe.Add(mBase, uint32(l0+int32(184)))) = v63
				v69 = *(*int64)(unsafe.Add(mBase, uint32(v42+int32(24))))
				*(*int64)(unsafe.Add(mBase, uint32(l0+int32(176)))) = v69
				v75 = *(*int64)(unsafe.Add(mBase, uint32(v42+int32(16))))
				*(*int64)(unsafe.Add(mBase, uint32(l0+int32(168)))) = v75
				v81 = *(*int64)(unsafe.Add(mBase, uint32(v42+int32(8))))
				*(*int64)(unsafe.Add(mBase, uint32(l0+int32(160)))) = v81
				v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(9)))))
				if v85&int32(2) == int32(0) {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v40 | int64(33554432)
					return
				}
			}
		}
	default:
		if v41 < int32(1) {
			v167 = v40
			*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v167 | int64(33554432)
			return
		} else {
			v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v98 = int32(0)
			v101 = v40
			v103 = v98
			v106 = int32(2147483647)
			v107 = v98
			for {
				v111 = v96 + v103*int32(48)
				v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
				if v112 != int32(2) {
					v148 = v101 | int64(33554432)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v148
					v150 = v148
					v152 = v106
					v153 = v107
				} else {
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
					if v115 != int32(2) {
						v148 = v101 | int64(33554432)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v148
						v150 = v148
						v152 = v106
						v153 = v107
					} else {
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
						if v118 != int32(1) {
							v148 = v101 | int64(33554432)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v148
							v150 = v148
							v152 = v106
							v153 = v107
						} else {
							if v107 == int32(0) {
								v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+9)))
								if v127&int32(2) == int32(0) {
									v135 = v101
								} else {
									v133 = v101 | int64(33554432)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v133
									v135 = v133
								}
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
								v137 = int32(0)
								v138 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
								if v136 < v137 {
									v141 = v137
								} else {
									v141 = v138
								}
								v142 = v136 + v141
								if base.Ui32(v142) < base.Ui32(v107) {
									v144 = v107
								} else {
									v144 = v142
								}
								if v106 < v138 {
									v146 = v106
								} else {
									v146 = v138
								}
								v150 = v135
								v152 = v146
								v153 = v144
							} else {
								v123 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
								if v107 != v123+int32(-1) {
									v148 = v101 | int64(33554432)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v148
									v150 = v148
									v152 = v106
									v153 = v107
								} else {
									v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+9)))
									if v127&int32(2) == int32(0) {
										v135 = v101
									} else {
										v133 = v101 | int64(33554432)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v133
										v135 = v133
									}
									v136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
									v137 = int32(0)
									v138 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
									if v136 < v137 {
										v141 = v137
									} else {
										v141 = v138
									}
									v142 = v136 + v141
									if base.Ui32(v142) < base.Ui32(v107) {
										v144 = v107
									} else {
										v144 = v142
									}
									if v106 < v138 {
										v146 = v106
									} else {
										v146 = v138
									}
									v150 = v135
									v152 = v146
									v153 = v144
								}
							}
						}
					}
				}
				v156 = v103 + int32(1)
				if v156 != v41 {
					v101 = v150
					v103 = v156
					v106 = v152
					v107 = v153
					continue
				} else {
					break
				}
				break
			}
			if v152 == int32(2147483647) {
				v167 = v150
				*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v167 | int64(33554432)
				return
			} else {
				if v152 != 0 {
					if v153 == int32(0) {
						F__serverAssert(m, int32(_a1356), int32(_a1240), int32(3288))
						mBase = m.M
						v206 = m.ExcPending
						if v206 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(1)
						v182 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v182
						*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v152
						*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v182
						v187 = int32(0)
						if v153 < v187 {
							v190 = v187
						} else {
							v190 = v152
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v153 - v190
						return
					}
				} else {
					F__serverAssert(m, int32(_a1357), int32(_a1240), int32(3287))
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
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
		}
	}
}
func F_postWriteToReplica(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v7 < int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a803), int32(_a774), int32(2476))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L13
	} else {
		goto L15
	}
L2:
	;
	return
L3:
	;
	v10 = int32(_a69)
	v12 = *(*int64)(unsafe.Add(mBase, _consts[430]))
	*(*int64)(unsafe.Add(mBase, _consts[430])) = v12 + base.I64_extend_i32_u(v7)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+188))
	v18 = v17 + v7
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+184))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if base.Ui32(v21) <= base.Ui32(v18) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if base.Ui32(v47) < base.Ui32(v45) {
		goto L1
	} else {
		goto L12
	}
L5:
	;
	v23 = v20
	v26 = v18
	v27 = v19
	v28 = v21
	goto L7
L6:
	;
	v43 = v19
	v45 = v18
	v47 = v21
	goto L4
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v43 = v29
	v45 = v39
	v47 = v40
	goto L4
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v30 + int32(-1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v35 + int32(1)
	v39 = v26 - v28
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
	if base.Ui32(v40) <= base.Ui32(v39) {
		v23 = v34
		v26 = v39
		v27 = v29
		v28 = v40
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v43 = v27
	v45 = v26
	v47 = v28
	goto L4
L11:
	;
	goto L8
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+188)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v16)+184)) = v43
	F_incrementalTrimReplicationBacklog(m, int32(64))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	goto L2
L15:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pqsort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	var v16 int32
	_ = v16
	F__pqsort(m, l0, l1, l2, l3, l0+l4*l2, l0+(l5+int32(1))*l2+int32(-1))
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		return
	}
}
func F_prefetchCommandsBatchInit(m *base.Module) {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, _consts[338]))
	if v5 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[339]))
		if v7 == int32(0) {
			return
		} else {
			v12 = F_valkey_calloc(m, int32(44))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[338])) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v7
				v17 = v7 << (uint(int32(2)) % 32)
				v18 = F_valkey_calloc(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, _consts[338]))
					*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v18
					v23 = F_valkey_calloc(m, v17)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, _consts[338]))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v23
						v28 = F_valkey_calloc(m, v17)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, _consts[338]))
							*(*int32)(unsafe.Add(mBase, uint32(v31)+36)) = v28
							v33 = F_valkey_calloc(m, v17)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, _consts[338]))
								*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v33
								v40 = F_valkey_calloc(m, v7*int32(48))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, _consts[338]))
									*(*int32)(unsafe.Add(mBase, uint32(v43)+40)) = v40
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_prepend_alloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var __phi95 int32
	_ = __phi95
	var v101 int32
	_ = v101
	var __phi101 int32
	_ = __phi101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v292 int32
	_ = v292
	v11 = int32(-8)
	v13 = int32(7)
	v15 = l0 + (v11-l0)&v13
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l2 | int32(3)
	v23 = l1 + (v11-l1)&v13
	v24 = v15 + l2
	v25 = v23 - v24
	v27 = *(*int32)(unsafe.Add(mBase, _consts[1213]))
	if v23 != v27 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v15 + int32(8)
L2:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[1208]))
	if v23 != v40 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1213])) = v24
	v33 = *(*int32)(unsafe.Add(mBase, _consts[1214]))
	v34 = v33 + v25
	*(*int32)(unsafe.Add(mBase, _consts[1214])) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v34 | int32(1)
	goto L1
L4:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v54&int32(3) != int32(1) {
		v170 = v25
		v171 = v54
		v174 = v23
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v42 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1208])) = v24
	v46 = *(*int32)(unsafe.Add(mBase, _consts[1210]))
	v47 = v46 + v25
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v47 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24+v47))) = v47
	goto L1
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = v171 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v170 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24+v170))) = v170
	if base.Ui32(int32(255)) < base.Ui32(v170) {
		goto L37
	} else {
		goto L38
	}
L7:
	;
	v60 = v54 & int32(-8)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if base.Ui32(int32(255)) < base.Ui32(v54) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v168 = v23 + v60
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v170 = v60 + v25
	v171 = v169
	v174 = v168
	goto L6
L9:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	if v61 == v23 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v61 != v64 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v64
	goto L8
L12:
	;
	v66 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = v68 & base.I32_rotl(int32(-2), int32(base.Ui32(v54)>>(uint(int32(3))%32)))
	goto L8
L13:
	;
	if v77 == int32(0) {
		goto L8
	} else {
		goto L25
	}
L14:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	if v82 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v79
	v115 = v61
	goto L13
L16:
	;
	v115 = int32(0)
	goto L13
L17:
	;
	__phi95 = v92
	__phi101 = v93
	v95 = __phi95
	v101 = __phi101
	goto L21
L18:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v87 == int32(0) {
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v92 = v82
	v93 = v23 + int32(20)
	goto L17
L20:
	;
	v92 = v87
	v93 = v23 + int32(16)
	goto L17
L21:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	if v106 != 0 {
		__phi95 = v106
		__phi101 = v95 + int32(20)
		v95 = __phi95
		v101 = __phi101
		goto L21
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = int32(0)
	v115 = v95
	goto L13
L23:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	if v109 != 0 {
		__phi95 = v109
		__phi101 = v95 + int32(16)
		v95 = __phi95
		v101 = __phi101
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v127 = v125 << (uint(int32(2)) % 32)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v127)+uint32(_consts[1211])))
	if v23 != v130 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+24)) = v77
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v147 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	if v140 != v23 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+uint32(_consts[1211]))) = v115
	if v115 != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v133 = int32(0)
	v135 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v135 & base.I32_rotl(int32(-2), v125)
	goto L8
L30:
	;
	if v115 == int32(0) {
		goto L8
	} else {
		goto L33
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = v115
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = v115
	goto L30
L33:
	;
	goto L26
L34:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	if v152 == int32(0) {
		goto L8
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+16)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v147)+24)) = v115
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+20)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v152)+24)) = v115
	goto L8
L37:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v170) {
		v226 = int32(31)
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v191 = v170 & int32(-8)
	v193 = v191 + int32(9128464)
	v195 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	v199 = int32(1) << (uint(int32(base.Ui32(v170)>>(uint(int32(3))%32))) % 32)
	if v195&v199 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+uint32(_consts[1215]))) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v205)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v205
	goto L1
L40:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v191)+uint32(_consts[1215])))
	v205 = v204
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1209])) = v195 | v199
	v205 = v193
	goto L39
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v226
	*(*int64)(unsafe.Add(mBase, uint32(v24)+16)) = int64(0)
	v231 = v226 << (uint(int32(2)) % 32)
	v235 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
	v237 = int32(1) << (uint(v226) % 32)
	if v235&v237 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v216 = base.I32_clz(int32(base.Ui32(v170) >> (uint(int32(8)) % 32)))
	v219 = int32(1)
	v226 = int32(base.Ui32(v170)>>(uint(int32(38)-v216)%32))&v219 - v216<<(uint(v219)%32) + int32(62)
	goto L42
L44:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v261)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v292)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v261)+8)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v292
	goto L1
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v24
	goto L1
L46:
	;
	if v226 == int32(31) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v235 | v237
	*(*int32)(unsafe.Add(mBase, uint32(v231)+uint32(_consts[1211]))) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v231 + int32(9128728)
	goto L45
L48:
	;
	v251 = int32(0)
	goto L50
L49:
	;
	v251 = int32(25) - int32(base.Ui32(v226)>>(uint(int32(1))%32))
	goto L50
L50:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v231)+uint32(_consts[1211])))
	v256 = v170 << (uint(v251) % 32)
	v261 = v253
	goto L51
L51:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if v264&int32(-8) == v170 {
		goto L44
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274+int32(16)))) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v261
	goto L45
L53:
	;
	v274 = v261 + int32(base.Ui32(v256)>>(uint(int32(29))%32))&int32(4)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+16))
	if v275 != 0 {
		v256 = v256 << (uint(int32(1)) % 32)
		v261 = v275
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
}
func F_preventCommandPropagation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v2 | int32(2097152)
	return
}
func F_primaryTryPartialResynchronization(m *base.Module, l0 int32, l1 int64) int32 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int64
	_ = v161
	var v163 int64
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v177 int64
	_ = v177
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int64
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int64
	_ = v257
	var v260 int64
	_ = v260
	var v262 int64
	_ = v262
	var v264 int64
	_ = v264
	var v267 int64
	_ = v267
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v273 int64
	_ = v273
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v469 int32
	_ = v469
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int64
	_ = v513
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int64
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int64
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int64
	_ = v626
	var v627 int64
	_ = v627
	var v630 int64
	_ = v630
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	v9 = m.G0
	v11 = v9 - int32(320)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = F_objectGetVal(m, v14)
	mBase = m.M
	v18 = int32(_a920)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v11 + int32(320)
	return v677
L2:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	if v158 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L3:
	;
	if v53-v55 == int32(0) {
		goto L2
	} else {
		goto L15
	}
L4:
	;
	v53 = F_tolower(m, v49)
	mBase = m.M
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v55 = F_tolower(m, v54)
	mBase = m.M
	goto L3
L5:
	;
	v23 = v15
	v24 = v18
	v25 = v21
	goto L8
L6:
	;
	v49 = int32(0)
	v50 = v18
	goto L4
L7:
	;
	v49 = v46 & int32(255)
	v50 = v45
	goto L4
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v27 == int32(0) {
		v45 = v24
		v46 = v25
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v45 = v39
	v46 = int32(0)
	goto L7
L10:
	;
	v31 = v25 & int32(255)
	if v31 == v27 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = int32(1)
	v39 = v24 + v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v40 != 0 {
		v23 = v23 + v38
		v24 = v39
		v25 = v40
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v33 = F_tolower(m, v31)
	mBase = m.M
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v35 = F_tolower(m, v34)
	mBase = m.M
	if v33 == v35 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v45 = v24
	v46 = v37
	goto L7
L14:
	;
	goto L9
L15:
	;
	v60 = *(*int64)(unsafe.Add(mBase, _consts[508]))
	v63 = int32(_a921)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v66 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v103 == int32(63) {
		goto L31
	} else {
		goto L32
	}
L17:
	;
	if v101 != 0 {
		goto L16
	} else {
		goto L29
	}
L18:
	;
	v98 = F_tolower(m, v94)
	mBase = m.M
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v100 = F_tolower(m, v99)
	mBase = m.M
	v101 = v98 - v100
	goto L17
L19:
	;
	v68 = v15
	v69 = v63
	v70 = v66
	goto L22
L20:
	;
	v94 = int32(0)
	v95 = v63
	goto L18
L21:
	;
	v94 = v91 & int32(255)
	v95 = v90
	goto L18
L22:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v72 == int32(0) {
		v90 = v69
		v91 = v70
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v90 = v84
	v91 = int32(0)
	goto L21
L24:
	;
	v76 = v70 & int32(255)
	if v76 == v72 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v83 = int32(1)
	v84 = v69 + v83
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	if v85 != 0 {
		v68 = v68 + v83
		v69 = v84
		v70 = v85
		goto L22
	} else {
		goto L28
	}
L26:
	;
	v78 = F_tolower(m, v76)
	mBase = m.M
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v80 = F_tolower(m, v79)
	mBase = m.M
	if v78 == v80 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v90 = v69
	v91 = v82
	goto L21
L28:
	;
	goto L23
L29:
	;
	if l1 <= v60 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	goto L16
L31:
	;
	v141 = int32(-1)
	v143 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v143 {
		v677 = v141
		goto L1
	} else {
		goto L40
	}
L32:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v101 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v130 = int32(-1)
	if int32(2) < v107 {
		v677 = v130
		goto L1
	} else {
		goto L38
	}
L34:
	;
	v110 = int32(-1)
	if int32(2) < v107 {
		v677 = v110
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v11)+168)) = int32(_a921)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = int32(_a920)
	F__serverLog(m, int32(2), int32(_a922), v11+int32(160))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	return int32(0)
L37:
	;
	v677 = v110
	goto L1
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+152)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v11)+144)) = l1
	F__serverLog(m, int32(2), int32(_a923), v11+int32(144))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v677 = v130
	goto L1
L40:
	;
	v146 = F_replicationGetReplicaName(m, l0)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v146
	F__serverLog(m, int32(2), int32(_a924), v11+int32(128))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v677 = v141
	goto L1
L43:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v232 != 0 {
		goto L62
	} else {
		goto L63
	}
L44:
	;
	v211 = int32(-1)
	v213 = *(*int64)(unsafe.Add(mBase, _consts[31]))
	if l1 <= v213 {
		v677 = v211
		goto L1
	} else {
		goto L56
	}
L45:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v193 {
		goto L44
	} else {
		goto L53
	}
L46:
	;
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v158)+24))
	if l1 < v161 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v167 {
		goto L44
	} else {
		goto L50
	}
L48:
	;
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v158)+16))
	if l1 <= v163+v161 {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v170 = F_replicationGetReplicaName(m, l0)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L36
	} else {
		goto L51
	}
L51:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v173)+16))
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v173)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(64)))) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(72)))) = v174 + v177
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = l1
	F__serverLog(m, int32(2), int32(_a925), v11+int32(48))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L36
	} else {
		goto L52
	}
L52:
	;
	goto L44
L53:
	;
	v196 = F_replicationGetReplicaName(m, l0)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L36
	} else {
		goto L54
	}
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v196
	F__serverLog(m, int32(2), int32(_a926), v11+int32(32))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L36
	} else {
		goto L55
	}
L55:
	;
	goto L44
L56:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v216 {
		v677 = v211
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v219 = F_replicationGetReplicaName(m, l0)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L36
	} else {
		goto L58
	}
L58:
	;
	v224 = *(*int64)(unsafe.Add(mBase, _consts[31]))
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(16)))) = v224
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v219
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = l1
	F__serverLog(m, int32(3), int32(_a927), v11)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L36
	} else {
		goto L59
	}
L59:
	;
	v677 = v211
	goto L1
L60:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v252 | int32(2)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v256)+168))
	if v257 == int64(0) {
		v508 = v256
		goto L75
	} else {
		goto L76
	}
L61:
	;
	goto L60
L62:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v236 != int32(1) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v233 == int32(0) {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v244 != int32(1) {
		goto L61
	} else {
		goto L70
	}
L66:
	;
	goto L67
L67:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v240 == int32(1) {
		goto L67
	} else {
		goto L69
	}
L68:
	;
	goto L65
L69:
	;
	goto L68
L70:
	;
	goto L71
L71:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v248 == int32(1) {
		goto L71
	} else {
		goto L73
	}
L72:
	;
	goto L61
L73:
	;
	goto L72
L74:
	;
	v512 = int32(_a69)
	v513 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v514)+4)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v514)+80)) = v513
	v519 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v520 = F_listAddNodeTail(m, v519, l0)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L36
	} else {
		goto L117
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v508))) = int32(9)
	goto L74
L76:
	;
	v260 = int64(56)
	v262 = int64(65280)
	v264 = int64(40)
	v267 = int64(16711680)
	v269 = int64(24)
	v271 = int64(4278190080)
	v273 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+176)) = v257<<(uint(v260)%64) | v257&v262<<(uint(v264)%64) | (v257&v267<<(uint(v269)%64) | v257&v271<<(uint(v273)%64)) | (int64(base.Ui64(v257)>>(uint(v273)%64))&v271 | int64(base.Ui64(v257)>>(uint(v269)%64))&v267 | (int64(base.Ui64(v257)>>(uint(v264)%64))&v262 | int64(base.Ui64(v257)>>(uint(v260)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+312)) = int32(0)
	v299 = *(*int32)(unsafe.Add(mBase, _consts[505]))
	v301 = v11 + int32(176)
	v302 = int32(8)
	v304 = v11 + int32(312)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	goto L80
L77:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v11)+312))
	if v501 == int32(0) {
		v508 = v500
		goto L75
	} else {
		goto L115
	}
L78:
	;
	if v457 != v302 {
		goto L105
	} else {
		goto L106
	}
L79:
	;
	v448 = int32(0)
	v454 = v313
	v455 = v314
	v457 = v448
	v461 = v448
	goto L78
L80:
	;
	if base.Ui32(v314) < base.Ui32(int32(8)) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v325 = v313
	v326 = v314
	v328 = int32(0)
	goto L83
L82:
	;
	v454 = v438
	v455 = v439
	v457 = v441
	v461 = base.B2i32(v444 != int32(0))
	goto L78
L83:
	;
	v334 = int32(base.Ui32(v326) >> (uint(int32(3)) % 32))
	v335 = int32(4)
	v336 = v325 + v335
	if v326&v335 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v438 = v429
	v439 = v430
	v441 = v414
	v444 = v419
	goto L82
L85:
	;
	v419 = int32(0)
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v336+v334+(v419-v334)&int32(3)+v407<<(uint(int32(2))%32))))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	if base.Ui32(v430) < base.Ui32(int32(8)) {
		v438 = v429
		v439 = v430
		v441 = v414
		v444 = v419
		goto L82
	} else {
		goto L103
	}
L86:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301+v328))))
	v385 = int32(0)
	goto L97
L87:
	;
	v341 = int32(0)
	if base.Ui32(v302) <= base.Ui32(v328) {
		v374 = v328
		v377 = v341
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v377 == v334 {
		v407 = v341
		v414 = v374
		goto L85
	} else {
		goto L95
	}
L89:
	;
	v351 = v328
	v354 = v341
	goto L90
L90:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336+v354))))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301+v351))))
	if v357 != v359 {
		v374 = v351
		v377 = v354
		goto L88
	} else {
		goto L92
	}
L91:
	;
	v374 = v362
	v377 = v364
	goto L88
L92:
	;
	v361 = int32(1)
	v362 = v351 + v361
	v364 = v354 + v361
	if base.Ui32(v334) <= base.Ui32(v364) {
		v374 = v362
		v377 = v364
		goto L88
	} else {
		goto L93
	}
L93:
	;
	if base.Ui32(v362) < base.Ui32(v302) {
		v351 = v362
		v354 = v364
		goto L90
	} else {
		goto L94
	}
L94:
	;
	goto L91
L95:
	;
	v438 = v325
	v439 = v326
	v441 = v374
	v444 = v377
	goto L82
L96:
	;
	if v385 != v334 {
		goto L101
	} else {
		goto L102
	}
L97:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336+v385))))
	if v398 == v382&int32(255) {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v400 = int32(1)
	v402 = v385 + v400
	if v402 != v334 {
		v385 = v402
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v454 = v325
	v455 = v326
	v457 = v328
	v461 = v400
	goto L78
L101:
	;
	v407 = v385
	v414 = v328 + int32(1)
	goto L85
L102:
	;
	v438 = v325
	v439 = v326
	v441 = v328
	v444 = v334
	goto L82
L103:
	;
	if base.Ui32(v414) < base.Ui32(v302) {
		v325 = v429
		v326 = v430
		v328 = v414
		goto L83
	} else {
		goto L104
	}
L104:
	;
	goto L84
L105:
	;
	goto L77
L106:
	;
	if v455&int32(1) == int32(0) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v469 = v455 & int32(4)
	if v461&base.B2i32(v469 != int32(0)) != 0 {
		goto L105
	} else {
		goto L108
	}
L108:
	;
	if v304 == int32(0) {
		goto L105
	} else {
		goto L109
	}
L109:
	;
	if v455&int32(2) != 0 {
		v495 = int32(0)
		goto L110
	} else {
		goto L111
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v304))) = v495
	goto L105
L111:
	;
	v479 = int32(3)
	v480 = int32(base.Ui32(v455) >> (uint(v479) % 32))
	if v469 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v490 = int32(4)
	goto L114
L113:
	;
	v490 = v480 << (uint(int32(2)) % 32)
	goto L114
L114:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v454+v480+(int32(0)-v480)&v479+v490+int32(4))))
	v495 = v494
	goto L110
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v500))) = int32(11)
	F_removeReplicaFromPsyncWait(m, l0)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L36
	} else {
		goto L116
	}
L116:
	;
	goto L74
L117:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522)+160)))
	if v523&int32(2) == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+68))
	v553 = m.T0[v552].(func(*base.Module, int32, int32, int32) int32)(m, v548, v11+int32(176), v547)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L36
	} else {
		goto L123
	}
L119:
	;
	v540 = int32(0)
	v541 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+184)) = v541
	v544 = *(*int64)(unsafe.Add(mBase, _consts[510]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+176)) = v544
	v547 = int32(11)
	goto L118
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = int32(_a920)
	v538 = F_snprintf(m, v11+int32(176), int32(128), int32(_a928), v11+int32(112))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L36
	} else {
		goto L121
	}
L121:
	;
	v547 = v538
	goto L118
L122:
	;
	v559 = F_addReplyReplicationBacklog(m, l0, l1)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L36
	} else {
		goto L126
	}
L123:
	;
	if v553 == v547 {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	F_freeClientAsync(m, l0)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L36
	} else {
		goto L125
	}
L125:
	;
	v677 = int32(0)
	goto L1
L126:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v562 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v580 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	if v580 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v565 = F_replicationGetReplicaName(m, l0)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L36
	} else {
		goto L129
	}
L129:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(96)))) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v565
	F__serverLog(m, int32(2), int32(_a929), v11+int32(80))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L36
	} else {
		goto L130
	}
L130:
	;
	goto L127
L131:
	;
	v666 = int32(0)
	F_moduleFireServerEvent(m, int64(6), v666, v666)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L36
	} else {
		goto L148
	}
L132:
	;
	v584 = *(*int32)(unsafe.Add(mBase, _consts[187]))
	if v584 == int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v588 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v590 = v11 + int32(312)
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v588)))
	*(*int32)(unsafe.Add(mBase, uint32(v590)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v590))) = v591
	goto L134
L134:
	;
	v595 = int32(0)
	v597 = v11 + int32(312)
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v597)))
	if v599 == v595 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, _consts[189])) = v651
	goto L131
L136:
	;
	if v599 == int32(0) {
		v651 = v595
		goto L135
	} else {
		goto L139
	}
L137:
	;
	goto L136
L138:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v599+base.B2i32(v602 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v597))) = v608
	goto L137
L139:
	;
	v612 = v599
	v615 = v595
	goto L140
L140:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v612)+8))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v620)+104))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v621)))
	if v622 != int32(9) {
		v633 = v615
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v651 = v633
	goto L135
L142:
	;
	v635 = v11 + int32(312)
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	if v637 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v625 = int32(_a69)
	v626 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	v627 = *(*int64)(unsafe.Add(mBase, uint32(v621)+80))
	v630 = int64(*(*int32)(unsafe.Add(mBase, _consts[187])))
	v633 = v615 + base.B2i32(v626-v627 <= v630)
	goto L142
L144:
	;
	if v637 != 0 {
		v612 = v637
		v615 = v633
		goto L140
	} else {
		goto L147
	}
L145:
	;
	goto L144
L146:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v637+base.B2i32(v640 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v635))) = v646
	goto L145
L147:
	;
	goto L141
L148:
	;
	v677 = v666
	goto L1
}
func F_proceedWithSlotMigration(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
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
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
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
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int64
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v309 int64
	_ = v309
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int64
	_ = v403
	var v404 int64
	_ = v404
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int64
	_ = v432
	var v433 int64
	_ = v433
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	v7 = m.G0
	v9 = v7 - int32(224)
	m.G0 = v9
	goto L9
L1:
	;
	F__serverAssert(m, int32(_a347), int32(_a325), int32(1406))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L16
	} else {
		goto L145
	}
L2:
	;
	F__serverAssert(m, int32(_a348), int32(_a325), int32(1346))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L16
	} else {
		goto L144
	}
L3:
	;
	m.G0 = v9 + int32(224)
	return
L4:
	;
	v432 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	v433 = F_mstime(m)
	mBase = m.M
	if v432 < v433 {
		goto L130
	} else {
		goto L131
	}
L5:
	;
	v403 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	v404 = F_mstime(m)
	mBase = m.M
	if v403 < v404 {
		goto L120
	} else {
		goto L121
	}
L6:
	;
	v398 = F_slotExportTryDoPause(m, l0)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L16
	} else {
		goto L117
	}
L7:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	goto L92
L8:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v238 != 0 {
		goto L1
	} else {
		goto L75
	}
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	switch v17 {
	case 0, 1, 2, 3, 6, 9, 11, 13, 14, 18, 19, 20:
		goto L3
	case 4:
		goto L14
	case 5:
		goto L13
	case 7:
		goto L12
	case 8:
		goto L11
	case 10:
		goto L8
	case 12:
		goto L7
	case 15:
		goto L6
	case 16:
		goto L5
	case 17:
		goto L4
	default:
		goto L9
	}
L11:
	;
	F_slotMigrationJobSendAuth(m, l0)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L16
	} else {
		goto L74
	}
L12:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v71 != 0 {
		goto L31
	} else {
		goto L32
	}
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v39 {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _consts[123])))
	if v19&int32(16) != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	F_performSlotImportJobFailover(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v25 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_finishSlotMigrationJob(m, l0, int32(20), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v28
	F__serverLog(m, int32(2), int32(_a349), v9)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L3
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_delKeysNotOwnedByMyself(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L16
	} else {
		goto L27
	}
L23:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if base.Ui32(int32(20)) < base.Ui32(v44) {
		v52 = int32(_a242)
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v42
	F__serverLog(m, int32(2), int32(_a350), v9+int32(16))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L16
	} else {
		goto L26
	}
L25:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44<<(uint(int32(2))%32))+uint32(_consts[145])))
	v52 = v51
	goto L24
L26:
	;
	goto L22
L27:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	F_finishSlotMigrationJob(m, l0, v67, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	goto L3
L29:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v164 {
		goto L60
	} else {
		goto L61
	}
L30:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v143 != 0 {
		goto L53
	} else {
		goto L54
	}
L31:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v134 != 0 {
		goto L2
	} else {
		goto L49
	}
L32:
	;
	v73 = l0 + int32(32)
	v75 = F_clusterLookupNode(m, v73, int32(40))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L16
	} else {
		goto L34
	}
L33:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v93 != 0 {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	if v75 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v78 {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v81
	F__serverLog(m, int32(3), int32(_a351), v9+int32(48))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L16
	} else {
		goto L37
	}
L37:
	;
	goto L30
L38:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v98 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v94 = int32(2328)
	goto L41
L40:
	;
	v94 = int32(2324)
	goto L41
L41:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v75+v94)))
	goto L38
L42:
	;
	v114 = F_connTypeOfReplication(m)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L16
	} else {
		goto L45
	}
L43:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v75 + int32(2256)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v101
	F__serverLog(m, int32(2), int32(_a352), v9+int32(64))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)+40))
	v117 = m.T0[v116].(func(*base.Module) int32)(m)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v117
	v123 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+56))
	v128 = m.T0[v127].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v117, v75+int32(2256), v96, v123, int32(0), int32(104))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L16
	} else {
		goto L47
	}
L47:
	;
	if v128 == int32(-1) {
		goto L30
	} else {
		goto L48
	}
L48:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+24)) = l0
	goto L3
L49:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v135 == int32(3) {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	if v135 == int32(1) {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	goto L30
L52:
	;
	v150 = F_sdsempty(m)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L16
	} else {
		goto L56
	}
L53:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+88))
	v147 = m.T0[v146].(func(*base.Module, int32) int32)(m, v143)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L16
	} else {
		goto L55
	}
L54:
	;
	v149 = int32(_a353)
	goto L52
L55:
	;
	v149 = v147
	goto L52
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v149
	v157 = F_sdscatfmt(m, v150, int32(_a354), v9+int32(32))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	F_finishSlotMigrationJob(m, l0, int32(18), v157)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L16
	} else {
		goto L58
	}
L58:
	;
	F_sdsfree(m, v157)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L16
	} else {
		goto L59
	}
L59:
	;
	goto L3
L60:
	;
	v175 = int32(_a69)
	v176 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v178 = *(*int32)(unsafe.Add(mBase, _consts[147]))
	if v178 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v167
	F__serverLog(m, int32(2), int32(_a355), v9+int32(112))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L16
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v233 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v229
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v233
	goto L9
L64:
	;
	v205 = int32(10)
	if int32(2) < v176 {
		v229 = v205
		goto L63
	} else {
		goto L70
	}
L65:
	;
	v181 = int32(8)
	if int32(2) < v176 {
		v229 = v181
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if base.Ui32(int32(20)) < base.Ui32(v185) {
		v193 = int32(_a242)
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+104)) = int32(_a356)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+100)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v194
	F__serverLog(m, int32(2), int32(_a330), v9+int32(96))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L16
	} else {
		goto L69
	}
L68:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v185<<(uint(int32(2))%32))+uint32(_consts[145])))
	v193 = v192
	goto L67
L69:
	;
	v229 = v181
	goto L63
L70:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if base.Ui32(int32(20)) < base.Ui32(v209) {
		v217 = int32(_a242)
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+88)) = int32(_a357)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v218
	F__serverLog(m, int32(2), int32(_a330), v9+int32(80))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L16
	} else {
		goto L73
	}
L72:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v209<<(uint(int32(2))%32))+uint32(_consts[145])))
	v217 = v216
	goto L71
L73:
	;
	v229 = v205
	goto L63
L74:
	;
	goto L9
L75:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v240 = F_createClient(m, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L16
	} else {
		goto L76
	}
L76:
	;
	v242 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v240)+328)) = v242
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v240)+204))
	goto L78
L77:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+216)) = l0
	F_initClientReplicationData(m, v264)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L16
	} else {
		goto L81
	}
L78:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+204)) = v249&int32(-25165825) | int32(_a14) | int32(16777216)
	goto L77
L81:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v269 = F_generateSyncSlotsEstablishCommand(m, l0)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L16
	} else {
		goto L82
	}
L82:
	;
	F_addReplySds(m, v268, v269)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L16
	} else {
		goto L83
	}
L83:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+8))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+84))
	v278 = m.T0[v277].(func(*base.Module, int32, int32) int32)(m, v274, int32(105))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L16
	} else {
		goto L84
	}
L84:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v281 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v309 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(11)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v309
	goto L3
L86:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if base.Ui32(int32(20)) < base.Ui32(v286) {
		v294 = int32(_a242)
		goto L87
	} else {
		goto L88
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+136)) = int32(_a358)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+132)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v284
	F__serverLog(m, int32(2), int32(_a330), v9+int32(128))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L16
	} else {
		goto L89
	}
L88:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v286<<(uint(int32(2))%32))+uint32(_consts[145])))
	v294 = v293
	goto L87
L89:
	;
	goto L85
L90:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v364 {
		goto L106
	} else {
		goto L107
	}
L91:
	;
	v326 = int32(_a69)
	v327 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	v331 = *(*int32)(unsafe.Add(mBase, _consts[149]))
	v332 = base.I32_div_s(int32(1000), v331)
	v334 = base.I32_div_s(int32(5000), base.I32_extend16_s(v332))
	v336 = base.I32_rem_s(v327, base.I32_extend16_s(v334))
	if v336 != 0 {
		goto L3
	} else {
		goto L97
	}
L92:
	;
	if v314 != int32(-1) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+202)))
	if v318&int32(64) != 0 {
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+223)))
	if v321 != 0 {
		goto L91
	} else {
		goto L95
	}
L95:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+222)))
	if v322 == int32(0) {
		goto L90
	} else {
		goto L96
	}
L96:
	;
	goto L91
L97:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v338 {
		goto L3
	} else {
		goto L98
	}
L98:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v344 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	goto L100
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+180)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v341
	F__serverLog(m, int32(2), int32(_a359), v9+int32(176))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L16
	} else {
		goto L105
	}
L100:
	;
	if v344 != int32(-1) {
		v354 = int32(_a360)
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+202)))
	if v350&int32(64) != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v353 = int32(_a361)
	goto L104
L103:
	;
	v353 = int32(_a362)
	goto L104
L104:
	;
	v354 = v353
	goto L99
L105:
	;
	goto L3
L106:
	;
	v375 = F_slotExportJobBeginSnapshotToTargetSocket(m, l0)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L16
	} else {
		goto L110
	}
L107:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v367
	F__serverLog(m, int32(2), int32(_a363), v9+int32(160))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L16
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	F_updateSlotMigrationJobState(m, l0, int32(13))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L16
	} else {
		goto L116
	}
L110:
	;
	if v375 != int32(-1) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v380 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_finishSlotMigrationJob(m, l0, int32(18), int32(_a364))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L16
	} else {
		goto L115
	}
L113:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v383
	F__serverLog(m, int32(3), int32(_a365), v9+int32(144))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L16
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	goto L3
L116:
	;
	goto L3
L117:
	;
	if v398 != 0 {
		goto L3
	} else {
		goto L118
	}
L118:
	;
	F_updateSlotMigrationJobState(m, l0, int32(16))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L16
	} else {
		goto L119
	}
L119:
	;
	goto L3
L120:
	;
	v413 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v413 {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	v411 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	goto L122
L122:
	;
	if v411 != 0 {
		goto L3
	} else {
		goto L123
	}
L123:
	;
	goto L120
L124:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = int64(0)
	F_updatePausedActions(m)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L16
	} else {
		goto L127
	}
L125:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+192)) = v416
	F__serverLog(m, int32(3), int32(_a366), v9+int32(192))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L16
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	F_finishSlotMigrationJob(m, l0, int32(18), int32(_a367))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L16
	} else {
		goto L128
	}
L128:
	;
	goto L3
L129:
	;
	v459 = F_checkSlotExportOwnership(m, l0, v9+int32(223))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L16
	} else {
		goto L139
	}
L130:
	;
	v442 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v442 {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	v440 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	goto L132
L132:
	;
	if v440 != 0 {
		goto L129
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
	;
	F_finishSlotMigrationJob(m, l0, int32(18), int32(_a368))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L16
	} else {
		goto L137
	}
L135:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+208)) = v445
	F__serverLog(m, int32(3), int32(_a369), v9+int32(208))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L16
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	goto L3
L138:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+223)))
	if v467 != int32(1) {
		goto L3
	} else {
		goto L142
	}
L139:
	;
	if v459 != int32(-1) {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	F_finishSlotMigrationJob(m, l0, int32(18), int32(_a370))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L16
	} else {
		goto L141
	}
L141:
	;
	goto L3
L142:
	;
	F_finishSlotMigrationJob(m, l0, int32(20), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L16
	} else {
		goto L143
	}
L143:
	;
	goto L3
L144:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_processCollectionElementEnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 < int32(0) {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = l0 + int32(1040)
		v12 = int32(2)
		v13 = v6 << (uint(v12) % 32)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11+v13)))
		if v15 != v12 {
			v30 = v15
			if v30 != int32(3) {
				v58 = v30
				if v58 != int32(1) {
					return
				} else {
					F_lua_settable(m, v9, int32(-3))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v34 = F_lua_checkstack(m, v9, int32(1))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					if v34 == int32(0) {
						F__serverPanic_2(m, int32(848))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(1)
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v46 + int32(16)
						F_lua_settable(m, v9, int32(-3))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v11+v53<<(uint(int32(2))%32))))
							v58 = v57
							if v58 != int32(1) {
								return
							} else {
								F_lua_settable(m, v9, int32(-3))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v13)+16)))
			if v19&int32(1) != 0 {
				return
			} else {
				F_lua_settable(m, v9, int32(-3))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v11+v25<<(uint(int32(2))%32))))
					v30 = v29
					if v30 != int32(3) {
						v58 = v30
						if v58 != int32(1) {
							return
						} else {
							F_lua_settable(m, v9, int32(-3))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v34 = F_lua_checkstack(m, v9, int32(1))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							if v34 == int32(0) {
								F__serverPanic_2(m, int32(848))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(1)
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v46 + int32(16)
								F_lua_settable(m, v9, int32(-3))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v11+v53<<(uint(int32(2))%32))))
									v58 = v57
									if v58 != int32(1) {
										return
									} else {
										F_lua_settable(m, v9, int32(-3))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
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
			}
		}
	}
}
func F_processEventsWhileBlocked(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v18 int64
	_ = v18
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v14 = F_ustime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v14
	v18 = base.I64_div_s(v14, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[32])) = v18
	v22 = base.I64_div_s(v14, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[37])) = v22
	v25 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	F_lrulfu_updateClockAndPolicy(m, v18, int32(base.Ui32(v25&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	m.G0 = v11 + int32(64)
	v47 = int32(_a69)
	v48 = *(*int64)(unsafe.Add(mBase, _consts[78]))
	v51 = *(*int64)(unsafe.Add(mBase, _consts[32]))
	*(*int64)(unsafe.Add(mBase, _consts[78])) = v51
	v53 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, _consts[435]))
	*(*int32)(unsafe.Add(mBase, _consts[435])) = v55 + int32(1)
	v60 = *(*int64)(unsafe.Add(mBase, _consts[436]))
	v62 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v64 = F_aeProcessEvents(m, v62, int32(29))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		return
	} else {
		v66 = int32(_a69)
		v68 = *(*int64)(unsafe.Add(mBase, _consts[436]))
		v70 = v68 + base.I64_extend_i32_s(v64)
		*(*int64)(unsafe.Add(mBase, _consts[436])) = v70
		if v70 == v60 {
			F_whileBlockedCron(m)
			mBase = m.M
			v112 = m.ExcPending
			if v112 != 0 {
				return
			} else {
				v113 = int32(0)
				v115 = *(*int32)(unsafe.Add(mBase, _consts[435]))
				*(*int32)(unsafe.Add(mBase, _consts[435])) = v115 + int32(-1)
				if v113 < v115 {
					*(*int64)(unsafe.Add(mBase, _consts[78])) = v48
					return
				} else {
					F__serverAssert(m, int32(_a816), int32(_a774), int32(6533))
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
		} else {
			v74 = *(*int32)(unsafe.Add(mBase, _consts[156]))
			v76 = F_aeProcessEvents(m, v74, int32(29))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return
			} else {
				v78 = int32(_a69)
				v80 = *(*int64)(unsafe.Add(mBase, _consts[436]))
				v82 = v80 + base.I64_extend_i32_s(v76)
				*(*int64)(unsafe.Add(mBase, _consts[436])) = v82
				if v82 == v70 {
					F_whileBlockedCron(m)
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						v113 = int32(0)
						v115 = *(*int32)(unsafe.Add(mBase, _consts[435]))
						*(*int32)(unsafe.Add(mBase, _consts[435])) = v115 + int32(-1)
						if v113 < v115 {
							*(*int64)(unsafe.Add(mBase, _consts[78])) = v48
							return
						} else {
							F__serverAssert(m, int32(_a816), int32(_a774), int32(6533))
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
				} else {
					v86 = *(*int32)(unsafe.Add(mBase, _consts[156]))
					v88 = F_aeProcessEvents(m, v86, int32(29))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						v90 = int32(_a69)
						v92 = *(*int64)(unsafe.Add(mBase, _consts[436]))
						v94 = v92 + base.I64_extend_i32_s(v88)
						*(*int64)(unsafe.Add(mBase, _consts[436])) = v94
						if v94 == v82 {
							F_whileBlockedCron(m)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return
							} else {
								v113 = int32(0)
								v115 = *(*int32)(unsafe.Add(mBase, _consts[435]))
								*(*int32)(unsafe.Add(mBase, _consts[435])) = v115 + int32(-1)
								if v113 < v115 {
									*(*int64)(unsafe.Add(mBase, _consts[78])) = v48
									return
								} else {
									F__serverAssert(m, int32(_a816), int32(_a774), int32(6533))
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
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, _consts[156]))
							v100 = F_aeProcessEvents(m, v98, int32(29))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return
							} else {
								v102 = int32(_a69)
								v104 = *(*int64)(unsafe.Add(mBase, _consts[436]))
								*(*int64)(unsafe.Add(mBase, _consts[436])) = v104 + base.I64_extend_i32_s(v100)
								F_whileBlockedCron(m)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return
								} else {
									v113 = int32(0)
									v115 = *(*int32)(unsafe.Add(mBase, _consts[435]))
									*(*int32)(unsafe.Add(mBase, _consts[435])) = v115 + int32(-1)
									if v113 < v115 {
										*(*int64)(unsafe.Add(mBase, _consts[78])) = v48
										return
									} else {
										F__serverAssert(m, int32(_a816), int32(_a774), int32(6533))
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
							}
						}
					}
				}
			}
		}
	}
}
func F_processInputBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	goto L2
L1:
	;
	return v194
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v15 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v194 = int32(-1)
	goto L1
L4:
	;
	v51 = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v52 != 0 {
		v194 = v51
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	if base.Ui32(v43) < base.Ui32(v44) {
		goto L4
	} else {
		goto L14
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(-1)))))
	switch v21 & int32(7) {
	case 0:
		goto L12
	case 1:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L9
	case 4:
		goto L8
	default:
		goto L5
	}
L7:
	;
	if base.Ui32(v18) < base.Ui32(v38) {
		goto L4
	} else {
		goto L13
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(-17))))
	v38 = v37
	goto L7
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(-9))))
	v38 = v34
	goto L7
L10:
	;
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+int32(-5)))))
	v38 = v31
	goto L7
L11:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(-3)))))
	v38 = v28
	goto L7
L12:
	;
	v38 = int32(base.Ui32(v21) >> (uint(int32(3)) % 32))
	goto L7
L13:
	;
	goto L5
L14:
	;
	return int32(0)
L15:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v53&int32(144) != 0 {
		v194 = v51
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
	if v56&int32(2) != 0 {
		v194 = v51
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v59 = F_scriptIsTimedout(m)
	mBase = m.M
	v60 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	goto L19
L18:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
	if v80&int32(1088) != 0 {
		v194 = v51
		goto L1
	} else {
		goto L27
	}
L19:
	;
	if base.B2i32(v59|v61 != v60) == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v68 = int32(1)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v69&v68 != 0 {
		v76 = v68
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v79 != 0 {
		v194 = v51
		goto L1
	} else {
		goto L26
	}
L22:
	;
	v79 = v76
	goto L21
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v72 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v74 = F_isImportSlotMigrationJob(m, v72)
	mBase = m.M
	v76 = v74
	goto L22
L25:
	;
	v79 = int32(0)
	goto L21
L26:
	;
	goto L18
L27:
	;
	v84 = int32(1)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v85&v84 != 0 {
		v92 = v84
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v96 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, _consts[414]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v102&int32(6) == int32(4) {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v95 = v92
	goto L28
L30:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v88 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v90 = F_isImportSlotMigrationJob(m, v88)
	mBase = m.M
	v92 = v90
	goto L29
L32:
	;
	v95 = int32(0)
	goto L28
L33:
	;
	v116 = v115 | base.B2i32(v95 != v96)<<(uint(int32(14))%32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v116
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	v120 = base.B2i32(base.Ui32(v119) <= base.Ui32(v118))
	if base.Ui32(v119) <= base.Ui32(v118) {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v115 = int32(0)
	goto L33
L35:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v108&int32(_a14) == int32(0) {
		v115 = int32(65536)
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	F_prefetchCommandQueueKeys(m, l0)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L41
	} else {
		goto L44
	}
L38:
	;
	F_parseInputBuffer(m, l0)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v122 = v118 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)) = uint16(v122)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v127 = v124 + v118*int32(40)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v128 | v116
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v127)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v137
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v127)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+240)) = v139
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v127)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v141
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v143
	if v122&int32(65535) != v119 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
	goto L37
L41:
	;
	return int32(0)
L42:
	;
	F_prepareCommandQueue(m, l0)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	goto L37
L44:
	;
	v160 = F_handleParseResults(m, l0)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	if base.Ui32(v119) <= base.Ui32(v118) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if v160 != 0 {
		v194 = v51
		goto L1
	} else {
		goto L49
	}
L47:
	;
	if v160 == int32(-2) {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v164 == int32(0) {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v169 = *(*int32)(unsafe.Add(mBase, _consts[407]))
	if v167 != v169 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v173 = int32(_a69)
	v174 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	*(*int32)(unsafe.Add(mBase, _consts[88])) = l0
	v177 = F_processCommand(m, l0)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L41
	} else {
		goto L55
	}
L52:
	;
	F_resetSharedQueryBuf(m, l0)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L41
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v186 = int32(_a69)
	v187 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v174
	if v187 != 0 {
		goto L2
	} else {
		goto L60
	}
L55:
	;
	if v177 != 0 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	F_commandProcessed(m, l0)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L41
	} else {
		goto L57
	}
L57:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v181 == int32(0) {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v184 = F_updateClientMemUsageAndBucket(m, l0)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L41
	} else {
		goto L59
	}
L59:
	;
	goto L54
L60:
	;
	goto L3
}
func F_processUnblockedClients(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v5 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v9 = v5
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v12 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L1
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v44 != 0 {
		v9 = v43
		goto L3
	} else {
		goto L20
	}
L6:
	;
	F__serverAssert(m, int32(_a133), int32(_a134), int32(164))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L19
	}
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	F_listDelNode(m, v9, v12)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+200)) = v18 & int32(-129)
	v23 = v18 & int32(16)
	if v18&int32(1073741824) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if v23 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	if v23 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	F_moduleCallCommandUnblockedHandler(m, v15)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L5
L14:
	;
	F_beforeNextClient(m, v15)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L18
	}
L15:
	;
	v30 = F_processPendingCommandAndInputBuffer(m, v15)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	if v30 == int32(-1) {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	goto L5
L19:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	goto L4
}
func F_propagateFieldsDeletion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(4112)
	m.G0 = v11
	v13 = int32(_a69)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[235]))
	*(*int32)(unsafe.Add(mBase, _consts[235])) = int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v21&int32(2) == v6 {
		v41 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v42 = F_createStringObjectFromSds(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L1
L3:
	;
	v35 = l1 + (v21&int32(4) ^ int32(12)) + v21<<(uint(int32(3))%32)&int32(8)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v41 = v35 + v36 + int32(1)
	goto L2
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v42
	v48 = *(*int32)(unsafe.Add(mBase, _consts[236]))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v48
	v50 = int32(1024)
	if base.Ui32(l2) < base.Ui32(v50) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v53 = l2
	goto L8
L7:
	;
	v53 = v50
	goto L8
L8:
	;
	if l2 == int32(0) {
		v67 = int32(2)
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_alsoPropagate(m, v68, v11, v67, int32(3), l4)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L14
	}
L10:
	;
	v60 = v53 << (uint(int32(2)) % 32)
	if v60 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v67 = v53 + int32(2)
	goto L9
L12:
	;
	goto L11
L13:
	;
	v63 = F__emscripten_memcpy_bulkmem(m, v11|int32(8), l3, v60)
	mBase = m.M
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[235])) = v14
	v77 = int32(0)
	goto L15
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v11+v77<<(uint(int32(2))%32))))
	F_decrRefCount(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	m.G0 = v11 + int32(4112)
	return v53
L17:
	;
	v90 = v77 + int32(1)
	if v90 != v67 {
		v77 = v90
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
}
func F_propagatemark(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
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
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v740 int32
	_ = v740
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1068 int32
	_ = v1068
	var v1076 int32
	_ = v1076
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1242 int32
	_ = v1242
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1253 int32
	_ = v1253
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1276 int32
	_ = v1276
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1413 int32
	_ = v1413
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1432 int32
	_ = v1432
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+5)))
	v11 = v9 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+5)) = uint8(v11)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
	switch v14 + int32(-5) {
	case 0:
		goto L5
	case 1:
		goto L4
	default:
		v1432 = int32(0)
		goto L1
	case 3:
		goto L3
	case 4:
		goto L2
	}
L1:
	;
	return v1432
L2:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v8)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1129
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	if v1131 == int32(0) {
		goto L314
	} else {
		goto L315
	}
L3:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v8)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v857
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+108)) = v859
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v8
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+5)))
	v864 = v862 & int32(251)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+5)) = uint8(v864)
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
	if v866 < int32(4) {
		goto L240
	} else {
		goto L241
	}
L4:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v490
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+5)))
	if v493&int32(3) == int32(0) {
		goto L136
	} else {
		goto L137
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v17
	v19 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v20 == v19 {
		v169 = v19
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)))
	return v480<<(uint(int32(4))%32) + int32(32)<<(uint(v484)%32) + int32(40)
L7:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+5)))
	v471 = v469 & int32(251)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+5)) = uint8(v471)
	goto L6
L8:
	;
	v280 = int32(-1)
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)))
	v289 = v280<<(uint(v281)%32) ^ v280
	goto L81
L9:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	if v172 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+5)))
	if v23&int32(3) == int32(0) {
		v99 = v20
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+6)))
	if v102&int32(8) != 0 {
		v169 = int32(0)
		goto L9
	} else {
		goto L34
	}
L12:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+5)))
	v32 = v20
	v33 = v30
	goto L21
L13:
	;
	v95 = int32(0)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v96 == v95 {
		v169 = v95
		goto L9
	} else {
		goto L33
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v32
	goto L14
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v86
	goto L15
L17:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v84
	goto L15
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v82
	goto L15
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v80
	goto L15
L20:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v60 < int32(4) {
		v71 = v59
		goto L29
	} else {
		goto L30
	}
L21:
	;
	v36 = v33 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+5)) = uint8(v36)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
	if v38 == int32(7) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v44 = v36 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+5)) = uint8(v44)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v46 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	switch v38 + int32(-5) {
	case 0:
		goto L18
	case 1:
		goto L19
	default:
		goto L14
	case 3:
		goto L17
	case 4:
		goto L16
	case 5:
		goto L20
	}
L25:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+5)))
	if v56&int32(3) != 0 {
		v32 = v55
		v33 = v56
		goto L21
	} else {
		goto L28
	}
L26:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+5)))
	if v49&int32(3) == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	F_reallymarkobject(m, l0, v46)
	mBase = m.M
	goto L25
L28:
	;
	goto L14
L29:
	;
	if v71 != v32+int32(16) {
		goto L14
	} else {
		goto L32
	}
L30:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
	if v64&int32(3) == int32(0) {
		v71 = v59
		goto L29
	} else {
		goto L31
	}
L31:
	;
	F_reallymarkobject(m, l0, v63)
	mBase = m.M
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v71 = v70
	goto L29
L32:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+5)))
	v78 = v76 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+5)) = uint8(v78)
	goto L13
L33:
	;
	v99 = v96
	goto L11
L34:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v108 = F_luaH_getstr(m, v99, v107)
	mBase = m.M
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	if v109 != 0 {
		v116 = v108
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v116 == int32(0) {
		v169 = int32(0)
		goto L9
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+6)))
	v113 = v110 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+6)) = uint8(v113)
	v116 = int32(0)
	goto L36
L38:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	if v120 != int32(4) {
		v169 = int32(0)
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v123 = int32(0)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v126 = v124 + int32(16)
	v127 = int32(107)
	v128 = F___strchrnul(m, v126, v127)
	mBase = m.M
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v130 == v127 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v135 = int32(118)
	v136 = F___strchrnul(m, v126, v135)
	mBase = m.M
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v138 == v135 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v134 = v128
	goto L43
L42:
	;
	v134 = v123
	goto L43
L43:
	;
	goto L40
L44:
	;
	if v134|v142 == int32(0) {
		v169 = v123
		goto L9
	} else {
		goto L48
	}
L45:
	;
	v142 = v136
	goto L47
L46:
	;
	v142 = int32(0)
	goto L47
L47:
	;
	goto L44
L48:
	;
	v146 = int32(0)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+5)))
	v158 = base.B2i32(v142 != v146)<<(uint(int32(4))%32) | base.B2i32(v134 != v146)<<(uint(int32(3))%32) | v155&int32(231)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+5)) = uint8(v158)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v8
	if v134 == v146 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v166 = base.B2i32(v134 != int32(0))
	if v142 != 0 {
		v276 = v166
		v278 = int32(1)
		goto L8
	} else {
		goto L52
	}
L50:
	;
	if v142 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v169 = v166
	goto L9
L53:
	;
	v276 = v169
	v278 = int32(0)
	goto L8
L54:
	;
	v177 = v172
	goto L55
L55:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v184 = v177 + int32(-1)
	v185 = int32(4)
	v187 = v182 + v184<<(uint(v185)%32)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	if v188 < v185 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L53
L57:
	;
	if v184 != 0 {
		v177 = v184
		goto L55
	} else {
		goto L80
	}
L58:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+5)))
	if v192&int32(3) == int32(0) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+5)))
	v201 = v191
	v202 = v199
	goto L68
L60:
	;
	goto L57
L61:
	;
	goto L60
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v201
	goto L61
L63:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v201)+68)) = v255
	goto L62
L64:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v201)+108)) = v253
	goto L62
L65:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v201)+32)) = v251
	goto L62
L66:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v201)+8)) = v249
	goto L62
L67:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+8))
	if v229 < int32(4) {
		v240 = v228
		goto L76
	} else {
		goto L77
	}
L68:
	;
	v205 = v202 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+5)) = uint8(v205)
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+4)))
	if v207 == int32(7) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v213 = v205 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+5)) = uint8(v213)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	if v215 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	switch v207 + int32(-5) {
	case 0:
		goto L65
	case 1:
		goto L66
	default:
		goto L61
	case 3:
		goto L64
	case 4:
		goto L63
	case 5:
		goto L67
	}
L72:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+5)))
	if v225&int32(3) != 0 {
		v201 = v224
		v202 = v225
		goto L68
	} else {
		goto L75
	}
L73:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+5)))
	if v218&int32(3) == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	F_reallymarkobject(m, l0, v215)
	mBase = m.M
	goto L72
L75:
	;
	goto L61
L76:
	;
	if v240 != v201+int32(16) {
		goto L61
	} else {
		goto L79
	}
L77:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+5)))
	if v233&int32(3) == int32(0) {
		v240 = v228
		goto L76
	} else {
		goto L78
	}
L78:
	;
	F_reallymarkobject(m, l0, v232)
	mBase = m.M
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	v240 = v239
	goto L76
L79:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+5)))
	v247 = v245 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+5)) = uint8(v247)
	goto L60
L80:
	;
	goto L56
L81:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	v295 = v292 + v289<<(uint(int32(5))%32)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+8))
	if v296 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	if v276|v278 != int32(1) {
		goto L6
	} else {
		goto L135
	}
L83:
	;
	if v289 != 0 {
		v289 = v289 + int32(-1)
		goto L81
	} else {
		goto L134
	}
L84:
	;
	if v276 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v295)+24))
	if v297 < int32(4) {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+24)) = int32(11)
	goto L83
L87:
	;
	if v278 != 0 {
		goto L83
	} else {
		goto L111
	}
L88:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v295)+24))
	if v302 < int32(4) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+5)))
	if v306&int32(3) == int32(0) {
		goto L87
	} else {
		goto L90
	}
L90:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+5)))
	v315 = v305
	v316 = v313
	goto L99
L91:
	;
	goto L87
L92:
	;
	goto L91
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v315
	goto L92
L94:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+68)) = v369
	goto L93
L95:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+108)) = v367
	goto L93
L96:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+32)) = v365
	goto L93
L97:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+8)) = v363
	goto L93
L98:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	if v343 < int32(4) {
		v354 = v342
		goto L107
	} else {
		goto L108
	}
L99:
	;
	v319 = v316 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v315)+5)) = uint8(v319)
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+4)))
	if v321 == int32(7) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v327 = v319 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v315)+5)) = uint8(v327)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	if v329 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	switch v321 + int32(-5) {
	case 0:
		goto L96
	case 1:
		goto L97
	default:
		goto L92
	case 3:
		goto L95
	case 4:
		goto L94
	case 5:
		goto L98
	}
L103:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v315)+12))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+5)))
	if v339&int32(3) != 0 {
		v315 = v338
		v316 = v339
		goto L99
	} else {
		goto L106
	}
L104:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+5)))
	if v332&int32(3) == int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	F_reallymarkobject(m, l0, v329)
	mBase = m.M
	goto L103
L106:
	;
	goto L92
L107:
	;
	if v354 != v315+int32(16) {
		goto L92
	} else {
		goto L110
	}
L108:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+5)))
	if v347&int32(3) == int32(0) {
		v354 = v342
		goto L107
	} else {
		goto L109
	}
L109:
	;
	F_reallymarkobject(m, l0, v346)
	mBase = m.M
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	v354 = v353
	goto L107
L110:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+5)))
	v361 = v359 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v315)+5)) = uint8(v361)
	goto L91
L111:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v295)+8))
	if v379 < int32(4) {
		goto L83
	} else {
		goto L112
	}
L112:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+5)))
	if v383&int32(3) == int32(0) {
		goto L83
	} else {
		goto L113
	}
L113:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+5)))
	v392 = v382
	v393 = v390
	goto L122
L114:
	;
	goto L83
L115:
	;
	goto L114
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v392
	goto L115
L117:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v392)+68)) = v446
	goto L116
L118:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v392)+108)) = v444
	goto L116
L119:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v392)+32)) = v442
	goto L116
L120:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v392)+8)) = v440
	goto L116
L121:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v392)+8))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	if v420 < int32(4) {
		v431 = v419
		goto L130
	} else {
		goto L131
	}
L122:
	;
	v396 = v393 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v392)+5)) = uint8(v396)
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+4)))
	if v398 == int32(7) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v404 = v396 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v392)+5)) = uint8(v404)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v392)+8))
	if v406 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	switch v398 + int32(-5) {
	case 0:
		goto L119
	case 1:
		goto L120
	default:
		goto L115
	case 3:
		goto L118
	case 4:
		goto L117
	case 5:
		goto L121
	}
L126:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v392)+12))
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+5)))
	if v416&int32(3) != 0 {
		v392 = v415
		v393 = v416
		goto L122
	} else {
		goto L129
	}
L127:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+5)))
	if v409&int32(3) == int32(0) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	F_reallymarkobject(m, l0, v406)
	mBase = m.M
	goto L126
L129:
	;
	goto L115
L130:
	;
	if v431 != v392+int32(16) {
		goto L115
	} else {
		goto L133
	}
L131:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+5)))
	if v424&int32(3) == int32(0) {
		v431 = v419
		goto L130
	} else {
		goto L132
	}
L132:
	;
	F_reallymarkobject(m, l0, v423)
	mBase = m.M
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v392)+8))
	v431 = v430
	goto L130
L133:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+5)))
	v438 = v436 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v392)+5)) = uint8(v438)
	goto L114
L134:
	;
	goto L82
L135:
	;
	goto L7
L136:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+6)))
	if v565 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L137:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+5)))
	v502 = v492
	v503 = v500
	goto L146
L138:
	;
	goto L136
L139:
	;
	goto L138
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v502
	goto L139
L141:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v502)+68)) = v556
	goto L140
L142:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v502)+108)) = v554
	goto L140
L143:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v502)+32)) = v552
	goto L140
L144:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v502)+8)) = v550
	goto L140
L145:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v502)+8))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)+8))
	if v530 < int32(4) {
		v541 = v529
		goto L154
	} else {
		goto L155
	}
L146:
	;
	v506 = v503 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v502)+5)) = uint8(v506)
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+4)))
	if v508 == int32(7) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v514 = v506 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v502)+5)) = uint8(v514)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v502)+8))
	if v516 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	switch v508 + int32(-5) {
	case 0:
		goto L143
	case 1:
		goto L144
	default:
		goto L139
	case 3:
		goto L142
	case 4:
		goto L141
	case 5:
		goto L145
	}
L150:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v502)+12))
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525)+5)))
	if v526&int32(3) != 0 {
		v502 = v525
		v503 = v526
		goto L146
	} else {
		goto L153
	}
L151:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+5)))
	if v519&int32(3) == int32(0) {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	F_reallymarkobject(m, l0, v516)
	mBase = m.M
	goto L150
L153:
	;
	goto L139
L154:
	;
	if v541 != v502+int32(16) {
		goto L139
	} else {
		goto L157
	}
L155:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+5)))
	if v534&int32(3) == int32(0) {
		v541 = v529
		goto L154
	} else {
		goto L156
	}
L156:
	;
	F_reallymarkobject(m, l0, v533)
	mBase = m.M
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v502)+8))
	v541 = v540
	goto L154
L157:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+5)))
	v548 = v546 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v502)+5)) = uint8(v548)
	goto L138
L158:
	;
	v843 = v839 & int32(255)
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+6)))
	if v844 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L159:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+5)))
	if v668&int32(3) == int32(0) {
		goto L189
	} else {
		goto L190
	}
L160:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)))
	if v568 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v575 = int32(0)
	v577 = v568
	goto L163
L162:
	;
	v839 = int32(0)
	goto L158
L163:
	;
	v580 = int32(4)
	v582 = v8 + int32(24) + v575<<(uint(v580)%32)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v582)+8))
	if v583 < v580 {
		v661 = v577
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v663 = v575 + int32(1)
	if base.Ui32(v663) < base.Ui32(v661&int32(255)) {
		v575 = v663
		v577 = v661
		goto L163
	} else {
		goto L188
	}
L166:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v582)))
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+5)))
	if v587&int32(3) == int32(0) {
		v661 = v577
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+5)))
	v596 = v586
	v597 = v594
	goto L176
L168:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)))
	v661 = v659
	goto L165
L169:
	;
	goto L168
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v596
	goto L169
L171:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v596)+68)) = v650
	goto L170
L172:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v596)+108)) = v648
	goto L170
L173:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v596)+32)) = v646
	goto L170
L174:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v596)+8)) = v644
	goto L170
L175:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v596)+8))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v623)+8))
	if v624 < int32(4) {
		v635 = v623
		goto L184
	} else {
		goto L185
	}
L176:
	;
	v600 = v597 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v596)+5)) = uint8(v600)
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+4)))
	if v602 == int32(7) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v608 = v600 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v596)+5)) = uint8(v608)
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v596)+8))
	if v610 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	switch v602 + int32(-5) {
	case 0:
		goto L173
	case 1:
		goto L174
	default:
		goto L169
	case 3:
		goto L172
	case 4:
		goto L171
	case 5:
		goto L175
	}
L180:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v596)+12))
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+5)))
	if v620&int32(3) != 0 {
		v596 = v619
		v597 = v620
		goto L176
	} else {
		goto L183
	}
L181:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610)+5)))
	if v613&int32(3) == int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	F_reallymarkobject(m, l0, v610)
	mBase = m.M
	goto L180
L183:
	;
	goto L169
L184:
	;
	if v635 != v596+int32(16) {
		goto L169
	} else {
		goto L187
	}
L185:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v623)))
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627)+5)))
	if v628&int32(3) == int32(0) {
		v635 = v623
		goto L184
	} else {
		goto L186
	}
L186:
	;
	F_reallymarkobject(m, l0, v627)
	mBase = m.M
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v596)+8))
	v635 = v634
	goto L184
L187:
	;
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+5)))
	v642 = v640 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v596)+5)) = uint8(v642)
	goto L168
L188:
	;
	v839 = v661
	goto L158
L189:
	;
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)))
	if v740 != 0 {
		goto L211
	} else {
		goto L212
	}
L190:
	;
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+5)))
	v677 = v667
	v678 = v675
	goto L199
L191:
	;
	goto L189
L192:
	;
	goto L191
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v677
	goto L192
L194:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v677)+68)) = v731
	goto L193
L195:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v677)+108)) = v729
	goto L193
L196:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v677)+32)) = v727
	goto L193
L197:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v677)+8)) = v725
	goto L193
L198:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v677)+8))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v704)+8))
	if v705 < int32(4) {
		v716 = v704
		goto L207
	} else {
		goto L208
	}
L199:
	;
	v681 = v678 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v677)+5)) = uint8(v681)
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+4)))
	if v683 == int32(7) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v689 = v681 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v677)+5)) = uint8(v689)
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v677)+8))
	if v691 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	switch v683 + int32(-5) {
	case 0:
		goto L196
	case 1:
		goto L197
	default:
		goto L192
	case 3:
		goto L195
	case 4:
		goto L194
	case 5:
		goto L198
	}
L203:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v677)+12))
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700)+5)))
	if v701&int32(3) != 0 {
		v677 = v700
		v678 = v701
		goto L199
	} else {
		goto L206
	}
L204:
	;
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v691)+5)))
	if v694&int32(3) == int32(0) {
		goto L203
	} else {
		goto L205
	}
L205:
	;
	F_reallymarkobject(m, l0, v691)
	mBase = m.M
	goto L203
L206:
	;
	goto L192
L207:
	;
	if v716 != v677+int32(16) {
		goto L192
	} else {
		goto L210
	}
L208:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708)+5)))
	if v709&int32(3) == int32(0) {
		v716 = v704
		goto L207
	} else {
		goto L209
	}
L209:
	;
	F_reallymarkobject(m, l0, v708)
	mBase = m.M
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v677)+8))
	v716 = v715
	goto L207
L210:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+5)))
	v723 = v721 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v677)+5)) = uint8(v723)
	goto L191
L211:
	;
	v747 = int32(0)
	v749 = v740
	goto L213
L212:
	;
	v839 = int32(0)
	goto L158
L213:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(20)+v747<<(uint(int32(2))%32))))
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+5)))
	if v756&int32(3) == int32(0) {
		v829 = v749
		goto L215
	} else {
		goto L216
	}
L214:
	;
	v839 = v829
	goto L158
L215:
	;
	v831 = v747 + int32(1)
	if base.Ui32(v831) < base.Ui32(v829&int32(255)) {
		v747 = v831
		v749 = v829
		goto L213
	} else {
		goto L237
	}
L216:
	;
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+5)))
	v765 = v755
	v766 = v763
	goto L225
L217:
	;
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)))
	v829 = v828
	goto L215
L218:
	;
	goto L217
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v765
	goto L218
L220:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v765)+68)) = v819
	goto L219
L221:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v765)+108)) = v817
	goto L219
L222:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v765)+32)) = v815
	goto L219
L223:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v765)+8)) = v813
	goto L219
L224:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v765)+8))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v792)+8))
	if v793 < int32(4) {
		v804 = v792
		goto L233
	} else {
		goto L234
	}
L225:
	;
	v769 = v766 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v765)+5)) = uint8(v769)
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765)+4)))
	if v771 == int32(7) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v777 = v769 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v765)+5)) = uint8(v777)
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v765)+8))
	if v779 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	switch v771 + int32(-5) {
	case 0:
		goto L222
	case 1:
		goto L223
	default:
		goto L218
	case 3:
		goto L221
	case 4:
		goto L220
	case 5:
		goto L224
	}
L229:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v765)+12))
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788)+5)))
	if v789&int32(3) != 0 {
		v765 = v788
		v766 = v789
		goto L225
	} else {
		goto L232
	}
L230:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779)+5)))
	if v782&int32(3) == int32(0) {
		goto L229
	} else {
		goto L231
	}
L231:
	;
	F_reallymarkobject(m, l0, v779)
	mBase = m.M
	goto L229
L232:
	;
	goto L218
L233:
	;
	if v804 != v765+int32(16) {
		goto L218
	} else {
		goto L236
	}
L234:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v792)))
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796)+5)))
	if v797&int32(3) == int32(0) {
		v804 = v792
		goto L233
	} else {
		goto L235
	}
L235:
	;
	F_reallymarkobject(m, l0, v796)
	mBase = m.M
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v765)+8))
	v804 = v803
	goto L233
L236:
	;
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765)+5)))
	v811 = v809 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v765)+5)) = uint8(v811)
	goto L217
L237:
	;
	goto L214
L238:
	;
	return v843<<(uint(int32(2))%32) + int32(20)
L239:
	;
	return v843<<(uint(int32(4))%32) + int32(24)
L240:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if base.Ui32(v945) < base.Ui32(v944) {
		v964 = v943
		goto L263
	} else {
		goto L264
	}
L241:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v869)+5)))
	if v870&int32(3) == int32(0) {
		goto L240
	} else {
		goto L242
	}
L242:
	;
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v869)+5)))
	v879 = v869
	v880 = v877
	goto L251
L243:
	;
	goto L240
L244:
	;
	goto L243
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v879
	goto L244
L246:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v879)+68)) = v933
	goto L245
L247:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v879)+108)) = v931
	goto L245
L248:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v879)+32)) = v929
	goto L245
L249:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v879)+8)) = v927
	goto L245
L250:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v879)+8))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v906)+8))
	if v907 < int32(4) {
		v918 = v906
		goto L259
	} else {
		goto L260
	}
L251:
	;
	v883 = v880 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v879)+5)) = uint8(v883)
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v879)+4)))
	if v885 == int32(7) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v891 = v883 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v879)+5)) = uint8(v891)
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v879)+8))
	if v893 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	switch v885 + int32(-5) {
	case 0:
		goto L248
	case 1:
		goto L249
	default:
		goto L244
	case 3:
		goto L247
	case 4:
		goto L246
	case 5:
		goto L250
	}
L255:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v879)+12))
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902)+5)))
	if v903&int32(3) != 0 {
		v879 = v902
		v880 = v903
		goto L251
	} else {
		goto L258
	}
L256:
	;
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893)+5)))
	if v896&int32(3) == int32(0) {
		goto L255
	} else {
		goto L257
	}
L257:
	;
	F_reallymarkobject(m, l0, v893)
	mBase = m.M
	goto L255
L258:
	;
	goto L244
L259:
	;
	if v918 != v879+int32(16) {
		goto L244
	} else {
		goto L262
	}
L260:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v906)))
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v910)+5)))
	if v911&int32(3) == int32(0) {
		v918 = v906
		goto L259
	} else {
		goto L261
	}
L261:
	;
	F_reallymarkobject(m, l0, v910)
	mBase = m.M
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v879)+8))
	v918 = v917
	goto L259
L262:
	;
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v879)+5)))
	v925 = v923 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v879)+5)) = uint8(v925)
	goto L243
L263:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	if base.Ui32(v943) <= base.Ui32(v967) {
		v1060 = v967
		goto L271
	} else {
		goto L272
	}
L264:
	;
	v949 = v944
	v951 = v943
	goto L265
L265:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v949)+8))
	if base.Ui32(v951) < base.Ui32(v954) {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	v964 = v956
	goto L263
L267:
	;
	v956 = v954
	goto L269
L268:
	;
	v956 = v951
	goto L269
L269:
	;
	v958 = v949 + int32(24)
	if base.Ui32(v958) <= base.Ui32(v945) {
		v949 = v958
		v951 = v956
		goto L265
	} else {
		goto L270
	}
L270:
	;
	goto L266
L271:
	;
	if base.Ui32(v964) < base.Ui32(v1060) {
		goto L299
	} else {
		goto L300
	}
L272:
	;
	v971 = v967
	v974 = v943
	goto L273
L273:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v971)+8))
	if v976 < int32(4) {
		v1054 = v974
		goto L275
	} else {
		goto L276
	}
L274:
	;
	v1060 = v1056
	goto L271
L275:
	;
	v1056 = v971 + int32(16)
	if base.Ui32(v1056) < base.Ui32(v1054) {
		v971 = v1056
		v974 = v1054
		goto L273
	} else {
		goto L298
	}
L276:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v971)))
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v979)+5)))
	if v980&int32(3) == int32(0) {
		v1054 = v974
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v979)+5)))
	v989 = v979
	v990 = v987
	goto L286
L278:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v1054 = v1052
	goto L275
L279:
	;
	goto L278
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v989
	goto L279
L281:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+68)) = v1043
	goto L280
L282:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+108)) = v1041
	goto L280
L283:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+32)) = v1039
	goto L280
L284:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v989)+8)) = v1037
	goto L280
L285:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v989)+8))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+8))
	if v1017 < int32(4) {
		v1028 = v1016
		goto L294
	} else {
		goto L295
	}
L286:
	;
	v993 = v990 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v989)+5)) = uint8(v993)
	v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v989)+4)))
	if v995 == int32(7) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1001 = v993 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v989)+5)) = uint8(v1001)
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v989)+8))
	if v1003 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	switch v995 + int32(-5) {
	case 0:
		goto L283
	case 1:
		goto L284
	default:
		goto L279
	case 3:
		goto L282
	case 4:
		goto L281
	case 5:
		goto L285
	}
L290:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v989)+12))
	v1013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1012)+5)))
	if v1013&int32(3) != 0 {
		v989 = v1012
		v990 = v1013
		goto L286
	} else {
		goto L293
	}
L291:
	;
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1003)+5)))
	if v1006&int32(3) == int32(0) {
		goto L290
	} else {
		goto L292
	}
L292:
	;
	F_reallymarkobject(m, l0, v1003)
	mBase = m.M
	goto L290
L293:
	;
	goto L279
L294:
	;
	if v1028 != v989+int32(16) {
		goto L279
	} else {
		goto L297
	}
L295:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v1016)))
	v1021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1020)+5)))
	if v1021&int32(3) == int32(0) {
		v1028 = v1016
		goto L294
	} else {
		goto L296
	}
L296:
	;
	F_reallymarkobject(m, l0, v1020)
	mBase = m.M
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v989)+8))
	v1028 = v1027
	goto L294
L297:
	;
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v989)+5)))
	v1035 = v1033 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v989)+5)) = uint8(v1035)
	goto L278
L298:
	;
	goto L274
L299:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
	if int32(20000) < v1085 {
		goto L304
	} else {
		goto L305
	}
L300:
	;
	v1068 = v1060
	goto L301
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1068)+8)) = int32(0)
	v1076 = v1068 + int32(16)
	if base.Ui32(v1076) <= base.Ui32(v964) {
		v1068 = v1076
		goto L301
	} else {
		goto L303
	}
L302:
	;
	goto L299
L303:
	;
	goto L302
L304:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
	return v1119<<(uint(int32(4))%32) + v1122*int32(24) + int32(120)
L305:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	v1094 = base.I32_div_s(v1090-v1091, int32(24))
	if v1085 < int32(17) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
	if v1108 <= (v964-v1088)>>(uint(int32(2))%32) {
		goto L304
	} else {
		goto L311
	}
L307:
	;
	if v1085 <= v1094<<(uint(int32(2))%32) {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	F_luaD_reallocCI(m, v8, int32(base.Ui32(v1085)>>(uint(int32(1))%32)))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	return int32(0)
L310:
	;
	goto L306
L311:
	;
	if v1108 < int32(91) {
		goto L304
	} else {
		goto L312
	}
L312:
	;
	F_luaD_reallocstack(m, v8, int32(base.Ui32(v1108)>>(uint(int32(1))%32)))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L309
	} else {
		goto L313
	}
L313:
	;
	goto L304
L314:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	if v1138 < int32(1) {
		goto L316
	} else {
		goto L317
	}
L315:
	;
	v1134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1131)+5)))
	v1136 = v1134 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v1131)+5)) = uint8(v1136)
	goto L314
L316:
	;
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	if v1242 < int32(1) {
		goto L344
	} else {
		goto L345
	}
L317:
	;
	v1144 = int32(0)
	v1145 = v1138
	goto L318
L318:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v1150 = int32(4)
	v1152 = v1149 + v1144<<(uint(v1150)%32)
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1152)+8))
	if v1153 < v1150 {
		v1230 = v1145
		goto L320
	} else {
		goto L321
	}
L319:
	;
	goto L316
L320:
	;
	v1233 = v1144 + int32(1)
	if v1233 < v1230 {
		v1144 = v1233
		v1145 = v1230
		goto L318
	} else {
		goto L343
	}
L321:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1152)))
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1156)+5)))
	if v1157&int32(3) == int32(0) {
		v1230 = v1145
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1156)+5)))
	v1166 = v1156
	v1167 = v1164
	goto L331
L323:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	v1230 = v1229
	goto L320
L324:
	;
	goto L323
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1166
	goto L324
L326:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1166)+68)) = v1220
	goto L325
L327:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1166)+108)) = v1218
	goto L325
L328:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1166)+32)) = v1216
	goto L325
L329:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1166)+8)) = v1214
	goto L325
L330:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+8))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+8))
	if v1194 < int32(4) {
		v1205 = v1193
		goto L339
	} else {
		goto L340
	}
L331:
	;
	v1170 = v1167 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v1166)+5)) = uint8(v1170)
	v1172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1166)+4)))
	if v1172 == int32(7) {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v1178 = v1170 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v1166)+5)) = uint8(v1178)
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+8))
	if v1180 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	switch v1172 + int32(-5) {
	case 0:
		goto L328
	case 1:
		goto L329
	default:
		goto L324
	case 3:
		goto L327
	case 4:
		goto L326
	case 5:
		goto L330
	}
L335:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+12))
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1189)+5)))
	if v1190&int32(3) != 0 {
		v1166 = v1189
		v1167 = v1190
		goto L331
	} else {
		goto L338
	}
L336:
	;
	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1180)+5)))
	if v1183&int32(3) == int32(0) {
		goto L335
	} else {
		goto L337
	}
L337:
	;
	F_reallymarkobject(m, l0, v1180)
	mBase = m.M
	goto L335
L338:
	;
	goto L324
L339:
	;
	if v1205 != v1166+int32(16) {
		goto L324
	} else {
		goto L342
	}
L340:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1193)))
	v1198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1197)+5)))
	if v1198&int32(3) == int32(0) {
		v1205 = v1193
		goto L339
	} else {
		goto L341
	}
L341:
	;
	F_reallymarkobject(m, l0, v1197)
	mBase = m.M
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+8))
	v1205 = v1204
	goto L339
L342:
	;
	v1210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1166)+5)))
	v1212 = v1210 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v1166)+5)) = uint8(v1212)
	goto L323
L343:
	;
	goto L319
L344:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
	if v1276 < int32(1) {
		v1374 = v1276
		goto L351
	} else {
		goto L352
	}
L345:
	;
	v1248 = int32(0)
	v1249 = v1242
	goto L346
L346:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1253+v1248<<(uint(int32(2))%32))))
	if v1257 == int32(0) {
		v1265 = v1249
		goto L348
	} else {
		goto L349
	}
L347:
	;
	goto L344
L348:
	;
	v1267 = v1248 + int32(1)
	if v1267 < v1265 {
		v1248 = v1267
		v1249 = v1265
		goto L346
	} else {
		goto L350
	}
L349:
	;
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1257)+5)))
	v1262 = v1260 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v1257)+5)) = uint8(v1262)
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	v1265 = v1264
	goto L348
L350:
	;
	goto L347
L351:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
	if v1378 < int32(1) {
		v1406 = v1378
		v1409 = v1374
		goto L379
	} else {
		goto L380
	}
L352:
	;
	v1282 = int32(0)
	v1283 = v1276
	goto L353
L353:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1287+v1282<<(uint(int32(2))%32))))
	if v1291 == int32(0) {
		v1367 = v1283
		goto L355
	} else {
		goto L356
	}
L354:
	;
	v1374 = v1367
	goto L351
L355:
	;
	v1369 = v1282 + int32(1)
	if v1369 < v1367 {
		v1282 = v1369
		v1283 = v1367
		goto L353
	} else {
		goto L378
	}
L356:
	;
	v1294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+5)))
	if v1294&int32(3) == int32(0) {
		v1367 = v1283
		goto L355
	} else {
		goto L357
	}
L357:
	;
	v1301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+5)))
	v1303 = v1291
	v1304 = v1301
	goto L366
L358:
	;
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
	v1367 = v1366
	goto L355
L359:
	;
	goto L358
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1303
	goto L359
L361:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+68)) = v1357
	goto L360
L362:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+108)) = v1355
	goto L360
L363:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+32)) = v1353
	goto L360
L364:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+8)) = v1351
	goto L360
L365:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+8))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1330)+8))
	if v1331 < int32(4) {
		v1342 = v1330
		goto L374
	} else {
		goto L375
	}
L366:
	;
	v1307 = v1304 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v1303)+5)) = uint8(v1307)
	v1309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303)+4)))
	if v1309 == int32(7) {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1315 = v1307 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v1303)+5)) = uint8(v1315)
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+8))
	if v1317 == int32(0) {
		goto L370
	} else {
		goto L371
	}
L369:
	;
	switch v1309 + int32(-5) {
	case 0:
		goto L363
	case 1:
		goto L364
	default:
		goto L359
	case 3:
		goto L362
	case 4:
		goto L361
	case 5:
		goto L365
	}
L370:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+12))
	v1327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1326)+5)))
	if v1327&int32(3) != 0 {
		v1303 = v1326
		v1304 = v1327
		goto L366
	} else {
		goto L373
	}
L371:
	;
	v1320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1317)+5)))
	if v1320&int32(3) == int32(0) {
		goto L370
	} else {
		goto L372
	}
L372:
	;
	F_reallymarkobject(m, l0, v1317)
	mBase = m.M
	goto L370
L373:
	;
	goto L359
L374:
	;
	if v1342 != v1303+int32(16) {
		goto L359
	} else {
		goto L377
	}
L375:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1330)))
	v1335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1334)+5)))
	if v1335&int32(3) == int32(0) {
		v1342 = v1330
		goto L374
	} else {
		goto L376
	}
L376:
	;
	F_reallymarkobject(m, l0, v1334)
	mBase = m.M
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+8))
	v1342 = v1341
	goto L374
L377:
	;
	v1347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303)+5)))
	v1349 = v1347 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v1303)+5)) = uint8(v1349)
	goto L358
L378:
	;
	goto L354
L379:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	v1432 = v1413<<(uint(int32(4))%32) + v1406*int32(12) + (v1409+v1419+v1421+v1423)<<(uint(int32(2))%32) + int32(76)
	goto L1
L380:
	;
	v1382 = v1378
	v1384 = int32(0)
	goto L381
L381:
	;
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1389+v1384*int32(12))))
	if v1393 == int32(0) {
		v1401 = v1382
		goto L383
	} else {
		goto L384
	}
L382:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
	v1406 = v1401
	v1409 = v1405
	goto L379
L383:
	;
	v1403 = v1384 + int32(1)
	if v1403 < v1401 {
		v1382 = v1401
		v1384 = v1403
		goto L381
	} else {
		goto L385
	}
L384:
	;
	v1396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1393)+5)))
	v1398 = v1396 & int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v1393)+5)) = uint8(v1398)
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
	v1401 = v1400
	goto L383
L385:
	;
	goto L382
}
func F_pruneLastBucket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if v7&int32(1) == int32(0) {
		F__serverAssert(m, int32(_a633), int32(_a626), int32(983))
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
		if v12 != l2 {
			F__serverAssert(m, int32(_a633), int32(_a626), int32(983))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
			if v14&int32(1) != 0 {
				F__serverAssert(m, int32(_a634), int32(_a626), int32(984))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v20 = int32(base.Ui32(v14)>>(uint(int32(1))%32)) & int32(4095)
				if v20&(v20+int32(-1)) != 0 {
					F__serverAssert(m, int32(_a635), int32(_a626), int32(985))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v25 = v7 & int32(65534)
					*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v25)
					if v7&int32(4096) != 0 {
						F__serverAssert(m, int32(_a636), int32(_a626), int32(971))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
						v33 = int32(base.Ui32(v29)>>(uint(int32(1))%32)) & int32(4095)
						if v33 == int32(0) {
							F_valkey_free(m, l2)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+36))
								if v43 == int32(0) {
									v53 = l0 + l3<<(uint(int32(2))%32) + int32(32)
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
									*(*int32)(unsafe.Add(mBase, uint32(v53))) = v54 + int32(-1)
									return
								} else {
									m.T0[v43].(func(*base.Module, int32, int32))(m, l0, int32(-64))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										v53 = l0 + l3<<(uint(int32(2))%32) + int32(32)
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
										*(*int32)(unsafe.Add(mBase, uint32(v53))) = v54 + int32(-1)
										return
									}
								}
							}
						} else {
							F_moveEntry(m, l1, int32(11), l2, base.I32_ctz(v33))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								F_valkey_free(m, l2)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+36))
									if v43 == int32(0) {
										v53 = l0 + l3<<(uint(int32(2))%32) + int32(32)
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
										*(*int32)(unsafe.Add(mBase, uint32(v53))) = v54 + int32(-1)
										return
									} else {
										m.T0[v43].(func(*base.Module, int32, int32))(m, l0, int32(-64))
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											v53 = l0 + l3<<(uint(int32(2))%32) + int32(32)
											v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
											*(*int32)(unsafe.Add(mBase, uint32(v53))) = v54 + int32(-1)
											return
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
func F_psetexCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v6&int32(1) != 0 {
		v13 = v4
		v14 = v5
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		v19 = int32(0)
		F_setGenericCommand(m, l0, int32(1032), v16, v14, v17, int32(1), v19, v19, v19)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			return
		}
	} else {
		v9 = F_tryObjectEncoding(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v9
			v13 = v11
			v14 = v9
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			v19 = int32(0)
			F_setGenericCommand(m, l0, int32(1032), v16, v14, v17, int32(1), v19, v19, v19)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_pttlCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v9 = F_lookupKeyReadWithFlags(m, v5, v7, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v9 != 0 {
			v14 = int64(-1)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			if v18&int32(1) == int32(0) {
				v29 = v14
			} else {
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v9+(v18&int32(4)^int32(12)))))
				v29 = v28
			}
			if v29 == int64(-1) {
				v39 = v14
			} else {
				v33 = *(*int64)(unsafe.Add(mBase, _consts[78]))
				v34 = v29 - v33
				v35 = int64(0)
				if v35 < v34 {
					v38 = v34
				} else {
					v38 = v35
				}
				v39 = v38
			}
			F_addReplyLongLong(m, l0, v39)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				return
			}
		} else {
			F_addReplyLongLong(m, l0, int64(-2))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_pvInsertAt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
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
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v14 = base.I32_wrap_i64(v13)
	v16 = v14 & int32(1073741823)
	if base.Ui32(v16) < base.Ui32(l2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a1669), int32(_a1670), int32(280))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L7
	} else {
		goto L53
	}
L2:
	;
	F__serverAssert(m, int32(_a1671), int32(_a1670), int32(452))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L52
	}
L3:
	;
	if v16 == int32(1073741823) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v23 = v14<<(uint(int32(2))%32) + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v23
	if base.Ui32(v23) <= base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v13)>>(uint(int64(30))%64)))) {
		v46 = l0
		v47 = v13
		v48 = v14
		v49 = v16
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if base.Ui32(v49) <= base.Ui32(l2) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v31 = F_zrealloc_usable(m, l0, v23, v11+int32(12))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	v38 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+12)))
	v41 = v35&int64(1073741823) | v38<<(uint(int64(30))%64)
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v41
	v43 = base.I32_wrap_i64(v41)
	v46 = v31
	v47 = v41
	v48 = v43
	v49 = v43 & int32(1073741823)
	goto L5
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46+l2<<(uint(int32(2))%32))+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v46))) = (v47+int64(1))&int64(1073741823) | v47&int64(-1073741824)
	m.G0 = v11 + int32(16)
	return v46
L10:
	;
	v51 = int32(2)
	v53 = v46 + l2<<(uint(v51)%32)
	v55 = v53 + int32(12)
	v57 = v53 + int32(8)
	v60 = (v48 - l2) << (uint(v51) % 32)
	if v55 == v57 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	goto L11
L13:
	;
	v64 = v60 + v55
	if base.Ui32(int32(0)-v60<<(uint(int32(1))%32)) < base.Ui32(v57-v64) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v74 = (v57 ^ v55) & int32(3)
	if base.Ui32(v57) <= base.Ui32(v55) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v71 = F___memcpy(m, v55, v57, v60)
	mBase = m.M
	goto L11
L16:
	;
	if v180 == int32(0) {
		goto L12
	} else {
		goto L48
	}
L17:
	;
	if base.Ui32(v158) <= base.Ui32(int32(3)) {
		v179 = v157
		v180 = v158
		v181 = v159
		goto L16
	} else {
		goto L44
	}
L18:
	;
	if v74 != 0 {
		v140 = v60
		goto L28
	} else {
		goto L29
	}
L19:
	;
	if v74 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v55&int32(3) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v179 = v57
	v180 = v60
	v181 = v55
	goto L16
L22:
	;
	v81 = v57
	v82 = v60
	v83 = v55
	goto L24
L23:
	;
	v157 = v57
	v158 = v60
	v159 = v55
	goto L17
L24:
	;
	if v82 == int32(0) {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	*(*uint8)(unsafe.Add(mBase, uint32(v83))) = uint8(v87)
	v89 = int32(1)
	v90 = v81 + v89
	v92 = v82 + int32(-1)
	v94 = v83 + v89
	if v94&int32(3) == int32(0) {
		v157 = v90
		v158 = v92
		v159 = v94
		goto L17
	} else {
		goto L27
	}
L27:
	;
	v81 = v90
	v82 = v92
	v83 = v94
	goto L24
L28:
	;
	if v140 == int32(0) {
		goto L12
	} else {
		goto L40
	}
L29:
	;
	if v64&int32(3) == int32(0) {
		v120 = v60
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if base.Ui32(v120) <= base.Ui32(int32(3)) {
		v140 = v120
		goto L28
	} else {
		goto L36
	}
L31:
	;
	v105 = v60
	goto L32
L32:
	;
	if v105 == int32(0) {
		goto L12
	} else {
		goto L34
	}
L33:
	;
	v120 = v111
	goto L30
L34:
	;
	v111 = v105 + int32(-1)
	v112 = v55 + v111
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+v111))))
	*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v114)
	if v112&int32(3) != 0 {
		v105 = v111
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v127 = v120
	goto L37
L37:
	;
	v131 = v127 + int32(-4)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v57+v131)))
	*(*int32)(unsafe.Add(mBase, uint32(v55+v131))) = v134
	if base.Ui32(int32(3)) < base.Ui32(v131) {
		v127 = v131
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v140 = v131
	goto L28
L39:
	;
	goto L38
L40:
	;
	v147 = v140
	goto L41
L41:
	;
	v151 = v147 + int32(-1)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+v151))))
	*(*uint8)(unsafe.Add(mBase, uint32(v55+v151))) = uint8(v154)
	if v151 != 0 {
		v147 = v151
		goto L41
	} else {
		goto L43
	}
L43:
	;
	goto L12
L44:
	;
	v164 = v157
	v165 = v158
	v166 = v159
	goto L45
L45:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v168
	v170 = int32(4)
	v171 = v164 + v170
	v173 = v166 + v170
	v175 = v165 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v175) {
		v164 = v171
		v165 = v175
		v166 = v173
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v179 = v171
	v180 = v175
	v181 = v173
	goto L16
L47:
	;
	goto L46
L48:
	;
	v186 = v179
	v187 = v180
	v188 = v181
	goto L49
L49:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v190)
	v192 = int32(1)
	v197 = v187 + int32(-1)
	if v197 != 0 {
		v186 = v186 + v192
		v187 = v197
		v188 = v188 + v192
		goto L49
	} else {
		goto L51
	}
L50:
	;
	goto L12
L51:
	;
	goto L50
L52:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pvRemoveAt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int64
	_ = v218
	var v221 int64
	_ = v221
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a1672), int32(_a1670), int32(554))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L56
	} else {
		goto L61
	}
L2:
	;
	F__serverAssert(m, int32(_a1673), int32(_a1670), int32(553))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L56
	} else {
		goto L60
	}
L3:
	;
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v17 = v15 & int64(1073741823)
	if v17 == int64(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v20 = base.I32_wrap_i64(v15)
	v22 = v20 & int32(1073741823)
	if base.Ui32(v22) <= base.Ui32(l1) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v22 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	m.G0 = v11 + int32(16)
	return v234
L7:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L56
	} else {
		goto L59
	}
L8:
	;
	if base.Ui32(v22+int32(-1)) <= base.Ui32(l1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v190 = int64(1073741823)
	v191 = v15 + v190
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v191&v190 | v15&int64(-1073741824)
	v200 = base.I32_wrap_i64(int64(base.Ui64(v15) >> (uint(int64(30)) % 64)))
	if v17 == int64(1) {
		goto L52
	} else {
		goto L53
	}
L10:
	;
	v29 = int32(2)
	v31 = l0 + l1<<(uint(v29)%32)
	v33 = v31 + int32(8)
	v35 = v31 + int32(12)
	v40 = (v20 + (l1 ^ int32(-1))) << (uint(v29) % 32)
	if v33 == v35 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	goto L11
L13:
	;
	v44 = v40 + v33
	if base.Ui32(int32(0)-v40<<(uint(int32(1))%32)) < base.Ui32(v35-v44) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v54 = (v35 ^ v33) & int32(3)
	if base.Ui32(v35) <= base.Ui32(v33) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v51 = F___memcpy(m, v33, v35, v40)
	mBase = m.M
	goto L11
L16:
	;
	if v160 == int32(0) {
		goto L12
	} else {
		goto L48
	}
L17:
	;
	if base.Ui32(v138) <= base.Ui32(int32(3)) {
		v159 = v137
		v160 = v138
		v161 = v139
		goto L16
	} else {
		goto L44
	}
L18:
	;
	if v54 != 0 {
		v120 = v40
		goto L28
	} else {
		goto L29
	}
L19:
	;
	if v54 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v33&int32(3) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v159 = v35
	v160 = v40
	v161 = v33
	goto L16
L22:
	;
	v61 = v35
	v62 = v40
	v63 = v33
	goto L24
L23:
	;
	v137 = v35
	v138 = v40
	v139 = v33
	goto L17
L24:
	;
	if v62 == int32(0) {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v67)
	v69 = int32(1)
	v70 = v61 + v69
	v72 = v62 + int32(-1)
	v74 = v63 + v69
	if v74&int32(3) == int32(0) {
		v137 = v70
		v138 = v72
		v139 = v74
		goto L17
	} else {
		goto L27
	}
L27:
	;
	v61 = v70
	v62 = v72
	v63 = v74
	goto L24
L28:
	;
	if v120 == int32(0) {
		goto L12
	} else {
		goto L40
	}
L29:
	;
	if v44&int32(3) == int32(0) {
		v100 = v40
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if base.Ui32(v100) <= base.Ui32(int32(3)) {
		v120 = v100
		goto L28
	} else {
		goto L36
	}
L31:
	;
	v85 = v40
	goto L32
L32:
	;
	if v85 == int32(0) {
		goto L12
	} else {
		goto L34
	}
L33:
	;
	v100 = v91
	goto L30
L34:
	;
	v91 = v85 + int32(-1)
	v92 = v33 + v91
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v91))))
	*(*uint8)(unsafe.Add(mBase, uint32(v92))) = uint8(v94)
	if v92&int32(3) != 0 {
		v85 = v91
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v107 = v100
	goto L37
L37:
	;
	v111 = v107 + int32(-4)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v35+v111)))
	*(*int32)(unsafe.Add(mBase, uint32(v33+v111))) = v114
	if base.Ui32(int32(3)) < base.Ui32(v111) {
		v107 = v111
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v120 = v111
	goto L28
L39:
	;
	goto L38
L40:
	;
	v127 = v120
	goto L41
L41:
	;
	v131 = v127 + int32(-1)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v131))))
	*(*uint8)(unsafe.Add(mBase, uint32(v33+v131))) = uint8(v134)
	if v131 != 0 {
		v127 = v131
		goto L41
	} else {
		goto L43
	}
L43:
	;
	goto L12
L44:
	;
	v144 = v137
	v145 = v138
	v146 = v139
	goto L45
L45:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v148
	v150 = int32(4)
	v151 = v144 + v150
	v153 = v146 + v150
	v155 = v145 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v155) {
		v144 = v151
		v145 = v155
		v146 = v153
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v159 = v151
	v160 = v155
	v161 = v153
	goto L16
L47:
	;
	goto L46
L48:
	;
	v166 = v159
	v167 = v160
	v168 = v161
	goto L49
L49:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v170)
	v172 = int32(1)
	v177 = v167 + int32(-1)
	if v177 != 0 {
		v166 = v166 + v172
		v167 = v177
		v168 = v168 + v172
		goto L49
	} else {
		goto L51
	}
L50:
	;
	goto L12
L51:
	;
	goto L50
L52:
	;
	if v200 == int32(0) {
		v234 = l0
		goto L6
	} else {
		goto L58
	}
L53:
	;
	v207 = base.I32_wrap_i64(v191)<<(uint(int32(2))%32) + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v207
	if base.Ui32(v200) <= base.Ui32(v207) {
		v234 = l0
		goto L6
	} else {
		goto L54
	}
L54:
	;
	if v207 == int32(0) {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	v214 = F_zrealloc_usable(m, l0, v207, v11+int32(12))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	return int32(0)
L57:
	;
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v214)))
	v221 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v214))) = v218&int64(1073741823) | v221<<(uint(int64(30))%64)
	v234 = v214
	goto L6
L58:
	;
	goto L7
L59:
	;
	v234 = int32(0)
	goto L6
L60:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pvSplit(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v57 int64
	_ = v57
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == v3 {
		v81 = v3
		m.G0 = v11 + int32(16)
		return v81
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v19 = v17 & int32(1073741823)
		if base.Ui32(v19) <= base.Ui32(l1) {
			v81 = v3
			m.G0 = v11 + int32(16)
			return v81
		} else {
			if l1 != 0 {
				v23 = v19 - l1
				v25 = v23 << (uint(int32(2)) % 32)
				v26 = int32(8)
				v30 = F_zmalloc_usable(m, v25+v26, v11+v26)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)))
					*(*int64)(unsafe.Add(mBase, uint32(v30))) = v34<<(uint(int64(30))%64) | base.I64_extend_i32_u(v23&int32(1073741823))
					v42 = int32(8)
					v47 = l1<<(uint(int32(2))%32) + v42
					if v25 == int32(0) {
					} else {
						v51 = F__emscripten_memcpy_bulkmem(m, v30+v42, v47+v14, v25)
						mBase = m.M
					}
					v53 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
					v57 = v53&int64(-1073741824) | base.I64_extend_i32_u(l1)
					*(*int64)(unsafe.Add(mBase, uint32(v14))) = v57
					*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v47
					if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v57)>>(uint(int64(30))%64)))) <= base.Ui32(v47) {
						v79 = v14
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v79
						v81 = v30
						m.G0 = v11 + int32(16)
						return v81
					} else {
						if v47 != 0 {
							v69 = F_zrealloc_usable(m, v14, v47, v11+int32(12))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
								v74 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+12)))
								*(*int64)(unsafe.Add(mBase, uint32(v69))) = v71&int64(1073741823) | v74<<(uint(int64(30))%64)
								v79 = v69
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v79
								v81 = v30
								m.G0 = v11 + int32(16)
								return v81
							}
						} else {
							F_valkey_free(m, v14)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								v79 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v79
								v81 = v30
								m.G0 = v11 + int32(16)
								return v81
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				v81 = v14
				m.G0 = v11 + int32(16)
				return v81
			}
		}
	}
}
