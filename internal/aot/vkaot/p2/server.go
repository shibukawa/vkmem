package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F__serverLog(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v5 = m.G0
	v7 = v5 - int32(1040)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+1036)) = l2
	v11 = F_vsnprintf(m, v7, int32(1024), l1, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		F_serverLogRaw(m, l0, v7)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			m.G0 = v7 + int32(1040)
			return
		}
	}
}
func F_appendServerSaveParams(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v5 = int32(_a44)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[228]))
	v9 = *(*int32)(unsafe.Add(mBase, _consts[229]))
	v14 = F_valkey_realloc(m, v7, v9<<(uint(int32(4))%32)+int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[228])) = v14
		v17 = int32(_a44)
		v18 = *(*int32)(unsafe.Add(mBase, _consts[229]))
		v21 = v14 + v18<<(uint(int32(4))%32)
		*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v21))) = l0
		*(*int32)(unsafe.Add(mBase, _consts[229])) = v18 + int32(1)
		return
	}
}
func F_initServer(m *base.Module) {
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
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v76 int32
	_ = v76
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v225 int32
	_ = v225
	var v226 int64
	_ = v226
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
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
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int64
	_ = v310
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int64
	_ = v352
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int64
	_ = v501
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int64
	_ = v543
	var v545 int64
	_ = v545
	var v548 int64
	_ = v548
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int64
	_ = v582
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int64
	_ = v601
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int64
	_ = v697
	var v705 int32
	_ = v705
	var v707 int64
	_ = v707
	var v710 int64
	_ = v710
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int64
	_ = v739
	var v757 int32
	_ = v757
	var v762 int64
	_ = v762
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v772 int64
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v800 int64
	_ = v800
	var v804 int32
	_ = v804
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int64
	_ = v852
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v887 int32
	_ = v887
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
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
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1047 int32
	_ = v1047
	var v1054 int32
	_ = v1054
	var v1061 int32
	_ = v1061
	var v1068 int32
	_ = v1068
	var v1075 int32
	_ = v1075
	v5 = m.G0
	v7 = v5 - int32(192)
	m.G0 = v7
	v12 = m.G0
	v13 = int32(288)
	v14 = v12 - v13
	m.G0 = v14
	v20 = F___memset(m, v14+int32(12), int32(0), int32(136))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+140)) = int32(268435456)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = int32(-2)
	v28 = F___sigaction(m, int32(1), v14+int32(8), v14+int32(148))
	mBase = m.M
	m.G0 = v14 + v13
	goto L2
L1:
	;
	v40 = m.G0
	v41 = int32(288)
	v42 = v40 - v41
	m.G0 = v42
	v48 = F___memset(m, v42+int32(12), int32(0), int32(136))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v42)+140)) = int32(268435456)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = int32(-2)
	v56 = F___sigaction(m, int32(13), v42+int32(8), v42+int32(148))
	mBase = m.M
	m.G0 = v42 + v41
	goto L6
L2:
	;
	goto L4
L4:
	;
	goto L1
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(56)))) = int64(0)
	goto L9
L6:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = int32(1023)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+184)) = int32(0)
	v76 = v7 + int32(52)
	goto L11
L10:
	;
	v104 = v7 + int32(52)
	goto L18
L11:
	;
	goto L13
L13:
	;
	if v76 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	goto L10
L16:
	;
	v99 = F___memcpy(m, int32(9119284), v76, int32(140))
	mBase = m.M
	goto L15
L17:
	;
	v131 = m.G0
	v132 = int32(144)
	v133 = v131 - v132
	m.G0 = v133
	F_setupSigSegvHandler(m)
	mBase = m.M
	v136 = int32(4)
	v140 = F_sigemptyset(m, v133+int32(8))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = int32(518)
	*(*int32)(unsafe.Add(mBase, uint32(v133)+136)) = v136
	v149 = F___sigaction(m, int32(14), v133+v136, int32(0))
	mBase = m.M
	m.G0 = v133 + v132
	goto L24
L18:
	;
	goto L20
L20:
	;
	if v104 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	goto L17
L23:
	;
	v127 = F___memcpy(m, int32(9117464), v104, int32(140))
	mBase = m.M
	goto L22
L24:
	;
	goto L25
L25:
	;
	goto L26
L26:
	;
	goto L27
L27:
	;
	v159 = int32(0)
	v160 = *(*int32)(unsafe.Add(mBase, _consts[795]))
	if v160 == v159 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v215 = int32(0)
	v216 = int64(-1)
	*(*int64)(unsafe.Add(mBase, _consts[875])) = v216
	*(*int64)(unsafe.Add(mBase, _consts[643])) = v216
	v225 = *(*int32)(unsafe.Add(mBase, _consts[876]))
	if v225 != 0 {
		goto L37
	} else {
		goto L38
	}
L29:
	;
	v163 = int32(0)
	v164 = *(*int32)(unsafe.Add(mBase, _consts[869]))
	v167 = *(*int32)(unsafe.Add(mBase, _consts[870]))
	v170 = m.G0
	v172 = v170 - int32(16)
	m.G0 = v172
	v177 = F_pthread_setcancelstate(m, int32(1), v172+int32(12))
	mBase = m.M
	F___lock(m, int32(9128376))
	mBase = m.M
	if v164 == v163 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L28
L31:
	;
	v194 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[872])) = v167
	*(*int32)(unsafe.Add(mBase, _consts[873])) = int32(25)
	goto L35
L32:
	;
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[871])) = uint8(v190)
	goto L31
L33:
	;
	v184 = F_strnlen(m, v164, int32(31))
	mBase = m.M
	v185 = F___memcpy(m, int32(9128384), v164, v184)
	mBase = m.M
	v188 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v184)+uint32(_consts[871]))) = uint8(v188)
	goto L31
L34:
	;
	F___unlock(m, int32(9128376))
	mBase = m.M
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v211 = F_pthread_setcancelstate(m, v209, int32(0))
	mBase = m.M
	m.G0 = v172 + int32(16)
	goto L30
L35:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _consts[874]))
	if int32(-1) < v203 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	F___openlog(m)
	mBase = m.M
	goto L34
L37:
	;
	v226 = int64(0)
	goto L39
L38:
	;
	v226 = v216
	goto L39
L39:
	;
	*(*int64)(unsafe.Add(mBase, _consts[31])) = v226
	v228 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[28])) = base.B2i32(v225 != v228)
	*(*int32)(unsafe.Add(mBase, _consts[400])) = v228
	v236 = F___get_tp(m)
	mBase = m.M
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, _consts[719])) = v236
	v238 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[67])) = v238
	v242 = F_raxNew(m)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	return
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[591])) = v242
	v245 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[95])) = v245
	v249 = F_listCreate(m)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v249
	v253 = F_raxNew(m)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, _consts[404])) = v253
	v257 = F_listCreate(m)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, _consts[593])) = v257
	v261 = F_listCreate(m)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, _consts[158])) = v261
	v265 = F_listCreate(m)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, _consts[413])) = v265
	v268 = F_raxNew(m)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L41
	} else {
		goto L48
	}
L48:
	;
	v270 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[398])) = int32(5)
	*(*int32)(unsafe.Add(mBase, _consts[654])) = v268
	v276 = F_listCreate(m)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L41
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, _consts[216])) = v276
	v279 = F_raxNew(m)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L41
	} else {
		goto L50
	}
L50:
	;
	v281 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[371])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[349])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[877])) = v279
	v290 = F_listCreate(m)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L41
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _consts[99])) = v290
	v294 = F_listCreate(m)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L41
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, _consts[92])) = v294
	v298 = F_listCreate(m)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L41
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, _consts[878])) = v298
	v302 = F_listCreate(m)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L41
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _consts[879])) = v302
	v306 = F_listCreate(m)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L41
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, _consts[880])) = v306
	v309 = int32(0)
	v310 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[597])) = v310
	*(*int64)(unsafe.Add(mBase, _consts[598])) = v310
	*(*int64)(unsafe.Add(mBase, _consts[881])) = v310
	*(*int64)(unsafe.Add(mBase, _consts[882])) = v310
	*(*int64)(unsafe.Add(mBase, _consts[883])) = v310
	*(*int64)(unsafe.Add(mBase, _consts[884])) = v310
	*(*int64)(unsafe.Add(mBase, _consts[213])) = v310
	*(*int64)(unsafe.Add(mBase, _consts[885])) = v310
	*(*int32)(unsafe.Add(mBase, _consts[208])) = v309
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v309
	v339 = F_listCreate(m)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L41
	} else {
		goto L56
	}
L56:
	;
	v341 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[886])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[887])) = v339
	v347 = F_sysconf(m, int32(85))
	mBase = m.M
	v349 = F_sysconf(m, int32(30))
	mBase = m.M
	goto L57
L57:
	;
	v351 = int32(0)
	v352 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[528])) = v352
	*(*int32)(unsafe.Add(mBase, _consts[888])) = v347 * v349
	*(*int32)(unsafe.Add(mBase, _consts[169])) = int32(-1)
	*(*int64)(unsafe.Add(mBase, _consts[395])) = int64(4294972296)
	*(*int64)(unsafe.Add(mBase, _consts[402])) = v352
	v367 = int32(*(*uint8)(unsafe.Add(mBase, _consts[201])))
	v369 = v367 & int32(249)
	*(*uint8)(unsafe.Add(mBase, _consts[201])) = uint8(v369)
	*(*int32)(unsafe.Add(mBase, _consts[527])) = v351
	*(*int32)(unsafe.Add(mBase, _consts[889])) = v351
	*(*int32)(unsafe.Add(mBase, _consts[393])) = v351
	F_resetReplicationBuffer(m)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L41
	} else {
		goto L58
	}
L58:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _consts[274]))
	v385 = F_setlocale(m, int32(3), v384)
	mBase = m.M
	if v385 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v455 = int32(0)
	v459 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v459 != 0 {
		goto L82
	} else {
		goto L83
	}
L60:
	;
	if int32(3) < v387 {
		goto L79
	} else {
		goto L80
	}
L61:
	;
	F_createSharedObjects(m)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L41
	} else {
		goto L66
	}
L62:
	;
	v386 = int32(0)
	v387 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v389 = *(*int32)(unsafe.Add(mBase, _consts[274]))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	if v390 != 0 {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	if int32(3) < v387 {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v393 = int32(3)
	v395 = F_setlocale(m, v393, int32(0))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v395
	F__serverLog(m, v393, int32(_a1600), v7+int32(32))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L41
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	F_adjustOpenFilesLimit(m)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L41
	} else {
		goto L67
	}
L67:
	;
	v409 = F_monotonicInit(m)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L41
	} else {
		goto L68
	}
L68:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v412 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v422 = int32(0)
	v424 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v427 = F_aeCreateEventLoop(m, v424+int32(128))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L41
	} else {
		goto L72
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v409
	F__serverLog(m, int32(2), int32(_a1599), v7+int32(16))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L41
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	*(*int32)(unsafe.Add(mBase, _consts[279])) = v427
	if v427 != 0 {
		goto L59
	} else {
		goto L73
	}
L73:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v431 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	goto L76
L76:
	;
	v435 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v436 = F___strerror_l(m, v435, v435)
	mBase = m.M
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v436
	F__serverLog(m, int32(3), int32(_a1598), v7)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L41
	} else {
		goto L78
	}
L78:
	;
	goto L74
L79:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v389
	F__serverLog(m, int32(3), int32(_a1601), v7+int32(48))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L41
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v460 = int32(_a1587)
	goto L84
L83:
	;
	v460 = int32(_a1588)
	goto L84
L84:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	*(*int32)(unsafe.Add(mBase, _consts[65])) = v461
	v466 = F_valkey_calloc(m, v461<<(uint(int32(2))%32))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L41
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v466
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v466)))
	if v469 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_evictionPoolAlloc(m)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L41
	} else {
		goto L89
	}
