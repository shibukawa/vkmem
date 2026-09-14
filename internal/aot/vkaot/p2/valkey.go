package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_genValkeyInfoStringLatencyStats(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
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
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v248 int32
	_ = v248
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(64)
	m.G0 = v9
	v12 = v9 + int32(16)
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v3
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v13)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(-1)
	if l1 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v35 = F_hashtableNext(m, v9+int32(16), v9+int32(12))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L1
L3:
	;
	goto L4
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v12
	goto L2
L5:
	;
	F_hashtableCleanupIterator(m, v9+int32(16))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L6
	} else {
		goto L59
	}
L6:
	;
	return int32(0)
L7:
	;
	if v35 == int32(0) {
		v239 = l0
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v41 = l0
	goto L9
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+148))
	if v48 == int32(0) {
		v223 = v41
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v239 = v232
	goto L5
L11:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v47)+200))
	if v227 == int32(0) {
		v232 = v223
		goto L54
	} else {
		goto L55
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-1)))))
	switch v55 & int32(7) {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L16
	case 3:
		goto L15
	case 4:
		goto L14
	default:
		v72 = int32(0)
		goto L13
	}
L13:
	;
	if v72 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-17))))
	v72 = v71
	goto L13
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-9))))
	v72 = v68
	goto L13
L16:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(-5)))))
	v72 = v65
	goto L13
L17:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-3)))))
	v72 = v62
	goto L13
L18:
	;
	v72 = int32(base.Ui32(v55) >> (uint(int32(3)) % 32))
	goto L13
L19:
	;
	v144 = F_valkey_malloc(m, v72+int32(1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L35
	}
L20:
	;
	if v137 != 0 {
		goto L19
	} else {
		goto L33
	}
L21:
	;
	goto L20
L22:
	;
	v137 = int32(0)
	goto L21
L23:
	;
	v86 = int32(0)
	goto L24
L24:
	;
	goto L27
L25:
	;
	goto L22
L26:
	;
	v120 = v86 + int32(1)
	if v120 != v72 {
		v86 = v120
		goto L24
	} else {
		goto L32
	}
L27:
	;
	v93 = v52 + v86
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	v102 = int32(0)
	goto L28
L28:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+uint32(_consts[904]))))
	if v94&int32(255) == v106 {
		v137 = v93
		goto L21
	} else {
		goto L30
	}
L29:
	;
	goto L26
L30:
	;
	v109 = v102 + int32(1)
	if v109 != int32(4) {
		v102 = v109
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	goto L25
L33:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v47)+148))
	v140 = F_fillPercentileDistributionLatencies(m, v41, v52, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v223 = v140
	goto L11
L35:
	;
	if v72 == int32(0) {
		v149 = v144
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v151 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v149+v72))) = uint8(v151)
	if v72 == v151 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L36
L38:
	;
	v148 = F__emscripten_memcpy_bulkmem(m, v144, v52, v72)
	mBase = m.M
	v149 = v148
	goto L37
L39:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v47)+148))
	v219 = F_fillPercentileDistributionLatencies(m, v41, v149, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L6
	} else {
		goto L52
	}
L40:
	;
	goto L39
L41:
	;
	v168 = int32(0)
	goto L42
L42:
	;
	goto L45
L43:
	;
	goto L40
L44:
	;
	v207 = v168 + int32(1)
	if v207 != v72 {
		v168 = v207
		goto L42
	} else {
		goto L51
	}
L45:
	;
	v175 = v149 + v168
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v185 = int32(0)
	goto L46
L46:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+uint32(_consts[904]))))
	if v176&int32(255) != v189 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L44
L48:
	;
	v195 = v185 + int32(1)
	if v195 != int32(4) {
		v185 = v195
		goto L46
	} else {
		goto L50
	}
L49:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+uint32(_consts[905]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v192)
	goto L44
L50:
	;
	goto L47
L51:
	;
	goto L43
L52:
	;
	F_valkey_free(m, v149)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	v223 = v219
	goto L11
L54:
	;
	v237 = F_hashtableNext(m, v9+int32(16), v9+int32(12))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L6
	} else {
		goto L57
	}
L55:
	;
	v230 = F_genValkeyInfoStringLatencyStats(m, v223, v227)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	v232 = v230
	goto L54
L57:
	;
	if v237 != 0 {
		v41 = v232
		goto L9
	} else {
		goto L58
	}
L58:
	;
	goto L10
