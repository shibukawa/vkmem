package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_strEncoding(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	if base.Ui32(int32(11)) < base.Ui32(l0) {
		v11 = int32(_a288)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[373])))
		v11 = v10
	}
	return v11
}
func F_str_byte(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = F_luaL_checklstring(m, l0, int32(1), v10+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v23 = F_luaL_optinteger(m, l0, int32(2), int32(1))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			v31 = v23 + v23>>(uint(int32(31))%32)&(v27+int32(1))
			v32 = int32(0)
			if v32 < v31 {
				v35 = v31
			} else {
				v35 = v32
			}
			v36 = F_luaL_optinteger(m, l0, int32(3), v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				v44 = v36 + v36>>(uint(int32(31))%32)&(v40+int32(1))
				v45 = int32(0)
				if v45 < v44 {
					v48 = v44
				} else {
					v48 = v45
				}
				if base.Ui32(v48) < base.Ui32(v40) {
					v50 = v48
				} else {
					v50 = v40
				}
				v51 = int32(1)
				if v51 < v31 {
					v54 = v31
				} else {
					v54 = v51
				}
				if base.Ui32(v50) < base.Ui32(v54) {
					v103 = int32(0)
					m.G0 = v10 + int32(16)
					return v103
				} else {
					v56 = v50 - v54
					v58 = v56 + int32(1)
					if v50 != int32(2147483647) {
						v67 = m.G3
						F_luaL_checkstack(m, l0, v58, v67+int32(_a2099))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							v82 = int32(0)
							for {
								v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v54+int32(-1)+v82))))
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(v88))) = base.F64_convert_i32_s(v86)
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v93 + int32(16)
								v98 = v82 + int32(1)
								if v98 != v56+int32(1) {
									v82 = v98
									continue
								} else {
									break
								}
								break
							}
							v103 = v58
							m.G0 = v10 + int32(16)
							return v103
						}
					} else {
						v61 = m.G3
						v65 = F_luaL_error(m, l0, v61+int32(_a2099), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v67 = m.G3
							F_luaL_checkstack(m, l0, v58, v67+int32(_a2099))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v82 = int32(0)
								for {
									v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v54+int32(-1)+v82))))
									v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = int32(3)
									*(*float64)(unsafe.Add(mBase, uint32(v88))) = base.F64_convert_i32_s(v86)
									v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v93 + int32(16)
									v98 = v82 + int32(1)
									if v98 != v56+int32(1) {
										v82 = v98
										continue
									} else {
										break
									}
									break
								}
								v103 = v58
								m.G0 = v10 + int32(16)
								return v103
							}
						}
					}
				}
			}
		}
	}
}
func F_str_find(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_str_find_aux(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_str_gsub(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v360 int32
	_ = v360
	var v361 int64
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v561 int32
	_ = v561
	var v573 int32
	_ = v573
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v732 int32
	_ = v732
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	v16 = m.G0
	v18 = v16 - int32(1328)
	m.G0 = v18
	v23 = F_luaL_checklstring(m, l0, int32(1), v18+int32(1320))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = F_luaL_checklstring(m, l0, int32(2), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	goto L7
L4:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1320))
	v93 = F_luaL_optinteger(m, l0, int32(4), v90+int32(1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L19
	}
L5:
	;
	v82 = m.G398
	if v39 != v82 {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v39 = v34 + int32(32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v39) < base.Ui32(v40) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v88 = int32(-1)
	goto L4
L17:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v88 = v85
	goto L4
L18:
	;
	v88 = int32(-1)
	goto L4
L19:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if base.Ui32(v88+int32(-3)) < base.Ui32(int32(4)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v110 = v18 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = v18 + int32(24)
	goto L23
L21:
	;
	v103 = m.G3
	v106 = F_luaL_argerror(m, l0, int32(3), v103+int32(_a2100))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1056)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1048)) = v23
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1320))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1052)) = v23 + v119
	v123 = v18 + int32(1048)
	v129 = v23
	v138 = int32(0)
	goto L25
L24:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1052))
	F_luaL_addlstring(m, v18+int32(12), v793, v798-v793)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L189
	}
L25:
	;
	if v138 < v93 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v792 = v777
	v793 = v778
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1060)) = int32(0)
	v147 = F_match(m, v18+int32(1048), v129, v29+base.B2i32(v95 == int32(94)))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v792 = v138
	v793 = v129
	goto L24
L29:
	;
	if v95 != int32(94) {
		v129 = v778
		v138 = v777
		goto L25
	} else {
		goto L188
	}
