package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaB_coresume(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int64
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int64
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int64
	_ = v266
	var v268 int32
	_ = v268
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v13 = v8 + int32(0)
	v14 = m.G398
	if base.Ui32(v13) < base.Ui32(v7) {
		v16 = v13
	} else {
		v16 = v14
	}
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v59 != int32(8) {
		v63 = int32(0)
	} else {
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		v63 = v62
	}
	if v63 != 0 {
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v79 = F_auxresume(m, l0, v63, (v72-v73)>>(uint(int32(4))%32)+int32(-1))
		mBase = m.M
		v80 = m.ExcPending
		if v80 != 0 {
			return int32(0)
		} else {
			if int32(-1) < v79 {
				v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v179))) = int32(1)
				v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v185 + int32(16)
				v190 = v79 ^ int32(-1)
				if v190 < int32(1) {
					if v190 < int32(-9999) {
						switch v190 + int32(10002) {
						case 0:
							v245 = l0 + int32(72)
						case 1:
							v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
							v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v219
							v245 = l0 + int32(88)
						case 2:
							v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v245 = v213 + int32(96)
						default:
							v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
							v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
							v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+7)))
							v231 = m.G398
							if base.Ui32(v230) < base.Ui32(int32(-10002)-v190) {
								v242 = v231
							} else {
								v242 = v229 + (int32(-10003)-v190)<<(uint(int32(4))%32) + int32(24)
							}
							v245 = v242
						}
					} else {
						v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v245 = v207 + v190<<(uint(int32(4))%32)
					}
				} else {
					v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v201 = v196 + v190<<(uint(int32(4))%32) + int32(-16)
					v202 = m.G398
					if base.Ui32(v201) < base.Ui32(v195) {
						v204 = v201
					} else {
						v204 = v202
					}
					v245 = v204
				}
				v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v246) <= base.Ui32(v245) {
					v263 = v246
				} else {
					v249 = v246
					for {
						v253 = v249 + int32(-16)
						v254 = *(*int64)(unsafe.Add(mBase, uint32(v253)))
						*(*int64)(unsafe.Add(mBase, uint32(v249))) = v254
						v258 = *(*int32)(unsafe.Add(mBase, uint32(v249+int32(-8))))
						*(*int32)(unsafe.Add(mBase, uint32(v249)+8)) = v258
						if base.Ui32(v245) < base.Ui32(v253) {
							v249 = v253
							continue
						} else {
							break
						}
						break
					}
					v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v263 = v261
				}
				v266 = *(*int64)(unsafe.Add(mBase, uint32(v263)))
				*(*int64)(unsafe.Add(mBase, uint32(v245))) = v266
				v268 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v245)+8)) = v268
				return v79 + int32(1)
			} else {
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(0)
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v91 + int32(16)
				v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v115 = v112 + int32(-32)
				v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v151) <= base.Ui32(v115) {
					v168 = v151
				} else {
					v154 = v151
					for {
						v158 = v154 + int32(-16)
						v159 = *(*int64)(unsafe.Add(mBase, uint32(v158)))
						*(*int64)(unsafe.Add(mBase, uint32(v154))) = v159
						v163 = *(*int32)(unsafe.Add(mBase, uint32(v154+int32(-8))))
						*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v163
						if base.Ui32(v115) < base.Ui32(v158) {
							v154 = v158
							continue
						} else {
							break
						}
						break
					}
					v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v168 = v166
				}
				v171 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
				*(*int64)(unsafe.Add(mBase, uint32(v115))) = v171
				v173 = *(*int32)(unsafe.Add(mBase, uint32(v168)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v115)+8)) = v173
				return int32(2)
			}
		}
	} else {
		v65 = m.G3
		v68 = F_luaL_argerror(m, l0, int32(1), v65+int32(_a2309))
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v79 = F_auxresume(m, l0, v63, (v72-v73)>>(uint(int32(4))%32)+int32(-1))
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int32(0)
			} else {
				if int32(-1) < v79 {
					v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v179))) = int32(1)
					v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v185 + int32(16)
					v190 = v79 ^ int32(-1)
					if v190 < int32(1) {
						if v190 < int32(-9999) {
							switch v190 + int32(10002) {
							case 0:
								v245 = l0 + int32(72)
							case 1:
								v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
								v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
								v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v219
								v245 = l0 + int32(88)
							case 2:
								v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v245 = v213 + int32(96)
							default:
								v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
								v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
								v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+7)))
								v231 = m.G398
								if base.Ui32(v230) < base.Ui32(int32(-10002)-v190) {
									v242 = v231
								} else {
									v242 = v229 + (int32(-10003)-v190)<<(uint(int32(4))%32) + int32(24)
								}
								v245 = v242
							}
						} else {
							v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v245 = v207 + v190<<(uint(int32(4))%32)
						}
					} else {
						v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v201 = v196 + v190<<(uint(int32(4))%32) + int32(-16)
						v202 = m.G398
						if base.Ui32(v201) < base.Ui32(v195) {
							v204 = v201
						} else {
							v204 = v202
						}
						v245 = v204
					}
					v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if base.Ui32(v246) <= base.Ui32(v245) {
						v263 = v246
					} else {
						v249 = v246
						for {
							v253 = v249 + int32(-16)
							v254 = *(*int64)(unsafe.Add(mBase, uint32(v253)))
							*(*int64)(unsafe.Add(mBase, uint32(v249))) = v254
							v258 = *(*int32)(unsafe.Add(mBase, uint32(v249+int32(-8))))
							*(*int32)(unsafe.Add(mBase, uint32(v249)+8)) = v258
							if base.Ui32(v245) < base.Ui32(v253) {
								v249 = v253
								continue
							} else {
								break
							}
							break
						}
						v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v263 = v261
					}
					v266 = *(*int64)(unsafe.Add(mBase, uint32(v263)))
					*(*int64)(unsafe.Add(mBase, uint32(v245))) = v266
					v268 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v245)+8)) = v268
					return v79 + int32(1)
				} else {
					v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(0)
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v91 + int32(16)
					v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v115 = v112 + int32(-32)
					v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if base.Ui32(v151) <= base.Ui32(v115) {
						v168 = v151
					} else {
						v154 = v151
						for {
							v158 = v154 + int32(-16)
							v159 = *(*int64)(unsafe.Add(mBase, uint32(v158)))
							*(*int64)(unsafe.Add(mBase, uint32(v154))) = v159
							v163 = *(*int32)(unsafe.Add(mBase, uint32(v154+int32(-8))))
							*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v163
							if base.Ui32(v115) < base.Ui32(v158) {
								v154 = v158
								continue
							} else {
								break
							}
							break
						}
						v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v168 = v166
					}
					v171 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
					*(*int64)(unsafe.Add(mBase, uint32(v115))) = v171
					v173 = *(*int32)(unsafe.Add(mBase, uint32(v168)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v115)+8)) = v173
					return int32(2)
				}
			}
		}
	}
}
func F_luaB_costatus(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = v14 + int32(0)
	v20 = m.G398
	if base.Ui32(v19) < base.Ui32(v13) {
		v22 = v19
	} else {
		v22 = v20
	}
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v65 != int32(8) {
		v69 = int32(0)
	} else {
		v68 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		v69 = v68
	}
	if v69 != 0 {
		if l0 != v69 {
			v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+6)))
			switch v80 {
			case 0:
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
				v120 = *(*int32)(unsafe.Add(mBase, uint32(v69)+40))
				if base.Ui32(v88) <= base.Ui32(v120) {
					v135 = int32(0)
				} else {
					v124 = base.I32_div_s(v88-v120, int32(24))
					*(*int32)(unsafe.Add(mBase, uint32(v7+int32(12))+96)) = v124
					v135 = int32(1)
				}
				if int32(0) < v135 {
					v147 = int32(2)
				} else {
					v140 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
					v141 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
					if (v140-v141)>>(uint(int32(4))%32) != 0 {
						v145 = int32(1)
					} else {
						v145 = int32(3)
					}
					v147 = v145
				}
			case 1:
				v147 = v80
			default:
				v147 = int32(3)
			}
		} else {
			v147 = int32(0)
		}
		v148 = m.G3
		v154 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(_a2311)+v147<<(uint(int32(2))%32))))
		F_lua_pushstring(m, l0, v154)
		mBase = m.M
		v156 = m.ExcPending
		if v156 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(112)
			return int32(1)
		}
	} else {
		v71 = m.G3
		v74 = F_luaL_argerror(m, l0, int32(1), v71+int32(_a2309))
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
			return int32(0)
		} else {
			if l0 != v69 {
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+6)))
				switch v80 {
				case 0:
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v69)+40))
					if base.Ui32(v88) <= base.Ui32(v120) {
						v135 = int32(0)
					} else {
						v124 = base.I32_div_s(v88-v120, int32(24))
						*(*int32)(unsafe.Add(mBase, uint32(v7+int32(12))+96)) = v124
						v135 = int32(1)
					}
					if int32(0) < v135 {
						v147 = int32(2)
					} else {
						v140 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
						v141 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
						if (v140-v141)>>(uint(int32(4))%32) != 0 {
							v145 = int32(1)
						} else {
							v145 = int32(3)
						}
						v147 = v145
					}
				case 1:
					v147 = v80
				default:
					v147 = int32(3)
				}
			} else {
				v147 = int32(0)
			}
			v148 = m.G3
			v154 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(_a2311)+v147<<(uint(int32(2))%32))))
			F_lua_pushstring(m, l0, v154)
			mBase = m.M
			v156 = m.ExcPending
			if v156 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(112)
				return int32(1)
			}
		}
	}
}
func F_luaB_getmetatable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	F_luaL_checkany(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = v13 + int32(0)
		v19 = m.G398
		if base.Ui32(v18) < base.Ui32(v12) {
			v21 = v18
		} else {
			v21 = v19
		}
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
		switch v63 + int32(-5) {
		case 0:
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v78 = v66 + int32(16)
		default:
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v78 = v72 + v63<<(uint(int32(2))%32) + int32(152)
		case 2:
			v69 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v78 = v69 + int32(8)
		}
		v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
		if v79 != 0 {
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = int32(5)
			*(*int32)(unsafe.Add(mBase, uint32(v81))) = v79
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v85 + int32(16)
		} else {
		}
		if v79 != 0 {
			v102 = m.G3
			v105 = F_luaL_getmetafield(m, l0, int32(1), v102+int32(_a2300))
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return int32(0)
			} else {
				return int32(1)
			}
		} else {
			v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v93 + int32(16)
			return int32(1)
		}
	}
}
func F_luaB_load(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int64
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	v6 = m.G3
	v10 = F_luaL_optlstring(m, l0, int32(2), v6+int32(_a2301), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		F_luaL_checktype(m, l0, int32(1), int32(6))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v26 = v23 + int32(48)
			if base.Ui32(v26) <= base.Ui32(v22) {
			} else {
				v30 = v22
				for {
					*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(0)
					v34 = v30 + int32(16)
					if base.Ui32(v34) < base.Ui32(v26) {
						v30 = v34
						continue
					} else {
						break
					}
					break
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v26
			v46 = m.G5
			v50 = F_lua_load(m, l0, v46+int32(1243), int32(0), v10)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				if v50 == int32(0) {
					v142 = int32(1)
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v55 + int32(16)
					v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v81 = v78 + int32(-32)
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if base.Ui32(v117) <= base.Ui32(v81) {
						v134 = v117
					} else {
						v120 = v117
						for {
							v124 = v120 + int32(-16)
							v125 = *(*int64)(unsafe.Add(mBase, uint32(v124)))
							*(*int64)(unsafe.Add(mBase, uint32(v120))) = v125
							v129 = *(*int32)(unsafe.Add(mBase, uint32(v120+int32(-8))))
							*(*int32)(unsafe.Add(mBase, uint32(v120)+8)) = v129
							if base.Ui32(v81) < base.Ui32(v124) {
								v120 = v124
								continue
							} else {
								break
							}
							break
						}
						v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v134 = v132
					}
					v137 = *(*int64)(unsafe.Add(mBase, uint32(v134)))
					*(*int64)(unsafe.Add(mBase, uint32(v81))) = v137
					v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v139
					v142 = int32(2)
				}
				return v142
			}
		}
	}
}
func F_luaB_pairs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		switch int32(-1) {
		case 0:
			v61 = l0 + int32(72)
		case 1:
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v37
			v61 = l0 + int32(88)
		case 2:
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v61 = v31 + int32(96)
		default:
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
			v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+7)))
			v49 = m.G398
			if base.Ui32(v48) < base.Ui32(int32(1)) {
				v60 = v49
			} else {
				v60 = v47 + int32(24)
			}
			v61 = v60
		}
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v65 = *(*int64)(unsafe.Add(mBase, uint32(v61)))
		*(*int64)(unsafe.Add(mBase, uint32(v64))) = v65
		v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v67
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v69 + int32(16)
		v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v84 = v79 + int32(0)
		v85 = m.G398
		if base.Ui32(v84) < base.Ui32(v78) {
			v87 = v84
		} else {
			v87 = v85
		}
		v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v130 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
		*(*int64)(unsafe.Add(mBase, uint32(v129))) = v130
		v132 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v129)+8)) = v132
		v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v134 + int32(16)
		v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v139)+8)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v139 + int32(16)
		return int32(3)
	}
}
func F_luaB_print(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v183 int32
	_ = v183
	var v197 int32
	_ = v197
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v257 int32
	_ = v257
	var v258 int64
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int64
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v363 int32
	_ = v363
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	v6 = m.G3
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = (v7 - v8) >> (uint(int32(4)) % 32)
	goto L1
L1:
	;
	F_lua_getfield(m, l0, int32(-10002), v6+int32(_a2183))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	v19 = m.G403
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v11 < int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v391 = m.G3
	v395 = F_luaL_error(m, l0, v391+int32(_a2304), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L2
	} else {
		goto L101
	}
L5:
	;
	v382 = F_fputc(m, int32(10), v20)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L2
	} else {
		goto L100
	}
