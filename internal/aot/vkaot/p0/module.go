package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_ModuleForkDoneHandler(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v10 {
		v22 = int32(0)
		v23 = *(*int32)(unsafe.Add(mBase, _consts[391]))
		if v23 == v22 {
			v30 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[392])) = v30
			*(*int32)(unsafe.Add(mBase, _consts[391])) = v30
			m.G0 = v7 + int32(16)
			return
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, _consts[392]))
			m.T0[v23].(func(*base.Module, int32, int32, int32))(m, l0, l1, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v30 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[392])) = v30
				*(*int32)(unsafe.Add(mBase, _consts[391])) = v30
				m.G0 = v7 + int32(16)
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
		v16 = *(*int32)(unsafe.Add(mBase, _consts[39]))
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v16
		F__serverLog(m, int32(2), int32(_a735), v7)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v22 = int32(0)
			v23 = *(*int32)(unsafe.Add(mBase, _consts[391]))
			if v23 == v22 {
				v30 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[392])) = v30
				*(*int32)(unsafe.Add(mBase, _consts[391])) = v30
				m.G0 = v7 + int32(16)
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, _consts[392]))
				m.T0[v23].(func(*base.Module, int32, int32, int32))(m, l0, l1, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v30 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[392])) = v30
					*(*int32)(unsafe.Add(mBase, _consts[391])) = v30
					m.G0 = v7 + int32(16)
					return
				}
			}
		}
	}
}
func F_addModuleBoolConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_sdsempty(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
		v17 = F_sdscatfmt(m, v12, int32(_a465), v10)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v20 = F_valkey_malloc(m, int32(88))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v22
				*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = l4
				v25 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = int32(416)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(417)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(418)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(419)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = l2 | int32(256)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v17
				*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = l3
				*(*int64)(unsafe.Add(mBase, uint32(v20+int32(48)))) = v22
				*(*int64)(unsafe.Add(mBase, uint32(v20+int32(56)))) = v22
				*(*int64)(unsafe.Add(mBase, uint32(v20+int32(64)))) = v22
				*(*int64)(unsafe.Add(mBase, uint32(v20+int32(72)))) = v22
				*(*int32)(unsafe.Add(mBase, uint32(v20+int32(80)))) = v25
				v65 = *(*int32)(unsafe.Add(mBase, _consts[171]))
				v66 = F_sdsnew(m, v17)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					v68 = F_dictAdd(m, v65, v66, v20)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						m.G0 = v10 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_addModuleNumericConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64, l7 int64) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = F_sdsempty(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
		v20 = F_sdscatfmt(m, v15, int32(_a465), v13)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v23 = F_valkey_malloc(m, int32(88))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = int32(1)
				*(*int64)(unsafe.Add(mBase, uint32(v23)+72)) = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v23)+64)) = l4
				*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = l7
				*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = l6
				*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = int64(4)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = l5
				v36 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v36
				*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = int32(428)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(429)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v36
				*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(430)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(431)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = l2 | int32(256)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v36
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = v20
				v55 = *(*int32)(unsafe.Add(mBase, _consts[171]))
				v56 = F_sdsnew(m, v20)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					v58 = F_dictAdd(m, v55, v56, v23)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						m.G0 = v13 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_addModuleStringConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_sdsempty(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
		v17 = F_sdscatfmt(m, v12, int32(_a465), v10)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v20 = F_valkey_malloc(m, int32(88))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v22
				*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = l4
				v25 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = int32(420)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(421)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(422)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(423)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = l2 | int32(256)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v17
				*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = int32(3)
				*(*int64)(unsafe.Add(mBase, uint32(v20+int32(48)))) = v22
				*(*int64)(unsafe.Add(mBase, uint32(v20+int32(56)))) = v22
				*(*int64)(unsafe.Add(mBase, uint32(v20+int32(64)))) = v22
				*(*int64)(unsafe.Add(mBase, uint32(v20+int32(72)))) = v22
				v63 = *(*int32)(unsafe.Add(mBase, _consts[171]))
				v64 = F_sdsnew(m, v17)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					v66 = F_dictAdd(m, v63, v64, v20)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						m.G0 = v10 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_addModuleUnsignedNumericConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64, l7 int64) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = F_sdsempty(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
		v20 = F_sdscatfmt(m, v15, int32(_a465), v13)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v23 = F_valkey_malloc(m, int32(88))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = int32(1)
				*(*int64)(unsafe.Add(mBase, uint32(v23)+72)) = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v23)+64)) = l4
				*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = l7
				*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = l6
				*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = int64(5)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = l5
				v36 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v36
				*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = int32(428)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(429)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v36
				*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(430)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(431)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = l2 | int32(256)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v36
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = v20
				v55 = *(*int32)(unsafe.Add(mBase, _consts[171]))
				v56 = F_sdsnew(m, v20)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					v58 = F_dictAdd(m, v55, v56, v23)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						m.G0 = v13 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_getModuleNumericConfig(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = m.T0[v4].(func(*base.Module, int32, int32) int64)(m, v2, v3)
	mBase = m.M
	return v5
}
func F_getModuleStringConfig(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = m.T0[v4].(func(*base.Module, int32, int32) int32)(m, v2, v3)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v11 = F_objectGetVal(m, v5)
			mBase = m.M
			v12 = F_sdsdup(m, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12
			}
		} else {
			return int32(0)
		}
	}
}
func F_isModuleClientUnblocked(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+48))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+48))
	return base.B2i32(v4 == int32(1))
}
func F_loadModuleConfigs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v145 int32
	_ = v145
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(64)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v2
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v16 = v9 + int32(56)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v17
	goto L1
L1:
	;
	v22 = v9 + int32(56)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v24 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	m.G0 = v9 + int32(64)
	return v145
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(1)
	v145 = v2
	goto L2
L4:
	;
	if v24 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24+base.B2i32(v27 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v33
	goto L5
L7:
	;
	v40 = v24
	goto L8
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v44 = F_sdsempty(m)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L3
L10:
	;
	return int32(0)
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v48
	v55 = F_sdscatfmt(m, v44, int32(_a465), v9+int32(32))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[399]))
	v59 = F_dictFind(m, v58, v55)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _consts[399]))
	v117 = F_dictDelete(m, v116, v55)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L10
	} else {
		goto L34
	}
L14:
	;
	v93 = F_performModuleConfigSetDefaultFromName(m, v55, v9+int32(52))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L10
	} else {
		goto L27
	}
L15:
	;
	if v59 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	goto L17
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	goto L18
L18:
	;
	v67 = F_performModuleConfigSetFromName(m, v63, v64, v9+int32(52))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	if v67 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v70 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_sdsfree(m, v55)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L10
	} else {
		goto L25
	}
L22:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v75
	F__serverLog(m, int32(3), int32(_a753), v9+int32(16))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[399]))
	F_dictEmpty(m, v86, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v145 = int32(1)
	goto L2
L27:
	;
	if v93 != 0 {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v96 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_sdsfree(m, v55)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L10
	} else {
		goto L32
	}
L30:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v101
	F__serverLog(m, int32(3), int32(_a754), v9)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[399]))
	F_dictEmpty(m, v110, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	v145 = int32(1)
	goto L2
L34:
	;
	F_sdsfree(m, v55)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	v122 = v9 + int32(56)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v124 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v124 != 0 {
		v40 = v124
		goto L8
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v124+base.B2i32(v127 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v133
	goto L37
L39:
	;
	goto L9
}
func F_moduleAllDatatypesHandleErrors(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	v4 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	v5 = F_dictGetIterator(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v5)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L2
	} else {
		goto L65
	}
L2:
	;
	return int32(0)
L3:
	;
	v16 = v5 + int32(20)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v112 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v23 = v16
	v24 = v20
	goto L8
L6:
	;
	v20 = int32(1)
	goto L5
L7:
	;
	v20 = int32(0)
	goto L5
L8:
	;
	switch v24 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v24 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v104
	if v104 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v28 != int32(-1) {
		v67 = v28
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v68 = int32(1)
	v69 = v67 + v68
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v69
	v71 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v75+int32(26)))))
	if v79 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if v32 != 0 {
		v67 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	if v34 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	if v61 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v41 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v33)+16)))
	v42 = int64(*(*int8)(unsafe.Add(mBase, uint32(v33)+27)))
	v43 = int64(*(*int32)(unsafe.Add(mBase, uint32(v33)+8)))
	v44 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v33)+12)))
	v45 = int64(*(*int8)(unsafe.Add(mBase, uint32(v33)+26)))
	v46 = int64(*(*int32)(unsafe.Add(mBase, uint32(v33)+4)))
	v47 = F_wangHash64(m, v46)
	mBase = m.M
	v49 = F_wangHash64(m, v45+v47)
	mBase = m.M
	v51 = F_wangHash64(m, v44+v49)
	mBase = m.M
	v53 = F_wangHash64(m, v43+v51)
	mBase = m.M
	v55 = F_wangHash64(m, v42+v53)
	mBase = m.M
	v57 = F_wangHash64(m, v41+v55)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v60 = v59
	goto L17
L19:
	;
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+24)))
	v39 = v37 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+24)) = uint16(v39)
	v60 = v33
	goto L17
L20:
	;
	v67 = v61 + int32(-1)
	goto L14
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v67 = v64
	goto L14
L22:
	;
	v94 = int32(2)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v74+v92<<(uint(v94)%32)+int32(4))))
	v23 = v99 + v93<<(uint(v94)%32)
	v24 = int32(1)
	goto L8
L23:
	;
	v83 = v71
	goto L25
L24:
	;
	v83 = v68 << (uint(v79) % 32)
	goto L25
L25:
	;
	if v69 < v83 {
		v92 = v75
		v93 = v69
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v75 != 0 {
		v112 = v71
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	if v85 == int32(-1) {
		v112 = v71
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5)+4)) = int64(4294967296)
	v92 = int32(1)
	v93 = int32(0)
	goto L22
L29:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v108
	v112 = v104
	goto L11
L30:
	;
	v119 = v112
	goto L31
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
	goto L34
L32:
	;
	goto L1
L33:
	;
	v139 = v5 + int32(20)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	if v140 != 0 {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+20))
	if v122 == int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+48)))
	if v125&int32(1) != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	F_dictReleaseIterator(m, v5)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	return int32(0)
L38:
	;
	if v235 != 0 {
		v119 = v235
		goto L31
	} else {
		goto L64
	}
L39:
	;
	v146 = v139
	v147 = v143
	goto L42
L40:
	;
	v143 = int32(1)
	goto L39
L41:
	;
	v143 = int32(0)
	goto L39
L42:
	;
	switch v147 {
	case 0:
		goto L47
	default:
		goto L46
	}
L44:
	;
	v147 = int32(0)
	goto L42
L45:
	;
	goto L38
L46:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v227
	if v227 == int32(0) {
		goto L44
	} else {
		goto L63
	}
L47:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v151 != int32(-1) {
		v190 = v151
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v191 = int32(1)
	v192 = v190 + v191
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v192
	v194 = int32(0)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197+v198+int32(26)))))
	if v202 == int32(255) {
		goto L57
	} else {
		goto L58
	}
L49:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if v155 != 0 {
		v190 = int32(-1)
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	if v157 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+20))
	if v184 != int32(-1) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v164 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v156)+16)))
	v165 = int64(*(*int8)(unsafe.Add(mBase, uint32(v156)+27)))
	v166 = int64(*(*int32)(unsafe.Add(mBase, uint32(v156)+8)))
	v167 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v156)+12)))
	v168 = int64(*(*int8)(unsafe.Add(mBase, uint32(v156)+26)))
	v169 = int64(*(*int32)(unsafe.Add(mBase, uint32(v156)+4)))
	v170 = F_wangHash64(m, v169)
	mBase = m.M
	v172 = F_wangHash64(m, v168+v170)
	mBase = m.M
	v174 = F_wangHash64(m, v167+v172)
	mBase = m.M
	v176 = F_wangHash64(m, v166+v174)
	mBase = m.M
	v178 = F_wangHash64(m, v165+v176)
	mBase = m.M
	v180 = F_wangHash64(m, v164+v178)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = v180
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v183 = v182
	goto L51
L53:
	;
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+24)))
	v162 = v160 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+24)) = uint16(v162)
	v183 = v156
	goto L51
L54:
	;
	v190 = v184 + int32(-1)
	goto L48
L55:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v190 = v187
	goto L48
L56:
	;
	v217 = int32(2)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v197+v215<<(uint(v217)%32)+int32(4))))
	v146 = v222 + v216<<(uint(v217)%32)
	v147 = int32(1)
	goto L42
L57:
	;
	v206 = v194
	goto L59
L58:
	;
	v206 = v191 << (uint(v202) % 32)
	goto L59
L59:
	;
	if v192 < v206 {
		v215 = v198
		v216 = v192
		goto L56
	} else {
		goto L60
	}
L60:
	;
	if v198 != 0 {
		v235 = v194
		goto L45
	} else {
		goto L61
	}
L61:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v197)+20))
	if v208 == int32(-1) {
		v235 = v194
		goto L45
	} else {
		goto L62
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5)+4)) = int64(4294967296)
	v215 = int32(1)
	v216 = int32(0)
	goto L56
L63:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v231
	v235 = v227
	goto L45
L64:
	;
	goto L32
L65:
	;
	return int32(1)
}
func F_moduleBlockClient(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int64
	_ = v185
	var v188 int32
	_ = v188
	var v200 int64
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v212 int64
	_ = v212
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+200))
	if v19&int32(17) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v44&int32(192) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	goto L11
L3:
	;
	if v19&int32(2) == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+216))
	if v32 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L5:
	;
	if v19&int32(262144) != 0 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	if v19&int32(262148) == int32(4) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L2
L8:
	;
	goto L4
L9:
	;
	goto L10
L10:
	;
	goto L2
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(138)
	return int32(0)
L12:
	;
	v54 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	v57 = base.B2i32(v55 != v54)
	goto L15
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(28)
	return int32(0)
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+205)))
	if v60&int32(16) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	F__serverAssert(m, int32(_a725), int32(_a694), int32(8350))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L24
	} else {
		goto L69
	}
L17:
	;
	v69 = int32(0)
	v70 = base.B2i32(v57|v59 != v69)
	v72 = v44 & int32(1024)
	if v72 == v69 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v57|v59 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	F_initClientBlockingState(m, v18)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	if v70 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(28)
	return int32(0)
L24:
	;
	return int32(0)
L25:
	;
	v87 = F_valkey_malloc(m, int32(72))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v18)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+48)) = v87
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+52)) = v92 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = l4
	v98 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+20)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v87)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v87)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v91
	if v57|v59 != v69 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v105 = v98
	goto L29
L28:
	;
	v105 = v18
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v105
	v107 = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, _consts[345]))
	if v108 == v107 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+36)) = v141
	if v142 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v127 = F_createClient(m, int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L24
	} else {
		goto L34
	}
L32:
	;
	v111 = int32(0)
	v113 = v108 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[345])) = v113
	v116 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116+v113<<(uint(int32(2))%32))))
	v122 = *(*int32)(unsafe.Add(mBase, _consts[349]))
	if base.Ui32(v122) <= base.Ui32(v113) {
		v141 = v120
		v142 = v113
		goto L30
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, _consts[349])) = v113
	v141 = v120
	v142 = v113
	goto L30
L34:
	;
	v129 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v127)+328)) = v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+200)) = v131 | int32(1073741824)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v127)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+204)) = v135 | int32(268435456)
	v140 = *(*int32)(unsafe.Add(mBase, _consts[345]))
	v141 = v127
	v142 = v140
	goto L30
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v177 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v162 = F_createClient(m, int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L24
	} else {
		goto L39
	}
L37:
	;
	v146 = int32(0)
	v148 = v142 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[345])) = v148
	v151 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151+v148<<(uint(int32(2))%32))))
	v157 = *(*int32)(unsafe.Add(mBase, _consts[349]))
	if base.Ui32(v157) <= base.Ui32(v148) {
		v175 = v155
		goto L35
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, _consts[349])) = v148
	v175 = v155
	goto L35
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+328)) = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v162)+200)) = v166 | int32(1073741824)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v162)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v162)+204)) = v170 | int32(268435456)
	v175 = v162
	goto L35
