package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_addRdbReplicaToPsyncWait(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int64
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	if v12 != 0 {
		v34 = int32(1)
		v36 = *(*int32)(unsafe.Add(mBase, _consts[314]))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
		if v37 != 0 {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
			if v39 == int32(0) {
				v47 = v37
				v48 = v34
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
				*(*int32)(unsafe.Add(mBase, uint32(v39))) = v42 + int32(1)
				v47 = v37
				v48 = int32(0)
			}
		} else {
			v47 = int32(0)
			v48 = v34
		}
		v51 = *(*int32)(unsafe.Add(mBase, _consts[15]))
		if int32(0) < v51 {
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v48 != 0 {
				v73 = int32(0)
			} else {
				v73 = v47
			}
			*(*int32)(unsafe.Add(mBase, uint32(v71)+184)) = v73
			v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v75 | int32(33554432)
			v79 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			v80 = int64(56)
			v82 = int64(65280)
			v84 = int64(40)
			v87 = int64(16711680)
			v89 = int64(24)
			v91 = int64(4278190080)
			v93 = int64(8)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v79<<(uint(v80)%64) | v79&v82<<(uint(v84)%64) | (v79&v87<<(uint(v89)%64) | v79&v91<<(uint(v93)%64)) | (int64(base.Ui64(v79)>>(uint(v93)%64))&v91 | int64(base.Ui64(v79)>>(uint(v89)%64))&v87 | (int64(base.Ui64(v79)>>(uint(v84)%64))&v82 | int64(base.Ui64(v79)>>(uint(v80)%64))))
			v117 = *(*int32)(unsafe.Add(mBase, _consts[505]))
			v122 = F_raxInsert(m, v117, v9+int32(24), int32(8), l0, int32(0))
			mBase = m.M
			v123 = m.ExcPending
			if v123 != 0 {
				return
			} else {
				m.G0 = v9 + int32(32)
				return
			}
		} else {
			v54 = F_replicationGetReplicaName(m, l0)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				if v48 != 0 {
					v61 = int32(_a908)
				} else {
					v61 = int32(_a909)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v61
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v54
				F__serverLog(m, int32(0), int32(_a910), v9)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					if v48 != 0 {
						v73 = int32(0)
					} else {
						v73 = v47
					}
					*(*int32)(unsafe.Add(mBase, uint32(v71)+184)) = v73
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v75 | int32(33554432)
					v79 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					v80 = int64(56)
					v82 = int64(65280)
					v84 = int64(40)
					v87 = int64(16711680)
					v89 = int64(24)
					v91 = int64(4278190080)
					v93 = int64(8)
					*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v79<<(uint(v80)%64) | v79&v82<<(uint(v84)%64) | (v79&v87<<(uint(v89)%64) | v79&v91<<(uint(v93)%64)) | (int64(base.Ui64(v79)>>(uint(v93)%64))&v91 | int64(base.Ui64(v79)>>(uint(v89)%64))&v87 | (int64(base.Ui64(v79)>>(uint(v84)%64))&v82 | int64(base.Ui64(v79)>>(uint(v80)%64))))
					v117 = *(*int32)(unsafe.Add(mBase, _consts[505]))
					v122 = F_raxInsert(m, v117, v9+int32(24), int32(8), l0, int32(0))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				}
			}
		}
	} else {
		v15 = F_valkey_malloc(m, int32(32))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[448])) = v15
			*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(0)
			v20 = F_raxNew(m)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = int32(_a69)
				v23 = *(*int32)(unsafe.Add(mBase, _consts[448]))
				*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v20
				v28 = *(*int64)(unsafe.Add(mBase, _consts[31]))
				*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v28 + int64(1)
				v47 = int32(0)
				v48 = int32(1)
				v51 = *(*int32)(unsafe.Add(mBase, _consts[15]))
				if int32(0) < v51 {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					if v48 != 0 {
						v73 = int32(0)
					} else {
						v73 = v47
					}
					*(*int32)(unsafe.Add(mBase, uint32(v71)+184)) = v73
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v75 | int32(33554432)
					v79 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					v80 = int64(56)
					v82 = int64(65280)
					v84 = int64(40)
					v87 = int64(16711680)
					v89 = int64(24)
					v91 = int64(4278190080)
					v93 = int64(8)
					*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v79<<(uint(v80)%64) | v79&v82<<(uint(v84)%64) | (v79&v87<<(uint(v89)%64) | v79&v91<<(uint(v93)%64)) | (int64(base.Ui64(v79)>>(uint(v93)%64))&v91 | int64(base.Ui64(v79)>>(uint(v89)%64))&v87 | (int64(base.Ui64(v79)>>(uint(v84)%64))&v82 | int64(base.Ui64(v79)>>(uint(v80)%64))))
					v117 = *(*int32)(unsafe.Add(mBase, _consts[505]))
					v122 = F_raxInsert(m, v117, v9+int32(24), int32(8), l0, int32(0))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				} else {
					v54 = F_replicationGetReplicaName(m, l0)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
						if v48 != 0 {
							v61 = int32(_a908)
						} else {
							v61 = int32(_a909)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v61
						*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v56
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v54
						F__serverLog(m, int32(0), int32(_a910), v9)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							if v48 != 0 {
								v73 = int32(0)
							} else {
								v73 = v47
							}
							*(*int32)(unsafe.Add(mBase, uint32(v71)+184)) = v73
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v75 | int32(33554432)
							v79 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							v80 = int64(56)
							v82 = int64(65280)
							v84 = int64(40)
							v87 = int64(16711680)
							v89 = int64(24)
							v91 = int64(4278190080)
							v93 = int64(8)
							*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v79<<(uint(v80)%64) | v79&v82<<(uint(v84)%64) | (v79&v87<<(uint(v89)%64) | v79&v91<<(uint(v93)%64)) | (int64(base.Ui64(v79)>>(uint(v93)%64))&v91 | int64(base.Ui64(v79)>>(uint(v89)%64))&v87 | (int64(base.Ui64(v79)>>(uint(v84)%64))&v82 | int64(base.Ui64(v79)>>(uint(v80)%64))))
							v117 = *(*int32)(unsafe.Add(mBase, _consts[505]))
							v122 = F_raxInsert(m, v117, v9+int32(24), int32(8), l0, int32(0))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_parseCheckRdbOptions(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if int32(1) < l0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v333
	v336 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	v338 = F_fiprintf(m, v336, int32(_a1657), v9)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L44
	} else {
		goto L89
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1117])) = int32(1)
	if l0 < int32(3) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if l2 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L2
L5:
	;
	v315 = F_getVersion(m)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L44
	} else {
		goto L86
	}
L6:
	;
	m.G0 = v9 + int32(32)
	return
L7:
	;
	v25 = int32(2)
	goto L8
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v30 != int32(45) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L6
L10:
	;
	v39 = int32(_a1658)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1118])))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v43 == int32(0) {
		v66 = v42
		v67 = v43
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	if v33 != int32(118) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
	if v36 == int32(0) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	if v67-v66&int32(255) == int32(0) {
		goto L5
	} else {
		goto L22
	}
L15:
	;
	goto L14
L16:
	;
	if v43 != v42&int32(255) {
		v66 = v42
		v67 = v43
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v49 = v29
	v50 = v39
	goto L18
L18:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if v54 == int32(0) {
		v66 = v53
		v67 = v54
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v66 = v53
	v67 = v54
	goto L15
L20:
	;
	v57 = int32(1)
	if v54 == v53&int32(255) {
		v49 = v49 + v57
		v50 = v50 + v57
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1+v25<<(uint(int32(2))%32))))
	v77 = int32(_a1659)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1119])))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v81 == int32(0) {
		v104 = v80
		v105 = v81
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v304 = v301 + int32(1)
	if v304 < l0 {
		v25 = v304
		goto L8
	} else {
		goto L85
	}
L24:
	;
	v112 = int32(_a1660)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1120])))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v116 == int32(0) {
		v139 = v115
		v140 = v116
		goto L36
	} else {
		goto L37
	}
L25:
	;
	if v105-v104&int32(255) != 0 {
		goto L24
	} else {
		goto L33
	}
L26:
	;
	goto L25
L27:
	;
	if v81 != v80&int32(255) {
		v104 = v80
		v105 = v81
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v87 = v76
	v88 = v77
	goto L29
L29:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	if v92 == int32(0) {
		v104 = v91
		v105 = v92
		goto L26
	} else {
		goto L31
	}
L30:
	;
	v104 = v91
	v105 = v92
	goto L26
L31:
	;
	v95 = int32(1)
	if v92 == v91&int32(255) {
		v87 = v87 + v95
		v88 = v88 + v95
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1121])) = int32(1)
	v301 = v25
	goto L23
L34:
	;
	v157 = int32(_a1661)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1122])))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v161 == int32(0) {
		v184 = v160
		v185 = v161
		goto L47
	} else {
		goto L48
	}
L35:
	;
	if v140-v139&int32(255) != 0 {
		goto L34
	} else {
		goto L43
	}
L36:
	;
	goto L35
L37:
	;
	if v116 != v115&int32(255) {
		v139 = v115
		v140 = v116
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v122 = v76
	v123 = v112
	goto L39
L39:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
	if v127 == int32(0) {
		v139 = v126
		v140 = v127
		goto L36
	} else {
		goto L41
	}
L40:
	;
	v139 = v126
	v140 = v127
	goto L36
L41:
	;
	v130 = int32(1)
	if v127 == v126&int32(255) {
		v122 = v122 + v130
		v123 = v123 + v130
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v145 = v25 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1+v145<<(uint(int32(2))%32))))
	v150 = F_zstrdup(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return
L45:
	;
	v152 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1123])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1124])) = v150
	v301 = v145
	goto L23
L46:
	;
	if v185-v184&int32(255) != 0 {
		goto L1
	} else {
		goto L54
	}
L47:
	;
	goto L46
L48:
	;
	if v161 != v160&int32(255) {
		v184 = v160
		v185 = v161
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v167 = v76
	v168 = v157
	goto L50
L50:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	if v172 == int32(0) {
		v184 = v171
		v185 = v172
		goto L47
	} else {
		goto L52
	}
L51:
	;
	v184 = v171
	v185 = v172
	goto L47
L52:
	;
	v175 = int32(1)
	if v172 == v171&int32(255) {
		v167 = v167 + v175
		v168 = v168 + v175
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	if v25 == l0+int32(-1) {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v191 = v25 + int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1+v191<<(uint(int32(2))%32))))
	v196 = int32(_a1662)
	v199 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1125])))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v200 == int32(0) {
		v223 = v199
		v224 = v200
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v231 = int32(_a1663)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1126])))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v235 == int32(0) {
		v258 = v234
		v259 = v235
		goto L68
	} else {
		goto L69
	}
L57:
	;
	if v224-v223&int32(255) != 0 {
		goto L56
	} else {
		goto L65
	}
L58:
	;
	goto L57
L59:
	;
	if v200 != v199&int32(255) {
		v223 = v199
		v224 = v200
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v206 = v195
	v207 = v196
	goto L61
L61:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if v211 == int32(0) {
		v223 = v210
		v224 = v211
		goto L58
	} else {
		goto L63
	}
L62:
	;
	v223 = v210
	v224 = v211
	goto L58
L63:
	;
	v214 = int32(1)
	if v211 == v210&int32(255) {
		v206 = v206 + v214
		v207 = v207 + v214
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1117])) = int32(1)
	v301 = v191
	goto L23
L66:
	;
	v266 = int32(_a1664)
	v269 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1127])))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v270 == int32(0) {
		v293 = v269
		v294 = v270
		goto L77
	} else {
		goto L78
	}
L67:
	;
	if v259-v258&int32(255) != 0 {
		goto L66
	} else {
		goto L75
	}
L68:
	;
	goto L67
L69:
	;
	if v235 != v234&int32(255) {
		v258 = v234
		v259 = v235
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v241 = v195
	v242 = v231
	goto L71
L71:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+1)))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+1)))
	if v246 == int32(0) {
		v258 = v245
		v259 = v246
		goto L68
	} else {
		goto L73
	}
L72:
	;
	v258 = v245
	v259 = v246
	goto L68
L73:
	;
	v249 = int32(1)
	if v246 == v245&int32(255) {
		v241 = v241 + v249
		v242 = v242 + v249
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1117])) = int32(2)
	v301 = v191
	goto L23
L76:
	;
	if v294-v293&int32(255) != 0 {
		goto L1
	} else {
		goto L84
	}
L77:
	;
	goto L76
L78:
	;
	if v270 != v269&int32(255) {
		v293 = v269
		v294 = v270
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v276 = v195
	v277 = v266
	goto L80
L80:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+1)))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+1)))
	if v281 == int32(0) {
		v293 = v280
		v294 = v281
		goto L77
	} else {
		goto L82
	}
L81:
	;
	v293 = v280
	v294 = v281
	goto L77
L82:
	;
	v284 = int32(1)
	if v281 == v280&int32(255) {
		v276 = v276 + v284
		v277 = v277 + v284
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v298 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1117])) = v298
	v301 = v191
	goto L23
L85:
	;
	goto L9
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v315
	v321 = F_iprintf(m, int32(_a1665), v9+int32(16))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L44
	} else {
		goto L87
	}
L87:
	;
	F_sdsfree(m, v315)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L44
	} else {
		goto L88
	}
L88:
	;
	m.Env.Exit(m, int32(0))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_rdbGetObjectType(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v22 int32
	_ = v22
	var v38 int32
	_ = v38
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = v5 & int32(15)
	switch v7 {
	case 0:
		v93 = v7
		return v93
	case 1:
		switch int32(base.Ui32(v5)>>(uint(int32(4))%32))&int32(15) + int32(-9) {
		case 0, 2:
			v93 = int32(18)
			return v93
		default:
			F__serverPanic_1(m, int32(_a875), int32(720), int32(_a852), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	case 2:
		switch int32(base.Ui32(v5)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
		case 0:
			v93 = int32(2)
			return v93
		default:
			F__serverPanic_1(m, int32(_a875), int32(729), int32(_a853), int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 4:
			v93 = int32(11)
			return v93
		case 9:
			return int32(20)
		}
	case 3:
		switch int32(base.Ui32(v5)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
		case 0:
			return int32(5)
		default:
			F__serverPanic_1(m, int32(_a875), int32(736), int32(_a854), int32(0))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 4:
			v93 = int32(17)
			return v93
		}
	case 4:
		switch int32(base.Ui32(v5)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
		case 0:
			v64 = F_hashTypeHasVolatileFields(m, l0)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				if v64 == int32(0) {
					v93 = int32(4)
					return v93
				} else {
					if int32(79) < l1 {
						v72 = int32(22)
					} else {
						v72 = int32(-1)
					}
					return v72
				}
			}
		default:
			F__serverPanic_1(m, int32(_a875), int32(749), int32(_a855), int32(0))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 9:
			v93 = int32(16)
			return v93
		}
	case 5:
		return int32(7)
	case 6:
		return int32(21)
	default:
		F__serverPanic_1(m, int32(_a875), int32(752), int32(_a123), int32(0))
		mBase = m.M
		v90 = m.ExcPending
		if v90 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_rdbIsVersionAccepted(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	v4 = int32(0)
	if l0 < int32(1) {
		v23 = v4
	} else {
		if l1&base.B2i32(l0 < int32(80)) != 0 {
			v23 = v4
		} else {
			if base.B2i32(int32(79) < l0)&l2 != 0 {
				v23 = v4
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _consts[474]))
				if v15 != 0 {
					v23 = int32(1)
				} else {
					if base.Ui32(int32(80)) < base.Ui32(l0) {
						v23 = v4
					} else {
						if base.Ui32(l0+int32(-12)) < base.Ui32(int32(68)) {
							v23 = v4
						} else {
							v23 = int32(1)
						}
					}
				}
			}
		}
	}
	return v23
}
func F_rdbLoadBinaryFloatValue(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v7&int32(5) != 0 {
		v44 = int32(-1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v44
L2:
	;
	v12 = l1
	v13 = int32(4)
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v16) < base.Ui32(v13) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v44 = int32(0)
	goto L1
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v31 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v18 = v16
	goto L8
L7:
	;
	v18 = v13
	goto L8
L8:
	;
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v19 = v18
	goto L11
L10:
	;
	v19 = v13
	goto L11
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = m.T0[v20].(func(*base.Module, int32, int32, int32) int32)(m, l0, v12, v19)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v21 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v25 | int64(1)
	return int32(-1)
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v36 + v19
	v40 = v13 - v19
	if v40 != 0 {
		v12 = v12 + v19
		v13 = v40
		goto L3
	} else {
		goto L18
	}
L16:
	;
	m.T0[v31].(func(*base.Module, int32, int32, int32))(m, l0, v12, v19)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L4
}
func F_rdbLoadCheckModuleValue(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v31 int64
	_ = v31
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int64
	_ = v185
	var v188 int64
	_ = v188
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v16 = F_rdbLoadLenByRef(m, l0, int32(0), v11+int32(56))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v11)+56))
	if v16 == int32(-1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v201 = F_createStringObject_1(m, int32(_a883), int32(18))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L66
	}
L4:
	;
	v24 = int64(-1)
	goto L6
L5:
	;
	v24 = v21
	goto L6
L6:
	;
	if v24 == int64(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v31 = v24
	goto L8
L8:
	;
	if base.Ui64(int64(2)) < base.Ui64(v31) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L3
L10:
	;
	v182 = F_rdbLoadLenByRef(m, l0, int32(0), v11+int32(56))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L61
	}
L11:
	;
	v51 = v31 + int64(-3)
	if base.Ui64(int64(2)) < base.Ui64(v51) {
		goto L10
	} else {
		goto L16
	}
L12:
	;
	v40 = F_rdbLoadLenByRef(m, l0, int32(0), v11+int32(56))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v40 != int32(-1) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
	F_rdbReportError(m, int32(1), int32(1742), int32(_a884), v11)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	switch base.I32_wrap_i64(v51) {
	default:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L20
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l1
	F_rdbReportError(m, int32(1), int32(1753), int32(_a885), v11+int32(32))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L60
	}
L18:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v104&int32(5) != 0 {
		goto L42
	} else {
		goto L43
	}
L19:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v69&int32(5) != 0 {
		goto L17
	} else {
		goto L26
	}
L20:
	;
	v55 = int32(0)
	v57 = F_rdbGenericLoadStringObject(m, l0, v55, v55)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	F_decrRefCount(m, v57)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L25
	}
