package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_addHashFieldToReply(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 == int32(0) {
		F_addReplyNull(m, l0)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	} else {
		v11 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(-1)
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(9223372036854775807)
		v22 = F_hashTypeGetValue(m, l1, l2, v7+int32(12), v7+int32(8), v7, v11)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			if v22 != 0 {
				F_addReplyNull(m, l0)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				if v24 == int32(0) {
					v30 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
					F_addReplyBulkLongLong(m, l0, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					F_addReplyBulkCBuffer(m, l0, v24, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_createHashObject(m *base.Module) int32 {
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
			*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(34359738372)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v23&int32(-241) | int32(176)
			m.G0 = v6 + int32(16)
			return v18
		}
	}
}
func F_hashHashtableTypeMetadataSize(m *base.Module) int32 {
	return int32(4)
}
func F_hashReplyFromListpackEntry(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v4 == int32(0) {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		F_addReplyBulkLongLong(m, l0, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		F_addReplyBulkCBuffer(m, l0, v4, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	}
}
func F_hashTypeDelete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a2356), int32(_a2349), int32(585))
		mBase = m.M
		v98 = m.ExcPending
		if v98 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v15&int32(15) != int32(4) {
			F__serverAssert(m, int32(_a2356), int32(_a2349), int32(585))
			mBase = m.M
			v98 = m.ExcPending
			if v98 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			switch int32(base.Ui32(v15)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
			case 0:
				v68 = F_objectGetVal(m, l0)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(0)
				v73 = F_hashtablePop(m, v68, l1, v11+int32(8))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					if v73 == int32(0) {
						v84 = v73
						m.G0 = v11 + int32(16)
						return v84
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
						F_hashTypeUntrackEntry(m, l0, v77)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
							F_entryFree(m, v80)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								v84 = v73
								m.G0 = v11 + int32(16)
								return v84
							}
						}
					}
				}
			default:
				F__serverPanic_1(m, int32(_a2349), int32(609), int32(_a2355), int32(0))
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			case 9:
				v26 = F_objectGetVal(m, l0)
				mBase = m.M
				v27 = F_lpFirst(m, v26)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					if v27 != 0 {
						v32 = int32(0)
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
						switch v36 & int32(7) {
						case 0:
							v53 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
						case 1:
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
							v53 = v43
						case 2:
							v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
							v53 = v46
						case 3:
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
							v53 = v49
						case 4:
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
							v53 = v52
						default:
							v53 = v32
						}
						v55 = F_lpFind(m, v26, v27, l1, v53, int32(1))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v55
							if v55 == int32(0) {
								v84 = v32
								m.G0 = v11 + int32(16)
								return v84
							} else {
								v63 = F_lpDeleteRangeWithEntry(m, v26, v11+int32(12), int32(2))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									F_objectSetVal(m, l0, v63)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										v84 = int32(1)
										m.G0 = v11 + int32(16)
										return v84
									}
								}
							}
						}
					} else {
						v84 = int32(0)
						m.G0 = v11 + int32(16)
						return v84
					}
				}
			}
		}
	}
}
func F_hashTypeDup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int64
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int64
	_ = v212
	var v214 int64
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	v7 = m.G0
	v9 = v7 - int32(688)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11&int32(15) != int32(4) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a2348), int32(_a2349), int32(828))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L7
	} else {
		goto L87
	}
L2:
	;
	switch int32(base.Ui32(v11)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		goto L6
	default:
		goto L5
	case 9:
		goto L4
	}
L3:
	;
	m.G0 = v9 + int32(688)
	return v279
L4:
	;
	v259 = F_objectGetVal(m, l0)
	mBase = m.M
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	goto L81
L5:
	;
	F__serverPanic_1(m, int32(_a2349), int32(859), int32(_a2355), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L7
	} else {
		goto L80
	}
L6:
	;
	v23 = F_hashtableCreate(m, int32(_a2352))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v27 = F_objectGetVal(m, l0)
	mBase = m.M
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	goto L9
L9:
	;
	v31 = F_hashtableExpand(m, v23, v28+v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v34 = F_createObject(m, int32(4), v23)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v36&int32(-241) | int32(32)
	F_hashTypeInitIterator(m, l0, v9+int32(8))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v48 = F_hashTypeNext(m, v9+int32(8))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v239 != int32(2) {
		v279 = v34
		goto L3
	} else {
		goto L75
	}
L14:
	;
	if v48 == int32(-1) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	goto L16
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+680))
	goto L18
