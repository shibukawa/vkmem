package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_raxCompressNode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(4)) <= base.Ui32(v11) {
		F__serverAssert(m, int32(_a1806), int32(_a1807), int32(382))
		mBase = m.M
		v152 = m.ExcPending
		if v152 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v15 = F_valkey_malloc(m, int32(4))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 != 0 {
				v23 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v23
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v15
				v30 = (v23 - l2) & int32(3)
				v31 = l2 + v30
				v33 = v31 + int32(8)
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v34&int32(1) == v23 {
					v66 = v33
					v67 = v23
				} else {
					v43 = v34 & int32(2)
					if v43 != 0 {
						v62 = int32(0)
					} else {
						v44 = int32(3)
						v45 = int32(base.Ui32(v34) >> (uint(v44) % 32))
						v52 = int32(4)
						if v34&v52 != 0 {
							v57 = v52
						} else {
							v57 = v45 << (uint(int32(2)) % 32)
						}
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l0+v45+(int32(0)-v45)&v44+v57+int32(4))))
						v62 = v61
					}
					if v43 != 0 {
						v65 = v33
					} else {
						v65 = v31 + int32(12)
					}
					v66 = v65
					v67 = v62
				}
				v70 = F_valkey_realloc(m, l0, v66)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					if v70 != 0 {
						v77 = int32(3)
						v78 = l2 << (uint(v77) % 32)
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
						v83 = int32(4)
						v84 = v78 | v79&v77 | v83
						*(*int32)(unsafe.Add(mBase, uint32(v70))) = v84
						if l2 == int32(0) {
						} else {
							v90 = F__emscripten_memcpy_bulkmem(m, v70+v83, l1, l2)
							mBase = m.M
						}
						if v79&int32(1) == int32(0) {
							v111 = v84
						} else {
							if v67 == int32(0) {
								v109 = v78 | int32(7)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v70+l2&int32(536870911)+v30+int32(8)))) = v67
								v109 = v78 | int32(5)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v70))) = v109
							v111 = v109
						}
						v112 = int32(3)
						v113 = int32(base.Ui32(v111) >> (uint(v112) % 32))
						v120 = int32(4)
						if v111&v120 != 0 {
							v125 = v120
						} else {
							v125 = v113 << (uint(int32(2)) % 32)
						}
						v127 = int32(1)
						v128 = v111 << (uint(v127) % 32)
						v134 = int32(0) - v111&v127
						v145 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						*(*int32)(unsafe.Add(mBase, uint32(v70+v113+(int32(0)-v113)&v112+v125+(v128^int32(-1))&v134&int32(4)+v134&(v128|int32(-5)+v127)))) = v145
						return v70
					} else {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						F_valkey_free(m, v72)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				}
			} else {
				v19 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v19
				return v19
			}
		}
	}
}
func F_raxGenericInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v348 int32
	_ = v348
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v382 int32
	_ = v382
	var v384 int64
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v682 int32
	_ = v682
	var v685 int64
	_ = v685
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v844 int64
	_ = v844
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v930 int32
	_ = v930
	var v935 int32
	_ = v935
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v985 int64
	_ = v985
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1014 int32
	_ = v1014
	var v1021 int32
	_ = v1021
	var v1035 int32
	_ = v1035
	var v1048 int64
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1072 int32
	_ = v1072
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1086 int32
	_ = v1086
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1178 int32
	_ = v1178
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1192 int32
	_ = v1192
	var v1205 int32
	_ = v1205
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1222 int64
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1246 int32
	_ = v1246
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1270 int32
	_ = v1270
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1285 int32
	_ = v1285
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1315 int64
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1331 int32
	_ = v1331
	var v1336 int32
	_ = v1336
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1372 int32
	_ = v1372
	var v1376 int32
	_ = v1376
	var v1383 int32
	_ = v1383
	var v1389 int64
	_ = v1389
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1435 int32
	_ = v1435
	var v1450 int32
	_ = v1450
	v7 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if l2 == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v206 = v193 & int32(4)
	if v194 != l2 {
		goto L40
	} else {
		goto L41
	}
L2:
	;
	v192 = v25
	v193 = v26
	v194 = v7
	v195 = l0
	v199 = int32(0)
	goto L1
L3:
	;
	if base.Ui32(v26) < base.Ui32(int32(8)) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v39 = v25
	v40 = v26
	v41 = v7
	v42 = l0
	goto L5
L5:
	;
	v53 = int32(base.Ui32(v40) >> (uint(int32(3)) % 32))
	v54 = int32(4)
	v55 = v39 + v54
	if v40&v54 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v169 = int32(0)
	v178 = v55 + v53 + (v169-v53)&int32(3) + v162<<(uint(int32(2))%32)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	if base.Ui32(v180) < base.Ui32(int32(8)) {
		v192 = v179
		v193 = v180
		v194 = v158
		v195 = v178
		v199 = v169
		goto L1
	} else {
		goto L25
	}
L8:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v41))))
	v133 = int32(0)
	goto L19
L9:
	;
	v60 = int32(0)
	if base.Ui32(l2) <= base.Ui32(v41) {
		v103 = v41
		v108 = v60
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if v108 == v53 {
		v158 = v103
		v162 = v60
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v72 = v41
	v77 = v60
	goto L12
L12:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v77))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v72))))
	if v84 != v86 {
		v103 = v72
		v108 = v77
		goto L10
	} else {
		goto L14
	}
L13:
	;
	v103 = v89
	v108 = v91
	goto L10
L14:
	;
	v88 = int32(1)
	v89 = v72 + v88
	v91 = v77 + v88
	if base.Ui32(v53) <= base.Ui32(v91) {
		v103 = v89
		v108 = v91
		goto L10
	} else {
		goto L15
	}
L15:
	;
	if base.Ui32(v89) < base.Ui32(l2) {
		v72 = v89
		v77 = v91
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v192 = v39
	v193 = v40
	v194 = v103
	v195 = v42
	v199 = v108
	goto L1
L18:
	;
	if v133 != v53 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v133))))
	if v141 == v117&int32(255) {
		goto L18
	} else {
		goto L21
	}
L20:
	;
	v192 = v39
	v193 = v40
	v194 = v41
	v195 = v42
	v199 = v53
	goto L1
L21:
	;
	v144 = v133 + int32(1)
	if v144 != v53 {
		v133 = v144
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v158 = v41 + int32(1)
	v162 = v133
	goto L7
L24:
	;
	v192 = v39
	v193 = v40
	v194 = v41
	v195 = v42
	v199 = v53
	goto L1
L25:
	;
	if base.Ui32(v158) < base.Ui32(l2) {
		v39 = v179
		v40 = v180
		v41 = v158
		v42 = v178
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v192 = v179
	v193 = v180
	v194 = v158
	v195 = v178
	v199 = v169
	goto L1
L27:
	;
	F__serverAssert(m, int32(_a1808), int32(_a1807), int32(887))
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L56
	} else {
		goto L255
	}
L28:
	;
	m.G0 = v23 + int32(16)
	return v1435
L29:
	;
	v1435 = int32(0)
	goto L28
L30:
	;
	if base.Ui32(l2) <= base.Ui32(v194) {
		v1257 = v194
		v1258 = v1113
		v1261 = v1116
		goto L211
	} else {
		goto L212
	}
L31:
	;
	if v436 != 0 {
		goto L199
	} else {
		goto L200
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v430
	v1014 = v195
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430))) = v1007
	goto L32
L34:
	;
	v1007 = v1002 | int32(3)
	goto L33
L35:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v730+int32(-8))))
	goto L192
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v730))) = v935
	goto L35
L37:
	;
	v935 = v930 | int32(3)
	goto L36
L38:
	;
	goto L191