L6:
	;
	goto L9
L7:
	;
	goto L26
L8:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	*(*int64)(unsafe.Add(mBase, uint32(v79))) = v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v84 + int32(16)
	goto L7
L9:
	;
	goto L15
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v43 = v40 + int32(-16)
	goto L8
L23:
	;
	v153 = int32(1)
	F_lua_call(m, l0, v153, v153)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L2
	} else {
		goto L39
	}
L24:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v102)))
	*(*int64)(unsafe.Add(mBase, uint32(v144))) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v144)+8)) = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v149 + int32(16)
	goto L23
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v99 = v94 + int32(0)
	v100 = m.G398
	if base.Ui32(v99) < base.Ui32(v93) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v102 = v99
	goto L29
L28:
	;
	v102 = v100
	goto L29
L29:
	;
	goto L24
L39:
	;
	v159 = F_lua_tolstring(m, l0, int32(-1), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	if v159 == int32(0) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v163 = F_fputs(m, v159, v20)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	goto L45
L43:
	;
	if v11 == int32(1) {
		goto L5
	} else {
		goto L51
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v183 + int32(-16)
	goto L43
L45:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L44
L51:
	;
	v197 = int32(2)
	goto L52
L52:
	;
	goto L56
L53:
	;
	goto L5
L54:
	;
	if v197 < int32(1) {
		goto L72
	} else {
		goto L73
	}
L55:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
	*(*int64)(unsafe.Add(mBase, uint32(v257))) = v258
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v221)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v257)+8)) = v260
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v262 + int32(16)
	goto L54