L17:
	;
	goto L13
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+680))
	v61 = v9 + int32(4)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+int32(-1)))))
	v69 = v67 & int32(7)
	if v69 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v9)+680))
	v176 = int64(-1)
	v178 = v172 + int32(-1)
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v179&int32(7) == int32(0) {
		v214 = v176
		goto L52
	} else {
		goto L53
	}
L20:
	;
	v171 = v163
	goto L19
L21:
	;
	v124 = F_sdsAllocPtr(m, v59)
	mBase = m.M
	v126 = v124 + int32(-4)
	if v67&int32(32) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L22:
	;
	switch v69 {
	case 0:
		goto L30
	case 1:
		goto L29
	case 2:
		goto L28
	case 3:
		goto L27
	case 4:
		goto L26
	default:
		v89 = int32(0)
		goto L25
	}
L23:
	;
	if v67&int32(16) != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v91 = int32(1)
	v92 = F_sdsHdrSize(m, v91)
	mBase = m.M
	v93 = v59 + v89 + v92
	v95 = v93 + v91
	if v61 == int32(0) {
		v163 = v95
		goto L20
	} else {
		goto L31
	}
L26:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v59+int32(-17))))
	v89 = v88
	goto L25
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v59+int32(-9))))
	v89 = v85
	goto L25
L28:
	;
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59+int32(-5)))))
	v89 = v82
	goto L25
L29:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+int32(-3)))))
	v89 = v79
	goto L25
L30:
	;
	v89 = int32(base.Ui32(v67) >> (uint(int32(3)) % 32))
	goto L25
L31:
	;
	v98 = int32(0)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v98))))
	switch v101 & int32(7) {
	case 0:
		goto L37
	case 1:
		goto L36
	case 2:
		goto L35
	case 3:
		goto L34
	case 4:
		goto L33
	default:
		v122 = v98
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v122
	v171 = v95
	goto L19
L33:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v93+int32(-16))))
	v122 = v121
	goto L32
L34:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v93+int32(-8))))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v117
	v171 = v95
	goto L19
L35:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93+int32(-4)))))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v113
	v171 = v95
	goto L19
L36:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+int32(-2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v109
	v171 = v95
	goto L19
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(base.Ui32(v101) >> (uint(int32(3)) % 32))
	v171 = v95
	goto L19
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v61 == int32(0) {
		v163 = v138
		goto L20
	} else {
		goto L44
	}
L39:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v131 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if v61 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v171 = int32(0)
	goto L19
L42:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v171 = v137
	goto L19
L43:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v135
	goto L42
L44:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138+int32(-1)))))
	switch v144 & int32(7) {
	case 0:
		goto L50
	case 1:
		goto L49
	case 2:
		goto L48
	case 3:
		goto L47
	case 4:
		goto L46
	default:
		v161 = int32(0)
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v161
	v163 = v138
	goto L20
L46:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v138+int32(-17))))
	v161 = v160
	goto L45
L47:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v138+int32(-9))))
	v161 = v157
	goto L45
L48:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138+int32(-5)))))
	v161 = v154
	goto L45
L49:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138+int32(-3)))))
	v161 = v151
	goto L45
L50:
	;
	v161 = int32(base.Ui32(v144) >> (uint(int32(3)) % 32))
	goto L45
L51:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v217 = F_sdsnewlen(m, v171, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L7
	} else {
		goto L67
	}
L52:
	;
	goto L51
L53:
	;
	if v179&int32(8) == int32(0) {
		v214 = v176
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v188 = F_sdsAllocPtr(m, v172)
	mBase = m.M
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if int32(base.Ui32(v191&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v196 = int32(-4)
	goto L57
L56:
	;
	v196 = int32(0)
	goto L57
L57:
	;
	v199 = v191 & int32(7)
	if v199 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v200 = v196
	goto L60
L59:
	;
	v200 = int32(0)
	goto L60
L60:
	;
	if int32(base.Ui32(v191&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v208 = int32(-8)
	goto L63
L62:
	;
	v208 = int32(0)
	goto L63
L63:
	;
	if v199 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v210 = v208
	goto L66
L65:
	;
	v210 = int32(0)
	goto L66
L66:
	;
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v188+v200+v210)))
	v214 = v212
	goto L52
L67:
	;
	v219 = F_entryCreate(m, v58, v217, v214)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	v221 = F_hashtableAdd(m, v23, v219)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	if v214 == int64(-1) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v229 = F_hashTypeNext(m, v9+int32(8))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L73
	}
L71:
	;
	F_hashTypeTrackEntry(m, v34, v219)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	if v229 != int32(-1) {
		goto L16
	} else {
		goto L74
	}
L74:
	;
	goto L17
L75:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
	if v242 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_vsetResetIterator(m, v9+int32(80))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L7
	} else {
		goto L79
	}
