package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_dictAdd(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_dictAddRaw(m, l0, l1, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l2
			return int32(0)
		} else {
			return int32(1)
		}
	}
}
func F_dictAddOrFind(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_dictAddRaw(m, l0, l1, v6+int32(12))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		m.G0 = v6 + int32(16)
		if v10 != 0 {
			v18 = v10
		} else {
			v18 = v14
		}
		return v18
	}
}
func F_dictAddRaw(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	v7 = F_dictFindPositionForInsert(m, l0, l1, l2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v14 == int32(0) {
				v19 = l1
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v22 = base.B2i32(v20 != int32(-1))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+v22<<(uint(int32(2))%32))+4))
				if base.Ui32(v7) < base.Ui32(v26) {
					F__serverAssert(m, int32(_a596), int32(_a597), int32(450))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v29 = int32(-1)
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v22)+26)))
					if v31 == int32(255) {
						v37 = int32(0)
					} else {
						v37 = v29<<(uint(v31)%32) ^ v29
					}
					if base.Ui32(v26+v37<<(uint(int32(2))%32)) < base.Ui32(v7) {
						F__serverAssert(m, int32(_a596), int32(_a597), int32(450))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
						v44 = F_valkey_malloc(m, int32(24))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v42
							*(*int32)(unsafe.Add(mBase, uint32(v44))) = v19
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v44
							v51 = l0 + v22<<(uint(int32(2))%32)
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v52 + int32(1)
							return v44
						}
					}
				}
			} else {
				v17 = m.T0[v14].(func(*base.Module, int32) int32)(m, l1)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = v17
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v22 = base.B2i32(v20 != int32(-1))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+v22<<(uint(int32(2))%32))+4))
					if base.Ui32(v7) < base.Ui32(v26) {
						F__serverAssert(m, int32(_a596), int32(_a597), int32(450))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v29 = int32(-1)
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v22)+26)))
						if v31 == int32(255) {
							v37 = int32(0)
						} else {
							v37 = v29<<(uint(v31)%32) ^ v29
						}
						if base.Ui32(v26+v37<<(uint(int32(2))%32)) < base.Ui32(v7) {
							F__serverAssert(m, int32(_a596), int32(_a597), int32(450))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
							v44 = F_valkey_malloc(m, int32(24))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v42
								*(*int32)(unsafe.Add(mBase, uint32(v44))) = v19
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v44
								v51 = l0 + v22<<(uint(int32(2))%32)
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v52 + int32(1)
								return v44
							}
						}
					}
				}
			}
		} else {
			return int32(0)
		}
	}
}
func F_dictBucketRehash(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int64
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _consts[302]))
	if v10 == int32(2) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v13 == int32(-1) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v10 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = base.I32_wrap_i64(l1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42+v43<<(uint(int32(2))%32))))
	if v47 != 0 {
		goto L19
	} else {
		goto L20
	}
L6:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	if v20 == int32(255) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if base.Ui32(v31) <= base.Ui32(v24) {
		goto L5
	} else {
		goto L16
	}
L8:
	;
	v24 = int32(0)
	goto L10
L9:
	;
	v24 = int32(1) << (uint(v20) % 32)
	goto L10
L10:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v27 == int32(255) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = int32(0)
	goto L13
L12:
	;
	v31 = int32(1) << (uint(v27) % 32)
	goto L13
L13:
	;
	if base.Ui32(v24) <= base.Ui32(v31) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(v24) < base.Ui32(v31<<(uint(int32(2))%32)) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L7
L16:
	;
	if base.Ui32(v31) < base.Ui32(v24<<(uint(int32(5))%32)) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L5
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100+v43<<(uint(int32(2))%32)))) = int32(0)
	if v102 != 0 {
		goto L1
	} else {
		goto L31
	}
L19:
	;
	v51 = v47
	goto L21
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v100 = v42
	v102 = v48
	goto L18
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+27)))
	v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v57 <= v58 {
		v66 = v57
		v67 = l1
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v100 = v96
	v102 = v90
	goto L18
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v70 = int32(-1)
	v71 = int32(255)
	v72 = v66 & v71
	if v72 == v71 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = m.T0[v62].(func(*base.Module, int32) int64)(m, v60)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return
L26:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	v66 = v65
	v67 = v63
	goto L23
L27:
	;
	v79 = int64(0)
	goto L29
L28:
	;
	v79 = base.I64_extend_i32_u(v70<<(uint(v72)%32) ^ v70)
	goto L29
L29:
	;
	v84 = v68 + base.I32_wrap_i64(v79&v67)<<(uint(int32(2))%32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v51
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v90 = v88 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v92 + int32(1)
	if v56 != 0 {
		v51 = v56
		goto L21
	} else {
		goto L30
	}
L30:
	;
	goto L22
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+28))
	if v110 == int32(0) {
		v116 = v100
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_valkey_free(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L25
	} else {
		goto L35
	}
L33:
	;
	m.T0[v110].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v116 = v115
	goto L32
L35:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v121
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	v124 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v124)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v123)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(-4294967296)
	goto L1
}
func F_dictCStrKeyCaseCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return base.B2i32(v37-v39 == int32(0))
L2:
	;
	v37 = F_tolower(m, v33)
	mBase = m.M
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v39 = F_tolower(m, v38)
	mBase = m.M
	goto L1
L3:
	;
	v7 = l0
	v8 = l1
	v9 = v5
	goto L6
L4:
	;
	v33 = int32(0)
	v34 = l1
	goto L2
L5:
	;
	v33 = v30 & int32(255)
	v34 = v29
	goto L2
L6:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v11 == int32(0) {
		v29 = v8
		v30 = v9
		goto L5
	} else {
		goto L8
	}
