package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_disableTracking(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int64
	_ = v23
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v326 int32
	_ = v326
	v4 = m.G0
	v6 = v4 - int32(320)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+316)) = l0
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v9&int32(16) == int32(0) {
		v302 = l0
		v304 = v9
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a1778), int32(_a2464), int32(77))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L5
	} else {
		goto L65
	}
L2:
	;
	if v304&int32(4) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L3:
	;
	v15 = v6 + int32(12)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(128)
	v23 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+12)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v15)+296)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v15)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v6 + int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+156)) = v6 + int32(180)
	goto L4
L4:
	;
	v38 = int32(0)
	v40 = F_raxSeek(m, v6+int32(12), int32(_a67), v38, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v44 = F_raxNext(m, v6+int32(12))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	F_raxStop(m, v6+int32(12))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L61
	}
L8:
	;
	if v44 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	goto L10
L10:
	;
	v51 = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, _consts[885]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
	v56 = v6 + int32(8)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v54 == v51 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L7
L12:
	;
	if v250 == int32(0) {
		goto L1
	} else {
		goto L50
	}
L13:
	;
	if v209 != v54 {
		v250 = v51
		goto L40
	} else {
		goto L41
	}
L14:
	;
	v200 = int32(0)
	v206 = v65
	v207 = v66
	v209 = v200
	v213 = v200
	goto L13
L15:
	;
	if base.Ui32(v66) < base.Ui32(int32(8)) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v77 = v65
	v78 = v66
	v80 = int32(0)
	goto L18
L17:
	;
	v206 = v190
	v207 = v191
	v209 = v193
	v213 = base.B2i32(v196 != int32(0))
	goto L13
L18:
	;
	v86 = int32(base.Ui32(v78) >> (uint(int32(3)) % 32))
	v87 = int32(4)
	v88 = v77 + v87
	if v78&v87 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v190 = v181
	v191 = v182
	v193 = v166
	v196 = v171
	goto L17
L20:
	;
	v171 = int32(0)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v88+v86+(v171-v86)&int32(3)+v159<<(uint(int32(2))%32))))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	if base.Ui32(v182) < base.Ui32(int32(8)) {
		v190 = v181
		v191 = v182
		v193 = v166
		v196 = v171
		goto L17
	} else {
		goto L38
	}
L21:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v80))))
	v137 = int32(0)
	goto L32
L22:
	;
	v93 = int32(0)
	if base.Ui32(v54) <= base.Ui32(v80) {
		v126 = v80
		v129 = v93
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v129 == v86 {
		v159 = v93
		v166 = v126
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v103 = v80
	v106 = v93
	goto L25
L25:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+v106))))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v103))))
	if v109 != v111 {
		v126 = v103
		v129 = v106
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v126 = v114
	v129 = v116
	goto L23
L27:
	;
	v113 = int32(1)
	v114 = v103 + v113
	v116 = v106 + v113
	if base.Ui32(v86) <= base.Ui32(v116) {
		v126 = v114
		v129 = v116
		goto L23
	} else {
		goto L28
	}
L28:
	;
	if base.Ui32(v114) < base.Ui32(v54) {
		v103 = v114
		v106 = v116
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	v190 = v77
	v191 = v78
	v193 = v126
	v196 = v129
	goto L17
L31:
	;
	if v137 != v86 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+v137))))
	if v150 == v134&int32(255) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v152 = int32(1)
	v154 = v137 + v152
	if v154 != v86 {
		v137 = v154
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v206 = v77
	v207 = v78
	v209 = v80
	v213 = v152
	goto L13
L36:
	;
	v159 = v137
	v166 = v80 + int32(1)
	goto L20
L37:
	;
	v190 = v77
	v191 = v78
	v193 = v80
	v196 = v86
	goto L17
L38:
	;
	if base.Ui32(v166) < base.Ui32(v54) {
		v77 = v181
		v78 = v182
		v80 = v166
		goto L18
	} else {
		goto L39
	}
L39:
	;
	goto L19
L40:
	;
	goto L12
L41:
	;
	v215 = int32(0)
	if v207&int32(1) == v215 {
		v250 = v215
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v221 = v207 & int32(4)
	if v213&base.B2i32(v221 != int32(0)) != 0 {
		v250 = v215
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v225 = int32(1)
	if v56 == int32(0) {
		v250 = v225
		goto L40
	} else {
		goto L44
	}
L44:
	;
	if v207&int32(2) != 0 {
		v247 = int32(0)
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v247
	v250 = v225
	goto L40
L46:
	;
	v231 = int32(3)
	v232 = int32(base.Ui32(v207) >> (uint(v231) % 32))
	if v221 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v242 = int32(4)
	goto L49
L48:
	;
	v242 = v232 << (uint(int32(2)) % 32)
	goto L49
L49:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v206+v232+(int32(0)-v232)&v231+v242+int32(4))))
	v247 = v246
	goto L45
L50:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	v260 = F_raxRemove(m, v255, v6+int32(316), int32(4), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v262)+8))
	goto L53
