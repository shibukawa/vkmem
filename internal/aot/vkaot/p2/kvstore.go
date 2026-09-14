package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_kvstoreCreate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	if int32(17) <= l1 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F__serverAssert(m, int32(_a863), int32(_a864), int32(296))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L11
	} else {
		goto L34
	}
L2:
	;
	F__serverAssert(m, int32(_a865), int32(_a864), int32(295))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L11
	} else {
		goto L33
	}
L3:
	;
	F__serverAssert(m, int32(_a866), int32(_a864), int32(294))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L11
	} else {
		goto L32
	}
L4:
	;
	F__serverAssert(m, int32(_a867), int32(_a864), int32(293))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L11
	} else {
		goto L31
	}
L5:
	;
	F__serverAssert(m, int32(_a868), int32(_a864), int32(289))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L11
	} else {
		goto L30
	}
L6:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 != int32(543) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v11 != int32(544) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v14 != int32(545) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v17 != int32(546) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v21 = F_valkey_calloc(m, int32(80))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = l0
	v29 = int32(1) << (uint(l1) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v29
	v33 = F_valkey_calloc(m, int32(4)<<(uint(l1)%32))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v33
	v37 = F_hashtableCreate(m, int32(_a869))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v37
	v40 = F_listCreate(m)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v40
	if l1 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v50
	if l2&int32(1) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v44 = int32(8)
	v48 = F_valkey_calloc(m, v44<<(uint(l1)%32)+v44)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L11
	} else {
		goto L19
	}
L18:
	;
	v50 = int32(0)
	goto L16
L19:
	;
	v50 = v48
	goto L16
L20:
	;
	return v21
L21:
	;
	v56 = int32(0)
	v59 = v29
	goto L22
L22:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v62 = v56 << (uint(int32(2)) % 32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v62)))
	if v64 != 0 {
		v87 = v59
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L20
L24:
	;
	v89 = v56 + int32(1)
	if v89 < v87 {
		v56 = v89
		v59 = v87
		goto L22
	} else {
		goto L29
	}
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v66 = F_hashtableCreate(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v68+v62))) = v66
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66+int32(44))+4)) = v21
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74+v62)))
	v77 = F_hashtableMemUsage(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v77 + v79
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v82 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v87 = v86
	goto L24
L29:
	;
	goto L23
L30:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
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
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_kvstoreEmpty(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 < int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	F_hashtableEmpty(m, v79, int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L10
	} else {
		goto L23
	}
L2:
	;
	v13 = int32(0)
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = v13 << (uint(int32(2)) % 32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v19)))
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v70 = v13 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v70 < v71 {
		v13 = v70
		goto L3
	} else {
		goto L22
	}
L6:
	;
	v25 = v21 + int32(44)
	goto L8
L7:
	;
	F_hashtableEmpty(m, v21, l1)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v26 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(0)
	goto L7
L10:
	;
	return
L11:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v33&int32(2) == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38+v19)))
	if v40 == int32(0) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	goto L14
L14:
	;
	if v43+v44 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46+v19)))
	if v48 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	F_hashtableRelease(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L21
	}
L17:
	;
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+26)))
	goto L19
L18:
	;
	v56 = int32(0)
	goto L16
L19:
	;
	if int32(0) < v50 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53+v19)))
	v56 = v55
	goto L16
L21:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v59+v19))) = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v63 + int32(-1)
	goto L5