L59:
	;
	m.G0 = v9 + int32(64)
	return v239
}
func F_valkeyAeDelRead(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2 == int32(0) {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		F_aeDeleteFileEvent(m, v7, v8, int32(1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_valkeyAppendCmdLen(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v22 int64
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v5 = F_sdscatlen(m, v4, l1, l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v5
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
			v11 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)) = uint8(v11)
			v13 = m.G3
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1002])))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v16
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[1003])))
			*(*int64)(unsafe.Add(mBase, uint32(l0+int32(13)))) = v22
			return int32(-1)
		}
	}
}
func F_valkeyAsyncConnectWithOptions(m *base.Module, l0 int32) int32 {
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
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v102 int32
	_ = v102
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = int32(32)
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0+v12)))
	*(*int64)(unsafe.Add(mBase, uint32(v10+v12))) = v16
	v18 = int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0+v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v10+v18))) = v22
	v24 = int32(24)
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l0+v24)))
	*(*int64)(unsafe.Add(mBase, uint32(v10+v24))) = v28
	v30 = int32(16)
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0+v30)))
	*(*int64)(unsafe.Add(mBase, uint32(v10+v30))) = v34
	v36 = int32(8)
	v40 = *(*int64)(unsafe.Add(mBase, uint32(l0+v36)))
	*(*int64)(unsafe.Add(mBase, uint32(v10+v36))) = v40
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v2
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v47 | int32(9)
	v51 = F_valkeyConnectWithOptions(m, v10)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		return int32(0)
	} else {
		if v51 == int32(0) {
			v142 = v2
			m.G0 = v10 + int32(48)
			return v142
		} else {
			v57 = int32(0)
			v60 = m.G3
			v63 = F_dictCreate(m, v60+int32(_a1883))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				if v63 == int32(0) {
					v87 = v57
					v88 = v57
					F_dictRelease(m, v63)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						F_dictRelease(m, v87)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							F_dictRelease(m, v88)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_valkeyFree(m, v51)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									v142 = v57
									m.G0 = v10 + int32(48)
									return v142
								}
							}
						}
					}
				} else {
					v67 = int32(0)
					v69 = m.G3
					v72 = F_dictCreate(m, v69+int32(_a1883))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						if v72 == int32(0) {
							v87 = v67
							v88 = v67
							F_dictRelease(m, v63)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								F_dictRelease(m, v87)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									F_dictRelease(m, v88)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										F_valkeyFree(m, v51)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											v142 = v57
											m.G0 = v10 + int32(48)
											return v142
										}
									}
								}
							}
						} else {
							v76 = m.G3
							v79 = F_dictCreate(m, v76+int32(_a1883))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								if v79 != 0 {
									v83 = m.G4
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
									v85 = m.T0[v84].(func(*base.Module, int32, int32) int32)(m, v51, int32(300))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										if v85 != 0 {
											v98 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v85)+272)) = v98
											*(*int64)(unsafe.Add(mBase, uint32(v85)+212)) = v98
											v102 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v85)+292)) = v102
											*(*int32)(unsafe.Add(mBase, uint32(v85)+288)) = v79
											*(*int32)(unsafe.Add(mBase, uint32(v85)+284)) = v72
											*(*int32)(unsafe.Add(mBase, uint32(v85)+280)) = v63
											*(*int64)(unsafe.Add(mBase, uint32(v85+int32(220)))) = v98
											*(*int64)(unsafe.Add(mBase, uint32(v85+int32(228)))) = v98
											*(*int64)(unsafe.Add(mBase, uint32(v85+int32(236)))) = v98
											*(*int64)(unsafe.Add(mBase, uint32(v85+int32(244)))) = v98
											*(*int64)(unsafe.Add(mBase, uint32(v85+int32(252)))) = v98
											*(*int32)(unsafe.Add(mBase, uint32(v85+int32(260)))) = v102
											v131 = *(*int32)(unsafe.Add(mBase, uint32(v85)+140))
											*(*int32)(unsafe.Add(mBase, uint32(v85)+140)) = v131 & int32(-3)
											v135 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
											*(*int32)(unsafe.Add(mBase, uint32(v85)+296)) = v135
											*(*int32)(unsafe.Add(mBase, uint32(v85)+208)) = v85 + int32(8)
											v140 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v85)+204)) = v140
											v142 = v85
											m.G0 = v10 + int32(48)
											return v142
										} else {
											v87 = v72
											v88 = v79
											F_dictRelease(m, v63)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int32(0)
											} else {
												F_dictRelease(m, v87)
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return int32(0)
												} else {
													F_dictRelease(m, v88)
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return int32(0)
													} else {
														F_valkeyFree(m, v51)
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															v142 = v57
															m.G0 = v10 + int32(48)
															return v142
														}
													}
												}
											}
										}
									}
								} else {
									v87 = v72
									v88 = int32(0)
									F_dictRelease(m, v63)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_dictRelease(m, v87)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return int32(0)
										} else {
											F_dictRelease(m, v88)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												F_valkeyFree(m, v51)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													v142 = v57
													m.G0 = v10 + int32(48)
													return v142
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
	}
}
func F_valkeyAsyncHandleConnectFailure(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	if v4 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v27
		if v27 != 0 {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v51 | int32(4)
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
			if v56 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
				if v64&int32(2) != 0 {
					return
				} else {
					F_valkeyAsyncFreeInternal(m, l0)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
				m.T0[v56].(func(*base.Module, int32))(m, v59)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
					v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
					if v64&int32(2) != 0 {
						return
					} else {
						F_valkeyAsyncFreeInternal(m, l0)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
			if v29 == int32(0) {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
				if v56 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
					v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
					if v64&int32(2) != 0 {
						return
					} else {
						F_valkeyAsyncFreeInternal(m, l0)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
					m.T0[v56].(func(*base.Module, int32))(m, v59)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
						if v64&int32(2) != 0 {
							return
						} else {
							F_valkeyAsyncFreeInternal(m, l0)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v32
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
				if v29 != v34 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
				}
				v38 = m.G4
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
				m.T0[v39].(func(*base.Module, int32))(m, v29)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					v42 = m.G3
					m.Env.X__assert_fail(m, v42+int32(_a1884), v42+int32(_a1885), int32(423), v42+int32(_a1886))
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
		if v7&int32(16) != 0 {
			m.T0[v4].(func(*base.Module, int32, int32))(m, l0, int32(-1))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v27
				if v27 != 0 {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v51 | int32(4)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
					if v56 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
						if v64&int32(2) != 0 {
							return
						} else {
							F_valkeyAsyncFreeInternal(m, l0)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
						m.T0[v56].(func(*base.Module, int32))(m, v59)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
							if v64&int32(2) != 0 {
								return
							} else {
								F_valkeyAsyncFreeInternal(m, l0)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
					if v29 == int32(0) {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
						if v56 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
							if v64&int32(2) != 0 {
								return
							} else {
								F_valkeyAsyncFreeInternal(m, l0)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
							m.T0[v56].(func(*base.Module, int32))(m, v59)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
								v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
								if v64&int32(2) != 0 {
									return
								} else {
									F_valkeyAsyncFreeInternal(m, l0)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v32
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
						if v29 != v34 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
						}
						v38 = m.G4
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
						m.T0[v39].(func(*base.Module, int32))(m, v29)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							v42 = m.G3
							m.Env.X__assert_fail(m, v42+int32(_a1884), v42+int32(_a1885), int32(423), v42+int32(_a1886))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v7 | int32(16)
			m.T0[v4].(func(*base.Module, int32, int32))(m, l0, int32(-1))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v16 & int32(-17)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v27
				if v27 != 0 {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v51 | int32(4)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
					if v56 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
						if v64&int32(2) != 0 {
							return
						} else {
							F_valkeyAsyncFreeInternal(m, l0)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
						m.T0[v56].(func(*base.Module, int32))(m, v59)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
							if v64&int32(2) != 0 {
								return
							} else {
								F_valkeyAsyncFreeInternal(m, l0)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
					if v29 == int32(0) {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
						if v56 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
							if v64&int32(2) != 0 {
								return
							} else {
								F_valkeyAsyncFreeInternal(m, l0)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
							m.T0[v56].(func(*base.Module, int32))(m, v59)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
								v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
								if v64&int32(2) != 0 {
									return
								} else {
									F_valkeyAsyncFreeInternal(m, l0)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v32
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
						if v29 != v34 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
						}
						v38 = m.G4
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
						m.T0[v39].(func(*base.Module, int32))(m, v29)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							v42 = m.G3
							m.Env.X__assert_fail(m, v42+int32(_a1884), v42+int32(_a1885), int32(423), v42+int32(_a1886))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_valkeyAsyncRead(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_valkeyBufferRead(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if v10 != int32(-1) {
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
			v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
			if v63&int32(2) == int32(0) {
				if v62 == int32(0) {
					v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
					if v102 == int32(0) {
						F_valkeyProcessCallbacks(m, l0)
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
						m.T0[v102].(func(*base.Module, int32))(m, v105)
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return
						} else {
							F_valkeyProcessCallbacks(m, l0)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
					if v79 == int32(0) {
						v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
						if v102 == int32(0) {
							F_valkeyProcessCallbacks(m, l0)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						} else {
							v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
							m.T0[v102].(func(*base.Module, int32))(m, v105)
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return
							} else {
								F_valkeyProcessCallbacks(m, l0)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					} else {
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
						if v82 != int64(0) {
							v88 = v79
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
							v90 = int32(8)
							v94 = *(*int64)(unsafe.Add(mBase, uint32(v88+v90)))
							*(*int64)(unsafe.Add(mBase, uint32(v8+v90))) = v94
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = v96
							m.T0[v62].(func(*base.Module, int32, int32))(m, v89, v8)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
								if v102 == int32(0) {
									F_valkeyProcessCallbacks(m, l0)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								} else {
									v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
									m.T0[v102].(func(*base.Module, int32))(m, v105)
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return
									} else {
										F_valkeyProcessCallbacks(m, l0)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									}
								}
							}
						} else {
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
							if v85 == int32(0) {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
								if v102 == int32(0) {
									F_valkeyProcessCallbacks(m, l0)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								} else {
									v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
									m.T0[v102].(func(*base.Module, int32))(m, v105)
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return
									} else {
										F_valkeyProcessCallbacks(m, l0)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									}
								}
							} else {
								v88 = v79
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
								v90 = int32(8)
								v94 = *(*int64)(unsafe.Add(mBase, uint32(v88+v90)))
								*(*int64)(unsafe.Add(mBase, uint32(v8+v90))) = v94
								v96 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = v96
								m.T0[v62].(func(*base.Module, int32, int32))(m, v89, v8)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
									if v102 == int32(0) {
										F_valkeyProcessCallbacks(m, l0)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									} else {
										v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
										m.T0[v102].(func(*base.Module, int32))(m, v105)
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return
										} else {
											F_valkeyProcessCallbacks(m, l0)
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
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
				}
			} else {
				if v62 == int32(0) {
					v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
					if v102 == int32(0) {
						F_valkeyProcessCallbacks(m, l0)
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
						m.T0[v102].(func(*base.Module, int32))(m, v105)
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return
						} else {
							F_valkeyProcessCallbacks(m, l0)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
					if v70 == int32(0) {
						v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
						if v102 == int32(0) {
							F_valkeyProcessCallbacks(m, l0)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						} else {
							v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
							m.T0[v102].(func(*base.Module, int32))(m, v105)
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return
							} else {
								F_valkeyProcessCallbacks(m, l0)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					} else {
						v73 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
						if v73 != int64(0) {
							v88 = v70
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
							v90 = int32(8)
							v94 = *(*int64)(unsafe.Add(mBase, uint32(v88+v90)))
							*(*int64)(unsafe.Add(mBase, uint32(v8+v90))) = v94
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = v96
							m.T0[v62].(func(*base.Module, int32, int32))(m, v89, v8)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
								if v102 == int32(0) {
									F_valkeyProcessCallbacks(m, l0)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								} else {
									v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
									m.T0[v102].(func(*base.Module, int32))(m, v105)
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return
									} else {
										F_valkeyProcessCallbacks(m, l0)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									}
								}
							}
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
							if v76 != 0 {
								v88 = v70
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
								v90 = int32(8)
								v94 = *(*int64)(unsafe.Add(mBase, uint32(v88+v90)))
								*(*int64)(unsafe.Add(mBase, uint32(v8+v90))) = v94
								v96 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = v96
								m.T0[v62].(func(*base.Module, int32, int32))(m, v89, v8)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
									if v102 == int32(0) {
										F_valkeyProcessCallbacks(m, l0)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									} else {
										v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
										m.T0[v102].(func(*base.Module, int32))(m, v105)
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return
										} else {
											F_valkeyProcessCallbacks(m, l0)
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return
											} else {
												m.G0 = v8 + int32(16)
												return
											}
										}
									}
								}
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
								if v102 == int32(0) {
									F_valkeyProcessCallbacks(m, l0)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								} else {
									v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
									m.T0[v102].(func(*base.Module, int32))(m, v105)
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return
									} else {
										F_valkeyProcessCallbacks(m, l0)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
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
			}
		} else {
			if l0 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v19
				v21 = v19
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, 204))
				v21 = v15
			}
			if v21 != 0 {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v44 | int32(4)
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
				if v49 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
					v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
					if v57&int32(2) != 0 {
						m.G0 = v8 + int32(16)
						return
					} else {
						F_valkeyAsyncFreeInternal(m, l0)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
					m.T0[v49].(func(*base.Module, int32))(m, v52)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
						if v57&int32(2) != 0 {
							m.G0 = v8 + int32(16)
							return
						} else {
							F_valkeyAsyncFreeInternal(m, l0)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
				if v22 == int32(0) {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
					if v49 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
						if v57&int32(2) != 0 {
							m.G0 = v8 + int32(16)
							return
						} else {
							F_valkeyAsyncFreeInternal(m, l0)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
						m.T0[v49].(func(*base.Module, int32))(m, v52)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
							if v57&int32(2) != 0 {
								m.G0 = v8 + int32(16)
								return
							} else {
								F_valkeyAsyncFreeInternal(m, l0)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v25
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
					if v22 != v27 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
					}
					v31 = m.G4
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
					m.T0[v32].(func(*base.Module, int32))(m, v22)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v35 = m.G3
						m.Env.X__assert_fail(m, v35+int32(_a1884), v35+int32(_a1885), int32(423), v35+int32(_a1886))
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_valkeyBufferRead(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
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
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	v6 = m.G0
	v8 = v6 - int32(16384)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16384)
	return v202
L2:
	;
	v202 = int32(-1)
	goto L1
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if v12 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v107 = int32(0)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v110 = m.T0[v109].(func(*base.Module, int32, int32, int32) int32)(m, l0, v8, int32(16384))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L36
	}
L5:
	;
	v15 = int32(-1)
	v16 = m.T0[v12].(func(*base.Module, int32, int32) int32)(m, l0, v8)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v16 < int32(0) {
		v202 = v15
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v16 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+28))
	v105 = m.T0[v104].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L6
	} else {
		goto L35
	}
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v26 = F_valkeyReaderFeed(m, v24, v25, v16)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	if v26 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v31
	v34 = v30 + int32(4)
	v36 = l0 + int32(8)
	if v34&int32(3) == int32(0) {
		v58 = v34
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v92 = int32(127)
	if base.Ui32(v91) < base.Ui32(v92) {
		goto L29
	} else {
		goto L30
	}
L14:
	;
	v91 = v83 - v34
	goto L13
L15:
	;
	v62 = v58
	goto L23
L16:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v44 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v47 = v34
	goto L19
L18:
	;
	v91 = v34 - v34
	goto L13
L19:
	;
	v51 = v47 + int32(1)
	if v51&int32(3) == int32(0) {
		v58 = v51
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v56 != 0 {
		v47 = v51
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v83 = v51
	goto L14
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v71 = int32(-2139062144)
	if (int32(16843008)-v68|v68)&v71 == v71 {
		v62 = v62 + int32(4)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v77 = v62
	goto L26
L25:
	;
	goto L24
L26:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v81 != 0 {
		v77 = v77 + int32(1)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v83 = v77
	goto L14
L28:
	;
	goto L27
L29:
	;
	v95 = v91
	goto L31
L30:
	;
	v95 = v92
	goto L31
L31:
	;
	if v95 == int32(0) {
		v99 = v36
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v99+v95))) = uint8(v101)
	v202 = v15
	goto L1
L33:
	;
	goto L32
L34:
	;
	v98 = F__emscripten_memcpy_bulkmem(m, v36, v34, v95)
	mBase = m.M
	v99 = v98
	goto L33
L35:
	;
	v202 = v105
	goto L1
L36:
	;
	if v110 < int32(0) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	if v110 == int32(0) {
		v202 = v107
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v117 = F_valkeyReaderFeed(m, v116, v8, v110)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	if v117 == int32(0) {
		v202 = v107
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v122
	v125 = v121 + int32(4)
	v127 = l0 + int32(8)
	if v125&int32(3) == int32(0) {
		v149 = v125
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v183 = int32(127)
	if base.Ui32(v182) < base.Ui32(v183) {
		goto L57
	} else {
		goto L58
	}
L42:
	;
	v182 = v174 - v125
	goto L41
L43:
	;
	v153 = v149
	goto L51
L44:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if v135 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v138 = v125
	goto L47
L46:
	;
	v182 = v125 - v125
	goto L41
L47:
	;
	v142 = v138 + int32(1)
	if v142&int32(3) == int32(0) {
		v149 = v142
		goto L43
	} else {
		goto L49
	}
L49:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v147 != 0 {
		v138 = v142
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v174 = v142
	goto L42
L51:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v162 = int32(-2139062144)
	if (int32(16843008)-v159|v159)&v162 == v162 {
		v153 = v153 + int32(4)
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v168 = v153
	goto L54
L53:
	;
	goto L52
L54:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v172 != 0 {
		v168 = v168 + int32(1)
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v174 = v168
	goto L42
L56:
	;
	goto L55
L57:
	;
	v186 = v182
	goto L59
L58:
	;
	v186 = v183
	goto L59
L59:
	;
	if v186 == int32(0) {
		v190 = v127
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v192 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v190+v186))) = uint8(v192)
	goto L2
L61:
	;
	goto L60
L62:
	;
	v189 = F__emscripten_memcpy_bulkmem(m, v127, v125, v186)
	mBase = m.M
	v190 = v189
	goto L61
}
func F_valkeyCheckSocketError(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v121 int32
	_ = v121
	v6 = m.G0
	v8 = v6 - int32(144)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
	v12 = int32(9116376)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v14 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v14
	v16 = int32(-1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v24 = F_getsockopt(m, v17, int32(1), v14, v8+int32(12), v8+int32(8))
	mBase = m.M
	if v24 != v16 {
		v78 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		if v78 != 0 {
			v81 = v78
			*(*int32)(unsafe.Add(mBase, _consts[5])) = v81
			v88 = F__emscripten_memset_bulkmem(m, v8+int32(16), base.I32_extend8_s(int32(0)), int32(128))
			mBase = m.M
			v90 = v8 + int32(16)
			v93 = F_strerror(m, v81)
			mBase = m.M
			v94 = F_strlen(m, v93)
			mBase = m.M
			if base.Ui32(v94) < base.Ui32(int32(128)) {
				v108 = F___memcpy(m, v90, v93, v94+int32(1))
				mBase = m.M
			} else {
				v101 = F___memcpy(m, v90, v93, int32(127))
				mBase = m.M
				v103 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(143)))) = uint8(v103)
			}
			F_valkeySetError(m, l0, int32(1), v8+int32(16))
			mBase = m.M
			v121 = int32(-1)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v13
			if v13 != 0 {
				v81 = v13
				*(*int32)(unsafe.Add(mBase, _consts[5])) = v81
				v88 = F__emscripten_memset_bulkmem(m, v8+int32(16), base.I32_extend8_s(int32(0)), int32(128))
				mBase = m.M
				v90 = v8 + int32(16)
				v93 = F_strerror(m, v81)
				mBase = m.M
				v94 = F_strlen(m, v93)
				mBase = m.M
				if base.Ui32(v94) < base.Ui32(int32(128)) {
					v108 = F___memcpy(m, v90, v93, v94+int32(1))
					mBase = m.M
				} else {
					v101 = F___memcpy(m, v90, v93, int32(127))
					mBase = m.M
					v103 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(143)))) = uint8(v103)
				}
				F_valkeySetError(m, l0, int32(1), v8+int32(16))
				mBase = m.M
				v121 = int32(-1)
			} else {
				v121 = int32(0)
			}
		}
		m.G0 = v8 + int32(144)
		return v121
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, _consts[5]))
		v33 = F__emscripten_memset_bulkmem(m, v8+int32(16), base.I32_extend8_s(int32(0)), int32(128))
		mBase = m.M
		v34 = m.G3
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v34 + int32(_a1894)
		v39 = v8 + int32(16)
		v45 = F_snprintf(m, v39, int32(128), v34+int32(_a1895), v8)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			v49 = v39 + v45
			v51 = int32(128) - v45
			v53 = F_strerror(m, v27)
			mBase = m.M
			v54 = F_strlen(m, v53)
			mBase = m.M
			if base.Ui32(v54) < base.Ui32(v51) {
				v68 = F___memcpy(m, v49, v53, v54+int32(1))
				mBase = m.M
			} else {
				if v51 == int32(0) {
				} else {
					v60 = v51 + int32(-1)
					v61 = F___memcpy(m, v49, v53, v60)
					mBase = m.M
					v63 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v49+v60))) = uint8(v63)
				}
			}
			F_valkeySetError(m, l0, int32(1), v8+int32(16))
			mBase = m.M
			v121 = v16
			m.G0 = v8 + int32(144)
			return v121
		}
	}
}
func F_valkeyContextConnectTcp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v328 int32
	_ = v328
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v529 int32
	_ = v529
	var v538 int32
	_ = v538
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v613 int32
	_ = v613
	var v635 int32
	_ = v635
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v653 int32
	_ = v653
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(256)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+204)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v3
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v19 == v27 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v15 + int32(256)
	return v653
L2:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v15)+204))
	if v642 == int32(0) {
		v653 = v635
		goto L1
	} else {
		goto L170
	}
L3:
	;
	v635 = int32(-1)
	goto L2
L4:
	;
	v613 = m.G3
	F_valkeySetError(m, l0, int32(5), v613+int32(_a1898))
	mBase = m.M
	goto L3
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v18 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v29 = m.G4
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	m.T0[v30].(func(*base.Module, int32))(m, v27)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v36 = m.T0[v35].(func(*base.Module, int32) int32)(m, v19)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v36
	if v36 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L5
L11:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v17 != 0 {
		goto L31
	} else {
		goto L32
	}
L12:
	;
	v102 = int32(2147483647)
	goto L11
L13:
	;
	v92 = m.G4
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	m.T0[v93].(func(*base.Module, int32))(m, v43)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L28
	}
L14:
	;
	if v43 == v18 {
		v65 = v43
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v65 == int32(0) {
		goto L12
	} else {
		goto L21
	}
L16:
	;
	if v43 != 0 {
		v55 = v43
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v56
	v58 = int32(8)
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v18+v58)))
	*(*int64)(unsafe.Add(mBase, uint32(v55+v58))) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v65 = v64
	goto L15
L18:
	;
	v48 = m.G4
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v50 = m.T0[v49].(func(*base.Module, int32) int32)(m, int32(16))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v50
	if v50 == int32(0) {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v55 = v50
	goto L17
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if int32(1000000) < v68 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v88 = m.G3
	F_valkeySetError(m, l0, int32(1), v88+int32(_a1899))
	mBase = m.M
	goto L3
L23:
	;
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	if int64(2147482) < v71 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v75 = int32(1000)
	v80 = base.I32_div_s(v68+int32(999), v75)
	v81 = base.I32_wrap_i64(v71)*v75 + v80
	v82 = int32(2147483647)
	if base.Ui32(v81) < base.Ui32(v82) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v85 = v81
	goto L27
L26:
	;
	v85 = v82
	goto L27
L27:
	;
	v102 = v85
	goto L11
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(0)
	goto L12
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v20
	v127 = m.G3
	v132 = F_snprintf(m, v15+int32(246), int32(6), v127+int32(_a77), v15+int32(48))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L7
	} else {
		goto L37
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v118
	goto L29
L31:
	;
	if v104 == v17 {
		goto L29
	} else {
		goto L34
	}
L32:
	;
	v105 = m.G4
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	m.T0[v106].(func(*base.Module, int32))(m, v104)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v118 = int32(0)
	goto L30
L34:
	;
	v111 = m.G4
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	m.T0[v112].(func(*base.Module, int32))(m, v104)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v116 = m.T0[v115].(func(*base.Module, int32) int32)(m, v17)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	v118 = v116
	goto L30
L37:
	;
	v135 = v15 + int32(216)
	v136 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v135))) = v136
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(232)))) = v136
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(224)))) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+208)) = v136
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v153&int32(4096) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v156 = int32(10)
	goto L40
L39:
	;
	v156 = int32(2)
	goto L40
L40:
	;
	v157 = int32(6144)
	if v153&v157 == v157 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v161 = int32(0)
	goto L43
L42:
	;
	v161 = v156
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+212)) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v170 = m.Env.Getaddrinfo(m, v163, v15+int32(246), v15+int32(208), v15+int32(204))
	mBase = m.M
	if v170 == int32(0) {
		v190 = v170
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if v190 != 0 {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v15)+212))
	if v173 == int32(0) {
		v190 = v170
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v177 = int32(2)
	if v173 == v177 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v180 = int32(10)
	goto L49
L48:
	;
	v180 = v177
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+212)) = v180
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v189 = m.Env.Getaddrinfo(m, v182, v15+int32(246), v15+int32(208), v15+int32(204))
	mBase = m.M
	v190 = v189
	goto L44
