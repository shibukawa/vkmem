package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F_VM_ACLCheckChannelPermissions(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v101 int32
	_ = v101
	if base.Ui32(l2) < base.Ui32(int32(16)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if base.Ui32(int32(7)) < base.Ui32(l2) {
		v101 = int32(0)
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_ACLCheckChannelPermissions[0])) = int32(28)
	return int32(1)
L4:
	;
	return v101
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = F_objectGetVal(m, l1)
	mBase = m.M
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	if v15 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v101 = base.B2i32(v88 != int32(0))
	goto L4
L7:
	;
	m.G0 = v30 + int32(16)
	goto L6
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	F_listRewind(m, v33, v30+int32(8))
	mBase = m.M
	goto L10
L9:
	;
	v88 = int32(0)
	goto L7
L10:
	;
	v61 = F_listNext(m, v30+int32(8))
	mBase = m.M
	if v61 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v88 = v71
	goto L7
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v64&int32(8) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v88 = int32(5)
	goto L7
L14:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v63)+144))
	v71 = int32(0)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-1)))))
	switch v73 & int32(7) {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	case 4:
		goto L17
	default:
		v82 = v71
		goto L16
	}
L15:
	;
	v88 = int32(0)
	goto L7
L16:
	;
	v83 = F_ACLCheckChannelAgainstList(m, v70, v16, v82, l2&int32(1))
	mBase = m.M
	if v83 != 0 {
		goto L10
	} else {
		goto L22
	}
L17:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
	v82 = v81
	goto L16
L18:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
	v82 = v80
	goto L16
L19:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
	v82 = v79
	goto L16
L20:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
	v82 = v78
	goto L16
