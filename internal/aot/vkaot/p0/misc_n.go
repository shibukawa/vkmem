package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_nanosleep(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	v3 = int32(0)
	v6 = F___clock_nanosleep(m, v3, v3, l0, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = v3 - v6
		if base.Ui32(v10) < base.Ui32(int32(-4095)) {
			v18 = v10
		} else {
			v13 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(0) - v10
			v18 = int32(-1)
		}
		return v18
	}
}
func F_new_localvar(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+50)))
	if l2+v24 < int32(200) {
		v61 = v23
		v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+56))
		v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+48)))
		if v65 <= v66 {
			v69 = m.G3
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
			v78 = F_luaM_growaux_(m, v70, v71, v64+int32(56), int32(12), int32(32767), v69+int32(_a2266))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = v78
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v64)+56))
				if v81 <= v65 {
					v209 = v78
				} else {
					v85 = (v81 - v65) & int32(7)
					if v85 == int32(0) {
						v124 = v65
					} else {
						v95 = int32(0)
						v96 = v65
						for {
							*(*int32)(unsafe.Add(mBase, uint32(v78+v96*int32(12)))) = int32(0)
							v112 = int32(1)
							v113 = v96 + v112
							v115 = v95 + v112
							if v115 != v85 {
								v95 = v115
								v96 = v113
								continue
							} else {
								break
							}
							break
						}
						v124 = v113
					}
					if base.Ui32(int32(-8)) < base.Ui32(v65-v81) {
						v209 = v78
					} else {
						v159 = v124
						for {
							v171 = v159 * int32(12)
							v173 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v78+v171))) = v173
							*(*int32)(unsafe.Add(mBase, uint32(v78+int32(12)+v171))) = v173
							*(*int32)(unsafe.Add(mBase, uint32(v78+int32(24)+v171))) = v173
							*(*int32)(unsafe.Add(mBase, uint32(v78+int32(36)+v171))) = v173
							*(*int32)(unsafe.Add(mBase, uint32(v78+int32(48)+v171))) = v173
							*(*int32)(unsafe.Add(mBase, uint32(v78+int32(60)+v171))) = v173
							*(*int32)(unsafe.Add(mBase, uint32(v78+int32(72)+v171))) = v173
							*(*int32)(unsafe.Add(mBase, uint32(v78+int32(84)+v171))) = v173
							v197 = v159 + int32(8)
							if v197 != v81 {
								v159 = v197
								continue
							} else {
								break
							}
							break
						}
						v209 = v78
					}
				}
				v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+48)))
				*(*int32)(unsafe.Add(mBase, uint32(v209+v217*int32(12)))) = l1
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
				if v222&int32(3) == int32(0) {
					v247 = v217
				} else {
					v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
					if v227&int32(4) == int32(0) {
						v247 = v217
					} else {
						v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
						v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+21)))
						if v234 != int32(1) {
							v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+20)))
							v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
							v244 = v238&int32(3) | v241&int32(248)
							*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)) = uint8(v244)
						} else {
							F_reallymarkobject(m, v233, l1)
							mBase = m.M
						}
						v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+48)))
						v247 = v246
					}
				}
				v248 = int32(1)
				v249 = v247 + v248
				*(*uint16)(unsafe.Add(mBase, uint32(v61)+48)) = uint16(v249)
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+50)))
				*(*uint16)(unsafe.Add(mBase, uint32(v23+(l2+v251)<<(uint(v248)%32)+int32(172)))) = uint16(v247)
				m.G0 = v21 + int32(32)
				return
			}
		} else {
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
			v209 = v68
			v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+48)))
			*(*int32)(unsafe.Add(mBase, uint32(v209+v217*int32(12)))) = l1
			v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
			if v222&int32(3) == int32(0) {
				v247 = v217
			} else {
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
				if v227&int32(4) == int32(0) {
					v247 = v217
				} else {
					v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
					v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+21)))
					if v234 != int32(1) {
						v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+20)))
						v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
						v244 = v238&int32(3) | v241&int32(248)
						*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)) = uint8(v244)
					} else {
						F_reallymarkobject(m, v233, l1)
						mBase = m.M
					}
					v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+48)))
					v247 = v246
				}
			}
			v248 = int32(1)
			v249 = v247 + v248
			*(*uint16)(unsafe.Add(mBase, uint32(v61)+48)) = uint16(v249)
			v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+50)))
			*(*uint16)(unsafe.Add(mBase, uint32(v23+(l2+v251)<<(uint(v248)%32)+int32(172)))) = uint16(v247)
			m.G0 = v21 + int32(32)
			return
		}
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
		if v30 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = int32(200)
			*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v30
			v44 = m.G3
			*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v44 + int32(_a2267)
			v52 = F_luaO_pushfstring(m, v28, v44+int32(_a2268), v21+int32(16))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				v54 = v52
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
				F_luaX_lexerror(m, v56, v54, int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v61 = v60
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+56))
					v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+48)))
					if v65 <= v66 {
						v69 = m.G3
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
						v78 = F_luaM_growaux_(m, v70, v71, v64+int32(56), int32(12), int32(32767), v69+int32(_a2266))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = v78
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v64)+56))
							if v81 <= v65 {
								v209 = v78
							} else {
								v85 = (v81 - v65) & int32(7)
								if v85 == int32(0) {
									v124 = v65
								} else {
									v95 = int32(0)
									v96 = v65
									for {
										*(*int32)(unsafe.Add(mBase, uint32(v78+v96*int32(12)))) = int32(0)
										v112 = int32(1)
										v113 = v96 + v112
										v115 = v95 + v112
										if v115 != v85 {
											v95 = v115
											v96 = v113
											continue
										} else {
											break
										}
										break
									}
									v124 = v113
								}
								if base.Ui32(int32(-8)) < base.Ui32(v65-v81) {
									v209 = v78
								} else {
									v159 = v124
									for {
										v171 = v159 * int32(12)
										v173 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v78+v171))) = v173
										*(*int32)(unsafe.Add(mBase, uint32(v78+int32(12)+v171))) = v173
										*(*int32)(unsafe.Add(mBase, uint32(v78+int32(24)+v171))) = v173
										*(*int32)(unsafe.Add(mBase, uint32(v78+int32(36)+v171))) = v173
										*(*int32)(unsafe.Add(mBase, uint32(v78+int32(48)+v171))) = v173
										*(*int32)(unsafe.Add(mBase, uint32(v78+int32(60)+v171))) = v173
										*(*int32)(unsafe.Add(mBase, uint32(v78+int32(72)+v171))) = v173
										*(*int32)(unsafe.Add(mBase, uint32(v78+int32(84)+v171))) = v173
										v197 = v159 + int32(8)
										if v197 != v81 {
											v159 = v197
											continue
										} else {
											break
										}
										break
									}
									v209 = v78
								}
							}
							v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+48)))
							*(*int32)(unsafe.Add(mBase, uint32(v209+v217*int32(12)))) = l1
							v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
							if v222&int32(3) == int32(0) {
								v247 = v217
							} else {
								v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
								if v227&int32(4) == int32(0) {
									v247 = v217
								} else {
									v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
									v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+21)))
									if v234 != int32(1) {
										v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+20)))
										v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
										v244 = v238&int32(3) | v241&int32(248)
										*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)) = uint8(v244)
									} else {
										F_reallymarkobject(m, v233, l1)
										mBase = m.M
									}
									v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+48)))
									v247 = v246
								}
							}
							v248 = int32(1)
							v249 = v247 + v248
							*(*uint16)(unsafe.Add(mBase, uint32(v61)+48)) = uint16(v249)
							v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+50)))
							*(*uint16)(unsafe.Add(mBase, uint32(v23+(l2+v251)<<(uint(v248)%32)+int32(172)))) = uint16(v247)
							m.G0 = v21 + int32(32)
							return
						}
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
						v209 = v68
						v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+48)))
						*(*int32)(unsafe.Add(mBase, uint32(v209+v217*int32(12)))) = l1
						v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
						if v222&int32(3) == int32(0) {
							v247 = v217
						} else {
							v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
							if v227&int32(4) == int32(0) {
								v247 = v217
							} else {
								v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
								v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+21)))
								if v234 != int32(1) {
									v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+20)))
									v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
									v244 = v238&int32(3) | v241&int32(248)
									*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)) = uint8(v244)
								} else {
									F_reallymarkobject(m, v233, l1)
									mBase = m.M
								}
								v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+48)))
								v247 = v246
							}
						}
						v248 = int32(1)
						v249 = v247 + v248
						*(*uint16)(unsafe.Add(mBase, uint32(v61)+48)) = uint16(v249)
						v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+50)))
						*(*uint16)(unsafe.Add(mBase, uint32(v23+(l2+v251)<<(uint(v248)%32)+int32(172)))) = uint16(v247)
						m.G0 = v21 + int32(32)
						return
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(200)
			v33 = m.G3
			*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v33 + int32(_a2267)
			v39 = F_luaO_pushfstring(m, v28, v33+int32(_a2269), v21)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				v54 = v39
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
				F_luaX_lexerror(m, v56, v54, int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v61 = v60
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+56))
					v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+48)))
					if v65 <= v66 {
						v69 = m.G3
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
						v78 = F_luaM_growaux_(m, v70, v71, v64+int32(56), int32(12), int32(32767), v69+int32(_a2266))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = v78
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v64)+56))
							if v81 <= v65 {
								v209 = v78
							} else {
								v85 = (v81 - v65) & int32(7)
								if v85 == int32(0) {
									v124 = v65
								} else {
									v95 = int32(0)
									v96 = v65
									for {
										*(*int32)(unsafe.Add(mBase, uint32(v78+v96*int32(12)))) = int32(0)
										v112 = int32(1)
										v113 = v96 + v112
										v115 = v95 + v112
										if v115 != v85 {
											v95 = v115
											v96 = v113
											continue
										} else {
											break
										}
										break
									}
									v124 = v113
								}
								if base.Ui32(int32(-8)) < base.Ui32(v65-v81) {
									v209 = v78
								} else {
									v159 = v124
									for {
										v171 = v159 * int32(12)
										v173 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v78+v171))) = v173
										*(*int32)(unsafe.Add(mBase, uint32(v78+int32(12)+v171))) = v173
										*(*int32)(unsafe.Add(mBase, uint32(v78+int32(24)+v171))) = v173
										*(*int32)(unsafe.Add(mBase, uint32(v78+int32(36)+v171))) = v173
										*(*int32)(unsafe.Add(mBase, uint32(v78+int32(48)+v171))) = v173
										*(*int32)(unsafe.Add(mBase, uint32(v78+int32(60)+v171))) = v173
										*(*int32)(unsafe.Add(mBase, uint32(v78+int32(72)+v171))) = v173
										*(*int32)(unsafe.Add(mBase, uint32(v78+int32(84)+v171))) = v173
										v197 = v159 + int32(8)
										if v197 != v81 {
											v159 = v197
											continue
										} else {
											break
										}
										break
									}
									v209 = v78
								}
							}
							v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+48)))
							*(*int32)(unsafe.Add(mBase, uint32(v209+v217*int32(12)))) = l1
							v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
							if v222&int32(3) == int32(0) {
								v247 = v217
							} else {
								v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
								if v227&int32(4) == int32(0) {
									v247 = v217
								} else {
									v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
									v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+21)))
									if v234 != int32(1) {
										v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+20)))
										v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
										v244 = v238&int32(3) | v241&int32(248)
										*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)) = uint8(v244)
									} else {
										F_reallymarkobject(m, v233, l1)
										mBase = m.M
									}
									v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+48)))
									v247 = v246
								}
							}
							v248 = int32(1)
							v249 = v247 + v248
							*(*uint16)(unsafe.Add(mBase, uint32(v61)+48)) = uint16(v249)
							v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+50)))
							*(*uint16)(unsafe.Add(mBase, uint32(v23+(l2+v251)<<(uint(v248)%32)+int32(172)))) = uint16(v247)
							m.G0 = v21 + int32(32)
							return
						}
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
						v209 = v68
						v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+48)))
						*(*int32)(unsafe.Add(mBase, uint32(v209+v217*int32(12)))) = l1
						v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
						if v222&int32(3) == int32(0) {
							v247 = v217
						} else {
							v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
							if v227&int32(4) == int32(0) {
								v247 = v217
							} else {
								v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
								v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+21)))
								if v234 != int32(1) {
									v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+20)))
									v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
									v244 = v238&int32(3) | v241&int32(248)
									*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)) = uint8(v244)
								} else {
									F_reallymarkobject(m, v233, l1)
									mBase = m.M
								}
								v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+48)))
								v247 = v246
							}
						}
						v248 = int32(1)
						v249 = v247 + v248
						*(*uint16)(unsafe.Add(mBase, uint32(v61)+48)) = uint16(v249)
						v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+50)))
						*(*uint16)(unsafe.Add(mBase, uint32(v23+(l2+v251)<<(uint(v248)%32)+int32(172)))) = uint16(v247)
						m.G0 = v21 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_nolocks_localtime(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v22 int64
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v77 int64
	_ = v77
	var v82 int64
	_ = v82
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v106 int64
	_ = v106
	var v111 int64
	_ = v111
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v122 int64
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int64
	_ = v136
	var v137 int32
	_ = v137
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v149 int32
	_ = v149
	var v153 int64
	_ = v153
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int64
	_ = v170
	v8 = m.G0
	v10 = v8 - int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l3
	v13 = int32(3600)
	v16 = l1 - l2 + base.I64_extend_i32_s(l3*v13)
	v17 = int64(86400)
	v18 = base.I64_div_s(v16, v17)
	v22 = base.I64_rem_s(v18+int64(4), int64(7))
	*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)) = uint32(v22)
	v27 = base.I32_wrap_i64(v16 - v18*v17)
	v29 = base.I32_div_s(v27, v13)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v29
	v33 = v27 - v29*v13
	v35 = int32(60)
	v36 = base.I32_div_s(base.I32_extend16_s(v33), v35)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = base.I32_extend16_s(v36)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = base.I32_extend16_s(v33 - v36*v35)
	v47 = v18
	v48 = int32(1970)
	goto L1