L56:
	;
	goto L62
L62:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v221 = v218 + int32(-16)
	goto L55
L70:
	;
	v330 = int32(1)
	F_lua_call(m, l0, v330, v330)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L2
	} else {
		goto L86
	}
L71:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v322 = *(*int64)(unsafe.Add(mBase, uint32(v318)))
	*(*int64)(unsafe.Add(mBase, uint32(v321))) = v322
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+8)) = v324
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v326 + int32(16)
	goto L70
L72:
	;
	if v197 < int32(-9999) {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v276 = v271 + v197<<(uint(int32(4))%32) + int32(-16)
	v277 = m.G398
	if base.Ui32(v276) < base.Ui32(v270) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v279 = v276
	goto L76
L75:
	;
	v279 = v277
	goto L76
L76:
	;
	v318 = v279
	goto L71
L77:
	;
	switch v197 + int32(10002) {
	case 0:
		goto L80
	case 1:
		goto L81
	case 2:
		goto L82
	default:
		goto L79
	}
L78:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v318 = v282 + v197<<(uint(int32(4))%32)
	goto L71
L79:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+4))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+7)))
	v306 = m.G398
	if base.Ui32(v305) < base.Ui32(int32(-10002)-v197) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v318 = l0 + int32(72)
	goto L71
L81:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v294
	v318 = l0 + int32(88)
	goto L71