L22:
	;
	if v57 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	F_rdbReportError(m, int32(1), int32(1747), int32(_a886), v11+int32(16))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	goto L10
L26:
	;
	v78 = int32(4)
	v80 = v11 + int32(56)
	goto L27
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v83) < base.Ui32(v78) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v94 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L30:
	;
	v85 = v83
	goto L32
L31:
	;
	v85 = v78
	goto L32
L32:
	;
	if v83 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v86 = v85
	goto L35
L34:
	;
	v86 = v78
	goto L35
L35:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = m.T0[v87].(func(*base.Module, int32, int32, int32) int32)(m, l0, v80, v86)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v88 != 0 {
		goto L29
	} else {
		goto L37
	}
L37:
	;
	v90 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v90 | int64(1)
	goto L17
L38:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v99 + v86
	v103 = v78 - v86
	if v103 != 0 {
		v78 = v103
		v80 = v80 + v86
		goto L27
	} else {
		goto L41
	}
L39:
	;
	m.T0[v94].(func(*base.Module, int32, int32, int32))(m, l0, v80, v86)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	goto L10
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = l1
	F_rdbReportError(m, int32(1), int32(1758), int32(_a887), v11+int32(48))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L59
	}
L43:
	;
	v113 = int32(8)
	v115 = v11 + int32(56)
	goto L44
L44:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v118) < base.Ui32(v113) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v129 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L47:
	;
	v120 = v118
	goto L49
L48:
	;
	v120 = v113
	goto L49
L49:
	;
	if v118 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v121 = v120
	goto L52
L51:
	;
	v121 = v113
	goto L52
L52:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v123 = m.T0[v122].(func(*base.Module, int32, int32, int32) int32)(m, l0, v115, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v123 != 0 {
		goto L46
	} else {
		goto L54
	}
L54:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v125 | int64(1)
	goto L42
L55:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v134 + v121
	v138 = v113 - v121
	if v138 != 0 {
		v113 = v138
		v115 = v115 + v121
		goto L44
	} else {
		goto L58
	}
L56:
	;
	m.T0[v129].(func(*base.Module, int32, int32, int32))(m, l0, v115, v121)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	goto L10
L59:
	;
	goto L10
L60:
	;
	goto L10
L61:
	;
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v11)+56))
	if v182 == int32(-1) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v188 = int64(-1)
	goto L64
L63:
	;
	v188 = v185
	goto L64
L64:
	;
	if v188 != int64(0) {
		v31 = v188
		goto L8
	} else {
		goto L65
	}
L65:
	;
	goto L9
L66:
	;
	m.G0 = v11 + int32(64)
	return v201
}
func F_rdbLoadDoubleValue(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 float64
	_ = v56
	var v58 int32
	_ = v58
	var v60 float64
	_ = v60
	var v62 int32
	_ = v62
	var v64 float64
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
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
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	v9 = m.G0
	v11 = v9 - int32(272)
	m.G0 = v11
	v13 = int32(-1)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v14&int32(5) != 0 {
		v131 = v13
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(272)
	return v131
L2:
	;
	v24 = v11 + int32(15)
	v25 = int32(1)
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v28) < base.Ui32(v25) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	switch v51 + int32(-253) {
	case 0:
		goto L20
	case 1:
		goto L21
	case 2:
		goto L22
	default:
		goto L19
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v41 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v30 = v28
	goto L8
L7:
	;
	v30 = v25
	goto L8
L8:
	;
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = v30
	goto L11
L10:
	;
	v31 = v25
	goto L11
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = m.T0[v32].(func(*base.Module, int32, int32, int32) int32)(m, l0, v24, v31)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v33 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v37 | int64(1)
	v131 = v13
	goto L1
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v46 + v31
	v50 = v25 - v31
	if v50 != 0 {
		v24 = v24 + v31
		v25 = v50
		goto L3
	} else {
		goto L18
	}
L16:
	;
	m.T0[v41].(func(*base.Module, int32, int32, int32))(m, l0, v24, v31)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L4
L19:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v66&int32(5) != 0 {
		v131 = v13
		goto L1
	} else {
		goto L23
	}
L20:
	;
	v62 = int32(0)
	v64 = *(*float64)(unsafe.Add(mBase, _consts[478]))
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v64
	v131 = v62
	goto L1
L21:
	;
	v58 = int32(0)
	v60 = *(*float64)(unsafe.Add(mBase, _consts[479]))
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v60
	v131 = v58
	goto L1
L22:
	;
	v54 = int32(0)
	v56 = *(*float64)(unsafe.Add(mBase, _consts[480]))
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v56
	v131 = v54
	goto L1
L23:
	;
	v69 = int32(0)
	if v51 == v69 {
		v110 = v69
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v113 = v11 + int32(16)
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v113+v110))) = uint8(v115)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
	v123 = F_sscanf(m, v113, int32(_a874), v11)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L12
	} else {
		goto L41
	}
L25:
	;
	v78 = v11 + int32(16)
	v79 = v51
	goto L26
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v82) < base.Ui32(v79) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	v110 = v103
	goto L24
L28:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v93 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L29:
	;
	v84 = v82
	goto L31
L30:
	;
	v84 = v79
	goto L31
L31:
	;
	if v82 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v85 = v84
	goto L34
L33:
	;
	v85 = v79
	goto L34
L34:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = m.T0[v86].(func(*base.Module, int32, int32, int32) int32)(m, l0, v78, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	if v87 != 0 {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v89 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v89 | int64(1)
	v131 = v13
	goto L1
L37:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v98 + v85
	v102 = v79 - v85
	if v102 != 0 {
		v78 = v78 + v85
		v79 = v102
		goto L26
	} else {
		goto L40
	}
L38:
	;
	m.T0[v93].(func(*base.Module, int32, int32, int32))(m, l0, v78, v85)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	goto L27
L41:
	;
	if v123 != int32(1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v127 = int32(-1)
	goto L44
L43:
	;
	v127 = v115
	goto L44
L44:
	;
	v131 = v127
	goto L1
}
func F_rdbLoadLen(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v22 int64
	_ = v22
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = F_rdbLoadLenByRef(m, l0, l1, v7+int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		m.G0 = v7 + int32(16)
		if v11 == int32(-1) {
			v22 = int64(-1)
		} else {
			v22 = v15
		}
		return v22
	}
}
func F_rdbLoadMillisecondTime(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v55 int64
	_ = v55
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int64(9223372036854775807)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v14&int32(5) != 0 {
		v55 = v13
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v55
L2:
	;
	v17 = int32(8)
	v24 = v11 + v17
	v25 = v17
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v28) < base.Ui32(v25) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v55 = v51
	goto L1
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v41 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v30 = v28
	goto L8
L7:
	;
	v30 = v25
	goto L8
L8:
	;
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = v30
	goto L11
L10:
	;
	v31 = v25
	goto L11
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = m.T0[v32].(func(*base.Module, int32, int32, int32) int32)(m, l0, v24, v31)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int64(0)
L13:
	;
	if v33 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v37 | int64(1)
	v55 = v13
	goto L1
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v46 + v31
	v50 = v25 - v31
	if v50 != 0 {
		v24 = v24 + v31
		v25 = v50
		goto L3
	} else {
		goto L18
	}
L16:
	;
	m.T0[v41].(func(*base.Module, int32, int32, int32))(m, l0, v24, v31)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L4
}
func F_rdbLoadTime(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int64(-1)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v13&int32(5) != 0 {
		v52 = v12
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v52
L2:
	;
	v22 = v10 + int32(12)
	v23 = int32(4)
	goto L3
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v26) < base.Ui32(v23) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10)+12)))
	v52 = v49
	goto L1
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v39 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v28 = v26
	goto L8
L7:
	;
	v28 = v23
	goto L8
L8:
	;
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v29 = v28
	goto L11
L10:
	;
	v29 = v23
	goto L11
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = m.T0[v30].(func(*base.Module, int32, int32, int32) int32)(m, l0, v22, v29)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int64(0)
L13:
	;
	if v31 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v35 | int64(1)
	v52 = v12
	goto L1
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v44 + v29
	v48 = v23 - v29
	if v48 != 0 {
		v22 = v22 + v29
		v23 = v48
		goto L3
	} else {
		goto L18
	}