L21:
	;
	v82 = int32(base.Ui32(v73) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	goto L11
}
func F_VM_ACLCheckPermissions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l3 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_VM_ACLCheckPermissions[0])) = int32(28)
		v53 = int32(1)
		m.G0 = v10 + int32(16)
		return v53
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ACLCheckPermissions[1]))
		if l3 < v15 {
			v21 = F_lookupCommand(m, l1, l2)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v21 != 0 {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v32 = F_ACLCheckAllUserCommandPerm(m, v29, v21, l1, l2, l3, v10+int32(12))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						if v32 != 0 {
							*(*int32)(unsafe.Add(mBase, _c_F_VM_ACLCheckPermissions[0])) = int32(2)
							v38 = int32(1)
							if l4 == int32(0) {
								v53 = v38
							} else {
								switch v32 + int32(-1) {
								case 0:
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(4)
									v53 = v38
								default:
									v49 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v49
									v53 = v49
								case 2:
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
									v53 = v38
								case 4:
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(3)
									v53 = v38
								}
							}
						} else {
							v53 = int32(0)
						}
						m.G0 = v10 + int32(16)
						return v53
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_VM_ACLCheckPermissions[0])) = int32(28)
					v53 = int32(1)
					m.G0 = v10 + int32(16)
					return v53
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_VM_ACLCheckPermissions[0])) = int32(28)
			v53 = int32(1)
			m.G0 = v10 + int32(16)
			return v53
		}
	}
}
func F_VM_Alloc(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_zmalloc_usable(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_VM_BlockClientOnAuth(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+108))
	if v7 == int32(0) {
		F_addReplyError(m, v6, int32(_a_F_VM_BlockClientOnAuth_0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		if v10 != 0 {
			v18 = int32(0)
			v25 = F_moduleBlockClient(m, l0, v18, l1, v18, l2, int64(0), v18, v18, v18, v18)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+200)))
				if v28&int32(16) == int32(0) {
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+204))
					*(*int32)(unsafe.Add(mBase, uint32(v27)+204)) = v33 | int32(2)
				}
				return v25
			}
		} else {
			F_addReplyError(m, v6, int32(_a_F_VM_BlockClientOnAuth_0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	}
}
func F_VM_BlockedClientMeasureTimeStart(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_VM_BlockedClientMeasureTimeStart[0]))
	v4 = m.T0[v3].(func(*base.Module) int64)(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v4
	return v2
}
func F_VM_CachedMicroseconds(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, _c_F_VM_CachedMicroseconds[0]))
	return v2
}
func F_VM_CallReplyAttribute(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	return v2
}
func F_VM_CallReplyPromiseAbort(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
	if v6 != 0 {
		F__serverAssert(m, int32(_a_F_VM_CallReplyPromiseAbort_0), int32(_a_F_VM_CallReplyPromiseAbort_1), int32(6258))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v7 = int32(1)
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
		if v8 == int32(0) {
			v33 = v7
			return v33
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+200)))
			if v11&int32(16) == int32(0) {
				v33 = v7
				return v33
			} else {
				if l1 == int32(0) {
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v18
				}
				v20 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v20
				F_unblockClient(m, v8, v20)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
					F_moduleReleaseTempClient(m, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = v20
						return v33
					}
				}
			}
		}
	}
}
func F_VM_CallReplySetElement(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_callReplyGetSetElement(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_VM_CallReplyType(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_callReplyType(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_VM_CallReplyVerbatim(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_callReplyGetVerbatim(m, l0, l1, l2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_VM_CloseKey(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_moduleCloseKey(m, l0)
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+28)))
	if v14&int32(1) == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L20
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	if v19 < int32(1) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v22 = int32(1)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v33 = int32(0)
	goto L9
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = int32(3)
	v59 = v19 + int32(-1)
	if v54 == v59 {
		v69 = v54
		goto L18
	} else {
		goto L19
	}
L9:
	;
	v38 = v19 + (v33 ^ int32(-1))
	v41 = v26 + v38<<(uint(int32(3))%32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v47 = v26 + v33<<(uint(int32(3))%32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v43 == l0 {
		v54 = v38
		v55 = v41
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v52 = v33 + int32(1)
	if v52 != int32(base.Ui32(v19+v22)>>(uint(v22)%32)) {
		v33 = v52
		goto L9
	} else {
		goto L17
	}
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v49 != l0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v54 = v33
	v55 = v47
	goto L8
L17:
	;
	goto L5
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v69
	goto L5
L19:
	;
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v26+v59<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v69 = v66 + int32(-1)
	goto L18
L20:
	;
	goto L1
}
func F_VM_CommandFilterArgsCount(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	return v2
}
func F_VM_CommandFilterGetClientId(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	return v3
}
func F_VM_CreateCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	v10 = int32(1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	if v12 == int32(0) {
		v73 = v10
		return v73
	} else {
		if l3 != 0 {
			v16 = F_commandFlagsFromString(m, l3)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v16 == int64(-1) {
					v73 = v10
					return v73
				} else {
					if v16&int64(4194304) == int64(0) {
						v28 = v16
						v30 = F_strcspn(m, l1, int32(_a_F_VM_CreateCommand_0))
						mBase = m.M
						v31 = l1 + v30
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
						if v33 != 0 {
							v34 = v31
						} else {
							v34 = int32(0)
						}
						if v34 != 0 {
							v73 = v10
							return v73
						} else {
							v35 = F_lookupCommandByCString(m, l1)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								if v35 != 0 {
									v73 = v10
									return v73
								} else {
									v37 = F_sdsnew(m, l1)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return int32(0)
									} else {
										v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v40 = F_sdsdup(m, v37)
										mBase = m.M
										v41 = m.ExcPending
										if v41 != 0 {
											return int32(0)
										} else {
											v42 = F_moduleCreateCommandProxy(m, v39, v37, v40, l2, v28, l4, l5, l6)
											mBase = m.M
											v43 = m.ExcPending
											if v43 != 0 {
												return int32(0)
											} else {
												v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
												if l2 != 0 {
													v47 = int32(-1)
												} else {
													v47 = int32(-2)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v44)+52)) = v47
												F_drainIOThreadsQueue(m)
												mBase = m.M
												v50 = m.ExcPending
												if v50 != 0 {
													return int32(0)
												} else {
													v52 = *(*int32)(unsafe.Add(mBase, _c_F_VM_CreateCommand[0]))
													v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
													v54 = F_hashtableAdd(m, v52, v53)
													mBase = m.M
													v55 = m.ExcPending
													if v55 != 0 {
														return int32(0)
													} else {
														if v54 == int32(0) {
															F__serverAssert(m, int32(_a_F_VM_CreateCommand_1), int32(_a_F_VM_CreateCommand_2), int32(1421))
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return int32(0)
															} else {
																F_abort(m)
																mBase = m.M
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															v59 = *(*int32)(unsafe.Add(mBase, _c_F_VM_CreateCommand[1]))
															v60 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
															v61 = F_hashtableAdd(m, v59, v60)
															mBase = m.M
															v62 = m.ExcPending
															if v62 != 0 {
																return int32(0)
															} else {
																if v61 == int32(0) {
																	F__serverAssert(m, int32(_a_F_VM_CreateCommand_3), int32(_a_F_VM_CreateCommand_2), int32(1422))
																	mBase = m.M
																	v86 = m.ExcPending
																	if v86 != 0 {
																		return int32(0)
																	} else {
																		F_abort(m)
																		mBase = m.M
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																} else {
																	v65 = F_ACLGetCommandID(m, v37)
																	mBase = m.M
																	v66 = m.ExcPending
																	if v66 != 0 {
																		return int32(0)
																	} else {
																		v67 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
																		*(*int32)(unsafe.Add(mBase, uint32(v67)+136)) = v65
																		F_invalidateCommandCache(m)
																		mBase = m.M
																		v70 = m.ExcPending
																		if v70 != 0 {
																			return int32(0)
																		} else {
																			v73 = int32(0)
																			return v73
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
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, _c_F_VM_CreateCommand[2]))
						if v27 != 0 {
							v73 = v10
							return v73
						} else {
							v28 = v16
							v30 = F_strcspn(m, l1, int32(_a_F_VM_CreateCommand_0))
							mBase = m.M
							v31 = l1 + v30
							v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
							if v33 != 0 {
								v34 = v31
							} else {
								v34 = int32(0)
							}
							if v34 != 0 {
								v73 = v10
								return v73
							} else {
								v35 = F_lookupCommandByCString(m, l1)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									if v35 != 0 {
										v73 = v10
										return v73
									} else {
										v37 = F_sdsnew(m, l1)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return int32(0)
										} else {
											v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v40 = F_sdsdup(m, v37)
											mBase = m.M
											v41 = m.ExcPending
											if v41 != 0 {
												return int32(0)
											} else {
												v42 = F_moduleCreateCommandProxy(m, v39, v37, v40, l2, v28, l4, l5, l6)
												mBase = m.M
												v43 = m.ExcPending
												if v43 != 0 {
													return int32(0)
												} else {
													v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
													if l2 != 0 {
														v47 = int32(-1)
													} else {
														v47 = int32(-2)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v44)+52)) = v47
													F_drainIOThreadsQueue(m)
													mBase = m.M
													v50 = m.ExcPending
													if v50 != 0 {
														return int32(0)
													} else {
														v52 = *(*int32)(unsafe.Add(mBase, _c_F_VM_CreateCommand[0]))
														v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
														v54 = F_hashtableAdd(m, v52, v53)
														mBase = m.M
														v55 = m.ExcPending
														if v55 != 0 {
															return int32(0)
														} else {
															if v54 == int32(0) {
																F__serverAssert(m, int32(_a_F_VM_CreateCommand_1), int32(_a_F_VM_CreateCommand_2), int32(1421))
																mBase = m.M
																v80 = m.ExcPending
																if v80 != 0 {
																	return int32(0)
																} else {
																	F_abort(m)
																	mBase = m.M
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																v59 = *(*int32)(unsafe.Add(mBase, _c_F_VM_CreateCommand[1]))
																v60 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
																v61 = F_hashtableAdd(m, v59, v60)
																mBase = m.M
																v62 = m.ExcPending
																if v62 != 0 {
																	return int32(0)
																} else {
																	if v61 == int32(0) {
																		F__serverAssert(m, int32(_a_F_VM_CreateCommand_3), int32(_a_F_VM_CreateCommand_2), int32(1422))
																		mBase = m.M
																		v86 = m.ExcPending
																		if v86 != 0 {
																			return int32(0)
																		} else {
																			F_abort(m)
																			mBase = m.M
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	} else {
																		v65 = F_ACLGetCommandID(m, v37)
																		mBase = m.M
																		v66 = m.ExcPending
																		if v66 != 0 {
																			return int32(0)
																		} else {
																			v67 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
																			*(*int32)(unsafe.Add(mBase, uint32(v67)+136)) = v65
																			F_invalidateCommandCache(m)
																			mBase = m.M
																			v70 = m.ExcPending
																			if v70 != 0 {
																				return int32(0)
																			} else {
																				v73 = int32(0)
																				return v73
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
					}
				}
			}
		} else {
			v28 = int64(0)
			v30 = F_strcspn(m, l1, int32(_a_F_VM_CreateCommand_0))
			mBase = m.M
			v31 = l1 + v30
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
			if v33 != 0 {
				v34 = v31
			} else {
				v34 = int32(0)
			}
			if v34 != 0 {
				v73 = v10
				return v73
			} else {
				v35 = F_lookupCommandByCString(m, l1)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					if v35 != 0 {
						v73 = v10
						return v73
					} else {
						v37 = F_sdsnew(m, l1)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v40 = F_sdsdup(m, v37)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v42 = F_moduleCreateCommandProxy(m, v39, v37, v40, l2, v28, l4, l5, l6)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
									if l2 != 0 {
										v47 = int32(-1)
									} else {
										v47 = int32(-2)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v44)+52)) = v47
									F_drainIOThreadsQueue(m)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, _c_F_VM_CreateCommand[0]))
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
										v54 = F_hashtableAdd(m, v52, v53)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											if v54 == int32(0) {
												F__serverAssert(m, int32(_a_F_VM_CreateCommand_1), int32(_a_F_VM_CreateCommand_2), int32(1421))
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v59 = *(*int32)(unsafe.Add(mBase, _c_F_VM_CreateCommand[1]))
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
												v61 = F_hashtableAdd(m, v59, v60)
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return int32(0)
												} else {
													if v61 == int32(0) {
														F__serverAssert(m, int32(_a_F_VM_CreateCommand_3), int32(_a_F_VM_CreateCommand_2), int32(1422))
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int32(0)
														} else {
															F_abort(m)
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														v65 = F_ACLGetCommandID(m, v37)
														mBase = m.M
														v66 = m.ExcPending
														if v66 != 0 {
															return int32(0)
														} else {
															v67 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v67)+136)) = v65
															F_invalidateCommandCache(m)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return int32(0)
															} else {
																v73 = int32(0)
																return v73
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
	}
}
func F_VM_CreateDataType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int64
	_ = v13
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int64
	_ = v68
	var v72 int64
	_ = v72
	var v113 int64
	_ = v113
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int64
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int64
	_ = v184
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	v5 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
	if v9 == v5 {
		v196 = v5
		return v196
	} else {
		v13 = int64(0)
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_VM_CreateDataType[0]))
		v25 = F_strlen(m, l1)
		mBase = m.M
		if base.Ui32(int32(1023)) < base.Ui32(l2) {
			v113 = v13
		} else {
			if v25 != int32(9) {
				v113 = v13
			} else {
				v30 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
				v31 = F_strchr(m, v23, v30)
				mBase = m.M
				if v31 == int32(0) {
					v113 = v13
				} else {
					v34 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+1)))
					v35 = F_strchr(m, v23, v34)
					mBase = m.M
					if v35 == int32(0) {
						v113 = v13
					} else {
						v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+2)))
						v39 = F_strchr(m, v23, v38)
						mBase = m.M
						if v39 == int32(0) {
							v113 = v13
						} else {
							v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+3)))
							v43 = F_strchr(m, v23, v42)
							mBase = m.M
							if v43 == int32(0) {
								v113 = v13
							} else {
								v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+4)))
								v47 = F_strchr(m, v23, v46)
								mBase = m.M
								if v47 == int32(0) {
									v113 = v13
								} else {
									v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+5)))
									v51 = F_strchr(m, v23, v50)
									mBase = m.M
									if v51 == int32(0) {
										v113 = v13
									} else {
										v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+6)))
										v55 = F_strchr(m, v23, v54)
										mBase = m.M
										if v55 == int32(0) {
											v113 = v13
										} else {
											v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+7)))
											v59 = F_strchr(m, v23, v58)
											mBase = m.M
											if v59 == int32(0) {
												v113 = v13
											} else {
												v62 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+8)))
												v63 = F_strchr(m, v23, v62)
												mBase = m.M
												if v63 == int32(0) {
													v113 = v13
												} else {
													v68 = int64(12)
													v72 = int64(6)
													v113 = ((((base.I64_extend_i32_u(v31-v23)<<(uint(v68)%64)|base.I64_extend_i32_u(v35-v23)<<(uint(v72)%64)|base.I64_extend_i32_u(v39-v23))<<(uint(v68)%64)|base.I64_extend_i32_u(v43-v23)<<(uint(v72)%64)|base.I64_extend_i32_u(v47-v23))<<(uint(v68)%64)|base.I64_extend_i32_u(v51-v23)<<(uint(v72)%64)|base.I64_extend_i32_u(v55-v23))<<(uint(v68)%64)|base.I64_extend_i32_u(v59-v23)<<(uint(v72)%64)|base.I64_extend_i32_u(v63-v23))<<(uint(int64(10))%64) | base.I64_extend_i32_u(l2)
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
		if v113 == int64(0) {
			v196 = v5
			return v196
		} else {
			v124 = int32(0)
			v126 = F_moduleTypeLookupModuleByNameInternal(m, l1, v124)
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return int32(0)
			} else {
				if v126 != 0 {
					v196 = v124
					return v196
				} else {
					v130 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					if v130 == int32(0) {
						v196 = v124
						return v196
					} else {
						v134 = F_valkey_calloc(m, int32(96))
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v134))) = v113
							v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v134)+8)) = v137
							v139 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = v139
							v141 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v134)+16)) = v141
							v143 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v134)+20)) = v143
							v145 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v134)+24)) = v145
							v147 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v134)+28)) = v147
							v149 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v134)+32)) = v149
							v151 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
							if base.Ui64(v151) < base.Ui64(int64(2)) {
							} else {
								v154 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v134)+52)) = v154
								v156 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v134)+56)) = v156
								v158 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v134)+80)) = v158
								if v151 == int64(2) {
								} else {
									v162 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
									*(*int32)(unsafe.Add(mBase, uint32(v134)+36)) = v162
									v164 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v134)+40)) = v164
									v166 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
									*(*int32)(unsafe.Add(mBase, uint32(v134)+44)) = v166
									v168 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
									*(*int32)(unsafe.Add(mBase, uint32(v134)+48)) = v168
									if base.Ui64(v151) < base.Ui64(int64(4)) {
									} else {
										v172 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
										*(*int32)(unsafe.Add(mBase, uint32(v134)+60)) = v172
										v174 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
										*(*int32)(unsafe.Add(mBase, uint32(v134)+68)) = v174
										v176 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v134)+64)) = v176
										v178 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
										*(*int32)(unsafe.Add(mBase, uint32(v134)+72)) = v178
										if v151 == int64(4) {
										} else {
											v182 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
											*(*int32)(unsafe.Add(mBase, uint32(v134)+76)) = v182
										}
									}
								}
							}
							v184 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
							*(*int64)(unsafe.Add(mBase, uint32(v134)+84)) = v184
							v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(8)))))
							*(*uint16)(unsafe.Add(mBase, uint32(v134+int32(92)))) = uint16(v190)
							v192 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
							v193 = F_listAddNodeTail(m, v192, v134)
							mBase = m.M
							v194 = m.ExcPending
							if v194 != 0 {
								return int32(0)
							} else {
								v196 = v134
								return v196
							}
						}
					}
				}
			}
		}
	}
}
func F_VM_CreateString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v5 = F_createStringObject_1(m, l1, l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if l0 == int32(0) {
			return v5
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
			if v11&int32(1) == int32(0) {
				return v5
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v16 == v17 {
					v20 = int32(8)
					if v20 < v16 {
						v23 = v16
					} else {
						v23 = v20
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v23 << (uint(int32(1)) % 32)
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v30 = F_valkey_realloc(m, v27, v23<<(uint(int32(4))%32))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v30
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v34 = v33
						v35 = v30
						v38 = v35 + v34<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v38))) = v5
						v40 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v40
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v34 + v40
						return v5
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v34 = v16
					v35 = v19
					v38 = v35 + v34<<(uint(int32(3))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v38))) = v5
					v40 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v40
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v34 + v40
					return v5
				}
			}
		}
	}
}
func F_VM_CreateStringFromCallReply(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	v6 = m.G0
	v8 = v6 - int32(64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = F_callReplyType(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		switch v12 {
		case 0, 1, 13:
			v16 = F_callReplyGetString(m, l0, v8)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v19 = F_createStringObject_1(m, v16, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v10 == int32(0) {
						v140 = v19
						m.G0 = v8 + int32(64)
						return v140
					} else {
						v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
						if v23&int32(1) == int32(0) {
							v140 = v19
							m.G0 = v8 + int32(64)
							return v140
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
							if v28 == v29 {
								v32 = int32(8)
								if v32 < v28 {
									v35 = v28
								} else {
									v35 = v32
								}
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v35 << (uint(int32(1)) % 32)
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
								v42 = F_valkey_realloc(m, v39, v35<<(uint(int32(4))%32))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v42
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
									v46 = v45
									v47 = v42
									v50 = v47 + v46<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v50))) = v19
									v52 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v52
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v46 + v52
									v140 = v19
									m.G0 = v8 + int32(64)
									return v140
								}
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
								v46 = v28
								v47 = v31
								v50 = v47 + v46<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v50))) = v19
								v52 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v46 + v52
								v140 = v19
								m.G0 = v8 + int32(64)
								return v140
							}
						}
					}
				}
			}
		case 2:
			v58 = F_callReplyGetLongLong(m, l0)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				if v58 <= int64(-1) {
					v68 = int32(45)
					*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v68)
					v72 = int32(1)
					v77 = v8 + v72
					v78 = int32(63)
					v79 = int64(0) - v58
					v80 = v72
				} else {
					v77 = v8
					v78 = int32(64)
					v79 = v58
					v80 = int32(0)
				}
				v81 = F_ull2string(m, v77, v78, v79)
				mBase = m.M
				if v81 == int32(0) {
					v100 = int32(0)
				} else {
					v100 = v81 + v80
				}
				v101 = F_createStringObject_1(m, v8, v100)
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return int32(0)
				} else {
					if v10 == int32(0) {
						v140 = v101
						m.G0 = v8 + int32(64)
						return v140
					} else {
						v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
						if v105&int32(1) == int32(0) {
							v140 = v101
							m.G0 = v8 + int32(64)
							return v140
						} else {
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
							if v110 == v111 {
								v114 = int32(8)
								if v114 < v110 {
									v117 = v110
								} else {
									v117 = v114
								}
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v117 << (uint(int32(1)) % 32)
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
								v124 = F_valkey_realloc(m, v121, v117<<(uint(int32(4))%32))
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v124
									v127 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
									v128 = v127
									v129 = v124
									v132 = v129 + v128<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v132))) = v101
									v134 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v132)+4)) = v134
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v128 + v134
									v140 = v101
									m.G0 = v8 + int32(64)
									return v140
								}
							} else {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
								v128 = v110
								v129 = v113
								v132 = v129 + v128<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v132))) = v101
								v134 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v132)+4)) = v134
								*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v128 + v134
								v140 = v101
								m.G0 = v8 + int32(64)
								return v140
							}
						}
					}
				}
			}
		default:
			v140 = int32(0)
			m.G0 = v8 + int32(64)
			return v140
		}
	}
}
func F_VM_CreateStringFromDouble(m *base.Module, l0 int32, l1 float64) int32 {
	mBase := m.M
	_ = mBase
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
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v7 = m.G0
	v8 = int32(128)
	v9 = v7 - v8
	m.G0 = v9
	v12 = F_d2string(m, v9, v8, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = F_createStringObject_1(m, v9, v12)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if l0 == int32(0) {
				m.G0 = v9 + int32(128)
				return v16
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
				if v20&int32(1) == int32(0) {
					m.G0 = v9 + int32(128)
					return v16
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v25 == v26 {
						v29 = int32(8)
						if v29 < v25 {
							v32 = v25
						} else {
							v32 = v29
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v32 << (uint(int32(1)) % 32)
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v39 = F_valkey_realloc(m, v36, v32<<(uint(int32(4))%32))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v39
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v43 = v42
							v44 = v39
							v47 = v44 + v43<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v47))) = v16
							v49 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v49
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v43 + v49
							m.G0 = v9 + int32(128)
							return v16
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v43 = v25
						v44 = v28
						v47 = v44 + v43<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v47))) = v16
						v49 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v49
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v43 + v49
						m.G0 = v9 + int32(128)
						return v16
					}
				}
			}
		}
	}
}
func F_VM_CreateStringFromLongLong(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l1 <= int64(-1) {
		v20 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v20)
		v24 = int32(1)
		v29 = v9 + v24
		v30 = int32(20)
		v31 = int64(0) - l1
		v32 = v24
	} else {
		v29 = v9
		v30 = int32(21)
		v31 = l1
		v32 = int32(0)
	}
	v33 = F_ull2string(m, v29, v30, v31)
	mBase = m.M
	if v33 == int32(0) {
		v52 = int32(0)
	} else {
		v52 = v33 + v32
	}
	v53 = F_createStringObject_1(m, v9, v52)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		return int32(0)
	} else {
		if l0 == int32(0) {
			m.G0 = v9 + int32(32)
			return v53
		} else {
			v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
			if v59&int32(1) == int32(0) {
				m.G0 = v9 + int32(32)
				return v53
			} else {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v64 == v65 {
					v68 = int32(8)
					if v68 < v64 {
						v71 = v64
					} else {
						v71 = v68
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v71 << (uint(int32(1)) % 32)
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v78 = F_valkey_realloc(m, v75, v71<<(uint(int32(4))%32))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v78
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v82 = v81
						v83 = v78
						v86 = v83 + v82<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v86))) = v53
						v88 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v88
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v82 + v88
						m.G0 = v9 + int32(32)
						return v53
					}
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v82 = v64
					v83 = v67
					v86 = v83 + v82<<(uint(int32(3))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v86))) = v53
					v88 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v88
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v82 + v88
					m.G0 = v9 + int32(32)
					return v53
				}
			}
		}
	}
}
func F_VM_CreateStringFromString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v5 = F_dupStringObject(m, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if l0 == int32(0) {
			return v5
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
			if v11&int32(1) == int32(0) {
				return v5
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v16 == v17 {
					v20 = int32(8)
					if v20 < v16 {
						v23 = v16
					} else {
						v23 = v20
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v23 << (uint(int32(1)) % 32)
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v30 = F_valkey_realloc(m, v27, v23<<(uint(int32(4))%32))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v30
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v34 = v33
						v35 = v30
						v38 = v35 + v34<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v38))) = v5
						v40 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v40
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v34 + v40
						return v5
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v34 = v16
					v35 = v19
					v38 = v35 + v34<<(uint(int32(3))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v38))) = v5
					v40 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v40
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v34 + v40
					return v5
				}
			}
		}
	}
}
func F_VM_CreateSubcommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v10 = int32(1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	if v12 == int32(0) {
		v67 = v10
		return v67
	} else {
		if l3 != 0 {
			v16 = F_commandFlagsFromString(m, l3)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v16 == int64(-1) {
					v67 = v10
					return v67
				} else {
					if v16&int64(4194304) == int64(0) {
						v28 = v16
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+204))
						if v30 != 0 {
							v67 = v10
							return v67
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+208))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
							if v32 != 0 {
								v67 = v10
								return v67
							} else {
								v34 = F_strcspn(m, l1, int32(_a_F_VM_CreateSubcommand_0))
								mBase = m.M
								v35 = l1 + v34
								v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
								if v37 != 0 {
									v38 = v35
								} else {
									v38 = int32(0)
								}
								if v38 != 0 {
									v67 = v10
									return v67
								} else {
									v39 = F_sdsnew(m, l1)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+200))
										if v41 == int32(0) {
											v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+140))
											v53 = F_catSubCommandFullname(m, v52, l1)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v56 = F_moduleCreateCommandProxy(m, v55, v39, v53, l2, v28, l4, l5, l6)
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int32(0)
												} else {
													v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v58)+52)) = int32(-2)
													F_commandAddSubcommand(m, v29, v58)
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return int32(0)
													} else {
														v67 = int32(0)
														return v67
													}
												}
											}
										} else {
											v44 = F_lookupSubcommand(m, v29, v39)
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int32(0)
											} else {
												if v44 == int32(0) {
													v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+140))
													v53 = F_catSubCommandFullname(m, v52, l1)
													mBase = m.M
													v54 = m.ExcPending
													if v54 != 0 {
														return int32(0)
													} else {
														v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														v56 = F_moduleCreateCommandProxy(m, v55, v39, v53, l2, v28, l4, l5, l6)
														mBase = m.M
														v57 = m.ExcPending
														if v57 != 0 {
															return int32(0)
														} else {
															v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v58)+52)) = int32(-2)
															F_commandAddSubcommand(m, v29, v58)
															mBase = m.M
															v62 = m.ExcPending
															if v62 != 0 {
																return int32(0)
															} else {
																v67 = int32(0)
																return v67
															}
														}
													}
												} else {
													F_sdsfree(m, v39)
													mBase = m.M
													v49 = m.ExcPending
													if v49 != 0 {
														return int32(0)
													} else {
														return int32(1)
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, _c_F_VM_CreateSubcommand[0]))
						if v27 != 0 {
							v67 = v10
							return v67
						} else {
							v28 = v16
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+204))
							if v30 != 0 {
								v67 = v10
								return v67
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+208))
								v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
								if v32 != 0 {
									v67 = v10
									return v67
								} else {
									v34 = F_strcspn(m, l1, int32(_a_F_VM_CreateSubcommand_0))
									mBase = m.M
									v35 = l1 + v34
									v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
									if v37 != 0 {
										v38 = v35
									} else {
										v38 = int32(0)
									}
									if v38 != 0 {
										v67 = v10
										return v67
									} else {
										v39 = F_sdsnew(m, l1)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int32(0)
										} else {
											v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+200))
											if v41 == int32(0) {
												v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+140))
												v53 = F_catSubCommandFullname(m, v52, l1)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return int32(0)
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v56 = F_moduleCreateCommandProxy(m, v55, v39, v53, l2, v28, l4, l5, l6)
													mBase = m.M
													v57 = m.ExcPending
													if v57 != 0 {
														return int32(0)
													} else {
														v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v58)+52)) = int32(-2)
														F_commandAddSubcommand(m, v29, v58)
														mBase = m.M
														v62 = m.ExcPending
														if v62 != 0 {
															return int32(0)
														} else {
															v67 = int32(0)
															return v67
														}
													}
												}
											} else {
												v44 = F_lookupSubcommand(m, v29, v39)
												mBase = m.M
												v45 = m.ExcPending
												if v45 != 0 {
													return int32(0)
												} else {
													if v44 == int32(0) {
														v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+140))
														v53 = F_catSubCommandFullname(m, v52, l1)
														mBase = m.M
														v54 = m.ExcPending
														if v54 != 0 {
															return int32(0)
														} else {
															v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															v56 = F_moduleCreateCommandProxy(m, v55, v39, v53, l2, v28, l4, l5, l6)
															mBase = m.M
															v57 = m.ExcPending
															if v57 != 0 {
																return int32(0)
															} else {
																v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v58)+52)) = int32(-2)
																F_commandAddSubcommand(m, v29, v58)
																mBase = m.M
																v62 = m.ExcPending
																if v62 != 0 {
																	return int32(0)
																} else {
																	v67 = int32(0)
																	return v67
																}
															}
														}
													} else {
														F_sdsfree(m, v39)
														mBase = m.M
														v49 = m.ExcPending
														if v49 != 0 {
															return int32(0)
														} else {
															return int32(1)
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
		} else {
			v28 = int64(0)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+204))
			if v30 != 0 {
				v67 = v10
				return v67
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+208))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				if v32 != 0 {
					v67 = v10
					return v67
				} else {
					v34 = F_strcspn(m, l1, int32(_a_F_VM_CreateSubcommand_0))
					mBase = m.M
					v35 = l1 + v34
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
					if v37 != 0 {
						v38 = v35
					} else {
						v38 = int32(0)
					}
					if v38 != 0 {
						v67 = v10
						return v67
					} else {
						v39 = F_sdsnew(m, l1)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+200))
							if v41 == int32(0) {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+140))
								v53 = F_catSubCommandFullname(m, v52, l1)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v56 = F_moduleCreateCommandProxy(m, v55, v39, v53, l2, v28, l4, l5, l6)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v58)+52)) = int32(-2)
										F_commandAddSubcommand(m, v29, v58)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											v67 = int32(0)
											return v67
										}
									}
								}
							} else {
								v44 = F_lookupSubcommand(m, v29, v39)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									if v44 == int32(0) {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+140))
										v53 = F_catSubCommandFullname(m, v52, l1)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int32(0)
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v56 = F_moduleCreateCommandProxy(m, v55, v39, v53, l2, v28, l4, l5, l6)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v58)+52)) = int32(-2)
												F_commandAddSubcommand(m, v29, v58)
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return int32(0)
												} else {
													v67 = int32(0)
													return v67
												}
											}
										}
									} else {
										F_sdsfree(m, v39)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											return int32(1)
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
func F_VM_DeauthenticateAndCloseClient(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v104 int32
	_ = v104
	v8 = m.G0
	v9 = int32(16)
	v10 = v8 - v9
	m.G0 = v10
	v12 = int64(56)
	v14 = int64(65280)
	v16 = int64(40)
	v19 = int64(16711680)
	v21 = int64(24)
	v23 = int64(4278190080)
	v25 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = l1<<(uint(v12)%64) | l1&v14<<(uint(v16)%64) | (l1&v19<<(uint(v21)%64) | l1&v23<<(uint(v25)%64)) | (int64(base.Ui64(l1)>>(uint(v25)%64))&v23 | int64(base.Ui64(l1)>>(uint(v21)%64))&v19 | (int64(base.Ui64(l1)>>(uint(v16)%64))&v14 | int64(base.Ui64(l1)>>(uint(v12)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_VM_DeauthenticateAndCloseClient[0]))
	v52 = int32(8)
	v57 = F_raxFind(m, v51, v10+v52, v52, v10+int32(4))
	mBase = m.M
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	m.G0 = v10 + v9
	if v58 != 0 {
		v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+108))
		if v64 == int32(0) {
		} else {
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
			if v67 == int32(0) {
			} else {
				v70 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
				m.T0[v67].(func(*base.Module, int64, int32))(m, v70, v71)
				mBase = m.M
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v58)+108))
				*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v73)+8)) = int64(0)
			}
		}
		v81 = *(*int32)(unsafe.Add(mBase, _c_F_VM_DeauthenticateAndCloseClient[1]))
		*(*int32)(unsafe.Add(mBase, uint32(v58)+328)) = v81
		v85 = *(*int32)(unsafe.Add(mBase, uint32(v58)+204))
		*(*int32)(unsafe.Add(mBase, uint32(v58)+204)) = v85 & int32(-8388609)
		F_freeClientOrCloseLater(m, v58, int32(1))
		mBase = m.M
		v104 = m.ExcPending
		if v104 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		return int32(1)
	}
}
func F_VM_DefragShouldStop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.B2i32(v3 == int64(0)) == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_VM_DefragShouldStop[0]))
		v12 = m.T0[v11].(func(*base.Module) int64)(m)
		mBase = m.M
		return base.B2i32(base.Ui64(v3) <= base.Ui64(v12))
	} else {
		return int32(0)
	}
}
func F_VM_DictCompareC(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
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
	var v43 int32
	_ = v43
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	v8 = l0 + int32(4)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v9&int32(2) != 0 {
		v72 = int32(1)
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v17 == int32(61) {
			v26 = int32(1)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			if v21 != int32(61) {
				v26 = int32(0)
			} else {
				v26 = int32(1)
			}
		}
		v29 = base.B2i32(v17 == int32(62))
		if v17 == int32(62) {
			v38 = int32(1)
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
			v41 = base.B2i32(base.Ui32(l3) < base.Ui32(v40))
			if base.Ui32(l3) < base.Ui32(v40) {
				v42 = l3
			} else {
				v42 = v40
			}
			v43 = F_memcmp(m, v39, l2, v42)
			mBase = m.M
			if v17 == int32(62) {
				if v43 != 0 {
					if v43 < int32(1) {
						v69 = v38 ^ int32(1)
					} else {
						v69 = base.B2i32(v17 == int32(62))
					}
				} else {
					if v26&base.B2i32(l3 == v40) == int32(0) {
						if v38 != 0 {
							v69 = v41 & base.B2i32(v17 == int32(62))
						} else {
							v69 = base.B2i32(base.Ui32(v40) < base.Ui32(l3))
						}
					} else {
						v69 = int32(1)
					}
				}
			} else {
				if v38 == int32(0) {
					if v43 != 0 {
						if v43 < int32(1) {
							v69 = v38 ^ int32(1)
						} else {
							v69 = base.B2i32(v17 == int32(62))
						}
					} else {
						if v26&base.B2i32(l3 == v40) == int32(0) {
							if v38 != 0 {
								v69 = v41 & base.B2i32(v17 == int32(62))
							} else {
								v69 = base.B2i32(base.Ui32(v40) < base.Ui32(l3))
							}
						} else {
							v69 = int32(1)
						}
					}
				} else {
					v69 = base.B2i32(v43 == int32(0)) & base.B2i32(l3 == v40)
				}
			}
		} else {
			if v17 == int32(60) {
				v38 = int32(0)
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
				v41 = base.B2i32(base.Ui32(l3) < base.Ui32(v40))
				if base.Ui32(l3) < base.Ui32(v40) {
					v42 = l3
				} else {
					v42 = v40
				}
				v43 = F_memcmp(m, v39, l2, v42)
				mBase = m.M
				if v17 == int32(62) {
					if v43 != 0 {
						if v43 < int32(1) {
							v69 = v38 ^ int32(1)
						} else {
							v69 = base.B2i32(v17 == int32(62))
						}
					} else {
						if v26&base.B2i32(l3 == v40) == int32(0) {
							if v38 != 0 {
								v69 = v41 & base.B2i32(v17 == int32(62))
							} else {
								v69 = base.B2i32(base.Ui32(v40) < base.Ui32(l3))
							}
						} else {
							v69 = int32(1)
						}
					}
				} else {
					if v38 == int32(0) {
						if v43 != 0 {
							if v43 < int32(1) {
								v69 = v38 ^ int32(1)
							} else {
								v69 = base.B2i32(v17 == int32(62))
							}
						} else {
							if v26&base.B2i32(l3 == v40) == int32(0) {
								if v38 != 0 {
									v69 = v41 & base.B2i32(v17 == int32(62))
								} else {
									v69 = base.B2i32(base.Ui32(v40) < base.Ui32(l3))
								}
							} else {
								v69 = int32(1)
							}
						}
					} else {
						v69 = base.B2i32(v43 == int32(0)) & base.B2i32(l3 == v40)
					}
				}
			} else {
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
				if v34 == int32(61) {
					v38 = int32(1)
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
					v41 = base.B2i32(base.Ui32(l3) < base.Ui32(v40))
					if base.Ui32(l3) < base.Ui32(v40) {
						v42 = l3
					} else {
						v42 = v40
					}
					v43 = F_memcmp(m, v39, l2, v42)
					mBase = m.M
					if v17 == int32(62) {
						if v43 != 0 {
							if v43 < int32(1) {
								v69 = v38 ^ int32(1)
							} else {
								v69 = base.B2i32(v17 == int32(62))
							}
						} else {
							if v26&base.B2i32(l3 == v40) == int32(0) {
								if v38 != 0 {
									v69 = v41 & base.B2i32(v17 == int32(62))
								} else {
									v69 = base.B2i32(base.Ui32(v40) < base.Ui32(l3))
								}
							} else {
								v69 = int32(1)
							}
						}
					} else {
						if v38 == int32(0) {
							if v43 != 0 {
								if v43 < int32(1) {
									v69 = v38 ^ int32(1)
								} else {
									v69 = base.B2i32(v17 == int32(62))
								}
							} else {
								if v26&base.B2i32(l3 == v40) == int32(0) {
									if v38 != 0 {
										v69 = v41 & base.B2i32(v17 == int32(62))
									} else {
										v69 = base.B2i32(base.Ui32(v40) < base.Ui32(l3))
									}
								} else {
									v69 = int32(1)
								}
							}
						} else {
							v69 = base.B2i32(v43 == int32(0)) & base.B2i32(l3 == v40)
						}
					}
				} else {
					v69 = int32(0)
				}
			}
		}
		v72 = base.B2i32(v69 == int32(0))
	}
	return v72
}
func F_VM_DictDel(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	v7 = F_objectGetVal(m, l1)
	mBase = m.M
	v9 = F_objectGetVal(m, l1)
	mBase = m.M
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-1)))))
	switch v12 & int32(7) {
	case 0:
		v29 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
	case 1:
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-3)))))
		v29 = v19
	case 2:
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9+int32(-5)))))
		v29 = v22
	case 3:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(-9))))
		v29 = v25
	case 4:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(-17))))
		v29 = v28
	default:
		v29 = int32(0)
	}
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = F_raxRemove(m, v30, v7, v29, l2)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v31 == int32(0))
	}
}
func F_VM_DictIteratorReseek(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v7 = F_objectGetVal(m, l2)
	mBase = m.M
	v9 = F_objectGetVal(m, l2)
	mBase = m.M
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-1)))))
	switch v12 & int32(7) {
	case 0:
		v29 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
	case 1:
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-3)))))
		v29 = v19
	case 2:
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9+int32(-5)))))
		v29 = v22
	case 3:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(-9))))
		v29 = v25
	case 4:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(-17))))
		v29 = v28
	default:
		v29 = int32(0)
	}
	v32 = F_raxSeek(m, l0+int32(4), l1, v7, v29)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		return v32
	}
}
func F_VM_DictIteratorReseekC(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_raxSeek(m, l0+int32(4), l1, l2, l3)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_VM_DictIteratorStartC(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int64
	_ = v21
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v8 = F_valkey_malloc(m, int32(308))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
		v14 = v8 + int32(4)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(2)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(128)
		v21 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+12)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v14)+296)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v14)+160)) = int64(137438953472)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v8 + int32(28)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+156)) = v8 + int32(172)
		v33 = F_raxSeek(m, v14, l1, l2, l3)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_VM_DictIteratorStop(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	F_raxStop(m, l0+int32(4))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_valkey_free(m, l0)
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	}
}
func F_VM_DictNext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	v4 = int32(0)
	v9 = F_raxNext(m, l1+int32(4))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			v63 = v4
			return v63
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			if l2 == int32(0) {
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
			}
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			if v20 == int32(0) {
				v63 = v4
				return v63
			} else {
				v23 = F_createStringObject_1(m, v20, v15)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if l0 == int32(0) {
						v63 = v23
						return v63
					} else {
						v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v27&int32(1) == int32(0) {
							v63 = v23
							return v63
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v32 == v33 {
								v36 = int32(8)
								if v36 < v32 {
									v39 = v32
								} else {
									v39 = v36
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v39 << (uint(int32(1)) % 32)
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v46 = F_valkey_realloc(m, v43, v39<<(uint(int32(4))%32))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v46
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v50 = v49
									v51 = v46
									v54 = v51 + v50<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v54))) = v23
									v56 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v56
									*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v50 + v56
									v63 = v23
									return v63
								}
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v50 = v32
								v51 = v35
								v54 = v51 + v50<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v54))) = v23
								v56 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v56
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v50 + v56
								v63 = v23
								return v63
							}
						}
					}
				}
			}
		}
	}
}
func F_VM_DictPrev(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	v4 = int32(0)
	v9 = F_raxPrev(m, l1+int32(4))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			v63 = v4
			return v63
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			if l2 == int32(0) {
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
			}
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			if v20 == int32(0) {
				v63 = v4
				return v63
			} else {
				v23 = F_createStringObject_1(m, v20, v15)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if l0 == int32(0) {
						v63 = v23
						return v63
					} else {
						v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v27&int32(1) == int32(0) {
							v63 = v23
							return v63
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v32 == v33 {
								v36 = int32(8)
								if v36 < v32 {
									v39 = v32
								} else {
									v39 = v36
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v39 << (uint(int32(1)) % 32)
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v46 = F_valkey_realloc(m, v43, v39<<(uint(int32(4))%32))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v46
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v50 = v49
									v51 = v46
									v54 = v51 + v50<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v54))) = v23
									v56 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v56
									*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v50 + v56
									v63 = v23
									return v63
								}
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v50 = v32
								v51 = v35
								v54 = v51 + v50<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v54))) = v23
								v56 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v56
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v50 + v56
								v63 = v23
								return v63
							}
						}
					}
				}
			}
		}
	}
}
func F_VM_DigestAddStringBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v5 = m.G0
	v6 = int32(96)
	v7 = v5 - v6
	m.G0 = v7
	F_xorDigest(m, l0, l1, l2)
	v11 = v7 + int32(4)
	F_SHA1Init(m, v11)
	F_SHA1Update(m, v11, l0, int32(20))
	F_SHA1Final(m, l0, v11)
	m.G0 = v7 + v6
	return
}
func F_VM_EmitAOF(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(48)
	return
L2:
	;
	v12 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v12
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v16 == v12 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l3
	v53 = F_moduleCreateArgvFromUserFormat(m, l1, l2, v9+int32(44), v9+int32(40), l3)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L8
	} else {
		goto L16
	}
