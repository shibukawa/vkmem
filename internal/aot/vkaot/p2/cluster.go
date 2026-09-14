package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_clusterAllReplicasThinkPrimaryIsFail(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	v1 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_clusterAllReplicasThinkPrimaryIsFail[0]))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+88)))
	if v7&int32(2) == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_clusterAllReplicasThinkPrimaryIsFail_0), int32(_a_F_clusterAllReplicasThinkPrimaryIsFail_1), int32(5505))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L15
	}
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+2172))
	if v12 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	return v48
L4:
	;
	v29 = int32(0)
	goto L11
L5:
	;
	F__serverAssert(m, int32(_a_F_clusterAllReplicasThinkPrimaryIsFail_2), int32(_a_F_clusterAllReplicasThinkPrimaryIsFail_1), int32(5506))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v15 = int32(1)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2164))
	if v16 < v15 {
		v48 = v15
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2168))
	goto L4
L8:
	;
	return int32(0)
L9:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	v48 = int32(0)
	goto L3
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v19+v29<<(uint(int32(2))%32))))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+89)))
	if v37&int32(32) == int32(0) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v42 = int32(1)
	v44 = v29 + v42
	if v44 == v16 {
		v48 = v42
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v29 = v44
	goto L11
L15:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterAllowFailoverCmd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_clusterAllowFailoverCmd[0]))
	if v3 != 0 {
		F_addReplyError(m, l0, int32(_a_F_clusterAllowFailoverCmd_0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		return int32(1)
	}
}
func F_clusterAutoFailoverOnShutdown(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(176)
	m.G0 = v7
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_clusterAutoFailoverOnShutdown[0]))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+88)))
	if v12&int32(1) == v1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_clusterAutoFailoverOnShutdown_0), int32(_a_F_clusterAutoFailoverOnShutdown_1), int32(1616))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L33
	} else {
		goto L43
	}
L2:
	;
	m.G0 = v7 + int32(176)
	return
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_clusterAutoFailoverOnShutdown[1]))
	v20 = v7 + int32(168)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v21
	goto L4
L4:
	;
	v26 = v7 + int32(168)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v28 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v83)+104))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v122
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = int64(171798691880)
	v132 = F_snprintf(m, v7+int32(32), int32(128), int32(_a_F_clusterAutoFailoverOnShutdown_2), v7+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L33
	} else {
		goto L37
	}
L6:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_clusterAutoFailoverOnShutdown[2]))
	if int32(2) < v113 {
		goto L2
	} else {
		goto L35
	}
L7:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_clusterAutoFailoverOnShutdown[2]))
	if int32(2) < v104 {
		goto L2
	} else {
		goto L32
	}
L8:
	;
	if v28 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28+base.B2i32(v31 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v37
	goto L9
L11:
	;
	v42 = v1
	v43 = v28
	goto L12
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+104))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+156))
	if v47 < int32(589824) {
		goto L6
	} else {
		goto L14
	}
L13:
	;
	if v83 != 0 {
		goto L5
	} else {
		goto L31
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v50 != int32(9) {
		v83 = v42
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v86 = v7 + int32(168)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v88 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L16:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v46)+64))
	v55 = *(*int64)(unsafe.Add(mBase, _c_F_clusterAutoFailoverOnShutdown[3]))
	if v53 != v55 {
		v83 = v42
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v46)+192))
	if v57 == int32(0) {
		v83 = v42
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v60 = int32(-1)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+v60))))
	switch v62&int32(7) + v60 {
	case 0:
		goto L23
	case 1:
		goto L22
	case 2:
		goto L21
	case 3:
		goto L20
	default:
		v83 = v42
		goto L15
	}
L19:
	;
	if v79 == int32(40) {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v57+int32(-17))))
	v79 = v78
	goto L19
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v57+int32(-9))))
	v79 = v75
	goto L19
L22:
	;
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57+int32(-5)))))
	v79 = v72
	goto L19
L23:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+int32(-3)))))
	v79 = v69
	goto L19
L24:
	;
	v82 = v45
	goto L26
L25:
	;
	v82 = v42
	goto L26
L26:
	;
	v83 = v82
	goto L15
L27:
	;
	if v88 != 0 {
		v42 = v83
		v43 = v88
		goto L12
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v88+base.B2i32(v91 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v97
	goto L28
L30:
	;
	goto L13
L31:
	;
	goto L7
L32:
	;
	F__serverLog(m, int32(2), int32(_a_F_clusterAutoFailoverOnShutdown_3), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return
L34:
	;
	goto L2
L35:
	;
	F__serverLog(m, int32(2), int32(_a_F_clusterAutoFailoverOnShutdown_4), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L2
L37:
	;
	if base.Ui32(int32(129)) <= base.Ui32(v132) {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v136 = F_prepareReplicasToWrite(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	F_feedReplicationBuffer(m, v7+int32(32), v132)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L33
	} else {
		goto L40
	}
L40:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_clusterAutoFailoverOnShutdown[2]))
	if int32(2) < v143 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v83)+104))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v147
	F__serverLog(m, int32(2), int32(_a_F_clusterAutoFailoverOnShutdown_5), v7)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L33
	} else {
		goto L42
	}
L42:
	;
	goto L2
L43:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterBeforeSleep(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v110 int32
	_ = v110
	v1 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_clusterBeforeSleep[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+uint32(_c_F_clusterBeforeSleep[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+uint32(_c_F_clusterBeforeSleep[1]))) = v1
	if v5&int32(2) == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v5&int32(16) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	F_clusterUpdateState(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	goto L1
L5:
	;
	if v5&int32(64) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	F_clusterHandleReplicaFailover(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L15
	}
L7:
	;
	if v5&int32(1) == int32(0) {
		goto L5
	} else {
		goto L12
	}
L8:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_clusterBeforeSleep[6]))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+88)))
	if v20&int32(2) == v18 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_clusterHandleManualFailover(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_clusterBeforeSleep[7])))
	if v28&int32(2) == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L5
L12:
	;
	v37 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_clusterBeforeSleep[6]))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+88)))
	if v39&int32(2) == v37 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_clusterBeforeSleep[7])))
	if v45&int32(2) != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L6
L15:
	;
	goto L5
L16:
	;
	if v5&int32(4) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	F_clusterSlotMigrationCron(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	if v5&int32(32) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L20:
	;
	v61 = v5 & int32(8)
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_clusterBeforeSleep[2]))
	switch v63 {
	case 0:
		goto L22
	case 1:
		goto L21
	default:
		goto L19
	}
L21:
	;
	v79 = F_clusterSaveConfig(m, v61)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L28
	}
L22:
	;
	v64 = F_clusterSaveConfig(m, v61)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	if v64 != int32(-1) {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_clusterBeforeSleep[3]))
	if int32(3) < v69 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	F__serverLog(m, int32(3), int32(_a_F_clusterBeforeSleep_0), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	if v79 != int32(-1) {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v84 = *(*int64)(unsafe.Add(mBase, _c_F_clusterBeforeSleep[4]))
	v86 = *(*int64)(unsafe.Add(mBase, _c_F_clusterBeforeSleep[5]))
	if v84-v86 < int64(31) {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_clusterBeforeSleep[3]))
	if int32(3) < v91 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v101 = *(*int64)(unsafe.Add(mBase, _c_F_clusterBeforeSleep[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_clusterBeforeSleep[5])) = v101
	goto L19
L32:
	;
	F__serverLog(m, int32(3), int32(_a_F_clusterBeforeSleep_1), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	return
L35:
	;
	F_clusterBroadcastPong(m, int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	goto L34
}
func F_clusterBroadcastMessage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
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
	var v47 int64
	_ = v47
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
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
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int64
	_ = v160
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
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v176 int64
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_clusterBroadcastMessage[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+32))
	v7 = F_dictGetSafeIterator(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v7)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L2
	} else {
		goto L64
	}
L2:
	;
	return
L3:
	;
	v16 = v7 + int32(20)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v112 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v23 = v16
	v24 = v20
	goto L8
L6:
	;
	v20 = int32(1)
	goto L5
L7:
	;
	v20 = int32(0)
	goto L5
L8:
	;
	switch v24 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v24 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v104
	if v104 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v28 != int32(-1) {
		v67 = v28
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v68 = int32(1)
	v69 = v67 + v68
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v69
	v71 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v75+int32(26)))))
	if v79 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v32 != 0 {
		v67 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v34 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	if v61 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v41 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v33)+16)))
	v42 = int64(*(*int8)(unsafe.Add(mBase, uint32(v33)+27)))
	v43 = int64(*(*int32)(unsafe.Add(mBase, uint32(v33)+8)))
	v44 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v33)+12)))
	v45 = int64(*(*int8)(unsafe.Add(mBase, uint32(v33)+26)))
	v46 = int64(*(*int32)(unsafe.Add(mBase, uint32(v33)+4)))
	v47 = F_wangHash64(m, v46)
	mBase = m.M
	v49 = F_wangHash64(m, v45+v47)
	mBase = m.M
	v51 = F_wangHash64(m, v44+v49)
	mBase = m.M
	v53 = F_wangHash64(m, v43+v51)
	mBase = m.M
	v55 = F_wangHash64(m, v42+v53)
	mBase = m.M
	v57 = F_wangHash64(m, v41+v55)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v60 = v59
	goto L17
L19:
	;
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+24)))
	v39 = v37 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+24)) = uint16(v39)
	v60 = v33
	goto L17
L20:
	;
	v67 = v61 + int32(-1)
	goto L14
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v67 = v64
	goto L14
L22:
	;
	v94 = int32(2)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v74+v92<<(uint(v94)%32)+int32(4))))
	v23 = v99 + v93<<(uint(v94)%32)
	v24 = int32(1)
	goto L8
L23:
	;
	v83 = v71
	goto L25
L24:
	;
	v83 = v68 << (uint(v79) % 32)
	goto L25
L25:
	;
	if v69 < v83 {
		v92 = v75
		v93 = v69
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v75 != 0 {
		v112 = v71
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	if v85 == int32(-1) {
		v112 = v71
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967296)
	v92 = int32(1)
	v93 = int32(0)
	goto L22
L29:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v108
	v112 = v104
	goto L11
L30:
	;
	v120 = v112
	goto L31
L31:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	goto L34
L32:
	;
	goto L1
L33:
	;
	v135 = v7 + int32(20)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v136 != 0 {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+88)))
	if v122&int32(48) != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121)+2344))
	F_clusterSendMessage(m, v125, l0)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	if v231 != 0 {
		v120 = v231
		goto L31
	} else {
		goto L63
	}
L38:
	;
	v142 = v135
	v143 = v139
	goto L41
L39:
	;
	v139 = int32(1)
	goto L38
L40:
	;
	v139 = int32(0)
	goto L38
L41:
	;
	switch v143 {
	case 0:
		goto L46
	default:
		goto L45
	}
L43:
	;
	v143 = int32(0)
	goto L41
L44:
	;
	goto L37
L45:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v223
	if v223 == int32(0) {
		goto L43
	} else {
		goto L62
	}
L46:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v147 != int32(-1) {
		v186 = v147
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v187 = int32(1)
	v188 = v186 + v187
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v188
	v190 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v194+int32(26)))))
	if v198 == int32(255) {
		goto L56
	} else {
		goto L57
	}
L48:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v151 != 0 {
		v186 = int32(-1)
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v153 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	if v180 != int32(-1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v160 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v152)+16)))
	v161 = int64(*(*int8)(unsafe.Add(mBase, uint32(v152)+27)))
	v162 = int64(*(*int32)(unsafe.Add(mBase, uint32(v152)+8)))
	v163 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v152)+12)))
	v164 = int64(*(*int8)(unsafe.Add(mBase, uint32(v152)+26)))
	v165 = int64(*(*int32)(unsafe.Add(mBase, uint32(v152)+4)))
	v166 = F_wangHash64(m, v165)
	mBase = m.M
	v168 = F_wangHash64(m, v164+v166)
	mBase = m.M
	v170 = F_wangHash64(m, v163+v168)
	mBase = m.M
	v172 = F_wangHash64(m, v162+v170)
	mBase = m.M
	v174 = F_wangHash64(m, v161+v172)
	mBase = m.M
	v176 = F_wangHash64(m, v160+v174)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v179 = v178
	goto L50
L52:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+24)))
	v158 = v156 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v152)+24)) = uint16(v158)
	v179 = v152
	goto L50
L53:
	;
	v186 = v180 + int32(-1)
	goto L47
L54:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v186 = v183
	goto L47
L55:
	;
	v213 = int32(2)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v193+v211<<(uint(v213)%32)+int32(4))))
	v142 = v218 + v212<<(uint(v213)%32)
	v143 = int32(1)
	goto L41
L56:
	;
	v202 = v190
	goto L58
L57:
	;
	v202 = v187 << (uint(v198) % 32)
	goto L58
L58:
	;
	if v188 < v202 {
		v211 = v194
		v212 = v188
		goto L55
	} else {
		goto L59
	}
L59:
	;
	if v194 != 0 {
		v231 = v190
		goto L44
	} else {
		goto L60
	}
L60:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v193)+20))
	if v204 == int32(-1) {
		v231 = v190
		goto L44
	} else {
		goto L61
	}
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967296)
	v211 = int32(1)
	v212 = int32(0)
	goto L55
L62:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v227
	v231 = v223
	goto L44
L63:
	;
	goto L32
L64:
	;
	return
}
func F_clusterCommandExtendedHelp(m *base.Module) int32 {
	return int32(_a_F_clusterCommandExtendedHelp_0)
}
func F_clusterCommandFlushslot(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandFlushslot[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v8 = F_getSlotOrReply(m, l0, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandFlushslot[1]))
	F_addReplyErrorObject(m, l0, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L3
	} else {
		goto L37
	}
L2:
	;
	return
L3:
	;
	return
L4:
	;
	if v8 == int32(-1) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v12 != int32(4) {
		v101 = v5
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v104 = F_delKeysInSlot(m, v8, v101, int32(0), int32(1))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L35
	}
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v17 = F_objectGetVal(m, v16)
	mBase = m.M
	v18 = int32(_a_F_clusterCommandFlushslot_0)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v21 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v60 = F_objectGetVal(m, v59)
	mBase = m.M
	v61 = int32(_a_F_clusterCommandFlushslot_1)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v64 != 0 {
		goto L24
	} else {
		goto L25
	}
L9:
	;
	if v53-v55 != 0 {
		goto L8
	} else {
		goto L21
	}
L10:
	;
	v53 = F_tolower(m, v49)
	mBase = m.M
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v55 = F_tolower(m, v54)
	mBase = m.M
	goto L9
L11:
	;
	v23 = v17
	v24 = v18
	v25 = v21
	goto L14
L12:
	;
	v49 = int32(0)
	v50 = v18
	goto L10
L13:
	;
	v49 = v46 & int32(255)
	v50 = v45
	goto L10
L14:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v27 == int32(0) {
		v45 = v24
		v46 = v25
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v45 = v39
	v46 = int32(0)
	goto L13
L16:
	;
	v31 = v25 & int32(255)
	if v31 == v27 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v38 = int32(1)
	v39 = v24 + v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v40 != 0 {
		v23 = v23 + v38
		v24 = v39
		v25 = v40
		goto L14
	} else {
		goto L20
	}
L18:
	;
	v33 = F_tolower(m, v31)
	mBase = m.M
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v35 = F_tolower(m, v34)
	mBase = m.M
	if v33 == v35 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v45 = v24
	v46 = v37
	goto L13
L20:
	;
	goto L15
L21:
	;
	v101 = int32(1)
	goto L6
L22:
	;
	if v96-v98 != 0 {
		goto L1
	} else {
		goto L34
	}
L23:
	;
	v96 = F_tolower(m, v92)
	mBase = m.M
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	v98 = F_tolower(m, v97)
	mBase = m.M
	goto L22
L24:
	;
	v66 = v60
	v67 = v61
	v68 = v64
	goto L27
L25:
	;
	v92 = int32(0)
	v93 = v61
	goto L23
L26:
	;
	v92 = v89 & int32(255)
	v93 = v88
	goto L23
L27:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v70 == int32(0) {
		v88 = v67
		v89 = v68
		goto L26
	} else {
		goto L29
	}
L28:
	;
	v88 = v82
	v89 = int32(0)
	goto L26
L29:
	;
	v74 = v68 & int32(255)
	if v74 == v70 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v81 = int32(1)
	v82 = v67 + v81
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v83 != 0 {
		v66 = v66 + v81
		v67 = v82
		v68 = v83
		goto L27
	} else {
		goto L33
	}
L31:
	;
	v76 = F_tolower(m, v74)
	mBase = m.M
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v78 = F_tolower(m, v77)
	mBase = m.M
	if v76 == v78 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v88 = v67
	v89 = v80
	goto L26
L33:
	;
	goto L28
L34:
	;
	v101 = int32(0)
	goto L6
L35:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandFlushslot[2]))
	F_addReply(m, l0, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	goto L2
L37:
	;
	return
}
func F_clusterCommandMigrateSlots(m *base.Module, l0 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v627 int32
	_ = v627
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = F_moduleVerifyAllAllowAtomicSlotMigrationOrReply(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(80)
	return
L2:
	;
	return
L3:
	;
	if v13 == int32(-1) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandMigrateSlots[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+88)))
	if v20&int32(1) != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandMigrateSlots[0]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+48))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	goto L9
L6:
	;
	F_addReplyError(m, l0, int32(_a_F_clusterCommandMigrateSlots_0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandMigrateSlots[0]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+44))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	goto L13
L9:
	;
	if base.B2i32(v30 != int32(0)-v32) == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_addReplyError(m, l0, int32(_a_F_clusterCommandMigrateSlots_1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	goto L1
L12:
	;
	v54 = F_listCreate(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L16
	}
L13:
	;
	if base.B2i32(v44 != int32(0)-v46) == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_addReplyError(m, l0, int32(_a_F_clusterCommandMigrateSlots_2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	goto L1
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = int32(103)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v58 < int32(3) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	F_listRelease(m, v54)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L2
	} else {
		goto L164
	}
L18:
	;
	if v130 == int32(0) {
		goto L17
	} else {
		goto L161
	}
L19:
	;
	v512 = v11 + int32(56)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v512)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v512))) = v513
	goto L138
L20:
	;
	v65 = int32(2)
	goto L21
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v65<<(uint(int32(2))%32))))
	v75 = F_objectGetVal(m, v74)
	mBase = m.M
	v76 = int32(_a_F_clusterCommandMigrateSlots_3)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v79 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	goto L19
L23:
	;
	v122 = v65 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = int32(0)
	v130 = F_parseSlotRangesOrReply(m, l0, v122, v11+int32(76), v11+int32(72))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L38
	}
L24:
	;
	if v111-v113 == int32(0) {
		goto L23
	} else {
		goto L36
	}
L25:
	;
	v111 = F_tolower(m, v107)
	mBase = m.M
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v113 = F_tolower(m, v112)
	mBase = m.M
	goto L24
L26:
	;
	v81 = v75
	v82 = v76
	v83 = v79
	goto L29
L27:
	;
	v107 = int32(0)
	v108 = v76
	goto L25
L28:
	;
	v107 = v104 & int32(255)
	v108 = v103
	goto L25
L29:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v85 == int32(0) {
		v103 = v82
		v104 = v83
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v103 = v97
	v104 = int32(0)
	goto L28
L31:
	;
	v89 = v83 & int32(255)
	if v89 == v85 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v96 = int32(1)
	v97 = v82 + v96
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	if v98 != 0 {
		v81 = v81 + v96
		v82 = v97
		v83 = v98
		goto L29
	} else {
		goto L35
	}
L33:
	;
	v91 = F_tolower(m, v89)
	mBase = m.M
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v93 = F_tolower(m, v92)
	mBase = m.M
	if v91 == v93 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v103 = v82
	v104 = v95
	goto L28
L35:
	;
	goto L30
L36:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandMigrateSlots[1]))
	F_addReplyErrorObject(m, l0, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	goto L17
L38:
	;
	if v130 == int32(0) {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandMigrateSlots[0]))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	if v134 == v137 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v143 = v11 + int32(64)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v144
	goto L43
L41:
	;
	F_addReplyError(m, l0, int32(_a_F_clusterCommandMigrateSlots_4))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	goto L18
L43:
	;
	v149 = v11 + int32(64)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	if v151 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v287 = v11 + int32(64)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v287)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = v288
	goto L80
L45:
	;
	if v151 == int32(0) {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v151+base.B2i32(v154 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v160
	goto L46
L48:
	;
	v167 = v151
	goto L50
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v179
	F_addReplyErrorFormat(m, l0, int32(_a_F_clusterCommandMigrateSlots_5), v11+int32(48))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L2
	} else {
		goto L79
	}
L50:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v174 < v173 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v257 = v11 + int32(64)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	if v259 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L53:
	;
	v179 = v173
	goto L54
L54:
	;
	v184 = int32(0)
	v187 = m.G0
	v189 = v187 - int32(16)
	m.G0 = v189
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandMigrateSlots[0]))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+uint32(_c_F_clusterCommandMigrateSlots[2])))
	F_listRewind(m, v193, v189)
	mBase = m.M
	v196 = F_listNext(m, v189)
	mBase = m.M
	if v196 == v184 {
		v239 = v184
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L52
L56:
	;
	if v239 != 0 {
		goto L49
	} else {
		goto L73
	}
L57:
	;
	m.G0 = v189 + int32(16)
	goto L56
L58:
	;
	v202 = v196
	goto L60
L59:
	;
	v239 = int32(1)
	goto L57
L60:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	if v204 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v235 = F_listNext(m, v189)
	mBase = m.M
	if v235 != 0 {
		v202 = v235
		goto L60
	} else {
		goto L72
	}
L63:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v203)+156))
	if base.Ui32(v205+int32(-18)) < base.Ui32(int32(3)) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v203)+164))
	v212 = v189 + int32(8)
	F_listRewind(m, v210, v212)
	mBase = m.M
	v216 = F_listNext(m, v212)
	mBase = m.M
	if v216 == int32(0) {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v222 = v216
	goto L66
L66:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	if v179 < v224 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L62
L68:
	;
	v230 = F_listNext(m, v189+int32(8))
	mBase = m.M
	if v230 != 0 {
		v222 = v230
		goto L66
	} else {
		goto L71
	}
L69:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v179 <= v226 {
		goto L59
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	goto L67
L72:
	;
	v239 = v184
	goto L57
L73:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v179 < v244 {
		v179 = v179 + int32(1)
		goto L54
	} else {
		goto L74
	}
L74:
	;
	goto L55
L75:
	;
	if v259 == int32(0) {
		goto L44
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v259+base.B2i32(v262 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = v268
	goto L76
L78:
	;
	v167 = v259
	goto L50
L79:
	;
	goto L18
L80:
	;
	goto L82
L81:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	v375 = v373 + int32(1)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v376 <= v375 {
		goto L103
	} else {
		goto L104
	}
L82:
	;
	v301 = v11 + int32(64)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	if v303 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	F_addReplyError(m, l0, int32(_a_F_clusterCommandMigrateSlots_6))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L2
	} else {
		goto L101
	}
L84:
	;
	if v303 == int32(0) {
		goto L81
	} else {
		goto L87
	}
L85:
	;
	goto L84
L86:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v303+base.B2i32(v306 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v301))) = v312
	goto L85
L87:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v303)+8))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)+164))
	v318 = int32(0)
	v321 = m.G0
	v323 = v321 - int32(16)
	m.G0 = v323
	v326 = v323 + int32(8)
	F_listRewind(m, v317, v326)
	mBase = m.M
	v331 = F_listNext(m, v326)
	mBase = m.M
	if v331 == v318 {
		v363 = v318
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if v363 == int32(0) {
		goto L82
	} else {
		goto L100
	}
L89:
	;
	m.G0 = v323 + int32(16)
	goto L88
L90:
	;
	v334 = v331
	goto L91
L91:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v334)+8))
	F_listRewind(m, v130, v323)
	mBase = m.M
	goto L93
L92:
	;
	v363 = int32(1)
	goto L89
L93:
	;
	v346 = F_listNext(m, v323)
	mBase = m.M
	if v346 != 0 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L92
L95:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v346)+8))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	if v352 < v354 {
		goto L93
	} else {
		goto L98
	}
L96:
	;
	v349 = F_listNext(m, v323+int32(8))
	mBase = m.M
	if v349 == int32(0) {
		v363 = v318
		goto L89
	} else {
		goto L97
	}
L97:
	;
	v334 = v349
	goto L91
L98:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	if v357 < v356 {
		goto L93
	} else {
		goto L99
	}
L99:
	;
	goto L94
L100:
	;
	goto L83
L101:
	;
	goto L18
L102:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v431 = v375 << (uint(int32(2)) % 32)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v429+v431)))
	v434 = F_objectGetVal(m, v433)
	mBase = m.M
	v435 = int32(-1)
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434+v435))))
	switch v437&int32(7) + v435 {
	case 0:
		goto L125
	case 1:
		goto L124
	case 2:
		goto L123
	case 3:
		goto L122
	default:
		goto L120
	}
L103:
	;
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandMigrateSlots[1]))
	F_addReplyErrorObject(m, l0, v426)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L2
	} else {
		goto L118
	}
L104:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v378+v373<<(uint(int32(2))%32))))
	v383 = F_objectGetVal(m, v382)
	mBase = m.M
	v384 = int32(_a_F_clusterCommandMigrateSlots_7)
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383))))
	if v387 != 0 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	if v419-v421 == int32(0) {
		goto L102
	} else {
		goto L117
	}
L106:
	;
	v419 = F_tolower(m, v415)
	mBase = m.M
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416))))
	v421 = F_tolower(m, v420)
	mBase = m.M
	goto L105
L107:
	;
	v389 = v383
	v390 = v384
	v391 = v387
	goto L110
L108:
	;
	v415 = int32(0)
	v416 = v384
	goto L106
L109:
	;
	v415 = v412 & int32(255)
	v416 = v411
	goto L106
L110:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390))))
	if v393 == int32(0) {
		v411 = v390
		v412 = v391
		goto L109
	} else {
		goto L112
	}
L111:
	;
	v411 = v405
	v412 = int32(0)
	goto L109
L112:
	;
	v397 = v391 & int32(255)
	if v397 == v393 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v404 = int32(1)
	v405 = v390 + v404
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+1)))
	if v406 != 0 {
		v389 = v389 + v404
		v390 = v405
		v391 = v406
		goto L110
	} else {
		goto L116
	}
L114:
	;
	v399 = F_tolower(m, v397)
	mBase = m.M
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390))))
	v401 = F_tolower(m, v400)
	mBase = m.M
	if v399 == v401 {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	v411 = v390
	v412 = v403
	goto L109
L116:
	;
	goto L111
L117:
	;
	goto L103
L118:
	;
	goto L18
L119:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v468+v431)))
	v471 = F_objectGetVal(m, v470)
	mBase = m.M
	v473 = F_clusterLookupNode(m, v471, int32(40))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L2
	} else {
		goto L129
	}
L120:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v458+v375<<(uint(int32(2))%32))))
	v463 = F_objectGetVal(m, v462)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v463
	F_addReplyErrorFormat(m, l0, int32(_a_F_clusterCommandMigrateSlots_8), v11)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L2
	} else {
		goto L127
	}
L121:
	;
	if v454 == int32(40) {
		goto L119
	} else {
		goto L126
	}
L122:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v434+int32(-17))))
	v454 = v453
	goto L121
L123:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v434+int32(-9))))
	v454 = v450
	goto L121
L124:
	;
	v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v434+int32(-5)))))
	v454 = v447
	goto L121
L125:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434+int32(-3)))))
	v454 = v444
	goto L121
L126:
	;
	goto L120
L127:
	;
	goto L18
L128:
	;
	v488 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandMigrateSlots[0]))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)))
	if v473 != v489 {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	if v473 != 0 {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v475+v375<<(uint(int32(2))%32))))
	v480 = F_objectGetVal(m, v479)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v480
	F_addReplyErrorFormat(m, l0, int32(_a_F_clusterCommandMigrateSlots_9), v11+int32(16))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	goto L18
L132:
	;
	v495 = v373 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v495
	v497 = F_createSlotExportJob(m, v473, v130)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L2
	} else {
		goto L135
	}
L133:
	;
	F_addReplyError(m, l0, int32(_a_F_clusterCommandMigrateSlots_10))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L2
	} else {
		goto L134
	}
L134:
	;
	goto L18
L135:
	;
	v499 = F_listAddNodeHead(m, v54, v497)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L2
	} else {
		goto L136
	}
L136:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v495 < v501 {
		v65 = v495
		goto L21
	} else {
		goto L137
	}
L137:
	;
	goto L22
L138:
	;
	v517 = F_sdsempty(m)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L2
	} else {
		goto L139
	}
L139:
	;
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandMigrateSlots[3]))
	v521 = F_catClientInfoShortString(m, v517, l0, v520)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L2
	} else {
		goto L140
	}
L140:
	;
	v524 = v11 + int32(56)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	if v526 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = int32(0)
	F_listRelease(m, v54)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L2
	} else {
		goto L158
	}
L142:
	;
	if v526 == int32(0) {
		goto L141
	} else {
		goto L145
	}
L143:
	;
	goto L142
L144:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v526+base.B2i32(v529 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v524))) = v535
	goto L143
L145:
	;
	v542 = v526
	goto L146
L146:
	;
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandMigrateSlots[0]))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+uint32(_c_F_clusterCommandMigrateSlots[2])))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v542)+8))
	v551 = F_listAddNodeHead(m, v549, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L2
	} else {
		goto L148
	}
L147:
	;
	goto L141
L148:
	;
	v554 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandMigrateSlots[4]))
	if int32(2) < v554 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	F_fireModuleSlotMigrationEvent(m, v550, int32(1))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L2
	} else {
		goto L152
	}
L150:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v550)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v557
	F__serverLog(m, int32(2), int32(_a_F_clusterCommandMigrateSlots_11), v11+int32(32))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L2
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	F_proceedWithSlotMigration(m, v550)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L2
	} else {
		goto L153
	}
L153:
	;
	v573 = v11 + int32(56)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	if v575 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	if v575 != 0 {
		v542 = v575
		goto L146
	} else {
		goto L157
	}
L155:
	;
	goto L154
L156:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v575+base.B2i32(v578 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v573))) = v584
	goto L155
L157:
	;
	goto L147
L158:
	;
	F_sdsfree(m, v521)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L2
	} else {
		goto L159
	}
L159:
	;
	v601 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandMigrateSlots[5]))
	F_addReply(m, l0, v601)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L2
	} else {
		goto L160
	}
L160:
	;
	goto L1
L161:
	;
	F_listRelease(m, v130)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L2
	} else {
		goto L162
	}
L162:
	;
	F_listRelease(m, v54)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L2
	} else {
		goto L163
	}
L163:
	;
	goto L1
L164:
	;
	goto L1
}
func F_clusterCommandSetSlot(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
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
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
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
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v609 int64
	_ = v609
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	v10 = m.G0
	v12 = v10 - int32(144)
	m.G0 = v12
	v20 = F_clusterParseSetSlotCommand(m, l0, v12+int32(140), v12+int32(124), v12+int32(128))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_clusterCommandSetSlot_0), int32(_a_F_clusterCommandSetSlot_1), int32(7912))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L5
	} else {
		goto L209
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_clusterCommandSetSlot_0), int32(_a_F_clusterCommandSetSlot_1), int32(7892))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L5
	} else {
		goto L208
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_clusterCommandSetSlot_2), int32(_a_F_clusterCommandSetSlot_1), int32(7623))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L5
	} else {
		goto L207
	}
L4:
	;
	m.G0 = v12 + int32(144)
	return
L5:
	;
	return
L6:
	;
	if v20 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v24 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[0]))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+88)))
	if v26&int32(1) == v24 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v126 = F_objectGetVal(m, v125)
	mBase = m.M
	v127 = int32(_a_F_clusterCommandSetSlot_3)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 != 0 {
		goto L33
	} else {
		goto L34
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+2164))
	if v31 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+206)))
	if v34&int32(64) != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[1]))
	v40 = v12 + int32(116)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v41
	goto L12
L12:
	;
	v46 = v12 + int32(116)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v48 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v48 == int32(0) {
		goto L8
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v48+base.B2i32(v51 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v57
	goto L14
L16:
	;
	v64 = v48
	v65 = int32(0)
	goto L17
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+104))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v73 != int32(9) {
		v80 = v65
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v80 == int32(0) {
		goto L8
	} else {
		goto L25
	}
L19:
	;
	v82 = v12 + int32(116)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v84 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+156))
	v80 = v65 + base.B2i32(int32(459519) < v76)
	goto L19
L21:
	;
	if v84 != 0 {
		v64 = v84
		v65 = v80
		goto L17
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84+base.B2i32(v87 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v93
	goto L22
L24:
	;
	goto L18
L25:
	;
	F_forceCommandPropagation(m, l0, int32(2))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v12)+128))
	v102 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[2]))
	F_blockClientForReplicaAck(m, l0, v100, v102+int64(1), v80, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v108 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[3])) = int32(1)
	goto L28
L28:
	;
	goto L4
L29:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L5
	} else {
		goto L205
	}
L30:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
	v209 = F_objectGetVal(m, v208)
	mBase = m.M
	v210 = int32(_a_F_clusterCommandSetSlot_4)
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v213 != 0 {
		goto L61
	} else {
		goto L62
	}
L31:
	;
	if v162-v164 != 0 {
		goto L30
	} else {
		goto L43
	}
L32:
	;
	v162 = F_tolower(m, v158)
	mBase = m.M
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v164 = F_tolower(m, v163)
	mBase = m.M
	goto L31
L33:
	;
	v132 = v126
	v133 = v127
	v134 = v130
	goto L36
L34:
	;
	v158 = int32(0)
	v159 = v127
	goto L32
L35:
	;
	v158 = v155 & int32(255)
	v159 = v154
	goto L32
L36:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v136 == int32(0) {
		v154 = v133
		v155 = v134
		goto L35
	} else {
		goto L38
	}
L37:
	;
	v154 = v148
	v155 = int32(0)
	goto L35
L38:
	;
	v140 = v134 & int32(255)
	if v140 == v136 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v147 = int32(1)
	v148 = v133 + v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v149 != 0 {
		v132 = v132 + v147
		v133 = v148
		v134 = v149
		goto L36
	} else {
		goto L42
	}
L40:
	;
	v142 = F_tolower(m, v140)
	mBase = m.M
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v144 = F_tolower(m, v143)
	mBase = m.M
	if v142 == v144 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v154 = v133
	v155 = v146
	goto L35
L42:
	;
	goto L37
L43:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[4]))
	if int32(2) < v167 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+44))
	v190 = F_dictFind(m, v189, v186)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L49
	}
L45:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v12)+140))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
	v185 = v184
	v186 = v183
	goto L44