L16:
	;
	m.T0[v39].(func(*base.Module, int32, int32, int32))(m, l0, v22, v29)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L4
}
func F_rdbReportError(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int64
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	v6 = m.G0
	v8 = v6 - int32(1168)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+104)) = l1
	v12 = *(*int64)(unsafe.Add(mBase, _consts[397]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+96)) = v12
	v20 = F_snprintf(m, v8+int32(128), int32(1024), int32(_a869), v8+int32(96))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+1164)) = l3
		v28 = F_vsnprintf(m, v20+(v8+int32(128)), int32(1024)-v20, l2, l3)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, _consts[88]))
			if v31 == int32(0) {
				v50 = int32(0)
				v51 = *(*int32)(unsafe.Add(mBase, _consts[475]))
				if v51 == v50 {
					v62 = int32(0)
					v63 = *(*int32)(unsafe.Add(mBase, _consts[476]))
					if v63 == v62 {
						v110 = *(*int32)(unsafe.Add(mBase, _consts[15]))
						if l0 == int32(0) {
							if int32(3) < v110 {
								m.G0 = v8 + int32(1168)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(128)
								F__serverLog(m, int32(3), int32(_a870), v8)
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return
								} else {
									m.G0 = v8 + int32(1168)
									return
								}
							}
						} else {
							if int32(3) < v110 {
								m.Env.Exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v8 + int32(128)
								F__serverLog(m, int32(3), int32(_a871), v8+int32(16))
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return
								} else {
									v145 = *(*int32)(unsafe.Add(mBase, _consts[15]))
									if int32(3) < v145 {
										m.Env.Exit(m, int32(1))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									} else {
										F__serverLog(m, int32(3), int32(_a872), int32(0))
										mBase = m.M
										v152 = m.ExcPending
										if v152 != 0 {
											return
										} else {
											m.Env.Exit(m, int32(1))
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, _consts[15]))
						if int32(3) < v67 {
							v81 = v63
							*(*int32)(unsafe.Add(mBase, uint32(v8)+124)) = v81
							*(*int32)(unsafe.Add(mBase, uint32(v8)+120)) = int32(_a188)
							v87 = m.G0
							v88 = int32(96)
							v89 = v87 - v88
							m.G0 = v89
							v91 = F_stat(m, v81, v89)
							mBase = m.M
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
							m.G0 = v89 + v88
							if base.B2i32(v91 != int32(-1))&base.B2i32(v92&int32(61440) == int32(4096)) != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v81
								F_rdbCheckError(m, int32(_a873), v8+int32(32))
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return
								} else {
									m.G0 = v8 + int32(1168)
									return
								}
							} else {
								v107 = F_redis_check_rdb_main(m, int32(2), v8+int32(120), int32(0))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									v145 = *(*int32)(unsafe.Add(mBase, _consts[15]))
									if int32(3) < v145 {
										m.Env.Exit(m, int32(1))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									} else {
										F__serverLog(m, int32(3), int32(_a872), int32(0))
										mBase = m.M
										v152 = m.ExcPending
										if v152 != 0 {
											return
										} else {
											m.Env.Exit(m, int32(1))
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v8 + int32(128)
							F__serverLog(m, int32(3), int32(_a57), v8+int32(48))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, _consts[476]))
								v81 = v80
								*(*int32)(unsafe.Add(mBase, uint32(v8)+124)) = v81
								*(*int32)(unsafe.Add(mBase, uint32(v8)+120)) = int32(_a188)
								v87 = m.G0
								v88 = int32(96)
								v89 = v87 - v88
								m.G0 = v89
								v91 = F_stat(m, v81, v89)
								mBase = m.M
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
								m.G0 = v89 + v88
								if base.B2i32(v91 != int32(-1))&base.B2i32(v92&int32(61440) == int32(4096)) != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v81
									F_rdbCheckError(m, int32(_a873), v8+int32(32))
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return
									} else {
										m.G0 = v8 + int32(1168)
										return
									}
								} else {
									v107 = F_redis_check_rdb_main(m, int32(2), v8+int32(120), int32(0))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										v145 = *(*int32)(unsafe.Add(mBase, _consts[15]))
										if int32(3) < v145 {
											m.Env.Exit(m, int32(1))
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										} else {
											F__serverLog(m, int32(3), int32(_a872), int32(0))
											mBase = m.M
											v152 = m.ExcPending
											if v152 != 0 {
												return
											} else {
												m.Env.Exit(m, int32(1))
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
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v8 + int32(128)
					F_rdbCheckError(m, int32(_a57), v8+int32(64))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						v145 = *(*int32)(unsafe.Add(mBase, _consts[15]))
						if int32(3) < v145 {
							m.Env.Exit(m, int32(1))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						} else {
							F__serverLog(m, int32(3), int32(_a872), int32(0))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return
							} else {
								m.Env.Exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v34 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
				if v34 == int64(-1) {
					v50 = int32(0)
					v51 = *(*int32)(unsafe.Add(mBase, _consts[475]))
					if v51 == v50 {
						v62 = int32(0)
						v63 = *(*int32)(unsafe.Add(mBase, _consts[476]))
						if v63 == v62 {
							v110 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							if l0 == int32(0) {
								if int32(3) < v110 {
									m.G0 = v8 + int32(1168)
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(128)
									F__serverLog(m, int32(3), int32(_a870), v8)
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return
									} else {
										m.G0 = v8 + int32(1168)
										return
									}
								}
							} else {
								if int32(3) < v110 {
									m.Env.Exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v8 + int32(128)
									F__serverLog(m, int32(3), int32(_a871), v8+int32(16))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return
									} else {
										v145 = *(*int32)(unsafe.Add(mBase, _consts[15]))
										if int32(3) < v145 {
											m.Env.Exit(m, int32(1))
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										} else {
											F__serverLog(m, int32(3), int32(_a872), int32(0))
											mBase = m.M
											v152 = m.ExcPending
											if v152 != 0 {
												return
											} else {
												m.Env.Exit(m, int32(1))
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							if int32(3) < v67 {
								v81 = v63
								*(*int32)(unsafe.Add(mBase, uint32(v8)+124)) = v81
								*(*int32)(unsafe.Add(mBase, uint32(v8)+120)) = int32(_a188)
								v87 = m.G0
								v88 = int32(96)
								v89 = v87 - v88
								m.G0 = v89
								v91 = F_stat(m, v81, v89)
								mBase = m.M
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
								m.G0 = v89 + v88
								if base.B2i32(v91 != int32(-1))&base.B2i32(v92&int32(61440) == int32(4096)) != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v81
									F_rdbCheckError(m, int32(_a873), v8+int32(32))
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return
									} else {
										m.G0 = v8 + int32(1168)
										return
									}
								} else {
									v107 = F_redis_check_rdb_main(m, int32(2), v8+int32(120), int32(0))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										v145 = *(*int32)(unsafe.Add(mBase, _consts[15]))
										if int32(3) < v145 {
											m.Env.Exit(m, int32(1))
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										} else {
											F__serverLog(m, int32(3), int32(_a872), int32(0))
											mBase = m.M
											v152 = m.ExcPending
											if v152 != 0 {
												return
											} else {
												m.Env.Exit(m, int32(1))
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v8 + int32(128)
								F__serverLog(m, int32(3), int32(_a57), v8+int32(48))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, _consts[476]))
									v81 = v80
									*(*int32)(unsafe.Add(mBase, uint32(v8)+124)) = v81
									*(*int32)(unsafe.Add(mBase, uint32(v8)+120)) = int32(_a188)
									v87 = m.G0
									v88 = int32(96)
									v89 = v87 - v88
									m.G0 = v89
									v91 = F_stat(m, v81, v89)
									mBase = m.M
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
									m.G0 = v89 + v88
									if base.B2i32(v91 != int32(-1))&base.B2i32(v92&int32(61440) == int32(4096)) != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v81
										F_rdbCheckError(m, int32(_a873), v8+int32(32))
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
											return
										} else {
											m.G0 = v8 + int32(1168)
											return
										}
									} else {
										v107 = F_redis_check_rdb_main(m, int32(2), v8+int32(120), int32(0))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return
										} else {
											v145 = *(*int32)(unsafe.Add(mBase, _consts[15]))
											if int32(3) < v145 {
												m.Env.Exit(m, int32(1))
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											} else {
												F__serverLog(m, int32(3), int32(_a872), int32(0))
												mBase = m.M
												v152 = m.ExcPending
												if v152 != 0 {
													return
												} else {
													m.Env.Exit(m, int32(1))
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
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v8 + int32(128)
						F_rdbCheckError(m, int32(_a57), v8+int32(64))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							v145 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							if int32(3) < v145 {
								m.Env.Exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							} else {
								F__serverLog(m, int32(3), int32(_a872), int32(0))
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return
								} else {
									m.Env.Exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, _consts[15]))
					if int32(1) < v38 {
						m.G0 = v8 + int32(1168)
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v8 + int32(128)
						F__serverLog(m, int32(1), int32(_a57), v8+int32(80))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							m.G0 = v8 + int32(1168)
							return
						}
					}
				}
			}
		}
	}
}
func F_rdbSaveAuxField(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v100 int32
	_ = v100
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int32(250)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v15)
	if l0 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v100
L2:
	;
	v100 = int32(-1)
	goto L1
L3:
	;
	v69 = F_rdbSaveRawString(m, l0, l1, l2)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L16
	} else {
		goto L22
	}
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v19&int32(6) != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v31 = v13 + int32(15)
	v32 = int32(1)
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v35) < base.Ui32(v32) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	v37 = v35
	goto L10
L9:
	;
	v37 = v32
	goto L10
L10:
	;
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = v37
	goto L13
L12:
	;
	v38 = v32
	goto L13
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v39 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v47 = m.T0[v46].(func(*base.Module, int32, int32, int32) int32)(m, l0, v31, v38)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L16
	} else {
		goto L19
	}
L15:
	;
	m.T0[v39].(func(*base.Module, int32, int32, int32))(m, l0, v31, v38)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	goto L14
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v53 + v38
	v57 = v32 - v38
	if v57 != 0 {
		v31 = v31 + v38
		v32 = v57
		goto L6
	} else {
		goto L21
	}
L19:
	;
	if v47 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v49 | int64(2)
	goto L2
L21:
	;
	goto L7
L22:
	;
	if v69 == int32(-1) {
		v100 = int32(-1)
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v74 = F_rdbSaveRawString(m, l0, l3, l4)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	if v74 == int32(-1) {
		v100 = int32(-1)
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v100 = v69 + v74 + int32(1)
	goto L1
}
func F_rdbSaveBinaryDoubleValue(m *base.Module, l0 int32, l1 float64) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = l1
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v66
L2:
	;
	v66 = int32(8)
	goto L1
L3:
	;
	v16 = int32(-1)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v17&int32(6) != 0 {
		v66 = v16
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = int32(8)
	v27 = v11 + v20
	v28 = v20
	goto L5
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v31) < base.Ui32(v28) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	v33 = v31
	goto L9
L8:
	;
	v33 = v28
	goto L9
L9:
	;
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v34 = v33
	goto L12
L11:
	;
	v34 = v28
	goto L12
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v35 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = m.T0[v42].(func(*base.Module, int32, int32, int32) int32)(m, l0, v27, v34)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L15
	} else {
		goto L18
	}
L14:
	;
	m.T0[v35].(func(*base.Module, int32, int32, int32))(m, l0, v27, v34)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	goto L13
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v49 + v34
	v53 = v28 - v34
	if v53 != 0 {
		v27 = v27 + v34
		v28 = v53
		goto L5
	} else {
		goto L20
	}
L18:
	;
	if v43 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v45 | int64(2)
	v66 = v16
	goto L1
L20:
	;
	goto L6
}
func F_rdbSaveBinaryFloatValue(m *base.Module, l0 int32, l1 float32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*float32)(unsafe.Add(mBase, uint32(v11)+12)) = l1
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v66
L2:
	;
	v66 = int32(4)
	goto L1
L3:
	;
	v16 = int32(-1)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v17&int32(6) != 0 {
		v66 = v16
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = v11 + int32(12)
	v28 = int32(4)
	goto L5
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v31) < base.Ui32(v28) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	v33 = v31
	goto L9
L8:
	;
	v33 = v28
	goto L9
L9:
	;
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v34 = v33
	goto L12
L11:
	;
	v34 = v28
	goto L12
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v35 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = m.T0[v42].(func(*base.Module, int32, int32, int32) int32)(m, l0, v27, v34)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L15
	} else {
		goto L18
	}
L14:
	;
	m.T0[v35].(func(*base.Module, int32, int32, int32))(m, l0, v27, v34)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	goto L13
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v49 + v34
	v53 = v28 - v34
	if v53 != 0 {
		v27 = v27 + v34
		v28 = v53
		goto L5
	} else {
		goto L20
	}
L18:
	;
	if v43 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v45 | int64(2)
	v66 = v16
	goto L1
L20:
	;
	goto L6
}
func F_rdbSaveDb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v391 int32
	_ = v391
	var v401 int64
	_ = v401
	var v402 int64
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v419 int64
	_ = v419
	var v421 int64
	_ = v421
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v465 int32
	_ = v465
	var v490 int32
	_ = v490
	v6 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	v24 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+l1<<(uint(int32(2))%32))))
	if v28 == v6 {
		v490 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(32)
	return v490
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v32 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)+72))
	goto L8
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v37 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v31)+40))
	v42 = v35
	goto L3
L6:
	;
	v39 = F_hashtableSize(m, v37)
	mBase = m.M
	v42 = base.I64_extend_i32_u(v39)
	goto L3
L7:
	;
	v42 = int64(0)
	goto L3
L8:
	;
	v45 = v42 + v44
	if v45 == int64(0) {
		v490 = v6
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v48 = int32(254)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+16)) = uint8(v48)
	if l0 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v490 = int32(-1)
	goto L1
L11:
	;
	v115 = int32(-1)
	v117 = F_rdbSaveLen(m, l0, base.I64_extend_i32_s(l1))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L24
	} else {
		goto L30
	}
L12:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v52&int32(6) != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v67 = v20 + int32(16)
	v68 = int32(1)
	goto L14
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v75) < base.Ui32(v68) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L11
L16:
	;
	v77 = v75
	goto L18
L17:
	;
	v77 = v68
	goto L18
L18:
	;
	if v75 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v78 = v77
	goto L21
L20:
	;
	v78 = v68
	goto L21
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v79 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v87 = m.T0[v86].(func(*base.Module, int32, int32, int32) int32)(m, l0, v67, v78)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L24
	} else {
		goto L27
	}
L23:
	;
	m.T0[v79].(func(*base.Module, int32, int32, int32))(m, l0, v67, v78)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	goto L22
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v93 + v78
	v97 = v68 - v78
	if v97 != 0 {
		v67 = v67 + v78
		v68 = v97
		goto L14
	} else {
		goto L29
	}
L27:
	;
	if v87 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v89 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v89 | int64(2)
	goto L10
L29:
	;
	goto L15
L30:
	;
	if v117 < int32(0) {
		v490 = v115
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	if v122 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v133)+72))
	goto L37
L33:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v127 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v121)+40))
	v132 = v125
	goto L32
L35:
	;
	v129 = F_hashtableSize(m, v127)
	mBase = m.M
	v132 = base.I64_extend_i32_u(v129)
	goto L32
L36:
	;
	v132 = int64(0)
	goto L32
L37:
	;
	v135 = int32(251)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+16)) = uint8(v135)
	if l0 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v200 = F_rdbSaveLen(m, l0, v45)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L24
	} else {
		goto L56
	}
L39:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v139&int32(6) != 0 {
		v490 = v115
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v154 = int32(1)
	v156 = v20 + int32(16)
	goto L41
L41:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v162) < base.Ui32(v154) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L38
L43:
	;
	v164 = v162
	goto L45
L44:
	;
	v164 = v154
	goto L45
L45:
	;
	if v162 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v165 = v164
	goto L48
L47:
	;
	v165 = v154
	goto L48
L48:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v166 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v172 = m.T0[v171].(func(*base.Module, int32, int32, int32) int32)(m, l0, v156, v165)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L24
	} else {
		goto L53
	}
L50:
	;
	m.T0[v166].(func(*base.Module, int32, int32, int32))(m, l0, v156, v165)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L24
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v178 + v165
	v182 = v154 - v165
	if v182 != 0 {
		v154 = v182
		v156 = v156 + v165
		goto L41
	} else {
		goto L55
	}
L53:
	;
	if v172 != 0 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v174 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v174 | int64(2)
	v490 = v115
	goto L1
L55:
	;
	goto L42
L56:
	;
	if v200 < int32(0) {
		v490 = v115
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v205 = F_rdbSaveLen(m, l0, v134+v132)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L24
	} else {
		goto L58
	}
L58:
	;
	if v205 < int32(0) {
		v490 = v115
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v212 = v117 + v200 + v205 + int32(2)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v215 = F_kvstoreIteratorInit(m, v213, int32(11))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L24
	} else {
		goto L61
	}
L60:
	;
	F_kvstoreIteratorRelease(m, v215)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L24
	} else {
		goto L124
	}
L61:
	;
	v219 = F_kvstoreIteratorNext(m, v215, v20+int32(28))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L24
	} else {
		goto L62
	}
L62:
	;
	if v219 == int32(0) {
		v453 = v212
		goto L60
	} else {
		goto L63
	}
L63:
	;
	if l2&int32(1) != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v227 = int32(_a88)
	goto L66
L65:
	;
	v227 = int32(_a880)
	goto L66
L66:
	;
	v235 = v212
	v244 = int32(-1)
	goto L69
L67:
	;
	v444 = int32(-1)
	if v215 == int32(0) {
		v490 = v444
		goto L1
	} else {
		goto L123
	}
L68:
	;
	F_sdsfree(m, v288)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L24
	} else {
		goto L122
	}
L69:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v247 = F_kvstoreIteratorGetCurrentHashtableIndex(m, v215)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L24
	} else {
		goto L71
	}
L71:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v250 == int32(0) {
		v355 = v235
		v357 = v244
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v358 = int32(0)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	if v361&int32(2) == v358 {
		v381 = v358
		goto L106
	} else {
		goto L107
	}
L73:
	;
	if v247 == v244 {
		v355 = v235
		v357 = v244
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v254 = F_sdsempty(m)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L24
	} else {
		goto L75
	}
L75:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257+v247<<(uint(int32(2))%32))))
	if v261 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+8))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v266+v247<<(uint(int32(2))%32))))
	if v270 != 0 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	v263 = F_hashtableSize(m, v261)
	mBase = m.M
	v264 = v263
	goto L76
L78:
	;
	v264 = int32(0)
	goto L76
L79:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+8))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275+v247<<(uint(int32(2))%32))))
	if v279 != 0 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v272 = F_hashtableSize(m, v270)
	mBase = m.M
	v273 = v272
	goto L79
L81:
	;
	v273 = int32(0)
	goto L79
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v273
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v247
	v288 = F_sdscatprintf(m, v254, int32(_a881), v20)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L24
	} else {
		goto L85
	}
L83:
	;
	v281 = F_hashtableSize(m, v279)
	mBase = m.M
	v282 = v281
	goto L82
L84:
	;
	v282 = int32(0)
	goto L82
L85:
	;
	if v288&int32(3) == int32(0) {
		v313 = v288
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v347 = F_rdbSaveAuxField(m, l0, int32(_a882), int32(9), v288, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L24
	} else {
		goto L102
	}
L87:
	;
	v346 = v338 - v288
	goto L86
L88:
	;
	v317 = v313
	goto L96
L89:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	if v299 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v302 = v288
	goto L92
L91:
	;
	v346 = v288 - v288
	goto L86
L92:
	;
	v306 = v302 + int32(1)
	if v306&int32(3) == int32(0) {
		v313 = v306
		goto L88
	} else {
		goto L94
	}
L94:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	if v311 != 0 {
		v302 = v306
		goto L92
	} else {
		goto L95
	}
L95:
	;
	v338 = v306
	goto L87
L96:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	v326 = int32(-2139062144)
	if (int32(16843008)-v323|v323)&v326 == v326 {
		v317 = v317 + int32(4)
		goto L96
	} else {
		goto L98
	}
L97:
	;
	v332 = v317
	goto L99
L98:
	;
	goto L97
L99:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	if v336 != 0 {
		v332 = v332 + int32(1)
		goto L99
	} else {
		goto L101
	}
L100:
	;
	v338 = v332
	goto L87
L101:
	;
	goto L100
L102:
	;
	if v347 < int32(0) {
		goto L68
	} else {
		goto L103
	}
L103:
	;
	F_sdsfree(m, v288)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L24
	} else {
		goto L104
	}
L104:
	;
	v355 = v347 + v235
	v357 = v247
	goto L72
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v381
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = int64(-68719476736)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	if v391&int32(1) == int32(0) {
		v402 = int64(-1)
		goto L109
	} else {
		goto L110
	}
L106:
	;
	goto L105
L107:
	;
	v375 = v246 + (v361&int32(4) ^ int32(12)) + v361<<(uint(int32(3))%32)&int32(8)
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375))))
	v381 = v375 + v376 + int32(1)
	goto L106
L108:
	;
	v403 = F_rdbSaveKeyValuePair(m, l0, v20+int32(16), v246, v402, l1, l3)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L24
	} else {
		goto L111
	}
L109:
	;
	goto L108
L110:
	;
	v401 = *(*int64)(unsafe.Add(mBase, uint32(v246+(v391&int32(4)^int32(12)))))
	v402 = v401
	goto L109
L111:
	;
	if v403 <= int32(-1) {
		goto L67
	} else {
		goto L112
	}
L112:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	if v408 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v413 + int32(1)
	if v413&int32(1023) != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	goto L113
L116:
	;
	v432 = v403 + v355
	v435 = F_kvstoreIteratorNext(m, v215, v20+int32(28))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L24
	} else {
		goto L120
	}
L117:
	;
	v419 = F_mstime(m)
	mBase = m.M
	v421 = *(*int64)(unsafe.Add(mBase, _consts[481]))
	if v419-v421 < int64(1000) {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	F_sendChildInfo(m, int32(0), v426, v227)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L24
	} else {
		goto L119
	}
L119:
	;
	*(*int64)(unsafe.Add(mBase, _consts[481])) = v419
	goto L116
L120:
	;
	if v435 != 0 {
		v235 = v432
		v244 = v357
		goto L69
	} else {
		goto L121
	}
L121:
	;
	v453 = v432
	goto L60
L122:
	;
	goto L67
L123:
	;
	v453 = v444
	goto L60
L124:
	;
	v490 = v453
	goto L1
}
func F_rdbSaveFunctions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	goto L1
L1:
	;
	v18 = F_dictGetIterator(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	v24 = v2
	goto L5
L4:
	;
	F_dictReleaseIterator(m, v18)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L2
	} else {
		goto L63
	}
L5:
	;
	v38 = v18 + int32(20)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v233 = int32(-1)
	goto L4
L7:
	;
	v138 = int32(245)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v138)
	if l0 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L8:
	;
	if v134 != 0 {
		goto L7
	} else {
		goto L34
	}
L9:
	;
	v45 = v38
	v46 = v42
	goto L12
L10:
	;
	v42 = int32(1)
	goto L9
