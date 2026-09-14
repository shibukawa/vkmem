package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_setTypeConvertAndExpand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v286 int32
	_ = v286
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12&int32(15) != int32(2) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v286
L2:
	;
	v286 = int32(0)
	goto L1
L3:
	;
	v221 = F_setTypeInitIterator(m, l0)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L10
	} else {
		goto L81
	}
L4:
	;
	F__serverPanic_1(m, int32(_a_F_setTypeConvertAndExpand_0), int32(548), int32(_a_F_setTypeConvertAndExpand_1), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L10
	} else {
		goto L80
	}
L5:
	;
	F__serverAssertWithInfo(m, int32(0), l0, int32(_a_F_setTypeConvertAndExpand_2), int32(_a_F_setTypeConvertAndExpand_0), int32(498))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L10
	} else {
		goto L79
	}
L6:
	;
	if int32(base.Ui32(v12)>>(uint(int32(4))%32))&int32(15) == l1 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	switch l1 + int32(-2) {
	case 0:
		goto L9
	default:
		goto L4
	case 9:
		goto L8
	}
L8:
	;
	v39 = l2 << (uint(int32(1)) % 32)
	if v12&int32(240) != int32(96) {
		v143 = v39
		goto L18
	} else {
		goto L19
	}
L9:
	;
	v25 = F_hashtableCreate(m, int32(_a_F_setTypeConvertAndExpand_3))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if l3 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v33 = F_hashtableTryExpand(m, v25, l2)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	v31 = F_hashtableExpand(m, v25, l2)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L3
L15:
	;
	if v33 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	F_hashtableRelease(m, v25)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v286 = int32(-1)
	goto L1
L18:
	;
	v145 = F_lpNew(m, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L10
	} else {
		goto L59
	}
L19:
	;
	v44 = F_setTypeSize(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	if v44 == int32(0) {
		v143 = v39
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v48 = F_objectGetVal(m, l0)
	mBase = m.M
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	switch v49 + int32(-4) {
	case 0:
		goto L24
	default:
		goto L23
	case 4:
		goto L25
	}
L22:
	;
	if base.Ui64(int64(128)) <= base.Ui64(v55) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v54 = int64(*(*int16)(unsafe.Add(mBase, uint32(v48)+8)))
	v55 = v54
	goto L22
L24:
	;
	v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(v48)+8)))
	v55 = v53
	goto L22
L25:
	;
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
	v55 = v52
	goto L22
L26:
	;
	v86 = F_objectGetVal(m, l0)
	mBase = m.M
	v91 = v86 + int32(8)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v94 = v92 + int32(-1)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	switch v95 + int32(-4) {
	case 0:
		goto L41
	default:
		goto L40
	case 4:
		goto L42
	}
L27:
	;
	v85 = v82*l2 + int32(7)
	goto L26
L28:
	;
	if base.Ui64(int64(8192)) <= base.Ui64(v55+int64(4096)) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v82 = int32(2)
	goto L27
L30:
	;
	if base.Ui64(int64(65536)) <= base.Ui64(v55+int64(32768)) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v82 = int32(3)
	goto L27
L32:
	;
	if base.Ui64(int64(16777216)) <= base.Ui64(v55+int64(8388608)) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v82 = int32(4)
	goto L27
L34:
	;
	if base.Ui64(v55+int64(2147483648)) < base.Ui64(int64(4294967296)) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v82 = int32(5)
	goto L27
L36:
	;
	v81 = int32(6)
	goto L38
L37:
	;
	v81 = int32(10)
	goto L38
L38:
	;
	v82 = v81
	goto L27
L39:
	;
	if base.Ui64(int64(128)) <= base.Ui64(v110) {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v109 = int64(*(*int16)(unsafe.Add(mBase, uint32(v91+v94<<(uint(int32(1))%32)))))
	v110 = v109
	goto L39
L41:
	;
	v105 = int64(*(*int32)(unsafe.Add(mBase, uint32(v91+v94<<(uint(int32(2))%32)))))
	v110 = v105
	goto L39
L42:
	;
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v86+v92<<(uint(int32(3))%32))))
	v110 = v101
	goto L39
L43:
	;
	if base.Ui32(v140) < base.Ui32(v85) {
		goto L56
	} else {
		goto L57
	}
L44:
	;
	v140 = v137*l2 + int32(7)
	goto L43
L45:
	;
	if base.Ui64(int64(8192)) <= base.Ui64(v110+int64(4096)) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v137 = int32(2)
	goto L44
L47:
	;
	if base.Ui64(int64(65536)) <= base.Ui64(v110+int64(32768)) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v137 = int32(3)
	goto L44
L49:
	;
	if base.Ui64(int64(16777216)) <= base.Ui64(v110+int64(8388608)) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v137 = int32(4)
	goto L44
L51:
	;
	if base.Ui64(v110+int64(2147483648)) < base.Ui64(int64(4294967296)) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v137 = int32(5)
	goto L44
L53:
	;
	v136 = int32(6)
	goto L55
L54:
	;
	v136 = int32(10)
	goto L55
L55:
	;
	v137 = v136
	goto L44
L56:
	;
	v142 = v85
	goto L58
L57:
	;
	v142 = v140
	goto L58
L58:
	;
	v143 = v142
	goto L18
L59:
	;
	v147 = F_setTypeInitIterator(m, l0)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L10
	} else {
		goto L61
	}
L60:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	if v189 != int32(2) {
		goto L73
	} else {
		goto L74
	}
