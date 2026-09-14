package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_rdbLoadObject(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int64) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int64
	_ = v73
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
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
	var v128 int64
	_ = v128
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
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
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
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
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int64
	_ = v435
	var v438 int64
	_ = v438
	var v439 int64
	_ = v439
	var v440 int64
	_ = v440
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v483 int64
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int64
	_ = v536
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v578 float64
	_ = v578
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int64
	_ = v666
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int64
	_ = v720
	var v723 int64
	_ = v723
	var v724 int64
	_ = v724
	var v725 int64
	_ = v725
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v734 int64
	_ = v734
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v761 int64
	_ = v761
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
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
	var v822 int64
	_ = v822
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1032 int64
	_ = v1032
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1077 int32
	_ = v1077
	var v1088 int64
	_ = v1088
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1113 int64
	_ = v1113
	var v1121 int32
	_ = v1121
	var v1132 int32
	_ = v1132
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int64
	_ = v1150
	var v1152 int64
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1164 int64
	_ = v1164
	var v1167 int64
	_ = v1167
	var v1185 int64
	_ = v1185
	var v1204 int64
	_ = v1204
	var v1208 int64
	_ = v1208
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1256 int64
	_ = v1256
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1287 int32
	_ = v1287
	var v1291 int64
	_ = v1291
	var v1297 int32
	_ = v1297
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1338 int64
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1356 int64
	_ = v1356
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1384 int32
	_ = v1384
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1405 int64
	_ = v1405
	var v1408 int64
	_ = v1408
	var v1409 int64
	_ = v1409
	var v1410 int64
	_ = v1410
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1424 int32
	_ = v1424
	var v1431 int64
	_ = v1431
	var v1435 int64
	_ = v1435
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1454 int64
	_ = v1454
	var v1457 int64
	_ = v1457
	var v1461 int32
	_ = v1461
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1470 int64
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1490 int64
	_ = v1490
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1556 int64
	_ = v1556
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1589 int32
	_ = v1589
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1635 int32
	_ = v1635
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1648 int32
	_ = v1648
	var v1651 int32
	_ = v1651
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1661 int32
	_ = v1661
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1679 int32
	_ = v1679
	var v1687 int32
	_ = v1687
	var v1691 int32
	_ = v1691
	var v1693 int32
	_ = v1693
	var v1701 int32
	_ = v1701
	var v1716 int32
	_ = v1716
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1743 int32
	_ = v1743
	var v1747 int32
	_ = v1747
	var v1754 int32
	_ = v1754
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1764 int32
	_ = v1764
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1785 int32
	_ = v1785
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1797 int32
	_ = v1797
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1812 int32
	_ = v1812
	var v1814 int32
	_ = v1814
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1887 int32
	_ = v1887
	var v1891 int32
	_ = v1891
	var v1898 int32
	_ = v1898
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1908 int32
	_ = v1908
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1929 int32
	_ = v1929
	var v1931 int32
	_ = v1931
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1941 int32
	_ = v1941
	var v1946 int32
	_ = v1946
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1956 int32
	_ = v1956
	var v1958 int32
	_ = v1958
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1976 int32
	_ = v1976
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2027 int32
	_ = v2027
	var v2029 int32
	_ = v2029
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2042 int32
	_ = v2042
	var v2044 int32
	_ = v2044
	var v2046 int32
	_ = v2046
	var v2055 int32
	_ = v2055
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2069 int32
	_ = v2069
	var v2071 int64
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2083 int32
	_ = v2083
	var v2092 int32
	_ = v2092
	var v2101 int32
	_ = v2101
	var v2106 int64
	_ = v2106
	var v2107 int64
	_ = v2107
	var v2108 int64
	_ = v2108
	var v2109 int64
	_ = v2109
	var v2119 int32
	_ = v2119
	var v2123 int64
	_ = v2123
	var v2128 int64
	_ = v2128
	var v2132 int64
	_ = v2132
	var v2136 int64
	_ = v2136
	var v2137 int64
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2141 int32
	_ = v2141
	var v2151 int32
	_ = v2151
	var v2165 int32
	_ = v2165
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2182 int32
	_ = v2182
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2189 int64
	_ = v2189
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2202 int32
	_ = v2202
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2229 int32
	_ = v2229
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2256 int32
	_ = v2256
	var v2258 int32
	_ = v2258
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2267 int32
	_ = v2267
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2282 int32
	_ = v2282
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2294 int64
	_ = v2294
	var v2298 int32
	_ = v2298
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2307 int32
	_ = v2307
	var v2309 int32
	_ = v2309
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2320 int32
	_ = v2320
	var v2322 int32
	_ = v2322
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2335 int32
	_ = v2335
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2344 int32
	_ = v2344
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2358 int32
	_ = v2358
	var v2360 int32
	_ = v2360
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2378 int32
	_ = v2378
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2390 int64
	_ = v2390
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2418 int32
	_ = v2418
	var v2422 int32
	_ = v2422
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2439 int32
	_ = v2439
	var v2441 int32
	_ = v2441
	var v2443 int32
	_ = v2443
	var v2445 int32
	_ = v2445
	var v2447 int32
	_ = v2447
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2462 int64
	_ = v2462
	var v2466 int32
	_ = v2466
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2474 int64
	_ = v2474
	var v2484 int64
	_ = v2484
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2509 int32
	_ = v2509
	var v2515 int32
	_ = v2515
	var v2522 int32
	_ = v2522
	var v2525 int32
	_ = v2525
	var v2528 int32
	_ = v2528
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2534 int32
	_ = v2534
	var v2542 int32
	_ = v2542
	var v2544 int32
	_ = v2544
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2555 int32
	_ = v2555
	var v2557 int32
	_ = v2557
	var v2558 int32
	_ = v2558
	var v2560 int64
	_ = v2560
	var v2564 int32
	_ = v2564
	var v2567 int32
	_ = v2567
	var v2568 int32
	_ = v2568
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2582 int32
	_ = v2582
	var v2583 int32
	_ = v2583
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2589 int32
	_ = v2589
	var v2593 int64
	_ = v2593
	var v2601 int32
	_ = v2601
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2630 int64
	_ = v2630
	var v2633 int64
	_ = v2633
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2641 int64
	_ = v2641
	var v2644 int64
	_ = v2644
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2652 int64
	_ = v2652
	var v2655 int64
	_ = v2655
	var v2662 int32
	_ = v2662
	var v2663 int32
	_ = v2663
	var v2665 int64
	_ = v2665
	var v2668 int64
	_ = v2668
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2676 int64
	_ = v2676
	var v2679 int64
	_ = v2679
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2687 int64
	_ = v2687
	var v2690 int64
	_ = v2690
	var v2695 int32
	_ = v2695
	var v2696 int32
	_ = v2696
	var v2698 int64
	_ = v2698
	var v2701 int64
	_ = v2701
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2709 int64
	_ = v2709
	var v2712 int64
	_ = v2712
	var v2714 int64
	_ = v2714
	var v2720 int64
	_ = v2720
	var v2722 int32
	_ = v2722
	var v2727 int32
	_ = v2727
	var v2729 int32
	_ = v2729
	var v2734 int32
	_ = v2734
	var v2739 int32
	_ = v2739
	var v2740 int64
	_ = v2740
	var v2746 int32
	_ = v2746
	var v2747 int64
	_ = v2747
	var v2750 int64
	_ = v2750
	var v2751 int64
	_ = v2751
	var v2752 int64
	_ = v2752
	var v2759 int32
	_ = v2759
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2767 int64
	_ = v2767
	var v2786 int64
	_ = v2786
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2807 int32
	_ = v2807
	var v2811 int32
	_ = v2811
	var v2812 int32
	_ = v2812
	var v2814 int64
	_ = v2814
	var v2817 int64
	_ = v2817
	var v2822 int32
	_ = v2822
	var v2823 int32
	_ = v2823
	var v2825 int64
	_ = v2825
	var v2828 int64
	_ = v2828
	var v2830 int32
	_ = v2830
	var v2835 int32
	_ = v2835
	var v2840 int32
	_ = v2840
	var v2842 int32
	_ = v2842
	var v2846 int32
	_ = v2846
	var v2847 int32
	_ = v2847
	var v2848 int32
	_ = v2848
	var v2852 int64
	_ = v2852
	var v2855 int64
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2861 int32
	_ = v2861
	var v2863 int32
	_ = v2863
	var v2865 int32
	_ = v2865
	var v2873 int64
	_ = v2873
	var v2879 int64
	_ = v2879
	var v2880 int64
	_ = v2880
	var v2883 int64
	_ = v2883
	var v2886 int64
	_ = v2886
	var v2887 int64
	_ = v2887
	var v2892 int64
	_ = v2892
	var v2893 int64
	_ = v2893
	var v2896 int64
	_ = v2896
	var v2898 int64
	_ = v2898
	var v2900 int64
	_ = v2900
	var v2902 int64
	_ = v2902
	var v2903 int64
	_ = v2903
	var v2904 int64
	_ = v2904
	var v2906 int64
	_ = v2906
	var v2907 int64
	_ = v2907
	var v2909 int64
	_ = v2909
	var v2912 int64
	_ = v2912
	var v2915 int64
	_ = v2915
	var v2916 int64
	_ = v2916
	var v2919 int64
	_ = v2919
	var v2920 int64
	_ = v2920
	var v2926 int32
	_ = v2926
	var v2928 int64
	_ = v2928
	var v2933 int64
	_ = v2933
	var v2934 int64
	_ = v2934
	var v2939 int32
	_ = v2939
	var v2942 int32
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2944 int64
	_ = v2944
	var v2948 int64
	_ = v2948
	var v2949 int64
	_ = v2949
	var v2955 int32
	_ = v2955
	var v2958 int32
	_ = v2958
	var v2961 int64
	_ = v2961
	var v2964 int32
	_ = v2964
	var v2967 int32
	_ = v2967
	var v2968 int64
	_ = v2968
	var v2980 int64
	_ = v2980
	var v2990 int64
	_ = v2990
	var v2992 int64
	_ = v2992
	var v2998 int32
	_ = v2998
	var v3005 int32
	_ = v3005
	var v3008 int32
	_ = v3008
	var v3011 int32
	_ = v3011
	var v3014 int32
	_ = v3014
	var v3015 int32
	_ = v3015
	var v3017 int32
	_ = v3017
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3023 int32
	_ = v3023
	var v3031 int32
	_ = v3031
	var v3039 int32
	_ = v3039
	var v3041 int32
	_ = v3041
	var v3043 int32
	_ = v3043
	var v3046 int32
	_ = v3046
	var v3050 int32
	_ = v3050
	var v3051 int32
	_ = v3051
	var v3054 int64
	_ = v3054
	var v3065 int64
	_ = v3065
	var v3084 int32
	_ = v3084
	var v3085 int32
	_ = v3085
	var v3086 int32
	_ = v3086
	var v3091 int32
	_ = v3091
	var v3093 int32
	_ = v3093
	var v3094 int32
	_ = v3094
	var v3095 int64
	_ = v3095
	var v3096 int32
	_ = v3096
	var v3101 int32
	_ = v3101
	var v3102 int32
	_ = v3102
	var v3104 int64
	_ = v3104
	var v3107 int64
	_ = v3107
	var v3109 int32
	_ = v3109
	var v3114 int32
	_ = v3114
	var v3119 int32
	_ = v3119
	var v3120 int32
	_ = v3120
	var v3125 int32
	_ = v3125
	var v3126 int32
	_ = v3126
	var v3132 int32
	_ = v3132
	var v3134 int64
	_ = v3134
	var v3162 int32
	_ = v3162
	var v3163 int32
	_ = v3163
	var v3166 int64
	_ = v3166
	var v3174 int32
	_ = v3174
	var v3179 int32
	_ = v3179
	var v3193 int64
	_ = v3193
	var v3204 int32
	_ = v3204
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3211 int32
	_ = v3211
	var v3212 int32
	_ = v3212
	var v3215 int32
	_ = v3215
	var v3216 int32
	_ = v3216
	var v3218 int32
	_ = v3218
	var v3224 int32
	_ = v3224
	var v3225 int64
	_ = v3225
	var v3226 int32
	_ = v3226
	var v3228 int32
	_ = v3228
	var v3233 int32
	_ = v3233
	var v3238 int32
	_ = v3238
	var v3239 int64
	_ = v3239
	var v3240 int32
	_ = v3240
	var v3242 int32
	_ = v3242
	var v3247 int32
	_ = v3247
	var v3252 int32
	_ = v3252
	var v3257 int32
	_ = v3257
	var v3258 int32
	_ = v3258
	var v3261 int64
	_ = v3261
	var v3272 int64
	_ = v3272
	var v3291 int32
	_ = v3291
	var v3292 int32
	_ = v3292
	var v3293 int32
	_ = v3293
	var v3298 int32
	_ = v3298
	var v3299 int32
	_ = v3299
	var v3301 int32
	_ = v3301
	var v3302 int32
	_ = v3302
	var v3304 int32
	_ = v3304
	var v3313 int32
	_ = v3313
	var v3314 int32
	_ = v3314
	var v3325 int32
	_ = v3325
	var v3326 int32
	_ = v3326
	var v3328 int32
	_ = v3328
	var v3334 int32
	_ = v3334
	var v3335 int32
	_ = v3335
	var v3336 int32
	_ = v3336
	var v3341 int32
	_ = v3341
	var v3351 int32
	_ = v3351
	var v3354 int32
	_ = v3354
	var v3357 int32
	_ = v3357
	var v3359 int32
	_ = v3359
	var v3361 int32
	_ = v3361
	var v3362 int32
	_ = v3362
	var v3364 int32
	_ = v3364
	var v3374 int32
	_ = v3374
	var v3377 int32
	_ = v3377
	var v3382 int32
	_ = v3382
	var v3385 int32
	_ = v3385
	var v3398 int32
	_ = v3398
	var v3400 int32
	_ = v3400
	var v3402 int32
	_ = v3402
	var v3407 int32
	_ = v3407
	var v3414 int32
	_ = v3414
	var v3419 int32
	_ = v3419
	var v3429 int32
	_ = v3429
	var v3430 int32
	_ = v3430
	var v3438 int32
	_ = v3438
	var v3439 int32
	_ = v3439
	var v3441 int32
	_ = v3441
	var v3444 int32
	_ = v3444
	var v3448 int32
	_ = v3448
	var v3454 int32
	_ = v3454
	var v3455 int32
	_ = v3455
	var v3457 int32
	_ = v3457
	var v3461 int32
	_ = v3461
	var v3463 int32
	_ = v3463
	var v3469 int32
	_ = v3469
	var v3473 int32
	_ = v3473
	var v3479 int32
	_ = v3479
	var v3480 int32
	_ = v3480
	var v3490 int32
	_ = v3490
	var v3494 int32
	_ = v3494
	var v3495 int32
	_ = v3495
	var v3498 int32
	_ = v3498
	var v3505 int32
	_ = v3505
	var v3506 int32
	_ = v3506
	var v3507 int32
	_ = v3507
	var v3516 int32
	_ = v3516
	var v3518 int32
	_ = v3518
	var v3523 int32
	_ = v3523
	var v3524 int32
	_ = v3524
	var v3528 int64
	_ = v3528
	var v3556 int64
	_ = v3556
	var v3564 int32
	_ = v3564
	var v3566 int32
	_ = v3566
	var v3571 int32
	_ = v3571
	var v3573 int32
	_ = v3573
	var v3578 int32
	_ = v3578
	var v3602 int64
	_ = v3602
	var v3604 int32
	_ = v3604
	var v3605 int32
	_ = v3605
	var v3611 int64
	_ = v3611
	var v3626 int32
	_ = v3626
	var v3628 int32
	_ = v3628
	var v3629 int32
	_ = v3629
	var v3654 int32
	_ = v3654
	var v3655 int32
	_ = v3655
	var v3658 int32
	_ = v3658
	var v3659 int32
	_ = v3659
	var v3663 int32
	_ = v3663
	var v3669 int32
	_ = v3669
	var v3673 int32
	_ = v3673
	var v3679 int32
	_ = v3679
	var v3684 int32
	_ = v3684
	var v3690 int32
	_ = v3690
	var v3692 int32
	_ = v3692
	var v3695 int32
	_ = v3695
	var v3701 int32
	_ = v3701
	var v3706 int32
	_ = v3706
	var v3707 int32
	_ = v3707
	var v3711 int32
	_ = v3711
	var v3712 int32
	_ = v3712
	var v3713 int32
	_ = v3713
	var v3718 int32
	_ = v3718
	var v3723 int32
	_ = v3723
	var v3725 int64
	_ = v3725
	var v3728 int64
	_ = v3728
	var v3729 int32
	_ = v3729
	var v3730 int32
	_ = v3730
	var v3731 int32
	_ = v3731
	var v3732 int32
	_ = v3732
	var v3736 int32
	_ = v3736
	var v3737 int32
	_ = v3737
	var v3740 int32
	_ = v3740
	var v3743 int32
	_ = v3743
	var v3746 int32
	_ = v3746
	var v3749 int32
	_ = v3749
	var v3756 int32
	_ = v3756
	var v3763 int32
	_ = v3763
	var v3771 int32
	_ = v3771
	var v3779 int32
	_ = v3779
	var v3787 int32
	_ = v3787
	var v3795 int32
	_ = v3795
	var v3803 int32
	_ = v3803
	var v3809 int32
	_ = v3809
	var v3813 int32
	_ = v3813
	var v3814 int32
	_ = v3814
	var v3816 int32
	_ = v3816
	var v3817 int32
	_ = v3817
	var v3820 int32
	_ = v3820
	var v3823 int32
	_ = v3823
	var v3826 int32
	_ = v3826
	var v3829 int32
	_ = v3829
	var v3836 int32
	_ = v3836
	var v3843 int32
	_ = v3843
	var v3851 int32
	_ = v3851
	var v3859 int32
	_ = v3859
	var v3867 int32
	_ = v3867
	var v3875 int32
	_ = v3875
	var v3883 int32
	_ = v3883
	var v3889 int32
	_ = v3889
	var v3900 int32
	_ = v3900
	var v3904 int32
	_ = v3904
	var v3922 int32
	_ = v3922
	var v3923 int32
	_ = v3923
	var v3924 int32
	_ = v3924
	var v3925 int32
	_ = v3925
	var v3929 int32
	_ = v3929
	var v3930 int32
	_ = v3930
	var v3932 int32
	_ = v3932
	var v3936 int32
	_ = v3936
	var v3937 int32
	_ = v3937
	var v3939 int64
	_ = v3939
	var v3942 int64
	_ = v3942
	var v3943 int64
	_ = v3943
	var v3944 int64
	_ = v3944
	var v3950 int32
	_ = v3950
	var v3951 int32
	_ = v3951
	var v3953 int32
	_ = v3953
	var v3955 int32
	_ = v3955
	var v3959 int32
	_ = v3959
	var v3962 int32
	_ = v3962
	var v3964 int32
	_ = v3964
	var v3972 int32
	_ = v3972
	var v3975 int32
	_ = v3975
	var v3976 int32
	_ = v3976
	var v3978 int32
	_ = v3978
	var v3980 int32
	_ = v3980
	var v3984 int32
	_ = v3984
	var v3987 int32
	_ = v3987
	var v3989 int32
	_ = v3989
	var v3997 int32
	_ = v3997
	var v3999 int32
	_ = v3999
	var v4000 int32
	_ = v4000
	var v4010 int32
	_ = v4010
	var v4027 int32
	_ = v4027
	var v4033 int32
	_ = v4033
	var v4035 int32
	_ = v4035
	var v4037 int32
	_ = v4037
	var v4040 int32
	_ = v4040
	var v4043 int32
	_ = v4043
	var v4045 int32
	_ = v4045
	var v4056 int32
	_ = v4056
	var v4070 int32
	_ = v4070
	var v4071 int32
	_ = v4071
	var v4073 int32
	_ = v4073
	var v4074 int32
	_ = v4074
	var v4088 int32
	_ = v4088
	var v4102 int32
	_ = v4102
	var v4176 int32
	_ = v4176
	var v4178 int32
	_ = v4178
	var v4180 int32
	_ = v4180
	var v4183 int32
	_ = v4183
	var v4185 int32
	_ = v4185
	var v4194 int32
	_ = v4194
	v23 = m.G0
	v25 = v23 - int32(544)
	m.G0 = v25
	if l4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	switch l0 {
	case 0:
		goto L32
	case 1:
		goto L31
	case 2:
		goto L30
	default:
		goto L29
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(3)
	goto L1
L3:
	;
	m.G0 = v25 + int32(544)
	return v4194
L4:
	;
	F_decrRefCount(m, v200)
	mBase = m.M
	v4183 = m.ExcPending
	if v4183 != 0 {
		goto L33
	} else {
		goto L1157
	}
L5:
	;
	F_decrRefCount(m, v444)
	mBase = m.M
	v4176 = m.ExcPending
	if v4176 != 0 {
		goto L33
	} else {
		goto L1155
	}
L6:
	;
	v4194 = int32(0)
	goto L3
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	v4194 = int32(0)
	goto L3
L8:
	;
	F_decrRefCount(m, v4088)
	mBase = m.M
	v4102 = m.ExcPending
	if v4102 != 0 {
		goto L33
	} else {
		goto L1153
	}
L9:
	;
	F_decrRefCount(m, v1598)
	mBase = m.M
	v4073 = m.ExcPending
	if v4073 != 0 {
		goto L33
	} else {
		goto L1150
	}
L10:
	;
	F_decrRefCount(m, v4056)
	mBase = m.M
	v4070 = m.ExcPending
	if v4070 != 0 {
		goto L33
	} else {
		goto L1149
	}
L11:
	;
	F_decrRefCount(m, v2451)
	mBase = m.M
	v4043 = m.ExcPending
	if v4043 != 0 {
		goto L33
	} else {
		goto L1147
	}
L12:
	;
	F_sdsfree(m, v2502)
	mBase = m.M
	v4040 = m.ExcPending
	if v4040 != 0 {
		goto L33
	} else {
		goto L1146
	}
L13:
	;
	F_decrRefCount(m, v2451)
	mBase = m.M
	v4035 = m.ExcPending
	if v4035 != 0 {
		goto L33
	} else {
		goto L1144
	}
L14:
	;
	v4027 = int32(0)
	F_rdbReportError(m, int32(1), int32(2920), int32(_a_F_rdbLoadObject_0), v4027)
	mBase = m.M
	v4033 = m.ExcPending
	if v4033 != 0 {
		goto L33
	} else {
		goto L1143
	}
L15:
	;
	if l4 == int32(0) {
		goto L1141
	} else {
		goto L1142
	}
L16:
	;
	v3707 = int32(0)
	v3711 = F_rdbLoadLenByRef(m, l1, v3707, v25+int32(192))
	mBase = m.M
	v3712 = m.ExcPending
	if v3712 != 0 {
		goto L33
	} else {
		goto L1093
	}
L17:
	;
	v3692 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[0]))
	if v3692 != int32(1) {
		goto L1089
	} else {
		goto L1090
	}
