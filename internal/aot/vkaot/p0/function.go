package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_functionDeleteCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v4 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v8 = F_objectGetVal(m, v7)
	mBase = m.M
	v9 = F_dictFetchValue(m, v5, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v9 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, _consts[18]))
			F_libraryUnlink(m, v15, v9)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				F_dictRelease(m, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					F_sdsfree(m, v21)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						F_sdsfree(m, v24)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							F_valkey_free(m, v9)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								v29 = int32(_a69)
								v31 = *(*int64)(unsafe.Add(mBase, _consts[60]))
								*(*int64)(unsafe.Add(mBase, _consts[60])) = v31 + int64(1)
								v36 = *(*int32)(unsafe.Add(mBase, _consts[77]))
								F_addReply(m, l0, v36)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		} else {
			F_addReplyError(m, l0, int32(_a580))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_functionFreeLibMetaData(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 == int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8 == int32(0) {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v13 == int32(0) {
				return
			} else {
				F_sdsfree(m, v13)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_sdsfree(m, v8)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v13 == int32(0) {
					return
				} else {
					F_sdsfree(m, v13)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		F_sdsfree(m, v3)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v8 == int32(0) {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v13 == int32(0) {
					return
				} else {
					F_sdsfree(m, v13)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				F_sdsfree(m, v8)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v13 == int32(0) {
						return
					} else {
						F_sdsfree(m, v13)
						mBase = m.M
						v17 = m.ExcPending
						if v17 != 0 {
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
func F_functionLoadCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(4) <= v10 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F__serverAssert(m, int32(_a583), int32(_a582), int32(1140))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L23
	} else {
		goto L44
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return
L3:
	;
	if v82 < v80 {
		goto L25
	} else {
		goto L26
	}
L4:
	;
	v20 = int32(2)
	goto L7
L5:
	;
	v80 = v10
	v81 = int32(0)
	v82 = int32(2)
	goto L3
L6:
	;
	v73 = F_objectGetVal(m, v25)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v73
	F_addReplyErrorFormat(m, l0, int32(_a584), v8)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v20<<(uint(int32(2))%32))))
	v26 = F_objectGetVal(m, v25)
	mBase = m.M
	v27 = int32(_a138)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v62-v64 != 0 {
		goto L6
	} else {
		goto L21
	}
L10:
	;
	v62 = F_tolower(m, v58)
	mBase = m.M
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v64 = F_tolower(m, v63)
	mBase = m.M
	goto L9
L11:
	;
	v32 = v26
	v33 = v27
	v34 = v30
	goto L14
L12:
	;
	v58 = int32(0)
	v59 = v27
	goto L10
L13:
	;
	v58 = v55 & int32(255)
	v59 = v54
	goto L10
L14:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v36 == int32(0) {
		v54 = v33
		v55 = v34
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v54 = v48
	v55 = int32(0)
	goto L13
L16:
	;
	v40 = v34 & int32(255)
	if v40 == v36 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v47 = int32(1)
	v48 = v33 + v47
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	if v49 != 0 {
		v32 = v32 + v47
		v33 = v48
		v34 = v49
		goto L14
	} else {
		goto L20
	}
L18:
	;
	v42 = F_tolower(m, v40)
	mBase = m.M
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v44 = F_tolower(m, v43)
	mBase = m.M
	if v42 == v44 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v54 = v33
	v55 = v46
	goto L13
L20:
	;
	goto L15
L21:
	;
	v66 = int32(1)
	v68 = v20 + v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v69+int32(-1) <= v68 {
		v80 = v69
		v81 = v66
		v82 = v68
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v20 = v68
	goto L7
L23:
	;
	return
L24:
	;
	goto L2
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v82<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
	v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v95 != int64(-1) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	F_addReplyError(m, l0, int32(_a585))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L2
L28:
	;
	v111 = F_objectGetVal(m, v91)
	mBase = m.M
	v114 = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	if v110 != 0 {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v99 = int32(1)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v100&v99 != 0 {
		v107 = v99
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v110 = int32(1)
	goto L28
L31:
	;
	v110 = v107
	goto L28
L32:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v103 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v105 = F_isImportSlotMigrationJob(m, v103)
	mBase = m.M
	v107 = v105
	goto L31
L34:
	;
	v110 = int32(0)
	goto L28
L35:
	;
	v126 = int32(_a69)
	v128 = *(*int64)(unsafe.Add(mBase, _consts[60]))
	*(*int64)(unsafe.Add(mBase, _consts[60])) = v128 + int64(1)
	F_addReplyBulkSds(m, l0, v119)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L23
	} else {
		goto L43
	}
L36:
	;
	v118 = v114
	goto L38
L37:
	;
	v118 = int32(500)
	goto L38
L38:
	;
	v119 = F_functionsCreateWithLibraryCtx(m, v111, v81, v8+int32(12), v115, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L23
	} else {
		goto L39
	}
L39:
	;
	if v119 != 0 {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v121 == int32(0) {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_addReplyErrorSds(m, l0, v121)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L23
	} else {
		goto L42
	}
L42:
	;
	goto L2
L43:
	;
	goto L2
L44:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
