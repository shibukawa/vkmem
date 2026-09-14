package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_addListQuicklistRangeReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v71 int32
	_ = v71
	v8 = m.G0
	v10 = v8 - int32(32)
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
	F__serverAssert(m, int32(_a2385), int32(_a2378), int32(666))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L3
	} else {
		goto L21
	}
L2:
	;
	m.G0 = v10 + int32(32)
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
	v22 = F_quicklistGetIteratorAtIdx(m, v18, base.B2i32(l4 != int32(0)), base.I64_extend_i32_s(l2))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if l3 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_quicklistReleaseIterator(m, v22)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L20
	}
L9:
	;
	v29 = l3
	goto L10
L10:
	;
	v33 = F_quicklistNext(m, v22, v10)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	if v33 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v37 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v47 = v29 + int32(-1)
	if v47 != 0 {
		v29 = v47
		goto L10
	} else {
		goto L19
	}
L15:
	;
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
	F_addWritePreparedReplyBulkLongLong(m, v12, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	F_addWritePreparedReplyBulkCBuffer(m, v12, v37, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	goto L14
L19:
	;
	goto L11
L20:
	;
	goto L2
L21:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_listAddNodeHead(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v20 int32
	_ = v20
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
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(0)
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v6
				v22 = v20
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
				v15 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v15
				v22 = v15
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v6
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v13 + int32(1)
			return l0
		} else {
			return int32(0)
		}
	}
}
func F_listCreate(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	v3 = F_valkey_malloc(m, int32(24))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
		} else {
			v9 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v3))) = v9
			*(*int64)(unsafe.Add(mBase, uint32(v3+int32(16)))) = v9
			*(*int64)(unsafe.Add(mBase, uint32(v3+int32(8)))) = v9
		}
		return v3
	}
}
func F_listDelNode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	F_listUnlinkNode(m, l0, l1)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v5 == int32(0) {
			F_valkey_free(m, l1)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			m.T0[v5].(func(*base.Module, int32))(m, v8)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_valkey_free(m, l1)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_listDup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
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
	var v43 int32
	_ = v43
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	v7 = F_valkey_malloc(m, int32(24))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v13 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v13
	v16 = v7 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v13
	v20 = v7 + int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v13
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v29 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	return int32(0)
L3:
	;
	if v7 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	return int32(0)
L5:
	;
	v33 = v29
	goto L9
L6:
	;
	return v7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(0)
	F_valkey_free(m, v7)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L40
	}
L8:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v86 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v38 == int32(0) {
		v63 = v36
		goto L11
	} else {
		goto L12
	}
L10:
	;
	return v7
L11:
	;
	v65 = F_valkey_malloc(m, int32(12))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L23
	}
