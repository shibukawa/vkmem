package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_expireCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = *(*int64)(unsafe.Add(mBase, _c_F_expireCommand[0]))
	F_expireGenericCommand(m, l0, v3, int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_expireReplicaKeys(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v112 int64
	_ = v112
	var v115 int64
	_ = v115
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v159 int64
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	v1 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_expireReplicaKeys[0]))
	if v15 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v18 == int32(0)-v20 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = F_mstime(m)
	mBase = m.M
	v24 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_expireReplicaKeys[0]))
	v28 = v25
	v30 = v24
	v31 = v24
	goto L4
L4:
	;
	v41 = F_dictGetRandomKey(m, v28)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	goto L8
L8:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v41)+8))
	goto L11
L9:
	;
	if int32(3) < v142 {
		goto L1
	} else {
		goto L33
	}
L10:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_expireReplicaKeys[0]))
	v138 = F_dictDelete(m, v137, v43)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L32
	}
L11:
	;
	if v44 == int64(0) {
		v125 = v30
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v47 = int64(0)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_expireReplicaKeys[1]))
	if v50 < int32(1) {
		v125 = v30
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v53 = v50
	v55 = v30
	v59 = v44
	v60 = v47
	v61 = v47
	goto L14
L14:
	;
	if v59&int64(1) == int64(0) {
		v103 = v53
		v104 = v55
		v105 = v61
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v105 == int64(0) {
		v125 = v104
		goto L10
	} else {
		goto L30
	}
L16:
	;
	if base.Ui64(v59) < base.Ui64(int64(2)) {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_expireReplicaKeys[2]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71+base.I32_wrap_i64(v60)<<(uint(int32(2))%32))))
	v77 = int64(0)
	v78 = int32(0)
	v79 = F_getKVStoreIndexForKey(m, v43)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	if v76 == int32(0) {
		v96 = v77
		v97 = v78
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_expireReplicaKeys[1]))
	v103 = v102
	v104 = v55 + v97
	v105 = v96 | v61
	goto L16
L20:
	;
	v83 = F_dbFindExpiresWithDictIndex(m, v76, v43, v79)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	if v83 == int32(0) {
		v96 = v77
		v97 = v78
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v87 = F_activeExpireCycleTryExpire(m, v76, v83, v23, v79)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L24
	}
L23:
	;
	v96 = int64(1) << (uint(v60) % 64)
	v97 = int32(1)
	goto L19
L24:
	;
	if v87 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_postExecutionUnitOperations(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v96 = v77
	v97 = v78
	goto L19
L27:
	;
	goto L15
L28:
	;
	v112 = int64(1)
	v115 = v60 + v112
	if v115 < base.I64_extend_i32_s(v103) {
		v53 = v103
		v55 = v104
		v59 = int64(base.Ui64(v59) >> (uint(v112) % 64))
		v60 = v115
		v61 = v105
		goto L14
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = v105
	goto L31
L31:
	;
	v142 = v104
	goto L9
L32:
	;
	v142 = v125
	goto L9
L33:
	;
	v156 = v31 + int32(1)
	if v156&int32(63) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v163 = int32(0)
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_expireReplicaKeys[0]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v164)+16))
	if v165 != v163-v167 {
		v28 = v164
		v30 = v142
		v31 = v156
		goto L4
	} else {
		goto L37
	}
