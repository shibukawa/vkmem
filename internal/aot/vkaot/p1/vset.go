package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_freeVsetBucket(m *base.Module, l0 int32) {
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	switch l0 + int32(1) {
	case 0:
		return
	case 1:
		F__serverAssert(m, int32(_a2513), int32(_a2500), int32(781))
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_abort(m)
			base.Wasm_trap_unreachable()
			for {
			}
		}
	default:
		if l0&int32(1) != 0 {
			return
		} else {
			switch l0 & int32(6) {
			default:
				F__serverPanic_1(m, int32(_a2500), int32(1136), int32(_a2532), int32(0))
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_abort(m)
					base.Wasm_trap_unreachable()
					for {
					}
				}
			case 1, 3, 5:
				base.Wasm_trap_unreachable()
				for {
				}
			case 2:
				v15 = l0 & int32(-8)
				if v15 == int32(0) {
					return
				} else {
					F_valkey_free(m, v15)
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						return
					}
				}
			case 4:
				F_hashtableRelease(m, l0&int32(-8))
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					return
				}
			case 6:
				F_raxFreeWithCallback(m, l0&int32(-8), int32(1129))
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_vsetInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(-1)
	return
}
func F_vsetNext(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v53 int64
	_ = v53
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	v9 = l0 + int32(360)
	v11 = l0 + int32(304)
	goto L2
L1:
	;
	return v241
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	v24 = v19
	goto L4
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = int32(-1)
	v241 = v27
	goto L1
L4:
	;
	v27 = int32(0)
	switch v24 + int32(1) {
	case 0:
		v241 = v27
		goto L1
	case 1:
		goto L7
	default:
		goto L6
	}
L5:
	;
	goto L3
L6:
	;
	if v24&int32(1) != 0 {
		goto L20
	} else {
		goto L21
	}
L7:
	;
	F__serverAssert(m, int32(_a2513), int32(_a2500), int32(781))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	goto L5
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+356)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(l0)+376)) = v234
	v24 = v205
	goto L4
L12:
	;
	v228 = int32(1)
	if v205&v228 != 0 {
		goto L56
	} else {
		goto L57
	}
L13:
	;
	F__serverAssert(m, int32(_a2513), int32(_a2500), int32(781))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L8
	} else {
		goto L55
	}
L14:
	;
	F__serverPanic_1(m, int32(_a2500), int32(2239), int32(_a2533), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L8
	} else {
		goto L54
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	if l1 == int32(0) {
		v241 = int32(1)
		goto L1
	} else {
		goto L53
	}
L17:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	switch v205 + int32(1) {
	case 0:
		v234 = v205
		goto L11
	case 1:
		goto L13
	default:
		goto L12
	}
L18:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
	if v174 != int32(-1) {
		goto L44
	} else {
		goto L45
	}
L19:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
	if v144 != int32(-1) {
		goto L39
	} else {
		goto L40
	}
L20:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
	if v132 != int32(-1) {
		goto L17
	} else {
		goto L36
	}
L21:
	;
	switch v24 & int32(6) {
	default:
		goto L14
	case 1, 3, 5:
		goto L15
	case 2:
		goto L19
	case 4:
		goto L18
	case 6:
		goto L22
	}
L22:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
	if v42 != int32(-1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v70 = F_raxNext(m, l0)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L27
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v24 & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(128)
	v53 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(l0)+296)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l0 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = l0 + int32(168)
	goto L25
L25:
	;
	v66 = int32(0)
	v68 = F_raxSeek(m, l0, int32(_a67), v66, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	if v70 == int32(0) {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	switch v74 + int32(1) {
	case 0:
		v89 = v74
		goto L29
	case 1:
		goto L31
	default:
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+376)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+376)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+356)) = v74
	v96 = int64(56)
	v98 = int64(65280)
	v100 = int64(40)
	v103 = int64(16711680)
	v105 = int64(24)
	v107 = int64(4278190080)
	v109 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+368)) = v92<<(uint(v96)%64) | v92&v98<<(uint(v100)%64) | (v92&v103<<(uint(v105)%64) | v92&v107<<(uint(v109)%64)) | (int64(base.Ui64(v92)>>(uint(v109)%64))&v107 | int64(base.Ui64(v92)>>(uint(v105)%64))&v103 | (int64(base.Ui64(v92)>>(uint(v100)%64))&v98 | int64(base.Ui64(v92)>>(uint(v96)%64))))
	goto L2
