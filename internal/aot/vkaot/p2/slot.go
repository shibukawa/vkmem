package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_createSlotImportJob(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v22 int64
	_ = v22
	var v28 int64
	_ = v28
	var v34 int64
	_ = v34
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v114 int64
	_ = v114
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int64
	_ = v143
	var v149 int64
	_ = v149
	var v155 int64
	_ = v155
	var v161 int64
	_ = v161
	var v167 int64
	_ = v167
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v249 int32
	_ = v249
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_valkey_calloc(m, int32(208))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(32))))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(144)))) = v22
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(24))))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(136)))) = v28
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(128)))) = v34
	v40 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(120)))) = v40
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+112)) = v42
	v45 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v45
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v45
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v45
	v52 = F_representSlotRangeList(m, l3)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+168)) = v52
	v55 = F_generateSlotMigrationJobDescription(m, v14, l1)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+188)) = v55
	v59 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v59 < int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v94 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v65 = int32(0)
	goto L7
L7:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v65<<(uint(int32(2))%32))))
	F_setSlotImportingStateInDb(m, v76, l3, int32(1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L5
L9:
	;
	v81 = v65 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v81 < v83 {
		v65 = v81
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	if l0 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L12:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v14)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v97
	F__serverLog(m, int32(2), int32(_a441), v11)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	F__serverAssert(m, int32(_a442), int32(_a443), int32(831))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L38
	}
L15:
	;
	m.G0 = v11 + int32(16)
	return v14
L16:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211+int32(-1)))))
	switch v215 & int32(7) {
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
		v232 = v169
		goto L32
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+16)) = v207
	v209 = v180
	v211 = v178
	goto L16
L18:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v178+int32(-17))))
	v207 = v206
	goto L17
L19:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v178+int32(-9))))
	v207 = v203
	goto L17
L20:
	;
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178+int32(-5)))))
	v207 = v200
	goto L17
L21:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+int32(-3)))))
	v207 = v197
	goto L17
L22:
	;
	v207 = int32(base.Ui32(v185) >> (uint(int32(3)) % 32))
	goto L17
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+156)) = int32(6)
	goto L15
L24:
	;
	v188 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+196)) = uint8(v188)
	goto L23
L25:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v105&int32(1) != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v108 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v110 = base.B2i32(v108 == int64(-1))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+196)) = uint8(v110)
	if v108 == int64(-1) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if l1 == int32(0) {
		goto L14
	} else {
		goto L28
	}
L28:
	;
	v114 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v114
	v118 = int32(40)
	v120 = *(*int64)(unsafe.Add(mBase, uint32(l1+v118)))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(104)))) = v120
	v124 = int32(32)
	v126 = *(*int64)(unsafe.Add(mBase, uint32(l1+v124)))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(96)))) = v126
	v130 = int32(24)
	v132 = *(*int64)(unsafe.Add(mBase, uint32(l1+v130)))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(88)))) = v132
	v136 = int32(16)
	v138 = *(*int64)(unsafe.Add(mBase, uint32(l1+v136)))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(80)))) = v138
	v141 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v142)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v143
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v142+v136)))
	*(*int64)(unsafe.Add(mBase, uint32(v14+v118))) = v149
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v142+v130)))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(48)))) = v155
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v142+v124)))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(56)))) = v161
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v142+v118)))
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(64)))) = v167
	v169 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+156)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v14)+152)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v14
	F_initClientReplicationData(m, l0)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v14)+152))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	if v177 != 0 {
		v209 = v176
		v211 = v177
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v178 = F_generateSyncSlotsEstablishCommand(m, v14)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v14)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+12)) = v178
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+int32(-1)))))
	switch v185 & int32(7) {
	case 0:
		goto L22
	case 1:
		goto L21
	case 2:
		goto L20
	case 3:
		goto L19
	case 4:
		goto L18
	default:
		v207 = int32(0)
		goto L17
	}
L32:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v209)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v233)+40)) = base.I64_extend_i32_u(v232)
	goto L15
L33:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v211+int32(-17))))
	v232 = v231
	goto L32
L34:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v211+int32(-9))))
	v232 = v228
	goto L32
L35:
	;
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211+int32(-5)))))
	v232 = v225
	goto L32
L36:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211+int32(-3)))))
	v232 = v222
	goto L32