L18:
	;
	v2451 = F_createStreamObject(m)
	mBase = m.M
	v2452 = m.ExcPending
	if v2452 != 0 {
		goto L33
	} else {
		goto L759
	}
L19:
	;
	v2431 = *(*int32)(unsafe.Add(mBase, uint32(v25)+512))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v2431
	F_rdbReportError(m, int32(1), int32(2398), int32(_a_F_rdbLoadObject_1), v25+int32(32))
	mBase = m.M
	v2439 = m.ExcPending
	if v2439 != 0 {
		goto L33
	} else {
		goto L753
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = l0
	F_rdbReportError(m, int32(1), int32(2595), int32(_a_F_rdbLoadObject_2), v25+int32(16))
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L33
	} else {
		goto L752
	}
L21:
	;
	v2388 = int32(_a_F_rdbLoadObject_3)
	v2390 = *(*int64)(unsafe.Add(mBase, _c_F_rdbLoadObject[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_rdbLoadObject[1])) = v2390 + int64(1)
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v25)+192))
	v2396 = F_lpValidateIntegrityAndDups(m, v1593, v2394, int32(1))
	mBase = m.M
	v2397 = m.ExcPending
	if v2397 != 0 {
		goto L33
	} else {
		goto L742
	}
L22:
	;
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(v25)+192))
	v2341 = F_lpNew(m, v2340)
	mBase = m.M
	v2342 = m.ExcPending
	if v2342 != 0 {
		goto L33
	} else {
		goto L722
	}
L23:
	;
	v2292 = int32(_a_F_rdbLoadObject_3)
	v2294 = *(*int64)(unsafe.Add(mBase, _c_F_rdbLoadObject[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_rdbLoadObject[1])) = v2294 + int64(1)
	v2298 = *(*int32)(unsafe.Add(mBase, uint32(v25)+192))
	v2300 = F_lpValidateIntegrityAndDups(m, v1593, v2298, int32(1))
	mBase = m.M
	v2301 = m.ExcPending
	if v2301 != 0 {
		goto L33
	} else {
		goto L706
	}
L24:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v25)+192))
	v2226 = F_lpNew(m, v2225)
	mBase = m.M
	v2227 = m.ExcPending
	if v2227 != 0 {
		goto L33
	} else {
		goto L681
	}
L25:
	;
	v2187 = int32(_a_F_rdbLoadObject_3)
	v2189 = *(*int64)(unsafe.Add(mBase, _c_F_rdbLoadObject[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_rdbLoadObject[1])) = v2189 + int64(1)
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(v25)+192))
	v2195 = F_lpValidateIntegrityAndDups(m, v1593, v2193, int32(0))
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L33
	} else {
		goto L668
	}
L26:
	;
	v2069 = int32(_a_F_rdbLoadObject_3)
	v2071 = *(*int64)(unsafe.Add(mBase, _c_F_rdbLoadObject[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_rdbLoadObject[1])) = v2071 + int64(1)
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v25)+192))
	if base.Ui32(v2075) < base.Ui32(int32(8)) {
		goto L638
	} else {
		goto L639
	}
L27:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v1598)))
	*(*int32)(unsafe.Add(mBase, uint32(v1598))) = v2055&int32(-16) | int32(1)
	F_objectSetVal(m, v1598, v2014)
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L33
	} else {
		goto L634
	}
L28:
	;
	if l4 != 0 {
		goto L7
	} else {
		goto L633
	}
L29:
	;
	switch l0 + int32(-3) {
	case 0, 2:
		goto L157
	default:
		goto L156
	}
L30:
	;
	v158 = int32(0)
	v162 = F_rdbLoadLenByRef(m, l1, v158, v25+int32(192))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L33
	} else {
		goto L62
	}
L31:
	;
	v43 = int32(0)
	v47 = F_rdbLoadLenByRef(m, l1, v43, v25+int32(192))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L33
	} else {
		goto L37
	}
L32:
	;
	v31 = int32(0)
	v34 = F_rdbGenericLoadStringObject(m, l1, int32(1), v31)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return int32(0)
L34:
	;
	if v34 == int32(0) {
		v4194 = v31
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v41 = F_tryObjectEncodingEx(m, v34, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v4010 = v41
	goto L15
L37:
	;
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v47 == int32(-1) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v59 = int32(_a_F_rdbLoadObject_3)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[2]))
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[3]))
	v63 = F_createQuicklistObject(m, v60, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L33
	} else {
		goto L43
	}
L39:
	;
	v53 = int64(-1)
	goto L41
L40:
	;
	v53 = v50
	goto L41
L41:
	;
	v54 = int64(1)
	v55 = v53 + v54
	if base.Ui64(v54) < base.Ui64(v55) {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	switch base.I32_wrap_i64(v55) {
	default:
		v4194 = v43
		goto L3
	case 1:
		goto L28
	}
L43:
	;
	if v50 == int64(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v153 = int32(0)
	F_listTypeTryConversion(m, v63, v153, v153, v153)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L33
	} else {
		goto L61
	}
L45:
	;
	v73 = v50
	goto L46
L46:
	;
	v91 = F_rdbGenericLoadStringObject(m, l1, int32(1), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L33
	} else {
		goto L48
	}
L47:
	;
	goto L44
L48:
	;
	if v91 == int32(0) {
		v4056 = v63
		goto L10
	} else {
		goto L49
	}
L49:
	;
	v95 = F_getDecodedObject(m, v91)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L33
	} else {
		goto L50
	}
L50:
	;
	v97 = F_objectGetVal(m, v63)
	mBase = m.M
	v98 = F_objectGetVal(m, v95)
	mBase = m.M
	v100 = F_objectGetVal(m, v95)
	mBase = m.M
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(-1)))))
	switch v103 & int32(7) {
	case 0:
		goto L56
	case 1:
		goto L55
	case 2:
		goto L54
	case 3:
		goto L53
	case 4:
		goto L52
	default:
		v120 = int32(0)
		goto L51
	}
L51:
	;
	v121 = F_quicklistPushTail(m, v97, v98, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L33
	} else {
		goto L57
	}
L52:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-17))))
	v120 = v119
	goto L51
L53:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(-9))))
	v120 = v116
	goto L51
L54:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100+int32(-5)))))
	v120 = v113
	goto L51
L55:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(-3)))))
	v120 = v110
	goto L51
L56:
	;
	v120 = int32(base.Ui32(v103) >> (uint(int32(3)) % 32))
	goto L51
L57:
	;
	F_decrRefCount(m, v95)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L33
	} else {
		goto L58
	}
L58:
	;
	F_decrRefCount(m, v91)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L33
	} else {
		goto L59
	}
L59:
	;
	v128 = v73 + int64(-1)
	if v128 != int64(0) {
		v73 = v128
		goto L46
	} else {
		goto L60
	}
L60:
	;
	goto L47
L61:
	;
	v4010 = v63
	goto L15
L62:
	;
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v162 == int32(-1) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[4]))
	v176 = int32(1073741824)
	if base.Ui32(v175) < base.Ui32(v176) {
		goto L70
	} else {
		goto L71
	}
L64:
	;
	v168 = int64(-1)
	goto L66
L65:
	;
	v168 = v165
	goto L66
L66:
	;
	v169 = int64(1)
	v170 = v168 + v169
	if base.Ui64(v169) < base.Ui64(v170) {
		goto L63
	} else {
		goto L67
	}
L67:
	;
	switch base.I32_wrap_i64(v170) {
	default:
		v4194 = v158
		goto L3
	case 1:
		goto L28
	}
L68:
	;
	v201 = base.I32_wrap_i64(v168)
	v202 = int32(0)
	v207 = v202
	v208 = v202
	v210 = v202
	goto L80
L69:
	;
	v198 = F_createIntsetObject(m)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L33
	} else {
		goto L79
	}
L70:
	;
	v179 = v175
	goto L72
L71:
	;
	v179 = v176
	goto L72
L72:
	;
	if base.Ui64(v168) <= base.Ui64(base.I64_extend_i32_u(v179)) {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v182 = F_createSetObject(m)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L33
	} else {
		goto L74
	}
L74:
	;
	v184 = F_objectGetVal(m, v182)
	mBase = m.M
	v186 = F_hashtableTryExpand(m, v184, base.I32_wrap_i64(v168))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L33
	} else {
		goto L75
	}
L75:
	;
	if v186 != 0 {
		v200 = v182
		goto L68
	} else {
		goto L76
	}
L76:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+176)) = v168
	F_rdbReportError(m, int32(1), int32(1968), int32(_a_F_rdbLoadObject_4), v25+int32(176))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L33
	} else {
		goto L77
	}
L77:
	;
	F_decrRefCount(m, v182)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L33
	} else {
		goto L78
	}
L78:
	;
	v4194 = v158
	goto L3
L79:
	;
	v200 = v198
	goto L68
L80:
	;
	v229 = F_rdbGenericLoadStringObject(m, l1, int32(4), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L33
	} else {
		goto L82
	}
L82:
	;
	if v229 == int32(0) {
		v4056 = v200
		goto L10
	} else {
		goto L83
	}
L83:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+int32(-1)))))
	switch v236 & int32(7) {
	case 0:
		goto L89
	case 1:
		goto L88
	case 2:
		goto L87
	case 3:
		goto L86
	case 4:
		goto L85
	default:
		v253 = int32(0)
		goto L84
	}
L84:
	;
	if base.Ui32(v210) < base.Ui32(v253) {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v229+int32(-17))))
	v253 = v252
	goto L84
L86:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v229+int32(-9))))
	v253 = v249
	goto L84
L87:
	;
	v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229+int32(-5)))))
	v253 = v246
	goto L84
L88:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+int32(-3)))))
	v253 = v243
	goto L84
L89:
	;
	v253 = int32(base.Ui32(v236) >> (uint(int32(3)) % 32))
	goto L84
L90:
	;
	v255 = v253
	goto L92
L91:
	;
	v255 = v210
	goto L92
L92:
	;
	v256 = v253 + v208
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	if v257&int32(240) != int32(96) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	if v343&int32(240) != int32(176) {
		goto L125
	} else {
		goto L126
	}
L94:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+int32(-1)))))
	switch v269 & int32(7) {
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
		v286 = int32(0)
		goto L97
	}
L95:
	;
	v307 = F_setTypeSize(m, v200)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L33
	} else {
		goto L112
	}
L96:
	;
	if v290 != 0 {
		goto L95
	} else {
		goto L106
	}
L97:
	;
	v289 = F_string2ll(m, v229, v286, v25+int32(192))
	mBase = m.M
	if v289 != 0 {
		goto L103
	} else {
		goto L104
	}
L98:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v229+int32(-17))))
	v286 = v285
	goto L97
L99:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v229+int32(-9))))
	v286 = v282
	goto L97
L100:
	;
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229+int32(-5)))))
	v286 = v279
	goto L97
L101:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+int32(-3)))))
	v286 = v276
	goto L97
L102:
	;
	v286 = int32(base.Ui32(v269) >> (uint(int32(3)) % 32))
	goto L97
L103:
	;
	v290 = int32(0)
	goto L105
L104:
	;
	v290 = int32(-1)
	goto L105
L105:
	;
	goto L96
L106:
	;
	v291 = F_objectGetVal(m, v200)
	mBase = m.M
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	v295 = F_intsetAdd(m, v291, v292, v25+int32(496))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L33
	} else {
		goto L107
	}
L107:
	;
	F_objectSetVal(m, v200, v295)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L33
	} else {
		goto L108
	}
L108:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+496)))
	if v299 != 0 {
		goto L93
	} else {
		goto L109
	}
L109:
	;
	v300 = int32(0)
	F_rdbReportError(m, int32(1), int32(1996), int32(_a_F_rdbLoadObject_5), v300)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L33
	} else {
		goto L110
	}
L110:
	;
	v4180 = v300
	goto L4
L111:
	;
	v329 = F_setTypeConvertAndExpand(m, v200, int32(2), v201, int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L33
	} else {
		goto L121
	}
L112:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[5]))
	if base.Ui32(v310) <= base.Ui32(v307) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[6]))
	if base.Ui32(v313) < base.Ui32(v255) {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L118
L115:
	;
	if base.B2i32(base.Ui32(int32(0)+v256) < base.Ui32(int32(1073741825))) == int32(0) {
		goto L111
	} else {
		goto L119
	}
L116:
	;
	goto L115
L118:
	;
	goto L116
L119:
	;
	F_setTypeConvert(m, v200, int32(11))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L33
	} else {
		goto L120
	}
L120:
	;
	goto L93
L121:
	;
	if v329 == int32(0) {
		goto L93
	} else {
		goto L122
	}
L122:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+160)) = v168
	F_rdbReportError(m, int32(1), int32(2008), int32(_a_F_rdbLoadObject_4), v25+int32(160))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L33
	} else {
		goto L123
	}
L123:
	;
	F_sdsfree(m, v229)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L33
	} else {
		goto L124
	}
L124:
	;
	v4056 = v200
	goto L10
L125:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	if v405&int32(240) != int32(32) {
		goto L149
	} else {
		goto L150
	}
L126:
	;
	v348 = F_setTypeSize(m, v200)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L33
	} else {
		goto L129
	}
L127:
	;
	v399 = F_objectGetVal(m, v200)
	mBase = m.M
	v400 = F_lpAppend(m, v399, v229, v253)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L33
	} else {
		goto L146
	}
L128:
	;
	v385 = F_setTypeConvertAndExpand(m, v200, int32(2), v201, int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L33
	} else {
		goto L142
	}
L129:
	;
	v351 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[5]))
	if base.Ui32(v351) <= base.Ui32(v348) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[6]))
	if base.Ui32(v354) < base.Ui32(v253) {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v356 = F_objectGetVal(m, v200)
	mBase = m.M
	if v356 != 0 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	if base.B2i32(base.Ui32(v359+v253) < base.Ui32(int32(1073741825))) == int32(0) {
		goto L128
	} else {
		goto L136
	}
L133:
	;
	goto L132
L134:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v359 = v358
	goto L133
L135:
	;
	v359 = int32(0)
	goto L133
L136:
	;
	v365 = F_objectGetVal(m, v200)
	mBase = m.M
	v366 = F_lpFirst(m, v365)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L33
	} else {
		goto L137
	}
L137:
	;
	if v366 == int32(0) {
		goto L127
	} else {
		goto L138
	}
L138:
	;
	v370 = F_objectGetVal(m, v200)
	mBase = m.M
	v372 = F_lpFind(m, v370, v366, v229, v253, int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L33
	} else {
		goto L139
	}
L139:
	;
	if v372 == int32(0) {
		goto L127
	} else {
		goto L140
	}
L140:
	;
	v376 = int32(0)
	F_rdbReportError(m, int32(1), int32(2022), int32(_a_F_rdbLoadObject_5), v376)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L33
	} else {
		goto L141
	}
L141:
	;
	v4180 = v376
	goto L4
L142:
	;
	if v385 == int32(0) {
		goto L125
	} else {
		goto L143
	}
L143:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+144)) = v168
	F_rdbReportError(m, int32(1), int32(2029), int32(_a_F_rdbLoadObject_4), v25+int32(144))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L33
	} else {
		goto L144
	}
L144:
	;
	F_sdsfree(m, v229)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L33
	} else {
		goto L145
	}
L145:
	;
	v4056 = v200
	goto L10
L146:
	;
	F_objectSetVal(m, v200, v400)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L33
	} else {
		goto L147
	}
L147:
	;
	goto L125
L148:
	;
	v423 = v207 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v423)) < base.Ui64(v168) {
		v207 = v423
		v208 = v256
		v210 = v255
		goto L80
	} else {
		goto L155
	}
L149:
	;
	F_sdsfree(m, v229)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L33
	} else {
		goto L154
	}
L150:
	;
	v410 = F_objectGetVal(m, v200)
	mBase = m.M
	v411 = F_hashtableAdd(m, v410, v229)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L33
	} else {
		goto L151
	}
L151:
	;
	if v411 != 0 {
		goto L148
	} else {
		goto L152
	}
L152:
	;
	v413 = int32(0)
	F_rdbReportError(m, int32(1), int32(2040), int32(_a_F_rdbLoadObject_5), v413)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L33
	} else {
		goto L153
	}
L153:
	;
	v4180 = v413
	goto L4
L154:
	;
	goto L148
L155:
	;
	v4010 = v200
	goto L15
L156:
	;
	switch l0 + int32(-4) {
	case 0, 18:
		goto L240
	default:
		goto L239
	}
L157:
	;
	v428 = int32(0)
	v432 = F_rdbLoadLenByRef(m, l1, v428, v25+int32(192))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L33
	} else {
		goto L158
	}
L158:
	;
	v435 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v432 == int32(-1) {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v444 = F_createZsetObject(m)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L33
	} else {
		goto L167
	}
L160:
	;
	v438 = int64(-1)
	goto L162
L161:
	;
	v438 = v435
	goto L162
L162:
	;
	v439 = int64(1)
	v440 = v438 + v439
	if base.Ui64(v439) < base.Ui64(v440) {
		goto L159
	} else {
		goto L163
	}
L163:
	;
	switch base.I32_wrap_i64(v440) {
	default:
		v4194 = v428
		goto L3
	case 1:
		goto L28
	}
L164:
	;
	v691 = F_zsetLength(m, v444)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L33
	} else {
		goto L230
	}
L165:
	;
	v471 = int32(0)
	v483 = v435
	v485 = v471
	v487 = v471
	goto L173
L166:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+128)) = v438
	F_rdbReportError(m, int32(1), int32(2062), int32(_a_F_rdbLoadObject_4), v25+int32(128))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L33
	} else {
		goto L171
	}
L167:
	;
	v446 = F_objectGetVal(m, v444)
	mBase = m.M
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)))
	v449 = F_hashtableTryExpand(m, v447, base.I32_wrap_i64(v438))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L33
	} else {
		goto L168
	}
L168:
	;
	if v449 == int32(0) {
		goto L166
	} else {
		goto L169
	}
L169:
	;
	if base.B2i32(v435 == int64(0)) == int32(0) {
		goto L165
	} else {
		goto L170
	}
L170:
	;
	v457 = int32(0)
	v681 = v457
	v683 = v457
	goto L164
L171:
	;
	F_decrRefCount(m, v444)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L33
	} else {
		goto L172
	}
L172:
	;
	v4194 = v428
	goto L3
L173:
	;
	v497 = F_rdbGenericLoadStringObject(m, l1, int32(4), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L33
	} else {
		goto L175
	}
L174:
	;
	v681 = v664
	v683 = v646
	goto L164
L175:
	;
	if v497 == int32(0) {
		v4056 = v444
		goto L10
	} else {
		goto L176
	}
L176:
	;
	if l0 != int32(5) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v578 = *(*float64)(unsafe.Add(mBase, uint32(v25)+192))
	if base.Ui64(base.I64_reinterpret_f64(v578)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L198
	} else {
		goto L199
	}
L178:
	;
	v552 = F_rdbLoadDoubleValue(m, l1, v25+int32(192))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L33
	} else {
		goto L196
	}
L179:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v501&int32(5) != 0 {
		goto L5
	} else {
		goto L180
	}
L180:
	;
	v512 = v25 + int32(192)
	v515 = int32(8)
	goto L181
L181:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if base.Ui32(v529) < base.Ui32(v515) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v540 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L184:
	;
	v531 = v529
	goto L186
L185:
	;
	v531 = v515
	goto L186
L186:
	;
	if v529 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v532 = v531
	goto L189
L188:
	;
	v532 = v515
	goto L189
L189:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v534 = m.T0[v533].(func(*base.Module, int32, int32, int32) int32)(m, l1, v512, v532)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L33
	} else {
		goto L190
	}
L190:
	;
	if v534 != 0 {
		goto L183
	} else {
		goto L191
	}
L191:
	;
	v536 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = v536 | int64(1)
	goto L5
L192:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v545 + v532
	v549 = v515 - v532
	if v549 != 0 {
		v512 = v512 + v532
		v515 = v549
		goto L181
	} else {
		goto L195
	}