L50:
	;
	v570 = int32(_a78)
	v572 = v190 + int32(1)
	if v572 == int32(0) {
		v592 = v570
		goto L161
	} else {
		goto L162
	}
L51:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v15)+204))
	if v192 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	goto L157
L53:
	;
	v196 = v21 & int32(128)
	v198 = v21 & int32(1)
	v203 = int32(0)
	v206 = v192
	goto L54
L54:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v216&int32(8192) != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L52
L56:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v206)+28))
	if v538 != 0 {
		v203 = v529
		v206 = v538
		goto L54
	} else {
		goto L156
	}
L57:
	;
	v219 = int32(262)
	goto L59
L58:
	;
	v219 = int32(6)
	goto L59
L59:
	;
	v220 = F_socket(m, v212, v213, v219)
	mBase = m.M
	if v220 == int32(-1) {
		v529 = v203
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v223 = int32(9)
	if v223 < v203 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v226 = v203
	goto L63
L62:
	;
	v226 = v223
	goto L63
L63:
	;
	v230 = v203
	v231 = v220
	goto L67
L64:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v521 | int32(2)
	v635 = int32(0)
	goto L2
L65:
	;
	if v198 == int32(0) {
		goto L64
	} else {
		goto L153
	}
L66:
	;
	v510 = F_valkeyContextWaitReady(m, l0, v102)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L7
	} else {
		goto L149
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v231
	v240 = int32(-1)
	v242 = F_valkeySetBlocking(m, l0, int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L7
	} else {
		goto L69
	}
L68:
	;
	if v198 == int32(0) {
		goto L64
	} else {
		goto L148
	}
L69:
	;
	if v242 != 0 {
		v635 = v240
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v244 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L71:
	;
	goto L68
L72:
	;
	if v441 != int32(4) {
		goto L66
	} else {
		goto L139
	}
L73:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v478 == int32(-1) {
		v529 = v230
		goto L56
	} else {
		goto L138
	}
L74:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v446)+20))
	F_emscripten_builtin_free(m, v457)
	mBase = m.M
	F_emscripten_builtin_free(m, v446)
	mBase = m.M
	goto L134