L46:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v12)+140))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
	v172 = F_humanNodename(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v171 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v170
	F__serverLog(m, int32(2), int32(_a_F_clusterCommandSetSlot_5), v12)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v185 = v171
	v186 = v170
	goto L44
L49:
	;
	if v185 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+44))
	if v190 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	if v190 == int32(0) {
		goto L29
	} else {
		goto L52
	}
L52:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+44))
	v197 = F_dictDelete(m, v196, v186)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	goto L29
L54:
	;
	v205 = F_dictAdd(m, v201, v186, v185)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L5
	} else {
		goto L57
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = v185
	goto L56
L56:
	;
	goto L29
L57:
	;
	goto L29
L58:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
	v294 = F_objectGetVal(m, v293)
	mBase = m.M
	v295 = int32(_a_F_clusterCommandSetSlot_6)
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v298 != 0 {
		goto L89
	} else {
		goto L90
	}
L59:
	;
	if v245-v247 != 0 {
		goto L58
	} else {
		goto L71
	}
L60:
	;
	v245 = F_tolower(m, v241)
	mBase = m.M
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	v247 = F_tolower(m, v246)
	mBase = m.M
	goto L59
L61:
	;
	v215 = v209
	v216 = v210
	v217 = v213
	goto L64
L62:
	;
	v241 = int32(0)
	v242 = v210
	goto L60
L63:
	;
	v241 = v238 & int32(255)
	v242 = v237
	goto L60
L64:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	if v219 == int32(0) {
		v237 = v216
		v238 = v217
		goto L63
	} else {
		goto L66
	}
L65:
	;
	v237 = v231
	v238 = int32(0)
	goto L63
L66:
	;
	v223 = v217 & int32(255)
	if v223 == v219 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v230 = int32(1)
	v231 = v216 + v230
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	if v232 != 0 {
		v215 = v215 + v230
		v216 = v231
		v217 = v232
		goto L64
	} else {
		goto L70
	}
L68:
	;
	v225 = F_tolower(m, v223)
	mBase = m.M
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v227 = F_tolower(m, v226)
	mBase = m.M
	if v225 == v227 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	v237 = v216
	v238 = v229
	goto L63
L70:
	;
	goto L65
L71:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[4]))
	if int32(2) < v250 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+48))
	v275 = F_dictFind(m, v274, v271)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L5
	} else {
		goto L77
	}
L73:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v12)+140))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
	v270 = v269
	v271 = v268
	goto L72
L74:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v12)+140))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
	v255 = F_humanNodename(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v254 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v253
	F__serverLog(m, int32(2), int32(_a_F_clusterCommandSetSlot_7), v12+int32(16))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	v270 = v254
	v271 = v253
	goto L72
L77:
	;
	if v270 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+48))
	if v275 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	if v275 == int32(0) {
		goto L29
	} else {
		goto L80
	}
L80:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)+48))
	v282 = F_dictDelete(m, v281, v271)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	goto L29
L82:
	;
	v290 = F_dictAdd(m, v286, v271, v270)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L85
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+8)) = v270
	goto L84
L84:
	;
	goto L29
L85:
	;
	goto L29
L86:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)+12))
	v372 = F_objectGetVal(m, v371)
	mBase = m.M
	v373 = int32(_a_F_clusterCommandSetSlot_8)
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	if v376 != 0 {
		goto L112
	} else {
		goto L113
	}
L87:
	;
	if v330-v332 != 0 {
		goto L86
	} else {
		goto L99
	}
L88:
	;
	v330 = F_tolower(m, v326)
	mBase = m.M
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	v332 = F_tolower(m, v331)
	mBase = m.M
	goto L87
L89:
	;
	v300 = v294
	v301 = v295
	v302 = v298
	goto L92
L90:
	;
	v326 = int32(0)
	v327 = v295
	goto L88
L91:
	;
	v326 = v323 & int32(255)
	v327 = v322
	goto L88
L92:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v304 == int32(0) {
		v322 = v301
		v323 = v302
		goto L91
	} else {
		goto L94
	}
L93:
	;
	v322 = v316
	v323 = int32(0)
	goto L91
L94:
	;
	v308 = v302 & int32(255)
	if v308 == v304 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v315 = int32(1)
	v316 = v301 + v315
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+1)))
	if v317 != 0 {
		v300 = v300 + v315
		v301 = v316
		v302 = v317
		goto L92
	} else {
		goto L98
	}
L96:
	;
	v310 = F_tolower(m, v308)
	mBase = m.M
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	v312 = F_tolower(m, v311)
	mBase = m.M
	if v310 == v312 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	v322 = v301
	v323 = v314
	goto L91
L98:
	;
	goto L93
L99:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v12)+140))
	v336 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[4]))
	if int32(2) < v336 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+48))
	v349 = F_dictFind(m, v348, v334)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L5
	} else {
		goto L104
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v334
	F__serverLog(m, int32(2), int32(_a_F_clusterCommandSetSlot_9), v12+int32(32))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+44))
	v361 = F_dictFind(m, v360, v334)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L5
	} else {
		goto L107
	}
L104:
	;
	if v349 == int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+48))
	v356 = F_dictDelete(m, v355, v334)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	goto L103
L107:
	;
	if v361 == int32(0) {
		goto L29
	} else {
		goto L108
	}
L108:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v366)+44))
	v368 = F_dictDelete(m, v367, v334)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	goto L29
L110:
	;
	if v408-v410 != 0 {
		goto L29
	} else {
		goto L122
	}
L111:
	;
	v408 = F_tolower(m, v404)
	mBase = m.M
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405))))
	v410 = F_tolower(m, v409)
	mBase = m.M
	goto L110
L112:
	;
	v378 = v372
	v379 = v373
	v380 = v376
	goto L115
L113:
	;
	v404 = int32(0)
	v405 = v373
	goto L111
L114:
	;
	v404 = v401 & int32(255)
	v405 = v400
	goto L111
L115:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	if v382 == int32(0) {
		v400 = v379
		v401 = v380
		goto L114
	} else {
		goto L117
	}
L116:
	;
	v400 = v394
	v401 = int32(0)
	goto L114
L117:
	;
	v386 = v380 & int32(255)
	if v386 == v382 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v393 = int32(1)
	v394 = v379 + v393
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+1)))
	if v395 != 0 {
		v378 = v378 + v393
		v379 = v394
		v380 = v395
		goto L115
	} else {
		goto L121
	}
L119:
	;
	v388 = F_tolower(m, v386)
	mBase = m.M
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	v390 = F_tolower(m, v389)
	mBase = m.M
	if v388 == v390 {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	v400 = v379
	v401 = v392
	goto L114
L121:
	;
	goto L116
L122:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v12)+140))
	v414 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[4]))
	if int32(2) < v414 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v443 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[6]))
	if int32(1) <= v443 {
		goto L130
	} else {
		goto L131
	}
L124:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
	v418 = F_humanNodename(m, v417)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L5
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v417 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = v418
	*(*int32)(unsafe.Add(mBase, uint32(v12)+100)) = v417 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v412
	F__serverLog(m, int32(2), int32(_a_F_clusterCommandSetSlot_10), v12+int32(96))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L5
	} else {
		goto L126
	}
L126:
	;
	goto L123
L127:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[0]))
	v497 = v493
	goto L142
L128:
	;
	if v480 != 0 {
		goto L127
	} else {
		goto L137
	}
L129:
	;
	goto L128
L130:
	;
	v448 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[7]))
	v449 = int32(0)
	v452 = v443
	v453 = v449
	v454 = v448
	v455 = v449
	goto L132
L131:
	;
	v480 = int32(0)
	goto L129
L132:
	;
	v458 = int32(0)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v454+v455<<(uint(int32(2))%32))))
	if v462 == v458 {
		v471 = v452
		v472 = v454
		v473 = v458
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v480 = v474
	goto L129
L134:
	;
	v474 = v473 + v453
	v476 = v455 + int32(1)
	if v476 < v471 {
		v452 = v471
		v453 = v474
		v454 = v472
		v455 = v476
		goto L132
	} else {
		goto L136
	}
L135:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	v466 = F_kvstoreHashtableSize(m, v465, v412)
	mBase = m.M
	v467 = int32(_a_F_clusterCommandSetSlot_11)
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[6]))
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[7]))
	v471 = v468
	v472 = v470
	v473 = v466
	goto L134
L136:
	;
	goto L133
L137:
	;
	v485 = F_getMigratingSlotDest(m, v412)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L5
	} else {
		goto L138
	}
L138:
	;
	if v485 == int32(0) {
		goto L127
	} else {
		goto L139
	}
L139:
	;
	F_setMigratingSlotDest(m, v412, int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L5
	} else {
		goto L140
	}
L140:
	;
	goto L127
L141:
	;
	v509 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[8]))
	if v509 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L142:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v497)+2172))
	if v503 == int32(0) {
		v507 = v497
		goto L141
	} else {
		goto L144
	}
L143:
	;
	v507 = v503
	goto L141
L144:
	;
	if v503 != v493 {
		v497 = v503
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v516 = v412 << (uint(int32(2)) % 32)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v514+v516+int32(52))))
	v521 = F_clusterDelSlot(m, v412)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L5
	} else {
		goto L149
	}
L147:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v507)+2172))
	if v512 != 0 {
		goto L3
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v525+v516+int32(52))))
	if v529 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v507)+2160))
	v623 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[9]))
	if v623 == int32(0) {
		v681 = v623
		goto L165
	} else {
		goto L166
	}
L151:
	;
	v533 = m.G0
	v535 = v533 - int32(32)
	m.G0 = v535
	v540 = int32(1) << (uint(v412&int32(7)) % 32)
	v542 = base.I32_div_s(v412, int32(8))
	v545 = v523 + v542 + int32(104)
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545))))
	if v540&v546 != 0 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v584 = int32(_a_F_clusterCommandSetSlot_11)
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v585+v516+int32(52)))) = v523
	v591 = base.I32_div_s(v412, int32(8))
	v592 = v585 + v591
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592)+uint32(_c_F_clusterCommandSetSlot[10]))))
	v600 = v595 & base.I32_rotl(int32(-2), v412&int32(7))
	*(*uint8)(unsafe.Add(mBase, uint32(v592)+uint32(_c_F_clusterCommandSetSlot[10]))) = uint8(v600)
	v603 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v606 = v603 + v412*int32(24)
	v609 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v606)+uint32(_c_F_clusterCommandSetSlot[11]))) = v609
	*(*int64)(unsafe.Add(mBase, uint32(v606)+uint32(_c_F_clusterCommandSetSlot[12]))) = v609
	*(*int64)(unsafe.Add(mBase, uint32(v606)+uint32(_c_F_clusterCommandSetSlot[13]))) = v609
	goto L164
L153:
	;
	m.G0 = v535 + int32(32)
	goto L152
L154:
	;
	v548 = v546 | v540
	*(*uint8)(unsafe.Add(mBase, uint32(v545))) = uint8(v548)
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v523)+2160))
	*(*int32)(unsafe.Add(mBase, uint32(v523)+2160)) = v550 + int32(1)
	if v550 != 0 {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)+32))
	F_dictInitIterator(m, v535, v556)
	mBase = m.M
	v558 = F_dictNext(m, v535)
	mBase = m.M
	if v558 == int32(0) {
		goto L153
	} else {
		goto L156
	}
L156:
	;
	v562 = v558
	goto L158
L157:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v523)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v523)+88)) = v572 | int32(256)
	goto L153
L158:
	;
	v566 = F_dictGetVal(m, v562)
	mBase = m.M
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566)+88)))
	if v567&int32(2) != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v571 = F_dictNext(m, v535)
	mBase = m.M
	if v571 != 0 {
		v562 = v571
		goto L158
	} else {
		goto L163
	}
L161:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v566)+2164))
	if v570 != 0 {
		goto L157
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	goto L153
L164:
	;
	goto L150
L165:
	;
	v683 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[0]))
	if v681 != 0 {
		v713 = v683
		goto L182
	} else {
		goto L183
	}
L166:
	;
	if v520 != v507 {
		v681 = v623
		goto L165
	} else {
		goto L167
	}
L167:
	;
	if v621 != 0 {
		v681 = v623
		goto L165
	} else {
		goto L168
	}
L168:
	;
	if v523 == v507 {
		goto L2
	} else {
		goto L169
	}
L169:
	;
	v629 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[4]))
	if int32(2) < v629 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v648 = int32(0)
	v649 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[0]))
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649)+88)))
	if v650&int32(2) == v648 {
		goto L174
	} else {
		goto L175
	}
L171:
	;
	v632 = F_humanNodename(m, v523)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L5
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = v523 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v632
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v523 + int32(8)
	F__serverLog(m, int32(2), int32(_a_F_clusterCommandSetSlot_12), v12+int32(80))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L5
	} else {
		goto L173
	}
L173:
	;
	goto L170
L174:
	;
	v657 = int32(1)
	F_clusterSetPrimary(m, v523, v657, v657)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L5
	} else {
		goto L177
	}
L175:
	;
	F_protectClient(m, l0)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L5
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	v661 = int32(0)
	v662 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[0]))
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+88)))
	if v663&int32(2) == v661 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L5
	} else {
		goto L181
	}
L179:
	;
	F_unprotectClient(m, l0)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L5
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	v672 = int32(_a_F_clusterCommandSetSlot_11)
	v673 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v673)+uint32(_c_F_clusterCommandSetSlot[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v673)+uint32(_c_F_clusterCommandSetSlot[14]))) = v674 | int32(46)
	v679 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[9]))
	v681 = v679
	goto L165
L182:
	;
	if v523 == v713 {
		goto L191
	} else {
		goto L192
	}
L183:
	;
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683)+88)))
	if v684&int32(1) == int32(0) {
		v713 = v683
		goto L182
	} else {
		goto L184
	}
L184:
	;
	if v520 != v507 {
		v713 = v683
		goto L182
	} else {
		goto L185
	}
L185:
	;
	if v621 != 0 {
		v713 = v683
		goto L182
	} else {
		goto L186
	}
L186:
	;
	if v523 == v507 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	v692 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[4]))
	if int32(2) < v692 {
		v713 = v683
		goto L182
	} else {
		goto L188
	}
L188:
	;
	v695 = F_humanNodename(m, v523)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L5
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v523 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v695
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v523 + int32(8)
	F__serverLog(m, int32(2), int32(_a_F_clusterCommandSetSlot_13), v12+int32(64))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L5
	} else {
		goto L190
	}
L190:
	;
	v711 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[0]))
	v713 = v711
	goto L182
L191:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v718)+48))
	v720 = F_dictFind(m, v719, v412)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L5
	} else {
		goto L194
	}
L192:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v713)+2172))
	if v523 != v715 {
		goto L29
	} else {
		goto L193
	}
L193:
	;
	goto L191
L194:
	;
	if v720 == int32(0) {
		goto L29
	} else {
		goto L195
	}
L195:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v720)+8))
	goto L196
L196:
	;
	if v724 == int32(0) {
		goto L29
	} else {
		goto L197
	}
L197:
	;
	F_setImportingSlotSource(m, v412, int32(0))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L5
	} else {
		goto L198
	}
L198:
	;
	v731 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[0]))
	if v523 != v731 {
		goto L29
	} else {
		goto L199
	}
L199:
	;
	v733 = F_clusterBumpConfigEpochWithoutConsensus(m)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L5
	} else {
		goto L201
	}
L200:
	;
	v747 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)+uint32(_c_F_clusterCommandSetSlot[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v747)+uint32(_c_F_clusterCommandSetSlot[14]))) = v748 | int32(32)
	goto L29
L201:
	;
	if v733 != 0 {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v736 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[4]))
	if int32(2) < v736 {
		goto L200
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v412
	F__serverLog(m, int32(2), int32(_a_F_clusterCommandSetSlot_14), v12+int32(48))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L5
	} else {
		goto L204
	}
L204:
	;
	goto L200
L205:
	;
	v764 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[5]))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v764)+uint32(_c_F_clusterCommandSetSlot[14])))
	*(*int32)(unsafe.Add(mBase, uint32(v764)+uint32(_c_F_clusterCommandSetSlot[14]))) = v765 | int32(6)
	v770 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSetSlot[15]))
	F_addReply(m, l0, v770)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L5
	} else {
		goto L206
	}
L206:
	;
	goto L4
L207:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L208:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L209:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterCommandSyncSlotsPaused(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v73 int64
	_ = v73
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v11 != 0 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+156))
		switch v15 + int32(-2) {
		case 0:
			F_sendSyncSlotsMessage(m, v11, int32(_a_F_clusterCommandSyncSlotsPaused_0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				v45 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSyncSlotsPaused[0]))
				if int32(2) < v45 {
					v73 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCommandSyncSlotsPaused[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v43)+156)) = int32(3)
					*(*int64)(unsafe.Add(mBase, uint32(v43)+16)) = v73
					m.G0 = v9 + int32(32)
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+188))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+156))
					if base.Ui32(int32(20)) < base.Ui32(v50) {
						v58 = int32(_a_F_clusterCommandSyncSlotsPaused_1)
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v50<<(uint(int32(2))%32))+uint32(_c_F_clusterCommandSyncSlotsPaused[2])))
						v58 = v57
					}
					*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a_F_clusterCommandSyncSlotsPaused_2)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v58
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v48
					F__serverLog(m, int32(2), int32(_a_F_clusterCommandSyncSlotsPaused_3), v9+int32(16))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						v73 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCommandSyncSlotsPaused[1]))
						*(*int32)(unsafe.Add(mBase, uint32(v43)+156)) = int32(3)
						*(*int64)(unsafe.Add(mBase, uint32(v43)+16)) = v73
						m.G0 = v9 + int32(32)
						return
					}
				}
			}
		default:
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCommandSyncSlotsPaused[0]))
			if int32(3) < v25 {
				v35 = v11
				F_finishSlotMigrationJob(m, v35, int32(18), int32(_a_F_clusterCommandSyncSlotsPaused_4))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					m.G0 = v9 + int32(32)
					return
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+188))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v28
				F__serverLog(m, int32(3), int32(_a_F_clusterCommandSyncSlotsPaused_5), v9)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					v35 = v34
					F_finishSlotMigrationJob(m, v35, int32(18), int32(_a_F_clusterCommandSyncSlotsPaused_4))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				}
			}
		case 3, 16, 17, 18:
			F__serverAssert(m, int32(_a_F_clusterCommandSyncSlotsPaused_6), int32(_a_F_clusterCommandSyncSlotsPaused_7), int32(692))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		F_addReplyError(m, l0, int32(_a_F_clusterCommandSyncSlotsPaused_8))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			m.G0 = v9 + int32(32)
			return
		}
	}
}
func F_clusterCreatePublishMsgBlock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	v10 = F_getDecodedObject(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_getDecodedObject(m, l1)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(0)
			v18 = F_objectGetVal(m, v10)
			mBase = m.M
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(-1)))))
			switch v21 & int32(7) {
			case 0:
				v38 = int32(base.Ui32(v21) >> (uint(int32(3)) % 32))
			case 1:
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(-3)))))
				v38 = v28
			case 2:
				v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18+int32(-5)))))
				v38 = v31
			case 3:
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(-9))))
				v38 = v34
			case 4:
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(-17))))
				v38 = v37
			default:
				v38 = v16
			}
			if l3 != 0 {
				v41 = int32(10)
			} else {
				v41 = int32(4)
			}
			v42 = F_objectGetVal(m, v14)
			mBase = m.M
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(-1)))))
			switch v45 & int32(7) {
			case 0:
				v62 = int32(base.Ui32(v45) >> (uint(int32(3)) % 32))
			case 1:
				v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(-3)))))
				v62 = v52
			case 2:
				v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42+int32(-5)))))
				v62 = v55
			case 3:
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(-9))))
				v62 = v58
			case 4:
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(-17))))
				v62 = v61
			default:
				v62 = v16
			}
			if l2 != 0 {
				v65 = v41 | int32(32768)
			} else {
				v65 = v41
			}
			if l2 != 0 {
				v68 = int32(24)
			} else {
				v68 = int32(2264)
			}
			v71 = F_createClusterMsgSendBlock(m, v65, v68+v38+v62)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				v73 = v71 + v68
				v74 = F___bswap_32_1(m, v38)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v73))) = v74
				v76 = F___bswap_32_1(m, v62)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v76
				v79 = v73 + int32(8)
				v80 = F_objectGetVal(m, v10)
				mBase = m.M
				v81 = int32(0)
				v83 = F_objectGetVal(m, v10)
				mBase = m.M
				v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+int32(-1)))))
				switch v86 & int32(7) {
				case 0:
					v103 = int32(base.Ui32(v86) >> (uint(int32(3)) % 32))
				case 1:
					v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+int32(-3)))))
					v103 = v93
				case 2:
					v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83+int32(-5)))))
					v103 = v96
				case 3:
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(-9))))
					v103 = v99
				case 4:
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(-17))))
					v103 = v102
				default:
					v103 = v81
				}
				if v103 == int32(0) {
					v107 = v79
				} else {
					v106 = F__emscripten_memcpy_bulkmem(m, v79, v80, v103)
					mBase = m.M
					v107 = v106
				}
				v108 = F_objectGetVal(m, v10)
				mBase = m.M
				v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+int32(-1)))))
				switch v111 & int32(7) {
				case 0:
					v128 = int32(base.Ui32(v111) >> (uint(int32(3)) % 32))
				case 1:
					v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+int32(-3)))))
					v128 = v118
				case 2:
					v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108+int32(-5)))))
					v128 = v121
				case 3:
					v124 = *(*int32)(unsafe.Add(mBase, uint32(v108+int32(-9))))
					v128 = v124
				case 4:
					v127 = *(*int32)(unsafe.Add(mBase, uint32(v108+int32(-17))))
					v128 = v127
				default:
					v128 = v81
				}
				v129 = F_objectGetVal(m, v14)
				mBase = m.M
				v131 = F_objectGetVal(m, v14)
				mBase = m.M
				v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+int32(-1)))))
				switch v134 & int32(7) {
				case 0:
					v151 = int32(base.Ui32(v134) >> (uint(int32(3)) % 32))
				case 1:
					v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+int32(-3)))))
					v151 = v141
				case 2:
					v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131+int32(-5)))))
					v151 = v144
				case 3:
					v147 = *(*int32)(unsafe.Add(mBase, uint32(v131+int32(-9))))
					v151 = v147
				case 4:
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v131+int32(-17))))
					v151 = v150
				default:
					v151 = int32(0)
				}
				if v151 == int32(0) {
				} else {
					v155 = F__emscripten_memcpy_bulkmem(m, v107+v128, v129, v151)
					mBase = m.M
				}
				F_decrRefCount(m, v10)
				mBase = m.M
				v158 = m.ExcPending
				if v158 != 0 {
					return int32(0)
				} else {
					F_decrRefCount(m, v14)
					mBase = m.M
					v160 = m.ExcPending
					if v160 != 0 {
						return int32(0)
					} else {
						return v71
					}
				}
			}
		}
	}
}
func F_clusterCron(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int64
	_ = v68
	var v72 int64
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v210 int32
	_ = v210
	var v212 int64
	_ = v212
	var v213 int32
	_ = v213
	var v218 int64
	_ = v218
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int64
	_ = v261
	var v266 int64
	_ = v266
	var v267 int64
	_ = v267
	var v274 int32
	_ = v274
	var v279 int64
	_ = v279
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int64
	_ = v318
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v344 int32
	_ = v344
	var v345 int64
	_ = v345
	var v354 int64
	_ = v354
	var v357 int64
	_ = v357
	var v358 int64
	_ = v358
	var v361 int64
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v389 int64
	_ = v389
	var v392 int64
	_ = v392
	var v393 int64
	_ = v393
	var v396 int64
	_ = v396
	var v400 int64
	_ = v400
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int64
	_ = v437
	var v440 int64
	_ = v440
	var v442 int64
	_ = v442
	var v444 int64
	_ = v444
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int64
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int64
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int64
	_ = v507
	var v510 int64
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int64
	_ = v530
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int64
	_ = v566
	var v567 int64
	_ = v567
	var v568 int64
	_ = v568
	var v569 int64
	_ = v569
	var v570 int64
	_ = v570
	var v571 int64
	_ = v571
	var v572 int64
	_ = v572
	var v574 int64
	_ = v574
	var v576 int64
	_ = v576
	var v578 int64
	_ = v578
	var v580 int64
	_ = v580
	var v582 int64
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int64
	_ = v669
	var v671 int64
	_ = v671
	var v673 int64
	_ = v673
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int64
	_ = v685
	var v686 int64
	_ = v686
	var v689 int32
	_ = v689
	var v692 int64
	_ = v692
	var v693 int32
	_ = v693
	var v694 int64
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int64
	_ = v702
	var v705 int32
	_ = v705
	var v710 int64
	_ = v710
	var v712 int64
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int64
	_ = v716
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int64
	_ = v725
	var v728 int32
	_ = v728
	var v733 int64
	_ = v733
	var v735 int64
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int64
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int64
	_ = v748
	var v751 int32
	_ = v751
	var v756 int64
	_ = v756
	var v758 int64
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int64
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int64
	_ = v773
	var v776 int32
	_ = v776
	var v779 int64
	_ = v779
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v846 int64
	_ = v846
	var v847 int64
	_ = v847
	var v848 int64
	_ = v848
	var v849 int64
	_ = v849
	var v850 int64
	_ = v850
	var v851 int64
	_ = v851
	var v852 int64
	_ = v852
	var v854 int64
	_ = v854
	var v856 int64
	_ = v856
	var v858 int64
	_ = v858
	var v860 int64
	_ = v860
	var v862 int64
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v931 int32
	_ = v931
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v951 int64
	_ = v951
	var v952 int32
	_ = v952
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v993 int32
	_ = v993
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1065 int32
	_ = v1065
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1100 int32
	_ = v1100
	var v1113 int32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1143 int64
	_ = v1143
	var v1144 int64
	_ = v1144
	var v1145 int64
	_ = v1145
	var v1146 int64
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1150 int64
	_ = v1150
	var v1153 int64
	_ = v1153
	var v1158 int64
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1166 int64
	_ = v1166
	var v1170 int64
	_ = v1170
	var v1172 int64
	_ = v1172
	var v1173 int64
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1177 int64
	_ = v1177
	var v1180 int64
	_ = v1180
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1188 int64
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1205 int64
	_ = v1205
	var v1209 int64
	_ = v1209
	var v1211 int64
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1246 int32
	_ = v1246
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1300 int64
	_ = v1300
	var v1301 int64
	_ = v1301
	var v1302 int64
	_ = v1302
	var v1303 int64
	_ = v1303
	var v1304 int64
	_ = v1304
	var v1305 int64
	_ = v1305
	var v1306 int64
	_ = v1306
	var v1308 int64
	_ = v1308
	var v1310 int64
	_ = v1310
	var v1312 int64
	_ = v1312
	var v1314 int64
	_ = v1314
	var v1316 int64
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	var v1363 int32
	_ = v1363
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1433 int64
	_ = v1433
	var v1436 int64
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1489 int32
	_ = v1489
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	v1 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(144)
	m.G0 = v24
	v26 = F_mstime(m)
	mBase = m.M
	v29 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCron[0]))
	*(*int64)(unsafe.Add(mBase, _c_F_clusterCron[0])) = v29 + int64(1)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[1]))
	if v34 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_clusterSlotMigrationCron(m)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L5
	}
L2:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[2]))
	F_updateSdsExtensionField(m, v34+int32(2312), v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	goto L1
L5:
	;
	v45 = int64(0)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+uint32(_c_F_clusterCron[4]))) = v45
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	v51 = F_dictGetSafeIterator(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v54 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCron[5]))
	if v54 < int64(2) {
		v72 = v45
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v80 = v51 + int32(20)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	if v81 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+32))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v68 = base.I64_div_u_s(base.I64_extend_i32_u((v60+v61)*int32(100)), int64(base.Ui64(v54)>>(uint(int64(1))%64)))
	v72 = v68 * int64(10)
	goto L7
L9:
	;
	F_dictReleaseIterator(m, v51)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L3
	} else {
		goto L157
	}
L10:
	;
	if v176 == int32(0) {
		goto L9
	} else {
		goto L36
	}
L11:
	;
	v87 = v80
	v88 = v84
	goto L14
L12:
	;
	v84 = int32(1)
	goto L11
L13:
	;
	v84 = int32(0)
	goto L11
L14:
	;
	switch v88 {
	case 0:
		goto L19
	default:
		goto L18
	}
L16:
	;
	v88 = int32(0)
	goto L14
L17:
	;
	goto L10
L18:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v168
	if v168 == int32(0) {
		goto L16
	} else {
		goto L35
	}
L19:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v92 != int32(-1) {
		v131 = v92
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v132 = int32(1)
	v133 = v131 + v132
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v133
	v135 = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138+v139+int32(26)))))
	if v143 == int32(255) {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if v96 != 0 {
		v131 = int32(-1)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	if v98 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	if v125 != int32(-1) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v105 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v97)+16)))
	v106 = int64(*(*int8)(unsafe.Add(mBase, uint32(v97)+27)))
	v107 = int64(*(*int32)(unsafe.Add(mBase, uint32(v97)+8)))
	v108 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v97)+12)))
	v109 = int64(*(*int8)(unsafe.Add(mBase, uint32(v97)+26)))
	v110 = int64(*(*int32)(unsafe.Add(mBase, uint32(v97)+4)))
	v111 = F_wangHash64(m, v110)
	mBase = m.M
	v113 = F_wangHash64(m, v109+v111)
	mBase = m.M
	v115 = F_wangHash64(m, v108+v113)
	mBase = m.M
	v117 = F_wangHash64(m, v107+v115)
	mBase = m.M
	v119 = F_wangHash64(m, v106+v117)
	mBase = m.M
	v121 = F_wangHash64(m, v105+v119)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v51)+24)) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v124 = v123
	goto L23
L25:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+24)))
	v103 = v101 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v97)+24)) = uint16(v103)
	v124 = v97
	goto L23
L26:
	;
	v131 = v125 + int32(-1)
	goto L20
L27:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v131 = v128
	goto L20
L28:
	;
	v158 = int32(2)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v138+v156<<(uint(v158)%32)+int32(4))))
	v87 = v163 + v157<<(uint(v158)%32)
	v88 = int32(1)
	goto L14
L29:
	;
	v147 = v135
	goto L31
L30:
	;
	v147 = v132 << (uint(v143) % 32)
	goto L31
L31:
	;
	if v133 < v147 {
		v156 = v139
		v157 = v133
		goto L28
	} else {
		goto L32
	}
L32:
	;
	if v139 != 0 {
		v176 = v135
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
	if v149 == int32(-1) {
		v176 = v135
		goto L17
	} else {
		goto L34
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51)+4)) = int64(4294967296)
	v156 = int32(1)
	v157 = int32(0)
	goto L28
L35:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v172
	v176 = v168
	goto L17
L36:
	;
	v191 = v176
	v192 = v72
	goto L37
L37:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	goto L39
L38:
	;
	goto L9
L39:
	;
	v212 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCron[6]))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v210)+2344))
	if v213 == int32(0) {
		v267 = v212
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v210)+2348))
	if v274 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L41:
	;
	if v212 == int64(0) {
		v267 = v212
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v213)+24))
	if base.Ui64(v218) <= base.Ui64(v212) {
		v267 = v212
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[7]))
	if int32(3) < v221 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_freeClusterLink(m, v213)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L3
	} else {
		goto L56
	}
L45:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v213)+48))
	if v226 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v227 = int32(_a_F_clusterCron_0)
	goto L48
L47:
	;
	v227 = int32(_a_F_clusterCron_1)
	goto L48
L48:
	;
	v228 = int32(_a_F_clusterCron_2)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v213)+44))
	if v229 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v233 = v229 + int32(8)
	goto L51
L50:
	;
	v233 = v228
	goto L51
L51:
	;
	if v229 == int32(0) {
		v238 = v228
		goto L52
	} else {
		goto L53
	}
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(128)))) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v24)+120)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v24)+116)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v24)+112)) = v227
	F__serverLog(m, int32(3), int32(_a_F_clusterCron_3), v24+int32(112))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L3
	} else {
		goto L55
	}
L53:
	;
	v236 = F_humanNodename(m, v229)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v238 = v236
	goto L52
L55:
	;
	goto L44
L56:
	;
	v259 = int32(_a_F_clusterCron_4)
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v260)+uint32(_c_F_clusterCron[8])))
	*(*int64)(unsafe.Add(mBase, uint32(v260)+uint32(_c_F_clusterCron[8]))) = v261 + int64(1)
	v266 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCron[6]))
	v267 = v266
	goto L40
L57:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_clusterCron[9])))
	if v329&int32(4) == int32(0) {
		goto L75
	} else {
		goto L76
	}
L58:
	;
	if v267 == int64(0) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v274)+24))
	if base.Ui64(v279) <= base.Ui64(v267) {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[7]))
	if int32(3) < v282 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_freeClusterLink(m, v274)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L3
	} else {
		goto L73
	}
L62:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v274)+48))
	if v287 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v288 = int32(_a_F_clusterCron_0)
	goto L65
L64:
	;
	v288 = int32(_a_F_clusterCron_1)
	goto L65
L65:
	;
	v289 = int32(_a_F_clusterCron_2)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v274)+44))
	if v290 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v294 = v290 + int32(8)
	goto L68
L67:
	;
	v294 = v289
	goto L68
L68:
	;
	if v290 == int32(0) {
		v299 = v289
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(96)))) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v24)+88)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v288
	F__serverLog(m, int32(3), int32(_a_F_clusterCron_3), v24+int32(80))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L3
	} else {
		goto L72
	}
L70:
	;
	v297 = F_humanNodename(m, v290)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	v299 = v297
	goto L69
L72:
	;
	goto L61
L73:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v317)+uint32(_c_F_clusterCron[8])))
	*(*int64)(unsafe.Add(mBase, uint32(v317)+uint32(_c_F_clusterCron[8]))) = v318 + int64(1)
	goto L57
L74:
	;
	v541 = v51 + int32(20)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	if v542 != 0 {
		goto L132
	} else {
		goto L133
	}
L75:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v210)+88))
	if v334&int32(80) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v530 = v192
	goto L74
L77:
	;
	if v334&int32(4) == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v530 = v192
	goto L74
L79:
	;
	if v334&int32(32) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v344)+uint32(_c_F_clusterCron[4])))
	*(*int64)(unsafe.Add(mBase, uint32(v344)+uint32(_c_F_clusterCron[4]))) = v345 + int64(1)
	goto L79
L81:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v210)+2344))
	if v334&int32(172) != 0 {
		v430 = v381
		goto L92
	} else {
		goto L93
	}
L82:
	;
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v210)))
	v357 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCron[5]))
	v358 = int64(1000)
	if v358 < v357 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v361 = v357
	goto L85
L84:
	;
	v361 = v358
	goto L85
L85:
	;
	if v26-v354 <= v361 {
		goto L81
	} else {
		goto L86
	}
L86:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[7]))
	if int32(3) < v364 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_clusterDelNode(m, v210)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L3
	} else {
		goto L90
	}