L52:
	;
	v283 = F_raxNext(m, v6+int32(12))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L5
	} else {
		goto L59
	}
L53:
	;
	if v263 != int64(0) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	F_raxFree(m, v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	F_raxFree(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_valkey_free(m, v254)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v274 = int32(0)
	v275 = *(*int32)(unsafe.Add(mBase, _consts[885]))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
	v279 = F_raxRemove(m, v275, v276, v277, v274)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	goto L52
L59:
	;
	if v283 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	goto L11
L61:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v6)+316))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+100))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+24))
	F_raxFree(m, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v6)+316))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+24)) = int32(0)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297)+204))
	v302 = v297
	v304 = v301
	goto L2
L63:
	;
	m.G0 = v6 + int32(320)
	return
L64:
	;
	v309 = int32(_a20)
	v311 = *(*int32)(unsafe.Add(mBase, _consts[780]))
	*(*int32)(unsafe.Add(mBase, _consts[780])) = v311 + int32(-1)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v302)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v302)+204)) = v315 & int32(-509)
	goto L63
L65:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sendTrackingMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
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
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v11 = v9 | int32(131072)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == int64(0) {
		v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
		if base.Ui32(int32(2)) < base.Ui32(v120) {
			v139 = l0
			v140 = v9
			F_addReplyPushLen(m, v139, int32(2))
			mBase = m.M
			v145 = m.ExcPending
			if v145 != 0 {
				return
			} else {
				F_addReplyBulkCBuffer(m, v139, int32(_a2465), int32(10))
				mBase = m.M
				v149 = m.ExcPending
				if v149 != 0 {
					return
				} else {
					v150 = v139
					v151 = v140
					if l3 == int32(0) {
						F_addReplyArrayLen(m, v150, int32(1))
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
							return
						} else {
							F_addReplyBulkCBuffer(m, v150, l1, l2)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return
							} else {
								v163 = F_updateClientMemUsageAndBucket(m, v150)
								mBase = m.M
								v164 = m.ExcPending
								if v164 != 0 {
									return
								} else {
									if v151&int32(131072) != 0 {
									} else {
										v167 = *(*int32)(unsafe.Add(mBase, uint32(v150)+204))
										*(*int32)(unsafe.Add(mBase, uint32(v150)+204)) = v167 & int32(-131073)
									}
									return
								}
							}
						}
					} else {
						F_addReplyProto(m, v150, l1, l2)
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return
						} else {
							v163 = F_updateClientMemUsageAndBucket(m, v150)
							mBase = m.M
							v164 = m.ExcPending
							if v164 != 0 {
								return
							} else {
								if v151&int32(131072) != 0 {
								} else {
									v167 = *(*int32)(unsafe.Add(mBase, uint32(v150)+204))
									*(*int32)(unsafe.Add(mBase, uint32(v150)+204)) = v167 & int32(-131073)
								}
								return
							}
						}
					}
				}
			}
		} else {
			v123 = l0
			v124 = v9
			v125 = v11
			if v124&int32(131072) != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v123)+204)) = v125 & int32(-131073)
				return
			}
		}
	} else {
		v19 = m.G0
		v20 = int32(16)
		v21 = v19 - v20
		m.G0 = v21
		v23 = int64(56)
		v25 = int64(65280)
		v27 = int64(40)
		v30 = int64(16711680)
		v32 = int64(24)
		v34 = int64(4278190080)
		v36 = int64(8)
		*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v14<<(uint(v23)%64) | v14&v25<<(uint(v27)%64) | (v14&v30<<(uint(v32)%64) | v14&v34<<(uint(v36)%64)) | (int64(base.Ui64(v14)>>(uint(v36)%64))&v34 | int64(base.Ui64(v14)>>(uint(v32)%64))&v30 | (int64(base.Ui64(v14)>>(uint(v27)%64))&v25 | int64(base.Ui64(v14)>>(uint(v23)%64))))
		*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(0)
		v62 = *(*int32)(unsafe.Add(mBase, _consts[444]))
		v63 = int32(8)
		v68 = F_raxFind(m, v62, v21+v63, v63, v21+int32(4))
		mBase = m.M
		v69 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
		m.G0 = v21 + v20
		if v69 == int32(0) {
			v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v80 | int32(8)
			v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
			if base.Ui32(v84) < base.Ui32(int32(3)) {
				if v9&int32(131072) != 0 {
					return
				} else {
					v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v100 & int32(-131073)
					return
				}
			} else {
				F_addReplyPushLen(m, l0, int32(2))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return
				} else {
					F_addReplyBulkCBuffer(m, l0, int32(_a2466), int32(21))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						v95 = *(*int64)(unsafe.Add(mBase, uint32(v94)+16))
						F_addReplyLongLong(m, l0, v95)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							if v9&int32(131072) != 0 {
								return
							} else {
								v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v100 & int32(-131073)
								return
							}
						}
					}
				}
			}
		} else {
			v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+200)))
			if v75&int32(1088) == int32(0) {
				if v9&int32(131072) != 0 {
				} else {
					v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v106 & int32(-131073)
				}
				v110 = *(*int32)(unsafe.Add(mBase, uint32(v69)+204))
				v112 = v110 | int32(131072)
				*(*int32)(unsafe.Add(mBase, uint32(v69)+204)) = v112
				v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+224)))
				if base.Ui32(v114) < base.Ui32(int32(3)) {
					v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+202)))
					if v117&int32(4) != 0 {
						v132 = int32(0)
						v133 = *(*int32)(unsafe.Add(mBase, _consts[886]))
						v136 = *(*int32)(unsafe.Add(mBase, _consts[887]))
						F_addReplyPubsubMessage(m, v69, v133, v132, v136)
						mBase = m.M
						v138 = m.ExcPending
						if v138 != 0 {
							return
						} else {
							v150 = v69
							v151 = v110
							if l3 == int32(0) {
								F_addReplyArrayLen(m, v150, int32(1))
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return
								} else {
									F_addReplyBulkCBuffer(m, v150, l1, l2)
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										v163 = F_updateClientMemUsageAndBucket(m, v150)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return
										} else {
											if v151&int32(131072) != 0 {
											} else {
												v167 = *(*int32)(unsafe.Add(mBase, uint32(v150)+204))
												*(*int32)(unsafe.Add(mBase, uint32(v150)+204)) = v167 & int32(-131073)
											}
											return
										}
									}
								}
							} else {
								F_addReplyProto(m, v150, l1, l2)
								mBase = m.M
								v157 = m.ExcPending
								if v157 != 0 {
									return
								} else {
									v163 = F_updateClientMemUsageAndBucket(m, v150)
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return
									} else {
										if v151&int32(131072) != 0 {
										} else {
											v167 = *(*int32)(unsafe.Add(mBase, uint32(v150)+204))
											*(*int32)(unsafe.Add(mBase, uint32(v150)+204)) = v167 & int32(-131073)
										}
										return
									}
								}
							}
						}
					} else {
						v123 = v69
						v124 = v110
						v125 = v112
						if v124&int32(131072) != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v123)+204)) = v125 & int32(-131073)
							return
						}
					}
				} else {
					v139 = v69
					v140 = v110
					F_addReplyPushLen(m, v139, int32(2))
					mBase = m.M
					v145 = m.ExcPending
					if v145 != 0 {
						return
					} else {
						F_addReplyBulkCBuffer(m, v139, int32(_a2465), int32(10))
						mBase = m.M
						v149 = m.ExcPending
						if v149 != 0 {
							return
						} else {
							v150 = v139
							v151 = v140
							if l3 == int32(0) {
								F_addReplyArrayLen(m, v150, int32(1))
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return
								} else {
									F_addReplyBulkCBuffer(m, v150, l1, l2)
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										v163 = F_updateClientMemUsageAndBucket(m, v150)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return
										} else {
											if v151&int32(131072) != 0 {
											} else {
												v167 = *(*int32)(unsafe.Add(mBase, uint32(v150)+204))
												*(*int32)(unsafe.Add(mBase, uint32(v150)+204)) = v167 & int32(-131073)
											}
											return
										}
									}
								}
							} else {
								F_addReplyProto(m, v150, l1, l2)
								mBase = m.M
								v157 = m.ExcPending
								if v157 != 0 {
									return
								} else {
									v163 = F_updateClientMemUsageAndBucket(m, v150)
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return
									} else {
										if v151&int32(131072) != 0 {
										} else {
											v167 = *(*int32)(unsafe.Add(mBase, uint32(v150)+204))
											*(*int32)(unsafe.Add(mBase, uint32(v150)+204)) = v167 & int32(-131073)
										}
										return
									}
								}
							}
						}
					}
				}
			} else {
				v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v80 | int32(8)
				v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
				if base.Ui32(v84) < base.Ui32(int32(3)) {
					if v9&int32(131072) != 0 {
						return
					} else {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v100 & int32(-131073)
						return
					}
				} else {
					F_addReplyPushLen(m, l0, int32(2))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						F_addReplyBulkCBuffer(m, l0, int32(_a2466), int32(21))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
							v95 = *(*int64)(unsafe.Add(mBase, uint32(v94)+16))
							F_addReplyLongLong(m, l0, v95)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return
							} else {
								if v9&int32(131072) != 0 {
									return
								} else {
									v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v100 & int32(-131073)
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
func F_trackingBroadcastInvalidationMessages(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int64
	_ = v34
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int64
	_ = v86
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v239 int32
	_ = v239
	v1 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(608)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _consts[888]))
	if v18 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(608)
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[780]))
	if v22 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = v15 + int32(304)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[885]))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = int32(128)
	v34 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+12)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v26)+296)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v26)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v15 + int32(328)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+156)) = v15 + int32(472)
	goto L4
