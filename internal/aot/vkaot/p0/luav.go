package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaV_execute(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v182 int32
	_ = v182
	var v193 int32
	_ = v193
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int64
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int64
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v388 int64
	_ = v388
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 float64
	_ = v429
	var v430 float64
	_ = v430
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v465 float64
	_ = v465
	var v466 float64
	_ = v466
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 float64
	_ = v501
	var v502 float64
	_ = v502
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v537 float64
	_ = v537
	var v538 float64
	_ = v538
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v573 float64
	_ = v573
	var v574 float64
	_ = v574
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v612 float64
	_ = v612
	var v613 float64
	_ = v613
	var v616 float64
	_ = v616
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 float64
	_ = v631
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v722 int32
	_ = v722
	var __phi722 int32
	_ = __phi722
	var v724 int32
	_ = v724
	var __phi724 int32
	_ = __phi724
	var v733 int32
	_ = v733
	var v739 float64
	_ = v739
	var v743 float64
	_ = v743
	var v744 int64
	_ = v744
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v761 float64
	_ = v761
	var v765 int32
	_ = v765
	var v773 int32
	_ = v773
	var v776 float64
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v817 float64
	_ = v817
	var v818 int64
	_ = v818
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v835 float64
	_ = v835
	var v839 int32
	_ = v839
	var v847 int32
	_ = v847
	var v850 float64
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	var v889 int32
	_ = v889
	var v895 float64
	_ = v895
	var v899 float64
	_ = v899
	var v900 int64
	_ = v900
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v917 float64
	_ = v917
	var v921 int32
	_ = v921
	var v929 int32
	_ = v929
	var v932 float64
	_ = v932
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v941 int32
	_ = v941
	var v948 int32
	_ = v948
	var v954 int32
	_ = v954
	var v971 int32
	_ = v971
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1018 int64
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1129 float64
	_ = v1129
	var v1130 float64
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1176 int32
	_ = v1176
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1216 int32
	_ = v1216
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1240 int32
	_ = v1240
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1273 int32
	_ = v1273
	var v1289 int32
	_ = v1289
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
	var v1313 int32
	_ = v1313
	var v1321 int32
	_ = v1321
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1341 int64
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1352 int32
	_ = v1352
	var v1359 int32
	_ = v1359
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1426 int32
	_ = v1426
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1444 int64
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1453 int32
	_ = v1453
	var v1461 int32
	_ = v1461
	var v1477 int32
	_ = v1477
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1507 int32
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1524 float64
	_ = v1524
	var v1525 float64
	_ = v1525
	var v1526 float64
	_ = v1526
	var v1527 float64
	_ = v1527
	var v1538 int32
	_ = v1538
	var v1552 int32
	_ = v1552
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1564 float64
	_ = v1564
	var v1569 int32
	_ = v1569
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1581 float64
	_ = v1581
	var v1586 int32
	_ = v1586
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1598 float64
	_ = v1598
	var v1602 int32
	_ = v1602
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1613 int32
	_ = v1613
	var v1614 int64
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1618 int64
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1622 int64
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1649 int64
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1659 int32
	_ = v1659
	var v1665 int32
	_ = v1665
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1726 int64
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1738 int32
	_ = v1738
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1783 int32
	_ = v1783
	var v1786 int32
	_ = v1786
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1810 int32
	_ = v1810
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1819 int32
	_ = v1819
	var v1821 int32
	_ = v1821
	var v1829 int32
	_ = v1829
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1864 int32
	_ = v1864
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1893 int32
	_ = v1893
	var v1905 int32
	_ = v1905
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1924 int32
	_ = v1924
	var v1925 int64
	_ = v1925
	var v1927 int32
	_ = v1927
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1944 int32
	_ = v1944
	var v1949 float64
	_ = v1949
	var v1950 float64
	_ = v1950
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v26 = l1
	goto L1
L1:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v49 = v47 + int32(20)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v59 = v50
	v60 = v51
	goto L4
L2:
	;
	m.G0 = v23 + int32(16)
	return
L3:
	;
	goto L2
L4:
	;
	v75 = v60 + int32(4)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v77&int32(12) == int32(0) {
		v152 = v59
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v160 = int32(base.Ui32(v76)>>(uint(int32(6))%32)) & int32(255)
	v163 = v152 + v160<<(uint(int32(4))%32)
	switch v76 & int32(63) {
	case 0:
		goto L67
	case 1:
		goto L66
	case 2:
		goto L65
	case 3:
		goto L64
	case 4:
		goto L63
	case 5:
		goto L62
	case 6:
		goto L61
	case 7:
		goto L60
	case 8:
		goto L59
	case 9:
		goto L58
	case 10:
		goto L57
	case 11:
		goto L56
	case 12:
		goto L55
	case 13:
		goto L54
	case 14:
		goto L53
	case 15:
		goto L52
	case 16:
		goto L51
	case 17:
		goto L50
	case 18:
		goto L49
	case 19:
		goto L48
	case 20:
		goto L47
	case 21:
		goto L46
	case 22:
		goto L45
	case 23:
		goto L44
	case 24:
		goto L43
	case 25:
		goto L42
	case 26:
		goto L41
	case 27:
		goto L40
	case 28:
		goto L39
	case 29:
		goto L38
	case 30:
		goto L37
	case 31:
		goto L36
	case 32:
		goto L35
	case 33:
		goto L34
	case 34:
		goto L33
	case 35:
		goto L32
	case 36:
		goto L31
	case 37:
		goto L30
	default:
		v59 = v152
		v60 = v75
		goto L4
	}
L7:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v84 = v82 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v84
	if v84 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v77&int32(4) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	if v77&int32(8) == int32(0) {
		v106 = v94
		goto L8
	} else {
		goto L12
	}
L10:
	;
	if v77&int32(4) == int32(0) {
		v152 = v59
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	v106 = v92
	goto L8
L12:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v100
	F_luaD_callhook(m, l0, int32(3), int32(-1))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	v106 = v94
	goto L8
L15:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v147 != int32(1) {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+16))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v120 = (v75-v115)>>(uint(int32(2))%32) + int32(-1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v114)+20))
	if v121 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_luaD_callhook(m, l0, int32(2), v139)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L13
	} else {
		goto L25
	}
L18:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v121+v120<<(uint(int32(2))%32))))
	if v120 == int32(0) {
		v139 = v129
		goto L17
	} else {
		goto L22
	}
L19:
	;
	v122 = int32(0)
	if v120 == v122 {
		v139 = v122
		goto L17
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32(v75) <= base.Ui32(v106) {
		v139 = v122
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	if base.Ui32(v75) <= base.Ui32(v106) {
		v139 = v129
		goto L17
	} else {
		goto L23
	}
L23:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v121+(v106-v115)+int32(-4))))
	if v129 == v137 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	v139 = v129
	goto L17
L25:
	;
	goto L15
L26:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v152 = v151
	goto L6
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v60
	goto L3
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(3)
	v1949 = *(*float64)(unsafe.Add(mBase, uint32(v163)))
	v1950 = *(*float64)(unsafe.Add(mBase, uint32(v163)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v163))) = base.F64_sub(v1949, v1950)
	v59 = v152
	v60 = v75 + int32(base.Ui32(v76)>>(uint(int32(12))%32))&int32(1048572) + int32(-524284)
	goto L4
