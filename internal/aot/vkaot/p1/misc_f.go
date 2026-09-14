package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F___fe_raise_inexact(m *base.Module) int32 {
	return int32(0)
}
func F___fixtfdi(m *base.Module, l0 int64, l1 int64) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v48 int64
	_ = v48
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v57 int64
	_ = v57
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v15 = base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(48))%64))) & int32(32767)
	if base.Ui32(v15) < base.Ui32(int32(16383)) {
		v63 = int64(0)
	} else {
		if base.Ui32(int32(-65)) < base.Ui32(v15+int32(-16447)) {
			v29 = l1&int64(281474976710655) | int64(281474976710656)
			v31 = int32(16495) - v15
			if v31&int32(64) == int32(0) {
				if v31 == int32(0) {
					v52 = l0
					v53 = v29
				} else {
					v48 = base.I64_extend_i32_u(v31)
					v52 = v29<<(uint(base.I64_extend_i32_u(int32(64)-v31))%64) | int64(base.Ui64(l0)>>(uint(v48)%64))
					v53 = int64(base.Ui64(v29) >> (uint(v48) % 64))
				}
			} else {
				v52 = int64(base.Ui64(v29) >> (uint(base.I64_extend_i32_u(v31+int32(-64))) % 64))
				v53 = int64(0)
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = v52
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v53
			v57 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
			if int64(-1) < l1 {
				v62 = v57
			} else {
				v62 = int64(0) - v57
			}
			v63 = v62
		} else {
			v63 = l1>>(uint(int64(63))%64) ^ int64(9223372036854775807)
		}
	}
	m.G0 = v8 + int32(16)
	return v63
}
func F___fpclassify(m *base.Module, l0 float64) int32 {
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	v5 = base.I64_reinterpret_f64(l0)
	v9 = int32(2047)
	v10 = base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(52))%64))) & v9
	if v10 == v9 {
		v24 = base.B2i32(v5&int64(4503599627370495) == int64(0))
		return v24
	} else {
		if v10 != 0 {
			v24 = int32(4)
			return v24
		} else {
			if base.F64_eq(l0, float64(0)) != 0 {
				v18 = int32(2)
			} else {
				v18 = int32(3)
			}
			return v18
		}
	}
}
func F___fpclassifyl(m *base.Module, l0 int64, l1 int64) int32 {
	var v7 int64
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v7 = l1 & int64(281474976710655)
	v11 = int32(32767)
	v12 = base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(48))%64))) & v11
	if v12 == v11 {
		v26 = base.B2i32(v7|l0 == int64(0))
		return v26
	} else {
		if v12 != 0 {
			v26 = int32(4)
			return v26
		} else {
			if v7|l0 == int64(0) {
				v21 = int32(2)
			} else {
				v21 = int32(3)
			}
			return v21
		}
	}
}
func F___fstat(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	if int32(-1) < l0 {
		v17 = F___fstatat(m, l0, int32(_a320), l1, int32(4096))
		mBase = m.M
		return v17
	} else {
		v8 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(8)
		return int32(-1)
	}
}
func F___fwritex(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v7 != 0 {
		v36 = v7
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v99
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if base.Ui32(l1) <= base.Ui32(v36-v38) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v8 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v10 + int32(-1) | v10
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v15&int32(8) == v8 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v34 != 0 {
		v99 = v8
		goto L1
	} else {
		goto L7
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v26
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v26 + v29
	v34 = int32(0)
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v15 | int32(32)
	v34 = int32(-1)
	goto L4
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v36 = v35
	goto L2
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	if v47 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v42 = m.T0[v41].(func(*base.Module, int32, int32, int32) int32)(m, l2, l0, l1)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	return v42
L12:
	;
	if v82 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L13:
	;
	v82 = l1
	v84 = int32(0)
	v85 = v38
	v86 = l0
	goto L12
L14:
	;
	if l1 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v55 = l1
	goto L17
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v69 = m.T0[v68].(func(*base.Module, int32, int32, int32) int32)(m, l2, l0, v55)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L10
	} else {
		goto L21
	}
L17:
	;
	v58 = l0 + v55
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+int32(-1)))))
	if v61 == int32(10) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v65 = v55 + int32(-1)
	if v65 == int32(0) {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v55 = v65
	goto L17
L21:
	;
	if base.Ui32(v69) < base.Ui32(v55) {
		v99 = v69
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v82 = l1 - v55
	v84 = v55
	v85 = v73
	v86 = v58
	goto L12
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v91 + v82
	v99 = v84 + v82
	goto L1
L24:
	;
	goto L23
L25:
	;
	v89 = F__emscripten_memcpy_bulkmem(m, v85, v86, v82)
	mBase = m.M
	goto L24
}
func F_f_luaopen(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = F_luaM_realloc_(m, l0, v3, v3, int32(192))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v11 + int32(168)
		v20 = int32(0)
		v23 = F_luaM_realloc_(m, l0, v20, v20, int32(720))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(45)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v23
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v23 + int32(624)
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v23
			v34 = v23 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34
			v36 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v36
			*(*int32)(unsafe.Add(mBase, uint32(v31))) = v34
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v34
			*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v23 + int32(336)
			v45 = F_luaH_new(m, l0, v36, int32(2))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v45
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v53 = F_luaH_new(m, l0, int32(0), int32(2))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v50)+104)) = int32(5)
					*(*int32)(unsafe.Add(mBase, uint32(v50)+96)) = v53
					F_luaS_resize(m, l0, int32(32))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						F_luaT_init(m, l0)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							F_luaX_init(m, l0)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								v65 = m.G3
								v69 = F_luaS_newlstr(m, l0, v65+int32(_a2661), int32(17))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+5)))
									v73 = v71 | int32(32)
									*(*uint8)(unsafe.Add(mBase, uint32(v69)+5)) = uint8(v73)
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v7)+68))
									*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v75 << (uint(int32(2)) % 32)
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
func F_fcallCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_fcallCommandGeneric(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_fcallCommandGeneric(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[336]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_replicationFeedMonitors(m, l0, v12, v14, v15, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
		if v19 != 0 {
			v31 = v19
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
			v39 = F_getLongLongFromObject(m, v36, v9+int32(40))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				if v39 == int32(0) {
					v46 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v46 <= base.I64_extend_i32_s(v47+int32(-3)) {
						if int64(-1) < v46 {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
							v62 = F_objectGetVal(m, v61)
							mBase = m.M
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
							v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)+24))
							v65 = F_scriptPrepareForRun(m, v9, v34, l0, v62, v64, l1)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								if v65 != 0 {
									m.G0 = v9 + int32(48)
									return
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
									v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v72 = v70 + int32(12)
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									F_scriptingEngineCallFunction(m, v34, v9, v67, v68, int32(1), v72, v73, v72+v73<<(uint(int32(2))%32), v77-v73+int32(-3))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										F_scriptResetRun(m, v9)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return
										} else {
											m.G0 = v9 + int32(48)
											return
										}
									}
								}
							}
						} else {
							F_addReplyError(m, l0, int32(_a610))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								m.G0 = v9 + int32(48)
								return
							}
						}
					} else {
						F_addReplyError(m, l0, int32(_a611))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							m.G0 = v9 + int32(48)
							return
						}
					}
				} else {
					F_addReplyError(m, l0, int32(_a624))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						m.G0 = v9 + int32(48)
						return
					}
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[332]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
			v25 = F_objectGetVal(m, v24)
			mBase = m.M
			v26 = F_dictFind(m, v22, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				if v26 != 0 {
					v31 = v26
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
					v39 = F_getLongLongFromObject(m, v36, v9+int32(40))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						if v39 == int32(0) {
							v46 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v46 <= base.I64_extend_i32_s(v47+int32(-3)) {
								if int64(-1) < v46 {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
									v62 = F_objectGetVal(m, v61)
									mBase = m.M
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
									v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)+24))
									v65 = F_scriptPrepareForRun(m, v9, v34, l0, v62, v64, l1)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										if v65 != 0 {
											m.G0 = v9 + int32(48)
											return
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
											v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v72 = v70 + int32(12)
											v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											F_scriptingEngineCallFunction(m, v34, v9, v67, v68, int32(1), v72, v73, v72+v73<<(uint(int32(2))%32), v77-v73+int32(-3))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												F_scriptResetRun(m, v9)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return
												} else {
													m.G0 = v9 + int32(48)
													return
												}
											}
										}
									}
								} else {
									F_addReplyError(m, l0, int32(_a610))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										m.G0 = v9 + int32(48)
										return
									}
								}
							} else {
								F_addReplyError(m, l0, int32(_a611))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									m.G0 = v9 + int32(48)
									return
								}
							}
						} else {
							F_addReplyError(m, l0, int32(_a624))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								m.G0 = v9 + int32(48)
								return
							}
						}
					}
				} else {
					F_addReplyError(m, l0, int32(_a625))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						m.G0 = v9 + int32(48)
						return
					}
				}
			}
		}
	}
}
func F_fchmod(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = m.Env.X__syscall_fchmod(m, l0, l1)
	mBase = m.M
	if v9 != int32(-8) {
		if base.Ui32(v9) < base.Ui32(int32(-4095)) {
			v20 = v9
		} else {
			v15 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(0) - v9
			v20 = int32(-1)
		}
		v92 = v20
	} else {
		v12 = F___wasi_fd_is_valid(m, l0)
		mBase = m.M
		if v12 != 0 {
			v27 = int32(0)
			for {
				v30 = v7 + v27
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1008]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v33)
				if v27 != int32(14) {
					v27 = v27 + int32(1)
					continue
				} else {
					break
				}
				break
			}
			if l0 == int32(0) {
				v74 = int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v74)
				v76 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v76)
			} else {
				v44 = int32(14)
				v46 = l0
				for {
					v48 = v44 + int32(1)
					v52 = base.I32_div_u_s(v46, int32(10))
					if base.Ui32(int32(9)) < base.Ui32(v46) {
						v44 = v48
						v46 = v52
						continue
					} else {
						break
					}
					break
				}
				v54 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v7+v48))) = uint8(v54)
				v57 = l0
				v58 = v48
				for {
					v62 = v58 + int32(-1)
					v64 = int32(10)
					v65 = base.I32_div_u_s(v57, v64)
					v70 = v57 - v65*v64 | int32(48)
					*(*uint8)(unsafe.Add(mBase, uint32(v7+v62))) = uint8(v70)
					if base.Ui32(int32(9)) < base.Ui32(v57) {
						v57 = v65
						v58 = v62
						continue
					} else {
						break
					}
					break
				}
			}
			v83 = m.Env.X__syscall_chmod(m, v7, l1)
			mBase = m.M
			if base.Ui32(v83) < base.Ui32(int32(-4095)) {
				v91 = v83
			} else {
				v86 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(0) - v83
				v91 = int32(-1)
			}
			v92 = v91
		} else {
			if base.Ui32(v9) < base.Ui32(int32(-4095)) {
				v20 = v9
			} else {
				v15 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(0) - v9
				v20 = int32(-1)
			}
			v92 = v20
		}
	}
	m.G0 = v7 + int32(32)
	return v92
}
func F_ffc_digit_comp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v72 int32
	_ = v72
	var v76 int64
	_ = v76
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v138 int32
	_ = v138
	var v142 int64
	_ = v142
	var v150 int32
	_ = v150
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v204 int32
	_ = v204
	var v208 int64
	_ = v208
	var v216 int32
	_ = v216
	var v217 int64
	_ = v217
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v258 int32
	_ = v258
	var v275 int64
	_ = v275
	var v279 int32
	_ = v279
	var v295 int32
	_ = v295
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v333 int32
	_ = v333
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v370 int64
	_ = v370
	var v371 int64
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v386 int64
	_ = v386
	var v398 int64
	_ = v398
	var v400 int64
	_ = v400
	var v405 int64
	_ = v405
	var v408 int64
	_ = v408
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v427 int64
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int64
	_ = v516
	var v525 int32
	_ = v525
	var v535 int64
	_ = v535
	var v537 int32
	_ = v537
	var v547 int32
	_ = v547
	var v560 int32
	_ = v560
	var v561 int64
	_ = v561
	var v563 int64
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int64
	_ = v567
	var v569 int64
	_ = v569
	var v571 int64
	_ = v571
	var v574 int32
	_ = v574
	var v575 int64
	_ = v575
	var v579 int64
	_ = v579
	var v582 int32
	_ = v582
	var v583 int64
	_ = v583
	var v587 int64
	_ = v587
	var v590 int64
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v605 int64
	_ = v605
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v624 int64
	_ = v624
	var v639 int32
	_ = v639
	var v649 int32
	_ = v649
	var v650 int64
	_ = v650
	var v652 int64
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int64
	_ = v656
	var v658 int64
	_ = v658
	var v660 int64
	_ = v660
	var v663 int32
	_ = v663
	var v664 int64
	_ = v664
	var v668 int64
	_ = v668
	var v671 int32
	_ = v671
	var v672 int64
	_ = v672
	var v676 int64
	_ = v676
	var v679 int64
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v693 int64
	_ = v693
	var v709 int64
	_ = v709
	var v717 int32
	_ = v717
	var v724 int64
	_ = v724
	var v736 int32
	_ = v736
	var v749 int32
	_ = v749
	var v750 int64
	_ = v750
	var v752 int64
	_ = v752
	var v754 int32
	_ = v754
	var v757 int64
	_ = v757
	var v759 int32
	_ = v759
	var v769 int64
	_ = v769
	var v785 int64
	_ = v785
	var v801 int32
	_ = v801
	var v811 int64
	_ = v811
	var v812 int64
	_ = v812
	var v814 int32
	_ = v814
	var v820 int32
	_ = v820
	var v827 int64
	_ = v827
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v905 int32
	_ = v905
	var v925 int32
	_ = v925
	var v938 int32
	_ = v938
	var v944 int32
	_ = v944
	var v965 int64
	_ = v965
	var v966 int64
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v981 int64
	_ = v981
	var v988 int32
	_ = v988
	var v1007 int32
	_ = v1007
	var v1020 int64
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1044 int32
	_ = v1044
	var v1076 int32
	_ = v1076
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1111 int32
	_ = v1111
	var v1125 int32
	_ = v1125
	var v1134 int64
	_ = v1134
	var v1135 int64
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1142 int32
	_ = v1142
	var v1150 int64
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1162 int64
	_ = v1162
	var v1164 int64
	_ = v1164
	var v1169 int64
	_ = v1169
	var v1172 int64
	_ = v1172
	var v1184 int32
	_ = v1184
	var v1190 int32
	_ = v1190
	var v1191 int64
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1216 int32
	_ = v1216
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1252 int32
	_ = v1252
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1280 int64
	_ = v1280
	var v1289 int32
	_ = v1289
	var v1299 int64
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1309 int32
	_ = v1309
	var v1324 int32
	_ = v1324
	var v1325 int64
	_ = v1325
	var v1327 int64
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int64
	_ = v1331
	var v1333 int64
	_ = v1333
	var v1335 int64
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1339 int64
	_ = v1339
	var v1343 int64
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1347 int64
	_ = v1347
	var v1351 int64
	_ = v1351
	var v1354 int64
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1369 int64
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1388 int64
	_ = v1388
	var v1393 int32
	_ = v1393
	var v1413 int32
	_ = v1413
	var v1414 int64
	_ = v1414
	var v1416 int64
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1420 int64
	_ = v1420
	var v1422 int64
	_ = v1422
	var v1424 int64
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1428 int64
	_ = v1428
	var v1432 int64
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1436 int64
	_ = v1436
	var v1440 int64
	_ = v1440
	var v1443 int64
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1457 int64
	_ = v1457
	var v1473 int64
	_ = v1473
	var v1481 int32
	_ = v1481
	var v1488 int64
	_ = v1488
	var v1498 int32
	_ = v1498
	var v1513 int32
	_ = v1513
	var v1514 int64
	_ = v1514
	var v1516 int64
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1521 int64
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1533 int64
	_ = v1533
	var v1549 int64
	_ = v1549
	var v1565 int32
	_ = v1565
	var v1575 int64
	_ = v1575
	var v1576 int64
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1591 int64
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1599 int32
	_ = v1599
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1670 int32
	_ = v1670
	var v1689 int32
	_ = v1689
	var v1702 int32
	_ = v1702
	var v1716 int32
	_ = v1716
	var v1729 int64
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1745 int64
	_ = v1745
	var v1760 int64
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1764 int32
	_ = v1764
	var v1785 int32
	_ = v1785
	var v1786 int64
	_ = v1786
	var v1788 int64
	_ = v1788
	var v1790 int32
	_ = v1790
	var v1793 int64
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1805 int64
	_ = v1805
	var v1821 int64
	_ = v1821
	var v1837 int32
	_ = v1837
	var v1847 int64
	_ = v1847
	var v1848 int64
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1870 int32
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1877 int32
	_ = v1877
	var v1889 int32
	_ = v1889
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1941 int32
	_ = v1941
	var v1960 int32
	_ = v1960
	var v1973 int32
	_ = v1973
	var v2010 int32
	_ = v2010
	var v2023 int64
	_ = v2023
	var v2027 int32
	_ = v2027
	var v2047 int32
	_ = v2047
	var v2077 int32
	_ = v2077
	var v2090 int32
	_ = v2090
	var v2094 int32
	_ = v2094
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2131 int32
	_ = v2131
	var v2140 int32
	_ = v2140
	var v2151 int64
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2158 int32
	_ = v2158
	var v2175 int32
	_ = v2175
	var v2176 int64
	_ = v2176
	var v2177 int64
	_ = v2177
	var v2179 int64
	_ = v2179
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2183 int64
	_ = v2183
	var v2186 int64
	_ = v2186
	var v2188 int64
	_ = v2188
	var v2191 int32
	_ = v2191
	var v2192 int64
	_ = v2192
	var v2197 int64
	_ = v2197
	var v2200 int32
	_ = v2200
	var v2201 int64
	_ = v2201
	var v2206 int64
	_ = v2206
	var v2209 int64
	_ = v2209
	var v2211 int32
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2223 int64
	_ = v2223
	var v2224 int64
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2255 int64
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2265 int32
	_ = v2265
	var v2279 int32
	_ = v2279
	var v2280 int64
	_ = v2280
	var v2283 int64
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2288 int64
	_ = v2288
	var v2290 int32
	_ = v2290
	var v2300 int64
	_ = v2300
	var v2301 int64
	_ = v2301
	var v2325 int32
	_ = v2325
	var v2333 int32
	_ = v2333
	var v2343 int64
	_ = v2343
	var v2344 int64
	_ = v2344
	var v2346 int32
	_ = v2346
	var v2365 int32
	_ = v2365
	var v2367 int32
	_ = v2367
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2418 int32
	_ = v2418
	var v2421 int32
	_ = v2421
	var v2425 int32
	_ = v2425
	var v2436 int64
	_ = v2436
	var v2438 int32
	_ = v2438
	var v2452 int64
	_ = v2452
	var v2467 int64
	_ = v2467
	var v2469 int32
	_ = v2469
	var v2478 int32
	_ = v2478
	var v2492 int32
	_ = v2492
	var v2493 int64
	_ = v2493
	var v2495 int64
	_ = v2495
	var v2497 int32
	_ = v2497
	var v2500 int64
	_ = v2500
	var v2502 int32
	_ = v2502
	var v2512 int64
	_ = v2512
	var v2528 int64
	_ = v2528
	var v2544 int32
	_ = v2544
	var v2554 int64
	_ = v2554
	var v2555 int64
	_ = v2555
	var v2557 int32
	_ = v2557
	var v2577 int32
	_ = v2577
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2584 int32
	_ = v2584
	var v2596 int32
	_ = v2596
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2647 int32
	_ = v2647
	var v2667 int32
	_ = v2667
	var v2680 int32
	_ = v2680
	var v2713 int32
	_ = v2713
	var v2730 int64
	_ = v2730
	var v2735 int32
	_ = v2735
	var v2751 int32
	_ = v2751
	var v2781 int32
	_ = v2781
	var v2798 int32
	_ = v2798
	var v2800 int32
	_ = v2800
	var v2802 int32
	_ = v2802
	var v2814 int32
	_ = v2814
	var v2846 int32
	_ = v2846
	var v2865 int32
	_ = v2865
	var v2884 int32
	_ = v2884
	var v2897 int64
	_ = v2897
	var v2901 int32
	_ = v2901
	var v2921 int32
	_ = v2921
	var v2951 int32
	_ = v2951
	var v2964 int32
	_ = v2964
	var v2968 int32
	_ = v2968
	var v3028 int32
	_ = v3028
	var v3030 int32
	_ = v3030
	var v3034 int32
	_ = v3034
	var v3043 int32
	_ = v3043
	var v3054 int64
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3061 int32
	_ = v3061
	var v3078 int32
	_ = v3078
	var v3079 int64
	_ = v3079
	var v3080 int64
	_ = v3080
	var v3082 int64
	_ = v3082
	var v3084 int32
	_ = v3084
	var v3085 int32
	_ = v3085
	var v3086 int64
	_ = v3086
	var v3089 int64
	_ = v3089
	var v3091 int64
	_ = v3091
	var v3094 int32
	_ = v3094
	var v3095 int64
	_ = v3095
	var v3100 int64
	_ = v3100
	var v3103 int32
	_ = v3103
	var v3104 int64
	_ = v3104
	var v3109 int64
	_ = v3109
	var v3112 int64
	_ = v3112
	var v3114 int32
	_ = v3114
	var v3116 int32
	_ = v3116
	var v3126 int64
	_ = v3126
	var v3127 int64
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3158 int64
	_ = v3158
	var v3159 int32
	_ = v3159
	var v3168 int32
	_ = v3168
	var v3182 int32
	_ = v3182
	var v3183 int64
	_ = v3183
	var v3186 int64
	_ = v3186
	var v3188 int32
	_ = v3188
	var v3191 int64
	_ = v3191
	var v3193 int32
	_ = v3193
	var v3203 int64
	_ = v3203
	var v3204 int64
	_ = v3204
	var v3228 int32
	_ = v3228
	var v3236 int32
	_ = v3236
	var v3246 int64
	_ = v3246
	var v3247 int64
	_ = v3247
	var v3249 int32
	_ = v3249
	var v3268 int32
	_ = v3268
	var v3270 int32
	_ = v3270
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3305 int32
	_ = v3305
	var v3306 int32
	_ = v3306
	var v3321 int32
	_ = v3321
	var v3324 int32
	_ = v3324
	var v3328 int32
	_ = v3328
	var v3339 int64
	_ = v3339
	var v3340 int64
	_ = v3340
	var v3341 int32
	_ = v3341
	var v3342 int32
	_ = v3342
	var v3362 int32
	_ = v3362
	var v3381 int32
	_ = v3381
	var v3400 int64
	_ = v3400
	var v3406 int32
	_ = v3406
	var v3410 int32
	_ = v3410
	var v3427 int32
	_ = v3427
	var v3442 int32
	_ = v3442
	var v3453 int64
	_ = v3453
	var v3454 int64
	_ = v3454
	var v3459 int32
	_ = v3459
	var v3468 int32
	_ = v3468
	var v3475 int32
	_ = v3475
	var v3478 int32
	_ = v3478
	var v3480 int32
	_ = v3480
	var v3487 int32
	_ = v3487
	var v3497 int64
	_ = v3497
	var v3499 int32
	_ = v3499
	var v3508 int32
	_ = v3508
	var v3522 int32
	_ = v3522
	var v3523 int64
	_ = v3523
	var v3524 int64
	_ = v3524
	var v3526 int64
	_ = v3526
	var v3528 int32
	_ = v3528
	var v3529 int32
	_ = v3529
	var v3530 int64
	_ = v3530
	var v3533 int64
	_ = v3533
	var v3535 int64
	_ = v3535
	var v3538 int32
	_ = v3538
	var v3539 int64
	_ = v3539
	var v3544 int64
	_ = v3544
	var v3547 int32
	_ = v3547
	var v3548 int64
	_ = v3548
	var v3553 int64
	_ = v3553
	var v3556 int64
	_ = v3556
	var v3558 int32
	_ = v3558
	var v3560 int32
	_ = v3560
	var v3570 int64
	_ = v3570
	var v3571 int64
	_ = v3571
	var v3572 int32
	_ = v3572
	var v3601 int64
	_ = v3601
	var v3603 int32
	_ = v3603
	var v3605 int32
	_ = v3605
	var v3626 int32
	_ = v3626
	var v3627 int64
	_ = v3627
	var v3630 int64
	_ = v3630
	var v3632 int32
	_ = v3632
	var v3635 int64
	_ = v3635
	var v3637 int32
	_ = v3637
	var v3647 int64
	_ = v3647
	var v3648 int64
	_ = v3648
	var v3689 int64
	_ = v3689
	var v3690 int64
	_ = v3690
	var v3704 int32
	_ = v3704
	var v3711 int32
	_ = v3711
	var v3722 int64
	_ = v3722
	var v3728 int32
	_ = v3728
	var v3737 int32
	_ = v3737
	var v3747 int32
	_ = v3747
	var v3750 int32
	_ = v3750
	var v3751 int32
	_ = v3751
	var v3756 int64
	_ = v3756
	var v3765 int32
	_ = v3765
	var v3777 int32
	_ = v3777
	var v3786 int32
	_ = v3786
	var v3791 int64
	_ = v3791
	var v3800 int32
	_ = v3800
	var v3801 int64
	_ = v3801
	var v3803 int64
	_ = v3803
	var v3805 int32
	_ = v3805
	var v3806 int32
	_ = v3806
	var v3807 int64
	_ = v3807
	var v3809 int64
	_ = v3809
	var v3811 int64
	_ = v3811
	var v3814 int32
	_ = v3814
	var v3815 int64
	_ = v3815
	var v3819 int64
	_ = v3819
	var v3822 int32
	_ = v3822
	var v3823 int64
	_ = v3823
	var v3827 int64
	_ = v3827
	var v3830 int64
	_ = v3830
	var v3832 int32
	_ = v3832
	var v3834 int32
	_ = v3834
	var v3837 int64
	_ = v3837
	var v3840 int32
	_ = v3840
	var v3842 int32
	_ = v3842
	var v3845 int32
	_ = v3845
	var v3848 int32
	_ = v3848
	var v3849 int32
	_ = v3849
	var v3852 int32
	_ = v3852
	var v3857 int64
	_ = v3857
	var v3865 int64
	_ = v3865
	var v3868 int64
	_ = v3868
	var v3871 int32
	_ = v3871
	var v3872 int64
	_ = v3872
	var v3878 int32
	_ = v3878
	var v3881 int32
	_ = v3881
	var v3882 int32
	_ = v3882
	var v3883 int64
	_ = v3883
	var v3884 int32
	_ = v3884
	var v3885 int32
	_ = v3885
	var v3887 int64
	_ = v3887
	var v3892 int32
	_ = v3892
	var v3895 int64
	_ = v3895
	var v3898 int64
	_ = v3898
	var v3901 int64
	_ = v3901
	var v3910 int64
	_ = v3910
	var v3921 int64
	_ = v3921
	var v3922 int64
	_ = v3922
	var v3932 int32
	_ = v3932
	var v3933 int64
	_ = v3933
	var v3934 int64
	_ = v3934
	var v3935 int32
	_ = v3935
	var v3937 int64
	_ = v3937
	var v3943 int32
	_ = v3943
	var v3946 int32
	_ = v3946
	var v3954 int32
	_ = v3954
	var v3956 int32
	_ = v3956
	var v3973 int32
	_ = v3973
	var v3992 int64
	_ = v3992
	var v3998 int32
	_ = v3998
	var v4002 int32
	_ = v4002
	var v4019 int32
	_ = v4019
	var v4036 int32
	_ = v4036
	var v4046 int64
	_ = v4046
	var v4051 int32
	_ = v4051
	var v4060 int32
	_ = v4060
	var v4067 int32
	_ = v4067
	var v4070 int32
	_ = v4070
	var v4072 int32
	_ = v4072
	var v4079 int32
	_ = v4079
	var v4089 int64
	_ = v4089
	var v4091 int32
	_ = v4091
	var v4100 int32
	_ = v4100
	var v4114 int32
	_ = v4114
	var v4115 int64
	_ = v4115
	var v4116 int64
	_ = v4116
	var v4118 int64
	_ = v4118
	var v4120 int32
	_ = v4120
	var v4121 int32
	_ = v4121
	var v4122 int64
	_ = v4122
	var v4125 int64
	_ = v4125
	var v4127 int64
	_ = v4127
	var v4130 int32
	_ = v4130
	var v4131 int64
	_ = v4131
	var v4136 int64
	_ = v4136
	var v4139 int32
	_ = v4139
	var v4140 int64
	_ = v4140
	var v4145 int64
	_ = v4145
	var v4148 int64
	_ = v4148
	var v4150 int32
	_ = v4150
	var v4152 int32
	_ = v4152
	var v4162 int64
	_ = v4162
	var v4163 int64
	_ = v4163
	var v4164 int32
	_ = v4164
	var v4193 int64
	_ = v4193
	var v4195 int32
	_ = v4195
	var v4197 int32
	_ = v4197
	var v4218 int32
	_ = v4218
	var v4219 int64
	_ = v4219
	var v4222 int64
	_ = v4222
	var v4224 int32
	_ = v4224
	var v4227 int64
	_ = v4227
	var v4229 int32
	_ = v4229
	var v4239 int64
	_ = v4239
	var v4240 int64
	_ = v4240
	var v4282 int64
	_ = v4282
	var v4296 int32
	_ = v4296
	var v4303 int32
	_ = v4303
	var v4321 int32
	_ = v4321
	var v4338 int32
	_ = v4338
	var v4341 int32
	_ = v4341
	var v4342 int32
	_ = v4342
	var v4347 int64
	_ = v4347
	var v4356 int32
	_ = v4356
	var v4366 int64
	_ = v4366
	var v4368 int32
	_ = v4368
	var v4377 int32
	_ = v4377
	var v4391 int32
	_ = v4391
	var v4392 int64
	_ = v4392
	var v4394 int64
	_ = v4394
	var v4396 int32
	_ = v4396
	var v4397 int32
	_ = v4397
	var v4398 int64
	_ = v4398
	var v4400 int64
	_ = v4400
	var v4402 int64
	_ = v4402
	var v4405 int32
	_ = v4405
	var v4406 int64
	_ = v4406
	var v4410 int64
	_ = v4410
	var v4413 int32
	_ = v4413
	var v4414 int64
	_ = v4414
	var v4418 int64
	_ = v4418
	var v4421 int64
	_ = v4421
	var v4423 int32
	_ = v4423
	var v4425 int32
	_ = v4425
	var v4435 int64
	_ = v4435
	var v4437 int32
	_ = v4437
	var v4451 int64
	_ = v4451
	var v4466 int64
	_ = v4466
	var v4468 int32
	_ = v4468
	var v4470 int32
	_ = v4470
	var v4491 int32
	_ = v4491
	var v4492 int64
	_ = v4492
	var v4494 int64
	_ = v4494
	var v4496 int32
	_ = v4496
	var v4499 int64
	_ = v4499
	var v4501 int32
	_ = v4501
	var v4511 int64
	_ = v4511
	var v4527 int64
	_ = v4527
	var v4567 int32
	_ = v4567
	var v4606 int32
	_ = v4606
	var v4609 int32
	_ = v4609
	var v4613 int32
	_ = v4613
	var v4614 int32
	_ = v4614
	var v4618 int32
	_ = v4618
	var v4622 int32
	_ = v4622
	var v4635 int32
	_ = v4635
	var v4643 int32
	_ = v4643
	var v4644 int32
	_ = v4644
	var v4661 int32
	_ = v4661
	var v4664 int32
	_ = v4664
	var v4666 int32
	_ = v4666
	var v4669 int32
	_ = v4669
	var v4671 int32
	_ = v4671
	var v4683 int32
	_ = v4683
	var v4688 int32
	_ = v4688
	var v4693 int32
	_ = v4693
	var v4698 int32
	_ = v4698
	var v4699 int32
	_ = v4699
	var v4701 int32
	_ = v4701
	var v4705 int32
	_ = v4705
	var v4707 int32
	_ = v4707
	var v4709 int32
	_ = v4709
	var v4711 int32
	_ = v4711
	var v4715 int32
	_ = v4715
	var v4716 int32
	_ = v4716
	var v4720 int32
	_ = v4720
	var v4724 int32
	_ = v4724
	var v4737 int32
	_ = v4737
	var v4745 int32
	_ = v4745
	var v4746 int32
	_ = v4746
	var v4763 int32
	_ = v4763
	var v4766 int32
	_ = v4766
	var v4768 int32
	_ = v4768
	var v4771 int32
	_ = v4771
	var v4773 int32
	_ = v4773
	var v4785 int32
	_ = v4785
	var v4790 int32
	_ = v4790
	var v4795 int32
	_ = v4795
	var v4800 int32
	_ = v4800
	var v4801 int32
	_ = v4801
	var v4803 int32
	_ = v4803
	var v4815 int32
	_ = v4815
	var v4824 int32
	_ = v4824
	var v4843 int32
	_ = v4843
	var v4848 int32
	_ = v4848
	var v4854 int32
	_ = v4854
	var v4859 int32
	_ = v4859
	var v4860 int32
	_ = v4860
	var v4872 int32
	_ = v4872
	var v4884 int32
	_ = v4884
	var v4906 int32
	_ = v4906
	var v4910 int32
	_ = v4910
	var v4915 int32
	_ = v4915
	var v4916 int32
	_ = v4916
	var v4917 int32
	_ = v4917
	var v4918 int32
	_ = v4918
	var v4922 int32
	_ = v4922
	var v4926 int32
	_ = v4926
	var v4933 int32
	_ = v4933
	var v4936 int32
	_ = v4936
	var v4943 int32
	_ = v4943
	var v4944 int32
	_ = v4944
	var v4945 int32
	_ = v4945
	var v4949 int32
	_ = v4949
	var v4951 int32
	_ = v4951
	var v4952 int32
	_ = v4952
	var v4954 int32
	_ = v4954
	var v4956 int32
	_ = v4956
	var v4967 int32
	_ = v4967
	var v4973 int32
	_ = v4973
	var v4974 int32
	_ = v4974
	var v4976 int32
	_ = v4976
	var v4982 int32
	_ = v4982
	var v4989 int32
	_ = v4989
	var v4993 int32
	_ = v4993
	var v4996 int32
	_ = v4996
	var v5002 int32
	_ = v5002
	var v5009 int32
	_ = v5009
	var v5013 int32
	_ = v5013
	var v5016 int32
	_ = v5016
	var v5019 int32
	_ = v5019
	var v5020 int32
	_ = v5020
	var v5021 int32
	_ = v5021
	var v5026 int32
	_ = v5026
	var v5027 int32
	_ = v5027
	var v5028 int32
	_ = v5028
	var v5030 int32
	_ = v5030
	var v5032 int32
	_ = v5032
	var v5033 int32
	_ = v5033
	var v5035 int32
	_ = v5035
	var v5037 int32
	_ = v5037
	var v5041 int32
	_ = v5041
	var v5042 int32
	_ = v5042
	var v5043 int32
	_ = v5043
	var v5048 int32
	_ = v5048
	var v5049 int32
	_ = v5049
	var v5050 int32
	_ = v5050
	var v5052 int32
	_ = v5052
	var v5054 int32
	_ = v5054
	var v5059 int32
	_ = v5059
	var v5075 int32
	_ = v5075
	var v5076 int32
	_ = v5076
	var v5077 int32
	_ = v5077
	var v5089 int32
	_ = v5089
	var v5098 int32
	_ = v5098
	var v5117 int32
	_ = v5117
	var v5122 int32
	_ = v5122
	var v5128 int32
	_ = v5128
	var v5133 int32
	_ = v5133
	var v5134 int32
	_ = v5134
	var v5146 int32
	_ = v5146
	var v5179 int32
	_ = v5179
	var v5183 int32
	_ = v5183
	var v5188 int32
	_ = v5188
	var v5189 int32
	_ = v5189
	var v5190 int32
	_ = v5190
	var v5191 int32
	_ = v5191
	var v5195 int32
	_ = v5195
	var v5199 int32
	_ = v5199
	var v5206 int32
	_ = v5206
	var v5209 int32
	_ = v5209
	var v5216 int32
	_ = v5216
	var v5217 int32
	_ = v5217
	var v5218 int32
	_ = v5218
	var v5222 int32
	_ = v5222
	var v5224 int32
	_ = v5224
	var v5225 int32
	_ = v5225
	var v5227 int32
	_ = v5227
	var v5229 int32
	_ = v5229
	var v5240 int32
	_ = v5240
	var v5246 int32
	_ = v5246
	var v5247 int32
	_ = v5247
	var v5249 int32
	_ = v5249
	var v5255 int32
	_ = v5255
	var v5262 int32
	_ = v5262
	var v5266 int32
	_ = v5266
	var v5269 int32
	_ = v5269
	var v5275 int32
	_ = v5275
	var v5282 int32
	_ = v5282
	var v5286 int32
	_ = v5286
	var v5289 int32
	_ = v5289
	var v5292 int32
	_ = v5292
	var v5293 int32
	_ = v5293
	var v5294 int32
	_ = v5294
	var v5299 int32
	_ = v5299
	var v5300 int32
	_ = v5300
	var v5301 int32
	_ = v5301
	var v5303 int32
	_ = v5303
	var v5305 int32
	_ = v5305
	var v5306 int32
	_ = v5306
	var v5308 int32
	_ = v5308
	var v5310 int32
	_ = v5310
	var v5314 int32
	_ = v5314
	var v5315 int32
	_ = v5315
	var v5316 int32
	_ = v5316
	var v5321 int32
	_ = v5321
	var v5322 int32
	_ = v5322
	var v5323 int32
	_ = v5323
	var v5325 int32
	_ = v5325
	var v5327 int32
	_ = v5327
	var v5332 int32
	_ = v5332
	var v5348 int32
	_ = v5348
	var v5349 int32
	_ = v5349
	var v5350 int32
	_ = v5350
	var v5382 int32
	_ = v5382
	var v5383 int32
	_ = v5383
	var v5384 int32
	_ = v5384
	var v5388 int32
	_ = v5388
	var v5399 int32
	_ = v5399
	var v5419 int32
	_ = v5419
	var v5429 int32
	_ = v5429
	var v5434 int32
	_ = v5434
	var v5436 int32
	_ = v5436
	var v5439 int32
	_ = v5439
	var v5440 int32
	_ = v5440
	var v5448 int32
	_ = v5448
	var v5450 int32
	_ = v5450
	var v5472 int32
	_ = v5472
	var v5502 int32
	_ = v5502
	var v5503 int32
	_ = v5503
	var v5515 int32
	_ = v5515
	var v5516 int32
	_ = v5516
	var v5519 int32
	_ = v5519
	var v5524 int64
	_ = v5524
	var v5531 int64
	_ = v5531
	var v5538 int32
	_ = v5538
	var v5542 int64
	_ = v5542
	var v5547 int64
	_ = v5547
	var v5551 int64
	_ = v5551
	var v5553 int32
	_ = v5553
	var v5554 int32
	_ = v5554
	var v5558 int64
	_ = v5558
	var v5565 int32
	_ = v5565
	var v5578 int64
	_ = v5578
	var v5580 int32
	_ = v5580
	var v5594 int64
	_ = v5594
	var v5611 int32
	_ = v5611
	var v5613 int32
	_ = v5613
	var v5625 int64
	_ = v5625
	var v5634 int32
	_ = v5634
	var v5635 int64
	_ = v5635
	var v5637 int64
	_ = v5637
	var v5639 int32
	_ = v5639
	var v5642 int64
	_ = v5642
	var v5644 int32
	_ = v5644
	var v5654 int64
	_ = v5654
	var v5670 int64
	_ = v5670
	var v5689 int32
	_ = v5689
	var v5721 int32
	_ = v5721
	var v5724 int32
	_ = v5724
	var v5728 int32
	_ = v5728
	var v5729 int32
	_ = v5729
	var v5733 int32
	_ = v5733
	var v5737 int32
	_ = v5737
	var v5750 int32
	_ = v5750
	var v5758 int32
	_ = v5758
	var v5759 int32
	_ = v5759
	var v5776 int32
	_ = v5776
	var v5779 int32
	_ = v5779
	var v5781 int32
	_ = v5781
	var v5784 int32
	_ = v5784
	var v5786 int32
	_ = v5786
	var v5798 int32
	_ = v5798
	var v5803 int32
	_ = v5803
	var v5808 int32
	_ = v5808
	var v5813 int32
	_ = v5813
	var v5814 int32
	_ = v5814
	var v5816 int32
	_ = v5816
	var v5828 int32
	_ = v5828
	var v5837 int32
	_ = v5837
	var v5856 int32
	_ = v5856
	var v5861 int32
	_ = v5861
	var v5867 int32
	_ = v5867
	var v5872 int32
	_ = v5872
	var v5873 int32
	_ = v5873
	var v5885 int32
	_ = v5885
	var v5918 int32
	_ = v5918
	var v5922 int32
	_ = v5922
	var v5927 int32
	_ = v5927
	var v5928 int32
	_ = v5928
	var v5929 int32
	_ = v5929
	var v5930 int32
	_ = v5930
	var v5934 int32
	_ = v5934
	var v5938 int32
	_ = v5938
	var v5945 int32
	_ = v5945
	var v5948 int32
	_ = v5948
	var v5955 int32
	_ = v5955
	var v5956 int32
	_ = v5956
	var v5957 int32
	_ = v5957
	var v5961 int32
	_ = v5961
	var v5963 int32
	_ = v5963
	var v5964 int32
	_ = v5964
	var v5966 int32
	_ = v5966
	var v5968 int32
	_ = v5968
	var v5979 int32
	_ = v5979
	var v5985 int32
	_ = v5985
	var v5986 int32
	_ = v5986
	var v5988 int32
	_ = v5988
	var v5994 int32
	_ = v5994
	var v6001 int32
	_ = v6001
	var v6005 int32
	_ = v6005
	var v6008 int32
	_ = v6008
	var v6014 int32
	_ = v6014
	var v6021 int32
	_ = v6021
	var v6025 int32
	_ = v6025
	var v6028 int32
	_ = v6028
	var v6031 int32
	_ = v6031
	var v6032 int32
	_ = v6032
	var v6033 int32
	_ = v6033
	var v6038 int32
	_ = v6038
	var v6039 int32
	_ = v6039
	var v6040 int32
	_ = v6040
	var v6042 int32
	_ = v6042
	var v6044 int32
	_ = v6044
	var v6045 int32
	_ = v6045
	var v6047 int32
	_ = v6047
	var v6049 int32
	_ = v6049
	var v6053 int32
	_ = v6053
	var v6054 int32
	_ = v6054
	var v6055 int32
	_ = v6055
	var v6060 int32
	_ = v6060
	var v6061 int32
	_ = v6061
	var v6062 int32
	_ = v6062
	var v6064 int32
	_ = v6064
	var v6066 int32
	_ = v6066
	var v6071 int32
	_ = v6071
	var v6087 int32
	_ = v6087
	var v6088 int32
	_ = v6088
	var v6119 int32
	_ = v6119
	var v6132 int32
	_ = v6132
	var v6149 int32
	_ = v6149
	var v6151 int32
	_ = v6151
	var v6152 int64
	_ = v6152
	var v6155 int32
	_ = v6155
	var v6158 int64
	_ = v6158
	var v6168 int32
	_ = v6168
	var v6175 int32
	_ = v6175
	var v6176 int32
	_ = v6176
	var v6179 int32
	_ = v6179
	var v6187 int32
	_ = v6187
	var v6193 int64
	_ = v6193
	var v6201 int32
	_ = v6201
	var v6209 int64
	_ = v6209
	var v6210 int64
	_ = v6210
	var v6211 int32
	_ = v6211
	var v6212 int32
	_ = v6212
	var v6219 int64
	_ = v6219
	var v6221 int64
	_ = v6221
	var v6237 int32
	_ = v6237
	var v6263 int32
	_ = v6263
	var v6275 int32
	_ = v6275
	var v6277 int32
	_ = v6277
	var v6279 int32
	_ = v6279
	var v6299 int32
	_ = v6299
	var v6318 int32
	_ = v6318
	var v6322 int32
	_ = v6322
	var v6325 int32
	_ = v6325
	var v6327 int32
	_ = v6327
	var v6334 int64
	_ = v6334
	var v6344 int32
	_ = v6344
	var v6348 int32
	_ = v6348
	var v6356 int64
	_ = v6356
	var v6383 int32
	_ = v6383
	var v6388 int32
	_ = v6388
	var v6398 int64
	_ = v6398
	var v6404 int32
	_ = v6404
	var v6412 int64
	_ = v6412
	var v6413 int32
	_ = v6413
	var v6415 int32
	_ = v6415
	var v6434 int32
	_ = v6434
	var v6439 int32
	_ = v6439
	var v6443 int64
	_ = v6443
	var v6444 int64
	_ = v6444
	var v6445 int64
	_ = v6445
	var v6449 int64
	_ = v6449
	var v6450 int64
	_ = v6450
	var v6451 int32
	_ = v6451
	var v6454 int64
	_ = v6454
	var v6468 int64
	_ = v6468
	var v6472 int64
	_ = v6472
	var v6474 int32
	_ = v6474
	var v6475 int32
	_ = v6475
	var v6479 int64
	_ = v6479
	var v6486 int32
	_ = v6486
	v9 = int64(0)
	v30 = m.G0
	v32 = v30 - int32(1536)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v36 = v34 + int32(32768)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui64(int64(9999)) < base.Ui64(v39) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if base.Ui64(int64(99)) < base.Ui64(v86) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v49 = v38
	v50 = v39
	goto L4
