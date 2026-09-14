package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_connCreateAcceptedSocket(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v5 = F_valkey_calloc(m, int32(40))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(1024)
		*(*uint16)(unsafe.Add(mBase, uint32(v5)+20)) = uint16(v9)
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a1477)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(2)
		return v5
	}
}
func F_connCreateAcceptedUnix(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v5 = F_valkey_calloc(m, int32(40))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(1024)
		*(*uint16)(unsafe.Add(mBase, uint32(v5)+20)) = uint16(v9)
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a1637)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(2)
		return v5
	}
}
func F_connRecvTimeout(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = F_anetRecvTimeout(m, int32(0), v4, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_connSocketAddr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = F_anetFdToString(m, v6, l1, l2, l3, l4)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v14
			return int32(-1)
		} else {
			return int32(0)
		}
	}
}
func F_connSocketBlockingConnect(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v6 = F_anetTcpNonBlockConnect(m, int32(0), l1, l2)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != int32(-1) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
			v22 = m.G0
			v24 = v22 - int32(16)
			m.G0 = v24
			*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v6
			v29 = int32(4)
			*(*uint16)(unsafe.Add(mBase, uint32(v24)+12)) = uint16(v29)
			v45 = int32(1)
			v47 = F_poll(m, v24+int32(8), v45, base.I32_wrap_i64(l3))
			mBase = m.M
			if v47 != v45 {
				v64 = v47
			} else {
				v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+14)))
				v51 = int32(1)
				v53 = int32(2)
				v57 = int32(base.Ui32(v50)>>(uint(v51)%32))&v53 | v50&v51
				if v50&int32(24) != 0 {
					v62 = v57 | v53
				} else {
					v62 = v57
				}
				v64 = v62
			}
			m.G0 = v24 + int32(16)
			if v64&int32(2) != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(3)
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(313532612613)
				return int32(-1)
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v15
			return int32(-1)
		}
	}
}
func F_connSocketCloseListener(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v6 < int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
	return
L2:
	;
	v11 = v6
	v12 = int32(0)
	goto L3
L3:
	;
	v17 = l0 + v12<<(uint(int32(2))%32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v18 == int32(-1) {
		v29 = v11
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v31 = v12 + int32(1)
	if v31 < v29 {
		v11 = v29
		v12 = v31
		goto L3
	} else {
		goto L9
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	F_aeDeleteFileEvent(m, v22, v18, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v27 = F_close(m, v26)
	mBase = m.M
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v29 = v28
	goto L5
L9:
	;
	goto L4
}
func F_connSocketEventHandler(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	v10 = l3 & int32(2)
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return
L2:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+52))
	m.T0[v202].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L19
	} else {
		goto L51
	}
L3:
	;
	v104 = int32(0)
	if l3&int32(1) == v104 {
		v135 = v104
		v136 = v101
		goto L28
	} else {
		goto L29
	}
L4:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v101 = v94
	v102 = v95
	v103 = base.B2i32(v97 != int32(0))
	goto L3
L5:
	;
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v89 = v87 & int32(2)
	v90 = int32(0)
	if v10 == v90 {
		v101 = v87
		v102 = v89
		v103 = v90
		goto L3
	} else {
		goto L27
	}
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v13 != int32(1) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v94 = v84
	v95 = v84 & int32(2)
	goto L4
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(0)
	v27 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v27
	v30 = v23 + int32(12)
	v37 = F_getsockopt(m, v18, int32(1), v27, v30, v23+int32(8))
	mBase = m.M
	if v37 != int32(-1) {
		v41 = v30
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v84 = v17
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v51 != 0 {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v42
	v49 = int32(5)
	goto L11
L13:
	;
	if v42 != 0 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	m.G0 = v23 + int32(16)
	goto L13
L15:
	;
	v40 = F___errno_location(m)
	mBase = m.M
	v41 = v40
	goto L14
L16:
	;
	v49 = int32(3)
	goto L11
L17:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)))
	v60 = v58 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)) = uint16(v60)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	if v62 == int32(0) {
		v70 = v58
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	F_aeDeleteFileEvent(m, v53, v54, int32(2))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	goto L17
L21:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)) = uint16(v70)
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	if v72&int32(1) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	m.T0[v62].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)))
	v70 = v67 + int32(-1)
	goto L21
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = int32(0)
	v84 = v72
	goto L8
L25:
	;
	if v70&int32(65535) == int32(0) {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	goto L1
L27:
	;
	v94 = v87
	v95 = v89
	goto L4
L28:
	;
	if v103 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v111 = base.B2i32(v109 != int32(0))
	if v102 != 0 {
		v135 = v111
		v136 = v101
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if v109 == int32(0) {
		v135 = v111
		v136 = v101
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v114 = int32(1)
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)))
	v117 = v115 + v114
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)) = uint16(v117)
	m.T0[v109].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)))
	v123 = v121 + int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)) = uint16(v123)
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	if v125&int32(1) == int32(0) {
		v135 = v114
		v136 = v125
		goto L28
	} else {
		goto L33
	}
