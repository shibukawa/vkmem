package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___vfprintf_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	v10 = m.G0
	v12 = v10 - int32(208)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = l2
	v20 = F__emscripten_memset_bulkmem(m, v12+int32(160), base.I32_extend8_s(int32(0)), int32(40))
	mBase = m.M
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v21
	v30 = F_printf_core(m, int32(0), l1, v12+int32(200), v12+int32(80), v12+int32(160), l3, l4)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v30 {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			if int32(0) <= v37 {
				v44 = int32(0)
			} else {
				v44 = int32(1)
			}
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v45 & int32(-33)
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v49 != 0 {
				v58 = int32(0)
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v59 != 0 {
					v89 = v58
					v96 = F_printf_core(m, l0, l1, v12+int32(200), v12+int32(80), v12+int32(160), l3, l4)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v98 = v96
						v99 = v89
						if v99 == int32(0) {
							v119 = v98
							v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v121 | v45&int32(32)
							if v121&int32(32) != 0 {
								v127 = int32(-1)
							} else {
								v127 = v119
							}
							if v44 != 0 {
								v130 = v127
							} else {
								v130 = v127
							}
							m.G0 = v12 + int32(208)
							return v130
						} else {
							v104 = int32(0)
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v107 = m.T0[v106].(func(*base.Module, int32, int32, int32) int32)(m, l0, v104, v104)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								v109 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v109
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v99
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v109
								v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
								if v114 != 0 {
									v118 = v98
								} else {
									v118 = int32(-1)
								}
								v119 = v118
								v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v121 | v45&int32(32)
								if v121&int32(32) != 0 {
									v127 = int32(-1)
								} else {
									v127 = v119
								}
								if v44 != 0 {
									v130 = v127
								} else {
									v130 = v127
								}
								m.G0 = v12 + int32(208)
								return v130
							}
						}
					}
				} else {
					v60 = v58
					v61 = int32(-1)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v63 + v61 | v63
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v68&int32(8) == int32(0) {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v79
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v79
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v79 + v82
						v87 = int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v68 | int32(32)
						v87 = int32(-1)
					}
					if v87 != 0 {
						v98 = v61
						v99 = v60
						if v99 == int32(0) {
							v119 = v98
							v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v121 | v45&int32(32)
							if v121&int32(32) != 0 {
								v127 = int32(-1)
							} else {
								v127 = v119
							}
							if v44 != 0 {
								v130 = v127
							} else {
								v130 = v127
							}
							m.G0 = v12 + int32(208)
							return v130
						} else {
							v104 = int32(0)
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v107 = m.T0[v106].(func(*base.Module, int32, int32, int32) int32)(m, l0, v104, v104)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								v109 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v109
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v99
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v109
								v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
								if v114 != 0 {
									v118 = v98
								} else {
									v118 = int32(-1)
								}
								v119 = v118
								v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v121 | v45&int32(32)
								if v121&int32(32) != 0 {
									v127 = int32(-1)
								} else {
									v127 = v119
								}
								if v44 != 0 {
									v130 = v127
								} else {
									v130 = v127
								}
								m.G0 = v12 + int32(208)
								return v130
							}
						}
					} else {
						v89 = v60
						v96 = F_printf_core(m, l0, l1, v12+int32(200), v12+int32(80), v12+int32(160), l3, l4)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							v98 = v96
							v99 = v89
							if v99 == int32(0) {
								v119 = v98
								v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v121 | v45&int32(32)
								if v121&int32(32) != 0 {
									v127 = int32(-1)
								} else {
									v127 = v119
								}
								if v44 != 0 {
									v130 = v127
								} else {
									v130 = v127
								}
								m.G0 = v12 + int32(208)
								return v130
							} else {
								v104 = int32(0)
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								v107 = m.T0[v106].(func(*base.Module, int32, int32, int32) int32)(m, l0, v104, v104)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									v109 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v109
									*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v99
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v109
									v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
									if v114 != 0 {
										v118 = v98
									} else {
										v118 = int32(-1)
									}
									v119 = v118
									v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v121 | v45&int32(32)
									if v121&int32(32) != 0 {
										v127 = int32(-1)
									} else {
										v127 = v119
									}
									if v44 != 0 {
										v130 = v127
									} else {
										v130 = v127
									}
									m.G0 = v12 + int32(208)
									return v130
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(80)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v12
				v60 = v56
				v61 = int32(-1)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v63 + v61 | v63
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v68&int32(8) == int32(0) {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v79
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v79
					v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v79 + v82
					v87 = int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v68 | int32(32)
					v87 = int32(-1)
				}
				if v87 != 0 {
					v98 = v61
					v99 = v60
					if v99 == int32(0) {
						v119 = v98
						v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v121 | v45&int32(32)
						if v121&int32(32) != 0 {
							v127 = int32(-1)
						} else {
							v127 = v119
						}
						if v44 != 0 {
							v130 = v127
						} else {
							v130 = v127
						}
						m.G0 = v12 + int32(208)
						return v130
					} else {
						v104 = int32(0)
						v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v107 = m.T0[v106].(func(*base.Module, int32, int32, int32) int32)(m, l0, v104, v104)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							v109 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v109
							*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v99
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v109
							v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
							if v114 != 0 {
								v118 = v98
							} else {
								v118 = int32(-1)
							}
							v119 = v118
							v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v121 | v45&int32(32)
							if v121&int32(32) != 0 {
								v127 = int32(-1)
							} else {
								v127 = v119
							}
							if v44 != 0 {
								v130 = v127
							} else {
								v130 = v127
							}
							m.G0 = v12 + int32(208)
							return v130
						}
					}
				} else {
					v89 = v60
					v96 = F_printf_core(m, l0, l1, v12+int32(200), v12+int32(80), v12+int32(160), l3, l4)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v98 = v96
						v99 = v89
						if v99 == int32(0) {
							v119 = v98
							v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v121 | v45&int32(32)
							if v121&int32(32) != 0 {
								v127 = int32(-1)
							} else {
								v127 = v119
							}
							if v44 != 0 {
								v130 = v127
							} else {
								v130 = v127
							}
							m.G0 = v12 + int32(208)
							return v130
						} else {
							v104 = int32(0)
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v107 = m.T0[v106].(func(*base.Module, int32, int32, int32) int32)(m, l0, v104, v104)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								v109 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v109
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v99
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v109
								v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
								if v114 != 0 {
									v118 = v98
								} else {
									v118 = int32(-1)
								}
								v119 = v118
								v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v121 | v45&int32(32)
								if v121&int32(32) != 0 {
									v127 = int32(-1)
								} else {
									v127 = v119
								}
								if v44 != 0 {
									v130 = v127
								} else {
									v130 = v127
								}
								m.G0 = v12 + int32(208)
								return v130
							}
						}
					}
				}
			}
		} else {
			v130 = int32(-1)
			m.G0 = v12 + int32(208)
			return v130
		}
	}
}
func F___vsyslog(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v26 int32
	_ = v26
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if base.Ui32(int32(1023)) < base.Ui32(l0) {
		m.G0 = v7 + int32(16)
		return
	} else {
		v11 = int32(0)
		v12 = *(*int32)(unsafe.Add(mBase, _consts[1264]))
		if v12&(int32(1)<<(uint(l0&int32(7))%32)) == v11 {
			m.G0 = v7 + int32(16)
			return
		} else {
			F__vsyslog(m, l0, l1, l2)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
func F_vectorGet(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(l1) < base.Ui32(v3) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		return v13 + v14*l1
	} else {
		F__serverAssert(m, int32(_a1668), int32(_a1667), int32(39))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
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
func F_vectorInit(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	if l2 == int32(0) {
		F__serverAssert(m, int32(_a1666), int32(_a1667), int32(21))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v7 = int32(0)
		if l1 == v7 {
			v13 = v7
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v13
			return
		} else {
			v11 = F_valkey_malloc(m, l2*l1)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = v11
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v13
				return
			}
		}
	}
}
func F_verbatimStringCallback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 < int32(0) {
		v33 = v8
	} else {
		v14 = l0 + v9<<(uint(int32(2))%32)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
		v16 = int32(1)
		v17 = v15 + v16
		*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v17
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1040))
		if v19 != v16 {
			v33 = v8
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v24))) = base.F64_convert_i32_u(v17)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v28 + int32(16)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v33 = v32
		}
	}
	v37 = F_lua_checkstack(m, v33, int32(5))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		return
	} else {
		if v37 != 0 {
			v42 = int32(0)
			F_lua_createtable(m, v33, v42, v42)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				v46 = m.G3
				F_lua_pushstring(m, v33, v46+int32(_a2235))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					v51 = int32(0)
					F_lua_createtable(m, v33, v51, v51)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						F_lua_pushstring(m, v33, v46+int32(_a2173))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_lua_pushlstring(m, v33, l1, l2)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								F_lua_settable(m, v33, int32(-3))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_lua_pushstring(m, v33, v46+int32(_a2236))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_lua_pushlstring(m, v33, l3, int32(3))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											F_lua_settable(m, v33, int32(-3))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												F_lua_settable(m, v33, int32(-3))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return
												} else {
													F_processCollectionElementEnd(m, l0)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return
													} else {
														return
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
			}
		} else {
			F__serverPanic_2(m, int32(1046))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_version2num(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v62 int32
	_ = v62
	v2 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v11 = l0
	v12 = v7
	v13 = v2
	v14 = v2
	v15 = v2
	goto L2
L1:
	;
	return v62
L2:
	;
	v20 = (v12 + int32(-48)) & int32(255)
	if base.Ui32(int32(9)) < base.Ui32(v20) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if v45 == int32(2) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v47 != 0 {
		v11 = v11 + int32(1)
		v12 = v47
		v13 = v43
		v14 = v44
		v15 = v45
		goto L2
	} else {
		goto L10
	}
L5:
	;
	v30 = int32(-1)
	if v12&int32(255) != int32(46) {
		v62 = v30
		goto L1
	} else {
		goto L8
	}
L6:
	;
	v25 = v14*int32(10) + v20
	if v25 <= int32(255) {
		v43 = v13
		v44 = v25
		v45 = v15
		goto L4
	} else {
		goto L7
	}
L7:
	;
	return int32(-1)
L8:
	;
	if int32(1) < v15 {
		v62 = v30
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v43 = v13<<(uint(int32(8))%32) | v14
	v44 = int32(0)
	v45 = v15 + int32(1)
	goto L4
L10:
	;
	goto L3
L11:
	;
	v62 = v43<<(uint(int32(8))%32) | v44
	goto L1
L12:
	;
	return int32(-1)
}
func F_vkmem_dlsym(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	if l0 != int32(_a0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v75 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(_a1)
	return v75
L2:
	;
	v5 = int32(_a2)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v9 == int32(0) {
		v32 = v8
		v33 = v9
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v40 = int32(_a3)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
	if v44 == int32(0) {
		v67 = v43
		v68 = v44
		goto L14
	} else {
		goto L15
	}
L4:
	;
	if v33-v32&int32(255) != 0 {
		goto L3
	} else {
		goto L12
	}
L5:
	;
	goto L4
L6:
	;
	if v9 != v8&int32(255) {
		v32 = v8
		v33 = v9
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v15 = v5
	v16 = l1
	goto L8
L8:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v20 == int32(0) {
		v32 = v19
		v33 = v20
		goto L5
	} else {
		goto L10
	}
L9:
	;
	v32 = v19
	v33 = v20
	goto L5
L10:
	;
	v23 = int32(1)
	if v20 == v19&int32(255) {
		v15 = v15 + v23
		v16 = v16 + v23
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	return v38
L13:
	;
	if v68-v67&int32(255) != 0 {
		goto L1
	} else {
		goto L21
	}
L14:
	;
	goto L13
L15:
	;
	if v44 != v43&int32(255) {
		v67 = v43
		v68 = v44
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v50 = v40
	v51 = l1
	goto L17
L17:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v55 == int32(0) {
		v67 = v54
		v68 = v55
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v67 = v54
	v68 = v55
	goto L14
L19:
	;
	v58 = int32(1)
	if v55 == v54&int32(255) {
		v50 = v50 + v58
		v51 = v51 + v58
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	return v73
}
func F_vsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v7 = m.G0
	v9 = v7 - int32(160)
	m.G0 = v9
	if l1 != 0 {
		v13 = l0
	} else {
		v13 = v9 + int32(158)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v9)+148)) = v13
	v17 = l1 + int32(-1)
	if base.Ui32(l1) < base.Ui32(v17) {
		v19 = int32(0)
	} else {
		v19 = v17
	}
	*(*int32)(unsafe.Add(mBase, uint32(v9)+152)) = v19
	v24 = F__emscripten_memset_bulkmem(m, v9, base.I32_extend8_s(int32(0)), int32(144))
	mBase = m.M
	v25 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+76)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = int32(1389)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v24 + int32(159)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = v24 + int32(148)
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v37)
	v39 = F_vfprintf(m, v24, l2, l3)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		return int32(0)
	} else {
		m.G0 = v24 + int32(160)
		return v39
	}
}
func F_vsnprintf_async_signal_safe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v171 int64
	_ = v171
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v180 int64
	_ = v180
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
	var v201 int32
	_ = v201
	var v203 int64
	_ = v203
	var v213 int32
	_ = v213
	var v220 int64
	_ = v220
	var v224 int32
	_ = v224
	var v225 int64
	_ = v225
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int64
	_ = v247
	var v250 int64
	_ = v250
	var v258 int64
	_ = v258
	var v260 int32
	_ = v260
	var v261 int64
	_ = v261
	var v271 int32
	_ = v271
	var v278 int64
	_ = v278
	var v281 int64
	_ = v281
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v361 int32
	_ = v361
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v386 int64
	_ = v386
	var v387 int32
	_ = v387
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	v22 = l0 + l1 + int32(-1)
	v23 = m.G0
	v25 = v23 - int32(32)
	v27 = v25 | int32(5)
	v29 = v25 + int32(20)
	v33 = l0
	v34 = l2
	v35 = l3
	goto L1
L1:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v51 == int32(37) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v67 = int32(1)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	if v68 == int32(108) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	if v51 == int32(0) {
		v57 = v33
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v51)
	v63 = int32(1)
	v33 = v33 + v63
	v34 = v34 + v63
	goto L1
L6:
	;
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v58)
	return v57 - l0
L7:
	;
	if v33 != v22 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v57 = v22
	goto L6
L9:
	;
	v86 = v83 & int32(255)
	v88 = v86 + int32(-100)
	if base.Ui32(int32(20)) < base.Ui32(v88) {
		v438 = v33
		v440 = v35
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+2)))
	if v73 == int32(108) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v82 = v34 + int32(1)
	v83 = v68
	v84 = v67
	goto L9
