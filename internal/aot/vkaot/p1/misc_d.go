package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___divtf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v65 int32
	_ = v65
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v76 int32
	_ = v76
	var v109 int64
	_ = v109
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v149 int64
	_ = v149
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v198 int64
	_ = v198
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v212 int64
	_ = v212
	var v213 int64
	_ = v213
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v227 int64
	_ = v227
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v245 int64
	_ = v245
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v249 int64
	_ = v249
	var v253 int64
	_ = v253
	var v260 int64
	_ = v260
	var v272 int32
	_ = v272
	var v273 int64
	_ = v273
	var v278 int64
	_ = v278
	var v279 int64
	_ = v279
	var v286 int64
	_ = v286
	var v287 int64
	_ = v287
	var v289 int64
	_ = v289
	var v292 int64
	_ = v292
	var v293 int64
	_ = v293
	var v295 int64
	_ = v295
	var v296 int64
	_ = v296
	var v300 int64
	_ = v300
	var v307 int64
	_ = v307
	var v319 int32
	_ = v319
	var v320 int64
	_ = v320
	var v327 int64
	_ = v327
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v337 int64
	_ = v337
	var v338 int64
	_ = v338
	var v340 int64
	_ = v340
	var v343 int64
	_ = v343
	var v344 int64
	_ = v344
	var v346 int64
	_ = v346
	var v347 int64
	_ = v347
	var v351 int64
	_ = v351
	var v358 int64
	_ = v358
	var v370 int32
	_ = v370
	var v371 int64
	_ = v371
	var v377 int64
	_ = v377
	var v378 int64
	_ = v378
	var v384 int64
	_ = v384
	var v385 int64
	_ = v385
	var v387 int64
	_ = v387
	var v390 int64
	_ = v390
	var v391 int64
	_ = v391
	var v393 int64
	_ = v393
	var v394 int64
	_ = v394
	var v398 int64
	_ = v398
	var v405 int64
	_ = v405
	var v417 int32
	_ = v417
	var v418 int64
	_ = v418
	var v425 int64
	_ = v425
	var v428 int64
	_ = v428
	var v429 int64
	_ = v429
	var v435 int64
	_ = v435
	var v436 int64
	_ = v436
	var v438 int64
	_ = v438
	var v441 int64
	_ = v441
	var v442 int64
	_ = v442
	var v444 int64
	_ = v444
	var v445 int64
	_ = v445
	var v449 int64
	_ = v449
	var v456 int64
	_ = v456
	var v468 int32
	_ = v468
	var v469 int64
	_ = v469
	var v475 int64
	_ = v475
	var v476 int64
	_ = v476
	var v482 int64
	_ = v482
	var v483 int64
	_ = v483
	var v485 int64
	_ = v485
	var v488 int64
	_ = v488
	var v489 int64
	_ = v489
	var v491 int64
	_ = v491
	var v492 int64
	_ = v492
	var v496 int64
	_ = v496
	var v503 int64
	_ = v503
	var v515 int32
	_ = v515
	var v516 int64
	_ = v516
	var v523 int64
	_ = v523
	var v526 int64
	_ = v526
	var v527 int64
	_ = v527
	var v533 int64
	_ = v533
	var v534 int64
	_ = v534
	var v536 int64
	_ = v536
	var v539 int64
	_ = v539
	var v540 int64
	_ = v540
	var v542 int64
	_ = v542
	var v543 int64
	_ = v543
	var v547 int64
	_ = v547
	var v554 int64
	_ = v554
	var v566 int32
	_ = v566
	var v567 int64
	_ = v567
	var v573 int64
	_ = v573
	var v574 int64
	_ = v574
	var v580 int64
	_ = v580
	var v581 int64
	_ = v581
	var v583 int64
	_ = v583
	var v586 int64
	_ = v586
	var v587 int64
	_ = v587
	var v589 int64
	_ = v589
	var v590 int64
	_ = v590
	var v594 int64
	_ = v594
	var v601 int64
	_ = v601
	var v613 int32
	_ = v613
	var v614 int64
	_ = v614
	var v615 int64
	_ = v615
	var v622 int64
	_ = v622
	var v627 int64
	_ = v627
	var v633 int64
	_ = v633
	var v634 int64
	_ = v634
	var v636 int64
	_ = v636
	var v639 int64
	_ = v639
	var v640 int64
	_ = v640
	var v642 int64
	_ = v642
	var v643 int64
	_ = v643
	var v647 int64
	_ = v647
	var v654 int64
	_ = v654
	var v666 int32
	_ = v666
	var v668 int64
	_ = v668
	var v669 int64
	_ = v669
	var v675 int64
	_ = v675
	var v676 int64
	_ = v676
	var v678 int64
	_ = v678
	var v681 int64
	_ = v681
	var v682 int64
	_ = v682
	var v684 int64
	_ = v684
	var v685 int64
	_ = v685
	var v689 int64
	_ = v689
	var v696 int64
	_ = v696
	var v708 int32
	_ = v708
	var v709 int64
	_ = v709
	var v715 int64
	_ = v715
	var v716 int64
	_ = v716
	var v721 int64
	_ = v721
	var v722 int64
	_ = v722
	var v730 int64
	_ = v730
	var v736 int64
	_ = v736
	var v737 int64
	_ = v737
	var v739 int64
	_ = v739
	var v742 int64
	_ = v742
	var v743 int64
	_ = v743
	var v745 int64
	_ = v745
	var v746 int64
	_ = v746
	var v750 int64
	_ = v750
	var v757 int64
	_ = v757
	var v769 int32
	_ = v769
	var v771 int64
	_ = v771
	var v772 int64
	_ = v772
	var v778 int64
	_ = v778
	var v779 int64
	_ = v779
	var v781 int64
	_ = v781
	var v784 int64
	_ = v784
	var v785 int64
	_ = v785
	var v787 int64
	_ = v787
	var v788 int64
	_ = v788
	var v792 int64
	_ = v792
	var v799 int64
	_ = v799
	var v811 int32
	_ = v811
	var v812 int64
	_ = v812
	var v813 int64
	_ = v813
	var v814 int64
	_ = v814
	var v815 int64
	_ = v815
	var v816 int64
	_ = v816
	var v822 int64
	_ = v822
	var v826 int64
	_ = v826
	var v828 int64
	_ = v828
	var v829 int64
	_ = v829
	var v830 int64
	_ = v830
	var v832 int64
	_ = v832
	var v834 int64
	_ = v834
	var v836 int64
	_ = v836
	var v837 int64
	_ = v837
	var v839 int64
	_ = v839
	var v841 int64
	_ = v841
	var v846 int64
	_ = v846
	var v862 int64
	_ = v862
	var v864 int64
	_ = v864
	var v866 int64
	_ = v866
	var v869 int64
	_ = v869
	var v870 int64
	_ = v870
	var v872 int64
	_ = v872
	var v877 int64
	_ = v877
	var v879 int64
	_ = v879
	var v885 int64
	_ = v885
	var v887 int64
	_ = v887
	var v898 int64
	_ = v898
	var v903 int64
	_ = v903
	var v904 int64
	_ = v904
	var v906 int64
	_ = v906
	var v910 int64
	_ = v910
	var v912 int64
	_ = v912
	var v916 int64
	_ = v916
	var v920 int64
	_ = v920
	var v922 int64
	_ = v922
	var v924 int64
	_ = v924
	var v926 int64
	_ = v926
	var v940 int64
	_ = v940
	var v944 int64
	_ = v944
	var v946 int64
	_ = v946
	var v954 int64
	_ = v954
	var v963 int64
	_ = v963
	var v966 int64
	_ = v966
	var v971 int32
	_ = v971
	var v976 int64
	_ = v976
	var v977 int64
	_ = v977
	var v979 int64
	_ = v979
	var v982 int64
	_ = v982
	var v983 int64
	_ = v983
	var v985 int64
	_ = v985
	var v986 int64
	_ = v986
	var v990 int64
	_ = v990
	var v997 int64
	_ = v997
	var v1014 int64
	_ = v1014
	var v1016 int64
	_ = v1016
	var v1017 int64
	_ = v1017
	var v1026 int32
	_ = v1026
	var v1027 int64
	_ = v1027
	var v1031 int64
	_ = v1031
	var v1033 int64
	_ = v1033
	var v1038 int64
	_ = v1038
	var v1039 int64
	_ = v1039
	var v1041 int64
	_ = v1041
	var v1044 int64
	_ = v1044
	var v1045 int64
	_ = v1045
	var v1047 int64
	_ = v1047
	var v1048 int64
	_ = v1048
	var v1052 int64
	_ = v1052
	var v1059 int64
	_ = v1059
	var v1076 int64
	_ = v1076
	var v1078 int64
	_ = v1078
	var v1079 int64
	_ = v1079
	var v1088 int64
	_ = v1088
	var v1089 int64
	_ = v1089
	var v1090 int64
	_ = v1090
	var v1091 int64
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int64
	_ = v1093
	var v1094 int64
	_ = v1094
	var v1102 int64
	_ = v1102
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1138 int64
	_ = v1138
	var v1142 int64
	_ = v1142
	var v1143 int64
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1167 int64
	_ = v1167
	var v1171 int64
	_ = v1171
	var v1172 int64
	_ = v1172
	var v1177 int32
	_ = v1177
	var v1178 int64
	_ = v1178
	var v1183 int64
	_ = v1183
	var v1188 int64
	_ = v1188
	var v1189 int64
	_ = v1189
	var v1191 int64
	_ = v1191
	var v1194 int64
	_ = v1194
	var v1195 int64
	_ = v1195
	var v1197 int64
	_ = v1197
	var v1198 int64
	_ = v1198
	var v1202 int64
	_ = v1202
	var v1209 int64
	_ = v1209
	var v1224 int64
	_ = v1224
	var v1229 int64
	_ = v1229
	var v1230 int64
	_ = v1230
	var v1232 int64
	_ = v1232
	var v1237 int64
	_ = v1237
	var v1239 int64
	_ = v1239
	var v1244 int64
	_ = v1244
	var v1245 int64
	_ = v1245
	var v1246 int64
	_ = v1246
	var v1247 int64
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1251 int64
	_ = v1251
	var v1252 int64
	_ = v1252
	var v1257 int64
	_ = v1257
	var v1260 int64
	_ = v1260
	var v1263 int64
	_ = v1263
	var v1266 int64
	_ = v1266
	var v1267 int64
	_ = v1267
	var v1271 int64
	_ = v1271
	var v1278 int64
	_ = v1278
	var v1289 int64
	_ = v1289
	var v1290 int64
	_ = v1290
	var v1295 int64
	_ = v1295
	var v1298 int64
	_ = v1298
	var v1301 int64
	_ = v1301
	var v1304 int64
	_ = v1304
	var v1305 int64
	_ = v1305
	var v1309 int64
	_ = v1309
	var v1316 int64
	_ = v1316
	var v1328 int64
	_ = v1328
	var v1329 int64
	_ = v1329
	var v1333 int64
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1338 int64
	_ = v1338
	var v1341 int64
	_ = v1341
	var v1344 int64
	_ = v1344
	var v1350 int64
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1356 int64
	_ = v1356
	var v1359 int64
	_ = v1359
	var v1362 int64
	_ = v1362
	var v1366 int64
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1372 int64
	_ = v1372
	var v1377 int64
	_ = v1377
	var v1383 int64
	_ = v1383
	v26 = m.G0
	v28 = v26 - int32(336)
	m.G0 = v28
	v30 = int64(281474976710655)
	v31 = l4 & v30
	v33 = l2 & v30
	v36 = (l4 ^ l2) & int64(-9223372036854775807-1)
	v37 = int64(48)
	v40 = int32(32767)
	v41 = base.I32_wrap_i64(int64(base.Ui64(l4)>>(uint(v37)%64))) & v40
	v46 = base.I32_wrap_i64(int64(base.Ui64(l2)>>(uint(v37)%64))) & v40
	if base.Ui32(v46+int32(-32767)) < base.Ui32(int32(-32766)) {
		v60 = l2 & int64(9223372036854775807)
		v61 = int64(9223090561878065152)
		if v60 == v61 {
			v65 = base.B2i32(l1 == int64(0))
		} else {
			v65 = base.B2i32(base.Ui64(v60) < base.Ui64(v61))
		}
		if v65 != 0 {
			v71 = l4 & int64(9223372036854775807)
			v72 = int64(9223090561878065152)
			if v71 == v72 {
				v76 = base.B2i32(l3 == int64(0))
			} else {
				v76 = base.B2i32(base.Ui64(v71) < base.Ui64(v72))
			}
			if v76 != 0 {
				if l1|(v60^int64(9223090561878065152)) != int64(0) {
					if l3|(v71^int64(9223090561878065152)) != int64(0) {
						if l1|v60 != int64(0) {
							if l3|v71 != int64(0) {
								if base.Ui64(int64(281474976710655)) < base.Ui64(v60) {
									v164 = l1
									v165 = v33
									v166 = int32(0)
								} else {
									v121 = v28 + int32(320)
									v123 = base.B2i32(v33 == int64(0))
									if v33 == int64(0) {
										v124 = l1
									} else {
										v124 = v33
									}
									v130 = base.I32_wrap_i64(base.I64_clz(v124) + base.I64_extend_i32_u(v123<<(uint(int32(6))%32)))
									v132 = v130 + int32(-15)
									if v132&int32(64) == int32(0) {
										if v132 == int32(0) {
											v153 = l1
											v154 = v33
										} else {
											v149 = base.I64_extend_i32_u(v132)
											v153 = l1 << (uint(v149) % 64)
											v154 = int64(base.Ui64(l1)>>(uint(base.I64_extend_i32_u(int32(64)-v132))%64)) | v33<<(uint(v149)%64)
										}
									} else {
										v153 = int64(0)
										v154 = l1 << (uint(base.I64_extend_i32_u(v130+int32(-79))) % 64)
									}
									*(*int64)(unsafe.Add(mBase, uint32(v121))) = v153
									*(*int64)(unsafe.Add(mBase, uint32(v121)+8)) = v154
									v162 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(328))))
									v163 = *(*int64)(unsafe.Add(mBase, uint32(v28)+320))
									v164 = v163
									v165 = v162
									v166 = int32(16) - v130
								}
								if base.Ui64(int64(281474976710655)) < base.Ui64(v71) {
									v214 = v164
									v216 = l3
									v217 = v31
									v218 = v165
									v219 = v166
								} else {
									v170 = v28 + int32(304)
									v172 = base.B2i32(v31 == int64(0))
									if v31 == int64(0) {
										v173 = l3
									} else {
										v173 = v31
									}
									v179 = base.I32_wrap_i64(base.I64_clz(v173) + base.I64_extend_i32_u(v172<<(uint(int32(6))%32)))
									v181 = v179 + int32(-15)
									if v181&int32(64) == int32(0) {
										if v181 == int32(0) {
											v202 = l3
											v203 = v31
										} else {
											v198 = base.I64_extend_i32_u(v181)
											v202 = l3 << (uint(v198) % 64)
											v203 = int64(base.Ui64(l3)>>(uint(base.I64_extend_i32_u(int32(64)-v181))%64)) | v31<<(uint(v198)%64)
										}
									} else {
										v202 = int64(0)
										v203 = l3 << (uint(base.I64_extend_i32_u(v179+int32(-79))) % 64)
									}
									*(*int64)(unsafe.Add(mBase, uint32(v170))) = v202
									*(*int64)(unsafe.Add(mBase, uint32(v170)+8)) = v203
									v212 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(312))))
									v213 = *(*int64)(unsafe.Add(mBase, uint32(v28)+304))
									v214 = v164
									v216 = v213
									v217 = v212
									v218 = v165
									v219 = v179 + v166 + int32(-16)
								}
								v223 = v28 + int32(288)
								v227 = v217 | int64(281474976710656)
								v230 = int64(base.Ui64(v216)>>(uint(int64(49))%64)) | v227<<(uint(int64(15))%64)
								v231 = int64(0)
								v233 = int64(8432131802713292800) - v230
								v239 = int64(32)
								v240 = int64(base.Ui64(v233) >> (uint(v239) % 64))
								v242 = int64(base.Ui64(v230) >> (uint(v239) % 64))
								v245 = int64(4294967295)
								v246 = v233 & v245
								v248 = v230 & v245
								v249 = v246 * v248
								v253 = int64(base.Ui64(v249)>>(uint(v239)%64)) + v246*v242
								v260 = v253&v245 + v240*v248
								*(*int64)(unsafe.Add(mBase, uint32(v223)+8)) = v231*v230 + v231*v233 + v240*v242 + int64(base.Ui64(v253)>>(uint(v239)%64)) + int64(base.Ui64(v260)>>(uint(v239)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v223))) = v260<<(uint(v239)%64) | v249&v245
								v272 = v28 + int32(272)
								v273 = int64(0)
								v278 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(296))))
								v279 = v273 - v278
								v286 = int64(32)
								v287 = int64(base.Ui64(v233) >> (uint(v286) % 64))
								v289 = int64(base.Ui64(v279) >> (uint(v286) % 64))
								v292 = int64(4294967295)
								v293 = v233 & v292
								v295 = v279 & v292
								v296 = v293 * v295
								v300 = int64(base.Ui64(v296)>>(uint(v286)%64)) + v293*v289
								v307 = v300&v292 + v287*v295
								*(*int64)(unsafe.Add(mBase, uint32(v272)+8)) = v273*v279 + v273*v233 + v287*v289 + int64(base.Ui64(v300)>>(uint(v286)%64)) + int64(base.Ui64(v307)>>(uint(v286)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v272))) = v307<<(uint(v286)%64) | v296&v292
								v319 = v28 + int32(256)
								v320 = *(*int64)(unsafe.Add(mBase, uint32(v28)+272))
								v327 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(280))))
								v330 = int64(base.Ui64(v320)>>(uint(int64(63))%64)) | v327<<(uint(int64(1))%64)
								v331 = int64(0)
								v337 = int64(32)
								v338 = int64(base.Ui64(v230) >> (uint(v337) % 64))
								v340 = int64(base.Ui64(v330) >> (uint(v337) % 64))
								v343 = int64(4294967295)
								v344 = v230 & v343
								v346 = v330 & v343
								v347 = v344 * v346
								v351 = int64(base.Ui64(v347)>>(uint(v337)%64)) + v344*v340
								v358 = v351&v343 + v338*v346
								*(*int64)(unsafe.Add(mBase, uint32(v319)+8)) = v331*v330 + v331*v230 + v338*v340 + int64(base.Ui64(v351)>>(uint(v337)%64)) + int64(base.Ui64(v358)>>(uint(v337)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v319))) = v358<<(uint(v337)%64) | v347&v343
								v370 = v28 + int32(240)
								v371 = int64(0)
								v377 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(264))))
								v378 = v371 - v377
								v384 = int64(32)
								v385 = int64(base.Ui64(v378) >> (uint(v384) % 64))
								v387 = int64(base.Ui64(v330) >> (uint(v384) % 64))
								v390 = int64(4294967295)
								v391 = v378 & v390
								v393 = v330 & v390
								v394 = v391 * v393
								v398 = int64(base.Ui64(v394)>>(uint(v384)%64)) + v391*v387
								v405 = v398&v390 + v385*v393
								*(*int64)(unsafe.Add(mBase, uint32(v370)+8)) = v371*v330 + v371*v378 + v385*v387 + int64(base.Ui64(v398)>>(uint(v384)%64)) + int64(base.Ui64(v405)>>(uint(v384)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v370))) = v405<<(uint(v384)%64) | v394&v390
								v417 = v28 + int32(224)
								v418 = *(*int64)(unsafe.Add(mBase, uint32(v28)+240))
								v425 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(248))))
								v428 = int64(base.Ui64(v418)>>(uint(int64(63))%64)) | v425<<(uint(int64(1))%64)
								v429 = int64(0)
								v435 = int64(32)
								v436 = int64(base.Ui64(v230) >> (uint(v435) % 64))
								v438 = int64(base.Ui64(v428) >> (uint(v435) % 64))
								v441 = int64(4294967295)
								v442 = v230 & v441
								v444 = v428 & v441
								v445 = v442 * v444
								v449 = int64(base.Ui64(v445)>>(uint(v435)%64)) + v442*v438
								v456 = v449&v441 + v436*v444
								*(*int64)(unsafe.Add(mBase, uint32(v417)+8)) = v429*v428 + v429*v230 + v436*v438 + int64(base.Ui64(v449)>>(uint(v435)%64)) + int64(base.Ui64(v456)>>(uint(v435)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v417))) = v456<<(uint(v435)%64) | v445&v441
								v468 = v28 + int32(208)
								v469 = int64(0)
								v475 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(232))))
								v476 = v469 - v475
								v482 = int64(32)
								v483 = int64(base.Ui64(v476) >> (uint(v482) % 64))
								v485 = int64(base.Ui64(v428) >> (uint(v482) % 64))
								v488 = int64(4294967295)
								v489 = v476 & v488
								v491 = v428 & v488
								v492 = v489 * v491
								v496 = int64(base.Ui64(v492)>>(uint(v482)%64)) + v489*v485
								v503 = v496&v488 + v483*v491
								*(*int64)(unsafe.Add(mBase, uint32(v468)+8)) = v469*v428 + v469*v476 + v483*v485 + int64(base.Ui64(v496)>>(uint(v482)%64)) + int64(base.Ui64(v503)>>(uint(v482)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v468))) = v503<<(uint(v482)%64) | v492&v488
								v515 = v28 + int32(192)
								v516 = *(*int64)(unsafe.Add(mBase, uint32(v28)+208))
								v523 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(216))))
								v526 = int64(base.Ui64(v516)>>(uint(int64(63))%64)) | v523<<(uint(int64(1))%64)
								v527 = int64(0)
								v533 = int64(32)
								v534 = int64(base.Ui64(v230) >> (uint(v533) % 64))
								v536 = int64(base.Ui64(v526) >> (uint(v533) % 64))
								v539 = int64(4294967295)
								v540 = v230 & v539
								v542 = v526 & v539
								v543 = v540 * v542
								v547 = int64(base.Ui64(v543)>>(uint(v533)%64)) + v540*v536
								v554 = v547&v539 + v534*v542
								*(*int64)(unsafe.Add(mBase, uint32(v515)+8)) = v527*v526 + v527*v230 + v534*v536 + int64(base.Ui64(v547)>>(uint(v533)%64)) + int64(base.Ui64(v554)>>(uint(v533)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v515))) = v554<<(uint(v533)%64) | v543&v539
								v566 = v28 + int32(176)
								v567 = int64(0)
								v573 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(200))))
								v574 = v567 - v573
								v580 = int64(32)
								v581 = int64(base.Ui64(v574) >> (uint(v580) % 64))
								v583 = int64(base.Ui64(v526) >> (uint(v580) % 64))
								v586 = int64(4294967295)
								v587 = v574 & v586
								v589 = v526 & v586
								v590 = v587 * v589
								v594 = int64(base.Ui64(v590)>>(uint(v580)%64)) + v587*v583
								v601 = v594&v586 + v581*v589
								*(*int64)(unsafe.Add(mBase, uint32(v566)+8)) = v567*v526 + v567*v574 + v581*v583 + int64(base.Ui64(v594)>>(uint(v580)%64)) + int64(base.Ui64(v601)>>(uint(v580)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v566))) = v601<<(uint(v580)%64) | v590&v586
								v613 = v28 + int32(160)
								v614 = int64(0)
								v615 = *(*int64)(unsafe.Add(mBase, uint32(v28)+176))
								v622 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(184))))
								v627 = int64(base.Ui64(v615)>>(uint(int64(63))%64)) | v622<<(uint(int64(1))%64) + int64(-1)
								v633 = int64(32)
								v634 = int64(base.Ui64(v627) >> (uint(v633) % 64))
								v636 = int64(base.Ui64(v230) >> (uint(v633) % 64))
								v639 = int64(4294967295)
								v640 = v627 & v639
								v642 = v230 & v639
								v643 = v640 * v642
								v647 = int64(base.Ui64(v643)>>(uint(v633)%64)) + v640*v636
								v654 = v647&v639 + v634*v642
								*(*int64)(unsafe.Add(mBase, uint32(v613)+8)) = v614*v230 + v614*v627 + v634*v636 + int64(base.Ui64(v647)>>(uint(v633)%64)) + int64(base.Ui64(v654)>>(uint(v633)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v613))) = v654<<(uint(v633)%64) | v643&v639
								v666 = v28 + int32(144)
								v668 = v216 << (uint(int64(15)) % 64)
								v669 = int64(0)
								v675 = int64(32)
								v676 = int64(base.Ui64(v627) >> (uint(v675) % 64))
								v678 = int64(base.Ui64(v668) >> (uint(v675) % 64))
								v681 = int64(4294967295)
								v682 = v627 & v681
								v684 = v668 & v681
								v685 = v682 * v684
								v689 = int64(base.Ui64(v685)>>(uint(v675)%64)) + v682*v678
								v696 = v689&v681 + v676*v684
								*(*int64)(unsafe.Add(mBase, uint32(v666)+8)) = v669*v668 + v669*v627 + v676*v678 + int64(base.Ui64(v689)>>(uint(v675)%64)) + int64(base.Ui64(v696)>>(uint(v675)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v666))) = v696<<(uint(v675)%64) | v685&v681
								v708 = v28 + int32(112)
								v709 = int64(0)
								v715 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(168))))
								v716 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
								v721 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(152))))
								v722 = v716 + v721
								v730 = v709 - (v715 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v722) < base.Ui64(v716))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(1)) < base.Ui64(v722))))
								v736 = int64(32)
								v737 = int64(base.Ui64(v730) >> (uint(v736) % 64))
								v739 = int64(base.Ui64(v627) >> (uint(v736) % 64))
								v742 = int64(4294967295)
								v743 = v730 & v742
								v745 = v627 & v742
								v746 = v743 * v745
								v750 = int64(base.Ui64(v746)>>(uint(v736)%64)) + v743*v739
								v757 = v750&v742 + v737*v745
								*(*int64)(unsafe.Add(mBase, uint32(v708)+8)) = v709*v627 + v709*v730 + v737*v739 + int64(base.Ui64(v750)>>(uint(v736)%64)) + int64(base.Ui64(v757)>>(uint(v736)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v708))) = v757<<(uint(v736)%64) | v746&v742
								v769 = v28 + int32(128)
								v771 = int64(1) - v722
								v772 = int64(0)
								v778 = int64(32)
								v779 = int64(base.Ui64(v627) >> (uint(v778) % 64))
								v781 = int64(base.Ui64(v771) >> (uint(v778) % 64))
								v784 = int64(4294967295)
								v785 = v627 & v784
								v787 = v771 & v784
								v788 = v785 * v787
								v792 = int64(base.Ui64(v788)>>(uint(v778)%64)) + v785*v781
								v799 = v792&v784 + v779*v787
								*(*int64)(unsafe.Add(mBase, uint32(v769)+8)) = v772*v771 + v772*v627 + v779*v781 + int64(base.Ui64(v792)>>(uint(v778)%64)) + int64(base.Ui64(v799)>>(uint(v778)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v769))) = v799<<(uint(v778)%64) | v788&v784
								v811 = v219 + (v46 - v41)
								v812 = *(*int64)(unsafe.Add(mBase, uint32(v28)+112))
								v813 = int64(1)
								v814 = v812 << (uint(v813) % 64)
								v815 = *(*int64)(unsafe.Add(mBase, uint32(v28)+128))
								v816 = int64(63)
								v822 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(136))))
								v826 = v814 + (int64(base.Ui64(v815)>>(uint(v816)%64)) | v822<<(uint(v813)%64))
								v828 = v826 + int64(-13927)
								v829 = int64(32)
								v830 = int64(base.Ui64(v828) >> (uint(v829) % 64))
								v832 = v218 | int64(281474976710656)
								v834 = v832 << (uint(v813) % 64)
								v836 = int64(base.Ui64(v834) >> (uint(v829) % 64))
								v837 = v830 * v836
								v839 = v214 << (uint(v813) % 64)
								v841 = int64(base.Ui64(v839) >> (uint(v829) % 64))
								v846 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(120))))
								v862 = v846<<(uint(v813)%64) | int64(base.Ui64(v812)>>(uint(v816)%64)) + int64(base.Ui64(v822)>>(uint(v816)%64)) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v826) < base.Ui64(v814))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v828) < base.Ui64(v826))) + int64(-1)
								v864 = int64(base.Ui64(v862) >> (uint(v829) % 64))
								v866 = v837 + v841*v864
								v869 = int64(4294967295)
								v870 = v862 & v869
								v872 = int64(base.Ui64(v214) >> (uint(v816) % 64))
								v877 = (v872 | v218<<(uint(v813)%64)) & v869
								v879 = v866 + v870*v877
								v885 = v870 * v836
								v887 = v885 + v877*v864
								v898 = v879 + v887<<(uint(v829)%64)
								v903 = v828 & v869
								v904 = v903 * v877
								v906 = v904 + v830*v841
								v910 = v839 & int64(4294967294)
								v912 = v906 + v870*v910
								v916 = v898 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v906) < base.Ui64(v904))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v912) < base.Ui64(v906))))
								v920 = v903 * v836
								v922 = v920 + v910*v864
								v924 = v922 + v830*v877
								v926 = v924 + v870*v841
								v940 = v916 + (int64(base.Ui64(v926)>>(uint(v829)%64)) | (base.I64_extend_i32_u(base.B2i32(base.Ui64(v922) < base.Ui64(v920)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v924) < base.Ui64(v922)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v926) < base.Ui64(v924))))<<(uint(v829)%64))
								v944 = v830 * v910
								v946 = v944 + v903*v841
								v954 = v912 + (int64(base.Ui64(v946)>>(uint(v829)%64)) | base.I64_extend_i32_u(base.B2i32(base.Ui64(v946) < base.Ui64(v944)))<<(uint(v829)%64))
								v963 = v940 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v954) < base.Ui64(v912))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v954+v926<<(uint(v829)%64)) < base.Ui64(v954))))
								v966 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v866) < base.Ui64(v837))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v879) < base.Ui64(v866))) + v864*v836 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v887) < base.Ui64(v885)))<<(uint(v829)%64) | int64(base.Ui64(v887)>>(uint(v829)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v898) < base.Ui64(v879))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v916) < base.Ui64(v898))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v940) < base.Ui64(v916))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v963) < base.Ui64(v940)))
								if base.Ui64(int64(562949953421311)) < base.Ui64(v966) {
									v1026 = v28 + int32(96)
									v1027 = int64(1)
									v1031 = int64(base.Ui64(v963)>>(uint(v1027)%64)) | v966<<(uint(int64(63))%64)
									v1033 = int64(base.Ui64(v966) >> (uint(v1027) % 64))
									v1038 = int64(32)
									v1039 = int64(base.Ui64(v216) >> (uint(v1038) % 64))
									v1041 = int64(base.Ui64(v1031) >> (uint(v1038) % 64))
									v1044 = int64(4294967295)
									v1045 = v216 & v1044
									v1047 = v1031 & v1044
									v1048 = v1045 * v1047
									v1052 = int64(base.Ui64(v1048)>>(uint(v1038)%64)) + v1045*v1041
									v1059 = v1052&v1044 + v1039*v1047
									*(*int64)(unsafe.Add(mBase, uint32(v1026)+8)) = v227*v1031 + v1033*v216 + v1039*v1041 + int64(base.Ui64(v1052)>>(uint(v1038)%64)) + int64(base.Ui64(v1059)>>(uint(v1038)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v1026))) = v1059<<(uint(v1038)%64) | v1048&v1044
									v1076 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(104))))
									v1078 = *(*int64)(unsafe.Add(mBase, uint32(v28)+96))
									v1079 = int64(0)
									v1088 = v1031
									v1089 = v1033
									v1090 = v214<<(uint(int64(48))%64) - v1076 - base.I64_extend_i32_u(base.B2i32(v1078 != v1079))
									v1091 = v1079 - v1078
									v1092 = v811 + int32(16383)
									v1093 = v832
									v1094 = v214
								} else {
									v971 = v28 + int32(80)
									v976 = int64(32)
									v977 = int64(base.Ui64(v216) >> (uint(v976) % 64))
									v979 = int64(base.Ui64(v963) >> (uint(v976) % 64))
									v982 = int64(4294967295)
									v983 = v216 & v982
									v985 = v963 & v982
									v986 = v983 * v985
									v990 = int64(base.Ui64(v986)>>(uint(v976)%64)) + v983*v979
									v997 = v990&v982 + v977*v985
									*(*int64)(unsafe.Add(mBase, uint32(v971)+8)) = v227*v963 + v966*v216 + v977*v979 + int64(base.Ui64(v990)>>(uint(v976)%64)) + int64(base.Ui64(v997)>>(uint(v976)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v971))) = v997<<(uint(v976)%64) | v986&v982
									v1014 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(88))))
									v1016 = *(*int64)(unsafe.Add(mBase, uint32(v28)+80))
									v1017 = int64(0)
									v1088 = v963
									v1089 = v966
									v1090 = v214<<(uint(int64(49))%64) - v1014 - base.I64_extend_i32_u(base.B2i32(v1016 != v1017))
									v1091 = v1017 - v1016
									v1092 = v811 + int32(16382)
									v1093 = v834 | v872
									v1094 = v839
								}
								if v1092 < int32(32767) {
									if v1092 < int32(1) {
										if int32(-113) < v1092 {
											v1118 = int32(64)
											v1119 = v28 + v1118
											v1121 = int32(1) - v1092
											if v1121&v1118 == int32(0) {
												if v1121 == int32(0) {
													v1142 = v1088
													v1143 = v1089
												} else {
													v1138 = base.I64_extend_i32_u(v1121)
													v1142 = v1089<<(uint(base.I64_extend_i32_u(int32(64)-v1121))%64) | int64(base.Ui64(v1088)>>(uint(v1138)%64))
													v1143 = int64(base.Ui64(v1089) >> (uint(v1138) % 64))
												}
											} else {
												v1142 = int64(base.Ui64(v1089) >> (uint(base.I64_extend_i32_u(v1121+int32(-64))) % 64))
												v1143 = int64(0)
											}
											*(*int64)(unsafe.Add(mBase, uint32(v1119))) = v1142
											*(*int64)(unsafe.Add(mBase, uint32(v1119)+8)) = v1143
											v1148 = v28 + int32(48)
											v1150 = v1092 + int32(112)
											if v1150&int32(64) == int32(0) {
												if v1150 == int32(0) {
													v1171 = v1094
													v1172 = v1093
												} else {
													v1167 = base.I64_extend_i32_u(v1150)
													v1171 = v1094 << (uint(v1167) % 64)
													v1172 = int64(base.Ui64(v1094)>>(uint(base.I64_extend_i32_u(int32(64)-v1150))%64)) | v1093<<(uint(v1167)%64)
												}
											} else {
												v1171 = int64(0)
												v1172 = v1094 << (uint(base.I64_extend_i32_u(v1092+int32(48))) % 64)
											}
											*(*int64)(unsafe.Add(mBase, uint32(v1148))) = v1171
											*(*int64)(unsafe.Add(mBase, uint32(v1148)+8)) = v1172
											v1177 = v28 + int32(32)
											v1178 = *(*int64)(unsafe.Add(mBase, uint32(v28)+64))
											v1183 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(72))))
											v1188 = int64(32)
											v1189 = int64(base.Ui64(v1178) >> (uint(v1188) % 64))
											v1191 = int64(base.Ui64(v216) >> (uint(v1188) % 64))
											v1194 = int64(4294967295)
											v1195 = v1178 & v1194
											v1197 = v216 & v1194
											v1198 = v1195 * v1197
											v1202 = int64(base.Ui64(v1198)>>(uint(v1188)%64)) + v1195*v1191
											v1209 = v1202&v1194 + v1189*v1197
											*(*int64)(unsafe.Add(mBase, uint32(v1177)+8)) = v1183*v216 + v227*v1178 + v1189*v1191 + int64(base.Ui64(v1202)>>(uint(v1188)%64)) + int64(base.Ui64(v1209)>>(uint(v1188)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v1177))) = v1209<<(uint(v1188)%64) | v1198&v1194
											v1224 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(56))))
											v1229 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(40))))
											v1230 = int64(1)
											v1232 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
											v1237 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
											v1239 = v1232 << (uint(v1230) % 64)
											v1244 = v1224 - (v1229<<(uint(v1230)%64) | int64(base.Ui64(v1232)>>(uint(int64(63))%64))) - base.I64_extend_i32_u(base.B2i32(base.Ui64(v1237) < base.Ui64(v1239)))
											v1245 = v1178
											v1246 = v1237 - v1239
											v1247 = v1183
											v1250 = v28 + int32(16)
											v1251 = int64(3)
											v1252 = int64(0)
											v1257 = int64(32)
											v1260 = int64(base.Ui64(v216) >> (uint(v1257) % 64))
											v1263 = int64(4294967295)
											v1266 = v216 & v1263
											v1267 = v1251 * v1266
											v1271 = int64(base.Ui64(v1267)>>(uint(v1257)%64)) + v1251*v1260
											v1278 = v1271&v1263 + v1252*v1266
											*(*int64)(unsafe.Add(mBase, uint32(v1250)+8)) = v1252*v216 + v227*v1251 + v1252*v1260 + int64(base.Ui64(v1271)>>(uint(v1257)%64)) + int64(base.Ui64(v1278)>>(uint(v1257)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v1250))) = v1278<<(uint(v1257)%64) | v1267&v1263
											v1289 = int64(5)
											v1290 = int64(0)
											v1295 = int64(32)
											v1298 = int64(base.Ui64(v216) >> (uint(v1295) % 64))
											v1301 = int64(4294967295)
											v1304 = v216 & v1301
											v1305 = v1289 * v1304
											v1309 = int64(base.Ui64(v1305)>>(uint(v1295)%64)) + v1289*v1298
											v1316 = v1309&v1301 + v1290*v1304
											*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v1290*v216 + v227*v1289 + v1290*v1298 + int64(base.Ui64(v1309)>>(uint(v1295)%64)) + int64(base.Ui64(v1316)>>(uint(v1295)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1316<<(uint(v1295)%64) | v1305&v1301
											v1328 = v1245 & int64(1)
											v1329 = v1328 + v1246
											v1333 = v1244 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1329) < base.Ui64(v1328)))
											if v1333 == v227 {
												v1336 = base.B2i32(base.Ui64(v216) < base.Ui64(v1329))
											} else {
												v1336 = base.B2i32(base.Ui64(v227) < base.Ui64(v1333))
											}
											v1338 = v1245 + base.I64_extend_i32_u(v1336)
											v1341 = v1247 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1338) < base.Ui64(v1245)))
											v1344 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
											v1350 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(24))))
											if v1333 == v1350 {
												v1353 = base.B2i32(base.Ui64(v1344) < base.Ui64(v1329))
											} else {
												v1353 = base.B2i32(base.Ui64(v1350) < base.Ui64(v1333))
											}
											v1356 = v1338 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1341) < base.Ui64(int64(9223090561878065152)))&v1353)
											v1359 = v1341 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1356) < base.Ui64(v1338)))
											v1362 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
											v1366 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(8))))
											if v1333 == v1366 {
												v1369 = base.B2i32(base.Ui64(v1362) < base.Ui64(v1329))
											} else {
												v1369 = base.B2i32(base.Ui64(v1366) < base.Ui64(v1333))
											}
											v1372 = v1356 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1359) < base.Ui64(int64(9223090561878065152)))&v1369)
											v1377 = v1372
											v1383 = v1359 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1372) < base.Ui64(v1356))) | v36
										} else {
											v1377 = int64(0)
											v1383 = v36
										}
									} else {
										v1102 = int64(1)
										v1244 = v1090<<(uint(v1102)%64) | int64(base.Ui64(v1091)>>(uint(int64(63))%64))
										v1245 = v1088
										v1246 = v1091 << (uint(v1102) % 64)
										v1247 = base.I64_extend_i32_u(v1092)<<(uint(int64(48))%64) | v1089&int64(281474976710655)
										v1250 = v28 + int32(16)
										v1251 = int64(3)
										v1252 = int64(0)
										v1257 = int64(32)
										v1260 = int64(base.Ui64(v216) >> (uint(v1257) % 64))
										v1263 = int64(4294967295)
										v1266 = v216 & v1263
										v1267 = v1251 * v1266
										v1271 = int64(base.Ui64(v1267)>>(uint(v1257)%64)) + v1251*v1260
										v1278 = v1271&v1263 + v1252*v1266
										*(*int64)(unsafe.Add(mBase, uint32(v1250)+8)) = v1252*v216 + v227*v1251 + v1252*v1260 + int64(base.Ui64(v1271)>>(uint(v1257)%64)) + int64(base.Ui64(v1278)>>(uint(v1257)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v1250))) = v1278<<(uint(v1257)%64) | v1267&v1263
										v1289 = int64(5)
										v1290 = int64(0)
										v1295 = int64(32)
										v1298 = int64(base.Ui64(v216) >> (uint(v1295) % 64))
										v1301 = int64(4294967295)
										v1304 = v216 & v1301
										v1305 = v1289 * v1304
										v1309 = int64(base.Ui64(v1305)>>(uint(v1295)%64)) + v1289*v1298
										v1316 = v1309&v1301 + v1290*v1304
										*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v1290*v216 + v227*v1289 + v1290*v1298 + int64(base.Ui64(v1309)>>(uint(v1295)%64)) + int64(base.Ui64(v1316)>>(uint(v1295)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1316<<(uint(v1295)%64) | v1305&v1301
										v1328 = v1245 & int64(1)
										v1329 = v1328 + v1246
										v1333 = v1244 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1329) < base.Ui64(v1328)))
										if v1333 == v227 {
											v1336 = base.B2i32(base.Ui64(v216) < base.Ui64(v1329))
										} else {
											v1336 = base.B2i32(base.Ui64(v227) < base.Ui64(v1333))
										}
										v1338 = v1245 + base.I64_extend_i32_u(v1336)
										v1341 = v1247 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1338) < base.Ui64(v1245)))
										v1344 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
										v1350 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(24))))
										if v1333 == v1350 {
											v1353 = base.B2i32(base.Ui64(v1344) < base.Ui64(v1329))
										} else {
											v1353 = base.B2i32(base.Ui64(v1350) < base.Ui64(v1333))
										}
										v1356 = v1338 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1341) < base.Ui64(int64(9223090561878065152)))&v1353)
										v1359 = v1341 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1356) < base.Ui64(v1338)))
										v1362 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
										v1366 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(8))))
										if v1333 == v1366 {
											v1369 = base.B2i32(base.Ui64(v1362) < base.Ui64(v1329))
										} else {
											v1369 = base.B2i32(base.Ui64(v1366) < base.Ui64(v1333))
										}
										v1372 = v1356 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1359) < base.Ui64(int64(9223090561878065152)))&v1369)
										v1377 = v1372
										v1383 = v1359 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1372) < base.Ui64(v1356))) | v36
									}
								} else {
									v1377 = int64(0)
									v1383 = v36 | int64(9223090561878065152)
								}
							} else {
								v1377 = int64(0)
								v1383 = v36 | int64(9223090561878065152)
							}
						} else {
							if l3|v71 == int64(0) {
								v109 = int64(9223231299366420480)
							} else {
								v109 = v36
							}
							v1377 = int64(0)
							v1383 = v109
						}
					} else {
						v1377 = int64(0)
						v1383 = v36
					}
				} else {
					if base.B2i32(l3|(v71^int64(9223090561878065152)) == int64(0)) == int32(0) {
						v1377 = int64(0)
						v1383 = v36 | int64(9223090561878065152)
					} else {
						v1377 = int64(0)
						v1383 = int64(9223231299366420480)
					}
				}
			} else {
				v1377 = l3
				v1383 = l4 | int64(140737488355328)
			}
		} else {
			v1377 = l1
			v1383 = l2 | int64(140737488355328)
		}
	} else {
		v52 = int32(-32767)
		if base.Ui32(v52) < base.Ui32(v41+v52) {
			v214 = l1
			v216 = l3
			v217 = v31
			v218 = v33
			v219 = int32(0)
			v223 = v28 + int32(288)
			v227 = v217 | int64(281474976710656)
			v230 = int64(base.Ui64(v216)>>(uint(int64(49))%64)) | v227<<(uint(int64(15))%64)
			v231 = int64(0)
			v233 = int64(8432131802713292800) - v230
			v239 = int64(32)
			v240 = int64(base.Ui64(v233) >> (uint(v239) % 64))
			v242 = int64(base.Ui64(v230) >> (uint(v239) % 64))
			v245 = int64(4294967295)
			v246 = v233 & v245
			v248 = v230 & v245
			v249 = v246 * v248
			v253 = int64(base.Ui64(v249)>>(uint(v239)%64)) + v246*v242
			v260 = v253&v245 + v240*v248
			*(*int64)(unsafe.Add(mBase, uint32(v223)+8)) = v231*v230 + v231*v233 + v240*v242 + int64(base.Ui64(v253)>>(uint(v239)%64)) + int64(base.Ui64(v260)>>(uint(v239)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v223))) = v260<<(uint(v239)%64) | v249&v245
			v272 = v28 + int32(272)
			v273 = int64(0)
			v278 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(296))))
			v279 = v273 - v278
			v286 = int64(32)
			v287 = int64(base.Ui64(v233) >> (uint(v286) % 64))
			v289 = int64(base.Ui64(v279) >> (uint(v286) % 64))
			v292 = int64(4294967295)
			v293 = v233 & v292
			v295 = v279 & v292
			v296 = v293 * v295
			v300 = int64(base.Ui64(v296)>>(uint(v286)%64)) + v293*v289
			v307 = v300&v292 + v287*v295
			*(*int64)(unsafe.Add(mBase, uint32(v272)+8)) = v273*v279 + v273*v233 + v287*v289 + int64(base.Ui64(v300)>>(uint(v286)%64)) + int64(base.Ui64(v307)>>(uint(v286)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v272))) = v307<<(uint(v286)%64) | v296&v292
			v319 = v28 + int32(256)
			v320 = *(*int64)(unsafe.Add(mBase, uint32(v28)+272))
			v327 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(280))))
			v330 = int64(base.Ui64(v320)>>(uint(int64(63))%64)) | v327<<(uint(int64(1))%64)
			v331 = int64(0)
			v337 = int64(32)
			v338 = int64(base.Ui64(v230) >> (uint(v337) % 64))
			v340 = int64(base.Ui64(v330) >> (uint(v337) % 64))
			v343 = int64(4294967295)
			v344 = v230 & v343
			v346 = v330 & v343
			v347 = v344 * v346
			v351 = int64(base.Ui64(v347)>>(uint(v337)%64)) + v344*v340
			v358 = v351&v343 + v338*v346
			*(*int64)(unsafe.Add(mBase, uint32(v319)+8)) = v331*v330 + v331*v230 + v338*v340 + int64(base.Ui64(v351)>>(uint(v337)%64)) + int64(base.Ui64(v358)>>(uint(v337)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v319))) = v358<<(uint(v337)%64) | v347&v343
			v370 = v28 + int32(240)
			v371 = int64(0)
			v377 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(264))))
			v378 = v371 - v377
			v384 = int64(32)
			v385 = int64(base.Ui64(v378) >> (uint(v384) % 64))
			v387 = int64(base.Ui64(v330) >> (uint(v384) % 64))
			v390 = int64(4294967295)
			v391 = v378 & v390
			v393 = v330 & v390
			v394 = v391 * v393
			v398 = int64(base.Ui64(v394)>>(uint(v384)%64)) + v391*v387
			v405 = v398&v390 + v385*v393
			*(*int64)(unsafe.Add(mBase, uint32(v370)+8)) = v371*v330 + v371*v378 + v385*v387 + int64(base.Ui64(v398)>>(uint(v384)%64)) + int64(base.Ui64(v405)>>(uint(v384)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v370))) = v405<<(uint(v384)%64) | v394&v390
			v417 = v28 + int32(224)
			v418 = *(*int64)(unsafe.Add(mBase, uint32(v28)+240))
			v425 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(248))))
			v428 = int64(base.Ui64(v418)>>(uint(int64(63))%64)) | v425<<(uint(int64(1))%64)
			v429 = int64(0)
			v435 = int64(32)
			v436 = int64(base.Ui64(v230) >> (uint(v435) % 64))
			v438 = int64(base.Ui64(v428) >> (uint(v435) % 64))
			v441 = int64(4294967295)
			v442 = v230 & v441
			v444 = v428 & v441
			v445 = v442 * v444
			v449 = int64(base.Ui64(v445)>>(uint(v435)%64)) + v442*v438
			v456 = v449&v441 + v436*v444
			*(*int64)(unsafe.Add(mBase, uint32(v417)+8)) = v429*v428 + v429*v230 + v436*v438 + int64(base.Ui64(v449)>>(uint(v435)%64)) + int64(base.Ui64(v456)>>(uint(v435)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v417))) = v456<<(uint(v435)%64) | v445&v441
			v468 = v28 + int32(208)
			v469 = int64(0)
			v475 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(232))))
			v476 = v469 - v475
			v482 = int64(32)
			v483 = int64(base.Ui64(v476) >> (uint(v482) % 64))
			v485 = int64(base.Ui64(v428) >> (uint(v482) % 64))
			v488 = int64(4294967295)
			v489 = v476 & v488
			v491 = v428 & v488
			v492 = v489 * v491
			v496 = int64(base.Ui64(v492)>>(uint(v482)%64)) + v489*v485
			v503 = v496&v488 + v483*v491
			*(*int64)(unsafe.Add(mBase, uint32(v468)+8)) = v469*v428 + v469*v476 + v483*v485 + int64(base.Ui64(v496)>>(uint(v482)%64)) + int64(base.Ui64(v503)>>(uint(v482)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v468))) = v503<<(uint(v482)%64) | v492&v488
			v515 = v28 + int32(192)
			v516 = *(*int64)(unsafe.Add(mBase, uint32(v28)+208))
			v523 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(216))))
			v526 = int64(base.Ui64(v516)>>(uint(int64(63))%64)) | v523<<(uint(int64(1))%64)
			v527 = int64(0)
			v533 = int64(32)
			v534 = int64(base.Ui64(v230) >> (uint(v533) % 64))
			v536 = int64(base.Ui64(v526) >> (uint(v533) % 64))
			v539 = int64(4294967295)
			v540 = v230 & v539
			v542 = v526 & v539
			v543 = v540 * v542
			v547 = int64(base.Ui64(v543)>>(uint(v533)%64)) + v540*v536
			v554 = v547&v539 + v534*v542
			*(*int64)(unsafe.Add(mBase, uint32(v515)+8)) = v527*v526 + v527*v230 + v534*v536 + int64(base.Ui64(v547)>>(uint(v533)%64)) + int64(base.Ui64(v554)>>(uint(v533)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v515))) = v554<<(uint(v533)%64) | v543&v539
			v566 = v28 + int32(176)
			v567 = int64(0)
			v573 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(200))))
			v574 = v567 - v573
			v580 = int64(32)
			v581 = int64(base.Ui64(v574) >> (uint(v580) % 64))
			v583 = int64(base.Ui64(v526) >> (uint(v580) % 64))
			v586 = int64(4294967295)
			v587 = v574 & v586
			v589 = v526 & v586
			v590 = v587 * v589
			v594 = int64(base.Ui64(v590)>>(uint(v580)%64)) + v587*v583
			v601 = v594&v586 + v581*v589
			*(*int64)(unsafe.Add(mBase, uint32(v566)+8)) = v567*v526 + v567*v574 + v581*v583 + int64(base.Ui64(v594)>>(uint(v580)%64)) + int64(base.Ui64(v601)>>(uint(v580)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v566))) = v601<<(uint(v580)%64) | v590&v586
			v613 = v28 + int32(160)
			v614 = int64(0)
			v615 = *(*int64)(unsafe.Add(mBase, uint32(v28)+176))
			v622 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(184))))
			v627 = int64(base.Ui64(v615)>>(uint(int64(63))%64)) | v622<<(uint(int64(1))%64) + int64(-1)
			v633 = int64(32)
			v634 = int64(base.Ui64(v627) >> (uint(v633) % 64))
			v636 = int64(base.Ui64(v230) >> (uint(v633) % 64))
			v639 = int64(4294967295)
			v640 = v627 & v639
			v642 = v230 & v639
			v643 = v640 * v642
			v647 = int64(base.Ui64(v643)>>(uint(v633)%64)) + v640*v636
			v654 = v647&v639 + v634*v642
			*(*int64)(unsafe.Add(mBase, uint32(v613)+8)) = v614*v230 + v614*v627 + v634*v636 + int64(base.Ui64(v647)>>(uint(v633)%64)) + int64(base.Ui64(v654)>>(uint(v633)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v613))) = v654<<(uint(v633)%64) | v643&v639
			v666 = v28 + int32(144)
			v668 = v216 << (uint(int64(15)) % 64)
			v669 = int64(0)
			v675 = int64(32)
			v676 = int64(base.Ui64(v627) >> (uint(v675) % 64))
			v678 = int64(base.Ui64(v668) >> (uint(v675) % 64))
			v681 = int64(4294967295)
			v682 = v627 & v681
			v684 = v668 & v681
			v685 = v682 * v684
			v689 = int64(base.Ui64(v685)>>(uint(v675)%64)) + v682*v678
			v696 = v689&v681 + v676*v684
			*(*int64)(unsafe.Add(mBase, uint32(v666)+8)) = v669*v668 + v669*v627 + v676*v678 + int64(base.Ui64(v689)>>(uint(v675)%64)) + int64(base.Ui64(v696)>>(uint(v675)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v666))) = v696<<(uint(v675)%64) | v685&v681
			v708 = v28 + int32(112)
			v709 = int64(0)
			v715 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(168))))
			v716 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
			v721 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(152))))
			v722 = v716 + v721
			v730 = v709 - (v715 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v722) < base.Ui64(v716))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(1)) < base.Ui64(v722))))
			v736 = int64(32)
			v737 = int64(base.Ui64(v730) >> (uint(v736) % 64))
			v739 = int64(base.Ui64(v627) >> (uint(v736) % 64))
			v742 = int64(4294967295)
			v743 = v730 & v742
			v745 = v627 & v742
			v746 = v743 * v745
			v750 = int64(base.Ui64(v746)>>(uint(v736)%64)) + v743*v739
			v757 = v750&v742 + v737*v745
			*(*int64)(unsafe.Add(mBase, uint32(v708)+8)) = v709*v627 + v709*v730 + v737*v739 + int64(base.Ui64(v750)>>(uint(v736)%64)) + int64(base.Ui64(v757)>>(uint(v736)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v708))) = v757<<(uint(v736)%64) | v746&v742
			v769 = v28 + int32(128)
			v771 = int64(1) - v722
			v772 = int64(0)
			v778 = int64(32)
			v779 = int64(base.Ui64(v627) >> (uint(v778) % 64))
			v781 = int64(base.Ui64(v771) >> (uint(v778) % 64))
			v784 = int64(4294967295)
			v785 = v627 & v784
			v787 = v771 & v784
			v788 = v785 * v787
			v792 = int64(base.Ui64(v788)>>(uint(v778)%64)) + v785*v781
			v799 = v792&v784 + v779*v787
			*(*int64)(unsafe.Add(mBase, uint32(v769)+8)) = v772*v771 + v772*v627 + v779*v781 + int64(base.Ui64(v792)>>(uint(v778)%64)) + int64(base.Ui64(v799)>>(uint(v778)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v769))) = v799<<(uint(v778)%64) | v788&v784
			v811 = v219 + (v46 - v41)
			v812 = *(*int64)(unsafe.Add(mBase, uint32(v28)+112))
			v813 = int64(1)
			v814 = v812 << (uint(v813) % 64)
			v815 = *(*int64)(unsafe.Add(mBase, uint32(v28)+128))
			v816 = int64(63)
			v822 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(136))))
			v826 = v814 + (int64(base.Ui64(v815)>>(uint(v816)%64)) | v822<<(uint(v813)%64))
			v828 = v826 + int64(-13927)
			v829 = int64(32)
			v830 = int64(base.Ui64(v828) >> (uint(v829) % 64))
			v832 = v218 | int64(281474976710656)
			v834 = v832 << (uint(v813) % 64)
			v836 = int64(base.Ui64(v834) >> (uint(v829) % 64))
			v837 = v830 * v836
			v839 = v214 << (uint(v813) % 64)
			v841 = int64(base.Ui64(v839) >> (uint(v829) % 64))
			v846 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(120))))
			v862 = v846<<(uint(v813)%64) | int64(base.Ui64(v812)>>(uint(v816)%64)) + int64(base.Ui64(v822)>>(uint(v816)%64)) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v826) < base.Ui64(v814))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v828) < base.Ui64(v826))) + int64(-1)
			v864 = int64(base.Ui64(v862) >> (uint(v829) % 64))
			v866 = v837 + v841*v864
			v869 = int64(4294967295)
			v870 = v862 & v869
			v872 = int64(base.Ui64(v214) >> (uint(v816) % 64))
			v877 = (v872 | v218<<(uint(v813)%64)) & v869
			v879 = v866 + v870*v877
			v885 = v870 * v836
			v887 = v885 + v877*v864
			v898 = v879 + v887<<(uint(v829)%64)
			v903 = v828 & v869
			v904 = v903 * v877
			v906 = v904 + v830*v841
			v910 = v839 & int64(4294967294)
			v912 = v906 + v870*v910
			v916 = v898 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v906) < base.Ui64(v904))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v912) < base.Ui64(v906))))
			v920 = v903 * v836
			v922 = v920 + v910*v864
			v924 = v922 + v830*v877
			v926 = v924 + v870*v841
			v940 = v916 + (int64(base.Ui64(v926)>>(uint(v829)%64)) | (base.I64_extend_i32_u(base.B2i32(base.Ui64(v922) < base.Ui64(v920)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v924) < base.Ui64(v922)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v926) < base.Ui64(v924))))<<(uint(v829)%64))
			v944 = v830 * v910
			v946 = v944 + v903*v841
			v954 = v912 + (int64(base.Ui64(v946)>>(uint(v829)%64)) | base.I64_extend_i32_u(base.B2i32(base.Ui64(v946) < base.Ui64(v944)))<<(uint(v829)%64))
			v963 = v940 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v954) < base.Ui64(v912))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v954+v926<<(uint(v829)%64)) < base.Ui64(v954))))
			v966 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v866) < base.Ui64(v837))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v879) < base.Ui64(v866))) + v864*v836 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v887) < base.Ui64(v885)))<<(uint(v829)%64) | int64(base.Ui64(v887)>>(uint(v829)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v898) < base.Ui64(v879))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v916) < base.Ui64(v898))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v940) < base.Ui64(v916))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v963) < base.Ui64(v940)))
			if base.Ui64(int64(562949953421311)) < base.Ui64(v966) {
				v1026 = v28 + int32(96)
				v1027 = int64(1)
				v1031 = int64(base.Ui64(v963)>>(uint(v1027)%64)) | v966<<(uint(int64(63))%64)
				v1033 = int64(base.Ui64(v966) >> (uint(v1027) % 64))
				v1038 = int64(32)
				v1039 = int64(base.Ui64(v216) >> (uint(v1038) % 64))
				v1041 = int64(base.Ui64(v1031) >> (uint(v1038) % 64))
				v1044 = int64(4294967295)
				v1045 = v216 & v1044
				v1047 = v1031 & v1044
				v1048 = v1045 * v1047
				v1052 = int64(base.Ui64(v1048)>>(uint(v1038)%64)) + v1045*v1041
				v1059 = v1052&v1044 + v1039*v1047
				*(*int64)(unsafe.Add(mBase, uint32(v1026)+8)) = v227*v1031 + v1033*v216 + v1039*v1041 + int64(base.Ui64(v1052)>>(uint(v1038)%64)) + int64(base.Ui64(v1059)>>(uint(v1038)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v1026))) = v1059<<(uint(v1038)%64) | v1048&v1044
				v1076 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(104))))
				v1078 = *(*int64)(unsafe.Add(mBase, uint32(v28)+96))
				v1079 = int64(0)
				v1088 = v1031
				v1089 = v1033
				v1090 = v214<<(uint(int64(48))%64) - v1076 - base.I64_extend_i32_u(base.B2i32(v1078 != v1079))
				v1091 = v1079 - v1078
				v1092 = v811 + int32(16383)
				v1093 = v832
				v1094 = v214
			} else {
				v971 = v28 + int32(80)
				v976 = int64(32)
				v977 = int64(base.Ui64(v216) >> (uint(v976) % 64))
				v979 = int64(base.Ui64(v963) >> (uint(v976) % 64))
				v982 = int64(4294967295)
				v983 = v216 & v982
				v985 = v963 & v982
				v986 = v983 * v985
				v990 = int64(base.Ui64(v986)>>(uint(v976)%64)) + v983*v979
				v997 = v990&v982 + v977*v985
				*(*int64)(unsafe.Add(mBase, uint32(v971)+8)) = v227*v963 + v966*v216 + v977*v979 + int64(base.Ui64(v990)>>(uint(v976)%64)) + int64(base.Ui64(v997)>>(uint(v976)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v971))) = v997<<(uint(v976)%64) | v986&v982
				v1014 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(88))))
				v1016 = *(*int64)(unsafe.Add(mBase, uint32(v28)+80))
				v1017 = int64(0)
				v1088 = v963
				v1089 = v966
				v1090 = v214<<(uint(int64(49))%64) - v1014 - base.I64_extend_i32_u(base.B2i32(v1016 != v1017))
				v1091 = v1017 - v1016
				v1092 = v811 + int32(16382)
				v1093 = v834 | v872
				v1094 = v839
			}
			if v1092 < int32(32767) {
				if v1092 < int32(1) {
					if int32(-113) < v1092 {
						v1118 = int32(64)
						v1119 = v28 + v1118
						v1121 = int32(1) - v1092
						if v1121&v1118 == int32(0) {
							if v1121 == int32(0) {
								v1142 = v1088
								v1143 = v1089
							} else {
								v1138 = base.I64_extend_i32_u(v1121)
								v1142 = v1089<<(uint(base.I64_extend_i32_u(int32(64)-v1121))%64) | int64(base.Ui64(v1088)>>(uint(v1138)%64))
								v1143 = int64(base.Ui64(v1089) >> (uint(v1138) % 64))
							}
						} else {
							v1142 = int64(base.Ui64(v1089) >> (uint(base.I64_extend_i32_u(v1121+int32(-64))) % 64))
							v1143 = int64(0)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v1119))) = v1142
						*(*int64)(unsafe.Add(mBase, uint32(v1119)+8)) = v1143
						v1148 = v28 + int32(48)
						v1150 = v1092 + int32(112)
						if v1150&int32(64) == int32(0) {
							if v1150 == int32(0) {
								v1171 = v1094
								v1172 = v1093
							} else {
								v1167 = base.I64_extend_i32_u(v1150)
								v1171 = v1094 << (uint(v1167) % 64)
								v1172 = int64(base.Ui64(v1094)>>(uint(base.I64_extend_i32_u(int32(64)-v1150))%64)) | v1093<<(uint(v1167)%64)
							}
						} else {
							v1171 = int64(0)
							v1172 = v1094 << (uint(base.I64_extend_i32_u(v1092+int32(48))) % 64)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v1148))) = v1171
						*(*int64)(unsafe.Add(mBase, uint32(v1148)+8)) = v1172
						v1177 = v28 + int32(32)
						v1178 = *(*int64)(unsafe.Add(mBase, uint32(v28)+64))
						v1183 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(72))))
						v1188 = int64(32)
						v1189 = int64(base.Ui64(v1178) >> (uint(v1188) % 64))
						v1191 = int64(base.Ui64(v216) >> (uint(v1188) % 64))
						v1194 = int64(4294967295)
						v1195 = v1178 & v1194
						v1197 = v216 & v1194
						v1198 = v1195 * v1197
						v1202 = int64(base.Ui64(v1198)>>(uint(v1188)%64)) + v1195*v1191
						v1209 = v1202&v1194 + v1189*v1197
						*(*int64)(unsafe.Add(mBase, uint32(v1177)+8)) = v1183*v216 + v227*v1178 + v1189*v1191 + int64(base.Ui64(v1202)>>(uint(v1188)%64)) + int64(base.Ui64(v1209)>>(uint(v1188)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v1177))) = v1209<<(uint(v1188)%64) | v1198&v1194
						v1224 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(56))))
						v1229 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(40))))
						v1230 = int64(1)
						v1232 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
						v1237 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
						v1239 = v1232 << (uint(v1230) % 64)
						v1244 = v1224 - (v1229<<(uint(v1230)%64) | int64(base.Ui64(v1232)>>(uint(int64(63))%64))) - base.I64_extend_i32_u(base.B2i32(base.Ui64(v1237) < base.Ui64(v1239)))
						v1245 = v1178
						v1246 = v1237 - v1239
						v1247 = v1183
						v1250 = v28 + int32(16)
						v1251 = int64(3)
						v1252 = int64(0)
						v1257 = int64(32)
						v1260 = int64(base.Ui64(v216) >> (uint(v1257) % 64))
						v1263 = int64(4294967295)
						v1266 = v216 & v1263
						v1267 = v1251 * v1266
						v1271 = int64(base.Ui64(v1267)>>(uint(v1257)%64)) + v1251*v1260
						v1278 = v1271&v1263 + v1252*v1266
						*(*int64)(unsafe.Add(mBase, uint32(v1250)+8)) = v1252*v216 + v227*v1251 + v1252*v1260 + int64(base.Ui64(v1271)>>(uint(v1257)%64)) + int64(base.Ui64(v1278)>>(uint(v1257)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v1250))) = v1278<<(uint(v1257)%64) | v1267&v1263
						v1289 = int64(5)
						v1290 = int64(0)
						v1295 = int64(32)
						v1298 = int64(base.Ui64(v216) >> (uint(v1295) % 64))
						v1301 = int64(4294967295)
						v1304 = v216 & v1301
						v1305 = v1289 * v1304
						v1309 = int64(base.Ui64(v1305)>>(uint(v1295)%64)) + v1289*v1298
						v1316 = v1309&v1301 + v1290*v1304
						*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v1290*v216 + v227*v1289 + v1290*v1298 + int64(base.Ui64(v1309)>>(uint(v1295)%64)) + int64(base.Ui64(v1316)>>(uint(v1295)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1316<<(uint(v1295)%64) | v1305&v1301
						v1328 = v1245 & int64(1)
						v1329 = v1328 + v1246
						v1333 = v1244 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1329) < base.Ui64(v1328)))
						if v1333 == v227 {
							v1336 = base.B2i32(base.Ui64(v216) < base.Ui64(v1329))
						} else {
							v1336 = base.B2i32(base.Ui64(v227) < base.Ui64(v1333))
						}
						v1338 = v1245 + base.I64_extend_i32_u(v1336)
						v1341 = v1247 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1338) < base.Ui64(v1245)))
						v1344 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
						v1350 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(24))))
						if v1333 == v1350 {
							v1353 = base.B2i32(base.Ui64(v1344) < base.Ui64(v1329))
						} else {
							v1353 = base.B2i32(base.Ui64(v1350) < base.Ui64(v1333))
						}
						v1356 = v1338 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1341) < base.Ui64(int64(9223090561878065152)))&v1353)
						v1359 = v1341 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1356) < base.Ui64(v1338)))
						v1362 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
						v1366 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(8))))
						if v1333 == v1366 {
							v1369 = base.B2i32(base.Ui64(v1362) < base.Ui64(v1329))
						} else {
							v1369 = base.B2i32(base.Ui64(v1366) < base.Ui64(v1333))
						}
						v1372 = v1356 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1359) < base.Ui64(int64(9223090561878065152)))&v1369)
						v1377 = v1372
						v1383 = v1359 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1372) < base.Ui64(v1356))) | v36
					} else {
						v1377 = int64(0)
						v1383 = v36
					}
				} else {
					v1102 = int64(1)
					v1244 = v1090<<(uint(v1102)%64) | int64(base.Ui64(v1091)>>(uint(int64(63))%64))
					v1245 = v1088
					v1246 = v1091 << (uint(v1102) % 64)
					v1247 = base.I64_extend_i32_u(v1092)<<(uint(int64(48))%64) | v1089&int64(281474976710655)
					v1250 = v28 + int32(16)
					v1251 = int64(3)
					v1252 = int64(0)
					v1257 = int64(32)
					v1260 = int64(base.Ui64(v216) >> (uint(v1257) % 64))
					v1263 = int64(4294967295)
					v1266 = v216 & v1263
					v1267 = v1251 * v1266
					v1271 = int64(base.Ui64(v1267)>>(uint(v1257)%64)) + v1251*v1260
					v1278 = v1271&v1263 + v1252*v1266
					*(*int64)(unsafe.Add(mBase, uint32(v1250)+8)) = v1252*v216 + v227*v1251 + v1252*v1260 + int64(base.Ui64(v1271)>>(uint(v1257)%64)) + int64(base.Ui64(v1278)>>(uint(v1257)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v1250))) = v1278<<(uint(v1257)%64) | v1267&v1263
					v1289 = int64(5)
					v1290 = int64(0)
					v1295 = int64(32)
					v1298 = int64(base.Ui64(v216) >> (uint(v1295) % 64))
					v1301 = int64(4294967295)
					v1304 = v216 & v1301
					v1305 = v1289 * v1304
					v1309 = int64(base.Ui64(v1305)>>(uint(v1295)%64)) + v1289*v1298
					v1316 = v1309&v1301 + v1290*v1304
					*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v1290*v216 + v227*v1289 + v1290*v1298 + int64(base.Ui64(v1309)>>(uint(v1295)%64)) + int64(base.Ui64(v1316)>>(uint(v1295)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1316<<(uint(v1295)%64) | v1305&v1301
					v1328 = v1245 & int64(1)
					v1329 = v1328 + v1246
					v1333 = v1244 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1329) < base.Ui64(v1328)))
					if v1333 == v227 {
						v1336 = base.B2i32(base.Ui64(v216) < base.Ui64(v1329))
					} else {
						v1336 = base.B2i32(base.Ui64(v227) < base.Ui64(v1333))
					}
					v1338 = v1245 + base.I64_extend_i32_u(v1336)
					v1341 = v1247 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1338) < base.Ui64(v1245)))
					v1344 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
					v1350 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(24))))
					if v1333 == v1350 {
						v1353 = base.B2i32(base.Ui64(v1344) < base.Ui64(v1329))
					} else {
						v1353 = base.B2i32(base.Ui64(v1350) < base.Ui64(v1333))
					}
					v1356 = v1338 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1341) < base.Ui64(int64(9223090561878065152)))&v1353)
					v1359 = v1341 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1356) < base.Ui64(v1338)))
					v1362 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
					v1366 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(8))))
					if v1333 == v1366 {
						v1369 = base.B2i32(base.Ui64(v1362) < base.Ui64(v1329))
					} else {
						v1369 = base.B2i32(base.Ui64(v1366) < base.Ui64(v1333))
					}
					v1372 = v1356 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1359) < base.Ui64(int64(9223090561878065152)))&v1369)
					v1377 = v1372
					v1383 = v1359 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1372) < base.Ui64(v1356))) | v36
				}
			} else {
				v1377 = int64(0)
				v1383 = v36 | int64(9223090561878065152)
			}
		} else {
			v60 = l2 & int64(9223372036854775807)
			v61 = int64(9223090561878065152)
			if v60 == v61 {
				v65 = base.B2i32(l1 == int64(0))
			} else {
				v65 = base.B2i32(base.Ui64(v60) < base.Ui64(v61))
			}
			if v65 != 0 {
				v71 = l4 & int64(9223372036854775807)
				v72 = int64(9223090561878065152)
				if v71 == v72 {
					v76 = base.B2i32(l3 == int64(0))
				} else {
					v76 = base.B2i32(base.Ui64(v71) < base.Ui64(v72))
				}
				if v76 != 0 {
					if l1|(v60^int64(9223090561878065152)) != int64(0) {
						if l3|(v71^int64(9223090561878065152)) != int64(0) {
							if l1|v60 != int64(0) {
								if l3|v71 != int64(0) {
									if base.Ui64(int64(281474976710655)) < base.Ui64(v60) {
										v164 = l1
										v165 = v33
										v166 = int32(0)
									} else {
										v121 = v28 + int32(320)
										v123 = base.B2i32(v33 == int64(0))
										if v33 == int64(0) {
											v124 = l1
										} else {
											v124 = v33
										}
										v130 = base.I32_wrap_i64(base.I64_clz(v124) + base.I64_extend_i32_u(v123<<(uint(int32(6))%32)))
										v132 = v130 + int32(-15)
										if v132&int32(64) == int32(0) {
											if v132 == int32(0) {
												v153 = l1
												v154 = v33
											} else {
												v149 = base.I64_extend_i32_u(v132)
												v153 = l1 << (uint(v149) % 64)
												v154 = int64(base.Ui64(l1)>>(uint(base.I64_extend_i32_u(int32(64)-v132))%64)) | v33<<(uint(v149)%64)
											}
										} else {
											v153 = int64(0)
											v154 = l1 << (uint(base.I64_extend_i32_u(v130+int32(-79))) % 64)
										}
										*(*int64)(unsafe.Add(mBase, uint32(v121))) = v153
										*(*int64)(unsafe.Add(mBase, uint32(v121)+8)) = v154
										v162 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(328))))
										v163 = *(*int64)(unsafe.Add(mBase, uint32(v28)+320))
										v164 = v163
										v165 = v162
										v166 = int32(16) - v130
									}
									if base.Ui64(int64(281474976710655)) < base.Ui64(v71) {
										v214 = v164
										v216 = l3
										v217 = v31
										v218 = v165
										v219 = v166
									} else {
										v170 = v28 + int32(304)
										v172 = base.B2i32(v31 == int64(0))
										if v31 == int64(0) {
											v173 = l3
										} else {
											v173 = v31
										}
										v179 = base.I32_wrap_i64(base.I64_clz(v173) + base.I64_extend_i32_u(v172<<(uint(int32(6))%32)))
										v181 = v179 + int32(-15)
										if v181&int32(64) == int32(0) {
											if v181 == int32(0) {
												v202 = l3
												v203 = v31
											} else {
												v198 = base.I64_extend_i32_u(v181)
												v202 = l3 << (uint(v198) % 64)
												v203 = int64(base.Ui64(l3)>>(uint(base.I64_extend_i32_u(int32(64)-v181))%64)) | v31<<(uint(v198)%64)
											}
										} else {
											v202 = int64(0)
											v203 = l3 << (uint(base.I64_extend_i32_u(v179+int32(-79))) % 64)
										}
										*(*int64)(unsafe.Add(mBase, uint32(v170))) = v202
										*(*int64)(unsafe.Add(mBase, uint32(v170)+8)) = v203
										v212 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(312))))
										v213 = *(*int64)(unsafe.Add(mBase, uint32(v28)+304))
										v214 = v164
										v216 = v213
										v217 = v212
										v218 = v165
										v219 = v179 + v166 + int32(-16)
									}
									v223 = v28 + int32(288)
									v227 = v217 | int64(281474976710656)
									v230 = int64(base.Ui64(v216)>>(uint(int64(49))%64)) | v227<<(uint(int64(15))%64)
									v231 = int64(0)
									v233 = int64(8432131802713292800) - v230
									v239 = int64(32)
									v240 = int64(base.Ui64(v233) >> (uint(v239) % 64))
									v242 = int64(base.Ui64(v230) >> (uint(v239) % 64))
									v245 = int64(4294967295)
									v246 = v233 & v245
									v248 = v230 & v245
									v249 = v246 * v248
									v253 = int64(base.Ui64(v249)>>(uint(v239)%64)) + v246*v242
									v260 = v253&v245 + v240*v248
									*(*int64)(unsafe.Add(mBase, uint32(v223)+8)) = v231*v230 + v231*v233 + v240*v242 + int64(base.Ui64(v253)>>(uint(v239)%64)) + int64(base.Ui64(v260)>>(uint(v239)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v223))) = v260<<(uint(v239)%64) | v249&v245
									v272 = v28 + int32(272)
									v273 = int64(0)
									v278 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(296))))
									v279 = v273 - v278
									v286 = int64(32)
									v287 = int64(base.Ui64(v233) >> (uint(v286) % 64))
									v289 = int64(base.Ui64(v279) >> (uint(v286) % 64))
									v292 = int64(4294967295)
									v293 = v233 & v292
									v295 = v279 & v292
									v296 = v293 * v295
									v300 = int64(base.Ui64(v296)>>(uint(v286)%64)) + v293*v289
									v307 = v300&v292 + v287*v295
									*(*int64)(unsafe.Add(mBase, uint32(v272)+8)) = v273*v279 + v273*v233 + v287*v289 + int64(base.Ui64(v300)>>(uint(v286)%64)) + int64(base.Ui64(v307)>>(uint(v286)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v272))) = v307<<(uint(v286)%64) | v296&v292
									v319 = v28 + int32(256)
									v320 = *(*int64)(unsafe.Add(mBase, uint32(v28)+272))
									v327 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(280))))
									v330 = int64(base.Ui64(v320)>>(uint(int64(63))%64)) | v327<<(uint(int64(1))%64)
									v331 = int64(0)
									v337 = int64(32)
									v338 = int64(base.Ui64(v230) >> (uint(v337) % 64))
									v340 = int64(base.Ui64(v330) >> (uint(v337) % 64))
									v343 = int64(4294967295)
									v344 = v230 & v343
									v346 = v330 & v343
									v347 = v344 * v346
									v351 = int64(base.Ui64(v347)>>(uint(v337)%64)) + v344*v340
									v358 = v351&v343 + v338*v346
									*(*int64)(unsafe.Add(mBase, uint32(v319)+8)) = v331*v330 + v331*v230 + v338*v340 + int64(base.Ui64(v351)>>(uint(v337)%64)) + int64(base.Ui64(v358)>>(uint(v337)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v319))) = v358<<(uint(v337)%64) | v347&v343
									v370 = v28 + int32(240)
									v371 = int64(0)
									v377 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(264))))
									v378 = v371 - v377
									v384 = int64(32)
									v385 = int64(base.Ui64(v378) >> (uint(v384) % 64))
									v387 = int64(base.Ui64(v330) >> (uint(v384) % 64))
									v390 = int64(4294967295)
									v391 = v378 & v390
									v393 = v330 & v390
									v394 = v391 * v393
									v398 = int64(base.Ui64(v394)>>(uint(v384)%64)) + v391*v387
									v405 = v398&v390 + v385*v393
									*(*int64)(unsafe.Add(mBase, uint32(v370)+8)) = v371*v330 + v371*v378 + v385*v387 + int64(base.Ui64(v398)>>(uint(v384)%64)) + int64(base.Ui64(v405)>>(uint(v384)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v370))) = v405<<(uint(v384)%64) | v394&v390
									v417 = v28 + int32(224)
									v418 = *(*int64)(unsafe.Add(mBase, uint32(v28)+240))
									v425 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(248))))
									v428 = int64(base.Ui64(v418)>>(uint(int64(63))%64)) | v425<<(uint(int64(1))%64)
									v429 = int64(0)
									v435 = int64(32)
									v436 = int64(base.Ui64(v230) >> (uint(v435) % 64))
									v438 = int64(base.Ui64(v428) >> (uint(v435) % 64))
									v441 = int64(4294967295)
									v442 = v230 & v441
									v444 = v428 & v441
									v445 = v442 * v444
									v449 = int64(base.Ui64(v445)>>(uint(v435)%64)) + v442*v438
									v456 = v449&v441 + v436*v444
									*(*int64)(unsafe.Add(mBase, uint32(v417)+8)) = v429*v428 + v429*v230 + v436*v438 + int64(base.Ui64(v449)>>(uint(v435)%64)) + int64(base.Ui64(v456)>>(uint(v435)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v417))) = v456<<(uint(v435)%64) | v445&v441
									v468 = v28 + int32(208)
									v469 = int64(0)
									v475 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(232))))
									v476 = v469 - v475
									v482 = int64(32)
									v483 = int64(base.Ui64(v476) >> (uint(v482) % 64))
									v485 = int64(base.Ui64(v428) >> (uint(v482) % 64))
									v488 = int64(4294967295)
									v489 = v476 & v488
									v491 = v428 & v488
									v492 = v489 * v491
									v496 = int64(base.Ui64(v492)>>(uint(v482)%64)) + v489*v485
									v503 = v496&v488 + v483*v491
									*(*int64)(unsafe.Add(mBase, uint32(v468)+8)) = v469*v428 + v469*v476 + v483*v485 + int64(base.Ui64(v496)>>(uint(v482)%64)) + int64(base.Ui64(v503)>>(uint(v482)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v468))) = v503<<(uint(v482)%64) | v492&v488
									v515 = v28 + int32(192)
									v516 = *(*int64)(unsafe.Add(mBase, uint32(v28)+208))
									v523 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(216))))
									v526 = int64(base.Ui64(v516)>>(uint(int64(63))%64)) | v523<<(uint(int64(1))%64)
									v527 = int64(0)
									v533 = int64(32)
									v534 = int64(base.Ui64(v230) >> (uint(v533) % 64))
									v536 = int64(base.Ui64(v526) >> (uint(v533) % 64))
									v539 = int64(4294967295)
									v540 = v230 & v539
									v542 = v526 & v539
									v543 = v540 * v542
									v547 = int64(base.Ui64(v543)>>(uint(v533)%64)) + v540*v536
									v554 = v547&v539 + v534*v542
									*(*int64)(unsafe.Add(mBase, uint32(v515)+8)) = v527*v526 + v527*v230 + v534*v536 + int64(base.Ui64(v547)>>(uint(v533)%64)) + int64(base.Ui64(v554)>>(uint(v533)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v515))) = v554<<(uint(v533)%64) | v543&v539
									v566 = v28 + int32(176)
									v567 = int64(0)
									v573 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(200))))
									v574 = v567 - v573
									v580 = int64(32)
									v581 = int64(base.Ui64(v574) >> (uint(v580) % 64))
									v583 = int64(base.Ui64(v526) >> (uint(v580) % 64))
									v586 = int64(4294967295)
									v587 = v574 & v586
									v589 = v526 & v586
									v590 = v587 * v589
									v594 = int64(base.Ui64(v590)>>(uint(v580)%64)) + v587*v583
									v601 = v594&v586 + v581*v589
									*(*int64)(unsafe.Add(mBase, uint32(v566)+8)) = v567*v526 + v567*v574 + v581*v583 + int64(base.Ui64(v594)>>(uint(v580)%64)) + int64(base.Ui64(v601)>>(uint(v580)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v566))) = v601<<(uint(v580)%64) | v590&v586
									v613 = v28 + int32(160)
									v614 = int64(0)
									v615 = *(*int64)(unsafe.Add(mBase, uint32(v28)+176))
									v622 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(184))))
									v627 = int64(base.Ui64(v615)>>(uint(int64(63))%64)) | v622<<(uint(int64(1))%64) + int64(-1)
									v633 = int64(32)
									v634 = int64(base.Ui64(v627) >> (uint(v633) % 64))
									v636 = int64(base.Ui64(v230) >> (uint(v633) % 64))
									v639 = int64(4294967295)
									v640 = v627 & v639
									v642 = v230 & v639
									v643 = v640 * v642
									v647 = int64(base.Ui64(v643)>>(uint(v633)%64)) + v640*v636
									v654 = v647&v639 + v634*v642
									*(*int64)(unsafe.Add(mBase, uint32(v613)+8)) = v614*v230 + v614*v627 + v634*v636 + int64(base.Ui64(v647)>>(uint(v633)%64)) + int64(base.Ui64(v654)>>(uint(v633)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v613))) = v654<<(uint(v633)%64) | v643&v639
									v666 = v28 + int32(144)
									v668 = v216 << (uint(int64(15)) % 64)
									v669 = int64(0)
									v675 = int64(32)
									v676 = int64(base.Ui64(v627) >> (uint(v675) % 64))
									v678 = int64(base.Ui64(v668) >> (uint(v675) % 64))
									v681 = int64(4294967295)
									v682 = v627 & v681
									v684 = v668 & v681
									v685 = v682 * v684
									v689 = int64(base.Ui64(v685)>>(uint(v675)%64)) + v682*v678
									v696 = v689&v681 + v676*v684
									*(*int64)(unsafe.Add(mBase, uint32(v666)+8)) = v669*v668 + v669*v627 + v676*v678 + int64(base.Ui64(v689)>>(uint(v675)%64)) + int64(base.Ui64(v696)>>(uint(v675)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v666))) = v696<<(uint(v675)%64) | v685&v681
									v708 = v28 + int32(112)
									v709 = int64(0)
									v715 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(168))))
									v716 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
									v721 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(152))))
									v722 = v716 + v721
									v730 = v709 - (v715 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v722) < base.Ui64(v716))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(1)) < base.Ui64(v722))))
									v736 = int64(32)
									v737 = int64(base.Ui64(v730) >> (uint(v736) % 64))
									v739 = int64(base.Ui64(v627) >> (uint(v736) % 64))
									v742 = int64(4294967295)
									v743 = v730 & v742
									v745 = v627 & v742
									v746 = v743 * v745
									v750 = int64(base.Ui64(v746)>>(uint(v736)%64)) + v743*v739
									v757 = v750&v742 + v737*v745
									*(*int64)(unsafe.Add(mBase, uint32(v708)+8)) = v709*v627 + v709*v730 + v737*v739 + int64(base.Ui64(v750)>>(uint(v736)%64)) + int64(base.Ui64(v757)>>(uint(v736)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v708))) = v757<<(uint(v736)%64) | v746&v742
									v769 = v28 + int32(128)
									v771 = int64(1) - v722
									v772 = int64(0)
									v778 = int64(32)
									v779 = int64(base.Ui64(v627) >> (uint(v778) % 64))
									v781 = int64(base.Ui64(v771) >> (uint(v778) % 64))
									v784 = int64(4294967295)
									v785 = v627 & v784
									v787 = v771 & v784
									v788 = v785 * v787
									v792 = int64(base.Ui64(v788)>>(uint(v778)%64)) + v785*v781
									v799 = v792&v784 + v779*v787
									*(*int64)(unsafe.Add(mBase, uint32(v769)+8)) = v772*v771 + v772*v627 + v779*v781 + int64(base.Ui64(v792)>>(uint(v778)%64)) + int64(base.Ui64(v799)>>(uint(v778)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v769))) = v799<<(uint(v778)%64) | v788&v784
									v811 = v219 + (v46 - v41)
									v812 = *(*int64)(unsafe.Add(mBase, uint32(v28)+112))
									v813 = int64(1)
									v814 = v812 << (uint(v813) % 64)
									v815 = *(*int64)(unsafe.Add(mBase, uint32(v28)+128))
									v816 = int64(63)
									v822 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(136))))
									v826 = v814 + (int64(base.Ui64(v815)>>(uint(v816)%64)) | v822<<(uint(v813)%64))
									v828 = v826 + int64(-13927)
									v829 = int64(32)
									v830 = int64(base.Ui64(v828) >> (uint(v829) % 64))
									v832 = v218 | int64(281474976710656)
									v834 = v832 << (uint(v813) % 64)
									v836 = int64(base.Ui64(v834) >> (uint(v829) % 64))
									v837 = v830 * v836
									v839 = v214 << (uint(v813) % 64)
									v841 = int64(base.Ui64(v839) >> (uint(v829) % 64))
									v846 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(120))))
									v862 = v846<<(uint(v813)%64) | int64(base.Ui64(v812)>>(uint(v816)%64)) + int64(base.Ui64(v822)>>(uint(v816)%64)) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v826) < base.Ui64(v814))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v828) < base.Ui64(v826))) + int64(-1)
									v864 = int64(base.Ui64(v862) >> (uint(v829) % 64))
									v866 = v837 + v841*v864
									v869 = int64(4294967295)
									v870 = v862 & v869
									v872 = int64(base.Ui64(v214) >> (uint(v816) % 64))
									v877 = (v872 | v218<<(uint(v813)%64)) & v869
									v879 = v866 + v870*v877
									v885 = v870 * v836
									v887 = v885 + v877*v864
									v898 = v879 + v887<<(uint(v829)%64)
									v903 = v828 & v869
									v904 = v903 * v877
									v906 = v904 + v830*v841
									v910 = v839 & int64(4294967294)
									v912 = v906 + v870*v910
									v916 = v898 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v906) < base.Ui64(v904))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v912) < base.Ui64(v906))))
									v920 = v903 * v836
									v922 = v920 + v910*v864
									v924 = v922 + v830*v877
									v926 = v924 + v870*v841
									v940 = v916 + (int64(base.Ui64(v926)>>(uint(v829)%64)) | (base.I64_extend_i32_u(base.B2i32(base.Ui64(v922) < base.Ui64(v920)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v924) < base.Ui64(v922)))+base.I64_extend_i32_u(base.B2i32(base.Ui64(v926) < base.Ui64(v924))))<<(uint(v829)%64))
									v944 = v830 * v910
									v946 = v944 + v903*v841
									v954 = v912 + (int64(base.Ui64(v946)>>(uint(v829)%64)) | base.I64_extend_i32_u(base.B2i32(base.Ui64(v946) < base.Ui64(v944)))<<(uint(v829)%64))
									v963 = v940 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v954) < base.Ui64(v912))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v954+v926<<(uint(v829)%64)) < base.Ui64(v954))))
									v966 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v866) < base.Ui64(v837))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v879) < base.Ui64(v866))) + v864*v836 + (base.I64_extend_i32_u(base.B2i32(base.Ui64(v887) < base.Ui64(v885)))<<(uint(v829)%64) | int64(base.Ui64(v887)>>(uint(v829)%64))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v898) < base.Ui64(v879))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v916) < base.Ui64(v898))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v940) < base.Ui64(v916))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v963) < base.Ui64(v940)))
									if base.Ui64(int64(562949953421311)) < base.Ui64(v966) {
										v1026 = v28 + int32(96)
										v1027 = int64(1)
										v1031 = int64(base.Ui64(v963)>>(uint(v1027)%64)) | v966<<(uint(int64(63))%64)
										v1033 = int64(base.Ui64(v966) >> (uint(v1027) % 64))
										v1038 = int64(32)
										v1039 = int64(base.Ui64(v216) >> (uint(v1038) % 64))
										v1041 = int64(base.Ui64(v1031) >> (uint(v1038) % 64))
										v1044 = int64(4294967295)
										v1045 = v216 & v1044
										v1047 = v1031 & v1044
										v1048 = v1045 * v1047
										v1052 = int64(base.Ui64(v1048)>>(uint(v1038)%64)) + v1045*v1041
										v1059 = v1052&v1044 + v1039*v1047
										*(*int64)(unsafe.Add(mBase, uint32(v1026)+8)) = v227*v1031 + v1033*v216 + v1039*v1041 + int64(base.Ui64(v1052)>>(uint(v1038)%64)) + int64(base.Ui64(v1059)>>(uint(v1038)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v1026))) = v1059<<(uint(v1038)%64) | v1048&v1044
										v1076 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(104))))
										v1078 = *(*int64)(unsafe.Add(mBase, uint32(v28)+96))
										v1079 = int64(0)
										v1088 = v1031
										v1089 = v1033
										v1090 = v214<<(uint(int64(48))%64) - v1076 - base.I64_extend_i32_u(base.B2i32(v1078 != v1079))
										v1091 = v1079 - v1078
										v1092 = v811 + int32(16383)
										v1093 = v832
										v1094 = v214
									} else {
										v971 = v28 + int32(80)
										v976 = int64(32)
										v977 = int64(base.Ui64(v216) >> (uint(v976) % 64))
										v979 = int64(base.Ui64(v963) >> (uint(v976) % 64))
										v982 = int64(4294967295)
										v983 = v216 & v982
										v985 = v963 & v982
										v986 = v983 * v985
										v990 = int64(base.Ui64(v986)>>(uint(v976)%64)) + v983*v979
										v997 = v990&v982 + v977*v985
										*(*int64)(unsafe.Add(mBase, uint32(v971)+8)) = v227*v963 + v966*v216 + v977*v979 + int64(base.Ui64(v990)>>(uint(v976)%64)) + int64(base.Ui64(v997)>>(uint(v976)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v971))) = v997<<(uint(v976)%64) | v986&v982
										v1014 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(88))))
										v1016 = *(*int64)(unsafe.Add(mBase, uint32(v28)+80))
										v1017 = int64(0)
										v1088 = v963
										v1089 = v966
										v1090 = v214<<(uint(int64(49))%64) - v1014 - base.I64_extend_i32_u(base.B2i32(v1016 != v1017))
										v1091 = v1017 - v1016
										v1092 = v811 + int32(16382)
										v1093 = v834 | v872
										v1094 = v839
									}
									if v1092 < int32(32767) {
										if v1092 < int32(1) {
											if int32(-113) < v1092 {
												v1118 = int32(64)
												v1119 = v28 + v1118
												v1121 = int32(1) - v1092
												if v1121&v1118 == int32(0) {
													if v1121 == int32(0) {
														v1142 = v1088
														v1143 = v1089
													} else {
														v1138 = base.I64_extend_i32_u(v1121)
														v1142 = v1089<<(uint(base.I64_extend_i32_u(int32(64)-v1121))%64) | int64(base.Ui64(v1088)>>(uint(v1138)%64))
														v1143 = int64(base.Ui64(v1089) >> (uint(v1138) % 64))
													}
												} else {
													v1142 = int64(base.Ui64(v1089) >> (uint(base.I64_extend_i32_u(v1121+int32(-64))) % 64))
													v1143 = int64(0)
												}
												*(*int64)(unsafe.Add(mBase, uint32(v1119))) = v1142
												*(*int64)(unsafe.Add(mBase, uint32(v1119)+8)) = v1143
												v1148 = v28 + int32(48)
												v1150 = v1092 + int32(112)
												if v1150&int32(64) == int32(0) {
													if v1150 == int32(0) {
														v1171 = v1094
														v1172 = v1093
													} else {
														v1167 = base.I64_extend_i32_u(v1150)
														v1171 = v1094 << (uint(v1167) % 64)
														v1172 = int64(base.Ui64(v1094)>>(uint(base.I64_extend_i32_u(int32(64)-v1150))%64)) | v1093<<(uint(v1167)%64)
													}
												} else {
													v1171 = int64(0)
													v1172 = v1094 << (uint(base.I64_extend_i32_u(v1092+int32(48))) % 64)
												}
												*(*int64)(unsafe.Add(mBase, uint32(v1148))) = v1171
												*(*int64)(unsafe.Add(mBase, uint32(v1148)+8)) = v1172
												v1177 = v28 + int32(32)
												v1178 = *(*int64)(unsafe.Add(mBase, uint32(v28)+64))
												v1183 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(72))))
												v1188 = int64(32)
												v1189 = int64(base.Ui64(v1178) >> (uint(v1188) % 64))
												v1191 = int64(base.Ui64(v216) >> (uint(v1188) % 64))
												v1194 = int64(4294967295)
												v1195 = v1178 & v1194
												v1197 = v216 & v1194
												v1198 = v1195 * v1197
												v1202 = int64(base.Ui64(v1198)>>(uint(v1188)%64)) + v1195*v1191
												v1209 = v1202&v1194 + v1189*v1197
												*(*int64)(unsafe.Add(mBase, uint32(v1177)+8)) = v1183*v216 + v227*v1178 + v1189*v1191 + int64(base.Ui64(v1202)>>(uint(v1188)%64)) + int64(base.Ui64(v1209)>>(uint(v1188)%64))
												*(*int64)(unsafe.Add(mBase, uint32(v1177))) = v1209<<(uint(v1188)%64) | v1198&v1194
												v1224 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(56))))
												v1229 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(40))))
												v1230 = int64(1)
												v1232 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
												v1237 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
												v1239 = v1232 << (uint(v1230) % 64)
												v1244 = v1224 - (v1229<<(uint(v1230)%64) | int64(base.Ui64(v1232)>>(uint(int64(63))%64))) - base.I64_extend_i32_u(base.B2i32(base.Ui64(v1237) < base.Ui64(v1239)))
												v1245 = v1178
												v1246 = v1237 - v1239
												v1247 = v1183
												v1250 = v28 + int32(16)
												v1251 = int64(3)
												v1252 = int64(0)
												v1257 = int64(32)
												v1260 = int64(base.Ui64(v216) >> (uint(v1257) % 64))
												v1263 = int64(4294967295)
												v1266 = v216 & v1263
												v1267 = v1251 * v1266
												v1271 = int64(base.Ui64(v1267)>>(uint(v1257)%64)) + v1251*v1260
												v1278 = v1271&v1263 + v1252*v1266
												*(*int64)(unsafe.Add(mBase, uint32(v1250)+8)) = v1252*v216 + v227*v1251 + v1252*v1260 + int64(base.Ui64(v1271)>>(uint(v1257)%64)) + int64(base.Ui64(v1278)>>(uint(v1257)%64))
												*(*int64)(unsafe.Add(mBase, uint32(v1250))) = v1278<<(uint(v1257)%64) | v1267&v1263
												v1289 = int64(5)
												v1290 = int64(0)
												v1295 = int64(32)
												v1298 = int64(base.Ui64(v216) >> (uint(v1295) % 64))
												v1301 = int64(4294967295)
												v1304 = v216 & v1301
												v1305 = v1289 * v1304
												v1309 = int64(base.Ui64(v1305)>>(uint(v1295)%64)) + v1289*v1298
												v1316 = v1309&v1301 + v1290*v1304
												*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v1290*v216 + v227*v1289 + v1290*v1298 + int64(base.Ui64(v1309)>>(uint(v1295)%64)) + int64(base.Ui64(v1316)>>(uint(v1295)%64))
												*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1316<<(uint(v1295)%64) | v1305&v1301
												v1328 = v1245 & int64(1)
												v1329 = v1328 + v1246
												v1333 = v1244 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1329) < base.Ui64(v1328)))
												if v1333 == v227 {
													v1336 = base.B2i32(base.Ui64(v216) < base.Ui64(v1329))
												} else {
													v1336 = base.B2i32(base.Ui64(v227) < base.Ui64(v1333))
												}
												v1338 = v1245 + base.I64_extend_i32_u(v1336)
												v1341 = v1247 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1338) < base.Ui64(v1245)))
												v1344 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
												v1350 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(24))))
												if v1333 == v1350 {
													v1353 = base.B2i32(base.Ui64(v1344) < base.Ui64(v1329))
												} else {
													v1353 = base.B2i32(base.Ui64(v1350) < base.Ui64(v1333))
												}
												v1356 = v1338 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1341) < base.Ui64(int64(9223090561878065152)))&v1353)
												v1359 = v1341 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1356) < base.Ui64(v1338)))
												v1362 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
												v1366 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(8))))
												if v1333 == v1366 {
													v1369 = base.B2i32(base.Ui64(v1362) < base.Ui64(v1329))
												} else {
													v1369 = base.B2i32(base.Ui64(v1366) < base.Ui64(v1333))
												}
												v1372 = v1356 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1359) < base.Ui64(int64(9223090561878065152)))&v1369)
												v1377 = v1372
												v1383 = v1359 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1372) < base.Ui64(v1356))) | v36
											} else {
												v1377 = int64(0)
												v1383 = v36
											}
										} else {
											v1102 = int64(1)
											v1244 = v1090<<(uint(v1102)%64) | int64(base.Ui64(v1091)>>(uint(int64(63))%64))
											v1245 = v1088
											v1246 = v1091 << (uint(v1102) % 64)
											v1247 = base.I64_extend_i32_u(v1092)<<(uint(int64(48))%64) | v1089&int64(281474976710655)
											v1250 = v28 + int32(16)
											v1251 = int64(3)
											v1252 = int64(0)
											v1257 = int64(32)
											v1260 = int64(base.Ui64(v216) >> (uint(v1257) % 64))
											v1263 = int64(4294967295)
											v1266 = v216 & v1263
											v1267 = v1251 * v1266
											v1271 = int64(base.Ui64(v1267)>>(uint(v1257)%64)) + v1251*v1260
											v1278 = v1271&v1263 + v1252*v1266
											*(*int64)(unsafe.Add(mBase, uint32(v1250)+8)) = v1252*v216 + v227*v1251 + v1252*v1260 + int64(base.Ui64(v1271)>>(uint(v1257)%64)) + int64(base.Ui64(v1278)>>(uint(v1257)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v1250))) = v1278<<(uint(v1257)%64) | v1267&v1263
											v1289 = int64(5)
											v1290 = int64(0)
											v1295 = int64(32)
											v1298 = int64(base.Ui64(v216) >> (uint(v1295) % 64))
											v1301 = int64(4294967295)
											v1304 = v216 & v1301
											v1305 = v1289 * v1304
											v1309 = int64(base.Ui64(v1305)>>(uint(v1295)%64)) + v1289*v1298
											v1316 = v1309&v1301 + v1290*v1304
											*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v1290*v216 + v227*v1289 + v1290*v1298 + int64(base.Ui64(v1309)>>(uint(v1295)%64)) + int64(base.Ui64(v1316)>>(uint(v1295)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1316<<(uint(v1295)%64) | v1305&v1301
											v1328 = v1245 & int64(1)
											v1329 = v1328 + v1246
											v1333 = v1244 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1329) < base.Ui64(v1328)))
											if v1333 == v227 {
												v1336 = base.B2i32(base.Ui64(v216) < base.Ui64(v1329))
											} else {
												v1336 = base.B2i32(base.Ui64(v227) < base.Ui64(v1333))
											}
											v1338 = v1245 + base.I64_extend_i32_u(v1336)
											v1341 = v1247 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1338) < base.Ui64(v1245)))
											v1344 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
											v1350 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(24))))
											if v1333 == v1350 {
												v1353 = base.B2i32(base.Ui64(v1344) < base.Ui64(v1329))
											} else {
												v1353 = base.B2i32(base.Ui64(v1350) < base.Ui64(v1333))
											}
											v1356 = v1338 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1341) < base.Ui64(int64(9223090561878065152)))&v1353)
											v1359 = v1341 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1356) < base.Ui64(v1338)))
											v1362 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
											v1366 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(8))))
											if v1333 == v1366 {
												v1369 = base.B2i32(base.Ui64(v1362) < base.Ui64(v1329))
											} else {
												v1369 = base.B2i32(base.Ui64(v1366) < base.Ui64(v1333))
											}
											v1372 = v1356 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1359) < base.Ui64(int64(9223090561878065152)))&v1369)
											v1377 = v1372
											v1383 = v1359 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1372) < base.Ui64(v1356))) | v36
										}
									} else {
										v1377 = int64(0)
										v1383 = v36 | int64(9223090561878065152)
									}
								} else {
									v1377 = int64(0)
									v1383 = v36 | int64(9223090561878065152)
								}
							} else {
								if l3|v71 == int64(0) {
									v109 = int64(9223231299366420480)
								} else {
									v109 = v36
								}
								v1377 = int64(0)
								v1383 = v109
							}
						} else {
							v1377 = int64(0)
							v1383 = v36
						}
					} else {
						if base.B2i32(l3|(v71^int64(9223090561878065152)) == int64(0)) == int32(0) {
							v1377 = int64(0)
							v1383 = v36 | int64(9223090561878065152)
						} else {
							v1377 = int64(0)
							v1383 = int64(9223231299366420480)
						}
					}
				} else {
					v1377 = l3
					v1383 = l4 | int64(140737488355328)
				}
			} else {
				v1377 = l1
				v1383 = l2 | int64(140737488355328)
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v1377
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v1383
	m.G0 = v28 + int32(336)
	return
}
func F_decfloat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v100 int64
	_ = v100
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
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v131 int32
	_ = v131
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v208 int64
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int64
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v230 int64
	_ = v230
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int64
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int64
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v304 int64
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int64
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v318 int64
	_ = v318
	var v325 int64
	_ = v325
	var v326 int32
	_ = v326
	var v331 int64
	_ = v331
	var v332 int64
	_ = v332
	var v335 int32
	_ = v335
	var v339 int64
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v356 int64
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int64
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v370 int64
	_ = v370
	var v373 int32
	_ = v373
	var v388 int64
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int64
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v432 int64
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v464 int64
	_ = v464
	var v466 int32
	_ = v466
	var v469 int64
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int64
	_ = v492
	var v494 int64
	_ = v494
	var v498 int64
	_ = v498
	var v518 int64
	_ = v518
	var v532 int32
	_ = v532
	var v541 int64
	_ = v541
	var v544 int64
	_ = v544
	var v545 int64
	_ = v545
	var v546 int64
	_ = v546
	var v547 int64
	_ = v547
	var v562 int64
	_ = v562
	var v563 int64
	_ = v563
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int64
	_ = v580
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v594 int64
	_ = v594
	var v609 int64
	_ = v609
	var v610 int64
	_ = v610
	var v611 int64
	_ = v611
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v627 int64
	_ = v627
	var v632 int32
	_ = v632
	var v639 int64
	_ = v639
	var v648 int64
	_ = v648
	var v650 int64
	_ = v650
	var v651 int64
	_ = v651
	var v659 int64
	_ = v659
	var v664 int64
	_ = v664
	var v665 int64
	_ = v665
	var v670 int64
	_ = v670
	var v676 int64
	_ = v676
	var v677 int64
	_ = v677
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v695 int64
	_ = v695
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v709 int64
	_ = v709
	var v724 int64
	_ = v724
	var v725 int64
	_ = v725
	var v726 int64
	_ = v726
	var v735 int64
	_ = v735
	var v740 int64
	_ = v740
	var v741 int64
	_ = v741
	var v742 int64
	_ = v742
	var v746 int64
	_ = v746
	var v751 int64
	_ = v751
	var v759 int64
	_ = v759
	var v760 int64
	_ = v760
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v778 int64
	_ = v778
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v792 int64
	_ = v792
	var v807 int64
	_ = v807
	var v808 int64
	_ = v808
	var v809 int64
	_ = v809
	var v818 int64
	_ = v818
	var v823 int64
	_ = v823
	var v824 int64
	_ = v824
	var v825 int64
	_ = v825
	var v829 int64
	_ = v829
	var v834 int64
	_ = v834
	var v842 int64
	_ = v842
	var v843 int64
	_ = v843
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v872 int32
	_ = v872
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v930 int32
	_ = v930
	var v938 int32
	_ = v938
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v956 int64
	_ = v956
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v970 int64
	_ = v970
	var v985 int64
	_ = v985
	var v986 int64
	_ = v986
	var v987 int64
	_ = v987
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1004 int64
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1016 int64
	_ = v1016
	var v1025 int64
	_ = v1025
	var v1027 int64
	_ = v1027
	var v1028 int64
	_ = v1028
	var v1036 int64
	_ = v1036
	var v1041 int64
	_ = v1041
	var v1042 int64
	_ = v1042
	var v1047 int64
	_ = v1047
	var v1053 int64
	_ = v1053
	var v1054 int64
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1067 int64
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1081 int64
	_ = v1081
	var v1096 int64
	_ = v1096
	var v1097 int64
	_ = v1097
	var v1098 int64
	_ = v1098
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1115 int64
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1127 int64
	_ = v1127
	var v1136 int64
	_ = v1136
	var v1138 int64
	_ = v1138
	var v1139 int64
	_ = v1139
	var v1147 int64
	_ = v1147
	var v1152 int64
	_ = v1152
	var v1153 int64
	_ = v1153
	var v1158 int64
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1177 int64
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1191 int64
	_ = v1191
	var v1206 int64
	_ = v1206
	var v1207 int64
	_ = v1207
	var v1208 int64
	_ = v1208
	var v1217 int64
	_ = v1217
	var v1222 int64
	_ = v1222
	var v1223 int64
	_ = v1223
	var v1228 int64
	_ = v1228
	var v1234 int64
	_ = v1234
	var v1235 int64
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1255 int64
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1269 int64
	_ = v1269
	var v1284 int64
	_ = v1284
	var v1285 int64
	_ = v1285
	var v1286 int64
	_ = v1286
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1302 int64
	_ = v1302
	var v1307 int32
	_ = v1307
	var v1314 int64
	_ = v1314
	var v1323 int64
	_ = v1323
	var v1325 int64
	_ = v1325
	var v1326 int64
	_ = v1326
	var v1334 int64
	_ = v1334
	var v1339 int64
	_ = v1339
	var v1340 int64
	_ = v1340
	var v1345 int64
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1362 int64
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1376 int64
	_ = v1376
	var v1391 int64
	_ = v1391
	var v1392 int64
	_ = v1392
	var v1393 int64
	_ = v1393
	var v1402 int64
	_ = v1402
	var v1407 int64
	_ = v1407
	var v1408 int64
	_ = v1408
	var v1413 int64
	_ = v1413
	var v1419 int64
	_ = v1419
	var v1420 int64
	_ = v1420
	var v1440 int32
	_ = v1440
	var v1451 int32
	_ = v1451
	var v1455 int32
	_ = v1455
	var v1460 int32
	_ = v1460
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1614 int32
	_ = v1614
	var v1641 int32
	_ = v1641
	var v1643 int32
	_ = v1643
	var v1655 int32
	_ = v1655
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1678 int32
	_ = v1678
	var v1689 int32
	_ = v1689
	var v1692 int32
	_ = v1692
	var v1693 int64
	_ = v1693
	var v1697 int64
	_ = v1697
	var v1701 int64
	_ = v1701
	var v1702 int64
	_ = v1702
	var v1707 int64
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1726 int32
	_ = v1726
	var v1732 int32
	_ = v1732
	var v1735 int32
	_ = v1735
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1799 int32
	_ = v1799
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1829 int32
	_ = v1829
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1857 int32
	_ = v1857
	var v1883 int32
	_ = v1883
	var v1887 int32
	_ = v1887
	var v1890 int32
	_ = v1890
	var v1895 int32
	_ = v1895
	var v1899 int32
	_ = v1899
	var v1907 int64
	_ = v1907
	var v1911 int32
	_ = v1911
	var v1921 int64
	_ = v1921
	var v1925 int32
	_ = v1925
	var v1926 int64
	_ = v1926
	var v1937 int32
	_ = v1937
	var v1942 int32
	_ = v1942
	var v1952 int32
	_ = v1952
	var v1954 int32
	_ = v1954
	var v1960 int32
	_ = v1960
	var v1964 int32
	_ = v1964
	var v1966 int32
	_ = v1966
	var v1968 int64
	_ = v1968
	var v1973 int32
	_ = v1973
	var v1980 int64
	_ = v1980
	var v1989 int64
	_ = v1989
	var v1991 int64
	_ = v1991
	var v1992 int64
	_ = v1992
	var v2005 int64
	_ = v2005
	var v2010 int64
	_ = v2010
	var v2011 int64
	_ = v2011
	var v2016 int64
	_ = v2016
	var v2022 int64
	_ = v2022
	var v2023 int64
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2029 int32
	_ = v2029
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2038 int64
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2046 int32
	_ = v2046
	var v2052 int64
	_ = v2052
	var v2067 int64
	_ = v2067
	var v2068 int64
	_ = v2068
	var v2069 int64
	_ = v2069
	var v2078 int64
	_ = v2078
	var v2083 int64
	_ = v2083
	var v2089 int64
	_ = v2089
	var v2091 int64
	_ = v2091
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2103 int64
	_ = v2103
	var v2109 int32
	_ = v2109
	var v2113 int32
	_ = v2113
	var v2119 int32
	_ = v2119
	var v2130 int32
	_ = v2130
	var v2132 int32
	_ = v2132
	var v2136 int32
	_ = v2136
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2169 int32
	_ = v2169
	var v2180 int32
	_ = v2180
	var v2185 int32
	_ = v2185
	var v2186 float64
	_ = v2186
	var v2188 int32
	_ = v2188
	var v2192 float64
	_ = v2192
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2208 float64
	_ = v2208
	var v2215 int32
	_ = v2215
	var v2218 int32
	_ = v2218
	var v2221 float64
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2240 int64
	_ = v2240
	var v2242 int64
	_ = v2242
	var v2246 int64
	_ = v2246
	var v2266 int64
	_ = v2266
	var v2280 int32
	_ = v2280
	var v2289 int64
	_ = v2289
	var v2292 int64
	_ = v2292
	var v2293 int64
	_ = v2293
	var v2294 int64
	_ = v2294
	var v2295 int64
	_ = v2295
	var v2309 int32
	_ = v2309
	var v2310 int64
	_ = v2310
	var v2315 int64
	_ = v2315
	var v2317 int64
	_ = v2317
	var v2339 int64
	_ = v2339
	var v2340 int64
	_ = v2340
	var v2342 int32
	_ = v2342
	var v2343 float64
	_ = v2343
	var v2345 int32
	_ = v2345
	var v2349 float64
	_ = v2349
	var v2356 int32
	_ = v2356
	var v2359 int32
	_ = v2359
	var v2365 float64
	_ = v2365
	var v2372 int32
	_ = v2372
	var v2375 int32
	_ = v2375
	var v2378 float64
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2393 int32
	_ = v2393
	var v2395 int32
	_ = v2395
	var v2397 int64
	_ = v2397
	var v2399 int64
	_ = v2399
	var v2403 int64
	_ = v2403
	var v2423 int64
	_ = v2423
	var v2437 int32
	_ = v2437
	var v2446 int64
	_ = v2446
	var v2449 int64
	_ = v2449
	var v2450 int64
	_ = v2450
	var v2451 int64
	_ = v2451
	var v2452 int64
	_ = v2452
	var v2467 int64
	_ = v2467
	var v2472 int64
	_ = v2472
	var v2475 int32
	_ = v2475
	var v2476 int64
	_ = v2476
	var v2481 int64
	_ = v2481
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2490 int64
	_ = v2490
	var v2493 int64
	_ = v2493
	var v2501 int64
	_ = v2501
	var v2506 int64
	_ = v2506
	var v2512 int64
	_ = v2512
	var v2513 int64
	_ = v2513
	var v2514 int64
	_ = v2514
	var v2515 int64
	_ = v2515
	var v2516 int64
	_ = v2516
	var v2517 int64
	_ = v2517
	var v2518 int64
	_ = v2518
	var v2519 int64
	_ = v2519
	var v2523 int32
	_ = v2523
	var v2530 int32
	_ = v2530
	var v2539 int32
	_ = v2539
	var v2549 int32
	_ = v2549
	var v2551 int32
	_ = v2551
	var v2553 int64
	_ = v2553
	var v2555 int64
	_ = v2555
	var v2559 int64
	_ = v2559
	var v2579 int64
	_ = v2579
	var v2593 int32
	_ = v2593
	var v2602 int64
	_ = v2602
	var v2605 int64
	_ = v2605
	var v2606 int64
	_ = v2606
	var v2607 int64
	_ = v2607
	var v2608 int64
	_ = v2608
	var v2623 int64
	_ = v2623
	var v2628 int64
	_ = v2628
	var v2634 int64
	_ = v2634
	var v2635 int64
	_ = v2635
	var v2639 int32
	_ = v2639
	var v2649 int32
	_ = v2649
	var v2651 int32
	_ = v2651
	var v2653 int64
	_ = v2653
	var v2655 int64
	_ = v2655
	var v2659 int64
	_ = v2659
	var v2679 int64
	_ = v2679
	var v2693 int32
	_ = v2693
	var v2702 int64
	_ = v2702
	var v2705 int64
	_ = v2705
	var v2706 int64
	_ = v2706
	var v2707 int64
	_ = v2707
	var v2708 int64
	_ = v2708
	var v2723 int64
	_ = v2723
	var v2728 int64
	_ = v2728
	var v2734 int64
	_ = v2734
	var v2735 int64
	_ = v2735
	var v2736 float64
	_ = v2736
	var v2743 int32
	_ = v2743
	var v2752 int32
	_ = v2752
	var v2754 int32
	_ = v2754
	var v2756 int64
	_ = v2756
	var v2758 int64
	_ = v2758
	var v2762 int64
	_ = v2762
	var v2782 int64
	_ = v2782
	var v2796 int32
	_ = v2796
	var v2805 int64
	_ = v2805
	var v2808 int64
	_ = v2808
	var v2809 int64
	_ = v2809
	var v2810 int64
	_ = v2810
	var v2811 int64
	_ = v2811
	var v2826 int64
	_ = v2826
	var v2831 int64
	_ = v2831
	var v2837 int64
	_ = v2837
	var v2838 int64
	_ = v2838
	var v2840 int32
	_ = v2840
	var v2849 int32
	_ = v2849
	var v2851 int32
	_ = v2851
	var v2853 int64
	_ = v2853
	var v2855 int64
	_ = v2855
	var v2859 int64
	_ = v2859
	var v2879 int64
	_ = v2879
	var v2893 int32
	_ = v2893
	var v2902 int64
	_ = v2902
	var v2905 int64
	_ = v2905
	var v2906 int64
	_ = v2906
	var v2907 int64
	_ = v2907
	var v2908 int64
	_ = v2908
	var v2923 int64
	_ = v2923
	var v2928 int64
	_ = v2928
	var v2934 int64
	_ = v2934
	var v2935 int64
	_ = v2935
	var v2936 int64
	_ = v2936
	var v2937 int64
	_ = v2937
	var v2943 int64
	_ = v2943
	var v2946 int64
	_ = v2946
	var v2951 int64
	_ = v2951
	var v2961 int64
	_ = v2961
	var v2962 int64
	_ = v2962
	var v2966 int32
	_ = v2966
	var v2990 int32
	_ = v2990
	var v3002 int32
	_ = v3002
	var v3011 int32
	_ = v3011
	var v3015 int32
	_ = v3015
	var v3025 int64
	_ = v3025
	var v3026 int64
	_ = v3026
	var v3027 int64
	_ = v3027
	var v3029 int64
	_ = v3029
	var v3035 int32
	_ = v3035
	var v3036 int64
	_ = v3036
	var v3041 int64
	_ = v3041
	var v3043 int32
	_ = v3043
	var v3044 int32
	_ = v3044
	var v3045 int32
	_ = v3045
	var v3050 int64
	_ = v3050
	var v3053 int64
	_ = v3053
	var v3063 int64
	_ = v3063
	var v3064 int64
	_ = v3064
	var v3071 int32
	_ = v3071
	var v3078 int64
	_ = v3078
	var v3081 int64
	_ = v3081
	var v3086 int64
	_ = v3086
	var v3088 int64
	_ = v3088
	var v3092 int32
	_ = v3092
	var v3096 int64
	_ = v3096
	var v3097 int64
	_ = v3097
	var v3101 int32
	_ = v3101
	var v3125 int32
	_ = v3125
	var v3134 int32
	_ = v3134
	var v3140 int32
	_ = v3140
	var v3144 int32
	_ = v3144
	var v3149 int64
	_ = v3149
	var v3151 int32
	_ = v3151
	var v3152 int64
	_ = v3152
	var v3153 int64
	_ = v3153
	var v3154 int64
	_ = v3154
	var v3155 int64
	_ = v3155
	var v3164 int64
	_ = v3164
	var v3165 int64
	_ = v3165
	var v3169 int32
	_ = v3169
	var v3193 int32
	_ = v3193
	var v3205 int32
	_ = v3205
	var v3214 int32
	_ = v3214
	var v3218 int32
	_ = v3218
	var v3219 int32
	_ = v3219
	var v3224 int32
	_ = v3224
	var v3237 int32
	_ = v3237
	var v3240 int64
	_ = v3240
	var v3241 int64
	_ = v3241
	var v3243 int32
	_ = v3243
	var v3245 int32
	_ = v3245
	var v3247 int32
	_ = v3247
	var v3260 int64
	_ = v3260
	var v3261 int64
	_ = v3261
	var v3271 int32
	_ = v3271
	var v3274 int32
	_ = v3274
	var v3281 int64
	_ = v3281
	var v3282 int64
	_ = v3282
	var v3294 int64
	_ = v3294
	var v3295 int64
	_ = v3295
	var v3305 int32
	_ = v3305
	var v3308 int32
	_ = v3308
	var v3315 int64
	_ = v3315
	var v3316 int64
	_ = v3316
	var v3317 int64
	_ = v3317
	var v3318 int64
	_ = v3318
	var v3319 int32
	_ = v3319
	var v3329 int64
	_ = v3329
	var v3331 int64
	_ = v3331
	var v3340 int64
	_ = v3340
	var v3341 int64
	_ = v3341
	var v3353 int64
	_ = v3353
	var v3358 int64
	_ = v3358
	v8 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(8976)
	m.G0 = v28
	v32 = v8 - l4
	v33 = v32 - l3
	v38 = l2
	v48 = v8
	goto L4