L87:
	;
	v471 = F_createDatabase(m, int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L41
	} else {
		goto L88
	}
L88:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*int32)(unsafe.Add(mBase, uint32(v474))) = v471
	goto L86
L89:
	;
	v479 = int32(0)
	v483 = F_kvstoreCreate(m, int32(_a1589), v479, int32(1))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L41
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, _consts[19])) = v483
	v488 = F_dictCreate(m, int32(_a1590))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L41
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = v488
	v493 = int32(0)
	v495 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v495 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v496 = int32(14)
	goto L94
L93:
	;
	v496 = v493
	goto L94
L94:
	;
	v498 = F_kvstoreCreate(m, int32(_a1589), v496, int32(3))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L41
	} else {
		goto L95
	}
L95:
	;
	v500 = int32(0)
	v501 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[623])) = v501
	*(*int32)(unsafe.Add(mBase, _consts[20])) = v498
	*(*int64)(unsafe.Add(mBase, _consts[45])) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, _consts[544])) = v501
	*(*int64)(unsafe.Add(mBase, _consts[645])) = v501
	*(*int64)(unsafe.Add(mBase, _consts[851])) = v501
	*(*int64)(unsafe.Add(mBase, _consts[649])) = v501
	*(*int64)(unsafe.Add(mBase, _consts[108])) = int64(-1)
	*(*int32)(unsafe.Add(mBase, _consts[599])) = v500
	*(*int32)(unsafe.Add(mBase, _consts[890])) = v500
	*(*int32)(unsafe.Add(mBase, _consts[220])) = v500
	*(*int32)(unsafe.Add(mBase, _consts[671])) = v500
	*(*int32)(unsafe.Add(mBase, _consts[109])) = v500
	v539 = F_sdsempty(m)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L41
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, _consts[34])) = v539
	v542 = int32(0)
	v543 = F___time(m, v542)
	mBase = m.M
	v545 = int64(-1)
	*(*int64)(unsafe.Add(mBase, _consts[891])) = v545
	v548 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[846])) = v548
	*(*int64)(unsafe.Add(mBase, _consts[845])) = v543
	*(*int64)(unsafe.Add(mBase, _consts[652])) = v545
	*(*int64)(unsafe.Add(mBase, _consts[637])) = v548
	*(*int64)(unsafe.Add(mBase, _consts[83])) = v548
	*(*int64)(unsafe.Add(mBase, _consts[640])) = v548
	*(*int64)(unsafe.Add(mBase, _consts[296])) = v548
	*(*int64)(unsafe.Add(mBase, _consts[297])) = v548
	*(*int32)(unsafe.Add(mBase, _consts[360])) = v542
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v542
	v580 = F__emscripten_memset_bulkmem(m, int32(_a525), base.I32_extend8_s(v542), int32(96))
	mBase = m.M
	goto L97
L97:
	;
	v581 = int32(0)
	v582 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[298])) = v582
	*(*int64)(unsafe.Add(mBase, _consts[299])) = v582
	*(*uint8)(unsafe.Add(mBase, _consts[300])) = uint8(v581)
	v594 = F__emscripten_memset_bulkmem(m, int32(_a526), base.I32_extend8_s(v581), int32(176))
	mBase = m.M
	goto L98
L98:
	;
	v599 = F__emscripten_memset_bulkmem(m, int32(_a527), base.I32_extend8_s(int32(0)), int32(80))
	mBase = m.M
	goto L99
L99:
	;
	v600 = int32(0)
	v601 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[301])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[302])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[124])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[303])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[304])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[305])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[306])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[94])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[307])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[308])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[309])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[310])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[311])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[312])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[313])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[314])) = v601
	*(*int64)(unsafe.Add(mBase, _consts[315])) = v601
	v655 = F__emscripten_memset_bulkmem(m, int32(_a528), base.I32_extend8_s(v600), int32(148))
	mBase = m.M
	goto L100
L100:
	;
	v660 = F__emscripten_memset_bulkmem(m, int32(_a529), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	goto L101
L101:
	;
	v665 = F__emscripten_memset_bulkmem(m, int32(_a530), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	goto L102
L102:
	;
	v670 = F__emscripten_memset_bulkmem(m, int32(_a531), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	goto L103
L103:
	;
	v675 = F__emscripten_memset_bulkmem(m, int32(_a532), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	goto L104
L104:
	;
	v680 = F__emscripten_memset_bulkmem(m, int32(_a533), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	goto L105
L105:
	;
	v685 = F__emscripten_memset_bulkmem(m, int32(_a534), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	goto L106
L106:
	;
	v690 = F__emscripten_memset_bulkmem(m, int32(_a535), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	goto L107
L107:
	;
	v695 = F__emscripten_memset_bulkmem(m, int32(_a536), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	goto L108
L108:
	;
	v696 = int32(0)
	v697 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[316])) = v697
	*(*int64)(unsafe.Add(mBase, _consts[317])) = v697
	*(*int32)(unsafe.Add(mBase, _consts[504])) = v696
	goto L109
L109:
	;
	v705 = int32(0)
	v707 = F___time(m, v705)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[892])) = v707
	v710 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[46])) = v710
	*(*int64)(unsafe.Add(mBase, _consts[811])) = v710
	*(*int64)(unsafe.Add(mBase, _consts[610])) = v710
	*(*int32)(unsafe.Add(mBase, _consts[847])) = v705
	*(*int32)(unsafe.Add(mBase, _consts[71])) = v705
	*(*int32)(unsafe.Add(mBase, _consts[668])) = v705
	*(*int32)(unsafe.Add(mBase, _consts[844])) = v705
	*(*int32)(unsafe.Add(mBase, _consts[612])) = v705
	v737 = F__emscripten_memset_bulkmem(m, int32(_a1591), base.I32_extend8_s(v705), int32(76))
	mBase = m.M
	goto L110
L110:
	;
	v738 = int32(0)
	v739 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[893])) = v739
	*(*int64)(unsafe.Add(mBase, _consts[894])) = v739
	*(*int64)(unsafe.Add(mBase, _consts[895])) = v739
	*(*int64)(unsafe.Add(mBase, _consts[896])) = v739
	*(*int64)(unsafe.Add(mBase, _consts[897])) = v739
	*(*int64)(unsafe.Add(mBase, _consts[898])) = v739
	v757 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v762 = F_aeCreateTimeEvent(m, v757, int64(1), int32(1024), v738, v738)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L41
	} else {
		goto L116
	}
L111:
	;
	F__serverPanic_1(m, int32(_a1555), int32(3092), int32(_a1596), int32(0))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L41
	} else {
		goto L172
	}
L112:
	;
	F__serverPanic_1(m, int32(_a1555), int32(3076), int32(_a1595), int32(0))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L41
	} else {
		goto L171
	}
L113:
	;
	F__serverPanic_1(m, int32(_a1555), int32(3056), int32(_a1594), int32(0))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L41
	} else {
		goto L170
	}
L114:
	;
	F__serverPanic_1(m, int32(_a1555), int32(3049), int32(_a1593), int32(0))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L41
	} else {
		goto L169
	}
L115:
	;
	F__serverPanic_1(m, int32(_a1555), int32(3043), int32(_a1592), int32(0))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L41
	} else {
		goto L168
	}
L116:
	;
	if v762 == int64(-1) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v766 = int32(0)
	v767 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v772 = F_aeCreateTimeEvent(m, v767, int64(1), int32(1025), v766, v766)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L41
	} else {
		goto L118
	}
L118:
	;
	if v772 == int64(-1) {
		goto L114
	} else {
		goto L119
	}
L119:
	;
	v776 = int32(0)
	v777 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v779 = *(*int32)(unsafe.Add(mBase, _consts[899]))
	v783 = F_aeCreateFileEvent(m, v777, v779, int32(1), int32(1026), v776)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L41
	} else {
		goto L120
	}
L120:
	;
	if v783 == int32(-1) {
		goto L113
	} else {
		goto L121
	}
L121:
	;
	v788 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	*(*int32)(unsafe.Add(mBase, uint32(v788)+36)) = int32(1027)
	goto L122
L122:
	;
	v792 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	*(*int32)(unsafe.Add(mBase, uint32(v792)+40)) = int32(1028)
	goto L123
L123:
	;
	v796 = *(*int32)(unsafe.Add(mBase, _consts[900]))
	if v796 != int32(32) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v818 = F_scriptingEngineManagerInit(m)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L41
	} else {
		goto L130
	}
L125:
	;
	v800 = *(*int64)(unsafe.Add(mBase, _consts[280]))
	if v800 != int64(0) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v804 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v804 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v812 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[97])) = int32(1792)
	*(*int64)(unsafe.Add(mBase, _consts[280])) = int64(3221225472)
	goto L124
L128:
	;
	F__serverLog(m, int32(3), int32(_a1597), int32(0))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L41
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	if v818 == int32(-1) {
		goto L112
	} else {
		goto L131
	}
L131:
	;
	v822 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[391])) = v822
	F_commandlogInit(m)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L41
	} else {
		goto L132
	}
L132:
	;
	F_latencyMonitorInit(m)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L41
	} else {
		goto L133
	}
L133:
	;
	F_initSharedQueryBuf(m)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L41
	} else {
		goto L134
	}
L134:
	;
	v832 = *(*int32)(unsafe.Add(mBase, _consts[753]))
	F_ACLUpdateDefaultUserPassword(m, v832)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L41
	} else {
		goto L135
	}
L135:
	;
	v835 = F_functionsInit(m)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L41
	} else {
		goto L136
	}
L136:
	;
	if v835 == int32(-1) {
		goto L111
	} else {
		goto L137
	}
L137:
	;
	F_evalInit(m)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L41
	} else {
		goto L138
	}
L138:
	;
	v844 = m.G0
	v846 = v844 - int32(32)
	m.G0 = v846
	v849 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	if v849 != 0 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v893 = int32(0)
	v894 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	if v894 == v893 {
		goto L145
	} else {
		goto L146
	}
L140:
	;
	m.G0 = v846 + int32(32)
	goto L139
L141:
	;
	v863 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v864 = base.I32_div_s(int32(1000), v863)
	v866 = v864 << (uint(int32(1)) % 32)
	if v866 <= v849 {
		v870 = v849
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v850 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v846)+24)) = v850
	v852 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v846)+16)) = v852
	*(*int32)(unsafe.Add(mBase, uint32(v846)+8)) = v850
	*(*int64)(unsafe.Add(mBase, uint32(v846))) = v852
	v860 = F_setitimer(m, v850, v846, v850)
	mBase = m.M
	goto L140
L143:
	;
	v871 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v846)+8)) = v871
	*(*int64)(unsafe.Add(mBase, uint32(v846))) = int64(0)
	v875 = int32(1000)
	v876 = base.I32_div_s(v870, v875)
	*(*int64)(unsafe.Add(mBase, uint32(v846)+16)) = base.I64_extend_i32_s(v876)
	*(*int32)(unsafe.Add(mBase, uint32(v846)+24)) = (v870 - v876*v875) * v875
	v887 = F_setitimer(m, v871, v846, v871)
	mBase = m.M
	goto L140
L144:
	;
	*(*int32)(unsafe.Add(mBase, _consts[275])) = v866
	v870 = v866
	goto L143
L145:
	;
	m.G0 = v7 + int32(192)
	return
L146:
	;
	v898 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	if v898 != 0 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v901 = F_valkey_malloc(m, int32(152))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L41
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, _consts[393])) = v901
	*(*int32)(unsafe.Add(mBase, uint32(v901)+4)) = int32(0)
	v906 = F_listCreate(m)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L41
	} else {
		goto L149
	}
L149:
	;
	v908 = int32(0)
	v909 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v909)+12)) = v908
	*(*int32)(unsafe.Add(mBase, uint32(v909))) = v906
	v913 = F_listCreate(m)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L41
	} else {
		goto L150
	}
