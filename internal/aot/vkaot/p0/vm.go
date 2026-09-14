package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_VM_ACLAddLogEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	if base.Ui32(int32(4)) < base.Ui32(l3) {
		v27 = int32(1)
		return v27
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_c_F_VM_ACLAddLogEntry[0])))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		v19 = F_objectGetVal(m, l2)
		mBase = m.M
		v20 = F_sdsdup(m, v19)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			F_addACLLogEntry(m, v9, v14, int32(3), int32(-1), v18, v20)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v27 = int32(0)
				return v27
			}
		}
	}
}
func F_VM_ACLAddLogEntryByUserName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	if base.Ui32(int32(4)) < base.Ui32(l3) {
		v26 = int32(1)
		return v26
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_c_F_VM_ACLAddLogEntryByUserName[0])))
		v17 = F_objectGetVal(m, l1)
		mBase = m.M
		v18 = F_objectGetVal(m, l2)
		mBase = m.M
		v19 = F_sdsdup(m, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_addACLLogEntry(m, v9, v14, int32(3), int32(-1), v17, v19)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v26 = int32(0)
				return v26
			}
		}
	}
}
func F_VM_ACLCheckKeyPrefixPermissions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v26 int64
	_ = v26
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v56 int64
	_ = v56
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v84 int64
	_ = v84
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v97 int64
	_ = v97
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v112 int64
	_ = v112
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v125 int64
	_ = v125
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v140 int64
	_ = v140
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v153 int64
	_ = v153
	var v161 int64
	_ = v161
	var v163 int64
	_ = v163
	var v168 int64
	_ = v168
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	v7 = int32(28)
	if l0 == int32(0) {
		v219 = v7
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v225
L2:
	;
	goto L38
L3:
	;
	if l3&int32(-241) != 0 {
		v219 = v7
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v12 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = base.I64_extend_i32_u(l3)
	v26 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[0]))
	if base.B2i32(v26&v15 == int64(0)) == v12 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v181 = int32(0)
	v183 = m.G0
	v185 = v183 - int32(16)
	m.G0 = v185
	if v14 == v181 {
		v213 = v181
		goto L30
	} else {
		goto L31
	}
L6:
	;
	v41 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[1]))
	if v41&v15 == int64(0) {
		v51 = v38
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v37 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[2]))
	v38 = v37
	goto L6
L8:
	;
	v38 = int64(0)
	goto L6
L9:
	;
	v56 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[3]))
	if v56&v15 == int64(0) {
		v66 = v51
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v49 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[4]))
	v51 = v49 | v38
	goto L9
L11:
	;
	v69 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[5]))
	if v69&v15 == int64(0) {
		v79 = v66
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v64 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[6]))
	v66 = v64 | v51
	goto L11
L13:
	;
	v84 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[7]))
	if v84&v15 == int64(0) {
		v94 = v79
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v77 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[8]))
	v79 = v77 | v66
	goto L13
L15:
	;
	v97 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[9]))
	if v97&v15 == int64(0) {
		v107 = v94
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v92 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[10]))
	v94 = v92 | v79
	goto L15
L17:
	;
	v112 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[11]))
	if v112&v15 == int64(0) {
		v122 = v107
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v105 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[12]))
	v107 = v105 | v94
	goto L17
L19:
	;
	v125 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[13]))
	if v125&v15 == int64(0) {
		v135 = v122
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v120 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[14]))
	v122 = v120 | v107
	goto L19
L21:
	;
	v140 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[15]))
	if v140&v15 == int64(0) {
		v150 = v135
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v133 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[16]))
	v135 = v133 | v122
	goto L21
L23:
	;
	v153 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[17]))
	if v153&v15 == int64(0) {
		v163 = v150
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v148 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[18]))
	v150 = v148 | v135
	goto L23
L25:
	;
	v168 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[19]))
	if v168&v15 == int64(0) {
		v178 = v163
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v161 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[20]))
	v163 = v161 | v150
	goto L25
L27:
	;
	goto L5
L28:
	;
	v176 = *(*int64)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[21]))
	v178 = v176 | v163
	goto L27
L29:
	;
	if v213 == int32(0) {
		v225 = v12
		goto L1
	} else {
		goto L37
	}
L30:
	;
	m.G0 = v185 + int32(16)
	goto L29
L31:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_listRewind(m, v190, v185+int32(8))
	mBase = m.M
	goto L32
L32:
	;
	v203 = F_listNext(m, v185+int32(8))
	mBase = m.M
	if v203 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v213 = v181
	goto L30
L34:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	v206 = F_ACLSelectorCheckKey(m, v205, l1, l2, base.I32_wrap_i64(v178), int32(1))
	mBase = m.M
	if v206 != 0 {
		goto L32
	} else {
		goto L36
	}
L35:
	;
	v213 = int32(3)
	goto L30
L36:
	;
	goto L33
L37:
	;
	v219 = int32(2)
	goto L2
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_ACLCheckKeyPrefixPermissions[22])) = v219
	v225 = int32(1)
	goto L1
}
func F_VM_AddACLCategory(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int64
	_ = v106
	var v113 int64
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v152 int32
	_ = v152
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v152
L2:
	;
	v152 = int32(1)
	goto L1
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v16&int32(255) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_AddACLCategory[0])) = int32(28)
	goto L2
L6:
	;
	goto L36
L7:
	;
	v25 = v16
	v26 = int32(0)
	goto L8
L8:
	;
	if base.Ui32((v25+int32(-48))&int32(255)) < base.Ui32(int32(10)) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v68 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_VM_AddACLCategory[1]))
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v74)+8))
	if base.B2i32(v75 == int64(0)) == v68 {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	v63 = v26 + int32(1)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v63))))
	if v65&int32(255) != 0 {
		v25 = v65
		v26 = v63
		goto L8
	} else {
		goto L18
	}
L11:
	;
	if base.Ui32((v25&int32(-33)+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v43 = v25 & int32(255)
	if v43 == int32(45) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	if v43 == int32(95) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_VM_AddACLCategory[2]))
	if int32(3) < v49 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = base.I32_extend8_s(v25)
	F__serverLog(m, int32(3), int32(_a_F_VM_AddACLCategory_0), v9)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	goto L6
L18:
	;
	goto L9
L19:
	;
	v120 = F_ACLAddCommandCategory(m, l1, int64(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L16
	} else {
		goto L33
	}
L20:
	;
	if v113 == int64(0) {
		goto L19
	} else {
		goto L30
	}
L21:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v82 = F_strcasecmp(m, l1, v81)
	mBase = m.M
	if v82 == int32(0) {
		v106 = v75
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v113 = int64(0)
	goto L20
L23:
	;
	v113 = v106
	goto L20
L24:
	;
	v86 = v68
	goto L25
L25:
	;
	v91 = v86 + int32(1)
	v94 = v74 + v91<<(uint(int32(4))%32)
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v94)+8))
	if base.B2i32(v95 == int64(0)) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v106 = v95
	goto L23
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v102 = F_strcasecmp(m, l1, v101)
	mBase = m.M
	if v102 != 0 {
		v86 = v91
		goto L25
	} else {
		goto L29
	}
L28:
	;
	v113 = int64(0)
	goto L20
L29:
	;
	goto L26
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_AddACLCategory[0])) = int32(10)
	goto L2
L32:
	;
	goto L35
L33:
	;
	if v120 == int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+76)) = v125 + int32(1)
	v152 = int32(0)
	goto L1
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_AddACLCategory[0])) = int32(48)
	goto L2
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_AddACLCategory[0])) = int32(28)
	goto L2
}
func F_VM_AuthenticateClientWithACLUser(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v7 = int32(0)
	v8 = m.G0
	v9 = int32(16)
	v10 = v8 - v9
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v7
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_VM_AuthenticateClientWithACLUser[0]))
	v18 = F_raxFind(m, v15, l1, l2, v10+int32(12))
	mBase = m.M
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	m.G0 = v10 + v9
	if v19 != 0 {
		v25 = F_authenticateClientWithUser(m, l0, v19, l3, l4, l5)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v25
		}
	} else {
		return int32(1)
	}
}
func F_VM_AuthenticateClientWithUser(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = F_authenticateClientWithUser(m, l0, v6, l2, l3, l4)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_VM_BlockClient(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) int32 {
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v6 = int32(0)
	v11 = F_moduleBlockClient(m, l0, l1, v6, l2, l3, l4, v6, v6, v6, v6)
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_VM_BlockClientGetPrivateData(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return v2
}
func F_VM_BlockClientOnKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int32, l7 int32) int32 {
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v9 = int32(0)
	v11 = F_moduleBlockClient(m, l0, l1, v9, l2, l3, l4, l5, l6, l7, v9)
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_VM_BlockClientOnKeysWithFlags(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v11 = F_moduleBlockClient(m, l0, l1, int32(0), l2, l3, l4, l5, l6, l7, l8)
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_VM_BlockClientSetPrivateData(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l1
	return
}
func F_VM_BlockedClientMeasureTimeEnd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if base.B2i32(v3 == int64(0)) == int32(0) {
		v10 = int32(0)
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_VM_BlockedClientMeasureTimeEnd[0]))
		v12 = m.T0[v11].(func(*base.Module) int64)(m)
		mBase = m.M
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v12 - v3 + v14
		return v10
	} else {
		return int32(1)
	}
}
func F_VM_CallReplyBigNumber(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_callReplyGetBigNumber(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_VM_CallReplyBool(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_callReplyGetBool(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_VM_CallReplyDouble(m *base.Module, l0 int32) float64 {
	var v2 float64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_callReplyGetDouble(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return float64(0)
	} else {
		return v2
	}
}
func F_VM_CallReplyInteger(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_callReplyGetLongLong(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_VM_CallReplyLength(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_callReplyGetLen(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_VM_CallReplyProto(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	return v5
}
func F_VM_CommandFilterArgDelete(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int64
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v157 int32
	_ = v157
	v8 = int32(1)
	if l1 < int32(0) {
		v157 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v157
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 <= l1 {
		v157 = v8
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+208)))
	if v14&int32(1) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64+l1<<(uint(int32(2))%32))))
	F_decrRefCount(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L15
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+192))
	if v19 != 0 {
		v64 = v18
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v64 = v17
	goto L4
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v13)+192)) = v18
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = F_valkey_malloc(m, v22<<(uint(int32(2))%32))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v25
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v30 < int32(1) {
		v64 = v25
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+192))
	v38 = int32(0)
	v41 = v34
	goto L11
L11:
	;
	v44 = v38 << (uint(int32(2)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41+v44)))
	F_incrRefCount(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v64 = v49
	goto L4
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+192))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52+v44)))
	*(*int32)(unsafe.Add(mBase, uint32(v49+v44))) = v54
	v57 = v38 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v57 < v58 {
		v38 = v57
		v41 = v52
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v75 = v73 + int32(-1)
	if v75 <= l1 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v75
	v157 = int32(0)
	goto L1
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v85 = (v73 + (l1 ^ int32(-1))) & int32(3)
	if v85 == int32(0) {
		v110 = l1
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if base.Ui32(v73-l1+int32(-2)) < base.Ui32(int32(3)) {
		goto L16
	} else {
		goto L23
	}
L19:
	;
	v90 = l1
	v92 = int32(0)
	goto L20
L20:
	;
	v96 = int32(2)
	v99 = int32(1)
	v100 = v90 + v99
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v80+v100<<(uint(v96)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v80+v90<<(uint(v96)%32)))) = v104
	v107 = v92 + v99
	if v107 != v85 {
		v90 = v100
		v92 = v107
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v110 = v100
	goto L18
L22:
	;
	goto L21
L23:
	;
	v119 = v110
	goto L24
L24:
	;
	v125 = int32(2)
	v127 = v80 + v119<<(uint(v125)%32)
	v128 = int32(4)
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v127+v128)))
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = v130
	v135 = v127 + int32(12)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	*(*int32)(unsafe.Add(mBase, uint32(v127+int32(8)))) = v136
	v139 = v119 + v128
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v80+v139<<(uint(v125)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v143
	if v139 != v75 {
		v119 = v139
		goto L24
	} else {
		goto L26
	}
L25:
	;
	goto L16
L26:
	;
	goto L25
}
func F_VM_CommandFilterArgInsert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int64
	_ = v142
	var v148 int64
	_ = v148
	var v151 int32
	_ = v151
	var v173 int32
	_ = v173
	v10 = int32(1)
	if l1 < int32(0) {
		v173 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v173
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v13 < l1 {
		v173 = v10
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= v13 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v81 <= l1 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v18 = int32(1)
	v19 = v13 + v18
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+208)))
	if v22&v18 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v81 = v13
	v82 = v17
	goto L4
L7:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v73 = F_valkey_realloc(m, v70, v19<<(uint(int32(2))%32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L17
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+192))
	if v28 != 0 {
		v81 = v13
		v82 = v27
		goto L4
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v27
	v33 = F_valkey_malloc(m, v19<<(uint(int32(2))%32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v33
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v38 < int32(1) {
		v81 = v38
		v82 = v33
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+192))
	v47 = int32(0)
	v51 = v42
	goto L13
L13:
	;
	v54 = v47 << (uint(int32(2)) % 32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51+v54)))
	F_incrRefCount(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+192))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62+v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v59+v54))) = v64
	v67 = v47 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v67 < v68 {
		v47 = v67
		v51 = v62
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v81 = v68
	v82 = v59
	goto L4
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v73
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = v76
	v82 = v73
	goto L4
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82+l1<<(uint(int32(2))%32)))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v81 + int32(1)
	v173 = int32(0)
	goto L1
L19:
	;
	v89 = (v81 - l1) & int32(3)
	if v89 == int32(0) {
		v117 = v81
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if base.Ui32(int32(-4)) < base.Ui32(l1-v81) {
		goto L18
	} else {
		goto L25
	}
L21:
	;
	v96 = v81
	v99 = int32(0)
	goto L22
L22:
	;
	v104 = v82 + v96<<(uint(int32(2))%32)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v104+int32(-4))))
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v107
	v110 = v96 + int32(-1)
	v112 = v99 + int32(1)
	if v112 != v89 {
		v96 = v110
		v99 = v112
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v117 = v110
	goto L20
L24:
	;
	goto L23
L25:
	;
	v129 = v117
	goto L26
L26:
	;
	v137 = v82 + v129<<(uint(int32(2))%32)
	v138 = int32(-4)
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v137+int32(-8))))
	*(*int64)(unsafe.Add(mBase, uint32(v137+v138))) = v142
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v137+int32(-16))))
	*(*int64)(unsafe.Add(mBase, uint32(v137+int32(-12)))) = v148
	v151 = v129 + v138
	if l1 < v151 {
		v129 = v151
		goto L26
	} else {
		goto L28
	}
L27:
	;
	goto L18
L28:
	;
	goto L27
}
func F_VM_CommandFilterArgReplace(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	v8 = int32(1)
	if l1 < int32(0) {
		v80 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v80
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 <= l1 {
		v80 = v8
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+208)))
	if v14&int32(1) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v68 = l1 << (uint(int32(2)) % 32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65+v68)))
	F_decrRefCount(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L15
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+192))
	if v19 != 0 {
		v65 = v18
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = v17
	goto L4
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v13)+192)) = v18
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = F_valkey_malloc(m, v22<<(uint(int32(2))%32))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v25
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v30 < int32(1) {
		v65 = v25
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+192))
	v39 = int32(0)
	v42 = v34
	goto L11
L11:
	;
	v44 = v39 << (uint(int32(2)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v44)))
	F_incrRefCount(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v65 = v49
	goto L4
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+192))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52+v44)))
	*(*int32)(unsafe.Add(mBase, uint32(v49+v44))) = v54
	v57 = v39 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v57 < v58 {
		v39 = v57
		v42 = v52
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v73+v68))) = l2
	v80 = int32(0)
	goto L1
}
func F_VM_CreateDict(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
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
	var v42 int32
	_ = v42
	v6 = F_valkey_malloc(m, int32(4))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_raxNew(m)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
			if l0 == int32(0) {
				return v6
			} else {
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
				if v15&int32(1) == int32(0) {
					return v6
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v20 == v21 {
						v24 = int32(8)
						if v24 < v20 {
							v27 = v20
						} else {
							v27 = v24
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v27 << (uint(int32(1)) % 32)
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v34 = F_valkey_realloc(m, v31, v27<<(uint(int32(4))%32))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v34
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v38 = v37
							v39 = v34
							v42 = v39 + v38<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v42))) = v6
							*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v38 + int32(1)
							return v6
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v38 = v20
						v39 = v23
						v42 = v39 + v38<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v42))) = v6
						*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v38 + int32(1)
						return v6
					}
				}
			}
		}
	}
}
func F_VM_CreateModuleUser(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v5 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_ACLCreateUnlinkedUser(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v9
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			F_sdsfree(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = F_sdsnew(m, l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
					return v5
				}
			}
		}
	}
}
func F_VM_CreateStringFromStreamID(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
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
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	v14 = F_createObjectFromStreamID(m, v8)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if l0 == int32(0) {
			m.G0 = v8 + int32(16)
			return v14
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
			if v20&int32(1) == int32(0) {
				m.G0 = v8 + int32(16)
				return v14
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
						*(*int32)(unsafe.Add(mBase, uint32(v47))) = v14
						v49 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v49
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v43 + v49
						m.G0 = v8 + int32(16)
						return v14
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v43 = v25
					v44 = v28
					v47 = v44 + v43<<(uint(int32(3))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v47))) = v14
					v49 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v49
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v43 + v49
					m.G0 = v8 + int32(16)
					return v14
				}
			}
		}
	}
}
func F_VM_CreateStringFromULongLong(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int64
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int64
	_ = v90
	var v92 int32
	_ = v92
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v16 = int32(1)
	if base.Ui64(l1) < base.Ui64(int64(10)) {
		v73 = v16
		v74 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v148 = F_createStringObject_1(m, v9, v147)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L45
	} else {
		goto L46
	}
L2:
	;
	v77 = v73 + v74
	if base.Ui32(int32(21)) <= base.Ui32(v77) {
		goto L33
	} else {
		goto L34
	}
L3:
	;
	v24 = v3
	v25 = l1
	goto L4
L4:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v25) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v73 = v16
	v74 = v65
	goto L2
L6:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v25) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v73 = int32(2)
	v74 = v24
	goto L2
L8:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v25) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v73 = int32(3)
	v74 = v24
	goto L2
L10:
	;
	v65 = v24 + int32(12)
	v69 = base.I64_div_u_s(v25, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v25) {
		v24 = v65
		v25 = v69
		goto L4
	} else {
		goto L32
	}
L11:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v25) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v25) {
		goto L24
	} else {
		goto L25
	}
L13:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v25) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v25) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v25) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v25) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v73 = int32(4)
	v74 = v24
	goto L2
L18:
	;
	v46 = int32(6)
	goto L20
L19:
	;
	v46 = int32(5)
	goto L20
L20:
	;
	v73 = v46
	v74 = v24
	goto L2
L21:
	;
	v51 = int32(8)
	goto L23
L22:
	;
	v51 = int32(7)
	goto L23
L23:
	;
	v73 = v51
	v74 = v24
	goto L2
L24:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v25) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v25) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v58 = int32(10)
	goto L28
L27:
	;
	v58 = int32(9)
	goto L28
L28:
	;
	v73 = v58
	v74 = v24
	goto L2
L29:
	;
	v63 = int32(12)
	goto L31
L30:
	;
	v63 = int32(11)
	goto L31
L31:
	;
	v73 = v63
	v74 = v24
	goto L2
L32:
	;
	goto L5
L33:
	;
	goto L44
L34:
	;
	v80 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9+v77))) = uint8(v80)
	v83 = v77 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(l1) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v119 = v9 + v116
	if base.Ui64(int64(9)) < base.Ui64(v117) {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	v90 = l1
	v92 = v83
	goto L38
L37:
	;
	v116 = v83
	v117 = l1
	goto L35
L38:
	;
	v96 = int64(100)
	v97 = base.I64_div_u_s(v90, v96)
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v90-v97*v96)<<(uint(int32(1))%32))+uint32(_c_F_VM_CreateStringFromULongLong[0]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v9+int32(-1)+v92))) = uint16(v106)
	v109 = v92 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v90) {
		v90 = v97
		v92 = v109
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v116 = v109
	v117 = v97
	goto L35
L40:
	;
	goto L39
L41:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v117)<<(uint(int32(1))%32))+uint32(_c_F_VM_CreateStringFromULongLong[0]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v119+int32(-1)))) = uint16(v133)
	v147 = v77
	goto L1
L42:
	;
	v124 = base.I32_wrap_i64(v117) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v124)
	v147 = v77
	goto L1