L12:
	;
	v41 = m.T0[v38].(func(*base.Module, int32) int32)(m, v36)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	if v41 != 0 {
		v63 = v41
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	if v43 == int32(0) {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v47 = v46
	v49 = v43
	goto L16
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v53 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_valkey_free(m, v47)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L2
	} else {
		goto L21
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	m.T0[v53].(func(*base.Module, int32))(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v62 = v49 + int32(-1)
	if v62 != 0 {
		v47 = v52
		v49 = v62
		goto L16
	} else {
		goto L22
	}
L22:
	;
	goto L7
L23:
	;
	if v65 == int32(0) {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v63
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	if v70 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v70 + int32(1)
	if v37 != 0 {
		v33 = v37
		goto L9
	} else {
		goto L28
	}
L26:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v65
	v79 = v75
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v65
	v72 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v72
	v79 = v72
	goto L25
L28:
	;
	goto L10
L29:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	if v91 == int32(0) {
		goto L7
	} else {
		goto L32
	}
L30:
	;
	m.T0[v86].(func(*base.Module, int32))(m, v63)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v95 = v94
	v97 = v91
	goto L33
L33:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v101 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L7
L35:
	;
	F_valkey_free(m, v95)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L38
	}
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	m.T0[v101].(func(*base.Module, int32))(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v110 = v97 + int32(-1)
	if v110 != 0 {
		v95 = v100
		v97 = v110
		goto L33
	} else {
		goto L39
	}
L39:
	;
	goto L34
L40:
	;
	return int32(0)
}
func F_listIndex(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	if l1 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v28
L2:
	;
	v21 = l0 + int32(4)
	v22 = l1 ^ int32(-1)
	goto L10
L3:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 == int32(0) {
		v28 = v5
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v5 == int32(0) {
		v28 = v5
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v10 = v5
	v11 = l1
	goto L6
L6:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v14 = v11 + int32(-1)
	if v14 == int32(0) {
		v28 = v12
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v12 != 0 {
		v10 = v12
		v11 = v14
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v28 = v12
	goto L1
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v22 == int32(0) {
		v28 = v23
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v28 = v23
	goto L1
L12:
	;
	if v23 != 0 {
		v21 = v23
		v22 = v22 + int32(-1)
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
}
func F_listLinkNodeHead(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v5 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
		v14 = v12
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
		v7 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v7
		v14 = v7
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v5 + int32(1)
	return
}
func F_listRelease(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v43 int32
	_ = v43
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L13
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v8 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	goto L1
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = v8
	v14 = v11
	goto L5
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v18 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	F_valkey_free(m, v14)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L11
	}
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	m.T0[v18].(func(*base.Module, int32))(m, v21)
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
	v27 = v13 + int32(-1)
	if v27 != 0 {
		v13 = v27
		v14 = v17
		goto L5
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	return
}
func F_listRewind(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3
	return
}
func F_listRewindTail(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3
	return
}
func F_listRotateHeadToTail(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v4) < base.Ui32(int32(2)) {
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v8
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	}
	return
}
func F_listSearchKey(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v37
L2:
	;
	v37 = int32(0)
	goto L1
L3:
	;
	v12 = v7
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v18 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L2
L6:
	;
	if v17 != 0 {
		v12 = v17
		goto L4
	} else {
		goto L13
	}
L7:
	;
	if l1 == v16 {
		v37 = v12
		goto L1
	} else {
		goto L12
	}
L8:
	;
	v21 = m.T0[v18].(func(*base.Module, int32, int32) int32)(m, v16, l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if v21 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v37 = v12
	goto L1
L12:
	;
	goto L6
L13:
	;
	goto L5
}
func F_listTypeDelRange(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v4)>>(uint(int32(4))%32))&int32(15) + int32(-9) {
	case 0:
		v23 = F_objectGetVal(m, l0)
		mBase = m.M
		v24 = F_quicklistDelRange(m, v23, l1, l2)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			return
		}
	default:
		F__serverPanic_1(m, int32(_a2378), int32(454), int32(_a2379), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 2:
		v11 = F_objectGetVal(m, l0)
		mBase = m.M
		v12 = F_lpDeleteRange(m, v11, l1, l2)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_objectSetVal(m, l0, v12)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_listTypeEqual(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
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
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch int32(base.Ui32(v6)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
		switch v21 + int32(-9) {
		case 0:
			v24 = F_objectGetVal(m, l1)
			mBase = m.M
			v26 = l0 + int32(8)
			v28 = F_objectGetVal(m, l1)
			mBase = m.M
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-1)))))
			switch v31 & int32(7) {
			case 0:
				v36 = F_quicklistCompare(m, v26, v24, int32(base.Ui32(v31)>>(uint(int32(3))%32)))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					return v36
				}
			case 1:
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))))
				v42 = F_quicklistCompare(m, v26, v24, v41)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					return v42
				}
			case 2:
				v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))))
				v48 = F_quicklistCompare(m, v26, v24, v47)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					return v48
				}
			case 3:
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9))))
				v54 = F_quicklistCompare(m, v26, v24, v53)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					return v54
				}
			case 4:
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-17))))
				v60 = v59
				v61 = F_quicklistCompare(m, v26, v24, v60)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					return v61
				}
			default:
				v60 = int32(0)
				v61 = F_quicklistCompare(m, v26, v24, v60)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					return v61
				}
			}
		default:
			F__serverPanic_1(m, int32(_a2378), int32(399), int32(_a2379), int32(0))
			mBase = m.M
			v108 = m.ExcPending
			if v108 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 2:
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v65 = F_objectGetVal(m, l1)
			mBase = m.M
			v67 = F_objectGetVal(m, l1)
			mBase = m.M
			v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+int32(-1)))))
			switch v70 & int32(7) {
			case 0:
				v75 = F_lpCompare(m, v64, v65, int32(base.Ui32(v70)>>(uint(int32(3))%32)))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					return v75
				}
			case 1:
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+int32(-3)))))
				v81 = F_lpCompare(m, v64, v65, v80)
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					return v81
				}
			case 2:
				v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+int32(-5)))))
				v87 = F_lpCompare(m, v64, v65, v86)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					return v87
				}
			case 3:
				v92 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(-9))))
				v93 = F_lpCompare(m, v64, v65, v92)
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					return v93
				}
			case 4:
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(-17))))
				v99 = v98
				v100 = F_lpCompare(m, v64, v65, v99)
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return int32(0)
				} else {
					return v100
				}
			default:
				v99 = int32(0)
				v100 = F_lpCompare(m, v64, v65, v99)
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return int32(0)
				} else {
					return v100
				}
			}
		}
	default:
		F__serverAssertWithInfo(m, int32(0), l1, int32(_a2381), int32(_a2378), int32(393))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_listTypeInitIterator(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v3 = l2
	v7 = F_valkey_malloc(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+5)) = uint8(v3)
		v17 = int32(base.Ui32(v12) >> (uint(int32(4)) % 32))
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v17)
		switch v17 + int32(-9) {
		case 0:
			v33 = F_objectGetVal(m, l0)
			mBase = m.M
			v37 = F_quicklistGetIteratorAtIdx(m, v33, base.B2i32(v3 == int32(0)), base.I64_extend_i32_s(l1))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v37
				return v7
			}
		default:
			F__serverPanic_1(m, int32(_a2378), int32(237), int32(_a2379), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 2:
			v21 = F_objectGetVal(m, l0)
			mBase = m.M
			v22 = F_lpSeek(m, v21, l1)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v22
				return v7
			}
		}
	}
}
func F_listTypeNext(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if int32(base.Ui32(v6)>>(uint(int32(4))%32))&int32(15) != v11 {
		F__serverAssert(m, int32(_a2380), int32(_a2378), int32(271))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = l0
		switch v11 + int32(-9) {
		case 0:
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v45 = F_quicklistNext(m, v42, l1+int32(8))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				return v45
			}
		default:
			F__serverPanic_1(m, int32(_a2378), int32(284), int32(_a2379), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 2:
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v16
			if v16 != 0 {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
				v21 = F_objectGetVal(m, v5)
				mBase = m.M
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v20 != int32(1) {
					v29 = F_lpPrev(m, v21, v22)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = v29
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v31
						return int32(1)
					}
				} else {
					v25 = F_lpNext(m, v21, v22)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v31 = v25
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v31
						return int32(1)
					}
				}
			} else {
				return int32(0)
			}
		}
	}
}
func F_listTypePush(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int64
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v250 int32
	_ = v250
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v14)>>(uint(int32(4))%32))&int32(15) + int32(-9) {
	case 0:
		goto L4
	default:
		goto L1
	case 2:
		goto L3
	}
