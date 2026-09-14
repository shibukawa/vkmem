package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_createSlotExportJob(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v63 int64
	_ = v63
	var v69 int64
	_ = v69
	var v75 int64
	_ = v75
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	v6 = F_valkey_calloc(m, int32(208))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = *(*int64)(unsafe.Add(mBase, _c_F_createSlotExportJob[0]))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v6)+164)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+156)) = int32(7)
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(0)
		v20 = F_representSlotRangeList(m, l1)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+168)) = v20
			F_getRandomHexChars(m, v6+int32(112), int32(40))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v30 = int32(40)
				v32 = *(*int64)(unsafe.Add(mBase, uint32(l0+v30)))
				*(*int64)(unsafe.Add(mBase, uint32(v6+int32(64)))) = v32
				v36 = int32(32)
				v38 = *(*int64)(unsafe.Add(mBase, uint32(l0+v36)))
				*(*int64)(unsafe.Add(mBase, uint32(v6+int32(56)))) = v38
				v42 = int32(24)
				v44 = *(*int64)(unsafe.Add(mBase, uint32(l0+v42)))
				*(*int64)(unsafe.Add(mBase, uint32(v6+int32(48)))) = v44
				v48 = int32(16)
				v50 = *(*int64)(unsafe.Add(mBase, uint32(l0+v48)))
				*(*int64)(unsafe.Add(mBase, uint32(v6+v30))) = v50
				v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v52
				v55 = *(*int32)(unsafe.Add(mBase, _c_F_createSlotExportJob[1]))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
				v57 = *(*int64)(unsafe.Add(mBase, uint32(v56)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v6)+72)) = v57
				v63 = *(*int64)(unsafe.Add(mBase, uint32(v56+v30)))
				*(*int64)(unsafe.Add(mBase, uint32(v6+int32(104)))) = v63
				v69 = *(*int64)(unsafe.Add(mBase, uint32(v56+v36)))
				*(*int64)(unsafe.Add(mBase, uint32(v6+int32(96)))) = v69
				v75 = *(*int64)(unsafe.Add(mBase, uint32(v56+v42)))
				*(*int64)(unsafe.Add(mBase, uint32(v6+int32(88)))) = v75
				v81 = *(*int64)(unsafe.Add(mBase, uint32(v56+v48)))
				*(*int64)(unsafe.Add(mBase, uint32(v6+int32(80)))) = v81
				v83 = F_generateSlotMigrationJobDescription(m, v6, l0)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+188)) = v83
					return v6
				}
			}
		}
	}
}
func F_generateSlotMigrationJobDescription(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	v7 = m.G0
	v9 = v7 - int32(64)
	m.G0 = v9
	if l1 != 0 {
		v13 = l1 + int32(8)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2316))
		if v14 == int32(0) {
			v72 = v13
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v75 = F_sdsempty(m)
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
				*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v79
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v72
				if v74 != 0 {
					v84 = int32(_a_F_generateSlotMigrationJobDescription_0)
				} else {
					v84 = int32(_a_F_generateSlotMigrationJobDescription_1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v84
				if v74 != 0 {
					v88 = int32(_a_F_generateSlotMigrationJobDescription_2)
				} else {
					v88 = int32(_a_F_generateSlotMigrationJobDescription_3)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v88
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0 + int32(112)
				v94 = F_sdscatprintf(m, v75, int32(_a_F_generateSlotMigrationJobDescription_4), v9)
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int32(0)
				} else {
					v96 = v94
					m.G0 = v9 + int32(64)
					return v96
				}
			}
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-1)))))
			switch v19 & int32(7) {
			case 0:
				v36 = int32(base.Ui32(v19) >> (uint(int32(3)) % 32))
				if v36 == int32(0) {
					v72 = v13
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v75 = F_sdsempty(m)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v79
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v72
						if v74 != 0 {
							v84 = int32(_a_F_generateSlotMigrationJobDescription_0)
						} else {
							v84 = int32(_a_F_generateSlotMigrationJobDescription_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v84
						if v74 != 0 {
							v88 = int32(_a_F_generateSlotMigrationJobDescription_2)
						} else {
							v88 = int32(_a_F_generateSlotMigrationJobDescription_3)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v88
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0 + int32(112)
						v94 = F_sdscatprintf(m, v75, int32(_a_F_generateSlotMigrationJobDescription_4), v9)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							v96 = v94
							m.G0 = v9 + int32(64)
							return v96
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v40 = F_sdsempty(m)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(56)))) = v46
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(52)))) = v14
						if v39 != 0 {
							v55 = int32(_a_F_generateSlotMigrationJobDescription_0)
						} else {
							v55 = int32(_a_F_generateSlotMigrationJobDescription_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(48)))) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v13
						*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v55
						if v39 != 0 {
							v61 = int32(_a_F_generateSlotMigrationJobDescription_2)
						} else {
							v61 = int32(_a_F_generateSlotMigrationJobDescription_3)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v61
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0 + int32(112)
						v69 = F_sdscatprintf(m, v40, int32(_a_F_generateSlotMigrationJobDescription_5), v9+int32(32))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v96 = v69
							m.G0 = v9 + int32(64)
							return v96
						}
					}
				}
			case 1:
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-3)))))
				v36 = v26
				if v36 == int32(0) {
					v72 = v13
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v75 = F_sdsempty(m)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v79
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v72
						if v74 != 0 {
							v84 = int32(_a_F_generateSlotMigrationJobDescription_0)
						} else {
							v84 = int32(_a_F_generateSlotMigrationJobDescription_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v84
						if v74 != 0 {
							v88 = int32(_a_F_generateSlotMigrationJobDescription_2)
						} else {
							v88 = int32(_a_F_generateSlotMigrationJobDescription_3)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v88
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0 + int32(112)
						v94 = F_sdscatprintf(m, v75, int32(_a_F_generateSlotMigrationJobDescription_4), v9)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							v96 = v94
							m.G0 = v9 + int32(64)
							return v96
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v40 = F_sdsempty(m)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(56)))) = v46
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(52)))) = v14
						if v39 != 0 {
							v55 = int32(_a_F_generateSlotMigrationJobDescription_0)
						} else {
							v55 = int32(_a_F_generateSlotMigrationJobDescription_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(48)))) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v13
						*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v55
						if v39 != 0 {
							v61 = int32(_a_F_generateSlotMigrationJobDescription_2)
						} else {
							v61 = int32(_a_F_generateSlotMigrationJobDescription_3)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v61
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0 + int32(112)
						v69 = F_sdscatprintf(m, v40, int32(_a_F_generateSlotMigrationJobDescription_5), v9+int32(32))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v96 = v69
							m.G0 = v9 + int32(64)
							return v96
						}
					}
				}
			case 2:
				v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(-5)))))
				v36 = v29
				if v36 == int32(0) {
					v72 = v13
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v75 = F_sdsempty(m)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v79
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v72
						if v74 != 0 {
							v84 = int32(_a_F_generateSlotMigrationJobDescription_0)
						} else {
							v84 = int32(_a_F_generateSlotMigrationJobDescription_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v84
						if v74 != 0 {
							v88 = int32(_a_F_generateSlotMigrationJobDescription_2)
						} else {
							v88 = int32(_a_F_generateSlotMigrationJobDescription_3)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v88
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0 + int32(112)
						v94 = F_sdscatprintf(m, v75, int32(_a_F_generateSlotMigrationJobDescription_4), v9)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							v96 = v94
							m.G0 = v9 + int32(64)
							return v96
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v40 = F_sdsempty(m)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(56)))) = v46
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(52)))) = v14
						if v39 != 0 {
							v55 = int32(_a_F_generateSlotMigrationJobDescription_0)
						} else {
							v55 = int32(_a_F_generateSlotMigrationJobDescription_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(48)))) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v13
						*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v55
						if v39 != 0 {
							v61 = int32(_a_F_generateSlotMigrationJobDescription_2)
						} else {
							v61 = int32(_a_F_generateSlotMigrationJobDescription_3)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v61
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0 + int32(112)
						v69 = F_sdscatprintf(m, v40, int32(_a_F_generateSlotMigrationJobDescription_5), v9+int32(32))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v96 = v69
							m.G0 = v9 + int32(64)
							return v96
						}
					}
				}
			case 3:
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-9))))
				v36 = v32
				if v36 == int32(0) {
					v72 = v13
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v75 = F_sdsempty(m)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v79
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v72
						if v74 != 0 {
							v84 = int32(_a_F_generateSlotMigrationJobDescription_0)
						} else {
							v84 = int32(_a_F_generateSlotMigrationJobDescription_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v84
						if v74 != 0 {
							v88 = int32(_a_F_generateSlotMigrationJobDescription_2)
						} else {
							v88 = int32(_a_F_generateSlotMigrationJobDescription_3)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v88
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0 + int32(112)
						v94 = F_sdscatprintf(m, v75, int32(_a_F_generateSlotMigrationJobDescription_4), v9)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							v96 = v94
							m.G0 = v9 + int32(64)
							return v96
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v40 = F_sdsempty(m)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(56)))) = v46
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(52)))) = v14
						if v39 != 0 {
							v55 = int32(_a_F_generateSlotMigrationJobDescription_0)
						} else {
							v55 = int32(_a_F_generateSlotMigrationJobDescription_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(48)))) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v13
						*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v55
						if v39 != 0 {
							v61 = int32(_a_F_generateSlotMigrationJobDescription_2)
						} else {
							v61 = int32(_a_F_generateSlotMigrationJobDescription_3)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v61
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0 + int32(112)
						v69 = F_sdscatprintf(m, v40, int32(_a_F_generateSlotMigrationJobDescription_5), v9+int32(32))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v96 = v69
							m.G0 = v9 + int32(64)
							return v96
						}
					}
				}
			case 4:
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-17))))
				v36 = v35
				if v36 == int32(0) {
					v72 = v13
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v75 = F_sdsempty(m)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v79
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v72
						if v74 != 0 {
							v84 = int32(_a_F_generateSlotMigrationJobDescription_0)
						} else {
							v84 = int32(_a_F_generateSlotMigrationJobDescription_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v84
						if v74 != 0 {
							v88 = int32(_a_F_generateSlotMigrationJobDescription_2)
						} else {
							v88 = int32(_a_F_generateSlotMigrationJobDescription_3)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v88
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0 + int32(112)
						v94 = F_sdscatprintf(m, v75, int32(_a_F_generateSlotMigrationJobDescription_4), v9)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							v96 = v94
							m.G0 = v9 + int32(64)
							return v96
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v40 = F_sdsempty(m)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(56)))) = v46
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(52)))) = v14
						if v39 != 0 {
							v55 = int32(_a_F_generateSlotMigrationJobDescription_0)
						} else {
							v55 = int32(_a_F_generateSlotMigrationJobDescription_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(48)))) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v13
						*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v55
						if v39 != 0 {
							v61 = int32(_a_F_generateSlotMigrationJobDescription_2)
						} else {
							v61 = int32(_a_F_generateSlotMigrationJobDescription_3)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v61
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0 + int32(112)
						v69 = F_sdscatprintf(m, v40, int32(_a_F_generateSlotMigrationJobDescription_5), v9+int32(32))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v96 = v69
							m.G0 = v9 + int32(64)
							return v96
						}
					}
				}
			default:
				v72 = v13
				v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v75 = F_sdsempty(m)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
					*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v79
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v72
					if v74 != 0 {
						v84 = int32(_a_F_generateSlotMigrationJobDescription_0)
					} else {
						v84 = int32(_a_F_generateSlotMigrationJobDescription_1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v84
					if v74 != 0 {
						v88 = int32(_a_F_generateSlotMigrationJobDescription_2)
					} else {
						v88 = int32(_a_F_generateSlotMigrationJobDescription_3)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v88
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0 + int32(112)
					v94 = F_sdscatprintf(m, v75, int32(_a_F_generateSlotMigrationJobDescription_4), v9)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v96 = v94
						m.G0 = v9 + int32(64)
						return v96
					}
				}
			}
		}
	} else {
		v72 = int32(_a_F_generateSlotMigrationJobDescription_6)
		v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v75 = F_sdsempty(m)
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return int32(0)
		} else {
			v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
			*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v79
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v72
			if v74 != 0 {
				v84 = int32(_a_F_generateSlotMigrationJobDescription_0)
			} else {
				v84 = int32(_a_F_generateSlotMigrationJobDescription_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v84
			if v74 != 0 {
				v88 = int32(_a_F_generateSlotMigrationJobDescription_2)
			} else {
				v88 = int32(_a_F_generateSlotMigrationJobDescription_3)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v88
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0 + int32(112)
			v94 = F_sdscatprintf(m, v75, int32(_a_F_generateSlotMigrationJobDescription_4), v9)
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return int32(0)
			} else {
				v96 = v94
				m.G0 = v9 + int32(64)
				return v96
			}
		}
	}
}
func F_rewriteSlotToAppendOnlyFileRio(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int64
	_ = v140
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int64
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v178 int32
	_ = v178
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v19 = int32(1)
	if l1 < v5 {
		v38 = v19
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(32)
	return v178
L2:
	;
	if v38 != 0 {
		v178 = v5
		goto L1
	} else {
		goto L7
	}
L3:
	;
	goto L2
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteSlotToAppendOnlyFileRio[0]))
	if v23 <= l1 {
		v38 = v19
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v25 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteSlotToAppendOnlyFileRio[1]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+l1<<(uint(int32(2))%32))))
	if v30 == v25 {
		v38 = v19
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v34 = F_kvstoreSize(m, v33)
	mBase = m.M
	v38 = base.B2i32(v34 == int64(0))
	goto L3
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_rewriteSlotToAppendOnlyFileRio[1]))
	v41 = int32(2)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+l1<<(uint(v41)%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+l2<<(uint(v41)%32))))
	if v50 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v53 == int32(0) {
		v178 = v5
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v52 = F_hashtableSize(m, v50)
	mBase = m.M
	v53 = v52
	goto L8
L10:
	;
	v53 = int32(0)
	goto L8
L11:
	;
	v56 = int32(16)
	v59 = int32(0)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_rewriteSlotToAppendOnlyFileRio[2])))
	*(*uint8)(unsafe.Add(mBase, uint32(v15+v56))) = uint8(v60)
	v63 = *(*int64)(unsafe.Add(mBase, _c_F_rewriteSlotToAppendOnlyFileRio[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v63
	v66 = *(*int64)(unsafe.Add(mBase, _c_F_rewriteSlotToAppendOnlyFileRio[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v66
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v68&int32(6) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v178 = int32(-1)
	goto L1
L13:
	;
	v78 = v56
	v79 = v15
	goto L14
L14:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v83) < base.Ui32(v78) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v107 = F_rioWriteBulkLongLong(m, l0, base.I64_extend_i32_s(l1))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L24
	} else {
		goto L31
	}
L16:
	;
	v85 = v83
	goto L18
L17:
	;
	v85 = v78
	goto L18
L18:
	;
	if v83 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v86 = v85
	goto L21
L20:
	;
	v86 = v78
	goto L21
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v87 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v95 = m.T0[v94].(func(*base.Module, int32, int32, int32) int32)(m, l0, v79, v86)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L24
	} else {
		goto L27
	}
L23:
	;
	m.T0[v87].(func(*base.Module, int32, int32, int32))(m, l0, v79, v86)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	goto L22
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v101 + v86
	v105 = v78 - v86
	if v105 != 0 {
		v78 = v105
		v79 = v79 + v86
		goto L14
	} else {
		goto L29
	}
L27:
	;
	if v95 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v97 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v97 | int64(2)
	goto L12
L29:
	;
	goto L15
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v112 = F_kvstoreGetHashtableIterator(m, v110, l2, int32(3))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L24
	} else {
		goto L33
	}
L31:
	;
	if v107 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v178 = int32(-1)
	goto L1
L33:
	;
	v125 = int64(0)
	goto L35
L34:
	;
	F_kvstoreReleaseHashtableIterator(m, v112)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L24
	} else {
		goto L46
	}
L35:
	;
	v127 = F_kvstoreHashtableIteratorNext(m, v112, v15)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L24
	} else {
		goto L37
	}