L150:
	;
	v915 = int32(0)
	v916 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v916)+20)) = v915
	*(*int32)(unsafe.Add(mBase, uint32(v916)+8)) = v913
	v920 = F_listCreate(m)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L41
	} else {
		goto L151
	}
L151:
	;
	v922 = int32(0)
	v923 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v923)+28)) = v922
	*(*int32)(unsafe.Add(mBase, uint32(v923)+16)) = v920
	v927 = F_listCreate(m)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L41
	} else {
		goto L152
	}
L152:
	;
	v929 = int32(0)
	v930 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v930)+36)) = v929
	*(*int32)(unsafe.Add(mBase, uint32(v930)+24)) = v927
	v934 = F_listCreate(m)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L41
	} else {
		goto L153
	}
L153:
	;
	v936 = int32(0)
	v937 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v937)+44)) = v936
	*(*int32)(unsafe.Add(mBase, uint32(v937)+32)) = v934
	v941 = F_listCreate(m)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L41
	} else {
		goto L154
	}
L154:
	;
	v943 = int32(0)
	v944 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v944)+52)) = v943
	*(*int32)(unsafe.Add(mBase, uint32(v944)+40)) = v941
	v948 = F_listCreate(m)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L41
	} else {
		goto L155
	}
L155:
	;
	v950 = int32(0)
	v951 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v951)+60)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v951)+48)) = v948
	v955 = F_listCreate(m)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L41
	} else {
		goto L156
	}
L156:
	;
	v957 = int32(0)
	v958 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v958)+68)) = v957
	*(*int32)(unsafe.Add(mBase, uint32(v958)+56)) = v955
	v962 = F_listCreate(m)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L41
	} else {
		goto L157
	}
L157:
	;
	v964 = int32(0)
	v965 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v965)+76)) = v964
	*(*int32)(unsafe.Add(mBase, uint32(v965)+64)) = v962
	v969 = F_listCreate(m)
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L41
	} else {
		goto L158
	}
L158:
	;
	v971 = int32(0)
	v972 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v972)+84)) = v971
	*(*int32)(unsafe.Add(mBase, uint32(v972)+72)) = v969
	v976 = F_listCreate(m)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L41
	} else {
		goto L159
	}
L159:
	;
	v978 = int32(0)
	v979 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v979)+92)) = v978
	*(*int32)(unsafe.Add(mBase, uint32(v979)+80)) = v976
	v983 = F_listCreate(m)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L41
	} else {
		goto L160
	}
L160:
	;
	v985 = int32(0)
	v986 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v986)+100)) = v985
	*(*int32)(unsafe.Add(mBase, uint32(v986)+88)) = v983
	v990 = F_listCreate(m)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L41
	} else {
		goto L161
	}
L161:
	;
	v992 = int32(0)
	v993 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v993)+108)) = v992
	*(*int32)(unsafe.Add(mBase, uint32(v993)+96)) = v990
	v997 = F_listCreate(m)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L41
	} else {
		goto L162
	}
L162:
	;
	v999 = int32(0)
	v1000 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v1000)+116)) = v999
	*(*int32)(unsafe.Add(mBase, uint32(v1000)+104)) = v997
	v1004 = F_listCreate(m)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L41
	} else {
		goto L163
	}
L163:
	;
	v1006 = int32(0)
	v1007 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v1007)+124)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v1007)+112)) = v1004
	v1011 = F_listCreate(m)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L41
	} else {
		goto L164
	}
L164:
	;
	v1013 = int32(0)
	v1014 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v1014)+132)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v1014)+120)) = v1011
	v1018 = F_listCreate(m)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L41
	} else {
		goto L165
	}
L165:
	;
	v1020 = int32(0)
	v1021 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v1021)+140)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v1021)+128)) = v1018
	v1025 = F_listCreate(m)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L41
	} else {
		goto L166
	}
L166:
	;
	v1027 = int32(0)
	v1028 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v1028)+148)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v1028)+136)) = v1025
	v1032 = F_listCreate(m)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L41
	} else {
		goto L167
	}
L167:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	*(*int32)(unsafe.Add(mBase, uint32(v1035)+144)) = v1032
	goto L145
L168:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L172:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_serverBuildId(m *base.Module) int64 {
	var v4 int64
	_ = v4
	v4 = F_crc64(m, int64(0), int32(_a1188), int64(73))
	return v4
}
func F_serverCron(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v86 int64
	_ = v86
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v126 int64
	_ = v126
	var v131 int64
	_ = v131
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v173 int64
	_ = v173
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int64
	_ = v199
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v207 int64
	_ = v207
	var v212 int64
	_ = v212
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v243 int64
	_ = v243
	var v248 int64
	_ = v248
	var v252 int64
	_ = v252
	var v253 int64
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v279 int64
	_ = v279
	var v284 int64
	_ = v284
	var v288 int64
	_ = v288
	var v289 int64
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v301 int32
	_ = v301
	var v304 int64
	_ = v304
	var v306 int64
	_ = v306
	var v307 int32
	_ = v307
	var v312 int64
	_ = v312
	var v314 int64
	_ = v314
	var v317 int64
	_ = v317
	var v322 int64
	_ = v322
	var v324 int64
	_ = v324
	var v325 int64
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int64
	_ = v355
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
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
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v384 int64
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int64
	_ = v430
	var v432 int64
	_ = v432
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v496 int32
	_ = v496
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int64
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v542 int64
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int64
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v558 int64
	_ = v558
	var v563 int32
	_ = v563
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v768 int64
	_ = v768
	var v773 int32
	_ = v773
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int64
	_ = v789
	var v791 int32
	_ = v791
	var v792 int64
	_ = v792
	var v794 int64
	_ = v794
	var v796 int64
	_ = v796
	var v798 int32
	_ = v798
	var v799 int64
	_ = v799
	var v801 int64
	_ = v801
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int64
	_ = v812
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v828 int64
	_ = v828
	var v833 int64
	_ = v833
	var v838 int64
	_ = v838
	var v843 int64
	_ = v843
	var v848 int64
	_ = v848
	var v853 int64
	_ = v853
	var v858 int64
	_ = v858
	var v861 int64
	_ = v861
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v930 int64
	_ = v930
	var v932 int64
	_ = v932
	var v937 int64
	_ = v937
	var v938 int64
	_ = v938
	var v941 int64
	_ = v941
	var v942 int64
	_ = v942
	var v944 int64
	_ = v944
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1010 int32
	_ = v1010
	var v1016 int64
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1042 int32
	_ = v1042
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1136 int64
	_ = v1136
	var v1138 int64
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1150 int64
	_ = v1150
	var v1155 int64
	_ = v1155
	var v1160 int64
	_ = v1160
	var v1165 int64
	_ = v1165
	var v1170 int64
	_ = v1170
	var v1175 int64
	_ = v1175
	var v1180 int64
	_ = v1180
	var v1183 int64
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1269 int32
	_ = v1269
	var v1270 int64
	_ = v1270
	var v1273 int64
	_ = v1273
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1297 int32
	_ = v1297
	v4 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(144)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	if v21 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v50 = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, _consts[394]))
	if v51 == v50 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v24 = int32(0)
	v26 = m.G0
	v27 = int32(32)
	v28 = v26 - v27
	m.G0 = v28
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = int64(0)
	v34 = int32(1000)
	v35 = base.I32_div_s(v21, v34)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = base.I64_extend_i32_s(v35)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = (v21 - v35*v34) * v34
	v46 = F_setitimer(m, v24, v28, v24)
	mBase = m.M
	m.G0 = v28 + v27
	goto L3
L3:
	;
	goto L1
L4:
	;
	m.G0 = v18 + int32(144)
	v1297 = base.I32_div_s(int32(1000), v1282)
	return base.I64_extend_i32_s(v1297)
L5:
	;
	v56 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v58 = m.T0[v57].(func(*base.Module) int64)(m)
	mBase = m.M
	v61 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v62 = base.I32_div_s(int32(1000), v61)
	if int32(99) < v62 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v1282 = v55
	goto L4
L7:
	;
	F_cronUpdateMemoryStats(m)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L46
	} else {
		goto L47
	}
L8:
	;
	v72 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v74 = m.T0[v73].(func(*base.Module) int64)(m)
	mBase = m.M
	v76 = *(*int64)(unsafe.Add(mBase, _consts[87]))
	v78 = *(*int64)(unsafe.Add(mBase, _consts[799]))
	if v78 < int64(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	v69 = base.I32_div_s(int32(100), base.I32_extend16_s(v62))
	v71 = base.I32_rem_s(v66, base.I32_extend16_s(v69))
	if v71 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v107 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[800])) = v76
	*(*int64)(unsafe.Add(mBase, _consts[799])) = v74
	v112 = *(*int64)(unsafe.Add(mBase, _consts[303]))
	v114 = *(*int64)(unsafe.Add(mBase, _consts[305]))
	v117 = *(*int64)(unsafe.Add(mBase, _consts[296]))
	v120 = *(*int64)(unsafe.Add(mBase, _consts[302]))
	v121 = v112 + v114 + v117 + v120
	v123 = *(*int64)(unsafe.Add(mBase, _consts[813]))
	if v123 < int64(1) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v81 = v74 - v78
	if int64(1) <= v81 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v92 = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, _consts[801]))
	*(*int64)(unsafe.Add(mBase, uint32(v93<<(uint(int32(3))%32))+uint32(_consts[814]))) = v91
	v103 = base.I32_rem_s(v93+int32(1), int32(16))
	*(*int32)(unsafe.Add(mBase, _consts[801])) = v103
	goto L11
L14:
	;
	v86 = *(*int64)(unsafe.Add(mBase, _consts[800]))
	v90 = base.I64_div_s((v76-v86)*int64(1000000), v81)
	v91 = v90
	goto L13
L15:
	;
	v91 = int64(0)
	goto L13
L16:
	;
	v152 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[815])) = v121
	*(*int64)(unsafe.Add(mBase, _consts[813])) = v74
	v157 = *(*int64)(unsafe.Add(mBase, _consts[124]))
	v159 = *(*int64)(unsafe.Add(mBase, _consts[304]))
	v162 = *(*int64)(unsafe.Add(mBase, _consts[301]))
	v163 = v157 + v159 + v162
	v165 = *(*int64)(unsafe.Add(mBase, _consts[816]))
	if v165 < int64(1) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v126 = v74 - v123
	if int64(1) <= v126 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v137 = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, _consts[817]))
	*(*int64)(unsafe.Add(mBase, uint32(v138<<(uint(int32(3))%32))+uint32(_consts[818]))) = v136
	v148 = base.I32_rem_s(v138+int32(1), int32(16))
	*(*int32)(unsafe.Add(mBase, _consts[817])) = v148
	goto L16
L19:
	;
	v131 = *(*int64)(unsafe.Add(mBase, _consts[815]))
	v135 = base.I64_div_s((v121-v131)*int64(1000000), v126)
	v136 = v135
	goto L18
L20:
	;
	v136 = int64(0)
	goto L18
L21:
	;
	v194 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[819])) = v163
	*(*int64)(unsafe.Add(mBase, _consts[816])) = v74
	v199 = *(*int64)(unsafe.Add(mBase, _consts[296]))
	v201 = *(*int64)(unsafe.Add(mBase, _consts[303]))
	v202 = v199 + v201
	v204 = *(*int64)(unsafe.Add(mBase, _consts[820]))
	if v204 < int64(1) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v168 = v74 - v165
	if int64(1) <= v168 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v179 = int32(0)
	v180 = *(*int32)(unsafe.Add(mBase, _consts[821]))
	*(*int64)(unsafe.Add(mBase, uint32(v180<<(uint(int32(3))%32))+uint32(_consts[822]))) = v178
	v190 = base.I32_rem_s(v180+int32(1), int32(16))
	*(*int32)(unsafe.Add(mBase, _consts[821])) = v190
	goto L21
