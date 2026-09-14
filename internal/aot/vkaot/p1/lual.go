package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaL_checkany(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v52 = l0 + int32(72)
				v53 = m.G398
				if v52 != v53 {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
					v59 = v56
				} else {
					v59 = int32(-1)
				}
			case 1:
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v25
				v52 = l0 + int32(88)
				v53 = m.G398
				if v52 != v53 {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
					v59 = v56
				} else {
					v59 = int32(-1)
				}
			case 2:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v52 = v48 + int32(96)
				v53 = m.G398
				if v52 != v53 {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
					v59 = v56
				} else {
					v59 = int32(-1)
				}
			default:
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+7)))
				if base.Ui32(v38) < base.Ui32(int32(-10002)-l1) {
					v59 = int32(-1)
				} else {
					v52 = v37 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
					v53 = m.G398
					if v52 != v53 {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
						v59 = v56
					} else {
						v59 = int32(-1)
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v52 = v16 + l1<<(uint(int32(4))%32)
			v53 = m.G398
			if v52 != v53 {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
				v59 = v56
			} else {
				v59 = int32(-1)
			}
		}
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v10 = v5 + l1<<(uint(int32(4))%32) + int32(-16)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v10) < base.Ui32(v11) {
			v52 = v10
			v53 = m.G398
			if v52 != v53 {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
				v59 = v56
			} else {
				v59 = int32(-1)
			}
		} else {
			v59 = int32(-1)
		}
	}
	if v59 != int32(-1) {
		return
	} else {
		v62 = m.G3
		v65 = F_luaL_argerror(m, l0, l1, v62+int32(_a2693))
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return
		} else {
			return
		}
	}
}
func F_luaL_checknumber(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_lua_tonumber(m, l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return float64(0)
	} else {
		if base.F64_ne(v10, float64(0)) != 0 {
			m.G0 = v8 + int32(16)
			return v10
		} else {
			v16 = F_lua_isnumber(m, l0, l1)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return float64(0)
			} else {
				if v16 != 0 {
					m.G0 = v8 + int32(16)
					return v10
				} else {
					v25 = m.G399
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(12))))
					if l1 < int32(1) {
						if l1 < int32(-9999) {
							switch l1 + int32(10002) {
							case 0:
								v80 = l0 + int32(72)
								v81 = m.G398
								if v80 != v81 {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
									v87 = v84
								} else {
									v87 = int32(-1)
								}
							case 1:
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v53
								v80 = l0 + int32(88)
								v81 = m.G398
								if v80 != v81 {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
									v87 = v84
								} else {
									v87 = int32(-1)
								}
							case 2:
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v80 = v76 + int32(96)
								v81 = m.G398
								if v80 != v81 {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
									v87 = v84
								} else {
									v87 = int32(-1)
								}
							default:
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
								v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+7)))
								if base.Ui32(v66) < base.Ui32(int32(-10002)-l1) {
									v87 = int32(-1)
								} else {
									v80 = v65 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
									v81 = m.G398
									if v80 != v81 {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
										v87 = v84
									} else {
										v87 = int32(-1)
									}
								}
							}
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v80 = v44 + l1<<(uint(int32(4))%32)
							v81 = m.G398
							if v80 != v81 {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
								v87 = v84
							} else {
								v87 = int32(-1)
							}
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v38 = v33 + l1<<(uint(int32(4))%32) + int32(-16)
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if base.Ui32(v38) < base.Ui32(v39) {
							v80 = v38
							v81 = m.G398
							if v80 != v81 {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
								v87 = v84
							} else {
								v87 = int32(-1)
							}
						} else {
							v87 = int32(-1)
						}
					}
					v89 = m.G3
					if v87 != int32(-1) {
						v94 = m.G399
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v87<<(uint(int32(2))%32))))
						v99 = v98
					} else {
						v99 = v89 + int32(_a2694)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v99
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v29
					v102 = m.G3
					v105 = F_lua_pushfstring(m, l0, v102+int32(_a2695), v8)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return float64(0)
					} else {
						v107 = F_luaL_argerror(m, l0, l1, v105)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return float64(0)
						} else {
							m.G0 = v8 + int32(16)
							return v10
						}
					}
				}
			}
		}
	}
}
func F_luaL_checkoption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v79 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v74 = F_luaL_checklstring(m, l0, l1, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L20
	} else {
		goto L21
	}