L30:
	;
	v83 = int32(1)
	if v74&v83 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	F__serverAssert(m, int32(_a2513), int32(_a2500), int32(781))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v88 = v83
	goto L35
L34:
	;
	v88 = v74 & int32(6)
	goto L35
L35:
	;
	v89 = v88
	goto L29
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v24
	v136 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+376)) = v136
	if l1 == int32(0) {
		v241 = v136
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v24
	return int32(1)
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v153
	v156 = v24 & int32(-8)
	if v156 == int32(0) {
		goto L17
	} else {
		goto L41
	}
L39:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v153 = v150 + int32(1)
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+376)) = int32(2)
	v153 = int32(0)
	goto L38
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if base.Ui32(v159&int32(1073741823)) <= base.Ui32(v153) {
		goto L17
	} else {
		goto L42
	}
L42:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v156+v153<<(uint(int32(2))%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v166
	if l1 == int32(0) {
		v241 = int32(1)
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v166
	return int32(1)
L44:
	;
	v199 = F_hashtableNext(m, v11, v9)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L50
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+376)) = int32(4)
	v180 = v24 & int32(-8)
	v181 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)) = uint8(v181)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v181
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v181)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(-1)
	if v180 == v181 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L44
L47:
	;
	goto L46
L48:
	;
	goto L47
L50:
	;
	if v199 != 0 {
		goto L16
	} else {
		goto L51
	}
L51:
	;
	F_hashtableCleanupIterator(m, v11)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	goto L17
L53:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v211
	return int32(1)
L54:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v233 = v228
	goto L58
L57:
	;
	v233 = v205 & int32(6)
	goto L58
L58:
	;
	v234 = v233
	goto L11
}
func F_vsetRemoveEntry(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v4 = m.T0[l1].(func(*base.Module, int32) int64)(m, l2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_vsetRemoveEntryWithExpiry(m, l0, l1, l2, v4)
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_vsetRemoveEntryWithExpiry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == int32(0) {
		F__serverAssert(m, int32(_a2520), int32(_a2500), int32(1840))
		mBase = m.M
		v89 = m.ExcPending
		if v89 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if v12 == int32(-1) {
			v78 = int32(0)
			m.G0 = v10 + int32(48)
			return v78 & int32(1)
		} else {
			if v12&int32(1) != 0 {
				v69 = base.B2i32(v12 == l2)
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v69)
				if v12 == l2 {
					v72 = int32(-1)
				} else {
					v72 = v12
				}
				v73 = v72
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v73
				v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
				v78 = v76
				m.G0 = v10 + int32(48)
				return v78 & int32(1)
			} else {
				switch v12 & int32(6) {
				default:
					F__serverPanic_1(m, int32(_a2500), int32(1859), int32(_a2521), int32(0))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 1, 3, 5:
					base.Wasm_trap_unreachable()
					for {
					}
				case 2:
					v25 = F_removeFromBucket_VECTOR(m, v12, l2, v10+int32(15), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v73 = v25
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v73
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						v78 = v76
						m.G0 = v10 + int32(48)
						return v78 & int32(1)
					}
				case 4:
					v31 = F_removeFromBucket_HASHTABLE(m, v12, l2, v10+int32(15))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v73 = v31
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v73
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						v78 = v76
						m.G0 = v10 + int32(48)
						return v78 & int32(1)
					}
				case 6:
					*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v12
					*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = int64(0)
					v46 = F_findBucket_2(m, v12&int32(-8), l3, v10+int32(32), v10+int32(20), v10+int32(24), v10+int32(16))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						if v46 == int32(-1) {
							F__serverAssert(m, int32(_a2522), int32(_a2500), int32(1465))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
							v54 = F_removeEntryFromRaxBucket(m, v12, l2, v46, v10+int32(32), v52, v53)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v54)
								F_shrinkRaxBucketIfPossible(m, v10+int32(44), l1)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
									v73 = v61
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v73
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
									v78 = v76
									m.G0 = v10 + int32(48)
									return v78 & int32(1)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_vsetRemoveExpired(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v69 int64
	_ = v69
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v366 int64
	_ = v366
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v407 int32
	_ = v407
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v425 int64
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v459 int64
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v479 int32
	_ = v479
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(336)
	m.G0 = v20
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v23 + int32(1) {
	case 0:
		v479 = v7
		goto L1
	case 1:
		goto L3
	default:
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(336)
	return v479
L2:
	;
	if v23&int32(1) != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F__serverAssert(m, int32(_a2513), int32(_a2500), int32(781))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L6:
	;
	if l4 == int32(0) {
		v479 = v7
		goto L1
	} else {
		goto L120
	}
L7:
	;
	switch v23 & int32(6) {
	default:
		goto L11
	case 1, 3, 5:
		goto L10
	case 2:
		goto L8
	case 4:
		goto L9
	case 6:
		goto L12
	}
L8:
	;
	v385 = v23 & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v385
	v387 = int32(0)
	if v385 == v387 {
		v479 = v387
		goto L1
	} else {
		goto L95
	}
L9:
	;
	v382 = F_vsetBucketRemoveExpired_HASHTABLE(m, l0, l2, l4, l5)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L4
	} else {
		goto L94
	}
L10:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L11:
	;
	F__serverPanic_1(m, int32(_a2500), int32(2135), int32(_a2529), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L4
	} else {
		goto L93
	}
L12:
	;
	v39 = v23 & int32(-8)
	if l4 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
	goto L89
L14:
	;
	v49 = int32(0)
	goto L16
L15:
	;
	v356 = int32(0)
	goto L13
L16:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
	goto L18
L17:
	;
	v356 = v333
	goto L13
L18:
	;
	if v59 == int64(0) {
		v356 = v49
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v63 = v20 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+20)) = int32(128)
	v69 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v63)+12)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v63)+296)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v63)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v20 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+156)) = v20 + int32(196)
	goto L20