L22:
	;
	goto L4
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = int64(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_listEmpty(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v88 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v88
	v90 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v90
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(48)))) = v88
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v98 == v90 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
	return
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v108 = F__emscripten_memset_bulkmem(m, v98, base.I32_extend8_s(int32(0)), v102<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	goto L27
L27:
	;
	goto L25
}
func F_kvstoreGetHashtableIterator(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v34 int32
	_ = v34
	v3 = l2
	v6 = F_valkey_malloc(m, int32(64))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = base.I64_extend_i32_s(l1)
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
		v14 = v6 + int32(16)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v15+l1<<(uint(int32(2))%32))))
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)) = uint8(v20)
		*(*int32)(unsafe.Add(mBase, uint32(v14))) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v20
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)) = uint8(v3)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = int32(-1)
		if v19 == v20 {
		} else {
			if v3&int32(1) == int32(0) {
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v34
				*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v14
			}
		}
		return v6
	}
}
func F_kvstoreGetNextNonEmptyHashtableIndex(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 != int32(1) {
		if int32(0) <= l1 {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v30 = l1 + int32(1)
			v32 = int64(0)
			for {
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v27+v30<<(uint(int32(3))%32))))
				v40 = v39 + v32
				v43 = (v30 + int32(-1)) & v30
				if v43 != 0 {
					v30 = v43
					v32 = v40
					continue
				} else {
					break
				}
				break
			}
			v49 = v40 + int64(1)
		} else {
			v49 = int64(1)
		}
		v53 = int32(-1)
		v54 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		if base.Ui64(v54) < base.Ui64(v49) {
			v95 = v53
			return v95
		} else {
			v56 = base.I32_wrap_i64(v49)
			if v56 == int32(0) {
				v95 = v53
				return v95
			} else {
				v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				if v59 == int64(0) {
					v95 = v53
					return v95
				} else {
					if base.Ui64(v59) < base.Ui64(v49&int64(4294967295)) {
						F__serverAssert(m, int32(_a874), int32(_a864), int32(576))
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
					} else {
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						v71 = int32(1) << (uint(v66) % 32)
						v72 = int32(0)
						v74 = v56
						for {
							v77 = v72 + v71
							v81 = *(*int64)(unsafe.Add(mBase, uint32(v68+v77<<(uint(int32(3))%32))))
							v83 = base.B2i32(base.Ui64(v81) < base.Ui64(base.I64_extend_i32_u(v74)))
							if base.Ui64(v81) < base.Ui64(base.I64_extend_i32_u(v74)) {
								v84 = v77
							} else {
								v84 = v72
							}
							if base.Ui64(v81) < base.Ui64(base.I64_extend_i32_u(v74)) {
								v87 = base.I32_wrap_i64(v81)
							} else {
								v87 = int32(0)
							}
							v89 = int32(1)
							if base.Ui32(v89) < base.Ui32(v71) {
								v71 = v71 >> (uint(v89) % 32)
								v72 = v84
								v74 = v74 - v87
								continue
							} else {
								break
							}
							break
						}
						v95 = v84
						return v95
					}
				}
			}
		}
	} else {
		if l1 == int32(0) {
			v95 = int32(-1)
			return v95
		} else {
			F__serverAssert(m, int32(_a875), int32(_a864), int32(606))
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
		}
	}
}
func F_kvstoreGetStats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v69 int64
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	v5 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v5)
	v16 = F_valkey_malloc(m, int32(80))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v22 == int32(1) {
		v69 = int64(0)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v74 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v74
	v76 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+72)) = uint8(v76)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v69
	v80 = v16 + int32(24)
	*(*uint8)(unsafe.Add(mBase, uint32(v80)+14)) = uint8(v74)
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = v74
	*(*uint8)(unsafe.Add(mBase, uint32(v80)+15)) = uint8(v76)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = int32(-1)
	goto L16
L4:
	;
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if v26 == int64(0) {
		v69 = int64(-1)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v29 = int32(1)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v35 = int32(0)
	v42 = v29 << (uint(v30) % 32)
	v44 = v29
	goto L6
L6:
	;
	v46 = v35 + v42
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v32+v46<<(uint(int32(3))%32))))
	v52 = base.B2i32(base.Ui64(v50) < base.Ui64(base.I64_extend_i32_u(v44)))
	if base.Ui64(v50) < base.Ui64(base.I64_extend_i32_u(v44)) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v69 = base.I64_extend_i32_s(v53)
	goto L3
L8:
	;
	v53 = v46
	goto L10
L9:
	;
	v53 = v35
	goto L10
L10:
	;
	if base.Ui64(v50) < base.Ui64(base.I64_extend_i32_u(v44)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v56 = base.I32_wrap_i64(v50)
	goto L13
L12:
	;
	v56 = int32(0)
	goto L13
L13:
	;
	v58 = int32(1)
	if base.Ui32(v58) < base.Ui32(v42) {
		v35 = v53
		v42 = v42 >> (uint(v58) % 32)
		v44 = v44 - v56
		goto L6
	} else {
		goto L14
	}
L14:
	;
	goto L7
L15:
	;
	v100 = F_kvstoreIteratorNextHashtable(m, v16)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L21
	}
