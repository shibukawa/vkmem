package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_json_append_data(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 float64
	_ = v292
	var v293 int32
	_ = v293
	var v320 int32
	_ = v320
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v486 int32
	_ = v486
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int64
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v556 int32
	_ = v556
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v614 int32
	_ = v614
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int64
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v684 int32
	_ = v684
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v744 int32
	_ = v744
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1031 int32
	_ = v1031
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
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
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	v12 = m.G0
	v14 = v12 - int32(64)
	m.G0 = v14
	goto L11
L1:
	;
	m.G0 = v14 + int32(64)
	return
L2:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1328))
	if v1195 != 0 {
		goto L312
	} else {
		goto L313
	}
L3:
	;
	goto L291
L4:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if base.Ui32(v1091-v1092) < base.Ui32(int32(-4)) {
		v1101 = v1091
		goto L286
	} else {
		goto L287
	}
L5:
	;
	v190 = l2 + int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1316))
	if v191 <= l2 {
		goto L55
	} else {
		goto L56
	}
L6:
	;
	goto L29
L7:
	;
	F_json_append_number(m, l0, l1, l3, int32(-1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L24
	} else {
		goto L26
	}
L8:
	;
	F_json_append_string(m, l0, l3, int32(-1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L24
	} else {
		goto L25
	}
L9:
	;
	switch v73 {
	case 0:
		goto L4
	case 1:
		goto L6
	case 2:
		goto L3
	case 3:
		goto L7
	case 4:
		goto L8
	case 5:
		goto L5
	default:
		goto L2
	}
L10:
	;
	v67 = m.G398
	if v33 != v67 {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	goto L15
L15:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v33 = v30 + int32(-16)
	goto L10
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v73 = v70
	goto L9
L23:
	;
	v73 = int32(-1)
	goto L9
L24:
	;
	return
L25:
	;
	goto L1
L26:
	;
	goto L1
L27:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v145 = v143 - v144
	if v142 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L28:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	switch v135 {
	case 0:
		v140 = v135
		goto L43
	case 1:
		goto L45
	default:
		goto L44
	}
L29:
	;
	goto L35
L35:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v99 = v96 + int32(-16)
	goto L28
L43:
	;
	v142 = v140
	goto L27
L44:
	;
	v140 = int32(1)
	goto L43
L45:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v142 = base.B2i32(v136 != int32(0))
	goto L27
L46:
	;
	if base.Ui32(v145) < base.Ui32(int32(-5)) {
		v171 = v143
		goto L51
	} else {
		goto L52
	}
L47:
	;
	if base.Ui32(v145) < base.Ui32(int32(-4)) {
		v155 = v143
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v156+v155))) = int32(1702195828)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v160 + int32(4)
	goto L1
L49:
	;
	F_strbuf_resize(m, l3, v143+int32(4))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L24
	} else {
		goto L50
	}
L50:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v155 = v154
	goto L48
L51:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v173 = v172 + v171
	v174 = m.G3
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v174)+uint32(_consts[1019])))
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v177
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+uint32(_consts[1020]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v173+int32(4)))) = uint8(v183)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v185 + int32(5)
	goto L1
L52:
	;
	F_strbuf_resize(m, l3, v143+int32(5))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L24
	} else {
		goto L53
	}
L53:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v171 = v170
	goto L51
L54:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v208 + int32(16)
	goto L63
L55:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1328))
	if v196 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v194 = F_lua_checkstack(m, l0, int32(3))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L24
	} else {
		goto L57
	}
L57:
	;
	if v194 != 0 {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v190
	v200 = m.G3
	v205 = F_luaL_error(m, l0, v200+int32(_a2127), v14+int32(48))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L24
	} else {
		goto L62
	}
L60:
	;
	F_strbuf_free(m, l3)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L24
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	goto L54
L63:
	;
	v214 = int32(0)
	v216 = F_lua_next(m, l0, int32(-2))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L24
	} else {
		goto L68
	}
L64:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v766 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v768 = v766 + int32(1)
	if v765 != v768 {
		v775 = v766
		v776 = v768
		goto L208
	} else {
		goto L209
	}
L65:
	;
	goto L202
L66:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1308))
	if v357 < int32(1) {
		goto L108
	} else {
		goto L109
	}
L67:
	;
	v222 = int32(0)
	v226 = v214
	goto L70
L68:
	;
	if v216 != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v348 = int32(0)
	v352 = v214
	goto L66
L70:
	;
	goto L74
L71:
	;
	v348 = v340
	v352 = v342
	goto L66