L39:
	;
	if l3 != 0 {
		goto L144
	} else {
		goto L145
	}
L40:
	;
	if v206 != 0 {
		goto L84
	} else {
		goto L85
	}
L41:
	;
	if v206 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v193&int32(1) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	if v199 != 0 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	if v283&int32(1) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L46:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v192+int32(-8))))
	goto L50
L47:
	;
	if l5 == int32(0) {
		v282 = v192
		v283 = v193
		goto L45
	} else {
		goto L48
	}
L48:
	;
	if v193&int32(2) == int32(0) {
		v282 = v192
		v283 = v193
		goto L45
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	if l3 == int32(0) {
		v265 = v192
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v265
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v265+int32(-8))))
	goto L59
L52:
	;
	v229 = int32(0)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v231 = int32(3)
	v232 = int32(base.Ui32(v230) >> (uint(v231) % 32))
	v238 = int32(4)
	if v230&v238 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v243 = v238
	goto L55
L54:
	;
	v243 = v232 << (uint(int32(2)) % 32)
	goto L55
L55:
	;
	v247 = int32(1)
	v259 = F_valkey_realloc(m, v192, v232+(v229-v232)&v231+v243+(v230^int32(-1))<<(uint(v247)%32)&(int32(0)-v230&v247)&int32(4)+int32(8))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	return int32(0)
L57:
	;
	if v259 == int32(0) {
		goto L38
	} else {
		goto L58
	}
L58:
	;
	v265 = v259
	goto L51
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v270 - (v222&int32(2147483647) + int32(8)) + (v274&int32(2147483647) + int32(8))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	v282 = v265
	v283 = v281
	goto L45
L60:
	;
	if l3 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L61:
	;
	if l4 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if l5 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L63:
	;
	if v283&int32(2) != 0 {
		v315 = int32(0)
		goto L64
	} else {
		goto L65
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v315
	goto L62
L65:
	;
	v297 = int32(3)
	v298 = int32(base.Ui32(v283) >> (uint(v297) % 32))
	v305 = int32(4)
	if v283&v305 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v310 = v305
	goto L68
L67:
	;
	v310 = v298 << (uint(int32(2)) % 32)
	goto L68
L68:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v282+v298+(int32(0)-v298)&v297+v310+int32(4))))
	v315 = v314
	goto L64
L69:
	;
	goto L77
L70:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	if l3 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = v348
	goto L69
L72:
	;
	v348 = v320 | int32(3)
	goto L71
L73:
	;
	v323 = int32(3)
	v324 = int32(base.Ui32(v320) >> (uint(v323) % 32))
	v331 = int32(4)
	if v320&v331 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v336 = v331
	goto L76
L75:
	;
	v336 = v324 << (uint(int32(2)) % 32)
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282+v324+(int32(0)-v324)&v323+v336+int32(4)))) = l3
	v348 = v320&int32(-4) | int32(1)
	goto L71
L77:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(0)
	v1435 = int32(0)
	goto L28
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = v382
	v384 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v384 + int64(1)
	v1435 = int32(1)
	goto L28
L79:
	;
	v382 = v283 | int32(3)
	goto L78
L80:
	;
	v358 = int32(3)
	v359 = int32(base.Ui32(v283) >> (uint(v358) % 32))
	v366 = int32(4)
	if v283&v366 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v371 = v366
	goto L83
L82:
	;
	v371 = v359 << (uint(int32(2)) % 32)
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282+v359+(int32(0)-v359)&v358+v371+int32(4)))) = l3
	v382 = v283&int32(-4) | int32(1)
	goto L78
L84:
	;
	v389 = int32(0)
	v390 = int32(3)
	v391 = int32(base.Ui32(v193) >> (uint(v390) % 32))
	v398 = int32(1)
	v399 = v193 << (uint(v398) % 32)
	v400 = int32(-1)
	v405 = v389 - v193&v398
	v407 = int32(4)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v192+v391+(v389-v391)&v390+(v399^v400)&v405&v407+v405&(v399|int32(-5)+v398)+v407)))
	v421 = int32(12)
	if v199 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v1113 = v195
	v1116 = v192
	goto L30
L86:
	;
	v436 = v391 + (v199 ^ v400)
	if v199 == int32(0) {
		v457 = v389
		goto L95
	} else {
		goto L96
	}
L87:
	;
	v424 = v421
	goto L89
L88:
	;
	v424 = int32(16)
	goto L89
L89:
	;
	if v193&int32(3) != int32(1) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v429 = v421
	goto L92
L91:
	;
	v429 = v424
	goto L92
L92:
	;
	v430 = F_valkey_malloc(m, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L56
	} else {
		goto L93
	}
L93:
	;
	if v430 == int32(0) {
		goto L86
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430))) = int32(8)
	goto L86
L95:
	;
	v458 = int32(0)
	if v436 == v458 {
		v470 = v458
		goto L101
	} else {
		goto L102
	}
L96:
	;
	v441 = int32(3)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	if v446&v441 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v451 = int32(8)
	goto L99
L98:
	;
	v451 = int32(4)
	goto L99
L99:
	;
	v455 = F_valkey_malloc(m, v199+(int32(0)-v199)&v441+v451+int32(4))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L56
	} else {
		goto L100
	}
L100:
	;
	v457 = v455
	goto L95
L101:
	;
	if v430 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v468 = F_valkey_malloc(m, v436+(int32(0)-v436)&int32(3)+int32(8))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L56
	} else {
		goto L103
	}
L103:
	;
	v470 = v468
	goto L101
L104:
	;
	v495 = v192 + int32(4)
	v496 = v495 + v199
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+4)) = uint8(v497)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v430+int32(-8))))
	goto L113
L105:
	;
	F_valkey_free(m, v430)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L56
	} else {
		goto L109
	}
L106:
	;
	v473 = int32(0)
	if base.B2i32(v199 != v473)&base.B2i32(v457 == v473) != 0 {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v478 = int32(0)
	if base.B2i32(v436 != v478)&base.B2i32(v470 == v478) == v478 {
		goto L104
	} else {
		goto L108
	}
L108:
	;
	goto L105
L109:
	;
	F_valkey_free(m, v457)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L56
	} else {
		goto L110
	}
L110:
	;
	F_valkey_free(m, v470)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L56
	} else {
		goto L111
	}
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(48)
	goto L29
L113:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v501&int32(2147483647) + int32(8) + v506
	if v199 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v566 = v199 << (uint(int32(3)) % 32)
	v567 = v562&int32(7) | v566
	*(*int32)(unsafe.Add(mBase, uint32(v457))) = v567
	if v199 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L115:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	if v509&int32(1) == int32(0) {
		goto L32
	} else {
		goto L116
	}
L116:
	;
	if v509&int32(2) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	v520 = int32(3)
	v521 = int32(base.Ui32(v509) >> (uint(v520) % 32))
	v528 = int32(4)
	if v509&v528 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	v1002 = v518
	goto L34
L119:
	;
	v533 = v528
	goto L121
L120:
	;
	v533 = v521 << (uint(int32(2)) % 32)
	goto L121
L121:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v192+v521+(int32(0)-v521)&v520+v533+int32(4))))
	if v537 == int32(0) {
		v1002 = v519
		goto L34
	} else {
		goto L122
	}
L122:
	;
	v540 = int32(3)
	v541 = int32(base.Ui32(v519) >> (uint(v540) % 32))
	v548 = int32(4)
	if v519&v548 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v553 = v548
	goto L125
L124:
	;
	v553 = v541 << (uint(int32(2)) % 32)
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430+v541+(int32(0)-v541)&v540+v553+int32(4)))) = v537
	v1007 = v519&int32(-4) | int32(1)
	goto L33