L3:
	;
	v84 = v38
	v86 = v39
	goto L1
L4:
	;
	v72 = v49 + int32(4)
	v76 = base.I64_div_u_s(v50, int64(10000))
	if base.Ui64(int64(99999999)) < base.Ui64(v50) {
		v49 = v72
		v50 = v76
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v84 = v72
	v86 = v76
	goto L1
L6:
	;
	goto L5
L7:
	;
	if base.Ui64(v151) <= base.Ui64(int64(9)) {
		v216 = v150
		v217 = v151
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v115 = v84
	v117 = v86
	goto L10
L9:
	;
	v150 = v84
	v151 = v86
	v152 = v86
	goto L7
L10:
	;
	v138 = v115 + int32(2)
	v142 = base.I64_div_u_s(v117, int64(100))
	if base.Ui64(int64(9999)) < base.Ui64(v117) {
		v115 = v138
		v117 = v142
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v150 = v138
	v151 = v142
	v152 = v142
	goto L7
L12:
	;
	goto L11
L13:
	;
	v238 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v238)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v243 = v241 + v242
	if base.Ui32(v242) < base.Ui32(int32(8)) {
		v295 = v241
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v181 = v150
	v182 = v151
	goto L15
L15:
	;
	v204 = v181 + int32(1)
	v208 = base.I64_div_u_s(v182, int64(10))
	if base.Ui64(int64(99)) < base.Ui64(v182) {
		v181 = v204
		v182 = v208
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v216 = v204
	v217 = v208
	goto L13
L17:
	;
	goto L16
L18:
	;
	if l3 != 0 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v258 = v241
	goto L20
L20:
	;
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v258)))
	if v275 != int64(3472328296227680304) {
		v295 = v258
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v295 = v279
	goto L18
L22:
	;
	v279 = v258 + int32(8)
	if base.Ui32(int32(7)) < base.Ui32(v243-v279) {
		v258 = v279
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v314 = int32(769)
	goto L26
L25:
	;
	v314 = int32(114)
	goto L26
L26:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v295 != v243 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v3362 = v216 - v3341 + int32(1)
	if v3362 < int32(0) {
		goto L290
	} else {
		goto L291
	}
L28:
	;
	if v496 == int32(0) {
		v2680 = v2557
		goto L218
	} else {
		goto L219
	}
L29:
	;
	if v511 == int32(0) {
		v2512 = v2436
		v2528 = v2452
		goto L211
	} else {
		goto L212
	}
L30:
	;
	if v316 == int32(0) {
		v3339 = v965
		v3340 = v966
		v3341 = v967
		v3342 = v968
		goto L27
	} else {
		goto L90
	}
L31:
	;
	v333 = v295
	goto L34
L32:
	;
	v965 = v217
	v966 = v152
	v967 = int32(0)
	v968 = v238
	v981 = v9
	goto L30
L33:
	;
	v357 = int32(0)
	if v356 != v243 {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	if v350 != int32(48) {
		v356 = v333
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v356 = v295 + (v243 - v295)
	goto L33
L36:
	;
	v354 = v333 + int32(1)
	if v354 != v243 {
		v333 = v354
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v360 = int32(0)
	v370 = v217
	v371 = v152
	v372 = v360
	v373 = v357
	v374 = v356
	v379 = v360
	v386 = v9
	goto L40
L39:
	;
	v965 = v217
	v966 = v152
	v967 = int32(0)
	v968 = v357
	v981 = v9
	goto L30
L40:
	;
	if v243-v374 < int32(8) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v965 = v811
	v966 = v812
	v967 = v488
	v968 = v938
	v981 = v827
	goto L30
L42:
	;
	if v426 != v243 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v426 = v374
	v427 = v371
	v428 = v372
	v429 = int32(0)
	v430 = int32(0)
	goto L42
L44:
	;
	if base.Ui32(v314-v372) < base.Ui32(int32(8)) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v374)))
	v400 = v398 + int64(-3472328296227680304)
	v405 = v400*int64(10) + int64(base.Ui64(v400)>>(uint(int64(8))%64))
	v408 = int64(1095216660735)
	v420 = int32(8)
	v426 = v374 + v420
	v427 = v405
	v428 = v372 + v420
	v429 = base.I32_wrap_i64(int64(base.Ui64(int64(base.Ui64(v405)>>(uint(int64(16))%64))&v408*int64(42949672960001)+v405&v408*int64(4294967296000100)) >> (uint(int64(32)) % 64)))
	v430 = v420
	goto L42
L46:
	;
	if v488 != v314 {
		goto L55
	} else {
		goto L56
	}
L47:
	;
	if base.Ui32(v314) <= base.Ui32(v428) {
		v488 = v428
		v490 = v426
		v496 = v429
		v498 = v430
		goto L46
	} else {
		goto L49
	}
L48:
	;
	v488 = v428
	v490 = v426
	v496 = v429
	v498 = v430
	goto L46
L49:
	;
	v434 = v426
	v443 = v428
	v451 = v429
	v452 = v430
	goto L50
L50:
	;
	v462 = int32(1)
	v463 = v443 + v462
	v465 = v452 + v462
	v467 = v434 + v462
	v470 = int32(*(*int8)(unsafe.Add(mBase, uint32(v434))))
	v473 = v451*int32(10) + v470 + int32(-48)
	if base.Ui32(int32(7)) < base.Ui32(v452) {
		v488 = v463
		v490 = v467
		v496 = v473
		v498 = v465
		goto L46
	} else {
		goto L52
	}
L51:
	;
	v488 = v463
	v490 = v467
	v496 = v473
	v498 = v465
	goto L46
L52:
	;
	if v467 == v243 {
		v488 = v463
		v490 = v467
		v496 = v473
		v498 = v465
		goto L46
	} else {
		goto L53
	}
L53:
	;
	if base.Ui32(v463) < base.Ui32(v314) {
		v434 = v467
		v443 = v463
		v451 = v473
		v452 = v465
		goto L50
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	v597 = v379 & int32(65535)
	if v597 != 0 {
		goto L65
	} else {
		goto L66
	}
L56:
	;
	v508 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	if v508 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v510 = int32(3)
	v511 = v508 & v510
	v516 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v498<<(uint(v510)%32))+uint32(_consts[921]))))
	if base.Ui32(int32(4)) <= base.Ui32(v508) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v2554 = v370
	v2555 = v427
	v2557 = int32(0)
	goto L28
L59:
	;
	v525 = int32(0)
	v535 = int64(0)
	v537 = v525
	v547 = v525
	goto L61
L60:
	;
	v2436 = int64(0)
	v2438 = int32(0)
	v2452 = v386
	goto L29
L61:
	;
	v560 = v32 + int32(24) + v537<<(uint(int32(2))%32)
	v561 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v560))))
	v563 = v516*v561 + v535
	*(*uint32)(unsafe.Add(mBase, uint32(v560))) = uint32(v563)
	v565 = int32(4)
	v566 = v560 + v565
	v567 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v566))))
	v569 = int64(32)
	v571 = v516*v567 + int64(base.Ui64(v563)>>(uint(v569)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v566))) = uint32(v571)
	v574 = v560 + int32(8)
	v575 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v574))))
	v579 = v516*v575 + int64(base.Ui64(v571)>>(uint(v569)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v574))) = uint32(v579)
	v582 = v560 + int32(12)
	v583 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v582))))
	v587 = v516*v583 + int64(base.Ui64(v579)>>(uint(v569)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v582))) = uint32(v587)
	v590 = int64(base.Ui64(v587) >> (uint(v569) % 64))
	v592 = v537 + v565
	v594 = v547 + v565
	if v594 != v508&int32(65532) {
		v535 = v590
		v537 = v592
		v547 = v594
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v2436 = v590
	v2438 = v592
	v2452 = v587
	goto L29
L64:
	;
	if v496 == int32(0) {
		v938 = v814
		v944 = v820
		goto L79
	} else {
		goto L80
	}
L65:
	;
	v599 = int32(3)
	v600 = v597 & v599
	v605 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v498<<(uint(v599)%32))+uint32(_consts[921]))))
	v607 = int32(0)
	if base.Ui32(v597) < base.Ui32(int32(4)) {
		v686 = v607
		v693 = int64(0)
		v709 = v386
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v811 = v370
	v812 = v427
	v814 = v373
	v820 = int32(0)
	v827 = v386
	goto L64
L67:
	;
	if v600 == int32(0) {
		v769 = v693
		v785 = v709
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v614 = int32(0)
	v617 = v614
	v624 = int64(0)
	v639 = v614
	goto L69
L69:
	;
	v649 = v32 + int32(24) + v617<<(uint(int32(2))%32)
	v650 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v649))))
	v652 = v605*v650 + v624
	*(*uint32)(unsafe.Add(mBase, uint32(v649))) = uint32(v652)
	v654 = int32(4)
	v655 = v649 + v654
	v656 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v655))))
	v658 = int64(32)
	v660 = v605*v656 + int64(base.Ui64(v652)>>(uint(v658)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v655))) = uint32(v660)
	v663 = v649 + int32(8)
	v664 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v663))))
	v668 = v605*v664 + int64(base.Ui64(v660)>>(uint(v658)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v663))) = uint32(v668)
	v671 = v649 + int32(12)
	v672 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v671))))
	v676 = v605*v672 + int64(base.Ui64(v668)>>(uint(v658)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v671))) = uint32(v676)
	v679 = int64(base.Ui64(v676) >> (uint(v658) % 64))
	v681 = v617 + v654
	v683 = v639 + v654
	if v683 != v597&int32(65532) {
		v617 = v681
		v624 = v679
		v639 = v683
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v686 = v681
	v693 = v679
	v709 = v676
	goto L67
L71:
	;
	goto L70
L72:
	;
	if base.Ui32(int32(124)) < base.Ui32(v597) {
		v811 = v769
		v812 = v605
		v814 = v373
		v820 = v379
		v827 = v785
		goto L64
	} else {
		goto L77
	}
L73:
	;
	v717 = v686
	v724 = v693
	v736 = v607
	goto L74
L74:
	;
	v749 = v32 + int32(24) + v717<<(uint(int32(2))%32)
	v750 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v749))))
	v752 = v605*v750 + v724
	*(*uint32)(unsafe.Add(mBase, uint32(v749))) = uint32(v752)
	v754 = int32(1)
	v757 = int64(base.Ui64(v752) >> (uint(int64(32)) % 64))
	v759 = v736 + v754
	if v759 != v600 {
		v717 = v717 + v754
		v724 = v757
		v736 = v759
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v769 = v757
	v785 = v752
	goto L72
L76:
	;
	goto L75
L77:
	;
	if base.Ui64(v785) <= base.Ui64(int64(4294967295)) {
		v811 = v769
		v812 = v605
		v814 = v373
		v820 = v379
		v827 = v785
		goto L64
	} else {
		goto L78
	}
L78:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v32+int32(24)+v597<<(uint(int32(2))%32)))) = uint32(v769)
	v801 = v379 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v801)
	v811 = v769
	v812 = v605
	v814 = v801
	v820 = v801
	v827 = v785
	goto L64
L79:
	;
	if v490 != v243 {
		v370 = v811
		v371 = v812
		v372 = v488
		v373 = v938
		v374 = v490
		v379 = v944
		v386 = v827
		goto L40
	} else {
		goto L89
	}
L80:
	;
	v835 = v820 & int32(65535)
	if v835 == int32(0) {
		v905 = v496
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if base.Ui32(int32(124)) < base.Ui32(v835) {
		v938 = v814
		v944 = v820
		goto L79
	} else {
		goto L88
	}
L82:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v839 = v838 + v496
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v839
	v841 = int32(1)
	if base.Ui32(v838) <= base.Ui32(v839) {
		v938 = v814
		v944 = v820
		goto L79
	} else {
		goto L83
	}
L83:
	;
	v845 = v841
	goto L84
L84:
	;
	if v845 == v835 {
		v905 = v841
		goto L81
	} else {
		goto L86
	}
L86:
	;
	v878 = v32 + int32(24) + v845<<(uint(int32(2))%32)
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v878)))
	v880 = int32(1)
	v881 = v879 + v880
	*(*int32)(unsafe.Add(mBase, uint32(v878))) = v881
	if v881 == int32(0) {
		v845 = v845 + v880
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v938 = v814
	v944 = v820
	goto L79
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(24)+v835<<(uint(int32(2))%32)))) = v905
	v925 = v820 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v925)
	v938 = v925
	v944 = v925
	goto L79
L89:
	;
	goto L41
L90:
	;
	v988 = v316 + v315
	if v967 != 0 {
		v1111 = v316
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if v1111 == v988 {
		v3339 = v965
		v3340 = v966
		v3341 = v967
		v3342 = v968
		goto L27
	} else {
		goto L104
	}
L92:
	;
	if base.Ui32(v315) < base.Ui32(int32(8)) {
		v1044 = v316
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v1044 == v988 {
		v3339 = v965
		v3340 = v966
		v3341 = v967
		v3342 = v968
		goto L27
	} else {
		goto L99
	}
L94:
	;
	v1007 = v316
	goto L95
L95:
	;
	v1020 = *(*int64)(unsafe.Add(mBase, uint32(v1007)))
	if v1020 != int64(3472328296227680304) {
		v1044 = v1007
		goto L93
	} else {
		goto L97
	}
L96:
	;
	v1044 = v1024
	goto L93
L97:
	;
	v1024 = v1007 + int32(8)
	if base.Ui32(int32(7)) < base.Ui32(v988-v1024) {
		v1007 = v1024
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v1076 = v1044
	goto L100
L100:
	;
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1076))))
	if v1089 != int32(48) {
		v1111 = v1076
		goto L91
	} else {
		goto L102
	}
L101:
	;
	v1111 = v1044 + (v988 - v1044)
	goto L91
L102:
	;
	v1093 = v1076 + int32(1)
	if v1093 != v988 {
		v1076 = v1093
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v1125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	v1134 = v965
	v1135 = v966
	v1136 = v967
	v1137 = v1125
	v1142 = v1111
	v1150 = v981
	v1151 = v1125
	goto L106
L105:
	;
	if v1261 == int32(0) {
		v1973 = v1850
		goto L164
	} else {
		goto L165
	}
L106:
	;
	if v988-v1142 < int32(8) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	if v1275 == int32(0) {
		v1805 = v1729
		v1821 = v1745
		goto L157
	} else {
		goto L158
	}
L108:
	;
	if v1190 != v988 {
		goto L113
	} else {
		goto L114
	}
L109:
	;
	v1190 = v1142
	v1191 = v1135
	v1192 = v1136
	v1193 = int32(0)
	v1194 = int32(0)
	goto L108
L110:
	;
	if base.Ui32(v314-v1136) < base.Ui32(int32(8)) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v1162 = *(*int64)(unsafe.Add(mBase, uint32(v1142)))
	v1164 = v1162 + int64(-3472328296227680304)
	v1169 = v1164*int64(10) + int64(base.Ui64(v1164)>>(uint(int64(8))%64))
	v1172 = int64(1095216660735)
	v1184 = int32(8)
	v1190 = v1142 + v1184
	v1191 = v1169
	v1192 = v1136 + v1184
	v1193 = v1184
	v1194 = base.I32_wrap_i64(int64(base.Ui64(int64(base.Ui64(v1169)>>(uint(int64(16))%64))&v1172*int64(42949672960001)+v1169&v1172*int64(4294967296000100)) >> (uint(int64(32)) % 64)))
	goto L108
L112:
	;
	if v1252 != v314 {
		goto L122
	} else {
		goto L123
	}
L113:
	;
	if base.Ui32(v314) <= base.Ui32(v1192) {
		v1252 = v1192
		v1258 = v1190
		v1260 = v1193
		v1261 = v1194
		goto L112
	} else {
		goto L115
	}
L114:
	;
	v1252 = v1192
	v1258 = v1190
	v1260 = v1193
	v1261 = v1194
	goto L112
L115:
	;
	v1198 = v1190
	v1207 = v1192
	v1209 = v1193
	v1216 = v1194
	goto L116
L116:
	;
	v1226 = int32(1)
	v1227 = v1207 + v1226
	v1229 = v1209 + v1226
	v1231 = v1198 + v1226
	v1234 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1198))))
	v1237 = v1216*int32(10) + v1234 + int32(-48)
	if base.Ui32(int32(7)) < base.Ui32(v1209) {
		v1252 = v1227
		v1258 = v1231
		v1260 = v1229
		v1261 = v1237
		goto L112
	} else {
		goto L118
	}
L117:
	;
	v1252 = v1227
	v1258 = v1231
	v1260 = v1229
	v1261 = v1237
	goto L112
L118:
	;
	if v1231 == v988 {
		v1252 = v1227
		v1258 = v1231
		v1260 = v1229
		v1261 = v1237
		goto L112
	} else {
		goto L119
	}
L119:
	;
	if base.Ui32(v1227) < base.Ui32(v314) {
		v1198 = v1231
		v1207 = v1227
		v1209 = v1229
		v1216 = v1237
		goto L116
	} else {
		goto L120
	}
L120:
	;
	goto L117
L121:
	;
	goto L107
L122:
	;
	v1361 = v1151 & int32(65535)
	if v1361 != 0 {
		goto L132
	} else {
		goto L133
	}
L123:
	;
	v1272 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	if v1272 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v1274 = int32(3)
	v1275 = v1272 & v1274
	v1280 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1260<<(uint(v1274)%32))+uint32(_consts[921]))))
	if base.Ui32(int32(4)) <= base.Ui32(v1272) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v1847 = v1134
	v1848 = v1191
	v1850 = int32(0)
	goto L105
L126:
	;
	v1289 = int32(0)
	v1299 = int64(0)
	v1301 = v1289
	v1309 = v1289
	goto L128
L127:
	;
	v1729 = int64(0)
	v1731 = int32(0)
	v1745 = v1150
	goto L121
L128:
	;
	v1324 = v32 + int32(24) + v1301<<(uint(int32(2))%32)
	v1325 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1324))))
	v1327 = v1280*v1325 + v1299
	*(*uint32)(unsafe.Add(mBase, uint32(v1324))) = uint32(v1327)
	v1329 = int32(4)
	v1330 = v1324 + v1329
	v1331 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1330))))
	v1333 = int64(32)
	v1335 = v1280*v1331 + int64(base.Ui64(v1327)>>(uint(v1333)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1330))) = uint32(v1335)
	v1338 = v1324 + int32(8)
	v1339 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1338))))
	v1343 = v1280*v1339 + int64(base.Ui64(v1335)>>(uint(v1333)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1338))) = uint32(v1343)
	v1346 = v1324 + int32(12)
	v1347 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1346))))
	v1351 = v1280*v1347 + int64(base.Ui64(v1343)>>(uint(v1333)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1346))) = uint32(v1351)
	v1354 = int64(base.Ui64(v1351) >> (uint(v1333) % 64))
	v1356 = v1301 + v1329
	v1358 = v1309 + v1329
	if v1358 != v1272&int32(65532) {
		v1299 = v1354
		v1301 = v1356
		v1309 = v1358
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v1729 = v1354
	v1731 = v1356
	v1745 = v1351
	goto L121
L131:
	;
	if v1261 == int32(0) {
		v1702 = v1578
		v1716 = v1592
		goto L146
	} else {
		goto L147
	}
L132:
	;
	v1363 = int32(3)
	v1364 = v1361 & v1363
	v1369 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1260<<(uint(v1363)%32))+uint32(_consts[921]))))
	v1371 = int32(0)
	if base.Ui32(v1361) < base.Ui32(int32(4)) {
		v1450 = v1371
		v1457 = int64(0)
		v1473 = v1150
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v1575 = v1134
	v1576 = v1191
	v1578 = v1137
	v1591 = v1150
	v1592 = int32(0)
	goto L131
L134:
	;
	if v1364 == int32(0) {
		v1533 = v1457
		v1549 = v1473
		goto L139
	} else {
		goto L140
	}
L135:
	;
	v1378 = int32(0)
	v1381 = v1378
	v1388 = int64(0)
	v1393 = v1378
	goto L136
L136:
	;
	v1413 = v32 + int32(24) + v1381<<(uint(int32(2))%32)
	v1414 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1413))))
	v1416 = v1369*v1414 + v1388
	*(*uint32)(unsafe.Add(mBase, uint32(v1413))) = uint32(v1416)
	v1418 = int32(4)
	v1419 = v1413 + v1418
	v1420 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1419))))
	v1422 = int64(32)
	v1424 = v1369*v1420 + int64(base.Ui64(v1416)>>(uint(v1422)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1419))) = uint32(v1424)
	v1427 = v1413 + int32(8)
	v1428 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1427))))
	v1432 = v1369*v1428 + int64(base.Ui64(v1424)>>(uint(v1422)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1427))) = uint32(v1432)
	v1435 = v1413 + int32(12)
	v1436 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1435))))
	v1440 = v1369*v1436 + int64(base.Ui64(v1432)>>(uint(v1422)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1435))) = uint32(v1440)
	v1443 = int64(base.Ui64(v1440) >> (uint(v1422) % 64))
	v1445 = v1381 + v1418
	v1447 = v1393 + v1418
	if v1447 != v1361&int32(65532) {
		v1381 = v1445
		v1388 = v1443
		v1393 = v1447
		goto L136
	} else {
		goto L138
	}
L137:
	;
	v1450 = v1445
	v1457 = v1443
	v1473 = v1440
	goto L134
L138:
	;
	goto L137
L139:
	;
	if base.Ui32(int32(124)) < base.Ui32(v1361) {
		v1575 = v1533
		v1576 = v1369
		v1578 = v1137
		v1591 = v1549
		v1592 = v1151
		goto L131
	} else {
		goto L144
	}
L140:
	;
	v1481 = v1450
	v1488 = v1457
	v1498 = v1371
	goto L141
L141:
	;
	v1513 = v32 + int32(24) + v1481<<(uint(int32(2))%32)
	v1514 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1513))))
	v1516 = v1369*v1514 + v1488
	*(*uint32)(unsafe.Add(mBase, uint32(v1513))) = uint32(v1516)
	v1518 = int32(1)
	v1521 = int64(base.Ui64(v1516) >> (uint(int64(32)) % 64))
	v1523 = v1498 + v1518
	if v1523 != v1364 {
		v1481 = v1481 + v1518
		v1488 = v1521
		v1498 = v1523
		goto L141
	} else {
		goto L143
	}
L142:
	;
	v1533 = v1521
	v1549 = v1516
	goto L139
L143:
	;
	goto L142
