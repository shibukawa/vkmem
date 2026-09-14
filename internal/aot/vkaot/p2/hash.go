package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_hashHashtableTypeValidate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v4 = F_getExpirationPolicyWithFlags(m, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 != 0 {
			v10 = int32(0)
			v16 = l1 + int32(-1)
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			if v17&int32(7) == v10 {
				v56 = v10
			} else {
				if v17&int32(8) == int32(0) {
					v56 = v10
				} else {
					v26 = int32(0)
					v27 = F_sdsAllocPtr(m, l1)
					mBase = m.M
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
					if int32(base.Ui32(v30&int32(16))>>(uint(int32(4))%32)) != 0 {
						v35 = int32(-4)
					} else {
						v35 = v26
					}
					v38 = v30 & int32(7)
					if v38 != 0 {
						v39 = v35
					} else {
						v39 = int32(0)
					}
					if int32(base.Ui32(v30&int32(8))>>(uint(int32(3))%32)) != 0 {
						v47 = int32(-8)
					} else {
						v47 = int32(0)
					}
					if v38 != 0 {
						v49 = v47
					} else {
						v49 = int32(0)
					}
					v51 = *(*int64)(unsafe.Add(mBase, uint32(v27+v39+v49)))
					if v51 == int64(-1) {
						v56 = v26
					} else {
						v54 = F_timestampIsExpired(m, v51)
						mBase = m.M
						v56 = v54
					}
				}
			}
			return v56 ^ int32(1)
		} else {
			return int32(1)
		}
	}
}
func F_hashTypeConvert(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = int32(base.Ui32(v4)>>(uint(int32(4))%32)) & int32(15)
	if v8 == int32(11) {
		F_hashTypeConvertListpack(m, l0, l1)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			return
		}
	} else {
		if v8 == int32(2) {
			F__serverPanic_1(m, int32(_a1652), int32(813), int32(_a1052), int32(0))
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
		} else {
			F__serverPanic_1(m, int32(_a1652), int32(815), int32(_a176), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_hashTypeConvertListpack(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
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
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	v6 = m.G0
	v8 = v6 - int32(688)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10&int32(240) != int32(176) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverPanic_1(m, int32(_a1652), int32(805), int32(_a176), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L38
	}
L2:
	;
	F__serverAssert(m, int32(_a1655), int32(_a1652), int32(771))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L37
	}
L3:
	;
	switch l1 + int32(-2) {
	case 0:
		goto L5
	default:
		goto L1
	case 9:
		goto L4
	}
L4:
	;
	m.G0 = v8 + int32(688)
	return
L5:
	;
	v18 = F_hashtableCreate(m, int32(_a1654))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v20 = F_hashTypeLength(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v22 = F_hashtableExpand(m, v18, v20)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_hashTypeInitIterator(m, l0, v8+int32(8))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L12
L11:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v78 != int32(2) {
		goto L29
	} else {
		goto L30
	}
L12:
	;
	v35 = F_hashTypeNext(m, v8+int32(8))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L14
	}
L13:
	;
	F_entryFree(m, v50)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L6
	} else {
		goto L22
	}
L14:
	;
	if v35 == int32(-1) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v42 = F_hashTypeCurrentObjectNewSds(m, v8+int32(8), int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v47 = F_hashTypeCurrentObjectNewSds(m, v8+int32(8), int32(2))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v50 = F_entryCreate(m, v42, v47, int64(-1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	F_sdsfree(m, v42)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v54 = F_hashtableAdd(m, v18, v50)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	if v54 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	goto L13
L22:
	;
	F_hashTypeResetIterator(m, v8+int32(8))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[209]))
	if v63 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F__serverPanic_1(m, int32(_a1652), int32(797), int32(_a1656), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L28
	}
L25:
	;
	v66 = F_objectGetVal(m, l0)
	mBase = m.M
	v67 = F_objectGetVal(m, l0)
	mBase = m.M
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	goto L26
L26:
	;
	F_serverLogHexDump(m, int32(3), int32(_a1657), v66, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v90 = F_objectGetVal(m, l0)
	mBase = m.M
	F_valkey_free(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L35
	}
L30:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
	if v81 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_vsetResetIterator(m, v8+int32(80))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L34
	}
L32:
	;
	F_hashtableCleanupIterator(m, v8+int32(32))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	goto L29
L35:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v93&int32(-241) | int32(32)
	F_objectSetVal(m, l0, v18)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	goto L4
L37:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hashTypeCurrentFromHashTable(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 != int32(2) {
		F__serverAssert(m, int32(_a1660), int32(_a1652), int32(729))
		mBase = m.M
		v161 = m.ExcPending
		if v161 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+672))
		if l1&int32(1) == int32(0) {
			v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-1)))))
			v51 = v49 & int32(7)
			if v51 == int32(0) {
				switch v51 {
				case 0:
					v71 = int32(base.Ui32(v49) >> (uint(int32(3)) % 32))
				case 1:
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-3)))))
					v71 = v61
				case 2:
					v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-5)))))
					v71 = v64
				case 3:
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-9))))
					v71 = v67
				case 4:
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-17))))
					v71 = v70
				default:
					v71 = int32(0)
				}
				v73 = int32(1)
				v74 = F_sdsHdrSize(m, v73)
				mBase = m.M
				v75 = v8 + v71 + v74
				v77 = v75 + v73
				if l2 == int32(0) {
					v145 = v77
					v153 = v145
				} else {
					v80 = int32(0)
					v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v80))))
					switch v83 & int32(7) {
					case 0:
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(base.Ui32(v83) >> (uint(int32(3)) % 32))
						v153 = v77
					case 1:
						v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+int32(-2)))))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v91
						v153 = v77
					case 2:
						v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75+int32(-4)))))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v95
						v153 = v77
					case 3:
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(-8))))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v99
						v153 = v77
					case 4:
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(-16))))
						v104 = v103
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v104
						v153 = v77
					default:
						v104 = v80
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v104
						v153 = v77
					}
				}
			} else {
				if v49&int32(16) != 0 {
					v106 = F_sdsAllocPtr(m, v8)
					mBase = m.M
					v108 = v106 + int32(-4)
					if v49&int32(32) == int32(0) {
						v120 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
						if l2 == int32(0) {
							v145 = v120
						} else {
							v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120+int32(-1)))))
							switch v126 & int32(7) {
							case 0:
								v143 = int32(base.Ui32(v126) >> (uint(int32(3)) % 32))
							case 1:
								v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120+int32(-3)))))
								v143 = v133
							case 2:
								v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120+int32(-5)))))
								v143 = v136
							case 3:
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v120+int32(-9))))
								v143 = v139
							case 4:
								v142 = *(*int32)(unsafe.Add(mBase, uint32(v120+int32(-17))))
								v143 = v142
							default:
								v143 = int32(0)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v143
							v145 = v120
						}
						v153 = v145
					} else {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
						if v113 != 0 {
							if l2 == int32(0) {
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v117
							}
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
							v153 = v119
						} else {
							v153 = int32(0)
						}
					}
				} else {
					switch v51 {
					case 0:
						v71 = int32(base.Ui32(v49) >> (uint(int32(3)) % 32))
					case 1:
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-3)))))
						v71 = v61
					case 2:
						v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-5)))))
						v71 = v64
					case 3:
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-9))))
						v71 = v67
					case 4:
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-17))))
						v71 = v70
					default:
						v71 = int32(0)
					}
					v73 = int32(1)
					v74 = F_sdsHdrSize(m, v73)
					mBase = m.M
					v75 = v8 + v71 + v74
					v77 = v75 + v73
					if l2 == int32(0) {
						v145 = v77
						v153 = v145
					} else {
						v80 = int32(0)
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v80))))
						switch v83 & int32(7) {
						case 0:
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(base.Ui32(v83) >> (uint(int32(3)) % 32))
							v153 = v77
						case 1:
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+int32(-2)))))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v91
							v153 = v77
						case 2:
							v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75+int32(-4)))))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v95
							v153 = v77
						case 3:
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(-8))))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v99
							v153 = v77
						case 4:
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(-16))))
							v104 = v103
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v104
							v153 = v77
						default:
							v104 = v80
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v104
							v153 = v77
						}
					}
				}
			}
			return v153
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-1)))))
			switch v16 & int32(7) {
			case 0:
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(base.Ui32(v16) >> (uint(int32(3)) % 32))
				return v8
			case 1:
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-3)))))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v25
				return v8
			case 2:
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-5)))))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
				return v8
			case 3:
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-9))))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v35
				return v8
			case 4:
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-17))))
				v41 = v40
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v41
				return v8
			default:
				v41 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v41
				return v8
			}
		}
	}
}
func F_hashTypeDeleteExpiredFields(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13&int32(240) != int32(32) {
		F__serverAssert(m, int32(_a1651), int32(_a1652), int32(2431))
		mBase = m.M
		v139 = m.ExcPending
		if v139 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v18 = int32(0)
		v19 = F_objectGetVal(m, l0)
		mBase = m.M
		v21 = v19 + int32(44)
		if v21 == int32(0) {
			v35 = int32(0)
		} else {
			v25 = int32(1)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			switch v26 + v25 {
			case 0:
				v35 = v25
			case 1:
				v35 = int32(0)
			default:
				if v26&int32(7) != 0 {
					v35 = v25
				} else {
					v35 = int32(0)
				}
			}
		}
		if v21 == int32(0) {
			v129 = v18
			m.G0 = v11 + int32(16)
			return v129
		} else {
			if v35 == int32(0) {
				v129 = v18
				m.G0 = v11 + int32(16)
				return v129
			} else {
				v40 = F_vsetIsEmpty(m, v21)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					if v40 != 0 {
						F__serverAssert(m, int32(_a1671), int32(_a1652), int32(2438))
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v44&int32(240) != int32(32) {
						} else {
							v49 = F_objectGetVal(m, l0)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v49))) = int32(_a1654)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l0
						v62 = F_vsetRemoveExpired(m, v21, int32(1087), int32(1088), l1, l2, v11+int32(4))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
							if base.Ui32(l2) < base.Ui32(v64) {
								F__serverAssert(m, int32(_a1672), int32(_a1652), int32(2443))
								mBase = m.M
								v151 = m.ExcPending
								if v151 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v66 = F_vsetIsEmpty(m, v21)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									if v66 == int32(0) {
										v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										if v97&int32(240) != int32(32) {
											v129 = v62
										} else {
											v104 = F_objectGetVal(m, l0)
											mBase = m.M
											v106 = v104 + int32(44)
											if v106 == int32(0) {
												v120 = int32(0)
											} else {
												v110 = int32(1)
												v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
												switch v111 + v110 {
												case 0:
													v120 = v110
												case 1:
													v120 = int32(0)
												default:
													if v111&int32(7) != 0 {
														v120 = v110
													} else {
														v120 = int32(0)
													}
												}
											}
											if v120 != 0 {
												v121 = int32(_a1653)
											} else {
												v121 = int32(_a1654)
											}
											if v106 != 0 {
												v123 = v121
											} else {
												v123 = int32(_a1654)
											}
											v124 = v123
											v125 = F_objectGetVal(m, l0)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v125))) = v124
											v129 = v62
										}
										m.G0 = v11 + int32(16)
										return v129
									} else {
										v70 = F_objectGetVal(m, l0)
										mBase = m.M
										v72 = v70 + int32(44)
										if v72 == int32(0) {
											v86 = int32(0)
										} else {
											v76 = int32(1)
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
											switch v77 + v76 {
											case 0:
												v86 = v76
											case 1:
												v86 = int32(0)
											default:
												if v77&int32(7) != 0 {
													v86 = v76
												} else {
													v86 = int32(0)
												}
											}
										}
										if v86 == int32(0) {
											v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											if v92&int32(240) == int32(32) {
												v124 = int32(_a1654)
												v125 = F_objectGetVal(m, l0)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v125))) = v124
												v129 = v62
											} else {
												v129 = v62
											}
											m.G0 = v11 + int32(16)
											return v129
										} else {
											F_vsetRelease(m, v72)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												if v92&int32(240) == int32(32) {
													v124 = int32(_a1654)
													v125 = F_objectGetVal(m, l0)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v125))) = v124
													v129 = v62
												} else {
													v129 = v62
												}
												m.G0 = v11 + int32(16)
												return v129
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_hashTypeExpireEntry(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v11&int32(240) != int32(32) {
		F__serverAssert(m, int32(_a1673), int32(_a1652), int32(2414))
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
		v16 = F_objectGetVal(m, v10)
		mBase = m.M
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
		if v17+v18 == int32(0) {
			F__serverAssert(m, int32(_a1673), int32(_a1652), int32(2414))
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
			v22 = F_objectGetVal(m, v10)
			mBase = m.M
			v23 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v23
			v28 = F_hashtablePop(m, v22, l0, v8+int32(12))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				if v28 == int32(0) {
					v59 = v23
					m.G0 = v8 + int32(16)
					return v59
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if v34 == int32(0) {
						v50 = int32(_a44)
						v52 = *(*int64)(unsafe.Add(mBase, _consts[914]))
						*(*int64)(unsafe.Add(mBase, _consts[914])) = v52 + int64(1)
						F_entryFree(m, l0)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v59 = int32(1)
							m.G0 = v8 + int32(16)
							return v59
						}
					} else {
						v37 = F_createStringObjectFromSds(m, l0)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v39 + int32(1)
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v43+v39<<(uint(int32(2))%32)))) = v37
							v50 = int32(_a44)
							v52 = *(*int64)(unsafe.Add(mBase, _consts[914]))
							*(*int64)(unsafe.Add(mBase, _consts[914])) = v52 + int64(1)
							F_entryFree(m, l0)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								v59 = int32(1)
								m.G0 = v8 + int32(16)
								return v59
							}
						}
					}
				}
			}
		}
	}
}
func F_hashTypeGetOrcreateVolatileSet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5&int32(240) != int32(32) {
		F__serverAssert(m, int32(_a1651), int32(_a1652), int32(94))
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
		v10 = F_objectGetVal(m, l0)
		mBase = m.M
		v12 = v10 + int32(44)
		if v12 == int32(0) {
			v26 = int32(0)
		} else {
			v16 = int32(1)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			switch v17 + v16 {
			case 0:
				v26 = v16
			case 1:
				v26 = int32(0)
			default:
				if v17&int32(7) != 0 {
					v26 = v16
				} else {
					v26 = int32(0)
				}
			}
		}
		if v26 != 0 {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(-1)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v29&int32(240) != int32(32) {
			} else {
				v34 = F_objectGetVal(m, l0)
				mBase = m.M
				v36 = v34 + int32(44)
				if v36 == int32(0) {
					v50 = int32(0)
				} else {
					v40 = int32(1)
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					switch v41 + v40 {
					case 0:
						v50 = v40
					case 1:
						v50 = int32(0)
					default:
						if v41&int32(7) != 0 {
							v50 = v40
						} else {
							v50 = int32(0)
						}
					}
				}
				v51 = F_objectGetVal(m, l0)
				mBase = m.M
				if v50 != 0 {
					v54 = int32(_a1653)
				} else {
					v54 = int32(_a1654)
				}
				if v36 != 0 {
					v56 = v54
				} else {
					v56 = int32(_a1654)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v51))) = v56
			}
		}
		return v12
	}
}
func F_hashTypeInitIterator(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v41 int32
	_ = v41
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = l0
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v3)
	v11 = int32(base.Ui32(v5)>>(uint(int32(4))%32)) & int32(15)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v11
	switch v11 + int32(-2) {
	case 0:
		v16 = l1 + int32(24)
		v17 = F_objectGetVal(m, l0)
		mBase = m.M
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+14)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v18
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(-1)
		if v17 == v18 {
		} else {
		}
		return
	default:
		F__serverPanic_1(m, int32(_a1652), int32(638), int32(_a176), int32(0))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 9:
		*(*int64)(unsafe.Add(mBase, uint32(l1)+12)) = int64(0)
		return
	}
}
func F_hashTypePersist(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(-2)
	if l0 == int32(0) {
		v98 = v10
		m.G0 = v8 + int32(16)
		return v98
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v13&int32(15) != int32(4) {
			v98 = v10
			m.G0 = v8 + int32(16)
			return v98
		} else {
			if v13&int32(240) != int32(176) {
				v38 = F_objectGetVal(m, l0)
				mBase = m.M
				v39 = F_hashtableFindRef(m, v38, l1)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					if v39 == int32(0) {
						v98 = v10
						m.G0 = v8 + int32(16)
						return v98
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
						v47 = int64(-1)
						v49 = v43 + int32(-1)
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
						if v50&int32(7) == int32(0) {
							v85 = v47
						} else {
							if v50&int32(8) == int32(0) {
								v85 = v47
							} else {
								v59 = F_sdsAllocPtr(m, v43)
								mBase = m.M
								v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
								if int32(base.Ui32(v62&int32(16))>>(uint(int32(4))%32)) != 0 {
									v67 = int32(-4)
								} else {
									v67 = int32(0)
								}
								v70 = v62 & int32(7)
								if v70 != 0 {
									v71 = v67
								} else {
									v71 = int32(0)
								}
								if int32(base.Ui32(v62&int32(8))>>(uint(int32(3))%32)) != 0 {
									v79 = int32(-8)
								} else {
									v79 = int32(0)
								}
								if v70 != 0 {
									v81 = v79
								} else {
									v81 = int32(0)
								}
								v83 = *(*int64)(unsafe.Add(mBase, uint32(v59+v71+v81)))
								v85 = v83
							}
						}
						if v85 != int64(-1) {
							F_hashTypeUntrackEntry(m, l0, v43)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								v93 = F_entrySetExpiry(m, v43, int64(-1))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v39))) = v93
									v98 = int32(1)
									m.G0 = v8 + int32(16)
									return v98
								}
							}
						} else {
							v98 = int32(-1)
							m.G0 = v8 + int32(16)
							return v98
						}
					}
				}
			} else {
				v22 = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v22
				*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
				v33 = F_hashTypeGetValue(m, l0, l1, v8+int32(12), v8+int32(8), v8, int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					if v33 != 0 {
						v37 = int32(-2)
					} else {
						v37 = v22
					}
					v98 = v37
					m.G0 = v8 + int32(16)
					return v98
				}
			}
		}
	}
}
func F_hashTypeRandomElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int64
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v14)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		goto L4
	default:
		goto L3
	case 9:
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v381
L2:
	;
	v372 = F_objectGetVal(m, l0)
	mBase = m.M
	F_lpRandomPair(m, v372, l1, l2, l3)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L11
	} else {
		goto L112
	}