L1:
	;
	v180 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+784)) = v180
	v184 = v157 + int32(-48)
	v186 = base.B2i32(v157 == int32(46))
	if v157 == int32(46) {
		goto L30
	} else {
		goto L31
	}
L2:
	;
	v86 = int64(0)
	if v85 != int32(48) {
		v131 = v85
		v140 = v86
		v141 = v48
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v83 = F___shgetc(m, l1)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L12
	} else {
		goto L14
	}
L4:
	;
	if v38 == int32(48) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v72 == v73 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	if v38 != int32(46) {
		v157 = v38
		v163 = v8
		v166 = int64(0)
		v167 = v48
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v65 == v66 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v65 + int32(1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v85 = v71
	goto L2
L10:
	;
	v81 = F___shgetc(m, l1)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v75 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v72 + v75
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	v38 = v79
	v48 = v75
	goto L4
L12:
	;
	return
L13:
	;
	v38 = v81
	v48 = int32(1)
	goto L4
L14:
	;
	v85 = v83
	goto L2
L15:
	;
	v157 = v131
	v163 = int32(1)
	v166 = v140
	v167 = v141
	goto L1
L16:
	;
	v100 = v86
	goto L17
L17:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v114 == v115 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v131 = v123
	v140 = v125
	v141 = int32(1)
	goto L15
L19:
	;
	v125 = v100 + int64(-1)
	if v123 == int32(48) {
		v100 = v125
		goto L17
	} else {
		goto L23
	}
L20:
	;
	v121 = F___shgetc(m, l1)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L12
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v114 + int32(1)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	v123 = v120
	goto L19
L22:
	;
	v123 = v121
	goto L19
L23:
	;
	goto L18
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v3353
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v3358
	m.G0 = v28 + int32(8976)
	return
L25:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v28)+784))
	if v478 != 0 {
		goto L76
	} else {
		goto L77
	}
