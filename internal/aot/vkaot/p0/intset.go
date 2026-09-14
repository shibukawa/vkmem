package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_createIntsetObject(m *base.Module) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_intsetNew(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
		v14 = int32(12)
		v17 = F_zmalloc_usable(m, v14, v6+v14)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v8
			*(*int64)(unsafe.Add(mBase, uint32(v17))) = int64(34359738370)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v22&int32(-241) | int32(96)
			m.G0 = v6 + int32(16)
			return v17
		}
	}
}
func F_intsetBlobLen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	return v2*v3 + int32(8)
}
func F_intsetRandom(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int64
	_ = v28
	var v33 int64
	_ = v33
	var v38 int64
	_ = v38
	var v46 int32
	_ = v46
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v4 == int32(0) {
		F__serverAssert(m, int32(_a659), int32(_a658), int32(263))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int64(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = l0 + int32(8)
		v10 = int32(0)
		v12 = *(*int64)(unsafe.Add(mBase, _consts[298]))
		v16 = v12*int64(6364136223846793005) + int64(1)
		*(*int64)(unsafe.Add(mBase, _consts[298])) = v16
		v21 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v16)>>(uint(int64(33))%64))), v4)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		switch v22 + int32(-4) {
		case 0:
			v33 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8+v21<<(uint(int32(2))%32)))))
			return v33
		default:
			v38 = int64(*(*int16)(unsafe.Add(mBase, uint32(v8+v21<<(uint(int32(1))%32)))))
			return v38
		case 4:
			v28 = *(*int64)(unsafe.Add(mBase, uint32(v8+v21<<(uint(int32(3))%32))))
			return v28
		}
	}
}
func F_intsetRemove(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int64
	_ = v58
	var v62 int64
	_ = v62
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v363 int64
	_ = v363
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v391 int32
	_ = v391
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui64(l1+int64(-32768)) < base.Ui64(int64(-65536)) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L1
L3:
	;
	F__serverAssert(m, int32(_a657), int32(_a658), int32(108))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L98
	} else {
		goto L100
	}
L4:
	;
	m.G0 = v15 + int32(16)
	return v374
L5:
	;
	v31 = int32(4)
	goto L7
L6:
	;
	v31 = int32(2)
	goto L7