L1:
	;
	if v48&int32(3) != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v47 = v47 - v170
	v48 = v48 + int32(1)
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v133
	v136 = int64(31)
	v137 = int32(0)
	if v47 < v136 {
		v161 = v137
		v162 = v134
		goto L19
	} else {
		goto L20
	}
L5:
	;
	v96 = base.I32_wrap_i64(v47)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v96
	v100 = int32(0)
	v101 = *(*int64)(unsafe.Add(mBase, _consts[326]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v101
	v106 = *(*int64)(unsafe.Add(mBase, _consts[327]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(32)))) = v106
	v111 = *(*int64)(unsafe.Add(mBase, _consts[328]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v111
	v116 = *(*int64)(unsafe.Add(mBase, _consts[329]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(16)))) = v116
	v119 = *(*int64)(unsafe.Add(mBase, _consts[330]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v119
	v122 = *(*int64)(unsafe.Add(mBase, _consts[331]))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v122
	v126 = base.I32_rem_u_s(v48, int32(100))
	if v126 != 0 {
		v133 = int32(29)
		v134 = v96
		goto L4
	} else {
		goto L15
	}
L6:
	;
	if int64(364) < v47 {
		v170 = int64(365)
		goto L3
	} else {
		goto L14
	}
L7:
	;
	v56 = base.I32_rem_u_s(v48, int32(100))
	if v56 != 0 {
		v62 = int64(366)
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v47 < v62 {
		goto L5
	} else {
		goto L13
	}
L9:
	;
	v60 = base.I32_rem_u_s(v48, int32(400))
	if v60 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v61 = int64(365)
	goto L12
L11:
	;
	v61 = int64(366)
	goto L12
L12:
	;
	v62 = v61
	goto L8
L13:
	;
	v170 = v62
	goto L3
L14:
	;
	v67 = base.I32_wrap_i64(v47)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v67
	v71 = int32(0)
	v72 = *(*int64)(unsafe.Add(mBase, _consts[326]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v72
	v77 = *(*int64)(unsafe.Add(mBase, _consts[327]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(32)))) = v77
	v82 = *(*int64)(unsafe.Add(mBase, _consts[328]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v82
	v87 = *(*int64)(unsafe.Add(mBase, _consts[329]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(16)))) = v87
	v90 = *(*int64)(unsafe.Add(mBase, _consts[330]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v90
	v93 = *(*int64)(unsafe.Add(mBase, _consts[331]))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v93
	v133 = int32(28)
	v134 = v67
	goto L4
L15:
	;
	v130 = base.I32_rem_u_s(v48, int32(400))
	if v130 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v131 = int32(28)
	goto L18
L17:
	;
	v131 = int32(29)
	goto L18
L18:
	;
	v133 = v131
	v134 = v96
	goto L4
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v48 + int32(-1900)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v162 + int32(1)
	return
L20:
	;
	v141 = v136
	v142 = v47
	v145 = v137
	goto L21
L21:
	;
	v147 = v142 - v141
	v149 = v145 + int32(1)
	v153 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10+v149<<(uint(int32(2))%32)))))
	if v153 <= v147 {
		v141 = v153
		v142 = v147
		v145 = v149
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v161 = v149
	v162 = base.I32_wrap_i64(v147)
	goto L19
L23:
	;
	goto L22
}
func F_notifyKeyspaceEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int64
	_ = v166
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int64
	_ = v239
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int64
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v313 int32
	_ = v313
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(_a69)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	v15 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v15 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F__serverAssert(m, int32(_a833), int32(_a834), int32(119))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L29
	} else {
		goto L92
	}
L2:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[439]))
	if v92&l0 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L3:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v13)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+204)) = v83 | int32(-2147483648)
	F_commitDeferredReplyBuffer(m, v13, int32(1))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L29
	} else {
		goto L33
	}
