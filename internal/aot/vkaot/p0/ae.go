package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_aeApiPoll(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v21 = F__emscripten_memcpy_bulkmem(m, v15+int32(256), v15, int32(128))
	mBase = m.M
	v25 = int32(128)
	v30 = F__emscripten_memcpy_bulkmem(m, v15+int32(384), v15+v25, v25)
	mBase = m.M
	v32 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = int32(1)
	v37 = F_select_(m, v33+v34, v21, v30, v32, l1)
	mBase = m.M
	if v37 < v34 {
		if v37 != int32(-1) {
			v133 = v32
			m.G0 = v13 + int32(16)
			return v133
		} else {
			v114 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			if v114 == int32(27) {
				v133 = v32
				m.G0 = v13 + int32(16)
				return v133
			} else {
				v117 = F___strerror_l(m, v114, v114)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v117
				F__serverPanic_1(m, int32(_a51), int32(99), int32(_a52), v13)
				mBase = m.M
				v125 = m.ExcPending
				if v125 != 0 {
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
	} else {
		v40 = int32(0)
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v41 < v40 {
			v133 = v40
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v45 = int32(0)
			v50 = v45
			v53 = v45
			for {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v44+v50<<(uint(int32(4))%32))))
				if v60 == int32(0) {
					v104 = v53
				} else {
					v63 = int32(0)
					if v60&int32(1) == v63 {
						v77 = v63
					} else {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(base.Ui32(v50)>>(uint(int32(3))%32))&int32(536870908))))
						v77 = int32(base.Ui32(v73)>>(uint(v50)%32)) & int32(1)
					}
					if v60&int32(2) == int32(0) {
						v94 = v77
					} else {
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(base.Ui32(v50)>>(uint(int32(3))%32))&int32(536870908))))
						if int32(base.Ui32(v89)>>(uint(v50)%32))&int32(1) != 0 {
							v93 = v77 | int32(2)
						} else {
							v93 = v77
						}
						v94 = v93
					}
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v98 = v95 + v53<<(uint(int32(3))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v94
					*(*int32)(unsafe.Add(mBase, uint32(v98))) = v50
					v104 = v53 + int32(1)
				}
				if base.B2i32(v50 == v41) == int32(0) {
					v50 = v50 + int32(1)
					v53 = v104
					continue
				} else {
					break
				}
				break
			}
			v133 = v104
		}
		m.G0 = v13 + int32(16)
		return v133
	}
}
func F_aeDeleteTimeEvent(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-1)
L2:
	;
	v6 = v3
	goto L3
L3:
	;
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v8 != l1 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
	if v14 != 0 {
		v6 = v14
		goto L3
	} else {
		goto L7
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(-1)
	return int32(0)
L7:
	;
	goto L4
}
func F_aePoll(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v3&int32(32) == int32(0) {
	} else {
	}
	v11 = F_aeApiPoll(m, l0, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
		if v15&int32(32) == int32(0) {
		} else {
		}
		return v11
	}
}
func F_aeSetBeforeSleepProc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l1
	return
}
func F_aeSetCustomPollProc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l1
	return
}
func F_aeSetPollProtect(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v3&int32(-33) | base.B2i32(l1 != int32(0))<<(uint(int32(5))%32)
	return
}
