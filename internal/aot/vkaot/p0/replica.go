package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_replicaAfterLoadPrimaryRDB(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v44 int64
	_ = v44
	var v51 int64
	_ = v51
	var v58 int64
	_ = v58
	var v65 int64
	_ = v65
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v87 int32
	_ = v87
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int64
	_ = v126
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	v6 = *(*int32)(unsafe.Add(mBase, _consts[527]))
	if l0 != v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v25 = int32(0)
	F_moduleFireServerEvent(m, int64(7), v25, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L8
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_replicationCreatePrimaryClientWithHandler(m, v11, v12, int32(107))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L6
	}
L3:
	;
	F_dualChannelSyncHandleRdbLoadCompletion(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
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
	v16 = int32(_a69)
	*(*int64)(unsafe.Add(mBase, _consts[519])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[518])) = int32(14)
	F_replicationSendAck(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	if v30 != int32(14) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v80 = int32(_a69)
	v81 = int64(3472328296227680304)
	*(*int64)(unsafe.Add(mBase, _consts[512])) = v81
	*(*int64)(unsafe.Add(mBase, _consts[508])) = int64(-1)
	v87 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[513])) = uint8(v87)
	*(*int64)(unsafe.Add(mBase, _consts[514])) = v81
	*(*int64)(unsafe.Add(mBase, _consts[515])) = v81
	*(*int64)(unsafe.Add(mBase, _consts[516])) = v81
	*(*int64)(unsafe.Add(mBase, _consts[517])) = v81
	v110 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	if v110 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v33 = int32(_a69)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+104))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v36)+104))
	*(*int64)(unsafe.Add(mBase, _consts[528])) = v37
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v36+int32(112))))
	*(*int64)(unsafe.Add(mBase, _consts[529])) = v44
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v36+int32(120))))
	*(*int64)(unsafe.Add(mBase, _consts[530])) = v51
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v36+int32(128))))
	*(*int64)(unsafe.Add(mBase, _consts[531])) = v58
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v36+int32(136))))
	*(*int64)(unsafe.Add(mBase, _consts[532])) = v65
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+int32(144)))))
	*(*uint8)(unsafe.Add(mBase, _consts[511])) = uint8(v72)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v35)+104))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v75)+48))
	*(*int64)(unsafe.Add(mBase, _consts[31])) = v76
	goto L9
L11:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v133 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v113 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[448])) = v113
	*(*int64)(unsafe.Add(mBase, uint32(v113))) = int64(0)
	v118 = F_raxNew(m)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v120 = int32(_a69)
	v121 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = v118
	v126 = *(*int64)(unsafe.Add(mBase, _consts[31]))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+24)) = v126 + int64(1)
	goto L11
L15:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _consts[533]))
	if v142 != int32(2) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	F__serverLog(m, int32(2), int32(_a948), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	if v148 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	goto L18
L21:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[527]))
	if l0 != v164 {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	if l2 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_restartAOFAfterSYNC(m)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L28
	}
L24:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[534]))
	if v154 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v157 = F_restartAOFWithSyncRdb(m)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	if v157 != int32(-1) {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	goto L21
L29:
	;
	return
L30:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+52))
	m.T0[v167].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[527])) = int32(0)
	goto L29
}
func F_replicaLoadPrimaryRDBFromDisk(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
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
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int64
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = int32(_a69)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[535]))
	v13 = int32(-1)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[539]))
	v16 = F_fsync(m, v15)
	mBase = m.M
	if v16 != v13 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(48)
	return v309
L2:
	;
	v33 = int32(_a69)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	v37 = F_open(m, v34, int32(2048), int32(0))
	mBase = m.M
	v39 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	v41 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	v42 = F_rename(m, v39, v41)
	mBase = m.M
	if v42 != int32(-1) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v20 {
		v309 = v13
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v25 = F___strerror_l(m, v24, v24)
	mBase = m.M
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v25
	F__serverLog(m, int32(3), int32(_a961), v9)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v309 = v13
	goto L1
L9:
	;
	if v37 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v46 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v63 = int32(-1)
	if v37 == v63 {
		v309 = v63
		goto L1
	} else {
		goto L16
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	goto L13
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v53 = F___strerror_l(m, v52, v52)
	mBase = m.M
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v50
	F__serverLog(m, int32(3), int32(_a962), v9+int32(16))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v66 = F_close(m, v37)
	mBase = m.M
	v309 = v63
	goto L1
L17:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	v78 = m.G0
	v80 = v78 - int32(4112)
	m.G0 = v80
	v82 = F_strlen(m, v74)
	mBase = m.M
	if base.Ui32(v82) < base.Ui32(int32(4097)) {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v69 = int32(0)
	F_bioCreateCloseJob(m, v37, v69, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	F_replicationAttachToNewPrimary(m)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L7
	} else {
		goto L39
	}
L21:
	;
	if v121 != int32(-1) {
		goto L20
	} else {
		goto L34
	}
L22:
	;
	m.G0 = v80 + int32(4112)
	goto L21
L23:
	;
	v91 = F___memcpy(m, v80, v74, v82+int32(1))
	mBase = m.M
	v92 = F_dirname(m, v91)
	mBase = m.M
	v93 = int32(0)
	v95 = F_open(m, v92, v93, v93)
	mBase = m.M
	if v95 != int32(-1) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v85 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(37)
	v121 = int32(-1)
	goto L22
L25:
	;
	v105 = F_fsync(m, v95)
	mBase = m.M
	if v105 != int32(-1) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v100 = F___errno_location(m)
	mBase = m.M
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v101 != int32(31) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v104 = int32(-1)
	goto L29
L28:
	;
	v104 = int32(0)
	goto L29
L29:
	;
	v121 = v104
	goto L22
L30:
	;
	v119 = F_close(m, v95)
	mBase = m.M
	v121 = int32(0)
	goto L22
L31:
	;
	v108 = F___errno_location(m)
	mBase = m.M
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v109 == int32(8) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	if v109 == int32(28) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v114 = F_close(m, v95)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v109
	v121 = int32(-1)
	goto L22
L34:
	;
	v129 = int32(-1)
	v131 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v131 {
		v309 = v129
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	goto L36
L36:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v138 = F___strerror_l(m, v137, v137)
	mBase = m.M
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v135
	F__serverLog(m, int32(3), int32(_a963), v9+int32(32))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	v309 = v129
	goto L1
L39:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v150 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	v161 = F_rdbLoad(m, v159, l0, int32(34))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L7
	} else {
		goto L44
	}
L41:
	;
	F__serverLog(m, int32(2), int32(_a950), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v243 = int32(_a69)
	v244 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	v245 = int32(0)
	v248 = *(*int32)(unsafe.Add(mBase, _consts[534]))
	if base.B2i32(v244 == v245)|base.B2i32(v248 == v245) != int32(1) {
		goto L69
	} else {
		goto L70
	}
L44:
	;
	if v161 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v166 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _consts[541]))
	if v175 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	F__serverLog(m, int32(3), int32(_a960), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v161 != int32(2) {
		goto L61
	} else {
		goto L62
	}
L50:
	;
	v178 = int32(0)
	v179 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	v181 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	goto L51
L51:
	;
	if base.B2i32(v179|v181 == v178) == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v188 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	v200 = F_open(m, v197, int32(2048), int32(0))
	mBase = m.M
	v201 = F_unlink(m, v197)
	mBase = m.M
	if v200 == int32(-1) {
		goto L49
	} else {
		goto L56
	}
L54:
	;
	F__serverLog(m, int32(2), int32(_a958), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	if v201 != int32(-1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v210 = int32(0)
	F_bioCreateCloseJob(m, v200, v210, v210)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L7
	} else {
		goto L60
	}
L58:
	;
	v206 = int32(9116376)
	goto L59
L59:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v208 = F_close(m, v200)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v207
	goto L49
L60:
	;
	goto L49
L61:
	;
	if int32(2) < v218 {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v221 = int32(-1)
	if int32(2) < v218 {
		v309 = v221
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F__serverLog(m, int32(2), int32(_a959), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	v309 = v221
	goto L1
L65:
	;
	v238 = int32(-1)
	v241 = F_emptyData(m, v238, base.B2i32(v12 != int32(0)), int32(968))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L7
	} else {
		goto L68
	}
L66:
	;
	F__serverLog(m, int32(2), int32(_a955), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v309 = v238
	goto L1
L69:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	F_valkey_free(m, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L7
	} else {
		goto L82
	}
L70:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _consts[541]))
	if v255 == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v258 = int32(0)
	v259 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	v261 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	goto L72
L72:
	;
	if base.B2i32(v259|v261 == v258) == int32(0) {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v268 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v277 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	v280 = F_open(m, v277, int32(2048), int32(0))
	mBase = m.M
	v281 = F_unlink(m, v277)
	mBase = m.M
	if v280 == int32(-1) {
		goto L69
	} else {
		goto L77
	}
L75:
	;
	F__serverLog(m, int32(2), int32(_a958), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	if v281 != int32(-1) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v290 = int32(0)
	F_bioCreateCloseJob(m, v280, v290, v290)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L7
	} else {
		goto L81
	}
L79:
	;
	v286 = int32(9116376)
	goto L80
L80:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v288 = F_close(m, v280)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v287
	goto L69
L81:
	;
	goto L69
L82:
	;
	v301 = int32(_a69)
	v302 = *(*int32)(unsafe.Add(mBase, _consts[539]))
	v303 = F_close(m, v302)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[539])) = int64(4294967295)
	v309 = int32(0)
	goto L1
}
func F_replicaLoadPrimaryRDBFromSocket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
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
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int64
	_ = v165
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int64
	_ = v220
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v300 int32
	_ = v300
	var v316 int32
	_ = v316
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int64
	_ = v404
	var v405 int32
	_ = v405
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int64
	_ = v456
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v482 int32
	_ = v482
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v17 = int32(_a69)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[535]))
	v20 = *(*int32)(unsafe.Add(mBase, _consts[536]))
	if v20 != int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _consts[537]))
	F_rioInitWithConn(m, v15+int32(16), l0, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L26
	}