L7:
	;
	if base.Ui64(l1+int64(-2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v34 = int32(8)
	goto L10
L9:
	;
	v34 = v31
	goto L10
L10:
	;
	if base.Ui32(v25) < base.Ui32(v34) {
		v374 = l0
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v37 = v15 + int32(12)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v43 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	if v158 == int32(0) {
		v374 = l0
		goto L4
	} else {
		goto L47
	}
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v148
	v158 = v149
	goto L13
L15:
	;
	v142 = int32(0)
	if v37 == v142 {
		v158 = v142
		goto L13
	} else {
		goto L46
	}
L16:
	;
	v137 = int32(0)
	goto L15
L17:
	;
	v47 = v43 + int32(-1)
	v49 = l0 + int32(8)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v50&int32(255) + int32(-4) {
	case 0:
		goto L20
	default:
		goto L19
	case 4:
		goto L21
	}
L18:
	;
	if v67 < l1 {
		v137 = v43
		goto L15
	} else {
		goto L22
	}
L19:
	;
	v66 = int64(*(*int16)(unsafe.Add(mBase, uint32(v49+v47<<(uint(int32(1))%32)))))
	v67 = v66
	goto L18
L20:
	;
	v62 = int64(*(*int32)(unsafe.Add(mBase, uint32(v49+v47<<(uint(int32(2))%32)))))
	v67 = v62
	goto L18
L21:
	;
	v58 = *(*int64)(unsafe.Add(mBase, uint32(l0+v43<<(uint(int32(3))%32))))
	v67 = v58
	goto L18
L22:
	;
	switch v50&int32(255) + int32(-4) {
	case 0:
		goto L25
	default:
		goto L24
	case 4:
		goto L26
	}
L23:
	;
	if l1 < v76 {
		goto L16
	} else {
		goto L27
	}
L24:
	;
	v75 = int64(*(*int16)(unsafe.Add(mBase, uint32(v49))))
	v76 = v75
	goto L23
L25:
	;
	v74 = int64(*(*int32)(unsafe.Add(mBase, uint32(v49))))
	v76 = v74
	goto L23
L26:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
	v76 = v73
	goto L23
L27:
	;
	v78 = int32(0)
	if v78 <= v47 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v128 = base.B2i32(l1 == v127)
	if v37 == int32(0) {
		v158 = v128
		goto L13
	} else {
		goto L42
	}
L29:
	;
	v90 = v78
	v91 = v47
	goto L31
L30:
	;
	v120 = int32(-1)
	v123 = v78
	v127 = int64(-1)
	goto L28
L31:
	;
	v97 = int32(base.Ui32(v91+v90) >> (uint(int32(1)) % 32))
	switch v50&int32(255) + int32(-4) {
	case 0:
		goto L35
	default:
		goto L34
	case 4:
		goto L36
	}
L32:
	;
	v120 = v97
	v123 = v117
	v127 = v110
	goto L28
L33:
	;
	if l1 <= v110 {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v109 = int64(*(*int16)(unsafe.Add(mBase, uint32(v49+v97<<(uint(int32(1))%32)))))
	v110 = v109
	goto L33
L35:
	;
	v105 = int64(*(*int32)(unsafe.Add(mBase, uint32(v49+v97<<(uint(int32(2))%32)))))
	v110 = v105
	goto L33
L36:
	;
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v49+v97<<(uint(int32(3))%32))))
	v110 = v101
	goto L33
L37:
	;
	if v117 <= v118 {
		v90 = v117
		v91 = v118
		goto L31
	} else {
		goto L41
	}
L38:
	;
	if v110 <= l1 {
		v120 = v97
		v123 = v90
		v127 = v110
		goto L28
	} else {
		goto L40
	}
L39:
	;
	v117 = v97 + int32(1)
	v118 = v91
	goto L37
L40:
	;
	v117 = v90
	v118 = v97 + int32(-1)
	goto L37
L41:
	;
	goto L32
L42:
	;
	if l1 == v127 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v131 = v120
	goto L45
L44:
	;
	v131 = v123
	goto L45
L45:
	;
	v148 = v131
	v149 = v128
	goto L14
L46:
	;
	v148 = v137
	v149 = v142
	goto L14
L47:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l2 == int32(0) {
		v170 = v25
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v173 = v164 + int32(-1)
	if base.Ui32(v173) <= base.Ui32(v171) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v170 = v169
	goto L48
L50:
	;
	v363 = base.I64_extend_i32_u(v170) * base.I64_extend_i32_u(v173)
	if base.Ui64(int64(4294967288)) <= base.Ui64(v363) {
		goto L3
	} else {
		goto L97
	}
L51:
	;
	v176 = l0 + int32(8)
	v177 = int32(1)
	v179 = v171 + v177
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v170 + int32(-4) {
	case 0:
		goto L54
	default:
		goto L53
	case 4:
		goto L55
	}
L52:
	;
	v207 = (v180 - v179) << (uint(v205) % 32)
	if v204 == v203 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v197 = int32(1)
	v203 = v176 + v179<<(uint(v197)%32)
	v204 = v176 + v171<<(uint(v197)%32)
	v205 = v177
	goto L52
L54:
	;
	v190 = int32(2)
	v203 = v176 + v179<<(uint(v190)%32)
	v204 = v176 + v171<<(uint(v190)%32)
	v205 = v190
	goto L52
L55:
	;
	v183 = int32(3)
	v203 = v176 + v179<<(uint(v183)%32)
	v204 = v176 + v171<<(uint(v183)%32)
	v205 = v183
	goto L52
L56:
	;
	goto L50
L57:
	;
	goto L56
L58:
	;
	v211 = v207 + v204
	if base.Ui32(int32(0)-v207<<(uint(int32(1))%32)) < base.Ui32(v203-v211) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v221 = (v203 ^ v204) & int32(3)
	if base.Ui32(v203) <= base.Ui32(v204) {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v218 = F___memcpy(m, v204, v203, v207)
	mBase = m.M
	goto L56
L61:
	;
	if v327 == int32(0) {
		goto L57
	} else {
		goto L93
	}
L62:
	;
	if base.Ui32(v305) <= base.Ui32(int32(3)) {
		v326 = v304
		v327 = v305
		v328 = v306
		goto L61
	} else {
		goto L89
	}
L63:
	;
	if v221 != 0 {
		v287 = v207
		goto L73
	} else {
		goto L74
	}
L64:
	;
	if v221 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if v204&int32(3) != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v326 = v203
	v327 = v207
	v328 = v204
	goto L61
L67:
	;
	v228 = v203
	v229 = v207
	v230 = v204
	goto L69
L68:
	;
	v304 = v203
	v305 = v207
	v306 = v204
	goto L62
L69:
	;
	if v229 == int32(0) {
		goto L57
	} else {
		goto L71
	}
L71:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	*(*uint8)(unsafe.Add(mBase, uint32(v230))) = uint8(v234)
	v236 = int32(1)
	v237 = v228 + v236
	v239 = v229 + int32(-1)
	v241 = v230 + v236
	if v241&int32(3) == int32(0) {
		v304 = v237
		v305 = v239
		v306 = v241
		goto L62
	} else {
		goto L72
	}
L72:
	;
	v228 = v237
	v229 = v239
	v230 = v241
	goto L69
L73:
	;
	if v287 == int32(0) {
		goto L57
	} else {
		goto L85
	}
L74:
	;
	if v211&int32(3) == int32(0) {
		v267 = v207
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if base.Ui32(v267) <= base.Ui32(int32(3)) {
		v287 = v267
		goto L73
	} else {
		goto L81
	}
L76:
	;
	v252 = v207
	goto L77
L77:
	;
	if v252 == int32(0) {
		goto L57
	} else {
		goto L79
	}
L78:
	;
	v267 = v258
	goto L75
L79:
	;
	v258 = v252 + int32(-1)
	v259 = v204 + v258
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203+v258))))
	*(*uint8)(unsafe.Add(mBase, uint32(v259))) = uint8(v261)
	if v259&int32(3) != 0 {
		v252 = v258
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v274 = v267
	goto L82
L82:
	;
	v278 = v274 + int32(-4)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v203+v278)))
	*(*int32)(unsafe.Add(mBase, uint32(v204+v278))) = v281
	if base.Ui32(int32(3)) < base.Ui32(v278) {
		v274 = v278
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v287 = v278
	goto L73
L84:
	;
	goto L83
L85:
	;
	v294 = v287
	goto L86
L86:
	;
	v298 = v294 + int32(-1)
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203+v298))))
	*(*uint8)(unsafe.Add(mBase, uint32(v204+v298))) = uint8(v301)
	if v298 != 0 {
		v294 = v298
		goto L86
	} else {
		goto L88
	}