L1:
	;
	F__serverPanic_1(m, int32(_a2378), int32(177), int32(_a2379), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L35
	} else {
		goto L76
	}
L2:
	;
	m.G0 = v12 + int32(32)
	return
L3:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v166 = F_objectGetVal(m, l0)
	mBase = m.M
	v167 = F_objectGetVal(m, l1)
	mBase = m.M
	if v165&int32(240) != int32(16) {
		goto L49
	} else {
		goto L50
	}
L4:
	;
	v21 = int32(0)
	if l2 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = int32(-1)
	goto L7
L6:
	;
	v24 = v21
	goto L7
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v25&int32(240) != int32(16) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v132 = F_objectGetVal(m, l0)
	mBase = m.M
	v133 = F_objectGetVal(m, l1)
	mBase = m.M
	v134 = F_objectGetVal(m, l1)
	mBase = m.M
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+int32(-1)))))
	switch v137 & int32(7) {
	case 0:
		goto L42
	case 1:
		goto L41
	case 2:
		goto L40
	case 3:
		goto L39
	case 4:
		goto L38
	default:
		v162 = v21
		goto L37
	}
L9:
	;
	v31 = F_objectGetVal(m, l1)
	mBase = m.M
	v32 = base.I64_extend_i32_s(v31)
	if v32 <= int64(-1) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v74 = F_objectGetVal(m, l0)
	mBase = m.M
	if v12&int32(3) == int32(0) {
		v96 = v12
		goto L21
	} else {
		goto L22
	}