L2:
	;
	F_replicationAttachToNewPrimary(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L24
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v27 = F_valkey_calloc(m, v24<<(uint(int32(2))%32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v31 = F_functionsLibCtxCreate(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v34 = int32(0)
	F_moduleFireServerEvent(m, int64(14), v34, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L12
L8:
	;
	v120 = v27
	v121 = v31
	v122 = base.B2i32(v108 == int32(0))
	v123 = v31
	v124 = v27
	goto L1
L9:
	;
	v108 = int32(0)
	goto L8
L10:
	;
	v80 = v75
	v81 = v76
	v82 = v77
	goto L20
L11:
	;
	if v65 == int32(0) {
		goto L9
	} else {
		goto L18
	}
L12:
	;
	goto L13
L13:
	;
	v52 = int32(_a920)
	v53 = int32(_a949)
	v54 = int32(40)
	goto L14
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v57 != v58 {
		v75 = v52
		v76 = v53
		v77 = v54
		goto L10
	} else {
		goto L16
	}
L15:
	;
	goto L11
L16:
	;
	v60 = int32(4)
	v61 = v53 + v60
	v63 = v52 + v60
	v65 = v54 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v65) {
		v52 = v63
		v53 = v61
		v54 = v65
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v75 = v63
	v76 = v61
	v77 = v65
	goto L10
L19:
	;
	v108 = v85 - v86
	goto L8
L20:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v85 != v86 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v88 = int32(1)
	v93 = v82 + int32(-1)
	if v93 == int32(0) {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v80 = v80 + v88
	v81 = v81 + v88
	v82 = v93
	goto L20
L24:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v116 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	goto L25
L25:
	;
	v117 = int32(0)
	v120 = v117
	v121 = v116
	v122 = v117
	v123 = v117
	v124 = v114
	goto L1
L26:
	;
	v131 = F_connBlock(m, l0)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[538]))
	v138 = F_connRecvTimeout(m, l0, base.I64_extend_i32_s(v134*int32(1000)))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v141 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[537]))
	F_startLoading(m, v150, int32(2), v122)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	F__serverLog(m, int32(2), int32(_a950), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v154 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v124
	v174 = int32(2)
	v177 = *(*int32)(unsafe.Add(mBase, _consts[536]))
	if v177 == v174 {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+132))
	if v158 == int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v161 = m.T0[v158].(func(*base.Module) int32)(m)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	if v161 == int32(0) {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v165 | int64(8)
	goto L33
L38:
	;
	m.G0 = v15 + int32(96)
	return v482
L39:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _consts[536]))
	if v419 != int32(2) {
		goto L103
	} else {
		goto L104
	}
L40:
	;
	F_stopLoading(m, int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L4
	} else {
		goto L86
	}
L41:
	;
	F__serverLog(m, int32(3), v328, int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L85
	}
L42:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v192 == int32(0) {
		goto L39
	} else {
		goto L49
	}
L43:
	;
	v180 = v174
	goto L45
L44:
	;
	v180 = int32(34)
	goto L45
L45:
	;
	v183 = F_rdbLoadRioWithLoadingCtxScopedRdb(m, v15+int32(16), v180, l4, v15+int32(8))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	if v183 == int32(0) {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v188 {
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v328 = int32(_a951)
	goto L41
L49:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+48)))
	if v195&int32(5) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v316 {
		goto L40
	} else {
		goto L84
	}
L51:
	;
	v203 = l1
	v207 = int32(40)
	goto L52
L52:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	if base.Ui32(v213) < base.Ui32(v207) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v236 = int32(40)
	goto L71
L54:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	if v224 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L55:
	;
	v215 = v213
	goto L57
L56:
	;
	v215 = v207
	goto L57
L57:
	;
	if v213 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v216 = v215
	goto L60
L59:
	;
	v216 = v207
	goto L60
L60:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v218 = m.T0[v217].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(16), v203, v216)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	if v218 != 0 {
		goto L54
	} else {
		goto L62
	}
L62:
	;
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v220 | int64(1)
	goto L50
L63:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v231 + v216
	v235 = v207 - v216
	if v235 != 0 {
		v203 = v203 + v216
		v207 = v235
		goto L52
	} else {
		goto L66
	}
L64:
	;
	m.T0[v224].(func(*base.Module, int32, int32, int32))(m, v15+int32(16), v203, v216)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	goto L53
L67:
	;
	if v300 == int32(0) {
		goto L39
	} else {
		goto L83
	}
L68:
	;
	v300 = int32(0)
	goto L67
L69:
	;
	v272 = v267
	v273 = v268
	v274 = v269
	goto L79
L70:
	;
	if v257 == int32(0) {
		goto L68
	} else {
		goto L77
	}
L71:
	;
	if (l2|l1)&int32(3) != 0 {
		v267 = l1
		v268 = l2
		v269 = v236
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v244 = l1
	v245 = l2
	v246 = v236
	goto L73
L73:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	if v249 != v250 {
		v267 = v244
		v268 = v245
		v269 = v246
		goto L69
	} else {
		goto L75
	}
L74:
	;
	goto L70
L75:
	;
	v252 = int32(4)
	v253 = v245 + v252
	v255 = v244 + v252
	v257 = v246 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v257) {
		v244 = v255
		v245 = v253
		v246 = v257
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v267 = v255
	v268 = v253
	v269 = v257
	goto L69
L78:
	;
	v300 = v277 - v278
	goto L67
L79:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	if v277 != v278 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v280 = int32(1)
	v285 = v274 + int32(-1)
	if v285 == int32(0) {
		goto L68
	} else {
		goto L82
	}
L82:
	;
	v272 = v272 + v280
	v273 = v273 + v280
	v274 = v285
	goto L79
L83:
	;
	goto L50
L84:
	;
	v328 = int32(_a952)
	goto L41
L85:
	;
	goto L40
L86:
	;
	F_rioFreeConn(m, v15+int32(16), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _consts[536]))
	if v357 != int32(2) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v183 != int32(2) {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	F_moduleFireServerEvent(m, int64(14), int32(1), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	F_discardTempDb(m, v120)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	F_freeFunctionsAsync(m, v123, int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v370 = int32(-1)
	v372 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v372 {
		v482 = v370
		goto L38
	} else {
		goto L93
	}
L93:
	;
	F__serverLog(m, int32(2), int32(_a953), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	v482 = v370
	goto L38
L95:
	;
	if int32(2) < v381 {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	v384 = int32(-1)
	if int32(2) < v381 {
		v482 = v384
		goto L38
	} else {
		goto L97
	}
L97:
	;
	F__serverLog(m, int32(2), int32(_a954), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	v482 = v384
	goto L38
L99:
	;
	v401 = int32(-1)
	v404 = F_emptyData(m, v401, base.B2i32(v18 != int32(0)), int32(968))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L102
	}
L100:
	;
	F__serverLog(m, int32(2), int32(_a955), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v482 = v401
	goto L38
L103:
	;
	v454 = int32(_a69)
	v456 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v456 + int64(1)
	F_stopLoading(m, int32(1))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L115
	}
L104:
	;
	F_replicationAttachToNewPrimary(m)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	v425 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v425 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	F_swapMainDbWithTempDb(m, v120)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L4
	} else {
		goto L109
	}
L107:
	;
	F__serverLog(m, int32(2), int32(_a956), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	F_functionsLibCtxSwapWithCurrent(m, v123, int32(1))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	F_moduleFireServerEvent(m, int64(14), int32(2), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	F_discardTempDb(m, v120)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	v446 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v446 {
		goto L103
	} else {
		goto L113
	}
L113:
	;
	F__serverLog(m, int32(2), int32(_a957), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	goto L103
L115:
	;
	v463 = int32(0)
	F_rioFreeConn(m, v15+int32(16), v463)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	v469 = F_connNonBlock(m, l0)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	v472 = F_connRecvTimeout(m, l0, int64(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	v482 = v463
	goto L38
}
func F_replicaProcessPsyncReply(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v189 int64
	_ = v189
	var v194 int32
	_ = v194
	var v216 int32
	_ = v216
	var v217 int64
	_ = v217
	var v224 int64
	_ = v224
	var v231 int64
	_ = v231
	var v238 int64
	_ = v238
	var v245 int64
	_ = v245
	var v248 int32
	_ = v248
	var v254 int64
	_ = v254
	var v257 int32
	_ = v257
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
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
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v358 int32
	_ = v358
	var v371 int32
	_ = v371
	var v376 int64
	_ = v376
	var v384 int64
	_ = v384
	var v392 int64
	_ = v392
	var v394 int64
	_ = v394
	var v398 int64
	_ = v398
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int64
	_ = v461
	var v463 int32
	_ = v463
	var v465 int64
	_ = v465
	var v466 int32
	_ = v466
	var v468 int64
	_ = v468
	var v469 int32
	_ = v469
	var v471 int64
	_ = v471
	var v472 int32
	_ = v472
	var v474 int64
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int64
	_ = v479
	var v503 int64
	_ = v503
	var v510 int64
	_ = v510
	var v516 int32
	_ = v516
	var v517 int64
	_ = v517
	var v523 int32
	_ = v523
	var v524 int64
	_ = v524
	var v530 int32
	_ = v530
	var v531 int64
	_ = v531
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int64
	_ = v541
	var v545 int32
	_ = v545
	var v549 int64
	_ = v549
	var v553 int64
	_ = v553
	var v557 int64
	_ = v557
	var v561 int64
	_ = v561
	var v564 int32
	_ = v564
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int64
	_ = v620
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v741 int32
	_ = v741
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v802 int32
	_ = v802
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v872 int32
	_ = v872
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	v14 = m.G0
	v16 = v14 - int32(352)
	m.G0 = v16
	v22 = *(*int32)(unsafe.Add(mBase, _consts[551]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+100))
	v28 = m.T0[v27].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v16+int32(96), int32(256), base.I64_extend_i32_s(v22*int32(1000)))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v16 + int32(352)
	return v885
L2:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(-1)))))
	switch v73 & int32(7) {
	case 0:
		goto L23
	case 1:
		goto L22
	case 2:
		goto L21
	case 3:
		goto L20
	case 4:
		goto L19
	default:
		goto L17
	}
L3:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+84))
	v59 = m.T0[v58].(func(*base.Module, int32, int32) int32)(m, l0, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L13
	}
L4:
	;
	v47 = int32(_a69)
	v49 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	*(*int64)(unsafe.Add(mBase, _consts[523])) = v49
	v53 = F_sdsnew(m, v16+int32(96))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L11
	}
L5:
	;
	return int32(0)
L6:
	;
	if v28 != int32(-1) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v35 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+88))
	v40 = m.T0[v39].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v40
	F__serverLog(m, int32(3), int32(_a1010), v16)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L3