L16:
	;
	goto L15
L19:
	;
	F_kvstoreIteratorRelease(m, v16)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L53
	}
L20:
	;
	v108 = v5
	v111 = v100
	v113 = int32(0)
	goto L23
L21:
	;
	if v100 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v243 = v5
	v248 = int32(0)
	goto L19
L23:
	;
	v116 = F_hashtableGetStatsHt(m, v111, int32(0), l3)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v243 = v235
	v248 = v172
	goto L19
L25:
	;
	if v113 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	goto L38
L27:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+8)) = v121 + v122
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = v125 + v126
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v116)+24))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v113)+24))
	if base.Ui32(v130) < base.Ui32(v129) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v172 = v116
	goto L26
L29:
	;
	F_hashtableFreeStats(m, v116)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L36
	}
L30:
	;
	v132 = v129
	goto L32
L31:
	;
	v132 = v130
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+24)) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v113)+16))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+16)) = v134 + v135
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v113)+20))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+20)) = v138 + v139
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v113)+28))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v116)+28))
	v146 = int32(0)
	goto L33
L33:
	;
	v150 = int32(2)
	v151 = v146 << (uint(v150) % 32)
	v152 = v142 + v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v143+v151)))
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v153 + v155
	v159 = v151 | int32(4)
	v160 = v142 + v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v143+v159)))
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v161 + v163
	v167 = v146 + v150
	if v167 != int32(50) {
		v146 = v167
		goto L33
	} else {
		goto L35
	}
L34:
	;
	goto L29
L35:
	;
	goto L34
L36:
	;
	v172 = v113
	goto L26
L37:
	;
	v237 = F_kvstoreIteratorNextHashtable(m, v16)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L51
	}
L38:
	;
	if base.B2i32(v173 != int32(-1)) == int32(0) {
		v235 = v108
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v179 = F_hashtableGetStatsHt(m, v111, int32(1), l3)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v108 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v108)+8)) = v184 + v185
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v108)+12)) = v188 + v189
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v179)+24))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
	if base.Ui32(v193) < base.Ui32(v192) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v235 = v179
	goto L37
L43:
	;
	F_hashtableFreeStats(m, v179)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L50
	}
L44:
	;
	v195 = v192
	goto L46
L45:
	;
	v195 = v193
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+24)) = v195
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = v197 + v198
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v108)+20)) = v201 + v202
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v108)+28))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v179)+28))
	v209 = int32(0)
	goto L47
L47:
	;
	v213 = int32(2)
	v214 = v209 << (uint(v213) % 32)
	v215 = v205 + v214
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v206+v214)))
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v216 + v218
	v222 = v214 | int32(4)
	v223 = v205 + v222
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v206+v222)))
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v224 + v226
	v230 = v209 + v213
	if v230 != int32(50) {
		v209 = v230
		goto L47
	} else {
		goto L49
	}
L48:
	;
	goto L43
L49:
	;
	goto L48
L50:
	;
	v235 = v108
	goto L37
L51:
	;
	if v237 != 0 {
		v108 = v235
		v111 = v237
		v113 = v172
		goto L23
	} else {
		goto L52
	}
L52:
	;
	goto L24