L72:
	;
	if v288 != int32(3) {
		goto L65
	} else {
		goto L87
	}
L73:
	;
	v282 = m.G398
	if v248 != v282 {
		goto L85
	} else {
		goto L86
	}
L74:
	;
	goto L78
L78:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v248 = v245 + int32(-32)
	goto L73
L85:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v248)+8))
	v288 = v285
	goto L72
L86:
	;
	v288 = int32(-1)
	goto L72
L87:
	;
	v292 = F_lua_tonumber(m, l0, int32(-2))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L24
	} else {
		goto L88
	}
L88:
	;
	if base.F64_eq(v292, float64(0)) != 0 {
		goto L65
	} else {
		goto L89
	}
L89:
	;
	if base.F64_ge(v292, float64(1)) == int32(0) {
		goto L65
	} else {
		goto L90
	}
L90:
	;
	if base.F64_ne(base.F64_floor(v292), v292) != 0 {
		goto L65
	} else {
		goto L91
	}
L91:
	;
	goto L94
L92:
	;
	if base.F64_lt(base.F64_abs(v292), float64(2.147483648e+09)) == int32(0) {
		goto L101
	} else {
		goto L102
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v320 + int32(-16)
	goto L92
L94:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L93
L100:
	;
	if base.F64_gt(v292, base.F64_convert_i32_s(v222)) != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v339 = int32(-2147483648)
	goto L100
L102:
	;
	v337 = base.I32_trunc_f64_s(v292)
	v339 = v337
	goto L100
L103:
	;
	v340 = v339
	goto L105
L104:
	;
	v340 = v222
	goto L105
L105:
	;
	v342 = v226 + int32(1)
	v344 = F_lua_next(m, l0, int32(-2))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L24
	} else {
		goto L106
	}
L106:
	;
	if v344 != 0 {
		v222 = v340
		v226 = v342
		goto L70
	} else {
		goto L107
	}
L107:
	;
	goto L71
L108:
	;
	if v348 < int32(1) {
		goto L64
	} else {
		goto L135
	}
L109:
	;
	if v348 <= v357*v352 {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1312))
	if v348 <= v362 {
		goto L108
	} else {
		goto L111
	}
L111:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1304))
	if v364 != 0 {
		goto L64
	} else {
		goto L112
	}
L112:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1328))
	if v365 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v368 = m.G3
	goto L118
L114:
	;
	F_strbuf_free(m, l3)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L24
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v428 = m.G3
	if v426 != int32(-1) {
		goto L132
	} else {
		goto L133
	}
L117:
	;
	v420 = m.G398
	if v386 != v420 {
		goto L129
	} else {
		goto L130
	}
L118:
	;
	goto L122
L122:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v386 = v383 + int32(-16)
	goto L117
L129:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v386)+8))
	v426 = v423
	goto L116
L130:
	;
	v426 = int32(-1)
	goto L116
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v368 + int32(_a2128)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v438
	v447 = F_luaL_error(m, l0, v368+int32(_a2129), v14+int32(32))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L24
	} else {
		goto L134
	}
L132:
	;
	v433 = m.G399
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v433+v426<<(uint(int32(2))%32))))
	v438 = v437
	goto L131
L133:
	;
	v438 = v428 + int32(_a2018)
	goto L131
L134:
	;
	goto L64
L135:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v454 = v452 + int32(1)
	if v451 != v454 {
		v461 = v452
		v462 = v454
		goto L136
	} else {
		goto L137
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v462
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v466 = int32(91)
	*(*uint8)(unsafe.Add(mBase, uint32(v464+v461))) = uint8(v466)
	goto L141
L137:
	;
	F_strbuf_resize(m, l3, v451)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L24
	} else {
		goto L138
	}
L138:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v461 = v458
	v462 = v458 + int32(1)
	goto L136
L139:
	;
	F_json_append_data(m, l0, l1, v190, l3)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L24
	} else {
		goto L155
	}
L140:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v486+int32(-16))))
	v526 = F_luaH_getnum(m, v525, int32(1))
	mBase = m.M
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v528 = *(*int64)(unsafe.Add(mBase, uint32(v526)))
	*(*int64)(unsafe.Add(mBase, uint32(v527))) = v528
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v526)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v527)+8)) = v530
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v532 + int32(16)
	goto L139
L141:
	;
	goto L147
L147:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L140
L155:
	;
	goto L158