L11:
	;
	v42 = int32(0)
	goto L9
L12:
	;
	switch v46 {
	case 0:
		goto L17
	default:
		goto L16
	}
L14:
	;
	v46 = int32(0)
	goto L12
L15:
	;
	goto L8
L16:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v126
	if v126 == int32(0) {
		goto L14
	} else {
		goto L33
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v50 != int32(-1) {
		v89 = v50
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v90 = int32(1)
	v91 = v89 + v90
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v91
	v93 = int32(0)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v97+int32(26)))))
	if v101 == int32(255) {
		goto L27
	} else {
		goto L28
	}
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v54 != 0 {
		v89 = int32(-1)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v56 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v83 != int32(-1) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v63 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v55)+16)))
	v64 = int64(*(*int8)(unsafe.Add(mBase, uint32(v55)+27)))
	v65 = int64(*(*int32)(unsafe.Add(mBase, uint32(v55)+8)))
	v66 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v55)+12)))
	v67 = int64(*(*int8)(unsafe.Add(mBase, uint32(v55)+26)))
	v68 = int64(*(*int32)(unsafe.Add(mBase, uint32(v55)+4)))
	v69 = F_wangHash64(m, v68)
	mBase = m.M
	v71 = F_wangHash64(m, v67+v69)
	mBase = m.M
	v73 = F_wangHash64(m, v66+v71)
	mBase = m.M
	v75 = F_wangHash64(m, v65+v73)
	mBase = m.M
	v77 = F_wangHash64(m, v64+v75)
	mBase = m.M
	v79 = F_wangHash64(m, v63+v77)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v82 = v81
	goto L21
L23:
	;
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+24)))
	v61 = v59 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v55)+24)) = uint16(v61)
	v82 = v55
	goto L21
L24:
	;
	v89 = v83 + int32(-1)
	goto L18
L25:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v89 = v86
	goto L18
L26:
	;
	v116 = int32(2)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v96+v114<<(uint(v116)%32)+int32(4))))
	v45 = v121 + v115<<(uint(v116)%32)
	v46 = int32(1)
	goto L12
L27:
	;
	v105 = v93
	goto L29
L28:
	;
	v105 = v90 << (uint(v101) % 32)
	goto L29
L29:
	;
	if v91 < v105 {
		v114 = v97
		v115 = v91
		goto L26
	} else {
		goto L30
	}
L30:
	;
	if v97 != 0 {
		v134 = v93
		goto L15
	} else {
		goto L31
	}
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	if v107 == int32(-1) {
		v134 = v93
		goto L15
	} else {
		goto L32
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+4)) = int64(4294967296)
	v114 = int32(1)
	v115 = int32(0)
	goto L26
L33:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v130
	v134 = v126
	goto L15
L34:
	;
	v233 = v24
	goto L4
L35:
	;
	goto L6
L36:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	goto L60
L37:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v142&int32(6) != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v154 = v12 + int32(15)
	v155 = int32(1)
	goto L39
L39:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v157) < base.Ui32(v155) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L36
L41:
	;
	v159 = v157
	goto L43
L42:
	;
	v159 = v155
	goto L43
L43:
	;
	if v157 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v160 = v159
	goto L46
L45:
	;
	v160 = v155
	goto L46
L46:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v161 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v167 = m.T0[v166].(func(*base.Module, int32, int32, int32) int32)(m, l0, v154, v160)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L51
	}
L48:
	;
	m.T0[v161].(func(*base.Module, int32, int32, int32))(m, l0, v154, v160)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v173 + v160
	v177 = v155 - v160
	if v177 != 0 {
		v154 = v154 + v160
		v155 = v177
		goto L39
	} else {
		goto L53
	}
L51:
	;
	if v167 != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v169 | int64(2)
	goto L35
L53:
	;
	goto L40
L54:
	;
	v212 = F_rdbSaveRawString(m, l0, v189, v209)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L61
	}
L55:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v189+int32(-17))))
	v209 = v208
	goto L54
L56:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v189+int32(-9))))
	v209 = v205
	goto L54
L57:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189+int32(-5)))))
	v209 = v202
	goto L54
L58:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189+int32(-3)))))
	v209 = v199
	goto L54
L59:
	;
	v209 = int32(base.Ui32(v192) >> (uint(int32(3)) % 32))
	goto L54
L60:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189+int32(-1)))))
	switch v192 & int32(7) {
	case 0:
		goto L59
	case 1:
		goto L58
	case 2:
		goto L57
	case 3:
		goto L56
	case 4:
		goto L55
	default:
		v209 = int32(0)
		goto L54
	}
L61:
	;
	v215 = int32(-1)
	if v215 < v212 {
		v24 = v24 + int32(1) + v212
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v233 = v215
	goto L4
L63:
	;
	m.G0 = v12 + int32(16)
	return v233
}
func F_rdbSaveKeyValuePair(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
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
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int64
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int64
	_ = v310
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int64
	_ = v399
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
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
	var v438 int32
	_ = v438
	var v463 int32
	_ = v463
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	if l3 == int64(-1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v16 + int32(48)
	return v463
L2:
	;
	v463 = int32(-1)
	goto L1
L3:
	;
	if v19&int32(1) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L4:
	;
	v22 = int32(252)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+40)) = uint8(v22)
	if l0 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v26&int32(6) != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v40 = v16 + int32(40)
	v41 = int32(1)
	goto L7
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v45) < base.Ui32(v41) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = l3
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v70&int32(6) != 0 {
		goto L2
	} else {
		goto L23
	}
L9:
	;
	v47 = v45
	goto L11
L10:
	;
	v47 = v41
	goto L11
L11:
	;
	if v45 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v48 = v47
	goto L14
L13:
	;
	v48 = v41
	goto L14
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v49 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v57 = m.T0[v56].(func(*base.Module, int32, int32, int32) int32)(m, l0, v40, v48)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L17
	} else {
		goto L20
	}
L16:
	;
	m.T0[v49].(func(*base.Module, int32, int32, int32))(m, l0, v40, v48)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	goto L15
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v64 + v48
	v68 = v41 - v48
	if v68 != 0 {
		v40 = v40 + v48
		v41 = v68
		goto L7
	} else {
		goto L22
	}
L20:
	;
	if v57 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v59 | int64(2)
	v463 = int32(-1)
	goto L1
L22:
	;
	goto L8
L23:
	;
	v84 = v16 + int32(40)
	v85 = int32(8)
	goto L24
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v89) < base.Ui32(v85) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L3
L26:
	;
	v91 = v89
	goto L28
L27:
	;
	v91 = v85
	goto L28
L28:
	;
	if v89 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v92 = v91
	goto L31
L30:
	;
	v92 = v85
	goto L31
L31:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v93 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v99 = m.T0[v98].(func(*base.Module, int32, int32, int32) int32)(m, l0, v84, v92)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L17
	} else {
		goto L36
	}
L33:
	;
	m.T0[v93].(func(*base.Module, int32, int32, int32))(m, l0, v84, v92)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L17
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v106 + v92
	v110 = v85 - v92
	if v110 != 0 {
		v84 = v84 + v92
		v85 = v110
		goto L24
	} else {
		goto L38
	}
L36:
	;
	if v99 != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v101 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v101 | int64(2)
	v463 = int32(-1)
	goto L1
L38:
	;
	goto L25
L39:
	;
	if v19&int32(2) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L40:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v131 = F_lru_getIdleSecs(m, int32(base.Ui32(v128)>>(uint(int32(8))%32)))
	mBase = m.M
	goto L41
L41:
	;
	v132 = int32(248)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+40)) = uint8(v132)
	if l0 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v192 = F_rdbSaveLen(m, l0, base.I64_extend_i32_u(v131))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L17
	} else {
		goto L60
	}
L43:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v136&int32(6) != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v150 = v16 + int32(40)
	v151 = int32(1)
	goto L45
L45:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v155) < base.Ui32(v151) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L42
L47:
	;
	v157 = v155
	goto L49
L48:
	;
	v157 = v151
	goto L49
L49:
	;
	if v155 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v158 = v157
	goto L52
L51:
	;
	v158 = v151
	goto L52
L52:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v159 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v165 = m.T0[v164].(func(*base.Module, int32, int32, int32) int32)(m, l0, v150, v158)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L17
	} else {
		goto L57
	}
L54:
	;
	m.T0[v159].(func(*base.Module, int32, int32, int32))(m, l0, v150, v158)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L17
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v172 + v158
	v176 = v151 - v158
	if v176 != 0 {
		v150 = v150 + v158
		v151 = v176
		goto L45
	} else {
		goto L59
	}
L57:
	;
	if v165 != 0 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v167 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v167 | int64(2)
	v463 = int32(-1)
	goto L1
L59:
	;
	goto L46
L60:
	;
	if v192 == int32(-1) {
		v463 = int32(-1)
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L39
L62:
	;
	v333 = F_rdbGetObjectType(m, l2, l5)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L17
	} else {
		goto L99
	}
L63:
	;
	v215 = m.G0
	v216 = int32(16)
	v217 = v215 - v216
	m.G0 = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v220 = int32(8)
	v224 = F_lfu_getFrequency(m, int32(base.Ui32(v219)>>(uint(v220)%32)), v217+int32(15))
	mBase = m.M
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v225 | v224<<(uint(v220)%32)
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+15)))
	m.G0 = v217 + v216
	goto L64
L64:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+39)) = uint8(v230)
	v235 = int32(249)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+40)) = uint8(v235)
	if l0 == int32(0) {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v239&int32(6) != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v253 = v16 + int32(40)
	v254 = int32(1)
	goto L67
L67:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v258) < base.Ui32(v254) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v279&int32(6) != 0 {
		goto L2
	} else {
		goto L82
	}
L69:
	;
	v260 = v258
	goto L71
L70:
	;
	v260 = v254
	goto L71
L71:
	;
	if v258 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v261 = v260
	goto L74
L73:
	;
	v261 = v254
	goto L74
L74:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v262 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v268 = m.T0[v267].(func(*base.Module, int32, int32, int32) int32)(m, l0, v253, v261)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L17
	} else {
		goto L79
	}
L76:
	;
	m.T0[v262].(func(*base.Module, int32, int32, int32))(m, l0, v253, v261)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L17
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v274 + v261
	v278 = v254 - v261
	if v278 != 0 {
		v253 = v253 + v261
		v254 = v278
		goto L67
	} else {
		goto L81
	}
L79:
	;
	if v268 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v270 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v270 | int64(2)
	goto L2
L81:
	;
	goto L68
L82:
	;
	v293 = v16 + int32(39)
	v294 = int32(1)
	goto L83
L83:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v298) < base.Ui32(v294) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L62
L85:
	;
	v300 = v298
	goto L87
L86:
	;
	v300 = v294
	goto L87
L87:
	;
	if v298 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v301 = v300
	goto L90
L89:
	;
	v301 = v294
	goto L90
L90:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v302 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v308 = m.T0[v307].(func(*base.Module, int32, int32, int32) int32)(m, l0, v293, v301)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L17
	} else {
		goto L95
	}
L92:
	;
	m.T0[v302].(func(*base.Module, int32, int32, int32))(m, l0, v293, v301)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L17
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v315 + v301
	v319 = v294 - v301
	if v319 != 0 {
		v293 = v293 + v301
		v294 = v319
		goto L83
	} else {
		goto L97
	}
L95:
	;
	if v308 != 0 {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v310 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v310 | int64(2)
	v463 = int32(-1)
	goto L1
L97:
	;
	goto L84
L98:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+40)) = uint8(v333)
	if l0 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L99:
	;
	if v333 != int32(-1) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v337 = int32(_a69)
	v338 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v340 = *(*int32)(unsafe.Add(mBase, _consts[151]))
	if v340 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v354 = int32(-1)
	if int32(3) < v338 {
		v463 = v354
		goto L1
	} else {
		goto L105
	}
L102:
	;
	v343 = int32(-1)
	if int32(3) < v338 {
		v463 = v343
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l4
	F__serverLog(m, int32(3), int32(_a877), v16+int32(16))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L17
	} else {
		goto L104
	}
L104:
	;
	v463 = v343
	goto L1
L105:
	;
	v357 = F_objectGetVal(m, l1)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v357
	F__serverLog(m, int32(3), int32(_a878), v16)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L17
	} else {
		goto L106
	}
L106:
	;
	v463 = v354
	goto L1
L107:
	;
	v423 = F_rdbSaveStringObject(m, l0, l1)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L17
	} else {
		goto L125
	}
L108:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v368&int32(6) != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	v382 = v16 + int32(40)
	v383 = int32(1)
	goto L110
L110:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v387) < base.Ui32(v383) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	goto L107
L112:
	;
	v389 = v387
	goto L114
L113:
	;
	v389 = v383
	goto L114
L114:
	;
	if v387 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v390 = v389
	goto L117
L116:
	;
	v390 = v383
	goto L117
L117:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v391 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v397 = m.T0[v396].(func(*base.Module, int32, int32, int32) int32)(m, l0, v382, v390)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L17
	} else {
		goto L122
	}
L119:
	;
	m.T0[v391].(func(*base.Module, int32, int32, int32))(m, l0, v382, v390)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L17
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v404 + v390
	v408 = v383 - v390
	if v408 != 0 {
		v382 = v382 + v390
		v383 = v408
		goto L110
	} else {
		goto L124
	}
L122:
	;
	if v397 != 0 {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v399 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v399 | int64(2)
	v463 = int32(-1)
	goto L1
L124:
	;
	goto L111
L125:
	;
	if v423 == int32(-1) {
		v463 = int32(-1)
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v428 = F_rdbSaveObject(m, l0, l2, l1, l4, v333)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L17
	} else {
		goto L127
	}
L127:
	;
	if v428 == int32(-1) {
		v463 = int32(-1)
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v432 = int32(1)
	v434 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	if v434 == int32(0) {
		v463 = v432
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_debugDelay(m, v434)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L17
	} else {
		goto L130
	}
L130:
	;
	v463 = v432
	goto L1
}
func F_rdbSaveLen(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v228 int64
	_ = v228
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int64
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v307 int32
	_ = v307
	v2 = l1
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if base.Ui64(int64(63)) < base.Ui64(v2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v307
L2:
	;
	if base.Ui64(int64(16383)) < base.Ui64(v2) {
		goto L23
	} else {
		goto L24
	}
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)) = uint8(v2)
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = int32(-1)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v18&int32(6) != 0 {
		v307 = v17
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v307 = int32(1)
	goto L1
L6:
	;
	v28 = v11 + int32(6)
	v29 = int32(1)
	goto L7
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v32) < base.Ui32(v29) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v307 = int32(1)
	goto L1
L9:
	;
	v34 = v32
	goto L11
L10:
	;
	v34 = v29
	goto L11
L11:
	;
	if v32 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v35 = v34
	goto L14
L13:
	;
	v35 = v29
	goto L14
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v36 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v44 = m.T0[v43].(func(*base.Module, int32, int32, int32) int32)(m, l0, v28, v35)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L17
	} else {
		goto L20
	}
L16:
	;
	m.T0[v36].(func(*base.Module, int32, int32, int32))(m, l0, v28, v35)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	goto L15
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v50 + v35
	v54 = v29 - v35
	if v54 != 0 {
		v28 = v28 + v35
		v29 = v54
		goto L7
	} else {
		goto L22
	}
L20:
	;
	if v44 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v46 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v46 | int64(2)
	v307 = v17
	goto L1
L22:
	;
	goto L8
L23:
	;
	if base.Ui64(int64(4294967295)) < base.Ui64(v2) {
		goto L43
	} else {
		goto L44
	}
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)) = uint8(v2)
	v63 = int32(base.Ui32(base.I32_wrap_i64(v2))>>(uint(int32(8))%32)) | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)) = uint8(v63)
	if l0 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v67&int32(6) != 0 {
		v307 = int32(-1)
		goto L1
	} else {
		goto L27
	}
L26:
	;
	v307 = int32(2)
	goto L1
L27:
	;
	v77 = v11 + int32(6)
	v78 = int32(2)
	goto L28
L28:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v81) < base.Ui32(v78) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v83 = v81
	goto L32
L31:
	;
	v83 = v78
	goto L32
L32:
	;
	if v81 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v84 = v83
	goto L35
L34:
	;
	v84 = v78
	goto L35
L35:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v85 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, l0, v77, v84)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L17
	} else {
		goto L40
	}
L37:
	;
	m.T0[v85].(func(*base.Module, int32, int32, int32))(m, l0, v77, v84)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L17
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v98 + v84
	v103 = v78 - v84
	if v103 != 0 {
		v77 = v77 + v84
		v78 = v103
		goto L28
	} else {
		goto L42
	}
