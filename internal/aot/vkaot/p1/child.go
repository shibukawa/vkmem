package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_childSnapshotForSyncSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
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
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v3
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_childSnapshotForSyncSlot[0]))
	if v17 < int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v174
L2:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v120&int32(6) != 0 {
		goto L28
	} else {
		goto L29
	}
L3:
	;
	v24 = v3
	goto L4
L4:
	;
	v28 = v10 + int32(4)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v29
	goto L6
L5:
	;
	goto L2
L6:
	;
	v34 = v10 + int32(4)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v36 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v109 = v24 + int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_childSnapshotForSyncSlot[0]))
	if v109 < v111 {
		v24 = v109
		goto L4
	} else {
		goto L27
	}
L8:
	;
	if v36 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v36+base.B2i32(v39 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v45
	goto L9
L11:
	;
	v50 = v36
	goto L12
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v58 < v57 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L7
L14:
	;
	v88 = v10 + int32(4)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v90 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	v61 = v57
	goto L17
L16:
	;
	v174 = int32(-1)
	goto L1
L17:
	;
	v69 = F_rewriteSlotToAppendOnlyFileRio(m, l0, v24, v61, v10+int32(12))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	if v69 == int32(-1) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v75 <= v61 {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	v61 = v61 + int32(1)
	goto L17
L23:
	;
	if v90 != 0 {
		v50 = v90
		goto L12
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v90+base.B2i32(v93 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v99
	goto L24
L26:
	;
	goto L13
L27:
	;
	goto L5
L28:
	;
	v162 = F_rioWriteBulkString(m, l0, int32(_a_F_childSnapshotForSyncSlot_0), int32(7))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L19
	} else {
		goto L45
	}
L29:
	;
	v130 = int32(_a_F_childSnapshotForSyncSlot_1)
	v131 = int32(4)
	goto L30
L30:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v132) < base.Ui32(v131) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L28
L32:
	;
	v134 = v132
	goto L34
L33:
	;
	v134 = v131
	goto L34
L34:
	;
	if v132 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v135 = v134
	goto L37
L36:
	;
	v135 = v131
	goto L37
L37:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v136 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v142 = m.T0[v141].(func(*base.Module, int32, int32, int32) int32)(m, l0, v130, v135)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L19
	} else {
		goto L42
	}
L39:
	;
	m.T0[v136].(func(*base.Module, int32, int32, int32))(m, l0, v130, v135)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v148 + v135
	v152 = v131 - v135
	if v152 != 0 {
		v130 = v130 + v135
		v131 = v152
		goto L30
	} else {
		goto L44
	}
L42:
	;
	if v142 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v144 | int64(2)
	goto L28
L44:
	;
	goto L31
L45:
	;
	v166 = F_rioWriteBulkString(m, l0, int32(_a_F_childSnapshotForSyncSlot_2), int32(9))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L19
	} else {
		goto L46
	}
L46:
	;
	v170 = F_rioWriteBulkString(m, l0, int32(_a_F_childSnapshotForSyncSlot_3), int32(12))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L19
	} else {
		goto L47
	}