L20:
	;
	v84 = int32(0)
	v86 = F_raxSeek(m, v20+int32(28), int32(_a67), v84, v84)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v90 = F_raxNext(m, v20+int32(28))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L26
	}
L22:
	;
	v346 = F_raxRemove(m, v39, v20+int32(16), v117, int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L86
	}
L23:
	;
	F__serverAssert(m, int32(_a2507), int32(_a2500), int32(600))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L85
	}
L24:
	;
	F__serverAssert(m, int32(_a2503), int32(_a2500), int32(797))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L84
	}
L25:
	;
	F__serverAssert(m, int32(_a2530), int32(_a2500), int32(1566))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L83
	}
L26:
	;
	if v90 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = int64(0)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v96
	switch v96 + int32(1) {
	case 0:
		v112 = v96
		goto L28
	case 1:
		goto L30
	default:
		goto L29
	}
L28:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v113)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v117 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v106 = int32(1)
	if v96&v106 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	F__serverAssert(m, int32(_a2513), int32(_a2500), int32(781))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	v111 = v106
	goto L34
L33:
	;
	v111 = v96 & int32(6)
	goto L34
L34:
	;
	v112 = v111
	goto L28
L35:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v20)+180))
	F_raxStop(m, v20+int32(28))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	v120 = F__emscripten_memcpy_bulkmem(m, v20+int32(16), v113, v117)
	mBase = m.M
	goto L36
L38:
	;
	v127 = int64(56)
	v129 = int64(65280)
	v131 = int64(40)
	v134 = int64(16711680)
	v136 = int64(24)
	v138 = int64(4278190080)
	v140 = int64(8)
	if l3 < v114<<(uint(v127)%64)|v114&v129<<(uint(v131)%64)|(v114&v134<<(uint(v136)%64)|v114&v138<<(uint(v140)%64))|(int64(base.Ui64(v114)>>(uint(v140)%64))&v138|int64(base.Ui64(v114)>>(uint(v136)%64))&v134|(int64(base.Ui64(v114)>>(uint(v131)%64))&v129|int64(base.Ui64(v114)>>(uint(v127)%64)))) {
		v356 = v49
		goto L13
	} else {
		goto L39
	}
L39:
	;
	switch v112 + int32(-1) {
	case 0:
		goto L44
	case 1:
		goto L43
	default:
		goto L42
	case 3:
		goto L41
	}
L40:
	;
	v273 = v267 + v49
	if v264 == int32(-1) {
		v333 = v273
		goto L22
	} else {
		goto L76
	}