L24:
	;
	v173 = *(*int64)(unsafe.Add(mBase, _consts[819]))
	v177 = base.I64_div_s((v163-v173)*int64(1000000), v168)
	v178 = v177
	goto L23
L25:
	;
	v178 = int64(0)
	goto L23
L26:
	;
	v233 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[823])) = v202
	*(*int64)(unsafe.Add(mBase, _consts[820])) = v74
	v238 = *(*int64)(unsafe.Add(mBase, _consts[124]))
	v240 = *(*int64)(unsafe.Add(mBase, _consts[824]))
	if v240 < int64(1) {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v207 = v74 - v204
	if int64(1) <= v207 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v218 = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, _consts[825]))
	*(*int64)(unsafe.Add(mBase, uint32(v219<<(uint(int32(3))%32))+uint32(_consts[826]))) = v217
	v229 = base.I32_rem_s(v219+int32(1), int32(16))
	*(*int32)(unsafe.Add(mBase, _consts[825])) = v229
	goto L26
L29:
	;
	v212 = *(*int64)(unsafe.Add(mBase, _consts[823]))
	v216 = base.I64_div_s((v202-v212)*int64(1000000), v207)
	v217 = v216
	goto L28
L30:
	;
	v217 = int64(0)
	goto L28
L31:
	;
	v269 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[827])) = v238
	*(*int64)(unsafe.Add(mBase, _consts[824])) = v74
	v274 = *(*int64)(unsafe.Add(mBase, _consts[828]))
	v276 = *(*int64)(unsafe.Add(mBase, _consts[829]))
	if v276 < int64(1) {
		v306 = v274
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v243 = v74 - v240
	if int64(1) <= v243 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v254 = int32(0)
	v255 = *(*int32)(unsafe.Add(mBase, _consts[830]))
	*(*int64)(unsafe.Add(mBase, uint32(v255<<(uint(int32(3))%32))+uint32(_consts[831]))) = v253
	v265 = base.I32_rem_s(v255+int32(1), int32(16))
	*(*int32)(unsafe.Add(mBase, _consts[830])) = v265
	goto L31
L34:
	;
	v248 = *(*int64)(unsafe.Add(mBase, _consts[827]))
	v252 = base.I64_div_s((v238-v248)*int64(1000000), v243)
	v253 = v252
	goto L33
L35:
	;
	v253 = int64(0)
	goto L33
L36:
	;
	v307 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[832])) = v274
	*(*int64)(unsafe.Add(mBase, _consts[829])) = v74
	v312 = *(*int64)(unsafe.Add(mBase, _consts[833]))
	v314 = *(*int64)(unsafe.Add(mBase, _consts[834]))
	if v314 < int64(1) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v279 = v74 - v276
	if int64(1) <= v279 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v290 = int32(0)
	v291 = *(*int32)(unsafe.Add(mBase, _consts[835]))
	*(*int64)(unsafe.Add(mBase, uint32(v291<<(uint(int32(3))%32))+uint32(_consts[836]))) = v289
	v301 = base.I32_rem_s(v291+int32(1), int32(16))
	*(*int32)(unsafe.Add(mBase, _consts[835])) = v301
	v304 = *(*int64)(unsafe.Add(mBase, _consts[828]))
	v306 = v304
	goto L36
L39:
	;
	v284 = *(*int64)(unsafe.Add(mBase, _consts[832]))
	v288 = base.I64_div_s((v274-v284)*int64(1000000), v279)
	v289 = v288
	goto L38
L40:
	;
	v289 = int64(0)
	goto L38
L41:
	;
	v341 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[837])) = v312
	*(*int64)(unsafe.Add(mBase, _consts[834])) = v306
	goto L7
L42:
	;
	v317 = v306 - v314
	if int64(1) <= v317 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v326 = int32(0)
	v327 = *(*int32)(unsafe.Add(mBase, _consts[838]))
	*(*int64)(unsafe.Add(mBase, uint32(v327<<(uint(int32(3))%32))+uint32(_consts[839]))) = v325
	v337 = base.I32_rem_s(v327+int32(1), int32(16))
	*(*int32)(unsafe.Add(mBase, _consts[838])) = v337
	goto L41
L44:
	;
	v322 = *(*int64)(unsafe.Add(mBase, _consts[837]))
	v324 = base.I64_div_s(v312-v322, v317)
	v325 = v324
	goto L43
L45:
	;
	v325 = int64(0)
	goto L43
L46:
	;
	return int64(0)
L47:
	;
	v354 = int32(0)
	v355 = *(*int64)(unsafe.Add(mBase, _consts[840]))
	v357 = *(*int32)(unsafe.Add(mBase, _consts[841]))
	if v357 == v354 {
		goto L52
	} else {
		goto L53
	}
L48:
	;
	v993 = int32(2)
	v994 = int32(0)
	v995 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	v1000 = *(*int32)(unsafe.Add(mBase, _consts[399]))
	if v1000 != 0 {
		goto L186
	} else {
		goto L187
	}
L49:
	;
	v964 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v965 = base.I32_div_s(int32(1000), v964)
	if int32(999) < v965 {
		goto L182
	} else {
		goto L183
	}
L50:
	;
	m.Env.Exit(m, int32(0))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v469 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(1) < v469 {
		goto L80
	} else {
		goto L81
	}
L52:
	;
	if v355 == int64(0) {
		goto L51
	} else {
		goto L66
	}
L53:
	;
	if v355 != int64(0) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v362 = int32(0)
	v364 = *(*int32)(unsafe.Add(mBase, _consts[842]))
	v366 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	v369 = *(*int32)(unsafe.Add(mBase, _consts[844]))
	if v369 == int32(15) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v372 = v366
	goto L57
L56:
	;
	v372 = v362
	goto L57
L57:
	;
	if v364 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v373 = v364
	goto L60
L59:
	;
	v373 = v372
	goto L60
L60:
	;
	if v369 == int32(2) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v376 = v373
	goto L63
L62:
	;
	v376 = v372
	goto L63
L63:
	;
	v377 = F_prepareForShutdown(m, v362, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L46
	} else {
		goto L64
	}
L64:
	;
	if v377 != 0 {
		goto L51
	} else {
		goto L65
	}
L65:
	;
	m.Env.Exit(m, int32(0))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v384 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	if v355 <= v384 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v449 = F_finishShutdown(m)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L46
	} else {
		goto L78
	}
L68:
	;
	v386 = int32(0)
	v387 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)+20))
	if v388 == v386 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v392 = v18 + int32(80)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	*(*int32)(unsafe.Add(mBase, uint32(v392)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v392))) = v393
	goto L70
L70:
	;
	goto L71
L71:
	;
	v413 = v18 + int32(80)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	if v415 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v415 == int32(0) {
		goto L67
	} else {
		goto L76
	}
L74:
	;
	goto L73
L75:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v413)+4))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v415+base.B2i32(v418 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v413))) = v424
	goto L74
L76:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v415)+8))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+104))
	v430 = *(*int64)(unsafe.Add(mBase, uint32(v429)+64))
	v432 = *(*int64)(unsafe.Add(mBase, _consts[40]))
	if v430 == v432 {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	goto L51
L78:
	;
	if v449 == int32(0) {
		goto L50
	} else {
		goto L79
	}
L79:
	;
	goto L51
L80:
	;
	v603 = *(*int32)(unsafe.Add(mBase, _consts[232]))
	if v603 != 0 {
		goto L108
	} else {
		goto L109
	}
L81:
	;
	v472 = int32(0)
	v473 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	v477 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v478 = base.I32_div_s(int32(1000), v477)
	v480 = base.I32_div_s(int32(5000), base.I32_extend16_s(v478))
	v482 = base.I32_rem_s(v473, base.I32_extend16_s(v480))
	if v482 != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v484 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v484 < int32(1) {
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v496 = int32(0)
	goto L84
L84:
	;
	v507 = int32(0)
	v508 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v510 = v496 << (uint(int32(2)) % 32)
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v508+v510)))
	if v512 == v507 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L80
L86:
	;
	v583 = v496 + int32(1)
	v585 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v583 < v585 {
		v496 = v583
		goto L84
	} else {
		goto L107
	}
L87:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v515)+12))
	if v516 == int32(1) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L93
L89:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v515)+8))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	if v521 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v515)+48))
	v525 = v519
	goto L88
L91:
	;
	v523 = F_hashtableBuckets(m, v521)
	mBase = m.M
	v525 = v523
	goto L88
L92:
	;
	v525 = int32(0)
	goto L88
L93:
	;
	v528 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v528+v510)))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v531)+12))
	if v532 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v544 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v544+v510)))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)+4))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	if v548 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v531)+8))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	if v537 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v535 = *(*int64)(unsafe.Add(mBase, uint32(v531)+40))
	v542 = v535
	goto L94
L97:
	;
	v539 = F_hashtableSize(m, v537)
	mBase = m.M
	v542 = base.I64_extend_i32_u(v539)
	goto L94
L98:
	;
	v542 = int64(0)
	goto L94
L99:
	;
	if v542|v558 == int64(0) {
		goto L86
	} else {
		goto L104
	}
L100:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v547)+8))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	if v553 != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v551 = *(*int64)(unsafe.Add(mBase, uint32(v547)+40))
	v558 = v551
	goto L99
L102:
	;
	v555 = F_hashtableSize(m, v553)
	mBase = m.M
	v558 = base.I64_extend_i32_u(v555)
	goto L99
L103:
	;
	v558 = int64(0)
	goto L99
L104:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(1) < v563 {
		goto L86
	} else {
		goto L105
	}
L105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(64)))) = v558
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(72)))) = base.I64_extend_i32_u(int32(12) * v525)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v496
	F__serverLog(m, int32(1), int32(_a1569), v18+int32(48))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L46
	} else {
		goto L106
	}
L106:
	;
	goto L86
L107:
	;
	goto L85
L108:
	;
	F_databasesCron(m)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L46
	} else {
		goto L130
	}
L109:
	;
	v604 = int32(0)
	v605 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	v609 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v610 = base.I32_div_s(int32(1000), v609)
	v612 = base.I32_div_s(int32(5000), base.I32_extend16_s(v610))
	v614 = base.I32_rem_s(v605, base.I32_extend16_s(v612))
	if v614 != 0 {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v618 = int32(0)
	v627 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	if v627 < int32(261) {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	F_bytesToHuman(m, v18+int32(80), int32(64), base.I64_extend_i32_u(v704))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L46
	} else {
		goto L127
	}
L112:
	;
	goto L111
L113:
	;
	v638 = v636 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v636) {
		goto L118
	} else {
		goto L119
	}
L114:
	;
	if v627 < int32(1) {
		v704 = v618
		goto L112
	} else {
		goto L116
	}
L115:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	v635 = v631
	v636 = int32(260)
	goto L113
L116:
	;
	v635 = v618
	v636 = v627
	goto L113
L117:
	;
	if v638 == int32(0) {
		v704 = v677
		goto L112
	} else {
		goto L123
	}
L118:
	;
	v645 = int32(0)
	v647 = v635
	v648 = v645
	v652 = v645
	goto L120
L119:
	;
	v677 = v635
	v678 = int32(0)
	goto L117
L120:
	;
	v655 = v648 << (uint(int32(2)) % 32)
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v655)+uint32(_consts[282])))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v655)+uint32(_consts[283])))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v655)+uint32(_consts[284])))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v655)+uint32(_consts[285])))
	v671 = v658 + (v661 + (v664 + (v667 + v647)))
	v672 = int32(4)
	v673 = v648 + v672
	v675 = v652 + v672
	if v675 != v636&int32(2147483644) {
		v647 = v671
		v648 = v673
		v652 = v675
		goto L120
	} else {
		goto L122
	}
L121:
	;
	v677 = v671
	v678 = v673
	goto L117