L40:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v18)+96))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+28))
	v185 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v87)+56)) = v185
	v188 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+48)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v87)+44)) = base.B2i32(l6 != v188)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v184
	*(*int64)(unsafe.Add(mBase, uint32(v87+int32(64)))) = v185
	if l5 == v185 {
		v212 = v185
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v87)+36))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+224)))
	*(*uint8)(unsafe.Add(mBase, uint32(v180)+224)) = uint8(v181)
	goto L40
L42:
	;
	if v70 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v200 = F_mstime(m)
	mBase = m.M
	if l5 <= v200^int64(9223372036854775807) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v212 = v200 + l5
	goto L42
L45:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v18)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v204)+48)) = int32(0)
	F_addReplyError(m, v18, int32(_a724))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L24
	} else {
		goto L46
	}
L46:
	;
	return v87
L47:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v224&int32(4) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v18)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+48)) = int32(0)
	if v55 != v54 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v220 = int32(_a722)
	goto L51
L50:
	;
	v220 = int32(_a723)
	goto L51
L51:
	;
	F_addReplyError(m, v18, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L24
	} else {
		goto L52
	}
L52:
	;
	return v87
L53:
	;
	if l2 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v18)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v229)+48)) = int32(0)
	F_addReplyError(m, v18, int32(_a721))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L24
	} else {
		goto L55
	}
L55:
	;
	return v87
L56:
	;
	if l6 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	if v236 == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	if v239 == int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v18)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v242)+48)) = int32(0)
	F_addReplyError(m, v18, int32(_a720))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L24
	} else {
		goto L60
	}
L60:
	;
	return v87
L61:
	;
	if v72 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v18)+116))
	*(*int64)(unsafe.Add(mBase, uint32(v257)+8)) = v212
	F_blockClient(m, v18, int32(3))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L24
	} else {
		goto L65
	}
L63:
	;
	F_blockForKeys(m, v18, int32(3), l6, l7, v212, l9&int32(1))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L24
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	goto L61
L66:
	;
	return v87
L67:
	;
	F_initDeferredReplyBuffer(m, v18)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L24
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_moduleCommand(m *base.Module, l0 int32) {
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
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int64
	_ = v63
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = F_objectGetVal(m, v11)
	mBase = m.M
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v13 != int32(2) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v8 + int32(80)
	return
L2:
	;
	v143 = int32(_a755)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v146 != 0 {
		goto L44
	} else {
		goto L45
	}
L3:
	;
	v80 = int32(_a756)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v83 != 0 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v16 = int32(_a757)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v51-v53 != 0 {
		goto L2
	} else {
		goto L17
	}
L6:
	;
	v51 = F_tolower(m, v47)
	mBase = m.M
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v53 = F_tolower(m, v52)
	mBase = m.M
	goto L5
L7:
	;
	v21 = v12
	v22 = v16
	v23 = v19
	goto L10
L8:
	;
	v47 = int32(0)
	v48 = v16
	goto L6
L9:
	;
	v47 = v44 & int32(255)
	v48 = v43
	goto L6
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v25 == int32(0) {
		v43 = v22
		v44 = v23
		goto L9
	} else {
		goto L12
	}
L11:
	;
	v43 = v37
	v44 = int32(0)
	goto L9
L12:
	;
	v29 = v23 & int32(255)
	if v29 == v25 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v36 = int32(1)
	v37 = v22 + v36
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v38 != 0 {
		v21 = v21 + v36
		v22 = v37
		v23 = v38
		goto L10
	} else {
		goto L16
	}
L14:
	;
	v31 = F_tolower(m, v29)
	mBase = m.M
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v33 = F_tolower(m, v32)
	mBase = m.M
	if v31 == v33 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v43 = v22
	v44 = v35
	goto L9
L16:
	;
	goto L11
L17:
	;
	v57 = int32(0)
	v58 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(64)))) = v58
	v63 = *(*int64)(unsafe.Add(mBase, _consts[401]))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(56)))) = v63
	v68 = *(*int64)(unsafe.Add(mBase, _consts[402]))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(48)))) = v68
	v71 = *(*int64)(unsafe.Add(mBase, _consts[403]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v71
	v74 = *(*int64)(unsafe.Add(mBase, _consts[404]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v74
	F_addReplyHelp(m, l0, v8+int32(32))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	goto L1
L20:
	;
	if v13 < int32(3) {
		goto L2
	} else {
		goto L32
	}
L21:
	;
	v115 = F_tolower(m, v111)
	mBase = m.M
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v117 = F_tolower(m, v116)
	mBase = m.M
	goto L20
L22:
	;
	v85 = v12
	v86 = v80
	v87 = v83
	goto L25
L23:
	;
	v111 = int32(0)
	v112 = v80
	goto L21
L24:
	;
	v111 = v108 & int32(255)
	v112 = v107
	goto L21
L25:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v89 == int32(0) {
		v107 = v86
		v108 = v87
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v107 = v101
	v108 = int32(0)
	goto L24
L27:
	;
	v93 = v87 & int32(255)
	if v93 == v89 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v100 = int32(1)
	v101 = v86 + v100
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v102 != 0 {
		v85 = v85 + v100
		v86 = v101
		v87 = v102
		goto L25
	} else {
		goto L31
	}
L29:
	;
	v95 = F_tolower(m, v93)
	mBase = m.M
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v97 = F_tolower(m, v96)
	mBase = m.M
	if v95 == v97 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v107 = v86
	v108 = v99
	goto L24
L31:
	;
	goto L26
L32:
	;
	if v115-v117 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	v123 = F_objectGetVal(m, v122)
	mBase = m.M
	if v13 == int32(3) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_addReplyError(m, l0, int32(_a758))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L18
	} else {
		goto L41
	}
L35:
	;
	v129 = int32(0)
	goto L37
L36:
	;
	v129 = v121 + int32(12)
	goto L37
L37:
	;
	v133 = F_moduleLoad(m, v123, v129, v13+int32(-3), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L18
	} else {
		goto L38
	}
L38:
	;
	if v133 != 0 {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L18
	} else {
		goto L40
	}
L40:
	;
	goto L1
L41:
	;
	goto L1
L42:
	;
	if v13 < int32(3) {
		goto L54
	} else {
		goto L55
	}
L43:
	;
	v178 = F_tolower(m, v174)
	mBase = m.M
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v180 = F_tolower(m, v179)
	mBase = m.M
	goto L42
L44:
	;
	v148 = v12
	v149 = v143
	v150 = v146
	goto L47
L45:
	;
	v174 = int32(0)
	v175 = v143
	goto L43
L46:
	;
	v174 = v171 & int32(255)
	v175 = v170
	goto L43
L47:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v152 == int32(0) {
		v170 = v149
		v171 = v150
		goto L46
	} else {
		goto L49
	}
L48:
	;
	v170 = v164
	v171 = int32(0)
	goto L46
L49:
	;
	v156 = v150 & int32(255)
	if v156 == v152 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v163 = int32(1)
	v164 = v149 + v163
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	if v165 != 0 {
		v148 = v148 + v163
		v149 = v164
		v150 = v165
		goto L47
	} else {
		goto L53
	}
L51:
	;
	v158 = F_tolower(m, v156)
	mBase = m.M
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	v160 = F_tolower(m, v159)
	mBase = m.M
	if v158 == v160 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	v170 = v149
	v171 = v162
	goto L46
L53:
	;
	goto L48
L54:
	;
	v223 = int32(_a759)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v226 != 0 {
		goto L69
	} else {
		goto L70
	}
L55:
	;
	if v178-v180 != 0 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v184 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v184
	if base.Ui32(v13) < base.Ui32(int32(4)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v201 = F_parseLoadexArguments(m, v8+int32(32), v8+int32(28))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L18
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v13 + int32(-3)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v193 + int32(12)
	goto L57
L59:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _consts[399]))
	F_dictEmpty(m, v216, int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L18
	} else {
		goto L65
	}
L60:
	;
	if v201 != 0 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	v205 = F_objectGetVal(m, v204)
	mBase = m.M
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	v209 = F_moduleLoad(m, v205, v206, v207, int32(1))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L18
	} else {
		goto L62
	}
L62:
	;
	if v209 != 0 {
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L18
	} else {
		goto L64
	}
L64:
	;
	goto L1
L65:
	;
	F_addReplyError(m, l0, int32(_a758))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L18
	} else {
		goto L66
	}
L66:
	;
	goto L1
L67:
	;
	if v13 != int32(3) {
		goto L79
	} else {
		goto L80
	}
L68:
	;
	v258 = F_tolower(m, v254)
	mBase = m.M
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	v260 = F_tolower(m, v259)
	mBase = m.M
	goto L67
L69:
	;
	v228 = v12
	v229 = v223
	v230 = v226
	goto L72
L70:
	;
	v254 = int32(0)
	v255 = v223
	goto L68
L71:
	;
	v254 = v251 & int32(255)
	v255 = v250
	goto L68
L72:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	if v232 == int32(0) {
		v250 = v229
		v251 = v230
		goto L71
	} else {
		goto L74
	}
L73:
	;
	v250 = v244
	v251 = int32(0)
	goto L71
L74:
	;
	v236 = v230 & int32(255)
	if v236 == v232 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v243 = int32(1)
	v244 = v229 + v243
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
	if v245 != 0 {
		v228 = v228 + v243
		v229 = v244
		v230 = v245
		goto L72
	} else {
		goto L78
	}
L76:
	;
	v238 = F_tolower(m, v236)
	mBase = m.M
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	v240 = F_tolower(m, v239)
	mBase = m.M
	if v238 == v240 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	v250 = v229
	v251 = v242
	goto L71
L78:
	;
	goto L73
L79:
	;
	v305 = int32(_a760)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v308 != 0 {
		goto L98
	} else {
		goto L99
	}
L80:
	;
	if v258-v260 != 0 {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v264 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v264
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
	v268 = F_objectGetVal(m, v267)
	mBase = m.M
	v270 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	v271 = F_dictFetchValue(m, v270, v268)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L18
	} else {
		goto L84
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v285
	F_addReplyErrorFormat(m, l0, int32(_a761), v8+int32(16))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L18
	} else {
		goto L93
	}
L83:
	;
	v276 = F_moduleUnloadInternal(m, v271, v8+int32(32))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L18
	} else {
		goto L87
	}
L84:
	;
	if v271 != 0 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v285 = int32(_a762)
	goto L82
L86:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	if v282 != 0 {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	if v276 != 0 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L18
	} else {
		goto L89
	}
L89:
	;
	goto L1
L90:
	;
	v284 = v282
	goto L92
L91:
	;
	v284 = int32(_a763)
	goto L92
L92:
	;
	v285 = v284
	goto L82
L93:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v293 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+8))
	v298 = F_objectGetVal(m, v297)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v298
	F__serverLog(m, int32(3), int32(_a764), v8)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L18
	} else {
		goto L95
	}
L95:
	;
	goto L1
L96:
	;
	if v13 != int32(2) {
		goto L108
	} else {
		goto L109
	}
L97:
	;
	v340 = F_tolower(m, v336)
	mBase = m.M
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	v342 = F_tolower(m, v341)
	mBase = m.M
	goto L96
L98:
	;
	v310 = v12
	v311 = v305
	v312 = v308
	goto L101
L99:
	;
	v336 = int32(0)
	v337 = v305
	goto L97
L100:
	;
	v336 = v333 & int32(255)
	v337 = v332
	goto L97
L101:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	if v314 == int32(0) {
		v332 = v311
		v333 = v312
		goto L100
	} else {
		goto L103
	}
L102:
	;
	v332 = v326
	v333 = int32(0)
	goto L100
L103:
	;
	v318 = v312 & int32(255)
	if v318 == v314 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v325 = int32(1)
	v326 = v311 + v325
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
	if v327 != 0 {
		v310 = v310 + v325
		v311 = v326
		v312 = v327
		goto L101
	} else {
		goto L107
	}
L105:
	;
	v320 = F_tolower(m, v318)
	mBase = m.M
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	v322 = F_tolower(m, v321)
	mBase = m.M
	if v320 == v322 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	v332 = v311
	v333 = v324
	goto L100
L107:
	;
	goto L102
L108:
	;
	F_addReplySubcommandSyntaxError(m, l0)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L18
	} else {
		goto L112
	}
L109:
	;
	if v340-v342 != 0 {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	F_addReplyLoadedModules(m, l0)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L18
	} else {
		goto L111
	}
L111:
	;
	goto L1
L112:
	;
	goto L1
}
func F_moduleConfigApplyConfig(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int64
	_ = v111
	var v115 int64
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v131 int64
	_ = v131
	var v133 int32
	_ = v133
	var v137 int64
	_ = v137
	var v141 int64
	_ = v141
	var v144 int32
	_ = v144
	var v152 int64
	_ = v152
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v261 int32
	_ = v261
	v14 = m.G0
	v16 = v14 - int32(96)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v18 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(96)
	return v261
L2:
	;
	v261 = int32(1)
	goto L1
L3:
	;
	v21 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v21
	v24 = v16 + int32(88)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v25
	goto L4
L4:
	;
	v30 = v16 + int32(88)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v32 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v32 == int32(0) {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v32+base.B2i32(v35 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v41
	goto L6
L8:
	;
	v61 = v32
	goto L9
L9:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	v76 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(72)))) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(64)))) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(56)))) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(48)))) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(40)))) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(32)))) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(24)))) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(16)))) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(561)
	v98 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	v99 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	v101 = m.T0[v100].(func(*base.Module) int64)(m)
	mBase = m.M
	if v98 == v99 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L2
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v115
	v119 = int32(0)
	v123 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	*(*int32)(unsafe.Add(mBase, _consts[275])) = v123 + int32(1)
	goto L16
L12:
	;
	v111 = *(*int64)(unsafe.Add(mBase, _consts[348]))
	v115 = v111*int64(1000) + v101
	goto L11
L13:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[149]))
	v107 = base.I32_div_s(int32(1000000), v106)
	v115 = v101 + base.I64_extend_i32_s(v107)
	goto L11
L14:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	v162 = m.T0[v161].(func(*base.Module, int32, int32, int32) int32)(m, v16+int32(8), v158, v16+int32(84))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	goto L14
L16:
	;
	if v123 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	goto L19
L18:
	;
	v133 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v131
	v137 = base.I64_div_s(v131, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[32])) = v137
	v141 = base.I64_div_s(v131, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[37])) = v141
	v144 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	F_lrulfu_updateClockAndPolicy(m, v137, int32(base.Ui32(v144&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v152 = *(*int64)(unsafe.Add(mBase, _consts[32]))
	*(*int64)(unsafe.Add(mBase, _consts[78])) = v152
	goto L15
L19:
	;
	v131 = F_ustime(m)
	mBase = m.M
	goto L18
L20:
	;
	F_moduleFreeContext(m, v16+int32(8))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L21
	} else {
		goto L42
	}
L21:
	;
	return int32(0)
L22:
	;
	if v162 == int32(0) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	if l2 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v16)+84))
	if v172 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v170
	goto L24
L26:
	;
	F_moduleFreeContext(m, v16+int32(8))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L21
	} else {
		goto L41
	}
L27:
	;
	v176 = F_objectGetVal(m, v172)
	mBase = m.M
	goto L31
L28:
	;
	F_decrRefCount(m, v172)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L21
	} else {
		goto L40
	}
L29:
	;
	goto L28
L30:
	;
	v207 = v186
	goto L37
L31:
	;
	v182 = int32(_a752)
	v184 = int32(256)
	v186 = v176
	goto L33
L32:
	;
	v197 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v182))) = uint8(v197)
	goto L30