L144:
	;
	if base.Ui64(v1549) <= base.Ui64(int64(4294967295)) {
		v1575 = v1533
		v1576 = v1369
		v1578 = v1137
		v1591 = v1549
		v1592 = v1151
		goto L131
	} else {
		goto L145
	}
L145:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v32+int32(24)+v1361<<(uint(int32(2))%32)))) = uint32(v1533)
	v1565 = v1151 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v1565)
	v1575 = v1533
	v1576 = v1369
	v1578 = v1565
	v1591 = v1549
	v1592 = v1565
	goto L131
L146:
	;
	if v1258 != v988 {
		v1134 = v1575
		v1135 = v1576
		v1136 = v1252
		v1137 = v1702
		v1142 = v1258
		v1150 = v1591
		v1151 = v1716
		goto L106
	} else {
		goto L156
	}
L147:
	;
	v1599 = v1592 & int32(65535)
	if v1599 == int32(0) {
		v1670 = v1261
		goto L148
	} else {
		goto L149
	}
L148:
	;
	if base.Ui32(int32(124)) < base.Ui32(v1599) {
		v1702 = v1578
		v1716 = v1592
		goto L146
	} else {
		goto L155
	}
L149:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v1603 = v1602 + v1261
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v1603
	v1605 = int32(1)
	if base.Ui32(v1602) <= base.Ui32(v1603) {
		v1702 = v1578
		v1716 = v1592
		goto L146
	} else {
		goto L150
	}
L150:
	;
	v1609 = v1605
	goto L151
L151:
	;
	if v1609 == v1599 {
		v1670 = v1605
		goto L148
	} else {
		goto L153
	}
L153:
	;
	v1642 = v32 + int32(24) + v1609<<(uint(int32(2))%32)
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1642)))
	v1644 = int32(1)
	v1645 = v1643 + v1644
	*(*int32)(unsafe.Add(mBase, uint32(v1642))) = v1645
	if v1645 == int32(0) {
		v1609 = v1609 + v1644
		goto L151
	} else {
		goto L154
	}
L154:
	;
	v1702 = v1578
	v1716 = v1592
	goto L146
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(24)+v1599<<(uint(int32(2))%32)))) = v1670
	v1689 = v1592 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v1689)
	v1702 = v1689
	v1716 = v1689
	goto L146
L156:
	;
	v3339 = v1575
	v3340 = v1576
	v3341 = v1252
	v3342 = v1702
	goto L27
L157:
	;
	if base.Ui32(int32(124)) < base.Ui32(v1272) {
		v1847 = v1805
		v1848 = v1280
		v1850 = v1272
		goto L105
	} else {
		goto L162
	}
L158:
	;
	v1760 = v1729
	v1762 = v1731
	v1764 = int32(0)
	goto L159
L159:
	;
	v1785 = v32 + int32(24) + v1762<<(uint(int32(2))%32)
	v1786 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1785))))
	v1788 = v1280*v1786 + v1760
	*(*uint32)(unsafe.Add(mBase, uint32(v1785))) = uint32(v1788)
	v1790 = int32(1)
	v1793 = int64(base.Ui64(v1788) >> (uint(int64(32)) % 64))
	v1795 = v1764 + v1790
	if v1795 != v1275 {
		v1760 = v1793
		v1762 = v1762 + v1790
		v1764 = v1795
		goto L159
	} else {
		goto L161
	}
L160:
	;
	v1805 = v1793
	v1821 = v1788
	goto L157
L161:
	;
	goto L160
L162:
	;
	if base.Ui64(v1821) <= base.Ui64(int64(4294967295)) {
		v1847 = v1805
		v1848 = v1280
		v1850 = v1272
		goto L105
	} else {
		goto L163
	}
L163:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v32+int32(24)+v1272<<(uint(int32(2))%32)))) = uint32(v1805)
	v1837 = v1272 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v1837)
	v1847 = v1805
	v1848 = v1280
	v1850 = v1837
	goto L105
L164:
	;
	if base.Ui32(v988-v1258) <= base.Ui32(int32(7)) {
		v2047 = v1258
		goto L175
	} else {
		goto L176
	}
L165:
	;
	v1870 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	if v1870 == int32(0) {
		v1941 = v1261
		goto L166
	} else {
		goto L167
	}
L166:
	;
	if base.Ui32(int32(124)) < base.Ui32(v1870) {
		v1973 = v1870
		goto L164
	} else {
		goto L173
	}
L167:
	;
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v1874 = v1873 + v1261
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v1874
	if base.Ui32(v1873) <= base.Ui32(v1874) {
		v1973 = v1870
		goto L164
	} else {
		goto L168
	}
L168:
	;
	v1877 = int32(1)
	v1889 = v1877
	goto L169
L169:
	;
	if v1889 == v1870 {
		v1941 = v1877
		goto L166
	} else {
		goto L171
	}
L171:
	;
	v1913 = v32 + int32(24) + v1889<<(uint(int32(2))%32)
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v1913)))
	v1915 = int32(1)
	v1916 = v1914 + v1915
	*(*int32)(unsafe.Add(mBase, uint32(v1913))) = v1916
	if v1916 == int32(0) {
		v1889 = v1889 + v1915
		goto L169
	} else {
		goto L172
	}
L172:
	;
	v1973 = v1870
	goto L164
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(24)+v1870<<(uint(int32(2))%32)))) = v1941
	v1960 = v1870 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v1960)
	v1973 = v1960
	goto L164
L174:
	;
	v2125 = int32(0)
	v2127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	if v2127 == v2125 {
		v2343 = v1847
		v2344 = v1848
		v2346 = v2125
		goto L187
	} else {
		goto L188
	}
L175:
	;
	if v2047 != v988 {
		goto L181
	} else {
		goto L182
	}
L176:
	;
	v2010 = v1258
	goto L177
L177:
	;
	v2023 = *(*int64)(unsafe.Add(mBase, uint32(v2010)))
	if v2023 != int64(3472328296227680304) {
		goto L174
	} else {
		goto L179
	}
L178:
	;
	v2047 = v2027
	goto L175
L179:
	;
	v2027 = v2010 + int32(8)
	if base.Ui32(int32(7)) < base.Ui32(v988-v2027) {
		v2010 = v2027
		goto L177
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	v2077 = v2047
	goto L183
L182:
	;
	v3339 = v1847
	v3340 = v1848
	v3341 = v314
	v3342 = v1973
	goto L27
L183:
	;
	v2090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2077))))
	if v2090 != int32(48) {
		goto L174
	} else {
		goto L185
	}
L184:
	;
	v3339 = v1847
	v3340 = v1848
	v3341 = v314
	v3342 = v1973
	goto L27
L185:
	;
	v2094 = v2077 + int32(1)
	if v2094 != v988 {
		v2077 = v2094
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v2365 = v2346 & int32(65535)
	v2367 = v2125
	goto L206
L188:
	;
	v2131 = v2127 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v2127) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	if v2131 == int32(0) {
		v2300 = v2223
		v2301 = v2224
		goto L195
	} else {
		goto L196
	}
L190:
	;
	v2140 = int32(0)
	v2151 = int64(0)
	v2152 = v2140
	v2158 = v2140
	goto L192
L191:
	;
	v2223 = v1847
	v2224 = int64(0)
	v2225 = int32(0)
	goto L189
L192:
	;
	v2175 = v32 + int32(24) + v2152<<(uint(int32(2))%32)
	v2176 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2175))))
	v2177 = int64(10)
	v2179 = v2176*v2177 + v2151
	*(*uint32)(unsafe.Add(mBase, uint32(v2175))) = uint32(v2179)
	v2181 = int32(4)
	v2182 = v2175 + v2181
	v2183 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2182))))
	v2186 = int64(32)
	v2188 = v2183*v2177 + int64(base.Ui64(v2179)>>(uint(v2186)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v2182))) = uint32(v2188)
	v2191 = v2175 + int32(8)
	v2192 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2191))))
	v2197 = v2192*v2177 + int64(base.Ui64(v2188)>>(uint(v2186)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v2191))) = uint32(v2197)
	v2200 = v2175 + int32(12)
	v2201 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2200))))
	v2206 = v2201*v2177 + int64(base.Ui64(v2197)>>(uint(v2186)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v2200))) = uint32(v2206)
	v2209 = int64(base.Ui64(v2206) >> (uint(v2186) % 64))
	v2211 = v2152 + v2181
	v2213 = v2158 + v2181
	if v2213 != v2127&int32(65532) {
		v2151 = v2209
		v2152 = v2211
		v2158 = v2213
		goto L192
	} else {
		goto L194
	}
L193:
	;
	v2223 = v2206
	v2224 = v2209
	v2225 = v2211
	goto L189
L194:
	;
	goto L193
L195:
	;
	if base.Ui32(int32(124)) < base.Ui32(v2127) {
		goto L201
	} else {
		goto L202
	}
L196:
	;
	v2255 = v2224
	v2256 = v2225
	v2265 = int32(0)
	goto L197
L197:
	;
	v2279 = v32 + int32(24) + v2256<<(uint(int32(2))%32)
	v2280 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2279))))
	v2283 = v2280*int64(10) + v2255
	*(*uint32)(unsafe.Add(mBase, uint32(v2279))) = uint32(v2283)
	v2285 = int32(1)
	v2288 = int64(base.Ui64(v2283) >> (uint(int64(32)) % 64))
	v2290 = v2265 + v2285
	if v2290 != v2131 {
		v2255 = v2288
		v2256 = v2256 + v2285
		v2265 = v2290
		goto L197
	} else {
		goto L199
	}
L198:
	;
	v2300 = v2283
	v2301 = v2288
	goto L195
L199:
	;
	goto L198
L200:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v32+int32(24)+v2127<<(uint(int32(2))%32)))) = uint32(v2301)
	v2333 = v2127 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v2333)
	v2343 = v2300
	v2344 = v2301
	v2346 = v2333
	goto L187
L201:
	;
	v2325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	v2343 = v2300
	v2344 = v2301
	v2346 = v2325
	goto L187
L202:
	;
	if base.Ui64(int64(4294967295)) < base.Ui64(v2300) {
		goto L200
	} else {
		goto L203
	}
L203:
	;
	goto L201
L204:
	;
	v3339 = v2343
	v3340 = v2344
	v3341 = v314 + int32(1)
	v3342 = v2425
	goto L27
L205:
	;
	if base.Ui32(int32(124)) < base.Ui32(v2346&int32(65535)) {
		v2425 = v2346
		goto L204
	} else {
		goto L210
	}
L206:
	;
	if v2367 == v2365 {
		goto L205
	} else {
		goto L208
	}
L208:
	;
	v2400 = v32 + int32(24) + v2367<<(uint(int32(2))%32)
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(v2400)))
	v2402 = int32(1)
	v2403 = v2401 + v2402
	*(*int32)(unsafe.Add(mBase, uint32(v2400))) = v2403
	if v2403 == int32(0) {
		v2367 = v2367 + v2402
		goto L206
	} else {
		goto L209
	}
L209:
	;
	v2425 = v2346
	goto L204
L210:
	;
	v2418 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(24)+v2365<<(uint(int32(2))%32)))) = v2418
	v2421 = v2346 + v2418
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v2421)
	v2425 = v2421
	goto L204
L211:
	;
	if base.Ui32(int32(124)) < base.Ui32(v508) {
		v2554 = v2512
		v2555 = v516
		v2557 = v508
		goto L28
	} else {
		goto L216
	}
L212:
	;
	v2467 = v2436
	v2469 = v2438
	v2478 = int32(0)
	goto L213
L213:
	;
	v2492 = v32 + int32(24) + v2469<<(uint(int32(2))%32)
	v2493 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2492))))
	v2495 = v516*v2493 + v2467
	*(*uint32)(unsafe.Add(mBase, uint32(v2492))) = uint32(v2495)
	v2497 = int32(1)
	v2500 = int64(base.Ui64(v2495) >> (uint(int64(32)) % 64))
	v2502 = v2478 + v2497
	if v2502 != v511 {
		v2467 = v2500
		v2469 = v2469 + v2497
		v2478 = v2502
		goto L213
	} else {
		goto L215
	}
L214:
	;
	v2512 = v2500
	v2528 = v2495
	goto L211
L215:
	;
	goto L214
L216:
	;
	if base.Ui64(v2528) <= base.Ui64(int64(4294967295)) {
		v2554 = v2512
		v2555 = v516
		v2557 = v508
		goto L28
	} else {
		goto L217
	}
L217:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v32+int32(24)+v508<<(uint(int32(2))%32)))) = uint32(v2512)
	v2544 = v508 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v2544)
	v2554 = v2512
	v2555 = v516
	v2557 = v2544
	goto L28
L218:
	;
	if base.Ui32(v243-v490) <= base.Ui32(int32(7)) {
		v2751 = v490
		goto L233
	} else {
		goto L234
	}
L219:
	;
	v2577 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	if v2577 == int32(0) {
		v2647 = v496
		goto L220
	} else {
		goto L221
	}
L220:
	;
	if base.Ui32(int32(124)) < base.Ui32(v2577) {
		v2680 = v2577
		goto L218
	} else {
		goto L227
	}
L221:
	;
	v2580 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v2581 = v2580 + v496
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v2581
	if base.Ui32(v2580) <= base.Ui32(v2581) {
		v2680 = v2577
		goto L218
	} else {
		goto L222
	}
L222:
	;
	v2584 = int32(1)
	v2596 = v2584
	goto L223
L223:
	;
	if v2596 == v2577 {
		v2647 = v2584
		goto L220
	} else {
		goto L225
	}
L225:
	;
	v2620 = v32 + int32(24) + v2596<<(uint(int32(2))%32)
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v2620)))
	v2622 = int32(1)
	v2623 = v2621 + v2622
	*(*int32)(unsafe.Add(mBase, uint32(v2620))) = v2623
	if v2623 == int32(0) {
		v2596 = v2596 + v2622
		goto L223
	} else {
		goto L226
	}
L226:
	;
	v2680 = v2577
	goto L218
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(24)+v2577<<(uint(int32(2))%32)))) = v2647
	v2667 = v2577 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v2667)
	v2680 = v2667
	goto L218
L228:
	;
	v3028 = int32(0)
	v3030 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	if v3030 == v3028 {
		v3246 = v2554
		v3247 = v2555
		v3249 = v3028
		goto L261
	} else {
		goto L262
	}
L229:
	;
	v2865 = v316 + v315
	if base.Ui32(v315) <= base.Ui32(int32(7)) {
		v2921 = v316
		goto L248
	} else {
		goto L249
	}
L230:
	;
	if v2814 != 0 {
		goto L228
	} else {
		goto L247
	}
L231:
	;
	if v316 != 0 {
		v2846 = int32(0)
		goto L229
	} else {
		goto L246
	}
L232:
	;
	if v316 == int32(0) {
		goto L230
	} else {
		goto L245
	}
L233:
	;
	if v2751 == v243 {
		goto L231
	} else {
		goto L240
	}
L234:
	;
	v2713 = v490
	goto L235
L235:
	;
	v2730 = *(*int64)(unsafe.Add(mBase, uint32(v2713)))
	if v2730 == int64(3472328296227680304) {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v2751 = v2735
	goto L233
L237:
	;
	v2735 = v2713 + int32(8)
	if base.Ui32(int32(7)) < base.Ui32(v243-v2735) {
		v2713 = v2735
		goto L235
	} else {
		goto L239
	}
L238:
	;
	v2814 = int32(1)
	goto L232
L239:
	;
	goto L236
L240:
	;
	v2781 = v2751
	goto L241
L241:
	;
	v2798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2781))))
	v2800 = base.B2i32(v2798 != int32(48))
	if v2798 != int32(48) {
		v2814 = v2800
		goto L232
	} else {
		goto L243
	}
L242:
	;
	v2814 = v2800
	goto L232
L243:
	;
	v2802 = v2781 + int32(1)
	if v2802 != v243 {
		v2781 = v2802
		goto L241
	} else {
		goto L244
	}
L244:
	;
	goto L242
L245:
	;
	v2846 = v2814
	goto L229
L246:
	;
	v3339 = v2554
	v3340 = v2555
	v3341 = v314
	v3342 = v2680
	goto L27
L247:
	;
	v3339 = v2554
	v3340 = v2555
	v3341 = v314
	v3342 = v2680
	goto L27
L248:
	;
	if v2921 == v2865 {
		goto L254
	} else {
		goto L255
	}
L249:
	;
	v2884 = v316
	goto L250
L250:
	;
	v2897 = *(*int64)(unsafe.Add(mBase, uint32(v2884)))
	if v2897 != int64(3472328296227680304) {
		goto L228
	} else {
		goto L252
	}
L251:
	;
	v2921 = v2901
	goto L248
L252:
	;
	v2901 = v2884 + int32(8)
	if base.Ui32(int32(7)) < base.Ui32(v2865-v2901) {
		v2884 = v2901
		goto L250
	} else {
		goto L253
	}
L253:
	;
	goto L251
L254:
	;
	if v2846 != 0 {
		goto L228
	} else {
		goto L260
	}
L255:
	;
	v2951 = v2921
	goto L256
L256:
	;
	v2964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2951))))
	if v2964 != int32(48) {
		goto L228
	} else {
		goto L258
	}
L257:
	;
	goto L254
L258:
	;
	v2968 = v2951 + int32(1)
	if v2968 != v2865 {
		v2951 = v2968
		goto L256
	} else {
		goto L259
	}
L259:
	;
	goto L257
L260:
	;
	v3339 = v2554
	v3340 = v2555
	v3341 = v314
	v3342 = v2680
	goto L27
L261:
	;
	v3268 = v3249 & int32(65535)
	v3270 = v3028
	goto L280
L262:
	;
	v3034 = v3030 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v3030) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	if v3034 == int32(0) {
		v3203 = v3126
		v3204 = v3127
		goto L269
	} else {
		goto L270
	}
L264:
	;
	v3043 = int32(0)
	v3054 = int64(0)
	v3055 = v3043
	v3061 = v3043
	goto L266
L265:
	;
	v3126 = v2554
	v3127 = int64(0)
	v3128 = int32(0)
	goto L263
L266:
	;
	v3078 = v32 + int32(24) + v3055<<(uint(int32(2))%32)
	v3079 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3078))))
	v3080 = int64(10)
	v3082 = v3079*v3080 + v3054
	*(*uint32)(unsafe.Add(mBase, uint32(v3078))) = uint32(v3082)
	v3084 = int32(4)
	v3085 = v3078 + v3084
	v3086 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3085))))
	v3089 = int64(32)
	v3091 = v3086*v3080 + int64(base.Ui64(v3082)>>(uint(v3089)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3085))) = uint32(v3091)
	v3094 = v3078 + int32(8)
	v3095 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3094))))
	v3100 = v3095*v3080 + int64(base.Ui64(v3091)>>(uint(v3089)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3094))) = uint32(v3100)
	v3103 = v3078 + int32(12)
	v3104 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3103))))
	v3109 = v3104*v3080 + int64(base.Ui64(v3100)>>(uint(v3089)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3103))) = uint32(v3109)
	v3112 = int64(base.Ui64(v3109) >> (uint(v3089) % 64))
	v3114 = v3055 + v3084
	v3116 = v3061 + v3084
	if v3116 != v3030&int32(65532) {
		v3054 = v3112
		v3055 = v3114
		v3061 = v3116
		goto L266
	} else {
		goto L268
	}
L267:
	;
	v3126 = v3109
	v3127 = v3112
	v3128 = v3114
	goto L263
L268:
	;
	goto L267
L269:
	;
	if base.Ui32(int32(124)) < base.Ui32(v3030) {
		goto L275
	} else {
		goto L276
	}
L270:
	;
	v3158 = v3127
	v3159 = v3128
	v3168 = int32(0)
	goto L271
L271:
	;
	v3182 = v32 + int32(24) + v3159<<(uint(int32(2))%32)
	v3183 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3182))))
	v3186 = v3183*int64(10) + v3158
	*(*uint32)(unsafe.Add(mBase, uint32(v3182))) = uint32(v3186)
	v3188 = int32(1)
	v3191 = int64(base.Ui64(v3186) >> (uint(int64(32)) % 64))
	v3193 = v3168 + v3188
	if v3193 != v3034 {
		v3158 = v3191
		v3159 = v3159 + v3188
		v3168 = v3193
		goto L271
	} else {
		goto L273
	}
L272:
	;
	v3203 = v3186
	v3204 = v3191
	goto L269
L273:
	;
	goto L272
L274:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v32+int32(24)+v3030<<(uint(int32(2))%32)))) = uint32(v3204)
	v3236 = v3030 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v3236)
	v3246 = v3203
	v3247 = v3204
	v3249 = v3236
	goto L261
L275:
	;
	v3228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	v3246 = v3203
	v3247 = v3204
	v3249 = v3228
	goto L261
L276:
	;
	if base.Ui64(int64(4294967295)) < base.Ui64(v3203) {
		goto L274
	} else {
		goto L277
	}
L277:
	;
	goto L275
L278:
	;
	v3339 = v3246
	v3340 = v3247
	v3341 = v314 + int32(1)
	v3342 = v3328
	goto L27
L279:
	;
	if base.Ui32(int32(124)) < base.Ui32(v3249&int32(65535)) {
		v3328 = v3249
		goto L278
	} else {
		goto L284
	}
L280:
	;
	if v3270 == v3268 {
		goto L279
	} else {
		goto L282
	}
L282:
	;
	v3303 = v32 + int32(24) + v3270<<(uint(int32(2))%32)
	v3304 = *(*int32)(unsafe.Add(mBase, uint32(v3303)))
	v3305 = int32(1)
	v3306 = v3304 + v3305
	*(*int32)(unsafe.Add(mBase, uint32(v3303))) = v3306
	if v3306 == int32(0) {
		v3270 = v3270 + v3305
		goto L280
	} else {
		goto L283
	}
L283:
	;
	v3328 = v3249
	goto L278
L284:
	;
	v3321 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(24)+v3268<<(uint(int32(2))%32)))) = v3321
	v3324 = v3249 + v3321
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v3324)
	v3328 = v3324
	goto L278
L285:
	;
	m.G0 = v32 + int32(1536)
	return
L286:
	;
	v6149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+526)))
	v6151 = v6132 & int32(65535)
	switch v6151 {
	case 0:
		goto L674
	case 1:
		goto L671
	case 2:
		goto L673
	default:
		goto L672
	}
L287:
	;
	v6119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	v6132 = v6119
	goto L286
L288:
	;
	v5721 = v3362 & int32(31)
	if v5721 == int32(0) {
		goto L601
	} else {
		goto L602
	}
L289:
	;
	if v3751 == int32(0) {
		v5654 = v5578
		v5670 = v5594
		goto L594
	} else {
		goto L595
	}
L290:
	;
	v3837 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	if l3 != 0 {
		goto L328
	} else {
		goto L329
	}
L291:
	;
	if base.Ui32(v3362) < base.Ui32(int32(135)) {
		v3427 = v3362
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v3442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	if base.Ui32(v3427) <= base.Ui32(int32(12)) {
		v3722 = v3339
		v3728 = v3427
		v3737 = v3442
		goto L299
	} else {
		goto L300
	}
L293:
	;
	v3381 = v3362
	goto L294
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+1036)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+1032)) = int32(_a2498)
	v3400 = *(*int64)(unsafe.Add(mBase, uint32(v32)+1032))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v3400
	v3406 = F_ffc_bigint_long_mul(m, v32+int32(24), v32+int32(8))
	mBase = m.M
	if v3406 == int32(0) {
		goto L287
	} else {
		goto L296
	}
L295:
	;
	v3427 = v3410
	goto L292
L296:
	;
	v3410 = v3381 + int32(-135)
	if base.Ui32(int32(134)) < base.Ui32(v3410) {
		v3381 = v3410
		goto L294
	} else {
		goto L297
	}
L297:
	;
	goto L295
L298:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v3468)
	goto L287
L299:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v3737)
	if v3728 == int32(0) {
		goto L288
	} else {
		goto L319
	}
L300:
	;
	v3453 = v3339
	v3454 = v3340
	v3459 = v3427
	v3468 = v3442
	goto L301
L301:
	;
	v3475 = v3468 & int32(65535)
	if v3475 != 0 {
		goto L304
	} else {
		goto L305
	}
L302:
	;
	v3722 = v3689
	v3728 = v3711
	v3737 = v3704
	goto L299
L303:
	;
	v3711 = v3459 + int32(-13)
	if base.Ui32(int32(12)) < base.Ui32(v3711) {
		v3453 = v3689
		v3454 = v3690
		v3459 = v3711
		v3468 = v3704
		goto L301
	} else {
		goto L318
	}
L304:
	;
	v3478 = v3475 & int32(3)
	v3480 = int32(0)
	if base.Ui32(v3475) < base.Ui32(int32(4)) {
		v3570 = int64(0)
		v3571 = v3454
		v3572 = v3480
		goto L306
	} else {
		goto L307
	}
L305:
	;
	v3689 = v3453
	v3690 = v3454
	v3704 = int32(0)
	goto L303
L306:
	;
	if v3478 == int32(0) {
		v3647 = v3570
		v3648 = v3571
		goto L311
	} else {
		goto L312
	}
L307:
	;
	v3487 = int32(0)
	v3497 = int64(0)
	v3499 = v3487
	v3508 = v3487
	goto L308
L308:
	;
	v3522 = v32 + int32(24) + v3499<<(uint(int32(2))%32)
	v3523 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3522))))
	v3524 = int64(1220703125)
	v3526 = v3523*v3524 + v3497
	*(*uint32)(unsafe.Add(mBase, uint32(v3522))) = uint32(v3526)
	v3528 = int32(4)
	v3529 = v3522 + v3528
	v3530 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3529))))
	v3533 = int64(32)
	v3535 = v3530*v3524 + int64(base.Ui64(v3526)>>(uint(v3533)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3529))) = uint32(v3535)
	v3538 = v3522 + int32(8)
	v3539 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3538))))
	v3544 = v3539*v3524 + int64(base.Ui64(v3535)>>(uint(v3533)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3538))) = uint32(v3544)
	v3547 = v3522 + int32(12)
	v3548 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3547))))
	v3553 = v3548*v3524 + int64(base.Ui64(v3544)>>(uint(v3533)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3547))) = uint32(v3553)
	v3556 = int64(base.Ui64(v3553) >> (uint(v3533) % 64))
	v3558 = v3499 + v3528
	v3560 = v3508 + v3528
	if v3560 != v3475&int32(65532) {
		v3497 = v3556
		v3499 = v3558
		v3508 = v3560
		goto L308
	} else {
		goto L310
	}
L309:
	;
	v3570 = v3556
	v3571 = v3553
	v3572 = v3558
	goto L306
L310:
	;
	goto L309
L311:
	;
	if base.Ui64(v3648) < base.Ui64(int64(4294967296)) {
		v3689 = v3647
		v3690 = v3648
		v3704 = v3468
		goto L303
	} else {
		goto L316
	}
L312:
	;
	v3601 = v3570
	v3603 = v3572
	v3605 = v3480
	goto L313
L313:
	;
	v3626 = v32 + int32(24) + v3603<<(uint(int32(2))%32)
	v3627 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3626))))
	v3630 = v3627*int64(1220703125) + v3601
	*(*uint32)(unsafe.Add(mBase, uint32(v3626))) = uint32(v3630)
	v3632 = int32(1)
	v3635 = int64(base.Ui64(v3630) >> (uint(int64(32)) % 64))
	v3637 = v3605 + v3632
	if v3637 != v3478 {
		v3601 = v3635
		v3603 = v3603 + v3632
		v3605 = v3637
		goto L313
	} else {
		goto L315
	}
L314:
	;
	v3647 = v3635
	v3648 = v3630
	goto L311
L315:
	;
	goto L314
L316:
	;
	if base.Ui32(int32(124)) < base.Ui32(v3475) {
		goto L298
	} else {
		goto L317
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(24)+v3475<<(uint(int32(2))%32)))) = base.I32_wrap_i64(v3647)
	v3689 = v3647
	v3690 = v3648
	v3704 = v3468 + int32(1)
	goto L303
L318:
	;
	goto L302
L319:
	;
	v3747 = v3737 & int32(65535)
	if v3747 == int32(0) {
		goto L288
	} else {
		goto L320
	}
L320:
	;
	v3750 = int32(3)
	v3751 = v3747 & v3750
	v3756 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3728<<(uint(v3750)%32))+uint32(_consts[922]))))
	if base.Ui32(int32(4)) <= base.Ui32(v3747) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v3765 = int32(0)
	v3777 = v3765
	v3786 = v3765
	v3791 = int64(0)
	goto L323
L322:
	;
	v5578 = v3722
	v5580 = int32(0)
	v5594 = int64(0)
	goto L289
L323:
	;
	v3800 = v32 + int32(24) + v3777<<(uint(int32(2))%32)
	v3801 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3800))))
	v3803 = v3756*v3801 + v3791
	*(*uint32)(unsafe.Add(mBase, uint32(v3800))) = uint32(v3803)
	v3805 = int32(4)
	v3806 = v3800 + v3805
	v3807 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3806))))
	v3809 = int64(32)
	v3811 = v3756*v3807 + int64(base.Ui64(v3803)>>(uint(v3809)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3806))) = uint32(v3811)
	v3814 = v3800 + int32(8)
	v3815 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3814))))
	v3819 = v3756*v3815 + int64(base.Ui64(v3811)>>(uint(v3809)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3814))) = uint32(v3819)
	v3822 = v3800 + int32(12)
	v3823 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3822))))
	v3827 = v3756*v3823 + int64(base.Ui64(v3819)>>(uint(v3809)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v3822))) = uint32(v3827)
	v3830 = int64(base.Ui64(v3827) >> (uint(v3809) % 64))
	v3832 = v3777 + v3805
	v3834 = v3786 + v3805
	if v3834 != v3747&int32(65532) {
		v3777 = v3832
		v3786 = v3834
		v3791 = v3830
		goto L323
	} else {
		goto L325
	}
L325:
	;
	v5578 = v3827
	v5580 = v3832
	v5594 = v3830
	goto L289
L326:
	;
	if l3 != 0 {
		goto L353
	} else {
		goto L354
	}
L327:
	;
	if l3 != 0 {
		goto L338
	} else {
		goto L339
	}
L328:
	;
	v3840 = int32(52)
	goto L330
L329:
	;
	v3840 = int32(23)
	goto L330
L330:
	;
	v3842 = v3840 ^ int32(63)
	v3845 = base.B2i32(int32(-32768)-v34 < v3842)
	if int32(-32768)-v34 < v3842 {
		goto L327
	} else {
		goto L331
	}
L331:
	;
	v3848 = int32(-32767) - v34
	v3849 = int32(64)
	if v3848 < v3849 {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v3852 = v3848
	goto L334
L333:
	;
	v3852 = v3849
	goto L334
L334:
	;
	if int32(63) < v3848 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v3857 = int64(0)
	goto L337
L336:
	;
	v3857 = int64(base.Ui64(v3837) >> (uint(base.I64_extend_i32_u(v3852)) % 64))
	goto L337
L337:
	;
	v3885 = base.B2i32(int64(base.Ui64(v3857)>>(uint(base.I64_extend_i32_u(v3840))%64)) != int64(0))
	v3887 = v3857
	goto L326
L338:
	;
	v3865 = int64(52)
	goto L340
L339:
	;
	v3865 = int64(23)
	goto L340
L340:
	;
	v3868 = int64(base.Ui64(v3837) >> (uint(base.I64_extend_i32_u(v3842)) % 64))
	v3871 = base.B2i32(base.Ui64(int64(2)<<(uint(v3865)%64)) <= base.Ui64(v3868))
	if base.Ui64(int64(2)<<(uint(v3865)%64)) <= base.Ui64(v3868) {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v3872 = int64(1) << (uint(v3865) % 64)
	goto L343
L342:
	;
	v3872 = v3868
	goto L343
L343:
	;
	v3878 = v36 + v3842 + v3871
	if l3 != 0 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v3881 = int32(2047)
	goto L346
L345:
	;
	v3881 = int32(255)
	goto L346
L346:
	;
	v3882 = base.B2i32(v3878 < v3881)
	if v3878 < v3881 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v3883 = v3872 & base.I64_rotl(int64(-2), v3865)
	goto L349
L348:
	;
	v3883 = int64(0)
	goto L349
L349:
	;
	if v3878 < v3881 {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v3884 = v3878
	goto L352
L351:
	;
	v3884 = v3881
	goto L352
L352:
	;
	v3885 = v3884
	v3887 = v3883
	goto L326
L353:
	;
	v3892 = int32(-1075)
	goto L355
L354:
	;
	v3892 = int32(-150)
	goto L355
L355:
	;
	if l3 != 0 {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v3895 = int64(4503599627370496)
	goto L358
L357:
	;
	v3895 = int64(8388608)
	goto L358
L358:
	;
	if l3 != 0 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v3898 = int64(4503599627370495)
	goto L361
L360:
	;
	v3898 = int64(8388607)
	goto L361
L361:
	;
	if l3 != 0 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v3901 = int64(9218868437227405312)
	goto L364
L363:
	;
	v3901 = int64(2139095040)
	goto L364
L364:
	;
	if l3&int32(1) == int32(0) {
		goto L366
	} else {
		goto L367
	}
L365:
	;
	v3935 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v3937 = int64(base.Ui64(v3933) >> (uint(int64(31)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v32)+532)) = uint32(v3937)
	if base.Ui64(v3933) < base.Ui64(int64(2147483648)) {
		goto L372
	} else {
		goto L373
	}
L366:
	;
	v3921 = v3887 & int64(4294967295)
	v3922 = v3901 & v3921
	if v3922 != int64(0) {
		goto L370
	} else {
		goto L371
	}
L367:
	;
	v3910 = (base.I64_extend_i32_u(v3885)<<(uint(int64(52))%64) | v3887) & v3901
	if v3910 != int64(0) {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v3932 = base.I32_wrap_i64(int64(base.Ui64(v3910) >> (uint(base.I64_extend_i32_u(v3840)) % 64)))
	v3933 = v3887&v3898 | v3895
	v3934 = v3910
	goto L365
L369:
	;
	v3932 = int32(1)
	v3933 = v3887 & v3898
	v3934 = v3910
	goto L365
L370:
	;
	v3932 = base.I32_wrap_i64(int64(base.Ui64(v3922) >> (uint(base.I64_extend_i32_u(v3840)) % 64)))
	v3933 = v3898&v3921 | v3895
	v3934 = v3922
	goto L365
L371:
	;
	v3932 = int32(1)
	v3933 = v3898 & v3921
	v3934 = v3922
	goto L365
L372:
	;
	v3943 = int32(1)
	goto L374
L373:
	;
	v3943 = int32(2)
	goto L374
L374:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1028)) = uint16(v3943)
	v3946 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+528)) = base.I32_wrap_i64(v3933)<<(uint(v3946)%32) | v3946
	v3954 = v3892 + (v3362 ^ int32(-1)) + v3932
	v3956 = int32(0) - v3362
	if base.Ui32(v3956) < base.Ui32(int32(135)) {
		v4019 = v3956
		goto L376
	} else {
		goto L377
	}
L375:
	;
	if v3954 < int32(1) {
		goto L421
	} else {
		goto L422
	}
L376:
	;
	if base.Ui32(v4019) < base.Ui32(int32(13)) {
		v4321 = v4019
		goto L383
	} else {
		goto L384
	}
L377:
	;
	v3973 = v3956
	goto L378
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+1036)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+1032)) = int32(_a2498)
	v3992 = *(*int64)(unsafe.Add(mBase, uint32(v32)+1032))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v3992
	v3998 = F_ffc_bigint_long_mul(m, v32+int32(528), v32+int32(16))
	mBase = m.M
	if v3998 == int32(0) {
		goto L375
	} else {
		goto L380
	}