L3:
	;
	if l1 < int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v70 < int32(1) {
		v78 = l2
		goto L1
	} else {
		goto L19
	}
L5:
	;
	v64 = m.G398
	if v63 != v64 {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	if l1 < int32(-9999) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v21 = v16 + l1<<(uint(int32(4))%32) + int32(-16)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v21) < base.Ui32(v22) {
		v63 = v21
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v70 = int32(-1)
	goto L4
L9:
	;
	switch l1 + int32(10002) {
	case 0:
		goto L13
	case 1:
		goto L14
	case 2:
		goto L11
	default:
		goto L12
	}
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v63 = v27 + l1<<(uint(int32(4))%32)
	goto L5
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v63 = v59 + int32(96)
	goto L5
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+7)))
	if base.Ui32(v49) < base.Ui32(int32(-10002)-l1) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v63 = l0 + int32(72)
	goto L5
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v36
	v63 = l0 + int32(88)
	goto L5
L15:
	;
	v70 = int32(-1)
	goto L4
L16:
	;
	v63 = v48 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
	goto L5
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	v70 = v67
	goto L4
L18:
	;
	v70 = int32(-1)
	goto L4
L19:
	;
	goto L2
L20:
	;
	return int32(0)
L21:
	;
	v78 = v74
	goto L1
L22:
	;
	m.G0 = v10 + int32(16)
	return v150
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v78
	v137 = m.G3
	v140 = F_lua_pushfstring(m, l0, v137+int32(_a2692), v10)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L20
	} else {
		goto L37
	}
L24:
	;
	v88 = v79
	v89 = int32(0)
	goto L25
L25:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v93 == int32(0) {
		v116 = v92
		v117 = v93
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L23
L27:
	;
	if v117-v116&int32(255) == int32(0) {
		v150 = v89
		goto L22
	} else {
		goto L35
	}
L28:
	;
	goto L27
L29:
	;
	if v93 != v92&int32(255) {
		v116 = v92
		v117 = v93
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v99 = v88
	v100 = v78
	goto L31
L31:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if v104 == int32(0) {
		v116 = v103
		v117 = v104
		goto L28
	} else {
		goto L33
	}
L32:
	;
	v116 = v103
	v117 = v104
	goto L28
L33:
	;
	v107 = int32(1)
	if v104 == v103&int32(255) {
		v99 = v99 + v107
		v100 = v100 + v107
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v124 = v89 + int32(1)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l3+v124<<(uint(int32(2))%32))))
	if v128 != 0 {
		v88 = v128
		v89 = v124
		goto L25
	} else {
		goto L36
	}
L36:
	;
	goto L26
L37:
	;
	v142 = F_luaL_argerror(m, l0, l1, v140)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	v150 = v142
	goto L22
}
func F_luaL_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l2
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	goto L6
L1:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v96 = F_lua_pushvfstring(m, l0, l1, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L17
	} else {
		goto L22
	}
L2:
	;
	v88 = m.G3
	F_lua_pushlstring(m, l0, v88+int32(_a320), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L17
	} else {
		goto L21
	}
L3:
	;
	if v63 == int32(0) {
		goto L2
	} else {
		goto L16
	}
L4:
	;
	goto L3
