package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_kvstoreExpand(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
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
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	if base.B2i32(l1 == int64(0)) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v17 < v16 {
		v120 = v16
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return int32(1)
L3:
	;
	return v120
L4:
	;
	v20 = base.I32_wrap_i64(l1)
	v28 = int32(0)
	goto L5
L5:
	;
	if l3 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v120 = v111
	goto L3
L7:
	;
	v111 = int32(1)
	v113 = v28 + v111
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v113 < v114 {
		v28 = v113
		goto L5
	} else {
		goto L30
	}
L8:
	;
	if l2 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v33 = m.T0[l3].(func(*base.Module, int32) int32)(m, v28)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v33 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	if v20 == int32(0) {
		goto L7
	} else {
		goto L23
	}
L14:
	;
	v39 = int32(0)
	if v20 == v39 {
		v120 = v39
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v44 = v28 << (uint(int32(2)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v44)))
	if v46 != 0 {
		v71 = v46
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v72 = F_hashtableTryExpand(m, v71, v20)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L21
	}
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v48 = F_hashtableCreate(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v50+v44))) = v48
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48+int32(44))+4)) = l0
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56+v44)))
	v59 = F_hashtableMemUsage(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v59 + v61
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v64 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v68+v44)))
	v71 = v70
	goto L16
L21:
	;
	if v72 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v120 = v39
	goto L3
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v78 = v28 << (uint(int32(2)) % 32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v78)))
	if v80 != 0 {
		v105 = v80
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v106 = F_hashtableExpand(m, v105, v20)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L10
	} else {
		goto L29
	}
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v82 = F_hashtableCreate(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v84+v78))) = v82
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82+int32(44))+4)) = l0
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90+v78)))
	v93 = F_hashtableMemUsage(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v93 + v95
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v98 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102+v78)))
	v105 = v104
	goto L24
L29:
	;
	goto L7