L29:
	;
	v1939 = m.G3
	F_luaG_runerror(m, l0, v1939+int32(_a2275), int32(0))
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L13
	} else {
		goto L439
	}
L30:
	;
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1854)))
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v1854)+4))
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v1861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1860)+73)))
	v1864 = (v1855-v1856)>>(uint(int32(4))%32) + (v1861 ^ int32(-1))
	if base.Ui32(v76) <= base.Ui32(int32(8388607)) {
		goto L427
	} else {
		goto L428
	}
L31:
	;
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(v1759)+16))
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v1760+int32(base.Ui32(v76)>>(uint(int32(12))%32))&int32(1048572))))
	v1767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1766)+72)))
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v1769 = F_luaF_newLclosure(m, l0, v1767, v1768)
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L13
	} else {
		goto L413
	}
L32:
	;
	F_luaF_close(m, l0, v163)
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L13
	} else {
		goto L412
	}
L33:
	;
	v1665 = int32(base.Ui32(v76)>>(uint(int32(14))%32)) & int32(511)
	if base.Ui32(v76) <= base.Ui32(int32(8388607)) {
		goto L394
	} else {
		goto L395
	}
L34:
	;
	v1614 = *(*int64)(unsafe.Add(mBase, uint32(v163)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v163)+80)) = v1614
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v163)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+88)) = v1616
	v1618 = *(*int64)(unsafe.Add(mBase, uint32(v163)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v163)+64)) = v1618
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v163)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+72)) = v1620
	v1622 = *(*int64)(unsafe.Add(mBase, uint32(v163)))
	*(*int64)(unsafe.Add(mBase, uint32(v163)+48)) = v1622
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+56)) = v1624
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v163 + int32(96)
	F_luaD_call(m, l0, v163+int32(48), int32(base.Ui32(v76)>>(uint(int32(14))%32))&int32(511))
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L13
	} else {
		goto L390
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	if v1552 == int32(3) {
		goto L373
	} else {
		goto L374
	}
L36:
	;
	v1524 = *(*float64)(unsafe.Add(mBase, uint32(v163)+32))
	v1525 = *(*float64)(unsafe.Add(mBase, uint32(v163)))
	v1526 = base.F64_add(v1524, v1525)
	v1527 = *(*float64)(unsafe.Add(mBase, uint32(v163)+16))
	if base.F64_gt(v1524, float64(0)) == int32(0) {
		goto L368
	} else {
		goto L369
	}
L37:
	;
	if base.Ui32(v76) < base.Ui32(int32(_a14)) {
		goto L359
	} else {
		goto L360
	}
L38:
	;
	if base.Ui32(v76) < base.Ui32(int32(_a14)) {
		goto L345
	} else {
		goto L346
	}
L39:
	;
	v1359 = int32(base.Ui32(v76)>>(uint(int32(14))%32)) & int32(511)
	if base.Ui32(v76) < base.Ui32(int32(_a14)) {
		goto L337
	} else {
		goto L338
	}
L40:
	;
	v1329 = v152 + int32(base.Ui32(v76)>>(uint(int32(19))%32))&int32(8176)
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1329)+8))
	switch v1330 {
	case 0:
		v1335 = int32(1)
		goto L332
	case 1:
		goto L334
	default:
		goto L333
	}
L41:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	switch v1302 {
	case 0:
		v1307 = int32(1)
		goto L327
	case 1:
		goto L329
	default:
		goto L328
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	v1106 = int32(base.Ui32(v76) >> (uint(int32(19)) % 32))
	if v76 < int32(0) {
		goto L283
	} else {
		goto L284
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	v1071 = int32(base.Ui32(v76) >> (uint(int32(19)) % 32))
	if v76 < int32(0) {
		goto L273
	} else {
		goto L274
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	v1030 = int32(0)
	v1032 = int32(base.Ui32(v76) >> (uint(int32(19)) % 32))
	if v76 < v1030 {
		goto L262
	} else {
		goto L263
	}
L45:
	;
	v59 = v152
	v60 = v75 + int32(base.Ui32(v76)>>(uint(int32(12))%32))&int32(1048572) + int32(-524284)
	goto L4
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	v997 = int32(base.Ui32(v76)>>(uint(int32(14))%32)) & int32(511)
	v999 = int32(base.Ui32(v76) >> (uint(int32(23)) % 32))
	F_luaV_concat(m, l0, v997-v999+int32(1), v997)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L13
	} else {
		goto L257
	}
L47:
	;
	v660 = v152 + int32(base.Ui32(v76)>>(uint(int32(19))%32))&int32(8176)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v660)+8))
	switch v661 + int32(-4) {
	case 0:
		goto L176
	case 1:
		goto L177
	default:
		goto L175
	}
L48:
	;
	v646 = v152 + int32(base.Ui32(v76)>>(uint(int32(19))%32))&int32(8176)
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v646)+8))
	switch v647 {
	case 0:
		v652 = int32(1)
		goto L172
	case 1:
		goto L174
	default:
		goto L173
	}
L49:
	;
	v627 = v152 + int32(base.Ui32(v76)>>(uint(int32(19))%32))&int32(8176)
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v627)+8))
	if v628 != int32(3) {
		goto L169
	} else {
		goto L170
	}
L50:
	;
	if v76&int32(4194304) != 0 {
		goto L159
	} else {
		goto L160
	}
L51:
	;
	if v76&int32(4194304) != 0 {
		goto L149
	} else {
		goto L150
	}
L52:
	;
	if v76&int32(4194304) != 0 {
		goto L139
	} else {
		goto L140
	}
L53:
	;
	if v76&int32(4194304) != 0 {
		goto L129
	} else {
		goto L130
	}
L54:
	;
	if v76&int32(4194304) != 0 {
		goto L119
	} else {
		goto L120
	}
L55:
	;
	if v76&int32(4194304) != 0 {
		goto L109
	} else {
		goto L110
	}
L56:
	;
	v387 = v152 + int32(base.Ui32(v76)>>(uint(int32(19))%32))&int32(8176)
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v387)))
	*(*int64)(unsafe.Add(mBase, uint32(v163)+16)) = v388
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v387)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+24)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	if v76&int32(4194304) != 0 {
		goto L105
	} else {
		goto L106
	}
L57:
	;
	v339 = int32(base.Ui32(v76) >> (uint(int32(23)) % 32))
	v348 = int32(base.Ui32(v339)>>(uint(int32(3))%32)) & int32(31)
	if v348 != 0 {
		goto L94
	} else {
		goto L95
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	v317 = int32(base.Ui32(v76) >> (uint(int32(19)) % 32))
	if v76 < int32(0) {
		goto L86
	} else {
		goto L87
	}
L59:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(base.Ui32(v76)>>(uint(int32(21))%32))&int32(2044))))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+8))
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v163)))
	*(*int64)(unsafe.Add(mBase, uint32(v284))) = v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v284)+8)) = v287
	if v287 < int32(4) {
		v59 = v152
		v60 = v75
		goto L4
	} else {
		goto L80
	}