L4:
	;
	v26 = F_lookupCommandByCString(m, l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	if v22&int32(16) != 0 {
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
	if v26 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_VM_EmitAOF[0]))
	if int32(3) < v29 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
	goto L14
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v32 + int32(84)
	F__serverLog(m, int32(3), int32(_a_F_VM_EmitAOF_0), v9)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_EmitAOF[1])) = int32(28)
	goto L1
L15:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v76 != 0 {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	if v53 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_VM_EmitAOF[0]))
	if int32(3) < v56 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
	goto L21
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v59 + int32(84)
	F__serverLog(m, int32(3), int32(_a_F_VM_EmitAOF_1), v9+int32(16))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_EmitAOF[1])) = int32(28)
	goto L1
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	if v84 < int32(1) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	v80 = F_rioWriteBulkCount(m, v77, int32(42), v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	if v80 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
	goto L22
L26:
	;
	F_valkey_free(m, v53)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L8
	} else {
		goto L36
	}
L27:
	;
	v89 = int32(0)
	goto L28
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v94 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L26
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v53+v89<<(uint(int32(2))%32))))
	F_decrRefCount(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L34
	}
L31:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v53+v89<<(uint(int32(2))%32))))
	v100 = F_rioWriteBulkObject(m, v95, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	if v100 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
	goto L30
L34:
	;
	v111 = v89 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	if v111 < v112 {
		v89 = v111
		goto L28
	} else {
		goto L35
	}
L35:
	;
	goto L29
L36:
	;
	goto L1
}
func F_VM_ExitFromChild(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_sendChildCowInfo(m, int32(3), int32(_a_F_VM_ExitFromChild_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F__exit(m, l0)
		base.Wasm_trap_unreachable()
		for {
		}
	}
}
func F_VM_Free(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_valkey_free(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_VM_FreeCallReply(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int64
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	v9 = F_callReplyType(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L3
L3:
	;
	if v9 == int32(12) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = int32(3)
	v98 = v34 + int32(-1)
	if v94 == v98 {
		v108 = v94
		goto L30
	} else {
		goto L31
	}
L5:
	;
	F__serverAssert(m, int32(_a_F_VM_FreeCallReply_0), int32(_a_F_VM_FreeCallReply_1), int32(732))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L29
	}
L6:
	;
	F__serverAssert(m, int32(_a_F_VM_FreeCallReply_2), int32(_a_F_VM_FreeCallReply_1), int32(6115))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L28
	}
L7:
	;
	F_freeCallReply(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L14
	}
L8:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if v14 != 0 {
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v24 = v11
	goto L7
L10:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v17 = v15 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if v17 != 0 {
		v24 = v19
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v20 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	F_valkey_free(m, v11)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v24 = v19
	goto L7
L14:
	;
	if v24 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+28)))
	if v29&int32(1) == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	if v34 < int32(1) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v37 = int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v45 = int32(0)
	goto L19
L19:
	;
	v53 = v34 + (v45 ^ int32(-1))
	v56 = v41 + v53<<(uint(int32(3))%32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v57 != int32(2) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L15
L21:
	;
	v64 = v41 + v45<<(uint(int32(3))%32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v65 != int32(2) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v60 == l0 {
		v93 = v56
		v94 = v53
		goto L4
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v71 = v45 + int32(1)
	if v71 != int32(base.Ui32(v34+v37)>>(uint(v37)%32)) {
		v45 = v71
		goto L19
	} else {
		goto L27
	}
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v68 != l0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v93 = v64
	v94 = v45
	goto L4
L27:
	;
	goto L20
L28:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v108
	return
L31:
	;
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v41+v98<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v93))) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v108 = v105 + int32(-1)
	goto L30
}
func F_VM_FreeServerInfo(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_raxFreeWithCallback(m, v80, int32(3))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v11&int32(1) == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v16 < int32(1) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v30 = int32(0)
	goto L6
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = int32(3)
	v60 = v16 + int32(-1)
	if v55 == v60 {
		v70 = v55
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v35 = v16 + (v30 ^ int32(-1))
	v38 = v23 + v35<<(uint(int32(3))%32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v39 != int32(5) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v46 = v23 + v30<<(uint(int32(3))%32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v47 != int32(5) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v42 == l1 {
		v55 = v35
		v56 = v38
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v53 = v30 + int32(1)
	if v53 != int32(base.Ui32(v16+v19)>>(uint(v19)%32)) {
		v30 = v53
		goto L6
	} else {
		goto L14
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v50 != l1 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v55 = v30
	v56 = v46
	goto L5
L14:
	;
	goto L1
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v70
	goto L1
L16:
	;
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v23+v60<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v56))) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v70 = v67 + int32(-1)
	goto L15
L17:
	;
	return
L18:
	;
	F_valkey_free(m, l1)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	return
}
func F_VM_FreeThreadSafeContext(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	F_moduleFreeContext(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		F_valkey_free(m, l0)
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			return
		}
	}
}
func F_VM_GetApi(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetApi[0]))
	v5 = F_dictFind(m, v4, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v11
			return int32(0)
		} else {
			return int32(1)
		}
	}
}
func F_VM_GetBlockedClientPrivateData(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	return v2
}
func F_VM_GetClientCertificate(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v115 int32
	_ = v115
	v3 = int32(0)
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	v13 = int64(56)
	v15 = int64(65280)
	v17 = int64(40)
	v20 = int64(16711680)
	v22 = int64(24)
	v24 = int64(4278190080)
	v26 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = l1<<(uint(v13)%64) | l1&v15<<(uint(v17)%64) | (l1&v20<<(uint(v22)%64) | l1&v24<<(uint(v26)%64)) | (int64(base.Ui64(l1)>>(uint(v26)%64))&v24 | int64(base.Ui64(l1)>>(uint(v22)%64))&v20 | (int64(base.Ui64(l1)>>(uint(v17)%64))&v15 | int64(base.Ui64(l1)>>(uint(v13)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v3
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetClientCertificate[0]))
	v53 = int32(8)
	v58 = F_raxFind(m, v52, v11+v53, v53, v11+int32(4))
	mBase = m.M
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	m.G0 = v11 + v10
	if v59 == int32(0) {
		v115 = v3
		return v115
	} else {
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
		v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
		v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+124))
		if v67 == int32(0) {
			v115 = v3
			return v115
		} else {
			v70 = m.T0[v67].(func(*base.Module, int32) int32)(m, v65)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				if v70 == int32(0) {
					v115 = v3
					return v115
				} else {
					v77 = F_createObject(m, int32(0), v70)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						if l0 == int32(0) {
							v115 = v77
							return v115
						} else {
							v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
							if v81&int32(1) == int32(0) {
								v115 = v77
								return v115
							} else {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v86 == v87 {
									v90 = int32(8)
									if v90 < v86 {
										v93 = v86
									} else {
										v93 = v90
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v93 << (uint(int32(1)) % 32)
									v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v100 = F_valkey_realloc(m, v97, v93<<(uint(int32(4))%32))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v100
										v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v104 = v103
										v105 = v100
										v108 = v105 + v104<<(uint(int32(3))%32)
										*(*int32)(unsafe.Add(mBase, uint32(v108))) = v77
										v110 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v110
										*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v104 + v110
										v115 = v77
										return v115
									}
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v104 = v86
									v105 = v89
									v108 = v105 + v104<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v108))) = v77
									v110 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v110
									*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v104 + v110
									v115 = v77
									return v115
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_VM_GetClientId(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2 != 0 {
		v5 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
		return v5
	} else {
		return int64(0)
	}
}
func F_VM_GetClusterNodeInfoForClient(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v11 = m.G0
	v12 = int32(16)
	v13 = v11 - v12
	m.G0 = v13
	v15 = int64(56)
	v17 = int64(65280)
	v19 = int64(40)
	v22 = int64(16711680)
	v24 = int64(24)
	v26 = int64(4278190080)
	v28 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = l1<<(uint(v15)%64) | l1&v17<<(uint(v19)%64) | (l1&v22<<(uint(v24)%64) | l1&v26<<(uint(v28)%64)) | (int64(base.Ui64(l1)>>(uint(v28)%64))&v26 | int64(base.Ui64(l1)>>(uint(v24)%64))&v22 | (int64(base.Ui64(l1)>>(uint(v19)%64))&v17 | int64(base.Ui64(l1)>>(uint(v15)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetClusterNodeInfoForClient[0]))
	v55 = int32(8)
	v60 = F_raxFind(m, v54, v13+v55, v55, v13+int32(4))
	mBase = m.M
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	m.G0 = v13 + v12
	if v61 != 0 {
		v67 = F_moduleGetClusterNodeInfoForClient(m, l6, v61, l2, l3, l4, l5, l6)
		mBase = m.M
		v70 = m.ExcPending
		if v70 != 0 {
			return int32(0)
		} else {
			return v67
		}
	} else {
		return int32(1)
	}
}
func F_VM_GetContextFromIO(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int64
	_ = v61
	var v65 int64
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v102 int64
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v5 != 0 {
		v107 = v5
		return v107
	} else {
		v7 = F_valkey_malloc(m, int32(72))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			v16 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7+int32(64)))) = v16
			*(*int64)(unsafe.Add(mBase, uint32(v7+int32(56)))) = v16
			*(*int64)(unsafe.Add(mBase, uint32(v7+int32(48)))) = v16
			*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v16
			*(*int64)(unsafe.Add(mBase, uint32(v7+int32(32)))) = v16
			*(*int64)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v16
			*(*int64)(unsafe.Add(mBase, uint32(v7+int32(16)))) = v16
			*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v13
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(561)
			v48 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetContextFromIO[0]))
			v49 = int32(0)
			v50 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetContextFromIO[1]))
			v51 = m.T0[v50].(func(*base.Module) int64)(m)
			mBase = m.M
			if v48 == v49 {
				v61 = *(*int64)(unsafe.Add(mBase, _c_F_VM_GetContextFromIO[2]))
				v65 = v61*int64(1000) + v51
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetContextFromIO[3]))
				v57 = base.I32_div_s(int32(1000000), v56)
				v65 = v51 + base.I64_extend_i32_s(v57)
			}
			*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = v65
			v69 = int32(0)
			v73 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetContextFromIO[4]))
			*(*int32)(unsafe.Add(mBase, _c_F_VM_GetContextFromIO[4])) = v73 + int32(1)
			if v73 != 0 {
			} else {
				v81 = F_ustime(m)
				mBase = m.M
				v83 = int32(0)
				*(*int64)(unsafe.Add(mBase, _c_F_VM_GetContextFromIO[5])) = v81
				v87 = base.I64_div_s(v81, int64(1000))
				*(*int64)(unsafe.Add(mBase, _c_F_VM_GetContextFromIO[6])) = v87
				v91 = base.I64_div_s(v81, int64(1000000))
				*(*int64)(unsafe.Add(mBase, _c_F_VM_GetContextFromIO[7])) = v91
				v94 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetContextFromIO[8]))
				F_lrulfu_updateClockAndPolicy(m, v87, int32(base.Ui32(v94&int32(2))>>(uint(int32(1))%32)))
				mBase = m.M
				v102 = *(*int64)(unsafe.Add(mBase, _c_F_VM_GetContextFromIO[6]))
				*(*int64)(unsafe.Add(mBase, _c_F_VM_GetContextFromIO[9])) = v102
			}
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v107 = v106
			return v107
		}
	}
}
func F_VM_GetCurrentUserName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_VM_GetCurrentUserName[0])) = int32(28)
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v7 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_VM_GetCurrentUserName[0])) = int32(28)
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+328))
			if v10 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _c_F_VM_GetCurrentUserName[0])) = int32(28)
				return int32(0)
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				if v13 != 0 {
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-1)))))
					switch v23 & int32(7) {
					case 0:
						v40 = int32(base.Ui32(v23) >> (uint(int32(3)) % 32))
					case 1:
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
						v40 = v30
					case 2:
						v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
						v40 = v33
					case 3:
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
						v40 = v36
					case 4:
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
						v40 = v39
					default:
						v40 = int32(0)
					}
					v41 = F_createStringObject_1(m, v13, v40)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v45&int32(1) == int32(0) {
							return v41
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v50 == v51 {
								v54 = int32(8)
								if v54 < v50 {
									v57 = v50
								} else {
									v57 = v54
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v57 << (uint(int32(1)) % 32)
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v64 = F_valkey_realloc(m, v61, v57<<(uint(int32(4))%32))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v64
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v68 = v67
									v69 = v64
									v72 = v69 + v68<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v72))) = v41
									v74 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v74
									*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v68 + v74
									return v41
								}
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v68 = v50
								v69 = v53
								v72 = v69 + v68<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v72))) = v41
								v74 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v74
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v68 + v74
								return v41
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_VM_GetCurrentUserName[0])) = int32(28)
					return int32(0)
				}
			}
		}
	}
}
func F_VM_GetDbIdFromDigest(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	return v2
}
func F_VM_GetDbIdFromIO(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	return v2
}
func F_VM_GetDbIdFromOptCtx(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	return v2
}
func F_VM_GetLFU(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(-1)
	v6 = int32(1)
	if l0 == int32(0) {
		v41 = v6
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v9 == int32(0) {
			v41 = v6
		} else {
			v12 = int32(0)
			v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_VM_GetLFU[0])))
			if v14 == int32(0) {
				v41 = v12
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v20 = m.G0
				v21 = int32(16)
				v22 = v20 - v21
				m.G0 = v22
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v25 = int32(8)
				v29 = F_lfu_getFrequency(m, int32(base.Ui32(v24)>>(uint(v25)%32)), v22+int32(15))
				mBase = m.M
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = v30 | v29<<(uint(v25)%32)
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+15)))
				m.G0 = v22 + v21
				*(*int64)(unsafe.Add(mBase, uint32(l1))) = base.I64_extend_i32_u(v35)
				v41 = v12
			}
		}
	}
	return v41
}
func F_VM_GetModuleUserFromUserName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	v2 = int32(0)
	v6 = F_objectGetVal(m, l0)
	mBase = m.M
	v9 = F_objectGetVal(m, l0)
	mBase = m.M
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-1)))))
	switch v12 & int32(7) {
	case 0:
		v29 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
	case 1:
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-3)))))
		v29 = v19
	case 2:
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9+int32(-5)))))
		v29 = v22
	case 3:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(-9))))
		v29 = v25
	case 4:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(-17))))
		v29 = v28
	default:
		v29 = v2
	}
	v30 = int32(0)
	v31 = m.G0
	v32 = int32(16)
	v33 = v31 - v32
	m.G0 = v33
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v30
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetModuleUserFromUserName[0]))
	v41 = F_raxFind(m, v38, v6, v29, v33+int32(12))
	mBase = m.M
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	m.G0 = v33 + v32
	if v42 == int32(0) {
		v56 = v2
		return v56
	} else {
		v49 = F_valkey_malloc(m, int32(8))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v49))) = v42
			v56 = v49
			return v56
		}
	}
}
func F_VM_GetMyClusterID(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetMyClusterID[0]))
	if v2 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetMyClusterID[1]))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		return v7 + int32(8)
	} else {
		return int32(0)
	}
}
func F_VM_GetRandomBytes(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	F_getRandomBytes(m, l0, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_VM_GetSharedAPI(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetSharedAPI[0]))
	v5 = F_dictFind(m, v4, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v15 = F_listSearchKey(m, v13, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v15 != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					return v27
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v20 = F_listAddNodeTail(m, v18, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						v25 = F_listAddNodeTail(m, v23, v24)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							return v27
						}
					}
				}
			}
		} else {
			return int32(0)
		}
	}
}
func F_VM_GetToKeyNameFromOptCtx(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	return v2
}
func F_VM_InfoAddFieldCString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v11 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v14 == int32(0) {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v27
			v32 = F_sdscatfmt(m, v13, int32(_a_F_VM_InfoAddFieldCString_0), v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = v32
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34
				v38 = int32(0)
				m.G0 = v9 + int32(32)
				return v38
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
			v22 = F_sdscatfmt(m, v13, int32(_a_F_VM_InfoAddFieldCString_1), v9+int32(16))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v34 = v22
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34
				v38 = int32(0)
				m.G0 = v9 + int32(32)
				return v38
			}
		}
	} else {
		v38 = int32(1)
		m.G0 = v9 + int32(32)
		return v38
	}
}
func F_VM_InfoAddFieldDouble(m *base.Module, l0 int32, l1 int32, l2 float64) int32 {
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
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v11 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v14 == int32(0) {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
			*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v27
			v32 = F_sdscatprintf(m, v13, int32(_a_F_VM_InfoAddFieldDouble_0), v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = v32
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34
				v38 = int32(0)
				m.G0 = v9 + int32(32)
				return v38
			}
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v9)+24)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
			v22 = F_sdscatprintf(m, v13, int32(_a_F_VM_InfoAddFieldDouble_1), v9+int32(16))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v34 = v22
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34
				v38 = int32(0)
				m.G0 = v9 + int32(32)
				return v38
			}
		}
	} else {
		v38 = int32(1)
		m.G0 = v9 + int32(32)
		return v38
	}
}
func F_VM_InfoAddSection(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = F_sdsdup(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l1 == int32(0) {
		v28 = v13
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v29 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v19 == int32(0) {
		v28 = v13
		goto L3
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	v26 = F_sdscatfmt(m, v13, int32(_a_F_VM_InfoAddSection_2), v9+int32(16))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v28 = v26
	goto L3
L7:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v74 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+int32(-1)))))
	switch v36 & int32(7) {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	case 4:
		goto L10
	default:
		v53 = int32(0)
		goto L9
	}
L9:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v53+int32(-1)))))
	if v57 != int32(44) {
		v64 = v33
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(-17))))
	v53 = v52
	goto L9
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(-9))))
	v53 = v49
	goto L9
L12:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33+int32(-5)))))
	v53 = v46
	goto L9
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+int32(-3)))))
	v53 = v43
	goto L9
L14:
	;
	v53 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
	goto L9
