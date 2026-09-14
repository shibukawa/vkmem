package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_getKeysFreeResult(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	if l0 == int32(0) {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v5 == l0+int32(12) {
			return
		} else {
			F_valkey_free(m, v5)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_getKeysPrepareResult(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v6 != 0 {
		v11 = v6
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if l1 <= v12 {
			v39 = v11
			return v39
		} else {
			v15 = l1 << (uint(int32(3)) % 32)
			v17 = l0 + int32(12)
			if v11 == v17 {
				v24 = F_valkey_malloc(m, v15)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v27 == int32(0) {
						v36 = v24
					} else {
						v31 = v27 << (uint(int32(3)) % 32)
						if v31 == int32(0) {
						} else {
							v34 = F__emscripten_memcpy_bulkmem(m, v24, v17, v31)
							mBase = m.M
						}
						v36 = v24
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
					v39 = v36
					return v39
				}
			} else {
				v19 = F_valkey_realloc(m, v11, v15)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v19
					v36 = v19
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
					v39 = v36
					return v39
				}
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 != 0 {
			F__serverAssert(m, int32(_a512), int32(_a474), int32(2292))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v9 = l0 + int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
			v11 = v9
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if l1 <= v12 {
				v39 = v11
				return v39
			} else {
				v15 = l1 << (uint(int32(3)) % 32)
				v17 = l0 + int32(12)
				if v11 == v17 {
					v24 = F_valkey_malloc(m, v15)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v27 == int32(0) {
							v36 = v24
						} else {
							v31 = v27 << (uint(int32(3)) % 32)
							if v31 == int32(0) {
							} else {
								v34 = F__emscripten_memcpy_bulkmem(m, v24, v17, v31)
								mBase = m.M
							}
							v36 = v24
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
						v39 = v36
						return v39
					}
				} else {
					v19 = F_valkey_realloc(m, v11, v15)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v19
						v36 = v19
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
						v39 = v36
						return v39
					}
				}
			}
		}
	}
}
func F_getKeysSubcommandImpl(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v91 int64
	_ = v91
	var v97 int64
	_ = v97
	var v104 int64
	_ = v104
	var v111 int64
	_ = v111
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int64
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v154 int64
	_ = v154
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int64
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v273 int32
	_ = v273
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(2064)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[822]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+2060)) = v3
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v19 = F_objectGetVal(m, v18)
	mBase = m.M
	v22 = F_hashtableFind(m, v14, v19, v10+int32(2060))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+2060))
	if v22 == int32(0) {
		v42 = v24
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v10 + int32(2064)
	return
L4:
	;
	v59 = int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v51)+84))
	if v60 != 0 {
		v162 = v59
		goto L16
	} else {
		goto L17
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(1099511627776)
	if v42 != 0 {
		v51 = v42
		goto L4
	} else {
		goto L12
	}
L6:
	;
	if v12 == int32(3) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v35 = F_objectGetVal(m, v34)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v24)+200))
	v39 = F_hashtableFind(m, v38, v35, v10)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L11
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(1099511627776)
	v51 = v24
	goto L4
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+200))
	if v29 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v42 = v41
	goto L5
L12:
	;
	F_addReplyError(m, l0, int32(_a1372))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L3
L14:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v51)+52))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v184 = v182 + int32(-2)
	if base.B2i32(int32(0) < v179)&base.B2i32(v179 != v184) != 0 {
		goto L35
	} else {
		goto L36
	}
L15:
	;
	if v175 != 0 {
		goto L14
	} else {
		goto L32
	}
L16:
	;
	v175 = v162
	goto L15
L17:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+58)))
	if v61&int32(32) != 0 {
		v162 = v59
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
	if int32(1) <= v64 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v69 = v64 & int32(3)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v51)+76))
	if base.Ui32(int32(4)) <= base.Ui32(v64) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v175 = int32(0)
	goto L15
L21:
	;
	if v69 == int32(0) {
		v154 = v126
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v78 = int32(0)
	v81 = v78
	v83 = v78
	v86 = int64(0)
	goto L24
L23:
	;
	v121 = int32(0)
	v126 = int64(0)
	goto L21
L24:
	;
	v88 = int32(48)
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v70+v81*v88)+8))
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v70+(v81|int32(1))*v88)+8))
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v70+(v81|int32(2))*v88)+8))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v70+(v81|int32(3))*v88)+8))
	v115 = v86 | (v91&v97&v104&v111 ^ int64(-1))
	v116 = int32(4)
	v117 = v81 + v116
	v119 = v83 + v116
	if v119 != v64&int32(2147483644) {
		v81 = v117
		v83 = v119
		v86 = v115
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v121 = v117
	v126 = v115
	goto L21
L26:
	;
	goto L25
L27:
	;
	v162 = int32(base.Ui32(base.I32_wrap_i64(v154))>>(uint(int32(8))%32)) & int32(1)
	goto L16
L28:
	;
	v130 = v121
	v134 = int32(0)
	v135 = v126
	goto L29
L29:
	;
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v70+v130*int32(48))+8))
	v143 = v135 | (v140 ^ int64(-1))
	v144 = int32(1)
	v147 = v134 + v144
	if v147 != v69 {
		v130 = v130 + v144
		v134 = v147
		v135 = v143
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v154 = v143
	goto L27
L31:
	;
	goto L30
L32:
	;
	F_addReplyError(m, l0, int32(_a1373))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L3
L34:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v197 = F_getKeysFromCommandWithSpecs(m, v51, v193+int32(8), v184, int32(0), v10)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L41
	}