L33:
	;
	v188 = v184 + int32(-1)
	if v188 == int32(0) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	*(*uint8)(unsafe.Add(mBase, uint32(v182))) = uint8(v191)
	v193 = int32(1)
	if v191 != 0 {
		v182 = v182 + v193
		v184 = v188
		v186 = v186 + v193
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L29
L37:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v209 != 0 {
		v207 = v207 + int32(1)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	goto L29
L39:
	;
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a752)
	goto L26
L41:
	;
	v261 = int32(0)
	goto L1
L42:
	;
	v234 = v16 + int32(88)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if v236 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v236 != 0 {
		v61 = v236
		goto L9
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v236+base.B2i32(v239 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v245
	goto L44
L46:
	;
	goto L10
}
func F_moduleCreateContext(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v15 int32
	_ = v15
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int64
	_ = v108
	var v112 int64
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v136 int64
	_ = v136
	var v140 int64
	_ = v140
	var v143 int32
	_ = v143
	var v151 int64
	_ = v151
	v5 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(24)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(561)
	v15 = int32(64)
	*(*int64)(unsafe.Add(mBase, uint32(l0+v15))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(56)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(48)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(40)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(32)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(16)))) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l2
	if l2&v15 == int32(0) {
		if l2&int32(128) == int32(0) {
			v95 = *(*int32)(unsafe.Add(mBase, _consts[116]))
			v96 = int32(0)
			v97 = *(*int32)(unsafe.Add(mBase, _consts[271]))
			v98 = m.T0[v97].(func(*base.Module) int64)(m)
			mBase = m.M
			if v95 == v96 {
				v108 = *(*int64)(unsafe.Add(mBase, _consts[348]))
				v112 = v108*int64(1000) + v98
			} else {
				v103 = *(*int32)(unsafe.Add(mBase, _consts[149]))
				v104 = base.I32_div_s(int32(1000000), v103)
				v112 = v98 + base.I64_extend_i32_s(v104)
			}
			*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v112
			if l2&int32(528) != 0 {
			} else {
				v118 = int32(0)
				v122 = *(*int32)(unsafe.Add(mBase, _consts[275]))
				*(*int32)(unsafe.Add(mBase, _consts[275])) = v122 + int32(1)
				if v122 != 0 {
				} else {
					v130 = F_ustime(m)
					mBase = m.M
					v132 = int32(0)
					*(*int64)(unsafe.Add(mBase, _consts[277])) = v130
					v136 = base.I64_div_s(v130, int64(1000))
					*(*int64)(unsafe.Add(mBase, _consts[32])) = v136
					v140 = base.I64_div_s(v130, int64(1000000))
					*(*int64)(unsafe.Add(mBase, _consts[37])) = v140
					v143 = *(*int32)(unsafe.Add(mBase, _consts[167]))
					F_lrulfu_updateClockAndPolicy(m, v136, int32(base.Ui32(v143&int32(2))>>(uint(int32(1))%32)))
					mBase = m.M
					v151 = *(*int64)(unsafe.Add(mBase, _consts[32]))
					*(*int64)(unsafe.Add(mBase, _consts[78])) = v151
				}
			}
			return
		} else {
			v85 = F_createClient(m, int32(0))
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v85
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+204))
				*(*int32)(unsafe.Add(mBase, uint32(v85)+204)) = v88 | int32(268435456)
				v95 = *(*int32)(unsafe.Add(mBase, _consts[116]))
				v96 = int32(0)
				v97 = *(*int32)(unsafe.Add(mBase, _consts[271]))
				v98 = m.T0[v97].(func(*base.Module) int64)(m)
				mBase = m.M
				if v95 == v96 {
					v108 = *(*int64)(unsafe.Add(mBase, _consts[348]))
					v112 = v108*int64(1000) + v98
				} else {
					v103 = *(*int32)(unsafe.Add(mBase, _consts[149]))
					v104 = base.I32_div_s(int32(1000000), v103)
					v112 = v98 + base.I64_extend_i32_s(v104)
				}
				*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v112
				if l2&int32(528) != 0 {
				} else {
					v118 = int32(0)
					v122 = *(*int32)(unsafe.Add(mBase, _consts[275]))
					*(*int32)(unsafe.Add(mBase, _consts[275])) = v122 + int32(1)
					if v122 != 0 {
					} else {
						v130 = F_ustime(m)
						mBase = m.M
						v132 = int32(0)
						*(*int64)(unsafe.Add(mBase, _consts[277])) = v130
						v136 = base.I64_div_s(v130, int64(1000))
						*(*int64)(unsafe.Add(mBase, _consts[32])) = v136
						v140 = base.I64_div_s(v130, int64(1000000))
						*(*int64)(unsafe.Add(mBase, _consts[37])) = v140
						v143 = *(*int32)(unsafe.Add(mBase, _consts[167]))
						F_lrulfu_updateClockAndPolicy(m, v136, int32(base.Ui32(v143&int32(2))>>(uint(int32(1))%32)))
						mBase = m.M
						v151 = *(*int64)(unsafe.Add(mBase, _consts[32]))
						*(*int64)(unsafe.Add(mBase, _consts[78])) = v151
					}
				}
				return
			}
		}
	} else {
		v44 = int32(0)
		v45 = *(*int32)(unsafe.Add(mBase, _consts[345]))
		if v45 == v44 {
			v65 = F_createClient(m, int32(0))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v65)+328)) = int32(0)
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+200))
				*(*int32)(unsafe.Add(mBase, uint32(v65)+200)) = v69 | int32(1073741824)
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v65)+204))
				*(*int32)(unsafe.Add(mBase, uint32(v65)+204)) = v73 | int32(268435456)
				v77 = v65
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v77
				v95 = *(*int32)(unsafe.Add(mBase, _consts[116]))
				v96 = int32(0)
				v97 = *(*int32)(unsafe.Add(mBase, _consts[271]))
				v98 = m.T0[v97].(func(*base.Module) int64)(m)
				mBase = m.M
				if v95 == v96 {
					v108 = *(*int64)(unsafe.Add(mBase, _consts[348]))
					v112 = v108*int64(1000) + v98
				} else {
					v103 = *(*int32)(unsafe.Add(mBase, _consts[149]))
					v104 = base.I32_div_s(int32(1000000), v103)
					v112 = v98 + base.I64_extend_i32_s(v104)
				}
				*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v112
				if l2&int32(528) != 0 {
				} else {
					v118 = int32(0)
					v122 = *(*int32)(unsafe.Add(mBase, _consts[275]))
					*(*int32)(unsafe.Add(mBase, _consts[275])) = v122 + int32(1)
					if v122 != 0 {
					} else {
						v130 = F_ustime(m)
						mBase = m.M
						v132 = int32(0)
						*(*int64)(unsafe.Add(mBase, _consts[277])) = v130
						v136 = base.I64_div_s(v130, int64(1000))
						*(*int64)(unsafe.Add(mBase, _consts[32])) = v136
						v140 = base.I64_div_s(v130, int64(1000000))
						*(*int64)(unsafe.Add(mBase, _consts[37])) = v140
						v143 = *(*int32)(unsafe.Add(mBase, _consts[167]))
						F_lrulfu_updateClockAndPolicy(m, v136, int32(base.Ui32(v143&int32(2))>>(uint(int32(1))%32)))
						mBase = m.M
						v151 = *(*int64)(unsafe.Add(mBase, _consts[32]))
						*(*int64)(unsafe.Add(mBase, _consts[78])) = v151
					}
				}
				return
			}
		} else {
			v48 = int32(0)
			v50 = v45 + int32(-1)
			*(*int32)(unsafe.Add(mBase, _consts[345])) = v50
			v53 = *(*int32)(unsafe.Add(mBase, _consts[347]))
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v50<<(uint(int32(2))%32))))
			v59 = *(*int32)(unsafe.Add(mBase, _consts[349]))
			if base.Ui32(v59) <= base.Ui32(v50) {
				v77 = v57
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v77
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[349])) = v50
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v57
			}
			v95 = *(*int32)(unsafe.Add(mBase, _consts[116]))
			v96 = int32(0)
			v97 = *(*int32)(unsafe.Add(mBase, _consts[271]))
			v98 = m.T0[v97].(func(*base.Module) int64)(m)
			mBase = m.M
			if v95 == v96 {
				v108 = *(*int64)(unsafe.Add(mBase, _consts[348]))
				v112 = v108*int64(1000) + v98
			} else {
				v103 = *(*int32)(unsafe.Add(mBase, _consts[149]))
				v104 = base.I32_div_s(int32(1000000), v103)
				v112 = v98 + base.I64_extend_i32_s(v104)
			}
			*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v112
			if l2&int32(528) != 0 {
			} else {
				v118 = int32(0)
				v122 = *(*int32)(unsafe.Add(mBase, _consts[275]))
				*(*int32)(unsafe.Add(mBase, _consts[275])) = v122 + int32(1)
				if v122 != 0 {
				} else {
					v130 = F_ustime(m)
					mBase = m.M
					v132 = int32(0)
					*(*int64)(unsafe.Add(mBase, _consts[277])) = v130
					v136 = base.I64_div_s(v130, int64(1000))
					*(*int64)(unsafe.Add(mBase, _consts[32])) = v136
					v140 = base.I64_div_s(v130, int64(1000000))
					*(*int64)(unsafe.Add(mBase, _consts[37])) = v140
					v143 = *(*int32)(unsafe.Add(mBase, _consts[167]))
					F_lrulfu_updateClockAndPolicy(m, v136, int32(base.Ui32(v143&int32(2))>>(uint(int32(1))%32)))
					mBase = m.M
					v151 = *(*int64)(unsafe.Add(mBase, _consts[32]))
					*(*int64)(unsafe.Add(mBase, _consts[78])) = v151
				}
			}
			return
		}
	}
}
func F_moduleCreateEmptyKey(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(1)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v10&int32(2) == int32(0) {
		v53 = v9
		m.G0 = v7 + int32(16)
		return v53
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v15 != 0 {
			v53 = v9
			m.G0 = v7 + int32(16)
			return v53
		} else {
			switch l1 + int32(-2) {
			case 0:
				v18 = F_createListListpackObject(m)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v28 = v18
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v28
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_dbAdd(m, v30, v31, v7+int32(12))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v36
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
						switch v39&int32(15) + int32(-3) {
						case 0:
							*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(4294967296)
							v46 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v46
							v53 = v46
						default:
							v53 = int32(0)
						case 3:
							v49 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v49
							v53 = v49
						}
						m.G0 = v7 + int32(16)
						return v53
					}
				}
			case 1:
				v24 = F_createHashObject(m)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v28 = v24
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v28
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_dbAdd(m, v30, v31, v7+int32(12))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v36
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
						switch v39&int32(15) + int32(-3) {
						case 0:
							*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(4294967296)
							v46 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v46
							v53 = v46
						default:
							v53 = int32(0)
						case 3:
							v49 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v49
							v53 = v49
						}
						m.G0 = v7 + int32(16)
						return v53
					}
				}
			default:
				v53 = v9
				m.G0 = v7 + int32(16)
				return v53
			case 3:
				v22 = F_createZsetListpackObject(m)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v28 = v22
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v28
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_dbAdd(m, v30, v31, v7+int32(12))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v36
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
						switch v39&int32(15) + int32(-3) {
						case 0:
							*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(4294967296)
							v46 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v46
							v53 = v46
						default:
							v53 = int32(0)
						case 3:
							v49 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v49
							v53 = v49
						}
						m.G0 = v7 + int32(16)
						return v53
					}
				}
			case 5:
				v26 = F_createStreamObject(m)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = v26
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v28
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_dbAdd(m, v30, v31, v7+int32(12))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v36
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
						switch v39&int32(15) + int32(-3) {
						case 0:
							*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(4294967296)
							v46 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v46
							v53 = v46
						default:
							v53 = int32(0)
						case 3:
							v49 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v49
							v53 = v49
						}
						m.G0 = v7 + int32(16)
						return v53
					}
				}
			}
		}
	}
}
func F_moduleFireCommandACLRejectedEvent(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v104 int64
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v13 = *(*int32)(unsafe.Add(mBase, _consts[390]))
	if v13 == v4 {
		m.G0 = v9 + int32(96)
		return
	} else {
		if l1&int64(-2) != int64(2) {
			v84 = v4
		} else {
			if l2 < int32(0) {
				v84 = v4
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v23 <= l2 {
					v84 = int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+l2<<(uint(int32(2))%32))))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					v31 = F_objectGetVal(m, v29)
					mBase = m.M
					if v30&int32(240) != int32(16) {
						v84 = v31
					} else {
						v37 = v9 + int32(64)
						v39 = base.I64_extend_i32_s(v31)
						if v39 <= int64(-1) {
							v48 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v48)
							v57 = v9 + int32(65)
							v58 = int32(20)
							v59 = int64(0) - v39
						} else {
							v57 = v37
							v58 = int32(21)
							v59 = v39
						}
						v61 = F_ull2string(m, v57, v58, v59)
						mBase = m.M
						if v61 == int32(0) {
						} else {
						}
						v84 = v9 + int32(64)
					}
				}
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(1)
		v87 = int32(0)
		v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		if v88 == v87 {
			v92 = v87
		} else {
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v88)+140))
			v92 = v91
		}
		v95 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9+int32(28)))) = v95
		*(*int32)(unsafe.Add(mBase, uint32(v9+int32(36)))) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+20)) = v95
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v92
		v104 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v104
		v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = int32(base.Ui32(v106)>>(uint(int32(30))%32)) & int32(1)
		v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v112
		v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = v84
		*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v114
		F_moduleFireServerEvent(m, int64(23), base.I32_wrap_i64(l1), v9+int32(8))
		mBase = m.M
		v122 = m.ExcPending
		if v122 != 0 {
			return
		} else {
			m.G0 = v9 + int32(96)
			return
		}
	}
}
func F_moduleFireCommandResultEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v130 int64
	_ = v130
	var v135 int32
	_ = v135
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v174 int32
	_ = v174
	v13 = m.G0
	v15 = v13 - int32(64)
	m.G0 = v15
	if l2 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(64)
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v26 != 0 {
		v29 = int32(188)
		v30 = v26
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v21 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[388]))
	if v22 == v21 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[389]))
	if v20 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L1
L6:
	;
	goto L2
L7:
	;
	v31 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0+v29)))
	if v31 < v33 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = int32(24)
	v30 = v27
	goto L7
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = int64(1)
	if l1 == int32(0) {
		v109 = v31
		goto L23
	} else {
		goto L24
	}
L10:
	;
	v48 = int32(0)
	goto L13
L11:
	;
	v99 = v30
	v101 = int32(0)
	goto L9
L12:
	;
	v66 = F_valkey_malloc(m, v33<<(uint(int32(2))%32))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v30+v48<<(uint(int32(2))%32))))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v54&int32(240) == int32(16) {
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v99 = v30
	v101 = int32(0)
	goto L9
L15:
	;
	v60 = v48 + int32(1)
	if v60 != v33 {
		v48 = v60
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	return
L18:
	;
	v78 = int32(0)
	goto L19
L19:
	;
	v81 = v78 << (uint(int32(2)) % 32)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v30+v81)))
	v85 = F_getDecodedObject(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v99 = v66
	v101 = v88
	goto L9
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66+v81))) = v85
	v88 = int32(1)
	v90 = v78 + v88
	if v90 != v33 {
		v78 = v90
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = l3
	v112 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v109
	v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = int32(base.Ui32(v118)>>(uint(int32(30))%32)) & int32(1)
	if l2 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v109 = v108
	goto L23
L25:
	;
	v130 = int64(21)
	goto L27
L26:
	;
	v130 = int64(20)
	goto L27
L27:
	;
	F_moduleFireServerEvent(m, v130, int32(0), v15+int32(8))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	if v101 == int32(0) {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v33 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_valkey_free(m, v99)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L17
	} else {
		goto L36
	}
L31:
	;
	v150 = v112
	goto L32
L32:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v99+v150<<(uint(int32(2))%32))))
	F_decrRefCount(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L17
	} else {
		goto L34
	}
L33:
	;
	goto L30