L53:
	;
	if l2 == int32(0) {
		v262 = l1
		v263 = l2
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if v243 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	if v248 == int32(0) {
		v262 = l1
		v263 = l2
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v256 = F_hashtableGetStatsMsg(m, l1, l2, v248, l3)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_hashtableFreeStats(m, v248)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v262 = l1 + v256
	v263 = l2 - v256
	goto L54
L59:
	;
	if l2 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	if v263 == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v268 = F_hashtableGetStatsMsg(m, v262, v263, v243, l3)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_hashtableFreeStats(m, v243)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	return
L65:
	;
	v277 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1+l2+int32(-1)))) = uint8(v277)
	goto L64
}
func F_kvstoreHashtableFindRef(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4+l1<<(uint(int32(2))%32))))
	if v8 != 0 {
		v11 = F_hashtableFindRef(m, v8, l2)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v11
		}
	} else {
		return int32(0)
	}
}
func F_kvstoreHashtableInsertAtPosition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5+l1<<(uint(int32(2))%32))))
	F_hashtableInsertAtPosition(m, v9, l2, l3)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		F_cumulativeKeyCountAdd(m, l0, l1, int32(1))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			return
		}
	}
}
func F_kvstoreHashtableIteratorNext(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4+v5<<(uint(int32(2))%32))))
	if v9 != 0 {
		v14 = F_hashtableNext(m, l0+int32(16), l1)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			return v14
		}
	} else {
		return int32(0)
	}
}
func F_kvstoreHashtableRehashingCount(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+20))
	return v3
}
func F_kvstoreHashtableRehashingStarted(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = l0 + int32(44)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v13 = F_listAddNodeTail(m, v12, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v16
		F_hashtableRehashingInfo(m, l0, v7+int32(12), v7+int32(8))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
			v25 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+8)))
			*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v24 + v25
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v28 + v29<<(uint(int32(6))%32)
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_kvstoreHashtableSampleEntries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5+l1<<(uint(int32(2))%32))))
	if v9 != 0 {
		v12 = F_hashtableSampleEntries(m, v9, l2, l3)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v12
		}
	} else {
		return int32(0)
	}
}
func F_kvstoreHashtableSize(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v3+l1<<(uint(int32(2))%32))))
	if v7 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
		return v10 + v11
	} else {
		return int32(0)
	}
}
func F_kvstoreHashtableTrackMemUsage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(44))+4))
	if v5 == int32(0) {
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+60))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+60)) = v8 + l1
	}
	return
}
func F_kvstoreIncrementallyRehash(m *base.Module, l0 int32, l1 int64) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v13 = m.T0[v12].(func(*base.Module) int64)(m)
	mBase = m.M
	v16 = int64(0)
	goto L4
L2:
	;
	return int64(0)
L3:
	;
	return v34
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v20 == int32(0) {
		v34 = v16
		goto L3
	} else {
		goto L6
	}
L5:
	;
	v34 = v32
	goto L3
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v25 = F_hashtableRehashMicroseconds(m, v23, l1-v16)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int64(0)
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v31 = m.T0[v30].(func(*base.Module) int64)(m)
	mBase = m.M
	v32 = v31 - v13
	if base.Ui64(v32) < base.Ui64(l1) {
		v16 = v32
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
func F_kvstoreOverheadHashtableLut(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	return v2
}
func F_kvstoreRelease(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v5 < int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v10 = v5
	v11 = int32(0)
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13+v11<<(uint(int32(2))%32))))
	if v17 == int32(0) {
		v30 = v10
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v32 = v11 + int32(1)
	if v32 < v30 {
		v10 = v30
		v11 = v32
		goto L3
	} else {
		goto L12
	}
L6:
	;
	v21 = v17 + int32(44)
	goto L8
L7:
	;
	F_hashtableRelease(m, v17)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v22 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	goto L7
L10:
	;
	return
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = v29
	goto L5
L12:
	;
	goto L4
L13:
	;
	F__serverAssert(m, int32(_a870), int32(_a864), int32(346))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L22
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_valkey_free(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	F_hashtableRelease(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_listRelease(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v48 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L21
	}
L19:
	;
	F_valkey_free(m, v48)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	return
L22:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_kvstoreScan(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v127 int64
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int64
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int64
	_ = v148
	var v152 int64
	_ = v152
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v17 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if int32(0) <= l2 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v21 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+16)))
	v27 = int64(base.Ui64(l1) >> (uint(v21) % 64))
	v28 = (v17 + int32(-1)) & base.I32_wrap_i64(l1)
	goto L1
L3:
	;
	v27 = l1
	v28 = int32(0)
	goto L1
L4:
	;
	F__serverAssert(m, int32(_a873), int32(_a864), int32(150))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L20
	} else {
		goto L51
	}