L3:
	;
	F__serverPanic_1(m, int32(_a1652), int32(911), int32(_a176), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L11
	} else {
		goto L111
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	if v14&int32(240) != int32(32) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v33 = l3 + int32(4)
	v36 = int32(100)
	goto L9
L6:
	;
	v27 = F_objectGetVal(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(_a1654)
	goto L7
L7:
	;
	goto L5
L8:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v334&int32(240) != int32(32) {
		v381 = v330
		goto L1
	} else {
		goto L96
	}
L9:
	;
	v44 = F_objectGetVal(m, l0)
	mBase = m.M
	v47 = F_hashtableFairRandomEntry(m, v44, v12+int32(12))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v330 = int32(0)
	goto L8
L11:
	;
	return int32(0)
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v52 = int32(0)
	v58 = v51 + int32(-1)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v59&int32(7) == v52 {
		v98 = v52
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	goto L54
L14:
	;
	if v98 == int32(0) {
		v175 = v36
		goto L13
	} else {
		goto L31
	}
L15:
	;
	goto L14
L16:
	;
	if v59&int32(8) == int32(0) {
		v98 = v52
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v68 = int32(0)
	v69 = F_sdsAllocPtr(m, v51)
	mBase = m.M
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if int32(base.Ui32(v72&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v77 = int32(-4)
	goto L20
L19:
	;
	v77 = v68
	goto L20
L20:
	;
	v80 = v72 & int32(7)
	if v80 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v81 = v77
	goto L23
L22:
	;
	v81 = int32(0)
	goto L23
L23:
	;
	if int32(base.Ui32(v72&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v89 = int32(-8)
	goto L26
L25:
	;
	v89 = int32(0)
	goto L26
L26:
	;
	if v80 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v91 = v89
	goto L29
L28:
	;
	v91 = int32(0)
	goto L29
L29:
	;
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v69+v81+v91)))
	if v93 == int64(-1) {
		v98 = v68
		goto L15
	} else {
		goto L30
	}
L30:
	;
	v96 = F_timestampIsExpired(m, v93)
	mBase = m.M
	v98 = v96
	goto L15
L31:
	;
	v104 = v36
	goto L32
L32:
	;
	v112 = int32(-1)
	v114 = v104 + v112
	if v114 == int32(0) {
		v330 = v112
		goto L8
	} else {
		goto L34
	}
L33:
	;
	v175 = v114
	goto L13
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v119 = F_objectGetVal(m, l0)
	mBase = m.M
	v122 = F_hashtableFairRandomEntry(m, v119, v12+int32(12))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v125 = int32(0)
	v131 = v124 + int32(-1)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v132&int32(7) == v125 {
		v171 = v125
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v171 != 0 {
		v104 = v114
		goto L32
	} else {
		goto L53
	}
L37:
	;
	goto L36
L38:
	;
	if v132&int32(8) == int32(0) {
		v171 = v125
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v141 = int32(0)
	v142 = F_sdsAllocPtr(m, v124)
	mBase = m.M
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if int32(base.Ui32(v145&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v150 = int32(-4)
	goto L42
L41:
	;
	v150 = v141
	goto L42
L42:
	;
	v153 = v145 & int32(7)
	if v153 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v154 = v150
	goto L45
L44:
	;
	v154 = int32(0)
	goto L45
L45:
	;
	if int32(base.Ui32(v145&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v162 = int32(-8)
	goto L48
L47:
	;
	v162 = int32(0)
	goto L48
L48:
	;
	if v153 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v164 = v162
	goto L51
L50:
	;
	v164 = int32(0)
	goto L51
L51:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v142+v154+v164)))
	if v166 == int64(-1) {
		v171 = v141
		goto L37
	} else {
		goto L52
	}
L52:
	;
	v169 = F_timestampIsExpired(m, v166)
	mBase = m.M
	v171 = v169
	goto L37
L53:
	;
	goto L33
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v183
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+int32(-1)))))
	switch v188 & int32(7) {
	case 0:
		goto L60
	case 1:
		goto L59
	case 2:
		goto L58
	case 3:
		goto L57
	case 4:
		goto L56
	default:
		v205 = int32(0)
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v205
	if l3 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v183+int32(-17))))
	v205 = v204
	goto L55
L57:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v183+int32(-9))))
	v205 = v201
	goto L55
