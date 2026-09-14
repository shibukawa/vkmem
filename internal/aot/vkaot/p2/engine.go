package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_engineStatsDispose(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_valkey_free(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_getEngineUsedMemory_2(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_scriptingEngineCallGetMemoryInfo(m, v6, l0, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v11 + v12
		m.G0 = v6 + int32(16)
		return
	}
}