L43:
	;
	v147 = int32(0)
	goto L1
L44:
	;
	v137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v137)
	goto L43
L45:
	;
	return int32(0)
L46:
	;
	if l0 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	m.G0 = v9 + int32(32)
	return v148
L48:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v154&int32(1) == int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v159 == v160 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v181 = v178 + v177<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v148
	v183 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v181)+4)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v177 + v183
	goto L47
L51:
	;
	v163 = int32(8)
	if v163 < v159 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v177 = v159
	v178 = v162
	goto L50
L53:
	;
	v166 = v159
	goto L55
L54:
	;
	v166 = v163
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v166 << (uint(int32(1)) % 32)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v173 = F_valkey_realloc(m, v170, v166<<(uint(int32(4))%32))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L45
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v173
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v177 = v176
	v178 = v173
	goto L50
}
func F_VM_CreateStringPrintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_sdsempty(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
		v16 = F_sdscatvprintf(m, v10, l1, l2)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = F_createObject(m, int32(0), v16)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if l0 == int32(0) {
					m.G0 = v8 + int32(16)
					return v18
				} else {
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v22&int32(1) == int32(0) {
						m.G0 = v8 + int32(16)
						return v18
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v27 == v28 {
							v31 = int32(8)
							if v31 < v27 {
								v34 = v27
							} else {
								v34 = v31
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v34 << (uint(int32(1)) % 32)
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v41 = F_valkey_realloc(m, v38, v34<<(uint(int32(4))%32))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v41
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v45 = v44
								v46 = v41
								v49 = v46 + v45<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v49))) = v18
								v51 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v51
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v45 + v51
								m.G0 = v8 + int32(16)
								return v18
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v45 = v27
							v46 = v30
							v49 = v46 + v45<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v49))) = v18
							v51 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v51
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v45 + v51
							m.G0 = v8 + int32(16)
							return v18
						}
					}
				}
			}
		}
	}
}
func F_VM_CreateTimer(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int64
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v302 int64
	_ = v302
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int64
	_ = v326
	var v327 int64
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int64
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int64
	_ = v338
	var v343 int32
	_ = v343
	var v356 int32
	_ = v356
	var v358 int64
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int64
	_ = v367
	var v368 int32
	_ = v368
	var v370 int64
	_ = v370
	v9 = m.G0
	v11 = v9 - int32(320)
	m.G0 = v11
	v14 = F_valkey_malloc(m, int32(16))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v18
	v22 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v23 == v22 {
		v28 = v22
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v28
	v30 = F_ustime(m)
	mBase = m.M
	v41 = v30 + l1*int64(1000)
	goto L5
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+96))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v28 = v27
	goto L3
L5:
	;
	v42 = int64(56)
	v44 = int64(65280)
	v46 = int64(40)
	v49 = int64(16711680)
	v51 = int64(24)
	v53 = int64(4278190080)
	v55 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+312)) = v41<<(uint(v42)%64) | v41&v44<<(uint(v46)%64) | (v41&v49<<(uint(v51)%64) | v41&v53<<(uint(v55)%64)) | (int64(base.Ui64(v41)>>(uint(v55)%64))&v53 | int64(base.Ui64(v41)>>(uint(v51)%64))&v49 | (int64(base.Ui64(v41)>>(uint(v46)%64))&v44 | int64(base.Ui64(v41)>>(uint(v42)%64))))
	v80 = int32(0)
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_VM_CreateTimer[0]))
	v83 = v11 + int32(312)
	v84 = int32(8)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	goto L10
L6:
	;
	v281 = int32(0)
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_VM_CreateTimer[0]))
	v287 = F_raxInsert(m, v282, v11+int32(312), int32(8), v14, v281)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L46
	}
L7:
	;
	if v279 != 0 {
		v41 = v41 + int64(1)
		goto L5
	} else {
		goto L45
	}
L8:
	;
	if v238 != v84 {
		v279 = v80
		goto L35
	} else {
		goto L36
	}
L9:
	;
	v229 = int32(0)
	v236 = v95
	v238 = v229
	v242 = v229
	goto L8
L10:
	;
	if base.Ui32(v95) < base.Ui32(int32(8)) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v106 = v94
	v107 = v95
	v109 = int32(0)
	goto L13
L12:
	;
	v236 = v220
	v238 = v222
	v242 = base.B2i32(v225 != int32(0))
	goto L8
L13:
	;
	v115 = int32(base.Ui32(v107) >> (uint(int32(3)) % 32))
	v116 = int32(4)
	v117 = v106 + v116
	if v107&v116 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v220 = v211
	v222 = v195
	v225 = v200
	goto L12
L15:
	;
	v200 = int32(0)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v117+v115+(v200-v115)&int32(3)+v188<<(uint(int32(2))%32))))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	if base.Ui32(v211) < base.Ui32(int32(8)) {
		v220 = v211
		v222 = v195
		v225 = v200
		goto L12
	} else {
		goto L33
	}
L16:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v109))))
	v166 = int32(0)
	goto L27
L17:
	;
	v122 = int32(0)
	if base.Ui32(v84) <= base.Ui32(v109) {
		v155 = v109
		v158 = v122
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v158 == v115 {
		v188 = v122
		v195 = v155
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v132 = v109
	v135 = v122
	goto L20
L20:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+v135))))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v132))))
	if v138 != v140 {
		v155 = v132
		v158 = v135
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v155 = v143
	v158 = v145
	goto L18
L22:
	;
	v142 = int32(1)
	v143 = v132 + v142
	v145 = v135 + v142
	if base.Ui32(v115) <= base.Ui32(v145) {
		v155 = v143
		v158 = v145
		goto L18
	} else {
		goto L23
	}
L23:
	;
	if base.Ui32(v143) < base.Ui32(v84) {
		v132 = v143
		v135 = v145
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v220 = v107
	v222 = v155
	v225 = v158
	goto L12
L26:
	;
	if v166 != v115 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+v166))))
	if v179 == v163&int32(255) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v181 = int32(1)
	v183 = v166 + v181
	if v183 != v115 {
		v166 = v183
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v236 = v107
	v238 = v109
	v242 = v181
	goto L8
L31:
	;
	v188 = v166
	v195 = v109 + int32(1)
	goto L15
L32:
	;
	v220 = v107
	v222 = v109
	v225 = v115
	goto L12
L33:
	;
	if base.Ui32(v195) < base.Ui32(v84) {
		v106 = v210
		v107 = v211
		v109 = v195
		goto L13
	} else {
		goto L34
	}
L34:
	;
	goto L14
L35:
	;
	goto L7
L36:
	;
	v244 = int32(0)
	if v236&int32(1) == v244 {
		v279 = v244
		goto L35
	} else {
		goto L37
	}
L37:
	;
	if v242&base.B2i32(v236&int32(4) != int32(0)) != 0 {
		v279 = v244
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v279 = int32(1)
	goto L35
L45:
	;
	goto L6
L46:
	;
	v290 = *(*int64)(unsafe.Add(mBase, _c_F_VM_CreateTimer[1]))
	if v290 == int64(-1) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v11)+312))
	m.G0 = v11 + int32(320)
	return v370
L48:
	;
	v361 = int32(0)
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_VM_CreateTimer[2]))
	v367 = F_aeCreateTimeEvent(m, v363, l1, int32(566), v361, v361)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L65
	}
L49:
	;
	v294 = v11 + int32(8)
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_VM_CreateTimer[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v294)+4)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v294)+20)) = int32(128)
	v302 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v294)+12)) = v302
	*(*int64)(unsafe.Add(mBase, uint32(v294)+296)) = v302
	*(*int64)(unsafe.Add(mBase, uint32(v294)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v294)+8)) = v11 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v294)+156)) = v11 + int32(176)
	goto L50
L50:
	;
	v317 = int32(0)
	v319 = F_raxSeek(m, v11+int32(8), int32(_a_F_VM_CreateTimer_0), v317, v317)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v323 = F_raxNext(m, v11+int32(8))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v325)))
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v11)+312))
	if v326 != v327 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_raxStop(m, v11+int32(8))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L63
	}
L54:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_VM_CreateTimer[2]))
	v331 = int32(0)
	v332 = *(*int64)(unsafe.Add(mBase, _c_F_VM_CreateTimer[1]))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v330)+24))
	if v333 == v331 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_VM_CreateTimer[1])) = int64(-1)
	goto L53
L56:
	;
	goto L55
L57:
	;
	v336 = v333
	goto L58
L58:
	;
	v338 = *(*int64)(unsafe.Add(mBase, uint32(v336)))
	if v338 != v332 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L56
L60:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v336)+32))
	if v343 != 0 {
		v336 = v343
		goto L58
	} else {
		goto L62
	}
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v336))) = int64(-1)
	goto L55
L62:
	;
	goto L59
L63:
	;
	v358 = *(*int64)(unsafe.Add(mBase, _c_F_VM_CreateTimer[1]))
	if v358 != int64(-1) {
		goto L47
	} else {
		goto L64
	}
L64:
	;
	goto L48
L65:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_VM_CreateTimer[1])) = v367
	goto L47
}
func F_VM_DefragCursorGet(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_VM_DefragCursorSet(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v3))) = l1
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_VM_DefragValkeyModuleString(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
func F_VM_DictDelC(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_raxRemove(m, v5, l1, l2, l3)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int32(0))
	}
}
func F_VM_DictIteratorStart(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int64
	_ = v44
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	v31 = F_valkey_malloc(m, int32(308))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v31))) = l0
		v37 = v31 + int32(4)
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v38
		*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(2)
		*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = int32(128)
		v44 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v37)+12)) = v44
		*(*int64)(unsafe.Add(mBase, uint32(v37)+296)) = v44
		*(*int64)(unsafe.Add(mBase, uint32(v37)+160)) = int64(137438953472)
		*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v31 + int32(28)
		*(*int32)(unsafe.Add(mBase, uint32(v37)+156)) = v31 + int32(172)
		v56 = F_raxSeek(m, v37, l1, v7, v29)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			return v31
		}
	}
}
func F_VM_DictNextC(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v6 = F_raxNext(m, l0+int32(4))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			if l1 == int32(0) {
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
			}
			if l2 == int32(0) {
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
			}
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			return v20
		} else {
			return int32(0)
		}
	}
}
func F_VM_DictReplace(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	v32 = F_raxInsert(m, v30, v7, v29, l2, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v32 != int32(1))
	}
}
func F_VM_DictSet(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	v32 = F_raxTryInsert(m, v30, v7, v29, l2, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v32 != int32(1))
	}
}
func F_VM_DictSetC(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = F_raxTryInsert(m, v5, l1, l2, l3, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(1))
	}
}
func F_VM_DigestAddLongLong(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	if l1 <= int64(-1) {
		v17 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v17)
		v21 = int32(1)
		v26 = v6 + v21
		v27 = int32(20)
		v28 = int64(0) - l1
		v29 = v21
	} else {
		v26 = v6
		v27 = int32(21)
		v28 = l1
		v29 = int32(0)
	}
	v30 = F_ull2string(m, v26, v27, v28)
	mBase = m.M
	if v30 == int32(0) {
		v49 = int32(0)
	} else {
		v49 = v30 + v29
	}
	v51 = m.G0
	v52 = int32(96)
	v53 = v51 - v52
	m.G0 = v53
	F_xorDigest(m, l0, v6, v49)
	mBase = m.M
	v57 = v53 + int32(4)
	F_SHA1Init(m, v57)
	mBase = m.M
	F_SHA1Update(m, v57, l0, int32(20))
	mBase = m.M
	F_SHA1Final(m, l0, v57)
	mBase = m.M
	m.G0 = v53 + v52
	m.G0 = v6 + int32(32)
	return
}
func F_VM_DigestEndSequence(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v39 int32
	_ = v39
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v108 int64
	_ = v108
	v2 = int32(20)
	v3 = l0 + v2
	v6 = m.G0
	v7 = int32(112)
	v8 = v6 - v7
	m.G0 = v8
	v11 = v8 + v2
	F_SHA1Init(m, v11)
	mBase = m.M
	F_SHA1Update(m, v11, l0, v2)
	mBase = m.M
	F_SHA1Final(m, v8, v11)
	mBase = m.M
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v21 = v19 ^ v20
	*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v21)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	v25 = v23 ^ v24
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)) = uint8(v25)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+2)))
	v29 = v27 ^ v28
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)) = uint8(v29)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+3)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+3)))
	v33 = v31 ^ v32
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+3)) = uint8(v33)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+4)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
	v37 = v35 ^ v36
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+4)) = uint8(v37)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+5)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+5)))
	v41 = v39 ^ v40
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+5)) = uint8(v41)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+6)))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+6)))
	v45 = v43 ^ v44
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+6)) = uint8(v45)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+7)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)))
	v49 = v47 ^ v48
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+7)) = uint8(v49)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+8)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+8)))
	v53 = v51 ^ v52
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+8)) = uint8(v53)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+9)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+9)))
	v57 = v55 ^ v56
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+9)) = uint8(v57)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+10)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+10)))
	v61 = v59 ^ v60
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+10)) = uint8(v61)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+11)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)))
	v65 = v63 ^ v64
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+11)) = uint8(v65)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+12)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)))
	v69 = v67 ^ v68
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+12)) = uint8(v69)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+13)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+13)))
	v73 = v71 ^ v72
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+13)) = uint8(v73)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+14)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)))
	v77 = v75 ^ v76
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+14)) = uint8(v77)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+15)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
	v81 = v79 ^ v80
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+15)) = uint8(v81)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+16)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
	v85 = v83 ^ v84
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+16)) = uint8(v85)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+17)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
	v89 = v87 ^ v88
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+17)) = uint8(v89)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+18)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
	v93 = v91 ^ v92
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+18)) = uint8(v93)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+19)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+19)))
	v97 = v95 ^ v96
	*(*uint8)(unsafe.Add(mBase, uint32(v3)+19)) = uint8(v97)
	m.G0 = v8 + v7
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(16)))) = int32(0)
	v108 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(8)))) = v108
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v108
	return
}
func F_VM_EventLoopDel(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v5 = int32(1)
	v6 = int32(68)
	if l0 < int32(0) {
		v54 = v5
		v55 = v6
		*(*int32)(unsafe.Add(mBase, _c_F_VM_EventLoopDel[0])) = v55
		return v54
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_VM_EventLoopDel[1]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v11 <= l0 {
			v54 = v5
			v55 = v6
			*(*int32)(unsafe.Add(mBase, _c_F_VM_EventLoopDel[0])) = v55
			return v54
		} else {
			if base.Ui32(int32(3)) < base.Ui32(l1) {
				v54 = v5
				v55 = int32(28)
				*(*int32)(unsafe.Add(mBase, _c_F_VM_EventLoopDel[0])) = v55
				return v54
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_VM_EventLoopDel[1]))
				v18 = int32(0)
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				if v20 <= l0 {
					v31 = v18
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
					v25 = v22 + l0<<(uint(int32(4))%32)
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					if v26 == int32(0) {
						v31 = v18
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
						v31 = v29
					}
				}
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_VM_EventLoopDel[1]))
				F_aeDeleteFileEvent(m, v33, l0, l1)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v38 = int32(0)
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_VM_EventLoopDel[1]))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
					if v43 <= l0 {
						v50 = v38
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+l0<<(uint(int32(4))%32))))
						v50 = v49
					}
					if v50 != 0 {
						v54 = int32(0)
						v55 = v38
						*(*int32)(unsafe.Add(mBase, _c_F_VM_EventLoopDel[0])) = v55
						return v54
					} else {
						F_valkey_free(m, v31)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v54 = int32(0)
							v55 = v38
							*(*int32)(unsafe.Add(mBase, _c_F_VM_EventLoopDel[0])) = v55
							return v54
						}
					}
				}
			}
		}
	}
}
func F_VM_ExportSharedAPI(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	v6 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v10
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ExportSharedAPI[0]))
		v15 = F_dictAdd(m, v14, l1, v6)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v15 != 0 {
				F_valkey_free(m, v6)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					return int32(1)
				}
			} else {
				return int32(0)
			}
		}
	}
}
func F_VM_FreeClusterNodesList(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v6 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L10
	}
L4:
	;
	v11 = v6
	v12 = int32(0)
	goto L5
L5:
	;
	F_valkey_free(m, v11)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	return
L8:
	;
	v16 = v12 + int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0+v16<<(uint(int32(2))%32))))
	if v20 != 0 {
		v11 = v20
		v12 = v16
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	goto L1
}
func F_VM_FreeDict(m *base.Module, l0 int32, l1 int32) {
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_raxFree(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
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
	if v39 != int32(4) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v46 = v23 + v30<<(uint(int32(3))%32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v47 != int32(4) {
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
	v84 = m.ExcPending
	if v84 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	return
}
func F_VM_FreeModuleUser(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2 == int32(0) {
		F_valkey_free(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_ACLFreeUserAndKillClients(m, v5)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_valkey_free(m, l0)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	}
}
func F_VM_FreeString(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	F_decrRefCount(m, l1)
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
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = int32(3)
	v70 = v18 + int32(-1)
	if v65 == v70 {
		v80 = v65
		goto L17
	} else {
		goto L18
	}
L4:
	;
	return
L5:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v13&int32(1) == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v18 < int32(1) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v21 = int32(1)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v32 = int32(0)
	goto L8
L8:
	;
	v37 = v18 + (v32 ^ int32(-1))
	v40 = v25 + v37<<(uint(int32(3))%32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v41 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L4
L10:
	;
	v48 = v25 + v32<<(uint(int32(3))%32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v49 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v44 == l1 {
		v65 = v37
		v66 = v40
		goto L3
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v55 = v32 + int32(1)
	if v55 != int32(base.Ui32(v18+v21)>>(uint(v21)%32)) {
		v32 = v55
		goto L8
	} else {
		goto L16
	}
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v52 != l1 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v65 = v32
	v66 = v48
	goto L3
L16:
	;
	goto L9
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v80
	return
L18:
	;
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v25+v70<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v66))) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v80 = v77 + int32(-1)
	goto L17
}
func F_VM_GetAbsExpire(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
		if v8&int32(1) == int32(0) {
			v19 = int64(-1)
		} else {
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v2+(v8&int32(4)^int32(12)))))
			v19 = v18
		}
		return v19
	} else {
		return int64(-1)
	}
}
func F_VM_GetBlockedClientHandle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	return v2
}
func F_VM_GetClientUserNameById(m *base.Module, l0 int32, l1 int64) int32 {
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
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetClientUserNameById[0]))
	v52 = int32(8)
	v57 = F_raxFind(m, v51, v10+v52, v52, v10+int32(4))
	mBase = m.M
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	m.G0 = v10 + v9
	if v58 != 0 {
		v67 = *(*int32)(unsafe.Add(mBase, uint32(v58)+328))
		if v67 != 0 {
			v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
			v75 = F_sdsnew(m, v74)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				v79 = F_createObject(m, int32(0), v75)
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v81&int32(1) == int32(0) {
						return v79
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
								*(*int32)(unsafe.Add(mBase, uint32(v108))) = v79
								v110 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v110
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v104 + v110
								return v79
							}
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v104 = v86
							v105 = v89
							v108 = v105 + v104<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v108))) = v79
							v110 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v110
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v104 + v110
							return v79
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_VM_GetClientUserNameById[1])) = int32(138)
			return int32(0)
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_VM_GetClientUserNameById[1])) = int32(44)
		return int32(0)
	}
}
func F_VM_GetClusterSize(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetClusterSize[0]))
	if v2 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetClusterSize[1]))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		return v9 + v10
	} else {
		return int32(0)
	}
}
func F_VM_GetCommand(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	v3 = int32(0)
	v5 = F_lookupCommandByCString(m, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v22 = v3
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+56)))
			if v11&int32(8) == int32(0) {
				v22 = v3
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+208))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v18 == v19 {
					v21 = v16
				} else {
					v21 = int32(0)
				}
				v22 = v21
			}
		}
		return v22
	}
}
func F_VM_GetCommandKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_VM_GetCommandKeysWithFlags(m, l3, l1, l2, l3, int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_VM_GetDbIdFromModuleKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	if l0 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
		return v5
	} else {
		return int32(-1)
	}
}
func F_VM_GetDetachedThreadSafeContext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int64
	_ = v69
	v5 = F_valkey_malloc(m, int32(72))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v12 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v5+int32(24)))) = v12
		*(*int64)(unsafe.Add(mBase, uint32(v5+int32(64)))) = v12
		*(*int64)(unsafe.Add(mBase, uint32(v5+int32(56)))) = v12
		*(*int64)(unsafe.Add(mBase, uint32(v5+int32(48)))) = v12
		*(*int64)(unsafe.Add(mBase, uint32(v5+int32(40)))) = v12
		*(*int64)(unsafe.Add(mBase, uint32(v5+int32(32)))) = v12
		*(*int64)(unsafe.Add(mBase, uint32(v5+int32(16)))) = v12
		*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v9
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(561)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = int32(144)
		v46 = F_createClient(m, int32(0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v46
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+204))
			*(*int32)(unsafe.Add(mBase, uint32(v46)+204)) = v49 | int32(268435456)
			v54 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetDetachedThreadSafeContext[0]))
			v55 = int32(0)
			v56 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetDetachedThreadSafeContext[1]))
			v57 = m.T0[v56].(func(*base.Module) int64)(m)
			mBase = m.M
			if v54 == v55 {
				v69 = *(*int64)(unsafe.Add(mBase, _c_F_VM_GetDetachedThreadSafeContext[2]))
				*(*int64)(unsafe.Add(mBase, uint32(v5)+56)) = v69*int64(1000) + v57
				return v5
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetDetachedThreadSafeContext[3]))
				v63 = base.I32_div_s(int32(1000000), v62)
				*(*int64)(unsafe.Add(mBase, uint32(v5)+56)) = v57 + base.I64_extend_i32_s(v63)
				return v5
			}
		}
	}
}
func F_VM_GetKeyNameFromDefragCtx(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	return v2
}
func F_VM_GetKeyNameFromIO(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v2
}
func F_VM_GetModuleUserACLString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	if l0 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = F_ACLDescribeUser(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v11
		}
	} else {
		F__serverAssert(m, int32(_a_F_VM_GetModuleUserACLString_0), int32(_a_F_VM_GetModuleUserACLString_1), int32(10409))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
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
func F_VM_GetRandomHexChars(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	F_getRandomHexChars(m, l0, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_VM_GetSelectedDb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+200)))
	v7 = v5 & int32(8)
	if v7 != 0 {
		v8 = int32(112)
	} else {
		v8 = int32(96)
	}
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v2+v8)))
	if v7 != 0 {
		v13 = int32(52)
	} else {
		v13 = int32(28)
	}
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10+v13)))
	return v15
}
func F_VM_GetServerInfo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
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
	var v62 int32
	_ = v62
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = F_valkey_malloc(m, int32(4))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = F_raxNew(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v21
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v62
	if l1 == v62 {
		v127 = v62
		goto L14
	} else {
		goto L15
	}
L5:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v26&int32(1) == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v31 == v32 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v53 = v50 + v49<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v49 + int32(1)
	goto L4
L8:
	;
	v35 = int32(8)
	if v35 < v31 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v49 = v31
	v50 = v34
	goto L7
L10:
	;
	v38 = v31
	goto L12
L11:
	;
	v38 = v35
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v38 << (uint(int32(1)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v45 = F_valkey_realloc(m, v42, v38<<(uint(int32(4))%32))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v45
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v49 = v48
	v50 = v45
	goto L7
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v127
	v131 = int32(0)
	v138 = F_genInfoSectionDict(m, v14+int32(4), base.B2i32(l1 != v131), v131, v14+int32(12), v14+int32(8))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L39
	}
L15:
	;
	if l1&int32(3) == int32(0) {
		v91 = l1
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v125 = F_createStringObject_1(m, l1, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L32
	}
L17:
	;
	v124 = v116 - l1
	goto L16
L18:
	;
	v95 = v91
	goto L26
L19:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v77 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v80 = l1
	goto L22
L21:
	;
	v124 = l1 - l1
	goto L16
L22:
	;
	v84 = v80 + int32(1)
	if v84&int32(3) == int32(0) {
		v91 = v84
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v89 != 0 {
		v80 = v84
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v116 = v84
	goto L17
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v104 = int32(-2139062144)
	if (int32(16843008)-v101|v101)&v104 == v104 {
		v95 = v95 + int32(4)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v110 = v95
	goto L29
L28:
	;
	goto L27
L29:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v114 != 0 {
		v110 = v110 + int32(1)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v116 = v110
	goto L17
L31:
	;
	goto L30
L32:
	;
	v127 = v125
	goto L14
L33:
	;
	v166 = F_sdssplitlen(m, v142, v163, int32(_a_F_VM_GetServerInfo_0), int32(2), v14)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L41
	}
L34:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v142+int32(-17))))
	v163 = v162
	goto L33
L35:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v142+int32(-9))))
	v163 = v159
	goto L33