L88:
	;
	goto L57
L89:
	;
	v311 = v304
	v312 = v305
	v313 = v306
	goto L90
L90:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = v315
	v317 = int32(4)
	v318 = v311 + v317
	v320 = v313 + v317
	v322 = v312 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v322) {
		v311 = v318
		v312 = v322
		v313 = v320
		goto L90
	} else {
		goto L92
	}
L91:
	;
	v326 = v318
	v327 = v322
	v328 = v320
	goto L61
L92:
	;
	goto L91
L93:
	;
	v333 = v326
	v334 = v327
	v335 = v328
	goto L94
L94:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	*(*uint8)(unsafe.Add(mBase, uint32(v335))) = uint8(v337)
	v339 = int32(1)
	v344 = v334 + int32(-1)
	if v344 != 0 {
		v333 = v333 + v339
		v334 = v344
		v335 = v335 + v339
		goto L94
	} else {
		goto L96
	}
L95:
	;
	goto L57
L96:
	;
	goto L95
L97:
	;
	v369 = F_valkey_realloc(m, l0, base.I32_wrap_i64(v363)+int32(8))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	return int32(0)
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+4)) = v173
	v374 = v369
	goto L4
L100:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_intsetSearch(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int64
	_ = v24
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v44 int32
	_ = v44
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int64
	_ = v67
	var v71 int64
	_ = v71
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int64
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v124
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v114
	v124 = v115
	goto L1
L3:
	;
	v108 = int32(0)
	if l2 == v108 {
		v124 = v108
		goto L1
	} else {
		goto L34
	}
L4:
	;
	v103 = int32(0)
	goto L3
L5:
	;
	v13 = v9 + int32(-1)
	v15 = l0 + int32(8)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v16&int32(255) + int32(-4) {
	case 0:
		goto L8
	default:
		goto L7
	case 4:
		goto L9
	}
L6:
	;
	if v33 < l1 {
		v103 = v9
		goto L3
	} else {
		goto L10
	}
L7:
	;
	v32 = int64(*(*int16)(unsafe.Add(mBase, uint32(v15+v13<<(uint(int32(1))%32)))))
	v33 = v32
	goto L6
L8:
	;
	v28 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15+v13<<(uint(int32(2))%32)))))
	v33 = v28
	goto L6