L60:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v265
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	F_luaV_settable(m, l0, v23, v53+int32(base.Ui32(v76)>>(uint(int32(10))%32))&int32(4194288), v163)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L13
	} else {
		goto L79
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	if v76&int32(4194304) != 0 {
		goto L75
	} else {
		goto L76
	}
L62:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v235
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	F_luaV_gettable(m, l0, v23, v53+int32(base.Ui32(v76)>>(uint(int32(10))%32))&int32(4194288), v163)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L13
	} else {
		goto L74
	}
L63:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(base.Ui32(v76)>>(uint(int32(21))%32))&int32(2044))))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+8))
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v230)))
	*(*int64)(unsafe.Add(mBase, uint32(v163))) = v231
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v230)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v233
	v59 = v152
	v60 = v75
	goto L4
L64:
	;
	v208 = v152 + int32(base.Ui32(v76)>>(uint(int32(19))%32))&int32(8176)
	goto L71
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = int32(base.Ui32(v76) >> (uint(int32(23)) % 32))
	if v76&int32(8372224) != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v179 = v53 + int32(base.Ui32(v76)>>(uint(int32(10))%32))&int32(4194288)
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v179)))
	*(*int64)(unsafe.Add(mBase, uint32(v163))) = v180
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v182
	v59 = v152
	v60 = v75
	goto L4
L67:
	;
	v170 = v152 + int32(base.Ui32(v76)>>(uint(int32(19))%32))&int32(8176)
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v170)))
	*(*int64)(unsafe.Add(mBase, uint32(v163))) = v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v173
	v59 = v152
	v60 = v75
	goto L4
L68:
	;
	v193 = v60 + int32(8)
	goto L70
L69:
	;
	v193 = v75
	goto L70
L70:
	;
	v59 = v152
	v60 = v193
	goto L4
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = int32(0)
	v222 = v208 + int32(-16)
	if base.Ui32(v163) <= base.Ui32(v222) {
		v208 = v222
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v59 = v152
	v60 = v75
	goto L4
L74:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v247
	v60 = v75
	goto L4
L75:
	;
	v256 = v53
	goto L77
L76:
	;
	v256 = v152
	goto L77
L77:
	;
	F_luaV_gettable(m, l0, v152+int32(base.Ui32(v76)>>(uint(int32(19))%32))&int32(8176), v256+int32(base.Ui32(v76)>>(uint(int32(10))%32))&int32(4080), v163)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L13
	} else {
		goto L78
	}
L78:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v264
	v60 = v75
	goto L4
L79:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v277
	v60 = v75
	goto L4
L80:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+5)))
	if v292&int32(3) == int32(0) {
		v59 = v152
		v60 = v75
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+5)))
	if v297&int32(4) == int32(0) {
		v59 = v152
		v60 = v75
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+21)))
	if v303 != int32(1) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v59 = v152
	v60 = v75
	goto L4
L84:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+20)))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+5)))
	v313 = v307&int32(3) | v310&int32(248)
	*(*uint8)(unsafe.Add(mBase, uint32(v283)+5)) = uint8(v313)
	goto L83
L85:
	;
	F_reallymarkobject(m, v302, v291)
	mBase = m.M
	goto L83
L86:
	;
	v326 = v53 + v317&int32(4080)
	goto L88
L87:
	;
	v326 = v152 + v317&int32(8176)
	goto L88
L88:
	;
	if v76&int32(4194304) != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v329 = v53
	goto L91
L90:
	;
	v329 = v152
	goto L91
L91:
	;
	F_luaV_settable(m, l0, v163, v326, v329+int32(base.Ui32(v76)>>(uint(int32(10))%32))&int32(4080))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L13
	} else {
		goto L92
	}
L92:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v337
	v60 = v75
	goto L4
L93:
	;
	v356 = int32(base.Ui32(v76)>>(uint(int32(14))%32)) & int32(511)
	v365 = int32(base.Ui32(v356)>>(uint(int32(3))%32)) & int32(31)
	if v365 != 0 {
		goto L98
	} else {
		goto L99
	}
L94:
	;
	v352 = (v339&int32(7) | int32(8)) << (uint(v348+int32(-1)) % 32)
	goto L96
L95:
	;
	v352 = v339
	goto L96
L96:
	;
	goto L93
L97:
	;
	v370 = F_luaH_new(m, l0, v352, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L13
	} else {
		goto L101
	}
L98:
	;
	v369 = (v356&int32(7) | int32(8)) << (uint(v365+int32(-1)) % 32)
	goto L100
L99:
	;
	v369 = v356
	goto L100
L100:
	;
	goto L97
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v370
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+68))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v376)+64))
	if base.Ui32(v377) < base.Ui32(v378) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v382
	v60 = v75
	goto L4
L103:
	;
	F_luaC_step(m, l0)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L13
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v395 = v53
	goto L107
L106:
	;
	v395 = v152
	goto L107
L107:
	;
	F_luaV_gettable(m, l0, v387, v395+int32(base.Ui32(v76)>>(uint(int32(10))%32))&int32(4080), v163)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L13
	} else {
		goto L108
	}
L108:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v403
	v60 = v75
	goto L4
L109:
	;
	v406 = v53
	goto L111
L110:
	;
	v406 = v152
	goto L111
L111:
	;
	v409 = int32(4080)
	v411 = v406 + int32(base.Ui32(v76)>>(uint(int32(10))%32))&v409
	v413 = int32(base.Ui32(v76) >> (uint(int32(19)) % 32))
	if v76 < int32(0) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	F_Arith(m, l0, v163, v422, v411, int32(5))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L13
	} else {
		goto L118
	}
L113:
	;
	v422 = v53 + v413&v409
	goto L115
L114:
	;
	v422 = v152 + v413&int32(8176)
	goto L115
L115:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)+8))
	if v423 != int32(3) {
		goto L112
	} else {
		goto L116
	}
L116:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v411)+8))
	if v426 != int32(3) {
		goto L112
	} else {
		goto L117
	}
L117:
	;
	v429 = *(*float64)(unsafe.Add(mBase, uint32(v411)))
	v430 = *(*float64)(unsafe.Add(mBase, uint32(v422)))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v163))) = base.F64_add(v430, v429)
	v59 = v152
	v60 = v75
	goto L4
L118:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v439
	v60 = v75
	goto L4
L119:
	;
	v442 = v53
	goto L121
L120:
	;
	v442 = v152
	goto L121
L121:
	;
	v445 = int32(4080)
	v447 = v442 + int32(base.Ui32(v76)>>(uint(int32(10))%32))&v445
	v449 = int32(base.Ui32(v76) >> (uint(int32(19)) % 32))
	if v76 < int32(0) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	F_Arith(m, l0, v163, v458, v447, int32(6))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L13
	} else {
		goto L128
	}
L123:
	;
	v458 = v53 + v449&v445
	goto L125
L124:
	;
	v458 = v152 + v449&int32(8176)
	goto L125
L125:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v458)+8))
	if v459 != int32(3) {
		goto L122
	} else {
		goto L126
	}
L126:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v447)+8))
	if v462 != int32(3) {
		goto L122
	} else {
		goto L127
	}
