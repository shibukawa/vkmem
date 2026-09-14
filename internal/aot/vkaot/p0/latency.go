package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_latencyCommandReplyWithLatestEvents(m *base.Module, l0 int32) {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v47 int32
	_ = v47
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
	var v54 int64
	_ = v54
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
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
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
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v153 int64
	_ = v153
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v205 int64
	_ = v205
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v278 int32
	_ = v278
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_latencyCommandReplyWithLatestEvents[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	F_addReplyArrayLen(m, l0, v8+v9)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_latencyCommandReplyWithLatestEvents[0]))
	v15 = F_dictGetIterator(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_dictReleaseIterator(m, v15)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L70
	}
L4:
	;
	v24 = v15 + int32(20)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v120 == int32(0) {
		goto L3
	} else {
		goto L31
	}
L6:
	;
	v31 = v24
	v32 = v28
	goto L9
L7:
	;
	v28 = int32(1)
	goto L6
L8:
	;
	v28 = int32(0)
	goto L6
L9:
	;
	switch v32 {
	case 0:
		goto L14
	default:
		goto L13
	}
L11:
	;
	v32 = int32(0)
	goto L9
L12:
	;
	goto L5
L13:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v112
	if v112 == int32(0) {
		goto L11
	} else {
		goto L30
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v36 != int32(-1) {
		v75 = v36
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v76 = int32(1)
	v77 = v75 + v76
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v77
	v79 = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82+v83+int32(26)))))
	if v87 == int32(255) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v40 != 0 {
		v75 = int32(-1)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v42 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	if v69 != int32(-1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v41)+16)))
	v50 = int64(*(*int8)(unsafe.Add(mBase, uint32(v41)+27)))
	v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v41)+8)))
	v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v41)+12)))
	v53 = int64(*(*int8)(unsafe.Add(mBase, uint32(v41)+26)))
	v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v41)+4)))
	v55 = F_wangHash64(m, v54)
	mBase = m.M
	v57 = F_wangHash64(m, v53+v55)
	mBase = m.M
	v59 = F_wangHash64(m, v52+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v51+v59)
	mBase = m.M
	v63 = F_wangHash64(m, v50+v61)
	mBase = m.M
	v65 = F_wangHash64(m, v49+v63)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v68 = v67
	goto L18
L20:
	;
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+24)))
	v47 = v45 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+24)) = uint16(v47)
	v68 = v41
	goto L18
L21:
	;
	v75 = v69 + int32(-1)
	goto L15
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v75 = v72
	goto L15
L23:
	;
	v102 = int32(2)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v82+v100<<(uint(v102)%32)+int32(4))))
	v31 = v107 + v101<<(uint(v102)%32)
	v32 = int32(1)
	goto L9
L24:
	;
	v91 = v79
	goto L26
L25:
	;
	v91 = v76 << (uint(v87) % 32)
	goto L26
L26:
	;
	if v77 < v91 {
		v100 = v83
		v101 = v77
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if v83 != 0 {
		v120 = v79
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v93 == int32(-1) {
		v120 = v79
		goto L12
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = int64(4294967296)
	v100 = int32(1)
	v101 = int32(0)
	goto L23
L30:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v116
	v120 = v112
	goto L12
L31:
	;
	v127 = v120
	goto L32
L32:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	goto L34
L33:
	;
	goto L3
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	goto L35
L35:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	F_addReplyArrayLen(m, l0, int32(6))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_addReplyBulkCString(m, l0, v131)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v142 = base.I32_rem_s(v133+int32(159), int32(160))
	v145 = v132 + v142<<(uint(int32(3))%32)
	v148 = int64(*(*int32)(unsafe.Add(mBase, uint32(v145+int32(16)))))
	F_addReplyLongLong(m, l0, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v153 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v145+int32(20)))))
	F_addReplyLongLong(m, l0, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v156 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v132)+4)))
	F_addReplyLongLong(m, l0, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v159 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v132)+8)))
	F_addReplyLongLong(m, l0, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v162 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v132)+12)))
	F_addReplyLongLong(m, l0, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v172 = v15 + int32(20)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v173 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	if v268 != 0 {
		v127 = v268
		goto L32
	} else {
		goto L69
	}
L44:
	;
	v179 = v172
	v180 = v176
	goto L47
L45:
	;
	v176 = int32(1)
	goto L44
L46:
	;
	v176 = int32(0)
	goto L44
L47:
	;
	switch v180 {
	case 0:
		goto L52
	default:
		goto L51
	}
L49:
	;
	v180 = int32(0)
	goto L47
L50:
	;
	goto L43
L51:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v260
	if v260 == int32(0) {
		goto L49
	} else {
		goto L68
	}
L52:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v184 != int32(-1) {
		v223 = v184
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v224 = int32(1)
	v225 = v223 + v224
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v225
	v227 = int32(0)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v231+int32(26)))))
	if v235 == int32(255) {
		goto L62
	} else {
		goto L63
	}