L5:
	;
	F__serverAssert(m, int32(_a872), int32(_a864), int32(430))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L20
	} else {
		goto L50
	}
L6:
	;
	F__serverAssert(m, int32(_a871), int32(_a864), int32(429))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L20
	} else {
		goto L49
	}
L7:
	;
	m.G0 = v15 + int32(16)
	return v152
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v38<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l4
	if v44 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	if l3 < l2 {
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v38 = v28
	v39 = v27
	goto L8
L11:
	;
	if v17 <= l3 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	if v28 < l2 {
		v38 = l2
		v39 = int64(0)
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if base.Ui32(l3) < base.Ui32(v28) {
		v152 = int64(0)
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v38 = v28
	v39 = v27
	goto L8
L15:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v145 == int32(1) {
		v152 = v137
		goto L7
	} else {
		goto L48
	}
L16:
	;
	if l2 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L17:
	;
	if l5 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v38 < int32(0) {
		goto L4
	} else {
		goto L23
	}
L19:
	;
	v52 = m.T0[l5].(func(*base.Module, int32) int32)(m, v44)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int64(0)
L21:
	;
	if v52 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v58 <= v38 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v62 = F_hashtableFind(m, v60, v38, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	if v62 != 0 {
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v68 = F_hashtableScan(m, v44, base.I32_wrap_i64(v39), int32(547), v15+int32(4))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v70&int32(2) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v68 == int32(0) {
		goto L16
	} else {
		goto L38
	}
L29:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v38<<(uint(int32(2))%32))))
	if v79 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	goto L31
L31:
	;
	if v82+v83 != 0 {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v85 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v88 = v38 << (uint(int32(2)) % 32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v88)))
	if v90 == v85 {
		v101 = v85
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_hashtableRelease(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L20
	} else {
		goto L37
	}
L34:
	;
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+26)))
	goto L35
L35:
	;
	if int32(0) < v93 {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+v38<<(uint(int32(2))%32))))
	v101 = v100
	goto L33
L37:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v104+v88))) = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v108 + int32(-1)
	goto L28
L38:
	;
	v137 = base.I64_extend_i32_u(v68)
	v143 = v38
	goto L15
L39:
	;
	v127 = int64(0)
	v130 = F_kvstoreGetNextNonEmptyHashtableIndex(m, l0, v38)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L20
	} else {
		goto L42
	}
L40:
	;
	if l3 <= v38 {
		v152 = int64(0)
		goto L7
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v133 = base.B2i32(int32(-1) < l2) & base.B2i32(l3 < v130)
	if v133 != 0 {
		v152 = v127
		goto L7
	} else {
		goto L43
	}
L43:
	;
	if v133 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v134 = v38
	goto L46
L45:
	;
	v134 = v130
	goto L46
L46:
	;
	if v134 == int32(-1) {
		v152 = v127
		goto L7
	} else {
		goto L47
	}
L47:
	;
	v137 = v127
	v143 = v134
	goto L15
L48:
	;
	v148 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+16)))
	v152 = v137<<(uint(v148)%64) | base.I64_extend_i32_s(v143)
	goto L7
L49:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_kvstoreSetIsImporting(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v44 int64
	_ = v44
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	if l1 < int32(0) {
		F__serverAssert(m, int32(_a873), int32(_a864), int32(952))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7 <= l1 {
			F__serverAssert(m, int32(_a873), int32(_a864), int32(952))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v9+l1<<(uint(int32(2))%32))))
			if l2 == int32(0) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				v25 = F_hashtableDelete(m, v24, l1)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					if v25 == int32(0) {
						return
					} else {
						if v13 == int32(0) {
							return
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
							if v31+v32 == int32(0) {
								return
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
								F_cumulativeKeyCountAdd(m, l0, l1, v36+v37)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
									v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v44 - base.I64_extend_i32_u(v41+v42)
									return
								}
							}
						}
					}
				}
			} else {
				if v13 == int32(0) {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					v22 = F_hashtableAdd(m, v21, l1)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						return
					}
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
					if v18+v19 != 0 {
						F__serverAssert(m, int32(_a876), int32(_a864), int32(958))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
						v22 = F_hashtableAdd(m, v21, l1)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
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