L127:
	;
	v465 = *(*float64)(unsafe.Add(mBase, uint32(v447)))
	v466 = *(*float64)(unsafe.Add(mBase, uint32(v458)))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v163))) = base.F64_sub(v466, v465)
	v59 = v152
	v60 = v75
	goto L4
L128:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v475
	v60 = v75
	goto L4
L129:
	;
	v478 = v53
	goto L131
L130:
	;
	v478 = v152
	goto L131
L131:
	;
	v481 = int32(4080)
	v483 = v478 + int32(base.Ui32(v76)>>(uint(int32(10))%32))&v481
	v485 = int32(base.Ui32(v76) >> (uint(int32(19)) % 32))
	if v76 < int32(0) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	F_Arith(m, l0, v163, v494, v483, int32(7))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L13
	} else {
		goto L138
	}
L133:
	;
	v494 = v53 + v485&v481
	goto L135
L134:
	;
	v494 = v152 + v485&int32(8176)
	goto L135
L135:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v494)+8))
	if v495 != int32(3) {
		goto L132
	} else {
		goto L136
	}
L136:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v483)+8))
	if v498 != int32(3) {
		goto L132
	} else {
		goto L137
	}
L137:
	;
	v501 = *(*float64)(unsafe.Add(mBase, uint32(v483)))
	v502 = *(*float64)(unsafe.Add(mBase, uint32(v494)))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v163))) = base.F64_mul(v502, v501)
	v59 = v152
	v60 = v75
	goto L4
L138:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v511
	v60 = v75
	goto L4
L139:
	;
	v514 = v53
	goto L141
L140:
	;
	v514 = v152
	goto L141
L141:
	;
	v517 = int32(4080)
	v519 = v514 + int32(base.Ui32(v76)>>(uint(int32(10))%32))&v517
	v521 = int32(base.Ui32(v76) >> (uint(int32(19)) % 32))
	if v76 < int32(0) {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	F_Arith(m, l0, v163, v530, v519, int32(8))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L13
	} else {
		goto L148
	}
L143:
	;
	v530 = v53 + v521&v517
	goto L145
L144:
	;
	v530 = v152 + v521&int32(8176)
	goto L145
L145:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v530)+8))
	if v531 != int32(3) {
		goto L142
	} else {
		goto L146
	}
L146:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v519)+8))
	if v534 != int32(3) {
		goto L142
	} else {
		goto L147
	}
L147:
	;
	v537 = *(*float64)(unsafe.Add(mBase, uint32(v519)))
	v538 = *(*float64)(unsafe.Add(mBase, uint32(v530)))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v163))) = base.F64_div(v538, v537)
	v59 = v152
	v60 = v75
	goto L4
L148:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v547
	v60 = v75
	goto L4
L149:
	;
	v550 = v53
	goto L151
L150:
	;
	v550 = v152
	goto L151
L151:
	;
	v553 = int32(4080)
	v555 = v550 + int32(base.Ui32(v76)>>(uint(int32(10))%32))&v553
	v557 = int32(base.Ui32(v76) >> (uint(int32(19)) % 32))
	if v76 < int32(0) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	F_Arith(m, l0, v163, v566, v555, int32(9))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L13
	} else {
		goto L158
	}
L153:
	;
	v566 = v53 + v557&v553
	goto L155
L154:
	;
	v566 = v152 + v557&int32(8176)
	goto L155
L155:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	if v567 != int32(3) {
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v555)+8))
	if v570 != int32(3) {
		goto L152
	} else {
		goto L157
	}
L157:
	;
	v573 = *(*float64)(unsafe.Add(mBase, uint32(v555)))
	v574 = *(*float64)(unsafe.Add(mBase, uint32(v566)))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v163))) = base.F64_sub(v574, base.F64_mul(v573, base.F64_floor(base.F64_div(v574, v573))))
	v59 = v152
	v60 = v75
	goto L4
L158:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v586
	v60 = v75
	goto L4
L159:
	;
	v589 = v53
	goto L161
L160:
	;
	v589 = v152
	goto L161
L161:
	;
	v592 = int32(4080)
	v594 = v589 + int32(base.Ui32(v76)>>(uint(int32(10))%32))&v592
	v596 = int32(base.Ui32(v76) >> (uint(int32(19)) % 32))
	if v76 < int32(0) {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	F_Arith(m, l0, v163, v605, v594, int32(10))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L13
	} else {
		goto L168
	}
L163:
	;
	v605 = v53 + v596&v592
	goto L165
L164:
	;
	v605 = v152 + v596&int32(8176)
	goto L165
L165:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)+8))
	if v606 != int32(3) {
		goto L162
	} else {
		goto L166
	}
L166:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v594)+8))
	if v609 != int32(3) {
		goto L162
	} else {
		goto L167
	}
L167:
	;
	v612 = *(*float64)(unsafe.Add(mBase, uint32(v594)))
	v613 = *(*float64)(unsafe.Add(mBase, uint32(v605)))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(3)
	v616 = F_pow(m, v613, v612)
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(v163))) = v616
	v59 = v152
	v60 = v75
	goto L4
L168:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v622
	v60 = v75
	goto L4
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	F_Arith(m, l0, v163, v627, v627, int32(11))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L13
	} else {
		goto L171
	}
L170:
	;
	v631 = *(*float64)(unsafe.Add(mBase, uint32(v627)))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v163))) = base.F64_neg(v631)
	v59 = v152
	v60 = v75
	goto L4
L171:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v640
	v60 = v75
	goto L4
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v652
	v59 = v152
	v60 = v75
	goto L4
L173:
	;
	v652 = int32(0)
	goto L172
L174:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v646)))
	v652 = base.B2i32(v648 == int32(0))
	goto L172
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	v983 = m.G398
	v985 = F_call_binTM(m, l0, v660, v983, v163, int32(12))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L13
	} else {
		goto L254
	}
L176:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v660)))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v976)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v163))) = base.F64_convert_i32_u(v977)
	v59 = v152
	v60 = v75
	goto L4
L177:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v660)))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v664)+36))
	if v674 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v163))) = base.F64_convert_i32_s(v971)
	v59 = v152
	v60 = v75
	goto L4
L179:
	;
	v971 = v954
	goto L178
L180:
	;
	v712 = m.G3
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v664)+24))
	if v713 != v712+int32(_a2240) {
		goto L193
	} else {
		goto L194
	}
L181:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v664)+20))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v677+v674<<(uint(int32(4))%32)+int32(-8))))
	if v683 != 0 {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v684 = int32(0)
	if v674 == int32(1) {
		v954 = v684
		goto L179
	} else {
		goto L183
	}
L183:
	;
	v690 = v674
	v692 = v684
	goto L184
L184:
	;
	v701 = int32(base.Ui32(v692+v690) >> (uint(int32(1)) % 32))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v677+int32(-8)+v701<<(uint(int32(4))%32))))
	if v705 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v706 = v690
	goto L188
L187:
	;
	v706 = v701
	goto L188
L188:
	;
	if v705 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v707 = v701
	goto L191
L190:
	;
	v707 = v692
	goto L191
L191:
	;
	if base.Ui32(int32(1)) < base.Ui32(v706-v707) {
		v690 = v706
		v692 = v707
		goto L184
	} else {
		goto L192
	}
L192:
	;
	v954 = v707
	goto L179