L41:
	;
	v253 = F_vsetBucketRemoveExpired_HASHTABLE(m, v20+int32(12), l2, l4-v49, l5)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L75
	}
L42:
	;
	F__serverPanic_1(m, int32(_a2500), int32(1589), int32(_a2531), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L74
	}
L43:
	;
	if v96&int32(7) != int32(2) {
		goto L24
	} else {
		goto L51
	}
L44:
	;
	if int64(0) <= l3 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_freeVsetBucket(m, v96)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	v264 = v96
	v267 = int32(0)
	goto L40
L47:
	;
	if l2 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v333 = v49 + int32(1)
	goto L22
L49:
	;
	v172 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v96, l5)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v181 = v96 & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+332)) = v181
	if v181 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v186 = v184 & int32(1073741823)
	v187 = l4 - v49
	if base.Ui32(v186) < base.Ui32(v187) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v264 = v96
	v267 = int32(0)
	goto L40
L54:
	;
	v204 = int32(0)
	goto L59
L55:
	;
	v189 = v186
	goto L57
L56:
	;
	v189 = v187
	goto L57
L57:
	;
	if v189 != 0 {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v264 = v96
	v267 = int32(0)
	goto L40
L59:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	if base.Ui32(v211&int32(1073741823)) <= base.Ui32(v204) {
		goto L23
	} else {
		goto L61
	}
L60:
	;
	v231 = F_pvSplit(m, v20+int32(332), v189)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L68
	}
L61:
	;
	if int64(0) <= l3 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if l2 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v264 = v96
	v267 = int32(0)
	goto L40
L64:
	;
	v227 = v204 + int32(1)
	if v227 != v189 {
		v204 = v227
		goto L59
	} else {
		goto L67
	}
L65:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v181+int32(8)+v204<<(uint(int32(2))%32))))
	v224 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v223, l5)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	goto L60
L68:
	;
	if v231 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v236 = v231 | int32(2)
	goto L71
L70:
	;
	v236 = int32(-1)
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v236
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v20)+332))
	if v238 == int32(0) {
		v264 = v236
		v267 = v189
		goto L40
	} else {
		goto L72
	}
L72:
	;
	F_valkey_free(m, v238)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v264 = v236
	v267 = v189
	goto L40
L74:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v264 = v255
	v267 = v253
	goto L40
L76:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v264 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v356 = v273
	goto L13
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v278 | int32(3)
	goto L77
L79:
	;
	v281 = int32(3)
	v282 = int32(base.Ui32(v278) >> (uint(v281) % 32))
	v289 = int32(4)
	if v278&v289 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v294 = v289
	goto L82
L81:
	;
	v294 = v282 << (uint(int32(2)) % 32)
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122+v282+(int32(0)-v282)&v281+v294+int32(4)))) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v278&int32(-4) | int32(1)
	goto L77
L83:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	if base.Ui32(v333) < base.Ui32(l4) {
		v49 = v333
		goto L16
	} else {
		goto L87
	}
L87:
	;
	goto L17
L88:
	;
	F_shrinkRaxBucketIfPossible(m, l0, l1)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L4
	} else {
		goto L92
	}
L89:
	;
	if v366 != int64(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	F_raxFree(m, v39)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(-1)
	v479 = v356
	goto L1
L92:
	;
	v479 = v356
	goto L1
L93:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	v479 = v382
	goto L1
L95:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	v392 = v390 & int32(1073741823)
	if base.Ui32(v392) < base.Ui32(l4) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v394 = v392
	goto L98
L97:
	;
	v394 = l4
	goto L98
L98:
	;
	if v394 == int32(0) {
		v479 = v387
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v407 = int32(0)
	goto L103
L100:
	;
	v445 = F_pvSplit(m, v20+int32(28), v442)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L4
	} else {
		goto L114
	}
L101:
	;
	F__serverAssert(m, int32(_a2507), int32(_a2500), int32(600))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L4
	} else {
		goto L113
	}
L102:
	;
	if v407 != 0 {
		v442 = v407
		goto L100
	} else {
		goto L112
	}
L103:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	if base.Ui32(v417&int32(1073741823)) <= base.Ui32(v407) {
		goto L101
	} else {
		goto L105
	}
L104:
	;
	v442 = v394
	goto L100