L7:
	;
	v29 = v23
	v30 = int32(0)
	goto L5
L8:
	;
	v15 = v9 & int32(255)
	if v15 == v11 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v22 = int32(1)
	v23 = v8 + v22
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	if v24 != 0 {
		v7 = v7 + v22
		v8 = v23
		v9 = v24
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v17 = F_tolower(m, v15)
	mBase = m.M
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v19 = F_tolower(m, v18)
	mBase = m.M
	if v17 == v19 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v29 = v8
	v30 = v21
	goto L5
L12:
	;
	goto L7
}
func F_dictEncObjHash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v101 int32
	_ = v101
	var v103 int64
	_ = v103
	var v108 int64
	_ = v108
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v11)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		v16 = F_objectGetVal(m, l0)
		mBase = m.M
		v18 = F_objectGetVal(m, l0)
		mBase = m.M
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(-1)))))
		switch v21 & int32(7) {
		case 0:
			v27 = F_siphash(m, v16, int32(base.Ui32(v21)>>(uint(int32(3))%32)), int32(_a245))
			mBase = m.M
			v108 = v27
		case 1:
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(-3)))))
			v32 = F_siphash(m, v16, v30, int32(_a245))
			mBase = m.M
			v108 = v32
		case 2:
			v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18+int32(-5)))))
			v37 = F_siphash(m, v16, v35, int32(_a245))
			mBase = m.M
			v108 = v37
		case 3:
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(-9))))
			v42 = F_siphash(m, v16, v40, int32(_a245))
			mBase = m.M
			v108 = v42
		case 4:
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(-17))))
			v46 = v45
			v48 = F_siphash(m, v16, v46, int32(_a245))
			mBase = m.M
			v108 = v48
		default:
			v46 = int32(0)
			v48 = F_siphash(m, v16, v46, int32(_a245))
			mBase = m.M
			v108 = v48
		}
		m.G0 = v9 + int32(32)
		return v108
	case 1:
		v59 = F_objectGetVal(m, l0)
		mBase = m.M
		v60 = base.I64_extend_i32_s(v59)
		if v60 <= int64(-1) {
			v69 = int32(45)
			*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v69)
			v73 = int32(1)
			v78 = v9 + v73
			v79 = int32(31)
			v80 = int64(0) - v60
			v81 = v73
		} else {
			v78 = v9
			v79 = int32(32)
			v80 = v60
			v81 = int32(0)
		}
		v82 = F_ull2string(m, v78, v79, v80)
		mBase = m.M
		if v82 == int32(0) {
			v101 = int32(0)
		} else {
			v101 = v82 + v81
		}
		v103 = F_siphash(m, v9, v101, int32(_a245))
		mBase = m.M
		v108 = v103
		m.G0 = v9 + int32(32)
		return v108
	default:
		F__serverPanic_1(m, int32(_a2157), int32(505), int32(_a2158), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int64(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_dictEntryMemUsage(m *base.Module, l0 int32) int32 {
	return int32(24)
}
func F_dictExpand(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 != int32(-1) {
		v25 = v5
		return v25
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if base.Ui32(l1) < base.Ui32(v9) {
			v25 = v5
			return v25
		} else {
			v11 = int32(1)
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
			if v14 == int32(255) {
				v18 = int32(0)
			} else {
				v18 = v11 << (uint(v14) % 32)
			}
			if base.Ui32(l1) <= base.Ui32(v18) {
				v25 = v11
				return v25
			} else {
				v21 = F_dictResizeWithOptionalCheck(m, l0, l1, int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = v21
					return v25
				}
			}
		}
	}
}
func F_dictFetchValue(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = F_dictFind(m, l0, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
			return v9
		} else {
			return int32(0)
		}
	}
}
func F_dictFindPositionForInsert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
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
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = m.T0[v15].(func(*base.Module, int32) int64)(m, l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L1
L3:
	;
	return int32(0)
L4:
	;
	v21 = int32(-1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v22 == int32(255) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v29 = int64(0)
	goto L7
L6:
	;
	v29 = base.I64_extend_i32_u(v21<<(uint(v22)%32) ^ v21)
	goto L7
L7:
	;
	v30 = v16 & v29
	v31 = base.I32_wrap_i64(v30)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v32 == int32(-1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	if int32(0) < v49 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	if v31 < v32 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
	if v45 != 0 {
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v31<<(uint(int32(2))%32))))
	if v40 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_dictBucketRehash(m, l0, v30)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L8
L14:
	;
	v47 = F_dictRehash(m, l0, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L8
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v31 < v54 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v52 = F_dictExpandIfNeeded(m, l0)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	return v199
L20:
	;
	v193 = int32(0)
	if l2 == v193 {
		v199 = v193
		goto L19
	} else {
		goto L52
	}
L21:
	;
	v177 = int32(2)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(4)+v170<<(uint(v177)%32))))
	v199 = v180 + v173<<(uint(v177)%32)
	goto L19
L22:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v117 = int32(-1)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	if v118 == int32(255) {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v58 = int32(-1)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v59 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v102 != int32(-1) {
		goto L22
	} else {
		goto L37
	}
L25:
	;
	v66 = int64(0)
	goto L27
L26:
	;
	v66 = base.I64_extend_i32_u(v58<<(uint(v59)%32) ^ v58)
	goto L27
L27:
	;
	v68 = base.I32_wrap_i64(v66 & v16)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v56+v68<<(uint(int32(2))%32))))
	if v72 == int32(0) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v79 = v72
	goto L29
L29:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if l1 == v84 {
		v188 = v79
		goto L20
	} else {
		goto L31
	}
L30:
	;
	goto L24