L193:
	;
	__phi722 = v674
	__phi724 = v674 + int32(1)
	v722 = __phi722
	v724 = __phi724
	goto L195
L194:
	;
	v971 = v674
	goto L178
L195:
	;
	if v724 < int32(1) {
		goto L200
	} else {
		goto L201
	}
L196:
	;
	v877 = v870
	goto L236
L197:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v785)+8))
	if v790 != 0 {
		goto L211
	} else {
		goto L212
	}
L198:
	;
	v765 = v760
	goto L205
L199:
	;
	v744 = base.I64_reinterpret_f64(v743)
	v749 = int32(-1)
	v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664)+12)))
	v756 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v744)>>(uint(int64(32))%64))+v744), v749<<(uint(v750)%32)^v749|int32(1))
	v760 = v713 + v756<<(uint(int32(5))%32)
	v761 = v743
	goto L198
L200:
	;
	v739 = base.F64_convert_i32_s(v724)
	if v724 == int32(0) {
		v760 = v713
		v761 = v739
		goto L198
	} else {
		goto L204
	}
L201:
	;
	if v724 <= v674 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v664)+20))
	v785 = v733 + v724<<(uint(int32(4))%32) + int32(-16)
	goto L197
L203:
	;
	v743 = base.F64_convert_i32_u(v724)
	goto L199
L204:
	;
	v743 = v739
	goto L199
L205:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v765)+24))
	if v773 != int32(3) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v785 = v779
	goto L197
L207:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v765)+28))
	v779 = m.G398
	if v778 != 0 {
		v765 = v778
		goto L205
	} else {
		goto L210
	}
L208:
	;
	v776 = *(*float64)(unsafe.Add(mBase, uint32(v765)+16))
	if base.F64_ne(v776, v761) != 0 {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v785 = v765
	goto L197
L210:
	;
	goto L206
L211:
	;
	v870 = int32(1)
	v872 = v724 << (uint(v870) % 32)
	if base.Ui32(v872) < base.Ui32(int32(2147483646)) {
		__phi722 = v724
		__phi724 = v872
		v722 = __phi722
		v724 = __phi724
		goto L195
	} else {
		goto L235
	}
L212:
	;
	if base.Ui32(v724-v722) <= base.Ui32(int32(1)) {
		v954 = v722
		goto L179
	} else {
		goto L213
	}
L213:
	;
	v797 = v722
	v798 = v724
	goto L214
L214:
	;
	v804 = v798 + v797
	v806 = int32(base.Ui32(v804) >> (uint(int32(1)) % 32))
	if base.Ui32(v804) < base.Ui32(int32(2)) {
		goto L219
	} else {
		goto L220
	}
L216:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v859)+8))
	if v864 != 0 {
		goto L228
	} else {
		goto L229
	}
L217:
	;
	v839 = v834
	goto L222
L218:
	;
	v817 = base.F64_convert_i32_u(v806)
	v818 = base.I64_reinterpret_f64(v817)
	v823 = int32(-1)
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664)+12)))
	v830 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v818)>>(uint(int64(32))%64))+v818), v823<<(uint(v824)%32)^v823|int32(1))
	v834 = v713 + v830<<(uint(int32(5))%32)
	v835 = v817
	goto L217
L219:
	;
	v834 = v713
	v835 = base.F64_convert_i32_u(v806)
	goto L217
L220:
	;
	if v674 < v806 {
		goto L218
	} else {
		goto L221
	}
L221:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v664)+20))
	v859 = v810 + v806<<(uint(int32(4))%32) + int32(-16)
	goto L216
L222:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v839)+24))
	if v847 != int32(3) {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v859 = v853
	goto L216
L224:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v839)+28))
	v853 = m.G398
	if v852 != 0 {
		v839 = v852
		goto L222
	} else {
		goto L227
	}
L225:
	;
	v850 = *(*float64)(unsafe.Add(mBase, uint32(v839)+16))
	if base.F64_ne(v850, v835) != 0 {
		goto L224
	} else {
		goto L226
	}
L226:
	;
	v859 = v839
	goto L216
L227:
	;
	goto L223
L228:
	;
	v865 = v798
	goto L230
L229:
	;
	v865 = v806
	goto L230
L230:
	;
	if v864 != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v866 = v806
	goto L233
L232:
	;
	v866 = v797
	goto L233
L233:
	;
	if base.Ui32(int32(1)) < base.Ui32(v865-v866) {
		v797 = v866
		v798 = v865
		goto L214
	} else {
		goto L234
	}
L234:
	;
	v954 = v866
	goto L179
L235:
	;
	goto L196
L236:
	;
	if v877 < int32(1) {
		goto L241
	} else {
		goto L242
	}
L237:
	;
	v971 = v877 + int32(-1)
	goto L178
L238:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v941)+8))
	if v948 != 0 {
		v877 = v877 + int32(1)
		goto L236
	} else {
		goto L252
	}
L239:
	;
	v921 = v916
	goto L246
L240:
	;
	v900 = base.I64_reinterpret_f64(v899)
	v905 = int32(-1)
	v906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664)+12)))
	v912 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v900)>>(uint(int64(32))%64))+v900), v905<<(uint(v906)%32)^v905|int32(1))
	v916 = v713 + v912<<(uint(int32(5))%32)
	v917 = v899
	goto L239
L241:
	;
	v895 = base.F64_convert_i32_s(v877)
	if v877 == int32(0) {
		v916 = v713
		v917 = v895
		goto L239
	} else {
		goto L245
	}
L242:
	;
	if v877 <= v674 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v664)+20))
	v941 = v889 + v877<<(uint(int32(4))%32) + int32(-16)
	goto L238
L244:
	;
	v899 = base.F64_convert_i32_u(v877)
	goto L240
L245:
	;
	v899 = v895
	goto L240
L246:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v921)+24))
	if v929 != int32(3) {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	v941 = v935
	goto L238
L248:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v921)+28))
	v935 = m.G398
	if v934 != 0 {
		v921 = v934
		goto L246
	} else {
		goto L251
	}
L249:
	;
	v932 = *(*float64)(unsafe.Add(mBase, uint32(v921)+16))
	if base.F64_ne(v932, v917) != 0 {
		goto L248
	} else {
		goto L250
	}
L250:
	;
	v941 = v921
	goto L238
L251:
	;
	goto L247
L252:
	;
	goto L237
L253:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v992
	v60 = v75
	goto L4
L254:
	;
	if v985 != 0 {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	v987 = m.G3
	F_luaG_typeerror(m, l0, v660, v987+int32(_a2276))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L13
	} else {
		goto L256
	}
L256:
	;
	goto L253
L257:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v1005)+68))
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v1005)+64))
	if base.Ui32(v1006) < base.Ui32(v1007) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1012 = int32(4)
	v1014 = v1011 + v160<<(uint(v1012)%32)
	v1017 = v1011 + v999<<(uint(v1012)%32)
	v1018 = *(*int64)(unsafe.Add(mBase, uint32(v1017)))
	*(*int64)(unsafe.Add(mBase, uint32(v1014))) = v1018
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v1017)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1014)+8)) = v1020
	v59 = v1011
	v60 = v75
	goto L4