L88:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v210)+2332))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v210 + int32(2256)
	F__serverLog(m, int32(3), int32(_a_F_clusterCron_5), v24+int32(64))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L3
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v530 = v192
	goto L74
L91:
	;
	if v435 != 0 {
		goto L113
	} else {
		goto L114
	}
L92:
	;
	if v430 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L93:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v210)+2348))
	if v381 == int32(0) {
		v435 = v384
		goto L91
	} else {
		goto L94
	}
L94:
	;
	if v384 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v389 = *(*int64)(unsafe.Add(mBase, uint32(v210)+2240))
	v392 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCron[5]))
	v393 = int64(1000)
	if v393 < v392 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v530 = v192
	goto L74
L97:
	;
	if v334&int32(4096) != 0 {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	v396 = v392
	goto L100
L99:
	;
	v396 = v393
	goto L100
L100:
	;
	if v396 < v26-v389 {
		goto L97
	} else {
		goto L101
	}
L101:
	;
	v530 = v192
	goto L74
L102:
	;
	v400 = *(*int64)(unsafe.Add(mBase, uint32(v210)+2208))
	if v396 < v26-v400 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v530 = v192
	goto L74
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210)+88)) = v334 | int32(128)
	v407 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[7]))
	if int32(2) < v407 {
		v423 = v381
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v530 = v192
	goto L74
L106:
	;
	F_clusterSendPing(m, v423, int32(2))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L3
	} else {
		goto L110
	}
L107:
	;
	v410 = F_humanNodename(m, v210)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L3
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v210 + int32(8)
	F__serverLog(m, int32(2), int32(_a_F_clusterCron_6), v24+int32(48))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L3
	} else {
		goto L109
	}
L109:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v210)+2344))
	v423 = v422
	goto L106
L110:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v210)+2344))
	v430 = v427
	goto L92
L111:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v210)+2348))
	v435 = v433
	goto L91
L112:
	;
	v530 = v192
	goto L74
L113:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v210)+2232)) = v26
	v450 = F_valkey_malloc(m, int32(56))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L3
	} else {
		goto L117
	}
L114:
	;
	v437 = *(*int64)(unsafe.Add(mBase, uint32(v210)+2232))
	v440 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCron[5]))
	v442 = base.I64_div_s(v440, int64(20))
	if v442 <= v26-v437 {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v444 = int64(0)
	if v192 == v444 {
		v530 = v444
		goto L74
	} else {
		goto L116
	}
L116:
	;
	goto L113
L117:
	;
	v452 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v450))) = v452
	v454 = F_listCreate(m)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L3
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v450)+12)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v454)+12)) = int32(63)
	v459 = int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v450)+40)) = v459
	*(*int64)(unsafe.Add(mBase, uint32(v450)+24)) = int64(24)
	*(*int32)(unsafe.Add(mBase, uint32(v450)+16)) = int32(0)
	v466 = F_valkey_malloc(m, v459)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L3
	} else {
		goto L119
	}
L119:
	;
	v468 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v450)+36)) = v468
	*(*int32)(unsafe.Add(mBase, uint32(v450)+32)) = v466
	v471 = int32(_a_F_clusterCron_4)
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_clusterCron[10])) = v473 + int32(1048)
	*(*int32)(unsafe.Add(mBase, uint32(v450)+48)) = v468
	*(*int32)(unsafe.Add(mBase, uint32(v450)+44)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v450)+8)) = v468
	*(*int32)(unsafe.Add(mBase, uint32(v210)+2344)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v450)+52)) = v468
	v485 = F_connTypeOfCluster(m)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v485)+40))
	v488 = m.T0[v487].(func(*base.Module) int32)(m)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L3
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v450)+8)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v488)+24)) = v450
	v493 = v192 + int64(-1)
	v495 = v210 + int32(2256)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v210)+2332))
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[11]))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v488)))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v501)+56))
	v503 = m.T0[v502].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v488, v495, v496, v498, int32(0), int32(66))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L3
	} else {
		goto L122
	}
L122:
	;
	if v503 != int32(-1) {
		v530 = v493
		goto L74
	} else {
		goto L123
	}
L123:
	;
	v507 = *(*int64)(unsafe.Add(mBase, uint32(v210)+2184))
	if v507 != int64(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v513 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[7]))
	if int32(0) < v513 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v510 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v210)+2184)) = v510
	goto L124
L126:
	;
	F_freeClusterLink(m, v450)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L3
	} else {
		goto L129
	}
L127:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v210)+2332))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = int32(_a_F_clusterCron_7)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v495
	F__serverLog(m, int32(0), int32(_a_F_clusterCron_8), v24+int32(32))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v530 = v493
	goto L74
L130:
	;
	if v637 != 0 {
		v191 = v637
		v192 = v530
		goto L37
	} else {
		goto L156
	}
L131:
	;
	v548 = v541
	v549 = v545
	goto L134
L132:
	;
	v545 = int32(1)
	goto L131
L133:
	;
	v545 = int32(0)
	goto L131
L134:
	;
	switch v549 {
	case 0:
		goto L139
	default:
		goto L138
	}
L136:
	;
	v549 = int32(0)
	goto L134
L137:
	;
	goto L130
L138:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v629
	if v629 == int32(0) {
		goto L136
	} else {
		goto L155
	}
L139:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v553 != int32(-1) {
		v592 = v553
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v593 = int32(1)
	v594 = v592 + v593
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v594
	v596 = int32(0)
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599+v600+int32(26)))))
	if v604 == int32(255) {
		goto L149
	} else {
		goto L150
	}
L141:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if v557 != 0 {
		v592 = int32(-1)
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	if v559 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v585)+20))
	if v586 != int32(-1) {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v566 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v558)+16)))
	v567 = int64(*(*int8)(unsafe.Add(mBase, uint32(v558)+27)))
	v568 = int64(*(*int32)(unsafe.Add(mBase, uint32(v558)+8)))
	v569 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v558)+12)))
	v570 = int64(*(*int8)(unsafe.Add(mBase, uint32(v558)+26)))
	v571 = int64(*(*int32)(unsafe.Add(mBase, uint32(v558)+4)))
	v572 = F_wangHash64(m, v571)
	mBase = m.M
	v574 = F_wangHash64(m, v570+v572)
	mBase = m.M
	v576 = F_wangHash64(m, v569+v574)
	mBase = m.M
	v578 = F_wangHash64(m, v568+v576)
	mBase = m.M
	v580 = F_wangHash64(m, v567+v578)
	mBase = m.M
	v582 = F_wangHash64(m, v566+v580)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v51)+24)) = v582
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v585 = v584
	goto L143
L145:
	;
	v562 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v558)+24)))
	v564 = v562 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v558)+24)) = uint16(v564)
	v585 = v558
	goto L143
L146:
	;
	v592 = v586 + int32(-1)
	goto L140
L147:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v592 = v589
	goto L140
L148:
	;
	v619 = int32(2)
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v599+v617<<(uint(v619)%32)+int32(4))))
	v548 = v624 + v618<<(uint(v619)%32)
	v549 = int32(1)
	goto L134
L149:
	;
	v608 = v596
	goto L151
L150:
	;
	v608 = v593 << (uint(v604) % 32)
	goto L151
L151:
	;
	if v594 < v608 {
		v617 = v600
		v618 = v594
		goto L148
	} else {
		goto L152
	}
L152:
	;
	if v600 != 0 {
		v637 = v596
		goto L137
	} else {
		goto L153
	}
L153:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v599)+20))
	if v610 == int32(-1) {
		v637 = v596
		goto L137
	} else {
		goto L154
	}
L154:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51)+4)) = int64(4294967296)
	v617 = int32(1)
	v618 = int32(0)
	goto L148
L155:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v629)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v541))) = v633
	v637 = v629
	goto L137
L156:
	;
	goto L38
L157:
	;
	v665 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_clusterCron[9])))
	if v665&int32(2) != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v810 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v810)+32))
	v812 = F_dictGetSafeIterator(m, v811)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L3
	} else {
		goto L236
	}
L159:
	;
	v668 = int32(0)
	v669 = int64(0)
	v671 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCron[0]))
	v673 = base.I64_rem_u_s(v671, int64(10))
	if v673 != v669 {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v677 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v677)+32))
	v679 = F_dictGetRandomKey(m, v678)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L3
	} else {
		goto L162
	}
L161:
	;
	v696 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v696)+32))
	v698 = F_dictGetRandomKey(m, v697)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L3
	} else {
		goto L169
	}
L162:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v679)+8))
	goto L163
L163:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v681)+2344))
	if v682 == int32(0) {
		v693 = v668
		v694 = v669
		goto L161
	} else {
		goto L164
	}
L164:
	;
	v685 = int64(0)
	v686 = *(*int64)(unsafe.Add(mBase, uint32(v681)+2184))
	if v686 != v685 {
		v693 = v668
		v694 = v685
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681)+88)))
	if v689&int32(48) != 0 {
		v693 = v668
		v694 = v685
		goto L161
	} else {
		goto L166
	}
L166:
	;
	v692 = *(*int64)(unsafe.Add(mBase, uint32(v681)+2192))
	v693 = v681
	v694 = v692
	goto L161
L167:
	;
	v719 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)+32))
	v721 = F_dictGetRandomKey(m, v720)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L3
	} else {
		goto L186
	}
L168:
	;
	v702 = *(*int64)(unsafe.Add(mBase, uint32(v700)+2184))
	if v702 == int64(0) {
		goto L172
	} else {
		goto L173
	}
L169:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v698)+8))
	goto L170
L170:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v700)+2344))
	if v701 != 0 {
		goto L168
	} else {
		goto L171
	}
L171:
	;
	v715 = v693
	v716 = v694
	goto L167
L172:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700)+88)))
	if v705&int32(48) == int32(0) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v715 = v693
	v716 = v694
	goto L167
L174:
	;
	v710 = *(*int64)(unsafe.Add(mBase, uint32(v700)+2192))
	if v693 != 0 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v715 = v693
	v716 = v694
	goto L167
L176:
	;
	if v694 < v710 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v715 = v700
	v716 = v710
	goto L167
L178:
	;
	v712 = v694
	goto L180
L179:
	;
	v712 = v710
	goto L180
L180:
	;
	if v710 < v694 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v714 = v700
	goto L183
L182:
	;
	v714 = v693
	goto L183
L183:
	;
	v715 = v714
	v716 = v712
	goto L167
L184:
	;
	v742 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v742)+32))
	v744 = F_dictGetRandomKey(m, v743)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L3
	} else {
		goto L203
	}
L185:
	;
	v725 = *(*int64)(unsafe.Add(mBase, uint32(v723)+2184))
	if v725 == int64(0) {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v721)+8))
	goto L187
L187:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v723)+2344))
	if v724 != 0 {
		goto L185
	} else {
		goto L188
	}
L188:
	;
	v738 = v715
	v739 = v716
	goto L184
L189:
	;
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723)+88)))
	if v728&int32(48) == int32(0) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v738 = v715
	v739 = v716
	goto L184
L191:
	;
	v733 = *(*int64)(unsafe.Add(mBase, uint32(v723)+2192))
	if v715 != 0 {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v738 = v715
	v739 = v716
	goto L184
L193:
	;
	if v716 < v733 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v738 = v723
	v739 = v733
	goto L184
L195:
	;
	v735 = v716
	goto L197
L196:
	;
	v735 = v733
	goto L197
L197:
	;
	if v733 < v716 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v737 = v723
	goto L200
L199:
	;
	v737 = v715
	goto L200
L200:
	;
	v738 = v737
	v739 = v735
	goto L184
L201:
	;
	v765 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v765)+32))
	v767 = F_dictGetRandomKey(m, v766)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L3
	} else {
		goto L220
	}
L202:
	;
	v748 = *(*int64)(unsafe.Add(mBase, uint32(v746)+2184))
	if v748 == int64(0) {
		goto L206
	} else {
		goto L207
	}
L203:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v744)+8))
	goto L204
L204:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v746)+2344))
	if v747 != 0 {
		goto L202
	} else {
		goto L205
	}
L205:
	;
	v761 = v738
	v762 = v739
	goto L201
L206:
	;
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746)+88)))
	if v751&int32(48) == int32(0) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v761 = v738
	v762 = v739
	goto L201
L208:
	;
	v756 = *(*int64)(unsafe.Add(mBase, uint32(v746)+2192))
	if v738 != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v761 = v738
	v762 = v739
	goto L201
L210:
	;
	if v739 < v756 {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	v761 = v746
	v762 = v756
	goto L201
L212:
	;
	v758 = v739
	goto L214
L213:
	;
	v758 = v756
	goto L214
L214:
	;
	if v756 < v739 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v760 = v746
	goto L217
L216:
	;
	v760 = v738
	goto L217
L217:
	;
	v761 = v760
	v762 = v758
	goto L201
L218:
	;
	v785 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[7]))
	if int32(0) < v785 {
		goto L229
	} else {
		goto L230
	}
L219:
	;
	if v761 == int32(0) {
		goto L158
	} else {
		goto L228
	}
L220:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v767)+8))
	goto L221
L221:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v769)+2344))
	if v770 == int32(0) {
		goto L219
	} else {
		goto L222
	}
L222:
	;
	v773 = *(*int64)(unsafe.Add(mBase, uint32(v769)+2184))
	if v773 != int64(0) {
		goto L219
	} else {
		goto L223
	}
L223:
	;
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+88)))
	if v776&int32(48) != 0 {
		goto L219
	} else {
		goto L224
	}
L224:
	;
	if v761 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v779 = *(*int64)(unsafe.Add(mBase, uint32(v769)+2192))
	if v762 <= v779 {
		v783 = v761
		goto L218
	} else {
		goto L227
	}
L226:
	;
	v783 = v769
	goto L218
L227:
	;
	v783 = v769
	goto L218
L228:
	;
	v783 = v761
	goto L218
L229:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v783)+2344))
	F_clusterSendPing(m, v800, int32(0))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L3
	} else {
		goto L233
	}
L230:
	;
	v788 = F_humanNodename(m, v783)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L3
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v783 + int32(8)
	F__serverLog(m, int32(0), int32(_a_F_clusterCron_9), v24+int32(16))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L3
	} else {
		goto L232
	}
L232:
	;
	goto L229
L233:
	;
	goto L158
L234:
	;
	F_dictReleaseIterator(m, v812)
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L3
	} else {
		goto L350
	}
L235:
	;
	v925 = int32(0)
	v931 = v917
	v943 = v925
	v944 = v925
	v945 = v925
	v946 = v925
	goto L264
L236:
	;
	v821 = v812 + int32(20)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v812)+16))
	if v822 != 0 {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	if v917 != 0 {
		goto L235
	} else {
		goto L263
	}
L238:
	;
	v828 = v821
	v829 = v825
	goto L241
L239:
	;
	v825 = int32(1)
	goto L238
L240:
	;
	v825 = int32(0)
	goto L238
L241:
	;
	switch v829 {
	case 0:
		goto L246
	default:
		goto L245
	}
L243:
	;
	v829 = int32(0)
	goto L241
L244:
	;
	goto L237
L245:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v828)))
	*(*int32)(unsafe.Add(mBase, uint32(v812)+16)) = v909
	if v909 == int32(0) {
		goto L243
	} else {
		goto L262
	}
L246:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v812)+4))
	if v833 != int32(-1) {
		v872 = v833
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v873 = int32(1)
	v874 = v872 + v873
	*(*int32)(unsafe.Add(mBase, uint32(v812)+4)) = v874
	v876 = int32(0)
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v812)))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v812)+8))
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v879+v880+int32(26)))))
	if v884 == int32(255) {
		goto L256
	} else {
		goto L257
	}
L248:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v812)+8))
	if v837 != 0 {
		v872 = int32(-1)
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v812)))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v812)+12))
	if v839 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v865)+20))
	if v866 != int32(-1) {
		goto L253
	} else {
		goto L254
	}
L251:
	;
	v846 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v838)+16)))
	v847 = int64(*(*int8)(unsafe.Add(mBase, uint32(v838)+27)))
	v848 = int64(*(*int32)(unsafe.Add(mBase, uint32(v838)+8)))
	v849 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v838)+12)))
	v850 = int64(*(*int8)(unsafe.Add(mBase, uint32(v838)+26)))
	v851 = int64(*(*int32)(unsafe.Add(mBase, uint32(v838)+4)))
	v852 = F_wangHash64(m, v851)
	mBase = m.M
	v854 = F_wangHash64(m, v850+v852)
	mBase = m.M
	v856 = F_wangHash64(m, v849+v854)
	mBase = m.M
	v858 = F_wangHash64(m, v848+v856)
	mBase = m.M
	v860 = F_wangHash64(m, v847+v858)
	mBase = m.M
	v862 = F_wangHash64(m, v846+v860)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v812)+24)) = v862
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v812)))
	v865 = v864
	goto L250
L252:
	;
	v842 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v838)+24)))
	v844 = v842 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v838)+24)) = uint16(v844)
	v865 = v838
	goto L250
L253:
	;
	v872 = v866 + int32(-1)
	goto L247
L254:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v812)+4))
	v872 = v869
	goto L247
L255:
	;
	v899 = int32(2)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v879+v897<<(uint(v899)%32)+int32(4))))
	v828 = v904 + v898<<(uint(v899)%32)
	v829 = int32(1)
	goto L241
L256:
	;
	v888 = v876
	goto L258
L257:
	;
	v888 = v873 << (uint(v884) % 32)
	goto L258
L258:
	;
	if v874 < v888 {
		v897 = v880
		v898 = v874
		goto L255
	} else {
		goto L259
	}
L259:
	;
	if v880 != 0 {
		v917 = v876
		goto L244
	} else {
		goto L260
	}
L260:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v879)+20))
	if v890 == int32(-1) {
		v917 = v876
		goto L244
	} else {
		goto L261
	}
L261:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v812)+4)) = int64(4294967296)
	v897 = int32(1)
	v898 = int32(0)
	goto L255
L262:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v909)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v821))) = v913
	v917 = v909
	goto L244
L263:
	;
	v922 = int32(0)
	v1381 = int32(1)
	v1383 = v922
	v1393 = v922
	v1394 = v922
	goto L234
L264:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v931)+8))
	goto L266
L265:
	;
	v1375 = int32(0)
	v1381 = base.B2i32(v1263 == v1375)
	v1383 = base.B2i32(v1264 != v1375)
	v1393 = v1261
	v1394 = v1262
	goto L234
L266:
	;
	v951 = F_mstime(m)
	mBase = m.M
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v950)+88))
	if v952&int32(112) != 0 {
		v1261 = v943
		v1262 = v944
		v1263 = v945
		v1264 = v946
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1275 = v812 + int32(20)
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v812)+16))
	if v1276 != 0 {
		goto L325
	} else {
		goto L326
	}
L268:
	;
	if v952&int32(9) != int32(1) {
		v1136 = v943
		v1137 = v944
		v1139 = v946
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1143 = *(*int64)(unsafe.Add(mBase, uint32(v950)+2200))
	v1144 = v951 - v1143
	v1145 = *(*int64)(unsafe.Add(mBase, uint32(v950)+2184))
	v1146 = v951 - v1145
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v950)+2344))
	if v1147 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L270:
	;
	v959 = int32(0)
	v960 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[1]))
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v960)+88))
	if v961&int32(2) == v959 {
		v1136 = v943
		v1137 = v944
		v1139 = v946
		goto L269
	} else {
		goto L271
	}
L271:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v950)+2164))
	if v966 < int32(1) {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	if v944 < v1100 {
		goto L284
	} else {
		goto L285
	}
L273:
	;
	v1087 = int32(0)
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v950)+2160))
	if v1088 < int32(1) {
		v1100 = v1087
		v1113 = v946
		goto L272
	} else {
		goto L283
	}
L274:
	;
	v969 = int32(1)
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v950)+2168))
	v972 = int32(0)
	if v966 == v969 {
		v1032 = v972
		v1034 = v972
		goto L275
	} else {
		goto L276
	}
L275:
	;
	if v966&v969 == int32(0) {
		v1065 = v1034
		goto L280
	} else {
		goto L281
	}
L276:
	;
	v978 = int32(0)
	v983 = v978
	v985 = v978
	v993 = v978
	goto L277
L277:
	;
	v1002 = int32(2)
	v1004 = v971 + v983<<(uint(v1002)%32)
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v1004+int32(4))))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1007)+88))
	v1009 = int32(-1)
	v1011 = int32(3)
	v1013 = int32(1)
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1004)))
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1015)+88))
	v1024 = int32(base.Ui32(v1008^v1009)>>(uint(v1011)%32))&v1013 + (int32(base.Ui32(v1016^v1009)>>(uint(v1011)%32))&v1013 + v985)
	v1026 = v983 + v1002
	v1028 = v993 + v1002
	if v1028 != v966&int32(2147483646) {
		v983 = v1026
		v985 = v1024
		v993 = v1028
		goto L277
	} else {
		goto L279
	}
L278:
	;
	v1032 = v1026
	v1034 = v1024
	goto L275
L279:
	;
	goto L278
L280:
	;
	if v1065 != 0 {
		v1100 = v1065
		v1113 = v946
		goto L272
	} else {
		goto L282
	}
L281:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v971+v1032<<(uint(int32(2))%32))))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+88))
	v1065 = int32(base.Ui32(v1057^int32(-1))>>(uint(int32(3))%32))&int32(1) + v1034
	goto L280
L282:
	;
	goto L273
L283:
	;
	v1100 = v1087
	v1113 = int32(base.Ui32(v952)>>(uint(int32(8))%32))&int32(1) + v946
	goto L272
L284:
	;
	v1118 = v1100
	goto L286
L285:
	;
	v1118 = v944
	goto L286
L286:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v960)+2172))
	if v1119 == v950 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1121 = v1100
	goto L289
L288:
	;
	v1121 = v943
	goto L289
L289:
	;
	v1136 = v1121
	v1137 = v1118
	v1139 = v1113
	goto L269
L290:
	;
	v1166 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCron[12]))
	if v1166 != int64(0) {
		v1173 = v1166
		goto L297
	} else {
		goto L298
	}
L291:
	;
	v1150 = *(*int64)(unsafe.Add(mBase, uint32(v1147)))
	v1153 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCron[5]))
	if v951-v1150 <= v1153 {
		goto L290
	} else {
		goto L292
	}
L292:
	;
	if v1145 == int64(0) {
		goto L290
	} else {
		goto L293
	}
L293:
	;
	v1158 = base.I64_div_s(v1153, int64(2))
	if v1146 <= v1158 {
		goto L290
	} else {
		goto L294
	}
L294:
	;
	if v1144 <= v1158 {
		goto L290
	} else {
		goto L295
	}
L295:
	;
	F_freeClusterLink(m, v1147)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L3
	} else {
		goto L296
	}
L296:
	;
	goto L290
L297:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v950)+2344))
	if v1174 == int32(0) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v1170 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCron[5]))
	v1172 = base.I64_div_s(v1170, int64(2))
	v1173 = v1172
	goto L297
L299:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v1188 = *(*int64)(unsafe.Add(mBase, uint32(v1187)+uint32(_c_F_clusterCron[13])))
	if v1188 == int64(0) {
		goto L304
	} else {
		goto L305
	}
L300:
	;
	v1177 = *(*int64)(unsafe.Add(mBase, uint32(v950)+2184))
	if v1177 != int64(0) {
		goto L299
	} else {
		goto L301
	}
L301:
	;
	v1180 = *(*int64)(unsafe.Add(mBase, uint32(v950)+2192))
	if v951-v1180 <= v1173 {
		goto L299
	} else {
		goto L302
	}
L302:
	;
	F_clusterSendPing(m, v1174, int32(0))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L3
	} else {
		goto L303
	}
L303:
	;
	v1261 = v1136
	v1262 = v1137
	v1263 = v945
	v1264 = v1139
	goto L267
L304:
	;
	v1205 = *(*int64)(unsafe.Add(mBase, uint32(v950)+2184))
	if v1205 == int64(0) {
		v1261 = v1136
		v1262 = v1137
		v1263 = v945
		v1264 = v1139
		goto L267
	} else {
		goto L310
	}
L305:
	;
	v1191 = int32(0)
	v1192 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[1]))
	v1193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1192)+88)))
	if v1193&int32(1) == v1191 {
		goto L304
	} else {
		goto L306
	}
L306:
	;
	if v1174 == int32(0) {
		goto L304
	} else {
		goto L307
	}
L307:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1187)+uint32(_c_F_clusterCron[14])))
	if v1200 != v950 {
		goto L304
	} else {
		goto L308
	}
L308:
	;
	F_clusterSendPing(m, v1174, int32(0))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L3
	} else {
		goto L309
	}
L309:
	;
	v1261 = v1136
	v1262 = v1137
	v1263 = v945
	v1264 = v1139
	goto L267
L310:
	;
	if v1146 < v1144 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1209 = v1146
	goto L313
L312:
	;
	v1209 = v1144
	goto L313
L313:
	;
	v1211 = *(*int64)(unsafe.Add(mBase, _c_F_clusterCron[5]))
	if v1209 <= v1211 {
		v1261 = v1136
		v1262 = v1137
		v1263 = v945
		v1264 = v1139
		goto L267
	} else {
		goto L314
	}
L314:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v950)+88))
	if v1213&int32(12) != 0 {
		v1261 = v1136
		v1262 = v1137
		v1263 = v945
		v1264 = v1139
		goto L267
	} else {
		goto L315
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v950)+88)) = v1213 | int32(4)
	v1219 = int32(0)
	v1220 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[1]))
	v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1220)+88)))
	if v1221&int32(1) == v1219 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v1232 = int32(1)
	v1234 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[7]))
	if int32(2) < v1234 {
		v1261 = v1136
		v1262 = v1137
		v1263 = v1232
		v1264 = v1139
		goto L267
	} else {
		goto L320
	}
L317:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1220)+2160))
	if v1226 == int32(0) {
		goto L316
	} else {
		goto L318
	}
L318:
	;
	F_markNodeAsFailingIfNeeded(m, v950)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L3
	} else {
		goto L319
	}
L319:
	;
	v1261 = v1136
	v1262 = v1137
	v1263 = int32(1)
	v1264 = v1139
	goto L267
L320:
	;
	v1237 = F_humanNodename(m, v950)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L3
	} else {
		goto L321
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v1237
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v950 + int32(8)
	F__serverLog(m, int32(2), int32(_a_F_clusterCron_10), v24)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L3
	} else {
		goto L322
	}
L322:
	;
	v1261 = v1136
	v1262 = v1137
	v1263 = v1232
	v1264 = v1139
	goto L267
L323:
	;
	if v1371 != 0 {
		v931 = v1371
		v943 = v1261
		v944 = v1262
		v945 = v1263
		v946 = v1264
		goto L264
	} else {
		goto L349
	}
L324:
	;
	v1282 = v1275
	v1283 = v1279
	goto L327
L325:
	;
	v1279 = int32(1)
	goto L324
L326:
	;
	v1279 = int32(0)
	goto L324
L327:
	;
	switch v1283 {
	case 0:
		goto L332
	default:
		goto L331
	}
L329:
	;
	v1283 = int32(0)
	goto L327
L330:
	;
	goto L323
L331:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1282)))
	*(*int32)(unsafe.Add(mBase, uint32(v812)+16)) = v1363
	if v1363 == int32(0) {
		goto L329
	} else {
		goto L348
	}
L332:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v812)+4))
	if v1287 != int32(-1) {
		v1326 = v1287
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v1327 = int32(1)
	v1328 = v1326 + v1327
	*(*int32)(unsafe.Add(mBase, uint32(v812)+4)) = v1328
	v1330 = int32(0)
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v812)))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v812)+8))
	v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1333+v1334+int32(26)))))
	if v1338 == int32(255) {
		goto L342
	} else {
		goto L343
	}
L334:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v812)+8))
	if v1291 != 0 {
		v1326 = int32(-1)
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v812)))
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v812)+12))
	if v1293 == int32(0) {
		goto L337
	} else {
		goto L338
	}
L336:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1319)+20))
	if v1320 != int32(-1) {
		goto L339
	} else {
		goto L340
	}
L337:
	;
	v1300 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1292)+16)))
	v1301 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1292)+27)))
	v1302 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1292)+8)))
	v1303 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1292)+12)))
	v1304 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1292)+26)))
	v1305 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1292)+4)))
	v1306 = F_wangHash64(m, v1305)
	mBase = m.M
	v1308 = F_wangHash64(m, v1304+v1306)
	mBase = m.M
	v1310 = F_wangHash64(m, v1303+v1308)
	mBase = m.M
	v1312 = F_wangHash64(m, v1302+v1310)
	mBase = m.M
	v1314 = F_wangHash64(m, v1301+v1312)
	mBase = m.M
	v1316 = F_wangHash64(m, v1300+v1314)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v812)+24)) = v1316
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v812)))
	v1319 = v1318
	goto L336
L338:
	;
	v1296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1292)+24)))
	v1298 = v1296 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1292)+24)) = uint16(v1298)
	v1319 = v1292
	goto L336
L339:
	;
	v1326 = v1320 + int32(-1)
	goto L333
L340:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v812)+4))
	v1326 = v1323
	goto L333
L341:
	;
	v1353 = int32(2)
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1333+v1351<<(uint(v1353)%32)+int32(4))))
	v1282 = v1358 + v1352<<(uint(v1353)%32)
	v1283 = int32(1)
	goto L327
L342:
	;
	v1342 = v1330
	goto L344
L343:
	;
	v1342 = v1327 << (uint(v1338) % 32)
	goto L344
L344:
	;
	if v1328 < v1342 {
		v1351 = v1334
		v1352 = v1328
		goto L341
	} else {
		goto L345
	}
L345:
	;
	if v1334 != 0 {
		v1371 = v1330
		goto L330
	} else {
		goto L346
	}
L346:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+20))
	if v1344 == int32(-1) {
		v1371 = v1330
		goto L330
	} else {
		goto L347
	}
L347:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v812)+4)) = int64(4294967296)
	v1351 = int32(1)
	v1352 = int32(0)
	goto L341
L348:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1363)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1275))) = v1367
	v1371 = v1363
	goto L330
L349:
	;
	goto L265
L350:
	;
	v1402 = int32(0)
	v1403 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[1]))
	v1404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1403)+88)))
	if v1404&int32(2) == v1402 {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v1433 = *(*int64)(unsafe.Add(mBase, uint32(v1432)+uint32(_c_F_clusterCron[13])))
	if v1433 == int64(0) {
		goto L360
	} else {
		goto L361
	}
L352:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[15]))
	if v1410 != 0 {
		goto L351
	} else {
		goto L353
	}
L353:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+2172))
	if v1411 == int32(0) {
		goto L351
	} else {
		goto L354
	}
L354:
	;
	v1414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1411)+88)))
	if v1414&int32(64) != 0 {
		goto L351
	} else {
		goto L355
	}
L355:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[16]))
	if v1422 != 0 {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v1423 = int32(2328)
	goto L358
L357:
	;
	v1423 = int32(2324)
	goto L358
L358:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1411+v1423)))
	v1426 = int32(0)
	F_replicationSetPrimary(m, v1411+int32(2256), v1425, v1426, v1426)
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L3
	} else {
		goto L359
	}
L359:
	;
	goto L351
L360:
	;
	v1467 = int32(0)
	v1468 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[1]))
	v1469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1468)+88)))
	if v1469&int32(2) == v1467 {
		goto L369
	} else {
		goto L370
	}
L361:
	;
	v1436 = F_mstime(m)
	mBase = m.M
	if v1436 <= v1433 {
		goto L360
	} else {
		goto L362
	}
L362:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[7]))
	if int32(3) < v1439 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1448 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v1448)+uint32(_c_F_clusterCron[14])))
	if v1449 == int32(0) {
		v1457 = v1448
		goto L366
	} else {
		goto L367
	}
L364:
	;
	F__serverLog(m, int32(3), int32(_a_F_clusterCron_11), int32(0))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L3
	} else {
		goto L365
	}
L365:
	;
	goto L363
L366:
	;
	v1458 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1457)+uint32(_c_F_clusterCron[17]))) = v1458
	*(*int64)(unsafe.Add(mBase, uint32(v1457)+uint32(_c_F_clusterCron[13]))) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1457)+uint32(_c_F_clusterCron[18]))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1457)+uint32(_c_F_clusterCron[14]))) = v1458
	goto L360
L367:
	;
	F_unpauseActions(m, int32(2))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L3
	} else {
		goto L368
	}
L368:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v1457 = v1456
	goto L366
L369:
	;
	if v1381 == int32(0) {
		goto L380
	} else {
		goto L381
	}
L370:
	;
	F_clusterHandleManualFailover(m)
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L3
	} else {
		goto L371
	}
L371:
	;
	v1477 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_clusterCron[19])))
	if v1477&int32(2) != 0 {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v1482 = int32(1)
	if v1383&base.B2i32(v1482 < v1394) != v1482 {
		goto L369
	} else {
		goto L375
	}
L373:
	;
	F_clusterHandleReplicaFailover(m)
	mBase = m.M
	v1481 = m.ExcPending
	if v1481 != 0 {
		goto L3
	} else {
		goto L374
	}
L374:
	;
	goto L372
L375:
	;
	if v1393 != v1394 {
		goto L369
	} else {
		goto L376
	}
L376:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[20]))
	if v1489 == int32(0) {
		goto L369
	} else {
		goto L377
	}
L377:
	;
	F_clusterHandleReplicaMigration(m, v1393)
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L3
	} else {
		goto L378
	}
L378:
	;
	goto L369
L379:
	;
	m.G0 = v24 + int32(144)
	return
L380:
	;
	F_clusterUpdateState(m)
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L3
	} else {
		goto L383
	}
L381:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, _c_F_clusterCron[3]))
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1497)+16))
	if v1498 != int32(1) {
		goto L379
	} else {
		goto L382
	}
L382:
	;
	goto L380