L40:
	;
	if v91 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v93 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v93 | int64(2)
	v307 = int32(-1)
	goto L1
L42:
	;
	v307 = int32(2)
	goto L1
L43:
	;
	v184 = int32(129)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)) = uint8(v184)
	if l0 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L44:
	;
	v106 = int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)) = uint8(v106)
	if l0 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v111 = int32(-1)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v112&int32(6) != 0 {
		v307 = v111
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v109 = F___bswap_32_1(m, base.I32_wrap_i64(v2))
	mBase = m.M
	goto L47
L47:
	;
	v307 = int32(5)
	goto L1
L48:
	;
	v122 = v11 + int32(6)
	v123 = int32(1)
	goto L49
L49:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v126) < base.Ui32(v123) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v148 = F___bswap_32_1(m, base.I32_wrap_i64(v2))
	mBase = m.M
	goto L64
L51:
	;
	v128 = v126
	goto L53
L52:
	;
	v128 = v123
	goto L53
L53:
	;
	if v126 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v129 = v128
	goto L56
L55:
	;
	v129 = v123
	goto L56
L56:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v130 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v136 = m.T0[v135].(func(*base.Module, int32, int32, int32) int32)(m, l0, v122, v129)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L17
	} else {
		goto L61
	}
L58:
	;
	m.T0[v130].(func(*base.Module, int32, int32, int32))(m, l0, v122, v129)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L17
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v142 + v129
	v146 = v123 - v129
	if v146 != 0 {
		v122 = v122 + v129
		v123 = v146
		goto L49
	} else {
		goto L63
	}
L61:
	;
	if v136 != 0 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v138 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v138 | int64(2)
	v307 = v111
	goto L1
L63:
	;
	goto L50
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v148
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v150&int32(6) != 0 {
		v307 = v111
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v158 = v11
	v159 = int32(4)
	goto L66
L66:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v162) < base.Ui32(v159) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v307 = int32(5)
	goto L1
L68:
	;
	v164 = v162
	goto L70
L69:
	;
	v164 = v159
	goto L70
L70:
	;
	if v162 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v165 = v164
	goto L73
L72:
	;
	v165 = v159
	goto L73
L73:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v166 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v172 = m.T0[v171].(func(*base.Module, int32, int32, int32) int32)(m, l0, v158, v165)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L17
	} else {
		goto L78
	}
L75:
	;
	m.T0[v166].(func(*base.Module, int32, int32, int32))(m, l0, v158, v165)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L17
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v178 + v165
	v182 = v159 - v165
	if v182 != 0 {
		v158 = v158 + v165
		v159 = v182
		goto L66
	} else {
		goto L80
	}
L78:
	;
	if v172 != 0 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v174 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v174 | int64(2)
	v307 = v111
	goto L1
L80:
	;
	goto L67
L81:
	;
	v307 = int32(9)
	goto L1
L82:
	;
	v188 = int32(-1)
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v189&int32(6) != 0 {
		v307 = v188
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v199 = v11 + int32(6)
	v200 = int32(1)
	goto L84
L84:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v203) < base.Ui32(v200) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v224 = int64(56)
	v226 = int64(65280)
	v228 = int64(40)
	v231 = int64(16711680)
	v233 = int64(24)
	v235 = int64(4278190080)
	v237 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v2<<(uint(v224)%64) | v2&v226<<(uint(v228)%64) | (v2&v231<<(uint(v233)%64) | v2&v235<<(uint(v237)%64)) | (int64(base.Ui64(v2)>>(uint(v237)%64))&v235 | int64(base.Ui64(v2)>>(uint(v233)%64))&v231 | (int64(base.Ui64(v2)>>(uint(v228)%64))&v226 | int64(base.Ui64(v2)>>(uint(v224)%64))))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v260&int32(6) != 0 {
		v307 = v188
		goto L1
	} else {
		goto L99
	}
L86:
	;
	v205 = v203
	goto L88
L87:
	;
	v205 = v200
	goto L88
L88:
	;
	if v203 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v206 = v205
	goto L91
L90:
	;
	v206 = v200
	goto L91
L91:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v207 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v213 = m.T0[v212].(func(*base.Module, int32, int32, int32) int32)(m, l0, v199, v206)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L17
	} else {
		goto L96
	}
L93:
	;
	m.T0[v207].(func(*base.Module, int32, int32, int32))(m, l0, v199, v206)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L17
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v219 + v206
	v223 = v200 - v206
	if v223 != 0 {
		v199 = v199 + v206
		v200 = v223
		goto L84
	} else {
		goto L98
	}
L96:
	;
	if v213 != 0 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v215 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v215 | int64(2)
	v307 = v188
	goto L1
L98:
	;
	goto L85
L99:
	;
	v263 = int32(8)
	v270 = v11 + v263
	v271 = v263
	goto L100
L100:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v274) < base.Ui32(v271) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L81
L102:
	;
	v276 = v274
	goto L104
L103:
	;
	v276 = v271
	goto L104
L104:
	;
	if v274 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v277 = v276
	goto L107
L106:
	;
	v277 = v271
	goto L107
L107:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v278 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v284 = m.T0[v283].(func(*base.Module, int32, int32, int32) int32)(m, l0, v270, v277)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L17
	} else {
		goto L112
	}
L109:
	;
	m.T0[v278].(func(*base.Module, int32, int32, int32))(m, l0, v270, v277)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L17
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v290 + v277
	v294 = v271 - v277
	if v294 != 0 {
		v270 = v270 + v277
		v271 = v294
		goto L100
	} else {
		goto L114
	}
L112:
	;
	if v284 != 0 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v286 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v286 | int64(2)
	v307 = v188
	goto L1
L114:
	;
	goto L101
}
func F_rdbSaveLzfStringObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v533 int32
	_ = v533
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	v4 = int32(0)
	if base.Ui32(l2) < base.Ui32(int32(5)) {
		v558 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v558
L2:
	;
	v11 = l2 + int32(-4)
	if base.Ui32(int32(8191)) < base.Ui32(v11) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v30 = int32(0)
	v43 = m.G0
	v45 = v43 - int32(262144)
	m.G0 = v45
	if l2 == v30 {
		v533 = v30
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v25 = F_valkey_malloc(m, l2+int32(-3))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L9
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[477]))
	if v15 != 0 {
		v29 = v15
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v18 = F_valkey_malloc(m, int32(8192))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[477])) = v18
	v29 = v18
	goto L3
L9:
	;
	if v25 == int32(0) {
		v558 = v4
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v29 = v25
	goto L3
L11:
	;
	v553 = *(*int32)(unsafe.Add(mBase, _consts[477]))
	if v29 == v553 {
		v558 = v551
		goto L1
	} else {
		goto L108
	}
L12:
	;
	if v533 == int32(0) {
		v551 = v30
		goto L11
	} else {
		goto L106
	}
L13:
	;
	m.G0 = v45 + int32(262144)
	goto L12
L14:
	;
	if v11 == int32(0) {
		v533 = v30
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v52 = v29 + v11
	v54 = v29 + int32(1)
	v56 = l1 + l2
	v58 = v56 + int32(-2)
	if base.Ui32(l1) < base.Ui32(v58) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if base.Ui32(v52) < base.Ui32(v406+int32(3)) {
		v533 = int32(0)
		goto L13
	} else {
		goto L87
	}
L17:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v68 = l1
	v70 = v54
	v76 = int32(0)
	v77 = v61<<(uint(int32(8))%32) | v64
	goto L19
L18:
	;
	v404 = l1
	v406 = v54
	v412 = int32(0)
	v415 = v30
	goto L16
L19:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+2)))
	v86 = v77<<(uint(int32(8))%32) | v85
	v94 = v45 + (v86*int32(65531)+v77)&int32(65535)<<(uint(int32(2))%32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v68
	if base.Ui32(v95) <= base.Ui32(l1) {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v404 = v387
	v406 = v389
	v412 = v395
	v415 = v398
	goto L16
L21:
	;
	if base.Ui32(v387) < base.Ui32(v58) {
		v68 = v387
		v70 = v389
		v76 = v395
		v77 = v396
		goto L19
	} else {
		goto L86
	}
L22:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v329))) = uint8(v100)
	v344 = v329 + int32(2)
	v345 = v68 + v335
	if base.Ui32(v345) < base.Ui32(v58) {
		goto L84
	} else {
		goto L85
	}
L23:
	;
	v317 = v309 + int32(-9)
	*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)) = uint8(v317)
	v322 = int32(base.Ui32(v100)>>(uint(int32(8))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v127))) = uint8(v322)
	v329 = v127 + int32(2)
	v335 = v309
	goto L22
L24:
	;
	if base.Ui32(v70) < base.Ui32(v52) {
		goto L80
	} else {
		goto L81
	}
L25:
	;
	v100 = v95 ^ int32(-1) + v68
	if base.Ui32(int32(8191)) < base.Ui32(v100) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+2)))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+2)))
	if v103 != v104 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95))))
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68))))
	if v106 != v107 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if base.Ui32(v70+int32(4)) < base.Ui32(v52) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v119 = int32(-1)
	v123 = v76 + v119
	*(*uint8)(unsafe.Add(mBase, uint32(v70+(v76^v119)))) = uint8(v123)
	v127 = v70 - base.B2i32(v76 == int32(0))
	v130 = v56 - v68 + int32(-2)
	if base.Ui32(int32(17)) <= base.Ui32(v130) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	if base.Ui32(v70-base.B2i32(v76 == int32(0))+int32(4)) < base.Ui32(v52) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v533 = int32(0)
	goto L13
L32:
	;
	v277 = v267<<(uint(int32(5))%32) | int32(base.Ui32(v100)>>(uint(int32(8))%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v127))) = uint8(v277)
	v329 = v127 + int32(1)
	v335 = v266
	goto L22
L33:
	;
	v217 = int32(264)
	if base.Ui32(v130) < base.Ui32(v217) {
		goto L67
	} else {
		goto L68
	}
L34:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+3)))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+3)))
	if v134 == v135 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v216 = int32(2)
	goto L33
L36:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+4)))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+4)))
	if v141 == v142 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v266 = int32(3)
	v267 = int32(1)
	goto L32
L38:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+5)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+5)))
	if v148 == v149 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v266 = int32(4)
	v267 = int32(2)
	goto L32
L40:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+6)))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+6)))
	if v155 == v156 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v266 = int32(5)
	v267 = int32(3)
	goto L32
L42:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+7)))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+7)))
	if v162 == v163 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v266 = int32(6)
	v267 = int32(4)
	goto L32
L44:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+8)))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+8)))
	if v169 == v170 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v266 = int32(7)
	v267 = int32(5)
	goto L32
L46:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+9)))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+9)))
	if v176 == v177 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v266 = int32(8)
	v267 = int32(6)
	goto L32
L48:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+10)))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+10)))
	if v180 == v181 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v309 = int32(9)
	goto L23
L50:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+11)))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+11)))
	if v184 == v185 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v309 = int32(10)
	goto L23
L52:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+12)))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+12)))
	if v188 == v189 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v309 = int32(11)
	goto L23
L54:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+13)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+13)))
	if v192 == v193 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v309 = int32(12)
	goto L23
L56:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+14)))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+14)))
	if v196 == v197 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v309 = int32(13)
	goto L23
L58:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+15)))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+15)))
	if v200 == v201 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v309 = int32(14)
	goto L23
L60:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+16)))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+16)))
	if v204 == v205 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v309 = int32(15)
	goto L23
L62:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+17)))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+17)))
	if v208 == v209 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v309 = int32(16)
	goto L23
L64:
	;
	v212 = int32(18)
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+18)))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+18)))
	if v213 != v214 {
		v309 = v212
		goto L23
	} else {
		goto L66
	}
L65:
	;
	v309 = int32(17)
	goto L23
L66:
	;
	v216 = v212
	goto L33
L67:
	;
	v220 = v130
	goto L69
L68:
	;
	v220 = v217
	goto L69
L69:
	;
	v222 = v216 | int32(1)
	if base.Ui32(v222) < base.Ui32(v220) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v224 = v220
	goto L72
L71:
	;
	v224 = v222
	goto L72
L72:
	;
	v236 = v216
	goto L74
L73:
	;
	v254 = v252 + int32(-1)
	if base.Ui32(int32(6)) < base.Ui32(v254) {
		v309 = v251
		goto L23
	} else {
		goto L79
	}
L74:
	;
	v244 = v236 + int32(1)
	if base.Ui32(v244) < base.Ui32(v220) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v251 = v244
	v252 = v236
	goto L73
L76:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v244))))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v244))))
	if v247 == v249 {
		v236 = v244
		goto L74
	} else {
		goto L78
	}
L77:
	;
	v251 = v224
	v252 = v224 + int32(-1)
	goto L73
L78:
	;
	goto L75
L79:
	;
	v266 = v251
	v267 = v254
	goto L32
L80:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v284)
	v286 = int32(1)
	v287 = v68 + v286
	v289 = v76 + v286
	if v289 == int32(32) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v533 = int32(0)
	goto L13
L82:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v70+int32(-32)))) = uint8(v76)
	v387 = v287
	v389 = v70 + int32(2)
	v395 = int32(0)
	v396 = v86
	v398 = v289
	goto L21
L83:
	;
	v387 = v287
	v389 = v70 + int32(1)
	v395 = v289
	v396 = v86
	v398 = v289
	goto L21
L84:
	;
	v349 = v345 + int32(-1)
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	v351 = int32(8)
	v354 = v345 + int32(-2)
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354))))
	v358 = v350<<(uint(v351)%32) | v355<<(uint(int32(16))%32)
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345))))
	v360 = v358 | v359
	v361 = int32(65531)
	v366 = int32(65535)
	v368 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v45+(v360*v361+int32(base.Ui32(v358)>>(uint(v351)%32)))&v366<<(uint(v368)%32)))) = v354
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+1)))
	v375 = v360<<(uint(v351)%32) | v374
	*(*int32)(unsafe.Add(mBase, uint32(v45+(v375*v361+v360)&v366<<(uint(v368)%32)))) = v349
	v387 = v345
	v389 = v344
	v395 = int32(0)
	v396 = v375
	v398 = v349
	goto L21
L85:
	;
	v404 = v345
	v406 = v344
	v412 = int32(0)
	v415 = v100
	goto L16
L86:
	;
	goto L20
L87:
	;
	if base.Ui32(v404) < base.Ui32(v56) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v518 = int32(-1)
	v522 = v514 + v518
	*(*uint8)(unsafe.Add(mBase, uint32(v505+(v514^v518)))) = uint8(v522)
	v533 = v505 - base.B2i32(v514 == int32(0)) - v29
	goto L13
L89:
	;
	v423 = int32(1)
	v425 = v56 - v404
	if v425&v423 != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v505 = v406
	v514 = v412
	goto L88
L91:
	;
	if v56 == v404+v423 {
		v505 = v445
		v514 = v448
		goto L88
	} else {
		goto L96
	}
L92:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	*(*uint8)(unsafe.Add(mBase, uint32(v406))) = uint8(v428)
	v430 = int32(1)
	v431 = v404 + v430
	v433 = v412 + v430
	if v433 == int32(32) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v445 = v406
	v446 = v412
	v447 = v404
	v448 = v415
	goto L91
L94:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v406+int32(-32)))) = uint8(v412)
	v443 = int32(0)
	v445 = v406 + int32(2)
	v446 = v443
	v447 = v431
	v448 = v443
	goto L91
L95:
	;
	v445 = v406 + int32(1)
	v446 = v433
	v447 = v431
	v448 = v433
	goto L91
L96:
	;
	v454 = v445
	v461 = v447
	v463 = v446
	goto L97
L97:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	*(*uint8)(unsafe.Add(mBase, uint32(v454))) = uint8(v467)
	v470 = v463 + int32(1)
	if v470 == int32(32) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v505 = v497
	v514 = v498
	goto L88
L99:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v481))) = uint8(v483)
	v486 = v482 + int32(1)
	if v486 == int32(32) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v454+int32(-32)))) = uint8(v463)
	v481 = v454 + int32(2)
	v482 = int32(0)
	goto L99
L101:
	;
	v481 = v454 + int32(1)
	v482 = v470
	goto L99
L102:
	;
	v500 = v461 + int32(2)
	if v500 != v404+v425 {
		v454 = v497
		v461 = v500
		v463 = v498
		goto L97
	} else {
		goto L105
	}
L103:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v481+int32(-32)))) = uint8(v482)
	v497 = v481 + int32(2)
	v498 = int32(0)
	goto L102
L104:
	;
	v497 = v481 + int32(1)
	v498 = v486
	goto L102
L105:
	;
	goto L98