L379:
	;
	v4019 = v4002
	goto L376
L380:
	;
	v4002 = v3973 + int32(-135)
	if base.Ui32(int32(134)) < base.Ui32(v4002) {
		v3973 = v4002
		goto L378
	} else {
		goto L381
	}
L381:
	;
	goto L379
L382:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1028)) = uint16(v4567)
	goto L375
L383:
	;
	if v4321 == int32(0) {
		goto L375
	} else {
		goto L403
	}
L384:
	;
	v4036 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+1028)))
	v4046 = v3933
	v4051 = v4019
	v4060 = v4036
	goto L385
L385:
	;
	v4067 = v4060 & int32(65535)
	if v4067 != 0 {
		goto L388
	} else {
		goto L389
	}
L386:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1028)) = uint16(v4296)
	v4321 = v4303
	goto L383
L387:
	;
	v4303 = v4051 + int32(-13)
	if base.Ui32(int32(12)) < base.Ui32(v4303) {
		v4046 = v4282
		v4051 = v4303
		v4060 = v4296
		goto L385
	} else {
		goto L402
	}
L388:
	;
	v4070 = v4067 & int32(3)
	v4072 = int32(0)
	if base.Ui32(v4067) < base.Ui32(int32(4)) {
		v4162 = int64(0)
		v4163 = v4046
		v4164 = v4072
		goto L390
	} else {
		goto L391
	}
L389:
	;
	v4282 = v4046
	v4296 = int32(0)
	goto L387
L390:
	;
	if v4070 == int32(0) {
		v4239 = v4162
		v4240 = v4163
		goto L395
	} else {
		goto L396
	}
L391:
	;
	v4079 = int32(0)
	v4089 = int64(0)
	v4091 = v4079
	v4100 = v4079
	goto L392
L392:
	;
	v4114 = v32 + int32(528) + v4091<<(uint(int32(2))%32)
	v4115 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4114))))
	v4116 = int64(1220703125)
	v4118 = v4115*v4116 + v4089
	*(*uint32)(unsafe.Add(mBase, uint32(v4114))) = uint32(v4118)
	v4120 = int32(4)
	v4121 = v4114 + v4120
	v4122 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4121))))
	v4125 = int64(32)
	v4127 = v4122*v4116 + int64(base.Ui64(v4118)>>(uint(v4125)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v4121))) = uint32(v4127)
	v4130 = v4114 + int32(8)
	v4131 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4130))))
	v4136 = v4131*v4116 + int64(base.Ui64(v4127)>>(uint(v4125)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v4130))) = uint32(v4136)
	v4139 = v4114 + int32(12)
	v4140 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4139))))
	v4145 = v4140*v4116 + int64(base.Ui64(v4136)>>(uint(v4125)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v4139))) = uint32(v4145)
	v4148 = int64(base.Ui64(v4145) >> (uint(v4125) % 64))
	v4150 = v4091 + v4120
	v4152 = v4100 + v4120
	if v4152 != v4067&int32(65532) {
		v4089 = v4148
		v4091 = v4150
		v4100 = v4152
		goto L392
	} else {
		goto L394
	}
L393:
	;
	v4162 = v4148
	v4163 = v4145
	v4164 = v4150
	goto L390
L394:
	;
	goto L393
L395:
	;
	if base.Ui64(v4240) < base.Ui64(int64(4294967296)) {
		v4282 = v4240
		v4296 = v4060
		goto L387
	} else {
		goto L400
	}
L396:
	;
	v4193 = v4162
	v4195 = v4164
	v4197 = v4072
	goto L397
L397:
	;
	v4218 = v32 + int32(528) + v4195<<(uint(int32(2))%32)
	v4219 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4218))))
	v4222 = v4219*int64(1220703125) + v4193
	*(*uint32)(unsafe.Add(mBase, uint32(v4218))) = uint32(v4222)
	v4224 = int32(1)
	v4227 = int64(base.Ui64(v4222) >> (uint(int64(32)) % 64))
	v4229 = v4197 + v4224
	if v4229 != v4070 {
		v4193 = v4227
		v4195 = v4195 + v4224
		v4197 = v4229
		goto L397
	} else {
		goto L399
	}
L398:
	;
	v4239 = v4227
	v4240 = v4222
	goto L395
L399:
	;
	goto L398
L400:
	;
	if base.Ui32(int32(124)) < base.Ui32(v4067) {
		v4567 = v4060
		goto L382
	} else {
		goto L401
	}
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(528)+v4067<<(uint(int32(2))%32)))) = base.I32_wrap_i64(v4239)
	v4282 = v4240
	v4296 = v4060 + int32(1)
	goto L387
L402:
	;
	goto L386
L403:
	;
	v4338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+1028)))
	if v4338 == int32(0) {
		goto L375
	} else {
		goto L404
	}
L404:
	;
	v4341 = int32(3)
	v4342 = v4338 & v4341
	v4347 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4321<<(uint(v4341)%32))+uint32(_consts[922]))))
	if base.Ui32(int32(4)) <= base.Ui32(v4338) {
		goto L406
	} else {
		goto L407
	}
L405:
	;
	if v4342 == int32(0) {
		v4511 = v4435
		v4527 = v4451
		goto L411
	} else {
		goto L412
	}
L406:
	;
	v4356 = int32(0)
	v4366 = int64(0)
	v4368 = v4356
	v4377 = v4356
	goto L408
L407:
	;
	v4435 = int64(0)
	v4437 = int32(0)
	v4451 = v3934
	goto L405
L408:
	;
	v4391 = v32 + int32(528) + v4368<<(uint(int32(2))%32)
	v4392 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4391))))
	v4394 = v4347*v4392 + v4366
	*(*uint32)(unsafe.Add(mBase, uint32(v4391))) = uint32(v4394)
	v4396 = int32(4)
	v4397 = v4391 + v4396
	v4398 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4397))))
	v4400 = int64(32)
	v4402 = v4347*v4398 + int64(base.Ui64(v4394)>>(uint(v4400)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v4397))) = uint32(v4402)
	v4405 = v4391 + int32(8)
	v4406 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4405))))
	v4410 = v4347*v4406 + int64(base.Ui64(v4402)>>(uint(v4400)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v4405))) = uint32(v4410)
	v4413 = v4391 + int32(12)
	v4414 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4413))))
	v4418 = v4347*v4414 + int64(base.Ui64(v4410)>>(uint(v4400)%64))
	*(*uint32)(unsafe.Add(mBase, uint32(v4413))) = uint32(v4418)
	v4421 = int64(base.Ui64(v4418) >> (uint(v4400) % 64))
	v4423 = v4368 + v4396
	v4425 = v4377 + v4396
	if v4425 != v4338&int32(65532) {
		v4366 = v4421
		v4368 = v4423
		v4377 = v4425
		goto L408
	} else {
		goto L410
	}
L409:
	;
	v4435 = v4421
	v4437 = v4423
	v4451 = v4418
	goto L405
L410:
	;
	goto L409
L411:
	;
	if base.Ui32(int32(124)) < base.Ui32(v4338) {
		goto L375
	} else {
		goto L416
	}
L412:
	;
	v4466 = v4435
	v4468 = v4437
	v4470 = int32(0)
	goto L413
L413:
	;
	v4491 = v32 + int32(528) + v4468<<(uint(int32(2))%32)
	v4492 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4491))))
	v4494 = v4347*v4492 + v4466
	*(*uint32)(unsafe.Add(mBase, uint32(v4491))) = uint32(v4494)
	v4496 = int32(1)
	v4499 = int64(base.Ui64(v4494) >> (uint(int64(32)) % 64))
	v4501 = v4470 + v4496
	if v4501 != v4342 {
		v4466 = v4499
		v4468 = v4468 + v4496
		v4470 = v4501
		goto L413
	} else {
		goto L415
	}
L414:
	;
	v4511 = v4499
	v4527 = v4494
	goto L411
L415:
	;
	goto L414
L416:
	;
	if base.Ui64(v4527) <= base.Ui64(int64(4294967295)) {
		goto L375
	} else {
		goto L417
	}
L417:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v32+int32(528)+v4338<<(uint(int32(2))%32)))) = uint32(v4511)
	v4567 = v4338 + int32(1)
	goto L382
L418:
	;
	v5382 = int32(1)
	v5383 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	v5384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+1028)))
	if base.Ui32(v5384) < base.Ui32(v5383) {
		v5502 = int32(0)
		v5503 = v5382
		goto L558
	} else {
		goto L559
	}
L419:
	;
	if base.Ui32(v3954) < base.Ui32(int32(32)) {
		goto L418
	} else {
		goto L513
	}
L420:
	;
	if v4609&v4614 == int32(0) {
		v5133 = v5098
		goto L506
	} else {
		goto L507
	}
L421:
	;
	if int32(-1) < v3954 {
		goto L418
	} else {
		goto L436
	}
L422:
	;
	v4606 = v3954 & int32(31)
	if v4606 == int32(0) {
		goto L419
	} else {
		goto L423
	}
L423:
	;
	v4609 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+1028)))
	if v4609 == int32(0) {
		goto L419
	} else {
		goto L424
	}
L424:
	;
	v4613 = int32(32) - v4606
	v4614 = int32(1)
	if v4609 != v4614 {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	v4622 = int32(0)
	v4635 = v4622
	v4643 = v4622
	v4644 = v4622
	goto L427
L426:
	;
	v4618 = int32(0)
	v5089 = v4618
	v5098 = v4618
	goto L420
L427:
	;
	goto L431
L429:
	;
	v4664 = v32 + int32(528)
	v4666 = v4635 << (uint(int32(2)) % 32)
	v4669 = v32 + int32(1032)
	v4671 = *(*int32)(unsafe.Add(mBase, uint32(v4669+v4666)))
	*(*int32)(unsafe.Add(mBase, uint32(v4664+v4666))) = v4671<<(uint(v4606)%32) | int32(base.Ui32(v4644)>>(uint(v4613)%32))
	goto L434
L430:
	;
	goto L429
L431:
	;
	v4661 = F__emscripten_memcpy_bulkmem(m, v32+int32(1032), v32+int32(528), int32(504))
	mBase = m.M
	goto L430
L432:
	;
	v4688 = v4666 | int32(4)
	v4693 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(1032)+v4688)))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(528)+v4688))) = v4693<<(uint(v4606)%32) | int32(base.Ui32(v4671)>>(uint(v4613)%32))
	v4698 = int32(2)
	v4699 = v4635 + v4698
	v4701 = v4643 + v4698
	if v4701 != v4609&int32(65534) {
		v4635 = v4699
		v4643 = v4701
		v4644 = v4693
		goto L427
	} else {
		goto L435
	}
L433:
	;
	goto L432
L434:
	;
	v4683 = F__emscripten_memcpy_bulkmem(m, v4669, v4664, int32(504))
	mBase = m.M
	goto L433
L435:
	;
	v5089 = v4699
	v5098 = v4693
	goto L420
L436:
	;
	v4705 = int32(0)
	v4707 = v4705 - v3954
	v4709 = v4707 & int32(31)
	if v4709 != 0 {
		goto L438
	} else {
		goto L439
	}
L437:
	;
	if base.Ui32(v4707) < base.Ui32(int32(32)) {
		goto L418
	} else {
		goto L461
	}
L438:
	;
	v4711 = v3342 & int32(65535)
	if v4711 == int32(0) {
		v4884 = v4705
		goto L437
	} else {
		goto L440
	}
L439:
	;
	v4884 = v3342
	goto L437
L440:
	;
	v4715 = int32(32) - v4709
	v4716 = int32(1)
	if v4711 != v4716 {
		goto L442
	} else {
		goto L443
	}
L441:
	;
	if v4711&v4716 == int32(0) {
		v4859 = v4824
		goto L453
	} else {
		goto L454
	}
L442:
	;
	v4724 = int32(0)
	v4737 = v4724
	v4745 = v4724
	v4746 = v4724
	goto L444
L443:
	;
	v4720 = int32(0)
	v4815 = v4720
	v4824 = v4720
	goto L441
L444:
	;
	goto L448
L445:
	;
	v4815 = v4801
	v4824 = v4795
	goto L441
L446:
	;
	v4766 = v32 + int32(24)
	v4768 = v4737 << (uint(int32(2)) % 32)
	v4771 = v32 + int32(1032)
	v4773 = *(*int32)(unsafe.Add(mBase, uint32(v4771+v4768)))
	*(*int32)(unsafe.Add(mBase, uint32(v4766+v4768))) = v4773<<(uint(v4709)%32) | int32(base.Ui32(v4746)>>(uint(v4715)%32))
	goto L451
L447:
	;
	goto L446
L448:
	;
	v4763 = F__emscripten_memcpy_bulkmem(m, v32+int32(1032), v32+int32(24), int32(504))
	mBase = m.M
	goto L447
L449:
	;
	v4790 = v4768 | int32(4)
	v4795 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(1032)+v4790)))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(24)+v4790))) = v4795<<(uint(v4709)%32) | int32(base.Ui32(v4773)>>(uint(v4715)%32))
	v4800 = int32(2)
	v4801 = v4737 + v4800
	v4803 = v4745 + v4800
	if v4803 != v4711&int32(65534) {
		v4737 = v4801
		v4745 = v4803
		v4746 = v4795
		goto L444
	} else {
		goto L452
	}
L450:
	;
	goto L449
L451:
	;
	v4785 = F__emscripten_memcpy_bulkmem(m, v4771, v4766, int32(504))
	mBase = m.M
	goto L450
L452:
	;
	goto L445
L453:
	;
	v4860 = int32(base.Ui32(v4859) >> (uint(v4715) % 32))
	if v4860 != 0 {
		goto L458
	} else {
		goto L459
	}
L454:
	;
	goto L457
L455:
	;
	v4848 = v4815 << (uint(int32(2)) % 32)
	v4854 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(1032)+v4848)))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(24)+v4848))) = int32(base.Ui32(v4824)>>(uint(v4715)%32)) | v4854<<(uint(v4709)%32)
	v4859 = v4854
	goto L453
L456:
	;
	goto L455
L457:
	;
	v4843 = F__emscripten_memcpy_bulkmem(m, v32+int32(1032), v32+int32(24), int32(504))
	mBase = m.M
	goto L456
L458:
	;
	if base.Ui32(int32(124)) < base.Ui32(v3342&int32(65535)) {
		goto L418
	} else {
		goto L460
	}
L459:
	;
	v4884 = v3342
	goto L437
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(24)+v4711<<(uint(int32(2))%32)))) = v4860
	v4872 = v3342 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v4872)
	v4884 = v4872
	goto L437
L461:
	;
	v4906 = v4884 & int32(65535)
	if v4906 == int32(0) {
		goto L418
	} else {
		goto L462
	}
L462:
	;
	v4910 = int32(base.Ui32(v4707) >> (uint(int32(5)) % 32))
	if base.Ui32(int32(125)) < base.Ui32(v4910+v4906) {
		goto L418
	} else {
		goto L463
	}
L463:
	;
	v4915 = v32 + int32(24)
	v4916 = int32(2)
	v4917 = v4910 << (uint(v4916) % 32)
	v4918 = v4915 + v4917
	v4922 = v4906 << (uint(v4916) % 32)
	if v4918 == v4915 {
		goto L465
	} else {
		goto L466
	}
L464:
	;
	v5075 = F__emscripten_memset_bulkmem(m, v32+int32(24), base.I32_extend8_s(int32(0)), v4917)
	mBase = m.M
	goto L505
L465:
	;
	goto L464
L466:
	;
	v4926 = v4922 + v4918
	if base.Ui32(int32(0)-v4922<<(uint(int32(1))%32)) < base.Ui32(v4915-v4926) {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	v4936 = (v4915 ^ v4918) & int32(3)
	if base.Ui32(v4915) <= base.Ui32(v4918) {
		goto L471
	} else {
		goto L472
	}
L468:
	;
	v4933 = F___memcpy(m, v4918, v4915, v4922)
	mBase = m.M
	goto L464
L469:
	;
	if v5042 == int32(0) {
		goto L465
	} else {
		goto L501
	}
L470:
	;
	if base.Ui32(v5020) <= base.Ui32(int32(3)) {
		v5041 = v5019
		v5042 = v5020
		v5043 = v5021
		goto L469
	} else {
		goto L497
	}
L471:
	;
	if v4936 != 0 {
		v5002 = v4922
		goto L481
	} else {
		goto L482
	}
L472:
	;
	if v4936 == int32(0) {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	if v4918&int32(3) != 0 {
		goto L475
	} else {
		goto L476
	}
L474:
	;
	v5041 = v4915
	v5042 = v4922
	v5043 = v4918
	goto L469
L475:
	;
	v4943 = v4915
	v4944 = v4922
	v4945 = v4918
	goto L477
L476:
	;
	v5019 = v4915
	v5020 = v4922
	v5021 = v4918
	goto L470
L477:
	;
	if v4944 == int32(0) {
		goto L465
	} else {
		goto L479
	}
L479:
	;
	v4949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4943))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4945))) = uint8(v4949)
	v4951 = int32(1)
	v4952 = v4943 + v4951
	v4954 = v4944 + int32(-1)
	v4956 = v4945 + v4951
	if v4956&int32(3) == int32(0) {
		v5019 = v4952
		v5020 = v4954
		v5021 = v4956
		goto L470
	} else {
		goto L480
	}
L480:
	;
	v4943 = v4952
	v4944 = v4954
	v4945 = v4956
	goto L477
L481:
	;
	if v5002 == int32(0) {
		goto L465
	} else {
		goto L493
	}
L482:
	;
	if v4926&int32(3) == int32(0) {
		v4982 = v4922
		goto L483
	} else {
		goto L484
	}
L483:
	;
	if base.Ui32(v4982) <= base.Ui32(int32(3)) {
		v5002 = v4982
		goto L481
	} else {
		goto L489
	}
L484:
	;
	v4967 = v4922
	goto L485
L485:
	;
	if v4967 == int32(0) {
		goto L465
	} else {
		goto L487
	}
L486:
	;
	v4982 = v4973
	goto L483
L487:
	;
	v4973 = v4967 + int32(-1)
	v4974 = v4918 + v4973
	v4976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4915+v4973))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4974))) = uint8(v4976)
	if v4974&int32(3) != 0 {
		v4967 = v4973
		goto L485
	} else {
		goto L488
	}
L488:
	;
	goto L486
L489:
	;
	v4989 = v4982
	goto L490
L490:
	;
	v4993 = v4989 + int32(-4)
	v4996 = *(*int32)(unsafe.Add(mBase, uint32(v4915+v4993)))
	*(*int32)(unsafe.Add(mBase, uint32(v4918+v4993))) = v4996
	if base.Ui32(int32(3)) < base.Ui32(v4993) {
		v4989 = v4993
		goto L490
	} else {
		goto L492
	}
L491:
	;
	v5002 = v4993
	goto L481
L492:
	;
	goto L491
L493:
	;
	v5009 = v5002
	goto L494
L494:
	;
	v5013 = v5009 + int32(-1)
	v5016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4915+v5013))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4918+v5013))) = uint8(v5016)
	if v5013 != 0 {
		v5009 = v5013
		goto L494
	} else {
		goto L496
	}
L496:
	;
	goto L465
L497:
	;
	v5026 = v5019
	v5027 = v5020
	v5028 = v5021
	goto L498
L498:
	;
	v5030 = *(*int32)(unsafe.Add(mBase, uint32(v5026)))
	*(*int32)(unsafe.Add(mBase, uint32(v5028))) = v5030
	v5032 = int32(4)
	v5033 = v5026 + v5032
	v5035 = v5028 + v5032
	v5037 = v5027 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v5037) {
		v5026 = v5033
		v5027 = v5037
		v5028 = v5035
		goto L498
	} else {
		goto L500
	}
L499:
	;
	v5041 = v5033
	v5042 = v5037
	v5043 = v5035
	goto L469
L500:
	;
	goto L499
L501:
	;
	v5048 = v5041
	v5049 = v5042
	v5050 = v5043
	goto L502
L502:
	;
	v5052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5048))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5050))) = uint8(v5052)
	v5054 = int32(1)
	v5059 = v5049 + int32(-1)
	if v5059 != 0 {
		v5048 = v5048 + v5054
		v5049 = v5059
		v5050 = v5050 + v5054
		goto L502
	} else {
		goto L504
	}
L503:
	;
	goto L465
L504:
	;
	goto L503
L505:
	;
	v5076 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	v5077 = v5076 + v4910
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v5077)
	goto L418
L506:
	;
	v5134 = int32(base.Ui32(v5133) >> (uint(v4613) % 32))
	if v5134 == int32(0) {
		goto L419
	} else {
		goto L511
	}
L507:
	;
	goto L510
L508:
	;
	v5122 = v5089 << (uint(int32(2)) % 32)
	v5128 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(1032)+v5122)))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(528)+v5122))) = int32(base.Ui32(v5098)>>(uint(v4613)%32)) | v5128<<(uint(v4606)%32)
	v5133 = v5128
	goto L506
L509:
	;
	goto L508
L510:
	;
	v5117 = F__emscripten_memcpy_bulkmem(m, v32+int32(1032), v32+int32(528), int32(504))
	mBase = m.M
	goto L509
L511:
	;
	if base.Ui32(int32(124)) < base.Ui32(v4609) {
		goto L418
	} else {
		goto L512
	}
L512:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(528)+v4609<<(uint(int32(2))%32)))) = v5134
	v5146 = v4609 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1028)) = uint16(v5146)
	goto L419
L513:
	;
	v5179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+1028)))
	if v5179 == int32(0) {
		goto L418
	} else {
		goto L514
	}
L514:
	;
	v5183 = int32(base.Ui32(v3954) >> (uint(int32(5)) % 32))
	if base.Ui32(int32(125)) < base.Ui32(v5183+v5179) {
		goto L418
	} else {
		goto L515
	}
L515:
	;
	v5188 = v32 + int32(528)
	v5189 = int32(2)
	v5190 = v5183 << (uint(v5189) % 32)
	v5191 = v5188 + v5190
	v5195 = v5179 << (uint(v5189) % 32)
	if v5191 == v5188 {
		goto L517
	} else {
		goto L518
	}
L516:
	;
	v5348 = F__emscripten_memset_bulkmem(m, v32+int32(528), base.I32_extend8_s(int32(0)), v5190)
	mBase = m.M
	goto L557
L517:
	;
	goto L516
L518:
	;
	v5199 = v5195 + v5191
	if base.Ui32(int32(0)-v5195<<(uint(int32(1))%32)) < base.Ui32(v5188-v5199) {
		goto L519
	} else {
		goto L520
	}
L519:
	;
	v5209 = (v5188 ^ v5191) & int32(3)
	if base.Ui32(v5188) <= base.Ui32(v5191) {
		goto L523
	} else {
		goto L524
	}
L520:
	;
	v5206 = F___memcpy(m, v5191, v5188, v5195)
	mBase = m.M
	goto L516
L521:
	;
	if v5315 == int32(0) {
		goto L517
	} else {
		goto L553
	}
L522:
	;
	if base.Ui32(v5293) <= base.Ui32(int32(3)) {
		v5314 = v5292
		v5315 = v5293
		v5316 = v5294
		goto L521
	} else {
		goto L549
	}
L523:
	;
	if v5209 != 0 {
		v5275 = v5195
		goto L533
	} else {
		goto L534
	}
L524:
	;
	if v5209 == int32(0) {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	if v5191&int32(3) != 0 {
		goto L527
	} else {
		goto L528
	}
L526:
	;
	v5314 = v5188
	v5315 = v5195
	v5316 = v5191
	goto L521
L527:
	;
	v5216 = v5188
	v5217 = v5195
	v5218 = v5191
	goto L529
L528:
	;
	v5292 = v5188
	v5293 = v5195
	v5294 = v5191
	goto L522
L529:
	;
	if v5217 == int32(0) {
		goto L517
	} else {
		goto L531
	}
L531:
	;
	v5222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5216))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5218))) = uint8(v5222)
	v5224 = int32(1)
	v5225 = v5216 + v5224
	v5227 = v5217 + int32(-1)
	v5229 = v5218 + v5224
	if v5229&int32(3) == int32(0) {
		v5292 = v5225
		v5293 = v5227
		v5294 = v5229
		goto L522
	} else {
		goto L532
	}
L532:
	;
	v5216 = v5225
	v5217 = v5227
	v5218 = v5229
	goto L529
L533:
	;
	if v5275 == int32(0) {
		goto L517
	} else {
		goto L545
	}
L534:
	;
	if v5199&int32(3) == int32(0) {
		v5255 = v5195
		goto L535
	} else {
		goto L536
	}
L535:
	;
	if base.Ui32(v5255) <= base.Ui32(int32(3)) {
		v5275 = v5255
		goto L533
	} else {
		goto L541
	}
L536:
	;
	v5240 = v5195
	goto L537
L537:
	;
	if v5240 == int32(0) {
		goto L517
	} else {
		goto L539
	}
L538:
	;
	v5255 = v5246
	goto L535
L539:
	;
	v5246 = v5240 + int32(-1)
	v5247 = v5191 + v5246
	v5249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5188+v5246))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5247))) = uint8(v5249)
	if v5247&int32(3) != 0 {
		v5240 = v5246
		goto L537
	} else {
		goto L540
	}
L540:
	;
	goto L538
L541:
	;
	v5262 = v5255
	goto L542
L542:
	;
	v5266 = v5262 + int32(-4)
	v5269 = *(*int32)(unsafe.Add(mBase, uint32(v5188+v5266)))
	*(*int32)(unsafe.Add(mBase, uint32(v5191+v5266))) = v5269
	if base.Ui32(int32(3)) < base.Ui32(v5266) {
		v5262 = v5266
		goto L542
	} else {
		goto L544
	}
L543:
	;
	v5275 = v5266
	goto L533
L544:
	;
	goto L543
L545:
	;
	v5282 = v5275
	goto L546
L546:
	;
	v5286 = v5282 + int32(-1)
	v5289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5188+v5286))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5191+v5286))) = uint8(v5289)
	if v5286 != 0 {
		v5282 = v5286
		goto L546
	} else {
		goto L548
	}
L548:
	;
	goto L517
L549:
	;
	v5299 = v5292
	v5300 = v5293
	v5301 = v5294
	goto L550
L550:
	;
	v5303 = *(*int32)(unsafe.Add(mBase, uint32(v5299)))
	*(*int32)(unsafe.Add(mBase, uint32(v5301))) = v5303
	v5305 = int32(4)
	v5306 = v5299 + v5305
	v5308 = v5301 + v5305
	v5310 = v5300 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v5310) {
		v5299 = v5306
		v5300 = v5310
		v5301 = v5308
		goto L550
	} else {
		goto L552
	}
L551:
	;
	v5314 = v5306
	v5315 = v5310
	v5316 = v5308
	goto L521
L552:
	;
	goto L551
L553:
	;
	v5321 = v5314
	v5322 = v5315
	v5323 = v5316
	goto L554
L554:
	;
	v5325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5321))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5323))) = uint8(v5325)
	v5327 = int32(1)
	v5332 = v5322 + int32(-1)
	if v5332 != 0 {
		v5321 = v5321 + v5327
		v5322 = v5332
		v5323 = v5323 + v5327
		goto L554
	} else {
		goto L556
	}
L555:
	;
	goto L517
L556:
	;
	goto L555
L557:
	;
	v5349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+1028)))
	v5350 = v5349 + v5183
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1028)) = uint16(v5350)
	goto L418
L558:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v3935
	if int32(-32768)-v34 < v3842 {
		goto L573
	} else {
		goto L574
	}
L559:
	;
	if base.Ui32(v5383) < base.Ui32(v5384) {
		v5472 = int32(0)
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v5502 = v5472
	v5503 = int32(0)
	goto L558
L561:
	;
	v5388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+526)))
	v5399 = v5383
	goto L562
L562:
	;
	v5419 = base.B2i32(v5399 == int32(0))
	if v5399 == int32(0) {
		v5472 = v5419
		goto L560
	} else {
		goto L564
	}
L563:
	;
	v5472 = v5419
	goto L560
L564:
	;
	goto L567
L565:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1534)) = uint16(v5388)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1532)) = uint16(v5383)
	v5434 = v32 + int32(1032)
	v5436 = v5399 + int32(-1)
	v5439 = v5434 + v5436<<(uint(int32(2))%32)
	v5440 = *(*int32)(unsafe.Add(mBase, uint32(v5439)))
	goto L570
L566:
	;
	goto L565
L567:
	;
	v5429 = F__emscripten_memcpy_bulkmem(m, v32+int32(1032), v32+int32(24), int32(500))
	mBase = m.M
	goto L566
L568:
	;
	v5450 = *(*int32)(unsafe.Add(mBase, uint32(v5439)))
	if base.Ui32(v5450) < base.Ui32(v5440) {
		v5502 = v5419
		v5503 = v5382
		goto L558
	} else {
		goto L571
	}
L569:
	;
	goto L568
L570:
	;
	v5448 = F__emscripten_memcpy_bulkmem(m, v5434, v32+int32(528), int32(504))
	mBase = m.M
	goto L569
L571:
	;
	if base.Ui32(v5450) <= base.Ui32(v5440) {
		v5399 = v5436
		goto L562
	} else {
		goto L572
	}
L572:
	;
	goto L563
L573:
	;
	v5538 = v36 + v3842
	v5542 = int64(base.Ui64(v3837) >> (uint(base.I64_extend_i32_u(v3842)) % 64))
	v5547 = v5542 + base.I64_extend_i32_u(v5503|v5502&base.I32_wrap_i64(v5542))
	if l3 != 0 {
		goto L581
	} else {
		goto L582
	}
L574:
	;
	v5515 = int32(-32767) - v34
	v5516 = int32(64)
	if v5515 < v5516 {
		goto L575
	} else {
		goto L576
	}
L575:
	;
	v5519 = v5515
	goto L577
L576:
	;
	v5519 = v5516
	goto L577
L577:
	;
	if int32(63) < v5515 {
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v5524 = int64(0)
	goto L580
L579:
	;
	v5524 = int64(base.Ui64(v3837) >> (uint(base.I64_extend_i32_u(v5519)) % 64))
	goto L580
L580:
	;
	v5531 = v5524 + base.I64_extend_i32_u(v5503|v5502&base.I32_wrap_i64(v5524))&int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v5531
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = base.B2i32(int64(base.Ui64(v5531)>>(uint(base.I64_extend_i32_u(v3840))%64)) != int64(0))
	goto L285
L581:
	;
	v5551 = int64(52)
	goto L583
L582:
	;
	v5551 = int64(23)
	goto L583
L583:
	;
	v5553 = base.B2i32(base.Ui64(v5547) < base.Ui64(int64(2)<<(uint(v5551)%64)))
	if base.Ui64(v5547) < base.Ui64(int64(2)<<(uint(v5551)%64)) {
		goto L584
	} else {
		goto L585
	}
L584:
	;
	v5554 = v5538
	goto L586
L585:
	;
	v5554 = v5538 + int32(1)
	goto L586
L586:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v5554
	if base.Ui64(v5547) < base.Ui64(int64(2)<<(uint(v5551)%64)) {
		goto L587
	} else {
		goto L588
	}
L587:
	;
	v5558 = v5547
	goto L589
L588:
	;
	v5558 = int64(1) << (uint(v5551) % 64)
	goto L589
L589:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v5558 & base.I64_rotl(int64(-2), v5551)
	if l3 != 0 {
		goto L590
	} else {
		goto L591
	}
L590:
	;
	v5565 = int32(2047)
	goto L592
L591:
	;
	v5565 = int32(255)
	goto L592
L592:
	;
	if v5554 < v5565 {
		goto L285
	} else {
		goto L593
	}
L593:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v5565
	goto L285
L594:
	;
	if base.Ui64(v5654) < base.Ui64(int64(4294967296)) {
		goto L288
	} else {
		goto L599
	}
L595:
	;
	v5611 = v5580
	v5613 = int32(0)
	v5625 = v5594
	goto L596
L596:
	;
	v5634 = v32 + int32(24) + v5611<<(uint(int32(2))%32)
	v5635 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v5634))))
	v5637 = v3756*v5635 + v5625
	*(*uint32)(unsafe.Add(mBase, uint32(v5634))) = uint32(v5637)
	v5639 = int32(1)
	v5642 = int64(base.Ui64(v5637) >> (uint(int64(32)) % 64))
	v5644 = v5613 + v5639
	if v5644 != v3751 {
		v5611 = v5611 + v5639
		v5613 = v5644
		v5625 = v5642
		goto L596
	} else {
		goto L598
	}
