package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_addClientToTimeoutTable(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
	if v9 == int64(0) {
		m.G0 = v6 + int32(16)
		return
	} else {
		v12 = int64(56)
		v14 = int64(65280)
		v16 = int64(40)
		v19 = int64(16711680)
		v21 = int64(24)
		v23 = int64(4278190080)
		v25 = int64(8)
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v9<<(uint(v12)%64) | v9&v14<<(uint(v16)%64) | (v9&v19<<(uint(v21)%64) | v9&v23<<(uint(v25)%64)) | (int64(base.Ui64(v9)>>(uint(v25)%64))&v23 | int64(base.Ui64(v9)>>(uint(v21)%64))&v19 | (int64(base.Ui64(v9)>>(uint(v16)%64))&v14 | int64(base.Ui64(v9)>>(uint(v12)%64))))
		v49 = v6 | int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v49))) = l0
		v54 = *(*int32)(unsafe.Add(mBase, _consts[883]))
		v56 = int32(0)
		v58 = F_raxTryInsert(m, v54, v6, int32(16), v56, v56)
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return
		} else {
			if v58 == int32(0) {
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v62 | int32(512)
			}
			m.G0 = v6 + int32(16)
			return
		}
	}
}
func F_applyClientMaxMemoryUsage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
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
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	if v9 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return int32(1)
L2:
	;
	if v9 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	if v13 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L2
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[62]))
	v23 = v6 + int32(8)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v24
	goto L9
L6:
	;
	F_initServerClientMemUsageBuckets(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L5
L9:
	;
	v29 = v6 + int32(8)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v31 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	if v73 != 0 {
		goto L1
	} else {
		goto L26
	}
L11:
	;
	if v31 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31+base.B2i32(v34 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v40
	goto L12
L14:
	;
	v46 = v31
	goto L15
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v49 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	if v49 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L10
L17:
	;
	v56 = v6 + int32(8)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v58 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v53 = F_updateClientMemUsageAndBucket(m, v47)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	F_removeClientFromMemUsageBucket(m, v47, int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	goto L17
L22:
	;
	if v58 != 0 {
		v46 = v58
		goto L15
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v58+base.B2i32(v61 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v67
	goto L23
L25:
	;
	goto L16
L26:
	;
	F_freeServerClientMemUsageBuckets(m)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L1
}
func F_clientCachingCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
	if v6&int32(4) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v20 = F_objectGetVal(m, v19)
	mBase = m.M
	v21 = int32(_a421)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	F_addReplyErrorLength(m, l0, int32(_a811), int32(103))
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
	F_afterErrorReply(m, l0, int32(_a811), int32(103), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v155 | int32(128)
	v160 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L3
	} else {
		goto L49
	}
L7:
	;
	v72 = int32(_a422)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v75 != 0 {
		goto L27
	} else {
		goto L28
	}
L8:
	;
	if v56-v58 != 0 {
		goto L7
	} else {
		goto L20
	}
L9:
	;
	v56 = F_tolower(m, v52)
	mBase = m.M
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v58 = F_tolower(m, v57)
	mBase = m.M
	goto L8
L10:
	;
	v26 = v20
	v27 = v21
	v28 = v24
	goto L13
L11:
	;
	v52 = int32(0)
	v53 = v21
	goto L9
L12:
	;
	v52 = v49 & int32(255)
	v53 = v48
	goto L9
L13:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v30 == int32(0) {
		v48 = v27
		v49 = v28
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v48 = v42
	v49 = int32(0)
	goto L12
L15:
	;
	v34 = v28 & int32(255)
	if v34 == v30 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v41 = int32(1)
	v42 = v27 + v41
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if v43 != 0 {
		v26 = v26 + v41
		v27 = v42
		v28 = v43
		goto L13
	} else {
		goto L19
	}
L17:
	;
	v36 = F_tolower(m, v34)
	mBase = m.M
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v38 = F_tolower(m, v37)
	mBase = m.M
	if v36 == v38 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v48 = v27
	v49 = v40
	goto L12
L19:
	;
	goto L14
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v60&int32(32) != 0 {
		v155 = v60
		goto L6
	} else {
		goto L21
	}
L21:
	;
	F_addReplyErrorLength(m, l0, int32(_a812), int32(72))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	F_afterErrorReply(m, l0, int32(_a812), int32(72), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	return
L24:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReply(m, l0, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L41
	}
L25:
	;
	if v107-v109 != 0 {
		goto L24
	} else {
		goto L37
	}
L26:
	;
	v107 = F_tolower(m, v103)
	mBase = m.M
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v109 = F_tolower(m, v108)
	mBase = m.M
	goto L25
L27:
	;
	v77 = v20
	v78 = v72
	v79 = v75
	goto L30
L28:
	;
	v103 = int32(0)
	v104 = v72
	goto L26
L29:
	;
	v103 = v100 & int32(255)
	v104 = v99
	goto L26
L30:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v81 == int32(0) {
		v99 = v78
		v100 = v79
		goto L29
	} else {
		goto L32
	}
L31:
	;
	v99 = v93
	v100 = int32(0)
	goto L29
L32:
	;
	v85 = v79 & int32(255)
	if v85 == v81 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v92 = int32(1)
	v93 = v78 + v92
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	if v94 != 0 {
		v77 = v77 + v92
		v78 = v93
		v79 = v94
		goto L30
	} else {
		goto L36
	}
L34:
	;
	v87 = F_tolower(m, v85)
	mBase = m.M
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v89 = F_tolower(m, v88)
	mBase = m.M
	if v87 == v89 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v99 = v78
	v100 = v91
	goto L29
L36:
	;
	goto L31
L37:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v111&int32(64) != 0 {
		v155 = v111
		goto L6
	} else {
		goto L38
	}
L38:
	;
	F_addReplyErrorLength(m, l0, int32(_a813), int32(72))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	F_afterErrorReply(m, l0, int32(_a813), int32(72), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	return
L41:
	;
	v127 = F_objectGetVal(m, v124)
	mBase = m.M
	v129 = F_objectGetVal(m, v124)
	mBase = m.M
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+int32(-1)))))
	switch v132 & int32(7) {
	case 0:
		goto L47
	case 1:
		goto L46
	case 2:
		goto L45
	case 3:
		goto L44
	case 4:
		goto L43
	default:
		v149 = int32(0)
		goto L42
	}
L42:
	;
	F_afterErrorReply(m, l0, v127, v149+int32(-2), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L48
	}
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v129+int32(-17))))
	v149 = v148
	goto L42
L44:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v129+int32(-9))))
	v149 = v145
	goto L42
L45:
	;
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129+int32(-5)))))
	v149 = v142
	goto L42
L46:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+int32(-3)))))
	v149 = v139
	goto L42
L47:
	;
	v149 = int32(base.Ui32(v132) >> (uint(int32(3)) % 32))
	goto L42
L48:
	;
	return
L49:
	;
	return
}
func F_clientGetNameCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	if v4 == int32(0) {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
		v10 = F_prepareClientToWrite(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			if v9 != int32(2) {
				if v10 != 0 {
					return
				} else {
					F__addReplyToBufferOrList(m, l0, int32(_a794), int32(3))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				if v10 != 0 {
					return
				} else {
					F__addReplyToBufferOrList(m, l0, int32(_a795), int32(5))
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		F_addReplyBulk(m, l0, v4)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_clientHelpCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_addExtendedReplyHelp(m, l0, int32(_a808), int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_clientInfoCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v6 = F_sdsempty(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v9 = F_catClientInfoString(m, v6, l0, int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v13 = F_sdscatlen(m, v9, int32(_a397), int32(1))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-1)))))
				switch v17 & int32(7) {
				case 0:
					v34 = int32(base.Ui32(v17) >> (uint(int32(3)) % 32))
				case 1:
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
					v34 = v24
				case 2:
					v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
					v34 = v27
				case 3:
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
					v34 = v30
				case 4:
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
					v34 = v33
				default:
					v34 = int32(0)
				}
				F_addReplyVerbatim(m, l0, v13, v34, int32(_a683))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_sdsfree(m, v13)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_clientKillCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
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
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
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
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
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
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	v14 = F__emscripten_memset_bulkmem(m, v9, base.I32_extend8_s(int32(0)), int32(120))
	mBase = m.M
	goto L1
L1:
	;
	v15 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v21 != int32(3) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	F_freeClientFilter(m, v14)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L8
	} else {
		goto L55
	}
L3:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReply(m, l0, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L8
	} else {
		goto L47
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[62]))
	v38 = v14 + int32(120)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v39
	goto L11
L5:
	;
	if v21 < int32(4) {
		goto L3
	} else {
		goto L7
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v26 = F_objectGetVal(m, v25)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v26
	goto L4
L7:
	;
	v32 = F_parseClientFiltersOrReply(m, l0, v14)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v32 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	v44 = v14 + int32(120)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v46 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	F_addReplyErrorLength(m, l0, int32(_a810), int32(14))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L8
	} else {
		goto L45
	}
L13:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v122 | int32(64)
	goto L2
L14:
	;
	if v79 != 0 {
		goto L38
	} else {
		goto L39
	}
L15:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v101 == int32(3) {
		goto L12
	} else {
		goto L36
	}
L16:
	;
	if v46 == int32(0) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46+base.B2i32(v49 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v55
	goto L17
L19:
	;
	v59 = int32(0)
	v63 = v46
	v64 = v59
	v65 = v59
	goto L20
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	v68 = F_clientMatchesFilter(m, v67, v14)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L23
	}
L21:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v95 == int32(3) {
		goto L14
	} else {
		goto L33
	}
L22:
	;
	v82 = v14 + int32(120)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v84 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	if v68 == int32(0) {
		v79 = v64
		v80 = v65
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if l0 != v67 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v79 = v64 + int32(1)
	v80 = v76
	goto L22
L26:
	;
	v74 = F_freeClient(m, v67)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L28
	}
L27:
	;
	v76 = int32(1)
	goto L25
L28:
	;
	v76 = v65
	goto L25
L29:
	;
	if v84 != 0 {
		v63 = v84
		v64 = v79
		v65 = v80
		goto L20
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84+base.B2i32(v87 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v93
	goto L30
L32:
	;
	goto L21
L33:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v79))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	if v80 != 0 {
		goto L13
	} else {
		goto L35
	}
L35:
	;
	goto L2
L36:
	;
	F_addReplyLongLong(m, l0, int64(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	goto L2
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L43
	}
L39:
	;
	F_addReplyErrorLength(m, l0, int32(_a810), int32(14))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	F_afterErrorReply(m, l0, int32(_a810), int32(14), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	if v80 != 0 {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	goto L2
L43:
	;
	if v80 == int32(0) {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	goto L13
L45:
	;
	F_afterErrorReply(m, l0, int32(_a810), int32(14), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	goto L2
L47:
	;
	v139 = F_objectGetVal(m, v136)
	mBase = m.M
	v141 = F_objectGetVal(m, v136)
	mBase = m.M
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+int32(-1)))))
	switch v144 & int32(7) {
	case 0:
		goto L53
	case 1:
		goto L52
	case 2:
		goto L51
	case 3:
		goto L50
	case 4:
		goto L49
	default:
		v161 = int32(0)
		goto L48
	}
L48:
	;
	F_afterErrorReply(m, l0, v139, v161+int32(-2), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L54
	}
L49:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v141+int32(-17))))
	v161 = v160
	goto L48
L50:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v141+int32(-9))))
	v161 = v157
	goto L48
L51:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141+int32(-5)))))
	v161 = v154
	goto L48
L52:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+int32(-3)))))
	v161 = v151
	goto L48
L53:
	;
	v161 = int32(base.Ui32(v144) >> (uint(int32(3)) % 32))
	goto L48
L54:
	;
	goto L2
L55:
	;
	m.G0 = v14 + int32(128)
	return
}
func F_clientListCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v11 < int32(4) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v9 + int32(128)
	return
L2:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+int32(-1)))))
	switch v153 & int32(7) {
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
		v170 = int32(0)
		goto L44
	}
L3:
	;
	v142 = F_getAllClientsInfoString(m, int32(-1), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L9
	} else {
		goto L43
	}