L9:
	;
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0+v9<<(uint(int32(3))%32))))
	v33 = v24
	goto L6
L10:
	;
	switch v16&int32(255) + int32(-4) {
	case 0:
		goto L13
	default:
		goto L12
	case 4:
		goto L14
	}
L11:
	;
	if l1 < v42 {
		goto L4
	} else {
		goto L15
	}
L12:
	;
	v41 = int64(*(*int16)(unsafe.Add(mBase, uint32(v15))))
	v42 = v41
	goto L11
L13:
	;
	v40 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15))))
	v42 = v40
	goto L11
L14:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v42 = v39
	goto L11
L15:
	;
	v44 = int32(0)
	if v44 <= v13 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v94 = base.B2i32(l1 == v93)
	if l2 == int32(0) {
		v124 = v94
		goto L1
	} else {
		goto L30
	}
L17:
	;
	v56 = v44
	v57 = v13
	goto L19
L18:
	;
	v86 = int32(-1)
	v89 = v44
	v93 = int64(-1)
	goto L16
L19:
	;
	v63 = int32(base.Ui32(v57+v56) >> (uint(int32(1)) % 32))
	switch v16&int32(255) + int32(-4) {
	case 0:
		goto L23
	default:
		goto L22
	case 4:
		goto L24
	}
L20:
	;
	v86 = v63
	v89 = v83
	v93 = v76
	goto L16
L21:
	;
	if l1 <= v76 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v75 = int64(*(*int16)(unsafe.Add(mBase, uint32(v15+v63<<(uint(int32(1))%32)))))
	v76 = v75
	goto L21
L23:
	;
	v71 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15+v63<<(uint(int32(2))%32)))))
	v76 = v71
	goto L21
L24:
	;
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v15+v63<<(uint(int32(3))%32))))
	v76 = v67
	goto L21
L25:
	;
	if v83 <= v84 {
		v56 = v83
		v57 = v84
		goto L19
	} else {
		goto L29
	}
L26:
	;
	if v76 <= l1 {
		v86 = v63
		v89 = v56
		v93 = v76
		goto L16
	} else {
		goto L28
	}
L27:
	;
	v83 = v63 + int32(1)
	v84 = v57
	goto L25
L28:
	;
	v83 = v56
	v84 = v63 + int32(-1)
	goto L25
L29:
	;
	goto L20
L30:
	;
	if l1 == v93 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v97 = v86
	goto L33
L32:
	;
	v97 = v89
	goto L33
L33:
	;
	v114 = v97
	v115 = v94
	goto L2
L34:
	;
	v114 = v103
	v115 = v108
	goto L2
}