L11:
	;
	if v53 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	goto L3
L13:
	;
	v61 = int32(5)
	v63 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v63 {
		v885 = v61
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F__serverLog(m, int32(3), int32(_a1011), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v885 = v61
	goto L1
L16:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+84))
	v98 = m.T0[v97].(func(*base.Module, int32, int32) int32)(m, l0, int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L26
	}
L17:
	;
	F_sdsfree(m, v53)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L25
	}
L18:
	;
	if v90 != 0 {
		goto L16
	} else {
		goto L24
	}
L19:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(-17))))
	v90 = v89
	goto L18
L20:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(-9))))
	v90 = v86
	goto L18
L21:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53+int32(-5)))))
	v90 = v83
	goto L18
L22:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(-3)))))
	v90 = v80
	goto L18
L23:
	;
	v90 = int32(base.Ui32(v73) >> (uint(int32(3)) % 32))
	goto L18
L24:
	;
	goto L17
L25:
	;
	v885 = int32(1)
	goto L1
L26:
	;
	v100 = int32(_a1012)
	goto L29
L27:
	;
	v278 = int32(_a1013)
	goto L68
L28:
	;
	if v134-v139 != 0 {
		goto L27
	} else {
		goto L41
	}
L29:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v105 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	goto L28
L32:
	;
	v107 = v53
	v108 = v100
	v109 = int32(11)
	v110 = v105
	goto L35
L33:
	;
	v134 = int32(0)
	v135 = v100
	goto L31
L34:
	;
	v134 = v131 & int32(255)
	v135 = v129
	goto L31
L35:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v110&int32(255) != v114 {
		v129 = v108
		v131 = v110
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v129 = v123
	v131 = int32(0)
	goto L34
L37:
	;
	if v114 == int32(0) {
		v129 = v108
		v131 = v110
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v119 = v109 + int32(-1)
	if v119 == int32(0) {
		v129 = v108
		v131 = v110
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v122 = int32(1)
	v123 = v108 + v122
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	if v124 != 0 {
		v107 = v107 + v122
		v108 = v123
		v109 = v119
		v110 = v124
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	v147 = int32(32)
	v148 = F___strchrnul(m, v53, v147)
	mBase = m.M
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v150 == v147 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	F_sdsfree(m, v53)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L5
	} else {
		goto L65
	}
L43:
	;
	v216 = int32(_a69)
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v158)))
	*(*int64)(unsafe.Add(mBase, _consts[579])) = v217
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v154+int32(33))))
	*(*int64)(unsafe.Add(mBase, _consts[580])) = v224
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v154+int32(25))))
	*(*int64)(unsafe.Add(mBase, _consts[581])) = v231
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v154+int32(17))))
	*(*int64)(unsafe.Add(mBase, _consts[582])) = v238
	v245 = *(*int64)(unsafe.Add(mBase, uint32(v154+int32(9))))
	*(*int64)(unsafe.Add(mBase, _consts[583])) = v245
	v248 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[584])) = uint8(v248)
	v254 = F_strtox_2(m, v170, v248, int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L62
L44:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v180 {
		goto L59
	} else {
		goto L60
	}
L45:
	;
	if v154 == int32(0) {
		goto L44
	} else {
		goto L49
	}
L46:
	;
	v154 = v148
	goto L48
L47:
	;
	v154 = int32(0)
	goto L48
L48:
	;
	goto L45
L49:
	;
	v158 = v154 + int32(1)
	v159 = int32(32)
	v160 = F___strchrnul(m, v158, v159)
	mBase = m.M
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if v162 == v159 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v166 == int32(0) {
		goto L44
	} else {
		goto L54
	}
L51:
	;
	v166 = v160
	goto L53
L52:
	;
	v166 = int32(0)
	goto L53
L53:
	;
	goto L50
L54:
	;
	v170 = v166 + int32(1)
	if v166 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v172 = v170
	goto L57
L56:
	;
	v172 = int32(0)
	goto L57
L57:
	;
	if v158-v172 == int32(-41) {
		goto L43
	} else {
		goto L58
	}
L58:
	;
	goto L44
L59:
	;
	v189 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[579])) = v189
	v194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[584])) = uint8(v194)
	*(*int64)(unsafe.Add(mBase, _consts[580])) = v189
	*(*int64)(unsafe.Add(mBase, _consts[581])) = v189
	*(*int64)(unsafe.Add(mBase, _consts[582])) = v189
	*(*int64)(unsafe.Add(mBase, _consts[583])) = v189
	goto L42
L60:
	;
	F__serverLog(m, int32(3), int32(_a1014), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	*(*int64)(unsafe.Add(mBase, _consts[574])) = v254
	v257 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v257 {
		goto L42
	} else {
		goto L63
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(_a949)
	F__serverLog(m, int32(2), int32(_a1015), v16+int32(16))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	goto L42
L65:
	;
	v885 = int32(3)
	goto L1
L66:
	;
	v644 = int32(_a1016)
	goto L121
L67:
	;
	if v312-v317 != 0 {
		goto L66
	} else {
		goto L80
	}
L68:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v283 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
	goto L67
L71:
	;
	v285 = v53
	v286 = v278
	v287 = int32(9)
	v288 = v283
	goto L74
L72:
	;
	v312 = int32(0)
	v313 = v278
	goto L70
L73:
	;
	v312 = v309 & int32(255)
	v313 = v307
	goto L70
L74:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v288&int32(255) != v292 {
		v307 = v286
		v309 = v288
		goto L73
	} else {
		goto L76
	}
L75:
	;
	v307 = v301
	v309 = int32(0)
	goto L73
L76:
	;
	if v292 == int32(0) {
		v307 = v286
		v309 = v288
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v297 = v287 + int32(-1)
	if v297 == int32(0) {
		v307 = v286
		v309 = v288
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v300 = int32(1)
	v301 = v286 + v300
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+1)))
	if v302 != 0 {
		v285 = v285 + v300
		v286 = v301
		v287 = v297
		v288 = v302
		goto L74
	} else {
		goto L79
	}
L79:
	;
	goto L75
L80:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _consts[572]))
	if v326 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v333 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	F_sdsfree(m, v53)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	v885 = int32(2)
	goto L1
L84:
	;
	v344 = v53 + int32(10)
	v348 = v53 + int32(9)
	goto L88
L85:
	;
	F__serverLog(m, int32(2), int32(_a1017), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	if v348-v344 != int32(40) {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348))))
	if base.Ui32(int32(13)) < base.Ui32(v358) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v348 = v348 + int32(1)
	goto L88
L91:
	;
	if int32(1)<<(uint(v358)%32)&int32(9217) != 0 {
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	F_sdsfree(m, v53)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L5
	} else {
		goto L108
	}
L94:
	;
	v371 = v16 + int32(96)
	v376 = *(*int64)(unsafe.Add(mBase, uint32(v53+int32(42))))
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(128)))) = v376
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v53+int32(34))))
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(120)))) = v384
	v392 = *(*int64)(unsafe.Add(mBase, uint32(v53+int32(26))))
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(112)))) = v392
	v394 = *(*int64)(unsafe.Add(mBase, uint32(v344)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v394
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v53+int32(18))))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+104)) = v398
	v400 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+136)) = uint8(v400)
	v405 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)+104))
	v408 = v406 + int32(104)
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408))))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371))))
	if v412 == v400 {
		v435 = v411
		v436 = v412
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v436-v435&int32(255) == int32(0) {
		goto L93
	} else {
		goto L103
	}
L96:
	;
	goto L95
L97:
	;
	if v412 != v411&int32(255) {
		v435 = v411
		v436 = v412
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v418 = v371
	v419 = v408
	goto L99
L99:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+1)))
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+1)))
	if v423 == int32(0) {
		v435 = v422
		v436 = v423
		goto L96
	} else {
		goto L101
	}
L100:
	;
	v435 = v422
	v436 = v423
	goto L96