L26:
	;
	v432 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v432
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = base.I64_extend_i32_s(v437 - v438)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	goto L73
L27:
	;
	if v392 == int32(0) {
		v464 = v388
		v466 = v390
		v469 = v393
		v470 = v394
		v471 = v395
		goto L25
	} else {
		goto L70
	}
L28:
	;
	v370 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v370 < int64(0) {
		v388 = v356
		v390 = v358
		v392 = v360
		v393 = v361
		v394 = v362
		v395 = v363
		goto L27
	} else {
		goto L69
	}
L29:
	;
	if v301 != 0 {
		goto L57
	} else {
		goto L58
	}
L30:
	;
	v194 = int32(0)
	v199 = v157
	v205 = v163
	v208 = v166
	v209 = v167
	v210 = v194
	v211 = v184
	v212 = v186
	v213 = int64(0)
	v214 = v194
	v215 = v194
	goto L33
L31:
	;
	if base.Ui32(v184) <= base.Ui32(int32(9)) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v190 = int32(0)
	v295 = v157
	v301 = v163
	v304 = v166
	v305 = v167
	v306 = v180
	v309 = int64(0)
	v310 = v190
	v311 = v190
	goto L29
L33:
	;
	if v212&int32(1) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v295 = v286
	v301 = v268
	v304 = v269
	v305 = v270
	v306 = v271
	v309 = v274
	v310 = v275
	v311 = v276
	goto L29
