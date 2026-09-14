package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_writeToClient(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v2 = int32(0)
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v4 != 0 {
		v39 = v2
		return v39
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
		if v5 != 0 {
			v39 = v2
			return v39
		} else {
			v6 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+220)) = uint16(v6)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v6
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			if v10&int32(1) != 0 {
				v30 = F__writeToClient(m, l0)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v37 = F_postWriteToClient(m, l0)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = v37
						return v39
					}
				}
			} else {
				if v10&int32(2) == int32(0) {
					if v10&int32(262144) != 0 {
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						if v23 == int32(0) {
						} else {
						}
					}
					v30 = F__writeToClient(m, l0)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v37 = F_postWriteToClient(m, l0)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = v37
							return v39
						}
					}
				} else {
					if v10&int32(4) == int32(0) {
						F_writeToReplica(m, l0)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v37 = F_postWriteToClient(m, l0)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = v37
								return v39
							}
						}
					} else {
						if v10&int32(262144) != 0 {
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							if v23 == int32(0) {
							} else {
							}
						}
						v30 = F__writeToClient(m, l0)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v37 = F_postWriteToClient(m, l0)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = v37
								return v39
							}
						}
					}
				}
			}
		}
	}
}