L58:
	;
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183+int32(-5)))))
	v205 = v198
	goto L55
L59:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+int32(-3)))))
	v205 = v195
	goto L55
L60:
	;
	v205 = int32(base.Ui32(v188) >> (uint(int32(3)) % 32))
	goto L55
L61:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v321 == int32(0) {
		v36 = v175
		goto L9
	} else {
		goto L95
	}
L62:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209+int32(-1)))))
	v217 = v215 & int32(7)
	if v217 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v319
	goto L61
L64:
	;
	v319 = v311
	goto L63
L65:
	;
	v272 = F_sdsAllocPtr(m, v209)
	mBase = m.M
	v274 = v272 + int32(-4)
	if v215&int32(32) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L66:
	;
	switch v217 {
	case 0:
		goto L74
	case 1:
		goto L73
	case 2:
		goto L72
	case 3:
		goto L71
	case 4:
		goto L70
	default:
		v237 = int32(0)
		goto L69
	}
L67:
	;
	if v215&int32(16) != 0 {
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v239 = int32(1)
	v240 = F_sdsHdrSize(m, v239)
	mBase = m.M
	v241 = v209 + v237 + v240
	v243 = v241 + v239
	if v33 == int32(0) {
		v311 = v243
		goto L64
	} else {
		goto L75
	}
L70:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v209+int32(-17))))
	v237 = v236
	goto L69