L35:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v277 == v278 {
		goto L52
	} else {
		goto L53
	}
L36:
	;
	v230 = v213 + int64(1)
	if int32(2044) < v214 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	if v205 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v356 = v208
	v358 = v210
	v360 = base.B2i32(v209 == int32(0))
	v361 = v213
	v362 = v214
	v363 = v215
	goto L28
L39:
	;
	v268 = int32(1)
	v269 = v213
	v270 = v209
	v271 = v210
	v274 = v213
	v275 = v214
	v276 = v215
	goto L35
L40:
	;
	if v199 == int32(48) {
		v268 = v205
		v269 = v208
		v270 = v209
		v271 = v210
		v274 = v230
		v275 = v214
		v276 = v215
		goto L35
	} else {
		goto L50
	}
L41:
	;
	v237 = v28 + int32(784) + v214<<(uint(int32(2))%32)
	if v215 == int32(0) {
		v246 = v211
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v199 == int32(48) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v246 = v199 + v240*int32(10) + int32(-48)
	goto L42
L44:
	;
	v250 = v210
	goto L46
L45:
	;
	v250 = base.I32_wrap_i64(v230)
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v246
	v252 = int32(1)
	v255 = v215 + v252
	v257 = base.B2i32(v255 == int32(9))
	if v255 == int32(9) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v258 = int32(0)
	goto L49
L48:
	;
	v258 = v255
	goto L49
L49:
	;
	v268 = v205
	v269 = v208
	v270 = v252
	v271 = v250
	v274 = v230
	v275 = v214 + v257
	v276 = v258
	goto L35
L50:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_decfloat[0])))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_decfloat[0]))) = v262 | int32(1)
	v268 = v205
	v269 = v208
	v270 = v209
	v271 = int32(18396)
	v274 = v230
	v275 = v214
	v276 = v215
	goto L35
L51:
	;
	v288 = v286 + int32(-48)
	v290 = base.B2i32(v286 == int32(46))
	if v286 == int32(46) {
		v199 = v286
		v205 = v268
		v208 = v269
		v209 = v270
		v210 = v271
		v211 = v288
		v212 = v290
		v213 = v274
		v214 = v275
		v215 = v276
		goto L33
	} else {
		goto L55
	}
L52:
	;
	v284 = F___shgetc(m, l1)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L12
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v277 + int32(1)
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	v286 = v283
	goto L51
L54:
	;
	v286 = v284
	goto L51
L55:
	;
	if base.Ui32(v288) < base.Ui32(int32(10)) {
		v199 = v286
		v205 = v268
		v208 = v269
		v209 = v270
		v210 = v271
		v211 = v288
		v212 = v290
		v213 = v274
		v214 = v275
		v215 = v276
		goto L33
	} else {
		goto L56
	}
L56:
	;
	goto L34
L57:
	;
	v318 = v304
	goto L59
L58:
	;
	v318 = v309
	goto L59
L59:
	;
	if v305 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v341 = int32(0)
	v342 = base.B2i32(v305 == v341)
	if v295 < v341 {
		v388 = v318
		v390 = v306
		v392 = v342
		v393 = v309
		v394 = v310
		v395 = v311
		goto L27
	} else {
		goto L68
	}
L61:
	;
	if v295&int32(-33) != int32(69) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v325 = F_scanexp(m, l1, l6)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L12
	} else {
		goto L64
	}
L63:
	;
	v464 = v339 + v318
	v466 = v306
	v469 = v309
	v470 = v310
	v471 = v311
	goto L25
L64:
	;
	if v325 != int64(-9223372036854775807-1) {
		v339 = v325
		goto L63
	} else {
		goto L65
	}
L65:
	;
	if l6 == int32(0) {
		goto L26
	} else {
		goto L66
	}
L66:
	;
	v331 = int64(0)
	v332 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v332 < v331 {
		v339 = v331
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v335 + int32(-1)
	v339 = v331
	goto L63
L68:
	;
	v356 = v318
	v358 = v306
	v360 = v342
	v361 = v309
	v362 = v310
	v363 = v311
	goto L28
L69:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v373 + int32(-1)
	v388 = v356
	v390 = v358
	v392 = v360
	v393 = v361
	v394 = v362
	v395 = v363
	goto L27
L70:
	;
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_decfloat[1])) = int32(28)
	goto L26
L72:
	;
	v3353 = int64(0)
	v3358 = v432
	goto L24
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v442
	goto L72
L76:
	;
	if int64(9) < v469 {
		goto L89
	} else {
		goto L90
	}
L77:
	;
	v488 = m.G0
	v490 = v488 - int32(16)
	m.G0 = v490
	v492 = base.I64_reinterpret_f64(base.F64_copysign(float64(0), base.F64_convert_i32_s(l5)))
	v494 = v492 & int64(4503599627370495)
	v498 = int64(base.Ui64(v492)>>(uint(int64(52))%64)) & int64(2047)
	if v498 == int64(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v562 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(8))))
	v563 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	v3353 = v562
	v3358 = v563
	goto L24
L79:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v545
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v546<<(uint(int64(48))%64) | v492&int64(-9223372036854775807-1) | v547
	m.G0 = v490 + int32(16)
	goto L78
L80:
	;
	if base.B2i32(v494 == int64(0)) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	if v498 == int64(2047) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v545 = v494 << (uint(int64(60)) % 64)
	v546 = int64(32767)
	v547 = int64(base.Ui64(v494) >> (uint(int64(4)) % 64))
	goto L79
L83:
	;
	v545 = v494 << (uint(int64(60)) % 64)
	v546 = v498 + int64(15360)
	v547 = int64(base.Ui64(v494) >> (uint(int64(4)) % 64))
	goto L79
L84:
	;
	if base.Ui64(v494) < base.Ui64(int64(4294967296)) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v518 = int64(0)
	v545 = v518
	v546 = v518
	v547 = v518
	goto L79
L86:
	;
	v532 = base.I32_clz(base.I32_wrap_i64(v492)) | int32(32)
	goto L88
L87:
	;
	v532 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v494) >> (uint(int64(32)) % 64))))
	goto L88
L88:
	;
	F___ashlti3(m, v490, v494, int64(0), v532+int32(49))
	mBase = m.M
	v541 = *(*int64)(unsafe.Add(mBase, uint32(v490+int32(8))))
	v544 = *(*int64)(unsafe.Add(mBase, uint32(v490)))
	v545 = v544
	v546 = base.I64_extend_i32_u(int32(15372) - v532)
	v547 = v541 ^ int64(281474976710656)
	goto L79
L89:
	;
	if v464 <= base.I64_extend_i32_u(int32(base.Ui32(v32)>>(uint(int32(1))%32))) {
		goto L103
	} else {
		goto L104
	}
L90:
	;
	if v464 != v469 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	if base.Ui32(int32(30)) < base.Ui32(l3) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v571 = v28 + int32(48)
	v576 = m.G0
	v578 = v576 - int32(16)
	m.G0 = v578
	if l5 != 0 {
		goto L97
	} else {
		goto L98
	}
L93:
	;
	if int32(base.Ui32(v478)>>(uint(l3)%32)) != 0 {
		goto L89
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v619 = v28 + int32(32)
	v623 = m.G0
	v625 = v623 - int32(16)
	m.G0 = v625
	if v478 != 0 {
		goto L101
	} else {
		goto L102
	}
L96:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v571))) = v610
	*(*int64)(unsafe.Add(mBase, uint32(v571)+8)) = v611
	m.G0 = v578 + int32(16)
	goto L95
L97:
	;
	v583 = l5 >> (uint(int32(31)) % 32)
	v585 = l5 ^ v583 - v583
	v588 = base.I32_clz(v585)
	F___ashlti3(m, v578, base.I64_extend_i32_u(v585), int64(0), v588+int32(81))
	mBase = m.M
	v594 = *(*int64)(unsafe.Add(mBase, uint32(v578+int32(8))))
	v609 = *(*int64)(unsafe.Add(mBase, uint32(v578)))
	v610 = v609
	v611 = v594 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v588)<<(uint(int64(48))%64) | base.I64_extend_i32_u(l5&int32(-2147483648))<<(uint(int64(32))%64)
	goto L96
L98:
	;
	v580 = int64(0)
	v610 = v580
	v611 = v580
	goto L96
L99:
	;
	v659 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
	v664 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(56))))
	v665 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
	v670 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(40))))
	F___multf3(m, v28+int32(16), v659, v664, v665, v670)
	mBase = m.M
	v676 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(24))))
	v677 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
	v3353 = v676
	v3358 = v677
	goto L24
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v619))) = v650
	*(*int64)(unsafe.Add(mBase, uint32(v619)+8)) = v651
	m.G0 = v625 + int32(16)
	goto L99
L101:
	;
	v632 = base.I32_clz(v478)
	F___ashlti3(m, v625, base.I64_extend_i32_u(v478), int64(0), int32(112)-(v632^int32(31)))
	mBase = m.M
	v639 = *(*int64)(unsafe.Add(mBase, uint32(v625+int32(8))))
	v648 = *(*int64)(unsafe.Add(mBase, uint32(v625)))
	v650 = v648
	v651 = v639 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v632)<<(uint(int64(48))%64)
	goto L100
L102:
	;
	v627 = int64(0)
	v650 = v627
	v651 = v627
	goto L100
L103:
	;
	if base.I64_extend_i32_s(l4+int32(-226)) <= v464 {
		goto L110
	} else {
		goto L111
	}
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_decfloat[1])) = int32(68)
	v686 = v28 + int32(96)
	v691 = m.G0
	v693 = v691 - int32(16)
	m.G0 = v693
	if l5 != 0 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v735 = *(*int64)(unsafe.Add(mBase, uint32(v28)+96))
	v740 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(104))))
	v741 = int64(-1)
	v742 = int64(9223090561878065151)
	F___multf3(m, v28+int32(80), v735, v740, v741, v742)
	mBase = m.M
	v746 = *(*int64)(unsafe.Add(mBase, uint32(v28)+80))
	v751 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(88))))
	F___multf3(m, v28+int32(64), v746, v751, v741, v742)
	mBase = m.M
	v759 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(72))))
	v760 = *(*int64)(unsafe.Add(mBase, uint32(v28)+64))
	v3353 = v759
	v3358 = v760
	goto L24
L107:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v686))) = v725
	*(*int64)(unsafe.Add(mBase, uint32(v686)+8)) = v726
	m.G0 = v693 + int32(16)
	goto L106
L108:
	;
	v698 = l5 >> (uint(int32(31)) % 32)
	v700 = l5 ^ v698 - v698
	v703 = base.I32_clz(v700)
	F___ashlti3(m, v693, base.I64_extend_i32_u(v700), int64(0), v703+int32(81))
	mBase = m.M
	v709 = *(*int64)(unsafe.Add(mBase, uint32(v693+int32(8))))
	v724 = *(*int64)(unsafe.Add(mBase, uint32(v693)))
	v725 = v724
	v726 = v709 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v703)<<(uint(int64(48))%64) | base.I64_extend_i32_u(l5&int32(-2147483648))<<(uint(int64(32))%64)
	goto L107
L109:
	;
	v695 = int64(0)
	v725 = v695
	v726 = v695
	goto L107
L110:
	;
	if v471 == int32(0) {
		v930 = v470
		goto L117
	} else {
		goto L118
	}
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_decfloat[1])) = int32(68)
	v769 = v28 + int32(144)
	v774 = m.G0
	v776 = v774 - int32(16)
	m.G0 = v776
	if l5 != 0 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v818 = *(*int64)(unsafe.Add(mBase, uint32(v28)+144))
	v823 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(152))))
	v824 = int64(0)
	v825 = int64(281474976710656)
	F___multf3(m, v28+int32(128), v818, v823, v824, v825)
	mBase = m.M
	v829 = *(*int64)(unsafe.Add(mBase, uint32(v28)+128))
	v834 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(136))))
	F___multf3(m, v28+int32(112), v829, v834, v824, v825)
	mBase = m.M
	v842 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(120))))
	v843 = *(*int64)(unsafe.Add(mBase, uint32(v28)+112))
	v3353 = v842
	v3358 = v843
	goto L24
L114:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v769))) = v808
	*(*int64)(unsafe.Add(mBase, uint32(v769)+8)) = v809
	m.G0 = v776 + int32(16)
	goto L113
L115:
	;
	v781 = l5 >> (uint(int32(31)) % 32)
	v783 = l5 ^ v781 - v781
	v786 = base.I32_clz(v783)
	F___ashlti3(m, v776, base.I64_extend_i32_u(v783), int64(0), v786+int32(81))
	mBase = m.M
	v792 = *(*int64)(unsafe.Add(mBase, uint32(v776+int32(8))))
	v807 = *(*int64)(unsafe.Add(mBase, uint32(v776)))
	v808 = v807
	v809 = v792 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v786)<<(uint(int64(48))%64) | base.I64_extend_i32_u(l5&int32(-2147483648))<<(uint(int64(32))%64)
	goto L114
L116:
	;
	v778 = int64(0)
	v808 = v778
	v809 = v778
	goto L114
L117:
	;
	v938 = base.I32_wrap_i64(v464)
	if int32(9) <= v466 {
		goto L124
	} else {
		goto L125
	}
L118:
	;
	if int32(8) < v471 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v930 = v470 + int32(1)
	goto L117
L120:
	;
	v852 = v28 + int32(784) + v470<<(uint(int32(2))%32)
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v852)))
	v855 = v853
	v872 = v471
	goto L121
L121:
	;
	v880 = v855 * int32(10)
	v882 = v872 + int32(1)
	if v882 != int32(9) {
		v855 = v880
		v872 = v882
		goto L121
	} else {
		goto L123
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v852))) = v880
	goto L119
L123:
	;
	goto L122
L124:
	;
	v1440 = v930
	goto L167
L125:
	;
	if int64(17) < v464 {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	if v938 < v466 {
		goto L124
	} else {
		goto L127
	}
L127:
	;
	if v464 != int64(9) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if int64(8) < v464 {
		goto L138
	} else {
		goto L139
	}
L129:
	;
	v947 = v28 + int32(192)
	v952 = m.G0
	v954 = v952 - int32(16)
	m.G0 = v954
	if l5 != 0 {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v995 = v28 + int32(176)
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v28)+784))
	v1000 = m.G0
	v1002 = v1000 - int32(16)
	m.G0 = v1002
	if v996 != 0 {
		goto L136
	} else {
		goto L137
	}
L131:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v947))) = v986
	*(*int64)(unsafe.Add(mBase, uint32(v947)+8)) = v987
	m.G0 = v954 + int32(16)
	goto L130
L132:
	;
	v959 = l5 >> (uint(int32(31)) % 32)
	v961 = l5 ^ v959 - v959
	v964 = base.I32_clz(v961)
	F___ashlti3(m, v954, base.I64_extend_i32_u(v961), int64(0), v964+int32(81))
	mBase = m.M
	v970 = *(*int64)(unsafe.Add(mBase, uint32(v954+int32(8))))
	v985 = *(*int64)(unsafe.Add(mBase, uint32(v954)))
	v986 = v985
	v987 = v970 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v964)<<(uint(int64(48))%64) | base.I64_extend_i32_u(l5&int32(-2147483648))<<(uint(int64(32))%64)
	goto L131
L133:
	;
	v956 = int64(0)
	v986 = v956
	v987 = v956
	goto L131