L126:
	;
	v577 = int32(1)
	v579 = int32(2)
	v580 = base.B2i32(v577 < v199) << (uint(v579) % 32)
	v581 = v567&int32(-5) | v580
	*(*int32)(unsafe.Add(mBase, uint32(v457))) = v581
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v588 = v581&int32(-2) | v585&v577
	*(*int32)(unsafe.Add(mBase, uint32(v457))) = v588
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v595 = v588&int32(-3) | v592&v579
	*(*int32)(unsafe.Add(mBase, uint32(v457))) = v595
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	if v597&int32(3) != v577 {
		v648 = v595
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L126
L128:
	;
	v573 = F__emscripten_memcpy_bulkmem(m, v457+int32(4), v495, v199)
	mBase = m.M
	goto L127
L129:
	;
	v650 = int32(3)
	v651 = int32(base.Ui32(v648) >> (uint(v650) % 32))
	v658 = int32(4)
	if v648&v658 != 0 {
		goto L140
	} else {
		goto L141
	}
L130:
	;
	v602 = int32(3)
	v603 = int32(base.Ui32(v597) >> (uint(v602) % 32))
	v610 = int32(4)
	if v597&v610 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v457))) = v645
	v648 = v645
	goto L129
L132:
	;
	v645 = v581 | int32(3)
	goto L131
L133:
	;
	v615 = v610
	goto L135
L134:
	;
	v615 = v603 << (uint(int32(2)) % 32)
	goto L135
L135:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v192+v603+(int32(0)-v603)&v602+v615+int32(4))))
	if v619 == int32(0) {
		goto L132
	} else {
		goto L136
	}
L136:
	;
	v623 = v199 & int32(536870911)
	v630 = int32(2)
	if v199 < v630 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v635 = v623 << (uint(v630) % 32)
	goto L139
L138:
	;
	v635 = int32(4)
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v457+v623+(int32(0)-v199)&int32(3)+v635+int32(4)))) = v619
	v645 = v566 | v580 | int32(1)
	goto L131
L140:
	;
	v663 = v658
	goto L142
L141:
	;
	v663 = v651 << (uint(int32(2)) % 32)
	goto L142
L142:
	;
	v665 = int32(1)
	v666 = v648 << (uint(v665) % 32)
	v672 = int32(0) - v648&v665
	v682 = v457 + v651 + (int32(0)-v651)&v650 + v663 + (v666^int32(-1))&v672&int32(4) + v672&(v666|int32(-5)+v665)
	*(*int32)(unsafe.Add(mBase, uint32(v682))) = v430
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v457
	v685 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v685 + int64(1)
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v457+int32(-8))))
	goto L143
L143:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v691&int32(2147483647) + int32(8) + v696
	v1014 = v682
	goto L31
L144:
	;
	v702 = int32(12)
	goto L146
L145:
	;
	v702 = int32(8)
	goto L146
L146:
	;
	v703 = int32(3)
	v705 = int32(base.Ui32(v193)>>(uint(v703)%32)) - v199
	v710 = (int32(0) - v705) & v703
	v712 = F_valkey_malloc(m, v702+v705+v710)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L56
	} else {
		goto L147
	}
L147:
	;
	v716 = int32(3)
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	if v721&v716 == int32(1) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v726 = int32(8)
	goto L150
L149:
	;
	v726 = int32(4)
	goto L150
L150:
	;
	v730 = F_valkey_malloc(m, v199+(int32(0)-v199)&v716+v726+int32(4))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L56
	} else {
		goto L151
	}
L151:
	;
	if v712 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v742 = int32(3)
	v743 = int32(base.Ui32(v741) >> (uint(v742) % 32))
	v750 = int32(4)
	if v741&v750 != 0 {
		goto L159
	} else {
		goto L160
	}
L153:
	;
	F_valkey_free(m, v712)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L56
	} else {
		goto L156
	}
L154:
	;
	if v730 != 0 {
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	F_valkey_free(m, v730)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L56
	} else {
		goto L157
	}
L157:
	;
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(48)
	v1435 = int32(0)
	goto L28
L159:
	;
	v755 = v750
	goto L161
L160:
	;
	v755 = v743 << (uint(int32(2)) % 32)
	goto L161
L161:
	;
	v757 = int32(1)
	v758 = v741 << (uint(v757) % 32)
	v764 = int32(0) - v741&v757
	v766 = int32(4)
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v192+v743+(int32(0)-v743)&v742+v755+(v758^int32(-1))&v764&v766+v764&(v758|int32(-5)+v757))))
	v779 = base.B2i32(base.Ui32(v705) < base.Ui32(int32(2)))
	if base.Ui32(v705) < base.Ui32(int32(2)) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v780 = v757
	goto L164
L163:
	;
	v780 = v766
	goto L164
L164:
	;
	v783 = v780 | v705<<(uint(int32(3))%32)
	v785 = v783 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v712))) = v785
	v787 = int32(4)
	v790 = v192 + v787
	if v705 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	if l3 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L166:
	;
	goto L165
L167:
	;
	v794 = F__emscripten_memcpy_bulkmem(m, v712+v787, v790+v199, v705)
	mBase = m.M
	goto L166
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v712))) = v812
	v816 = int32(3)
	v817 = int32(base.Ui32(v812) >> (uint(v816) % 32))
	v824 = int32(4)
	if v812&v824 != 0 {
		goto L174
	} else {
		goto L175
	}
L169:
	;
	v812 = v783 | int32(3)
	goto L168
L170:
	;
	v799 = v705 & int32(536870911)
	if base.Ui32(v705) < base.Ui32(int32(2)) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v805 = v799 << (uint(int32(2)) % 32)
	goto L173
L172:
	;
	v805 = int32(4)
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v712+v799+v710+v805+int32(4)))) = l3
	v812 = v785
	goto L168
L174:
	;
	v829 = v824
	goto L176
L175:
	;
	v829 = v817 << (uint(int32(2)) % 32)
	goto L176
L176:
	;
	v831 = int32(1)
	v832 = v812 << (uint(v831) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v712+v817+(int32(0)-v817)&v816+v829+(v832^int32(-1))&int32(4)+(v832|int32(-5))+v831))) = v775
	v844 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v844 + int64(1)
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v712+int32(-8))))
	goto L177
L177:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v850&int32(2147483647) + int32(8) + v855
	*(*int32)(unsafe.Add(mBase, uint32(v730))) = base.B2i32(int32(1) < v199)<<(uint(int32(2))%32) | v199<<(uint(int32(3))%32)
	if v199 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v730
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	if v873&int32(1) == int32(0) {
		goto L35
	} else {
		goto L181
	}
L179:
	;
	goto L178
L180:
	;
	v870 = F__emscripten_memcpy_bulkmem(m, v730+int32(4), v790, v199)
	mBase = m.M
	goto L179
L181:
	;
	if v873&int32(2) == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	v884 = int32(3)
	v885 = int32(base.Ui32(v873) >> (uint(v884) % 32))
	v892 = int32(4)
	if v873&v892 != 0 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	v930 = v882
	goto L37
L184:
	;
	v897 = v892
	goto L186
L185:
	;
	v897 = v885 << (uint(int32(2)) % 32)
	goto L186
L186:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v192+v885+(int32(0)-v885)&v884+v897+int32(4))))
	if v901 == int32(0) {
		v930 = v883
		goto L37
	} else {
		goto L187
	}
L187:
	;
	v904 = int32(3)
	v905 = int32(base.Ui32(v883) >> (uint(v904) % 32))
	v912 = int32(4)
	if v883&v912 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v917 = v912
	goto L190
L189:
	;
	v917 = v905 << (uint(int32(2)) % 32)
	goto L190
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v730+v905+(int32(0)-v905)&v904+v917+int32(4)))) = v901
	v935 = v883&int32(-4) | int32(1)
	goto L36