L75:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v15)+200))
	v446 = v444
	goto L74
L76:
	;
	v416 = m.G4
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v416)+16))
	m.T0[v418].(func(*base.Module, int32))(m, v417)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L7
	} else {
		goto L126
	}
L77:
	;
	v247 = int32(0)
	v252 = m.Env.Getaddrinfo(m, v244, v247, v15+int32(208), v15+int32(200))
	mBase = m.M
	if v252 == v247 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	if v196 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L79:
	;
	v257 = int32(_a78)
	v259 = v252 + int32(1)
	if v259 == int32(0) {
		v279 = v257
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v279 + base.B2i32(v281 == int32(0))
	v289 = m.G3
	v294 = F_snprintf(m, v15+int32(64), int32(128), v289+int32(_a1900), v15+int32(32))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L7
	} else {
		goto L90
	}
L81:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	goto L80
L82:
	;
	v263 = v257
	v264 = v259
	goto L83
L83:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	if v265 == int32(0) {
		v279 = v263
		goto L81
	} else {
		goto L85
	}
L84:
	;
	v279 = v275
	goto L81
L85:
	;
	v269 = v263
	goto L86
L86:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+1)))
	if v273 != 0 {
		v269 = v269 + int32(1)
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v275 = v269 + int32(2)
	v277 = v264 + int32(1)
	if v277 != 0 {
		v263 = v275
		v264 = v277
		goto L83
	} else {
		goto L89
	}
L88:
	;
	goto L87
L89:
	;
	goto L84
L90:
	;
	F_valkeySetError(m, l0, int32(2), v15+int32(64))
	mBase = m.M
	v635 = v240
	goto L2
L91:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v15)+200))
	if v378 != 0 {
		goto L118
	} else {
		goto L119
	}
L92:
	;
	v302 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v302
	v305 = int32(2)
	v307 = v15 + int32(252)
	v308 = int32(4)
	v312 = m.G0
	v314 = v312 - int32(16)
	m.G0 = v314
	v317 = F___syscall_setsockopt(m, v231, v302, v305, v307, v308, int32(0))
	mBase = m.M
	goto L96
L93:
	;
	if int32(-1) < v367 {
		goto L91
	} else {
		goto L116
	}
L94:
	;
	m.G0 = v314 + int32(16)
	goto L93
L95:
	;
	v366 = F___syscall_ret(m, v364)
	mBase = m.M
	v367 = v366
	goto L94
L96:
	;
	if v317 != int32(-50) {
		v364 = v317
		goto L95
	} else {
		goto L97
	}
L97:
	;
	switch int32(-61) {
	case 0, 1:
		goto L98
	default:
		v364 = int32(-50)
		goto L95
	case 3, 4:
		goto L99
	}
L98:
	;
	goto L111
L99:
	;
	goto L101
L101:
	;
	v328 = F___syscall_ret(m, int32(-28))
	mBase = m.M
	v367 = v328
	goto L94
L111:
	;
	goto L112
L112:
	;
	goto L114
L114:
	;
	goto L115
L115:
	;
	v363 = F___syscall_setsockopt(m, v231, int32(1), v305, v307, v308, int32(0))
	mBase = m.M
	v364 = v363
	goto L95
L116:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v15)+200))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)+20))
	F_emscripten_builtin_free(m, v375)
	mBase = m.M
	F_emscripten_builtin_free(m, v374)
	mBase = m.M
	goto L117
L117:
	;
	v635 = v240
	goto L2
L118:
	;
	v381 = v378
	goto L121
L119:
	;
	v446 = int32(0)
	goto L74
L120:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v15)+200))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+20))
	F_emscripten_builtin_free(m, v401)
	mBase = m.M
	F_emscripten_builtin_free(m, v400)
	mBase = m.M
	goto L125
L121:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v381)+20))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v381)+16))
	v394 = F_bind(m, v231, v392, v393)
	mBase = m.M
	if v394 != int32(-1) {
		goto L120
	} else {
		goto L123
	}
L123:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v381)+28))
	if v397 == int32(0) {
		goto L75
	} else {
		goto L124
	}
L124:
	;
	v381 = v397
	goto L121
L125:
	;
	goto L76
L126:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v206)+16))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	v423 = m.T0[v422].(func(*base.Module, int32) int32)(m, v421)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v423
	if v423 == int32(0) {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v206)+20))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v206)+16))
	if v429 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v206)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v434
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v206)+20))
	v437 = F_connect(m, v231, v436, v434)
	mBase = m.M
	if v437 != int32(-1) {
		goto L65
	} else {
		goto L132
	}
L130:
	;
	goto L129
L131:
	;
	v432 = F__emscripten_memcpy_bulkmem(m, v423, v428, v429)
	mBase = m.M
	goto L130
L132:
	;
	goto L133
L133:
	;
	v441 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	switch v441 + int32(-23) {
	case 0:
		goto L73
	default:
		goto L72
	case 3:
		goto L71
	}
L134:
	;
	goto L135
L135:
	;
	v461 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v462 = F___strerror_l(m, v461, v461)
	mBase = m.M
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v462
	v467 = m.G3
	v472 = F_snprintf(m, v15+int32(64), int32(128), v467+int32(_a1901), v15+int32(16))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L7
	} else {
		goto L137
	}
L137:
	;
	F_valkeySetError(m, l0, int32(2), v15+int32(64))
	mBase = m.M
	v635 = v240
	goto L2
L138:
	;
	v481 = F_close(m, v478)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = int32(-1)
	v529 = v230
	goto L56
L139:
	;
	if v196 == int32(0) {
		goto L66
	} else {
		goto L140
	}
