package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_auxAnnounceClientIpV4Present(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2304))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+int32(-1)))))
	switch v8 & int32(7) {
	case 0:
		return base.B2i32(int32(base.Ui32(v8)>>(uint(int32(3))%32)) != int32(0))
	case 1:
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+int32(-3)))))
		return base.B2i32(v18 != int32(0))
	case 2:
		v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5+int32(-5)))))
		return base.B2i32(v24 != int32(0))
	case 3:
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v5+int32(-9))))
		return base.B2i32(v30 != int32(0))
	case 4:
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v5+int32(-17))))
		v37 = v36
		return base.B2i32(v37 != int32(0))
	default:
		v37 = int32(0)
		return base.B2i32(v37 != int32(0))
	}
}
func F_auxAnnounceClientIpV6Setter(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
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
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2308))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-1)))))
	switch v16 & int32(7) {
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
		v33 = int32(0)
		goto L1
	}
L1:
	;
	if v33 != l2 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
	v33 = v32
	goto L1
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
	v33 = v29
	goto L1
L4:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
	v33 = v26
	goto L1
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
	v33 = v23
	goto L1
L6:
	;
	v33 = int32(base.Ui32(v16) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	m.G0 = v10 + int32(32)
	return v98
L8:
	;
	v98 = int32(0)
	goto L7
L9:
	;
	if l2 == int32(0) {
		v90 = v13
		goto L25
	} else {
		goto L26
	}
L10:
	;
	if l2 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v79 == int32(0) {
		goto L8
	} else {
		goto L24
	}
L12:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v38 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v79 = int32(0)
	goto L11
L14:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v79 = v67 - v72
	goto L11
L15:
	;
	v40 = l1
	v41 = v13
	v42 = l2
	v43 = v38
	goto L18
L16:
	;
	v67 = int32(0)
	v68 = v13
	goto L14
L17:
	;
	v67 = v64 & int32(255)
	v68 = v62
	goto L14
L18:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v43&int32(255) != v47 {
		v62 = v41
		v64 = v43
		goto L17
	} else {
		goto L20
	}
L19:
	;
	v62 = v56
	v64 = int32(0)
	goto L17
L20:
	;
	if v47 == int32(0) {
		v62 = v41
		v64 = v43
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v52 = v42 + int32(-1)
	if v52 == int32(0) {
		v62 = v41
		v64 = v43
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v55 = int32(1)
	v56 = v41 + v55
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v57 != 0 {
		v40 = v40 + v55
		v41 = v56
		v42 = v52
		v43 = v57
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	goto L9
L25:
	;
	v91 = F_sdscpylen(m, v90, l1, l2)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v87 = F_inet_pton(m, int32(10), l1, v10+int32(12))
	mBase = m.M
	if v87 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2308))
	v90 = v89
	goto L25
L28:
	;
	v98 = int32(-1)
	goto L7
L29:
	;
	return int32(0)
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2308)) = v91
	goto L8
}
func F_auxAnnounceClientTcpPortSetter(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	v7 = m.G0
	if base.Ui32(l2+int32(-6)) < base.Ui32(int32(-5)) {
		v82 = int32(-1)
	} else {
		v17 = v7 - (l2+int32(16))&int32(-16)
		m.G0 = v17
		if l2 == int32(0) {
			v22 = v17
		} else {
			v21 = F__emscripten_memcpy_bulkmem(m, v17, l1, l2)
			mBase = m.M
			v22 = v21
		}
		v24 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v22+l2))) = uint8(v24)
		v29 = v22
		for {
			v34 = v29 + int32(1)
			v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v29))))
			v36 = F___isspace_1(m, v35)
			mBase = m.M
			if v36 != 0 {
				v29 = v34
				continue
			} else {
				break
			}
			break
		}
		v37 = int32(1)
		switch v35&int32(255) + int32(-43) {
		case 0:
			v43 = v37
			v44 = int32(*(*int8)(unsafe.Add(mBase, uint32(v34))))
			v45 = v34
			v46 = v44
			v47 = v43
		default:
			v45 = v29
			v46 = v35
			v47 = v37
		case 2:
			v43 = int32(0)
			v44 = int32(*(*int8)(unsafe.Add(mBase, uint32(v34))))
			v45 = v34
			v46 = v44
			v47 = v43
		}
		v50 = v46 + int32(-48)
		if base.Ui32(int32(9)) < base.Ui32(v50) {
			v68 = int32(0)
		} else {
			v54 = int32(0)
			v55 = v45
			v56 = v50
			for {
				v58 = int32(10)
				v60 = v54*v58 - v56
				v61 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55)+1)))
				v65 = v61 + int32(-48)
				if base.Ui32(v65) < base.Ui32(v58) {
					v54 = v60
					v55 = v55 + int32(1)
					v56 = v65
					continue
				} else {
					break
				}
				break
			}
			v68 = v60
		}
		if v47 != 0 {
			v74 = int32(0) - v68
		} else {
			v74 = v68
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+2336)) = v74
		if base.Ui32(int32(65535)) < base.Ui32(v74) {
			v80 = int32(-1)
		} else {
			v80 = int32(0)
		}
		v82 = v80
	}
	m.G0 = v7
	return v82
}
func F_auxAnnounceClientTlsPortPresent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2340))
	return base.B2i32(base.Ui32(v2+int32(-1)) < base.Ui32(int32(65535)))
}
func F_auxAvailabilityZoneGetter(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2320))
	v4 = F_sdscat(m, l1, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_auxAvailabilityZonePresent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2320))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+int32(-1)))))
	switch v8 & int32(7) {
	case 0:
		return int32(base.Ui32(v8) >> (uint(int32(3)) % 32))
	case 1:
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+int32(-3)))))
		return v16
	case 2:
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5+int32(-5)))))
		return v20
	case 3:
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v5+int32(-9))))
		return v24
	case 4:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v5+int32(-17))))
		v29 = v28
		return v29
	default:
		v29 = int32(0)
		return v29
	}
}
func F_auxShardIdPresent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	v3 = l0 + int32(48)
	if v3&int32(3) == int32(0) {
		v25 = v3
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v58
L2:
	;
	v58 = v50 - v3
	goto L1
L3:
	;
	v29 = v25
	goto L11
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v11 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v14 = v3
	goto L7
L6:
	;
	v58 = v3 - v3
	goto L1
L7:
	;
	v18 = v14 + int32(1)
	if v18&int32(3) == int32(0) {
		v25 = v18
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v23 != 0 {
		v14 = v18
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v50 = v18
	goto L2
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v38 = int32(-2139062144)
	if (int32(16843008)-v35|v35)&v38 == v38 {
		v29 = v29 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v44 = v29
	goto L14
L13:
	;
	goto L12
L14:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v48 != 0 {
		v44 = v44 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v50 = v44
	goto L2
L16:
	;
	goto L15
}
func F_auxTlsPortGetter(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2328))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v11 = F_sdscatfmt(m, l1, int32(_a248), v6)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v11
	}
}
func F_auxTlsPortSetter(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	v7 = m.G0
	if base.Ui32(l2+int32(-6)) < base.Ui32(int32(-5)) {
		v82 = int32(-1)
	} else {
		v17 = v7 - (l2+int32(16))&int32(-16)
		m.G0 = v17
		if l2 == int32(0) {
			v22 = v17
		} else {
			v21 = F__emscripten_memcpy_bulkmem(m, v17, l1, l2)
			mBase = m.M
			v22 = v21
		}
		v24 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v22+l2))) = uint8(v24)
		v29 = v22
		for {
			v34 = v29 + int32(1)
			v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v29))))
			v36 = F___isspace_1(m, v35)
			mBase = m.M
			if v36 != 0 {
				v29 = v34
				continue
			} else {
				break
			}
			break
		}
		v37 = int32(1)
		switch v35&int32(255) + int32(-43) {
		case 0:
			v43 = v37
			v44 = int32(*(*int8)(unsafe.Add(mBase, uint32(v34))))
			v45 = v34
			v46 = v44
			v47 = v43
		default:
			v45 = v29
			v46 = v35
			v47 = v37
		case 2:
			v43 = int32(0)
			v44 = int32(*(*int8)(unsafe.Add(mBase, uint32(v34))))
			v45 = v34
			v46 = v44
			v47 = v43
		}
		v50 = v46 + int32(-48)
		if base.Ui32(int32(9)) < base.Ui32(v50) {
			v68 = int32(0)
		} else {
			v54 = int32(0)
			v55 = v45
			v56 = v50
			for {
				v58 = int32(10)
				v60 = v54*v58 - v56
				v61 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55)+1)))
				v65 = v61 + int32(-48)
				if base.Ui32(v65) < base.Ui32(v58) {
					v54 = v60
					v55 = v55 + int32(1)
					v56 = v65
					continue
				} else {
					break
				}
				break
			}
			v68 = v60
		}
		if v47 != 0 {
			v74 = int32(0) - v68
		} else {
			v74 = v68
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+2328)) = v74
		if base.Ui32(int32(65535)) < base.Ui32(v74) {
			v80 = int32(-1)
		} else {
			v80 = int32(0)
		}
		v82 = v80
	}
	m.G0 = v7
	return v82
}
