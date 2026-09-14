package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F__listZiplistEntryConvertAndValidate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int64
	_ = v12
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v12 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(24)))) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(16)))) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v12
	v28 = F_ziplistGet(m, l0, v8+int32(44), v8+int32(40), v8+int32(32))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		if v28 != 0 {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
			if v33 == int32(0) {
				v38 = *(*int64)(unsafe.Add(mBase, uint32(v8)+32))
				if v38 <= int64(-1) {
					v47 = int32(45)
					*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v47)
					v51 = int32(1)
					v56 = v8 + v51
					v57 = int32(31)
					v58 = int64(0) - v38
					v59 = v51
				} else {
					v56 = v8
					v57 = int32(32)
					v58 = v38
					v59 = int32(0)
				}
				v60 = F_ull2string(m, v56, v57, v58)
				mBase = m.M
				if v60 == int32(0) {
					v79 = int32(0)
				} else {
					v79 = v60 + v59
				}
				*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v79
				*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v8
				v82 = v8
				v83 = v79
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
				v82 = v33
				v83 = v36
			}
			v84 = F_quicklistPushTail(m, l2, v82, v83)
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return int32(0)
			} else {
				v87 = int32(1)
				m.G0 = v8 + int32(48)
				return v87
			}
		} else {
			v87 = int32(0)
			m.G0 = v8 + int32(48)
			return v87
		}
	}
}
func F_addListRangeReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	v9 = F_listTypeLength(m, l1)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v12 = l2>>(uint(int32(31))%32)&v9 + l2
		v13 = int32(0)
		if v13 < v12 {
			v16 = v12
		} else {
			v16 = v13
		}
		v20 = l3>>(uint(int32(31))%32)&v9 + l3
		if v20 < v16 {
			v24 = *(*int32)(unsafe.Add(mBase, _consts[289]))
			F_addReply(m, l0, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				return
			}
		} else {
			if v16 < v9 {
				if base.Ui32(v20) < base.Ui32(v9) {
					v30 = v20
				} else {
					v30 = v9 + int32(-1)
				}
				if l4 != 0 {
					v31 = v30
				} else {
					v31 = v16
				}
				v34 = v30 - v16 + int32(1)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				switch int32(base.Ui32(v35)>>(uint(int32(4))%32))&int32(15) + int32(-9) {
				case 0:
					F_addListQuicklistRangeReply(m, l0, l1, v31, v34, l4)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						return
					}
				default:
					F__serverPanic_1(m, int32(_a1497), int32(729), int32(_a852), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 2:
					F_addListListpackRangeReply(m, l0, l1, v31, v34, l4)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _consts[289]))
				F_addReply(m, l0, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_createListListpackObject(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_lpNew(m, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
		v15 = int32(12)
		v18 = F_zmalloc_usable(m, v15, v6+v15)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v9
			*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(34359738369)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v23&int32(-241) | int32(176)
			m.G0 = v6 + int32(16)
			return v18
		}
	}
}
func F_freeListObject(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v5)>>(uint(int32(4))%32))&int32(15) + int32(-9) {
	case 0:
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v12&int32(4) == int32(0) {
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_quicklistRelease(m, v73)
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return
			} else {
				return
			}
		} else {
			if v12&int32(1) != 0 {
				v21 = int32(16)
			} else {
				v21 = int32(8)
			}
			v22 = l0 + v21
			if v12&int32(2) == int32(0) {
				v53 = v22
			} else {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v28 = v22 + v27
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
				switch v32 & int32(7) {
				case 0:
					v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
				case 1:
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-2)))))
					v49 = v39
				case 2:
					v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-4)))))
					v49 = v42
				case 3:
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-8))))
					v49 = v45
				case 4:
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-16))))
					v49 = v48
				default:
					v49 = int32(0)
				}
				v53 = v28 + int32(1) + v49 + int32(1)
			}
			v68 = *(*int32)(unsafe.Add(mBase, _consts[249]))
			F_quicklistRelease(m, v53+v68)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				return
			}
		}
	default:
		F__serverPanic_1(m, int32(_a838), int32(567), int32(_a844), int32(0))
		mBase = m.M
		v145 = m.ExcPending
		if v145 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 2:
		v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v76&int32(4) == int32(0) {
			v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_lpFree(m, v137)
			mBase = m.M
			v139 = m.ExcPending
			if v139 != 0 {
				return
			} else {
				return
			}
		} else {
			if v76&int32(1) != 0 {
				v85 = int32(16)
			} else {
				v85 = int32(8)
			}
			v86 = l0 + v85
			if v76&int32(2) == int32(0) {
				v117 = v86
			} else {
				v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
				v92 = v86 + v91
				v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
				switch v96 & int32(7) {
				case 0:
					v113 = int32(base.Ui32(v96) >> (uint(int32(3)) % 32))
				case 1:
					v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+int32(-2)))))
					v113 = v103
				case 2:
					v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92+int32(-4)))))
					v113 = v106
				case 3:
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v92+int32(-8))))
					v113 = v109
				case 4:
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v92+int32(-16))))
					v113 = v112
				default:
					v113 = int32(0)
				}
				v117 = v92 + int32(1) + v113 + int32(1)
			}
			v132 = *(*int32)(unsafe.Add(mBase, _consts[249]))
			F_lpFree(m, v117+v132)
			mBase = m.M
			v136 = m.ExcPending
			if v136 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_listCommandHandler(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = m.G6
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+280))
	v16 = int32(5)
	if base.Ui32(l1) < base.Ui32(int32(2)) {
		v136 = v15
		v137 = v16
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v139 = m.G6
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+276))
	if v140 < int32(1) {
		goto L40
	} else {
		goto L41
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = m.G7
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v23 = m.T0[v22].(func(*base.Module, int32, int32) int32)(m, v19, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v30 = v23
	goto L6
L5:
	;
	if int32(0) < v75 {
		goto L20
	} else {
		goto L21
	}
L6:
	;
	v35 = v30 + int32(1)
	v36 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30))))
	v37 = F___isspace_1(m, v36)
	mBase = m.M
	if v37 != 0 {
		v30 = v35
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v38 = int32(1)
	switch v36&int32(255) + int32(-43) {
	case 0:
		v44 = v38
		goto L10
	default:
		v46 = v30
		v47 = v36
		v48 = v38
		goto L9
	case 2:
		goto L11
	}
L8:
	;
	goto L7
L9:
	;
	v51 = v47 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v51) {
		v69 = int32(0)
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35))))
	v46 = v35
	v47 = v45
	v48 = v44
	goto L9