L140:
	;
	if v230 == v226 {
		v635 = v240
		goto L2
	} else {
		goto L141
	}
L141:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v489 == int32(-1) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v496 = v230 + int32(1)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v501&int32(8192) != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v492 = F_close(m, v489)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = int32(-1)
	goto L142
L144:
	;
	v504 = int32(262)
	goto L146
L145:
	;
	v504 = int32(6)
	goto L146
L146:
	;
	v505 = F_socket(m, v497, v498, v504)
	mBase = m.M
	if v505 != int32(-1) {
		v230 = v496
		v231 = v505
		goto L67
	} else {
		goto L147
	}
L147:
	;
	v529 = v496
	goto L56
L148:
	;
	goto L66
L149:
	;
	if v510 != 0 {
		v635 = v240
		goto L2
	} else {
		goto L150
	}
L150:
	;
	v512 = F_valkeySetTcpNoDelay(m, l0)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L7
	} else {
		goto L151
	}
L151:
	;
	if v512 != 0 {
		v635 = v240
		goto L2
	} else {
		goto L152
	}
L152:
	;
	goto L65
L153:
	;
	v518 = F_valkeySetBlocking(m, l0, int32(1))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L7
	} else {
		goto L154
	}
L154:
	;
	if v518 != 0 {
		v635 = v240
		goto L2
	} else {
		goto L155
	}
L155:
	;
	goto L64
L156:
	;
	goto L55
L157:
	;
	v552 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v553 = F___strerror_l(m, v552, v552)
	mBase = m.M
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v553
	v558 = m.G3
	v561 = F_snprintf(m, v15+int32(64), int32(128), v558+int32(_a1902), v15)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L7
	} else {
		goto L159
	}
L159:
	;
	F_valkeySetError(m, l0, int32(2), v15+int32(64))
	mBase = m.M
	goto L3
L160:
	;
	F_valkeySetError(m, l0, int32(2), v592+base.B2i32(v594 == int32(0)))
	mBase = m.M
	v653 = int32(-1)
	goto L1
L161:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592))))
	goto L160
L162:
	;
	v576 = v570
	v577 = v572
	goto L163
L163:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576))))
	if v578 == int32(0) {
		v592 = v576
		goto L161
	} else {
		goto L165
	}
L164:
	;
	v592 = v588
	goto L161
L165:
	;
	v582 = v576
	goto L166
L166:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+1)))
	if v586 != 0 {
		v582 = v582 + int32(1)
		goto L166
	} else {
		goto L168
	}
L167:
	;
	v588 = v582 + int32(2)
	v590 = v577 + int32(1)
	if v590 != 0 {
		v576 = v588
		v577 = v590
		goto L163
	} else {
		goto L169
	}
L168:
	;
	goto L167
L169:
	;
	goto L164
L170:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v642)+20))
	F_emscripten_builtin_free(m, v645)
	mBase = m.M
	F_emscripten_builtin_free(m, v642)
	mBase = m.M
	goto L171
L171:
	;
	v653 = v635
	goto L1
}
func F_valkeyContextRegisterFuncs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		v15 = m.G3
		m.Env.X__assert_fail(m, v15+int32(_a1889), v15+int32(_a1890), int32(85), v15+int32(_a1891))
		mBase = m.M
		base.Wasm_trap_unreachable()
		for {
		}
	} else {
		v5 = m.G3
		v10 = v5 + int32(_a1892) + l1<<(uint(int32(2))%32)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		if v11 != 0 {
			v24 = m.G3
			m.Env.X__assert_fail(m, v24+int32(_a1893), v24+int32(_a1890), int32(86), v24+int32(_a1891))
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
			return int32(0)
		}
	}
}
func F_valkeyContextRegisterUnixFuncs(m *base.Module) {
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	v1 = m.G3
	v5 = F_valkeyContextRegisterFuncs(m, v1+int32(_a1904), int32(1))
	return
}
func F_valkeyContextWaitReady(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	v12 = int32(9116376)
	goto L5
L1:
	;
	m.G0 = v10 + int32(144)
	return v313
L2:
	;
	v306 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v306
	v313 = v306
	goto L1
L3:
	;
	v313 = int32(-1)
	goto L1
L4:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v59 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v58
	v62 = int32(0)
	if l1 < v62 {
		v77 = v62
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v13 == int32(26) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v21 = F__emscripten_memset_bulkmem(m, v10+int32(16), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L7
L7:
	;
	v23 = v10 + int32(16)
	v26 = F_strerror(m, v13)
	mBase = m.M
	v27 = F_strlen(m, v26)
	mBase = m.M
	if base.Ui32(v27) < base.Ui32(int32(128)) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	F_valkeySetError(m, l0, int32(1), v10+int32(16))
	mBase = m.M
	if l0 == int32(0) {
		goto L3
	} else {
		goto L13
	}
L9:
	;
	goto L8
L10:
	;
	v41 = F___memcpy(m, v23, v26, v27+int32(1))
	mBase = m.M
	goto L9
L11:
	;
	goto L12
L12:
	;
	v34 = F___memcpy(m, v23, v26, int32(127))
	mBase = m.M
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(143)))) = uint8(v36)
	goto L8
L13:
	;
	v53 = int32(-1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v54 == v53 {
		v313 = v53
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v57 = F_close(m, v54)
	mBase = m.M
	goto L2
L15:
	;
	v81 = F_poll(m, v10+int32(8), int32(1), l1)
	mBase = m.M
	if int32(0) < v81 {
		v223 = v81
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v68 = F___clock_gettime(m, int32(1), v10+int32(16))
	mBase = m.M
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v71 = base.I32_div_s(v69, int32(1000000))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v77 = v71 + l1 + v73*int32(1000)
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v223
	v227 = int32(16)
	v228 = v10 + v227
	v233 = m.G0
	v235 = v233 - v227
	m.G0 = v235
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v240 = F_connect(m, v237, v238, v239)
	mBase = m.M
	if v240 != 0 {
		goto L50
	} else {
		goto L51
	}
L18:
	;
	v90 = v81
	goto L19
L19:
	;
	if int32(-1) < v90 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v223 = v216
	goto L17
L21:
	;
	v215 = int32(1)
	v216 = F_poll(m, v10+int32(8), v215, l1)
	mBase = m.M
	if v216 < v215 {
		v90 = v216
		goto L19
	} else {
		goto L46
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(73)
	v175 = F__emscripten_memset_bulkmem(m, v10+int32(16), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L38
L23:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v95 == int32(27) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if l1 < int32(0) {
		goto L21
	} else {
		goto L36
	}
L25:
	;
	v103 = F__emscripten_memset_bulkmem(m, v10+int32(16), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L26
L26:
	;
	v104 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v104 + int32(_a1903)
	v109 = v10 + int32(16)
	v115 = F_snprintf(m, v109, int32(128), v104+int32(_a1895), v10)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	v119 = v109 + v115
	v121 = int32(128) - v115
	v123 = F_strerror(m, v95)
	mBase = m.M
	v124 = F_strlen(m, v123)
	mBase = m.M
	if base.Ui32(v124) < base.Ui32(v121) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	F_valkeySetError(m, l0, int32(1), v10+int32(16))
	mBase = m.M
	if l0 == int32(0) {
		goto L3
	} else {
		goto L34
	}
L30:
	;
	goto L29
L31:
	;
	v138 = F___memcpy(m, v119, v123, v124+int32(1))
	mBase = m.M
	goto L30
L32:
	;
	if v121 == int32(0) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v130 = v121 + int32(-1)
	v131 = F___memcpy(m, v119, v123, v130)
	mBase = m.M
	v133 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v119+v130))) = uint8(v133)
	goto L29
L34:
	;
	v150 = int32(-1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v151 == v150 {
		v313 = v150
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v154 = F_close(m, v151)
	mBase = m.M
	goto L2
L36:
	;
	v158 = F___clock_gettime(m, int32(1), v10+int32(16))
	mBase = m.M
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v164 = base.I32_div_s(v162, int32(1000000))
	if v159*int32(1000)+v164 < v77 {
		goto L21
	} else {
		goto L37
	}
L37:
	;
	goto L22
L38:
	;
	v178 = v10 + int32(16)
	v181 = F_strerror(m, int32(73))
	mBase = m.M
	v182 = F_strlen(m, v181)
	mBase = m.M
	if base.Ui32(v182) < base.Ui32(int32(128)) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	F_valkeySetError(m, l0, int32(1), v10+int32(16))
	mBase = m.M
	if l0 == int32(0) {
		goto L3
	} else {
		goto L44
	}
L40:
	;
	goto L39
L41:
	;
	v196 = F___memcpy(m, v178, v181, v182+int32(1))
	mBase = m.M
	goto L40
L42:
	;
	goto L43
L43:
	;
	v189 = F___memcpy(m, v178, v181, int32(127))
	mBase = m.M
	v191 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(143)))) = uint8(v191)
	goto L39
L44:
	;
	v208 = int32(-1)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v209 == v208 {
		v313 = v208
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v212 = F_close(m, v209)
	mBase = m.M
	goto L2
L46:
	;
	goto L20
L47:
	;
	v289 = F_valkeyCheckSocketError(m, l0)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L27
	} else {
		goto L62
	}
L48:
	;
	if v278 != 0 {
		goto L47
	} else {
		goto L60
	}
L49:
	;
	m.G0 = v235 + int32(16)
	goto L48
L50:
	;
	v244 = int32(26)
	v245 = F___errno_location(m)
	mBase = m.M
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	if v246 != v244 {
		v265 = v246
		goto L53
	} else {
		goto L54
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = int32(1)
	v278 = int32(0)
	goto L49
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = int32(1)
	v278 = int32(0)
	goto L49
L53:
	;
	switch v265 + int32(-6) {
	case 0, 1:
		goto L58
	default:
		v278 = int32(-1)
		goto L49
	case 24:
		goto L59
	}
L54:
	;
	v249 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v235)+8)) = v249
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v258 = F_getsockopt(m, v251, int32(1), v249, v235+int32(12), v235+int32(8))
	mBase = m.M
	if v258 != 0 {
		v263 = v244
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v265 = v263
	goto L53
L56:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v235)+12))
	if v259 == int32(0) {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v259
	v263 = v259
	goto L55
L58:
	;
	v272 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v272
	v278 = v272
	goto L49
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = int32(1)
	v278 = int32(0)
	goto L49
L60:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v285 == int32(0) {
		goto L47
	} else {
		goto L61
	}
L61:
	;
	v313 = int32(0)
	goto L1
L62:
	;
	goto L3
}
func F_valkeyNetClose(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	if l0 == int32(0) {
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
		if v5 == int32(-1) {
		} else {
			v8 = F_close(m, v5)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = int32(-1)
		}
	}
	return
}
func F_valkeyReaderFeed(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v147 int64
	_ = v147
	var v153 int64
	_ = v153
	var v156 int32
	_ = v156
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 != 0 {
		v156 = int32(-1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v156
L2:
	;
	v9 = int32(0)
	if l1 == v9 {
		v156 = v9
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l2 == int32(0) {
		v156 = v9
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v116 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L6:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v72 = F_sdscatlen(m, v71, l1, l2)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L15
	} else {
		goto L19
	}
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v15 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v19 = int32(-1)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v19))))
	switch v21&int32(7) + v19 {
	case 0:
		goto L13
	case 1:
		goto L12
	case 2:
		goto L11
	case 3:
		goto L10
	default:
		goto L6
	}
