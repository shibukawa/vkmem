package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_configKeyCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v8 == int32(0) {
		v31 = v7
		v32 = v8
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v32 - v31&int32(255)
L2:
	;
	goto L1
L3:
	;
	if v8 != v7&int32(255) {
		v31 = v7
		v32 = v8
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v14 = v3
	v15 = v4
	goto L5
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v19 == int32(0) {
		v31 = v18
		v32 = v19
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v31 = v18
	v32 = v19
	goto L2
L7:
	;
	v22 = int32(1)
	if v19 == v18&int32(255) {
		v14 = v14 + v22
		v15 = v15 + v22
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
func F_configRewriteCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[178]))
	if v9 != 0 {
		v14 = F_rewriteConfig(m, v9, int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			if v14 != int32(-1) {
				v38 = *(*int32)(unsafe.Add(mBase, _consts[15]))
				if int32(2) < v38 {
					v47 = *(*int32)(unsafe.Add(mBase, _consts[77]))
					F_addReply(m, l0, v47)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						m.G0 = v6 + int32(32)
						return
					}
				} else {
					F__serverLog(m, int32(2), int32(_a466), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, _consts[77]))
						F_addReply(m, l0, v47)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							m.G0 = v6 + int32(32)
							return
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				v21 = *(*int32)(unsafe.Add(mBase, _consts[15]))
				if int32(3) < v21 {
					v32 = F___strerror_l(m, v19, v19)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v32
					F_addReplyErrorFormat(m, l0, int32(_a467), v6)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						m.G0 = v6 + int32(32)
						return
					}
				} else {
					v24 = F___strerror_l(m, v19, v19)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v24
					F__serverLog(m, int32(3), int32(_a468), v6+int32(16))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						v32 = F___strerror_l(m, v19, v19)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v32
						F_addReplyErrorFormat(m, l0, int32(_a467), v6)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							m.G0 = v6 + int32(32)
							return
						}
					}
				}
			}
		}
	} else {
		F_addReplyError(m, l0, int32(_a469))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			m.G0 = v6 + int32(32)
			return
		}
	}
}
func F_configSetCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v496 int32
	_ = v496
	var v508 int32
	_ = v508
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v755 int32
	_ = v755
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(112)
	m.G0 = v23
	*(*int32)(unsafe.Add(mBase, uint32(v23)+104)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = v2
	v30 = int32(1)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v31&v30 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v23 + int32(112)
	return
L2:
	;
	v40 = F_listCreate(m)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L6
	}
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReplyErrorObject(m, l0, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
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
	v43 = v31 + int32(-2)
	v45 = v43 << (uint(int32(1)) % 32)
	v46 = F_valkey_calloc(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v48 = F_valkey_calloc(m, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v50 = F_valkey_malloc(m, v45)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v52 = F_valkey_calloc(m, v45)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v54 = F_valkey_calloc(m, v45)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v56 = F_valkey_malloc(m, v45)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v58 = int32(1)
	v59 = v43 >> (uint(v58) % 32)
	if v59 < v58 {
		v609 = v2
		v610 = v30
		goto L17
	} else {
		goto L18
	}
L13:
	;
	F_valkey_free(m, v46)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L4
	} else {
		goto L125
	}
L14:
	;
	if v681 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L15:
	;
	v676 = *(*int32)(unsafe.Add(mBase, _consts[170]))
	F_addReplyErrorObject(m, l0, v676)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L4
	} else {
		goto L117
	}
L16:
	;
	if v265 == int32(0) {
		v681 = v253
		goto L14
	} else {
		goto L116
	}
L17:
	;
	v631 = F_moduleConfigApplyConfig(m, v40, v23+int32(104), v23+int32(100))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L4
	} else {
		goto L109
	}
L18:
	;
	v62 = int32(0)
	v68 = v62
	v70 = v62
	v80 = v62
	v81 = v62
	goto L19
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v88 = v70 << (uint(int32(1)) % 32)
	v89 = int32(2)
	v92 = (v88 + v89) << (uint(v89) % 32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v86+v92)))
	v95 = F_objectGetVal(m, v94)
	mBase = m.M
	v97 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	v98 = F_dictFind(m, v97, v95)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L24
	}
L20:
	;
	if v266 != 0 {
		goto L16
	} else {
		goto L58
	}
L21:
	;
	v272 = v70 + int32(1)
	if v272 != v59 {
		v68 = v253
		v70 = v272
		v80 = v265
		v81 = v266
		goto L19
	} else {
		goto L57
	}
L22:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+8)))
	if v112&int32(2) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	if v81 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	if v98 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	goto L26
L26:
	;
	if v102 != 0 {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v107+v92)))
	v110 = F_objectGetVal(m, v109)
	mBase = m.M
	v253 = v110
	v265 = v80
	v266 = int32(1)
	goto L21
L29:
	;
	v253 = v68
	v265 = v80
	v266 = int32(1)
	goto L21
L30:
	;
	if v81 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	F_redactClientCommandArgument(m, l0, v88+int32(3))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	if v124&int32(1) != 0 {
		v150 = v124
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v253 = v68
	v265 = v80
	v266 = int32(1)
	goto L21
L35:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	if v166 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L36:
	;
	v151 = int32(1)
	if v150&v151 != 0 {
		goto L44
	} else {
		goto L45
	}
L37:
	;
	if base.I64_extend_i32_u(v124)&int64(32) == int64(0) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	switch v133 + int32(-1) {
	case 0:
		goto L35
	case 1:
		goto L39
	default:
		v150 = v124
		goto L36
	}
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v136 == int32(0) {
		v150 = v124
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
	if v140 == int32(0) {
		v150 = v124
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v143 = m.T0[v140].(func(*base.Module, int32) int32)(m, v136)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	if v143 == int32(1) {
		goto L35
	} else {
		goto L43
	}
L43:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	v150 = v147
	goto L36
L44:
	;
	v156 = int32(_a398)
	goto L46
L45:
	;
	v156 = int32(_a399)
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+104)) = v156
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v158+v92)))
	v161 = F_objectGetVal(m, v160)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = v161
	v253 = v68
	v265 = v80
	v266 = v151
	goto L21
L47:
	;
	v176 = int32(0)
	if v70 == v176 {
		v231 = v176
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+8)))
	if v169&int32(64) == int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v174 = int32(1)
	v253 = v68
	v265 = v174
	v266 = v174
	goto L21