L4:
	;
	F_freeClientFilter(m, v17)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L41
	}
L5:
	;
	if v11 == int32(2) {
		goto L3
	} else {
		goto L32
	}
L6:
	;
	v17 = F__emscripten_memset_bulkmem(m, v9, base.I32_extend8_s(int32(0)), int32(120))
	mBase = m.M
	goto L7
L7:
	;
	v18 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v18
	v22 = F_parseClientFiltersOrReply(m, l0, v17)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v28 = F_sdsempty(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L9
	} else {
		goto L13
	}
L9:
	;
	return
L10:
	;
	if v22 == int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	F_freeClientFilter(m, v17)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L1
L13:
	;
	v32 = v28 + int32(-1)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	switch v33 & int32(7) {
	case 0:
		goto L20
	case 1:
		goto L19
	case 2:
		goto L18
	case 3:
		goto L17
	case 4:
		goto L16
	default:
		goto L15
	}
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[62]))
	v59 = v17 + int32(120)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v60
	goto L21
L15:
	;
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v54)
	goto L14
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28+int32(-17)))) = int64(0)
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9)))) = int32(0)
	goto L15
L18:
	;
	v44 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))) = uint16(v44)
	goto L15
L19:
	;
	v40 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))) = uint8(v40)
	goto L15
L20:
	;
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v36)
	goto L15
L21:
	;
	v67 = v28
	goto L22
L22:
	;
	v71 = v17 + int32(120)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v73 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v73 == int32(0) {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v73+base.B2i32(v76 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v82
	goto L25
L27:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v87 = F_clientMatchesFilter(m, v86, v17)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	if v87 == int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v92 = F_catClientInfoString(m, v67, v86, int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v96 = F_sdscatlen(m, v92, int32(_a397), int32(1))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	v67 = v96
	goto L22
L32:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReply(m, l0, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	v104 = F_objectGetVal(m, v101)
	mBase = m.M
	v106 = F_objectGetVal(m, v101)
	mBase = m.M
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106+int32(-1)))))
	switch v109 & int32(7) {
	case 0:
		goto L39
	case 1:
		goto L38
	case 2:
		goto L37
	case 3:
		goto L36
	case 4:
		goto L35
	default:
		v126 = int32(0)
		goto L34
	}
L34:
	;
	F_afterErrorReply(m, l0, v104, v126+int32(-2), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L9
	} else {
		goto L40
	}
L35:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v106+int32(-17))))
	v126 = v125
	goto L34
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v106+int32(-9))))
	v126 = v122
	goto L34
L37:
	;
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106+int32(-5)))))
	v126 = v119
	goto L34
L38:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106+int32(-3)))))
	v126 = v116
	goto L34
L39:
	;
	v126 = int32(base.Ui32(v109) >> (uint(int32(3)) % 32))
	goto L34
L40:
	;
	goto L1
L41:
	;
	if v67 != 0 {
		v147 = v67
		goto L2
	} else {
		goto L42
	}
L42:
	;
	goto L3
L43:
	;
	v147 = v142
	goto L2
L44:
	;
	F_addReplyVerbatim(m, l0, v147, v170, int32(_a683))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L9
	} else {
		goto L50
	}
L45:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v147+int32(-17))))
	v170 = v169
	goto L44
L46:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v147+int32(-9))))
	v170 = v166
	goto L44
L47:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147+int32(-5)))))
	v170 = v163
	goto L44
L48:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+int32(-3)))))
	v170 = v160
	goto L44
L49:
	;
	v170 = int32(base.Ui32(v153) >> (uint(int32(3)) % 32))
	goto L44
L50:
	;
	F_sdsfree(m, v147)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	goto L1
}
func F_clientMatchesFilter(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int64
	_ = v169
	var v173 int64
	_ = v173
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
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
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v286 int32
	_ = v286
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int64
	_ = v443
	var v446 int32
	_ = v446
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v570 int32
	_ = v570
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v49 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v9 = F_getClientPeerId(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v17 == int32(0) {
		v40 = v16
		v41 = v17
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v41-v40&int32(255) == int32(0) {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	if v17 != v16&int32(255) {
		v40 = v16
		v41 = v17
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v23 = v9
	v24 = v13
	goto L9
L9:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v28 == int32(0) {
		v40 = v27
		v41 = v28
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v40 = v27
	v41 = v28
	goto L6
L11:
	;
	v31 = int32(1)
	if v28 == v27&int32(255) {
		v23 = v23 + v31
		v24 = v24 + v31
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	return int32(0)
L14:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v90 == int32(-1) {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	v52 = F_getClientSockname(m, l0)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v58 == int32(0) {
		v81 = v57
		v82 = v58
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v82-v81&int32(255) == int32(0) {
		goto L14
	} else {
		goto L25
	}
L18:
	;
	goto L17
L19:
	;
	if v58 != v57&int32(255) {
		v81 = v57
		v82 = v58
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v64 = v52
	v65 = v54
	goto L21
L21:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v69 == int32(0) {
		v81 = v68
		v82 = v69
		goto L18
	} else {
		goto L23
	}
L22:
	;
	v81 = v68
	v82 = v69
	goto L18
L23:
	;
	v72 = int32(1)
	if v69 == v68&int32(255) {
		v64 = v64 + v72
		v65 = v65 + v72
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	return int32(0)
L26:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v127 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L27:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v93&int32(1) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v120 == v118 {
		goto L26
	} else {
		goto L41
	}
L29:
	;
	v99 = int32(2)
	if v93&v99 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v118 = v90
	v120 = int32(3)
	goto L28
L31:
	;
	if v93&int32(262144) != 0 {
		v118 = v90
		v120 = v99
		goto L28
	} else {
		goto L34
	}
L32:
	;
	if v93&int32(4) != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v118 = v90
	v120 = int32(1)
	goto L28
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v109 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	goto L37
L36:
	;
	v118 = v90
	v120 = int32(0)
	goto L28
L37:
	;
	if v113 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v116 = int32(4)
	goto L40
L39:
	;
	v116 = int32(5)
	goto L40
L40:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v118 = v117
	v120 = v116
	goto L28
L41:
	;
	return int32(0)
L42:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v155 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L43:
	;
	v130 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	if base.Ui64(v130+int64(-32768)) < base.Ui64(int64(-65536)) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v152 != 0 {
		goto L42
	} else {
		goto L53
	}
L45:
	;
	goto L44
L46:
	;
	v141 = int32(4)
	goto L48
L47:
	;
	v141 = int32(2)
	goto L48
L48:
	;
	if base.Ui64(v130+int64(-2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v146 = int32(8)
	goto L51
L50:
	;
	v146 = v141
	goto L51
L51:
	;
	if base.Ui32(v133) < base.Ui32(v146) {
		v152 = int32(0)
		goto L45
	} else {
		goto L52
	}
L52:
	;
	v148 = int32(0)
	v149 = F_intsetSearch(m, v127, v130, v148)
	mBase = m.M
	v152 = base.B2i32(v149 != v148)
	goto L45
L53:
	;
	return int32(0)
L54:
	;
	v162 = int32(0)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v163 == v162 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v158 == v155 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	return int32(0)
L57:
	;
	return v611
L58:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if v169 == int64(0) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	if l0 == v167 {
		v611 = v162
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v180 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	if v180 == int64(0) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v173 = *(*int64)(unsafe.Add(mBase, _consts[78]))
	goto L63
L63:
	;
	v175 = base.I64_div_s(v173, int64(1000))
	v176 = *(*int64)(unsafe.Add(mBase, uint32(l0)+368))
	v178 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if v175-v176 < v178 {
		v611 = v162
		goto L57
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v191 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v184 = *(*int64)(unsafe.Add(mBase, _consts[78]))
	goto L67
L67:
	;
	v186 = base.I64_div_s(v184, int64(1000))
	v187 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	if v186-v187 < v189 {
		v611 = v162
		goto L57
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v197 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v194 = F_clientMatchesFlagFilter(m, l0, v191)
	mBase = m.M
	if v194 == int32(0) {
		v611 = v162
		goto L57
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v241 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L73:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	if v200 == int32(0) {
		v611 = v162
		goto L57
	} else {
		goto L74
	}
L74:
	;
	v203 = F_objectGetVal(m, v200)
	mBase = m.M
	if v203 == int32(0) {
		v611 = v162
		goto L57
	} else {
		goto L75
	}
L75:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v207 = F_objectGetVal(m, v206)
	mBase = m.M
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v212 == int32(0) {
		v235 = v211
		v236 = v212
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if v236-v235&int32(255) != 0 {
		v611 = v162
		goto L57
	} else {
		goto L84
	}
L77:
	;
	goto L76
L78:
	;
	if v212 != v211&int32(255) {
		v235 = v211
		v236 = v212
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v218 = v207
	v219 = v208
	goto L80
L80:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+1)))
	if v223 == int32(0) {
		v235 = v222
		v236 = v223
		goto L77
	} else {
		goto L82
	}
L81:
	;
	v235 = v222
	v236 = v223
	goto L77
L82:
	;
	v226 = int32(1)
	if v223 == v222&int32(255) {
		v218 = v218 + v226
		v219 = v219 + v226
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	goto L72
L85:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v250 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L86:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	if v244 == int32(0) {
		v611 = v162
		goto L57
	} else {
		goto L87
	}
L87:
	;
	v247 = F_compareStringObjects(m, v244, v241)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	if v247 != 0 {
		v611 = v162
		goto L57
	} else {
		goto L89
	}
L89:
	;
	goto L85
L90:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	if v259 == int32(-1) {
		goto L95
	} else {
		goto L96
	}
L91:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v253 == int32(0) {
		v611 = v162
		goto L57
	} else {
		goto L92
	}
L92:
	;
	v256 = F_compareStringObjects(m, v253, v250)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L3
	} else {
		goto L93
	}
L93:
	;
	if v256 != 0 {
		v611 = v162
		goto L57
	} else {
		goto L94
	}
L94:
	;
	goto L90
L95:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v265 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+28))
	if v263 != v259 {
		v611 = v162
		goto L57
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	if v321 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L99:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+int32(-1)))))
	v298 = int32(0)
	goto L101
L100:
	;
	if v308 == int32(0) {
		v611 = v162
		goto L57
	} else {
		goto L113
	}
L101:
	;
	switch v286 & int32(7) {
	case 0:
		goto L108
	case 1:
		goto L107
	case 2:
		goto L106
	case 3:
		goto L105
	case 4:
		goto L104
	default:
		v307 = int32(0)
		goto L103
	}
L102:
	;
	goto L100
L103:
	;
	v308 = base.B2i32(base.Ui32(v307) <= base.Ui32(v298))
	if base.Ui32(v307) <= base.Ui32(v298) {
		goto L109
	} else {
		goto L110
	}
L104:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v265+int32(-17))))
	v307 = v306
	goto L103
L105:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v265+int32(-9))))
	v307 = v305
	goto L103
L106:
	;
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v265+int32(-5)))))
	v307 = v304
	goto L103
L107:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+int32(-3)))))
	v307 = v303
	goto L103
L108:
	;
	v307 = int32(base.Ui32(v286) >> (uint(int32(3)) % 32))
	goto L103
L109:
	;
	goto L102
L110:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+v298))))
	if v310 != int32(114) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v313 = int32(1)
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+227)))
	if v315&v313 != 0 {
		v298 = v298 + v313
		goto L101
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	goto L98
L114:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v328 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v324 = F_clientMatchesIpFilter(m, l0, v321)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	if v324 == int32(0) {
		v611 = v162
		goto L57
	} else {
		goto L117
	}
L117:
	;
	goto L114
L118:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v367 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L119:
	;
	v331 = F_getClientPeerId(m, l0)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v337 == int32(0) {
		v360 = v336
		v361 = v337
		goto L122
	} else {
		goto L123
	}
L121:
	;
	if v361-v360&int32(255) == int32(0) {
		v611 = v162
		goto L57
	} else {
		goto L129
	}
L122:
	;
	goto L121
L123:
	;
	if v337 != v336&int32(255) {
		v360 = v336
		v361 = v337
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v343 = v331
	v344 = v333
	goto L125
L125:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+1)))
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+1)))
	if v348 == int32(0) {
		v360 = v347
		v361 = v348
		goto L122
	} else {
		goto L127
	}