L15:
	;
	v66 = F_sdscat(m, v64, int32(_a_F_VM_InfoAddSection_1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	F_sdsIncrLen(m, v33, int32(-1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v64 = v63
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v66
	goto L7
L19:
	;
	m.G0 = v9 + int32(32)
	return v115
L20:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v93 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v93 == int32(0) {
		v104 = v97
		goto L29
	} else {
		goto L30
	}
L21:
	;
	if v28 == int32(0) {
		v82 = v74
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v85 = F_dictFind(m, v82, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L26
	}
L23:
	;
	v79 = F_dictFind(m, v74, v28)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v79 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v82 = v81
	goto L22
L26:
	;
	if v85 != 0 {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	F_sdsfree(m, v28)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	v115 = int32(1)
	goto L19
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v28
	v107 = F_sdscatfmt(m, v104, int32(_a_F_VM_InfoAddSection_0), v9)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	v101 = F_sdscat(m, v97, int32(_a_F_VM_InfoAddSection_1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v101
	v104 = v101
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v107
	F_sdsfree(m, v28)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v115 = int32(0)
	goto L19
}
func F_VM_InfoEndDictField(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v5 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-1)))))
		switch v12 & int32(7) {
		case 0:
			v29 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
		case 1:
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-3)))))
			v29 = v19
		case 2:
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9+int32(-5)))))
			v29 = v22
		case 3:
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(-9))))
			v29 = v25
		case 4:
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(-17))))
			v29 = v28
		default:
			v29 = int32(0)
		}
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v29+int32(-1)))))
		if v33 != int32(44) {
			v42 = v9
			v44 = F_sdscat(m, v42, int32(_a_F_VM_InfoEndDictField_0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				v46 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v46
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v44
				return v46
			}
		} else {
			F_sdsIncrLen(m, v9, int32(-1))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v42 = v41
				v44 = F_sdscat(m, v42, int32(_a_F_VM_InfoEndDictField_0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v46 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v44
					return v46
				}
			}
		}
	} else {
		return int32(1)
	}
}
func F_VM_IsIOError(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	return v2
}
func F_VM_IsKeysPositionRequest(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(1)
	return int32(base.Ui32(v2)>>(uint(v3)%32)) & v3
}
func F_VM_IsModuleNameBusy(m *base.Module, l0 int32) int32 {
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = F_sdsnew(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_VM_IsModuleNameBusy[0]))
		v9 = F_dictFind(m, v8, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_sdsfree(m, v3)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v9 != int32(0))
			}
		}
	}
}
func F_VM_IsSubEventSupported(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int64
	_ = v6
	var v69 int32
	_ = v69
	v3 = int32(0)
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui64(int64(23)) < base.Ui64(v6) {
		v69 = v3
		return v69
	} else {
		switch base.I32_wrap_i64(v6) {
		default:
			return base.B2i32(l1 < int64(2))
		case 1:
			return base.B2i32(l1 < int64(6))
		case 2:
			return base.B2i32(l1 < int64(2))
		case 3:
			return base.B2i32(l1 < int64(5))
		case 4:
			return base.B2i32(l1 < int64(2))
		case 5:
			return base.I32_wrap_i64(int64(base.Ui64(l1) >> (uint(int64(63)) % 64)))
		case 6:
			return base.B2i32(l1 < int64(2))
		case 7:
			return base.B2i32(l1 < int64(2))
		case 8:
			return base.I32_wrap_i64(int64(base.Ui64(l1) >> (uint(int64(63)) % 64)))
		case 9:
			return base.B2i32(l1 < int64(2))
		case 10:
			return base.B2i32(l1 < int64(2))
		case 11:
			return base.I32_wrap_i64(int64(base.Ui64(l1) >> (uint(int64(63)) % 64)))
		case 12, 18, 19, 20, 21:
			v69 = v3
			return v69
		case 13:
			return base.B2i32(l1 < int64(2))
		case 14:
			return base.B2i32(l1 < int64(3))
		case 15:
			return base.B2i32(l1 < int64(2))
		case 16:
			return base.B2i32(l1 < int64(1))
		case 17:
			return base.B2i32(l1 < int64(4))
		case 22:
			return base.B2i32(l1 == int64(0))
		case 23:
			v69 = base.B2i32(l1 < int64(5))
			return v69
		}
	}
}
func F_VM_KeyAtPosWithFlags(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v44 int64
	_ = v44
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v74 int64
	_ = v74
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v102 int64
	_ = v102
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v115 int64
	_ = v115
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v130 int64
	_ = v130
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v158 int64
	_ = v158
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
	var v186 int64
	_ = v186
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v6&int32(2) == int32(0) {
		return
	} else {
		if l1 < int32(1) {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v13 == int32(0) {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				if v16 != v17 {
					v27 = v16
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					v31 = v28 + v27<<(uint(int32(3))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = l1
					v33 = base.I64_extend_i32_s(l2)
					v44 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[0]))
					if base.B2i32(v44&v33 == int64(0)) == int32(0) {
						v55 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[1]))
						v56 = v55
					} else {
						v56 = int64(0)
					}
					v59 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[2]))
					if v59&v33 == int64(0) {
						v69 = v56
					} else {
						v67 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[3]))
						v69 = v67 | v56
					}
					v74 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[4]))
					if v74&v33 == int64(0) {
						v84 = v69
					} else {
						v82 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[5]))
						v84 = v82 | v69
					}
					v87 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[6]))
					if v87&v33 == int64(0) {
						v97 = v84
					} else {
						v95 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[7]))
						v97 = v95 | v84
					}
					v102 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[8]))
					if v102&v33 == int64(0) {
						v112 = v97
					} else {
						v110 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[9]))
						v112 = v110 | v97
					}
					v115 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[10]))
					if v115&v33 == int64(0) {
						v125 = v112
					} else {
						v123 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[11]))
						v125 = v123 | v112
					}
					v130 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[12]))
					if v130&v33 == int64(0) {
						v140 = v125
					} else {
						v138 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[13]))
						v140 = v138 | v125
					}
					v143 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[14]))
					if v143&v33 == int64(0) {
						v153 = v140
					} else {
						v151 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[15]))
						v153 = v151 | v140
					}
					v158 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[16]))
					if v158&v33 == int64(0) {
						v168 = v153
					} else {
						v166 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[17]))
						v168 = v166 | v153
					}
					v171 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[18]))
					if v171&v33 == int64(0) {
						v181 = v168
					} else {
						v179 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[19]))
						v181 = v179 | v168
					}
					v186 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[20]))
					if v186&v33 == int64(0) {
						v196 = v181
					} else {
						v194 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[21]))
						v196 = v194 | v181
					}
					*(*uint32)(unsafe.Add(mBase, uint32(v31)+4)) = uint32(v196)
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v27 + int32(1)
					return
				} else {
					v19 = int32(8192)
					if v16 < v19 {
						v22 = v16
					} else {
						v22 = v19
					}
					v24 = F_getKeysPrepareResult(m, v13, v22+v16)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
						v27 = v26
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						v31 = v28 + v27<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = l1
						v33 = base.I64_extend_i32_s(l2)
						v44 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[0]))
						if base.B2i32(v44&v33 == int64(0)) == int32(0) {
							v55 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[1]))
							v56 = v55
						} else {
							v56 = int64(0)
						}
						v59 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[2]))
						if v59&v33 == int64(0) {
							v69 = v56
						} else {
							v67 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[3]))
							v69 = v67 | v56
						}
						v74 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[4]))
						if v74&v33 == int64(0) {
							v84 = v69
						} else {
							v82 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[5]))
							v84 = v82 | v69
						}
						v87 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[6]))
						if v87&v33 == int64(0) {
							v97 = v84
						} else {
							v95 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[7]))
							v97 = v95 | v84
						}
						v102 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[8]))
						if v102&v33 == int64(0) {
							v112 = v97
						} else {
							v110 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[9]))
							v112 = v110 | v97
						}
						v115 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[10]))
						if v115&v33 == int64(0) {
							v125 = v112
						} else {
							v123 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[11]))
							v125 = v123 | v112
						}
						v130 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[12]))
						if v130&v33 == int64(0) {
							v140 = v125
						} else {
							v138 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[13]))
							v140 = v138 | v125
						}
						v143 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[14]))
						if v143&v33 == int64(0) {
							v153 = v140
						} else {
							v151 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[15]))
							v153 = v151 | v140
						}
						v158 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[16]))
						if v158&v33 == int64(0) {
							v168 = v153
						} else {
							v166 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[17]))
							v168 = v166 | v153
						}
						v171 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[18]))
						if v171&v33 == int64(0) {
							v181 = v168
						} else {
							v179 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[19]))
							v181 = v179 | v168
						}
						v186 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[20]))
						if v186&v33 == int64(0) {
							v196 = v181
						} else {
							v194 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPosWithFlags[21]))
							v196 = v194 | v181
						}
						*(*uint32)(unsafe.Add(mBase, uint32(v31)+4)) = uint32(v196)
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = v27 + int32(1)
						return
					}
				}
			}
		}
	}
}
func F_VM_KeyExists(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+96))
	v6 = F_lookupKeyReadWithFlags(m, v4, l1, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_VM_KeyType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v2 = int32(0)
	if l0 == v2 {
		v20 = v2
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v6 == int32(0) {
			v20 = v2
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v11 = v9 & int32(15)
			if base.Ui32(int32(6)) < base.Ui32(v11) {
				v20 = v2
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v11<<(uint(int32(2))%32))+uint32(_c_F_VM_KeyType[0])))
				v20 = v18
			}
		}
	}
	return v20
}
func F_VM_KillForkChild(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_VM_KillForkChild[0]))
	if v10 != int32(4) {
		v63 = v8
		m.G0 = v6 + int32(16)
		return v63
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_VM_KillForkChild[1]))
		if v14 != l0 {
			v63 = v8
			m.G0 = v6 + int32(16)
			return v63
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_VM_KillForkChild[2]))
			if int32(1) < v17 {
				v29 = l0
				v31 = F_kill(m, v29, int32(10))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v31 == int32(-1) {
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, _c_F_VM_KillForkChild[1]))
						v37 = v36
						for {
							v42 = int32(0)
							v44 = F___syscall_wait4(m, v37, v6+int32(12), v42, v42)
							mBase = m.M
							v45 = F___syscall_ret(m, v44)
							mBase = m.M
							v47 = *(*int32)(unsafe.Add(mBase, _c_F_VM_KillForkChild[1]))
							if v45 != v47 {
								v37 = v47
								continue
							} else {
								break
							}
							break
						}
					}
					F_resetChildState(m)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_VM_KillForkChild[3])) = v54
						*(*int32)(unsafe.Add(mBase, _c_F_VM_KillForkChild[4])) = v54
						v63 = v54
						m.G0 = v6 + int32(16)
						return v63
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F__serverLog(m, int32(1), int32(_a_F_VM_KillForkChild_0), v6)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, _c_F_VM_KillForkChild[1]))
					v29 = v28
					v31 = F_kill(m, v29, int32(10))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 == int32(-1) {
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, _c_F_VM_KillForkChild[1]))
							v37 = v36
							for {
								v42 = int32(0)
								v44 = F___syscall_wait4(m, v37, v6+int32(12), v42, v42)
								mBase = m.M
								v45 = F___syscall_ret(m, v44)
								mBase = m.M
								v47 = *(*int32)(unsafe.Add(mBase, _c_F_VM_KillForkChild[1]))
								if v45 != v47 {
									v37 = v47
									continue
								} else {
									break
								}
								break
							}
						}
						F_resetChildState(m)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v54 = int32(0)
							*(*int32)(unsafe.Add(mBase, _c_F_VM_KillForkChild[3])) = v54
							*(*int32)(unsafe.Add(mBase, _c_F_VM_KillForkChild[4])) = v54
							v63 = v54
							m.G0 = v6 + int32(16)
							return v63
						}
					}
				}
			}
		}
	}
}
func F_VM_ListGet(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	v6 = F_moduleListIteratorSeek(m, l0, l1, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v14 = F_listTypeGet(m, l0+int32(24))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = F_getDecodedObject(m, v14)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					F_decrRefCount(m, v14)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+28)))
						if v21&int32(1) == int32(0) {
							return v16
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
							if v26 == v27 {
								v30 = int32(8)
								if v30 < v26 {
									v33 = v26
								} else {
									v33 = v30
								}
								*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v33 << (uint(int32(1)) % 32)
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
								v40 = F_valkey_realloc(m, v37, v33<<(uint(int32(4))%32))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v40
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
									v44 = v43
									v45 = v40
									v48 = v45 + v44<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v48))) = v16
									v50 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v50
									*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v44 + v50
									return v16
								}
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
								v44 = v26
								v45 = v29
								v48 = v45 + v44<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v48))) = v16
								v50 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v50
								*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v44 + v50
								return v16
							}
						}
					}
				}
			}
		} else {
			return int32(0)
		}
	}
}
func F_VM_LoadStringBuffer(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_moduleLoadString(m, l0, int32(1), l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_VM_MallocUsableSize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-8))))
	return v4 & int32(2147483647)
}
func F_VM_Microseconds(m *base.Module) int64 {
	var v1 int64
	_ = v1
	v1 = F_ustime(m)
	return v1
}
func F_VM_ModuleTypeSetValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(1)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v11&int32(2) == int32(0) {
		v43 = v10
		m.G0 = v8 + int32(16)
		return v43
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v16 != 0 {
			v43 = v10
			m.G0 = v8 + int32(16)
			return v43
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v17 == int32(0) {
				v28 = F_createModuleObject(m, l1, l2)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v28
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_setKey(m, v32, v33, v34, v8+int32(12), int32(10))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v40
						v43 = int32(0)
						m.G0 = v8 + int32(16)
						return v43
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v22 = F_dbDelete(m, v20, v21)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
					v28 = F_createModuleObject(m, l1, l2)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v28
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						F_setKey(m, v32, v33, v34, v8+int32(12), int32(10))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v40
							v43 = int32(0)
							m.G0 = v8 + int32(16)
							return v43
						}
					}
				}
			}
		}
	}
}
func F_VM_PublishMessage(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_pubsubPublishMessageAndPropagateToCluster(m, l1, l2, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_VM_RdbStreamCreateFromFile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v4 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(1)
		v10 = F_zstrdup(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v10
			return v4
		}
	}
}
func F_VM_Realloc(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_zrealloc_usable(m, l0, l1, int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_VM_RedactClientCommandArgument(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	if l0 != 0 {
		v6 = int32(1)
		if l1 < v6 {
			v20 = v6
			return v20
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v9 == int32(0) {
				v20 = v6
				return v20
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
				if v12 <= l1 {
					v20 = v6
					return v20
				} else {
					F_redactClientCommandArgument(m, v9, l1)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return int32(0)
					} else {
						v20 = int32(0)
						return v20
					}
				}
			}
		}
	} else {
		return int32(1)
	}
}
func F_VM_RegisterBoolConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v38 int32
	_ = v38
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_moduleConfigValidityCheck(m, v12, l1, l3, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v38 = int32(1)
			return v38
		} else {
			v19 = F_valkey_malloc(m, int32(24))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = F_sdsnew(m, l1)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l6
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v12
					*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l7
					*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l4
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
					v30 = F_listAddNodeTail(m, v29, v19)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
						F_addModuleBoolConfig(m, v32, l1, l3&int32(113), v19, l2)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v38 = int32(0)
							return v38
						}
					}
				}
			}
		}
	}
}
func F_VM_RegisterClusterMessageReceiver(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int64
	_ = v72
	var v76 int64
	_ = v76
	var v117 int64
	_ = v117
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int64
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v145 int64
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_VM_RegisterClusterMessageReceiver[0]))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v17 = int64(0)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_VM_RegisterClusterMessageReceiver[1]))
	v29 = F_strlen(m, v14)
	mBase = m.M
	goto L5
L3:
	;
	v127 = l1 << (uint(int32(2)) % 32)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v127)+uint32(_c_F_VM_RegisterClusterMessageReceiver[2])))
	if v130 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	goto L3
L5:
	;
	if v29 != int32(9) {
		v117 = v17
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v34 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14))))
	v35 = F_strchr(m, v27, v34)
	mBase = m.M
	if v35 == int32(0) {
		v117 = v17
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+1)))
	v39 = F_strchr(m, v27, v38)
	mBase = m.M
	if v39 == int32(0) {
		v117 = v17
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+2)))
	v43 = F_strchr(m, v27, v42)
	mBase = m.M
	if v43 == int32(0) {
		v117 = v17
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+3)))
	v47 = F_strchr(m, v27, v46)
	mBase = m.M
	if v47 == int32(0) {
		v117 = v17
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+4)))
	v51 = F_strchr(m, v27, v50)
	mBase = m.M
	if v51 == int32(0) {
		v117 = v17
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+5)))
	v55 = F_strchr(m, v27, v54)
	mBase = m.M
	if v55 == int32(0) {
		v117 = v17
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+6)))
	v59 = F_strchr(m, v27, v58)
	mBase = m.M
	if v59 == int32(0) {
		v117 = v17
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v62 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+7)))
	v63 = F_strchr(m, v27, v62)
	mBase = m.M
	if v63 == int32(0) {
		v117 = v17
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+8)))
	v67 = F_strchr(m, v27, v66)
	mBase = m.M
	if v67 == int32(0) {
		v117 = v17
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v72 = int64(12)
	v76 = int64(6)
	v117 = ((((base.I64_extend_i32_u(v35-v27)<<(uint(v72)%64)|base.I64_extend_i32_u(v39-v27)<<(uint(v76)%64)|base.I64_extend_i32_u(v43-v27))<<(uint(v72)%64)|base.I64_extend_i32_u(v47-v27)<<(uint(v76)%64)|base.I64_extend_i32_u(v51-v27))<<(uint(v72)%64)|base.I64_extend_i32_u(v55-v27)<<(uint(v76)%64)|base.I64_extend_i32_u(v59-v27))<<(uint(v72)%64)|base.I64_extend_i32_u(v63-v27)<<(uint(v76)%64)|base.I64_extend_i32_u(v67-v27))<<(uint(int64(10))%64) | base.I64_extend_i32_u(v12)
	goto L4
L16:
	;
	if l2 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L17:
	;
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v130)))
	if v133 != v117 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if l2 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v136 = v130
	goto L21
L20:
	;
	v150 = v12
	v153 = v130
	goto L18
L21:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	if v142 == int32(0) {
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v150 = v136
	v153 = v142
	goto L18
L23:
	;
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v142)))
	if v145 != v117 {
		v136 = v142
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
	if v150 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = l2
	return
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+uint32(_c_F_VM_RegisterClusterMessageReceiver[2]))) = v157
	F_valkey_free(m, v153)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L29
	} else {
		goto L31
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+16)) = v157
	F_valkey_free(m, v153)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return
L30:
	;
	return
L31:
	;
	return