L71:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v209+int32(-9))))
	v237 = v233
	goto L69
L72:
	;
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209+int32(-5)))))
	v237 = v230
	goto L69
L73:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209+int32(-3)))))
	v237 = v227
	goto L69
L74:
	;
	v237 = int32(base.Ui32(v215) >> (uint(int32(3)) % 32))
	goto L69
L75:
	;
	v246 = int32(0)
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241+v246))))
	switch v249 & int32(7) {
	case 0:
		goto L81
	case 1:
		goto L80
	case 2:
		goto L79
	case 3:
		goto L78
	case 4:
		goto L77
	default:
		v270 = v246
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v270
	v319 = v243
	goto L63
L77:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v241+int32(-16))))
	v270 = v269
	goto L76
L78:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v241+int32(-8))))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v265
	v319 = v243
	goto L63
L79:
	;
	v261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241+int32(-4)))))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v261
	v319 = v243
	goto L63
L80:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241+int32(-2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v257
	v319 = v243
	goto L63
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(base.Ui32(v249) >> (uint(int32(3)) % 32))
	v319 = v243
	goto L63
L82:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	if v33 == int32(0) {
		v311 = v286
		goto L64
	} else {
		goto L88
	}
L83:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	if v279 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	if v33 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v319 = int32(0)
	goto L63
L86:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v319 = v285
	goto L63
L87:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v283
	goto L86
L88:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286+int32(-1)))))
	switch v292 & int32(7) {
	case 0:
		goto L94
	case 1:
		goto L93
	case 2:
		goto L92
	case 3:
		goto L91
	case 4:
		goto L90
	default:
		v309 = int32(0)
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v309
	v311 = v286
	goto L64
L90:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v286+int32(-17))))
	v309 = v308
	goto L89
