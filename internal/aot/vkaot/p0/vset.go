package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_vsetAddEntry(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int64
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = m.T0[l1].(func(*base.Module, int32) int64)(m, l2)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v22 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	F__serverAssert(m, int32(_a1674), int32(_a1670), int32(600))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L82
	}
L4:
	;
	F__serverAssert(m, int32(_a1675), int32(_a1670), int32(816))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L81
	}
L5:
	;
	F__serverAssert(m, int32(_a1675), int32(_a1670), int32(816))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L80
	}
L6:
	;
	F__serverAssert(m, int32(_a1674), int32(_a1670), int32(600))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L79
	}
L7:
	;
	F__serverAssert(m, int32(_a1674), int32(_a1670), int32(600))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L78
	}
L8:
	;
	F__serverAssert(m, int32(_a1675), int32(_a1670), int32(816))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L77
	}
L9:
	;
	F__serverAssert(m, int32(_a1676), int32(_a1670), int32(1784))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L76
	}
L10:
	;
	if v22 == int32(-1) {
		v268 = l2
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v268
	m.G0 = v16 + int32(16)
	return int32(1)
L12:
	;
	if v22&int32(1) != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v264 = F_insertToBucket_RAX(m, l1, v22, l2, v18)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L75
	}
L14:
	;
	F__serverPanic_1(m, int32(_a1670), int32(1830), int32(_a1677), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L74
	}
L15:
	;
	v72 = int32(0)
	v74 = v22 & int32(-8)
	if v74 == v72 {
		v249 = v72
		goto L32
	} else {
		goto L33
	}
L16:
	;
	v34 = F_valkey_malloc(m, int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	switch v22&int32(6) + int32(-2) {
	case 0:
		goto L15
	default:
		goto L14
	case 4:
		goto L13
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v34))) = int64(17179869184)
	v38 = m.T0[l1].(func(*base.Module, int32) int64)(m, v22)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	if v65 == int32(0) {
		goto L8
	} else {
		goto L31
	}
L20:
	;
	v53 = int32(0)
	v55 = F_pvInsertAt(m, v34, l2, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L28
	}