L11:
	;
	goto L10
L13:
	;
	v54 = F_ull2string(m, v50, v51, v52)
	mBase = m.M
	if v54 == int32(0) {
		goto L11
	} else {
		goto L17
	}
L14:
	;
	goto L16
L15:
	;
	v50 = v12
	v51 = int32(32)
	v52 = v32
	goto L13
L16:
	;
	v41 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v41)
	v50 = v12 + int32(1)
	v51 = int32(31)
	v52 = int64(0) - v32
	goto L13
L17:
	;
	goto L10
L19:
	;
	F_quicklistPush(m, v74, v12, v129, v24)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L35
	} else {
		goto L36
	}
L20:
	;
	v129 = v121 - v12
	goto L19
L21:
	;
	v100 = v96
	goto L29
L22:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v82 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v85 = v12
	goto L25
L24:
	;
	v129 = v12 - v12
	goto L19
L25:
	;
	v89 = v85 + int32(1)
	if v89&int32(3) == int32(0) {
		v96 = v89
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v94 != 0 {
		v85 = v89
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v121 = v89
	goto L20
L29:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v109 = int32(-2139062144)
	if (int32(16843008)-v106|v106)&v109 == v109 {
		v100 = v100 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v115 = v100
	goto L32
L31:
	;
	goto L30
L32:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if v119 != 0 {
		v115 = v115 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v121 = v115
	goto L20
L34:
	;
	goto L33
L35:
	;
	return
L36:
	;
	goto L2
L37:
	;
	F_quicklistPush(m, v132, v133, v162, v24)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L35
	} else {
		goto L47
	}
L38:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v134+int32(-17))))
	v162 = v161
	goto L37
L39:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v134+int32(-9))))
	F_quicklistPush(m, v132, v133, v156, v24)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L35
	} else {
		goto L46
	}
L40:
	;
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134+int32(-5)))))
	F_quicklistPush(m, v132, v133, v151, v24)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L35
	} else {
		goto L45
	}
L41:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+int32(-3)))))
	F_quicklistPush(m, v132, v133, v146, v24)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L35
	} else {
		goto L44
	}
L42:
	;
	F_quicklistPush(m, v132, v133, int32(base.Ui32(v137)>>(uint(int32(3))%32)), v24)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L35
	} else {
		goto L43
	}
L43:
	;
	goto L2
L44:
	;
	goto L2
L45:
	;
	goto L2
L46:
	;
	goto L2
L47:
	;
	goto L2
L48:
	;
	F_objectSetVal(m, l0, v227)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L35
	} else {
		goto L75
	}
L49:
	;
	v177 = F_objectGetVal(m, l1)
	mBase = m.M
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+int32(-1)))))
	v182 = v180 & int32(7)
	if l2 != 0 {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v172 = base.I64_extend_i32_s(v167)
	if l2 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v175 = F_lpAppendInteger(m, v166, v172)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L35
	} else {
		goto L54
	}
L52:
	;
	v173 = F_lpPrependInteger(m, v166, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L35
	} else {
		goto L53
	}
L53:
	;
	v227 = v173
	goto L48
L54:
	;
	v227 = v175
	goto L48
L55:
	;
	switch v182 {
	case 0:
		goto L73
	case 1:
		goto L72
	case 2:
		goto L71
	case 3:
		goto L70
	case 4:
		goto L69
	default:
		v224 = int32(0)
		goto L68
	}