L77:
	;
	F_hashtableCleanupIterator(m, v9+int32(32))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L7
	} else {
		goto L78
	}
L78:
	;
	v279 = v34
	goto L3
L79:
	;
	v279 = v34
	goto L3
L80:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	v261 = F_valkey_malloc(m, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	if v260 == int32(0) {
		v266 = v261
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v267 = F_createObject(m, int32(4), v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L7
	} else {
		goto L86
	}
L84:
	;
	goto L83
L85:
	;
	v265 = F__emscripten_memcpy_bulkmem(m, v261, v259, v260)
	mBase = m.M
	v266 = v265
	goto L84
L86:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = v269&int32(-241) | int32(176)
	v279 = v267
	goto L3
L87:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hashTypeGetFromListpack(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10&int32(240) != int32(176) {
		F__serverAssert(m, int32(_a2350), int32(_a2349), int32(201))
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v15 = F_objectGetVal(m, l0)
		mBase = m.M
		v16 = F_lpFirst(m, v15)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v16 != 0 {
				v22 = int32(-1)
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v22))))
				switch v26 & int32(7) {
				case 0:
					v43 = int32(base.Ui32(v26) >> (uint(int32(3)) % 32))
				case 1:
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
					v43 = v33
				case 2:
					v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
					v43 = v36
				case 3:
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
					v43 = v39
				case 4:
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
					v43 = v42
				default:
					v43 = int32(0)
				}
				v45 = F_lpFind(m, v15, v16, l1, v43, int32(1))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					if v45 == int32(0) {
						v58 = v22
						return v58
					} else {
						v49 = F_lpNext(m, v15, v45)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 == int32(0) {
								F__serverAssert(m, int32(_a2351), int32(_a2349), int32(210))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v53 = F_lpGetValue(m, v49, l3, l4)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53
									v58 = int32(0)
									return v58
								}
							}
						}
					}
				}
			} else {
				return int32(-1)
			}
		}
	}
}
func F_hashTypeHasVolatileFields(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	v2 = int32(0)
	if l0 == v2 {
		v47 = v2
		return v47
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7&int32(15) != int32(4) {
			F__serverAssert(m, int32(_a2348), int32(_a2349), int32(72))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			if v7&int32(240) != int32(32) {
				v47 = v2
				return v47
			} else {
				v16 = F_objectGetVal(m, l0)
				mBase = m.M
				v18 = v16 + int32(44)
				if v18 == int32(0) {
					v32 = int32(0)
				} else {
					v22 = int32(1)
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					switch v23 + v22 {
					case 0:
						v32 = v22
					case 1:
						v32 = int32(0)
					default:
						if v23&int32(7) != 0 {
							v32 = v22
						} else {
							v32 = int32(0)
						}
					}
				}
				if v18 == int32(0) {
					v47 = int32(0)
					return v47
				} else {
					if v32 == int32(0) {
						v47 = int32(0)
						return v47
					} else {
						v38 = F_vsetIsEmpty(m, v18)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							if v38 == int32(0) {
								v47 = int32(1)
							} else {
								v47 = int32(0)
							}
							return v47
						}
					}
				}
			}
		}
	}
}
func F_hashTypeNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v5 + int32(-2) {
	case 0:
		v32 = l0 + int32(672)
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		if v33 != 0 {
			v42 = F_vsetNext(m, l0+int32(72), v32)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				if v42 != 0 {
					v75 = int32(0)
					return v75
				} else {
					return int32(-1)
				}
			}
		} else {
			v36 = F_hashtableNext(m, l0+int32(24), v32)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				if v36 != 0 {
					v75 = int32(0)
					return v75
				} else {
					return int32(-1)
				}
			}
		}
	default:
		F__serverPanic_1(m, int32(_a2349), int32(704), int32(_a2355), int32(0))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 9:
		v8 = int32(-1)
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		if v9 != 0 {
			v75 = v8
			return v75
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v11 = F_objectGetVal(m, v10)
			mBase = m.M
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v13 != 0 {
				if v12 == int32(0) {
					F__serverAssert(m, int32(_a2351), int32(_a2349), int32(685))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v20 = F_lpNext(m, v11, v12)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = v20
						if v22 == int32(0) {
							v75 = v8
							return v75
						} else {
							v25 = F_lpNext(m, v11, v22)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								if v25 == int32(0) {
									F__serverAssert(m, int32(_a2351), int32(_a2349), int32(692))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v25
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v22
									v75 = int32(0)
									return v75
								}
							}
						}
					}
				}
			} else {
				if v12 != 0 {
					F__serverAssert(m, int32(_a2357), int32(_a2349), int32(681))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v14 = F_lpFirst(m, v11)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return int32(0)
					} else {
						v22 = v14
						if v22 == int32(0) {
							v75 = v8
							return v75
						} else {
							v25 = F_lpNext(m, v11, v22)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								if v25 == int32(0) {
									F__serverAssert(m, int32(_a2351), int32(_a2349), int32(692))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v25
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v22
									v75 = int32(0)
									return v75
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_hashTypeSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int64
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int64
	_ = v293
	var v295 int64
	_ = v295
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int64
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int64
	_ = v326
	var v329 int64
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v17&int32(240) != int32(176) {
		v80 = v17
		goto L1
	} else {
		goto L2
	}
L1:
	;
	switch int32(base.Ui32(v80)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		goto L29
	default:
		goto L26
	case 9:
		goto L30
	}
L2:
	;
	if l3 != int64(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_hashTypeConvert(m, l0, int32(2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v26 & int32(7) {
	case 0:
		goto L11
	case 1:
		goto L10
	case 2:
		goto L9
	case 3:
		goto L8
	case 4:
		goto L7
	default:
		goto L12
	}
L5:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
	switch v53 & int32(7) {
	case 0:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	default:
		v80 = v17
		goto L1
	}
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[550]))
	if base.Ui32(v47) < base.Ui32(v45) {
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v45 = v44
	goto L6
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v45 = v41
	goto L6
L9:
	;
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v45 = v38
	goto L6
L10:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v45 = v35
	goto L6
L11:
	;
	v45 = int32(base.Ui32(v26) >> (uint(int32(3)) % 32))
	goto L6
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[550]))
	v50 = v30
	goto L5
L13:
	;
	v50 = v47
	goto L5
L14:
	;
	if base.Ui32(v70) <= base.Ui32(v50) {
		v80 = v17
		goto L1
	} else {
		goto L20
	}
L15:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
	v70 = v69
	goto L14
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
	v70 = v66
	goto L14
L17:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
	v70 = v63
	goto L14
L18:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
	v70 = v60
	goto L14
L19:
	;
	v70 = int32(base.Ui32(v53) >> (uint(int32(3)) % 32))
	goto L14
L20:
	;
	goto L3
L21:
	;
	return int32(0)
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v80 = v79
	goto L1
L23:
	;
	if l1 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L24:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v371&int32(240) != int32(32) {
		v402 = v223
		v404 = v366
		v405 = v367
		goto L23
	} else {
		goto L127
	}
L25:
	;
	v363 = int32(0)
	v366 = v363
	v367 = v363
	goto L24
L26:
	;
	F__serverPanic_1(m, int32(_a2349), int32(454), int32(_a2355), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L21
	} else {
		goto L126
	}
L27:
	;
	F__serverAssert(m, int32(_a2354), int32(_a2349), int32(444))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L21
	} else {
		goto L125
	}
L28:
	;
	F__serverAssert(m, int32(_a2351), int32(_a2349), int32(389))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L21
	} else {
		goto L124
	}
L29:
	;
	v215 = F_objectGetVal(m, l0)
	mBase = m.M
	if l4&int32(2) == int32(0) {
		goto L72
	} else {
		goto L73
	}
L30:
	;
	v89 = F_objectGetVal(m, l0)
	mBase = m.M
	v90 = F_lpFirst(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L21
	} else {
		goto L33
	}
L31:
	;
	F_objectSetVal(m, l0, v200)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L21
	} else {
		goto L66
	}
L32:
	;
	v153 = int32(0)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v157 & int32(7) {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	case 4:
		goto L53
	default:
		v174 = v153
		goto L52
	}
L33:
	;
	if v90 == int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v97 & int32(7) {
	case 0:
		goto L40
	case 1:
		goto L39
	case 2:
		goto L38
	case 3:
		goto L37
	case 4:
		goto L36
	default:
		v114 = int32(0)
		goto L35
	}
L35:
	;
	v116 = F_lpFind(m, v89, v90, l1, v114, int32(1))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L21
	} else {
		goto L41
	}