L30:
	;
	goto L6
}
func F_kvstoreGetFairRandomHashtableIndex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 == int32(1) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v19 = base.I64_extend_i32_u(v15 + v16)
			if base.B2i32(v19 == int64(0)) == int32(0) {
				v26 = int32(0)
				F___lock(m, int32(9116960))
				mBase = m.M
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_kvstoreGetFairRandomHashtableIndex[0]))
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_kvstoreGetFairRandomHashtableIndex[1]))
				if v35 != 0 {
					v39 = int32(0)
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_kvstoreGetFairRandomHashtableIndex[2]))
					v41 = int32(2)
					v43 = v33 + v40<<(uint(v41)%32)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_kvstoreGetFairRandomHashtableIndex[3]))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v33+v46<<(uint(v41)%32))))
					v51 = v44 + v50
					*(*int32)(unsafe.Add(mBase, uint32(v43))) = v51
					v56 = v46 + int32(1)
					if v56 == v35 {
						v58 = v39
					} else {
						v58 = v56
					}
					*(*int32)(unsafe.Add(mBase, _c_F_kvstoreGetFairRandomHashtableIndex[3])) = v58
					v60 = int32(0)
					v63 = v40 + int32(1)
					if v63 == v35 {
						v65 = v60
					} else {
						v65 = v63
					}
					*(*int32)(unsafe.Add(mBase, _c_F_kvstoreGetFairRandomHashtableIndex[2])) = v65
					v70 = int32(base.Ui32(v51) >> (uint(int32(1)) % 32))
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
					v37 = F_lcg31(m, v36)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v33))) = v37
					v70 = v37
				}
				F___unlock(m, int32(9116960))
				mBase = m.M
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v76 == int32(1) {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
					v86 = base.I64_extend_i32_u(v82 + v83)
				} else {
					v79 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
					v86 = v79
				}
				v87 = base.I64_rem_u_s(base.I64_extend_i32_s(v70), v86)
				v93 = base.I32_wrap_i64(v87) + int32(1)
			} else {
				v93 = int32(0)
			}
		} else {
			v93 = int32(0)
		}
	} else {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = v11
		if base.B2i32(v19 == int64(0)) == int32(0) {
			v26 = int32(0)
			F___lock(m, int32(9116960))
			mBase = m.M
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_kvstoreGetFairRandomHashtableIndex[0]))
			v35 = *(*int32)(unsafe.Add(mBase, _c_F_kvstoreGetFairRandomHashtableIndex[1]))
			if v35 != 0 {
				v39 = int32(0)
				v40 = *(*int32)(unsafe.Add(mBase, _c_F_kvstoreGetFairRandomHashtableIndex[2]))
				v41 = int32(2)
				v43 = v33 + v40<<(uint(v41)%32)
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_kvstoreGetFairRandomHashtableIndex[3]))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v33+v46<<(uint(v41)%32))))
				v51 = v44 + v50
				*(*int32)(unsafe.Add(mBase, uint32(v43))) = v51
				v56 = v46 + int32(1)
				if v56 == v35 {
					v58 = v39
				} else {
					v58 = v56
				}
				*(*int32)(unsafe.Add(mBase, _c_F_kvstoreGetFairRandomHashtableIndex[3])) = v58
				v60 = int32(0)
				v63 = v40 + int32(1)
				if v63 == v35 {
					v65 = v60
				} else {
					v65 = v63
				}
				*(*int32)(unsafe.Add(mBase, _c_F_kvstoreGetFairRandomHashtableIndex[2])) = v65
				v70 = int32(base.Ui32(v51) >> (uint(int32(1)) % 32))
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
				v37 = F_lcg31(m, v36)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v33))) = v37
				v70 = v37
			}
			F___unlock(m, int32(9116960))
			mBase = m.M
			v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v76 == int32(1) {
				v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
				v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
				v86 = base.I64_extend_i32_u(v82 + v83)
			} else {
				v79 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				v86 = v79
			}
			v87 = base.I64_rem_u_s(base.I64_extend_i32_s(v70), v86)
			v93 = base.I32_wrap_i64(v87) + int32(1)
		} else {
			v93 = int32(0)
		}
	}
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v96 == int32(1) {
		v140 = int32(0)
		return v140
	} else {
		v99 = int32(-1)
		if v93 == int32(0) {
			v140 = v99
			return v140
		} else {
			v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
			if v102 == int64(0) {
				v140 = v99
				return v140
			} else {
				if base.Ui64(v102) < base.Ui64(base.I64_extend_i32_u(v93)) {
					F__serverAssert(m, int32(_a_F_kvstoreGetFairRandomHashtableIndex_0), int32(_a_F_kvstoreGetFairRandomHashtableIndex_1), int32(576))
					mBase = m.M
					v149 = m.ExcPending
					if v149 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v114 = int32(1) << (uint(v108) % 32)
					v115 = v93
					v117 = int32(0)
					for {
						v119 = v117 + v114
						v123 = *(*int64)(unsafe.Add(mBase, uint32(v110+v119<<(uint(int32(3))%32))))
						v125 = base.B2i32(base.Ui64(v123) < base.Ui64(base.I64_extend_i32_u(v115)))
						if base.Ui64(v123) < base.Ui64(base.I64_extend_i32_u(v115)) {
							v126 = v119
						} else {
							v126 = v117
						}
						if base.Ui64(v123) < base.Ui64(base.I64_extend_i32_u(v115)) {
							v129 = base.I32_wrap_i64(v123)
						} else {
							v129 = int32(0)
						}
						v131 = int32(1)
						if base.Ui32(v131) < base.Ui32(v114) {
							v114 = v114 >> (uint(v131) % 32)
							v115 = v115 - v129
							v117 = v126
							continue
						} else {
							break
						}
						break
					}
					v140 = v126
					return v140
				}
			}
		}
	}
}
func F_kvstoreHashtableAdd(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = l1 << (uint(int32(2)) % 32)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6+v8)))
	if v10 != 0 {
		v37 = v10
		v38 = F_hashtableAdd(m, v37, l2)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			if v38 == int32(0) {
				return v38
			} else {
				F_cumulativeKeyCountAdd(m, l0, l1, int32(1))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					return v38
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v12 = F_hashtableCreate(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v16+v8))) = v12
			*(*int32)(unsafe.Add(mBase, uint32(v12+int32(44))+4)) = l0
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v22+v8)))
			v25 = F_hashtableMemUsage(m, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v25 + v27
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v30 + int32(1)
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v34+v8)))
				v37 = v36
				v38 = F_hashtableAdd(m, v37, l2)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					if v38 == int32(0) {
						return v38
					} else {
						F_cumulativeKeyCountAdd(m, l0, l1, int32(1))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							return v38
						}
					}
				}
			}
		}
	}
}
func F_kvstoreHashtableFind(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
		v12 = F_hashtableFind(m, v9, l2, l3)
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
func F_kvstoreHashtableFindPositionForInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = l1 << (uint(int32(2)) % 32)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7+v9)))
	if v11 != 0 {
		v38 = v11
		v39 = F_hashtableFindPositionForInsert(m, v38, l2, l3, l4)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			return v39
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v13 = F_hashtableCreate(m, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v17+v9))) = v13
			*(*int32)(unsafe.Add(mBase, uint32(v13+int32(44))+4)) = l0
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v23+v9)))
			v26 = F_hashtableMemUsage(m, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v26 + v28
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v31 + int32(1)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v35+v9)))
				v38 = v37
				v39 = F_hashtableFindPositionForInsert(m, v38, l2, l3, l4)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					return v39
				}
			}
		}
	}
}
func F_kvstoreHashtableMetadataSize(m *base.Module) int32 {
	return int32(8)
}
func F_kvstoreHashtableRandomEntry(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
		v11 = F_hashtableRandomEntry(m, v8, l2)
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
func F_kvstoreHashtableRehashingCompleted(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = l0 + int32(44)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v13 == int32(0) {
		F_hashtableRehashingInfo(m, l0, v8+int32(12), v8+int32(8))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v27 - base.I64_extend_i32_u(v28)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v32 - v28<<(uint(int32(6))%32)
			m.G0 = v8 + int32(16)
			return
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
		F_listDelNode(m, v16, v13)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
			F_hashtableRehashingInfo(m, l0, v8+int32(12), v8+int32(8))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v27 - base.I64_extend_i32_u(v28)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v32 - v28<<(uint(int32(6))%32)
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_kvstoreIteratorGetCurrentHashtableIndex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v15 int32
	_ = v15
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 < int64(0) {
		F__serverAssert(m, int32(_a_F_kvstoreIteratorGetCurrentHashtableIndex_0), int32(_a_F_kvstoreIteratorGetCurrentHashtableIndex_1), int32(699))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v7 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+12)))
		if v3 < v7 {
			return base.I32_wrap_i64(v3)
		} else {
			F__serverAssert(m, int32(_a_F_kvstoreIteratorGetCurrentHashtableIndex_0), int32(_a_F_kvstoreIteratorGetCurrentHashtableIndex_1), int32(699))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
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
func F_kvstoreIteratorInit(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int64
	_ = v59
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v2 = l1
	v10 = F_valkey_malloc(m, int32(80))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v18 == int32(1) {
			v59 = int64(0)
		} else {
			v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
			if v22 == int64(0) {
				v59 = int64(-1)
			} else {
				v25 = int32(1)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v31 = int32(0)
				v35 = v25 << (uint(v26) % 32)
				v37 = v25
				for {
					v39 = v31 + v35
					v43 = *(*int64)(unsafe.Add(mBase, uint32(v28+v39<<(uint(int32(3))%32))))
					v45 = base.B2i32(base.Ui64(v43) < base.Ui64(base.I64_extend_i32_u(v37)))
					if base.Ui64(v43) < base.Ui64(base.I64_extend_i32_u(v37)) {
						v46 = v39
					} else {
						v46 = v31
					}
					if base.Ui64(v43) < base.Ui64(base.I64_extend_i32_u(v37)) {
						v49 = base.I32_wrap_i64(v43)
					} else {
						v49 = int32(0)
					}
					v51 = int32(1)
					if base.Ui32(v51) < base.Ui32(v35) {
						v31 = v46
						v35 = v35 >> (uint(v51) % 32)
						v37 = v37 - v49
						continue
					} else {
						break
					}
					break
				}
				v59 = base.I64_extend_i32_s(v46)
			}
		}
		v64 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+76)) = v64
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+72)) = uint8(v2)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v59
		v69 = v10 + int32(24)
		*(*uint8)(unsafe.Add(mBase, uint32(v69)+14)) = uint8(v64)
		*(*int32)(unsafe.Add(mBase, uint32(v69))) = v64
		*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = v64
		*(*uint8)(unsafe.Add(mBase, uint32(v69)+15)) = uint8(v2)
		*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = int32(-1)
		return v10
	}
}
func F_kvstoreIteratorNext(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v4 == int64(-1) {
		v17 = F_kvstoreIteratorNextHashtable(m, l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v17 != 0 {
				v22 = l0 + int32(24)
				F_hashtableRetargetIterator(m, v22, v17)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = F_hashtableNext(m, v22, l1)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return v25
					}
				}
			} else {
				return int32(0)
			}
		}
	} else {
		v9 = F_hashtableNext(m, l0+int32(24), l1)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v9 == int32(0) {
				v17 = F_kvstoreIteratorNextHashtable(m, l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if v17 != 0 {
						v22 = l0 + int32(24)
						F_hashtableRetargetIterator(m, v22, v17)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							v25 = F_hashtableNext(m, v22, l1)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								return v25
							}
						}
					} else {
						return int32(0)
					}
				}
			} else {
				return int32(1)
			}
		}
	}
}
func F_kvstoreIteratorNextHashtable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v77 int64
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v14 != int32(-1) {
		v70 = v14
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v164
L2:
	;
	v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v77 == int64(-1) {
		goto L21
	} else {
		goto L22
	}
L3:
	;
	v17 = int32(0)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v18&int32(8) == v17 {
		v164 = v17
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v23 != 0 {
		v32 = v23
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v35 = F_hashtableNext(m, v32, v12+int32(12))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L9
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	v27 = F_hashtableCreateIterator(m, v25, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v27
	v32 = v27
	goto L5
L9:
	;
	if v35 == int32(0) {
		v164 = v17
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L12
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v65 == int32(-1) {
		v164 = v17
		goto L1
	} else {
		goto L20
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49+v50<<(uint(int32(2))%32))))
	if v54 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v63 = F_hashtableNext(m, v60, v12+int32(12))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L18
	}
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	goto L16
L16:
	;
	if v57+v58 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	if v63 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v164 = v17
	goto L1
L20:
	;
	v70 = v65
	goto L2
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = base.I64_extend_i32_s(v70)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v145 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v145 == int64(-1) {
		v154 = v70
		v155 = v144
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81+base.I32_wrap_i64(v77)<<(uint(int32(2))%32))))
	if v86 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	F_hashtableCleanupIterator(m, l0+int32(24))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v94&int32(2) == int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99+v100<<(uint(int32(2))%32))))
	if v104 == int32(0) {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v104)+20))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	goto L27
