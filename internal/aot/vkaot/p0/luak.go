package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaK_code(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_patchlistaux(m, l0, v8, v9, int32(255), v9)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(-1)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
		if v18 <= v17 {
			v21 = m.G3
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v30 = F_luaM_growaux_(m, v22, v23, v7+int32(44), int32(4), int32(2147483645), v21+int32(_a_F_luaK_code_0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v30
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v34 = v33
				v35 = v30
				*(*int32)(unsafe.Add(mBase, uint32(v35+v34<<(uint(int32(2))%32)))) = l1
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
				if v41 <= v40 {
					v44 = m.G3
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
					v53 = F_luaM_growaux_(m, v45, v46, v7+int32(48), int32(4), int32(2147483645), v44+int32(_a_F_luaK_code_0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v53
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v57 = v53
						v58 = v56
						*(*int32)(unsafe.Add(mBase, uint32(v57+v58<<(uint(int32(2))%32)))) = l2
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v63 + int32(1)
						return v63
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
					v57 = v43
					v58 = v40
					*(*int32)(unsafe.Add(mBase, uint32(v57+v58<<(uint(int32(2))%32)))) = l2
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v63 + int32(1)
					return v63
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v34 = v17
			v35 = v20
			*(*int32)(unsafe.Add(mBase, uint32(v35+v34<<(uint(int32(2))%32)))) = l1
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
			if v41 <= v40 {
				v44 = m.G3
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
				v53 = F_luaM_growaux_(m, v45, v46, v7+int32(48), int32(4), int32(2147483645), v44+int32(_a_F_luaK_code_0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v53
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v57 = v53
					v58 = v56
					*(*int32)(unsafe.Add(mBase, uint32(v57+v58<<(uint(int32(2))%32)))) = l2
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v63 + int32(1)
					return v63
				}
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
				v57 = v43
				v58 = v40
				*(*int32)(unsafe.Add(mBase, uint32(v57+v58<<(uint(int32(2))%32)))) = l2
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v63 + int32(1)
				return v63
			}
		}
	}
}
func F_luaK_codeABC(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v17 = F_luaK_code(m, l0, l2<<(uint(int32(6))%32)|l1|l3<<(uint(int32(23))%32)|l4<<(uint(int32(14))%32), v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		return v17
	}
}
func F_luaK_codeABx(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v13 = F_luaK_code(m, l0, l2<<(uint(int32(6))%32)|l1|l3<<(uint(int32(14))%32), v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		return v13
	}
}
func F_luaK_exp2nextreg(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	F_luaK_dischargevars(m, l0, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v8 != int32(12) {
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v11&int32(256) != 0 {
			} else {
				v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
				if v11 < v14 {
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v16 + int32(-1)
				}
			}
		}
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v23 = v21 + int32(1)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+75)))
		if v25 <= v21 {
			if base.Ui32(v23) < base.Ui32(int32(250)) {
				v39 = v23
				v40 = v24
				*(*uint8)(unsafe.Add(mBase, uint32(v40)+75)) = uint8(v23)
				v42 = v39
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v42
				F_exp2reg(m, l0, l1, v42+int32(-1))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					return
				}
			} else {
				v29 = m.G3
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_luaX_syntaxerror(m, v30, v29+int32(_a_F_luaK_exp2nextreg_0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v39 = v35 + int32(1)
					v40 = v38
					*(*uint8)(unsafe.Add(mBase, uint32(v40)+75)) = uint8(v23)
					v42 = v39
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v42
					F_exp2reg(m, l0, l1, v42+int32(-1))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v42 = v23
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v42
			F_exp2reg(m, l0, l1, v42+int32(-1))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_luaK_getlabel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v3
	return v3
}
func F_luaK_nil(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v7 <= v8 {
		v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
		v63 = F_luaK_code(m, l0, l1<<(uint(int32(6))%32)|((l2+l1)<<(uint(int32(23))%32)+int32(-8388608))|int32(3), v62)
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return
		} else {
			return
		}
	} else {
		if v7 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
			v18 = v13 + v7<<(uint(int32(2))%32) + int32(-4)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			if v19&int32(63) != int32(3) {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
				v63 = F_luaK_code(m, l0, l1<<(uint(int32(6))%32)|((l2+l1)<<(uint(int32(23))%32)+int32(-8388608))|int32(3), v62)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					return
				}
			} else {
				if l1 < int32(base.Ui32(v19)>>(uint(int32(6))%32))&int32(255) {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
					v63 = F_luaK_code(m, l0, l1<<(uint(int32(6))%32)|((l2+l1)<<(uint(int32(23))%32)+int32(-8388608))|int32(3), v62)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						return
					}
				} else {
					v30 = int32(base.Ui32(v19) >> (uint(int32(23)) % 32))
					if v30+int32(1) < l1 {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
						v63 = F_luaK_code(m, l0, l1<<(uint(int32(6))%32)|((l2+l1)<<(uint(int32(23))%32)+int32(-8388608))|int32(3), v62)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							return
						}
					} else {
						v36 = l2 + l1 + int32(-1)
						if v36 <= v30 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v18))) = v19&int32(8388547) | v36<<(uint(int32(23))%32)
						}
						return
					}
				}
			}
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
			if v10 <= l1 {
				return
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
				v63 = F_luaK_code(m, l0, l1<<(uint(int32(6))%32)|((l2+l1)<<(uint(int32(23))%32)+int32(-8388608))|int32(3), v62)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_luaK_patchlist(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if l2 != v8 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	return
L2:
	;
	return
L3:
	;
	F_patchlistaux(m, l0, l1, l2, int32(255), l2)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L14
	} else {
		goto L16
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l2
	if l1 == int32(-1) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v13 == int32(-1) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v20 = v13
	goto L8
L7:
	;
	v43 = l1 + (v20 ^ int32(-1))
	v45 = v43 >> (uint(int32(31)) % 32)
	if base.Ui32(v43^v45-v45) < base.Ui32(int32(131072)) {
		v58 = v28
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v27 = v17 + v20<<(uint(int32(2))%32)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v32 = int32(base.Ui32(v28)>>(uint(int32(14))%32)) + int32(-131071)
	if v32 == int32(-1) {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v37 = v20 + v32 + int32(1)
	if v37 != int32(-1) {
		v20 = v37
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v43<<(uint(int32(14))%32) | v58&int32(16383) + int32(2147467264)
	return
L13:
	;
	v50 = m.G3
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v51, v50+int32(_a_F_luaK_patchlist_0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v58 = v56
	goto L12
L16:
	;
	goto L2
}
func F_luaK_patchtohere(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v8
	if l1 == int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	return
L2:
	;
	return
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v12 == int32(-1) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v19 = v12
	goto L6
L5:
	;
	v42 = l1 + (v19 ^ int32(-1))
	v44 = v42 >> (uint(int32(31)) % 32)
	if base.Ui32(v42^v44-v44) < base.Ui32(int32(131072)) {
		v57 = v27
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v26 = v16 + v19<<(uint(int32(2))%32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v31 = int32(base.Ui32(v27)>>(uint(int32(14))%32)) + int32(-131071)
	if v31 == int32(-1) {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v36 = v19 + v31 + int32(1)
	if v36 != int32(-1) {
		v19 = v36
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v42<<(uint(int32(14))%32) | v57&int32(16383) + int32(2147467264)
	goto L2
L11:
	;
	v49 = m.G3
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v50, v49+int32(_a_F_luaK_patchtohere_0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v57 = v55
	goto L10
}
func F_luaK_posfix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
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
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v462 int64
	_ = v462
	var v464 int32
	_ = v464
	var v468 int64
	_ = v468
	var v470 int32
	_ = v470
	var v474 int64
	_ = v474
	var v485 int64
	_ = v485
	var v487 int32
	_ = v487
	var v491 int64
	_ = v491
	var v493 int32
	_ = v493
	var v497 int64
	_ = v497
	switch l1 {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L13
	case 5:
		goto L12
	case 6:
		goto L18
	case 7:
		goto L10
	case 8:
		goto L11
	case 9:
		goto L9
	case 10:
		goto L8
	case 11:
		goto L7
	case 12:
		goto L6
	case 13:
		goto L20
	case 14:
		goto L19
	default:
		goto L5
	}
L1:
	;
	v485 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v485
	v487 = int32(16)
	v491 = *(*int64)(unsafe.Add(mBase, uint32(l3+v487)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v487))) = v491
	v493 = int32(8)
	v497 = *(*int64)(unsafe.Add(mBase, uint32(l3+v493)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v493))) = v497
	return
L2:
	;
	v462 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v462
	v464 = int32(16)
	v468 = *(*int64)(unsafe.Add(mBase, uint32(l3+v464)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v464))) = v468
	v470 = int32(8)
	v474 = *(*int64)(unsafe.Add(mBase, uint32(l3+v470)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v470))) = v474
	return
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v73
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v12
	goto L1
L5:
	;
	return
L6:
	;
	v411 = F_luaK_exp2RK(m, l0, l2)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L21
	} else {
		goto L123
	}
L7:
	;
	v374 = F_luaK_exp2RK(m, l0, l2)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L21
	} else {
		goto L112
	}
L8:
	;
	v337 = F_luaK_exp2RK(m, l0, l2)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L21
	} else {
		goto L101
	}
L9:
	;
	v300 = F_luaK_exp2RK(m, l0, l2)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L21
	} else {
		goto L90
	}
L10:
	;
	v263 = F_luaK_exp2RK(m, l0, l2)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L21
	} else {
		goto L79
	}
L11:
	;
	v226 = F_luaK_exp2RK(m, l0, l2)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L21
	} else {
		goto L68
	}