L54:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v188 != 0 {
		v223 = int32(-1)
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v190 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+20))
	if v217 != int32(-1) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v197 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v189)+16)))
	v198 = int64(*(*int8)(unsafe.Add(mBase, uint32(v189)+27)))
	v199 = int64(*(*int32)(unsafe.Add(mBase, uint32(v189)+8)))
	v200 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v189)+12)))
	v201 = int64(*(*int8)(unsafe.Add(mBase, uint32(v189)+26)))
	v202 = int64(*(*int32)(unsafe.Add(mBase, uint32(v189)+4)))
	v203 = F_wangHash64(m, v202)
	mBase = m.M
	v205 = F_wangHash64(m, v201+v203)
	mBase = m.M
	v207 = F_wangHash64(m, v200+v205)
	mBase = m.M
	v209 = F_wangHash64(m, v199+v207)
	mBase = m.M
	v211 = F_wangHash64(m, v198+v209)
	mBase = m.M
	v213 = F_wangHash64(m, v197+v211)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v213
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v216 = v215
	goto L56
L58:
	;
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189)+24)))
	v195 = v193 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v189)+24)) = uint16(v195)
	v216 = v189
	goto L56
L59:
	;
	v223 = v217 + int32(-1)
	goto L53
L60:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v223 = v220
	goto L53
L61:
	;
	v250 = int32(2)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v230+v248<<(uint(v250)%32)+int32(4))))
	v179 = v255 + v249<<(uint(v250)%32)
	v180 = int32(1)
	goto L47
L62:
	;
	v239 = v227
	goto L64
L63:
	;
	v239 = v224 << (uint(v235) % 32)
	goto L64
L64:
	;
	if v225 < v239 {
		v248 = v231
		v249 = v225
		goto L61
	} else {
		goto L65
	}
L65:
	;
	if v231 != 0 {
		v268 = v227
		goto L50
	} else {
		goto L66
	}
L66:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v230)+20))
	if v241 == int32(-1) {
		v268 = v227
		goto L50
	} else {
		goto L67
	}
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = int64(4294967296)
	v248 = int32(1)
	v249 = int32(0)
	goto L61
L68:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v264
	v268 = v260
	goto L50
L69:
	;
	goto L33
L70:
	;
	return
}
func F_latencyResetEvent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
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
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int64
	_ = v207
	var v208 int64
	_ = v208
	var v209 int64
	_ = v209
	var v210 int64
	_ = v210
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v213 int64
	_ = v213
	var v215 int64
	_ = v215
	var v217 int64
	_ = v217
	var v219 int64
	_ = v219
	var v221 int64
	_ = v221
	var v223 int64
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_latencyResetEvent[0]))
	v7 = F_dictGetSafeIterator(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v125 = v114
	v126 = int32(0)
	goto L32
L2:
	;
	return int32(0)
L3:
	;
	v18 = v7 + int32(20)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v114 != 0 {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v25 = v18
	v26 = v22
	goto L8
L6:
	;
	v22 = int32(1)
	goto L5
L7:
	;
	v22 = int32(0)
	goto L5
L8:
	;
	switch v26 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v26 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v106
	if v106 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v30 != int32(-1) {
		v69 = v30
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v70 = int32(1)
	v71 = v69 + v70
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v71
	v73 = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v77+int32(26)))))
	if v81 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v34 != 0 {
		v69 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v36 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	if v63 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v43 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35)+16)))
	v44 = int64(*(*int8)(unsafe.Add(mBase, uint32(v35)+27)))
	v45 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35)+8)))
	v46 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35)+12)))
	v47 = int64(*(*int8)(unsafe.Add(mBase, uint32(v35)+26)))
	v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35)+4)))
	v49 = F_wangHash64(m, v48)
	mBase = m.M
	v51 = F_wangHash64(m, v47+v49)
	mBase = m.M
	v53 = F_wangHash64(m, v46+v51)
	mBase = m.M
	v55 = F_wangHash64(m, v45+v53)
	mBase = m.M
	v57 = F_wangHash64(m, v44+v55)
	mBase = m.M
	v59 = F_wangHash64(m, v43+v57)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v62 = v61
	goto L17
L19:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)))
	v41 = v39 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v41)
	v62 = v35
	goto L17
L20:
	;
	v69 = v63 + int32(-1)
	goto L14
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v69 = v66
	goto L14