L597:
	;
	v5654 = v5637
	v5670 = v5642
	goto L594
L598:
	;
	goto L597
L599:
	;
	if base.Ui32(int32(124)) < base.Ui32(v3737&int32(65535)) {
		goto L287
	} else {
		goto L600
	}
L600:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(24)+v3747<<(uint(int32(2))%32)))) = base.I32_wrap_i64(v5670)
	v5689 = v3737 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v5689)
	goto L288
L601:
	;
	if base.Ui32(v3362) < base.Ui32(int32(32)) {
		goto L287
	} else {
		goto L623
	}
L602:
	;
	v5724 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	if v5724 == int32(0) {
		goto L601
	} else {
		goto L603
	}
L603:
	;
	v5728 = int32(32) - v5721
	v5729 = int32(1)
	if v5724 != v5729 {
		goto L605
	} else {
		goto L606
	}
L604:
	;
	if v5724&v5729 == int32(0) {
		v5872 = v5837
		goto L616
	} else {
		goto L617
	}
L605:
	;
	v5737 = int32(0)
	v5750 = v5737
	v5758 = v5737
	v5759 = v5737
	goto L607
L606:
	;
	v5733 = int32(0)
	v5828 = v5733
	v5837 = v5733
	goto L604
L607:
	;
	goto L611
L608:
	;
	v5828 = v5814
	v5837 = v5808
	goto L604
L609:
	;
	v5779 = v32 + int32(24)
	v5781 = v5750 << (uint(int32(2)) % 32)
	v5784 = v32 + int32(1032)
	v5786 = *(*int32)(unsafe.Add(mBase, uint32(v5784+v5781)))
	*(*int32)(unsafe.Add(mBase, uint32(v5779+v5781))) = v5786<<(uint(v5721)%32) | int32(base.Ui32(v5759)>>(uint(v5728)%32))
	goto L614
L610:
	;
	goto L609
L611:
	;
	v5776 = F__emscripten_memcpy_bulkmem(m, v32+int32(1032), v32+int32(24), int32(504))
	mBase = m.M
	goto L610
L612:
	;
	v5803 = v5781 | int32(4)
	v5808 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(1032)+v5803)))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(24)+v5803))) = v5808<<(uint(v5721)%32) | int32(base.Ui32(v5786)>>(uint(v5728)%32))
	v5813 = int32(2)
	v5814 = v5750 + v5813
	v5816 = v5758 + v5813
	if v5816 != v5724&int32(65534) {
		v5750 = v5814
		v5758 = v5816
		v5759 = v5808
		goto L607
	} else {
		goto L615
	}
L613:
	;
	goto L612
L614:
	;
	v5798 = F__emscripten_memcpy_bulkmem(m, v5784, v5779, int32(504))
	mBase = m.M
	goto L613
L615:
	;
	goto L608
L616:
	;
	v5873 = int32(base.Ui32(v5872) >> (uint(v5728) % 32))
	if v5873 == int32(0) {
		goto L601
	} else {
		goto L621
	}
L617:
	;
	goto L620
L618:
	;
	v5861 = v5828 << (uint(int32(2)) % 32)
	v5867 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(1032)+v5861)))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(24)+v5861))) = int32(base.Ui32(v5837)>>(uint(v5728)%32)) | v5867<<(uint(v5721)%32)
	v5872 = v5867
	goto L616
L619:
	;
	goto L618
L620:
	;
	v5856 = F__emscripten_memcpy_bulkmem(m, v32+int32(1032), v32+int32(24), int32(504))
	mBase = m.M
	goto L619
L621:
	;
	if base.Ui32(int32(124)) < base.Ui32(v5724) {
		goto L287
	} else {
		goto L622
	}
L622:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(24)+v5724<<(uint(int32(2))%32)))) = v5873
	v5885 = v5724 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)) = uint16(v5885)
	goto L601
L623:
	;
	v5918 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	if v5918 == int32(0) {
		goto L287
	} else {
		goto L624
	}
L624:
	;
	v5922 = int32(base.Ui32(v3362) >> (uint(int32(5)) % 32))
	if base.Ui32(int32(125)) < base.Ui32(v5922+v5918) {
		goto L287
	} else {
		goto L625
	}
L625:
	;
	v5927 = v32 + int32(24)
	v5928 = int32(2)
	v5929 = v5922 << (uint(v5928) % 32)
	v5930 = v5927 + v5929
	v5934 = v5918 << (uint(v5928) % 32)
	if v5930 == v5927 {
		goto L627
	} else {
		goto L628
	}
L626:
	;
	v6087 = F__emscripten_memset_bulkmem(m, v32+int32(24), base.I32_extend8_s(int32(0)), v5929)
	mBase = m.M
	goto L667
L627:
	;
	goto L626
L628:
	;
	v5938 = v5934 + v5930
	if base.Ui32(int32(0)-v5934<<(uint(int32(1))%32)) < base.Ui32(v5927-v5938) {
		goto L629
	} else {
		goto L630
	}
L629:
	;
	v5948 = (v5927 ^ v5930) & int32(3)
	if base.Ui32(v5927) <= base.Ui32(v5930) {
		goto L633
	} else {
		goto L634
	}
L630:
	;
	v5945 = F___memcpy(m, v5930, v5927, v5934)
	mBase = m.M
	goto L626
L631:
	;
	if v6054 == int32(0) {
		goto L627
	} else {
		goto L663
	}
L632:
	;
	if base.Ui32(v6032) <= base.Ui32(int32(3)) {
		v6053 = v6031
		v6054 = v6032
		v6055 = v6033
		goto L631
	} else {
		goto L659
	}
L633:
	;
	if v5948 != 0 {
		v6014 = v5934
		goto L643
	} else {
		goto L644
	}
L634:
	;
	if v5948 == int32(0) {
		goto L635
	} else {
		goto L636
	}
L635:
	;
	if v5930&int32(3) != 0 {
		goto L637
	} else {
		goto L638
	}
L636:
	;
	v6053 = v5927
	v6054 = v5934
	v6055 = v5930
	goto L631
L637:
	;
	v5955 = v5927
	v5956 = v5934
	v5957 = v5930
	goto L639
L638:
	;
	v6031 = v5927
	v6032 = v5934
	v6033 = v5930
	goto L632
L639:
	;
	if v5956 == int32(0) {
		goto L627
	} else {
		goto L641
	}
L641:
	;
	v5961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5955))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5957))) = uint8(v5961)
	v5963 = int32(1)
	v5964 = v5955 + v5963
	v5966 = v5956 + int32(-1)
	v5968 = v5957 + v5963
	if v5968&int32(3) == int32(0) {
		v6031 = v5964
		v6032 = v5966
		v6033 = v5968
		goto L632
	} else {
		goto L642
	}
L642:
	;
	v5955 = v5964
	v5956 = v5966
	v5957 = v5968
	goto L639
L643:
	;
	if v6014 == int32(0) {
		goto L627
	} else {
		goto L655
	}
L644:
	;
	if v5938&int32(3) == int32(0) {
		v5994 = v5934
		goto L645
	} else {
		goto L646
	}
L645:
	;
	if base.Ui32(v5994) <= base.Ui32(int32(3)) {
		v6014 = v5994
		goto L643
	} else {
		goto L651
	}
L646:
	;
	v5979 = v5934
	goto L647
L647:
	;
	if v5979 == int32(0) {
		goto L627
	} else {
		goto L649
	}
L648:
	;
	v5994 = v5985
	goto L645
L649:
	;
	v5985 = v5979 + int32(-1)
	v5986 = v5930 + v5985
	v5988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5927+v5985))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5986))) = uint8(v5988)
	if v5986&int32(3) != 0 {
		v5979 = v5985
		goto L647
	} else {
		goto L650
	}
L650:
	;
	goto L648
L651:
	;
	v6001 = v5994
	goto L652
L652:
	;
	v6005 = v6001 + int32(-4)
	v6008 = *(*int32)(unsafe.Add(mBase, uint32(v5927+v6005)))
	*(*int32)(unsafe.Add(mBase, uint32(v5930+v6005))) = v6008
	if base.Ui32(int32(3)) < base.Ui32(v6005) {
		v6001 = v6005
		goto L652
	} else {
		goto L654
	}
L653:
	;
	v6014 = v6005
	goto L643
L654:
	;
	goto L653
L655:
	;
	v6021 = v6014
	goto L656
L656:
	;
	v6025 = v6021 + int32(-1)
	v6028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5927+v6025))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5930+v6025))) = uint8(v6028)
	if v6025 != 0 {
		v6021 = v6025
		goto L656
	} else {
		goto L658
	}
L658:
	;
	goto L627
L659:
	;
	v6038 = v6031
	v6039 = v6032
	v6040 = v6033
	goto L660
L660:
	;
	v6042 = *(*int32)(unsafe.Add(mBase, uint32(v6038)))
	*(*int32)(unsafe.Add(mBase, uint32(v6040))) = v6042
	v6044 = int32(4)
	v6045 = v6038 + v6044
	v6047 = v6040 + v6044
	v6049 = v6039 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v6049) {
		v6038 = v6045
		v6039 = v6049
		v6040 = v6047
		goto L660
	} else {
		goto L662
	}
L661:
	;
	v6053 = v6045
	v6054 = v6049
	v6055 = v6047
	goto L631
L662:
	;
	goto L661
L663:
	;
	v6060 = v6053
	v6061 = v6054
	v6062 = v6055
	goto L664
L664:
	;
	v6064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6060))))
	*(*uint8)(unsafe.Add(mBase, uint32(v6062))) = uint8(v6064)
	v6066 = int32(1)
	v6071 = v6061 + int32(-1)
	if v6071 != 0 {
		v6060 = v6060 + v6066
		v6061 = v6071
		v6062 = v6062 + v6066
		goto L664
	} else {
		goto L666
	}
L665:
	;
	goto L627
L666:
	;
	goto L665
L667:
	;
	v6088 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+524)))
	v6132 = v6088 + v5922
	goto L286
L668:
	;
	if l3 != 0 {
		goto L703
	} else {
		goto L704
	}
L669:
	;
	goto L702
L670:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v6334
	goto L699
L671:
	;
	v6327 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v6334 = base.I64_extend_i32_u(v6327) << (uint(base.I64_extend_i32_u(base.I32_clz(v6327)+int32(32))) % 64)
	goto L670
L672:
	;
	goto L677
L673:
	;
	v6158 = *(*int64)(unsafe.Add(mBase, uint32(v32)+24))
	v6334 = v6158 << (uint(base.I64_clz(v6158)) % 64)
	goto L670
L674:
	;
	v6152 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v6152
	v6155 = int32(0)
	v6404 = int32(1)
	v6412 = v6152
	v6413 = v6155
	v6415 = v6155
	goto L668
L675:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1534)) = uint16(v6149)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1532)) = uint16(v6132)
	v6175 = v32 + int32(1032)
	v6176 = v6151<<(uint(int32(2))%32) + v6175
	v6179 = *(*int32)(unsafe.Add(mBase, uint32(v6176+int32(-4))))
	goto L680
L676:
	;
	goto L675
L677:
	;
	v6168 = F__emscripten_memcpy_bulkmem(m, v32+int32(1032), v32+int32(24), int32(500))
	mBase = m.M
	goto L676
L678:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1534)) = uint16(v6149)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1532)) = uint16(v6132)
	v6193 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v6176+int32(-8)))))
	goto L683
L679:
	;
	goto L678
L680:
	;
	v6187 = F__emscripten_memcpy_bulkmem(m, v6175, v32+int32(24), int32(500))
	mBase = m.M
	goto L679
L681:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1534)) = uint16(v6149)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1532)) = uint16(v6132)
	v6209 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v6176+int32(-12)))))
	v6210 = v6193<<(uint(int64(32))%64) | v6209
	v6211 = int32(32)
	v6212 = base.I32_clz(v6179)
	v6219 = base.I64_extend_i32_u(v6212 + v6211)
	v6221 = int64(base.Ui64(v6210)>>(uint(base.I64_extend_i32_u(v6211-v6212))%64)) | base.I64_extend_i32_u(v6179)<<(uint(v6219)%64)
	if base.Ui32(v6151) < base.Ui32(int32(4)) {
		v6299 = int32(1)
		goto L684
	} else {
		goto L685
	}
L682:
	;
	goto L681
L683:
	;
	v6201 = F__emscripten_memcpy_bulkmem(m, v32+int32(1032), v32+int32(24), int32(500))
	mBase = m.M
	goto L682
L684:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v6221
	goto L695
L685:
	;
	v6237 = int32(3)
	goto L686
L686:
	;
	goto L690
L687:
	;
	v6299 = v6277
	goto L684
L688:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1534)) = uint16(v6149)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1532)) = uint16(v6132)
	v6275 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(1032)+(v6237^int32(-1)+v6151)<<(uint(int32(2))%32))))
	v6277 = base.B2i32(v6275 == int32(0))
	if v6275 != 0 {
		v6299 = v6277
		goto L684
	} else {
		goto L691
	}
L689:
	;
	goto L688
L690:
	;
	v6263 = F__emscripten_memcpy_bulkmem(m, v32+int32(1032), v32+int32(24), int32(500))
	mBase = m.M
	goto L689
L691:
	;
	v6279 = v6237 + int32(1)
	if v6279 != v6151 {
		v6237 = v6279
		goto L686
	} else {
		goto L692
	}
L692:
	;
	goto L687
L693:
	;
	v6322 = base.B2i32(v6210<<(uint(v6219)%64) == int64(0)) & v6299
	if v6132&int32(65535) != 0 {
		v6348 = v6322
		v6356 = v6221
		goto L669
	} else {
		goto L696
	}
L694:
	;
	goto L693
L695:
	;
	v6318 = F__emscripten_memcpy_bulkmem(m, v32+int32(528), v32+int32(24), int32(500))
	mBase = m.M
	goto L694
L696:
	;
	v6325 = int32(0)
	v6404 = v6322
	v6412 = v6221
	v6413 = v6325
	v6415 = v6325
	goto L668
L697:
	;
	v6348 = int32(1)
	v6356 = v6334
	goto L669
L698:
	;
	goto L697
L699:
	;
	v6344 = F__emscripten_memcpy_bulkmem(m, v32+int32(528), v32+int32(24), int32(500))
	mBase = m.M
	goto L698
L700:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1534)) = uint16(v6149)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+1532)) = uint16(v6132)
	v6388 = v6132 & int32(65535)
	v6398 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v6388<<(uint(int32(2))%32)+(v32+int32(1032))+int32(-4)))))
	v6404 = v6348
	v6412 = v6356
	v6413 = v6388 << (uint(int32(5)) % 32)
	v6415 = base.I32_wrap_i64(base.I64_clz(v6398 << (uint(int64(32)) % 64)))
	goto L668
L701:
	;
	goto L700
L702:
	;
	v6383 = F__emscripten_memcpy_bulkmem(m, v32+int32(1032), v32+int32(528), int32(500))
	mBase = m.M
	goto L701
L703:
	;
	v6434 = int32(1011)
	goto L705
L704:
	;
	v6434 = int32(86)
	goto L705
L705:
	;
	if l3 != 0 {
		goto L706
	} else {
		goto L707
	}
L706:
	;
	v6439 = int32(11)
	goto L708
L707:
	;
	v6439 = int32(40)
	goto L708
L708:
	;
	if l3 != 0 {
		goto L709
	} else {
		goto L710
	}
L709:
	;
	v6443 = int64(11)
	goto L711
L710:
	;
	v6443 = int64(40)
	goto L711
L711:
	;
	v6444 = int64(base.Ui64(v6412) >> (uint(v6443) % 64))
	v6445 = int64(-1)
	v6449 = v6412 & (v6445<<(uint(v6443)%64) ^ v6445)
	v6450 = int64(1)
	v6451 = int32(-1)
	v6454 = v6450 << (uint(base.I64_extend_i32_u(v6439+v6451)) % 64)
	v6468 = v6444 + base.I64_extend_i32_u(base.B2i32(v6449 == v6454)&base.I32_wrap_i64(v6444)|(base.B2i32(base.Ui64(v6454) < base.Ui64(v6449))|(v6404|base.B2i32(v6449 != v6454)^v6451)))&v6450
	if l3 != 0 {
		goto L712
	} else {
		goto L713
	}
L712:
	;
	v6472 = int64(52)
	goto L714
L713:
	;
	v6472 = int64(23)
	goto L714
L714:
	;
	v6474 = base.B2i32(base.Ui64(int64(2)<<(uint(v6472)%64)) <= base.Ui64(v6468))
	v6475 = v6434 + v6413 - v6415 + v6439 + v6474
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v6475
	if base.Ui64(int64(2)<<(uint(v6472)%64)) <= base.Ui64(v6468) {
		goto L715
	} else {
		goto L716
	}
L715:
	;
	v6479 = int64(1) << (uint(v6472) % 64)
	goto L717
L716:
	;
	v6479 = v6468
	goto L717
L717:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v6479 & base.I64_rotl(int64(-2), v6472)
	if l3 != 0 {
		goto L718
	} else {
		goto L719
	}
L718:
	;
	v6486 = int32(2047)
	goto L720
L719:
	;
	v6486 = int32(255)
	goto L720
L720:
	;
	if base.Ui32(v6475) < base.Ui32(v6486) {
		goto L285
	} else {
		goto L721
	}
L721:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v6486
	goto L285
}
func F_fgets(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	if v9 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v237
L2:
	;
	v46 = l0
	v48 = l1 + int32(-1)
	goto L15
L3:
	;
	if l1 != int32(1) {
		v237 = int32(0)
		goto L1
	} else {
		goto L12
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v25 + int32(-1) | v25
	goto L10
L5:
	;
	v17 = int32(1)
	if v17 < l1 {
		v42 = v17
		goto L2
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	if l1 < int32(2) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v42 = int32(0)
	goto L2
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v20 + int32(-1) | v20
	goto L3
L10:
	;
	goto L11
L11:
	;
	goto L3
L12:
	;
	v38 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v38)
	return l0
L13:
	;
	if v42 != 0 {
		v237 = v230
		goto L1
	} else {
		goto L66
	}
L14:
	;
	if l0 != 0 {
		goto L64
	} else {
		goto L65
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v53 == v54 {
		v191 = v46
		v192 = v48
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v221 = v214
	goto L14
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v209)
	v214 = v207 + int32(1)
	if v209&int32(255) == int32(10) {
		v221 = v214
		goto L14
	} else {
		goto L62
	}
L18:
	;
	v196 = F___uflow(m, l2)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L57
	} else {
		goto L58
	}
L19:
	;
	v57 = v54 - v53
	v58 = int32(0)
	v61 = base.B2i32(v57 != v58)
	if v53&int32(3) == v58 {
		v87 = v53
		v89 = v57
		v90 = v61
		goto L25
	} else {
		goto L26
	}
L20:
	;
	if base.Ui32(v170) < base.Ui32(v48) {
		goto L48
	} else {
		goto L49
	}
L21:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v170 = v167 - v168
	v171 = v168
	goto L20
L22:
	;
	if v160 == int32(0) {
		goto L21
	} else {
		goto L47
	}
L23:
	;
	v160 = int32(0)
	goto L22
L24:
	;
	v138 = v131
	v140 = v133
	goto L42
L25:
	;
	if v90 == int32(0) {
		goto L23
	} else {
		goto L33
	}
L26:
	;
	if v57 == int32(0) {
		v87 = v53
		v89 = v57
		v90 = v61
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v70 = v53
	v72 = v57
	goto L28
L28:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v75 == int32(10) {
		v131 = v70
		v133 = v72
		goto L24
	} else {
		goto L30
	}
L29:
	;
	v87 = v82
	v89 = v78
	v90 = v80
	goto L25
L30:
	;
	v78 = v72 + int32(-1)
	v79 = int32(0)
	v80 = base.B2i32(v78 != v79)
	v82 = v70 + int32(1)
	if v82&int32(3) == v79 {
		v87 = v82
		v89 = v78
		v90 = v80
		goto L25
	} else {
		goto L31
	}
L31:
	;
	if v78 != 0 {
		v70 = v82
		v72 = v78
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v94 == int32(10) {
		v124 = v87
		v126 = v89
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v126 == int32(0) {
		goto L23
	} else {
		goto L41
	}
L35:
	;
	if base.Ui32(v89) < base.Ui32(int32(4)) {
		v124 = v87
		v126 = v89
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v104 = v87
	v106 = v89
	goto L37
L37:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v111 = v110 ^ int32(168430090)
	v114 = int32(-2139062144)
	if (int32(16843008)-v111|v111)&v114 != v114 {
		v131 = v104
		v133 = v106
		goto L24
	} else {
		goto L39
	}
L38:
	;
	v124 = v119
	v126 = v121
	goto L34
L39:
	;
	v119 = v104 + int32(4)
	v121 = v106 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v121) {
		v104 = v119
		v106 = v121
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v131 = v124
	v133 = v126
	goto L24
L42:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v143 != int32(10) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L23
L44:
	;
	v148 = v140 + int32(-1)
	if v148 != 0 {
		v138 = v138 + int32(1)
		v140 = v148
		goto L42
	} else {
		goto L46
	}
L45:
	;
	v160 = v138
	goto L22
L46:
	;
	goto L43
L47:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v170 = v160 - v163 + int32(1)
	v171 = v163
	goto L20
L48:
	;
	v173 = v170
	goto L50
L49:
	;
	v173 = v48
	goto L50
L50:
	;
	if v173 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v179 = v178 + v173
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v179
	v181 = v46 + v173
	if v160 != 0 {
		v221 = v181
		goto L14
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	v176 = F__emscripten_memcpy_bulkmem(m, v46, v171, v173)
	mBase = m.M
	goto L52
L54:
	;
	v182 = v48 - v173
	if v182 == int32(0) {
		v221 = v181
		goto L14
	} else {
		goto L55
	}
L55:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v179 == v185 {
		v191 = v181
		v192 = v182
		goto L18
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v179 + int32(1)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v207 = v181
	v208 = v182
	v209 = v190
	goto L17
L57:
	;
	return int32(0)
L58:
	;
	if int32(-1) < v196 {
		v207 = v191
		v208 = v192
		v209 = v196
		goto L17
	} else {
		goto L59
	}
L59:
	;
	v202 = int32(0)
	if v191 == l0 {
		v230 = v202
		goto L13
	} else {
		goto L60
	}
L60:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v204&int32(16) != 0 {
		v221 = v191
		goto L14
	} else {
		goto L61
	}
L61:
	;
	v230 = v202
	goto L13
L62:
	;
	v220 = v208 + int32(-1)
	if v220 != 0 {
		v46 = v214
		v48 = v220
		goto L15
	} else {
		goto L63
	}
L63:
	;
	goto L16
L64:
	;
	v227 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v221))) = uint8(v227)
	v230 = l0
	goto L13
L65:
	;
	v230 = int32(0)
	goto L13
L66:
	;
	goto L67
L67:
	;
	v237 = v230
	goto L1
}
func F_fillPercentileDistributionLatencies(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 float64
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 float64
	_ = v77
	var v85 float64
	_ = v85
	var v88 float64
	_ = v88
	var v91 int64
	_ = v91
	var v95 float64
	_ = v95
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v117 int64
	_ = v117
	var v121 int32
	_ = v121
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v165 int64
	_ = v165
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v186 int32
	_ = v186
	var v195 int64
	_ = v195
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	v6 = m.G0
	v8 = v6 - int32(176)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l1
	v15 = F_sdscatfmt(m, l0, int32(_a2277), v8+int32(32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	if v20 <= v19 {
		v222 = v15
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v229 = F_sdscatprintf(m, v222, int32(_a823), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L47
	}
L4:
	;
	v23 = v15
	v24 = int32(0)
	goto L5
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	v31 = v24 << (uint(int32(3)) % 32)
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v29+v31)))
	*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v33
	v36 = v8 + int32(48)
	v43 = F_snprintf(m, v36, int32(128), int32(_a546), v8+int32(16))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v222 = v217
	goto L3
L7:
	;
	v48 = F_strchr(m, v36, int32(46))
	mBase = m.M
	if v48 == int32(0) {
		v68 = v43
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	v77 = *(*float64)(unsafe.Add(mBase, uint32(v75+v31)))
	v85 = float64(100)
	if base.F64_lt(v77, v85) != 0 {
		goto L19
	} else {
		goto L20
	}
L9:
	;
	v72 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36+v68))) = uint8(v72)
	goto L8
L10:
	;
	v53 = v43
	v54 = v36 + v43
	goto L11
L11:
	;
	v57 = v54 + int32(-1)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v58 == int32(48) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v53 = v53 + int32(-1)
	v54 = v57
	goto L11
L14:
	;
	if v58 != int32(46) {
		v68 = v53
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v68 = v53 + int32(-1)
	goto L9
L16:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = base.F64_div(base.F64_convert_i64_s(v195), float64(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(48)
	v204 = F_sdscatprintf(m, v23, int32(_a2278), v8)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L42
	}
L17:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	if v104 < int32(1) {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v103 = int64(-9223372036854775807 - 1)
	goto L17
L19:
	;
	v88 = v77
	goto L21
L20:
	;
	v88 = v85
	goto L21
L21:
	;
	v91 = *(*int64)(unsafe.Add(mBase, uint32(l2)+88))
	v95 = base.F64_add(base.F64_mul(base.F64_div(v88, float64(100)), base.F64_convert_i64_s(v91)), float64(0.5))
	if base.F64_lt(base.F64_abs(v95), float64(9.223372036854776e+18)) == int32(0) {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v101 = base.I64_trunc_f64_s(v95)
	v103 = v101
	goto L17
L23:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v173 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
	v178 = int32(63) - (v172 + base.I32_wrap_i64(base.I64_clz(v173|v165)))
	v179 = base.I64_extend_i32_u(v178)
	v180 = v165 >> (uint(v179) % 64)
	v182 = base.I64_extend32_s(v180) << (uint(v179) % 64)
	if base.F64_eq(v77, float64(0)) != 0 {
		v195 = v182
		goto L40
	} else {
		goto L41
	}
L24:
	;
	v165 = int64(0)
	goto L23
L25:
	;
	v107 = int64(1)
	if v107 < v103 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v110 = v103
	goto L28
L27:
	;
	v110 = v107
	goto L28
L28:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
	v117 = int64(0)
	v121 = int32(0)
	goto L29
L29:
	;
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v111+v121<<(uint(int32(3))%32))))
	v127 = v126 + v117
	if v127 < v110 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L24
L31:
	;
	v150 = v121 + int32(1)
	if v150 != v104 {
		v117 = v127
		v121 = v150
		goto L29
	} else {
		goto L39
	}
L32:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v135 = int32(base.Ui32(v121) >> (uint(v134) % 32))
	if v135 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v136 = v129
	goto L35
L34:
	;
	v136 = int32(0)
	goto L35
L35:
	;
	v139 = int32(1)
	if v139 < v135 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v142 = v135
	goto L38
L37:
	;
	v142 = v139
	goto L38
L38:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v165 = base.I64_extend_i32_s((v129+int32(2147483647))&v121+v136) << (uint(base.I64_extend_i32_u(v142+v143+int32(-1))) % 64)
	goto L23
L39:
	;
	goto L30
L40:
	;
	goto L16
L41:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v195 = v182 + int64(1)<<(uint(base.I64_extend_i32_u(v178+base.B2i32(v186 <= base.I32_wrap_i64(v180))))%64) + int64(-1)
	goto L40
L42:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	if v24 == v207+int32(-1) {
		v217 = v204
		v218 = v207
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v220 = v24 + int32(1)
	if v220 < v218 {
		v23 = v217
		v24 = v220
		goto L5
	} else {
		goto L46
	}
L44:
	;
	v213 = F_sdscatlen(m, v204, int32(_a15), int32(1))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	v217 = v213
	v218 = v216
	goto L43
L46:
	;
	goto L6
L47:
	;
	m.G0 = v8 + int32(176)
	return v229
}
func F_findBucket_2(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int64
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	v8 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(304)
	m.G0 = v12
	v14 = int64(56)
	v16 = int64(65280)
	v18 = int64(40)
	v21 = int64(16711680)
	v23 = int64(24)
	v25 = int64(4278190080)
	v27 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = l1<<(uint(v14)%64) | l1&v16<<(uint(v18)%64) | (l1&v21<<(uint(v23)%64) | l1&v25<<(uint(v27)%64)) | (int64(base.Ui64(l1)>>(uint(v27)%64))&v25 | int64(base.Ui64(l1)>>(uint(v23)%64))&v21 | (int64(base.Ui64(l1)>>(uint(v18)%64))&v16 | int64(base.Ui64(l1)>>(uint(v14)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+12)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v12)+296)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v12)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v12 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v12 + int32(168)
	if l1 == int64(9223372036854775807) {
		v73 = int32(_a2517)
	} else {
		v73 = int32(_a104)
	}
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v75 = F_raxSeek(m, v12, v73, l2, v74)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		return int32(0)
	} else {
		v79 = int32(-1)
		v80 = F_raxNext(m, v12)
		mBase = m.M
		v81 = m.ExcPending
		if v81 != 0 {
			return int32(0)
		} else {
			if v80 == int32(0) {
				v144 = v79
				F_raxStop(m, v12)
				mBase = m.M
				v148 = m.ExcPending
				if v148 != 0 {
					return int32(0)
				} else {
					m.G0 = v12 + int32(304)
					return v144
				}
			} else {
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				v85 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
				v86 = int64(56)
				v88 = int64(65280)
				v90 = int64(40)
				v93 = int64(16711680)
				v95 = int64(24)
				v97 = int64(4278190080)
				v99 = int64(8)
				v120 = v85<<(uint(v86)%64) | v85&v88<<(uint(v90)%64) | (v85&v93<<(uint(v95)%64) | v85&v97<<(uint(v99)%64)) | (int64(base.Ui64(v85)>>(uint(v99)%64))&v97 | int64(base.Ui64(v85)>>(uint(v95)%64))&v93 | (int64(base.Ui64(v85)>>(uint(v90)%64))&v88 | int64(base.Ui64(v85)>>(uint(v86)%64))))
				v122 = l1 & int64(-8192)
				if v122 == int64(9223372036854767616) {
					v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
					v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
					if v129&int32(1) == int32(0) {
						F__serverAssert(m, int32(_a2518), int32(_a2500), int32(1106))
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v134 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l5))) = v128
						v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
						if v136 != int32(8) {
							F__serverAssert(m, int32(_a2519), int32(_a2500), int32(1109))
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v139 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
							*(*int64)(unsafe.Add(mBase, uint32(l2))) = v139
							*(*int64)(unsafe.Add(mBase, uint32(l4))) = v120
							v144 = v134
							F_raxStop(m, v12)
							mBase = m.M
							v148 = m.ExcPending
							if v148 != 0 {
								return int32(0)
							} else {
								m.G0 = v12 + int32(304)
								return v144
							}
						}
					}
				} else {
					if v122+int64(8192) < v120 {
						v144 = v79
						F_raxStop(m, v12)
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return int32(0)
						} else {
							m.G0 = v12 + int32(304)
							return v144
						}
					} else {
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
						v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
						if v129&int32(1) == int32(0) {
							F__serverAssert(m, int32(_a2518), int32(_a2500), int32(1106))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v134 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v128
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
							if v136 != int32(8) {
								F__serverAssert(m, int32(_a2519), int32(_a2500), int32(1109))
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v139 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
								*(*int64)(unsafe.Add(mBase, uint32(l2))) = v139
								*(*int64)(unsafe.Add(mBase, uint32(l4))) = v120
								v144 = v134
								F_raxStop(m, v12)
								mBase = m.M
								v148 = m.ExcPending
								if v148 != 0 {
									return int32(0)
								} else {
									m.G0 = v12 + int32(304)
									return v144
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_finishShutdown(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v74 int32
	_ = v74
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int64
	_ = v373
	var v378 int64
	_ = v378
	var v383 int64
	_ = v383
	var v388 int64
	_ = v388
	var v393 int64
	_ = v393
	var v398 int64
	_ = v398
	var v403 int64
	_ = v403
	var v406 int64
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int64
	_ = v474
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	v1 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(144)
	m.G0 = v16
	v20 = *(*int32)(unsafe.Add(mBase, _consts[711]))
	v22 = *(*int32)(unsafe.Add(mBase, _consts[78]))
	v24 = v16 + int32(136)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v1
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v25
	goto L1
L1:
	;
	v32 = v16 + int32(136)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v34 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v171 = v20 & int32(8)
	if v20&int32(16) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L3:
	;
	if v34 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v34+base.B2i32(v37 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v43
	goto L4
L6:
	;
	v55 = v1
	v58 = v34
	v61 = int32(0)
	goto L7
L7:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+104))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)+64))
	v71 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	if v69 == v71 {
		v122 = v61
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v145 {
		goto L2
	} else {
		goto L23
	}
L9:
	;
	v129 = v55 + int32(1)
	v131 = v16 + int32(136)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	if v133 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L10:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v74 != int32(9) {
		v83 = int32(0)
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v85 = v61 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v87 {
		v122 = v85
		goto L9
	} else {
		goto L13
	}
L12:
	;
	v78 = F___time(m, int32(0))
	mBase = m.M
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v67)+104))
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)+80))
	v83 = base.I32_wrap_i64(v78 - v80)
	goto L11
L13:
	;
	v90 = F_replicationGetReplicaName(m, v67)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v67)+104))
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v94)+64))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v99 = v97 + int32(-6)
	if base.Ui32(int32(5)) < base.Ui32(v99) {
		v107 = int32(_a320)
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v109 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(68)))) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(64)))) = v83
	*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = v109 - v95
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v90
	F__serverLog(m, int32(2), int32(_a2163), v16+int32(48))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v99<<(uint(int32(2))%32))+uint32(_consts[712])))
	v107 = v106
	goto L16
L18:
	;
	v122 = v85
	goto L9