L34:
	;
	v159 = v150 + int32(1)
	if v159 != v33 {
		v150 = v159
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	goto L1
}
func F_moduleFreeKeyIterator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v3 == int32(0) {
		F__serverAssert(m, int32(_a693), int32(_a694), int32(795))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		switch v7&int32(15) + int32(-1) {
		case 0:
			F_listTypeReleaseIterator(m, v3)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
				return
			}
		default:
			F__serverAssert(m, int32(_a107), int32(_a694), int32(802))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 5:
			F_streamIteratorStop(m, v3)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				F_valkey_free(m, v14)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
					return
				}
			}
		}
	}
}
func F_moduleGetCommandChannelsViaAPI(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int64
	_ = v63
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v89 int64
	_ = v89
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v104 int64
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	v5 = int32(0)
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(32)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(72)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(64)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(56)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(48)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(561)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(256)
	v50 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	v52 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	v53 = m.T0[v52].(func(*base.Module) int64)(m)
	mBase = m.M
	if v50 == v5 {
		v63 = *(*int64)(unsafe.Add(mBase, _consts[348]))
		v67 = v63*int64(1000) + v53
	} else {
		v58 = *(*int32)(unsafe.Add(mBase, _consts[149]))
		v59 = base.I32_div_s(int32(1000000), v58)
		v67 = v53 + base.I64_extend_i32_s(v59)
	}
	*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v67
	v71 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	*(*int32)(unsafe.Add(mBase, _consts[275])) = v75 + int32(1)
	if v75 != 0 {
	} else {
		v83 = F_ustime(m)
		mBase = m.M
		v85 = int32(0)
		*(*int64)(unsafe.Add(mBase, _consts[277])) = v83
		v89 = base.I64_div_s(v83, int64(1000))
		*(*int64)(unsafe.Add(mBase, _consts[32])) = v89
		v93 = base.I64_div_s(v83, int64(1000000))
		*(*int64)(unsafe.Add(mBase, _consts[37])) = v93
		v96 = *(*int32)(unsafe.Add(mBase, _consts[167]))
		F_lrulfu_updateClockAndPolicy(m, v89, int32(base.Ui32(v96&int32(2))>>(uint(int32(1))%32)))
		mBase = m.M
		v104 = *(*int64)(unsafe.Add(mBase, _consts[32]))
		*(*int64)(unsafe.Add(mBase, _consts[78])) = v104
	}
	v109 = F_getKeysPrepareResult(m, l3, int32(256))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = l3
		v116 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v117 = m.T0[v116].(func(*base.Module, int32, int32, int32) int32)(m, v10+int32(8), l1, l2)
		mBase = m.M
		v118 = m.ExcPending
		if v118 != 0 {
			return int32(0)
		} else {
			F_moduleFreeContext(m, v10+int32(8))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return int32(0)
			} else {
				v123 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				m.G0 = v10 + int32(80)
				return v123
			}
		}
	}
}
func F_moduleGetMemUsage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_objectGetVal(m, l1)
	mBase = m.M
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	if v14 == int32(0) {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
		if v28 != 0 {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v31 = m.T0[v28].(func(*base.Module, int32) int32)(m, v30)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = v31
				m.G0 = v10 + int32(16)
				return v33
			}
		} else {
			v33 = int32(0)
			m.G0 = v10 + int32(16)
			return v33
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v24 = m.T0[v14].(func(*base.Module, int32, int32, int32) int32)(m, v10, v23, l2)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v33 = v24
			m.G0 = v10 + int32(16)
			return v33
		}
	}
}
func F_moduleInitModulesSystemLast(m *base.Module) {
	return
}
func F_moduleLoad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
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
	v8 = m.G0
	v10 = v8 - int32(160)
	m.G0 = v10
	v16 = F___fstatat(m, int32(-100), l0, v10+int32(64), int32(0))
	mBase = m.M
	if v16 != 0 {
		if l0 != 0 {
			v34 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(_a739)
			v38 = v34
		} else {
			v38 = int32(_a0)
		}
		if v38 == int32(0) {
			v59 = int32(-1)
			v61 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			if int32(3) < v61 {
				v122 = v59
				m.G0 = v10 + int32(160)
				return v122
			} else {
				v64 = int32(0)
				v66 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v64
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v66
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
				F__serverLog(m, int32(3), int32(_a740), v10+int32(16))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					v122 = v59
					m.G0 = v10 + int32(160)
					return v122
				}
			}
		} else {
			v41 = int32(_a741)
			if v38 != int32(_a0) {
				v52 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(_a1)
				v56 = v52
			} else {
				v45 = F_strcmp(m, int32(_a2), v41)
				mBase = m.M
				if v45 != 0 {
					v49 = F_strcmp(m, int32(_a3), v41)
					mBase = m.M
					if v49 != 0 {
						v52 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(_a1)
						v56 = v52
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, _consts[4]))
						v56 = v51
					}
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, _consts[3]))
					v56 = v47
				}
			}
			if v56 == int32(0) {
				v78 = int32(_a742)
				if v38 != int32(_a0) {
					v89 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(_a1)
					v93 = v89
				} else {
					v82 = F_strcmp(m, int32(_a2), v78)
					mBase = m.M
					if v82 != 0 {
						v86 = F_strcmp(m, int32(_a3), v78)
						mBase = m.M
						if v86 != 0 {
							v89 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(_a1)
							v93 = v89
						} else {
							v88 = *(*int32)(unsafe.Add(mBase, _consts[4]))
							v93 = v88
						}
					} else {
						v84 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						v93 = v84
					}
				}
				if v93 != 0 {
					v108 = *(*int32)(unsafe.Add(mBase, _consts[15]))
					if int32(2) < v108 {
						v118 = v93
						v120 = F_moduleInitPostOnLoadResolved(m, v118, v38, l0, l1, l2, l3, int32(0))
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							v122 = v120
							m.G0 = v10 + int32(160)
							return v122
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l0
						F__serverLog(m, int32(2), int32(_a743), v10+int32(48))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return int32(0)
						} else {
							v118 = v93
							v120 = F_moduleInitPostOnLoadResolved(m, v118, v38, l0, l1, l2, l3, int32(0))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								v122 = v120
								m.G0 = v10 + int32(160)
								return v122
							}
						}
					}
				} else {
					v95 = int32(-1)
					v97 = *(*int32)(unsafe.Add(mBase, _consts[15]))
					if int32(3) < v97 {
						v122 = v95
						m.G0 = v10 + int32(160)
						return v122
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
						F__serverLog(m, int32(3), int32(_a744), v10+int32(32))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							v122 = v95
							m.G0 = v10 + int32(160)
							return v122
						}
					}
				}
			} else {
				v118 = v56
				v120 = F_moduleInitPostOnLoadResolved(m, v118, v38, l0, l1, l2, l3, int32(0))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return int32(0)
				} else {
					v122 = v120
					m.G0 = v10 + int32(160)
					return v122
				}
			}
		}
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+68)))
		if v17&int32(73) != 0 {
			if l0 != 0 {
				v34 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(_a739)
				v38 = v34
			} else {
				v38 = int32(_a0)
			}
			if v38 == int32(0) {
				v59 = int32(-1)
				v61 = *(*int32)(unsafe.Add(mBase, _consts[15]))
				if int32(3) < v61 {
					v122 = v59
					m.G0 = v10 + int32(160)
					return v122
				} else {
					v64 = int32(0)
					v66 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v64
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v66
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
					F__serverLog(m, int32(3), int32(_a740), v10+int32(16))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						v122 = v59
						m.G0 = v10 + int32(160)
						return v122
					}
				}
			} else {
				v41 = int32(_a741)
				if v38 != int32(_a0) {
					v52 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(_a1)
					v56 = v52
				} else {
					v45 = F_strcmp(m, int32(_a2), v41)
					mBase = m.M
					if v45 != 0 {
						v49 = F_strcmp(m, int32(_a3), v41)
						mBase = m.M
						if v49 != 0 {
							v52 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(_a1)
							v56 = v52
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, _consts[4]))
							v56 = v51
						}
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						v56 = v47
					}
				}
				if v56 == int32(0) {
					v78 = int32(_a742)
					if v38 != int32(_a0) {
						v89 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(_a1)
						v93 = v89
					} else {
						v82 = F_strcmp(m, int32(_a2), v78)
						mBase = m.M
						if v82 != 0 {
							v86 = F_strcmp(m, int32(_a3), v78)
							mBase = m.M
							if v86 != 0 {
								v89 = int32(0)
								*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(_a1)
								v93 = v89
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, _consts[4]))
								v93 = v88
							}
						} else {
							v84 = *(*int32)(unsafe.Add(mBase, _consts[3]))
							v93 = v84
						}
					}
					if v93 != 0 {
						v108 = *(*int32)(unsafe.Add(mBase, _consts[15]))
						if int32(2) < v108 {
							v118 = v93
							v120 = F_moduleInitPostOnLoadResolved(m, v118, v38, l0, l1, l2, l3, int32(0))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								v122 = v120
								m.G0 = v10 + int32(160)
								return v122
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l0
							F__serverLog(m, int32(2), int32(_a743), v10+int32(48))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
							} else {
								v118 = v93
								v120 = F_moduleInitPostOnLoadResolved(m, v118, v38, l0, l1, l2, l3, int32(0))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									v122 = v120
									m.G0 = v10 + int32(160)
									return v122
								}
							}
						}
					} else {
						v95 = int32(-1)
						v97 = *(*int32)(unsafe.Add(mBase, _consts[15]))
						if int32(3) < v97 {
							v122 = v95
							m.G0 = v10 + int32(160)
							return v122
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
							F__serverLog(m, int32(3), int32(_a744), v10+int32(32))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								v122 = v95
								m.G0 = v10 + int32(160)
								return v122
							}
						}
					}
				} else {
					v118 = v56
					v120 = F_moduleInitPostOnLoadResolved(m, v118, v38, l0, l1, l2, l3, int32(0))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						v122 = v120
						m.G0 = v10 + int32(160)
						return v122
					}
				}
			}
		} else {
			v20 = int32(-1)
			v22 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			if int32(3) < v22 {
				v122 = v20
				m.G0 = v10 + int32(160)
				return v122
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F__serverLog(m, int32(3), int32(_a745), v10)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v122 = v20
					m.G0 = v10 + int32(160)
					return v122
				}
			}
		}
	}
}
func F_moduleLogRaw(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	v6 = m.G0
	v8 = v6 - int32(1040)
	m.G0 = v8
	v10 = int32(_a713)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v13 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v172 < v174 {
		goto L59
	} else {
		goto L60
	}
L2:
	;
	v50 = int32(_a714)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v53 != 0 {
		goto L19
	} else {
		goto L20
	}
L3:
	;
	if v45-v47 != 0 {
		goto L2
	} else {
		goto L15
	}
L4:
	;
	v45 = F_tolower(m, v41)
	mBase = m.M
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v47 = F_tolower(m, v46)
	mBase = m.M
	goto L3
L5:
	;
	v15 = l1
	v16 = v10
	v17 = v13
	goto L8
L6:
	;
	v41 = int32(0)
	v42 = v10
	goto L4
L7:
	;
	v41 = v38 & int32(255)
	v42 = v37
	goto L4
L8:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v19 == int32(0) {
		v37 = v16
		v38 = v17
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v37 = v31
	v38 = int32(0)
	goto L7
L10:
	;
	v23 = v17 & int32(255)
	if v23 == v19 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v30 = int32(1)
	v31 = v16 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v32 != 0 {
		v15 = v15 + v30
		v16 = v31
		v17 = v32
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v25 = F_tolower(m, v23)
	mBase = m.M
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v27 = F_tolower(m, v26)
	mBase = m.M
	if v25 == v27 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v37 = v16
	v38 = v29
	goto L7
L14:
	;
	goto L9
L15:
	;
	v172 = int32(0)
	goto L1
L16:
	;
	v90 = int32(_a715)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v93 != 0 {
		goto L33
	} else {
		goto L34
	}
L17:
	;
	if v85-v87 != 0 {
		goto L16
	} else {
		goto L29
	}
L18:
	;
	v85 = F_tolower(m, v81)
	mBase = m.M
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v87 = F_tolower(m, v86)
	mBase = m.M
	goto L17
L19:
	;
	v55 = l1
	v56 = v50
	v57 = v53
	goto L22
L20:
	;
	v81 = int32(0)
	v82 = v50
	goto L18
L21:
	;
	v81 = v78 & int32(255)
	v82 = v77
	goto L18
L22:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v59 == int32(0) {
		v77 = v56
		v78 = v57
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v77 = v71
	v78 = int32(0)
	goto L21
L24:
	;
	v63 = v57 & int32(255)
	if v63 == v59 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v70 = int32(1)
	v71 = v56 + v70
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v72 != 0 {
		v55 = v55 + v70
		v56 = v71
		v57 = v72
		goto L22
	} else {
		goto L28
	}
L26:
	;
	v65 = F_tolower(m, v63)
	mBase = m.M
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	v67 = F_tolower(m, v66)
	mBase = m.M
	if v65 == v67 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v77 = v56
	v78 = v69
	goto L21
L28:
	;
	goto L23
L29:
	;
	v172 = int32(1)
	goto L1
L30:
	;
	v132 = int32(_a716)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v135 != 0 {
		goto L46
	} else {
		goto L47
	}
L31:
	;
	if v125-v127 != 0 {
		goto L30
	} else {
		goto L43
	}
L32:
	;
	v125 = F_tolower(m, v121)
	mBase = m.M
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v127 = F_tolower(m, v126)
	mBase = m.M
	goto L31
L33:
	;
	v95 = l1
	v96 = v90
	v97 = v93
	goto L36
L34:
	;
	v121 = int32(0)
	v122 = v90
	goto L32
L35:
	;
	v121 = v118 & int32(255)
	v122 = v117
	goto L32
L36:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v99 == int32(0) {
		v117 = v96
		v118 = v97
		goto L35
	} else {
		goto L38
	}
L37:
	;
	v117 = v111
	v118 = int32(0)
	goto L35
L38:
	;
	v103 = v97 & int32(255)
	if v103 == v99 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v110 = int32(1)
	v111 = v96 + v110
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	if v112 != 0 {
		v95 = v95 + v110
		v96 = v111
		v97 = v112
		goto L36
	} else {
		goto L42
	}
L40:
	;
	v105 = F_tolower(m, v103)
	mBase = m.M
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	v107 = F_tolower(m, v106)
	mBase = m.M
	if v105 == v107 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v117 = v96
	v118 = v109
	goto L35
L42:
	;
	goto L37
L43:
	;
	v172 = int32(2)
	goto L1
L44:
	;
	if v167-v169 != 0 {
		goto L56
	} else {
		goto L57
	}
L45:
	;
	v167 = F_tolower(m, v163)
	mBase = m.M
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	v169 = F_tolower(m, v168)
	mBase = m.M
	goto L44
L46:
	;
	v137 = l1
	v138 = v132
	v139 = v135
	goto L49
L47:
	;
	v163 = int32(0)
	v164 = v132
	goto L45
L48:
	;
	v163 = v160 & int32(255)
	v164 = v159
	goto L45
L49:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v141 == int32(0) {
		v159 = v138
		v160 = v139
		goto L48
	} else {
		goto L51
	}
L50:
	;
	v159 = v153
	v160 = int32(0)
	goto L48
L51:
	;
	v145 = v139 & int32(255)
	if v145 == v141 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v152 = int32(1)
	v153 = v138 + v152
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	if v154 != 0 {
		v137 = v137 + v152
		v138 = v153
		v139 = v154
		goto L49
	} else {
		goto L55
	}
L53:
	;
	v147 = F_tolower(m, v145)
	mBase = m.M
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v149 = F_tolower(m, v148)
	mBase = m.M
	if v147 == v149 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	v159 = v138
	v160 = v151
	goto L48
L55:
	;
	goto L50
L56:
	;
	v171 = int32(1)
	goto L58
L57:
	;
	v171 = int32(3)
	goto L58
L58:
	;
	v172 = v171
	goto L1
L59:
	;
	m.G0 = v8 + int32(1040)
	return
L60:
	;
	if l0 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v178
	v181 = v8 + int32(16)
	v186 = F_snprintf(m, v181, int32(1024), int32(_a717), v8)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v178 = v177
	goto L61
L63:
	;
	v178 = int32(_a718)
	goto L61
L64:
	;
	return
L65:
	;
	v191 = F_vsnprintf(m, v181+v186, int32(1024)-v186, l2, l3)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	F_serverLogRaw(m, v172, v8+int32(16))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L59
}
func F_moduleNotifyKeyspaceEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int64
	_ = v220
	var v224 int64
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v240 int64
	_ = v240
	var v242 int32
	_ = v242
	var v246 int64
	_ = v246
	var v250 int64
	_ = v250
	var v253 int32
	_ = v253
	var v261 int64
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int64
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int64
	_ = v280
	var v284 int64
	_ = v284
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v300 int64
	_ = v300
	var v302 int32
	_ = v302
	var v306 int64
	_ = v306
	var v310 int64
	_ = v310
	var v313 int32
	_ = v313
	var v321 int64
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	v5 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(80)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _consts[380]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v25 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v21 + int32(80)
	return
L2:
	;
	v28 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	*(*int32)(unsafe.Add(mBase, _consts[275])) = v34 + int32(1)
	goto L4
L3:
	;
	v67 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, _consts[380]))
	v70 = v21 + int32(72)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v71
	goto L9
L4:
	;
	goto L3
L9:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v76 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v82 = v21 + int32(72)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v84 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v76)+96))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+28))
	v80 = v79
	goto L10