L61:
	;
	v153 = F_setTypeNext(m, v147, v10+int32(4), v10, v10+int32(8))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	if v153 == int32(-1) {
		v183 = v145
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v158 = v145
	goto L64
L64:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v164 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v183 = v173
	goto L60
L66:
	;
	v178 = F_setTypeNext(m, v147, v10+int32(4), v10, v10+int32(8))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L10
	} else {
		goto L71
	}
L67:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v171 = F_lpAppendInteger(m, v158, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L10
	} else {
		goto L70
	}
L68:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v168 = F_lpAppend(m, v158, v164, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	v173 = v168
	goto L66
L70:
	;
	v173 = v171
	goto L66
L71:
	;
	if v178 != int32(-1) {
		v158 = v173
		goto L64
	} else {
		goto L72
	}
L72:
	;
	goto L65
L73:
	;
	F_valkey_free(m, v147)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L10
	} else {
		goto L76
	}
L74:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v147)+12))
	F_hashtableReleaseIterator(m, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	F_freeSetObject(m, l0)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L10
	} else {
		goto L77
	}
L77:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v199&int32(-241) | int32(176)
	F_objectSetVal(m, l0, v183)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L10
	} else {
		goto L78
	}
L78:
	;
	goto L2
L79:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	goto L83
L82:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v259 != int32(2) {
		goto L96
	} else {
		goto L97
	}
L83:
	;
	v234 = F_setTypeNext(m, v221, v10+int32(4), v10, v10+int32(8))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L10
	} else {
		goto L85
	}
L84:
	;
	F__serverAssert(m, int32(_a_F_setTypeConvertAndExpand_4), int32(_a_F_setTypeConvertAndExpand_0), int32(515))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L10
	} else {
		goto L95
	}
L85:
	;
	if v234 == int32(-1) {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v238 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v247 == int32(0) {
		goto L82
	} else {
		goto L92
	}
L88:
	;
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v245 = F_sdsfromlonglong(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L10
	} else {
		goto L91
	}
L89:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v242 = F_sdsnewlen(m, v238, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	v247 = v242
	goto L87
L91:
	;
	v247 = v245
	goto L87
L92:
	;
	v250 = F_hashtableAdd(m, v25, v247)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	if v250 != 0 {
		goto L83
	} else {
		goto L94
	}
L94:
	;
	goto L84
L95:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_valkey_free(m, v221)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L10
	} else {
		goto L99
	}
L97:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	F_hashtableReleaseIterator(m, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	F_freeSetObject(m, l0)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L10
	} else {
		goto L100
	}
L100:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v269&int32(-241) | int32(32)
	F_objectSetVal(m, l0, v25)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L10
	} else {
		goto L101
	}
L101:
	;
	goto L2
}
func F_setTypePopRandom(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11&int32(240) != int32(176) {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
		v61 = F_setTypeRandomElement(m, l0, v9+int32(12), v9+int32(8), v9)
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return int32(0)
		} else {
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			if v63 == int32(0) {
				v69 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
				v70 = F_createStringObjectFromLongLong(m, v69)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					v73 = v70
					v74 = v72
					v75 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
					v78 = F_setTypeRemoveAux(m, l0, v63, v74, v75, base.B2i32(v61 == int32(2)))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						v81 = v73
						m.G0 = v9 + int32(16)
						return v81
					}
				}
			} else {
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				v67 = F_createStringObject_1(m, v63, v66)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					v73 = v67
					v74 = v66
					v75 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
					v78 = F_setTypeRemoveAux(m, l0, v63, v74, v75, base.B2i32(v61 == int32(2)))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						v81 = v73
						m.G0 = v9 + int32(16)
						return v81
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
		v18 = F_objectGetVal(m, l0)
		mBase = m.M
		v19 = F_objectGetVal(m, l0)
		mBase = m.M
		v20 = F_lpFirst(m, v19)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v28 = F_lpNextRandom(m, v18, v20, v9+int32(12), int32(1), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
				v36 = F_lpGetValue(m, v28, v9+int32(8), v9)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					if v36 == int32(0) {
						v43 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
						v44 = F_createStringObjectFromLongLong(m, v43)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = v44
							v47 = F_objectGetVal(m, l0)
							mBase = m.M
							v49 = F_lpDelete(m, v47, v28, int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_objectSetVal(m, l0, v49)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v81 = v46
									m.G0 = v9 + int32(16)
									return v81
								}
							}
						}
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						v41 = F_createStringObject_1(m, v36, v40)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v46 = v41
							v47 = F_objectGetVal(m, l0)
							mBase = m.M
							v49 = F_lpDelete(m, v47, v28, int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_objectSetVal(m, l0, v49)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v81 = v46
									m.G0 = v9 + int32(16)
									return v81
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_typeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v8 = F_lookupKey(m, v4, v6, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		if v8 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v15 = v13 & int32(15)
			if base.Ui32(int32(7)) <= base.Ui32(v15) {
				F__serverAssert(m, int32(_a_F_typeCommand_0), int32(_a_F_typeCommand_1), int32(1140))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if v15 != int32(5) {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(2))%32))+uint32(_c_F_typeCommand[0])))
					F_addReplyStatus(m, l0, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						return
					}
				} else {
					v20 = F_objectGetVal(m, v8)
					mBase = m.M
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					F_addReplyStatus(m, l0, v21+int32(84))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			F_addReplyStatus(m, l0, int32(_a_F_typeCommand_2))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		}
	}
}