L156:
	;
	if v348 == int32(1) {
		goto L164
	} else {
		goto L165
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v556 + int32(-16)
	goto L156
L158:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L157
L164:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v711 = v709 + int32(1)
	if v708 != v711 {
		v718 = v709
		v719 = v711
		goto L197
	} else {
		goto L198
	}
L165:
	;
	v575 = int32(2)
	goto L166
L166:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v583 = v581 + int32(1)
	if v580 != v583 {
		v590 = v581
		v591 = v583
		goto L168
	} else {
		goto L169
	}
L167:
	;
	goto L164
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v591
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v595 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v593+v590))) = uint8(v595)
	goto L173
L169:
	;
	F_strbuf_resize(m, l3, v580)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L24
	} else {
		goto L170
	}
L170:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v590 = v587
	v591 = v587 + int32(1)
	goto L168
L171:
	;
	F_json_append_data(m, l0, l1, v190, l3)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L24
	} else {
		goto L187
	}
L172:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v614+int32(-16))))
	v654 = F_luaH_getnum(m, v653, v575)
	mBase = m.M
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v656 = *(*int64)(unsafe.Add(mBase, uint32(v654)))
	*(*int64)(unsafe.Add(mBase, uint32(v655))) = v656
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v654)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v655)+8)) = v658
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v660 + int32(16)
	goto L171
L173:
	;
	goto L179
L179:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L172
L187:
	;
	goto L190
L188:
	;
	if v348 != v575 {
		v575 = v575 + int32(1)
		goto L166
	} else {
		goto L196
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v684 + int32(-16)
	goto L188
L190:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L189
L196:
	;
	goto L167
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v719
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v723 = int32(93)
	*(*uint8)(unsafe.Add(mBase, uint32(v721+v718))) = uint8(v723)
	goto L1
L198:
	;
	F_strbuf_resize(m, l3, v708)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L24
	} else {
		goto L199
	}
L199:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v718 = v715
	v719 = v715 + int32(1)
	goto L197
L200:
	;
	goto L64
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v744 + int32(-32)
	goto L200
L202:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L201
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v776
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v780 = int32(123)
	*(*uint8)(unsafe.Add(mBase, uint32(v778+v775))) = uint8(v780)
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v783)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v783 + int32(16)
	goto L211
L209:
	;
	F_strbuf_resize(m, l3, v765)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L24
	} else {
		goto L210
	}
L210:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v775 = v772
	v776 = v772 + int32(1)
	goto L208
L211:
	;
	v790 = F_lua_next(m, l0, int32(-2))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L24
	} else {
		goto L213
	}
L212:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v1077 = v1075 + int32(1)
	if v1074 != v1077 {
		v1084 = v1075
		v1085 = v1077
		goto L283
	} else {
		goto L284
	}
L213:
	;
	if v790 == int32(0) {
		goto L212
	} else {
		goto L214
	}
L214:
	;
	goto L215
L215:
	;
	goto L223
L217:
	;
	F_json_append_data(m, l0, l1, v190, l3)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L24
	} else {
		goto L269
	}
L218:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1328))
	if v924 != 0 {
		goto L247
	} else {
		goto L248
	}
L219:
	;
	F_json_append_string(m, l0, l3, int32(-2))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L24
	} else {
		goto L243
	}
L220:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v868 = v866 + int32(1)
	if v865 != v868 {
		v875 = v866
		v876 = v868
		goto L236
	} else {
		goto L237
	}
L221:
	;
	switch v862 + int32(-3) {
	case 0:
		goto L220
	case 1:
		goto L219
	default:
		goto L218
	}
L222:
	;
	v856 = m.G398
	if v822 != v856 {
		goto L234
	} else {
		goto L235
	}
L223:
	;
	goto L227
L227:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v822 = v819 + int32(-32)
	goto L222
L234:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v822)+8))
	v862 = v859
	goto L221
L235:
	;
	v862 = int32(-1)
	goto L221
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v876
	v878 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v880 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v878+v875))) = uint8(v880)
	F_json_append_number(m, l0, l1, l3, int32(-2))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L24
	} else {
		goto L239
	}
L237:
	;
	F_strbuf_resize(m, l3, v865)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L24
	} else {
		goto L238
	}
L238:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v875 = v872
	v876 = v872 + int32(1)
	goto L236
L239:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if base.Ui32(v885-v886) < base.Ui32(int32(-2)) {
		v895 = v885
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v898 = int32(14882)
	*(*uint16)(unsafe.Add(mBase, uint32(v896+v895))) = uint16(v898)
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v900 + int32(2)
	goto L217
L241:
	;
	F_strbuf_resize(m, l3, v885+int32(2))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L24
	} else {
		goto L242
	}