L32:
	;
	v176 = F_valkey_malloc(m, int32(24))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v176))) = v117
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v176)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v176)+12)) = v179
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v127)+uint32(_c_F_VM_RegisterClusterMessageReceiver[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v176)+16)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v127)+uint32(_c_F_VM_RegisterClusterMessageReceiver[2]))) = v176
	goto L1
}
func F_VM_RegisterStringConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
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
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	v11 = int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_moduleConfigValidityCheck(m, v12, l1, l3, v11)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v47 = v11
			return v47
		} else {
			v19 = F_valkey_malloc(m, int32(24))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = F_sdsnew(m, l1)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l6
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v12
					*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l7
					*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l4
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
					v30 = F_listAddNodeTail(m, v29, v19)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
						v35 = int32(0)
						if l2 == v35 {
							v41 = v35
							F_addModuleStringConfig(m, v34, l1, l3&int32(113), v19, v41)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v47 = v35
								return v47
							}
						} else {
							v39 = F_sdsnew(m, l2)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = v39
								F_addModuleStringConfig(m, v34, l1, l3&int32(113), v19, v41)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v47 = v35
									return v47
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_VM_ReplyWithCallReply(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
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
	v5 = m.G0
	v6 = int32(16)
	v7 = v5 - v6
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v9&v6 == int32(0) {
		v21 = l0 + int32(8)
		v22 = int32(0)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		if v23 == v22 {
			v50 = v22
			m.G0 = v7 + int32(16)
			return v50
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+224)))
			if v26 != int32(2) {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(12)))) = v37
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				F_addReplyProto(m, v23, v39, v40)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					if v45 == int32(0) {
						v50 = v22
						m.G0 = v7 + int32(16)
						return v50
					} else {
						F_deferredAfterErrorReply(m, v23, v45)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							v50 = v22
							m.G0 = v7 + int32(16)
							return v50
						}
					}
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				if v29&int32(4) == int32(0) {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v7+int32(12)))) = v37
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					F_addReplyProto(m, v23, v39, v40)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						if v45 == int32(0) {
							v50 = v22
							m.G0 = v7 + int32(16)
							return v50
						} else {
							F_deferredAfterErrorReply(m, v23, v45)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v50 = v22
								m.G0 = v7 + int32(16)
								return v50
							}
						}
					}
				} else {
					v50 = int32(1)
					m.G0 = v7 + int32(16)
					return v50
				}
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v14 != 0 {
			v21 = v14 + int32(36)
			v22 = int32(0)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			if v23 == v22 {
				v50 = v22
				m.G0 = v7 + int32(16)
				return v50
			} else {
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+224)))
				if v26 != int32(2) {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v7+int32(12)))) = v37
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					F_addReplyProto(m, v23, v39, v40)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						if v45 == int32(0) {
							v50 = v22
							m.G0 = v7 + int32(16)
							return v50
						} else {
							F_deferredAfterErrorReply(m, v23, v45)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v50 = v22
								m.G0 = v7 + int32(16)
								return v50
							}
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v29&int32(4) == int32(0) {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v7+int32(12)))) = v37
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						F_addReplyProto(m, v23, v39, v40)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							if v45 == int32(0) {
								v50 = v22
								m.G0 = v7 + int32(16)
								return v50
							} else {
								F_deferredAfterErrorReply(m, v23, v45)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									v50 = v22
									m.G0 = v7 + int32(16)
									return v50
								}
							}
						}
					} else {
						v50 = int32(1)
						m.G0 = v7 + int32(16)
						return v50
					}
				}
			}
		} else {
			v50 = int32(0)
			m.G0 = v7 + int32(16)
			return v50
		}
	}
}
func F_VM_ReplyWithCustomErrorFormat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
	F_moduleReplyErrorFormatInternal(m, l0, base.B2i32(l1 == int32(0))|int32(2), l2, l3)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return int32(0)
	}
}
func F_VM_SaveDataTypeToString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
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
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v12 = F_sdsempty(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v18 = F___memcpy(m, v8+int32(32), int32(_a_F_VM_SaveDataTypeToString_0), int32(80))
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v12
		*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = int64(-4294967296)
		v24 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v24
		*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v8 + int32(32)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l2
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
		m.T0[v35].(func(*base.Module, int32, int32))(m, v8, l1)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
			if v38 == int32(0) {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				if v46 != 0 {
					v89 = v24
					m.G0 = v8 + int32(112)
					return v89
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
					v49 = F_createObject(m, int32(0), v48)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						if l0 == int32(0) {
							v89 = v49
							m.G0 = v8 + int32(112)
							return v89
						} else {
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
							if v53&int32(1) == int32(0) {
								v89 = v49
								m.G0 = v8 + int32(112)
								return v89
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v58 == v59 {
									v62 = int32(8)
									if v62 < v58 {
										v65 = v58
									} else {
										v65 = v62
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v65 << (uint(int32(1)) % 32)
									v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v72 = F_valkey_realloc(m, v69, v65<<(uint(int32(4))%32))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v72
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v76 = v72
										v77 = v75
										v80 = v76 + v77<<(uint(int32(3))%32)
										*(*int32)(unsafe.Add(mBase, uint32(v80))) = v49
										v82 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v82
										*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v77 + v82
										v89 = v49
										m.G0 = v8 + int32(112)
										return v89
									}
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v76 = v61
									v77 = v58
									v80 = v76 + v77<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v80))) = v49
									v82 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v82
									*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v77 + v82
									v89 = v49
									m.G0 = v8 + int32(112)
									return v89
								}
							}
						}
					}
				}
			} else {
				F_moduleFreeContext(m, v38)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
					F_valkey_free(m, v43)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						if v46 != 0 {
							v89 = v24
							m.G0 = v8 + int32(112)
							return v89
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							v49 = F_createObject(m, int32(0), v48)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								if l0 == int32(0) {
									v89 = v49
									m.G0 = v8 + int32(112)
									return v89
								} else {
									v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
									if v53&int32(1) == int32(0) {
										v89 = v49
										m.G0 = v8 + int32(112)
										return v89
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v58 == v59 {
											v62 = int32(8)
											if v62 < v58 {
												v65 = v58
											} else {
												v65 = v62
											}
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v65 << (uint(int32(1)) % 32)
											v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											v72 = F_valkey_realloc(m, v69, v65<<(uint(int32(4))%32))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v72
												v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
												v76 = v72
												v77 = v75
												v80 = v76 + v77<<(uint(int32(3))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v80))) = v49
												v82 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v82
												*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v77 + v82
												v89 = v49
												m.G0 = v8 + int32(112)
												return v89
											}
										} else {
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											v76 = v61
											v77 = v58
											v80 = v76 + v77<<(uint(int32(3))%32)
											*(*int32)(unsafe.Add(mBase, uint32(v80))) = v49
											v82 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v82
											*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v77 + v82
											v89 = v49
											m.G0 = v8 + int32(112)
											return v89
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
func F_VM_SaveFloat(m *base.Module, l0 int32, l1 float32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v8 == int32(0) {
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v51 = F_rdbSaveLen(m, v49, int64(3))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				if v51 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
					return
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v55 + v51
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v59 = F_rdbSaveBinaryFloatValue(m, v58, l1)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						if v59 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
							return
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v63 + v59
							return
						}
					}
				}
			}
		} else {
			v11 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v11
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-1)))))
			switch v17 & int32(7) {
			case 0:
				v34 = int32(base.Ui32(v17) >> (uint(int32(3)) % 32))
			case 1:
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-3)))))
				v34 = v24
			case 2:
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-5)))))
				v34 = v27
			case 3:
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-9))))
				v34 = v30
			case 4:
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-17))))
				v34 = v33
			default:
				v34 = v11
			}
			v35 = F_rdbWriteRaw(m, v13, v8, v34)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_sdsfree(m, v8)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					if v35 < int32(0) {
						if v35 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
							return
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v51 = F_rdbSaveLen(m, v49, int64(3))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								if v51 == int32(-1) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
									return
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v55 + v51
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v59 = F_rdbSaveBinaryFloatValue(m, v58, l1)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										if v59 == int32(-1) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
											return
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = v63 + v59
											return
										}
									}
								}
							}
						}
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v41 + v35
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v51 = F_rdbSaveLen(m, v49, int64(3))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							if v51 == int32(-1) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
								return
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v55 + v51
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v59 = F_rdbSaveBinaryFloatValue(m, v58, l1)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									if v59 == int32(-1) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
										return
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v63 + v59
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
func F_VM_SaveString(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v8 == int32(0) {
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v51 = F_rdbSaveLen(m, v49, int64(5))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				if v51 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
					return
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v55 + v51
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v59 = F_rdbSaveStringObject(m, v58, l1)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						if v59 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
							return
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v63 + v59
							return
						}
					}
				}
			}
		} else {
			v11 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v11
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-1)))))
			switch v17 & int32(7) {
			case 0:
				v34 = int32(base.Ui32(v17) >> (uint(int32(3)) % 32))
			case 1:
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-3)))))
				v34 = v24
			case 2:
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-5)))))
				v34 = v27
			case 3:
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-9))))
				v34 = v30
			case 4:
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-17))))
				v34 = v33
			default:
				v34 = v11
			}
			v35 = F_rdbWriteRaw(m, v13, v8, v34)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_sdsfree(m, v8)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					if v35 < int32(0) {
						if v35 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
							return
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v51 = F_rdbSaveLen(m, v49, int64(5))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								if v51 == int32(-1) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
									return
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v55 + v51
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v59 = F_rdbSaveStringObject(m, v58, l1)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										if v59 == int32(-1) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
											return
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = v63 + v59
											return
										}
									}
								}
							}
						}
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v41 + v35
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v51 = F_rdbSaveLen(m, v49, int64(5))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							if v51 == int32(-1) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
								return
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v55 + v51
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v59 = F_rdbSaveStringObject(m, v58, l1)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									if v59 == int32(-1) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
										return
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v63 + v59
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
func F_VM_SaveStringBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v9 == int32(0) {
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v52 = F_rdbSaveLen(m, v50, int64(5))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				if v52 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
					return
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v56 + v52
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v60 = F_rdbSaveRawString(m, v59, l1, l2)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						if v60 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
							return
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v64 + v60
							return
						}
					}
				}
			}
		} else {
			v12 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v12
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-1)))))
			switch v18 & int32(7) {
			case 0:
				v35 = int32(base.Ui32(v18) >> (uint(int32(3)) % 32))
			case 1:
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(-3)))))
				v35 = v25
			case 2:
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9+int32(-5)))))
				v35 = v28
			case 3:
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(-9))))
				v35 = v31
			case 4:
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(-17))))
				v35 = v34
			default:
				v35 = v12
			}
			v36 = F_rdbWriteRaw(m, v14, v9, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				F_sdsfree(m, v9)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					if v36 < int32(0) {
						if v36 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
							return
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v52 = F_rdbSaveLen(m, v50, int64(5))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								if v52 == int32(-1) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
									return
								} else {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v56 + v52
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v60 = F_rdbSaveRawString(m, v59, l1, l2)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										if v60 == int32(-1) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
											return
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = v64 + v60
											return
										}
									}
								}
							}
						}
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v42 + v36
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v52 = F_rdbSaveLen(m, v50, int64(5))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							if v52 == int32(-1) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
								return
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v56 + v52
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v60 = F_rdbSaveRawString(m, v59, l1, l2)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									if v60 == int32(-1) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
										return
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v64 + v60
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
func F_VM_ScanCursorDestroy(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_valkey_free(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_VM_ScriptingEngineDebuggerLog(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if l1 == int32(0) {
		F_scriptingEngineDebuggerLog(m, l0)
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	} else {
		F_scriptingEngineDebuggerLogWithMaxLen(m, l0)
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_VM_ScriptingEngineDebuggerLogRespReply(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v5+int32(12)))) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_scriptingEngineDebuggerLogRespReplyStr(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_VM_SendChildHeartbeat(m *base.Module, l0 float64) {
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	v2 = int32(0)
	F_sendChildInfoGeneric(m, v2, v2, v2, l0, int32(_a_F_VM_SendChildHeartbeat_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_VM_ServerInfoGetField(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
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
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l2&int32(3) == int32(0) {
		v32 = l2
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v327
L2:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+int32(-1)))))
	switch v268 & int32(7) {
	case 0:
		goto L63
	case 1:
		goto L62
	case 2:
		goto L61
	case 3:
		goto L60
	case 4:
		goto L59
	default:
		v285 = int32(0)
		goto L58
	}
L3:
	;
	v67 = v8 + int32(12)
	v68 = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v65 == v68 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v65 = v57 - l2
	goto L3
L5:
	;
	v36 = v32
	goto L13
L6:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v21 = l2
	goto L9
L8:
	;
	v65 = l2 - l2
	goto L3
L9:
	;
	v25 = v21 + int32(1)
	if v25&int32(3) == int32(0) {
		v32 = v25
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v30 != 0 {
		v21 = v25
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v57 = v25
	goto L4
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v45 = int32(-2139062144)
	if (int32(16843008)-v42|v42)&v45 == v45 {
		v36 = v36 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v51 = v36
	goto L16
L15:
	;
	goto L14
L16:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 != 0 {
		v51 = v51 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v57 = v51
	goto L4
L18:
	;
	goto L17
L19:
	;
	if v261 != 0 {
		goto L2
	} else {
		goto L57
	}
L20:
	;
	if v220 != v65 {
		v261 = v68
		goto L47
	} else {
		goto L48
	}
L21:
	;
	v211 = int32(0)
	v217 = v76
	v218 = v77
	v220 = v211
	v224 = v211
	goto L20
L22:
	;
	if base.Ui32(v77) < base.Ui32(int32(8)) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v88 = v76
	v89 = v77
	v91 = int32(0)
	goto L25
L24:
	;
	v217 = v201
	v218 = v202
	v220 = v204
	v224 = base.B2i32(v207 != int32(0))
	goto L20
L25:
	;
	v97 = int32(base.Ui32(v89) >> (uint(int32(3)) % 32))
	v98 = int32(4)
	v99 = v88 + v98
	if v89&v98 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v201 = v192
	v202 = v193
	v204 = v177
	v207 = v182
	goto L24
L27:
	;
	v182 = int32(0)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v99+v97+(v182-v97)&int32(3)+v170<<(uint(int32(2))%32))))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	if base.Ui32(v193) < base.Ui32(int32(8)) {
		v201 = v192
		v202 = v193
		v204 = v177
		v207 = v182
		goto L24
	} else {
		goto L45
	}
L28:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v91))))
	v148 = int32(0)
	goto L39
L29:
	;
	v104 = int32(0)
	if base.Ui32(v65) <= base.Ui32(v91) {
		v137 = v91
		v140 = v104
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v140 == v97 {
		v170 = v104
		v177 = v137
		goto L27
	} else {
		goto L37
	}
L31:
	;
	v114 = v91
	v117 = v104
	goto L32
L32:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99+v117))))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v114))))
	if v120 != v122 {
		v137 = v114
		v140 = v117
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v137 = v125
	v140 = v127
	goto L30
L34:
	;
	v124 = int32(1)
	v125 = v114 + v124
	v127 = v117 + v124
	if base.Ui32(v97) <= base.Ui32(v127) {
		v137 = v125
		v140 = v127
		goto L30
	} else {
		goto L35
	}
L35:
	;
	if base.Ui32(v125) < base.Ui32(v65) {
		v114 = v125
		v117 = v127
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v201 = v88
	v202 = v89
	v204 = v137
	v207 = v140
	goto L24
L38:
	;
	if v148 != v97 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99+v148))))
	if v161 == v145&int32(255) {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v163 = int32(1)
	v165 = v148 + v163
	if v165 != v97 {
		v148 = v165
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v217 = v88
	v218 = v89
	v220 = v91
	v224 = v163
	goto L20
L43:
	;
	v170 = v148
	v177 = v91 + int32(1)
	goto L27
L44:
	;
	v201 = v88
	v202 = v89
	v204 = v91
	v207 = v97
	goto L24
L45:
	;
	if base.Ui32(v177) < base.Ui32(v65) {
		v88 = v192
		v89 = v193
		v91 = v177
		goto L25
	} else {
		goto L46
	}
L46:
	;
	goto L26
L47:
	;
	goto L19
L48:
	;
	v226 = int32(0)
	if v218&int32(1) == v226 {
		v261 = v226
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v232 = v218 & int32(4)
	if v224&base.B2i32(v232 != int32(0)) != 0 {
		v261 = v226
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v236 = int32(1)
	if v67 == int32(0) {
		v261 = v236
		goto L47
	} else {
		goto L51
	}
L51:
	;
	if v218&int32(2) != 0 {
		v258 = int32(0)
		goto L52
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v258
	v261 = v236
	goto L47
L53:
	;
	v242 = int32(3)
	v243 = int32(base.Ui32(v218) >> (uint(v242) % 32))
	if v232 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v253 = int32(4)
	goto L56
L55:
	;
	v253 = v243 << (uint(int32(2)) % 32)
	goto L56
L56:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v217+v243+(int32(0)-v243)&v242+v253+int32(4))))
	v258 = v257
	goto L52
L57:
	;
	v327 = int32(0)
	goto L1
L58:
	;
	v286 = F_createStringObject_1(m, v265, v285)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v265+int32(-17))))
	v285 = v284
	goto L58
L60:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v265+int32(-9))))
	v285 = v281
	goto L58
L61:
	;
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v265+int32(-5)))))
	v285 = v278
	goto L58
L62:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+int32(-3)))))
	v285 = v275
	goto L58
L63:
	;
	v285 = int32(base.Ui32(v268) >> (uint(int32(3)) % 32))
	goto L58
L64:
	;
	return int32(0)
L65:
	;
	if l0 == int32(0) {
		v327 = v286
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v292&int32(1) == int32(0) {
		v327 = v286
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v297 == v298 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v319 = v316 + v315<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v319))) = v286
	v321 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v319)+4)) = v321
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v315 + v321
	v327 = v286
	goto L1
L69:
	;
	v301 = int32(8)
	if v301 < v297 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v315 = v297
	v316 = v300
	goto L68
L71:
	;
	v304 = v297
	goto L73
L72:
	;
	v304 = v301
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v304 << (uint(int32(1)) % 32)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v311 = F_valkey_realloc(m, v308, v304<<(uint(int32(4))%32))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L64
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v311
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v315 = v314
	v316 = v311
	goto L68
}
func F_VM_ServerInfoGetFieldDouble(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 float64
	_ = v264
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v307 float64
	_ = v307
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v331 int64
	_ = v331
	var v332 int64
	_ = v332
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v353 float64
	_ = v353
	var v362 float64
	_ = v362
	var v365 float64
	_ = v365
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1&int32(3) == int32(0) {
		v33 = l1
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v365
L2:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270+int32(-1)))))
	switch v273 & int32(7) {
	case 0:
		goto L64
	case 1:
		goto L63
	case 2:
		goto L62
	case 3:
		goto L61
	case 4:
		goto L60
	default:
		v290 = int32(0)
		goto L59
	}
L3:
	;
	v68 = v9 + int32(4)
	v69 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v66 == v69 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v66 = v58 - l1
	goto L3
L5:
	;
	v37 = v33
	goto L13
L6:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v22 = l1
	goto L9
L8:
	;
	v66 = l1 - l1
	goto L3
L9:
	;
	v26 = v22 + int32(1)
	if v26&int32(3) == int32(0) {
		v33 = v26
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v31 != 0 {
		v22 = v26
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v58 = v26
	goto L4
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v46 = int32(-2139062144)
	if (int32(16843008)-v43|v43)&v46 == v46 {
		v37 = v37 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v52 = v37
	goto L16
L15:
	;
	goto L14
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 != 0 {
		v52 = v52 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v58 = v52
	goto L4
L18:
	;
	goto L17
L19:
	;
	if v262 != 0 {
		goto L2
	} else {
		goto L57
	}
L20:
	;
	if v221 != v66 {
		v262 = v69
		goto L47
	} else {
		goto L48
	}
L21:
	;
	v212 = int32(0)
	v218 = v77
	v219 = v78
	v221 = v212
	v225 = v212
	goto L20
L22:
	;
	if base.Ui32(v78) < base.Ui32(int32(8)) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v89 = v77
	v90 = v78
	v92 = int32(0)
	goto L25
L24:
	;
	v218 = v202
	v219 = v203
	v221 = v205
	v225 = base.B2i32(v208 != int32(0))
	goto L20
L25:
	;
	v98 = int32(base.Ui32(v90) >> (uint(int32(3)) % 32))
	v99 = int32(4)
	v100 = v89 + v99
	if v90&v99 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v202 = v193
	v203 = v194
	v205 = v178
	v208 = v183
	goto L24
L27:
	;
	v183 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v100+v98+(v183-v98)&int32(3)+v171<<(uint(int32(2))%32))))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	if base.Ui32(v194) < base.Ui32(int32(8)) {
		v202 = v193
		v203 = v194
		v205 = v178
		v208 = v183
		goto L24
	} else {
		goto L45
	}
L28:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v92))))
	v149 = int32(0)
	goto L39
L29:
	;
	v105 = int32(0)
	if base.Ui32(v66) <= base.Ui32(v92) {
		v138 = v92
		v141 = v105
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v141 == v98 {
		v171 = v105
		v178 = v138
		goto L27
	} else {
		goto L37
	}
L31:
	;
	v115 = v92
	v118 = v105
	goto L32
L32:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v118))))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v115))))
	if v121 != v123 {
		v138 = v115
		v141 = v118
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v138 = v126
	v141 = v128
	goto L30
L34:
	;
	v125 = int32(1)
	v126 = v115 + v125
	v128 = v118 + v125
	if base.Ui32(v98) <= base.Ui32(v128) {
		v138 = v126
		v141 = v128
		goto L30
	} else {
		goto L35
	}
L35:
	;
	if base.Ui32(v126) < base.Ui32(v66) {
		v115 = v126
		v118 = v128
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v202 = v89
	v203 = v90
	v205 = v138
	v208 = v141
	goto L24
L38:
	;
	if v149 != v98 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v149))))
	if v162 == v146&int32(255) {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v164 = int32(1)
	v166 = v149 + v164
	if v166 != v98 {
		v149 = v166
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v218 = v89
	v219 = v90
	v221 = v92
	v225 = v164
	goto L20
L43:
	;
	v171 = v149
	v178 = v92 + int32(1)
	goto L27
L44:
	;
	v202 = v89
	v203 = v90
	v205 = v92
	v208 = v98
	goto L24
L45:
	;
	if base.Ui32(v178) < base.Ui32(v66) {
		v89 = v193
		v90 = v194
		v92 = v178
		goto L25
	} else {
		goto L46
	}
L46:
	;
	goto L26
L47:
	;
	goto L19
L48:
	;
	v227 = int32(0)
	if v219&int32(1) == v227 {
		v262 = v227
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v233 = v219 & int32(4)
	if v225&base.B2i32(v233 != int32(0)) != 0 {
		v262 = v227
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v237 = int32(1)
	if v68 == int32(0) {
		v262 = v237
		goto L47
	} else {
		goto L51
	}
L51:
	;
	if v219&int32(2) != 0 {
		v259 = int32(0)
		goto L52
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v259
	v262 = v237
	goto L47
L53:
	;
	v243 = int32(3)
	v244 = int32(base.Ui32(v219) >> (uint(v243) % 32))
	if v233 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v254 = int32(4)
	goto L56
L55:
	;
	v254 = v244 << (uint(int32(2)) % 32)
	goto L56
L56:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v218+v244+(int32(0)-v244)&v243+v254+int32(4))))
	v259 = v258
	goto L52
L57:
	;
	v264 = float64(0)
	if l2 == int32(0) {
		v365 = v264
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
	v365 = v264
	goto L1
L59:
	;
	v292 = v9 + int32(8)
	v293 = int32(0)
	v298 = m.G0
	v300 = v298 - int32(16)
	m.G0 = v300
	v302 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v302))) = v293
	v307 = F_valkey_strtod_n(m, v270, v290, v300+int32(12))
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(v292))) = v307
	if v290 == v293 {
		goto L68
	} else {
		goto L69
	}
L60:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v270+int32(-17))))
	v290 = v289
	goto L59
L61:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v270+int32(-9))))
	v290 = v286
	goto L59
L62:
	;
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v270+int32(-5)))))
	v290 = v283
	goto L59
L63:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270+int32(-3)))))
	v290 = v280
	goto L59
L64:
	;
	v290 = int32(base.Ui32(v273) >> (uint(int32(3)) % 32))
	goto L59
L65:
	;
	if l2 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L66:
	;
	if v347 != 0 {
		goto L65
	} else {
		goto L80
	}
L67:
	;
	m.G0 = v300 + int32(16)
	goto L66
L68:
	;
	v344 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v302))) = v344
	v347 = v344
	goto L67
L69:
	;
	v311 = int32(*(*int8)(unsafe.Add(mBase, uint32(v270))))
	if v311 == int32(32) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	if base.Ui32(int32(-6)) < base.Ui32(v311+int32(-14)) {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v300)+12))
	if v318-v270 != v290 {
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	if v321 == int32(68) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v332&int64(9223372036854775807)) {
		goto L68
	} else {
		goto L78
	}
L74:
	;
	if base.F64_eq(base.F64_abs(v307), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L68
	} else {
		goto L76
	}
L75:
	;
	v332 = base.I64_reinterpret_f64(v307)
	goto L73
L76:
	;
	v328 = F___fpclassify(m, v307)
	mBase = m.M
	if v328 == int32(2) {
		goto L68
	} else {
		goto L77
	}
L77:
	;
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v292)))
	v332 = v331
	goto L73
L78:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	if v338 != int32(28) {
		v347 = int32(1)
		goto L67
	} else {
		goto L79
	}
L79:
	;
	goto L68
