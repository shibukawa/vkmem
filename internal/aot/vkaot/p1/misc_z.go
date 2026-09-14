package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_zaddCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_zaddGenericCommand(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_zcalloc_num(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	if l1 == int32(0) {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[803]))
		m.T0[v17].(func(*base.Module, int32))(m, int32(-1))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		if base.I32_wrap_i64(int64(base.Ui64(base.I64_extend_i32_u(l1)*base.I64_extend_i32_u(l0))>>(uint(int64(32))%64))) == int32(0) {
			v24 = l1 * l0
			if base.Ui32(int32(2147483646)) < base.Ui32(v24) {
				v101 = *(*int32)(unsafe.Add(mBase, _consts[803]))
				m.T0[v101].(func(*base.Module, int32))(m, v24)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			} else {
				v27 = int32(1)
				if v24 != 0 {
					v29 = v24
				} else {
					v29 = int32(4)
				}
				v31 = v29 + int32(8)
				v37 = base.I64_extend_i32_u(v27) * base.I64_extend_i32_u(v31)
				v38 = base.I32_wrap_i64(v37)
				if base.Ui32(v31|v27) < base.Ui32(int32(65536)) {
					v49 = v38
				} else {
					if base.I32_wrap_i64(int64(base.Ui64(v37)>>(uint(int64(32))%64))) != int32(0) {
						v48 = int32(-1)
					} else {
						v48 = v38
					}
					v49 = v48
				}
				v51 = F_emscripten_builtin_malloc(m, v49)
				mBase = m.M
				if v51 == int32(0) {
				} else {
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+int32(-4)))))
					if v56&int32(3) == int32(0) {
					} else {
						v62 = F___memset(m, v51, int32(0), v49)
						mBase = m.M
					}
				}
				if v51 == int32(0) {
					v101 = *(*int32)(unsafe.Add(mBase, _consts[803]))
					m.T0[v101].(func(*base.Module, int32))(m, v24)
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v51))) = v29
					v67 = *(*int32)(unsafe.Add(mBase, _consts[525]))
					if v67 != int32(-1) {
						v78 = v67
					} else {
						v70 = int32(0)
						v72 = *(*int32)(unsafe.Add(mBase, _consts[315]))
						*(*int32)(unsafe.Add(mBase, _consts[525])) = v72
						*(*int32)(unsafe.Add(mBase, _consts[315])) = v72 + int32(1)
						v78 = v72
					}
					if v78 < int32(260) {
						v87 = v78 << (uint(int32(2)) % 32)
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_consts[320])))
						*(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_consts[320]))) = v90 + v31
					} else {
						v81 = int32(0)
						v83 = *(*int32)(unsafe.Add(mBase, _consts[316]))
						*(*int32)(unsafe.Add(mBase, _consts[316])) = v83 + v31
					}
					return v51 + int32(8)
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[803]))
			m.T0[v17].(func(*base.Module, int32))(m, int32(-1))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	}
}
func F_zdiff(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int64
	_ = v130
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int64
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 float64
	_ = v273
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v293 int64
	_ = v293
	var v297 int64
	_ = v297
	var v298 int64
	_ = v298
	var v306 int64
	_ = v306
	var v308 int64
	_ = v308
	var v312 int32
	_ = v312
	var v316 int64
	_ = v316
	var v319 int64
	_ = v319
	var v321 int64
	_ = v321
	var v324 int64
	_ = v324
	var v335 int64
	_ = v335
	var v338 int64
	_ = v338
	var v349 int64
	_ = v349
	var v352 int64
	_ = v352
	var v365 int64
	_ = v365
	var v372 int64
	_ = v372
	var v377 int32
	_ = v377
	var v380 int64
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int64
	_ = v396
	var v404 int64
	_ = v404
	var v407 int64
	_ = v407
	var v419 int64
	_ = v419
	var v421 int32
	_ = v421
	var v423 int64
	_ = v423
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v440 int64
	_ = v440
	var v448 int64
	_ = v448
	var v451 int64
	_ = v451
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int64
	_ = v465
	var v473 int64
	_ = v473
	var v475 int64
	_ = v475
	var v480 int64
	_ = v480
	var v489 int32
	_ = v489
	var v490 int64
	_ = v490
	var v501 int64
	_ = v501
	var v506 int64
	_ = v506
	var v511 int64
	_ = v511
	var v514 int64
	_ = v514
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v662 int32
	_ = v662
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int64
	_ = v691
	var v712 int32
	_ = v712
	var v719 int32
	_ = v719
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int64
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 float64
	_ = v764
	var v765 int32
	_ = v765
	var v775 int32
	_ = v775
	var v784 int64
	_ = v784
	var v788 int64
	_ = v788
	var v789 int64
	_ = v789
	var v797 int64
	_ = v797
	var v799 int64
	_ = v799
	var v803 int32
	_ = v803
	var v807 int64
	_ = v807
	var v810 int64
	_ = v810
	var v812 int64
	_ = v812
	var v815 int64
	_ = v815
	var v826 int64
	_ = v826
	var v829 int64
	_ = v829
	var v840 int64
	_ = v840
	var v843 int64
	_ = v843
	var v856 int64
	_ = v856
	var v863 int64
	_ = v863
	var v868 int32
	_ = v868
	var v871 int64
	_ = v871
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v887 int64
	_ = v887
	var v895 int64
	_ = v895
	var v898 int64
	_ = v898
	var v910 int64
	_ = v910
	var v912 int32
	_ = v912
	var v914 int64
	_ = v914
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v931 int64
	_ = v931
	var v939 int64
	_ = v939
	var v942 int64
	_ = v942
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int64
	_ = v956
	var v964 int64
	_ = v964
	var v966 int64
	_ = v966
	var v971 int64
	_ = v971
	var v980 int32
	_ = v980
	var v981 int64
	_ = v981
	var v992 int64
	_ = v992
	var v997 int64
	_ = v997
	var v1002 int64
	_ = v1002
	var v1005 int64
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int64
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1051 int32
	_ = v1051
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1077 int32
	_ = v1077
	var v1092 int32
	_ = v1092
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1229 int32
	_ = v1229
	v21 = m.G0
	v23 = v21 - int32(128)
	m.G0 = v23
	v25 = F_zuiLength(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v23 + int32(128)
	return
L2:
	;
	return
L3:
	;
	if v25 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = int32(1)
	if l1 < v29 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v644 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v643)+28)))
	v646 = v644 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v643)+28)) = uint16(v646)
	goto L91
L6:
	;
	v120 = int32(40)
	F_qsort(m, l0+v120, l1+int32(-1), v120, int32(1096))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L19
	}
L7:
	;
	v32 = F_zuiLength(m, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v34 = base.I64_extend_i32_u(v32)
	v35 = F_zuiLength(m, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v37 = base.I64_extend_i32_u(v35)
	if l1 == int32(1) {
		v84 = v34
		v85 = v37
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if base.Ui64(v85) < base.Ui64(int64(base.Ui64(v84)>>(uint(int64(1))%64))) {
		goto L5
	} else {
		goto L18
	}
L11:
	;
	v46 = v29
	v47 = v34
	v48 = v37
	goto L12
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v63 = l0 + v46*int32(40)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v60 == v64 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v84 = v69
	v85 = v73
	goto L10
L14:
	;
	v66 = F_zuiLength(m, l0)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v69 = v47 + base.I64_extend_i32_u(v66)
	v70 = F_zuiLength(m, v63)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v73 = v48 + base.I64_extend_i32_u(v70)
	v75 = v46 + int32(1)
	if v75 != l1 {
		v46 = v75
		v47 = v69
		v48 = v73
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	goto L6
L19:
	;
	v130 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(64)))) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(56)))) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(48)))) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(40)))) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(32)))) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(24)))) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(16)))) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v130
	F_zuiInitIterator(m, l0)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v166 = F_zuiNext(m, l0, v23+int32(8))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L22
	}
L21:
	;
	F_zuiClearIterator(m, l0)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L2
	} else {
		goto L90
	}
L22:
	;
	if v166 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	goto L24
L24:
	;
	v190 = int32(1)
	if l1 <= v190 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L21
L26:
	;
	v619 = F_zuiNext(m, l0, v23+int32(8))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L2
	} else {
		goto L88
	}
L27:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v249&int32(1) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L28:
	;
	v199 = v190
	goto L29
L29:
	;
	v215 = l0 + v199*int32(40)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v216 == v217 {
		goto L26
	} else {
		goto L31
	}
L30:
	;
	goto L27
L31:
	;
	v223 = F_zuiFind(m, v215, v23+int32(8), v23+int32(80))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	if v223 != 0 {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v226 = v199 + int32(1)
	if v226 != l1 {
		v199 = v226
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	v273 = *(*float64)(unsafe.Add(mBase, uint32(v23)+64))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v284 = *(*int32)(unsafe.Add(mBase, _consts[466]))
	if int32(311) < v284 {
		goto L47
	} else {
		goto L48
	}
L36:
	;
	if v248 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v249 & int32(-2)
	v272 = v248
	goto L35
L38:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	if v263 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v261 = F_sdsdup(m, v248)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v272 = v261
	goto L35
L41:
	;
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	v270 = F_sdsfromlonglong(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L2
	} else {
		goto L44
	}
L42:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	v267 = F_sdsnewlen(m, v263, v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v272 = v267
	goto L35
L44:
	;
	v272 = v270
	goto L35
L45:
	;
	v517 = int32(1)
	if v514 == int64(0) {
		goto L61
	} else {
		goto L62
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, _consts[466])) = v489
	v501 = int64(base.Ui64(v490)>>(uint(int64(29))%64))&int64(22906492245) ^ v490
	v506 = v501<<(uint(int64(17))%64)&int64(8202884508482404352) ^ v501
	v511 = v506<<(uint(int64(37))%64)&int64(-2270628950310912) ^ v506
	v514 = int64(base.Ui64(v511)>>(uint(int64(43))%64)) ^ v511
	goto L45
L47:
	;
	if v284 == int32(313) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v284<<(uint(int32(3))%32))+uint32(_consts[467])))
	v489 = v284 + int32(1)
	v490 = v293
	goto L46
L49:
	;
	v377 = int32(0)
	v380 = v372
	goto L55
L50:
	;
	v298 = int64(5489)
	*(*int64)(unsafe.Add(mBase, _consts[467])) = v298
	v306 = int64(1)
	v308 = v298
	goto L52
L51:
	;
	v297 = *(*int64)(unsafe.Add(mBase, _consts[467]))
	v372 = v297
	goto L49
L52:
	;
	v312 = int32(3)
	v316 = int64(62)
	v319 = int64(6364136223846793005)
	v321 = (int64(base.Ui64(v308)>>(uint(v316)%64))^v308)*v319 + v306
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v306)<<(uint(v312)%32))+uint32(_consts[467]))) = v321
	v324 = v306 + int64(1)
	v335 = (int64(base.Ui64(v321)>>(uint(v316)%64))^v321)*v319 + v324
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v324)<<(uint(v312)%32))+uint32(_consts[467]))) = v335
	v338 = v306 + int64(2)
	v349 = (int64(base.Ui64(v335)>>(uint(v316)%64))^v335)*v319 + v338
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v338)<<(uint(v312)%32))+uint32(_consts[467]))) = v349
	v352 = v306 + int64(3)
	if v352 == int64(312) {
		v372 = v298
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v365 = (int64(base.Ui64(v349)>>(uint(int64(62))%64))^v349)*int64(6364136223846793005) + v352
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v352)<<(uint(int32(3))%32))+uint32(_consts[467]))) = v365
	v306 = v306 + int64(4)
	v308 = v365
	goto L52
