package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_connKeepAlive(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v4 = int32(-1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v5 == v4 {
		v13 = v4
		return v13
	} else {
		v9 = F_anetKeepAlive(m, int32(0), v5, l1)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = v9
			return v13
		}
	}
}
func F_connNonBlock(m *base.Module, l0 int32) int32 {
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
		v8 = F_anetNonBlock(m, int32(0), v4)
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
func F_connSendTimeout(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = F_anetSendTimeout(m, int32(0), v4, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_connSocketConnect(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
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
	v8 = F_anetTcpNonBlockBestEffortBindConnect(m, int32(0), l1, l2, l3, l4)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != int32(-1) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l5
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
			v26 = *(*int32)(unsafe.Add(mBase, _consts[279]))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
			v30 = F_aeCreateFileEvent(m, v26, v8, int32(2), v29, l0)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
			v17 = *(*int32)(unsafe.Add(mBase, _consts[5]))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17
			return int32(-1)
		}
	}
}
func F_connSocketSetReadHandler(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v28 int32
	_ = v28
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if l1 == v7 {
		v28 = int32(0)
		return v28
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l1
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v12 = *(*int32)(unsafe.Add(mBase, _consts[279]))
		if l1 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			v22 = F_aeCreateFileEvent(m, v12, v10, int32(1), v21, l0)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(-1) {
					v28 = int32(-1)
				} else {
					v28 = int32(0)
				}
				return v28
			}
		} else {
			F_aeDeleteFileEvent(m, v12, v10, int32(1))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v28 = int32(0)
				return v28
			}
		}
	}
}
func F_connSocketSetWriteHandler(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if l1 == v5 {
		v37 = int32(0)
		return v37
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
		v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		v15 = v8&int32(65533) | base.B2i32(l2 != int32(0))<<(uint(int32(1))%32)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v15)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v19 = *(*int32)(unsafe.Add(mBase, _consts[279]))
		if l1 != 0 {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
			v29 = F_aeCreateFileEvent(m, v19, v17, int32(2), v28, l0)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if v29 == int32(-1) {
					v37 = int32(-1)
				} else {
					v37 = int32(0)
				}
				return v37
			}
		} else {
			F_aeDeleteFileEvent(m, v19, v17, int32(2))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v37 = int32(0)
				return v37
			}
		}
	}
}
func F_connSocketSyncWrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = F_mstime(m)
	mBase = m.M
	v13 = l1
	v17 = l2
	v18 = l3
	goto L5
L1:
	;
	return v53
L2:
	;
	if v53 != int32(-1) {
		goto L1
	} else {
		goto L16
	}
L3:
	;
	goto L2
L4:
	;
	v53 = int32(-1)
	goto L3
L5:
	;
	v21 = F_write(m, v5, v13, v17)
	mBase = m.M
	if v21 != int32(-1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v44 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(73)
	goto L4
L7:
	;
	if v31 == int32(0) {
		v53 = l2
		goto L3
	} else {
		goto L11
	}
L8:
	;
	v30 = v13 + v21
	v31 = v17 - v21
	goto L7
L9:
	;
	v24 = F___errno_location(m)
	mBase = m.M
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v25 == int32(6) {
		v30 = v13
		v31 = v17
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	v35 = int64(10)
	if v35 < v18 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v38 = v18
	goto L14
L13:
	;
	v38 = v35
	goto L14
L14:
	;
	v39 = F_aeWait(m, v5, int32(2), v38)
	mBase = m.M
	v40 = F_mstime(m)
	mBase = m.M
	v41 = v40 - v11
	if v41 < l3 {
		v13 = v30
		v17 = v31
		v18 = l3 - v41
		goto L5
	} else {
		goto L15
	}
L15:
	;
	goto L6
L16:
	;
	goto L17
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v60
	goto L1
}
func F_connTypeCleanupAll(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _consts[319]))
	if v3 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[322]))
	if v13 == v12 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	m.T0[v6].(func(*base.Module))(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
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
	v22 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[323]))
	if v23 == v22 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v16 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	m.T0[v16].(func(*base.Module))(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	v32 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[324]))
	if v33 == v32 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v26 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	m.T0[v26].(func(*base.Module))(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	return
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	if v36 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	m.T0[v36].(func(*base.Module))(m)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	goto L14
}
func F_connTypeRegister(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = m.T0[v9].(func(*base.Module) int32)(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if base.Ui32(int32(3)) < base.Ui32(v10) {
			F__serverAssert(m, int32(_a537), int32(_a538), int32(34))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v17 = v10 << (uint(int32(2)) % 32)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[319])))
			if v20 != 0 {
				F__serverAssert(m, int32(_a537), int32(_a538), int32(34))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				if int32(1) < v22 {
					*(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[319]))) = l0
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v36 == int32(0) {
						m.G0 = v7 + int32(16)
						return int32(0)
					} else {
						m.T0[v36].(func(*base.Module))(m)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(16)
							return int32(0)
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v10<<(uint(int32(2))%32))+uint32(_consts[320])))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v29
					F__serverLog(m, int32(1), int32(_a539), v7)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[319]))) = l0
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v36 == int32(0) {
							m.G0 = v7 + int32(16)
							return int32(0)
						} else {
							m.T0[v36].(func(*base.Module))(m)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(16)
								return int32(0)
							}
						}
					}
				}
			}
		}
	}
}
func F_connUnixAcceptHandler(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	v11 = m.G0
	v13 = v11 - int32(64)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _consts[954]))
	v20 = v13 + int32(56)
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(2048)
	if v16 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(64)
	return
