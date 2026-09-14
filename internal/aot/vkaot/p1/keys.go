package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_getKeysFromCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+58)))
	if v6&int32(32) == int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		if v16 == int32(0) {
			v22 = F_getKeysUsingLegacyRangeSpec(m, l0, l0, l2, l3)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return v22
			}
		} else {
			v19 = m.T0[v16].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v19
			}
		}
	} else {
		v11 = F_moduleGetCommandKeysViaAPI(m, l0, l1, l2, l3)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_getKeysFromCommandWithSpecs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v54 int64
	_ = v54
	var v61 int64
	_ = v61
	var v68 int64
	_ = v68
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v111 int64
	_ = v111
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v129 int64
	_ = v129
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v162 int32
	_ = v162
	var v165 int64
	_ = v165
	var v171 int64
	_ = v171
	var v177 int64
	_ = v177
	var v181 int64
	_ = v181
	var v185 int64
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v204 int64
	_ = v204
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int64
	_ = v220
	var v224 int64
	_ = v224
	var v225 int64
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v244 int64
	_ = v244
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v15 < int32(1) {
		v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+58)))
		if v273&int32(32) == int32(0) {
			v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			if v281 != 0 {
				v284 = m.T0[v281].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l4)
				mBase = m.M
				v285 = m.ExcPending
				if v285 != 0 {
					return int32(0)
				} else {
					v293 = v284
					return v293
				}
			} else {
				return int32(0)
			}
		} else {
			v278 = F_moduleGetCommandKeysViaAPI(m, l0, l1, l2, l4)
			mBase = m.M
			v279 = m.ExcPending
			if v279 != 0 {
				return int32(0)
			} else {
				return v278
			}
		}
	} else {
		v19 = v15 & int32(3)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		if base.Ui32(int32(4)) <= base.Ui32(v15) {
			v28 = int32(0)
			v40 = int64(0)
			v41 = v28
			v43 = v28
			for {
				v45 = int32(48)
				v48 = *(*int64)(unsafe.Add(mBase, uint32(v20+v41*v45)+8))
				v54 = *(*int64)(unsafe.Add(mBase, uint32(v20+(v41|int32(1))*v45)+8))
				v61 = *(*int64)(unsafe.Add(mBase, uint32(v20+(v41|int32(2))*v45)+8))
				v68 = *(*int64)(unsafe.Add(mBase, uint32(v20+(v41|int32(3))*v45)+8))
				v72 = v40 | (v48&v54&v61&v68 ^ int64(-1))
				v73 = int32(4)
				v74 = v41 + v73
				v76 = v43 + v73
				if v76 != v15&int32(2147483644) {
					v40 = v72
					v41 = v74
					v43 = v76
					continue
				} else {
					break
				}
				break
			}
			v87 = v72
			v88 = v74
		} else {
			v87 = int64(0)
			v88 = int32(0)
		}
		if v19 == int32(0) {
			v129 = v87
		} else {
			v102 = int32(0)
			v103 = v87
			v104 = v88
			for {
				v111 = *(*int64)(unsafe.Add(mBase, uint32(v20+v104*int32(48))+8))
				v114 = v103 | (v111 ^ int64(-1))
				v115 = int32(1)
				v118 = v102 + v115
				if v118 != v19 {
					v102 = v118
					v103 = v114
					v104 = v104 + v115
					continue
				} else {
					break
				}
				break
			}
			v129 = v114
		}
		v135 = v15 & int32(3)
		if base.Ui32(int32(4)) <= base.Ui32(v15) {
			v143 = int32(0)
			v156 = v143
			v158 = v143
			v159 = int64(0)
			for {
				v162 = int32(48)
				v165 = *(*int64)(unsafe.Add(mBase, uint32(v20+(v156|int32(3))*v162)+8))
				v171 = *(*int64)(unsafe.Add(mBase, uint32(v20+(v156|int32(2))*v162)+8))
				v177 = *(*int64)(unsafe.Add(mBase, uint32(v20+(v156|int32(1))*v162)+8))
				v181 = *(*int64)(unsafe.Add(mBase, uint32(v20+v156*v162)+8))
				v185 = v165 | (v171 | (v177 | (v181 | v159)))
				v186 = int32(4)
				v187 = v156 + v186
				v189 = v158 + v186
				if v189 != v15&int32(2147483644) {
					v156 = v187
					v158 = v189
					v159 = v185
					continue
				} else {
					break
				}
				break
			}
			v201 = v187
			v204 = v185
		} else {
			v201 = int32(0)
			v204 = int64(0)
		}
		if v135 == int32(0) {
			v244 = v204
		} else {
			v215 = int32(0)
			v217 = v201
			v220 = v204
			for {
				v224 = *(*int64)(unsafe.Add(mBase, uint32(v20+v217*int32(48))+8))
				v225 = v224 | v220
				v226 = int32(1)
				v229 = v215 + v226
				if v229 != v135 {
					v215 = v229
					v217 = v217 + v226
					v220 = v225
					continue
				} else {
					break
				}
				break
			}
			v244 = v225
		}
		if v129&int64(256) == int64(0) {
			v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+58)))
			if v273&int32(32) == int32(0) {
				v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
				if v281 != 0 {
					v284 = m.T0[v281].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l4)
					mBase = m.M
					v285 = m.ExcPending
					if v285 != 0 {
						return int32(0)
					} else {
						v293 = v284
						return v293
					}
				} else {
					return int32(0)
				}
			} else {
				v278 = F_moduleGetCommandKeysViaAPI(m, l0, l1, l2, l4)
				mBase = m.M
				v279 = m.ExcPending
				if v279 != 0 {
					return int32(0)
				} else {
					return v278
				}
			}
		} else {
			if v244&int64(1024) != int64(0) {
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+58)))
				if v273&int32(32) == int32(0) {
					v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
					if v281 != 0 {
						v284 = m.T0[v281].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l4)
						mBase = m.M
						v285 = m.ExcPending
						if v285 != 0 {
							return int32(0)
						} else {
							v293 = v284
							return v293
						}
					} else {
						return int32(0)
					}
				} else {
					v278 = F_moduleGetCommandKeysViaAPI(m, l0, l1, l2, l4)
					mBase = m.M
					v279 = m.ExcPending
					if v279 != 0 {
						return int32(0)
					} else {
						return v278
					}
				}
			} else {
				v253 = F_getKeysUsingKeySpecs(m, l0, l1, l2, l3, l4)
				mBase = m.M
				v256 = m.ExcPending
				if v256 != 0 {
					return int32(0)
				} else {
					if int32(-1) < v253 {
						v293 = v253
						return v293
					} else {
						v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+58)))
						if v273&int32(32) == int32(0) {
							v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							if v281 != 0 {
								v284 = m.T0[v281].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l4)
								mBase = m.M
								v285 = m.ExcPending
								if v285 != 0 {
									return int32(0)
								} else {
									v293 = v284
									return v293
								}
							} else {
								return int32(0)
							}
						} else {
							v278 = F_moduleGetCommandKeysViaAPI(m, l0, l1, l2, l4)
							mBase = m.M
							v279 = m.ExcPending
							if v279 != 0 {
								return int32(0)
							} else {
								return v278
							}
						}
					}
				}
			}
		}
	}
}
func F_getKeysUsingKeySpecs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v285 int64
	_ = v285
	var v291 int32
	_ = v291
	var v293 int64
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v307 int64
	_ = v307
	var v312 int64
	_ = v312
	var v316 int32
	_ = v316
	var v318 int64
	_ = v318
	var v320 int32
	_ = v320
	var v328 int64
	_ = v328
	var v352 int64
	_ = v352
	var v371 int32
	_ = v371
	var v380 int64
	_ = v380
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int64
	_ = v390
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v481 int64
	_ = v481
	var v488 int32
	_ = v488
	var v490 int64
	_ = v490
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v22 + int32(16)
	return v581
