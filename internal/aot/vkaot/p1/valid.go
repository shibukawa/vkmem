package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_isValidAnnouncedNodename(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v8 & int32(7) {
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
		v25 = int32(0)
		goto L1
	}
L1:
	;
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v25 = v24
	goto L1
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v25 = v21
	goto L1
L4:
	;
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v25 = v18
	goto L1
L5:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v25 = v15
	goto L1
L6:
	;
	v25 = int32(base.Ui32(v8) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	return v62
L8:
	;
	if v58 != 0 {
		v62 = int32(1)
		goto L7
	} else {
		goto L18
	}
L9:
	;
	v34 = int32(0)
	goto L12
L10:
	;
	v58 = int32(1)
	goto L8
L11:
	;
	v58 = v52
	goto L8
L12:
	;
	v37 = int32(0)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v34))))
	if base.Ui32(v39) < base.Ui32(int32(45)) {
		v52 = v37
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v52 = v47
	goto L11
L14:
	;
	if v39 == int32(127) {
		v52 = v37
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v46 = F_memchr(m, int32(_a515), v39, int32(15))
	mBase = m.M
	if v46 != 0 {
		v52 = v37
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v47 = int32(1)
	v49 = v34 + v47
	if v49 != v25 {
		v34 = v49
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a516)
	v62 = int32(0)
	goto L7
}
func F_isValidClusterConfigFile(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 != 0 {
		v9 = int32(1)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a514)
		v9 = int32(0)
	}
	return v9
}
func F_isValidIpV4(m *base.Module, l0 int32, l1 int32) int32 {
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
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(1)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 == int32(0) {
		v20 = v9
	} else {
		v16 = F_inet_pton(m, int32(2), l0, v7+int32(4))
		mBase = m.M
		if v16 != 0 {
			v20 = v9
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a513)
			v20 = int32(0)
		}
	}
	m.G0 = v7 + int32(16)
	return v20
}
func F_isValidProcTitleTemplate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v5 = F_validateProcTitleTemplate(m, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v12 = int32(1)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a517)
			v12 = int32(0)
		}
		return v12
	}
}
func F_isValidShutdownOnSigFlags(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	v5 = int32(3)
	if l0&v5 != v5 {
		v12 = int32(1)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a519)
		v12 = int32(0)
	}
	return v12
}
