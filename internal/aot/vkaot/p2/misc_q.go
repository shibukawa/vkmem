package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F___qsort_r(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var __phi26 int32
	_ = __phi26
	var v33 int32
	_ = v33
	var __phi33 int32
	_ = __phi33
	var v34 int32
	_ = v34
	var __phi34 int32
	_ = __phi34
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	v11 = m.G0
	v13 = v11 - int32(208)
	m.G0 = v13
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(1)
	v17 = l2 * l1
	if v17 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(208)
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = l2
	v23 = int32(0) - l2
	__phi26 = l2
	__phi33 = l2
	__phi34 = int32(2)
	v26 = __phi26
	v33 = __phi33
	v34 = __phi34
	goto L3
L3:
	;
	v41 = v33 + l2 + v26
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(16)+v34<<(uint(int32(2))%32)))) = v41
	if base.Ui32(v41) < base.Ui32(v17) {
		__phi26 = v41
		__phi33 = v26
		__phi34 = v34 + int32(1)
		v26 = __phi26
		v33 = __phi33
		v34 = __phi34
		goto L3
	} else {
		goto L5
	}
L4:
	;
	v47 = l0 + v17 + v23
	if base.Ui32(l0) < base.Ui32(v47) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L4
L6:
	;
	F_trinkle(m, v178, l2, l3, l4, v13+int32(8), v179, int32(0), v13+int32(16))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L14
	} else {
		goto L36
	}
L7:
	;
	v50 = int32(1)
	v52 = l0
	v53 = v50
	v61 = v50
	goto L9
L8:
	;
	v178 = l0
	v179 = int32(1)
	goto L6
L9:
	;
	v62 = int32(3)
	if v61&v62 != v62 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v178 = v176
	v179 = v170
	goto L6
L11:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v174 = v172 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v174
	v176 = v52 + l2
	if base.Ui32(v176) < base.Ui32(v47) {
		v52 = v176
		v53 = v170
		v61 = v174
		goto L9
	} else {
		goto L35
	}
L12:
	;
	v99 = v53 + int32(-1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(16)+v99<<(uint(int32(2))%32))))
	if base.Ui32(v103) < base.Ui32(v47-v52) {
		goto L21
	} else {
		goto L22
	}
L13:
	;
	F_sift(m, v52, l2, l3, l4, v53, v13+int32(16))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	v71 = v13 + int32(8)
	v72 = int32(2)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	goto L19