L122:
	;
	goto L121
L123:
	;
	v686 = v677
	v687 = v678
	v689 = int32(0)
	goto L124
L124:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v687<<(uint(int32(2))%32))+uint32(_consts[285])))
	v698 = v697 + v686
	v699 = int32(1)
	v702 = v689 + v699
	if v702 != v638 {
		v686 = v698
		v687 = v687 + v699
		v689 = v702
		goto L124
	} else {
		goto L126
	}
L125:
	;
	v704 = v698
	goto L112
L126:
	;
	goto L125
L127:
	;
	v714 = int32(0)
	v715 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v714 < v715 {
		goto L108
	} else {
		goto L128
	}
L128:
	;
	v718 = int32(0)
	v719 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)+20))
	v722 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v720 - v723
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v18 + int32(80)
	F__serverLog(m, v718, int32(_a1570), v18+int32(32))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L46
	} else {
		goto L129
	}
L129:
	;
	goto L108
L130:
	;
	v743 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	if v743 != int32(-1) {
		v756 = v743
		goto L131
	} else {
		goto L132
	}
L131:
	;
	if v756 != int32(-1) {
		goto L49
	} else {
		goto L138
	}
L132:
	;
	v746 = int32(0)
	v747 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	if v747 == v746 {
		v756 = v743
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v750 = F_aofRewriteLimited(m)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L46
	} else {
		goto L135
	}
L134:
	;
	v755 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	v756 = v755
	goto L131
L135:
	;
	if v750 != 0 {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v752 = F_rewriteAppendOnlyFileBackground(m)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L46
	} else {
		goto L137
	}
L137:
	;
	goto L134
L138:
	;
	v760 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v760)+20))
	goto L139
L139:
	;
	if v761 != 0 {
		goto L49
	} else {
		goto L140
	}
L140:
	;
	v762 = int32(0)
	v764 = *(*int32)(unsafe.Add(mBase, _consts[229]))
	if v764 <= v762 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v918 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v918 != int32(1) {
		goto L48
	} else {
		goto L167
	}
L142:
	;
	v768 = *(*int64)(unsafe.Add(mBase, _consts[83]))
	v773 = v762
	goto L143
L143:
	;
	v785 = *(*int32)(unsafe.Add(mBase, _consts[228]))
	v788 = v785 + v773<<(uint(int32(4))%32)
	v789 = int64(*(*int32)(unsafe.Add(mBase, uint32(v788)+8)))
	if v768 < v789 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	goto L141
L145:
	;
	v898 = v773 + int32(1)
	v900 = *(*int32)(unsafe.Add(mBase, _consts[229]))
	if v898 < v900 {
		v773 = v898
		goto L143
	} else {
		goto L166
	}
L146:
	;
	v791 = int32(0)
	v792 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v794 = *(*int64)(unsafe.Add(mBase, _consts[845]))
	v796 = *(*int64)(unsafe.Add(mBase, uint32(v788)))
	if v792-v794 <= v796 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v798 = int32(0)
	v799 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v801 = *(*int64)(unsafe.Add(mBase, _consts[846]))
	if int64(5) < v799-v801 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v808 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v808 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v806 = *(*int32)(unsafe.Add(mBase, _consts[847]))
	if v806 != 0 {
		goto L145
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v823 = v18 + int32(80)
	v824 = int32(0)
	v828 = *(*int64)(unsafe.Add(mBase, _consts[361]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(136)))) = v828
	v833 = *(*int64)(unsafe.Add(mBase, _consts[362]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(128)))) = v833
	v838 = *(*int64)(unsafe.Add(mBase, _consts[363]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(120)))) = v838
	v843 = *(*int64)(unsafe.Add(mBase, _consts[364]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(112)))) = v843
	v848 = *(*int64)(unsafe.Add(mBase, _consts[365]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(104)))) = v848
	v853 = *(*int64)(unsafe.Add(mBase, _consts[366]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(96)))) = v853
	v858 = *(*int64)(unsafe.Add(mBase, _consts[367]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(88)))) = v858
	v861 = *(*int64)(unsafe.Add(mBase, _consts[368]))
	*(*int64)(unsafe.Add(mBase, uint32(v823))) = v861
	v864 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v864 != 0 {
		goto L155
	} else {
		goto L156
	}
L152:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v788)+8))
	v812 = *(*int64)(unsafe.Add(mBase, uint32(v788)))
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+20)) = uint32(v812)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v811
	F__serverLog(m, int32(2), int32(_a1571), v18+int32(16))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L46
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	v891 = int32(0)
	v893 = *(*int32)(unsafe.Add(mBase, _consts[369]))
	v895 = F_rdbSaveBackground(m, v891, v893, v890, v891)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L46
	} else {
		goto L165
	}
L155:
	;
	v877 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	if v877 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L156:
	;
	v866 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	if v866 == int32(0) {
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v871 = *(*int32)(unsafe.Add(mBase, _consts[371]))
	if v871 == int32(-1) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v874 = int32(0)
	goto L160
L159:
	;
	v874 = v871
	goto L160
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v823))) = v874
	v890 = v823
	goto L154
L161:
	;
	v884 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	if v884 != 0 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v877)+96))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v880)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v823))) = v881
	v890 = v823
	goto L154
L163:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v884)+96))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v886)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v823))) = v887
	v890 = v823
	goto L154
L164:
	;
	v890 = int32(0)
	goto L154
L165:
	;
	goto L141
L166:
	;
	goto L144
L167:
	;
	v922 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	if v922 != int32(-1) {
		goto L48
	} else {
		goto L168
	}
L168:
	;
	v925 = int32(0)
	v926 = *(*int32)(unsafe.Add(mBase, _consts[848]))
	if v926 == v925 {
		goto L48
	} else {
		goto L169
	}
L169:
	;
	v929 = int32(0)
	v930 = *(*int64)(unsafe.Add(mBase, _consts[50]))
	v932 = *(*int64)(unsafe.Add(mBase, _consts[849]))
	if v930 <= v932 {
		goto L48
	} else {
		goto L170
	}
L170:
	;
	v937 = *(*int64)(unsafe.Add(mBase, _consts[75]))
	v938 = int64(1)
	if base.Ui64(v938) < base.Ui64(v937) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v941 = v937
	goto L173
L172:
	;
	v941 = v938
	goto L173
L173:
	;
	v942 = base.I64_div_s(v930*int64(100), v941)
	v944 = v942 + int64(-100)
	if v944 < base.I64_extend_i32_s(v926) {
		goto L48
	} else {
		goto L174
	}
L174:
	;
	v947 = F_aofRewriteLimited(m)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L46
	} else {
		goto L175
	}
L175:
	;
	if v947 != 0 {
		goto L48
	} else {
		goto L176
	}
L176:
	;
	v950 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v950 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v958 = F_rewriteAppendOnlyFileBackground(m)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L46
	} else {
		goto L180
	}
L178:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v944
	F__serverLog(m, int32(2), int32(_a1572), v18)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L46
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	goto L48
L181:
	;
	F_checkChildrenDone(m)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L46
	} else {
		goto L185
	}
L182:
	;
	F_receiveChildInfo(m)
	mBase = m.M
	goto L181
L183:
	;
	v969 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	v972 = base.I32_div_s(int32(1000), base.I32_extend16_s(v965))
	v974 = base.I32_rem_s(v969, base.I32_extend16_s(v972))
	if v974 != 0 {
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	goto L48
L186:
	;
	v1001 = base.B2i32(v995 != int32(-1))
	goto L188
L187:
	;
	v1001 = v993
	goto L188
L188:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	if v1003 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1004 = v993
	goto L191
L190:
	;
	v1004 = v1001
	goto L191
L191:
	;
	*(*int32)(unsafe.Add(mBase, _consts[798])) = v1004
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, _consts[427])) = v1004
	goto L193
L193:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if base.Ui32(int32(1)) < base.Ui32(v1010+int32(-1)) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v1025 = base.I32_div_s(int32(1000), v1024)
	if int32(999) < v1025 {
		goto L199
	} else {
		goto L200
	}
L195:
	;
	v1016 = *(*int64)(unsafe.Add(mBase, _consts[42]))
	if v1016 == int64(0) {
		goto L194
	} else {
		goto L196
	}
L196:
	;
	F_flushAppendOnlyFile(m, int32(0))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L46
	} else {
		goto L197
	}
L197:
	;
	goto L194
L198:
	;
	F_updatePausedActions(m)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L46
	} else {
		goto L205
	}
L199:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if base.Ui32(int32(1)) < base.Ui32(v1036+int32(-1)) {
		goto L198
	} else {
		goto L202
	}
L200:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	v1032 = base.I32_div_s(int32(1000), base.I32_extend16_s(v1025))
	v1034 = base.I32_rem_s(v1029, base.I32_extend16_s(v1032))
	if v1034 != 0 {
		goto L198
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	if v1042 != int32(-1) {
		goto L198
	} else {
		goto L203
	}
L203:
	;
	F_flushAppendOnlyFile(m, int32(0))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L46
	} else {
		goto L204
	}
L204:
	;
	goto L198
L205:
	;
	v1051 = int32(0)
	v1052 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v1053 = base.I32_div_s(int32(1000), v1052)
	v1055 = *(*int32)(unsafe.Add(mBase, _consts[658]))
	if v1055 == v1051 {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	v1080 = int32(0)
	v1081 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v1081 == v1080 {
		goto L215
	} else {
		goto L216
	}
L207:
	;
	F_replicationCron(m)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L46
	} else {
		goto L214
	}
L208:
	;
	if int32(999) < v1053 {
		goto L207
	} else {
		goto L212
	}
L209:
	;
	if int32(99) < v1053 {
		goto L207
	} else {
		goto L210
	}
L210:
	;
	v1060 = int32(0)
	v1061 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	v1064 = base.I32_div_s(int32(100), base.I32_extend16_s(v1053))
	v1066 = base.I32_rem_s(v1061, base.I32_extend16_s(v1064))
	if v1066 == v1060 {
		goto L207
	} else {
		goto L211
	}
L211:
	;
	goto L206
L212:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	v1075 = base.I32_div_s(int32(1000), base.I32_extend16_s(v1053))
	v1077 = base.I32_rem_s(v1072, base.I32_extend16_s(v1075))
	if v1077 != 0 {
		goto L206
	} else {
		goto L213
	}
L213:
	;
	goto L207
L214:
	;
	goto L206
L215:
	;
	v1100 = int32(0)
	v1101 = *(*int32)(unsafe.Add(mBase, _consts[232]))
	if v1101 == v1100 {
		goto L221
	} else {
		goto L222
	}
L216:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v1087 = base.I32_div_s(int32(1000), v1086)
	if int32(99) < v1087 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	F_clusterCron(m)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L46
	} else {
		goto L220
	}
L218:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	v1094 = base.I32_div_s(int32(100), base.I32_extend16_s(v1087))
	v1096 = base.I32_rem_s(v1091, base.I32_extend16_s(v1094))
	if v1096 != 0 {
		goto L215
	} else {
		goto L219
	}
L219:
	;
	goto L217
L220:
	;
	goto L215
L221:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v1109 = base.I32_div_s(int32(1000), v1108)
	if int32(999) < v1109 {
		goto L225
	} else {
		goto L226
	}
L222:
	;
	F_sentinelTimer(m)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L46
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	v1121 = int32(0)
	v1122 = *(*int32)(unsafe.Add(mBase, _consts[850]))
	if v1122 == v1121 {
		goto L229
	} else {
		goto L230
	}
L225:
	;
	F_migrateCloseTimedoutSockets(m)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L46
	} else {
		goto L228
	}