L12:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+3)))
	v82 = v34 + int32(3)
	v83 = v78
	v84 = int32(0)
	goto L9
L13:
	;
	v82 = v34 + int32(2)
	v83 = v73
	v84 = v67
	goto L9
L14:
	;
	v33 = v438
	v34 = v82 + int32(1)
	v35 = v440
	goto L1
L15:
	;
	if int32(1)<<(uint(v88)%32)&int32(1179681) != 0 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v398 == int32(120) {
		goto L81
	} else {
		goto L82
	}
L17:
	;
	v236 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+21)) = uint8(v236)
	v241 = base.B2i32(v86 == int32(120))
	v243 = base.B2i32(v184 < int64(0)) & (v241 | v189)
	if v243 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L18:
	;
	v201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+21)) = uint8(v201)
	v203 = base.I64_extend_i32_u(v195)
	v213 = v25 + int32(21)
	v220 = v200
	goto L44
L19:
	;
	v189 = base.B2i32(v86 == int32(112))
	if v86 == int32(112) {
		goto L40
	} else {
		goto L41
	}
L20:
	;
	v180 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35))))
	v183 = v35 + int32(4)
	v184 = v180
	v185 = int32(1)
	goto L19
L21:
	;
	if v84 != 0 {
		goto L36
	} else {
		goto L37
	}