L50:
	;
	v236 = int32(2)
	v237 = v70 << (uint(v236) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v46+v237))) = v102
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	*(*int32)(unsafe.Add(mBase, uint32(v48+v237))) = v241
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244+v88<<(uint(v236)%32))+12))
	v249 = F_objectGetVal(m, v248)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v50+v237))) = v249
	v253 = v68
	v265 = v80
	v266 = v231
	goto L21
L51:
	;
	v199 = v176
	goto L53
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+104)) = int32(_a400)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v210+v92)))
	v213 = F_objectGetVal(m, v212)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = v213
	v231 = int32(1)
	goto L50
L53:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v46+v199<<(uint(int32(2))%32))))
	if v203 == v102 {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v206 = v199 + int32(1)
	if v206 == v70 {
		v231 = v176
		goto L50
	} else {
		goto L56
	}
L56:
	;
	v199 = v206
	goto L53
L57:
	;
	goto L20
L58:
	;
	if v59 == int32(1) {
		v328 = int32(0)
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if v43&int32(2) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	v279 = int32(0)
	v285 = v279
	v299 = v279
	goto L61
L61:
	;
	v302 = v285 << (uint(int32(2)) % 32)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v46+v302)))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+24))
	v307 = m.T0[v306].(func(*base.Module, int32) int32)(m, v305)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L4
	} else {
		goto L63
	}
L62:
	;
	v328 = v320
	goto L59
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52+v302))) = v307
	v311 = v302 | int32(4)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v46+v311)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+24))
	v316 = m.T0[v315].(func(*base.Module, int32) int32)(m, v314)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52+v311))) = v316
	v319 = int32(2)
	v320 = v285 + v319
	v322 = v299 + v319
	if v322 != v59&int32(2147483646) {
		v285 = v320
		v299 = v322
		goto L61
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	v358 = int32(0)
	v359 = base.B2i32(v265 == v358)
	v376 = v358
	goto L69
L67:
	;
	v349 = v328 << (uint(int32(2)) % 32)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v46+v349)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+24))
	v354 = m.T0[v353].(func(*base.Module, int32) int32)(m, v352)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52+v349))) = v354
	goto L66
L69:
	;
	v382 = v376 << (uint(int32(2)) % 32)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v46+v382)))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v50+v382)))
	v389 = F_performInterfaceSet(m, v384, v386, v23+int32(104))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L4
	} else {
		goto L74
	}
L70:
	;
	v543 = int32(0)
	goto L96
L71:
	;
	v536 = v376 + int32(1)
	if v536 != v59 {
		v376 = v536
		goto L69
	} else {
		goto L95
	}
L72:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+9)))
	if v444&int32(1) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+108)) = int32(_a401)
	v411 = int32(0)
	goto L75
L74:
	;
	switch v389 {
	case 0:
		goto L73
	case 1:
		goto L72
	default:
		goto L71
	}
L75:
	;
	v415 = v411 << (uint(int32(2)) % 32)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v46+v415)))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v52+v415)))
	v422 = F_performInterfaceSet(m, v417, v419, v23+int32(108))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L4
	} else {
		goto L78
	}
L76:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = v442
	if v265 != 0 {
		goto L15
	} else {
		goto L83
	}
L77:
	;
	if v411 != v376 {
		v411 = v411 + int32(1)
		goto L75
	} else {
		goto L82
	}
L78:
	;
	if v422 != 0 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v425 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v425 {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v419
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v23)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v431
	F__serverLog(m, int32(3), int32(_a402), v23+int32(64))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	goto L77
L82:
	;
	goto L76
L83:
	;
	v681 = v253
	goto L14
L84:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v384)+20))
	if v452 == int32(0) {
		goto L71
	} else {
		goto L87
	}
L85:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v384)+84))
	F_addModuleConfigApply(m, v40, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	goto L71
L87:
	;
	v455 = int32(0)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v457 == v455 {
		v496 = v54
		v508 = v455
		goto L88
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v496))) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v56+v508<<(uint(int32(2))%32)))) = v376
	goto L71
L89:
	;
	v464 = v455
	v477 = v457
	goto L90
L90:
	;
	if v477 == v452 {
		goto L71
	} else {
		goto L92
	}
L91:
	;
	v496 = v485
	v508 = v482
	goto L88
L92:
	;
	v482 = v464 + int32(1)
	v485 = v54 + v482<<(uint(int32(2))%32)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	if v486 == int32(0) {
		v496 = v485
		v508 = v482
		goto L88
	} else {
		goto L93
	}
L93:
	;
	if base.Ui32(v464) < base.Ui32(v376) {
		v464 = v482
		v477 = v486
		goto L90
	} else {
		goto L94
	}
L94:
	;
	goto L91
L95:
	;
	goto L70
L96:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v54+v543<<(uint(int32(2))%32))))
	if v562 == int32(0) {
		v609 = v253
		v610 = v359
		goto L17
	} else {
		goto L98
	}
L97:
	;
	v609 = v253
	v610 = v359
	goto L17
L98:
	;
	v567 = m.T0[v562].(func(*base.Module, int32) int32)(m, v23+int32(104))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L4
	} else {
		goto L100
	}
L99:
	;
	v605 = v543 + int32(1)
	if v605 != v59 {
		v543 = v605
		goto L96
	} else {
		goto L107
	}
L100:
	;
	if v567 != 0 {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v570 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v570 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_restoreBackupConfig(m, v46, v52, v59, v54, int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L4
	} else {
		goto L105
	}
L103:
	;
	v573 = int32(2)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v56+v543<<(uint(v573)%32))))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v46+v576<<(uint(v573)%32))))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v580)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v581
	F__serverLog(m, int32(3), int32(_a403), v23+int32(48))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v592 = int32(2)
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v56+v543<<(uint(v592)%32))))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v46+v595<<(uint(v592)%32))))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v599)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = v600
	if v265 == int32(0) {
		v681 = v253
		goto L14
	} else {
		goto L106
	}
L106:
	;
	goto L15
L107:
	;
	goto L97
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v23)+80)) = int64(0)
	F_moduleFireServerEvent(m, int64(16), int32(0), v23+int32(80))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L4
	} else {
		goto L114
	}