L242:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v895 = v894
	goto L240
L243:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v910 = v908 + int32(1)
	if v907 != v910 {
		v917 = v908
		v918 = v910
		goto L244
	} else {
		goto L245
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v918
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v922 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v920+v917))) = uint8(v922)
	goto L217
L245:
	;
	F_strbuf_resize(m, l3, v907)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L24
	} else {
		goto L246
	}
L246:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v917 = v914
	v918 = v914 + int32(1)
	goto L244
L247:
	;
	v927 = m.G3
	goto L252
L248:
	;
	F_strbuf_free(m, l3)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L24
	} else {
		goto L249
	}
L249:
	;
	goto L247
L250:
	;
	v987 = m.G3
	if v985 != int32(-1) {
		goto L266
	} else {
		goto L267
	}
L251:
	;
	v979 = m.G398
	if v945 != v979 {
		goto L263
	} else {
		goto L264
	}
L252:
	;
	goto L256
L256:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v945 = v942 + int32(-32)
	goto L251
L263:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v945)+8))
	v985 = v982
	goto L250
L264:
	;
	v985 = int32(-1)
	goto L250
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v927 + int32(_a2130)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v997
	v1006 = F_luaL_error(m, l0, v927+int32(_a2129), v14+int32(16))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L24
	} else {
		goto L268
	}
L266:
	;
	v992 = m.G399
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v992+v985<<(uint(int32(2))%32))))
	v997 = v996
	goto L265
L267:
	;
	v997 = v987 + int32(_a2018)
	goto L265
L268:
	;
	goto L217
L269:
	;
	goto L272
L270:
	;
	v1042 = F_lua_next(m, l0, int32(-2))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L24
	} else {
		goto L278
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1031 + int32(-16)
	goto L270
L272:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L271
L278:
	;
	if v1042 == int32(0) {
		goto L212
	} else {
		goto L279
	}
L279:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v1049 = v1047 + int32(1)
	if v1046 != v1049 {
		v1056 = v1047
		v1057 = v1049
		goto L280
	} else {
		goto L281
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v1057
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v1061 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v1059+v1056))) = uint8(v1061)
	goto L215
L281:
	;
	F_strbuf_resize(m, l3, v1046)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L24
	} else {
		goto L282
	}
L282:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v1056 = v1053
	v1057 = v1053 + int32(1)
	goto L280
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v1085
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v1089 = int32(125)
	*(*uint8)(unsafe.Add(mBase, uint32(v1087+v1084))) = uint8(v1089)
	goto L1
L284:
	;
	F_strbuf_resize(m, l3, v1074)
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L24
	} else {
		goto L285
	}
L285:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v1084 = v1081
	v1085 = v1081 + int32(1)
	goto L283
L286:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v1102+v1101))) = int32(1819047278)
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v1106 + int32(4)
	goto L1
L287:
	;
	F_strbuf_resize(m, l3, v1091+int32(4))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L24
	} else {
		goto L288
	}
L288:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v1101 = v1100
	goto L286
L289:
	;
	if v1175 != 0 {
		goto L2
	} else {
		goto L308
	}
L290:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+8))
	switch v1166 + int32(-2) {
	case 0:
		goto L306
	default:
		v1173 = int32(0)
		goto L305
	case 5:
		goto L307
	}
L291:
	;
	goto L297
L297:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1129 = v1126 + int32(-16)
	goto L290
L305:
	;
	v1175 = v1173
	goto L289
L306:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1129)))
	v1173 = v1172
	goto L305
L307:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1129)))
	v1175 = v1169 + int32(24)
	goto L289
L308:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if base.Ui32(v1176-v1177) < base.Ui32(int32(-4)) {
		v1186 = v1176
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v1187+v1186))) = int32(1819047278)
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v1191 + int32(4)
	goto L1
L310:
	;
	F_strbuf_resize(m, l3, v1176+int32(4))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L24
	} else {
		goto L311
	}
L311:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v1186 = v1185
	goto L309
L312:
	;
	v1198 = m.G3
	goto L317
L313:
	;
	F_strbuf_free(m, l3)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L24
	} else {
		goto L314
	}
L314:
	;
	goto L312
L315:
	;
	v1258 = m.G3
	if v1256 != int32(-1) {
		goto L331
	} else {
		goto L332
	}
L316:
	;
	v1250 = m.G398
	if v1216 != v1250 {
		goto L328
	} else {
		goto L329
	}
L317:
	;
	goto L321
L321:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1216 = v1213 + int32(-16)
	goto L316
L328:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1216)+8))
	v1256 = v1253
	goto L315