L91:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v286+int32(-9))))
	v309 = v305
	goto L89
L92:
	;
	v302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v286+int32(-5)))))
	v309 = v302
	goto L89
L93:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286+int32(-3)))))
	v309 = v299
	goto L89
L94:
	;
	v309 = int32(base.Ui32(v292) >> (uint(int32(3)) % 32))
	goto L89
L95:
	;
	goto L10
L96:
	;
	v339 = F_objectGetVal(m, l0)
	mBase = m.M
	v341 = v339 + int32(44)
	goto L97
L97:
	;
	if v341 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v356 = F_objectGetVal(m, l0)
	mBase = m.M
	if v355 != 0 {
		goto L104
	} else {
		goto L105
	}
L99:
	;
	goto L98
L100:
	;
	v355 = int32(0)
	goto L99
L101:
	;
	v345 = int32(1)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	switch v346 + v345 {
	case 0:
		v355 = v345
		goto L99
	case 1:
		goto L100
	default:
		goto L102
	}
L102:
	;
	if v346&int32(7) != 0 {
		v355 = v345
		goto L99
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	v359 = int32(_a1653)
	goto L106
L105:
	;
	v359 = int32(_a1654)
	goto L106
L106:
	;
	if v341 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v361 = v359
	goto L109
L108:
	;
	v361 = int32(_a1654)
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v356))) = v361
	goto L110
L110:
	;
	v381 = v330
	goto L1