L56:
	;
	switch v182 {
	case 0:
		goto L62
	case 1:
		goto L61
	case 2:
		goto L60
	case 3:
		goto L59
	case 4:
		goto L58
	default:
		v206 = int32(0)
		goto L57
	}
L57:
	;
	v207 = F_lpPrepend(m, v166, v167, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L35
	} else {
		goto L67
	}
L58:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v177+int32(-17))))
	v206 = v205
	goto L57
L59:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v177+int32(-9))))
	v201 = F_lpPrepend(m, v166, v167, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L35
	} else {
		goto L66
	}
L60:
	;
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177+int32(-5)))))
	v196 = F_lpPrepend(m, v166, v167, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L35
	} else {
		goto L65
	}
L61:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+int32(-3)))))
	v191 = F_lpPrepend(m, v166, v167, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L35
	} else {
		goto L64
	}
L62:
	;
	v186 = F_lpPrepend(m, v166, v167, int32(base.Ui32(v180)>>(uint(int32(3))%32)))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L35
	} else {
		goto L63
	}
L63:
	;
	v227 = v186
	goto L48
L64:
	;
	v227 = v191
	goto L48
L65:
	;
	v227 = v196
	goto L48
L66:
	;
	v227 = v201
	goto L48
L67:
	;
	v227 = v207
	goto L48
L68:
	;
	v225 = F_lpAppend(m, v166, v167, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L35
	} else {
		goto L74
	}
L69:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v177+int32(-17))))
	v224 = v223
	goto L68
L70:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v177+int32(-9))))
	v224 = v220
	goto L68
L71:
	;
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177+int32(-5)))))
	v224 = v217
	goto L68
L72:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+int32(-3)))))
	v224 = v214
	goto L68
L73:
	;
	v224 = int32(base.Ui32(v180) >> (uint(int32(3)) % 32))
	goto L68
L74:
	;
	v227 = v225
	goto L48
L75:
	;
	goto L2