L101:
	;
	v426 = int32(1)
	if v423 == v422&int32(255) {
		v418 = v418 + v426
		v419 = v419 + v426
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v443 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v443 {
		v458 = v406
		v459 = v405
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v460 = int32(_a69)
	v461 = *(*int64)(unsafe.Add(mBase, uint32(v458)+104))
	*(*int64)(unsafe.Add(mBase, _consts[512])) = v461
	v463 = int32(112)
	v465 = *(*int64)(unsafe.Add(mBase, uint32(v458+v463)))
	v466 = int32(120)
	v468 = *(*int64)(unsafe.Add(mBase, uint32(v458+v466)))
	v469 = int32(128)
	v471 = *(*int64)(unsafe.Add(mBase, uint32(v458+v469)))
	v472 = int32(136)
	v474 = *(*int64)(unsafe.Add(mBase, uint32(v458+v472)))
	v475 = int32(144)
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458+v475))))
	v479 = *(*int64)(unsafe.Add(mBase, uint32(v16)+96))
	*(*int64)(unsafe.Add(mBase, _consts[528])) = v479
	*(*uint8)(unsafe.Add(mBase, _consts[513])) = uint8(v477)
	*(*int64)(unsafe.Add(mBase, _consts[514])) = v474
	*(*int64)(unsafe.Add(mBase, _consts[515])) = v471
	*(*int64)(unsafe.Add(mBase, _consts[516])) = v468
	*(*int64)(unsafe.Add(mBase, _consts[517])) = v465
	v503 = *(*int64)(unsafe.Add(mBase, _consts[31]))
	*(*int64)(unsafe.Add(mBase, _consts[508])) = v503 + int64(1)
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v16)+104))
	*(*int64)(unsafe.Add(mBase, _consts[529])) = v510
	v516 = v16 + v463
	v517 = *(*int64)(unsafe.Add(mBase, uint32(v516)))
	*(*int64)(unsafe.Add(mBase, _consts[530])) = v517
	v523 = v16 + v466
	v524 = *(*int64)(unsafe.Add(mBase, uint32(v523)))
	*(*int64)(unsafe.Add(mBase, _consts[531])) = v524
	v530 = v16 + v469
	v531 = *(*int64)(unsafe.Add(mBase, uint32(v530)))
	*(*int64)(unsafe.Add(mBase, _consts[532])) = v531
	v537 = v16 + v472
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537))))
	*(*uint8)(unsafe.Add(mBase, _consts[511])) = uint8(v538)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v459)+104))
	v541 = *(*int64)(unsafe.Add(mBase, uint32(v16)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v540)+104)) = v541
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537))))
	*(*uint8)(unsafe.Add(mBase, uint32(v540+v475))) = uint8(v545)
	v549 = *(*int64)(unsafe.Add(mBase, uint32(v530)))
	*(*int64)(unsafe.Add(mBase, uint32(v540+v472))) = v549
	v553 = *(*int64)(unsafe.Add(mBase, uint32(v523)))
	*(*int64)(unsafe.Add(mBase, uint32(v540+v469))) = v553
	v557 = *(*int64)(unsafe.Add(mBase, uint32(v516)))
	*(*int64)(unsafe.Add(mBase, uint32(v540+v466))) = v557
	v561 = *(*int64)(unsafe.Add(mBase, uint32(v16)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v540+v463))) = v561
	F_disconnectReplicas(m)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L5
	} else {
		goto L107
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v16 + int32(96)
	F__serverLog(m, int32(2), int32(_a1018), v16+int32(32))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	v456 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)+104))
	v458 = v457
	v459 = v456
	goto L104
L107:
	;
	goto L93
L108:
	;
	v577 = int32(_a69)
	v578 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	v580 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[134])) = v580
	*(*int32)(unsafe.Add(mBase, _consts[133])) = v578
	*(*int32)(unsafe.Add(mBase, uint32(v578)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v578
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v578)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v578)+200)) = v586 & int32(-65)
	v591 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v591)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v591)+200)) = v592 & int32(-1025)
	v597 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	*(*int32)(unsafe.Add(mBase, uint32(v597)+328)) = v580
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v597)+204))
	goto L110
L109:
	;
	v617 = int32(_a69)
	v618 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v620 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	*(*int64)(unsafe.Add(mBase, uint32(v618)+88)) = v620
	*(*int64)(unsafe.Add(mBase, _consts[519])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[518])) = int32(14)
	v629 = int32(0)
	F_moduleFireServerEvent(m, int64(7), v629, v629)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L5
	} else {
		goto L113
	}
L110:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v597)+204)) = v602&int32(-25165825) | int32(_a14) | int32(16777216)
	goto L109
L113:
	;
	v634 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	F_linkClient(m, v634)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	F_replicationSteadyStateInit(m)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	v639 = int32(2)
	v641 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	if v641 != 0 {
		v885 = v639
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_createReplicationBacklog(m)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	v885 = v639
	goto L1
L118:
	;
	v754 = int32(_a1019)
	goto L154
L119:
	;
	v741 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v741 {
		goto L148
	} else {
		goto L149
	}
L120:
	;
	if v678-v683 == int32(0) {
		goto L119
	} else {
		goto L133
	}
L121:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v649 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679))))
	goto L120
L124:
	;
	v651 = v53
	v652 = v644
	v653 = int32(13)
	v654 = v649
	goto L127
L125:
	;
	v678 = int32(0)
	v679 = v644
	goto L123
L126:
	;
	v678 = v675 & int32(255)
	v679 = v673
	goto L123
L127:
	;
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652))))
	if v654&int32(255) != v658 {
		v673 = v652
		v675 = v654
		goto L126
	} else {
		goto L129
	}
L128:
	;
	v673 = v667
	v675 = int32(0)
	goto L126
L129:
	;
	if v658 == int32(0) {
		v673 = v652
		v675 = v654
		goto L126
	} else {
		goto L130
	}
L130:
	;
	v663 = v653 + int32(-1)
	if v663 == int32(0) {
		v673 = v652
		v675 = v654
		goto L126
	} else {
		goto L131
	}
L131:
	;
	v666 = int32(1)
	v667 = v652 + v666
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651)+1)))
	if v668 != 0 {
		v651 = v651 + v666
		v652 = v667
		v653 = v663
		v654 = v668
		goto L127
	} else {
		goto L132
	}
L132:
	;
	goto L128
L133:
	;
	v693 = int32(_a1020)
	goto L135
L134:
	;
	if v727-v732 != 0 {
		goto L118
	} else {
		goto L147
	}
L135:
	;
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v698 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728))))
	goto L134
L138:
	;
	v700 = v53
	v701 = v693
	v702 = int32(8)
	v703 = v698
	goto L141
L139:
	;
	v727 = int32(0)
	v728 = v693
	goto L137
L140:
	;
	v727 = v724 & int32(255)
	v728 = v722
	goto L137
L141:
	;
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701))))
	if v703&int32(255) != v707 {
		v722 = v701
		v724 = v703
		goto L140
	} else {
		goto L143
	}
L142:
	;
	v722 = v716
	v724 = int32(0)
	goto L140
L143:
	;
	if v707 == int32(0) {
		v722 = v701
		v724 = v703
		goto L140
	} else {
		goto L144
	}
L144:
	;
	v712 = v702 + int32(-1)
	if v712 == int32(0) {
		v722 = v701
		v724 = v703
		goto L140
	} else {
		goto L145
	}
L145:
	;
	v715 = int32(1)
	v716 = v701 + v715
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700)+1)))
	if v717 != 0 {
		v700 = v700 + v715
		v701 = v716
		v702 = v712
		v703 = v717
		goto L141
	} else {
		goto L146
	}
L146:
	;
	goto L142
L147:
	;
	goto L119
L148:
	;
	F_sdsfree(m, v53)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L5
	} else {
		goto L151
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v53
	F__serverLog(m, int32(2), int32(_a1021), v16+int32(48))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L5
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v885 = int32(5)
	goto L1
L152:
	;
	v814 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v815 = int32(_a1022)
	goto L174
L153:
	;
	if v788-v793 != 0 {
		goto L152
	} else {
		goto L166
	}
L154:
	;
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v759 != 0 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v789))))
	goto L153
L157:
	;
	v761 = v53
	v762 = v754
	v763 = int32(16)
	v764 = v759
	goto L160
L158:
	;
	v788 = int32(0)
	v789 = v754
	goto L156
L159:
	;
	v788 = v785 & int32(255)
	v789 = v783
	goto L156
L160:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762))))
	if v764&int32(255) != v768 {
		v783 = v762
		v785 = v764
		goto L159
	} else {
		goto L162
	}
L161:
	;
	v783 = v777
	v785 = int32(0)
	goto L159
L162:
	;
	if v768 == int32(0) {
		v783 = v762
		v785 = v764
		goto L159
	} else {
		goto L163
	}
L163:
	;
	v773 = v763 + int32(-1)
	if v773 == int32(0) {
		v783 = v762
		v785 = v764
		goto L159
	} else {
		goto L164
	}
L164:
	;
	v776 = int32(1)
	v777 = v762 + v776
	v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761)+1)))
	if v778 != 0 {
		v761 = v761 + v776
		v762 = v777
		v763 = v773
		v764 = v778
		goto L160
	} else {
		goto L165
	}
L165:
	;
	goto L161
L166:
	;
	v802 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v802 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	F_sdsfree(m, v53)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L5
	} else {
		goto L170
	}
L168:
	;
	F__serverLog(m, int32(2), int32(_a1023), int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L5
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v885 = int32(6)
	goto L1
L171:
	;
	F_sdsfree(m, v53)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L5
	} else {
		goto L191
	}
L172:
	;
	if int32(2) < v814 {
		goto L171
	} else {
		goto L189
	}
L173:
	;
	if v849-v854 == int32(0) {
		goto L172
	} else {
		goto L186
	}
L174:
	;
	v820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v820 != 0 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850))))
	goto L173
L177:
	;
	v822 = v53
	v823 = v815
	v824 = int32(4)
	v825 = v820
	goto L180
L178:
	;
	v849 = int32(0)
	v850 = v815
	goto L176