L16:
	;
	v170 = v53 + int32(2)
	goto L11
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = int32(base.Ui32(v76) >> (uint(v72) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v76<<(uint(int32(30))%32) | int32(base.Ui32(v79)>>(uint(v72)%32))
	goto L16
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	goto L17
L20:
	;
	if v53 != int32(1) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	F_sift(m, v52, l2, l3, l4, v53, v13+int32(16))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L14
	} else {
		goto L24
	}
L22:
	;
	F_trinkle(m, v52, l2, l3, l4, v13+int32(8), v53, int32(0), v13+int32(16))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	goto L20
L25:
	;
	v146 = v13 + int32(8)
	if base.Ui32(int32(31)) < base.Ui32(v99) {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	v120 = v13 + int32(8)
	v121 = int32(1)
	goto L30
L27:
	;
	v170 = int32(0)
	goto L11
L28:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v128 << (uint(v121) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = int32(base.Ui32(v128)>>(uint(int32(31))%32)) | v135<<(uint(v121)%32)
	goto L27
L30:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	goto L28
L31:
	;
	v170 = int32(1)
	goto L11
L32:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v159 << (uint(v157) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v146)+4)) = int32(base.Ui32(v159)>>(uint(int32(32)-v157)%32)) | v160<<(uint(v157)%32)
	goto L31
L33:
	;
	v157 = v53 + int32(-33)
	v158 = v146
	v159 = int32(0)
	goto L32
L34:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v157 = v99
	v158 = v13 + int32(12)
	v159 = v153
	goto L32
L35:
	;
	goto L10
L36:
	;
	if v179 != int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v203 = v178
	v204 = v179
	goto L41
L38:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v197 != int32(1) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v200 == int32(0) {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	if int32(1) < v204 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L1
L43:
	;
	v365 = v203 + v23
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v362 != int32(1) {
		v203 = v365
		v204 = v362
		goto L41
	} else {
		goto L70
	}
L44:
	;
	v255 = v13 + int32(8)
	v256 = int32(2)
	goto L59
L45:
	;
	v216 = v13 + int32(8)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v223 = F___builtin_ctz(m, v220+int32(-1))
	mBase = m.M
	if v223 != 0 {
		v231 = v223
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if base.Ui32(int32(31)) < base.Ui32(v231) {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	goto L46
L48:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v225 = F___builtin_ctz(m, v224)
	mBase = m.M
	if v225 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v229 = v225 | int32(32)
	goto L51
L50:
	;
	v229 = int32(0)
	goto L51
L51:
	;
	v231 = v229
	goto L47
L52:
	;
	v362 = v231 + v204
	goto L43
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+4)) = int32(base.Ui32(v244) >> (uint(v242) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v216))) = v244<<(uint(int32(32)-v242)%32) | int32(base.Ui32(v243)>>(uint(v242)%32))
	goto L52
L54:
	;
	v242 = v231 + int32(-32)
	v243 = v235
	v244 = int32(0)
	goto L53
L55:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v242 = v231
	v243 = v238
	v244 = v235
	goto L53
L56:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v279 ^ int32(7)
	v284 = v13 + int32(8)
	v285 = int32(1)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	goto L63
L57:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v263 << (uint(v256) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+4)) = int32(base.Ui32(v263)>>(uint(int32(30))%32)) | v270<<(uint(v256)%32)
	goto L56
L59:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	goto L57
L60:
	;
	v307 = v203 + v23
	v309 = v13 + int32(16)
	v311 = v204 + int32(-2)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v309+v311<<(uint(int32(2))%32))))
	F_trinkle(m, v307-v315, l2, l3, l4, v13+int32(8), v204+int32(-1), int32(1), v309)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L14
	} else {
		goto L64
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+4)) = int32(base.Ui32(v289) >> (uint(v285) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v284))) = v289<<(uint(int32(31))%32) | int32(base.Ui32(v292)>>(uint(v285)%32))
	goto L60
L63:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	goto L61
L64:
	;
	v327 = v13 + int32(8)
	v328 = int32(1)
	goto L68
L65:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v352 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v351 | v352
	F_trinkle(m, v307, l2, l3, l4, v13+int32(8), v311, v352, v13+int32(16))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L14
	} else {
		goto L69
	}
L66:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v327))) = v335 << (uint(v328) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v327)+4)) = int32(base.Ui32(v335)>>(uint(int32(31))%32)) | v342<<(uint(v328)%32)
	goto L65
L68:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	goto L66
L69:
	;
	v362 = v311
	goto L43
L70:
	;
	if v367 != int32(1) {
		v203 = v365
		v204 = v362
		goto L41
	} else {
		goto L71
	}
L71:
	;
	if v366 != 0 {
		v203 = v365
		v204 = v362
		goto L41
	} else {
		goto L72
	}
L72:
	;
	goto L42
}
func F_queueSentinelConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
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
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	v8 = *(*int32)(unsafe.Add(mBase, _consts[745]))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v40 = F_valkey_malloc(m, int32(16))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L3
	} else {
		goto L8
	}
L2:
	;
	v11 = F_valkey_malloc(m, int32(12))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[745])) = v11
	v14 = F_listCreate(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[745]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v14
	v19 = F_listCreate(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[745]))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v19
	v24 = F_listCreate(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[745]))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v24
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v30 = int32(1005)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v30
	goto L1
L8:
	;
	v44 = F_valkey_malloc(m, l1<<(uint(int32(2))%32))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v44
	v49 = F_sdsdup(m, l3)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v49
	if l1 < int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v79 = int32(_a1363)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v82 != 0 {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	v58 = int32(0)
	goto L13
L13:
	;
	v62 = v58 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0+v62)))
	v66 = F_sdsdup(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L15
	}
