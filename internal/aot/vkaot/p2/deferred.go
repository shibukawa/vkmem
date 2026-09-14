package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_commitDeferredReplyBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	if l1 == int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
		if v9 == int64(-1) {
			F_listEmpty(m, v8)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+384)) = int64(-1)
				return
			}
		} else {
			if v8 == int32(0) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
				if v23 == int32(0) {
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v27
					if v27 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v26
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v26
					}
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v33
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v35 + v23
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(0)
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(0)
				}
				v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
				v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v44 + v45
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
				F_listEmpty(m, v48)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+384)) = int64(-1)
					v53 = F_prepareClientToWrite(m, l0)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						if v53 != 0 {
						} else {
						}
						return
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
				if v14 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
					if v23 == int32(0) {
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v27
						if v27 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v26
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v26
						}
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v33
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v35 + v23
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(0)
					}
					v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
					v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v44 + v45
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
					F_listEmpty(m, v48)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+384)) = int64(-1)
						v53 = F_prepareClientToWrite(m, l0)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							if v53 != 0 {
							} else {
							}
							return
						}
					}
				} else {
					F_listEmpty(m, v8)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+384)) = int64(-1)
						return
					}
				}
			}
		}
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
		if v5&int32(16) != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
			v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
			if v9 == int64(-1) {
				F_listEmpty(m, v8)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+384)) = int64(-1)
					return
				}
			} else {
				if v8 == int32(0) {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
					if v23 == int32(0) {
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v27
						if v27 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v26
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v26
						}
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v33
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v35 + v23
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(0)
					}
					v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
					v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v44 + v45
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
					F_listEmpty(m, v48)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+384)) = int64(-1)
						v53 = F_prepareClientToWrite(m, l0)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							if v53 != 0 {
							} else {
							}
							return
						}
					}
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
					if v14 != 0 {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
						if v23 == int32(0) {
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v26))) = v27
							if v27 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v19))) = v26
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v26
							}
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v33
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v35 + v23
							*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(0)
						}
						v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
						v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v44 + v45
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
						F_listEmpty(m, v48)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+384)) = int64(-1)
							v53 = F_prepareClientToWrite(m, l0)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								if v53 != 0 {
								} else {
								}
								return
							}
						}
					} else {
						F_listEmpty(m, v8)
						mBase = m.M
						v16 = m.ExcPending
						if v16 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+384)) = int64(-1)
							return
						}
					}
				}
			}
		}
	}
}
func F_setDeferredAttributeLen(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(128)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if base.Ui32(v9) <= base.Ui32(int32(2)) {
		F__serverAssert(m, int32(_a_F_setDeferredAttributeLen_0), int32(_a_F_setDeferredAttributeLen_1), int32(1294))
		mBase = m.M
		v80 = m.ExcPending
		if v80 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if l2 <= int32(-1) {
			F__serverAssert(m, int32(_a_F_setDeferredAttributeLen_2), int32(_a_F_setDeferredAttributeLen_1), int32(1246))
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			if l1 == int32(0) {
				m.G0 = v7 + int32(128)
				return
			} else {
				v16 = int32(124)
				*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v16)
				v19 = v7 | int32(1)
				v21 = base.I64_extend_i32_u(l2)
				if v21 <= int64(-1) {
					v30 = int32(45)
					*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v30)
					v34 = int32(1)
					v39 = v19 + v34
					v40 = int32(126)
					v41 = int64(0) - v21
					v42 = v34
				} else {
					v39 = v19
					v40 = int32(127)
					v41 = v21
					v42 = int32(0)
				}
				v43 = F_ull2string(m, v39, v40, v41)
				mBase = m.M
				if v43 == int32(0) {
					v62 = int32(0)
				} else {
					v62 = v43 + v42
				}
				v66 = int32(2573)
				*(*uint16)(unsafe.Add(mBase, uint32(v62+v7+int32(1)))) = uint16(v66)
				F_setDeferredReply(m, l0, l1, v7, v62+int32(3))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					m.G0 = v7 + int32(128)
					return
				}
			}
		}
	}
}
func F_setDeferredSetLen(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v6 == int32(2) {
		v9 = int32(42)
	} else {
		v9 = int32(126)
	}
	F_setDeferredAggregateLen(m, l0, l1, l2, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		return
	}
}