L191:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(48)
	v1435 = v229
	goto L28
L192:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v942&int32(2147483647) + int32(8) + v947
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	v951 = int32(3)
	v952 = int32(base.Ui32(v950) >> (uint(v951) % 32))
	v959 = int32(4)
	if v950&v959 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v964 = v959
	goto L195
L194:
	;
	v964 = v952 << (uint(int32(2)) % 32)
	goto L195
L195:
	;
	v966 = int32(1)
	v967 = v950 << (uint(v966) % 32)
	v973 = int32(0) - v950&v966
	*(*int32)(unsafe.Add(mBase, uint32(v730+v952+(int32(0)-v952)&v951+v964+(v967^int32(-1))&v973&int32(4)+v973&(v967|int32(-5)+v966)))) = v712
	v985 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v985 + int64(1)
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v192+int32(-8))))
	goto L196
L196:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v996 - (v991&int32(2147483647) + int32(8))
	F_valkey_free(m, v192)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L56
	} else {
		goto L197
	}
L197:
	;
	v1435 = int32(1)
	goto L28
L198:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	v1064 = int32(3)
	v1065 = int32(base.Ui32(v1063) >> (uint(v1064) % 32))
	v1072 = int32(4)
	if v1063&v1072 != 0 {
		goto L205
	} else {
		goto L206
	}
L199:
	;
	v1021 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = base.B2i32(v436 != v1021)<<(uint(int32(2))%32) | v436<<(uint(int32(3))%32)
	if v436 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	v1062 = v418
	goto L198
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470+v436&int32(536870911)+(int32(0)-v436)&int32(3)+int32(4)))) = v418
	v1048 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v1048 + int64(1)
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v470+int32(-8))))
	goto L204
L202:
	;
	goto L201
L203:
	;
	v1035 = F__emscripten_memcpy_bulkmem(m, v470+int32(4), v496+v1021, v436)
	mBase = m.M
	goto L202
L204:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v1054&int32(2147483647) + int32(8) + v1059
	v1062 = v470
	goto L198
L205:
	;
	v1077 = v1072
	goto L207
L206:
	;
	v1077 = v1065 << (uint(int32(2)) % 32)
	goto L207
L207:
	;
	v1079 = int32(1)
	v1080 = v1063 << (uint(v1079) % 32)
	v1086 = int32(0) - v1063&v1079
	*(*int32)(unsafe.Add(mBase, uint32(v430+v1065+(int32(0)-v1065)&v1064+v1077+(v1080^int32(-1))&v1086&int32(4)+v1086&(v1080|int32(-5)+v1079)))) = v1062
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v192+int32(-8))))
	goto L208
L208:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v1105 - (v1100&int32(2147483647) + int32(8))
	F_valkey_free(m, v192)
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L56
	} else {
		goto L209
	}
L209:
	;
	v1113 = v1014
	v1116 = v430
	goto L30
L210:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1376)))
	if base.Ui32(int32(7)) < base.Ui32(v1383) {
		goto L250
	} else {
		goto L251
	}
L211:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1261+int32(-8))))
	goto L233
L212:
	;
	v1133 = v194
	v1134 = v1113
	v1137 = v1116
	goto L213
L213:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1137+int32(-8))))
	goto L215
L214:
	;
	v1257 = v1246
	v1258 = v1218
	v1261 = v1236
	goto L211
L215:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1137)))
	if base.Ui32(int32(7)) < base.Ui32(v1151) {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v1222 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v1222 + int64(1)
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1221+int32(-8))))
	goto L230
L217:
	;
	v1205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v1133))))
	v1210 = F_raxAddChild(m, v1137, v1205, v23+int32(12), v23+int32(8))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L56
	} else {
		goto L228
	}
L218:
	;
	v1154 = l2 - v1133
	if base.Ui32(v1154) < base.Ui32(int32(2)) {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v1158 = int32(536870911)
	if base.Ui32(v1154) < base.Ui32(v1158) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v1161 = v1154
	goto L222
L221:
	;
	v1161 = v1158
	goto L222
L222:
	;
	v1164 = F_raxCompressNode(m, v1137, l1+v1133, v1161, v23+int32(12))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L56
	} else {
		goto L223
	}
L223:
	;
	if v1164 == int32(0) {
		v1372 = v1133
		v1376 = v1137
		goto L210
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1134))) = v1164
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1164)))
	v1170 = int32(3)
	v1171 = int32(base.Ui32(v1169) >> (uint(v1170) % 32))
	v1178 = int32(4)
	if v1169&v1178 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1183 = v1178
	goto L227
L226:
	;
	v1183 = v1171 << (uint(int32(2)) % 32)
	goto L227
L227:
	;
	v1185 = int32(1)
	v1186 = v1169 << (uint(v1185) % 32)
	v1192 = int32(0) - v1169&v1185
	v1218 = v1164 + v1171 + (int32(0)-v1171)&v1170 + v1183 + (v1186^int32(-1))&v1192&int32(4) + v1192&(v1186|int32(-5)+v1185)
	v1219 = v1161
	v1221 = v1164
	goto L216
L228:
	;
	if v1210 == int32(0) {
		v1372 = v1133
		v1376 = v1137
		goto L210
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1134))) = v1210
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1218 = v1215
	v1219 = int32(1)
	v1221 = v1210
	goto L216
L230:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1236+int32(-8))))
	goto L231
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v1226 - (v1146&int32(2147483647) + int32(8)) + (v1230&int32(2147483647) + int32(8)) + (v1239&int32(2147483647) + int32(8))
	v1246 = v1219 + v1133
	if base.Ui32(v1246) < base.Ui32(l2) {
		v1133 = v1246
		v1134 = v1218
		v1137 = v1236
		goto L213
	} else {
		goto L232
	}
L232:
	;
	goto L214
L233:
	;
	if l3 == int32(0) {
		v1309 = v1261
		goto L234
	} else {
		goto L235
	}
L234:
	;
	if v1309 == int32(0) {
		v1372 = v1257
		v1376 = v1261
		goto L210
	} else {
		goto L240
	}
L235:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1261)))
	v1278 = int32(3)
	v1279 = int32(base.Ui32(v1277) >> (uint(v1278) % 32))
	v1285 = int32(4)
	if v1277&v1285 != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1290 = v1285
	goto L238
L237:
	;
	v1290 = v1279 << (uint(int32(2)) % 32)
	goto L238
L238:
	;
	v1294 = int32(1)
	v1306 = F_valkey_realloc(m, v1261, v1279+(int32(0)-v1279)&v1278+v1290+(v1277^int32(-1))<<(uint(v1294)%32)&(int32(0)-v1277&v1294)&int32(4)+int32(8))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L56
	} else {
		goto L239
	}
L239:
	;
	v1309 = v1306
	goto L234
L240:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1309)))
	if v1312&int32(1) != 0 {
		v1320 = v1312
		goto L241
	} else {
		goto L242
	}
L241:
	;
	if l3 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	v1315 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v1315 + int64(1)
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1309)))
	v1320 = v1319
	goto L241
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1309))) = v1348
	*(*int32)(unsafe.Add(mBase, uint32(v1258))) = v1309
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1309+int32(-8))))
	goto L249
L244:
	;
	v1348 = v1320 | int32(3)
	goto L243
L245:
	;
	v1323 = int32(3)
	v1324 = int32(base.Ui32(v1320) >> (uint(v1323) % 32))
	v1331 = int32(4)
	if v1320&v1331 != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1336 = v1331
	goto L248
