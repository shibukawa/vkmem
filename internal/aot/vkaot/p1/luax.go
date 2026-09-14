package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaX_init(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
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
	var v77 int32
	_ = v77
	v7 = int32(0)
	goto L1
L1:
	;
	v10 = m.G401
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10+v7<<(uint(int32(2))%32))))
	if v14&int32(3) == int32(0) {
		v36 = v14
		goto L5
	} else {
		goto L6
	}
L2:
	;
	return
L3:
	;
	v70 = F_luaS_newlstr(m, l0, v14, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v69 = v61 - v14
	goto L3
L5:
	;
	v40 = v36
	goto L13
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v25 = v14
	goto L9
L8:
	;
	v69 = v14 - v14
	goto L3
L9:
	;
	v29 = v25 + int32(1)
	if v29&int32(3) == int32(0) {
		v36 = v29
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v34 != 0 {
		v25 = v29
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v61 = v29
	goto L4
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v49 = int32(-2139062144)
	if (int32(16843008)-v46|v46)&v49 == v49 {
		v40 = v40 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v55 = v40
	goto L16
L15:
	;
	goto L14
L16:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v59 != 0 {
		v55 = v55 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v61 = v55
	goto L4
L18:
	;
	goto L17
L19:
	;
	return
L20:
	;
	v73 = v7 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+6)) = uint8(v73)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+5)))
	v77 = v75 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+5)) = uint8(v77)
	if v73 != int32(21) {
		v7 = v73
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L2
}
func F_luaX_lexerror(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
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
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	v7 = m.G0
	v9 = v7 - int32(144)
	m.G0 = v9
	v12 = v9 + int32(64)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v15 = v13 + int32(16)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	switch v19 + int32(-61) {
	case 0:
		v24 = F_strncpy(m, v12, v13+int32(17), int32(80))
		mBase = m.M
		v28 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(79)))) = uint8(v28)
	default:
		v46 = m.G3
		v49 = F_strcspn(m, v15, v46+int32(_a2653))
		mBase = m.M
		v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[988]))))
		*(*uint16)(unsafe.Add(mBase, uint32(v9+int32(72)))) = uint16(v56)
		v58 = *(*int64)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[989])))
		*(*int64)(unsafe.Add(mBase, uint32(v12))) = v58
		v61 = int32(63)
		if base.Ui32(v49) < base.Ui32(v61) {
			v63 = v49
		} else {
			v63 = v61
		}
		v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v63))))
		if v65 == int32(0) {
			v73 = F_strcat(m, v12, v15)
			mBase = m.M
		} else {
			v68 = F_strncat(m, v12, v15, v63)
			mBase = m.M
			v69 = F_strlen(m, v68)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v68+v69))) = int32(3026478)
		}
		v75 = F_strlen(m, v12)
		mBase = m.M
		v76 = v12 + v75
		v77 = m.G3
		v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+uint32(_consts[990]))))
		*(*uint16)(unsafe.Add(mBase, uint32(v76))) = uint16(v80)
		v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+uint32(_consts[991]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v76+int32(2)))) = uint8(v86)
	case 3:
		v31 = v13 + int32(17)
		v32 = F_strlen(m, v31)
		mBase = m.M
		v33 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v33)
		v36 = int32(72)
		if base.Ui32(v32) <= base.Ui32(v36) {
			v44 = v31
		} else {
			v38 = F_strlen(m, v12)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v12+v38))) = int32(3026478)
			v44 = v31 + (v32 - v36)
		}
		v45 = F_strcat(m, v12, v44)
		mBase = m.M
	}
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v94
	v97 = m.G3
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v9 + int32(64)
	v105 = F_luaO_pushfstring(m, v93, v97+int32(_a2654), v9+int32(48))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		return
	} else {
		if l2 == int32(0) {
			v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			F_luaD_throw(m, v161, int32(3))
			mBase = m.M
			v164 = m.ExcPending
			if v164 != 0 {
				return
			} else {
				m.G0 = v9 + int32(144)
				return
			}
		} else {
			v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			if base.Ui32(int32(2)) < base.Ui32(l2+int32(-284)) {
				if int32(256) < l2 {
					v146 = m.G401
					v150 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32)+v146+int32(-1028))))
					v151 = v150
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v151
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v105
					v154 = m.G3
					v157 = F_luaO_pushfstring(m, v109, v154+int32(_a2671), v9)
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
						return
					} else {
						v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						F_luaD_throw(m, v161, int32(3))
						mBase = m.M
						v164 = m.ExcPending
						if v164 != 0 {
							return
						} else {
							m.G0 = v9 + int32(144)
							return
						}
					}
				} else {
					if base.B2i32(base.Ui32(l2) < base.Ui32(int32(32)))|base.B2i32(l2 == int32(127)) == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l2
						v137 = m.G3
						v142 = F_luaO_pushfstring(m, v109, v137+int32(_a2669), v9+int32(16))
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return
						} else {
							v151 = v142
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v151
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v105
							v154 = m.G3
							v157 = F_luaO_pushfstring(m, v109, v154+int32(_a2671), v9)
							mBase = m.M
							v158 = m.ExcPending
							if v158 != 0 {
								return
							} else {
								v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								F_luaD_throw(m, v161, int32(3))
								mBase = m.M
								v164 = m.ExcPending
								if v164 != 0 {
									return
								} else {
									m.G0 = v9 + int32(144)
									return
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l2
						v129 = m.G3
						v134 = F_luaO_pushfstring(m, v109, v129+int32(_a2670), v9+int32(32))
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return
						} else {
							v151 = v134
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v151
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v105
							v154 = m.G3
							v157 = F_luaO_pushfstring(m, v109, v154+int32(_a2671), v9)
							mBase = m.M
							v158 = m.ExcPending
							if v158 != 0 {
								return
							} else {
								v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								F_luaD_throw(m, v161, int32(3))
								mBase = m.M
								v164 = m.ExcPending
								if v164 != 0 {
									return
								} else {
									m.G0 = v9 + int32(144)
									return
								}
							}
						}
					}
				}
			} else {
				F_save(m, l0, int32(0))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return
				} else {
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
					v151 = v118
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v151
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v105
					v154 = m.G3
					v157 = F_luaO_pushfstring(m, v109, v154+int32(_a2671), v9)
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
						return
					} else {
						v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						F_luaD_throw(m, v161, int32(3))
						mBase = m.M
						v164 = m.ExcPending
						if v164 != 0 {
							return
						} else {
							m.G0 = v9 + int32(144)
							return
						}
					}
				}
			}
		}
	}
}
func F_luaX_newstring(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v6 = F_luaS_newlstr(m, v5, l1, l2)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v12 = F_luaH_setstr(m, v5, v11, v6)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if v14 != 0 {
				return v6
			} else {
				v15 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v15
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v15
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
				if base.Ui32(v20) < base.Ui32(v21) {
					return v6
				} else {
					F_luaC_step(m, v5)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						return v6
					}
				}
			}
		}
	}
}
func F_luaX_token2str(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	if int32(256) < l1 {
		v36 = m.G401
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32)+v36+int32(-1028))))
		v42 = v40
		m.G0 = v6 + int32(32)
		return v42
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		if base.B2i32(base.Ui32(l1) < base.Ui32(int32(32)))|base.B2i32(l1 == int32(127)) == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
			v29 = m.G3
			v32 = F_luaO_pushfstring(m, v10, v29+int32(_a2669), v6)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v42 = v32
				m.G0 = v6 + int32(32)
				return v42
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l1
			v19 = m.G3
			v24 = F_luaO_pushfstring(m, v10, v19+int32(_a2670), v6+int32(16))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v42 = v24
				m.G0 = v6 + int32(32)
				return v42
			}
		}
	}
}