L109:
	;
	if v631 != 0 {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	F_serverLogRaw(m, int32(3), int32(_a404))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	F_restoreBackupConfig(m, v46, v52, v59, v54, v40)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	if v610 != 0 {
		v681 = v609
		goto L14
	} else {
		goto L113
	}
L113:
	;
	goto L15
L114:
	;
	v650 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v650)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	goto L13
L116:
	;
	goto L15
L117:
	;
	goto L13
L118:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v23)+100))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v23)+104))
	if v708 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v681
	F_addReplyErrorFormat(m, l0, int32(_a405), v23+int32(32))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	goto L13
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v707
	F_addReplyErrorFormat(m, l0, int32(_a406), v23)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L4
	} else {
		goto L124
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v707
	F_addReplyErrorFormat(m, l0, int32(_a407), v23+int32(16))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	goto L13
L124:
	;
	goto L13
L125:
	;
	F_valkey_free(m, v48)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	F_valkey_free(m, v50)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	if v59 < int32(1) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_valkey_free(m, v52)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L4
	} else {
		goto L134
	}
L129:
	;
	v755 = int32(0)
	goto L130
L130:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v52+v755<<(uint(int32(2))%32))))
	F_sdsfree(m, v774)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L4
	} else {
		goto L132
	}
L131:
	;
	goto L128
L132:
	;
	v778 = v755 + int32(1)
	if v778 != v59 {
		v755 = v778
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	F_valkey_free(m, v54)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	F_valkey_free(m, v56)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	F_listRelease(m, v40)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	goto L1
}
func F_getConfigClientOutputBufferLimitOption(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v36 int64
	_ = v36
	var v40 int64
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v74 int64
	_ = v74
	var v78 int64
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v108 int64
	_ = v108
	var v112 int64
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	v5 = m.G0
	v7 = v5 - int32(96)
	m.G0 = v7
	v9 = F_sdsempty(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _consts[193]))
		v28 = int32(_a69)
		v29 = *(*int64)(unsafe.Add(mBase, _consts[194]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(80)))) = v29
		v36 = *(*int64)(unsafe.Add(mBase, _consts[195]))
		*(*uint32)(unsafe.Add(mBase, uint32(v7+int32(88)))) = uint32(v36)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v22
		v40 = *(*int64)(unsafe.Add(mBase, _consts[196]))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+72)) = v40
		v45 = F_sdscatprintf(m, v9, int32(_a456), v7+int32(64))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			v49 = F_sdscatlen(m, v45, int32(_a10), int32(1))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, _consts[197]))
				v66 = int32(_a69)
				v67 = *(*int64)(unsafe.Add(mBase, _consts[198]))
				*(*int64)(unsafe.Add(mBase, uint32(v7+int32(48)))) = v67
				v74 = *(*int64)(unsafe.Add(mBase, _consts[199]))
				*(*uint32)(unsafe.Add(mBase, uint32(v7+int32(56)))) = uint32(v74)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v60
				v78 = *(*int64)(unsafe.Add(mBase, _consts[200]))
				*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = v78
				v83 = F_sdscatprintf(m, v49, int32(_a456), v7+int32(32))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					v87 = F_sdscatlen(m, v83, int32(_a10), int32(1))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int32(0)
					} else {
						v98 = *(*int32)(unsafe.Add(mBase, _consts[201]))
						v102 = int32(_a69)
						v103 = *(*int64)(unsafe.Add(mBase, _consts[202]))
						*(*int64)(unsafe.Add(mBase, uint32(v7+int32(16)))) = v103
						v108 = *(*int64)(unsafe.Add(mBase, _consts[203]))
						*(*uint32)(unsafe.Add(mBase, uint32(v7+int32(24)))) = uint32(v108)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v98
						v112 = *(*int64)(unsafe.Add(mBase, _consts[204]))
						*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v112
						v115 = F_sdscatprintf(m, v87, int32(_a456), v7)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(96)
							return v115
						}
					}
				}
			}
		}
	}
}
func F_getConfigSaveOption(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_sdsempty(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	if v16 < int32(1) {
		v56 = v11
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v9 + int32(16)
	return v56
L4:
	;
	v22 = v11
	v23 = int32(0)
	goto L5
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v30 = v27 + v23<<(uint(int32(4))%32)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v31
	v36 = F_sdscatprintf(m, v22, int32(_a455), v9)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v56 = v49
	goto L3
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	if v23 == v39+int32(-1) {
		v49 = v36
		v50 = v39
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v52 = v23 + int32(1)
	if v52 < v50 {
		v22 = v49
		v23 = v52
		goto L5
	} else {
		goto L11
	}
L9:
	;
	v45 = F_sdscatlen(m, v36, int32(_a10), int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	v49 = v45
	v50 = v48
	goto L8
L11:
	;
	goto L6
}
func F_rewriteConfig(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v193 int64
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
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
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
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
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	v6 = F_rewriteConfigReadOldFile(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	return int32(0)
L3:
	;
	if v6 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	return int32(-1)
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	v18 = F_dictGetIterator(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(1)
	goto L5
L7:
	;
	F_dictReleaseIterator(m, v18)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L2
	} else {
		goto L71
	}
L8:
	;
	v27 = v18 + int32(20)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v123 == int32(0) {
		goto L7
	} else {
		goto L35
	}
L10:
	;
	v34 = v27
	v35 = v31
	goto L13
L11:
	;
	v31 = int32(1)
	goto L10
L12:
	;
	v31 = int32(0)
	goto L10
L13:
	;
	switch v35 {
	case 0:
		goto L18
	default:
		goto L17
	}
L15:
	;
	v35 = int32(0)
	goto L13
L16:
	;
	goto L9
L17:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v115
	if v115 == int32(0) {
		goto L15
	} else {
		goto L34
	}
L18:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v39 != int32(-1) {
		v78 = v39
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v79 = int32(1)
	v80 = v78 + v79
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v80
	v82 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v86+int32(26)))))
	if v90 == int32(255) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v43 != 0 {
		v78 = int32(-1)
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v45 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	if v72 != int32(-1) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v44)+16)))
	v53 = int64(*(*int8)(unsafe.Add(mBase, uint32(v44)+27)))
	v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v44)+8)))
	v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v44)+12)))
	v56 = int64(*(*int8)(unsafe.Add(mBase, uint32(v44)+26)))
	v57 = int64(*(*int32)(unsafe.Add(mBase, uint32(v44)+4)))
	v58 = F_wangHash64(m, v57)
	mBase = m.M
	v60 = F_wangHash64(m, v56+v58)
	mBase = m.M
	v62 = F_wangHash64(m, v55+v60)
	mBase = m.M
	v64 = F_wangHash64(m, v54+v62)
	mBase = m.M
	v66 = F_wangHash64(m, v53+v64)
	mBase = m.M
	v68 = F_wangHash64(m, v52+v66)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v71 = v70
	goto L22
