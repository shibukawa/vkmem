package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___nl_langinfo_l(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	if l0 != int32(14) {
		v14 = l0 >> (uint(int32(16)) % 32)
		v15 = int32(65535)
		v16 = l0 & v15
		if v16 != v15 {
			v30 = int32(_a_F___nl_langinfo_l_0)
			switch v14 + int32(-1) {
			case 0:
				if base.Ui32(int32(1)) < base.Ui32(v16) {
					v58 = v30
					return v58
				} else {
					v42 = int32(_a_F___nl_langinfo_l_1)
					if v16 != 0 {
						v44 = v42
						v47 = v16
						for {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
							v51 = v44 + int32(1)
							if v49 != 0 {
								v44 = v51
								continue
							} else {
							}
							v53 = v47 + int32(-1)
							if v53 != 0 {
								v44 = v51
								v47 = v53
								continue
							} else {
								break
							}
							break
						}
						v58 = v51
						return v58
					} else {
						return v42
					}
				}
			case 1:
				if base.Ui32(int32(49)) < base.Ui32(v16) {
					v58 = v30
					return v58
				} else {
					v42 = int32(_a_F___nl_langinfo_l_2)
					if v16 != 0 {
						v44 = v42
						v47 = v16
						for {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
							v51 = v44 + int32(1)
							if v49 != 0 {
								v44 = v51
								continue
							} else {
							}
							v53 = v47 + int32(-1)
							if v53 != 0 {
								v44 = v51
								v47 = v53
								continue
							} else {
								break
							}
							break
						}
						v58 = v51
						return v58
					} else {
						return v42
					}
				}
			default:
				v58 = v30
				return v58
			case 4:
				if base.Ui32(int32(3)) < base.Ui32(v16) {
					v58 = v30
					return v58
				} else {
					v42 = int32(_a_F___nl_langinfo_l_3)
					if v16 != 0 {
						v44 = v42
						v47 = v16
						for {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
							v51 = v44 + int32(1)
							if v49 != 0 {
								v44 = v51
								continue
							} else {
							}
							v53 = v47 + int32(-1)
							if v53 != 0 {
								v44 = v51
								v47 = v53
								continue
							} else {
								break
							}
							break
						}
						v58 = v51
						return v58
					} else {
						return v42
					}
				}
			}
		} else {
			if int32(5) < v14 {
				v30 = int32(_a_F___nl_langinfo_l_0)
				switch v14 + int32(-1) {
				case 0:
					if base.Ui32(int32(1)) < base.Ui32(v16) {
						v58 = v30
						return v58
					} else {
						v42 = int32(_a_F___nl_langinfo_l_1)
						if v16 != 0 {
							v44 = v42
							v47 = v16
							for {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
								v51 = v44 + int32(1)
								if v49 != 0 {
									v44 = v51
									continue
								} else {
								}
								v53 = v47 + int32(-1)
								if v53 != 0 {
									v44 = v51
									v47 = v53
									continue
								} else {
									break
								}
								break
							}
							v58 = v51
							return v58
						} else {
							return v42
						}
					}
				case 1:
					if base.Ui32(int32(49)) < base.Ui32(v16) {
						v58 = v30
						return v58
					} else {
						v42 = int32(_a_F___nl_langinfo_l_2)
						if v16 != 0 {
							v44 = v42
							v47 = v16
							for {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
								v51 = v44 + int32(1)
								if v49 != 0 {
									v44 = v51
									continue
								} else {
								}
								v53 = v47 + int32(-1)
								if v53 != 0 {
									v44 = v51
									v47 = v53
									continue
								} else {
									break
								}
								break
							}
							v58 = v51
							return v58
						} else {
							return v42
						}
					}
				default:
					v58 = v30
					return v58
				case 4:
					if base.Ui32(int32(3)) < base.Ui32(v16) {
						v58 = v30
						return v58
					} else {
						v42 = int32(_a_F___nl_langinfo_l_3)
						if v16 != 0 {
							v44 = v42
							v47 = v16
							for {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
								v51 = v44 + int32(1)
								if v49 != 0 {
									v44 = v51
									continue
								} else {
								}
								v53 = v47 + int32(-1)
								if v53 != 0 {
									v44 = v51
									v47 = v53
									continue
								} else {
									break
								}
								break
							}
							v58 = v51
							return v58
						} else {
							return v42
						}
					}
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1+v14<<(uint(int32(2))%32))))
				if v24 != 0 {
					v28 = v24 + int32(8)
				} else {
					v28 = int32(_a_F___nl_langinfo_l_4)
				}
				return v28
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v10 != 0 {
			v11 = int32(_a_F___nl_langinfo_l_5)
		} else {
			v11 = int32(_a_F___nl_langinfo_l_6)
		}
		return v11
	}
}
func F_nextArgument(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v143 int64
	_ = v143
	var v147 int32
	_ = v147
	var v161 int64
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v206 int32
	_ = v206
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v206 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v206
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v206
	return v206
L2:
	;
	if l1 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v14 == int32(36) {
		v123 = l0
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v124 = l0 + l1
	v126 = v123 + int32(1)
	if v124-v126 <= int32(-1) {
		goto L32
	} else {
		goto L33
	}
L5:
	;
	v18 = int32(0)
	v21 = base.B2i32(l1 != v18)
	if l0&int32(3) == v18 {
		v47 = l0
		v49 = l1
		v50 = v21
		goto L9
	} else {
		goto L10
	}
L6:
	;
	if v120 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L7:
	;
	v120 = int32(0)
	goto L6
L8:
	;
	v98 = v91
	v100 = v93
	goto L26
L9:
	;
	if v50 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L10:
	;
	if l1 == int32(0) {
		v47 = l0
		v49 = l1
		v50 = v21
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v30 = l0
	v32 = l1
	goto L12
L12:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v35 == int32(36) {
		v91 = v30
		v93 = v32
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v47 = v42
	v49 = v38
	v50 = v40
	goto L9
L14:
	;
	v38 = v32 + int32(-1)
	v39 = int32(0)
	v40 = base.B2i32(v38 != v39)
	v42 = v30 + int32(1)
	if v42&int32(3) == v39 {
		v47 = v42
		v49 = v38
		v50 = v40
		goto L9
	} else {
		goto L15
	}
L15:
	;
	if v38 != 0 {
		v30 = v42
		v32 = v38
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v54 == int32(36) {
		v84 = v47
		v86 = v49
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v86 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L19:
	;
	if base.Ui32(v49) < base.Ui32(int32(4)) {
		v84 = v47
		v86 = v49
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v64 = v47
	v66 = v49
	goto L21
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v71 = v70 ^ int32(606348324)
	v74 = int32(-2139062144)
	if (int32(16843008)-v71|v71)&v74 != v74 {
		v91 = v64
		v93 = v66
		goto L8
	} else {
		goto L23
	}
L22:
	;
	v84 = v79
	v86 = v81
	goto L18
L23:
	;
	v79 = v64 + int32(4)
	v81 = v66 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v81) {
		v64 = v79
		v66 = v81
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v91 = v84
	v93 = v86
	goto L8
L26:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v103 != int32(36) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L7
L28:
	;
	v108 = v100 + int32(-1)
	if v108 != 0 {
		v98 = v98 + int32(1)
		v100 = v108
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v120 = v98
	goto L6
L30:
	;
	goto L27
L31:
	;
	v123 = v120
	goto L4
L32:
	;
	v188 = m.G3
	m.Env.X__assert_fail(m, v188+int32(_a_F_nextArgument_0), v188+int32(_a_F_nextArgument_1), int32(822), v188+int32(_a_F_nextArgument_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	if v124 == v126 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v131 = int32(*(*int8)(unsafe.Add(mBase, uint32(v126))))
	if base.Ui32(v131+int32(-58)) < base.Ui32(int32(-10)) {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v141 = v126
	v143 = int64(0)
	goto L37
L36:
	;
	v165 = int32(2)
	v166 = v141 + v165
	v167 = base.I32_wrap_i64(v143)
	v170 = v166 + v167 + v165
	v171 = v170 - l0
	if base.Ui32(l1) < base.Ui32(v171) {
		goto L1
	} else {
		goto L42
	}
L37:
	;
	if base.Ui32(v124) <= base.Ui32(v141) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v141))))
	if base.Ui32(v147+int32(-58)) < base.Ui32(int32(-10)) {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v161 = (base.I64_extend_i32_s(v147)+int64(4294967248))&int64(4294967295) + v143*int64(10)
	if base.Ui64(v161) < base.Ui64(int64(536870913)) {
		v141 = v141 + int32(1)
		v143 = v161
		goto L37
	} else {
		goto L41
	}
L41:
	;
	goto L1
L42:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v173 != int32(13) {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+v167+int32(2)))))
	if v179 != int32(13) {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v166
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v167
	if v171 == l1 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v186 = int32(0)
	goto L47
L46:
	;
	v186 = v170
	goto L47
L47:
	;
	return v186
}
func F_ntohl(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = int32(24)
	v4 = int32(65280)
	v6 = int32(8)
	return l0<<(uint(v2)%32) | l0&v4<<(uint(v6)%32) | (int32(base.Ui32(l0)>>(uint(v6)%32))&v4 | int32(base.Ui32(l0)>>(uint(v2)%32)))
}
