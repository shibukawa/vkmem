package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_aeCreateTimeEvent(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v33 int32
	_ = v33
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v9 + int64(1)
	v14 = F_valkey_malloc(m, int32(40))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		if v14 != 0 {
			*(*int64)(unsafe.Add(mBase, uint32(v14))) = v9
			v21 = int32(0)
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_aeCreateTimeEvent[0]))
			v23 = m.T0[v22].(func(*base.Module) int64)(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l2
			*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v23 + l1*int64(1000)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v33
			if v33 == v21 {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = v14
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v14
			return v9
		} else {
			return int64(-1)
		}
	}
}
func F_aeDeleteEventLoop(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_valkey_free(m, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_valkey_free(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_valkey_free(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v14 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	v18 = v14
	goto L7
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if v22 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	F_valkey_free(m, v18)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	m.T0[v22].(func(*base.Module, int32, int32))(m, l0, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	if v21 != 0 {
		v18 = v21
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L8
L14:
	;
	return
}
func F_aeDeleteFileEvent(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v108 int32
	_ = v108
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v8&int32(32) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 <= l1 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	goto L6
L6:
	;
	goto L3
L7:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v108&int32(32) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v21 = v18 + l1<<(uint(int32(4))%32)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v22 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v30 = v22 & (l2<<(uint(int32(1))%32)&int32(4) | l2)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v30 ^ v22
	if v30 != v22 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if v30&int32(3) == int32(0) {
		goto L7
	} else {
		goto L19
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 != v34 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v43 = l1
	goto L14
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v58
	goto L10
L14:
	;
	if int32(1) <= v43 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v58 = v51
	goto L13
L16:
	;
	v51 = v43 + int32(-1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v18+v51<<(uint(int32(4))%32))))
	if v55 == int32(0) {
		v43 = v51
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v58 = l1>>(uint(int32(31))%32)&l1 + int32(-1)
	goto L13
L18:
	;
	goto L15
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v30&int32(1) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v30&int32(2) == int32(0) {
		goto L7
	} else {
		goto L22
	}
L21:
	;
	v80 = v71 + int32(base.Ui32(l1)>>(uint(int32(3))%32))&int32(536870908)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v81 & base.I32_rotl(int32(-2), l1)
	goto L20
L22:
	;
	v95 = v71 + int32(base.Ui32(l1)>>(uint(int32(3))%32))&int32(536870908)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+128)) = v96 & base.I32_rotl(int32(-2), l1)
	goto L7
L23:
	;
	return
L24:
	;
	goto L25
L25:
	;
	goto L26
L26:
	;
	goto L23
}
func F_aeGetApiName(m *base.Module) int32 {
	return int32(_a_F_aeGetApiName_0)
}
func F_aeGetFileEvents(m *base.Module, l0 int32, l1 int32) int32 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 <= l1 {
		v12 = int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7+l1<<(uint(int32(4))%32))))
		v12 = v11
	}
	return v12
}
func F_aeGetSetSize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	return v2
}
func F_aeResizeSetSize(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v14&int32(32) == int32(0) {
	} else {
	}
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l1 == v23 {
		v160 = int32(0)
		v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
		if v171&int32(32) == int32(0) {
		} else {
		}
		return v160
	} else {
		v25 = int32(-1)
		if int32(1023) < l1 {
			v160 = v25
			v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
			if v171&int32(32) == int32(0) {
			} else {
			}
			return v160
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if l1 <= v28 {
				v160 = v25
				v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
				if v171&int32(32) == int32(0) {
				} else {
				}
				return v160
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v33 = F_valkey_realloc(m, v30, l1<<(uint(int32(4))%32))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v33
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v41 = F_valkey_realloc(m, v38, l1<<(uint(int32(3))%32))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v41
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v48 = v46 + int32(1)
						if l1 <= v48 {
							v160 = int32(0)
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v58 = (l1 + (v46 ^ int32(-1))) & int32(7)
							if v58 == int32(0) {
								v89 = v48
							} else {
								v65 = int32(0)
								v66 = v48
								for {
									*(*int32)(unsafe.Add(mBase, uint32(v53+v66<<(uint(int32(4))%32)))) = int32(0)
									v80 = int32(1)
									v81 = v66 + v80
									v83 = v65 + v80
									if v83 != v58 {
										v65 = v83
										v66 = v81
										continue
									} else {
										break
									}
									break
								}
								v89 = v81
							}
							if base.Ui32(l1-v46+int32(-2)) < base.Ui32(int32(7)) {
								v160 = int32(0)
							} else {
								v119 = v89
								for {
									v128 = int32(0)
									v130 = v119 << (uint(int32(4)) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(v53+v130))) = v128
									*(*int32)(unsafe.Add(mBase, uint32(v53+int32(16)+v130))) = v128
									*(*int32)(unsafe.Add(mBase, uint32(v53+int32(32)+v130))) = v128
									*(*int32)(unsafe.Add(mBase, uint32(v53+int32(48)+v130))) = v128
									*(*int32)(unsafe.Add(mBase, uint32(v53+int32(64)+v130))) = v128
									*(*int32)(unsafe.Add(mBase, uint32(v53+int32(80)+v130))) = v128
									*(*int32)(unsafe.Add(mBase, uint32(v53+int32(96)+v130))) = v128
									*(*int32)(unsafe.Add(mBase, uint32(v53+int32(112)+v130))) = v128
									v156 = v119 + int32(8)
									if v156 != l1 {
										v119 = v156
										continue
									} else {
										break
									}
									break
								}
								v160 = v128
							}
						}
						v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
						if v171&int32(32) == int32(0) {
						} else {
						}
						return v160
					}
				}
			}
		}
	}
}
func F_aeSetAfterSleepProc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l1
	return
}
func F_aeSetDontWait(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v3&int32(-5) | base.B2i32(l1 != int32(0))<<(uint(int32(2))%32)
	return
}