L37:
	;
	if v127 == int32(0) {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if l3 == int32(0) {
		v150 = v125
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v153 = F_rewriteObjectRio(m, l0, v131, l1)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L24
	} else {
		goto L44
	}
L40:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v134 + int32(1)
	if v134&int32(1023) != 0 {
		v150 = v125
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v140 = F_mstime(m)
	mBase = m.M
	if v140-v125 < int64(1000) {
		v150 = v125
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_sendChildInfo(m, int32(0), v145, int32(_a_F_rewriteSlotToAppendOnlyFileRio_0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L24
	} else {
		goto L43
	}
L43:
	;
	v150 = v140
	goto L39
L44:
	;
	if v153 != int32(-1) {
		v125 = v150
		goto L35
	} else {
		goto L45
	}
L45:
	;
	v178 = int32(-1)
	goto L1
L46:
	;
	v178 = int32(0)
	goto L1
}
func F_slotMigrationPipeReadHandler(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[0]))
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[1]))
	if v22 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v18 = F_valkey_malloc(m, int32(16384))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[0])) = v18
	goto L1
L5:
	;
	m.G0 = v12 + int32(32)
	return
L6:
	;
	goto L7
L7:
	;
	v34 = int32(_a_F_slotMigrationPipeReadHandler_0)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[0]))
	v38 = F_read(m, l1, v36, int32(16384))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[2])) = v38
	if int32(-1) < v38 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+80))
	v130 = m.T0[v129].(func(*base.Module, int32, int32, int32) int32)(m, v78, int32(972), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L3
	} else {
		goto L33
	}