L226:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	v1116 = base.I32_div_s(int32(1000), base.I32_extend16_s(v1109))
	v1118 = base.I32_rem_s(v1113, base.I32_extend16_s(v1116))
	if v1118 != 0 {
		goto L224
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	goto L224
L229:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	if v1128 != int32(-1) {
		goto L232
	} else {
		goto L233
	}
L230:
	;
	F_trackingLimitUsedSlots(m)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L46
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, _consts[541]))
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+16))
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+12))
	goto L251
L233:
	;
	v1131 = int32(0)
	v1132 = *(*int32)(unsafe.Add(mBase, _consts[851]))
	if v1132 == v1131 {
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v1135 = int32(0)
	v1136 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v1138 = *(*int64)(unsafe.Add(mBase, _consts[846]))
	if int64(5) < v1136-v1138 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v1145 = v18 + int32(80)
	v1146 = int32(0)
	v1150 = *(*int64)(unsafe.Add(mBase, _consts[361]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(136)))) = v1150
	v1155 = *(*int64)(unsafe.Add(mBase, _consts[362]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(128)))) = v1155
	v1160 = *(*int64)(unsafe.Add(mBase, _consts[363]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(120)))) = v1160
	v1165 = *(*int64)(unsafe.Add(mBase, _consts[364]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(112)))) = v1165
	v1170 = *(*int64)(unsafe.Add(mBase, _consts[365]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(104)))) = v1170
	v1175 = *(*int64)(unsafe.Add(mBase, _consts[366]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(96)))) = v1175
	v1180 = *(*int64)(unsafe.Add(mBase, _consts[367]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(88)))) = v1180
	v1183 = *(*int64)(unsafe.Add(mBase, _consts[368]))
	*(*int64)(unsafe.Add(mBase, uint32(v1145))) = v1183
	v1186 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v1186 != 0 {
		goto L239
	} else {
		goto L240
	}
L236:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, _consts[847]))
	if v1143 != 0 {
		goto L232
	} else {
		goto L237
	}
L237:
	;
	goto L235
L238:
	;
	v1213 = int32(0)
	v1215 = *(*int32)(unsafe.Add(mBase, _consts[369]))
	v1217 = F_rdbSaveBackground(m, v1213, v1215, v1212, v1213)
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L46
	} else {
		goto L249
	}
L239:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	if v1199 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L240:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	if v1188 == int32(0) {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, _consts[371]))
	if v1193 == int32(-1) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1196 = int32(0)
	goto L244
L243:
	;
	v1196 = v1193
	goto L244
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1145))) = v1196
	v1212 = v1145
	goto L238
L245:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	if v1206 != 0 {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+96))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1202)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1145))) = v1203
	v1212 = v1145
	goto L238
L247:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+96))
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1145))) = v1209
	v1212 = v1145
	goto L238
L248:
	;
	v1212 = int32(0)
	goto L238
L249:
	;
	if v1217 != 0 {
		goto L232
	} else {
		goto L250
	}
L250:
	;
	v1219 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[851])) = v1219
	goto L232
L251:
	;
	v1229 = int32(0)
	v1230 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	if v1226+v1227 == v1229 {
		v1248 = v1230
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1250 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v1250
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v1248
	*(*int64)(unsafe.Add(mBase, uint32(v18)+80)) = int64(1)
	F_moduleFireServerEvent(m, int64(8), v1250, v18+int32(80))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L46
	} else {
		goto L258
	}
L253:
	;
	v1234 = base.I32_div_s(int32(1000), v1230)
	if int32(99) < v1234 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	F_modulesCron(m)
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L46
	} else {
		goto L257
	}
L255:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	v1241 = base.I32_div_s(int32(100), base.I32_extend16_s(v1234))
	v1243 = base.I32_rem_s(v1238, base.I32_extend16_s(v1241))
	if v1243 != 0 {
		v1248 = v1230
		goto L252
	} else {
		goto L256
	}
L256:
	;
	goto L254
L257:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v1248 = v1247
	goto L252
L258:
	;
	v1261 = int32(0)
	v1263 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	*(*int32)(unsafe.Add(mBase, _consts[220])) = v1263 + int32(1)
	v1269 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v1270 = m.T0[v1269].(func(*base.Module) int64)(m)
	mBase = m.M
	v1273 = *(*int64)(unsafe.Add(mBase, _consts[804]))
	*(*int64)(unsafe.Add(mBase, _consts[804])) = v1270 - v58 + v1273
	v1277 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v1282 = v1277
	goto L4
}
func F_serverLogFromHandler(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v176 int64
	_ = v176
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v185 int64
	_ = v185
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int64
	_ = v204
	var v205 int64
	_ = v205
	var v206 int32
	_ = v206
	var v208 int64
	_ = v208
	var v218 int32
	_ = v218
	var v225 int64
	_ = v225
	var v229 int32
	_ = v229
	var v230 int64
	_ = v230
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int64
	_ = v252
	var v255 int64
	_ = v255
	var v263 int64
	_ = v263
	var v265 int32
	_ = v265
	var v266 int64
	_ = v266
	var v276 int32
	_ = v276
	var v283 int64
	_ = v283
	var v286 int64
	_ = v286
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v366 int32
	_ = v366
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v391 int64
	_ = v391
	var v392 int32
	_ = v392
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int64
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	v5 = m.G0
	v7 = v5 - int32(1040)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+1036)) = l2
	v28 = v7 + int32(1023)
	v29 = m.G0
	v31 = v29 - int32(32)
	v33 = v31 | int32(5)
	v35 = v31 + int32(20)
	v39 = v7
	v40 = l1
	v41 = l2
	goto L2
L1:
	;
	v466 = m.G0
	v468 = v466 - int32(80)
	m.G0 = v468
	v473 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if l0&int32(255) < v473 {
		goto L95
	} else {
		goto L96
	}
L2:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v57 == int32(37) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v72 = int32(1)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v73 == int32(108) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	if v57 == int32(0) {
		v63 = v39
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v57)
	v68 = int32(1)
	v39 = v39 + v68
	v40 = v40 + v68
	goto L2
L7:
	;
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v64)
	goto L1
L8:
	;
	if v39 != v28 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v63 = v28
	goto L7
L10:
	;
	v91 = v88 & int32(255)
	v93 = v91 + int32(-100)
	if base.Ui32(int32(20)) < base.Ui32(v93) {
		v443 = v39
		v445 = v41
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)))
	if v78 == int32(108) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v87 = v40 + int32(1)
	v88 = v73
	v89 = v72
	goto L10
L13:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+3)))
	v87 = v40 + int32(3)
	v88 = v83
	v89 = int32(0)
	goto L10
L14:
	;
	v87 = v40 + int32(2)
	v88 = v78
	v89 = v72
	goto L10
L15:
	;
	v39 = v443
	v40 = v87 + int32(1)
	v41 = v445
	goto L2
L16:
	;
	if int32(1)<<(uint(v93)%32)&int32(1179681) != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v403 == int32(120) {
		goto L82
	} else {
		goto L83
	}
L18:
	;
	v241 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+21)) = uint8(v241)
	v246 = base.B2i32(v91 == int32(120))
	v248 = base.B2i32(v189 < int64(0)) & (v246 | v194)
	if v248 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L19:
	;
	v206 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+21)) = uint8(v206)
	v208 = base.I64_extend_i32_u(v200)
	v218 = v31 + int32(21)
	v225 = v205
	goto L45
L20:
	;
	v194 = base.B2i32(v91 == int32(112))
	if v91 == int32(112) {
		goto L41
	} else {
		goto L42
	}
L21:
	;
	v185 = int64(*(*int32)(unsafe.Add(mBase, uint32(v41))))
	v188 = v41 + int32(4)
	v189 = v185
	v190 = int32(1)
	goto L20
L22:
	;
	if v89 != 0 {
		goto L37
	} else {
		goto L38
	}
L23:
	;
	if v93 == int32(12) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	if v93 != int32(15) {
		v443 = v39
		v445 = v41
		goto L15
	} else {
		goto L25
	}
L25:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v106 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v39 = v142
	v40 = v87 + int32(1)
	v41 = v41 + int32(4)
	goto L2
L27:
	;
	v108 = v106
	goto L29
L28:
	;
	v108 = int32(_a939)
	goto L29
L29:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v109 == int32(0) {
		v142 = v39
		goto L26
	} else {
		goto L30
	}
L30:
	;
	if base.Ui32(v28) <= base.Ui32(v39) {
		v142 = v39
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v114 = v39
	v122 = v108
	v124 = v109
	goto L32
L32:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v124)
	v134 = v114 + int32(1)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
	if v135 == int32(0) {
		v142 = v134
		goto L26
	} else {
		goto L34
	}
L33:
	;
	v142 = v134
	goto L26
L34:
	;
	if base.Ui32(v134) < base.Ui32(v28) {
		v114 = v134
		v122 = v122 + int32(1)
		v124 = v135
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v199 = v89
	v200 = int32(10)
	v201 = v178
	v204 = int64(0)
	v205 = v179
	goto L19
L37:
	;
	if v91 != int32(117) {
		goto L21
	} else {
		goto L40
	}
L38:
	;
	v165 = (v41 + int32(7)) & int32(-8)
	v167 = v165 + int32(8)
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v165)))
	if v91 == int32(117) {
		v178 = v167
		v179 = v168
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v188 = v167
	v189 = v168
	v190 = int32(0)
	goto L20
L40:
	;
	v176 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v41))))
	v178 = v41 + int32(4)
	v179 = v176
	goto L36
L41:
	;
	v195 = int32(16)
	goto L43
L42:
	;
	v195 = int32(10)
	goto L43
L43:
	;
	if v91 != int32(117) {
		goto L18
	} else {
		goto L44
	}
L44:
	;
	v199 = v190
	v200 = v195
	v201 = v188
	v204 = v189
	v205 = int64(0)
	goto L19
L45:
	;
	v229 = v218 + int32(-1)
	v230 = base.I64_div_u_s(v225, v208)
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v225-v230*v208))+uint32(_consts[797]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v229))) = uint8(v236)
	if base.B2i32(base.Ui64(v225) < base.Ui64(v208)) == int32(0) {
		v218 = v229
		v225 = v230
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v381 = v201
	v387 = v229
	v391 = v204
	v392 = v199
	goto L17
L47:
	;
	goto L46
L48:
	;
	if v91 == int32(120) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v255 = int64(3472328296227680304)
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v255
	*(*int64)(unsafe.Add(mBase, uint32(v33+int32(8)))) = v255
	v263 = v189 ^ int64(-1)
	goto L48
L50:
	;
	v252 = v189 >> (uint(int64(63)) % 64)
	v263 = v189 ^ v252 - v252
	goto L48
L51:
	;
	v265 = int32(16)
	goto L53
L52:
	;
	v265 = v195
	goto L53
L53:
	;
	v266 = base.I64_extend_i32_u(v265)
	v276 = v35
	v283 = v263
	goto L54
L54:
	;
	v286 = base.I64_div_s(v283, v266)
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v283-v286*v266))+uint32(_consts[797]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v276))) = uint8(v292)
	v295 = v276 + int32(-1)
	if v286 != int64(0) {
		v276 = v295
		v283 = v286
		goto L54
	} else {
		goto L56
	}
L55:
	;
	if int64(-1) < v189 {
		v306 = v295
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	v307 = int32(0)
	if v248 == v307 {
		v366 = v306
		goto L60
	} else {
		goto L61
	}
L58:
	;
	if v265 != int32(10) {
		v306 = v295
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v302 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v295))) = uint8(v302)
	v306 = v276 + int32(-2)
	goto L57
L60:
	;
	v381 = v188
	v387 = v366 + int32(1)
	v391 = v189
	v392 = v190
	goto L17
L61:
	;
	v320 = v307
	v321 = v35
	goto L62