L24:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+24)))
	v50 = v48 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+24)) = uint16(v50)
	v71 = v44
	goto L22
L25:
	;
	v78 = v72 + int32(-1)
	goto L19
L26:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v78 = v75
	goto L19
L27:
	;
	v105 = int32(2)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v85+v103<<(uint(v105)%32)+int32(4))))
	v34 = v110 + v104<<(uint(v105)%32)
	v35 = int32(1)
	goto L13
L28:
	;
	v94 = v82
	goto L30
L29:
	;
	v94 = v79 << (uint(v90) % 32)
	goto L30
L30:
	;
	if v80 < v94 {
		v103 = v86
		v104 = v80
		goto L27
	} else {
		goto L31
	}
L31:
	;
	if v86 != 0 {
		v123 = v82
		goto L16
	} else {
		goto L32
	}
L32:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v85)+20))
	if v96 == int32(-1) {
		v123 = v82
		goto L16
	} else {
		goto L33
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+4)) = int64(4294967296)
	v103 = int32(1)
	v104 = int32(0)
	goto L27
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v119
	v123 = v115
	goto L16
L35:
	;
	v129 = v123
	goto L36
L36:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	goto L39
L37:
	;
	goto L7
L38:
	;
	v152 = v18 + int32(20)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v153 != 0 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+8)))
	if v135&int32(128) != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v134)+28))
	if v138 == int32(0) {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	goto L42
L42:
	;
	m.T0[v138].(func(*base.Module, int32, int32, int32))(m, v134, v141, v6)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	goto L38
L44:
	;
	if v248 != 0 {
		v129 = v248
		goto L36
	} else {
		goto L70
	}
L45:
	;
	v159 = v152
	v160 = v156
	goto L48
L46:
	;
	v156 = int32(1)
	goto L45
L47:
	;
	v156 = int32(0)
	goto L45
L48:
	;
	switch v160 {
	case 0:
		goto L53
	default:
		goto L52
	}
L50:
	;
	v160 = int32(0)
	goto L48
L51:
	;
	goto L44
L52:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v240
	if v240 == int32(0) {
		goto L50
	} else {
		goto L69
	}
L53:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v164 != int32(-1) {
		v203 = v164
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v204 = int32(1)
	v205 = v203 + v204
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v205
	v207 = int32(0)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210+v211+int32(26)))))
	if v215 == int32(255) {
		goto L63
	} else {
		goto L64
	}
L55:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v168 != 0 {
		v203 = int32(-1)
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v170 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+20))
	if v197 != int32(-1) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v177 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v169)+16)))
	v178 = int64(*(*int8)(unsafe.Add(mBase, uint32(v169)+27)))
	v179 = int64(*(*int32)(unsafe.Add(mBase, uint32(v169)+8)))
	v180 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v169)+12)))
	v181 = int64(*(*int8)(unsafe.Add(mBase, uint32(v169)+26)))
	v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v169)+4)))
	v183 = F_wangHash64(m, v182)
	mBase = m.M
	v185 = F_wangHash64(m, v181+v183)
	mBase = m.M
	v187 = F_wangHash64(m, v180+v185)
	mBase = m.M
	v189 = F_wangHash64(m, v179+v187)
	mBase = m.M
	v191 = F_wangHash64(m, v178+v189)
	mBase = m.M
	v193 = F_wangHash64(m, v177+v191)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v193
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v196 = v195
	goto L57
L59:
	;
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+24)))
	v175 = v173 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v169)+24)) = uint16(v175)
	v196 = v169
	goto L57
L60:
	;
	v203 = v197 + int32(-1)
	goto L54
L61:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v203 = v200
	goto L54
L62:
	;
	v230 = int32(2)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v210+v228<<(uint(v230)%32)+int32(4))))
	v159 = v235 + v229<<(uint(v230)%32)
	v160 = int32(1)
	goto L48
L63:
	;
	v219 = v207
	goto L65
L64:
	;
	v219 = v204 << (uint(v215) % 32)
	goto L65
L65:
	;
	if v205 < v219 {
		v228 = v211
		v229 = v205
		goto L62
	} else {
		goto L66
	}
L66:
	;
	if v211 != 0 {
		v248 = v207
		goto L51
	} else {
		goto L67
	}
L67:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v210)+20))
	if v221 == int32(-1) {
		v248 = v207
		goto L51
	} else {
		goto L68
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+4)) = int64(4294967296)
	v228 = int32(1)
	v229 = int32(0)
	goto L62
L69:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v244
	v248 = v240
	goto L51
L70:
	;
	goto L37
L71:
	;
	F_rewriteConfigUserOption(m, v6)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	F_rewriteConfigLoadmoduleOption(m, v6)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	if v264 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_rewriteConfigRemoveOrphaned(m, v6)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L2
	} else {
		goto L77
	}