L9:
	;
	if v38 != 0 {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	goto L11
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[3]))
	if v43 == int32(6) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[4]))
	if int32(3) < v47 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[1]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	v59 = F_freeClient(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L17
	}
L14:
	;
	v50 = F___strerror_l(m, v43, v43)
	mBase = m.M
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v50
	F__serverLog(m, int32(3), int32(_a_F_slotMigrationPipeReadHandler_1), v12)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[1])) = int32(0)
	goto L5
L18:
	;
	v77 = int32(_a_F_slotMigrationPipeReadHandler_0)
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[1]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+24))
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[0]))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+68))
	v84 = m.T0[v83].(func(*base.Module, int32, int32, int32) int32)(m, v78, v81, v38)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L3
	} else {
		goto L23
	}
L19:
	;
	v64 = int32(_a_F_slotMigrationPipeReadHandler_0)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[5]))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[6]))
	F_aeDeleteFileEvent(m, v65, v67, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v71 = int32(_a_F_slotMigrationPipeReadHandler_0)
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[7]))
	v73 = F_close(m, v72)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[7])) = int32(-1)
	goto L5
L21:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[2]))
	if v84 == v124 {
		goto L7
	} else {
		goto L32
	}
L22:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v79)+104))
	v115 = base.I64_extend_i32_s(v84)
	*(*int64)(unsafe.Add(mBase, uint32(v114)+16)) = v115
	v117 = int32(_a_F_slotMigrationPipeReadHandler_0)
	v119 = *(*int64)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[8]))
	*(*int64)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[8])) = v119 + v115
	goto L21