L329:
	;
	v1256 = int32(-1)
	goto L315
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v1198 + int32(_a2131)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v1268
	v1275 = F_luaL_error(m, l0, v1198+int32(_a2129), v14)
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L24
	} else {
		goto L333
	}
L331:
	;
	v1263 = m.G399
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1263+v1256<<(uint(int32(2))%32))))
	v1268 = v1267
	goto L330
L332:
	;
	v1268 = v1258 + int32(_a2018)
	goto L330
L333:
	;
	goto L1
}
func F_json_arg_init(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if (v3-v4)>>(uint(int32(4))%32) <= l1 {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if l1 <= (v18-v19)>>(uint(int32(4))%32) {
		} else {
			for {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27 + int32(16)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if (v33-v34)>>(uint(int32(4))%32) < l1 {
					continue
				} else {
					break
				}
				break
			}
		}
		switch int32(-1) {
		case 0:
			v94 = l0 + int32(72)
		case 1:
			v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
			v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v69
			v94 = l0 + int32(88)
		case 2:
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v94 = v63 + int32(96)
		default:
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
			v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
			v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+7)))
			v81 = m.G398
			if base.Ui32(v80) < base.Ui32(int32(1)) {
				v92 = v81
			} else {
				v92 = v79 + int32(24)
			}
			v94 = v92
		}
		v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
		switch v97 + int32(-2) {
		case 0:
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
			v104 = v103
			v106 = v104
		default:
			v104 = int32(0)
			v106 = v104
		case 5:
			v100 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
			v106 = v100 + int32(24)
		}
		if v106 != 0 {
			return v106
		} else {
			v107 = m.G3
			v111 = F_luaL_error(m, l0, v107+int32(_a2124), int32(0))
			mBase = m.M
			v112 = m.ExcPending
			if v112 != 0 {
				return int32(0)
			} else {
				return v106
			}
		}
	} else {
		v11 = m.G3
		v14 = F_luaL_argerror(m, l0, l1+int32(1), v11+int32(_a2140))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if l1 <= (v18-v19)>>(uint(int32(4))%32) {
			} else {
				for {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27 + int32(16)
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if (v33-v34)>>(uint(int32(4))%32) < l1 {
						continue
					} else {
						break
					}
					break
				}
			}
			switch int32(-1) {
			case 0:
				v94 = l0 + int32(72)
			case 1:
				v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v69
				v94 = l0 + int32(88)
			case 2:
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v94 = v63 + int32(96)
			default:
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+7)))
				v81 = m.G398
				if base.Ui32(v80) < base.Ui32(int32(1)) {
					v92 = v81
				} else {
					v92 = v79 + int32(24)
				}
				v94 = v92
			}
			v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
			switch v97 + int32(-2) {
			case 0:
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
				v104 = v103
				v106 = v104
			default:
				v104 = int32(0)
				v106 = v104
			case 5:
				v100 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
				v106 = v100 + int32(24)
			}
			if v106 != 0 {
				return v106
			} else {
				v107 = m.G3
				v111 = F_luaL_error(m, l0, v107+int32(_a2124), int32(0))
				mBase = m.M
				v112 = m.ExcPending
				if v112 != 0 {
					return int32(0)
				} else {
					return v106
				}
			}
		}
	}
}
func F_json_cfg_decode_max_depth(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v10 = F_json_arg_init(m, l0, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v22 = v17 + int32(0)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v22) < base.Ui32(v23) {
			v65 = m.G398
			if v22 != v65 {
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
				v71 = v68
			} else {
				v71 = int32(-1)
			}
		} else {
			v71 = int32(-1)
		}
		if v71 != 0 {
			v74 = F_luaL_checkinteger(m, l0, int32(1))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(9223372032559808513)
				v81 = m.G3
				v84 = F_snprintf(m, v7+int32(16), int32(64), v81+int32(_a2125), v7)
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int32(0)
				} else {
					if int32(0) < v74 {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+1336)) = v74
						v94 = v74
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(v96))) = base.F64_convert_i32_s(v94)
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v101 + int32(16)
						m.G0 = v7 + int32(80)
						return int32(1)
					} else {
						v91 = F_luaL_argerror(m, l0, int32(1), v7+int32(16))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+1336)) = v74
							v94 = v74
							v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = int32(3)
							*(*float64)(unsafe.Add(mBase, uint32(v96))) = base.F64_convert_i32_s(v94)
							v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v101 + int32(16)
							m.G0 = v7 + int32(80)
							return int32(1)
						}
					}
				}
			}
		} else {
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1336))
			v94 = v72
			v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v96))) = base.F64_convert_i32_s(v94)
			v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v101 + int32(16)
			m.G0 = v7 + int32(80)
			return int32(1)
		}
	}
}
func F_json_cfg_encode_invalid_numbers(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	v3 = m.G3
	v5 = F_json_arg_init(m, l0, int32(1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		F_json_enum_option(m, l0, v5+int32(1320), v3+int32(_a2126))
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return int32(1)
		}
	}
}
func F_json_cfg_encode_keep_buffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	v5 = F_json_arg_init(m, l0, int32(1))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+1328))
		F_json_enum_option(m, l0, v5+int32(1328), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)+1328))
			if v9 == v15 {
				return int32(1)
			} else {
				v18 = v5 + int32(1280)
				if v15 == int32(0) {
					F_strbuf_free(m, v18)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						return int32(1)
					}
				} else {
					F_strbuf_init(m, v18, int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						return int32(1)
					}
				}
			}
		}
	}
}
func F_json_destroy_config(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = v7 + int32(0)
	v13 = m.G398
	if base.Ui32(v12) < base.Ui32(v6) {
		v15 = v12
	} else {
		v15 = v13
	}
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	switch v58 + int32(-2) {
	case 0:
		v64 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		v65 = v64
		v67 = v65
	default:
		v65 = int32(0)
		v67 = v65
	case 5:
		v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		v67 = v61 + int32(24)
	}
	if v67 == int32(0) {
		return int32(0)
	} else {
		F_strbuf_free(m, v67+int32(1280))
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_json_encode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	switch int32(-1) {
	case 0:
		v62 = l0 + int32(72)
	case 1:
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v37
		v62 = l0 + int32(88)
	case 2:
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v62 = v31 + int32(96)
	default:
		v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
		v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+7)))
		v49 = m.G398
		if base.Ui32(v48) < base.Ui32(int32(1)) {
			v60 = v49
		} else {
			v60 = v47 + int32(24)
		}
		v62 = v60
	}
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	switch v65 + int32(-2) {
	case 0:
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
		v72 = v71
		v74 = v72
	default:
		v72 = int32(0)
		v74 = v72
	case 5:
		v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
		v74 = v68 + int32(24)
	}
	if v74 != 0 {
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if (v83-v84)>>(uint(int32(4))%32) == int32(1) {
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v74)+1328))
			if v96 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v74)+1288)) = int32(0)
				v108 = v74 + int32(1280)
				F_json_append_data(m, l0, v74, int32(0), v108)
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int32(0)
				} else {
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
					F_lua_pushlstring(m, l0, v112, v113)
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return int32(0)
					} else {
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v74)+1328))
						if v116 != 0 {
							m.G0 = v7 + int32(32)
							return int32(1)
						} else {
							F_strbuf_free(m, v108)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(32)
								return int32(1)
							}
						}
					}
				}
			} else {
				F_strbuf_init(m, v7+int32(8), int32(0))
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return int32(0)
				} else {
					v108 = v7 + int32(8)
					F_json_append_data(m, l0, v74, int32(0), v108)
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return int32(0)
					} else {
						v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
						F_lua_pushlstring(m, l0, v112, v113)
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v74)+1328))
							if v116 != 0 {
								m.G0 = v7 + int32(32)
								return int32(1)
							} else {
								F_strbuf_free(m, v108)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(32)
									return int32(1)
								}
							}
						}
					}
				}
			}
		} else {
			v91 = m.G3
			v94 = F_luaL_argerror(m, l0, int32(1), v91+int32(_a2123))
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return int32(0)
			} else {
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v74)+1328))
				if v96 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v74)+1288)) = int32(0)
					v108 = v74 + int32(1280)
					F_json_append_data(m, l0, v74, int32(0), v108)
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return int32(0)
					} else {
						v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
						F_lua_pushlstring(m, l0, v112, v113)
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v74)+1328))
							if v116 != 0 {
								m.G0 = v7 + int32(32)
								return int32(1)
							} else {
								F_strbuf_free(m, v108)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(32)
									return int32(1)
								}
							}
						}
					}
				} else {
					F_strbuf_init(m, v7+int32(8), int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						v108 = v7 + int32(8)
						F_json_append_data(m, l0, v74, int32(0), v108)
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return int32(0)
						} else {
							v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
							F_lua_pushlstring(m, l0, v112, v113)
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return int32(0)
							} else {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v74)+1328))
								if v116 != 0 {
									m.G0 = v7 + int32(32)
									return int32(1)
								} else {
									F_strbuf_free(m, v108)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										m.G0 = v7 + int32(32)
										return int32(1)
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v75 = m.G3
		v79 = F_luaL_error(m, l0, v75+int32(_a2124), int32(0))
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return int32(0)
		} else {
			v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if (v83-v84)>>(uint(int32(4))%32) == int32(1) {
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v74)+1328))
				if v96 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v74)+1288)) = int32(0)
					v108 = v74 + int32(1280)
					F_json_append_data(m, l0, v74, int32(0), v108)
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return int32(0)
					} else {
						v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
						F_lua_pushlstring(m, l0, v112, v113)
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v74)+1328))
							if v116 != 0 {
								m.G0 = v7 + int32(32)
								return int32(1)
							} else {
								F_strbuf_free(m, v108)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(32)
									return int32(1)
								}
							}
						}
					}
				} else {
					F_strbuf_init(m, v7+int32(8), int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						v108 = v7 + int32(8)
						F_json_append_data(m, l0, v74, int32(0), v108)
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return int32(0)
						} else {
							v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
							F_lua_pushlstring(m, l0, v112, v113)
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return int32(0)
							} else {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v74)+1328))
								if v116 != 0 {
									m.G0 = v7 + int32(32)
									return int32(1)
								} else {
									F_strbuf_free(m, v108)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										m.G0 = v7 + int32(32)
										return int32(1)
									}
								}
							}
						}
					}
				}
			} else {
				v91 = m.G3
				v94 = F_luaL_argerror(m, l0, int32(1), v91+int32(_a2123))
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int32(0)
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v74)+1328))
					if v96 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v74)+1288)) = int32(0)
						v108 = v74 + int32(1280)
						F_json_append_data(m, l0, v74, int32(0), v108)
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return int32(0)
						} else {
							v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
							F_lua_pushlstring(m, l0, v112, v113)
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return int32(0)
							} else {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v74)+1328))
								if v116 != 0 {
									m.G0 = v7 + int32(32)
									return int32(1)
								} else {
									F_strbuf_free(m, v108)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										m.G0 = v7 + int32(32)
										return int32(1)
									}
								}
							}
						}
					} else {
						F_strbuf_init(m, v7+int32(8), int32(0))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							v108 = v7 + int32(8)
							F_json_append_data(m, l0, v74, int32(0), v108)
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
								F_lua_pushlstring(m, l0, v112, v113)
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									v116 = *(*int32)(unsafe.Add(mBase, uint32(v74)+1328))
									if v116 != 0 {
										m.G0 = v7 + int32(32)
										return int32(1)
									} else {
										F_strbuf_free(m, v108)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											m.G0 = v7 + int32(32)
											return int32(1)
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
func F_json_next_number_token(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(5)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = F_fpconv_strtod(m, v12, v8+int32(12))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v15
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		if v18 != v19 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
			v23 = m.G3
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v23 + int32(_a2141)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v18 - v24
		}
		m.G0 = v8 + int32(16)
		return
	}
}
func F_json_process_value(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	v8 = m.G0
	v10 = v8 - int32(128)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v12 {
	case 0:
		goto L5
	default:
		goto L2
	case 2:
		goto L4
	case 4:
		goto L8
	case 5:
		goto L7
	case 6:
		goto L6
	case 7:
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(128)
	return
L2:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_strbuf_free(m, v342)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L9
	} else {
		goto L79
	}
L3:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v334)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v334))) = int32(0)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v338 + int32(16)
	goto L78