L4:
	;
	F_moduleNotifyKeyspaceEvent(m, l0, l1, l2, l3)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L29
	} else {
		goto L31
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[380]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	goto L6
L6:
	;
	if v20 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if l0&int32(1276) == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	if v13 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	if v29 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_moduleNotifyKeyspaceEvent(m, l0, l1, l2, l3)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L29
	} else {
		goto L30
	}
L11:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+56)))
	if v32&int32(1) == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+204))
	if v37&int32(-1073741824) != int32(1073741824) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	if v42 == int64(-1) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
	if v47&int32(1) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v73 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L16:
	;
	v53 = int32(2)
	if v47&v53 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v73 = int32(3)
	goto L15
L18:
	;
	if v47&int32(262144) != 0 {
		v70 = v53
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v47&int32(4) != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v73 = int32(1)
	goto L15
L21:
	;
	v73 = v70
	goto L15
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+216))
	if v63 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v67 = F_isImportSlotMigrationJob(m, v63)
	mBase = m.M
	if v67 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v73 = int32(0)
	goto L15
L25:
	;
	v68 = int32(4)
	goto L27
L26:
	;
	v68 = int32(5)
	goto L27
L27:
	;
	v70 = v68
	goto L21
L28:
	;
	goto L10
L29:
	;
	return
L30:
	;
	goto L3
L31:
	;
	if v13 == int32(0) {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	goto L3
L33:
	;
	goto L2
L34:
	;
	m.G0 = v10 + int32(32)
	return
L35:
	;
	if l1&int32(3) == int32(0) {
		v117 = l1
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v151 = F_createStringObject_1(m, l1, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L29
	} else {
		goto L52
	}
L37:
	;
	v150 = v142 - l1
	goto L36
L38:
	;
	v121 = v117
	goto L46
L39:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v103 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v106 = l1
	goto L42
L41:
	;
	v150 = l1 - l1
	goto L36
L42:
	;
	v110 = v106 + int32(1)
	if v110&int32(3) == int32(0) {
		v117 = v110
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v115 != 0 {
		v106 = v110
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v142 = v110
	goto L37
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v130 = int32(-2139062144)
	if (int32(16843008)-v127|v127)&v130 == v130 {
		v121 = v121 + int32(4)
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v136 = v121
	goto L49
L48:
	;
	goto L47
L49:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v140 != 0 {
		v136 = v136 + int32(1)
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v142 = v136
	goto L37
L51:
	;
	goto L50
L52:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[439]))
	if v155&int32(1) == int32(0) {
		v226 = int32(-1)
		v227 = v155
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v227&int32(2) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L54:
	;
	v163 = F_sdsnewlen(m, int32(_a835), int32(11))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L29
	} else {
		goto L55
	}
L55:
	;
	v166 = base.I64_extend_i32_s(l3)
	if v166 <= int64(-1) {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v208 = F_sdscatlen(m, v163, v10, v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L29
	} else {
		goto L65
	}
L57:
	;
	v207 = int32(0)
	goto L56
L59:
	;
	v188 = F_ull2string(m, v184, v185, v186)
	mBase = m.M
	if v188 == int32(0) {
		goto L57
	} else {
		goto L63
	}
L60:
	;
	goto L62
L61:
	;
	v184 = v10
	v185 = int32(24)
	v186 = v166
	v187 = int32(0)
	goto L59
L62:
	;
	v175 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v175)
	v179 = int32(1)
	v184 = v10 + v179
	v185 = int32(23)
	v186 = int64(0) - v166
	v187 = v179
	goto L59
L63:
	;
	v207 = v188 + v187
	goto L56
L65:
	;
	v212 = F_sdscatlen(m, v208, int32(_a836), int32(3))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L29
	} else {
		goto L66
	}
L66:
	;
	v214 = F_objectGetVal(m, l2)
	mBase = m.M
	v215 = F_sdscatsds(m, v212, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L29
	} else {
		goto L67
	}
L67:
	;
	v217 = F_createObject(m, int32(0), v215)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L29
	} else {
		goto L68
	}
L68:
	;
	v220 = F_pubsubPublishMessage(m, v217, v151, int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L29
	} else {
		goto L69
	}
L69:
	;
	F_decrRefCount(m, v217)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L29
	} else {
		goto L70
	}
L70:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _consts[439]))
	v226 = v207
	v227 = v225
	goto L53
L71:
	;
	F_decrRefCount(m, v151)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L29
	} else {
		goto L91
	}
L72:
	;
	v234 = F_sdsnewlen(m, int32(_a837), int32(11))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L29
	} else {
		goto L73
	}
L73:
	;
	if v226 != int32(-1) {
		v281 = v226
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v283 = F_sdscatlen(m, v234, v10, v281)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L29
	} else {
		goto L85
	}