L5:
	;
	if v35 != 0 {
		v54 = int32(0)
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v22 = int32(1)
	v24 = v16
	goto L7
L7:
	;
	if base.Ui32(v24) <= base.Ui32(v19) {
		v63 = int32(0)
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L5
L9:
	;
	v29 = v22 + int32(-1)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+6)))
	if v32 != 0 {
		v35 = v29
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = v24 + int32(-24)
	if int32(0) < v35 {
		v22 = v35
		v24 = v37
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v35 = v29 - v33
	goto L10
L12:
	;
	goto L8
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7+int32(12))+96)) = v54
	v63 = int32(1)
	goto L4
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v37) <= base.Ui32(v48) {
		v63 = int32(0)
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v52 = base.I32_div_s(v37-v48, int32(24))
	v54 = v52
	goto L13
L16:
	;
	v66 = m.G3
	v71 = F_lua_getinfo(m, l0, v66+int32(_a2690), v7+int32(12))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	if v75 < int32(1) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(48)
	v82 = m.G3
	v85 = F_lua_pushfstring(m, l0, v82+int32(_a2691), v7)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L1
L21:
	;
	goto L1
L22:
	;
	F_lua_concat(m, l0, int32(2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	v101 = F_lua_error(m, l0)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	m.G0 = v7 + int32(112)
	return v101
}
func F_luaL_findtable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v162 int32
	_ = v162
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v213 int32
	_ = v213
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v291 int32
	_ = v291
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v457 int32
	_ = v457
	var v485 int32
	_ = v485
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var __phi529 int32
	_ = __phi529
	var v530 int32
	_ = v530
	var __phi530 int32
	_ = __phi530
	var v532 int64
	_ = v532
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	if l1 < int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v72 = l2
	goto L17
L2:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
	*(*int64)(unsafe.Add(mBase, uint32(v61))) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v66 + int32(16)
	goto L1
L3:
	;
	if l1 < int32(-9999) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = v11 + l1<<(uint(int32(4))%32) + int32(-16)
	v17 = m.G398
	if base.Ui32(v16) < base.Ui32(v10) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v19 = v16
	goto L7
L6:
	;
	v19 = v17
	goto L7
L7:
	;
	v58 = v19
	goto L2
L8:
	;
	switch l1 + int32(10002) {
	case 0:
		goto L11
	case 1:
		goto L12
	case 2:
		goto L13
	default:
		goto L10
	}
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v58 = v22 + l1<<(uint(int32(4))%32)
	goto L2
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+7)))
	v46 = m.G398
	if base.Ui32(v45) < base.Ui32(int32(-10002)-l1) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v58 = l0 + int32(72)
	goto L2
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v34
	v58 = l0 + int32(88)
	goto L2
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v58 = v28 + int32(96)
	goto L2
L14:
	;
	v57 = v46
	goto L16
L15:
	;
	v57 = v44 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
	goto L16
L16:
	;
	v58 = v57
	goto L2
L17:
	;
	v75 = int32(46)
	v76 = F___strchrnul(m, v72, v75)
	mBase = m.M
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v78 == v75 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	return int32(0)
L19:
	;
	v140 = v139 - v72
	F_lua_pushlstring(m, l0, v72, v140)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L41
	} else {
		goto L42
	}
L20:
	;
	if v82 != 0 {
		v139 = v82
		goto L19
	} else {
		goto L24
	}
L21:
	;
	v82 = v76
	goto L23
L22:
	;
	v82 = int32(0)
	goto L23
L23:
	;
	goto L20
L24:
	;
	if v72&int32(3) == int32(0) {
		v104 = v72
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v139 = v72 + v137
	goto L19
L26:
	;
	v137 = v129 - v72
	goto L25
L27:
	;
	v108 = v104
	goto L35
L28:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v90 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v93 = v72
	goto L31
L30:
	;
	v137 = v72 - v72
	goto L25
L31:
	;
	v97 = v93 + int32(1)
	if v97&int32(3) == int32(0) {
		v104 = v97
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v102 != 0 {
		v93 = v97
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v129 = v97
	goto L26
L35:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v117 = int32(-2139062144)
	if (int32(16843008)-v114|v114)&v117 == v117 {
		v108 = v108 + int32(4)
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v123 = v108
	goto L38
L37:
	;
	goto L36
L38:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v127 != 0 {
		v123 = v123 + int32(1)
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v129 = v123
	goto L26
L40:
	;
	goto L39
L41:
	;
	return int32(0)
L42:
	;
	goto L45
L43:
	;
	goto L63
L44:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v162+int32(-32))))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v203 = int32(-16)
	v205 = F_luaH_get(m, v201, v202+v203)
	mBase = m.M
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v205)))
	*(*int64)(unsafe.Add(mBase, uint32(v206+v203))) = v209
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v205)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v206+int32(-8)))) = v213
	goto L43