L12:
	;
	v80 = int32(-1)
	goto L10
L13:
	;
	if v76 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L14:
	;
	if v84 == int32(0) {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84+base.B2i32(v87 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v93
	goto L15
L17:
	;
	v98 = l0 & int32(-4)
	v120 = v84
	goto L18
L18:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	if v132&v98 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L13
L20:
	;
	v362 = v21 + int32(72)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	if v364 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L21:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	if v136 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v148 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(64)))) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(56)))) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(48)))) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(40)))) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(32)))) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(24)))) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(16)))) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(8)))) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(561)
	if v146 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+48)))
	if v140&int32(8) == int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v331 = F_selectDb(m, v328, l3)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L32
	} else {
		goto L52
	}
L26:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	v268 = int32(0)
	v269 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	v270 = m.T0[v269].(func(*base.Module) int64)(m)
	mBase = m.M
	if v267 == v268 {
		goto L44
	} else {
		goto L45
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = int32(64)
	v171 = int32(0)
	v172 = *(*int32)(unsafe.Add(mBase, _consts[345]))
	if v172 == v171 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v203
	v207 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	v208 = int32(0)
	v209 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	v210 = m.T0[v209].(func(*base.Module) int64)(m)
	mBase = m.M
	if v207 == v208 {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	v191 = F_createClient(m, int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v175 = int32(0)
	v177 = v172 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[345])) = v177
	v180 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v180+v177<<(uint(int32(2))%32))))
	v186 = *(*int32)(unsafe.Add(mBase, _consts[349]))
	if base.Ui32(v186) <= base.Ui32(v177) {
		v203 = v184
		goto L28
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[349])) = v177
	v203 = v184
	goto L28
L32:
	;
	return
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+328)) = int32(0)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v191)+200)) = v195 | int32(1073741824)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v191)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v191)+204)) = v199 | int32(268435456)
	v203 = v191
	goto L28
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v224
	v228 = int32(0)
	v232 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	*(*int32)(unsafe.Add(mBase, _consts[275])) = v232 + int32(1)
	goto L39
L35:
	;
	v220 = *(*int64)(unsafe.Add(mBase, _consts[348]))
	v224 = v220*int64(1000) + v210
	goto L34
L36:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _consts[149]))
	v216 = base.I32_div_s(int32(1000000), v215)
	v224 = v210 + base.I64_extend_i32_s(v216)
	goto L34
L37:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v328 = v265
	goto L25
L38:
	;
	goto L37
L39:
	;
	if v232 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	goto L42
L41:
	;
	v242 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v240
	v246 = base.I64_div_s(v240, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[32])) = v246
	v250 = base.I64_div_s(v240, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[37])) = v250
	v253 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	F_lrulfu_updateClockAndPolicy(m, v246, int32(base.Ui32(v253&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v261 = *(*int64)(unsafe.Add(mBase, _consts[32]))
	*(*int64)(unsafe.Add(mBase, _consts[78])) = v261
	goto L38
L42:
	;
	v240 = F_ustime(m)
	mBase = m.M
	goto L41
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v284
	v288 = int32(0)
	v292 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	*(*int32)(unsafe.Add(mBase, _consts[275])) = v292 + int32(1)
	goto L48
L44:
	;
	v280 = *(*int64)(unsafe.Add(mBase, _consts[348]))
	v284 = v280*int64(1000) + v270
	goto L43
L45:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _consts[149]))
	v276 = base.I32_div_s(int32(1000000), v275)
	v284 = v270 + base.I64_extend_i32_s(v276)
	goto L43
L46:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v326
	v328 = v326
	goto L25
L47:
	;
	goto L46
L48:
	;
	if v292 != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	goto L51
L50:
	;
	v302 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v300
	v306 = base.I64_div_s(v300, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[32])) = v306
	v310 = base.I64_div_s(v300, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[37])) = v310
	v313 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	F_lrulfu_updateClockAndPolicy(m, v306, int32(base.Ui32(v313&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v321 = *(*int64)(unsafe.Add(mBase, _consts[32]))
	*(*int64)(unsafe.Add(mBase, _consts[78])) = v321
	goto L47
L51:
	;
	v300 = F_ustime(m)
	mBase = m.M
	goto L50
L52:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v333 | int32(1024)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v338 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v131)+12)) = v338
	v340 = int32(_a69)
	v342 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	*(*int32)(unsafe.Add(mBase, _consts[217])) = v342 + v338
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v347 = m.T0[v346].(func(*base.Module, int32, int32, int32, int32) int32)(m, v21, v98, l1, l2)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L32
	} else {
		goto L53
	}
L53:
	;
	v349 = int32(_a69)
	v351 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	*(*int32)(unsafe.Add(mBase, _consts[217])) = v351 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v131)+12)) = v337
	F_moduleFreeContext(m, v21)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L32
	} else {
		goto L54
	}
L54:
	;
	goto L20
L55:
	;
	if v364 != 0 {
		v120 = v364
		goto L18
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v364+base.B2i32(v367 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = v373
	goto L56
L58:
	;
	goto L19
L59:
	;
	v397 = int32(0)
	v399 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	*(*int32)(unsafe.Add(mBase, _consts[275])) = v399 + int32(-1)
	goto L62
L60:
	;
	v395 = F_selectDb(m, v76, v80)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L32
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	goto L1
}
func F_moduleNotifyKeyspaceSubscribersCnt(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _consts[380]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+20))
	return v3
}
func F_modulePipeReadable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v8 = m.G0
	v10 = v8 - int32(128)
	m.G0 = v10
	goto L1
L1:
	;
	v19 = int32(128)
	v20 = F_read(m, l1, v10, v19)
	mBase = m.M
	if v20 == v19 {
		goto L1
	} else {
		goto L3
	}
L2:
	;
	goto L4
L3:
	;
	goto L2
