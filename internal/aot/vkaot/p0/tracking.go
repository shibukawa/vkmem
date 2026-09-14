package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_freeTrackingRadixTreeCallback(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_raxFree(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_trackingGetTotalPrefixes(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int64
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_trackingGetTotalPrefixes[0]))
	if v3 != 0 {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v3)+8))
		return v6
	} else {
		return int64(0)
	}
}
func F_trackingHandlePendingKeyInvalidations(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_trackingHandlePendingKeyInvalidations[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v15 == int32(0) {
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
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_trackingHandlePendingKeyInvalidations[1]))
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = v11 + int32(8)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
	goto L4
L4:
	;
	v27 = v11 + int32(8)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v29 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_trackingHandlePendingKeyInvalidations[0]))
	F_listEmpty(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L24
	} else {
		goto L47
	}
L6:
	;
	if v29 == int32(0) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29+base.B2i32(v32 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v38
	goto L7
L9:
	;
	v46 = v29
	goto L10
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_trackingHandlePendingKeyInvalidations[2]))
	if v55 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L5
L12:
	;
	v162 = v11 + int32(8)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v164 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L13:
	;
	F_decrRefCount(m, v53)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L24
	} else {
		goto L42
	}
L14:
	;
	if v53 == int32(0) {
		goto L12
	} else {
		goto L41
	}
L15:
	;
	if v53 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+224)))
	v99 = int32(2)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98<<(uint(v99)%32))+uint32(_c_F_trackingHandlePendingKeyInvalidations[3])))
	v103 = F_objectGetVal(m, v102)
	mBase = m.M
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_trackingHandlePendingKeyInvalidations[2]))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+224)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107<<(uint(v99)%32))+uint32(_c_F_trackingHandlePendingKeyInvalidations[3])))
	v112 = F_objectGetVal(m, v111)
	mBase = m.M
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+int32(-1)))))
	switch v115 & int32(7) {
	case 0:
		goto L35
	case 1:
		goto L34
	case 2:
		goto L33
	case 3:
		goto L32
	case 4:
		goto L31
	default:
		v144 = int32(0)
		goto L30
	}
L17:
	;
	v60 = F_objectGetVal(m, v53)
	mBase = m.M
	v62 = F_objectGetVal(m, v53)
	mBase = m.M
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+int32(-1)))))
	switch v65 & int32(7) {
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
		v94 = int32(0)
		goto L18
	}
L18:
	;
	F_sendTrackingMessage(m, v55, v60, v94, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L24
	} else {
		goto L29
	}
L19:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v62+int32(-17))))
	v94 = v93
	goto L18
L20:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v62+int32(-9))))
	F_sendTrackingMessage(m, v55, v60, v87, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L24
	} else {
		goto L28
	}
L21:
	;
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62+int32(-5)))))
	F_sendTrackingMessage(m, v55, v60, v81, int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L24
	} else {
		goto L27
	}
L22:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+int32(-3)))))
	F_sendTrackingMessage(m, v55, v60, v75, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L24
	} else {
		goto L26
	}
L23:
	;
	F_sendTrackingMessage(m, v55, v60, int32(base.Ui32(v65)>>(uint(int32(3))%32)), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return
L25:
	;
	goto L13
L26:
	;
	goto L13
L27:
	;
	goto L13
L28:
	;
	goto L13
L29:
	;
	goto L13
L30:
	;
	F_sendTrackingMessage(m, v55, v103, v144, int32(1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L24
	} else {
		goto L40
	}
L31:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v112+int32(-17))))
	v144 = v143
	goto L30
L32:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v112+int32(-9))))
	F_sendTrackingMessage(m, v55, v103, v137, int32(1))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L24
	} else {
		goto L39
	}
L33:
	;
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112+int32(-5)))))
	F_sendTrackingMessage(m, v55, v103, v131, int32(1))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L24
	} else {
		goto L38
	}
L34:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+int32(-3)))))
	F_sendTrackingMessage(m, v55, v103, v125, int32(1))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L24
	} else {
		goto L37
	}
L35:
	;
	F_sendTrackingMessage(m, v55, v103, int32(base.Ui32(v115)>>(uint(int32(3))%32)), int32(1))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L24
	} else {
		goto L36
	}
L36:
	;
	goto L12
L37:
	;
	goto L12
L38:
	;
	goto L12
L39:
	;
	goto L12
L40:
	;
	goto L12
L41:
	;
	goto L13
L42:
	;
	goto L12