L22:
	;
	if v88 == int32(12) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	if v88 != int32(15) {
		v438 = v33
		v440 = v35
		goto L14
	} else {
		goto L24
	}
L24:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v101 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v33 = v137
	v34 = v82 + int32(1)
	v35 = v35 + int32(4)
	goto L1
L26:
	;
	v103 = v101
	goto L28
L27:
	;
	v103 = int32(_a711)
	goto L28
L28:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v104 == int32(0) {
		v137 = v33
		goto L25
	} else {
		goto L29
	}
L29:
	;
	if base.Ui32(v22) <= base.Ui32(v33) {
		v137 = v33
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v109 = v33
	v117 = v103
	v119 = v104
	goto L31
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v119)
	v129 = v109 + int32(1)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	if v130 == int32(0) {
		v137 = v129
		goto L25
	} else {
		goto L33
	}
L32:
	;
	v137 = v129
	goto L25
L33:
	;
	if base.Ui32(v129) < base.Ui32(v22) {
		v109 = v129
		v117 = v117 + int32(1)
		v119 = v130
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v194 = v84
	v195 = int32(10)
	v196 = v173
	v199 = int64(0)
	v200 = v174
	goto L18
L36:
	;
	if v86 != int32(117) {
		goto L20
	} else {
		goto L39
	}
L37:
	;
	v160 = (v35 + int32(7)) & int32(-8)
	v162 = v160 + int32(8)
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v160)))
	if v86 == int32(117) {
		v173 = v162
		v174 = v163
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v183 = v162
	v184 = v163
	v185 = int32(0)
	goto L19
L39:
	;
	v171 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35))))
	v173 = v35 + int32(4)
	v174 = v171
	goto L35