L35:
	;
	v159 = F_mstime(m)
	mBase = m.M
	if int64(1) < v159-v23 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	goto L5
}
func F_expireScanCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v9&int32(1) == int32(0) {
		v20 = int64(-1)
	} else {
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l1+(v9&int32(4)^int32(12)))))
		v20 = v19
	}
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = v20 - v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = F_activeExpireCycleTryExpire(m, v23, l1, v21, l2)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return
	} else {
		if v24 == int32(0) {
			if v22 < int64(1) {
			} else {
				v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v36 + v22
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v39 + int32(1)
			}
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v43 + int32(1)
			return
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v28 + int32(1)
			F_postExecutionUnitOperations(m)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				if v22 < int64(1) {
				} else {
					v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v36 + v22
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v39 + int32(1)
				}
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v43 + int32(1)
				return
			}
		}
	}
}
func F_expireShouldSkipTableForSamplingCb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v10 == int32(255) {
		v14 = int32(0)
	} else {
		v14 = int32(1) << (uint(v10) % 32)
	}
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v17 == int32(255) {
		v21 = int32(0)
	} else {
		v21 = int32(1) << (uint(v17) % 32)
	}
	v22 = v14 + v21
	if v22 == int32(0) {
		v33 = int32(0)
	} else {
		if base.Ui64(base.I64_extend_i32_u(v4+v5)*int64(100)) < base.Ui64(base.I64_extend_i32_u(v22)) {
			v33 = int32(1)
		} else {
			v33 = int32(0)
		}
	}
	return v33
}
func F_removeExpire(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_objectGetVal(m, l1)
	mBase = m.M
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_removeExpire[0]))
	if v15 == v3 {
		v22 = v3
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v24 = F_objectGetVal(m, l1)
		mBase = m.M
		v27 = F_kvstoreHashtablePop(m, v23, v22, v24, v9+int32(8))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			if v27 == int32(0) {
				v80 = v3
				m.G0 = v9 + int32(16)
				return v80
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				v33 = F_objectSetExpire(m, v31, int64(-1))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					if v31 != v33 {
						F__serverAssert(m, int32(_a_F_removeExpire_0), int32(_a_F_removeExpire_1), int32(1880))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v36 = int32(1)
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_removeExpire[1]))
						if v38 == int32(0) {
							v80 = v36
							m.G0 = v9 + int32(16)
							return v80
						} else {
							v41 = F_objectGetVal(m, l1)
							mBase = m.M
							v42 = int32(0)
							v44 = *(*int32)(unsafe.Add(mBase, _c_F_removeExpire[0]))
							if v44 == v42 {
								v49 = v42
								v50 = F_objectGetVal(m, l1)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
								v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v56 = F_kvstoreHashtableFind(m, v53, v49, v50, v9+int32(12))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
									if v58 == int32(0) {
										v80 = v36
										m.G0 = v9 + int32(16)
										return v80
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
										if v64&int32(1) == int32(0) {
											v75 = int64(-1)
										} else {
											v74 = *(*int64)(unsafe.Add(mBase, uint32(v58+(v64&int32(4)^int32(12)))))
											v75 = v74
										}
										if v75 != int64(-1) {
											F__serverAssert(m, int32(_a_F_removeExpire_2), int32(_a_F_removeExpire_1), int32(1881))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v80 = v36
											m.G0 = v9 + int32(16)
											return v80
										}
									}
								}
							} else {
								v47 = F_getKeySlot(m, v41)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v49 = v47
									v50 = F_objectGetVal(m, l1)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v56 = F_kvstoreHashtableFind(m, v53, v49, v50, v9+int32(12))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
										if v58 == int32(0) {
											v80 = v36
											m.G0 = v9 + int32(16)
											return v80
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
											if v64&int32(1) == int32(0) {
												v75 = int64(-1)
											} else {
												v74 = *(*int64)(unsafe.Add(mBase, uint32(v58+(v64&int32(4)^int32(12)))))
												v75 = v74
											}
											if v75 != int64(-1) {
												F__serverAssert(m, int32(_a_F_removeExpire_2), int32(_a_F_removeExpire_1), int32(1881))
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int32(0)
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v80 = v36
												m.G0 = v9 + int32(16)
												return v80
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
	} else {
		v18 = F_getKeySlot(m, v11)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = v18
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v24 = F_objectGetVal(m, l1)
			mBase = m.M
			v27 = F_kvstoreHashtablePop(m, v23, v22, v24, v9+int32(8))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 == int32(0) {
					v80 = v3
					m.G0 = v9 + int32(16)
					return v80
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					v33 = F_objectSetExpire(m, v31, int64(-1))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						if v31 != v33 {
							F__serverAssert(m, int32(_a_F_removeExpire_0), int32(_a_F_removeExpire_1), int32(1880))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v36 = int32(1)
							v38 = *(*int32)(unsafe.Add(mBase, _c_F_removeExpire[1]))
							if v38 == int32(0) {
								v80 = v36
								m.G0 = v9 + int32(16)
								return v80
							} else {
								v41 = F_objectGetVal(m, l1)
								mBase = m.M
								v42 = int32(0)
								v44 = *(*int32)(unsafe.Add(mBase, _c_F_removeExpire[0]))
								if v44 == v42 {
									v49 = v42
									v50 = F_objectGetVal(m, l1)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v56 = F_kvstoreHashtableFind(m, v53, v49, v50, v9+int32(12))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
										if v58 == int32(0) {
											v80 = v36
											m.G0 = v9 + int32(16)
											return v80
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
											if v64&int32(1) == int32(0) {
												v75 = int64(-1)
											} else {
												v74 = *(*int64)(unsafe.Add(mBase, uint32(v58+(v64&int32(4)^int32(12)))))
												v75 = v74
											}
											if v75 != int64(-1) {
												F__serverAssert(m, int32(_a_F_removeExpire_2), int32(_a_F_removeExpire_1), int32(1881))
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int32(0)
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v80 = v36
												m.G0 = v9 + int32(16)
												return v80
											}
										}
									}
								} else {
									v47 = F_getKeySlot(m, v41)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										v49 = v47
										v50 = F_objectGetVal(m, l1)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
										v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v56 = F_kvstoreHashtableFind(m, v53, v49, v50, v9+int32(12))
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
											if v58 == int32(0) {
												v80 = v36
												m.G0 = v9 + int32(16)
												return v80
											} else {
												v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
												if v64&int32(1) == int32(0) {
													v75 = int64(-1)
												} else {
													v74 = *(*int64)(unsafe.Add(mBase, uint32(v58+(v64&int32(4)^int32(12)))))
													v75 = v74
												}
												if v75 != int64(-1) {
													F__serverAssert(m, int32(_a_F_removeExpire_2), int32(_a_F_removeExpire_1), int32(1881))
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return int32(0)
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v80 = v36
													m.G0 = v9 + int32(16)
													return v80
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
}
