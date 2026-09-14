package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_keyspaceEventsFlagsToString(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	v3 = F_sdsempty(m)
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v7 = int32(10236)
	if l0&v7 != v7 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if l0&int32(1) == int32(0) {
		v122 = v113
		goto L39
	} else {
		goto L40
	}
L4:
	;
	if l0&int32(4) == int32(0) {
		v23 = v3
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v13 = F_sdscatlen(m, v3, int32(_a719), int32(1))
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v113 = v13
	goto L3
L7:
	;
	if l0&int32(8) == int32(0) {
		v32 = v23
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v21 = F_sdscatlen(m, v3, int32(_a827), int32(1))
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v23 = v21
	goto L7
L10:
	;
	if l0&int32(16) == int32(0) {
		v41 = v32
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v30 = F_sdscatlen(m, v23, int32(_a828), int32(1))
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v32 = v30
	goto L10
L13:
	;
	if l0&int32(32) == int32(0) {
		v50 = v41
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v39 = F_sdscatlen(m, v32, int32(_a829), int32(1))
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v41 = v39
	goto L13
L16:
	;
	if l0&int32(64) == int32(0) {
		v59 = v50
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v48 = F_sdscatlen(m, v41, int32(_a830), int32(1))
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v50 = v48
	goto L16
L19:
	;
	if l0&int32(128) == int32(0) {
		v68 = v59
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v57 = F_sdscatlen(m, v50, int32(_a831), int32(1))
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v59 = v57
	goto L19
L22:
	;
	if l0&int32(256) == int32(0) {
		v77 = v68
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v66 = F_sdscatlen(m, v59, int32(_a832), int32(1))
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v68 = v66
	goto L22
L25:
	;
	if l0&int32(512) == int32(0) {
		v86 = v77
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v75 = F_sdscatlen(m, v68, int32(_a116), int32(1))
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v77 = v75
	goto L25
L28:
	;
	if l0&int32(1024) == int32(0) {
		v95 = v86
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v84 = F_sdscatlen(m, v77, int32(_a826), int32(1))
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v86 = v84
	goto L28
L31:
	;
	if l0&int32(8192) == int32(0) {
		v104 = v95
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v93 = F_sdscatlen(m, v86, int32(_a825), int32(1))
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v95 = v93
	goto L31
L34:
	;
	if l0&int32(16384) == int32(0) {
		v113 = v104
		goto L3
	} else {
		goto L37
	}
L35:
	;
	v102 = F_sdscatlen(m, v95, int32(_a824), int32(1))
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v104 = v102
	goto L34
L37:
	;
	v111 = F_sdscatlen(m, v104, int32(_a823), int32(1))
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v113 = v111
	goto L3
L39:
	;
	if l0&int32(2) == int32(0) {
		v131 = v122
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v120 = F_sdscatlen(m, v113, int32(_a822), int32(1))
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v122 = v120
	goto L39
L42:
	;
	if l0&int32(2048) == int32(0) {
		v140 = v131
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v129 = F_sdscatlen(m, v122, int32(_a821), int32(1))
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v131 = v129
	goto L42
L45:
	;
	return v140
L46:
	;
	v138 = F_sdscatlen(m, v131, int32(_a820), int32(1))
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v140 = v138
	goto L45
}
func F_kill(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = F___syscall_getpid(m)
	mBase = m.M
	if l0 != v3 {
		*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(63)
		return int32(-1)
	} else {
		v5 = F_raise(m, l1)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v5
		}
	}
}
func F_killAppendOnlyChild(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v8 != int32(2) {
		m.G0 = v5 + int32(16)
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[15]))
		if int32(2) < v12 {
			v23 = *(*int32)(unsafe.Add(mBase, _consts[39]))
			v25 = F_kill(m, v23, int32(10))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				if v25 == int32(-1) {
					v42 = *(*int32)(unsafe.Add(mBase, _consts[39]))
					v44 = v42
				} else {
					for {
						v34 = int32(0)
						v36 = F___syscall_wait4(m, int32(-1), v5+int32(12), v34, v34)
						mBase = m.M
						v37 = F___syscall_ret(m, v36)
						mBase = m.M
						v39 = *(*int32)(unsafe.Add(mBase, _consts[39]))
						if v37 != v39 {
							continue
						} else {
							break
						}
						break
					}
					v44 = v37
				}
				F_aofRemoveTempFile(m, v44, int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					F_resetChildState(m)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, _consts[40])) = int64(-1)
						m.G0 = v5 + int32(16)
						return
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[39]))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v16
			F__serverLog(m, int32(2), int32(_a93), v5)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, _consts[39]))
				v25 = F_kill(m, v23, int32(10))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					if v25 == int32(-1) {
						v42 = *(*int32)(unsafe.Add(mBase, _consts[39]))
						v44 = v42
					} else {
						for {
							v34 = int32(0)
							v36 = F___syscall_wait4(m, int32(-1), v5+int32(12), v34, v34)
							mBase = m.M
							v37 = F___syscall_ret(m, v36)
							mBase = m.M
							v39 = *(*int32)(unsafe.Add(mBase, _consts[39]))
							if v37 != v39 {
								continue
							} else {
								break
							}
							break
						}
						v44 = v37
					}
					F_aofRemoveTempFile(m, v44, int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						F_resetChildState(m)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, _consts[40])) = int64(-1)
							m.G0 = v5 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_killRDBChild(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v7 != int32(1) {
		m.G0 = v4 + int32(16)
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[15]))
		if int32(2) < v11 {
			v22 = *(*int32)(unsafe.Add(mBase, _consts[39]))
			v24 = F_kill(m, v22, int32(10))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				m.G0 = v4 + int32(16)
				return
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _consts[39]))
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = v15
			F__serverLog(m, int32(2), int32(_a888), v4)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _consts[39]))
				v24 = F_kill(m, v22, int32(10))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					m.G0 = v4 + int32(16)
					return
				}
			}
		}
	}
}