L14:
	;
	goto L11
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44+v62))) = v66
	v70 = v58 + int32(1)
	if v70 != l1 {
		v58 = v70
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v179 = F_listAddNodeTail(m, v178, v40)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L3
	} else {
		goto L42
	}
L18:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[745]))
	v124 = int32(0)
	v125 = int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, _consts[736]))
	v128 = F_strcasecmp(m, v127, v78)
	mBase = m.M
	if v128 == v124 {
		v171 = v125
		goto L33
	} else {
		goto L34
	}
L19:
	;
	if v114-v116 != 0 {
		goto L18
	} else {
		goto L31
	}
L20:
	;
	v114 = F_tolower(m, v110)
	mBase = m.M
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v116 = F_tolower(m, v115)
	mBase = m.M
	goto L19
L21:
	;
	v84 = v78
	v85 = v79
	v86 = v82
	goto L24
L22:
	;
	v110 = int32(0)
	v111 = v79
	goto L20
L23:
	;
	v110 = v107 & int32(255)
	v111 = v106
	goto L20
L24:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v88 == int32(0) {
		v106 = v85
		v107 = v86
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v106 = v100
	v107 = int32(0)
	goto L23
L26:
	;
	v92 = v86 & int32(255)
	if v92 == v88 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v99 = int32(1)
	v100 = v85 + v99
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	if v101 != 0 {
		v84 = v84 + v99
		v85 = v100
		v86 = v101
		goto L24
	} else {
		goto L30
	}
L28:
	;
	v94 = F_tolower(m, v92)
	mBase = m.M
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v96 = F_tolower(m, v95)
	mBase = m.M
	if v94 == v96 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	v106 = v85
	v107 = v98
	goto L23
L30:
	;
	goto L25
L31:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[745]))
	v177 = v119 + int32(4)
	goto L17
L32:
	;
	v177 = v123 + base.B2i32(v171 == int32(0))<<(uint(int32(3))%32)
	goto L17
L33:
	;
	goto L32
L34:
	;
	v131 = int32(0)
	v132 = *(*int32)(unsafe.Add(mBase, _consts[737]))
	v133 = F_strcasecmp(m, v132, v78)
	mBase = m.M
	if v133 == v131 {
		v171 = v125
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v136 = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, _consts[738]))
	v138 = F_strcasecmp(m, v137, v78)
	mBase = m.M
	if v138 == v136 {
		v171 = v125
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v141 = int32(0)
	v142 = *(*int32)(unsafe.Add(mBase, _consts[739]))
	v143 = F_strcasecmp(m, v142, v78)
	mBase = m.M
	if v143 == v141 {
		v171 = v125
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v146 = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, _consts[740]))
	v148 = F_strcasecmp(m, v147, v78)
	mBase = m.M
	if v148 == v146 {
		v171 = v125
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v151 = int32(0)
	v152 = *(*int32)(unsafe.Add(mBase, _consts[741]))
	v153 = F_strcasecmp(m, v152, v78)
	mBase = m.M
	if v153 == v151 {
		v171 = v125
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v156 = int32(0)
	v157 = *(*int32)(unsafe.Add(mBase, _consts[742]))
	v158 = F_strcasecmp(m, v157, v78)
	mBase = m.M
	if v158 == v156 {
		v171 = v125
		goto L33
	} else {
		goto L40
	}
L40:
	;
	v161 = int32(0)
	v162 = *(*int32)(unsafe.Add(mBase, _consts[743]))
	v163 = F_strcasecmp(m, v162, v78)
	mBase = m.M
	if v163 == v161 {
		v171 = v125
		goto L33
	} else {
		goto L41
	}
L41:
	;
	v166 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, _consts[744]))
	v168 = F_strcasecmp(m, v167, v78)
	mBase = m.M
	v171 = base.B2i32(v168 == v166)
	goto L33
L42:
	;
	return
}