L75:
	;
	v239 = base.I64_extend_i32_s(l3)
	if v239 <= int64(-1) {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v281 = v280
	goto L74
L77:
	;
	v280 = int32(0)
	goto L76
L79:
	;
	v261 = F_ull2string(m, v257, v258, v259)
	mBase = m.M
	if v261 == int32(0) {
		goto L77
	} else {
		goto L83
	}
L80:
	;
	goto L82
L81:
	;
	v257 = v10
	v258 = int32(24)
	v259 = v239
	v260 = int32(0)
	goto L79
L82:
	;
	v248 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v248)
	v252 = int32(1)
	v257 = v10 + v252
	v258 = int32(23)
	v259 = int64(0) - v239
	v260 = v252
	goto L79
L83:
	;
	v280 = v261 + v260
	goto L76
L85:
	;
	v287 = F_sdscatlen(m, v283, int32(_a836), int32(3))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L29
	} else {
		goto L86
	}
L86:
	;
	v289 = F_objectGetVal(m, v151)
	mBase = m.M
	v290 = F_sdscatsds(m, v287, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L29
	} else {
		goto L87
	}
L87:
	;
	v292 = F_createObject(m, int32(0), v290)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L29
	} else {
		goto L88
	}
L88:
	;
	v295 = F_pubsubPublishMessage(m, v292, l2, int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L29
	} else {
		goto L89
	}
L89:
	;
	F_decrRefCount(m, v292)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L29
	} else {
		goto L90
	}
L90:
	;
	goto L71
L91:
	;
	goto L34
L92:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_nullArrayCallback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 < int32(0) {
		v30 = v5
	} else {
		v11 = l0 + v6<<(uint(int32(2))%32)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		v13 = int32(1)
		v14 = v12 + v13
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1040))
		if v16 != v13 {
			v30 = v5
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v21))) = base.F64_convert_i32_u(v14)
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v25 + int32(16)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v30 = v29
		}
	}
	v34 = F_lua_checkstack(m, v30, int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return
	} else {
		if v34 != 0 {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(0)
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v47 + int32(16)
			F_processCollectionElementEnd(m, l0)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				return
			}
		} else {
			F__serverPanic_2(m, int32(891))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_numericConfigGet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v56 int64
	_ = v56
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v131 int32
	_ = v131
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int64
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v238 int64
	_ = v238
	var v240 int32
	_ = v240
	var v244 int64
	_ = v244
	var v245 int64
	_ = v245
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int64
	_ = v265
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v329 int64
	_ = v329
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int64
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v394 int64
	_ = v394
	var v396 int32
	_ = v396
	var v400 int64
	_ = v400
	var v401 int64
	_ = v401
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v421 int64
	_ = v421
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int64
	_ = v474
	var v476 int32
	_ = v476
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	v4 = m.G0
	v6 = v4 - int32(144)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	switch v9 {
	case 0:
		goto L4
	case 1, 3, 6:
		goto L6
	case 2:
		goto L12
	case 4:
		goto L11
	case 5:
		goto L10
	case 7:
		goto L9
	case 8:
		goto L8
	case 9:
		goto L7
	default:
		v38 = int64(0)
		goto L5
	}
L1:
	;
	v500 = F_sdsnew(m, v6+int32(16))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L88
	} else {
		goto L145
	}
L2:
	;
	if v151&int32(1) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v43&int32(2) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v41 = int64(*(*int32)(unsafe.Add(mBase, uint32(v40))))
	v42 = v41
	goto L3
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v151 = v39
	v152 = v38
	goto L2
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v37 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v36))))
	v38 = v37
	goto L5
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)))
	v42 = v35
	goto L3
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
	v42 = v33
	goto L3
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v31 = int64(*(*int32)(unsafe.Add(mBase, uint32(v30))))
	v42 = v31
	goto L3
L10:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v21&int32(1) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v12&int32(1) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v11 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10))))
	v42 = v11
	goto L3
L13:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	v42 = v20
	goto L3
L14:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v18 = F_getModuleNumericConfig(m, v17)
	mBase = m.M
	v42 = v18
	goto L3
L15:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	v42 = v29
	goto L3
L16:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v27 = F_getModuleUnsignedNumericConfig(m, v26)
	mBase = m.M
	v42 = v27
	goto L3
L17:
	;
	if v43&int32(16) == int32(0) {
		v151 = v43
		v152 = v42
		goto L2
	} else {
		goto L29
	}
L18:
	;
	if int64(-1) < v42 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v51 = v6 + int32(16)
	v56 = int64(0) - v42
	if v56 <= int64(-1) {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v99 = int32(37)
	*(*uint16)(unsafe.Add(mBase, uint32(v51+v97))) = uint16(v99)
	goto L1
L21:
	;
	v97 = int32(0)
	goto L20
L23:
	;
	v78 = F_ull2string(m, v74, v75, v76)
	mBase = m.M
	if v78 == int32(0) {
		goto L21
	} else {
		goto L27
	}
L24:
	;
	goto L26
L25:
	;
	v74 = v51
	v75 = int32(128)
	v76 = v56
	v77 = int32(0)
	goto L23
L26:
	;
	v65 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v65)
	v74 = v6 + int32(17)
	v75 = int32(127)
	v76 = int64(0) - v56
	v77 = int32(1)
	goto L23
L27:
	;
	v97 = v78 + v77
	goto L20
L29:
	;
	if int64(-1) < v42 {
		v151 = v43
		v152 = v42
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v108 = v6 + int32(16)
	if v42 <= int64(-1) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	goto L1
L32:
	;
	goto L31
L34:
	;
	v131 = F_ull2string(m, v127, v128, v129)
	mBase = m.M
	if v131 == int32(0) {
		goto L32
	} else {
		goto L38
	}
L35:
	;
	goto L37
L36:
	;
	v127 = v108
	v128 = int32(128)
	v129 = v42
	goto L34
L37:
	;
	v118 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v118)
	v127 = v6 + int32(17)
	v128 = int32(127)
	v129 = int64(0) - v42
	goto L34
L38:
	;
	goto L31
L40:
	;
	if v151&int32(4) == int32(0) {
		goto L86
	} else {
		goto L87
	}
L41:
	;
	v158 = v6 + int32(16)
	v160 = int32(0)
	v164 = int32(1)
	if base.Ui64(v152) < base.Ui64(int64(10)) {
		v221 = v164
		v222 = v160
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L1
L43:
	;
	v225 = v221 + v222
	if base.Ui32(int32(128)) <= base.Ui32(v225) {
		goto L74
	} else {
		goto L75
	}
L44:
	;
	v172 = v160
	v173 = v152
	goto L45
L45:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v173) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v221 = v164
	v222 = v213
	goto L43
L47:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v173) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v221 = int32(2)
	v222 = v172
	goto L43
L49:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v173) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v221 = int32(3)
	v222 = v172
	goto L43
L51:
	;
	v213 = v172 + int32(12)
	v217 = base.I64_div_u_s(v173, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v173) {
		v172 = v213
		v173 = v217
		goto L45
	} else {
		goto L73
	}
L52:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v173) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v173) {
		goto L65
	} else {
		goto L66
	}
L54:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v173) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v173) {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v173) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v173) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v221 = int32(4)
	v222 = v172
	goto L43
L59:
	;
	v194 = int32(6)
	goto L61
L60:
	;
	v194 = int32(5)
	goto L61
L61:
	;
	v221 = v194
	v222 = v172
	goto L43
L62:
	;
	v199 = int32(8)
	goto L64
L63:
	;
	v199 = int32(7)
	goto L64
L64:
	;
	v221 = v199
	v222 = v172
	goto L43
L65:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v173) {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v173) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v206 = int32(10)
	goto L69
L68:
	;
	v206 = int32(9)
	goto L69
L69:
	;
	v221 = v206
	v222 = v172
	goto L43
L70:
	;
	v211 = int32(12)
	goto L72
L71:
	;
	v211 = int32(11)
	goto L72
L72:
	;
	v221 = v211
	v222 = v172
	goto L43
L73:
	;
	goto L46