L31:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	if v87 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	if v92 != 0 {
		v79 = v92
		goto L29
	} else {
		goto L36
	}
L33:
	;
	v90 = m.T0[v87].(func(*base.Module, int32, int32) int32)(m, l1, v84)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	if v90 != 0 {
		v188 = v79
		goto L20
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	goto L30
L37:
	;
	v170 = int32(0)
	v173 = v68
	goto L21
L38:
	;
	v170 = int32(1)
	v173 = v127
	goto L21
L39:
	;
	v125 = int64(0)
	goto L41
L40:
	;
	v125 = base.I64_extend_i32_u(v117<<(uint(v118)%32) ^ v117)
	goto L41
L41:
	;
	v127 = base.I32_wrap_i64(v125 & v16)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v115+v127<<(uint(int32(2))%32))))
	if v131 == int32(0) {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v138 = v131
	goto L43
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	if l1 == v143 {
		v188 = v138
		goto L20
	} else {
		goto L45
	}
L44:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v152 != int32(-1) {
		goto L38
	} else {
		goto L51
	}
L45:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	if v146 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	if v151 != 0 {
		v138 = v151
		goto L43
	} else {
		goto L50
	}
L47:
	;
	v149 = m.T0[v146].(func(*base.Module, int32, int32) int32)(m, l1, v143)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	if v149 != 0 {
		v188 = v138
		goto L20
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	goto L44
L51:
	;
	v170 = int32(0)
	v173 = v127
	goto L21
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v188
	return int32(0)
}
func F_dictFreeUnlinkedEntry(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	if l1 == int32(0) {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if v8 == int32(0) {
			v15 = v7
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			if v16 == int32(0) {
				F_valkey_free(m, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				m.T0[v16].(func(*base.Module, int32))(m, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_valkey_free(m, l1)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			m.T0[v8].(func(*base.Module, int32))(m, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v15 = v14
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
				if v16 == int32(0) {
					F_valkey_free(m, l1)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						return
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					m.T0[v16].(func(*base.Module, int32))(m, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						F_valkey_free(m, l1)
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
func F_dictGetSafeIterator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+4)) = int64(4294967295)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+12)) = int64(1)
		return v4
	}
}
func F_dictGetSomeKeys(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v410 int32
	_ = v410
	var v425 int32
	_ = v425
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = v17 + v18
	if base.Ui32(v19) < base.Ui32(l2) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v119 = int32(0)
	F___lock(m, int32(9116960))
	mBase = m.M
	v126 = *(*int32)(unsafe.Add(mBase, _consts[303]))
	v128 = *(*int32)(unsafe.Add(mBase, _consts[304]))
	if v128 != 0 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	v79 = int32(0)
	v81 = int32(-1)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v82 == int32(255) {
		goto L19
	} else {
		goto L20
	}
L3:
	;
	v21 = v19
	goto L5
L4:
	;
	v21 = l2
	goto L5
L5:
	;
	if v21 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v27 = int32(0)
	goto L7
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v41 != int32(-1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
	if v54 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v44 = int32(0)
	v46 = int32(-1)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v47 == int32(255) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = v44
	goto L13
L12:
	;
	v53 = v46<<(uint(v47)%32) ^ v46
	goto L13
L13:
	;
	v108 = v44
	v109 = v53
	goto L1
L14:
	;
	v61 = v27 + int32(1)
	if v61 != v21 {
		v27 = v61
		goto L7
	} else {
		goto L18
	}
L15:
	;
	v56 = F_dictRehash(m, l0, int32(1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	goto L14
L18:
	;
	goto L8
L19:
	;
	v88 = v79
	goto L21
L20:
	;
	v88 = v81<<(uint(v82)%32) ^ v81
	goto L21
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v89 == int32(-1) {
		v108 = v79
		v109 = v88
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v92 = int32(-1)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	v96 = v92<<(uint(v93)%32) ^ v92
	if base.Ui32(v96) < base.Ui32(v88) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v98 = v88
	goto L25
L24:
	;
	v98 = v96
	goto L25
L25:
	;
	if v93 == int32(255) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v101 = v88
	goto L28
L27:
	;
	v101 = v98
	goto L28
L28:
	;
	v108 = int32(1)
	v109 = v101
	goto L1
L29:
	;
	if v21 != 0 {
		goto L40
	} else {
		goto L41
	}
L30:
	;
	F___unlock(m, int32(9116960))
	mBase = m.M
	goto L29
L31:
	;
	v132 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, _consts[305]))
	v134 = int32(2)
	v136 = v126 + v133<<(uint(v134)%32)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v139 = *(*int32)(unsafe.Add(mBase, _consts[306]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v126+v139<<(uint(v134)%32))))
	v144 = v137 + v143
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v144
	v149 = v139 + int32(1)
	if v149 == v128 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v130 = F_lcg31(m, v129)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v130
	v163 = v130
	goto L30
L33:
	;
	v151 = v132
	goto L35
L34:
	;
	v151 = v149
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, _consts[306])) = v151
	v153 = int32(0)
	v156 = v133 + int32(1)
	if v156 == v128 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v158 = v153
	goto L38
L37:
	;
	v158 = v156
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, _consts[305])) = v158
	v163 = int32(base.Ui32(v144) >> (uint(int32(1)) % 32))
	goto L30
L39:
	;
	if base.Ui32(v410) < base.Ui32(v21) {
		goto L96
	} else {
		goto L97
	}
L40:
	;
	v171 = int32(4)
	if base.Ui32(v171) < base.Ui32(v21) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v410 = int32(0)
	goto L39
L42:
	;
	v174 = v21
	goto L44
L43:
	;
	v174 = v171
	goto L44
L44:
	;
	v179 = int32(0)
	v183 = v179
	v184 = v163
	v188 = v21 * int32(10)
	v192 = v179
	goto L45
L45:
	;
	if v188 == int32(0) {
		v410 = v183
		goto L39
	} else {
		goto L47
	}
L46:
	;
	v410 = v389
	goto L39
L47:
	;
	v206 = v183
	v207 = v184 & v109
	v215 = v192
	v216 = int32(0)
	v217 = int32(1)
	goto L48
L48:
	;
	v220 = v108 & v217
	if v220 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	if base.Ui32(v389) < base.Ui32(v21) {
		v183 = v389
		v184 = v390 + int32(1)
		v188 = v188 + int32(-1)
		v192 = v398
		goto L45
	} else {
		goto L95
	}
L50:
	;
	if v220 != 0 {
		v206 = v389
		v207 = v390
		v215 = v398
		v216 = int32(1)
		v217 = int32(0)
		goto L48
	} else {
		goto L94
	}
L51:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(26)+v216))))
	if v234 != int32(255) {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v207) < base.Ui32(v223) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v231 = v207
	goto L51
