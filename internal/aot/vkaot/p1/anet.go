package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F__anetTcpServer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	var v24 int32
	_ = v24
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	v9 = m.G0
	v11 = v9 - int32(128)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = l1
	v20 = F_snprintf(m, v11+int32(118), int32(6), int32(_a114), v11+int32(64))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(112)))) = v24
	v31 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(104)))) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = v31
	v35 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v11)+88)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v35
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v102 = m.Env.Getaddrinfo(m, v94, v11+int32(118), v11+int32(84), v11+int32(80))
	mBase = m.M
	if v102 != 0 {
		goto L26
	} else {
		goto L27
	}
L4:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v42 == int32(42) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v94 = v24
	v95 = base.B2i32(l3 == int32(10))
	goto L3
L6:
	;
	if v50 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	v50 = int32(0) - v48
	goto L6
L8:
	;
	v50 = int32(42) - v42
	goto L6
L9:
	;
	v52 = l2
	goto L11
L10:
	;
	v52 = int32(0)
	goto L11
L11:
	;
	v53 = int32(10)
	v54 = base.B2i32(l3 == v53)
	if l3 != v53 {
		v94 = v52
		v95 = v54
		goto L3
	} else {
		goto L12
	}
L12:
	;
	if v52 == int32(0) {
		v94 = v52
		v95 = v54
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v59 = int32(0)
	v60 = int32(_a115)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[35])))
	if v64 == v59 {
		v87 = v63
		v88 = v64
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v88-v87&int32(255) != 0 {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	goto L14
L16:
	;
	if v64 != v63&int32(255) {
		v87 = v63
		v88 = v64
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v70 = v60
	v71 = v52
	goto L18
L18:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v75 == int32(0) {
		v87 = v74
		v88 = v75
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v87 = v74
	v88 = v75
	goto L15
L20:
	;
	v78 = int32(1)
	if v75 == v74&int32(255) {
		v70 = v70 + v78
		v71 = v71 + v78
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v92 = v52
	goto L24
L23:
	;
	v92 = v59
	goto L24
L24:
	;
	v94 = v92
	v95 = int32(1)
	goto L3
L25:
	;
	m.G0 = v11 + int32(128)
	return v372
L26:
	;
	v336 = int32(_a116)
	v338 = v102 + int32(1)
	if v338 == int32(0) {
		v358 = v336
		goto L105
	} else {
		goto L106
	}
L27:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	if v103 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+20))
	F_emscripten_builtin_free(m, v331)
	mBase = m.M
	F_emscripten_builtin_free(m, v330)
	mBase = m.M
	goto L103
L29:
	;
	v323 = int32(-1)
	goto L28
L30:
	;
	goto L101
L31:
	;
	if l5 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v108 = int32(262)
	goto L34
L33:
	;
	v108 = int32(6)
	goto L34
L34:
	;
	v110 = v103
	goto L35
L35:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	v119 = F_socket(m, v117, v118, v108)
	mBase = m.M
	if v119 == int32(-1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L30
L37:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v110)+28))
	if v298 != 0 {
		v110 = v298
		goto L35
	} else {
		goto L100
	}
L38:
	;
	if v95 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v206 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+124)) = v206
	v209 = int32(2)
	v211 = v11 + int32(124)
	v212 = int32(4)
	v216 = m.G0
	v218 = v216 - int32(16)
	m.G0 = v218
	v221 = F___syscall_setsockopt(m, v119, v206, v209, v211, v212, int32(0))
	mBase = m.M
	goto L72
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+124)) = int32(1)
	v134 = m.G0
	v136 = v134 - int32(16)
	m.G0 = v136
	v139 = F___syscall_setsockopt(m, v119, int32(41), int32(26), v11+int32(124), int32(4), int32(0))
	mBase = m.M
	goto L43
L41:
	;
	if v188 != int32(-1) {
		goto L39
	} else {
		goto L64
	}
L42:
	;
	m.G0 = v136 + int32(16)
	goto L41
L43:
	;
	v188 = F___syscall_ret(m, v139)
	mBase = m.M
	goto L42
L64:
	;
	goto L65
L65:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v198 = F___strerror_l(m, v197, v197)
	mBase = m.M
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v198
	F_anetSetError(m, l0, int32(_a117), v11+int32(16))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v205 = F_close(m, v119)
	mBase = m.M
	goto L29
L68:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	v291 = int32(0)
	v293 = F_anetListen(m, l0, v119, v289, v290, l4, v291, v291)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L96
	}
L69:
	;
	if v271 != int32(-1) {
		goto L68
	} else {
		goto L92
	}
