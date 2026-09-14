package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_jumponcond(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v7 + int32(-11) {
	case 0:
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v11+v12<<(uint(int32(2))%32))))
		if v16&int32(63) != int32(19) {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v39 = v37 + int32(1)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+75)))
			if v41 <= v37 {
				if base.Ui32(v39) < base.Ui32(int32(250)) {
					v55 = v39
					v56 = v40
					*(*uint8)(unsafe.Add(mBase, uint32(v56)+75)) = uint8(v39)
					v58 = v55
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v58
					F_discharge2reg(m, l0, l1, v58+int32(-1))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if v65 != int32(12) {
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							if v71&int32(256) != 0 {
							} else {
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
								if v71 < v74 {
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v76 + int32(-1)
								}
							}
						}
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v86 = F_condjump(m, l0, int32(27), int32(255), v85, l2)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							return v86
						}
					}
				} else {
					v45 = m.G3
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					F_luaX_syntaxerror(m, v46, v45+int32(_a2263))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v55 = v51 + int32(1)
						v56 = v54
						*(*uint8)(unsafe.Add(mBase, uint32(v56)+75)) = uint8(v39)
						v58 = v55
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v58
						F_discharge2reg(m, l0, l1, v58+int32(-1))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if v65 != int32(12) {
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								if v71&int32(256) != 0 {
								} else {
									v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
									if v71 < v74 {
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v76 + int32(-1)
									}
								}
							}
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v86 = F_condjump(m, l0, int32(27), int32(255), v85, l2)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								return v86
							}
						}
					}
				}
			} else {
				v58 = v39
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v58
				F_discharge2reg(m, l0, l1, v58+int32(-1))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v65 != int32(12) {
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						if v71&int32(256) != 0 {
						} else {
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
							if v71 < v74 {
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v76 + int32(-1)
							}
						}
					}
					v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v86 = F_condjump(m, l0, int32(27), int32(255), v85, l2)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						return v86
					}
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v21 + int32(-1)
			v31 = F_condjump(m, l0, int32(26), int32(base.Ui32(v16)>>(uint(int32(23))%32)), int32(0), l2^int32(1))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				return v31
			}
		}
	case 1:
		v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v71&int32(256) != 0 {
		} else {
			v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
			if v71 < v74 {
			} else {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v76 + int32(-1)
			}
		}
		v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v86 = F_condjump(m, l0, int32(27), int32(255), v85, l2)
		mBase = m.M
		v87 = m.ExcPending
		if v87 != 0 {
			return int32(0)
		} else {
			return v86
		}
	default:
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v39 = v37 + int32(1)
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+75)))
		if v41 <= v37 {
			if base.Ui32(v39) < base.Ui32(int32(250)) {
				v55 = v39
				v56 = v40
				*(*uint8)(unsafe.Add(mBase, uint32(v56)+75)) = uint8(v39)
				v58 = v55
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v58
				F_discharge2reg(m, l0, l1, v58+int32(-1))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v65 != int32(12) {
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						if v71&int32(256) != 0 {
						} else {
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
							if v71 < v74 {
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v76 + int32(-1)
							}
						}
					}
					v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v86 = F_condjump(m, l0, int32(27), int32(255), v85, l2)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						return v86
					}
				}
			} else {
				v45 = m.G3
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_luaX_syntaxerror(m, v46, v45+int32(_a2263))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v55 = v51 + int32(1)
					v56 = v54
					*(*uint8)(unsafe.Add(mBase, uint32(v56)+75)) = uint8(v39)
					v58 = v55
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v58
					F_discharge2reg(m, l0, l1, v58+int32(-1))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if v65 != int32(12) {
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							if v71&int32(256) != 0 {
							} else {
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
								if v71 < v74 {
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v76 + int32(-1)
								}
							}
						}
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v86 = F_condjump(m, l0, int32(27), int32(255), v85, l2)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							return v86
						}
					}
				}
			}
		} else {
			v58 = v39
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v58
			F_discharge2reg(m, l0, l1, v58+int32(-1))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v65 != int32(12) {
				} else {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if v71&int32(256) != 0 {
					} else {
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
						if v71 < v74 {
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v76 + int32(-1)
						}
					}
				}
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v86 = F_condjump(m, l0, int32(27), int32(255), v85, l2)
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					return v86
				}
			}
		}
	}
}