L40:
	;
	v190 = int32(16)
	goto L42
L41:
	;
	v190 = int32(10)
	goto L42
L42:
	;
	if v86 != int32(117) {
		goto L17
	} else {
		goto L43
	}
L43:
	;
	v194 = v185
	v195 = v190
	v196 = v183
	v199 = v184
	v200 = int64(0)
	goto L18
L44:
	;
	v224 = v213 + int32(-1)
	v225 = base.I64_div_u_s(v220, v203)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v220-v225*v203))+uint32(_consts[1114]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v231)
	if base.B2i32(base.Ui64(v220) < base.Ui64(v203)) == int32(0) {
		v213 = v224
		v220 = v225
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v376 = v196
	v382 = v224
	v386 = v199
	v387 = v194
	goto L16
L46:
	;
	goto L45
L47:
	;
	if v86 == int32(120) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v250 = int64(3472328296227680304)
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = v250
	*(*int64)(unsafe.Add(mBase, uint32(v27+int32(8)))) = v250
	v258 = v184 ^ int64(-1)
	goto L47
L49:
	;
	v247 = v184 >> (uint(int64(63)) % 64)
	v258 = v184 ^ v247 - v247
	goto L47
L50:
	;
	v260 = int32(16)
	goto L52
L51:
	;
	v260 = v190
	goto L52
L52:
	;
	v261 = base.I64_extend_i32_u(v260)
	v271 = v29
	v278 = v258
	goto L53
L53:
	;
	v281 = base.I64_div_s(v278, v261)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v278-v281*v261))+uint32(_consts[1114]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v271))) = uint8(v287)
	v290 = v271 + int32(-1)
	if v281 != int64(0) {
		v271 = v290
		v278 = v281
		goto L53
	} else {
		goto L55
	}