L383:
	;
	goto L379
}
func F_clusterEncodeOpenSlotsAuxField(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
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
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
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
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int64
	_ = v314
	var v315 int64
	_ = v315
	var v316 int64
	_ = v316
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v320 int64
	_ = v320
	var v322 int64
	_ = v322
	var v324 int64
	_ = v324
	var v326 int64
	_ = v326
	var v328 int64
	_ = v328
	var v330 int64
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int64
	_ = v448
	var v449 int64
	_ = v449
	var v450 int64
	_ = v450
	var v451 int64
	_ = v451
	var v452 int64
	_ = v452
	var v453 int64
	_ = v453
	var v454 int64
	_ = v454
	var v456 int64
	_ = v456
	var v458 int64
	_ = v458
	var v460 int64
	_ = v460
	var v462 int64
	_ = v462
	var v464 int64
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l0&int32(2) == v2 {
		v532 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return v532
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_clusterEncodeOpenSlotsAuxField[0]))
	if v16 == int32(0) {
		v532 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_clusterEncodeOpenSlotsAuxField[1]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v23 = F_dictGetIterator(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_dictReleaseIterator(m, v23)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L5
	} else {
		goto L71
	}
L5:
	;
	return int32(0)
L6:
	;
	v34 = v23 + int32(20)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v130 == int32(0) {
		v272 = v19
		goto L4
	} else {
		goto L33
	}
L8:
	;
	v41 = v34
	v42 = v38
	goto L11
L9:
	;
	v38 = int32(1)
	goto L8
L10:
	;
	v38 = int32(0)
	goto L8
L11:
	;
	switch v42 {
	case 0:
		goto L16
	default:
		goto L15
	}
L13:
	;
	v42 = int32(0)
	goto L11
L14:
	;
	goto L7
L15:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v122
	if v122 == int32(0) {
		goto L13
	} else {
		goto L32
	}
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v46 != int32(-1) {
		v85 = v46
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v86 = int32(1)
	v87 = v85 + v86
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v87
	v89 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v93+int32(26)))))
	if v97 == int32(255) {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v50 != 0 {
		v85 = int32(-1)
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v52 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	if v79 != int32(-1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v59 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v51)+16)))
	v60 = int64(*(*int8)(unsafe.Add(mBase, uint32(v51)+27)))
	v61 = int64(*(*int32)(unsafe.Add(mBase, uint32(v51)+8)))
	v62 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v51)+12)))
	v63 = int64(*(*int8)(unsafe.Add(mBase, uint32(v51)+26)))
	v64 = int64(*(*int32)(unsafe.Add(mBase, uint32(v51)+4)))
	v65 = F_wangHash64(m, v64)
	mBase = m.M
	v67 = F_wangHash64(m, v63+v65)
	mBase = m.M
	v69 = F_wangHash64(m, v62+v67)
	mBase = m.M
	v71 = F_wangHash64(m, v61+v69)
	mBase = m.M
	v73 = F_wangHash64(m, v60+v71)
	mBase = m.M
	v75 = F_wangHash64(m, v59+v73)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v78 = v77
	goto L20
L22:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
	v57 = v55 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v57)
	v78 = v51
	goto L20
L23:
	;
	v85 = v79 + int32(-1)
	goto L17
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v85 = v82
	goto L17
L25:
	;
	v112 = int32(2)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v92+v110<<(uint(v112)%32)+int32(4))))
	v41 = v117 + v111<<(uint(v112)%32)
	v42 = int32(1)
	goto L11
L26:
	;
	v101 = v89
	goto L28
L27:
	;
	v101 = v86 << (uint(v97) % 32)
	goto L28
L28:
	;
	if v87 < v101 {
		v110 = v93
		v111 = v87
		goto L25
	} else {
		goto L29
	}
L29:
	;
	if v93 != 0 {
		v130 = v89
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
	if v103 == int32(-1) {
		v130 = v89
		goto L14
	} else {
		goto L31
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+4)) = int64(4294967296)
	v110 = int32(1)
	v111 = int32(0)
	goto L25
L32:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v126
	v130 = v122
	goto L14
L33:
	;
	v136 = v130
	v138 = v19
	goto L34
L34:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	goto L36
L35:
	;
	v272 = v161
	goto L4
L36:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	goto L37
L37:
	;
	if v138 != 0 {
		v145 = v138
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_clusterEncodeOpenSlotsAuxField_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v141
	v152 = F_sdscatfmt(m, v145, int32(_a_F_clusterEncodeOpenSlotsAuxField_1), v8+int32(16))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	v143 = F_sdsempty(m)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v145 = v143
	goto L38
L41:
	;
	v157 = F_sdscatlen(m, v152, v142+int32(8), int32(40))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v161 = F_sdscatlen(m, v157, int32(_a_F_clusterEncodeOpenSlotsAuxField_2), int32(1))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v170 = v23 + int32(20)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v171 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v266 != 0 {
		v136 = v266
		v138 = v161
		goto L34
	} else {
		goto L70
	}
L45:
	;
	v177 = v170
	v178 = v174
	goto L48
L46:
	;
	v174 = int32(1)
	goto L45
L47:
	;
	v174 = int32(0)
	goto L45
L48:
	;
	switch v178 {
	case 0:
		goto L53
	default:
		goto L52
	}
L50:
	;
	v178 = int32(0)
	goto L48
L51:
	;
	goto L44
L52:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v258
	if v258 == int32(0) {
		goto L50
	} else {
		goto L69
	}
L53:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v182 != int32(-1) {
		v221 = v182
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v222 = int32(1)
	v223 = v221 + v222
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v223
	v225 = int32(0)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228+v229+int32(26)))))
	if v233 == int32(255) {
		goto L63
	} else {
		goto L64
	}
L55:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v186 != 0 {
		v221 = int32(-1)
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v188 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+20))
	if v215 != int32(-1) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v195 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v187)+16)))
	v196 = int64(*(*int8)(unsafe.Add(mBase, uint32(v187)+27)))
	v197 = int64(*(*int32)(unsafe.Add(mBase, uint32(v187)+8)))
	v198 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v187)+12)))
	v199 = int64(*(*int8)(unsafe.Add(mBase, uint32(v187)+26)))
	v200 = int64(*(*int32)(unsafe.Add(mBase, uint32(v187)+4)))
	v201 = F_wangHash64(m, v200)
	mBase = m.M
	v203 = F_wangHash64(m, v199+v201)
	mBase = m.M
	v205 = F_wangHash64(m, v198+v203)
	mBase = m.M
	v207 = F_wangHash64(m, v197+v205)
	mBase = m.M
	v209 = F_wangHash64(m, v196+v207)
	mBase = m.M
	v211 = F_wangHash64(m, v195+v209)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v211
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v214 = v213
	goto L57
L59:
	;
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v187)+24)))
	v193 = v191 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v187)+24)) = uint16(v193)
	v214 = v187
	goto L57
L60:
	;
	v221 = v215 + int32(-1)
	goto L54
L61:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v221 = v218
	goto L54
L62:
	;
	v248 = int32(2)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v228+v246<<(uint(v248)%32)+int32(4))))
	v177 = v253 + v247<<(uint(v248)%32)
	v178 = int32(1)
	goto L48
L63:
	;
	v237 = v225
	goto L65
L64:
	;
	v237 = v222 << (uint(v233) % 32)
	goto L65
L65:
	;
	if v223 < v237 {
		v246 = v229
		v247 = v223
		goto L62
	} else {
		goto L66
	}
L66:
	;
	if v229 != 0 {
		v266 = v225
		goto L51
	} else {
		goto L67
	}
L67:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v228)+20))
	if v239 == int32(-1) {
		v266 = v225
		goto L51
	} else {
		goto L68
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+4)) = int64(4294967296)
	v246 = int32(1)
	v247 = int32(0)
	goto L62
L69:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v258)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v262
	v266 = v258
	goto L51
L70:
	;
	goto L35
L71:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_clusterEncodeOpenSlotsAuxField[1]))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+44))
	v280 = F_dictGetIterator(m, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L73
	}
L72:
	;
	F_dictReleaseIterator(m, v280)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L5
	} else {
		goto L138
	}
L73:
	;
	v289 = v280 + int32(20)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v280)+16))
	if v290 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	if v385 == int32(0) {
		v525 = v272
		goto L72
	} else {
		goto L100
	}
L75:
	;
	v296 = v289
	v297 = v293
	goto L78
L76:
	;
	v293 = int32(1)
	goto L75
L77:
	;
	v293 = int32(0)
	goto L75
L78:
	;
	switch v297 {
	case 0:
		goto L83
	default:
		goto L82
	}
L80:
	;
	v297 = int32(0)
	goto L78
L81:
	;
	goto L74
L82:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+16)) = v377
	if v377 == int32(0) {
		goto L80
	} else {
		goto L99
	}
L83:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	if v301 != int32(-1) {
		v340 = v301
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v341 = int32(1)
	v342 = v340 + v341
	*(*int32)(unsafe.Add(mBase, uint32(v280)+4)) = v342
	v344 = int32(0)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+v348+int32(26)))))
	if v352 == int32(255) {
		goto L93
	} else {
		goto L94
	}
L85:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	if v305 != 0 {
		v340 = int32(-1)
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v280)+12))
	if v307 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+20))
	if v334 != int32(-1) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v314 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v306)+16)))
	v315 = int64(*(*int8)(unsafe.Add(mBase, uint32(v306)+27)))
	v316 = int64(*(*int32)(unsafe.Add(mBase, uint32(v306)+8)))
	v317 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v306)+12)))
	v318 = int64(*(*int8)(unsafe.Add(mBase, uint32(v306)+26)))
	v319 = int64(*(*int32)(unsafe.Add(mBase, uint32(v306)+4)))
	v320 = F_wangHash64(m, v319)
	mBase = m.M
	v322 = F_wangHash64(m, v318+v320)
	mBase = m.M
	v324 = F_wangHash64(m, v317+v322)
	mBase = m.M
	v326 = F_wangHash64(m, v316+v324)
	mBase = m.M
	v328 = F_wangHash64(m, v315+v326)
	mBase = m.M
	v330 = F_wangHash64(m, v314+v328)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v280)+24)) = v330
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v333 = v332
	goto L87
L89:
	;
	v310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v306)+24)))
	v312 = v310 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+24)) = uint16(v312)
	v333 = v306
	goto L87
L90:
	;
	v340 = v334 + int32(-1)
	goto L84
L91:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	v340 = v337
	goto L84
L92:
	;
	v367 = int32(2)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v347+v365<<(uint(v367)%32)+int32(4))))
	v296 = v372 + v366<<(uint(v367)%32)
	v297 = int32(1)
	goto L78
L93:
	;
	v356 = v344
	goto L95
L94:
	;
	v356 = v341 << (uint(v352) % 32)
	goto L95
L95:
	;
	if v342 < v356 {
		v365 = v348
		v366 = v342
		goto L92
	} else {
		goto L96
	}
L96:
	;
	if v348 != 0 {
		v385 = v344
		goto L81
	} else {
		goto L97
	}
L97:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v347)+20))
	if v358 == int32(-1) {
		v385 = v344
		goto L81
	} else {
		goto L98
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v280)+4)) = int64(4294967296)
	v365 = int32(1)
	v366 = int32(0)
	goto L92
L99:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v377)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v381
	v385 = v377
	goto L81
L100:
	;
	v391 = v385
	v393 = v272
	goto L101
L101:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	goto L103
L102:
	;
	v525 = v414
	goto L72
L103:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v391)+8))
	goto L104
L104:
	;
	if v393 != 0 {
		v400 = v393
		goto L105
	} else {
		goto L106
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(_a_F_clusterEncodeOpenSlotsAuxField_3)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v396
	v405 = F_sdscatfmt(m, v400, int32(_a_F_clusterEncodeOpenSlotsAuxField_1), v8)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L5
	} else {
		goto L108
	}
L106:
	;
	v398 = F_sdsempty(m)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v400 = v398
	goto L105
L108:
	;
	v410 = F_sdscatlen(m, v405, v397+int32(8), int32(40))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	v414 = F_sdscatlen(m, v410, int32(_a_F_clusterEncodeOpenSlotsAuxField_2), int32(1))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v423 = v280 + int32(20)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v280)+16))
	if v424 != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	if v519 != 0 {
		v391 = v519
		v393 = v414
		goto L101
	} else {
		goto L137
	}
L112:
	;
	v430 = v423
	v431 = v427
	goto L115
L113:
	;
	v427 = int32(1)
	goto L112
L114:
	;
	v427 = int32(0)
	goto L112
L115:
	;
	switch v431 {
	case 0:
		goto L120
	default:
		goto L119
	}
L117:
	;
	v431 = int32(0)
	goto L115
L118:
	;
	goto L111
L119:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+16)) = v511
	if v511 == int32(0) {
		goto L117
	} else {
		goto L136
	}
L120:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	if v435 != int32(-1) {
		v474 = v435
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v475 = int32(1)
	v476 = v474 + v475
	*(*int32)(unsafe.Add(mBase, uint32(v280)+4)) = v476
	v478 = int32(0)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481+v482+int32(26)))))
	if v486 == int32(255) {
		goto L130
	} else {
		goto L131
	}
L122:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	if v439 != 0 {
		v474 = int32(-1)
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v280)+12))
	if v441 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+20))
	if v468 != int32(-1) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v448 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v440)+16)))
	v449 = int64(*(*int8)(unsafe.Add(mBase, uint32(v440)+27)))
	v450 = int64(*(*int32)(unsafe.Add(mBase, uint32(v440)+8)))
	v451 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v440)+12)))
	v452 = int64(*(*int8)(unsafe.Add(mBase, uint32(v440)+26)))
	v453 = int64(*(*int32)(unsafe.Add(mBase, uint32(v440)+4)))
	v454 = F_wangHash64(m, v453)
	mBase = m.M
	v456 = F_wangHash64(m, v452+v454)
	mBase = m.M
	v458 = F_wangHash64(m, v451+v456)
	mBase = m.M
	v460 = F_wangHash64(m, v450+v458)
	mBase = m.M
	v462 = F_wangHash64(m, v449+v460)
	mBase = m.M
	v464 = F_wangHash64(m, v448+v462)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v280)+24)) = v464
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v467 = v466
	goto L124
L126:
	;
	v444 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v440)+24)))
	v446 = v444 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v440)+24)) = uint16(v446)
	v467 = v440
	goto L124
L127:
	;
	v474 = v468 + int32(-1)
	goto L121
L128:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	v474 = v471
	goto L121
L129:
	;
	v501 = int32(2)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v481+v499<<(uint(v501)%32)+int32(4))))
	v430 = v506 + v500<<(uint(v501)%32)
	v431 = int32(1)
	goto L115
L130:
	;
	v490 = v478
	goto L132
L131:
	;
	v490 = v475 << (uint(v486) % 32)
	goto L132
L132:
	;
	if v476 < v490 {
		v499 = v482
		v500 = v476
		goto L129
	} else {
		goto L133
	}
L133:
	;
	if v482 != 0 {
		v519 = v478
		goto L118
	} else {
		goto L134
	}
L134:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v481)+20))
	if v492 == int32(-1) {
		v519 = v478
		goto L118
	} else {
		goto L135
	}
L135:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v280)+4)) = int64(4294967296)
	v499 = int32(1)
	v500 = int32(0)
	goto L129
L136:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v511)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v423))) = v515
	v519 = v511
	goto L118
L137:
	;
	goto L102
L138:
	;
	v532 = v525
	goto L1
}
func F_clusterGenNodesDescription(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
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
	var v56 int64
	_ = v56
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
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v197 int64
	_ = v197
	var v199 int64
	_ = v199
	var v201 int64
	_ = v201
	var v203 int64
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	v8 = F_sdsempty(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_clusterGenNodesSlotsInfo(m, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_clusterGenNodesDescription[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v17 = F_dictGetSafeIterator(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_dictReleaseIterator(m, v17)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L70
	}
L5:
	;
	v26 = v17 + int32(20)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v122 == int32(0) {
		v265 = v8
		goto L4
	} else {
		goto L32
	}
L7:
	;
	v33 = v26
	v34 = v30
	goto L10
L8:
	;
	v30 = int32(1)
	goto L7
L9:
	;
	v30 = int32(0)
	goto L7
L10:
	;
	switch v34 {
	case 0:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v34 = int32(0)
	goto L10
L13:
	;
	goto L6
L14:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v114
	if v114 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v38 != int32(-1) {
		v77 = v38
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v78 = int32(1)
	v79 = v77 + v78
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v79
	v81 = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v85+int32(26)))))
	if v89 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v42 != 0 {
		v77 = int32(-1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v44 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	if v71 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v51 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43)+16)))
	v52 = int64(*(*int8)(unsafe.Add(mBase, uint32(v43)+27)))
	v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(v43)+8)))
	v54 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43)+12)))
	v55 = int64(*(*int8)(unsafe.Add(mBase, uint32(v43)+26)))
	v56 = int64(*(*int32)(unsafe.Add(mBase, uint32(v43)+4)))
	v57 = F_wangHash64(m, v56)
	mBase = m.M
	v59 = F_wangHash64(m, v55+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v54+v59)
	mBase = m.M
	v63 = F_wangHash64(m, v53+v61)
	mBase = m.M
	v65 = F_wangHash64(m, v52+v63)
	mBase = m.M
	v67 = F_wangHash64(m, v51+v65)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v70 = v69
	goto L19
L21:
	;
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+24)))
	v49 = v47 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+24)) = uint16(v49)
	v70 = v43
	goto L19
L22:
	;
	v77 = v71 + int32(-1)
	goto L16
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v77 = v74
	goto L16
L24:
	;
	v104 = int32(2)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v84+v102<<(uint(v104)%32)+int32(4))))
	v33 = v109 + v103<<(uint(v104)%32)
	v34 = int32(1)
	goto L10
L25:
	;
	v93 = v81
	goto L27
L26:
	;
	v93 = v78 << (uint(v89) % 32)
	goto L27
L27:
	;
	if v79 < v93 {
		v102 = v85
		v103 = v79
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v85 != 0 {
		v122 = v81
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	if v95 == int32(-1) {
		v122 = v81
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+4)) = int64(4294967296)
	v102 = int32(1)
	v103 = int32(0)
	goto L24
L31:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v118
	v122 = v114
	goto L13
L32:
	;
	v131 = v8
	v133 = v122
	goto L33
L33:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	goto L36
L34:
	;
	v265 = v153
	goto L4
L35:
	;
	v162 = v17 + int32(20)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v163 != 0 {
		goto L45
	} else {
		goto L46
	}
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+88))
	if v136&l1 != 0 {
		v153 = v131
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v138 = F_clusterGenNodeDescription(m, l0, v135, l2)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v140 = F_sdscatsds(m, v131, v138)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_sdsfree(m, v138)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v146 = F_sdscatlen(m, v140, int32(_a_F_clusterGenNodesDescription_0), int32(1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v135)+2152))
	F_valkey_free(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v135)+2152)) = int64(0)
	v153 = v146
	goto L35
L43:
	;
	if v258 != 0 {
		v131 = v153
		v133 = v258
		goto L33
	} else {
		goto L69
	}
L44:
	;
	v169 = v162
	v170 = v166
	goto L47
L45:
	;
	v166 = int32(1)
	goto L44
L46:
	;
	v166 = int32(0)
	goto L44
L47:
	;
	switch v170 {
	case 0:
		goto L52
	default:
		goto L51
	}
L49:
	;
	v170 = int32(0)
	goto L47
L50:
	;
	goto L43
L51:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v250
	if v250 == int32(0) {
		goto L49
	} else {
		goto L68
	}
L52:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v174 != int32(-1) {
		v213 = v174
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v214 = int32(1)
	v215 = v213 + v214
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v215
	v217 = int32(0)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220+v221+int32(26)))))
	if v225 == int32(255) {
		goto L62
	} else {
		goto L63
	}
L54:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v178 != 0 {
		v213 = int32(-1)
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v180 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+20))
	if v207 != int32(-1) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v187 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v179)+16)))
	v188 = int64(*(*int8)(unsafe.Add(mBase, uint32(v179)+27)))
	v189 = int64(*(*int32)(unsafe.Add(mBase, uint32(v179)+8)))
	v190 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v179)+12)))
	v191 = int64(*(*int8)(unsafe.Add(mBase, uint32(v179)+26)))
	v192 = int64(*(*int32)(unsafe.Add(mBase, uint32(v179)+4)))
	v193 = F_wangHash64(m, v192)
	mBase = m.M
	v195 = F_wangHash64(m, v191+v193)
	mBase = m.M
	v197 = F_wangHash64(m, v190+v195)
	mBase = m.M
	v199 = F_wangHash64(m, v189+v197)
	mBase = m.M
	v201 = F_wangHash64(m, v188+v199)
	mBase = m.M
	v203 = F_wangHash64(m, v187+v201)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v203
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v206 = v205
	goto L56
L58:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v179)+24)))
	v185 = v183 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v179)+24)) = uint16(v185)
	v206 = v179
	goto L56
L59:
	;
	v213 = v207 + int32(-1)
	goto L53
L60:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v213 = v210
	goto L53
L61:
	;
	v240 = int32(2)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v220+v238<<(uint(v240)%32)+int32(4))))
	v169 = v245 + v239<<(uint(v240)%32)
	v170 = int32(1)
	goto L47
L62:
	;
	v229 = v217
	goto L64
L63:
	;
	v229 = v214 << (uint(v225) % 32)
	goto L64
L64:
	;
	if v215 < v229 {
		v238 = v221
		v239 = v215
		goto L61
	} else {
		goto L65
	}
L65:
	;
	if v221 != 0 {
		v258 = v217
		goto L50
	} else {
		goto L66
	}
L66:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v220)+20))
	if v231 == int32(-1) {
		v258 = v217
		goto L50
	} else {
		goto L67
	}
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+4)) = int64(4294967296)
	v238 = int32(1)
	v239 = int32(0)
	goto L61
L68:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v254
	v258 = v250
	goto L50
L69:
	;
	goto L34
L70:
	;
	return v265
}
func F_clusterGetTotalSlotExportBufferMemory(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_clusterGetTotalSlotExportBufferMemory[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_clusterGetTotalSlotExportBufferMemory[1])))
	v12 = v6 + int32(8)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v13
	v17 = int32(0)
	v19 = v6 + int32(8)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v21 == v17 {
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v21+base.B2i32(v24 == int32(0))<<(uint(int32(2))%32))))
		*(*int32)(unsafe.Add(mBase, uint32(v19))) = v30
	}
	if v21 == int32(0) {
		v126 = v17
	} else {
		v35 = v17
		v36 = v21
		for {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
			if v38 != 0 {
				v109 = v35
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+152))
				if v39 == int32(0) {
					v109 = v35
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)+200))
					if v44&int32(1) != 0 {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v39)+132))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v39)+160))
						v67 = v63*int32(28) + v66
						v68 = *(*int64)(unsafe.Add(mBase, uint32(v39)+384))
						if v68 == int64(-1) {
							v78 = v67
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v39)+376))
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
							v78 = v67 + base.I32_wrap_i64(v68) + v74*int32(28)
						}
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v39)+272))
						v107 = v79 + v78
					} else {
						if v44&int32(2) == int32(0) {
							if v44&int32(262144) != 0 {
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v39)+216))
								if v57 == int32(0) {
								} else {
									v60 = F_isImportSlotMigrationJob(m, v57)
									mBase = m.M
								}
							}
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v39)+132))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v39)+160))
							v67 = v63*int32(28) + v66
							v68 = *(*int64)(unsafe.Add(mBase, uint32(v39)+384))
							if v68 == int64(-1) {
								v78 = v67
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v39)+376))
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
								v78 = v67 + base.I32_wrap_i64(v68) + v74*int32(28)
							}
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v39)+272))
							v107 = v79 + v78
						} else {
							if v44&int32(4) == int32(0) {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v39)+104))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+184))
								if v82 != 0 {
									v85 = *(*int32)(unsafe.Add(mBase, _c_F_clusterGetTotalSlotExportBufferMemory[2]))
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
									v88 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
									v89 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v87)+24)))
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
									v92 = *(*int64)(unsafe.Add(mBase, uint32(v91)+16))
									v95 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
									v96 = *(*int64)(unsafe.Add(mBase, uint32(v91)+8))
									v99 = int32(44)
									v107 = base.I32_wrap_i64(v88+v89-v92) + base.I32_wrap_i64(v95-v96)*v99 + v99
								} else {
									v107 = int32(0)
								}
							} else {
								if v44&int32(262144) != 0 {
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v39)+216))
									if v57 == int32(0) {
									} else {
										v60 = F_isImportSlotMigrationJob(m, v57)
										mBase = m.M
									}
								}
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v39)+132))
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v39)+160))
								v67 = v63*int32(28) + v66
								v68 = *(*int64)(unsafe.Add(mBase, uint32(v39)+384))
								if v68 == int64(-1) {
									v78 = v67
								} else {
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v39)+376))
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
									v78 = v67 + base.I32_wrap_i64(v68) + v74*int32(28)
								}
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v39)+272))
								v107 = v79 + v78
							}
						}
					}
					v109 = v107 + v35
				}
			}
			v112 = v6 + int32(8)
			v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
			if v114 == int32(0) {
			} else {
				v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
				v123 = *(*int32)(unsafe.Add(mBase, uint32(v114+base.B2i32(v117 == int32(0))<<(uint(int32(2))%32))))
				*(*int32)(unsafe.Add(mBase, uint32(v112))) = v123
			}
			if v114 != 0 {
				v35 = v109
				v36 = v114
				continue
			} else {
				break
			}
			break
		}
		v126 = v109
	}
	m.G0 = v6 + int32(16)
	return v126
}
func F_clusterHandleReplicaFailover(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v106 int32
	_ = v106
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int64
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v230 int64
	_ = v230
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int64
	_ = v252
	var v257 int64
	_ = v257
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int64
	_ = v286
	var v288 int64
	_ = v288
	var v290 int32
	_ = v290
	var v293 int64
	_ = v293
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v319 int64
	_ = v319
	var v329 int32
	_ = v329
	var v333 int64
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int64
	_ = v340
	var v342 int32
	_ = v342
	var v345 int64
	_ = v345
	var v347 int32
	_ = v347
	var v348 int64
	_ = v348
	var v351 int64
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int64
	_ = v360
	var v365 int64
	_ = v365
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int64
	_ = v388
	var v391 int64
	_ = v391
	var v395 int32
	_ = v395
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int64
	_ = v429
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int64
	_ = v439
	var v441 int64
	_ = v441
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int64
	_ = v452
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int64
	_ = v492
	var v493 int64
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int64
	_ = v510
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = F_mstime(m)
	mBase = m.M
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v21 = base.I32_div_s(v19, int32(2))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_clusterHandleReplicaFailover[1])))
	if v23 == int64(0) {
		v29 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_clusterHandleReplicaFailover[3])))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_clusterHandleReplicaFailover[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_clusterHandleReplicaFailover[4]))) = v31 & int32(-2)
	v36 = *(*int64)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[5]))
	v38 = base.I64_div_s(v36, int64(30))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[6]))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+88)))
	if v41&int32(1) != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_clusterHandleReplicaFailover[2])))
	v29 = base.B2i32(v26 != int32(0))
	goto L1
L3:
	;
	m.G0 = v14 + int32(96)
	return
L4:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[9]))
	if v66 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_clusterHandleReplicaFailover[7]))) = int32(0)
	goto L3
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+2172))
	if v44 == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+88)))
	if (int32(base.Ui32(v47)>>(uint(int32(3))%32))|v29)&int32(1) == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[8]))
	if (base.B2i32(v56 == int32(0))|v29)&int32(1) != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	if v36 < int64(15000) {
		goto L21
	} else {
		goto L22
	}
L11:
	;
	v69 = int32(_a_F_clusterHandleReplicaFailover_0)
	v70 = *(*int64)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[24]))
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[18]))
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[25]))
	if v79 == int32(14) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v82 = v72 + int32(88)
	goto L14
L13:
	;
	v82 = int32(_a_F_clusterHandleReplicaFailover_6)
	goto L14
L14:
	;
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
	v86 = (v70 - v83) * int64(1000)
	v88 = *(*int64)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[5]))
	if v88 < v86 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v91 = v88
	goto L17
L16:
	;
	v91 = int64(0)
	goto L17
L17:
	;
	v94 = int64(*(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[26])))
	if (base.B2i32(v86-v91 <= v94*int64(1000)+v88*base.I64_extend_i32_s(v66))|v29)&int32(1) != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	F_clusterLogCantFailover(m, int32(1))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	goto L3
L21:
	;
	v112 = v38
	goto L23
L22:
	;
	v112 = int64(500)
	goto L23
L23:
	;
	v113 = v16 - v30
	v115 = v36 << (uint(int64(1)) % 64)
	v116 = int64(2000)
	if v116 < v115 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v342)+uint32(_c_F_clusterHandleReplicaFailover[22])))
	if v347 != 0 {
		v425 = v342
		goto L84
	} else {
		goto L85
	}
L25:
	;
	v126 = int64(0)
	if v112 == v126 {
		v180 = v126
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v119 = v115
	goto L28
L27:
	;
	v119 = v116
	goto L28
L28:
	;
	if v119<<(uint(int64(1))%64) < v113 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	v342 = v124
	v345 = v113
	goto L24
L30:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_clusterHandleReplicaFailover[14]))) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_clusterHandleReplicaFailover[3]))) = v112 + v16 + v180
	v187 = F_clusterGetReplicaRank(m)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L19
	} else {
		goto L42
	}
L31:
	;
	v129 = int32(0)
	F___lock(m, int32(9116960))
	mBase = m.M
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[10]))
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[11]))
	if v138 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v179 = base.I64_rem_s(base.I64_extend_i32_s(v173), v112)
	v180 = v179
	goto L30
L33:
	;
	F___unlock(m, int32(9116960))
	mBase = m.M
	goto L32
L34:
	;
	v142 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[12]))
	v144 = int32(2)
	v146 = v136 + v143<<(uint(v144)%32)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[13]))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v136+v149<<(uint(v144)%32))))
	v154 = v147 + v153
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v154
	v159 = v149 + int32(1)
	if v159 == v138 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v140 = F_lcg31(m, v139)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v140
	v173 = v140
	goto L33
L36:
	;
	v161 = v142
	goto L38
L37:
	;
	v161 = v159
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[13])) = v161
	v163 = int32(0)
	v166 = v143 + int32(1)
	if v166 == v138 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v168 = v163
	goto L41
L40:
	;
	v168 = v166
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[12])) = v168
	v173 = int32(base.Ui32(v154) >> (uint(int32(1)) % 32))
	goto L33
L42:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+uint32(_c_F_clusterHandleReplicaFailover[15]))) = v187
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v190)+uint32(_c_F_clusterHandleReplicaFailover[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v190)+uint32(_c_F_clusterHandleReplicaFailover[3]))) = v192 + v112*base.I64_extend_i32_s(v187)<<(uint(int64(1))%64)
	v199 = int32(0)
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[6]))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+88))
	if v201&int32(16) == v199 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v237 != int64(0) {
		goto L57
	} else {
		goto L58
	}
L44:
	;
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v200)+2248))
	v237 = v236
	goto L43
L45:
	;
	if v201&int32(2) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v235 = *(*int64)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[16]))
	v237 = v235
	goto L43
L47:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[17]))
	if v214 == int32(0) {
		v228 = int64(0)
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v237 = v233
	goto L43
L49:
	;
	v230 = int64(0)
	if v230 < v228 {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[18]))
	if v218 != 0 {
		v225 = v218
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+104))
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
	v228 = v227
	goto L49
L52:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[19]))
	if v221 == int32(0) {
		v228 = int64(0)
		goto L49
	} else {
		goto L53
	}
L53:
	;
	v225 = v221
	goto L51
L54:
	;
	v233 = v228
	goto L56
L55:
	;
	v233 = v230
	goto L56
L56:
	;
	goto L48
L57:
	;
	v247 = F_clusterGetFailedPrimaryRank(m)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L19
	} else {
		goto L59
	}
L58:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v241)+uint32(_c_F_clusterHandleReplicaFailover[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v241)+uint32(_c_F_clusterHandleReplicaFailover[3]))) = v242 + int64(500)
	goto L57
L59:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v250)+uint32(_c_F_clusterHandleReplicaFailover[20]))) = v247
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v250)+uint32(_c_F_clusterHandleReplicaFailover[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+uint32(_c_F_clusterHandleReplicaFailover[3]))) = v252 + v112*base.I64_extend_i32_s(v247)
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v250)+uint32(_c_F_clusterHandleReplicaFailover[1])))
	if v257 == int64(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	F_clusterBroadcastPong(m, int32(1))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L19
	} else {
		goto L82
	}
L61:
	;
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_clusterHandleReplicaFailover[3])))
	if v286 == v16 {
		goto L68
	} else {
		goto L69
	}
L62:
	;
	v263 = F_myselfIsBestRankedReplica(m)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L19
	} else {
		goto L64
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v250)+uint32(_c_F_clusterHandleReplicaFailover[15]))) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v250)+uint32(_c_F_clusterHandleReplicaFailover[3]))) = v16
	v282 = v250
	goto L61
L64:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	if v263 == int32(0) {
		v282 = v266
		goto L61
	} else {
		goto L65
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v266)+uint32(_c_F_clusterHandleReplicaFailover[3]))) = v16
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[21]))
	if int32(2) < v272 {
		v333 = int64(0)
		goto L60
	} else {
		goto L66
	}
L66:
	;
	F__serverLog(m, int32(2), int32(_a_F_clusterHandleReplicaFailover_7), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L19
	} else {
		goto L67
	}
L67:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	v282 = v281
	goto L61
L68:
	;
	v288 = int64(0)
	goto L70
L69:
	;
	v288 = v113
	goto L70
L70:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[21]))
	if int32(2) < v290 {
		v333 = v288
		goto L60
	} else {
		goto L71
	}
L71:
	;
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_clusterHandleReplicaFailover[15])))
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[17]))
	if v300 == int32(0) {
		v314 = int64(0)
		goto L73
	} else {
		goto L74
	}
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(80)))) = v319
	*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v293
	*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v286 - v16
	F__serverLog(m, int32(2), int32(_a_F_clusterHandleReplicaFailover_8), v14+int32(64))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L19
	} else {
		goto L81
	}
L73:
	;
	v316 = int64(0)
	if v316 < v314 {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[18]))
	if v304 != 0 {
		v311 = v304
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+104))
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v312)+48))
	v314 = v313
	goto L73
L76:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[19]))
	if v307 == int32(0) {
		v314 = int64(0)
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v311 = v307
	goto L75
L78:
	;
	v319 = v314
	goto L80
L79:
	;
	v319 = v316
	goto L80
L80:
	;
	goto L72
L81:
	;
	v333 = v288
	goto L60
L82:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	v340 = *(*int64)(unsafe.Add(mBase, uint32(v339)+uint32(_c_F_clusterHandleReplicaFailover[3])))
	if v16 < v340 {
		goto L3
	} else {
		goto L83
	}
L83:
	;
	v342 = v339
	v345 = v333
	goto L24
L84:
	;
	v429 = *(*int64)(unsafe.Add(mBase, uint32(v425)+uint32(_c_F_clusterHandleReplicaFailover[3])))
	if v429 <= v16 {
		goto L102
	} else {
		goto L103
	}
L85:
	;
	v348 = *(*int64)(unsafe.Add(mBase, uint32(v342)+uint32(_c_F_clusterHandleReplicaFailover[1])))
	if v348 != int64(0) {
		v425 = v342
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v342)+uint32(_c_F_clusterHandleReplicaFailover[3])))
	if v351 == v16 {
		v425 = v342
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v353 = F_clusterGetReplicaRank(m)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L19
	} else {
		goto L89
	}
L88:
	;
	v381 = F_clusterGetFailedPrimaryRank(m)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L19
	} else {
		goto L94
	}