L11:
	;
	v44 = int32(0)
	goto L10
L12:
	;
	if v48 != 0 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v55 = int32(0)
	v56 = v46
	v57 = v51
	goto L14
L14:
	;
	v59 = int32(10)
	v61 = v55*v59 - v57
	v62 = int32(*(*int8)(unsafe.Add(mBase, uint32(v56)+1)))
	v66 = v62 + int32(-48)
	if base.Ui32(v66) < base.Ui32(v59) {
		v55 = v61
		v56 = v56 + int32(1)
		v57 = v66
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v69 = v61
	goto L12
L16:
	;
	goto L15
L17:
	;
	v75 = int32(0) - v69
	goto L19
L18:
	;
	v75 = v69
	goto L19
L19:
	;
	goto L5
L20:
	;
	v78 = v75
	goto L22
L21:
	;
	v78 = v15
	goto L22
L22:
	;
	if l1 == int32(2) {
		v136 = v78
		v137 = v16
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v83 = m.G7
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v85 = m.T0[v84].(func(*base.Module, int32, int32) int32)(m, v81, int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v90 = v85
	goto L26
L25:
	;
	v136 = v78
	v137 = v135
	goto L1
L26:
	;
	v95 = v90 + int32(1)
	v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v90))))
	v97 = F___isspace_1(m, v96)
	mBase = m.M
	if v97 != 0 {
		v90 = v95
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v98 = int32(1)
	switch v96&int32(255) + int32(-43) {
	case 0:
		v104 = v98
		goto L30
	default:
		v106 = v90
		v107 = v96
		v108 = v98
		goto L29
	case 2:
		goto L31
	}
L28:
	;
	goto L27
L29:
	;
	v111 = v107 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v111) {
		v129 = int32(0)
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v105 = int32(*(*int8)(unsafe.Add(mBase, uint32(v95))))
	v106 = v95
	v107 = v105
	v108 = v104
	goto L29
L31:
	;
	v104 = int32(0)
	goto L30
L32:
	;
	if v108 != 0 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v115 = int32(0)
	v116 = v106
	v117 = v111
	goto L34
L34:
	;
	v119 = int32(10)
	v121 = v115*v119 - v117
	v122 = int32(*(*int8)(unsafe.Add(mBase, uint32(v116)+1)))
	v126 = v122 + int32(-48)
	if base.Ui32(v126) < base.Ui32(v119) {
		v115 = v121
		v116 = v116 + int32(1)
		v117 = v126
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v129 = v121
	goto L32
L36:
	;
	goto L35
L37:
	;
	v135 = int32(0) - v129
	goto L39
L38:
	;
	v135 = v129
	goto L39
L39:
	;
	goto L25
L40:
	;
	v258 = m.G14
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	m.T0[v259].(func(*base.Module))(m)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L3
	} else {
		goto L62
	}
L41:
	;
	v144 = int32(1)
	v150 = v140
	goto L42
L42:
	;
	if v136 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L40
L44:
	;
	if v144 < v243 {
		v144 = v144 + int32(1)
		v150 = v243
		goto L42
	} else {
		goto L61
	}
L45:
	;
	v162 = m.G3
	v163 = int32(0)
	v164 = m.G6
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+272))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v165+v144<<(uint(int32(2))%32)+int32(-4))))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v164)+260))
	if v163 < v172 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v155 = v136 - v144
	v157 = v155 >> (uint(int32(31)) % 32)
	if v137 < v155^v157-v157 {
		v243 = v150
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v144
	v218 = m.G6
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+280))
	if v219 == v144 {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	v180 = v163
	goto L52
L50:
	;
	v208 = v162 + int32(_a1717)
	v213 = v162 + int32(_a1718)
	goto L48
L51:
	;
	v208 = v188 + int32(_a1719)
	v213 = v188 + int32(_a1720)
	goto L48
L52:
	;
	v188 = m.G3
	v189 = m.G6
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189+v180<<(uint(int32(2))%32))+4))
	if v193 == v144 {
		goto L51
	} else {
		goto L54
	}