L2:
	;
	v38 = l3 & int32(2)
	v40 = l4 + int32(12)
	v56 = int32(0)
	goto L8
L3:
	;
	F__serverAssert(m, int32(_a_F_getKeysUsingKeySpecs_0), int32(_a_F_getKeysUsingKeySpecs_1), int32(2333))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if int32(1) <= v25 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v581 = int32(0)
	goto L1
L6:
	;
	return int32(0)
L7:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L8:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v68 = v65 + v56*int32(48)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	if v69 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v581 = v577
	goto L1
L10:
	;
	v574 = v56 + int32(1)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v574 < v575 {
		v56 = v574
		goto L8
	} else {
		goto L132
	}
L11:
	;
	F__serverAssert(m, int32(_a_F_getKeysUsingKeySpecs_2), int32(_a_F_getKeysUsingKeySpecs_1), int32(2292))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L6
	} else {
		goto L131
	}
L12:
	;
	F__serverAssert(m, int32(_a_F_getKeysUsingKeySpecs_3), int32(_a_F_getKeysUsingKeySpecs_1), int32(2374))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L6
	} else {
		goto L130
	}
L13:
	;
	F__serverAssert(m, int32(_a_F_getKeysUsingKeySpecs_4), int32(_a_F_getKeysUsingKeySpecs_1), int32(2337))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L6
	} else {
		goto L129
	}