L193:
	;
	m.T0[v540].(func(*base.Module, int32, int32, int32))(m, l1, v512, v532)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L33
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	goto L177
L196:
	;
	if v552 == int32(-1) {
		goto L5
	} else {
		goto L197
	}
L197:
	;
	goto L177
L198:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497+int32(-1)))))
	v600 = v598 & int32(7)
	switch v600 {
	case 0:
		goto L209
	case 1:
		goto L208
	case 2:
		goto L207
	case 3:
		goto L206
	case 4:
		goto L205
	default:
		v645 = int32(0)
		v646 = v487
		goto L203
	}
L199:
	;
	v584 = int32(0)
	F_rdbReportError(m, int32(1), int32(2093), int32(_a_F_rdbLoadObject_6), v584)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L33
	} else {
		goto L200
	}
L200:
	;
	F_decrRefCount(m, v444)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L33
	} else {
		goto L201
	}
L201:
	;
	F_sdsfree(m, v497)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L33
	} else {
		goto L202
	}
L202:
	;
	v4194 = v584
	goto L3
L203:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v446)+4))
	v648 = F_zslInsert(m, v647, v578, v497)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L33
	} else {
		goto L222
	}
L204:
	;
	if base.Ui32(v615) <= base.Ui32(v487) {
		goto L210
	} else {
		goto L211
	}
L205:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v497+int32(-17))))
	v615 = v614
	goto L204
L206:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v497+int32(-9))))
	v615 = v611
	goto L204
L207:
	;
	v608 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v497+int32(-5)))))
	v615 = v608
	goto L204
L208:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497+int32(-3)))))
	v615 = v605
	goto L204
L209:
	;
	v615 = int32(base.Ui32(v598) >> (uint(int32(3)) % 32))
	goto L204
L210:
	;
	switch v600 {
	default:
		goto L221
	case 1:
		goto L220
	case 2:
		goto L219
	case 3:
		goto L218
	case 4:
		goto L217
	}
L211:
	;
	switch v600 {
	default:
		goto L216
	case 1:
		goto L215
	case 2:
		goto L214
	case 3:
		goto L213
	case 4:
		goto L212
	}
L212:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v497+int32(-17))))
	v645 = v630
	v646 = v630
	goto L203
L213:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v497+int32(-9))))
	v645 = v627
	v646 = v627
	goto L203
L214:
	;
	v624 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v497+int32(-5)))))
	v645 = v624
	v646 = v624
	goto L203
L215:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497+int32(-3)))))
	v645 = v621
	v646 = v621
	goto L203
L216:
	;
	v618 = int32(base.Ui32(v598) >> (uint(int32(3)) % 32))
	v645 = v618
	v646 = v618
	goto L203
L217:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v497+int32(-17))))
	v645 = v644
	v646 = v487
	goto L203
L218:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v497+int32(-9))))
	v645 = v641
	v646 = v487
	goto L203
L219:
	;
	v638 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v497+int32(-5)))))
	v645 = v638
	v646 = v487
	goto L203
L220:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497+int32(-3)))))
	v645 = v635
	v646 = v487
	goto L203
L221:
	;
	v645 = int32(base.Ui32(v598) >> (uint(int32(3)) % 32))
	v646 = v487
	goto L203
L222:
	;
	F_sdsfree(m, v497)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L33
	} else {
		goto L223
	}
L223:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v446)))
	v653 = F_hashtableAdd(m, v652, v648)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L33
	} else {
		goto L225
	}
L224:
	;
	v664 = v645 + v485
	v666 = v483 + int64(-1)
	if v666 != int64(0) {
		v483 = v666
		v485 = v664
		v487 = v646
		goto L173
	} else {
		goto L229
	}
L225:
	;
	if v653 != 0 {
		goto L224
	} else {
		goto L226
	}
L226:
	;
	v655 = int32(0)
	F_rdbReportError(m, int32(1), int32(2106), int32(_a_F_rdbLoadObject_7), v655)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L33
	} else {
		goto L227
	}
L227:
	;
	F_decrRefCount(m, v444)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L33
	} else {
		goto L228
	}
L228:
	;
	v4194 = v655
	goto L3
L229:
	;
	goto L174
L230:
	;
	v694 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[7]))
	if base.Ui32(v694) < base.Ui32(v691) {
		v4010 = v444
		goto L15
	} else {
		goto L231
	}
L231:
	;
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[8]))
	if base.Ui32(v697) < base.Ui32(v683) {
		v4010 = v444
		goto L15
	} else {
		goto L232
	}
L232:
	;
	goto L236
L233:
	;
	if base.B2i32(base.Ui32(int32(0)+v681) < base.Ui32(int32(1073741825))) == int32(0) {
		v4010 = v444
		goto L15
	} else {
		goto L237
	}
L234:
	;
	goto L233
L236:
	;
	goto L234
L237:
	;
	F_zsetConvert(m, v444, int32(11))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L33
	} else {
		goto L238
	}
L238:
	;
	v4010 = v444
	goto L15
L239:
	;
	switch l0 + int32(-6) {
	case 0:
		goto L14
	case 1:
		goto L16
	default:
		goto L17
	case 3, 4, 5, 6, 7, 10, 11, 14:
		goto L432
	case 8, 12:
		goto L433
	case 9, 13, 15:
		goto L18
	}
L240:
	;
	v713 = int32(0)
	v717 = F_rdbLoadLenByRef(m, l1, v713, v25+int32(192))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L33
	} else {
		goto L241
	}
L241:
	;
	v720 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v717 == int32(-1) {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v729 = F_createHashObject(m)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L33
	} else {
		goto L247
	}
L243:
	;
	v723 = int64(-1)
	goto L245
L244:
	;
	v723 = v720
	goto L245
L245:
	;
	v724 = int64(1)
	v725 = v723 + v724
	if base.Ui64(v724) < base.Ui64(v725) {
		goto L242
	} else {
		goto L246
	}
L246:
	;
	switch base.I32_wrap_i64(v725) {
	default:
		v4194 = v713
		goto L3
	case 1:
		goto L28
	}
L247:
	;
	if l0 == int32(22) {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v729)))
	if v744&int32(240) != int32(176) {
		v1032 = v720
		goto L255
	} else {
		goto L256
	}
L249:
	;
	v741 = F_hashtableCreate(m, int32(_a_F_rdbLoadObject_8))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L33
	} else {
		goto L254
	}
L250:
	;
	F_hashTypeConvert(m, v729, int32(2))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L33
	} else {
		goto L253
	}
L251:
	;
	v734 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_rdbLoadObject[9])))
	if base.Ui64(v723) <= base.Ui64(v734) {
		goto L249
	} else {
		goto L252
	}
L252:
	;
	goto L250
L253:
	;
	v743 = int32(0)
	goto L248
L254:
	;
	v743 = v741
	goto L248
L255:
	;
	if v743 == int32(0) {
		goto L352
	} else {
		goto L353
	}
L256:
	;
	if v720 == int64(0) {
		v1032 = v720
		goto L255
	} else {
		goto L257
	}
L257:
	;
	v761 = v720
	goto L258
L258:
	;
	v775 = F_rdbGenericLoadStringObject(m, l1, int32(4), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L33
	} else {
		goto L261
	}
L259:
	;
	v1032 = v822
	goto L255
L260:
	;
	v786 = F_rdbGenericLoadStringObject(m, l1, int32(4), int32(0))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L33
	} else {
		goto L267
	}
L261:
	;
	if v775 != 0 {
		goto L260
	} else {
		goto L262
	}
L262:
	;
	F_decrRefCount(m, v729)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L33
	} else {
		goto L263
	}
L263:
	;
	v779 = int32(0)
	if v743 == v779 {
		v4194 = v779
		goto L3
	} else {
		goto L264
	}
L264:
	;
	F_hashtableRelease(m, v743)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L33
	} else {
		goto L265
	}
L265:
	;
	v4194 = v779
	goto L3
L266:
	;
	if v743 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L267:
	;
	if v786 != 0 {
		goto L266
	} else {
		goto L268
	}
L268:
	;
	F_sdsfree(m, v775)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L33
	} else {
		goto L269
	}
L269:
	;
	F_decrRefCount(m, v729)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L33
	} else {
		goto L270
	}
L270:
	;
	v792 = int32(0)
	if v743 == v792 {
		v4194 = v792
		goto L3
	} else {
		goto L271
	}
L271:
	;
	F_hashtableRelease(m, v743)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L33
	} else {
		goto L272
	}
L272:
	;
	v4194 = v792
	goto L3
L273:
	;
	v822 = v761 + int64(-1)
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v729)))
	if v823&int32(240) == int32(32) {
		goto L284
	} else {
		goto L285
	}
L274:
	;
	v799 = F_sdsdup(m, v775)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L33
	} else {
		goto L275
	}
L275:
	;
	v801 = F_hashtableAdd(m, v743, v799)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L33
	} else {
		goto L276
	}
L276:
	;
	if v801 != 0 {
		goto L273
	} else {
		goto L277
	}
L277:
	;
	v803 = int32(0)
	F_rdbReportError(m, int32(1), int32(2160), int32(_a_F_rdbLoadObject_9), v803)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L33
	} else {
		goto L278
	}
L278:
	;
	F_hashtableRelease(m, v743)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L33
	} else {
		goto L279
	}
L279:
	;
	F_decrRefCount(m, v729)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L33
	} else {
		goto L280
	}
L280:
	;
	F_sdsfree(m, v799)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L33
	} else {
		goto L281
	}
L281:
	;
	F_sdsfree(m, v775)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L33
	} else {
		goto L282
	}
L282:
	;
	F_sdsfree(m, v786)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L33
	} else {
		goto L283
	}
L283:
	;
	v4194 = v803
	goto L3
L284:
	;
	v959 = F_objectGetVal(m, v729)
	mBase = m.M
	v960 = int32(0)
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775+int32(-1)))))
	switch v964 & int32(7) {
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
		v981 = v960
		goto L332
	}
L285:
	;
	v829 = v775 + int32(-1)
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829))))
	switch v830 & int32(7) {
	case 0:
		goto L293
	case 1:
		goto L292
	case 2:
		goto L291
	case 3:
		goto L290
	case 4:
		goto L289
	default:
		goto L294
	}
L286:
	;
	F_hashTypeConvert(m, v729, int32(2))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L33
	} else {
		goto L321
	}
L287:
	;
	v856 = v786 + int32(-1)
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856))))
	switch v857 & int32(7) {
	case 0:
		goto L302
	case 1:
		goto L301
	case 2:
		goto L300
	case 3:
		goto L299
	case 4:
		goto L298
	default:
		goto L296
	}
L288:
	;
	v851 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[10]))
	if base.Ui32(v851) < base.Ui32(v849) {
		goto L286
	} else {
		goto L295
	}
L289:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v775+int32(-17))))
	v849 = v848
	goto L288
L290:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v775+int32(-9))))
	v849 = v845
	goto L288
L291:
	;
	v842 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v775+int32(-5)))))
	v849 = v842
	goto L288
L292:
	;
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775+int32(-3)))))
	v849 = v839
	goto L288
L293:
	;
	v849 = int32(base.Ui32(v830) >> (uint(int32(3)) % 32))
	goto L288
L294:
	;
	v834 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[10]))
	v853 = v834
	goto L287
L295:
	;
	v853 = v851
	goto L287
L296:
	;
	v877 = F_objectGetVal(m, v729)
	mBase = m.M
	v878 = int32(0)
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829))))
	switch v880 & int32(7) {
	case 0:
		goto L309
	case 1:
		goto L308
	case 2:
		goto L307
	case 3:
		goto L306
	case 4:
		goto L305
	default:
		v897 = v878
		goto L304
	}
L297:
	;
	if base.Ui32(v853) < base.Ui32(v874) {
		goto L286
	} else {
		goto L303
	}
L298:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v786+int32(-17))))
	v874 = v873
	goto L297
L299:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v786+int32(-9))))
	v874 = v870
	goto L297
L300:
	;
	v867 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v786+int32(-5)))))
	v874 = v867
	goto L297
L301:
	;
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786+int32(-3)))))
	v874 = v864
	goto L297
L302:
	;
	v874 = int32(base.Ui32(v857) >> (uint(int32(3)) % 32))
	goto L297
L303:
	;
	goto L296
L304:
	;
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856))))
	switch v898 & int32(7) {
	case 0:
		goto L315
	case 1:
		goto L314
	case 2:
		goto L313
	case 3:
		goto L312
	case 4:
		goto L311
	default:
		v915 = v878
		goto L310
	}
L305:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v775+int32(-17))))
	v897 = v896
	goto L304
L306:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v775+int32(-9))))
	v897 = v893
	goto L304
L307:
	;
	v890 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v775+int32(-5)))))
	v897 = v890
	goto L304
L308:
	;
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775+int32(-3)))))
	v897 = v887
	goto L304
L309:
	;
	v897 = int32(base.Ui32(v880) >> (uint(int32(3)) % 32))
	goto L304
L310:
	;
	if v877 != 0 {
		goto L318
	} else {
		goto L319
	}
L311:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v786+int32(-17))))
	v915 = v914
	goto L310
L312:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v786+int32(-9))))
	v915 = v911
	goto L310
L313:
	;
	v908 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v786+int32(-5)))))
	v915 = v908
	goto L310
L314:
	;
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786+int32(-3)))))
	v915 = v905
	goto L310
L315:
	;
	v915 = int32(base.Ui32(v898) >> (uint(int32(3)) % 32))
	goto L310
L316:
	;
	if base.Ui32(v919+(v915+v897)) < base.Ui32(int32(1073741825)) {
		goto L284
	} else {
		goto L320
	}
L317:
	;
	goto L316
L318:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v877)))
	v919 = v918
	goto L317
L319:
	;
	v919 = int32(0)
	goto L317
L320:
	;
	goto L286
L321:
	;
	v932 = F_entryCreate(m, v775, v786, int64(-1))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L33
	} else {
		goto L322
	}
L322:
	;
	F_sdsfree(m, v775)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L33
	} else {
		goto L323
	}
L323:
	;
	v936 = F_objectGetVal(m, v729)
	mBase = m.M
	v937 = F_hashtableAdd(m, v936, v932)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L33
	} else {
		goto L324
	}
L324:
	;
	if v937 != 0 {
		v1032 = v822
		goto L255
	} else {
		goto L325
	}
L325:
	;
	v939 = int32(0)
	F_rdbReportError(m, int32(1), int32(2178), int32(_a_F_rdbLoadObject_10), v939)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L33
	} else {
		goto L326
	}
L326:
	;
	if v743 == int32(0) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	F_entryFree(m, v932)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L33
	} else {
		goto L330
	}
L328:
	;
	F_hashtableRelease(m, v743)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L33
	} else {
		goto L329
	}
L329:
	;
	goto L327
L330:
	;
	F_decrRefCount(m, v729)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L33
	} else {
		goto L331
	}
L331:
	;
	v4194 = v939
	goto L3
L332:
	;
	v982 = F_lpAppend(m, v959, v775, v981)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L33
	} else {
		goto L338
	}
L333:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v775+int32(-17))))
	v981 = v980
	goto L332
L334:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v775+int32(-9))))
	v981 = v977
	goto L332
L335:
	;
	v974 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v775+int32(-5)))))
	v981 = v974
	goto L332
L336:
	;
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775+int32(-3)))))
	v981 = v971
	goto L332
L337:
	;
	v981 = int32(base.Ui32(v964) >> (uint(int32(3)) % 32))
	goto L332
L338:
	;
	F_objectSetVal(m, v729, v982)
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L33
	} else {
		goto L339
	}
L339:
	;
	v986 = F_objectGetVal(m, v729)
	mBase = m.M
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786+int32(-1)))))
	switch v989 & int32(7) {
	case 0:
		goto L345
	case 1:
		goto L344
	case 2:
		goto L343
	case 3:
		goto L342
	case 4:
		goto L341
	default:
		v1006 = v960
		goto L340
	}
L340:
	;
	v1007 = F_lpAppend(m, v986, v786, v1006)
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L33
	} else {
		goto L346
	}
L341:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v786+int32(-17))))
	v1006 = v1005
	goto L340
L342:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v786+int32(-9))))
	v1006 = v1002
	goto L340
L343:
	;
	v999 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v786+int32(-5)))))
	v1006 = v999
	goto L340
L344:
	;
	v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786+int32(-3)))))
	v1006 = v996
	goto L340
L345:
	;
	v1006 = int32(base.Ui32(v989) >> (uint(int32(3)) % 32))
	goto L340
L346:
	;
	F_objectSetVal(m, v729, v1007)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L33
	} else {
		goto L347
	}
L347:
	;
	F_sdsfree(m, v775)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L33
	} else {
		goto L348
	}
L348:
	;
	F_sdsfree(m, v786)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L33
	} else {
		goto L349
	}
L349:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v729)))
	if v1015&int32(240) != int32(176) {
		v1032 = v822
		goto L255
	} else {
		goto L350
	}
L350:
	;
	if v822 != int64(0) {
		v761 = v822
		goto L258
	} else {
		goto L351
	}
L351:
	;
	goto L259
L352:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v729)))
	if v1048&int32(240) != int32(32) {
		v1060 = v1048
		goto L356
	} else {
		goto L357
	}
L353:
	;
	F_hashtableRelease(m, v743)
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L33
	} else {
		goto L354
	}
L354:
	;
	goto L352
L355:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+112)) = v1032
	F_rdbReportError(m, int32(1), int32(2205), int32(_a_F_rdbLoadObject_4), v25+int32(112))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L33
	} else {
		goto L430
	}
L356:
	;
	if v1060&int32(240) != int32(32) {
		v1356 = v1032
		goto L360
	} else {
		goto L361
	}
L357:
	;
	v1053 = F_objectGetVal(m, v729)
	mBase = m.M
	v1055 = F_hashtableTryExpand(m, v1053, base.I32_wrap_i64(v1032))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L33
	} else {
		goto L358
	}
L358:
	;
	if v1055 == int32(0) {
		goto L355
	} else {
		goto L359
	}
L359:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v729)))
	v1060 = v1059
	goto L356
L360:
	;
	if base.B2i32(v1356 == int64(0)) == int32(0) {
		goto L423
	} else {
		goto L424
	}
L361:
	;
	if v1032 == int64(0) {
		v1356 = v1032
		goto L360
	} else {
		goto L362
	}
L362:
	;
	v1077 = base.B2i32(l0 != int32(22))
	v1088 = v1032
	goto L363
L363:
	;
	v1102 = F_rdbGenericLoadStringObject(m, l1, int32(4), int32(0))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L33
	} else {
		goto L365
	}
L364:
	;
	v1356 = v1338
	goto L360
L365:
	;
	if v1102 == int32(0) {
		v4056 = v729
		goto L10
	} else {
		goto L366
	}
L366:
	;
	v1108 = F_rdbGenericLoadStringObject(m, l1, int32(4), int32(0))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L33
	} else {
		goto L368
	}
L367:
	;
	if l0 != int32(22) {
		v1256 = int64(-1)
		goto L371
	} else {
		goto L372
	}
L368:
	;
	if v1108 != 0 {
		goto L367
	} else {
		goto L369
	}
L369:
	;
	F_sdsfree(m, v1102)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L33
	} else {
		goto L370
	}
L370:
	;
	v4056 = v729
	goto L10
L371:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[11]))
	if v1266 != 0 {
		goto L399
	} else {
		goto L400
	}
L372:
	;
	v1113 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	if v1113&int64(5) != int64(0) {
		v1185 = v1113
		goto L375
	} else {
		goto L376
	}
L373:
	;
	F_sdsfree(m, v1102)
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L33
	} else {
		goto L394
	}
L374:
	;
	if v1208&int64(1) == int64(0) {
		v1256 = v1204
		goto L371
	} else {
		goto L393
	}
L375:
	;
	v1204 = int64(9223372036854775807)
	v1208 = v1185
	goto L374
L376:
	;
	v1121 = int32(8)
	v1132 = v25 + int32(192)
	goto L377
L377:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if base.Ui32(v1143) < base.Ui32(v1121) {
		goto L380
	} else {
		goto L381
	}
L378:
	;
	v1164 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v1164 < int64(-1) {
		goto L373
	} else {
		goto L392
	}
L379:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1154 == int32(0) {
		goto L388
	} else {
		goto L389
	}
L380:
	;
	v1145 = v1143
	goto L382
L381:
	;
	v1145 = v1121
	goto L382
L382:
	;
	if v1143 != 0 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1146 = v1145
	goto L385
L384:
	;
	v1146 = v1121
	goto L385
L385:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1148 = m.T0[v1147].(func(*base.Module, int32, int32, int32) int32)(m, l1, v1132, v1146)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L33
	} else {
		goto L386
	}
L386:
	;
	if v1148 != 0 {
		goto L379
	} else {
		goto L387
	}
L387:
	;
	v1150 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	v1152 = v1150 | int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = v1152
	v1185 = v1152
	goto L375
L388:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v1159 + v1146
	v1163 = v1121 - v1146
	if v1163 != 0 {
		v1121 = v1163
		v1132 = v1132 + v1146
		goto L377
	} else {
		goto L391
	}
L389:
	;
	m.T0[v1154].(func(*base.Module, int32, int32, int32))(m, l1, v1132, v1146)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L33
	} else {
		goto L390
	}
L390:
	;
	goto L388
L391:
	;
	goto L378
L392:
	;
	v1167 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	v1204 = v1164
	v1208 = v1167
	goto L374
L393:
	;
	goto L373
L394:
	;
	F_sdsfree(m, v1108)
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L33
	} else {
		goto L395
	}
L395:
	;
	v4056 = v729
	goto L10