L22:
	;
	v96 = int32(2)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v76+v94<<(uint(v96)%32)+int32(4))))
	v25 = v101 + v95<<(uint(v96)%32)
	v26 = int32(1)
	goto L8
L23:
	;
	v85 = v73
	goto L25
L24:
	;
	v85 = v70 << (uint(v81) % 32)
	goto L25
L25:
	;
	if v71 < v85 {
		v94 = v77
		v95 = v71
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v77 != 0 {
		v114 = v73
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	if v87 == int32(-1) {
		v114 = v73
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967296)
	v94 = int32(1)
	v95 = int32(0)
	goto L22
L29:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v110
	v114 = v106
	goto L11
L30:
	;
	F_dictReleaseIterator(m, v7)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	return int32(0)
L32:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	goto L34
L33:
	;
	F_dictReleaseIterator(m, v7)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L2
	} else {
		goto L79
	}
L34:
	;
	if l0 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v182 = v7 + int32(20)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v183 != 0 {
		goto L54
	} else {
		goto L55
	}
L36:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_latencyResetEvent[0]))
	v170 = F_dictDelete(m, v169, v127)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L51
	}
L37:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v132 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	if v164-v166 != 0 {
		v174 = v126
		goto L35
	} else {
		goto L50
	}
L39:
	;
	v164 = F_tolower(m, v160)
	mBase = m.M
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	v166 = F_tolower(m, v165)
	mBase = m.M
	goto L38
L40:
	;
	v134 = v127
	v135 = l0
	v136 = v132
	goto L43
L41:
	;
	v160 = int32(0)
	v161 = l0
	goto L39
L42:
	;
	v160 = v157 & int32(255)
	v161 = v156
	goto L39
L43:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v138 == int32(0) {
		v156 = v135
		v157 = v136
		goto L42
	} else {
		goto L45
	}
L44:
	;
	v156 = v150
	v157 = int32(0)
	goto L42
L45:
	;
	v142 = v136 & int32(255)
	if v142 == v138 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v149 = int32(1)
	v150 = v135 + v149
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
	if v151 != 0 {
		v134 = v134 + v149
		v135 = v150
		v136 = v151
		goto L43
	} else {
		goto L49
	}
L47:
	;
	v144 = F_tolower(m, v142)
	mBase = m.M
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	v146 = F_tolower(m, v145)
	mBase = m.M
	if v144 == v146 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v156 = v135
	v157 = v148
	goto L42
L49:
	;
	goto L44
L50:
	;
	goto L36
L51:
	;
	v174 = v126 + int32(1)
	goto L35
L52:
	;
	if v278 != 0 {
		v125 = v278
		v126 = v174
		goto L32
	} else {
		goto L78
	}
L53:
	;
	v189 = v182
	v190 = v186
	goto L56
L54:
	;
	v186 = int32(1)
	goto L53
L55:
	;
	v186 = int32(0)
	goto L53
L56:
	;
	switch v190 {
	case 0:
		goto L61
	default:
		goto L60
	}
L58:
	;
	v190 = int32(0)
	goto L56
L59:
	;
	goto L52
L60:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v270
	if v270 == int32(0) {
		goto L58
	} else {
		goto L77
	}
L61:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v194 != int32(-1) {
		v233 = v194
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v234 = int32(1)
	v235 = v233 + v234
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v235
	v237 = int32(0)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+v241+int32(26)))))
	if v245 == int32(255) {
		goto L71
	} else {
		goto L72
	}
L63:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v198 != 0 {
		v233 = int32(-1)
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v200 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
	if v227 != int32(-1) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v207 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v199)+16)))
	v208 = int64(*(*int8)(unsafe.Add(mBase, uint32(v199)+27)))
	v209 = int64(*(*int32)(unsafe.Add(mBase, uint32(v199)+8)))
	v210 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v199)+12)))
	v211 = int64(*(*int8)(unsafe.Add(mBase, uint32(v199)+26)))
	v212 = int64(*(*int32)(unsafe.Add(mBase, uint32(v199)+4)))
	v213 = F_wangHash64(m, v212)
	mBase = m.M
	v215 = F_wangHash64(m, v211+v213)
	mBase = m.M
	v217 = F_wangHash64(m, v210+v215)
	mBase = m.M
	v219 = F_wangHash64(m, v209+v217)
	mBase = m.M
	v221 = F_wangHash64(m, v208+v219)
	mBase = m.M
	v223 = F_wangHash64(m, v207+v221)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v223
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v226 = v225
	goto L65
L67:
	;
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+24)))
	v205 = v203 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v199)+24)) = uint16(v205)
	v226 = v199
	goto L65
L68:
	;
	v233 = v227 + int32(-1)
	goto L62