L55:
	;
	v386 = int32(3)
	v387 = v377 << (uint(v386) % 32)
	v390 = int32(1)
	v391 = v377 + v390
	v396 = *(*int64)(unsafe.Add(mBase, uint32(v391<<(uint(v386)%32))+uint32(_consts[467])))
	v404 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v396)&v390<<(uint(v386)%32))+uint32(_consts[468])))
	v407 = *(*int64)(unsafe.Add(mBase, uint32(v387)+uint32(_consts[469])))
	*(*int64)(unsafe.Add(mBase, uint32(v387)+uint32(_consts[467]))) = v404 ^ v407 ^ int64(base.Ui64(v380&int64(-2147483648)|v396&int64(2147483646))>>(uint(int64(1))%64))
	if v391 != int32(156) {
		v377 = v391
		v380 = v396
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v419 = *(*int64)(unsafe.Add(mBase, _consts[469]))
	v421 = int32(156)
	v423 = v419
	goto L58
L57:
	;
	goto L56
L58:
	;
	v430 = int32(3)
	v431 = v421 << (uint(v430) % 32)
	v434 = int32(1)
	v435 = v421 + v434
	v440 = *(*int64)(unsafe.Add(mBase, uint32(v435<<(uint(v430)%32))+uint32(_consts[467])))
	v448 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v440)&v434<<(uint(v430)%32))+uint32(_consts[468])))
	v451 = *(*int64)(unsafe.Add(mBase, uint32(v431)+uint32(_consts[470])))
	*(*int64)(unsafe.Add(mBase, uint32(v431)+uint32(_consts[467]))) = v448 ^ v451 ^ int64(base.Ui64(v423&int64(-2147483648)|v440&int64(2147483646))>>(uint(int64(1))%64))
	if v435 != int32(311) {
		v421 = v435
		v423 = v440
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v462 = int32(1)
	v463 = int32(0)
	v465 = *(*int64)(unsafe.Add(mBase, _consts[467]))
	v473 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v465)&v462<<(uint(int32(3))%32))+uint32(_consts[468])))
	v475 = *(*int64)(unsafe.Add(mBase, _consts[471]))
	v480 = *(*int64)(unsafe.Add(mBase, _consts[472]))
	*(*int64)(unsafe.Add(mBase, _consts[472])) = v473 ^ v475 ^ int64(base.Ui64(v465&int64(2147483646)|v480&int64(-2147483648))>>(uint(int64(1))%64))
	v489 = v462
	v490 = v465
	goto L46
L60:
	;
	goto L59
L61:
	;
	v523 = int32(32)
	goto L63
L62:
	;
	v523 = int32(base.Ui32(base.I32_wrap_i64(base.I64_clz(v514)))>>(uint(v517)%32)) + v517
	goto L63
L63:
	;
	v524 = F_zslCreateNode(m, v523, v273, v272)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v526 = F_zslInsertNode(m, v274, v524)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v529 = F_hashtableAdd(m, v528, v526)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v532 = v272 + int32(-1)
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532))))
	v535 = v533 & int32(7)
	switch v535 {
	case 0:
		goto L73
	case 1:
		goto L72
	case 2:
		goto L71
	case 3:
		goto L70
	case 4:
		goto L69
	default:
		v570 = v533
		goto L67
	}
L67:
	;
	switch v570 & int32(7) {
	case 0:
		goto L86
	case 1:
		goto L85
	case 2:
		goto L84
	case 3:
		goto L83
	case 4:
		goto L82
	default:
		v591 = int32(0)
		goto L81
	}
L68:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if base.Ui32(v550) <= base.Ui32(v551) {
		v570 = v533
		goto L67
	} else {
		goto L74
	}
L69:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v272+int32(-17))))
	v550 = v549
	goto L68
L70:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v272+int32(-9))))
	v550 = v546
	goto L68
L71:
	;
	v543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v272+int32(-5)))))
	v550 = v543
	goto L68
L72:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272+int32(-3)))))
	v550 = v540
	goto L68
L73:
	;
	v550 = int32(base.Ui32(v533) >> (uint(int32(3)) % 32))
	goto L68
L74:
	;
	switch v535 {
	default:
		goto L80
	case 1:
		goto L79
	case 2:
		goto L78
	case 3:
		goto L77
	case 4:
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v567
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532))))
	v570 = v569
	goto L67
L76:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v272+int32(-17))))
	v567 = v566
	goto L75
L77:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v272+int32(-9))))
	v567 = v563
	goto L75
L78:
	;
	v560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v272+int32(-5)))))
	v567 = v560
	goto L75
L79:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272+int32(-3)))))
	v567 = v557
	goto L75
L80:
	;
	v567 = int32(base.Ui32(v533) >> (uint(int32(3)) % 32))
	goto L75
L81:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v592 + v591
	F_sdsfree(m, v272)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L2
	} else {
		goto L87
	}
L82:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v272+int32(-17))))
	v591 = v590
	goto L81
L83:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v272+int32(-9))))
	v591 = v587
	goto L81
L84:
	;
	v584 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v272+int32(-5)))))
	v591 = v584
	goto L81
L85:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272+int32(-3)))))
	v591 = v581
	goto L81
L86:
	;
	v591 = int32(base.Ui32(v570&int32(248)) >> (uint(int32(3)) % 32))
	goto L81
L87:
	;
	goto L26
L88:
	;
	if v619 != 0 {
		goto L24
	} else {
		goto L89
	}
L89:
	;
	goto L25
L90:
	;
	goto L1
L91:
	;
	v662 = int32(0)
	v670 = v662
	v677 = v662
	goto L93
L92:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_hashtableResumeAutoShrink(m, v1114)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L2
	} else {
		goto L156
	}
L93:
	;
	v686 = l0 + v677*int32(40)
	v687 = F_zuiLength(m, v686)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L2
	} else {
		goto L96
	}
L94:
	;
	goto L92
L95:
	;
	v1092 = v677 + int32(1)
	if v1092 != l1 {
		v670 = v1077
		v677 = v1092
		goto L93
	} else {
		goto L155
	}
L96:
	;
	if v687 == int32(0) {
		v1077 = v670
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v691 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(64)))) = v691
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(56)))) = v691
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(48)))) = v691
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(40)))) = v691
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(32)))) = v691
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(24)))) = v691
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(16)))) = v691
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v691
	F_zuiInitIterator(m, v686)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L2
	} else {
		goto L98
	}
L98:
	;
	v719 = v670
	goto L100
L99:
	;
	F_zuiClearIterator(m, v686)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L2
	} else {
		goto L153
	}
L100:
	;
	v735 = F_zuiNext(m, v686, v23+int32(8))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L2
	} else {
		goto L102
	}
L101:
	;
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)))
	if v1051&int32(1) == int32(0) {
		goto L149
	} else {
		goto L150
	}
L102:
	;
	if v735 == int32(0) {
		goto L99
	} else {
		goto L103
	}
L103:
	;
	if v677 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v1046 != 0 {
		v719 = v1046
		goto L100
	} else {
		goto L148
	}
L105:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	if v1026 != 0 {
		v1042 = v1026
		goto L140
	} else {
		goto L141
	}
L106:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v740&int32(1) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v764 = *(*float64)(unsafe.Add(mBase, uint32(v23)+64))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v775 = *(*int32)(unsafe.Add(mBase, _consts[466]))
	if int32(311) < v775 {
		goto L119
	} else {
		goto L120
	}
L108:
	;
	if v739 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v740 & int32(-2)
	v763 = v739
	goto L107
L110:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	if v754 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	v752 = F_sdsdup(m, v739)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	v763 = v752
	goto L107
L113:
	;
	v760 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	v761 = F_sdsfromlonglong(m, v760)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L2
	} else {
		goto L116
	}
L114:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	v758 = F_sdsnewlen(m, v754, v757)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L2
	} else {
		goto L115
	}
L115:
	;
	v763 = v758
	goto L107
L116:
	;
	v763 = v761
	goto L107
L117:
	;
	v1008 = int32(1)
	if v1005 == int64(0) {
		goto L133
	} else {
		goto L134
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, _consts[466])) = v980
	v992 = int64(base.Ui64(v981)>>(uint(int64(29))%64))&int64(22906492245) ^ v981
	v997 = v992<<(uint(int64(17))%64)&int64(8202884508482404352) ^ v992
	v1002 = v997<<(uint(int64(37))%64)&int64(-2270628950310912) ^ v997
	v1005 = int64(base.Ui64(v1002)>>(uint(int64(43))%64)) ^ v1002
	goto L117
L119:
	;
	if v775 == int32(313) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v784 = *(*int64)(unsafe.Add(mBase, uint32(v775<<(uint(int32(3))%32))+uint32(_consts[467])))
	v980 = v775 + int32(1)
	v981 = v784
	goto L118
L121:
	;
	v868 = int32(0)
	v871 = v863
	goto L127
L122:
	;
	v789 = int64(5489)
	*(*int64)(unsafe.Add(mBase, _consts[467])) = v789
	v797 = int64(1)
	v799 = v789
	goto L124
L123:
	;
	v788 = *(*int64)(unsafe.Add(mBase, _consts[467]))
	v863 = v788
	goto L121
L124:
	;
	v803 = int32(3)
	v807 = int64(62)
	v810 = int64(6364136223846793005)
	v812 = (int64(base.Ui64(v799)>>(uint(v807)%64))^v799)*v810 + v797
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v797)<<(uint(v803)%32))+uint32(_consts[467]))) = v812
	v815 = v797 + int64(1)
	v826 = (int64(base.Ui64(v812)>>(uint(v807)%64))^v812)*v810 + v815
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v815)<<(uint(v803)%32))+uint32(_consts[467]))) = v826
	v829 = v797 + int64(2)
	v840 = (int64(base.Ui64(v826)>>(uint(v807)%64))^v826)*v810 + v829
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v829)<<(uint(v803)%32))+uint32(_consts[467]))) = v840
	v843 = v797 + int64(3)
	if v843 == int64(312) {
		v863 = v789
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v856 = (int64(base.Ui64(v840)>>(uint(int64(62))%64))^v840)*int64(6364136223846793005) + v843
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v843)<<(uint(int32(3))%32))+uint32(_consts[467]))) = v856
	v797 = v797 + int64(4)
	v799 = v856
	goto L124
L127:
	;
	v877 = int32(3)
	v878 = v868 << (uint(v877) % 32)
	v881 = int32(1)
	v882 = v868 + v881
	v887 = *(*int64)(unsafe.Add(mBase, uint32(v882<<(uint(v877)%32))+uint32(_consts[467])))
	v895 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v887)&v881<<(uint(v877)%32))+uint32(_consts[468])))
	v898 = *(*int64)(unsafe.Add(mBase, uint32(v878)+uint32(_consts[469])))
	*(*int64)(unsafe.Add(mBase, uint32(v878)+uint32(_consts[467]))) = v895 ^ v898 ^ int64(base.Ui64(v871&int64(-2147483648)|v887&int64(2147483646))>>(uint(int64(1))%64))
	if v882 != int32(156) {
		v868 = v882
		v871 = v887
		goto L127
	} else {
		goto L129
	}
L128:
	;
	v910 = *(*int64)(unsafe.Add(mBase, _consts[469]))
	v912 = int32(156)
	v914 = v910
	goto L130
L129:
	;
	goto L128