L126:
	;
	v360 = v347
	v361 = v348
	goto L122
L127:
	;
	v351 = int32(1)
	if v348 == v347&int32(255) {
		v343 = v343 + v351
		v344 = v344 + v351
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	goto L118
L130:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v406 == int32(-1) {
		goto L142
	} else {
		goto L143
	}
L131:
	;
	v370 = F_getClientSockname(m, l0)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L3
	} else {
		goto L132
	}
L132:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
	if v376 == int32(0) {
		v399 = v375
		v400 = v376
		goto L134
	} else {
		goto L135
	}
L133:
	;
	if v400-v399&int32(255) == int32(0) {
		v611 = v162
		goto L57
	} else {
		goto L141
	}
L134:
	;
	goto L133
L135:
	;
	if v376 != v375&int32(255) {
		v399 = v375
		v400 = v376
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v382 = v370
	v383 = v372
	goto L137
L137:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+1)))
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+1)))
	if v387 == int32(0) {
		v399 = v386
		v400 = v387
		goto L134
	} else {
		goto L139
	}
L138:
	;
	v399 = v386
	v400 = v387
	goto L134
L139:
	;
	v390 = int32(1)
	if v387 == v386&int32(255) {
		v382 = v382 + v390
		v383 = v383 + v390
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	goto L130
L142:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v440 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L143:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v411&int32(1) == int32(0) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v437 == v438 {
		v611 = v162
		goto L57
	} else {
		goto L157
	}
L145:
	;
	v417 = int32(2)
	if v411&v417 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	v437 = int32(3)
	goto L144
L147:
	;
	if v411&int32(262144) != 0 {
		v434 = v417
		goto L150
	} else {
		goto L151
	}
L148:
	;
	if v411&int32(4) != 0 {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v437 = int32(1)
	goto L144
L150:
	;
	v437 = v434
	goto L144
L151:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v427 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v431 = F_isImportSlotMigrationJob(m, v427)
	mBase = m.M
	if v431 != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v437 = int32(0)
	goto L144
L154:
	;
	v432 = int32(4)
	goto L156
L155:
	;
	v432 = int32(5)
	goto L156
L156:
	;
	v434 = v432
	goto L150
L157:
	;
	goto L142
L158:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v466 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L159:
	;
	v443 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	if base.Ui64(v443+int64(-32768)) < base.Ui64(int64(-65536)) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	if v465 != 0 {
		v611 = v162
		goto L57
	} else {
		goto L169
	}
L161:
	;
	goto L160
L162:
	;
	v454 = int32(4)
	goto L164
L163:
	;
	v454 = int32(2)
	goto L164
L164:
	;
	if base.Ui64(v443+int64(-2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v459 = int32(8)
	goto L167
L166:
	;
	v459 = v454
	goto L167
L167:
	;
	if base.Ui32(v446) < base.Ui32(v459) {
		v465 = int32(0)
		goto L161
	} else {
		goto L168
	}
L168:
	;
	v461 = int32(0)
	v462 = F_intsetSearch(m, v440, v443, v461)
	mBase = m.M
	v465 = base.B2i32(v462 != v461)
	goto L161
L169:
	;
	goto L158
L170:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v471 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v469 == v466 {
		v611 = v162
		goto L57
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v475 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	v474 = F_clientMatchesFlagFilter(m, l0, v471)
	mBase = m.M
	if v474 != 0 {
		v611 = v162
		goto L57
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v521 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L177:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	if v478 == int32(0) {
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v481 = F_objectGetVal(m, v478)
	mBase = m.M
	if v481 == int32(0) {
		goto L176
	} else {
		goto L179
	}
L179:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v485 = F_objectGetVal(m, v484)
	mBase = m.M
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486))))
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	if v490 == int32(0) {
		v513 = v489
		v514 = v490
		goto L181
	} else {
		goto L182
	}
L180:
	;
	if v514-v513&int32(255) == int32(0) {
		v611 = v162
		goto L57
	} else {
		goto L188
	}
L181:
	;
	goto L180
L182:
	;
	if v490 != v489&int32(255) {
		v513 = v489
		v514 = v490
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v496 = v485
	v497 = v486
	goto L184
L184:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+1)))
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)))
	if v501 == int32(0) {
		v513 = v500
		v514 = v501
		goto L181
	} else {
		goto L186
	}
L185:
	;
	v513 = v500
	v514 = v501
	goto L181
L186:
	;
	v504 = int32(1)
	if v501 == v500&int32(255) {
		v496 = v496 + v504
		v497 = v497 + v504
		goto L184
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	goto L176
L189:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	if v532 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L190:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	if v524 == int32(0) {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v527 = F_compareStringObjects(m, v524, v521)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L3
	} else {
		goto L192
	}
L192:
	;
	if v527 == int32(0) {
		v611 = v162
		goto L57
	} else {
		goto L193
	}
L193:
	;
	goto L189
L194:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v543 == int32(-1) {
		goto L199
	} else {
		goto L200
	}
L195:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v535 == int32(0) {
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v538 = F_compareStringObjects(m, v535, v532)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L3
	} else {
		goto L197
	}
L197:
	;
	if v538 == int32(0) {
		v611 = v162
		goto L57
	} else {
		goto L198
	}
L198:
	;
	goto L194
L199:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v549 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)+28))
	if v547 == v543 {
		v611 = v162
		goto L57
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if v603 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L203:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549+int32(-1)))))
	v582 = int32(0)
	goto L205
L204:
	;
	if base.Ui32(v591) <= base.Ui32(v582) {
		v611 = v162
		goto L57
	} else {
		goto L217
	}
L205:
	;
	switch v570 & int32(7) {
	case 0:
		goto L212
	case 1:
		goto L211
	case 2:
		goto L210
	case 3:
		goto L209
	case 4:
		goto L208
	default:
		v591 = int32(0)
		goto L207
	}
L206:
	;
	goto L204
L207:
	;
	v592 = base.B2i32(base.Ui32(v591) <= base.Ui32(v582))
	if base.Ui32(v591) <= base.Ui32(v582) {
		goto L213
	} else {
		goto L214
	}
L208:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v549+int32(-17))))
	v591 = v590
	goto L207
L209:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v549+int32(-9))))
	v591 = v589
	goto L207
L210:
	;
	v588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v549+int32(-5)))))
	v591 = v588
	goto L207
L211:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549+int32(-3)))))
	v591 = v587
	goto L207
L212:
	;
	v591 = int32(base.Ui32(v570) >> (uint(int32(3)) % 32))
	goto L207
L213:
	;
	goto L206
L214:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549+v582))))
	if v594 != int32(114) {
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v597 = int32(1)
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+227)))
	if v599&v597 != 0 {
		v582 = v582 + v597
		goto L205
	} else {
		goto L216
	}
L216:
	;
	goto L213
L217:
	;
	goto L202
L218:
	;
	v611 = int32(1)
	goto L57
L219:
	;
	v606 = F_clientMatchesIpFilter(m, l0, v603)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L3
	} else {
		goto L220
	}
L220:
	;
	if v606 != 0 {
		v611 = v162
		goto L57
	} else {
		goto L221
	}
L221:
	;
	goto L218
}
func F_clientReplyCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v8 = F_objectGetVal(m, v7)
	mBase = m.M
	v9 = int32(_a25)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v58 = F_objectGetVal(m, v57)
	mBase = m.M
	v59 = int32(_a26)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	if v44-v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L3:
	;
	v44 = F_tolower(m, v40)
	mBase = m.M
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v46 = F_tolower(m, v45)
	mBase = m.M
	goto L2
L4:
	;
	v14 = v8
	v15 = v9
	v16 = v12
	goto L7
L5:
	;
	v40 = int32(0)
	v41 = v9
	goto L3
L6:
	;
	v40 = v37 & int32(255)
	v41 = v36
	goto L3
L7:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v18 == int32(0) {
		v36 = v15
		v37 = v16
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v36 = v30
	v37 = int32(0)
	goto L6
L9:
	;
	v22 = v16 & int32(255)
	if v22 == v18 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v29 = int32(1)
	v30 = v15 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v31 != 0 {
		v14 = v14 + v29
		v15 = v30
		v16 = v31
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v24 = F_tolower(m, v22)
	mBase = m.M
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v26 = F_tolower(m, v25)
	mBase = m.M
	if v24 == v26 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v36 = v15
	v37 = v28
	goto L6
L13:
	;
	goto L8
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v48 & int32(-167772161)
	v53 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	return
L17:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	v104 = F_objectGetVal(m, v103)
	mBase = m.M
	v105 = int32(_a809)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v108 != 0 {
		goto L35
	} else {
		goto L36
	}
L18:
	;
	if v94-v96 != 0 {
		goto L17
	} else {
		goto L30
	}
L19:
	;
	v94 = F_tolower(m, v90)
	mBase = m.M
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v96 = F_tolower(m, v95)
	mBase = m.M
	goto L18
L20:
	;
	v64 = v58
	v65 = v59
	v66 = v62
	goto L23
L21:
	;
	v90 = int32(0)
	v91 = v59
	goto L19
L22:
	;
	v90 = v87 & int32(255)
	v91 = v86
	goto L19
L23:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v68 == int32(0) {
		v86 = v65
		v87 = v66
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v86 = v80
	v87 = int32(0)
	goto L22
L25:
	;
	v72 = v66 & int32(255)
	if v72 == v68 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v79 = int32(1)
	v80 = v65 + v79
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v81 != 0 {
		v64 = v64 + v79
		v65 = v80
		v66 = v81
		goto L23
	} else {
		goto L29
	}
L27:
	;
	v74 = F_tolower(m, v72)
	mBase = m.M
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v76 = F_tolower(m, v75)
	mBase = m.M
	if v74 == v76 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v86 = v65
	v87 = v78
	goto L22
L29:
	;
	goto L24
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v98 | int32(33554432)
	return
L31:
	;
	return
L32:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	F_addReply(m, l0, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L15
	} else {
		goto L47
	}
L33:
	;
	if v140-v142 != 0 {
		goto L32
	} else {
		goto L45
	}
L34:
	;
	v140 = F_tolower(m, v136)
	mBase = m.M
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	v142 = F_tolower(m, v141)
	mBase = m.M
	goto L33
L35:
	;
	v110 = v104
	v111 = v105
	v112 = v108
	goto L38
L36:
	;
	v136 = int32(0)
	v137 = v105
	goto L34
L37:
	;
	v136 = v133 & int32(255)
	v137 = v132
	goto L34
L38:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v114 == int32(0) {
		v132 = v111
		v133 = v112
		goto L37
	} else {
		goto L40
	}
L39:
	;
	v132 = v126
	v133 = int32(0)
	goto L37
L40:
	;
	v118 = v112 & int32(255)
	if v118 == v114 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v125 = int32(1)
	v126 = v111 + v125
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	if v127 != 0 {
		v110 = v110 + v125
		v111 = v126
		v112 = v127
		goto L38
	} else {
		goto L44
	}
L42:
	;
	v120 = F_tolower(m, v118)
	mBase = m.M
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v122 = F_tolower(m, v121)
	mBase = m.M
	if v120 == v122 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v132 = v111
	v133 = v124
	goto L37
L44:
	;
	goto L39
L45:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v144&int32(33554432) != 0 {
		goto L31
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v144 | int32(67108864)
	return
L47:
	;
	v154 = F_objectGetVal(m, v151)
	mBase = m.M
	v156 = F_objectGetVal(m, v151)
	mBase = m.M
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+int32(-1)))))
	switch v159 & int32(7) {
	case 0:
		goto L53
	case 1:
		goto L52
	case 2:
		goto L51
	case 3:
		goto L50
	case 4:
		goto L49
	default:
		v176 = int32(0)
		goto L48
	}
L48:
	;
	F_afterErrorReply(m, l0, v154, v176+int32(-2), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L15
	} else {
		goto L54
	}