L62:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	switch v330 + int32(-48) {
	case 0:
		v348 = int32(102)
		goto L65
	case 1:
		goto L80
	case 2:
		goto L79
	case 3:
		goto L78
	case 4:
		goto L77
	case 5:
		goto L76
	case 6:
		goto L75
	case 7:
		goto L74
	case 8:
		goto L73
	case 9:
		goto L72
	default:
		goto L64
	case 49:
		goto L71
	case 50:
		goto L70
	case 51:
		goto L69
	case 52:
		goto L68
	case 53:
		goto L67
	case 54:
		goto L66
	}
L63:
	;
	v366 = v352
	goto L60
L64:
	;
	v352 = v321 + int32(-1)
	v354 = v320 + int32(1)
	if v354 != int32(16) {
		v320 = v354
		v321 = v352
		goto L62
	} else {
		goto L81
	}
L65:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v321))) = uint8(v348)
	goto L64
L66:
	;
	v348 = int32(48)
	goto L65
L67:
	;
	v348 = int32(49)
	goto L65
L68:
	;
	v348 = int32(50)
	goto L65
L69:
	;
	v348 = int32(51)
	goto L65
L70:
	;
	v348 = int32(52)
	goto L65
L71:
	;
	v348 = int32(53)
	goto L65
L72:
	;
	v348 = int32(54)
	goto L65
L73:
	;
	v348 = int32(55)
	goto L65
L74:
	;
	v348 = int32(56)
	goto L65
L75:
	;
	v348 = int32(57)
	goto L65
L76:
	;
	v348 = int32(97)
	goto L65
L77:
	;
	v348 = int32(98)
	goto L65
L78:
	;
	v348 = int32(99)
	goto L65
L79:
	;
	v348 = int32(100)
	goto L65
L80:
	;
	v348 = int32(101)
	goto L65
L81:
	;
	goto L63
L82:
	;
	v406 = base.I32_wrap_i64(int64(base.Ui64(v391)>>(uint(int64(60))%64))) & int32(8)
	goto L84
L83:
	;
	v406 = int32(0)
	goto L84
L84:
	;
	if v392 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v408 = v406
	goto L87
L86:
	;
	v408 = int32(0)
	goto L87
L87:
	;
	v409 = v387 + v408
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	if v410 == int32(0) {
		v443 = v39
		v445 = v381
		goto L15
	} else {
		goto L88
	}
L88:
	;
	if base.Ui32(v28) <= base.Ui32(v39) {
		v443 = v39
		v445 = v381
		goto L15
	} else {
		goto L89
	}
L89:
	;
	v415 = v39
	v423 = v409
	v425 = v410
	goto L90
L90:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v415))) = uint8(v425)
	v435 = v415 + int32(1)
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+1)))
	if v436 == int32(0) {
		v443 = v435
		v445 = v381
		goto L15
	} else {
		goto L92
	}
L91:
	;
	v443 = v435
	v445 = v381
	goto L15
L92:
	;
	if base.Ui32(v435) < base.Ui32(v28) {
		v415 = v435
		v423 = v423 + int32(1)
		v425 = v436
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	m.G0 = v7 + int32(1040)
	return
L95:
	;
	m.G0 = v468 + int32(80)
	goto L94
L96:
	;
	v476 = *(*int32)(unsafe.Add(mBase, _consts[780]))
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
	if v477&int32(255) != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	if v477&int32(255) != 0 {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v481 = *(*int32)(unsafe.Add(mBase, _consts[355]))
	if v481 != 0 {
		goto L95
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	if l0&int32(1024) == int32(0) {
		goto L105
	} else {
		goto L106
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v468))) = int32(420)
	v488 = F_open(m, v476, int32(1089), v468)
	mBase = m.M
	if v488 == int32(-1) {
		goto L95
	} else {
		goto L103
	}
L102:
	;
	v491 = int32(1)
	goto L100
L103:
	;
	v491 = v488
	goto L100
L104:
	;
	if v477&int32(255) == int32(0) {
		goto L95
	} else {
		goto L112
	}
L105:
	;
	v499 = v468 + int32(16)
	v501 = F_getpid(m)
	mBase = m.M
	v503 = F_ll2string(m, v499, int32(64), base.I64_extend_i32_s(v501))
	mBase = m.M
	v508 = F_strlen(m, v499)
	mBase = m.M
	v509 = F_write(m, v491, v499, v508)
	mBase = m.M
	if v509 == int32(-1) {
		goto L104
	} else {
		goto L107
	}
L106:
	;
	v496 = F_strlen(m, v7)
	mBase = m.M
	v497 = F_write(m, v491, v7, v496)
	mBase = m.M
	goto L104
L107:
	;
	v514 = F_write(m, v491, int32(_a1566), int32(17))
	mBase = m.M
	if v514 == int32(-1) {
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v518 = v468 + int32(16)
	v521 = F___time(m, int32(0))
	mBase = m.M
	v522 = F_ll2string(m, v518, int32(64), v521)
	mBase = m.M
	v527 = F_strlen(m, v518)
	mBase = m.M
	v528 = F_write(m, v491, v518, v527)
	mBase = m.M
	if v528 == int32(-1) {
		goto L104
	} else {
		goto L109
	}
L109:
	;
	v533 = F_write(m, v491, int32(_a1567), int32(2))
	mBase = m.M
	if v533 == int32(-1) {
		goto L104
	} else {
		goto L110
	}
L110:
	;
	v536 = F_strlen(m, v7)
	mBase = m.M
	v537 = F_write(m, v491, v7, v536)
	mBase = m.M
	if v537 == int32(-1) {
		goto L104
	} else {
		goto L111
	}
L111:
	;
	v542 = F_write(m, v491, int32(_a247), int32(1))
	mBase = m.M
	goto L104
L112:
	;
	v547 = F_close(m, v491)
	mBase = m.M
	goto L95
}
func F_serverLogHexDump(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v14 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if l0&int32(255) < v14 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l3 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	F__serverLog(m, l0, int32(_a760), v9)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	goto L1
L5:
	;
	F_serverLogRaw(m, v75, int32(_a247))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
	} else {
		goto L16
	}
L6:
	;
	v24 = l0 | int32(1024)
	v27 = v9 + int32(16)
	v29 = l2
	v30 = l3
	goto L8
L7:
	;
	v75 = l0 | int32(1024)
	goto L5
L8:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v33)>>(uint(int32(4))%32)))+uint32(_consts[405]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v38)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v41 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+2)) = uint8(v41)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40&int32(15))+uint32(_consts[405]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)) = uint8(v47)
	v50 = v30 + int32(-1)
	if v50 == v41 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v75 = v24
	goto L5
L10:
	;
	if v50 != 0 {
		v27 = v67
		v29 = v29 + int32(1)
		v30 = v50
		goto L8
	} else {
		goto L15
	}
L11:
	;
	F_serverLogRaw(m, v24, v9+int32(16))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L3
	} else {
		goto L14
	}
L12:
	;
	v54 = v27 + int32(2)
	if v54-(v9+int32(16)) != int32(64) {
		v67 = v54
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v67 = v9 + int32(16)
	goto L10
L15:
	;
	goto L9
L16:
	;
	m.G0 = v9 + int32(96)
	return
}
func F_serverLogRaw(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v71 int64
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v126 int64
	_ = v126
	var v131 int64
	_ = v131
	var v136 int64
	_ = v136
	var v139 int64
	_ = v139
	var v142 int64
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v155 int64
	_ = v155
	var v160 int64
	_ = v160
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int64
	_ = v185
	var v186 int32
	_ = v186
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v194 int32
	_ = v194
	var v196 int64
	_ = v196
	var v198 int32
	_ = v198
	var v202 int64
	_ = v202
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int64
	_ = v219
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v619 int32
	_ = v619
	v14 = m.G0
	v16 = v14 - int32(1360)
	m.G0 = v16
	v19 = l0 & int32(255)
	v21 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v19 < v21 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a1556), int32(_a1555), int32(129))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L11
	} else {
		goto L100
	}
L2:
	;
	m.G0 = v16 + int32(1360)
	return
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[780]))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v25&int32(255) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v32 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	v31 = F_fopen(m, v24, int32(_a1557))
	mBase = m.M
	v32 = v31
	goto L4
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[359]))
	v32 = v29
	goto L4
L7:
	;
	if l0&int32(1024) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v578 = F_fflush(m, v32)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L11
	} else {
		goto L94
	}
L9:
	;
	v42 = F___syscall_getpid(m)
	mBase = m.M
	goto L13
L10:
	;
	v39 = F_fputs(m, l1, v32)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	goto L8
L13:
	;
	v43 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, _consts[781]))
	v48 = F___gettimeofday(m, v16+int32(1280), v43)
	mBase = m.M
	v50 = v16 + int32(1236)
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v16)+1280))
	v53 = *(*int64)(unsafe.Add(mBase, _consts[782]))
	v57 = m.G0
	v59 = v57 - int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+32)) = v44
	v62 = int32(3600)
	v65 = v51 - v53 + base.I64_extend_i32_s(v44*v62)
	v66 = int64(86400)
	v67 = base.I64_div_s(v65, v66)
	v71 = base.I64_rem_s(v67+int64(4), int64(7))
	*(*uint32)(unsafe.Add(mBase, uint32(v50)+24)) = uint32(v71)
	v76 = base.I32_wrap_i64(v65 - v67*v66)
	v78 = base.I32_div_s(v76, v62)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v78
	v82 = v76 - v78*v62
	v84 = int32(60)
	v85 = base.I32_div_s(base.I32_extend16_s(v82), v84)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = base.I32_extend16_s(v85)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = base.I32_extend16_s(v82 - v85*v84)
	v96 = v67
	v97 = int32(1970)
	goto L15
L14:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[783]))
	switch v224 {
	case 0:
		goto L41
	case 1:
		goto L40
	case 2:
		goto L39
	default:
		goto L38
	}
L15:
	;
	if v97&int32(3) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v96 = v96 - v219
	v97 = v97 + int32(1)
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v182
	v185 = int64(31)
	v186 = int32(0)
	if v96 < v185 {
		v210 = v186
		v211 = v183
		goto L33
	} else {
		goto L34
	}
L19:
	;
	v145 = base.I32_wrap_i64(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+28)) = v145
	v149 = int32(0)
	v150 = *(*int64)(unsafe.Add(mBase, _consts[784]))
	*(*int64)(unsafe.Add(mBase, uint32(v59+int32(40)))) = v150
	v155 = *(*int64)(unsafe.Add(mBase, _consts[785]))
	*(*int64)(unsafe.Add(mBase, uint32(v59+int32(32)))) = v155
	v160 = *(*int64)(unsafe.Add(mBase, _consts[786]))
	*(*int64)(unsafe.Add(mBase, uint32(v59+int32(24)))) = v160
	v165 = *(*int64)(unsafe.Add(mBase, _consts[787]))
	*(*int64)(unsafe.Add(mBase, uint32(v59+int32(16)))) = v165
	v168 = *(*int64)(unsafe.Add(mBase, _consts[788]))
	*(*int64)(unsafe.Add(mBase, uint32(v59)+8)) = v168
	v171 = *(*int64)(unsafe.Add(mBase, _consts[789]))
	*(*int64)(unsafe.Add(mBase, uint32(v59))) = v171
	v175 = base.I32_rem_u_s(v97, int32(100))
	if v175 != 0 {
		v182 = int32(29)
		v183 = v145
		goto L18
	} else {
		goto L29
	}
L20:
	;
	if int64(364) < v96 {
		v219 = int64(365)
		goto L17
	} else {
		goto L28
	}