L35:
	;
	F_addReplyError(m, l0, int32(_a1374))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	if int32(0)-v179 <= v184 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	goto L3
L39:
	;
	F_getKeysFreeResult(m, v10)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L59
	}
L40:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	F_addReplyArrayLen(m, l0, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L47
	}
L41:
	;
	if v197 != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+58)))
	if v199&int32(8) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_addReplyError(m, l0, int32(_a1375))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	F_addReplyArrayLen(m, l0, int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L39
L46:
	;
	goto L39
L47:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v213 < int32(1) {
		goto L39
	} else {
		goto L48
	}
L48:
	;
	v221 = int32(0)
	goto L49
L49:
	;
	if l1 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L39
L51:
	;
	v262 = v221 + int32(1)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v262 < v263 {
		v221 = v262
		goto L49
	} else {
		goto L58
	}
L52:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+v221<<(uint(int32(3))%32))))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v224+v229<<(uint(int32(2))%32)+int32(8))))
	F_addReplyBulk(m, l0, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v244 = v221 << (uint(int32(3)) % 32)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242+v244)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v241+v246<<(uint(int32(2))%32)+int32(8))))
	F_addReplyBulk(m, l0, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v257 = int64(*(*int32)(unsafe.Add(mBase, uint32(v255+v244)+4)))
	F_addReplyFlagsForKeyArgs(m, l0, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L51
L58:
	;
	goto L50
L59:
	;
	goto L3
}
func F_getKeysUsingLegacyRangeSpec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v140 int32
	_ = v140
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a512), int32(_a474), int32(2292))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L16
	} else {
		goto L38
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v124
	m.G0 = v15 + int32(16)
	return v124
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v22 != 0 {
		v27 = v22
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v124 = int32(0)
	goto L2
L5:
	;
	v28 = int32(0)
	if v19 < v28 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v23 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v25 = l3 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v25
	v27 = v25
	goto L5
L8:
	;
	v71 = int32(0)
	if v36 < v20 {
		v124 = v71
		goto L2
	} else {
		goto L23
	}
L9:
	;
	v31 = v28
	goto L11
L10:
	;
	v31 = v20
	goto L11
L11:
	;
	v32 = v31 + v19
	v36 = v32>>(uint(int32(31))%32)&l2 + v32
	v37 = v36 - v20
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v37 < v38 {
		v69 = v27
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v41 = v37 + int32(1)
	v43 = v41 << (uint(int32(3)) % 32)
	v45 = l3 + int32(12)
	if v27 == v45 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v41
	v69 = v65
	goto L8
L14:
	;
	v52 = F_valkey_malloc(m, v43)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L16
	} else {
		goto L18
	}
L15:
	;
	v47 = F_valkey_realloc(m, v27, v43)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v47
	v65 = v47
	goto L13
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v55 == int32(0) {
		v65 = v52
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v59 = v55 << (uint(int32(3)) % 32)
	if v59 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v65 = v52
	goto L13
L21:
	;
	goto L20
L22:
	;
	v62 = F__emscripten_memcpy_bulkmem(m, v52, v45, v59)
	mBase = m.M
	goto L21
L23:
	;
	v78 = v71
	v79 = v20
	goto L24
L24:
	;
	if l2 <= v79 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v124 = v116
	goto L2
L26:
	;
	v111 = v69 + v78<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v79
	v116 = v78 + int32(1)
	v117 = v79 + v21
	if v117 <= v36 {
		v78 = v116
		v79 = v117
		goto L24
	} else {
		goto L37
	}
L27:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v87&int32(8) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if v20 <= v79 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v93 = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v94 < v93 {
		v124 = v93
		goto L2
	} else {
		goto L32
	}
L31:
	;
	v124 = int32(0)
	goto L2
L32:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[237]))
	if v100 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v101 = int32(_a513)
	goto L35
L34:
	;
	v101 = int32(_a184)
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v101
	F__serverPanic_1(m, int32(_a474), int32(2614), int32(_a514), v15)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L16
	} else {
		goto L36
	}
L36:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	goto L25
L38:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