L4:
	;
	v25 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[398]))
	if v26 == v25 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	goto L17
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v29 == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v37 = v26
	goto L8
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	F_listDelNode(m, v37, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L5
L10:
	;
	return
L11:
	;
	goto L12
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	m.T0[v46].(func(*base.Module, int32))(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	F_valkey_free(m, v40)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L15
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[398]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	if v55 != 0 {
		v37 = v54
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	m.G0 = v10 + int32(128)
	return
}
func F_modulePopulateClientInfoStructure(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v80 int32
	_ = v80
	var v85 int64
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int64
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int64
	_ = v125
	var v129 int32
	_ = v129
	var v134 int64
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int64
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int64
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v180 int32
	_ = v180
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(1)
	if l2 != v11 {
		v180 = v11
		m.G0 = v9 + int32(16)
		return v180
	} else {
		v19 = F__emscripten_memset_bulkmem(m, l0+int32(8), base.I32_extend8_s(int32(0)), int32(72))
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(1)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
		if v23&int32(8) == int32(0) {
			v32 = v23
			v33 = int64(0)
		} else {
			v28 = int64(32)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v28
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
			v32 = v31
			v33 = v28
		}
		if v32&int32(262144) == int32(0) {
			v42 = v32
			v43 = v33
		} else {
			v39 = v33 | int64(2)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v39
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
			v42 = v41
			v43 = v39
		}
		if v42&int32(2048) == int32(0) {
			v51 = v43
		} else {
			v49 = v43 | int64(16)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v49
			v51 = v49
		}
		v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+204)))
		if v52&int32(4) == int32(0) {
			v60 = v51
		} else {
			v58 = v51 | int64(8)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v58
			v60 = v58
		}
		v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+200)))
		if v61&int32(16) == int32(0) {
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v60 | int64(4)
		}
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
		v71 = F_connectionTypeTls(m)
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return int32(0)
		} else {
			if v70 != v71 {
			} else {
				v76 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v76 | int64(1)
			}
			v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
			if v80&int32(131072) == int32(0) {
				v90 = v80
			} else {
				v85 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v85 | int64(64)
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
				v90 = v89
			}
			if v90&int32(1) == int32(0) {
				v100 = v90
			} else {
				v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v95 | int64(128)
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
				v100 = v99
			}
			if v100&int32(2) == int32(0) {
				v110 = v100
			} else {
				v105 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105 | int64(256)
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
				v110 = v109
			}
			if v110&int32(4) == int32(0) {
				v120 = v110
			} else {
				v115 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v115 | int64(512)
				v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
				v120 = v119
			}
			if v120&int32(1073741824) == int32(0) {
			} else {
				v125 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v125 | int64(1024)
			}
			v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
			if v129&int32(_a14) == int32(0) {
				v139 = v129
			} else {
				v134 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v134 | int64(2048)
				v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
				v139 = v138
			}
			if v139&int32(16777216) == int32(0) {
				v149 = v139
			} else {
				v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v144 | int64(4096)
				v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
				v149 = v148
			}
			if v149&int32(268435456) == int32(0) {
			} else {
				v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v154 | int64(8192)
			}
			v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v158 != 0 {
				v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
				v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+24))
				if v160 != 0 {
					v167 = m.T0[v160].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v158, l0+int32(24), int32(46), v9+int32(12), int32(1))
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int32(0)
					} else {
						v169 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						v171 = v169
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(v171)
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+28))
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v174)
						v176 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v176
						v180 = int32(0)
						m.G0 = v9 + int32(16)
						return v180
					}
				} else {
					v171 = v158
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(v171)
					v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+28))
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v174)
					v176 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v176
					v180 = int32(0)
					m.G0 = v9 + int32(16)
					return v180
				}
			} else {
				v171 = v158
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(v171)
				v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+28))
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v174)
				v176 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v176
				v180 = int32(0)
				m.G0 = v9 + int32(16)
				return v180
			}
		}
	}
}
func F_moduleReleaseTempClient(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
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
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[345]))
	v8 = *(*int32)(unsafe.Add(mBase, _consts[346]))
	if v6 != v8 {
		F_clearClientConnectionState(m, l0)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			F_listEmpty(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(0)
				F_resetClient(m, l0)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v36
					v38 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v38
					*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v38
					*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = int64(1152921505680588800)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v36
					*(*int64)(unsafe.Add(mBase, uint32(l0+int32(76)))) = v36
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					if v50 == v38 {
						v71 = int32(0)
						v73 = *(*int32)(unsafe.Add(mBase, _consts[345]))
						*(*int32)(unsafe.Add(mBase, _consts[345])) = v73 + int32(1)
						v78 = *(*int32)(unsafe.Add(mBase, _consts[347]))
						*(*int32)(unsafe.Add(mBase, uint32(v78+v73<<(uint(int32(2))%32)))) = l0
						return
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
						if v53 == int32(0) {
							v71 = int32(0)
							v73 = *(*int32)(unsafe.Add(mBase, _consts[345]))
							*(*int32)(unsafe.Add(mBase, _consts[345])) = v73 + int32(1)
							v78 = *(*int32)(unsafe.Add(mBase, _consts[347]))
							*(*int32)(unsafe.Add(mBase, uint32(v78+v73<<(uint(int32(2))%32)))) = l0
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = int32(0)
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
							v60 = v58 + int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v53))) = v60
							if v60 != 0 {
								v65 = v50
								*(*int32)(unsafe.Add(mBase, uint32(v65)+52)) = int32(0)
								v71 = int32(0)
								v73 = *(*int32)(unsafe.Add(mBase, _consts[345]))
								*(*int32)(unsafe.Add(mBase, _consts[345])) = v73 + int32(1)
								v78 = *(*int32)(unsafe.Add(mBase, _consts[347]))
								*(*int32)(unsafe.Add(mBase, uint32(v78+v73<<(uint(int32(2))%32)))) = l0
								return
							} else {
								F_valkey_free(m, v53)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
									v65 = v64
									*(*int32)(unsafe.Add(mBase, uint32(v65)+52)) = int32(0)
									v71 = int32(0)
									v73 = *(*int32)(unsafe.Add(mBase, _consts[345]))
									*(*int32)(unsafe.Add(mBase, _consts[345])) = v73 + int32(1)
									v78 = *(*int32)(unsafe.Add(mBase, _consts[347]))
									*(*int32)(unsafe.Add(mBase, uint32(v78+v73<<(uint(int32(2))%32)))) = l0
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		if v6 != 0 {
			v14 = v6 << (uint(int32(1)) % 32)
		} else {
			v14 = int32(32)
		}
		*(*int32)(unsafe.Add(mBase, _consts[346])) = v14
		v16 = int32(0)
		v18 = *(*int32)(unsafe.Add(mBase, _consts[347]))
		v21 = F_valkey_realloc(m, v18, v14<<(uint(int32(2))%32))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[347])) = v21
			F_clearClientConnectionState(m, l0)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				F_listEmpty(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(0)
					F_resetClient(m, l0)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v36 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v36
						v38 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v38
						*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v38
						*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = int64(1152921505680588800)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v36
						*(*int64)(unsafe.Add(mBase, uint32(l0+int32(76)))) = v36
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
						if v50 == v38 {
							v71 = int32(0)
							v73 = *(*int32)(unsafe.Add(mBase, _consts[345]))
							*(*int32)(unsafe.Add(mBase, _consts[345])) = v73 + int32(1)
							v78 = *(*int32)(unsafe.Add(mBase, _consts[347]))
							*(*int32)(unsafe.Add(mBase, uint32(v78+v73<<(uint(int32(2))%32)))) = l0
							return
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
							if v53 == int32(0) {
								v71 = int32(0)
								v73 = *(*int32)(unsafe.Add(mBase, _consts[345]))
								*(*int32)(unsafe.Add(mBase, _consts[345])) = v73 + int32(1)
								v78 = *(*int32)(unsafe.Add(mBase, _consts[347]))
								*(*int32)(unsafe.Add(mBase, uint32(v78+v73<<(uint(int32(2))%32)))) = l0
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = int32(0)
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
								v60 = v58 + int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v53))) = v60
								if v60 != 0 {
									v65 = v50
									*(*int32)(unsafe.Add(mBase, uint32(v65)+52)) = int32(0)
									v71 = int32(0)
									v73 = *(*int32)(unsafe.Add(mBase, _consts[345]))
									*(*int32)(unsafe.Add(mBase, _consts[345])) = v73 + int32(1)
									v78 = *(*int32)(unsafe.Add(mBase, _consts[347]))
									*(*int32)(unsafe.Add(mBase, uint32(v78+v73<<(uint(int32(2))%32)))) = l0
									return
								} else {
									F_valkey_free(m, v53)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
										v65 = v64
										*(*int32)(unsafe.Add(mBase, uint32(v65)+52)) = int32(0)
										v71 = int32(0)
										v73 = *(*int32)(unsafe.Add(mBase, _consts[345]))
										*(*int32)(unsafe.Add(mBase, _consts[345])) = v73 + int32(1)
										v78 = *(*int32)(unsafe.Add(mBase, _consts[347]))
										*(*int32)(unsafe.Add(mBase, uint32(v78+v73<<(uint(int32(2))%32)))) = l0
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
func F_moduleReplyErrorFormatInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v13&v10 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v26 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L3:
	;
	v25 = l0 + int32(8)
	goto L2
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v18 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v25 = v18 + int32(36)
	goto L2
L6:
	;
	if l2&int32(3) == int32(0) {
		v50 = l2
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v88 = v11 - (v83+int32(17))&int32(-16)
	m.G0 = v88
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
	v94 = F_snprintf(m, v88, v83+int32(2), int32(_a704), v11)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L23
	} else {
		goto L24
	}
L8:
	;
	v83 = v75 - l2
	goto L7
L9:
	;
	v54 = v50
	goto L17
L10:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = l2
	goto L13
L12:
	;
	v83 = l2 - l2
	goto L7
L13:
	;
	v43 = v39 + int32(1)
	if v43&int32(3) == int32(0) {
		v50 = v43
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v48 != 0 {
		v39 = v43
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v75 = v43
	goto L8
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v63 = int32(-2139062144)
	if (int32(16843008)-v60|v60)&v63 == v63 {
		v54 = v54 + int32(4)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v69 = v54
	goto L20
L19:
	;
	goto L18
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v73 != 0 {
		v69 = v69 + int32(1)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v75 = v69
	goto L8
L22:
	;
	goto L21
L23:
	;
	return
L24:
	;
	F_addReplyErrorFormatInternal(m, v26, l1, v88, l3)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	goto L1
}
func F_moduleReplyWithCollection(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v10&int32(16) == int32(0) {
		v22 = l0 + int32(8)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		if v23 == int32(0) {
			m.G0 = v8 + int32(32)
			return int32(0)
		} else {
			switch l1 + int32(1) {
			case 0:
				v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v87 = F_valkey_realloc(m, v81, v82<<(uint(int32(2))%32)+int32(4))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v87
					v90 = F_addReplyDeferredLen(m, v23)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v92+v93<<(uint(int32(2))%32)))) = v90
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v93 + int32(1)
						m.G0 = v8 + int32(32)
						return int32(0)
					}
				}
			case 1:
				switch l2 + int32(-1) {
				case 0:
					v31 = *(*int32)(unsafe.Add(mBase, _consts[289]))
					F_addReply(m, v23, v31)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(32)
						return int32(0)
					}
				case 1:
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+224)))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v37<<(uint(int32(2))%32))+uint32(_consts[372])))
					F_addReply(m, v23, v41)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(32)
						return int32(0)
					}
				case 2:
					v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+224)))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v45<<(uint(int32(2))%32))+uint32(_consts[373])))
					F_addReply(m, v23, v49)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(32)
						return int32(0)
					}
				case 3:
					F_addReplyAttributeLen(m, v23, int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(32)
						return int32(0)
					}
				default:
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
					F__serverPanic_1(m, int32(_a694), int32(3326), int32(_a705), v8+int32(16))
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
				}
			default:
				switch l2 + int32(-1) {
				case 0:
					F_addReplyArrayLen(m, v23, l1)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(32)
						return int32(0)
					}
				case 1:
					F_addReplyMapLen(m, v23, l1)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(32)
						return int32(0)
					}
				case 2:
					F_addReplySetLen(m, v23, l1)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(32)
						return int32(0)
					}
				case 3:
					F_addReplyAttributeLen(m, v23, l1)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(32)
						return int32(0)
					}
				default:
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l2
					F__serverPanic_1(m, int32(_a694), int32(3334), int32(_a706), v8)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
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
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v15 == int32(0) {
			m.G0 = v8 + int32(32)
			return int32(0)
		} else {
			v22 = v15 + int32(36)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			if v23 == int32(0) {
				m.G0 = v8 + int32(32)
				return int32(0)
			} else {
				switch l1 + int32(1) {
				case 0:
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v87 = F_valkey_realloc(m, v81, v82<<(uint(int32(2))%32)+int32(4))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v87
						v90 = F_addReplyDeferredLen(m, v23)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v92+v93<<(uint(int32(2))%32)))) = v90
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v93 + int32(1)
							m.G0 = v8 + int32(32)
							return int32(0)
						}
					}
				case 1:
					switch l2 + int32(-1) {
					case 0:
						v31 = *(*int32)(unsafe.Add(mBase, _consts[289]))
						F_addReply(m, v23, v31)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(32)
							return int32(0)
						}
					case 1:
						v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+224)))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v37<<(uint(int32(2))%32))+uint32(_consts[372])))
						F_addReply(m, v23, v41)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(32)
							return int32(0)
						}
					case 2:
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+224)))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v45<<(uint(int32(2))%32))+uint32(_consts[373])))
						F_addReply(m, v23, v49)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(32)
							return int32(0)
						}
					case 3:
						F_addReplyAttributeLen(m, v23, int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(32)
							return int32(0)
						}
					default:
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
						F__serverPanic_1(m, int32(_a694), int32(3326), int32(_a705), v8+int32(16))
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
					}
				default:
					switch l2 + int32(-1) {
					case 0:
						F_addReplyArrayLen(m, v23, l1)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(32)
							return int32(0)
						}
					case 1:
						F_addReplyMapLen(m, v23, l1)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(32)
							return int32(0)
						}
					case 2:
						F_addReplySetLen(m, v23, l1)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(32)
							return int32(0)
						}
					case 3:
						F_addReplyAttributeLen(m, v23, l1)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(32)
							return int32(0)
						}
					default:
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = l2
						F__serverPanic_1(m, int32(_a694), int32(3334), int32(_a706), v8)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
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
		}
	}
}
func F_moduleScanKeyHashtableCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 float64
	_ = v154
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v171 int64
	_ = v171
	var v191 int64
	_ = v191
	var v205 int32
	_ = v205
	var v214 int64
	_ = v214
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v219 int64
	_ = v219
	var v220 int64
	_ = v220
	var v233 int64
	_ = v233
	var v236 int64
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	switch v15&int32(15) + int32(-2) {
	case 0:
		v240 = l1
		v241 = v3
		v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+int32(-1)))))
		switch v245 & int32(7) {
		case 0:
			v262 = int32(base.Ui32(v245) >> (uint(int32(3)) % 32))
		case 1:
			v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+int32(-3)))))
			v262 = v252
		case 2:
			v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240+int32(-5)))))
			v262 = v255
		case 3:
			v258 = *(*int32)(unsafe.Add(mBase, uint32(v240+int32(-9))))
			v262 = v258
		case 4:
			v261 = *(*int32)(unsafe.Add(mBase, uint32(v240+int32(-17))))
			v262 = v261
		default:
			v262 = v3
		}
		v263 = F_createStringObject_1(m, v240, v262)
		mBase = m.M
		v264 = m.ExcPending
		if v264 != 0 {
			return
		} else {
			v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			m.T0[v267].(func(*base.Module, int32, int32, int32, int32))(m, v265, v263, v241, v266)
			mBase = m.M
			v269 = m.ExcPending
			if v269 != 0 {
				return
			} else {
				F_decrRefCount(m, v263)
				mBase = m.M
				v271 = m.ExcPending
				if v271 != 0 {
					return
				} else {
					if v241 == int32(0) {
						m.G0 = v9 + int32(32)
						return
					} else {
						F_decrRefCount(m, v241)
						mBase = m.M
						v275 = m.ExcPending
						if v275 != 0 {
							return
						} else {
							m.G0 = v9 + int32(32)
							return
						}
					}
				}
			}
		}
	case 1:
		v143 = l1 + int32(16)
		v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
		v147 = v143 + v144<<(uint(int32(3))%32)
		v148 = int32(*(*int8)(unsafe.Add(mBase, uint32(v147))))
		v153 = v9 + int32(8)
		v154 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
		v161 = m.G0
		v163 = v161 - int32(16)
		m.G0 = v163
		v165 = base.I64_reinterpret_f64(v154)
		v167 = v165 & int64(4503599627370495)
		v171 = int64(base.Ui64(v165)>>(uint(int64(52))%64)) & int64(2047)
		if v171 == int64(0) {
			if base.B2i32(v167 == int64(0)) == int32(0) {
				if base.Ui64(v167) < base.Ui64(int64(4294967296)) {
					v205 = base.I32_clz(base.I32_wrap_i64(v165)) | int32(32)
				} else {
					v205 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v167) >> (uint(int64(32)) % 64))))
				}
				F___ashlti3(m, v163, v167, int64(0), v205+int32(49))
				mBase = m.M
				v214 = *(*int64)(unsafe.Add(mBase, uint32(v163+int32(8))))
				v217 = *(*int64)(unsafe.Add(mBase, uint32(v163)))
				v218 = v217
				v219 = base.I64_extend_i32_u(int32(15372) - v205)
				v220 = v214 ^ int64(281474976710656)
			} else {
				v191 = int64(0)
				v218 = v191
				v219 = v191
				v220 = v191
			}
		} else {
			if v171 == int64(2047) {
				v218 = v167 << (uint(int64(60)) % 64)
				v219 = int64(32767)
				v220 = int64(base.Ui64(v167) >> (uint(int64(4)) % 64))
			} else {
				v218 = v167 << (uint(int64(60)) % 64)
				v219 = v171 + int64(15360)
				v220 = int64(base.Ui64(v167) >> (uint(int64(4)) % 64))
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(v153))) = v218
		*(*int64)(unsafe.Add(mBase, uint32(v153)+8)) = v219<<(uint(int64(48))%64) | v165&int64(-9223372036854775807-1) | v220
		m.G0 = v163 + int32(16)
		v233 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
		v236 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(16))))
		v238 = F_createStringObjectFromLongDouble(m, v233, v236, int32(0))
		mBase = m.M
		v239 = m.ExcPending
		if v239 != 0 {
			return
		} else {
			v240 = v147 + v148 + int32(1)
			v241 = v238
			v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+int32(-1)))))
			switch v245 & int32(7) {
			case 0:
				v262 = int32(base.Ui32(v245) >> (uint(int32(3)) % 32))
			case 1:
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+int32(-3)))))
				v262 = v252
			case 2:
				v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240+int32(-5)))))
				v262 = v255
			case 3:
				v258 = *(*int32)(unsafe.Add(mBase, uint32(v240+int32(-9))))
				v262 = v258
			case 4:
				v261 = *(*int32)(unsafe.Add(mBase, uint32(v240+int32(-17))))
				v262 = v261
			default:
				v262 = v3
			}
			v263 = F_createStringObject_1(m, v240, v262)
			mBase = m.M
			v264 = m.ExcPending
			if v264 != 0 {
				return
			} else {
				v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				m.T0[v267].(func(*base.Module, int32, int32, int32, int32))(m, v265, v263, v241, v266)
				mBase = m.M
				v269 = m.ExcPending
				if v269 != 0 {
					return
				} else {
					F_decrRefCount(m, v263)
					mBase = m.M
					v271 = m.ExcPending
					if v271 != 0 {
						return
					} else {
						if v241 == int32(0) {
							m.G0 = v9 + int32(32)
							return
						} else {
							F_decrRefCount(m, v241)
							mBase = m.M
							v275 = m.ExcPending
							if v275 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	case 2:
		v21 = v9 + int32(28)
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
		v29 = v27 & int32(7)
		if v29 == int32(0) {
			switch v29 {
			case 0:
				v49 = int32(base.Ui32(v27) >> (uint(int32(3)) % 32))
			case 1:
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
				v49 = v39
			case 2:
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
				v49 = v42
			case 3:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
				v49 = v45
			case 4:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
				v49 = v48
			default:
				v49 = int32(0)
			}
			v51 = int32(1)
			v52 = F_sdsHdrSize(m, v51)
			mBase = m.M
			v53 = l1 + v49 + v52
			v55 = v53 + v51
			if v21 == int32(0) {
				v123 = v55
				v131 = v123
			} else {
				v58 = int32(0)
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v58))))
				switch v61 & int32(7) {
				case 0:
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(base.Ui32(v61) >> (uint(int32(3)) % 32))
					v131 = v55
				case 1:
					v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(-2)))))
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v69
					v131 = v55
				case 2:
					v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53+int32(-4)))))
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v73
					v131 = v55
				case 3:
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(-8))))
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v77
					v131 = v55
				case 4:
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(-16))))
					v82 = v81
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v82
					v131 = v55
				default:
					v82 = v58
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v82
					v131 = v55
				}
			}
		} else {
			if v27&int32(16) != 0 {
				v84 = F_sdsAllocPtr(m, l1)
				mBase = m.M
				v86 = v84 + int32(-4)
				if v27&int32(32) == int32(0) {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
					if v21 == int32(0) {
						v123 = v98
					} else {
						v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+int32(-1)))))
						switch v104 & int32(7) {
						case 0:
							v121 = int32(base.Ui32(v104) >> (uint(int32(3)) % 32))
						case 1:
							v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+int32(-3)))))
							v121 = v111
						case 2:
							v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98+int32(-5)))))
							v121 = v114
						case 3:
							v117 = *(*int32)(unsafe.Add(mBase, uint32(v98+int32(-9))))
							v121 = v117
						case 4:
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v98+int32(-17))))
							v121 = v120
						default:
							v121 = int32(0)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v121
						v123 = v98
					}
					v131 = v123
				} else {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
					if v91 != 0 {
						if v21 == int32(0) {
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = v95
						}
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
						v131 = v97
					} else {
						v131 = int32(0)
					}
				}
			} else {
				switch v29 {
				case 0:
					v49 = int32(base.Ui32(v27) >> (uint(int32(3)) % 32))
				case 1:
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
					v49 = v39
				case 2:
					v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
					v49 = v42
				case 3:
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
					v49 = v45
				case 4:
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
					v49 = v48
				default:
					v49 = int32(0)
				}
				v51 = int32(1)
				v52 = F_sdsHdrSize(m, v51)
				mBase = m.M
				v53 = l1 + v49 + v52
				v55 = v53 + v51
				if v21 == int32(0) {
					v123 = v55
					v131 = v123
				} else {
					v58 = int32(0)
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v58))))
					switch v61 & int32(7) {
					case 0:
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(base.Ui32(v61) >> (uint(int32(3)) % 32))
						v131 = v55
					case 1:
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(-2)))))
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v69
						v131 = v55
					case 2:
						v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53+int32(-4)))))
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v73
						v131 = v55
					case 3:
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(-8))))
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v77
						v131 = v55
					case 4:
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(-16))))
						v82 = v81
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v82
						v131 = v55
					default:
						v82 = v58
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v82
						v131 = v55
					}
				}
			}
		}
		v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
		v133 = F_createStringObject_1(m, v131, v132)
		mBase = m.M
		v134 = m.ExcPending
		if v134 != 0 {
			return
		} else {
			v240 = l1
			v241 = v133
			v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+int32(-1)))))
			switch v245 & int32(7) {
			case 0:
				v262 = int32(base.Ui32(v245) >> (uint(int32(3)) % 32))
			case 1:
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+int32(-3)))))
				v262 = v252
			case 2:
				v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240+int32(-5)))))
				v262 = v255
			case 3:
				v258 = *(*int32)(unsafe.Add(mBase, uint32(v240+int32(-9))))
				v262 = v258
			case 4:
				v261 = *(*int32)(unsafe.Add(mBase, uint32(v240+int32(-17))))
				v262 = v261
			default:
				v262 = v3
			}
			v263 = F_createStringObject_1(m, v240, v262)
			mBase = m.M
			v264 = m.ExcPending
			if v264 != 0 {
				return
			} else {
				v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				m.T0[v267].(func(*base.Module, int32, int32, int32, int32))(m, v265, v263, v241, v266)
				mBase = m.M
				v269 = m.ExcPending
				if v269 != 0 {
					return
				} else {
					F_decrRefCount(m, v263)
					mBase = m.M
					v271 = m.ExcPending
					if v271 != 0 {
						return
					} else {
						if v241 == int32(0) {
							m.G0 = v9 + int32(32)
							return
						} else {
							F_decrRefCount(m, v241)
							mBase = m.M
							v275 = m.ExcPending
							if v275 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	default:
		F__serverPanic_1(m, int32(_a694), int32(12044), int32(_a734), int32(0))
		mBase = m.M
		v140 = m.ExcPending
		if v140 != 0 {
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
func F_moduleTimerHandler(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int64
	_ = v205
	var v209 int64
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v225 int64
	_ = v225
	var v227 int32
	_ = v227
	var v231 int64
	_ = v231
	var v235 int64
	_ = v235
	var v238 int32
	_ = v238
	var v246 int64
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v511 int64
	_ = v511
	var v514 int64
	_ = v514
	var v527 int64
	_ = v527
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int64
	_ = v538
	var v541 int64
	_ = v541
	var v544 int64
	_ = v544
	var v545 int64
	_ = v545
	var v549 int64
	_ = v549
	v5 = int64(0)
	v18 = m.G0
	v20 = v18 - int32(400)
	m.G0 = v20
	v23 = v20 + int32(96)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[383]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+12)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v23)+296)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v23)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v20 + int32(120)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+156)) = v20 + int32(264)
	goto L1