L111:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	v381 = int32(0)
	goto L1
}
func F_hashTypeUntrackEntry(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	if base.B2i32(v6&int32(7) != int32(0))&int32(base.Ui32(v6&int32(8))>>(uint(int32(3))%32)) == int32(0) {
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v18&int32(240) != int32(32) {
			F__serverAssert(m, int32(_a1651), int32(_a1652), int32(65))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v23 = F_objectGetVal(m, l0)
			mBase = m.M
			v25 = v23 + int32(44)
			v26 = int32(0)
			if v25 == v26 {
				v40 = int32(0)
			} else {
				v30 = int32(1)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				switch v31 + v30 {
				case 0:
					v40 = v30
				case 1:
					v40 = int32(0)
				default:
					if v31&int32(7) != 0 {
						v40 = v30
					} else {
						v40 = int32(0)
					}
				}
			}
			if v40 != 0 {
				v41 = v25
			} else {
				v41 = v26
			}
			v43 = *(*int32)(unsafe.Add(mBase, _consts[88]))
			if v43 == int32(0) {
				v49 = F_vsetRemoveEntry(m, v41, int32(1087), l1)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					if v49 == int32(0) {
						F__serverAssert(m, int32(_a1659), int32(_a1652), int32(126))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v53 = F_vsetIsEmpty(m, v41)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							if v53 == int32(0) {
								return
							} else {
								v57 = F_objectGetVal(m, l0)
								mBase = m.M
								v59 = v57 + int32(44)
								if v59 == int32(0) {
									v73 = int32(0)
								} else {
									v63 = int32(1)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
									switch v64 + v63 {
									case 0:
										v73 = v63
									case 1:
										v73 = int32(0)
									default:
										if v64&int32(7) != 0 {
											v73 = v63
										} else {
											v73 = int32(0)
										}
									}
								}
								if v73 == int32(0) {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									if v78&int32(240) != int32(32) {
									} else {
										v83 = F_objectGetVal(m, l0)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(_a1654)
									}
									return
								} else {
									F_vsetRelease(m, v59)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										if v78&int32(240) != int32(32) {
										} else {
											v83 = F_objectGetVal(m, l0)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(_a1654)
										}
										return
									}
								}
							}
						}
					}
				}
			} else {
				if v41 == int32(0) {
					F__serverAssert(m, int32(_a188), int32(_a1652), int32(125))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v49 = F_vsetRemoveEntry(m, v41, int32(1087), l1)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						if v49 == int32(0) {
							F__serverAssert(m, int32(_a1659), int32(_a1652), int32(126))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v53 = F_vsetIsEmpty(m, v41)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								if v53 == int32(0) {
									return
								} else {
									v57 = F_objectGetVal(m, l0)
									mBase = m.M
									v59 = v57 + int32(44)
									if v59 == int32(0) {
										v73 = int32(0)
									} else {
										v63 = int32(1)
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
										switch v64 + v63 {
										case 0:
											v73 = v63
										case 1:
											v73 = int32(0)
										default:
											if v64&int32(7) != 0 {
												v73 = v63
											} else {
												v73 = int32(0)
											}
										}
									}
									if v73 == int32(0) {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										if v78&int32(240) != int32(32) {
										} else {
											v83 = F_objectGetVal(m, l0)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(_a1654)
										}
										return
									} else {
										F_vsetRelease(m, v59)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											if v78&int32(240) != int32(32) {
											} else {
												v83 = F_objectGetVal(m, l0)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(_a1654)
											}
											return
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_hashTypeUpdateAsStringRef(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(9223372036854775807)
	v22 = F_hashTypeGetValue(m, l0, l1, v10+int32(12), v10+int32(8), v10, int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		if v22 != 0 {
			v92 = v12
			m.G0 = v10 + int32(16)
			return v92
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v26&int32(240) != int32(176) {
				v34 = F_objectGetVal(m, l0)
				mBase = m.M
				v35 = F_hashtableFindRef(m, v34, l1)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
					v41 = int64(-1)
					v43 = v37 + int32(-1)
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
					if v44&int32(7) == int32(0) {
						v79 = v41
					} else {
						if v44&int32(8) == int32(0) {
							v79 = v41
						} else {
							v53 = F_sdsAllocPtr(m, v37)
							mBase = m.M
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
							if int32(base.Ui32(v56&int32(16))>>(uint(int32(4))%32)) != 0 {
								v61 = int32(-4)
							} else {
								v61 = int32(0)
							}
							v64 = v56 & int32(7)
							if v64 != 0 {
								v65 = v61
							} else {
								v65 = int32(0)
							}
							if int32(base.Ui32(v56&int32(8))>>(uint(int32(3))%32)) != 0 {
								v73 = int32(-8)
							} else {
								v73 = int32(0)
							}
							if v64 != 0 {
								v75 = v73
							} else {
								v75 = int32(0)
							}
							v77 = *(*int64)(unsafe.Add(mBase, uint32(v53+v65+v75)))
							v79 = v77
						}
					}
					v81 = F_entryUpdateAsStringRef(m, v37, l2, l3, v79)
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						v83 = F_hashtableReplaceReallocatedEntry(m, v34, v37, v81)
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							if v83 == int32(0) {
								F__serverAssert(m, int32(_a1658), int32(_a1652), int32(345))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								F_hashTypeTrackUpdateEntry(m, l0, v37, v81, v79, v79)
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									v92 = int32(0)
									m.G0 = v10 + int32(16)
									return v92
								}
							}
						}
					}
				}
			} else {
				F_hashTypeConvert(m, l0, int32(2))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = F_objectGetVal(m, l0)
					mBase = m.M
					v35 = F_hashtableFindRef(m, v34, l1)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
						v41 = int64(-1)
						v43 = v37 + int32(-1)
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
						if v44&int32(7) == int32(0) {
							v79 = v41
						} else {
							if v44&int32(8) == int32(0) {
								v79 = v41
							} else {
								v53 = F_sdsAllocPtr(m, v37)
								mBase = m.M
								v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
								if int32(base.Ui32(v56&int32(16))>>(uint(int32(4))%32)) != 0 {
									v61 = int32(-4)
								} else {
									v61 = int32(0)
								}
								v64 = v56 & int32(7)
								if v64 != 0 {
									v65 = v61
								} else {
									v65 = int32(0)
								}
								if int32(base.Ui32(v56&int32(8))>>(uint(int32(3))%32)) != 0 {
									v73 = int32(-8)
								} else {
									v73 = int32(0)
								}
								if v64 != 0 {
									v75 = v73
								} else {
									v75 = int32(0)
								}
								v77 = *(*int64)(unsafe.Add(mBase, uint32(v53+v65+v75)))
								v79 = v77
							}
						}
						v81 = F_entryUpdateAsStringRef(m, v37, l2, l3, v79)
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							v83 = F_hashtableReplaceReallocatedEntry(m, v34, v37, v81)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v83 == int32(0) {
									F__serverAssert(m, int32(_a1658), int32(_a1652), int32(345))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									F_hashTypeTrackUpdateEntry(m, l0, v37, v81, v79, v79)
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int32(0)
									} else {
										v92 = int32(0)
										m.G0 = v10 + int32(16)
										return v92
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_rewriteHashObject(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v31 int64
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v280 int64
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v301 int64
	_ = v301
	var v302 int64
	_ = v302
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v329 int64
	_ = v329
	var v332 int64
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v374 int64
	_ = v374
	var v377 int64
	_ = v377
	var v383 int32
	_ = v383
	var v391 int32
	_ = v391
	v11 = m.G0
	v13 = v11 - int32(688)
	m.G0 = v13
	v16 = F_hashTypeHasVolatileFields(m, l2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(688)
	return v391
L2:
	;
	v286 = F_hashTypeLength(m, l2)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L3
	} else {
		goto L88
	}
L3:
	;
	return int32(0)
L4:
	;
	if v16 == int32(0) {
		v280 = int64(0)
		goto L2
	} else {
		goto L5
	}
L5:
	;
	F_hashTypeInitVolatileIterator(m, l2, v13+int32(8))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v31 = int64(0)
	goto L8
L7:
	;
	F_hashTypeResetIterator(m, v13+int32(8))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L3
	} else {
		goto L87
	}
L8:
	;
	v39 = F_hashTypeNext(m, v13+int32(8))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v39 == int32(-1) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v13)+680))
	v47 = int64(-1)
	v49 = v43 + int32(-1)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v50&int32(7) == int32(0) {
		v85 = v47
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)+680))
	goto L28
L13:
	;
	goto L12
L14:
	;
	if v50&int32(8) == int32(0) {
		v85 = v47
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v59 = F_sdsAllocPtr(m, v43)
	mBase = m.M
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if int32(base.Ui32(v62&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v67 = int32(-4)
	goto L18
L17:
	;
	v67 = int32(0)
	goto L18
L18:
	;
	v70 = v62 & int32(7)
	if v70 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v71 = v67
	goto L21
L20:
	;
	v71 = int32(0)
	goto L21
L21:
	;
	if int32(base.Ui32(v62&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v79 = int32(-8)
	goto L24
L23:
	;
	v79 = int32(0)
	goto L24
L24:
	;
	if v70 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v81 = v79
	goto L27
L26:
	;
	v81 = int32(0)
	goto L27
L27:
	;
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v59+v71+v81)))
	v85 = v83
	goto L13
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v13)+680))
	v90 = v13 + int32(4)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+int32(-1)))))
	v98 = v96 & int32(7)
	if v98 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v203 = F_rioWriteBulkCount(m, l0, int32(42), int32(8))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L3
	} else {
		goto L63
	}