L4:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v223 + int32(1)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+1336))
	if v228 <= v223 {
		goto L55
	} else {
		goto L56
	}
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v39 + int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+1336))
	if v44 <= v39 {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = base.B2i32(v27 != int32(0))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v35 + int32(16)
	goto L12
L7:
	;
	v17 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v19))) = v17
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23 + int32(16)
	goto L11
L8:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	F_lua_pushlstring(m, l0, v13, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	goto L1
L11:
	;
	goto L1
L12:
	;
	goto L1
L13:
	;
	v67 = int32(0)
	F_lua_createtable(m, l0, v67, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L9
	} else {
		goto L20
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_strbuf_free(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L18
	}
L15:
	;
	v47 = F_lua_checkstack(m, l0, int32(3))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	if v47 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v53 - v52
	v58 = m.G3
	v63 = F_luaL_error(m, l0, v58+int32(_a2132), v10+int32(64))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	goto L13
L20:
	;
	F_json_next_token(m, l1, v10+int32(104))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v10)+104))
	if v75 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v219 + int32(-1)
	goto L1
L23:
	;
	v81 = v10 + int32(112)
	v84 = v75
	goto L24
L24:
	;
	if v84 == int32(4) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v10)+112))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v10)+120))
	F_lua_pushlstring(m, l0, v120, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L9
	} else {
		goto L33
	}
