package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_initDeferredReplyBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int64
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, _consts[441]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	if v5 == int32(0) {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
		if v8 != 0 {
			v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
			if v15 != int64(-1) {
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+384)) = int64(0)
			}
			return
		} else {
			v9 = F_listCreate(m)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+376)) = v9
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(953)
				v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
				if v15 != int64(-1) {
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+384)) = int64(0)
				}
				return
			}
		}
	}
}
func F_setDeferredArrayLen(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v6 int32
	_ = v6
	F_setDeferredAggregateLen(m, l0, l1, l2, int32(42))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_setDeferredMapLen(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v7 = base.B2i32(v5 == int32(2))
	if v5 == int32(2) {
		v11 = int32(42)
	} else {
		v11 = int32(37)
	}
	F_setDeferredAggregateLen(m, l0, l1, l2<<(uint(v7)%32), v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		return
	}
}
func F_setDeferredReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v243 int32
	_ = v243
	var v245 int64
	_ = v245
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v277 int64
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int64
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v309 int32
	_ = v309
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	if l1 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a1651), int32(_a1630), int32(1181))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L23
	} else {
		goto L86
	}
L2:
	;
	m.G0 = v14 + int32(16)
	return
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v18 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
	v23 = base.B2i32(v21 == int64(-1))
	if v21 == int64(-1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = int32(132)
	goto L7
L6:
	;
	v24 = int32(376)
	goto L7
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+v24)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v27 == int32(0) {
		v64 = l2
		v65 = l3
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v70 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v30 == int32(0) {
		v64 = l2
		v65 = l3
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if base.Ui32(v33) <= base.Ui32(v34) {
		v64 = l2
		v65 = l3
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v36 == int32(1) {
		v64 = l2
		v65 = l3
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v40 = v30 + int32(13)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v40 == v41 {
		v64 = l2
		v65 = l3
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+12)))
	if v43&int32(1) != 0 {
		v64 = l2
		v65 = l3
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v47 = v33 - v34
	if base.Ui32(v47) < base.Ui32(l3) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v49 = v47
	goto L17
L16:
	;
	v49 = l3
	goto L17
L17:
	;
	if v49 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+264)) = v54 + base.I64_extend_i32_u(v49)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v49 + v34
	v60 = l3 - v49
	if v60 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L18
L20:
	;
	v52 = F__emscripten_memcpy_bulkmem(m, v40+v34, l2, v49)
	mBase = m.M
	goto L19
L21:
	;
	v64 = l2 + v49
	v65 = v60
	goto L8
L22:
	;
	F_listDelNode(m, v26, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return
L24:
	;
	goto L2
L25:
	;
	v260 = F_zmalloc_usable(m, v65+int32(16), v14+int32(12))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L23
	} else {
		goto L78
	}
L26:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	if v73 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if base.Ui32(int32(65535)) < base.Ui32(v76) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if base.Ui32(v79-v76) < base.Ui32(v65) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v82 == int32(1) {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v86 = v73 + int32(13)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v86 == v87 {
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+12)))
	if v89&int32(1) != 0 {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	v92 = v86 + v65
	if v92 == v86 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v65 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L34:
	;
	goto L33
L35:
	;
	v96 = v76 + v92
	if base.Ui32(int32(0)-v76<<(uint(int32(1))%32)) < base.Ui32(v86-v96) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v106 = (v86 ^ v92) & int32(3)
	if base.Ui32(v86) <= base.Ui32(v92) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v103 = F___memcpy(m, v92, v86, v76)
	mBase = m.M
	goto L33
L38:
	;
	if v212 == int32(0) {
		goto L34
	} else {
		goto L70
	}
L39:
	;
	if base.Ui32(v190) <= base.Ui32(int32(3)) {
		v211 = v189
		v212 = v190
		v213 = v191
		goto L38
	} else {
		goto L66
	}
L40:
	;
	if v106 != 0 {
		v172 = v76
		goto L50
	} else {
		goto L51
	}
L41:
	;
	if v106 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v92&int32(3) != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v211 = v86
	v212 = v76
	v213 = v92
	goto L38
L44:
	;
	v113 = v86
	v114 = v76
	v115 = v92
	goto L46
L45:
	;
	v189 = v86
	v190 = v76
	v191 = v92
	goto L39
L46:
	;
	if v114 == int32(0) {
		goto L34
	} else {
		goto L48
	}
L48:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v119)
	v121 = int32(1)
	v122 = v113 + v121
	v124 = v114 + int32(-1)
	v126 = v115 + v121
	if v126&int32(3) == int32(0) {
		v189 = v122
		v190 = v124
		v191 = v126
		goto L39
	} else {
		goto L49
	}