L259:
	;
	F_luaC_step(m, l0)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L13
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	if v1055 != v160 {
		v1065 = v75
		goto L270
	} else {
		goto L271
	}
L262:
	;
	v1041 = v53 + v1032&int32(4080)
	goto L264
L263:
	;
	v1041 = v152 + v1032&int32(8176)
	goto L264
L264:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+8))
	if v76&int32(4194304) != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1045 = v53
	goto L267
L266:
	;
	v1045 = v152
	goto L267
L267:
	;
	v1050 = v1045 + int32(base.Ui32(v76)>>(uint(int32(10))%32))&int32(4080)
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+8))
	if v1042 != v1051 {
		v1055 = v1030
		goto L261
	} else {
		goto L268
	}
L268:
	;
	v1053 = F_luaV_equalval(m, l0, v1041, v1050)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L13
	} else {
		goto L269
	}
L269:
	;
	v1055 = v1053
	goto L261
L270:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v1068
	v60 = v1065 + int32(4)
	goto L4
L271:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v1065 = v75 + int32(base.Ui32(v1057)>>(uint(int32(12))%32))&int32(1048572) + int32(-524284)
	goto L270
L272:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v1103
	v60 = v1100 + int32(4)
	goto L4
L273:
	;
	v1080 = v53 + v1071&int32(4080)
	goto L275
L274:
	;
	v1080 = v152 + v1071&int32(8176)
	goto L275
L275:
	;
	if v76&int32(4194304) != 0 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1083 = v53
	goto L278
L277:
	;
	v1083 = v152
	goto L278
L278:
	;
	v1089 = F_luaV_lessthan(m, l0, v1080, v1083+int32(base.Ui32(v76)>>(uint(int32(10))%32))&int32(4080))
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L13
	} else {
		goto L279
	}
L279:
	;
	if v1089 != v160 {
		v1100 = v75
		goto L272
	} else {
		goto L280
	}
L280:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v1100 = v75 + int32(base.Ui32(v1092)>>(uint(int32(12))%32))&int32(1048572) + int32(-524284)
	goto L272
L281:
	;
	if v1273 != v160 {
		v1297 = v75
		goto L325
	} else {
		goto L326
	}
L282:
	;
	v1266 = F_luaG_ordererror(m, l0, v1115, v1124)
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L13
	} else {
		goto L324
	}
L283:
	;
	v1115 = v53 + v1106&int32(4080)
	goto L285
L284:
	;
	v1115 = v152 + v1106&int32(8176)
	goto L285
L285:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+8))
	if v76&int32(4194304) != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1119 = v53
	goto L288
L287:
	;
	v1119 = v152
	goto L288
L288:
	;
	v1124 = v1119 + int32(base.Ui32(v76)>>(uint(int32(10))%32))&int32(4080)
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+8))
	if v1116 != v1125 {
		goto L282
	} else {
		goto L289
	}
L289:
	;
	switch v1116 + int32(-3) {
	case 0:
		goto L292
	case 1:
		goto L291
	default:
		goto L290
	}
L290:
	;
	v1254 = F_call_orderTM(m, l0, v1115, v1124, int32(14))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L13
	} else {
		goto L320
	}
L291:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1115)))
	v1133 = int32(16)
	v1134 = v1132 + v1133
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1124)))
	v1137 = v1135 + v1133
	v1138 = F___get_tp(m)
	mBase = m.M
	v1139 = F___strcoll_l(m, v1134, v1137, v1137)
	mBase = m.M
	goto L294
L292:
	;
	v1129 = *(*float64)(unsafe.Add(mBase, uint32(v1115)))
	v1130 = *(*float64)(unsafe.Add(mBase, uint32(v1124)))
	v1273 = base.F64_le(v1129, v1130)
	goto L281
L293:
	;
	v1273 = base.B2i32(v1240 < int32(1))
	goto L281
L294:
	;
	if v1139 != 0 {
		v1240 = v1139
		goto L293
	} else {
		goto L295
	}
L295:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+12))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+12))
	v1147 = v1141
	v1150 = v1140
	v1152 = v1134
	v1155 = v1137
	goto L296
L296:
	;
	if v1152&int32(3) == int32(0) {
		v1183 = v1152
		goto L301
	} else {
		goto L302
	}
L297:
	;
	v1240 = v1228
	goto L293
L298:
	;
	if v1216 != v1150 {
		goto L316
	} else {
		goto L317
	}
L299:
	;
	if v1216 != v1147 {
		goto L298
	} else {
		goto L315
	}
L300:
	;
	v1216 = v1208 - v1152
	goto L299
L301:
	;
	v1187 = v1183
	goto L309
L302:
	;
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1152))))
	if v1169 != 0 {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v1172 = v1152
	goto L305
L304:
	;
	v1216 = v1152 - v1152
	goto L299
L305:
	;
	v1176 = v1172 + int32(1)
	if v1176&int32(3) == int32(0) {
		v1183 = v1176
		goto L301
	} else {
		goto L307
	}
L307:
	;
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1176))))
	if v1181 != 0 {
		v1172 = v1176
		goto L305
	} else {
		goto L308
	}
L308:
	;
	v1208 = v1176
	goto L300
L309:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1187)))
	v1196 = int32(-2139062144)
	if (int32(16843008)-v1193|v1193)&v1196 == v1196 {
		v1187 = v1187 + int32(4)
		goto L309
	} else {
		goto L311
	}
L310:
	;
	v1202 = v1187
	goto L312
L311:
	;
	goto L310
L312:
	;
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1202))))
	if v1206 != 0 {
		v1202 = v1202 + int32(1)
		goto L312
	} else {
		goto L314
	}
L313:
	;
	v1208 = v1202
	goto L300
L314:
	;
	goto L313
L315:
	;
	v1240 = base.B2i32(v1147 != v1150)
	goto L293
L316:
	;
	v1222 = v1216 + int32(1)
	v1225 = v1152 + v1222
	v1226 = v1155 + v1222
	v1227 = F___get_tp(m)
	mBase = m.M
	v1228 = F___strcoll_l(m, v1225, v1226, v1226)
	mBase = m.M
	goto L318
L317:
	;
	v1240 = int32(-1)
	goto L293
L318:
	;
	if v1228 == int32(0) {
		v1147 = v1147 - v1222
		v1150 = v1150 - v1222
		v1152 = v1225
		v1155 = v1226
		goto L296
	} else {
		goto L319
	}
L319:
	;
	goto L297
L320:
	;
	if v1254 != int32(-1) {
		v1273 = v1254
		goto L281
	} else {
		goto L321
	}
L321:
	;
	v1259 = F_call_orderTM(m, l0, v1124, v1115, int32(13))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L13
	} else {
		goto L322
	}
L322:
	;
	if v1259 == int32(-1) {
		goto L282
	} else {
		goto L323
	}
L323:
	;
	v1273 = base.B2i32(v1259 == int32(0))
	goto L281
L324:
	;
	v1273 = v1266
	goto L281
L325:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v1300
	v60 = v1297 + int32(4)
	goto L4
L326:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v1297 = v75 + int32(base.Ui32(v1289)>>(uint(int32(12))%32))&int32(1048572) + int32(-524284)
	goto L325
