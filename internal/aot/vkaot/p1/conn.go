package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_connBlock(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = int32(-1)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4 == v3 {
		v12 = v3
		return v12
	} else {
		v8 = F_anetBlock(m, int32(0), v4)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = v8
			return v12
		}
	}
}
func F_connCreateSocket(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = F_valkey_calloc(m, int32(40))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(1024)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+20)) = uint16(v7)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v3))) = int32(_a_F_connCreateSocket_0)
		return v3
	}
}
func F_connCreateUnix(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = F_valkey_calloc(m, int32(40))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(1024)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+20)) = uint16(v7)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v3))) = int32(_a_F_connCreateUnix_0)
		return v3
	}
}
func F_connSocketAccept(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v4 != int32(2) {
		v39 = int32(-1)
		return v39
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(3)
		v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		v11 = v9 + int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v11)
		if l1 == int32(0) {
			v22 = v9
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v22)
			v24 = int32(0)
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
			if v25&int32(1) == v24 {
				v39 = v24
				return v39
			} else {
				if v22&int32(65535) != 0 {
					v39 = int32(-1)
					return v39
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
					m.T0[v33].(func(*base.Module, int32))(m, l0)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v39 = int32(-1)
						return v39
					}
				}
			}
		} else {
			m.T0[l1].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
				v22 = v19 + int32(-1)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v22)
				v24 = int32(0)
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
				if v25&int32(1) == v24 {
					v39 = v24
					return v39
				} else {
					if v22&int32(65535) != 0 {
						v39 = int32(-1)
						return v39
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
						m.T0[v33].(func(*base.Module, int32))(m, l0)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v39 = int32(-1)
							return v39
						}
					}
				}
			}
		}
	}
}
func F_connSocketAcceptHandler(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v35 int32
	_ = v35
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
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
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
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	v11 = m.G0
	v13 = v11 - int32(128)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_connSocketAcceptHandler[0]))
	v20 = v13 + int32(112)
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v13)+104)) = v21
	if v16 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(128)
	return
L2:
	;
	v35 = v16
	goto L3
L3:
	;
	v45 = F_anetTcpAccept(m, int32(_a_F_connSocketAcceptHandler_0), l1, v13+int32(48), int32(46), v13+int32(124))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L1
L5:
	;
	v116 = v35 + int32(-1)
	if v116 != 0 {
		v35 = v116
		goto L3
	} else {
		goto L24
	}
L6:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_connSocketAcceptHandler[1]))
	if int32(1) < v69 {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	return
L8:
	;
	if v45 != int32(-1) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v49 = int32(9116376)
	goto L10
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_connSocketAcceptHandler[2]))
	goto L11
L11:
	;
	if v50 == int32(13) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_connSocketAcceptHandler[2]))
	if v53 == int32(6) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_connSocketAcceptHandler[1]))
	if int32(3) < v57 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_connSocketAcceptHandler_0)
	F__serverLog(m, int32(3), int32(_a_F_connSocketAcceptHandler_1), v13)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L1
L16:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_connSocketAcceptHandler[3]))
	if v84 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v13 + int32(48)
	F__serverLog(m, int32(1), int32(_a_F_connSocketAcceptHandler_2), v13+int32(32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v91 = F_valkey_calloc(m, int32(40))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	v88 = F_anetKeepAlive(m, int32(0), v45, v84)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v93 = int32(1024)
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+20)) = uint16(v93)
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(_a_F_connSocketAcceptHandler_3)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = int32(2)
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(24)))) = v104
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v13)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v106
	F_acceptCommonHandler(m, v91, v13+int32(16), v13+int32(48))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	goto L5