L247:
	;
	v1336 = v1324 << (uint(int32(2)) % 32)
	goto L248
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1309+v1324+(int32(0)-v1324)&v1323+v1336+int32(4)))) = l3
	v1348 = v1320&int32(-4) | int32(1)
	goto L243
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v1351 - (v1270&int32(2147483647) + int32(8)) + (v1355&int32(2147483647) + int32(8))
	v1435 = int32(1)
	goto L28
L250:
	;
	goto L254
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1376))) = v1383 | int32(3)
	v1389 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v1389 + int64(1)
	v1394 = F_raxRemove(m, l0, l1, v1372, int32(0))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L56
	} else {
		goto L252
	}
L252:
	;
	if v1394 == int32(0) {
		goto L27
	} else {
		goto L253
	}
L253:
	;
	goto L250
L254:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(48)
	goto L29
L255:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_raxGetData(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5&int32(2) != 0 {
		v29 = int32(0)
	} else {
		v8 = int32(3)
		v9 = int32(base.Ui32(v5) >> (uint(v8) % 32))
		v16 = int32(4)
		if v5&v16 != 0 {
			v21 = v16
		} else {
			v21 = v9 << (uint(int32(2)) % 32)
		}
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0+v9+(int32(0)-v9)&v8+v21+v5<<(uint(int32(2))%32)&int32(4))))
		v29 = v28
	}
	return v29
}
func F_raxNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	v5 = F_raxIteratorNextStep(m, l0, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v11 = int32(0)
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v12&int32(2) == v11 {
				v23 = int32(1)
			} else {
				v17 = v11
				*(*int32)(unsafe.Add(mBase, _consts[18])) = v17
				v23 = int32(0)
			}
		} else {
			v17 = int32(48)
			*(*int32)(unsafe.Add(mBase, _consts[18])) = v17
			v23 = int32(0)
		}
		return v23
	}
}
func F_raxPrev(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	v5 = F_raxIteratorPrevStep(m, l0, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v11 = int32(0)
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v12&int32(2) == v11 {
				v23 = int32(1)
			} else {
				v17 = v11
				*(*int32)(unsafe.Add(mBase, _consts[18])) = v17
				v23 = int32(0)
			}
		} else {
			v17 = int32(48)
			*(*int32)(unsafe.Add(mBase, _consts[18])) = v17
			v23 = int32(0)
		}
		return v23
	}
}
func F_raxRecursiveFree(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = int32(base.Ui32(v9) >> (uint(int32(3)) % 32))
	v13 = v9 & int32(4)
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l2 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v14 = int32(1)
	goto L4
L3:
	;
	v14 = v11
	goto L4
L4:
	;
	if v14 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v13 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v26 = int32(4)
	goto L8
L7:
	;
	v26 = v11 << (uint(int32(2)) % 32)
	goto L8
L8:
	;
	v28 = int32(1)
	v29 = v9 << (uint(v28) % 32)
	v35 = int32(0) - v9&v28
	v49 = l1 + v11 + (int32(0)-v11)&int32(3) + v26 + (v29^int32(-1))&v35&int32(4) + v35&(v29|int32(-5)+v28)
	v52 = v14
	goto L9
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	F_raxRecursiveFree(m, l0, v53, l2)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L1
L11:
	;
	return
L12:
	;
	v59 = v52 + int32(-1)
	if v59 != 0 {
		v49 = v49 + int32(-4)
		v52 = v59
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	F_valkey_free(m, l1)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L11
	} else {
		goto L21
	}
L15:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v69&int32(3) != int32(1) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v74 = int32(3)
	v75 = int32(base.Ui32(v69) >> (uint(v74) % 32))
	v82 = int32(4)
	if v69&v82 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v87 = v82
	goto L19
L18:
	;
	v87 = v75 << (uint(int32(2)) % 32)
	goto L19
L19:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1+v75+(int32(0)-v75)&v74+v87+int32(4))))
	m.T0[l2].(func(*base.Module, int32))(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	v98 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v98 + int64(-1)
	return
}
func F_raxRemove(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v241 int32
	_ = v241
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var __phi282 int32
	_ = __phi282
	var v283 int32
	_ = v283
	var __phi283 int32
	_ = __phi283
	var v286 int32
	_ = v286
	var __phi286 int32
	_ = __phi286
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v398 int32
	_ = v398
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v429 int64
	_ = v429
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int64
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v535 int32
	_ = v535
	var v546 int32
	_ = v546
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v600 int32
	_ = v600
	var v615 int32
	_ = v615
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(160)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+152)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = int64(137438953472)
	v24 = v16 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v5
	v35 = F_raxLowWalk(m, l0, l1, l2, v16+int32(156), v5, v16+int32(8), v16+int32(12))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v16 + int32(160)
	return v655
L2:
	;
	F_valkey_free(m, v635)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L5
	} else {
		goto L110
	}
L3:
	;
	if l3 == int32(0) {
		v84 = v41
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v56 != v24 {
		v635 = v56
		v640 = v5
		goto L2
	} else {
		goto L10
	}
L5:
	;
	return int32(0)
L6:
	;
	if v35 != l2 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v16)+156))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v41&int32(1) == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v49 = int32(0)
	if int32(base.Ui32(v41)>>(uint(int32(2))%32))&base.B2i32(v48 != v49) == v49 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L4
L10:
	;
	v655 = v5
	goto L1
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v84 & int32(-2)
	v89 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v89 + int64(-1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v16)+156))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if base.Ui32(int32(7)) < base.Ui32(v94) {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	if v41&int32(2) != 0 {
		v81 = int32(0)
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v84 = v83
	goto L11
L14:
	;
	v63 = int32(3)
	v64 = int32(base.Ui32(v41) >> (uint(v63) % 32))
	v71 = int32(4)
	if v41&v71 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v76 = v71
	goto L17
L16:
	;
	v76 = v64 << (uint(int32(2)) % 32)
	goto L17
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v40+v64+(int32(0)-v64)&v63+v76+int32(4))))
	v81 = v80
	goto L13
L18:
	;
	v632 = int32(1)
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v633 == v24 {
		v655 = v632
		goto L1
	} else {
		goto L109
	}
L19:
	;
	if v257 != 0 {
		goto L18
	} else {
		goto L48
	}
L20:
	;
	v254 = v93
	v257 = base.B2i32(v94&int32(-8) != int32(8))
	goto L19
L21:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v93 == v97 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v101 = v93
	goto L23
L23:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v101+int32(-8))))
	goto L25
L24:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v138+int32(-8))))
	goto L36
L25:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v119 - (v114&int32(2147483647) + int32(8))
	F_valkey_free(m, v101)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v124 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v124 + int64(-1)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v128 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+156)) = v138
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	if v140&int32(1) != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v131 = v128 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133+v131<<(uint(int32(2))%32))))
	v138 = v137
	goto L27
L29:
	;
	v138 = int32(0)
	goto L27
L30:
	;
	goto L24
L31:
	;
	if v140&int32(4) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v138 != v149 {
		v101 = v138
		goto L23
	} else {
		goto L35
	}
L33:
	;
	if v140&int32(-8) != int32(8) {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	goto L30
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v16)+156))
	v159 = F_raxRemoveChild(m, v158, v101)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v159+int32(-8))))
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v161 - (v153&int32(2147483647) + int32(8)) + (v165&int32(2147483647) + int32(8))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v16)+156))
	if v159 == v172 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v241&int32(-7) != int32(8) {
		goto L18
	} else {
		goto L47
	}
L40:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v174 == int32(0) {
		v219 = l0
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v159
	goto L39
L42:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v177+v174<<(uint(int32(2))%32)+int32(-4))))
	if v183 == int32(0) {
		v219 = l0
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v187 = int32(3)
	v188 = int32(base.Ui32(v186) >> (uint(v187) % 32))
	v198 = v183 + v188 + (int32(0)-v188)&v187 + int32(4)
	goto L44
L44:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	if v212 != v172 {
		v198 = v198 + int32(4)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v219 = v198
	goto L41
L46:
	;
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+156)) = v159
	v254 = v159
	v257 = int32(0)
	goto L19