L36:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142+int32(-5)))))
	v163 = v156
	goto L33
L37:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142+int32(-3)))))
	v163 = v153
	goto L33
L38:
	;
	v163 = int32(base.Ui32(v146) >> (uint(int32(3)) % 32))
	goto L33
L39:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v142 = F_genValkeyInfoString(m, v138, v140, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142+int32(-1)))))
	switch v146 & int32(7) {
	case 0:
		goto L38
	case 1:
		goto L37
	case 2:
		goto L36
	case 3:
		goto L35
	case 4:
		goto L34
	default:
		v163 = v62
		goto L33
	}
L41:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v168 < int32(1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_sdsfree(m, v142)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L64
	}
L43:
	;
	v172 = int32(0)
	goto L44
L44:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v166+v172<<(uint(int32(2))%32))))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v187 == int32(35) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L42
L46:
	;
	v240 = v172 + int32(1)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v240 < v241 {
		v172 = v240
		goto L44
	} else {
		goto L63
	}
L47:
	;
	v190 = int32(58)
	v191 = F___strchrnul(m, v186, v190)
	mBase = m.M
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v193 == v190 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v197 == int32(0) {
		goto L46
	} else {
		goto L52
	}
L49:
	;
	v197 = v191
	goto L51
L50:
	;
	v197 = int32(0)
	goto L51
L51:
	;
	goto L48
L52:
	;
	v202 = v197 - v186
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+int32(-1)))))
	switch v206 & int32(7) {
	case 0:
		goto L58
	case 1:
		goto L57
	case 2:
		goto L56
	case 3:
		goto L55
	case 4:
		goto L54
	default:
		v223 = int32(0)
		goto L53
	}
L53:
	;
	v227 = F_sdsnewlen(m, v197+int32(1), v223+(v202^int32(-1)))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L59
	}
L54:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v186+int32(-17))))
	v223 = v222
	goto L53
L55:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v186+int32(-9))))
	v223 = v219
	goto L53
L56:
	;
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186+int32(-5)))))
	v223 = v216
	goto L53
L57:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+int32(-3)))))
	v223 = v213
	goto L53
L58:
	;
	v223 = int32(base.Ui32(v206) >> (uint(int32(3)) % 32))
	goto L53
L59:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v231 = F_raxTryInsert(m, v229, v186, v202, v227, int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v231 != 0 {
		goto L46
	} else {
		goto L61
	}
L61:
	;
	F_sdsfree(m, v227)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	goto L46
L63:
	;
	goto L45
L64:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	F_sdsfreesplitres(m, v166, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_releaseInfoSectionDict(m, v138)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v261 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	m.G0 = v14 + int32(16)
	return v17
L68:
	;
	F_decrRefCount(m, v261)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	goto L67
}
func F_VM_GetServerVersion(m *base.Module) int32 {
	return int32(590082)
}
func F_VM_GetThreadSafeContext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v25 int64
	_ = v25
	var v27 int32
	_ = v27
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int64
	_ = v88
	var v92 int64
	_ = v92
	v6 = F_valkey_malloc(m, int32(72))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if l0 == int32(0) {
			v32 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v6+int32(24)))) = v32
			*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v32
			v36 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v36
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(561)
			*(*int64)(unsafe.Add(mBase, uint32(v6+int32(64)))) = v32
			*(*int64)(unsafe.Add(mBase, uint32(v6+int32(56)))) = v32
			*(*int64)(unsafe.Add(mBase, uint32(v6+int32(48)))) = v32
			*(*int64)(unsafe.Add(mBase, uint32(v6+int32(40)))) = v32
			*(*int64)(unsafe.Add(mBase, uint32(v6+int32(32)))) = v32
			*(*int64)(unsafe.Add(mBase, uint32(v6+int32(16)))) = v32
			*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(144)
			v67 = F_createClient(m, v36)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v67
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+204))
				*(*int32)(unsafe.Add(mBase, uint32(v67)+204)) = v70 | int32(268435456)
				v75 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetThreadSafeContext[0]))
				v76 = int32(0)
				v77 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetThreadSafeContext[1]))
				v78 = m.T0[v77].(func(*base.Module) int64)(m)
				mBase = m.M
				if v75 == v76 {
					v88 = *(*int64)(unsafe.Add(mBase, _c_F_VM_GetThreadSafeContext[2]))
					v92 = v88*int64(1000) + v78
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetThreadSafeContext[3]))
					v84 = base.I32_div_s(int32(1000000), v83)
					v92 = v78 + base.I64_extend_i32_s(v84)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = v92
				return v6
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			F_moduleCreateContext(m, v6, v12, int32(16))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v17
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v20 = F_selectDb(m, v17, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v22 == int32(0) {
						return v6
					} else {
						v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
						*(*int64)(unsafe.Add(mBase, uint32(v17))) = v25
						v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+224)))
						*(*uint8)(unsafe.Add(mBase, uint32(v17)+224)) = uint8(v27)
						return v6
					}
				}
			}
		}
	}
}
func F_VM_GetTimerInfo(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int64
	_ = v224
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v229 int64
	_ = v229
	var v232 int64
	_ = v232
	var v234 int64
	_ = v234
	var v236 int64
	_ = v236
	var v238 int64
	_ = v238
	var v260 int64
	_ = v260
	var v261 int64
	_ = v261
	var v262 int64
	_ = v262
	var v265 int64
	_ = v265
	var v267 int64
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = l1
	v13 = int32(1)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_VM_GetTimerInfo[0]))
	v16 = int32(8)
	v17 = v10 + v16
	v20 = v10 + int32(4)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	goto L5
L1:
	;
	m.G0 = v10 + int32(16)
	return v276
L2:
	;
	if v214 == int32(0) {
		v276 = v13
		goto L1
	} else {
		goto L40
	}
L3:
	;
	if v173 != v16 {
		v214 = v5
		goto L30
	} else {
		goto L31
	}
L4:
	;
	v164 = int32(0)
	v170 = v29
	v171 = v30
	v173 = v164
	v177 = v164
	goto L3
L5:
	;
	if base.Ui32(v30) < base.Ui32(int32(8)) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v41 = v29
	v42 = v30
	v44 = int32(0)
	goto L8
L7:
	;
	v170 = v154
	v171 = v155
	v173 = v157
	v177 = base.B2i32(v160 != int32(0))
	goto L3
L8:
	;
	v50 = int32(base.Ui32(v42) >> (uint(int32(3)) % 32))
	v51 = int32(4)
	v52 = v41 + v51
	if v42&v51 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v154 = v145
	v155 = v146
	v157 = v130
	v160 = v135
	goto L7
L10:
	;
	v135 = int32(0)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v52+v50+(v135-v50)&int32(3)+v123<<(uint(int32(2))%32))))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	if base.Ui32(v146) < base.Ui32(int32(8)) {
		v154 = v145
		v155 = v146
		v157 = v130
		v160 = v135
		goto L7
	} else {
		goto L28
	}
L11:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v44))))
	v101 = int32(0)
	goto L22
L12:
	;
	v57 = int32(0)
	if base.Ui32(v16) <= base.Ui32(v44) {
		v90 = v44
		v93 = v57
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v93 == v50 {
		v123 = v57
		v130 = v90
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v67 = v44
	v70 = v57
	goto L15
L15:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v70))))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v67))))
	if v73 != v75 {
		v90 = v67
		v93 = v70
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v90 = v78
	v93 = v80
	goto L13
L17:
	;
	v77 = int32(1)
	v78 = v67 + v77
	v80 = v70 + v77
	if base.Ui32(v50) <= base.Ui32(v80) {
		v90 = v78
		v93 = v80
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if base.Ui32(v78) < base.Ui32(v16) {
		v67 = v78
		v70 = v80
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v154 = v41
	v155 = v42
	v157 = v90
	v160 = v93
	goto L7
L21:
	;
	if v101 != v50 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v101))))
	if v114 == v98&int32(255) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v116 = int32(1)
	v118 = v101 + v116
	if v118 != v50 {
		v101 = v118
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v170 = v41
	v171 = v42
	v173 = v44
	v177 = v116
	goto L3
L26:
	;
	v123 = v101
	v130 = v44 + int32(1)
	goto L10
L27:
	;
	v154 = v41
	v155 = v42
	v157 = v44
	v160 = v50
	goto L7
L28:
	;
	if base.Ui32(v130) < base.Ui32(v16) {
		v41 = v145
		v42 = v146
		v44 = v130
		goto L8
	} else {
		goto L29
	}
L29:
	;
	goto L9
L30:
	;
	goto L2
L31:
	;
	v179 = int32(0)
	if v171&int32(1) == v179 {
		v214 = v179
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v185 = v171 & int32(4)
	if v177&base.B2i32(v185 != int32(0)) != 0 {
		v214 = v179
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v189 = int32(1)
	if v20 == int32(0) {
		v214 = v189
		goto L30
	} else {
		goto L34
	}
L34:
	;
	if v171&int32(2) != 0 {
		v211 = int32(0)
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v211
	v214 = v189
	goto L30
L36:
	;
	v195 = int32(3)
	v196 = int32(base.Ui32(v171) >> (uint(v195) % 32))
	if v185 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v206 = int32(4)
	goto L39
L38:
	;
	v206 = v196 << (uint(int32(2)) % 32)
	goto L39
L39:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v170+v196+(int32(0)-v196)&v195+v206+int32(4))))
	v211 = v210
	goto L35
L40:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v219 != v220 {
		v276 = v13
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if l2 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v270 = int32(0)
	if l3 == v270 {
		v276 = v270
		goto L1
	} else {
		goto L47
	}
L43:
	;
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v225 = int64(56)
	v227 = int64(65280)
	v229 = int64(40)
	v232 = int64(16711680)
	v234 = int64(24)
	v236 = int64(4278190080)
	v238 = int64(8)
	v260 = F_ustime(m)
	mBase = m.M
	v261 = v224<<(uint(v225)%64) | v224&v227<<(uint(v229)%64) | (v224&v232<<(uint(v234)%64) | v224&v236<<(uint(v238)%64)) | (int64(base.Ui64(v224)>>(uint(v238)%64))&v236 | int64(base.Ui64(v224)>>(uint(v234)%64))&v232 | (int64(base.Ui64(v224)>>(uint(v229)%64))&v227 | int64(base.Ui64(v224)>>(uint(v225)%64)))) - v260
	v262 = int64(0)
	if v262 < v261 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v265 = v261
	goto L46
L45:
	;
	v265 = v262
	goto L46
L46:
	;
	v267 = base.I64_div_u_s(v265, int64(1000))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v267
	goto L42
L47:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v218)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v273
	v276 = v270
	goto L1
}
func F_VM_GetTypeMethodVersion(m *base.Module) int32 {
	return int32(5)
}
func F_VM_HashHasStringRef(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v4 = int32(1)
	if l0 == int32(0) {
		v21 = v4
		return v21
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7 == int32(0) {
			v21 = v4
			return v21
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			if v10&int32(15) != int32(4) {
				v21 = v4
				return v21
			} else {
				v15 = F_objectGetVal(m, l1)
				mBase = m.M
				v16 = F_hashTypeHasStringRef(m, v7, v15)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v21 = v16
					return v21
				}
			}
		}
	}
}
func F_VM_InfoAddFieldString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v11 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v14 == int32(0) {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			v29 = F_objectGetVal(m, l2)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v29
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v28
			v34 = F_sdscatfmt(m, v13, int32(_a_F_VM_InfoAddFieldString_0), v9)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				v36 = v34
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v36
				v40 = int32(0)
				m.G0 = v9 + int32(32)
				return v40
			}
		} else {
			v17 = F_objectGetVal(m, l2)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
			v23 = F_sdscatfmt(m, v13, int32(_a_F_VM_InfoAddFieldString_1), v9+int32(16))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v36 = v23
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v36
				v40 = int32(0)
				m.G0 = v9 + int32(32)
				return v40
			}
		}
	} else {
		v40 = int32(1)
		m.G0 = v9 + int32(32)
		return v40
	}
}
func F_VM_InfoAddFieldULongLong(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
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
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v27
			v32 = F_sdscatfmt(m, v13, int32(_a_F_VM_InfoAddFieldULongLong_0), v9)
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
			*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
			v22 = F_sdscatfmt(m, v13, int32(_a_F_VM_InfoAddFieldULongLong_1), v9+int32(16))
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
func F_VM_InfoBeginDictField(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v199
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v14 == int32(0) {
		v57 = v13
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v199 = int32(1)
	goto L1
L4:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v61&int32(3) == int32(0) {
		v83 = v61
		goto L19
	} else {
		goto L20
	}
L5:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-1)))))
	switch v20 & int32(7) {
	case 0:
		goto L11
	case 1:
		goto L10
	case 2:
		goto L9
	case 3:
		goto L8
	case 4:
		goto L7
	default:
		v37 = int32(0)
		goto L6
	}