L179:
	;
	v849 = v846 & int32(255)
	v850 = v844
	goto L176
L180:
	;
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v823))))
	if v825&int32(255) != v829 {
		v844 = v823
		v846 = v825
		goto L179
	} else {
		goto L182
	}
L181:
	;
	v844 = v838
	v846 = int32(0)
	goto L179
L182:
	;
	if v829 == int32(0) {
		v844 = v823
		v846 = v825
		goto L179
	} else {
		goto L183
	}
L183:
	;
	v834 = v824 + int32(-1)
	if v834 == int32(0) {
		v844 = v823
		v846 = v825
		goto L179
	} else {
		goto L184
	}
L184:
	;
	v837 = int32(1)
	v838 = v823 + v837
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+1)))
	if v839 != 0 {
		v822 = v822 + v837
		v823 = v838
		v824 = v834
		v825 = v839
		goto L180
	} else {
		goto L185
	}
L185:
	;
	goto L181
L186:
	;
	if int32(3) < v814 {
		goto L171
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v53
	F__serverLog(m, int32(3), int32(_a1024), v16+int32(80))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	goto L171
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v53
	F__serverLog(m, int32(2), int32(_a1025), v16+int32(64))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L5
	} else {
		goto L190
	}
L190:
	;
	goto L171
L191:
	;
	v885 = int32(4)
	goto L1
}
func F_replicaPutOnline(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	v5 = m.G0
	v6 = int32(32)
	v7 = v5 - v6
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
	if v10&v6 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(9)
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v36 = int32(_a69)
		v37 = *(*int64)(unsafe.Add(mBase, _consts[37]))
		*(*int64)(unsafe.Add(mBase, uint32(v35)+80)) = v37
		v40 = *(*int32)(unsafe.Add(mBase, _consts[186]))
		if v40 == int32(0) {
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, _consts[187]))
			if v44 == int32(0) {
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, _consts[188]))
				v50 = v7 + int32(24)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
				*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v50))) = v51
				v55 = int32(0)
				v57 = v7 + int32(24)
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
				if v59 == v55 {
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v59+base.B2i32(v62 == int32(0))<<(uint(int32(2))%32))))
					*(*int32)(unsafe.Add(mBase, uint32(v57))) = v68
				}
				if v59 == int32(0) {
					v107 = v55
				} else {
					v74 = v59
					v75 = v55
					for {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+104))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
						if v78 != int32(9) {
							v89 = v75
						} else {
							v81 = int32(_a69)
							v82 = *(*int64)(unsafe.Add(mBase, _consts[37]))
							v83 = *(*int64)(unsafe.Add(mBase, uint32(v77)+80))
							v86 = int64(*(*int32)(unsafe.Add(mBase, _consts[187])))
							v89 = v75 + base.B2i32(v82-v83 <= v86)
						}
						v91 = v7 + int32(24)
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
						if v93 == int32(0) {
						} else {
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v93+base.B2i32(v96 == int32(0))<<(uint(int32(2))%32))))
							*(*int32)(unsafe.Add(mBase, uint32(v91))) = v102
						}
						if v93 != 0 {
							v74 = v93
							v75 = v89
							continue
						} else {
							break
						}
						break
					}
					v107 = v89
				}
				*(*int32)(unsafe.Add(mBase, _consts[189])) = v107
			}
		}
		v115 = int32(0)
		F_moduleFireServerEvent(m, int64(6), v115, v115)
		mBase = m.M
		v118 = m.ExcPending
		if v118 != 0 {
			return int32(0)
		} else {
			v119 = int32(1)
			v121 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			if int32(2) < v121 {
				v133 = v119
				m.G0 = v7 + int32(32)
				return v133
			} else {
				v124 = F_replicationGetReplicaName(m, l0)
				mBase = m.M
				v125 = m.ExcPending
				if v125 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v124
					F__serverLog(m, int32(2), int32(_a933), v7)
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return int32(0)
					} else {
						v133 = v119
						m.G0 = v7 + int32(32)
						return v133
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(10)
		v17 = int32(0)
		v19 = *(*int32)(unsafe.Add(mBase, _consts[15]))
		if int32(2) < v19 {
			v133 = v17
			m.G0 = v7 + int32(32)
			return v133
		} else {
			v22 = F_replicationGetReplicaName(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v22
				F__serverLog(m, int32(2), int32(_a934), v7+int32(16))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v133 = v17
					m.G0 = v7 + int32(32)
					return v133
				}
			}
		}
	}
}
func F_replicaRdbVersion(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+160)))
	if v3&int32(1) != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v2)+156))
		if v8 <= int32(589823) {
			if v8 < int32(459264) {
				return int32(11)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _consts[503]))
				return v17
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[504]))
			return v12
		}
	} else {
		return int32(80)
	}
}
func F_replicaReceiveRDBFromPrimaryToDisk(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int64
	_ = v93
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v140 int64
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
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
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v207 int32
	_ = v207
	var v210 int64
	_ = v210
	var v218 int64
	_ = v218
	var v226 int64
	_ = v226
	var v228 int64
	_ = v228
	var v232 int64
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int64
	_ = v473
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int64
	_ = v502
	var v503 int64
	_ = v503
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v519 int64
	_ = v519
	var v520 int64
	_ = v520
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int64
	_ = v527
	var v528 int32
	_ = v528
	var v529 int64
	_ = v529
	var v531 int64
	_ = v531
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v564 int32
	_ = v564
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	v11 = m.G0
	v13 = v11 - int32(16544)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v16 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v595 = F_connNonBlock(m, l0)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L8
	} else {
		goto L146
	}
L2:
	;
	v577 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v577 {
		v589 = v112
		goto L1
	} else {
		goto L144
	}
L3:
	;
	F__serverAssert(m, int32(_a974), int32(_a913), int32(2716))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L8
	} else {
		goto L143
	}
L4:
	;
	v25 = F_connBlock(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v19 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[550]))
	goto L6
L6:
	;
	if base.B2i32(v20 != v19) == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	return
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[551]))
	v32 = F_connRecvTimeout(m, l0, base.I64_extend_i32_s(v28*int32(1000)))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[552])) = int32(1)
	goto L11
L11:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, _consts[553])))
	if v51&int32(1) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v74 != int32(-1) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v74 = F_tryReadBulkPayloadMetadata(m, l0, v13+int32(112), v13+int32(64), v13+int32(16496), v13+int32(16540), int32(_a975))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L17
	}
L14:
	;
	v56 = int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v58 {
		v589 = v56
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F__serverLog(m, int32(3), int32(_a976), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v589 = v56
	goto L1
L17:
	;
	if v74 == int32(-2) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	v90 = int32(_a69)
	v91 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v93 = *(*int64)(unsafe.Add(mBase, _consts[554]))
	if v93 != int64(0) {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v80 = int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v82 {
		v589 = v80
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F__serverLog(m, int32(3), int32(_a977), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v589 = v80
	goto L1
L23:
	;
	v112 = int32(1)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, _consts[553])))
	if v114&v112 != 0 {
		goto L2
	} else {
		goto L30
	}
L24:
	;
	if int32(2) < v91 {
		goto L23
	} else {
		goto L28
	}
L25:
	;
	if int32(2) < v91 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	F__serverLog(m, int32(2), int32(_a978), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v93
	F__serverLog(m, int32(2), int32(_a979), v13+int32(48))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	v128 = int64(0)
	goto L31
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[555])))
	if v131 != 0 {
		v142 = int32(16384)
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+76))
	v148 = m.T0[v147].(func(*base.Module, int32, int32, int32) int32)(m, l0, v13+int32(112), v142)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L8
	} else {
		goto L39
	}
L34:
	;
	v132 = int32(_a69)
	v133 = *(*int64)(unsafe.Add(mBase, _consts[554]))
	v135 = *(*int64)(unsafe.Add(mBase, _consts[556]))
	v136 = v133 - v135
	v137 = int64(16384)
	if v136 < v137 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v140 = v136
	goto L37
L36:
	;
	v140 = v137
	goto L37
L37:
	;
	v142 = base.I32_wrap_i64(v140)
	goto L33
L38:
	;
	v194 = int32(_a69)
	v196 = *(*int64)(unsafe.Add(mBase, _consts[557]))
	v197 = base.I64_extend_i32_u(v148)
	*(*int64)(unsafe.Add(mBase, _consts[557])) = v196 + v197
	if v131 != 0 {
		goto L53
	} else {
		goto L54
	}
L39:
	;
	if int32(0) < v148 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v152 != int32(3) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v167 = int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v169 {
		v589 = v167
		goto L1
	} else {
		goto L45
	}
L42:
	;
	v160 = F__emscripten_memset_bulkmem(m, v13+int32(112), base.I32_extend8_s(int32(0)), int32(16384))
	mBase = m.M
	goto L43
L43:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, _consts[553])))
	if v162&int32(1) == int32(0) {
		goto L31
	} else {
		goto L44
	}
L44:
	;
	goto L2
L45:
	;
	if v148 != int32(-1) {
		v179 = int32(_a980)
		goto L46
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v179
	F__serverLog(m, int32(3), int32(_a981), v13)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L8
	} else {
		goto L49
	}
L47:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+88))
	v177 = m.T0[v176].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	v179 = v177
	goto L46