L327:
	;
	if v1307 == int32(base.Ui32(v76)>>(uint(int32(14))%32))&int32(511) {
		v1321 = v75
		goto L330
	} else {
		goto L331
	}
L328:
	;
	v1307 = int32(0)
	goto L327
L329:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v1307 = base.B2i32(v1303 == int32(0))
	goto L327
L330:
	;
	v59 = v152
	v60 = v1321 + int32(4)
	goto L4
L331:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v1321 = v75 + int32(base.Ui32(v1313)>>(uint(int32(12))%32))&int32(1048572) + int32(-524284)
	goto L330
L332:
	;
	if v1335 == int32(base.Ui32(v76)>>(uint(int32(14))%32))&int32(511) {
		v1352 = v75
		goto L335
	} else {
		goto L336
	}
L333:
	;
	v1335 = int32(0)
	goto L332
L334:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1329)))
	v1335 = base.B2i32(v1331 == int32(0))
	goto L332
L335:
	;
	v59 = v152
	v60 = v1352 + int32(4)
	goto L4
L336:
	;
	v1341 = *(*int64)(unsafe.Add(mBase, uint32(v1329)))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v1330
	*(*int64)(unsafe.Add(mBase, uint32(v163))) = v1341
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v1352 = v75 + int32(base.Ui32(v1344)>>(uint(int32(12))%32))&int32(1048572) + int32(-524284)
	goto L335
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	v1371 = F_luaD_precall(m, l0, v163, v1359+int32(-1))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L13
	} else {
		goto L340
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v163 + int32(base.Ui32(v76)>>(uint(int32(19))%32))&int32(8176)
	goto L337
L339:
	;
	if v1359 == int32(0) {
		goto L343
	} else {
		goto L344
	}
L340:
	;
	if v1371 == int32(1) {
		goto L339
	} else {
		goto L341
	}
L341:
	;
	if v1371 != 0 {
		goto L3
	} else {
		goto L342
	}
L342:
	;
	v26 = v26 + int32(1)
	goto L1
L343:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v1382
	v60 = v75
	goto L4
L344:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1379)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1380
	goto L343
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	v1393 = F_luaD_precall(m, l0, v163, int32(-1))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L13
	} else {
		goto L348
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v163 + int32(base.Ui32(v76)>>(uint(int32(19))%32))&int32(8176)
	goto L345
L347:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v1496
	v60 = v75
	goto L4
L348:
	;
	if v1393 == int32(1) {
		goto L347
	} else {
		goto L349
	}
L349:
	;
	if v1393 != 0 {
		goto L3
	} else {
		goto L350
	}
L350:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1399 = v1397 + int32(-24)
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1397)+4))
	v1402 = v1397 + int32(-20)
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1402)))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1404 == int32(0) {
		v1411 = v1403
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v1397)))
	v1414 = v1411 + (v1412 - v1400)
	*(*int32)(unsafe.Add(mBase, uint32(v1399))) = v1414
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1414
	v1417 = int32(0)
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v1418) <= base.Ui32(v1400) {
		v1461 = v1417
		goto L354
	} else {
		goto L355
	}
L352:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1399)))
	F_luaF_close(m, l0, v1407)
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L13
	} else {
		goto L353
	}
L353:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v1402)))
	v1411 = v1410
	goto L351
L354:
	;
	v1477 = v1403 + v1461<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v1397+int32(-16)))) = v1477
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1397+int32(-12)))) = v1484
	v1487 = v1397 + int32(-4)
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1487)))
	*(*int32)(unsafe.Add(mBase, uint32(v1487))) = v1488 + int32(1)
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1492 + int32(-24)
	goto L1
L355:
	;
	v1426 = v1417
	goto L356
L356:
	;
	v1440 = int32(4)
	v1441 = v1426 << (uint(v1440) % 32)
	v1442 = v1403 + v1441
	v1443 = v1400 + v1441
	v1444 = *(*int64)(unsafe.Add(mBase, uint32(v1443)))
	*(*int64)(unsafe.Add(mBase, uint32(v1442))) = v1444
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1442)+8)) = v1446
	v1449 = v1426 + int32(1)
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v1400+v1449<<(uint(v1440)%32)) < base.Ui32(v1453) {
		v1426 = v1449
		goto L356
	} else {
		goto L358
	}
L357:
	;
	v1461 = v1449
	goto L354
L358:
	;
	goto L357
L359:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1507 == int32(0) {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v163 + int32(base.Ui32(v76)>>(uint(int32(19))%32))&int32(8176) + int32(-16)
	goto L359
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	v1513 = F_luaD_poscall(m, l0, v163)
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L13
	} else {
		goto L364
	}
L362:
	;
	F_luaF_close(m, l0, v152)
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L13
	} else {
		goto L363
	}
L363:
	;
	goto L361
L364:
	;
	v1516 = v26 + int32(-1)
	if v1516 == int32(0) {
		goto L3
	} else {
		goto L365
	}
L365:
	;
	if v1513 == int32(0) {
		v26 = v1516
		goto L1
	} else {
		goto L366
	}
L366:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1521)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1522
	v26 = v1516
	goto L1
L367:
	;
	v1538 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+56)) = v1538
	*(*float64)(unsafe.Add(mBase, uint32(v163)+48)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v1538
	*(*float64)(unsafe.Add(mBase, uint32(v163))) = v1526
	v59 = v152
	v60 = v75 + int32(base.Ui32(v76)>>(uint(int32(12))%32))&int32(1048572) + int32(-524284)
	goto L4
L368:
	;
	if base.F64_le(v1527, v1526) == int32(0) {
		v59 = v152
		v60 = v75
		goto L4
	} else {
		goto L371
	}
L369:
	;
	if base.F64_le(v1526, v1527) == int32(0) {
		v59 = v152
		v60 = v75
		goto L4
	} else {
		goto L370
	}
L370:
	;
	goto L367
L371:
	;
	goto L367
L372:
	;
	v1608 = m.G3
	F_luaG_runerror(m, l0, v1608+int32(_a2277), int32(0))
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L13
	} else {
		goto L389
	}
L373:
	;
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v163)+24))
	if v1569 == int32(3) {
		goto L378
	} else {
		goto L379
	}
L374:
	;
	if v1552 != int32(4) {
		goto L372
	} else {
		goto L375
	}
L375:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v1560 = F_luaO_str2d(m, v1557+int32(16), v23)
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L13
	} else {
		goto L376
	}
L376:
	;
	if v1560 == int32(0) {
		goto L372
	} else {
		goto L377
	}
L377:
	;
	v1564 = *(*float64)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v163))) = v1564
	goto L373
L378:
	;
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v163)+40))
	if v1586 == int32(3) {
		goto L28
	} else {
		goto L383
	}
L379:
	;
	if v1569 != int32(4) {
		goto L29
	} else {
		goto L380
	}
L380:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	v1577 = F_luaO_str2d(m, v1574+int32(16), v23)
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L13
	} else {
		goto L381
	}
L381:
	;
	if v1577 == int32(0) {
		goto L29
	} else {
		goto L382
	}
L382:
	;
	v1581 = *(*float64)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+24)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v163)+16)) = v1581
	goto L378
L383:
	;
	if v1586 != int32(4) {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1602 = m.G3
	F_luaG_runerror(m, l0, v1602+int32(_a2278), int32(0))
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L13
	} else {
		goto L388
	}