L134:
	;
	v1036 = *(*int64)(unsafe.Add(mBase, uint32(v28)+192))
	v1041 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(200))))
	v1042 = *(*int64)(unsafe.Add(mBase, uint32(v28)+176))
	v1047 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(184))))
	F___multf3(m, v28+int32(160), v1036, v1041, v1042, v1047)
	mBase = m.M
	v1053 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(168))))
	v1054 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
	v3353 = v1053
	v3358 = v1054
	goto L24
L135:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v995))) = v1027
	*(*int64)(unsafe.Add(mBase, uint32(v995)+8)) = v1028
	m.G0 = v1002 + int32(16)
	goto L134
L136:
	;
	v1009 = base.I32_clz(v996)
	F___ashlti3(m, v1002, base.I64_extend_i32_u(v996), int64(0), int32(112)-(v1009^int32(31)))
	mBase = m.M
	v1016 = *(*int64)(unsafe.Add(mBase, uint32(v1002+int32(8))))
	v1025 = *(*int64)(unsafe.Add(mBase, uint32(v1002)))
	v1027 = v1025
	v1028 = v1016 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1009)<<(uint(int64(48))%64)
	goto L135
L137:
	;
	v1004 = int64(0)
	v1027 = v1004
	v1028 = v1004
	goto L135
L138:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v28)+784))
	v1241 = l3 + v938*int32(-3) + int32(27)
	if int32(30) < v1241 {
		goto L152
	} else {
		goto L153
	}
L139:
	;
	v1058 = v28 + int32(272)
	v1063 = m.G0
	v1065 = v1063 - int32(16)
	m.G0 = v1065
	if l5 != 0 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v1106 = v28 + int32(256)
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v28)+784))
	v1111 = m.G0
	v1113 = v1111 - int32(16)
	m.G0 = v1113
	if v1107 != 0 {
		goto L146
	} else {
		goto L147
	}
L141:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1058))) = v1097
	*(*int64)(unsafe.Add(mBase, uint32(v1058)+8)) = v1098
	m.G0 = v1065 + int32(16)
	goto L140
L142:
	;
	v1070 = l5 >> (uint(int32(31)) % 32)
	v1072 = l5 ^ v1070 - v1070
	v1075 = base.I32_clz(v1072)
	F___ashlti3(m, v1065, base.I64_extend_i32_u(v1072), int64(0), v1075+int32(81))
	mBase = m.M
	v1081 = *(*int64)(unsafe.Add(mBase, uint32(v1065+int32(8))))
	v1096 = *(*int64)(unsafe.Add(mBase, uint32(v1065)))
	v1097 = v1096
	v1098 = v1081 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1075)<<(uint(int64(48))%64) | base.I64_extend_i32_u(l5&int32(-2147483648))<<(uint(int64(32))%64)
	goto L141
L143:
	;
	v1067 = int64(0)
	v1097 = v1067
	v1098 = v1067
	goto L141
L144:
	;
	v1147 = *(*int64)(unsafe.Add(mBase, uint32(v28)+272))
	v1152 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(280))))
	v1153 = *(*int64)(unsafe.Add(mBase, uint32(v28)+256))
	v1158 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(264))))
	F___multf3(m, v28+int32(240), v1147, v1152, v1153, v1158)
	mBase = m.M
	v1161 = v28 + int32(224)
	v1168 = *(*int32)(unsafe.Add(mBase, uint32((int32(8)-v938)<<(uint(int32(2))%32))+uint32(_c_F_decfloat[2])))
	v1173 = m.G0
	v1175 = v1173 - int32(16)
	m.G0 = v1175
	if v1168 != 0 {
		goto L150
	} else {
		goto L151
	}
L145:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1106))) = v1138
	*(*int64)(unsafe.Add(mBase, uint32(v1106)+8)) = v1139
	m.G0 = v1113 + int32(16)
	goto L144
L146:
	;
	v1120 = base.I32_clz(v1107)
	F___ashlti3(m, v1113, base.I64_extend_i32_u(v1107), int64(0), int32(112)-(v1120^int32(31)))
	mBase = m.M
	v1127 = *(*int64)(unsafe.Add(mBase, uint32(v1113+int32(8))))
	v1136 = *(*int64)(unsafe.Add(mBase, uint32(v1113)))
	v1138 = v1136
	v1139 = v1127 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1120)<<(uint(int64(48))%64)
	goto L145
L147:
	;
	v1115 = int64(0)
	v1138 = v1115
	v1139 = v1115
	goto L145
L148:
	;
	v1217 = *(*int64)(unsafe.Add(mBase, uint32(v28)+240))
	v1222 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(248))))
	v1223 = *(*int64)(unsafe.Add(mBase, uint32(v28)+224))
	v1228 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(232))))
	F___divtf3(m, v28+int32(208), v1217, v1222, v1223, v1228)
	mBase = m.M
	v1234 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(216))))
	v1235 = *(*int64)(unsafe.Add(mBase, uint32(v28)+208))
	v3353 = v1234
	v3358 = v1235
	goto L24
L149:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1161))) = v1207
	*(*int64)(unsafe.Add(mBase, uint32(v1161)+8)) = v1208
	m.G0 = v1175 + int32(16)
	goto L148
L150:
	;
	v1180 = v1168 >> (uint(int32(31)) % 32)
	v1182 = v1168 ^ v1180 - v1180
	v1185 = base.I32_clz(v1182)
	F___ashlti3(m, v1175, base.I64_extend_i32_u(v1182), int64(0), v1185+int32(81))
	mBase = m.M
	v1191 = *(*int64)(unsafe.Add(mBase, uint32(v1175+int32(8))))
	v1206 = *(*int64)(unsafe.Add(mBase, uint32(v1175)))
	v1207 = v1206
	v1208 = v1191 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1185)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v1168&int32(-2147483648))<<(uint(int64(32))%64)
	goto L149
L151:
	;
	v1177 = int64(0)
	v1207 = v1177
	v1208 = v1177
	goto L149
L152:
	;
	v1246 = v28 + int32(352)
	v1251 = m.G0
	v1253 = v1251 - int32(16)
	m.G0 = v1253
	if l5 != 0 {
		goto L157
	} else {
		goto L158
	}
L153:
	;
	if int32(base.Ui32(v1236)>>(uint(v1241)%32)) != 0 {
		goto L124
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v1294 = v28 + int32(336)
	v1298 = m.G0
	v1300 = v1298 - int32(16)
	m.G0 = v1300
	if v1236 != 0 {
		goto L161
	} else {
		goto L162
	}
L156:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1246))) = v1285
	*(*int64)(unsafe.Add(mBase, uint32(v1246)+8)) = v1286
	m.G0 = v1253 + int32(16)
	goto L155
L157:
	;
	v1258 = l5 >> (uint(int32(31)) % 32)
	v1260 = l5 ^ v1258 - v1258
	v1263 = base.I32_clz(v1260)
	F___ashlti3(m, v1253, base.I64_extend_i32_u(v1260), int64(0), v1263+int32(81))
	mBase = m.M
	v1269 = *(*int64)(unsafe.Add(mBase, uint32(v1253+int32(8))))
	v1284 = *(*int64)(unsafe.Add(mBase, uint32(v1253)))
	v1285 = v1284
	v1286 = v1269 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1263)<<(uint(int64(48))%64) | base.I64_extend_i32_u(l5&int32(-2147483648))<<(uint(int64(32))%64)
	goto L156
L158:
	;
	v1255 = int64(0)
	v1285 = v1255
	v1286 = v1255
	goto L156
L159:
	;
	v1334 = *(*int64)(unsafe.Add(mBase, uint32(v28)+352))
	v1339 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(360))))
	v1340 = *(*int64)(unsafe.Add(mBase, uint32(v28)+336))
	v1345 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(344))))
	F___multf3(m, v28+int32(320), v1334, v1339, v1340, v1345)
	mBase = m.M
	v1348 = v28 + int32(304)
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v938<<(uint(int32(2))%32))+uint32(_c_F_decfloat[3])))
	v1358 = m.G0
	v1360 = v1358 - int32(16)
	m.G0 = v1360
	if v1353 != 0 {
		goto L165
	} else {
		goto L166
	}
L160:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1294))) = v1325
	*(*int64)(unsafe.Add(mBase, uint32(v1294)+8)) = v1326
	m.G0 = v1300 + int32(16)
	goto L159
L161:
	;
	v1307 = base.I32_clz(v1236)
	F___ashlti3(m, v1300, base.I64_extend_i32_u(v1236), int64(0), int32(112)-(v1307^int32(31)))
	mBase = m.M
	v1314 = *(*int64)(unsafe.Add(mBase, uint32(v1300+int32(8))))
	v1323 = *(*int64)(unsafe.Add(mBase, uint32(v1300)))
	v1325 = v1323
	v1326 = v1314 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1307)<<(uint(int64(48))%64)
	goto L160
L162:
	;
	v1302 = int64(0)
	v1325 = v1302
	v1326 = v1302
	goto L160
L163:
	;
	v1402 = *(*int64)(unsafe.Add(mBase, uint32(v28)+320))
	v1407 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(328))))
	v1408 = *(*int64)(unsafe.Add(mBase, uint32(v28)+304))
	v1413 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(312))))
	F___multf3(m, v28+int32(288), v1402, v1407, v1408, v1413)
	mBase = m.M
	v1419 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(296))))
	v1420 = *(*int64)(unsafe.Add(mBase, uint32(v28)+288))
	v3353 = v1419
	v3358 = v1420
	goto L24
L164:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1348))) = v1392
	*(*int64)(unsafe.Add(mBase, uint32(v1348)+8)) = v1393
	m.G0 = v1360 + int32(16)
	goto L163
L165:
	;
	v1365 = v1353 >> (uint(int32(31)) % 32)
	v1367 = v1353 ^ v1365 - v1365
	v1370 = base.I32_clz(v1367)
	F___ashlti3(m, v1360, base.I64_extend_i32_u(v1367), int64(0), v1370+int32(81))
	mBase = m.M
	v1376 = *(*int64)(unsafe.Add(mBase, uint32(v1360+int32(8))))
	v1391 = *(*int64)(unsafe.Add(mBase, uint32(v1360)))
	v1392 = v1391
	v1393 = v1376 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1370)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v1353&int32(-2147483648))<<(uint(int64(32))%64)
	goto L164
L166:
	;
	v1362 = int64(0)
	v1392 = v1362
	v1393 = v1362
	goto L164
L167:
	;
	v1451 = v1440 + int32(-1)
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(784)+v1451<<(uint(int32(2))%32))))
	if v1455 == int32(0) {
		v1440 = v1451
		goto L167
	} else {
		goto L169
	}
L168:
	;
	v1460 = base.I32_rem_s(v938, int32(9))
	if v1460 != 0 {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	goto L168
L170:
	;
	v1609 = int32(0)
	v1610 = v1585
	v1611 = v1586
	v1614 = v1589
	goto L189
L171:
	;
	if v464 < int64(0) {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v1585 = int32(0)
	v1586 = v1440
	v1589 = v938
	goto L170
L173:
	;
	v1466 = v1460 + int32(9)
	goto L175
L174:
	;
	v1466 = v1460
	goto L175
L175:
	;
	if v1440 != 0 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v1585 = v1557
	v1586 = v1558
	v1589 = v1561 - v1466 + int32(9)
	goto L170
L177:
	;
	v1476 = *(*int32)(unsafe.Add(mBase, uint32((int32(8)-v1466)<<(uint(int32(2))%32))+uint32(_c_F_decfloat[2])))
	v1477 = base.I32_div_s(int32(1000000000), v1476)
	v1478 = int32(0)
	v1482 = v1478
	v1483 = v1478
	v1495 = v1478
	v1499 = v938
	goto L179
L178:
	;
	v1467 = int32(0)
	v1557 = v1467
	v1558 = v1467
	v1561 = v938
	goto L176
L179:
	;
	v1510 = v28 + int32(784) + v1482<<(uint(int32(2))%32)
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1510)))
	v1512 = base.I32_div_u_s(v1511, v1476)
	v1513 = v1512 + v1483
	*(*int32)(unsafe.Add(mBase, uint32(v1510))) = v1513
	v1522 = base.B2i32(v1482 == v1495) & base.B2i32(v1513 == int32(0))
	if v1522 != 0 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	if v1529 == int32(0) {
		v1557 = v1523
		v1558 = v1440
		v1561 = v1526
		goto L176
	} else {
		goto L188
	}
L181:
	;
	v1523 = (v1495 + int32(1)) & int32(2047)
	goto L183
L182:
	;
	v1523 = v1495
	goto L183
L183:
	;
	if v1522 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1526 = v1499 + int32(-9)
	goto L186
L185:
	;
	v1526 = v1499
	goto L186
L186:
	;
	v1529 = v1477 * (v1511 - v1512*v1476)
	v1531 = v1482 + int32(1)
	if v1531 != v1440 {
		v1482 = v1531
		v1483 = v1529
		v1495 = v1523
		v1499 = v1526
		goto L179
	} else {
		goto L187
	}
L187:
	;
	goto L180
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28+int32(784)+v1440<<(uint(int32(2))%32)))) = v1529
	v1557 = v1523
	v1558 = v1440 + int32(1)
	v1561 = v1526
	goto L176
L189:
	;
	v1641 = v1609
	v1643 = v1611
	goto L192
L190:
	;
	v1775 = v1641
	v1776 = v1610
	v1777 = v1643
	v1780 = v1614
	goto L218
L191:
	;
	goto L190
L192:
	;
	if v1614 < int32(36) {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v1732 = (v1610 + int32(-1)) & int32(2047)
	if v1732 == v1676 {
		goto L215
	} else {
		goto L216
	}
L194:
	;
	v1673 = int32(0)
	v1676 = v1643
	v1678 = v1643 + int32(2047)
	goto L198
L195:
	;
	if v1614 != int32(36) {
		goto L191
	} else {
		goto L196
	}
L196:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(784)+v1610<<(uint(int32(2))%32))))
	if base.Ui32(int32(10384593)) <= base.Ui32(v1655) {
		goto L191
	} else {
		goto L197
	}
L197:
	;
	goto L194
L198:
	;
	v1689 = v1678 & int32(2047)
	v1692 = v28 + int32(784) + v1689<<(uint(int32(2))%32)
	v1693 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1692))))
	v1697 = v1693<<(uint(int64(29))%64) + base.I64_extend_i32_u(v1673)
	if base.Ui64(int64(1000000001)) <= base.Ui64(v1697) {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v1726 = v1641 + int32(-29)
	if v1708 == int32(0) {
		v1641 = v1726
		v1643 = v1676
		goto L192
	} else {
		goto L213
	}
L200:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v1692))) = uint32(v1707)
	if v1707 == int64(0) {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	v1701 = int64(1000000000)
	v1702 = base.I64_div_u_s(v1697, v1701)
	v1707 = v1697 - v1702*v1701
	v1708 = base.I32_wrap_i64(v1702)
	goto L200
L202:
	;
	v1707 = v1697
	v1708 = int32(0)
	goto L200
L203:
	;
	v1713 = v1689
	goto L205
L204:
	;
	v1713 = v1676
	goto L205
L205:
	;
	if v1689 == v1610 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v1715 = v1676
	goto L208
L207:
	;
	v1715 = v1713
	goto L208
L208:
	;
	v1719 = (v1676 + int32(-1)) & int32(2047)
	if v1689 != v1719 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v1721 = v1676
	goto L211
L210:
	;
	v1721 = v1715
	goto L211
L211:
	;
	if v1689 != v1610 {
		v1673 = v1708
		v1676 = v1721
		v1678 = v1689 + int32(-1)
		goto L198
	} else {
		goto L212
	}
L212:
	;
	goto L199
L213:
	;
	goto L193
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28+int32(784)+v1732<<(uint(int32(2))%32)))) = v1708
	v1609 = v1726
	v1610 = v1732
	v1611 = v1753
	v1614 = v1614 + int32(9)
	goto L189
L215:
	;
	v1735 = v28 + int32(784)
	v1740 = int32(2)
	v1742 = v1735 + (v1676+int32(2046))&int32(2047)<<(uint(v1740)%32)
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v1742)))
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1735+v1719<<(uint(v1740)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1742))) = v1743 | v1749
	v1753 = v1719
	goto L214
L216:
	;
	v1753 = v1676
	goto L214
L217:
	;
	v2523 = (v1844 + int32(4)) & int32(2047)
	if v2523 == v1952 {
		v3027 = v2514
		v3029 = v2517
		goto L324
	} else {
		goto L325
	}
L218:
	;
	v1789 = int32(2047)
	v1790 = (v1777 + int32(1)) & v1789
	v1799 = v28 + int32(784) + (v1777+int32(-1))&v1789<<(uint(int32(2))%32)
	v1813 = v1775
	v1814 = v1776
	v1818 = v1780
	goto L220
L219:
	;
	v2185 = v28 + int32(656)
	v2186 = float64(1)
	v2188 = int32(225) - v2100
	if v2188 < int32(1024) {
		goto L272
	} else {
		goto L273
	}
L220:
	;
	if int32(45) < v1818 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	goto L219
L222:
	;
	v1829 = int32(9)
	goto L224
L223:
	;
	v1829 = int32(1)
	goto L224
L224:
	;
	v1843 = v1813
	v1844 = v1814
	goto L226
L225:
	;
	goto L221
L226:
	;
	v1857 = int32(0)
	goto L230
L227:
	;
	v2113 = int32(-1)
	v2119 = int32(0)
	v2130 = v1844
	v2132 = v1844
	v2136 = v1818
	goto L258
L228:
	;
	v2109 = v1829 + v1843
	if v1844 == v1777 {
		v1843 = v2109
		v1844 = v1777
		goto L226
	} else {
		goto L257
	}
L229:
	;
	if v1818 != int32(36) {
		goto L228
	} else {
		goto L236
	}
L230:
	;
	v1883 = (v1857 + v1844) & int32(2047)
	if v1883 == v1777 {
		goto L229
	} else {
		goto L232
	}
L231:
	;
	goto L229
L232:
	;
	v1887 = int32(2)
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(784)+v1883<<(uint(v1887)%32))))
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1857<<(uint(v1887)%32))+uint32(_c_F_decfloat[4])))
	if base.Ui32(v1890) < base.Ui32(v1895) {
		goto L229
	} else {
		goto L233
	}
L233:
	;
	if base.Ui32(v1895) < base.Ui32(v1890) {
		goto L228
	} else {
		goto L234
	}
L234:
	;
	v1899 = v1857 + int32(1)
	if v1899 != int32(4) {
		v1857 = v1899
		goto L230
	} else {
		goto L235
	}
L235:
	;
	goto L231
L236:
	;
	v1907 = int64(0)
	v1911 = int32(0)
	v1921 = v1907
	v1925 = v1777
	v1926 = v1907
	goto L237
L237:
	;
	v1937 = (v1911 + v1844) & int32(2047)
	if v1937 != v1925 {
		v1952 = v1925
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v2029 = v28 + int32(720)
	v2034 = m.G0
	v2036 = v2034 - int32(16)
	m.G0 = v2036
	if l5 != 0 {
		goto L248
	} else {
		goto L249
	}
L239:
	;
	v1954 = v28 + int32(768)
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(784)+v1937<<(uint(int32(2))%32))))
	v1964 = m.G0
	v1966 = v1964 - int32(16)
	m.G0 = v1966
	if v1960 != 0 {
		goto L243
	} else {
		goto L244
	}
L240:
	;
	v1942 = (v1925 + int32(1)) & int32(2047)
	*(*int32)(unsafe.Add(mBase, uint32(v1942<<(uint(int32(2))%32)+(v28+int32(784))+int32(-4)))) = int32(0)
	v1952 = v1942
	goto L239
L241:
	;
	F___multf3(m, v28+int32(752), v1921, v1926, int64(0), int64(4619810130798575616))
	mBase = m.M
	v2005 = *(*int64)(unsafe.Add(mBase, uint32(v28)+752))
	v2010 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(760))))
	v2011 = *(*int64)(unsafe.Add(mBase, uint32(v28)+768))
	v2016 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(776))))
	F___addtf3(m, v28+int32(736), v2005, v2010, v2011, v2016)
	mBase = m.M
	v2022 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(744))))
	v2023 = *(*int64)(unsafe.Add(mBase, uint32(v28)+736))
	v2025 = v1911 + int32(1)
	if v2025 != int32(4) {
		v1911 = v2025
		v1921 = v2023
		v1925 = v1952
		v1926 = v2022
		goto L237
	} else {
		goto L245
	}
L242:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1954))) = v1991
	*(*int64)(unsafe.Add(mBase, uint32(v1954)+8)) = v1992
	m.G0 = v1966 + int32(16)
	goto L241
L243:
	;
	v1973 = base.I32_clz(v1960)
	F___ashlti3(m, v1966, base.I64_extend_i32_u(v1960), int64(0), int32(112)-(v1973^int32(31)))
	mBase = m.M
	v1980 = *(*int64)(unsafe.Add(mBase, uint32(v1966+int32(8))))
	v1989 = *(*int64)(unsafe.Add(mBase, uint32(v1966)))
	v1991 = v1989
	v1992 = v1980 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1973)<<(uint(int64(48))%64)
	goto L242
L244:
	;
	v1968 = int64(0)
	v1991 = v1968
	v1992 = v1968
	goto L242
L245:
	;
	goto L238
L246:
	;
	v2078 = *(*int64)(unsafe.Add(mBase, uint32(v28)+720))
	v2083 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(728))))
	F___multf3(m, v28+int32(704), v2023, v2022, v2078, v2083)
	mBase = m.M
	v2089 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(712))))
	v2091 = *(*int64)(unsafe.Add(mBase, uint32(v28)+704))
	v2093 = v1843 + int32(113)
	v2094 = v2093 - l4
	v2095 = int32(0)
	if v2095 < v2094 {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2029))) = v2068
	*(*int64)(unsafe.Add(mBase, uint32(v2029)+8)) = v2069
	m.G0 = v2036 + int32(16)
	goto L246
L248:
	;
	v2041 = l5 >> (uint(int32(31)) % 32)
	v2043 = l5 ^ v2041 - v2041
	v2046 = base.I32_clz(v2043)
	F___ashlti3(m, v2036, base.I64_extend_i32_u(v2043), int64(0), v2046+int32(81))
	mBase = m.M
	v2052 = *(*int64)(unsafe.Add(mBase, uint32(v2036+int32(8))))
	v2067 = *(*int64)(unsafe.Add(mBase, uint32(v2036)))
	v2068 = v2067
	v2069 = v2052 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v2046)<<(uint(int64(48))%64) | base.I64_extend_i32_u(l5&int32(-2147483648))<<(uint(int64(32))%64)
	goto L247
L249:
	;
	v2038 = int64(0)
	v2068 = v2038
	v2069 = v2038
	goto L247
L250:
	;
	v2098 = v2094
	goto L252
L251:
	;
	v2098 = v2095
	goto L252
L252:
	;
	v2099 = base.B2i32(v2094 < l3)
	if v2094 < l3 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v2100 = v2098
	goto L255
L254:
	;
	v2100 = l3
	goto L255
L255:
	;
	if base.Ui32(v2100) <= base.Ui32(int32(112)) {
		goto L225
	} else {
		goto L256
	}
L256:
	;
	v2103 = int64(0)
	v2514 = int64(0)
	v2515 = v2089
	v2516 = v2091
	v2517 = v2103
	v2518 = v2103
	v2519 = v2103
	goto L217
L257:
	;
	goto L227
L258:
	;
	v2147 = v28 + int32(784) + v2130<<(uint(int32(2))%32)
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(v2147)))
	v2150 = int32(base.Ui32(v2148)>>(uint(v1829)%32)) + v2119
	*(*int32)(unsafe.Add(mBase, uint32(v2147))) = v2150
	v2159 = base.B2i32(v2130 == v2132) & base.B2i32(v2150 == int32(0))
	if v2159 != 0 {
		goto L260
	} else {
		goto L261
	}
L259:
	;
	if v2165 == int32(0) {
		v1813 = v2109
		v1814 = v2160
		v1818 = v2163
		goto L220
	} else {
		goto L267
	}
L260:
	;
	v2160 = (v2132 + int32(1)) & int32(2047)
	goto L262
L261:
	;
	v2160 = v2132
	goto L262
L262:
	;
	if v2159 != 0 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v2163 = v2136 + int32(-9)
	goto L265
L264:
	;
	v2163 = v2136
	goto L265
L265:
	;
	v2165 = v2148 & (v2113<<(uint(v1829)%32) ^ v2113) * int32(base.Ui32(int32(1000000000))>>(uint(v1829)%32))
	v2169 = (v2130 + int32(1)) & int32(2047)
	if v2169 != v1777 {
		v2119 = v2165
		v2130 = v2169
		v2132 = v2160
		v2136 = v2163
		goto L258
	} else {
		goto L266
	}
L266:
	;
	goto L259
L267:
	;
	if v1790 == v2160 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(v1799)))
	*(*int32)(unsafe.Add(mBase, uint32(v1799))) = v2180 | int32(1)
	v1813 = v2109
	v1814 = v2160
	v1818 = v2163
	goto L220
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28+int32(784)+v1777<<(uint(int32(2))%32)))) = v2165
	v1775 = v2109
	v1776 = v2160
	v1777 = v1790
	v1780 = v2163
	goto L218
L270:
	;
	v2236 = m.G0
	v2238 = v2236 - int32(16)
	m.G0 = v2238
	v2240 = base.I64_reinterpret_f64(base.F64_mul(v2221, base.F64_reinterpret_i64(base.I64_extend_i32_u(v2222+int32(1023))<<(uint(int64(52))%64))))
	v2242 = v2240 & int64(4503599627370495)
	v2246 = int64(base.Ui64(v2240)>>(uint(int64(52))%64)) & int64(2047)
	if v2246 == int64(0) {
		goto L287
	} else {
		goto L288
	}
L271:
	;
	goto L270
L272:
	;
	if int32(-1023) < v2188 {
		v2221 = v2186
		v2222 = v2188
		goto L271
	} else {
		goto L279
	}
L273:
	;
	v2192 = base.F64_mul(v2186, float64(8.98846567431158e+307))
	if base.Ui32(int32(2047)) <= base.Ui32(v2188) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v2199 = int32(3069)
	if base.Ui32(v2188) < base.Ui32(v2199) {
		goto L276
	} else {
		goto L277
	}
L275:
	;
	v2221 = v2192
	v2222 = v2188 + int32(-1023)
	goto L271
L276:
	;
	v2202 = v2188
	goto L278
L277:
	;
	v2202 = v2199
	goto L278
L278:
	;
	v2221 = base.F64_mul(v2192, float64(8.98846567431158e+307))
	v2222 = v2202 + int32(-2046)
	goto L271
L279:
	;
	v2208 = base.F64_mul(v2186, float64(2.004168360008973e-292))
	if base.Ui32(v2188) <= base.Ui32(int32(-1992)) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v2215 = int32(-2960)
	if base.Ui32(v2215) < base.Ui32(v2188) {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	v2221 = v2208
	v2222 = v2188 + int32(969)
	goto L271
L282:
	;
	v2218 = v2188
	goto L284
L283:
	;
	v2218 = v2215
	goto L284
L284:
	;
	v2221 = base.F64_mul(v2208, float64(2.004168360008973e-292))
	v2222 = v2218 + int32(1938)
	goto L271
L285:
	;
	v2309 = v28 + int32(688)
	v2310 = *(*int64)(unsafe.Add(mBase, uint32(v28)+656))
	v2315 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(664))))
	*(*int64)(unsafe.Add(mBase, uint32(v2309))) = v2310
	v2317 = int64(48)
	*(*int64)(unsafe.Add(mBase, uint32(v2309)+8)) = base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(v2089)>>(uint(v2317)%64)))&int32(32768)|base.I32_wrap_i64(int64(base.Ui64(v2315&int64(9223090561878065152))>>(uint(v2317)%64))))<<(uint(v2317)%64) | v2315&int64(281474976710655)
	goto L296
L286:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2185))) = v2293
	*(*int64)(unsafe.Add(mBase, uint32(v2185)+8)) = v2294<<(uint(int64(48))%64) | v2240&int64(-9223372036854775807-1) | v2295
	m.G0 = v2238 + int32(16)
	goto L285
L287:
	;
	if base.B2i32(v2242 == int64(0)) == int32(0) {
		goto L291
	} else {
		goto L292
	}
L288:
	;
	if v2246 == int64(2047) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v2293 = v2242 << (uint(int64(60)) % 64)
	v2294 = int64(32767)
	v2295 = int64(base.Ui64(v2242) >> (uint(int64(4)) % 64))
	goto L286
L290:
	;
	v2293 = v2242 << (uint(int64(60)) % 64)
	v2294 = v2246 + int64(15360)
	v2295 = int64(base.Ui64(v2242) >> (uint(int64(4)) % 64))
	goto L286
L291:
	;
	if base.Ui64(v2242) < base.Ui64(int64(4294967296)) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	v2266 = int64(0)
	v2293 = v2266
	v2294 = v2266
	v2295 = v2266
	goto L286
L293:
	;
	v2280 = base.I32_clz(base.I32_wrap_i64(v2240)) | int32(32)
	goto L295
L294:
	;
	v2280 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v2242) >> (uint(int64(32)) % 64))))
	goto L295
L295:
	;
	F___ashlti3(m, v2238, v2242, int64(0), v2280+int32(49))
	mBase = m.M
	v2289 = *(*int64)(unsafe.Add(mBase, uint32(v2238+int32(8))))
	v2292 = *(*int64)(unsafe.Add(mBase, uint32(v2238)))
	v2293 = v2292
	v2294 = base.I64_extend_i32_u(int32(15372) - v2280)
	v2295 = v2289 ^ int64(281474976710656)
	goto L286
L296:
	;
	v2339 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(696))))
	v2340 = *(*int64)(unsafe.Add(mBase, uint32(v28)+688))
	v2342 = v28 + int32(640)
	v2343 = float64(1)
	v2345 = int32(113) - v2100
	if v2345 < int32(1024) {
		goto L299
	} else {
		goto L300
	}
L297:
	;
	v2393 = m.G0
	v2395 = v2393 - int32(16)
	m.G0 = v2395
	v2397 = base.I64_reinterpret_f64(base.F64_mul(v2378, base.F64_reinterpret_i64(base.I64_extend_i32_u(v2379+int32(1023))<<(uint(int64(52))%64))))
	v2399 = v2397 & int64(4503599627370495)
	v2403 = int64(base.Ui64(v2397)>>(uint(int64(52))%64)) & int64(2047)
	if v2403 == int64(0) {
		goto L314
	} else {
		goto L315
	}
L298:
	;
	goto L297
L299:
	;
	if int32(-1023) < v2345 {
		v2378 = v2343
		v2379 = v2345
		goto L298
	} else {
		goto L306
	}
L300:
	;
	v2349 = base.F64_mul(v2343, float64(8.98846567431158e+307))
	if base.Ui32(int32(2047)) <= base.Ui32(v2345) {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v2356 = int32(3069)
	if base.Ui32(v2345) < base.Ui32(v2356) {
		goto L303
	} else {
		goto L304
	}
L302:
	;
	v2378 = v2349
	v2379 = v2345 + int32(-1023)
	goto L298
L303:
	;
	v2359 = v2345
	goto L305
L304:
	;
	v2359 = v2356
	goto L305
L305:
	;
	v2378 = base.F64_mul(v2349, float64(8.98846567431158e+307))
	v2379 = v2359 + int32(-2046)
	goto L298
L306:
	;
	v2365 = base.F64_mul(v2343, float64(2.004168360008973e-292))
	if base.Ui32(v2345) <= base.Ui32(int32(-1992)) {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v2372 = int32(-2960)
	if base.Ui32(v2372) < base.Ui32(v2345) {
		goto L309
	} else {
		goto L310
	}
L308:
	;
	v2378 = v2365
	v2379 = v2345 + int32(969)
	goto L298
L309:
	;
	v2375 = v2345
	goto L311
L310:
	;
	v2375 = v2372
	goto L311
L311:
	;
	v2378 = base.F64_mul(v2365, float64(2.004168360008973e-292))
	v2379 = v2375 + int32(1938)
	goto L298
L312:
	;
	v2467 = *(*int64)(unsafe.Add(mBase, uint32(v28)+640))
	v2472 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(648))))
	F_fmodl(m, v28+int32(672), v2091, v2089, v2467, v2472)
	mBase = m.M
	v2475 = v28 + int32(624)
	v2476 = *(*int64)(unsafe.Add(mBase, uint32(v28)+672))
	v2481 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(680))))
	v2483 = m.G0
	v2484 = int32(16)
	v2485 = v2483 - v2484
	m.G0 = v2485
	F___addtf3(m, v2485, v2091, v2089, v2476, v2481^int64(-9223372036854775807-1))
	mBase = m.M
	v2490 = *(*int64)(unsafe.Add(mBase, uint32(v2485)))
	v2493 = *(*int64)(unsafe.Add(mBase, uint32(v2485+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v2475)+8)) = v2493
	*(*int64)(unsafe.Add(mBase, uint32(v2475))) = v2490
	m.G0 = v2485 + v2484
	goto L323
L313:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2342))) = v2450
	*(*int64)(unsafe.Add(mBase, uint32(v2342)+8)) = v2451<<(uint(int64(48))%64) | v2397&int64(-9223372036854775807-1) | v2452
	m.G0 = v2395 + int32(16)
	goto L312
L314:
	;
	if base.B2i32(v2399 == int64(0)) == int32(0) {
		goto L318
	} else {
		goto L319
	}
L315:
	;
	if v2403 == int64(2047) {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v2450 = v2399 << (uint(int64(60)) % 64)
	v2451 = int64(32767)
	v2452 = int64(base.Ui64(v2399) >> (uint(int64(4)) % 64))
	goto L313
L317:
	;
	v2450 = v2399 << (uint(int64(60)) % 64)
	v2451 = v2403 + int64(15360)
	v2452 = int64(base.Ui64(v2399) >> (uint(int64(4)) % 64))
	goto L313
L318:
	;
	if base.Ui64(v2399) < base.Ui64(int64(4294967296)) {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	v2423 = int64(0)
	v2450 = v2423
	v2451 = v2423
	v2452 = v2423
	goto L313
L320:
	;
	v2437 = base.I32_clz(base.I32_wrap_i64(v2397)) | int32(32)
	goto L322
L321:
	;
	v2437 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v2399) >> (uint(int64(32)) % 64))))
	goto L322
L322:
	;
	F___ashlti3(m, v2395, v2399, int64(0), v2437+int32(49))
	mBase = m.M
	v2446 = *(*int64)(unsafe.Add(mBase, uint32(v2395+int32(8))))
	v2449 = *(*int64)(unsafe.Add(mBase, uint32(v2395)))
	v2450 = v2449
	v2451 = base.I64_extend_i32_u(int32(15372) - v2437)
	v2452 = v2446 ^ int64(281474976710656)
	goto L313