L6:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v37+int32(-1)))))
	if v41 != int32(44) {
		v50 = v13
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
	v37 = v36
	goto L6
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
	v37 = v33
	goto L6
L9:
	;
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
	v37 = v30
	goto L6
L10:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
	v37 = v27
	goto L6
L11:
	;
	v37 = int32(base.Ui32(v20) >> (uint(int32(3)) % 32))
	goto L6
L12:
	;
	v52 = F_sdscat(m, v50, int32(_a_F_VM_InfoBeginDictField_0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L14
	} else {
		goto L16
	}
L13:
	;
	F_sdsIncrLen(m, v13, int32(-1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v50 = v49
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v52
	v57 = v52
	goto L4
L17:
	;
	v119 = F_getSafeInfoString(m, v61, v116, v9+int32(12))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L14
	} else {
		goto L33
	}
L18:
	;
	v116 = v108 - v61
	goto L17
L19:
	;
	v87 = v83
	goto L27
L20:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v69 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v72 = v61
	goto L23
L22:
	;
	v116 = v61 - v61
	goto L17
L23:
	;
	v76 = v72 + int32(1)
	if v76&int32(3) == int32(0) {
		v83 = v76
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v81 != 0 {
		v72 = v76
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v108 = v76
	goto L18
L27:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v96 = int32(-2139062144)
	if (int32(16843008)-v93|v93)&v96 == v96 {
		v87 = v87 + int32(4)
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v102 = v87
	goto L30
L29:
	;
	goto L28
L30:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v106 != 0 {
		v102 = v102 + int32(1)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v108 = v102
	goto L18
L32:
	;
	goto L31
L33:
	;
	if l1&int32(3) == int32(0) {
		v142 = l1
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v178 = F_getSafeInfoString(m, l1, v175, v9+int32(8))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L14
	} else {
		goto L50
	}
L35:
	;
	v175 = v167 - l1
	goto L34
L36:
	;
	v146 = v142
	goto L44
L37:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v128 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v131 = l1
	goto L40
L39:
	;
	v175 = l1 - l1
	goto L34
L40:
	;
	v135 = v131 + int32(1)
	if v135&int32(3) == int32(0) {
		v142 = v135
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v140 != 0 {
		v131 = v135
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v167 = v135
	goto L35
L44:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v155 = int32(-2139062144)
	if (int32(16843008)-v152|v152)&v155 == v155 {
		v146 = v146 + int32(4)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v161 = v146
	goto L47
L46:
	;
	goto L45
L47:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v165 != 0 {
		v161 = v161 + int32(1)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v167 = v161
	goto L35
L49:
	;
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v119
	v183 = F_sdscatfmt(m, v57, int32(_a_F_VM_InfoBeginDictField_1), v9)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v183
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v186 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v191 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	F_valkey_free(m, v186)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L14
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(1)
	v199 = int32(0)
	goto L1
L56:
	;
	F_valkey_free(m, v191)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L14
	} else {
		goto L57
	}
L57:
	;
	goto L55
}
func F_VM_KeyAtPos(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v17 int64
	_ = v17
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v47 int64
	_ = v47
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v75 int64
	_ = v75
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v103 int64
	_ = v103
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v131 int64
	_ = v131
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v144 int64
	_ = v144
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v159 int64
	_ = v159
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int64
	_ = v197
	var v208 int64
	_ = v208
	var v219 int64
	_ = v219
	var v220 int64
	_ = v220
	var v223 int64
	_ = v223
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v238 int64
	_ = v238
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v251 int64
	_ = v251
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v266 int64
	_ = v266
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v279 int64
	_ = v279
	var v287 int64
	_ = v287
	var v289 int64
	_ = v289
	var v294 int64
	_ = v294
	var v302 int64
	_ = v302
	var v304 int64
	_ = v304
	var v307 int64
	_ = v307
	var v315 int64
	_ = v315
	var v317 int64
	_ = v317
	var v322 int64
	_ = v322
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v335 int64
	_ = v335
	var v343 int64
	_ = v343
	var v345 int64
	_ = v345
	var v350 int64
	_ = v350
	var v358 int64
	_ = v358
	var v360 int64
	_ = v360
	v6 = int64(50)
	v17 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[0]))
	if base.B2i32(v17&v6 == int64(0)) == int32(0) {
		v28 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[1]))
		v29 = v28
	} else {
		v29 = int64(0)
	}
	v32 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[2]))
	if v32&v6 == int64(0) {
		v42 = v29
	} else {
		v40 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[3]))
		v42 = v40 | v29
	}
	v47 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[4]))
	if v47&v6 == int64(0) {
		v57 = v42
	} else {
		v55 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[5]))
		v57 = v55 | v42
	}
	v60 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[6]))
	if v60&v6 == int64(0) {
		v70 = v57
	} else {
		v68 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[7]))
		v70 = v68 | v57
	}
	v75 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[8]))
	if v75&v6 == int64(0) {
		v85 = v70
	} else {
		v83 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[9]))
		v85 = v83 | v70
	}
	v88 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[10]))
	if v88&v6 == int64(0) {
		v98 = v85
	} else {
		v96 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[11]))
		v98 = v96 | v85
	}
	v103 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[12]))
	if v103&v6 == int64(0) {
		v113 = v98
	} else {
		v111 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[13]))
		v113 = v111 | v98
	}
	v116 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[14]))
	if v116&v6 == int64(0) {
		v126 = v113
	} else {
		v124 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[15]))
		v126 = v124 | v113
	}
	v131 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[16]))
	if v131&v6 == int64(0) {
		v141 = v126
	} else {
		v139 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[17]))
		v141 = v139 | v126
	}
	v144 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[18]))
	if v144&v6 == int64(0) {
		v154 = v141
	} else {
		v152 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[19]))
		v154 = v152 | v141
	}
	v159 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[20]))
	if v159&v6 == int64(0) {
		v169 = v154
	} else {
		v167 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[21]))
		v169 = v167 | v154
	}
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v170&int32(2) == int32(0) {
		return
	} else {
		if l1 < int32(1) {
			return
		} else {
			v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v177 == int32(0) {
				return
			} else {
				v180 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
				v181 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
				if v180 != v181 {
					v191 = v180
					v192 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
					v195 = v192 + v191<<(uint(int32(3))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v195))) = l1
					v197 = base.I64_extend32_s(v169)
					v208 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[1]))
					if base.B2i32(v208&v197 == int64(0)) == int32(0) {
						v219 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[0]))
						v220 = v219
					} else {
						v220 = int64(0)
					}
					v223 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[3]))
					if v223&v197 == int64(0) {
						v233 = v220
					} else {
						v231 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[2]))
						v233 = v231 | v220
					}
					v238 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[5]))
					if v238&v197 == int64(0) {
						v248 = v233
					} else {
						v246 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[4]))
						v248 = v246 | v233
					}
					v251 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[7]))
					if v251&v197 == int64(0) {
						v261 = v248
					} else {
						v259 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[6]))
						v261 = v259 | v248
					}
					v266 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[9]))
					if v266&v197 == int64(0) {
						v276 = v261
					} else {
						v274 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[8]))
						v276 = v274 | v261
					}
					v279 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[11]))
					if v279&v197 == int64(0) {
						v289 = v276
					} else {
						v287 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[10]))
						v289 = v287 | v276
					}
					v294 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[13]))
					if v294&v197 == int64(0) {
						v304 = v289
					} else {
						v302 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[12]))
						v304 = v302 | v289
					}
					v307 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[15]))
					if v307&v197 == int64(0) {
						v317 = v304
					} else {
						v315 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[14]))
						v317 = v315 | v304
					}
					v322 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[17]))
					if v322&v197 == int64(0) {
						v332 = v317
					} else {
						v330 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[16]))
						v332 = v330 | v317
					}
					v335 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[19]))
					if v335&v197 == int64(0) {
						v345 = v332
					} else {
						v343 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[18]))
						v345 = v343 | v332
					}
					v350 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[21]))
					if v350&v197 == int64(0) {
						v360 = v345
					} else {
						v358 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[20]))
						v360 = v358 | v345
					}
					*(*uint32)(unsafe.Add(mBase, uint32(v195)+4)) = uint32(v360)
					*(*int32)(unsafe.Add(mBase, uint32(v177))) = v191 + int32(1)
					return
				} else {
					v183 = int32(8192)
					if v180 < v183 {
						v186 = v180
					} else {
						v186 = v183
					}
					v188 = F_getKeysPrepareResult(m, v177, v186+v180)
					mBase = m.M
					v189 = m.ExcPending
					if v189 != 0 {
						return
					} else {
						v190 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
						v191 = v190
						v192 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
						v195 = v192 + v191<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v195))) = l1
						v197 = base.I64_extend32_s(v169)
						v208 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[1]))
						if base.B2i32(v208&v197 == int64(0)) == int32(0) {
							v219 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[0]))
							v220 = v219
						} else {
							v220 = int64(0)
						}
						v223 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[3]))
						if v223&v197 == int64(0) {
							v233 = v220
						} else {
							v231 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[2]))
							v233 = v231 | v220
						}
						v238 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[5]))
						if v238&v197 == int64(0) {
							v248 = v233
						} else {
							v246 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[4]))
							v248 = v246 | v233
						}
						v251 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[7]))
						if v251&v197 == int64(0) {
							v261 = v248
						} else {
							v259 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[6]))
							v261 = v259 | v248
						}
						v266 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[9]))
						if v266&v197 == int64(0) {
							v276 = v261
						} else {
							v274 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[8]))
							v276 = v274 | v261
						}
						v279 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[11]))
						if v279&v197 == int64(0) {
							v289 = v276
						} else {
							v287 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[10]))
							v289 = v287 | v276
						}
						v294 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[13]))
						if v294&v197 == int64(0) {
							v304 = v289
						} else {
							v302 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[12]))
							v304 = v302 | v289
						}
						v307 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[15]))
						if v307&v197 == int64(0) {
							v317 = v304
						} else {
							v315 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[14]))
							v317 = v315 | v304
						}
						v322 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[17]))
						if v322&v197 == int64(0) {
							v332 = v317
						} else {
							v330 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[16]))
							v332 = v330 | v317
						}
						v335 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[19]))
						if v335&v197 == int64(0) {
							v345 = v332
						} else {
							v343 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[18]))
							v345 = v343 | v332
						}
						v350 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[21]))
						if v350&v197 == int64(0) {
							v360 = v345
						} else {
							v358 = *(*int64)(unsafe.Add(mBase, _c_F_VM_KeyAtPos[20]))
							v360 = v358 | v345
						}
						*(*uint32)(unsafe.Add(mBase, uint32(v195)+4)) = uint32(v360)
						*(*int32)(unsafe.Add(mBase, uint32(v177))) = v191 + int32(1)
						return
					}
				}
			}
		}
	}
}
func F_VM_ListDelete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	v5 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v13 = l0 + int32(24)
			F_listTypeDelete(m, v11, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = F_moduleDelKeyIfEmpty(m, l0)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					if v16 != 0 {
						return int32(0)
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						F_listTypeTryConversion(m, v18, int32(2), int32(565), l0)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							if v23 == int32(0) {
								return int32(0)
							} else {
								v26 = F_listTypeNext(m, v23, v13)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return int32(0)
								} else {
									if v26 == int32(0) {
										F_moduleFreeKeyIterator(m, l0)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return int32(0)
										} else {
											return int32(0)
										}
									} else {
										v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+5)))
										v32 = int32(0)
										if v31 != 0 {
											v36 = v32
										} else {
											v36 = int32(-1)
										}
										v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
										if v37 < int32(0) {
											v40 = base.B2i32(v31 != v32)
										} else {
											v40 = v36
										}
										*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v40 + v37
										return int32(0)
									}
								}
							}
						}
					}
				}
			}
		} else {
			return int32(1)
		}
	}
}
func F_VM_ListInsert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	if l2 != 0 {
		if l0 == int32(0) {
			v62 = int32(0)
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_listTypeTryConversionAppend(m, v63, v8+int32(12), v62, v62, int32(565), l0)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v72 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					if v72 != 0 {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						F_listTypeInsert(m, l0+int32(24), v77, int32(base.Ui32(l1)>>(uint(int32(31))%32)))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							F_moduleFreeKeyIterator(m, l0)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								v84 = v62
								m.G0 = v8 + int32(16)
								return v84
							}
						}
					} else {
						v84 = int32(1)
						m.G0 = v8 + int32(16)
						return v84
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = int32(1)
			if base.Ui32(v18) < base.Ui32(l1+v18) {
				if v17 == int32(0) {
					v62 = int32(0)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					F_listTypeTryConversionAppend(m, v63, v8+int32(12), v62, v62, int32(565), l0)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						v72 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							if v72 != 0 {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								F_listTypeInsert(m, l0+int32(24), v77, int32(base.Ui32(l1)>>(uint(int32(31))%32)))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									F_moduleFreeKeyIterator(m, l0)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										v84 = v62
										m.G0 = v8 + int32(16)
										return v84
									}
								}
							} else {
								v84 = int32(1)
								m.G0 = v8 + int32(16)
								return v84
							}
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					if v29&int32(15) != int32(1) {
						v45 = v17
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
						if v46&int32(15) != int32(1) {
							v62 = int32(0)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							F_listTypeTryConversionAppend(m, v63, v8+int32(12), v62, v62, int32(565), l0)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v72 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									if v72 != 0 {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
										F_listTypeInsert(m, l0+int32(24), v77, int32(base.Ui32(l1)>>(uint(int32(31))%32)))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											F_moduleFreeKeyIterator(m, l0)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												v84 = v62
												m.G0 = v8 + int32(16)
												return v84
											}
										}
									} else {
										v84 = int32(1)
										m.G0 = v8 + int32(16)
										return v84
									}
								}
							}
						} else {
							if l1 == int32(0) {
								v59 = F_VM_ListPush(m, l0, int32(0), l2)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v84 = v59
									m.G0 = v8 + int32(16)
									return v84
								}
							} else {
								v53 = F_listTypeLength(m, v45)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									if v53^l1 != int32(-1) {
										v62 = int32(0)
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										F_listTypeTryConversionAppend(m, v63, v8+int32(12), v62, v62, int32(565), l0)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											v72 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												if v72 != 0 {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
													F_listTypeInsert(m, l0+int32(24), v77, int32(base.Ui32(l1)>>(uint(int32(31))%32)))
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														F_moduleFreeKeyIterator(m, l0)
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
															return int32(0)
														} else {
															v84 = v62
															m.G0 = v8 + int32(16)
															return v84
														}
													}
												} else {
													v84 = int32(1)
													m.G0 = v8 + int32(16)
													return v84
												}
											}
										}
									} else {
										v59 = F_VM_ListPush(m, l0, int32(0), l2)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v84 = v59
											m.G0 = v8 + int32(16)
											return v84
										}
									}
								}
							}
						}
					} else {
						v34 = F_listTypeLength(m, v17)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							if l1 == int32(-1) {
								v40 = F_VM_ListPush(m, l0, int32(1), l2)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v84 = v40
									m.G0 = v8 + int32(16)
									return v84
								}
							} else {
								if l1 != v34 {
									v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v42 == int32(0) {
										v62 = int32(0)
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										F_listTypeTryConversionAppend(m, v63, v8+int32(12), v62, v62, int32(565), l0)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											v72 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												if v72 != 0 {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
													F_listTypeInsert(m, l0+int32(24), v77, int32(base.Ui32(l1)>>(uint(int32(31))%32)))
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														F_moduleFreeKeyIterator(m, l0)
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
															return int32(0)
														} else {
															v84 = v62
															m.G0 = v8 + int32(16)
															return v84
														}
													}
												} else {
													v84 = int32(1)
													m.G0 = v8 + int32(16)
													return v84
												}
											}
										}
									} else {
										v45 = v42
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
										if v46&int32(15) != int32(1) {
											v62 = int32(0)
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											F_listTypeTryConversionAppend(m, v63, v8+int32(12), v62, v62, int32(565), l0)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v72 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													if v72 != 0 {
														v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
														F_listTypeInsert(m, l0+int32(24), v77, int32(base.Ui32(l1)>>(uint(int32(31))%32)))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															F_moduleFreeKeyIterator(m, l0)
															mBase = m.M
															v83 = m.ExcPending
															if v83 != 0 {
																return int32(0)
															} else {
																v84 = v62
																m.G0 = v8 + int32(16)
																return v84
															}
														}
													} else {
														v84 = int32(1)
														m.G0 = v8 + int32(16)
														return v84
													}
												}
											}
										} else {
											if l1 == int32(0) {
												v59 = F_VM_ListPush(m, l0, int32(0), l2)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													v84 = v59
													m.G0 = v8 + int32(16)
													return v84
												}
											} else {
												v53 = F_listTypeLength(m, v45)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return int32(0)
												} else {
													if v53^l1 != int32(-1) {
														v62 = int32(0)
														v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														F_listTypeTryConversionAppend(m, v63, v8+int32(12), v62, v62, int32(565), l0)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															v72 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
															mBase = m.M
															v73 = m.ExcPending
															if v73 != 0 {
																return int32(0)
															} else {
																if v72 != 0 {
																	v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																	F_listTypeInsert(m, l0+int32(24), v77, int32(base.Ui32(l1)>>(uint(int32(31))%32)))
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		F_moduleFreeKeyIterator(m, l0)
																		mBase = m.M
																		v83 = m.ExcPending
																		if v83 != 0 {
																			return int32(0)
																		} else {
																			v84 = v62
																			m.G0 = v8 + int32(16)
																			return v84
																		}
																	}
																} else {
																	v84 = int32(1)
																	m.G0 = v8 + int32(16)
																	return v84
																}
															}
														}
													} else {
														v59 = F_VM_ListPush(m, l0, int32(0), l2)
														mBase = m.M
														v60 = m.ExcPending
														if v60 != 0 {
															return int32(0)
														} else {
															v84 = v59
															m.G0 = v8 + int32(16)
															return v84
														}
													}
												}
											}
										}
									}
								} else {
									v40 = F_VM_ListPush(m, l0, int32(1), l2)
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										v84 = v40
										m.G0 = v8 + int32(16)
										return v84
									}
								}
							}
						}
					}
				}
			} else {
				if v17 != 0 {
					if v17 == int32(0) {
						v62 = int32(0)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						F_listTypeTryConversionAppend(m, v63, v8+int32(12), v62, v62, int32(565), l0)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v72 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								if v72 != 0 {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
									F_listTypeInsert(m, l0+int32(24), v77, int32(base.Ui32(l1)>>(uint(int32(31))%32)))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										F_moduleFreeKeyIterator(m, l0)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											v84 = v62
											m.G0 = v8 + int32(16)
											return v84
										}
									}
								} else {
									v84 = int32(1)
									m.G0 = v8 + int32(16)
									return v84
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
						if v29&int32(15) != int32(1) {
							v45 = v17
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
							if v46&int32(15) != int32(1) {
								v62 = int32(0)
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								F_listTypeTryConversionAppend(m, v63, v8+int32(12), v62, v62, int32(565), l0)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v72 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										if v72 != 0 {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
											F_listTypeInsert(m, l0+int32(24), v77, int32(base.Ui32(l1)>>(uint(int32(31))%32)))
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												F_moduleFreeKeyIterator(m, l0)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													v84 = v62
													m.G0 = v8 + int32(16)
													return v84
												}
											}
										} else {
											v84 = int32(1)
											m.G0 = v8 + int32(16)
											return v84
										}
									}
								}
							} else {
								if l1 == int32(0) {
									v59 = F_VM_ListPush(m, l0, int32(0), l2)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v84 = v59
										m.G0 = v8 + int32(16)
										return v84
									}
								} else {
									v53 = F_listTypeLength(m, v45)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										if v53^l1 != int32(-1) {
											v62 = int32(0)
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											F_listTypeTryConversionAppend(m, v63, v8+int32(12), v62, v62, int32(565), l0)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v72 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													if v72 != 0 {
														v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
														F_listTypeInsert(m, l0+int32(24), v77, int32(base.Ui32(l1)>>(uint(int32(31))%32)))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															F_moduleFreeKeyIterator(m, l0)
															mBase = m.M
															v83 = m.ExcPending
															if v83 != 0 {
																return int32(0)
															} else {
																v84 = v62
																m.G0 = v8 + int32(16)
																return v84
															}
														}
													} else {
														v84 = int32(1)
														m.G0 = v8 + int32(16)
														return v84
													}
												}
											}
										} else {
											v59 = F_VM_ListPush(m, l0, int32(0), l2)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int32(0)
											} else {
												v84 = v59
												m.G0 = v8 + int32(16)
												return v84
											}
										}
									}
								}
							}
						} else {
							v34 = F_listTypeLength(m, v17)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								if l1 == int32(-1) {
									v40 = F_VM_ListPush(m, l0, int32(1), l2)
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										v84 = v40
										m.G0 = v8 + int32(16)
										return v84
									}
								} else {
									if l1 != v34 {
										v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v42 == int32(0) {
											v62 = int32(0)
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											F_listTypeTryConversionAppend(m, v63, v8+int32(12), v62, v62, int32(565), l0)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v72 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													if v72 != 0 {
														v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
														F_listTypeInsert(m, l0+int32(24), v77, int32(base.Ui32(l1)>>(uint(int32(31))%32)))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															F_moduleFreeKeyIterator(m, l0)
															mBase = m.M
															v83 = m.ExcPending
															if v83 != 0 {
																return int32(0)
															} else {
																v84 = v62
																m.G0 = v8 + int32(16)
																return v84
															}
														}
													} else {
														v84 = int32(1)
														m.G0 = v8 + int32(16)
														return v84
													}
												}
											}
										} else {
											v45 = v42
											v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
											if v46&int32(15) != int32(1) {
												v62 = int32(0)
												v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												F_listTypeTryConversionAppend(m, v63, v8+int32(12), v62, v62, int32(565), l0)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													v72 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int32(0)
													} else {
														if v72 != 0 {
															v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
															F_listTypeInsert(m, l0+int32(24), v77, int32(base.Ui32(l1)>>(uint(int32(31))%32)))
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																F_moduleFreeKeyIterator(m, l0)
																mBase = m.M
																v83 = m.ExcPending
																if v83 != 0 {
																	return int32(0)
																} else {
																	v84 = v62
																	m.G0 = v8 + int32(16)
																	return v84
																}
															}
														} else {
															v84 = int32(1)
															m.G0 = v8 + int32(16)
															return v84
														}
													}
												}
											} else {
												if l1 == int32(0) {
													v59 = F_VM_ListPush(m, l0, int32(0), l2)
													mBase = m.M
													v60 = m.ExcPending
													if v60 != 0 {
														return int32(0)
													} else {
														v84 = v59
														m.G0 = v8 + int32(16)
														return v84
													}
												} else {
													v53 = F_listTypeLength(m, v45)
													mBase = m.M
													v54 = m.ExcPending
													if v54 != 0 {
														return int32(0)
													} else {
														if v53^l1 != int32(-1) {
															v62 = int32(0)
															v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															F_listTypeTryConversionAppend(m, v63, v8+int32(12), v62, v62, int32(565), l0)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return int32(0)
															} else {
																v72 = F_moduleListIteratorSeek(m, l0, l1, int32(2))
																mBase = m.M
																v73 = m.ExcPending
																if v73 != 0 {
																	return int32(0)
																} else {
																	if v72 != 0 {
																		v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																		F_listTypeInsert(m, l0+int32(24), v77, int32(base.Ui32(l1)>>(uint(int32(31))%32)))
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return int32(0)
																		} else {
																			F_moduleFreeKeyIterator(m, l0)
																			mBase = m.M
																			v83 = m.ExcPending
																			if v83 != 0 {
																				return int32(0)
																			} else {
																				v84 = v62
																				m.G0 = v8 + int32(16)
																				return v84
																			}
																		}
																	} else {
																		v84 = int32(1)
																		m.G0 = v8 + int32(16)
																		return v84
																	}
																}
															}
														} else {
															v59 = F_VM_ListPush(m, l0, int32(0), l2)
															mBase = m.M
															v60 = m.ExcPending
															if v60 != 0 {
																return int32(0)
															} else {
																v84 = v59
																m.G0 = v8 + int32(16)
																return v84
															}
														}
													}
												}
											}
										}
									} else {
										v40 = F_VM_ListPush(m, l0, int32(1), l2)
										mBase = m.M
										v41 = m.ExcPending
										if v41 != 0 {
											return int32(0)
										} else {
											v84 = v40
											m.G0 = v8 + int32(16)
											return v84
										}
									}
								}
							}
						}
					}
				} else {
					v23 = F_VM_ListPush(m, l0, int32(1), l2)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v84 = v23
						m.G0 = v8 + int32(16)
						return v84
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_VM_ListInsert[0])) = int32(28)
		v84 = int32(1)
		m.G0 = v8 + int32(16)
		return v84
	}
}
func F_VM_LoadDataTypeFromString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
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
	var v42 int32
	_ = v42
	v4 = m.G0
	v6 = v4 - int32(112)
	m.G0 = v6
	v10 = F_objectGetVal(m, l0)
	mBase = m.M
	v13 = F___memcpy(m, v6+int32(32), int32(_a_F_VM_LoadDataTypeFromString_0), int32(80))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v6)+20)) = int64(-4294967296)
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v6)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v6 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v31 = m.T0[v30].(func(*base.Module, int32, int32) int32)(m, v6, v19)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
		if v35 == int32(0) {
			m.G0 = v6 + int32(112)
			return v31
		} else {
			F_moduleFreeContext(m, v35)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
				F_valkey_free(m, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(112)
					return v31
				}
			}
		}
	}
}
func F_VM_LoadDouble(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v61 float64
	_ = v61
	v3 = float64(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v12 != 0 {
		v61 = v3
		m.G0 = v9 + int32(32)
		return v61
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v15 = F_rdbLoadLen(m, v13, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return float64(0)
		} else {
			if v15 != int64(4) {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+48)))
				if v31&int32(1) != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
					v61 = v3
					m.G0 = v9 + int32(32)
					return v61
				} else {
					v34 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v36 != 0 {
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_VM_LoadDouble[0]))
						if v39 == int32(0) {
							v43 = F_objectGetVal(m, v36)
							mBase = m.M
							v44 = v43
						} else {
							v44 = int32(_a_F_VM_LoadDouble_0)
						}
					} else {
						v44 = int32(_a_F_VM_LoadDouble_1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v44
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v35
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 + int32(84)
					F__serverPanic_1(m, int32(_a_F_VM_LoadDouble_2), int32(7585), int32(_a_F_VM_LoadDouble_3), v9)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return float64(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v24 = F_rdbLoadBinaryDoubleValue(m, v21, v9+int32(24))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return float64(0)
				} else {
					if v24 == int32(-1) {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+48)))
						if v31&int32(1) != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
							v61 = v3
							m.G0 = v9 + int32(32)
							return v61
						} else {
							v34 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v36 != 0 {
								v39 = *(*int32)(unsafe.Add(mBase, _c_F_VM_LoadDouble[0]))
								if v39 == int32(0) {
									v43 = F_objectGetVal(m, v36)
									mBase = m.M
									v44 = v43
								} else {
									v44 = int32(_a_F_VM_LoadDouble_0)
								}
							} else {
								v44 = int32(_a_F_VM_LoadDouble_1)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v44
							*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 + int32(84)
							F__serverPanic_1(m, int32(_a_F_VM_LoadDouble_2), int32(7585), int32(_a_F_VM_LoadDouble_3), v9)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return float64(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					} else {
						v28 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
						v61 = v28
						m.G0 = v9 + int32(32)
						return v61
					}
				}
			}
		}
	}
}
func F_VM_LoadSigned(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_VM_LoadUnsigned(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_VM_LoadString(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = F_moduleLoadString(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_VM_LoadUnsigned(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v61 int64
	_ = v61
	v3 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 != 0 {
		v61 = v3
		m.G0 = v8 + int32(32)
		return v61
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v14 = F_rdbLoadLen(m, v12, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			if v14 != int64(2) {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+48)))
				if v31&int32(1) != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
					v61 = v3
					m.G0 = v8 + int32(32)
					return v61
				} else {
					v34 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v36 != 0 {
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_VM_LoadUnsigned[0]))
						if v39 == int32(0) {
							v43 = F_objectGetVal(m, v36)
							mBase = m.M
							v44 = v43
						} else {
							v44 = int32(_a_F_VM_LoadUnsigned_0)
						}
					} else {
						v44 = int32(_a_F_VM_LoadUnsigned_1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8+int32(16)))) = v44
					*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v35
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v29 + int32(84)
					F__serverPanic_1(m, int32(_a_F_VM_LoadUnsigned_2), int32(7585), int32(_a_F_VM_LoadUnsigned_3), v8)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int64(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v24 = F_rdbLoadLenByRef(m, v20, int32(0), v8+int32(24))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					if v24 == int32(-1) {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+48)))
						if v31&int32(1) != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
							v61 = v3
							m.G0 = v8 + int32(32)
							return v61
						} else {
							v34 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v36 != 0 {
								v39 = *(*int32)(unsafe.Add(mBase, _c_F_VM_LoadUnsigned[0]))
								if v39 == int32(0) {
									v43 = F_objectGetVal(m, v36)
									mBase = m.M
									v44 = v43
								} else {
									v44 = int32(_a_F_VM_LoadUnsigned_0)
								}
							} else {
								v44 = int32(_a_F_VM_LoadUnsigned_1)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v8+int32(16)))) = v44
							*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v29 + int32(84)
							F__serverPanic_1(m, int32(_a_F_VM_LoadUnsigned_2), int32(7585), int32(_a_F_VM_LoadUnsigned_3), v8)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int64(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					} else {
						v28 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
						v61 = v28
						m.G0 = v8 + int32(32)
						return v61
					}
				}
			}
		}
	}
}
func F_VM_LogIOError(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v14 int32
	_ = v14
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	F_moduleLogRaw(m, v12, l1, l2, l3)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		m.G0 = v8 + int32(16)
		return
	}
}
func F_VM_MallocSizeString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v2&int32(15) == int32(0) {
		v15 = F_getStringObjectSdsUsedMemory(m, l0)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15 + int32(12)
		}
	} else {
		F__serverAssert(m, int32(_a_F_VM_MallocSizeString_0), int32(_a_F_VM_MallocSizeString_1), int32(11860))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
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
func F_VM_Milliseconds(m *base.Module) int64 {
	var v1 int64
	_ = v1
	v1 = F_mstime(m)
	return v1
}
func F_VM_ModuleTypeGetValue(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v2 = int32(0)
	if l0 == v2 {
		v17 = v2
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v6 == int32(0) {
			v17 = v2
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			if v9&int32(15) != int32(5) {
				v17 = v2
			} else {
				v14 = F_objectGetVal(m, v6)
				mBase = m.M
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				v17 = v15
			}
		}
	}
	return v17
}
func F_VM_ModuleTypeReplaceValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	v6 = int32(1)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v7&int32(2) == int32(0) {
		v31 = v6
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v12 != 0 {
			v31 = v6
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v13 == int32(0) {
				v31 = v6
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v16&int32(15) != int32(5) {
					v31 = v6
				} else {
					v21 = F_objectGetVal(m, v13)
					mBase = m.M
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
					if v22 != l1 {
						v31 = v6
					} else {
						if l3 == int32(0) {
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v26
						}
						*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = l2
						v31 = int32(0)
					}
				}
			}
		}
	}
	return v31
}
func F_VM_MustObeyClient(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v10 int64
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v2 = int32(0)
	if l0 == v2 {
		v27 = v2
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v6 == int32(0) {
			v27 = v2
		} else {
			v10 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
			if v10 != int64(-1) {
				v14 = int32(1)
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+200)))
				if v15&v14 != 0 {
					v22 = v14
					v25 = v22
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+216))
					if v18 != 0 {
						v20 = F_isImportSlotMigrationJob(m, v18)
						mBase = m.M
						v22 = v20
						v25 = v22
					} else {
						v25 = int32(0)
					}
				}
			} else {
				v25 = int32(1)
			}
			v27 = v25
		}
	}
	return v27
}
func F_VM_RandomKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+96))
	v7 = F_dbRandomKey(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
		if v11&int32(1) == int32(0) {
			return v7
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
					*(*int32)(unsafe.Add(mBase, uint32(v38))) = v7
					v40 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v40
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v34 + v40
					return v7
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v34 = v16
				v35 = v19
				v38 = v35 + v34<<(uint(int32(3))%32)
				*(*int32)(unsafe.Add(mBase, uint32(v38))) = v7
				v40 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v40
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v34 + v40
				return v7
			}
		}
	}
}
func F_VM_RdbLoad(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	v6 = int32(1)
	v7 = int32(28)
	if l1 == int32(0) {
		v75 = v6
		v76 = v7
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_VM_RdbLoad_0), int32(_a_F_VM_RdbLoad_1), int32(14427))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L7
	} else {
		goto L36
	}
