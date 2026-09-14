package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_engineFunctionDispose(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	if l0 == int32(0) {
		return
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_scriptingEngineCallFreeFunction(m, v5, int32(1), v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_valkey_free(m, l0)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_engineLibraryDispose(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	if l0 == int32(0) {
		return
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_dictRelease(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			F_sdsfree(m, v7)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_sdsfree(m, v10)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v14 = m.ExcPending
					if v14 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_engineLibraryFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	if l0 == int32(0) {
		return
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_dictRelease(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			F_sdsfree(m, v7)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_sdsfree(m, v10)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v14 = m.ExcPending
					if v14 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_getEngineUsedMemory_1(m *base.Module, l0 int32, l1 int32) {
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
	F_scriptingEngineCallGetMemoryInfo(m, v6, l0, int32(0))
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
func F_resetEngineEvalEnvCallback(m *base.Module, l0 int32, l1 int32) {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v3 = int32(0)
	v6 = F_scriptingEngineCallResetEnvFunc(m, l0, v3, base.B2i32(l1 != v3))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		if l1 == int32(0) {
			return
		} else {
			v10 = F_listAddNodeTail(m, l1, v6)
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_resetEngineOrCollectResetCallbacks(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v6 = F_scriptingEngineCallResetEnvFunc(m, l0, int32(1), base.B2i32(l1 != int32(0)))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		if l1 == int32(0) {
			return
		} else {
			v10 = F_listAddNodeTail(m, l1, v6)
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				return
			}
		}
	}
}