L14:
	;
	if l3&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	switch v69 + int32(-2) {
	case 0:
		goto L22
	case 1:
		goto L21
	default:
		goto L19
	}
L16:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
	if v72&int64(256) != int64(0) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	v581 = int32(-1)
	goto L1
L19:
	;
	if v38 != 0 {
		goto L10
	} else {
		goto L128
	}
L20:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
	switch v186 + int32(-2) {
	case 0:
		goto L54
	case 1:
		goto L53
	default:
		goto L19
	}
L21:
	;
	v80 = int32(0)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	if v80 < v81 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v179 = v79
	goto L20
L23:
	;
	v84 = v80
	goto L25
L24:
	;
	v84 = l2
	goto L25
L25:
	;
	v85 = v84 + v81
	v86 = int32(1)
	if v81 < v86 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v89 = v86
	goto L28
L27:
	;
	v89 = l2 + int32(-1)
	goto L28
L28:
	;
	if v85 == v89 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	if v89 < v85 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v94 = int32(-1)
	goto L32
L31:
	;
	v94 = int32(1)
	goto L32
L32:
	;
	v98 = v85
	goto L33
L33:
	;
	if l2 <= v98 {
		goto L10
	} else {
		goto L35
	}
L34:
	;
	v179 = v98 + int32(1)
	goto L20
L35:
	;
	if v98 < int32(1) {
		goto L10
	} else {
		goto L36
	}
L36:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1+v98<<(uint(int32(2))%32))))
	v121 = F_objectGetVal(m, v120)
	mBase = m.M
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if v125 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L34
L38:
	;
	if v157-v159 == int32(0) {
		goto L37
	} else {
		goto L50
	}
L39:
	;
	v157 = F_tolower(m, v153)
	mBase = m.M
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	v159 = F_tolower(m, v158)
	mBase = m.M
	goto L38
L40:
	;
	v127 = v121
	v128 = v122
	v129 = v125
	goto L43
L41:
	;
	v153 = int32(0)
	v154 = v122
	goto L39
L42:
	;
	v153 = v150 & int32(255)
	v154 = v149
	goto L39
L43:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v131 == int32(0) {
		v149 = v128
		v150 = v129
		goto L42
	} else {
		goto L45
	}
L44:
	;
	v149 = v143
	v150 = int32(0)
	goto L42
L45:
	;
	v135 = v129 & int32(255)
	if v135 == v131 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v142 = int32(1)
	v143 = v128 + v142
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	if v144 != 0 {
		v127 = v127 + v142
		v128 = v143
		v129 = v144
		goto L43
	} else {
		goto L49
	}
L47:
	;
	v137 = F_tolower(m, v135)
	mBase = m.M
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	v139 = F_tolower(m, v138)
	mBase = m.M
	if v137 == v139 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v149 = v128
	v150 = v141
	goto L42
L49:
	;
	goto L44
L50:
	;
	v163 = v98 + v94
	if v163 != v89 {
		v98 = v163
		goto L33
	} else {
		goto L51
	}
L51:
	;
	goto L10
L52:
	;
	if l2 <= v398 {
		goto L19
	} else {
		goto L97
	}
L53:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
	if l2 <= v203 {
		goto L19
	} else {
		goto L60
	}
L54:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v68)+36))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
	if v190 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v68)+40))
	if v194 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v397 = v179
	v398 = v190 + v179
	v399 = v189
	goto L52
L57:
	;
	if v190 != int32(-1) {
		goto L12
	} else {
		goto L59
	}
L58:
	;
	v397 = v179
	v398 = v190 + l2
	v399 = v189
	goto L52
L59:
	;
	v199 = base.I32_div_s(l2-v179, v194)
	v397 = v179
	v398 = v179 + v199 + int32(-1)
	v399 = v189
	goto L52