L2:
	;
	v37 = v16
	goto L3
L3:
	;
	v42 = F_anetUnixAccept(m, int32(_a336), l1)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L1
L5:
	;
	v103 = v37 + int32(-1)
	if v103 != 0 {
		v37 = v103
		goto L3
	} else {
		goto L21
	}
L6:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(1) < v66 {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	return
L8:
	;
	if v42 != int32(-1) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v46 = int32(9116376)
	goto L10
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	goto L11
L11:
	;
	if v47 == int32(13) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v50 == int32(6) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v54 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a336)
	F__serverLog(m, int32(3), int32(_a1756), v13)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L1
L16:
	;
	v79 = F_valkey_calloc(m, int32(40))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L7
	} else {
		goto L19
	}
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[868]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v70
	F__serverLog(m, int32(1), int32(_a1757), v13+int32(32))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v81 = int32(1024)
	*(*uint16)(unsafe.Add(mBase, uint32(v79)+20)) = uint16(v81)
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(_a1755)
	*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = int32(2)
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(24)))) = v92
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v94
	F_acceptCommonHandler(m, v79, v13+int32(16), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	goto L5
L21:
	;
	goto L4
}
func F_connUnixCloseListener(m *base.Module, l0 int32) {
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
		v4 = *(*int32)(unsafe.Add(mBase, uint32(v2)+36))
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
func F_connUnixEventHandler(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	v5 = F_connectionTypeTcp(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
		m.T0[v7].(func(*base.Module, int32, int32, int32, int32))(m, l0, l1, l2, l3)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	}
}
func F_connUnixGetLastError(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3 = F___strerror_l(m, v2, v2)
	mBase = m.M
	return v3
}
func F_connUnixListen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v13 < int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v25 = int32(0)
	goto L3
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v25<<(uint(int32(2))%32))))
	v36 = F_unlink(m, v35)
	mBase = m.M
	v38 = *(*int32)(unsafe.Add(mBase, _consts[955]))
	v39 = F_anetUnixServer(m, int32(_a336), v35, v18, v38, v17)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L1
L5:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v60 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v59 + v60
	*(*int32)(unsafe.Add(mBase, uint32(l0+v59<<(uint(int32(2))%32)))) = v39
	v68 = v25 + v60
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v68 < v69 {
		v25 = v68
		goto L3
	} else {
		goto L12
	}
L6:
	;
	return int32(0)
L7:
	;
	if v39 != int32(-1) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v46 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a336)
	F__serverLog(m, int32(3), int32(_a1758), v11)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	goto L4
}
func F_connUnixShutdown(m *base.Module, l0 int32) {
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
		v4 = *(*int32)(unsafe.Add(mBase, uint32(v2)+48))
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
func F_connUnixSyncWrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = F_mstime(m)
	mBase = m.M
	v13 = l1
	v17 = l2
	v18 = l3
	goto L4
L1:
	;
	return v53
L2:
	;
	goto L1
L3:
	;
	v53 = int32(-1)
	goto L2
L4:
	;
	v21 = F_write(m, v5, v13, v17)
	mBase = m.M
	if v21 != int32(-1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v44 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(73)
	goto L3
L6:
	;
	if v31 == int32(0) {
		v53 = l2
		goto L2
	} else {
		goto L10
	}
L7:
	;
	v30 = v13 + v21
	v31 = v17 - v21
	goto L6
L8:
	;
	v24 = F___errno_location(m)
	mBase = m.M
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v25 == int32(6) {
		v30 = v13
		v31 = v17
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L3
L10:
	;
	v35 = int64(10)
	if v35 < v18 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = v18
	goto L13
L12:
	;
	v38 = v35
	goto L13
L13:
	;
	v39 = F_aeWait(m, v5, int32(2), v38)
	mBase = m.M
	v40 = F_mstime(m)
	mBase = m.M
	v41 = v40 - v11
	if v41 < l3 {
		v13 = v30
		v17 = v31
		v18 = l3 - v41
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L5
}
func F_connUnixWritev(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = F_connectionTypeTcp(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+72))
		v9 = m.T0[v8].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