L130:
	;
	v921 = int32(3)
	v922 = v912 << (uint(v921) % 32)
	v925 = int32(1)
	v926 = v912 + v925
	v931 = *(*int64)(unsafe.Add(mBase, uint32(v926<<(uint(v921)%32))+uint32(_consts[467])))
	v939 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v931)&v925<<(uint(v921)%32))+uint32(_consts[468])))
	v942 = *(*int64)(unsafe.Add(mBase, uint32(v922)+uint32(_consts[470])))
	*(*int64)(unsafe.Add(mBase, uint32(v922)+uint32(_consts[467]))) = v939 ^ v942 ^ int64(base.Ui64(v914&int64(-2147483648)|v931&int64(2147483646))>>(uint(int64(1))%64))
	if v926 != int32(311) {
		v912 = v926
		v914 = v931
		goto L130
	} else {
		goto L132
	}
L131:
	;
	v953 = int32(1)
	v954 = int32(0)
	v956 = *(*int64)(unsafe.Add(mBase, _consts[467]))
	v964 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v956)&v953<<(uint(int32(3))%32))+uint32(_consts[468])))
	v966 = *(*int64)(unsafe.Add(mBase, _consts[471]))
	v971 = *(*int64)(unsafe.Add(mBase, _consts[472]))
	*(*int64)(unsafe.Add(mBase, _consts[472])) = v964 ^ v966 ^ int64(base.Ui64(v956&int64(2147483646)|v971&int64(-2147483648))>>(uint(int64(1))%64))
	v980 = v953
	v981 = v956
	goto L118
L132:
	;
	goto L131
L133:
	;
	v1014 = int32(32)
	goto L135
L134:
	;
	v1014 = int32(base.Ui32(base.I32_wrap_i64(base.I64_clz(v1005)))>>(uint(v1008)%32)) + v1008
	goto L135
L135:
	;
	v1015 = F_zslCreateNode(m, v1014, v764, v763)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L2
	} else {
		goto L136
	}
L136:
	;
	v1017 = F_zslInsertNode(m, v765, v1015)
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L2
	} else {
		goto L137
	}
L137:
	;
	F_sdsfree(m, v763)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L2
	} else {
		goto L138
	}
L138:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1022 = F_hashtableAdd(m, v1021, v1017)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L2
	} else {
		goto L139
	}
L139:
	;
	v1046 = v719 + int32(1)
	goto L104
L140:
	;
	v1043 = F_zsetRemoveFromSkiplist(m, l2, v1042)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L2
	} else {
		goto L147
	}
L141:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	if v1027 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v1036
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v1038 | int32(1)
	v1042 = v1036
	goto L140
L143:
	;
	v1033 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	v1034 = F_sdsfromlonglong(m, v1033)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L2
	} else {
		goto L146
	}
L144:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	v1031 = F_sdsnewlen(m, v1027, v1030)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L2
	} else {
		goto L145
	}
L145:
	;
	v1036 = v1031
	goto L142
L146:
	;
	v1036 = v1034
	goto L142
L147:
	;
	v1046 = v719 - v1043
	goto L104
L148:
	;
	goto L101
L149:
	;
	F_zuiClearIterator(m, v686)
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L2
	} else {
		goto L152
	}
L150:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	F_sdsfree(m, v1056)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L2
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = int32(0)
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v1061 & int32(-2)
	goto L149
L152:
	;
	goto L92
L153:
	;
	if v719 == int32(0) {
		goto L92
	} else {
		goto L154
	}
L154:
	;
	v1077 = v719
	goto L95
L155:
	;
	goto L94
L156:
	;
	v1117 = int32(0)
	v1119 = v23 + int32(80)
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1119)+14)) = uint8(v1117)
	*(*int32)(unsafe.Add(mBase, uint32(v1119))) = v1120
	*(*int32)(unsafe.Add(mBase, uint32(v1119)+24)) = v1117
	*(*uint8)(unsafe.Add(mBase, uint32(v1119)+15)) = uint8(v1117)
	*(*int32)(unsafe.Add(mBase, uint32(v1119)+8)) = int32(-1)
	if v1120 == v1117 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v1143 = F_hashtableNext(m, v23+int32(80), v23+int32(76))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L2
	} else {
		goto L162
	}
L158:
	;
	goto L157
L159:
	;
	goto L158
L161:
	;
	F_hashtableCleanupIterator(m, v23+int32(80))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L2
	} else {
		goto L177
	}
L162:
	;
	if v1143 == int32(0) {
		v1206 = v1117
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v1147 = v1117
	goto L164
L164:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v1170 = v1168 + int32(16)
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v1170)))
	v1174 = v1170 + v1171<<(uint(int32(3))%32)
	v1175 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1174))))
	v1176 = v1174 + v1175
	v1177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1176))))
	switch v1177 & int32(7) {
	case 0:
		goto L171
	case 1:
		goto L170
	case 2:
		goto L169
	case 3:
		goto L168
	case 4:
		goto L167
	default:
		v1194 = int32(0)
		goto L166
	}
L165:
	;
	v1206 = v1199
	goto L161
L166:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1195 + v1194
	if base.Ui32(v1147) < base.Ui32(v1194) {
		goto L172
	} else {
		goto L173
	}
L167:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1176+int32(-16))))
	v1194 = v1193
	goto L166
L168:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1176+int32(-8))))
	v1194 = v1190
	goto L166
L169:
	;
	v1187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1176+int32(-4)))))
	v1194 = v1187
	goto L166
L170:
	;
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1176+int32(-2)))))
	v1194 = v1184
	goto L166
L171:
	;
	v1194 = int32(base.Ui32(v1177) >> (uint(int32(3)) % 32))
	goto L166
L172:
	;
	v1199 = v1194
	goto L174
L173:
	;
	v1199 = v1147
	goto L174
L174:
	;
	v1204 = F_hashtableNext(m, v23+int32(80), v23+int32(76))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L2
	} else {
		goto L175
	}
L175:
	;
	if v1204 != 0 {
		v1147 = v1199
		goto L164
	} else {
		goto L176
	}
L176:
	;
	goto L165
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1206
	goto L1
}
func F_zeroinfnan(m *base.Module, l0 int64) int32 {
	return base.B2i32(base.Ui64(l0<<(uint(int64(1))%64)+int64(9007199254740992)) < base.Ui64(int64(9007199254740993)))
}
func F_zinterCardCommand(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = int32(1)
	F_zunionInterDiffGenericCommand(m, l0, int32(0), v3, int32(2), v3)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_zipEntrySafe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v81 int32
	_ = v81
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v323 int32
	_ = v323
	v14 = l0 + l1 + int32(-1)
	v16 = l0 + int32(10)
	v17 = base.B2i32(base.Ui32(l2) < base.Ui32(v16))
	if base.Ui32(l2) < base.Ui32(v16) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v323
L2:
	;
	v323 = int32(1)
	goto L1
L3:
	;
	v140 = int32(0)
	if base.Ui32(l2) < base.Ui32(v16) {
		v323 = v140
		goto L1
	} else {
		goto L37
	}
L4:
	;
	if base.Ui32(v14) <= base.Ui32(l2+int32(10)) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.Ui32(v23) < base.Ui32(int32(254)) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v26 = int32(1)
	goto L8
L7:
	;
	v26 = int32(5)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v26
	if base.Ui32(int32(253)) < base.Ui32(v23) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v34
	v36 = l2 + v26
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v38 = int32(192)
	if base.Ui32(v37) < base.Ui32(v38) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(1))))
	v34 = v33
	goto L9
L11:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v34 = v30
	goto L9
L12:
	;
	v42 = v37 & v38
	goto L14
L13:
	;
	v42 = v37
	goto L14
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)) = uint8(v42)
	if base.Ui32(int32(191)) < base.Ui32(v37) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = l2
	v128 = v126 + v26
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v128
	v130 = int32(0)
	v132 = l2 + v128 + v124
	if base.Ui32(v132) < base.Ui32(v16) {
		v323 = v130
		goto L1
	} else {
		goto L32
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v120
	v124 = v120
	v126 = v122
	goto L15
L17:
	;
	v114 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v114
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v120 = v117 & int32(63)
	v122 = v114
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v26
	return int32(0)
L19:
	;
	v81 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v81
	switch v37 + int32(-208) {
	case 0:
		goto L26
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L24
	case 16:
		goto L25
	default:
		goto L27
	}
L20:
	;
	v46 = int32(2)
	switch int32(base.Ui32(v37)>>(uint(int32(6))%32)) ^ v46 {
	default:
		goto L22
	case 1:
		goto L21
	case 2:
		goto L17
	case 3:
		goto L23
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = int64(0)
	goto L18
L22:
	;
	v60 = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v60
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v36)+1))
	v64 = int32(24)
	v66 = int32(65280)
	v68 = int32(8)
	v120 = v63<<(uint(v64)%32) | v63&v66<<(uint(v68)%32) | (int32(base.Ui32(v63)>>(uint(v68)%32))&v66 | int32(base.Ui32(v63)>>(uint(v64)%32)))
	v122 = v60
	goto L16
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = int32(2)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	v120 = v53&int32(63)<<(uint(int32(8))%32) | v58
	v122 = v46
	goto L16
L24:
	;
	v96 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v96
	v99 = int32(1)
	if base.Ui32(int32(241)) < base.Ui32((v37+v99)&int32(255)) {
		v124 = v96
		v126 = v99
		goto L15
	} else {
		goto L31
	}
L25:
	;
	v120 = int32(8)
	v122 = v81
	goto L16
L26:
	;
	v120 = int32(4)
	v122 = v81
	goto L16
L27:
	;
	switch v37 + int32(-240) {
	case 0:
		goto L28
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
		goto L24
	case 14:
		v120 = int32(1)
		v122 = v81
		goto L16
	default:
		goto L29
	}
L28:
	;
	v120 = int32(3)
	v122 = v81
	goto L16