L106:
	;
	v549 = F_rdbSaveLzfBlob(m, l0, v29, v533, l2)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	v551 = v549
	goto L11
L108:
	;
	F_valkey_free(m, v29)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	v558 = v551
	goto L1
}
func F_rdbSaveObject(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 float64
	_ = v404
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int64
	_ = v439
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v535 int32
	_ = v535
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v697 int64
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int64
	_ = v733
	var v735 int64
	_ = v735
	var v740 int32
	_ = v740
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v800 int32
	_ = v800
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v832 int32
	_ = v832
	var v833 int64
	_ = v833
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int64
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v857 int64
	_ = v857
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
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
	var v910 int32
	_ = v910
	var v925 int32
	_ = v925
	var v926 int64
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int64
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int64
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int64
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v946 int64
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int64
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v956 int64
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v961 int64
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v968 int64
	_ = v968
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v993 int64
	_ = v993
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1036 int64
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1041 int64
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int64
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1085 int32
	_ = v1085
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int64
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int64
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1185 int32
	_ = v1185
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1204 int32
	_ = v1204
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1293 int32
	_ = v1293
	var v1334 int32
	_ = v1334
	v16 = m.G0
	v18 = v16 - int32(320)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v20 & int32(15) {
	case 0:
		goto L9
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L13
	case 5:
		goto L11
	case 6:
		goto L12
	default:
		goto L10
	}
L1:
	;
	m.G0 = v18 + int32(320)
	return v1334
L2:
	;
	v1334 = int32(-1)
	goto L1
L3:
	;
	v1334 = int32(-1)
	goto L1
L4:
	;
	F_hashtableCleanupIterator(m, v18)
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L20
	} else {
		goto L350
	}
L5:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1265 == int32(0) {
		goto L343
	} else {
		goto L344
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(1)
	goto L5
L7:
	;
	v1334 = v1221
	goto L1
L8:
	;
	F_hashtableCleanupIterator(m, v18)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L20
	} else {
		goto L342
	}
L9:
	;
	v1197 = F_rdbSaveStringObject(m, l0, l1)
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L20
	} else {
		goto L340
	}
L10:
	;
	F__serverPanic_1(m, int32(_a875), int32(1165), int32(_a123), int32(0))
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L20
	} else {
		goto L339
	}
L11:
	;
	v1104 = F_objectGetVal(m, l1)
	mBase = m.M
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1104)))
	v1106 = *(*int64)(unsafe.Add(mBase, uint32(v1105)))
	v1107 = F_rdbSaveLen(m, l0, v1106)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L20
	} else {
		goto L318
	}
L12:
	;
	v844 = int32(-1)
	v845 = F_objectGetVal(m, l1)
	mBase = m.M
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v845)+4))
	v847 = *(*int64)(unsafe.Add(mBase, uint32(v846)+8))
	goto L253
L13:
	;
	switch int32(base.Ui32(v20)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		goto L150
	default:
		goto L149
	case 9:
		goto L151
	}
L14:
	;
	switch int32(base.Ui32(v20)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L108
	default:
		goto L107
	case 4:
		goto L109
	}
L15:
	;
	switch int32(base.Ui32(v20)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		goto L78
	default:
		goto L75
	case 4:
		goto L77
	case 9:
		goto L76
	}
L16:
	;
	switch int32(base.Ui32(v20)>>(uint(int32(4))%32))&int32(15) + int32(-9) {
	case 0:
		goto L19
	default:
		goto L17
	case 2:
		goto L18
	}
L17:
	;
	F__serverPanic_1(m, int32(_a875), int32(909), int32(_a852), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L20
	} else {
		goto L74
	}
L18:
	;
	v95 = F_objectGetVal(m, l1)
	mBase = m.M
	v96 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v96)
	if l0 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L19:
	;
	v29 = F_objectGetVal(m, l1)
	mBase = m.M
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v32 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v29)+12)))
	v33 = F_rdbSaveLen(m, l0, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	if v33 == int32(-1) {
		v1334 = int32(-1)
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v30 == int32(0) {
		v1221 = v33
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v42 = v30
	v44 = v33
	goto L24
L24:
	;
	v57 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v42)+16)))
	v62 = F_rdbSaveLen(m, l0, int64(base.Ui64(v57)>>(uint(int64(18))%64))&int64(3))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	if v62 == int32(-1) {
		v1334 = int32(-1)
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v66&int32(196608) != int32(131072) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v93 = v62 + v44 + v90
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v94 != 0 {
		v42 = v94
		v44 = v93
		goto L24
	} else {
		goto L36
	}
L29:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v86 = F_rdbSaveRawString(m, l0, v84, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L20
	} else {
		goto L34
	}
L30:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v71 + int32(4)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	goto L31
L31:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v79 = F_rdbSaveLzfBlob(m, l0, v77, v75, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	if v79 != int32(-1) {
		v90 = v79
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v1334 = int32(-1)
	goto L1
L34:
	;
	if v86 == int32(-1) {
		v1334 = int32(-1)
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v90 = v86
	goto L28
L36:
	;
	v1221 = v93
	goto L7
L37:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	goto L71
L38:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v100&int32(6) != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v110 = int32(1)
	v111 = v18
	goto L40
L40:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v119) < base.Ui32(v110) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v140 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v140)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v142&int32(6) != 0 {
		goto L3
	} else {
		goto L55
	}
L42:
	;
	v121 = v119
	goto L44
L43:
	;
	v121 = v110
	goto L44
L44:
	;
	if v119 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v122 = v121
	goto L47
L46:
	;
	v122 = v110
	goto L47
L47:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v123 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v129 = m.T0[v128].(func(*base.Module, int32, int32, int32) int32)(m, l0, v111, v122)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L20
	} else {
		goto L52
	}
L49:
	;
	m.T0[v123].(func(*base.Module, int32, int32, int32))(m, l0, v111, v122)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L20
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v135 + v122
	v139 = v110 - v122
	if v139 != 0 {
		v110 = v139
		v111 = v111 + v122
		goto L40
	} else {
		goto L54
	}
L52:
	;
	if v129 != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v131 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v131 | int64(2)
	goto L2
L54:
	;
	goto L41
L55:
	;
	v152 = int32(1)
	v153 = v18
	goto L56
L56:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v161) < base.Ui32(v152) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L37
L58:
	;
	v163 = v161
	goto L60
L59:
	;
	v163 = v152
	goto L60
L60:
	;
	if v161 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v164 = v163
	goto L63
L62:
	;
	v164 = v152
	goto L63
L63:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v165 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v171 = m.T0[v170].(func(*base.Module, int32, int32, int32) int32)(m, l0, v153, v164)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L20
	} else {
		goto L68
	}
L65:
	;
	m.T0[v165].(func(*base.Module, int32, int32, int32))(m, l0, v153, v164)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L20
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v177 + v164
	v181 = v152 - v164
	if v181 != 0 {
		v152 = v181
		v153 = v153 + v164
		goto L56
	} else {
		goto L70
	}
L68:
	;
	if v171 != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v173 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v173 | int64(2)
	goto L3
L70:
	;
	goto L57
L71:
	;
	v199 = F_rdbSaveRawString(m, l0, v95, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	if v199 == int32(-1) {
		v1334 = int32(-1)
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v1221 = v199 + int32(2)
	goto L7
L74:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F__serverPanic_1(m, int32(_a875), int32(943), int32(_a853), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L20
	} else {
		goto L106
	}
L76:
	;
	v313 = F_objectGetVal(m, l1)
	mBase = m.M
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	goto L103
L77:
	;
	v301 = F_objectGetVal(m, l1)
	mBase = m.M
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	goto L100
L78:
	;
	v219 = F_objectGetVal(m, l1)
	mBase = m.M
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	goto L79
L79:
	;
	v224 = F_rdbSaveLen(m, l0, base.I64_extend_i32_u(v220+v221))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L20
	} else {
		goto L80
	}
L80:
	;
	if v224 == int32(-1) {
		v1334 = int32(-1)
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v228 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+14)) = uint8(v228)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v228
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)) = uint8(v228)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(-1)
	if v219 == v228 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v248 = F_hashtableNext(m, v18, v18+int32(312))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L20
	} else {
		goto L86
	}
L83:
	;
	goto L82
L84:
	;
	goto L83
L86:
	;
	if v248 == int32(0) {
		v1204 = v224
		goto L8
	} else {
		goto L87
	}
L87:
	;
	v255 = v224
	goto L88
L88:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v18)+312))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+int32(-1)))))
	switch v271 & int32(7) {
	case 0:
		goto L95
	case 1:
		goto L94
	case 2:
		goto L93
	case 3:
		goto L92
	case 4:
		goto L91
	default:
		v288 = int32(0)
		goto L90
	}
L90:
	;
	v289 = F_rdbSaveRawString(m, l0, v268, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L20
	} else {
		goto L96
	}
L91:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v268+int32(-17))))
	v288 = v287
	goto L90
L92:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v268+int32(-9))))
	v288 = v284
	goto L90
L93:
	;
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v268+int32(-5)))))
	v288 = v281
	goto L90
L94:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268+int32(-3)))))
	v288 = v278
	goto L90
L95:
	;
	v288 = int32(base.Ui32(v271) >> (uint(int32(3)) % 32))
	goto L90
L96:
	;
	if v289 == int32(-1) {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	v293 = v289 + v255
	v296 = F_hashtableNext(m, v18, v18+int32(312))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L20
	} else {
		goto L98
	}
L98:
	;
	if v296 == int32(0) {
		v1204 = v293
		goto L8
	} else {
		goto L99
	}
L99:
	;
	v255 = v293
	goto L88
L100:
	;
	v307 = F_objectGetVal(m, l1)
	mBase = m.M
	v308 = F_rdbSaveRawString(m, l0, v307, v302*v303+int32(8))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L20
	} else {
		goto L101
	}
L101:
	;
	if v308 != int32(-1) {
		v1221 = v308
		goto L7
	} else {
		goto L102
	}
L102:
	;
	v1334 = int32(-1)
	goto L1
L103:
	;
	v315 = F_objectGetVal(m, l1)
	mBase = m.M
	v316 = F_rdbSaveRawString(m, l0, v315, v314)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L20
	} else {
		goto L104
	}
L104:
	;
	if v316 != int32(-1) {
		v1221 = v316
		goto L7
	} else {
		goto L105
	}
L105:
	;
	v1334 = int32(-1)
	goto L1
L106:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	F__serverPanic_1(m, int32(_a875), int32(977), int32(_a854), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L20
	} else {
		goto L148
	}
L108:
	;
	v342 = F_objectGetVal(m, l1)
	mBase = m.M
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+4))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	goto L113
L109:
	;
	v334 = F_objectGetVal(m, l1)
	mBase = m.M
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	goto L110
L110:
	;
	v336 = F_objectGetVal(m, l1)
	mBase = m.M
	v337 = F_rdbSaveRawString(m, l0, v336, v335)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L20
	} else {
		goto L111
	}
L111:
	;
	if v337 != int32(-1) {
		v1221 = v337
		goto L7
	} else {
		goto L112
	}
L112:
	;
	v1334 = int32(-1)
	goto L1
L113:
	;
	v346 = F_rdbSaveLen(m, l0, base.I64_extend_i32_u(v344))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L20
	} else {
		goto L114
	}
L114:
	;
	if v346 == int32(-1) {
		v1334 = int32(-1)
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v343)+8))
	goto L116
L116:
	;
	if v350 == int32(0) {
		v1221 = v346
		goto L7
	} else {
		goto L117
	}
L117:
	;
	v356 = v346
	v361 = v350
	goto L118
L118:
	;
	v370 = v361 + int32(16)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	v374 = v370 + v371<<(uint(int32(3))%32)
	v375 = int32(*(*int8)(unsafe.Add(mBase, uint32(v374))))
	v376 = v374 + v375
	goto L126
L120:
	;
	v399 = int32(-1)
	v400 = F_rdbSaveRawString(m, l0, v376+int32(1), v398)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L20
	} else {
		goto L127
	}
L121:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v376+int32(-16))))
	v398 = v397
	goto L120
L122:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v376+int32(-8))))
	v398 = v394
	goto L120
L123:
	;
	v391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v376+int32(-4)))))
	v398 = v391
	goto L120
L124:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376+int32(-2)))))
	v398 = v388
	goto L120
L125:
	;
	v398 = int32(base.Ui32(v381) >> (uint(int32(3)) % 32))
	goto L120
L126:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376+int32(0)))))
	switch v381 & int32(7) {
	case 0:
		goto L125
	case 1:
		goto L124
	case 2:
		goto L123
	case 3:
		goto L122
	case 4:
		goto L121
	default:
		v398 = int32(0)
		goto L120
	}
L127:
	;
	if v400 == int32(-1) {
		v1334 = v399
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v404 = *(*float64)(unsafe.Add(mBase, uint32(v361)))
	*(*float64)(unsafe.Add(mBase, uint32(v18))) = v404
	if l0 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v465 = v356 + v400 + int32(8)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v361)+8))
	if v466 != 0 {
		v356 = v465
		v361 = v466
		goto L118
	} else {
		goto L147
	}
L130:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v408&int32(6) != 0 {
		v1334 = v399
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v416 = v18
	v418 = int32(8)
	goto L132
L132:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v427) < base.Ui32(v418) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	goto L129
L134:
	;
	v429 = v427
	goto L136
L135:
	;
	v429 = v418
	goto L136
L136:
	;
	if v427 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v430 = v429
	goto L139
L138:
	;
	v430 = v418
	goto L139
L139:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v431 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v437 = m.T0[v436].(func(*base.Module, int32, int32, int32) int32)(m, l0, v416, v430)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L20
	} else {
		goto L144
	}
L141:
	;
	m.T0[v431].(func(*base.Module, int32, int32, int32))(m, l0, v416, v430)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L20
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v443 + v430
	v447 = v418 - v430
	if v447 != 0 {
		v416 = v416 + v430
		v418 = v447
		goto L132
	} else {
		goto L146
	}
L144:
	;
	if v437 != 0 {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v439 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v439 | int64(2)
	v1334 = v399
	goto L1
L146:
	;
	goto L133
L147:
	;
	v1221 = v465
	goto L7
L148:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	F__serverPanic_1(m, int32(_a875), int32(1026), int32(_a855), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L20
	} else {
		goto L252
	}
L150:
	;
	switch l4 + int32(-4) {
	case 0, 18:
		goto L155
	default:
		goto L156
	}
L151:
	;
	v481 = F_objectGetVal(m, l1)
	mBase = m.M
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)))
	goto L152
L152:
	;
	v483 = F_objectGetVal(m, l1)
	mBase = m.M
	v484 = F_rdbSaveRawString(m, l0, v483, v482)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L20
	} else {
		goto L153
	}
L153:
	;
	if v484 != int32(-1) {
		v1221 = v484
		goto L7
	} else {
		goto L154
	}
L154:
	;
	v1334 = int32(-1)
	goto L1
L155:
	;
	v497 = F_objectGetVal(m, l1)
	mBase = m.M
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+20))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v497)+16))
	goto L158
L156:
	;
	F__serverAssert(m, int32(_a876), int32(_a875), int32(987))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L20
	} else {
		goto L157
	}
L157:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	v502 = F_rdbSaveLen(m, l0, base.I64_extend_i32_u(v498+v499))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L20
	} else {
		goto L159
	}
L159:
	;
	if v502 == int32(-1) {
		v1334 = int32(-1)
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v506 = int32(4)
	v507 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+14)) = uint8(v507)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v507
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)) = uint8(v506)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(-1)
	if v497 == v507 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v526 = F_hashtableNext(m, v18, v18+int32(308))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L20
	} else {
		goto L167
	}
L162:
	;
	goto L161
L163:
	;
	goto L162
L165:
	;
	v833 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v833 | int64(2)
	goto L4
L166:
	;
	F_hashtableCleanupIterator(m, v18)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L20
	} else {
		goto L251
	}
L167:
	;
	if v526 == int32(0) {
		v819 = v502
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v535 = v502
	goto L169
L169:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v18)+308))
	goto L171
L170:
	;
	v819 = v800
	goto L166
L171:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v18)+308))
	v550 = v18 + int32(304)
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548+int32(-1)))))
	v558 = v556 & int32(7)
	if v558 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547+int32(-1)))))
	switch v664 & int32(7) {
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
		v681 = int32(0)
		goto L204
	}
L173:
	;
	v660 = v652
	goto L172
L174:
	;
	v613 = F_sdsAllocPtr(m, v548)
	mBase = m.M
	v615 = v613 + int32(-4)
	if v556&int32(32) == int32(0) {
		goto L191
	} else {
		goto L192
	}
