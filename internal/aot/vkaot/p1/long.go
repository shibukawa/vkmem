package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_getLongLongFromObjectOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_getLongLongFromObject(m, l1, v8+int32(8))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v27
			v30 = int32(0)
			m.G0 = v8 + int32(16)
			return v30
		} else {
			if l3 == int32(0) {
				F_addReplyError(m, l0, int32(_a1762))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v30 = int32(-1)
					m.G0 = v8 + int32(16)
					return v30
				}
			} else {
				F_addReplyError(m, l0, l3)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v30 = int32(-1)
					m.G0 = v8 + int32(16)
					return v30
				}
			}
		}
	}
}