L29:
	;
	if v37 != int32(192) {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v120 = int32(2)
	v122 = v81
	goto L16
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = int32(0)
	goto L18
L32:
	;
	if base.Ui32(v14) < base.Ui32(v132) {
		v323 = v130
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if l4 == int32(0) {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v137 = l2 - v34
	if base.Ui32(v137) < base.Ui32(v16) {
		v323 = v130
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if base.Ui32(v137) <= base.Ui32(v14) {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v323 = v130
	goto L1
L37:
	;
	if base.Ui32(v14) < base.Ui32(l2) {
		v323 = v140
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.Ui32(v144) < base.Ui32(int32(254)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v147 = int32(1)
	goto L41
L40:
	;
	v147 = int32(5)
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v147
	v149 = l2 + v147
	if base.Ui32(v149) < base.Ui32(v16) {
		v323 = v140
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if base.Ui32(v14) < base.Ui32(v149) {
		v323 = v140
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	v153 = int32(192)
	if base.Ui32(v152) < base.Ui32(v153) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v157 = v152 & v153
	goto L46
L45:
	;
	v157 = v152
	goto L46
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)) = uint8(v157)
	v159 = int32(1)
	switch v157 + int32(-208) {
	case 0, 16:
		v184 = v159
		goto L47
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L48
	default:
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v184
	v186 = v149 + v184
	if base.Ui32(v186) < base.Ui32(v16) {
		v323 = v140
		goto L1
	} else {
		goto L58
	}
L48:
	;
	if base.Ui32((v157+int32(15))&int32(255)) < base.Ui32(int32(13)) {
		v184 = v159
		goto L47
	} else {
		goto L52
	}
L49:
	;
	switch v157 + int32(-240) {
	case 0, 14:
		v184 = v159
		goto L47
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
		goto L48
	default:
		goto L50
	}
L50:
	;
	if v157 == int32(192) {
		v184 = v159
		goto L47
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	if v157 == int32(0) {
		v184 = v159
		goto L47
	} else {
		goto L53
	}
L53:
	;
	if v157 == int32(128) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v184 = int32(5)
	goto L47
L55:
	;
	if v157 != int32(64) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = int32(255)
	return int32(0)
L57:
	;
	v184 = int32(2)
	goto L47
L58:
	;
	if base.Ui32(v14) < base.Ui32(v186) {
		v323 = v140
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.Ui32(v191) < base.Ui32(int32(254)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v194 = int32(1)
	goto L62
L61:
	;
	v194 = int32(5)
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v194
	if base.Ui32(int32(253)) < base.Ui32(v191) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v202
	if base.Ui32(int32(191)) < base.Ui32(v152) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(1))))
	v202 = v201
	goto L63
L65:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v202 = v198
	goto L63
L66:
	;
	v298 = v296 + v194
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v298
	v301 = l2 + v298 + v297
	if base.Ui32(v301) < base.Ui32(v16) {
		v323 = v140
		goto L1
	} else {
		goto L84
	}
L67:
	;
	v292 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v292
	v296 = v254
	v297 = v292
	goto L66
L68:
	;
	v296 = int32(0)
	v297 = v290
	goto L66
L69:
	;
	v254 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v254
	switch v152 + int32(-208) {
	case 0:
		goto L78
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L75
	case 16:
		goto L77
	default:
		goto L79
	}
L70:
	;
	v206 = int32(2)
	switch int32(base.Ui32(v152)>>(uint(int32(6))%32)) ^ v206 {
	default:
		goto L72
	case 1:
		goto L71
	case 2:
		goto L74
	case 3:
		goto L73
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = int64(0)
	v290 = int32(0)
	goto L68
L72:
	;
	v230 = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v230
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l2+v194)+1))
	v235 = int32(24)
	v237 = int32(65280)
	v239 = int32(8)
	v249 = v234<<(uint(v235)%32) | v234&v237<<(uint(v239)%32) | (int32(base.Ui32(v234)>>(uint(v239)%32))&v237 | int32(base.Ui32(v234)>>(uint(v235)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v249
	v296 = v230
	v297 = v249
	goto L66
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = int32(2)
	v221 = l2 + v194
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	v228 = v222&int32(63)<<(uint(int32(8))%32) | v227
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v228
	v296 = v206
	v297 = v228
	goto L66
L74:
	;
	v211 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v211
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v194))))
	v217 = v215 & int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v217
	v296 = v211
	v297 = v217
	goto L66
L75:
	;
	v276 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v276
	v279 = int32(1)
	if base.Ui32(int32(241)) < base.Ui32((v152+v279)&int32(255)) {
		v296 = v279
		v297 = v276
		goto L66
	} else {
		goto L83
	}
L76:
	;
	if v152 == int32(192) {
		goto L67
	} else {
		goto L82
	}
L77:
	;
	v271 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v271
	v296 = v254
	v297 = v271
	goto L66
L78:
	;
	v268 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v268
	v296 = v254
	v297 = v268
	goto L66
L79:
	;
	switch v152 + int32(-240) {
	case 0:
		goto L80
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
		goto L75
	case 14:
		goto L81
	default:
		goto L76
	}
L80:
	;
	v265 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v265
	v296 = v254
	v297 = v265
	goto L66
L81:
	;
	v261 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v261
	v296 = v261
	v297 = v261
	goto L66
L82:
	;
	goto L75
L83:
	;
	v286 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v286
	v290 = v286
	goto L68
L84:
	;
	if base.Ui32(v14) < base.Ui32(v301) {
		v323 = v140
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if l4 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = l2
	goto L2
L87:
	;
	v306 = l2 - v202
	if base.Ui32(v306) < base.Ui32(v16) {
		v323 = v140
		goto L1
	} else {
		goto L88
	}
L88:
	;
	if base.Ui32(v14) < base.Ui32(v306) {
		v323 = v140
		goto L1
	} else {
		goto L89
	}
L89:
	;
	goto L86
}
func F_ziplistGet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	if l0 == v5 {
		v88 = v5
		m.G0 = v11 + int32(32)
		return v88
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v16 == int32(255) {
			v88 = v5
			m.G0 = v11 + int32(32)
			return v88
		} else {
			if l1 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
				F_zipEntry(m, l0, v11+int32(4))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)))
					if base.Ui32(int32(191)) < base.Ui32(v35) {
						v45 = v35
						v46 = int32(1)
						if l3 == int32(0) {
							v88 = v46
							m.G0 = v11 + int32(32)
							return v88
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
							v50 = l0 + v49
							switch v45 + int32(-208) {
							case 0:
								v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v50))))
								v85 = v55
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
								v88 = v46
								m.G0 = v11 + int32(32)
								return v88
							case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
								if base.Ui32(int32(13)) < base.Ui32((v45+int32(15))&int32(255)) {
									F__serverAssert(m, int32(_a1646), int32(_a2534), int32(604))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v85 = base.I64_extend_i32_u(v45)&int64(15) + int64(-1)
									*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
									v88 = v46
									m.G0 = v11 + int32(32)
									return v88
								}
							case 16:
								v63 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
								v85 = v63
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
								v88 = v46
								m.G0 = v11 + int32(32)
								return v88
							default:
								switch v45 + int32(-240) {
								case 0:
									v56 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v50))))
									v59 = int64(*(*int8)(unsafe.Add(mBase, uint32(v50+int32(2)))))
									v85 = v56 | v59<<(uint(int64(16))%64)
									*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
									v88 = v46
									m.G0 = v11 + int32(32)
									return v88
								case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
									if base.Ui32(int32(13)) < base.Ui32((v45+int32(15))&int32(255)) {
										F__serverAssert(m, int32(_a1646), int32(_a2534), int32(604))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v85 = base.I64_extend_i32_u(v45)&int64(15) + int64(-1)
										*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
										v88 = v46
										m.G0 = v11 + int32(32)
										return v88
									}
								case 14:
									v83 = int64(*(*int8)(unsafe.Add(mBase, uint32(v50))))
									v85 = v83
									*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
									v88 = v46
									m.G0 = v11 + int32(32)
									return v88
								default:
									if v45 == int32(192) {
										v84 = int64(*(*int16)(unsafe.Add(mBase, uint32(v50))))
										v85 = v84
										*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
										v88 = v46
										m.G0 = v11 + int32(32)
										return v88
									} else {
										if base.Ui32(int32(13)) < base.Ui32((v45+int32(15))&int32(255)) {
											F__serverAssert(m, int32(_a1646), int32(_a2534), int32(604))
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v85 = base.I64_extend_i32_u(v45)&int64(15) + int64(-1)
											*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
											v88 = v46
											m.G0 = v11 + int32(32)
											return v88
										}
									}
								}
							}
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v38
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = l0 + v40
						v88 = int32(1)
						m.G0 = v11 + int32(32)
						return v88
					}
				}
			} else {
				F_zipEntry(m, l0, v11+int32(4))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)))
					if base.Ui32(v26) <= base.Ui32(int32(191)) {
						v88 = int32(1)
						m.G0 = v11 + int32(32)
						return v88
					} else {
						v45 = v26
						v46 = int32(1)
						if l3 == int32(0) {
							v88 = v46
							m.G0 = v11 + int32(32)
							return v88
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
							v50 = l0 + v49
							switch v45 + int32(-208) {
							case 0:
								v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v50))))
								v85 = v55
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
								v88 = v46
								m.G0 = v11 + int32(32)
								return v88
							case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
								if base.Ui32(int32(13)) < base.Ui32((v45+int32(15))&int32(255)) {
									F__serverAssert(m, int32(_a1646), int32(_a2534), int32(604))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v85 = base.I64_extend_i32_u(v45)&int64(15) + int64(-1)
									*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
									v88 = v46
									m.G0 = v11 + int32(32)
									return v88
								}
							case 16:
								v63 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
								v85 = v63
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
								v88 = v46
								m.G0 = v11 + int32(32)
								return v88
							default:
								switch v45 + int32(-240) {
								case 0:
									v56 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v50))))
									v59 = int64(*(*int8)(unsafe.Add(mBase, uint32(v50+int32(2)))))
									v85 = v56 | v59<<(uint(int64(16))%64)
									*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
									v88 = v46
									m.G0 = v11 + int32(32)
									return v88
								case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
									if base.Ui32(int32(13)) < base.Ui32((v45+int32(15))&int32(255)) {
										F__serverAssert(m, int32(_a1646), int32(_a2534), int32(604))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v85 = base.I64_extend_i32_u(v45)&int64(15) + int64(-1)
										*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
										v88 = v46
										m.G0 = v11 + int32(32)
										return v88
									}
								case 14:
									v83 = int64(*(*int8)(unsafe.Add(mBase, uint32(v50))))
									v85 = v83
									*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
									v88 = v46
									m.G0 = v11 + int32(32)
									return v88
								default:
									if v45 == int32(192) {
										v84 = int64(*(*int16)(unsafe.Add(mBase, uint32(v50))))
										v85 = v84
										*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
										v88 = v46
										m.G0 = v11 + int32(32)
										return v88
									} else {
										if base.Ui32(int32(13)) < base.Ui32((v45+int32(15))&int32(255)) {
											F__serverAssert(m, int32(_a1646), int32(_a2534), int32(604))
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v85 = base.I64_extend_i32_u(v45)&int64(15) + int64(-1)
											*(*int64)(unsafe.Add(mBase, uint32(l3))) = v85
											v88 = v46
											m.G0 = v11 + int32(32)
											return v88
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
func F_ziplistValidateIntegrity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
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
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v125 int32
	_ = v125
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v365 int32
	_ = v365
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v420 int32
	_ = v420
	var v431 int32
	_ = v431
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if base.Ui32(l1) < base.Ui32(int32(11)) {
		v431 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v431
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v18 != l1 {
		v431 = v6
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = l1 + int32(-1)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v21))))
	if v23 != int32(255) {
		v431 = v6
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v21) < base.Ui32(v26) {
		v431 = v6
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if l2 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	if v30 == int32(255) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v431 = int32(1)
	goto L1
L8:
	;
	v431 = base.B2i32(v29 == int32(65535)) | base.B2i32(v420 == v29)
	goto L1
L9:
	;
	v407 = int32(0)
	if l1 != int32(11) {
		v431 = v407
		goto L1
	} else {
		goto L114
	}
L10:
	;
	v35 = int32(0)
	v43 = v35
	v44 = l0 + int32(10)
	v46 = v35
	goto L11
L11:
	;
	v48 = v13 + int32(4)
	v58 = l0 + l1 + int32(-1)
	v60 = l0 + int32(10)
	v61 = base.B2i32(base.Ui32(v44) < base.Ui32(v60))
	if base.Ui32(v44) < base.Ui32(v60) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	v399 = int32(0)
	if v395 != l0+l1+int32(-1) {
		v431 = v399
		goto L1
	} else {
		goto L112
	}
L13:
	;
	v391 = v46 + int32(1)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v394 = v392 + v393
	v395 = v44 + v394
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	if v396 != int32(255) {
		v43 = v394
		v44 = v395
		v46 = v391
		goto L11
	} else {
		goto L111
	}
L14:
	;
	v431 = int32(0)
	goto L1
L15:
	;
	if v378 == int32(0) {
		goto L14
	} else {
		goto L105
	}
L16:
	;
	v378 = v365
	goto L15
L17:
	;
	v365 = int32(1)
	goto L16
L18:
	;
	v183 = int32(0)
	if base.Ui32(v44) < base.Ui32(v60) {
		v365 = v183
		goto L16
	} else {
		goto L52
	}
L19:
	;
	if base.Ui32(v58) <= base.Ui32(v44+int32(10)) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if base.Ui32(v67) < base.Ui32(int32(254)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v70 = int32(1)
	goto L23
L22:
	;
	v70 = int32(5)
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v70
	if base.Ui32(int32(253)) < base.Ui32(v67) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v78
	v80 = v44 + v70
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v82 = int32(192)
	if base.Ui32(v81) < base.Ui32(v82) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v44+int32(1))))
	v78 = v77
	goto L24