L21:
	;
	v105 = base.I32_rem_u_s(v97, int32(100))
	if v105 != 0 {
		v111 = int64(366)
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v96 < v111 {
		goto L19
	} else {
		goto L27
	}
L23:
	;
	v109 = base.I32_rem_u_s(v97, int32(400))
	if v109 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v110 = int64(365)
	goto L26
L25:
	;
	v110 = int64(366)
	goto L26
L26:
	;
	v111 = v110
	goto L22
L27:
	;
	v219 = v111
	goto L17
L28:
	;
	v116 = base.I32_wrap_i64(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+28)) = v116
	v120 = int32(0)
	v121 = *(*int64)(unsafe.Add(mBase, _consts[784]))
	*(*int64)(unsafe.Add(mBase, uint32(v59+int32(40)))) = v121
	v126 = *(*int64)(unsafe.Add(mBase, _consts[785]))
	*(*int64)(unsafe.Add(mBase, uint32(v59+int32(32)))) = v126
	v131 = *(*int64)(unsafe.Add(mBase, _consts[786]))
	*(*int64)(unsafe.Add(mBase, uint32(v59+int32(24)))) = v131
	v136 = *(*int64)(unsafe.Add(mBase, _consts[787]))
	*(*int64)(unsafe.Add(mBase, uint32(v59+int32(16)))) = v136
	v139 = *(*int64)(unsafe.Add(mBase, _consts[788]))
	*(*int64)(unsafe.Add(mBase, uint32(v59)+8)) = v139
	v142 = *(*int64)(unsafe.Add(mBase, _consts[789]))
	*(*int64)(unsafe.Add(mBase, uint32(v59))) = v142
	v182 = int32(28)
	v183 = v116
	goto L18
L29:
	;
	v179 = base.I32_rem_u_s(v97, int32(400))
	if v179 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v180 = int32(28)
	goto L32
L31:
	;
	v180 = int32(29)
	goto L32
L32:
	;
	v182 = v180
	v183 = v145
	goto L18
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v50)+20)) = v97 + int32(-1900)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = v211 + int32(1)
	goto L14
L34:
	;
	v190 = v185
	v191 = v96
	v194 = v186
	goto L35
L35:
	;
	v196 = v191 - v190
	v198 = v194 + int32(1)
	v202 = int64(*(*int32)(unsafe.Add(mBase, uint32(v59+v198<<(uint(int32(2))%32)))))
	if v202 <= v196 {
		v190 = v202
		v191 = v196
		v194 = v198
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v210 = v198
	v211 = base.I32_wrap_i64(v196)
	goto L33
L37:
	;
	goto L36
L38:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _consts[232]))
	if v355 != 0 {
		v365 = int32(0)
		goto L51
	} else {
		goto L52
	}
L39:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1288))
	v337 = base.I32_div_s(v335, int32(1000))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+176)) = v51*int64(1000) + base.I64_extend_i32_s(v337)
	v347 = F_snprintf(m, v16+int32(1296), int32(64), int32(_a1558), v16+int32(176))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L11
	} else {
		goto L50
	}
L40:
	;
	v253 = F_strftime(m, v16+int32(1296), int32(64), int32(_a1559), v16+int32(1236))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L11
	} else {
		goto L44
	}
L41:
	;
	v231 = F_strftime(m, v16+int32(1296), int32(64), int32(_a1560), v16+int32(1236))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L11
	} else {
		goto L42
	}
L42:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1288))
	v235 = base.I32_div_s(v233, int32(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v235
	v245 = F_snprintf(m, v231+(v16+int32(1296)), int32(64)-v231, int32(_a1561), v16+int32(144))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	goto L38
L44:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _consts[782]))
	if base.Ui32(int32(93601)) <= base.Ui32(v256+int32(50400)) {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v261 = int32(0)
	v262 = *(*int32)(unsafe.Add(mBase, _consts[781]))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1235)) = uint8(v261)
	v265 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1232)) = uint8(v265)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1288))
	v269 = base.I32_div_s(v267, int32(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+160)) = v269
	v275 = v262*int32(3600) - v256
	if int32(-1) < v275 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v278 = int32(43)
	goto L48
L47:
	;
	v278 = int32(45)
	goto L48
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1229)) = uint8(v278)
	v280 = int32(3600)
	v281 = base.I32_div_s(v275, v280)
	v282 = int32(31)
	v283 = v281 >> (uint(v282) % 32)
	v285 = v281 ^ v283 - v283
	v286 = int32(10)
	v287 = base.I32_div_u_s(v285, v286)
	v288 = int32(48)
	v289 = v287 + v288
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1230)) = uint8(v289)
	v295 = v285 - v287*v286 | v288
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1231)) = uint8(v295)
	v299 = v275 - v281*v280
	v301 = v299 >> (uint(v282) % 32)
	v305 = (v299 ^ v301 - v301) & int32(65535)
	v307 = base.I32_div_u_s(v305, int32(600))
	v309 = v307 + v288
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1233)) = uint8(v309)
	v312 = base.I32_div_u_s(v305, int32(60))
	v316 = base.I32_rem_u_s(v312&int32(255), v286)
	v318 = v316 | v288
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1234)) = uint8(v318)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+164)) = v16 + int32(1229)
	v331 = F_snprintf(m, v16+int32(1296)+v253, int32(64)-v253, int32(_a1562), v16+int32(160))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	goto L38
L50:
	;
	goto L38
L51:
	;
	v366 = int32(0)
	v368 = *(*int32)(unsafe.Add(mBase, _consts[790]))
	switch v368 {
	case 0:
		goto L57
	case 1:
		goto L59
	case 2:
		goto L58
	default:
		goto L8
	}
L52:
	;
	v358 = *(*int32)(unsafe.Add(mBase, _consts[635]))
	if v42 != v358 {
		v365 = int32(1)
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v363 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v364 = int32(2)
	goto L56
L55:
	;
	v364 = int32(3)
	goto L56
L56:
	;
	v365 = v364
	goto L51
L57:
	;
	v544 = F___syscall_getpid(m)
	mBase = m.M
	goto L92
L58:
	;
	v458 = F_sdsempty(m)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L11
	} else {
		goto L71
	}
L59:
	;
	if l1 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v433 = F___syscall_getpid(m)
	mBase = m.M
	goto L69
L61:
	;
	v371 = v366
	goto L63
L62:
	;
	F_filterInvalidLogfmtChar(m, v16+int32(192), int32(1024), l1)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L11
	} else {
		goto L66
	}
L63:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v371))))
	switch v385 {
	case 0:
		goto L60
	default:
		goto L65
	case 10, 13, 34:
		goto L62
	}
L65:
	;
	v371 = v371 + int32(1)
	goto L63
L66:
	;
	v393 = F___syscall_getpid(m)
	mBase = m.M
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(64)))) = v16 + int32(192)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v393
	v400 = int32(2)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v19<<(uint(v400)%32))+uint32(_consts[791])))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v404
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v365<<(uint(v400)%32))+uint32(_consts[792])))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v16 + int32(1296)
	v418 = F_fiprintf(m, v32, int32(_a1563), v16+int32(48))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	goto L8
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(32)))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v433
	v438 = int32(2)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v19<<(uint(v438)%32))+uint32(_consts[791])))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v442
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v365<<(uint(v438)%32))+uint32(_consts[792])))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v16 + int32(1296)
	v456 = F_fiprintf(m, v32, int32(_a1563), v16+int32(16))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L11
	} else {
		goto L70
	}
L70:
	;
	goto L8
L71:
	;
	if l1&int32(3) == int32(0) {
		v481 = l1
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v515 = F_escapeJsonString(m, v458, l1, v514)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L11
	} else {
		goto L88
	}
L73:
	;
	v514 = v506 - l1
	goto L72
L74:
	;
	v485 = v481
	goto L82
L75:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v467 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v470 = l1
	goto L78
L77:
	;
	v514 = l1 - l1
	goto L72
L78:
	;
	v474 = v470 + int32(1)
	if v474&int32(3) == int32(0) {
		v481 = v474
		goto L74
	} else {
		goto L80
	}
L80:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	if v479 != 0 {
		v470 = v474
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v506 = v474
	goto L73
L82:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	v494 = int32(-2139062144)
	if (int32(16843008)-v491|v491)&v494 == v494 {
		v485 = v485 + int32(4)
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v500 = v485
	goto L85
L84:
	;
	goto L83
L85:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500))))
	if v504 != 0 {
		v500 = v500 + int32(1)
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v506 = v500
	goto L73
L87:
	;
	goto L86
L88:
	;
	v517 = F___syscall_getpid(m)
	mBase = m.M
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(96)))) = v515
	v521 = int32(2)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v19<<(uint(v521)%32))+uint32(_consts[791])))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v525
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v365<<(uint(v521)%32))+uint32(_consts[792])))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v16 + int32(1296)
	v540 = F_fiprintf(m, v32, int32(_a1564), v16+int32(80))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L11
	} else {
		goto L90
	}
L90:
	;
	F_sdsfree(m, v515)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L11
	} else {
		goto L91
	}
L91:
	;
	goto L8
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(128)))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = v544
	v551 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[793]))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+124)) = v551
	v555 = int32(*(*int8)(unsafe.Add(mBase, uint32(v365)+uint32(_consts[794]))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = v555
	*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v16 + int32(1296)
	v563 = F_fiprintf(m, v32, int32(_a1565), v16+int32(112))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L11
	} else {
		goto L93
	}
L93:
	;
	goto L8
L94:
	;
	if v25&int32(255) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v586 = int32(0)
	v587 = *(*int32)(unsafe.Add(mBase, _consts[795]))
	if v587 == v586 {
		goto L2
	} else {
		goto L98
	}
L96:
	;
	v584 = F_fclose(m, v32)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L11
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l1
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v19<<(uint(int32(2))%32))+uint32(_consts[796])))
	F_syslog(m, v595, int32(_a79), v16)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L11
	} else {
		goto L99
	}
L99:
	;
	goto L2
L100:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_serverLrand48(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v59 int32
	_ = v59
	var v73 int32
	_ = v73
	v1 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[629]))
	v12 = int32(58989)
	v13 = v11 * v12
	v16 = int32(65535)
	*(*int32)(unsafe.Add(mBase, _consts[629])) = (v13 + int32(11)) & v16
	v20 = int32(16)
	v26 = int32(base.Ui32(v13)>>(uint(v20)%32)) + base.B2i32(base.Ui32(int32(65524)) < base.Ui32(v13&v16))
	v30 = *(*int32)(unsafe.Add(mBase, _consts[630]))
	v32 = v30 * v12
	v35 = v26&v16 + v32&v16
	v38 = int32(57068)
	v39 = v11 * v38
	v42 = v35&v16 + v39&int32(65532)
	v44 = v42 & v16
	*(*int32)(unsafe.Add(mBase, _consts[630])) = v44
	v59 = *(*int32)(unsafe.Add(mBase, _consts[631]))
	v73 = (int32(base.Ui32(v39)>>(uint(v20)%32)) + v11*int32(5) + v30*v38 + int32(base.Ui32(v32)>>(uint(v20)%32)) + v59*v12 + base.B2i32(base.Ui32(v16) < base.Ui32(v26)) + base.B2i32(base.Ui32(v16) < base.Ui32(v35)) + base.B2i32(base.Ui32(v16) < base.Ui32(v42))) & v16
	*(*int32)(unsafe.Add(mBase, _consts[631])) = v73
	return v73<<(uint(int32(15))%32) | int32(base.Ui32(v44)>>(uint(int32(1))%32))
}
func F_serverOutOfMemoryHandler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v8 {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		F__serverPanic_1(m, int32(_a1555), int32(7298), int32(_a1649), v5)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = l0
		F__serverLog(m, int32(3), int32(_a1650), v5+int32(16))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
			F__serverPanic_1(m, int32(_a1555), int32(7298), int32(_a1649), v5)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
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
func F_serverSetCpuAffinity(m *base.Module, l0 int32) {
	return
}
func F_serverSetProcTitle(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F_serverSrand48(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[629])) = int32(13070)
	*(*int32)(unsafe.Add(mBase, _consts[630])) = l0 & int32(65535)
	*(*int32)(unsafe.Add(mBase, _consts[631])) = int32(base.Ui32(l0) >> (uint(int32(16)) % 32))
	return
}