L49:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v186 {
		v589 = v167
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F__serverLog(m, int32(3), int32(_a982), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v589 = v167
	goto L1
L52:
	;
	v471 = int32(_a69)
	v473 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	*(*int64)(unsafe.Add(mBase, _consts[523])) = v473
	v476 = *(*int32)(unsafe.Add(mBase, _consts[539]))
	v479 = F_write(m, v476, v13+int32(112), v148)
	mBase = m.M
	if v479 == v148 {
		goto L118
	} else {
		goto L119
	}
L53:
	;
	if base.Ui32(v148) < base.Ui32(int32(40)) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v470 = int32(0)
	goto L52
L55:
	;
	v400 = v13 + int32(16496)
	v402 = v13 + int32(64)
	v403 = int32(40)
	goto L106
L56:
	;
	v235 = v13 + int32(16496)
	v238 = v235 + v148
	v240 = int32(40) - v148
	if v235 == v238 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v207 = v13 + int32(72) + v148
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v207+int32(32))))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[558]))) = v210
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v207+int32(24))))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[559]))) = v218
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v207+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[560]))) = v226
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v207)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[561]))) = v228
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v207+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[562]))) = v232
	goto L55
L58:
	;
	if v148 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L59:
	;
	goto L58
L60:
	;
	v244 = v240 + v235
	if base.Ui32(int32(0)-v240<<(uint(int32(1))%32)) < base.Ui32(v238-v244) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v254 = (v238 ^ v235) & int32(3)
	if base.Ui32(v238) <= base.Ui32(v235) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v251 = F___memcpy(m, v235, v238, v240)
	mBase = m.M
	goto L58
L63:
	;
	if v360 == int32(0) {
		goto L59
	} else {
		goto L95
	}
L64:
	;
	if base.Ui32(v338) <= base.Ui32(int32(3)) {
		v359 = v337
		v360 = v338
		v361 = v339
		goto L63
	} else {
		goto L91
	}
L65:
	;
	if v254 != 0 {
		v320 = v240
		goto L75
	} else {
		goto L76
	}
L66:
	;
	if v254 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if v235&int32(3) != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v359 = v238
	v360 = v240
	v361 = v235
	goto L63
L69:
	;
	v261 = v238
	v262 = v240
	v263 = v235
	goto L71
L70:
	;
	v337 = v238
	v338 = v240
	v339 = v235
	goto L64
L71:
	;
	if v262 == int32(0) {
		goto L59
	} else {
		goto L73
	}
L73:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	*(*uint8)(unsafe.Add(mBase, uint32(v263))) = uint8(v267)
	v269 = int32(1)
	v270 = v261 + v269
	v272 = v262 + int32(-1)
	v274 = v263 + v269
	if v274&int32(3) == int32(0) {
		v337 = v270
		v338 = v272
		v339 = v274
		goto L64
	} else {
		goto L74
	}
L74:
	;
	v261 = v270
	v262 = v272
	v263 = v274
	goto L71
L75:
	;
	if v320 == int32(0) {
		goto L59
	} else {
		goto L87
	}
L76:
	;
	if v244&int32(3) == int32(0) {
		v300 = v240
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if base.Ui32(v300) <= base.Ui32(int32(3)) {
		v320 = v300
		goto L75
	} else {
		goto L83
	}
L78:
	;
	v285 = v240
	goto L79
L79:
	;
	if v285 == int32(0) {
		goto L59
	} else {
		goto L81
	}
L80:
	;
	v300 = v291
	goto L77
L81:
	;
	v291 = v285 + int32(-1)
	v292 = v235 + v291
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238+v291))))
	*(*uint8)(unsafe.Add(mBase, uint32(v292))) = uint8(v294)
	if v292&int32(3) != 0 {
		v285 = v291
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v307 = v300
	goto L84
L84:
	;
	v311 = v307 + int32(-4)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v238+v311)))
	*(*int32)(unsafe.Add(mBase, uint32(v235+v311))) = v314
	if base.Ui32(int32(3)) < base.Ui32(v311) {
		v307 = v311
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v320 = v311
	goto L75
L86:
	;
	goto L85
L87:
	;
	v327 = v320
	goto L88
L88:
	;
	v331 = v327 + int32(-1)
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238+v331))))
	*(*uint8)(unsafe.Add(mBase, uint32(v235+v331))) = uint8(v334)
	if v331 != 0 {
		v327 = v331
		goto L88
	} else {
		goto L90
	}
L90:
	;
	goto L59
L91:
	;
	v344 = v337
	v345 = v338
	v346 = v339
	goto L92
L92:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	*(*int32)(unsafe.Add(mBase, uint32(v346))) = v348
	v350 = int32(4)
	v351 = v344 + v350
	v353 = v346 + v350
	v355 = v345 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v355) {
		v344 = v351
		v345 = v355
		v346 = v353
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v359 = v351
	v360 = v355
	v361 = v353
	goto L63
L94:
	;
	goto L93
L95:
	;
	v366 = v359
	v367 = v360
	v368 = v361
	goto L96
L96:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	*(*uint8)(unsafe.Add(mBase, uint32(v368))) = uint8(v370)
	v372 = int32(1)
	v377 = v367 + int32(-1)
	if v377 != 0 {
		v366 = v366 + v372
		v367 = v377
		v368 = v368 + v372
		goto L96
	} else {
		goto L98
	}
L97:
	;
	goto L59
L98:
	;
	goto L97
L99:
	;
	goto L55
L100:
	;
	goto L99
L101:
	;
	v396 = F__emscripten_memcpy_bulkmem(m, v13+int32(16496)+v240, v13+int32(112), v148)
	mBase = m.M
	goto L100
L102:
	;
	v470 = base.B2i32(v467 == int32(0))
	goto L52
L103:
	;
	v467 = int32(0)
	goto L102
L104:
	;
	v439 = v434
	v440 = v435
	v441 = v436
	goto L114
L105:
	;
	if v424 == int32(0) {
		goto L103
	} else {
		goto L112
	}
L106:
	;
	if (v402|v400)&int32(3) != 0 {
		v434 = v400
		v435 = v402
		v436 = v403
		goto L104
	} else {
		goto L107
	}
L107:
	;
	v411 = v400
	v412 = v402
	v413 = v403
	goto L108
L108:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	if v416 != v417 {
		v434 = v411
		v435 = v412
		v436 = v413
		goto L104
	} else {
		goto L110
	}
L109:
	;
	goto L105
L110:
	;
	v419 = int32(4)
	v420 = v412 + v419
	v422 = v411 + v419
	v424 = v413 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v424) {
		v411 = v422
		v412 = v420
		v413 = v424
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v434 = v422
	v435 = v420
	v436 = v424
	goto L104
L113:
	;
	v467 = v444 - v445
	goto L102
L114:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439))))
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
	if v444 != v445 {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v447 = int32(1)
	v452 = v441 + int32(-1)
	if v452 == int32(0) {
		goto L103
	} else {
		goto L117
	}
L117:
	;
	v439 = v439 + v447
	v440 = v440 + v447
	v441 = v452
	goto L114
L118:
	;
	v500 = int32(_a69)
	v502 = *(*int64)(unsafe.Add(mBase, _consts[556]))
	v503 = v502 + v197
	*(*int64)(unsafe.Add(mBase, _consts[556])) = v503
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[555])))
	if base.B2i32(v505 != int32(0))&v470 != int32(1) {
		v520 = v503
		goto L130
	} else {
		goto L131
	}
L119:
	;
	v481 = int32(1)
	v483 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v483 {
		v589 = v481
		goto L1
	} else {
		goto L120
	}
L120:
	;
	if v479 != int32(-1) {
		v492 = int32(_a983)
		goto L121
	} else {
		goto L122
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v492
	F__serverLog(m, int32(3), int32(_a984), v13+int32(32))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L8
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v491 = F___strerror_l(m, v490, v490)
	mBase = m.M
	goto L124
L124:
	;
	v492 = v491
	goto L121
L125:
	;
	v589 = v481
	goto L1
L126:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, _consts[553])))
	if v555&int32(1) == int32(0) {
		v128 = v527
		goto L31
	} else {
		goto L142
	}
L127:
	;
	v589 = int32(0)
	goto L1
L128:
	;
	if v470 == int32(0) {
		goto L126
	} else {
		goto L141
	}
L129:
	;
	v536 = int32(1)
	v538 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v538 {
		v589 = v536
		goto L1
	} else {
		goto L137
	}
L130:
	;
	if v520 < v128+int64(8388608) {
		v527 = v128
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v512 = *(*int32)(unsafe.Add(mBase, _consts[539]))
	v515 = F_ftruncate(m, v512, v503+int64(-40))
	mBase = m.M
	if v515 == int32(-1) {
		goto L129
	} else {
		goto L132
	}
L132:
	;
	v519 = *(*int64)(unsafe.Add(mBase, _consts[556]))
	v520 = v519
	goto L130
L133:
	;
	if v505 != 0 {
		goto L128
	} else {
		goto L135
	}
L134:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _consts[539]))
	v526 = F_fsync(m, v525)
	mBase = m.M
	v527 = v520
	goto L133
L135:
	;
	v528 = int32(_a69)
	v529 = *(*int64)(unsafe.Add(mBase, _consts[556]))
	v531 = *(*int64)(unsafe.Add(mBase, _consts[554]))
	if (base.B2i32(v529 == v531)|v470)&int32(1) != 0 {
		goto L127
	} else {
		goto L136
	}
L136:
	;
	goto L126
L137:
	;
	goto L138
L138:
	;
	v542 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v543 = F___strerror_l(m, v542, v542)
	mBase = m.M
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v543
	F__serverLog(m, int32(3), int32(_a985), v13+int32(16))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L8
	} else {
		goto L140
	}
L140:
	;
	v589 = v536
	goto L1
L141:
	;
	goto L127
L142:
	;
	goto L2
L143:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F__serverLog(m, int32(3), int32(_a986), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L8
	} else {
		goto L145
	}
L145:
	;
	v589 = v112
	goto L1
L146:
	;
	v598 = F_connRecvTimeout(m, l0, int64(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L8
	} else {
		goto L147
	}
L147:
	;
	if l1 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v589 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, _consts[521])) = l0
	goto L148