L48:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v16)+152))
	if v265 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	v266 = int32(0)
	v267 = int32(1)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v268 == v266 {
		v322 = v254
		v327 = v266
		v328 = v267
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	if base.Ui32(v333) < base.Ui32(int32(8)) {
		goto L18
	} else {
		goto L63
	}
L51:
	;
	v272 = v268 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v272
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v274+v272<<(uint(int32(2))%32))))
	if v278 == int32(0) {
		v322 = v254
		v327 = v266
		v328 = v267
		goto L50
	} else {
		goto L52
	}
L52:
	;
	__phi282 = v278
	__phi283 = v254
	__phi286 = v272
	v282 = __phi282
	v283 = __phi283
	v286 = __phi286
	goto L53
L53:
	;
	v294 = int32(0)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	if v295&int32(1) == v294 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v322 = v282
	v327 = v307
	v328 = v308
	goto L50
L55:
	;
	if v295&int32(4) != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v322 = v283
	v327 = v282
	v328 = v294
	goto L50
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+156)) = v282
	v307 = int32(0)
	v308 = int32(1)
	if v286 == v307 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	if v295&int32(-8) == int32(8) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v322 = v283
	v327 = v282
	v328 = v294
	goto L50
L60:
	;
	goto L54
L61:
	;
	v312 = v286 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v312
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v274+v312<<(uint(int32(2))%32))))
	if v317 == int32(0) {
		v322 = v282
		v327 = v307
		v328 = v308
		goto L50
	} else {
		goto L62
	}
L62:
	;
	__phi282 = v317
	__phi283 = v282
	__phi286 = v312
	v282 = __phi282
	v283 = __phi283
	v286 = __phi286
	goto L53
L63:
	;
	v340 = v322
	v344 = v333
	v348 = int32(base.Ui32(v333) >> (uint(int32(3)) % 32))
	v349 = int32(1)
	goto L66
L64:
	;
	v411 = int32(0)
	v419 = F_valkey_malloc(m, v408+(v411-v408)&int32(3)+int32(8))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L5
	} else {
		goto L79
	}
L65:
	;
	if base.Ui32(v349) < base.Ui32(int32(2)) {
		goto L18
	} else {
		goto L77
	}
L66:
	;
	v352 = int32(3)
	v353 = int32(base.Ui32(v344) >> (uint(v352) % 32))
	v360 = int32(4)
	if v344&v360 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v365 = v360
	goto L70
L69:
	;
	v365 = v353 << (uint(int32(2)) % 32)
	goto L70
L70:
	;
	v367 = int32(1)
	v368 = v344 << (uint(v367) % 32)
	v374 = int32(0) - v344&v367
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v340+v353+(int32(0)-v353)&v352+v365+(v368^int32(-1))&v374&int32(4)+v374&(v368|int32(-5)+v367))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+156)) = v385
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	if v387&v367 != 0 {
		goto L65
	} else {
		goto L71
	}
L71:
	;
	if v387&int32(4) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v398 = int32(base.Ui32(v387)>>(uint(int32(3))%32)) + v348
	if base.Ui32(int32(536870911)) < base.Ui32(v398) {
		goto L65
	} else {
		goto L75
	}
L73:
	;
	if v387&int32(-8) != int32(8) {
		goto L65
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	if base.Ui32(int32(8)) <= base.Ui32(v387) {
		v340 = v385
		v344 = v387
		v348 = v398
		v349 = v349 + int32(1)
		goto L66
	} else {
		goto L76
	}
L76:
	;
	v408 = v398
	goto L64
L77:
	;
	v408 = v348
	goto L64
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v419))) = v408<<(uint(int32(3))%32) | int32(4)
	v429 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v429 + int64(1)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v419+int32(-8))))
	goto L82
L79:
	;
	if v419 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v421 = int32(1)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v422 != v24 {
		v635 = v422
		v640 = v421
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v655 = v421
	goto L1
L82:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v435&int32(2147483647) + int32(8) + v440
	*(*int32)(unsafe.Add(mBase, uint32(v16)+156)) = v322
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	if base.Ui32(v444) < base.Ui32(int32(8)) {
		v546 = v322
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	v555 = int32(3)
	v556 = int32(base.Ui32(v554) >> (uint(v555) % 32))
	v563 = int32(4)
	if v554&v563 != 0 {
		goto L101
	} else {
		goto L102
	}
L84:
	;
	v450 = v444
	v454 = v322
	v458 = v411
	goto L85
L85:
	;
	v466 = int32(base.Ui32(v450) >> (uint(int32(3)) % 32))
	if v466 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v546 = v523
	goto L83
L87:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	v472 = int32(3)
	v473 = int32(base.Ui32(v471) >> (uint(v472) % 32))
	v480 = int32(4)
	if v471&v480 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L87
L89:
	;
	v469 = F__emscripten_memcpy_bulkmem(m, v419+int32(4)+v458, v454+int32(4), v466)
	mBase = m.M
	goto L88
L90:
	;
	v485 = v480
	goto L92
L91:
	;
	v485 = v473 << (uint(int32(2)) % 32)
	goto L92
L92:
	;
	v487 = int32(1)
	v488 = v471 << (uint(v487) % 32)
	v494 = int32(0) - v471&v487
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v454+v473+(int32(0)-v473)&v472+v485+(v488^int32(-1))&v494&int32(4)+v494&(v488|int32(-5)+v487))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+156)) = v505
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v454+int32(-8))))
	goto L93
L93:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v514 - (v509&int32(2147483647) + int32(8))
	F_valkey_free(m, v454)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	v519 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v519 + int64(-1)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v16)+156))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	if v524&int32(1) != 0 {
		v546 = v523
		goto L83
	} else {
		goto L95
	}
L95:
	;
	if v524&int32(4) != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v535 = v473 + v458
	if base.Ui32(int32(536870911)) < base.Ui32(int32(base.Ui32(v524)>>(uint(int32(3))%32))+v535) {
		v546 = v523
		goto L83
	} else {
		goto L99
	}
L97:
	;
	if v524&int32(-8) != int32(8) {
		v546 = v523
		goto L83
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	if base.Ui32(int32(7)) < base.Ui32(v524) {
		v450 = v524
		v454 = v523
		v458 = v535
		goto L85
	} else {
		goto L100
	}
L100:
	;
	goto L86
L101:
	;
	v568 = v563
	goto L103
L102:
	;
	v568 = v556 << (uint(int32(2)) % 32)
	goto L103
L103:
	;
	v570 = int32(1)
	v571 = v554 << (uint(v570) % 32)
	v577 = int32(0) - v554&v570
	*(*int32)(unsafe.Add(mBase, uint32(v419+v556+(int32(0)-v556)&v555+v568+(v571^int32(-1))&v577&int32(4)+v577&(v571|int32(-5)+v570)))) = v546
	if v328 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v419
	goto L18
L105:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	v590 = int32(3)
	v591 = int32(base.Ui32(v589) >> (uint(v590) % 32))
	v600 = v327 + v591 + (int32(0)-v591)&v590 + int32(4)
	goto L106
L106:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v600)))
	if v615 != v322 {
		v600 = v600 + int32(4)
		goto L106
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v600))) = v419
	goto L18
L108:
	;
	goto L107
L109:
	;
	v635 = v633
	v640 = v632
	goto L2