L80:
	;
	v353 = float64(0)
	if l2 == int32(0) {
		v365 = v353
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
	v365 = v353
	goto L1
L82:
	;
	v362 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
	v365 = v362
	goto L1
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L82
}
func F_VM_ServerInfoGetFieldSigned(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int64
	_ = v264
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v341 int64
	_ = v341
	var v347 int32
	_ = v347
	var v349 int64
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v363 int64
	_ = v363
	var v368 int64
	_ = v368
	var v372 int32
	_ = v372
	var v374 int64
	_ = v374
	var v376 int32
	_ = v376
	var v384 int64
	_ = v384
	var v408 int64
	_ = v408
	var v427 int32
	_ = v427
	var v434 int64
	_ = v434
	var v443 int64
	_ = v443
	var v446 int64
	_ = v446
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1&int32(3) == int32(0) {
		v33 = l1
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v446
L2:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270+int32(-1)))))
	switch v273 & int32(7) {
	case 0:
		goto L64
	case 1:
		goto L63
	case 2:
		goto L62
	case 3:
		goto L61
	case 4:
		goto L60
	default:
		v290 = int32(0)
		goto L59
	}
L3:
	;
	v68 = v9 + int32(4)
	v69 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v66 == v69 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v66 = v58 - l1
	goto L3
L5:
	;
	v37 = v33
	goto L13
L6:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v22 = l1
	goto L9
L8:
	;
	v66 = l1 - l1
	goto L3
L9:
	;
	v26 = v22 + int32(1)
	if v26&int32(3) == int32(0) {
		v33 = v26
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v31 != 0 {
		v22 = v26
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v58 = v26
	goto L4
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v46 = int32(-2139062144)
	if (int32(16843008)-v43|v43)&v46 == v46 {
		v37 = v37 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v52 = v37
	goto L16
L15:
	;
	goto L14
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 != 0 {
		v52 = v52 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v58 = v52
	goto L4
L18:
	;
	goto L17
L19:
	;
	if v262 != 0 {
		goto L2
	} else {
		goto L57
	}
L20:
	;
	if v221 != v66 {
		v262 = v69
		goto L47
	} else {
		goto L48
	}
L21:
	;
	v212 = int32(0)
	v218 = v77
	v219 = v78
	v221 = v212
	v225 = v212
	goto L20
L22:
	;
	if base.Ui32(v78) < base.Ui32(int32(8)) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v89 = v77
	v90 = v78
	v92 = int32(0)
	goto L25
L24:
	;
	v218 = v202
	v219 = v203
	v221 = v205
	v225 = base.B2i32(v208 != int32(0))
	goto L20
L25:
	;
	v98 = int32(base.Ui32(v90) >> (uint(int32(3)) % 32))
	v99 = int32(4)
	v100 = v89 + v99
	if v90&v99 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v202 = v193
	v203 = v194
	v205 = v178
	v208 = v183
	goto L24
L27:
	;
	v183 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v100+v98+(v183-v98)&int32(3)+v171<<(uint(int32(2))%32))))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	if base.Ui32(v194) < base.Ui32(int32(8)) {
		v202 = v193
		v203 = v194
		v205 = v178
		v208 = v183
		goto L24
	} else {
		goto L45
	}
L28:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v92))))
	v149 = int32(0)
	goto L39
L29:
	;
	v105 = int32(0)
	if base.Ui32(v66) <= base.Ui32(v92) {
		v138 = v92
		v141 = v105
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v141 == v98 {
		v171 = v105
		v178 = v138
		goto L27
	} else {
		goto L37
	}
L31:
	;
	v115 = v92
	v118 = v105
	goto L32
L32:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v118))))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v115))))
	if v121 != v123 {
		v138 = v115
		v141 = v118
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v138 = v126
	v141 = v128
	goto L30
L34:
	;
	v125 = int32(1)
	v126 = v115 + v125
	v128 = v118 + v125
	if base.Ui32(v98) <= base.Ui32(v128) {
		v138 = v126
		v141 = v128
		goto L30
	} else {
		goto L35
	}
L35:
	;
	if base.Ui32(v126) < base.Ui32(v66) {
		v115 = v126
		v118 = v128
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v202 = v89
	v203 = v90
	v205 = v138
	v208 = v141
	goto L24
L38:
	;
	if v149 != v98 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v149))))
	if v162 == v146&int32(255) {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v164 = int32(1)
	v166 = v149 + v164
	if v166 != v98 {
		v149 = v166
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v218 = v89
	v219 = v90
	v221 = v92
	v225 = v164
	goto L20
L43:
	;
	v171 = v149
	v178 = v92 + int32(1)
	goto L27
L44:
	;
	v202 = v89
	v203 = v90
	v205 = v92
	v208 = v98
	goto L24
L45:
	;
	if base.Ui32(v178) < base.Ui32(v66) {
		v89 = v193
		v90 = v194
		v92 = v178
		goto L25
	} else {
		goto L46
	}
L46:
	;
	goto L26
L47:
	;
	goto L19
L48:
	;
	v227 = int32(0)
	if v219&int32(1) == v227 {
		v262 = v227
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v233 = v219 & int32(4)
	if v225&base.B2i32(v233 != int32(0)) != 0 {
		v262 = v227
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v237 = int32(1)
	if v68 == int32(0) {
		v262 = v237
		goto L47
	} else {
		goto L51
	}
L51:
	;
	if v219&int32(2) != 0 {
		v259 = int32(0)
		goto L52
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v259
	v262 = v237
	goto L47
L53:
	;
	v243 = int32(3)
	v244 = int32(base.Ui32(v219) >> (uint(v243) % 32))
	if v233 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v254 = int32(4)
	goto L56
L55:
	;
	v254 = v244 << (uint(int32(2)) % 32)
	goto L56
L56:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v218+v244+(int32(0)-v244)&v243+v254+int32(4))))
	v259 = v258
	goto L52
L57:
	;
	v264 = int64(0)
	if l2 == int32(0) {
		v446 = v264
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
	v446 = v264
	goto L1
L59:
	;
	v292 = v9 + int32(8)
	v293 = int32(0)
	if base.Ui32(v290+int32(-21)) < base.Ui32(int32(-20)) {
		v427 = v293
		goto L67
	} else {
		goto L68
	}
L60:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v270+int32(-17))))
	v290 = v289
	goto L59
L61:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v270+int32(-9))))
	v290 = v286
	goto L59
L62:
	;
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v270+int32(-5)))))
	v290 = v283
	goto L59
L63:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270+int32(-3)))))
	v290 = v280
	goto L59
L64:
	;
	v290 = int32(base.Ui32(v273) >> (uint(int32(3)) % 32))
	goto L59
L65:
	;
	if l2 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L66:
	;
	if v427 != 0 {
		goto L65
	} else {
		goto L93
	}
L67:
	;
	goto L66
L68:
	;
	v305 = int32(1)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if v290 != v305 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v427 = int32(1)
	goto L67
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v292))) = v408
	goto L69
L71:
	;
	if v306&int32(255) == int32(45) {
		goto L76
	} else {
		goto L77
	}
L72:
	;
	v310 = v306 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v310&int32(255)) {
		v427 = v293
		goto L67
	} else {
		goto L73
	}
L73:
	;
	if v292 == int32(0) {
		goto L69
	} else {
		goto L74
	}
L74:
	;
	v408 = base.I64_extend_i32_u(v310) & int64(255)
	goto L70
L75:
	;
	if base.Ui32(int32(8)) < base.Ui32((v329+int32(-49))&int32(255)) {
		v427 = v293
		goto L67
	} else {
		goto L78
	}
L76:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+1)))
	v328 = int32(2)
	v329 = v326
	v330 = v270 + int32(1)
	goto L75
L77:
	;
	v328 = v305
	v329 = v306
	v330 = v270
	goto L75
L78:
	;
	v341 = base.I64_extend_i32_u(v329+int32(-48)) & int64(255)
	if base.Ui32(v290) <= base.Ui32(v328) {
		v384 = v341
		goto L79
	} else {
		goto L80
	}
L79:
	;
	if v306&int32(255) != int32(45) {
		goto L87
	} else {
		goto L88
	}
L80:
	;
	v347 = v328
	v349 = v341
	v351 = v330
	goto L81
L81:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351)+1)))
	if base.Ui32((v353+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v427 = v293
		goto L67
	} else {
		goto L83
	}
L82:
	;
	v384 = v374
	goto L79
L83:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v349) {
		v427 = v293
		goto L67
	} else {
		goto L84
	}
L84:
	;
	v363 = v349 * int64(10)
	v368 = base.I64_extend_i32_u(v353+int32(-48)) & int64(255)
	if base.Ui64(v368^int64(-1)) < base.Ui64(v363) {
		v427 = v293
		goto L67
	} else {
		goto L85
	}
L85:
	;
	v372 = int32(1)
	v374 = v363 + v368
	v376 = v347 + v372
	if v376 != v290 {
		v347 = v376
		v349 = v374
		v351 = v351 + v372
		goto L81
	} else {
		goto L86
	}
L86:
	;
	goto L82
L87:
	;
	if v384 < int64(0) {
		v427 = v293
		goto L67
	} else {
		goto L91
	}
L88:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v384) {
		v427 = v293
		goto L67
	} else {
		goto L89
	}
L89:
	;
	if v292 == int32(0) {
		goto L69
	} else {
		goto L90
	}
L90:
	;
	v408 = int64(0) - v384
	goto L70
L91:
	;
	if v292 == int32(0) {
		goto L69
	} else {
		goto L92
	}
L92:
	;
	v408 = v384
	goto L70
L93:
	;
	v434 = int64(0)
	if l2 == int32(0) {
		v446 = v434
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
	v446 = v434
	goto L1
L95:
	;
	v443 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	v446 = v443
	goto L1
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L95
}
func F_VM_SetAbsExpire(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v4 = int32(1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v5&int32(2) == int32(0) {
		v31 = v4
		return v31
	} else {
		if l1 < int64(-1) {
			v31 = v4
			return v31
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v12 == int32(0) {
				v31 = v4
				return v31
			} else {
				if l1 == int64(-1) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v28 = F_removeExpire(m, v26, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v31 = int32(0)
						return v31
					}
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v21 = F_setExpire(m, v18, v19, v20, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v21
						v31 = int32(0)
						return v31
					}
				}
			}
		}
	}
}
func F_VM_SetCommandACLCategories(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v6 = int32(1)
	if l0 == int32(0) {
		v32 = v6
		return v32
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v9 == int32(0) {
			v32 = v6
			return v32
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
			if v12 == int32(0) {
				v32 = v6
				return v32
			} else {
				if l1 != 0 {
					v16 = F_categoryFlagsFromString(m, l1)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						if v16 == int64(-1) {
							v32 = v6
						} else {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v23 = v22
							v24 = v16
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v25)+64)) = v24
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
							*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v27 + int32(1)
							v32 = int32(0)
						}
						return v32
					}
				} else {
					v23 = v9
					v24 = int64(0)
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v25)+64)) = v24
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v27 + int32(1)
					v32 = int32(0)
					return v32
				}
			}
		}
	}
}
func F_VM_SetCommandInfo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v123 int64
	_ = v123
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
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
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
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
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
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
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v344 int32
	_ = v344
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v464 int32
	_ = v464
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v520 int32
	_ = v520
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int64
	_ = v544
	var v555 int64
	_ = v555
	var v566 int64
	_ = v566
	var v567 int64
	_ = v567
	var v570 int64
	_ = v570
	var v578 int64
	_ = v578
	var v580 int64
	_ = v580
	var v585 int64
	_ = v585
	var v593 int64
	_ = v593
	var v595 int64
	_ = v595
	var v598 int64
	_ = v598
	var v606 int64
	_ = v606
	var v608 int64
	_ = v608
	var v613 int64
	_ = v613
	var v621 int64
	_ = v621
	var v623 int64
	_ = v623
	var v626 int64
	_ = v626
	var v634 int64
	_ = v634
	var v636 int64
	_ = v636
	var v641 int64
	_ = v641
	var v649 int64
	_ = v649
	var v651 int64
	_ = v651
	var v654 int64
	_ = v654
	var v662 int64
	_ = v662
	var v664 int64
	_ = v664
	var v669 int64
	_ = v669
	var v677 int64
	_ = v677
	var v679 int64
	_ = v679
	var v682 int64
	_ = v682
	var v690 int64
	_ = v690
	var v692 int64
	_ = v692
	var v697 int64
	_ = v697
	var v705 int64
	_ = v705
	var v707 int64
	_ = v707
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
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
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v792 int32
	_ = v792
	var v806 int32
	_ = v806
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v835 int32
	_ = v835
	var v847 int32
	_ = v847
	var v853 int32
	_ = v853
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_VM_SetCommandInfo_0), int32(_a_F_VM_SetCommandInfo_1), int32(2060))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L9
	} else {
		goto L178
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_VM_SetCommandInfo_2), int32(_a_F_VM_SetCommandInfo_1), int32(2028))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L9
	} else {
		goto L177
	}
L3:
	;
	m.G0 = v14 + int32(96)
	return v835
L4:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	if v214 != 0 {
		goto L55
	} else {
		goto L56
	}
L5:
	;
	goto L53
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v28 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[0]))
	if int32(3) < v18 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	F__serverLog(m, int32(3), int32(_a_F_VM_SetCommandInfo_3), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L5
L11:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v75 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v31 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v42 = int32(0)
	v43 = v28
	goto L14
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v47 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L11
L16:
	;
	v60 = v42 + int32(1)
	v62 = v28 + v60*v34
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v63 != 0 {
		v42 = v60
		v43 = v62
		goto L14
	} else {
		goto L20
	}
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[0]))
	if int32(3) < v49 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v42
	F__serverLog(m, int32(3), int32(_a_F_VM_SetCommandInfo_4), v14+int32(80))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	goto L5
L20:
	;
	goto L15
L21:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v196 = F_moduleValidateCommandArgs(m, v195, v16)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L9
	} else {
		goto L51
	}
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	if v78 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v87 = v78
	v89 = int32(0)
	v90 = v75
	goto L24
L24:
	;
	if v89 != int32(2147483647) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L21
L26:
	;
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v90)+8))
	if base.I64_popcnt(v105&int64(15)) == int64(1) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[0]))
	if int32(3) < v97 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	F__serverLog(m, int32(3), int32(_a_F_VM_SetCommandInfo_5), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	goto L5
L30:
	;
	v123 = v105 & int64(224)
	if v123&(v123+int64(-1)) == int64(0) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[0]))
	if int32(3) < v112 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v89
	F__serverLog(m, int32(3), int32(_a_F_VM_SetCommandInfo_6), v14+int32(64))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	goto L5
L34:
	;
	switch v87 + int32(-1) {
	case 0, 1:
		goto L38
	case 2:
		goto L40
	default:
		goto L39
	}
L35:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[0]))
	if int32(3) < v130 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v89
	F__serverLog(m, int32(3), int32(_a_F_VM_SetCommandInfo_7), v14+int32(48))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	goto L5
L38:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v90)+28))
	if base.Ui32(v164) < base.Ui32(int32(4)) {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[0]))
	if int32(3) < v155 {
		goto L5
	} else {
		goto L44
	}
L40:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	if v142 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[0]))
	if int32(3) < v144 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v89
	F__serverLog(m, int32(3), int32(_a_F_VM_SetCommandInfo_8), v14+int32(32))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	goto L5
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v89
	F__serverLog(m, int32(3), int32(_a_F_VM_SetCommandInfo_9), v14)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L45
	}
L45:
	;
	goto L5
L46:
	;
	v180 = v89 + int32(1)
	v182 = v75 + v180*v81
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+16))
	if v183 != 0 {
		v87 = v183
		v89 = v180
		v90 = v182
		goto L24
	} else {
		goto L50
	}
L47:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[0]))
	if int32(3) < v168 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v89
	F__serverLog(m, int32(3), int32(_a_F_VM_SetCommandInfo_10), v14+int32(16))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	goto L5
L50:
	;
	goto L25
L51:
	;
	if v196 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	goto L5
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[1])) = int32(28)
	v835 = int32(1)
	goto L3
L54:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v234 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L55:
	;
	goto L65
L56:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v213)+8))
	if v215 != 0 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	if v216 != 0 {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213)+32))
	if v217 != 0 {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v213)+40))
	if v218 != 0 {
		goto L55
	} else {
		goto L60
	}
L60:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v213)+96))
	if v219 != 0 {
		goto L55
	} else {
		goto L61
	}
L61:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v213)+80))
	switch v220 {
	case 0:
		goto L54
	case 1:
		goto L62
	default:
		goto L55
	}
L62:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
	if v222 != int32(2) {
		goto L55
	} else {
		goto L63
	}
L63:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v221)+28))
	if v225 == int32(2) {
		goto L54
	} else {
		goto L64
	}
L64:
	;
	goto L55
L65:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[1])) = int32(20)
	v835 = int32(1)
	goto L3
L66:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v240 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v237 = F_zstrdup(m, v234)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+4)) = v237
	goto L66
L69:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v246 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v243 = F_zstrdup(m, v240)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+8)) = v243
	goto L69
L72:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v253 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v249 = F_zstrdup(m, v246)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+12)) = v249
	goto L72
L75:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v344 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L76:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	v262 = int32(0)
	goto L77
L77:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v253+v256*v262)))
	if v273 != 0 {
		v262 = v262 + int32(1)
		goto L77
	} else {
		goto L79
	}
L78:
	;
	if base.Ui32(int32(536870911)) <= base.Ui32(v262) {
		goto L2
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v277 = v262 << (uint(int32(3)) % 32)
	v280 = F_valkey_malloc(m, v277+int32(8))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+32)) = v280
	if v262 == int32(0) {
		v321 = v280
		goto L82
	} else {
		goto L83
	}
L82:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v321+v277))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v213)+36)) = v262
	goto L75
L83:
	;
	v290 = int32(0)
	goto L84
L84:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	v300 = v297 + v298*v290
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	v302 = F_zstrdup(m, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L9
	} else {
		goto L86
	}
L85:
	;
	v321 = v312
	goto L82
L86:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v213)+32))
	v306 = v290 << (uint(int32(3)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v304+v306))) = v302
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	v310 = F_zstrdup(m, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v213)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v312+v306)+4)) = v310
	v316 = v290 + int32(1)
	if v316 != v262 {
		v290 = v316
		goto L84
	} else {
		goto L88
	}
L88:
	;
	goto L85
L89:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v476 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L90:
	;
	if v344&int32(3) == int32(0) {
		v368 = v344
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v406 = F_sdssplitlen(m, v344, v401, int32(_a_F_VM_SetCommandInfo_11), int32(1), v14+int32(92))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L9
	} else {
		goto L107
	}
L92:
	;
	v401 = v393 - v344
	goto L91
L93:
	;
	v372 = v368
	goto L101
L94:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	if v354 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v357 = v344
	goto L97
L96:
	;
	v401 = v344 - v344
	goto L91
L97:
	;
	v361 = v357 + int32(1)
	if v361&int32(3) == int32(0) {
		v368 = v361
		goto L93
	} else {
		goto L99
	}
L99:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	if v366 != 0 {
		v357 = v361
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v393 = v361
	goto L92
L101:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	v381 = int32(-2139062144)
	if (int32(16843008)-v378|v378)&v381 == v381 {
		v372 = v372 + int32(4)
		goto L101
	} else {
		goto L103
	}
L102:
	;
	v387 = v372
	goto L104
L103:
	;
	goto L102
L104:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v391 != 0 {
		v387 = v387 + int32(1)
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v393 = v387
	goto L92
L106:
	;
	goto L105
L107:
	;
	if v406 == int32(0) {
		goto L89
	} else {
		goto L108
	}
L108:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	v415 = F_valkey_malloc(m, v410<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L9
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+40)) = v415
	v418 = int32(0)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	if v419 <= v418 {
		v450 = v419
		v454 = v415
		goto L110
	} else {
		goto L111
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v454+v450<<(uint(int32(2))%32)))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v213)+44)) = v450
	F_sdsfreesplitres(m, v406, v450)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L9
	} else {
		goto L116
	}
L111:
	;
	v429 = v418
	goto L112
L112:
	;
	v434 = v429 << (uint(int32(2)) % 32)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v406+v434)))
	v437 = F_zstrdup(m, v436)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L9
	} else {
		goto L114
	}
L113:
	;
	v450 = v444
	v454 = v439
	goto L110
L114:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v213)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v439+v434))) = v437
	v443 = v429 + int32(1)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	if v443 < v444 {
		v429 = v443
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	goto L89
L117:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v480 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+52)) = v476
	goto L117
L119:
	;
	v818 = int32(0)
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v819 == v818 {
		v835 = v818
		goto L3
	} else {
		goto L174
	}
L120:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v252)+8))
	v494 = int32(0)
	goto L121
L121:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v480+int32(16)+v485*v494)))
	if v502 != 0 {
		v494 = v494 + int32(1)
		goto L121
	} else {
		goto L123
	}
L122:
	;
	if base.Ui32(int32(2147483647)) <= base.Ui32(v494) {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	F_valkey_free(m, v505)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	v510 = F_valkey_malloc(m, v494*int32(48))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L9
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+80)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v213)+76)) = v510
	if v494 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	F_populateCommandLegacyRangeSpec(m, v213)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L9
	} else {
		goto L173
	}
L128:
	;
	v520 = int32(0)
	goto L129
L129:
	;
	v528 = int32(0)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v252)+8))
	v532 = v529 + v530*v520
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	if v533 == v528 {
		v538 = v528
		goto L131
	} else {
		goto L132
	}
L130:
	;
	goto L127
L131:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	v541 = v520 * int32(48)
	v542 = v539 + v541
	*(*int32)(unsafe.Add(mBase, uint32(v542))) = v538
	v544 = *(*int64)(unsafe.Add(mBase, uint32(v532)+8))
	v555 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[2]))
	if base.B2i32(v555&v544 == int64(0)) == int32(0) {
		goto L136
	} else {
		goto L137
	}
L132:
	;
	v536 = F_zstrdup(m, v533)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L9
	} else {
		goto L133
	}
L133:
	;
	v538 = v536
	goto L131
L134:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v542)+8)) = v707
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v532)+16))
	switch v709 + int32(-1) {
	case 0:
		goto L159
	case 1:
		goto L162
	case 2:
		goto L161
	default:
		goto L160
	}