L89:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+uint32(_c_F_clusterHandleReplicaFailover[15])))
	if v353 == v357 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v356)+uint32(_c_F_clusterHandleReplicaFailover[15]))) = v353
	v360 = *(*int64)(unsafe.Add(mBase, uint32(v356)+uint32(_c_F_clusterHandleReplicaFailover[3])))
	v365 = v112 * base.I64_extend_i32_s(v353-v357) << (uint(int64(1)) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(v356)+uint32(_c_F_clusterHandleReplicaFailover[3]))) = v360 + v365
	v369 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[21]))
	if int32(2) < v369 {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v353
	F__serverLog(m, int32(2), int32(_a_F_clusterHandleReplicaFailover_9), v14+int32(48))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L19
	} else {
		goto L92
	}
L92:
	;
	goto L88
L93:
	;
	v407 = F_myselfIsBestRankedReplica(m)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L19
	} else {
		goto L98
	}
L94:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v384)+uint32(_c_F_clusterHandleReplicaFailover[20])))
	if v381 == v385 {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384)+uint32(_c_F_clusterHandleReplicaFailover[20]))) = v381
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v384)+uint32(_c_F_clusterHandleReplicaFailover[3])))
	v391 = v112 * base.I64_extend_i32_s(v381-v385)
	*(*int64)(unsafe.Add(mBase, uint32(v384)+uint32(_c_F_clusterHandleReplicaFailover[3]))) = v388 + v391
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[21]))
	if int32(2) < v395 {
		goto L93
	} else {
		goto L96
	}
L96:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v381
	F__serverLog(m, int32(2), int32(_a_F_clusterHandleReplicaFailover_5), v14+int32(32))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L19
	} else {
		goto L97
	}
L97:
	;
	goto L93
L98:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	if v407 == int32(0) {
		v425 = v410
		goto L84
	} else {
		goto L99
	}
L99:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v410)+uint32(_c_F_clusterHandleReplicaFailover[3]))) = v16
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[21]))
	if int32(2) < v415 {
		v425 = v410
		goto L84
	} else {
		goto L100
	}
L100:
	;
	F__serverLog(m, int32(2), int32(_a_F_clusterHandleReplicaFailover_4), int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L19
	} else {
		goto L101
	}
L101:
	;
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	v425 = v424
	goto L84
L102:
	;
	if v345 <= v119 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	F_clusterLogCantFailover(m, int32(2))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L19
	} else {
		goto L104
	}
L104:
	;
	goto L3
L105:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v425)+uint32(_c_F_clusterHandleReplicaFailover[22])))
	if v438 != 0 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	F_clusterLogCantFailover(m, int32(3))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L19
	} else {
		goto L107
	}
L107:
	;
	goto L3
L108:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v425)+uint32(_c_F_clusterHandleReplicaFailover[14])))
	if v476 <= v21 {
		goto L118
	} else {
		goto L119
	}
L109:
	;
	v439 = *(*int64)(unsafe.Add(mBase, uint32(v425)+8))
	v441 = v439 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v425)+uint32(_c_F_clusterHandleReplicaFailover[23]))) = v441
	*(*int64)(unsafe.Add(mBase, uint32(v425)+8)) = v441
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[21]))
	if int32(2) < v445 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	F_clusterRequestFailoverAuth(m)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L19
	} else {
		goto L116
	}
L111:
	;
	v449 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[6]))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)+2172))
	if v450 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v451 = v450
	goto L114
L113:
	;
	v451 = v449
	goto L114
L114:
	;
	v452 = *(*int64)(unsafe.Add(mBase, uint32(v451)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v441
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v452
	F__serverLog(m, int32(2), int32(_a_F_clusterHandleReplicaFailover_3), v14)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L19
	} else {
		goto L115
	}
L115:
	;
	goto L110
L116:
	;
	v465 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v465)+uint32(_c_F_clusterHandleReplicaFailover[22]))) = int32(1)
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L19
	} else {
		goto L117
	}
L117:
	;
	v471 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)+uint32(_c_F_clusterHandleReplicaFailover[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v471)+uint32(_c_F_clusterHandleReplicaFailover[4]))) = v472 | int32(14)
	goto L3
L118:
	;
	F_clusterLogCantFailover(m, int32(4))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L19
	} else {
		goto L129
	}
L119:
	;
	v479 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[21]))
	if int32(2) < v479 {
		v489 = v425
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v491 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[6]))
	v492 = *(*int64)(unsafe.Add(mBase, uint32(v491)+96))
	v493 = *(*int64)(unsafe.Add(mBase, uint32(v489)+uint32(_c_F_clusterHandleReplicaFailover[23])))
	if base.Ui64(v493) <= base.Ui64(v492) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	F__serverLog(m, int32(2), int32(_a_F_clusterHandleReplicaFailover_2), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L19
	} else {
		goto L122
	}
L122:
	;
	v488 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	v489 = v488
	goto L120
L123:
	;
	F_clusterFailoverReplaceYourPrimary(m)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L19
	} else {
		goto L128
	}
L124:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v491)+96)) = v493
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L19
	} else {
		goto L125
	}
L125:
	;
	v498 = int32(_a_F_clusterHandleReplicaFailover_0)
	v499 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[0]))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v499)+uint32(_c_F_clusterHandleReplicaFailover[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v499)+uint32(_c_F_clusterHandleReplicaFailover[4]))) = v500 | int32(44)
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[21]))
	if int32(2) < v505 {
		goto L123
	} else {
		goto L126
	}
L126:
	;
	v509 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaFailover[6]))
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v509)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v510
	F__serverLog(m, int32(2), int32(_a_F_clusterHandleReplicaFailover_1), v14+int32(16))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L19
	} else {
		goto L127
	}
L127:
	;
	goto L123
L128:
	;
	goto L3
L129:
	;
	goto L3
}
func F_clusterHandleReplicaMigration(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
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
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
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
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int64
	_ = v370
	var v373 int64
	_ = v373
	var v380 int32
	_ = v380
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int64
	_ = v550
	var v551 int64
	_ = v551
	var v552 int64
	_ = v552
	var v553 int64
	_ = v553
	var v554 int64
	_ = v554
	var v555 int64
	_ = v555
	var v556 int64
	_ = v556
	var v558 int64
	_ = v558
	var v560 int64
	_ = v560
	var v562 int64
	_ = v562
	var v564 int64
	_ = v564
	var v566 int64
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v628 int32
	_ = v628
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v646 int64
	_ = v646
	var v647 int64
	_ = v647
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaMigration[0]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return
L2:
	;
	v22 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaMigration[1]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+2172))
	if v24 == v22 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+2164))
	if int32(1) <= v27 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaMigration[2]))
	if v112 <= v122 {
		goto L1
	} else {
		goto L14
	}
L5:
	;
	v31 = int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+2168))
	if v27 != v31 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v112 = int32(0)
	goto L4
L7:
	;
	if v27&v31 == int32(0) {
		v112 = v86
		goto L4
	} else {
		goto L13
	}
L8:
	;
	v40 = int32(0)
	v47 = v40
	v48 = v40
	v52 = v40
	goto L10
L9:
	;
	v36 = int32(0)
	v85 = v36
	v86 = v36
	goto L7
L10:
	;
	v57 = int32(2)
	v59 = v33 + v47<<(uint(v57)%32)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+88)))
	v62 = int32(12)
	v64 = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59+int32(4))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+88)))
	v75 = v48 + base.B2i32(v61&v62 == v64) + base.B2i32(v70&v62 == v64)
	v77 = v47 + v57
	v79 = v52 + v57
	if v79 != v27&int32(2147483646) {
		v47 = v77
		v48 = v75
		v52 = v79
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v85 = v77
	v86 = v75
	goto L7
L12:
	;
	goto L11
L13:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v33+v85<<(uint(int32(2))%32))))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+88)))
	v112 = v86 + base.B2i32(v101&int32(12) == int32(0))
	goto L4
L14:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v125 = F_dictGetSafeIterator(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	F_dictReleaseIterator(m, v125)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L17
	} else {
		goto L121
	}
L16:
	;
	v239 = v23
	v240 = v230
	v248 = int32(0)
	goto L46
L17:
	;
	return
L18:
	;
	v134 = v125 + int32(20)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	if v135 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v230 != 0 {
		goto L16
	} else {
		goto L45
	}
L20:
	;
	v141 = v134
	v142 = v138
	goto L23
L21:
	;
	v138 = int32(1)
	goto L20
L22:
	;
	v138 = int32(0)
	goto L20
L23:
	;
	switch v142 {
	case 0:
		goto L28
	default:
		goto L27
	}
L25:
	;
	v142 = int32(0)
	goto L23
L26:
	;
	goto L19
L27:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+16)) = v222
	if v222 == int32(0) {
		goto L25
	} else {
		goto L44
	}
L28:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v146 != int32(-1) {
		v185 = v146
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v186 = int32(1)
	v187 = v185 + v186
	*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = v187
	v189 = int32(0)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192+v193+int32(26)))))
	if v197 == int32(255) {
		goto L38
	} else {
		goto L39
	}
L30:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	if v150 != 0 {
		v185 = int32(-1)
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	if v152 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+20))
	if v179 != int32(-1) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v159 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v151)+16)))
	v160 = int64(*(*int8)(unsafe.Add(mBase, uint32(v151)+27)))
	v161 = int64(*(*int32)(unsafe.Add(mBase, uint32(v151)+8)))
	v162 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v151)+12)))
	v163 = int64(*(*int8)(unsafe.Add(mBase, uint32(v151)+26)))
	v164 = int64(*(*int32)(unsafe.Add(mBase, uint32(v151)+4)))
	v165 = F_wangHash64(m, v164)
	mBase = m.M
	v167 = F_wangHash64(m, v163+v165)
	mBase = m.M
	v169 = F_wangHash64(m, v162+v167)
	mBase = m.M
	v171 = F_wangHash64(m, v161+v169)
	mBase = m.M
	v173 = F_wangHash64(m, v160+v171)
	mBase = m.M
	v175 = F_wangHash64(m, v159+v173)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v125)+24)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v178 = v177
	goto L32
L34:
	;
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+24)))
	v157 = v155 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v151)+24)) = uint16(v157)
	v178 = v151
	goto L32
L35:
	;
	v185 = v179 + int32(-1)
	goto L29
L36:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v185 = v182
	goto L29
L37:
	;
	v212 = int32(2)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v192+v210<<(uint(v212)%32)+int32(4))))
	v141 = v217 + v211<<(uint(v212)%32)
	v142 = int32(1)
	goto L23
L38:
	;
	v201 = v189
	goto L40
L39:
	;
	v201 = v186 << (uint(v197) % 32)
	goto L40
L40:
	;
	if v187 < v201 {
		v210 = v193
		v211 = v187
		goto L37
	} else {
		goto L41
	}
L41:
	;
	if v193 != 0 {
		v230 = v189
		goto L26
	} else {
		goto L42
	}
L42:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v192)+20))
	if v203 == int32(-1) {
		v230 = v189
		goto L26
	} else {
		goto L43
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v125)+4)) = int64(4294967296)
	v210 = int32(1)
	v211 = int32(0)
	goto L37
L44:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v226
	v230 = v222
	goto L26
L45:
	;
	v628 = v23
	v637 = int32(0)
	goto L15
L46:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v240)+8))
	goto L51
L47:
	;
	v628 = v507
	v637 = v403
	goto L15
L48:
	;
	if v396 != l0 {
		v507 = v239
		goto L69
	} else {
		goto L70
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v251)+2224)) = int64(0)
	v396 = v380
	v403 = v248
	goto L48
L50:
	;
	if v252&int32(266) != int32(256) {
		v380 = v351
		goto L49
	} else {
		goto L62
	}
L51:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+88))
	if v252&int32(1) == int32(0) {
		v351 = int32(0)
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v251)+2164))
	if v258 < int32(1) {
		v351 = int32(0)
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v261 = int32(1)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v251)+2168))
	v264 = int32(0)
	if v258 == v261 {
		v319 = v264
		v320 = v264
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if v258&v261 == int32(0) {
		v343 = v320
		goto L59
	} else {
		goto L60
	}
L55:
	;
	v270 = int32(0)
	v277 = v270
	v278 = v270
	v282 = v270
	goto L56
L56:
	;
	v287 = int32(2)
	v289 = v263 + v277<<(uint(v287)%32)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v289+int32(4))))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+88))
	v294 = int32(-1)
	v296 = int32(3)
	v298 = int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+88))
	v309 = int32(base.Ui32(v293^v294)>>(uint(v296)%32))&v298 + (int32(base.Ui32(v301^v294)>>(uint(v296)%32))&v298 + v278)
	v311 = v277 + v287
	v313 = v282 + v287
	if v313 != v258&int32(2147483646) {
		v277 = v311
		v278 = v309
		v282 = v313
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v319 = v311
	v320 = v309
	goto L54
L58:
	;
	goto L57
L59:
	;
	if int32(0) < v343 {
		v380 = v343
		goto L49
	} else {
		goto L61
	}
L60:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v263+v319<<(uint(int32(2))%32))))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+88))
	v343 = int32(base.Ui32(v335^int32(-1))>>(uint(int32(3))%32))&int32(1) + v320
	goto L59
L61:
	;
	v351 = v343
	goto L50
L62:
	;
	if v248 != 0 {
		v369 = v248
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v251)+2224))
	if v370 != int64(0) {
		v396 = v351
		v403 = v369
		goto L48
	} else {
		goto L68
	}
L64:
	;
	v364 = int32(0)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v251)+2160))
	if v364 < v365 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v368 = v251
	goto L67
L66:
	;
	v368 = v364
	goto L67
L67:
	;
	v369 = v368
	goto L63
L68:
	;
	v373 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v251)+2224)) = v373
	v396 = v351
	v403 = v369
	goto L48
L69:
	;
	v525 = v125 + int32(20)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	if v526 != 0 {
		goto L96
	} else {
		goto L97
	}
L70:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v251)+2164))
	if v406 < int32(1) {
		v507 = v239
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v251)+2168))
	v414 = v239
	v415 = int32(0)
	goto L72
L72:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v409+v415<<(uint(int32(2))%32))))
	v429 = int32(8)
	v430 = v428 + v429
	v432 = v414 + v429
	v433 = int32(40)
	goto L78
L73:
	;
	v507 = v500
	goto L69
L74:
	;
	if v497 < int32(0) {
		goto L90
	} else {
		goto L91
	}
L75:
	;
	v497 = int32(0)
	goto L74
L76:
	;
	v469 = v464
	v470 = v465
	v471 = v466
	goto L86
L77:
	;
	if v454 == int32(0) {
		goto L75
	} else {
		goto L84
	}
L78:
	;
	if (v432|v430)&int32(3) != 0 {
		v464 = v430
		v465 = v432
		v466 = v433
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v441 = v430
	v442 = v432
	v443 = v433
	goto L80
L80:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	if v446 != v447 {
		v464 = v441
		v465 = v442
		v466 = v443
		goto L76
	} else {
		goto L82
	}
L81:
	;
	goto L77
L82:
	;
	v449 = int32(4)
	v450 = v442 + v449
	v452 = v441 + v449
	v454 = v443 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v454) {
		v441 = v452
		v442 = v450
		v443 = v454
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v464 = v452
	v465 = v450
	v466 = v454
	goto L76
L85:
	;
	v497 = v474 - v475
	goto L74
L86:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469))))
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470))))
	if v474 != v475 {
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v477 = int32(1)
	v482 = v471 + int32(-1)
	if v482 == int32(0) {
		goto L75
	} else {
		goto L89
	}
L89:
	;
	v469 = v469 + v477
	v470 = v470 + v477
	v471 = v482
	goto L86
L90:
	;
	v500 = v428
	goto L92
L91:
	;
	v500 = v414
	goto L92
L92:
	;
	v502 = v415 + int32(1)
	if v502 != v406 {
		v414 = v500
		v415 = v502
		goto L72
	} else {
		goto L93
	}
L93:
	;
	goto L73
L94:
	;
	if v621 != 0 {
		v239 = v507
		v240 = v621
		v248 = v403
		goto L46
	} else {
		goto L120
	}
L95:
	;
	v532 = v525
	v533 = v529
	goto L98
L96:
	;
	v529 = int32(1)
	goto L95
L97:
	;
	v529 = int32(0)
	goto L95
L98:
	;
	switch v533 {
	case 0:
		goto L103
	default:
		goto L102
	}
L100:
	;
	v533 = int32(0)
	goto L98
L101:
	;
	goto L94
L102:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+16)) = v613
	if v613 == int32(0) {
		goto L100
	} else {
		goto L119
	}
L103:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v537 != int32(-1) {
		v576 = v537
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v577 = int32(1)
	v578 = v576 + v577
	*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = v578
	v580 = int32(0)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583+v584+int32(26)))))
	if v588 == int32(255) {
		goto L113
	} else {
		goto L114
	}
L105:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	if v541 != 0 {
		v576 = int32(-1)
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	if v543 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)+20))
	if v570 != int32(-1) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v550 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v542)+16)))
	v551 = int64(*(*int8)(unsafe.Add(mBase, uint32(v542)+27)))
	v552 = int64(*(*int32)(unsafe.Add(mBase, uint32(v542)+8)))
	v553 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v542)+12)))
	v554 = int64(*(*int8)(unsafe.Add(mBase, uint32(v542)+26)))
	v555 = int64(*(*int32)(unsafe.Add(mBase, uint32(v542)+4)))
	v556 = F_wangHash64(m, v555)
	mBase = m.M
	v558 = F_wangHash64(m, v554+v556)
	mBase = m.M
	v560 = F_wangHash64(m, v553+v558)
	mBase = m.M
	v562 = F_wangHash64(m, v552+v560)
	mBase = m.M
	v564 = F_wangHash64(m, v551+v562)
	mBase = m.M
	v566 = F_wangHash64(m, v550+v564)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v125)+24)) = v566
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v569 = v568
	goto L107
L109:
	;
	v546 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v542)+24)))
	v548 = v546 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v542)+24)) = uint16(v548)
	v569 = v542
	goto L107
L110:
	;
	v576 = v570 + int32(-1)
	goto L104
L111:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v576 = v573
	goto L104
L112:
	;
	v603 = int32(2)
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v583+v601<<(uint(v603)%32)+int32(4))))
	v532 = v608 + v602<<(uint(v603)%32)
	v533 = int32(1)
	goto L98
L113:
	;
	v592 = v580
	goto L115
L114:
	;
	v592 = v577 << (uint(v588) % 32)
	goto L115
L115:
	;
	if v578 < v592 {
		v601 = v584
		v602 = v578
		goto L112
	} else {
		goto L116
	}
L116:
	;
	if v584 != 0 {
		v621 = v580
		goto L101
	} else {
		goto L117
	}
L117:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v583)+20))
	if v594 == int32(-1) {
		v621 = v580
		goto L101
	} else {
		goto L118
	}
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v125)+4)) = int64(4294967296)
	v601 = int32(1)
	v602 = int32(0)
	goto L112
L119:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v613)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v525))) = v617
	v621 = v613
	goto L101
L120:
	;
	goto L47
L121:
	;
	if v637 == int32(0) {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v644 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaMigration[1]))
	if v628 != v644 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v646 = F_mstime(m)
	mBase = m.M
	v647 = *(*int64)(unsafe.Add(mBase, uint32(v637)+2224))
	if v646-v647 < int64(5001) {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v652 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_clusterHandleReplicaMigration[3])))
	if v652&int32(2) != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v656 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaMigration[4]))
	if int32(2) < v656 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v673 = int32(1)
	F_clusterSetPrimary(m, v637, v673, v673)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L17
	} else {
		goto L130
	}
L127:
	;
	v659 = F_humanNodename(m, v637)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L17
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v637 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v659
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v637 + int32(8)
	F__serverLog(m, int32(2), int32(_a_F_clusterHandleReplicaMigration_0), v17)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L17
	} else {
		goto L129
	}
L129:
	;
	goto L126
L130:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L17
	} else {
		goto L131
	}
L131:
	;
	v680 = *(*int32)(unsafe.Add(mBase, _c_F_clusterHandleReplicaMigration[0]))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v680)+uint32(_c_F_clusterHandleReplicaMigration[5])))
	*(*int32)(unsafe.Add(mBase, uint32(v680)+uint32(_c_F_clusterHandleReplicaMigration[5]))) = v681 | int32(44)
	goto L1
}
func F_clusterIsSlotImporting(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
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
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_clusterIsSlotImporting[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_clusterIsSlotImporting[1])))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
	goto L1
L1:
	;
	v16 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v18 == v16 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return v109
L3:
	;
	if v18 == int32(0) {
		v109 = v16
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18+base.B2i32(v21 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v27
	goto L4
L6:
	;
	v34 = v18
	goto L8
L7:
	;
	v109 = int32(1)
	goto L2
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v95 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+156))
	if base.Ui32(v39+int32(-18)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)+164))
	v46 = v7 + int32(8)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v47
	goto L13
L13:
	;
	v52 = v7 + int32(8)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v54 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v54 == int32(0) {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54+base.B2i32(v57 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v63
	goto L15
L17:
	;
	v70 = v54
	goto L18
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if l0 < v72 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L10
L20:
	;
	v77 = v7 + int32(8)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v79 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if l0 <= v74 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v79 != 0 {
		v70 = v79
		goto L18
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79+base.B2i32(v82 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v88
	goto L24
L26:
	;
	goto L19
L27:
	;
	if v95 != 0 {
		v34 = v95
		goto L8
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v95+base.B2i32(v98 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v104
	goto L28
L30:
	;
	v109 = v16
	goto L2
}
func F_clusterIsValidPacket(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
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
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v417 int32
	_ = v417
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v459 int32
	_ = v459
	var v470 int32
	_ = v470
	v10 = m.G0
	v12 = v10 - int32(192)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = F___bswap_32_2(m, v15)
	mBase = m.M
	goto L1
L1:
	;
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+12)))
	v18 = F___bswap_16_2(m, v17)
	mBase = m.M
	goto L2
L2:
	;
	v19 = base.I32_extend16_s(v18)
	v21 = v18 & int32(32767)
	if base.Ui32(int32(10)) < base.Ui32(v21) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if int32(-1) < v19 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_clusterIsValidPacket[0]))
	v28 = v25 + v21<<(uint(int32(3))%32)
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_clusterIsValidPacket[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_clusterIsValidPacket[1]))) = v29 + int64(1)
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_clusterIsValidPacket[2])))
	v34 = base.I64_extend_i32_u(v16)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_clusterIsValidPacket[2]))) = v33 + v34
	switch v21 + int32(-4) {
	case 0, 6:
		v41 = int32(65880)
		goto L5
	default:
		goto L3
	case 5:
		goto L6
	}
L5:
	;
	v42 = v25 + v41
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = v43 + v34
	goto L3
L6:
	;
	v41 = int32(65896)
	goto L5
L7:
	;
	F__serverAssert(m, int32(_a_F_clusterIsValidPacket_0), int32(_a_F_clusterIsValidPacket_1), int32(145))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L17
	} else {
		goto L125
	}
L8:
	;
	m.G0 = v12 + int32(192)
	return v459
L9:
	;
	v81 = int32(0)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_clusterIsValidPacket[3]))
	if v81 < v83 {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	if base.Ui32(int32(10)) < base.Ui32(v21) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v57 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_clusterIsValidPacket[3]))
	if int32(2) < v59 {
		v459 = v57
		goto L8
	} else {
		goto L14
	}
L12:
	;
	if int32(1)<<(uint(v21)%32)&int32(1552) != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	if base.Ui32(int32(10)) < base.Ui32(v21) {
		v70 = int32(_a_F_clusterIsValidPacket_2)
		goto L15
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+180)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v12)+176)) = v70
	F__serverLog(m, int32(2), int32(_a_F_clusterIsValidPacket_3), v12+int32(176))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_c_F_clusterIsValidPacket[4])))
	v70 = v69
	goto L15
L17:
	;
	return int32(0)
L18:
	;
	v459 = v57
	goto L8
L19:
	;
	if base.Ui32(v16) < base.Ui32(int32(16)) {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	if base.Ui32(int32(10)) < base.Ui32(v21) {
		v94 = int32(_a_F_clusterIsValidPacket_2)
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v94
	F__serverLog(m, int32(0), int32(_a_F_clusterIsValidPacket_4), v12+int32(160))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L17
	} else {
		goto L23
	}
L22:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_c_F_clusterIsValidPacket[4])))
	v94 = v93
	goto L21
L23:
	;
	goto L19
L24:
	;
	v459 = int32(0)
	goto L8
L25:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if base.Ui32(v106) < base.Ui32(v16) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+8)))
	v109 = F___bswap_16_2(m, v108)
	mBase = m.M
	goto L27
L27:
	;
	if v109 != int32(1) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_clusterIsValidPacket[5]))
	if v113 == v21 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if base.Ui32(int32(2)) < base.Ui32(v21) {
		goto L38
	} else {
		goto L39
	}
L30:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_clusterIsValidPacket[3]))
	if int32(3) < v118 {
		v459 = v81
		goto L8
	} else {
		goto L33
	}
L31:
	;
	if v113 != int32(-2) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	if base.Ui32(int32(10)) < base.Ui32(v21) {
		v129 = int32(_a_F_clusterIsValidPacket_2)
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v129
	F__serverLog(m, int32(3), int32(_a_F_clusterIsValidPacket_5), v12)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L17
	} else {
		goto L36
	}
L35:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_c_F_clusterIsValidPacket[4])))
	v129 = v128
	goto L34
L36:
	;
	goto L24
L37:
	;
	if v16 != v407 {
		goto L118
	} else {
		goto L119
	}
L38:
	;
	if v21 != int32(3) {
		goto L75
	} else {
		goto L76
	}
L39:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+12)))
	v139 = F___bswap_16_2(m, v138)
	mBase = m.M
	goto L40
L40:
	;
	if base.I32_extend16_s(v139) <= int32(-1) {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+2214)))
	v144 = F___bswap_16_2(m, v143)
	mBase = m.M
	goto L42
L42:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+14)))
	v146 = F___bswap_16_2(m, v145)
	mBase = m.M
	goto L44
L43:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+2253)))
	if v176&int32(4) == int32(0) {
		v407 = v150
		goto L37
	} else {
		goto L51
	}
L44:
	;
	v150 = v146*int32(104) + int32(2256)
	if base.Ui32(v150) <= base.Ui32(v16) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_clusterIsValidPacket[3]))
	if int32(3) < v153 {
		v459 = v81
		goto L8
	} else {
		goto L46
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = base.I64_extend_i32_u(v16)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v146
	if base.Ui32(int32(10)) < base.Ui32(v21) {
		v168 = int32(_a_F_clusterIsValidPacket_2)
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v168
	F__serverLog(m, int32(3), int32(_a_F_clusterIsValidPacket_6), v12+int32(16))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L17
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_c_F_clusterIsValidPacket[4])))
	v168 = v167
	goto L48
L50:
	;
	goto L24
L51:
	;
	if v144 == int32(0) {
		v407 = v150
		goto L37
	} else {
		goto L52
	}
L52:
	;
	v188 = v137 + v146*int32(104) + int32(2256)
	v190 = v150
	v193 = v144
	goto L53
L53:
	;
	v197 = v16 - v190
	if base.Ui32(int32(7)) < base.Ui32(v197) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v224 = F___bswap_32_2(m, v223)
	mBase = m.M
	goto L63
L56:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_clusterIsValidPacket[3]))
	if int32(3) < v201 {
		goto L24
	} else {
		goto L57
	}
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = base.I64_extend_i32_u(v16)
	if base.Ui32(int32(10)) < base.Ui32(v21) {
		v215 = int32(_a_F_clusterIsValidPacket_2)
		goto L59
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v215
	F__serverLog(m, int32(3), int32(_a_F_clusterIsValidPacket_7), v12+int32(64))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L17
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_c_F_clusterIsValidPacket[4])))
	v215 = v214
	goto L59
L61:
	;
	goto L24
L62:
	;
	if base.Ui32(v224) <= base.Ui32(v197) {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	if v224&int32(7) == int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_clusterIsValidPacket[3]))
	if int32(3) < v230 {
		goto L24
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+100)) = v224
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_c_F_clusterIsValidPacket[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v238
	F__serverLog(m, int32(3), int32(_a_F_clusterIsValidPacket_8), v12+int32(96))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L17
	} else {
		goto L66
	}
L66:
	;
	goto L24
L67:
	;
	v271 = v224 + v190
	v273 = v193 + int32(-1)
	if v273&int32(65535) != 0 {
		v188 = v188 + v224
		v190 = v271
		v193 = v273
		goto L53
	} else {
		goto L74
	}
L68:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_clusterIsValidPacket[3]))
	if int32(3) < v248 {
		goto L24
	} else {
		goto L69
	}
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = base.I64_extend_i32_u(v16)
	if base.Ui32(int32(10)) < base.Ui32(v21) {
		v262 = int32(_a_F_clusterIsValidPacket_2)
		goto L71
	} else {
		goto L72
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v262
	F__serverLog(m, int32(3), int32(_a_F_clusterIsValidPacket_7), v12+int32(80))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L17
	} else {
		goto L73
	}
L71:
	;
	goto L70
L72:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_c_F_clusterIsValidPacket[4])))
	v262 = v261
	goto L71
L73:
	;
	goto L24
L74:
	;
	v407 = v271
	goto L37
L75:
	;
	v280 = v19 & int32(65535)
	switch v280 + int32(-4) {
	case 0, 6:
		goto L78
	case 1, 2, 3, 4, 5:
		goto L77
	default:
		goto L79
	}
L76:
	;
	v407 = int32(2296)
	goto L37
L77:
	;
	v358 = int32(2256)
	if v21 == int32(8) {
		v407 = v358
		goto L37
	} else {
		goto L103
	}
L78:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if int32(-1) < v19 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	switch v280 + int32(-32772) {
	case 0, 6:
		goto L78
	default:
		goto L77
	}
L80:
	;
	if base.Ui32(v16) < base.Ui32(v297) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v292 = F_toClusterMsg(m, v285)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L17
	} else {
		goto L84
	}
L82:
	;
	v288 = F_toClusterMsgLight(m, v285)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L17
	} else {
		goto L83
	}
L83:
	;
	v296 = int32(16)
	v297 = int32(24)
	goto L80
L84:
	;
	v296 = int32(2256)
	v297 = int32(2264)
	goto L80
L85:
	;
	v329 = v302 + v297
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	v332 = F___bswap_32_2(m, v331)
	mBase = m.M
	goto L96
L86:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_clusterIsValidPacket[3]))
	if int32(3) < v307 {
		goto L24
	} else {
		goto L90
	}
L87:
	;
	v300 = v285 + v296
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	v302 = F___bswap_32_2(m, v301)
	mBase = m.M
	goto L88
L88:
	;
	if base.Ui32(v302) <= base.Ui32(v16-v297) {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	goto L86
L90:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+136)) = base.I64_extend_i32_u(v16)
	if base.Ui32(int32(10)) < base.Ui32(v21) {
		v321 = int32(_a_F_clusterIsValidPacket_2)
		goto L92
	} else {
		goto L93
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v321
	F__serverLog(m, int32(3), int32(_a_F_clusterIsValidPacket_9), v12+int32(128))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L17
	} else {
		goto L94
	}
L92:
	;
	goto L91
L93:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_c_F_clusterIsValidPacket[4])))
	v321 = v320
	goto L92
L94:
	;
	goto L24
L95:
	;
	v407 = v332 + v329
	goto L37
L96:
	;
	if base.Ui32(v332) <= base.Ui32(v16-v329) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_clusterIsValidPacket[3]))
	if int32(3) < v335 {
		goto L24
	} else {
		goto L98
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+152)) = base.I64_extend_i32_u(v16)
	if base.Ui32(int32(10)) < base.Ui32(v21) {
		v349 = int32(_a_F_clusterIsValidPacket_2)
		goto L100
	} else {
		goto L101
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v349
	F__serverLog(m, int32(3), int32(_a_F_clusterIsValidPacket_10), v12+int32(144))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L17
	} else {
		goto L102
	}
L100:
	;
	goto L99
L101:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_c_F_clusterIsValidPacket[4])))
	v349 = v348
	goto L100
L102:
	;
	goto L24
L103:
	;
	if base.Ui32((v21+int32(-5))&int32(65535)) < base.Ui32(int32(2)) {
		v407 = v358
		goto L37
	} else {
		goto L104
	}
L104:
	;
	switch v21 + int32(-7) {
	case 0:
		v407 = int32(4352)
		goto L37
	default:
		v459 = int32(1)
		goto L8
	case 2:
		goto L105
	}
L105:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if int32(-1) < v19 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v371+v382)))
	v386 = F___bswap_32_2(m, v385)
	mBase = m.M
	goto L111
L107:
	;
	v378 = F_toClusterMsg(m, v371)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L17
	} else {
		goto L110
	}
L108:
	;
	v374 = F_toClusterMsgLight(m, v371)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L17
	} else {
		goto L109
	}
L109:
	;
	v382 = int32(24)
	v383 = int32(29)
	goto L106
L110:
	;
	v382 = int32(2264)
	v383 = int32(2269)
	goto L106
L111:
	;
	if base.Ui32(v16) < base.Ui32(v383) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v407 = v386 + v383
	goto L37
L113:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_clusterIsValidPacket[3]))
	if int32(3) < v391 {
		goto L24
	} else {
		goto L116
	}
L114:
	;
	if base.Ui32(v386) <= base.Ui32(v16-v383) {
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = int32(_a_F_clusterIsValidPacket_11)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+120)) = base.I64_extend_i32_u(v16)
	F__serverLog(m, int32(3), int32(_a_F_clusterIsValidPacket_12), v12+int32(112))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L17
	} else {
		goto L117
	}
L117:
	;
	goto L24
L118:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_clusterIsValidPacket[3]))
	if int32(3) < v417 {
		goto L24
	} else {
		goto L120
	}
L119:
	;
	v459 = int32(1)
	goto L8
L120:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(48)))) = base.I64_extend_i32_u(v407)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = base.I64_extend_i32_u(v16)
	if base.Ui32(int32(10)) < base.Ui32(v21) {
		v435 = int32(_a_F_clusterIsValidPacket_2)
		goto L122
	} else {
		goto L123
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v435
	F__serverLog(m, int32(3), int32(_a_F_clusterIsValidPacket_13), v12+int32(32))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L17
	} else {
		goto L124
	}
L122:
	;
	goto L121
L123:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_c_F_clusterIsValidPacket[4])))
	v435 = v434
	goto L122
L124:
	;
	goto L24