L49:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v156+int32(-17))))
	v176 = v175
	goto L48
L50:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v156+int32(-9))))
	v176 = v172
	goto L48
L51:
	;
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156+int32(-5)))))
	v176 = v169
	goto L48
L52:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+int32(-3)))))
	v176 = v166
	goto L48
L53:
	;
	v176 = int32(base.Ui32(v159) >> (uint(int32(3)) % 32))
	goto L48
L54:
	;
	goto L31
}
func F_clientShardSubscriptionsCount(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
	return v4 + v5
}
func F_clientUnpauseCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v2 = int32(_a69)
	*(*int32)(unsafe.Add(mBase, _consts[432])) = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[433])) = int64(0)
	F_updatePausedActions(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[77]))
		F_addReply(m, l0, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	}
}
func F_freeClient(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v263 int64
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v312 int64
	_ = v312
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v519 int32
	_ = v519
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v11 < v2 {
		v27 = v11
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F__serverAssert(m, int32(_a782), int32(_a774), int32(2241))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L17
	} else {
		goto L160
	}
L2:
	;
	F__serverAssert(m, int32(_a133), int32(_a774), int32(2167))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L17
	} else {
		goto L159
	}
L3:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a778), int32(_a774), int32(2277))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L17
	} else {
		goto L158
	}
L4:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a778), int32(_a774), int32(2277))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L17
	} else {
		goto L157
	}
L5:
	;
	m.G0 = v8 + int32(32)
	return v507
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[416]))
	if v48 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L7:
	;
	if v27&int32(1280) != 0 {
		v507 = v2
		goto L5
	} else {
		goto L14
	}
L8:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+207)))
	if v14&int32(2) != 0 {
		v27 = v11
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v19 != 0 {
		v23 = int32(1)
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v23 == int32(0) {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	v23 = base.B2i32(v20 != int32(0))
	goto L11
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v27 = v26
	goto L7
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v27 | int32(1024)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v34 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v45 = F_listAddNodeTail(m, v44, l0)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L17
	} else {
		goto L20
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v39 = F_listSearchKey(m, v38, l0)
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
	if v39 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v507 = v2
	goto L5
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v78 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v51&int32(1) == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v56 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[416])) = v56
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v60&int32(1280) != 0 {
		v507 = v56
		goto L5
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v60 | int32(1024)
	v67 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v67 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v76 = F_listAddNodeTail(m, v75, l0)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L17
	} else {
		goto L29
	}
L26:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v72 = F_listSearchKey(m, v71, l0)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	if v72 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v507 = v56
	goto L5
L30:
	;
	F_moduleNotifyUserChanged(m, l0)
	mBase = m.M
	F_freeClientModuleData(m, l0)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L17
	} else {
		goto L33
	}
L31:
	;
	F_moduleFireServerEvent(m, int64(4), int32(1), l0)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L17
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v88&int32(1024) == int32(0) {
		v104 = v88
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	if v106 == int32(0) {
		v135 = v104
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v95 = F_listSearchKey(m, v94, l0)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L17
	} else {
		goto L36
	}
L36:
	;
	if v95 == int32(0) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	F_listDelNode(m, v100, v95)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L17
	} else {
		goto L38
	}
L38:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v104 = v103
	goto L34
L39:
	;
	if v135&int32(1) != 0 {
		goto L47
	} else {
		goto L48
	}
L40:
	;
	if v104&int32(1) == int32(0) {
		v135 = v104
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v114 {
		v123 = v104
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v124&int32(134218752)|v123&int32(16) != 0 {
		v135 = v123
		goto L39
	} else {
		goto L45
	}
L43:
	;
	F__serverLog(m, int32(2), int32(_a783), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L17
	} else {
		goto L44
	}
L44:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v123 = v122
	goto L42
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v123 & int32(-1105)
	F_replicationCachePrimary(m, l0)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L17
	} else {
		goto L46
	}
L46:
	;
	v507 = v2
	goto L5
L47:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v182 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L48:
	;
	if v135&int32(2) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+207)))
	if v156&int32(4) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	if v135&int32(262144) != 0 {
		goto L47
	} else {
		goto L53
	}
L51:
	;
	if v135&int32(4) == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v148 == int32(0) {
		goto L47
	} else {
		goto L54
	}
L54:
	;
	goto L55
L55:
	;
	goto L47
L56:
	;
	if int32(2) < v155 {
		goto L47
	} else {
		goto L61
	}
L57:
	;
	if int32(2) < v155 {
		goto L47
	} else {
		goto L58
	}
L58:
	;
	v163 = F_replicationGetReplicaName(m, l0)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L17
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v163
	F__serverLog(m, int32(2), int32(_a784), v8+int32(16))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L17
	} else {
		goto L60
	}
L60:
	;
	goto L47
L61:
	;
	v174 = F_replicationGetReplicaName(m, l0)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L17
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v174
	F__serverLog(m, int32(2), int32(_a785), v8)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L17
	} else {
		goto L63
	}
L63:
	;
	goto L47
L64:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v187 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	F_clusterHandleSlotMigrationClientClose(m, v182)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L17
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v221 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v221
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v221
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v225&int32(16) == v221 {
		goto L79
	} else {
		goto L80
	}
L68:
	;
	F_sdsfree(m, v187)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L17
	} else {
		goto L78
	}
L69:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _consts[407]))
	if v187 != v191 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v195 = v187 + int32(-1)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	switch v196 & int32(7) {
	case 0:
		goto L77
	case 1:
		goto L76
	case 2:
		goto L75
	case 3:
		goto L74
	case 4:
		goto L73
	default:
		goto L72
	}
L71:
	;
	goto L67
L72:
	;
	v217 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v187))) = uint8(v217)
	goto L71
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v187+int32(-17)))) = int64(0)
	goto L72
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187+int32(-9)))) = int32(0)
	goto L72
L75:
	;
	v207 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v187+int32(-5)))) = uint16(v207)
	goto L72
L76:
	;
	v203 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v187+int32(-3)))) = uint8(v203)
	goto L72
L77:
	;
	v199 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v199)
	goto L72
L78:
	;
	goto L67
L79:
	;
	F_freeClientBlockingState(m, l0)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L17
	} else {
		goto L82
	}
L80:
	;
	F_unblockClient(m, l0, int32(1))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L17
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	F_freeClientPubSubData(m, l0)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L17
	} else {
		goto L83
	}
L83:
	;
	F_releaseReplyReferences(m, l0)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L17
	} else {
		goto L84
	}
L84:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_listRelease(m, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L17
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	F_zfree_with_size(m, v244, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L17
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(0)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
	F_listRelease(m, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L17
	} else {
		goto L87
	}
L87:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v254&int32(1) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	v376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	if base.Ui32(v376) <= base.Ui32(v375) {
		goto L117
	} else {
		goto L118
	}
L89:
	;
	v308 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v308
	v312 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = v312
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v312
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v318 == v308 {
		goto L88
	} else {
		goto L104
	}
L90:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v275 < int32(1) {
		goto L97
	} else {
		goto L98
	}
L91:
	;
	if v253 != 0 {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	if v253 != 0 {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v259 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v259
	v263 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = v263
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v263
	goto L88
L94:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v271 = F_tryOffloadFreeArgvToIOThreads(m, l0, v269, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L17
	} else {
		goto L95
	}
L95:
	;
	if v271 != int32(-1) {
		goto L89
	} else {
		goto L96
	}
L96:
	;
	goto L90
L97:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_valkey_free(m, v300)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L17
	} else {
		goto L103
	}
L98:
	;
	v282 = int32(0)
	goto L99
L99:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v284+v282<<(uint(int32(2))%32))))
	F_decrRefCount(m, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L17
	} else {
		goto L101
	}
L100:
	;
	goto L97
L101:
	;
	v292 = v282 + int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v292 < v293 {
		v282 = v292
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	goto L89
L104:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v321&int32(1) == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v329 = F_tryOffloadFreeArgvToIOThreads(m, l0, v328, v318)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L17
	} else {
		goto L108
	}
L106:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
	goto L88
L107:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
	goto L88
L108:
	;
	if v329 != int32(-1) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v333 = int32(0)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v334 <= v333 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	F_valkey_free(m, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L17
	} else {
		goto L116
	}
L111:
	;
	v340 = v333
	goto L112
L112:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v342+v340<<(uint(int32(2))%32))))
	F_decrRefCount(m, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L17
	} else {
		goto L114
	}
L113:
	;
	goto L110
L114:
	;
	v350 = v340 + int32(1)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v350 < v351 {
		v340 = v350
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	goto L107
L117:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_valkey_free(m, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L17
	} else {
		goto L129
	}
L118:
	;
	v381 = v375
	goto L119
L119:
	;
	v383 = int32(1)
	v384 = v381 + v383
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)) = uint16(v384)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v391 = v386 + v381&int32(65535)*int32(40)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v392 < v383 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	goto L117
L121:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v391)+8))
	F_valkey_free(m, v417)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L17
	} else {
		goto L127
	}
L122:
	;
	v399 = int32(0)
	goto L123
L123:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v391)+8))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v401+v399<<(uint(int32(2))%32))))
	F_decrRefCount(m, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L17
	} else {
		goto L125
	}
L124:
	;
	goto L121
L125:
	;
	v409 = v399 + int32(1)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v409 < v410 {
		v399 = v409
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	v420 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	if base.Ui32(v420) < base.Ui32(v421) {
		v381 = v420
		goto L119
	} else {
		goto L128
	}
L128:
	;
	goto L120
L129:
	;
	v433 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(64)))) = uint16(v433)
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(56)))) = int64(0)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	if v437 == v433 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v442 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v442
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
	if v444 == v442 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	F_listRelease(m, v437)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L17
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	F_unlinkClient(m, l0)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L17
	} else {
		goto L139
	}
L134:
	;
	v448 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v448 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+226)))
	v458 = v456 << (uint(int32(2)) % 32)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v458)+uint32(_consts[417])))
	*(*int32)(unsafe.Add(mBase, uint32(v458)+uint32(_consts[417]))) = v460 - v444
	goto L133
L136:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v451 != 0 {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v453 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	if l0 != v453 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	goto L135
L139:
	;
	F_freeClientReplicationData(m, l0)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L17
	} else {
		goto L140
	}
L140:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	if v468 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	if v479 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+4)) = v471 - v472
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
	F_listDelNode(m, v475, v476)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L17
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	if v484 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	F_decrRefCount(m, v479)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L17
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v489 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	F_decrRefCount(m, v484)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L17
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	F_freeClientMultiState(m, l0)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L17
	} else {
		goto L153
	}
L151:
	;
	F_decrRefCount(m, v489)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L17
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
	F_sdsfree(m, v496)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L17
	} else {
		goto L154
	}
L154:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	F_sdsfree(m, v499)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L17
	} else {
		goto L155
	}
L155:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L17
	} else {
		goto L156
	}
L156:
	;
	v507 = int32(1)
	goto L5