L9:
	;
	if base.Ui32(v55) <= base.Ui32(v15) {
		goto L6
	} else {
		goto L14
	}
L10:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v18+int32(-9))))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v18+int32(-17))))
	v55 = base.I32_wrap_i64(v49 - v52)
	goto L9
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(-5))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(-9))))
	v55 = v42 - v45
	goto L9
L12:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18+int32(-3)))))
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18+int32(-5)))))
	v55 = v35 - v38
	goto L9
L13:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(-2)))))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(-3)))))
	v55 = v28 - v31
	goto L9
L14:
	;
	F_sdsfree(m, v18)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	v61 = F_sdsempty(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v61
	if v61 == int32(0) {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = int32(0)
	goto L6
L19:
	;
	if v72 == int32(0) {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v72
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+int32(-1)))))
	switch v80 & int32(7) {
	case 0:
		goto L26
	case 1:
		goto L25
	case 2:
		goto L24
	case 3:
		goto L23
	case 4:
		goto L22
	default:
		v109 = int32(0)
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v109
	return int32(0)
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v72+int32(-17))))
	v109 = v108
	goto L21
L23:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v72+int32(-9))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v102
	return int32(0)
L24:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72+int32(-5)))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v96
	return int32(0)
L25:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+int32(-3)))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v90
	return int32(0)
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(base.Ui32(v80) >> (uint(int32(3)) % 32))
	return int32(0)
L27:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_sdsfree(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L15
	} else {
		goto L32
	}
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v119 == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+24))
	if v122 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	m.T0[v122].(func(*base.Module, int32))(m, v116)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L15
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	goto L27
L32:
	;
	v133 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v133
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v137 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(5)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v133)
	v144 = m.G3
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v144)+uint32(_consts[1002])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v147
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v144)+uint32(_consts[1003])))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(9)))) = v153
	v156 = v137
	goto L1
}
func F_valkeyTcpSetTimeout(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v35 int64
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v156 int64
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v256 int32
	_ = v256
	v6 = m.G0
	v8 = v6 - int32(160)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v13 = int32(66)
	v14 = int32(16)
	v18 = m.G0
	v20 = v18 - v14
	m.G0 = v20
	v23 = F___syscall_setsockopt(m, v11, int32(1), v13, l1, v14, int32(0))
	mBase = m.M
	if v23 != int32(-50) {
		v70 = v23
		v72 = F___syscall_ret(m, v70)
		mBase = m.M
		v73 = v72
	} else {
		switch int32(3) {
		case 0, 1:
			v69 = F___syscall_setsockopt(m, v11, int32(1), v13, l1, v14, int32(0))
			mBase = m.M
			v70 = v69
			v72 = F___syscall_ret(m, v70)
			mBase = m.M
			v73 = v72
		default:
			v70 = int32(-50)
			v72 = F___syscall_ret(m, v70)
			mBase = m.M
			v73 = v72
		case 3, 4:
			v35 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			if base.Ui64(v35+int64(2147483648)) < base.Ui64(int64(4294967296)) {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v42
				*(*uint32)(unsafe.Add(mBase, uint32(v20)+8)) = uint32(v35)
				v54 = int32(8)
				v58 = F___syscall_setsockopt(m, v11, int32(1), int32(20), v20+v54, v54, int32(0))
				mBase = m.M
				v70 = v58
				v72 = F___syscall_ret(m, v70)
				mBase = m.M
				v73 = v72
			} else {
				v41 = F___syscall_ret(m, int32(-138))
				mBase = m.M
				v73 = v41
			}
		}
	}
	m.G0 = v20 + int32(16)
	if v73 != int32(-1) {
		v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
		v134 = int32(67)
		v135 = int32(16)
		v139 = m.G0
		v141 = v139 - v135
		m.G0 = v141
		v144 = F___syscall_setsockopt(m, v132, int32(1), v134, l1, v135, int32(0))
		mBase = m.M
		if v144 != int32(-50) {
			v191 = v144
			v193 = F___syscall_ret(m, v191)
			mBase = m.M
			v194 = v193
		} else {
			switch int32(4) {
			case 0, 1:
				v190 = F___syscall_setsockopt(m, v132, int32(1), v134, l1, v135, int32(0))
				mBase = m.M
				v191 = v190
				v193 = F___syscall_ret(m, v191)
				mBase = m.M
				v194 = v193
			default:
				v191 = int32(-50)
				v193 = F___syscall_ret(m, v191)
				mBase = m.M
				v194 = v193
			case 3, 4:
				v156 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				if base.Ui64(v156+int64(2147483648)) < base.Ui64(int64(4294967296)) {
					v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v141)+12)) = v163
					*(*uint32)(unsafe.Add(mBase, uint32(v141)+8)) = uint32(v156)
					v175 = int32(8)
					v179 = F___syscall_setsockopt(m, v132, int32(1), int32(21), v141+v175, v175, int32(0))
					mBase = m.M
					v191 = v179
					v193 = F___syscall_ret(m, v191)
					mBase = m.M
					v194 = v193
				} else {
					v162 = F___syscall_ret(m, int32(-138))
					mBase = m.M
					v194 = v162
				}
			}
		}
		m.G0 = v141 + int32(16)
		if v194 != int32(-1) {
			v256 = int32(0)
			m.G0 = v8 + int32(160)
			return v256
		} else {
			v203 = *(*int32)(unsafe.Add(mBase, _consts[5]))
			v209 = F__emscripten_memset_bulkmem(m, v8+int32(32), base.I32_extend8_s(int32(0)), int32(128))
			mBase = m.M
			v210 = m.G3
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v210 + int32(_a1896)
			v215 = v8 + int32(32)
			v223 = F_snprintf(m, v215, int32(128), v210+int32(_a1895), v8+int32(16))
			mBase = m.M
			v224 = m.ExcPending
			if v224 != 0 {
				return int32(0)
			} else {
				v225 = v215 + v223
				v227 = int32(128) - v223
				v229 = F_strerror(m, v203)
				mBase = m.M
				v230 = F_strlen(m, v229)
				mBase = m.M
				if base.Ui32(v230) < base.Ui32(v227) {
					v244 = F___memcpy(m, v225, v229, v230+int32(1))
					mBase = m.M
				} else {
					if v227 == int32(0) {
					} else {
						v236 = v227 + int32(-1)
						v237 = F___memcpy(m, v225, v229, v236)
						mBase = m.M
						v239 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v225+v236))) = uint8(v239)
					}
				}
				F_valkeySetError(m, l0, int32(1), v8+int32(32))
				mBase = m.M
				v256 = int32(-1)
				m.G0 = v8 + int32(160)
				return v256
			}
		}
	} else {
		v81 = *(*int32)(unsafe.Add(mBase, _consts[5]))
		v87 = F__emscripten_memset_bulkmem(m, v8+int32(32), base.I32_extend8_s(int32(0)), int32(128))
		mBase = m.M
		v88 = m.G3
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v88 + int32(_a1897)
		v93 = v8 + int32(32)
		v99 = F_snprintf(m, v93, int32(128), v88+int32(_a1895), v8)
		mBase = m.M
		v102 = m.ExcPending
		if v102 != 0 {
			return int32(0)
		} else {
			v103 = v93 + v99
			v105 = int32(128) - v99
			v107 = F_strerror(m, v81)
			mBase = m.M
			v108 = F_strlen(m, v107)
			mBase = m.M
			if base.Ui32(v108) < base.Ui32(v105) {
				v122 = F___memcpy(m, v103, v107, v108+int32(1))
				mBase = m.M
			} else {
				if v105 == int32(0) {
				} else {
					v114 = v105 + int32(-1)
					v115 = F___memcpy(m, v103, v107, v114)
					mBase = m.M
					v117 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v103+v114))) = uint8(v117)
				}
			}
			F_valkeySetError(m, l0, int32(1), v8+int32(32))
			mBase = m.M
			v256 = int32(-1)
			m.G0 = v8 + int32(160)
			return v256
		}
	}
}
func F_valkey_calloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	if base.Ui32(int32(2147483646)) < base.Ui32(l0) {
		v81 = *(*int32)(unsafe.Add(mBase, _consts[1001]))
		m.T0[v81].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v85 = m.ExcPending
		if v85 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		v7 = int32(1)
		if l0 != 0 {
			v9 = l0
		} else {
			v9 = int32(4)
		}
		v11 = v9 + int32(8)
		v17 = base.I64_extend_i32_u(v7) * base.I64_extend_i32_u(v11)
		v18 = base.I32_wrap_i64(v17)
		if base.Ui32(v11|v7) < base.Ui32(int32(65536)) {
			v29 = v18
		} else {
			if base.I32_wrap_i64(int64(base.Ui64(v17)>>(uint(int64(32))%64))) != int32(0) {
				v28 = int32(-1)
			} else {
				v28 = v18
			}
			v29 = v28
		}
		v31 = F_emscripten_builtin_malloc(m, v29)
		mBase = m.M
		if v31 == int32(0) {
		} else {
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+int32(-4)))))
			if v36&int32(3) == int32(0) {
			} else {
				v42 = F___memset(m, v31, int32(0), v29)
				mBase = m.M
			}
		}
		if v31 == int32(0) {
			v81 = *(*int32)(unsafe.Add(mBase, _consts[1001]))
			m.T0[v81].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v31))) = v9
			v47 = *(*int32)(unsafe.Add(mBase, _consts[411]))
			if v47 != int32(-1) {
				v58 = v47
			} else {
				v50 = int32(0)
				v52 = *(*int32)(unsafe.Add(mBase, _consts[281]))
				*(*int32)(unsafe.Add(mBase, _consts[411])) = v52
				*(*int32)(unsafe.Add(mBase, _consts[281])) = v52 + int32(1)
				v58 = v52
			}
			if v58 < int32(260) {
				v67 = v58 << (uint(int32(2)) % 32)
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[285])))
				*(*int32)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[285]))) = v70 + v11
			} else {
				v61 = int32(0)
				v63 = *(*int32)(unsafe.Add(mBase, _consts[286]))
				*(*int32)(unsafe.Add(mBase, _consts[286])) = v63 + v11
			}
			return v31 + int32(8)
		}
	}
}
func F_valkey_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
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
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var __phi147 int32
	_ = __phi147
	var v148 int32
	_ = v148
	var __phi148 int32
	_ = __phi148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var __phi314 int32
	_ = __phi314
	var v315 int32
	_ = v315
	var __phi315 int32
	_ = __phi315
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v396 int32
	_ = v396
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v560 int32
	_ = v560
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a1860), int32(_a1879), int32(463))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L116
	} else {
		goto L117
	}