L110:
	;
	v655 = v640
	goto L1
}
func F_raxRemoveChild(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10&int32(4) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return l0
L2:
	;
	v45 = l0 + int32(4)
	v46 = int32(3)
	v47 = int32(base.Ui32(v10) >> (uint(v46) % 32))
	v53 = v45 + v47 + (int32(0)-v47)&v46
	v56 = v53
	v57 = v45
	goto L8
L3:
	;
	if v10&int32(1) == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v10&int32(2) != 0 {
		v41 = int32(3)
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v41
	return l0
L6:
	;
	v22 = int32(3)
	v24 = int32(base.Ui32(v10) >> (uint(v22) % 32))
	v26 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0+v24+(v26-v24)&v22+int32(8))))
	if v33 == v26 {
		v41 = v22
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(4)))) = v33
	v41 = int32(1)
	goto L5
L8:
	;
	v64 = v57 + int32(1)
	v66 = v56 + int32(4)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v67 != l1 {
		v56 = v66
		v57 = v64
		goto L8
	} else {
		goto L10
	}
L9:
	;
	v70 = v45 - v57 + v47
	v72 = v70 + int32(-1)
	if v57 == v64 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L9
L11:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v225 = v223 & int32(24)
	if v225 == int32(8) {
		goto L52
	} else {
		goto L53
	}
L12:
	;
	goto L11
L13:
	;
	v76 = v72 + v57
	if base.Ui32(int32(0)-v72<<(uint(int32(1))%32)) < base.Ui32(v64-v76) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v86 = (v64 ^ v57) & int32(3)
	if base.Ui32(v64) <= base.Ui32(v57) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v83 = F___memcpy(m, v57, v64, v72)
	mBase = m.M
	goto L11
L16:
	;
	if v192 == int32(0) {
		goto L12
	} else {
		goto L48
	}
L17:
	;
	if base.Ui32(v170) <= base.Ui32(int32(3)) {
		v191 = v169
		v192 = v170
		v193 = v171
		goto L16
	} else {
		goto L44
	}
L18:
	;
	if v86 != 0 {
		v152 = v72
		goto L28
	} else {
		goto L29
	}
L19:
	;
	if v86 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v57&int32(3) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v191 = v64
	v192 = v72
	v193 = v57
	goto L16
L22:
	;
	v93 = v64
	v94 = v72
	v95 = v57
	goto L24
L23:
	;
	v169 = v64
	v170 = v72
	v171 = v57
	goto L17
L24:
	;
	if v94 == int32(0) {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v99)
	v101 = int32(1)
	v102 = v93 + v101
	v104 = v94 + int32(-1)
	v106 = v95 + v101
	if v106&int32(3) == int32(0) {
		v169 = v102
		v170 = v104
		v171 = v106
		goto L17
	} else {
		goto L27
	}
L27:
	;
	v93 = v102
	v94 = v104
	v95 = v106
	goto L24
L28:
	;
	if v152 == int32(0) {
		goto L12
	} else {
		goto L40
	}
L29:
	;
	if v76&int32(3) == int32(0) {
		v132 = v72
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if base.Ui32(v132) <= base.Ui32(int32(3)) {
		v152 = v132
		goto L28
	} else {
		goto L36
	}
L31:
	;
	v117 = v72
	goto L32
L32:
	;
	if v117 == int32(0) {
		goto L12
	} else {
		goto L34
	}
L33:
	;
	v132 = v123
	goto L30
L34:
	;
	v123 = v117 + int32(-1)
	v124 = v57 + v123
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v123))))
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v126)
	if v124&int32(3) != 0 {
		v117 = v123
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v139 = v132
	goto L37
L37:
	;
	v143 = v139 + int32(-4)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v64+v143)))
	*(*int32)(unsafe.Add(mBase, uint32(v57+v143))) = v146
	if base.Ui32(int32(3)) < base.Ui32(v143) {
		v139 = v143
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v152 = v143
	goto L28
L39:
	;
	goto L38
L40:
	;
	v159 = v152
	goto L41
L41:
	;
	v163 = v159 + int32(-1)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v163))))
	*(*uint8)(unsafe.Add(mBase, uint32(v57+v163))) = uint8(v166)
	if v163 != 0 {
		v159 = v163
		goto L41
	} else {
		goto L43
	}
L43:
	;
	goto L12
L44:
	;
	v176 = v169
	v177 = v170
	v178 = v171
	goto L45
L45:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v180
	v182 = int32(4)
	v183 = v176 + v182
	v185 = v178 + v182
	v187 = v177 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v187) {
		v176 = v183
		v177 = v187
		v178 = v185
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v191 = v183
	v192 = v187
	v193 = v185
	goto L16
L47:
	;
	goto L46
L48:
	;
	v198 = v191
	v199 = v192
	v200 = v193
	goto L49
L49:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	*(*uint8)(unsafe.Add(mBase, uint32(v200))) = uint8(v202)
	v204 = int32(1)
	v209 = v199 + int32(-1)
	if v209 != 0 {
		v198 = v198 + v204
		v199 = v209
		v200 = v200 + v204
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
	v228 = int32(-4)
	goto L54
L53:
	;
	v228 = int32(0)
	goto L54
L54:
	;
	if v225 != int32(8) {
		v387 = v223
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v388 = v56 + v228
	v391 = int32(1)
	v402 = (v387^int32(-1))<<(uint(v391)%32)&(int32(0)-v387&v391)&int32(4) + v72<<(uint(int32(2))%32)
	if v388 == v66 {
		goto L99
	} else {
		goto L100
	}
L56:
	;
	v232 = v53 + int32(-4)
	v237 = (int32(base.Ui32(v223)>>(uint(int32(3))%32)) - v70) << (uint(int32(2)) % 32)
	if v232 == v53 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v387 = v386
	goto L55
L58:
	;
	goto L57
L59:
	;
	v241 = v237 + v232
	if base.Ui32(int32(0)-v237<<(uint(int32(1))%32)) < base.Ui32(v53-v241) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v251 = (v53 ^ v232) & int32(3)
	if base.Ui32(v53) <= base.Ui32(v232) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v248 = F___memcpy(m, v232, v53, v237)
	mBase = m.M
	goto L57
L62:
	;
	if v357 == int32(0) {
		goto L58
	} else {
		goto L94
	}
L63:
	;
	if base.Ui32(v335) <= base.Ui32(int32(3)) {
		v356 = v334
		v357 = v335
		v358 = v336
		goto L62
	} else {
		goto L90
	}
L64:
	;
	if v251 != 0 {
		v317 = v237
		goto L74
	} else {
		goto L75
	}
L65:
	;
	if v251 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	if v232&int32(3) != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v356 = v53
	v357 = v237
	v358 = v232
	goto L62
L68:
	;
	v258 = v53
	v259 = v237
	v260 = v232
	goto L70
L69:
	;
	v334 = v53
	v335 = v237
	v336 = v232
	goto L63
L70:
	;
	if v259 == int32(0) {
		goto L58
	} else {
		goto L72
	}
L72:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	*(*uint8)(unsafe.Add(mBase, uint32(v260))) = uint8(v264)
	v266 = int32(1)
	v267 = v258 + v266
	v269 = v259 + int32(-1)
	v271 = v260 + v266
	if v271&int32(3) == int32(0) {
		v334 = v267
		v335 = v269
		v336 = v271
		goto L63
	} else {
		goto L73
	}
L73:
	;
	v258 = v267
	v259 = v269
	v260 = v271
	goto L70
L74:
	;
	if v317 == int32(0) {
		goto L58
	} else {
		goto L86
	}
L75:
	;
	if v241&int32(3) == int32(0) {
		v297 = v237
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if base.Ui32(v297) <= base.Ui32(int32(3)) {
		v317 = v297
		goto L74
	} else {
		goto L82
	}
L77:
	;
	v282 = v237
	goto L78
L78:
	;
	if v282 == int32(0) {
		goto L58
	} else {
		goto L80
	}
L79:
	;
	v297 = v288
	goto L76
L80:
	;
	v288 = v282 + int32(-1)
	v289 = v232 + v288
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v288))))
	*(*uint8)(unsafe.Add(mBase, uint32(v289))) = uint8(v291)
	if v289&int32(3) != 0 {
		v282 = v288
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v304 = v297
	goto L83
L83:
	;
	v308 = v304 + int32(-4)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v53+v308)))
	*(*int32)(unsafe.Add(mBase, uint32(v232+v308))) = v311
	if base.Ui32(int32(3)) < base.Ui32(v308) {
		v304 = v308
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v317 = v308
	goto L74
L85:
	;
	goto L84
L86:
	;
	v324 = v317
	goto L87
L87:
	;
	v328 = v324 + int32(-1)
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v328))))
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v328))) = uint8(v331)
	if v328 != 0 {
		v324 = v328
		goto L87
	} else {
		goto L89
	}