L36:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v114 = v113
	goto L35
L37:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v114 = v110
	goto L35
L38:
	;
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v114 = v107
	goto L35
L39:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v114 = v104
	goto L35
L40:
	;
	v114 = int32(base.Ui32(v97) >> (uint(int32(3)) % 32))
	goto L35
L41:
	;
	if v116 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L42:
	;
	v120 = F_lpNext(m, v89, v116)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L21
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v120
	if v120 == int32(0) {
		goto L28
	} else {
		goto L44
	}
L44:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
	switch v128 & int32(7) {
	case 0:
		goto L50
	case 1:
		goto L49
	case 2:
		goto L48
	case 3:
		goto L47
	case 4:
		goto L46
	default:
		v145 = int32(0)
		goto L45
	}
L45:
	;
	v149 = F_lpReplace(m, v89, v15+int32(16), l2, v145)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L21
	} else {
		goto L51
	}
L46:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
	v145 = v144
	goto L45
L47:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
	v145 = v141
	goto L45
L48:
	;
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
	v145 = v138
	goto L45
L49:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
	v145 = v135
	goto L45
L50:
	;
	v145 = int32(base.Ui32(v128) >> (uint(int32(3)) % 32))
	goto L45
L51:
	;
	v200 = v149
	v201 = int32(1)
	goto L31