L150:
	;
	*(*int32)(unsafe.Add(mBase, _consts[527])) = l0
	goto L148
L151:
	;
	*(*int32)(unsafe.Add(mBase, _consts[552])) = v628
	m.G0 = v13 + int32(16544)
	return
L152:
	;
	v619 = int32(2)
	if v619 < v607 {
		v628 = v619
		goto L151
	} else {
		goto L156
	}
L153:
	;
	v610 = int32(3)
	if v610 < v607 {
		v628 = v610
		goto L151
	} else {
		goto L154
	}
L154:
	;
	v613 = int32(3)
	F__serverLog(m, v613, int32(_a987), int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L8
	} else {
		goto L155
	}
L155:
	;
	v628 = v613
	goto L151
L156:
	;
	v622 = int32(2)
	F__serverLog(m, v622, int32(_a988), int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L8
	} else {
		goto L157
	}
L157:
	;
	v628 = v622
	goto L151
}
func F_replicaReceiveRDBFromPrimaryToMemory(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v23 int64
	_ = v23
	var v28 int64
	_ = v28
	var v33 int64
	_ = v33
	var v38 int64
	_ = v38
	var v43 int64
	_ = v43
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int64
	_ = v73
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16464)
	m.G0 = v7
	v9 = F_scriptIsTimedout(m)
	mBase = m.M
	v11 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	goto L2
L1:
	;
	m.G0 = v7 + int32(16464)
	return