L37:
	;
	v232 = int32(base.Ui32(v215) >> (uint(int32(3)) % 32))
	goto L32
L38:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_finishSlotMigrationJob(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
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
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v247 int32
	_ = v247
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v13 = F_sdsnew(m, l2)
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v13
	if v12 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v20 != 0 {
		v82 = v20
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_sdsfree(m, v12)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	v85 = int32(0)
	if v82 != int32(1) {
		v109 = l1
		v111 = v85
		goto L21
	} else {
		goto L22
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = int64(0)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[213]))
	goto L9
L8:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v82 = v77
	goto L6
L9:
	;
	if v28 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_consts[212])))
	v35 = v10 + int32(24)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v36
	goto L11
L11:
	;
	goto L13
L12:
	;
	F_unpauseActions(m, int32(3))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L20
	}
L13:
	;
	v48 = v10 + int32(24)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v50 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v50 == int32(0) {
		goto L12
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50+base.B2i32(v53 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v59
	goto L16
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)+176))
	if v64 == int64(0) {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L8
L20:
	;
	goto L8
L21:
	;
	F_updateSlotMigrationJobState(m, l0, v109)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L28
	}
L22:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+88)))
	if v91&int32(1) == int32(0) {
		v109 = l1
		v111 = v85
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v96 == int32(5) {
		v109 = l1
		v111 = v85
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	if l1 == int32(20) {
		v109 = l1
		v111 = v85
		goto L21
	} else {
		goto L25
	}
L25:
	;
	if v100 != 0 {
		v109 = l1
		v111 = v85
		goto L21
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = l1
	F_clusterDoBeforeSleep(m, int32(64))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v109 = int32(5)
	v111 = int32(1)
	goto L21
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v115 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	F__serverAssert(m, int32(_a444), int32(_a443), int32(2186))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L72
	}
L30:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_sdsfree(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L37
	}
L31:
	;
	if v114 == int32(0) {
		goto L30
	} else {
		goto L35
	}
L32:
	;
	if v114 != 0 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+216)) = int32(0)
	F_freeClientAsync(m, v115)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
	goto L30
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+52))
	m.T0[v127].(func(*base.Module, int32))(m, v114)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = int32(0)
	goto L30
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = int32(0)
	if v111 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	m.G0 = v10 + int32(32)
	return
L39:
	;
	if v109 == int32(18) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v167 != int32(1) {
		v219 = v167
		goto L51
	} else {
		goto L52
	}
L41:
	;
	v141 = int32(3)
	goto L43
L42:
	;
	v141 = int32(2)
	goto L43
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v141 < v143 {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if base.Ui32(int32(20)) < base.Ui32(v109) {
		v154 = int32(_a288)
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v145
	if v155 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v109<<(uint(int32(2))%32))+uint32(_consts[214])))
	v154 = v153
	goto L45
L47:
	;
	v159 = v155
	goto L49
L48:
	;
	v159 = int32(_a445)
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v159
	F__serverLog(m, v141, int32(_a446), v10)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L40
L51:
	;
	v225 = base.B2i32(v109 == int32(20))
	if v109 == int32(20) {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v171 < int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+88)))
	if v207&int32(1) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L54:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v180 = int32(0)
	goto L55