L1:
	;
	v43 = F_ustime(m)
	mBase = m.M
	v47 = int32(0)
	v49 = F_raxSeek(m, v20+int32(96), int32(_a4), v47, v47)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int64(0)
L3:
	;
	v55 = F_raxNext(m, v20+int32(96))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L7
	}
L4:
	;
	F_raxStop(m, v20+int32(96))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L2
	} else {
		goto L75
	}
L5:
	;
	v511 = F_ustime(m)
	mBase = m.M
	v514 = base.I64_div_s(v128-v511, int64(1000))
	v527 = v514
	goto L4
L6:
	;
	v527 = int64(0)
	goto L4
L7:
	;
	if v55 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
	v94 = int64(56)
	v96 = int64(65280)
	v98 = int64(40)
	v101 = int64(16711680)
	v103 = int64(24)
	v105 = int64(4278190080)
	v107 = int64(8)
	v128 = v93<<(uint(v94)%64) | v93&v96<<(uint(v98)%64) | (v93&v101<<(uint(v103)%64) | v93&v105<<(uint(v107)%64)) | (int64(base.Ui64(v93)>>(uint(v107)%64))&v105 | int64(base.Ui64(v93)>>(uint(v103)%64))&v101 | (int64(base.Ui64(v93)>>(uint(v98)%64))&v96 | int64(base.Ui64(v93)>>(uint(v94)%64))))
	if base.Ui64(v43) < base.Ui64(v128) {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L6
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+88)) = v93
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v133 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(80)))) = v133
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(72)))) = v133
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(64)))) = v133
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(56)))) = v133
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(48)))) = v133
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(40)))) = v133
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(32)))) = v133
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(24)))) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(561)
	v156 = int32(0)
	v157 = *(*int32)(unsafe.Add(mBase, _consts[345]))
	if v157 == v156 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v188
	v192 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	v193 = int32(0)
	v194 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	v195 = m.T0[v194].(func(*base.Module) int64)(m)
	mBase = m.M
	if v192 == v193 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v176 = F_createClient(m, int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L16
	}
L14:
	;
	v160 = int32(0)
	v162 = v157 + int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[345])) = v162
	v165 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165+v162<<(uint(int32(2))%32))))
	v171 = *(*int32)(unsafe.Add(mBase, _consts[349]))
	if base.Ui32(v171) <= base.Ui32(v162) {
		v188 = v169
		goto L12
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[349])) = v162
	v188 = v169
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+328)) = int32(0)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v176)+200)) = v180 | int32(1073741824)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v176)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v176)+204)) = v184 | int32(268435456)
	v188 = v176
	goto L12
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+72)) = v209
	v213 = int32(0)
	v217 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	*(*int32)(unsafe.Add(mBase, _consts[275])) = v217 + int32(1)
	goto L22
L18:
	;
	v205 = *(*int64)(unsafe.Add(mBase, _consts[348]))
	v209 = v205*int64(1000) + v195
	goto L17
L19:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _consts[149]))
	v201 = base.I32_div_s(int32(1000000), v200)
	v209 = v195 + base.I64_extend_i32_s(v201)
	goto L17
L20:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v252 = F_selectDb(m, v250, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L2
	} else {
		goto L26
	}
L21:
	;
	goto L20
L22:
	;
	if v217 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	goto L25
L24:
	;
	v227 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v225
	v231 = base.I64_div_s(v225, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[32])) = v231
	v235 = base.I64_div_s(v225, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[37])) = v235
	v238 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	F_lrulfu_updateClockAndPolicy(m, v231, int32(base.Ui32(v238&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v246 = *(*int64)(unsafe.Add(mBase, _consts[32]))
	*(*int64)(unsafe.Add(mBase, _consts[78])) = v246
	goto L21
L25:
	;
	v225 = F_ustime(m)
	mBase = m.M
	goto L24
L26:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	m.T0[v257].(func(*base.Module, int32, int32))(m, v20+int32(16), v256)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	F_moduleFreeContext(m, v20+int32(16))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v264 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v264
	v267 = *(*int32)(unsafe.Add(mBase, _consts[383]))
	v269 = v20 + int32(88)
	v270 = int32(8)
	v272 = v20 + int32(12)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	goto L33
L29:
	;
	v485 = int32(0)
	v487 = F_raxSeek(m, v20+int32(96), int32(_a4), v485, v485)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L2
	} else {
		goto L72
	}
L30:
	;
	if v466 == int32(0) {
		goto L29
	} else {
		goto L68
	}
L31:
	;
	if v425 != v270 {
		v466 = v264
		goto L58
	} else {
		goto L59
	}
L32:
	;
	v416 = int32(0)
	v422 = v281
	v423 = v282
	v425 = v416
	v429 = v416
	goto L31
L33:
	;
	if base.Ui32(v282) < base.Ui32(int32(8)) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v293 = v281
	v294 = v282
	v296 = int32(0)
	goto L36
L35:
	;
	v422 = v406
	v423 = v407
	v425 = v409
	v429 = base.B2i32(v412 != int32(0))
	goto L31
L36:
	;
	v302 = int32(base.Ui32(v294) >> (uint(int32(3)) % 32))
	v303 = int32(4)
	v304 = v293 + v303
	if v294&v303 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v406 = v397
	v407 = v398
	v409 = v382
	v412 = v387
	goto L35
L38:
	;
	v387 = int32(0)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v304+v302+(v387-v302)&int32(3)+v375<<(uint(int32(2))%32))))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	if base.Ui32(v398) < base.Ui32(int32(8)) {
		v406 = v397
		v407 = v398
		v409 = v382
		v412 = v387
		goto L35
	} else {
		goto L56
	}
L39:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269+v296))))
	v353 = int32(0)
	goto L50
L40:
	;
	v309 = int32(0)
	if base.Ui32(v270) <= base.Ui32(v296) {
		v342 = v296
		v345 = v309
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if v345 == v302 {
		v375 = v309
		v382 = v342
		goto L38
	} else {
		goto L48
	}
L42:
	;
	v319 = v296
	v322 = v309
	goto L43
L43:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304+v322))))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269+v319))))
	if v325 != v327 {
		v342 = v319
		v345 = v322
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v342 = v330
	v345 = v332
	goto L41
L45:
	;
	v329 = int32(1)
	v330 = v319 + v329
	v332 = v322 + v329
	if base.Ui32(v302) <= base.Ui32(v332) {
		v342 = v330
		v345 = v332
		goto L41
	} else {
		goto L46
	}
L46:
	;
	if base.Ui32(v330) < base.Ui32(v270) {
		v319 = v330
		v322 = v332
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v406 = v293
	v407 = v294
	v409 = v342
	v412 = v345
	goto L35
L49:
	;
	if v353 != v302 {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304+v353))))
	if v366 == v350&int32(255) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v368 = int32(1)
	v370 = v353 + v368
	if v370 != v302 {
		v353 = v370
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v422 = v293
	v423 = v294
	v425 = v296
	v429 = v368
	goto L31
L54:
	;
	v375 = v353
	v382 = v296 + int32(1)
	goto L38
L55:
	;
	v406 = v293
	v407 = v294
	v409 = v296
	v412 = v302
	goto L35
L56:
	;
	if base.Ui32(v382) < base.Ui32(v270) {
		v293 = v397
		v294 = v398
		v296 = v382
		goto L36
	} else {
		goto L57
	}
L57:
	;
	goto L37
L58:
	;
	goto L30
L59:
	;
	v431 = int32(0)
	if v423&int32(1) == v431 {
		v466 = v431
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v437 = v423 & int32(4)
	if v429&base.B2i32(v437 != int32(0)) != 0 {
		v466 = v431
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v441 = int32(1)
	if v272 == int32(0) {
		v466 = v441
		goto L58
	} else {
		goto L62
	}
L62:
	;
	if v423&int32(2) != 0 {
		v463 = int32(0)
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = v463
	v466 = v441
	goto L58
L64:
	;
	v447 = int32(3)
	v448 = int32(base.Ui32(v423) >> (uint(v447) % 32))
	if v437 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v458 = int32(4)
	goto L67
L66:
	;
	v458 = v448 << (uint(int32(2)) % 32)
	goto L67
L67:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v422+v448+(int32(0)-v448)&v447+v458+int32(4))))
	v463 = v462
	goto L63
L68:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v470 != v131 {
		goto L29
	} else {
		goto L69
	}
L69:
	;
	v472 = int32(0)
	v473 = *(*int32)(unsafe.Add(mBase, _consts[383]))
	v478 = F_raxRemove(m, v473, v20+int32(88), int32(8), v472)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	F_valkey_free(m, v131)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	goto L29
L72:
	;
	v491 = F_raxNext(m, v20+int32(96))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	if v491 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	goto L10
L75:
	;
	v537 = *(*int32)(unsafe.Add(mBase, _consts[383]))
	v538 = *(*int64)(unsafe.Add(mBase, uint32(v537)+8))
	goto L78
L76:
	;
	m.G0 = v20 + int32(400)
	return v549
L77:
	;
	v545 = int64(-1)
	*(*int64)(unsafe.Add(mBase, _consts[384])) = v545
	v549 = v545
	goto L76
L78:
	;
	if v538 == int64(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v541 = int64(1)
	if v541 < v527 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v544 = v527
	goto L82
L81:
	;
	v544 = v541
	goto L82
L82:
	;
	v549 = v544
	goto L76
}
func F_moduleTypeLookupModuleByNameInternal(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int64
	_ = v324
	var v325 int64
	_ = v325
	var v326 int64
	_ = v326
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v334 int64
	_ = v334
	var v336 int64
	_ = v336
	var v338 int64
	_ = v338
	var v340 int64
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	v13 = F_dictGetIterator(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v413
L2:
	;
	F_dictReleaseIterator(m, v13)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L3
	} else {
		goto L109
	}
L3:
	;
	return int32(0)
L4:
	;
	v24 = v13 + int32(20)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v120 == int32(0) {
		goto L2
	} else {
		goto L31
	}
L6:
	;
	v31 = v24
	v32 = v28
	goto L9
L7:
	;
	v28 = int32(1)
	goto L6
L8:
	;
	v28 = int32(0)
	goto L6
L9:
	;
	switch v32 {
	case 0:
		goto L14
	default:
		goto L13
	}
L11:
	;
	v32 = int32(0)
	goto L9
L12:
	;
	goto L5
L13:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v112
	if v112 == int32(0) {
		goto L11
	} else {
		goto L30
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v36 != int32(-1) {
		v75 = v36
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v76 = int32(1)
	v77 = v75 + v76
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v77
	v79 = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82+v83+int32(26)))))
	if v87 == int32(255) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v40 != 0 {
		v75 = int32(-1)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v42 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	if v69 != int32(-1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v41)+16)))
	v50 = int64(*(*int8)(unsafe.Add(mBase, uint32(v41)+27)))
	v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v41)+8)))
	v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v41)+12)))
	v53 = int64(*(*int8)(unsafe.Add(mBase, uint32(v41)+26)))
	v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v41)+4)))
	v55 = F_wangHash64(m, v54)
	mBase = m.M
	v57 = F_wangHash64(m, v53+v55)
	mBase = m.M
	v59 = F_wangHash64(m, v52+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v51+v59)
	mBase = m.M
	v63 = F_wangHash64(m, v50+v61)
	mBase = m.M
	v65 = F_wangHash64(m, v49+v63)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v68 = v67
	goto L18
L20:
	;
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+24)))
	v47 = v45 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+24)) = uint16(v47)
	v68 = v41
	goto L18
L21:
	;
	v75 = v69 + int32(-1)
	goto L15
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v75 = v72
	goto L15
L23:
	;
	v102 = int32(2)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v82+v100<<(uint(v102)%32)+int32(4))))
	v31 = v107 + v101<<(uint(v102)%32)
	v32 = int32(1)
	goto L9
L24:
	;
	v91 = v79
	goto L26
L25:
	;
	v91 = v76 << (uint(v87) % 32)
	goto L26
L26:
	;
	if v77 < v91 {
		v100 = v83
		v101 = v77
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if v83 != 0 {
		v120 = v79
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v93 == int32(-1) {
		v120 = v79
		goto L12
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+4)) = int64(4294967296)
	v100 = int32(1)
	v101 = int32(0)
	goto L23
L30:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v116
	v120 = v112
	goto L12
L31:
	;
	v130 = v120
	goto L32
L32:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	goto L34
L33:
	;
	goto L2
L34:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+16))
	v135 = v9 + int32(8)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v136
	goto L35
L35:
	;
	v141 = v9 + int32(8)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v143 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v299 = v13 + int32(20)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v300 != 0 {
		goto L84
	} else {
		goto L85
	}
L37:
	;
	if v143 == int32(0) {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v143+base.B2i32(v146 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v152
	goto L38
L40:
	;
	v160 = v143
	goto L41
L41:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	v164 = v162 + int32(84)
	if l1 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	goto L36
L43:
	;
	v273 = v9 + int32(8)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	if v275 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L44:
	;
	F_dictReleaseIterator(m, v13)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L3
	} else {
		goto L77
	}
L45:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v234 != 0 {
		goto L66
	} else {
		goto L67
	}
L46:
	;
	v165 = int32(10)
	goto L51
L47:
	;
	if v229 == int32(0) {
		goto L44
	} else {
		goto L63
	}
L48:
	;
	v229 = int32(0)
	goto L47
L49:
	;
	v201 = v196
	v202 = v197
	v203 = v198
	goto L59
L50:
	;
	if v186 == int32(0) {
		goto L48
	} else {
		goto L57
	}
L51:
	;
	if (v164|l0)&int32(3) != 0 {
		v196 = l0
		v197 = v164
		v198 = v165
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v173 = l0
	v174 = v164
	v175 = v165
	goto L53
L53:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if v178 != v179 {
		v196 = v173
		v197 = v174
		v198 = v175
		goto L49
	} else {
		goto L55
	}
L54:
	;
	goto L50
L55:
	;
	v181 = int32(4)
	v182 = v174 + v181
	v184 = v173 + v181
	v186 = v175 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v186) {
		v173 = v184
		v174 = v182
		v175 = v186
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v196 = v184
	v197 = v182
	v198 = v186
	goto L49
L58:
	;
	v229 = v206 - v207
	goto L47
L59:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v206 != v207 {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v209 = int32(1)
	v214 = v203 + int32(-1)
	if v214 == int32(0) {
		goto L48
	} else {
		goto L62
	}
L62:
	;
	v201 = v201 + v209
	v202 = v202 + v209
	v203 = v214
	goto L59
L63:
	;
	goto L43
L64:
	;
	if v266-v268 != 0 {
		goto L43
	} else {
		goto L76
	}
L65:
	;
	v266 = F_tolower(m, v262)
	mBase = m.M
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	v268 = F_tolower(m, v267)
	mBase = m.M
	goto L64
L66:
	;
	v236 = l0
	v237 = v164
	v238 = v234
	goto L69
L67:
	;
	v262 = int32(0)
	v263 = v164
	goto L65
L68:
	;
	v262 = v259 & int32(255)
	v263 = v258
	goto L65
L69:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	if v240 == int32(0) {
		v258 = v237
		v259 = v238
		goto L68
	} else {
		goto L71
	}
L70:
	;
	v258 = v252
	v259 = int32(0)
	goto L68
L71:
	;
	v244 = v238 & int32(255)
	if v244 == v240 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v251 = int32(1)
	v252 = v237 + v251
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+1)))
	if v253 != 0 {
		v236 = v236 + v251
		v237 = v252
		v238 = v253
		goto L69
	} else {
		goto L75
	}
L73:
	;
	v246 = F_tolower(m, v244)
	mBase = m.M
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	v248 = F_tolower(m, v247)
	mBase = m.M
	if v246 == v248 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
	v258 = v237
	v259 = v250
	goto L68
L75:
	;
	goto L70
L76:
	;
	goto L44
L77:
	;
	v413 = v162
	goto L1
L78:
	;
	if v275 != 0 {
		v160 = v275
		goto L41
	} else {
		goto L81
	}
L79:
	;
	goto L78