L30:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1052))
	if base.Ui32(v129) < base.Ui32(v749) {
		goto L183
	} else {
		goto L184
	}
L31:
	;
	if v147 == int32(0) {
		v745 = v138
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1056))
	goto L41
L33:
	;
	v732 = v138 + int32(1)
	if base.Ui32(v129) < base.Ui32(v147) {
		v777 = v732
		v778 = v147
		goto L29
	} else {
		goto L182
	}
L34:
	;
	goto L119
L35:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1060))
	if int32(0) < v422 {
		goto L105
	} else {
		goto L106
	}
L36:
	;
	goto L77
L37:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1056))
	v216 = F_lua_tolstring(m, v212, int32(3), v18+int32(1324))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L53
	}
L38:
	;
	switch v209 + int32(-3) {
	case 0, 1:
		goto L37
	case 2:
		goto L35
	case 3:
		goto L36
	default:
		goto L34
	}
L39:
	;
	v203 = m.G398
	if v160 != v203 {
		goto L51
	} else {
		goto L52
	}
L41:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	v160 = v155 + int32(32)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	if base.Ui32(v160) < base.Ui32(v161) {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v209 = int32(-1)
	goto L38
L51:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	v209 = v206
	goto L38
L52:
	;
	v209 = int32(-1)
	goto L38
L53:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1324))
	if v218 == int32(0) {
		goto L33
	} else {
		goto L54
	}
L54:
	;
	v226 = int32(0)
	goto L55
L55:
	;
	v238 = v216 + v226
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	if v239 == int32(37) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v301 = v296 + int32(1)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1324))
	if base.Ui32(v301) < base.Ui32(v302) {
		v226 = v301
		goto L55
	} else {
		goto L73
	}
L58:
	;
	v257 = v226 + int32(1)
	v258 = v216 + v257
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	if base.Ui32((v259+int32(-48))&int32(255)) < base.Ui32(int32(10)) {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if base.Ui32(v242) < base.Ui32(v123) {
		v250 = v239
		v251 = v242
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v251 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v251))) = uint8(v250)
	v296 = v226
	goto L57
L61:
	;
	v246 = F_luaL_prepbuffer(m, v18+int32(12))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v250 = v248
	v251 = v249
	goto L60
L63:
	;
	if v259 != int32(48) {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if base.Ui32(v266) < base.Ui32(v123) {
		v274 = v259
		v275 = v266
		goto L65
	} else {
		goto L66
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v275 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v275))) = uint8(v274)
	v296 = v257
	goto L57
L66:
	;
	v270 = F_luaL_prepbuffer(m, v18+int32(12))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v274 = v272
	v275 = v273
	goto L65
L68:
	;
	F_push_onecapture(m, v18+int32(1048), v259+int32(-49), v129, v147)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	F_luaL_addlstring(m, v18+int32(12), v129, v147-v129)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v296 = v257
	goto L57
L71:
	;
	F_luaL_addvalue(m, v18+int32(12))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v296 = v257
	goto L57
L73:
	;
	goto L33
L74:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1060))
	v370 = m.G3
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1056))
	if v369 != 0 {
		goto L90
	} else {
		goto L91
	}
L75:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v318)))
	*(*int64)(unsafe.Add(mBase, uint32(v360))) = v361
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v360)+8)) = v363
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = v365 + int32(16)
	goto L74
L77:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	v315 = v310 + int32(32)
	v316 = m.G398
	if base.Ui32(v315) < base.Ui32(v309) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v318 = v315
	goto L80
L79:
	;
	v318 = v316
	goto L80
L80:
	;
	goto L75
L90:
	;
	v373 = v369
	goto L92
L91:
	;
	v373 = int32(1)
	goto L92
L92:
	;
	if v129 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v374 = v373
	goto L95
L94:
	;
	v374 = v369
	goto L95
L95:
	;
	F_luaL_checkstack(m, v371, v374, v370+int32(_a2101))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	if v374 < int32(1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	F_lua_call(m, v151, v374, int32(1))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L103
	}
L98:
	;
	v385 = int32(0)
	goto L99
L99:
	;
	F_push_onecapture(m, v18+int32(1048), v385, v129, v147)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L101
	}
L100:
	;
	goto L97
L101:
	;
	v402 = v385 + int32(1)
	if v402 != v374 {
		v385 = v402
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	goto L34
L104:
	;
	F_lua_gettable(m, v151, int32(3))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L114
	}
L105:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1068))
	switch v429 + int32(2) {
	case 0:
		goto L110
	case 1:
		goto L109
	default:
		goto L108
	}