L2:
	;
	return
L3:
	;
	v7 = l0 + int32(-8)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if int32(-1) < v8 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v77 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[411]))
	if v48 != int32(-1) {
		v59 = v48
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-12))))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[411]))
	if v19 != int32(-1) {
		v30 = v19
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v32 = v8&int32(2147483647) + int32(8)
	if v30 < int32(260) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v22 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	*(*int32)(unsafe.Add(mBase, _consts[411])) = v24
	*(*int32)(unsafe.Add(mBase, _consts[281])) = v24 + int32(1)
	v30 = v24
	goto L8
L10:
	;
	v41 = v30 << (uint(int32(2)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[285])))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[285]))) = v44 - v32
	v77 = v13
	goto L4
L11:
	;
	v35 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	*(*int32)(unsafe.Add(mBase, _consts[286])) = v37 - v32
	v77 = v13
	goto L4
L12:
	;
	v61 = v8 + int32(8)
	if v59 < int32(260) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v51 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	*(*int32)(unsafe.Add(mBase, _consts[411])) = v53
	*(*int32)(unsafe.Add(mBase, _consts[281])) = v53 + int32(1)
	v59 = v53
	goto L12
L14:
	;
	v70 = v59 << (uint(int32(2)) % 32)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[285])))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[285]))) = v73 - v61
	v77 = v7
	goto L4
L15:
	;
	v64 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	*(*int32)(unsafe.Add(mBase, _consts[286])) = v66 - v61
	v77 = v7
	goto L4
L16:
	;
	goto L2
L17:
	;
	goto L16
L18:
	;
	v89 = int32(-8)
	v90 = v77 + v89
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v77+int32(-4))))
	v95 = v93 & v89
	v96 = v90 + v95
	if v93&int32(1) != 0 {
		v220 = v95
		v221 = v90
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if base.Ui32(v96) <= base.Ui32(v221) {
		goto L17
	} else {
		goto L54
	}
L20:
	;
	if v93&int32(2) == int32(0) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v104 = v90 - v103
	v106 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	if base.Ui32(v104) < base.Ui32(v106) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v108 = v103 + v95
	v110 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v104 == v110 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	if v126 == int32(0) {
		v220 = v108
		v221 = v104
		goto L19
	} else {
		goto L42
	}
L24:
	;
	v179 = int32(0)
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+12)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v115
	v220 = v108
	v221 = v104
	goto L19
L26:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v161 = int32(3)
	if v160&v161 != v161 {
		v220 = v108
		v221 = v104
		goto L19
	} else {
		goto L41
	}
L27:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	if base.Ui32(int32(255)) < base.Ui32(v103) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v104)+24))
	if v112 == v104 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	if v112 != v115 {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v117 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v119 & base.I32_rotl(int32(-2), int32(base.Ui32(v103)>>(uint(int32(3))%32)))
	v220 = v108
	v221 = v104
	goto L19
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v104)+20))
	if v131 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v128
	v179 = v112
	goto L23
L33:
	;
	__phi147 = v141
	__phi148 = v142
	v147 = __phi147
	v148 = __phi148
	goto L37
L34:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	if v136 == int32(0) {
		goto L24
	} else {
		goto L36
	}
L35:
	;
	v141 = v131
	v142 = v104 + int32(20)
	goto L33
L36:
	;
	v141 = v136
	v142 = v104 + int32(16)
	goto L33
L37:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v147)+20))
	if v154 != 0 {
		__phi147 = v154
		__phi148 = v147 + int32(20)
		v147 = __phi147
		v148 = __phi148
		goto L37
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = int32(0)
	v179 = v147
	goto L23
L39:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v147)+16))
	if v157 != 0 {
		__phi147 = v157
		__phi148 = v147 + int32(16)
		v147 = __phi147
		v148 = __phi148
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v160 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v108 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v108
	goto L16
L42:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v104)+28))
	v190 = v188 << (uint(int32(2)) % 32)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[519])))
	if v104 != v193 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+24)) = v126
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	if v210 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L44:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	if v203 != v104 {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[519]))) = v179
	if v179 != 0 {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v196 = int32(0)
	v198 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v198 & base.I32_rotl(int32(-2), v188)
	v220 = v108
	v221 = v104
	goto L19
L47:
	;
	if v179 == int32(0) {
		v220 = v108
		v221 = v104
		goto L19
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+20)) = v179
	goto L47
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+16)) = v179
	goto L47
L50:
	;
	goto L43
L51:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v104)+20))
	if v215 == int32(0) {
		v220 = v108
		v221 = v104
		goto L19
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+16)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v210)+24)) = v179
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+20)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v215)+24)) = v179
	v220 = v108
	v221 = v104
	goto L19
L54:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v230&int32(1) == int32(0) {
		goto L17
	} else {
		goto L55
	}
L55:
	;
	if v230&int32(2) != 0 {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	if base.Ui32(int32(255)) < base.Ui32(v396) {
		goto L94
	} else {
		goto L95
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v221)+4)) = v276 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v221+v276))) = v276
	if v221 != v260 {
		v396 = v276
		goto L56
	} else {
		goto L93
	}
L58:
	;
	if v293 == int32(0) {
		goto L57
	} else {
		goto L81
	}
L59:
	;
	v338 = int32(0)
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v230 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v221)+4)) = v220 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v221+v220))) = v220
	v396 = v220
	goto L56
L61:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	if v96 != v238 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v96 != v260 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v240 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v221
	v244 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v245 = v244 + v220
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v221)+4)) = v245 | int32(1)
	v251 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v221 != v251 {
		goto L17
	} else {
		goto L64
	}
L64:
	;
	v253 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v253
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v253
	goto L16
L65:
	;
	v276 = v230&int32(-8) + v220
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	if base.Ui32(int32(255)) < base.Ui32(v230) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v262 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v221
	v266 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	v267 = v266 + v220
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v221)+4)) = v267 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v221+v267))) = v267
	goto L16
L67:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v96)+24))
	if v277 == v96 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	if v277 != v280 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+12)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v277)+8)) = v280
	goto L57
L70:
	;
	v282 = int32(0)
	v284 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v284 & base.I32_rotl(int32(-2), int32(base.Ui32(v230)>>(uint(int32(3))%32)))
	goto L57
L71:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	if v298 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v277)+8)) = v295
	v338 = v277
	goto L58
L73:
	;
	__phi314 = v308
	__phi315 = v309
	v314 = __phi314
	v315 = __phi315
	goto L77
L74:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	if v303 == int32(0) {
		goto L59
	} else {
		goto L76
	}
L75:
	;
	v308 = v298
	v309 = v96 + int32(20)
	goto L73