L105:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v385+int32(8)+v407<<(uint(int32(2))%32))))
	v425 = m.T0[l1].(func(*base.Module, int32) int64)(m, v424)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	if l3 < v425 {
		goto L102
	} else {
		goto L107
	}
L107:
	;
	if l2 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v433 = v407 + int32(1)
	if v433 != v394 {
		v407 = v433
		goto L103
	} else {
		goto L111
	}
L109:
	;
	v430 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v424, l5)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	goto L104
L112:
	;
	v479 = int32(0)
	goto L1
L113:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	if v445 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v450 = v445 | int32(2)
	goto L117
L116:
	;
	v450 = int32(-1)
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v450
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if v452 == int32(0) {
		v479 = v442
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_valkey_free(m, v452)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	v479 = v442
	goto L1
L120:
	;
	v459 = m.T0[l1].(func(*base.Module, int32) int64)(m, v23)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	if l3 < v459 {
		v479 = v7
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_freeVsetBucket(m, v462)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(-1)
	v467 = int32(1)
	if l2 == int32(0) {
		v479 = v467
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v470 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v23, l5)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	v479 = v467
	goto L1
}
func F_vsetUpdateEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int64) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v317 int32
	_ = v317
	var v335 int32
	_ = v335
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v17 + int32(32)
	return v335
L2:
	;
	if v317 == int32(-1) {
		v335 = int32(0)
		goto L1
	} else {
		goto L91
	}
L3:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	if v271 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L4:
	;
	if v110 == int32(0) {
		v317 = v126
		goto L2
	} else {
		goto L84
	}
L5:
	;
	F__serverAssert(m, int32(_a2525), int32(_a2500), int32(663))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L19
	} else {
		goto L83
	}
L6:
	;
	F__serverAssert(m, int32(_a2525), int32(_a2500), int32(663))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L19
	} else {
		goto L82
	}
L7:
	;
	F__serverAssert(m, int32(_a2526), int32(_a2500), int32(2044))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L19
	} else {
		goto L81
	}