L33:
	;
	if v123&int32(65535) == int32(0) {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	goto L1
L35:
	;
	v176 = v172 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)) = uint16(v176)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	if v178 == int32(0) {
		v187 = v172
		v188 = v173
		goto L46
	} else {
		goto L47
	}
L36:
	;
	if base.B2i32(v102 != int32(0))&v135 != int32(1) {
		goto L1
	} else {
		goto L45
	}
L37:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)))
	v141 = v139 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)) = uint16(v141)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v143 == int32(0) {
		v152 = v139
		v153 = v136
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)) = uint16(v152)
	if v153&int32(1) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	m.T0[v143].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)))
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v152 = v148 + int32(-1)
	v153 = v151
	goto L38
L41:
	;
	if base.B2i32(v102 != int32(0))&v135 != 0 {
		v172 = v152
		v173 = v153
		goto L35
	} else {
		goto L44
	}
L42:
	;
	if v152&int32(65535) == int32(0) {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	goto L1
L44:
	;
	goto L1
L45:
	;
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)))
	v172 = v171
	v173 = v136
	goto L35
L46:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)) = uint16(v187)
	if v188&int32(1) == int32(0) {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	m.T0[v178].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L19
	} else {
		goto L48
	}
L48:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)))
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v187 = v183 + int32(-1)
	v188 = v186
	goto L46
L49:
	;
	if v187&int32(65535) != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L2
L51:
	;
	goto L1
}
func F_connSocketGetLastError(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3 = F___strerror_l(m, v2, v2)
	mBase = m.M
	return v3
}
func F_connSocketIsLocal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int64
	_ = v9
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v9 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(39)))) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(32)))) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(24)))) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(16)))) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v9
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v31 = F_anetFdToString(m, v27, v5, int32(46), int32(0), int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		if v31 == int32(0) {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v47 = base.B2i32(v41 == int32(775369265)) | base.B2i32(v41 == int32(3226170))
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38
			v47 = int32(-1)
		}
		m.G0 = v5 + int32(48)
		return v47
	}
}
func F_connSocketListen(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_listenToPort(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_connSocketRead(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v53 int32
	_ = v53
	v6 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v6 == int32(0) {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v25 = F_read(m, v24, l1, l2)
		mBase = m.M
		if v25 != 0 {
			if int32(-1) < v25 {
			} else {
				v31 = int32(9116376)
				v32 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				if v32 == int32(6) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32
					v36 = *(*int32)(unsafe.Add(mBase, _consts[9]))
					if v36 == int32(27) {
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v39 != int32(3) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
						}
					}
				}
			}
			return v25
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(4)
			return v25
		}
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if v9&int32(4) == int32(0) {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v25 = F_read(m, v24, l1, l2)
			mBase = m.M
			if v25 != 0 {
				if int32(-1) < v25 {
				} else {
					v31 = int32(9116376)
					v32 = *(*int32)(unsafe.Add(mBase, _consts[9]))
					if v32 == int32(6) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32
						v36 = *(*int32)(unsafe.Add(mBase, _consts[9]))
						if v36 == int32(27) {
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v39 != int32(3) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
							}
						}
					}
				}
				return v25
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(4)
				return v25
			}
		} else {
			v14 = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[299]))
			if base.B2i32(v15 == v14) == int32(0) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v25 = F_read(m, v24, l1, l2)
				mBase = m.M
				if v25 != 0 {
					if int32(-1) < v25 {
					} else {
						v31 = int32(9116376)
						v32 = *(*int32)(unsafe.Add(mBase, _consts[9]))
						if v32 == int32(6) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32
							v36 = *(*int32)(unsafe.Add(mBase, _consts[9]))
							if v36 == int32(27) {
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v39 != int32(3) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
								}
							}
						}
					}
					return v25
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(4)
					return v25
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+222)))
				if v21 == int32(1) {
					F__serverAssert(m, int32(_a1480), int32(_a1479), int32(193))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v25 = F_read(m, v24, l1, l2)
					mBase = m.M
					if v25 != 0 {
						if int32(-1) < v25 {
						} else {
							v31 = int32(9116376)
							v32 = *(*int32)(unsafe.Add(mBase, _consts[9]))
							if v32 == int32(6) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32
								v36 = *(*int32)(unsafe.Add(mBase, _consts[9]))
								if v36 == int32(27) {
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v39 != int32(3) {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
									}
								}
							}
						}
						return v25
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(4)
						return v25
					}
				}
			}
		}
	}
}
func F_connSocketShutdown(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2 == int32(-1) {
	} else {
		v6 = int32(0)
		v10 = F___syscall_shutdown(m, v2, int32(2), v6, v6, v6, v6)
		mBase = m.M
		v11 = F___syscall_ret(m, v10)
		mBase = m.M
	}
	return
}
func F_connSocketSyncReadLine(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = l2 + int32(-1)
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v64
L2:
	;
	if v64 != int32(-1) {
		goto L1
	} else {
		goto L15
	}
L3:
	;
	m.G0 = v12 + int32(16)
	goto L2
L4:
	;
	v19 = l1
	v20 = v16
	v23 = int32(0)
	goto L6
L5:
	;
	v64 = int32(0)
	goto L3
L6:
	;
	v26 = int32(-1)
	v30 = F_syncRead(m, v5, v12+int32(15), int32(1), l3)
	mBase = m.M
	if v30 == v26 {
		v64 = v26
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v64 = v16
	goto L3
L8:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v33 != int32(10) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v48 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)) = uint8(v48)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v33)
	v51 = int32(1)
	v56 = v20 + int32(-1)
	if v56 != 0 {
		v19 = v19 + v51
		v20 = v56
		v23 = v23 + v51
		goto L6
	} else {
		goto L14
	}