L135:
	;
	v570 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[3]))
	if v570&v544 == int64(0) {
		v580 = v567
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v566 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[4]))
	v567 = v566
	goto L135
L137:
	;
	v567 = int64(0)
	goto L135
L138:
	;
	v585 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[5]))
	if v585&v544 == int64(0) {
		v595 = v580
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v578 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[6]))
	v580 = v578 | v567
	goto L138
L140:
	;
	v598 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[7]))
	if v598&v544 == int64(0) {
		v608 = v595
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v593 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[8]))
	v595 = v593 | v580
	goto L140
L142:
	;
	v613 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[9]))
	if v613&v544 == int64(0) {
		v623 = v608
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v606 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[10]))
	v608 = v606 | v595
	goto L142
L144:
	;
	v626 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[11]))
	if v626&v544 == int64(0) {
		v636 = v623
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v621 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[12]))
	v623 = v621 | v608
	goto L144
L146:
	;
	v641 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[13]))
	if v641&v544 == int64(0) {
		v651 = v636
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v634 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[14]))
	v636 = v634 | v623
	goto L146
L148:
	;
	v654 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[15]))
	if v654&v544 == int64(0) {
		v664 = v651
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v649 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[16]))
	v651 = v649 | v636
	goto L148
L150:
	;
	v669 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[17]))
	if v669&v544 == int64(0) {
		v679 = v664
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v662 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[18]))
	v664 = v662 | v651
	goto L150
L152:
	;
	v682 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[19]))
	if v682&v544 == int64(0) {
		v692 = v679
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v677 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[20]))
	v679 = v677 | v664
	goto L152
L154:
	;
	v697 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[21]))
	if v697&v544 == int64(0) {
		v707 = v692
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v690 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[22]))
	v692 = v690 | v679
	goto L154
L156:
	;
	goto L134
L157:
	;
	v705 = *(*int64)(unsafe.Add(mBase, _c_F_VM_SetCommandInfo[23]))
	v707 = v705 | v692
	goto L156
L158:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v532)+28))
	switch v738 {
	case 0:
		goto L166
	case 1:
		goto L170
	case 2:
		goto L169
	case 3:
		goto L168
	default:
		goto L167
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+16)) = int32(1)
	goto L158
L160:
	;
	F__serverPanic_1(m, int32(_a_F_VM_SetCommandInfo_1), int32(2083), int32(_a_F_VM_SetCommandInfo_12), int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L9
	} else {
		goto L164
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+16)) = int32(3)
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v532)+20))
	v719 = F_zstrdup(m, v718)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L9
	} else {
		goto L163
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+16)) = int32(2)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v532)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v542)+20)) = v714
	goto L158
L163:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v721+v541)+20)) = v719
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v532)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v724+v541)+24)) = v726
	goto L158
L164:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	v792 = v520 + int32(1)
	if v792 != v494 {
		v520 = v792
		goto L129
	} else {
		goto L172
	}
L166:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	*(*int64)(unsafe.Add(mBase, uint32(v778+v541)+28)) = int64(2)
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v782+v541)+36)) = int32(1)
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v786+v541)+40)) = int32(0)
	goto L165
L167:
	;
	F__serverPanic_1(m, int32(_a_F_VM_SetCommandInfo_1), int32(2109), int32(_a_F_VM_SetCommandInfo_13), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L9
	} else {
		goto L171
	}
L168:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	v758 = v757 + v541
	*(*int32)(unsafe.Add(mBase, uint32(v758)+28)) = int32(3)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v532)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v758)+32)) = v761
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v532)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v763+v541)+36)) = v765
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v532)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v767+v541)+40)) = v769
	goto L165
L169:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	v744 = v743 + v541
	*(*int32)(unsafe.Add(mBase, uint32(v744)+28)) = int32(2)
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v532)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v744)+32)) = v747
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v532)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v749+v541)+36)) = v751
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v532)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v753+v541)+40)) = v755
	goto L165
L170:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v213)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v739+v541)+28)) = int32(1)
	goto L165
L171:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L172:
	;
	goto L130
L173:
	;
	goto L119
L174:
	;
	v822 = F_moduleCopyCommandArgs(m, v819, v252)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L9
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+96)) = v822
	v825 = F_populateArgsStructure(m, v822)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L9
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+88)) = v825
	v835 = v818
	goto L3
L177:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L178:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_VM_SetDisconnectCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	return
}
func F_VM_SetModuleAttribs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int64
	_ = v36
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 != 0 {
		return
	} else {
		v8 = F_valkey_malloc(m, int32(84))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = F_sdsnew(m, l1)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v10
				v15 = F_listCreate(m)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v15
					v18 = F_listCreate(m)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v18
						v21 = F_listCreate(m)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v21
							v24 = F_listCreate(m)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v24
								v27 = F_listCreate(m)
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v27
									*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = int32(563)
									*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = int32(564)
									v36 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v8+int32(44)))) = v36
									*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = v36
									*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v36
									*(*int64)(unsafe.Add(mBase, uint32(v8+int32(64)))) = v36
									*(*int64)(unsafe.Add(mBase, uint32(v8)+72)) = int64(1)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8
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
func F_VM_SignalModifiedKey(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+96))
	F_signalModifiedKey(m, v3, v4, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_VM_Strdup(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_zstrdup(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_VM_StreamDelete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamDelete[0])) = int32(28)
		v52 = int32(1)
		m.G0 = v7 + int32(16)
		return v52
	} else {
		if l1 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v14 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamDelete[0])) = int32(138)
				v52 = int32(1)
				m.G0 = v7 + int32(16)
				return v52
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				if v17&int32(15) == int32(6) {
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
					if v25&int32(2) == int32(0) {
						*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamDelete[0])) = int32(8)
						v52 = int32(1)
						m.G0 = v7 + int32(16)
						return v52
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v30 == int32(0) {
							v36 = F_objectGetVal(m, v14)
							mBase = m.M
							v37 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
							*(*int64)(unsafe.Add(mBase, uint32(v7))) = v37
							v39 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v39
							v42 = F_streamDeleteItem(m, v36, v7)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								if v42 != 0 {
									v52 = int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamDelete[0])) = int32(44)
									v52 = int32(1)
								}
								m.G0 = v7 + int32(16)
								return v52
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamDelete[0])) = int32(8)
							v52 = int32(1)
							m.G0 = v7 + int32(16)
							return v52
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamDelete[0])) = int32(138)
					v52 = int32(1)
					m.G0 = v7 + int32(16)
					return v52
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamDelete[0])) = int32(28)
			v52 = int32(1)
			m.G0 = v7 + int32(16)
			return v52
		}
	}
}
func F_VM_StreamIteratorStart(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v88 int32
	_ = v88
	var v92 int64
	_ = v92
	var v95 int64
	_ = v95
	var v98 int64
	_ = v98
	var v105 int64
	_ = v105
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int64
	_ = v141
	var v152 int32
	_ = v152
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorStart[0])) = int32(28)
		v152 = int32(1)
		m.G0 = v12 + int32(32)
		return v152
	} else {
		if base.Ui32(l1) < base.Ui32(int32(4)) {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v22 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorStart[0])) = int32(138)
				v152 = int32(1)
				m.G0 = v12 + int32(32)
				return v152
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				if v25&int32(15) == int32(6) {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v34 == int32(0) {
						if l2 == int32(0) {
						} else {
							v43 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							v44 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v44
							*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v43
						}
						if l3 == int32(0) {
						} else {
							v50 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
							v51 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v51
							*(*int64)(unsafe.Add(mBase, uint32(v12))) = v50
						}
						if l1&int32(1) == int32(0) {
							v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v125 = F_objectGetVal(m, v124)
							mBase = m.M
							v127 = F_valkey_malloc(m, int32(448))
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int32(0)
							} else {
								if l2 != 0 {
									v134 = v12 + int32(16)
								} else {
									v134 = int32(0)
								}
								if l3 != 0 {
									v136 = v12
								} else {
									v136 = int32(0)
								}
								F_streamIteratorStart(m, v127, v125, v134, v136, l1&int32(2))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									v141 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v141
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v127
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(32)))) = v141
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(40)))) = v141
									v152 = int32(0)
									m.G0 = v12 + int32(32)
									return v152
								}
							}
						} else {
							if l2 == int32(0) {
								if l3 == int32(0) {
									v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v125 = F_objectGetVal(m, v124)
									mBase = m.M
									v127 = F_valkey_malloc(m, int32(448))
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return int32(0)
									} else {
										if l2 != 0 {
											v134 = v12 + int32(16)
										} else {
											v134 = int32(0)
										}
										if l3 != 0 {
											v136 = v12
										} else {
											v136 = int32(0)
										}
										F_streamIteratorStart(m, v127, v125, v134, v136, l1&int32(2))
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int32(0)
										} else {
											v141 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v141
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v127
											*(*int64)(unsafe.Add(mBase, uint32(l0+int32(32)))) = v141
											*(*int64)(unsafe.Add(mBase, uint32(l0+int32(40)))) = v141
											v152 = int32(0)
											m.G0 = v12 + int32(32)
											return v152
										}
									}
								} else {
									v92 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
									if v92 != int64(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v92 + int64(-1)
										v116 = int32(0)
									} else {
										v95 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
										if v95 != int64(0) {
											v105 = int64(-1)
											*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v105
											*(*int64)(unsafe.Add(mBase, uint32(v12))) = v95 + v105
											v116 = int32(0)
										} else {
											v98 = int64(-1)
											*(*int64)(unsafe.Add(mBase, uint32(v12))) = v98
											*(*int64)(unsafe.Add(mBase, uint32(v12+int32(8)))) = v98
											v116 = int32(-1)
										}
									}
									if v116 == int32(0) {
										v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v125 = F_objectGetVal(m, v124)
										mBase = m.M
										v127 = F_valkey_malloc(m, int32(448))
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int32(0)
										} else {
											if l2 != 0 {
												v134 = v12 + int32(16)
											} else {
												v134 = int32(0)
											}
											if l3 != 0 {
												v136 = v12
											} else {
												v136 = int32(0)
											}
											F_streamIteratorStart(m, v127, v125, v134, v136, l1&int32(2))
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												v141 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v141
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v127
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(32)))) = v141
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(40)))) = v141
												v152 = int32(0)
												m.G0 = v12 + int32(32)
												return v152
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorStart[0])) = int32(18)
										v152 = int32(1)
										m.G0 = v12 + int32(32)
										return v152
									}
								}
							} else {
								v62 = v12 + int32(16)
								v64 = *(*int64)(unsafe.Add(mBase, uint32(v62)+8))
								if v64 != int64(-1) {
									*(*int64)(unsafe.Add(mBase, uint32(v62)+8)) = v64 + int64(1)
									v88 = int32(0)
								} else {
									v67 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
									if v67 != int64(-1) {
										*(*int64)(unsafe.Add(mBase, uint32(v62)+8)) = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v62))) = v67 + int64(1)
										v88 = int32(0)
									} else {
										v70 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v62))) = v70
										*(*int64)(unsafe.Add(mBase, uint32(v12+int32(24)))) = v70
										v88 = int32(-1)
									}
								}
								if v88 != 0 {
									*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorStart[0])) = int32(18)
									v152 = int32(1)
									m.G0 = v12 + int32(32)
									return v152
								} else {
									if l3 == int32(0) {
										v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v125 = F_objectGetVal(m, v124)
										mBase = m.M
										v127 = F_valkey_malloc(m, int32(448))
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int32(0)
										} else {
											if l2 != 0 {
												v134 = v12 + int32(16)
											} else {
												v134 = int32(0)
											}
											if l3 != 0 {
												v136 = v12
											} else {
												v136 = int32(0)
											}
											F_streamIteratorStart(m, v127, v125, v134, v136, l1&int32(2))
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												v141 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v141
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v127
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(32)))) = v141
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(40)))) = v141
												v152 = int32(0)
												m.G0 = v12 + int32(32)
												return v152
											}
										}
									} else {
										v92 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
										if v92 != int64(0) {
											*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v92 + int64(-1)
											v116 = int32(0)
										} else {
											v95 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
											if v95 != int64(0) {
												v105 = int64(-1)
												*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v105
												*(*int64)(unsafe.Add(mBase, uint32(v12))) = v95 + v105
												v116 = int32(0)
											} else {
												v98 = int64(-1)
												*(*int64)(unsafe.Add(mBase, uint32(v12))) = v98
												*(*int64)(unsafe.Add(mBase, uint32(v12+int32(8)))) = v98
												v116 = int32(-1)
											}
										}
										if v116 == int32(0) {
											v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v125 = F_objectGetVal(m, v124)
											mBase = m.M
											v127 = F_valkey_malloc(m, int32(448))
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return int32(0)
											} else {
												if l2 != 0 {
													v134 = v12 + int32(16)
												} else {
													v134 = int32(0)
												}
												if l3 != 0 {
													v136 = v12
												} else {
													v136 = int32(0)
												}
												F_streamIteratorStart(m, v127, v125, v134, v136, l1&int32(2))
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													v141 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v141
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v127
													*(*int64)(unsafe.Add(mBase, uint32(l0+int32(32)))) = v141
													*(*int64)(unsafe.Add(mBase, uint32(l0+int32(40)))) = v141
													v152 = int32(0)
													m.G0 = v12 + int32(32)
													return v152
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorStart[0])) = int32(18)
											v152 = int32(1)
											m.G0 = v12 + int32(32)
											return v152
										}
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorStart[0])) = int32(8)
						v152 = int32(1)
						m.G0 = v12 + int32(32)
						return v152
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorStart[0])) = int32(138)
					v152 = int32(1)
					m.G0 = v12 + int32(32)
					return v152
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorStart[0])) = int32(28)
			v152 = int32(1)
			m.G0 = v12 + int32(32)
			return v152
		}
	}
}
func F_VM_StringPtrLen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	if l0 != 0 {
		if l1 == int32(0) {
		} else {
			v16 = F_objectGetVal(m, l0)
			mBase = m.M
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-1)))))
			switch v19 & int32(7) {
			case 0:
				v36 = int32(base.Ui32(v19) >> (uint(int32(3)) % 32))
			case 1:
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
				v36 = v26
			case 2:
				v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
				v36 = v29
			case 3:
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
				v36 = v32
			case 4:
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
				v36 = v35
			default:
				v36 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v36
		}
		v41 = F_objectGetVal(m, l0)
		mBase = m.M
		v42 = v41
		return v42
	} else {
		if l1 == int32(0) {
			v42 = int32(_a_F_VM_StringPtrLen_0)
			return v42
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(40)
			return int32(_a_F_VM_StringPtrLen_0)
		}
	}
}
func F_VM_StringToStreamID(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_streamParseID(m, l0, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v19 = int32(1)
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v14
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v16
			v19 = int32(0)
		}
		m.G0 = v7 + int32(16)
		return v19
	}
}
func F_VM_SubscribeToKeyspaceEvents(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v6 = F_valkey_malloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
		v13 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l2
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToKeyspaceEvents[0]))
		v18 = F_listAddNodeTail(m, v17, v6)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_VM_SubscribeToServerEvent(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v25 int64
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int64
	_ = v94
	var v100 int64
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v184 int32
	_ = v184
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == int32(0) {
		v184 = v12
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v184
L2:
	;
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui64(int64(23)) < base.Ui64(v16) {
		v184 = v12
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v16)<<(uint(int32(3))%32))+uint32(_c_F_VM_SubscribeToServerEvent[0])))
	if base.Ui64(v25) < base.Ui64(v19) {
		v184 = v12
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = int32(0)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[1]))
	v30 = v10 + int32(8)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v31
	goto L5
L5:
	;
	v36 = v10 + int32(8)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v38 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v184 = int32(0)
	goto L1
L7:
	;
	v166 = int32(0)
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[2])) = v168 + int32(1)
	goto L6
L8:
	;
	v160 = int32(0)
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[3])) = v162 + int32(1)
	goto L6
L9:
	;
	v154 = int32(0)
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[4])) = v156 + int32(1)
	goto L6
L10:
	;
	v148 = int32(0)
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[5])) = v150 + int32(1)
	goto L6
L11:
	;
	if l2 != 0 {
		goto L31
	} else {
		goto L32
	}
L12:
	;
	if l2 == int32(0) {
		goto L6
	} else {
		goto L26
	}
L13:
	;
	if v38 == int32(0) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38+base.B2i32(v41 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v47
	goto L14
L16:
	;
	v55 = v38
	goto L17
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v59 != v60 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L12
L19:
	;
	v65 = v10 + int32(8)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v67 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v58)+8))
	if v62 == v16 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if v67 != 0 {
		v55 = v67
		goto L17
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v67+base.B2i32(v70 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v76
	goto L23
L25:
	;
	goto L18
L26:
	;
	v88 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v92
	v94 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v88)+8)) = v94
	v100 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v88+int32(16)))) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = l2
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[1]))
	v105 = F_listAddNodeTail(m, v104, v88)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v108 = v16 + int64(-20)
	if base.Ui64(int64(3)) < base.Ui64(v108) {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	switch base.I32_wrap_i64(v108) {
	default:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	case 3:
		goto L7
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = l2
	goto L6
L32:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[1]))
	F_listDelNode(m, v113, v55)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	F_valkey_free(m, v58)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v119 = v16 + int64(-20)
	if base.Ui64(int64(3)) < base.Ui64(v119) {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	switch base.I32_wrap_i64(v119) {
	default:
		goto L39
	case 1:
		goto L38
	case 2:
		goto L37
	case 3:
		goto L36
	}
L36:
	;
	v141 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[2])) = v143 + int32(-1)
	goto L6
L37:
	;
	v135 = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[3])) = v137 + int32(-1)
	goto L6
L38:
	;
	v129 = int32(0)
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[4])) = v131 + int32(-1)
	goto L6
