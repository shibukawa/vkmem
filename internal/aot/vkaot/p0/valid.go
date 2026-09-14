package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_isValidAOFdirname(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v4 != 0 {
		v10 = F_strchr(m, l0, int32(47))
		mBase = m.M
		if v10 != 0 {
			v15 = int32(0)
		} else {
			v12 = F_strchr(m, l0, int32(92))
			mBase = m.M
			v15 = base.B2i32(v12 == int32(0))
		}
		if v15 != 0 {
			v22 = int32(1)
		} else {
			v17 = int32(_a435)
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
			v22 = int32(0)
		}
	} else {
		v17 = int32(_a436)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
		v22 = int32(0)
	}
	return v22
}
func F_isValidAOFfilename(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v4 != 0 {
		v10 = F_strchr(m, l0, int32(47))
		mBase = m.M
		if v10 != 0 {
			v15 = int32(0)
		} else {
			v12 = F_strchr(m, l0, int32(92))
			mBase = m.M
			v15 = base.B2i32(v12 == int32(0))
		}
		if v15 != 0 {
			v22 = int32(1)
		} else {
			v17 = int32(_a433)
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
			v22 = int32(0)
		}
	} else {
		v17 = int32(_a434)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
		v22 = int32(0)
	}
	return v22
}
func F_isValidAnnouncedHostname(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	if l0&int32(3) == int32(0) {
		v27 = l0
		goto L6
	} else {
		goto L7
	}
L1:
	;
	return v110
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v101
	v110 = int32(0)
	goto L1
L3:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v64 != 0 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	if base.Ui32(v60) <= base.Ui32(int32(255)) {
		goto L3
	} else {
		goto L20
	}
L5:
	;
	v60 = v52 - l0
	goto L4
L6:
	;
	v31 = v27
	goto L14
L7:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v13 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v16 = l0
	goto L10
L9:
	;
	v60 = l0 - l0
	goto L4
L10:
	;
	v20 = v16 + int32(1)
	if v20&int32(3) == int32(0) {
		v27 = v20
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v25 != 0 {
		v16 = v20
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v52 = v20
	goto L5
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v40 = int32(-2139062144)
	if (int32(16843008)-v37|v37)&v40 == v40 {
		v31 = v31 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v46 = v31
	goto L17
L16:
	;
	goto L15
L17:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 != 0 {
		v46 = v46 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v52 = v46
	goto L5
L19:
	;
	goto L18
L20:
	;
	v101 = int32(_a429)
	goto L2
L21:
	;
	v70 = v64
	v71 = int32(0)
	goto L23
L22:
	;
	return int32(1)
L23:
	;
	if base.Ui32((v70&int32(223)+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v94 = int32(1)
	v96 = v71 + v94
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v96))))
	if v98 != 0 {
		v70 = v98
		v71 = v96
		goto L23
	} else {
		goto L29
	}
L26:
	;
	if base.Ui32((v70+int32(-48))&int32(255)) < base.Ui32(int32(10)) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32((v70+int32(-45))&int32(255)) <= base.Ui32(int32(1)) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v101 = int32(_a430)
	goto L2
L29:
	;
	v110 = v94
	goto L1
}
func F_isValidAnnouncedIp(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	v11 = v9 & int32(7)
	switch v11 {
	case 0:
		goto L8
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	case 4:
		goto L10
	default:
		v44 = int32(0)
		goto L3
	}
L1:
	;
	return v87
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v80
	v87 = int32(0)
	goto L1
L3:
	;
	if v44 != 0 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v44 = v43
	goto L3
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v44 = v40
	goto L3
L6:
	;
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v44 = v37
	goto L3
L7:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v44 = v34
	goto L3
L8:
	;
	v44 = int32(base.Ui32(v9) >> (uint(int32(3)) % 32))
	goto L3
L9:
	;
	if base.Ui32(v24) <= base.Ui32(int32(45)) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v24 = v23
	goto L9
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v24 = v20
	goto L9
L12:
	;
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v24 = v17
	goto L9
L13:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v24 = v14
	goto L9
L14:
	;
	switch v11 + int32(-1) {
	default:
		goto L7
	case 1:
		goto L6
	case 2:
		goto L5
	case 3:
		goto L4
	}
L15:
	;
	v80 = int32(_a427)
	goto L2
L16:
	;
	if v78 != 0 {
		v87 = int32(1)
		goto L1
	} else {
		goto L26
	}
L17:
	;
	v54 = int32(0)
	goto L20
L18:
	;
	v78 = int32(1)
	goto L16
L19:
	;
	v78 = v72
	goto L16
L20:
	;
	v57 = int32(0)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v54))))
	if base.Ui32(v59) < base.Ui32(int32(45)) {
		v72 = v57
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v72 = v67
	goto L19
L22:
	;
	if v59 == int32(127) {
		v72 = v57
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v66 = F_memchr(m, int32(_a198), v59, int32(15))
	mBase = m.M
	if v66 != 0 {
		v72 = v57
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v67 = int32(1)
	v69 = v54 + v67
	if v69 != v44 {
		v54 = v69
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v80 = int32(_a428)
	goto L2
}
func F_isValidDBfilename(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v4 != 0 {
		v10 = F_strchr(m, l0, int32(47))
		mBase = m.M
		if v10 != 0 {
			v15 = int32(0)
		} else {
			v12 = F_strchr(m, l0, int32(92))
			mBase = m.M
			v15 = base.B2i32(v12 == int32(0))
		}
		if v15 != 0 {
			v22 = int32(1)
		} else {
			v17 = int32(_a431)
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
			v22 = int32(0)
		}
	} else {
		v17 = int32(_a432)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
		v22 = int32(0)
	}
	return v22
}