L54:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	if v225 == int32(255) {
		v231 = v223
		goto L51
	} else {
		goto L56
	}
L55:
	;
	v231 = v207
	goto L51
L56:
	;
	if int32(base.Ui32(v207)>>(uint(v225)%32)) == int32(0) {
		v389 = v206
		v390 = v207
		v398 = v215
		goto L50
	} else {
		goto L57
	}
L57:
	;
	v231 = v223
	goto L51
L58:
	;
	if int32(base.Ui32(v231)>>(uint(v234)%32)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v389 = v206
	v390 = v231
	v398 = v215
	goto L50
L60:
	;
	v240 = int32(2)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(4)+v216<<(uint(v240)%32))))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v243+v231<<(uint(v240)%32))))
	if v247 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v389 = v206
	v390 = v231
	v398 = v215
	goto L50
L62:
	;
	v304 = v206
	v305 = v247
	goto L76
L63:
	;
	v249 = v215 + int32(1)
	if base.Ui32(v174) < base.Ui32(v249) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v251 = int32(0)
	F___lock(m, int32(9116960))
	mBase = m.M
	v258 = *(*int32)(unsafe.Add(mBase, _consts[303]))
	v260 = *(*int32)(unsafe.Add(mBase, _consts[304]))
	if v260 != 0 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v389 = v206
	v390 = v231
	v398 = v249
	goto L50
L66:
	;
	v389 = v206
	v390 = v295 & v109
	v398 = int32(0)
	goto L50
L67:
	;
	F___unlock(m, int32(9116960))
	mBase = m.M
	goto L66
L68:
	;
	v264 = int32(0)
	v265 = *(*int32)(unsafe.Add(mBase, _consts[305]))
	v266 = int32(2)
	v268 = v258 + v265<<(uint(v266)%32)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v271 = *(*int32)(unsafe.Add(mBase, _consts[306]))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v258+v271<<(uint(v266)%32))))
	v276 = v269 + v275
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v276
	v281 = v271 + int32(1)
	if v281 == v260 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v262 = F_lcg31(m, v261)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v262
	v295 = v262
	goto L67
L70:
	;
	v283 = v264
	goto L72
L71:
	;
	v283 = v281
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, _consts[306])) = v283
	v285 = int32(0)
	v288 = v265 + int32(1)
	if v288 == v260 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v290 = v285
	goto L75
L74:
	;
	v290 = v288
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, _consts[305])) = v290
	v295 = int32(base.Ui32(v276) >> (uint(int32(1)) % 32))
	goto L67
L76:
	;
	if base.Ui32(v21) <= base.Ui32(v304) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	if base.Ui32(v21) <= base.Ui32(v382) {
		v410 = v382
		goto L39
	} else {
		goto L93
	}
L78:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v305)+16))
	if v384 != 0 {
		v304 = v382
		v305 = v384
		goto L76
	} else {
		goto L92
	}
L79:
	;
	v325 = int32(0)
	F___lock(m, int32(9116960))
	mBase = m.M
	v332 = *(*int32)(unsafe.Add(mBase, _consts[303]))
	v334 = *(*int32)(unsafe.Add(mBase, _consts[304]))
	if v334 != 0 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v304<<(uint(int32(2))%32)))) = v305
	v382 = v304 + int32(1)
	goto L78
L81:
	;
	v375 = v304 + int32(1)
	v376 = base.I32_rem_u_s(v369, v375)
	if base.Ui32(v21) <= base.Ui32(v376) {
		v382 = v375
		goto L78
	} else {
		goto L91
	}
L82:
	;
	F___unlock(m, int32(9116960))
	mBase = m.M
	goto L81
L83:
	;
	v338 = int32(0)
	v339 = *(*int32)(unsafe.Add(mBase, _consts[305]))
	v340 = int32(2)
	v342 = v332 + v339<<(uint(v340)%32)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
	v345 = *(*int32)(unsafe.Add(mBase, _consts[306]))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v332+v345<<(uint(v340)%32))))
	v350 = v343 + v349
	*(*int32)(unsafe.Add(mBase, uint32(v342))) = v350
	v355 = v345 + int32(1)
	if v355 == v334 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v336 = F_lcg31(m, v335)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v332))) = v336
	v369 = v336
	goto L82
L85:
	;
	v357 = v338
	goto L87
L86:
	;
	v357 = v355
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, _consts[306])) = v357
	v359 = int32(0)
	v362 = v339 + int32(1)
	if v362 == v334 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v364 = v359
	goto L90
L89:
	;
	v364 = v362
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, _consts[305])) = v364
	v369 = int32(base.Ui32(v350) >> (uint(int32(1)) % 32))
	goto L82
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v376<<(uint(int32(2))%32)))) = v305
	v382 = v375
	goto L78