L157:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
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
L160:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_freeClientAsync(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v3&int32(1280) != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v3 | int32(1024)
		v10 = *(*int32)(unsafe.Add(mBase, _consts[122]))
		if v10 == int32(0) {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[413]))
			v19 = F_listAddNodeTail(m, v18, l0)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				return
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _consts[413]))
			v15 = F_listSearchKey(m, v14, l0)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				if v15 != 0 {
					F__serverAssertWithInfo(m, l0, int32(0), int32(_a778), int32(_a774), int32(2277))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, _consts[413]))
					v19 = F_listAddNodeTail(m, v18, l0)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_freeClientBlockingState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v3 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
		F_dictRelease(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
			F_valkey_free(m, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(0)
				return
			}
		}
	}
}
func F_freeClientOrCloseLater(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	v4 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	if l0 != v4 {
		if l1 == int32(0) {
			v30 = F_freeClient(m, l0)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				return
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			if v12&int32(1280) != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v12 | int32(1024)
				v19 = *(*int32)(unsafe.Add(mBase, _consts[122]))
				if v19 == int32(0) {
					v27 = *(*int32)(unsafe.Add(mBase, _consts[413]))
					v28 = F_listAddNodeTail(m, v27, l0)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						return
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, _consts[413]))
					v24 = F_listSearchKey(m, v23, l0)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						if v24 != 0 {
							F__serverAssertWithInfo(m, l0, int32(0), int32(_a778), int32(_a774), int32(2277))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, _consts[413]))
							v28 = F_listAddNodeTail(m, v27, l0)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v6 | int32(2048)
		return
	}
}
func F_getClientMemoryUsage(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
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
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v6&int32(1) != 0 {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
		v29 = v25*int32(28) + v28
		v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
		if v30 == int64(-1) {
			v40 = v29
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
			v40 = v29 + base.I32_wrap_i64(v30) + v36*int32(28)
		}
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
		v69 = v41 + v40
	} else {
		if v6&int32(2) == int32(0) {
			if v6&int32(262144) != 0 {
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				if v19 == int32(0) {
				} else {
					v22 = F_isImportSlotMigrationJob(m, v19)
					mBase = m.M
				}
			}
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
			v29 = v25*int32(28) + v28
			v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
			if v30 == int64(-1) {
				v40 = v29
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
				v40 = v29 + base.I32_wrap_i64(v30) + v36*int32(28)
			}
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
			v69 = v41 + v40
		} else {
			if v6&int32(4) == int32(0) {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+184))
				if v44 != 0 {
					v47 = *(*int32)(unsafe.Add(mBase, _consts[314]))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
					v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)+16))
					v51 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v49)+24)))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
					v54 = *(*int64)(unsafe.Add(mBase, uint32(v53)+16))
					v57 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
					v58 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
					v61 = int32(44)
					v69 = base.I32_wrap_i64(v50+v51-v54) + base.I32_wrap_i64(v57-v58)*v61 + v61
				} else {
					v69 = int32(0)
				}
			} else {
				if v6&int32(262144) != 0 {
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					if v19 == int32(0) {
					} else {
						v22 = F_isImportSlotMigrationJob(m, v19)
						mBase = m.M
					}
				}
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
				v29 = v25*int32(28) + v28
				v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
				if v30 == int64(-1) {
					v40 = v29
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
					v40 = v29 + base.I32_wrap_i64(v30) + v36*int32(28)
				}
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
				v69 = v41 + v40
			}
		}
	}
	if l1 == int32(0) {
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v69
	}
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v73 != 0 {
		v81 = v73 + int32(-1)
		v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
		v84 = v82 & int32(7)
		switch v84 {
		case 0:
			v85 = F_zmalloc_usable_size(m, v81)
			mBase = m.M
			v114 = v85
		case 1:
			v90 = int32(4)
			switch v84 {
			case 0:
				v114 = v90 + int32(base.Ui32(v82)>>(uint(int32(3))%32))
			case 1:
				v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+int32(-2)))))
				v114 = v90 + v97
			case 2:
				v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73+int32(-3)))))
				v114 = v90 + v101
			case 3:
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v73+int32(-5))))
				v114 = v90 + v105
			case 4:
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v73+int32(-9))))
				v110 = v109
				v114 = v90 + v110
			default:
				v110 = int32(0)
				v114 = v90 + v110
			}
		case 2:
			v90 = int32(6)
			switch v84 {
			case 0:
				v114 = v90 + int32(base.Ui32(v82)>>(uint(int32(3))%32))
			case 1:
				v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+int32(-2)))))
				v114 = v90 + v97
			case 2:
				v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73+int32(-3)))))
				v114 = v90 + v101
			case 3:
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v73+int32(-5))))
				v114 = v90 + v105
			case 4:
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v73+int32(-9))))
				v110 = v109
				v114 = v90 + v110
			default:
				v110 = int32(0)
				v114 = v90 + v110
			}
		case 3:
			v90 = int32(10)
			switch v84 {
			case 0:
				v114 = v90 + int32(base.Ui32(v82)>>(uint(int32(3))%32))
			case 1:
				v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+int32(-2)))))
				v114 = v90 + v97
			case 2:
				v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73+int32(-3)))))
				v114 = v90 + v101
			case 3:
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v73+int32(-5))))
				v114 = v90 + v105
			case 4:
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v73+int32(-9))))
				v110 = v109
				v114 = v90 + v110
			default:
				v110 = int32(0)
				v114 = v90 + v110
			}
		case 4:
			v90 = int32(18)
			switch v84 {
			case 0:
				v114 = v90 + int32(base.Ui32(v82)>>(uint(int32(3))%32))
			case 1:
				v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+int32(-2)))))
				v114 = v90 + v97
			case 2:
				v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73+int32(-3)))))
				v114 = v90 + v101
			case 3:
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v73+int32(-5))))
				v114 = v90 + v105
			case 4:
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v73+int32(-9))))
				v110 = v109
				v114 = v90 + v110
			default:
				v110 = int32(0)
				v114 = v90 + v110
			}
		default:
			v90 = int32(1)
			switch v84 {
			case 0:
				v114 = v90 + int32(base.Ui32(v82)>>(uint(int32(3))%32))
			case 1:
				v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+int32(-2)))))
				v114 = v90 + v97
			case 2:
				v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73+int32(-3)))))
				v114 = v90 + v101
			case 3:
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v73+int32(-5))))
				v114 = v90 + v105
			case 4:
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v73+int32(-9))))
				v110 = v109
				v114 = v90 + v110
			default:
				v110 = int32(0)
				v114 = v90 + v110
			}
		}
		v115 = v114
	} else {
		v115 = int32(0)
	}
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-8))))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v133 = F_multiStateMemOverhead(m, l0)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		return int32(0)
	} else {
		v138 = F_pubsubMemOverhead(m, l0)
		mBase = m.M
		v139 = m.ExcPending
		if v139 != 0 {
			return int32(0)
		} else {
			v140 = v115 + v69 + (v119&int32(2147483647) + int32(8)) + v125 + v127 + v129<<(uint(int32(2))%32) + v133 + v138
			v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			if v141 == int32(0) {
				v152 = v140
			} else {
				v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+24))
				if v144 == int32(0) {
					v152 = v140
				} else {
					v147 = *(*int32)(unsafe.Add(mBase, uint32(v144)+16))
					v152 = v147<<(uint(int32(3))%32) + v140
				}
			}
			return v152
		}
	}
}
func F_getClientPortFromClusterMsg(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v9 = *(*int32)(unsafe.Add(mBase, _consts[89]))
	if v9 != 0 {
		v10 = int32(2246)
	} else {
		v10 = int32(10)
	}
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v10))))
	if v9 != 0 {
		v15 = int32(10)
	} else {
		v15 = int32(2246)
	}
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v15))))
	v18 = F___bswap_16_2(m, v17)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v18
	v20 = F___bswap_16_2(m, v12)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v20
	return
}
func F_getClientSockname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v5 = m.G0
	v7 = v5 - int32(288)
	m.G0 = v7
	v14 = F__emscripten_memset_bulkmem(m, v7+int32(16), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	if v15 != 0 {
		v65 = v15
		m.G0 = v7 + int32(288)
		return v65
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v16 == int32(0) {
			v62 = F_sdsnew(m, v7+int32(16))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = v62
				v65 = v62
				m.G0 = v7 + int32(288)
				return v65
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
			if v20 == int32(0) {
				v62 = F_sdsnew(m, v7+int32(16))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = v62
					v65 = v62
					m.G0 = v7 + int32(288)
					return v65
				}
			} else {
				v29 = m.T0[v20].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v16, v7+int32(160), int32(128), v7+int32(156), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v29 < int32(0) {
						v62 = F_sdsnew(m, v7+int32(16))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = v62
							v65 = v62
							m.G0 = v7 + int32(288)
							return v65
						}
					} else {
						v37 = int32(58)
						v38 = F___strchrnul(m, v7+int32(160), v37)
						mBase = m.M
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
						if v40 == v37 {
							v44 = v38
						} else {
							v44 = int32(0)
						}
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v7)+156))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v45
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(160)
						if v44 != 0 {
							v55 = int32(_a798)
						} else {
							v55 = int32(_a799)
						}
						v56 = F_snprintf(m, v7+int32(16), int32(128), v55, v7)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v62 = F_sdsnew(m, v7+int32(16))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = v62
								v65 = v62
								m.G0 = v7 + int32(288)
								return v65
							}
						}
					}
				}
			}
		}
	}
}
func F_getClientType_3(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v4&int32(1) == int32(0) {
		v11 = int32(2)
		if v4&v11 == int32(0) {
			if v4&int32(262144) != 0 {
				v32 = v11
				return v32
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				if v22 != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					if v27 == int32(1) {
						v30 = int32(4)
					} else {
						v30 = int32(5)
					}
					v32 = v30
					return v32
				} else {
					return int32(0)
				}
			}
		} else {
			if v4&int32(4) != 0 {
				if v4&int32(262144) != 0 {
					v32 = v11
					return v32
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					if v22 != 0 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
						if v27 == int32(1) {
							v30 = int32(4)
						} else {
							v30 = int32(5)
						}
						v32 = v30
						return v32
					} else {
						return int32(0)
					}
				}
			} else {
				return int32(1)
			}
		}
	} else {
		return int32(3)
	}
}
func F_installClientWriteHandler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = int32(_a69)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	v7 = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[415]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	v16 = m.T0[v15].(func(*base.Module, int32, int32, int32) int32)(m, v3, int32(954), base.B2i32(v6 == v7)&base.B2i32(v10 == v7))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if v16 != int32(-1) {
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			if v20&int32(1280) != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v20 | int32(1024)
				v27 = *(*int32)(unsafe.Add(mBase, _consts[122]))
				if v27 == int32(0) {
					v35 = *(*int32)(unsafe.Add(mBase, _consts[413]))
					v36 = F_listAddNodeTail(m, v35, l0)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						return
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _consts[413]))
					v32 = F_listSearchKey(m, v31, l0)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						if v32 != 0 {
							F__serverAssertWithInfo(m, l0, int32(0), int32(_a778), int32(_a774), int32(2277))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, _consts[413]))
							v36 = F_listAddNodeTail(m, v35, l0)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
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
}
func F_isClientConnIpV6(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	if v6 != 0 {
		v7 = v6
	} else {
		v7 = l0
	}
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+207)))
	if int32(base.Ui32(v8&int32(16))>>(uint(int32(4))%32)) != 0 {
		v13 = v7
	} else {
		v13 = l0
	}
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+207)))
	if v14&int32(16) != 0 {
		v27 = v2
		return v27
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		if v17 == int32(0) {
			v27 = v2
			return v27
		} else {
			v20 = F_getClientPeerId(m, v13)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v27 = base.B2i32(v24 == int32(91))
				return v27
			}
		}
	}
}
func F_prepareClientToWrite(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v4&int32(1073742080) != 0 {
		v101 = int32(0)
		return v101
	} else {
		v7 = int32(-1)
		if v4&int32(1024) != 0 {
			v101 = v7
			return v101
		} else {
			if v4&int32(167772160) == int32(0) {
				if v4&int32(8193) == int32(1) {
					v101 = v7
					return v101
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
					v26 = v25
					if v26&int32(268435456) == int32(0) {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v34 == int32(0) {
							F__serverAssert(m, int32(_a781), int32(_a774), int32(494))
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v37 = F_clientHasPendingReplies(m, l0)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								if v37 != 0 {
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
									if v41&int32(4194304) != 0 {
									} else {
										v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
										if v44 == int32(0) {
											v52 = int32(1)
											v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
											if v53 == int32(0) {
												v60 = v52
											} else {
												v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
												if v56 != 0 {
													v60 = v52
												} else {
													v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+156))
													v60 = base.B2i32(v57 != int32(13))
												}
											}
											if v60 == int32(0) {
											} else {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v63 | int32(4194304)
												v68 = *(*int32)(unsafe.Add(mBase, _consts[317]))
												v70 = l0 + int32(168)
												v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
												if v73 != 0 {
													*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(0)
													v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
													*(*int32)(unsafe.Add(mBase, uint32(v80))) = v70
													v82 = v80
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v70
													v75 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
													v82 = v75
												}
												*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
												*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v82
												*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v73 + int32(1)
											}
										} else {
											v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
											switch v47 {
											case 0:
												v52 = int32(1)
												v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
												if v53 == int32(0) {
													v60 = v52
												} else {
													v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
													if v56 != 0 {
														v60 = v52
													} else {
														v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+156))
														v60 = base.B2i32(v57 != int32(13))
													}
												}
												if v60 == int32(0) {
												} else {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v63 | int32(4194304)
													v68 = *(*int32)(unsafe.Add(mBase, _consts[317]))
													v70 = l0 + int32(168)
													v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
													if v73 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(0)
														v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
														*(*int32)(unsafe.Add(mBase, uint32(v80))) = v70
														v82 = v80
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v70
														v75 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
														v82 = v75
													}
													*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
													*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v82
													*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v73 + int32(1)
												}
											default:
											case 9, 11:
												if v41&int32(1024) != 0 {
												} else {
													v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
													if v50 != 0 {
													} else {
														v52 = int32(1)
														v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
														if v53 == int32(0) {
															v60 = v52
														} else {
															v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
															if v56 != 0 {
																v60 = v52
															} else {
																v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+156))
																v60 = base.B2i32(v57 != int32(13))
															}
														}
														if v60 == int32(0) {
														} else {
															v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v63 | int32(4194304)
															v68 = *(*int32)(unsafe.Add(mBase, _consts[317]))
															v70 = l0 + int32(168)
															v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
															if v73 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(0)
																v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
																*(*int32)(unsafe.Add(mBase, uint32(v80))) = v70
																v82 = v80
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v70
																v75 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
																v82 = v75
															}
															*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
															*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v82
															*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v73 + int32(1)
														}
													}
												}
											}
										}
									}
								}
								v90 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
								if v90 != int64(-1) {
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v93 | int32(1073741824)
								}
								v101 = int32(0)
								return v101
							}
						}
					} else {
						v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
						if v31 != int64(-2) {
							v101 = v7
							return v101
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v34 == int32(0) {
								F__serverAssert(m, int32(_a781), int32(_a774), int32(494))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v37 = F_clientHasPendingReplies(m, l0)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									if v37 != 0 {
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
										if v41&int32(4194304) != 0 {
										} else {
											v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
											if v44 == int32(0) {
												v52 = int32(1)
												v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
												if v53 == int32(0) {
													v60 = v52
												} else {
													v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
													if v56 != 0 {
														v60 = v52
													} else {
														v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+156))
														v60 = base.B2i32(v57 != int32(13))
													}
												}
												if v60 == int32(0) {
												} else {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v63 | int32(4194304)
													v68 = *(*int32)(unsafe.Add(mBase, _consts[317]))
													v70 = l0 + int32(168)
													v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
													if v73 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(0)
														v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
														*(*int32)(unsafe.Add(mBase, uint32(v80))) = v70
														v82 = v80
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v70
														v75 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
														v82 = v75
													}
													*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
													*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v82
													*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v73 + int32(1)
												}
											} else {
												v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
												switch v47 {
												case 0:
													v52 = int32(1)
													v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
													if v53 == int32(0) {
														v60 = v52
													} else {
														v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
														if v56 != 0 {
															v60 = v52
														} else {
															v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+156))
															v60 = base.B2i32(v57 != int32(13))
														}
													}
													if v60 == int32(0) {
													} else {
														v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v63 | int32(4194304)
														v68 = *(*int32)(unsafe.Add(mBase, _consts[317]))
														v70 = l0 + int32(168)
														v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
														if v73 != 0 {
															*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(0)
															v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
															*(*int32)(unsafe.Add(mBase, uint32(v80))) = v70
															v82 = v80
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v70
															v75 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
															v82 = v75
														}
														*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
														*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v82
														*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v73 + int32(1)
													}
												default:
												case 9, 11:
													if v41&int32(1024) != 0 {
													} else {
														v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
														if v50 != 0 {
														} else {
															v52 = int32(1)
															v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
															if v53 == int32(0) {
																v60 = v52
															} else {
																v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
																if v56 != 0 {
																	v60 = v52
																} else {
																	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+156))
																	v60 = base.B2i32(v57 != int32(13))
																}
															}
															if v60 == int32(0) {
															} else {
																v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v63 | int32(4194304)
																v68 = *(*int32)(unsafe.Add(mBase, _consts[317]))
																v70 = l0 + int32(168)
																v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
																if v73 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(0)
																	v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
																	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v70
																	v82 = v80
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v70
																	v75 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
																	v82 = v75
																}
																*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
																*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v82
																*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v73 + int32(1)
															}
														}
													}
												}
											}
										}
									}
									v90 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
									if v90 != int64(-1) {
									} else {
										v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v93 | int32(1073741824)
									}
									v101 = int32(0)
									return v101
								}
							}
						}
					}
				}
			} else {
				if v4&int32(8193) == int32(1) {
					v101 = v7
					return v101
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
					if v18&int32(131072) != 0 {
						v26 = v18
						if v26&int32(268435456) == int32(0) {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v34 == int32(0) {
								F__serverAssert(m, int32(_a781), int32(_a774), int32(494))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v37 = F_clientHasPendingReplies(m, l0)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									if v37 != 0 {
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
										if v41&int32(4194304) != 0 {
										} else {
											v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
											if v44 == int32(0) {
												v52 = int32(1)
												v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
												if v53 == int32(0) {
													v60 = v52
												} else {
													v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
													if v56 != 0 {
														v60 = v52
													} else {
														v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+156))
														v60 = base.B2i32(v57 != int32(13))
													}
												}
												if v60 == int32(0) {
												} else {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v63 | int32(4194304)
													v68 = *(*int32)(unsafe.Add(mBase, _consts[317]))
													v70 = l0 + int32(168)
													v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
													if v73 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(0)
														v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
														*(*int32)(unsafe.Add(mBase, uint32(v80))) = v70
														v82 = v80
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v70
														v75 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
														v82 = v75
													}
													*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
													*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v82
													*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v73 + int32(1)
												}
											} else {
												v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
												switch v47 {
												case 0:
													v52 = int32(1)
													v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
													if v53 == int32(0) {
														v60 = v52
													} else {
														v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
														if v56 != 0 {
															v60 = v52
														} else {
															v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+156))
															v60 = base.B2i32(v57 != int32(13))
														}
													}
													if v60 == int32(0) {
													} else {
														v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v63 | int32(4194304)
														v68 = *(*int32)(unsafe.Add(mBase, _consts[317]))
														v70 = l0 + int32(168)
														v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
														if v73 != 0 {
															*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(0)
															v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
															*(*int32)(unsafe.Add(mBase, uint32(v80))) = v70
															v82 = v80
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v70
															v75 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
															v82 = v75
														}
														*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
														*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v82
														*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v73 + int32(1)
													}
												default:
												case 9, 11:
													if v41&int32(1024) != 0 {
													} else {
														v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
														if v50 != 0 {
														} else {
															v52 = int32(1)
															v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
															if v53 == int32(0) {
																v60 = v52
															} else {
																v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
																if v56 != 0 {
																	v60 = v52
																} else {
																	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+156))
																	v60 = base.B2i32(v57 != int32(13))
																}
															}
															if v60 == int32(0) {
															} else {
																v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v63 | int32(4194304)
																v68 = *(*int32)(unsafe.Add(mBase, _consts[317]))
																v70 = l0 + int32(168)
																v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
																if v73 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(0)
																	v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
																	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v70
																	v82 = v80
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v70
																	v75 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
																	v82 = v75
																}
																*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
																*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v82
																*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v73 + int32(1)
															}
														}
													}
												}
											}
										}
									}
									v90 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
									if v90 != int64(-1) {
									} else {
										v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v93 | int32(1073741824)
									}
									v101 = int32(0)
									return v101
								}
							}
						} else {
							v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							if v31 != int64(-2) {
								v101 = v7
								return v101
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v34 == int32(0) {
									F__serverAssert(m, int32(_a781), int32(_a774), int32(494))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v37 = F_clientHasPendingReplies(m, l0)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										if v37 != 0 {
										} else {
											v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
											if v41&int32(4194304) != 0 {
											} else {
												v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
												if v44 == int32(0) {
													v52 = int32(1)
													v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
													if v53 == int32(0) {
														v60 = v52
													} else {
														v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
														if v56 != 0 {
															v60 = v52
														} else {
															v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+156))
															v60 = base.B2i32(v57 != int32(13))
														}
													}
													if v60 == int32(0) {
													} else {
														v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v63 | int32(4194304)
														v68 = *(*int32)(unsafe.Add(mBase, _consts[317]))
														v70 = l0 + int32(168)
														v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
														if v73 != 0 {
															*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(0)
															v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
															*(*int32)(unsafe.Add(mBase, uint32(v80))) = v70
															v82 = v80
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v70
															v75 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
															v82 = v75
														}
														*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
														*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v82
														*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v73 + int32(1)
													}
												} else {
													v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
													switch v47 {
													case 0:
														v52 = int32(1)
														v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
														if v53 == int32(0) {
															v60 = v52
														} else {
															v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
															if v56 != 0 {
																v60 = v52
															} else {
																v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+156))
																v60 = base.B2i32(v57 != int32(13))
															}
														}
														if v60 == int32(0) {
														} else {
															v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v63 | int32(4194304)
															v68 = *(*int32)(unsafe.Add(mBase, _consts[317]))
															v70 = l0 + int32(168)
															v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
															if v73 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(0)
																v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
																*(*int32)(unsafe.Add(mBase, uint32(v80))) = v70
																v82 = v80
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v70
																v75 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
																v82 = v75
															}
															*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
															*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v82
															*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v73 + int32(1)
														}
													default:
													case 9, 11:
														if v41&int32(1024) != 0 {
														} else {
															v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
															if v50 != 0 {
															} else {
																v52 = int32(1)
																v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
																if v53 == int32(0) {
																	v60 = v52
																} else {
																	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
																	if v56 != 0 {
																		v60 = v52
																	} else {
																		v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+156))
																		v60 = base.B2i32(v57 != int32(13))
																	}
																}
																if v60 == int32(0) {
																} else {
																	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v63 | int32(4194304)
																	v68 = *(*int32)(unsafe.Add(mBase, _consts[317]))
																	v70 = l0 + int32(168)
																	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
																	if v73 != 0 {
																		*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(0)
																		v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
																		*(*int32)(unsafe.Add(mBase, uint32(v80))) = v70
																		v82 = v80
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v70
																		v75 = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
																		v82 = v75
																	}
																	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
																	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v82
																	*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v73 + int32(1)
																}
															}
														}
													}
												}
											}
										}
										v90 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
										if v90 != int64(-1) {
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v93 | int32(1073741824)
										}
										v101 = int32(0)
										return v101
									}
								}
							}
						}
					} else {
						v101 = v7
						return v101
					}
				}
			}
		}
	}
}
func F_removeClientFromPendingCommandsBatch(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[338]))
	if v7 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	v16 = int32(0)
	goto L5
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(0)
	goto L1
