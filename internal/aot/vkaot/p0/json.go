package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_json_append_number(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_lua_tonumber(m, l0, l3)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v17 = base.I64_reinterpret_f64(v13) & int64(9223372036854775807)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1320))
		switch v18 {
		case 0:
			if base.Ui64(v17) < base.Ui64(int64(9218868437227405312)) {
				v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v156 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				if base.Ui32(v155-v156) < base.Ui32(int32(-32)) {
					v165 = v155
					v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1324))
					v169 = F_fpconv_g_fmt(m, v166+v165, v13, v168)
					mBase = m.M
					v170 = m.ExcPending
					if v170 != 0 {
						return
					} else {
						v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v169 + v171
						m.G0 = v11 + int32(16)
						return
					}
				} else {
					F_strbuf_resize(m, l2, v155+int32(32))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return
					} else {
						v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						v165 = v164
						v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1324))
						v169 = F_fpconv_g_fmt(m, v166+v165, v13, v168)
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return
						} else {
							v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v169 + v171
							m.G0 = v11 + int32(16)
							return
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1328))
				if v21 != 0 {
					v24 = m.G3
					if l3 < int32(1) {
						if l3 < int32(-9999) {
							switch l3 + int32(10002) {
							case 0:
								v74 = l0 + int32(72)
								v75 = m.G398
								if v74 != v75 {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
									v81 = v78
								} else {
									v81 = int32(-1)
								}
							case 1:
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v47
								v74 = l0 + int32(88)
								v75 = m.G398
								if v74 != v75 {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
									v81 = v78
								} else {
									v81 = int32(-1)
								}
							case 2:
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v74 = v70 + int32(96)
								v75 = m.G398
								if v74 != v75 {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
									v81 = v78
								} else {
									v81 = int32(-1)
								}
							default:
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
								v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+7)))
								if base.Ui32(v60) < base.Ui32(int32(-10002)-l3) {
									v81 = int32(-1)
								} else {
									v74 = v59 + (int32(-10003)-l3)<<(uint(int32(4))%32) + int32(24)
									v75 = m.G398
									if v74 != v75 {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
										v81 = v78
									} else {
										v81 = int32(-1)
									}
								}
							}
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v74 = v38 + l3<<(uint(int32(4))%32)
							v75 = m.G398
							if v74 != v75 {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
								v81 = v78
							} else {
								v81 = int32(-1)
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v32 = v27 + l3<<(uint(int32(4))%32) + int32(-16)
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if base.Ui32(v32) < base.Ui32(v33) {
							v74 = v32
							v75 = m.G398
							if v74 != v75 {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
								v81 = v78
							} else {
								v81 = int32(-1)
							}
						} else {
							v81 = int32(-1)
						}
					}
					v83 = m.G3
					if v81 != int32(-1) {
						v88 = m.G399
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v81<<(uint(int32(2))%32))))
						v93 = v92
					} else {
						v93 = v83 + int32(_a2282)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v24 + int32(_a2335)
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v93
					v100 = F_luaL_error(m, l0, v24+int32(_a2336), v11)
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return
					} else {
						v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						v156 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
						if base.Ui32(v155-v156) < base.Ui32(int32(-32)) {
							v165 = v155
							v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1324))
							v169 = F_fpconv_g_fmt(m, v166+v165, v13, v168)
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
								return
							} else {
								v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v169 + v171
								m.G0 = v11 + int32(16)
								return
							}
						} else {
							F_strbuf_resize(m, l2, v155+int32(32))
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return
							} else {
								v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								v165 = v164
								v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1324))
								v169 = F_fpconv_g_fmt(m, v166+v165, v13, v168)
								mBase = m.M
								v170 = m.ExcPending
								if v170 != 0 {
									return
								} else {
									v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v169 + v171
									m.G0 = v11 + int32(16)
									return
								}
							}
						}
					}
				} else {
					F_strbuf_free(m, l2)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = m.G3
						if l3 < int32(1) {
							if l3 < int32(-9999) {
								switch l3 + int32(10002) {
								case 0:
									v74 = l0 + int32(72)
									v75 = m.G398
									if v74 != v75 {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
										v81 = v78
									} else {
										v81 = int32(-1)
									}
								case 1:
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v47
									v74 = l0 + int32(88)
									v75 = m.G398
									if v74 != v75 {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
										v81 = v78
									} else {
										v81 = int32(-1)
									}
								case 2:
									v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v74 = v70 + int32(96)
									v75 = m.G398
									if v74 != v75 {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
										v81 = v78
									} else {
										v81 = int32(-1)
									}
								default:
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+7)))
									if base.Ui32(v60) < base.Ui32(int32(-10002)-l3) {
										v81 = int32(-1)
									} else {
										v74 = v59 + (int32(-10003)-l3)<<(uint(int32(4))%32) + int32(24)
										v75 = m.G398
										if v74 != v75 {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
											v81 = v78
										} else {
											v81 = int32(-1)
										}
									}
								}
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v74 = v38 + l3<<(uint(int32(4))%32)
								v75 = m.G398
								if v74 != v75 {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
									v81 = v78
								} else {
									v81 = int32(-1)
								}
							}
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v32 = v27 + l3<<(uint(int32(4))%32) + int32(-16)
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if base.Ui32(v32) < base.Ui32(v33) {
								v74 = v32
								v75 = m.G398
								if v74 != v75 {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
									v81 = v78
								} else {
									v81 = int32(-1)
								}
							} else {
								v81 = int32(-1)
							}
						}
						v83 = m.G3
						if v81 != int32(-1) {
							v88 = m.G399
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v81<<(uint(int32(2))%32))))
							v93 = v92
						} else {
							v93 = v83 + int32(_a2282)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v24 + int32(_a2335)
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v93
						v100 = F_luaL_error(m, l0, v24+int32(_a2336), v11)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							v156 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							if base.Ui32(v155-v156) < base.Ui32(int32(-32)) {
								v165 = v155
								v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1324))
								v169 = F_fpconv_g_fmt(m, v166+v165, v13, v168)
								mBase = m.M
								v170 = m.ExcPending
								if v170 != 0 {
									return
								} else {
									v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v169 + v171
									m.G0 = v11 + int32(16)
									return
								}
							} else {
								F_strbuf_resize(m, l2, v155+int32(32))
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
									return
								} else {
									v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									v165 = v164
									v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1324))
									v169 = F_fpconv_g_fmt(m, v166+v165, v13, v168)
									mBase = m.M
									v170 = m.ExcPending
									if v170 != 0 {
										return
									} else {
										v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v169 + v171
										m.G0 = v11 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			}
		case 1:
			if base.Ui64(v17) < base.Ui64(int64(9218868437227405313)) {
				v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v156 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				if base.Ui32(v155-v156) < base.Ui32(int32(-32)) {
					v165 = v155
					v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1324))
					v169 = F_fpconv_g_fmt(m, v166+v165, v13, v168)
					mBase = m.M
					v170 = m.ExcPending
					if v170 != 0 {
						return
					} else {
						v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v169 + v171
						m.G0 = v11 + int32(16)
						return
					}
				} else {
					F_strbuf_resize(m, l2, v155+int32(32))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return
					} else {
						v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						v165 = v164
						v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1324))
						v169 = F_fpconv_g_fmt(m, v166+v165, v13, v168)
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return
						} else {
							v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v169 + v171
							m.G0 = v11 + int32(16)
							return
						}
					}
				}
			} else {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				if base.Ui32(v104-v105) < base.Ui32(int32(-3)) {
					v114 = v104
					v115 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v116 = v115 + v114
					v117 = m.G3
					v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+uint32(_consts[1109]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v120)
					v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+uint32(_consts[1108]))))
					*(*uint8)(unsafe.Add(mBase, uint32(v116+int32(2)))) = uint8(v126)
					v128 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v128 + int32(3)
					m.G0 = v11 + int32(16)
					return
				} else {
					F_strbuf_resize(m, l2, v104+int32(3))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						v114 = v113
						v115 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v116 = v115 + v114
						v117 = m.G3
						v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+uint32(_consts[1109]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v120)
						v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+uint32(_consts[1108]))))
						*(*uint8)(unsafe.Add(mBase, uint32(v116+int32(2)))) = uint8(v126)
						v128 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v128 + int32(3)
						m.G0 = v11 + int32(16)
						return
					}
				}
			}
		default:
			if base.Ui64(v17) < base.Ui64(int64(9218868437227405312)) {
				v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v156 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				if base.Ui32(v155-v156) < base.Ui32(int32(-32)) {
					v165 = v155
					v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1324))
					v169 = F_fpconv_g_fmt(m, v166+v165, v13, v168)
					mBase = m.M
					v170 = m.ExcPending
					if v170 != 0 {
						return
					} else {
						v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v169 + v171
						m.G0 = v11 + int32(16)
						return
					}
				} else {
					F_strbuf_resize(m, l2, v155+int32(32))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return
					} else {
						v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						v165 = v164
						v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+1324))
						v169 = F_fpconv_g_fmt(m, v166+v165, v13, v168)
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return
						} else {
							v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v169 + v171
							m.G0 = v11 + int32(16)
							return
						}
					}
				}
			} else {
				v134 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v135 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				if base.Ui32(v134-v135) < base.Ui32(int32(-4)) {
					v144 = v134
					v145 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					*(*int32)(unsafe.Add(mBase, uint32(v145+v144))) = int32(1819047278)
					v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v149 + int32(4)
					m.G0 = v11 + int32(16)
					return
				} else {
					F_strbuf_resize(m, l2, v134+int32(4))
					mBase = m.M
					v142 = m.ExcPending
					if v142 != 0 {
						return
					} else {
						v143 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						v144 = v143
						v145 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						*(*int32)(unsafe.Add(mBase, uint32(v145+v144))) = int32(1819047278)
						v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v149 + int32(4)
						m.G0 = v11 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_json_decode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	v5 = m.G0
	v7 = v5 - int32(64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L2
L1:
	;
	goto L9
L2:
	;
	if (v9-v10)>>(uint(int32(4))%32) == int32(1) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = m.G3
	v20 = F_luaL_argerror(m, l0, int32(1), v17+int32(_a2333))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L1
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v89
	v100 = F_luaL_checklstring(m, l0, int32(1), v7+int32(12))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L28
	}
L7:
	;
	if v89 != 0 {
		goto L6
	} else {
		goto L26
	}
L8:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	switch v80 + int32(-2) {
	case 0:
		goto L24
	default:
		v87 = int32(0)
		goto L23
	case 5:
		goto L25
	}
L9:
	;
	goto L14
L14:
	;
	switch int32(-1) {
	case 0:
		goto L17
	case 1:
		goto L18
	case 2:
		goto L19
	default:
		goto L16
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+7)))
	v64 = m.G398
	if base.Ui32(v63) < base.Ui32(int32(1)) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v77 = l0 + int32(72)
	goto L8
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v52
	v77 = l0 + int32(88)
	goto L8
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v77 = v46 + int32(96)
	goto L8
L20:
	;
	v75 = v64
	goto L22
L21:
	;
	v75 = v62 + int32(24)
	goto L22
L22:
	;
	v77 = v75
	goto L8
L23:
	;
	v89 = v87
	goto L7
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v87 = v86
	goto L23
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v89 = v83 + int32(24)
	goto L7
L26:
	;
	v90 = m.G3
	v94 = F_luaL_error(m, l0, v90+int32(_a2332), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	goto L6
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v100
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if base.Ui32(v106) < base.Ui32(int32(2)) {
		v120 = v106
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v121 = F_strbuf_new(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L35
	}
L30:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v109 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v113 = m.G3
	v117 = F_luaL_error(m, l0, v113+int32(_a2331), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	if v112 != 0 {
		v120 = v106
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v120 = v119
	goto L29
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v121
	F_json_next_token(m, v7+int32(44), v7+int32(16))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_json_process_value(m, l0, v7+int32(44), v7+int32(16))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_json_next_token(m, v7+int32(44), v7+int32(16))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v142 == int32(10) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v7)+52))
	F_strbuf_free(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L46
	}
L40:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v7)+52))
	F_strbuf_free(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v150 = m.G3
	if v142 == int32(12) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v158 = v7 + int32(24)
	goto L44
L43:
	;
	v158 = v150 + int32(_a2328) + v142<<(uint(int32(2))%32)
	goto L44
L44:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v150 + int32(_a2329)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v159
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v164 + int32(1)
	v170 = F_luaL_error(m, l0, v150+int32(_a2330), v7)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	goto L39
L46:
	;
	m.G0 = v7 + int32(64)
	return int32(1)
}
func F_json_enum_option(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	v5 = m.G3
	if l2 != 0 {
		v8 = l2
	} else {
		v8 = v5 + int32(_a2334)
	}
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = v12 + int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v17) < base.Ui32(v18) {
		v60 = m.G398
		if v17 != v60 {
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
			v66 = v63
		} else {
			v66 = int32(-1)
		}
	} else {
		v66 = int32(-1)
	}
	if v66 == int32(0) {
		v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v199 = v198
		if base.Ui32(int32(1)) < base.Ui32(v199) {
			v216 = *(*int32)(unsafe.Add(mBase, uint32(v8+v199<<(uint(int32(2))%32))))
			F_lua_pushstring(m, l0, v216)
			mBase = m.M
			v218 = m.ExcPending
			if v218 != 0 {
				return
			} else {
				return
			}
		} else {
			v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v203)+8)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v203))) = base.B2i32(v199 != int32(0))
			v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v209 + int32(16)
			return
		}
	} else {
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v77 = v72 + int32(0)
		v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v77) < base.Ui32(v78) {
			v120 = m.G398
			if v77 != v120 {
				v123 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
				v126 = v123
			} else {
				v126 = int32(-1)
			}
		} else {
			v126 = int32(-1)
		}
		if v126 != int32(1) {
			v195 = F_luaL_checkoption(m, l0, int32(1), int32(0), v8)
			mBase = m.M
			v196 = m.ExcPending
			if v196 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v195
				v199 = v195
				if base.Ui32(int32(1)) < base.Ui32(v199) {
					v216 = *(*int32)(unsafe.Add(mBase, uint32(v8+v199<<(uint(int32(2))%32))))
					F_lua_pushstring(m, l0, v216)
					mBase = m.M
					v218 = m.ExcPending
					if v218 != 0 {
						return
					} else {
						return
					}
				} else {
					v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v203)+8)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v203))) = base.B2i32(v199 != int32(0))
					v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v209 + int32(16)
					return
				}
			}
		} else {
			v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v139 = v134 + int32(0)
			v140 = m.G398
			if base.Ui32(v139) < base.Ui32(v133) {
				v142 = v139
			} else {
				v142 = v140
			}
			v184 = *(*int32)(unsafe.Add(mBase, uint32(v142)+8))
			switch v184 {
			case 0:
				v189 = v184
				v191 = v189
			case 1:
				v185 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
				v191 = base.B2i32(v185 != int32(0))
			default:
				v189 = int32(1)
				v191 = v189
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v191
			v199 = v191
			if base.Ui32(int32(1)) < base.Ui32(v199) {
				v216 = *(*int32)(unsafe.Add(mBase, uint32(v8+v199<<(uint(int32(2))%32))))
				F_lua_pushstring(m, l0, v216)
				mBase = m.M
				v218 = m.ExcPending
				if v218 != 0 {
					return
				} else {
					return
				}
			} else {
				v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v203)+8)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v203))) = base.B2i32(v199 != int32(0))
				v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v209 + int32(16)
				return
			}
		}
	}
}
