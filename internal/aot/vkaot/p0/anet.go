package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_anetCreateSocket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v15 = l4 & int32(1)
	if v15 != 0 {
		v21 = l4 & int32(6)
	} else {
		v21 = l4
	}
	v27 = F_socket(m, l1, v15<<(uint(int32(19))%32)|l2|v21<<(uint(int32(10))%32)&int32(2048), l3)
	mBase = m.M
	if v27 != int32(-1) {
		if base.Ui32(int32(4)) <= base.Ui32(v21) {
			v41 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v41
			v45 = int32(2)
			v47 = v11 + int32(28)
			v48 = int32(4)
			v52 = m.G0
			v54 = v52 - int32(16)
			m.G0 = v54
			v57 = F___syscall_setsockopt(m, v27, v41, v45, v47, v48, int32(0))
			mBase = m.M
			if v57 != int32(-50) {
				v104 = v57
				v106 = F___syscall_ret(m, v104)
				mBase = m.M
				v107 = v106
			} else {
				switch int32(-61) {
				case 0, 1:
					v103 = F___syscall_setsockopt(m, v27, int32(1), v45, v47, v48, int32(0))
					mBase = m.M
					v104 = v103
					v106 = F___syscall_ret(m, v104)
					mBase = m.M
					v107 = v106
				default:
					v104 = int32(-50)
					v106 = F___syscall_ret(m, v104)
					mBase = m.M
					v107 = v106
				case 3, 4:
					v68 = F___syscall_ret(m, int32(-28))
					mBase = m.M
					v107 = v68
				}
			}
			m.G0 = v54 + int32(16)
			if v107 == int32(-1) {
				v115 = *(*int32)(unsafe.Add(mBase, _c_F_anetCreateSocket[0]))
				v116 = F___strerror_l(m, v115, v115)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v116
				F_anetSetError(m, l0, int32(_a_F_anetCreateSocket_0), v11+int32(16))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					v123 = F_close(m, v27)
					mBase = m.M
					v124 = int32(-1)
					m.G0 = v11 + int32(32)
					return v124
				}
			} else {
				v124 = v27
				m.G0 = v11 + int32(32)
				return v124
			}
		} else {
			v124 = v27
			m.G0 = v11 + int32(32)
			return v124
		}
	} else {
		v31 = *(*int32)(unsafe.Add(mBase, _c_F_anetCreateSocket[0]))
		v32 = F___strerror_l(m, v31, v31)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v32
		F_anetSetError(m, l0, int32(_a_F_anetCreateSocket_1), v11)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			v124 = int32(-1)
			m.G0 = v11 + int32(32)
			return v124
		}
	}
}
func F_anetDisableTcpNoDelay(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v3 = int32(0)
	v5 = m.G0
	v6 = int32(16)
	v7 = v5 - v6
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v3
	v20 = m.G0
	v22 = v20 - v6
	m.G0 = v22
	v25 = F___syscall_setsockopt(m, l1, int32(6), int32(1), v7+int32(12), int32(4), v3)
	mBase = m.M
	v74 = F___syscall_ret(m, v25)
	mBase = m.M
	m.G0 = v22 + int32(16)
	if v74 != int32(-1) {
		v92 = v3
		m.G0 = v7 + int32(16)
		return v92
	} else {
		v83 = *(*int32)(unsafe.Add(mBase, _c_F_anetDisableTcpNoDelay[0]))
		v84 = F___strerror_l(m, v83, v83)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v84
		F_anetSetError(m, l0, int32(_a_F_anetDisableTcpNoDelay_0), v7)
		mBase = m.M
		v90 = m.ExcPending
		if v90 != 0 {
			return int32(0)
		} else {
			v92 = int32(-1)
			m.G0 = v7 + int32(16)
			return v92
		}
	}
}
func F_anetGenericAccept(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	goto L5
L1:
	;
	m.G0 = v9 + int32(80)
	return v124
L2:
	;
	v90 = int32(-1)
	v93 = F_fcntl(m, v17, int32(3), int32(0))
	mBase = m.M
	if v93 != v90 {
		goto L28
	} else {
		goto L29
	}
L3:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v75 = F___strerror_l(m, v74, v74)
	mBase = m.M
	goto L25
L4:
	;
	if v28&int32(1) != 0 {
		goto L2
	} else {
		goto L19
	}
L5:
	;
	v17 = F_accept(m, l1, l2, l3)
	mBase = m.M
	if v17 == int32(-1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v39 = F___strerror_l(m, v36, v36)
	mBase = m.M
	goto L16
L7:
	;
	goto L14
L8:
	;
	goto L9
L9:
	;
	v28 = F_fcntl(m, v17, int32(1), int32(0))
	mBase = m.M
	if v28 != int32(-1) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v31 = int32(9116376)
	goto L12
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_anetGenericAccept[0]))
	if v32 == int32(27) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v71 = v31
	goto L3
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_anetGenericAccept[0]))
	if v36 == int32(27) {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	goto L6
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v39
	F_anetSetError(m, l0, int32(_a_F_anetGenericAccept_0), v9)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	v124 = int32(-1)
	goto L1
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v28 | int32(1)
	v61 = F_fcntl(m, v17, int32(2), v9+int32(64))
	mBase = m.M
	if v61 != int32(-1) {
		goto L2
	} else {
		goto L22
	}
L21:
	;
	v71 = v64
	goto L3
L22:
	;
	v64 = int32(9116376)
	goto L23
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_anetGenericAccept[0]))
	if v65 == int32(27) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v75
	F_anetSetError(m, l0, int32(_a_F_anetGenericAccept_1), v9+int32(16))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	v82 = F_close(m, v17)
	mBase = m.M
	v124 = int32(-1)
	goto L1