L323:
	;
	v2501 = *(*int64)(unsafe.Add(mBase, uint32(v28)+624))
	v2506 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(632))))
	F___addtf3(m, v28+int32(608), v2340, v2339, v2501, v2506)
	mBase = m.M
	v2512 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(616))))
	v2513 = *(*int64)(unsafe.Add(mBase, uint32(v28)+608))
	v2514 = v2476
	v2515 = v2512
	v2516 = v2513
	v2517 = v2481
	v2518 = v2340
	v2519 = v2339
	goto L217
L324:
	;
	F___addtf3(m, v28+int32(432), v2516, v2515, v3027, v3029)
	mBase = m.M
	v3035 = v28 + int32(416)
	v3036 = *(*int64)(unsafe.Add(mBase, uint32(v28)+432))
	v3041 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(440))))
	v3043 = m.G0
	v3044 = int32(16)
	v3045 = v3043 - v3044
	m.G0 = v3045
	F___addtf3(m, v3045, v3036, v3041, v2518, v2519^int64(-9223372036854775807-1))
	mBase = m.M
	v3050 = *(*int64)(unsafe.Add(mBase, uint32(v3045)))
	v3053 = *(*int64)(unsafe.Add(mBase, uint32(v3045+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v3035)+8)) = v3053
	*(*int64)(unsafe.Add(mBase, uint32(v3035))) = v3050
	m.G0 = v3045 + v3044
	goto L406
L325:
	;
	v2530 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(784)+v2523<<(uint(int32(2))%32))))
	if base.Ui32(int32(499999999)) < base.Ui32(v2530) {
		goto L327
	} else {
		goto L328
	}
L326:
	;
	if base.Ui32(int32(111)) < base.Ui32(v2100) {
		v3027 = v2936
		v3029 = v2937
		goto L324
	} else {
		goto L380
	}
L327:
	;
	if v2530 == int32(500000000) {
		goto L343
	} else {
		goto L344
	}
L328:
	;
	if v2530 != 0 {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v2539 = v28 + int32(496)
	v2549 = m.G0
	v2551 = v2549 - int32(16)
	m.G0 = v2551
	v2553 = base.I64_reinterpret_f64(base.F64_mul(base.F64_convert_i32_s(l5), float64(0.25)))
	v2555 = v2553 & int64(4503599627370495)
	v2559 = int64(base.Ui64(v2553)>>(uint(int64(52))%64)) & int64(2047)
	if v2559 == int64(0) {
		goto L334
	} else {
		goto L335
	}
L330:
	;
	if (v1844+int32(5))&int32(2047) == v1952 {
		v2936 = v2514
		v2937 = v2517
		goto L326
	} else {
		goto L331
	}
L331:
	;
	goto L329
L332:
	;
	v2623 = *(*int64)(unsafe.Add(mBase, uint32(v28)+496))
	v2628 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(504))))
	F___addtf3(m, v28+int32(480), v2514, v2517, v2623, v2628)
	mBase = m.M
	v2634 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(488))))
	v2635 = *(*int64)(unsafe.Add(mBase, uint32(v28)+480))
	v2936 = v2635
	v2937 = v2634
	goto L326
L333:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2539))) = v2606
	*(*int64)(unsafe.Add(mBase, uint32(v2539)+8)) = v2607<<(uint(int64(48))%64) | v2553&int64(-9223372036854775807-1) | v2608
	m.G0 = v2551 + int32(16)
	goto L332
L334:
	;
	if base.B2i32(v2555 == int64(0)) == int32(0) {
		goto L338
	} else {
		goto L339
	}
L335:
	;
	if v2559 == int64(2047) {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v2606 = v2555 << (uint(int64(60)) % 64)
	v2607 = int64(32767)
	v2608 = int64(base.Ui64(v2555) >> (uint(int64(4)) % 64))
	goto L333
L337:
	;
	v2606 = v2555 << (uint(int64(60)) % 64)
	v2607 = v2559 + int64(15360)
	v2608 = int64(base.Ui64(v2555) >> (uint(int64(4)) % 64))
	goto L333
L338:
	;
	if base.Ui64(v2555) < base.Ui64(int64(4294967296)) {
		goto L340
	} else {
		goto L341
	}
L339:
	;
	v2579 = int64(0)
	v2606 = v2579
	v2607 = v2579
	v2608 = v2579
	goto L333
L340:
	;
	v2593 = base.I32_clz(base.I32_wrap_i64(v2553)) | int32(32)
	goto L342
L341:
	;
	v2593 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v2555) >> (uint(int64(32)) % 64))))
	goto L342
L342:
	;
	F___ashlti3(m, v2551, v2555, int64(0), v2593+int32(49))
	mBase = m.M
	v2602 = *(*int64)(unsafe.Add(mBase, uint32(v2551+int32(8))))
	v2605 = *(*int64)(unsafe.Add(mBase, uint32(v2551)))
	v2606 = v2605
	v2607 = base.I64_extend_i32_u(int32(15372) - v2593)
	v2608 = v2602 ^ int64(281474976710656)
	goto L333
L343:
	;
	v2736 = base.F64_convert_i32_s(l5)
	if (v1844+int32(5))&int32(2047) != v1952 {
		goto L356
	} else {
		goto L357
	}
L344:
	;
	v2639 = v28 + int32(592)
	v2649 = m.G0
	v2651 = v2649 - int32(16)
	m.G0 = v2651
	v2653 = base.I64_reinterpret_f64(base.F64_mul(base.F64_convert_i32_s(l5), float64(0.75)))
	v2655 = v2653 & int64(4503599627370495)
	v2659 = int64(base.Ui64(v2653)>>(uint(int64(52))%64)) & int64(2047)
	if v2659 == int64(0) {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	v2723 = *(*int64)(unsafe.Add(mBase, uint32(v28)+592))
	v2728 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(600))))
	F___addtf3(m, v28+int32(576), v2514, v2517, v2723, v2728)
	mBase = m.M
	v2734 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(584))))
	v2735 = *(*int64)(unsafe.Add(mBase, uint32(v28)+576))
	v2936 = v2735
	v2937 = v2734
	goto L326
L346:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2639))) = v2706
	*(*int64)(unsafe.Add(mBase, uint32(v2639)+8)) = v2707<<(uint(int64(48))%64) | v2653&int64(-9223372036854775807-1) | v2708
	m.G0 = v2651 + int32(16)
	goto L345
L347:
	;
	if base.B2i32(v2655 == int64(0)) == int32(0) {
		goto L351
	} else {
		goto L352
	}
L348:
	;
	if v2659 == int64(2047) {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v2706 = v2655 << (uint(int64(60)) % 64)
	v2707 = int64(32767)
	v2708 = int64(base.Ui64(v2655) >> (uint(int64(4)) % 64))
	goto L346
L350:
	;
	v2706 = v2655 << (uint(int64(60)) % 64)
	v2707 = v2659 + int64(15360)
	v2708 = int64(base.Ui64(v2655) >> (uint(int64(4)) % 64))
	goto L346
L351:
	;
	if base.Ui64(v2655) < base.Ui64(int64(4294967296)) {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	v2679 = int64(0)
	v2706 = v2679
	v2707 = v2679
	v2708 = v2679
	goto L346
L353:
	;
	v2693 = base.I32_clz(base.I32_wrap_i64(v2653)) | int32(32)
	goto L355
L354:
	;
	v2693 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v2655) >> (uint(int64(32)) % 64))))
	goto L355
L355:
	;
	F___ashlti3(m, v2651, v2655, int64(0), v2693+int32(49))
	mBase = m.M
	v2702 = *(*int64)(unsafe.Add(mBase, uint32(v2651+int32(8))))
	v2705 = *(*int64)(unsafe.Add(mBase, uint32(v2651)))
	v2706 = v2705
	v2707 = base.I64_extend_i32_u(int32(15372) - v2693)
	v2708 = v2702 ^ int64(281474976710656)
	goto L346
L356:
	;
	v2840 = v28 + int32(560)
	v2849 = m.G0
	v2851 = v2849 - int32(16)
	m.G0 = v2851
	v2853 = base.I64_reinterpret_f64(base.F64_mul(v2736, float64(0.75)))
	v2855 = v2853 & int64(4503599627370495)
	v2859 = int64(base.Ui64(v2853)>>(uint(int64(52))%64)) & int64(2047)
	if v2859 == int64(0) {
		goto L371
	} else {
		goto L372
	}
L357:
	;
	v2743 = v28 + int32(528)
	v2752 = m.G0
	v2754 = v2752 - int32(16)
	m.G0 = v2754
	v2756 = base.I64_reinterpret_f64(base.F64_mul(v2736, float64(0.5)))
	v2758 = v2756 & int64(4503599627370495)
	v2762 = int64(base.Ui64(v2756)>>(uint(int64(52))%64)) & int64(2047)
	if v2762 == int64(0) {
		goto L360
	} else {
		goto L361
	}
L358:
	;
	v2826 = *(*int64)(unsafe.Add(mBase, uint32(v28)+528))
	v2831 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(536))))
	F___addtf3(m, v28+int32(512), v2514, v2517, v2826, v2831)
	mBase = m.M
	v2837 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(520))))
	v2838 = *(*int64)(unsafe.Add(mBase, uint32(v28)+512))
	v2936 = v2838
	v2937 = v2837
	goto L326
L359:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2743))) = v2809
	*(*int64)(unsafe.Add(mBase, uint32(v2743)+8)) = v2810<<(uint(int64(48))%64) | v2756&int64(-9223372036854775807-1) | v2811
	m.G0 = v2754 + int32(16)
	goto L358
L360:
	;
	if base.B2i32(v2758 == int64(0)) == int32(0) {
		goto L364
	} else {
		goto L365
	}
L361:
	;
	if v2762 == int64(2047) {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v2809 = v2758 << (uint(int64(60)) % 64)
	v2810 = int64(32767)
	v2811 = int64(base.Ui64(v2758) >> (uint(int64(4)) % 64))
	goto L359
L363:
	;
	v2809 = v2758 << (uint(int64(60)) % 64)
	v2810 = v2762 + int64(15360)
	v2811 = int64(base.Ui64(v2758) >> (uint(int64(4)) % 64))
	goto L359
L364:
	;
	if base.Ui64(v2758) < base.Ui64(int64(4294967296)) {
		goto L366
	} else {
		goto L367
	}
L365:
	;
	v2782 = int64(0)
	v2809 = v2782
	v2810 = v2782
	v2811 = v2782
	goto L359
L366:
	;
	v2796 = base.I32_clz(base.I32_wrap_i64(v2756)) | int32(32)
	goto L368
L367:
	;
	v2796 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v2758) >> (uint(int64(32)) % 64))))
	goto L368
L368:
	;
	F___ashlti3(m, v2754, v2758, int64(0), v2796+int32(49))
	mBase = m.M
	v2805 = *(*int64)(unsafe.Add(mBase, uint32(v2754+int32(8))))
	v2808 = *(*int64)(unsafe.Add(mBase, uint32(v2754)))
	v2809 = v2808
	v2810 = base.I64_extend_i32_u(int32(15372) - v2796)
	v2811 = v2805 ^ int64(281474976710656)
	goto L359
L369:
	;
	v2923 = *(*int64)(unsafe.Add(mBase, uint32(v28)+560))
	v2928 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(568))))
	F___addtf3(m, v28+int32(544), v2514, v2517, v2923, v2928)
	mBase = m.M
	v2934 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(552))))
	v2935 = *(*int64)(unsafe.Add(mBase, uint32(v28)+544))
	v2936 = v2935
	v2937 = v2934
	goto L326
L370:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2840))) = v2906
	*(*int64)(unsafe.Add(mBase, uint32(v2840)+8)) = v2907<<(uint(int64(48))%64) | v2853&int64(-9223372036854775807-1) | v2908
	m.G0 = v2851 + int32(16)
	goto L369
L371:
	;
	if base.B2i32(v2855 == int64(0)) == int32(0) {
		goto L375
	} else {
		goto L376
	}
L372:
	;
	if v2859 == int64(2047) {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v2906 = v2855 << (uint(int64(60)) % 64)
	v2907 = int64(32767)
	v2908 = int64(base.Ui64(v2855) >> (uint(int64(4)) % 64))
	goto L370
L374:
	;
	v2906 = v2855 << (uint(int64(60)) % 64)
	v2907 = v2859 + int64(15360)
	v2908 = int64(base.Ui64(v2855) >> (uint(int64(4)) % 64))
	goto L370
L375:
	;
	if base.Ui64(v2855) < base.Ui64(int64(4294967296)) {
		goto L377
	} else {
		goto L378
	}
L376:
	;
	v2879 = int64(0)
	v2906 = v2879
	v2907 = v2879
	v2908 = v2879
	goto L370
L377:
	;
	v2893 = base.I32_clz(base.I32_wrap_i64(v2853)) | int32(32)
	goto L379
L378:
	;
	v2893 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v2855) >> (uint(int64(32)) % 64))))
	goto L379
L379:
	;
	F___ashlti3(m, v2851, v2855, int64(0), v2893+int32(49))
	mBase = m.M
	v2902 = *(*int64)(unsafe.Add(mBase, uint32(v2851+int32(8))))
	v2905 = *(*int64)(unsafe.Add(mBase, uint32(v2851)))
	v2906 = v2905
	v2907 = base.I64_extend_i32_u(int32(15372) - v2893)
	v2908 = v2902 ^ int64(281474976710656)
	goto L370
L380:
	;
	v2943 = int64(0)
	F_fmodl(m, v28+int32(464), v2936, v2937, v2943, int64(4611404543450677248))
	mBase = m.M
	v2946 = *(*int64)(unsafe.Add(mBase, uint32(v28)+464))
	v2951 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(472))))
	v2961 = v2951 & int64(9223372036854775807)
	v2962 = int64(9223090561878065152)
	if v2961 == v2962 {
		goto L383
	} else {
		goto L384
	}
L381:
	;
	if v3015 != 0 {
		v3027 = v2936
		v3029 = v2937
		goto L324
	} else {
		goto L405
	}
L382:
	;
	v3015 = v3011
	goto L381
L383:
	;
	v2966 = base.B2i32(v2946 != v2943)
	goto L385
L384:
	;
	v2966 = base.B2i32(base.Ui64(v2962) < base.Ui64(v2961))
	goto L385
L385:
	;
	if v2966 != 0 {
		v3011 = int32(1)
		goto L382
	} else {
		goto L386
	}
L386:
	;
	goto L388
L388:
	;
	goto L389
L389:
	;
	goto L390
L390:
	;
	if base.B2i32(v2943|v2946|(int64(0)|v2961) == int64(0)) == int32(0) {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	if v2943&v2951 < int64(0) {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	v3015 = int32(0)
	goto L381
L393:
	;
	if v2951 == v2943 {
		goto L401
	} else {
		goto L402
	}
L394:
	;
	if v2951 == v2943 {
		goto L396
	} else {
		goto L397
	}
L395:
	;
	v3015 = base.B2i32(v2946^v2943|(v2951^v2943) != int64(0))
	goto L381
L396:
	;
	v2990 = base.B2i32(base.Ui64(v2946) < base.Ui64(v2943))
	goto L398
L397:
	;
	v2990 = base.B2i32(v2951 < v2943)
	goto L398
L398:
	;
	if v2990 == int32(0) {
		goto L395
	} else {
		goto L399
	}
L399:
	;
	v3015 = int32(-1)
	goto L381
L400:
	;
	v3011 = base.B2i32(v2946^v2943|(v2951^v2943) != int64(0))
	goto L382
L401:
	;
	v3002 = base.B2i32(base.Ui64(v2943) < base.Ui64(v2946))
	goto L403
L402:
	;
	v3002 = base.B2i32(v2943 < v2951)
	goto L403
L403:
	;
	if v3002 == int32(0) {
		goto L400
	} else {
		goto L404
	}
L404:
	;
	v3015 = int32(-1)
	goto L381
L405:
	;
	F___addtf3(m, v28+int32(448), v2936, v2937, int64(0), int64(4611404543450677248))
	mBase = m.M
	v3025 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(456))))
	v3026 = *(*int64)(unsafe.Add(mBase, uint32(v28)+448))
	v3027 = v3026
	v3029 = v3025
	goto L324
L406:
	;
	v3063 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(424))))
	v3064 = *(*int64)(unsafe.Add(mBase, uint32(v28)+416))
	if v2093&int32(2147483647) <= v33+int32(-2) {
		v3237 = v1843
		v3240 = v3063
		v3241 = v3064
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v3243 = v28 + int32(368)
	v3245 = m.G0
	v3247 = v3245 - int32(80)
	m.G0 = v3247
	if v3237 < int32(16384) {
		goto L468
	} else {
		goto L469
	}
L408:
	;
	v3071 = v28 + int32(400)
	*(*int64)(unsafe.Add(mBase, uint32(v3071)+8)) = v3063 & int64(9223372036854775807)
	*(*int64)(unsafe.Add(mBase, uint32(v3071))) = v3064
	goto L409
L409:
	;
	v3078 = int64(0)
	F___multf3(m, v28+int32(384), v3064, v3063, v3078, int64(4611123068473966592))
	mBase = m.M
	v3081 = *(*int64)(unsafe.Add(mBase, uint32(v28)+400))
	v3086 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(408))))
	v3088 = int64(4643211215818981376)
	v3092 = int32(-1)
	v3096 = v3086 & int64(9223372036854775807)
	v3097 = int64(9223090561878065152)
	if v3096 == v3097 {
		goto L412
	} else {
		goto L413
	}
L410:
	;
	v3149 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(392))))
	v3151 = base.B2i32(int32(-1) < v3144)
	if int32(-1) < v3144 {
		goto L432
	} else {
		goto L433
	}
L411:
	;
	v3144 = v3140
	goto L410
L412:
	;
	v3101 = base.B2i32(v3081 != v3078)
	goto L414
L413:
	;
	v3101 = base.B2i32(base.Ui64(v3097) < base.Ui64(v3096))
	goto L414
L414:
	;
	if v3101 != 0 {
		v3140 = v3092
		goto L411
	} else {
		goto L415
	}
L415:
	;
	goto L417
L417:
	;
	goto L418
L418:
	;
	goto L419
L419:
	;
	if base.B2i32(v3078|v3081|(int64(4643211215818981376)|v3096) == int64(0)) == int32(0) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	if v3088&v3086 < int64(0) {
		goto L422
	} else {
		goto L423
	}
L421:
	;
	v3144 = int32(0)
	goto L410
L422:
	;
	if v3086 == v3088 {
		goto L428
	} else {
		goto L429
	}
L423:
	;
	if v3086 == v3088 {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v3125 = base.B2i32(base.Ui64(v3081) < base.Ui64(v3078))
	goto L426
L425:
	;
	v3125 = base.B2i32(v3086 < v3088)
	goto L426
L426:
	;
	if v3125 != 0 {
		v3140 = v3092
		goto L411
	} else {
		goto L427
	}
L427:
	;
	v3144 = base.B2i32(v3081^v3078|(v3086^v3088) != int64(0))
	goto L410
L428:
	;
	v3134 = base.B2i32(base.Ui64(v3078) < base.Ui64(v3081))
	goto L430
L429:
	;
	v3134 = base.B2i32(v3088 < v3086)
	goto L430
L430:
	;
	if v3134 != 0 {
		v3140 = v3092
		goto L411
	} else {
		goto L431
	}
L431:
	;
	v3140 = base.B2i32(v3081^v3078|(v3086^v3088) != int64(0))
	goto L411
L432:
	;
	v3152 = v3149
	goto L434
L433:
	;
	v3152 = v3063
	goto L434
L434:
	;
	v3153 = *(*int64)(unsafe.Add(mBase, uint32(v28)+384))
	if int32(-1) < v3144 {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v3154 = v3153
	goto L437
L436:
	;
	v3154 = v3064
	goto L437
L437:
	;
	v3155 = int64(0)
	v3164 = v3029 & int64(9223372036854775807)
	v3165 = int64(9223090561878065152)
	if v3164 == v3165 {
		goto L440
	} else {
		goto L441
	}
L438:
	;
	v3219 = v1843 + v3151
	if v33 < v3219+int32(110) {
		goto L462
	} else {
		goto L463
	}
L439:
	;
	v3218 = v3214
	goto L438
L440:
	;
	v3169 = base.B2i32(v3027 != v3155)
	goto L442
L441:
	;
	v3169 = base.B2i32(base.Ui64(v3165) < base.Ui64(v3164))
	goto L442
L442:
	;
	if v3169 != 0 {
		v3214 = int32(1)
		goto L439
	} else {
		goto L443
	}
L443:
	;
	goto L445
L445:
	;
	goto L446
L446:
	;
	goto L447
L447:
	;
	if base.B2i32(v3155|v3027|(int64(0)|v3164) == int64(0)) == int32(0) {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	if v3155&v3029 < int64(0) {
		goto L450
	} else {
		goto L451
	}
L449:
	;
	v3218 = int32(0)
	goto L438
L450:
	;
	if v3029 == v3155 {
		goto L458
	} else {
		goto L459
	}
L451:
	;
	if v3029 == v3155 {
		goto L453
	} else {
		goto L454
	}
L452:
	;
	v3218 = base.B2i32(v3027^v3155|(v3029^v3155) != int64(0))
	goto L438
L453:
	;
	v3193 = base.B2i32(base.Ui64(v3027) < base.Ui64(v3155))
	goto L455
L454:
	;
	v3193 = base.B2i32(v3029 < v3155)
	goto L455
L455:
	;
	if v3193 == int32(0) {
		goto L452
	} else {
		goto L456
	}
L456:
	;
	v3218 = int32(-1)
	goto L438
L457:
	;
	v3214 = base.B2i32(v3027^v3155|(v3029^v3155) != int64(0))
	goto L439
L458:
	;
	v3205 = base.B2i32(base.Ui64(v3155) < base.Ui64(v3027))
	goto L460
L459:
	;
	v3205 = base.B2i32(v3155 < v3029)
	goto L460
L460:
	;
	if v3205 == int32(0) {
		goto L457
	} else {
		goto L461
	}
L461:
	;
	v3218 = int32(-1)
	goto L438
L462:
	;
	goto L465
L463:
	;
	v3224 = int32(0)
	if v2099&(base.B2i32(v2100 != v2094)|base.B2i32(v3144 < v3224))&base.B2i32(v3218 != v3224) == v3224 {
		v3237 = v3219
		v3240 = v3152
		v3241 = v3154
		goto L407
	} else {
		goto L464
	}
L464:
	;
	goto L462
L465:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_decfloat[1])) = int32(68)
	v3237 = v3219
	v3240 = v3152
	v3241 = v3154
	goto L407
L466:
	;
	v3340 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(376))))
	v3341 = *(*int64)(unsafe.Add(mBase, uint32(v28)+368))
	v3353 = v3340
	v3358 = v3341
	goto L24
L467:
	;
	F___multf3(m, v3247, v3317, v3318, int64(0), base.I64_extend_i32_u(v3319+int32(16383))<<(uint(int64(48))%64))
	mBase = m.M
	v3329 = *(*int64)(unsafe.Add(mBase, uint32(v3247+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v3243)+8)) = v3329
	v3331 = *(*int64)(unsafe.Add(mBase, uint32(v3247)))
	*(*int64)(unsafe.Add(mBase, uint32(v3243))) = v3331
	m.G0 = v3247 + int32(80)
	goto L466
L468:
	;
	if int32(-16383) < v3237 {
		v3317 = v3241
		v3318 = v3240
		v3319 = v3237
		goto L467
	} else {
		goto L475
	}
L469:
	;
	F___multf3(m, v3247+int32(32), v3241, v3240, int64(0), int64(9222809086901354496))
	mBase = m.M
	v3260 = *(*int64)(unsafe.Add(mBase, uint32(v3247+int32(40))))
	v3261 = *(*int64)(unsafe.Add(mBase, uint32(v3247)+32))
	if base.Ui32(int32(32767)) <= base.Ui32(v3237) {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	F___multf3(m, v3247+int32(16), v3261, v3260, int64(0), int64(9222809086901354496))
	mBase = m.M
	v3271 = int32(49149)
	if base.Ui32(v3237) < base.Ui32(v3271) {
		goto L472
	} else {
		goto L473
	}
L471:
	;
	v3317 = v3261
	v3318 = v3260
	v3319 = v3237 + int32(-16383)
	goto L467
L472:
	;
	v3274 = v3237
	goto L474
L473:
	;
	v3274 = v3271
	goto L474
L474:
	;
	v3281 = *(*int64)(unsafe.Add(mBase, uint32(v3247+int32(24))))
	v3282 = *(*int64)(unsafe.Add(mBase, uint32(v3247)+16))
	v3317 = v3282
	v3318 = v3281
	v3319 = v3274 + int32(-32766)
	goto L467
L475:
	;
	F___multf3(m, v3247+int32(64), v3241, v3240, int64(0), int64(32088147345014784))
	mBase = m.M
	v3294 = *(*int64)(unsafe.Add(mBase, uint32(v3247+int32(72))))
	v3295 = *(*int64)(unsafe.Add(mBase, uint32(v3247)+64))
	if base.Ui32(v3237) <= base.Ui32(int32(-32652)) {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	F___multf3(m, v3247+int32(48), v3295, v3294, int64(0), int64(32088147345014784))
	mBase = m.M
	v3305 = int32(-48920)
	if base.Ui32(v3305) < base.Ui32(v3237) {
		goto L478
	} else {
		goto L479
	}
L477:
	;
	v3317 = v3295
	v3318 = v3294
	v3319 = v3237 + int32(16269)
	goto L467
L478:
	;
	v3308 = v3237
	goto L480
L479:
	;
	v3308 = v3305
	goto L480
L480:
	;
	v3315 = *(*int64)(unsafe.Add(mBase, uint32(v3247+int32(56))))
	v3316 = *(*int64)(unsafe.Add(mBase, uint32(v3247)+48))
	v3317 = v3316
	v3318 = v3315
	v3319 = v3308 + int32(32538)
	goto L467
}
func F_decrCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_incrDecrCommand(m, l0, int64(-1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_delKeysInSlot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v144 int64
	_ = v144
	var v146 int32
	_ = v146
	var v150 int64
	_ = v150
	var v154 int64
	_ = v154
	var v157 int32
	_ = v157
	var v165 int64
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int64
	_ = v252
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[0]))
	if int32(1) <= v25 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v314
L2:
	;
	v68 = int32(_a_F_delKeysInSlot_0)
	*(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[1])) = int32(1)
	v71 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[0]))
	if v71 < v73 {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	if v62 != 0 {
		goto L2
	} else {
		goto L12
	}
L4:
	;
	goto L3
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[2]))
	v31 = int32(0)
	v34 = v25
	v35 = v31
	v36 = v30
	v37 = v31
	goto L7
L6:
	;
	v62 = int32(0)
	goto L4
L7:
	;
	v40 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v36+v37<<(uint(int32(2))%32))))
	if v44 == v40 {
		v53 = v34
		v54 = v36
		v55 = v40
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v62 = v56
	goto L4
L9:
	;
	v56 = v55 + v35
	v58 = v37 + int32(1)
	if v58 < v53 {
		v34 = v53
		v35 = v56
		v36 = v54
		v37 = v58
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v48 = F_kvstoreHashtableSize(m, v47, l0)
	mBase = m.M
	v49 = int32(_a_F_delKeysInSlot_0)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[0]))
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[2]))
	v53 = v50
	v54 = v52
	v55 = v48
	goto L9
L11:
	;
	goto L8
L12:
	;
	v314 = int32(0)
	goto L1
L13:
	;
	v79 = int32(_a_F_delKeysInSlot_0)
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[3]))
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[2]))
	v83 = int32(0)
	v90 = v83
	v91 = v73
	v93 = v82
	v94 = v83
	goto L15
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[1])) = int32(0)
	v314 = v71
	goto L1
L15:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94<<(uint(int32(2))%32))))
	if v101 == int32(0) {
		v286 = v90
		v287 = v91
		v289 = v93
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v297 = int32(_a_F_delKeysInSlot_0)
	*(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[1])) = int32(0)
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[3]))
	if v301 == v80 {
		v314 = v286
		goto L1
	} else {
		goto L63
	}
L17:
	;
	v295 = v94 + int32(1)
	if v295 < v287 {
		v90 = v286
		v91 = v287
		v93 = v289
		v94 = v295
		goto L15
	} else {
		goto L62
	}
L18:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v106 = F_kvstoreGetHashtableIterator(m, v104, l0, int32(1))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	F_kvstoreReleaseHashtableIterator(m, v106)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L20
	} else {
		goto L61
	}
L20:
	;
	return int32(0)
L21:
	;
	v112 = F_kvstoreHashtableIteratorNext(m, v106, v16+int32(12))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if v112 == int32(0) {
		v267 = v90
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v121 = v90
	goto L24
L24:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v132 = int32(0)
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[3])) = v136 + int32(1)
	goto L28
L25:
	;
	v267 = v257
	goto L19
L26:
	;
	v169 = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v173&int32(2) == v169 {
		v193 = v169
		goto L39
	} else {
		goto L40
	}
L27:
	;
	goto L26
L28:
	;
	if v136 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	goto L31
L30:
	;
	v146 = int32(0)
	*(*int64)(unsafe.Add(mBase, _c_F_delKeysInSlot[4])) = v144
	v150 = base.I64_div_s(v144, int64(1000))
	*(*int64)(unsafe.Add(mBase, _c_F_delKeysInSlot[5])) = v150
	v154 = base.I64_div_s(v144, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _c_F_delKeysInSlot[6])) = v154
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[7]))
	F_lrulfu_updateClockAndPolicy(m, v150, int32(base.Ui32(v157&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v165 = *(*int64)(unsafe.Add(mBase, _c_F_delKeysInSlot[5]))
	*(*int64)(unsafe.Add(mBase, _c_F_delKeysInSlot[8])) = v165
	goto L27
L31:
	;
	v144 = F_ustime(m)
	mBase = m.M
	goto L30
L32:
	;
	v214 = F_createStringObject_1(m, v193, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L20
	} else {
		goto L41
	}
L33:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v193+int32(-17))))
	v213 = v212
	goto L32
L34:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v193+int32(-9))))
	v213 = v209
	goto L32
L35:
	;
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193+int32(-5)))))
	v213 = v206
	goto L32
L36:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+int32(-3)))))
	v213 = v203
	goto L32
L37:
	;
	v213 = int32(base.Ui32(v196) >> (uint(int32(3)) % 32))
	goto L32
L38:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+int32(-1)))))
	switch v196 & int32(7) {
	case 0:
		goto L37
	case 1:
		goto L36
	case 2:
		goto L35
	case 3:
		goto L34
	case 4:
		goto L33
	default:
		v213 = v169
		goto L32
	}
L39:
	;
	goto L38
L40:
	;
	v187 = v129 + (v173&int32(4) ^ int32(12)) + v173<<(uint(int32(3))%32)&int32(8)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v193 = v187 + v188 + int32(1)
	goto L39
L41:
	;
	if l1 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if l2 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v220 = F_dbSyncDelete(m, v101, v214)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L20
	} else {
		goto L46
	}
L44:
	;
	v218 = F_dbAsyncDelete(m, v101, v214)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L20
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	goto L42
L47:
	;
	F_signalModifiedKey(m, int32(0), v101, v214)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L20
	} else {
		goto L50
	}
L48:
	;
	F_propagateDeletion(m, v101, v214, l1, l0)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L20
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	if l3 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v240 = int32(0)
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[3])) = v242 + int32(-1)
	goto L56
L52:
	;
	F_moduleNotifyKeyspaceEvent(m, int32(4), int32(_a_F_delKeysInSlot_1), v214, v229)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L20
	} else {
		goto L55
	}
L53:
	;
	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_delKeysInSlot_1), v214, v229)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L20
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	goto L51
L56:
	;
	F_postExecutionUnitOperations(m)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L20
	} else {
		goto L57
	}
L57:
	;
	F_decrRefCount(m, v214)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	v250 = int32(_a_F_delKeysInSlot_0)
	v252 = *(*int64)(unsafe.Add(mBase, _c_F_delKeysInSlot[9]))
	*(*int64)(unsafe.Add(mBase, _c_F_delKeysInSlot[9])) = v252 + int64(1)
	v257 = v121 + int32(1)
	v260 = F_kvstoreHashtableIteratorNext(m, v106, v16+int32(12))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L20
	} else {
		goto L59
	}
L59:
	;
	if v260 != 0 {
		v121 = v257
		goto L24
	} else {
		goto L60
	}
L60:
	;
	goto L25
L61:
	;
	v277 = int32(_a_F_delKeysInSlot_0)
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[0]))
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_delKeysInSlot[2]))
	v286 = v267
	v287 = v278
	v289 = v280
	goto L17
L62:
	;
	goto L16
L63:
	;
	F__serverAssert(m, int32(_a_F_delKeysInSlot_2), int32(_a_F_delKeysInSlot_3), int32(7503))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L20
	} else {
		goto L64
	}