L19:
	;
	if v133 != 0 {
		v55 = v129
		v58 = v133
		v61 = v122
		goto L7
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v133+base.B2i32(v136 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v142
	goto L20
L22:
	;
	goto L8
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v129 - v122
	F__serverLog(m, int32(2), int32(_a2164), v16+int32(32))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	goto L2
L25:
	;
	m.G0 = v16 + int32(144)
	return v553
L26:
	;
	v489 = int32(0)
	v490 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	if v490 == v489 {
		goto L118
	} else {
		goto L119
	}
L27:
	;
	v473 = int32(0)
	v474 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[713])) = v474
	*(*int64)(unsafe.Add(mBase, _consts[714])) = v474
	*(*int32)(unsafe.Add(mBase, _consts[715])) = v473
	F_replyToClientsBlockedOnShutdown(m)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L14
	} else {
		goto L116
	}
L28:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v464 {
		goto L27
	} else {
		goto L114
	}
L29:
	;
	v295 = int32(0)
	v296 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v296 == v295 {
		goto L72
	} else {
		goto L73
	}
L30:
	;
	F_killAppendOnlyChild(m)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L14
	} else {
		goto L71
	}
L31:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v283 {
		goto L30
	} else {
		goto L69
	}
L32:
	;
	v263 = int32(0)
	v264 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v171 == v263 {
		goto L63
	} else {
		goto L64
	}
L33:
	;
	if int32(3) < v195 {
		goto L27
	} else {
		goto L61
	}
L34:
	;
	F_scriptingEngineDebuggerKillForkedSessions(m)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L14
	} else {
		goto L45
	}
L35:
	;
	v174 = int32(0)
	v175 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v175 == v174 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	goto L37
L37:
	;
	v181 = int32(0)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+88)))
	if v183&int32(1) == v181 {
		v191 = v181
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v191 == int32(0) {
		goto L34
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v180)+2160))
	v191 = base.B2i32(v188 != int32(0))
	goto L39
L41:
	;
	v194 = int32(0)
	v195 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v171 == v194 {
		goto L33
	} else {
		goto L42
	}
L42:
	;
	if int32(3) < v195 {
		goto L34
	} else {
		goto L43
	}
L43:
	;
	F__serverLog(m, int32(3), int32(_a2165), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	goto L34
L45:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _consts[60]))
	if v209 != int32(1) {
		v230 = v209
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if v230 != int32(4) {
		v249 = v230
		goto L53
	} else {
		goto L54
	}
L47:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v213 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_killRDBChild(m)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L14
	} else {
		goto L51
	}
L49:
	;
	F__serverLog(m, int32(3), int32(_a2166), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v223 = int32(0)
	v224 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	F_rdbRemoveTempFile(m, v224, v223)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L14
	} else {
		goto L52
	}
L52:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _consts[60]))
	v230 = v229
	goto L46
L53:
	;
	if v249 != int32(2) {
		goto L29
	} else {
		goto L59
	}
L54:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v234 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v242 = int32(0)
	v243 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v245 = F_TerminateModuleForkChild(m, v243, v242)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L14
	} else {
		goto L58
	}
L56:
	;
	F__serverLog(m, int32(3), int32(_a2167), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L14
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _consts[60]))
	v249 = v248
	goto L53
L59:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v253 == int32(2) {
		goto L32
	} else {
		goto L60
	}
L60:
	;
	goto L31
L61:
	;
	F__serverLog(m, int32(3), int32(_a2168), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L14
	} else {
		goto L62
	}
L62:
	;
	goto L28
L63:
	;
	if int32(3) < v264 {
		goto L27
	} else {
		goto L67
	}
L64:
	;
	if int32(3) < v264 {
		goto L30
	} else {
		goto L65
	}
L65:
	;
	F__serverLog(m, int32(3), int32(_a2169), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L14
	} else {
		goto L66
	}
L66:
	;
	goto L31
L67:
	;
	F__serverLog(m, int32(3), int32(_a2170), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L14
	} else {
		goto L68
	}
L68:
	;
	goto L28
L69:
	;
	F__serverLog(m, int32(3), int32(_a2171), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L14
	} else {
		goto L70
	}
L70:
	;
	goto L30
L71:
	;
	goto L29
L72:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _consts[60]))
	if v333 != int32(5) {
		goto L83
	} else {
		goto L84
	}
L73:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v300 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_flushAppendOnlyFile(m, int32(1))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L14
	} else {
		goto L77
	}
L75:
	;
	F__serverLog(m, int32(2), int32(_a2172), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L14
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v313 = F_fsync(m, v312)
	mBase = m.M
	if v313 != int32(-1) {
		goto L72
	} else {
		goto L78
	}
L78:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v317 {
		goto L72
	} else {
		goto L79
	}
L79:
	;
	goto L80
L80:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v322 = F___strerror_l(m, v321, v321)
	mBase = m.M
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v322
	F__serverLog(m, int32(3), int32(_a2173), v16+int32(16))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L14
	} else {
		goto L82
	}
L82:
	;
	goto L72
L83:
	;
	v347 = int32(0)
	v348 = *(*int32)(unsafe.Add(mBase, _consts[59]))
	if base.B2i32(v347 < v348)&base.B2i32(v20&int32(2) == v347) != 0 {
		goto L89
	} else {
		goto L90
	}
L84:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v337 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	F_killSlotMigrationChild(m)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L14
	} else {
		goto L88
	}
L86:
	;
	F__serverLog(m, int32(3), int32(_a2174), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L14
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	goto L83
L89:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v359 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	if v20&int32(1) == int32(0) {
		goto L26
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v368 = v16 + int32(72)
	v369 = int32(0)
	v373 = *(*int64)(unsafe.Add(mBase, _consts[285]))
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(128)))) = v373
	v378 = *(*int64)(unsafe.Add(mBase, _consts[286]))
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(120)))) = v378
	v383 = *(*int64)(unsafe.Add(mBase, _consts[287]))
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(112)))) = v383
	v388 = *(*int64)(unsafe.Add(mBase, _consts[288]))
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(104)))) = v388
	v393 = *(*int64)(unsafe.Add(mBase, _consts[289]))
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(96)))) = v393
	v398 = *(*int64)(unsafe.Add(mBase, _consts[290]))
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(88)))) = v398
	v403 = *(*int64)(unsafe.Add(mBase, _consts[291]))
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(80)))) = v403
	v406 = *(*int64)(unsafe.Add(mBase, _consts[292]))
	*(*int64)(unsafe.Add(mBase, uint32(v368))) = v406
	v409 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v409 != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	F__serverLog(m, int32(2), int32(_a2175), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L14
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v436 = int32(0)
	v438 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	v440 = F_rdbSave(m, v436, v438, v435, int32(16))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L14
	} else {
		goto L106
	}
L96:
	;
	v422 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	if v422 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L97:
	;
	v411 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	if v411 == int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v416 = *(*int32)(unsafe.Add(mBase, _consts[294]))
	if v416 == int32(-1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v419 = int32(0)
	goto L101
L100:
	;
	v419 = v416
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368))) = v419
	v435 = v368
	goto L95
L102:
	;
	v429 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	if v429 != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v422)+96))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v368))) = v426
	v435 = v368
	goto L95
L104:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v429)+96))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v368))) = v432
	v435 = v368
	goto L95
L105:
	;
	v435 = int32(0)
	goto L95
L106:
	;
	if v440 == int32(0) {
		goto L26
	} else {
		goto L107
	}
L107:
	;
	v444 = int32(0)
	v445 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v171 == v444 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	if int32(3) < v445 {
		goto L28
	} else {
		goto L112
	}
L109:
	;
	if int32(3) < v445 {
		goto L26
	} else {
		goto L110
	}
L110:
	;
	F__serverLog(m, int32(3), int32(_a2176), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L14
	} else {
		goto L111
	}
L111:
	;
	goto L26
L112:
	;
	F__serverLog(m, int32(3), int32(_a2177), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L14
	} else {
		goto L113
	}
L113:
	;
	goto L28
L114:
	;
	F__serverLog(m, int32(3), int32(_a2178), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L14
	} else {
		goto L115
	}
L115:
	;
	goto L27
L116:
	;
	F_unpauseActions(m, int32(1))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L14
	} else {
		goto L117
	}
L117:
	;
	v553 = int32(-1)
	goto L25
L118:
	;
	v496 = int32(0)
	F_moduleFireServerEvent(m, int64(5), v496, v496)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L14
	} else {
		goto L121
	}
L119:
	;
	F_aofManifestFree(m, v490)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L14
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v500 = int32(0)
	v501 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v503 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	if v503 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v520 = int32(0)
	v522 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v522 == v520 {
		goto L129
	} else {
		goto L130
	}
L123:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v507 {
		v517 = v501
		goto L126
	} else {
		goto L127
	}
L124:
	;
	if v501 == int32(0) {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v518 = F_unlink(m, v517)
	mBase = m.M
	goto L122
L127:
	;
	F__serverLog(m, int32(2), int32(_a2179), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L14
	} else {
		goto L128
	}
L128:
	;
	v516 = *(*int32)(unsafe.Add(mBase, _consts[716]))
	v517 = v516
	goto L126
L129:
	;
	F_flushReplicasOutputBuffers(m)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L14
	} else {
		goto L132
	}
L130:
	;
	F_clusterHandleServerShutdown(m, base.B2i32(v20&int32(32) != int32(0)))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L14
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	F_closeListeningSockets(m, int32(1))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L14
	} else {
		goto L133
	}
L133:
	;
	F_moduleUnloadAllModules(m)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L14
	} else {
		goto L134
	}
L134:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v539 {
		v553 = v520
		goto L25
	} else {
		goto L135
	}
L135:
	;
	v542 = int32(0)
	v546 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	if v546 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v547 = int32(_a2180)
	goto L138
L137:
	;
	v547 = int32(_a256)
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v547
	F__serverLog(m, int32(3), int32(_a2181), v16)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L14
	} else {
		goto L139
	}
L139:
	;
	v553 = v542
	goto L25
}
func F_firePostExecutionUnitJobs(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int64
	_ = v166
	var v170 int64
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v186 int64
	_ = v186
	var v188 int32
	_ = v188
	var v192 int64
	_ = v192
	var v196 int64
	_ = v196
	var v199 int32
	_ = v199
	var v207 int64
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	v1 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v23 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	*(*int32)(unsafe.Add(mBase, _consts[177])) = v23 + int32(1)
	goto L2
L1:
	;
	v56 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, _consts[442]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	if v58 == v56 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L7:
	;
	v248 = int32(0)
	v250 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	*(*int32)(unsafe.Add(mBase, _consts[177])) = v250 + int32(-1)
	goto L35
L8:
	;
	v78 = v57
	goto L9
L9:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	F_listDelNode(m, v78, v89)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L7
L11:
	;
	return
L12:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v94 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(72)))) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(64)))) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(56)))) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(48)))) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(40)))) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(32)))) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(24)))) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(16)))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(561)
	v117 = int32(0)
	v118 = *(*int32)(unsafe.Add(mBase, _consts[383]))
	if v118 == v117 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v149
	v153 = *(*int32)(unsafe.Add(mBase, _consts[384]))
	v154 = int32(0)
	v155 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v156 = m.T0[v155].(func(*base.Module) int64)(m)
	mBase = m.M
	if v153 == v154 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v137 = F_createClient(m, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L11
	} else {
		goto L17
	}
L15:
	;
	v121 = int32(0)
	v123 = v118 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[383])) = v123
	v126 = *(*int32)(unsafe.Add(mBase, _consts[385]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126+v123<<(uint(int32(2))%32))))
	v132 = *(*int32)(unsafe.Add(mBase, _consts[386]))
	if base.Ui32(v132) <= base.Ui32(v123) {
		v149 = v130
		goto L13
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[386])) = v123
	v149 = v130
	goto L13
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+328)) = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+200)) = v141 | int32(1073741824)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v137)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+204)) = v145 | int32(268435456)
	v149 = v137
	goto L13
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v170
	v174 = int32(0)
	v178 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	*(*int32)(unsafe.Add(mBase, _consts[177])) = v178 + int32(1)
	goto L23
L19:
	;
	v166 = *(*int64)(unsafe.Add(mBase, _consts[387]))
	v170 = v166*int64(1000) + v156
	goto L18
L20:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	v162 = base.I32_div_s(int32(1000000), v161)
	v170 = v156 + base.I64_extend_i32_s(v162)
	goto L18
L21:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v213 = F_selectDb(m, v211, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L11
	} else {
		goto L27
	}
L22:
	;
	goto L21
L23:
	;
	if v178 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	goto L26
L25:
	;
	v188 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[178])) = v186
	v192 = base.I64_div_s(v186, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[54])) = v192
	v196 = base.I64_div_s(v186, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[109])) = v196
	v199 = *(*int32)(unsafe.Add(mBase, _consts[179]))
	F_lrulfu_updateClockAndPolicy(m, v192, int32(base.Ui32(v199&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v207 = *(*int64)(unsafe.Add(mBase, _consts[54]))
	*(*int64)(unsafe.Add(mBase, _consts[12])) = v207
	goto L22
L26:
	;
	v186 = F_ustime(m)
	mBase = m.M
	goto L25
L27:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	m.T0[v218].(func(*base.Module, int32, int32))(m, v15+int32(8), v217)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	if v221 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_moduleFreeContext(m, v15+int32(8))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L11
	} else {
		goto L32
	}
L30:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	m.T0[v221].(func(*base.Module, int32))(m, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L11
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	F_valkey_free(m, v90)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _consts[442]))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+20))
	if v235 != 0 {
		v78 = v234
		goto L9
	} else {
		goto L34
	}
L34:
	;
	goto L10
L35:
	;
	m.G0 = v15 + int32(80)
	return
}
func F_flagTransaction(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v3&int32(8) == int32(0) {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v3 | int32(4096)
		F_resetClientMultiState(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_flock(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
func F_flushAllDataAndResetRDB(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v42 int64
	_ = v42
	var v47 int64
	_ = v47
	var v52 int64
	_ = v52
	var v57 int64
	_ = v57
	var v62 int64
	_ = v62
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v3 = m.G0
	v5 = v3 - int32(64)
	m.G0 = v5
	v10 = F_emptyData(m, int32(-1), l0, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = int32(_a20)
		v13 = *(*int64)(unsafe.Add(mBase, _consts[180]))
		*(*int64)(unsafe.Add(mBase, _consts[180])) = v10 + v13
		v17 = *(*int32)(unsafe.Add(mBase, _consts[60]))
		if v17 != int32(1) {
			v24 = v17
			if v24 != int32(5) {
				v30 = *(*int32)(unsafe.Add(mBase, _consts[59]))
				if v30 < int32(1) {
					m.G0 = v5 + int32(64)
					return
				} else {
					v33 = int32(0)
					v37 = *(*int64)(unsafe.Add(mBase, _consts[285]))
					*(*int64)(unsafe.Add(mBase, uint32(v5+int32(56)))) = v37
					v42 = *(*int64)(unsafe.Add(mBase, _consts[286]))
					*(*int64)(unsafe.Add(mBase, uint32(v5+int32(48)))) = v42
					v47 = *(*int64)(unsafe.Add(mBase, _consts[287]))
					*(*int64)(unsafe.Add(mBase, uint32(v5+int32(40)))) = v47
					v52 = *(*int64)(unsafe.Add(mBase, _consts[288]))
					*(*int64)(unsafe.Add(mBase, uint32(v5+int32(32)))) = v52
					v57 = *(*int64)(unsafe.Add(mBase, _consts[289]))
					*(*int64)(unsafe.Add(mBase, uint32(v5+int32(24)))) = v57
					v62 = *(*int64)(unsafe.Add(mBase, _consts[290]))
					*(*int64)(unsafe.Add(mBase, uint32(v5+int32(16)))) = v62
					v67 = *(*int64)(unsafe.Add(mBase, _consts[291]))
					*(*int64)(unsafe.Add(mBase, uint32(v5+int32(8)))) = v67
					v70 = *(*int64)(unsafe.Add(mBase, _consts[292]))
					*(*int64)(unsafe.Add(mBase, uint32(v5))) = v70
					v73 = *(*int32)(unsafe.Add(mBase, _consts[166]))
					if v73 != 0 {
						v86 = *(*int32)(unsafe.Add(mBase, _consts[167]))
						if v86 == int32(0) {
							v93 = *(*int32)(unsafe.Add(mBase, _consts[168]))
							if v93 != 0 {
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+96))
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = v96
								v99 = v5
							} else {
								v99 = int32(0)
							}
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+96))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v5))) = v90
							v99 = v5
						}
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, _consts[293]))
						if v75 == int32(0) {
							v86 = *(*int32)(unsafe.Add(mBase, _consts[167]))
							if v86 == int32(0) {
								v93 = *(*int32)(unsafe.Add(mBase, _consts[168]))
								if v93 != 0 {
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+96))
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
									*(*int32)(unsafe.Add(mBase, uint32(v5))) = v96
									v99 = v5
								} else {
									v99 = int32(0)
								}
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+96))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = v90
								v99 = v5
							}
						} else {
							v80 = *(*int32)(unsafe.Add(mBase, _consts[294]))
							if v80 == int32(-1) {
								v83 = int32(0)
							} else {
								v83 = v80
							}
							*(*int32)(unsafe.Add(mBase, uint32(v5))) = v83
							v99 = v5
						}
					}
					v100 = int32(0)
					v102 = *(*int32)(unsafe.Add(mBase, _consts[46]))
					v104 = F_rdbSave(m, v100, v102, v99, v100)
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return
					} else {
						m.G0 = v5 + int32(64)
						return
					}
				}
			} else {
				F_killSlotMigrationChild(m)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, _consts[59]))
					if v30 < int32(1) {
						m.G0 = v5 + int32(64)
						return
					} else {
						v33 = int32(0)
						v37 = *(*int64)(unsafe.Add(mBase, _consts[285]))
						*(*int64)(unsafe.Add(mBase, uint32(v5+int32(56)))) = v37
						v42 = *(*int64)(unsafe.Add(mBase, _consts[286]))
						*(*int64)(unsafe.Add(mBase, uint32(v5+int32(48)))) = v42
						v47 = *(*int64)(unsafe.Add(mBase, _consts[287]))
						*(*int64)(unsafe.Add(mBase, uint32(v5+int32(40)))) = v47
						v52 = *(*int64)(unsafe.Add(mBase, _consts[288]))
						*(*int64)(unsafe.Add(mBase, uint32(v5+int32(32)))) = v52
						v57 = *(*int64)(unsafe.Add(mBase, _consts[289]))
						*(*int64)(unsafe.Add(mBase, uint32(v5+int32(24)))) = v57
						v62 = *(*int64)(unsafe.Add(mBase, _consts[290]))
						*(*int64)(unsafe.Add(mBase, uint32(v5+int32(16)))) = v62
						v67 = *(*int64)(unsafe.Add(mBase, _consts[291]))
						*(*int64)(unsafe.Add(mBase, uint32(v5+int32(8)))) = v67
						v70 = *(*int64)(unsafe.Add(mBase, _consts[292]))
						*(*int64)(unsafe.Add(mBase, uint32(v5))) = v70
						v73 = *(*int32)(unsafe.Add(mBase, _consts[166]))
						if v73 != 0 {
							v86 = *(*int32)(unsafe.Add(mBase, _consts[167]))
							if v86 == int32(0) {
								v93 = *(*int32)(unsafe.Add(mBase, _consts[168]))
								if v93 != 0 {
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+96))
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
									*(*int32)(unsafe.Add(mBase, uint32(v5))) = v96
									v99 = v5
								} else {
									v99 = int32(0)
								}
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+96))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = v90
								v99 = v5
							}
						} else {
							v75 = *(*int32)(unsafe.Add(mBase, _consts[293]))
							if v75 == int32(0) {
								v86 = *(*int32)(unsafe.Add(mBase, _consts[167]))
								if v86 == int32(0) {
									v93 = *(*int32)(unsafe.Add(mBase, _consts[168]))
									if v93 != 0 {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+96))
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
										*(*int32)(unsafe.Add(mBase, uint32(v5))) = v96
										v99 = v5
									} else {
										v99 = int32(0)
									}
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+96))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
									*(*int32)(unsafe.Add(mBase, uint32(v5))) = v90
									v99 = v5
								}
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, _consts[294]))
								if v80 == int32(-1) {
									v83 = int32(0)
								} else {
									v83 = v80
								}
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = v83
								v99 = v5
							}
						}
						v100 = int32(0)
						v102 = *(*int32)(unsafe.Add(mBase, _consts[46]))
						v104 = F_rdbSave(m, v100, v102, v99, v100)
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return
						} else {
							m.G0 = v5 + int32(64)
							return
						}
					}
				}
			}
		} else {
			F_killRDBChild(m)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, _consts[60]))
				v24 = v23
				if v24 != int32(5) {
					v30 = *(*int32)(unsafe.Add(mBase, _consts[59]))
					if v30 < int32(1) {
						m.G0 = v5 + int32(64)
						return
					} else {
						v33 = int32(0)
						v37 = *(*int64)(unsafe.Add(mBase, _consts[285]))
						*(*int64)(unsafe.Add(mBase, uint32(v5+int32(56)))) = v37
						v42 = *(*int64)(unsafe.Add(mBase, _consts[286]))
						*(*int64)(unsafe.Add(mBase, uint32(v5+int32(48)))) = v42
						v47 = *(*int64)(unsafe.Add(mBase, _consts[287]))
						*(*int64)(unsafe.Add(mBase, uint32(v5+int32(40)))) = v47
						v52 = *(*int64)(unsafe.Add(mBase, _consts[288]))
						*(*int64)(unsafe.Add(mBase, uint32(v5+int32(32)))) = v52
						v57 = *(*int64)(unsafe.Add(mBase, _consts[289]))
						*(*int64)(unsafe.Add(mBase, uint32(v5+int32(24)))) = v57
						v62 = *(*int64)(unsafe.Add(mBase, _consts[290]))
						*(*int64)(unsafe.Add(mBase, uint32(v5+int32(16)))) = v62
						v67 = *(*int64)(unsafe.Add(mBase, _consts[291]))
						*(*int64)(unsafe.Add(mBase, uint32(v5+int32(8)))) = v67
						v70 = *(*int64)(unsafe.Add(mBase, _consts[292]))
						*(*int64)(unsafe.Add(mBase, uint32(v5))) = v70
						v73 = *(*int32)(unsafe.Add(mBase, _consts[166]))
						if v73 != 0 {
							v86 = *(*int32)(unsafe.Add(mBase, _consts[167]))
							if v86 == int32(0) {
								v93 = *(*int32)(unsafe.Add(mBase, _consts[168]))
								if v93 != 0 {
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+96))
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
									*(*int32)(unsafe.Add(mBase, uint32(v5))) = v96
									v99 = v5
								} else {
									v99 = int32(0)
								}
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+96))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = v90
								v99 = v5
							}
						} else {
							v75 = *(*int32)(unsafe.Add(mBase, _consts[293]))
							if v75 == int32(0) {
								v86 = *(*int32)(unsafe.Add(mBase, _consts[167]))
								if v86 == int32(0) {
									v93 = *(*int32)(unsafe.Add(mBase, _consts[168]))
									if v93 != 0 {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+96))
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
										*(*int32)(unsafe.Add(mBase, uint32(v5))) = v96
										v99 = v5
									} else {
										v99 = int32(0)
									}
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+96))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
									*(*int32)(unsafe.Add(mBase, uint32(v5))) = v90
									v99 = v5
								}
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, _consts[294]))
								if v80 == int32(-1) {
									v83 = int32(0)
								} else {
									v83 = v80
								}
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = v83
								v99 = v5
							}
						}
						v100 = int32(0)
						v102 = *(*int32)(unsafe.Add(mBase, _consts[46]))
						v104 = F_rdbSave(m, v100, v102, v99, v100)
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return
						} else {
							m.G0 = v5 + int32(64)
							return
						}
					}
				} else {
					F_killSlotMigrationChild(m)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, _consts[59]))
						if v30 < int32(1) {
							m.G0 = v5 + int32(64)
							return
						} else {
							v33 = int32(0)
							v37 = *(*int64)(unsafe.Add(mBase, _consts[285]))
							*(*int64)(unsafe.Add(mBase, uint32(v5+int32(56)))) = v37
							v42 = *(*int64)(unsafe.Add(mBase, _consts[286]))
							*(*int64)(unsafe.Add(mBase, uint32(v5+int32(48)))) = v42
							v47 = *(*int64)(unsafe.Add(mBase, _consts[287]))
							*(*int64)(unsafe.Add(mBase, uint32(v5+int32(40)))) = v47
							v52 = *(*int64)(unsafe.Add(mBase, _consts[288]))
							*(*int64)(unsafe.Add(mBase, uint32(v5+int32(32)))) = v52
							v57 = *(*int64)(unsafe.Add(mBase, _consts[289]))
							*(*int64)(unsafe.Add(mBase, uint32(v5+int32(24)))) = v57
							v62 = *(*int64)(unsafe.Add(mBase, _consts[290]))
							*(*int64)(unsafe.Add(mBase, uint32(v5+int32(16)))) = v62
							v67 = *(*int64)(unsafe.Add(mBase, _consts[291]))
							*(*int64)(unsafe.Add(mBase, uint32(v5+int32(8)))) = v67
							v70 = *(*int64)(unsafe.Add(mBase, _consts[292]))
							*(*int64)(unsafe.Add(mBase, uint32(v5))) = v70
							v73 = *(*int32)(unsafe.Add(mBase, _consts[166]))
							if v73 != 0 {
								v86 = *(*int32)(unsafe.Add(mBase, _consts[167]))
								if v86 == int32(0) {
									v93 = *(*int32)(unsafe.Add(mBase, _consts[168]))
									if v93 != 0 {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+96))
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
										*(*int32)(unsafe.Add(mBase, uint32(v5))) = v96
										v99 = v5
									} else {
										v99 = int32(0)
									}
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+96))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
									*(*int32)(unsafe.Add(mBase, uint32(v5))) = v90
									v99 = v5
								}
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, _consts[293]))
								if v75 == int32(0) {
									v86 = *(*int32)(unsafe.Add(mBase, _consts[167]))
									if v86 == int32(0) {
										v93 = *(*int32)(unsafe.Add(mBase, _consts[168]))
										if v93 != 0 {
											v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+96))
											v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
											*(*int32)(unsafe.Add(mBase, uint32(v5))) = v96
											v99 = v5
										} else {
											v99 = int32(0)
										}
									} else {
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+96))
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
										*(*int32)(unsafe.Add(mBase, uint32(v5))) = v90
										v99 = v5
									}
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, _consts[294]))
									if v80 == int32(-1) {
										v83 = int32(0)
									} else {
										v83 = v80
									}
									*(*int32)(unsafe.Add(mBase, uint32(v5))) = v83
									v99 = v5
								}
							}
							v100 = int32(0)
							v102 = *(*int32)(unsafe.Add(mBase, _consts[46]))
							v104 = F_rdbSave(m, v100, v102, v99, v100)
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return
							} else {
								m.G0 = v5 + int32(64)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_flushReplicasOutputBuffers(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[78]))
	v12 = v7 + int32(8)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v13
	goto L1
L1:
	;
	v18 = v7 + int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v20 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return
L3:
	;
	if v20 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20+base.B2i32(v23 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v29
	goto L4
L6:
	;
	v34 = v20
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	if v40 != 0 {
		v46 = int32(0)
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	switch v48 + int32(-9) {
	case 0, 2:
		goto L12
	default:
		goto L11
	}
L10:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+202)))
	v46 = base.B2i32(v41&int32(64) == int32(0))
	goto L9