L55:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184+v180<<(uint(int32(2))%32))))
	F_setSlotImportingStateInDb(m, v188, v174, int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	goto L53
L57:
	;
	v193 = v180 + int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v193 < v195 {
		v180 = v193
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v219 = v214
	goto L51
L60:
	;
	F_propagateSyncSlotsFinish(m, l0)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v226 = int32(4)
	goto L64
L63:
	;
	v226 = int32(2)
	goto L64
L64:
	;
	if v109 == int32(20) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v229 = int32(5)
	goto L67
L66:
	;
	v229 = int32(3)
	goto L67
L67:
	;
	if v219 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v230 = v226
	goto L70
L69:
	;
	v230 = v229
	goto L70
L70:
	;
	F_fireModuleSlotMigrationEvent(m, l0, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	goto L38
L72:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_getSlotOrReply(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = F_getLongLongFromObject(m, l1, v7+int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != 0 {
			F_addReplyError(m, l0, int32(_a343))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v24 = int32(-1)
				m.G0 = v7 + int32(16)
				return v24
			}
		} else {
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
			if base.Ui64(v15) < base.Ui64(int64(16384)) {
				v24 = base.I32_wrap_i64(v15)
				m.G0 = v7 + int32(16)
				return v24
			} else {
				F_addReplyError(m, l0, int32(_a343))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v24 = int32(-1)
					m.G0 = v7 + int32(16)
					return v24
				}
			}
		}
	}
}
func F_parseSlotRangesOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	v13 = m.G0
	v15 = v13 - int32(64)
	m.G0 = v15
	v17 = F_listCreate(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(102)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v25
	if v25 <= l1 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v15 + int32(64)
	return v241
L4:
	;
	F_listRelease(m, v17)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L56
	}
L5:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v217 != 0 {
		v241 = v17
		goto L3
	} else {
		goto L54
	}
L6:
	;
	v29 = l1
	goto L8
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v54
	F_addReplyErrorFormat(m, l0, int32(_a435), v15)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L53
	}
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = v29 << (uint(int32(2)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v42)))
	v46 = F_getLongLongFromObject(m, v44, int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v54
	F_addReplyErrorFormat(m, l0, int32(_a436), v15+int32(32))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L51
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51+v42)))
	v54 = F_getSlotOrReply(m, l0, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	if v46 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v29
	goto L5
L13:
	;
	if v54 == int32(-1) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v59 = v29 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v59 < v60 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v59<<(uint(int32(2))%32))))
	v70 = F_getSlotOrReply(m, l0, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	F_addReplyError(m, l0, int32(_a437))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L4
L18:
	;
	if v70 == int32(-1) {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	if v70 < v54 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v85 = v54
	goto L21
L21:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v76+int32(52)+v85<<(uint(int32(2))%32))))
	if v94 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v111 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L32
	}
L23:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v101 != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v85
	F_addReplyErrorFormat(m, l0, int32(_a438), v15+int32(16))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	goto L4
L26:
	;
	if v85 != v70 {
		v85 = v85 + int32(1)
		goto L21
	} else {
		goto L31
	}
L27:
	;
	if v101 == v94 {
		goto L26
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v94
	goto L26
L29:
	;
	F_addReplyError(m, l0, int32(_a439))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L4
L31:
	;
	goto L22
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v54
	v116 = v15 + int32(56)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v117
	goto L33
L33:
	;
	v122 = v15 + int32(56)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v124 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L9
L35:
	;
	v183 = F_listAddNodeTail(m, v17, v111)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L49
	}