L54:
	;
	if int64(-1) < v184 {
		v301 = v290
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v302 = int32(0)
	if v243 == v302 {
		v361 = v301
		goto L59
	} else {
		goto L60
	}
L57:
	;
	if v260 != int32(10) {
		v301 = v290
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v297 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v297)
	v301 = v271 + int32(-2)
	goto L56
L59:
	;
	v376 = v183
	v382 = v361 + int32(1)
	v386 = v184
	v387 = v185
	goto L16
L60:
	;
	v315 = v302
	v316 = v29
	goto L61
L61:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316))))
	switch v325 + int32(-48) {
	case 0:
		v343 = int32(102)
		goto L64
	case 1:
		goto L79
	case 2:
		goto L78
	case 3:
		goto L77
	case 4:
		goto L76
	case 5:
		goto L75
	case 6:
		goto L74
	case 7:
		goto L73
	case 8:
		goto L72
	case 9:
		goto L71
	default:
		goto L63
	case 49:
		goto L70
	case 50:
		goto L69
	case 51:
		goto L68
	case 52:
		goto L67
	case 53:
		goto L66
	case 54:
		goto L65
	}
L62:
	;
	v361 = v347
	goto L59
L63:
	;
	v347 = v316 + int32(-1)
	v349 = v315 + int32(1)
	if v349 != int32(16) {
		v315 = v349
		v316 = v347
		goto L61
	} else {
		goto L80
	}
