package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_isValidActiveDefrag(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	if l0 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a_F_isValidActiveDefrag_0)
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_isValidAuxString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = int32(0)
	goto L4
L2:
	;
	return int32(1)
L3:
	;
	return v131
L4:
	;
	v14 = int32(0)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v11))))
	if base.Ui32(v16) < base.Ui32(int32(45)) {
		v131 = v14
		goto L3
	} else {
		goto L6
	}
L5:
	;
	v131 = v126
	goto L3
L6:
	;
	if v16 == int32(127) {
		v131 = v14
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L12
L8:
	;
	if v125 != 0 {
		v131 = v14
		goto L3
	} else {
		goto L33
	}
L9:
	;
	v125 = int32(0)
	goto L8
L10:
	;
	v103 = v96
	v105 = v98
	goto L28
L11:
	;
	if base.B2i32(v43 != v44) == int32(0) {
		goto L9
	} else {
		goto L19
	}
L12:
	;
	goto L13
L13:
	;
	v35 = int32(_a_F_isValidAuxString_0)
	v37 = int32(15)
	goto L14
L14:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v40 == v16&int32(255) {
		v96 = v35
		v98 = v37
		goto L10
	} else {
		goto L16
	}
L15:
	;
	goto L11
L16:
	;
	v43 = v37 + int32(-1)
	v44 = int32(0)
	v47 = v35 + int32(1)
	if v47&int32(3) == v44 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	if v43 != 0 {
		v35 = v47
		v37 = v43
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v59 == v16&int32(255) {
		v89 = v47
		v91 = v43
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v91 == int32(0) {
		goto L9
	} else {
		goto L27
	}
L21:
	;
	if base.Ui32(v43) < base.Ui32(int32(4)) {
		v89 = v47
		v91 = v43
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v69 = v47
	v71 = v43
	goto L23
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v76 = v75 ^ v16&int32(255)*int32(16843009)
	v79 = int32(-2139062144)
	if (int32(16843008)-v76|v76)&v79 != v79 {
		v96 = v69
		v98 = v71
		goto L10
	} else {
		goto L25
	}
L24:
	;
	v89 = v84
	v91 = v86
	goto L20
L25:
	;
	v84 = v69 + int32(4)
	v86 = v71 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v86) {
		v69 = v84
		v71 = v86
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v96 = v89
	v98 = v91
	goto L10
L28:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v108 != v16&int32(255) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L9
L30:
	;
	v113 = v105 + int32(-1)
	if v113 != 0 {
		v103 = v103 + int32(1)
		v105 = v113
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v125 = v103
	goto L8
L32:
	;
	goto L29
L33:
	;
	v126 = int32(1)
	v128 = v11 + v126
	if v128 != l1 {
		v11 = v128
		goto L4
	} else {
		goto L34
	}
L34:
	;
	goto L5
}
func F_isValidDbHashSeed(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	v4 = int32(1)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v7&int32(7) + int32(-2) {
	case 0:
		v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v21 = v14
		if base.Ui32(v21) < base.Ui32(int32(257)) {
			v28 = v4
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a_F_isValidDbHashSeed_0)
			v28 = int32(0)
		}
	case 1:
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v21 = v17
		if base.Ui32(v21) < base.Ui32(int32(257)) {
			v28 = v4
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a_F_isValidDbHashSeed_0)
			v28 = int32(0)
		}
	case 2:
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v21 = v20
		if base.Ui32(v21) < base.Ui32(int32(257)) {
			v28 = v4
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a_F_isValidDbHashSeed_0)
			v28 = int32(0)
		}
	default:
		v28 = v4
	}
	return v28
}
func F_isValidIpV6(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = int32(1)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 == int32(0) {
		v20 = v9
	} else {
		v16 = F_inet_pton(m, int32(10), l0, v7+int32(12))
		mBase = m.M
		if v16 != 0 {
			v20 = v9
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a_F_isValidIpV6_0)
			v20 = int32(0)
		}
	}
	m.G0 = v7 + int32(32)
	return v20
}
func F_isValidMptcp(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(1)
}