L76:
	;
	v308 = v303
	v309 = v96 + int32(16)
	goto L73
L77:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v314)+20))
	if v321 != 0 {
		__phi314 = v321
		__phi315 = v314 + int32(20)
		v314 = __phi314
		v315 = __phi315
		goto L77
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315))) = int32(0)
	v338 = v314
	goto L58
L79:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v314)+16))
	if v324 != 0 {
		__phi314 = v324
		__phi315 = v314 + int32(16)
		v314 = __phi314
		v315 = __phi315
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
	v349 = v347 << (uint(int32(2)) % 32)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v349)+uint32(_consts[519])))
	if v96 != v352 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338)+24)) = v293
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	if v369 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L83:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v293)+16))
	if v362 != v96 {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v349)+uint32(_consts[519]))) = v338
	if v338 != 0 {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v355 = int32(0)
	v357 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v357 & base.I32_rotl(int32(-2), v347)
	goto L57
L86:
	;
	if v338 == int32(0) {
		goto L57
	} else {
		goto L89
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+20)) = v338
	goto L86
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+16)) = v338
	goto L86
L89:
	;
	goto L82
L90:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	if v374 == int32(0) {
		goto L57
	} else {
		goto L92
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338)+16)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v369)+24)) = v338
	goto L90
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338)+20)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v374)+24)) = v338
	goto L57
L93:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v276
	goto L16
L94:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v396) {
		v443 = int32(31)
		goto L99
	} else {
		goto L100
	}
L95:
	;
	v408 = v396 & int32(-8)
	v410 = v408 + int32(9128464)
	v412 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v416 = int32(1) << (uint(int32(base.Ui32(v396)>>(uint(int32(3))%32))) % 32)
	if v412&v416 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v408)+uint32(_consts[523]))) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v422)+12)) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v221)+12)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v221)+8)) = v422
	goto L16
L97:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v408)+uint32(_consts[523])))
	v422 = v421
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v412 | v416
	v422 = v410
	goto L96
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v221)+28)) = v443
	*(*int64)(unsafe.Add(mBase, uint32(v221)+16)) = int64(0)
	v448 = v443 << (uint(int32(2)) % 32)
	v452 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	v454 = int32(1) << (uint(v443) % 32)
	if v452&v454 != 0 {
		goto L104
	} else {
		goto L105
	}
L100:
	;
	v433 = base.I32_clz(int32(base.Ui32(v396) >> (uint(int32(8)) % 32)))
	v436 = int32(1)
	v443 = int32(base.Ui32(v396)>>(uint(int32(38)-v433)%32))&v436 - v433<<(uint(v436)%32) + int32(62)
	goto L99
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v221+v515))) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v221)+12)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v221+v513))) = v516
	v527 = int32(0)
	v529 = *(*int32)(unsafe.Add(mBase, _consts[524]))
	v530 = int32(-1)
	v531 = v529 + v530
	if v531 != 0 {
		goto L113
	} else {
		goto L114
	}
L102:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v477)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v507)+12)) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v477)+8)) = v221
	v513 = int32(24)
	v515 = int32(8)
	v516 = int32(0)
	v517 = v477
	v518 = v507
	goto L101
L103:
	;
	v513 = v498
	v515 = v500
	v516 = v221
	v517 = v221
	v518 = v503
	goto L101
L104:
	;
	if v443 == int32(31) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, _consts[520])) = v452 | v454
	*(*int32)(unsafe.Add(mBase, uint32(v448)+uint32(_consts[519]))) = v221
	v498 = int32(8)
	v500 = int32(24)
	v503 = v448 + int32(9128728)
	goto L103
L106:
	;
	v469 = int32(0)
	goto L108
L107:
	;
	v469 = int32(25) - int32(base.Ui32(v443)>>(uint(int32(1))%32))
	goto L108
L108:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v448)+uint32(_consts[519])))
	v474 = v396 << (uint(v469) % 32)
	v477 = v471
	goto L109
L109:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v477)+4))
	if v481&int32(-8) == v396 {
		goto L102
	} else {
		goto L111
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491+int32(16)))) = v221
	v498 = int32(8)
	v500 = int32(24)
	v503 = v477
	goto L103
L111:
	;
	v491 = v477 + int32(base.Ui32(v474)>>(uint(int32(29))%32))&int32(4)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)+16))
	if v492 != 0 {
		v474 = v474 << (uint(int32(1)) % 32)
		v477 = v492
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v533 = v531
	goto L115
L114:
	;
	v533 = v530
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, _consts[524])) = v533
	goto L17
L116:
	;
	return
L117:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_valkey_malloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	if base.Ui32(int32(2147483646)) < base.Ui32(l0) {
		v50 = *(*int32)(unsafe.Add(mBase, _consts[1001]))
		m.T0[v50].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		if l0 != 0 {
			v8 = l0
		} else {
			v8 = int32(4)
		}
		v10 = v8 + int32(8)
		v11 = F_emscripten_builtin_malloc(m, v10)
		mBase = m.M
		if v11 == int32(0) {
			v50 = *(*int32)(unsafe.Add(mBase, _consts[1001]))
			m.T0[v50].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v8
			v16 = *(*int32)(unsafe.Add(mBase, _consts[411]))
			if v16 != int32(-1) {
				v27 = v16
			} else {
				v19 = int32(0)
				v21 = *(*int32)(unsafe.Add(mBase, _consts[281]))
				*(*int32)(unsafe.Add(mBase, _consts[411])) = v21
				*(*int32)(unsafe.Add(mBase, _consts[281])) = v21 + int32(1)
				v27 = v21
			}
			if v27 < int32(260) {
				v36 = v27 << (uint(int32(2)) % 32)
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[285])))
				*(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[285]))) = v39 + v10
			} else {
				v30 = int32(0)
				v32 = *(*int32)(unsafe.Add(mBase, _consts[286]))
				*(*int32)(unsafe.Add(mBase, _consts[286])) = v32 + v10
			}
			return v11 + int32(8)
		}
	}
}
func F_valkey_malloc_cache_aligned(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	if l0 != 0 {
		v6 = l0
	} else {
		v6 = int32(4)
	}
	if base.Ui32(v6) < base.Ui32(int32(2147483647)) {
		v19 = F_emscripten_builtin_malloc(m, v6+int32(75))
		mBase = m.M
		if v19 != 0 {
			v27 = (v19 + int32(75)) & int32(-64)
			*(*int32)(unsafe.Add(mBase, uint32(v27+int32(-8)))) = v6 | int32(-2147483648)
			*(*int32)(unsafe.Add(mBase, uint32(v27+int32(-12)))) = v19
			v37 = *(*int32)(unsafe.Add(mBase, _consts[411]))
			if v37 != int32(-1) {
				v48 = v37
			} else {
				v40 = int32(0)
				v42 = *(*int32)(unsafe.Add(mBase, _consts[281]))
				*(*int32)(unsafe.Add(mBase, _consts[411])) = v42
				*(*int32)(unsafe.Add(mBase, _consts[281])) = v42 + int32(1)
				v48 = v42
			}
			v50 = v6 + int32(8)
			if v48 < int32(260) {
				v60 = v48 << (uint(int32(2)) % 32)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+uint32(_consts[285])))
				*(*int32)(unsafe.Add(mBase, uint32(v60)+uint32(_consts[285]))) = v63 + v50
				return v27
			} else {
				v53 = int32(0)
				v55 = *(*int32)(unsafe.Add(mBase, _consts[286]))
				*(*int32)(unsafe.Add(mBase, _consts[286])) = v55 + v50
				return v27
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[1001]))
			m.T0[v21].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v27 = (v19 + int32(75)) & int32(-64)
				*(*int32)(unsafe.Add(mBase, uint32(v27+int32(-8)))) = v6 | int32(-2147483648)
				*(*int32)(unsafe.Add(mBase, uint32(v27+int32(-12)))) = v19
				v37 = *(*int32)(unsafe.Add(mBase, _consts[411]))
				if v37 != int32(-1) {
					v48 = v37
				} else {
					v40 = int32(0)
					v42 = *(*int32)(unsafe.Add(mBase, _consts[281]))
					*(*int32)(unsafe.Add(mBase, _consts[411])) = v42
					*(*int32)(unsafe.Add(mBase, _consts[281])) = v42 + int32(1)
					v48 = v42
				}
				v50 = v6 + int32(8)
				if v48 < int32(260) {
					v60 = v48 << (uint(int32(2)) % 32)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+uint32(_consts[285])))
					*(*int32)(unsafe.Add(mBase, uint32(v60)+uint32(_consts[285]))) = v63 + v50
					return v27
				} else {
					v53 = int32(0)
					v55 = *(*int32)(unsafe.Add(mBase, _consts[286]))
					*(*int32)(unsafe.Add(mBase, _consts[286])) = v55 + v50
					return v27
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[1001]))
		m.T0[v10].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_valkey_strtod_n(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int64
	_ = v17
	var v22 int64
	_ = v22
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 float64
	_ = v43
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(9116376)
	v12 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[5])) = v12
	v17 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(8)))) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(0)
	v22 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v22
	F_ffc_from_chars_double_options(m, v9+int32(16), l0, l0+l1, v9+int32(24), v9)
	mBase = m.M
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if v30 == v12 {
	} else {
		if v30 == int32(2) {
			v37 = int32(68)
		} else {
			v37 = int32(28)
		}
		*(*int32)(unsafe.Add(mBase, _consts[5])) = v37
	}
	if l2 == int32(0) {
	} else {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v41
	}
	v43 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
	m.G0 = v9 + int32(32)
	return v43
}