L30:
	;
	v200 = v192
	goto L29
L31:
	;
	v153 = F_sdsAllocPtr(m, v88)
	mBase = m.M
	v155 = v153 + int32(-4)
	if v96&int32(32) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L32:
	;
	switch v98 {
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
		v118 = int32(0)
		goto L35
	}
L33:
	;
	if v96&int32(16) != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v120 = int32(1)
	v121 = F_sdsHdrSize(m, v120)
	mBase = m.M
	v122 = v88 + v118 + v121
	v124 = v122 + v120
	if v90 == int32(0) {
		v192 = v124
		goto L30
	} else {
		goto L41
	}
L36:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v88+int32(-17))))
	v118 = v117
	goto L35
L37:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v88+int32(-9))))
	v118 = v114
	goto L35
L38:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88+int32(-5)))))
	v118 = v111
	goto L35
L39:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+int32(-3)))))
	v118 = v108
	goto L35
L40:
	;
	v118 = int32(base.Ui32(v96) >> (uint(int32(3)) % 32))
	goto L35
L41:
	;
	v127 = int32(0)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v127))))
	switch v130 & int32(7) {
	case 0:
		goto L47
	case 1:
		goto L46
	case 2:
		goto L45
	case 3:
		goto L44
	case 4:
		goto L43
	default:
		v151 = v127
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v151
	v200 = v124
	goto L29
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v122+int32(-16))))
	v151 = v150
	goto L42
L44:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v122+int32(-8))))
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v146
	v200 = v124
	goto L29
L45:
	;
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122+int32(-4)))))
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v142
	v200 = v124
	goto L29
L46:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+int32(-2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v138
	v200 = v124
	goto L29
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(base.Ui32(v130) >> (uint(int32(3)) % 32))
	v200 = v124
	goto L29
L48:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	if v90 == int32(0) {
		v192 = v167
		goto L30
	} else {
		goto L54
	}
L49:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	if v160 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v90 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v200 = int32(0)
	goto L29
L52:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v200 = v166
	goto L29
L53:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v164
	goto L52
L54:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+int32(-1)))))
	switch v173 & int32(7) {
	case 0:
		goto L60
	case 1:
		goto L59
	case 2:
		goto L58
	case 3:
		goto L57
	case 4:
		goto L56
	default:
		v190 = int32(0)
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v190
	v192 = v167
	goto L30
L56:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v167+int32(-17))))
	v190 = v189
	goto L55
L57:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v167+int32(-9))))
	v190 = v186
	goto L55
L58:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167+int32(-5)))))
	v190 = v183
	goto L55
