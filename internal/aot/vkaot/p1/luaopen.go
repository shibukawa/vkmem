package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaopen_cjson(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	v2 = F_lua_cjson_new(m, l0)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v26 = v23 + int32(-16)
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v63 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
		*(*int64)(unsafe.Add(mBase, uint32(v62))) = v63
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v65
		v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67 + int32(16)
		v72 = m.G3
		F_lua_setfield(m, l0, int32(-10002), v72+int32(_a2739))
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return int32(0)
		} else {
			return int32(1)
		}
	}
}
func F_luaopen_cmsgpack(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	v2 = F_luaopen_create(m, l0)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v26 = v23 + int32(-16)
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v63 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
		*(*int64)(unsafe.Add(mBase, uint32(v62))) = v63
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v65
		v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67 + int32(16)
		v72 = m.G3
		F_lua_setfield(m, l0, int32(-10002), v72+int32(_a2759))
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return int32(0)
		} else {
			return int32(1)
		}
	}
}
func F_luaopen_string(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v117 int32
	_ = v117
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v215 int32
	_ = v215
	v3 = m.G3
	F_luaL_register(m, l0, v3+int32(_a2645), v3+int32(_a2717))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_lua_getfield(m, l0, int32(-1), v3+int32(_a2718))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			F_lua_setfield(m, l0, int32(-2), v3+int32(_a2719))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_lua_createtable(m, l0, int32(0), int32(1))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_lua_pushlstring(m, l0, v3+int32(_a320), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v51 = v48 + int32(-32)
						v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v88 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
						*(*int64)(unsafe.Add(mBase, uint32(v87))) = v88
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = v90
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(16)
						v97 = F_lua_setmetatable(m, l0, int32(-2))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int32(0)
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v117 + int32(-16)
							v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v147 = v144 + int32(-32)
							v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v184 = *(*int64)(unsafe.Add(mBase, uint32(v147)))
							*(*int64)(unsafe.Add(mBase, uint32(v183))) = v184
							v186 = *(*int32)(unsafe.Add(mBase, uint32(v147)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v183)+8)) = v186
							v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v188 + int32(16)
							F_lua_setfield(m, l0, int32(-2), v3+int32(_a2626))
							mBase = m.M
							v196 = m.ExcPending
							if v196 != 0 {
								return int32(0)
							} else {
								v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v215 + int32(-16)
								return int32(1)
							}
						}
					}
				}
			}
		}
	}
}
func F_luaopen_table(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	v3 = m.G3
	F_luaL_register(m, l0, v3+int32(_a2713), v3+int32(_a2714))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