L23:
	;
	if v84 != int32(-1) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v88 == int32(3) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v79)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v111)+16)) = int64(0)
	goto L21
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[4]))
	if int32(3) < v92 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v106 = F_freeClient(m, v79)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L31
	}
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+88))
	v97 = m.T0[v96].(func(*base.Module, int32) int32)(m, v78)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v97
	F__serverLog(m, int32(3), int32(_a_F_slotMigrationPipeReadHandler_2), v12+int32(16))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[1])) = int32(0)
	goto L5
L32:
	;
	goto L8
L33:
	;
	v132 = int32(_a_F_slotMigrationPipeReadHandler_0)
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[5]))
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeReadHandler[6]))
	F_aeDeleteFileEvent(m, v133, v135, int32(1))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	goto L5
}
func F_slotMigrationPipeWriteHandler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeWriteHandler[0]))
	if v13 <= int32(0) {
		F__serverAssert(m, int32(_a_F_slotMigrationPipeWriteHandler_0), int32(_a_F_slotMigrationPipeWriteHandler_1), int32(1880))
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
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeWriteHandler[1]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
		v25 = m.T0[v24].(func(*base.Module, int32, int32, int32) int32)(m, l0, v17+v20, v13-v20)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			if v25 != int32(-1) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
				v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)+16))
				v49 = base.I64_extend_i32_s(v25)
				v50 = v48 + v49
				*(*int64)(unsafe.Add(mBase, uint32(v47)+16)) = v50
				v52 = int32(_a_F_slotMigrationPipeWriteHandler_2)
				v54 = *(*int64)(unsafe.Add(mBase, _c_F_slotMigrationPipeWriteHandler[2]))
				*(*int64)(unsafe.Add(mBase, _c_F_slotMigrationPipeWriteHandler[2])) = v54 + v49
				v58 = int64(*(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeWriteHandler[0])))
				if v58 <= v50 {
					v64 = int32(0)
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+80))
					v68 = m.T0[v67].(func(*base.Module, int32, int32, int32) int32)(m, l0, v64, v64)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
						*(*int64)(unsafe.Add(mBase, uint32(v70)+88)) = int64(0)
						v73 = int32(_a_F_slotMigrationPipeWriteHandler_2)
						v74 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeWriteHandler[3]))
						v76 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeWriteHandler[4]))
						v80 = F_aeCreateFileEvent(m, v74, v76, int32(1), int32(108), int32(0))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							if v80 == int32(-1) {
								F__serverPanic_1(m, int32(_a_F_slotMigrationPipeWriteHandler_1), int32(1902), int32(_a_F_slotMigrationPipeWriteHandler_3), int32(0))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						}
					}
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
					v62 = *(*int64)(unsafe.Add(mBase, _c_F_slotMigrationPipeWriteHandler[5]))
					*(*int64)(unsafe.Add(mBase, uint32(v60)+88)) = v62
					m.G0 = v10 + int32(16)
					return
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v29 == int32(3) {
					m.G0 = v10 + int32(16)
					return
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_slotMigrationPipeWriteHandler[6]))
					if int32(3) < v33 {
						v45 = F_freeClient(m, v18)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v10 + int32(16)
							return
						}
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
						v38 = m.T0[v37].(func(*base.Module, int32) int32)(m, l0)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v38
							F__serverLog(m, int32(3), int32(_a_F_slotMigrationPipeWriteHandler_4), v10)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v45 = F_freeClient(m, v18)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
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
		}
	}
}
