package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_addListListpackRangeReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int64
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_prepareClientForFutureWrites(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_addListListpackRangeReply_0), int32(_a_F_addListListpackRangeReply_1), int32(690))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L24
	}
L2:
	;
	m.G0 = v10 + int32(16)
	return
L3:
	;
	return
L4:
	;
	if v12 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	F_addWritePreparedReplyArrayLen(m, v12, l3)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v18 = F_objectGetVal(m, l1)
	mBase = m.M
	v19 = F_lpSeek(m, v18, l2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if l3 == int32(0) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v23 = v19
	v26 = l3
	goto L9
L9:
	;
	if v23 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L2
L11:
	;
	v34 = F_lpGetValue(m, v23, v10+int32(12), v10)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L14
	}
L12:
	;
	v44 = F_objectGetVal(m, l1)
	mBase = m.M
	if l4 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	F_addWritePreparedReplyBulkLongLong(m, v12, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L17
	}
L14:
	;
	if v34 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_addWritePreparedReplyBulkCBuffer(m, v12, v34, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	goto L12
L18:
	;
	v53 = v26 + int32(-1)
	if v53 != 0 {
		v23 = v51
		v26 = v53
		goto L9
	} else {
		goto L23
	}
L19:
	;
	v49 = F_lpNext(m, v44, v23)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L22
	}
L20:
	;
	v47 = F_lpPrev(m, v44, v23)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v51 = v47
	goto L18
L22:
	;
	v51 = v49
	goto L18
L23:
	;
	goto L10
L24:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_listAddNodeTail(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	v6 = F_valkey_malloc(m, int32(12))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v13 != 0 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v6
				v22 = v18
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v6
				v15 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v15
				v22 = v15
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v22
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v13 + int32(1)
			return l0
		} else {
			return int32(0)
		}
	}
}
func F_listGetIterator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	v5 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0+base.B2i32(l1 != int32(0))<<(uint(int32(2))%32))))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v16
		}
		return v5
	}
}
func F_listNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 == int32(0) {
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v3+base.B2i32(v6 == int32(0))<<(uint(int32(2))%32))))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v12
	}
	return v3
}
func F_listReleaseIterator(m *base.Module, l0 int32) {
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
func F_listTypeReplaceAtIndex(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = F_getDecodedObject(m, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = F_objectGetVal(m, v13)
		mBase = m.M
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-1)))))
		switch v20 & int32(7) {
		case 0:
			v37 = int32(base.Ui32(v20) >> (uint(int32(3)) % 32))
		case 1:
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-3)))))
			v37 = v27
		case 2:
			v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(-5)))))
			v37 = v30
		case 3:
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-9))))
			v37 = v33
		case 4:
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-17))))
			v37 = v36
		default:
			v37 = int32(0)
		}
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch int32(base.Ui32(v38)>>(uint(int32(4))%32))&int32(15) + int32(-9) {
		case 0:
			v65 = F_objectGetVal(m, l0)
			mBase = m.M
			v66 = F_quicklistReplaceAtIndex(m, v65, l1, v17, v37)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				v68 = v66
				F_decrRefCount(m, v13)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					m.G0 = v10 + int32(16)
					return v68
				}
			}
		default:
			F__serverPanic_1(m, int32(_a_F_listTypeReplaceAtIndex_0), int32(384), int32(_a_F_listTypeReplaceAtIndex_1), int32(0))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 2:
			v45 = F_objectGetVal(m, l0)
			mBase = m.M
			v46 = F_lpSeek(m, v45, l1)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v46
				if v46 != 0 {
					v50 = F_objectGetVal(m, l0)
					mBase = m.M
					v53 = F_lpReplace(m, v50, v10+int32(12), v17, v37)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						F_objectSetVal(m, l0, v53)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v68 = int32(1)
							F_decrRefCount(m, v13)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(16)
								return v68
							}
						}
					}
				} else {
					v68 = int32(0)
					F_decrRefCount(m, v13)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(16)
						return v68
					}
				}
			}
		}
	}
}
func F_listTypeTryConversion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v5 = int32(0)
	F_listTypeTryConversionRaw(m, l0, l1, v5, v5, v5, l2, l3)
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_list_add(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	v7 = m.G9
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v9 = m.T0[v8].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(8))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v14 != 0 {
			v17 = v14
			for {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				if v20 != 0 {
					v17 = v20
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v9
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v9
		}
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v26 + int32(1)
		return
	}
}
func F_list_create(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = m.G22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = m.T0[v4].(func(*base.Module, int32) int32)(m, int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = int64(0)
		return v5
	}
}
func F_list_destroy(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = m.G11
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	m.T0[v19].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	v8 = v4
	goto L3
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v11 = m.G11
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	m.T0[v12].(func(*base.Module, int32))(m, v8)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	if v10 != 0 {
		v8 = v10
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
}