L2:
	;
	if v9|v11 != v2 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = int32(0)
	v18 = *(*int64)(unsafe.Add(mBase, _consts[542]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(72)))) = v18
	v23 = *(*int64)(unsafe.Add(mBase, _consts[543]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(64)))) = v23
	v28 = *(*int64)(unsafe.Add(mBase, _consts[544]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(56)))) = v28
	v33 = *(*int64)(unsafe.Add(mBase, _consts[545]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(48)))) = v33
	v38 = *(*int64)(unsafe.Add(mBase, _consts[546]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v38
	v43 = *(*int64)(unsafe.Add(mBase, _consts[547]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(32)))) = v43
	v48 = *(*int64)(unsafe.Add(mBase, _consts[548]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v48
	v51 = *(*int64)(unsafe.Add(mBase, _consts[549]))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v51
	v54 = *(*int64)(unsafe.Add(mBase, _consts[537]))
	v55 = int64(1)
	v56 = v54 + v55
	if base.Ui64(v55) < base.Ui64(v56) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v123 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L5:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v114 {
		goto L4
	} else {
		goto L24
	}
L6:
	;
	switch base.I32_wrap_i64(v56) {
	default:
		goto L8
	case 1:
		goto L7
	}
L7:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v105 {
		goto L4
	} else {
		goto L22
	}
L8:
	;
	v68 = F_tryReadBulkPayloadMetadata(m, l0, v7+int32(80), int32(_a964), int32(_a968), int32(_a965), int32(_a969))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v93 {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	v83 = *(*int64)(unsafe.Add(mBase, _consts[537]))
	if v83 == int64(-1) {
		goto L1
	} else {
		goto L16
	}
L11:
	;
	v73 = *(*int64)(unsafe.Add(mBase, _consts[537]))
	if int64(-1) < v73 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	return
L13:
	;
	switch v68 + int32(2) {
	case 0:
		goto L10
	default:
		goto L9
	case 2:
		goto L11
	}
L14:
	;
	F__serverAssert(m, int32(_a972), int32(_a913), int32(2650))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	F__serverAssert(m, int32(_a970), int32(_a913), int32(2653))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	v102 = F_cancelReplicationHandshake(m, int32(1))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L12
	} else {
		goto L21
	}
L19:
	;
	F__serverLog(m, int32(3), int32(_a971), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L1
L22:
	;
	F__serverLog(m, int32(2), int32(_a973), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	goto L4
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v54
	F__serverLog(m, int32(2), int32(_a967), v7)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	goto L4
L26:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v129 != int32(1) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	F_stopAppendOnly(m)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+84))
	v137 = m.T0[v136].(func(*base.Module, int32, int32) int32)(m, l0, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L12
	} else {
		goto L32
	}
L30:
	;
	F_killRDBChild(m)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v145 = F_replicaLoadPrimaryRDBFromSocket(m, l0, v7+int32(80), int32(_a964), int32(_a965), v7+int32(16))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L12
	} else {
		goto L34
	}
L33:
	;
	F_replicaAfterLoadPrimaryRDB(m, l0, v7+int32(16), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L12
	} else {
		goto L40
	}
L34:
	;
	if v145 != int32(-1) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v150 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v159 = F_cancelReplicationHandshake(m, int32(1))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L12
	} else {
		goto L39
	}
L37:
	;
	F__serverLog(m, int32(3), int32(_a966), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L1
L40:
	;
	goto L1
}
func F_replicaSendPsyncCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	v4 = m.G0
	v6 = v4 - int32(160)
	m.G0 = v6
	v8 = int32(_a69)
	*(*int64)(unsafe.Add(mBase, _consts[574])) = int64(-1)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[572]))
	if v12 == int32(0) {
		v51 = *(*int32)(unsafe.Add(mBase, _consts[134]))
		if v51 == int32(0) {
			v84 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			if int32(2) < v84 {
				v92 = int32(0)
				v93 = int32(*(*uint8)(unsafe.Add(mBase, _consts[575])))
				*(*uint8)(unsafe.Add(mBase, uint32(v6)+130)) = uint8(v93)
				v96 = int32(*(*uint16)(unsafe.Add(mBase, _consts[576])))
				*(*uint16)(unsafe.Add(mBase, uint32(v6)+128)) = uint16(v96)
				v99 = int32(_a176)
				v101 = *(*int32)(unsafe.Add(mBase, _consts[577]))
				if v101 != int32(2) {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v99
					*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(_a1002)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = v6 + int32(128)
					v130 = F_sendCommand(m, l0, v6+int32(48))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return int32(0)
					} else {
						v132 = v130
						if v132 != 0 {
							v135 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							if int32(3) < v135 {
								F_sdsfree(m, v132)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return int32(0)
								} else {
									v145 = int32(0)
									v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
									v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
										return int32(0)
									} else {
										v151 = v145
										m.G0 = v6 + int32(160)
										return v151
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v132
								F__serverLog(m, int32(3), int32(_a1003), v6)
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return int32(0)
								} else {
									F_sdsfree(m, v132)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										v145 = int32(0)
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
										v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int32(0)
										} else {
											v151 = v145
											m.G0 = v6 + int32(160)
											return v151
										}
									}
								}
							}
						} else {
							v151 = int32(1)
							m.G0 = v6 + int32(160)
							return v151
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6+int32(32)))) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a1004)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v99
					*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a1002)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v6 + int32(128)
					v118 = F_sendCommand(m, l0, v6+int32(16))
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						v132 = v118
						if v132 != 0 {
							v135 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							if int32(3) < v135 {
								F_sdsfree(m, v132)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return int32(0)
								} else {
									v145 = int32(0)
									v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
									v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
										return int32(0)
									} else {
										v151 = v145
										m.G0 = v6 + int32(160)
										return v151
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v132
								F__serverLog(m, int32(3), int32(_a1003), v6)
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return int32(0)
								} else {
									F_sdsfree(m, v132)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										v145 = int32(0)
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
										v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int32(0)
										} else {
											v151 = v145
											m.G0 = v6 + int32(160)
											return v151
										}
									}
								}
							}
						} else {
							v151 = int32(1)
							m.G0 = v6 + int32(160)
							return v151
						}
					}
				}
			} else {
				F__serverLog(m, int32(2), int32(_a1005), int32(0))
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					v92 = int32(0)
					v93 = int32(*(*uint8)(unsafe.Add(mBase, _consts[575])))
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+130)) = uint8(v93)
					v96 = int32(*(*uint16)(unsafe.Add(mBase, _consts[576])))
					*(*uint16)(unsafe.Add(mBase, uint32(v6)+128)) = uint16(v96)
					v99 = int32(_a176)
					v101 = *(*int32)(unsafe.Add(mBase, _consts[577]))
					if v101 != int32(2) {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v99
						*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(_a1002)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = v6 + int32(128)
						v130 = F_sendCommand(m, l0, v6+int32(48))
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return int32(0)
						} else {
							v132 = v130
							if v132 != 0 {
								v135 = *(*int32)(unsafe.Add(mBase, _consts[15]))
								if int32(3) < v135 {
									F_sdsfree(m, v132)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										v145 = int32(0)
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
										v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int32(0)
										} else {
											v151 = v145
											m.G0 = v6 + int32(160)
											return v151
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v132
									F__serverLog(m, int32(3), int32(_a1003), v6)
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return int32(0)
									} else {
										F_sdsfree(m, v132)
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											v145 = int32(0)
											v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
											v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return int32(0)
											} else {
												v151 = v145
												m.G0 = v6 + int32(160)
												return v151
											}
										}
									}
								}
							} else {
								v151 = int32(1)
								m.G0 = v6 + int32(160)
								return v151
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6+int32(32)))) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a1004)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v99
						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a1002)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v6 + int32(128)
						v118 = F_sendCommand(m, l0, v6+int32(16))
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							v132 = v118
							if v132 != 0 {
								v135 = *(*int32)(unsafe.Add(mBase, _consts[15]))
								if int32(3) < v135 {
									F_sdsfree(m, v132)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										v145 = int32(0)
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
										v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int32(0)
										} else {
											v151 = v145
											m.G0 = v6 + int32(160)
											return v151
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v132
									F__serverLog(m, int32(3), int32(_a1003), v6)
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return int32(0)
									} else {
										F_sdsfree(m, v132)
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											v145 = int32(0)
											v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
											v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return int32(0)
											} else {
												v151 = v145
												m.G0 = v6 + int32(160)
												return v151
											}
										}
									}
								}
							} else {
								v151 = int32(1)
								m.G0 = v6 + int32(160)
								return v151
							}
						}
					}
				}
			}
		} else {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+104))
			v55 = *(*int64)(unsafe.Add(mBase, uint32(v54)+48))
			*(*int64)(unsafe.Add(mBase, uint32(v6)+80)) = v55 + int64(1)
			v65 = F_snprintf(m, v6+int32(128), int32(32), int32(_a1006), v6+int32(80))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				v68 = v54 + int32(104)
				v70 = *(*int32)(unsafe.Add(mBase, _consts[15]))
				if int32(2) < v70 {
					v99 = v68
					v101 = *(*int32)(unsafe.Add(mBase, _consts[577]))
					if v101 != int32(2) {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v99
						*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(_a1002)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = v6 + int32(128)
						v130 = F_sendCommand(m, l0, v6+int32(48))
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return int32(0)
						} else {
							v132 = v130
							if v132 != 0 {
								v135 = *(*int32)(unsafe.Add(mBase, _consts[15]))
								if int32(3) < v135 {
									F_sdsfree(m, v132)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										v145 = int32(0)
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
										v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int32(0)
										} else {
											v151 = v145
											m.G0 = v6 + int32(160)
											return v151
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v132
									F__serverLog(m, int32(3), int32(_a1003), v6)
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return int32(0)
									} else {
										F_sdsfree(m, v132)
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											v145 = int32(0)
											v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
											v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return int32(0)
											} else {
												v151 = v145
												m.G0 = v6 + int32(160)
												return v151
											}
										}
									}
								}
							} else {
								v151 = int32(1)
								m.G0 = v6 + int32(160)
								return v151
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6+int32(32)))) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a1004)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v99
						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a1002)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v6 + int32(128)
						v118 = F_sendCommand(m, l0, v6+int32(16))
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							v132 = v118
							if v132 != 0 {
								v135 = *(*int32)(unsafe.Add(mBase, _consts[15]))
								if int32(3) < v135 {
									F_sdsfree(m, v132)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										v145 = int32(0)
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
										v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int32(0)
										} else {
											v151 = v145
											m.G0 = v6 + int32(160)
											return v151
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v132
									F__serverLog(m, int32(3), int32(_a1003), v6)
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return int32(0)
									} else {
										F_sdsfree(m, v132)
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											v145 = int32(0)
											v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
											v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return int32(0)
											} else {
												v151 = v145
												m.G0 = v6 + int32(160)
												return v151
											}
										}
									}
								}
							} else {
								v151 = int32(1)
								m.G0 = v6 + int32(160)
								return v151
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v68
					*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = v6 + int32(128)
					F__serverLog(m, int32(2), int32(_a1007), v6+int32(64))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						v99 = v68
						v101 = *(*int32)(unsafe.Add(mBase, _consts[577]))
						if v101 != int32(2) {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v99
							*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(_a1002)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = v6 + int32(128)
							v130 = F_sendCommand(m, l0, v6+int32(48))
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int32(0)
							} else {
								v132 = v130
								if v132 != 0 {
									v135 = *(*int32)(unsafe.Add(mBase, _consts[15]))
									if int32(3) < v135 {
										F_sdsfree(m, v132)
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											v145 = int32(0)
											v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
											v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return int32(0)
											} else {
												v151 = v145
												m.G0 = v6 + int32(160)
												return v151
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6))) = v132
										F__serverLog(m, int32(3), int32(_a1003), v6)
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int32(0)
										} else {
											F_sdsfree(m, v132)
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return int32(0)
											} else {
												v145 = int32(0)
												v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
												v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
												mBase = m.M
												v150 = m.ExcPending
												if v150 != 0 {
													return int32(0)
												} else {
													v151 = v145
													m.G0 = v6 + int32(160)
													return v151
												}
											}
										}
									}
								} else {
									v151 = int32(1)
									m.G0 = v6 + int32(160)
									return v151
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6+int32(32)))) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a1004)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v99
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a1002)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v6 + int32(128)
							v118 = F_sendCommand(m, l0, v6+int32(16))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								v132 = v118
								if v132 != 0 {
									v135 = *(*int32)(unsafe.Add(mBase, _consts[15]))
									if int32(3) < v135 {
										F_sdsfree(m, v132)
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											v145 = int32(0)
											v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
											v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return int32(0)
											} else {
												v151 = v145
												m.G0 = v6 + int32(160)
												return v151
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6))) = v132
										F__serverLog(m, int32(3), int32(_a1003), v6)
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int32(0)
										} else {
											F_sdsfree(m, v132)
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return int32(0)
											} else {
												v145 = int32(0)
												v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
												v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
												mBase = m.M
												v150 = m.ExcPending
												if v150 != 0 {
													return int32(0)
												} else {
													v151 = v145
													m.G0 = v6 + int32(160)
													return v151
												}
											}
										}
									}
								} else {
									v151 = int32(1)
									m.G0 = v6 + int32(160)
									return v151
								}
							}
						}
					}
				}
			}
		}
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, _consts[578]))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+112)) = v16 + int64(1)
		v26 = F_snprintf(m, v6+int32(128), int32(32), int32(_a1006), v6+int32(112))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			if v31 <= int32(2) {
				v39 = int32(_a1008)
				*(*int32)(unsafe.Add(mBase, uint32(v6)+96)) = v39
				*(*int32)(unsafe.Add(mBase, uint32(v6)+100)) = v6 + int32(128)
				F__serverLog(m, int32(2), int32(_a1009), v6+int32(96))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v99 = v39
					v101 = *(*int32)(unsafe.Add(mBase, _consts[577]))
					if v101 != int32(2) {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v99
						*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(_a1002)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = v6 + int32(128)
						v130 = F_sendCommand(m, l0, v6+int32(48))
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return int32(0)
						} else {
							v132 = v130
							if v132 != 0 {
								v135 = *(*int32)(unsafe.Add(mBase, _consts[15]))
								if int32(3) < v135 {
									F_sdsfree(m, v132)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										v145 = int32(0)
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
										v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int32(0)
										} else {
											v151 = v145
											m.G0 = v6 + int32(160)
											return v151
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v132
									F__serverLog(m, int32(3), int32(_a1003), v6)
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return int32(0)
									} else {
										F_sdsfree(m, v132)
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											v145 = int32(0)
											v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
											v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return int32(0)
											} else {
												v151 = v145
												m.G0 = v6 + int32(160)
												return v151
											}
										}
									}
								}
							} else {
								v151 = int32(1)
								m.G0 = v6 + int32(160)
								return v151
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6+int32(32)))) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a1004)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v99
						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a1002)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v6 + int32(128)
						v118 = F_sendCommand(m, l0, v6+int32(16))
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							v132 = v118
							if v132 != 0 {
								v135 = *(*int32)(unsafe.Add(mBase, _consts[15]))
								if int32(3) < v135 {
									F_sdsfree(m, v132)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										v145 = int32(0)
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
										v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int32(0)
										} else {
											v151 = v145
											m.G0 = v6 + int32(160)
											return v151
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v132
									F__serverLog(m, int32(3), int32(_a1003), v6)
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return int32(0)
									} else {
										F_sdsfree(m, v132)
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											v145 = int32(0)
											v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
											v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return int32(0)
											} else {
												v151 = v145
												m.G0 = v6 + int32(160)
												return v151
											}
										}
									}
								}
							} else {
								v151 = int32(1)
								m.G0 = v6 + int32(160)
								return v151
							}
						}
					}
				}
			} else {
				v99 = int32(_a1008)
				v101 = *(*int32)(unsafe.Add(mBase, _consts[577]))
				if v101 != int32(2) {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v99
					*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(_a1002)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = v6 + int32(128)
					v130 = F_sendCommand(m, l0, v6+int32(48))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return int32(0)
					} else {
						v132 = v130
						if v132 != 0 {
							v135 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							if int32(3) < v135 {
								F_sdsfree(m, v132)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return int32(0)
								} else {
									v145 = int32(0)
									v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
									v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
										return int32(0)
									} else {
										v151 = v145
										m.G0 = v6 + int32(160)
										return v151
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v132
								F__serverLog(m, int32(3), int32(_a1003), v6)
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return int32(0)
								} else {
									F_sdsfree(m, v132)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										v145 = int32(0)
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
										v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int32(0)
										} else {
											v151 = v145
											m.G0 = v6 + int32(160)
											return v151
										}
									}
								}
							}
						} else {
							v151 = int32(1)
							m.G0 = v6 + int32(160)
							return v151
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6+int32(32)))) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a1004)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v99
					*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a1002)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v6 + int32(128)
					v118 = F_sendCommand(m, l0, v6+int32(16))
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						v132 = v118
						if v132 != 0 {
							v135 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							if int32(3) < v135 {
								F_sdsfree(m, v132)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return int32(0)
								} else {
									v145 = int32(0)
									v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
									v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
										return int32(0)
									} else {
										v151 = v145
										m.G0 = v6 + int32(160)
										return v151
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v132
								F__serverLog(m, int32(3), int32(_a1003), v6)
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return int32(0)
								} else {
									F_sdsfree(m, v132)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										v145 = int32(0)
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+84))
										v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l0, v145)
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int32(0)
										} else {
											v151 = v145
											m.G0 = v6 + int32(160)
											return v151
										}
									}
								}
							}
						} else {
							v151 = int32(1)
							m.G0 = v6 + int32(160)
							return v151
						}
					}
				}
			}
		}
	}
}