L52:
	;
	v175 = F_lpAppend(m, v89, l1, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L21
	} else {
		goto L58
	}
L53:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v174 = v173
	goto L52
L54:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v174 = v170
	goto L52
L55:
	;
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v174 = v167
	goto L52
L56:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v174 = v164
	goto L52
L57:
	;
	v174 = int32(base.Ui32(v157) >> (uint(int32(3)) % 32))
	goto L52
L58:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
	switch v179 & int32(7) {
	case 0:
		goto L64
	case 1:
		goto L63
	case 2:
		goto L62
	case 3:
		goto L61
	case 4:
		goto L60
	default:
		v196 = v153
		goto L59
	}
L59:
	;
	v198 = F_lpAppend(m, v175, l2, v196)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L21
	} else {
		goto L65
	}
L60:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
	v196 = v195
	goto L59
L61:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
	v196 = v192
	goto L59
L62:
	;
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
	v196 = v189
	goto L59
L63:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
	v196 = v186
	goto L59
L64:
	;
	v196 = int32(base.Ui32(v179) >> (uint(int32(3)) % 32))
	goto L59
L65:
	;
	v200 = v198
	v201 = int32(0)
	goto L31
L66:
	;
	v206 = F_hashTypeLength(m, l0)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L21
	} else {
		goto L68
	}
L67:
	;
	v402 = l2
	v404 = int32(0)
	v405 = v201
	goto L23
L68:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _consts[549]))
	if base.Ui32(v206) <= base.Ui32(v209) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	F_hashTypeConvert(m, l0, int32(2))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L21
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v225&int32(240) != int32(32) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v221 = F_sdsdup(m, l2)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L21
	} else {
		goto L74
	}
L73:
	;
	v223 = int32(0)
	v224 = l2
	goto L71
L74:
	;
	v223 = l2
	v224 = v221
	goto L71
L75:
	;
	v239 = F_hashtableFindPositionForInsert(m, v215, l1, v15+int32(16), v15+int32(12))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L21
	} else {
		goto L79
	}
L76:
	;
	v230 = F_objectGetVal(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v230))) = int32(_a2352)
	goto L77
L77:
	;
	goto L75
L78:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v257 = int64(-1)
	v259 = v253 + int32(-1)
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v260&int32(7) == int32(0) {
		v295 = v257
		goto L88
	} else {
		goto L89
	}
L79:
	;
	if v239 == int32(0) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v243 = F_entryCreate(m, l1, v224, l3)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L21
	} else {
		goto L81
	}
L81:
	;
	F_hashtableInsertAtPosition(m, v215, v243, v15+int32(16))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L21
	} else {
		goto L82
	}
L82:
	;
	if l3 == int64(-1) {
		goto L25
	} else {
		goto L83
	}
L83:
	;
	F_hashTypeTrackEntry(m, l0, v243)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L21
	} else {
		goto L84
	}
L84:
	;
	goto L25
L85:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v324 != 0 {
		goto L113
	} else {
		goto L114
	}
L86:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v303 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L87:
	;
	if v295 != int64(-1) {
		goto L86
	} else {
		goto L103
	}