L60:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v68)+40))
	v207 = int32(2)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1+v179<<(uint(v207)%32)+v203<<(uint(v207)%32))))
	v214 = F_objectGetVal(m, v213)
	mBase = m.M
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214+int32(-1)))))
	switch v217 & int32(7) {
	case 0:
		goto L66
	case 1:
		goto L65
	case 2:
		goto L64
	case 3:
		goto L63
	case 4:
		goto L62
	default:
		v234 = int32(0)
		goto L61
	}
L61:
	;
	v236 = v22 + int32(8)
	v237 = int32(0)
	if base.Ui32(v234+int32(-21)) < base.Ui32(int32(-20)) {
		v371 = v237
		goto L68
	} else {
		goto L69
	}
L62:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v214+int32(-17))))
	v234 = v233
	goto L61
L63:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v214+int32(-9))))
	v234 = v230
	goto L61
L64:
	;
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214+int32(-5)))))
	v234 = v227
	goto L61
L65:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214+int32(-3)))))
	v234 = v224
	goto L61
L66:
	;
	v234 = int32(base.Ui32(v217) >> (uint(int32(3)) % 32))
	goto L61
L67:
	;
	if v371 == int32(0) {
		goto L19
	} else {
		goto L94
	}
L68:
	;
	goto L67
L69:
	;
	v249 = int32(1)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	if v234 != v249 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v371 = int32(1)
	goto L68
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v236))) = v352
	goto L70
L72:
	;
	if v250&int32(255) == int32(45) {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	v254 = v250 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v254&int32(255)) {
		v371 = v237
		goto L68
	} else {
		goto L74
	}
L74:
	;
	if v236 == int32(0) {
		goto L70
	} else {
		goto L75
	}
L75:
	;
	v352 = base.I64_extend_i32_u(v254) & int64(255)
	goto L71
L76:
	;
	if base.Ui32(int32(8)) < base.Ui32((v273+int32(-49))&int32(255)) {
		v371 = v237
		goto L68
	} else {
		goto L79
	}
L77:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+1)))
	v272 = int32(2)
	v273 = v270
	v274 = v214 + int32(1)
	goto L76
L78:
	;
	v272 = v249
	v273 = v250
	v274 = v214
	goto L76
L79:
	;
	v285 = base.I64_extend_i32_u(v273+int32(-48)) & int64(255)
	if base.Ui32(v234) <= base.Ui32(v272) {
		v328 = v285
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if v250&int32(255) != int32(45) {
		goto L88
	} else {
		goto L89
	}
L81:
	;
	v291 = v272
	v293 = v285
	v295 = v274
	goto L82
L82:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+1)))
	if base.Ui32((v297+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v371 = v237
		goto L68
	} else {
		goto L84
	}
L83:
	;
	v328 = v318
	goto L80
L84:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v293) {
		v371 = v237
		goto L68
	} else {
		goto L85
	}
L85:
	;
	v307 = v293 * int64(10)
	v312 = base.I64_extend_i32_u(v297+int32(-48)) & int64(255)
	if base.Ui64(v312^int64(-1)) < base.Ui64(v307) {
		v371 = v237
		goto L68
	} else {
		goto L86
	}
L86:
	;
	v316 = int32(1)
	v318 = v307 + v312
	v320 = v291 + v316
	if v320 != v234 {
		v291 = v320
		v293 = v318
		v295 = v295 + v316
		goto L82
	} else {
		goto L87
	}
L87:
	;
	goto L83
L88:
	;
	if v328 < int64(0) {
		v371 = v237
		goto L68
	} else {
		goto L92
	}
L89:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v328) {
		v371 = v237
		goto L68
	} else {
		goto L90
	}
L90:
	;
	if v236 == int32(0) {
		goto L70
	} else {
		goto L91
	}
L91:
	;
	v352 = int64(0) - v328
	goto L71
L92:
	;
	if v236 == int32(0) {
		goto L70
	} else {
		goto L93
	}
L93:
	;
	v352 = v328
	goto L71
L94:
	;
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	if v380 < int64(0) {
		goto L19
	} else {
		goto L95
	}
L95:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v68)+36))
	v388 = v387 + v179
	v390 = (v380+int64(-1))*base.I64_extend_i32_s(v205) + base.I64_extend_i32_s(v388)
	if base.Ui64(v390+int64(-2147483648)) <= base.Ui64(int64(-4294967297)) {
		goto L19
	} else {
		goto L96
	}
L96:
	;
	v397 = v388
	v398 = base.I32_wrap_i64(v390)
	v399 = v205
	goto L52