L74:
	;
	goto L85
L75:
	;
	v228 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v158+v225))) = uint8(v228)
	v231 = v225 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v152) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v267 = v158 + v264
	if base.Ui64(int64(9)) < base.Ui64(v265) {
		goto L82
	} else {
		goto L83
	}
L77:
	;
	v238 = v152
	v240 = v231
	goto L79
L78:
	;
	v264 = v231
	v265 = v152
	goto L76
L79:
	;
	v244 = int64(100)
	v245 = base.I64_div_u_s(v238, v244)
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v238-v245*v244)<<(uint(int32(1))%32))+uint32(_consts[185]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v6+int32(15)+v240))) = uint16(v254)
	v257 = v240 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v238) {
		v238 = v245
		v240 = v257
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v264 = v257
	v265 = v245
	goto L76
L81:
	;
	goto L80
L82:
	;
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v265)<<(uint(int32(1))%32))+uint32(_consts[185]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v267+int32(-1)))) = uint16(v281)
	goto L42
L83:
	;
	v272 = base.I32_wrap_i64(v265) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v267))) = uint8(v272)
	goto L42
L84:
	;
	goto L42
L85:
	;
	v285 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v285)
	goto L84
L86:
	;
	if v151&int32(8) == int32(0) {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v152
	v305 = F_snprintf(m, v6+int32(16), int32(128), int32(_a449), v6)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	return int32(0)
L89:
	;
	goto L1
L90:
	;
	v453 = v6 + int32(16)
	if v152 <= int64(-1) {
		goto L140
	} else {
		goto L141
	}
L91:
	;
	v314 = v6 + int32(16)
	v316 = int32(0)
	v320 = int32(1)
	if base.Ui64(v152) < base.Ui64(int64(10)) {
		v377 = v320
		v378 = v316
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L1
L93:
	;
	v381 = v377 + v378
	if base.Ui32(int32(128)) <= base.Ui32(v381) {
		goto L124
	} else {
		goto L125
	}
L94:
	;
	v328 = v316
	v329 = v152
	goto L95
L95:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v329) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v377 = v320
	v378 = v369
	goto L93
L97:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v329) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v377 = int32(2)
	v378 = v328
	goto L93
L99:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v329) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v377 = int32(3)
	v378 = v328
	goto L93
L101:
	;
	v369 = v328 + int32(12)
	v373 = base.I64_div_u_s(v329, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v329) {
		v328 = v369
		v329 = v373
		goto L95
	} else {
		goto L123
	}
L102:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v329) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v329) {
		goto L115
	} else {
		goto L116
	}
L104:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v329) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v329) {
		goto L112
	} else {
		goto L113
	}
L106:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v329) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v329) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v377 = int32(4)
	v378 = v328
	goto L93
L109:
	;
	v350 = int32(6)
	goto L111
L110:
	;
	v350 = int32(5)
	goto L111
L111:
	;
	v377 = v350
	v378 = v328
	goto L93
L112:
	;
	v355 = int32(8)
	goto L114
L113:
	;
	v355 = int32(7)
	goto L114
L114:
	;
	v377 = v355
	v378 = v328
	goto L93
L115:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v329) {
		goto L120
	} else {
		goto L121
	}
L116:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v329) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v362 = int32(10)
	goto L119
L118:
	;
	v362 = int32(9)
	goto L119
L119:
	;
	v377 = v362
	v378 = v328
	goto L93
L120:
	;
	v367 = int32(12)
	goto L122
L121:
	;
	v367 = int32(11)
	goto L122
L122:
	;
	v377 = v367
	v378 = v328
	goto L93
L123:
	;
	goto L96
L124:
	;
	goto L135
L125:
	;
	v384 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v314+v381))) = uint8(v384)
	v387 = v381 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v152) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v423 = v314 + v420
	if base.Ui64(int64(9)) < base.Ui64(v421) {
		goto L132
	} else {
		goto L133
	}
L127:
	;
	v394 = v152
	v396 = v387
	goto L129
L128:
	;
	v420 = v387
	v421 = v152
	goto L126
L129:
	;
	v400 = int64(100)
	v401 = base.I64_div_u_s(v394, v400)
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v394-v401*v400)<<(uint(int32(1))%32))+uint32(_consts[185]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v6+int32(15)+v396))) = uint16(v410)
	v413 = v396 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v394) {
		v394 = v401
		v396 = v413
		goto L129
	} else {
		goto L131
	}
L130:
	;
	v420 = v413
	v421 = v401
	goto L126
L131:
	;
	goto L130
L132:
	;
	v437 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v421)<<(uint(int32(1))%32))+uint32(_consts[185]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v423+int32(-1)))) = uint16(v437)
	goto L92
L133:
	;
	v428 = base.I32_wrap_i64(v421) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v423))) = uint8(v428)
	goto L92
L134:
	;
	goto L92
L135:
	;
	v441 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v314))) = uint8(v441)
	goto L134
L136:
	;
	goto L1
L137:
	;
	goto L136
L139:
	;
	v476 = F_ull2string(m, v472, v473, v474)
	mBase = m.M
	if v476 == int32(0) {
		goto L137
	} else {
		goto L143
	}
L140:
	;
	goto L142
L141:
	;
	v472 = v453
	v473 = int32(128)
	v474 = v152
	goto L139
L142:
	;
	v463 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v453))) = uint8(v463)
	v472 = v6 + int32(17)
	v473 = int32(127)
	v474 = int64(0) - v152
	goto L139
L143:
	;
	goto L136
L145:
	;
	m.G0 = v6 + int32(144)
	return v500
}
func F_numericConfigSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v105 int64
	_ = v105
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v127 int64
	_ = v127
	var v132 int64
	_ = v132
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v140 int32
	_ = v140
	var v148 int64
	_ = v148
	var v172 int64
	_ = v172
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v315 int64
	_ = v315
	var v321 int32
	_ = v321
	var v323 int64
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v337 int64
	_ = v337
	var v342 int64
	_ = v342
	var v346 int32
	_ = v346
	var v348 int64
	_ = v348
	var v350 int32
	_ = v350
	var v358 int64
	_ = v358
	var v382 int64
	_ = v382
	var v401 int32
	_ = v401
	var v410 int64
	_ = v410
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v432 int64
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v517 int64
	_ = v517
	var v523 int32
	_ = v523
	var v525 int64
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v539 int64
	_ = v539
	var v544 int64
	_ = v544
	var v548 int32
	_ = v548
	var v550 int64
	_ = v550
	var v552 int32
	_ = v552
	var v560 int64
	_ = v560
	var v571 int64
	_ = v571
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v597 int64
	_ = v597
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v617 int64
	_ = v617
	var v624 int32
	_ = v624
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v710 int64
	_ = v710
	var v716 int32
	_ = v716
	var v718 int64
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v732 int64
	_ = v732
	var v737 int64
	_ = v737
	var v741 int32
	_ = v741
	var v743 int64
	_ = v743
	var v745 int32
	_ = v745
	var v753 int64
	_ = v753
	var v777 int64
	_ = v777
	var v796 int32
	_ = v796
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v837 int64
	_ = v837
	var v843 int64
	_ = v843
	var v844 int32
	_ = v844
	var v853 int64
	_ = v853
	var v854 int64
	_ = v854
	var v857 int32
	_ = v857
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v880 int64
	_ = v880
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int64
	_ = v890
	var v891 int64
	_ = v891
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v921 int64
	_ = v921
	var v922 int32
	_ = v922
	var v923 int64
	_ = v923
	var v924 int32
	_ = v924
	var v925 int64
	_ = v925
	var v926 int32
	_ = v926
	var v927 int64
	_ = v927
	var v928 int32
	_ = v928
	var v933 int32
	_ = v933
	var v934 int64
	_ = v934
	var v935 int32
	_ = v935
	var v936 int64
	_ = v936
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v943 int64
	_ = v943
	var v944 int32
	_ = v944
	var v945 int64
	_ = v945
	var v946 int32
	_ = v946
	var v947 int64
	_ = v947
	var v948 int32
	_ = v948
	var v949 int64
	_ = v949
	var v950 int32
	_ = v950
	var v951 int64
	_ = v951
	var v952 int32
	_ = v952
	var v953 int64
	_ = v953
	var v954 int64
	_ = v954
	var v955 int64
	_ = v955
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	v12 = m.G0
	v14 = v12 - int32(64)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v17&int32(1) == int32(0) {
		v199 = v17
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(64)
	return v965
L2:
	;
	v843 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(int32(6)) < base.Ui32(v844) {
		goto L182
	} else {
		goto L183
	}
L3:
	;
	if v199&int32(2) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L4:
	;
	v24 = F_memtoull(m, v16, v14+int32(56))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	if v26 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v29&int32(16) == int32(0) {
		v199 = v29
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-1)))))
	switch v37 & int32(7) {
	case 0:
		goto L12
	case 1:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L9
	case 4:
		goto L8
	default:
		v54 = int32(0)
		goto L7
	}