L396:
	;
	v1338 = v1088 + int64(-1)
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v729)))
	if v1339&int32(240) != int32(32) {
		v1356 = v1338
		goto L360
	} else {
		goto L421
	}
L397:
	;
	v1314 = F_entryCreate(m, v1102, v1108, v1256)
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L33
	} else {
		goto L410
	}
L398:
	;
	if v1275 == int32(0) {
		goto L397
	} else {
		goto L401
	}
L399:
	;
	v1271 = F_getMyClusterNode(m)
	mBase = m.M
	v1272 = F_clusterNodeIsPrimary(m, v1271)
	mBase = m.M
	v1275 = base.B2i32(v1272 != int32(0))
	goto L398
L400:
	;
	v1267 = int32(0)
	v1268 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[12]))
	v1275 = base.B2i32(v1268 == v1267)
	goto L398
L401:
	;
	if l6 <= v1256 {
		goto L397
	} else {
		goto L402
	}
L402:
	;
	if base.B2i32(l5&int32(1) == int32(0))&base.B2i32(l6 != int64(0))&base.B2i32(v1256 != int64(-1)) == int32(0) {
		goto L397
	} else {
		goto L403
	}
L403:
	;
	if l5&int32(8) == int32(0) {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	F_sdsfree(m, v1102)
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L33
	} else {
		goto L408
	}
L405:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[13]))
	if v1287 == int32(0) {
		goto L404
	} else {
		goto L406
	}
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+200)) = l2
	v1291 = int64(-68719476736)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+192)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v25)+504)) = v1102
	*(*int64)(unsafe.Add(mBase, uint32(v25)+496)) = v1291
	v1297 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+528)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v25)+536)) = v25 + int32(496)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+532)) = v25 + int32(192)
	F_replicationFeedReplicas(m, l3, v25+int32(528), int32(3))
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L33
	} else {
		goto L407
	}
L407:
	;
	goto L404
L408:
	;
	F_sdsfree(m, v1108)
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L33
	} else {
		goto L409
	}
L409:
	;
	goto L396
L410:
	;
	F_sdsfree(m, v1102)
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L33
	} else {
		goto L411
	}
L411:
	;
	v1318 = F_objectGetVal(m, v729)
	mBase = m.M
	v1319 = F_hashtableAdd(m, v1318, v1314)
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L33
	} else {
		goto L413
	}
L412:
	;
	if l0 != int32(22) {
		goto L396
	} else {
		goto L418
	}
L413:
	;
	if v1319 != 0 {
		goto L412
	} else {
		goto L414
	}
L414:
	;
	v1321 = int32(0)
	F_rdbReportError(m, int32(1), int32(2261), int32(_a_F_rdbLoadObject_10), v1321)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L33
	} else {
		goto L415
	}
L415:
	;
	F_entryFree(m, v1314)
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L33
	} else {
		goto L416
	}
L416:
	;
	F_decrRefCount(m, v729)
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L33
	} else {
		goto L417
	}
L417:
	;
	v4194 = v1321
	goto L3
L418:
	;
	if v1256 == int64(-1) {
		goto L396
	} else {
		goto L419
	}
L419:
	;
	F_hashTypeTrackEntry(m, v729, v1314)
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L33
	} else {
		goto L420
	}
L420:
	;
	goto L396
L421:
	;
	if v1338 != int64(0) {
		v1088 = v1338
		goto L363
	} else {
		goto L422
	}
L422:
	;
	goto L364
L423:
	;
	F__serverAssert(m, int32(_a_F_rdbLoadObject_11), int32(_a_F_rdbLoadObject_12), int32(2273))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L33
	} else {
		goto L429
	}
L424:
	;
	v1372 = F_hashTypeLength(m, v729)
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L33
	} else {
		goto L425
	}
L425:
	;
	if v1372 != 0 {
		v4010 = v729
		goto L15
	} else {
		goto L426
	}
L426:
	;
	F_decrRefCount(m, v729)
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L33
	} else {
		goto L427
	}
L427:
	;
	if l4 == int32(0) {
		v4194 = v713
		goto L3
	} else {
		goto L428
	}
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(4)
	v4194 = v713
	goto L3
L429:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L430:
	;
	F_decrRefCount(m, v729)
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L33
	} else {
		goto L431
	}
L431:
	;
	v4194 = v713
	goto L3
L432:
	;
	v1593 = F_rdbGenericLoadStringObject(m, l1, int32(2), v25+int32(192))
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L33
	} else {
		goto L491
	}
L433:
	;
	v1398 = int32(0)
	v1402 = F_rdbLoadLenByRef(m, l1, v1398, v25+int32(192))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L33
	} else {
		goto L434
	}
L434:
	;
	v1405 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v1402 == int32(-1) {
		goto L436
	} else {
		goto L437
	}
L435:
	;
	v1414 = int32(_a_F_rdbLoadObject_3)
	v1415 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[2]))
	v1417 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[3]))
	v1418 = F_createQuicklistObject(m, v1415, v1417)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L33
	} else {
		goto L440
	}
L436:
	;
	v1408 = int64(-1)
	goto L438
L437:
	;
	v1408 = v1405
	goto L438
L438:
	;
	v1409 = int64(1)
	v1410 = v1408 + v1409
	if base.Ui64(v1409) < base.Ui64(v1410) {
		goto L435
	} else {
		goto L439
	}
L439:
	;
	switch base.I32_wrap_i64(v1410) {
	default:
		v4194 = v1398
		goto L3
	case 1:
		goto L28
	}
L440:
	;
	if v1405 == int64(0) {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	v1581 = F_objectGetVal(m, v1418)
	mBase = m.M
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1581)+8))
	goto L488
L442:
	;
	v1424 = base.B2i32(l0 != int32(18))
	v1431 = v1405
	v1435 = int64(2)
	goto L443
L443:
	;
	if l0 != int32(18) {
		v1470 = v1435
		goto L445
	} else {
		goto L446
	}
L444:
	;
	goto L441
L445:
	;
	v1474 = F_rdbGenericLoadStringObject(m, l1, int32(2), v25+int32(528))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L33
	} else {
		goto L455
	}
L446:
	;
	v1450 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L33
	} else {
		goto L447
	}
L447:
	;
	if v1450 == int32(-1) {
		v4056 = v1418
		goto L10
	} else {
		goto L448
	}
L448:
	;
	v1454 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v1454 == int64(-1) {
		v4056 = v1418
		goto L10
	} else {
		goto L449
	}
L449:
	;
	v1457 = int64(-3)
	if base.Ui64(v1457) < base.Ui64(v1454+v1457) {
		v1470 = v1454
		goto L445
	} else {
		goto L450
	}
L450:
	;
	v1461 = int32(0)
	F_rdbReportError(m, int32(1), int32(2298), int32(_a_F_rdbLoadObject_13), v1461)
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L33
	} else {
		goto L451
	}
L451:
	;
	F_decrRefCount(m, v1418)
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L33
	} else {
		goto L452
	}
L452:
	;
	v4194 = v1461
	goto L3
L453:
	;
	if v1470 != int64(1) {
		goto L460
	} else {
		goto L461
	}
L454:
	;
	F_valkey_free(m, v1474)
	mBase = m.M
	v1481 = m.ExcPending
	if v1481 != 0 {
		goto L33
	} else {
		goto L458
	}
L455:
	;
	if v1474 == int32(0) {
		goto L454
	} else {
		goto L456
	}
L456:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v25)+528))
	if v1478 != 0 {
		goto L453
	} else {
		goto L457
	}
L457:
	;
	goto L454
L458:
	;
	v4056 = v1418
	goto L10
L459:
	;
	v1556 = v1431 + int64(-1)
	if v1556 != int64(0) {
		v1431 = v1556
		v1435 = v1470
		goto L443
	} else {
		goto L487
	}
L460:
	;
	if l0 != int32(18) {
		goto L464
	} else {
		goto L465
	}
L461:
	;
	v1484 = F_objectGetVal(m, v1418)
	mBase = m.M
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v25)+528))
	F_quicklistAppendPlainNode(m, v1484, v1474, v1485)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L33
	} else {
		goto L462
	}
L462:
	;
	goto L459
L463:
	;
	v1545 = F_lpLength(m, v1544)
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L33
	} else {
		goto L483
	}
L464:
	;
	v1514 = F_lpNew(m, v1478)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L33
	} else {
		goto L472
	}
L465:
	;
	v1488 = int32(_a_F_rdbLoadObject_3)
	v1490 = *(*int64)(unsafe.Add(mBase, _c_F_rdbLoadObject[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_rdbLoadObject[1])) = v1490 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+496)) = v1474
	v1495 = int32(0)
	v1497 = F_lpValidateIntegrity(m, v1474, v1478, v1495, v1495)
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L33
	} else {
		goto L467
	}
L466:
	;
	v1502 = int32(0)
	F_rdbReportError(m, int32(1), int32(2320), int32(_a_F_rdbLoadObject_14), v1502)
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L33
	} else {
		goto L469
	}
L467:
	;
	if v1497 == int32(0) {
		goto L466
	} else {
		goto L468
	}
L468:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v25)+496))
	v1544 = v1501
	goto L463
L469:
	;
	F_decrRefCount(m, v1418)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L33
	} else {
		goto L470
	}
L470:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v25)+496))
	F_valkey_free(m, v1511)
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L33
	} else {
		goto L471
	}
L471:
	;
	v4194 = v1502
	goto L3
L472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+496)) = v1514
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v25)+528))
	v1522 = F_ziplistValidateIntegrity(m, v1474, v1517, int32(1), int32(965), v25+int32(496))
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L33
	} else {
		goto L474
	}
L473:
	;
	F_valkey_free(m, v1474)
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L33
	} else {
		goto L480
	}
L474:
	;
	if v1522 != 0 {
		goto L473
	} else {
		goto L475
	}
L475:
	;
	v1524 = int32(0)
	F_rdbReportError(m, int32(1), int32(2328), int32(_a_F_rdbLoadObject_15), v1524)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L33
	} else {
		goto L476
	}
L476:
	;
	F_decrRefCount(m, v1418)
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L33
	} else {
		goto L477
	}
L477:
	;
	F_valkey_free(m, v1474)
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L33
	} else {
		goto L478
	}
L478:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v25)+496))
	F_valkey_free(m, v1535)
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L33
	} else {
		goto L479
	}
L479:
	;
	v4194 = v1524
	goto L3
L480:
	;
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v25)+496))
	v1541 = F_lpShrinkToFit(m, v1540)
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L33
	} else {
		goto L481
	}
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+496)) = v1541
	v1544 = v1541
	goto L463
L482:
	;
	v1550 = F_objectGetVal(m, v1418)
	mBase = m.M
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v25)+496))
	F_quicklistAppendListpack(m, v1550, v1551)
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L33
	} else {
		goto L486
	}
L483:
	;
	if v1545 != 0 {
		goto L482
	} else {
		goto L484
	}
L484:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v25)+496))
	F_valkey_free(m, v1547)
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L33
	} else {
		goto L485
	}
L485:
	;
	goto L459
L486:
	;
	goto L459
L487:
	;
	goto L444
L488:
	;
	if v1582 == int32(0) {
		v4088 = v1418
		goto L8
	} else {
		goto L489
	}
L489:
	;
	v1585 = int32(0)
	F_listTypeTryConversion(m, v1418, v1585, v1585, v1585)
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		goto L33
	} else {
		goto L490
	}
L490:
	;
	v4010 = v1418
	goto L15
L491:
	;
	if v1593 == int32(0) {
		goto L6
	} else {
		goto L492
	}
L492:
	;
	v1598 = F_createObject(m, int32(0), v1593)
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L33
	} else {
		goto L493
	}
L493:
	;
	switch l0 + int32(-9) {
	case 0:
		goto L495
	case 1:
		goto L494
	case 2:
		goto L26
	case 3:
		goto L24
	case 4:
		goto L22
	default:
		goto L20
	case 7:
		goto L21
	case 8:
		goto L23
	case 11:
		goto L25
	}
L494:
	;
	v2010 = int32(_a_F_rdbLoadObject_3)
	v2011 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[2]))
	v2013 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[3]))
	v2014 = F_quicklistNew(m, v2011, v2013)
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L33
	} else {
		goto L619
	}
L495:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v25)+192))
	v1604 = int32(0)
	if base.Ui32(v1602) < base.Ui32(int32(2)) {
		v1701 = v1604
		goto L498
	} else {
		goto L499
	}
L496:
	;
	v1728 = int32(0)
	v1730 = F_lpNew(m, v1728)
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L33
	} else {
		goto L532
	}
L497:
	;
	if v1716 != 0 {
		goto L496
	} else {
		goto L528
	}
L498:
	;
	v1716 = v1701
	goto L497
L499:
	;
	v1615 = v1593 + v1602 + int32(-1)
	v1616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1615))))
	if v1616 != int32(255) {
		v1701 = v1604
		goto L498
	} else {
		goto L500
	}
L500:
	;
	goto L501
L501:
	;
	v1626 = int32(0)
	v1627 = int32(1)
	goto L504
L503:
	;
	if v1626 != 0 {
		goto L526
	} else {
		goto L527
	}
L504:
	;
	v1631 = v1593 + v1627
	v1632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1631))))
	if v1632 == int32(255) {
		goto L503
	} else {
		goto L506
	}
L506:
	;
	v1635 = int32(0)
	v1639 = base.B2i32(v1632 == int32(254))
	if v1632 == int32(254) {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v1640 = int32(5)
	goto L509
L508:
	;
	v1640 = int32(1)
	goto L509
L509:
	;
	v1641 = v1640 + v1627
	if v1641 < int32(2) {
		v1701 = v1635
		goto L498
	} else {
		goto L510
	}
L510:
	;
	v1644 = v1593 + v1641
	if base.Ui32(v1615) < base.Ui32(v1644) {
		v1701 = v1635
		goto L498
	} else {
		goto L511
	}
L511:
	;
	if v1639 == int32(0) {
		v1651 = v1632
		goto L512
	} else {
		goto L513
	}
L512:
	;
	if base.Ui64(base.I64_extend_i32_s(v1615-v1644)) < base.Ui64(base.I64_extend_i32_u(v1651)) {
		v1701 = v1635
		goto L498
	} else {
		goto L515
	}
L513:
	;
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+1))
	if base.Ui32(v1648) < base.Ui32(int32(254)) {
		v1701 = v1635
		goto L498
	} else {
		goto L514
	}
L514:
	;
	v1651 = v1648
	goto L512
L515:
	;
	v1656 = v1651 + v1641
	v1657 = v1593 + v1656
	if base.Ui32(v1615) < base.Ui32(v1657) {
		v1701 = v1635
		goto L498
	} else {
		goto L516
	}
L516:
	;
	v1661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1657))))
	v1663 = base.B2i32(base.Ui32(int32(253)) < base.Ui32(v1661))
	if base.Ui32(int32(253)) < base.Ui32(v1661) {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v1664 = int32(5)
	goto L519
L518:
	;
	v1664 = int32(1)
	goto L519
L519:
	;
	v1665 = v1664 + v1656
	v1666 = v1593 + v1665
	if base.Ui32(v1615) < base.Ui32(v1666) {
		v1701 = v1635
		goto L498
	} else {
		goto L520
	}
L520:
	;
	if v1663 == int32(0) {
		v1673 = v1661
		goto L521
	} else {
		goto L522
	}
L521:
	;
	v1674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1666))))
	v1679 = v1665 + int32(1)
	if base.Ui64(base.I64_extend_i32_s(v1615-(v1593+v1679))) < base.Ui64(base.I64_extend_i32_u(v1674)+base.I64_extend_i32_u(v1673)) {
		v1701 = v1635
		goto L498
	} else {
		goto L524
	}
L522:
	;
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v1657)+1))
	if base.Ui32(v1670) < base.Ui32(int32(254)) {
		v1701 = v1635
		goto L498
	} else {
		goto L523
	}
L523:
	;
	v1673 = v1670
	goto L521
L524:
	;
	v1687 = v1673 + v1679 + v1674
	if base.Ui32(v1593+v1687) <= base.Ui32(v1615) {
		v1626 = v1626 + int32(1)
		v1627 = v1687
		goto L504
	} else {
		goto L525
	}
L525:
	;
	v1701 = v1635
	goto L498
L526:
	;
	v1691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1593))))
	v1693 = v1691 & int32(255)
	v1701 = base.B2i32(v1693 == int32(254)) | base.B2i32(v1626 == v1693)
	goto L498
L527:
	;
	v1716 = int32(0)
	goto L497
L528:
	;
	F_rdbReportError(m, int32(1), int32(2374), int32(_a_F_rdbLoadObject_16), int32(0))
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L33
	} else {
		goto L529
	}
L529:
	;
	F_valkey_free(m, v1593)
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L33
	} else {
		goto L530
	}
L530:
	;
	F_objectSetVal(m, v1598, int32(0))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L33
	} else {
		goto L531
	}
L531:
	;
	v4056 = v1598
	goto L10
L532:
	;
	v1732 = F_objectGetVal(m, v1598)
	mBase = m.M
	v1734 = v1732 + int32(1)
	goto L533
L533:
	;
	v1736 = F_hashtableCreate(m, int32(_a_F_rdbLoadObject_8))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L33
	} else {
		goto L534
	}
L534:
	;
	v1739 = v25 + int32(496)
	v1743 = v25 + int32(528)
	v1747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1734))))
	if v1747 != int32(255) {
		goto L537
	} else {
		goto L538
	}
L535:
	;
	F_hashtableRelease(m, v1736)
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L33
	} else {
		goto L611
	}
L536:
	;
	if v1823 == int32(0) {
		v1968 = v1728
		v1976 = v1730
		goto L535
	} else {
		goto L563
	}
L537:
	;
	if v1739 == int32(0) {
		v1768 = v1747
		goto L539
	} else {
		goto L540
	}
L538:
	;
	v1823 = int32(0)
	goto L536
L539:
	;
	v1770 = v1768 & int32(255)
	if base.Ui32(v1770) <= base.Ui32(int32(253)) {
		v1774 = v1770
		goto L546
	} else {
		goto L547
	}
L540:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1739))) = v1734
	v1754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1734))))
	if base.Ui32(v1754) <= base.Ui32(int32(253)) {
		v1758 = v1754
		goto L541
	} else {
		goto L542
	}
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+int32(512)))) = v1758
	if base.Ui32(v1758) < base.Ui32(int32(254)) {
		goto L543
	} else {
		goto L544
	}
L542:
	;
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1734)+1))
	v1758 = v1757
	goto L541
L543:
	;
	v1764 = int32(1)
	goto L545
L544:
	;
	v1764 = int32(5)
	goto L545
L545:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1739))) = v1734 + v1764
	v1767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1734))))
	v1768 = v1767
	goto L539
L546:
	;
	if base.Ui32(v1774) < base.Ui32(int32(254)) {
		goto L548
	} else {
		goto L549
	}
L547:
	;
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v1734)+1))
	v1774 = v1773
	goto L546
L548:
	;
	v1779 = int32(1)
	goto L550
L549:
	;
	v1779 = int32(5)
	goto L550
L550:
	;
	v1781 = v1734 + v1779 + v1774
	if v1743 == int32(0) {
		goto L551
	} else {
		goto L552
	}
L551:
	;
	v1802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1781))))
	if base.Ui32(v1802) <= base.Ui32(int32(253)) {
		v1806 = v1802
		goto L558
	} else {
		goto L559
	}
L552:
	;
	v1785 = v1781 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1743))) = v1785
	v1787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1781))))
	if base.Ui32(v1787) <= base.Ui32(int32(253)) {
		v1791 = v1787
		goto L553
	} else {
		goto L554
	}
L553:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+int32(524)))) = v1791
	if base.Ui32(v1791) < base.Ui32(int32(254)) {
		goto L555
	} else {
		goto L556
	}
L554:
	;
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(v1785)))
	v1791 = v1790
	goto L553
L555:
	;
	v1797 = int32(1)
	goto L557
L556:
	;
	v1797 = int32(5)
	goto L557
L557:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1743))) = v1785 + v1797
	goto L551
L558:
	;
	if base.Ui32(v1806) < base.Ui32(int32(254)) {
		goto L560
	} else {
		goto L561
	}
L559:
	;
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(v1781)+1))
	v1806 = v1805
	goto L558
L560:
	;
	v1812 = int32(1)
	goto L562
L561:
	;
	v1812 = int32(5)
	goto L562
L562:
	;
	v1814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1781+v1812))))
	v1823 = v1781 + v1806 + v1814 + v1812 + int32(1)
	goto L536
L563:
	;
	v1826 = v1728
	v1834 = v1730
	v1837 = v1823
	goto L564
L564:
	;
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v25)+524))
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v25)+496))
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v25)+512))
	v1851 = F_sdstrynewlen(m, v1849, v1850)
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L33
	} else {
		goto L566
	}
L565:
	;
	v1968 = v1873
	v1976 = v1880
	goto L535
L566:
	;
	if v1851 == int32(0) {
		goto L19
	} else {
		goto L567
	}
L567:
	;
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v25)+524))
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v25)+512))
	if v1834 != 0 {
		goto L570
	} else {
		goto L571
	}
L568:
	;
	if base.B2i32(base.Ui32(v1860+(v1855+v1856)) < base.Ui32(int32(1073741825))) == int32(0) {
		goto L19
	} else {
		goto L572
	}
L569:
	;
	goto L568
L570:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1834)))
	v1860 = v1859
	goto L569
L571:
	;
	v1860 = int32(0)
	goto L569