L12:
	;
	F_codearith(m, l0, int32(17), l2, l3)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L21
	} else {
		goto L67
	}
L13:
	;
	F_codearith(m, l0, int32(16), l2, l3)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L21
	} else {
		goto L66
	}
L14:
	;
	F_codearith(m, l0, int32(15), l2, l3)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L21
	} else {
		goto L65
	}
L15:
	;
	F_codearith(m, l0, int32(14), l2, l3)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L21
	} else {
		goto L64
	}
L16:
	;
	F_codearith(m, l0, int32(13), l2, l3)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L21
	} else {
		goto L63
	}
L17:
	;
	F_codearith(m, l0, int32(12), l2, l3)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L21
	} else {
		goto L62
	}
L18:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	F_luaK_dischargevars(m, l0, l3)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L21
	} else {
		goto L44
	}
L19:
	;
	F_luaK_dischargevars(m, l0, l3)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L21
	} else {
		goto L33
	}
L20:
	;
	F_luaK_dischargevars(m, l0, l3)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return
L22:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v12 == int32(-1) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v15 == int32(-1) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v21 = v15
	goto L26
L25:
	;
	v47 = v12 + (v21 ^ int32(-1))
	v49 = v47 >> (uint(int32(31)) % 32)
	if base.Ui32(v47^v49-v49) < base.Ui32(int32(131072)) {
		v62 = v32
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v31 = v19 + v21<<(uint(int32(2))%32)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v36 = int32(base.Ui32(v32)>>(uint(int32(14))%32)) + int32(-131071)
	if v36 == int32(-1) {
		goto L25
	} else {
		goto L28
	}
L27:
	;
	goto L25
L28:
	;
	v41 = v21 + v36 + int32(1)
	if v41 != int32(-1) {
		v21 = v41
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v47<<(uint(int32(14))%32) | v62&int32(16383) + int32(2147467264)
	goto L1
L31:
	;
	v54 = m.G3
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v55, v54+int32(_a_F_luaK_posfix_0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L21
	} else {
		goto L32
	}
L32:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v62 = v60
	goto L30
L33:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v73 == int32(-1) {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v76 == int32(-1) {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v82 = v76
	goto L37
L36:
	;
	v108 = v73 + (v82 ^ int32(-1))
	v110 = v108 >> (uint(int32(31)) % 32)
	if base.Ui32(v108^v110-v110) < base.Ui32(int32(131072)) {
		v123 = v93
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v92 = v80 + v82<<(uint(int32(2))%32)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v97 = int32(base.Ui32(v93)>>(uint(int32(14))%32)) + int32(-131071)
	if v97 == int32(-1) {
		goto L36
	} else {
		goto L39
	}
L38:
	;
	goto L36
L39:
	;
	v102 = v82 + v97 + int32(1)
	if v102 != int32(-1) {
		v82 = v102
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v108<<(uint(int32(14))%32) | v123&int32(16383) + int32(2147467264)
	goto L2
L42:
	;
	v115 = m.G3
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v116, v115+int32(_a_F_luaK_posfix_0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L21
	} else {
		goto L43
	}
L43:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v123 = v121
	goto L41
L44:
	;
	if v133 == v132 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v152 != int32(11) {
		goto L53
	} else {
		goto L54
	}
L46:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v137 != int32(12) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_luaK_exp2nextreg(m, l0, l3)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L21
	} else {
		goto L52
	}
L48:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v140 == v141 {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v143 < v144 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	F_exp2reg(m, l0, l3, v143)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L21
	} else {
		goto L51
	}
L51:
	;
	goto L45
L52:
	;
	goto L45
L53:
	;
	F_luaK_exp2nextreg(m, l0, l3)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L21
	} else {
		goto L60
	}
L54:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v156+v157<<(uint(int32(2))%32))))
	if v161&int32(63) != int32(21) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v166 != int32(12) {
		v183 = v157
		v184 = v161
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v156+v183<<(uint(int32(2))%32)))) = v189<<(uint(int32(23))%32) | v184&int32(8388607)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v198
	return