L7:
	;
	v56 = v14 + int32(48)
	v57 = int32(0)
	if base.Ui32(v54+int32(-21)) < base.Ui32(int32(-20)) {
		v191 = v57
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
	v54 = v53
	goto L7
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
	v54 = v50
	goto L7
L10:
	;
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
	v54 = v47
	goto L7
L11:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
	v54 = v44
	goto L7
L12:
	;
	v54 = int32(base.Ui32(v37) >> (uint(int32(3)) % 32))
	goto L7
L13:
	;
	if v191 != 0 {
		goto L2
	} else {
		goto L40
	}
L14:
	;
	goto L13
L15:
	;
	v69 = int32(1)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v54 != v69 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v191 = int32(1)
	goto L14
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v56))) = v172
	goto L16
L18:
	;
	if v70&int32(255) == int32(45) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v74 = v70 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v74&int32(255)) {
		v191 = v57
		goto L14
	} else {
		goto L20
	}
L20:
	;
	if v56 == int32(0) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v172 = base.I64_extend_i32_u(v74) & int64(255)
	goto L17
L22:
	;
	if base.Ui32(int32(8)) < base.Ui32((v93+int32(-49))&int32(255)) {
		v191 = v57
		goto L14
	} else {
		goto L25
	}
L23:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v92 = int32(2)
	v93 = v90
	v94 = v16 + int32(1)
	goto L22
L24:
	;
	v92 = v69
	v93 = v70
	v94 = v16
	goto L22
L25:
	;
	v105 = base.I64_extend_i32_u(v93+int32(-48)) & int64(255)
	if base.Ui32(v54) <= base.Ui32(v92) {
		v148 = v105
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if v70&int32(255) != int32(45) {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	v111 = v92
	v113 = v105
	v115 = v94
	goto L28
L28:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
	if base.Ui32((v117+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v191 = v57
		goto L14
	} else {
		goto L30
	}
L29:
	;
	v148 = v138
	goto L26
L30:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v113) {
		v191 = v57
		goto L14
	} else {
		goto L31
	}
L31:
	;
	v127 = v113 * int64(10)
	v132 = base.I64_extend_i32_u(v117+int32(-48)) & int64(255)
	if base.Ui64(v132^int64(-1)) < base.Ui64(v127) {
		v191 = v57
		goto L14
	} else {
		goto L32
	}
L32:
	;
	v136 = int32(1)
	v138 = v127 + v132
	v140 = v111 + v136
	if v140 != v54 {
		v111 = v140
		v113 = v138
		v115 = v115 + v136
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	if v148 < int64(0) {
		v191 = v57
		goto L14
	} else {
		goto L38
	}
L35:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v148) {
		v191 = v57
		goto L14
	} else {
		goto L36
	}
L36:
	;
	if v56 == int32(0) {
		goto L16
	} else {
		goto L37
	}
L37:
	;
	v172 = int64(0) - v148
	goto L17
L38:
	;
	if v56 == int32(0) {
		goto L16
	} else {
		goto L39
	}
L39:
	;
	v172 = v148
	goto L17
L40:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v199 = v198
	goto L3
L41:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v420&int32(4) == int32(0) {
		v440 = v420
		goto L92
	} else {
		goto L93
	}
L42:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-1)))))
	v209 = v207 & int32(7)
	switch v209 {
	case 0:
		goto L48
	case 1:
		goto L47
	case 2:
		goto L46
	case 3:
		goto L45
	case 4:
		goto L44
	default:
		goto L41
	}
L43:
	;
	if base.Ui32(v224) < base.Ui32(int32(2)) {
		goto L41
	} else {
		goto L49
	}
L44:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
	v224 = v223
	goto L43
L45:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
	v224 = v220
	goto L43
L46:
	;
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
	v224 = v217
	goto L43
L47:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
	v224 = v214
	goto L43
L48:
	;
	v224 = int32(base.Ui32(v207) >> (uint(int32(3)) % 32))
	goto L43
L49:
	;
	switch v209 {
	default:
		goto L55
	case 1:
		goto L54
	case 2:
		goto L53
	case 3:
		goto L52
	case 4:
		goto L51
	}
L50:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v241+int32(-1)))))
	if v245 != int32(37) {
		goto L41
	} else {
		goto L56
	}
L51:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
	v241 = v240
	goto L50
L52:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
	v241 = v237
	goto L50
L53:
	;
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
	v241 = v234
	goto L50
L54:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
	v241 = v231
	goto L50
L55:
	;
	v241 = int32(base.Ui32(v207) >> (uint(int32(3)) % 32))
	goto L50
L56:
	;
	switch v209 {
	default:
		goto L62
	case 1:
		goto L61
	case 2:
		goto L60
	case 3:
		goto L59
	case 4:
		goto L58
	}
L57:
	;
	v264 = v262 + int32(-1)
	v266 = v14 + int32(48)
	v267 = int32(0)
	if base.Ui32(v262+int32(-22)) < base.Ui32(int32(-20)) {
		v401 = v267
		goto L64
	} else {
		goto L65
	}
L58:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
	v262 = v261
	goto L57
L59:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
	v262 = v258
	goto L57
L60:
	;
	v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
	v262 = v255
	goto L57
L61:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
	v262 = v252
	goto L57
L62:
	;
	v262 = int32(base.Ui32(v207) >> (uint(int32(3)) % 32))
	goto L57
L63:
	;
	if v401 == int32(0) {
		goto L41
	} else {
		goto L90
	}
L64:
	;
	goto L63
L65:
	;
	v279 = int32(1)
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v264 != v279 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v401 = int32(1)
	goto L64
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v266))) = v382
	goto L66
L68:
	;
	if v280&int32(255) == int32(45) {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	v284 = v280 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v284&int32(255)) {
		v401 = v267
		goto L64
	} else {
		goto L70
	}
L70:
	;
	if v266 == int32(0) {
		goto L66
	} else {
		goto L71
	}
L71:
	;
	v382 = base.I64_extend_i32_u(v284) & int64(255)
	goto L67
L72:
	;
	if base.Ui32(int32(8)) < base.Ui32((v303+int32(-49))&int32(255)) {
		v401 = v267
		goto L64
	} else {
		goto L75
	}
L73:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v302 = int32(2)
	v303 = v300
	v304 = v16 + int32(1)
	goto L72
L74:
	;
	v302 = v279
	v303 = v280
	v304 = v16
	goto L72