L572:
	;
	v1866 = F_hashtableAdd(m, v1736, v1851)
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L33
	} else {
		goto L573
	}
L573:
	;
	if v1866 == int32(0) {
		goto L19
	} else {
		goto L574
	}
L574:
	;
	if base.Ui32(v1826) < base.Ui32(v1850) {
		goto L575
	} else {
		goto L576
	}
L575:
	;
	v1871 = v1850
	goto L577
L576:
	;
	v1871 = v1826
	goto L577
L577:
	;
	if base.Ui32(v1871) < base.Ui32(v1848) {
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v1873 = v1848
	goto L580
L579:
	;
	v1873 = v1871
	goto L580
L580:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v25)+496))
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v25)+512))
	v1876 = F_lpAppend(m, v1834, v1874, v1875)
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L33
	} else {
		goto L581
	}
L581:
	;
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v25)+528))
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v25)+524))
	v1880 = F_lpAppend(m, v1876, v1878, v1879)
	mBase = m.M
	v1881 = m.ExcPending
	if v1881 != 0 {
		goto L33
	} else {
		goto L582
	}
L582:
	;
	v1883 = v25 + int32(496)
	v1887 = v25 + int32(528)
	v1891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1837))))
	if v1891 != int32(255) {
		goto L584
	} else {
		goto L585
	}
L583:
	;
	if v1967 != 0 {
		v1826 = v1873
		v1834 = v1880
		v1837 = v1967
		goto L564
	} else {
		goto L610
	}
L584:
	;
	if v1883 == int32(0) {
		v1912 = v1891
		goto L586
	} else {
		goto L587
	}
L585:
	;
	v1967 = int32(0)
	goto L583
L586:
	;
	v1914 = v1912 & int32(255)
	if base.Ui32(v1914) <= base.Ui32(int32(253)) {
		v1918 = v1914
		goto L593
	} else {
		goto L594
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1883))) = v1837
	v1898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1837))))
	if base.Ui32(v1898) <= base.Ui32(int32(253)) {
		v1902 = v1898
		goto L588
	} else {
		goto L589
	}
L588:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+int32(512)))) = v1902
	if base.Ui32(v1902) < base.Ui32(int32(254)) {
		goto L590
	} else {
		goto L591
	}
L589:
	;
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(v1837)+1))
	v1902 = v1901
	goto L588
L590:
	;
	v1908 = int32(1)
	goto L592
L591:
	;
	v1908 = int32(5)
	goto L592
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1883))) = v1837 + v1908
	v1911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1837))))
	v1912 = v1911
	goto L586
L593:
	;
	if base.Ui32(v1918) < base.Ui32(int32(254)) {
		goto L595
	} else {
		goto L596
	}
L594:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v1837)+1))
	v1918 = v1917
	goto L593
L595:
	;
	v1923 = int32(1)
	goto L597
L596:
	;
	v1923 = int32(5)
	goto L597
L597:
	;
	v1925 = v1837 + v1923 + v1918
	if v1887 == int32(0) {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	v1946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1925))))
	if base.Ui32(v1946) <= base.Ui32(int32(253)) {
		v1950 = v1946
		goto L605
	} else {
		goto L606
	}
L599:
	;
	v1929 = v1925 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1887))) = v1929
	v1931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1925))))
	if base.Ui32(v1931) <= base.Ui32(int32(253)) {
		v1935 = v1931
		goto L600
	} else {
		goto L601
	}
L600:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+int32(524)))) = v1935
	if base.Ui32(v1935) < base.Ui32(int32(254)) {
		goto L602
	} else {
		goto L603
	}
L601:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1929)))
	v1935 = v1934
	goto L600
L602:
	;
	v1941 = int32(1)
	goto L604
L603:
	;
	v1941 = int32(5)
	goto L604
L604:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1887))) = v1929 + v1941
	goto L598
L605:
	;
	if base.Ui32(v1950) < base.Ui32(int32(254)) {
		goto L607
	} else {
		goto L608
	}
L606:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v1925)+1))
	v1950 = v1949
	goto L605
L607:
	;
	v1956 = int32(1)
	goto L609
L608:
	;
	v1956 = int32(5)
	goto L609
L609:
	;
	v1958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1925+v1956))))
	v1967 = v1925 + v1950 + v1958 + v1956 + int32(1)
	goto L583
L610:
	;
	goto L565
L611:
	;
	v1992 = F_objectGetVal(m, v1598)
	mBase = m.M
	F_valkey_free(m, v1992)
	mBase = m.M
	v1994 = m.ExcPending
	if v1994 != 0 {
		goto L33
	} else {
		goto L612
	}
L612:
	;
	F_objectSetVal(m, v1598, v1976)
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L33
	} else {
		goto L613
	}
L613:
	;
	v1997 = int32(180)
	*(*uint8)(unsafe.Add(mBase, uint32(v1598))) = uint8(v1997)
	v1999 = F_hashTypeLength(m, v1598)
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L33
	} else {
		goto L615
	}
L614:
	;
	F_hashTypeConvert(m, v1598, int32(2))
	mBase = m.M
	v2009 = m.ExcPending
	if v2009 != 0 {
		goto L33
	} else {
		goto L618
	}
L615:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[9]))
	if base.Ui32(v2002) < base.Ui32(v1999) {
		goto L614
	} else {
		goto L616
	}
L616:
	;
	v2005 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[10]))
	if base.Ui32(v1968) <= base.Ui32(v2005) {
		v4010 = v1598
		goto L15
	} else {
		goto L617
	}
L617:
	;
	goto L614
L618:
	;
	v4010 = v1598
	goto L15
L619:
	;
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v25)+192))
	v2019 = F_ziplistValidateIntegrity(m, v1593, v2016, int32(1), int32(966), v2014)
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L33
	} else {
		goto L621
	}
L620:
	;
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v2014)+12))
	F_valkey_free(m, v1593)
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L33
	} else {
		goto L628
	}
L621:
	;
	if v2019 != 0 {
		goto L620
	} else {
		goto L622
	}
L622:
	;
	v2021 = int32(0)
	F_rdbReportError(m, int32(1), int32(2427), int32(_a_F_rdbLoadObject_17), v2021)
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L33
	} else {
		goto L623
	}
L623:
	;
	F_valkey_free(m, v1593)
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L33
	} else {
		goto L624
	}
L624:
	;
	F_objectSetVal(m, v1598, int32(0))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L33
	} else {
		goto L625
	}
L625:
	;
	F_decrRefCount(m, v1598)
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L33
	} else {
		goto L626
	}
L626:
	;
	F_quicklistRelease(m, v2014)
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L33
	} else {
		goto L627
	}
L627:
	;
	v4194 = v2021
	goto L3
L628:
	;
	if v2037 != 0 {
		goto L27
	} else {
		goto L629
	}
L629:
	;
	F_objectSetVal(m, v1598, int32(0))
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L33
	} else {
		goto L630
	}
L630:
	;
	F_decrRefCount(m, v1598)
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L33
	} else {
		goto L631
	}
L631:
	;
	F_quicklistRelease(m, v2014)
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L33
	} else {
		goto L632
	}
L632:
	;
	goto L28
L633:
	;
	goto L6
L634:
	;
	v2063 = *(*int32)(unsafe.Add(mBase, uint32(v1598)))
	*(*int32)(unsafe.Add(mBase, uint32(v1598))) = v2063&int32(-241) | int32(144)
	v4010 = v1598
	goto L15
L635:
	;
	v2177 = int32(98)
	*(*uint8)(unsafe.Add(mBase, uint32(v1598))) = uint8(v2177)
	v2179 = F_objectGetVal(m, v1598)
	mBase = m.M
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(v2179)+4))
	goto L664
L636:
	;
	if v2165 != 0 {
		goto L635
	} else {
		goto L660
	}
L637:
	;
	v2165 = v2151
	goto L636
L638:
	;
	v2151 = int32(0)
	goto L637
L639:
	;
	v2083 = *(*int32)(unsafe.Add(mBase, uint32(v1593)))
	if base.Ui32(int32(8)) < base.Ui32(v2083) {
		goto L638
	} else {
		goto L640
	}
L640:
	;
	if int32(1)<<(uint(v2083)%32)&int32(276) == int32(0) {
		goto L638
	} else {
		goto L641
	}
L641:
	;
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+4))
	if v2092 == int32(0) {
		goto L638
	} else {
		goto L642
	}
L642:
	;
	if v2092*v2083+int32(8) != v2075 {
		goto L638
	} else {
		goto L643
	}
L643:
	;
	goto L644
L644:
	;
	v2101 = v1593 + int32(8)
	switch v2083&int32(255) + int32(-4) {
	case 0:
		goto L648
	default:
		goto L647
	case 4:
		goto L649
	}
L646:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v2092) {
		goto L650
	} else {
		goto L651
	}
L647:
	;
	v2108 = int64(*(*int16)(unsafe.Add(mBase, uint32(v2101))))
	v2109 = v2108
	goto L646
L648:
	;
	v2107 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2101))))
	v2109 = v2107
	goto L646
L649:
	;
	v2106 = *(*int64)(unsafe.Add(mBase, uint32(v2101)))
	v2109 = v2106
	goto L646
L650:
	;
	v2119 = int32(1)
	v2123 = v2109
	goto L652
L651:
	;
	v2165 = int32(1)
	goto L636
L652:
	;
	switch v2083&int32(255) + int32(-4) {
	case 0:
		goto L656
	default:
		goto L655
	case 4:
		goto L657
	}
L654:
	;
	if v2137 <= v2123 {
		goto L638
	} else {
		goto L658
	}
L655:
	;
	v2136 = int64(*(*int16)(unsafe.Add(mBase, uint32(v2101+v2119<<(uint(int32(1))%32)))))
	v2137 = v2136
	goto L654
L656:
	;
	v2132 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2101+v2119<<(uint(int32(2))%32)))))
	v2137 = v2132
	goto L654
L657:
	;
	v2128 = *(*int64)(unsafe.Add(mBase, uint32(v2101+v2119<<(uint(int32(3))%32))))
	v2137 = v2128
	goto L654
L658:
	;
	v2139 = int32(1)
	v2141 = v2119 + v2139
	if v2141 == v2092 {
		v2151 = v2139
		goto L637
	} else {
		goto L659
	}
L659:
	;
	v2119 = v2141
	v2123 = v2137
	goto L652
L660:
	;
	F_rdbReportError(m, int32(1), int32(2452), int32(_a_F_rdbLoadObject_18), int32(0))
	mBase = m.M
	v2171 = m.ExcPending
	if v2171 != 0 {
		goto L33
	} else {
		goto L661
	}
L661:
	;
	F_valkey_free(m, v1593)
	mBase = m.M
	v2173 = m.ExcPending
	if v2173 != 0 {
		goto L33
	} else {
		goto L662
	}
L662:
	;
	F_objectSetVal(m, v1598, int32(0))
	mBase = m.M
	v2176 = m.ExcPending
	if v2176 != 0 {
		goto L33
	} else {
		goto L663
	}
L663:
	;
	v4056 = v1598
	goto L10
L664:
	;
	v2182 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[4]))
	if base.Ui32(v2180) <= base.Ui32(v2182) {
		v4010 = v1598
		goto L15
	} else {
		goto L665
	}
L665:
	;
	F_setTypeConvert(m, v1598, int32(2))
	mBase = m.M
	v2186 = m.ExcPending
	if v2186 != 0 {
		goto L33
	} else {
		goto L666
	}
L666:
	;
	v4010 = v1598
	goto L15
L667:
	;
	v2208 = int32(178)
	*(*uint8)(unsafe.Add(mBase, uint32(v1598))) = uint8(v2208)
	v2210 = F_setTypeSize(m, v1598)
	mBase = m.M
	v2211 = m.ExcPending
	if v2211 != 0 {
		goto L33
	} else {
		goto L674
	}
L668:
	;
	if v2195 != 0 {
		goto L667
	} else {
		goto L669
	}
L669:
	;
	F_rdbReportError(m, int32(1), int32(2465), int32(_a_F_rdbLoadObject_19), int32(0))
	mBase = m.M
	v2202 = m.ExcPending
	if v2202 != 0 {
		goto L33
	} else {
		goto L670
	}
L670:
	;
	F_valkey_free(m, v1593)
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L33
	} else {
		goto L671
	}
L671:
	;
	F_objectSetVal(m, v1598, int32(0))
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		goto L33
	} else {
		goto L672
	}
L672:
	;
	v4056 = v1598
	goto L10
L673:
	;
	v2217 = F_setTypeSize(m, v1598)
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L33
	} else {
		goto L678
	}
L674:
	;
	if v2210 != 0 {
		goto L673
	} else {
		goto L675
	}
L675:
	;
	F_valkey_free(m, v1593)
	mBase = m.M
	v2213 = m.ExcPending
	if v2213 != 0 {
		goto L33
	} else {
		goto L676
	}
L676:
	;
	F_objectSetVal(m, v1598, int32(0))
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L33
	} else {
		goto L677
	}
L677:
	;
	v4088 = v1598
	goto L8
L678:
	;
	v2220 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[5]))
	if base.Ui32(v2217) <= base.Ui32(v2220) {
		v4010 = v1598
		goto L15
	} else {
		goto L679
	}
L679:
	;
	F_setTypeConvert(m, v1598, int32(2))
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		goto L33
	} else {
		goto L680
	}
L680:
	;
	v4010 = v1598
	goto L15
L681:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+496)) = v2226
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(v25)+192))
	v2232 = F_ziplistPairsConvertAndValidateIntegrity(m, v1593, v2229, v25+int32(496))
	mBase = m.M
	v2233 = m.ExcPending
	if v2233 != 0 {
		goto L33
	} else {
		goto L685
	}
L682:
	;
	v2258 = F_objectGetVal(m, v1598)
	mBase = m.M
	F_valkey_free(m, v2258)
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L33
	} else {
		goto L694
	}
L683:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v25)+496))
	F_valkey_free(m, v2249)
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L33
	} else {
		goto L691
	}
L684:
	;
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v25)+496))
	v2241 = F_zzlValidateScores(m, v2240)
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L33
	} else {
		goto L688
	}
L685:
	;
	if v2232 != 0 {
		goto L684
	} else {
		goto L686
	}
L686:
	;
	F_rdbReportError(m, int32(1), int32(2485), int32(_a_F_rdbLoadObject_20), int32(0))
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L33
	} else {
		goto L687
	}
L687:
	;
	goto L683
L688:
	;
	if v2241 != 0 {
		goto L682
	} else {
		goto L689
	}
L689:
	;
	F_rdbReportError(m, int32(1), int32(2497), int32(_a_F_rdbLoadObject_21), int32(0))
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L33
	} else {
		goto L690
	}
L690:
	;
	goto L683
L691:
	;
	F_valkey_free(m, v1593)
	mBase = m.M
	v2253 = m.ExcPending
	if v2253 != 0 {
		goto L33
	} else {
		goto L692
	}
L692:
	;
	F_objectSetVal(m, v1598, int32(0))
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L33
	} else {
		goto L693
	}
L693:
	;
	v4071 = int32(1)
	goto L9
L694:
	;
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(v1598)))
	*(*int32)(unsafe.Add(mBase, uint32(v1598))) = v2261&int32(-16) | int32(3)
	v2267 = *(*int32)(unsafe.Add(mBase, uint32(v25)+496))
	F_objectSetVal(m, v1598, v2267)
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L33
	} else {
		goto L695
	}
L695:
	;
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(v1598)))
	*(*int32)(unsafe.Add(mBase, uint32(v1598))) = v2270&int32(-241) | int32(176)
	v2276 = F_zsetLength(m, v1598)
	mBase = m.M
	v2277 = m.ExcPending
	if v2277 != 0 {
		goto L33
	} else {
		goto L697
	}
L696:
	;
	v2279 = F_zsetLength(m, v1598)
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L33
	} else {
		goto L700
	}
L697:
	;
	if v2276 != 0 {
		goto L696
	} else {
		goto L698
	}
L698:
	;
	v4071 = int32(2)
	goto L9
L699:
	;
	v2287 = F_objectGetVal(m, v1598)
	mBase = m.M
	v2288 = F_lpShrinkToFit(m, v2287)
	mBase = m.M
	v2289 = m.ExcPending
	if v2289 != 0 {
		goto L33
	} else {
		goto L703
	}
L700:
	;
	v2282 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[7]))
	if base.Ui32(v2279) <= base.Ui32(v2282) {
		goto L699
	} else {
		goto L701
	}
L701:
	;
	F_zsetConvert(m, v1598, int32(7))
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		goto L33
	} else {
		goto L702
	}
L702:
	;
	v4010 = v1598
	goto L15
L703:
	;
	F_objectSetVal(m, v1598, v2288)
	mBase = m.M
	v2291 = m.ExcPending
	if v2291 != 0 {
		goto L33
	} else {
		goto L704
	}
L704:
	;
	v4010 = v1598
	goto L15
L705:
	;
	v2313 = F_zzlValidateScores(m, v1593)
	mBase = m.M
	v2314 = m.ExcPending
	if v2314 != 0 {
		goto L33
	} else {
		goto L712
	}
L706:
	;
	if v2300 != 0 {
		goto L705
	} else {
		goto L707
	}
L707:
	;
	F_rdbReportError(m, int32(1), int32(2523), int32(_a_F_rdbLoadObject_22), int32(0))
	mBase = m.M
	v2307 = m.ExcPending
	if v2307 != 0 {
		goto L33
	} else {
		goto L708
	}
L708:
	;
	F_valkey_free(m, v1593)
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L33
	} else {
		goto L709
	}
L709:
	;
	F_objectSetVal(m, v1598, int32(0))
	mBase = m.M
	v2312 = m.ExcPending
	if v2312 != 0 {
		goto L33
	} else {
		goto L710
	}
L710:
	;
	v4056 = v1598
	goto L10
L711:
	;
	v2326 = int32(179)
	*(*uint8)(unsafe.Add(mBase, uint32(v1598))) = uint8(v2326)
	v2328 = F_zsetLength(m, v1598)
	mBase = m.M
	v2329 = m.ExcPending
	if v2329 != 0 {
		goto L33
	} else {
		goto L717
	}
L712:
	;
	if v2313 != 0 {
		goto L711
	} else {
		goto L713
	}
L713:
	;
	F_rdbReportError(m, int32(1), int32(2534), int32(_a_F_rdbLoadObject_23), int32(0))
	mBase = m.M
	v2320 = m.ExcPending
	if v2320 != 0 {
		goto L33
	} else {
		goto L714
	}
L714:
	;
	F_valkey_free(m, v1593)
	mBase = m.M
	v2322 = m.ExcPending
	if v2322 != 0 {
		goto L33
	} else {
		goto L715
	}
L715:
	;
	F_objectSetVal(m, v1598, int32(0))
	mBase = m.M
	v2325 = m.ExcPending
	if v2325 != 0 {
		goto L33
	} else {
		goto L716
	}
L716:
	;
	v4056 = v1598
	goto L10
L717:
	;
	if v2328 == int32(0) {
		v4088 = v1598
		goto L8
	} else {
		goto L718
	}
L718:
	;
	v2332 = F_zsetLength(m, v1598)
	mBase = m.M
	v2333 = m.ExcPending
	if v2333 != 0 {
		goto L33
	} else {
		goto L719
	}
L719:
	;
	v2335 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[7]))
	if base.Ui32(v2332) <= base.Ui32(v2335) {
		v4010 = v1598
		goto L15
	} else {
		goto L720
	}
L720:
	;
	F_zsetConvert(m, v1598, int32(7))
	mBase = m.M
	v2339 = m.ExcPending
	if v2339 != 0 {
		goto L33
	} else {
		goto L721
	}
L721:
	;
	v4010 = v1598
	goto L15
L722:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+496)) = v2341
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(v25)+192))
	v2347 = F_ziplistPairsConvertAndValidateIntegrity(m, v1593, v2344, v25+int32(496))
	mBase = m.M
	v2348 = m.ExcPending
	if v2348 != 0 {
		goto L33
	} else {
		goto L724
	}
L723:
	;
	v2364 = F_objectGetVal(m, v1598)
	mBase = m.M
	F_valkey_free(m, v2364)
	mBase = m.M
	v2366 = m.ExcPending
	if v2366 != 0 {
		goto L33
	} else {
		goto L730
	}
L724:
	;
	if v2347 != 0 {
		goto L723
	} else {
		goto L725
	}
L725:
	;
	v2349 = int32(1)
	F_rdbReportError(m, v2349, int32(2552), int32(_a_F_rdbLoadObject_24), int32(0))
	mBase = m.M
	v2355 = m.ExcPending
	if v2355 != 0 {
		goto L33
	} else {
		goto L726
	}
L726:
	;
	v2356 = *(*int32)(unsafe.Add(mBase, uint32(v25)+496))
	F_valkey_free(m, v2356)
	mBase = m.M
	v2358 = m.ExcPending
	if v2358 != 0 {
		goto L33
	} else {
		goto L727
	}
L727:
	;
	F_valkey_free(m, v1593)
	mBase = m.M
	v2360 = m.ExcPending
	if v2360 != 0 {
		goto L33
	} else {
		goto L728
	}
L728:
	;
	F_objectSetVal(m, v1598, int32(0))
	mBase = m.M
	v2363 = m.ExcPending
	if v2363 != 0 {
		goto L33
	} else {
		goto L729
	}
L729:
	;
	v4071 = v2349
	goto L9
L730:
	;
	v2367 = *(*int32)(unsafe.Add(mBase, uint32(v25)+496))
	F_objectSetVal(m, v1598, v2367)
	mBase = m.M
	v2369 = m.ExcPending
	if v2369 != 0 {
		goto L33
	} else {
		goto L731
	}