L21:
	;
	if v18 <= v38 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v41 = int32(0)
	v43 = F_pvInsertAt(m, v34, v22, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v51 = F_pvInsertAt(m, v43, l2, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	if v43 == int32(0) {
		v50 = v41
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v50 = v47 & int32(1073741823)
	goto L23
L26:
	;
	v65 = v51
	goto L19
L27:
	;
	v63 = F_pvInsertAt(m, v55, v22, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	if v55 == int32(0) {
		v62 = v53
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v62 = v59 & int32(1073741823)
	goto L27
L30:
	;
	v65 = v63
	goto L19
L31:
	;
	v268 = v65 | int32(2)
	goto L11
L32:
	;
	v255 = F_insertToBucket_VECTOR(m, v22, l2, v249)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L73
	}
L33:
	;
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
	if v77&int64(1073741823) != int64(127) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v201 = base.I32_wrap_i64(v77) & int32(1073741823)
	if v201 == int32(0) {
		v249 = v72
		goto L32
	} else {
		goto L58
	}
L35:
	;
	v82 = F_raxNew(m)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v84&int32(1073741823) == int32(0) {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	v90 = m.T0[l1].(func(*base.Module, int32) int64)(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
	if base.Ui64(v92&int64(1073741823)) <= base.Ui64(int64(126)) {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v74)+512))
	v98 = m.T0[l1].(func(*base.Module, int32) int64)(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	if v82 == int32(0) {
		goto L4
	} else {
		goto L49
	}
L41:
	;
	v100 = int64(-8192)
	v103 = v90 & v100
	if v98&v100 != v103 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v107 = v103 + int64(8192)
	v108 = int64(56)
	v110 = int64(65280)
	v112 = int64(40)
	v115 = int64(16711680)
	v117 = int64(24)
	v119 = int64(4278190080)
	v121 = int64(8)
	if v103 == int64(9223372036854767616) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v145 = int64(-129)
	goto L45
L44:
	;
	v145 = v107<<(uint(v108)%64) | v107&v110<<(uint(v112)%64) | (v107&v115<<(uint(v117)%64) | v107&v119<<(uint(v121)%64)) | (int64(base.Ui64(v107)>>(uint(v121)%64))&v119 | int64(base.Ui64(v107)>>(uint(v117)%64))&v115 | (int64(base.Ui64(v107)>>(uint(v112)%64))&v110 | int64(base.Ui64(v107)>>(uint(v108)%64))))
	goto L45
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v145
	v147 = int32(8)
	v151 = F_raxInsert(m, v82, v16+v147, v147, v22, int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v82 == int32(0) {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	v157 = F_insertToBucket_RAX(m, l1, v82|int32(6), l2, v18)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v268 = v157
	goto L11
L49:
	;
	v171 = int32(0)
	v173 = v82 | int32(6)
	goto L50
L50:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if base.Ui32(v179&int32(1073741823)) <= base.Ui32(v171) {
		goto L3
	} else {
		goto L52
	}
L51:
	;
	F_valkey_free(m, v74)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L56
	}
L52:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(8)+v171<<(uint(int32(2))%32))))
	v187 = m.T0[l1].(func(*base.Module, int32) int64)(m, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v189 = F_insertToBucket_RAX(m, l1, v173, v186, v187)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v192 = v171 + int32(1)
	if v192 != int32(127) {
		v171 = v192
		v173 = v189
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	v197 = F_insertToBucket_RAX(m, l1, v189, l2, v18)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v268 = v197
	goto L11
L58:
	;
	v213 = v201
	v214 = int32(0)
	goto L59
L59:
	;
	v222 = int32(base.Ui32(v213+v214) >> (uint(int32(1)) % 32))
	v225 = v74 + int32(8) + v222<<(uint(int32(2))%32)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v227 = m.T0[l1].(func(*base.Module, int32) int64)(m, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L63
	}
L60:
	;
	v249 = v239
	goto L32
L61:
	;
	if base.Ui32(v239) < base.Ui32(v238) {
		v213 = v238
		v214 = v239
		goto L59
	} else {
		goto L72
	}
L62:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v233 = m.T0[l1].(func(*base.Module, int32) int64)(m, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	if v227 <= v18 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v238 = v222
	v239 = v214
	goto L61
L65:
	;
	v235 = base.B2i32(v18 == v233)
	if v18 == v233 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v236 = v214
	goto L68
L67:
	;
	v236 = v222 + int32(1)
	goto L68
L68:
	;
	if v18 == v233 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v237 = v222
	goto L71
L70:
	;
	v237 = v213
	goto L71
L71:
	;
	v238 = v237
	v239 = v236
	goto L61
L72:
	;
	goto L60
L73:
	;
	v268 = v255
	goto L11
L74:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	v268 = v264
	goto L11
L76:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
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
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vsetCompareEntries(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, _consts[1128]))
	v8 = m.T0[v7].(func(*base.Module, int32) int64)(m, v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v13 = m.T0[v7].(func(*base.Module, int32) int64)(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return base.B2i32(v13 < v8) - base.B2i32(v8 < v13)
		}
	}
}
func F_vsetResetIterator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	switch v5 + int32(1) {
	case 0:
		v18 = int32(0)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
		switch v19 + int32(1) {
		case 0:
			if v18 == int32(0) {
				return
			} else {
				F_hashtableCleanupIterator(m, l0+int32(304))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					return
				}
			}
		case 1:
			F__serverAssert(m, int32(_a1638), int32(_a1670), int32(781))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		default:
			if v19&int32(7) != int32(6) {
				if v18 == int32(0) {
					return
				} else {
					F_hashtableCleanupIterator(m, l0+int32(304))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				F_raxStop(m, l0)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v18 == int32(0) {
						return
					} else {
						F_hashtableCleanupIterator(m, l0+int32(304))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	case 1:
		F__serverAssert(m, int32(_a1638), int32(_a1670), int32(781))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	default:
		v18 = base.B2i32(v5&int32(7) == int32(4))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
		switch v19 + int32(1) {
		case 0:
			if v18 == int32(0) {
				return
			} else {
				F_hashtableCleanupIterator(m, l0+int32(304))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					return
				}
			}
		case 1:
			F__serverAssert(m, int32(_a1638), int32(_a1670), int32(781))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		default:
			if v19&int32(7) != int32(6) {
				if v18 == int32(0) {
					return
				} else {
					F_hashtableCleanupIterator(m, l0+int32(304))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				F_raxStop(m, l0)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v18 == int32(0) {
						return
					} else {
						F_hashtableCleanupIterator(m, l0+int32(304))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