L76:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_listTypeReplace(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v10 = F_getDecodedObject(m, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = F_objectGetVal(m, v10)
		mBase = m.M
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-1)))))
		switch v15 & int32(7) {
		case 0:
			v32 = int32(base.Ui32(v15) >> (uint(int32(3)) % 32))
		case 1:
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
			v32 = v22
		case 2:
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-5)))))
			v32 = v25
		case 3:
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
			v32 = v28
		case 4:
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-17))))
			v32 = v31
		default:
			v32 = int32(0)
		}
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
		switch v34 + int32(-9) {
		case 0:
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
			F_quicklistReplaceEntry(m, v53, l0+int32(8), v12, v32)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				F_decrRefCount(m, v10)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					return
				}
			}
		default:
			F__serverPanic_1(m, int32(_a2378), int32(358), int32(_a2379), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 2:
			v37 = F_objectGetVal(m, v8)
			mBase = m.M
			v40 = F_lpReplace(m, v37, l0+int32(4), v12, v32)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_objectSetVal(m, v8, v40)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					F_decrRefCount(m, v10)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_listTypeTryConversionAppend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	var v9 int32
	_ = v9
	F_listTypeTryConversionRaw(m, l0, int32(1), l1, l2, l3, l4, l5)
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_listTypeTryConversionRaw(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v17)>>(uint(int32(4))%32))&int32(15) + int32(-9) {
	case 0:
		if l1 == int32(1) {
			m.G0 = v15 + int32(16)
			return
		} else {
			if v17&int32(240) != int32(144) {
				F__serverAssert(m, int32(_a2377), int32(_a2378), int32(83))
				mBase = m.M
				v277 = m.ExcPending
				if v277 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v30 = F_objectGetVal(m, l0)
				mBase = m.M
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				if v31 != int32(1) {
					m.G0 = v15 + int32(16)
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
					if v35&int32(786432) != int32(524288) {
						m.G0 = v15 + int32(16)
						return
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, _consts[542]))
						v43 = v15 + int32(12)
						v45 = v15 + int32(8)
						v46 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v43))) = v46
						*(*int32)(unsafe.Add(mBase, uint32(v45))) = v46
						if v41 < int32(0) {
							v57 = int32(-5)
							if base.Ui32(v57) < base.Ui32(v41) {
								v60 = v41
							} else {
								v60 = v57
							}
							v67 = *(*int32)(unsafe.Add(mBase, uint32((v60^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[524])))
							*(*int32)(unsafe.Add(mBase, uint32(v43))) = v67
						} else {
							v52 = int32(1)
							if base.Ui32(v52) < base.Ui32(v41) {
								v55 = v41
							} else {
								v55 = v52
							}
							*(*int32)(unsafe.Add(mBase, uint32(v45))) = v55
						}
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
						if l1 != int32(2) {
							v79 = v69
						} else {
							v72 = int32(1)
							v73 = int32(base.Ui32(v69) >> (uint(v72) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v73
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(base.Ui32(v75) >> (uint(v72) % 32))
							v79 = v73
						}
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
						if base.Ui32(v79) < base.Ui32(v81) {
							m.G0 = v15 + int32(16)
							return
						} else {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
							if base.Ui32(v84) < base.Ui32(v83) {
								m.G0 = v15 + int32(16)
								return
							} else {
								if l5 == int32(0) {
									v91 = v80
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
									F_objectSetVal(m, l0, v92)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
										*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = int32(0)
										F_quicklistRelease(m, v30)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = v100&int32(-241) | int32(176)
											m.G0 = v15 + int32(16)
											return
										}
									}
								} else {
									m.T0[l5].(func(*base.Module, int32))(m, l6)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
										v91 = v90
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
										F_objectSetVal(m, l0, v92)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return
										} else {
											v95 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
											*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = int32(0)
											F_quicklistRelease(m, v30)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												*(*int32)(unsafe.Add(mBase, uint32(l0))) = v100&int32(-241) | int32(176)
												m.G0 = v15 + int32(16)
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
	default:
		F__serverPanic_1(m, int32(_a2378), int32(135), int32(_a2379), int32(0))
		mBase = m.M
		v290 = m.ExcPending
		if v290 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 2:
		if l1 == int32(2) {
			m.G0 = v15 + int32(16)
			return
		} else {
			if v17&int32(240) != int32(176) {
				F__serverAssert(m, int32(_a2350), int32(_a2378), int32(43))
				mBase = m.M
				v283 = m.ExcPending
				if v283 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if l2 != 0 {
					v114 = int32(0)
					if l4 < l3 {
						v178 = v114
					} else {
						v117 = l3
						v125 = v114
						for {
							v131 = *(*int32)(unsafe.Add(mBase, uint32(l2+v117<<(uint(int32(2))%32))))
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
							switch int32(base.Ui32(v132)>>(uint(int32(4))%32)) & int32(15) {
							case 0, 8:
								v138 = F_objectGetVal(m, v131)
								mBase = m.M
								v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138+int32(-1)))))
								switch v141 & int32(7) {
								case 0:
									v158 = int32(base.Ui32(v141) >> (uint(int32(3)) % 32))
								case 1:
									v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138+int32(-3)))))
									v158 = v148
								case 2:
									v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138+int32(-5)))))
									v158 = v151
								case 3:
									v154 = *(*int32)(unsafe.Add(mBase, uint32(v138+int32(-9))))
									v158 = v154
								case 4:
									v157 = *(*int32)(unsafe.Add(mBase, uint32(v138+int32(-17))))
									v158 = v157
								default:
									v158 = int32(0)
								}
								v161 = v158 + v125
							default:
								v161 = v125
							}
							if base.B2i32(v117 == l4) == int32(0) {
								v117 = v117 + int32(1)
								v125 = v161
								continue
							} else {
								break
							}
							break
						}
						v178 = v161
					}
					v185 = l4 - l3 + int32(1)
					v193 = v178
				} else {
					v112 = int32(0)
					v185 = v112
					v193 = v112
				}
				v197 = *(*int32)(unsafe.Add(mBase, _consts[542]))
				v198 = F_objectGetVal(m, l0)
				mBase = m.M
				v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
				v200 = v199 + v193
				v201 = F_objectGetVal(m, l0)
				mBase = m.M
				v202 = F_lpLength(m, v201)
				mBase = m.M
				v203 = m.ExcPending
				if v203 != 0 {
					return
				} else {
					if int32(-1) < v197 {
						v221 = int32(1)
						if base.Ui32(v221) < base.Ui32(v197) {
							v224 = v197
						} else {
							v224 = v221
						}
						v227 = base.B2i32(base.Ui32(int32(8192)) < base.Ui32(v200)) | base.B2i32(base.Ui32(v224) < base.Ui32(v202+v185))
					} else {
						v207 = int32(-5)
						if base.Ui32(v207) < base.Ui32(v197) {
							v210 = v197
						} else {
							v210 = v207
						}
						v217 = *(*int32)(unsafe.Add(mBase, uint32((v210^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[524])))
						v227 = base.B2i32(base.Ui32(v217) < base.Ui32(v200))
					}
					if v227 == int32(0) {
						m.G0 = v15 + int32(16)
						return
					} else {
						if l5 == int32(0) {
							v234 = int32(_a20)
							v235 = *(*int32)(unsafe.Add(mBase, _consts[542]))
							v237 = *(*int32)(unsafe.Add(mBase, _consts[543]))
							v238 = F_quicklistNew(m, v235, v237)
							mBase = m.M
							v239 = m.ExcPending
							if v239 != 0 {
								return
							} else {
								v240 = F_objectGetVal(m, l0)
								mBase = m.M
								v241 = F_lpLength(m, v240)
								mBase = m.M
								v242 = m.ExcPending
								if v242 != 0 {
									return
								} else {
									v243 = F_objectGetVal(m, l0)
									mBase = m.M
									if v241 == int32(0) {
										F_lpFree(m, v243)
										mBase = m.M
										v249 = m.ExcPending
										if v249 != 0 {
											return
										} else {
											F_objectSetVal(m, l0, v238)
											mBase = m.M
											v251 = m.ExcPending
											if v251 != 0 {
												return
											} else {
												v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												*(*int32)(unsafe.Add(mBase, uint32(l0))) = v252&int32(-241) | int32(144)
												m.G0 = v15 + int32(16)
												return
											}
										}
									} else {
										F_quicklistAppendListpack(m, v238, v243)
										mBase = m.M
										v247 = m.ExcPending
										if v247 != 0 {
											return
										} else {
											F_objectSetVal(m, l0, v238)
											mBase = m.M
											v251 = m.ExcPending
											if v251 != 0 {
												return
											} else {
												v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												*(*int32)(unsafe.Add(mBase, uint32(l0))) = v252&int32(-241) | int32(144)
												m.G0 = v15 + int32(16)
												return
											}
										}
									}
								}
							}
						} else {
							m.T0[l5].(func(*base.Module, int32))(m, l6)
							mBase = m.M
							v233 = m.ExcPending
							if v233 != 0 {
								return
							} else {
								v234 = int32(_a20)
								v235 = *(*int32)(unsafe.Add(mBase, _consts[542]))
								v237 = *(*int32)(unsafe.Add(mBase, _consts[543]))
								v238 = F_quicklistNew(m, v235, v237)
								mBase = m.M
								v239 = m.ExcPending
								if v239 != 0 {
									return
								} else {
									v240 = F_objectGetVal(m, l0)
									mBase = m.M
									v241 = F_lpLength(m, v240)
									mBase = m.M
									v242 = m.ExcPending
									if v242 != 0 {
										return
									} else {
										v243 = F_objectGetVal(m, l0)
										mBase = m.M
										if v241 == int32(0) {
											F_lpFree(m, v243)
											mBase = m.M
											v249 = m.ExcPending
											if v249 != 0 {
												return
											} else {
												F_objectSetVal(m, l0, v238)
												mBase = m.M
												v251 = m.ExcPending
												if v251 != 0 {
													return
												} else {
													v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													*(*int32)(unsafe.Add(mBase, uint32(l0))) = v252&int32(-241) | int32(144)
													m.G0 = v15 + int32(16)
													return
												}
											}
										} else {
											F_quicklistAppendListpack(m, v238, v243)
											mBase = m.M
											v247 = m.ExcPending
											if v247 != 0 {
												return
											} else {
												F_objectSetVal(m, l0, v238)
												mBase = m.M
												v251 = m.ExcPending
												if v251 != 0 {
													return
												} else {
													v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													*(*int32)(unsafe.Add(mBase, uint32(l0))) = v252&int32(-241) | int32(144)
													m.G0 = v15 + int32(16)
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
	}
}
func F_listUnlinkNode(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 == int32(0) {
		F__serverAssert(m, int32(_a105), int32(_a106), int32(192))
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
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v9 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v16 != l1 {
				F__serverAssert(m, int32(_a107), int32(_a106), int32(198))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18
				v20 = v18
				if v20 == int32(0) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v26 != l1 {
						F__serverAssert(m, int32(_a108), int32(_a106), int32(206))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
						*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6 + int32(-1)
						return
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					if v23 != l1 {
						F__serverAssert(m, int32(_a109), int32(_a106), int32(203))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v9
						*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6 + int32(-1)
						return
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			if v12 != l1 {
				F__serverAssert(m, int32(_a110), int32(_a106), int32(195))
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
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v14
				v20 = v14
				if v20 == int32(0) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v26 != l1 {
						F__serverAssert(m, int32(_a108), int32(_a106), int32(206))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
						*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6 + int32(-1)
						return
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					if v23 != l1 {
						F__serverAssert(m, int32(_a109), int32(_a106), int32(203))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v9
						*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6 + int32(-1)
						return
					}
				}
			}
		}
	}
}
func F_list_iter_next(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v8
		return v7
	} else {
		return int32(0)
	}
}
func F_rewriteListObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	v10 = m.G0
	v12 = v10 - int32(64)
	m.G0 = v12
	v15 = F_listTypeLength(m, l2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = F_listTypeInitIterator(m, l2, int32(0), int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v12 + int32(64)
	return v115
L4:
	;
	F_listTypeReleaseIterator(m, v21)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L36
	}
L5:
	;
	F_listTypeReleaseIterator(m, v21)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L35
	}
L6:
	;
	v25 = F_listTypeNext(m, v21, v12+int32(24))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v25 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v38 = base.I64_extend_i32_u(v15)
	v39 = int64(0)
	goto L9
L9:
	;
	if v39 != int64(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L5
L11:
	;
	v70 = F_listTypeGetValue(m, v12+int32(24), v12+int32(20), v12+int32(8))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L24
	}
L12:
	;
	v43 = int64(64)
	if v38 < v43 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v38
	goto L15
L14:
	;
	v46 = v43
	goto L15
L15:
	;
	v50 = F_rioWriteBulkCount(m, l0, int32(42), base.I32_wrap_i64(v46)+int32(2))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v50 == int32(0) {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v56 = F_rioWriteBulkString(m, l0, int32(_a160), int32(5))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v56 == int32(0) {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v60 = F_rioWriteBulkObject(m, l0, l1)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v60 == int32(0) {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	goto L11
L22:
	;
	v86 = v39 + int64(1)
	if v86 == int64(64) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	v80 = F_rioWriteBulkLongLong(m, l0, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L28
	}
L24:
	;
	if v70 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v75 = F_rioWriteBulkString(m, l0, v70, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v75 == int32(0) {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	goto L22
L28:
	;
	if v80 == int32(0) {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	goto L22
L30:
	;
	v89 = int64(0)
	goto L32
L31:
	;
	v89 = v86
	goto L32
L32:
	;
	v94 = F_listTypeNext(m, v21, v12+int32(24))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v94 != 0 {
		v38 = v38 + int64(-1)
		v39 = v89
		goto L9
	} else {
		goto L34
	}
L34:
	;
	goto L10
L35:
	;
	v115 = int32(1)
	goto L3
L36:
	;
	v115 = int32(0)
	goto L3
}