L5:
	;
	v22 = v13 + v16<<(uint(int32(2))%32)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v23 == l0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v26 = v16 + int32(1)
	if v26 == v10 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v16 = v26
	goto L5
}
func F_rewriteClientCommandArgument(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if l1 < v6 {
		v10 = v6
	} else {
		v10 = l1 + int32(1)
	}
	F_backupAndUpdateClientArgv(m, l0, v10, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14+l1<<(uint(int32(2))%32))))
		if v18 == int32(0) {
			if l2 == int32(0) {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v34+l1<<(uint(int32(2))%32)))) = l2
				F_incrRefCount(m, l2)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					if v18 == int32(0) {
						if l1 != 0 {
							return
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v45 & int32(-1073741825)
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v51 = F_lookupCommandOrOriginal(m, v49, v50)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v51
								if v51 != 0 {
									return
								} else {
									F__serverAssertWithInfo(m, l0, int32(0), int32(_a814), int32(_a774), int32(6119))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						F_decrRefCount(m, v18)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							if l1 != 0 {
								return
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v45 & int32(-1073741825)
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v51 = F_lookupCommandOrOriginal(m, v49, v50)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v51
									if v51 != 0 {
										return
									} else {
										F__serverAssertWithInfo(m, l0, int32(0), int32(_a814), int32(_a774), int32(6119))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											F_abort(m)
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
				v29 = F_getStringObjectLen(m, l2)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v29 + v31
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v34+l1<<(uint(int32(2))%32)))) = l2
					F_incrRefCount(m, l2)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						if v18 == int32(0) {
							if l1 != 0 {
								return
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v45 & int32(-1073741825)
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v51 = F_lookupCommandOrOriginal(m, v49, v50)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v51
									if v51 != 0 {
										return
									} else {
										F__serverAssertWithInfo(m, l0, int32(0), int32(_a814), int32(_a774), int32(6119))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							F_decrRefCount(m, v18)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								if l1 != 0 {
									return
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v45 & int32(-1073741825)
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v51 = F_lookupCommandOrOriginal(m, v49, v50)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v51
										if v51 != 0 {
											return
										} else {
											F__serverAssertWithInfo(m, l0, int32(0), int32(_a814), int32(_a774), int32(6119))
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												F_abort(m)
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
				}
			}
		} else {
			v21 = F_getStringObjectLen(m, v18)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v23 - v21
				if l2 == int32(0) {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v34+l1<<(uint(int32(2))%32)))) = l2
					F_incrRefCount(m, l2)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						if v18 == int32(0) {
							if l1 != 0 {
								return
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v45 & int32(-1073741825)
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v51 = F_lookupCommandOrOriginal(m, v49, v50)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v51
									if v51 != 0 {
										return
									} else {
										F__serverAssertWithInfo(m, l0, int32(0), int32(_a814), int32(_a774), int32(6119))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							F_decrRefCount(m, v18)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								if l1 != 0 {
									return
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v45 & int32(-1073741825)
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v51 = F_lookupCommandOrOriginal(m, v49, v50)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v51
										if v51 != 0 {
											return
										} else {
											F__serverAssertWithInfo(m, l0, int32(0), int32(_a814), int32(_a774), int32(6119))
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												F_abort(m)
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
					v29 = F_getStringObjectLen(m, l2)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v29 + v31
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v34+l1<<(uint(int32(2))%32)))) = l2
						F_incrRefCount(m, l2)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							if v18 == int32(0) {
								if l1 != 0 {
									return
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v45 & int32(-1073741825)
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v51 = F_lookupCommandOrOriginal(m, v49, v50)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v51
										if v51 != 0 {
											return
										} else {
											F__serverAssertWithInfo(m, l0, int32(0), int32(_a814), int32(_a774), int32(6119))
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								F_decrRefCount(m, v18)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									if l1 != 0 {
										return
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v45 & int32(-1073741825)
										v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v51 = F_lookupCommandOrOriginal(m, v49, v50)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v51
											if v51 != 0 {
												return
											} else {
												F__serverAssertWithInfo(m, l0, int32(0), int32(_a814), int32(_a774), int32(6119))
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return
												} else {
													F_abort(m)
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
					}
				}
			}
		}
	}
}
func F_rewriteClientCommandVector(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_valkey_malloc(m, l1<<(uint(int32(2))%32))
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
	if l1 < int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_replaceClientCommandVector(m, l0, l1, v13)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L9
	}
L4:
	;
	v21 = int32(0)
	goto L5
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v25 + int32(4)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v13+v21<<(uint(int32(2))%32)))) = v32
	F_incrRefCount(m, v32)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L3
L7:
	;
	v37 = v21 + int32(1)
	if v37 != l1 {
		v21 = v37
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	m.G0 = v9 + int32(16)
	return
}
func F_trimClientQueryBuffer(m *base.Module, l0 int32) {
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
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
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v151 int32
	_ = v151
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = *(*int32)(unsafe.Add(mBase, _consts[407]))
	if v6 != v8 {
		v13 = v6
		if v13 == int32(0) {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-1)))))
			switch v20 & int32(7) {
			case 0:
				v37 = int32(base.Ui32(v20) >> (uint(int32(3)) % 32))
			case 1:
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
				v37 = v27
			case 2:
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
				v37 = v30
			case 3:
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
				v37 = v33
			case 4:
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
				v37 = v36
			default:
				v37 = int32(0)
			}
			if base.Ui32(v37) < base.Ui32(v16) {
				F__serverAssert(m, int32(_a800), int32(_a774), int32(2341))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if v16 == int32(0) {
				} else {
					v41 = int32(-1)
					v49 = v13 + v41
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
					v52 = v50 & int32(7)
					switch v52 {
					case 0:
						v67 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
						if v67 == int32(0) {
						} else {
							v73 = int32(-1)&v67 + v41
							v77 = v16>>(uint(int32(31))%32)&v67 + v16
							v80 = v73 - v77 + int32(1)
							switch v52 {
							default:
								v95 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
							case 1:
								v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
								v95 = v85
							case 2:
								v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
								v95 = v88
							case 3:
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
								v95 = v91
							case 4:
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
								v95 = v94
							}
							v96 = int32(0)
							v98 = base.B2i32(base.Ui32(v77) < base.Ui32(v95))
							if base.Ui32(v77) < base.Ui32(v95) {
								v99 = v77
							} else {
								v99 = v96
							}
							v100 = v95 - v99
							if base.Ui32(v80) < base.Ui32(v100) {
								v102 = v80
							} else {
								v102 = v100
							}
							if v73 < v77 {
								v104 = v96
							} else {
								v104 = v102
							}
							if base.Ui32(v77) < base.Ui32(v95) {
								v106 = v104
							} else {
								v106 = int32(0)
							}
							if v106 == int32(0) {
							} else {
								v110 = F_memmove(m, v13, v13+v99, v106)
								mBase = m.M
							}
							v112 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v13+v106))) = uint8(v112)
							switch v52 {
							default:
								v115 = v106 << (uint(int32(3)) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v115)
							case 1:
								*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))) = uint8(v106)
							case 2:
								*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))) = uint16(v106)
							case 3:
								*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9)))) = v106
							case 4:
								*(*int64)(unsafe.Add(mBase, uint32(v13+int32(-17)))) = base.I64_extend_i32_u(v106)
							}
						}
					case 1:
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
						v67 = v57
						if v67 == int32(0) {
						} else {
							v73 = int32(-1)&v67 + v41
							v77 = v16>>(uint(int32(31))%32)&v67 + v16
							v80 = v73 - v77 + int32(1)
							switch v52 {
							default:
								v95 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
							case 1:
								v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
								v95 = v85
							case 2:
								v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
								v95 = v88
							case 3:
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
								v95 = v91
							case 4:
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
								v95 = v94
							}
							v96 = int32(0)
							v98 = base.B2i32(base.Ui32(v77) < base.Ui32(v95))
							if base.Ui32(v77) < base.Ui32(v95) {
								v99 = v77
							} else {
								v99 = v96
							}
							v100 = v95 - v99
							if base.Ui32(v80) < base.Ui32(v100) {
								v102 = v80
							} else {
								v102 = v100
							}
							if v73 < v77 {
								v104 = v96
							} else {
								v104 = v102
							}
							if base.Ui32(v77) < base.Ui32(v95) {
								v106 = v104
							} else {
								v106 = int32(0)
							}
							if v106 == int32(0) {
							} else {
								v110 = F_memmove(m, v13, v13+v99, v106)
								mBase = m.M
							}
							v112 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v13+v106))) = uint8(v112)
							switch v52 {
							default:
								v115 = v106 << (uint(int32(3)) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v115)
							case 1:
								*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))) = uint8(v106)
							case 2:
								*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))) = uint16(v106)
							case 3:
								*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9)))) = v106
							case 4:
								*(*int64)(unsafe.Add(mBase, uint32(v13+int32(-17)))) = base.I64_extend_i32_u(v106)
							}
						}
					case 2:
						v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
						v67 = v60
						if v67 == int32(0) {
						} else {
							v73 = int32(-1)&v67 + v41
							v77 = v16>>(uint(int32(31))%32)&v67 + v16
							v80 = v73 - v77 + int32(1)
							switch v52 {
							default:
								v95 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
							case 1:
								v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
								v95 = v85
							case 2:
								v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
								v95 = v88
							case 3:
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
								v95 = v91
							case 4:
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
								v95 = v94
							}
							v96 = int32(0)
							v98 = base.B2i32(base.Ui32(v77) < base.Ui32(v95))
							if base.Ui32(v77) < base.Ui32(v95) {
								v99 = v77
							} else {
								v99 = v96
							}
							v100 = v95 - v99
							if base.Ui32(v80) < base.Ui32(v100) {
								v102 = v80
							} else {
								v102 = v100
							}
							if v73 < v77 {
								v104 = v96
							} else {
								v104 = v102
							}
							if base.Ui32(v77) < base.Ui32(v95) {
								v106 = v104
							} else {
								v106 = int32(0)
							}
							if v106 == int32(0) {
							} else {
								v110 = F_memmove(m, v13, v13+v99, v106)
								mBase = m.M
							}
							v112 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v13+v106))) = uint8(v112)
							switch v52 {
							default:
								v115 = v106 << (uint(int32(3)) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v115)
							case 1:
								*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))) = uint8(v106)
							case 2:
								*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))) = uint16(v106)
							case 3:
								*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9)))) = v106
							case 4:
								*(*int64)(unsafe.Add(mBase, uint32(v13+int32(-17)))) = base.I64_extend_i32_u(v106)
							}
						}
					case 3:
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
						v67 = v63
						if v67 == int32(0) {
						} else {
							v73 = int32(-1)&v67 + v41
							v77 = v16>>(uint(int32(31))%32)&v67 + v16
							v80 = v73 - v77 + int32(1)
							switch v52 {
							default:
								v95 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
							case 1:
								v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
								v95 = v85
							case 2:
								v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
								v95 = v88
							case 3:
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
								v95 = v91
							case 4:
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
								v95 = v94
							}
							v96 = int32(0)
							v98 = base.B2i32(base.Ui32(v77) < base.Ui32(v95))
							if base.Ui32(v77) < base.Ui32(v95) {
								v99 = v77
							} else {
								v99 = v96
							}
							v100 = v95 - v99
							if base.Ui32(v80) < base.Ui32(v100) {
								v102 = v80
							} else {
								v102 = v100
							}
							if v73 < v77 {
								v104 = v96
							} else {
								v104 = v102
							}
							if base.Ui32(v77) < base.Ui32(v95) {
								v106 = v104
							} else {
								v106 = int32(0)
							}
							if v106 == int32(0) {
							} else {
								v110 = F_memmove(m, v13, v13+v99, v106)
								mBase = m.M
							}
							v112 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v13+v106))) = uint8(v112)
							switch v52 {
							default:
								v115 = v106 << (uint(int32(3)) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v115)
							case 1:
								*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))) = uint8(v106)
							case 2:
								*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))) = uint16(v106)
							case 3:
								*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9)))) = v106
							case 4:
								*(*int64)(unsafe.Add(mBase, uint32(v13+int32(-17)))) = base.I64_extend_i32_u(v106)
							}
						}
					case 4:
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
						v67 = v66
						if v67 == int32(0) {
						} else {
							v73 = int32(-1)&v67 + v41
							v77 = v16>>(uint(int32(31))%32)&v67 + v16
							v80 = v73 - v77 + int32(1)
							switch v52 {
							default:
								v95 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
							case 1:
								v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
								v95 = v85
							case 2:
								v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
								v95 = v88
							case 3:
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
								v95 = v91
							case 4:
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
								v95 = v94
							}
							v96 = int32(0)
							v98 = base.B2i32(base.Ui32(v77) < base.Ui32(v95))
							if base.Ui32(v77) < base.Ui32(v95) {
								v99 = v77
							} else {
								v99 = v96
							}
							v100 = v95 - v99
							if base.Ui32(v80) < base.Ui32(v100) {
								v102 = v80
							} else {
								v102 = v100
							}
							if v73 < v77 {
								v104 = v96
							} else {
								v104 = v102
							}
							if base.Ui32(v77) < base.Ui32(v95) {
								v106 = v104
							} else {
								v106 = int32(0)
							}
							if v106 == int32(0) {
							} else {
								v110 = F_memmove(m, v13, v13+v99, v106)
								mBase = m.M
							}
							v112 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v13+v106))) = uint8(v112)
							switch v52 {
							default:
								v115 = v106 << (uint(int32(3)) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v115)
							case 1:
								*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))) = uint8(v106)
							case 2:
								*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))) = uint16(v106)
							case 3:
								*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9)))) = v106
							case 4:
								*(*int64)(unsafe.Add(mBase, uint32(v13+int32(-17)))) = base.I64_extend_i32_u(v106)
							}
						}
					default:
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
				}
				return
			}
		}
	} else {
		F_resetSharedQueryBuf(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v13 = v12
			if v13 == int32(0) {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-1)))))
				switch v20 & int32(7) {
				case 0:
					v37 = int32(base.Ui32(v20) >> (uint(int32(3)) % 32))
				case 1:
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
					v37 = v27
				case 2:
					v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
					v37 = v30
				case 3:
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
					v37 = v33
				case 4:
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
					v37 = v36
				default:
					v37 = int32(0)
				}
				if base.Ui32(v37) < base.Ui32(v16) {
					F__serverAssert(m, int32(_a800), int32(_a774), int32(2341))
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					if v16 == int32(0) {
					} else {
						v41 = int32(-1)
						v49 = v13 + v41
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
						v52 = v50 & int32(7)
						switch v52 {
						case 0:
							v67 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
							if v67 == int32(0) {
							} else {
								v73 = int32(-1)&v67 + v41
								v77 = v16>>(uint(int32(31))%32)&v67 + v16
								v80 = v73 - v77 + int32(1)
								switch v52 {
								default:
									v95 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
								case 1:
									v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
									v95 = v85
								case 2:
									v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
									v95 = v88
								case 3:
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
									v95 = v91
								case 4:
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
									v95 = v94
								}
								v96 = int32(0)
								v98 = base.B2i32(base.Ui32(v77) < base.Ui32(v95))
								if base.Ui32(v77) < base.Ui32(v95) {
									v99 = v77
								} else {
									v99 = v96
								}
								v100 = v95 - v99
								if base.Ui32(v80) < base.Ui32(v100) {
									v102 = v80
								} else {
									v102 = v100
								}
								if v73 < v77 {
									v104 = v96
								} else {
									v104 = v102
								}
								if base.Ui32(v77) < base.Ui32(v95) {
									v106 = v104
								} else {
									v106 = int32(0)
								}
								if v106 == int32(0) {
								} else {
									v110 = F_memmove(m, v13, v13+v99, v106)
									mBase = m.M
								}
								v112 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v13+v106))) = uint8(v112)
								switch v52 {
								default:
									v115 = v106 << (uint(int32(3)) % 32)
									*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v115)
								case 1:
									*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))) = uint8(v106)
								case 2:
									*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))) = uint16(v106)
								case 3:
									*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9)))) = v106
								case 4:
									*(*int64)(unsafe.Add(mBase, uint32(v13+int32(-17)))) = base.I64_extend_i32_u(v106)
								}
							}
						case 1:
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
							v67 = v57
							if v67 == int32(0) {
							} else {
								v73 = int32(-1)&v67 + v41
								v77 = v16>>(uint(int32(31))%32)&v67 + v16
								v80 = v73 - v77 + int32(1)
								switch v52 {
								default:
									v95 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
								case 1:
									v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
									v95 = v85
								case 2:
									v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
									v95 = v88
								case 3:
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
									v95 = v91
								case 4:
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
									v95 = v94
								}
								v96 = int32(0)
								v98 = base.B2i32(base.Ui32(v77) < base.Ui32(v95))
								if base.Ui32(v77) < base.Ui32(v95) {
									v99 = v77
								} else {
									v99 = v96
								}
								v100 = v95 - v99
								if base.Ui32(v80) < base.Ui32(v100) {
									v102 = v80
								} else {
									v102 = v100
								}
								if v73 < v77 {
									v104 = v96
								} else {
									v104 = v102
								}
								if base.Ui32(v77) < base.Ui32(v95) {
									v106 = v104
								} else {
									v106 = int32(0)
								}
								if v106 == int32(0) {
								} else {
									v110 = F_memmove(m, v13, v13+v99, v106)
									mBase = m.M
								}
								v112 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v13+v106))) = uint8(v112)
								switch v52 {
								default:
									v115 = v106 << (uint(int32(3)) % 32)
									*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v115)
								case 1:
									*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))) = uint8(v106)
								case 2:
									*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))) = uint16(v106)
								case 3:
									*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9)))) = v106
								case 4:
									*(*int64)(unsafe.Add(mBase, uint32(v13+int32(-17)))) = base.I64_extend_i32_u(v106)
								}
							}
						case 2:
							v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
							v67 = v60
							if v67 == int32(0) {
							} else {
								v73 = int32(-1)&v67 + v41
								v77 = v16>>(uint(int32(31))%32)&v67 + v16
								v80 = v73 - v77 + int32(1)
								switch v52 {
								default:
									v95 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
								case 1:
									v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
									v95 = v85
								case 2:
									v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
									v95 = v88
								case 3:
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
									v95 = v91
								case 4:
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
									v95 = v94
								}
								v96 = int32(0)
								v98 = base.B2i32(base.Ui32(v77) < base.Ui32(v95))
								if base.Ui32(v77) < base.Ui32(v95) {
									v99 = v77
								} else {
									v99 = v96
								}
								v100 = v95 - v99
								if base.Ui32(v80) < base.Ui32(v100) {
									v102 = v80
								} else {
									v102 = v100
								}
								if v73 < v77 {
									v104 = v96
								} else {
									v104 = v102
								}
								if base.Ui32(v77) < base.Ui32(v95) {
									v106 = v104
								} else {
									v106 = int32(0)
								}
								if v106 == int32(0) {
								} else {
									v110 = F_memmove(m, v13, v13+v99, v106)
									mBase = m.M
								}
								v112 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v13+v106))) = uint8(v112)
								switch v52 {
								default:
									v115 = v106 << (uint(int32(3)) % 32)
									*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v115)
								case 1:
									*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))) = uint8(v106)
								case 2:
									*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))) = uint16(v106)
								case 3:
									*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9)))) = v106
								case 4:
									*(*int64)(unsafe.Add(mBase, uint32(v13+int32(-17)))) = base.I64_extend_i32_u(v106)
								}
							}
						case 3:
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
							v67 = v63
							if v67 == int32(0) {
							} else {
								v73 = int32(-1)&v67 + v41
								v77 = v16>>(uint(int32(31))%32)&v67 + v16
								v80 = v73 - v77 + int32(1)
								switch v52 {
								default:
									v95 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
								case 1:
									v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
									v95 = v85
								case 2:
									v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
									v95 = v88
								case 3:
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
									v95 = v91
								case 4:
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
									v95 = v94
								}
								v96 = int32(0)
								v98 = base.B2i32(base.Ui32(v77) < base.Ui32(v95))
								if base.Ui32(v77) < base.Ui32(v95) {
									v99 = v77
								} else {
									v99 = v96
								}
								v100 = v95 - v99
								if base.Ui32(v80) < base.Ui32(v100) {
									v102 = v80
								} else {
									v102 = v100
								}
								if v73 < v77 {
									v104 = v96
								} else {
									v104 = v102
								}
								if base.Ui32(v77) < base.Ui32(v95) {
									v106 = v104
								} else {
									v106 = int32(0)
								}
								if v106 == int32(0) {
								} else {
									v110 = F_memmove(m, v13, v13+v99, v106)
									mBase = m.M
								}
								v112 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v13+v106))) = uint8(v112)
								switch v52 {
								default:
									v115 = v106 << (uint(int32(3)) % 32)
									*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v115)
								case 1:
									*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))) = uint8(v106)
								case 2:
									*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))) = uint16(v106)
								case 3:
									*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9)))) = v106
								case 4:
									*(*int64)(unsafe.Add(mBase, uint32(v13+int32(-17)))) = base.I64_extend_i32_u(v106)
								}
							}
						case 4:
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
							v67 = v66
							if v67 == int32(0) {
							} else {
								v73 = int32(-1)&v67 + v41
								v77 = v16>>(uint(int32(31))%32)&v67 + v16
								v80 = v73 - v77 + int32(1)
								switch v52 {
								default:
									v95 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
								case 1:
									v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
									v95 = v85
								case 2:
									v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
									v95 = v88
								case 3:
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
									v95 = v91
								case 4:
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
									v95 = v94
								}
								v96 = int32(0)
								v98 = base.B2i32(base.Ui32(v77) < base.Ui32(v95))
								if base.Ui32(v77) < base.Ui32(v95) {
									v99 = v77
								} else {
									v99 = v96
								}
								v100 = v95 - v99
								if base.Ui32(v80) < base.Ui32(v100) {
									v102 = v80
								} else {
									v102 = v100
								}
								if v73 < v77 {
									v104 = v96
								} else {
									v104 = v102
								}
								if base.Ui32(v77) < base.Ui32(v95) {
									v106 = v104
								} else {
									v106 = int32(0)
								}
								if v106 == int32(0) {
								} else {
									v110 = F_memmove(m, v13, v13+v99, v106)
									mBase = m.M
								}
								v112 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v13+v106))) = uint8(v112)
								switch v52 {
								default:
									v115 = v106 << (uint(int32(3)) % 32)
									*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v115)
								case 1:
									*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))) = uint8(v106)
								case 2:
									*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))) = uint16(v106)
								case 3:
									*(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9)))) = v106
								case 4:
									*(*int64)(unsafe.Add(mBase, uint32(v13+int32(-17)))) = base.I64_extend_i32_u(v106)
								}
							}
						default:
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
					}
					return
				}
			}
		}
	}
}
func F_updateClientMemoryUsage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v5 == int32(0) {
		F__serverAssert(m, int32(_a781), int32(_a1240), int32(1069))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v9 = F_getClientMemoryUsage(m, l0, int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			if v11&int32(1) == int32(0) {
				v17 = int32(2)
				if v11&v17 == int32(0) {
					if v11&int32(262144) != 0 {
						v35 = v17
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						if v27 != 0 {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
							if v31 == int32(1) {
								v34 = int32(4)
							} else {
								v34 = int32(5)
							}
							v35 = v34
						} else {
							v35 = int32(0)
						}
					}
				} else {
					if v11&int32(4) != 0 {
						if v11&int32(262144) != 0 {
							v35 = v17
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							if v27 != 0 {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
								if v31 == int32(1) {
									v34 = int32(4)
								} else {
									v34 = int32(5)
								}
								v35 = v34
							} else {
								v35 = int32(0)
							}
						}
					} else {
						v35 = int32(1)
					}
				}
			} else {
				v35 = int32(3)
			}
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+226)))
			v37 = int32(2)
			v38 = v36 << (uint(v37) % 32)
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[417])))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
			*(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[417]))) = v41 - v42
			v46 = v35 << (uint(v37) % 32)
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[417])))
			*(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[417]))) = v49 + v9
			*(*int32)(unsafe.Add(mBase, uint32(l0)+304)) = v9
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+226)) = uint8(v35)
			return
		}
	}
}
func F_validateClientCapaFilter(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
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
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	v35 = int32(0)
	goto L1
L1:
	;
	switch v21 & int32(7) {
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
		v42 = int32(0)
		goto L3
	}
L2:
	;
	if base.Ui32(v35) < base.Ui32(v42) {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	if base.Ui32(v42) <= base.Ui32(v35) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v42 = v41
	goto L3
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v42 = v40
	goto L3
L6:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v42 = v39
	goto L3
L7:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v42 = v38
	goto L3
L8:
	;
	v42 = int32(base.Ui32(v21) >> (uint(int32(3)) % 32))
	goto L3
L9:
	;
	goto L2
L10:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v35))))
	if v47 == int32(114) {
		v35 = v35 + int32(1)
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v54 = int32(-1)
	goto L14
L13:
	;
	v54 = int32(0)
	goto L14
L14:
	;
	return v54
}
