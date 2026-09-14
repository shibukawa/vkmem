package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_bioCreateFsyncJob(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v6 = F_allocBioJob(m, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
		v10 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)))
		v19 = v12&int32(253) | l2<<(uint(v10)%32)&int32(2)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)) = uint8(v19)
		v21 = int32(0)
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_bioCreateFsyncJob[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_bioCreateFsyncJob[0])) = v23 + v10
		F_bioExecuteJob(m, v6)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			return
		}
	}
}
func F_bioInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v20 int32
	_ = v20
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	v2 = m.G0
	v4 = v2 - int32(48)
	m.G0 = v4
	v7 = F_mutexQueueCreate(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_bioInit[0])) = v7
		v11 = F_mutexQueueCreate(m)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_bioInit[1])) = v11
			v15 = F_mutexQueueCreate(m)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_bioInit[2])) = v15
				v19 = F_mutexQueueCreate(m)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_bioInit[3])) = v19
					v23 = F_mutexQueueCreate(m)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_bioInit[4])) = v23
						v27 = v4 + int32(4)
						v31 = m.G0
						v33 = v31 - int32(16)
						m.G0 = v33
						v35 = F_pthread_attr_init(m, v27)
						mBase = m.M
						v38 = F_pthread_attr_getstacksize(m, v27, v33+int32(12))
						mBase = m.M
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
						v40 = int32(1)
						if base.Ui32(v40) < base.Ui32(v39) {
							v43 = v39
						} else {
							v43 = v40
						}
						v47 = v43
						for {
							if base.Ui32(v47) < base.Ui32(int32(4194304)) {
								v47 = v47 << (uint(int32(1)) % 32)
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v47
						v53 = F_pthread_attr_setstacksize(m, v27, v47)
						mBase = m.M
						m.G0 = v33 + int32(16)
						m.G0 = v4 + int32(48)
						return
					}
				}
			}
		}
	}
}
func F_handleBioThreadFinishedRDBDownload(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v71 int64
	_ = v71
	var v76 int64
	_ = v76
	var v81 int64
	_ = v81
	var v86 int64
	_ = v86
	var v91 int64
	_ = v91
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
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
	var v166 int32
	_ = v166
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[0]))
	if base.Ui32(v10) < base.Ui32(int32(2)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_handleBioThreadFinishedRDBDownload_4), int32(_a_F_handleBioThreadFinishedRDBDownload_5), int32(5242))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L9
	} else {
		goto L45
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_handleBioThreadFinishedRDBDownload_7), int32(_a_F_handleBioThreadFinishedRDBDownload_5), int32(5221))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L9
	} else {
		goto L44
	}
L3:
	;
	m.G0 = v7 + int32(80)
	return
L4:
	;
	if v10 != int32(3) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if v10 == int32(2) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[3]))
	if int32(3) < v16 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v25 = F_cancelReplicationHandshake(m, int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L9
	} else {
		goto L11
	}
L8:
	;
	F__serverLog(m, int32(3), int32(_a_F_handleBioThreadFinishedRDBDownload_8), int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	goto L7
L11:
	;
	goto L3
L12:
	;
	v31 = int32(_a_F_handleBioThreadFinishedRDBDownload_0)
	*(*int32)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[0])) = int32(0)
	v35 = *(*int64)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[1]))
	v37 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[1])) = v37
	v41 = int32(_a_F_handleBioThreadFinishedRDBDownload_1)
	v42 = *(*int64)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[2]))
	*(*int64)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[2])) = v37
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[3]))
	if int32(2) < v46 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[14]))
	if v30 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v65 = int32(0)
	v66 = *(*int64)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(72)))) = v66
	v71 = *(*int64)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(64)))) = v71
	v76 = *(*int64)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(56)))) = v76
	v81 = *(*int64)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[7]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(48)))) = v81
	v86 = *(*int64)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v86
	v91 = *(*int64)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(32)))) = v91
	v96 = *(*int64)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[10]))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v96
	v99 = *(*int64)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[11]))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v99
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[12]))
	if v105 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	F__serverLog(m, int32(2), int32(_a_F_handleBioThreadFinishedRDBDownload_6), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[3]))
	if int32(2) < v55 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	F__serverLog(m, int32(2), int32(_a_F_handleBioThreadFinishedRDBDownload_9), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v106 = int32(4516)
	goto L22
L21:
	;
	v106 = int32(4512)
	goto L22
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)+uint32(_c_F_handleBioThreadFinishedRDBDownload[13])))
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[14]))
	if v110 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[15]))
	if v116 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	if v108 == int32(0) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[16]))
	if v122 != int32(1) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	F_stopAppendOnly(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+84))
	v141 = m.T0[v140].(func(*base.Module, int32, int32) int32)(m, v108, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L35
	}
L30:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[3]))
	if int32(2) < v126 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_killRDBChild(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L9
	} else {
		goto L34
	}
L32:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v130
	F__serverLog(m, int32(2), int32(_a_F_handleBioThreadFinishedRDBDownload_3), v7)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L29
L35:
	;
	v145 = F_replicaLoadPrimaryRDBFromDisk(m, v7+int32(16))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L9
	} else {
		goto L37
	}
L36:
	;
	F_replicaAfterLoadPrimaryRDB(m, v108, v7+int32(16), int32(1))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L43
	}
L37:
	;
	if v145 != int32(-1) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[3]))
	if int32(3) < v150 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v159 = F_cancelReplicationHandshake(m, int32(1))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L9
	} else {
		goto L42
	}
L40:
	;
	F__serverLog(m, int32(3), int32(_a_F_handleBioThreadFinishedRDBDownload_2), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L3
L43:
	;
	v166 = int32(_a_F_handleBioThreadFinishedRDBDownload_0)
	*(*int64)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[17])) = v42
	*(*int64)(unsafe.Add(mBase, _c_F_handleBioThreadFinishedRDBDownload[18])) = v35
	goto L3
L44:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