L731:
	;
	v2370 = int32(180)
	*(*uint8)(unsafe.Add(mBase, uint32(v1598))) = uint8(v2370)
	v2372 = F_hashTypeLength(m, v1598)
	mBase = m.M
	v2373 = m.ExcPending
	if v2373 != 0 {
		goto L33
	} else {
		goto L733
	}
L732:
	;
	v2375 = F_hashTypeLength(m, v1598)
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L33
	} else {
		goto L736
	}
L733:
	;
	if v2372 != 0 {
		goto L732
	} else {
		goto L734
	}
L734:
	;
	v4071 = int32(2)
	goto L9
L735:
	;
	v2383 = F_objectGetVal(m, v1598)
	mBase = m.M
	v2384 = F_lpShrinkToFit(m, v2383)
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L33
	} else {
		goto L739
	}
L736:
	;
	v2378 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[9]))
	if base.Ui32(v2375) <= base.Ui32(v2378) {
		goto L735
	} else {
		goto L737
	}
L737:
	;
	F_hashTypeConvert(m, v1598, int32(2))
	mBase = m.M
	v2382 = m.ExcPending
	if v2382 != 0 {
		goto L33
	} else {
		goto L738
	}
L738:
	;
	v4010 = v1598
	goto L15
L739:
	;
	F_objectSetVal(m, v1598, v2384)
	mBase = m.M
	v2387 = m.ExcPending
	if v2387 != 0 {
		goto L33
	} else {
		goto L740
	}
L740:
	;
	v4010 = v1598
	goto L15
L741:
	;
	v2409 = int32(180)
	*(*uint8)(unsafe.Add(mBase, uint32(v1598))) = uint8(v2409)
	v2411 = F_hashTypeLength(m, v1598)
	mBase = m.M
	v2412 = m.ExcPending
	if v2412 != 0 {
		goto L33
	} else {
		goto L747
	}
L742:
	;
	if v2396 != 0 {
		goto L741
	} else {
		goto L743
	}
L743:
	;
	F_rdbReportError(m, int32(1), int32(2578), int32(_a_F_rdbLoadObject_25), int32(0))
	mBase = m.M
	v2403 = m.ExcPending
	if v2403 != 0 {
		goto L33
	} else {
		goto L744
	}
L744:
	;
	F_valkey_free(m, v1593)
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		goto L33
	} else {
		goto L745
	}
L745:
	;
	F_objectSetVal(m, v1598, int32(0))
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		goto L33
	} else {
		goto L746
	}
L746:
	;
	v4056 = v1598
	goto L10
L747:
	;
	if v2411 == int32(0) {
		v4088 = v1598
		goto L8
	} else {
		goto L748
	}
L748:
	;
	v2415 = F_hashTypeLength(m, v1598)
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L33
	} else {
		goto L749
	}
L749:
	;
	v2418 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[9]))
	if base.Ui32(v2415) <= base.Ui32(v2418) {
		v4010 = v1598
		goto L15
	} else {
		goto L750
	}
L750:
	;
	F_hashTypeConvert(m, v1598, int32(2))
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
		goto L33
	} else {
		goto L751
	}
L751:
	;
	v4010 = v1598
	goto L15
L752:
	;
	v4010 = v1598
	goto L15
L753:
	;
	F_hashtableRelease(m, v1736)
	mBase = m.M
	v2441 = m.ExcPending
	if v2441 != 0 {
		goto L33
	} else {
		goto L754
	}
L754:
	;
	F_sdsfree(m, v1851)
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L33
	} else {
		goto L755
	}
L755:
	;
	F_lpFree(m, v1834)
	mBase = m.M
	v2445 = m.ExcPending
	if v2445 != 0 {
		goto L33
	} else {
		goto L756
	}
L756:
	;
	F_valkey_free(m, v1593)
	mBase = m.M
	v2447 = m.ExcPending
	if v2447 != 0 {
		goto L33
	} else {
		goto L757
	}
L757:
	;
	F_objectSetVal(m, v1598, int32(0))
	mBase = m.M
	v2450 = m.ExcPending
	if v2450 != 0 {
		goto L33
	} else {
		goto L758
	}
L758:
	;
	v4056 = v1598
	goto L10
L759:
	;
	v2453 = F_objectGetVal(m, v2451)
	mBase = m.M
	v2454 = int32(0)
	v2458 = F_rdbLoadLenByRef(m, l1, v2454, v25+int32(192))
	mBase = m.M
	v2459 = m.ExcPending
	if v2459 != 0 {
		goto L33
	} else {
		goto L762
	}
L760:
	;
	v2474 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+512)) = v2474
	if v2462 == v2474 {
		goto L767
	} else {
		goto L768
	}
L761:
	;
	v2466 = int32(0)
	F_rdbReportError(m, v2466, int32(2604), int32(_a_F_rdbLoadObject_26), v2466)
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L33
	} else {
		goto L765
	}
L762:
	;
	if v2458 == int32(-1) {
		goto L761
	} else {
		goto L763
	}
L763:
	;
	v2462 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v2462 != int64(-1) {
		goto L760
	} else {
		goto L764
	}
L764:
	;
	goto L761
L765:
	;
	F_decrRefCount(m, v2451)
	mBase = m.M
	v2473 = m.ExcPending
	if v2473 != 0 {
		goto L33
	} else {
		goto L766
	}
L766:
	;
	v4194 = v2454
	goto L3
L767:
	;
	v2627 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v2628 = m.ExcPending
	if v2628 != 0 {
		goto L33
	} else {
		goto L805
	}
L768:
	;
	v2484 = v2462
	goto L769
L769:
	;
	v2502 = F_rdbGenericLoadStringObject(m, l1, int32(4), int32(0))
	mBase = m.M
	v2503 = m.ExcPending
	if v2503 != 0 {
		goto L33
	} else {
		goto L772
	}
L770:
	;
	F_rdbReportError(m, int32(1), int32(2665), int32(_a_F_rdbLoadObject_27), int32(0))
	mBase = m.M
	v2601 = m.ExcPending
	if v2601 != 0 {
		goto L33
	} else {
		goto L804
	}
L771:
	;
	v2515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2502+int32(-1)))))
	switch v2515 & int32(7) {
	case 0:
		goto L782
	case 1:
		goto L781
	case 2:
		goto L780
	case 3:
		goto L779
	case 4:
		goto L778
	default:
		v2532 = int32(0)
		goto L777
	}
L772:
	;
	if v2502 != 0 {
		goto L771
	} else {
		goto L773
	}
L773:
	;
	v2504 = int32(0)
	F_rdbReportError(m, v2504, int32(2619), int32(_a_F_rdbLoadObject_28), v2504)
	mBase = m.M
	v2509 = m.ExcPending
	if v2509 != 0 {
		goto L33
	} else {
		goto L774
	}
L774:
	;
	v4056 = v2451
	goto L10
L775:
	;
	v2548 = F_rdbGenericLoadStringObject(m, l1, int32(2), v25+int32(192))
	mBase = m.M
	v2549 = m.ExcPending
	if v2549 != 0 {
		goto L33
	} else {
		goto L787
	}
L776:
	;
	if v2534 == int32(16) {
		goto L775
	} else {
		goto L783
	}
L777:
	;
	v2534 = v2532
	goto L776
L778:
	;
	v2531 = *(*int32)(unsafe.Add(mBase, uint32(v2502+int32(-17))))
	v2532 = v2531
	goto L777
L779:
	;
	v2528 = *(*int32)(unsafe.Add(mBase, uint32(v2502+int32(-9))))
	v2534 = v2528
	goto L776
L780:
	;
	v2525 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2502+int32(-5)))))
	v2534 = v2525
	goto L776
L781:
	;
	v2522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2502+int32(-3)))))
	v2534 = v2522
	goto L776
L782:
	;
	v2534 = int32(base.Ui32(v2515) >> (uint(int32(3)) % 32))
	goto L776
L783:
	;
	F_rdbReportError(m, int32(1), int32(2625), int32(_a_F_rdbLoadObject_29), int32(0))
	mBase = m.M
	v2542 = m.ExcPending
	if v2542 != 0 {
		goto L33
	} else {
		goto L784
	}
L784:
	;
	F_sdsfree(m, v2502)
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L33
	} else {
		goto L785
	}
L785:
	;
	v4056 = v2451
	goto L10
L786:
	;
	v2558 = int32(_a_F_rdbLoadObject_3)
	v2560 = *(*int64)(unsafe.Add(mBase, _c_F_rdbLoadObject[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_rdbLoadObject[1])) = v2560 + int64(1)
	v2564 = *(*int32)(unsafe.Add(mBase, uint32(v25)+192))
	v2567 = F_streamValidateListpackIntegrity(m, v2548, v2564, v25+int32(512))
	mBase = m.M
	v2568 = m.ExcPending
	if v2568 != 0 {
		goto L33
	} else {
		goto L792
	}
L787:
	;
	if v2548 != 0 {
		goto L786
	} else {
		goto L788
	}
L788:
	;
	v2550 = int32(0)
	F_rdbReportError(m, v2550, int32(2635), int32(_a_F_rdbLoadObject_30), v2550)
	mBase = m.M
	v2555 = m.ExcPending
	if v2555 != 0 {
		goto L33
	} else {
		goto L789
	}
L789:
	;
	F_sdsfree(m, v2502)
	mBase = m.M
	v2557 = m.ExcPending
	if v2557 != 0 {
		goto L33
	} else {
		goto L790
	}
L790:
	;
	v4056 = v2451
	goto L10
L791:
	;
	v2575 = F_lpFirst(m, v2548)
	mBase = m.M
	v2576 = m.ExcPending
	if v2576 != 0 {
		goto L33
	} else {
		goto L796
	}
L792:
	;
	if v2567 != 0 {
		goto L791
	} else {
		goto L793
	}
L793:
	;
	F_rdbReportError(m, int32(1), int32(2642), int32(_a_F_rdbLoadObject_31), int32(0))
	mBase = m.M
	v2574 = m.ExcPending
	if v2574 != 0 {
		goto L33
	} else {
		goto L794
	}
L794:
	;
	goto L12
L795:
	;
	v2583 = *(*int32)(unsafe.Add(mBase, uint32(v2453)+4))
	v2586 = F_raxTryInsert(m, v2583, v2502, int32(16), v2548, int32(0))
	mBase = m.M
	v2587 = m.ExcPending
	if v2587 != 0 {
		goto L33
	} else {
		goto L799
	}
L796:
	;
	if v2575 != 0 {
		goto L795
	} else {
		goto L797
	}
L797:
	;
	F_rdbReportError(m, int32(1), int32(2654), int32(_a_F_rdbLoadObject_32), int32(0))
	mBase = m.M
	v2582 = m.ExcPending
	if v2582 != 0 {
		goto L33
	} else {
		goto L798
	}
L798:
	;
	goto L12
L799:
	;
	F_sdsfree(m, v2502)
	mBase = m.M
	v2589 = m.ExcPending
	if v2589 != 0 {
		goto L33
	} else {
		goto L800
	}
L800:
	;
	if v2586 == int32(0) {
		goto L801
	} else {
		goto L802
	}
L801:
	;
	goto L770
L802:
	;
	v2593 = v2484 + int64(-1)
	if v2593 == int64(0) {
		goto L767
	} else {
		goto L803
	}
L803:
	;
	v2484 = v2593
	goto L769
L804:
	;
	goto L11
L805:
	;
	v2630 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v2627 == int32(-1) {
		goto L806
	} else {
		goto L807
	}
L806:
	;
	v2633 = int64(-1)
	goto L808
L807:
	;
	v2633 = v2630
	goto L808
L808:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2453)+8)) = v2633
	v2638 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v2639 = m.ExcPending
	if v2639 != 0 {
		goto L33
	} else {
		goto L809
	}
L809:
	;
	v2641 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v2638 == int32(-1) {
		goto L810
	} else {
		goto L811
	}
L810:
	;
	v2644 = int64(-1)
	goto L812
L811:
	;
	v2644 = v2641
	goto L812
L812:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2453)+16)) = v2644
	v2649 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v2650 = m.ExcPending
	if v2650 != 0 {
		goto L33
	} else {
		goto L813
	}
L813:
	;
	v2652 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v2649 == int32(-1) {
		goto L814
	} else {
		goto L815
	}
L814:
	;
	v2655 = int64(-1)
	goto L816
L815:
	;
	v2655 = v2652
	goto L816
L816:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2453)+24)) = v2655
	if l0 < int32(19) {
		goto L818
	} else {
		goto L819
	}
L817:
	;
	v2729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v2729&int32(1) == int32(0) {
		goto L841
	} else {
		goto L842
	}
L818:
	;
	v2714 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2453)+48)) = v2714
	*(*int64)(unsafe.Add(mBase, uint32(v2453+int32(56)))) = v2714
	v2720 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2453)+64)) = v2720
	v2722 = int32(1)
	F_streamGetEdgeID(m, v2453, v2722, v2722, v2453+int32(32))
	mBase = m.M
	v2727 = m.ExcPending
	if v2727 != 0 {
		goto L33
	} else {
		goto L840
	}
L819:
	;
	v2662 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v2663 = m.ExcPending
	if v2663 != 0 {
		goto L33
	} else {
		goto L820
	}
L820:
	;
	v2665 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v2662 == int32(-1) {
		goto L821
	} else {
		goto L822
	}
L821:
	;
	v2668 = int64(-1)
	goto L823
L822:
	;
	v2668 = v2665
	goto L823
L823:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2453)+32)) = v2668
	v2673 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v2674 = m.ExcPending
	if v2674 != 0 {
		goto L33
	} else {
		goto L824
	}
L824:
	;
	v2676 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v2673 == int32(-1) {
		goto L825
	} else {
		goto L826
	}
L825:
	;
	v2679 = int64(-1)
	goto L827
L826:
	;
	v2679 = v2676
	goto L827
L827:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2453)+40)) = v2679
	v2684 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v2685 = m.ExcPending
	if v2685 != 0 {
		goto L33
	} else {
		goto L828
	}
L828:
	;
	v2687 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v2684 == int32(-1) {
		goto L829
	} else {
		goto L830
	}
L829:
	;
	v2690 = int64(-1)
	goto L831
L830:
	;
	v2690 = v2687
	goto L831
L831:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2453)+48)) = v2690
	v2695 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v2696 = m.ExcPending
	if v2696 != 0 {
		goto L33
	} else {
		goto L832
	}
L832:
	;
	v2698 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v2695 == int32(-1) {
		goto L833
	} else {
		goto L834
	}
L833:
	;
	v2701 = int64(-1)
	goto L835
L834:
	;
	v2701 = v2698
	goto L835
L835:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2453)+56)) = v2701
	v2706 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v2707 = m.ExcPending
	if v2707 != 0 {
		goto L33
	} else {
		goto L836
	}
L836:
	;
	v2709 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v2706 == int32(-1) {
		goto L837
	} else {
		goto L838
	}
L837:
	;
	v2712 = int64(-1)
	goto L839
L838:
	;
	v2712 = v2709
	goto L839
L839:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2453)+64)) = v2712
	goto L817
L840:
	;
	goto L817
L841:
	;
	v2740 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+8))
	if base.B2i32(v2740 == int64(0)) == int32(0) {
		goto L846
	} else {
		goto L847
	}
L842:
	;
	v2734 = int32(0)
	F_rdbReportError(m, v2734, int32(2703), int32(_a_F_rdbLoadObject_33), v2734)
	mBase = m.M
	v2739 = m.ExcPending
	if v2739 != 0 {
		goto L33
	} else {
		goto L843
	}
L843:
	;
	v4056 = v2451
	goto L10
L844:
	;
	F_rdbReportError(m, int32(1), int32(2709), int32(_a_F_rdbLoadObject_34), int32(0))
	mBase = m.M
	v3690 = m.ExcPending
	if v3690 != 0 {
		goto L33
	} else {
		goto L1088
	}
L845:
	;
	v2752 = *(*int64)(unsafe.Add(mBase, uint32(v25)+512))
	if v2751 == v2752 {
		goto L850
	} else {
		goto L851
	}
L846:
	;
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v2453)+4))
	v2747 = *(*int64)(unsafe.Add(mBase, uint32(v2746)+8))
	goto L848
L847:
	;
	v2751 = int64(0)
	goto L845
L848:
	;
	if v2747 == int64(0) {
		goto L844
	} else {
		goto L849
	}
L849:
	;
	v2750 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+8))
	v2751 = v2750
	goto L845
L850:
	;
	v2763 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v2764 = m.ExcPending
	if v2764 != 0 {
		goto L33
	} else {
		goto L854
	}
L851:
	;
	F_rdbReportError(m, int32(1), int32(2722), int32(_a_F_rdbLoadObject_35), int32(0))
	mBase = m.M
	v2759 = m.ExcPending
	if v2759 != 0 {
		goto L33
	} else {
		goto L852
	}
L852:
	;
	v4056 = v2451
	goto L10
L853:
	;
	v3679 = int32(0)
	F_rdbReportError(m, v3679, int32(2730), int32(_a_F_rdbLoadObject_36), v3679)
	mBase = m.M
	v3684 = m.ExcPending
	if v3684 != 0 {
		goto L33
	} else {
		goto L1087
	}
L854:
	;
	if v2763 == int32(-1) {
		goto L853
	} else {
		goto L855
	}
L855:
	;
	v2767 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v2767 == int64(-1) {
		goto L853
	} else {
		goto L856
	}
L856:
	;
	if v2767 == int64(0) {
		v4010 = v2451
		goto L15
	} else {
		goto L857
	}
L857:
	;
	v2786 = v2767
	goto L858
L858:
	;
	v2800 = F_rdbGenericLoadStringObject(m, l1, int32(4), int32(0))
	mBase = m.M
	v2801 = m.ExcPending
	if v2801 != 0 {
		goto L33
	} else {
		goto L861
	}
L860:
	;
	v2811 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v2812 = m.ExcPending
	if v2812 != 0 {
		goto L33
	} else {
		goto L864
	}
L861:
	;
	if v2800 != 0 {
		goto L860
	} else {
		goto L862
	}
L862:
	;
	v2802 = int32(0)
	F_rdbReportError(m, v2802, int32(2741), int32(_a_F_rdbLoadObject_37), v2802)
	mBase = m.M
	v2807 = m.ExcPending
	if v2807 != 0 {
		goto L33
	} else {
		goto L863
	}
L863:
	;
	v4056 = v2451
	goto L10
L864:
	;
	v2814 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v2811 == int32(-1) {
		goto L865
	} else {
		goto L866
	}
L865:
	;
	v2817 = int64(-1)
	goto L867
L866:
	;
	v2817 = v2814
	goto L867
L867:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+496)) = v2817
	v2822 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v2823 = m.ExcPending
	if v2823 != 0 {
		goto L33
	} else {
		goto L868
	}
L868:
	;
	v2825 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v2822 == int32(-1) {
		goto L869
	} else {
		goto L870
	}
L869:
	;
	v2828 = int64(-1)
	goto L871
L870:
	;
	v2828 = v2825
	goto L871
L871:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+504)) = v2828
	v2830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v2830&int32(1) == int32(0) {
		goto L872
	} else {
		goto L873
	}
L872:
	;
	if l0 < int32(19) {
		goto L877
	} else {
		goto L878
	}
L873:
	;
	v2835 = int32(0)
	F_rdbReportError(m, v2835, int32(2749), int32(_a_F_rdbLoadObject_38), v2835)
	mBase = m.M
	v2840 = m.ExcPending
	if v2840 != 0 {
		goto L33
	} else {
		goto L874
	}
L874:
	;
	F_sdsfree(m, v2800)
	mBase = m.M
	v2842 = m.ExcPending
	if v2842 != 0 {
		goto L33
	} else {
		goto L875
	}
L875:
	;
	v4056 = v2451
	goto L10
L876:
	;
	v2998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2800+int32(-1)))))
	switch v2998 & int32(7) {
	case 0:
		goto L944
	case 1:
		goto L943
	case 2:
		goto L942
	case 3:
		goto L941
	case 4:
		goto L940
	default:
		v3015 = int32(0)
		goto L939
	}
L877:
	;
	v2865 = v25 + int32(496)
	v2873 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+64))
	if base.B2i32(v2873 == int64(0)) == int32(0) {
		goto L888
	} else {
		goto L889
	}
L878:
	;
	v2846 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v2847 = m.ExcPending
	if v2847 != 0 {
		goto L33
	} else {
		goto L879
	}
L879:
	;
	v2848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v2848&int32(1) != 0 {
		goto L880
	} else {
		goto L881
	}
L880:
	;
	v2856 = int32(0)
	F_rdbReportError(m, v2856, int32(2760), int32(_a_F_rdbLoadObject_39), v2856)
	mBase = m.M
	v2861 = m.ExcPending
	if v2861 != 0 {
		goto L33
	} else {
		goto L885
	}
L881:
	;
	v2852 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v2846 == int32(-1) {
		goto L882
	} else {
		goto L883
	}
L882:
	;
	v2855 = int64(-1)
	goto L884
L883:
	;
	v2855 = v2852
	goto L884
L884:
	;
	v2992 = v2855
	goto L876
L885:
	;
	F_sdsfree(m, v2800)
	mBase = m.M
	v2863 = m.ExcPending
	if v2863 != 0 {
		goto L33
	} else {
		goto L886
	}
L886:
	;
	v4056 = v2451
	goto L10
L887:
	;
	v2992 = v2990
	goto L876