L57:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v169&int32(256) != 0 {
		v183 = v157
		v184 = v161
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v169 < v172 {
		v183 = v157
		v184 = v161
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v174 + int32(-1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v156+v178<<(uint(int32(2))%32))))
	v183 = v178
	v184 = v182
	goto L56
L60:
	;
	F_codearith(m, l0, int32(21), l2, l3)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L21
	} else {
		goto L61
	}
L61:
	;
	return
L62:
	;
	return
L63:
	;
	return
L64:
	;
	return
L65:
	;
	return
L66:
	;
	return
L67:
	;
	return
L68:
	;
	v228 = F_luaK_exp2RK(m, l0, l3)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L21
	} else {
		goto L69
	}
L69:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v230 != int32(12) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v243 != int32(12) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v233&int32(256) != 0 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v233 < v236 {
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v238 + int32(-1)
	goto L70
L74:
	;
	v258 = F_condjump(m, l0, int32(23), int32(1), v226, v228)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L21
	} else {
		goto L78
	}
L75:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v246&int32(256) != 0 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v246 < v249 {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v251 + int32(-1)
	goto L74
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v258
	return
L79:
	;
	v265 = F_luaK_exp2RK(m, l0, l3)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L21
	} else {
		goto L80
	}
L80:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v267 != int32(12) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v280 != int32(12) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v270&int32(256) != 0 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v270 < v273 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v275 + int32(-1)
	goto L81