L75:
	;
	F_rewriteConfigSentinelOption(m, v6)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v271 = F_rewriteConfigGetContentFromState(m, v6)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _consts[178]))
	v275 = F_rewriteConfigOverwriteFile(m, v274, v271)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	F_sdsfree(m, v271)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	F_sdsfreesplitres(m, v279, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	F_dictRelease(m, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	F_dictRelease(m, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	F_valkey_free(m, v6)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	return v275
}
func F_rewriteConfigDirOption(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	v7 = m.G0
	v8 = int32(1024)
	v9 = v7 - v8
	m.G0 = v9
	v12 = F_getcwd(m, v9, v8)
	mBase = m.M
	v13 = F_sdsnew(m, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v9 + int32(1024)
	return
L4:
	;
	v25 = F_sdscatlen(m, v13, int32(_a10), int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L9
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v17 = F_dictAdd(m, v15, v13, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v17 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	F_sdsfree(m, v13)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L3
L9:
	;
	if v9&int32(3) == int32(0) {
		v48 = v9
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v82 = F_sdscatrepr(m, v25, v9, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L11:
	;
	v81 = v73 - v9
	goto L10
L12:
	;
	v52 = v48
	goto L20
L13:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v34 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v37 = v9
	goto L16
L15:
	;
	v81 = v9 - v9
	goto L10
L16:
	;
	v41 = v37 + int32(1)
	if v41&int32(3) == int32(0) {
		v48 = v41
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v46 != 0 {
		v37 = v41
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v73 = v41
	goto L11
L20:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v61 = int32(-2139062144)
	if (int32(16843008)-v58|v58)&v61 == v61 {
		v52 = v52 + int32(4)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v67 = v52
	goto L23
L22:
	;
	goto L21
L23:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v71 != 0 {
		v67 = v67 + int32(1)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v73 = v67
	goto L11
L25:
	;
	goto L24
L26:
	;
	v85 = F_rewriteConfigRewriteLine(m, l2, l1, v82, int32(1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L3
}
func F_rewriteConfigFormatMemory(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(64)
	m.G0 = v8
	v11 = base.B2i32(l2 == int64(0))
	if l2 == int64(0) {
		if l2 == int64(0) {
			if l2 == int64(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = l2
				v52 = F_snprintf(m, l0, l1, int32(_a408), v8+int32(48))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = v52
					m.G0 = v8 + int32(64)
					return v54
				}
			} else {
				if l2&int64(1023) != int64(0) {
					*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = l2
					v52 = F_snprintf(m, l0, l1, int32(_a408), v8+int32(48))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = v52
						m.G0 = v8 + int32(64)
						return v54
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(base.Ui64(l2) >> (uint(int64(10)) % 64))
					v46 = F_snprintf(m, l0, l1, int32(_a409), v8+int32(32))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v54 = v46
						m.G0 = v8 + int32(64)
						return v54
					}
				}
			}
		} else {
			if l2&int64(1048575) != int64(0) {
				if l2 == int64(0) {
					*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = l2
					v52 = F_snprintf(m, l0, l1, int32(_a408), v8+int32(48))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = v52
						m.G0 = v8 + int32(64)
						return v54
					}
				} else {
					if l2&int64(1023) != int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = l2
						v52 = F_snprintf(m, l0, l1, int32(_a408), v8+int32(48))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v54 = v52
							m.G0 = v8 + int32(64)
							return v54
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(base.Ui64(l2) >> (uint(int64(10)) % 64))
						v46 = F_snprintf(m, l0, l1, int32(_a409), v8+int32(32))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v54 = v46
							m.G0 = v8 + int32(64)
							return v54
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = int64(base.Ui64(l2) >> (uint(int64(20)) % 64))
				v34 = F_snprintf(m, l0, l1, int32(_a410), v8+int32(16))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v54 = v34
					m.G0 = v8 + int32(64)
					return v54
				}
			}
		}
	} else {
		if l2&int64(1073741823) != int64(0) {
			if l2 == int64(0) {
				if l2 == int64(0) {
					*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = l2
					v52 = F_snprintf(m, l0, l1, int32(_a408), v8+int32(48))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = v52
						m.G0 = v8 + int32(64)
						return v54
					}
				} else {
					if l2&int64(1023) != int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = l2
						v52 = F_snprintf(m, l0, l1, int32(_a408), v8+int32(48))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v54 = v52
							m.G0 = v8 + int32(64)
							return v54
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(base.Ui64(l2) >> (uint(int64(10)) % 64))
						v46 = F_snprintf(m, l0, l1, int32(_a409), v8+int32(32))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v54 = v46
							m.G0 = v8 + int32(64)
							return v54
						}
					}
				}
			} else {
				if l2&int64(1048575) != int64(0) {
					if l2 == int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = l2
						v52 = F_snprintf(m, l0, l1, int32(_a408), v8+int32(48))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v54 = v52
							m.G0 = v8 + int32(64)
							return v54
						}
					} else {
						if l2&int64(1023) != int64(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = l2
							v52 = F_snprintf(m, l0, l1, int32(_a408), v8+int32(48))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v54 = v52
								m.G0 = v8 + int32(64)
								return v54
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(base.Ui64(l2) >> (uint(int64(10)) % 64))
							v46 = F_snprintf(m, l0, l1, int32(_a409), v8+int32(32))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v54 = v46
								m.G0 = v8 + int32(64)
								return v54
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = int64(base.Ui64(l2) >> (uint(int64(20)) % 64))
					v34 = F_snprintf(m, l0, l1, int32(_a410), v8+int32(16))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v54 = v34
						m.G0 = v8 + int32(64)
						return v54
					}
				}
			}
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(base.Ui64(l2) >> (uint(int64(30)) % 64))
			v20 = F_snprintf(m, l0, l1, int32(_a411), v8)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v54 = v20
				m.G0 = v8 + int32(64)
				return v54
			}
		}
	}
}
func F_rewriteConfigLatencyTrackingInfoPercentilesOutputOption(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 float64
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	v12 = F_sdsnew(m, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[180]))
	if v15 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v92 = F_rewriteConfigRewriteLine(m, l2, l1, v88, int32(1))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v82 = F_sdscat(m, v12, int32(_a419))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L20
	}
L5:
	;
	v18 = int32(0)
	if v15 <= v18 {
		v88 = v12
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v25 = v12
	v27 = v18
	goto L7
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v29+v27<<(uint(int32(3))%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v10))) = v33
	v36 = v10 + int32(16)
	v43 = F_snprintf(m, v36, int32(128), int32(_a420), v10)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v48 = F_strchr(m, v36, int32(46))
	mBase = m.M
	if v48 == int32(0) {
		v68 = v43
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v74 = F_sdscatlen(m, v25, v36, v68)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L18
	}
L11:
	;
	v72 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36+v68))) = uint8(v72)
	goto L10
L12:
	;
	v53 = v43
	v54 = v36 + v43
	goto L13
L13:
	;
	v57 = v54 + int32(-1)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v58 == int32(48) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v53 = v53 + int32(-1)
	v54 = v57
	goto L13