L888:
	;
	v2879 = *(*int64)(unsafe.Add(mBase, uint32(v2865)))
	v2880 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+8))
	if v2880 != int64(0) {
		goto L895
	} else {
		goto L896
	}
L889:
	;
	v2990 = int64(0)
	goto L887
L890:
	;
	v2990 = v2980
	goto L887
L891:
	;
	if base.Ui64(v2915) < base.Ui64(v2916) {
		goto L908
	} else {
		goto L909
	}
L892:
	;
	v2912 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+16))
	if base.Ui64(v2909) <= base.Ui64(v2912) {
		v2915 = v2909
		v2916 = v2912
		goto L891
	} else {
		goto L907
	}
L893:
	;
	v2904 = int64(-1)
	if base.Ui64(v2902) < base.Ui64(v2903) {
		v2980 = v2904
		goto L890
	} else {
		goto L905
	}
L894:
	;
	v2900 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+48))
	if base.Ui64(v2900) < base.Ui64(v2879) {
		v2909 = v2879
		goto L892
	} else {
		goto L904
	}
L895:
	;
	if v2879 != int64(0) {
		goto L894
	} else {
		goto L901
	}
L896:
	;
	v2883 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+16))
	if base.Ui64(v2883) < base.Ui64(v2879) {
		goto L894
	} else {
		goto L897
	}
L897:
	;
	if base.Ui64(v2883) <= base.Ui64(v2879) {
		goto L898
	} else {
		goto L899
	}
L898:
	;
	v2886 = *(*int64)(unsafe.Add(mBase, uint32(v2865)+8))
	v2887 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+24))
	if base.Ui64(v2887) < base.Ui64(v2886) {
		goto L895
	} else {
		goto L900
	}
L899:
	;
	v2990 = v2873
	goto L887
L900:
	;
	v2990 = v2873
	goto L887
L901:
	;
	v2892 = int64(0)
	v2893 = *(*int64)(unsafe.Add(mBase, uint32(v2865)+8))
	if v2893 != v2892 {
		goto L902
	} else {
		goto L903
	}
L902:
	;
	v2898 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+48))
	v2902 = v2892
	v2903 = v2898
	goto L893
L903:
	;
	v2896 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+16))
	v2915 = int64(0)
	v2916 = v2896
	goto L891
L904:
	;
	v2902 = v2879
	v2903 = v2900
	goto L893
L905:
	;
	v2906 = *(*int64)(unsafe.Add(mBase, uint32(v2865)+8))
	v2907 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+56))
	if base.Ui64(v2906) < base.Ui64(v2907) {
		v2980 = v2904
		goto L890
	} else {
		goto L906
	}
L906:
	;
	v2909 = v2902
	goto L892
L907:
	;
	v2990 = int64(-1)
	goto L887
L908:
	;
	v2926 = int32(1)
	v2928 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+32))
	if base.Ui64(v2928) < base.Ui64(v2915) {
		v2942 = v2926
		goto L913
	} else {
		goto L914
	}
L909:
	;
	v2919 = *(*int64)(unsafe.Add(mBase, uint32(v2865)+8))
	v2920 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+24))
	if base.Ui64(v2919) <= base.Ui64(v2920) {
		goto L910
	} else {
		goto L911
	}
L910:
	;
	if base.Ui64(v2919) < base.Ui64(v2920) {
		goto L908
	} else {
		goto L912
	}
L911:
	;
	v2990 = int64(-1)
	goto L887
L912:
	;
	v2990 = v2873
	goto L887
L913:
	;
	v2943 = int32(0)
	v2944 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+48))
	if base.Ui64(v2928) < base.Ui64(v2944) {
		v2964 = v2943
		v2967 = v2926
		goto L920
	} else {
		goto L921
	}
L914:
	;
	if base.Ui64(v2915) < base.Ui64(v2928) {
		v2942 = int32(-1)
		goto L913
	} else {
		goto L915
	}
L915:
	;
	v2933 = *(*int64)(unsafe.Add(mBase, uint32(v2865)+8))
	v2934 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+40))
	if base.Ui64(v2934) < base.Ui64(v2933) {
		v2942 = int32(1)
		goto L913
	} else {
		goto L916
	}
L916:
	;
	if base.Ui64(v2933) < base.Ui64(v2934) {
		goto L917
	} else {
		goto L918
	}
L917:
	;
	v2939 = int32(-1)
	goto L919
L918:
	;
	v2939 = int32(0)
	goto L919
L919:
	;
	v2942 = v2939
	goto L913
L920:
	;
	v2968 = int64(-1)
	if v2964 != 0 {
		goto L931
	} else {
		goto L932
	}
L921:
	;
	if base.Ui64(v2928) <= base.Ui64(v2944) {
		goto L923
	} else {
		goto L924
	}
L922:
	;
	if v2944 != int64(0) {
		v2964 = v2943
		v2967 = v2958
		goto L920
	} else {
		goto L930
	}
L923:
	;
	v2948 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+56))
	v2949 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+40))
	if base.Ui64(v2948) <= base.Ui64(v2949) {
		goto L925
	} else {
		goto L926
	}
L924:
	;
	v2958 = int32(-1)
	goto L922
L925:
	;
	if base.Ui64(v2948) < base.Ui64(v2949) {
		goto L927
	} else {
		goto L928
	}
L926:
	;
	v2958 = int32(1)
	goto L922
L927:
	;
	v2955 = int32(-1)
	goto L929
L928:
	;
	v2955 = int32(0)
	goto L929
L929:
	;
	v2958 = v2955
	goto L922
L930:
	;
	v2961 = *(*int64)(unsafe.Add(mBase, uint32(v2453)+56))
	v2964 = base.B2i32(v2961 == int64(0))
	v2967 = v2958
	goto L920
L931:
	;
	if int32(-1) < v2942 {
		goto L934
	} else {
		goto L935
	}
L932:
	;
	if int32(-1) < v2967 {
		v2980 = v2968
		goto L890
	} else {
		goto L933
	}
L933:
	;
	goto L931
L934:
	;
	if v2942 != 0 {
		v2980 = v2968
		goto L890
	} else {
		goto L936
	}
L935:
	;
	v2990 = v2873 - v2880
	goto L887
L936:
	;
	v2980 = v2873 - v2880 + int64(1)
	goto L890
L937:
	;
	F_sdsfree(m, v2800)
	mBase = m.M
	v3046 = m.ExcPending
	if v3046 != 0 {
		goto L33
	} else {
		goto L954
	}
L938:
	;
	v3020 = F_streamCreateCG(m, v2453, v2800, v3017, v25+int32(496), v2992)
	mBase = m.M
	v3021 = m.ExcPending
	if v3021 != 0 {
		goto L33
	} else {
		goto L945
	}
L939:
	;
	v3017 = v3015
	goto L938
L940:
	;
	v3014 = *(*int32)(unsafe.Add(mBase, uint32(v2800+int32(-17))))
	v3015 = v3014
	goto L939
L941:
	;
	v3011 = *(*int32)(unsafe.Add(mBase, uint32(v2800+int32(-9))))
	v3017 = v3011
	goto L938
L942:
	;
	v3008 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2800+int32(-5)))))
	v3017 = v3008
	goto L938
L943:
	;
	v3005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2800+int32(-3)))))
	v3017 = v3005
	goto L938
L944:
	;
	v3017 = int32(base.Ui32(v2998) >> (uint(int32(3)) % 32))
	goto L938
L945:
	;
	if v3020 != 0 {
		goto L937
	} else {
		goto L946
	}
L946:
	;
	v3023 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[15]))
	if v3023 == int32(0) {
		goto L948
	} else {
		goto L949
	}
L947:
	;
	F_decrRefCount(m, v2451)
	mBase = m.M
	v3041 = m.ExcPending
	if v3041 != 0 {
		goto L33
	} else {
		goto L952
	}
L948:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v2800
	F_rdbReportError(m, int32(1), int32(2774), int32(_a_F_rdbLoadObject_40), v25+int32(48))
	mBase = m.M
	v3039 = m.ExcPending
	if v3039 != 0 {
		goto L33
	} else {
		goto L951
	}
L949:
	;
	F_rdbReportError(m, int32(1), int32(2772), int32(_a_F_rdbLoadObject_41), int32(0))
	mBase = m.M
	v3031 = m.ExcPending
	if v3031 != 0 {
		goto L33
	} else {
		goto L950
	}
L950:
	;
	goto L947
L951:
	;
	goto L947
L952:
	;
	F_sdsfree(m, v2800)
	mBase = m.M
	v3043 = m.ExcPending
	if v3043 != 0 {
		goto L33
	} else {
		goto L953
	}
L953:
	;
	v4194 = int32(0)
	goto L3
L954:
	;
	v3050 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v3051 = m.ExcPending
	if v3051 != 0 {
		goto L33
	} else {
		goto L958
	}
L955:
	;
	v3602 = v2786 + int64(-1)
	v3604 = v25 + int32(192)
	v3605 = *(*int32)(unsafe.Add(mBase, uint32(v3020)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3604)+4)) = v3605
	*(*int32)(unsafe.Add(mBase, uint32(v3604))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v3604)+20)) = int32(128)
	v3611 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3604)+12)) = v3611
	*(*int64)(unsafe.Add(mBase, uint32(v3604)+296)) = v3611
	*(*int64)(unsafe.Add(mBase, uint32(v3604)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v3604)+8)) = v25 + int32(216)
	*(*int32)(unsafe.Add(mBase, uint32(v3604)+156)) = v25 + int32(360)
	goto L1075
L956:
	;
	v3573 = int32(0)
	F_rdbReportError(m, v3573, int32(2864), int32(_a_F_rdbLoadObject_42), v3573)
	mBase = m.M
	v3578 = m.ExcPending
	if v3578 != 0 {
		goto L33
	} else {
		goto L1074
	}
L957:
	;
	v3566 = int32(0)
	F_rdbReportError(m, v3566, int32(2789), int32(_a_F_rdbLoadObject_43), v3566)
	mBase = m.M
	v3571 = m.ExcPending
	if v3571 != 0 {
		goto L33
	} else {
		goto L1073
	}
L958:
	;
	if v3050 == int32(-1) {
		goto L957
	} else {
		goto L959
	}
L959:
	;
	v3054 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v3054 == int64(-1) {
		goto L957
	} else {
		goto L960
	}
L960:
	;
	if v3054 == int64(0) {
		goto L961
	} else {
		goto L962
	}
L961:
	;
	v3162 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v3163 = m.ExcPending
	if v3163 != 0 {
		goto L33
	} else {
		goto L985
	}
L962:
	;
	v3065 = v3054
	goto L963
L963:
	;
	v3084 = F_rioRead_1(m, l1, v25+int32(192), int32(16))
	mBase = m.M
	v3085 = m.ExcPending
	if v3085 != 0 {
		goto L33
	} else {
		goto L966
	}
L964:
	;
	goto L961
L965:
	;
	v3093 = F_streamCreateNACK(m, int32(0))
	mBase = m.M
	v3094 = m.ExcPending
	if v3094 != 0 {
		goto L33
	} else {
		goto L969
	}
L966:
	;
	if v3084 != 0 {
		goto L965
	} else {
		goto L967
	}
L967:
	;
	v3086 = int32(0)
	F_rdbReportError(m, v3086, int32(2796), int32(_a_F_rdbLoadObject_44), v3086)
	mBase = m.M
	v3091 = m.ExcPending
	if v3091 != 0 {
		goto L33
	} else {
		goto L968
	}
L968:
	;
	v4056 = v2451
	goto L10
L969:
	;
	v3095 = F_rdbLoadMillisecondTime(m, l1, l1)
	mBase = m.M
	v3096 = m.ExcPending
	if v3096 != 0 {
		goto L33
	} else {
		goto L970
	}
L970:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3093))) = v3095
	v3101 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(528))
	mBase = m.M
	v3102 = m.ExcPending
	if v3102 != 0 {
		goto L33
	} else {
		goto L971
	}
L971:
	;
	v3104 = *(*int64)(unsafe.Add(mBase, uint32(v25)+528))
	if v3101 == int32(-1) {
		goto L972
	} else {
		goto L973
	}
L972:
	;
	v3107 = int64(-1)
	goto L974
L973:
	;
	v3107 = v3104
	goto L974
L974:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3093)+8)) = v3107
	v3109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v3109&int32(1) == int32(0) {
		goto L975
	} else {
		goto L976
	}
L975:
	;
	v3120 = *(*int32)(unsafe.Add(mBase, uint32(v3020)+24))
	v3125 = F_raxTryInsert(m, v3120, v25+int32(192), int32(16), v3093, int32(0))
	mBase = m.M
	v3126 = m.ExcPending
	if v3126 != 0 {
		goto L33
	} else {
		goto L979
	}
L976:
	;
	v3114 = int32(0)
	F_rdbReportError(m, v3114, int32(2804), int32(_a_F_rdbLoadObject_45), v3114)
	mBase = m.M
	v3119 = m.ExcPending
	if v3119 != 0 {
		goto L33
	} else {
		goto L977
	}
L977:
	;
	goto L13
L978:
	;
	v3134 = v3065 + int64(-1)
	if v3134 != int64(0) {
		v3065 = v3134
		goto L963
	} else {
		goto L982
	}
L979:
	;
	if v3125 != 0 {
		goto L978
	} else {
		goto L980
	}
L980:
	;
	F_rdbReportError(m, int32(1), int32(2811), int32(_a_F_rdbLoadObject_46), int32(0))
	mBase = m.M
	v3132 = m.ExcPending
	if v3132 != 0 {
		goto L33
	} else {
		goto L981
	}
L981:
	;
	goto L13
L982:
	;
	goto L964
L983:
	;
	v3193 = v3166
	goto L990
L984:
	;
	v3174 = int32(0)
	F_rdbReportError(m, v3174, int32(2822), int32(_a_F_rdbLoadObject_47), v3174)
	mBase = m.M
	v3179 = m.ExcPending
	if v3179 != 0 {
		goto L33
	} else {
		goto L989
	}
L985:
	;
	if v3162 == int32(-1) {
		goto L984
	} else {
		goto L986
	}
L986:
	;
	v3166 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v3166 == int64(-1) {
		goto L984
	} else {
		goto L987
	}
L987:
	;
	if base.B2i32(v3166 == int64(0)) == int32(0) {
		goto L983
	} else {
		goto L988
	}
L988:
	;
	goto L955
L989:
	;
	v4056 = v2451
	goto L10
L990:
	;
	v3204 = F_rdbGenericLoadStringObject(m, l1, int32(4), int32(0))
	mBase = m.M
	v3205 = m.ExcPending
	if v3205 != 0 {
		goto L33
	} else {
		goto L993
	}
L991:
	;
	F_rdbReportError(m, int32(1), int32(2897), int32(_a_F_rdbLoadObject_48), int32(0))
	mBase = m.M
	v3564 = m.ExcPending
	if v3564 != 0 {
		goto L33
	} else {
		goto L1072
	}
L992:
	;
	v3212 = int32(0)
	v3215 = F_streamCreateConsumer(m, v3020, v3204, v3212, v3212, int32(3))
	mBase = m.M
	v3216 = m.ExcPending
	if v3216 != 0 {
		goto L33
	} else {
		goto L996
	}
L993:
	;
	if v3204 != 0 {
		goto L992
	} else {
		goto L994
	}
L994:
	;
	v3206 = int32(0)
	F_rdbReportError(m, v3206, int32(2829), int32(_a_F_rdbLoadObject_49), v3206)
	mBase = m.M
	v3211 = m.ExcPending
	if v3211 != 0 {
		goto L33
	} else {
		goto L995
	}
L995:
	;
	v4056 = v2451
	goto L10
L996:
	;
	F_sdsfree(m, v3204)
	mBase = m.M
	v3218 = m.ExcPending
	if v3218 != 0 {
		goto L33
	} else {
		goto L997
	}
L997:
	;
	if v3215 != 0 {
		goto L998
	} else {
		goto L999
	}
L998:
	;
	v3225 = F_rdbLoadMillisecondTime(m, l1, l1)
	mBase = m.M
	v3226 = m.ExcPending
	if v3226 != 0 {
		goto L33
	} else {
		goto L1001
	}
L999:
	;
	F_rdbReportError(m, int32(1), int32(2836), int32(_a_F_rdbLoadObject_50), int32(0))
	mBase = m.M
	v3224 = m.ExcPending
	if v3224 != 0 {
		goto L33
	} else {
		goto L1000
	}
L1000:
	;
	v4056 = v2451
	goto L10
L1001:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3215))) = v3225
	v3228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v3228&int32(1) == int32(0) {
		goto L1002
	} else {
		goto L1003
	}
L1002:
	;
	if l0 < int32(21) {
		goto L1006
	} else {
		goto L1007
	}
L1003:
	;
	v3233 = int32(0)
	F_rdbReportError(m, v3233, int32(2843), int32(_a_F_rdbLoadObject_51), v3233)
	mBase = m.M
	v3238 = m.ExcPending
	if v3238 != 0 {
		goto L33
	} else {
		goto L1004
	}
L1004:
	;
	v4056 = v2451
	goto L10
L1005:
	;
	v3257 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(192))
	mBase = m.M
	v3258 = m.ExcPending
	if v3258 != 0 {
		goto L33
	} else {
		goto L1011
	}
L1006:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3215)+8)) = v3225
	goto L1005
L1007:
	;
	v3239 = F_rdbLoadMillisecondTime(m, l1, l1)
	mBase = m.M
	v3240 = m.ExcPending
	if v3240 != 0 {
		goto L33
	} else {
		goto L1008
	}
L1008:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3215)+8)) = v3239
	v3242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v3242&int32(1) == int32(0) {
		goto L1005
	} else {
		goto L1009
	}
L1009:
	;
	v3247 = int32(0)
	F_rdbReportError(m, v3247, int32(2851), int32(_a_F_rdbLoadObject_52), v3247)
	mBase = m.M
	v3252 = m.ExcPending
	if v3252 != 0 {
		goto L33
	} else {
		goto L1010
	}
L1010:
	;
	v4056 = v2451
	goto L10
L1011:
	;
	if v3257 == int32(-1) {
		goto L956
	} else {
		goto L1012
	}
L1012:
	;
	v3261 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v3261 == int64(-1) {
		goto L956
	} else {
		goto L1013
	}
L1013:
	;
	if v3261 == int64(0) {
		goto L1015
	} else {
		goto L1016
	}
L1014:
	;
	goto L991
L1015:
	;
	v3556 = v3193 + int64(-1)
	if v3556 == int64(0) {
		goto L955
	} else {
		goto L1071
	}
L1016:
	;
	v3272 = v3261
	goto L1017
L1017:
	;
	v3291 = F_rioRead_1(m, l1, v25+int32(192), int32(16))
	mBase = m.M
	v3292 = m.ExcPending
	if v3292 != 0 {
		goto L33
	} else {
		goto L1020
	}
L1018:
	;
	goto L1015
L1019:
	;
	v3299 = *(*int32)(unsafe.Add(mBase, uint32(v3020)+24))
	v3301 = v25 + int32(192)
	v3302 = int32(16)
	v3304 = v25 + int32(528)
	v3313 = *(*int32)(unsafe.Add(mBase, uint32(v3299)))
	v3314 = *(*int32)(unsafe.Add(mBase, uint32(v3313)))
	goto L1027
L1020:
	;
	if v3291 != 0 {
		goto L1019
	} else {
		goto L1021
	}
L1021:
	;
	v3293 = int32(0)
	F_rdbReportError(m, v3293, int32(2871), int32(_a_F_rdbLoadObject_53), v3293)
	mBase = m.M
	v3298 = m.ExcPending
	if v3298 != 0 {
		goto L33
	} else {
		goto L1022
	}
L1022:
	;
	v4056 = v2451
	goto L10
L1023:
	;
	v3506 = *(*int32)(unsafe.Add(mBase, uint32(v25)+528))
	v3507 = *(*int32)(unsafe.Add(mBase, uint32(v3506)+16))
	if v3507 == int32(0) {
		goto L1064
	} else {
		goto L1065
	}
L1024:
	;
	if v3498 != 0 {
		goto L1023
	} else {
		goto L1062
	}
L1025:
	;
	if v3457 != v3302 {
		v3498 = int32(0)
		goto L1052
	} else {
		goto L1053
	}
L1026:
	;
	v3448 = int32(0)
	v3454 = v3313
	v3455 = v3314
	v3457 = v3448
	v3461 = v3448
	goto L1025
L1027:
	;
	if base.Ui32(v3314) < base.Ui32(int32(8)) {
		goto L1026
	} else {
		goto L1028
	}
L1028:
	;
	v3325 = v3313
	v3326 = v3314
	v3328 = int32(0)
	goto L1030
L1029:
	;
	v3454 = v3438
	v3455 = v3439
	v3457 = v3441
	v3461 = base.B2i32(v3444 != int32(0))
	goto L1025
L1030:
	;
	v3334 = int32(base.Ui32(v3326) >> (uint(int32(3)) % 32))
	v3335 = int32(4)
	v3336 = v3325 + v3335
	if v3326&v3335 == int32(0) {
		goto L1033
	} else {
		goto L1034
	}
L1031:
	;
	v3438 = v3429
	v3439 = v3430
	v3441 = v3414
	v3444 = v3419
	goto L1029
L1032:
	;
	v3419 = int32(0)
	v3429 = *(*int32)(unsafe.Add(mBase, uint32(v3336+v3334+(v3419-v3334)&int32(3)+v3407<<(uint(int32(2))%32))))
	v3430 = *(*int32)(unsafe.Add(mBase, uint32(v3429)))
	if base.Ui32(v3430) < base.Ui32(int32(8)) {
		v3438 = v3429
		v3439 = v3430
		v3441 = v3414
		v3444 = v3419
		goto L1029
	} else {
		goto L1050
	}