L97:
	;
	if l2 <= v397 {
		goto L19
	} else {
		goto L98
	}
L98:
	;
	if v398 < v397 {
		goto L19
	} else {
		goto L99
	}
L99:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v406 != 0 {
		v408 = v406
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v412 = v398 - v397 + v405 + int32(1)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v412 <= v413 {
		v437 = v408
		goto L103
	} else {
		goto L104
	}
L101:
	;
	if v405 != 0 {
		goto L11
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v40
	v408 = v40
	goto L100
L103:
	;
	v441 = v397
	goto L114
L104:
	;
	v416 = v412 << (uint(int32(3)) % 32)
	if v408 == v40 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v412
	v437 = v434
	goto L103
L106:
	;
	v421 = F_valkey_malloc(m, v416)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L6
	} else {
		goto L109
	}
L107:
	;
	v418 = F_valkey_realloc(m, v408, v416)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v418
	v434 = v418
	goto L105
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v421
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v424 == int32(0) {
		v434 = v421
		goto L105
	} else {
		goto L110
	}
L110:
	;
	v428 = v424 << (uint(int32(3)) % 32)
	if v428 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v434 = v421
	goto L105
L112:
	;
	goto L111
L113:
	;
	v431 = F__emscripten_memcpy_bulkmem(m, v421, v40, v428)
	mBase = m.M
	goto L112
L114:
	;
	if v397 <= v441 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	if v38 != 0 {
		goto L10
	} else {
		goto L126
	}
L116:
	;
	v488 = v441 + v399
	if v488 <= v398 {
		v441 = v488
		goto L114
	} else {
		goto L125
	}
L117:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v479 = v437 + v476<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v479))) = v441
	v481 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v479)+4)) = uint32(v481)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v476 + int32(1)
	goto L116
L118:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v458&int32(8) != 0 {
		goto L116
	} else {
		goto L119
	}
L119:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v461 < int32(0) {
		goto L116
	} else {
		goto L120
	}
L120:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _c_F_getKeysUsingKeySpecs[0]))
	if v467 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v468 = int32(_a_F_getKeysUsingKeySpecs_5)
	goto L123
L122:
	;
	v468 = int32(_a_F_getKeysUsingKeySpecs_6)
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v468
	F__serverPanic_1(m, int32(_a_F_getKeysUsingKeySpecs_1), int32(2419), int32(_a_F_getKeysUsingKeySpecs_7), v22)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	goto L115
L126:
	;
	v490 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
	if v490&int64(512) == int64(0) {
		goto L10
	} else {
		goto L127
	}
L127:
	;
	goto L18
L128:
	;
	goto L18
L129:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	goto L9
}
func F_setGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v103 int32
	_ = v103
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v8 != 0 {
		v13 = v8
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_setGetKeys_0), int32(_a_F_setGetKeys_1), int32(2292))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L10
	} else {
		goto L27
	}
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if int32(0) < v14 {
		v43 = v13
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v9 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v11 = l3 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v11
	v13 = v11
	goto L2
L5:
	;
	v46 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v46
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v46
	if l2 < int32(4) {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	v18 = l3 + int32(12)
	if v13 == v18 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v43 = v39
	goto L5
L8:
	;
	v27 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L10
	} else {
		goto L12
	}
L9:
	;
	v21 = F_valkey_realloc(m, v13, int32(8))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v21
	v39 = v21
	goto L7
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v27
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v30 == int32(0) {
		v39 = v27
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v34 = v30 << (uint(int32(3)) % 32)
	if v34 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v39 = v27
	goto L7
L15:
	;
	goto L14
L16:
	;
	v37 = F__emscripten_memcpy_bulkmem(m, v27, v18, v34)
	mBase = m.M
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = int32(36)
	return int32(1)
L18:
	;
	v56 = int32(3)
	goto L19
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1+v56<<(uint(int32(2))%32))))
	v64 = F_objectGetVal(m, v63)
	mBase = m.M
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v65|int32(32) != int32(103) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L17
L21:
	;
	v86 = v56 + int32(1)
	if v86 != l2 {
		v56 = v86
		goto L19
	} else {
		goto L26
	}
L22:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v70|int32(32) != int32(101) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+2)))
	if v75|int32(32) != int32(116) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+3)))
	if v80 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = int32(50)
	return int32(1)
L26:
	;
	goto L20
L27:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