L64:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v316))) = uint8(v343)
	goto L63
L65:
	;
	v343 = int32(48)
	goto L64
L66:
	;
	v343 = int32(49)
	goto L64
L67:
	;
	v343 = int32(50)
	goto L64
L68:
	;
	v343 = int32(51)
	goto L64
L69:
	;
	v343 = int32(52)
	goto L64
L70:
	;
	v343 = int32(53)
	goto L64
L71:
	;
	v343 = int32(54)
	goto L64
L72:
	;
	v343 = int32(55)
	goto L64
L73:
	;
	v343 = int32(56)
	goto L64
L74:
	;
	v343 = int32(57)
	goto L64
L75:
	;
	v343 = int32(97)
	goto L64
L76:
	;
	v343 = int32(98)
	goto L64
L77:
	;
	v343 = int32(99)
	goto L64
L78:
	;
	v343 = int32(100)
	goto L64
L79:
	;
	v343 = int32(101)
	goto L64
L80:
	;
	goto L62
L81:
	;
	v401 = base.I32_wrap_i64(int64(base.Ui64(v386)>>(uint(int64(60))%64))) & int32(8)
	goto L83
L82:
	;
	v401 = int32(0)
	goto L83
L83:
	;
	if v387 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v403 = v401
	goto L86
L85:
	;
	v403 = int32(0)
	goto L86
L86:
	;
	v404 = v382 + v403
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	if v405 == int32(0) {
		v438 = v33
		v440 = v376
		goto L14
	} else {
		goto L87
	}
L87:
	;
	if base.Ui32(v22) <= base.Ui32(v33) {
		v438 = v33
		v440 = v376
		goto L14
	} else {
		goto L88
	}
L88:
	;
	v410 = v33
	v418 = v404
	v420 = v405
	goto L89
L89:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v410))) = uint8(v420)
	v430 = v410 + int32(1)
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+1)))
	if v431 == int32(0) {
		v438 = v430
		v440 = v376
		goto L14
	} else {
		goto L91
	}
L90:
	;
	v438 = v430
	v440 = v376
	goto L14
L91:
	;
	if base.Ui32(v430) < base.Ui32(v22) {
		v410 = v430
		v418 = v418 + int32(1)
		v420 = v431
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
}
