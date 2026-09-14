package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_killSlotMigrationChild(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_killSlotMigrationChild[0]))
	if v8 != int32(5) {
		m.G0 = v5 + int32(32)
		return
	} else {
		v11 = int32(_a_F_killSlotMigrationChild_0)
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_killSlotMigrationChild[1]))
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_killSlotMigrationChild[2]))
		if v14 != int32(-1) {
			if int32(2) < v12 {
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_killSlotMigrationChild[3]))
				v40 = F_kill(m, v38, int32(10))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					m.G0 = v5 + int32(32)
					return
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_killSlotMigrationChild[3]))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v29
				F__serverLog(m, int32(2), int32(_a_F_killSlotMigrationChild_1), v5+int32(16))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, _c_F_killSlotMigrationChild[3]))
					v40 = F_kill(m, v38, int32(10))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						m.G0 = v5 + int32(32)
						return
					}
				}
			}
		} else {
			if int32(2) < v12 {
				m.G0 = v5 + int32(32)
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, _c_F_killSlotMigrationChild[3]))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v20
				F__serverLog(m, int32(2), int32(_a_F_killSlotMigrationChild_2), v5)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					m.G0 = v5 + int32(32)
					return
				}
			}
		}
	}
}