L26:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v78 = v74
	goto L24
L27:
	;
	v86 = v81 & v82
	goto L29
L28:
	;
	v86 = v81
	goto L29
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+20)) = uint8(v86)
	if base.Ui32(int32(191)) < base.Ui32(v81) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v44
	v171 = v169 + v70
	*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v171
	v173 = int32(0)
	v175 = v44 + v171 + v167
	if base.Ui32(v175) < base.Ui32(v60) {
		v365 = v173
		goto L16
	} else {
		goto L47
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v163
	v167 = v163
	v169 = v165
	goto L30
L32:
	;
	v157 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v157
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v163 = v160 & int32(63)
	v165 = v157
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v70
	v378 = int32(0)
	goto L15
L34:
	;
	v125 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v125
	switch v81 + int32(-208) {
	case 0:
		goto L41
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L39
	case 16:
		goto L40
	default:
		goto L42
	}
L35:
	;
	v90 = int32(2)
	switch int32(base.Ui32(v81)>>(uint(int32(6))%32)) ^ v90 {
	default:
		goto L37
	case 1:
		goto L36
	case 2:
		goto L32
	case 3:
		goto L38
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48)+8)) = int64(0)
	goto L33
L37:
	;
	v104 = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v104
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v80)+1))
	v108 = int32(24)
	v110 = int32(65280)
	v112 = int32(8)
	v163 = v107<<(uint(v108)%32) | v107&v110<<(uint(v112)%32) | (int32(base.Ui32(v107)>>(uint(v112)%32))&v110 | int32(base.Ui32(v107)>>(uint(v108)%32)))
	v165 = v104
	goto L31
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = int32(2)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	v163 = v97&int32(63)<<(uint(int32(8))%32) | v102
	v165 = v90
	goto L31
L39:
	;
	v140 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v140
	v143 = int32(1)
	if base.Ui32(int32(241)) < base.Ui32((v81+v143)&int32(255)) {
		v167 = v140
		v169 = v143
		goto L30
	} else {
		goto L46
	}
L40:
	;
	v163 = int32(8)
	v165 = v125
	goto L31
L41:
	;
	v163 = int32(4)
	v165 = v125
	goto L31
L42:
	;
	switch v81 + int32(-240) {
	case 0:
		goto L43
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
		goto L39
	case 14:
		v163 = int32(1)
		v165 = v125
		goto L31
	default:
		goto L44
	}
L43:
	;
	v163 = int32(3)
	v165 = v125
	goto L31
L44:
	;
	if v81 != int32(192) {
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v163 = int32(2)
	v165 = v125
	goto L31
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = int32(0)
	goto L33
L47:
	;
	if base.Ui32(v58) < base.Ui32(v175) {
		v365 = v173
		goto L16
	} else {
		goto L48
	}
L48:
	;
	goto L49
L49:
	;
	v180 = v44 - v78
	if base.Ui32(v180) < base.Ui32(v60) {
		v365 = v173
		goto L16
	} else {
		goto L50
	}
L50:
	;
	if base.Ui32(v180) <= base.Ui32(v58) {
		goto L17
	} else {
		goto L51
	}
L51:
	;
	v365 = v173
	goto L16
L52:
	;
	if base.Ui32(v58) < base.Ui32(v44) {
		v365 = v183
		goto L16
	} else {
		goto L53
	}
L53:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if base.Ui32(v187) < base.Ui32(int32(254)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v190 = int32(1)
	goto L56
L55:
	;
	v190 = int32(5)
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v190
	v192 = v44 + v190
	if base.Ui32(v192) < base.Ui32(v60) {
		v365 = v183
		goto L16
	} else {
		goto L57
	}
L57:
	;
	if base.Ui32(v58) < base.Ui32(v192) {
		v365 = v183
		goto L16
	} else {
		goto L58
	}
L58:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	v196 = int32(192)
	if base.Ui32(v195) < base.Ui32(v196) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v200 = v195 & v196
	goto L61
L60:
	;
	v200 = v195
	goto L61
L61:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+20)) = uint8(v200)
	v202 = int32(1)
	switch v200 + int32(-208) {
	case 0, 16:
		v226 = v202
		goto L62
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L63
	default:
		goto L64
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v226
	v228 = v192 + v226
	if base.Ui32(v228) < base.Ui32(v60) {
		v365 = v183
		goto L16
	} else {
		goto L73
	}
L63:
	;
	if base.Ui32((v200+int32(15))&int32(255)) < base.Ui32(int32(13)) {
		v226 = v202
		goto L62
	} else {
		goto L67
	}
L64:
	;
	switch v200 + int32(-240) {
	case 0, 14:
		v226 = v202
		goto L62
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
		goto L63
	default:
		goto L65
	}
L65:
	;
	if v200 == int32(192) {
		v226 = v202
		goto L62
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	if v200 == int32(0) {
		v226 = v202
		goto L62
	} else {
		goto L68
	}
L68:
	;
	if v200 == int32(128) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v226 = int32(5)
	goto L62
L70:
	;
	if v200 != int32(64) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = int32(255)
	v378 = int32(0)
	goto L15
L72:
	;
	v226 = int32(2)
	goto L62
L73:
	;
	if base.Ui32(v58) < base.Ui32(v228) {
		v365 = v183
		goto L16
	} else {
		goto L74
	}
L74:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if base.Ui32(v233) < base.Ui32(int32(254)) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v236 = int32(1)
	goto L77
L76:
	;
	v236 = int32(5)
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v236
	if base.Ui32(int32(253)) < base.Ui32(v233) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v244
	if base.Ui32(int32(191)) < base.Ui32(v195) {
		goto L84
	} else {
		goto L85
	}
L79:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v44+int32(1))))
	v244 = v243
	goto L78
L80:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v244 = v240
	goto L78
L81:
	;
	v340 = v338 + v236
	*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v340
	v343 = v44 + v340 + v339
	if base.Ui32(v343) < base.Ui32(v60) {
		v365 = v183
		goto L16
	} else {
		goto L99
	}
L82:
	;
	v334 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v334
	v338 = v296
	v339 = v334
	goto L81
L83:
	;
	v338 = int32(0)
	v339 = v332
	goto L81
L84:
	;
	v296 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v296
	switch v195 + int32(-208) {
	case 0:
		goto L93
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L90
	case 16:
		goto L92
	default:
		goto L94
	}
L85:
	;
	v248 = int32(2)
	switch int32(base.Ui32(v195)>>(uint(int32(6))%32)) ^ v248 {
	default:
		goto L87
	case 1:
		goto L86
	case 2:
		goto L89
	case 3:
		goto L88
	}
L86:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48)+8)) = int64(0)
	v332 = int32(0)
	goto L83
L87:
	;
	v272 = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v272
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v44+v236)+1))
	v277 = int32(24)
	v279 = int32(65280)
	v281 = int32(8)
	v291 = v276<<(uint(v277)%32) | v276&v279<<(uint(v281)%32) | (int32(base.Ui32(v276)>>(uint(v281)%32))&v279 | int32(base.Ui32(v276)>>(uint(v277)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v291
	v338 = v272
	v339 = v291
	goto L81
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = int32(2)
	v263 = v44 + v236
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+1)))
	v270 = v264&int32(63)<<(uint(int32(8))%32) | v269
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v270
	v338 = v248
	v339 = v270
	goto L81
L89:
	;
	v253 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v253
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v236))))
	v259 = v257 & int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v259
	v338 = v253
	v339 = v259
	goto L81
L90:
	;
	v318 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v318
	v321 = int32(1)
	if base.Ui32(int32(241)) < base.Ui32((v195+v321)&int32(255)) {
		v338 = v321
		v339 = v318
		goto L81
	} else {
		goto L98
	}
L91:
	;
	if v195 == int32(192) {
		goto L82
	} else {
		goto L97
	}
L92:
	;
	v313 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v313
	v338 = v296
	v339 = v313
	goto L81
L93:
	;
	v310 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v310
	v338 = v296
	v339 = v310
	goto L81
L94:
	;
	switch v195 + int32(-240) {
	case 0:
		goto L95
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
		goto L90
	case 14:
		goto L96
	default:
		goto L91
	}
L95:
	;
	v307 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v307
	v338 = v296
	v339 = v307
	goto L81
L96:
	;
	v303 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v303
	v338 = v303
	v339 = v303
	goto L81
L97:
	;
	goto L90
L98:
	;
	v328 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v328
	v332 = v328
	goto L83
L99:
	;
	if base.Ui32(v58) < base.Ui32(v343) {
		v365 = v183
		goto L16
	} else {
		goto L100
	}
L100:
	;
	goto L102
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v44
	goto L17
L102:
	;
	v348 = v44 - v244
	if base.Ui32(v348) < base.Ui32(v60) {
		v365 = v183
		goto L16
	} else {
		goto L103
	}
L103:
	;
	if base.Ui32(v58) < base.Ui32(v348) {
		v365 = v183
		goto L16
	} else {
		goto L104
	}
L104:
	;
	goto L101
L105:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v381 != v43 {
		goto L14
	} else {
		goto L106
	}
L106:
	;
	if l3 == int32(0) {
		goto L13
	} else {
		goto L107
	}
L107:
	;
	v385 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, v44, v29, l4)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	return int32(0)
L109:
	;
	if v385 != 0 {
		goto L13
	} else {
		goto L110
	}
L110:
	;
	goto L14
L111:
	;
	goto L12