L2:
	;
	goto L35
L3:
	;
	if l2 != 0 {
		v75 = v6
		v76 = v7
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_VM_RdbLoad[1]))
	if v11 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_disconnectReplicas(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v75 = v6
	v76 = int32(138)
	goto L2
L7:
	;
	return int32(0)
L8:
	;
	F_freeReplicationBacklog(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_VM_RdbLoad[2]))
	if v22 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_VM_RdbLoad[3]))
	if v28 != int32(1) {
		v35 = v28
		goto L13
	} else {
		goto L14
	}
L11:
	;
	F_stopAppendOnly(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	if v35 != int32(5) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	F_killRDBChild(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_VM_RdbLoad[3]))
	v35 = v34
	goto L13
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_VM_RdbLoad[4]))
	if v41 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	F_killSlotMigrationChild(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v46 != int32(1) {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	F_protectClient(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v52 = F_rdbLoad(m, v49, int32(0), int32(32))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_VM_RdbLoad[4]))
	if v55 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_VM_RdbLoad[5]))
	if v61 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	F_unprotectClient(m, v55)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	if v52 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v64 = F_startAppendOnly(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v68 = int32(1)
	if v52 == v68 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v66 = int32(0)
	v75 = v66
	v76 = v66
	goto L2
L32:
	;
	v73 = int32(44)
	goto L34
L33:
	;
	v73 = int32(29)
	goto L34
L34:
	;
	v75 = v68
	v76 = v73
	goto L2
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_RdbLoad[0])) = v76
	return v75
L36:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_VM_RegisterCommandFilter(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	v7 = F_valkey_malloc(m, int32(12))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v11
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_VM_RegisterCommandFilter[0]))
		v17 = F_listAddNodeTail(m, v16, v7)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
			v21 = F_listAddNodeTail(m, v20, v7)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return v7
			}
		}
	}
}
func F_VM_RegisterEnumConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v114 int32
	_ = v114
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = F_moduleConfigValidityCheck(m, v16, l1, l3, int32(4))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v114
L2:
	;
	return int32(0)
L3:
	;
	if v18 != 0 {
		v114 = int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v23 = F_valkey_malloc(m, int32(24))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v25 = F_sdsnew(m, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = l9
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = l7
	v35 = l6 << (uint(int32(3)) % 32)
	v38 = F_valkey_malloc(m, v35+int32(8))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if l6 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38+v35))) = int64(0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	v90 = F_listAddNodeTail(m, v89, v23)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L14
	}
L9:
	;
	v53 = int32(0)
	goto L10