L89:
	;
	goto L58
L90:
	;
	v341 = v334
	v342 = v335
	v343 = v336
	goto L91
L91:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	*(*int32)(unsafe.Add(mBase, uint32(v343))) = v345
	v347 = int32(4)
	v348 = v341 + v347
	v350 = v343 + v347
	v352 = v342 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v352) {
		v341 = v348
		v342 = v352
		v343 = v350
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v356 = v348
	v357 = v352
	v358 = v350
	goto L62
L93:
	;
	goto L92
L94:
	;
	v363 = v356
	v364 = v357
	v365 = v358
	goto L95
L95:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
	*(*uint8)(unsafe.Add(mBase, uint32(v365))) = uint8(v367)
	v369 = int32(1)
	v374 = v364 + int32(-1)
	if v374 != 0 {
		v363 = v363 + v369
		v364 = v374
		v365 = v365 + v369
		goto L95
	} else {
		goto L97
	}
L96:
	;
	goto L58
L97:
	;
	goto L96
L98:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v552 = int32(-8)
	v555 = v551&v552 + v552
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v555 | v551&int32(7)
	v560 = int32(3)
	v561 = int32(base.Ui32(v555) >> (uint(v560) % 32))
	v567 = int32(4)
	if v551&v567 != 0 {
		goto L139
	} else {
		goto L140
	}
L99:
	;
	goto L98
L100:
	;
	v406 = v402 + v388
	if base.Ui32(int32(0)-v402<<(uint(int32(1))%32)) < base.Ui32(v66-v406) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v416 = (v66 ^ v388) & int32(3)
	if base.Ui32(v66) <= base.Ui32(v388) {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v413 = F___memcpy(m, v388, v66, v402)
	mBase = m.M
	goto L98
L103:
	;
	if v522 == int32(0) {
		goto L99
	} else {
		goto L135
	}
L104:
	;
	if base.Ui32(v500) <= base.Ui32(int32(3)) {
		v521 = v499
		v522 = v500
		v523 = v501
		goto L103
	} else {
		goto L131
	}
L105:
	;
	if v416 != 0 {
		v482 = v402
		goto L115
	} else {
		goto L116
	}
L106:
	;
	if v416 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	if v388&int32(3) != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v521 = v66
	v522 = v402
	v523 = v388
	goto L103
L109:
	;
	v423 = v66
	v424 = v402
	v425 = v388
	goto L111
L110:
	;
	v499 = v66
	v500 = v402
	v501 = v388
	goto L104
L111:
	;
	if v424 == int32(0) {
		goto L99
	} else {
		goto L113
	}
L113:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423))))
	*(*uint8)(unsafe.Add(mBase, uint32(v425))) = uint8(v429)
	v431 = int32(1)
	v432 = v423 + v431
	v434 = v424 + int32(-1)
	v436 = v425 + v431
	if v436&int32(3) == int32(0) {
		v499 = v432
		v500 = v434
		v501 = v436
		goto L104
	} else {
		goto L114
	}
L114:
	;
	v423 = v432
	v424 = v434
	v425 = v436
	goto L111
L115:
	;
	if v482 == int32(0) {
		goto L99
	} else {
		goto L127
	}
L116:
	;
	if v406&int32(3) == int32(0) {
		v462 = v402
		goto L117
	} else {
		goto L118
	}
L117:
	;
	if base.Ui32(v462) <= base.Ui32(int32(3)) {
		v482 = v462
		goto L115
	} else {
		goto L123
	}
L118:
	;
	v447 = v402
	goto L119
L119:
	;
	if v447 == int32(0) {
		goto L99
	} else {
		goto L121
	}
L120:
	;
	v462 = v453
	goto L117
L121:
	;
	v453 = v447 + int32(-1)
	v454 = v388 + v453
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v453))))
	*(*uint8)(unsafe.Add(mBase, uint32(v454))) = uint8(v456)
	if v454&int32(3) != 0 {
		v447 = v453
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v469 = v462
	goto L124
L124:
	;
	v473 = v469 + int32(-4)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v66+v473)))
	*(*int32)(unsafe.Add(mBase, uint32(v388+v473))) = v476
	if base.Ui32(int32(3)) < base.Ui32(v473) {
		v469 = v473
		goto L124
	} else {
		goto L126
	}
L125:
	;
	v482 = v473
	goto L115
L126:
	;
	goto L125
L127:
	;
	v489 = v482
	goto L128
L128:
	;
	v493 = v489 + int32(-1)
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v493))))
	*(*uint8)(unsafe.Add(mBase, uint32(v388+v493))) = uint8(v496)
	if v493 != 0 {
		v489 = v493
		goto L128
	} else {
		goto L130
	}
L130:
	;
	goto L99
L131:
	;
	v506 = v499
	v507 = v500
	v508 = v501
	goto L132
L132:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v506)))
	*(*int32)(unsafe.Add(mBase, uint32(v508))) = v510
	v512 = int32(4)
	v513 = v506 + v512
	v515 = v508 + v512
	v517 = v507 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v517) {
		v506 = v513
		v507 = v517
		v508 = v515
		goto L132
	} else {
		goto L134
	}
L133:
	;
	v521 = v513
	v522 = v517
	v523 = v515
	goto L103
L134:
	;
	goto L133
L135:
	;
	v528 = v521
	v529 = v522
	v530 = v523
	goto L136
L136:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528))))
	*(*uint8)(unsafe.Add(mBase, uint32(v530))) = uint8(v532)
	v534 = int32(1)
	v539 = v529 + int32(-1)
	if v539 != 0 {
		v528 = v528 + v534
		v529 = v539
		v530 = v530 + v534
		goto L136
	} else {
		goto L138
	}
L137:
	;
	goto L99
L138:
	;
	goto L137
L139:
	;
	v572 = v567
	goto L141
L140:
	;
	v572 = int32(base.Ui32(v555) >> (uint(int32(1)) % 32))
	goto L141
L141:
	;
	v576 = int32(1)
	v583 = int32(4)
	v588 = F_valkey_realloc(m, l0, v561+(int32(0)-v561)&v560+v572+(v551^int32(-1))<<(uint(v576)%32)&(int32(0)-v551&v576)&v583+v583)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	return int32(0)
L143:
	;
	if v588 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v592 = v588
	goto L146
L145:
	;
	v592 = l0
	goto L146
L146:
	;
	return v592
}
func F_raxSize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	return v2
}
func F_raxStop(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 == l0+int32(24) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
		if v9 == l0+int32(168) {
			return
		} else {
			F_valkey_free(m, v9)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_valkey_free(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
			if v9 == l0+int32(168) {
				return
			} else {
				F_valkey_free(m, v9)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