L64:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_deleteCachedResponseClient(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_valkey_free(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
		v7 = F_freeClient(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_deleteExpiredKeyAndPropagateWithDictIndex(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int64
	_ = v84
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[0]))
	if base.B2i32(v15 == int64(0)) == int32(0) {
		v21 = F_ustime(m)
		mBase = m.M
		v22 = v21
	} else {
		v22 = int64(0)
	}
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[1]))
	v26 = F_dbGenericDeleteWithDictIndex(m, l0, l1, v24, int32(2), l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return
	} else {
		v29 = *(*int64)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[0]))
		if v29 == int64(0) {
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			F_notifyKeyspaceEvent(m, int32(256), int32(_a_F_deleteExpiredKeyAndPropagateWithDictIndex_0), l1, v49)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				F_touchWatchedKey(m, l0, l1)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					F_trackingInvalidateKey(m, int32(0), l1, int32(1))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						v58 = int32(_a_F_deleteExpiredKeyAndPropagateWithDictIndex_1)
						v59 = *(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[2]))
						*(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[2])) = int32(1)
						v64 = *(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[1]))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
						if v64 != 0 {
							v69 = int32(244)
						} else {
							v69 = int32(240)
						}
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_c_F_deleteExpiredKeyAndPropagateWithDictIndex[3])))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v71
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						F_alsoPropagate(m, v73, v12+int32(8), int32(2), int32(3), l2)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							v80 = int32(_a_F_deleteExpiredKeyAndPropagateWithDictIndex_1)
							*(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[2])) = v59
							v84 = *(*int64)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[4]))
							*(*int64)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[4])) = v84 + int64(1)
							m.G0 = v12 + int32(16)
							return
						}
					}
				}
			}
		} else {
			v32 = F_ustime(m)
			mBase = m.M
			v34 = *(*int64)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[0]))
			if v34 == int64(0) {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				F_notifyKeyspaceEvent(m, int32(256), int32(_a_F_deleteExpiredKeyAndPropagateWithDictIndex_0), l1, v49)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					F_touchWatchedKey(m, l0, l1)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_trackingInvalidateKey(m, int32(0), l1, int32(1))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							v58 = int32(_a_F_deleteExpiredKeyAndPropagateWithDictIndex_1)
							v59 = *(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[2]))
							*(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[2])) = int32(1)
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[1]))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
							if v64 != 0 {
								v69 = int32(244)
							} else {
								v69 = int32(240)
							}
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_c_F_deleteExpiredKeyAndPropagateWithDictIndex[3])))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v71
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							F_alsoPropagate(m, v73, v12+int32(8), int32(2), int32(3), l2)
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								v80 = int32(_a_F_deleteExpiredKeyAndPropagateWithDictIndex_1)
								*(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[2])) = v59
								v84 = *(*int64)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[4]))
								*(*int64)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[4])) = v84 + int64(1)
								m.G0 = v12 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v37 = v32 - v22
				if v37 < v34*int64(1000) {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					F_notifyKeyspaceEvent(m, int32(256), int32(_a_F_deleteExpiredKeyAndPropagateWithDictIndex_0), l1, v49)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						F_touchWatchedKey(m, l0, l1)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							F_trackingInvalidateKey(m, int32(0), l1, int32(1))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								v58 = int32(_a_F_deleteExpiredKeyAndPropagateWithDictIndex_1)
								v59 = *(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[2]))
								*(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[2])) = int32(1)
								v64 = *(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[1]))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
								if v64 != 0 {
									v69 = int32(244)
								} else {
									v69 = int32(240)
								}
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_c_F_deleteExpiredKeyAndPropagateWithDictIndex[3])))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v71
								v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								F_alsoPropagate(m, v73, v12+int32(8), int32(2), int32(3), l2)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									v80 = int32(_a_F_deleteExpiredKeyAndPropagateWithDictIndex_1)
									*(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[2])) = v59
									v84 = *(*int64)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[4]))
									*(*int64)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[4])) = v84 + int64(1)
									m.G0 = v12 + int32(16)
									return
								}
							}
						}
					}
				} else {
					F_latencyAddSample(m, int32(_a_F_deleteExpiredKeyAndPropagateWithDictIndex_2), v37)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						F_notifyKeyspaceEvent(m, int32(256), int32(_a_F_deleteExpiredKeyAndPropagateWithDictIndex_0), l1, v49)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							F_touchWatchedKey(m, l0, l1)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								F_trackingInvalidateKey(m, int32(0), l1, int32(1))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									v58 = int32(_a_F_deleteExpiredKeyAndPropagateWithDictIndex_1)
									v59 = *(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[2]))
									*(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[2])) = int32(1)
									v64 = *(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[1]))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
									if v64 != 0 {
										v69 = int32(244)
									} else {
										v69 = int32(240)
									}
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_c_F_deleteExpiredKeyAndPropagateWithDictIndex[3])))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v71
									v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									F_alsoPropagate(m, v73, v12+int32(8), int32(2), int32(3), l2)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										v80 = int32(_a_F_deleteExpiredKeyAndPropagateWithDictIndex_1)
										*(*int32)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[2])) = v59
										v84 = *(*int64)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[4]))
										*(*int64)(unsafe.Add(mBase, _c_F_deleteExpiredKeyAndPropagateWithDictIndex[4])) = v84 + int64(1)
										m.G0 = v12 + int32(16)
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
func F_detectAndUpdateCachedNodeHealth(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_detectAndUpdateCachedNodeHealth[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+12)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967295)
	goto L1
L1:
	;
	v21 = int32(0)
	goto L3
L2:
	;
	m.G0 = v7 + int32(32)
	return v21
L3:
	;
	v31 = v7 + int32(20)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v32 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v127 == int32(0) {
		goto L2
	} else {
		goto L31
	}
L6:
	;
	v38 = v31
	v39 = v35
	goto L9
L7:
	;
	v35 = int32(1)
	goto L6
L8:
	;
	v35 = int32(0)
	goto L6
L9:
	;
	switch v39 {
	case 0:
		goto L14
	default:
		goto L13
	}
L11:
	;
	v39 = int32(0)
	goto L9
L12:
	;
	goto L5
L13:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v119
	if v119 == int32(0) {
		goto L11
	} else {
		goto L30
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v43 != int32(-1) {
		v82 = v43
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v83 = int32(1)
	v84 = v82 + v83
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v84
	v86 = int32(0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+v90+int32(26)))))
	if v94 == int32(255) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v47 != 0 {
		v82 = int32(-1)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v49 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	if v76 != int32(-1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v56 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v48)+16)))
	v57 = int64(*(*int8)(unsafe.Add(mBase, uint32(v48)+27)))
	v58 = int64(*(*int32)(unsafe.Add(mBase, uint32(v48)+8)))
	v59 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v48)+12)))
	v60 = int64(*(*int8)(unsafe.Add(mBase, uint32(v48)+26)))
	v61 = int64(*(*int32)(unsafe.Add(mBase, uint32(v48)+4)))
	v62 = F_wangHash64(m, v61)
	mBase = m.M
	v64 = F_wangHash64(m, v60+v62)
	mBase = m.M
	v66 = F_wangHash64(m, v59+v64)
	mBase = m.M
	v68 = F_wangHash64(m, v58+v66)
	mBase = m.M
	v70 = F_wangHash64(m, v57+v68)
	mBase = m.M
	v72 = F_wangHash64(m, v56+v70)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v75 = v74
	goto L18
L20:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+24)))
	v54 = v52 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+24)) = uint16(v54)
	v75 = v48
	goto L18
L21:
	;
	v82 = v76 + int32(-1)
	goto L15
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v82 = v79
	goto L15
L23:
	;
	v109 = int32(2)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v89+v107<<(uint(v109)%32)+int32(4))))
	v38 = v114 + v108<<(uint(v109)%32)
	v39 = int32(1)
	goto L9
L24:
	;
	v98 = v86
	goto L26
L25:
	;
	v98 = v83 << (uint(v94) % 32)
	goto L26
L26:
	;
	if v84 < v98 {
		v107 = v90
		v108 = v84
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if v90 != 0 {
		v127 = v86
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v100 == int32(-1) {
		v127 = v86
		goto L12
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967296)
	v107 = int32(1)
	v108 = int32(0)
	goto L23
L30:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v123
	v127 = v119
	goto L12
L31:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	goto L32
L32:
	;
	v136 = F_clusterNodeIsFailing(m, v133)
	mBase = m.M
	if v136 != 0 {
		v140 = int32(0)
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v133)+2356))
	if v140 == v141 {
		goto L3
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	v137 = F_getNodeReplicationOffset(m, v133)
	mBase = m.M
	v140 = base.B2i32(v137 != int64(0))
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+2356)) = v140
	v21 = int32(1)
	goto L3
}
func F_die(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1
	v9 = m.G397
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = F_vfprintf(m, v10, l0, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v14 = F_fputc(m, int32(10), v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
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
func F_digits10(m *base.Module, l0 int64) int32 {
	var v2 int32
	_ = v2
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	v2 = int32(0)
	if base.Ui64(l0) < base.Ui64(int64(10)) {
		v68 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1) + v68
L2:
	;
	v7 = l0
	v8 = v2
	goto L3
L3:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v7) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v68 = v62
	goto L1
L5:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v7) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	return int32(2) + v8
L7:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v7) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return int32(3) + v8
L9:
	;
	v62 = v8 + int32(12)
	v66 = base.I64_div_u_s(v7, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v7) {
		v7 = v66
		v8 = v62
		goto L3
	} else {
		goto L31
	}
L10:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v7) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v7) {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v7) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v7) {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v7) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v7) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	return int32(4) + v8
L17:
	;
	v35 = int32(6)
	goto L19
L18:
	;
	v35 = int32(5)
	goto L19
L19:
	;
	return v35 + v8
L20:
	;
	v42 = int32(8)
	goto L22
L21:
	;
	v42 = int32(7)
	goto L22
L22:
	;
	return v42 + v8
L23:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v7) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v7) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v51 = int32(10)
	goto L27
L26:
	;
	v51 = int32(9)
	goto L27
L27:
	;
	return v51 + v8
L28:
	;
	v58 = int32(12)
	goto L30
L29:
	;
	v58 = int32(11)
	goto L30
L30:
	;
	return v58 + v8
L31:
	;
	goto L4
}
func F_dirCreateIfMissing(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(96)
	m.G0 = v7
	v11 = F_mkdir(m, l0, int32(493))
	mBase = m.M
	if v11 == v2 {
		v29 = v2
	} else {
		v14 = int32(9116376)
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_dirCreateIfMissing[0]))
		if v15 != int32(20) {
			v29 = int32(-1)
		} else {
			v20 = F___fstatat(m, int32(-100), l0, v7, int32(0))
			mBase = m.M
			if v20 != 0 {
				*(*int32)(unsafe.Add(mBase, _c_F_dirCreateIfMissing[0])) = int32(54)
				v29 = int32(-1)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				if v21&int32(61440) == int32(16384) {
					v29 = v2
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_dirCreateIfMissing[0])) = int32(54)
					v29 = int32(-1)
				}
			}
		}
	}
	m.G0 = v7 + int32(96)
	return v29
}
func F_discardTransaction(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	F_resetClientMultiState(m, l0)
	mBase = m.M
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v4 & int32(-4137)
		F_unwatchAllKeys(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	}
}
func F_disconnectReplicas(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_disconnectReplicas[0]))
	v10 = v5 + int32(8)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v11
	goto L1
L1:
	;
	v16 = v5 + int32(8)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v5 + int32(16)
	return
L3:
	;
	if v18 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18+base.B2i32(v21 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v27
	goto L4
L6:
	;
	v32 = v18
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v34 & int32(-33554433)
	v38 = F_freeClient(m, v33)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	return
L10:
	;
	v41 = v5 + int32(8)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v43 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v43 != 0 {
		v32 = v43
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v43+base.B2i32(v46 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v52
	goto L12
L14:
	;
	goto L8
}
func F_dispose_chunk(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var __phi59 int32
	_ = __phi59
	var v60 int32
	_ = v60
	var __phi60 int32
	_ = __phi60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var __phi221 int32
	_ = __phi221
	var v222 int32
	_ = v222
	var __phi222 int32
	_ = __phi222
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v304 int32
	_ = v304
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v414 int32
	_ = v414
	v10 = l0 + l1
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11&int32(1) != 0 {
		v132 = l0
		v133 = l1
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v141&int32(2) != 0 {
		goto L40
	} else {
		goto L41
	}
L3:
	;
	if v11&int32(2) == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = v18 + l1
	v20 = l0 - v18
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[0]))
	if v20 == v22 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	if v38 == int32(0) {
		v132 = v20
		v133 = v19
		goto L2
	} else {
		goto L24
	}
L6:
	;
	v92 = int32(0)
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v27
	v132 = v20
	v133 = v19
	goto L2
L8:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v73 = int32(3)
	if v72&v73 != v73 {
		v132 = v20
		v133 = v19
		goto L2
	} else {
		goto L23
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if base.Ui32(int32(255)) < base.Ui32(v18) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v24 == v20 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v24 != v27 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v29 = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[1])) = v31 & base.I32_rotl(int32(-2), int32(base.Ui32(v18)>>(uint(int32(3))%32)))
	v132 = v20
	v133 = v19
	goto L2
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v43 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v40
	v92 = v24
	goto L5
L15:
	;
	__phi59 = v53
	__phi60 = v54
	v59 = __phi59
	v60 = __phi60
	goto L19
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v48 == int32(0) {
		goto L6
	} else {
		goto L18
	}
L17:
	;
	v53 = v43
	v54 = v20 + int32(20)
	goto L15
L18:
	;
	v53 = v48
	v54 = v20 + int32(16)
	goto L15
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	if v66 != 0 {
		__phi59 = v66
		__phi60 = v59 + int32(20)
		v59 = __phi59
		v60 = __phi60
		goto L19
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(0)
	v92 = v59
	goto L5
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	if v69 != 0 {
		__phi59 = v69
		__phi60 = v59 + int32(16)
		v59 = __phi59
		v60 = __phi60
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[2])) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v72 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v19 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v19
	return
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v102 = v100 << (uint(int32(2)) % 32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_dispose_chunk[3])))
	if v20 != v105 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = v38
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v115 != v20 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_dispose_chunk[3]))) = v92
	if v92 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v108 = int32(0)
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[4])) = v110 & base.I32_rotl(int32(-2), v100)
	v132 = v20
	v133 = v19
	goto L2
L29:
	;
	if v92 == int32(0) {
		v132 = v20
		v133 = v19
		goto L2
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v92
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v92
	goto L29
L32:
	;
	goto L25
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v127 == int32(0) {
		v132 = v20
		v133 = v19
		goto L2
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v122)+24)) = v92
	goto L33
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v127)+24)) = v92
	v132 = v20
	v133 = v19
	goto L2
L36:
	;
	if base.Ui32(int32(255)) < base.Ui32(v304) {
		goto L74
	} else {
		goto L75
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+4)) = v183 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v132+v183))) = v183
	if v132 != v167 {
		v304 = v183
		goto L36
	} else {
		goto L73
	}
L38:
	;
	if v200 == int32(0) {
		goto L37
	} else {
		goto L61
	}
L39:
	;
	v246 = int32(0)
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v141 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v132)+4)) = v133 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v132+v133))) = v133
	v304 = v133
	goto L36
L41:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[5]))
	if v10 != v145 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[0]))
	if v10 != v167 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v147 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[5])) = v132
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[6]))
	v152 = v151 + v133
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[6])) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v132)+4)) = v152 | int32(1)
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[0]))
	if v132 != v158 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v160 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[2])) = v160
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[0])) = v160
	return
L45:
	;
	v183 = v141&int32(-8) + v133
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if base.Ui32(int32(255)) < base.Ui32(v141) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v169 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[0])) = v132
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[2]))
	v174 = v173 + v133
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[2])) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v132)+4)) = v174 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v132+v174))) = v174
	return
L47:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v184 == v10 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v184 != v187 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187)+12)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v184)+8)) = v187
	goto L37
L50:
	;
	v189 = int32(0)
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[1])) = v191 & base.I32_rotl(int32(-2), int32(base.Ui32(v141)>>(uint(int32(3))%32)))
	goto L37
L51:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v205 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+12)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v184)+8)) = v202
	v246 = v184
	goto L38
L53:
	;
	__phi221 = v215
	__phi222 = v216
	v221 = __phi221
	v222 = __phi222
	goto L57
L54:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v210 == int32(0) {
		goto L39
	} else {
		goto L56
	}
L55:
	;
	v215 = v205
	v216 = v10 + int32(20)
	goto L53
L56:
	;
	v215 = v210
	v216 = v10 + int32(16)
	goto L53
L57:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
	if v228 != 0 {
		__phi221 = v228
		__phi222 = v221 + int32(20)
		v221 = __phi221
		v222 = __phi222
		goto L57
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = int32(0)
	v246 = v221
	goto L38
L59:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
	if v231 != 0 {
		__phi221 = v231
		__phi222 = v221 + int32(16)
		v221 = __phi221
		v222 = __phi222
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v256 = v254 << (uint(int32(2)) % 32)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v256)+uint32(_c_F_dispose_chunk[3])))
	if v10 != v259 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+24)) = v200
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v276 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L63:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v200)+16))
	if v269 != v10 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+uint32(_c_F_dispose_chunk[3]))) = v246
	if v246 != 0 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v262 = int32(0)
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[4])) = v264 & base.I32_rotl(int32(-2), v254)
	goto L37
L66:
	;
	if v246 == int32(0) {
		goto L37
	} else {
		goto L69
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+20)) = v246
	goto L66
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = v246
	goto L66
L69:
	;
	goto L62
L70:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v281 == int32(0) {
		goto L37
	} else {
		goto L72
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+16)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v276)+24)) = v246
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+20)) = v281
	*(*int32)(unsafe.Add(mBase, uint32(v281)+24)) = v246
	goto L37
L73:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[2])) = v183
	return
L74:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v304) {
		v350 = int32(31)
		goto L79
	} else {
		goto L80
	}
L75:
	;
	v315 = v304 & int32(-8)
	v317 = v315 + int32(9128464)
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[1]))
	v323 = int32(1) << (uint(int32(base.Ui32(v304)>>(uint(int32(3))%32))) % 32)
	if v319&v323 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315)+uint32(_c_F_dispose_chunk[7]))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v329)+12)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v132)+12)) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v329
	return
L77:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v315)+uint32(_c_F_dispose_chunk[7])))
	v329 = v328
	goto L76
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[1])) = v319 | v323
	v329 = v317
	goto L76
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+28)) = v350
	*(*int64)(unsafe.Add(mBase, uint32(v132)+16)) = int64(0)
	v355 = v350 << (uint(int32(2)) % 32)
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[4]))
	v361 = int32(1) << (uint(v350) % 32)
	if v359&v361 != 0 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v340 = base.I32_clz(int32(base.Ui32(v304) >> (uint(int32(8)) % 32)))
	v343 = int32(1)
	v350 = int32(base.Ui32(v304)>>(uint(int32(38)-v340)%32))&v343 - v340<<(uint(v343)%32) + int32(62)
	goto L79
L81:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v383)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v414)+12)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v383)+8)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v132)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v132)+12)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v414
	goto L1
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+12)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v132
	return
L83:
	;
	if v350 == int32(31) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[4])) = v359 | v361
	*(*int32)(unsafe.Add(mBase, uint32(v355)+uint32(_c_F_dispose_chunk[3]))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v132)+24)) = v355 + int32(9128728)
	goto L82
L85:
	;
	v375 = int32(0)
	goto L87
L86:
	;
	v375 = int32(25) - int32(base.Ui32(v350)>>(uint(int32(1))%32))
	goto L87
L87:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v355)+uint32(_c_F_dispose_chunk[3])))
	v381 = v304 << (uint(v375) % 32)
	v383 = v377
	goto L88
L88:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v387&int32(-8) == v304 {
		goto L81
	} else {
		goto L90
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v397+int32(16)))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v132)+24)) = v383
	goto L82
L90:
	;
	v397 = v383 + int32(base.Ui32(v381)>>(uint(int32(29))%32))&int32(4)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)+16))
	if v398 != 0 {
		v381 = v381 << (uint(int32(1)) % 32)
		v383 = v398
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
}
func F_do_getc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v3 < int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v13 == v14 {
			v21 = F___uflow(m, l0)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return v21
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v13 + int32(1)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			return v19
		}
	} else {
		if v3 == int32(0) {
			v26 = F_locking_getc(m, l0)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				return v26
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_do_getc[0]))
			if v3&int32(1073741823) != v11 {
				v26 = F_locking_getc(m, l0)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					return v26
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v13 == v14 {
					v21 = F___uflow(m, l0)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						return v21
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v13 + int32(1)
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
					return v19
				}
			}
		}
	}
}
func F_do_glob(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var __phi44 int32
	_ = __phi44
	var v45 int32
	_ = v45
	var __phi45 int32
	_ = __phi45
	var v46 int32
	_ = v46
	var __phi46 int32
	_ = __phi46
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
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
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
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
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v667 int32
	_ = v667
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v759 int32
	_ = v759
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v808 int32
	_ = v808
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v858 int32
	_ = v858
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v912 int32
	_ = v912
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v976 int32
	_ = v976
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	v8 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(96)
	m.G0 = v24
	v28 = int32(2)
	v29 = l4 & v28
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v34 = l2
	goto L3
L2:
	;
	v34 = v29<<(uint(v28)%32) ^ int32(8)
	goto L3
L3:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v36 = v8
	goto L6
L5:
	;
	v36 = v34
	goto L6
L6:
	;
	v40 = l1 + int32(1)
	if base.Ui32(int32(4095)) < base.Ui32(v40) {
		v78 = l1
		v80 = l3
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v34 != int32(4) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	__phi44 = l1
	__phi45 = v40
	__phi46 = l3
	v44 = __phi44
	v45 = __phi45
	v46 = __phi46
	goto L9
L9:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v64 != int32(47) {
		v78 = v44
		v80 = v46
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v78 = int32(4095)
	v80 = v71
	goto L7
L11:
	;
	v68 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v44))) = uint8(v68)
	v70 = int32(1)
	v71 = v46 + v70
	v73 = v45 + v70
	if v73 != int32(4096) {
		__phi44 = v45
		__phi45 = v73
		__phi46 = v71
		v44 = __phi44
		v45 = __phi45
		v46 = __phi46
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v98 = v36
	goto L15
L14:
	;
	v98 = v34
	goto L15
L15:
	;
	v100 = l4 & int32(64)
	v101 = int32(0)
	v105 = v78
	v107 = v80
	v112 = v8
	v114 = v101
	v115 = v101
	v117 = v98
	v119 = v101
	goto L21
L16:
	;
	m.G0 = v24 + int32(96)
	return v1023
L17:
	;
	v987 = F_emscripten_builtin_malloc(m, v209+int32(6))
	mBase = m.M
	if v987 != 0 {
		goto L254
	} else {
		goto L255
	}
L18:
	;
	v1023 = int32(0)
	goto L16
L19:
	;
	v212 = l0 + v209
	v213 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v212))) = uint8(v213)
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v215 != 0 {
		goto L48
	} else {
		goto L49
	}
L20:
	;
	v209 = v105
	v211 = v107
	goto L19
L21:
	;
	v125 = v107 + v114
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if v126 == int32(42) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	if v126 == int32(63) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	if v119 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	switch v126 + int32(-91) {
	case 0:
		goto L30
	case 1:
		goto L33
	default:
		goto L34
	}
L26:
	;
	if v126 == int32(93) {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v105 = v197
	v107 = v199
	v112 = v202
	v114 = v200 + int32(1)
	v115 = v201
	v117 = int32(0)
	v119 = v204
	goto L21
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v186+v190))) = uint8(v187)
	v197 = v186
	v199 = v188
	v200 = v189
	v201 = v192
	v202 = v112
	v204 = v193
	goto L28
L30:
	;
	v178 = int32(1)
	v182 = v115 + v178
	if base.Ui32(int32(4096)) <= base.Ui32(v105+v182) {
		v197 = v105
		v199 = v107
		v200 = v114
		v201 = v115
		v202 = v178
		v204 = v178
		goto L28
	} else {
		goto L47
	}
L31:
	;
	v173 = v115 + int32(1)
	if base.Ui32(v105+v173) < base.Ui32(int32(4096)) {
		v186 = v105
		v187 = v170
		v188 = v107
		v189 = v171
		v190 = v115
		v192 = v173
		v193 = v119
		goto L29
	} else {
		goto L45
	}
L32:
	;
	if v152 != int32(47) {
		v170 = v152
		v171 = v153
		goto L31
	} else {
		goto L42
	}
L33:
	;
	if v100 != 0 {
		v170 = int32(92)
		v171 = v114
		goto L31
	} else {
		goto L37
	}
L34:
	;
	if v126 != 0 {
		v152 = v126
		v153 = v114
		goto L32
	} else {
		goto L35
	}
L35:
	;
	if v112 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	v209 = v105 + v115
	v211 = v125
	goto L19
L37:
	;
	if v119 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v147 = v114 + int32(1)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v147))))
	if v149 == int32(0) {
		goto L18
	} else {
		goto L41
	}
L39:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125+int32(1)))))
	if v143 == int32(93) {
		goto L20
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v152 = v149
	v153 = v147
	goto L32
L42:
	;
	v156 = int32(0)
	if v112 != 0 {
		v1023 = v156
		goto L16
	} else {
		goto L43
	}
L43:
	;
	v159 = v115 + v105 + int32(1)
	if base.Ui32(int32(4095)) < base.Ui32(v159) {
		v1023 = v156
		goto L16
	} else {
		goto L44
	}
L44:
	;
	v162 = v107 + v153
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v166 = int32(-1)
	v167 = int32(0)
	v186 = v159
	v187 = v165
	v188 = v162 + int32(1)
	v189 = v166
	v190 = v166
	v192 = v167
	v193 = v167
	goto L29
L45:
	;
	if v119 != 0 {
		v197 = v105
		v199 = v107
		v200 = v171
		v201 = v115
		v202 = int32(1)
		v204 = v119
		goto L28
	} else {
		goto L46
	}
L46:
	;
	goto L18
L47:
	;
	v186 = v105
	v187 = int32(91)
	v188 = v107
	v189 = v114
	v190 = v115
	v192 = v182
	v193 = v178
	goto L29
L48:
	;
	v249 = int32(47)
	v251 = F___strchrnul(m, v211, v249)
	mBase = m.M
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if v253 == v249 {
		goto L71
	} else {
		goto L72
	}
L49:
	;
	if v29 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v117 != 0 {
		v976 = v117
		goto L17
	} else {
		goto L60
	}
L51:
	;
	if v117 == int32(10) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v222 = F___fstatat(m, int32(-100), l0, v24, int32(0))
	mBase = m.M
	goto L55
L53:
	;
	if v117 != 0 {
		v976 = v117
		goto L17
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	if v222 != 0 {
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v225&int32(61440) == int32(16384) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v230 = int32(4)
	goto L59
L58:
	;
	v230 = int32(8)
	goto L59
L59:
	;
	v976 = v230
	goto L17
L60:
	;
	v233 = F___fstatat(m, int32(-100), l0, v24, int32(256))
	mBase = m.M
	goto L62
L61:
	;
	goto L64
L62:
	;
	if v233 != 0 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v976 = int32(0)
	goto L17
L64:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_do_glob[0]))
	if v236 == int32(44) {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	v240 = m.T0[l5].(func(*base.Module, int32, int32) int32)(m, l0, v236)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	return int32(0)
L67:
	;
	if v240|l4&int32(1) == int32(0) {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	v1023 = int32(2)
	goto L16
L69:
	;
	if v209 != 0 {
		goto L84
	} else {
		goto L85
	}
L70:
	;
	if v257 == int32(0) {
		v298 = v257
		v304 = v249
		goto L69
	} else {
		goto L74
	}
L71:
	;
	v257 = v251
	goto L73
L72:
	;
	v257 = int32(0)
	goto L73
L73:
	;
	goto L70
L74:
	;
	if v100 != 0 {
		v298 = v257
		v304 = v249
		goto L69
	} else {
		goto L75
	}
L75:
	;
	v262 = v257
	goto L77
L76:
	;
	v292 = (v257 - v262) & int32(1)
	if v292 != 0 {
		goto L81
	} else {
		goto L82
	}
L77:
	;
	if base.Ui32(v262) <= base.Ui32(v211) {
		goto L76
	} else {
		goto L79
	}
L78:
	;
	goto L76
L79:
	;
	v283 = v262 + int32(-1)
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	if v284 == int32(92) {
		v262 = v283
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v293 = int32(92)
	goto L83
L82:
	;
	v293 = int32(47)
	goto L83
L83:
	;
	v298 = v257 - v292
	v304 = v293
	goto L69
L84:
	;
	v317 = l0
	goto L86
L85:
	;
	v317 = int32(_a_F_do_glob_0)
	goto L86
L86:
	;
	v318 = F_opendir(m, v317)
	mBase = m.M
	v319 = int32(9116376)
	goto L87
L87:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_do_glob[0]))
	if v318 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v330 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_do_glob[0])) = v330
	v332 = F_readdir(m, v318)
	mBase = m.M
	if v332 == v330 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v321 = m.T0[l5].(func(*base.Module, int32, int32) int32)(m, l0, v320)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L66
	} else {
		goto L90
	}
L90:
	;
	v323 = int32(1)
	v1023 = base.B2i32(v321|l4&v323 != int32(0)) << (uint(v323) % 32)
	goto L16
L91:
	;
	v934 = *(*int32)(unsafe.Add(mBase, _c_F_do_glob[0]))
	if v298 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L92:
	;
	v336 = int32(4096) - v209
	v339 = int32(0)
	v347 = int32(base.Ui32(l4)>>(uint(int32(5))%32)) & int32(6)
	v349 = v347 ^ int32(4)
	v352 = v332
	goto L93
L93:
	;
	if v298 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L91
L95:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_glob[0])) = int32(0)
	v912 = F_readdir(m, v318)
	mBase = m.M
	if v912 != 0 {
		v352 = v912
		goto L93
	} else {
		goto L245
	}
L96:
	;
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+18)))
	v898 = F_do_glob(m, l0, v894+v209, v897, v893, l4, l5, l6)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L66
	} else {
		goto L242
	}
L97:
	;
	v704 = v352 + int32(19)
	if v704&int32(3) == int32(0) {
		v726 = v704
		goto L193
	} else {
		goto L194
	}
L98:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+18)))
	switch v373 {
	case 0, 4:
		goto L99
	case 1, 2, 3:
		goto L95
	default:
		goto L100
	}
L99:
	;
	v377 = v352 + int32(19)
	if v377&int32(3) == int32(0) {
		v399 = v377
		goto L104
	} else {
		goto L105
	}
L100:
	;
	if v373 != int32(10) {
		goto L95
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	if base.Ui32(v336) <= base.Ui32(v432) {
		goto L95
	} else {
		goto L118
	}
L103:
	;
	v432 = v424 - v377
	goto L102
L104:
	;
	v403 = v399
	goto L112
L105:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377))))
	if v385 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v388 = v377
	goto L108
L107:
	;
	v432 = v377 - v377
	goto L102
L108:
	;
	v392 = v388 + int32(1)
	if v392&int32(3) == int32(0) {
		v399 = v392
		goto L104
	} else {
		goto L110
	}
L110:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392))))
	if v397 != 0 {
		v388 = v392
		goto L108
	} else {
		goto L111
	}
L111:
	;
	v424 = v392
	goto L103
L112:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	v412 = int32(-2139062144)
	if (int32(16843008)-v409|v409)&v412 == v412 {
		v403 = v403 + int32(4)
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v418 = v403
	goto L115
L114:
	;
	goto L113
L115:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418))))
	if v422 != 0 {
		v418 = v418 + int32(1)
		goto L115
	} else {
		goto L117
	}
L116:
	;
	v424 = v418
	goto L103
L117:
	;
	goto L116
L118:
	;
	v434 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v434)
	v442 = m.G0
	v444 = v442 - int32(16)
	m.G0 = v444
	v447 = v349 & int32(8)
	if v349&int32(1) == v434 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	if v550 != 0 {
		goto L95
	} else {
		goto L149
	}
L120:
	;
	m.G0 = v444 + int32(16)
	goto L119
L121:
	;
	if v447 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L122:
	;
	v452 = v211
	v453 = v377
	goto L123
L123:
	;
	v466 = v453
	goto L126
L125:
	;
	v483 = v452
	goto L131
L126:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466))))
	if v470 == int32(47) {
		goto L125
	} else {
		goto L128
	}
L128:
	;
	if v470 == int32(0) {
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v466 = v466 + int32(1)
	goto L126
L130:
	;
	v496 = int32(*(*int8)(unsafe.Add(mBase, uint32(v466))))
	if v489 == v496 {
		goto L135
	} else {
		goto L136
	}
L131:
	;
	v489 = F_pat_next(m, v483, int32(-1), v444+int32(12), v349)
	mBase = m.M
	if v489 == int32(0) {
		goto L130
	} else {
		goto L133
	}
L133:
	;
	if v489 == int32(47) {
		goto L130
	} else {
		goto L134
	}
L134:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v444)+12))
	v483 = v483 + v494
	goto L131
L135:
	;
	v508 = F_fnmatch_internal(m, v452, v483-v452, v453, v466-v453, v349)
	mBase = m.M
	if v489 == int32(0) {
		v550 = v508
		goto L120
	} else {
		goto L138
	}
L136:
	;
	v498 = int32(0)
	if base.B2i32(v447 == v498)|base.B2i32(v496 == v498) == v498 {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v550 = int32(1)
	goto L120
L138:
	;
	if v508 != 0 {
		v550 = v508
		goto L120
	} else {
		goto L139
	}
L139:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v444)+12))
	v452 = v483 + v511
	v453 = v466 + int32(1)
	goto L123
L140:
	;
	v546 = int32(-1)
	v548 = F_fnmatch_internal(m, v211, v546, v377, v546, v349)
	mBase = m.M
	v550 = v548
	goto L120
L141:
	;
	v522 = v377
	goto L142
L142:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	if v526 == int32(47) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v522 = v522 + int32(1)
	goto L142
L145:
	;
	v533 = F_fnmatch_internal(m, v211, int32(-1), v377, v522-v377, v349)
	mBase = m.M
	if v533 != 0 {
		goto L144
	} else {
		goto L148
	}
L146:
	;
	if v526 == int32(0) {
		goto L140
	} else {
		goto L147
	}
L147:
	;
	goto L144
L148:
	;
	v550 = int32(0)
	goto L120
L149:
	;
	if base.B2i32(l4&int32(128) == v339)|base.B2i32(v298 == v339) != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v697 = v432 + int32(1)
	if v697 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L151:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377))))
	if v561 != int32(46) {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+20)))
	if v564 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v576 = m.G0
	v578 = v576 - int32(16)
	m.G0 = v578
	v581 = v347 & int32(8)
	if v347&int32(1) == int32(0) {
		goto L159
	} else {
		goto L160
	}
L154:
	;
	if v564 != int32(46) {
		goto L150
	} else {
		goto L155
	}
L155:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+21)))
	if v569 != 0 {
		goto L150
	} else {
		goto L156
	}
L156:
	;
	goto L153
L157:
	;
	if v684 != 0 {
		goto L95
	} else {
		goto L187
	}
L158:
	;
	m.G0 = v578 + int32(16)
	goto L157
L159:
	;
	if v581 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L160:
	;
	v586 = v211
	v587 = v377
	goto L161
L161:
	;
	v600 = v587
	goto L164
L163:
	;
	v617 = v586
	goto L169
L164:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600))))
	if v604 == int32(47) {
		goto L163
	} else {
		goto L166
	}
L166:
	;
	if v604 == int32(0) {
		goto L163
	} else {
		goto L167
	}
L167:
	;
	v600 = v600 + int32(1)
	goto L164
L168:
	;
	v630 = int32(*(*int8)(unsafe.Add(mBase, uint32(v600))))
	if v623 == v630 {
		goto L173
	} else {
		goto L174
	}
L169:
	;
	v623 = F_pat_next(m, v617, int32(-1), v578+int32(12), v347)
	mBase = m.M
	if v623 == int32(0) {
		goto L168
	} else {
		goto L171
	}
L171:
	;
	if v623 == int32(47) {
		goto L168
	} else {
		goto L172
	}
L172:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v578)+12))
	v617 = v617 + v628
	goto L169
L173:
	;
	v642 = F_fnmatch_internal(m, v586, v617-v586, v587, v600-v587, v347)
	mBase = m.M
	if v623 == int32(0) {
		v684 = v642
		goto L158
	} else {
		goto L176
	}