L80:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v275+base.B2i32(v278 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v273))) = v284
	goto L79
L81:
	;
	goto L42
L82:
	;
	if v395 != 0 {
		v130 = v395
		goto L32
	} else {
		goto L108
	}
L83:
	;
	v306 = v299
	v307 = v303
	goto L86
L84:
	;
	v303 = int32(1)
	goto L83
L85:
	;
	v303 = int32(0)
	goto L83
L86:
	;
	switch v307 {
	case 0:
		goto L91
	default:
		goto L90
	}
L88:
	;
	v307 = int32(0)
	goto L86
L89:
	;
	goto L82
L90:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v387
	if v387 == int32(0) {
		goto L88
	} else {
		goto L107
	}
L91:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v311 != int32(-1) {
		v350 = v311
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v351 = int32(1)
	v352 = v350 + v351
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v352
	v354 = int32(0)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357+v358+int32(26)))))
	if v362 == int32(255) {
		goto L101
	} else {
		goto L102
	}
L93:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v315 != 0 {
		v350 = int32(-1)
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v317 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+20))
	if v344 != int32(-1) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v324 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v316)+16)))
	v325 = int64(*(*int8)(unsafe.Add(mBase, uint32(v316)+27)))
	v326 = int64(*(*int32)(unsafe.Add(mBase, uint32(v316)+8)))
	v327 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v316)+12)))
	v328 = int64(*(*int8)(unsafe.Add(mBase, uint32(v316)+26)))
	v329 = int64(*(*int32)(unsafe.Add(mBase, uint32(v316)+4)))
	v330 = F_wangHash64(m, v329)
	mBase = m.M
	v332 = F_wangHash64(m, v328+v330)
	mBase = m.M
	v334 = F_wangHash64(m, v327+v332)
	mBase = m.M
	v336 = F_wangHash64(m, v326+v334)
	mBase = m.M
	v338 = F_wangHash64(m, v325+v336)
	mBase = m.M
	v340 = F_wangHash64(m, v324+v338)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v340
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v343 = v342
	goto L95
L97:
	;
	v320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v316)+24)))
	v322 = v320 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v316)+24)) = uint16(v322)
	v343 = v316
	goto L95
L98:
	;
	v350 = v344 + int32(-1)
	goto L92
L99:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v350 = v347
	goto L92
L100:
	;
	v377 = int32(2)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v357+v375<<(uint(v377)%32)+int32(4))))
	v306 = v382 + v376<<(uint(v377)%32)
	v307 = int32(1)
	goto L86
L101:
	;
	v366 = v354
	goto L103
L102:
	;
	v366 = v351 << (uint(v362) % 32)
	goto L103
L103:
	;
	if v352 < v366 {
		v375 = v358
		v376 = v352
		goto L100
	} else {
		goto L104
	}
L104:
	;
	if v358 != 0 {
		v395 = v354
		goto L89
	} else {
		goto L105
	}
L105:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v357)+20))
	if v368 == int32(-1) {
		v395 = v354
		goto L89
	} else {
		goto L106
	}
L106:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+4)) = int64(4294967296)
	v375 = int32(1)
	v376 = int32(0)
	goto L100
L107:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v387)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v299))) = v391
	v395 = v387
	goto L89
L108:
	;
	goto L33
L109:
	;
	v413 = int32(0)
	goto L1
}
func F_moduleUnblockClient(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+48))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+44))
	if v7 != 0 {
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = int32(0)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v4)+48)) = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[378]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v14 != 0 {
		v22 = v13
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[379]))
		v19 = F_write(m, v16, int32(_a719), int32(1))
		mBase = m.M
		v21 = *(*int32)(unsafe.Add(mBase, _consts[378]))
		v22 = v21
	}
	v23 = F_listAddNodeTail(m, v22, v4)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		return
	}
}
func F_moduleUnsubscribeAllServerEvents(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	v14 = v9 + int32(8)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v15
	goto L1
L1:
	;
	v20 = v9 + int32(8)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v22 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return
L3:
	;
	if v22 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22+base.B2i32(v25 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v31
	goto L4
L6:
	;
	v37 = v22
	goto L7
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v42 != l0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v69 = v9 + int32(8)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v71 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v41)+8))
	v46 = v44 + int64(-20)
	if base.Ui64(int64(3)) < base.Ui64(v46) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	F_listDelNode(m, v61, v37)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v46)<<(uint(int32(2))%32))+uint32(_consts[394])))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v55 + int32(-1)
	goto L11
L13:
	;
	return
L14:
	;
	F_valkey_free(m, v41)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	if v71 != 0 {
		v37 = v71
		goto L7
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v71+base.B2i32(v74 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v80
	goto L17
L19:
	;
	goto L8
}
func F_moduleValidateCommandArgs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	v13 = int32(1)
	if l0 == int32(0) {
		v154 = v13
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(112)
	return v154
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 == int32(0) {
		v154 = v13
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = v16
	v25 = int32(0)
	v26 = l0
	goto L4
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if base.Ui32(v28) < base.Ui32(int32(9)) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v154 = v144
	goto L1
L6:
	;
	switch v28 + int32(-3) {
	case 0:
		goto L13
	default:
		goto L12
	case 3:
		goto L14
	}
L7:
	;
	v31 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v33 {
		v154 = v31
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v24
	F__serverLog(m, int32(3), int32(_a697), v11+int32(96))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v154 = v31
	goto L1
L11:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if base.Ui32(v95) < base.Ui32(int32(8)) {
		goto L24
	} else {
		goto L25
	}
L12:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if base.Ui32(v80+int32(1)) < base.Ui32(int32(2)) {
		goto L11
	} else {
		goto L21
	}
L13:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if int32(-1) < v65 {
		goto L11
	} else {
		goto L18
	}
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v48 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v49 = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v51 {
		v154 = v49
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0+v25*int32(40))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v57
	F__serverLog(m, int32(3), int32(_a698), v11+int32(64))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v154 = v49
	goto L1
L18:
	;
	v68 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v70 {
		v154 = v68
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v24
	F__serverLog(m, int32(3), int32(_a699), v11+int32(80))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v154 = v68
	goto L1
L21:
	;
	v85 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v87 {
		v154 = v85
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v24
	F__serverLog(m, int32(3), int32(_a700), v11)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v154 = v85
	goto L1
L24:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if base.Ui32(int32(1)) < base.Ui32(v28+int32(-7)) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v98 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v100 {
		v154 = v98
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v24
	F__serverLog(m, int32(3), int32(_a701), v11+int32(48))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v154 = v98
	goto L1
L28:
	;
	v144 = int32(1)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v147 = v25 + v144
	v149 = l0 + v145*v147
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	if v150 != 0 {
		v24 = v150
		v25 = v147
		v26 = v149
		goto L4
	} else {
		goto L40
	}
L29:
	;
	if v110 == int32(0) {
		goto L28
	} else {
		goto L37
	}
L30:
	;
	if v110 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v127 = F_moduleValidateCommandArgs(m, v110, l1)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L9
	} else {
		goto L35
	}
L32:
	;
	v115 = int32(0)
	v117 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v117 {
		v154 = v115
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v24
	F__serverLog(m, int32(3), int32(_a702), v11+int32(16))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v154 = v115
	goto L1
L35:
	;
	if v127 != 0 {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v154 = int32(0)
	goto L1
L37:
	;
	v132 = int32(0)
	v134 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v134 {
		v154 = v132
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v24
	F__serverLog(m, int32(3), int32(_a703), v11+int32(32))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v154 = v132
	goto L1
L40:
	;
	goto L5
}
func F_moduleVerifyAllAllowAtomicSlotMigrationOrReply(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, _consts[177]))
	v13 = F_dictGetIterator(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L4
L3:
	;
	F_dictReleaseIterator(m, v13)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L36
	}
L4:
	;
	v29 = v13 + int32(20)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v30 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v135
	F_addReplyErrorFormat(m, l0, int32(_a709), v8)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L35
	}
L6:
	;
	if v125 == int32(0) {
		v141 = v2
		goto L3
	} else {
		goto L32
	}
L7:
	;
	v36 = v29
	v37 = v33
	goto L10
L8:
	;
	v33 = int32(1)
	goto L7
L9:
	;
	v33 = int32(0)
	goto L7
L10:
	;
	switch v37 {
	case 0:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v37 = int32(0)
	goto L10
L13:
	;
	goto L6
L14:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v117
	if v117 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v41 != int32(-1) {
		v80 = v41
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v81 = int32(1)
	v82 = v80 + v81
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v82
	v84 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v88+int32(26)))))
	if v92 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v45 != 0 {
		v80 = int32(-1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v47 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	if v74 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v54 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v46)+16)))
	v55 = int64(*(*int8)(unsafe.Add(mBase, uint32(v46)+27)))
	v56 = int64(*(*int32)(unsafe.Add(mBase, uint32(v46)+8)))
	v57 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v46)+12)))
	v58 = int64(*(*int8)(unsafe.Add(mBase, uint32(v46)+26)))
	v59 = int64(*(*int32)(unsafe.Add(mBase, uint32(v46)+4)))
	v60 = F_wangHash64(m, v59)
	mBase = m.M
	v62 = F_wangHash64(m, v58+v60)
	mBase = m.M
	v64 = F_wangHash64(m, v57+v62)
	mBase = m.M
	v66 = F_wangHash64(m, v56+v64)
	mBase = m.M
	v68 = F_wangHash64(m, v55+v66)
	mBase = m.M
	v70 = F_wangHash64(m, v54+v68)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v73 = v72
	goto L19
L21:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+24)))
	v52 = v50 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v46)+24)) = uint16(v52)
	v73 = v46
	goto L19
L22:
	;
	v80 = v74 + int32(-1)
	goto L16
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v80 = v77
	goto L16
L24:
	;
	v107 = int32(2)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v87+v105<<(uint(v107)%32)+int32(4))))
	v36 = v112 + v106<<(uint(v107)%32)
	v37 = int32(1)
	goto L10
L25:
	;
	v96 = v84
	goto L27
L26:
	;
	v96 = v81 << (uint(v92) % 32)
	goto L27
L27:
	;
	if v82 < v96 {
		v105 = v88
		v106 = v82
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v88 != 0 {
		v125 = v84
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v87)+20))
	if v98 == int32(-1) {
		v125 = v84
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+4)) = int64(4294967296)
	v105 = int32(1)
	v106 = int32(0)
	goto L24
L31:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v121
	v125 = v117
	goto L13
L32:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	goto L33
L33:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+48)))
	if v132&int32(32) != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	goto L5
L35:
	;
	v141 = int32(-1)
	goto L3
L36:
	;
	m.G0 = v8 + int32(16)
	return v141
}
func F_processModuleLoadingProgressEvent(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	v13 = *(*int64)(unsafe.Add(mBase, _consts[395]))
	if v11 < v13 {
		m.G0 = v8 + int32(16)
		return
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, _consts[396]))
		if base.B2i32(v16 == int64(0)) == int32(0) {
			v23 = *(*int64)(unsafe.Add(mBase, _consts[397]))
			v26 = base.I64_div_s(v23<<(uint(int64(10))%64), v16)
			v28 = base.I32_wrap_i64(v26)
		} else {
			v28 = int32(-1)
		}
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(1)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v28
		v33 = *(*int32)(unsafe.Add(mBase, _consts[149]))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v33
		F_moduleFireServerEvent(m, int64(10), base.B2i32(l0 != int32(0)), v8)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, _consts[149]))
			v44 = base.I32_div_s(int32(1000000), v43)
			*(*int64)(unsafe.Add(mBase, _consts[395])) = v11 + base.I64_extend_i32_s(v44)
			m.G0 = v8 + int32(16)
			return
		}
	}
}
func F_setModuleBoolConfig(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = m.T0[v15].(func(*base.Module, int32, int32, int32, int32) int32)(m, v11, l1, v12, v7+int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v20 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v7 + int32(16)
	return base.B2i32(v16 == int32(0))
L4:
	;
	v24 = F_objectGetVal(m, v20)
	mBase = m.M
	goto L8
L5:
	;
	F_decrRefCount(m, v20)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L17
	}
L6:
	;
	goto L5
L7:
	;
	v55 = v34
	goto L14
L8:
	;
	v30 = int32(_a752)
	v32 = int32(256)
	v34 = v24
	goto L10
L9:
	;
	v45 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v45)
	goto L7
L10:
	;
	v36 = v32 + int32(-1)
	if v36 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v39)
	v41 = int32(1)
	if v39 != 0 {
		v30 = v30 + v41
		v32 = v36
		v34 = v34 + v41
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L6
L14:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v57 != 0 {
		v55 = v55 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	goto L6
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(_a752)
	goto L3
}
func F_setModuleEnumConfig(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = m.T0[v15].(func(*base.Module, int32, int32, int32, int32) int32)(m, v11, l1, v12, v7+int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v20 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v7 + int32(16)
	return base.B2i32(v16 == int32(0))
L4:
	;
	v24 = F_objectGetVal(m, v20)
	mBase = m.M
	goto L8
L5:
	;
	F_decrRefCount(m, v20)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L17
	}
L6:
	;
	goto L5
L7:
	;
	v55 = v34
	goto L14
L8:
	;
	v30 = int32(_a752)
	v32 = int32(256)
	v34 = v24
	goto L10
L9:
	;
	v45 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v45)
	goto L7
L10:
	;
	v36 = v32 + int32(-1)
	if v36 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v39)
	v41 = int32(1)
	if v39 != 0 {
		v30 = v30 + v41
		v32 = v36
		v34 = v34 + v41
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L6
L14:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v57 != 0 {
		v55 = v55 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	goto L6
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(_a752)
	goto L3
}
func F_setModuleNumericConfig(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v4
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = m.T0[v16].(func(*base.Module, int32, int64, int32, int32) int32)(m, v12, l1, v13, v8+int32(12))
	mBase = m.M
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v18 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return base.B2i32(v17 == int32(0))
L2:
	;
	v22 = F_objectGetVal(m, v18)
	mBase = m.M
	goto L6
L3:
	;
	F_decrRefCount(m, v18)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	goto L3
L5:
	;
	v53 = v32
	goto L12
L6:
	;
	v28 = int32(_a752)
	v30 = int32(256)
	v32 = v22
	goto L8
L7:
	;
	v43 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v43)
	goto L5
L8:
	;
	v34 = v30 + int32(-1)
	if v34 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v37)
	v39 = int32(1)
	if v37 != 0 {
		v28 = v28 + v39
		v30 = v34
		v32 = v32 + v39
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L4
L12:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v55 != 0 {
		v53 = v53 + int32(1)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	goto L4
L14:
	;
	goto L13
L15:
	;
	return int32(0)
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(_a752)
	goto L1
}
func F_setModuleStringConfig(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v4
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v16 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v33 = v4
		goto L1
	}
L1:
	;
	v34 = F_createStringObject_1(m, l1, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v33 = v32
	goto L1
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v33 = v29
	goto L1
L4:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v33 = v26
	goto L1
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v33 = v23
	goto L1
L6:
	;
	v33 = int32(base.Ui32(v16) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	return int32(0)
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = m.T0[v42].(func(*base.Module, int32, int32, int32, int32) int32)(m, v38, v34, v39, v9+int32(12))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v45 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_decrRefCount(m, v34)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L7
	} else {
		goto L25
	}
L11:
	;
	v49 = F_objectGetVal(m, v45)
	mBase = m.M
	goto L15
L12:
	;
	F_decrRefCount(m, v45)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L7
	} else {
		goto L24
	}
L13:
	;
	goto L12
L14:
	;
	v80 = v59
	goto L21
L15:
	;
	v55 = int32(_a752)
	v57 = int32(256)
	v59 = v49
	goto L17
L16:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v70)
	goto L14
L17:
	;
	v61 = v57 + int32(-1)
	if v61 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v64)
	v66 = int32(1)
	if v64 != 0 {
		v55 = v55 + v66
		v57 = v61
		v59 = v59 + v66
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L13
L21:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v82 != 0 {
		v80 = v80 + int32(1)
		goto L21
	} else {
		goto L23
	}
L22:
	;
	goto L13
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(_a752)
	goto L10
L25:
	;
	m.G0 = v9 + int32(16)
	return base.B2i32(v43 == int32(0))
}