L10:
	;
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v36)
	if v23 == v36 {
		v64 = v36
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v42 = v19 + int32(-1)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v43 != int32(13) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v64 = v23
	goto L3
L13:
	;
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v46)
	goto L12
L14:
	;
	goto L7
L15:
	;
	goto L16
L16:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v71
	goto L1
}
func F_connSocketWrite(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v50 int32
	_ = v50
	v6 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v6 == int32(0) {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v25 = F_write(m, v24, l1, l2)
		mBase = m.M
		if int32(-1) < v25 {
		} else {
			v28 = int32(9116376)
			v29 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			if v29 == int32(6) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v29
				v33 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				if v33 == int32(27) {
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v36 != int32(3) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
					}
				}
			}
		}
		return v25
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if v9&int32(4) == int32(0) {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v25 = F_write(m, v24, l1, l2)
			mBase = m.M
			if int32(-1) < v25 {
			} else {
				v28 = int32(9116376)
				v29 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				if v29 == int32(6) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v29
					v33 = *(*int32)(unsafe.Add(mBase, _consts[9]))
					if v33 == int32(27) {
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v36 != int32(3) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
						}
					}
				}
			}
			return v25
		} else {
			v14 = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[299]))
			if base.B2i32(v15 == v14) == int32(0) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v25 = F_write(m, v24, l1, l2)
				mBase = m.M
				if int32(-1) < v25 {
				} else {
					v28 = int32(9116376)
					v29 = *(*int32)(unsafe.Add(mBase, _consts[9]))
					if v29 == int32(6) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v29
						v33 = *(*int32)(unsafe.Add(mBase, _consts[9]))
						if v33 == int32(27) {
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v36 != int32(3) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
							}
						}
					}
				}
				return v25
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+223)))
				if v21 == int32(1) {
					F__serverAssert(m, int32(_a1478), int32(_a1479), int32(161))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v25 = F_write(m, v24, l1, l2)
					mBase = m.M
					if int32(-1) < v25 {
					} else {
						v28 = int32(9116376)
						v29 = *(*int32)(unsafe.Add(mBase, _consts[9]))
						if v29 == int32(6) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v29
							v33 = *(*int32)(unsafe.Add(mBase, _consts[9]))
							if v33 == int32(27) {
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v36 != int32(3) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
								}
							}
						}
					}
					return v25
				}
			}
		}
	}
}
func F_connTypeHasPendingData(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	if v3 == v1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v45
L2:
	;
	v14 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[208]))
	if v15 == v14 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)+104))
	if v6 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v9 = m.T0[v6].(func(*base.Module) int32)(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if v9 != 0 {
		v45 = v9
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L2
L8:
	;
	v24 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[209]))
	if v25 == v24 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+104))
	if v18 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v21 = m.T0[v18].(func(*base.Module) int32)(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	if v21 != 0 {
		v45 = v21
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	v34 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[210]))
	if v35 == v34 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+104))
	if v28 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v31 = m.T0[v28].(func(*base.Module) int32)(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	if v31 != 0 {
		v45 = v31
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	v45 = int32(0)
	goto L1
L19:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+104))
	if v38 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v41 = m.T0[v38].(func(*base.Module) int32)(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	if v41 != 0 {
		v45 = v41
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L18
}
func F_connTypeInitialize(m *base.Module) int32 {
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	v1 = F_RedisRegisterConnectionTypeSocket(m)
	v4 = m.ExcPending
	if v4 != 0 {
		return int32(0)
	} else {
		if v1 != 0 {
			F__serverAssert(m, int32(_a470), int32(_a471), int32(48))
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v5 = F_RedisRegisterConnectionTypeUnix(m)
			v6 = m.ExcPending
			if v6 != 0 {
				return int32(0)
			} else {
				if v5 != 0 {
					F__serverAssert(m, int32(_a472), int32(_a471), int32(51))
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v7 = F_RedisRegisterConnectionTypeTLS(m)
					v8 = m.ExcPending
					if v8 != 0 {
						return int32(0)
					} else {
						v9 = F_RegisterConnectionTypeRdma(m)
						v10 = m.ExcPending
						if v10 != 0 {
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
func F_connTypeOfReplication(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v2 == int32(0) {
		v10 = F_connectionTypeTcp(m)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	} else {
		v5 = F_connectionTypeTls(m)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v5
		}
	}
}
func F_connTypeProcessPendingData(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v1 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	if v5 == v1 {
		v16 = v1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[208]))
	if v19 == v18 {
		v28 = v16
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v8 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+108))
	if v9 == v8 {
		v16 = v8
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = m.T0[v9].(func(*base.Module) int32)(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v16 = v12
	goto L1
L6:
	;
	v30 = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, _consts[209]))
	if v31 == v30 {
		v40 = v28
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v22 == int32(0) {
		v28 = v16
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v25 = m.T0[v22].(func(*base.Module) int32)(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v28 = v25 + v16
	goto L6
L10:
	;
	v42 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, _consts[210]))
	if v43 == v42 {
		v52 = v40
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+108))
	if v34 == int32(0) {
		v40 = v28
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v37 = m.T0[v34].(func(*base.Module) int32)(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v40 = v37 + v28
	goto L10
L14:
	;
	return v52
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+108))
	if v46 == int32(0) {
		v52 = v40
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v49 = m.T0[v46].(func(*base.Module) int32)(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v52 = v49 + v40
	goto L14
}
func F_connUnixAccept(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = F_connectionTypeTcp(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+64))
		v8 = m.T0[v7].(func(*base.Module, int32, int32) int32)(m, l0, l1)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_connUnixAddr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
	v15 = F_snprintf(m, l1, l2, int32(_a1636), v9)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if l3 == int32(0) {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
		}
		m.G0 = v9 + int32(16)
		return int32(0)
	}
}
func F_connUnixIsLocal(m *base.Module, l0 int32) int32 {
	return int32(1)
}
func F_connUnixRead(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = F_connectionTypeTcp(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+76))
		v9 = m.T0[v8].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_connUnixSetReadHandler(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = F_connectionTypeTcp(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+84))
		v8 = m.T0[v7].(func(*base.Module, int32, int32) int32)(m, l0, l1)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_connUnixSetWriteHandler(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = F_connectionTypeTcp(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+80))
		v9 = m.T0[v8].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_connUnixSyncRead(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v27 int32
	_ = v27
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
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = int32(0)
	v13 = F_mstime(m)
	mBase = m.M
	if l2 == v6 {
		v72 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v72
L2:
	;
	goto L1
L3:
	;
	v17 = l1
	v18 = l2
	v20 = v6
	v22 = l3
	goto L6
L4:
	;
	v72 = int32(-1)
	goto L2
L5:
	;
	v59 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v57
	goto L4
L6:
	;
	v27 = F_read(m, v5, v17, v18)
	mBase = m.M
	switch v27 + int32(1) {
	case 0:
		goto L10
	case 1:
		v57 = int32(15)
		goto L5
	default:
		goto L9
	}
L7:
	;
	v57 = int32(73)
	goto L5
L8:
	;
	v43 = int64(10)
	if v43 < v22 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v34 = v27 + v20
	v35 = v18 - v27
	if v35 == int32(0) {
		v72 = v34
		goto L2
	} else {
		goto L12
	}
L10:
	;
	v30 = F___errno_location(m)
	mBase = m.M
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v31 == int32(6) {
		v39 = v17
		v40 = v18
		v41 = v20
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L4
L12:
	;
	v39 = v17 + v27
	v40 = v35
	v41 = v34
	goto L8
L13:
	;
	v46 = v22
	goto L15
L14:
	;
	v46 = v43
	goto L15
L15:
	;
	v47 = F_aeWait(m, v5, int32(1), v46)
	mBase = m.M
	v48 = F_mstime(m)
	mBase = m.M
	v49 = v48 - v13
	if v49 < l3 {
		v17 = v39
		v18 = v40
		v20 = v41
		v22 = l3 - v49
		goto L6
	} else {
		goto L16
	}
L16:
	;
	goto L7
}
func F_connUnixWrite(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = F_connectionTypeTcp(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+68))
		v9 = m.T0[v8].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