L24:
	;
	goto L4
}
func F_connSocketClose(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3 == int32(-1) {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v15 == int32(0) {
			F_valkey_free(m, l0)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				return
			}
		} else {
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
			v20 = v18 | int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v20)
			return
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_connSocketClose[0]))
		F_aeDeleteFileEvent(m, v7, v3, int32(3))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v12 = F_close(m, v11)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(-1)
			v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
			if v15 == int32(0) {
				F_valkey_free(m, l0)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			} else {
				v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
				v20 = v18 | int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v20)
				return
			}
		}
	}
}
func F_connSocketGetType(m *base.Module) int32 {
	return int32(0)
}
func F_connSocketSyncRead(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = int32(0)
	v13 = F_mstime(m)
	mBase = m.M
	if l2 == v6 {
		v72 = v6
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v72
L2:
	;
	if v72 != int32(-1) {
		goto L1
	} else {
		goto L18
	}
L3:
	;
	goto L2
L4:
	;
	v17 = l1
	v18 = l2
	v20 = v6
	v22 = l3
	goto L7
L5:
	;
	v72 = int32(-1)
	goto L3
L6:
	;
	v59 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v57
	goto L5
L7:
	;
	v27 = F_read(m, v5, v17, v18)
	mBase = m.M
	switch v27 + int32(1) {
	case 0:
		goto L11
	case 1:
		v57 = int32(15)
		goto L6
	default:
		goto L10
	}
L8:
	;
	v57 = int32(73)
	goto L6
L9:
	;
	v43 = int64(10)
	if v43 < v22 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v34 = v27 + v20
	v35 = v18 - v27
	if v35 == int32(0) {
		v72 = v34
		goto L3
	} else {
		goto L13
	}
L11:
	;
	v30 = F___errno_location(m)
	mBase = m.M
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v31 == int32(6) {
		v39 = v17
		v40 = v18
		v41 = v20
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L5
L13:
	;
	v39 = v17 + v27
	v40 = v35
	v41 = v34
	goto L9
L14:
	;
	v46 = v22
	goto L16
L15:
	;
	v46 = v43
	goto L16
L16:
	;
	v47 = F_aeWait(m, v5, int32(1), v46)
	mBase = m.M
	v48 = F_mstime(m)
	mBase = m.M
	v49 = v48 - v13
	if v49 < l3 {
		v17 = v39
		v18 = v40
		v20 = v41
		v22 = l3 - v49
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L8
L18:
	;
	goto L19
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_connSocketSyncRead[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v81
	goto L1
}
func F_connSocketWritev(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = F_writev(m, v5, l1, l2)
	mBase = m.M
	if int32(-1) < v6 {
	} else {
		v9 = int32(9116376)
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_connSocketWritev[0]))
		if v10 == int32(6) {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v10
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_connSocketWritev[0]))
			if v14 == int32(27) {
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v17 != int32(3) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
				}
			}
		}
	}
	return v6
}
func F_connTypeOfCluster(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_connTypeOfCluster[0]))
	if v2 == int32(0) {
		v10 = F_connectionTypeTcp(m)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	} else {
		v5 = F_connectionTypeTls(m)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v5
		}
	}
}
func F_connUnixClose(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = F_connectionTypeTcp(m)
	mBase = m.M
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(v2)+52))
		m.T0[v4].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_connUnixGetType(m *base.Module) int32 {
	return int32(1)
}
func F_connUnixSyncReadLine(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = l2 + int32(-1)
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v64
L2:
	;
	m.G0 = v12 + int32(16)
	goto L1
L3:
	;
	v19 = l1
	v20 = v16
	v23 = int32(0)
	goto L5
L4:
	;
	v64 = int32(0)
	goto L2
L5:
	;
	v26 = int32(-1)
	v30 = F_syncRead(m, v5, v12+int32(15), int32(1), l3)
	mBase = m.M
	if v30 == v26 {
		v64 = v26
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v64 = v16
	goto L2
L7:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v33 != int32(10) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v48 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)) = uint8(v48)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v33)
	v51 = int32(1)
	v56 = v20 + int32(-1)
	if v56 != 0 {
		v19 = v19 + v51
		v20 = v56
		v23 = v23 + v51
		goto L5
	} else {
		goto L13
	}
L9:
	;
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v36)
	if v23 == v36 {
		v64 = v36
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v42 = v19 + int32(-1)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v43 != int32(13) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v64 = v23
	goto L2
L12:
	;
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v46)
	goto L11
L13:
	;
	goto L6
}