L92:
	;
	goto L77
L93:
	;
	v389 = v382
	v390 = v231
	v398 = int32(0)
	goto L50
L94:
	;
	goto L49
L95:
	;
	goto L46
L96:
	;
	v425 = v410
	goto L98
L97:
	;
	v425 = v21
	goto L98
L98:
	;
	return v425
}
func F_dictHashtableDestructor(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_hashtableRelease(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_dictIncrUnsignedIntegerVal(m *base.Module, l0 int32, l1 int64) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = v3 + l1
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v4
	return v4
}
func F_dictMemUsage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	if v5 == int32(255) {
		v9 = int32(0)
	} else {
		v9 = int32(1) << (uint(v5) % 32)
	}
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v12 == int32(255) {
		v16 = int32(0)
	} else {
		v16 = int32(1) << (uint(v12) % 32)
	}
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	return (v9+v16)<<(uint(int32(2))%32) + (v20+v21)*int32(24)
}
func F_dictNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v44 int64
	_ = v44
	var v49 int64
	_ = v49
	var v54 int64
	_ = v54
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v65 int64
	_ = v65
	var v70 int64
	_ = v70
	var v75 int64
	_ = v75
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v86 int64
	_ = v86
	var v91 int64
	_ = v91
	var v96 int64
	_ = v96
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v107 int64
	_ = v107
	var v112 int64
	_ = v112
	var v117 int64
	_ = v117
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v128 int64
	_ = v128
	var v133 int64
	_ = v133
	var v138 int64
	_ = v138
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v149 int64
	_ = v149
	var v154 int64
	_ = v154
	var v159 int64
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	v9 = l0 + int32(20)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v16 = v9
	v17 = v13
	goto L4
L2:
	;
	v13 = int32(1)
	goto L1
L3:
	;
	v13 = int32(0)
	goto L1
L4:
	;
	switch v17 {
	case 0:
		goto L9
	default:
		goto L8
	}
L6:
	;
	v17 = int32(0)
	goto L4
L7:
	;
	return v219
L8:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v211
	if v211 == int32(0) {
		goto L6
	} else {
		goto L31
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 != int32(-1) {
		v174 = v21
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v175 = int32(1)
	v176 = v174 + v175
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176
	v178 = int32(0)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v182+int32(26)))))
	if v186 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v25 != 0 {
		v174 = int32(-1)
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v27 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	if v168 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	v34 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v26)+16)))
	v35 = int64(*(*int8)(unsafe.Add(mBase, uint32(v26)+27)))
	v36 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+8)))
	v37 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v26)+12)))
	v38 = int64(*(*int8)(unsafe.Add(mBase, uint32(v26)+26)))
	v39 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+4)))
	v40 = int64(21)
	v44 = v39<<(uint(v40)%64) + (v39 ^ int64(-1))
	v49 = (int64(base.Ui64(v44)>>(uint(int64(24))%64)) ^ v44) * int64(265)
	v54 = (int64(base.Ui64(v49)>>(uint(int64(14))%64)) ^ v49) * v40
	goto L16
L15:
	;
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+24)))
	v32 = v30 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+24)) = uint16(v32)
	v167 = v26
	goto L13
L16:
	;
	v60 = v38 + (int64(base.Ui64(v54)>>(uint(int64(28))%64))^v54)*int64(2147483649)
	v61 = int64(21)
	v65 = v60<<(uint(v61)%64) + (v60 ^ int64(-1))
	v70 = (int64(base.Ui64(v65)>>(uint(int64(24))%64)) ^ v65) * int64(265)
	v75 = (int64(base.Ui64(v70)>>(uint(int64(14))%64)) ^ v70) * v61
	goto L17
L17:
	;
	v81 = v37 + (int64(base.Ui64(v75)>>(uint(int64(28))%64))^v75)*int64(2147483649)
	v82 = int64(21)
	v86 = v81<<(uint(v82)%64) + (v81 ^ int64(-1))
	v91 = (int64(base.Ui64(v86)>>(uint(int64(24))%64)) ^ v86) * int64(265)
	v96 = (int64(base.Ui64(v91)>>(uint(int64(14))%64)) ^ v91) * v82
	goto L18
L18:
	;
	v102 = v36 + (int64(base.Ui64(v96)>>(uint(int64(28))%64))^v96)*int64(2147483649)
	v103 = int64(21)
	v107 = v102<<(uint(v103)%64) + (v102 ^ int64(-1))
	v112 = (int64(base.Ui64(v107)>>(uint(int64(24))%64)) ^ v107) * int64(265)
	v117 = (int64(base.Ui64(v112)>>(uint(int64(14))%64)) ^ v112) * v103
	goto L19
L19:
	;
	v123 = v35 + (int64(base.Ui64(v117)>>(uint(int64(28))%64))^v117)*int64(2147483649)
	v124 = int64(21)
	v128 = v123<<(uint(v124)%64) + (v123 ^ int64(-1))
	v133 = (int64(base.Ui64(v128)>>(uint(int64(24))%64)) ^ v128) * int64(265)
	v138 = (int64(base.Ui64(v133)>>(uint(int64(14))%64)) ^ v133) * v124
	goto L20
L20:
	;
	v144 = v34 + (int64(base.Ui64(v138)>>(uint(int64(28))%64))^v138)*int64(2147483649)
	v145 = int64(21)
	v149 = v144<<(uint(v145)%64) + (v144 ^ int64(-1))
	v154 = (int64(base.Ui64(v149)>>(uint(int64(24))%64)) ^ v149) * int64(265)
	v159 = (int64(base.Ui64(v154)>>(uint(int64(14))%64)) ^ v154) * v145
	goto L21
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = (int64(base.Ui64(v159)>>(uint(int64(28))%64)) ^ v159) * int64(2147483649)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v167 = v166
	goto L13