L39:
	;
	v123 = int32(0)
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_VM_SubscribeToServerEvent[5])) = v125 + int32(-1)
	goto L6
}
func F_VM_ThreadSafeContextLock(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v34 int64
	_ = v34
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v49 int64
	_ = v49
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ThreadSafeContextLock[0]))
	if v5 == int32(0) {
		v16 = int32(0)
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ThreadSafeContextLock[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_VM_ThreadSafeContextLock[0])) = v20 + int32(1)
		if v20 != 0 {
		} else {
			v28 = F_ustime(m)
			mBase = m.M
			v30 = int32(0)
			*(*int64)(unsafe.Add(mBase, _c_F_VM_ThreadSafeContextLock[1])) = v28
			v34 = base.I64_div_s(v28, int64(1000))
			*(*int64)(unsafe.Add(mBase, _c_F_VM_ThreadSafeContextLock[2])) = v34
			v38 = base.I64_div_s(v28, int64(1000000))
			*(*int64)(unsafe.Add(mBase, _c_F_VM_ThreadSafeContextLock[3])) = v38
			v41 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ThreadSafeContextLock[4]))
			F_lrulfu_updateClockAndPolicy(m, v34, int32(base.Ui32(v41&int32(2))>>(uint(int32(1))%32)))
			mBase = m.M
			v49 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ThreadSafeContextLock[2]))
			*(*int64)(unsafe.Add(mBase, _c_F_VM_ThreadSafeContextLock[5])) = v49
		}
		return
	} else {
		F__serverAssert(m, int32(_a_F_VM_ThreadSafeContextLock_0), int32(_a_F_VM_ThreadSafeContextLock_1), int32(9239))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
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
func F_VM_ThreadSafeContextUnlock(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ThreadSafeContextUnlock[0]))
	if v3 == int32(1) {
		v12 = int32(0)
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ThreadSafeContextUnlock[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_VM_ThreadSafeContextUnlock[0])) = v14 + int32(-1)
		F_postExecutionUnitOperations(m)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			return
		}
	} else {
		F__serverAssert(m, int32(_a_F_VM_ThreadSafeContextUnlock_0), int32(_a_F_VM_ThreadSafeContextUnlock_1), int32(9276))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
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
func F_VM_TryRealloc(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ztryrealloc_usable(m, l0, l1, int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_VM_UnblockClient(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int64
	_ = v80
	var v84 int64
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	if l0 != 0 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v16 == int32(0) {
			v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v114 != 0 {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l1
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
			v118 = int32(0)
			v120 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[0]))
			v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
			if v121 != 0 {
				v129 = v120
			} else {
				v123 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[1]))
				v126 = F_write(m, v123, int32(_a_F_VM_UnblockClient_0), int32(1))
				mBase = m.M
				v128 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[0]))
				v129 = v128
			}
			v130 = F_listAddNodeTail(m, v129, l0)
			mBase = m.M
			v131 = m.ExcPending
			if v131 != 0 {
				return int32(0)
			} else {
				v135 = v118
				m.G0 = v10 + int32(80)
				return v135
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v19 != 0 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if v25 != 0 {
					v135 = int32(0)
					m.G0 = v10 + int32(80)
					return v135
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v26 == int32(0) {
						v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						if v114 != 0 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l1
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
						v118 = int32(0)
						v120 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[0]))
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
						if v121 != 0 {
							v129 = v120
						} else {
							v123 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[1]))
							v126 = F_write(m, v123, int32(_a_F_VM_UnblockClient_0), int32(1))
							mBase = m.M
							v128 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[0]))
							v129 = v128
						}
						v130 = F_listAddNodeTail(m, v129, l0)
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return int32(0)
						} else {
							v135 = v118
							m.G0 = v10 + int32(80)
							return v135
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+116))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
						if v31 != 0 {
							v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v114 != 0 {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l1
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
							v118 = int32(0)
							v120 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[0]))
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
							if v121 != 0 {
								v129 = v120
							} else {
								v123 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[1]))
								v126 = F_write(m, v123, int32(_a_F_VM_UnblockClient_0), int32(1))
								mBase = m.M
								v128 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[0]))
								v129 = v128
							}
							v130 = F_listAddNodeTail(m, v129, l0)
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int32(0)
							} else {
								v135 = v118
								m.G0 = v10 + int32(80)
								return v135
							}
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
							v37 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v10+int32(32)))) = v37
							*(*int64)(unsafe.Add(mBase, uint32(v10+int32(72)))) = v37
							*(*int64)(unsafe.Add(mBase, uint32(v10+int32(64)))) = v37
							*(*int64)(unsafe.Add(mBase, uint32(v10+int32(56)))) = v37
							*(*int64)(unsafe.Add(mBase, uint32(v10+int32(48)))) = v37
							*(*int64)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v37
							*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v37
							*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v32
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(561)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(24)
							v67 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[2]))
							v68 = int32(0)
							v69 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[3]))
							v70 = m.T0[v69].(func(*base.Module) int64)(m)
							mBase = m.M
							if v67 == v68 {
								v80 = *(*int64)(unsafe.Add(mBase, _c_F_VM_UnblockClient[4]))
								v84 = v80*int64(1000) + v70
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[5]))
								v76 = base.I32_div_s(int32(1000000), v75)
								v84 = v70 + base.I64_extend_i32_s(v76)
							}
							*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v84
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v30
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v86
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v89
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
							if v91 == int32(0) {
								F_moduleFreeContext(m, v10+int32(8))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = int32(0)
									v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v114 != 0 {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l1
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
									v118 = int32(0)
									v120 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[0]))
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
									if v121 != 0 {
										v129 = v120
									} else {
										v123 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[1]))
										v126 = F_write(m, v123, int32(_a_F_VM_UnblockClient_0), int32(1))
										mBase = m.M
										v128 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[0]))
										v129 = v128
									}
									v130 = F_listAddNodeTail(m, v129, l0)
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int32(0)
									} else {
										v135 = v118
										m.G0 = v10 + int32(80)
										return v135
									}
								}
							} else {
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
								v98 = m.T0[v91].(func(*base.Module, int32, int32, int32) int32)(m, v10+int32(8), v96, v97)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									F_moduleFreeContext(m, v10+int32(8))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = int32(0)
										v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v114 != 0 {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l1
										}
										*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
										v118 = int32(0)
										v120 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[0]))
										v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
										if v121 != 0 {
											v129 = v120
										} else {
											v123 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[1]))
											v126 = F_write(m, v123, int32(_a_F_VM_UnblockClient_0), int32(1))
											mBase = m.M
											v128 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[0]))
											v129 = v128
										}
										v130 = F_listAddNodeTail(m, v129, l0)
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return int32(0)
										} else {
											v135 = v118
											m.G0 = v10 + int32(80)
											return v135
										}
									}
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[6])) = int32(138)
				v135 = int32(1)
				m.G0 = v10 + int32(80)
				return v135
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_VM_UnblockClient[6])) = int32(28)
		v135 = int32(1)
		m.G0 = v10 + int32(80)
		return v135
	}
}
func F_VM_UnlinkKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v16 int32
	_ = v16
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v2&int32(2) != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v12 = F_dbAsyncDelete(m, v10, v11)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v16
				return v16
			}
		} else {
			return int32(0)
		}
	} else {
		return int32(1)
	}
}
func F_VM_UnregisterScriptingEngine(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_scriptingEngineManagerUnregister(m, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v3 != int32(0))
	}
}
func F_VM_Yield(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v45 int64
	_ = v45
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v60 int64
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
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
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int64
	_ = v124
	var v128 int64
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0]))
	if v7 != 0 {
		return
	} else {
		v8 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0])) = int32(1)
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[1]))
		v13 = m.T0[v12].(func(*base.Module) int64)(m)
		mBase = m.M
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		if v13 < v14 {
			v133 = int32(0)
			v135 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0])) = v135 + int32(-1)
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[2]))
			if v17 == int32(0) {
				v22 = int32(_a_F_VM_Yield_0)
				v23 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[3]))
				*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[3])) = l2
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4]))
				if v27 != 0 {
					if l1&int32(2) == int32(0) {
					} else {
						v75 = int32(_a_F_VM_Yield_0)
						v77 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4]))
						*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4])) = v77 | int32(2)
					}
					v82 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[5]))
					v83 = F___get_tp(m)
					mBase = m.M
					if v82 == v83 {
						F_processEventsWhileBlocked(m)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v100 = int32(_a_F_VM_Yield_0)
							*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[3])) = v23
							v104 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4]))
							*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4])) = v104 & int32(-3)
							v111 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[2]))
							v112 = int32(0)
							v113 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[1]))
							v114 = m.T0[v113].(func(*base.Module) int64)(m)
							mBase = m.M
							if v111 == v112 {
								v124 = *(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[6]))
								v128 = v124*int64(1000) + v114
							} else {
								v119 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[7]))
								v120 = base.I32_div_s(int32(1000000), v119)
								v128 = v114 + base.I64_extend_i32_s(v120)
							}
							*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v128
							v133 = int32(0)
							v135 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0])) = v135 + int32(-1)
							return
						}
					} else {
						v86 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[8]))
						if v86 != 0 {
							v94 = F_sched_yield(m)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								v100 = int32(_a_F_VM_Yield_0)
								*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[3])) = v23
								v104 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4]))
								*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4])) = v104 & int32(-3)
								v111 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[2]))
								v112 = int32(0)
								v113 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[1]))
								v114 = m.T0[v113].(func(*base.Module) int64)(m)
								mBase = m.M
								if v111 == v112 {
									v124 = *(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[6]))
									v128 = v124*int64(1000) + v114
								} else {
									v119 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[7]))
									v120 = base.I32_div_s(int32(1000000), v119)
									v128 = v114 + base.I64_extend_i32_s(v120)
								}
								*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v128
								v133 = int32(0)
								v135 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0]))
								*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0])) = v135 + int32(-1)
								return
							}
						} else {
							v88 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[9]))
							v91 = F_write(m, v88, int32(_a_F_VM_Yield_1), int32(1))
							mBase = m.M
							v100 = int32(_a_F_VM_Yield_0)
							*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[3])) = v23
							v104 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4]))
							*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4])) = v104 & int32(-3)
							v111 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[2]))
							v112 = int32(0)
							v113 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[1]))
							v114 = m.T0[v113].(func(*base.Module) int64)(m)
							mBase = m.M
							if v111 == v112 {
								v124 = *(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[6]))
								v128 = v124*int64(1000) + v114
							} else {
								v119 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[7]))
								v120 = base.I32_div_s(int32(1000000), v119)
								v128 = v114 + base.I64_extend_i32_s(v120)
							}
							*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v128
							v133 = int32(0)
							v135 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0])) = v135 + int32(-1)
							return
						}
					}
				} else {
					v29 = int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4])) = v29
					v31 = int32(0)
					v36 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[10]))
					*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[10])) = v36 + v29
					if v36 != 0 {
					} else {
						v40 = int32(0)
						v41 = F_ustime(m)
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[11])) = v41
						v45 = base.I64_div_s(v41, int64(1000))
						*(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[12])) = v45
						v49 = base.I64_div_s(v41, int64(1000000))
						*(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[13])) = v49
						v52 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[14]))
						F_lrulfu_updateClockAndPolicy(m, v45, int32(base.Ui32(v52&int32(2))>>(uint(int32(1))%32)))
						mBase = m.M
						v60 = *(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[12]))
						*(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[15])) = v60
					}
					v65 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[16]))
					if v65 == int32(0) {
						if l1&int32(2) == int32(0) {
						} else {
							v75 = int32(_a_F_VM_Yield_0)
							v77 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4]))
							*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4])) = v77 | int32(2)
						}
						v82 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[5]))
						v83 = F___get_tp(m)
						mBase = m.M
						if v82 == v83 {
							F_processEventsWhileBlocked(m)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								v100 = int32(_a_F_VM_Yield_0)
								*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[3])) = v23
								v104 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4]))
								*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4])) = v104 & int32(-3)
								v111 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[2]))
								v112 = int32(0)
								v113 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[1]))
								v114 = m.T0[v113].(func(*base.Module) int64)(m)
								mBase = m.M
								if v111 == v112 {
									v124 = *(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[6]))
									v128 = v124*int64(1000) + v114
								} else {
									v119 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[7]))
									v120 = base.I32_div_s(int32(1000000), v119)
									v128 = v114 + base.I64_extend_i32_s(v120)
								}
								*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v128
								v133 = int32(0)
								v135 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0]))
								*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0])) = v135 + int32(-1)
								return
							}
						} else {
							v86 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[8]))
							if v86 != 0 {
								v94 = F_sched_yield(m)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									v100 = int32(_a_F_VM_Yield_0)
									*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[3])) = v23
									v104 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4]))
									*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4])) = v104 & int32(-3)
									v111 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[2]))
									v112 = int32(0)
									v113 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[1]))
									v114 = m.T0[v113].(func(*base.Module) int64)(m)
									mBase = m.M
									if v111 == v112 {
										v124 = *(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[6]))
										v128 = v124*int64(1000) + v114
									} else {
										v119 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[7]))
										v120 = base.I32_div_s(int32(1000000), v119)
										v128 = v114 + base.I64_extend_i32_s(v120)
									}
									*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v128
									v133 = int32(0)
									v135 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0]))
									*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0])) = v135 + int32(-1)
									return
								}
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[9]))
								v91 = F_write(m, v88, int32(_a_F_VM_Yield_1), int32(1))
								mBase = m.M
								v100 = int32(_a_F_VM_Yield_0)
								*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[3])) = v23
								v104 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4]))
								*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4])) = v104 & int32(-3)
								v111 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[2]))
								v112 = int32(0)
								v113 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[1]))
								v114 = m.T0[v113].(func(*base.Module) int64)(m)
								mBase = m.M
								if v111 == v112 {
									v124 = *(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[6]))
									v128 = v124*int64(1000) + v114
								} else {
									v119 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[7]))
									v120 = base.I32_div_s(int32(1000000), v119)
									v128 = v114 + base.I64_extend_i32_s(v120)
								}
								*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v128
								v133 = int32(0)
								v135 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0]))
								*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0])) = v135 + int32(-1)
								return
							}
						}
					} else {
						F_protectClient(m, v65)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							if l1&int32(2) == int32(0) {
							} else {
								v75 = int32(_a_F_VM_Yield_0)
								v77 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4]))
								*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4])) = v77 | int32(2)
							}
							v82 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[5]))
							v83 = F___get_tp(m)
							mBase = m.M
							if v82 == v83 {
								F_processEventsWhileBlocked(m)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									v100 = int32(_a_F_VM_Yield_0)
									*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[3])) = v23
									v104 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4]))
									*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4])) = v104 & int32(-3)
									v111 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[2]))
									v112 = int32(0)
									v113 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[1]))
									v114 = m.T0[v113].(func(*base.Module) int64)(m)
									mBase = m.M
									if v111 == v112 {
										v124 = *(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[6]))
										v128 = v124*int64(1000) + v114
									} else {
										v119 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[7]))
										v120 = base.I32_div_s(int32(1000000), v119)
										v128 = v114 + base.I64_extend_i32_s(v120)
									}
									*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v128
									v133 = int32(0)
									v135 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0]))
									*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0])) = v135 + int32(-1)
									return
								}
							} else {
								v86 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[8]))
								if v86 != 0 {
									v94 = F_sched_yield(m)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										v100 = int32(_a_F_VM_Yield_0)
										*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[3])) = v23
										v104 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4]))
										*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4])) = v104 & int32(-3)
										v111 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[2]))
										v112 = int32(0)
										v113 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[1]))
										v114 = m.T0[v113].(func(*base.Module) int64)(m)
										mBase = m.M
										if v111 == v112 {
											v124 = *(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[6]))
											v128 = v124*int64(1000) + v114
										} else {
											v119 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[7]))
											v120 = base.I32_div_s(int32(1000000), v119)
											v128 = v114 + base.I64_extend_i32_s(v120)
										}
										*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v128
										v133 = int32(0)
										v135 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0]))
										*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0])) = v135 + int32(-1)
										return
									}
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[9]))
									v91 = F_write(m, v88, int32(_a_F_VM_Yield_1), int32(1))
									mBase = m.M
									v100 = int32(_a_F_VM_Yield_0)
									*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[3])) = v23
									v104 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4]))
									*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[4])) = v104 & int32(-3)
									v111 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[2]))
									v112 = int32(0)
									v113 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[1]))
									v114 = m.T0[v113].(func(*base.Module) int64)(m)
									mBase = m.M
									if v111 == v112 {
										v124 = *(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[6]))
										v128 = v124*int64(1000) + v114
									} else {
										v119 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[7]))
										v120 = base.I32_div_s(int32(1000000), v119)
										v128 = v114 + base.I64_extend_i32_s(v120)
									}
									*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v128
									v133 = int32(0)
									v135 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0]))
									*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0])) = v135 + int32(-1)
									return
								}
							}
						}
					}
				}
			} else {
				F_processEventsWhileBlocked(m)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					v111 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[2]))
					v112 = int32(0)
					v113 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[1]))
					v114 = m.T0[v113].(func(*base.Module) int64)(m)
					mBase = m.M
					if v111 == v112 {
						v124 = *(*int64)(unsafe.Add(mBase, _c_F_VM_Yield[6]))
						v128 = v124*int64(1000) + v114
					} else {
						v119 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[7]))
						v120 = base.I32_div_s(int32(1000000), v119)
						v128 = v114 + base.I64_extend_i32_s(v120)
					}
					*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v128
					v133 = int32(0)
					v135 = *(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0]))
					*(*int32)(unsafe.Add(mBase, _c_F_VM_Yield[0])) = v135 + int32(-1)
					return
				}
			}
		}
	}
}
func F_VM_ZsetFirstInScoreRange(m *base.Module, l0 int32, l1 float64, l2 float64, l3 int32, l4 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_zsetInitScoreRange(m, l0, l1, l2, l3, l4, int32(1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_VM_ZsetLastInScoreRange(m *base.Module, l0 int32, l1 float64, l2 float64, l3 int32, l4 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_zsetInitScoreRange(m, l0, l1, l2, l3, l4, int32(0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_VM_ZsetRangeEndReached(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v4 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v5 == int32(0) {
		v14 = v4
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		if v8&int32(15) != int32(3) {
			v14 = v4
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			v14 = v13
		}
	}
	return v14
}
func F_VM_ZsetRangeNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 float64
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 float64
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v77 float64
	_ = v77
	var v79 int32
	_ = v79
	var v81 float64
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 == v2 {
		v165 = v2
		return v165
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		if v11&int32(15) != int32(3) {
			v165 = v2
			return v165
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v16 == int32(0) {
				v165 = v2
				return v165
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
				if v19 == int32(0) {
					v165 = v2
					return v165
				} else {
					switch int32(base.Ui32(v11)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
					case 0:
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						if v70 != 0 {
							if v16 != int32(2) {
								v89 = v16
								if v89 != int32(1) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v70
									v165 = int32(1)
									return v165
								} else {
									v93 = v70 + int32(16)
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
									v97 = v93 + v94<<(uint(int32(3))%32)
									v98 = int32(*(*int8)(unsafe.Add(mBase, uint32(v97))))
									v101 = v97 + v98 + int32(1)
									v103 = l0 + int32(56)
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
									if v108 == int32(0) {
										v127 = int32(1)
										if v101 == v107 {
											v141 = v127
										} else {
											v130 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangeNext[0]))
											if v101 == v130 {
												v141 = v127
											} else {
												v133 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangeNext[1]))
												if v107 == v133 {
													v141 = v127
												} else {
													v135 = int32(0)
													if v107 == v130 {
														v141 = v135
													} else {
														if v101 == v133 {
															v141 = v135
														} else {
															v138 = F_sdscmp(m, v101, v107)
															mBase = m.M
															v141 = base.B2i32(v138 < int32(1))
														}
													}
												}
											}
										}
										v147 = v141
									} else {
										if v101 != v107 {
											v113 = int32(1)
											v115 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangeNext[0]))
											if v101 == v115 {
												v141 = v113
												v147 = v141
											} else {
												v118 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangeNext[1]))
												if v107 == v118 {
													v141 = v113
													v147 = v141
												} else {
													if v107 != v115 {
														if v101 == v118 {
															v141 = int32(0)
															v147 = v141
														} else {
															v124 = F_sdscmp(m, v101, v107)
															mBase = m.M
															v147 = int32(base.Ui32(v124) >> (uint(int32(31)) % 32))
														}
													} else {
														v147 = int32(0)
													}
												}
											}
										} else {
											v147 = int32(0)
										}
									}
									if v147 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v70
										v165 = int32(1)
										return v165
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
										return int32(0)
									}
								}
							} else {
								v77 = *(*float64)(unsafe.Add(mBase, uint32(v70)))
								v79 = l0 + int32(32)
								v81 = *(*float64)(unsafe.Add(mBase, uint32(v79)+8))
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
								if v84 != 0 {
									v85 = base.F64_lt(v77, v81)
								} else {
									v85 = base.F64_le(v77, v81)
								}
								if v85 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
									return int32(0)
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v89 = v88
									if v89 != int32(1) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v70
										v165 = int32(1)
										return v165
									} else {
										v93 = v70 + int32(16)
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
										v97 = v93 + v94<<(uint(int32(3))%32)
										v98 = int32(*(*int8)(unsafe.Add(mBase, uint32(v97))))
										v101 = v97 + v98 + int32(1)
										v103 = l0 + int32(56)
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
										if v108 == int32(0) {
											v127 = int32(1)
											if v101 == v107 {
												v141 = v127
											} else {
												v130 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangeNext[0]))
												if v101 == v130 {
													v141 = v127
												} else {
													v133 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangeNext[1]))
													if v107 == v133 {
														v141 = v127
													} else {
														v135 = int32(0)
														if v107 == v130 {
															v141 = v135
														} else {
															if v101 == v133 {
																v141 = v135
															} else {
																v138 = F_sdscmp(m, v101, v107)
																mBase = m.M
																v141 = base.B2i32(v138 < int32(1))
															}
														}
													}
												}
											}
											v147 = v141
										} else {
											if v101 != v107 {
												v113 = int32(1)
												v115 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangeNext[0]))
												if v101 == v115 {
													v141 = v113
													v147 = v141
												} else {
													v118 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangeNext[1]))
													if v107 == v118 {
														v141 = v113
														v147 = v141
													} else {
														if v107 != v115 {
															if v101 == v118 {
																v141 = int32(0)
																v147 = v141
															} else {
																v124 = F_sdscmp(m, v101, v107)
																mBase = m.M
																v147 = int32(base.Ui32(v124) >> (uint(int32(31)) % 32))
															}
														} else {
															v147 = int32(0)
														}
													}
												}
											} else {
												v147 = int32(0)
											}
										}
										if v147 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v70
											v165 = int32(1)
											return v165
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
											return int32(0)
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
							return int32(0)
						}
					default:
						F__serverPanic_1(m, int32(_a_F_VM_ZsetRangeNext_0), int32(5313), int32(_a_F_VM_ZsetRangeNext_1), int32(0))
						mBase = m.M
						v158 = m.ExcPending
						if v158 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					case 4:
						v28 = F_objectGetVal(m, v8)
						mBase = m.M
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
						v30 = F_lpNext(m, v28, v29)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							if v30 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
								return int32(0)
							} else {
								v36 = F_lpNext(m, v28, v30)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									if v36 != 0 {
										v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										switch v43 + int32(-1) {
										case 0:
											v64 = F_zzlLexValueLteMax(m, v36, l0+int32(56))
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												if v64 != 0 {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v36
													v165 = int32(1)
													return v165
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
													return int32(0)
												}
											}
										case 1:
											v46 = F_lpNext(m, v28, v36)
											mBase = m.M
											v47 = m.ExcPending
											if v47 != 0 {
												return int32(0)
											} else {
												v48 = F_zzlGetScore(m, v46)
												mBase = m.M
												v49 = m.ExcPending
												if v49 != 0 {
													return int32(0)
												} else {
													v51 = l0 + int32(32)
													v53 = *(*float64)(unsafe.Add(mBase, uint32(v51)+8))
													v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
													if v56 != 0 {
														v57 = base.F64_lt(v48, v53)
													} else {
														v57 = base.F64_le(v48, v53)
													}
													if v57 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v36
														v165 = int32(1)
														return v165
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
														return int32(0)
													}
												}
											}
										default:
											*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v36
											v165 = int32(1)
											return v165
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
										return int32(0)
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
func F_VM_ZsetRem(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v6 = int32(1)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v7&int32(2) == int32(0) {
		v40 = v6
		return v40
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v12 == int32(0) {
			v35 = int32(0)
			if l2 == v35 {
				v40 = v35
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
				v40 = v35
			}
			return v40
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			if v15&int32(15) != int32(3) {
				v40 = v6
				return v40
			} else {
				v20 = F_objectGetVal(m, l1)
				mBase = m.M
				v21 = F_zsetDel(m, v12, v20)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v21 == int32(0) {
						v35 = int32(0)
						if l2 == v35 {
							v40 = v35
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
							v40 = v35
						}
						return v40
					} else {
						if l2 == int32(0) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
						}
						v31 = F_moduleDelKeyIfEmpty(m, l0)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				}
			}
		}
	}
}
func F_VM_ZsetScore(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v5 = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(0) {
		v21 = v5
		return v21
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		if v9&int32(15) != int32(3) {
			v21 = v5
			return v21
		} else {
			v14 = F_objectGetVal(m, l1)
			mBase = m.M
			v15 = F_zsetScore(m, v6, v14, l2)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v21 = base.B2i32(v15 == int32(-1))
				return v21
			}
		}
	}
}
func F_VM__Assert(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v5 int32
	_ = v5
	F__serverAssert(m, l0, l1, l2)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