L85:
	;
	v295 = F_condjump(m, l0, int32(23), int32(0), v263, v265)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L21
	} else {
		goto L89
	}
L86:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v283&int32(256) != 0 {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v283 < v286 {
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v288 + int32(-1)
	goto L85
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v295
	return
L90:
	;
	v302 = F_luaK_exp2RK(m, l0, l3)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L21
	} else {
		goto L91
	}
L91:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v304 != int32(12) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v317 != int32(12) {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v307&int32(256) != 0 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v307 < v310 {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v312 + int32(-1)
	goto L92
L96:
	;
	v332 = F_condjump(m, l0, int32(24), int32(1), v300, v302)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L21
	} else {
		goto L100
	}
L97:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v320&int32(256) != 0 {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v320 < v323 {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v325 + int32(-1)
	goto L96
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v332
	return
L101:
	;
	v339 = F_luaK_exp2RK(m, l0, l3)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L21
	} else {
		goto L102
	}
L102:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v341 != int32(12) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v354 != int32(12) {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v344&int32(256) != 0 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v344 < v347 {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v349 + int32(-1)
	goto L103
L107:
	;
	v369 = F_condjump(m, l0, int32(25), int32(1), v337, v339)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L21
	} else {
		goto L111
	}
L108:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v357&int32(256) != 0 {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v357 < v360 {
		goto L107
	} else {
		goto L110
	}
L110:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v362 + int32(-1)
	goto L107
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v369
	return
L112:
	;
	v376 = F_luaK_exp2RK(m, l0, l3)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L21
	} else {
		goto L113
	}
L113:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v378 != int32(12) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v391 != int32(12) {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v381&int32(256) != 0 {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v381 < v384 {
		goto L114
	} else {
		goto L117
	}
L117:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v386 + int32(-1)
	goto L114
L118:
	;
	v406 = F_condjump(m, l0, int32(24), int32(1), v376, v374)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L21
	} else {
		goto L122
	}
L119:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v394&int32(256) != 0 {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v394 < v397 {
		goto L118
	} else {
		goto L121
	}
L121:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v399 + int32(-1)
	goto L118
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v406
	return
L123:
	;
	v413 = F_luaK_exp2RK(m, l0, l3)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L21
	} else {
		goto L124
	}
L124:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v415 != int32(12) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v428 != int32(12) {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v418&int32(256) != 0 {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v418 < v421 {
		goto L125
	} else {
		goto L128
	}
L128:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v423 + int32(-1)
	goto L125
L129:
	;
	v443 = F_condjump(m, l0, int32(25), int32(1), v413, v411)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L21
	} else {
		goto L133
	}
L130:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v431&int32(256) != 0 {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v431 < v434 {
		goto L129
	} else {
		goto L132
	}
L132:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v436 + int32(-1)
	goto L129
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v443
	goto L5
}
func F_luaK_prefix(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(5)
	switch l1 {
	case 0:
		goto L5
	case 1:
		goto L4
	case 2:
		goto L3
	default:
		goto L1
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return
L2:
	;
	F_codearith(m, l0, v296, l2, v11+int32(8))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L10
	} else {
		goto L72
	}
L3:
	;
	F_luaK_dischargevars(m, l0, l2)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L10
	} else {
		goto L65
	}
L4:
	;
	F_luaK_dischargevars(m, l0, l2)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L18
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v19 != int32(5) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_luaK_dischargevars(m, l0, l2)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v22 != int32(-1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v25 != int32(-1) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v296 = int32(18)
	goto L2
L10:
	;
	return
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v31 != int32(12) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_luaK_exp2nextreg(m, l0, l2)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L17
	}
L13:
	;
	v34 = int32(18)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v35 == v36 {
		v296 = v34
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v38 < v39 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	F_exp2reg(m, l0, l2, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v296 = v34
	goto L2
L17:
	;
	v296 = int32(18)
	goto L2
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v50 + int32(-1) {
	case 0, 2:
		goto L25
	case 1, 3, 4:
		goto L24
	default:
		goto L19
	case 9:
		goto L23
	case 10:
		goto L22
	case 11:
		goto L21
	}
L19:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v151
	if v152 == int32(-1) {
		v217 = v151
		goto L41
	} else {
		goto L42
	}
L20:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	v143 = F_luaK_code(m, l0, v136<<(uint(int32(23))%32)|int32(19), v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L10
	} else {
		goto L40
	}
L21:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v124&int32(256) != 0 {
		goto L20
	} else {
		goto L38
	}
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v92 = v90 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+75)))
	if v94 <= v90 {
		goto L31
	} else {
		goto L32
	}
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v62 = v58 + v59<<(uint(int32(2))%32)
	if v59 < int32(1) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(3)
	goto L19
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(2)
	goto L19
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = base.B2i32(v78&int32(16320) == int32(0))<<(uint(int32(6))%32) | v78&int32(-16321)
	goto L19
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v78 = v77
	v79 = v62
	goto L26
L28:
	;
	v66 = v62 + int32(-4)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = m.G400
	v72 = int32(*(*int8)(unsafe.Add(mBase, uint32(v68+v67&int32(63)))))
	if v72 < int32(0) {
		v78 = v67
		v79 = v66
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v111
	F_discharge2reg(m, l0, l2, v111+int32(-1))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L10
	} else {
		goto L36
	}
L31:
	;
	if base.Ui32(v92) < base.Ui32(int32(250)) {
		v108 = v92
		v109 = v93
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v111 = v92
	goto L30
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+75)) = uint8(v92)
	v111 = v108
	goto L30
L34:
	;
	v98 = m.G3
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v99, v98+int32(_a_F_luaK_prefix_0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v108 = v104 + int32(1)
	v109 = v107
	goto L33
L36:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v118 != int32(12) {
		goto L20
	} else {
		goto L37
	}
L37:
	;
	goto L21
L38:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v124 < v127 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v129 + int32(-1)
	goto L20
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v143
	goto L19
L41:
	;
	if v217 == int32(-1) {
		goto L1
	} else {
		goto L54
	}
L42:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v164 = v152
	goto L43
L43:
	;
	v169 = v158 + v164<<(uint(int32(2))%32)
	if v164 < int32(1) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v217 = v212
	goto L41
L45:
	;
	if v185&int32(63) != int32(27) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v185 = v184
	v186 = v169
	goto L45
L47:
	;
	v173 = v169 + int32(-4)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v175 = m.G400
	v179 = int32(*(*int8)(unsafe.Add(mBase, uint32(v175+v174&int32(63)))))
	if v179 < int32(0) {
		v185 = v174
		v186 = v173
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v203 = int32(base.Ui32(v199)>>(uint(int32(14))%32)) + int32(-131071)
	if v203 == int32(-1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = int32(base.Ui32(v185)>>(uint(int32(17))%32))&int32(32704) | v185&int32(8372250)
	goto L49
L51:
	;
	goto L44
L52:
	;
	v208 = v164 + v203 + int32(1)
	if v208 != int32(-1) {
		v164 = v208
		goto L43
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v229 = v217
	goto L55
L55:
	;
	v235 = v224 + v229<<(uint(int32(2))%32)
	if v229 < int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if v251&int32(63) != int32(27) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	v251 = v250
	v252 = v235
	goto L57
L59:
	;
	v239 = v235 + int32(-4)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	v241 = m.G400
	v245 = int32(*(*int8)(unsafe.Add(mBase, uint32(v241+v240&int32(63)))))
	if v245 < int32(0) {
		v251 = v240
		v252 = v239
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	v269 = int32(base.Ui32(v265)>>(uint(int32(14))%32)) + int32(-131071)
	if v269 == int32(-1) {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = int32(base.Ui32(v251)>>(uint(int32(17))%32))&int32(32704) | v251&int32(8372250)
	goto L61
L63:
	;
	v274 = v229 + v269 + int32(1)
	if v274 != int32(-1) {
		v229 = v274
		goto L55
	} else {
		goto L64
	}
L64:
	;
	goto L1
L65:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v279 != int32(12) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	F_luaK_exp2nextreg(m, l0, l2)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L10
	} else {
		goto L71
	}
L67:
	;
	v282 = int32(20)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v283 == v284 {
		v296 = v282
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v286 < v287 {
		goto L66
	} else {
		goto L69
	}
L69:
	;
	F_exp2reg(m, l0, l2, v286)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L10
	} else {
		goto L70
	}
L70:
	;
	v296 = v282
	goto L2
L71:
	;
	v296 = int32(20)
	goto L2
L72:
	;
	goto L1
}
func F_luaK_ret(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v15 = F_luaK_code(m, l0, l1<<(uint(int32(6))%32)|(l2<<(uint(int32(23))%32)+int32(_a_F_luaK_ret_0))|int32(30), v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		return
	}
}
func F_luaK_setlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	if l3 == int32(-1) {
		v12 = int32(0)
	} else {
		v12 = l3 << (uint(int32(23)) % 32)
	}
	v14 = l1 << (uint(int32(6)) % 32)
	v18 = base.I32_div_s(l2+int32(-1), int32(50))
	v20 = v18 + int32(1)
	if int32(25550) < l2 {
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
		v34 = F_luaK_code(m, l0, v14|v12|int32(34), v33)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			v36 = v20
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
			v39 = F_luaK_code(m, l0, v36, v38)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l1 + int32(1)
				return
			}
		}
	} else {
		v36 = v12 | v20<<(uint(int32(14))%32) | v14 | int32(34)
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
		v39 = F_luaK_code(m, l0, v36, v38)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l1 + int32(1)
			return
		}
	}
}
func F_luaK_setreturns(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v7 + int32(-13) {
	case 0:
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v15 = v11 + v12<<(uint(int32(2))%32)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = v16&int32(-8372225) | (l2<<(uint(int32(14))%32)+int32(16384))&int32(8372224)
		return
	case 1:
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v30 = int32(2)
		v32 = v28 + v29<<(uint(v30)%32)
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
		*(*int32)(unsafe.Add(mBase, uint32(v32))) = l2<<(uint(int32(23))%32) | v35&int32(8388607) + int32(_a_F_luaK_setreturns_0)
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v45 = v28 + v42<<(uint(v30)%32)
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
		*(*int32)(unsafe.Add(mBase, uint32(v45))) = v46<<(uint(int32(6))%32)&int32(16320) | v51&int32(-16321)
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v58 = v56 + int32(1)
		v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+75)))
		if v59 <= v56 {
			if base.Ui32(v58) < base.Ui32(int32(250)) {
				v73 = v58
				v74 = v27
				*(*uint8)(unsafe.Add(mBase, uint32(v74)+75)) = uint8(v58)
				v76 = v73
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v76
				return
			} else {
				v63 = m.G3
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_luaX_syntaxerror(m, v64, v63+int32(_a_F_luaK_setreturns_1))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v73 = v69 + int32(1)
					v74 = v72
					*(*uint8)(unsafe.Add(mBase, uint32(v74)+75)) = uint8(v58)
					v76 = v73
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v76
					return
				}
			}
		} else {
			v76 = v58
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v76
			return
		}
	default:
		return
	}
}