L45:
	;
	goto L51
L51:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L44
L59:
	;
	goto L133
L60:
	;
	goto L109
L61:
	;
	if v272 != 0 {
		goto L60
	} else {
		goto L76
	}
L62:
	;
	v266 = m.G398
	if v232 != v266 {
		goto L74
	} else {
		goto L75
	}
L63:
	;
	goto L67
L67:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v232 = v229 + int32(-16)
	goto L62
L74:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v232)+8))
	v272 = v269
	goto L61
L75:
	;
	v272 = int32(-1)
	goto L61
L76:
	;
	goto L79
L77:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v303 == int32(46) {
		goto L85
	} else {
		goto L86
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v291 + int32(-16)
	goto L77
L79:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L78
L85:
	;
	v306 = int32(1)
	goto L87
L86:
	;
	v306 = l3
	goto L87
L87:
	;
	F_lua_createtable(m, l0, int32(0), v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L41
	} else {
		goto L88
	}
L88:
	;
	F_lua_pushlstring(m, l0, v72, v140)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L41
	} else {
		goto L89
	}
L89:
	;
	goto L92
L90:
	;
	F_lua_settable(m, l0, int32(-4))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L41
	} else {
		goto L106
	}
L91:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v331)))
	*(*int64)(unsafe.Add(mBase, uint32(v367))) = v368
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v331)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v367)+8)) = v370
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v372 + int32(16)
	goto L90
L92:
	;
	goto L98
L98:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v331 = v328 + int32(-32)
	goto L91
L106:
	;
	goto L59
L107:
	;
	if v436 == int32(5) {
		goto L59
	} else {
		goto L122
	}
L108:
	;
	v430 = m.G398
	if v396 != v430 {
		goto L120
	} else {
		goto L121
	}
L109:
	;
	goto L113
L113:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v396 = v393 + int32(-16)
	goto L108
L120:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v396)+8))
	v436 = v433
	goto L107
L121:
	;
	v436 = int32(-1)
	goto L107
L122:
	;
	goto L125
L123:
	;
	return v72
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v457 + int32(-32)
	goto L123
L125:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L124
L131:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v549 == int32(46) {
		v72 = v139 + int32(1)
		goto L17
	} else {
		goto L152
	}
L132:
	;
	v525 = v485 + int32(-16)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v526) <= base.Ui32(v525) {
		v543 = v526
		goto L147
	} else {
		goto L148
	}
L133:
	;
	goto L139
L139:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L132
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v543 + int32(-16)
	goto L131
L148:
	;
	__phi529 = v485 + int32(-32)
	__phi530 = v525
	v529 = __phi529
	v530 = __phi530
	goto L149
L149:
	;
	v532 = *(*int64)(unsafe.Add(mBase, uint32(v529)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v529))) = v532
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v529)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v529)+8)) = v534
	v537 = v530 + int32(16)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v537) < base.Ui32(v538) {
		__phi529 = v530
		__phi530 = v537
		v529 = __phi529
		v530 = __phi530
		goto L149
	} else {
		goto L151
	}
L150:
	;
	v543 = v538
	goto L147
L151:
	;
	goto L150