L59:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+int32(-3)))))
	v190 = v180
	goto L55
L60:
	;
	v190 = int32(base.Ui32(v173) >> (uint(int32(3)) % 32))
	goto L55
L61:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v266 = F_rioWriteBulkString(m, l0, v200, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L3
	} else {
		goto L85
	}
L62:
	;
	v391 = int32(0)
	goto L1
L63:
	;
	if v203 == int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v209 = F_rioWriteBulkString(m, l0, int32(_a172), int32(6))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	if v209 == int32(0) {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v213 = F_rioWriteBulkObject(m, l0, l1)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	if v213 == int32(0) {
		goto L62
	} else {
		goto L68
	}
L68:
	;
	v219 = F_rioWriteBulkString(m, l0, int32(_a173), int32(4))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	if v219 == int32(0) {
		goto L62
	} else {
		goto L70
	}
L70:
	;
	v223 = F_rioWriteBulkLongLong(m, l0, v85)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	if v223 == int32(0) {
		goto L62
	} else {
		goto L72
	}
L72:
	;
	v229 = F_rioWriteBulkString(m, l0, int32(_a174), int32(6))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	if v229 == int32(0) {
		goto L62
	} else {
		goto L74
	}
L74:
	;
	v234 = F_rioWriteBulkLongLong(m, l0, int64(1))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	if v234 == int32(0) {
		goto L62
	} else {
		goto L76
	}
L76:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+int32(-1)))))
	switch v241 & int32(7) {
	case 0:
		goto L82
	case 1:
		goto L81
	case 2:
		goto L80
	case 3:
		goto L79
	case 4:
		goto L78
	default:
		v258 = int32(0)
		goto L77
	}
L77:
	;
	v259 = F_rioWriteBulkString(m, l0, v87, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L3
	} else {
		goto L83
	}
L78:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v87+int32(-17))))
	v258 = v257
	goto L77
L79:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v87+int32(-9))))
	v258 = v254
	goto L77
L80:
	;
	v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87+int32(-5)))))
	v258 = v251
	goto L77
L81:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+int32(-3)))))
	v258 = v248
	goto L77
L82:
	;
	v258 = int32(base.Ui32(v241) >> (uint(int32(3)) % 32))
	goto L77
L83:
	;
	if v259 != 0 {
		goto L61
	} else {
		goto L84
	}
L84:
	;
	goto L62
L85:
	;
	if v266 != 0 {
		v31 = v31 + base.I64_extend_i32_u(base.B2i32(v266 != int32(0)))
		goto L8
	} else {
		goto L86
	}
L86:
	;
	v391 = int32(0)
	goto L1
L87:
	;
	v280 = v31
	goto L2
L88:
	;
	F_hashTypeInitIterator(m, l2, v13+int32(8))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L3
	} else {
		goto L89
	}
L89:
	;
	v301 = int64(0)
	v302 = base.I64_extend_i32_u(v286) - v280
	goto L91
L90:
	;
	F_hashTypeResetIterator(m, v13+int32(8))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L3
	} else {
		goto L122
	}
L91:
	;
	v309 = F_hashTypeNext(m, v13+int32(8))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L3
	} else {
		goto L93
	}
L93:
	;
	if v309 == int32(-1) {
		goto L90
	} else {
		goto L94
	}
L94:
	;
	if v280 < int64(1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if v301 != int64(0) {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v13)+680))
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313+int32(-1)))))
	goto L97
L97:
	;
	if base.B2i32(v316&int32(7) != int32(0))&int32(base.Ui32(v316&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L91
	} else {
		goto L98
	}
L98:
	;
	goto L95
L99:
	;
	v358 = F_rioWriteHashIteratorCursor(m, l0, v13+int32(8), int32(1))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L3
	} else {
		goto L114
	}
L100:
	;
	v329 = int64(64)
	if v302 < v329 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	F_hashTypeResetIterator(m, v13+int32(8))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L3
	} else {
		goto L111
	}
L102:
	;
	v332 = v302
	goto L104
L103:
	;
	v332 = v329
	goto L104
L104:
	;
	v338 = F_rioWriteBulkCount(m, l0, int32(42), base.I32_wrap_i64(v332)<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L3
	} else {
		goto L105
	}
L105:
	;
	if v338 == int32(0) {
		goto L101
	} else {
		goto L106
	}
L106:
	;
	v344 = F_rioWriteBulkString(m, l0, int32(_a175), int32(5))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L3
	} else {
		goto L107
	}
L107:
	;
	if v344 == int32(0) {
		goto L101
	} else {
		goto L108
	}
L108:
	;
	v348 = F_rioWriteBulkObject(m, l0, l1)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L3
	} else {
		goto L109
	}
L109:
	;
	if v348 != 0 {
		goto L99
	} else {
		goto L110
	}
L110:
	;
	goto L101
L111:
	;
	v391 = int32(0)
	goto L1
L112:
	;
	v374 = v301 + int64(1)
	if v374 == int64(64) {
		goto L119
	} else {
		goto L120
	}
L113:
	;
	F_hashTypeResetIterator(m, v13+int32(8))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L3
	} else {
		goto L118
	}
L114:
	;
	if v358 == int32(0) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v365 = F_rioWriteHashIteratorCursor(m, l0, v13+int32(8), int32(2))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	if v365 != 0 {
		goto L112
	} else {
		goto L117
	}
L117:
	;
	goto L113
L118:
	;
	v391 = int32(0)
	goto L1
L119:
	;
	v377 = int64(0)
	goto L121
L120:
	;
	v377 = v374
	goto L121
L121:
	;
	v301 = v377
	v302 = v302 + int64(-1)
	goto L91
L122:
	;
	v391 = int32(1)
	goto L1
}