L10:
	;
	v58 = v38 + v53<<(uint(int32(3))%32)
	v60 = v53 << (uint(int32(2)) % 32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l4+v60)))
	v63 = F_zstrdup(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v63
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l5+v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v67
	v70 = v53 + int32(1)
	if v70 != l6 {
		v53 = v70
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	F_addModuleEnumConfig(m, v92, l1, int32(base.Ui32(l3)>>(uint(int32(5))%32))&int32(8)|l3&int32(113), v23, l2, v38)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v114 = int32(0)
	goto L1
}
func F_VM_RegisterInfoFunc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = l1
	return int32(0)
}
func F_VM_RegisterScriptingEngine(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_VM_RegisterScriptingEngine[0]))
	if int32(0) < v12 {
		v24 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
		if base.Ui64(v24) < base.Ui64(int64(5)) {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v40 = F_scriptingEngineManagerRegister(m, l1, v39, l2, l3)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v44 = base.B2i32(v40 != int32(0))
				m.G0 = v9 + int32(32)
				return v44
			}
		} else {
			v27 = int32(1)
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_VM_RegisterScriptingEngine[0]))
			if int32(3) < v29 {
				v44 = v27
				m.G0 = v9 + int32(32)
				return v44
			} else {
				*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v24)
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(4)
				F__serverLog(m, int32(3), int32(_a_F_VM_RegisterScriptingEngine_0), v9)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v44 = v27
					m.G0 = v9 + int32(32)
					return v44
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
		F__serverLog(m, int32(0), int32(_a_F_VM_RegisterScriptingEngine_1), v9+int32(16))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
			if base.Ui64(v24) < base.Ui64(int64(5)) {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v40 = F_scriptingEngineManagerRegister(m, l1, v39, l2, l3)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v44 = base.B2i32(v40 != int32(0))
					m.G0 = v9 + int32(32)
					return v44
				}
			} else {
				v27 = int32(1)
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_VM_RegisterScriptingEngine[0]))
				if int32(3) < v29 {
					v44 = v27
					m.G0 = v9 + int32(32)
					return v44
				} else {
					*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v24)
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(4)
					F__serverLog(m, int32(3), int32(_a_F_VM_RegisterScriptingEngine_0), v9)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v44 = v27
						m.G0 = v9 + int32(32)
						return v44
					}
				}
			}
		}
	}
}
func F_VM_ReplyWithAttribute(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+224)))
	if v6 == int32(2) {
		v15 = int32(1)
		return v15
	} else {
		v10 = F_moduleReplyWithCollection(m, l0, l1, int32(4))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = int32(0)
			return v15
		}
	}
}
func F_VM_ReplyWithBigNumber(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v4&int32(16) == int32(0) {
		v16 = l0 + int32(8)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v17 == int32(0) {
			return int32(0)
		} else {
			F_addReplyBigNum(m, v17, l1, l2)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v9 == int32(0) {
			return int32(0)
		} else {
			v16 = v9 + int32(36)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v17 == int32(0) {
				return int32(0)
			} else {
				F_addReplyBigNum(m, v17, l1, l2)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_VM_ReplyWithDouble(m *base.Module, l0 int32, l1 float64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v3&int32(16) == int32(0) {
		v15 = l0 + int32(8)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		if v16 == int32(0) {
			return int32(0)
		} else {
			F_addReplyDouble(m, v16, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 == int32(0) {
			return int32(0)
		} else {
			v15 = v8 + int32(36)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if v16 == int32(0) {
				return int32(0)
			} else {
				F_addReplyDouble(m, v16, l1)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_VM_ReplyWithEmptyString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v2&int32(16) == int32(0) {
		v14 = l0 + int32(8)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		if v15 == int32(0) {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ReplyWithEmptyString[0]))
			F_addReply(m, v15, v19)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7 == int32(0) {
			return int32(0)
		} else {
			v14 = v7 + int32(36)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			if v15 == int32(0) {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ReplyWithEmptyString[0]))
				F_addReply(m, v15, v19)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_VM_ReplyWithError(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	v4 = m.G0
	v5 = int32(16)
	v6 = v4 - v5
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v8&v5 == int32(0) {
		v20 = l0 + int32(8)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		if v21 == int32(0) {
			m.G0 = v6 + int32(16)
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
			F_addReplyErrorFormat(m, v21, int32(_a_F_VM_ReplyWithError_0), v6)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return int32(0)
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v13 == int32(0) {
			m.G0 = v6 + int32(16)
			return int32(0)
		} else {
			v20 = v13 + int32(36)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			if v21 == int32(0) {
				m.G0 = v6 + int32(16)
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
				F_addReplyErrorFormat(m, v21, int32(_a_F_VM_ReplyWithError_0), v6)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return int32(0)
				}
			}
		}
	}
}
func F_VM_ReplyWithErrorFormat(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	F_moduleReplyErrorFormatInternal(m, l0, int32(0), l1, l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return int32(0)
	}
}
func F_VM_ReplyWithLongLong(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v3&int32(16) == int32(0) {
		v15 = l0 + int32(8)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		if v16 == int32(0) {
			return int32(0)
		} else {
			F_addReplyLongLong(m, v16, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 == int32(0) {
			return int32(0)
		} else {
			v15 = v8 + int32(36)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if v16 == int32(0) {
				return int32(0)
			} else {
				F_addReplyLongLong(m, v16, l1)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_VM_ReplyWithMap(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_moduleReplyWithCollection(m, l0, l1, int32(2))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_VM_ReplyWithNull(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v2&int32(16) == int32(0) {
		v14 = l0 + int32(8)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		if v15 == int32(0) {
			return int32(0)
		} else {
			F_addReplyNull(m, v15)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7 == int32(0) {
			return int32(0)
		} else {
			v14 = v7 + int32(36)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			if v15 == int32(0) {
				return int32(0)
			} else {
				F_addReplyNull(m, v15)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_VM_ReplyWithNullArray(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v2&int32(16) == int32(0) {
		v14 = l0 + int32(8)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		if v15 == int32(0) {
			return int32(0)
		} else {
			F_addReplyNullArray(m, v15)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7 == int32(0) {
			return int32(0)
		} else {
			v14 = v7 + int32(36)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			if v15 == int32(0) {
				return int32(0)
			} else {
				F_addReplyNullArray(m, v15)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_VM_ReplyWithStringBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v4&int32(16) == int32(0) {
		v16 = l0 + int32(8)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v17 == int32(0) {
			return int32(0)
		} else {
			F_addReplyBulkCBuffer(m, v17, l1, l2)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v9 == int32(0) {
			return int32(0)
		} else {
			v16 = v9 + int32(36)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v17 == int32(0) {
				return int32(0)
			} else {
				F_addReplyBulkCBuffer(m, v17, l1, l2)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_VM_ReplyWithVerbatimString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v4&int32(16) == int32(0) {
		v16 = l0 + int32(8)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v17 == int32(0) {
			return int32(0)
		} else {
			F_addReplyVerbatim(m, v17, l1, l2, int32(_a_F_VM_ReplyWithVerbatimString_0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v9 == int32(0) {
			return int32(0)
		} else {
			v16 = v9 + int32(36)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v17 == int32(0) {
				return int32(0)
			} else {
				F_addReplyVerbatim(m, v17, l1, l2, int32(_a_F_VM_ReplyWithVerbatimString_0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_VM_ReplyWithVerbatimStringType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v5&int32(16) == int32(0) {
		v17 = l0 + int32(8)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		if v18 == int32(0) {
			return int32(0)
		} else {
			F_addReplyVerbatim(m, v18, l1, l2, l3)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v10 == int32(0) {
			return int32(0)
		} else {
			v17 = v10 + int32(36)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			if v18 == int32(0) {
				return int32(0)
			} else {
				F_addReplyVerbatim(m, v18, l1, l2, l3)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_VM_ResetDataset(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	if l0 == int32(0) {
		F_flushAllDataAndResetRDB(m, base.B2i32(l1 != int32(0))|int32(2))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if l0 == int32(0) {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ResetDataset[0]))
				if v20 == int32(0) {
					return
				} else {
					F_restartAOFAfterSYNC(m)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ResetDataset[1]))
		if v6 == int32(0) {
			F_flushAllDataAndResetRDB(m, base.B2i32(l1 != int32(0))|int32(2))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				if l0 == int32(0) {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ResetDataset[0]))
					if v20 == int32(0) {
						return
					} else {
						F_restartAOFAfterSYNC(m)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			F_stopAppendOnly(m)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_flushAllDataAndResetRDB(m, base.B2i32(l1 != int32(0))|int32(2))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					if l0 == int32(0) {
						return
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ResetDataset[0]))
						if v20 == int32(0) {
							return
						} else {
							F_restartAOFAfterSYNC(m)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
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
func F_VM_SaveDouble(m *base.Module, l0 int32, l1 float64) {
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
			v51 = F_rdbSaveLen(m, v49, int64(4))
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
					v59 = F_rdbSaveBinaryDoubleValue(m, v58, l1)
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
							v51 = F_rdbSaveLen(m, v49, int64(4))
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
									v59 = F_rdbSaveBinaryDoubleValue(m, v58, l1)
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
						v51 = F_rdbSaveLen(m, v49, int64(4))
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
								v59 = F_rdbSaveBinaryDoubleValue(m, v58, l1)
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
func F_VM_SaveLongDouble(m *base.Module, l0 int32, l1 int64, l2 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v5 = m.G0
	v7 = v5 - int32(5120)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		m.G0 = v7 + int32(5120)
		return
	} else {
		v12 = F_ld2string(m, v7, int32(5120), l1, l2, int32(2))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_VM_SaveStringBuffer(m, l0, v7, v12)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				m.G0 = v7 + int32(5120)
				return
			}
		}
	}
}
func F_VM_Scan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v11 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+96))
		v23 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		v27 = F_dbScan(m, v22, v23, int32(569), v9+int32(4))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v27
			if v27 != int64(0) {
				v38 = int32(1)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
				v38 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, _c_F_VM_Scan[0])) = int32(0)
			v42 = v38
			m.G0 = v9 + int32(16)
			return v42
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_VM_Scan[0])) = int32(44)
		v42 = int32(0)
		m.G0 = v9 + int32(16)
		return v42
	}
}
func F_VM_ScanCursorRestart(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	return
}
func F_VM_ScriptingEngineDebuggerLogRespReplyStr(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_scriptingEngineDebuggerLogRespReplyStr(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_VM_ScriptingEngineDebuggerProcessCommands(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	F_scriptingEngineDebuggerProcessCommands(m, l0, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_VM_SendClusterMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int64
	_ = v69
	var v73 int64
	_ = v73
	var v114 int64
	_ = v114
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SendClusterMessage[0]))
	if v7 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v12 = int32(0)
		v14 = int64(0)
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SendClusterMessage[1]))
		v26 = F_strlen(m, v11)
		mBase = m.M
		if v26 != int32(9) {
			v114 = v14
		} else {
			v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11))))
			v32 = F_strchr(m, v24, v31)
			mBase = m.M
			if v32 == int32(0) {
				v114 = v14
			} else {
				v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+1)))
				v36 = F_strchr(m, v24, v35)
				mBase = m.M
				if v36 == int32(0) {
					v114 = v14
				} else {
					v39 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+2)))
					v40 = F_strchr(m, v24, v39)
					mBase = m.M
					if v40 == int32(0) {
						v114 = v14
					} else {
						v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+3)))
						v44 = F_strchr(m, v24, v43)
						mBase = m.M
						if v44 == int32(0) {
							v114 = v14
						} else {
							v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+4)))
							v48 = F_strchr(m, v24, v47)
							mBase = m.M
							if v48 == int32(0) {
								v114 = v14
							} else {
								v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+5)))
								v52 = F_strchr(m, v24, v51)
								mBase = m.M
								if v52 == int32(0) {
									v114 = v14
								} else {
									v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+6)))
									v56 = F_strchr(m, v24, v55)
									mBase = m.M
									if v56 == int32(0) {
										v114 = v14
									} else {
										v59 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+7)))
										v60 = F_strchr(m, v24, v59)
										mBase = m.M
										if v60 == int32(0) {
											v114 = v14
										} else {
											v63 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+8)))
											v64 = F_strchr(m, v24, v63)
											mBase = m.M
											if v64 == int32(0) {
												v114 = v14
											} else {
												v69 = int64(12)
												v73 = int64(6)
												v114 = ((((base.I64_extend_i32_u(v32-v24)<<(uint(v69)%64)|base.I64_extend_i32_u(v36-v24)<<(uint(v73)%64)|base.I64_extend_i32_u(v40-v24))<<(uint(v69)%64)|base.I64_extend_i32_u(v44-v24)<<(uint(v73)%64)|base.I64_extend_i32_u(v48-v24))<<(uint(v69)%64)|base.I64_extend_i32_u(v52-v24)<<(uint(v73)%64)|base.I64_extend_i32_u(v56-v24))<<(uint(v69)%64)|base.I64_extend_i32_u(v60-v24)<<(uint(v73)%64)|base.I64_extend_i32_u(v64-v24))<<(uint(int64(10))%64) | base.I64_extend_i32_u(v12)
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
		v123 = F_clusterSendModuleMessageToTarget(m, l1, v114, l2, l3, l4)
		mBase = m.M
		v126 = m.ExcPending
		if v126 != 0 {
			return int32(0)
		} else {
			return base.B2i32(v123 != int32(0))
		}
	} else {
		return int32(1)
	}
}
func F_VM_ServerInfoGetFieldC(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v232 int32
	_ = v232
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
	var v263 int32
	_ = v263
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v3
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1&int32(3) == v3 {
		v32 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v67 = v6 + int32(12)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v65 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v65 = v57 - l1
	goto L1
L3:
	;
	v36 = v32
	goto L11
L4:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = l1
	goto L7
L6:
	;
	v65 = l1 - l1
	goto L1
L7:
	;
	v25 = v21 + int32(1)
	if v25&int32(3) == int32(0) {
		v32 = v25
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v30 != 0 {
		v21 = v25
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v57 = v25
	goto L2
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v45 = int32(-2139062144)
	if (int32(16843008)-v42|v42)&v45 == v45 {
		v36 = v36 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v51 = v36
	goto L14
L13:
	;
	goto L12
L14:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 != 0 {
		v51 = v51 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v57 = v51
	goto L2
L16:
	;
	goto L15
L17:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	m.G0 = v6 + int32(16)
	return v263
L18:
	;
	if v220 != v65 {
		goto L45
	} else {
		goto L46
	}
L19:
	;
	v211 = int32(0)
	v217 = v76
	v218 = v77
	v220 = v211
	v224 = v211
	goto L18
L20:
	;
	if base.Ui32(v77) < base.Ui32(int32(8)) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v88 = v76
	v89 = v77
	v91 = int32(0)
	goto L23
L22:
	;
	v217 = v201
	v218 = v202
	v220 = v204
	v224 = base.B2i32(v207 != int32(0))
	goto L18
L23:
	;
	v97 = int32(base.Ui32(v89) >> (uint(int32(3)) % 32))
	v98 = int32(4)
	v99 = v88 + v98
	if v89&v98 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v201 = v192
	v202 = v193
	v204 = v177
	v207 = v182
	goto L22
L25:
	;
	v182 = int32(0)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v99+v97+(v182-v97)&int32(3)+v170<<(uint(int32(2))%32))))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	if base.Ui32(v193) < base.Ui32(int32(8)) {
		v201 = v192
		v202 = v193
		v204 = v177
		v207 = v182
		goto L22
	} else {
		goto L43
	}
L26:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v91))))
	v148 = int32(0)
	goto L37
L27:
	;
	v104 = int32(0)
	if base.Ui32(v65) <= base.Ui32(v91) {
		v137 = v91
		v140 = v104
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v140 == v97 {
		v170 = v104
		v177 = v137
		goto L25
	} else {
		goto L35
	}
L29:
	;
	v114 = v91
	v117 = v104
	goto L30
L30:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99+v117))))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v114))))
	if v120 != v122 {
		v137 = v114
		v140 = v117
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v137 = v125
	v140 = v127
	goto L28
L32:
	;
	v124 = int32(1)
	v125 = v114 + v124
	v127 = v117 + v124
	if base.Ui32(v97) <= base.Ui32(v127) {
		v137 = v125
		v140 = v127
		goto L28
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32(v125) < base.Ui32(v65) {
		v114 = v125
		v117 = v127
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v201 = v88
	v202 = v89
	v204 = v137
	v207 = v140
	goto L22
L36:
	;
	if v148 != v97 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99+v148))))
	if v161 == v145&int32(255) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v163 = int32(1)
	v165 = v148 + v163
	if v165 != v97 {
		v148 = v165
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v217 = v88
	v218 = v89
	v220 = v91
	v224 = v163
	goto L18
L41:
	;
	v170 = v148
	v177 = v91 + int32(1)
	goto L25
L42:
	;
	v201 = v88
	v202 = v89
	v204 = v91
	v207 = v97
	goto L22
L43:
	;
	if base.Ui32(v177) < base.Ui32(v65) {
		v88 = v192
		v89 = v193
		v91 = v177
		goto L23
	} else {
		goto L44
	}
L44:
	;
	goto L24
L45:
	;
	goto L17
L46:
	;
	if v218&int32(1) == int32(0) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v232 = v218 & int32(4)
	if v224&base.B2i32(v232 != int32(0)) != 0 {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	if v67 == int32(0) {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	if v218&int32(2) != 0 {
		v258 = int32(0)
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v258
	goto L45
L51:
	;
	v242 = int32(3)
	v243 = int32(base.Ui32(v218) >> (uint(v242) % 32))
	if v232 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v253 = int32(4)
	goto L54
L53:
	;
	v253 = v243 << (uint(int32(2)) % 32)
	goto L54
L54:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v217+v243+(int32(0)-v243)&v242+v253+int32(4))))
	v258 = v257
	goto L50
}
func F_VM_ServerInfoGetFieldUnsigned(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v342 int64
	_ = v342
	var v348 int32
	_ = v348
	var v350 int64
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v364 int64
	_ = v364
	var v369 int64
	_ = v369
	var v373 int32
	_ = v373
	var v375 int64
	_ = v375
	var v377 int32
	_ = v377
	var v385 int64
	_ = v385
	var v396 int64
	_ = v396
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v422 int64
	_ = v422
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v442 int64
	_ = v442
	var v449 int32
	_ = v449
	var v461 int64
	_ = v461
	var v470 int64
	_ = v470
	var v473 int64
	_ = v473
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
	return v473
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
		v473 = v264
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
	v473 = v264
	goto L1
L59:
	;
	v292 = v9 + int32(8)
	v300 = m.G0
	v302 = v300 - int32(16)
	m.G0 = v302
	if base.Ui32(v290+int32(-21)) < base.Ui32(int32(-20)) {
		goto L69
	} else {
		goto L70
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
	if v449 != 0 {
		goto L65
	} else {
		goto L93
	}
L67:
	;
	m.G0 = v302 + int32(16)
	goto L66
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v292))) = v442
	v449 = int32(1)
	goto L67
L69:
	;
	v413 = int32(0)
	v414 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v414))) = v413
	*(*int32)(unsafe.Add(mBase, uint32(v302)+12)) = v413
	v422 = F_strtoull(m, v270, v302+int32(12), int32(10))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v292))) = v422
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v414)))
	if v424 == int32(28) {
		v449 = v413
		goto L67
	} else {
		goto L90
	}
L70:
	;
	v308 = int32(1)
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if v290 != v308 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	if v309&int32(255) != int32(45) {
		v329 = v308
		v330 = v309
		v331 = v270
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v313 = v309 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v313&int32(255)) {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v442 = base.I64_extend_i32_u(v313) & int64(255)
	goto L68
L74:
	;
	if base.Ui32(int32(8)) < base.Ui32((v330+int32(-49))&int32(255)) {
		goto L69
	} else {
		goto L76
	}
L75:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+1)))
	v329 = int32(2)
	v330 = v327
	v331 = v270 + int32(1)
	goto L74
L76:
	;
	v342 = base.I64_extend_i32_u(v330+int32(-48)) & int64(255)
	if base.Ui32(v290) <= base.Ui32(v329) {
		v385 = v342
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if v309&int32(255) != int32(45) {
		goto L85
	} else {
		goto L86
	}
L78:
	;
	v348 = v329
	v350 = v342
	v352 = v331
	goto L79
L79:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+1)))
	if base.Ui32((v354+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L69
	} else {
		goto L81
	}
L80:
	;
	v385 = v375
	goto L77
L81:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v350) {
		goto L69
	} else {
		goto L82
	}
L82:
	;
	v364 = v350 * int64(10)
	v369 = base.I64_extend_i32_u(v354+int32(-48)) & int64(255)
	if base.Ui64(v369^int64(-1)) < base.Ui64(v364) {
		goto L69
	} else {
		goto L83
	}
L83:
	;
	v373 = int32(1)
	v375 = v364 + v369
	v377 = v348 + v373
	if v377 != v290 {
		v348 = v377
		v350 = v375
		v352 = v352 + v373
		goto L79
	} else {
		goto L84
	}
L84:
	;
	goto L80
L85:
	;
	if int64(0) <= v385 {
		v442 = v385
		goto L68
	} else {
		goto L89
	}
L86:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v385) {
		goto L69
	} else {
		goto L87
	}
L87:
	;
	v396 = int64(-1)
	if v396 < v385+v396 {
		v449 = int32(0)
		goto L67
	} else {
		goto L88
	}
L88:
	;
	v442 = int64(0)
	goto L68
L89:
	;
	goto L69
L90:
	;
	if v424 == int32(68) {
		v449 = v413
		goto L67
	} else {
		goto L91
	}
L91:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if v429 == int32(0) {
		v449 = v413
		goto L67
	} else {
		goto L92
	}
L92:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v302)+12))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	v449 = base.B2i32(v433 == int32(0))
	goto L67
L93:
	;
	v461 = int64(0)
	if l2 == int32(0) {
		v473 = v461
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
	v473 = v461
	goto L1
L95:
	;
	v470 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	v473 = v470
	goto L1
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L95
}
func F_VM_SetClientNameById(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	v11 = int64(56)
	v13 = int64(65280)
	v15 = int64(40)
	v18 = int64(16711680)
	v20 = int64(24)
	v22 = int64(4278190080)
	v24 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = l0<<(uint(v11)%64) | l0&v13<<(uint(v15)%64) | (l0&v18<<(uint(v20)%64) | l0&v22<<(uint(v24)%64)) | (int64(base.Ui64(l0)>>(uint(v24)%64))&v22 | int64(base.Ui64(l0)>>(uint(v20)%64))&v18 | (int64(base.Ui64(l0)>>(uint(v15)%64))&v13 | int64(base.Ui64(l0)>>(uint(v11)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SetClientNameById[0]))
	v51 = int32(8)
	v56 = F_raxFind(m, v50, v9+v51, v51, v9+int32(4))
	mBase = m.M
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	m.G0 = v9 + v8
	if v57 != 0 {
		v62 = int32(0)
		v64 = F_clientSetName(m, v57, l1, v62)
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return int32(0)
		} else {
			if v64 != int32(-1) {
				v75 = v62
			} else {
				v71 = int32(28)
				*(*int32)(unsafe.Add(mBase, _c_F_VM_SetClientNameById[1])) = v71
				v75 = int32(1)
			}
			return v75
		}
	} else {
		v71 = int32(44)
		*(*int32)(unsafe.Add(mBase, _c_F_VM_SetClientNameById[1])) = v71
		v75 = int32(1)
		return v75
	}
}
func F_VM_SetClusterFlags(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	if l1&int64(2) == int64(0) {
	} else {
		v7 = int32(_a_F_VM_SetClusterFlags_0)
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SetClusterFlags[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_VM_SetClusterFlags[0])) = v9 | int32(2)
	}
	if l1&int64(4) == int64(0) {
	} else {
		v17 = int32(_a_F_VM_SetClusterFlags_0)
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_VM_SetClusterFlags[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_VM_SetClusterFlags[0])) = v19 | int32(4)
	}
	return
}
func F_VM_SetContextUser(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l1
	return
}
func F_VM_SetLFU(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	v4 = int32(1)
	if l0 == int32(0) {
		v18 = v4
		return v18
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7 == int32(0) {
			v18 = v4
			return v18
		} else {
			v11 = F_objectSetLRUOrLFU(m, v7, l1, int64(-1))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v18 = base.B2i32(v11 == int32(0))
				return v18
			}
		}
	}
}
func F_VM_SetLRU(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	v4 = int32(1)
	if l0 == int32(0) {
		v20 = v4
		return v20
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7 == int32(0) {
			v20 = v4
			return v20
		} else {
			v13 = F_objectSetLRUOrLFU(m, v7, int64(-1), l1*int64(1000))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v20 = base.B2i32(v13 == int32(0))
				return v20
			}
		}
	}
}
func F_VM_SetModuleUserACLString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 == int32(0) {
		F__serverAssert(m, int32(_a_F_VM_SetModuleUserACLString_0), int32(_a_F_VM_SetModuleUserACLString_1), int32(10382))
		mBase = m.M
		v90 = m.ExcPending
		if v90 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v15 = F_sdssplitargs(m, l2, v9+int32(12))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(0)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			v23 = F_ACLStringSetUser(m, v20, v19, v15, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				F_sdsfreesplitres(m, v15, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					if v23 == int32(0) {
						v81 = v19
						m.G0 = v9 + int32(16)
						return v81
					} else {
						if l3 == int32(0) {
							F_sdsfree(m, v23)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v81 = int32(1)
								m.G0 = v9 + int32(16)
								return v81
							}
						} else {
							v33 = F_createObject(m, int32(0), v23)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v33
								if l0 == int32(0) {
									v81 = int32(1)
									m.G0 = v9 + int32(16)
									return v81
								} else {
									v38 = int32(1)
									v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
									if v39&v38 == int32(0) {
										v81 = v38
										m.G0 = v9 + int32(16)
										return v81
									} else {
										v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v44 == v45 {
											v48 = int32(8)
											if v48 < v44 {
												v51 = v44
											} else {
												v51 = v48
											}
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v51 << (uint(int32(1)) % 32)
											v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											v58 = F_valkey_realloc(m, v55, v51<<(uint(int32(4))%32))
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v58
												v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
												v62 = v61
												v63 = v58
												v66 = v63 + v62<<(uint(int32(3))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v66))) = v33
												v68 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v68
												*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v62 + v68
												v81 = v68
												m.G0 = v9 + int32(16)
												return v81
											}
										} else {
											v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											v62 = v44
											v63 = v47
											v66 = v63 + v62<<(uint(int32(3))%32)
											*(*int32)(unsafe.Add(mBase, uint32(v66))) = v33
											v68 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v68
											*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v62 + v68
											v81 = v68
											m.G0 = v9 + int32(16)
											return v81
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
func F_VM_SignalKeyAsReady(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+96))
	F_signalKeyAsReady(m, v4, l1, int32(5))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_VM_StopTimer(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
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
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = l1
	v12 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_VM_StopTimer[0]))
	v15 = int32(8)
	v16 = v9 + v15
	v19 = v9 + int32(4)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	goto L5
L1:
	;
	m.G0 = v9 + int32(16)
	return v238
L2:
	;
	if v213 == int32(0) {
		v238 = v12
		goto L1
	} else {
		goto L40
	}
L3:
	;
	if v172 != v15 {
		v213 = v4
		goto L30
	} else {
		goto L31
	}
L4:
	;
	v163 = int32(0)
	v169 = v28
	v170 = v29
	v172 = v163
	v176 = v163
	goto L3
L5:
	;
	if base.Ui32(v29) < base.Ui32(int32(8)) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v40 = v28
	v41 = v29
	v43 = int32(0)
	goto L8
L7:
	;
	v169 = v153
	v170 = v154
	v172 = v156
	v176 = base.B2i32(v159 != int32(0))
	goto L3
L8:
	;
	v49 = int32(base.Ui32(v41) >> (uint(int32(3)) % 32))
	v50 = int32(4)
	v51 = v40 + v50
	if v41&v50 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v153 = v144
	v154 = v145
	v156 = v129
	v159 = v134
	goto L7
L10:
	;
	v134 = int32(0)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v51+v49+(v134-v49)&int32(3)+v122<<(uint(int32(2))%32))))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	if base.Ui32(v145) < base.Ui32(int32(8)) {
		v153 = v144
		v154 = v145
		v156 = v129
		v159 = v134
		goto L7
	} else {
		goto L28
	}
L11:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v43))))
	v100 = int32(0)
	goto L22