L125:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterLinkConnectHandler(m *base.Module, l0 int32) {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == int32(3) {
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+84))
		v45 = m.T0[v44].(func(*base.Module, int32, int32) int32)(m, l0, int32(64))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return
		} else {
			v47 = *(*int64)(unsafe.Add(mBase, uint32(v12)+2184))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
			F_clusterSendPing(m, v11, int32(base.Ui32(v48)>>(uint(int32(6))%32))&int32(2))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				if v47 == int64(0) {
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v12)+2184)) = v47
				}
				v59 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLinkConnectHandler[0]))
				if int32(0) < v59 {
					m.G0 = v9 + int32(32)
					return
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2332))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v62
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v12 + int32(2256)
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12 + int32(8)
					F__serverLog(m, int32(0), int32(_a_F_clusterLinkConnectHandler_0), v9)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLinkConnectHandler[0]))
		if int32(1) < v17 {
			F_freeClusterLink(m, v11)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				m.G0 = v9 + int32(32)
				return
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2332))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+88))
			v23 = m.T0[v22].(func(*base.Module, int32) int32)(m, l0)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v12 + int32(2256)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v12 + int32(8)
				F__serverLog(m, int32(1), int32(_a_F_clusterLinkConnectHandler_1), v9+int32(16))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_freeClusterLink(m, v11)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
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
func F_clusterLockConfig(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(64)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = int32(420)
	v12 = int32(-1)
	v16 = F_open(m, l0, int32(524353), v8+int32(48))
	mBase = m.M
	if v16 != v12 {
		*(*int32)(unsafe.Add(mBase, _c_F_clusterLockConfig[0])) = v16
		v69 = int32(0)
		m.G0 = v8 + int32(64)
		return v69
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLockConfig[1]))
		if int32(3) < v20 {
			v69 = v12
			m.G0 = v8 + int32(64)
			return v69
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLockConfig[2]))
			v25 = F___strerror_l(m, v24, v24)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v25
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
			F__serverLog(m, int32(3), int32(_a_F_clusterLockConfig_0), v8)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v69 = v12
				m.G0 = v8 + int32(64)
				return v69
			}
		}
	}
}
func F_clusterLogCantFailover(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F___time(m, v2)
	mBase = m.M
	v14 = *(*int64)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[0]))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[1]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_clusterLogCantFailover[2])))
	if l0 != v17 {
		if l0 != int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_clusterLogCantFailover[2]))) = l0
			v32 = l0 + int32(-1)
			if base.Ui32(int32(4)) <= base.Ui32(v32) {
				F__serverPanic_1(m, int32(_a_F_clusterLogCantFailover_0), int32(5563), int32(_a_F_clusterLogCantFailover_1), int32(0))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_c_F_clusterLogCantFailover[3])))
				v40 = v39
				v41 = int32(0)
				v43 = F___time(m, v41)
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[0])) = v43
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[4]))
				if int32(2) < v46 {
					m.G0 = v9 + int32(32)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v40
					F__serverLog(m, int32(2), int32(_a_F_clusterLogCantFailover_2), v9+int32(16))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						if base.Ui32(l0) < base.Ui32(int32(3)) {
							m.G0 = v9 + int32(32)
							return
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[4]))
							if int32(2) < v59 {
								m.G0 = v9 + int32(32)
								return
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[1]))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+28))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_clusterLogCantFailover[5])))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v65
								v67 = int32(2)
								v68 = base.I32_div_s(v64, v67)
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v68 + int32(1)
								F__serverLog(m, v67, int32(_a_F_clusterLogCantFailover_3), v9)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
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
		} else {
			if v17 != int32(1) {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_clusterLogCantFailover[2]))) = l0
				v32 = l0 + int32(-1)
				if base.Ui32(int32(4)) <= base.Ui32(v32) {
					F__serverPanic_1(m, int32(_a_F_clusterLogCantFailover_0), int32(5563), int32(_a_F_clusterLogCantFailover_1), int32(0))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_c_F_clusterLogCantFailover[3])))
					v40 = v39
					v41 = int32(0)
					v43 = F___time(m, v41)
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[0])) = v43
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[4]))
					if int32(2) < v46 {
						m.G0 = v9 + int32(32)
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v40
						F__serverLog(m, int32(2), int32(_a_F_clusterLogCantFailover_2), v9+int32(16))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							if base.Ui32(l0) < base.Ui32(int32(3)) {
								m.G0 = v9 + int32(32)
								return
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[4]))
								if int32(2) < v59 {
									m.G0 = v9 + int32(32)
									return
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[1]))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+28))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_clusterLogCantFailover[5])))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v65
									v67 = int32(2)
									v68 = base.I32_div_s(v64, v67)
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v68 + int32(1)
									F__serverLog(m, v67, int32(_a_F_clusterLogCantFailover_3), v9)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
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
			} else {
				if v12-v14 < int64(10) {
					m.G0 = v9 + int32(32)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_clusterLogCantFailover[2]))) = int32(1)
					v40 = int32(_a_F_clusterLogCantFailover_4)
					v41 = int32(0)
					v43 = F___time(m, v41)
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[0])) = v43
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[4]))
					if int32(2) < v46 {
						m.G0 = v9 + int32(32)
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v40
						F__serverLog(m, int32(2), int32(_a_F_clusterLogCantFailover_2), v9+int32(16))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							if base.Ui32(l0) < base.Ui32(int32(3)) {
								m.G0 = v9 + int32(32)
								return
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[4]))
								if int32(2) < v59 {
									m.G0 = v9 + int32(32)
									return
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[1]))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+28))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_clusterLogCantFailover[5])))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v65
									v67 = int32(2)
									v68 = base.I32_div_s(v64, v67)
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v68 + int32(1)
									F__serverLog(m, v67, int32(_a_F_clusterLogCantFailover_3), v9)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
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
	} else {
		if v12 <= v14 {
			m.G0 = v9 + int32(32)
			return
		} else {
			if l0 != int32(1) {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_clusterLogCantFailover[2]))) = l0
				v32 = l0 + int32(-1)
				if base.Ui32(int32(4)) <= base.Ui32(v32) {
					F__serverPanic_1(m, int32(_a_F_clusterLogCantFailover_0), int32(5563), int32(_a_F_clusterLogCantFailover_1), int32(0))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_c_F_clusterLogCantFailover[3])))
					v40 = v39
					v41 = int32(0)
					v43 = F___time(m, v41)
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[0])) = v43
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[4]))
					if int32(2) < v46 {
						m.G0 = v9 + int32(32)
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v40
						F__serverLog(m, int32(2), int32(_a_F_clusterLogCantFailover_2), v9+int32(16))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							if base.Ui32(l0) < base.Ui32(int32(3)) {
								m.G0 = v9 + int32(32)
								return
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[4]))
								if int32(2) < v59 {
									m.G0 = v9 + int32(32)
									return
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[1]))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+28))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_clusterLogCantFailover[5])))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v65
									v67 = int32(2)
									v68 = base.I32_div_s(v64, v67)
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v68 + int32(1)
									F__serverLog(m, v67, int32(_a_F_clusterLogCantFailover_3), v9)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
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
			} else {
				if v17 != int32(1) {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_clusterLogCantFailover[2]))) = l0
					v32 = l0 + int32(-1)
					if base.Ui32(int32(4)) <= base.Ui32(v32) {
						F__serverPanic_1(m, int32(_a_F_clusterLogCantFailover_0), int32(5563), int32(_a_F_clusterLogCantFailover_1), int32(0))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_c_F_clusterLogCantFailover[3])))
						v40 = v39
						v41 = int32(0)
						v43 = F___time(m, v41)
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[0])) = v43
						v46 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[4]))
						if int32(2) < v46 {
							m.G0 = v9 + int32(32)
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v40
							F__serverLog(m, int32(2), int32(_a_F_clusterLogCantFailover_2), v9+int32(16))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								if base.Ui32(l0) < base.Ui32(int32(3)) {
									m.G0 = v9 + int32(32)
									return
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[4]))
									if int32(2) < v59 {
										m.G0 = v9 + int32(32)
										return
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[1]))
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+28))
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_clusterLogCantFailover[5])))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v65
										v67 = int32(2)
										v68 = base.I32_div_s(v64, v67)
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v68 + int32(1)
										F__serverLog(m, v67, int32(_a_F_clusterLogCantFailover_3), v9)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
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
				} else {
					if v12-v14 < int64(10) {
						m.G0 = v9 + int32(32)
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_clusterLogCantFailover[2]))) = int32(1)
						v40 = int32(_a_F_clusterLogCantFailover_4)
						v41 = int32(0)
						v43 = F___time(m, v41)
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[0])) = v43
						v46 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[4]))
						if int32(2) < v46 {
							m.G0 = v9 + int32(32)
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v40
							F__serverLog(m, int32(2), int32(_a_F_clusterLogCantFailover_2), v9+int32(16))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								if base.Ui32(l0) < base.Ui32(int32(3)) {
									m.G0 = v9 + int32(32)
									return
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[4]))
									if int32(2) < v59 {
										m.G0 = v9 + int32(32)
										return
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, _c_F_clusterLogCantFailover[1]))
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+28))
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_clusterLogCantFailover[5])))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v65
										v67 = int32(2)
										v68 = base.I32_div_s(v64, v67)
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v68 + int32(1)
										F__serverLog(m, v67, int32(_a_F_clusterLogCantFailover_3), v9)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
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
}
func F_clusterMoveNodeSlots(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int64
	_ = v206
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
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
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v346 int32
	_ = v346
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	v22 = int32(48)
	v23 = l0 + v22
	v25 = l1 + v22
	v26 = int32(40)
	goto L6
L1:
	;
	F__serverAssert(m, int32(_a_F_clusterMoveNodeSlots_0), int32(_a_F_clusterMoveNodeSlots_1), int32(6564))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L23
	} else {
		goto L89
	}
L2:
	;
	if v90 != 0 {
		goto L1
	} else {
		goto L18
	}
L3:
	;
	v90 = int32(0)
	goto L2
L4:
	;
	v62 = v57
	v63 = v58
	v64 = v59
	goto L14
L5:
	;
	if v47 == int32(0) {
		goto L3
	} else {
		goto L12
	}
L6:
	;
	if (v25|v23)&int32(3) != 0 {
		v57 = v23
		v58 = v25
		v59 = v26
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v34 = v23
	v35 = v25
	v36 = v26
	goto L8
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v39 != v40 {
		v57 = v34
		v58 = v35
		v59 = v36
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L5
L10:
	;
	v42 = int32(4)
	v43 = v35 + v42
	v45 = v34 + v42
	v47 = v36 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v47) {
		v34 = v45
		v35 = v43
		v36 = v47
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v57 = v45
	v58 = v43
	v59 = v47
	goto L4
L13:
	;
	v90 = v67 - v68
	goto L2
L14:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v67 != v68 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v70 = int32(1)
	v75 = v64 + int32(-1)
	if v75 == int32(0) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v62 = v62 + v70
	v63 = v63 + v70
	v64 = v75
	goto L14
L18:
	;
	v92 = l1 + int32(8)
	v95 = int32(0)
	v108 = v95
	v109 = v95
	v110 = v95
	v111 = v95
	goto L19
L19:
	;
	v119 = int32(1) << (uint(v108&int32(7)) % 32)
	v121 = int32(base.Ui32(v108) >> (uint(int32(3)) % 32))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(104)+v121))))
	if v119&v123 == int32(0) {
		v220 = v111
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if l2 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L21:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_clusterMoveNodeSlots[0]))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+48))
	v227 = F_dictFind(m, v226, v108)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L23
	} else {
		goto L42
	}
L22:
	;
	v127 = F_clusterDelSlot(m, v108)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return
L24:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_clusterMoveNodeSlots[0]))
	v132 = v108 << (uint(int32(2)) % 32)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v130+v132)+52))
	if v134 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v220 = v111 + int32(1)
	goto L21
L26:
	;
	v138 = m.G0
	v140 = v138 - int32(32)
	m.G0 = v140
	v145 = int32(1) << (uint(v108&int32(7)) % 32)
	v147 = base.I32_div_s(v108, int32(8))
	v150 = l1 + v147 + int32(104)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	if v145&v151 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v189 = int32(_a_F_clusterMoveNodeSlots_2)
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_clusterMoveNodeSlots[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v190+v132)+52)) = l1
	v193 = v190 + v121
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+uint32(_c_F_clusterMoveNodeSlots[1]))))
	v197 = v194 & (v119 ^ int32(-1))
	*(*uint8)(unsafe.Add(mBase, uint32(v193)+uint32(_c_F_clusterMoveNodeSlots[1]))) = uint8(v197)
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_clusterMoveNodeSlots[0]))
	v203 = v200 + v108*int32(24)
	v206 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_clusterMoveNodeSlots[2]))) = v206
	*(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_clusterMoveNodeSlots[3]))) = v206
	*(*int64)(unsafe.Add(mBase, uint32(v203)+uint32(_c_F_clusterMoveNodeSlots[4]))) = v206
	goto L39
L28:
	;
	m.G0 = v140 + int32(32)
	goto L27
L29:
	;
	v153 = v151 | v145
	*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v153)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2160))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+2160)) = v155 + int32(1)
	if v155 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_clusterMoveNodeSlots[0]))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+32))
	F_dictInitIterator(m, v140, v161)
	mBase = m.M
	v163 = F_dictNext(m, v140)
	mBase = m.M
	if v163 == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v167 = v163
	goto L33
L32:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v177 | int32(256)
	goto L28
L33:
	;
	v171 = F_dictGetVal(m, v167)
	mBase = m.M
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+88)))
	if v172&int32(2) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v176 = F_dictNext(m, v140)
	mBase = m.M
	if v176 != 0 {
		v167 = v176
		goto L33
	} else {
		goto L38
	}
L36:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v171)+2164))
	if v175 != 0 {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	goto L28
L39:
	;
	goto L25
L40:
	;
	if v231 != l0 {
		v273 = v110
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	goto L44
L42:
	;
	if v227 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v231 = int32(0)
	goto L40
L44:
	;
	v231 = v230
	goto L40
L45:
	;
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_clusterMoveNodeSlots[0]))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+44))
	v279 = F_dictFind(m, v278, v108)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L23
	} else {
		goto L63
	}
L46:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_clusterMoveNodeSlots[5]))
	if int32(1) < v234 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_clusterMoveNodeSlots[0]))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+48))
	v253 = F_dictFind(m, v252, v108)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L23
	} else {
		goto L51
	}
L48:
	;
	v237 = F_humanNodename(m, l1)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L23
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v108
	F__serverLog(m, int32(1), int32(_a_F_clusterMoveNodeSlots_3), v20+int32(16))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L23
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	if l1 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v273 = v110 + int32(1)
	goto L45
L53:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_clusterMoveNodeSlots[0]))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+48))
	if v253 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	if v253 == int32(0) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_clusterMoveNodeSlots[0]))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+48))
	v260 = F_dictDelete(m, v259, v108)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L23
	} else {
		goto L56
	}
L56:
	;
	goto L52
L57:
	;
	v268 = F_dictAdd(m, v264, v108, l1)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L23
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+8)) = l1
	goto L59
L59:
	;
	goto L52
L60:
	;
	goto L52
L61:
	;
	if v283 != l0 {
		v323 = v109
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v279)+8))
	goto L65
L63:
	;
	if v279 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v283 = int32(0)
	goto L61
L65:
	;
	v283 = v282
	goto L61
L66:
	;
	v327 = v108 + int32(1)
	if v327 != int32(16384) {
		v108 = v327
		v109 = v323
		v110 = v273
		v111 = v220
		goto L19
	} else {
		goto L82
	}
L67:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_clusterMoveNodeSlots[5]))
	if int32(1) < v286 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_clusterMoveNodeSlots[0]))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+44))
	v303 = F_dictFind(m, v302, v108)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L23
	} else {
		goto L72
	}
L69:
	;
	v289 = F_humanNodename(m, l1)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L23
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v108
	F__serverLog(m, int32(1), int32(_a_F_clusterMoveNodeSlots_4), v20)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L23
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	if l1 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v323 = v109 + int32(1)
	goto L66
L74:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_clusterMoveNodeSlots[0]))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+44))
	if v303 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	if v303 == int32(0) {
		goto L73
	} else {
		goto L76
	}
L76:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_clusterMoveNodeSlots[0]))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+44))
	v310 = F_dictDelete(m, v309, v108)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L23
	} else {
		goto L77
	}
L77:
	;
	goto L73
L78:
	;
	v318 = F_dictAdd(m, v314, v108, l1)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L23
	} else {
		goto L81
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+8)) = l1
	goto L80
L80:
	;
	goto L73
L81:
	;
	goto L73
L82:
	;
	goto L20
L83:
	;
	if l3 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v220
	goto L83
L85:
	;
	if l4 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v273
	goto L85
L87:
	;
	m.G0 = v20 + int32(32)
	return
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v323
	goto L87
L89:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterNodeAddFailureReport(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v28 int64
	_ = v28
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v118 int64
	_ = v118
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	v8 = m.G0
	v10 = v8 - int32(320)
	m.G0 = v10
	v12 = F_mstime(m)
	mBase = m.M
	v14 = base.I64_rem_u_s(v12, int64(1000))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v16&int32(8) != 0 {
		v167 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(320)
	return v167
L2:
	;
	v21 = v12 - v14 + int64(1000)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2352))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(128)
	v28 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+12)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v10)+296)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v10)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v10 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+156)) = v10 + int32(168)
	goto L3
L3:
	;
	v41 = int32(0)
	v43 = F_raxSeek(m, v10, int32(_a_F_clusterNodeAddFailureReport_0), v41, v41)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L7
L6:
	;
	F_raxStop(m, v10)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L17
	}
L7:
	;
	v54 = F_raxNext(m, v10)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	v61 = int64(56)
	v63 = int64(65280)
	v65 = int64(40)
	v68 = int64(16711680)
	v70 = int64(24)
	v72 = int64(4278190080)
	v74 = int64(8)
	if v60<<(uint(v61)%64)|v60&v63<<(uint(v65)%64)|(v60&v68<<(uint(v70)%64)|v60&v72<<(uint(v74)%64))|(int64(base.Ui64(v60)>>(uint(v74)%64))&v72|int64(base.Ui64(v60)>>(uint(v70)%64))&v68|(int64(base.Ui64(v60)>>(uint(v65)%64))&v63|int64(base.Ui64(v60)>>(uint(v61)%64)))) != v21 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if l1 != v58 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	if v54 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v107 = int32(1)
	goto L6
L12:
	;
	goto L8
L13:
	;
	v100 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2352))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v104 = F_raxRemove(m, v101, v57, v102, v100)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	F_raxStop(m, v10)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v167 = int32(0)
	goto L1
L16:
	;
	v107 = v100
	goto L6
L17:
	;
	v111 = int64(56)
	v113 = int64(65280)
	v115 = int64(40)
	v118 = int64(16711680)
	v120 = int64(24)
	v122 = int64(4278190080)
	v124 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+304)) = v21<<(uint(v111)%64) | v21&v113<<(uint(v115)%64) | (v21&v118<<(uint(v120)%64) | v21&v122<<(uint(v124)%64)) | (int64(base.Ui64(v21)>>(uint(v124)%64))&v122 | int64(base.Ui64(v21)>>(uint(v120)%64))&v118 | (int64(base.Ui64(v21)>>(uint(v115)%64))&v113 | int64(base.Ui64(v21)>>(uint(v111)%64))))
	v150 = v10 + int32(304) | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = l1
	goto L18
L18:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2352))
	v158 = int32(0)
	v160 = F_raxInsert(m, v154, v10+int32(304), int32(16), v158, v158)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v167 = v107
	goto L1
}
func F_clusterNodeClientPort(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	if l1 == int32(0) {
		if l1 == int32(0) {
			if l2 == int32(0) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2324))
				v21 = v20
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2336))
				if v18 != 0 {
					v21 = v18
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2324))
					v21 = v20
				}
			}
			return v21
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2328))
			return v14
		}
	} else {
		if l2 == int32(0) {
			if l1 == int32(0) {
				if l2 == int32(0) {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2324))
					v21 = v20
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2336))
					if v18 != 0 {
						v21 = v18
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2324))
						v21 = v20
					}
				}
				return v21
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2328))
				return v14
			}
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2340))
			if v8 == int32(0) {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2328))
				return v14
			} else {
				v21 = v8
				return v21
			}
		}
	}
}
func F_clusterNodeGetName(m *base.Module, l0 int32) int32 {
	return l0 + int32(8)
}
func F_clusterNodeGetShardId(m *base.Module, l0 int32) int32 {
	return l0 + int32(48)
}
func F_clusterNodeIsMyself(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_clusterNodeIsMyself[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	return base.B2i32(l0 == v4)
}
func F_clusterNodeIsVotingPrimary(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v2 = int32(0)
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v4&int32(1) == v2 {
		v12 = v2
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2160))
		v12 = base.B2i32(v9 != int32(0))
	}
	return v12
}
func F_clusterNodeIterNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v8 {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	default:
		goto L3
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return v150
L2:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	goto L43
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	F__serverPanic_1(m, int32(_a_F_clusterNodeIterNext_0), int32(354), int32(_a_F_clusterNodeIterNext_1), v6)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L41
	} else {
		goto L42
	}
L4:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v135 != 0 {
		goto L39
	} else {
		goto L40
	}
L5:
	;
	v120 = l0 + int32(8)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v122 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L6:
	;
	v10 = l0 + int32(8)
	v18 = l0 + int32(28)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v114 != 0 {
		goto L2
	} else {
		goto L33
	}
L8:
	;
	v25 = v18
	v26 = v22
	goto L11
L9:
	;
	v22 = int32(1)
	goto L8
L10:
	;
	v22 = int32(0)
	goto L8
L11:
	;
	switch v26 {
	case 0:
		goto L16
	default:
		goto L15
	}
L13:
	;
	v26 = int32(0)
	goto L11
L14:
	;
	goto L7
L15:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v106
	if v106 == int32(0) {
		goto L13
	} else {
		goto L32
	}
L16:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v30 != int32(-1) {
		v69 = v30
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v70 = int32(1)
	v71 = v69 + v70
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v71
	v73 = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v77+int32(26)))))
	if v81 == int32(255) {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v34 != 0 {
		v69 = int32(-1)
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v36 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	if v63 != int32(-1) {
		goto L23
	} else {
		goto L24
	}
L21:
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
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v62 = v61
	goto L20
L22:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)))
	v41 = v39 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v41)
	v62 = v35
	goto L20
L23:
	;
	v69 = v63 + int32(-1)
	goto L17
L24:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v69 = v66
	goto L17
L25:
	;
	v96 = int32(2)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v76+v94<<(uint(v96)%32)+int32(4))))
	v25 = v101 + v95<<(uint(v96)%32)
	v26 = int32(1)
	goto L11
L26:
	;
	v85 = v73
	goto L28
L27:
	;
	v85 = v70 << (uint(v81) % 32)
	goto L28
L28:
	;
	if v71 < v85 {
		v94 = v77
		v95 = v71
		goto L25
	} else {
		goto L29
	}
L29:
	;
	if v77 != 0 {
		v114 = v73
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	if v87 == int32(-1) {
		v114 = v73
		goto L14
	} else {
		goto L31
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(4294967296)
	v94 = int32(1)
	v95 = int32(0)
	goto L25
L32:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v110
	v114 = v106
	goto L14
L33:
	;
	v150 = int32(0)
	goto L1
L34:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v150 = v134
	goto L1
L35:
	;
	if v122 != 0 {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v122+base.B2i32(v125 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v131
	goto L36
L38:
	;
	v150 = int32(0)
	goto L1
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	v150 = v135
	goto L1
L40:
	;
	v150 = int32(0)
	goto L1
L41:
	;
	return int32(0)
L42:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	v150 = v148
	goto L1
}
func F_clusterNodeNumReplicas(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2164))
	return v2
}
func F_clusterNodePending(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	return v2 & int32(96)
}
func F_clusterRedirectBlockedClientIfNeeded(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v300 int32
	_ = v300
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_clusterRedirectBlockedClientIfNeeded[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	goto L1
L1:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v11&int32(16) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_clusterRedirectBlockedClientIfNeeded_0), int32(_a_F_clusterRedirectBlockedClientIfNeeded_1), int32(1382))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L12
	} else {
		goto L92
	}
L3:
	;
	return int32(0)
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	switch v17 + int32(-1) {
	case 0, 3, 4:
		goto L5
	default:
		goto L3
	case 2:
		goto L6
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_clusterRedirectBlockedClientIfNeeded[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	goto L10
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	goto L7
L7:
	;
	if v22 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v39 = F_dictGetIterator(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L15
	}
L10:
	;
	if v27 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_addReplyError(m, l0, int32(_a_F_clusterRedirectBlockedClientIfNeeded_2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	return int32(1)
L14:
	;
	F_dictReleaseIterator(m, v39)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L12
	} else {
		goto L91
	}
L15:
	;
	v48 = v39 + int32(20)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v49 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v144 == int32(0) {
		goto L14
	} else {
		goto L42
	}
L17:
	;
	v55 = v48
	v56 = v52
	goto L20
L18:
	;
	v52 = int32(1)
	goto L17
L19:
	;
	v52 = int32(0)
	goto L17
L20:
	;
	switch v56 {
	case 0:
		goto L25
	default:
		goto L24
	}
L22:
	;
	v56 = int32(0)
	goto L20
L23:
	;
	goto L16
L24:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v136
	if v136 == int32(0) {
		goto L22
	} else {
		goto L41
	}
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v60 != int32(-1) {
		v99 = v60
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v100 = int32(1)
	v101 = v99 + v100
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v101
	v103 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106+v107+int32(26)))))
	if v111 == int32(255) {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if v64 != 0 {
		v99 = int32(-1)
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v66 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
	if v93 != int32(-1) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v73 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v65)+16)))
	v74 = int64(*(*int8)(unsafe.Add(mBase, uint32(v65)+27)))
	v75 = int64(*(*int32)(unsafe.Add(mBase, uint32(v65)+8)))
	v76 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v65)+12)))
	v77 = int64(*(*int8)(unsafe.Add(mBase, uint32(v65)+26)))
	v78 = int64(*(*int32)(unsafe.Add(mBase, uint32(v65)+4)))
	v79 = F_wangHash64(m, v78)
	mBase = m.M
	v81 = F_wangHash64(m, v77+v79)
	mBase = m.M
	v83 = F_wangHash64(m, v76+v81)
	mBase = m.M
	v85 = F_wangHash64(m, v75+v83)
	mBase = m.M
	v87 = F_wangHash64(m, v74+v85)
	mBase = m.M
	v89 = F_wangHash64(m, v73+v87)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v39)+24)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v92 = v91
	goto L29
L31:
	;
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+24)))
	v71 = v69 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v65)+24)) = uint16(v71)
	v92 = v65
	goto L29
L32:
	;
	v99 = v93 + int32(-1)
	goto L26
L33:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v99 = v96
	goto L26
L34:
	;
	v126 = int32(2)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v106+v124<<(uint(v126)%32)+int32(4))))
	v55 = v131 + v125<<(uint(v126)%32)
	v56 = int32(1)
	goto L20
L35:
	;
	v115 = v103
	goto L37
L36:
	;
	v115 = v100 << (uint(v111) % 32)
	goto L37
L37:
	;
	if v101 < v115 {
		v124 = v107
		v125 = v101
		goto L34
	} else {
		goto L38
	}
L38:
	;
	if v107 != 0 {
		v144 = v103
		goto L23
	} else {
		goto L39
	}
L39:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	if v117 == int32(-1) {
		v144 = v103
		goto L23
	} else {
		goto L40
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+4)) = int64(4294967296)
	v124 = int32(1)
	v125 = int32(0)
	goto L34
L41:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v140
	v144 = v136
	goto L23
L42:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	goto L43
L43:
	;
	v151 = F_objectGetVal(m, v150)
	mBase = m.M
	v153 = F_objectGetVal(m, v150)
	mBase = m.M
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153+int32(-1)))))
	switch v156 & int32(7) {
	case 0:
		goto L49
	case 1:
		goto L48
	case 2:
		goto L47
	case 3:
		goto L46
	case 4:
		goto L45
	default:
		v173 = int32(0)
		goto L44
	}
L44:
	;
	v174 = int32(0)
	if v173 < int32(1) {
		v194 = v174
		goto L54
	} else {
		goto L55
	}
L45:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v153+int32(-17))))
	v173 = v172
	goto L44
L46:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v153+int32(-9))))
	v173 = v169
	goto L44
L47:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153+int32(-5)))))
	v173 = v166
	goto L44
L48:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153+int32(-3)))))
	v173 = v163
	goto L44
L49:
	;
	v173 = int32(base.Ui32(v156) >> (uint(int32(3)) % 32))
	goto L44
L50:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v240 != v241 {
		goto L2
	} else {
		goto L71
	}
L51:
	;
	v240 = v236 & int32(16383)
	goto L50
L52:
	;
	v205 = v194 + int32(1)
	if v173 <= v205 {
		goto L62
	} else {
		goto L63
	}
L53:
	;
	v203 = F_crc16(m, v151, v173)
	mBase = m.M
	v236 = v203
	goto L51
L54:
	;
	if v194 != v173 {
		goto L52
	} else {
		goto L60
	}
L55:
	;
	v182 = v174
	goto L56
L56:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v182))))
	if v186 == int32(123) {
		v194 = v182
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v190 = v182 + int32(1)
	if v190 != v173 {
		v182 = v190
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L53
L60:
	;
	goto L53
L61:
	;
	v233 = F_crc16(m, v151+v194+int32(1), v211+(v194^int32(-1)))
	mBase = m.M
	v236 = v233
	goto L51
L62:
	;
	v226 = F_crc16(m, v151, v173)
	mBase = m.M
	v236 = v226
	goto L51
L63:
	;
	v211 = v205
	goto L65
L64:
	;
	if v211 == v173 {
		goto L62
	} else {
		goto L69
	}
L65:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v211))))
	if v213 == int32(125) {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v217 = v211 + int32(1)
	if v217 != v173 {
		v211 = v217
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L62
L69:
	;
	if v211 != v205 {
		goto L61
	} else {
		goto L70
	}
L70:
	;
	goto L62
L71:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_clusterRedirectBlockedClientIfNeeded[0]))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v244+v240<<(uint(int32(2))%32)+int32(52))))
	goto L72
L72:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+202)))
	if v251&int32(2) == int32(0) {
		v269 = v250
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if v269 == v10 {
		goto L14
	} else {
		goto L82
	}
L74:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+56)))
	if v257&int32(1) != 0 {
		v269 = v250
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
	goto L76
L76:
	;
	if v260&int32(2) == int32(0) {
		v269 = v250
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v265 = F_clusterNodeGetPrimary(m, v10)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L12
	} else {
		goto L78
	}
L78:
	;
	if v265 == v250 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v268 = v10
	goto L81
L80:
	;
	v268 = v250
	goto L81
L81:
	;
	v269 = v268
	goto L73
L82:
	;
	v271 = F_getImportingSlotSource(m, v240)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L12
	} else {
		goto L83
	}
L83:
	;
	if v271 != 0 {
		goto L14
	} else {
		goto L84
	}
L84:
	;
	if v269 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	F_dictReleaseIterator(m, v39)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L12
	} else {
		goto L90
	}
L86:
	;
	F_clusterRedirectClient(m, l0, v269, v240, int32(4))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L12
	} else {
		goto L89
	}
L87:
	;
	F_addReplyError(m, l0, int32(_a_F_clusterRedirectBlockedClientIfNeeded_3))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L12
	} else {
		goto L88
	}
L88:
	;
	goto L85
L89:
	;
	goto L85
L90:
	;
	return int32(1)
L91:
	;
	goto L3
L92:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterReset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int64
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v324 int32
	_ = v324
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[0]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	F_dictEmpty(m, v12, v2)
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
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	F_dictEmpty(m, v18, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[0]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_clusterReset[1])))
	if v24 == int32(0) {
		v32 = v23
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_clusterReset[2]))) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_clusterReset[3]))) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_clusterReset[4]))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_clusterReset[1]))) = v33
	v43 = v2
	goto L7
L5:
	;
	F_unpauseActions(m, int32(2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[0]))
	v32 = v31
	goto L4
L7:
	;
	v45 = F_clusterDelSlot(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[0]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+36))
	F_dictEmpty(m, v53, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v48 = v43 + int32(1)
	if v48 != int32(16384) {
		v43 = v48
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[0]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+32))
	v60 = F_dictGetSafeIterator(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L14
L13:
	;
	F_dictReleaseIterator(m, v60)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L46
	}
L14:
	;
	v73 = v60 + int32(20)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	if v74 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v169 == int32(0) {
		goto L13
	} else {
		goto L42
	}
L17:
	;
	v80 = v73
	v81 = v77
	goto L20
L18:
	;
	v77 = int32(1)
	goto L17
L19:
	;
	v77 = int32(0)
	goto L17
L20:
	;
	switch v81 {
	case 0:
		goto L25
	default:
		goto L24
	}
L22:
	;
	v81 = int32(0)
	goto L20
L23:
	;
	goto L16
L24:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v161
	if v161 == int32(0) {
		goto L22
	} else {
		goto L41
	}
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v85 != int32(-1) {
		v124 = v85
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v125 = int32(1)
	v126 = v124 + v125
	*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v126
	v128 = int32(0)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v132+int32(26)))))
	if v136 == int32(255) {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v89 != 0 {
		v124 = int32(-1)
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if v91 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+20))
	if v118 != int32(-1) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v98 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v90)+16)))
	v99 = int64(*(*int8)(unsafe.Add(mBase, uint32(v90)+27)))
	v100 = int64(*(*int32)(unsafe.Add(mBase, uint32(v90)+8)))
	v101 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v90)+12)))
	v102 = int64(*(*int8)(unsafe.Add(mBase, uint32(v90)+26)))
	v103 = int64(*(*int32)(unsafe.Add(mBase, uint32(v90)+4)))
	v104 = F_wangHash64(m, v103)
	mBase = m.M
	v106 = F_wangHash64(m, v102+v104)
	mBase = m.M
	v108 = F_wangHash64(m, v101+v106)
	mBase = m.M
	v110 = F_wangHash64(m, v100+v108)
	mBase = m.M
	v112 = F_wangHash64(m, v99+v110)
	mBase = m.M
	v114 = F_wangHash64(m, v98+v112)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v117 = v116
	goto L29
L31:
	;
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+24)))
	v96 = v94 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v90)+24)) = uint16(v96)
	v117 = v90
	goto L29
L32:
	;
	v124 = v118 + int32(-1)
	goto L26
L33:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v124 = v121
	goto L26
L34:
	;
	v151 = int32(2)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v131+v149<<(uint(v151)%32)+int32(4))))
	v80 = v156 + v150<<(uint(v151)%32)
	v81 = int32(1)
	goto L20
L35:
	;
	v140 = v128
	goto L37
L36:
	;
	v140 = v125 << (uint(v136) % 32)
	goto L37
L37:
	;
	if v126 < v140 {
		v149 = v132
		v150 = v126
		goto L34
	} else {
		goto L38
	}
L38:
	;
	if v132 != 0 {
		v169 = v128
		goto L23
	} else {
		goto L39
	}
L39:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v131)+20))
	if v142 == int32(-1) {
		v169 = v128
		goto L23
	} else {
		goto L40
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v60)+4)) = int64(4294967296)
	v149 = int32(1)
	v150 = int32(0)
	goto L34