L49:
	;
	v113 = v122
	v114 = v124
	v115 = v126
	goto L46
L50:
	;
	if v172 == int32(0) {
		goto L34
	} else {
		goto L62
	}
L51:
	;
	if v96&int32(3) == int32(0) {
		v152 = v76
		goto L52
	} else {
		goto L53
	}
L52:
	;
	if base.Ui32(v152) <= base.Ui32(int32(3)) {
		v172 = v152
		goto L50
	} else {
		goto L58
	}
L53:
	;
	v137 = v76
	goto L54
L54:
	;
	if v137 == int32(0) {
		goto L34
	} else {
		goto L56
	}
L55:
	;
	v152 = v143
	goto L52
L56:
	;
	v143 = v137 + int32(-1)
	v144 = v92 + v143
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+v143))))
	*(*uint8)(unsafe.Add(mBase, uint32(v144))) = uint8(v146)
	if v144&int32(3) != 0 {
		v137 = v143
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v159 = v152
	goto L59
L59:
	;
	v163 = v159 + int32(-4)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v86+v163)))
	*(*int32)(unsafe.Add(mBase, uint32(v92+v163))) = v166
	if base.Ui32(int32(3)) < base.Ui32(v163) {
		v159 = v163
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v172 = v163
	goto L50
L61:
	;
	goto L60
L62:
	;
	v179 = v172
	goto L63
L63:
	;
	v183 = v179 + int32(-1)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+v183))))
	*(*uint8)(unsafe.Add(mBase, uint32(v92+v183))) = uint8(v186)
	if v183 != 0 {
		v179 = v183
		goto L63
	} else {
		goto L65
	}
L65:
	;
	goto L34
L66:
	;
	v196 = v189
	v197 = v190
	v198 = v191
	goto L67
L67:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v200
	v202 = int32(4)
	v203 = v196 + v202
	v205 = v198 + v202
	v207 = v197 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v207) {
		v196 = v203
		v197 = v207
		v198 = v205
		goto L67
	} else {
		goto L69
	}
L68:
	;
	v211 = v203
	v212 = v207
	v213 = v205
	goto L38
L69:
	;
	goto L68
L70:
	;
	v218 = v211
	v219 = v212
	v220 = v213
	goto L71
L71:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	*(*uint8)(unsafe.Add(mBase, uint32(v220))) = uint8(v222)
	v224 = int32(1)
	v229 = v219 + int32(-1)
	if v229 != 0 {
		v218 = v218 + v224
		v219 = v229
		v220 = v220 + v224
		goto L71
	} else {
		goto L73
	}
L72:
	;
	goto L34
L73:
	;
	goto L72
L74:
	;
	v245 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+264)) = v245 + base.I64_extend_i32_u(v65)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v76 + v65
	F_listDelNode(m, v26, l1)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L23
	} else {
		goto L77
	}
L75:
	;
	goto L74
L76:
	;
	v243 = F__emscripten_memcpy_bulkmem(m, v86, v64, v65)
	mBase = m.M
	goto L75
L77:
	;
	goto L2
L78:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v260)+4)) = v65
	v265 = v262 + int32(-16)
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v265
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+12)))
	v269 = v267 & int32(254)
	*(*uint8)(unsafe.Add(mBase, uint32(v260)+12)) = uint8(v269)
	if v65 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v277 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+264)) = v277 + base.I64_extend_i32_u(v65)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v260
	if v21 == int64(-1) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L79
L81:
	;
	v275 = F__emscripten_memcpy_bulkmem(m, v260+int32(13), v64, v65)
	mBase = m.M
	goto L80
L82:
	;
	v284 = int32(160)
	goto L84
L83:
	;
	v284 = int32(384)
	goto L84
L84:
	;
	v285 = l0 + v284
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v285)))
	*(*int64)(unsafe.Add(mBase, uint32(v285))) = v286 + base.I64_extend_i32_u(v265)
	v291 = F_closeClientOnOutputBufferLimitReached(m, l0, int32(1))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L23
	} else {
		goto L85
	}
L85:
	;
	goto L2
L86:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