L70:
	;
	m.G0 = v218 + int32(16)
	goto L69
L71:
	;
	v270 = F___syscall_ret(m, v268)
	mBase = m.M
	v271 = v270
	goto L70
L72:
	;
	if v221 != int32(-50) {
		v268 = v221
		goto L71
	} else {
		goto L73
	}
L73:
	;
	switch int32(-61) {
	case 0, 1:
		goto L74
	default:
		v268 = int32(-50)
		goto L71
	case 3, 4:
		goto L75
	}
L74:
	;
	goto L87
L75:
	;
	goto L77
L77:
	;
	v232 = F___syscall_ret(m, int32(-28))
	mBase = m.M
	v271 = v232
	goto L70
L87:
	;
	goto L88
L88:
	;
	goto L90
L90:
	;
	goto L91
L91:
	;
	v267 = F___syscall_setsockopt(m, v119, int32(1), v209, v211, v212, int32(0))
	mBase = m.M
	v268 = v267
	goto L71
L92:
	;
	goto L93
L93:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v280 = F___strerror_l(m, v279, v279)
	mBase = m.M
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v280
	F_anetSetError(m, l0, int32(_a118), v11+int32(32))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v287 = F_close(m, v119)
	mBase = m.M
	goto L29
L96:
	;
	if v293 == int32(-1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v297 = int32(-1)
	goto L99
L98:
	;
	v297 = v119
	goto L99
L99:
	;
	v323 = v297
	goto L28
L100:
	;
	goto L36
L101:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v308
	F_anetSetError(m, l0, int32(_a119), v11)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	goto L29
L103:
	;
	v372 = v323
	goto L25
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v358 + base.B2i32(v360 == int32(0))
	F_anetSetError(m, l0, int32(_a16), v11+int32(48))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L114
	}
L105:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
	goto L104
L106:
	;
	v342 = v336
	v343 = v338
	goto L107
L107:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	if v344 == int32(0) {
		v358 = v342
		goto L105
	} else {
		goto L109
	}
L108:
	;
	v358 = v354
	goto L105
L109:
	;
	v348 = v342
	goto L110