L27:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_strbuf_free(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v94 = m.G3
	if v84 == int32(12) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v102 = v81
	goto L31
L30:
	;
	v102 = v94 + int32(_a2133) + v84<<(uint(int32(2))%32)
	goto L31
L31:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v94 + int32(_a2134)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v103
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v108 + int32(1)
	v116 = F_luaL_error(m, l0, v94+int32(_a2135), v10+int32(48))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	goto L26
L33:
	;
	F_json_next_token(m, l1, v10+int32(104))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v10)+104))
	if v128 == int32(8) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_json_next_token(m, l1, v10+int32(104))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L42
	}
L36:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_strbuf_free(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	v134 = m.G3
	if v128 == int32(12) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v142 = v81
	goto L40
L39:
	;
	v142 = v134 + int32(_a2133) + v128<<(uint(int32(2))%32)
	goto L40
L40:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v134 + int32(_a2136)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v143
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v148 + int32(1)
	v156 = F_luaL_error(m, l0, v134+int32(_a2135), v10+int32(32))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	goto L35
L42:
	;
	F_json_process_value(m, l0, l1, v10+int32(104))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	F_lua_rawset(m, l0, int32(-3))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	F_json_next_token(m, l1, v10+int32(104))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L9
	} else {
		goto L45
	}
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v10)+104))
	switch v175 + int32(-1) {
	case 0:
		goto L22
	default:
		goto L47
	case 8:
		goto L46
	}