L4:
	;
	v49 = int32(0)
	v51 = F_raxSeek(m, v15+int32(304), int32(_a67), v49, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v55 = F_raxNext(m, v15+int32(304))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	F_raxStop(m, v15+int32(304))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L5
	} else {
		goto L51
	}
L8:
	;
	if v55 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	goto L10
L10:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v15)+316))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
	goto L13
L11:
	;
	goto L7
L12:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	F_raxFree(m, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L47
	}
L13:
	;
	if v73 == int64(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v78 = F_trackingBuildBroadcastReply(m, int32(0), v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(128)
	v86 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+12)) = v86
	*(*int64)(unsafe.Add(mBase, uint32(v15)+296)) = v86
	*(*int64)(unsafe.Add(mBase, uint32(v15)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v15 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+156)) = v15 + int32(168)
	goto L16
L16:
	;
	v99 = int32(0)
	v101 = F_raxSeek(m, v15, int32(_a67), v99, v99)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v103 = F_raxNext(m, v15)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L19
	}
L18:
	;
	F_raxStop(m, v15)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L5
	} else {
		goto L45
	}
L19:
	;
	if v103 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	goto L21
L21:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+205)))
	if v131&int32(1) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L18
L23:
	;
	v184 = F_raxNext(m, v15)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L5
	} else {
		goto L43
	}