L16:
	;
	if v58 != int32(46) {
		v68 = v53
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v68 = v53 + int32(-1)
	goto L11
L18:
	;
	v77 = v27 + int32(1)
	v79 = *(*int32)(unsafe.Add(mBase, _consts[180]))
	if v77 < v79 {
		v25 = v74
		v27 = v77
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v88 = v74
	goto L3
L20:
	;
	v88 = v82
	goto L3
L21:
	;
	m.G0 = v10 + int32(144)
	return
}
func F_rewriteConfigLoadmoduleOption(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
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
	var v40 int64
	_ = v40
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
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
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
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
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
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v173 int64
	_ = v173
	var v175 int64
	_ = v175
	var v177 int64
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	v5 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	v6 = F_dictGetIterator(m, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v6)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L2
	} else {
		goto L65
	}
L2:
	;
	return
L3:
	;
	v15 = v6 + int32(20)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v111 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v22 = v15
	v23 = v19
	goto L8
L6:
	;
	v19 = int32(1)
	goto L5
L7:
	;
	v19 = int32(0)
	goto L5
L8:
	;
	switch v23 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v23 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v103
	if v103 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v27 != int32(-1) {
		v66 = v27
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v67 = int32(1)
	v68 = v66 + v67
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v68
	v70 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+v74+int32(26)))))
	if v78 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if v31 != 0 {
		v66 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v33 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	if v60 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v40 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v32)+16)))
	v41 = int64(*(*int8)(unsafe.Add(mBase, uint32(v32)+27)))
	v42 = int64(*(*int32)(unsafe.Add(mBase, uint32(v32)+8)))
	v43 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v32)+12)))
	v44 = int64(*(*int8)(unsafe.Add(mBase, uint32(v32)+26)))
	v45 = int64(*(*int32)(unsafe.Add(mBase, uint32(v32)+4)))
	v46 = F_wangHash64(m, v45)
	mBase = m.M
	v48 = F_wangHash64(m, v44+v46)
	mBase = m.M
	v50 = F_wangHash64(m, v43+v48)
	mBase = m.M
	v52 = F_wangHash64(m, v42+v50)
	mBase = m.M
	v54 = F_wangHash64(m, v41+v52)
	mBase = m.M
	v56 = F_wangHash64(m, v40+v54)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v59 = v58
	goto L17
L19:
	;
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+24)))
	v38 = v36 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+24)) = uint16(v38)
	v59 = v32
	goto L17
L20:
	;
	v66 = v60 + int32(-1)
	goto L14
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v66 = v63
	goto L14
L22:
	;
	v93 = int32(2)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v73+v91<<(uint(v93)%32)+int32(4))))
	v22 = v98 + v92<<(uint(v93)%32)
	v23 = int32(1)
	goto L8
L23:
	;
	v82 = v70
	goto L25
L24:
	;
	v82 = v67 << (uint(v78) % 32)
	goto L25
L25:
	;
	if v68 < v82 {
		v91 = v74
		v92 = v68
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v74 != 0 {
		v111 = v70
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	if v84 == int32(-1) {
		v111 = v70
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = int64(4294967296)
	v91 = int32(1)
	v92 = int32(0)
	goto L22
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v107
	v111 = v103
	goto L11
L30:
	;
	v119 = v111
	goto L31
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
	goto L34
L32:
	;
	goto L1
L33:
	;
	v136 = v6 + int32(20)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v137 != 0 {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+80))
	if v121 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v122 = int32(_a417)
	v124 = F_moduleLoadQueueEntryToLoadmoduleOptionStr(m, v120, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v127 = F_rewriteConfigRewriteLine(m, l0, v122, v124, int32(1))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	if v232 != 0 {
		v119 = v232
		goto L31
	} else {
		goto L64
	}
L39:
	;
	v143 = v136
	v144 = v140
	goto L42
L40:
	;
	v140 = int32(1)
	goto L39
L41:
	;
	v140 = int32(0)
	goto L39
L42:
	;
	switch v144 {
	case 0:
		goto L47
	default:
		goto L46
	}
L44:
	;
	v144 = int32(0)
	goto L42
L45:
	;
	goto L38
L46:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v224
	if v224 == int32(0) {
		goto L44
	} else {
		goto L63
	}
L47:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v148 != int32(-1) {
		v187 = v148
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v188 = int32(1)
	v189 = v187 + v188
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v189
	v191 = int32(0)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v195+int32(26)))))
	if v199 == int32(255) {
		goto L57
	} else {
		goto L58
	}
L49:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if v152 != 0 {
		v187 = int32(-1)
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v154 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
	if v181 != int32(-1) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v161 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v153)+16)))
	v162 = int64(*(*int8)(unsafe.Add(mBase, uint32(v153)+27)))
	v163 = int64(*(*int32)(unsafe.Add(mBase, uint32(v153)+8)))
	v164 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v153)+12)))
	v165 = int64(*(*int8)(unsafe.Add(mBase, uint32(v153)+26)))
	v166 = int64(*(*int32)(unsafe.Add(mBase, uint32(v153)+4)))
	v167 = F_wangHash64(m, v166)
	mBase = m.M
	v169 = F_wangHash64(m, v165+v167)
	mBase = m.M
	v171 = F_wangHash64(m, v164+v169)
	mBase = m.M
	v173 = F_wangHash64(m, v163+v171)
	mBase = m.M
	v175 = F_wangHash64(m, v162+v173)
	mBase = m.M
	v177 = F_wangHash64(m, v161+v175)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v180 = v179
	goto L51
L53:
	;
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+24)))
	v159 = v157 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v153)+24)) = uint16(v159)
	v180 = v153
	goto L51
L54:
	;
	v187 = v181 + int32(-1)
	goto L48
L55:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v187 = v184
	goto L48
L56:
	;
	v214 = int32(2)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v194+v212<<(uint(v214)%32)+int32(4))))
	v143 = v219 + v213<<(uint(v214)%32)
	v144 = int32(1)
	goto L42
L57:
	;
	v203 = v191
	goto L59
L58:
	;
	v203 = v188 << (uint(v199) % 32)
	goto L59
L59:
	;
	if v189 < v203 {
		v212 = v195
		v213 = v189
		goto L56
	} else {
		goto L60
	}
L60:
	;
	if v195 != 0 {
		v232 = v191
		goto L45
	} else {
		goto L61
	}
L61:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v194)+20))
	if v205 == int32(-1) {
		v232 = v191
		goto L45
	} else {
		goto L62
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = int64(4294967296)
	v212 = int32(1)
	v213 = int32(0)
	goto L56
L63:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v224)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v228
	v232 = v224
	goto L45