L110:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+1)))
	if v352 != 0 {
		v348 = v348 + int32(1)
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v354 = v348 + int32(2)
	v356 = v343 + int32(1)
	if v356 != 0 {
		v342 = v354
		v343 = v356
		goto L107
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	goto L108
L114:
	;
	v372 = int32(-1)
	goto L25
}
func F_anetBlock(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v13 = F_fcntl(m, l1, int32(3), v3)
	mBase = m.M
	if v13 != int32(-1) {
		if v13&int32(2048) == int32(0) {
			v41 = v3
			m.G0 = v8 + int32(32)
			return v41
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v13 & int32(-2049)
			v27 = F_fcntl(m, l1, int32(4), v8+int32(16))
			mBase = m.M
			if v27 != int32(-1) {
				v41 = v3
				m.G0 = v8 + int32(32)
				return v41
			} else {
				v31 = int32(_a111)
				v33 = *(*int32)(unsafe.Add(mBase, _consts[18]))
				v34 = F___strerror_l(m, v33, v33)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v34
				F_anetSetError(m, l0, v31, v8)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v41 = int32(-1)
					m.G0 = v8 + int32(32)
					return v41
				}
			}
		}
	} else {
		v31 = int32(_a112)
		v33 = *(*int32)(unsafe.Add(mBase, _consts[18]))
		v34 = F___strerror_l(m, v33, v33)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v34
		F_anetSetError(m, l0, v31, v8)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			v41 = int32(-1)
			m.G0 = v8 + int32(32)
			return v41
		}
	}
}
func F_anetCloexec(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	goto L4
L1:
	;
	m.G0 = v7 + int32(16)
	return v46
L2:
	;
	v46 = int32(-1)
	goto L1
L3:
	;
	if v15&int32(1) != 0 {
		v46 = v15
		goto L1
	} else {
		goto L9
	}
L4:
	;
	v15 = F_fcntl(m, l0, int32(1), int32(0))
	mBase = m.M
	if v15 != int32(-1) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L7
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	if v19 == int32(27) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L2
L9:
	;
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v15 | int32(1)
	v32 = F_fcntl(m, l0, int32(2), v7)
	mBase = m.M
	if v32 != int32(-1) {
		v46 = v32
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L2
L12:
	;
	goto L13
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	if v36 == int32(27) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
}
func F_anetFdToString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	v7 = m.G0
	v9 = v7 - int32(144)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(128)
	if l4 == int32(0) {
		v26 = F_getsockname(m, l0, v9+int32(16), v9+int32(12))
		mBase = m.M
		if v26 == int32(-1) {
			if l1 == int32(0) {
			} else {
				if base.Ui32(l2) < base.Ui32(int32(2)) {
					if l2 != int32(1) {
					} else {
						v78 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v78)
					}
				} else {
					v74 = int32(63)
					*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v74)
				}
			}
			v81 = int32(0)
			v82 = int32(-1)
			if l3 == v81 {
				v90 = v82
			} else {
				v86 = v82
				v87 = v81
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
				v90 = v86
			}
			m.G0 = v9 + int32(144)
			return v90
		} else {
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+16)))
			switch v29 + int32(-1) {
			case 0:
				if l1 == int32(0) {
					v64 = int32(0)
					if l3 != 0 {
						v86 = v64
						v87 = v64
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
						v90 = v86
					} else {
						v90 = v64
					}
					m.G0 = v9 + int32(144)
					return v90
				} else {
					v58 = F_snprintf(m, l1, l2, int32(_a120), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						if v58 < int32(0) {
							if base.Ui32(l2) < base.Ui32(int32(2)) {
								if l2 != int32(1) {
								} else {
									v78 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v78)
								}
							} else {
								v74 = int32(63)
								*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v74)
							}
							v81 = int32(0)
							v82 = int32(-1)
							if l3 == v81 {
								v90 = v82
							} else {
								v86 = v82
								v87 = v81
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
								v90 = v86
							}
						} else {
							if base.Ui32(l2) <= base.Ui32(v58) {
								if base.Ui32(l2) < base.Ui32(int32(2)) {
									if l2 != int32(1) {
									} else {
										v78 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v78)
									}
								} else {
									v74 = int32(63)
									*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v74)
								}
								v81 = int32(0)
								v82 = int32(-1)
								if l3 == v81 {
									v90 = v82
								} else {
									v86 = v82
									v87 = v81
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
									v90 = v86
								}
							} else {
								v64 = int32(0)
								if l3 != 0 {
									v86 = v64
									v87 = v64
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
									v90 = v86
								} else {
									v90 = v64
								}
							}
						}
						m.G0 = v9 + int32(144)
						return v90
					}
				}
			case 1:
				if l1 == int32(0) {
					if l3 != 0 {
						v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+18)))
						v68 = F___bswap_16_2(m, v67)
						mBase = m.M
						v86 = int32(0)
						v87 = v68
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
						v90 = v86
					} else {
						v90 = int32(0)
					}
					m.G0 = v9 + int32(144)
					return v90
				} else {
					v37 = F_inet_ntop(m, int32(2), v9+int32(20), l1, l2)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						if v37 == int32(0) {
							if base.Ui32(l2) < base.Ui32(int32(2)) {
								if l2 != int32(1) {
								} else {
									v78 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v78)
								}
							} else {
								v74 = int32(63)
								*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v74)
							}
							v81 = int32(0)
							v82 = int32(-1)
							if l3 == v81 {
								v90 = v82
							} else {
								v86 = v82
								v87 = v81
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
								v90 = v86
							}
						} else {
							if l3 != 0 {
								v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+18)))
								v68 = F___bswap_16_2(m, v67)
								mBase = m.M
								v86 = int32(0)
								v87 = v68
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
								v90 = v86
							} else {
								v90 = int32(0)
							}
						}
						m.G0 = v9 + int32(144)
						return v90
					}
				}
			default:
				if l1 == int32(0) {
				} else {
					if base.Ui32(l2) < base.Ui32(int32(2)) {
						if l2 != int32(1) {
						} else {
							v78 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v78)
						}
					} else {
						v74 = int32(63)
						*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v74)
					}
				}
				v81 = int32(0)
				v82 = int32(-1)
				if l3 == v81 {
					v90 = v82
				} else {
					v86 = v82
					v87 = v81
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
					v90 = v86
				}
				m.G0 = v9 + int32(144)
				return v90
			case 9:
				if l1 == int32(0) {
					if l3 != 0 {
						v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+18)))
						v68 = F___bswap_16_2(m, v67)
						mBase = m.M
						v86 = int32(0)
						v87 = v68
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
						v90 = v86
					} else {
						v90 = int32(0)
					}
					m.G0 = v9 + int32(144)
					return v90
				} else {
					v49 = F_inet_ntop(m, int32(10), v9+int32(24), l1, l2)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						if v49 == int32(0) {
							if base.Ui32(l2) < base.Ui32(int32(2)) {
								if l2 != int32(1) {
								} else {
									v78 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v78)
								}
							} else {
								v74 = int32(63)
								*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v74)
							}
							v81 = int32(0)
							v82 = int32(-1)
							if l3 == v81 {
								v90 = v82
							} else {
								v86 = v82
								v87 = v81
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
								v90 = v86
							}
						} else {
							if l3 != 0 {
								v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+18)))
								v68 = F___bswap_16_2(m, v67)
								mBase = m.M
								v86 = int32(0)
								v87 = v68
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
								v90 = v86
							} else {
								v90 = int32(0)
							}
						}
						m.G0 = v9 + int32(144)
						return v90
					}
				}
			}
		}
	} else {
		v19 = F_getpeername(m, l0, v9+int32(16), v9+int32(12))
		mBase = m.M
		if v19 != int32(-1) {
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+16)))
			switch v29 + int32(-1) {
			case 0:
				if l1 == int32(0) {
					v64 = int32(0)
					if l3 != 0 {
						v86 = v64
						v87 = v64
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
						v90 = v86
					} else {
						v90 = v64
					}
					m.G0 = v9 + int32(144)
					return v90
				} else {
					v58 = F_snprintf(m, l1, l2, int32(_a120), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						if v58 < int32(0) {
							if base.Ui32(l2) < base.Ui32(int32(2)) {
								if l2 != int32(1) {
								} else {
									v78 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v78)
								}
							} else {
								v74 = int32(63)
								*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v74)
							}
							v81 = int32(0)
							v82 = int32(-1)
							if l3 == v81 {
								v90 = v82
							} else {
								v86 = v82
								v87 = v81
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
								v90 = v86
							}
						} else {
							if base.Ui32(l2) <= base.Ui32(v58) {
								if base.Ui32(l2) < base.Ui32(int32(2)) {
									if l2 != int32(1) {
									} else {
										v78 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v78)
									}
								} else {
									v74 = int32(63)
									*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v74)
								}
								v81 = int32(0)
								v82 = int32(-1)
								if l3 == v81 {
									v90 = v82
								} else {
									v86 = v82
									v87 = v81
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
									v90 = v86
								}
							} else {
								v64 = int32(0)
								if l3 != 0 {
									v86 = v64
									v87 = v64
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
									v90 = v86
								} else {
									v90 = v64
								}
							}
						}
						m.G0 = v9 + int32(144)
						return v90
					}
				}
			case 1:
				if l1 == int32(0) {
					if l3 != 0 {
						v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+18)))
						v68 = F___bswap_16_2(m, v67)
						mBase = m.M
						v86 = int32(0)
						v87 = v68
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
						v90 = v86
					} else {
						v90 = int32(0)
					}
					m.G0 = v9 + int32(144)
					return v90
				} else {
					v37 = F_inet_ntop(m, int32(2), v9+int32(20), l1, l2)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						if v37 == int32(0) {
							if base.Ui32(l2) < base.Ui32(int32(2)) {
								if l2 != int32(1) {
								} else {
									v78 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v78)
								}
							} else {
								v74 = int32(63)
								*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v74)
							}
							v81 = int32(0)
							v82 = int32(-1)
							if l3 == v81 {
								v90 = v82
							} else {
								v86 = v82
								v87 = v81
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
								v90 = v86
							}
						} else {
							if l3 != 0 {
								v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+18)))
								v68 = F___bswap_16_2(m, v67)
								mBase = m.M
								v86 = int32(0)
								v87 = v68
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
								v90 = v86
							} else {
								v90 = int32(0)
							}
						}
						m.G0 = v9 + int32(144)
						return v90
					}
				}
			default:
				if l1 == int32(0) {
				} else {
					if base.Ui32(l2) < base.Ui32(int32(2)) {
						if l2 != int32(1) {
						} else {
							v78 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v78)
						}
					} else {
						v74 = int32(63)
						*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v74)
					}
				}
				v81 = int32(0)
				v82 = int32(-1)
				if l3 == v81 {
					v90 = v82
				} else {
					v86 = v82
					v87 = v81
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
					v90 = v86
				}
				m.G0 = v9 + int32(144)
				return v90
			case 9:
				if l1 == int32(0) {
					if l3 != 0 {
						v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+18)))
						v68 = F___bswap_16_2(m, v67)
						mBase = m.M
						v86 = int32(0)
						v87 = v68
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
						v90 = v86
					} else {
						v90 = int32(0)
					}
					m.G0 = v9 + int32(144)
					return v90
				} else {
					v49 = F_inet_ntop(m, int32(10), v9+int32(24), l1, l2)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						if v49 == int32(0) {
							if base.Ui32(l2) < base.Ui32(int32(2)) {
								if l2 != int32(1) {
								} else {
									v78 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v78)
								}
							} else {
								v74 = int32(63)
								*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v74)
							}
							v81 = int32(0)
							v82 = int32(-1)
							if l3 == v81 {
								v90 = v82
							} else {
								v86 = v82
								v87 = v81
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
								v90 = v86
							}
						} else {
							if l3 != 0 {
								v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+18)))
								v68 = F___bswap_16_2(m, v67)
								mBase = m.M
								v86 = int32(0)
								v87 = v68
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
								v90 = v86
							} else {
								v90 = int32(0)
							}
						}
						m.G0 = v9 + int32(144)
						return v90
					}
				}
			}
		} else {
			if l1 == int32(0) {
			} else {
				if base.Ui32(l2) < base.Ui32(int32(2)) {
					if l2 != int32(1) {
					} else {
						v78 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v78)
					}
				} else {
					v74 = int32(63)
					*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v74)
				}
			}
			v81 = int32(0)
			v82 = int32(-1)
			if l3 == v81 {
				v90 = v82
			} else {
				v86 = v82
				v87 = v81
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
				v90 = v86
			}
			m.G0 = v9 + int32(144)
			return v90
		}
	}
}
func F_anetIsFifo(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v4 = m.G0
	v6 = v4 - int32(96)
	m.G0 = v6
	v10 = F___fstatat(m, int32(-100), l0, v6, int32(0))
	mBase = m.M
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	m.G0 = v6 + int32(96)
	return base.B2i32(v10 != int32(-1)) & base.B2i32(v11&int32(61440) == int32(4096))
}
func F_anetRecvTimeout(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v47 int64
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int64(1000)
	v12 = base.I64_div_s(l2, v11)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = base.I32_wrap_i64(l2-v12*v11) * int32(1000)
	v23 = int32(66)
	v24 = int32(16)
	v25 = v9 + v24
	v30 = m.G0
	v32 = v30 - v24
	m.G0 = v32
	v35 = F___syscall_setsockopt(m, l1, int32(1), v23, v25, v24, v4)
	mBase = m.M
	if v35 != int32(-50) {
		v82 = v35
		v84 = F___syscall_ret(m, v82)
		mBase = m.M
		v85 = v84
	} else {
		switch int32(3) {
		case 0, 1:
			v81 = F___syscall_setsockopt(m, l1, int32(1), v23, v25, v24, int32(0))
			mBase = m.M
			v82 = v81
			v84 = F___syscall_ret(m, v82)
			mBase = m.M
			v85 = v84
		default:
			v82 = int32(-50)
			v84 = F___syscall_ret(m, v82)
			mBase = m.M
			v85 = v84
		case 3, 4:
			v47 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
			if base.Ui64(v47+int64(2147483648)) < base.Ui64(int64(4294967296)) {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v54
				*(*uint32)(unsafe.Add(mBase, uint32(v32)+8)) = uint32(v47)
				v66 = int32(8)
				v70 = F___syscall_setsockopt(m, l1, int32(1), int32(20), v32+v66, v66, int32(0))
				mBase = m.M
				v82 = v70
				v84 = F___syscall_ret(m, v82)
				mBase = m.M
				v85 = v84
			} else {
				v53 = F___syscall_ret(m, int32(-138))
				mBase = m.M
				v85 = v53
			}
		}
	}
	m.G0 = v32 + int32(16)
	if v85 != int32(-1) {
		v102 = v4
		m.G0 = v9 + int32(32)
		return v102
	} else {
		v93 = *(*int32)(unsafe.Add(mBase, _consts[18]))
		v94 = F___strerror_l(m, v93, v93)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v94
		F_anetSetError(m, l0, int32(_a113), v9)
		mBase = m.M
		v100 = m.ExcPending
		if v100 != 0 {
			return int32(0)
		} else {
			v102 = int32(-1)
			m.G0 = v9 + int32(32)
			return v102
		}
	}
}
func F_anetSetSockMarkId(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v9 int32
	_ = v9
	F_anetSetError(m, l0, int32(_a121), int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_anetTcpNonBlockConnect(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_anetTcpGenericConnect(m, l0, l1, l2, int32(0), int32(1))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_anetTcpServer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F__anetTcpServer(m, l0, l1, l2, int32(2), l3, l4)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anetUnixAccept(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(110)
	v14 = F_anetGenericAccept(m, l0, l1, v6+int32(18), v6+int32(12))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(128)
		return v14
	}
}