L1033:
	;
	v3382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3301+v3328))))
	v3385 = int32(0)
	goto L1044
L1034:
	;
	v3341 = int32(0)
	if base.Ui32(v3302) <= base.Ui32(v3328) {
		v3374 = v3328
		v3377 = v3341
		goto L1035
	} else {
		goto L1036
	}
L1035:
	;
	if v3377 == v3334 {
		v3407 = v3341
		v3414 = v3374
		goto L1032
	} else {
		goto L1042
	}
L1036:
	;
	v3351 = v3328
	v3354 = v3341
	goto L1037
L1037:
	;
	v3357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3336+v3354))))
	v3359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3301+v3351))))
	if v3357 != v3359 {
		v3374 = v3351
		v3377 = v3354
		goto L1035
	} else {
		goto L1039
	}
L1038:
	;
	v3374 = v3362
	v3377 = v3364
	goto L1035
L1039:
	;
	v3361 = int32(1)
	v3362 = v3351 + v3361
	v3364 = v3354 + v3361
	if base.Ui32(v3334) <= base.Ui32(v3364) {
		v3374 = v3362
		v3377 = v3364
		goto L1035
	} else {
		goto L1040
	}
L1040:
	;
	if base.Ui32(v3362) < base.Ui32(v3302) {
		v3351 = v3362
		v3354 = v3364
		goto L1037
	} else {
		goto L1041
	}
L1041:
	;
	goto L1038
L1042:
	;
	v3438 = v3325
	v3439 = v3326
	v3441 = v3374
	v3444 = v3377
	goto L1029
L1043:
	;
	if v3385 != v3334 {
		goto L1048
	} else {
		goto L1049
	}
L1044:
	;
	v3398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3336+v3385))))
	if v3398 == v3382&int32(255) {
		goto L1043
	} else {
		goto L1046
	}
L1046:
	;
	v3400 = int32(1)
	v3402 = v3385 + v3400
	if v3402 != v3334 {
		v3385 = v3402
		goto L1044
	} else {
		goto L1047
	}
L1047:
	;
	v3454 = v3325
	v3455 = v3326
	v3457 = v3328
	v3461 = v3400
	goto L1025
L1048:
	;
	v3407 = v3385
	v3414 = v3328 + int32(1)
	goto L1032
L1049:
	;
	v3438 = v3325
	v3439 = v3326
	v3441 = v3328
	v3444 = v3334
	goto L1029
L1050:
	;
	if base.Ui32(v3414) < base.Ui32(v3302) {
		v3325 = v3429
		v3326 = v3430
		v3328 = v3414
		goto L1030
	} else {
		goto L1051
	}
L1051:
	;
	goto L1031
L1052:
	;
	goto L1024
L1053:
	;
	v3463 = int32(0)
	if v3455&int32(1) == v3463 {
		v3498 = v3463
		goto L1052
	} else {
		goto L1054
	}
L1054:
	;
	v3469 = v3455 & int32(4)
	if v3461&base.B2i32(v3469 != int32(0)) != 0 {
		v3498 = v3463
		goto L1052
	} else {
		goto L1055
	}
L1055:
	;
	v3473 = int32(1)
	if v3304 == int32(0) {
		v3498 = v3473
		goto L1052
	} else {
		goto L1056
	}
L1056:
	;
	if v3455&int32(2) != 0 {
		v3495 = int32(0)
		goto L1057
	} else {
		goto L1058
	}
L1057:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3304))) = v3495
	v3498 = v3473
	goto L1052
L1058:
	;
	v3479 = int32(3)
	v3480 = int32(base.Ui32(v3455) >> (uint(v3479) % 32))
	if v3469 != 0 {
		goto L1059
	} else {
		goto L1060
	}
L1059:
	;
	v3490 = int32(4)
	goto L1061
L1060:
	;
	v3490 = v3480 << (uint(int32(2)) % 32)
	goto L1061
L1061:
	;
	v3494 = *(*int32)(unsafe.Add(mBase, uint32(v3454+v3480+(int32(0)-v3480)&v3479+v3490+int32(4))))
	v3495 = v3494
	goto L1057
L1062:
	;
	F_rdbReportError(m, int32(1), int32(2878), int32(_a_F_rdbLoadObject_54), int32(0))
	mBase = m.M
	v3505 = m.ExcPending
	if v3505 != 0 {
		goto L33
	} else {
		goto L1063
	}
L1063:
	;
	v4056 = v2451
	goto L10
L1064:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3506)+16)) = v3215
	v3518 = *(*int32)(unsafe.Add(mBase, uint32(v3215)+20))
	v3523 = F_raxTryInsert(m, v3518, v25+int32(192), int32(16), v3506, int32(0))
	mBase = m.M
	v3524 = m.ExcPending
	if v3524 != 0 {
		goto L33
	} else {
		goto L1068
	}
L1065:
	;
	if v3507 == v3215 {
		goto L1064
	} else {
		goto L1066
	}
L1066:
	;
	F_rdbReportError(m, int32(1), int32(2889), int32(_a_F_rdbLoadObject_55), int32(0))
	mBase = m.M
	v3516 = m.ExcPending
	if v3516 != 0 {
		goto L33
	} else {
		goto L1067
	}
L1067:
	;
	v4056 = v2451
	goto L10
L1068:
	;
	if v3523 == int32(0) {
		goto L1014
	} else {
		goto L1069
	}
L1069:
	;
	v3528 = v3272 + int64(-1)
	if base.B2i32(v3528 == int64(0)) == int32(0) {
		v3272 = v3528
		goto L1017
	} else {
		goto L1070
	}
L1070:
	;
	goto L1018
L1071:
	;
	v3193 = v3556
	goto L990
L1072:
	;
	v4056 = v2451
	goto L10
L1073:
	;
	v4056 = v2451
	goto L10
L1074:
	;
	v4056 = v2451
	goto L10
L1075:
	;
	v3626 = int32(0)
	v3628 = F_raxSeek(m, v25+int32(192), int32(_a_F_rdbLoadObject_56), v3626, v3626)
	mBase = m.M
	v3629 = m.ExcPending
	if v3629 != 0 {
		goto L33
	} else {
		goto L1076
	}
L1076:
	;
	goto L1078
L1077:
	;
	F_raxStop(m, v25+int32(192))
	mBase = m.M
	v3673 = m.ExcPending
	if v3673 != 0 {
		goto L33
	} else {
		goto L1085
	}
L1078:
	;
	v3654 = F_raxNext(m, v25+int32(192))
	mBase = m.M
	v3655 = m.ExcPending
	if v3655 != 0 {
		goto L33
	} else {
		goto L1080
	}
L1079:
	;
	F_raxStop(m, v25+int32(192))
	mBase = m.M
	v3663 = m.ExcPending
	if v3663 != 0 {
		goto L33
	} else {
		goto L1083
	}
L1080:
	;
	if v3654 == int32(0) {
		goto L1077
	} else {
		goto L1081
	}
L1081:
	;
	v3658 = *(*int32)(unsafe.Add(mBase, uint32(v25)+204))
	v3659 = *(*int32)(unsafe.Add(mBase, uint32(v3658)+16))
	if v3659 != 0 {
		goto L1078
	} else {
		goto L1082
	}
L1082:
	;
	goto L1079
L1083:
	;
	F_rdbReportError(m, int32(1), int32(2912), int32(_a_F_rdbLoadObject_57), int32(0))
	mBase = m.M
	v3669 = m.ExcPending
	if v3669 != 0 {
		goto L33
	} else {
		goto L1084
	}
L1084:
	;
	v4056 = v2451
	goto L10
L1085:
	;
	if base.B2i32(v3602 == int64(0)) == int32(0) {
		v2786 = v3602
		goto L858
	} else {
		goto L1086
	}
L1086:
	;
	v4010 = v2451
	goto L15
L1087:
	;
	v4056 = v2451
	goto L10
L1088:
	;
	v4056 = v2451
	goto L10
L1089:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = l0
	v3701 = int32(0)
	F_rdbReportError(m, v3701, int32(2986), int32(_a_F_rdbLoadObject_2), v25)
	mBase = m.M
	v3706 = m.ExcPending
	if v3706 != 0 {
		goto L33
	} else {
		goto L1092
	}
L1090:
	;
	v3695 = int32(0)
	if l4 == v3695 {
		v4194 = v3695
		goto L3
	} else {
		goto L1091
	}
L1091:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
	v4194 = v3695
	goto L3
L1092:
	;
	v4194 = v3701
	goto L3
L1093:
	;
	v3713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v3713&int32(1) == int32(0) {
		goto L1094
	} else {
		goto L1095
	}
L1094:
	;
	v3725 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v3711 == int32(-1) {
		goto L1097
	} else {
		goto L1098
	}
L1095:
	;
	v3718 = int32(0)
	F_rdbReportError(m, v3718, int32(2925), int32(_a_F_rdbLoadObject_58), v3718)
	mBase = m.M
	v3723 = m.ExcPending
	if v3723 != 0 {
		goto L33
	} else {
		goto L1096
	}
L1096:
	;
	v4194 = v3707
	goto L3
L1097:
	;
	v3728 = int64(-1)
	goto L1099
L1098:
	;
	v3728 = v3725
	goto L1099
L1099:
	;
	v3729 = F_moduleTypeLookupModuleByID(m, v3728)
	mBase = m.M
	v3730 = m.ExcPending
	if v3730 != 0 {
		goto L33
	} else {
		goto L1100
	}
L1100:
	;
	v3731 = int32(0)
	v3732 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[16]))
	if v3732 == v3731 {
		goto L1101
	} else {
		goto L1102
	}
L1101:
	;
	if v3729 != 0 {
		goto L1105
	} else {
		goto L1106
	}
L1102:
	;
	v3736 = v25 + int32(192)
	v3737 = int32(0)
	v3740 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[17]))
	*(*uint8)(unsafe.Add(mBase, uint32(v3736)+9)) = uint8(v3737)
	v3743 = base.I32_wrap_i64(v3728)
	v3746 = int32(63)
	v3749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3740+int32(base.Ui32(v3743)>>(uint(int32(10))%32))&v3746))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3736)+8)) = uint8(v3749)
	v3756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3740+int32(base.Ui32(v3743)>>(uint(int32(16))%32))&v3746))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3736)+7)) = uint8(v3756)
	v3763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3740+int32(base.Ui32(v3743)>>(uint(int32(22))%32))&v3746))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3736)+6)) = uint8(v3763)
	v3771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3740+base.I32_wrap_i64(int64(base.Ui64(v3728)>>(uint(int64(28))%64)))&v3746))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3736)+5)) = uint8(v3771)
	v3779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3740+base.I32_wrap_i64(int64(base.Ui64(v3728)>>(uint(int64(34))%64)))&v3746))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3736)+4)) = uint8(v3779)
	v3787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3740+base.I32_wrap_i64(int64(base.Ui64(v3728)>>(uint(int64(40))%64)))&v3746))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3736)+3)) = uint8(v3787)
	v3795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3740+base.I32_wrap_i64(int64(base.Ui64(v3728)>>(uint(int64(46))%64)))&v3746))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3736)+2)) = uint8(v3795)
	v3803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3740+base.I32_wrap_i64(int64(base.Ui64(v3728)>>(uint(int64(52))%64)))&v3746))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3736)+1)) = uint8(v3803)
	v3809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3740+base.I32_wrap_i64(int64(base.Ui64(v3728)>>(uint(int64(58))%64)))))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3736))) = uint8(v3809)
	goto L1103
L1103:
	;
	v3813 = F_rdbLoadCheckModuleValue(m, l1, v25+int32(192))
	mBase = m.M
	v3814 = m.ExcPending
	if v3814 != 0 {
		goto L33
	} else {
		goto L1104
	}
L1104:
	;
	v4194 = v3813
	goto L3
L1105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+196)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v25)+216)) = l3
	v3904 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+192)) = v3904
	*(*int32)(unsafe.Add(mBase, uint32(v25)+220)) = v3904
	*(*int64)(unsafe.Add(mBase, uint32(v25)+204)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+504)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v25)+496)) = int64(-68719476736)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+212)) = v25 + int32(496)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+200)) = v3729
	v3922 = *(*int32)(unsafe.Add(mBase, uint32(v3729)+12))
	v3923 = m.T0[v3922].(func(*base.Module, int32, int32) int32)(m, v25+int32(192), base.I32_wrap_i64(v3728)&int32(1023))
	mBase = m.M
	v3924 = m.ExcPending
	if v3924 != 0 {
		goto L33
	} else {
		goto L1109
	}
L1106:
	;
	v3816 = v25 + int32(192)
	v3817 = int32(0)
	v3820 = *(*int32)(unsafe.Add(mBase, _c_F_rdbLoadObject[17]))
	*(*uint8)(unsafe.Add(mBase, uint32(v3816)+9)) = uint8(v3817)
	v3823 = base.I32_wrap_i64(v3728)
	v3826 = int32(63)
	v3829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3820+int32(base.Ui32(v3823)>>(uint(int32(10))%32))&v3826))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3816)+8)) = uint8(v3829)
	v3836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3820+int32(base.Ui32(v3823)>>(uint(int32(16))%32))&v3826))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3816)+7)) = uint8(v3836)
	v3843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3820+int32(base.Ui32(v3823)>>(uint(int32(22))%32))&v3826))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3816)+6)) = uint8(v3843)
	v3851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3820+base.I32_wrap_i64(int64(base.Ui64(v3728)>>(uint(int64(28))%64)))&v3826))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3816)+5)) = uint8(v3851)
	v3859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3820+base.I32_wrap_i64(int64(base.Ui64(v3728)>>(uint(int64(34))%64)))&v3826))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3816)+4)) = uint8(v3859)
	v3867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3820+base.I32_wrap_i64(int64(base.Ui64(v3728)>>(uint(int64(40))%64)))&v3826))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3816)+3)) = uint8(v3867)
	v3875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3820+base.I32_wrap_i64(int64(base.Ui64(v3728)>>(uint(int64(46))%64)))&v3826))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3816)+2)) = uint8(v3875)
	v3883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3820+base.I32_wrap_i64(int64(base.Ui64(v3728)>>(uint(int64(52))%64)))&v3826))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3816)+1)) = uint8(v3883)
	v3889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3820+base.I32_wrap_i64(int64(base.Ui64(v3728)>>(uint(int64(58))%64)))))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3816))) = uint8(v3889)
	goto L1107
L1107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = v25 + int32(192)
	F_rdbReportError(m, int32(1), int32(2939), int32(_a_F_rdbLoadObject_59), v25+int32(64))
	mBase = m.M
	v3900 = m.ExcPending
	if v3900 != 0 {
		goto L33
	} else {
		goto L1108
	}
L1108:
	;
	v4194 = int32(0)
	goto L3
L1109:
	;
	v3925 = *(*int32)(unsafe.Add(mBase, uint32(v25)+208))
	if v3925 == int32(0) {
		goto L1110
	} else {
		goto L1111
	}
L1110:
	;
	v3936 = F_rdbLoadLenByRef(m, l1, int32(0), v25+int32(528))
	mBase = m.M
	v3937 = m.ExcPending
	if v3937 != 0 {
		goto L33
	} else {
		goto L1114
	}
L1111:
	;
	F_moduleFreeContext(m, v3925)
	mBase = m.M
	v3929 = m.ExcPending
	if v3929 != 0 {
		goto L33
	} else {
		goto L1112
	}
L1112:
	;
	v3930 = *(*int32)(unsafe.Add(mBase, uint32(v25)+208))
	F_valkey_free(m, v3930)
	mBase = m.M
	v3932 = m.ExcPending
	if v3932 != 0 {
		goto L33
	} else {
		goto L1113
	}
L1113:
	;
	goto L1110
L1114:
	;
	v3939 = *(*int64)(unsafe.Add(mBase, uint32(v25)+528))
	if v3936 == int32(-1) {
		goto L1117
	} else {
		goto L1118
	}
L1115:
	;
	if v3923 != 0 {
		goto L1133
	} else {
		goto L1134
	}
L1116:
	;
	v3955 = int32(0)
	if v3729 == v3955 {
		v3964 = v3955
		goto L1126
	} else {
		goto L1127
	}
L1117:
	;
	v3942 = int64(-1)
	goto L1119
L1118:
	;
	v3942 = v3939
	goto L1119
L1119:
	;
	v3943 = int64(1)
	v3944 = v3942 + v3943
	if base.Ui64(v3943) < base.Ui64(v3944) {
		goto L1116
	} else {
		goto L1120
	}
L1120:
	;
	switch base.I32_wrap_i64(v3944) {
	default:
		goto L1121
	case 1:
		goto L1115
	}
L1121:
	;
	if v3923 == int32(0) {
		goto L6
	} else {
		goto L1122
	}
L1122:
	;
	v3950 = F_createModuleObject(m, v3729, v3923)
	mBase = m.M
	v3951 = m.ExcPending
	if v3951 != 0 {
		goto L33
	} else {
		goto L1123
	}
L1123:
	;
	F_decrRefCount(m, v3950)
	mBase = m.M
	v3953 = m.ExcPending
	if v3953 != 0 {
		goto L33
	} else {
		goto L1124
	}
L1124:
	;
	v4194 = int32(0)
	goto L3
L1125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+80)) = v3964
	F_rdbReportError(m, int32(1), int32(2966), int32(_a_F_rdbLoadObject_60), v25+int32(80))
	mBase = m.M
	v3972 = m.ExcPending
	if v3972 != 0 {
		goto L33
	} else {
		goto L1129
	}
L1126:
	;
	goto L1125
L1127:
	;
	v3959 = *(*int32)(unsafe.Add(mBase, uint32(v3729)+8))
	if v3959 == int32(0) {
		v3964 = v3955
		goto L1126
	} else {
		goto L1128
	}
L1128:
	;
	v3962 = *(*int32)(unsafe.Add(mBase, uint32(v3959)+4))
	v3964 = v3962
	goto L1126
L1129:
	;
	if v3923 == int32(0) {
		goto L6
	} else {
		goto L1130
	}
L1130:
	;
	v3975 = F_createModuleObject(m, v3729, v3923)
	mBase = m.M
	v3976 = m.ExcPending
	if v3976 != 0 {
		goto L33
	} else {
		goto L1131
	}
L1131:
	;
	F_decrRefCount(m, v3975)
	mBase = m.M
	v3978 = m.ExcPending
	if v3978 != 0 {
		goto L33
	} else {
		goto L1132
	}
L1132:
	;
	v4194 = int32(0)
	goto L3
L1133:
	;
	v3999 = F_createModuleObject(m, v3729, v3923)
	mBase = m.M
	v4000 = m.ExcPending
	if v4000 != 0 {
		goto L33
	} else {
		goto L1140
	}
L1134:
	;
	v3980 = int32(0)
	if v3729 == v3980 {
		v3989 = v3980
		goto L1136
	} else {
		goto L1137
	}
L1135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+96)) = v3989
	F_rdbReportError(m, int32(1), int32(2977), int32(_a_F_rdbLoadObject_61), v25+int32(96))
	mBase = m.M
	v3997 = m.ExcPending
	if v3997 != 0 {
		goto L33
	} else {
		goto L1139
	}
L1136:
	;
	goto L1135
L1137:
	;
	v3984 = *(*int32)(unsafe.Add(mBase, uint32(v3729)+8))
	if v3984 == int32(0) {
		v3989 = v3980
		goto L1136
	} else {
		goto L1138
	}
L1138:
	;
	v3987 = *(*int32)(unsafe.Add(mBase, uint32(v3984)+4))
	v3989 = v3987
	goto L1136
L1139:
	;
	v4194 = int32(0)
	goto L3
L1140:
	;
	v4010 = v3999
	goto L15
L1141:
	;
	v4194 = v4010
	goto L3
L1142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L1141
L1143:
	;
	v4194 = v4027
	goto L3
L1144:
	;
	F_streamFreeNACK(m, v3093)
	mBase = m.M
	v4037 = m.ExcPending
	if v4037 != 0 {
		goto L33
	} else {
		goto L1145
	}
L1145:
	;
	v4194 = int32(0)
	goto L3
L1146:
	;
	goto L11
L1147:
	;
	F_valkey_free(m, v2548)
	mBase = m.M
	v4045 = m.ExcPending
	if v4045 != 0 {
		goto L33
	} else {
		goto L1148
	}
L1148:
	;
	v4194 = int32(0)
	goto L3
L1149:
	;
	goto L6
L1150:
	;
	v4074 = int32(0)
	if l4 == v4074 {
		v4194 = v4074
		goto L3
	} else {
		goto L1151
	}
L1151:
	;
	if v4071 != int32(2) {
		v4194 = v4074
		goto L3
	} else {
		goto L1152
	}
L1152:
	;
	goto L7
L1153:
	;
	if l4 == int32(0) {
		goto L6
	} else {
		goto L1154
	}
L1154:
	;
	goto L7
L1155:
	;
	F_sdsfree(m, v497)
	mBase = m.M
	v4178 = m.ExcPending
	if v4178 != 0 {
		goto L33
	} else {
		goto L1156
	}
L1156:
	;
	v4194 = int32(0)
	goto L3
L1157:
	;
	F_sdsfree(m, v229)
	mBase = m.M
	v4185 = m.ExcPending
	if v4185 != 0 {
		goto L33
	} else {
		goto L1158
	}
L1158:
	;
	v4194 = v4180
	goto L3
}