L11:
	;
	v65 = v7 + int32(8)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v67 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+201)))
	if (int32(base.Ui32(v51)>>(uint(int32(2))%32))|v46)&int32(1) != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v57 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v58 = F_clientHasPendingReplies(m, v38)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	if v58 == int32(0) {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v62 = F_writeToClient(m, v38)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	if v67 != 0 {
		v34 = v67
		goto L7
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v67+base.B2i32(v70 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v76
	goto L20
L22:
	;
	goto L8
}
func F_flushdbCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = F_getFlushCommandFlags(m, l0, v5+int32(12))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v9 == int32(-1) {
			m.G0 = v5 + int32(16)
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			v20 = F_emptyData(m, v15, v16|int32(2), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v23 = *(*int64)(unsafe.Add(mBase, _consts[180]))
				*(*int64)(unsafe.Add(mBase, _consts[180])) = v20 + v23
				F_forceCommandPropagation(m, l0, int32(3))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, _consts[27]))
					F_addReply(m, l0, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						m.G0 = v5 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_fmod(m *base.Module, l0 float64, l1 float64) float64 {
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 float64
	_ = v29
	var v33 int64
	_ = v33
	var v38 float64
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v69 int32
	_ = v69
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v114 int32
	_ = v114
	var v125 int64
	_ = v125
	var v130 int32
	_ = v130
	var v135 int64
	_ = v135
	var v137 int32
	_ = v137
	var v140 int64
	_ = v140
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v152 int32
	_ = v152
	var v157 int64
	_ = v157
	var v159 int32
	_ = v159
	var v162 int64
	_ = v162
	var v170 int64
	_ = v170
	var v176 int64
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int64
	_ = v186
	var v192 int32
	_ = v192
	var v193 int64
	_ = v193
	var v209 int64
	_ = v209
	v9 = base.I64_reinterpret_f64(l1)
	v11 = v9 << (uint(int64(1)) % 64)
	if v11 == int64(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v33 = v19 << (uint(int64(1)) % 64)
	if base.Ui64(v11) < base.Ui64(v33) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v29 = base.F64_mul(l0, l1)
	return base.F64_div(v29, v29)
L3:
	;
	goto L4
L4:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l1)&int64(9223372036854775807)) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v19 = base.I64_reinterpret_f64(l0)
	v23 = int32(2047)
	v24 = base.I32_wrap_i64(int64(base.Ui64(v19)>>(uint(int64(52))%64))) & v23
	if v24 != v23 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L2
L7:
	;
	v44 = base.I32_wrap_i64(int64(base.Ui64(v9)>>(uint(int64(52))%64))) & int32(2047)
	if v24 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	if v33 == v11 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v38 = base.F64_mul(l0, float64(0))
	goto L11
L10:
	;
	v38 = l0
	goto L11
L11:
	;
	return v38
L12:
	;
	if v44 != 0 {
		goto L21
	} else {
		goto L22
	}
L13:
	;
	v83 = v19&int64(4503599627370495) | int64(4503599627370496)
	v85 = v24
	goto L12
L14:
	;
	v45 = int32(0)
	v47 = v19 << (uint(int64(12)) % 64)
	if v47 < int64(0) {
		v69 = v45
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v83 = v19 << (uint(base.I64_extend_i32_u(int32(1)-v69)) % 64)
	v85 = v69
	goto L12
L16:
	;
	v53 = v47
	v55 = v45
	goto L17
L17:
	;
	v59 = v55 + int32(-1)
	v61 = v53 << (uint(int64(1)) % 64)
	if int64(-1) < v61 {
		v53 = v61
		v55 = v59
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v69 = v59
	goto L15
L19:
	;
	goto L18
L20:
	;
	if v85 <= v130 {
		v157 = v83
		v159 = v85
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v125 = v9&int64(4503599627370495) | int64(4503599627370496)
	v130 = v44
	goto L20
L22:
	;
	v88 = int32(0)
	v90 = v9 << (uint(int64(12)) % 64)
	if v90 < int64(0) {
		v114 = v88
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v125 = v9 << (uint(base.I64_extend_i32_u(int32(1)-v114)) % 64)
	v130 = v114
	goto L20
L24:
	;
	v99 = v90
	v100 = v88
	goto L25
L25:
	;
	v102 = v100 + int32(-1)
	v104 = v99 << (uint(int64(1)) % 64)
	if int64(-1) < v104 {
		v99 = v104
		v100 = v102
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v114 = v102
	goto L23
L27:
	;
	goto L26
L28:
	;
	v162 = v157 - v125
	if v162 < int64(0) {
		v170 = v157
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v135 = v83
	v137 = v85
	goto L30
L30:
	;
	v140 = v135 - v125
	if v140 < int64(0) {
		v148 = v135
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v157 = v150
	v159 = v130
	goto L28
L32:
	;
	v150 = v148 << (uint(int64(1)) % 64)
	v152 = v137 + int32(-1)
	if v130 < v152 {
		v135 = v150
		v137 = v152
		goto L30
	} else {
		goto L35
	}
L33:
	;
	if v140 != int64(0) {
		v148 = v140
		goto L32
	} else {
		goto L34
	}
L34:
	;
	return base.F64_mul(l0, float64(0))
L35:
	;
	goto L31
L36:
	;
	if base.Ui64(v170) <= base.Ui64(int64(4503599627370495)) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	if v162 != int64(0) {
		v170 = v162
		goto L36
	} else {
		goto L38
	}
L38:
	;
	return base.F64_mul(l0, float64(0))
L39:
	;
	if v192 < int32(1) {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	v176 = v170
	v178 = v159
	goto L42
L41:
	;
	v192 = v159
	v193 = v170
	goto L39
L42:
	;
	v182 = v178 + int32(-1)
	v186 = v176 << (uint(int64(1)) % 64)
	if base.Ui64(v176) < base.Ui64(int64(2251799813685248)) {
		v176 = v186
		v178 = v182
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v192 = v182
	v193 = v186
	goto L39
L44:
	;
	goto L43
L45:
	;
	return base.F64_reinterpret_i64(v209 | v19&int64(-9223372036854775807-1))
L46:
	;
	v209 = int64(base.Ui64(v193) >> (uint(base.I64_extend_i32_u(int32(1)-v192)) % 64))
	goto L45
L47:
	;
	v209 = v193 + int64(-4503599627370496) | base.I64_extend_i32_u(v192)<<(uint(int64(52))%64)
	goto L45
}
func F_fmt_fp(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int64
	_ = v36
	var v41 float64
	_ = v41
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 float64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int64
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 float64
	_ = v115
	var v116 int32
	_ = v116
	var v119 float64
	_ = v119
	var v120 int32
	_ = v120
	var v130 float64
	_ = v130
	var v133 float64
	_ = v133
	var v134 float64
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v163 float64
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 float64
	_ = v177
	var v188 int32
	_ = v188
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 float64
	_ = v220
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v270 int64
	_ = v270
	var v273 int32
	_ = v273
	var v290 int64
	_ = v290
	var v294 int64
	_ = v294
	var v295 int64
	_ = v295
	var v296 int64
	_ = v296
	var v299 int64
	_ = v299
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v329 int32
	_ = v329
	var v349 int32
	_ = v349
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v563 int32
	_ = v563
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v643 int32
	_ = v643
	var v655 int32
	_ = v655
	var v662 int32
	_ = v662
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v704 int32
	_ = v704
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v738 int32
	_ = v738
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v763 float64
	_ = v763
	var v769 int32
	_ = v769
	var v776 float64
	_ = v776
	var v781 float64
	_ = v781
	var v784 int32
	_ = v784
	var v786 float64
	_ = v786
	var v788 float64
	_ = v788
	var v789 int32
	_ = v789
	var v794 float64
	_ = v794
	var v795 float64
	_ = v795
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v817 int32
	_ = v817
	var v824 int32
	_ = v824
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v861 int32
	_ = v861
	var v868 int32
	_ = v868
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v897 int32
	_ = v897
	var v904 int32
	_ = v904
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v931 int32
	_ = v931
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v961 int32
	_ = v961
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v989 int32
	_ = v989
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1055 int32
	_ = v1055
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1092 int32
	_ = v1092
	var v1113 int32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1144 int32
	_ = v1144
	var v1155 int32
	_ = v1155
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1193 int64
	_ = v1193
	var v1200 int64
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1207 int32
	_ = v1207
	var v1208 int64
	_ = v1208
	var v1209 int64
	_ = v1209
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1221 int64
	_ = v1221
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1242 int32
	_ = v1242
	var v1247 int32
	_ = v1247
	var v1266 int32
	_ = v1266
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1301 int32
	_ = v1301
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1345 int32
	_ = v1345
	var v1357 int32
	_ = v1357
	var v1363 int32
	_ = v1363
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1377 int32
	_ = v1377
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1406 int32
	_ = v1406
	var v1414 int64
	_ = v1414
	var v1421 int64
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1428 int32
	_ = v1428
	var v1429 int64
	_ = v1429
	var v1430 int64
	_ = v1430
	var v1436 int32
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1442 int64
	_ = v1442
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1463 int32
	_ = v1463
	var v1468 int32
	_ = v1468
	var v1488 int32
	_ = v1488
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1528 int32
	_ = v1528
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1556 int32
	_ = v1556
	var v1576 int32
	_ = v1576
	var v1580 int32
	_ = v1580
	var v1588 int64
	_ = v1588
	var v1595 int64
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1602 int32
	_ = v1602
	var v1603 int64
	_ = v1603
	var v1604 int64
	_ = v1604
	var v1610 int32
	_ = v1610
	var v1615 int32
	_ = v1615
	var v1616 int64
	_ = v1616
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1637 int32
	_ = v1637
	var v1642 int32
	_ = v1642
	var v1661 int32
	_ = v1661
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1696 int32
	_ = v1696
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1731 int32
	_ = v1731
	var v1735 int32
	_ = v1735
	var v1748 int32
	_ = v1748
	var v1752 int32
	_ = v1752
	var v1764 int64
	_ = v1764
	var v1771 int64
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1778 int32
	_ = v1778
	var v1779 int64
	_ = v1779
	var v1780 int64
	_ = v1780
	var v1786 int32
	_ = v1786
	var v1791 int32
	_ = v1791
	var v1792 int64
	_ = v1792
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1813 int32
	_ = v1813
	var v1818 int32
	_ = v1818
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1828 int32
	_ = v1828
	var v1844 int32
	_ = v1844
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1879 int32
	_ = v1879
	var v1891 int32
	_ = v1891
	var v1908 int32
	_ = v1908
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1915 int32
	_ = v1915
	var v1935 int32
	_ = v1935
	var v1948 int32
	_ = v1948
	var v1953 int32
	_ = v1953
	var v1956 int32
	_ = v1956
	var v1968 int32
	_ = v1968
	var v1986 int32
	_ = v1986
	var v1991 int32
	_ = v1991
	var v2024 int32
	_ = v2024
	var v2026 int32
	_ = v2026
	var v2035 int32
	_ = v2035
	var v2052 int32
	_ = v2052
	var v2068 float64
	_ = v2068
	var v2070 float64
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2083 float64
	_ = v2083
	var v2110 int32
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2115 int64
	_ = v2115
	var v2122 int64
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2129 int32
	_ = v2129
	var v2130 int64
	_ = v2130
	var v2131 int64
	_ = v2131
	var v2137 int32
	_ = v2137
	var v2142 int32
	_ = v2142
	var v2143 int64
	_ = v2143
	var v2151 int32
	_ = v2151
	var v2153 int32
	_ = v2153
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2164 int32
	_ = v2164
	var v2169 int32
	_ = v2169
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2187 int32
	_ = v2187
	var v2189 int32
	_ = v2189
	var v2197 int32
	_ = v2197
	var v2209 float64
	_ = v2209
	var v2220 int32
	_ = v2220
	var v2241 int32
	_ = v2241
	var v2243 int32
	_ = v2243
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2252 float64
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2263 int32
	_ = v2263
	var v2267 int32
	_ = v2267
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2281 int32
	_ = v2281
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2289 int32
	_ = v2289
	var v2291 int32
	_ = v2291
	var v2296 int32
	_ = v2296
	var v2300 int32
	_ = v2300
	var v2303 int32
	_ = v2303
	var v2306 int32
	_ = v2306
	var v2308 int32
	_ = v2308
	var v2313 int32
	_ = v2313
	var v2315 int32
	_ = v2315
	var v2329 int32
	_ = v2329
	v7 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(560)
	m.G0 = v31
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v7
	v36 = base.I64_reinterpret_f64(l1)
	goto L3
L1:
	;
	v61 = int64(9218868437227405312)
	if v58&v61 != v61 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	if l4&int32(2048) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if int64(-1) < v36 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v41 = base.F64_neg(l1)
	goto L5
L5:
	;
	v56 = v41
	v57 = v7
	v58 = base.I64_reinterpret_f64(v41)
	v59 = int32(1)
	v60 = int32(_a2783)
	goto L1
L6:
	;
	v52 = l4 & int32(1)
	if v52 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v56 = l1
	v57 = v7
	v58 = v36
	v59 = int32(1)
	v60 = int32(_a2784)
	goto L1
L8:
	;
	v53 = int32(_a2785)
	goto L10
L9:
	;
	v53 = int32(_a2786)
	goto L10
L10:
	;
	v56 = l1
	v57 = base.B2i32(v52 == int32(0))
	v58 = v36
	v59 = v52
	v60 = v53
	goto L1
L11:
	;
	m.G0 = v31 + int32(560)
	return v2329
L12:
	;
	v97 = v31 + int32(16)
	v99 = v31 + int32(44)
	v102 = base.I64_reinterpret_f64(v56)
	v106 = int32(2047)
	v107 = base.I32_wrap_i64(int64(base.Ui64(v102)>>(uint(int64(52))%64))) & v106
	if v107 == v106 {
		v130 = v56
		goto L36
	} else {
		goto L37
	}
L13:
	;
	v67 = v59 + int32(3)
	F_pad(m, l0, int32(32), l2, v67, l4&int32(-65537))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	F_out(m, l0, v60, v59)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v79 = l5 & int32(32)
	if v79 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v80 = int32(_a2581)
	goto L19
L18:
	;
	v80 = int32(_a2787)
	goto L19
L19:
	;
	if v79 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v83 = int32(_a2583)
	goto L22
L21:
	;
	v83 = int32(_a2788)
	goto L22
L22:
	;
	if base.F64_ne(v56, v56) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v85 = v80
	goto L25
L24:
	;
	v85 = v83
	goto L25
L25:
	;
	F_out(m, l0, v85, int32(3))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L14
	} else {
		goto L26
	}
L26:
	;
	F_pad(m, l0, int32(32), l2, v67, l4^int32(8192))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	if v67 < l2 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v95 = l2
	goto L30
L29:
	;
	v95 = v67
	goto L30
L30:
	;
	v2329 = v95
	goto L11
L31:
	;
	v2035 = v60 + l5<<(uint(int32(26))%32)>>(uint(int32(31))%32)&int32(9)
	if base.Ui32(int32(11)) < base.Ui32(l3) {
		v2083 = v134
		goto L339
	} else {
		goto L340
	}
L32:
	;
	v170 = int32(0)
	if v167 < v170 {
		goto L52
	} else {
		goto L53
	}
L33:
	;
	v155 = v137 + int32(-29)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v155
	if l3 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L34:
	;
	v146 = l5 | int32(32)
	if v146 == int32(97) {
		goto L31
	} else {
		goto L45
	}
L35:
	;
	v134 = base.F64_add(v133, v133)
	if base.F64_eq(v134, float64(0)) != 0 {
		goto L34
	} else {
		goto L43
	}
L36:
	;
	v133 = v130
	goto L35
L37:
	;
	if v107 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v107 + int32(-1022)
	v130 = base.F64_reinterpret_i64(v102&int64(-9218868437227405313) | int64(4602678819172646912))
	goto L36
L39:
	;
	if base.F64_ne(v56, float64(0)) != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v120
	v133 = v119
	goto L35
L41:
	;
	v115 = F_frexp(m, base.F64_mul(v56, float64(1.8446744073709552e+19)), v99)
	mBase = m.M
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v119 = v115
	v120 = v116 + int32(-64)
	goto L40
L42:
	;
	v119 = v56
	v120 = int32(0)
	goto L40
L43:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v137 + int32(-1)
	v142 = l5 | int32(32)
	if v142 != int32(97) {
		goto L33
	} else {
		goto L44
	}
L44:
	;
	goto L31
L45:
	;
	if l3 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v152 = int32(6)
	goto L48
L47:
	;
	v152 = l3
	goto L48
L48:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v163 = v134
	v165 = v146
	v166 = v152
	v167 = v153
	goto L32
L49:
	;
	v160 = int32(6)
	goto L51
L50:
	;
	v160 = l3
	goto L51
L51:
	;
	v163 = base.F64_mul(v134, float64(2.68435456e+08))
	v165 = v142
	v166 = v160
	v167 = v155
	goto L32
L52:
	;
	v174 = v170
	goto L54
L53:
	;
	v174 = int32(288)
	goto L54
L54:
	;
	v175 = v31 + int32(48) + v174
	v177 = v163
	v188 = v175
	goto L55
L55:
	;
	if base.F64_lt(v177, float64(4.294967296e+09))&base.F64_ge(v177, float64(0)) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	if int32(1) <= v167 {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v213
	v216 = v188 + int32(4)
	v220 = base.F64_mul(base.F64_sub(v177, base.F64_convert_i32_u(v213)), float64(1e+09))
	if base.F64_ne(v220, float64(0)) != 0 {
		v177 = v220
		v188 = v216
		goto L55
	} else {
		goto L60
	}
L58:
	;
	v213 = int32(0)
	goto L57
L59:
	;
	v211 = base.I32_trunc_f64_u(v177)
	v213 = v211
	goto L57
L60:
	;
	goto L56
L61:
	;
	if int32(-1) < v396 {
		v563 = v388
		v572 = v397
		v575 = v7
		goto L81
	} else {
		goto L82
	}
L62:
	;
	v237 = v216
	v244 = v167
	v245 = v175
	goto L64
L63:
	;
	v388 = v216
	v396 = v167
	v397 = v175
	goto L61
L64:
	;
	v253 = int32(29)
	if base.Ui32(v244) < base.Ui32(v253) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v388 = v349
	v396 = v373
	v397 = v329
	goto L61
L66:
	;
	v256 = v244
	goto L68
L67:
	;
	v256 = v253
	goto L68
L68:
	;
	v258 = v237 + int32(-4)
	if base.Ui32(v258) < base.Ui32(v245) {
		v329 = v245
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v349 = v237
	goto L76
L70:
	;
	v270 = int64(0)
	v273 = v258
	goto L71
L71:
	;
	v290 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v273))))
	v294 = v290<<(uint(base.I64_extend_i32_u(v256))%64) + v270&int64(4294967295)
	v295 = int64(1000000000)
	v296 = base.I64_div_u_s(v294, v295)
	v299 = v294 - v296*v295
	*(*uint32)(unsafe.Add(mBase, uint32(v273))) = uint32(v299)
	v302 = v273 + int32(-4)
	if base.Ui32(v245) <= base.Ui32(v302) {
		v270 = v296
		v273 = v302
		goto L71
	} else {
		goto L73
	}
L72:
	;
	if base.Ui64(v294) < base.Ui64(int64(1000000000)) {
		v329 = v245
		goto L69
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v307 = v245 + int32(-4)
	*(*uint32)(unsafe.Add(mBase, uint32(v307))) = uint32(v296)
	v329 = v307
	goto L69
L75:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v373 = v372 - v256
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v373
	if int32(0) < v373 {
		v237 = v349
		v244 = v373
		v245 = v329
		goto L64
	} else {
		goto L80
	}
L76:
	;
	if base.Ui32(v349) <= base.Ui32(v329) {
		goto L75
	} else {
		goto L78
	}
L77:
	;
	goto L75
L78:
	;
	v367 = v349 + int32(-4)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	if v368 == int32(0) {
		v349 = v367
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	goto L65
L81:
	;
	if base.Ui32(v563) <= base.Ui32(v572) {
		v643 = int32(0)
		goto L102
	} else {
		goto L103
	}
L82:
	;
	v410 = base.I32_div_u_s(v166+int32(25), int32(9))
	v412 = v410 + int32(1)
	v426 = v388
	v434 = v396
	v435 = v397
	goto L83
L83:
	;
	v444 = int32(0) - v434
	v445 = int32(9)
	if base.Ui32(v444) < base.Ui32(v445) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v563 = v549
	v572 = v540
	v575 = v412
	goto L81
L85:
	;
	v448 = v444
	goto L87
L86:
	;
	v448 = v445
	goto L87
L87:
	;
	if base.Ui32(v435) < base.Ui32(v426) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v538 = v537 + v448
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v538
	v540 = v435 + v521
	if v165 == int32(102) {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	v457 = int32(-1)
	v474 = v435
	v481 = int32(0)
	goto L91
L90:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
	v520 = v426
	v521 = base.B2i32(v450 == int32(0)) << (uint(int32(2)) % 32)
	goto L88
L91:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v474)))
	*(*int32)(unsafe.Add(mBase, uint32(v474))) = int32(base.Ui32(v490)>>(uint(v448)%32)) + v481
	v495 = v490 & (v457<<(uint(v448)%32) ^ v457) * int32(base.Ui32(int32(1000000000))>>(uint(v448)%32))
	v497 = v474 + int32(4)
	if base.Ui32(v497) < base.Ui32(v426) {
		v474 = v497
		v481 = v495
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
	v500 = int32(0)
	v503 = base.B2i32(v499 == v500) << (uint(int32(2)) % 32)
	if v495 == v500 {
		v520 = v426
		v521 = v503
		goto L88
	} else {
		goto L94
	}
L93:
	;
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v426))) = v495
	v520 = v426 + int32(4)
	v521 = v503
	goto L88
L95:
	;
	v541 = v175
	goto L97
L96:
	;
	v541 = v540
	goto L97
L97:
	;
	v542 = int32(2)
	if v412 < (v520-v541)>>(uint(v542)%32) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v549 = v541 + v412<<(uint(v542)%32)
	goto L100
L99:
	;
	v549 = v520
	goto L100
L100:
	;
	if v538 < int32(0) {
		v426 = v549
		v434 = v538
		v435 = v540
		goto L83
	} else {
		goto L101
	}
L101:
	;
	goto L84
L102:
	;
	if v165 == int32(102) {
		goto L109
	} else {
		goto L110
	}
L103:
	;
	v586 = (v175 - v572) >> (uint(int32(2)) % 32) * int32(9)
	v587 = int32(10)
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	if base.Ui32(v588) < base.Ui32(v587) {
		v643 = v586
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v603 = v587
	v610 = v586
	goto L105
L105:
	;
	v620 = v610 + int32(1)
	v622 = v603 * int32(10)
	if base.Ui32(v622) <= base.Ui32(v588) {
		v603 = v622
		v610 = v620
		goto L105
	} else {
		goto L107
	}
L106:
	;
	v643 = v620
	goto L102
L107:
	;
	goto L106
L108:
	;
	v989 = v961
	goto L159
L109:
	;
	v655 = int32(0)
	goto L111
L110:
	;
	v655 = v643
	goto L111
L111:
	;
	v662 = v166 - v655 - base.B2i32(v166 != int32(0))&base.B2i32(v165 == int32(103))
	if (v563-v175)>>(uint(int32(2))%32)*int32(9)+int32(-9) <= v662 {
		v961 = v563
		v969 = v643
		v970 = v572
		v973 = v575
		goto L108
	} else {
		goto L112
	}
L112:
	;
	if v167 < int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v677 = int32(-4092)
	goto L115
L114:
	;
	v677 = int32(-3804)
	goto L115
L115:
	;
	v680 = v662 + int32(9216)
	v681 = int32(9)
	v682 = base.I32_div_s(v680, v681)
	v685 = v31 + int32(48) + v677 + v682<<(uint(int32(2))%32)
	v686 = int32(10)
	v689 = v680 - v682*v681
	if int32(7) < v689 {
		v738 = v686
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v755 = v685 + int32(4)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v685)))
	v757 = base.I32_div_u_s(v756, v738)
	v759 = v756 - v757*v738
	if v759 != 0 {
		goto L122
	} else {
		goto L123
	}
L117:
	;
	v695 = v689
	v704 = v686
	goto L118
L118:
	;
	v721 = v704 * int32(10)
	v723 = v695 + int32(1)
	if v723 != int32(8) {
		v695 = v723
		v704 = v721
		goto L118
	} else {
		goto L120
	}
L119:
	;
	v738 = v721
	goto L116
L120:
	;
	goto L119
L121:
	;
	v947 = v931 + int32(4)
	if base.Ui32(v947) < base.Ui32(v563) {
		goto L155
	} else {
		goto L156
	}
L122:
	;
	if v757&int32(1) != 0 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	if v755 == v563 {
		v931 = v685
		v937 = v643
		v938 = v572
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	if v755 == v563 {
		goto L131
	} else {
		goto L132
	}
L126:
	;
	v776 = float64(9.007199254740994e+15)
	goto L125
L127:
	;
	v763 = float64(9.007199254740992e+15)
	if v738 != int32(1000000000) {
		v776 = v763
		goto L125
	} else {
		goto L128
	}
L128:
	;
	if base.Ui32(v685) <= base.Ui32(v572) {
		v776 = v763
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685+int32(-4)))))
	if v769&int32(1) == int32(0) {
		v776 = v763
		goto L125
	} else {
		goto L130
	}
L130:
	;
	goto L126
L131:
	;
	v781 = float64(1)
	goto L133
L132:
	;
	v781 = float64(1.5)
	goto L133
L133:
	;
	v784 = int32(base.Ui32(v738) >> (uint(int32(1)) % 32))
	if v759 == v784 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v786 = v781
	goto L136
L135:
	;
	v786 = float64(1.5)
	goto L136
L136:
	;
	if base.Ui32(v759) < base.Ui32(v784) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v788 = float64(0.5)
	goto L139
L138:
	;
	v788 = v786
	goto L139
L139:
	;
	if v57 != 0 {
		v794 = v776
		v795 = v788
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v796 = v756 - v759
	*(*int32)(unsafe.Add(mBase, uint32(v685))) = v796
	if base.F64_eq(base.F64_add(v794, v795), v794) != 0 {
		v931 = v685
		v937 = v643
		v938 = v572
		goto L121
	} else {
		goto L143
	}
L141:
	;
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v789 != int32(45) {
		v794 = v776
		v795 = v788
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v794 = base.F64_neg(v776)
	v795 = base.F64_neg(v788)
	goto L140
L143:
	;
	v800 = v796 + v738
	*(*int32)(unsafe.Add(mBase, uint32(v685))) = v800
	if base.Ui32(v800) < base.Ui32(int32(1000000000)) {
		v861 = v685
		v868 = v572
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v880 = (v175 - v868) >> (uint(int32(2)) % 32) * int32(9)
	v881 = int32(10)
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	if base.Ui32(v882) < base.Ui32(v881) {
		v931 = v861
		v937 = v880
		v938 = v868
		goto L121
	} else {
		goto L151
	}
L145:
	;
	v817 = v685
	v824 = v572
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v817))) = int32(0)
	v835 = v817 + int32(-4)
	if base.Ui32(v824) <= base.Ui32(v835) {
		v841 = v824
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v861 = v835
	v868 = v841
	goto L144
L148:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v835)))
	v844 = v842 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v835))) = v844
	if base.Ui32(int32(999999999)) < base.Ui32(v844) {
		v817 = v835
		v824 = v841
		goto L146
	} else {
		goto L150
	}
L149:
	;
	v838 = v824 + int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(v838))) = int32(0)
	v841 = v838
	goto L148
L150:
	;
	goto L147
L151:
	;
	v897 = v881
	v904 = v880
	goto L152
L152:
	;
	v914 = v904 + int32(1)
	v916 = v897 * int32(10)
	if base.Ui32(v916) <= base.Ui32(v882) {
		v897 = v916
		v904 = v914
		goto L152
	} else {
		goto L154
	}
L153:
	;
	v931 = v861
	v937 = v914
	v938 = v868
	goto L121
L154:
	;
	goto L153
L155:
	;
	v949 = v947
	goto L157
L156:
	;
	v949 = v563
	goto L157
L157:
	;
	v961 = v949
	v969 = v937
	v970 = v938
	v973 = v757
	goto L108
L158:
	;
	if v165 == int32(103) {
		goto L164
	} else {
		goto L165
	}
L159:
	;
	v1006 = base.B2i32(base.Ui32(v989) <= base.Ui32(v970))
	if base.Ui32(v989) <= base.Ui32(v970) {
		goto L158
	} else {
		goto L161
	}
L160:
	;
	goto L158
L161:
	;
	v1008 = v989 + int32(-4)
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v1008)))
	if v1009 == int32(0) {
		v989 = v1008
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	v1167 = int32(-1)
	v1170 = v1155 | v1164
	if v1170 != 0 {
		goto L197
	} else {
		goto L198
	}
L164:
	;
	v1017 = int32(-1)
	if v166 != 0 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v1144 = l5
	v1155 = v166
	v1164 = l4 & int32(8)
	goto L163
L166:
	;
	v1021 = v166
	goto L168
L167:
	;
	v1021 = int32(1)
	goto L168
L168:
	;
	v1025 = base.B2i32(v969 < v1021) & base.B2i32(int32(-5) < v969)
	if v1025 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1026 = v969 ^ v1017
	goto L171
L170:
	;
	v1026 = v1017
	goto L171
L171:
	;
	v1027 = v1026 + v1021
	if v1025 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v1030 = int32(-1)
	goto L174
L173:
	;
	v1030 = int32(-2)
	goto L174
L174:
	;
	v1031 = v1030 + l5
	v1033 = l4 & int32(8)
	if v1033 != 0 {
		v1144 = v1031
		v1155 = v1027
		v1164 = v1033
		goto L163
	} else {
		goto L175
	}
L175:
	;
	v1034 = int32(-9)
	if base.Ui32(v989) <= base.Ui32(v970) {
		v1092 = v1034
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v1113 = (v989 - v175) >> (uint(int32(2)) % 32) * int32(9)
	if v1031&int32(-33) != int32(70) {
		goto L183
	} else {
		goto L184
	}
L177:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v989+int32(-4))))
	if v1037 == int32(0) {
		v1092 = v1034
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v1040 = int32(10)
	v1041 = int32(0)
	v1043 = base.I32_rem_u_s(v1037, v1040)
	if v1043 != 0 {
		v1092 = v1041
		goto L176
	} else {
		goto L179
	}
L179:
	;
	v1047 = v1040
	v1055 = v1041
	goto L180
L180:
	;
	v1075 = v1047 * int32(10)
	v1076 = base.I32_rem_u_s(v1037, v1075)
	if v1076 == int32(0) {
		v1047 = v1075
		v1055 = v1055 + int32(1)
		goto L180
	} else {
		goto L182
	}
L181:
	;
	v1092 = v1055 ^ int32(-1)
	goto L176
L182:
	;
	goto L181
L183:
	;
	v1128 = int32(0)
	v1132 = v969 + v1113 + v1092 + int32(-9)
	if v1128 < v1132 {
		goto L191
	} else {
		goto L192
	}
L184:
	;
	v1118 = int32(0)
	v1121 = v1113 + v1092 + int32(-9)
	if v1118 < v1121 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v1125 = v1121
	goto L187
L186:
	;
	v1125 = v1118
	goto L187
L187:
	;
	if v1027 < v1125 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1127 = v1027
	goto L190
L189:
	;
	v1127 = v1125
	goto L190
L190:
	;
	v1144 = v1031
	v1155 = v1127
	v1164 = v1118
	goto L163
L191:
	;
	v1136 = v1132
	goto L193
L192:
	;
	v1136 = v1128
	goto L193
L193:
	;
	if v1027 < v1136 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1138 = v1027
	goto L196
L195:
	;
	v1138 = v1136
	goto L196
L196:
	;
	v1144 = v1031
	v1155 = v1138
	v1164 = v1128
	goto L163
L197:
	;
	v1171 = int32(2147483645)
	goto L199
L198:
	;
	v1171 = int32(2147483646)
	goto L199
L199:
	;
	if v1171 < v1155 {
		v2329 = v1167
		goto L11
	} else {
		goto L200
	}
L200:
	;
	v1177 = v1155 + base.B2i32(v1170 != int32(0)) + int32(1)
	v1179 = v1144 & int32(-33)
	if v1179 != int32(70) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v1363 = v1345 + v1177
	if v59^int32(2147483647) < v1363 {
		v2329 = int32(-1)
		goto L11
	} else {
		goto L229
	}
L202:
	;
	v1190 = v969 >> (uint(int32(31)) % 32)
	v1193 = base.I64_extend_i32_u(v969 ^ v1190 - v1190)
	if base.Ui64(int64(4294967296)) <= base.Ui64(v1193) {
		goto L211
	} else {
		goto L212
	}
L203:
	;
	if v1177^int32(2147483647) < v969 {
		v2329 = v1167
		goto L11
	} else {
		goto L204
	}
L204:
	;
	v1185 = int32(0)
	if v1185 < v969 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v1188 = v969
	goto L207
L206:
	;
	v1188 = v1185
	goto L207
L207:
	;
	v1345 = v1188
	v1357 = v973
	goto L201
L208:
	;
	v1319 = v1301 + int32(-2)
	*(*uint8)(unsafe.Add(mBase, uint32(v1319))) = uint8(v1144)
	v1321 = int32(-1)
	if v969 < int32(0) {
		goto L225
	} else {
		goto L226
	}
L209:
	;
	if int32(1) < v97-v1247 {
		v1301 = v1247
		goto L208
	} else {
		goto L221
	}
L210:
	;
	if v1221 == int64(0) {
		v1247 = v1220
		goto L216
	} else {
		goto L217
	}
L211:
	;
	v1200 = v1193
	v1201 = v97
	goto L213
L212:
	;
	v1220 = v97
	v1221 = v1193
	goto L210
L213:
	;
	v1207 = v1201 + int32(-1)
	v1208 = int64(10)
	v1209 = base.I64_div_u_s(v1200, v1208)
	v1215 = base.I32_wrap_i64(v1200-v1209*v1208) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1207))) = uint8(v1215)
	if base.Ui64(int64(42949672959)) < base.Ui64(v1200) {
		v1200 = v1209
		v1201 = v1207
		goto L213
	} else {
		goto L215
	}
L214:
	;
	v1220 = v1207
	v1221 = v1209
	goto L210
L215:
	;
	goto L214
L216:
	;
	goto L209
L217:
	;
	v1229 = v1220
	v1231 = base.I32_wrap_i64(v1221)
	goto L218
L218:
	;
	v1235 = v1229 + int32(-1)
	v1236 = int32(10)
	v1237 = base.I32_div_u_s(v1231, v1236)
	v1242 = v1231 - v1237*v1236 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1235))) = uint8(v1242)
	if base.Ui32(int32(9)) < base.Ui32(v1231) {
		v1229 = v1235
		v1231 = v1237
		goto L218
	} else {
		goto L220
	}
L219:
	;
	v1247 = v1235
	goto L216
L220:
	;
	goto L219
L221:
	;
	v1266 = v1247
	goto L222
L222:
	;
	v1284 = v1266 + int32(-1)
	v1285 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1284))) = uint8(v1285)
	if v97-v1284 < int32(2) {
		v1266 = v1284
		goto L222
	} else {
		goto L224
	}
L223:
	;
	v1301 = v1284
	goto L208
L224:
	;
	goto L223
L225:
	;
	v1328 = int32(45)
	goto L227
L226:
	;
	v1328 = int32(43)
	goto L227
L227:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1301+v1321))) = uint8(v1328)
	v1330 = v97 - v1319
	if v1177^int32(2147483647) < v1330 {
		v2329 = v1321
		goto L11
	} else {
		goto L228
	}
L228:
	;
	v1345 = v1330
	v1357 = v1319
	goto L201
L229:
	;
	v1368 = v1363 + v59
	F_pad(m, l0, int32(32), l2, v1368, l4)
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L14
	} else {
		goto L230
	}
L230:
	;
	F_out(m, l0, v60, v59)
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L14
	} else {
		goto L231
	}
L231:
	;
	F_pad(m, l0, int32(48), l2, v1368, l4^int32(65536))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L14
	} else {
		goto L232
	}
L232:
	;
	if v1179 != int32(70) {
		goto L236
	} else {
		goto L237
	}
L233:
	;
	F_pad(m, l0, int32(32), l2, v1368, l4^int32(8192))
	mBase = m.M
	v2024 = m.ExcPending
	if v2024 != 0 {
		goto L14
	} else {
		goto L335
	}
L234:
	;
	v1986 = int32(9)
	F_pad(m, l0, int32(48), v1968+v1986, v1986, int32(0))
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L14
	} else {
		goto L334
	}
L235:
	;
	v1968 = v1155
	goto L234
L236:
	;
	if v1155 < int32(0) {
		v1935 = v1155
		goto L295
	} else {
		goto L296
	}
L237:
	;
	v1383 = v31 + int32(16) | int32(9)
	if base.Ui32(v175) < base.Ui32(v970) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v1385 = v175
	goto L240
L239:
	;
	v1385 = v970
	goto L240
L240:
	;
	v1406 = v1385
	goto L241
L241:
	;
	v1414 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1406))))
	if base.Ui64(int64(4294967296)) <= base.Ui64(v1414) {
		goto L245
	} else {
		goto L246
	}
L242:
	;
	if v1170 == int32(0) {
		goto L265
	} else {
		goto L266
	}
L243:
	;
	if v1406 == v1385 {
		goto L256
	} else {
		goto L257
	}
L244:
	;
	if v1442 == int64(0) {
		v1468 = v1441
		goto L250
	} else {
		goto L251
	}
L245:
	;
	v1421 = v1414
	v1422 = v1383
	goto L247
L246:
	;
	v1441 = v1383
	v1442 = v1414
	goto L244
L247:
	;
	v1428 = v1422 + int32(-1)
	v1429 = int64(10)
	v1430 = base.I64_div_u_s(v1421, v1429)
	v1436 = base.I32_wrap_i64(v1421-v1430*v1429) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1428))) = uint8(v1436)
	if base.Ui64(int64(42949672959)) < base.Ui64(v1421) {
		v1421 = v1430
		v1422 = v1428
		goto L247
	} else {
		goto L249
	}
L248:
	;
	v1441 = v1428
	v1442 = v1430
	goto L244
L249:
	;
	goto L248
L250:
	;
	goto L243
L251:
	;
	v1450 = v1441
	v1452 = base.I32_wrap_i64(v1442)
	goto L252
L252:
	;
	v1456 = v1450 + int32(-1)
	v1457 = int32(10)
	v1458 = base.I32_div_u_s(v1452, v1457)
	v1463 = v1452 - v1458*v1457 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1456))) = uint8(v1463)
	if base.Ui32(int32(9)) < base.Ui32(v1452) {
		v1450 = v1456
		v1452 = v1458
		goto L252
	} else {
		goto L254
	}
L253:
	;
	v1468 = v1456
	goto L250
L254:
	;
	goto L253
L255:
	;
	F_out(m, l0, v1528, v1383-v1528)
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L14
	} else {
		goto L263
	}
L256:
	;
	if v1468 != v1383 {
		v1528 = v1468
		goto L255
	} else {
		goto L262
	}
L257:
	;
	if base.Ui32(v1468) <= base.Ui32(v31+int32(16)) {
		v1528 = v1468
		goto L255
	} else {
		goto L258
	}
L258:
	;
	v1488 = v1468
	goto L259
L259:
	;
	v1506 = v1488 + int32(-1)
	v1507 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1506))) = uint8(v1507)
	if base.Ui32(v31+int32(16)) < base.Ui32(v1506) {
		v1488 = v1506
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v1528 = v1506
	goto L255
L262:
	;
	v1514 = v1468 + int32(-1)
	v1515 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1514))) = uint8(v1515)
	v1528 = v1514
	goto L255
L263:
	;
	v1549 = v1406 + int32(4)
	if base.Ui32(v1549) <= base.Ui32(v175) {
		v1406 = v1549
		goto L241
	} else {
		goto L264
	}
L264:
	;
	goto L242
L265:
	;
	if base.Ui32(v989) <= base.Ui32(v1549) {
		goto L235
	} else {
		goto L268
	}
L266:
	;
	F_out(m, l0, int32(_a1964), int32(1))
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L14
	} else {
		goto L267
	}
L267:
	;
	goto L265
L268:
	;
	if v1155 < int32(1) {
		goto L235
	} else {
		goto L269
	}