L27:
	;
	if v107+v108 != 0 {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	v110 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	v113 = v100 << (uint(int32(2)) % 32)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111+v113)))
	if v115 == v110 {
		v126 = v110
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_hashtableRelease(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L33
	}
L30:
	;
	v118 = int32(*(*int16)(unsafe.Add(mBase, uint32(v115)+26)))
	goto L31
L31:
	;
	if int32(0) < v118 {
		goto L21
	} else {
		goto L32
	}
L32:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v100<<(uint(int32(2))%32))))
	v126 = v125
	goto L29
L33:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v129+v113))) = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v93)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v133 + int32(-1)
	goto L21
L34:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156+v154<<(uint(int32(2))%32))))
	v164 = v160
	goto L1
L35:
	;
	v148 = F_kvstoreGetNextNonEmptyHashtableIndex(m, v144, v70)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = base.I64_extend_i32_s(v148)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v154 = v152
	v155 = v153
	goto L34
}
func F_kvstoreNumHashtables(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	return v2
}
func F_kvstoreOverheadHashtableRehashing(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	return v2
}
func F_kvstoreReleaseHashtableIterator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8+v9<<(uint(int32(2))%32))))
	if v13 == int32(0) {
		F_valkey_free(m, l0)
		mBase = m.M
		v70 = m.ExcPending
		if v70 != 0 {
			return
		} else {
			return
		}
	} else {
		F_hashtableCleanupIterator(m, l0+int32(16))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			if v21&int32(2) == int32(0) {
				F_valkey_free(m, l0)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					return
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v26+v27<<(uint(int32(2))%32))))
				if v31 == int32(0) {
					F_valkey_free(m, l0)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						return
					}
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
					if v34+v35 != 0 {
						F_valkey_free(m, l0)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							return
						}
					} else {
						v37 = int32(0)
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
						v40 = v27 << (uint(int32(2)) % 32)
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v40)))
						if v42 == v37 {
							v53 = v37
							F_hashtableRelease(m, v53)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v56+v40))) = int32(0)
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v60 + int32(-1)
								F_valkey_free(m, l0)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+26)))
							if int32(0) < v45 {
								F_valkey_free(m, l0)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									return
								}
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v27<<(uint(int32(2))%32))))
								v53 = v52
								F_hashtableRelease(m, v53)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v56+v40))) = int32(0)
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
									*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v60 + int32(-1)
									F_valkey_free(m, l0)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
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
		}
	}
}
func F_kvstoreSize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2 == int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		if v8 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
			return base.I64_extend_i32_u(v11 + v12)
		} else {
			return int64(0)
		}
	} else {
		v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		return v5
	}
}
func F_kvstoreTryResizeHashtables(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if l1 < v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v9 = l1
	goto L4
L3:
	;
	v9 = v7
	goto L4
L4:
	;
	if v9 < int32(1) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = v12
	v16 = v7
	v18 = int32(0)
	goto L6
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v15<<(uint(int32(2))%32))))
	if v24 == int32(0) {
		v30 = v16
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	v31 = int32(1)
	v33 = base.I32_rem_s(v15+v31, v30)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v33
	v36 = v18 + v31
	if v36 != v9 {
		v15 = v33
		v16 = v30
		v18 = v36
		goto L6
	} else {
		goto L12
	}
L9:
	;
	v27 = F_hashtableRightsizeIfNeeded(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = v29
	goto L8
L12:
	;
	goto L7
}