L53:
	;
	v208 = v195 + int32(_a1717)
	v213 = v195 + int32(_a1718)
	goto L48
L54:
	;
	v195 = m.G3
	v197 = v180 + int32(1)
	if v197 != v172 {
		v180 = v197
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v221 = v208
	goto L58
L57:
	;
	v221 = v213
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v221
	v223 = m.G3
	v224 = m.G15
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v226 = m.G12
	v230 = m.T0[v225].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v223+int32(_a1721), v12)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	m.T0[v233].(func(*base.Module, int32, int32))(m, v230, int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v218)+276))
	v243 = v236
	goto L44
L61:
	;
	goto L43
L62:
	;
	m.G0 = v12 + int32(16)
	return int32(1)
}
func F_listElementsRemoved(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	if l2 != 0 {
		v10 = int32(_a1501)
	} else {
		v10 = int32(_a1502)
	}
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_notifyKeyspaceEvent(m, int32(16), v10, l1, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = F_listTypeLength(m, l3)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v15 != 0 {
				v31 = int32(0)
				F_listTypeTryConversionRaw(m, l3, int32(2), v31, v31, v31, v31, v31)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					if l5 == int32(0) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
					}
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					F_signalModifiedKey(m, l0, v42, l1)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = int32(_a69)
						v47 = *(*int64)(unsafe.Add(mBase, _consts[60]))
						*(*int64)(unsafe.Add(mBase, _consts[60])) = v47 + base.I64_extend_i32_s(l4)
						return
					}
				}
			} else {
				if l5 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(1)
				}
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v22 = F_dbDelete(m, v21, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
					F_notifyKeyspaceEvent(m, int32(4), int32(_a132), l1, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						F_signalModifiedKey(m, l0, v42, l1)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							v45 = int32(_a69)
							v47 = *(*int64)(unsafe.Add(mBase, _consts[60]))
							*(*int64)(unsafe.Add(mBase, _consts[60])) = v47 + base.I64_extend_i32_s(l4)
							return
						}
					}
				}
			}
		}
	}
}
func F_listEmpty(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v8 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	goto L1
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = v8
	v14 = v11
	goto L5
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v18 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	F_valkey_free(m, v14)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L11
	}
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	m.T0[v18].(func(*base.Module, int32))(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	goto L7
L11:
	;
	v27 = v13 + int32(-1)
	if v27 != 0 {
		v13 = v27
		v14 = v17
		goto L5
	} else {
		goto L12
	}
L12:
	;
	goto L6
}
func F_listInitNode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	return
}
func F_listJoin(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v6 == int32(0) {
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v10
		if v10 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v9
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v9
		}
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v18 + v6
		*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
	}
	return
}
func F_listLinkNodeTail(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v5 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
		v16 = v11
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
		v8 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v8
		v16 = v8
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v16
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v5 + int32(1)
	return
}
func F_listPopSaver(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_createStringObject_1(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_listReleaseVoid(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v43 int32
	_ = v43
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L13
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v8 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	goto L1
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = v8
	v14 = v11
	goto L5
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v18 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	F_valkey_free(m, v14)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L11
	}
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	m.T0[v18].(func(*base.Module, int32))(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	goto L7
L11:
	;
	v27 = v13 + int32(-1)
	if v27 != 0 {
		v13 = v27
		v14 = v17
		goto L5
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	return
}
func F_listTypeDelete(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v57 int32
	_ = v57
	var v63 int64
	_ = v63
	var v68 int32
	_ = v68
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v86 int64
	_ = v86
	var v91 int32
	_ = v91
	var v103 int64
	_ = v103
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int64
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
	switch v10 + int32(-9) {
	case 0:
		v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_quicklistDelEntry(m, v141, l1+int32(8))
		mBase = m.M
		v145 = m.ExcPending
		if v145 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	default:
		F__serverPanic_1(m, int32(_a1497), int32(424), int32(_a852), int32(0))
		mBase = m.M
		v139 = m.ExcPending
		if v139 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 2:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = F_objectGetVal(m, v15)
		mBase = m.M
		v19 = F_lpDelete(m, v16, v13, v7+int32(12))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_objectSetVal(m, v15, v19)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
				if v24 == int32(1) {
					v131 = v23
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v131
					m.G0 = v7 + int32(16)
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v28 = F_objectGetVal(m, v27)
					mBase = m.M
					if v23 == int32(0) {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						if v40 == int32(7) {
							v127 = int32(0)
						} else {
							v43 = int32(-1)
							v44 = v28 + v40
							v49 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44+int32(-2)))))
							v52 = base.I64_extend_i32_u(v49 & int32(127))
							if v43 < v49 {
								v120 = v43
								v122 = v52
							} else {
								v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44+int32(-3)))))
								v63 = base.I64_extend_i32_u(v57&int32(127))<<(uint(int64(7))%64) | v52
								if int32(-1) < v57 {
									v103 = v63
									if base.Ui64(int64(128)) <= base.Ui64(v103) {
										if base.Ui64(int64(16384)) <= base.Ui64(v103) {
											if base.Ui64(int64(2097152)) <= base.Ui64(v103) {
												if base.Ui64(v103) < base.Ui64(int64(268435456)) {
													v117 = int32(-4)
												} else {
													v117 = int32(-5)
												}
												v120 = v117
												v122 = v103
											} else {
												v120 = int32(-3)
												v122 = v103
											}
										} else {
											v120 = int32(-2)
											v122 = v103
										}
									} else {
										v120 = int32(-1)
										v122 = v103
									}
								} else {
									v68 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44+int32(-4)))))
									v74 = base.I64_extend_i32_u(v68&int32(127))<<(uint(int64(14))%64) | v63
									if int32(-1) < v68 {
										v103 = v74
										if base.Ui64(int64(128)) <= base.Ui64(v103) {
											if base.Ui64(int64(16384)) <= base.Ui64(v103) {
												if base.Ui64(int64(2097152)) <= base.Ui64(v103) {
													if base.Ui64(v103) < base.Ui64(int64(268435456)) {
														v117 = int32(-4)
													} else {
														v117 = int32(-5)
													}
													v120 = v117
													v122 = v103
												} else {
													v120 = int32(-3)
													v122 = v103
												}
											} else {
												v120 = int32(-2)
												v122 = v103
											}
										} else {
											v120 = int32(-1)
											v122 = v103
										}
									} else {
										v77 = int32(-5)
										v80 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44+v77))))
										v86 = base.I64_extend_i32_u(v80&int32(127))<<(uint(int64(21))%64) | v74
										if int32(-1) < v80 {
											v103 = v86
											if base.Ui64(int64(128)) <= base.Ui64(v103) {
												if base.Ui64(int64(16384)) <= base.Ui64(v103) {
													if base.Ui64(int64(2097152)) <= base.Ui64(v103) {
														if base.Ui64(v103) < base.Ui64(int64(268435456)) {
															v117 = int32(-4)
														} else {
															v117 = int32(-5)
														}
														v120 = v117
														v122 = v103
													} else {
														v120 = int32(-3)
														v122 = v103
													}
												} else {
													v120 = int32(-2)
													v122 = v103
												}
											} else {
												v120 = int32(-1)
												v122 = v103
											}
										} else {
											v91 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44+int32(-6)))))
											if v91 < int32(0) {
												v120 = v77
												v122 = int64(-1)
											} else {
												v103 = base.I64_extend_i32_u(v91&int32(127))<<(uint(int64(28))%64) | v86
												if base.Ui64(int64(128)) <= base.Ui64(v103) {
													if base.Ui64(int64(16384)) <= base.Ui64(v103) {
														if base.Ui64(int64(2097152)) <= base.Ui64(v103) {
															if base.Ui64(v103) < base.Ui64(int64(268435456)) {
																v117 = int32(-4)
															} else {
																v117 = int32(-5)
															}
															v120 = v117
															v122 = v103
														} else {
															v120 = int32(-3)
															v122 = v103
														}
													} else {
														v120 = int32(-2)
														v122 = v103
													}
												} else {
													v120 = int32(-1)
													v122 = v103
												}
											}
										}
									}
								}
							}
							v127 = v44 + v43 + (v120 - base.I32_wrap_i64(v122))
						}
						v131 = v127
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v131
						m.G0 = v7 + int32(16)
						return
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						v32 = F_lpPrev(m, v28, v31)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_listTypeDup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3&int32(15) != int32(1) {
		F__serverAssert(m, int32(_a1498), int32(_a1497), int32(436))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		switch int32(base.Ui32(v3)>>(uint(int32(4))%32))&int32(15) + int32(-9) {
		case 0:
			v14 = F_objectGetVal(m, l0)
			mBase = m.M
			v15 = F_quicklistDup(m, v14)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v29 = v15
				v31 = F_createObject(m, int32(1), v29)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = v33&int32(-241) | v36&int32(240)
					return v31
				}
			}
		default:
			F__serverPanic_1(m, int32(_a1497), int32(441), int32(_a852), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 2:
			v26 = F_objectGetVal(m, l0)
			mBase = m.M
			v27 = F_lpDup(m, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = v27
				v31 = F_createObject(m, int32(1), v29)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = v33&int32(-241) | v36&int32(240)
					return v31
				}
			}
		}
	}
}
func F_listTypeGet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
	switch v10 + int32(-9) {
	case 0:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v13 == int32(0) {
			v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v38
			v43 = v38
			v44 = F_createStringObjectFromLongLong(m, v43)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				v46 = v44
				m.G0 = v7 + int32(16)
				return v46
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v17 = F_createStringObject_1(m, v13, v16)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v46 = v17
				m.G0 = v7 + int32(16)
				return v46
			}
		}
	default:
		F__serverPanic_1(m, int32(_a1497), int32(307), int32(_a852), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 2:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v31 = F_lpGetValue(m, v28, v7+int32(12), v7)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			if v31 == int32(0) {
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
				v43 = v40
				v44 = F_createStringObjectFromLongLong(m, v43)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v46 = v44
					m.G0 = v7 + int32(16)
					return v46
				}
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v36 = F_createStringObject_1(m, v31, v35)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v46 = v36
					m.G0 = v7 + int32(16)
					return v46
				}
			}
		}
	}
}
func F_listTypeGetValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+4)))
	switch v11 + int32(-9) {
	case 0:
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v14 == int32(0) {
			v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v35
			v38 = int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
			v38 = v14
		}
		m.G0 = v8 + int32(16)
		return v38
	default:
		F__serverPanic_1(m, int32(_a1497), int32(307), int32(_a852), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 2:
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v22 = F_lpGetValue(m, v19, v8+int32(12), l2)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26
			v38 = v22
			m.G0 = v8 + int32(16)
			return v38
		}
	}
}
func F_listTypeInsert(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v11 = F_getDecodedObject(m, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = F_objectGetVal(m, v11)
		mBase = m.M
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-1)))))
		switch v16 & int32(7) {
		case 0:
			v33 = int32(base.Ui32(v16) >> (uint(int32(3)) % 32))
		case 1:
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
			v33 = v23
		case 2:
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
			v33 = v26
		case 3:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
			v33 = v29
		case 4:
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
			v33 = v32
		default:
			v33 = int32(0)
		}
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
		switch v35 + int32(-9) {
		case 0:
			switch l2 {
			case 0:
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
				F_quicklistInsertBefore(m, v38, l0+int32(8), v13, v33)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_decrRefCount(m, v11)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						return
					}
				}
			case 1:
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
				F_quicklistInsertAfter(m, v64, l0+int32(8), v13, v33)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					F_decrRefCount(m, v11)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						return
					}
				}
			default:
				F_decrRefCount(m, v11)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					return
				}
			}
		default:
			F__serverPanic_1(m, int32(_a1497), int32(341), int32(_a852), int32(0))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 2:
			v45 = F_objectGetVal(m, v9)
			mBase = m.M
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v51 = F_lpInsertString(m, v45, v13, v33, v46, base.B2i32(l2 == int32(1)), l0+int32(4))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				F_objectSetVal(m, v9, v51)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_decrRefCount(m, v11)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_listTypeLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v2)>>(uint(int32(4))%32))&int32(15) + int32(-9) {
	case 0:
		v22 = F_objectGetVal(m, l0)
		mBase = m.M
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
		return v23
	default:
		F__serverPanic_1(m, int32(_a1497), int32(218), int32(_a852), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 2:
		v9 = F_objectGetVal(m, l0)
		mBase = m.M
		v10 = F_lpLength(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_listTypePop(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v61 int32
	_ = v61
	var v67 int64
	_ = v67
	var v72 int32
	_ = v72
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int64
	_ = v90
	var v95 int32
	_ = v95
	var v107 int64
	_ = v107
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v11)>>(uint(int32(4))%32))&int32(15) + int32(-9) {
	case 0:
		v18 = F_objectGetVal(m, l0)
		mBase = m.M
		if l1 != 0 {
			v21 = int32(-1)
		} else {
			v21 = int32(0)
		}
		v26 = F_quicklistPopCustom(m, v18, v21, v7+int32(44), int32(0), v7, int32(1089))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
			if v26 == int32(0) {
				v151 = v30
				m.G0 = v7 + int32(48)
				return v151
			} else {
				if v30 != 0 {
					v151 = v30
					m.G0 = v7 + int32(48)
					return v151
				} else {
					v33 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
					v34 = F_createStringObjectFromLongLong(m, v33)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v151 = v34
						m.G0 = v7 + int32(48)
						return v151
					}
				}
			}
		}
	default:
		F__serverPanic_1(m, int32(_a1497), int32(207), int32(_a852), int32(0))
		mBase = m.M
		v162 = m.ExcPending
		if v162 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 2:
		v36 = F_objectGetVal(m, l0)
		mBase = m.M
		if l1 != 0 {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
			if v44 == int32(7) {
				v131 = int32(0)
			} else {
				v47 = int32(-1)
				v48 = v36 + v44
				v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48+int32(-2)))))
				v56 = base.I64_extend_i32_u(v53 & int32(127))
				if v47 < v53 {
					v124 = v47
					v126 = v56
				} else {
					v61 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48+int32(-3)))))
					v67 = base.I64_extend_i32_u(v61&int32(127))<<(uint(int64(7))%64) | v56
					if int32(-1) < v61 {
						v107 = v67
						if base.Ui64(int64(128)) <= base.Ui64(v107) {
							if base.Ui64(int64(16384)) <= base.Ui64(v107) {
								if base.Ui64(int64(2097152)) <= base.Ui64(v107) {
									if base.Ui64(v107) < base.Ui64(int64(268435456)) {
										v121 = int32(-4)
									} else {
										v121 = int32(-5)
									}
									v124 = v121
									v126 = v107
								} else {
									v124 = int32(-3)
									v126 = v107
								}
							} else {
								v124 = int32(-2)
								v126 = v107
							}
						} else {
							v124 = int32(-1)
							v126 = v107
						}
					} else {
						v72 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48+int32(-4)))))
						v78 = base.I64_extend_i32_u(v72&int32(127))<<(uint(int64(14))%64) | v67
						if int32(-1) < v72 {
							v107 = v78
							if base.Ui64(int64(128)) <= base.Ui64(v107) {
								if base.Ui64(int64(16384)) <= base.Ui64(v107) {
									if base.Ui64(int64(2097152)) <= base.Ui64(v107) {
										if base.Ui64(v107) < base.Ui64(int64(268435456)) {
											v121 = int32(-4)
										} else {
											v121 = int32(-5)
										}
										v124 = v121
										v126 = v107
									} else {
										v124 = int32(-3)
										v126 = v107
									}
								} else {
									v124 = int32(-2)
									v126 = v107
								}
							} else {
								v124 = int32(-1)
								v126 = v107
							}
						} else {
							v81 = int32(-5)
							v84 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48+v81))))
							v90 = base.I64_extend_i32_u(v84&int32(127))<<(uint(int64(21))%64) | v78
							if int32(-1) < v84 {
								v107 = v90
								if base.Ui64(int64(128)) <= base.Ui64(v107) {
									if base.Ui64(int64(16384)) <= base.Ui64(v107) {
										if base.Ui64(int64(2097152)) <= base.Ui64(v107) {
											if base.Ui64(v107) < base.Ui64(int64(268435456)) {
												v121 = int32(-4)
											} else {
												v121 = int32(-5)
											}
											v124 = v121
											v126 = v107
										} else {
											v124 = int32(-3)
											v126 = v107
										}
									} else {
										v124 = int32(-2)
										v126 = v107
									}
								} else {
									v124 = int32(-1)
									v126 = v107
								}
							} else {
								v95 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48+int32(-6)))))
								if v95 < int32(0) {
									v124 = v81
									v126 = int64(-1)
								} else {
									v107 = base.I64_extend_i32_u(v95&int32(127))<<(uint(int64(28))%64) | v90
									if base.Ui64(int64(128)) <= base.Ui64(v107) {
										if base.Ui64(int64(16384)) <= base.Ui64(v107) {
											if base.Ui64(int64(2097152)) <= base.Ui64(v107) {
												if base.Ui64(v107) < base.Ui64(int64(268435456)) {
													v121 = int32(-4)
												} else {
													v121 = int32(-5)
												}
												v124 = v121
												v126 = v107
											} else {
												v124 = int32(-3)
												v126 = v107
											}
										} else {
											v124 = int32(-2)
											v126 = v107
										}
									} else {
										v124 = int32(-1)
										v126 = v107
									}
								}
							}
						}
					}
				}
				v131 = v48 + v47 + (v124 - base.I32_wrap_i64(v126))
			}
			v135 = v131
			if v135 != 0 {
				v139 = F_lpGet(m, v135, v7+int32(32), v7)
				mBase = m.M
				v140 = m.ExcPending
				if v140 != 0 {
					return int32(0)
				} else {
					v141 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
					v142 = F_createStringObject_1(m, v139, v141)
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return int32(0)
					} else {
						v144 = F_objectGetVal(m, l0)
						mBase = m.M
						v146 = F_lpDelete(m, v144, v135, int32(0))
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
							return int32(0)
						} else {
							F_objectSetVal(m, l0, v146)
							mBase = m.M
							v149 = m.ExcPending
							if v149 != 0 {
								return int32(0)
							} else {
								v151 = v142
								m.G0 = v7 + int32(48)
								return v151
							}
						}
					}
				}
			} else {
				v151 = int32(0)
				m.G0 = v7 + int32(48)
				return v151
			}
		} else {
			v37 = F_lpFirst(m, v36)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v135 = v37
				if v135 != 0 {
					v139 = F_lpGet(m, v135, v7+int32(32), v7)
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int32(0)
					} else {
						v141 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
						v142 = F_createStringObject_1(m, v139, v141)
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return int32(0)
						} else {
							v144 = F_objectGetVal(m, l0)
							mBase = m.M
							v146 = F_lpDelete(m, v144, v135, int32(0))
							mBase = m.M
							v147 = m.ExcPending
							if v147 != 0 {
								return int32(0)
							} else {
								F_objectSetVal(m, l0, v146)
								mBase = m.M
								v149 = m.ExcPending
								if v149 != 0 {
									return int32(0)
								} else {
									v151 = v142
									m.G0 = v7 + int32(48)
									return v151
								}
							}
						}
					}
				} else {
					v151 = int32(0)
					m.G0 = v7 + int32(48)
					return v151
				}
			}
		}
	}
}
func F_listTypeReleaseIterator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v2 != int32(9) {
		F_valkey_free(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_quicklistReleaseIterator(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_valkey_free(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_listTypeSetIteratorDirection(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v3 = l2
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v5 == v3 {
		return
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v3)
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		switch v8 + int32(-9) {
		case 0:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = base.B2i32(v3 == int32(0))
			return
		default:
			F__serverPanic_1(m, int32(_a1497), int32(256), int32(_a852), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 2:
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v12 = F_objectGetVal(m, v11)
			mBase = m.M
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v3 != int32(1) {
				v19 = F_lpPrev(m, v12, v13)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v19
					return
				}
			} else {
				v16 = F_lpNext(m, v12, v13)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v16
					return
				}
			}
		}
	}
}
func F_list_get_iter(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v5 = m.G9
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = m.T0[v6].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(4))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v11
		return v7
	}
}
func F_list_length(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	return v2
}
func F_list_release_iter(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = m.G11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	m.T0[v3].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