L41:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v165
	v169 = v161
	goto L23
L42:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
	goto L43
L43:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[5]))
	if v175 == v177 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	F_clusterDelNode(m, v175)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L14
L46:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[0]))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+40))
	F_dictEmpty(m, v185, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_clusterUpdateSlotExportsOnOwnershipChange(m)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_clusterUpdateSlotImportsOnOwnershipChange(m)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if l0 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	F__serverAssert(m, int32(_a_F_clusterReset_0), int32(_a_F_clusterReset_1), int32(2220))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L76
	}
L51:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[5]))
	F_clusterRemoveNodeFromShard(m, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L65
	}
L52:
	;
	v195 = int32(_a_F_clusterReset_2)
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[0]))
	v197 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v196)+uint32(_c_F_clusterReset[6]))) = v197
	*(*int64)(unsafe.Add(mBase, uint32(v196)+8)) = v197
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v202)+96)) = v197
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[7]))
	if int32(2) < v206 {
		v216 = v202
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v220 = F_sdsnewlen(m, v216+int32(8), int32(40))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	F__serverLog(m, int32(2), int32(_a_F_clusterReset_3), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[5]))
	v216 = v215
	goto L53
L56:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[0]))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+32))
	v225 = F_dictDelete(m, v224, v220)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_sdsfree(m, v220)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[5]))
	F_getRandomHexChars(m, v230+int32(8), int32(40))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[0]))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+32))
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[5]))
	v244 = F_sdsnewlen(m, v240+int32(8), int32(40))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v246 = F_dictAdd(m, v238, v244, v240)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v246 != 0 {
		goto L50
	} else {
		goto L62
	}
L62:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[7]))
	if int32(2) < v249 {
		goto L51
	} else {
		goto L63
	}
L63:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v253 + int32(8)
	F__serverLog(m, int32(2), int32(_a_F_clusterReset_4), v7+int32(16))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	goto L51
L65:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[5]))
	F_getRandomHexChars(m, v269+int32(48), int32(40))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[5]))
	F_clusterAddNodeToShard(m, v276+int32(48), v276)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[7]))
	if int32(2) < v282 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v294 = int32(0)
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[5]))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+88)))
	if v296&int32(2) == v294 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v286 + int32(48)
	F__serverLog(m, int32(2), int32(_a_F_clusterReset_5), v7)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L75
	}
L72:
	;
	F_clusterSetNodeAsPrimary(m, v295)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[8]))
	F_flushAllDataAndResetRDB(m, base.B2i32(v304 != int32(0)))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_clusterReset[0]))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+uint32(_c_F_clusterReset[9])))
	*(*int32)(unsafe.Add(mBase, uint32(v312)+uint32(_c_F_clusterReset[9]))) = v313 | int32(14)
	m.G0 = v7 + int32(32)
	return
L76:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterSaveConfig(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int64
	_ = v79
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v118 int64
	_ = v118
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v208 int64
	_ = v208
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v216 int64
	_ = v216
	var v222 int32
	_ = v222
	var v224 int64
	_ = v224
	var v225 int64
	_ = v225
	var v228 int64
	_ = v228
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v264 int64
	_ = v264
	var v271 int64
	_ = v271
	var v272 int32
	_ = v272
	var v274 int64
	_ = v274
	var v277 int64
	_ = v277
	var v283 int32
	_ = v283
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v290 int32
	_ = v290
	var v296 int64
	_ = v296
	var v300 int64
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v324 int64
	_ = v324
	var v327 int64
	_ = v327
	var v329 int64
	_ = v329
	var v332 int64
	_ = v332
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int64
	_ = v344
	var v350 int64
	_ = v350
	var v351 int64
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v425 int64
	_ = v425
	var v431 int32
	_ = v431
	var v433 int64
	_ = v433
	var v434 int32
	_ = v434
	var v436 int64
	_ = v436
	var v443 int64
	_ = v443
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v467 int64
	_ = v467
	var v472 int32
	_ = v472
	var v473 int64
	_ = v473
	var v475 int32
	_ = v475
	var v487 int64
	_ = v487
	var v492 int32
	_ = v492
	var v493 int64
	_ = v493
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v504 int64
	_ = v504
	var v507 int64
	_ = v507
	var v509 int64
	_ = v509
	var v512 int64
	_ = v512
	var v518 int32
	_ = v518
	var v528 int32
	_ = v528
	var v539 int64
	_ = v539
	var v545 int64
	_ = v545
	var v546 int64
	_ = v546
	var v547 int32
	_ = v547
	var v549 int64
	_ = v549
	var v552 int64
	_ = v552
	var v554 int64
	_ = v554
	var v557 int64
	_ = v557
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(128)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[0]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_clusterSaveConfig[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_clusterSaveConfig[1]))) = v21 & int32(-5)
	v29 = F_clusterGenNodesDescription(m, v2, int32(32), v2)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[0]))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_clusterSaveConfig[2])))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v34)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v35
	v42 = F_sdscatfmt(m, v29, int32(_a_F_clusterSaveConfig_0), v17+int32(112))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L9
	}
L3:
	;
	v64 = F_sdsempty(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L10
	}
L4:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(-17))))
	v63 = v62
	goto L3
L5:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(-9))))
	v63 = v59
	goto L3
L6:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42+int32(-5)))))
	v63 = v56
	goto L3
L7:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(-3)))))
	v63 = v53
	goto L3
L8:
	;
	v63 = int32(base.Ui32(v46) >> (uint(int32(3)) % 32))
	goto L3
L9:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(-1)))))
	switch v46 & int32(7) {
	case 0:
		goto L8
	case 1:
		goto L7
	case 2:
		goto L6
	case 3:
		goto L5
	case 4:
		goto L4
	default:
		v63 = v2
		goto L3
	}
L10:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[3]))
	v68 = F___syscall_getpid(m)
	mBase = m.M
	goto L11
L11:
	;
	v69 = F_mstime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v17)+104)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v67
	v76 = F_sdscatfmt(m, v64, int32(_a_F_clusterSaveConfig_1), v17+int32(96))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v79 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if base.B2i32(v79 == int64(0)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = int32(420)
	v89 = int32(-1)
	v93 = F_open(m, v76, int32(65), v17+int32(80))
	mBase = m.M
	if v93 != v89 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v85 = F_ustime(m)
	mBase = m.M
	v86 = v85
	goto L13
L15:
	;
	v86 = int64(0)
	goto L13
L16:
	;
	F_sdsfree(m, v76)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L130
	}
L17:
	;
	v539 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if base.B2i32(v539 == int64(0)) == int32(0) {
		goto L124
	} else {
		goto L125
	}
L18:
	;
	v108 = int64(0)
	v110 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if v110 == v108 {
		v133 = v108
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[5]))
	if int32(3) < v97 {
		v528 = v89
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L21
L21:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[6]))
	v102 = F___strerror_l(m, v101, v101)
	mBase = m.M
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v102
	F__serverLog(m, int32(3), int32(_a_F_clusterSaveConfig_2), v17)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v528 = v89
	goto L17
L24:
	;
	if v63 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L25:
	;
	v113 = F_ustime(m)
	mBase = m.M
	v115 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if v115 == int64(0) {
		v128 = v115
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if v128 == int64(0) {
		v133 = v108
		goto L24
	} else {
		goto L30
	}
L27:
	;
	v118 = v113 - v86
	if v118 < v115*int64(1000) {
		v128 = v115
		goto L26
	} else {
		goto L28
	}
L28:
	;
	F_latencyAddSample(m, int32(_a_F_clusterSaveConfig_3), v118)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v126 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	v128 = v126
	goto L26
L30:
	;
	v131 = F_ustime(m)
	mBase = m.M
	v133 = v131
	goto L24
L31:
	;
	v502 = F_close(m, v93)
	mBase = m.M
	v504 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if v504 == int64(0) {
		goto L117
	} else {
		goto L118
	}
L32:
	;
	if base.B2i32(v473 == int64(0)) == int32(0) {
		goto L115
	} else {
		goto L116
	}
L33:
	;
	v467 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	v472 = v456
	v473 = v467
	v475 = v459
	goto L32
L34:
	;
	v208 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if v208 == int64(0) {
		goto L52
	} else {
		goto L53
	}
L35:
	;
	v143 = v63
	v151 = int32(0)
	goto L36
L36:
	;
	goto L39
L37:
	;
	goto L34
L38:
	;
	v190 = v168 + v151
	if base.Ui32(v190) < base.Ui32(v63) {
		v143 = v63 - v190
		v151 = v190
		goto L36
	} else {
		goto L47
	}
L39:
	;
	v168 = F_write(m, v93, v42+v151, v143)
	mBase = m.M
	if int32(0) < v168 {
		goto L38
	} else {
		goto L41
	}
L40:
	;
	v175 = int32(-1)
	v176 = int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[5]))
	if int32(3) < v178 {
		v456 = v175
		v459 = v176
		goto L33
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[6]))
	if v172 == int32(27) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	v181 = F___strerror_l(m, v172, v172)
	mBase = m.M
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v151
	F__serverLog(m, int32(3), int32(_a_F_clusterSaveConfig_4), v17+int32(64))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v456 = v175
	v459 = v176
	goto L33
L47:
	;
	goto L37
L48:
	;
	v302 = int32(-1)
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[3]))
	v305 = F_rename(m, v76, v304)
	mBase = m.M
	if v305 != v302 {
		goto L75
	} else {
		goto L76
	}
L49:
	;
	if base.B2i32(v287 == int64(0)) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L50:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[0]))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+uint32(_c_F_clusterSaveConfig[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v239)+uint32(_c_F_clusterSaveConfig[1]))) = v240 & int32(-9)
	v244 = int32(-1)
	v245 = F_fsync(m, v93)
	mBase = m.M
	if v245 != v244 {
		goto L62
	} else {
		goto L63
	}
L51:
	;
	if base.B2i32(v225 == int64(0)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L52:
	;
	v228 = int64(0)
	if l0 != 0 {
		v237 = v228
		goto L50
	} else {
		goto L59
	}
L53:
	;
	v211 = F_ustime(m)
	mBase = m.M
	v213 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if v213 == int64(0) {
		v225 = v213
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if l0 != 0 {
		goto L51
	} else {
		goto L58
	}
L55:
	;
	v216 = v211 - v133
	if v216 < v213*int64(1000) {
		v225 = v213
		goto L54
	} else {
		goto L56
	}
L56:
	;
	F_latencyAddSample(m, int32(_a_F_clusterSaveConfig_5), v216)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v224 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	v225 = v224
	goto L54
L58:
	;
	v287 = v225
	v290 = int32(1)
	goto L49
L59:
	;
	v300 = v228
	v301 = int32(1)
	goto L48
L60:
	;
	v235 = F_ustime(m)
	mBase = m.M
	v237 = v235
	goto L50
L61:
	;
	v237 = int64(0)
	goto L50
L62:
	;
	v264 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if base.B2i32(v264 == int64(0)) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	v248 = int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[5]))
	if int32(3) < v250 {
		v456 = v244
		v459 = v248
		goto L33
	} else {
		goto L64
	}
L64:
	;
	goto L65
L65:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[6]))
	v255 = F___strerror_l(m, v254, v254)
	mBase = m.M
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v255
	F__serverLog(m, int32(3), int32(_a_F_clusterSaveConfig_6), v17+int32(48))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v456 = v244
	v459 = v248
	goto L33
L68:
	;
	v271 = F_ustime(m)
	mBase = m.M
	v272 = int32(0)
	v274 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if v274 == int64(0) {
		v287 = v274
		v290 = v272
		goto L49
	} else {
		goto L70
	}
L69:
	;
	v300 = int64(0)
	v301 = int32(0)
	goto L48
L70:
	;
	v277 = v271 - v237
	if v277 < v274*int64(1000) {
		v287 = v274
		v290 = v272
		goto L49
	} else {
		goto L71
	}
L71:
	;
	F_latencyAddSample(m, int32(_a_F_clusterSaveConfig_7), v277)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v285 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	v287 = v285
	v290 = v272
	goto L49
L73:
	;
	v296 = F_ustime(m)
	mBase = m.M
	v300 = v296
	v301 = v290
	goto L48
L74:
	;
	v300 = int64(0)
	v301 = v290
	goto L48
L75:
	;
	v324 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if v324 == int64(0) {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	v308 = int32(1)
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[5]))
	if int32(3) < v310 {
		v456 = v302
		v459 = v308
		goto L33
	} else {
		goto L77
	}
L77:
	;
	goto L78
L78:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[6]))
	v315 = F___strerror_l(m, v314, v314)
	mBase = m.M
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v315
	F__serverLog(m, int32(3), int32(_a_F_clusterSaveConfig_8), v17+int32(16))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v456 = v302
	v459 = v308
	goto L33
L81:
	;
	v341 = int32(0)
	if v301 != 0 {
		v456 = v341
		v459 = v341
		goto L33
	} else {
		goto L86
	}
L82:
	;
	v327 = F_ustime(m)
	mBase = m.M
	v329 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if v329 == int64(0) {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v332 = v327 - v300
	if v332 < v329*int64(1000) {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	F_latencyAddSample(m, int32(_a_F_clusterSaveConfig_9), v332)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	goto L81
L86:
	;
	v344 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if base.B2i32(v344 == int64(0)) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v352 = int32(-1)
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[3]))
	v358 = m.G0
	v360 = v358 - int32(4112)
	m.G0 = v360
	v362 = F_strlen(m, v354)
	mBase = m.M
	if base.Ui32(v362) < base.Ui32(int32(4097)) {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v350 = F_ustime(m)
	mBase = m.M
	v351 = v350
	goto L87
L89:
	;
	v351 = int64(0)
	goto L87
L90:
	;
	v425 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if base.B2i32(v425 == int64(0)) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L91:
	;
	if v401 != int32(-1) {
		goto L90
	} else {
		goto L104
	}
L92:
	;
	m.G0 = v360 + int32(4112)
	goto L91
L93:
	;
	v371 = F___memcpy(m, v360, v354, v362+int32(1))
	mBase = m.M
	v372 = F_dirname(m, v371)
	mBase = m.M
	v373 = int32(0)
	v375 = F_open(m, v372, v373, v373)
	mBase = m.M
	if v375 != int32(-1) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v365 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = int32(37)
	v401 = int32(-1)
	goto L92
L95:
	;
	v385 = F_fsync(m, v375)
	mBase = m.M
	if v385 != int32(-1) {
		goto L100
	} else {
		goto L101
	}
L96:
	;
	v380 = F___errno_location(m)
	mBase = m.M
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	if v381 != int32(31) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v384 = int32(-1)
	goto L99
L98:
	;
	v384 = int32(0)
	goto L99
L99:
	;
	v401 = v384
	goto L92
L100:
	;
	v399 = F_close(m, v375)
	mBase = m.M
	v401 = int32(0)
	goto L92
L101:
	;
	v388 = F___errno_location(m)
	mBase = m.M
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	if v389 == int32(8) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	if v389 == int32(28) {
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v394 = F_close(m, v375)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v388))) = v389
	v401 = int32(-1)
	goto L92
L104:
	;
	v409 = int32(1)
	v411 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[5]))
	if int32(3) < v411 {
		v456 = v352
		v459 = v409
		goto L33
	} else {
		goto L105
	}
L105:
	;
	goto L106
L106:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSaveConfig[6]))
	v416 = F___strerror_l(m, v415, v415)
	mBase = m.M
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v416
	F__serverLog(m, int32(3), int32(_a_F_clusterSaveConfig_10), v17+int32(32))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v456 = v352
	v459 = v409
	goto L33
L109:
	;
	v433 = F_ustime(m)
	mBase = m.M
	v434 = int32(0)
	v436 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if base.B2i32(v436 == int64(0)) == v434 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v431 = int32(0)
	v492 = v431
	v493 = int64(0)
	v495 = v431
	goto L31
L111:
	;
	v443 = v433 - v351
	if v443 < v436*int64(1000) {
		v472 = int32(0)
		v473 = v436
		v475 = v434
		goto L32
	} else {
		goto L113
	}
L112:
	;
	v472 = int32(0)
	v473 = v436
	v475 = v434
	goto L32
L113:
	;
	F_latencyAddSample(m, int32(_a_F_clusterSaveConfig_11), v443)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v450 = int32(0)
	v456 = v450
	v459 = v450
	goto L33
L115:
	;
	v487 = F_ustime(m)
	mBase = m.M
	v492 = v472
	v493 = v487
	v495 = v475
	goto L31
L116:
	;
	v492 = v472
	v493 = int64(0)
	v495 = v475
	goto L31
L117:
	;
	if v495 == int32(0) {
		v568 = v492
		goto L16
	} else {
		goto L122
	}
L118:
	;
	v507 = F_ustime(m)
	mBase = m.M
	v509 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if v509 == int64(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v512 = v507 - v493
	if v512 < v509*int64(1000) {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	F_latencyAddSample(m, int32(_a_F_clusterSaveConfig_12), v512)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	goto L117
L122:
	;
	v528 = v492
	goto L17
L123:
	;
	v547 = F_unlink(m, v76)
	mBase = m.M
	v549 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if v549 == int64(0) {
		v568 = v528
		goto L16
	} else {
		goto L126
	}
L124:
	;
	v545 = F_ustime(m)
	mBase = m.M
	v546 = v545
	goto L123
L125:
	;
	v546 = int64(0)
	goto L123
L126:
	;
	v552 = F_ustime(m)
	mBase = m.M
	v554 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSaveConfig[4]))
	if v554 == int64(0) {
		v568 = v528
		goto L16
	} else {
		goto L127
	}
L127:
	;
	v557 = v552 - v546
	if v557 < v554*int64(1000) {
		v568 = v528
		goto L16
	} else {
		goto L128
	}
L128:
	;
	F_latencyAddSample(m, int32(_a_F_clusterSaveConfig_13), v557)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v568 = v528
	goto L16
L130:
	;
	F_sdsfree(m, v42)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	m.G0 = v17 + int32(128)
	return v568
}
func F_clusterSendMFStart(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2344))
	if v4 == int32(0) {
		return
	} else {
		v9 = F_createClusterMsgSendBlock(m, int32(8), int32(2256))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2344))
			F_clusterSendMessage(m, v11, v9)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				v16 = v14 + int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v16
				if v14 <= int32(0) {
					F__serverAssert(m, int32(_a_F_clusterSendMFStart_0), int32(_a_F_clusterSendMFStart_1), int32(1756))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					if v16 != 0 {
						return
					} else {
						v20 = int32(_a_F_clusterSendMFStart_2)
						v22 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSendMFStart[0]))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
						*(*int32)(unsafe.Add(mBase, _c_F_clusterSendMFStart[0])) = v22 - v23
						F_valkey_free(m, v9)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_clusterSendPing(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v64 int32
	_ = v64
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int64
	_ = v194
	var v196 int32
	_ = v196
	var v200 int64
	_ = v200
	var v202 int32
	_ = v202
	var v206 int64
	_ = v206
	var v208 int32
	_ = v208
	var v212 int64
	_ = v212
	var v216 int64
	_ = v216
	var v220 int64
	_ = v220
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v224 int32
	_ = v224
	var v228 int64
	_ = v228
	var v230 int64
	_ = v230
	var v232 int32
	_ = v232
	var v236 int64
	_ = v236
	var v242 int64
	_ = v242
	var v248 int64
	_ = v248
	var v250 int32
	_ = v250
	var v254 int64
	_ = v254
	var v260 int64
	_ = v260
	var v266 int64
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int64
	_ = v365
	var v366 int64
	_ = v366
	var v367 int64
	_ = v367
	var v368 int64
	_ = v368
	var v369 int64
	_ = v369
	var v370 int64
	_ = v370
	var v371 int64
	_ = v371
	var v373 int64
	_ = v373
	var v375 int64
	_ = v375
	var v377 int64
	_ = v377
	var v379 int64
	_ = v379
	var v381 int64
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int64
	_ = v471
	var v473 int32
	_ = v473
	var v477 int64
	_ = v477
	var v479 int32
	_ = v479
	var v483 int64
	_ = v483
	var v485 int32
	_ = v485
	var v489 int64
	_ = v489
	var v493 int64
	_ = v493
	var v497 int64
	_ = v497
	var v498 int64
	_ = v498
	var v499 int64
	_ = v499
	var v501 int32
	_ = v501
	var v505 int64
	_ = v505
	var v507 int64
	_ = v507
	var v509 int32
	_ = v509
	var v513 int64
	_ = v513
	var v519 int64
	_ = v519
	var v525 int64
	_ = v525
	var v527 int32
	_ = v527
	var v531 int64
	_ = v531
	var v537 int64
	_ = v537
	var v543 int64
	_ = v543
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int64
	_ = v617
	var v618 int64
	_ = v618
	var v619 int64
	_ = v619
	var v620 int64
	_ = v620
	var v621 int64
	_ = v621
	var v622 int64
	_ = v622
	var v623 int64
	_ = v623
	var v625 int64
	_ = v625
	var v627 int64
	_ = v627
	var v629 int64
	_ = v629
	var v631 int64
	_ = v631
	var v633 int64
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v700 int32
	_ = v700
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSendPing[0]))
	if int32(0) < v19 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v58 = int32(0)
	v60 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSendPing[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_clusterSendPing[1])) = v60 + int64(1)
	v64 = int32(_a_F_clusterSendPing_0)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSendPing[2]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v69 = v67 + v68
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSendPing[3]))
	v74 = base.I32_div_u_s(v69*v71, int32(100))
	v75 = int32(3)
	if base.Ui32(v75) < base.Ui32(v74) {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	if base.Ui32(int32(10)) < base.Ui32(l1) {
		v30 = int32(_a_F_clusterSendPing_1)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v31 = int32(_a_F_clusterSendPing_2)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32))+uint32(_c_F_clusterSendPing[4])))
	v30 = v29
	goto L3
L5:
	;
	v36 = v32 + int32(8)
	goto L7
L6:
	;
	v36 = v31
	goto L7
L7:
	;
	if v32 == int32(0) {
		v41 = v31
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v30
	if v42 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v39 = F_humanNodename(m, v32)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v41 = v39
	goto L8
L12:
	;
	v48 = int32(_a_F_clusterSendPing_3)
	goto L14
L13:
	;
	v48 = int32(_a_F_clusterSendPing_4)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v48
	F__serverLog(m, int32(0), int32(_a_F_clusterSendPing_5), v16)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L1
L16:
	;
	v78 = v74
	goto L18
L17:
	;
	v78 = v75
	goto L18
L18:
	;
	v80 = v69 + int32(-2)
	if v78 < v80 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v82 = v78
	goto L21
L20:
	;
	v82 = v80
	goto L21
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_c_F_clusterSendPing[5])))
	v88 = (v82+v83)*int32(104) + int32(2256)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v89&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v108 = int32(4352)
	if v108 < v107 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v101 = int32(0)
	v103 = F_writePingExtensions(m, v101, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L27
	}
L24:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v92 == int32(0) {
		v107 = v88
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+89)))
	if v95&int32(4) == int32(0) {
		v107 = v88
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v107 = v103 + v88
	goto L22
L28:
	;
	v111 = v107
	goto L30
L29:
	;
	v111 = v108
	goto L30
L30:
	;
	v112 = F_createClusterMsgSendBlock(m, l1, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v114 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v125 = v112 + int32(8)
	v127 = v82 + int32(2)
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSendPing[2]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+32))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v133 = v131 + v132
	if v127 < v133 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	switch l1 {
	case 0:
		v117 = int32(2184)
		goto L34
	default:
		goto L32
	case 2:
		goto L35
	}
L34:
	;
	v118 = F_mstime(m)
	mBase = m.M
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v119+v117))) = v118
	goto L32
L35:
	;
	v117 = int32(2208)
	goto L34
L36:
	;
	v135 = v127
	goto L38
L37:
	;
	v135 = v133
	goto L38
L38:
	;
	v138 = F_valkey_malloc(m, v135<<(uint(int32(2))%32))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	v140 = int32(0)
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSendPing[2]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+32))
	v144 = F_dictGetSomeKeys(m, v143, v138, v135)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L10
	} else {
		goto L41
	}
L40:
	;
	F_valkey_free(m, v138)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L10
	} else {
		goto L64
	}
L41:
	;
	if v144 == int32(0) {
		v315 = v140
		goto L40
	} else {
		goto L42
	}
L42:
	;
	if v80 < int32(1) {
		v315 = v140
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v150 = int32(0)
	v156 = v150
	v157 = v150
	goto L44
L44:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v138+v157<<(uint(int32(2))%32))))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+8))
	goto L47
L45:
	;
	v315 = v305
	goto L40
L46:
	;
	v308 = v157 + int32(1)
	if base.Ui32(v144) <= base.Ui32(v308) {
		v315 = v305
		goto L40
	} else {
		goto L62
	}
L47:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSendPing[6]))
	if v169 == v171 {
		v305 = v156
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v169 == v173 {
		v305 = v156
		goto L46
	} else {
		goto L49
	}
L49:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+88)))
	if v175&int32(100) != 0 {
		v305 = v156
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v169)+2344))
	if v178 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v169)+2176))
	v184 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSendPing[1]))
	if v182 == v184 {
		v305 = v156
		goto L46
	} else {
		goto L54
	}
L52:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v169)+2160))
	if v179 == int32(0) {
		v305 = v156
		goto L46
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v189 = v125 + v156*int32(104)
	v190 = int32(2288)
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v169+int32(40))))
	*(*int64)(unsafe.Add(mBase, uint32(v189+v190))) = v194
	v196 = int32(2280)
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v169+int32(32))))
	*(*int64)(unsafe.Add(mBase, uint32(v189+v196))) = v200
	v202 = int32(2272)
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v169+int32(24))))
	*(*int64)(unsafe.Add(mBase, uint32(v189+v202))) = v206
	v208 = int32(2264)
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v169+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v189+v208))) = v212
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v169)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v189+int32(2256)))) = v216
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v169)+2184))
	v221 = int64(1000)
	v222 = base.I64_div_s(v220, v221)
	v224 = F_htonl(m, base.I32_wrap_i64(v222))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v189+int32(2296)))) = v224
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v169)+2192))
	v230 = base.I64_div_s(v228, v221)
	v232 = F_htonl(m, base.I32_wrap_i64(v230))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v189+int32(2300)))) = v232
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v169)+2256))
	*(*int64)(unsafe.Add(mBase, uint32(v189+int32(2304)))) = v236
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v169+v208)))
	*(*int64)(unsafe.Add(mBase, uint32(v189+int32(2312)))) = v242
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v169+v202)))
	*(*int64)(unsafe.Add(mBase, uint32(v189+int32(2320)))) = v248
	v250 = int32(2328)
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v169+v196)))
	*(*int64)(unsafe.Add(mBase, uint32(v189+v250))) = v254
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v169+v190)))
	*(*int64)(unsafe.Add(mBase, uint32(v189+int32(2336)))) = v260
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v169+int32(2294))))
	*(*int64)(unsafe.Add(mBase, uint32(v189+int32(2342)))) = v266
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSendPing[7]))
	if v271 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v169)+2176)) = v184
	v305 = v156 + int32(1)
	goto L46
L56:
	;
	v272 = v250
	goto L58
L57:
	;
	v272 = int32(2324)
	goto L58
L58:
	;
	v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169+v272))))
	v275 = F_htons(m, v274)
	mBase = m.M
	if v271 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v280 = int32(2324)
	goto L61
L60:
	;
	v280 = int32(2328)
	goto L61
L61:
	;
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169+v280))))
	v283 = F_htons(m, v282)
	mBase = m.M
	*(*uint16)(unsafe.Add(mBase, uint32(v189+int32(2356)))) = uint16(v283)
	*(*uint16)(unsafe.Add(mBase, uint32(v189+int32(2350)))) = uint16(v275)
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+2332)))
	v291 = F_htons(m, v290)
	mBase = m.M
	*(*uint16)(unsafe.Add(mBase, uint32(v189+int32(2352)))) = uint16(v291)
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+88)))
	v296 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v189+int32(2358)))) = uint16(v296)
	v300 = F_htons(m, v293)
	mBase = m.M
	*(*uint16)(unsafe.Add(mBase, uint32(v189+int32(2354)))) = uint16(v300)
	goto L55
L62:
	;
	if v305 < v82 {
		v156 = v305
		v157 = v308
		goto L44
	} else {
		goto L63
	}
L63:
	;
	goto L45
L64:
	;
	if v83 == int32(0) {
		v715 = v315
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v724&int32(1) != 0 {
		goto L140
	} else {
		goto L141
	}
L66:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSendPing[2]))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+32))
	v331 = F_dictGetSafeIterator(m, v330)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L10
	} else {
		goto L68
	}
L67:
	;
	F_dictReleaseIterator(m, v331)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L10
	} else {
		goto L137
	}
L68:
	;
	v340 = v331 + int32(20)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v331)+16))
	if v341 != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	if v436 == int32(0) {
		v700 = v315
		goto L67
	} else {
		goto L95
	}
L70:
	;
	v347 = v340
	v348 = v344
	goto L73
L71:
	;
	v344 = int32(1)
	goto L70
L72:
	;
	v344 = int32(0)
	goto L70
L73:
	;
	switch v348 {
	case 0:
		goto L78
	default:
		goto L77
	}
L75:
	;
	v348 = int32(0)
	goto L73
L76:
	;
	goto L69
L77:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	*(*int32)(unsafe.Add(mBase, uint32(v331)+16)) = v428
	if v428 == int32(0) {
		goto L75
	} else {
		goto L94
	}
L78:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v352 != int32(-1) {
		v391 = v352
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v392 = int32(1)
	v393 = v391 + v392
	*(*int32)(unsafe.Add(mBase, uint32(v331)+4)) = v393
	v395 = int32(0)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v331)+8))
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398+v399+int32(26)))))
	if v403 == int32(255) {
		goto L88
	} else {
		goto L89
	}
L80:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v331)+8))
	if v356 != 0 {
		v391 = int32(-1)
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v331)+12))
	if v358 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v384)+20))
	if v385 != int32(-1) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v365 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v357)+16)))
	v366 = int64(*(*int8)(unsafe.Add(mBase, uint32(v357)+27)))
	v367 = int64(*(*int32)(unsafe.Add(mBase, uint32(v357)+8)))
	v368 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v357)+12)))
	v369 = int64(*(*int8)(unsafe.Add(mBase, uint32(v357)+26)))
	v370 = int64(*(*int32)(unsafe.Add(mBase, uint32(v357)+4)))
	v371 = F_wangHash64(m, v370)
	mBase = m.M
	v373 = F_wangHash64(m, v369+v371)
	mBase = m.M
	v375 = F_wangHash64(m, v368+v373)
	mBase = m.M
	v377 = F_wangHash64(m, v367+v375)
	mBase = m.M
	v379 = F_wangHash64(m, v366+v377)
	mBase = m.M
	v381 = F_wangHash64(m, v365+v379)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v331)+24)) = v381
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v384 = v383
	goto L82
L84:
	;
	v361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v357)+24)))
	v363 = v361 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v357)+24)) = uint16(v363)
	v384 = v357
	goto L82
L85:
	;
	v391 = v385 + int32(-1)
	goto L79
L86:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	v391 = v388
	goto L79
L87:
	;
	v418 = int32(2)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v398+v416<<(uint(v418)%32)+int32(4))))
	v347 = v423 + v417<<(uint(v418)%32)
	v348 = int32(1)
	goto L73
L88:
	;
	v407 = v395
	goto L90
L89:
	;
	v407 = v392 << (uint(v403) % 32)
	goto L90
L90:
	;
	if v393 < v407 {
		v416 = v399
		v417 = v393
		goto L87
	} else {
		goto L91
	}
L91:
	;
	if v399 != 0 {
		v436 = v395
		goto L76
	} else {
		goto L92
	}
L92:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v398)+20))
	if v409 == int32(-1) {
		v436 = v395
		goto L76
	} else {
		goto L93
	}
L93:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v331)+4)) = int64(4294967296)
	v416 = int32(1)
	v417 = int32(0)
	goto L87
L94:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v428)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v340))) = v432
	v436 = v428
	goto L76
L95:
	;
	if v83 < int32(1) {
		v700 = v315
		goto L67
	} else {
		goto L96
	}
L96:
	;
	v447 = v83
	v448 = v315
	v449 = v436
	goto L97
L97:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v449)+8))
	goto L100
L98:
	;
	v700 = v584
	goto L67
L99:
	;
	v592 = v331 + int32(20)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v331)+16))
	if v593 != 0 {
		goto L111
	} else {
		goto L112
	}
L100:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)+88))
	if v458&int32(100) != int32(4) {
		v583 = v447
		v584 = v448
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v466 = v125 + v448*int32(104)
	v467 = int32(2288)
	v471 = *(*int64)(unsafe.Add(mBase, uint32(v457+int32(40))))
	*(*int64)(unsafe.Add(mBase, uint32(v466+v467))) = v471
	v473 = int32(2280)
	v477 = *(*int64)(unsafe.Add(mBase, uint32(v457+int32(32))))
	*(*int64)(unsafe.Add(mBase, uint32(v466+v473))) = v477
	v479 = int32(2272)
	v483 = *(*int64)(unsafe.Add(mBase, uint32(v457+int32(24))))
	*(*int64)(unsafe.Add(mBase, uint32(v466+v479))) = v483
	v485 = int32(2264)
	v489 = *(*int64)(unsafe.Add(mBase, uint32(v457+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v466+v485))) = v489
	v493 = *(*int64)(unsafe.Add(mBase, uint32(v457)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v466+int32(2256)))) = v493
	v497 = *(*int64)(unsafe.Add(mBase, uint32(v457)+2184))
	v498 = int64(1000)
	v499 = base.I64_div_s(v497, v498)
	v501 = F_htonl(m, base.I32_wrap_i64(v499))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v466+int32(2296)))) = v501
	v505 = *(*int64)(unsafe.Add(mBase, uint32(v457)+2192))
	v507 = base.I64_div_s(v505, v498)
	v509 = F_htonl(m, base.I32_wrap_i64(v507))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v466+int32(2300)))) = v509
	v513 = *(*int64)(unsafe.Add(mBase, uint32(v457)+2256))
	*(*int64)(unsafe.Add(mBase, uint32(v466+int32(2304)))) = v513
	v519 = *(*int64)(unsafe.Add(mBase, uint32(v457+v485)))
	*(*int64)(unsafe.Add(mBase, uint32(v466+int32(2312)))) = v519
	v525 = *(*int64)(unsafe.Add(mBase, uint32(v457+v479)))
	*(*int64)(unsafe.Add(mBase, uint32(v466+int32(2320)))) = v525
	v527 = int32(2328)
	v531 = *(*int64)(unsafe.Add(mBase, uint32(v457+v473)))
	*(*int64)(unsafe.Add(mBase, uint32(v466+v527))) = v531
	v537 = *(*int64)(unsafe.Add(mBase, uint32(v457+v467)))
	*(*int64)(unsafe.Add(mBase, uint32(v466+int32(2336)))) = v537
	v543 = *(*int64)(unsafe.Add(mBase, uint32(v457+int32(2294))))
	*(*int64)(unsafe.Add(mBase, uint32(v466+int32(2342)))) = v543
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSendPing[7]))
	if v548 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v583 = v447 + int32(-1)
	v584 = v448 + int32(1)
	goto L99