L112:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v44 != l0+v404 {
		v431 = v399
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v420 = v391
	goto L8
L114:
	;
	v420 = v407
	goto L8
}
func F_zipmapNext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v7 != int32(255) {
		if l1 == int32(0) {
			v29 = v7
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = l0
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if base.Ui32(v15) <= base.Ui32(int32(253)) {
				v19 = v15
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1))
				v19 = v18
			}
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
			if base.Ui32(v19) < base.Ui32(int32(254)) {
				v25 = int32(1)
			} else {
				v25 = int32(5)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = l0 + v25
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			v29 = v28
		}
		v31 = v29 & int32(255)
		if base.Ui32(v31) <= base.Ui32(int32(253)) {
			v35 = v31
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1))
			v35 = v34
		}
		if base.Ui32(v35) < base.Ui32(int32(254)) {
			v40 = int32(1)
		} else {
			v40 = int32(5)
		}
		v42 = l0 + v40 + v35
		if l3 == int32(0) {
		} else {
			v46 = v42 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v46
			v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
			if base.Ui32(v48) <= base.Ui32(int32(253)) {
				v52 = v48
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
				v52 = v51
			}
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v52
			if base.Ui32(v52) < base.Ui32(int32(254)) {
				v58 = int32(1)
			} else {
				v58 = int32(5)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v46 + v58
		}
		v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
		if base.Ui32(v63) <= base.Ui32(int32(253)) {
			v67 = v63
		} else {
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v42)+1))
			v67 = v66
		}
		if base.Ui32(v67) < base.Ui32(int32(254)) {
			v73 = int32(1)
		} else {
			v73 = int32(5)
		}
		v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v73))))
		return v42 + v67 + v75 + v73 + int32(1)
	} else {
		return int32(0)
	}
}
func F_zipmapValidateIntegrity(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	v4 = int32(0)
	if base.Ui32(l1) < base.Ui32(int32(2)) {
		v103 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v103
L2:
	;
	v15 = l0 + l1 + int32(-1)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v16 != int32(255) {
		v103 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = int32(0)
	v28 = int32(1)
	goto L7
L5:
	;
	return int32(1)
L6:
	;
	if v27 != 0 {
		goto L29
	} else {
		goto L30
	}
L7:
	;
	v32 = l0 + v28
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v33 == int32(255) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v36 = int32(0)
	v40 = base.B2i32(v33 == int32(254))
	if v33 == int32(254) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v41 = int32(5)
	goto L12
L11:
	;
	v41 = int32(1)
	goto L12
L12:
	;
	v42 = v41 + v28
	if v42 < int32(2) {
		v103 = v36
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v45 = l0 + v42
	if base.Ui32(v15) < base.Ui32(v45) {
		v103 = v36
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v40 == int32(0) {
		v52 = v33
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if base.Ui64(base.I64_extend_i32_s(v15-v45)) < base.Ui64(base.I64_extend_i32_u(v52)) {
		v103 = v36
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v32)+1))
	if base.Ui32(v49) < base.Ui32(int32(254)) {
		v103 = v36
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v52 = v49
	goto L15
L18:
	;
	v57 = v52 + v42
	v58 = l0 + v57
	if base.Ui32(v15) < base.Ui32(v58) {
		v103 = v36
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v64 = base.B2i32(base.Ui32(int32(253)) < base.Ui32(v62))
	if base.Ui32(int32(253)) < base.Ui32(v62) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v65 = int32(5)
	goto L22
L21:
	;
	v65 = int32(1)
	goto L22
L22:
	;
	v66 = v65 + v57
	v67 = l0 + v66
	if base.Ui32(v15) < base.Ui32(v67) {
		v103 = v36
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v64 == int32(0) {
		v74 = v62
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v80 = v66 + int32(1)
	if base.Ui64(base.I64_extend_i32_s(v15-(l0+v80))) < base.Ui64(base.I64_extend_i32_u(v75)+base.I64_extend_i32_u(v74)) {
		v103 = v36
		goto L1
	} else {
		goto L27
	}
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v58)+1))
	if base.Ui32(v71) < base.Ui32(int32(254)) {
		v103 = v36
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v74 = v71
	goto L24
L27:
	;
	v88 = v74 + v80 + v75
	if base.Ui32(l0+v88) <= base.Ui32(v15) {
		v27 = v27 + int32(1)
		v28 = v88
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v103 = v36
	goto L1
L29:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v95 = v93 & int32(255)
	v103 = base.B2i32(v95 == int32(254)) | base.B2i32(v27 == v95)
	goto L1
L30:
	;
	return int32(0)
}
func F_zmadvise_dontneed_range(m *base.Module, l0 int32, l1 int32) {
	return
}
func F_zmscoreCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
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
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = F_lookupKeyRead(m, v9, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return
L2:
	;
	return
L3:
	;
	v15 = F_checkType(m, l0, v12, int32(3))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v15 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_addReplyArrayLen(m, l0, v17+int32(-2))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v22 < int32(3) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v29 = int32(2)
	goto L8
L8:
	;
	if v12 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L1
L10:
	;
	v50 = v29 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v50 < v51 {
		v29 = v50
		goto L8
	} else {
		goto L18
	}
L11:
	;
	v46 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
	F_addReplyDouble(m, l0, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L17
	}
L12:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L16
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v29<<(uint(int32(2))%32))))
	v37 = F_objectGetVal(m, v36)
	mBase = m.M
	v40 = F_zsetScore(m, v12, v37, v7+int32(8))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if v40 != int32(-1) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	goto L10
L17:
	;
	goto L10
L18:
	;
	goto L9
}
func F_zpopmaxCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 < int32(4) {
		v16 = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v16
		if v9 != int32(3) {
			v29 = v16
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v33 = int32(1)
			v35 = int32(0)
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
			F_genericZpopCommand(m, l0, v30+int32(4), v33, v33, v35, v29, base.B2i32(v29 != int32(-1))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(v38)), v35, v35)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
			v26 = F_getPositiveLongFromObjectOrReply(m, l0, v22, v7+int32(12), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				if v26 != 0 {
					m.G0 = v7 + int32(16)
					return
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v29 = v28
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v33 = int32(1)
					v35 = int32(0)
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
					F_genericZpopCommand(m, l0, v30+int32(4), v33, v33, v35, v29, base.B2i32(v29 != int32(-1))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(v38)), v35, v35)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[33]))
		F_addReplyErrorObject(m, l0, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_zrandmemberCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 < int32(3) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return
L2:
	;
	F_addReplyError(m, l0, int32(_a2375))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L43
	}
L3:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91<<(uint(int32(2))%32))+uint32(_consts[339])))
	v96 = F_lookupKeyReadOrReply(m, l0, v89, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L33
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v16 = F_getRangeLongFromObjectOrReply(m, l0, v12, int32(-2147483647), int32(2147483647), v6, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	if v16 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(4) < v18 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v76 == int32(4) {
		goto L28
	} else {
		goto L29
	}
L9:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L27
	}
L10:
	;
	if v18 == int32(4) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v29 = F_objectGetVal(m, v28)
	mBase = m.M
	v30 = int32(_a2452)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v33 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	F_zrandmemberWithCountCommand(m, l0, v23, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	goto L1
L14:
	;
	if v65-v67 == int32(0) {
		goto L8
	} else {
		goto L26
	}
L15:
	;
	v65 = F_tolower(m, v61)
	mBase = m.M
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v67 = F_tolower(m, v66)
	mBase = m.M
	goto L14
L16:
	;
	v35 = v29
	v36 = v30
	v37 = v33
	goto L19
L17:
	;
	v61 = int32(0)
	v62 = v30
	goto L15
L18:
	;
	v61 = v58 & int32(255)
	v62 = v57
	goto L15
L19:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v39 == int32(0) {
		v57 = v36
		v58 = v37
		goto L18
	} else {
		goto L21
	}
L20:
	;
	v57 = v51
	v58 = int32(0)
	goto L18
L21:
	;
	v43 = v37 & int32(255)
	if v43 == v39 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v50 = int32(1)
	v51 = v36 + v50
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	if v52 != 0 {
		v35 = v35 + v50
		v36 = v51
		v37 = v52
		goto L19
	} else {
		goto L25
	}
L23:
	;
	v45 = F_tolower(m, v43)
	mBase = m.M
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v47 = F_tolower(m, v46)
	mBase = m.M
	if v45 == v47 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v57 = v36
	v58 = v49
	goto L18
L25:
	;
	goto L20
L26:
	;
	goto L9
L27:
	;
	goto L1
L28:
	;
	if base.Ui32(v75+int32(-1073741824)) <= base.Ui32(int32(-2147483648)) {
		goto L2
	} else {
		goto L31
	}
L29:
	;
	F_zrandmemberWithCountCommand(m, l0, v75, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	goto L1
L31:
	;
	F_zrandmemberWithCountCommand(m, l0, v75, int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	goto L1
L33:
	;
	if v96 == int32(0) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v101 = F_checkType(m, l0, v96, int32(3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	if v101 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v103 = F_zsetLength(m, v96)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	F_zsetTypeRandomElement(m, v96, v103, v6, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v108 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	F_addReplyBulkLongLong(m, l0, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	F_addReplyBulkCBuffer(m, l0, v108, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	goto L1
L42:
	;
	goto L1
L43:
	;
	goto L1
}
func F_zrandmemberReplyWithListpack(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int64
	_ = v63
	var v68 int64
	_ = v68
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v89 float64
	_ = v89
	var v93 int64
	_ = v93
	var v95 float64
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = int32(0)
	goto L3
L3:
	;
	if l3 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v31 = v17 << (uint(int32(4)) % 32)
	v32 = l2 + v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v33 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if base.Ui32(v24) < base.Ui32(int32(3)) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	goto L5
L10:
	;
	if l3 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
	F_addReplyBulkLongLong(m, l0, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	F_addReplyBulkCBuffer(m, l0, v33, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	goto L10
L15:
	;
	v102 = v17 + int32(1)
	if v102 != l1 {
		v17 = v102
		goto L3
	} else {
		goto L29
	}
L16:
	;
	v44 = l3 + v31
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v45 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_addReplyDouble(m, l0, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L28
	}
L18:
	;
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v44)+8))
	v95 = base.F64_convert_i64_s(v93)
	goto L17
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v49 = int32(0)
	v53 = m.G0
	v55 = v53 - int32(32)
	m.G0 = v55
	v57 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v49
	v63 = *(*int64)(unsafe.Add(mBase, _consts[648]))
	*(*int64)(unsafe.Add(mBase, uint32(v55+int32(8)))) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v55)+24)) = int64(0)
	v68 = *(*int64)(unsafe.Add(mBase, _consts[649]))
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v68
	F_ffc_from_chars_double_options(m, v55+int32(16), v45, v45+v48, v55+int32(24), v55)
	mBase = m.M
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	if v76 == v49 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v95 = v89
	goto L17
L21:
	;
	goto L26
L22:
	;
	if v76 == int32(2) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v83 = int32(68)
	goto L25
L24:
	;
	v83 = int32(28)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v83
	goto L21
L26:
	;
	v89 = *(*float64)(unsafe.Add(mBase, uint32(v55)+24))
	m.G0 = v55 + int32(32)
	goto L20
L28:
	;
	goto L15
L29:
	;
	goto L4
}
func F_zrankGenericCommand(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v17 float64
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v154 int32
	_ = v154
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
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int64
	_ = v203
	var v208 int64
	_ = v208
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v229 float64
	_ = v229
	var v233 int64
	_ = v233
	var v236 float64
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 float64
	_ = v319
	var v320 float64
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v382 int32
	_ = v382
	var v394 float64
	_ = v394
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	v17 = float64(0)
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v22 < int32(5) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v20 + int32(32)
	return
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v22 == int32(4) {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	F_addReplyErrorArity(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
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
	if v382 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L7:
	;
	v382 = int32(-1)
	v394 = v17
	goto L6
L8:
	;
	F__serverPanic_1(m, int32(_a2436), int32(1704), int32(_a576), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L94
	}
L9:
	;
	F__serverAssert(m, int32(_a2458), int32(_a2436), int32(1697))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L93
	}
L10:
	;
	F__serverAssert(m, int32(_a575), int32(_a2436), int32(857))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L4
	} else {
		goto L92
	}
L11:
	;
	F__serverAssert(m, int32(_a575), int32(_a2436), int32(1670))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L4
	} else {
		goto L91
	}
L12:
	;
	F__serverAssert(m, int32(_a574), int32(_a2436), int32(1668))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L90
	}
L13:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L4
	} else {
		goto L89
	}
L14:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79+v80<<(uint(int32(2))%32))))
	v85 = F_lookupKeyReadOrReply(m, l0, v29, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L30
	}
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v36 = F_objectGetVal(m, v35)
	mBase = m.M
	v37 = int32(_a2459)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v40 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v79 = int32(_a2460)
	goto L14