L8:
	;
	v22 = int32(1)
	if l2 != l3 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v26 = base.B2i32(l5 == int64(-1))
	if l5 == int64(-1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if l4 == l5 {
		v335 = v22
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	if l5 == int64(-1) {
		goto L72
	} else {
		goto L73
	}
L13:
	;
	if l3 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if l2 == int32(0) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if l4 == int64(-1) {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v33 = int32(0)
	switch v19 + int32(1) {
	case 0:
		v335 = v33
		goto L1
	case 1:
		goto L18
	default:
		goto L17
	}
L17:
	;
	if v19&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	F__serverAssert(m, int32(_a2513), int32(_a2500), int32(781))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = int64(0)
	v110 = F_findBucket_2(m, v19&int32(-8), l4, v17+int32(24), v17+int32(20), v17+int32(8), v17+int32(4))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L19
	} else {
		goto L40
	}
L22:
	;
	if l4 == l5 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	if v19 != l2 {
		v335 = v33
		goto L1
	} else {
		goto L25
	}
L24:
	;
	switch v19&int32(6) + int32(-2) {
	case 0:
		goto L22
	default:
		v335 = v33
		goto L1
	case 4:
		goto L21
	}
L25:
	;
	v317 = l3
	goto L2
L26:
	;
	v60 = v19 & int32(-8)
	if v60 == int32(0) {
		v335 = v33
		goto L1
	} else {
		goto L31
	}
L27:
	;
	v52 = F_vsetRemoveEntryWithExpiry(m, l0, l1, l2, l4)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	if v52 == int32(0) {
		v335 = v33
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v56 = F_vsetAddEntry(m, l0, l1, l3)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	v335 = int32(1)
	goto L1
L31:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v65 = v63 & int32(1073741823)
	if v65 == int32(0) {
		v335 = v33
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v72 = int32(0)
	goto L34
L33:
	;
	if v72 == v65 {
		v335 = int32(0)
		goto L1
	} else {
		goto L38
	}
L34:
	;
	v87 = v60 + int32(8) + v72<<(uint(int32(2))%32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v88 == l2 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v92 = v72 + int32(1)
	if v92 != v65 {
		v72 = v92
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v335 = int32(0)
	goto L1
L38:
	;
	if base.Ui32(v65) <= base.Ui32(v72) {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = l3
	v317 = v19
	goto L2
L40:
	;
	if base.Ui64(l5^l4) < base.Ui64(int64(16)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v126 = int32(-1)
	switch v110 + int32(1) {
	case 0:
		v317 = v126
		goto L2
	case 1:
		goto L47
	default:
		goto L46
	}
L42:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v120 = F_removeEntryFromRaxBucket(m, v19, l2, v110, v17+int32(24), v118, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	if v120 == int32(0) {
		v317 = int32(-1)
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v124 = F_insertToBucket_RAX(m, l1, v19, l3, l5)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L19
	} else {
		goto L45
	}
L45:
	;
	v317 = v124
	goto L2
L46:
	;
	if v110&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L47:
	;
	F__serverAssert(m, int32(_a2513), int32(_a2500), int32(781))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L19
	} else {
		goto L48
	}
L48:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F__serverPanic_1(m, int32(_a2500), int32(1991), int32(_a2527), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L19
	} else {
		goto L71
	}
L50:
	;
	if l2 == l3 {
		goto L4
	} else {
		goto L66
	}
L51:
	;
	v144 = int32(-1)
	v146 = v110 & int32(-8)
	if v146 == int32(0) {
		v271 = v144
		goto L3
	} else {
		goto L57
	}
L52:
	;
	if v110 == l2 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	switch v110&int32(6) + int32(-2) {
	case 0:
		goto L51
	default:
		goto L49
	case 2:
		goto L50
	}
L54:
	;
	v143 = l3
	goto L56
L55:
	;
	v143 = int32(-1)
	goto L56
L56:
	;
	v271 = v143
	goto L3
L57:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v151 = v149 & int32(1073741823)
	if v151 == int32(0) {
		v271 = v144
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v158 = int32(0)
	goto L60
L59:
	;
	if v158 == v151 {
		v271 = v144
		goto L3
	} else {
		goto L64
	}
L60:
	;
	v173 = v146 + int32(8) + v158<<(uint(int32(2))%32)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	if v174 == l2 {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v177 = v158 + int32(1)
	if v177 == v151 {
		v271 = v144
		goto L3
	} else {
		goto L63
	}
L63:
	;
	v158 = v177
	goto L60
L64:
	;
	if base.Ui32(v151) <= base.Ui32(v158) {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = l3
	goto L4
L66:
	;
	v184 = v110 & int32(-8)
	v185 = F_hashtableDelete(m, v184, l2)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L19
	} else {
		goto L67
	}
L67:
	;
	v187 = F_hashtableAdd(m, v184, l3)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L19
	} else {
		goto L68
	}
L68:
	;
	if v187 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F__serverAssert(m, int32(_a2528), int32(_a2500), int32(1955))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L19
	} else {
		goto L70
	}
L70:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	v213 = int32(0)
	if l4 == int64(-1) {
		v335 = v213
		goto L1
	} else {
		goto L77
	}
L73:
	;
	if l3 == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v204 = int32(0)
	if base.B2i32(l2 == v204)|base.B2i32(l4 == int64(-1)) == v204 {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v211 = F_vsetAddEntry(m, l0, l1, l3)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L19
	} else {
		goto L76
	}
L76:
	;
	v335 = v22
	goto L1
L77:
	;
	if l2 == int32(0) {
		v335 = v213
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v218 = int32(0)
	if base.B2i32(l3 == v218)|base.B2i32(l5 == int64(-1)) == v218 {
		v335 = v213
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v225 = F_vsetRemoveEntryWithExpiry(m, l0, l1, l2, l4)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L19
	} else {
		goto L80
	}
L80:
	;
	v335 = v225
	goto L1
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
L83:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	v271 = v110
	goto L3
L85:
	;
	v317 = v19
	goto L2
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275))) = v278 | int32(3)
	goto L85
L87:
	;
	v281 = int32(3)
	v282 = int32(base.Ui32(v278) >> (uint(v281) % 32))
	v289 = int32(4)
	if v278&v289 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v294 = v289
	goto L90
L89:
	;
	v294 = v282 << (uint(int32(2)) % 32)
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275+v282+(int32(0)-v282)&v281+v294+int32(4)))) = v271
	*(*int32)(unsafe.Add(mBase, uint32(v275))) = v278&int32(-4) | int32(1)
	goto L85
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v317
	v335 = int32(1)
	goto L1
}
