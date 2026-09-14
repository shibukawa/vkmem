package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_auxAnnounceClientIpV4Setter(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2304))
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
	m.G0 = v10 + int32(16)
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
	v87 = F_inet_pton(m, int32(2), l1, v10+int32(4))
	mBase = m.M
	if v87 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2304))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2304)) = v91
	goto L8
}
func F_auxAnnounceClientIpV6Present(m *base.Module, l0 int32) int32 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2308))
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
func F_auxAnnounceClientTcpPortGetter(m *base.Module, l0 int32, l1 int32) int32 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2336))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v11 = F_sdscatfmt(m, l1, int32(_a232), v6)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v11
	}
}
func F_auxAnnounceClientTcpPortPresent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2336))
	return base.B2i32(base.Ui32(v2+int32(-1)) < base.Ui32(int32(65535)))
}
func F_auxAnnounceClientTlsPortGetter(m *base.Module, l0 int32, l1 int32) int32 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2340))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v11 = F_sdscatfmt(m, l1, int32(_a232), v6)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v11
	}
}
func F_auxHumanNodenameSetter(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2316))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2316)) = v77
	goto L7
}
func F_auxShardIdGetter(m *base.Module, l0 int32, l1 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_sdscatlen(m, l1, l0+int32(48), int32(40))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_auxShardIdSetter(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v60 int64
	_ = v60
	var v66 int64
	_ = v66
	var v72 int64
	_ = v72
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v210 int64
	_ = v210
	var v214 int64
	_ = v214
	var v218 int64
	_ = v218
	var v220 int32
	_ = v220
	var v224 int64
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = int32(-1)
	if l2 != int32(40) {
		v51 = v18
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v257
L2:
	;
	if v51 == int32(-1) {
		v257 = v18
		goto L1
	} else {
		goto L10
	}
L3:
	;
	goto L2
L4:
	;
	v26 = int32(0)
	goto L6
L5:
	;
	v51 = int32(0) - v41
	goto L3
L6:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v26))))
	v31 = int32(255)
	v41 = base.B2i32(base.Ui32((v28+int32(-123))&v31) < base.Ui32(int32(230))) & base.B2i32(base.Ui32((v28+int32(-58))&v31) < base.Ui32(int32(246)))
	if v41 != 0 {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v43 = v26 + int32(1)
	if v43 != int32(40) {
		v26 = v43
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v54
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(32))))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(80)))) = v60
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(24))))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(72)))) = v66
	v72 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(64)))) = v72
	v78 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(56)))) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2164))
	if v80 < int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_clusterAddNodeToShard(m, l1, l0)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L35
	} else {
		goto L40
	}
L12:
	;
	v87 = l0 + int32(48)
	v99 = v80
	v101 = int32(0)
	goto L13
L13:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2168))
	v109 = v101 << (uint(int32(2)) % 32)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v109)))
	v113 = v111 + int32(48)
	v114 = int32(40)
	goto L20
L14:
	;
	goto L11
L15:
	;
	v235 = v101 + int32(1)
	if v235 < v233 {
		v99 = v233
		v101 = v235
		goto L13
	} else {
		goto L39
	}
L16:
	;
	if v178 == int32(0) {
		v233 = v99
		goto L15
	} else {
		goto L32
	}
L17:
	;
	v178 = int32(0)
	goto L16
L18:
	;
	v150 = v145
	v151 = v146
	v152 = v147
	goto L28
L19:
	;
	if v135 == int32(0) {
		goto L17
	} else {
		goto L26
	}
L20:
	;
	if (v87|v113)&int32(3) != 0 {
		v145 = v113
		v146 = v87
		v147 = v114
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v122 = v113
	v123 = v87
	v124 = v114
	goto L22
L22:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v127 != v128 {
		v145 = v122
		v146 = v123
		v147 = v124
		goto L18
	} else {
		goto L24
	}
L23:
	;
	goto L19
L24:
	;
	v130 = int32(4)
	v131 = v123 + v130
	v133 = v122 + v130
	v135 = v124 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v135) {
		v122 = v133
		v123 = v131
		v124 = v135
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v145 = v133
	v146 = v131
	v147 = v135
	goto L18
L27:
	;
	v178 = v155 - v156
	goto L16
L28:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v155 != v156 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v158 = int32(1)
	v163 = v152 + int32(-1)
	if v163 == int32(0) {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	v150 = v150 + v158
	v151 = v151 + v158
	v152 = v163
	goto L28
L32:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v182 {
		v200 = v111
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_clusterRemoveNodeFromShard(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L35
	} else {
		goto L37
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l0 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v111 + int32(8)
	F__serverLog(m, int32(2), int32(_a231), v16)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	return int32(0)
L36:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2168))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v197+v109)))
	v200 = v199
	goto L33
L37:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2168))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v203+v109)))
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
	*(*int64)(unsafe.Add(mBase, uint32(v205)+48)) = v206
	v210 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(80))))
	*(*int64)(unsafe.Add(mBase, uint32(v205+int32(80)))) = v210
	v214 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(72))))
	*(*int64)(unsafe.Add(mBase, uint32(v205+int32(72)))) = v214
	v218 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(64))))
	*(*int64)(unsafe.Add(mBase, uint32(v205+int32(64)))) = v218
	v220 = int32(56)
	v224 = *(*int64)(unsafe.Add(mBase, uint32(l0+v220)))
	*(*int64)(unsafe.Add(mBase, uint32(v205+v220))) = v224
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2168))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v226+v109)))
	F_clusterAddNodeToShard(m, v87, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2164))
	v233 = v231
	goto L15
L39:
	;
	goto L14
L40:
	;
	v257 = int32(0)
	goto L1
}
func F_auxTcpPortGetter(m *base.Module, l0 int32, l1 int32) int32 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2324))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v11 = F_sdscatfmt(m, l1, int32(_a232), v6)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v11
	}
}
func F_auxTcpPortPresent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2324))
	return base.B2i32(base.Ui32(v2) < base.Ui32(int32(65536)))
}
func F_auxTlsPortPresent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2328))
	return base.B2i32(base.Ui32(v2) < base.Ui32(int32(65536)))
}