L64:
	;
	goto L32
L65:
	;
	v242 = F_sdsnew(m, int32(_a417))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v246 = F_dictAdd(m, v244, v242, int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L68
	}
L67:
	;
	return
L68:
	;
	if v246 == int32(0) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	F_sdsfree(m, v242)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	goto L67
}
func F_rewriteConfigRdmaBindOption(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	v6 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	if v6 == int32(0) {
		return
	} else {
		F_rewriteConfigBindOption(m, l1, l2, int32(_a462), v6)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	}
}
func F_rewriteConfigRemoveOrphaned(m *base.Module, l0 int32) {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
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
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v214 int64
	_ = v214
	var v215 int64
	_ = v215
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v219 int64
	_ = v219
	var v221 int64
	_ = v221
	var v223 int64
	_ = v223
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v229 int64
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v296 int32
	_ = v296
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = F_dictGetIterator(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v13)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L2
	} else {
		goto L75
	}
L2:
	;
	return
L3:
	;
	v22 = v13 + int32(20)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v118 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v29 = v22
	v30 = v26
	goto L8
L6:
	;
	v26 = int32(1)
	goto L5
L7:
	;
	v26 = int32(0)
	goto L5
L8:
	;
	switch v30 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v30 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v110
	if v110 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v34 != int32(-1) {
		v73 = v34
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v74 = int32(1)
	v75 = v73 + v74
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v75
	v77 = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v81+int32(26)))))
	if v85 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v38 != 0 {
		v73 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v40 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	if v67 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v47 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v39)+16)))
	v48 = int64(*(*int8)(unsafe.Add(mBase, uint32(v39)+27)))
	v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(v39)+8)))
	v50 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v39)+12)))
	v51 = int64(*(*int8)(unsafe.Add(mBase, uint32(v39)+26)))
	v52 = int64(*(*int32)(unsafe.Add(mBase, uint32(v39)+4)))
	v53 = F_wangHash64(m, v52)
	mBase = m.M
	v55 = F_wangHash64(m, v51+v53)
	mBase = m.M
	v57 = F_wangHash64(m, v50+v55)
	mBase = m.M
	v59 = F_wangHash64(m, v49+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v48+v59)
	mBase = m.M
	v63 = F_wangHash64(m, v47+v61)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v66 = v65
	goto L17
L19:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+24)))
	v45 = v43 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+24)) = uint16(v45)
	v66 = v39
	goto L17
L20:
	;
	v73 = v67 + int32(-1)
	goto L14
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v73 = v70
	goto L14
L22:
	;
	v100 = int32(2)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v80+v98<<(uint(v100)%32)+int32(4))))
	v29 = v105 + v99<<(uint(v100)%32)
	v30 = int32(1)
	goto L8
L23:
	;
	v89 = v77
	goto L25
L24:
	;
	v89 = v74 << (uint(v85) % 32)
	goto L25
L25:
	;
	if v75 < v89 {
		v98 = v81
		v99 = v75
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v81 != 0 {
		v118 = v77
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	if v91 == int32(-1) {
		v118 = v77
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+4)) = int64(4294967296)
	v98 = int32(1)
	v99 = int32(0)
	goto L22
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v114
	v118 = v110
	goto L11
L30:
	;
	v127 = v118
	goto L31
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	goto L33
L32:
	;
	goto L1
L33:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	goto L34
L34:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v134 = F_dictFind(m, v133, v132)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L2
	} else {
		goto L37
	}
L35:
	;
	v188 = v13 + int32(20)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v189 != 0 {
		goto L50
	} else {
		goto L51
	}
L36:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(0) < v166 {
		goto L35
	} else {
		goto L46
	}
L37:
	;
	if v134 == int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v131)+20))
	if v138 == int32(0) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	goto L40
L40:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
	v152 = v150 << (uint(int32(2)) % 32)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v148+v152)))
	F_sdsfree(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v157 = F_sdsempty(m)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v159+v152))) = v157
	F_listDelNode(m, v131, v149)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v131)+20))
	if v164 != 0 {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L35
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v132
	F__serverLog(m, int32(0), int32(_a418), v10)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	goto L35
L48:
	;
	if v284 != 0 {
		v127 = v284
		goto L31
	} else {
		goto L74
	}
L49:
	;
	v195 = v188
	v196 = v192
	goto L52
L50:
	;
	v192 = int32(1)
	goto L49
L51:
	;
	v192 = int32(0)
	goto L49
L52:
	;
	switch v196 {
	case 0:
		goto L57
	default:
		goto L56
	}
L54:
	;
	v196 = int32(0)
	goto L52
L55:
	;
	goto L48
L56:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v276
	if v276 == int32(0) {
		goto L54
	} else {
		goto L73
	}
L57:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v200 != int32(-1) {
		v239 = v200
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v240 = int32(1)
	v241 = v239 + v240
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v241
	v243 = int32(0)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246+v247+int32(26)))))
	if v251 == int32(255) {
		goto L67
	} else {
		goto L68
	}
L59:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v204 != 0 {
		v239 = int32(-1)
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v206 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+20))
	if v233 != int32(-1) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v213 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v205)+16)))
	v214 = int64(*(*int8)(unsafe.Add(mBase, uint32(v205)+27)))
	v215 = int64(*(*int32)(unsafe.Add(mBase, uint32(v205)+8)))
	v216 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v205)+12)))
	v217 = int64(*(*int8)(unsafe.Add(mBase, uint32(v205)+26)))
	v218 = int64(*(*int32)(unsafe.Add(mBase, uint32(v205)+4)))
	v219 = F_wangHash64(m, v218)
	mBase = m.M
	v221 = F_wangHash64(m, v217+v219)
	mBase = m.M
	v223 = F_wangHash64(m, v216+v221)
	mBase = m.M
	v225 = F_wangHash64(m, v215+v223)
	mBase = m.M
	v227 = F_wangHash64(m, v214+v225)
	mBase = m.M
	v229 = F_wangHash64(m, v213+v227)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v229
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v232 = v231
	goto L61
L63:
	;
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205)+24)))
	v211 = v209 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v205)+24)) = uint16(v211)
	v232 = v205
	goto L61
L64:
	;
	v239 = v233 + int32(-1)
	goto L58
L65:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v239 = v236
	goto L58