L24:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+int32(-1)))))
	switch v168 & int32(7) {
	case 0:
		goto L41
	case 1:
		goto L40
	case 2:
		goto L39
	case 3:
		goto L38
	case 4:
		goto L37
	default:
		v177 = int32(0)
		goto L36
	}
L25:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v137 = F_trackingBuildBroadcastReply(m, v130, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	if v137 == int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+int32(-1)))))
	switch v144 & int32(7) {
	case 0:
		goto L33
	case 1:
		goto L32
	case 2:
		goto L31
	case 3:
		goto L30
	case 4:
		goto L29
	default:
		v161 = int32(0)
		goto L28
	}
L28:
	;
	F_sendTrackingMessage(m, v130, v137, v161, int32(1))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L34
	}
L29:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v137+int32(-17))))
	v161 = v160
	goto L28
L30:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v137+int32(-9))))
	v161 = v157
	goto L28
L31:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137+int32(-5)))))
	v161 = v154
	goto L28
L32:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+int32(-3)))))
	v161 = v151
	goto L28
L33:
	;
	v161 = int32(base.Ui32(v144) >> (uint(int32(3)) % 32))
	goto L28
L34:
	;
	F_sdsfree(m, v137)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	goto L23
L36:
	;
	F_sendTrackingMessage(m, v130, v78, v177, int32(1))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L42
	}