L82:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v318 = v288 + int32(96)
	goto L71
L83:
	;
	v317 = v306
	goto L85
L84:
	;
	v317 = v304 + (int32(-10003)-v197)<<(uint(int32(4))%32) + int32(24)
	goto L85
L85:
	;
	v318 = v317
	goto L71
L86:
	;
	v336 = F_lua_tolstring(m, l0, int32(-1), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	if v336 == int32(0) {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	v341 = F_fputc(m, int32(9), v20)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	v343 = F_fputs(m, v336, v20)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L2
	} else {
		goto L90
	}
L90:
	;
	goto L93
L91:
	;
	if v197 != v11 {
		v197 = v197 + int32(1)
		goto L52
	} else {
		goto L99
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v363 + int32(-16)
	goto L91
L93:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L92
L99:
	;
	goto L53
L100:
	;
	return int32(0)
L101:
	;
	return v395
}
func F_luaB_rawequal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	F_luaL_checkany(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_luaL_checkany(m, l0, int32(2))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v22 = v17 + int32(0)
			v23 = m.G398
			if base.Ui32(v22) < base.Ui32(v16) {
				v25 = v22
			} else {
				v25 = v23
			}
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v75 = v70 + int32(16)
			v76 = m.G398
			if base.Ui32(v75) < base.Ui32(v69) {
				v78 = v75
			} else {
				v78 = v76
			}
			v120 = int32(0)
			v121 = m.G398
			if v25 == v121 {
				v126 = v120
			} else {
				v123 = m.G398
				if v78 == v123 {
					v126 = v120
				} else {
					v125 = F_luaO_rawequalObj(m, v25, v78)
					mBase = m.M
					v126 = v125
				}
			}
			v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v128))) = base.B2i32(v126 != int32(0))
			v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v134 + int32(16)
			return int32(1)
		}
	}
}
func F_luaB_rawget(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int64
	_ = v103
	var v107 int32
	_ = v107
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_luaL_checkany(m, l0, int32(2))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v19 = v16 + int32(32)
			if base.Ui32(v19) <= base.Ui32(v15) {
			} else {
				v23 = v15
				for {
					*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(0)
					v27 = v23 + int32(16)
					if base.Ui32(v27) < base.Ui32(v19) {
						v23 = v27
						continue
					} else {
						break
					}
					break
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v19
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v50 = v45 + int32(0)
			v51 = m.G398
			if base.Ui32(v50) < base.Ui32(v44) {
				v53 = v50
			} else {
				v53 = v51
			}
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v97 = int32(-16)
			v99 = F_luaH_get(m, v95, v96+v97)
			mBase = m.M
			v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v103 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
			*(*int64)(unsafe.Add(mBase, uint32(v100+v97))) = v103
			v107 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v100+int32(-8)))) = v107
			return int32(1)
		}
	}
}
func F_luaB_rawset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	F_luaL_checktype(m, l0, int32(1), int32(5))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_luaL_checkany(m, l0, int32(2))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_luaL_checkany(m, l0, int32(3))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v22 = v19 + int32(48)
				if base.Ui32(v22) <= base.Ui32(v18) {
				} else {
					v26 = v18
					for {
						*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(0)
						v30 = v26 + int32(16)
						if base.Ui32(v30) < base.Ui32(v22) {
							v26 = v30
							continue
						} else {
							break
						}
						break
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v22
				F_lua_rawset(m, l0, int32(1))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					return int32(1)
				}
			}
		}
	}
}
func F_luaB_setfenv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 float64
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	F_luaL_checktype(m, l0, int32(2), int32(5))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		F_getfunc(m, l0, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v23 = v18 + int32(16)
			v24 = m.G398
			if base.Ui32(v23) < base.Ui32(v17) {
				v26 = v23
			} else {
				v26 = v24
			}
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v69 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
			*(*int64)(unsafe.Add(mBase, uint32(v68))) = v69
			v71 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = v71
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73 + int32(16)
			v78 = F_lua_isnumber(m, l0, int32(1))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				if v78 == int32(0) {
					v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v306 = v303 + int32(-32)
					v343 = *(*int32)(unsafe.Add(mBase, uint32(v306)+8))
					if v343 != int32(6) {
						v350 = int32(0)
					} else {
						v346 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
						v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+6)))
						v350 = base.B2i32(v347 != int32(0))
					}
					if v350 != 0 {
						v459 = m.G3
						v463 = F_luaL_error(m, l0, v459+int32(_a2305), int32(0))
						mBase = m.M
						v464 = m.ExcPending
						if v464 != 0 {
							return int32(0)
						} else {
							v466 = int32(1)
							return v466
						}
					} else {
						v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v372 = v369 + int32(-32)
						v409 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
						switch v409 + int32(-6) {
						case 0:
							v412 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
							v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v416 = *(*int32)(unsafe.Add(mBase, uint32(v413+int32(-16))))
							*(*int32)(unsafe.Add(mBase, uint32(v412)+12)) = v416
							v434 = int32(1)
							v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v438 = *(*int32)(unsafe.Add(mBase, uint32(v435+int32(-16))))
							v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
							if v439&int32(3) == int32(0) {
								v452 = v434
							} else {
								v444 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
								v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+5)))
								if v445&int32(4) == int32(0) {
									v452 = v434
								} else {
									F_luaC_barrierf(m, l0, v444, v438)
									mBase = m.M
									v452 = v434
								}
							}
						case 1:
							v418 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
							v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v422 = *(*int32)(unsafe.Add(mBase, uint32(v419+int32(-16))))
							*(*int32)(unsafe.Add(mBase, uint32(v418)+12)) = v422
							v434 = int32(1)
							v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v438 = *(*int32)(unsafe.Add(mBase, uint32(v435+int32(-16))))
							v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
							if v439&int32(3) == int32(0) {
								v452 = v434
							} else {
								v444 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
								v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+5)))
								if v445&int32(4) == int32(0) {
									v452 = v434
								} else {
									F_luaC_barrierf(m, l0, v444, v438)
									mBase = m.M
									v452 = v434
								}
							}
						case 2:
							v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v427 = *(*int32)(unsafe.Add(mBase, uint32(v424+int32(-16))))
							v428 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
							*(*int32)(unsafe.Add(mBase, uint32(v428)+80)) = int32(5)
							*(*int32)(unsafe.Add(mBase, uint32(v428)+72)) = v427
							v434 = int32(1)
							v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v438 = *(*int32)(unsafe.Add(mBase, uint32(v435+int32(-16))))
							v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
							if v439&int32(3) == int32(0) {
								v452 = v434
							} else {
								v444 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
								v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+5)))
								if v445&int32(4) == int32(0) {
									v452 = v434
								} else {
									F_luaC_barrierf(m, l0, v444, v438)
									mBase = m.M
									v452 = v434
								}
							}
						default:
							v452 = int32(0)
						}
						v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v454 + int32(-16)
						if v452 != 0 {
							v466 = int32(1)
							return v466
						} else {
							v459 = m.G3
							v463 = F_luaL_error(m, l0, v459+int32(_a2305), int32(0))
							mBase = m.M
							v464 = m.ExcPending
							if v464 != 0 {
								return int32(0)
							} else {
								v466 = int32(1)
								return v466
							}
						}
					}
				} else {
					v83 = F_lua_tonumber(m, l0, int32(1))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						if base.F64_ne(v83, float64(0)) != 0 {
							v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v306 = v303 + int32(-32)
							v343 = *(*int32)(unsafe.Add(mBase, uint32(v306)+8))
							if v343 != int32(6) {
								v350 = int32(0)
							} else {
								v346 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
								v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+6)))
								v350 = base.B2i32(v347 != int32(0))
							}
							if v350 != 0 {
								v459 = m.G3
								v463 = F_luaL_error(m, l0, v459+int32(_a2305), int32(0))
								mBase = m.M
								v464 = m.ExcPending
								if v464 != 0 {
									return int32(0)
								} else {
									v466 = int32(1)
									return v466
								}
							} else {
								v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v372 = v369 + int32(-32)
								v409 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
								switch v409 + int32(-6) {
								case 0:
									v412 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
									v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v416 = *(*int32)(unsafe.Add(mBase, uint32(v413+int32(-16))))
									*(*int32)(unsafe.Add(mBase, uint32(v412)+12)) = v416
									v434 = int32(1)
									v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v438 = *(*int32)(unsafe.Add(mBase, uint32(v435+int32(-16))))
									v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
									if v439&int32(3) == int32(0) {
										v452 = v434
									} else {
										v444 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
										v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+5)))
										if v445&int32(4) == int32(0) {
											v452 = v434
										} else {
											F_luaC_barrierf(m, l0, v444, v438)
											mBase = m.M
											v452 = v434
										}
									}
								case 1:
									v418 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
									v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v422 = *(*int32)(unsafe.Add(mBase, uint32(v419+int32(-16))))
									*(*int32)(unsafe.Add(mBase, uint32(v418)+12)) = v422
									v434 = int32(1)
									v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v438 = *(*int32)(unsafe.Add(mBase, uint32(v435+int32(-16))))
									v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
									if v439&int32(3) == int32(0) {
										v452 = v434
									} else {
										v444 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
										v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+5)))
										if v445&int32(4) == int32(0) {
											v452 = v434
										} else {
											F_luaC_barrierf(m, l0, v444, v438)
											mBase = m.M
											v452 = v434
										}
									}
								case 2:
									v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v427 = *(*int32)(unsafe.Add(mBase, uint32(v424+int32(-16))))
									v428 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
									*(*int32)(unsafe.Add(mBase, uint32(v428)+80)) = int32(5)
									*(*int32)(unsafe.Add(mBase, uint32(v428)+72)) = v427
									v434 = int32(1)
									v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v438 = *(*int32)(unsafe.Add(mBase, uint32(v435+int32(-16))))
									v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+5)))
									if v439&int32(3) == int32(0) {
										v452 = v434
									} else {
										v444 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
										v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+5)))
										if v445&int32(4) == int32(0) {
											v452 = v434
										} else {
											F_luaC_barrierf(m, l0, v444, v438)
											mBase = m.M
											v452 = v434
										}
									}
								default:
									v452 = int32(0)
								}
								v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v454 + int32(-16)
								if v452 != 0 {
									v466 = int32(1)
									return v466
								} else {
									v459 = m.G3
									v463 = F_luaL_error(m, l0, v459+int32(_a2305), int32(0))
									mBase = m.M
									v464 = m.ExcPending
									if v464 != 0 {
										return int32(0)
									} else {
										v466 = int32(1)
										return v466
									}
								}
							}
						} else {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v88))) = l0
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92 + int32(16)
							v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v119 = v116 + int32(-32)
							v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if base.Ui32(v155) <= base.Ui32(v119) {
								v172 = v155
							} else {
								v158 = v155
								for {
									v162 = v158 + int32(-16)
									v163 = *(*int64)(unsafe.Add(mBase, uint32(v162)))
									*(*int64)(unsafe.Add(mBase, uint32(v158))) = v163
									v167 = *(*int32)(unsafe.Add(mBase, uint32(v158+int32(-8))))
									*(*int32)(unsafe.Add(mBase, uint32(v158)+8)) = v167
									if base.Ui32(v119) < base.Ui32(v162) {
										v158 = v162
										continue
									} else {
										break
									}
									break
								}
								v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v172 = v170
							}
							v175 = *(*int64)(unsafe.Add(mBase, uint32(v172)))
							*(*int64)(unsafe.Add(mBase, uint32(v119))) = v175
							v177 = *(*int32)(unsafe.Add(mBase, uint32(v172)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v177
							v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v199 = v196 + int32(-32)
							v236 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
							switch v236 + int32(-6) {
							case 0:
								v239 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
								v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v243 = *(*int32)(unsafe.Add(mBase, uint32(v240+int32(-16))))
								*(*int32)(unsafe.Add(mBase, uint32(v239)+12)) = v243
								v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v265 = *(*int32)(unsafe.Add(mBase, uint32(v262+int32(-16))))
								v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+5)))
								if v266&int32(3) == int32(0) {
								} else {
									v271 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
									v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+5)))
									if v272&int32(4) == int32(0) {
									} else {
										F_luaC_barrierf(m, l0, v271, v265)
										mBase = m.M
									}
								}
							case 1:
								v245 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
								v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v249 = *(*int32)(unsafe.Add(mBase, uint32(v246+int32(-16))))
								*(*int32)(unsafe.Add(mBase, uint32(v245)+12)) = v249
								v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v265 = *(*int32)(unsafe.Add(mBase, uint32(v262+int32(-16))))
								v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+5)))
								if v266&int32(3) == int32(0) {
								} else {
									v271 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
									v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+5)))
									if v272&int32(4) == int32(0) {
									} else {
										F_luaC_barrierf(m, l0, v271, v265)
										mBase = m.M
									}
								}
							case 2:
								v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v254 = *(*int32)(unsafe.Add(mBase, uint32(v251+int32(-16))))
								v255 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
								*(*int32)(unsafe.Add(mBase, uint32(v255)+80)) = int32(5)
								*(*int32)(unsafe.Add(mBase, uint32(v255)+72)) = v254
								v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v265 = *(*int32)(unsafe.Add(mBase, uint32(v262+int32(-16))))
								v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+5)))
								if v266&int32(3) == int32(0) {
								} else {
									v271 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
									v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+5)))
									if v272&int32(4) == int32(0) {
									} else {
										F_luaC_barrierf(m, l0, v271, v265)
										mBase = m.M
									}
								}
							default:
							}
							v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v281 + int32(-16)
							return int32(0)
						}
					}
				}
			}
		}
	}
}
func F_luaB_tonumber(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_luaL_optinteger(m, l0, int32(2), int32(10))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != int32(10) {
			v40 = F_luaL_checklstring(m, l0, int32(1), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				if base.Ui32(v12+int32(-2)) < base.Ui32(int32(35)) {
					v55 = F_strtox_2(m, v40, v8+int32(12), v12, int64(4294967295))
					mBase = m.M
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					if v40 == v57 {
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v89 + int32(16)
					} else {
						v61 = v57
						for {
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
							if base.Ui32(v64+int32(-9)) < base.Ui32(int32(5)) {
								v61 = v61 + int32(1)
								continue
							} else {
							}
							if v64 == int32(32) {
								v61 = v61 + int32(1)
								continue
							} else {
								break
							}
							break
						}
						if v64 != 0 {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v89 + int32(16)
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = int32(3)
							*(*float64)(unsafe.Add(mBase, uint32(v73))) = base.F64_convert_i32_u(base.I32_wrap_i64(v55))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v77 + int32(16)
						}
					}
					m.G0 = v8 + int32(16)
					return int32(1)
				} else {
					v47 = m.G3
					v50 = F_luaL_argerror(m, l0, int32(2), v47+int32(_a2306))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v55 = F_strtox_2(m, v40, v8+int32(12), v12, int64(4294967295))
						mBase = m.M
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						if v40 == v57 {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v89 + int32(16)
						} else {
							v61 = v57
							for {
								v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
								if base.Ui32(v64+int32(-9)) < base.Ui32(int32(5)) {
									v61 = v61 + int32(1)
									continue
								} else {
								}
								if v64 == int32(32) {
									v61 = v61 + int32(1)
									continue
								} else {
									break
								}
								break
							}
							if v64 != 0 {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v89 + int32(16)
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = int32(3)
								*(*float64)(unsafe.Add(mBase, uint32(v73))) = base.F64_convert_i32_u(base.I32_wrap_i64(v55))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v77 + int32(16)
							}
						}
						m.G0 = v8 + int32(16)
						return int32(1)
					}
				}
			}
		} else {
			F_luaL_checkany(m, l0, int32(1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v22 = F_lua_isnumber(m, l0, int32(1))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					if v22 == int32(0) {
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v89 + int32(16)
						m.G0 = v8 + int32(16)
						return int32(1)
					} else {
						v27 = F_lua_tonumber(m, l0, int32(1))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(3)
							*(*float64)(unsafe.Add(mBase, uint32(v30))) = v27
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34 + int32(16)
							m.G0 = v8 + int32(16)
							return int32(1)
						}
					}
				}
			}
		}
	}
}
func F_luaB_tostring(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_luaL_checkany(m, l0, int32(1))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = m.G3
		v17 = F_luaL_callmeta(m, l0, int32(1), v14+int32(_a2307))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v17 != 0 {
				m.G0 = v6 + int32(16)
				return int32(1)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v27 = v22 + int32(0)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v27) < base.Ui32(v28) {
					v70 = m.G398
					if v27 != v70 {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
						v76 = v73
					} else {
						v76 = int32(-1)
					}
				} else {
					v76 = int32(-1)
				}
				switch v76 {
				case 0:
					v218 = m.G3
					F_lua_pushlstring(m, l0, v218+int32(_a1724), int32(3))
					mBase = m.M
					v223 = m.ExcPending
					if v223 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(16)
						return int32(1)
					}
				case 1:
					v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v158 = v153 + int32(0)
					v159 = m.G398
					if base.Ui32(v158) < base.Ui32(v152) {
						v161 = v158
					} else {
						v161 = v159
					}
					v203 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
					switch v203 {
					case 0:
						v208 = v203
						v210 = v208
					case 1:
						v204 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
						v210 = base.B2i32(v204 != int32(0))
					default:
						v208 = int32(1)
						v210 = v208
					}
					v211 = m.G3
					if v210 != 0 {
						v214 = int32(_a1725)
					} else {
						v214 = int32(_a1726)
					}
					F_lua_pushstring(m, l0, v211+v214)
					mBase = m.M
					v217 = m.ExcPending
					if v217 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(16)
						return int32(1)
					}
				default:
					v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v232 = v227 + int32(0)
					v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if base.Ui32(v232) < base.Ui32(v233) {
						v275 = m.G398
						if v232 != v275 {
							v278 = *(*int32)(unsafe.Add(mBase, uint32(v232)+8))
							v281 = v278
						} else {
							v281 = int32(-1)
						}
					} else {
						v281 = int32(-1)
					}
					v283 = m.G3
					if v281 != int32(-1) {
						v288 = m.G399
						v292 = *(*int32)(unsafe.Add(mBase, uint32(v288+v281<<(uint(int32(2))%32))))
						v293 = v292
					} else {
						v293 = v283 + int32(_a2282)
					}
					v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v305 = v300 + int32(0)
					v306 = m.G398
					if base.Ui32(v305) < base.Ui32(v299) {
						v308 = v305
					} else {
						v308 = v306
					}
					v349 = int32(0)
					v350 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
					switch v350 + int32(-2) {
					case 0, 5:
						v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v362 = v357 + int32(0)
						v363 = m.G398
						if base.Ui32(v362) < base.Ui32(v356) {
							v365 = v362
						} else {
							v365 = v363
						}
						v407 = *(*int32)(unsafe.Add(mBase, uint32(v365)+8))
						switch v407 + int32(-2) {
						case 0:
							v413 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
							v417 = v413
							v422 = v417
						default:
							v417 = v349
							v422 = v417
						case 5:
							v410 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
							v422 = v410 + int32(24)
						}
					default:
						v417 = v349
						v422 = v417
					case 3, 4, 6:
						v353 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
						v422 = v353
					}
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v422
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v293
					v425 = m.G3
					v428 = F_lua_pushfstring(m, l0, v425+int32(_a2308), v6)
					mBase = m.M
					v429 = m.ExcPending
					if v429 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(16)
						return int32(1)
					}
				case 3:
					v79 = F_lua_tolstring(m, l0, int32(1), int32(0))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						F_lua_pushstring(m, l0, v79)
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							m.G0 = v6 + int32(16)
							return int32(1)
						}
					}
				case 4:
					v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v94 = v89 + int32(0)
					v95 = m.G398
					if base.Ui32(v94) < base.Ui32(v88) {
						v97 = v94
					} else {
						v97 = v95
					}
					v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v140 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
					*(*int64)(unsafe.Add(mBase, uint32(v139))) = v140
					v142 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v139)+8)) = v142
					v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v144 + int32(16)
					m.G0 = v6 + int32(16)
					return int32(1)
				}
			}
		}
	}
}
func F_luaB_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	F_luaL_checkany(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v15 = v10 + int32(0)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v15) < base.Ui32(v16) {
			v58 = m.G398
			if v15 != v58 {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v64 = v61
			} else {
				v64 = int32(-1)
			}
		} else {
			v64 = int32(-1)
		}
		v66 = m.G3
		if v64 != int32(-1) {
			v71 = m.G399
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v64<<(uint(int32(2))%32))))
			v76 = v75
		} else {
			v76 = v66 + int32(_a2282)
		}
		F_lua_pushstring(m, l0, v76)
		mBase = m.M
		v78 = m.ExcPending
		if v78 != 0 {
			return int32(0)
		} else {
			return int32(1)
		}
	}
}
func F_luaB_xpcall(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	F_luaL_checkany(m, l0, int32(2))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v15 = v12 + int32(32)
		if base.Ui32(v15) <= base.Ui32(v11) {
		} else {
			v19 = v11
			for {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(0)
				v23 = v19 + int32(16)
				if base.Ui32(v23) < base.Ui32(v15) {
					v19 = v23
					continue
				} else {
					break
				}
				break
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v15
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v46 = v41 + int32(0)
		v47 = m.G398
		if base.Ui32(v46) < base.Ui32(v40) {
			v49 = v46
		} else {
			v49 = v47
		}
		v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v91) <= base.Ui32(v49) {
			v108 = v91
		} else {
			v94 = v91
			for {
				v98 = v94 + int32(-16)
				v99 = *(*int64)(unsafe.Add(mBase, uint32(v98)))
				*(*int64)(unsafe.Add(mBase, uint32(v94))) = v99
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v94+int32(-8))))
				*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v103
				if base.Ui32(v49) < base.Ui32(v98) {
					v94 = v98
					continue
				} else {
					break
				}
				break
			}
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v108 = v106
		}
		v111 = *(*int64)(unsafe.Add(mBase, uint32(v108)))
		*(*int64)(unsafe.Add(mBase, uint32(v49))) = v111
		v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v113
		v118 = F_lua_pcall(m, l0, int32(0), int32(-1), int32(1))
		mBase = m.M
		v119 = m.ExcPending
		if v119 != 0 {
			return int32(0)
		} else {
			v120 = int32(0)
			v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v123)+8)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v123))) = base.B2i32(base.B2i32(v118 == v120) != v120)
			v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v129 + int32(16)
			F_lua_replace(m, l0, int32(1))
			mBase = m.M
			v135 = m.ExcPending
			if v135 != 0 {
				return int32(0)
			} else {
				v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				return (v136 - v137) >> (uint(int32(4)) % 32)
			}
		}
	}
}
func F_luaB_yield(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = F_lua_yield(m, l0, (v2-v3)>>(uint(int32(4))%32))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
