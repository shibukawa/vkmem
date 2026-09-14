package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_auxAnnounceClientIpV4Getter(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2304))
	v4 = F_sdscat(m, l1, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_auxAnnounceClientIpV6Getter(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2308))
	v4 = F_sdscat(m, l1, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_auxAnnounceClientTlsPortSetter(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
		*(*int32)(unsafe.Add(mBase, uint32(l0)+2340)) = v74
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
func F_auxAvailabilityZoneSetter(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2320))
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
	if v28 != l2 {
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
	return int32(0)
L8:
	;
	v77 = F_sdscpylen(m, v8, l1, l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L24
	} else {
		goto L25
	}
L9:
	;
	if l2 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v74 == int32(0) {
		goto L7
	} else {
		goto L23
	}
L11:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v33 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v74 = int32(0)
	goto L10
L13:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v74 = v62 - v67
	goto L10
L14:
	;
	v35 = l1
	v36 = v8
	v37 = l2
	v38 = v33
	goto L17
L15:
	;
	v62 = int32(0)
	v63 = v8
	goto L13
L16:
	;
	v62 = v59 & int32(255)
	v63 = v57
	goto L13
L17:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v38&int32(255) != v42 {
		v57 = v36
		v59 = v38
		goto L16
	} else {
		goto L19
	}
L18:
	;
	v57 = v51
	v59 = int32(0)
	goto L16
L19:
	;
	if v42 == int32(0) {
		v57 = v36
		v59 = v38
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v47 = v37 + int32(-1)
	if v47 == int32(0) {
		v57 = v36
		v59 = v38
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v50 = int32(1)
	v51 = v36 + v50
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	if v52 != 0 {
		v35 = v35 + v50
		v36 = v51
		v37 = v47
		v38 = v52
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	goto L8
L24:
	;
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2320)) = v77
	goto L7
}
func F_auxHumanNodenameGetter(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2316))
	v4 = F_sdscat(m, l1, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_auxHumanNodenamePresent(m *base.Module, l0 int32) int32 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2316))
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
func F_auxTcpPortSetter(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
		*(*int32)(unsafe.Add(mBase, uint32(l0)+2324)) = v74
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