L22:
	;
	v174 = v168 + int32(-1)
	goto L10
L23:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v174 = v171
	goto L10
L24:
	;
	v201 = int32(2)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v181+v199<<(uint(v201)%32)+int32(4))))
	v16 = v206 + v200<<(uint(v201)%32)
	v17 = int32(1)
	goto L4
L25:
	;
	v190 = v178
	goto L27
L26:
	;
	v190 = v175 << (uint(v186) % 32)
	goto L27
L27:
	;
	if v176 < v190 {
		v199 = v182
		v200 = v176
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v182 != 0 {
		v219 = v178
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
	if v192 == int32(-1) {
		v219 = v178
		goto L7
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(4294967296)
	v199 = int32(1)
	v200 = int32(0)
	goto L24
L31:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v215
	v219 = v211
	goto L7
}
func F_dictObjHash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	v5 = F_objectGetVal(m, l0)
	mBase = m.M
	v7 = F_objectGetVal(m, l0)
	mBase = m.M
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-1)))))
	switch v10 & int32(7) {
	case 0:
		v16 = F_siphash(m, v5, int32(base.Ui32(v10)>>(uint(int32(3))%32)), int32(_a245))
		mBase = m.M
		return v16
	case 1:
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-3)))))
		v22 = F_siphash(m, v5, v20, int32(_a245))
		mBase = m.M
		return v22
	case 2:
		v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(-5)))))
		v28 = F_siphash(m, v5, v26, int32(_a245))
		mBase = m.M
		return v28
	case 3:
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-9))))
		v34 = F_siphash(m, v5, v32, int32(_a245))
		mBase = m.M
		return v34
	case 4:
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-17))))
		v39 = v38
		v41 = F_siphash(m, v5, v39, int32(_a245))
		mBase = m.M
		return v41
	default:
		v39 = int32(0)
		v41 = F_siphash(m, v5, v39, int32(_a245))
		mBase = m.M
		return v41
	}
}
func F_dictPtrHash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int64
	_ = v13
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	v13 = F_siphash(m, v6+int32(12), int32(4), int32(_a245))
	mBase = m.M
	m.G0 = v6 + int32(16)
	return v13
}
func F_dictReplace(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_dictAddRaw(m, l0, l1, v8+int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = l2
			v23 = int32(0)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
			if v25 == v23 {
				v32 = v23
				m.G0 = v8 + int32(16)
				return v32
			} else {
				m.T0[v25].(func(*base.Module, int32))(m, v21)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v32 = v23
					m.G0 = v8 + int32(16)
					return v32
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = l2
			v32 = int32(1)
			m.G0 = v8 + int32(16)
			return v32
		}
	}
}
func F_dictResetIterator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v37 int64
	_ = v37
	var v42 int64
	_ = v42
	var v47 int64
	_ = v47
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v58 int64
	_ = v58
	var v63 int64
	_ = v63
	var v68 int64
	_ = v68
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v79 int64
	_ = v79
	var v84 int64
	_ = v84
	var v89 int64
	_ = v89
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v100 int64
	_ = v100
	var v105 int64
	_ = v105
	var v110 int64
	_ = v110
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v121 int64
	_ = v121
	var v126 int64
	_ = v126
	var v131 int64
	_ = v131
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v142 int64
	_ = v142
	var v147 int64
	_ = v147
	var v152 int64
	_ = v152
	var v163 int32
	_ = v163
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2 != int32(-1) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 == int32(0) {
			v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v27 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v26)+16)))
			v28 = int64(*(*int8)(unsafe.Add(mBase, uint32(v26)+27)))
			v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+8)))
			v30 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v26)+12)))
			v31 = int64(*(*int8)(unsafe.Add(mBase, uint32(v26)+26)))
			v32 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+4)))
			v33 = int64(21)
			v37 = v32<<(uint(v33)%64) + (v32 ^ int64(-1))
			v42 = (int64(base.Ui64(v37)>>(uint(int64(24))%64)) ^ v37) * int64(265)
			v47 = (int64(base.Ui64(v42)>>(uint(int64(14))%64)) ^ v42) * v33
			v53 = v31 + (int64(base.Ui64(v47)>>(uint(int64(28))%64))^v47)*int64(2147483649)
			v54 = int64(21)
			v58 = v53<<(uint(v54)%64) + (v53 ^ int64(-1))
			v63 = (int64(base.Ui64(v58)>>(uint(int64(24))%64)) ^ v58) * int64(265)
			v68 = (int64(base.Ui64(v63)>>(uint(int64(14))%64)) ^ v63) * v54
			v74 = v30 + (int64(base.Ui64(v68)>>(uint(int64(28))%64))^v68)*int64(2147483649)
			v75 = int64(21)
			v79 = v74<<(uint(v75)%64) + (v74 ^ int64(-1))
			v84 = (int64(base.Ui64(v79)>>(uint(int64(24))%64)) ^ v79) * int64(265)
			v89 = (int64(base.Ui64(v84)>>(uint(int64(14))%64)) ^ v84) * v75
			v95 = v29 + (int64(base.Ui64(v89)>>(uint(int64(28))%64))^v89)*int64(2147483649)
			v96 = int64(21)
			v100 = v95<<(uint(v96)%64) + (v95 ^ int64(-1))
			v105 = (int64(base.Ui64(v100)>>(uint(int64(24))%64)) ^ v100) * int64(265)
			v110 = (int64(base.Ui64(v105)>>(uint(int64(14))%64)) ^ v105) * v96
			v116 = v28 + (int64(base.Ui64(v110)>>(uint(int64(28))%64))^v110)*int64(2147483649)
			v117 = int64(21)
			v121 = v116<<(uint(v117)%64) + (v116 ^ int64(-1))
			v126 = (int64(base.Ui64(v121)>>(uint(int64(24))%64)) ^ v121) * int64(265)
			v131 = (int64(base.Ui64(v126)>>(uint(int64(14))%64)) ^ v126) * v117
			v137 = v27 + (int64(base.Ui64(v131)>>(uint(int64(28))%64))^v131)*int64(2147483649)
			v138 = int64(21)
			v142 = v137<<(uint(v138)%64) + (v137 ^ int64(-1))
			v147 = (int64(base.Ui64(v142)>>(uint(int64(24))%64)) ^ v142) * int64(265)
			v152 = (int64(base.Ui64(v147)>>(uint(int64(14))%64)) ^ v147) * v138
			if v25 == (int64(base.Ui64(v152)>>(uint(int64(28))%64))^v152)*int64(2147483649) {
				return
			} else {
				F__serverAssert(m, int32(_a598), int32(_a597), int32(833))
				mBase = m.M
				v163 = m.ExcPending
				if v163 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)))
			v13 = int32(-1)
			v14 = v12 + v13
			*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)) = uint16(v14)
			if v13 < base.I32_extend16_s(v14) {
				return
			} else {
				F__serverAssert(m, int32(_a599), int32(_a597), int32(831))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
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
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v5 == int32(0) {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v8 == int32(0) {
				v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v27 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v26)+16)))
				v28 = int64(*(*int8)(unsafe.Add(mBase, uint32(v26)+27)))
				v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+8)))
				v30 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v26)+12)))
				v31 = int64(*(*int8)(unsafe.Add(mBase, uint32(v26)+26)))
				v32 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+4)))
				v33 = int64(21)
				v37 = v32<<(uint(v33)%64) + (v32 ^ int64(-1))
				v42 = (int64(base.Ui64(v37)>>(uint(int64(24))%64)) ^ v37) * int64(265)
				v47 = (int64(base.Ui64(v42)>>(uint(int64(14))%64)) ^ v42) * v33
				v53 = v31 + (int64(base.Ui64(v47)>>(uint(int64(28))%64))^v47)*int64(2147483649)
				v54 = int64(21)
				v58 = v53<<(uint(v54)%64) + (v53 ^ int64(-1))
				v63 = (int64(base.Ui64(v58)>>(uint(int64(24))%64)) ^ v58) * int64(265)
				v68 = (int64(base.Ui64(v63)>>(uint(int64(14))%64)) ^ v63) * v54
				v74 = v30 + (int64(base.Ui64(v68)>>(uint(int64(28))%64))^v68)*int64(2147483649)
				v75 = int64(21)
				v79 = v74<<(uint(v75)%64) + (v74 ^ int64(-1))
				v84 = (int64(base.Ui64(v79)>>(uint(int64(24))%64)) ^ v79) * int64(265)
				v89 = (int64(base.Ui64(v84)>>(uint(int64(14))%64)) ^ v84) * v75
				v95 = v29 + (int64(base.Ui64(v89)>>(uint(int64(28))%64))^v89)*int64(2147483649)
				v96 = int64(21)
				v100 = v95<<(uint(v96)%64) + (v95 ^ int64(-1))
				v105 = (int64(base.Ui64(v100)>>(uint(int64(24))%64)) ^ v100) * int64(265)
				v110 = (int64(base.Ui64(v105)>>(uint(int64(14))%64)) ^ v105) * v96
				v116 = v28 + (int64(base.Ui64(v110)>>(uint(int64(28))%64))^v110)*int64(2147483649)
				v117 = int64(21)
				v121 = v116<<(uint(v117)%64) + (v116 ^ int64(-1))
				v126 = (int64(base.Ui64(v121)>>(uint(int64(24))%64)) ^ v121) * int64(265)
				v131 = (int64(base.Ui64(v126)>>(uint(int64(14))%64)) ^ v126) * v117
				v137 = v27 + (int64(base.Ui64(v131)>>(uint(int64(28))%64))^v131)*int64(2147483649)
				v138 = int64(21)
				v142 = v137<<(uint(v138)%64) + (v137 ^ int64(-1))
				v147 = (int64(base.Ui64(v142)>>(uint(int64(24))%64)) ^ v142) * int64(265)
				v152 = (int64(base.Ui64(v147)>>(uint(int64(14))%64)) ^ v147) * v138
				if v25 == (int64(base.Ui64(v152)>>(uint(int64(28))%64))^v152)*int64(2147483649) {
					return
				} else {
					F__serverAssert(m, int32(_a598), int32(_a597), int32(833))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)))
				v13 = int32(-1)
				v14 = v12 + v13
				*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)) = uint16(v14)
				if v13 < base.I32_extend16_s(v14) {
					return
				} else {
					F__serverAssert(m, int32(_a599), int32(_a597), int32(831))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
	}
}
func F_dictSetUnsignedIntegerVal(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = l1
	return
}
func F_dictSetVal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
	return
}
func F_dictShrink(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v5 = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 != int32(-1) {
		v20 = v5
		return v20
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if base.Ui32(l1) < base.Ui32(v9) {
			v20 = v5
			return v20
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
			if v11 == int32(255) {
				v20 = v5
				return v20
			} else {
				if int32(base.Ui32(l1)>>(uint(v11)%32)) != 0 {
					v20 = v5
					return v20
				} else {
					v16 = F_dictResizeWithOptionalCheck(m, l0, l1, int32(0))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						v20 = v16
						return v20
					}
				}
			}
		}
	}
}
func F_dictShrinkIfNeeded(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v7 != int32(-1) {
		v75 = v2
		return v75
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
		if base.Ui32((v10+int32(1))&int32(255)) < base.Ui32(int32(4)) {
			v75 = v2
			return v75
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[302]))
			switch v19 {
			case 0:
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v27 = int32(1) << (uint(v10) % 32)
				if base.Ui32(v23<<(uint(int32(3))%32)) <= base.Ui32(v27) {
					v37 = v23
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
					if v39 == int32(0) {
						v69 = v37
						v70 = F_dictShrink(m, l0, v69)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							v75 = int32(0)
							return v75
						}
					} else {
						if base.Ui32(int32(5)) <= base.Ui32(v37) {
							if base.Ui32(v37) <= base.Ui32(int32(2147483646)) {
								v53 = int32(32) - base.I32_clz(v37+int32(-1))
							} else {
								v53 = int32(31)
							}
						} else {
							v53 = int32(2)
						}
						v61 = m.T0[v39].(func(*base.Module, int32, float64) int32)(m, int32(4)<<(uint(v53)%32), base.F64_div(base.F64_convert_i32_u(v37), base.F64_convert_i32_u(int32(1)<<(uint(v10)%32))))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							if v61 == int32(0) {
								v75 = int32(0)
								return v75
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v69 = v67
								v70 = F_dictShrink(m, l0, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									v75 = int32(0)
									return v75
								}
							}
						}
					}
				} else {
					v29 = v27
					v30 = v23
					if base.Ui32(v29) < base.Ui32(v30<<(uint(int32(5))%32)) {
						v75 = int32(1)
						return v75
					} else {
						v37 = v30
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
						if v39 == int32(0) {
							v69 = v37
							v70 = F_dictShrink(m, l0, v69)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v75 = int32(0)
								return v75
							}
						} else {
							if base.Ui32(int32(5)) <= base.Ui32(v37) {
								if base.Ui32(v37) <= base.Ui32(int32(2147483646)) {
									v53 = int32(32) - base.I32_clz(v37+int32(-1))
								} else {
									v53 = int32(31)
								}
							} else {
								v53 = int32(2)
							}
							v61 = m.T0[v39].(func(*base.Module, int32, float64) int32)(m, int32(4)<<(uint(v53)%32), base.F64_div(base.F64_convert_i32_u(v37), base.F64_convert_i32_u(int32(1)<<(uint(v10)%32))))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								if v61 == int32(0) {
									v75 = int32(0)
									return v75
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v69 = v67
									v70 = F_dictShrink(m, l0, v69)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v75 = int32(0)
										return v75
									}
								}
							}
						}
					}
				}
			default:
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v29 = int32(1) << (uint(v10) % 32)
				v30 = v22
				if base.Ui32(v29) < base.Ui32(v30<<(uint(int32(5))%32)) {
					v75 = int32(1)
					return v75
				} else {
					v37 = v30
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
					if v39 == int32(0) {
						v69 = v37
						v70 = F_dictShrink(m, l0, v69)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							v75 = int32(0)
							return v75
						}
					} else {
						if base.Ui32(int32(5)) <= base.Ui32(v37) {
							if base.Ui32(v37) <= base.Ui32(int32(2147483646)) {
								v53 = int32(32) - base.I32_clz(v37+int32(-1))
							} else {
								v53 = int32(31)
							}
						} else {
							v53 = int32(2)
						}
						v61 = m.T0[v39].(func(*base.Module, int32, float64) int32)(m, int32(4)<<(uint(v53)%32), base.F64_div(base.F64_convert_i32_u(v37), base.F64_convert_i32_u(int32(1)<<(uint(v10)%32))))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							if v61 == int32(0) {
								v75 = int32(0)
								return v75
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v69 = v67
								v70 = F_dictShrink(m, l0, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									v75 = int32(0)
									return v75
								}
							}
						}
					}
				}
			case 2:
				v75 = int32(1)
				return v75
			}
		}
	}
}
func F_dictStrCaseHash_3(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	if l0&int32(3) == int32(0) {
		v23 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v58 = F_siphash_nocase(m, l0, v56, int32(_a245))
	mBase = m.M
	goto L17
L2:
	;
	v56 = v48 - l0
	goto L1
L3:
	;
	v27 = v23
	goto L11
L4:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v12 = l0
	goto L7
L6:
	;
	v56 = l0 - l0
	goto L1
L7:
	;
	v16 = v12 + int32(1)
	if v16&int32(3) == int32(0) {
		v23 = v16
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v21 != 0 {
		v12 = v16
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v48 = v16
	goto L2
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v36 = int32(-2139062144)
	if (int32(16843008)-v33|v33)&v36 == v36 {
		v27 = v27 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v42 = v27
	goto L14
L13:
	;
	goto L12
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v46 != 0 {
		v42 = v42 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v48 = v42
	goto L2
L16:
	;
	goto L15
L17:
	;
	return v58
}
func F_dictStringKeyCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v6 == int32(0) {
		v29 = v5
		v30 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.B2i32(v30-v29&int32(255) == int32(0))
L2:
	;
	goto L1
L3:
	;
	if v6 != v5&int32(255) {
		v29 = v5
		v30 = v6
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v12 = l0
	v13 = l1
	goto L5
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v17 == int32(0) {
		v29 = v16
		v30 = v17
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v29 = v16
	v30 = v17
	goto L2
L7:
	;
	v20 = int32(1)
	if v17 == v16&int32(255) {
		v12 = v12 + v20
		v13 = v13 + v20
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
func F_dictVanillaFree(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_valkey_free(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