L27:
	;
	goto L34
L28:
	;
	if v93&int32(2048) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v111 = int32(_a_F_anetGenericAccept_2)
	goto L27
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v93 | int32(2048)
	v107 = F_fcntl(m, v17, int32(4), v9+int32(48))
	mBase = m.M
	if v107 == int32(-1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v124 = v17
	goto L1
L32:
	;
	v111 = int32(_a_F_anetGenericAccept_3)
	goto L27
L33:
	;
	v124 = v17
	goto L1
L34:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_anetGenericAccept[0]))
	v114 = F___strerror_l(m, v113, v113)
	mBase = m.M
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v114
	F_anetSetError(m, l0, v111, v9+int32(32))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L17
	} else {
		goto L36
	}
L36:
	;
	v120 = F_close(m, v17)
	mBase = m.M
	v124 = v90
	goto L1
}
func F_anetGetError(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
	v10 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v10
	v13 = v6 + int32(12)
	v20 = F_getsockopt(m, l0, int32(1), v10, v13, v6+int32(8))
	mBase = m.M
	if v20 != int32(-1) {
		v24 = v13
	} else {
		v24 = int32(9116376)
	}
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	m.G0 = v6 + int32(16)
	return v25
}
func F_anetNonBlock(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v13 = F_fcntl(m, l1, int32(3), v3)
	mBase = m.M
	if v13 != int32(-1) {
		if v13&int32(2048) != 0 {
			v39 = v3
			m.G0 = v8 + int32(32)
			return v39
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v13 | int32(2048)
			v25 = F_fcntl(m, l1, int32(4), v8+int32(16))
			mBase = m.M
			if v25 != int32(-1) {
				v39 = v3
				m.G0 = v8 + int32(32)
				return v39
			} else {
				v29 = int32(_a_F_anetNonBlock_0)
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_anetNonBlock[0]))
				v32 = F___strerror_l(m, v31, v31)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v32
				F_anetSetError(m, l0, v29, v8)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v39 = int32(-1)
					m.G0 = v8 + int32(32)
					return v39
				}
			}
		}
	} else {
		v29 = int32(_a_F_anetNonBlock_1)
		v31 = *(*int32)(unsafe.Add(mBase, _c_F_anetNonBlock[0]))
		v32 = F___strerror_l(m, v31, v31)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v32
		F_anetSetError(m, l0, v29, v8)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			v39 = int32(-1)
			m.G0 = v8 + int32(32)
			return v39
		}
	}
}
func F_anetResolve(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v14 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(32)))) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v14
	if l4&int32(1) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(0)
	v35 = l4 & int32(6)
	switch v35 + int32(-2) {
	case 0:
		v39 = v35
		goto L4
	default:
		goto L3
	case 2:
		goto L5
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(4)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(1)
	v44 = int32(0)
	v49 = m.Env.Getaddrinfo(m, l1, v44, v10+int32(16), v10+int32(12))
	mBase = m.M
	if v49 == v44 {
		v67 = v49
		goto L6
	} else {
		goto L7
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v39
	goto L3
L5:
	;
	v39 = int32(10)
	goto L4
L6:
	;
	if v67 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v52 == int32(0) {
		v67 = v49
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v56 = int32(2)
	if v52 == v56 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v59 = int32(10)
	goto L11
L10:
	;
	v59 = v56
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v59
	v66 = m.Env.Getaddrinfo(m, l1, int32(0), v10+int32(16), v10+int32(12))
	mBase = m.M
	v67 = v66
	goto L6
L12:
	;
	m.G0 = v10 + int32(48)
	return v128
L13:
	;
	v108 = int32(2)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v113 = base.B2i32(v111 == v108)
	if v111 == v108 {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v73 = int32(_a_F_anetResolve_0)
	v75 = v67 + int32(1)
	if v75 == int32(0) {
		v95 = v73
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v95 + base.B2i32(v97 == int32(0))
	F_anetSetError(m, l0, int32(_a_F_anetResolve_1), v10)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L25
	} else {
		goto L26
	}
L16:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	goto L15
L17:
	;
	v79 = v73
	v80 = v75
	goto L18
L18:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v81 == int32(0) {
		v95 = v79
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v95 = v91
	goto L16
L20:
	;
	v85 = v79
	goto L21
L21:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v89 != 0 {
		v85 = v85 + int32(1)
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v91 = v85 + int32(2)
	v93 = v80 + int32(1)
	if v93 != 0 {
		v79 = v91
		v80 = v93
		goto L18
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	goto L19
L25:
	;
	return int32(0)
L26:
	;
	v128 = int32(-1)
	goto L12
L27:
	;
	v114 = v108
	goto L29
L28:
	;
	v114 = int32(10)
	goto L29
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	if v111 == v108 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v118 = int32(4)
	goto L32
L31:
	;
	v118 = int32(8)
	goto L32
L32:
	;
	v120 = F_inet_ntop(m, v114, v115+v118, l2, l3)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+20))
	F_emscripten_builtin_free(m, v123)
	mBase = m.M
	F_emscripten_builtin_free(m, v122)
	mBase = m.M
	goto L34
L34:
	;
	v128 = int32(0)
	goto L12
}
func F_anetSetError(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == int32(0) {
		m.G0 = v7 + int32(16)
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
		v13 = F_vsnprintf(m, l0, int32(256), l1, l2)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_anetUnixServer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	v9 = m.G0
	v11 = v9 - int32(128)
	m.G0 = v11
	if l1&int32(3) == int32(0) {
		v34 = l1
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v11 + int32(128)
	return v148
L2:
	;
	v80 = int32(1)
	v84 = F_anetCreateSocket(m, l0, v80, v80, int32(0), int32(3))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L20
	} else {
		goto L22
	}
L3:
	;
	if base.Ui32(v67) < base.Ui32(int32(108)) {
		goto L2
	} else {
		goto L19
	}
L4:
	;
	v67 = v59 - l1
	goto L3
L5:
	;
	v38 = v34
	goto L13
L6:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v23 = l1
	goto L9
L8:
	;
	v67 = l1 - l1
	goto L3
L9:
	;
	v27 = v23 + int32(1)
	if v27&int32(3) == int32(0) {
		v34 = v27
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		v23 = v27
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v59 = v27
	goto L4
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v47 = int32(-2139062144)
	if (int32(16843008)-v44|v44)&v47 == v47 {
		v38 = v38 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v53 = v38
	goto L16
L15:
	;
	goto L14
L16:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 != 0 {
		v53 = v53 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v59 = v53
	goto L4
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(108)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
	F_anetSetError(m, l0, int32(_a_F_anetUnixServer_0), v11)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	v148 = int32(-1)
	goto L1
L22:
	;
	if v84 == int32(-1) {
		v148 = int32(-1)
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v93 = F__emscripten_memset_bulkmem(m, v11+int32(20), base.I32_extend8_s(int32(0)), int32(108))
	mBase = m.M
	goto L24
L24:
	;
	v94 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+18)) = uint16(v94)
	goto L28
L25:
	;
	v143 = F_anetListen(m, l0, v84, v11+int32(18), int32(110), l3, l2, l4)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L20
	} else {
		goto L37
	}
L26:
	;
	goto L25
L27:
	;
	v126 = v105
	goto L34
L28:
	;
	v101 = v93
	v103 = int32(108)
	v105 = l1
	goto L30
L29:
	;
	v116 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v116)
	goto L27
L30:
	;
	v107 = v103 + int32(-1)
	if v107 == int32(0) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v110)
	v112 = int32(1)
	if v110 != 0 {
		v101 = v101 + v112
		v103 = v107
		v105 = v105 + v112
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L26
L34:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v128 != 0 {
		v126 = v126 + int32(1)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	goto L26
L36:
	;
	goto L35
L37:
	;
	if v143 == int32(-1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v147 = int32(-1)
	goto L40
L39:
	;
	v147 = v84
	goto L40
L40:
	;
	v148 = v147
	goto L1
}