L47:
	;
	v174 = int32(0)
	goto L1
}
func F_closeChildInfoPipe(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_closeChildInfoPipe[0]))
	if v3 != int32(-1) {
		v10 = F_close(m, v3)
		mBase = m.M
		v11 = int32(_a_F_closeChildInfoPipe_0)
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_closeChildInfoPipe[1]))
		v13 = F_close(m, v12)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, _c_F_closeChildInfoPipe[2])) = int32(0)
		*(*int64)(unsafe.Add(mBase, _c_F_closeChildInfoPipe[0])) = int64(-1)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_closeChildInfoPipe[1]))
		if v7 == int32(-1) {
		} else {
			v10 = F_close(m, v3)
			mBase = m.M
			v11 = int32(_a_F_closeChildInfoPipe_0)
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_closeChildInfoPipe[1]))
			v13 = F_close(m, v12)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, _c_F_closeChildInfoPipe[2])) = int32(0)
			*(*int64)(unsafe.Add(mBase, _c_F_closeChildInfoPipe[0])) = int64(-1)
		}
	}
	return
}
func F_openChildInfoPipe(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v4 = int32(_a_F_openChildInfoPipe_0)
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v14 = F_pipe(m, v4)
	mBase = m.M
	if v14 != 0 {
		v71 = int32(-1)
	} else {
		v43 = *(*int32)(unsafe.Add(mBase, _c_F_openChildInfoPipe[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(2048)
		v48 = F_fcntl(m, v43, int32(4), v11+int32(16))
		mBase = m.M
		if v48 != 0 {
			v64 = *(*int32)(unsafe.Add(mBase, _c_F_openChildInfoPipe[0]))
			v65 = F_close(m, v64)
			mBase = m.M
			v66 = *(*int32)(unsafe.Add(mBase, _c_F_openChildInfoPipe[1]))
			v67 = F_close(m, v66)
			mBase = m.M
			v71 = int32(-1)
		} else {
			v71 = int32(0)
		}
	}
	m.G0 = v11 + int32(64)
	if v71 != int32(-1) {
		*(*int32)(unsafe.Add(mBase, _c_F_openChildInfoPipe[2])) = int32(0)
	} else {
		v78 = *(*int32)(unsafe.Add(mBase, _c_F_openChildInfoPipe[0]))
		if v78 != int32(-1) {
			v85 = F_close(m, v78)
			mBase = m.M
			v86 = int32(_a_F_openChildInfoPipe_1)
			v87 = *(*int32)(unsafe.Add(mBase, _c_F_openChildInfoPipe[1]))
			v88 = F_close(m, v87)
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, _c_F_openChildInfoPipe[0])) = int64(-1)
			*(*int32)(unsafe.Add(mBase, _c_F_openChildInfoPipe[2])) = int32(0)
		} else {
			v82 = *(*int32)(unsafe.Add(mBase, _c_F_openChildInfoPipe[1]))
			if v82 == int32(-1) {
			} else {
				v85 = F_close(m, v78)
				mBase = m.M
				v86 = int32(_a_F_openChildInfoPipe_1)
				v87 = *(*int32)(unsafe.Add(mBase, _c_F_openChildInfoPipe[1]))
				v88 = F_close(m, v87)
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, _c_F_openChildInfoPipe[0])) = int64(-1)
				*(*int32)(unsafe.Add(mBase, _c_F_openChildInfoPipe[2])) = int32(0)
			}
		}
	}
	return
}
func F_sendChildInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v7 int32
	_ = v7
	F_sendChildInfoGeneric(m, l0, l1, int32(0), float64(-1), l2)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_sendChildInfoGeneric(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v30 int64
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int64
	_ = v79
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[0]))
	if v18 == int32(-1) {
		m.G0 = v15 + int32(80)
		return
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[1]))
		v23 = m.T0[v22].(func(*base.Module) int64)(m)
		mBase = m.M
		if l0 != 0 {
			v35 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2])) = v35
			v39 = int32(0)
			v41 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[1]))
			v42 = m.T0[v41].(func(*base.Module) int64)(m)
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[3])) = v42
			*(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[4])) = v42 - v23
			v48 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2]))
			v50 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[5]))
			if base.Ui32(v48) <= base.Ui32(v50) {
				v54 = v50
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[5])) = v48
				v54 = v48
			}
			v55 = int32(0)
			v57 = *(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[6]))
			v59 = v57 + base.I64_extend_i32_u(v48)
			*(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[6])) = v59
			v63 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[7]))
			v65 = v63 + int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[7])) = v65
			if v48|l0 == v55 {
				v98 = v42
				*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l0
				*(*float64)(unsafe.Add(mBase, uint32(v15)+64)) = l3
				*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v98
				v109 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2]))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v109
				v112 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[0]))
				v115 = int32(32)
				v116 = F_write(m, v112, v15+int32(48), v115)
				mBase = m.M
				if v116 == v115 {
					m.G0 = v15 + int32(80)
					return
				} else {
					v120 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[8]))
					if int32(3) < v120 {
						F__exit(m, int32(1))
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					} else {
						v124 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[9]))
						v125 = F___strerror_l(m, v124, v124)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v125
						F__serverLog(m, int32(3), int32(_a_F_sendChildInfoGeneric_0), v15)
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return
						} else {
							F__exit(m, int32(1))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if l0 != 0 {
					v72 = int32(2)
				} else {
					v72 = int32(1)
				}
				v74 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[8]))
				if v72 < v74 {
					v98 = v42
					*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l0
					*(*float64)(unsafe.Add(mBase, uint32(v15)+64)) = l3
					*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v98
					v109 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v109
					v112 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[0]))
					v115 = int32(32)
					v116 = F_write(m, v112, v15+int32(48), v115)
					mBase = m.M
					if v116 == v115 {
						m.G0 = v15 + int32(80)
						return
					} else {
						v120 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[8]))
						if int32(3) < v120 {
							F__exit(m, int32(1))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						} else {
							v124 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[9]))
							v125 = F___strerror_l(m, v124, v124)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v15))) = v125
							F__serverLog(m, int32(3), int32(_a_F_sendChildInfoGeneric_0), v15)
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return
							} else {
								F__exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v79 = base.I64_div_u_s(v59, base.I64_extend_i32_u(v65))
					*(*int64)(unsafe.Add(mBase, uint32(v15+int32(32)))) = int64(base.Ui64(v79) >> (uint(int64(20)) % 64))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l4
					v84 = int32(20)
					*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(base.Ui32(v54) >> (uint(v84) % 32))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(base.Ui32(v48) >> (uint(v84) % 32))
					F__serverLog(m, v72, int32(_a_F_sendChildInfoGeneric_1), v15+int32(16))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						v96 = *(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[3]))
						v98 = v96
						*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l0
						*(*float64)(unsafe.Add(mBase, uint32(v15)+64)) = l3
						*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v98
						v109 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v109
						v112 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[0]))
						v115 = int32(32)
						v116 = F_write(m, v112, v15+int32(48), v115)
						mBase = m.M
						if v116 == v115 {
							m.G0 = v15 + int32(80)
							return
						} else {
							v120 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[8]))
							if int32(3) < v120 {
								F__exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							} else {
								v124 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[9]))
								v125 = F___strerror_l(m, v124, v124)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = v125
								F__serverLog(m, int32(3), int32(_a_F_sendChildInfoGeneric_0), v15)
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return
								} else {
									F__exit(m, int32(1))
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
		} else {
			v25 = *(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[3]))
			if v25 == int64(0) {
				v35 = int32(0)
				*(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2])) = v35
				v39 = int32(0)
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[1]))
				v42 = m.T0[v41].(func(*base.Module) int64)(m)
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[3])) = v42
				*(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[4])) = v42 - v23
				v48 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2]))
				v50 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[5]))
				if base.Ui32(v48) <= base.Ui32(v50) {
					v54 = v50
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[5])) = v48
					v54 = v48
				}
				v55 = int32(0)
				v57 = *(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[6]))
				v59 = v57 + base.I64_extend_i32_u(v48)
				*(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[6])) = v59
				v63 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[7]))
				v65 = v63 + int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[7])) = v65
				if v48|l0 == v55 {
					v98 = v42
					*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l0
					*(*float64)(unsafe.Add(mBase, uint32(v15)+64)) = l3
					*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v98
					v109 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v109
					v112 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[0]))
					v115 = int32(32)
					v116 = F_write(m, v112, v15+int32(48), v115)
					mBase = m.M
					if v116 == v115 {
						m.G0 = v15 + int32(80)
						return
					} else {
						v120 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[8]))
						if int32(3) < v120 {
							F__exit(m, int32(1))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						} else {
							v124 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[9]))
							v125 = F___strerror_l(m, v124, v124)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v15))) = v125
							F__serverLog(m, int32(3), int32(_a_F_sendChildInfoGeneric_0), v15)
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return
							} else {
								F__exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					if l0 != 0 {
						v72 = int32(2)
					} else {
						v72 = int32(1)
					}
					v74 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[8]))
					if v72 < v74 {
						v98 = v42
						*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l0
						*(*float64)(unsafe.Add(mBase, uint32(v15)+64)) = l3
						*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v98
						v109 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v109
						v112 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[0]))
						v115 = int32(32)
						v116 = F_write(m, v112, v15+int32(48), v115)
						mBase = m.M
						if v116 == v115 {
							m.G0 = v15 + int32(80)
							return
						} else {
							v120 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[8]))
							if int32(3) < v120 {
								F__exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							} else {
								v124 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[9]))
								v125 = F___strerror_l(m, v124, v124)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = v125
								F__serverLog(m, int32(3), int32(_a_F_sendChildInfoGeneric_0), v15)
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return
								} else {
									F__exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v79 = base.I64_div_u_s(v59, base.I64_extend_i32_u(v65))
						*(*int64)(unsafe.Add(mBase, uint32(v15+int32(32)))) = int64(base.Ui64(v79) >> (uint(int64(20)) % 64))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l4
						v84 = int32(20)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(base.Ui32(v54) >> (uint(v84) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(base.Ui32(v48) >> (uint(v84) % 32))
						F__serverLog(m, v72, int32(_a_F_sendChildInfoGeneric_1), v15+int32(16))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							v96 = *(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[3]))
							v98 = v96
							*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l0
							*(*float64)(unsafe.Add(mBase, uint32(v15)+64)) = l3
							*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v98
							v109 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v109
							v112 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[0]))
							v115 = int32(32)
							v116 = F_write(m, v112, v15+int32(48), v115)
							mBase = m.M
							if v116 == v115 {
								m.G0 = v15 + int32(80)
								return
							} else {
								v120 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[8]))
								if int32(3) < v120 {
									F__exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								} else {
									v124 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[9]))
									v125 = F___strerror_l(m, v124, v124)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v15))) = v125
									F__serverLog(m, int32(3), int32(_a_F_sendChildInfoGeneric_0), v15)
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return
									} else {
										F__exit(m, int32(1))
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
			} else {
				v30 = *(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[4]))
				if base.Ui64(v23-v25) <= base.Ui64(v30*int64(100)) {
					v98 = v25
					*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l0
					*(*float64)(unsafe.Add(mBase, uint32(v15)+64)) = l3
					*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v98
					v109 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v109
					v112 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[0]))
					v115 = int32(32)
					v116 = F_write(m, v112, v15+int32(48), v115)
					mBase = m.M
					if v116 == v115 {
						m.G0 = v15 + int32(80)
						return
					} else {
						v120 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[8]))
						if int32(3) < v120 {
							F__exit(m, int32(1))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						} else {
							v124 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[9]))
							v125 = F___strerror_l(m, v124, v124)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v15))) = v125
							F__serverLog(m, int32(3), int32(_a_F_sendChildInfoGeneric_0), v15)
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return
							} else {
								F__exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v35 = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2])) = v35
					v39 = int32(0)
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[1]))
					v42 = m.T0[v41].(func(*base.Module) int64)(m)
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[3])) = v42
					*(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[4])) = v42 - v23
					v48 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2]))
					v50 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[5]))
					if base.Ui32(v48) <= base.Ui32(v50) {
						v54 = v50
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[5])) = v48
						v54 = v48
					}
					v55 = int32(0)
					v57 = *(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[6]))
					v59 = v57 + base.I64_extend_i32_u(v48)
					*(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[6])) = v59
					v63 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[7]))
					v65 = v63 + int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[7])) = v65
					if v48|l0 == v55 {
						v98 = v42
						*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l0
						*(*float64)(unsafe.Add(mBase, uint32(v15)+64)) = l3
						*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v98
						v109 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v109
						v112 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[0]))
						v115 = int32(32)
						v116 = F_write(m, v112, v15+int32(48), v115)
						mBase = m.M
						if v116 == v115 {
							m.G0 = v15 + int32(80)
							return
						} else {
							v120 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[8]))
							if int32(3) < v120 {
								F__exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							} else {
								v124 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[9]))
								v125 = F___strerror_l(m, v124, v124)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = v125
								F__serverLog(m, int32(3), int32(_a_F_sendChildInfoGeneric_0), v15)
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return
								} else {
									F__exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						if l0 != 0 {
							v72 = int32(2)
						} else {
							v72 = int32(1)
						}
						v74 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[8]))
						if v72 < v74 {
							v98 = v42
							*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l0
							*(*float64)(unsafe.Add(mBase, uint32(v15)+64)) = l3
							*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v98
							v109 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v109
							v112 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[0]))
							v115 = int32(32)
							v116 = F_write(m, v112, v15+int32(48), v115)
							mBase = m.M
							if v116 == v115 {
								m.G0 = v15 + int32(80)
								return
							} else {
								v120 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[8]))
								if int32(3) < v120 {
									F__exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								} else {
									v124 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[9]))
									v125 = F___strerror_l(m, v124, v124)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v15))) = v125
									F__serverLog(m, int32(3), int32(_a_F_sendChildInfoGeneric_0), v15)
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return
									} else {
										F__exit(m, int32(1))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v79 = base.I64_div_u_s(v59, base.I64_extend_i32_u(v65))
							*(*int64)(unsafe.Add(mBase, uint32(v15+int32(32)))) = int64(base.Ui64(v79) >> (uint(int64(20)) % 64))
							*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l4
							v84 = int32(20)
							*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(base.Ui32(v54) >> (uint(v84) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(base.Ui32(v48) >> (uint(v84) % 32))
							F__serverLog(m, v72, int32(_a_F_sendChildInfoGeneric_1), v15+int32(16))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								v96 = *(*int64)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[3]))
								v98 = v96
								*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l0
								*(*float64)(unsafe.Add(mBase, uint32(v15)+64)) = l3
								*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v98
								v109 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v109
								v112 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[0]))
								v115 = int32(32)
								v116 = F_write(m, v112, v15+int32(48), v115)
								mBase = m.M
								if v116 == v115 {
									m.G0 = v15 + int32(80)
									return
								} else {
									v120 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[8]))
									if int32(3) < v120 {
										F__exit(m, int32(1))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									} else {
										v124 = *(*int32)(unsafe.Add(mBase, _c_F_sendChildInfoGeneric[9]))
										v125 = F___strerror_l(m, v124, v124)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v15))) = v125
										F__serverLog(m, int32(3), int32(_a_F_sendChildInfoGeneric_0), v15)
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return
										} else {
											F__exit(m, int32(1))
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
func F_shouldStartChildReplication(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int64
	_ = v169
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	v4 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_shouldStartChildReplication[0]))
	if v22 != int32(-1) {
		v187 = v4
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_shouldStartChildReplication[1]))
		v28 = v18 + int32(8)
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v28))) = v29
		v33 = int32(0)
		v35 = v18 + int32(8)
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
		if v37 == v33 {
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v37+base.B2i32(v40 == int32(0))<<(uint(int32(2))%32))))
			*(*int32)(unsafe.Add(mBase, uint32(v35))) = v46
		}
		if v37 == int32(0) {
			v187 = v33
		} else {
			v58 = v37
			v59 = int32(1)
			v60 = int64(0)
			v61 = int32(0)
			v64 = v4
			v65 = v4
			v67 = v4
			for {
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+104))
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
				if v70 != int32(6) {
					v133 = v59
					v134 = v60
					v135 = v61
					v137 = v64
					v138 = v65
					v140 = v67
				} else {
					v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v69)+162)))
					if v59 == int32(0) {
						v89 = int32(0)
						if v65 != v73 {
							v133 = v89
							v134 = v60
							v135 = v61
							v137 = v64
							v138 = v65
							v140 = v67
						} else {
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+160)))
							if v91&int32(1) != 0 {
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v69)+156))
								if v95 <= int32(589823) {
									if v95 < int32(459264) {
										v119 = int32(11)
									} else {
										v102 = int32(_a_F_shouldStartChildReplication_0)
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
										v119 = v103
									}
								} else {
									v102 = int32(_a_F_shouldStartChildReplication_1)
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
									v119 = v103
								}
							} else {
								v119 = int32(80)
							}
							if v64 != v119 {
								v133 = v89
								v134 = v60
								v135 = v61
								v137 = v64
								v138 = v65
								v140 = v67
							} else {
								v122 = *(*int64)(unsafe.Add(mBase, _c_F_shouldStartChildReplication[2]))
								v123 = *(*int64)(unsafe.Add(mBase, uint32(v68)+88))
								v124 = v122 - v123
								if v60 < v124 {
									v126 = v124
								} else {
									v126 = v60
								}
								v129 = *(*int32)(unsafe.Add(mBase, uint32(v68)+104))
								v130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129)+160)))
								v133 = v89
								v134 = v126
								v135 = v61 + int32(1)
								v137 = v64
								v138 = v65
								v140 = v67 & v130
							}
						}
					} else {
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+160)))
						if v76&int32(1) != 0 {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v69)+156))
							if v80 <= int32(589823) {
								if v80 < int32(459264) {
									v106 = int32(11)
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, _c_F_shouldStartChildReplication[3]))
									v106 = v88
								}
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, _c_F_shouldStartChildReplication[4]))
								v106 = v84
							}
						} else {
							v106 = int32(80)
						}
						v108 = *(*int64)(unsafe.Add(mBase, _c_F_shouldStartChildReplication[2]))
						v109 = *(*int64)(unsafe.Add(mBase, uint32(v68)+88))
						v110 = v108 - v109
						if v60 < v110 {
							v112 = v110
						} else {
							v112 = v60
						}
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v68)+104))
						v116 = int32(*(*int16)(unsafe.Add(mBase, uint32(v115)+160)))
						v133 = int32(0)
						v134 = v112
						v135 = v61 + int32(1)
						v137 = v106
						v138 = v73
						v140 = v116
					}
				}
				v142 = v18 + int32(8)
				v144 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
				if v144 == int32(0) {
				} else {
					v147 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
					v153 = *(*int32)(unsafe.Add(mBase, uint32(v144+base.B2i32(v147 == int32(0))<<(uint(int32(2))%32))))
					*(*int32)(unsafe.Add(mBase, uint32(v142))) = v153
				}
				if v144 != 0 {
					v58 = v144
					v59 = v133
					v60 = v134
					v61 = v135
					v64 = v137
					v65 = v138
					v67 = v140
					continue
				} else {
					break
				}
				break
			}
			if v135 == int32(0) {
				v187 = v33
			} else {
				v158 = *(*int32)(unsafe.Add(mBase, _c_F_shouldStartChildReplication[5]))
				if v158 == int32(0) {
					if l0 == int32(0) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v140
					}
					if l1 == int32(0) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v138
					}
					v179 = int32(1)
					if l2 == int32(0) {
						v187 = v179
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v137
						v187 = v179
					}
				} else {
					v161 = int32(0)
					v163 = *(*int32)(unsafe.Add(mBase, _c_F_shouldStartChildReplication[6]))
					if base.B2i32(v161 < v163)&base.B2i32(v163 <= v135) != 0 {
						if l0 == int32(0) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v140
						}
						if l1 == int32(0) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v138
						}
						v179 = int32(1)
						if l2 == int32(0) {
							v187 = v179
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v137
							v187 = v179
						}
					} else {
						v169 = int64(*(*int32)(unsafe.Add(mBase, _c_F_shouldStartChildReplication[7])))
						if v134 < v169 {
							v187 = v161
						} else {
							if l0 == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v140
							}
							if l1 == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v138
							}
							v179 = int32(1)
							if l2 == int32(0) {
								v187 = v179
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v137
								v187 = v179
							}
						}
					}
				}
			}
		}
	}
	m.G0 = v18 + int32(16)
	return v187
}