L46:
	;
	F_json_next_token(m, l1, v10+int32(104))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L9
	} else {
		goto L53
	}
L47:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_strbuf_free(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	v181 = m.G3
	if v175 == int32(12) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v189 = v81
	goto L51
L50:
	;
	v189 = v181 + int32(_a2133) + v175<<(uint(int32(2))%32)
	goto L51
L51:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v181 + int32(_a2137)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v190
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v195 + int32(1)
	v203 = F_luaL_error(m, l0, v181+int32(_a2135), v10+int32(16))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	goto L46
L53:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v10)+104))
	v84 = v211
	goto L24
L54:
	;
	v251 = int32(0)
	F_lua_createtable(m, l0, v251, v251)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L9
	} else {
		goto L61
	}
L55:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_strbuf_free(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L9
	} else {
		goto L59
	}
L56:
	;
	v231 = F_lua_checkstack(m, l0, int32(2))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	if v231 != 0 {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = v237 - v236
	v242 = m.G3
	v247 = F_luaL_error(m, l0, v242+int32(_a2132), v10+int32(96))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	goto L54
L61:
	;
	F_json_next_token(m, l1, v10+int32(104))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v10)+104))
	if v259 == int32(3) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v328 + int32(-1)
	goto L1
L64:
	;
	v267 = int32(1)
	goto L65
L65:
	;
	F_json_process_value(m, l0, l1, v10+int32(104))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	F_lua_rawseti(m, l0, int32(-2), v267)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	F_json_next_token(m, l1, v10+int32(104))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v10)+104))
	switch v283 + int32(-3) {
	case 0:
		goto L63
	default:
		goto L71
	case 6:
		goto L70
	}
L70:
	;
	F_json_next_token(m, l1, v10+int32(104))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L9
	} else {
		goto L77
	}
L71:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_strbuf_free(m, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	v289 = m.G3
	if v283 == int32(12) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v297 = v10 + int32(112)
	goto L75
L74:
	;
	v297 = v289 + int32(_a2133) + v283<<(uint(int32(2))%32)
	goto L75
L75:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v289 + int32(_a2138)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v298
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = v303 + int32(1)
	v311 = F_luaL_error(m, l0, v289+int32(_a2135), v10+int32(80))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	goto L70
L77:
	;
	v267 = v267 + int32(1)
	goto L65
L78:
	;
	goto L1
L79:
	;
	v347 = m.G3
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v350 == int32(12) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v356 = l2 + int32(8)
	goto L82
L81:
	;
	v356 = v347 + int32(_a2133) + v350<<(uint(int32(2))%32)
	goto L82
L82:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v358 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v347 + int32(_a2139)
	v368 = F_luaL_error(m, l0, v347+int32(_a2135), v10)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L9
	} else {
		goto L83
	}
L83:
	;
	goto L1
}