L66:
	;
	v266 = int32(2)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v246+v264<<(uint(v266)%32)+int32(4))))
	v195 = v271 + v265<<(uint(v266)%32)
	v196 = int32(1)
	goto L52
L67:
	;
	v255 = v243
	goto L69
L68:
	;
	v255 = v240 << (uint(v251) % 32)
	goto L69
L69:
	;
	if v241 < v255 {
		v264 = v247
		v265 = v241
		goto L66
	} else {
		goto L70
	}
L70:
	;
	if v247 != 0 {
		v284 = v243
		goto L55
	} else {
		goto L71
	}
L71:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v246)+20))
	if v257 == int32(-1) {
		v284 = v243
		goto L55
	} else {
		goto L72
	}
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+4)) = int64(4294967296)
	v264 = int32(1)
	v265 = int32(0)
	goto L66
L73:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v276)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v280
	v284 = v276
	goto L55
L74:
	;
	goto L32
L75:
	;
	m.G0 = v10 + int32(16)
	return
}
func F_rewriteConfigReplicaOfOption(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v11 != 0 {
		v14 = F_sdsnew(m, l1)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v18 = F_dictAdd(m, v16, v14, int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				if v18 == int32(0) {
					m.G0 = v8 + int32(16)
					return
				} else {
					F_sdsfree(m, v14)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[64]))
		if v13 != 0 {
			v24 = F_sdsempty(m)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
				v28 = *(*int64)(unsafe.Add(mBase, _consts[64]))
				*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = v28
				v31 = F_sdscatprintf(m, v24, int32(_a416), v8)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v34 = F_rewriteConfigRewriteLine(m, l2, l1, v31, int32(1))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		} else {
			v14 = F_sdsnew(m, l1)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				v18 = F_dictAdd(m, v16, v14, int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					if v18 == int32(0) {
						m.G0 = v8 + int32(16)
						return
					} else {
						F_sdsfree(m, v14)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_rewriteConfigSaveOption(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	if v14 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	F_sdsfree(m, v83)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L23
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v17 = F_sdsnew(m, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v21 = F_dictAdd(m, v19, v17, int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v21 != 0 {
		v83 = v17
		goto L2
	} else {
		goto L8
	}
L8:
	;
	goto L1
L9:
	;
	v69 = F_rewriteConfigRewriteLine(m, l2, l1, v65, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L19
	}
L10:
	;
	v28 = F_sdsnew(m, l1)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	v26 = F_sdsnew(m, int32(_a412))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v65 = v26
	goto L9
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	if v31 < int32(1) {
		v65 = v28
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v39 = int32(0)
	v40 = v28
	goto L15
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v47 = v44 + v39<<(uint(int32(4))%32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
	*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v49)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v48
	v53 = F_sdscatprintf(m, v40, int32(_a413), v11)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L17
	}
L16:
	;
	v65 = v53
	goto L9
L17:
	;
	v56 = v39 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	if v56 < v58 {
		v39 = v56
		v40 = v53
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v71 = F_sdsnew(m, l1)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v75 = F_dictAdd(m, v73, v71, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	if v75 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v83 = v71
	goto L2
L23:
	;
	goto L1
}
func F_rewriteConfigUserOption(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int64
	_ = v32
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	var v69 int32
	_ = v69
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
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	v5 = m.G0
	v7 = v5 - int32(304)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v11 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(304)
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(128)
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+12)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v7)+296)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v7)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v7 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+156)) = v7 + int32(168)
	goto L9
L3:
	;
	v15 = F_sdsnew(m, int32(_a414))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = F_dictAdd(m, v17, v15, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v19 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_sdsfree(m, v15)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L1
L9:
	;
	v45 = int32(0)
	v47 = F_raxSeek(m, v7, int32(_a4), v45, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v49 = F_raxNext(m, v7)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	F_raxStop(m, v7)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L25
	}
L12:
	;
	if v49 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	goto L14
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v59 = F_sdsnew(m, int32(_a415))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L16
	}
L15:
	;
	goto L11
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v62 = F_sdscatsds(m, v59, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v66 = F_sdscatlen(m, v62, int32(_a10), int32(1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v68 = F_ACLDescribeUser(m, v57)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v70 = F_objectGetVal(m, v68)
	mBase = m.M
	v71 = F_sdscatsds(m, v66, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	F_decrRefCount(m, v68)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v77 = F_rewriteConfigRewriteLine(m, l0, int32(_a414), v71, int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v79 = F_raxNext(m, v7)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	if v79 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	goto L15
L25:
	;
	v88 = F_sdsnew(m, int32(_a414))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v92 = F_dictAdd(m, v90, v88, int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	if v92 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_sdsfree(m, v88)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	goto L1
}
func F_setConfigBindOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	if l1 < int32(17) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if l1 != int32(1) {
		v41 = l1
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(_a458)
	return int32(0)
L3:
	;
	v43 = int32(0)
	if v13 <= v43 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-1)))))
	switch v19 & int32(7) {
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
		goto L5
	}
L5:
	;
	v41 = int32(0)
	goto L3
L6:
	;
	if v36 != 0 {
		v41 = int32(1)
		goto L3
	} else {
		goto L12
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
	v36 = v35
	goto L6
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
	v36 = v32
	goto L6
L9:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
	v36 = v29
	goto L6
L10:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
	v36 = v26
	goto L6
L11:
	;
	v36 = int32(base.Ui32(v19) >> (uint(int32(3)) % 32))
	goto L6
L12:
	;
	goto L5
L13:
	;
	v69 = int32(0)
	if v41 <= v69 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v48 = v43
	goto L15
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l3+v48<<(uint(int32(2))%32))))
	F_valkey_free(m, v55)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L13
L17:
	;
	return int32(0)
L18:
	;
	v61 = v48 + int32(1)
	if v61 != v13 {
		v48 = v61
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v41
	return int32(1)
L21:
	;
	v74 = v69
	goto L22
L22:
	;
	v79 = v74 << (uint(int32(2)) % 32)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0+v79)))
	v83 = F_zstrdup(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L17
	} else {
		goto L24
	}
L23:
	;
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v79))) = v83
	v87 = v74 + int32(1)
	if v87 != v41 {
		v74 = v87
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
}
func F_setConfigSocketBindOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v11 = F_setConfigBindOption(m, l1, l2, l3, int32(_a452), int32(_a457))
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