L17:
	;
	if v72-v74 != 0 {
		goto L13
	} else {
		goto L29
	}
L18:
	;
	v72 = F_tolower(m, v68)
	mBase = m.M
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v74 = F_tolower(m, v73)
	mBase = m.M
	goto L17
L19:
	;
	v42 = v36
	v43 = v37
	v44 = v40
	goto L22
L20:
	;
	v68 = int32(0)
	v69 = v37
	goto L18
L21:
	;
	v68 = v65 & int32(255)
	v69 = v64
	goto L18
L22:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v46 == int32(0) {
		v64 = v43
		v65 = v44
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v64 = v58
	v65 = int32(0)
	goto L21
L24:
	;
	v50 = v44 & int32(255)
	if v50 == v46 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v57 = int32(1)
	v58 = v43 + v57
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v59 != 0 {
		v42 = v42 + v57
		v43 = v58
		v44 = v59
		goto L22
	} else {
		goto L28
	}
L26:
	;
	v52 = F_tolower(m, v50)
	mBase = m.M
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v54 = F_tolower(m, v53)
	mBase = m.M
	if v52 == v54 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v64 = v43
	v65 = v56
	goto L21
L28:
	;
	goto L23
L29:
	;
	v79 = int32(_a2461)
	goto L14
L30:
	;
	if v85 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v90 = F_checkType(m, l0, v85, int32(3))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	if v90 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	switch int32(base.Ui32(v92)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		goto L34
	default:
		goto L35
	}
L34:
	;
	v103 = F_objectGetVal(m, v28)
	mBase = m.M
	v104 = F_zsetLength(m, v85)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	F__serverAssertWithInfo(m, l0, v28, int32(_a2462), int32(_a2436), int32(3819))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	switch int32(base.Ui32(v106)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L38
	default:
		goto L8
	case 4:
		goto L39
	}
L38:
	;
	v242 = F_objectGetVal(m, v85)
	mBase = m.M
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v246 = F_hashtableFind(m, v243, v103, v20+int32(16))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L73
	}
L39:
	;
	v113 = F_objectGetVal(m, v85)
	mBase = m.M
	v115 = F_lpSeek(m, v113, int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v115
	if v115 == int32(0) {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	v120 = F_lpNext(m, v113, v115)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v120
	if v120 == int32(0) {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	v141 = int32(1)
	v142 = v115
	goto L44
L44:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103+int32(-1)))))
	switch v154 & int32(7) {
	case 0:
		goto L51
	case 1:
		goto L50
	case 2:
		goto L49
	case 3:
		goto L48
	case 4:
		goto L47
	default:
		v163 = int32(0)
		goto L46
	}
L45:
	;
	if v22 != int32(4) {
		v236 = v17
		goto L57
	} else {
		goto L58
	}
L46:
	;
	v164 = F_lpCompare(m, v142, v103, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L53
	}
L47:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v103+int32(-17))))
	v163 = v162
	goto L46
L48:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v103+int32(-9))))
	v163 = v161
	goto L46
L49:
	;
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103+int32(-5)))))
	v163 = v160
	goto L46
L50:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103+int32(-3)))))
	v163 = v159
	goto L46
L51:
	;
	v163 = int32(base.Ui32(v154) >> (uint(int32(3)) % 32))
	goto L46
L52:
	;
	goto L45
L53:
	;
	if v164 != 0 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	F_zzlNext(m, v113, v20+int32(12), v20+int32(8))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v174 != 0 {
		v141 = v141 + int32(1)
		v142 = v174
		goto L44
	} else {
		goto L56
	}
L56:
	;
	goto L7
L57:
	;
	if l1 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L58:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v177 == int32(0) {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	v184 = F_lpGetValue(m, v177, v20+int32(28), v20+int32(16))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L61
	}
L60:
	;
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v20)+16))
	v236 = base.F64_convert_i64_s(v233)
	goto L57
L61:
	;
	if v184 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v189 = int32(0)
	v193 = m.G0
	v195 = v193 - int32(32)
	m.G0 = v195
	v197 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v189
	v203 = *(*int64)(unsafe.Add(mBase, _consts[648]))
	*(*int64)(unsafe.Add(mBase, uint32(v195+int32(8)))) = v203
	*(*int64)(unsafe.Add(mBase, uint32(v195)+24)) = int64(0)
	v208 = *(*int64)(unsafe.Add(mBase, _consts[649]))
	*(*int64)(unsafe.Add(mBase, uint32(v195))) = v208
	F_ffc_from_chars_double_options(m, v195+int32(16), v184, v184+v188, v195+int32(24), v195)
	mBase = m.M
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v195)+20))
	if v216 == v189 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v236 = v229
	goto L57
L64:
	;
	goto L69
L65:
	;
	if v216 == int32(2) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v223 = int32(68)
	goto L68
L67:
	;
	v223 = int32(28)
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v223
	goto L64
L69:
	;
	v229 = *(*float64)(unsafe.Add(mBase, uint32(v195)+24))
	m.G0 = v195 + int32(32)
	goto L63
L71:
	;
	v382 = v141 + int32(-1)
	v394 = v236
	goto L6
L72:
	;
	v382 = v104 - v141
	v394 = v236
	goto L6
L73:
	;
	if v246 == int32(0) {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	v251 = int32(0)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v252 == v251 {
		v303 = v251
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v314 == v303 {
		goto L9
	} else {
		goto L83
	}
L76:
	;
	v259 = v252
	v261 = v251
	goto L77
L77:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v259)+16))
	v274 = v272 + int32(-1)
	if v274 < int32(1) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v303 = v296
	goto L75
L79:
	;
	v296 = v294 + v261
	if v293 != 0 {
		v259 = v293
		v261 = v296
		goto L77
	} else {
		goto L82
	}
L80:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v259+v274<<(uint(int32(3))%32)+int32(12))))
	v293 = v290
	v294 = base.B2i32(v290 != int32(0))
	goto L79
L81:
	;
	v280 = v274 << (uint(int32(3)) % 32)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v259+v280)+12))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v259+int32(16)+v280)))
	v293 = v282
	v294 = v284
	goto L79
L82:
	;
	goto L78
L83:
	;
	v316 = v314 - v303
	if v22 != int32(4) {
		v320 = v17
		goto L84
	} else {
		goto L85
	}
L84:
	;
	if l1 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v319 = *(*float64)(unsafe.Add(mBase, uint32(v252)))
	v320 = v319
	goto L84
L86:
	;
	v324 = v104 - v316
	goto L88
L87:
	;
	v324 = v316 + int32(-1)
	goto L88
L88:
	;
	v382 = v324
	v394 = v320
	goto L6
L89:
	;
	goto L1
L90:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v382))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L4
	} else {
		goto L106
	}
L96:
	;
	if v22 != int32(4) {
		goto L102
	} else {
		goto L103
	}
L97:
	;
	if v22 != int32(4) {
		goto L95
	} else {
		goto L98
	}
L98:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v382))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	F_addReplyDouble(m, l0, v394)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	goto L1
L102:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L4
	} else {
		goto L105
	}
L103:
	;
	F_addReplyNullArray(m, l0)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	goto L1
L105:
	;
	goto L1
L106:
	;
	goto L1
}
func F_zremCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v11 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v12 = F_lookupKeyWriteOrReply(m, l0, v9, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	return
L3:
	;
	if v12 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v17 = F_checkType(m, l0, v12, int32(3))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v17 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v19&int32(240) != int32(112) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v30 = int32(2)
	v31 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v33 <= v30 {
		v70 = v31
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v24 = F_objectGetVal(m, v12)
	mBase = m.M
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+28)))
	v28 = v26 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+28)) = uint16(v28)
	goto L9
L9:
	;
	goto L7
L10:
	;
	if v86 != 0 {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v72&int32(240) != int32(112) {
		v85 = v31
		v86 = v70
		goto L10
	} else {
		goto L21
	}
L12:
	;
	v40 = v30
	v42 = int32(0)
	goto L14
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v62 = F_dbDelete(m, v61, v9)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L2
	} else {
		goto L20
	}
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v40<<(uint(int32(2))%32))))
	v49 = F_objectGetVal(m, v48)
	mBase = m.M
	v50 = F_zsetDel(m, v12, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v52 = v50 + v42
	v53 = F_zsetLength(m, v12)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	if v53 == int32(0) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v58 = v40 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v59 <= v58 {
		v70 = v52
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v40 = v58
	v42 = v52
	goto L14
L20:
	;
	v85 = int32(1)
	v86 = v52
	goto L10
L21:
	;
	v77 = F_objectGetVal(m, v12)
	mBase = m.M
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	F_hashtableResumeAutoShrink(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v85 = v31
	v86 = v70
	goto L10
L23:
	;
	F_addReplyLongLong(m, l0, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L2
	} else {
		goto L31
	}
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	F_notifyKeyspaceEvent(m, int32(128), int32(_a2444), v9, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L26
	}
L25:
	;
	v112 = int64(0)
	goto L23
L26:
	;
	if v85 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v103, v9)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a308), v9, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v106 = int32(_a20)
	v108 = *(*int64)(unsafe.Add(mBase, _consts[180]))
	v109 = base.I64_extend_i32_s(v86)
	*(*int64)(unsafe.Add(mBase, _consts[180])) = v108 + v109
	v112 = v109
	goto L23
L31:
	;
	goto L1
}
func F_zremrangeGenericCommand(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
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
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
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
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v234 int64
	_ = v234
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	v11 = m.G0
	v13 = v11 - int32(64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = int32(0)
	if l1 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(64)
	return
L2:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v66 = F_lookupKeyWriteOrReply(m, l0, v16, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L22
	}
L3:
	;
	switch l1 + int32(-2) {
	case 0:
		goto L12
	case 1:
		goto L11
	default:
		goto L10
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v25 = F_getLongFromObjectOrReply(m, l0, v21, v13+int32(12), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	if v25 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v32 = F_getLongFromObjectOrReply(m, l0, v28, v13+int32(8), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v32 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v63 = int32(_a2447)
	goto L2
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	F__serverPanic_1(m, int32(_a2436), int32(2012), int32(_a2448), v13)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L20
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v50 = F_zsetParseLexRange(m, v46, v47, v13+int32(16))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L17
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v41 = F_zslParseRange(m, v37, v38, v13+int32(32))
	mBase = m.M
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_addReplyError(m, l0, int32(_a2445))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v63 = int32(_a2446)
	goto L2
L15:
	;
	goto L1
L16:
	;
	F_addReplyError(m, l0, int32(_a2449))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	if v50 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v63 = int32(_a2450)
	goto L2
L19:
	;
	goto L1
L20:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	if l1 != int32(3) {
		goto L1
	} else {
		goto L78
	}
L22:
	;
	if v66 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v71 = F_checkType(m, l0, v66, int32(3))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	if v71 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	if l1 != int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	switch int32(base.Ui32(v108)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L44
	default:
		goto L43
	case 4:
		goto L45
	}
L27:
	;
	v75 = F_zsetLength(m, v66)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if int32(-1) < v77 {
		v82 = v77
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if int32(-1) < v83 {
		v88 = v83
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v80 = v77 + v75
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v80
	v82 = v80
	goto L29
L31:
	;
	if int32(-1) < v82 {
		v94 = v82
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v86 = v83 + v75
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v86
	v88 = v86
	goto L31
L33:
	;
	if v88 < v94 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v91 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v91
	v94 = v91
	goto L33
L35:
	;
	if base.Ui32(v88) < base.Ui32(v75) {
		goto L26
	} else {
		goto L40
	}
L36:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	F_addReply(m, l0, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L39
	}
L37:
	;
	if v94 < v75 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L1
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v75 + int32(-1)
	goto L26
L41:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	if v212 != 0 {
		goto L70
	} else {
		goto L71
	}
L42:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v205 = F_dbDelete(m, v204, v16)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L5
	} else {
		goto L68
	}
L43:
	;
	F__serverPanic_1(m, int32(_a2436), int32(2061), int32(_a576), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L5
	} else {
		goto L67
	}
L44:
	;
	v153 = F_objectGetVal(m, v66)
	mBase = m.M
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+28)))
	v157 = v155 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v154)+28)) = uint16(v157)
	goto L56
L45:
	;
	v115 = F_objectGetVal(m, v66)
	mBase = m.M
	switch l1 + int32(-1) {
	default:
		goto L49
	case 1:
		goto L48
	case 2:
		goto L47
	}
L46:
	;
	F_objectSetVal(m, v66, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L5
	} else {
		goto L53
	}
L47:
	;
	v140 = F_zzlDeleteRangeByLex(m, v115, v13+int32(16), v13+int32(60))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L52
	}