L174:
	;
	v632 = int32(0)
	if base.B2i32(v581 == v632)|base.B2i32(v630 == v632) == v632 {
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v684 = int32(1)
	goto L158
L176:
	;
	if v642 != 0 {
		v684 = v642
		goto L158
	} else {
		goto L177
	}
L177:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v578)+12))
	v586 = v617 + v645
	v587 = v600 + int32(1)
	goto L161
L178:
	;
	v680 = int32(-1)
	v682 = F_fnmatch_internal(m, v211, v680, v377, v680, v347)
	mBase = m.M
	v684 = v682
	goto L158
L179:
	;
	v656 = v377
	goto L180
L180:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656))))
	if v660 == int32(47) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v656 = v656 + int32(1)
	goto L180
L183:
	;
	v667 = F_fnmatch_internal(m, v211, int32(-1), v377, v656-v377, v347)
	mBase = m.M
	if v667 != 0 {
		goto L182
	} else {
		goto L186
	}
L184:
	;
	if v660 == int32(0) {
		goto L178
	} else {
		goto L185
	}
L185:
	;
	goto L182
L186:
	;
	v684 = int32(0)
	goto L158
L187:
	;
	goto L150
L188:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v304)
	v893 = v298
	v894 = v432
	goto L96
L189:
	;
	goto L188
L190:
	;
	v700 = F__emscripten_memcpy_bulkmem(m, v212, v377, v697)
	mBase = m.M
	goto L189
L191:
	;
	if base.Ui32(v336) <= base.Ui32(v759) {
		goto L95
	} else {
		goto L207
	}
L192:
	;
	v759 = v751 - v704
	goto L191
L193:
	;
	v730 = v726
	goto L201
L194:
	;
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704))))
	if v712 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v715 = v704
	goto L197
L196:
	;
	v759 = v704 - v704
	goto L191
L197:
	;
	v719 = v715 + int32(1)
	if v719&int32(3) == int32(0) {
		v726 = v719
		goto L193
	} else {
		goto L199
	}
L199:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v719))))
	if v724 != 0 {
		v715 = v719
		goto L197
	} else {
		goto L200
	}
L200:
	;
	v751 = v719
	goto L192
L201:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	v739 = int32(-2139062144)
	if (int32(16843008)-v736|v736)&v739 == v739 {
		v730 = v730 + int32(4)
		goto L201
	} else {
		goto L203
	}
L202:
	;
	v745 = v730
	goto L204
L203:
	;
	goto L202
L204:
	;
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v745))))
	if v749 != 0 {
		v745 = v745 + int32(1)
		goto L204
	} else {
		goto L206
	}
L205:
	;
	v751 = v745
	goto L192
L206:
	;
	goto L205
L207:
	;
	v767 = m.G0
	v769 = v767 - int32(16)
	m.G0 = v769
	v772 = v349 & int32(8)
	if v349&int32(1) == int32(0) {
		goto L210
	} else {
		goto L211
	}
L208:
	;
	if v875 != 0 {
		goto L95
	} else {
		goto L238
	}
L209:
	;
	m.G0 = v769 + int32(16)
	goto L208
L210:
	;
	if v772 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L211:
	;
	v777 = v211
	v778 = v704
	goto L212
L212:
	;
	v791 = v778
	goto L215
L214:
	;
	v808 = v777
	goto L220
L215:
	;
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v791))))
	if v795 == int32(47) {
		goto L214
	} else {
		goto L217
	}
L217:
	;
	if v795 == int32(0) {
		goto L214
	} else {
		goto L218
	}
L218:
	;
	v791 = v791 + int32(1)
	goto L215
L219:
	;
	v821 = int32(*(*int8)(unsafe.Add(mBase, uint32(v791))))
	if v814 == v821 {
		goto L224
	} else {
		goto L225
	}
L220:
	;
	v814 = F_pat_next(m, v808, int32(-1), v769+int32(12), v349)
	mBase = m.M
	if v814 == int32(0) {
		goto L219
	} else {
		goto L222
	}
L222:
	;
	if v814 == int32(47) {
		goto L219
	} else {
		goto L223
	}
L223:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v769)+12))
	v808 = v808 + v819
	goto L220
L224:
	;
	v833 = F_fnmatch_internal(m, v777, v808-v777, v778, v791-v778, v349)
	mBase = m.M
	if v814 == int32(0) {
		v875 = v833
		goto L209
	} else {
		goto L227
	}
L225:
	;
	v823 = int32(0)
	if base.B2i32(v772 == v823)|base.B2i32(v821 == v823) == v823 {
		goto L224
	} else {
		goto L226
	}
L226:
	;
	v875 = int32(1)
	goto L209
L227:
	;
	if v833 != 0 {
		v875 = v833
		goto L209
	} else {
		goto L228
	}
L228:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v769)+12))
	v777 = v808 + v836
	v778 = v791 + int32(1)
	goto L212
L229:
	;
	v871 = int32(-1)
	v873 = F_fnmatch_internal(m, v211, v871, v704, v871, v349)
	mBase = m.M
	v875 = v873
	goto L209
L230:
	;
	v847 = v704
	goto L231
L231:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847))))
	if v851 == int32(47) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	v847 = v847 + int32(1)
	goto L231
L234:
	;
	v858 = F_fnmatch_internal(m, v211, int32(-1), v704, v847-v704, v349)
	mBase = m.M
	if v858 != 0 {
		goto L233
	} else {
		goto L237
	}
L235:
	;
	if v851 == int32(0) {
		goto L229
	} else {
		goto L236
	}
L236:
	;
	goto L233
L237:
	;
	v875 = int32(0)
	goto L209
L238:
	;
	v887 = v759 + int32(1)
	if v887 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	v893 = int32(_a_F_do_glob_1)
	v894 = v759
	goto L96
L240:
	;
	goto L239
L241:
	;
	v890 = F__emscripten_memcpy_bulkmem(m, v212, v704, v887)
	mBase = m.M
	goto L240
L242:
	;
	if v898 == int32(0) {
		goto L95
	} else {
		goto L243
	}
L243:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	v904 = F_close(m, v903)
	mBase = m.M
	F_emscripten_builtin_free(m, v318)
	mBase = m.M
	goto L244
L244:
	;
	v1023 = v898
	goto L16
L245:
	;
	goto L94
L246:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	v940 = F_close(m, v939)
	mBase = m.M
	F_emscripten_builtin_free(m, v318)
	mBase = m.M
	goto L248
L247:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v304)
	goto L246
L248:
	;
	if v934 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_glob[0])) = v320
	goto L18
L250:
	;
	v944 = *(*int32)(unsafe.Add(mBase, _c_F_do_glob[0]))
	v945 = m.T0[l5].(func(*base.Module, int32, int32) int32)(m, l0, v944)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L66
	} else {
		goto L251
	}
L251:
	;
	if v945|l4&int32(1) == int32(0) {
		goto L249
	} else {
		goto L252
	}
L252:
	;
	v1023 = int32(2)
	goto L16
L253:
	;
	v1023 = base.B2i32(v1018 != int32(0))
	goto L16
L254:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(v989))) = v987
	v991 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v987))) = v991
	v994 = v987 + int32(4)
	v996 = v209 + int32(1)
	v997 = F___memcpy(m, v994, l0, v996)
	mBase = m.M
	if v209 == v991 {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	v1018 = int32(-1)
	goto L253
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v987
	v1018 = int32(0)
	goto L253
L257:
	;
	if base.B2i32(v29 != int32(0))&base.B2i32(v976 == int32(4)) == int32(0) {
		goto L256
	} else {
		goto L258
	}
L258:
	;
	v1005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v209+int32(-1)))))
	if v1005 == int32(47) {
		goto L256
	} else {
		goto L259
	}
L259:
	;
	v1009 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v994+v209))) = uint8(v1009)
	v1012 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v994+v996))) = uint8(v1012)
	goto L256
}
func F_doesCommandHaveKeys(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v47 int64
	_ = v47
	var v54 int64
	_ = v54
	var v61 int64
	_ = v61
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int64
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int64
	_ = v104
	var v112 int32
	_ = v112
	v8 = int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v9 != 0 {
		v112 = v8
		return v112
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+58)))
		if v10&int32(32) != 0 {
			v112 = v8
			return v112
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			if int32(1) <= v13 {
				v19 = v13 & int32(3)
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				if base.Ui32(int32(4)) <= base.Ui32(v13) {
					v28 = int32(0)
					v31 = v28
					v33 = v28
					v36 = int64(0)
					for {
						v38 = int32(48)
						v41 = *(*int64)(unsafe.Add(mBase, uint32(v20+v31*v38)+8))
						v47 = *(*int64)(unsafe.Add(mBase, uint32(v20+(v31|int32(1))*v38)+8))
						v54 = *(*int64)(unsafe.Add(mBase, uint32(v20+(v31|int32(2))*v38)+8))
						v61 = *(*int64)(unsafe.Add(mBase, uint32(v20+(v31|int32(3))*v38)+8))
						v65 = v36 | (v41&v47&v54&v61 ^ int64(-1))
						v66 = int32(4)
						v67 = v31 + v66
						v69 = v33 + v66
						if v69 != v13&int32(2147483644) {
							v31 = v67
							v33 = v69
							v36 = v65
							continue
						} else {
							break
						}
						break
					}
					v71 = v67
					v76 = v65
				} else {
					v71 = int32(0)
					v76 = int64(0)
				}
				if v19 == int32(0) {
					v104 = v76
				} else {
					v80 = v71
					v84 = int32(0)
					v85 = v76
					for {
						v90 = *(*int64)(unsafe.Add(mBase, uint32(v20+v80*int32(48))+8))
						v93 = v85 | (v90 ^ int64(-1))
						v94 = int32(1)
						v97 = v84 + v94
						if v97 != v19 {
							v80 = v80 + v94
							v84 = v97
							v85 = v93
							continue
						} else {
							break
						}
						break
					}
					v104 = v93
				}
				v112 = int32(base.Ui32(base.I32_wrap_i64(v104))>>(uint(int32(8))%32)) & int32(1)
				return v112
			} else {
				return int32(0)
			}
		}
	}
}
func F_dualChannelFullSyncWithPrimary(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v230 int32
	_ = v230
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int64
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int64
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int64
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
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
	var v561 int32
	_ = v561
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int64
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int64
	_ = v645
	var v648 int64
	_ = v648
	var v651 int64
	_ = v651
	var v655 int32
	_ = v655
	var v660 int64
	_ = v660
	var v667 int64
	_ = v667
	var v674 int64
	_ = v674
	var v681 int64
	_ = v681
	var v688 int32
	_ = v688
	var v695 int32
	_ = v695
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int64
	_ = v714
	var v717 int64
	_ = v717
	var v732 int64
	_ = v732
	var v735 int32
	_ = v735
	var v742 int64
	_ = v742
	var v746 int32
	_ = v746
	var v750 int64
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int64
	_ = v760
	var v762 int64
	_ = v762
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int64
	_ = v768
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v825 int32
	_ = v825
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v888 int32
	_ = v888
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	v9 = m.G0
	v11 = v9 - int32(624)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[0]))
	if l0 != v14 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	m.G0 = v11 + int32(624)
	return
L2:
	;
	v896 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[1]))
	if v896 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L3:
	;
	F_sdsfree(m, v882)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L14
	} else {
		goto L201
	}
L4:
	;
	if int32(3) < v866 {
		v882 = v867
		goto L3
	} else {
		goto L199
	}
L5:
	;
	v865 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	v866 = v865
	v867 = v859
	goto L4
L6:
	;
	F__serverAssert(m, int32(_a_F_dualChannelFullSyncWithPrimary_0), int32(_a_F_dualChannelFullSyncWithPrimary_1), int32(3181))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L14
	} else {
		goto L198
	}
L7:
	;
	F__serverAssert(m, int32(_a_F_dualChannelFullSyncWithPrimary_2), int32(_a_F_dualChannelFullSyncWithPrimary_1), int32(3166))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L14
	} else {
		goto L197
	}
L8:
	;
	F__serverAssert(m, int32(_a_F_dualChannelFullSyncWithPrimary_3), int32(_a_F_dualChannelFullSyncWithPrimary_1), int32(3195))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L14
	} else {
		goto L196
	}
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[3]))
	if v17 == int32(0) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 == int32(3) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[4]))
	switch v39 + int32(-1) {
	case 0:
		goto L22
	case 1:
		goto L21
	case 2:
		goto L20
	case 3:
		goto L19
	default:
		goto L17
	}
L12:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v24 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
	v29 = m.T0[v28].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+336)) = v29
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_4), v11+int32(336))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	goto L2
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v39
	F__serverPanic_1(m, int32(_a_F_dualChannelFullSyncWithPrimary_1), int32(3231), int32(_a_F_dualChannelFullSyncWithPrimary_5), v11)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L14
	} else {
		goto L195
	}
L18:
	;
	F_sdsfree(m, v825)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L14
	} else {
		goto L194
	}
L19:
	;
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[5]))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)+100))
	v575 = m.T0[v574].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v11+int32(368), int32(256), base.I64_extend_i32_s(v569*int32(1000)))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L14
	} else {
		goto L145
	}
L20:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[5]))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)+100))
	v415 = m.T0[v414].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v11+int32(368), int32(256), base.I64_extend_i32_s(v409*int32(1000)))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L14
	} else {
		goto L107
	}
L21:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[6]))
	if v335 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L22:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(0) < v43 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[6]))
	if v52 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v46 = int32(0)
	F__serverLog(m, v46, int32(_a_F_dualChannelFullSyncWithPrimary_6), v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[7]))
	if v195 != 0 {
		v204 = v195
		goto L57
	} else {
		goto L58
	}
L27:
	;
	v55 = int32(0)
	v59 = v11 + int32(376)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v61
	v64 = *(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+368)) = v64
	v69 = v11 + int32(352)
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v71
	v74 = *(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[11]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+344)) = v74
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[12]))
	if v77 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v52
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-1)))))
	switch v151 & int32(7) {
	case 0:
		goto L52
	case 1:
		goto L51
	case 2:
		goto L50
	case 3:
		goto L49
	case 4:
		goto L48
	default:
		v168 = v55
		goto L47
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+372)) = v77
	if v77&int32(3) == int32(0) {
		v109 = v77
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v80 = int32(4)
	v145 = v11 + int32(368) | v80
	v146 = v11 + int32(344) | v80
	v147 = int32(2)
	goto L28
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+348)) = v142
	v145 = v59
	v146 = v69
	v147 = int32(3)
	goto L28
L32:
	;
	v142 = v134 - v77
	goto L31
L33:
	;
	v113 = v109
	goto L41
L34:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v95 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v98 = v77
	goto L37
L36:
	;
	v142 = v77 - v77
	goto L31
L37:
	;
	v102 = v98 + int32(1)
	if v102&int32(3) == int32(0) {
		v109 = v102
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v107 != 0 {
		v98 = v102
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v134 = v102
	goto L32
L41:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v122 = int32(-2139062144)
	if (int32(16843008)-v119|v119)&v122 == v122 {
		v113 = v113 + int32(4)
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v128 = v113
	goto L44
L43:
	;
	goto L42
L44:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v132 != 0 {
		v128 = v128 + int32(1)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v134 = v128
	goto L32
L46:
	;
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v168
	v174 = F_sendCommandArgv(m, l0, v147, v11+int32(368), v11+int32(344))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L14
	} else {
		goto L53
	}
L48:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-17))))
	v168 = v167
	goto L47
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-9))))
	v168 = v164
	goto L47
L50:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(-5)))))
	v168 = v161
	goto L47
L51:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-3)))))
	v168 = v158
	goto L47
L52:
	;
	v168 = int32(base.Ui32(v151) >> (uint(int32(3)) % 32))
	goto L47
L53:
	;
	if v174 == int32(0) {
		goto L26
	} else {
		goto L54
	}
L54:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v179 {
		v882 = v174
		goto L3
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v174
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_7), v11+int32(160))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L14
	} else {
		goto L56
	}
L56:
	;
	v859 = v174
	goto L5
L57:
	;
	v207 = F_sdsfromlonglong(m, base.I64_extend_i32_s(v204))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L14
	} else {
		goto L65
	}
L58:
	;
	v196 = int32(_a_F_dualChannelFullSyncWithPrimary_8)
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[13]))
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[14]))
	if v197 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v200 = v197
	goto L61
L60:
	;
	v200 = v199
	goto L61
L61:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[15]))
	if v202 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v203 = v200
	goto L64
L63:
	;
	v203 = v199
	goto L64
L64:
	;
	v204 = v203
	goto L57
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(156)))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(152)))) = int32(_a_F_dualChannelFullSyncWithPrimary_9)
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(148)))) = int32(_a_F_dualChannelFullSyncWithPrimary_10)
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(144)))) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(140)))) = int32(_a_F_dualChannelFullSyncWithPrimary_11)
	v230 = int32(_a_F_dualChannelFullSyncWithPrimary_12)
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(136)))) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(132)))) = int32(_a_F_dualChannelFullSyncWithPrimary_13)
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(128)))) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v11)+124)) = int32(_a_F_dualChannelFullSyncWithPrimary_14)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+120)) = int32(_a_F_dualChannelFullSyncWithPrimary_15)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+116)) = int32(_a_F_dualChannelFullSyncWithPrimary_16)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = int32(_a_F_dualChannelFullSyncWithPrimary_17)
	v250 = F_sendCommand(m, l0, v11+int32(112))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L14
	} else {
		goto L66
	}
L66:
	;
	F_sdsfree(m, v207)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L14
	} else {
		goto L67
	}
L67:
	;
	if v250 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[16]))
	if v268 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v257 {
		v882 = v250
		goto L3
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v250
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_7), v11+int32(96))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L14
	} else {
		goto L71
	}
L71:
	;
	v859 = v250
	goto L5
L72:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+84))
	v299 = m.T0[v298].(func(*base.Module, int32, int32) int32)(m, l0, int32(975))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L14
	} else {
		goto L79
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+88)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = int32(_a_F_dualChannelFullSyncWithPrimary_18)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = int32(_a_F_dualChannelFullSyncWithPrimary_17)
	v280 = F_sendCommand(m, l0, v11+int32(80))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L14
	} else {
		goto L74
	}
L74:
	;
	if v280 == int32(0) {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v285 {
		v882 = v280
		goto L3
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v280
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_7), v11+int32(64))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L14
	} else {
		goto L77
	}
L77:
	;
	v859 = v280
	goto L5
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[4])) = int32(2)
	v825 = int32(0)
	goto L18
L79:
	;
	if v299 != int32(-1) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v304 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	goto L82
L82:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[17]))
	v309 = F___strerror_l(m, v308, v308)
	mBase = m.M
	goto L83
L83:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v310
	v318 = F_snprintf(m, v11+int32(368), int32(31), int32(_a_F_dualChannelFullSyncWithPrimary_19), v11+int32(32))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L14
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v11 + int32(368)
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_20), v11+int32(16))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L14
	} else {
		goto L85
	}
L85:
	;
	goto L2
L86:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[4])) = int32(3)
	goto L20
L87:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[5]))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)+100))
	v348 = m.T0[v347].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v11+int32(368), int32(256), base.I64_extend_i32_s(v342*int32(1000)))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L14
	} else {
		goto L91
	}
L88:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	if v385 != int32(45) {
		goto L100
	} else {
		goto L101
	}
L89:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v377 {
		goto L2
	} else {
		goto L98
	}
L90:
	;
	v367 = int32(_a_F_dualChannelFullSyncWithPrimary_8)
	v369 = *(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[18]))
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[19])) = v369
	v373 = F_sdsnew(m, v11+int32(368))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L14
	} else {
		goto L96
	}
L91:
	;
	if v348 != int32(-1) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v353 {
		goto L89
	} else {
		goto L93
	}
L93:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+88))
	v358 = m.T0[v357].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L14
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+256)) = v358
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_21), v11+int32(256))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L14
	} else {
		goto L95
	}
L95:
	;
	goto L89
L96:
	;
	if v373 != 0 {
		goto L88
	} else {
		goto L97
	}
L97:
	;
	goto L89
L98:
	;
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_22), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L14
	} else {
		goto L99
	}
L99:
	;
	goto L2
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[4])) = int32(3)
	v825 = v373
	goto L18
L101:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v389 {
		v882 = v373
		goto L3
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+272)) = v373
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_23), v11+int32(272))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L14
	} else {
		goto L103
	}
L103:
	;
	v859 = v373
	goto L5
L104:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
	if v452 != int32(45) {
		goto L116
	} else {
		goto L117
	}
L105:
	;
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v444 {
		goto L2
	} else {
		goto L114
	}
L106:
	;
	v434 = int32(_a_F_dualChannelFullSyncWithPrimary_8)
	v436 = *(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[18]))
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[19])) = v436
	v440 = F_sdsnew(m, v11+int32(368))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L14
	} else {
		goto L112
	}
L107:
	;
	if v415 != int32(-1) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v420 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v420 {
		goto L105
	} else {
		goto L109
	}
L109:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)+88))
	v425 = m.T0[v424].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L14
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+176)) = v425
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_21), v11+int32(176))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L14
	} else {
		goto L111
	}
L111:
	;
	goto L105
L112:
	;
	if v440 != 0 {
		goto L104
	} else {
		goto L113
	}
L113:
	;
	goto L105
L114:
	;
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_24), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L14
	} else {
		goto L115
	}
L115:
	;
	goto L2
L116:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[16]))
	if v467 == int32(0) {
		v533 = v440
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(2) < v456 {
		v866 = v456
		v867 = v440
		goto L4
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v440
	F__serverLog(m, int32(2), int32(_a_F_dualChannelFullSyncWithPrimary_25), v11+int32(192))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L14
	} else {
		goto L119
	}
L119:
	;
	v859 = v440
	goto L5
L120:
	;
	v537 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[5]))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+92))
	v543 = m.T0[v542].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, int32(_a_F_dualChannelFullSyncWithPrimary_26), int32(6), base.I64_extend_i32_s(v537*int32(1000)))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L14
	} else {
		goto L139
	}
L121:
	;
	F_sdsfree(m, v440)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L14
	} else {
		goto L122
	}
L122:
	;
	v476 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[5]))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)+100))
	v482 = m.T0[v481].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v11+int32(368), int32(256), base.I64_extend_i32_s(v476*int32(1000)))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L14
	} else {
		goto L126
	}
L123:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507))))
	if v519 != int32(45) {
		v533 = v507
		goto L120
	} else {
		goto L135
	}
L124:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v511 {
		goto L2
	} else {
		goto L133
	}
L125:
	;
	v501 = int32(_a_F_dualChannelFullSyncWithPrimary_8)
	v503 = *(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[18]))
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[19])) = v503
	v507 = F_sdsnew(m, v11+int32(368))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L14
	} else {
		goto L131
	}
L126:
	;
	if v482 != int32(-1) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v487 {
		goto L124
	} else {
		goto L128
	}
L128:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)+88))
	v492 = m.T0[v491].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L14
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+224)) = v492
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_21), v11+int32(224))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L14
	} else {
		goto L130
	}
L130:
	;
	goto L124
L131:
	;
	if v507 != 0 {
		goto L123
	} else {
		goto L132
	}
L132:
	;
	goto L124
L133:
	;
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_27), int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L14
	} else {
		goto L134
	}
L134:
	;
	goto L2
L135:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v523 {
		v882 = v507
		goto L3
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+240)) = v507
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_28), v11+int32(240))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L14
	} else {
		goto L137
	}
L137:
	;
	v859 = v507
	goto L5
L138:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[4])) = int32(4)
	v825 = v533
	goto L18
L139:
	;
	if v543 != int32(-1) {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v548 {
		v882 = v533
		goto L3
	} else {
		goto L141
	}
L141:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+88))
	v553 = m.T0[v552].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L14
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v553
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_29), v11+int32(208))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L14
	} else {
		goto L143
	}
L143:
	;
	v859 = v533
	goto L5
L144:
	;
	v594 = int32(_a_F_dualChannelFullSyncWithPrimary_8)
	v596 = *(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[18]))
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[19])) = v596
	v600 = F_sdsnew(m, v11+int32(368))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L14
	} else {
		goto L150
	}
L145:
	;
	if v575 != int32(-1) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v580 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v580 {
		goto L2
	} else {
		goto L147
	}
L147:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v583)+88))
	v585 = m.T0[v584].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L14
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+288)) = v585
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_21), v11+int32(288))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L14
	} else {
		goto L149
	}
L149:
	;
	goto L2
L150:
	;
	if v600 == int32(0) {
		goto L2
	} else {
		goto L151
	}
L151:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600))))
	if v604 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+320)) = v11 + int32(360)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+324)) = v11 + int32(368)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+328)) = v11 + int32(356)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+332)) = v11 + int32(344)
	v629 = F_sscanf(m, v600, int32(_a_F_dualChannelFullSyncWithPrimary_30), v11+int32(320))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L14
	} else {
		goto L157
	}
L153:
	;
	v606 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(0) < v606 {
		v825 = v600
		goto L18
	} else {
		goto L154
	}
L154:
	;
	v609 = int32(0)
	F__serverLog(m, v609, int32(_a_F_dualChannelFullSyncWithPrimary_31), v609)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L14
	} else {
		goto L155
	}
L155:
	;
	v825 = v600
	goto L18
L156:
	;
	v644 = int32(_a_F_dualChannelFullSyncWithPrimary_8)
	v645 = *(*int64)(unsafe.Add(mBase, uint32(v11)+344))
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[20])) = v645
	v648 = *(*int64)(unsafe.Add(mBase, uint32(v11)+360))
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[21])) = v648
	v651 = *(*int64)(unsafe.Add(mBase, uint32(v11)+368))
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[22])) = v651
	v655 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[23])) = v655
	v660 = *(*int64)(unsafe.Add(mBase, uint32(v11)+376))
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[24])) = v660
	v667 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(384))))
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[25])) = v667
	v674 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(392))))
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[26])) = v674
	v681 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(400))))
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[27])) = v681
	v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(408)))))
	*(*uint8)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[28])) = uint8(v688)
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[29])) = v648
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[30])) = v648
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v11)+356))
	*(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[31])) = v695
	*(*int32)(unsafe.Add(mBase, uint32(v655)+4)) = int32(3)
	*(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[3])) = int32(4)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v655)))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v703)+84))
	v705 = m.T0[v704].(func(*base.Module, int32, int32) int32)(m, v655, int32(978))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L14
	} else {
		goto L161
	}
L157:
	;
	if v629 == int32(4) {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v634 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(3) < v634 {
		v859 = v600
		goto L5
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+304)) = v600
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_32), v11+int32(304))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L14
	} else {
		goto L160
	}
L160:
	;
	v859 = v600
	goto L5
L161:
	;
	if v705 == int32(-1) {
		goto L7
	} else {
		goto L162
	}
L162:
	;
	v710 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[1]))
	F_dualChannelSetupMainConnForPsync(m, v710)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L14
	} else {
		goto L163
	}
L163:
	;
	v713 = int32(_a_F_dualChannelFullSyncWithPrimary_8)
	v714 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[32])) = v714
	v717 = int64(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[33])) = v717
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[34])) = v714
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[35])) = v714
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[36])) = v717
	v732 = *(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[18]))
	*(*int64)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[19])) = v732
	v735 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[37]))
	if v735&int32(-2) == int32(2) {
		goto L167
	} else {
		goto L168
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[4])) = int32(5)
	v825 = v600
	goto L18
L165:
	;
	v812 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[0]))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v812)))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v814)+84))
	v816 = m.T0[v815].(func(*base.Module, int32, int32) int32)(m, v812, int32(977))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L14
	} else {
		goto L192
	}
L166:
	;
	v797 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[0]))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v797)))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v799)+84))
	v801 = m.T0[v800].(func(*base.Module, int32, int32) int32)(m, v797, int32(979))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L14
	} else {
		goto L189
	}
L167:
	;
	v772 = F_moduleAllDatatypesHandleErrors(m)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L14
	} else {
		goto L181
	}
L168:
	;
	if v735 != int32(1) {
		goto L166
	} else {
		goto L169
	}
L169:
	;
	v742 = int64(0)
	v746 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[38]))
	if v746 < int32(1) {
		v768 = v742
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v768 != int64(0) {
		goto L166
	} else {
		goto L178
	}
L171:
	;
	goto L170
L172:
	;
	v750 = v742
	v751 = int32(0)
	goto L173
L173:
	;
	v752 = F_dbHasNoKeys(m, v751)
	mBase = m.M
	if v752 != 0 {
		v762 = v750
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v768 = v762
	goto L171
L175:
	;
	v764 = v751 + int32(1)
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[38]))
	if v764 < v766 {
		v750 = v762
		v751 = v764
		goto L173
	} else {
		goto L177
	}
L176:
	;
	v754 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[39]))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v754+v751<<(uint(int32(2))%32))))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	v760 = F_kvstoreSize(m, v759)
	mBase = m.M
	v762 = v760 + v750
	goto L175
L177:
	;
	goto L174
L178:
	;
	goto L167
L179:
	;
	F__serverLog(m, int32(2), v790, int32(0))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L14
	} else {
		goto L188
	}
L180:
	;
	v780 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[37]))
	if v780 != int32(2) {
		goto L165
	} else {
		goto L184
	}
L181:
	;
	if v772 != 0 {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v775 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(2) < v775 {
		goto L166
	} else {
		goto L183
	}
L183:
	;
	v790 = int32(_a_F_dualChannelFullSyncWithPrimary_33)
	goto L179
L184:
	;
	v783 = F_moduleAllModulesHandleReplAsyncLoad(m)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L14
	} else {
		goto L185
	}
L185:
	;
	if v783 != 0 {
		goto L165
	} else {
		goto L186
	}
L186:
	;
	v786 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[2]))
	if int32(2) < v786 {
		goto L166
	} else {
		goto L187
	}
L187:
	;
	v790 = int32(_a_F_dualChannelFullSyncWithPrimary_34)
	goto L179
L188:
	;
	goto L166
L189:
	;
	if v801 != int32(-1) {
		goto L164
	} else {
		goto L190
	}
L190:
	;
	F__serverAssert(m, int32(_a_F_dualChannelFullSyncWithPrimary_35), int32(_a_F_dualChannelFullSyncWithPrimary_1), int32(3179))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L14
	} else {
		goto L191
	}
L191:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L192:
	;
	if v816 == int32(-1) {
		goto L6
	} else {
		goto L193
	}
L193:
	;
	goto L164
L194:
	;
	goto L1
L195:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v867
	F__serverLog(m, int32(3), int32(_a_F_dualChannelFullSyncWithPrimary_36), v11+int32(48))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L14
	} else {
		goto L200
	}
L200:
	;
	v882 = v867
	goto L3
L201:
	;
	goto L2
L202:
	;
	F_replicationAbortDualChannelSyncTransfer(m)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L14
	} else {
		goto L205
	}
L203:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v896)))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v899)+52))
	m.T0[v900].(func(*base.Module, int32))(m, v896)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L14
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[1])) = int32(0)
	goto L202
L205:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dualChannelFullSyncWithPrimary[3])) = int32(1)
	goto L1
}
func F_dualChannelReplMainConnRecvPsyncReply(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_replicaProcessPsyncReply(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		switch v12 + int32(-1) {
		default:
			v49 = v3
			m.G0 = v8 + int32(16)
			return v49
		case 1:
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelReplMainConnRecvPsyncReply[0]))
			if int32(2) < v19 {
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelReplMainConnRecvPsyncReply[1]))
				if v33 != int32(2) {
				} else {
				}
				v38 = F_dualChannelSyncHandlePsync(m)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v49 = v3
					m.G0 = v8 + int32(16)
					return v49
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelReplMainConnRecvPsyncReply[2]))
				if v25 != 0 {
					v26 = int32(_a_F_dualChannelReplMainConnRecvPsyncReply_0)
				} else {
					v26 = int32(_a_F_dualChannelReplMainConnRecvPsyncReply_1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v26
				F__serverLog(m, int32(2), int32(_a_F_dualChannelReplMainConnRecvPsyncReply_2), v8)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelReplMainConnRecvPsyncReply[1]))
					if v33 != int32(2) {
					} else {
					}
					v38 = F_dualChannelSyncHandlePsync(m)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v49 = v3
						m.G0 = v8 + int32(16)
						return v49
					}
				}
			}
		case 2:
			v43 = int32(_a_F_dualChannelReplMainConnRecvPsyncReply_3)
			v44 = F_sdsnew(m, v43)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v44
				v49 = int32(-1)
				m.G0 = v8 + int32(16)
				return v49
			}
		case 3:
			v43 = int32(_a_F_dualChannelReplMainConnRecvPsyncReply_4)
			v44 = F_sdsnew(m, v43)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v44
				v49 = int32(-1)
				m.G0 = v8 + int32(16)
				return v49
			}
		case 4:
			v43 = int32(_a_F_dualChannelReplMainConnRecvPsyncReply_5)
			v44 = F_sdsnew(m, v43)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v44
				v49 = int32(-1)
				m.G0 = v8 + int32(16)
				return v49
			}
		case 5:
			v43 = int32(_a_F_dualChannelReplMainConnRecvPsyncReply_6)
			v44 = F_sdsnew(m, v43)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v44
				v49 = int32(-1)
				m.G0 = v8 + int32(16)
				return v49
			}
		}
	}
}
func F_dualChannelSyncHandleRdbLoadCompletion(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
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
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSyncHandleRdbLoadCompletion[0]))
	if v3 != int32(5) {
		F__serverAssert(m, int32(_a_F_dualChannelSyncHandleRdbLoadCompletion_0), int32(_a_F_dualChannelSyncHandleRdbLoadCompletion_1), int32(3442))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSyncHandleRdbLoadCompletion[1]))
		if int32(12) < v7 {
			if v7 != int32(13) {
				F__serverAssert(m, int32(_a_F_dualChannelSyncHandleRdbLoadCompletion_2), int32(_a_F_dualChannelSyncHandleRdbLoadCompletion_1), int32(3448))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_dualChannelSyncHandleRdbLoadCompletion[2]))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
				v20 = m.T0[v19].(func(*base.Module, int32, int32) int32)(m, v16, int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_dualChannelSyncSuccess(m)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_dualChannelSyncHandleRdbLoadCompletion[0])) = int32(6)
			return
		}
	}
}
func F_dummy_1(m *base.Module, l0 int32) int32 {
	return l0
}
func F_dummy_4(m *base.Module, l0 int32) {
	return
}
func F_durationAddSample(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int64
	_ = v13
	var v19 int64
	_ = v19
	var v24 int64
	_ = v24
	if int32(3) < l0 {
	} else {
		v9 = l0 * int32(24)
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_durationAddSample[0])))
		*(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_durationAddSample[0]))) = v13 + int64(1)
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_durationAddSample[1])))
		*(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_durationAddSample[1]))) = v19 + l1
		v24 = *(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_durationAddSample[2])))
		if base.Ui64(l1) <= base.Ui64(v24) {
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_durationAddSample[2]))) = l1
		}
	}
	return
}