L75:
	;
	v315 = base.I64_extend_i32_u(v303+int32(-48)) & int64(255)
	if base.Ui32(v264) <= base.Ui32(v302) {
		v358 = v315
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if v280&int32(255) != int32(45) {
		goto L84
	} else {
		goto L85
	}
L77:
	;
	v321 = v302
	v323 = v315
	v325 = v304
	goto L78
L78:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+1)))
	if base.Ui32((v327+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v401 = v267
		goto L64
	} else {
		goto L80
	}
L79:
	;
	v358 = v348
	goto L76
L80:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v323) {
		v401 = v267
		goto L64
	} else {
		goto L81
	}
L81:
	;
	v337 = v323 * int64(10)
	v342 = base.I64_extend_i32_u(v327+int32(-48)) & int64(255)
	if base.Ui64(v342^int64(-1)) < base.Ui64(v337) {
		v401 = v267
		goto L64
	} else {
		goto L82
	}
L82:
	;
	v346 = int32(1)
	v348 = v337 + v342
	v350 = v321 + v346
	if v350 != v264 {
		v321 = v350
		v323 = v348
		v325 = v325 + v346
		goto L78
	} else {
		goto L83
	}
L83:
	;
	goto L79
L84:
	;
	if v358 < int64(0) {
		v401 = v267
		goto L64
	} else {
		goto L88
	}
L85:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v358) {
		v401 = v267
		goto L64
	} else {
		goto L86
	}
L86:
	;
	if v266 == int32(0) {
		goto L66
	} else {
		goto L87
	}
L87:
	;
	v382 = int64(0) - v358
	goto L67
L88:
	;
	if v266 == int32(0) {
		goto L66
	} else {
		goto L89
	}
L89:
	;
	v382 = v358
	goto L67
L90:
	;
	v410 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
	if v410 < int64(0) {
		goto L41
	} else {
		goto L91
	}
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = int64(0) - v410
	goto L2
L92:
	;
	if v440&int32(8) == int32(0) {
		v637 = v440
		goto L100
	} else {
		goto L101
	}
L93:
	;
	v425 = int32(9116376)
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = int32(0)
	v432 = F_strtox_2(m, v16, v14+int32(56), int32(8), int64(-9223372036854775807-1))
	mBase = m.M
	goto L95
L95:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = v432
	v434 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	if v434 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v440 = v439
	goto L92
L97:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435))))
	if v436 == int32(0) {
		goto L2
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v837 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = v837
	goto L2
L100:
	;
	if v637 != 0 {
		v804 = v637
		goto L136
	} else {
		goto L137
	}
L101:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-1)))))
	switch v448 & int32(7) {
	case 0:
		goto L107
	case 1:
		goto L106
	case 2:
		goto L105
	case 3:
		goto L104
	case 4:
		goto L103
	default:
		v465 = int32(0)
		goto L102
	}
L102:
	;
	v467 = v14 + int32(56)
	v475 = m.G0
	v477 = v475 - int32(16)
	m.G0 = v477
	if base.Ui32(v465+int32(-21)) < base.Ui32(int32(-20)) {
		goto L111
	} else {
		goto L112
	}
L103:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
	v465 = v464
	goto L102
L104:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
	v465 = v461
	goto L102
L105:
	;
	v458 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
	v465 = v458
	goto L102
L106:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
	v465 = v455
	goto L102
L107:
	;
	v465 = int32(base.Ui32(v448) >> (uint(int32(3)) % 32))
	goto L102
L108:
	;
	if v624 != 0 {
		goto L99
	} else {
		goto L135
	}
L109:
	;
	m.G0 = v477 + int32(16)
	goto L108
L110:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v467))) = v617
	v624 = int32(1)
	goto L109
L111:
	;
	v588 = int32(0)
	v589 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v589))) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v477)+12)) = v588
	v597 = F_strtoull(m, v16, v477+int32(12), int32(10))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v467))) = v597
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v589)))
	if v599 == int32(28) {
		v624 = v588
		goto L109
	} else {
		goto L132
	}
L112:
	;
	v483 = int32(1)
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v465 != v483 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	if v484&int32(255) != int32(45) {
		v504 = v483
		v505 = v484
		v506 = v16
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v488 = v484 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v488&int32(255)) {
		goto L111
	} else {
		goto L115
	}
L115:
	;
	v617 = base.I64_extend_i32_u(v488) & int64(255)
	goto L110
L116:
	;
	if base.Ui32(int32(8)) < base.Ui32((v505+int32(-49))&int32(255)) {
		goto L111
	} else {
		goto L118
	}
L117:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v504 = int32(2)
	v505 = v502
	v506 = v16 + int32(1)
	goto L116
L118:
	;
	v517 = base.I64_extend_i32_u(v505+int32(-48)) & int64(255)
	if base.Ui32(v465) <= base.Ui32(v504) {
		v560 = v517
		goto L119
	} else {
		goto L120
	}
L119:
	;
	if v484&int32(255) != int32(45) {
		goto L127
	} else {
		goto L128
	}
L120:
	;
	v523 = v504
	v525 = v517
	v527 = v506
	goto L121
L121:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527)+1)))
	if base.Ui32((v529+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L111
	} else {
		goto L123
	}
L122:
	;
	v560 = v550
	goto L119
L123:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v525) {
		goto L111
	} else {
		goto L124
	}
L124:
	;
	v539 = v525 * int64(10)
	v544 = base.I64_extend_i32_u(v529+int32(-48)) & int64(255)
	if base.Ui64(v544^int64(-1)) < base.Ui64(v539) {
		goto L111
	} else {
		goto L125
	}
L125:
	;
	v548 = int32(1)
	v550 = v539 + v544
	v552 = v523 + v548
	if v552 != v465 {
		v523 = v552
		v525 = v550
		v527 = v527 + v548
		goto L121
	} else {
		goto L126
	}
L126:
	;
	goto L122
L127:
	;
	if int64(0) <= v560 {
		v617 = v560
		goto L110
	} else {
		goto L131
	}
L128:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v560) {
		goto L111
	} else {
		goto L129
	}
L129:
	;
	v571 = int64(-1)
	if v571 < v560+v571 {
		v624 = int32(0)
		goto L109
	} else {
		goto L130
	}
L130:
	;
	v617 = int64(0)
	goto L110
L131:
	;
	goto L111
L132:
	;
	if v599 == int32(68) {
		v624 = v588
		goto L109
	} else {
		goto L133
	}
L133:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v604 == int32(0) {
		v624 = v588
		goto L109
	} else {
		goto L134
	}
L134:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v477)+12))
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607))))
	v624 = base.B2i32(v608 == int32(0))
	goto L109
L135:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v637 = v636
	goto L100
L136:
	;
	v806 = int32(3)
	if v804&v806 != v806 {
		goto L172
	} else {
		goto L173
	}
L137:
	;
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-1)))))
	switch v642 & int32(7) {
	case 0:
		goto L143
	case 1:
		goto L142
	case 2:
		goto L141
	case 3:
		goto L140
	case 4:
		goto L139
	default:
		v659 = int32(0)
		goto L138
	}
L138:
	;
	v661 = v14 + int32(48)
	v662 = int32(0)
	if base.Ui32(v659+int32(-21)) < base.Ui32(int32(-20)) {
		v796 = v662
		goto L145
	} else {
		goto L146
	}
L139:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
	v659 = v658
	goto L138
L140:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
	v659 = v655
	goto L138
L141:
	;
	v652 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
	v659 = v652
	goto L138
L142:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
	v659 = v649
	goto L138
L143:
	;
	v659 = int32(base.Ui32(v642) >> (uint(int32(3)) % 32))
	goto L138
L144:
	;
	if v796 != 0 {
		goto L2
	} else {
		goto L171
	}