L43:
	;
	if v164 != 0 {
		v46 = v164
		goto L10
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v164+base.B2i32(v167 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v173
	goto L44
L46:
	;
	goto L11
L47:
	;
	goto L1
}
func F_trackingLimitUsedSlots(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int64
	_ = v32
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(304)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_trackingLimitUsedSlots[0]))
	if v10 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(304)
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_trackingLimitUsedSlots[1]))
	if v14 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	goto L5
L4:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_trackingLimitUsedSlots[2]))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_trackingLimitUsedSlots[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(128)
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+12)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v7)+296)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v7)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v7 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+156)) = v7 + int32(168)
	goto L7
L5:
	;
	v18 = base.I64_extend_i32_u(v14)
	if base.Ui64(v18) < base.Ui64(v17) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v20 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_trackingLimitUsedSlots[2])) = v20
	goto L1
L7:
	;
	v44 = int32(100)
	v49 = v24*v44 + v44
	goto L9
L8:
	;
	F_raxStop(m, v7)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L12
	} else {
		goto L23
	}
L9:
	;
	if v49 < int32(1) {
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v81 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_trackingLimitUsedSlots[2])) = v81
	F_raxStop(m, v7)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L12
	} else {
		goto L22
	}
L11:
	;
	v55 = int32(0)
	v57 = F_raxSeek(m, v7, int32(_a_F_trackingLimitUsedSlots_0), v55, v55)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	v60 = F_raxRandomWalk(m, v7, int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	goto L15
L15:
	;
	if v62&int32(2) != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v68 = F_createStringObject_1(m, v66, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	F_trackingInvalidateKey(m, int32(0), v68, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	F_decrRefCount(m, v68)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_trackingLimitUsedSlots[0]))
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+8))
	goto L20
L20:
	;
	if base.Ui64(v18) < base.Ui64(v79) {
		v49 = v49 + int32(-1)
		goto L9
	} else {
		goto L21
	}
L21:
	;
	goto L10
L22:
	;
	goto L1
L23:
	;
	v88 = int32(0)
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_trackingLimitUsedSlots[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_trackingLimitUsedSlots[2])) = v90 + int32(1)
	goto L1
}
func F_trackingRememberKeyToBroadcast(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int64
	_ = v17
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	v6 = m.G0
	v8 = v6 - int32(304)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_trackingRememberKeyToBroadcast[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(128)
	v17 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v8)+296)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v8)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v8 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+156)) = v8 + int32(168)
	goto L1
L1:
	;
	v30 = int32(0)
	v32 = F_raxSeek(m, v8, int32(_a_F_trackingRememberKeyToBroadcast_0), v30, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	v34 = F_raxNext(m, v8)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L5
	}
L4:
	;
	F_raxStop(m, v8)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L33
	}
L5:
	;
	if v34 == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	goto L7
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if base.Ui32(l2) < base.Ui32(v43) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L4
L9:
	;
	v117 = F_raxNext(m, v8)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L31
	}
L10:
	;
	if v43 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v115 = F_raxInsert(m, v113, l1, l2, l0, int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L30
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if base.Ui32(v43) < base.Ui32(int32(4)) {
		v71 = v47
		v72 = l1
		v73 = v43
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if v111 != 0 {
		goto L9
	} else {
		goto L29
	}
L14:
	;
	v111 = int32(0)
	goto L13
L15:
	;
	v83 = v78
	v84 = v79
	v85 = v80
	goto L25
L16:
	;
	if v73 == int32(0) {
		goto L14
	} else {
		goto L23
	}
L17:
	;
	if (l1|v47)&int32(3) != 0 {
		v78 = v47
		v79 = l1
		v80 = v43
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v55 = v47
	v56 = l1
	v57 = v43
	goto L19
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v60 != v61 {
		v78 = v55
		v79 = v56
		v80 = v57
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v71 = v66
	v72 = v64
	v73 = v68
	goto L16
L21:
	;
	v63 = int32(4)
	v64 = v56 + v63
	v66 = v55 + v63
	v68 = v57 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v68) {
		v55 = v66
		v56 = v64
		v57 = v68
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v78 = v71
	v79 = v72
	v80 = v73
	goto L15
L24:
	;
	v111 = v88 - v89
	goto L13
L25:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v88 != v89 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v91 = int32(1)
	v96 = v85 + int32(-1)
	if v96 == int32(0) {
		goto L14
	} else {
		goto L28
	}
L28:
	;
	v83 = v83 + v91
	v84 = v84 + v91
	v85 = v96
	goto L25
L29:
	;
	goto L11
L30:
	;
	goto L9
L31:
	;
	if v117 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	goto L8
L33:
	;
	m.G0 = v8 + int32(304)
	return
}