L69:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v233 = v230
	goto L62
L70:
	;
	v260 = int32(2)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v240+v258<<(uint(v260)%32)+int32(4))))
	v189 = v265 + v259<<(uint(v260)%32)
	v190 = int32(1)
	goto L56
L71:
	;
	v249 = v237
	goto L73
L72:
	;
	v249 = v234 << (uint(v245) % 32)
	goto L73
L73:
	;
	if v235 < v249 {
		v258 = v241
		v259 = v235
		goto L70
	} else {
		goto L74
	}
L74:
	;
	if v241 != 0 {
		v278 = v237
		goto L59
	} else {
		goto L75
	}
L75:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v240)+20))
	if v251 == int32(-1) {
		v278 = v237
		goto L59
	} else {
		goto L76
	}
L76:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967296)
	v258 = int32(1)
	v259 = int32(0)
	goto L70
L77:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v274
	v278 = v270
	goto L59
L78:
	;
	goto L33
L79:
	;
	return v174
}
func F_latencySpecificCommandsFillCDF(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
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
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	v10 = m.G0
	v12 = v10 - int32(64)
	m.G0 = v12
	v14 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v17 < int32(3) {
		v189 = int32(0)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_setDeferredMapLen(m, l0, v14, v189)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L44
	}
L4:
	;
	v25 = int32(0)
	v26 = int32(2)
	goto L5
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v26<<(uint(int32(2))%32))))
	v36 = F_objectGetVal(m, v35)
	mBase = m.M
	v37 = F_lookupCommandBySds(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	v189 = v176
	goto L3
L7:
	;
	v183 = v26 + int32(1)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v183 < v184 {
		v25 = v176
		v26 = v183
		goto L5
	} else {
		goto L43
	}
L8:
	;
	if v37 == int32(0) {
		v176 = v25
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+148))
	if v41 == int32(0) {
		v73 = v25
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v37)+200))
	if v77 == int32(0) {
		v176 = v73
		goto L7
	} else {
		goto L20
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)+140))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+int32(-1)))))
	switch v48 & int32(7) {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L13
	default:
		v65 = int32(0)
		goto L12
	}
L12:
	;
	F_addReplyBulkCBuffer(m, l0, v45, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(-17))))
	v65 = v64
	goto L12
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(-9))))
	v65 = v61
	goto L12
L15:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45+int32(-5)))))
	v65 = v58
	goto L12
L16:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+int32(-3)))))
	v65 = v55
	goto L12
L17:
	;
	v65 = int32(base.Ui32(v48) >> (uint(int32(3)) % 32))
	goto L12
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v37)+148))
	F_fillCommandCDF(m, l0, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v73 = v25 + int32(1)
	goto L10
L20:
	;
	v81 = v12 + int32(16)
	v82 = int32(1)
	v83 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v81)+24)) = v83
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+15)) = uint8(v82)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = int32(-1)
	if v77 == v83 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v104 = F_hashtableNext(m, v12+int32(16), v12+int32(12))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L26
	}
L22:
	;
	goto L21
L23:
	;
	goto L24
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v77)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+24)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v77)+40)) = v81
	goto L22
L25:
	;
	F_hashtableCleanupIterator(m, v12+int32(16))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L42
	}
L26:
	;
	if v104 == int32(0) {
		v163 = v73
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v111 = v73
	goto L28
L28:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+148))
	if v118 == int32(0) {
		v150 = v111
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v163 = v150
	goto L25
L30:
	;
	v158 = F_hashtableNext(m, v12+int32(16), v12+int32(12))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L40
	}
L31:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117)+140))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+int32(-1)))))
	switch v125 & int32(7) {
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
		v142 = int32(0)
		goto L32
	}
L32:
	;
	F_addReplyBulkCBuffer(m, l0, v122, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L38
	}
L33:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v122+int32(-17))))
	v142 = v141
	goto L32
L34:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v122+int32(-9))))
	v142 = v138
	goto L32
L35:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122+int32(-5)))))
	v142 = v135
	goto L32
L36:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+int32(-3)))))
	v142 = v132
	goto L32
L37:
	;
	v142 = int32(base.Ui32(v125) >> (uint(int32(3)) % 32))
	goto L32
L38:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v117)+148))
	F_fillCommandCDF(m, l0, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v150 = v111 + int32(1)
	goto L30
L40:
	;
	if v158 != 0 {
		v111 = v150
		goto L28
	} else {
		goto L41
	}
L41:
	;
	goto L29
L42:
	;
	v176 = v163
	goto L7
L43:
	;
	goto L6
L44:
	;
	m.G0 = v12 + int32(64)
	return
}