L269:
	;
	v1576 = v1155
	v1580 = v1549
	goto L270
L270:
	;
	v1588 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1580))))
	if base.Ui64(int64(4294967296)) <= base.Ui64(v1588) {
		goto L275
	} else {
		goto L276
	}
L272:
	;
	v1713 = int32(9)
	if v1576 < v1713 {
		goto L289
	} else {
		goto L290
	}
L273:
	;
	if base.Ui32(v1642) <= base.Ui32(v31+int32(16)) {
		v1696 = v1642
		goto L272
	} else {
		goto L285
	}
L274:
	;
	if v1616 == int64(0) {
		v1642 = v1615
		goto L280
	} else {
		goto L281
	}
L275:
	;
	v1595 = v1588
	v1596 = v1383
	goto L277
L276:
	;
	v1615 = v1383
	v1616 = v1588
	goto L274
L277:
	;
	v1602 = v1596 + int32(-1)
	v1603 = int64(10)
	v1604 = base.I64_div_u_s(v1595, v1603)
	v1610 = base.I32_wrap_i64(v1595-v1604*v1603) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1602))) = uint8(v1610)
	if base.Ui64(int64(42949672959)) < base.Ui64(v1595) {
		v1595 = v1604
		v1596 = v1602
		goto L277
	} else {
		goto L279
	}
L278:
	;
	v1615 = v1602
	v1616 = v1604
	goto L274
L279:
	;
	goto L278
L280:
	;
	goto L273
L281:
	;
	v1624 = v1615
	v1626 = base.I32_wrap_i64(v1616)
	goto L282
L282:
	;
	v1630 = v1624 + int32(-1)
	v1631 = int32(10)
	v1632 = base.I32_div_u_s(v1626, v1631)
	v1637 = v1626 - v1632*v1631 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1630))) = uint8(v1637)
	if base.Ui32(int32(9)) < base.Ui32(v1626) {
		v1624 = v1630
		v1626 = v1632
		goto L282
	} else {
		goto L284
	}
L283:
	;
	v1642 = v1630
	goto L280
L284:
	;
	goto L283
L285:
	;
	v1661 = v1642
	goto L286
L286:
	;
	v1679 = v1661 + int32(-1)
	v1680 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1679))) = uint8(v1680)
	if base.Ui32(v31+int32(16)) < base.Ui32(v1679) {
		v1661 = v1679
		goto L286
	} else {
		goto L288
	}
L287:
	;
	v1696 = v1679
	goto L272
L288:
	;
	goto L287
L289:
	;
	v1716 = v1576
	goto L291
L290:
	;
	v1716 = v1713
	goto L291
L291:
	;
	F_out(m, l0, v1696, v1716)
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L14
	} else {
		goto L292
	}
L292:
	;
	v1720 = v1576 + int32(-9)
	v1722 = v1580 + int32(4)
	if base.Ui32(v989) <= base.Ui32(v1722) {
		v1968 = v1720
		goto L234
	} else {
		goto L293
	}
L293:
	;
	if int32(9) < v1576 {
		v1576 = v1720
		v1580 = v1722
		goto L270
	} else {
		goto L294
	}
L294:
	;
	v1968 = v1720
	goto L234
L295:
	;
	v1948 = int32(18)
	F_pad(m, l0, int32(48), v1935+v1948, v1948, int32(0))
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L14
	} else {
		goto L332
	}
L296:
	;
	if base.Ui32(v970) < base.Ui32(v989) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1731 = v989
	goto L299
L298:
	;
	v1731 = v970 + int32(4)
	goto L299
L299:
	;
	v1735 = v31 + int32(16) | int32(9)
	v1748 = v970
	v1752 = v1155
	goto L300
L300:
	;
	v1764 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1748))))
	if base.Ui64(int64(4294967296)) <= base.Ui64(v1764) {
		goto L305
	} else {
		goto L306
	}
L301:
	;
	v1935 = v1913
	goto L295
L302:
	;
	if v1748 == v970 {
		goto L317
	} else {
		goto L318
	}
L303:
	;
	if v1818 != v1735 {
		v1828 = v1818
		goto L302
	} else {
		goto L315
	}
L304:
	;
	if v1792 == int64(0) {
		v1818 = v1791
		goto L310
	} else {
		goto L311
	}
L305:
	;
	v1771 = v1764
	v1772 = v1735
	goto L307
L306:
	;
	v1791 = v1735
	v1792 = v1764
	goto L304
L307:
	;
	v1778 = v1772 + int32(-1)
	v1779 = int64(10)
	v1780 = base.I64_div_u_s(v1771, v1779)
	v1786 = base.I32_wrap_i64(v1771-v1780*v1779) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1778))) = uint8(v1786)
	if base.Ui64(int64(42949672959)) < base.Ui64(v1771) {
		v1771 = v1780
		v1772 = v1778
		goto L307
	} else {
		goto L309
	}
L308:
	;
	v1791 = v1778
	v1792 = v1780
	goto L304
L309:
	;
	goto L308
L310:
	;
	goto L303
L311:
	;
	v1800 = v1791
	v1802 = base.I32_wrap_i64(v1792)
	goto L312
L312:
	;
	v1806 = v1800 + int32(-1)
	v1807 = int32(10)
	v1808 = base.I32_div_u_s(v1802, v1807)
	v1813 = v1802 - v1808*v1807 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1806))) = uint8(v1813)
	if base.Ui32(int32(9)) < base.Ui32(v1802) {
		v1800 = v1806
		v1802 = v1808
		goto L312
	} else {
		goto L314
	}
L313:
	;
	v1818 = v1806
	goto L310
L314:
	;
	goto L313
L315:
	;
	v1825 = v1818 + int32(-1)
	v1826 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1825))) = uint8(v1826)
	v1828 = v1825
	goto L302
L316:
	;
	v1908 = v1735 - v1891
	if v1908 < v1752 {
		goto L326
	} else {
		goto L327
	}
L317:
	;
	F_out(m, l0, v1828, int32(1))
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L14
	} else {
		goto L323
	}
L318:
	;
	if base.Ui32(v1828) <= base.Ui32(v31+int32(16)) {
		v1891 = v1828
		goto L316
	} else {
		goto L319
	}
L319:
	;
	v1844 = v1828
	goto L320
L320:
	;
	v1862 = v1844 + int32(-1)
	v1863 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1862))) = uint8(v1863)
	if base.Ui32(v31+int32(16)) < base.Ui32(v1862) {
		v1844 = v1862
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1891 = v1862
	goto L316
L323:
	;
	v1872 = v1828 + int32(1)
	if v1752|v1164 == int32(0) {
		v1891 = v1872
		goto L316
	} else {
		goto L324
	}
L324:
	;
	F_out(m, l0, int32(_a1964), int32(1))
	mBase = m.M
	v1879 = m.ExcPending
	if v1879 != 0 {
		goto L14
	} else {
		goto L325
	}
L325:
	;
	v1891 = v1872
	goto L316
L326:
	;
	v1910 = v1908
	goto L328
L327:
	;
	v1910 = v1752
	goto L328
L328:
	;
	F_out(m, l0, v1891, v1910)
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L14
	} else {
		goto L329
	}
L329:
	;
	v1913 = v1752 - v1908
	v1915 = v1748 + int32(4)
	if base.Ui32(v1731) <= base.Ui32(v1915) {
		v1935 = v1913
		goto L295
	} else {
		goto L330
	}
L330:
	;
	if int32(-1) < v1913 {
		v1748 = v1915
		v1752 = v1913
		goto L300
	} else {
		goto L331
	}
L331:
	;
	goto L301
L332:
	;
	F_out(m, l0, v1357, v97-v1357)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L14
	} else {
		goto L333
	}
L333:
	;
	goto L233
L334:
	;
	goto L233
L335:
	;
	if v1368 < l2 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v2026 = l2
	goto L338
L337:
	;
	v2026 = v1368
	goto L338
L338:
	;
	v2329 = v2026
	goto L11
L339:
	;
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v2112 = v2110 >> (uint(int32(31)) % 32)
	v2115 = base.I64_extend_i32_u(v2110 ^ v2112 - v2112)
	if base.Ui64(int64(4294967296)) <= base.Ui64(v2115) {
		goto L349
	} else {
		goto L350
	}
L340:
	;
	v2052 = int32(12) - l3
	v2068 = float64(16)
	goto L341
L341:
	;
	v2070 = base.F64_mul(v2068, float64(16))
	v2072 = v2052 + int32(-1)
	if v2072 != 0 {
		v2052 = v2072
		v2068 = v2070
		goto L341
	} else {
		goto L343
	}
L342:
	;
	v2073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2035))))
	if v2073 != int32(45) {
		goto L344
	} else {
		goto L345
	}
L343:
	;
	goto L342
L344:
	;
	v2083 = base.F64_sub(base.F64_add(v134, v2070), v2070)
	goto L339
L345:
	;
	v2083 = base.F64_neg(base.F64_add(v2070, base.F64_sub(base.F64_neg(v134), v2070)))
	goto L339
L346:
	;
	v2183 = v59 | int32(2)
	v2187 = v2180 + int32(-2)
	v2189 = l5 + int32(15)
	*(*uint8)(unsafe.Add(mBase, uint32(v2187))) = uint8(v2189)
	if v2181 < int32(0) {
		goto L360
	} else {
		goto L361
	}
L347:
	;
	if v2169 != v97 {
		v2180 = v2169
		v2181 = v2110
		goto L346
	} else {
		goto L359
	}
L348:
	;
	if v2143 == int64(0) {
		v2169 = v2142
		goto L354
	} else {
		goto L355
	}
L349:
	;
	v2122 = v2115
	v2123 = v97
	goto L351
L350:
	;
	v2142 = v97
	v2143 = v2115
	goto L348
L351:
	;
	v2129 = v2123 + int32(-1)
	v2130 = int64(10)
	v2131 = base.I64_div_u_s(v2122, v2130)
	v2137 = base.I32_wrap_i64(v2122-v2131*v2130) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2129))) = uint8(v2137)
	if base.Ui64(int64(42949672959)) < base.Ui64(v2122) {
		v2122 = v2131
		v2123 = v2129
		goto L351
	} else {
		goto L353
	}
L352:
	;
	v2142 = v2129
	v2143 = v2131
	goto L348
L353:
	;
	goto L352
L354:
	;
	goto L347
L355:
	;
	v2151 = v2142
	v2153 = base.I32_wrap_i64(v2143)
	goto L356
L356:
	;
	v2157 = v2151 + int32(-1)
	v2158 = int32(10)
	v2159 = base.I32_div_u_s(v2153, v2158)
	v2164 = v2153 - v2159*v2158 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2157))) = uint8(v2164)
	if base.Ui32(int32(9)) < base.Ui32(v2153) {
		v2151 = v2157
		v2153 = v2159
		goto L356
	} else {
		goto L358
	}
L357:
	;
	v2169 = v2157
	goto L354
L358:
	;
	goto L357
L359:
	;
	v2176 = v2169 + int32(-1)
	v2177 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2176))) = uint8(v2177)
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v2180 = v2176
	v2181 = v2179
	goto L346
L360:
	;
	v2197 = int32(45)
	goto L362
L361:
	;
	v2197 = int32(43)
	goto L362
L362:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2180+int32(-1)))) = uint8(v2197)
	v2209 = v2083
	v2220 = v31 + int32(16)
	goto L363
L363:
	;
	if base.F64_lt(base.F64_abs(v2209), float64(2.147483648e+09)) == int32(0) {
		goto L366
	} else {
		goto L367
	}
L364:
	;
	v2272 = v97 - v2187
	v2273 = v2183 + v2272
	if int32(2147483645)-v2273 < l3 {
		v2329 = int32(-1)
		goto L11
	} else {
		goto L372
	}
L365:
	;
	v2246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2243)+uint32(_consts[1069]))))
	v2247 = v2246 | l5&int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v2220))) = uint8(v2247)
	v2252 = base.F64_mul(base.F64_sub(v2209, base.F64_convert_i32_s(v2243)), float64(16))
	v2253 = int32(1)
	v2254 = v2220 + v2253
	if v2254-(v31+int32(16)) != v2253 {
		v2267 = v2254
		goto L368
	} else {
		goto L369
	}
L366:
	;
	v2243 = int32(-2147483648)
	goto L365
L367:
	;
	v2241 = base.I32_trunc_f64_s(v2209)
	v2243 = v2241
	goto L365
L368:
	;
	if base.F64_ne(v2252, float64(0)) != 0 {
		v2209 = v2252
		v2220 = v2267
		goto L363
	} else {
		goto L371
	}
L369:
	;
	if base.F64_eq(v2252, float64(0))&(base.B2i32(l3 < int32(1))&base.B2i32(l4&int32(8) == int32(0))) != 0 {
		v2267 = v2254
		goto L368
	} else {
		goto L370
	}
L370:
	;
	v2263 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2220)+1)) = uint8(v2263)
	v2267 = v2220 + int32(2)
	goto L368
L371:
	;
	goto L364
L372:
	;
	v2281 = v2267 - (v31 + int32(16))
	if v2281+int32(-2) < l3 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v2285 = l3 + int32(2)
	goto L375
L374:
	;
	v2285 = v2281
	goto L375
L375:
	;
	if l3 != 0 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v2286 = v2285
	goto L378
L377:
	;
	v2286 = v2281
	goto L378
L378:
	;
	v2287 = v2273 + v2286
	F_pad(m, l0, int32(32), l2, v2287, l4)
	mBase = m.M
	v2289 = m.ExcPending
	if v2289 != 0 {
		goto L14
	} else {
		goto L379
	}
L379:
	;
	F_out(m, l0, v2035, v2183)
	mBase = m.M
	v2291 = m.ExcPending
	if v2291 != 0 {
		goto L14
	} else {
		goto L380
	}
L380:
	;
	F_pad(m, l0, int32(48), l2, v2287, l4^int32(65536))
	mBase = m.M
	v2296 = m.ExcPending
	if v2296 != 0 {
		goto L14
	} else {
		goto L381
	}
L381:
	;
	F_out(m, l0, v31+int32(16), v2281)
	mBase = m.M
	v2300 = m.ExcPending
	if v2300 != 0 {
		goto L14
	} else {
		goto L382
	}
L382:
	;
	v2303 = int32(0)
	F_pad(m, l0, int32(48), v2286-v2281, v2303, v2303)
	mBase = m.M
	v2306 = m.ExcPending
	if v2306 != 0 {
		goto L14
	} else {
		goto L383
	}
L383:
	;
	F_out(m, l0, v2187, v2272)
	mBase = m.M
	v2308 = m.ExcPending
	if v2308 != 0 {
		goto L14
	} else {
		goto L384
	}
L384:
	;
	F_pad(m, l0, int32(32), l2, v2287, l4^int32(8192))
	mBase = m.M
	v2313 = m.ExcPending
	if v2313 != 0 {
		goto L14
	} else {
		goto L385
	}
L385:
	;
	if v2287 < l2 {
		goto L386
	} else {
		goto L387
	}
L386:
	;
	v2315 = l2
	goto L388
L387:
	;
	v2315 = v2287
	goto L388
L388:
	;
	v2329 = v2315
	goto L11
}
func F_fmt_x(m *base.Module, l0 int64, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	if l0 == int64(0) {
		v26 = l1
	} else {
		v7 = l0
		v8 = l1
		for {
			v12 = v8 + int32(-1)
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v7)&int32(15))+uint32(_consts[1069]))))
			v19 = v18 | l2
			*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v19)
			if base.Ui64(int64(15)) < base.Ui64(v7) {
				v7 = int64(base.Ui64(v7) >> (uint(int64(4)) % 64))
				v8 = v12
				continue
			} else {
				break
			}
			break
		}
		v26 = v12
	}
	return v26
}
func F_fopen(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	v12 = F___strchrnul(m, int32(_a2763), v11)
	mBase = m.M
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v14 == v11&int32(255) {
		v18 = v12
	} else {
		v18 = int32(0)
	}
	if v18 != 0 {
		v25 = F_strchr(m, l1, int32(43))
		mBase = m.M
		if v25 != 0 {
			v29 = int32(2)
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v29 = base.B2i32(v26 != int32(114))
		}
		v33 = F_strchr(m, l1, int32(120))
		mBase = m.M
		if v33 != 0 {
			v34 = v29 | int32(128)
		} else {
			v34 = v29
		}
		v38 = F_strchr(m, l1, int32(101))
		mBase = m.M
		if v38 != 0 {
			v39 = v34 | int32(524288)
		} else {
			v39 = v34
		}
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v42 == int32(114) {
			v45 = v39
		} else {
			v45 = v39 | int32(64)
		}
		if v42 == int32(119) {
			v50 = v45 | int32(512)
		} else {
			v50 = v45
		}
		if v42 == int32(97) {
			v55 = v50 | int32(1024)
		} else {
			v55 = v50
		}
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(438)
		v62 = m.Env.X__syscall_openat(m, int32(-100), l0, v55|int32(32768), v8)
		mBase = m.M
		if base.Ui32(v62) < base.Ui32(int32(-4095)) {
			v70 = v62
		} else {
			v65 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v65))) = int32(0) - v62
			v70 = int32(-1)
		}
		if v70 < int32(0) {
			v81 = int32(0)
		} else {
			v73 = F___fdopen(m, v70, l1)
			mBase = m.M
			if v73 != 0 {
				v81 = v73
			} else {
				v74 = m.Wasi_snapshot_preview1.Fd_close(m, v70)
				mBase = m.M
				v81 = int32(0)
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(28)
		v81 = int32(0)
	}
	m.G0 = v8 + int32(16)
	return v81
}
func F_fp_barrier_3(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = m.G0
	*(*float64)(unsafe.Add(mBase, uint32(v3-int32(16))+8)) = l0
	return l0
}
func F_fp_barrier_4(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = m.G0
	*(*float64)(unsafe.Add(mBase, uint32(v3-int32(16))+8)) = l0
	return l0
}
func F_fpconv_g_fmt(m *base.Module, l0 int32, l1 float64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v8 = m.G0
	v10 = v8 - int32(64)
	m.G0 = v10
	if base.Ui32(int32(14)) <= base.Ui32(l2+int32(-1)) {
		v96 = m.G3
		m.Env.X__assert_fail(m, v96+int32(_a2735), v96+int32(_a2736), int32(157), v96+int32(_a2737))
		mBase = m.M
		base.Wasm_trap_unreachable()
		for {
		}
	} else {
		v16 = int32(11813)
		*(*uint16)(unsafe.Add(mBase, uint32(v10)+26)) = uint16(v16)
		v19 = base.B2i32(base.Ui32(l2) < base.Ui32(int32(10)))
		if v19 == int32(0) {
			v26 = int32(49)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)) = uint8(v26)
			v28 = v10 + int32(29)
		} else {
			v28 = v10 + int32(28)
		}
		v29 = int32(103)
		*(*uint16)(unsafe.Add(mBase, uint32(v28)+1)) = uint16(v29)
		if base.Ui32(l2) < base.Ui32(int32(10)) {
			v33 = l2
		} else {
			v33 = l2 + int32(246)
		}
		v35 = v33 | int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v35)
		v37 = m.G3
		v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[993]))))
		if v40 != int32(46) {
			*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = l1
			v52 = m.G3
			v53 = int32(32)
			v60 = F_snprintf(m, v10+v53, v53, v10+int32(26), v10+int32(16))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_consts[993]))))
				v69 = l0
				v71 = v10 + int32(32)
				for {
					v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
					if v77 == v66&int32(255) {
						v79 = int32(46)
					} else {
						v79 = v77
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v79)
					v81 = int32(1)
					if v77 != 0 {
						v69 = v69 + v81
						v71 = v71 + v81
						continue
					} else {
						break
					}
					break
				}
				v91 = v60
				m.G0 = v10 + int32(64)
				return v91
			}
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v10))) = l1
			v47 = F_snprintf(m, l0, int32(32), v10+int32(26), v10)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				v91 = v47
				m.G0 = v10 + int32(64)
				return v91
			}
		}
	}
}
func F_fputc(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_do_putc_1(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_fread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	if int32(0) <= v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v17 = l2 * l1
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+72)) = v18 + int32(-1) | v18
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v23 != v24 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L4
L3:
	;
	v16 = int32(1)
	goto L1
L4:
	;
	v16 = int32(0)
	goto L1
L5:
	;
	if v39 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v26 = v24 - v23
	if base.Ui32(v26) < base.Ui32(v17) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v38 = l0
	v39 = v17
	goto L5
L8:
	;
	v28 = v26
	goto L10
L9:
	;
	v28 = v17
	goto L10
L10:
	;
	if v28 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v33 + v28
	v38 = l0 + v28
	v39 = v17 - v28
	goto L5
L12:
	;
	goto L11
L13:
	;
	v31 = F__emscripten_memcpy_bulkmem(m, l0, v23, v28)
	mBase = m.M
	goto L12
L14:
	;
	if l1 != 0 {
		goto L29
	} else {
		goto L30
	}
L15:
	;
	v43 = v38
	v49 = v39
	goto L16
L16:
	;
	v51 = F___toread(m, l3)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L14
L18:
	;
	v63 = v49 - v56
	if v63 != 0 {
		v43 = v43 + v56
		v49 = v63
		goto L16
	} else {
		goto L28
	}
L19:
	;
	if v16 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	return int32(0)
L21:
	;
	if v51 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v56 = m.T0[v55].(func(*base.Module, int32, int32, int32) int32)(m, l3, v43, v49)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	if v56 != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	goto L19
L25:
	;
	v60 = base.I32_div_u_s(v17-v49, l1)
	return v60
L26:
	;
	goto L27
L27:
	;
	goto L25
L28:
	;
	goto L17
L29:
	;
	v73 = l2
	goto L31
L30:
	;
	v73 = int32(0)
	goto L31
L31:
	;
	if v16 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return v73
L33:
	;
	goto L34
L34:
	;
	goto L32
}
func F_freeEvalScripts(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	F_dictRelease(m, l0)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_listRelease(m, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l2 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v11 = F_listGetIterator(m, l2, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L8
L7:
	;
	F_listReleaseIterator(m, v11)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L17
	}
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v17 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v17 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v17+base.B2i32(v20 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v26
	goto L11
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v30 == int32(0) {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	m.T0[v34].(func(*base.Module, int32))(m, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_valkey_free(m, v30)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L8
L17:
	;
	F_listRelease(m, l2)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L4
}
func F_freeLogEntry(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_decrRefCount(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_freePendingReplDataBufAsync(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v8) < base.Ui32(int32(65)) {
		F_listRelease(m, l0)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	} else {
		v11 = int32(0)
		v13 = *(*int32)(unsafe.Add(mBase, _consts[373]))
		*(*int32)(unsafe.Add(mBase, _consts[373])) = v13 + v8
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
		F_bioCreateLazyFreeJob(m, int32(559), int32(1), v6)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	}
}
func F_freelist(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var __phi77 int32
	_ = __phi77
	var v78 int32
	_ = v78
	var __phi78 int32
	_ = __phi78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var __phi244 int32
	_ = __phi244
	var v245 int32
	_ = v245
	var __phi245 int32
	_ = __phi245
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v326 int32
	_ = v326
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v6 = v3
	goto L3
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v6 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L1
L5:
	;
	if v8 != 0 {
		v6 = v8
		goto L3
	} else {
		goto L105
	}
L6:
	;
	goto L5
L7:
	;
	v19 = int32(-8)
	v20 = v6 + v19
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-4))))
	v25 = v23 & v19
	v26 = v20 + v25
	if v23&int32(1) != 0 {
		v150 = v25
		v151 = v20
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if base.Ui32(v26) <= base.Ui32(v151) {
		goto L6
	} else {
		goto L43
	}
L9:
	;
	if v23&int32(2) == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v34 = v20 - v33
	v36 = *(*int32)(unsafe.Add(mBase, _consts[994]))
	if base.Ui32(v34) < base.Ui32(v36) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v38 = v33 + v25
	v40 = *(*int32)(unsafe.Add(mBase, _consts[995]))
	if v34 == v40 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v56 == int32(0) {
		v150 = v38
		v151 = v34
		goto L8
	} else {
		goto L31
	}
L13:
	;
	v109 = int32(0)
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v45
	v150 = v38
	v151 = v34
	goto L8
L15:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v91 = int32(3)
	if v90&v91 != v91 {
		v150 = v38
		v151 = v34
		goto L8
	} else {
		goto L30
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if base.Ui32(int32(255)) < base.Ui32(v33) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	if v42 == v34 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	if v42 != v45 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v47 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, _consts[996]))
	*(*int32)(unsafe.Add(mBase, _consts[996])) = v49 & base.I32_rotl(int32(-2), int32(base.Ui32(v33)>>(uint(int32(3))%32)))
	v150 = v38
	v151 = v34
	goto L8
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	if v61 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v58
	v109 = v42
	goto L12
L22:
	;
	__phi77 = v71
	__phi78 = v72
	v77 = __phi77
	v78 = __phi78
	goto L26
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	if v66 == int32(0) {
		goto L13
	} else {
		goto L25
	}
L24:
	;
	v71 = v61
	v72 = v34 + int32(20)
	goto L22
L25:
	;
	v71 = v66
	v72 = v34 + int32(16)
	goto L22
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	if v84 != 0 {
		__phi77 = v84
		__phi78 = v77 + int32(20)
		v77 = __phi77
		v78 = __phi78
		goto L26
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(0)
	v109 = v77
	goto L12
L28:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	if v87 != 0 {
		__phi77 = v87
		__phi78 = v77 + int32(16)
		v77 = __phi77
		v78 = __phi78
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	*(*int32)(unsafe.Add(mBase, _consts[997])) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v90 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v38 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v38
	goto L5
L31:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
	v120 = v118 << (uint(int32(2)) % 32)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)+uint32(_consts[998])))
	if v34 != v123 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = v56
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	if v140 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	if v133 != v34 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+uint32(_consts[998]))) = v109
	if v109 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v126 = int32(0)
	v128 = *(*int32)(unsafe.Add(mBase, _consts[999]))
	*(*int32)(unsafe.Add(mBase, _consts[999])) = v128 & base.I32_rotl(int32(-2), v118)
	v150 = v38
	v151 = v34
	goto L8
L36:
	;
	if v109 == int32(0) {
		v150 = v38
		v151 = v34
		goto L8
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = v109
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+16)) = v109
	goto L36
L39:
	;
	goto L32
L40:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	if v145 == int32(0) {
		v150 = v38
		v151 = v34
		goto L8
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v140)+24)) = v109
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v145)+24)) = v109
	v150 = v38
	v151 = v34
	goto L8
L43:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v160&int32(1) == int32(0) {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	if v160&int32(2) != 0 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	if base.Ui32(int32(255)) < base.Ui32(v326) {
		goto L83
	} else {
		goto L84
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v206 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v151+v206))) = v206
	if v151 != v190 {
		v326 = v206
		goto L45
	} else {
		goto L82
	}
L47:
	;
	if v223 == int32(0) {
		goto L46
	} else {
		goto L70
	}
L48:
	;
	v268 = int32(0)
	goto L47
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v160 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v150 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v151+v150))) = v150
	v326 = v150
	goto L45
L50:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _consts[1000]))
	if v26 != v168 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _consts[995]))
	if v26 != v190 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v170 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1000])) = v151
	v174 = *(*int32)(unsafe.Add(mBase, _consts[1001]))
	v175 = v174 + v150
	*(*int32)(unsafe.Add(mBase, _consts[1001])) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v175 | int32(1)
	v181 = *(*int32)(unsafe.Add(mBase, _consts[995]))
	if v151 != v181 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	v183 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[997])) = v183
	*(*int32)(unsafe.Add(mBase, _consts[995])) = v183
	goto L5
L54:
	;
	v206 = v160&int32(-8) + v150
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if base.Ui32(int32(255)) < base.Ui32(v160) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v192 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[995])) = v151
	v196 = *(*int32)(unsafe.Add(mBase, _consts[997]))
	v197 = v196 + v150
	*(*int32)(unsafe.Add(mBase, _consts[997])) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v197 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v151+v197))) = v197
	goto L5
L56:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v207 == v26 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v207 != v210 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210)+12)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v207)+8)) = v210
	goto L46
L59:
	;
	v212 = int32(0)
	v214 = *(*int32)(unsafe.Add(mBase, _consts[996]))
	*(*int32)(unsafe.Add(mBase, _consts[996])) = v214 & base.I32_rotl(int32(-2), int32(base.Ui32(v160)>>(uint(int32(3))%32)))
	goto L46
L60:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v228 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v225)+12)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v207)+8)) = v225
	v268 = v207
	goto L47
L62:
	;
	__phi244 = v238
	__phi245 = v239
	v244 = __phi244
	v245 = __phi245
	goto L66
L63:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v233 == int32(0) {
		goto L48
	} else {
		goto L65
	}
L64:
	;
	v238 = v228
	v239 = v26 + int32(20)
	goto L62
L65:
	;
	v238 = v233
	v239 = v26 + int32(16)
	goto L62
L66:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v244)+20))
	if v251 != 0 {
		__phi244 = v251
		__phi245 = v244 + int32(20)
		v244 = __phi244
		v245 = __phi245
		goto L66
	} else {
		goto L68
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = int32(0)
	v268 = v244
	goto L47
L68:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v244)+16))
	if v254 != 0 {
		__phi244 = v254
		__phi245 = v244 + int32(16)
		v244 = __phi244
		v245 = __phi245
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v279 = v277 << (uint(int32(2)) % 32)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v279)+uint32(_consts[998])))
	if v26 != v282 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268)+24)) = v223
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v299 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	if v292 != v26 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+uint32(_consts[998]))) = v268
	if v268 != 0 {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v285 = int32(0)
	v287 = *(*int32)(unsafe.Add(mBase, _consts[999]))
	*(*int32)(unsafe.Add(mBase, _consts[999])) = v287 & base.I32_rotl(int32(-2), v277)
	goto L46
L75:
	;
	if v268 == int32(0) {
		goto L46
	} else {
		goto L78
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+20)) = v268
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+16)) = v268
	goto L75
L78:
	;
	goto L71
L79:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v304 == int32(0) {
		goto L46
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268)+16)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v299)+24)) = v268
	goto L79
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268)+20)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v304)+24)) = v268
	goto L46
L82:
	;
	*(*int32)(unsafe.Add(mBase, _consts[997])) = v206
	goto L5
L83:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v326) {
		v373 = int32(31)
		goto L88
	} else {
		goto L89
	}
L84:
	;
	v338 = v326 & int32(-8)
	v340 = v338 + int32(9128464)
	v342 = *(*int32)(unsafe.Add(mBase, _consts[996]))
	v346 = int32(1) << (uint(int32(base.Ui32(v326)>>(uint(int32(3))%32))) % 32)
	if v342&v346 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338)+uint32(_consts[1002]))) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v352)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v151)+12)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = v352
	goto L5
L86:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v338)+uint32(_consts[1002])))
	v352 = v351
	goto L85
L87:
	;
	*(*int32)(unsafe.Add(mBase, _consts[996])) = v342 | v346
	v352 = v340
	goto L85
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+28)) = v373
	*(*int64)(unsafe.Add(mBase, uint32(v151)+16)) = int64(0)
	v378 = v373 << (uint(int32(2)) % 32)
	v382 = *(*int32)(unsafe.Add(mBase, _consts[999]))
	v384 = int32(1) << (uint(v373) % 32)
	if v382&v384 != 0 {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	v363 = base.I32_clz(int32(base.Ui32(v326) >> (uint(int32(8)) % 32)))
	v366 = int32(1)
	v373 = int32(base.Ui32(v326)>>(uint(int32(38)-v363)%32))&v366 - v363<<(uint(v366)%32) + int32(62)
	goto L88
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151+v445))) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v151)+12)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v151+v443))) = v446
	v457 = int32(0)
	v459 = *(*int32)(unsafe.Add(mBase, _consts[1003]))
	v460 = int32(-1)
	v461 = v459 + v460
	if v461 != 0 {
		goto L102
	} else {
		goto L103
	}
L91:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v407)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v437)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v407)+8)) = v151
	v443 = int32(24)
	v445 = int32(8)
	v446 = int32(0)
	v447 = v407
	v448 = v437
	goto L90
L92:
	;
	v443 = v428
	v445 = v430
	v446 = v151
	v447 = v151
	v448 = v433
	goto L90
L93:
	;
	if v373 == int32(31) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, _consts[999])) = v382 | v384
	*(*int32)(unsafe.Add(mBase, uint32(v378)+uint32(_consts[998]))) = v151
	v428 = int32(8)
	v430 = int32(24)
	v433 = v378 + int32(9128728)
	goto L92
L95:
	;
	v399 = int32(0)
	goto L97
L96:
	;
	v399 = int32(25) - int32(base.Ui32(v373)>>(uint(int32(1))%32))
	goto L97
L97:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v378)+uint32(_consts[998])))
	v404 = v326 << (uint(v399) % 32)
	v407 = v401
	goto L98
L98:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	if v411&int32(-8) == v326 {
		goto L91
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v421+int32(16)))) = v151
	v428 = int32(8)
	v430 = int32(24)
	v433 = v407
	goto L92
L100:
	;
	v421 = v407 + int32(base.Ui32(v404)>>(uint(int32(29))%32))&int32(4)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)+16))
	if v422 != 0 {
		v404 = v404 << (uint(int32(1)) % 32)
		v407 = v422
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v463 = v461
	goto L104
L103:
	;
	v463 = v460
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1003])) = v463
	goto L6
L105:
	;
	goto L4
}
func F_fsync(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = m.Wasi_snapshot_preview1.Fd_sync(m, l0)
	mBase = m.M
	if v2 != 0 {
		v4 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = v2
		v7 = int32(-1)
	} else {
		v7 = int32(0)
	}
	return v7
}
func F_ftruncate(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v3 = m.Env.X__syscall_ftruncate64(m, l0, l1)
	mBase = m.M
	if base.Ui32(v3) < base.Ui32(int32(-4095)) {
		v11 = v3
	} else {
		v6 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(0) - v3
		v11 = int32(-1)
	}
	return v11
}