L103:
	;
	v549 = v527
	goto L105
L104:
	;
	v549 = int32(2324)
	goto L105
L105:
	;
	v551 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v457+v549))))
	v552 = F_htons(m, v551)
	mBase = m.M
	if v548 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v557 = int32(2324)
	goto L108
L107:
	;
	v557 = int32(2328)
	goto L108
L108:
	;
	v559 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v457+v557))))
	v560 = F_htons(m, v559)
	mBase = m.M
	*(*uint16)(unsafe.Add(mBase, uint32(v466+int32(2356)))) = uint16(v560)
	*(*uint16)(unsafe.Add(mBase, uint32(v466+int32(2350)))) = uint16(v552)
	v567 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v457)+2332)))
	v568 = F_htons(m, v567)
	mBase = m.M
	*(*uint16)(unsafe.Add(mBase, uint32(v466+int32(2352)))) = uint16(v568)
	v570 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v457)+88)))
	v573 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v466+int32(2358)))) = uint16(v573)
	v577 = F_htons(m, v570)
	mBase = m.M
	*(*uint16)(unsafe.Add(mBase, uint32(v466+int32(2354)))) = uint16(v577)
	goto L102
L109:
	;
	if v688 == int32(0) {
		v700 = v584
		goto L67
	} else {
		goto L135
	}
L110:
	;
	v599 = v592
	v600 = v596
	goto L113
L111:
	;
	v596 = int32(1)
	goto L110
L112:
	;
	v596 = int32(0)
	goto L110
L113:
	;
	switch v600 {
	case 0:
		goto L118
	default:
		goto L117
	}
L115:
	;
	v600 = int32(0)
	goto L113
L116:
	;
	goto L109
L117:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v599)))
	*(*int32)(unsafe.Add(mBase, uint32(v331)+16)) = v680
	if v680 == int32(0) {
		goto L115
	} else {
		goto L134
	}
L118:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v604 != int32(-1) {
		v643 = v604
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v644 = int32(1)
	v645 = v643 + v644
	*(*int32)(unsafe.Add(mBase, uint32(v331)+4)) = v645
	v647 = int32(0)
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v331)+8))
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650+v651+int32(26)))))
	if v655 == int32(255) {
		goto L128
	} else {
		goto L129
	}
L120:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v331)+8))
	if v608 != 0 {
		v643 = int32(-1)
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v331)+12))
	if v610 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v636)+20))
	if v637 != int32(-1) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	v617 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v609)+16)))
	v618 = int64(*(*int8)(unsafe.Add(mBase, uint32(v609)+27)))
	v619 = int64(*(*int32)(unsafe.Add(mBase, uint32(v609)+8)))
	v620 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v609)+12)))
	v621 = int64(*(*int8)(unsafe.Add(mBase, uint32(v609)+26)))
	v622 = int64(*(*int32)(unsafe.Add(mBase, uint32(v609)+4)))
	v623 = F_wangHash64(m, v622)
	mBase = m.M
	v625 = F_wangHash64(m, v621+v623)
	mBase = m.M
	v627 = F_wangHash64(m, v620+v625)
	mBase = m.M
	v629 = F_wangHash64(m, v619+v627)
	mBase = m.M
	v631 = F_wangHash64(m, v618+v629)
	mBase = m.M
	v633 = F_wangHash64(m, v617+v631)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v331)+24)) = v633
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v636 = v635
	goto L122
L124:
	;
	v613 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v609)+24)))
	v615 = v613 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v609)+24)) = uint16(v615)
	v636 = v609
	goto L122
L125:
	;
	v643 = v637 + int32(-1)
	goto L119
L126:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	v643 = v640
	goto L119
L127:
	;
	v670 = int32(2)
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v650+v668<<(uint(v670)%32)+int32(4))))
	v599 = v675 + v669<<(uint(v670)%32)
	v600 = int32(1)
	goto L113
L128:
	;
	v659 = v647
	goto L130
L129:
	;
	v659 = v644 << (uint(v655) % 32)
	goto L130
L130:
	;
	if v645 < v659 {
		v668 = v651
		v669 = v645
		goto L127
	} else {
		goto L131
	}
L131:
	;
	if v651 != 0 {
		v688 = v647
		goto L116
	} else {
		goto L132
	}
L132:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v650)+20))
	if v661 == int32(-1) {
		v688 = v647
		goto L116
	} else {
		goto L133
	}
L133:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v331)+4)) = int64(4294967296)
	v668 = int32(1)
	v669 = int32(0)
	goto L127
L134:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v680)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v592))) = v684
	v688 = v680
	goto L116
L135:
	;
	if int32(0) < v583 {
		v447 = v583
		v448 = v584
		v449 = v688
		goto L97
	} else {
		goto L136
	}
L136:
	;
	goto L98
L137:
	;
	v715 = v700
	goto L65
L138:
	;
	if int32(65535) <= v715 {
		goto L149
	} else {
		goto L150
	}
L139:
	;
	v741 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSendPing[0]))
	if int32(0) < v741 {
		goto L145
	} else {
		goto L146
	}
L140:
	;
	v736 = F_writePingExtensions(m, v125, v715)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L10
	} else {
		goto L144
	}
L141:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v727 == int32(0) {
		goto L139
	} else {
		goto L142
	}
L142:
	;
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v727)+89)))
	if v730&int32(4) == int32(0) {
		goto L139
	} else {
		goto L143
	}
L143:
	;
	goto L140
L144:
	;
	v754 = v736 + int32(2256)
	goto L138
L145:
	;
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+2261)))
	v751 = v749 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+2261)) = uint8(v751)
	v754 = int32(2256)
	goto L138
L146:
	;
	v744 = int32(0)
	F__serverLog(m, v744, int32(_a_F_clusterSendPing_6), v744)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L10
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	F__serverAssert(m, int32(_a_F_clusterSendPing_7), int32(_a_F_clusterSendPing_8), int32(1756))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L10
	} else {
		goto L159
	}
L149:
	;
	F__serverAssert(m, int32(_a_F_clusterSendPing_9), int32(_a_F_clusterSendPing_8), int32(4994))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L10
	} else {
		goto L158
	}
L150:
	;
	v759 = F___bswap_16_1(m, v715&int32(65535))
	mBase = m.M
	goto L151
L151:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v112)+22)) = uint16(v759)
	v764 = F___bswap_32_1(m, v754+v715*int32(104))
	mBase = m.M
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+12)) = v764
	F_clusterSendMessage(m, l0, v112)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L10
	} else {
		goto L153
	}
L153:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v770 = v768 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v770
	if v768 <= int32(0) {
		goto L148
	} else {
		goto L154
	}
L154:
	;
	if v770 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	m.G0 = v16 + int32(16)
	return
L156:
	;
	v774 = int32(_a_F_clusterSendPing_0)
	v776 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSendPing[8]))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	*(*int32)(unsafe.Add(mBase, _c_F_clusterSendPing[8])) = v776 - v777
	F_valkey_free(m, v112)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L10
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterSetNodeAsPrimary(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v13&int32(1) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSetNodeAsPrimary[0]))
	if int32(2) < v17 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2172))
	if v34 != 0 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v20 = F_humanNodename(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l0 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0 + int32(8)
	F__serverLog(m, int32(2), int32(_a_F_clusterSetNodeAsPrimary_0), v11)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	goto L3
L8:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L68
	}
L9:
	;
	F_replicationUnsetPrimary(m)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L5
	} else {
		goto L67
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2172)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v250&int32(-4) | int32(1)
	if l0 != v248 {
		goto L8
	} else {
		goto L66
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+2164))
	if v38 < int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSetNodeAsPrimary[1]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v248 = v36
	v250 = v37
	goto L10
L13:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSetNodeAsPrimary[1]))
	if l0 != v235 {
		goto L64
	} else {
		goto L65
	}
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+2168))
	v45 = int32(0)
	goto L15
L15:
	;
	v52 = v45 + int32(1)
	v55 = v41 + v45<<(uint(int32(2))%32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v56 != l0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L13
L17:
	;
	if v52 != v38 {
		v45 = v52
		goto L15
	} else {
		goto L63
	}
L18:
	;
	if v38 <= v52 {
		v216 = v38
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v218 = v216 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+2164)) = v218
	if v218 != 0 {
		goto L13
	} else {
		goto L62
	}
L20:
	;
	v59 = int32(2)
	v61 = v41 + v52<<(uint(v59)%32)
	v66 = (v38 + (v45 ^ int32(-1))) << (uint(v59) % 32)
	if v55 == v61 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v34)+2164))
	v216 = v215
	goto L19
L22:
	;
	goto L21
L23:
	;
	v70 = v66 + v55
	if base.Ui32(int32(0)-v66<<(uint(int32(1))%32)) < base.Ui32(v61-v70) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v80 = (v61 ^ v55) & int32(3)
	if base.Ui32(v61) <= base.Ui32(v55) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v77 = F___memcpy(m, v55, v61, v66)
	mBase = m.M
	goto L21
L26:
	;
	if v186 == int32(0) {
		goto L22
	} else {
		goto L58
	}
L27:
	;
	if base.Ui32(v164) <= base.Ui32(int32(3)) {
		v185 = v163
		v186 = v164
		v187 = v165
		goto L26
	} else {
		goto L54
	}
L28:
	;
	if v80 != 0 {
		v146 = v66
		goto L38
	} else {
		goto L39
	}
L29:
	;
	if v80 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v55&int32(3) != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v185 = v61
	v186 = v66
	v187 = v55
	goto L26
L32:
	;
	v87 = v61
	v88 = v66
	v89 = v55
	goto L34
L33:
	;
	v163 = v61
	v164 = v66
	v165 = v55
	goto L27
L34:
	;
	if v88 == int32(0) {
		goto L22
	} else {
		goto L36
	}
L36:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v93)
	v95 = int32(1)
	v96 = v87 + v95
	v98 = v88 + int32(-1)
	v100 = v89 + v95
	if v100&int32(3) == int32(0) {
		v163 = v96
		v164 = v98
		v165 = v100
		goto L27
	} else {
		goto L37
	}
L37:
	;
	v87 = v96
	v88 = v98
	v89 = v100
	goto L34
L38:
	;
	if v146 == int32(0) {
		goto L22
	} else {
		goto L50
	}
L39:
	;
	if v70&int32(3) == int32(0) {
		v126 = v66
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if base.Ui32(v126) <= base.Ui32(int32(3)) {
		v146 = v126
		goto L38
	} else {
		goto L46
	}
L41:
	;
	v111 = v66
	goto L42
L42:
	;
	if v111 == int32(0) {
		goto L22
	} else {
		goto L44
	}
L43:
	;
	v126 = v117
	goto L40
L44:
	;
	v117 = v111 + int32(-1)
	v118 = v55 + v117
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v117))))
	*(*uint8)(unsafe.Add(mBase, uint32(v118))) = uint8(v120)
	if v118&int32(3) != 0 {
		v111 = v117
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v133 = v126
	goto L47
L47:
	;
	v137 = v133 + int32(-4)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v61+v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v55+v137))) = v140
	if base.Ui32(int32(3)) < base.Ui32(v137) {
		v133 = v137
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v146 = v137
	goto L38
L49:
	;
	goto L48
L50:
	;
	v153 = v146
	goto L51
L51:
	;
	v157 = v153 + int32(-1)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v157))))
	*(*uint8)(unsafe.Add(mBase, uint32(v55+v157))) = uint8(v160)
	if v157 != 0 {
		v153 = v157
		goto L51
	} else {
		goto L53
	}
L53:
	;
	goto L22
L54:
	;
	v170 = v163
	v171 = v164
	v172 = v165
	goto L55
L55:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v174
	v176 = int32(4)
	v177 = v170 + v176
	v179 = v172 + v176
	v181 = v171 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v181) {
		v170 = v177
		v171 = v181
		v172 = v179
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v185 = v177
	v186 = v181
	v187 = v179
	goto L26
L57:
	;
	goto L56
L58:
	;
	v192 = v185
	v193 = v186
	v194 = v187
	goto L59
L59:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v196)
	v198 = int32(1)
	v203 = v193 + int32(-1)
	if v203 != 0 {
		v192 = v192 + v198
		v193 = v203
		v194 = v194 + v198
		goto L59
	} else {
		goto L61
	}
L60:
	;
	goto L22
L61:
	;
	goto L60
L62:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+88)) = v220 & int32(-257)
	goto L13
L63:
	;
	goto L16
L64:
	;
	v248 = v235
	v250 = v233 | int32(256)
	goto L10
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2172)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v233&int32(-4) | int32(1)
	goto L9
L66:
	;
	goto L9
L67:
	;
	goto L8
L68:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSetNodeAsPrimary[2]))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+uint32(_c_F_clusterSetNodeAsPrimary[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v283)+uint32(_c_F_clusterSetNodeAsPrimary[3]))) = v284 | int32(6)
	goto L1
}
func F_clusterSlotFailoverGranted(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSlotFailoverGranted[0]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_clusterSlotFailoverGranted[1])))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v13
	goto L1
L1:
	;
	v17 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v19 == v17 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v115
L3:
	;
	if v19 == int32(0) {
		v115 = v17
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19+base.B2i32(v22 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v28
	goto L4
L6:
	;
	v35 = v19
	goto L8
L7:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v37)+156))
	v115 = base.B2i32(v110 == int32(17))
	goto L2
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v38 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v99 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+156))
	if base.Ui32(int32(20)) < base.Ui32(v39) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)+164))
	v48 = v8 + int32(8)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v49
	goto L15
L13:
	;
	if int32(1)<<(uint(v39)%32)&int32(1835040) != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v54 = v8 + int32(8)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v56 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v56 == int32(0) {
		goto L10
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v56+base.B2i32(v59 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v65
	goto L17
L19:
	;
	v72 = v56
	goto L20
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if l0 < v75 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L10
L22:
	;
	v80 = v8 + int32(8)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v82 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if l0 <= v77 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	if v82 != 0 {
		v72 = v82
		goto L20
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v82+base.B2i32(v85 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v91
	goto L26
L28:
	;
	goto L21
L29:
	;
	if v99 != 0 {
		v35 = v99
		goto L8
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v99+base.B2i32(v102 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v108
	goto L30
L32:
	;
	v115 = v17
	goto L2
}
func F_clusterSlotMigrationCron(m *base.Module) {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v243 int32
	_ = v243
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSlotMigrationCron[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_clusterSlotMigrationCron[1])))
	v13 = v7 + int32(16)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v14
	goto L1
L1:
	;
	v19 = v7 + int32(16)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v21 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSlotMigrationCron[2]))
	if v139 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L3:
	;
	if v21 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21+base.B2i32(v24 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v30
	goto L4
L6:
	;
	v35 = v21
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+156))
	if base.Ui32(int32(20)) < base.Ui32(v39) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L2
L9:
	;
	v121 = v7 + int32(16)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v123 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L10:
	;
	F_proceedWithSlotMigration(m, v38)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L19
	} else {
		goto L27
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v46 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	if int32(1)<<(uint(v39)%32)&int32(1966176) != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
	if v61 == int64(0) {
		goto L10
	} else {
		goto L21
	}
L15:
	;
	v60 = v49 + int32(88)
	goto L14
L16:
	;
	v60 = v38 + int32(24)
	goto L14
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+152))
	if v49 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	F__serverAssert(m, int32(_a_F_clusterSlotMigrationCron_0), int32(_a_F_clusterSlotMigrationCron_1), int32(2533))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	v64 = int32(_a_F_clusterSlotMigrationCron_2)
	v65 = *(*int64)(unsafe.Add(mBase, _c_F_clusterSlotMigrationCron[3]))
	v68 = int64(*(*int32)(unsafe.Add(mBase, _c_F_clusterSlotMigrationCron[4])))
	if v65-v61 <= v68 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSlotMigrationCron[5]))
	if int32(3) < v71 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_finishSlotMigrationJob(m, v38, int32(18), int32(_a_F_clusterSlotMigrationCron_3))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L19
	} else {
		goto L26
	}
L24:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v38)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v74
	F__serverLog(m, int32(3), int32(_a_F_clusterSlotMigrationCron_4), v7)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	goto L9
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v38)+156))
	if base.Ui32(int32(20)) < base.Ui32(v88) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v95 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if int32(1)<<(uint(v88)%32)&int32(1978337) != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSlotMigrationCron[6]))
	v105 = base.I32_div_s(int32(1000), v104)
	if int32(999) < v105 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v38)+152))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+202)))
	if v99&int32(64) != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	F_sendSyncSlotsMessage(m, v38, int32(_a_F_clusterSlotMigrationCron_5))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L19
	} else {
		goto L37
	}
L35:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSlotMigrationCron[7]))
	v112 = base.I32_div_s(int32(1000), base.I32_extend16_s(v105))
	v114 = base.I32_rem_s(v109, base.I32_extend16_s(v112))
	if v114 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	goto L9
L38:
	;
	if v123 != 0 {
		v35 = v123
		goto L7
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v123+base.B2i32(v126 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v132
	goto L39
L41:
	;
	goto L8
L42:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSlotMigrationCron[8]))
	goto L58
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSlotMigrationCron[0]))
	if v143 == int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSlotMigrationCron[9]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v143)+uint32(_c_F_clusterSlotMigrationCron[1])))
	v150 = v7 + int32(24)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v151
	goto L45
L45:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSlotMigrationCron[0]))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+uint32(_c_F_clusterSlotMigrationCron[1])))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+20))
	if base.Ui32(v158) <= base.Ui32(v147) {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L47
L47:
	;
	v165 = v7 + int32(24)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if v167 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L42
L49:
	;
	if v167 == int32(0) {
		goto L42
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v167+base.B2i32(v170 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v176
	goto L50
L52:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+156))
	if base.Ui32(int32(2)) < base.Ui32(v181+int32(-18)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSlotMigrationCron[0]))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+uint32(_c_F_clusterSlotMigrationCron[1])))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+20))
	if base.Ui32(v147) < base.Ui32(v194) {
		goto L47
	} else {
		goto L56
	}
L54:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSlotMigrationCron[0]))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_clusterSlotMigrationCron[1])))
	F_listDelNode(m, v188, v167)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L19
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	goto L48
L57:
	;
	m.G0 = v7 + int32(32)
	return
L58:
	;
	if v205 == int32(0) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSlotMigrationCron[0]))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+uint32(_c_F_clusterSlotMigrationCron[1])))
	v212 = v7 + int32(24)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v212)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v213
	goto L60
L60:
	;
	goto L62
L61:
	;
	F_unpauseActions(m, int32(3))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L19
	} else {
		goto L69
	}
L62:
	;
	v222 = v7 + int32(24)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	if v224 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v224 == int32(0) {
		goto L61
	} else {
		goto L67
	}
L65:
	;
	goto L64
L66:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v224+base.B2i32(v227 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = v233
	goto L65
L67:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v224)+8))
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v237)+176))
	if v238 == int64(0) {
		goto L62
	} else {
		goto L68
	}
L68:
	;
	goto L57
L69:
	;
	goto L57
}
func F_clusterSlotMigrationShouldInstallWriteHandler(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	v3 = int32(1)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v4 == int32(0) {
		v11 = v3
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		if v7 != 0 {
			v11 = v3
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+156))
			v11 = base.B2i32(v8 != int32(13))
		}
	}
	return v11
}
func F_clusterSlotStatResetAll(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_clusterSlotStatResetAll[0]))
	v8 = F__emscripten_memset_bulkmem(m, v2+int32(67968), base.I32_extend8_s(int32(0)), int32(393216))
	mBase = m.M
	return
}
func F_clusterUpdateMyselfAvailabilityZone(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_clusterUpdateMyselfAvailabilityZone[0]))
	if v3 == v1 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_clusterUpdateMyselfAvailabilityZone[1]))
		F_updateSdsExtensionField(m, v3+int32(2320), v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_clusterUpdateMyselfClientIpV4(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_clusterUpdateMyselfClientIpV4[0]))
	if v3 == v1 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_clusterUpdateMyselfClientIpV4[1]))
		F_updateSdsExtensionField(m, v3+int32(2304), v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_clusterUpdateMyselfFlags(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v1 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_clusterUpdateMyselfFlags[0]))
	if v5 == v1 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+88))
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_clusterUpdateMyselfFlags[1]))
		v19 = v8&int32(-7681) | base.B2i32(v12 != int32(0))<<(uint(int32(9))%32) | int32(7168)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+88)) = v19
		if v19 == v8 {
			return
		} else {
			F_clearCachedClusterSlotsResponse(m)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_clusterUpdateMyselfFlags[2]))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_clusterUpdateMyselfFlags[3])))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_clusterUpdateMyselfFlags[3]))) = v26 | int32(6)
				return
			}
		}
	}
}
func F_createClusterNode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v33 int64
	_ = v33
	var v39 int64
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int64
	_ = v116
	v7 = F_valkey_malloc(m, int32(2360))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = v7 + int32(8)
		if l0 == int32(0) {
			F_getRandomHexChars(m, v12, int32(40))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				F_getRandomHexChars(m, v7+int32(48), int32(40))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v49 = F_mstime(m)
					mBase = m.M
					v50 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+96)) = v50
					*(*int64)(unsafe.Add(mBase, uint32(v7))) = v49
					*(*int64)(unsafe.Add(mBase, uint32(v7)+2344)) = v50
					*(*int64)(unsafe.Add(mBase, uint32(v7)+2232)) = v50
					*(*int32)(unsafe.Add(mBase, uint32(v7)+88)) = l1
					v63 = F__emscripten_memset_bulkmem(m, v7+int32(104), base.I32_extend8_s(int32(0)), int32(2120))
					mBase = m.M
					v64 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+2256)) = v64
					*(*int64)(unsafe.Add(mBase, uint32(v7)+2240)) = v49
					*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2264)))) = v64
					*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2272)))) = v64
					*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2280)))) = v64
					*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2288)))) = v64
					*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2294)))) = v64
					v87 = F_sdsempty(m)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+2304)) = v87
						v90 = F_sdsempty(m)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+2308)) = v90
							v93 = F_sdsempty(m)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+2312)) = v93
								v96 = F_sdsempty(m)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+2316)) = v96
									v99 = F_sdsempty(m)
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int32(0)
									} else {
										v101 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v7)+2324)) = v101
										*(*int32)(unsafe.Add(mBase, uint32(v7)+2320)) = v99
										*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2332)))) = v101
										*(*int32)(unsafe.Add(mBase, uint32(v7+int32(2340)))) = int32(0)
										v112 = F_raxNew(m)
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+2356)) = int32(0)
											v116 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v7)+2248)) = v116
											*(*int64)(unsafe.Add(mBase, uint32(v7)+2224)) = v116
											*(*int32)(unsafe.Add(mBase, uint32(v7)+2352)) = v112
											return v7
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v12))) = v15
			v17 = int32(32)
			v21 = *(*int64)(unsafe.Add(mBase, uint32(l0+v17)))
			*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v21
			v23 = int32(24)
			v27 = *(*int64)(unsafe.Add(mBase, uint32(l0+v23)))
			*(*int64)(unsafe.Add(mBase, uint32(v7+v17))) = v27
			v29 = int32(16)
			v33 = *(*int64)(unsafe.Add(mBase, uint32(l0+v29)))
			*(*int64)(unsafe.Add(mBase, uint32(v7+v23))) = v33
			v39 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(8))))
			*(*int64)(unsafe.Add(mBase, uint32(v7+v29))) = v39
			F_getRandomHexChars(m, v7+int32(48), int32(40))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				v49 = F_mstime(m)
				mBase = m.M
				v50 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+96)) = v50
				*(*int64)(unsafe.Add(mBase, uint32(v7))) = v49
				*(*int64)(unsafe.Add(mBase, uint32(v7)+2344)) = v50
				*(*int64)(unsafe.Add(mBase, uint32(v7)+2232)) = v50
				*(*int32)(unsafe.Add(mBase, uint32(v7)+88)) = l1
				v63 = F__emscripten_memset_bulkmem(m, v7+int32(104), base.I32_extend8_s(int32(0)), int32(2120))
				mBase = m.M
				v64 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+2256)) = v64
				*(*int64)(unsafe.Add(mBase, uint32(v7)+2240)) = v49
				*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2264)))) = v64
				*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2272)))) = v64
				*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2280)))) = v64
				*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2288)))) = v64
				*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2294)))) = v64
				v87 = F_sdsempty(m)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+2304)) = v87
					v90 = F_sdsempty(m)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+2308)) = v90
						v93 = F_sdsempty(m)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+2312)) = v93
							v96 = F_sdsempty(m)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+2316)) = v96
								v99 = F_sdsempty(m)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									v101 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v7)+2324)) = v101
									*(*int32)(unsafe.Add(mBase, uint32(v7)+2320)) = v99
									*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2332)))) = v101
									*(*int32)(unsafe.Add(mBase, uint32(v7+int32(2340)))) = int32(0)
									v112 = F_raxNew(m)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+2356)) = int32(0)
										v116 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v7)+2248)) = v116
										*(*int64)(unsafe.Add(mBase, uint32(v7)+2224)) = v116
										*(*int32)(unsafe.Add(mBase, uint32(v7)+2352)) = v112
										return v7
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
func F_generateClusterSlotResponse(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v8 = F_createCachedResponseClient(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = F_addReplyDeferredLen(m, v8)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = int32(0)
	v17 = v15
	v18 = int32(0)
	v21 = int32(-1)
	v22 = v15
	goto L5
L4:
	;
	v73 = F_aggregateClientOutputBuffer(m, v8)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L23
	}
L5:
	;
	if v22 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v17 = v17 + int32(1)
	v18 = v68
	v21 = v69
	v22 = v70
	goto L5
L8:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_generateClusterSlotResponse[0]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61+v17<<(uint(int32(2))%32)+int32(52))))
	goto L22
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_generateClusterSlotResponse[0]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45+v17<<(uint(int32(2))%32)+int32(52))))
	goto L19
L10:
	;
	F_setDeferredArrayLen(m, v8, v12, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L16
	}
L11:
	;
	if v17 != int32(16384) {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	if v17 == int32(16384) {
		v32 = v18
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v59 = v18
	goto L8
L14:
	;
	F_addNodeReplyForClusterSlot(m, v8, v22, v21, int32(16383))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v32 = v18 + int32(1)
	goto L10
L16:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+180))
	if v35 == int32(0) {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	F__serverAssert(m, int32(_a_F_generateClusterSlotResponse_0), int32(_a_F_generateClusterSlotResponse_1), int32(1546))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	if v22 == v51 {
		v68 = v18
		v69 = v21
		v70 = v22
		goto L7
	} else {
		goto L20
	}
L20:
	;
	F_addNodeReplyForClusterSlot(m, v8, v22, v21, v17+int32(-1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v59 = v18 + int32(1)
	goto L8
L22:
	;
	v68 = v59
	v69 = v17
	v70 = v67
	goto L7
L23:
	;
	F_deleteCachedResponseClient(m, v8)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	return v73
}
func F_getClusterNodesList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int64
	_ = v158
	var v160 int32
	_ = v160
	var v164 int64
	_ = v164
	var v166 int32
	_ = v166
	var v170 int64
	_ = v170
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v219 int64
	_ = v219
	var v220 int64
	_ = v220
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v228 int64
	_ = v228
	var v230 int64
	_ = v230
	var v232 int64
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_getClusterNodesList[0]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v17 = F_valkey_malloc(m, (v10+v11)<<(uint(int32(2))%32)+int32(4))
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
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_getClusterNodesList[0]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	v25 = F_dictGetIterator(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v17+v294<<(uint(int32(2))%32)))) = int32(0)
	F_dictReleaseIterator(m, v25)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L65
	}
L4:
	;
	v34 = v25 + int32(20)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v130 == int32(0) {
		v294 = int32(0)
		goto L3
	} else {
		goto L31
	}
L6:
	;
	v41 = v34
	v42 = v38
	goto L9
L7:
	;
	v38 = int32(1)
	goto L6
L8:
	;
	v38 = int32(0)
	goto L6
L9:
	;
	switch v42 {
	case 0:
		goto L14
	default:
		goto L13
	}
L11:
	;
	v42 = int32(0)
	goto L9
L12:
	;
	goto L5
L13:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v122
	if v122 == int32(0) {
		goto L11
	} else {
		goto L30
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v46 != int32(-1) {
		v85 = v46
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v86 = int32(1)
	v87 = v85 + v86
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v87
	v89 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v93+int32(26)))))
	if v97 == int32(255) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v50 != 0 {
		v85 = int32(-1)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v52 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	if v79 != int32(-1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v59 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v51)+16)))
	v60 = int64(*(*int8)(unsafe.Add(mBase, uint32(v51)+27)))
	v61 = int64(*(*int32)(unsafe.Add(mBase, uint32(v51)+8)))
	v62 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v51)+12)))
	v63 = int64(*(*int8)(unsafe.Add(mBase, uint32(v51)+26)))
	v64 = int64(*(*int32)(unsafe.Add(mBase, uint32(v51)+4)))
	v65 = F_wangHash64(m, v64)
	mBase = m.M
	v67 = F_wangHash64(m, v63+v65)
	mBase = m.M
	v69 = F_wangHash64(m, v62+v67)
	mBase = m.M
	v71 = F_wangHash64(m, v61+v69)
	mBase = m.M
	v73 = F_wangHash64(m, v60+v71)
	mBase = m.M
	v75 = F_wangHash64(m, v59+v73)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v78 = v77
	goto L18
L20:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
	v57 = v55 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v57)
	v78 = v51
	goto L18
L21:
	;
	v85 = v79 + int32(-1)
	goto L15
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v85 = v82
	goto L15
L23:
	;
	v112 = int32(2)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v92+v110<<(uint(v112)%32)+int32(4))))
	v41 = v117 + v111<<(uint(v112)%32)
	v42 = int32(1)
	goto L9
L24:
	;
	v101 = v89
	goto L26
L25:
	;
	v101 = v86 << (uint(v97) % 32)
	goto L26
L26:
	;
	if v87 < v101 {
		v110 = v93
		v111 = v87
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if v93 != 0 {
		v130 = v89
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
	if v103 == int32(-1) {
		v130 = v89
		goto L12
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+4)) = int64(4294967296)
	v110 = int32(1)
	v111 = int32(0)
	goto L23
L30:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v126
	v130 = v122
	goto L12
L31:
	;
	v138 = v130
	v140 = int32(0)
	goto L32
L32:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	goto L35
L33:
	;
	v294 = v182
	goto L3
L34:
	;
	v191 = v25 + int32(20)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v192 != 0 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+88)))
	if v144&int32(96) != 0 {
		v182 = v140
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v151 = F_valkey_malloc(m, int32(40))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17+v140<<(uint(int32(2))%32)))) = v151
	v154 = int32(32)
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v143+int32(40))))
	*(*int64)(unsafe.Add(mBase, uint32(v151+v154))) = v158
	v160 = int32(24)
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v143+v154)))
	*(*int64)(unsafe.Add(mBase, uint32(v151+v160))) = v164
	v166 = int32(16)
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v143+v160)))
	*(*int64)(unsafe.Add(mBase, uint32(v151+v166))) = v170
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v143+v166)))
	*(*int64)(unsafe.Add(mBase, uint32(v151+int32(8)))) = v176
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v151))) = v178
	v182 = v140 + int32(1)
	goto L34
L38:
	;
	if v287 != 0 {
		v138 = v287
		v140 = v182
		goto L32
	} else {
		goto L64
	}
L39:
	;
	v198 = v191
	v199 = v195
	goto L42
L40:
	;
	v195 = int32(1)
	goto L39
L41:
	;
	v195 = int32(0)
	goto L39
L42:
	;
	switch v199 {
	case 0:
		goto L47
	default:
		goto L46
	}
L44:
	;
	v199 = int32(0)
	goto L42
L45:
	;
	goto L38
L46:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v279
	if v279 == int32(0) {
		goto L44
	} else {
		goto L63
	}
L47:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v203 != int32(-1) {
		v242 = v203
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v243 = int32(1)
	v244 = v242 + v243
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v244
	v246 = int32(0)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249+v250+int32(26)))))
	if v254 == int32(255) {
		goto L57
	} else {
		goto L58
	}
L49:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v207 != 0 {
		v242 = int32(-1)
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v209 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+20))
	if v236 != int32(-1) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v216 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v208)+16)))
	v217 = int64(*(*int8)(unsafe.Add(mBase, uint32(v208)+27)))
	v218 = int64(*(*int32)(unsafe.Add(mBase, uint32(v208)+8)))
	v219 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v208)+12)))
	v220 = int64(*(*int8)(unsafe.Add(mBase, uint32(v208)+26)))
	v221 = int64(*(*int32)(unsafe.Add(mBase, uint32(v208)+4)))
	v222 = F_wangHash64(m, v221)
	mBase = m.M
	v224 = F_wangHash64(m, v220+v222)
	mBase = m.M
	v226 = F_wangHash64(m, v219+v224)
	mBase = m.M
	v228 = F_wangHash64(m, v218+v226)
	mBase = m.M
	v230 = F_wangHash64(m, v217+v228)
	mBase = m.M
	v232 = F_wangHash64(m, v216+v230)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v232
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v235 = v234
	goto L51
L53:
	;
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208)+24)))
	v214 = v212 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v208)+24)) = uint16(v214)
	v235 = v208
	goto L51
L54:
	;
	v242 = v236 + int32(-1)
	goto L48
L55:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v242 = v239
	goto L48
L56:
	;
	v269 = int32(2)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v249+v267<<(uint(v269)%32)+int32(4))))
	v198 = v274 + v268<<(uint(v269)%32)
	v199 = int32(1)
	goto L42
L57:
	;
	v258 = v246
	goto L59
L58:
	;
	v258 = v243 << (uint(v254) % 32)
	goto L59
L59:
	;
	if v244 < v258 {
		v267 = v250
		v268 = v244
		goto L56
	} else {
		goto L60
	}
L60:
	;
	if v250 != 0 {
		v287 = v246
		goto L45
	} else {
		goto L61
	}
L61:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v249)+20))
	if v260 == int32(-1) {
		v287 = v246
		goto L45
	} else {
		goto L62
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+4)) = int64(4294967296)
	v267 = int32(1)
	v268 = int32(0)
	goto L56
L63:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v283
	v287 = v279
	goto L45
L64:
	;
	goto L33
L65:
	;
	return v17
}
func F_isClusterHealthy(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_isClusterHealthy[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	return base.B2i32(v3 == int32(0))
}
func F_updateClusterClientIpV6(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	F_clusterUpdateMyselfClientIpV6(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