L12:
	;
	v56 = int32(0)
	if base.Ui32(v15) <= base.Ui32(v43) {
		v89 = v43
		v92 = v56
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v92 == v49 {
		v122 = v56
		v129 = v89
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v66 = v43
	v69 = v56
	goto L15
L15:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v69))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v66))))
	if v72 != v74 {
		v89 = v66
		v92 = v69
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v89 = v77
	v92 = v79
	goto L13
L17:
	;
	v76 = int32(1)
	v77 = v66 + v76
	v79 = v69 + v76
	if base.Ui32(v49) <= base.Ui32(v79) {
		v89 = v77
		v92 = v79
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if base.Ui32(v77) < base.Ui32(v15) {
		v66 = v77
		v69 = v79
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v153 = v40
	v154 = v41
	v156 = v89
	v159 = v92
	goto L7
L21:
	;
	if v100 != v49 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v100))))
	if v113 == v97&int32(255) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v115 = int32(1)
	v117 = v100 + v115
	if v117 != v49 {
		v100 = v117
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v169 = v40
	v170 = v41
	v172 = v43
	v176 = v115
	goto L3
L26:
	;
	v122 = v100
	v129 = v43 + int32(1)
	goto L10
L27:
	;
	v153 = v40
	v154 = v41
	v156 = v43
	v159 = v49
	goto L7
L28:
	;
	if base.Ui32(v129) < base.Ui32(v15) {
		v40 = v144
		v41 = v145
		v43 = v129
		goto L8
	} else {
		goto L29
	}
L29:
	;
	goto L9
L30:
	;
	goto L2
L31:
	;
	v178 = int32(0)
	if v170&int32(1) == v178 {
		v213 = v178
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v184 = v170 & int32(4)
	if v176&base.B2i32(v184 != int32(0)) != 0 {
		v213 = v178
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v188 = int32(1)
	if v19 == int32(0) {
		v213 = v188
		goto L30
	} else {
		goto L34
	}
L34:
	;
	if v170&int32(2) != 0 {
		v210 = int32(0)
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v210
	v213 = v188
	goto L30
L36:
	;
	v194 = int32(3)
	v195 = int32(base.Ui32(v170) >> (uint(v194) % 32))
	if v184 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v205 = int32(4)
	goto L39
L38:
	;
	v205 = v195 << (uint(int32(2)) % 32)
	goto L39
L39:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v169+v195+(int32(0)-v195)&v194+v205+int32(4))))
	v210 = v209
	goto L35
L40:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v218 != v219 {
		v238 = v12
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if l2 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v225 = int32(0)
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_VM_StopTimer[0]))
	v228 = int32(8)
	v232 = F_raxRemove(m, v227, v9+v228, v228, v225)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v223
	goto L42
L44:
	;
	return int32(0)
L45:
	;
	F_valkey_free(m, v217)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v238 = v225
	goto L1
}
func F_VM_StreamAdd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v104 int32
	_ = v104
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	if l0 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v104
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v29 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	goto L9
L4:
	;
	if base.Ui32(int32(1)) < base.Ui32(l1) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v18 = int32(0)
	if base.B2i32(l3 != v18)|base.B2i32(l4 == v18) == v18 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if l1 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if l2 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	goto L3
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamAdd[0])) = int32(28)
	v104 = int32(1)
	goto L1
L10:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v41&int32(2) != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v32&int32(15) == int32(6) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamAdd[0])) = int32(138)
	v104 = int32(1)
	goto L1
L14:
	;
	if l1 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamAdd[0])) = int32(8)
	v104 = int32(1)
	goto L1
L17:
	;
	if v29 != 0 {
		v64 = v29
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	if v48 != int64(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v51 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	if v51 != int64(0) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamAdd[0])) = int32(18)
	v104 = int32(1)
	goto L1
L22:
	;
	v65 = F_objectGetVal(m, v64)
	mBase = m.M
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v65)+16))
	if v66 != int64(-1) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v59 = F_moduleCreateEmptyKey(m, l0, int32(7))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v64 = v63
	goto L22
L26:
	;
	if l1 != 0 {
		v81 = int32(0)
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v65)+24))
	if v69 != int64(-1) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamAdd[0])) = int32(22)
	v104 = int32(1)
	goto L1
L30:
	;
	v82 = int32(1)
	v87 = F_streamAppendItem(m, v65, l3, base.I64_extend_i32_s(l4), v12+int32(16), v81, v82)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L24
	} else {
		goto L33
	}
L31:
	;
	v77 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v77
	v79 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v79
	v81 = v12
	goto L30
L32:
	;
	if v29 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	if v87 != int32(-1) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if v29 != 0 {
		v104 = v82
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v91 = F_moduleDelKeyIfEmpty(m, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L24
	} else {
		goto L36
	}
L36:
	;
	v104 = v82
	goto L1
L37:
	;
	v97 = int32(0)
	if l2 == v97 {
		v104 = v97
		goto L1
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
	goto L37
L39:
	;
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v102
	v104 = v97
	goto L1
}
func F_VM_StreamIteratorDelete(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	if l0 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorDelete[0])) = int32(138)
			return int32(1)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if v11&int32(15) == int32(6) {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
				if v21&int32(2) == int32(0) {
					*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorDelete[0])) = int32(8)
					return int32(1)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v26 != 0 {
						v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
						if v33 != int64(0) {
							v45 = l0 + int32(24)
							F_streamIteratorRemoveEntry(m, v26, v45)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v52 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(40)))) = v52
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(32)))) = v52
								*(*int64)(unsafe.Add(mBase, uint32(v45))) = v52
								return int32(0)
							}
						} else {
							v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
							if v36 != int64(0) {
								v45 = l0 + int32(24)
								F_streamIteratorRemoveEntry(m, v26, v45)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									v52 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(40)))) = v52
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(32)))) = v52
									*(*int64)(unsafe.Add(mBase, uint32(v45))) = v52
									return int32(0)
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorDelete[0])) = int32(44)
								return int32(1)
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorDelete[0])) = int32(8)
						return int32(1)
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorDelete[0])) = int32(138)
				return int32(1)
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorDelete[0])) = int32(28)
		return int32(1)
	}
}
func F_VM_StreamIteratorNextField(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v28 int32
	_ = v28
	var v33 int64
	_ = v33
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v144 int64
	_ = v144
	var v149 int32
	_ = v149
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if l0 != 0 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v16 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorNextField[0])) = int32(138)
			v149 = int32(1)
			m.G0 = v10 + int32(32)
			return v149
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v19&int32(15) == int32(6) {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v28 != 0 {
					v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
					if int64(0) < v33 {
						F_streamIteratorGetField(m, v28, v10+int32(28), v10+int32(24), v10+int32(16), v10+int32(8))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							if l1 == int32(0) {
								if l2 == int32(0) {
									v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
									v149 = int32(0)
									m.G0 = v10 + int32(32)
									return v149
								} else {
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
									v102 = F_createRawStringObject(m, v100, v101)
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v102
										v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+28)))
										if v106&int32(1) == int32(0) {
											v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
											v149 = int32(0)
											m.G0 = v10 + int32(32)
											return v149
										} else {
											v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
											v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
											if v111 == v112 {
												v115 = int32(8)
												if v115 < v111 {
													v118 = v111
												} else {
													v118 = v115
												}
												*(*int32)(unsafe.Add(mBase, uint32(v105)+20)) = v118 << (uint(int32(1)) % 32)
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
												v125 = F_valkey_realloc(m, v122, v118<<(uint(int32(4))%32))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v105)+16)) = v125
													v128 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
													v129 = v128
													v130 = v125
													v133 = v130 + v129<<(uint(int32(3))%32)
													*(*int32)(unsafe.Add(mBase, uint32(v133))) = v102
													v135 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v135
													*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v129 + v135
													v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
													v149 = int32(0)
													m.G0 = v10 + int32(32)
													return v149
												}
											} else {
												v114 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
												v129 = v111
												v130 = v114
												v133 = v130 + v129<<(uint(int32(3))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v133))) = v102
												v135 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v135
												*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v129 + v135
												v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
												v149 = int32(0)
												m.G0 = v10 + int32(32)
												return v149
											}
										}
									}
								}
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
								v56 = F_createRawStringObject(m, v54, v55)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+28)))
									if v60&int32(1) == int32(0) {
										if l2 == int32(0) {
											v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
											v149 = int32(0)
											m.G0 = v10 + int32(32)
											return v149
										} else {
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
											v102 = F_createRawStringObject(m, v100, v101)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v102
												v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+28)))
												if v106&int32(1) == int32(0) {
													v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
													v149 = int32(0)
													m.G0 = v10 + int32(32)
													return v149
												} else {
													v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
													v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
													if v111 == v112 {
														v115 = int32(8)
														if v115 < v111 {
															v118 = v111
														} else {
															v118 = v115
														}
														*(*int32)(unsafe.Add(mBase, uint32(v105)+20)) = v118 << (uint(int32(1)) % 32)
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
														v125 = F_valkey_realloc(m, v122, v118<<(uint(int32(4))%32))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v105)+16)) = v125
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
															v129 = v128
															v130 = v125
															v133 = v130 + v129<<(uint(int32(3))%32)
															*(*int32)(unsafe.Add(mBase, uint32(v133))) = v102
															v135 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v135
															*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v129 + v135
															v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
															*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
															v149 = int32(0)
															m.G0 = v10 + int32(32)
															return v149
														}
													} else {
														v114 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
														v129 = v111
														v130 = v114
														v133 = v130 + v129<<(uint(int32(3))%32)
														*(*int32)(unsafe.Add(mBase, uint32(v133))) = v102
														v135 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v135
														*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v129 + v135
														v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
														v149 = int32(0)
														m.G0 = v10 + int32(32)
														return v149
													}
												}
											}
										}
									} else {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
										if v65 == v66 {
											v69 = int32(8)
											if v69 < v65 {
												v72 = v65
											} else {
												v72 = v69
											}
											*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v72 << (uint(int32(1)) % 32)
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
											v79 = F_valkey_realloc(m, v76, v72<<(uint(int32(4))%32))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v79
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
												v83 = v82
												v84 = v79
												v87 = v84 + v83<<(uint(int32(3))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v87))) = v56
												v89 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v89
												*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = v83 + v89
												if l2 == int32(0) {
													v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
													v149 = int32(0)
													m.G0 = v10 + int32(32)
													return v149
												} else {
													v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
													v101 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
													v102 = F_createRawStringObject(m, v100, v101)
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l2))) = v102
														v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+28)))
														if v106&int32(1) == int32(0) {
															v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
															*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
															v149 = int32(0)
															m.G0 = v10 + int32(32)
															return v149
														} else {
															v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
															v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
															if v111 == v112 {
																v115 = int32(8)
																if v115 < v111 {
																	v118 = v111
																} else {
																	v118 = v115
																}
																*(*int32)(unsafe.Add(mBase, uint32(v105)+20)) = v118 << (uint(int32(1)) % 32)
																v122 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
																v125 = F_valkey_realloc(m, v122, v118<<(uint(int32(4))%32))
																mBase = m.M
																v126 = m.ExcPending
																if v126 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+16)) = v125
																	v128 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
																	v129 = v128
																	v130 = v125
																	v133 = v130 + v129<<(uint(int32(3))%32)
																	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v102
																	v135 = int32(1)
																	*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v135
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v129 + v135
																	v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
																	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
																	v149 = int32(0)
																	m.G0 = v10 + int32(32)
																	return v149
																}
															} else {
																v114 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
																v129 = v111
																v130 = v114
																v133 = v130 + v129<<(uint(int32(3))%32)
																*(*int32)(unsafe.Add(mBase, uint32(v133))) = v102
																v135 = int32(1)
																*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v135
																*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v129 + v135
																v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
																*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
																v149 = int32(0)
																m.G0 = v10 + int32(32)
																return v149
															}
														}
													}
												}
											}
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
											v83 = v65
											v84 = v68
											v87 = v84 + v83<<(uint(int32(3))%32)
											*(*int32)(unsafe.Add(mBase, uint32(v87))) = v56
											v89 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v89
											*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = v83 + v89
											if l2 == int32(0) {
												v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
												v149 = int32(0)
												m.G0 = v10 + int32(32)
												return v149
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
												v102 = F_createRawStringObject(m, v100, v101)
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = v102
													v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+28)))
													if v106&int32(1) == int32(0) {
														v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
														v149 = int32(0)
														m.G0 = v10 + int32(32)
														return v149
													} else {
														v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
														v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
														if v111 == v112 {
															v115 = int32(8)
															if v115 < v111 {
																v118 = v111
															} else {
																v118 = v115
															}
															*(*int32)(unsafe.Add(mBase, uint32(v105)+20)) = v118 << (uint(int32(1)) % 32)
															v122 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
															v125 = F_valkey_realloc(m, v122, v118<<(uint(int32(4))%32))
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v105)+16)) = v125
																v128 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
																v129 = v128
																v130 = v125
																v133 = v130 + v129<<(uint(int32(3))%32)
																*(*int32)(unsafe.Add(mBase, uint32(v133))) = v102
																v135 = int32(1)
																*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v135
																*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v129 + v135
																v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
																*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
																v149 = int32(0)
																m.G0 = v10 + int32(32)
																return v149
															}
														} else {
															v114 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
															v129 = v111
															v130 = v114
															v133 = v130 + v129<<(uint(int32(3))%32)
															*(*int32)(unsafe.Add(mBase, uint32(v133))) = v102
															v135 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v135
															*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v129 + v135
															v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
															*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v144 + int64(-1)
															v149 = int32(0)
															m.G0 = v10 + int32(32)
															return v149
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
						*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorNextField[0])) = int32(44)
						v149 = int32(1)
						m.G0 = v10 + int32(32)
						return v149
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorNextField[0])) = int32(8)
					v149 = int32(1)
					m.G0 = v10 + int32(32)
					return v149
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorNextField[0])) = int32(138)
				v149 = int32(1)
				m.G0 = v10 + int32(32)
				return v149
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorNextField[0])) = int32(28)
		v149 = int32(1)
		m.G0 = v10 + int32(32)
		return v149
	}
}
func F_VM_StreamIteratorNextID(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v48 int64
	_ = v48
	var v65 int32
	_ = v65
	if l0 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v10 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorNextID[0])) = int32(138)
			v65 = int32(1)
			return v65
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v13&int32(15) == int32(6) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v21 != 0 {
					v26 = l0 + int32(24)
					v28 = l0 + int32(40)
					v29 = F_streamIteratorGetID(m, v21, v26, v28)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v29 == int32(0) {
							v48 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v26))) = v48
							*(*int64)(unsafe.Add(mBase, uint32(l0+int32(40)))) = v48
							*(*int64)(unsafe.Add(mBase, uint32(l0+int32(32)))) = v48
							*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorNextID[0])) = int32(44)
							v65 = int32(1)
							return v65
						} else {
							if l1 == int32(0) {
							} else {
								v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
								*(*int64)(unsafe.Add(mBase, uint32(l1))) = v37
								v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
								*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v39
							}
							v41 = int32(0)
							if l2 == v41 {
								v65 = v41
								return v65
							} else {
								v44 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
								*(*uint32)(unsafe.Add(mBase, uint32(l2))) = uint32(v44)
								return int32(0)
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorNextID[0])) = int32(8)
					v65 = int32(1)
					return v65
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorNextID[0])) = int32(138)
				v65 = int32(1)
				return v65
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamIteratorNextID[0])) = int32(28)
		v65 = int32(1)
		return v65
	}
}
func F_VM_StreamTrimByID(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v38 int64
	_ = v38
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamTrimByID[0])) = int32(28)
		v59 = int64(-1)
		m.G0 = v9 + int32(32)
		return v59
	} else {
		if base.Ui32(int32(1)) < base.Ui32(l1) {
			*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamTrimByID[0])) = int32(28)
			v59 = int64(-1)
			m.G0 = v9 + int32(32)
			return v59
		} else {
			if l2 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v19 == int32(0) {
					*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamTrimByID[0])) = int32(138)
					v59 = int64(-1)
					m.G0 = v9 + int32(32)
					return v59
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					if v22&int32(15) == int32(6) {
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v31&int32(2) != 0 {
							v38 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v38
							v43 = v9 + int32(24)
							v44 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v43))) = v44
							v46 = F_objectGetVal(m, v19)
							mBase = m.M
							v49 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
							*(*int64)(unsafe.Add(mBase, uint32(v9+int32(8)))) = v49
							v51 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v9))) = v51
							v53 = F_streamTrimByID(m, v46, v9, l1)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								v59 = v53
								m.G0 = v9 + int32(32)
								return v59
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamTrimByID[0])) = int32(8)
							v59 = int64(-1)
							m.G0 = v9 + int32(32)
							return v59
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamTrimByID[0])) = int32(138)
						v59 = int64(-1)
						m.G0 = v9 + int32(32)
						return v59
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_VM_StreamTrimByID[0])) = int32(28)
				v59 = int64(-1)
				m.G0 = v9 + int32(32)
				return v59
			}
		}
	}
}
func F_VM_StringDMA(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		if v13&int32(15) != 0 {
			v58 = int32(0)
			return v58
		} else {
			if v13&int32(240)|l2&int32(2) == int32(0) {
				v30 = v7
				v32 = F_objectGetVal(m, v30)
				mBase = m.M
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-1)))))
				switch v35 & int32(7) {
				case 0:
					v52 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
				case 1:
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-3)))))
					v52 = v42
				case 2:
					v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+int32(-5)))))
					v52 = v45
				case 3:
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(-9))))
					v52 = v48
				case 4:
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(-17))))
					v52 = v51
				default:
					v52 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v52
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v55 = F_objectGetVal(m, v54)
				mBase = m.M
				v58 = v55
				return v58
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v25 = F_dbUnshareStringValue(m, v23, v24, v7)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v25
					v30 = v25
					v32 = F_objectGetVal(m, v30)
					mBase = m.M
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-1)))))
					switch v35 & int32(7) {
					case 0:
						v52 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
					case 1:
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-3)))))
						v52 = v42
					case 2:
						v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+int32(-5)))))
						v52 = v45
					case 3:
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(-9))))
						v52 = v48
					case 4:
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(-17))))
						v52 = v51
					default:
						v52 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v52
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v55 = F_objectGetVal(m, v54)
					mBase = m.M
					v58 = v55
					return v58
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
		return int32(_a_F_VM_StringDMA_0)
	}
}
func F_VM_StringSet(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v10 = int32(1)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v11&int32(2) == int32(0) {
		v42 = v10
		m.G0 = v7 + int32(16)
		return v42
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v16 != 0 {
			v42 = v10
			m.G0 = v7 + int32(16)
			return v42
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v17 == int32(0) {
				F_incrRefCount(m, l1)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_setKey(m, v31, v32, v33, v7+int32(12), int32(10))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v39
						v42 = int32(0)
						m.G0 = v7 + int32(16)
						return v42
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
					F_incrRefCount(m, l1)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						F_setKey(m, v31, v32, v33, v7+int32(12), int32(10))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v39
							v42 = int32(0)
							m.G0 = v7 + int32(16)
							return v42
						}
					}
				}
			}
		}
	}
}
func F_VM_StringToDouble(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_getDoubleFromObject(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v3 != int32(0))
	}
}
func F_VM_StringToLongDouble(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v6 = F_objectGetVal(m, l0)
	mBase = m.M
	v8 = F_objectGetVal(m, l0)
	mBase = m.M
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-1)))))
	switch v11 & int32(7) {
	case 0:
		v28 = int32(base.Ui32(v11) >> (uint(int32(3)) % 32))
	case 1:
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-3)))))
		v28 = v18
	case 2:
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-5)))))
		v28 = v21
	case 3:
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-9))))
		v28 = v24
	case 4:
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-17))))
		v28 = v27
	default:
		v28 = int32(0)
	}
	v29 = F_string2ld(m, v6, v28, l1)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v29 == int32(0))
	}
}
func F_VM_StringToLongLong(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v77 int64
	_ = v77
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int64
	_ = v99
	var v104 int64
	_ = v104
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	var v112 int32
	_ = v112
	var v120 int64
	_ = v120
	var v144 int64
	_ = v144
	var v163 int32
	_ = v163
	v6 = F_objectGetVal(m, l0)
	mBase = m.M
	v8 = F_objectGetVal(m, l0)
	mBase = m.M
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-1)))))
	switch v11 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v28 = int32(0)
		goto L1
	}
L1:
	;
	v29 = int32(0)
	if base.Ui32(v28+int32(-21)) < base.Ui32(int32(-20)) {
		v163 = v29
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-17))))
	v28 = v27
	goto L1
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-9))))
	v28 = v24
	goto L1
L4:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-5)))))
	v28 = v21
	goto L1
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-3)))))
	v28 = v18
	goto L1
L6:
	;
	v28 = int32(base.Ui32(v11) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	return base.B2i32(v163 == int32(0))
L8:
	;
	goto L7
L9:
	;
	v41 = int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v28 != v41 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v163 = int32(1)
	goto L8
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v144
	goto L10
L12:
	;
	if v42&int32(255) == int32(45) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v46 = v42 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v46&int32(255)) {
		v163 = v29
		goto L8
	} else {
		goto L14
	}
L14:
	;
	if l1 == int32(0) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v144 = base.I64_extend_i32_u(v46) & int64(255)
	goto L11
L16:
	;
	if base.Ui32(int32(8)) < base.Ui32((v65+int32(-49))&int32(255)) {
		v163 = v29
		goto L8
	} else {
		goto L19
	}
L17:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	v64 = int32(2)
	v65 = v62
	v66 = v6 + int32(1)
	goto L16
L18:
	;
	v64 = v41
	v65 = v42
	v66 = v6
	goto L16