L36:
	;
	if v124 == int32(0) {
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v124+base.B2i32(v127 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v133
	goto L37
L39:
	;
	v143 = v124
	goto L40
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if v149 < v151 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L35
L42:
	;
	v158 = v15 + int32(56)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v160 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	if v153 <= v154 {
		goto L34
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	if v160 != 0 {
		v143 = v160
		goto L40
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v160+base.B2i32(v163 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v169
	goto L46
L48:
	;
	goto L41
L49:
	;
	v186 = v29 + int32(2)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v186 < v187 {
		v29 = v186
		goto L8
	} else {
		goto L50
	}
L50:
	;
	goto L5
L51:
	;
	F_valkey_free(m, v111)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L4
L53:
	;
	goto L4
L54:
	;
	F_addReplyError(m, l0, int32(_a440))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	goto L4
L56:
	;
	v241 = int32(0)
	goto L3
}
func F_setSlotImportingStateInDb(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v14 = v9 + int32(8)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v15
	goto L3
L3:
	;
	v20 = v9 + int32(8)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v22 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v22 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22+base.B2i32(v25 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v31
	goto L5
L7:
	;
	v36 = v22
	goto L8
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v43 < v42 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L1
L10:
	;
	v71 = v9 + int32(8)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v73 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L11:
	;
	v46 = v42
	goto L12
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_kvstoreSetIsImporting(m, v51, v46, l2)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L10
L14:
	;
	return
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_kvstoreSetIsImporting(m, v54, v46, l2)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_kvstoreSetIsImporting(m, v57, v46, l2)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v46 < v60 {
		v46 = v46 + int32(1)
		goto L12
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	if v73 != 0 {
		v36 = v73
		goto L8
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v73+base.B2i32(v76 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v82
	goto L20
L22:
	;
	goto L9
}
func F_slotExportBeginStreaming(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11 != 0 {
		F__serverAssert(m, int32(_a469), int32(_a443), int32(1443))
		mBase = m.M
		v88 = m.ExcPending
		if v88 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[6]))
		if int32(2) < v13 {
			v41 = *(*int64)(unsafe.Add(mBase, _consts[47]))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(14)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v41
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+200))
			if v48&int32(4194304) != 0 {
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+104))
				if v51 == int32(0) {
					v58 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v45)
					mBase = m.M
					if v58 == int32(0) {
					} else {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v45)+200))
						*(*int32)(unsafe.Add(mBase, uint32(v45)+200)) = v61 | int32(4194304)
						v66 = *(*int32)(unsafe.Add(mBase, _consts[216]))
						F_listLinkNodeHead(m, v66, v45+int32(168))
						mBase = m.M
					}
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
					switch v54 {
					case 0:
						v58 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v45)
						mBase = m.M
						if v58 == int32(0) {
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v45)+200))
							*(*int32)(unsafe.Add(mBase, uint32(v45)+200)) = v61 | int32(4194304)
							v66 = *(*int32)(unsafe.Add(mBase, _consts[216]))
							F_listLinkNodeHead(m, v66, v45+int32(168))
							mBase = m.M
						}
					default:
					case 9, 11:
						if v48&int32(1024) != 0 {
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
							if v57 != 0 {
							} else {
								v58 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v45)
								mBase = m.M
								if v58 == int32(0) {
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v45)+200))
									*(*int32)(unsafe.Add(mBase, uint32(v45)+200)) = v61 | int32(4194304)
									v66 = *(*int32)(unsafe.Add(mBase, _consts[216]))
									F_listLinkNodeHead(m, v66, v45+int32(168))
									mBase = m.M
								}
							}
						}
					}
				}
			}
			v72 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(2) < v72 {
				m.G0 = v9 + int32(32)
				return
			} else {
				v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v75
				F__serverLog(m, int32(2), int32(_a477), v9)
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					m.G0 = v9 + int32(32)
					return
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
			if base.Ui32(int32(20)) < base.Ui32(v18) {
				v26 = int32(_a288)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v18<<(uint(int32(2))%32))+uint32(_consts[214])))
				v26 = v25
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a478)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v26
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v16
			F__serverLog(m, int32(2), int32(_a450), v9+int32(16))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				v41 = *(*int64)(unsafe.Add(mBase, _consts[47]))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(14)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v41
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+200))
				if v48&int32(4194304) != 0 {
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+104))
					if v51 == int32(0) {
						v58 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v45)
						mBase = m.M
						if v58 == int32(0) {
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v45)+200))
							*(*int32)(unsafe.Add(mBase, uint32(v45)+200)) = v61 | int32(4194304)
							v66 = *(*int32)(unsafe.Add(mBase, _consts[216]))
							F_listLinkNodeHead(m, v66, v45+int32(168))
							mBase = m.M
						}
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
						switch v54 {
						case 0:
							v58 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v45)
							mBase = m.M
							if v58 == int32(0) {
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v45)+200))
								*(*int32)(unsafe.Add(mBase, uint32(v45)+200)) = v61 | int32(4194304)
								v66 = *(*int32)(unsafe.Add(mBase, _consts[216]))
								F_listLinkNodeHead(m, v66, v45+int32(168))
								mBase = m.M
							}
						default:
						case 9, 11:
							if v48&int32(1024) != 0 {
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
								if v57 != 0 {
								} else {
									v58 = F_clusterSlotMigrationShouldInstallWriteHandler(m, v45)
									mBase = m.M
									if v58 == int32(0) {
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v45)+200))
										*(*int32)(unsafe.Add(mBase, uint32(v45)+200)) = v61 | int32(4194304)
										v66 = *(*int32)(unsafe.Add(mBase, _consts[216]))
										F_listLinkNodeHead(m, v66, v45+int32(168))
										mBase = m.M
									}
								}
							}
						}
					}
				}
				v72 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				if int32(2) < v72 {
					m.G0 = v9 + int32(32)
					return
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v75
					F__serverLog(m, int32(2), int32(_a477), v9)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
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
func F_slotMigrationJobReadAuthResponse(m *base.Module, l0 int32) {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v78 int64
	_ = v78
	var v83 int32
	_ = v83
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+200))
	v13 = F_receiveSynchronousResponse(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v19 != int32(45) {
				F_sdsfree(m, v13)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					if int32(2) < v38 {
						v78 = *(*int64)(unsafe.Add(mBase, _consts[47]))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = int32(10)
						*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v78
						F_proceedWithSlotMigration(m, v11)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							m.G0 = v9 + int32(48)
							return
						}
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v11)+188))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v41
						F__serverLog(m, int32(2), int32(_a473), v9+int32(32))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, _consts[6]))
							if int32(2) < v50 {
								v78 = *(*int64)(unsafe.Add(mBase, _consts[47]))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = int32(10)
								*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v78
								F_proceedWithSlotMigration(m, v11)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									m.G0 = v9 + int32(48)
									return
								}
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+188))
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+156))
								if base.Ui32(int32(20)) < base.Ui32(v55) {
									v63 = int32(_a288)
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v55<<(uint(int32(2))%32))+uint32(_consts[214])))
									v63 = v62
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a474)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v53
								F__serverLog(m, int32(2), int32(_a450), v9+int32(16))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									v78 = *(*int64)(unsafe.Add(mBase, _consts[47]))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = int32(10)
									*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v78
									F_proceedWithSlotMigration(m, v11)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										m.G0 = v9 + int32(48)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v22 = F_sdsempty(m)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13
					v27 = F_sdscatfmt(m, v22, int32(_a475), v9)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						F_finishSlotMigrationJob(m, v11, int32(18), v27)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							F_sdsfree(m, v13)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_sdsfree(m, v27)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									m.G0 = v9 + int32(48)
									return
								}
							}
						}
					}
				}
			}
		} else {
			F_finishSlotMigrationJob(m, v11, int32(18), int32(_a476))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				m.G0 = v9 + int32(48)
				return
			}
		}
	}
}
func F_slotMigrationJobSendAuth(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v69 int64
	_ = v69
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11 != 0 {
		F__serverAssert(m, int32(_a469), int32(_a443), int32(1387))
		mBase = m.M
		v84 = m.ExcPending
		if v84 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[215]))
		if v13 == int32(0) {
			F__serverAssert(m, int32(_a470), int32(_a443), int32(1388))
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			v17 = F_replicationSendAuth(m, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				if v17 == int32(0) {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+84))
					v40 = m.T0[v39].(func(*base.Module, int32, int32) int32)(m, v36, int32(106))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if int32(2) < v43 {
							v69 = *(*int64)(unsafe.Add(mBase, _consts[47]))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(9)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v69
							m.G0 = v9 + int32(32)
							return
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
							if base.Ui32(int32(20)) < base.Ui32(v48) {
								v56 = int32(_a288)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v48<<(uint(int32(2))%32))+uint32(_consts[214])))
								v56 = v55
							}
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(_a471)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v56
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
							F__serverLog(m, int32(2), int32(_a450), v9)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								v69 = *(*int64)(unsafe.Add(mBase, _consts[47]))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(9)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v69
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				} else {
					v21 = F_sdsempty(m)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
						v28 = F_sdscatfmt(m, v21, int32(_a472), v9+int32(16))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							F_finishSlotMigrationJob(m, l0, int32(18), v28)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_sdsfree(m, v17)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									F_sdsfree(m, v28)
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
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
	}
}
func F_updateSlotMigrationJobState(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v47 int64
	_ = v47
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v13 {
		v47 = *(*int64)(unsafe.Add(mBase, _consts[47]))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v47
		m.G0 = v10 + int32(16)
		return
	} else {
		v16 = int32(_a288)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
		if base.Ui32(int32(20)) < base.Ui32(v18) {
			v26 = v16
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v18<<(uint(int32(2))%32))+uint32(_consts[214])))
			v26 = v25
		}
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
		if base.Ui32(int32(20)) < base.Ui32(l1) {
			v35 = v16
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32))+uint32(_consts[214])))
			v35 = v34
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v35
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v26
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v27
		F__serverLog(m, int32(2), int32(_a450), v10)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			v47 = *(*int64)(unsafe.Add(mBase, _consts[47]))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = l1
			*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v47
			m.G0 = v10 + int32(16)
			return
		}
	}
}