L145:
	;
	goto L144
L146:
	;
	v674 = int32(1)
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v659 != v674 {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	v796 = int32(1)
	goto L145
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v661))) = v777
	goto L147
L149:
	;
	if v675&int32(255) == int32(45) {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	v679 = v675 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v679&int32(255)) {
		v796 = v662
		goto L145
	} else {
		goto L151
	}
L151:
	;
	if v661 == int32(0) {
		goto L147
	} else {
		goto L152
	}
L152:
	;
	v777 = base.I64_extend_i32_u(v679) & int64(255)
	goto L148
L153:
	;
	if base.Ui32(int32(8)) < base.Ui32((v698+int32(-49))&int32(255)) {
		v796 = v662
		goto L145
	} else {
		goto L156
	}
L154:
	;
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v697 = int32(2)
	v698 = v695
	v699 = v16 + int32(1)
	goto L153
L155:
	;
	v697 = v674
	v698 = v675
	v699 = v16
	goto L153
L156:
	;
	v710 = base.I64_extend_i32_u(v698+int32(-48)) & int64(255)
	if base.Ui32(v659) <= base.Ui32(v697) {
		v753 = v710
		goto L157
	} else {
		goto L158
	}
L157:
	;
	if v675&int32(255) != int32(45) {
		goto L165
	} else {
		goto L166
	}
L158:
	;
	v716 = v697
	v718 = v710
	v720 = v699
	goto L159
L159:
	;
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v720)+1)))
	if base.Ui32((v722+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v796 = v662
		goto L145
	} else {
		goto L161
	}
L160:
	;
	v753 = v743
	goto L157
L161:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v718) {
		v796 = v662
		goto L145
	} else {
		goto L162
	}
L162:
	;
	v732 = v718 * int64(10)
	v737 = base.I64_extend_i32_u(v722+int32(-48)) & int64(255)
	if base.Ui64(v737^int64(-1)) < base.Ui64(v732) {
		v796 = v662
		goto L145
	} else {
		goto L163
	}
L163:
	;
	v741 = int32(1)
	v743 = v732 + v737
	v745 = v716 + v741
	if v745 != v659 {
		v716 = v745
		v718 = v743
		v720 = v720 + v741
		goto L159
	} else {
		goto L164
	}
L164:
	;
	goto L160
L165:
	;
	if v753 < int64(0) {
		v796 = v662
		goto L145
	} else {
		goto L169
	}
L166:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v753) {
		v796 = v662
		goto L145
	} else {
		goto L167
	}
L167:
	;
	if v661 == int32(0) {
		goto L147
	} else {
		goto L168
	}
L168:
	;
	v777 = int64(0) - v753
	goto L148
L169:
	;
	if v661 == int32(0) {
		goto L147
	} else {
		goto L170
	}
L170:
	;
	v777 = v753
	goto L148
L171:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v804 = v803
	goto L136
L172:
	;
	if v804&int32(1) == int32(0) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a440)
	v965 = int32(0)
	goto L1
L174:
	;
	if v804&int32(4) == int32(0) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a441)
	v965 = int32(0)
	goto L1
L176:
	;
	if v804&int32(8) == int32(0) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a442)
	v965 = int32(0)
	goto L1
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a443)
	v965 = int32(0)
	goto L1
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a444)
	v965 = int32(0)
	goto L1
L180:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v911 == int32(0) {
		v918 = v844
		goto L202
	} else {
		goto L203
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a439)
	v965 = int32(0)
	goto L1
L182:
	;
	if int64(-1) < v843 {
		goto L193
	} else {
		goto L194
	}
L183:
	;
	if int32(1)<<(uint(v844)%32)&int32(106) == int32(0) {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v853 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	v854 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui64(v854) < base.Ui64(v843) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v854
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v853
	if v857&int32(4) != 0 {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	if base.Ui64(v853) <= base.Ui64(v843) {
		goto L180
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	v866 = int32(_a445)
	goto L190
L189:
	;
	v866 = int32(_a446)
	goto L190
L190:
	;
	v869 = F_snprintf(m, int32(_a439), int32(256), v866, v14+int32(32))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	return int32(0)
L192:
	;
	goto L181
L193:
	;
	v890 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	v891 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if v891 < v843 {
		goto L198
	} else {
		goto L199
	}
L194:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v875&int32(2) == int32(0) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v880 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if v880 <= v843 {
		goto L180
	} else {
		goto L196
	}
L196:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = int64(0) - v880
	v888 = F_snprintf(m, int32(_a439), int32(256), int32(_a447), v14)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L191
	} else {
		goto L197
	}
L197:
	;
	goto L181
L198:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v891
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v890
	v901 = F_snprintf(m, int32(_a439), int32(256), int32(_a448), v14+int32(16))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L191
	} else {
		goto L201
	}
L199:
	;
	if v890 <= v843 {
		goto L180
	} else {
		goto L200
	}
L200:
	;
	goto L198
L201:
	;
	goto L181
L202:
	;
	switch v918 {
	case 0:
		goto L217
	case 1:
		goto L216
	case 2:
		goto L215
	case 3:
		goto L214
	case 4:
		goto L213
	case 5:
		goto L212
	case 6:
		goto L211
	case 7:
		goto L210
	case 8:
		goto L209
	case 9:
		goto L208
	default:
		v954 = int64(0)
		goto L207
	}
L203:
	;
	v914 = m.T0[v911].(func(*base.Module, int64, int32) int32)(m, v843, l3)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L191
	} else {
		goto L205
	}
L204:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v918 = v917
	goto L202
L205:
	;
	if v914 != 0 {
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v965 = int32(0)
	goto L1
L207:
	;
	v955 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
	if v954 == v955 {
		goto L222
	} else {
		goto L223
	}
L208:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v953 = *(*int64)(unsafe.Add(mBase, uint32(v952)))
	v954 = v953
	goto L207
L209:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v951 = *(*int64)(unsafe.Add(mBase, uint32(v950)))
	v954 = v951
	goto L207
L210:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v949 = int64(*(*int32)(unsafe.Add(mBase, uint32(v948))))
	v954 = v949
	goto L207
L211:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v947 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v946))))
	v954 = v947
	goto L207
L212:
	;
	v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v937&int32(1) == int32(0) {
		goto L220
	} else {
		goto L221
	}
L213:
	;
	v928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v928&int32(1) == int32(0) {
		goto L218
	} else {
		goto L219
	}
L214:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v927 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v926))))
	v954 = v927
	goto L207
L215:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v925 = int64(*(*int32)(unsafe.Add(mBase, uint32(v924))))
	v954 = v925
	goto L207
L216:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v923 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v922))))
	v954 = v923
	goto L207
L217:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v921 = int64(*(*int32)(unsafe.Add(mBase, uint32(v920))))
	v954 = v921
	goto L207
L218:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v936 = *(*int64)(unsafe.Add(mBase, uint32(v935)))
	v954 = v936
	goto L207
L219:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v934 = F_getModuleNumericConfig(m, v933)
	mBase = m.M
	v954 = v934
	goto L207
L220:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v945 = *(*int64)(unsafe.Add(mBase, uint32(v944)))
	v954 = v945
	goto L207
L221:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v943 = F_getModuleUnsignedNumericConfig(m, v942)
	mBase = m.M
	v954 = v943
	goto L207
L222:
	;
	v960 = int32(2)
	v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v961&v960 != 0 {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	v957 = F_setNumericType(m, l0, v955, l3)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L191
	} else {
		goto L224
	}
L224:
	;
	v965 = v957
	goto L1
L225:
	;
	v964 = int32(1)
	goto L227
L226:
	;
	v964 = v960
	goto L227
L227:
	;
	v965 = v964
	goto L1
}