L88:
	;
	goto L87
L89:
	;
	if v260&int32(8) == int32(0) {
		v295 = v257
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v269 = F_sdsAllocPtr(m, v253)
	mBase = m.M
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if int32(base.Ui32(v272&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v277 = int32(-4)
	goto L93
L92:
	;
	v277 = int32(0)
	goto L93
L93:
	;
	v280 = v272 & int32(7)
	if v280 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v281 = v277
	goto L96
L95:
	;
	v281 = int32(0)
	goto L96
L96:
	;
	if int32(base.Ui32(v272&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v289 = int32(-8)
	goto L99
L98:
	;
	v289 = int32(0)
	goto L99
L99:
	;
	if v280 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v291 = v289
	goto L102
L101:
	;
	v291 = int32(0)
	goto L102
L102:
	;
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v269+v281+v291)))
	v295 = v293
	goto L88
L103:
	;
	v324 = int32(0)
	goto L85
L104:
	;
	v324 = base.B2i32(v321 != int32(0))
	goto L85
L105:
	;
	goto L104
L106:
	;
	v309 = int32(0)
	v310 = F_commandTimeSnapshot(m)
	mBase = m.M
	if v310 < v295 {
		v321 = v309
		goto L105
	} else {
		goto L109
	}
L107:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v303)+216))
	if v307 != 0 {
		v321 = int32(0)
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v312 = int32(_a20)
	v313 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v315 = *(*int32)(unsafe.Add(mBase, _consts[384]))
	if v315 != 0 {
		v321 = v309
		goto L105
	} else {
		goto L110
	}
L110:
	;
	if v313 != 0 {
		v321 = v309
		goto L105
	} else {
		goto L111
	}
L111:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _consts[505]))
	v321 = base.B2i32(v317 == int32(0))
	goto L105
L112:
	;
	F_hashTypeTrackUpdateEntry(m, l0, v339, v330, v295, v329)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L21
	} else {
		goto L123
	}
L113:
	;
	v326 = l3
	goto L115
L114:
	;
	v326 = v295
	goto L115
L115:
	;
	if l4&int32(4) != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v329 = v326
	goto L118
L117:
	;
	v329 = l3
	goto L118
L118:
	;
	v330 = F_entryUpdate(m, v325, v224, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L21
	} else {
		goto L119
	}
L119:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v330 == v332 {
		v339 = v332
		goto L112
	} else {
		goto L120
	}
L120:
	;
	v334 = F_hashtableReplaceReallocatedEntry(m, v215, v332, v330)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L21
	} else {
		goto L121
	}
L121:
	;
	if v334 == int32(0) {
		goto L27
	} else {
		goto L122
	}
L122:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v339 = v338
	goto L112
L123:
	;
	v366 = v324
	v367 = v324 ^ int32(1)
	goto L24
L124:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	v376 = F_objectGetVal(m, l0)
	mBase = m.M
	v378 = v376 + int32(44)
	goto L128
L128:
	;
	if v378 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	v393 = F_objectGetVal(m, l0)
	mBase = m.M
	if v392 != 0 {
		goto L135
	} else {
		goto L136
	}
L130:
	;
	goto L129
L131:
	;
	v392 = int32(0)
	goto L130
L132:
	;
	v382 = int32(1)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	switch v383 + v382 {
	case 0:
		v392 = v382
		goto L130
	case 1:
		goto L131
	default:
		goto L133
	}
L133:
	;
	if v383&int32(7) != 0 {
		v392 = v382
		goto L130
	} else {
		goto L134
	}
L134:
	;
	goto L131
L135:
	;
	v396 = int32(_a2353)
	goto L137
L136:
	;
	v396 = int32(_a2352)
	goto L137
L137:
	;
	if v378 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v398 = v396
	goto L140
L139:
	;
	v398 = int32(_a2352)
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v393))) = v398
	goto L141
L141:
	;
	v402 = v223
	v404 = v366
	v405 = v367
	goto L23
L142:
	;
	if l4&int32(2) == int32(0) {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	if l4&int32(1) == int32(0) {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	F_sdsfree(m, l1)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L21
	} else {
		goto L145
	}
L145:
	;
	goto L142
L146:
	;
	if l5 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	if v402 == int32(0) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	F_sdsfree(m, v402)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L21
	} else {
		goto L149
	}
L149:
	;
	goto L146
L150:
	;
	m.G0 = v15 + int32(32)
	return v405
L151:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v404)
	goto L150
}