L106:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1056))
	F_lua_pushlstring(m, v425, v129, v147-v129)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	goto L104
L108:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1056))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1064))
	F_lua_pushlstring(m, v456, v457, v429)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L113
	}
L109:
	;
	v448 = m.G3
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1056))
	v453 = F_luaL_error(m, v449, v448+int32(_a2102), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1056))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1064))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1048))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v432)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v439)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v439))) = base.F64_convert_i32_s(v433 - v434 + int32(1))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v432)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+8)) = v444 + int32(16)
	goto L111
L111:
	;
	goto L104
L112:
	;
	goto L108
L113:
	;
	goto L104
L114:
	;
	goto L34
L115:
	;
	F_luaL_addvalue(m, v18+int32(12))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L181
	}
L116:
	;
	goto L148
L117:
	;
	if v542 != 0 {
		goto L116
	} else {
		goto L136
	}
L118:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v499)+8))
	switch v535 {
	case 0:
		v540 = v535
		goto L133
	case 1:
		goto L135
	default:
		goto L134
	}
L119:
	;
	goto L125
L125:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	v499 = v496 + int32(-16)
	goto L118
L133:
	;
	v542 = v540
	goto L117
L134:
	;
	v540 = int32(1)
	goto L133
L135:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	v542 = base.B2i32(v536 != int32(0))
	goto L117
L136:
	;
	goto L139
L137:
	;
	F_lua_pushlstring(m, v151, v129, v147-v129)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L145
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = v561 + int32(-16)
	goto L137
L139:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	goto L138
L145:
	;
	goto L115
L146:
	;
	if v635 != 0 {
		goto L115
	} else {
		goto L161
	}
L147:
	;
	v625 = m.G398
	if v591 != v625 {
		goto L159
	} else {
		goto L160
	}
L148:
	;
	goto L152
L152:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	v591 = v588 + int32(-16)
	goto L147
L159:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v591)+8))
	v635 = base.B2i32(base.Ui32(v628+int32(-3)) < base.Ui32(int32(2)))
	goto L146
L160:
	;
	v635 = int32(0)
	goto L146
L161:
	;
	goto L164
L162:
	;
	v695 = m.G3
	if v693 != int32(-1) {
		goto L178
	} else {
		goto L179
	}
L163:
	;
	v687 = m.G398
	if v653 != v687 {
		goto L175
	} else {
		goto L176
	}
L164:
	;
	goto L168
L168:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	v653 = v650 + int32(-16)
	goto L163
L175:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v653)+8))
	v693 = v690
	goto L162
L176:
	;
	v693 = int32(-1)
	goto L162
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v705
	v707 = m.G3
	v710 = F_luaL_error(m, v151, v707+int32(_a2103), v18)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L180
	}
L178:
	;
	v700 = m.G399
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v700+v693<<(uint(int32(2))%32))))
	v705 = v704
	goto L177
L179:
	;
	v705 = v695 + int32(_a2018)
	goto L177
L180:
	;
	goto L115
L181:
	;
	goto L33
L182:
	;
	v745 = v732
	goto L30
L183:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if base.Ui32(v751) < base.Ui32(v123) {
		v758 = v751
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v792 = v745
	v793 = v129
	goto L24
L185:
	;
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	v760 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v758 + v760
	*(*uint8)(unsafe.Add(mBase, uint32(v758))) = uint8(v759)
	v777 = v745
	v778 = v129 + v760
	goto L29
L186:
	;
	v755 = F_luaL_prepbuffer(m, v18+int32(12))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v758 = v757
	goto L185
L188:
	;
	goto L26
L189:
	;
	F_luaL_pushresult(m, v18+int32(12))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v807)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v807))) = base.F64_convert_i32_s(v792)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v812 + int32(16)
	goto L191
L191:
	;
	m.G0 = v18 + int32(1328)
	return int32(2)
}
func F_str_rep(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(1040)
	m.G0 = v7
	v12 = F_luaL_checklstring(m, l0, int32(1), v7+int32(1036))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = F_luaL_checkinteger(m, l0, int32(2))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(12)
	goto L4
L4:
	;
	if v17 < int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_luaL_pushresult(m, v7)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L11
	}
L6:
	;
	v30 = v17
	goto L7
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1036))
	F_luaL_addlstring(m, v7, v12, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L5
L9:
	;
	if base.Ui32(int32(1)) < base.Ui32(v30) {
		v30 = v30 + int32(-1)
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	m.G0 = v7 + int32(1040)
	return int32(1)
}