L37:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v78+int32(-17))))
	v177 = v176
	goto L36
L38:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v78+int32(-9))))
	v177 = v175
	goto L36
L39:
	;
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78+int32(-5)))))
	v177 = v174
	goto L36
L40:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+int32(-3)))))
	v177 = v173
	goto L36
L41:
	;
	v177 = int32(base.Ui32(v168) >> (uint(int32(3)) % 32))
	goto L36
L42:
	;
	goto L23
L43:
	;
	if v184 != 0 {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	goto L22
L45:
	;
	F_sdsfree(m, v78)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	goto L12
L47:
	;
	v217 = F_raxNew(m)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v217
	v222 = F_raxNext(m, v15+int32(304))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	if v222 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	goto L11
L51:
	;
	goto L1
}
func F_trackingInvalidateKeysOnFlush(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[780]))
	if v14 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v128 = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, _consts[888]))
	if v129 == v128 {
		goto L27
	} else {
		goto L28
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v20 = v11 + int32(8)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v21
	goto L3
L3:
	;
	v26 = v11 + int32(8)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v28 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v28 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28+base.B2i32(v31 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v37
	goto L5
L7:
	;
	v46 = v28
	goto L8
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+204)))
	if v53&int32(4) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L1
L10:
	;
	v107 = v11 + int32(8)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v109 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v52 != v59 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+224)))
	v67 = int32(2)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66<<(uint(v67)%32))+uint32(_consts[339])))
	v71 = F_objectGetVal(m, v70)
	mBase = m.M
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+224)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73<<(uint(v67)%32))+uint32(_consts[339])))
	v78 = F_objectGetVal(m, v77)
	mBase = m.M
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+int32(-1)))))
	switch v81 & int32(7) {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	case 4:
		goto L17
	default:
		v98 = int32(0)
		goto L16
	}
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[724]))
	v64 = F_listAddNodeTail(m, v62, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	goto L10
L16:
	;
	F_sendTrackingMessage(m, v52, v71, v98, int32(1))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L14
	} else {
		goto L22
	}
L17:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v78+int32(-17))))
	v98 = v97
	goto L16
L18:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v78+int32(-9))))
	v98 = v94
	goto L16
L19:
	;
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78+int32(-5)))))
	v98 = v91
	goto L16
L20:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+int32(-3)))))
	v98 = v88
	goto L16
L21:
	;
	v98 = int32(base.Ui32(v81) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	goto L10
L23:
	;
	if v109 != 0 {
		v46 = v109
		goto L8
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v109+base.B2i32(v112 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v118
	goto L24
L26:
	;
	goto L9
L27:
	;
	m.G0 = v11 + int32(16)
	return
L28:
	;
	if l0 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v139 = F_raxNew(m)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L14
	} else {
		goto L34
	}
L30:
	;
	F_raxFreeWithCallback(m, v129, int32(1105))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L14
	} else {
		goto L33
	}
L31:
	;
	F_freeTrackingRadixTreeAsync(m, v129)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	goto L29
L34:
	;
	v141 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[889])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[888])) = v139
	goto L27
}