L385:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v163)+32))
	v1594 = F_luaO_str2d(m, v1591+int32(16), v23)
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L13
	} else {
		goto L386
	}
L386:
	;
	if v1594 == int32(0) {
		goto L384
	} else {
		goto L387
	}
L387:
	;
	v1598 = *(*float64)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+40)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v163)+32)) = v1598
	goto L28
L388:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L389:
	;
	goto L28
L390:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1639
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1644 = v1641 + v160<<(uint(int32(4))%32)
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1644)+56))
	if v1645 == int32(0) {
		v1659 = v75
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v59 = v1641
	v60 = v1659 + int32(4)
	goto L4
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1644)+40)) = v1645
	v1649 = *(*int64)(unsafe.Add(mBase, uint32(v1644)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1644)+32)) = v1649
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v1659 = v75 + int32(base.Ui32(v1651)>>(uint(int32(12))%32))&int32(1048572) + int32(-524284)
	goto L391
L393:
	;
	if v1665 != 0 {
		v1684 = v75
		v1685 = v1665
		goto L396
	} else {
		goto L397
	}
L394:
	;
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1671)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1672
	v1680 = (v1670-v163)>>(uint(int32(4))%32) + int32(-1)
	goto L393
L395:
	;
	v1680 = int32(base.Ui32(v76) >> (uint(int32(23)) % 32))
	goto L393
L396:
	;
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	if v1686 != int32(5) {
		v59 = v152
		v60 = v1684
		goto L4
	} else {
		goto L398
	}
L397:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v1684 = v60 + int32(8)
	v1685 = v1683
	goto L396
L398:
	;
	v1693 = v1680 + v1685*int32(50) + int32(-50)
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+36))
	if v1693 <= v1695 {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	if v1680 < int32(1) {
		v59 = v152
		v60 = v1684
		goto L4
	} else {
		goto L402
	}
L400:
	;
	F_luaH_resizearray(m, l0, v1694, v1693)
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L13
	} else {
		goto L401
	}
L401:
	;
	goto L399
L402:
	;
	v1710 = v1693
	v1712 = v1680
	goto L403
L403:
	;
	v1721 = F_luaH_setnum(m, l0, v1694, v1710)
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L13
	} else {
		goto L405
	}
L405:
	;
	v1723 = int32(4)
	v1725 = v163 + v1712<<(uint(v1723)%32)
	v1726 = *(*int64)(unsafe.Add(mBase, uint32(v1725)))
	*(*int64)(unsafe.Add(mBase, uint32(v1721))) = v1726
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1725)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1721)+8)) = v1728
	if v1728 < v1723 {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v1751 = int32(-1)
	if v1712 <= int32(1) {
		v59 = v152
		v60 = v1684
		goto L4
	} else {
		goto L411
	}
L407:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v1725)))
	v1733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1732)+5)))
	if v1733&int32(3) == int32(0) {
		goto L406
	} else {
		goto L408
	}
L408:
	;
	v1738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1694)+5)))
	if v1738&int32(4) == int32(0) {
		goto L406
	} else {
		goto L409
	}
L409:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1694)+5)))
	v1746 = v1744 & int32(251)
	*(*uint8)(unsafe.Add(mBase, uint32(v1694)+5)) = uint8(v1746)
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v1743)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1694)+32)) = v1748
	*(*int32)(unsafe.Add(mBase, uint32(v1743)+40)) = v1694
	goto L410
L410:
	;
	goto L406
L411:
	;
	v1710 = v1710 + v1751
	v1712 = v1712 + v1751
	goto L403
L412:
	;
	v59 = v152
	v60 = v75
	goto L4
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1769)+16)) = v1766
	if v1767 == int32(0) {
		v1829 = v75
		goto L414
	} else {
		goto L415
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v1769
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v1829
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1847)+68))
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1847)+64))
	if base.Ui32(v1848) < base.Ui32(v1849) {
		goto L423
	} else {
		goto L424
	}
L415:
	;
	v1783 = v75
	v1786 = int32(0)
	goto L416
L416:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v1783)))
	v1802 = int32(base.Ui32(v1800) >> (uint(int32(23)) % 32))
	if v1800&int32(63) != int32(4) {
		goto L419
	} else {
		goto L420
	}
L417:
	;
	v1829 = v1819
	goto L414
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1769+int32(20)+v1786<<(uint(int32(2))%32)))) = v1816
	v1819 = v1783 + int32(4)
	v1821 = v1786 + int32(1)
	if v1821 != v1767 {
		v1783 = v1819
		v1786 = v1821
		goto L416
	} else {
		goto L422
	}
L419:
	;
	v1814 = F_luaF_findupval(m, l0, v152+v1802<<(uint(int32(4))%32))
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L13
	} else {
		goto L421
	}
L420:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v49+v1802<<(uint(int32(2))%32))))
	v1816 = v1810
	goto L418
L421:
	;
	v1816 = v1814
	goto L418
L422:
	;
	goto L417
L423:
	;
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v1853
	v60 = v1829
	goto L4
L424:
	;
	F_luaC_step(m, l0)
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L13
	} else {
		goto L425
	}
L425:
	;
	goto L423
L426:
	;
	if v1890 < int32(1) {
		v59 = v1888
		v60 = v75
		goto L4
	} else {
		goto L432
	}
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v75
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1864<<(uint(int32(4))%32) < v1872-v1873 {
		goto L429
	} else {
		goto L430
	}
L428:
	;
	v1888 = v152
	v1889 = v163
	v1890 = int32(base.Ui32(v76)>>(uint(int32(23))%32)) + int32(-1)
	goto L426
L429:
	;
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1881 = int32(4)
	v1883 = v1880 + v160<<(uint(v1881)%32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1883 + v1864<<(uint(v1881)%32)
	v1888 = v1880
	v1889 = v1883
	v1890 = v1864
	goto L426
L430:
	;
	F_luaD_growstack(m, l0, v1864)
	mBase = m.M
	v1879 = m.ExcPending
	if v1879 != 0 {
		goto L13
	} else {
		goto L431
	}
L431:
	;
	goto L429
L432:
	;
	v1893 = int32(0)
	v1905 = v1893
	goto L433
L433:
	;
	if v1864 <= v1905 {
		goto L436
	} else {
		goto L437
	}
L435:
	;
	v1937 = v1905 + int32(1)
	if v1937 == v1890 {
		v59 = v1888
		v60 = v75
		goto L4
	} else {
		goto L438
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1889+v1905<<(uint(int32(4))%32))+8)) = int32(0)
	goto L435
L437:
	;
	v1917 = int32(4)
	v1918 = v1905 << (uint(v1917) % 32)
	v1919 = v1889 + v1918
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v1854)))
	v1924 = v1920 + (v1893-v1864)<<(uint(v1917)%32) + v1918
	v1925 = *(*int64)(unsafe.Add(mBase, uint32(v1924)))
	*(*int64)(unsafe.Add(mBase, uint32(v1919))) = v1925
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1924)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1919)+8)) = v1927
	goto L435
L438:
	;
	v1905 = v1937
	goto L433
L439:
	;
	goto L28
}