L152:
	;
	goto L18
}
func F_luaL_getmetafield(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v113 int32
	_ = v113
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int64
	_ = v160
	var v164 int32
	_ = v164
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v242 int32
	_ = v242
	var v271 int32
	_ = v271
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var __phi315 int32
	_ = __phi315
	var v316 int32
	_ = v316
	var __phi316 int32
	_ = __phi316
	var v318 int64
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	if l1 < int32(1) {
		if l1 < int32(-9999) {
			switch l1 + int32(10002) {
			case 0:
				v58 = l0 + int32(72)
			case 1:
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v34
				v58 = l0 + int32(88)
			case 2:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v58 = v28 + int32(96)
			default:
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+7)))
				v46 = m.G398
				if base.Ui32(v45) < base.Ui32(int32(-10002)-l1) {
					v57 = v46
				} else {
					v57 = v44 + (int32(-10003)-l1)<<(uint(int32(4))%32) + int32(24)
				}
				v58 = v57
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v58 = v22 + l1<<(uint(int32(4))%32)
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v16 = v11 + l1<<(uint(int32(4))%32) + int32(-16)
		v17 = m.G398
		if base.Ui32(v16) < base.Ui32(v10) {
			v19 = v16
		} else {
			v19 = v17
		}
		v58 = v19
	}
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	switch v61 + int32(-5) {
	case 0:
		v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
		v76 = v64 + int32(16)
	default:
		v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v76 = v70 + v61<<(uint(int32(2))%32) + int32(152)
	case 2:
		v67 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
		v76 = v67 + int32(8)
	}
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v77 != 0 {
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = int32(5)
		*(*int32)(unsafe.Add(mBase, uint32(v79))) = v77
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v83 + int32(16)
		v89 = int32(1)
	} else {
		v89 = int32(0)
	}
	if v89 == int32(0) {
		v334 = int32(0)
		return v334
	} else {
		F_lua_pushstring(m, l0, l2)
		mBase = m.M
		v95 = m.ExcPending
		if v95 != 0 {
			return int32(0)
		} else {
			v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v152 = *(*int32)(unsafe.Add(mBase, uint32(v113+int32(-32))))
			v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v154 = int32(-16)
			v156 = F_luaH_get(m, v152, v153+v154)
			mBase = m.M
			v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v160 = *(*int64)(unsafe.Add(mBase, uint32(v156)))
			*(*int64)(unsafe.Add(mBase, uint32(v157+v154))) = v160
			v164 = *(*int32)(unsafe.Add(mBase, uint32(v156)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v157+int32(-8)))) = v164
			v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v183 = v180 + int32(-16)
			v217 = m.G398
			if v183 != v217 {
				v220 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
				v223 = v220
			} else {
				v223 = int32(-1)
			}
			if v223 != 0 {
				v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v311 = v271 + int32(-16)
				v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v312) <= base.Ui32(v311) {
					v329 = v312
				} else {
					__phi315 = v271 + int32(-32)
					__phi316 = v311
					v315 = __phi315
					v316 = __phi316
					for {
						v318 = *(*int64)(unsafe.Add(mBase, uint32(v315)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v315))) = v318
						v320 = *(*int32)(unsafe.Add(mBase, uint32(v315)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v315)+8)) = v320
						v323 = v316 + int32(16)
						v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if base.Ui32(v323) < base.Ui32(v324) {
							__phi315 = v316
							__phi316 = v323
							v315 = __phi315
							v316 = __phi316
							continue
						} else {
							break
						}
						break
					}
					v329 = v324
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v329 + int32(-16)
				v334 = int32(1)
				return v334
			} else {
				v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v242 + int32(-32)
				return int32(0)
			}
		}
	}
}
func F_luaL_prepbuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = l0 + int32(12)
	if v10 == v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v12
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_lua_pushlstring(m, v14, v12, v10-v12)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v12
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21 + v22
	if v21 < v22 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = F_lua_objlen(m, v27, int32(-1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v33 = int32(1)
	v36 = v29
	goto L8
L7:
	;
	if v52 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v43 = F_lua_objlen(m, v27, v33^int32(-1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v45 = int32(1)
	v46 = v33 + v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v52 = base.B2i32(int32(8) < v47-v33) | base.B2i32(base.Ui32(v43) < base.Ui32(v36))
	if v52 != v45 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	if v46 < v47 {
		v33 = v46
		v36 = v43 + v36
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v58 = v46
	goto L15
L14:
	;
	v58 = v33
	goto L15
L15:
	;
	F_lua_concat(m, v27, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61 - v58 + int32(1)
	goto L1
}
func F_luaL_pushresult(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = l0 + int32(12)
	if v4 != v6 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		F_lua_pushlstring(m, v9, v6, v4-v6)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v6
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v16 = v14 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v16
			v18 = v16
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_lua_concat(m, v19, v18)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
				return
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v18 = v8
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		F_lua_concat(m, v19, v18)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
			return
		}
	}
}
func F_luaL_register(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v6 int32
	_ = v6
	F_luaL_openlib(m, l0, l1, l2, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
