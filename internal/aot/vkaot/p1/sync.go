package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_syncWithPrimaryHandleConnectingState(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v9 {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
		v22 = m.T0[v21].(func(*base.Module, int32, int32) int32)(m, l0, int32(970))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = int32(0)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
			v29 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, l0, v24, v24)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a1970)
				v37 = F_sendCommand(m, l0, v6+int32(16))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					if v37 == int32(0) {
						v53 = v24
						m.G0 = v6 + int32(32)
						return v53
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, _consts[28]))
						if int32(3) < v42 {
							F_sdsfree(m, v37)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v53 = int32(-1)
								m.G0 = v6 + int32(32)
								return v53
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v37
							F__serverLog(m, int32(3), int32(_a1971), v6)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_sdsfree(m, v37)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v53 = int32(-1)
									m.G0 = v6 + int32(32)
									return v53
								}
							}
						}
					}
				}
			}
		}
	} else {
		F__serverLog(m, int32(2), int32(_a1972), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
			v22 = m.T0[v21].(func(*base.Module, int32, int32) int32)(m, l0, int32(970))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = int32(0)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
				v29 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, l0, v24, v24)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a1970)
					v37 = F_sendCommand(m, l0, v6+int32(16))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						if v37 == int32(0) {
							v53 = v24
							m.G0 = v6 + int32(32)
							return v53
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, _consts[28]))
							if int32(3) < v42 {
								F_sdsfree(m, v37)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v53 = int32(-1)
									m.G0 = v6 + int32(32)
									return v53
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v37
								F__serverLog(m, int32(3), int32(_a1971), v6)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									F_sdsfree(m, v37)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										v53 = int32(-1)
										m.G0 = v6 + int32(32)
										return v53
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
func F_syncWrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v98 int32
	_ = v98
	v10 = F_mstime(m)
	mBase = m.M
	v12 = l1
	v16 = l2
	v17 = l3
	goto L3
L1:
	;
	return v98
L2:
	;
	v98 = int32(-1)
	goto L1
L3:
	;
	v20 = F_write(m, l0, v12, v16)
	mBase = m.M
	if v20 != int32(-1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L25
L5:
	;
	if v30 == int32(0) {
		v98 = l2
		goto L1
	} else {
		goto L10
	}
L6:
	;
	v29 = v12 + v20
	v30 = v16 - v20
	goto L5
L7:
	;
	goto L8
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	if v24 == int32(6) {
		v29 = v12
		v30 = v16
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L2
L10:
	;
	v34 = int64(10)
	if v34 < v17 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = v17
	goto L13
L12:
	;
	v37 = v34
	goto L13
L13:
	;
	v39 = m.G0
	v41 = v39 - int32(16)
	m.G0 = v41
	*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = l0
	v46 = int32(4)
	goto L15
L14:
	;
	v85 = F_mstime(m)
	mBase = m.M
	v86 = v85 - v10
	if v86 < l3 {
		v12 = v29
		v16 = v30
		v17 = l3 - v86
		goto L3
	} else {
		goto L24
	}
L15:
	;
	goto L18
L17:
	;
	v62 = int32(1)
	v64 = F_poll(m, v41+int32(8), v62, base.I32_wrap_i64(v37))
	mBase = m.M
	if v64 != v62 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)) = uint16(v46)
	goto L17
L19:
	;
	m.G0 = v41 + int32(16)
	goto L14
L20:
	;
	goto L21
L21:
	;
	goto L23
L23:
	;
	goto L19
L24:
	;
	goto L4
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(73)
	goto L2
}