L19:
	;
	v77 = base.I64_extend_i32_u(v65+int32(-48)) & int64(255)
	if base.Ui32(v28) <= base.Ui32(v64) {
		v120 = v77
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v42&int32(255) != int32(45) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v83 = v64
	v85 = v77
	v87 = v66
	goto L22
L22:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	if base.Ui32((v89+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v163 = v29
		goto L8
	} else {
		goto L24
	}
L23:
	;
	v120 = v110
	goto L20
L24:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v85) {
		v163 = v29
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v99 = v85 * int64(10)
	v104 = base.I64_extend_i32_u(v89+int32(-48)) & int64(255)
	if base.Ui64(v104^int64(-1)) < base.Ui64(v99) {
		v163 = v29
		goto L8
	} else {
		goto L26
	}
L26:
	;
	v108 = int32(1)
	v110 = v99 + v104
	v112 = v83 + v108
	if v112 != v28 {
		v83 = v112
		v85 = v110
		v87 = v87 + v108
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	if v120 < int64(0) {
		v163 = v29
		goto L8
	} else {
		goto L32
	}
L29:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v120) {
		v163 = v29
		goto L8
	} else {
		goto L30
	}
L30:
	;
	if l1 == int32(0) {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v144 = int64(0) - v120
	goto L11
L32:
	;
	if l1 == int32(0) {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	v144 = v120
	goto L11
}
func F_VM_StringToULongLong(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v78 int64
	_ = v78
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v100 int64
	_ = v100
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v121 int64
	_ = v121
	var v132 int64
	_ = v132
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int64
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int64
	_ = v178
	var v185 int32
	_ = v185
	v6 = F_objectGetVal(m, l0)
	mBase = m.M
	v8 = F_objectGetVal(m, l0)
	mBase = m.M
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-1)))))
	switch v11 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v28 = int32(0)
		goto L1
	}
L1:
	;
	v36 = m.G0
	v38 = v36 - int32(16)
	m.G0 = v38
	if base.Ui32(v28+int32(-21)) < base.Ui32(int32(-20)) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-17))))
	v28 = v27
	goto L1
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-9))))
	v28 = v24
	goto L1
L4:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-5)))))
	v28 = v21
	goto L1
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-3)))))
	v28 = v18
	goto L1
L6:
	;
	v28 = int32(base.Ui32(v11) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	return base.B2i32(v185 == int32(0))
L8:
	;
	m.G0 = v38 + int32(16)
	goto L7
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v178
	v185 = int32(1)
	goto L8
L10:
	;
	v149 = int32(0)
	v150 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v149
	v158 = F_strtoull(m, v6, v38+int32(12), int32(10))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if v160 == int32(28) {
		v185 = v149
		goto L8
	} else {
		goto L31
	}
L11:
	;
	v44 = int32(1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v28 != v44 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v45&int32(255) != int32(45) {
		v65 = v44
		v66 = v45
		v67 = v6
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v49 = v45 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v49&int32(255)) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v178 = base.I64_extend_i32_u(v49) & int64(255)
	goto L9
L15:
	;
	if base.Ui32(int32(8)) < base.Ui32((v66+int32(-49))&int32(255)) {
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	v65 = int32(2)
	v66 = v63
	v67 = v6 + int32(1)
	goto L15
L17:
	;
	v78 = base.I64_extend_i32_u(v66+int32(-48)) & int64(255)
	if base.Ui32(v28) <= base.Ui32(v65) {
		v121 = v78
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v45&int32(255) != int32(45) {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v84 = v65
	v86 = v78
	v88 = v67
	goto L20
L20:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	if base.Ui32((v90+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L10
	} else {
		goto L22
	}
L21:
	;
	v121 = v111
	goto L18
L22:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v86) {
		goto L10
	} else {
		goto L23
	}
L23:
	;
	v100 = v86 * int64(10)
	v105 = base.I64_extend_i32_u(v90+int32(-48)) & int64(255)
	if base.Ui64(v105^int64(-1)) < base.Ui64(v100) {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v109 = int32(1)
	v111 = v100 + v105
	v113 = v84 + v109
	if v113 != v28 {
		v84 = v113
		v86 = v111
		v88 = v88 + v109
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	if int64(0) <= v121 {
		v178 = v121
		goto L9
	} else {
		goto L30
	}
L27:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v121) {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v132 = int64(-1)
	if v132 < v121+v132 {
		v185 = int32(0)
		goto L8
	} else {
		goto L29
	}
L29:
	;
	v178 = int64(0)
	goto L9
L30:
	;
	goto L10
L31:
	;
	if v160 == int32(68) {
		v185 = v149
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v165 == int32(0) {
		v185 = v149
		goto L8
	} else {
		goto L33
	}
L33:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	v185 = base.B2i32(v169 == int32(0))
	goto L8
}
func F_VM_TrimStringAllocation(m *base.Module, l0 int32) {
	var v6 int32
	_ = v6
	if l0 == int32(0) {
		return
	} else {
		F_trimStringObjectIfNeeded(m, l0, int32(1))
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_VM_TryCalloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	v3 = l1 * l0
	if base.Ui32(int32(2147483646)) < base.Ui32(v3) {
		v53 = int32(0)
	} else {
		if v3 != 0 {
			v14 = v3
		} else {
			v14 = int32(4)
		}
		v16 = v14 + int32(8)
		v17 = F_emscripten_builtin_calloc(m, int32(1), v16)
		mBase = m.M
		if v17 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v14
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_VM_TryCalloc[0]))
			if v22 != int32(-1) {
				v33 = v22
			} else {
				v25 = int32(0)
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_VM_TryCalloc[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_VM_TryCalloc[0])) = v27
				*(*int32)(unsafe.Add(mBase, _c_F_VM_TryCalloc[1])) = v27 + int32(1)
				v33 = v27
			}
			if v33 < int32(260) {
				v42 = v33 << (uint(int32(2)) % 32)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_VM_TryCalloc[2])))
				*(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_VM_TryCalloc[2]))) = v45 + v16
			} else {
				v36 = int32(0)
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_VM_TryCalloc[3]))
				*(*int32)(unsafe.Add(mBase, _c_F_VM_TryCalloc[3])) = v38 + v16
			}
			v53 = v17 + int32(8)
		} else {
			v53 = int32(0)
		}
	}
	return v53
}
func F_VM_UnregisterCommandFilter(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v5 = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 != v7 {
		v34 = v5
		return v34
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnregisterCommandFilter[0]))
		v11 = F_listSearchKey(m, v10, l1)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if v11 == int32(0) {
				v34 = v5
				return v34
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_VM_UnregisterCommandFilter[0]))
				F_listDelNode(m, v18, v11)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
					v23 = F_listSearchKey(m, v22, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						if v23 == int32(0) {
							v34 = v5
							return v34
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
							F_listDelNode(m, v28, v23)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								F_valkey_free(m, l1)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v34 = int32(0)
									return v34
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_VM_UpdateRuntimeArgs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+64))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v8 < int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	F_valkey_free(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	v12 = int32(0)
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v12<<(uint(int32(2))%32))))
	F_decrRefCount(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	v27 = v12 + int32(1)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v27 < v28 {
		v12 = v27
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	v39 = l2 + int32(-1)
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v47 = F_valkey_malloc(m, v39<<(uint(int32(2))%32))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v39
	v41 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v41
	return v41
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v47
	v51 = int32(1)
	if l2 <= v51 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v54 = v51
	goto L14
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v61 = v54 << (uint(int32(2)) % 32)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1+v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v59+v61+int32(-4)))) = v66
	F_incrRefCount(m, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	goto L12
L16:
	;
	v71 = v54 + int32(1)
	if v71 != l2 {
		v54 = v71
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
}
func F_VM_WrongArity(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_addReplyErrorArity(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_VM_ZsetFirstInLexRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_zsetInitLexRange(m, l0, l1, l2, int32(1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_VM_ZsetIncrby(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v6
	v15 = int32(1)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v16&int32(2) == v6 {
		v98 = v15
		m.G0 = v11 + int32(16)
		return v98
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v21 == int32(0) {
			v29 = F_createZsetListpackObject(m)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v29
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				F_dbAdd(m, v34, v35, v11+int32(12))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v40
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					switch v42&int32(15) + int32(-3) {
					case 0:
						*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(4294967296)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
						v53 = v40
					default:
						v53 = v40
					case 3:
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
						v53 = v40
					}
					if l3 != 0 {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						v60 = int32(29)
						v72 = int32(base.Ui32(v55)>>(uint(int32(2))%32))&int32(24) | int32(base.Ui32(v55<<(uint(v60)%32)&int32(1610612736)|v55<<(uint(int32(31))%32))>>(uint(v60)%32)) | int32(1)
					} else {
						v72 = int32(1)
					}
					v73 = F_objectGetVal(m, l2)
					mBase = m.M
					v76 = F_zsetAdd(m, v53, l1, v73, v72, v11+int32(8), l4)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						if v76 != 0 {
							v85 = int32(0)
							if l3 == v85 {
								v98 = v85
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v88<<(uint(int32(4))%32)&int32(16) | v88&int32(12)
								v98 = v85
							}
							m.G0 = v11 + int32(16)
							return v98
						} else {
							if l3 == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
							}
							v82 = F_moduleDelKeyIfEmpty(m, l0)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								v98 = int32(1)
								m.G0 = v11 + int32(16)
								return v98
							}
						}
					}
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			if v24&int32(15) == int32(3) {
				v53 = v21
				if l3 != 0 {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					v60 = int32(29)
					v72 = int32(base.Ui32(v55)>>(uint(int32(2))%32))&int32(24) | int32(base.Ui32(v55<<(uint(v60)%32)&int32(1610612736)|v55<<(uint(int32(31))%32))>>(uint(v60)%32)) | int32(1)
				} else {
					v72 = int32(1)
				}
				v73 = F_objectGetVal(m, l2)
				mBase = m.M
				v76 = F_zsetAdd(m, v53, l1, v73, v72, v11+int32(8), l4)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					if v76 != 0 {
						v85 = int32(0)
						if l3 == v85 {
							v98 = v85
						} else {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v88<<(uint(int32(4))%32)&int32(16) | v88&int32(12)
							v98 = v85
						}
						m.G0 = v11 + int32(16)
						return v98
					} else {
						if l3 == int32(0) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
						}
						v82 = F_moduleDelKeyIfEmpty(m, l0)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v98 = int32(1)
							m.G0 = v11 + int32(16)
							return v98
						}
					}
				}
			} else {
				v98 = v15
				m.G0 = v11 + int32(16)
				return v98
			}
		}
	}
}
func F_VM_ZsetRangeCurrentElement(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 float64
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
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
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 == v3 {
		v117 = v3
		return v117
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		if v10&int32(15) != int32(3) {
			v117 = v3
			return v117
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			if v15 == int32(0) {
				v117 = v3
				return v117
			} else {
				switch int32(base.Ui32(v10)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
				case 0:
					if l1 == int32(0) {
					} else {
						v42 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
						*(*float64)(unsafe.Add(mBase, uint32(l1))) = v42
					}
					v46 = v15 + int32(16)
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
					v50 = v46 + v47<<(uint(int32(3))%32)
					v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
					v52 = v50 + v51
					v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(0)))))
					switch v57 & int32(7) {
					case 0:
						v74 = int32(base.Ui32(v57) >> (uint(int32(3)) % 32))
					case 1:
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-2)))))
						v74 = v64
					case 2:
						v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(-4)))))
						v74 = v67
					case 3:
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-8))))
						v74 = v70
					case 4:
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-16))))
						v74 = v73
					default:
						v74 = int32(0)
					}
					v75 = F_createStringObject_1(m, v52+int32(1), v74)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v78 = v75
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+28)))
						if v81&int32(1) == int32(0) {
							v117 = v78
							return v117
						} else {
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v80)+24))
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
							if v86 == v87 {
								v90 = int32(8)
								if v90 < v86 {
									v93 = v86
								} else {
									v93 = v90
								}
								*(*int32)(unsafe.Add(mBase, uint32(v80)+20)) = v93 << (uint(int32(1)) % 32)
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
								v100 = F_valkey_realloc(m, v97, v93<<(uint(int32(4))%32))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v100
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v80)+24))
									v104 = v100
									v105 = v103
									v108 = v104 + v105<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v108))) = v78
									v110 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v110
									*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = v105 + v110
									v117 = v78
									return v117
								}
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
								v104 = v89
								v105 = v86
								v108 = v104 + v105<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v108))) = v78
								v110 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v110
								*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = v105 + v110
								v117 = v78
								return v117
							}
						}
					}
				default:
					F__serverPanic_1(m, int32(_a_F_VM_ZsetRangeCurrentElement_0), int32(5249), int32(_a_F_VM_ZsetRangeCurrentElement_1), int32(0))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 4:
					v24 = F_lpGetObject(m, v15)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						if l1 == int32(0) {
							v38 = F_createObject(m, int32(0), v24)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v78 = v38
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+28)))
								if v81&int32(1) == int32(0) {
									v117 = v78
									return v117
								} else {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v80)+24))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
									if v86 == v87 {
										v90 = int32(8)
										if v90 < v86 {
											v93 = v86
										} else {
											v93 = v90
										}
										*(*int32)(unsafe.Add(mBase, uint32(v80)+20)) = v93 << (uint(int32(1)) % 32)
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
										v100 = F_valkey_realloc(m, v97, v93<<(uint(int32(4))%32))
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v100
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v80)+24))
											v104 = v100
											v105 = v103
											v108 = v104 + v105<<(uint(int32(3))%32)
											*(*int32)(unsafe.Add(mBase, uint32(v108))) = v78
											v110 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v110
											*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = v105 + v110
											v117 = v78
											return v117
										}
									} else {
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
										v104 = v89
										v105 = v86
										v108 = v104 + v105<<(uint(int32(3))%32)
										*(*int32)(unsafe.Add(mBase, uint32(v108))) = v78
										v110 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v110
										*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = v105 + v110
										v117 = v78
										return v117
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v31 = F_objectGetVal(m, v30)
							mBase = m.M
							v32 = F_lpNext(m, v31, v15)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = F_zzlGetScore(m, v32)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(l1))) = v34
									v38 = F_createObject(m, int32(0), v24)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										v78 = v38
										v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+28)))
										if v81&int32(1) == int32(0) {
											v117 = v78
											return v117
										} else {
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v80)+24))
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
											if v86 == v87 {
												v90 = int32(8)
												if v90 < v86 {
													v93 = v86
												} else {
													v93 = v90
												}
												*(*int32)(unsafe.Add(mBase, uint32(v80)+20)) = v93 << (uint(int32(1)) % 32)
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
												v100 = F_valkey_realloc(m, v97, v93<<(uint(int32(4))%32))
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v100
													v103 = *(*int32)(unsafe.Add(mBase, uint32(v80)+24))
													v104 = v100
													v105 = v103
													v108 = v104 + v105<<(uint(int32(3))%32)
													*(*int32)(unsafe.Add(mBase, uint32(v108))) = v78
													v110 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = v105 + v110
													v117 = v78
													return v117
												}
											} else {
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
												v104 = v89
												v105 = v86
												v108 = v104 + v105<<(uint(int32(3))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v108))) = v78
												v110 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = v105 + v110
												v117 = v78
												return v117
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
func F_VM_ZsetRangePrev(m *base.Module, l0 int32) int32 {
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
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 == v2 {
		v166 = v2
		return v166
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		if v11&int32(15) != int32(3) {
			v166 = v2
			return v166
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v16 == int32(0) {
				v166 = v2
				return v166
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
				if v19 == int32(0) {
					v166 = v2
					return v166
				} else {
					switch int32(base.Ui32(v11)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
					case 0:
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
						if v70 != 0 {
							if v16 != int32(2) {
								v89 = v16
								if v89 != int32(1) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v70
									v166 = int32(1)
									return v166
								} else {
									v93 = v70 + int32(16)
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
									v97 = v93 + v94<<(uint(int32(3))%32)
									v98 = int32(*(*int8)(unsafe.Add(mBase, uint32(v97))))
									v101 = v97 + v98 + int32(1)
									v103 = l0 + int32(56)
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
									if v108 == int32(0) {
										if v101 != v107 {
											v125 = int32(0)
											v127 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangePrev[0]))
											if v101 == v127 {
												v142 = v125
												v148 = v142
											} else {
												v130 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangePrev[1]))
												if v107 == v130 {
													v142 = v125
													v148 = v142
												} else {
													if v107 != v127 {
														if v101 == v130 {
															v142 = int32(1)
														} else {
															v137 = int32(-1)
															v140 = F_sdscmp(m, v101, v107)
															mBase = m.M
															v142 = base.B2i32(v137 < v140)
														}
														v148 = v142
													} else {
														v148 = int32(1)
													}
												}
											}
										} else {
											v148 = int32(1)
										}
									} else {
										v111 = int32(0)
										if v101 == v107 {
											v142 = v111
										} else {
											v114 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangePrev[0]))
											if v101 == v114 {
												v142 = v111
											} else {
												v117 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangePrev[1]))
												if v107 == v117 {
													v142 = v111
												} else {
													v119 = int32(1)
													if v107 == v114 {
														v142 = v119
													} else {
														if v101 == v117 {
															v142 = v119
														} else {
															v137 = int32(0)
															v140 = F_sdscmp(m, v101, v107)
															mBase = m.M
															v142 = base.B2i32(v137 < v140)
														}
													}
												}
											}
										}
										v148 = v142
									}
									if v148 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v70
										v166 = int32(1)
										return v166
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
										return int32(0)
									}
								}
							} else {
								v77 = *(*float64)(unsafe.Add(mBase, uint32(v70)))
								v79 = l0 + int32(32)
								v81 = *(*float64)(unsafe.Add(mBase, uint32(v79)))
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
								if v84 != 0 {
									v85 = base.F64_gt(v77, v81)
								} else {
									v85 = base.F64_ge(v77, v81)
								}
								if v85 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
									return int32(0)
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v89 = v88
									if v89 != int32(1) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v70
										v166 = int32(1)
										return v166
									} else {
										v93 = v70 + int32(16)
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
										v97 = v93 + v94<<(uint(int32(3))%32)
										v98 = int32(*(*int8)(unsafe.Add(mBase, uint32(v97))))
										v101 = v97 + v98 + int32(1)
										v103 = l0 + int32(56)
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
										if v108 == int32(0) {
											if v101 != v107 {
												v125 = int32(0)
												v127 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangePrev[0]))
												if v101 == v127 {
													v142 = v125
													v148 = v142
												} else {
													v130 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangePrev[1]))
													if v107 == v130 {
														v142 = v125
														v148 = v142
													} else {
														if v107 != v127 {
															if v101 == v130 {
																v142 = int32(1)
															} else {
																v137 = int32(-1)
																v140 = F_sdscmp(m, v101, v107)
																mBase = m.M
																v142 = base.B2i32(v137 < v140)
															}
															v148 = v142
														} else {
															v148 = int32(1)
														}
													}
												}
											} else {
												v148 = int32(1)
											}
										} else {
											v111 = int32(0)
											if v101 == v107 {
												v142 = v111
											} else {
												v114 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangePrev[0]))
												if v101 == v114 {
													v142 = v111
												} else {
													v117 = *(*int32)(unsafe.Add(mBase, _c_F_VM_ZsetRangePrev[1]))
													if v107 == v117 {
														v142 = v111
													} else {
														v119 = int32(1)
														if v107 == v114 {
															v142 = v119
														} else {
															if v101 == v117 {
																v142 = v119
															} else {
																v137 = int32(0)
																v140 = F_sdscmp(m, v101, v107)
																mBase = m.M
																v142 = base.B2i32(v137 < v140)
															}
														}
													}
												}
											}
											v148 = v142
										}
										if v148 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v70
											v166 = int32(1)
											return v166
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
						F__serverPanic_1(m, int32(_a_F_VM_ZsetRangePrev_0), int32(5375), int32(_a_F_VM_ZsetRangePrev_1), int32(0))
						mBase = m.M
						v159 = m.ExcPending
						if v159 != 0 {
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
						v30 = F_lpPrev(m, v28, v29)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							if v30 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
								return int32(0)
							} else {
								v36 = F_lpPrev(m, v28, v30)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									if v36 != 0 {
										v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										switch v43 + int32(-1) {
										case 0:
											v64 = F_zzlLexValueGteMin(m, v36, l0+int32(56))
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												if v64 != 0 {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v36
													v166 = int32(1)
													return v166
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
													v53 = *(*float64)(unsafe.Add(mBase, uint32(v51)))
													v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
													if v56 != 0 {
														v57 = base.F64_gt(v48, v53)
													} else {
														v57 = base.F64_ge(v48, v53)
													}
													if v57 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v36
														v166 = int32(1)
														return v166
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1)
														return int32(0)
													}
												}
											}
										default:
											*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v36
											v166 = int32(1)
											return v166
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