L175:
	;
	switch v558 {
	case 0:
		goto L183
	case 1:
		goto L182
	case 2:
		goto L181
	case 3:
		goto L180
	case 4:
		goto L179
	default:
		v578 = int32(0)
		goto L178
	}
L176:
	;
	if v556&int32(16) != 0 {
		goto L174
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	v580 = int32(1)
	v581 = F_sdsHdrSize(m, v580)
	mBase = m.M
	v582 = v548 + v578 + v581
	v584 = v582 + v580
	if v550 == int32(0) {
		v652 = v584
		goto L173
	} else {
		goto L184
	}
L179:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v548+int32(-17))))
	v578 = v577
	goto L178
L180:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v548+int32(-9))))
	v578 = v574
	goto L178
L181:
	;
	v571 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v548+int32(-5)))))
	v578 = v571
	goto L178
L182:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548+int32(-3)))))
	v578 = v568
	goto L178
L183:
	;
	v578 = int32(base.Ui32(v556) >> (uint(int32(3)) % 32))
	goto L178
L184:
	;
	v587 = int32(0)
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582+v587))))
	switch v590 & int32(7) {
	case 0:
		goto L190
	case 1:
		goto L189
	case 2:
		goto L188
	case 3:
		goto L187
	case 4:
		goto L186
	default:
		v611 = v587
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = v611
	v660 = v584
	goto L172
L186:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v582+int32(-16))))
	v611 = v610
	goto L185
L187:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v582+int32(-8))))
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = v606
	v660 = v584
	goto L172
L188:
	;
	v602 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v582+int32(-4)))))
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = v602
	v660 = v584
	goto L172
L189:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582+int32(-2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = v598
	v660 = v584
	goto L172
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = int32(base.Ui32(v590) >> (uint(int32(3)) % 32))
	v660 = v584
	goto L172
L191:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v615)))
	if v550 == int32(0) {
		v652 = v627
		goto L173
	} else {
		goto L197
	}
L192:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v615)))
	if v620 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	if v550 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v660 = int32(0)
	goto L172
L195:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v660 = v626
	goto L172
L196:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v620)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = v624
	goto L195
L197:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627+int32(-1)))))
	switch v633 & int32(7) {
	case 0:
		goto L203
	case 1:
		goto L202
	case 2:
		goto L201
	case 3:
		goto L200
	case 4:
		goto L199
	default:
		v650 = int32(0)
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = v650
	v652 = v627
	goto L173
L199:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v627+int32(-17))))
	v650 = v649
	goto L198
L200:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v627+int32(-9))))
	v650 = v646
	goto L198
L201:
	;
	v643 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v627+int32(-5)))))
	v650 = v643
	goto L198
L202:
	;
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627+int32(-3)))))
	v650 = v640
	goto L198
L203:
	;
	v650 = int32(base.Ui32(v633) >> (uint(int32(3)) % 32))
	goto L198
L204:
	;
	v682 = F_rdbSaveRawString(m, l0, v547, v681)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L20
	} else {
		goto L210
	}
L205:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v547+int32(-17))))
	v681 = v680
	goto L204
L206:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v547+int32(-9))))
	v681 = v677
	goto L204
L207:
	;
	v674 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v547+int32(-5)))))
	v681 = v674
	goto L204
L208:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547+int32(-3)))))
	v681 = v671
	goto L204
L209:
	;
	v681 = int32(base.Ui32(v664) >> (uint(int32(3)) % 32))
	goto L204
L210:
	;
	if v682 == int32(-1) {
		goto L4
	} else {
		goto L211
	}
L211:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v18)+304))
	v687 = F_rdbSaveRawString(m, l0, v660, v686)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L20
	} else {
		goto L212
	}
L212:
	;
	if v687 == int32(-1) {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	v692 = v682 + v535 + v687
	if l4 != int32(22) {
		v800 = v692
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v814 = F_hashtableNext(m, v18, v18+int32(308))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L20
	} else {
		goto L249
	}
L215:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v18)+308))
	v697 = int64(-1)
	v699 = v693 + int32(-1)
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699))))
	if v700&int32(7) == int32(0) {
		v735 = v697
		goto L217
	} else {
		goto L218
	}
L216:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+312)) = v735
	if l0 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L217:
	;
	goto L216
L218:
	;
	if v700&int32(8) == int32(0) {
		v735 = v697
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v709 = F_sdsAllocPtr(m, v693)
	mBase = m.M
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699))))
	if int32(base.Ui32(v712&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v717 = int32(-4)
	goto L222
L221:
	;
	v717 = int32(0)
	goto L222
L222:
	;
	v720 = v712 & int32(7)
	if v720 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v721 = v717
	goto L225
L224:
	;
	v721 = int32(0)
	goto L225
L225:
	;
	if int32(base.Ui32(v712&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v729 = int32(-8)
	goto L228
L227:
	;
	v729 = int32(0)
	goto L228
L228:
	;
	if v720 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v731 = v729
	goto L231
L230:
	;
	v731 = int32(0)
	goto L231
L231:
	;
	v733 = *(*int64)(unsafe.Add(mBase, uint32(v709+v721+v731)))
	v735 = v733
	goto L217
L232:
	;
	v800 = v692 + int32(8)
	goto L214
L233:
	;
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v740&int32(6) != 0 {
		goto L4
	} else {
		goto L234
	}
L234:
	;
	v752 = int32(8)
	v753 = v18 + int32(312)
	goto L235
L235:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v761) < base.Ui32(v752) {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	goto L232
L237:
	;
	v763 = v761
	goto L239
L238:
	;
	v763 = v752
	goto L239
L239:
	;
	if v761 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v764 = v763
	goto L242
L241:
	;
	v764 = v752
	goto L242
L242:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v765 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v771 = m.T0[v770].(func(*base.Module, int32, int32, int32) int32)(m, l0, v753, v764)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L20
	} else {
		goto L246
	}
L244:
	;
	m.T0[v765].(func(*base.Module, int32, int32, int32))(m, l0, v753, v764)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L20
	} else {
		goto L245
	}
L245:
	;
	goto L243
L246:
	;
	if v771 == int32(0) {
		goto L165
	} else {
		goto L247
	}
L247:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v775 + v764
	v779 = v752 - v764
	if v779 != 0 {
		v752 = v779
		v753 = v753 + v764
		goto L235
	} else {
		goto L248
	}
L248:
	;
	goto L236
L249:
	;
	if v814 != 0 {
		v535 = v800
		goto L169
	} else {
		goto L250
	}
L250:
	;
	goto L170
L251:
	;
	v1221 = v819
	goto L7
L252:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	v848 = F_rdbSaveLen(m, l0, v847)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L20
	} else {
		goto L254
	}
L254:
	;
	if v848 == int32(-1) {
		v1334 = v844
		goto L1
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(128)
	v857 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+12)) = v857
	*(*int64)(unsafe.Add(mBase, uint32(v18)+296)) = v857
	*(*int64)(unsafe.Add(mBase, uint32(v18)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v18 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+156)) = v18 + int32(168)
	goto L256
L256:
	;
	v870 = int32(0)
	v872 = F_raxSeek(m, v18, int32(_a4), v870, v870)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L20
	} else {
		goto L257
	}
L257:
	;
	v874 = F_raxNext(m, v18)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L20
	} else {
		goto L260
	}
L258:
	;
	F_raxStop(m, v18)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L20
	} else {
		goto L317
	}
L259:
	;
	F_raxStop(m, v18)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L20
	} else {
		goto L271
	}
L260:
	;
	if v874 == int32(0) {
		v910 = v848
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v879 = v848
	goto L262
L262:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v893)))
	goto L264
L263:
	;
	v910 = v906
	goto L259
L264:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v897 = F_rdbSaveRawString(m, l0, v895, v896)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L20
	} else {
		goto L265
	}
L265:
	;
	if v897 == int32(-1) {
		goto L258
	} else {
		goto L266
	}
L266:
	;
	v901 = F_rdbSaveRawString(m, l0, v893, v894)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L20
	} else {
		goto L267
	}
L267:
	;
	if v901 == int32(-1) {
		goto L258
	} else {
		goto L268
	}
L268:
	;
	v906 = v897 + v879 + v901
	v907 = F_raxNext(m, v18)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L20
	} else {
		goto L269
	}
L269:
	;
	if v907 != 0 {
		v879 = v906
		goto L262
	} else {
		goto L270
	}
L270:
	;
	goto L263
L271:
	;
	v926 = *(*int64)(unsafe.Add(mBase, uint32(v845)+8))
	v927 = F_rdbSaveLen(m, l0, v926)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L20
	} else {
		goto L272
	}
L272:
	;
	if v927 == int32(-1) {
		v1334 = v844
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v931 = *(*int64)(unsafe.Add(mBase, uint32(v845)+16))
	v932 = F_rdbSaveLen(m, l0, v931)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L20
	} else {
		goto L274
	}
L274:
	;
	if v932 == int32(-1) {
		v1334 = v844
		goto L1
	} else {
		goto L275
	}
L275:
	;
	v936 = *(*int64)(unsafe.Add(mBase, uint32(v845)+24))
	v937 = F_rdbSaveLen(m, l0, v936)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L20
	} else {
		goto L276
	}
L276:
	;
	if v937 == int32(-1) {
		v1334 = v844
		goto L1
	} else {
		goto L277
	}
L277:
	;
	v941 = *(*int64)(unsafe.Add(mBase, uint32(v845)+32))
	v942 = F_rdbSaveLen(m, l0, v941)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L20
	} else {
		goto L278
	}
L278:
	;
	if v942 == int32(-1) {
		v1334 = v844
		goto L1
	} else {
		goto L279
	}
L279:
	;
	v946 = *(*int64)(unsafe.Add(mBase, uint32(v845)+40))
	v947 = F_rdbSaveLen(m, l0, v946)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L20
	} else {
		goto L280
	}
L280:
	;
	if v947 == int32(-1) {
		v1334 = v844
		goto L1
	} else {
		goto L281
	}
L281:
	;
	v951 = *(*int64)(unsafe.Add(mBase, uint32(v845)+48))
	v952 = F_rdbSaveLen(m, l0, v951)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L20
	} else {
		goto L282
	}
L282:
	;
	if v952 == int32(-1) {
		v1334 = v844
		goto L1
	} else {
		goto L283
	}
L283:
	;
	v956 = *(*int64)(unsafe.Add(mBase, uint32(v845)+56))
	v957 = F_rdbSaveLen(m, l0, v956)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L20
	} else {
		goto L284
	}
L284:
	;
	if v957 == int32(-1) {
		v1334 = v844
		goto L1
	} else {
		goto L285
	}
L285:
	;
	v961 = *(*int64)(unsafe.Add(mBase, uint32(v845)+64))
	v962 = F_rdbSaveLen(m, l0, v961)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L20
	} else {
		goto L286
	}
L286:
	;
	if v962 == int32(-1) {
		v1334 = v844
		goto L1
	} else {
		goto L287
	}
L287:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v845)))
	if v966 != 0 {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	v972 = F_rdbSaveLen(m, l0, base.I64_extend_i32_u(v970))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L20
	} else {
		goto L292
	}
L289:
	;
	v968 = *(*int64)(unsafe.Add(mBase, uint32(v966)+8))
	goto L291
L290:
	;
	v970 = int32(0)
	goto L288
L291:
	;
	v970 = base.I32_wrap_i64(v968)
	goto L288
L292:
	;
	if v972 == int32(-1) {
		v1334 = v844
		goto L1
	} else {
		goto L293
	}
L293:
	;
	v984 = v927 + v910 + v932 + v937 + v942 + v947 + v952 + v957 + v962 + v972
	if v970 == int32(0) {
		v1221 = v984
		goto L7
	} else {
		goto L294
	}
L294:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v845)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v987
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(128)
	v993 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+12)) = v993
	*(*int64)(unsafe.Add(mBase, uint32(v18)+296)) = v993
	*(*int64)(unsafe.Add(mBase, uint32(v18)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v18 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+156)) = v18 + int32(168)
	goto L295
L295:
	;
	v1006 = int32(0)
	v1008 = F_raxSeek(m, v18, int32(_a4), v1006, v1006)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L20
	} else {
		goto L296
	}
L296:
	;
	v1010 = F_raxNext(m, v18)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L20
	} else {
		goto L298
	}
L297:
	;
	F_raxStop(m, v18)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L20
	} else {
		goto L316
	}
L298:
	;
	if v1010 == int32(0) {
		v1072 = v984
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1017 = v984
	goto L300
L300:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v1032 = F_rdbSaveRawString(m, l0, v1030, v1031)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L20
	} else {
		goto L302
	}
L301:
	;
	v1072 = v1066
	goto L297
L302:
	;
	if v1032 == int32(-1) {
		goto L258
	} else {
		goto L303
	}
L303:
	;
	v1036 = *(*int64)(unsafe.Add(mBase, uint32(v1029)))
	v1037 = F_rdbSaveLen(m, l0, v1036)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L20
	} else {
		goto L304
	}
L304:
	;
	if v1037 == int32(-1) {
		goto L258
	} else {
		goto L305
	}
L305:
	;
	v1041 = *(*int64)(unsafe.Add(mBase, uint32(v1029)+8))
	v1042 = F_rdbSaveLen(m, l0, v1041)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L20
	} else {
		goto L306
	}
L306:
	;
	if v1042 == int32(-1) {
		goto L258
	} else {
		goto L307
	}
L307:
	;
	v1046 = *(*int64)(unsafe.Add(mBase, uint32(v1029)+16))
	v1047 = F_rdbSaveLen(m, l0, v1046)
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L20
	} else {
		goto L308
	}
L308:
	;
	if v1047 == int32(-1) {
		goto L258
	} else {
		goto L309
	}
L309:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+24))
	v1053 = F_rdbSaveStreamPEL(m, l0, v1051, int32(1))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L20
	} else {
		goto L310
	}
L310:
	;
	if v1053 == int32(-1) {
		goto L258
	} else {
		goto L311
	}
L311:
	;
	v1057 = F_rdbSaveStreamConsumers(m, l0, v1029)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L20
	} else {
		goto L312
	}
L312:
	;
	if v1057 == int32(-1) {
		goto L258
	} else {
		goto L313
	}
L313:
	;
	v1066 = v1032 + v1017 + v1037 + v1042 + v1047 + v1053 + v1057
	v1067 = F_raxNext(m, v18)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L20
	} else {
		goto L314
	}
L314:
	;
	if v1067 != 0 {
		v1017 = v1066
		goto L300
	} else {
		goto L315
	}
L315:
	;
	goto L301
L316:
	;
	v1221 = v1072
	goto L7
L317:
	;
	v1334 = v844
	goto L1
L318:
	;
	if v1107 == int32(-1) {
		v1334 = int32(-1)
		goto L1
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v1105
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+4))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+16))
	m.T0[v1121].(func(*base.Module, int32, int32))(m, v18, v1120)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L20
	} else {
		goto L320
	}
L320:
	;
	v1124 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+312)) = uint8(v1124)
	if l0 == v1124 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1185 + int32(1)
	goto L5
L322:
	;
	v1128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v1128&int32(6) != 0 {
		goto L6
	} else {
		goto L323
	}
L323:
	;
	v1140 = int32(1)
	v1141 = v18 + int32(312)
	goto L324
L324:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v1149) < base.Ui32(v1140) {
		goto L326
	} else {
		goto L327
	}
L325:
	;
	goto L321
L326:
	;
	v1151 = v1149
	goto L328
L327:
	;
	v1151 = v1140
	goto L328
L328:
	;
	if v1149 != 0 {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1152 = v1151
	goto L331
L330:
	;
	v1152 = v1140
	goto L331
L331:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v1153 == int32(0) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1159 = m.T0[v1158].(func(*base.Module, int32, int32, int32) int32)(m, l0, v1141, v1152)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L20
	} else {
		goto L336
	}
L333:
	;
	m.T0[v1153].(func(*base.Module, int32, int32, int32))(m, l0, v1141, v1152)
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L20
	} else {
		goto L334
	}
L334:
	;
	goto L332
L335:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1165 + v1152
	v1169 = v1140 - v1152
	if v1169 != 0 {
		v1140 = v1169
		v1141 = v1141 + v1152
		goto L324
	} else {
		goto L338
	}
L336:
	;
	if v1159 != 0 {
		goto L335
	} else {
		goto L337
	}
L337:
	;
	v1161 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v1161 | int64(2)
	goto L6
L338:
	;
	goto L325
L339:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L340:
	;
	if v1197 == int32(-1) {
		v1334 = int32(-1)
		goto L1
	} else {
		goto L341
	}
L341:
	;
	v1221 = v1197
	goto L7