L48:
	;
	v134 = F_zzlDeleteRangeByScore(m, v115, v13+int32(32), v13+int32(60))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L51
	}
L49:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v121 = int32(1)
	v122 = v118 - v119 + v121
	*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v122
	v128 = F_lpDeleteRange(m, v115, v119<<(uint(v121)%32), v122<<(uint(v121)%32))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	v143 = v128
	goto L46
L51:
	;
	v143 = v134
	goto L46
L52:
	;
	v143 = v140
	goto L46
L53:
	;
	v147 = F_objectGetVal(m, v66)
	mBase = m.M
	v148 = F_lpLength(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	if base.Ui32(v148) < base.Ui32(int32(2)) {
		goto L42
	} else {
		goto L55
	}
L55:
	;
	v208 = int32(1)
	goto L41
L56:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	switch l1 + int32(-1) {
	default:
		goto L60
	case 1:
		goto L59
	case 2:
		goto L58
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	F_hashtableResumeAutoShrink(m, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L5
	} else {
		goto L64
	}
L58:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v179 = F_zslDeleteRangeByLex(m, v159, v13+int32(16), v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L63
	}
L59:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v174 = F_zslDeleteRangeByScore(m, v159, v13+int32(32), v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L62
	}
L60:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v163 = int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v169 = F_zslDeleteRangeByRank(m, v159, v162+v163, v165+v163, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	v181 = v169
	goto L57
L62:
	;
	v181 = v174
	goto L57
L63:
	;
	v181 = v179
	goto L57
L64:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
	goto L65
L65:
	;
	if v187+v188 == int32(0) {
		goto L42
	} else {
		goto L66
	}
L66:
	;
	v208 = int32(1)
	goto L41
L67:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	v208 = int32(0)
	goto L41
L69:
	;
	F_addReplyLongLong(m, l0, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L77
	}
L70:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v214, v16)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L72
	}
L71:
	;
	v234 = int64(0)
	goto L69
L72:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+28))
	F_notifyKeyspaceEvent(m, int32(128), v63, v16, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	if v208 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v228 = int32(_a20)
	v230 = *(*int64)(unsafe.Add(mBase, _consts[180]))
	v231 = base.I64_extend_i32_u(v212)
	*(*int64)(unsafe.Add(mBase, _consts[180])) = v230 + v231
	v234 = v231
	goto L69
L75:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a308), v16, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	goto L21
L78:
	;
	v244 = int32(_a388)
	v245 = *(*int32)(unsafe.Add(mBase, _consts[420]))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v248 = *(*int32)(unsafe.Add(mBase, _consts[419]))
	if v246 == v248 {
		v257 = v245
		v258 = v248
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v259 == v258 {
		goto L1
	} else {
		goto L83
	}
L80:
	;
	if v246 == v245 {
		v257 = v245
		v258 = v248
		goto L79
	} else {
		goto L81
	}
L81:
	;
	F_sdsfree(m, v246)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	v253 = int32(_a388)
	v254 = *(*int32)(unsafe.Add(mBase, _consts[420]))
	v256 = *(*int32)(unsafe.Add(mBase, _consts[419]))
	v257 = v254
	v258 = v256
	goto L79
L83:
	;
	if v259 == v257 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_sdsfree(m, v259)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	goto L1
}
func F_zremrangebyscoreCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_zremrangeGenericCommand(m, l0, int32(2))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_zrevrangebylexCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int64
	_ = v13
	var v35 int32
	_ = v35
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5+int32(24)))) = v2
	v13 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(16)))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(8)))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v5)+40)) = int32(1101)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = int32(1102)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = int32(1103)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = int32(1104)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l0
	F_zrangeGenericCommand(m, v5, int32(1), v2, int32(3), int32(2))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return
	} else {
		m.G0 = v5 + int32(48)
		return
	}
}
func F_zrevrankCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_zrankGenericCommand(m, l0, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_zstrdup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	if l0&int32(3) == int32(0) {
		v24 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v59 = v57 + int32(1)
	v60 = F_valkey_malloc(m, v59)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v57 = v49 - l0
	goto L1
L3:
	;
	v28 = v24
	goto L11
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v13 = l0
	goto L7
L6:
	;
	v57 = l0 - l0
	goto L1
L7:
	;
	v17 = v13 + int32(1)
	if v17&int32(3) == int32(0) {
		v24 = v17
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v22 != 0 {
		v13 = v17
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v49 = v17
	goto L2
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v37 = int32(-2139062144)
	if (int32(16843008)-v34|v34)&v37 == v37 {
		v28 = v28 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v43 = v28
	goto L14
L13:
	;
	goto L12
L14:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v47 != 0 {
		v43 = v43 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v49 = v43
	goto L2
L16:
	;
	goto L15
L17:
	;
	return int32(0)
L18:
	;
	if v59 == int32(0) {
		v67 = v60
		goto L20
	} else {
		goto L21
	}
L19:
	;
	return v67
L20:
	;
	goto L19
L21:
	;
	v66 = F__emscripten_memcpy_bulkmem(m, v60, l0, v59)
	mBase = m.M
	v67 = v66
	goto L20
}
func F_ztrycalloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	v2 = int32(0)
	if base.Ui32(int32(2147483646)) < base.Ui32(l0) {
		v78 = v2
	} else {
		v8 = int32(1)
		if l0 != 0 {
			v10 = l0
		} else {
			v10 = int32(4)
		}
		v12 = v10 + int32(8)
		v18 = base.I64_extend_i32_u(v8) * base.I64_extend_i32_u(v12)
		v19 = base.I32_wrap_i64(v18)
		if base.Ui32(v12|v8) < base.Ui32(int32(65536)) {
			v30 = v19
		} else {
			if base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) != int32(0) {
				v29 = int32(-1)
			} else {
				v29 = v19
			}
			v30 = v29
		}
		v32 = F_emscripten_builtin_malloc(m, v30)
		mBase = m.M
		if v32 == int32(0) {
		} else {
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-4)))))
			if v37&int32(3) == int32(0) {
			} else {
				v43 = F___memset(m, v32, int32(0), v30)
				mBase = m.M
			}
		}
		if v32 == int32(0) {
			v78 = v2
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v32))) = v10
			v48 = *(*int32)(unsafe.Add(mBase, _consts[525]))
			if v48 != int32(-1) {
				v59 = v48
			} else {
				v51 = int32(0)
				v53 = *(*int32)(unsafe.Add(mBase, _consts[315]))
				*(*int32)(unsafe.Add(mBase, _consts[525])) = v53
				*(*int32)(unsafe.Add(mBase, _consts[315])) = v53 + int32(1)
				v59 = v53
			}
			if v59 < int32(260) {
				v68 = v59 << (uint(int32(2)) % 32)
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_consts[320])))
				*(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_consts[320]))) = v71 + v12
			} else {
				v62 = int32(0)
				v64 = *(*int32)(unsafe.Add(mBase, _consts[316]))
				*(*int32)(unsafe.Add(mBase, _consts[316])) = v64 + v12
			}
			v78 = v32 + int32(8)
		}
	}
	return v78
}
func F_ztrycalloc_usable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	v3 = int32(0)
	if base.Ui32(int32(2147483646)) < base.Ui32(l0) {
		v80 = v3
		v81 = v3
	} else {
		v10 = int32(1)
		if l0 != 0 {
			v12 = l0
		} else {
			v12 = int32(4)
		}
		v14 = v12 + int32(8)
		v20 = base.I64_extend_i32_u(v10) * base.I64_extend_i32_u(v14)
		v21 = base.I32_wrap_i64(v20)
		if base.Ui32(v14|v10) < base.Ui32(int32(65536)) {
			v32 = v21
		} else {
			if base.I32_wrap_i64(int64(base.Ui64(v20)>>(uint(int64(32))%64))) != int32(0) {
				v31 = int32(-1)
			} else {
				v31 = v21
			}
			v32 = v31
		}
		v34 = F_emscripten_builtin_malloc(m, v32)
		mBase = m.M
		if v34 == int32(0) {
		} else {
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+int32(-4)))))
			if v39&int32(3) == int32(0) {
			} else {
				v45 = F___memset(m, v34, int32(0), v32)
				mBase = m.M
			}
		}
		if v34 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v34))) = v12
			v50 = *(*int32)(unsafe.Add(mBase, _consts[525]))
			if v50 != int32(-1) {
				v61 = v50
			} else {
				v53 = int32(0)
				v55 = *(*int32)(unsafe.Add(mBase, _consts[315]))
				*(*int32)(unsafe.Add(mBase, _consts[525])) = v55
				*(*int32)(unsafe.Add(mBase, _consts[315])) = v55 + int32(1)
				v61 = v55
			}
			if v61 < int32(260) {
				v70 = v61 << (uint(int32(2)) % 32)
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[320])))
				*(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[320]))) = v73 + v14
			} else {
				v64 = int32(0)
				v66 = *(*int32)(unsafe.Add(mBase, _consts[316]))
				*(*int32)(unsafe.Add(mBase, _consts[316])) = v66 + v14
			}
			v80 = v12
			v81 = v34 + int32(8)
		} else {
			v46 = int32(0)
			v80 = v46
			v81 = v46
		}
	}
	if l1 == int32(0) {
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v80
	}
	return v81
}
func F_zuiLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		switch v6 + int32(-2) {
		case 0:
			v38 = F_setTypeSize(m, v3)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				return v38
			}
		case 1:
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			switch v9 + int32(-7) {
			case 0:
				v20 = F_objectGetVal(m, v3)
				mBase = m.M
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				return v22
			default:
				F__serverPanic_1(m, int32(_a2436), int32(2237), int32(_a576), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			case 4:
				v12 = F_objectGetVal(m, v3)
				mBase = m.M
				v13 = F_lpLength(m, v12)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					return int32(base.Ui32(v13) >> (uint(int32(1)) % 32))
				}
			}
		default:
			F__serverPanic_1(m, int32(_a2436), int32(2240), int32(_a2451), int32(0))
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
		}
	} else {
		return int32(0)
	}
}