L342:
	;
	v1221 = v1204
	goto L7
L343:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v1275 != 0 {
		goto L347
	} else {
		goto L348
	}
L344:
	;
	F_moduleFreeContext(m, v1265)
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L20
	} else {
		goto L345
	}
L345:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	F_valkey_free(m, v1270)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L20
	} else {
		goto L346
	}
L346:
	;
	goto L343
L347:
	;
	v1276 = int32(-1)
	goto L349
L348:
	;
	v1276 = v1274
	goto L349
L349:
	;
	v1334 = v1276
	goto L1
L350:
	;
	v1334 = int32(-1)
	goto L1
}
func F_rdbSaveSingleModuleAux(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(128)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+100)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(v12)+104)) = int64(0)
	v26 = F_sdsempty(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = F___memcpy(m, v12+int32(8), int32(_a157), int32(80))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v32)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v26
	goto L3
L3:
	;
	v36 = int32(247)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+126)) = uint8(v36)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+40)))
	if v38&int32(6) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v12 + int32(128)
	return v366
L5:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	F_sdsfree(m, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L114
	}
L6:
	;
	v48 = v12 + int32(126)
	v49 = int32(1)
	goto L7
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	if base.Ui32(v53) < base.Ui32(v49) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v80 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v81 = F_rdbSaveLen(m, v12+int32(8), v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L22
	}
L9:
	;
	v55 = v53
	goto L11
L10:
	;
	v55 = v49
	goto L11
L11:
	;
	if v53 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v56 = v55
	goto L14
L13:
	;
	v56 = v49
	goto L14
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if v57 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v67 = m.T0[v66].(func(*base.Module, int32, int32, int32) int32)(m, v12+int32(8), v48, v56)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	m.T0[v57].(func(*base.Module, int32, int32, int32))(m, v12+int32(8), v48, v56)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v73 + v56
	v77 = v49 - v56
	if v77 != 0 {
		v48 = v48 + v56
		v49 = v77
		goto L7
	} else {
		goto L21
	}
L19:
	;
	if v67 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v69 | int64(2)
	goto L5
L21:
	;
	goto L8
L22:
	;
	if v81 == int32(-1) {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v85 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+126)) = uint8(v85)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+40)))
	if v87&int32(6) != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v97 = v12 + int32(126)
	v98 = int32(1)
	goto L25
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	if base.Ui32(v102) < base.Ui32(v98) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v130 = F_rdbSaveLen(m, v12+int32(8), base.I64_extend_i32_s(l1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L40
	}
L27:
	;
	v104 = v102
	goto L29
L28:
	;
	v104 = v98
	goto L29
L29:
	;
	if v102 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v105 = v104
	goto L32
L31:
	;
	v105 = v98
	goto L32
L32:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if v106 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v116 = m.T0[v115].(func(*base.Module, int32, int32, int32) int32)(m, v12+int32(8), v97, v105)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L37
	}
L34:
	;
	m.T0[v106].(func(*base.Module, int32, int32, int32))(m, v12+int32(8), v97, v105)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v122 + v105
	v126 = v98 - v105
	if v126 != 0 {
		v97 = v97 + v105
		v98 = v126
		goto L25
	} else {
		goto L39
	}
L37:
	;
	if v116 != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v118 | int64(2)
	goto L5
L39:
	;
	goto L26
L40:
	;
	if v130 == int32(-1) {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	if v135 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v256 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+126)) = uint8(v256)
	if l0 == v256 {
		goto L82
	} else {
		goto L83
	}
L43:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+int32(-1)))))
	switch v162 & int32(7) {
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
		v179 = int32(0)
		goto L51
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = v134
	m.T0[v135].(func(*base.Module, int32, int32))(m, v12+int32(92), l1)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
	if v143 == int32(0) {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	F_sdsfree(m, v143)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v148 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = v148
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	if v151 == v148 {
		v366 = v148
		goto L4
	} else {
		goto L48
	}
L48:
	;
	F_moduleFreeContext(m, v151)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	F_valkey_free(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v366 = v148
	goto L4
L51:
	;
	if l0 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v134+int32(-17))))
	v179 = v178
	goto L51
L53:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v134+int32(-9))))
	v179 = v175
	goto L51
L54:
	;
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134+int32(-5)))))
	v179 = v172
	goto L51
L55:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+int32(-3)))))
	v179 = v169
	goto L51
L56:
	;
	v179 = int32(base.Ui32(v162) >> (uint(int32(3)) % 32))
	goto L51
L57:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = v237 + v236
	F_sdsfree(m, v234)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L79
	}
L58:
	;
	if v179 == int32(-1) {
		goto L5
	} else {
		goto L78
	}
L59:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v182&int32(6) != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	if v179 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v190 = v179
	v192 = v134
	goto L63
L62:
	;
	v234 = v134
	v236 = int32(0)
	goto L57
L63:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v195) < base.Ui32(v190) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L58
L65:
	;
	v197 = v195
	goto L67
L66:
	;
	v197 = v190
	goto L67
L67:
	;
	if v195 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v198 = v197
	goto L70
L69:
	;
	v198 = v190
	goto L70
L70:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v199 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v205 = m.T0[v204].(func(*base.Module, int32, int32, int32) int32)(m, l0, v192, v198)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L75
	}
L72:
	;
	m.T0[v199].(func(*base.Module, int32, int32, int32))(m, l0, v192, v198)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v211 + v198
	v215 = v190 - v198
	if v215 != 0 {
		v190 = v215
		v192 = v192 + v198
		goto L63
	} else {
		goto L77
	}
L75:
	;
	if v205 != 0 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v207 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v207 | int64(2)
	goto L5
L77:
	;
	goto L64
L78:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v234 = v227
	v236 = v179
	goto L57
L79:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	m.T0[v244].(func(*base.Module, int32, int32))(m, v12+int32(92), l1)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	goto L42
L81:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
	if v321 != 0 {
		goto L101
	} else {
		goto L102
	}
L82:
	;
	v313 = int32(1)
	v314 = v256
	goto L81
L83:
	;
	v261 = int32(-1)
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v262&int32(6) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v275 = v12 + int32(126)
	v276 = int32(1)
	goto L86
L85:
	;
	v313 = v261
	v314 = int32(1)
	goto L81
L86:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v280) < base.Ui32(v276) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L82
L88:
	;
	v282 = v280
	goto L90
L89:
	;
	v282 = v276
	goto L90
L90:
	;
	if v280 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v283 = v282
	goto L93
L92:
	;
	v283 = v276
	goto L93
L93:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v284 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v290 = m.T0[v289].(func(*base.Module, int32, int32, int32) int32)(m, l0, v275, v283)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L98
	}
L95:
	;
	m.T0[v284].(func(*base.Module, int32, int32, int32))(m, l0, v275, v283)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v297 + v283
	v301 = v276 - v283
	if v301 != 0 {
		v275 = v275 + v283
		v276 = v301
		goto L86
	} else {
		goto L100
	}
L98:
	;
	if v290 != 0 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v292 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v292 | int64(2)
	v313 = v261
	v314 = int32(1)
	goto L81
L100:
	;
	goto L87
L101:
	;
	F__serverAssert(m, int32(_a879), int32(_a875), int32(1347))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L113
	}
L102:
	;
	if v314 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	if v329 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = v326 + v313
	goto L103
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = int32(1)
	goto L103
L106:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
	if v339 != 0 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	F_moduleFreeContext(m, v329)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	F_valkey_free(m, v334)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L106
L110:
	;
	v340 = int32(-1)
	goto L112
L111:
	;
	v340 = v338
	goto L112
L112:
	;
	v366 = v340
	goto L4
L113:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	v366 = int32(-1)
	goto L4
}
func F_rdbSaveStreamConsumers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int64
	_ = v32
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
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
	var v73 int64
	_ = v73
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
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
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	v11 = m.G0
	v13 = v11 - int32(320)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	goto L2
L1:
	;
	m.G0 = v13 + int32(320)
	return v195
L2:
	;
	v18 = F_rdbSaveLen(m, l0, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v18 == int32(-1) {
		v195 = int32(-1)
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v25 = v13 + int32(8)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = int32(128)
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+12)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v25)+296)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v25)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v13 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+156)) = v13 + int32(176)
	goto L6
L6:
	;
	v47 = int32(0)
	v49 = F_raxSeek(m, v13+int32(8), int32(_a4), v47, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v53 = F_raxNext(m, v13+int32(8))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	F_raxStop(m, v13+int32(8))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L3
	} else {
		goto L54
	}
L9:
	;
	v59 = v18
	goto L12
L10:
	;
	if v53 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v181 = v18
	goto L8
L12:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v66 = int32(-1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v69 = F_rdbSaveRawString(m, l0, v67, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L3
	} else {
		goto L14
	}
L13:
	;
	v181 = v173
	goto L8
L14:
	;
	if v69 == int32(-1) {
		v181 = v66
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+312)) = v73
	if l0 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v166 = F_rdbSaveStreamPEL(m, l0, v164, int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L3
	} else {
		goto L50
	}
L17:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v77&int32(6) != 0 {
		v181 = v66
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v90 = v13 + int32(312)
	v91 = int32(8)
	goto L19
L19:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v93) < base.Ui32(v91) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+312)) = v114
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v116&int32(6) != 0 {
		v181 = v66
		goto L8
	} else {
		goto L34
	}
L21:
	;
	v95 = v93
	goto L23
L22:
	;
	v95 = v91
	goto L23
L23:
	;
	if v93 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v96 = v95
	goto L26
L25:
	;
	v96 = v91
	goto L26
L26:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v97 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v103 = m.T0[v102].(func(*base.Module, int32, int32, int32) int32)(m, l0, v90, v96)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L3
	} else {
		goto L31
	}
L28:
	;
	m.T0[v97].(func(*base.Module, int32, int32, int32))(m, l0, v90, v96)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v109 + v96
	v113 = v91 - v96
	if v113 != 0 {
		v90 = v90 + v96
		v91 = v113
		goto L19
	} else {
		goto L33
	}
L31:
	;
	if v103 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v105 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v105 | int64(2)
	v181 = v66
	goto L8
L33:
	;
	goto L20
L34:
	;
	v129 = v13 + int32(312)
	v130 = int32(8)
	goto L35
L35:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v132) < base.Ui32(v130) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L16
L37:
	;
	v134 = v132
	goto L39
L38:
	;
	v134 = v130
	goto L39
L39:
	;
	if v132 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v135 = v134
	goto L42
L41:
	;
	v135 = v130
	goto L42
L42:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v136 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v142 = m.T0[v141].(func(*base.Module, int32, int32, int32) int32)(m, l0, v129, v135)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
	} else {
		goto L47
	}
L44:
	;
	m.T0[v136].(func(*base.Module, int32, int32, int32))(m, l0, v129, v135)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v148 + v135
	v152 = v130 - v135
	if v152 != 0 {
		v129 = v129 + v135
		v130 = v152
		goto L35
	} else {
		goto L49
	}
L47:
	;
	if v142 != 0 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v144 | int64(2)
	v181 = v66
	goto L8
L49:
	;
	goto L36
L50:
	;
	if v166 == int32(-1) {
		v181 = int32(-1)
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v173 = v59 + v69 + v166 + int32(16)
	v176 = F_raxNext(m, v13+int32(8))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	if v176 != 0 {
		v59 = v173
		goto L12
	} else {
		goto L53
	}
L53:
	;
	goto L13
L54:
	;
	v195 = v181
	goto L1
}
func F_rdbSaveStreamPEL(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int64
	_ = v30
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int64
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	v11 = m.G0
	v13 = v11 - int32(320)
	m.G0 = v13
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	goto L2
L1:
	;
	m.G0 = v13 + int32(320)
	return v212
L2:
	;
	v17 = F_rdbSaveLen(m, l0, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v17 == int32(-1) {
		v212 = int32(-1)
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v24 = v13 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = int32(128)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+12)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v24)+296)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v24)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v13 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+156)) = v13 + int32(176)
	goto L6
L6:
	;
	v45 = int32(0)
	v47 = F_raxSeek(m, v13+int32(8), int32(_a4), v45, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v51 = F_raxNext(m, v13+int32(8))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	F_raxStop(m, v13+int32(8))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L60
	}
L9:
	;
	v57 = v17
	goto L13
L10:
	;
	if v51 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v198 = v17
	goto L8
L12:
	;
	v198 = int32(-1)
	goto L8
L13:
	;
	if l0 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v181 = F_raxNext(m, v13+int32(8))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L3
	} else {
		goto L58
	}
L16:
	;
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v159)+8))
	v162 = F_rdbSaveLen(m, l0, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L3
	} else {
		goto L56
	}
L17:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v158 = v108
	v159 = v149
	goto L16
L18:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+312)) = v110
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v112&int32(6) != 0 {
		goto L12
	} else {
		goto L40
	}
L19:
	;
	v108 = v57 + int32(16)
	if l2 != 0 {
		goto L17
	} else {
		goto L39
	}
L20:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v65&int32(6) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v78 = int32(16)
	v79 = v72
	goto L23
L22:
	;
	v198 = int32(-1)
	goto L8
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v83) < base.Ui32(v78) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v106 = v57 + int32(16)
	if l2 != 0 {
		goto L18
	} else {
		goto L38
	}
L25:
	;
	v85 = v83
	goto L27
L26:
	;
	v85 = v78
	goto L27
L27:
	;
	if v83 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v86 = v85
	goto L30
L29:
	;
	v86 = v78
	goto L30
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v87 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v93 = m.T0[v92].(func(*base.Module, int32, int32, int32) int32)(m, l0, v79, v86)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L3
	} else {
		goto L35
	}
L32:
	;
	m.T0[v87].(func(*base.Module, int32, int32, int32))(m, l0, v79, v86)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v100 + v86
	v104 = v78 - v86
	if v104 != 0 {
		v78 = v104
		v79 = v79 + v86
		goto L23
	} else {
		goto L37
	}
L35:
	;
	if v93 != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v95 | int64(2)
	v198 = int32(-1)
	goto L8
L37:
	;
	goto L24
L38:
	;
	v173 = v106
	goto L15
L39:
	;
	v173 = v108
	goto L15
L40:
	;
	v123 = int32(8)
	v124 = v13 + int32(312)
	goto L41
L41:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v128) < base.Ui32(v123) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v130 = v128
	goto L45
L44:
	;
	v130 = v123
	goto L45
L45:
	;
	if v128 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v131 = v130
	goto L48
L47:
	;
	v131 = v123
	goto L48
L48:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v132 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v138 = m.T0[v137].(func(*base.Module, int32, int32, int32) int32)(m, l0, v124, v131)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L3
	} else {
		goto L53
	}
L50:
	;
	m.T0[v132].(func(*base.Module, int32, int32, int32))(m, l0, v124, v131)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + v131
	v148 = v123 - v131
	if v148 != 0 {
		v123 = v148
		v124 = v124 + v131
		goto L41
	} else {
		goto L55
	}
L53:
	;
	if v138 != 0 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v140 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v140 | int64(2)
	goto L12
L55:
	;
	v158 = v106
	v159 = v109
	goto L16
L56:
	;
	if v162 == int32(-1) {
		v198 = int32(-1)
		goto L8
	} else {
		goto L57
	}
L57:
	;
	v173 = v158 + v162 + int32(8)
	goto L15
L58:
	;
	if v181 != 0 {
		v57 = v173
		goto L13
	} else {
		goto L59
	}
L59:
	;
	v198 = v173
	goto L8
L60:
	;
	v212 = v198
	goto L1
}
func F_rdbSaveType(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	v2 = l1
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v2)
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v63
L2:
	;
	v63 = int32(1)
	goto L1
L3:
	;
	v15 = int32(-1)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v16&int32(6) != 0 {
		v63 = v15
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = v10 + int32(15)
	v27 = int32(1)
	goto L5
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v29) < base.Ui32(v27) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	v31 = v29
	goto L9
L8:
	;
	v31 = v27
	goto L9
L9:
	;
	if v29 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v32 = v31
	goto L12
L11:
	;
	v32 = v27
	goto L12
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v33 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v41 = m.T0[v40].(func(*base.Module, int32, int32, int32) int32)(m, l0, v26, v32)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L15
	} else {
		goto L18
	}
L14:
	;
	m.T0[v33].(func(*base.Module, int32, int32, int32))(m, l0, v26, v32)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	goto L13
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v47 + v32
	v51 = v27 - v32
	if v51 != 0 {
		v26 = v26 + v32
		v27 = v51
		goto L5
	} else {
		goto L20
	}
L18:
	;
	if v41 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v43 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v43 | int64(2)
	v63 = v15
	goto L1
L20:
	;
	goto L6
}
