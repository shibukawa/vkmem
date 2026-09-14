package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___isspace_2(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0 == int32(32)) | base.B2i32(base.Ui32(l0+int32(-9)) < base.Ui32(int32(5)))
}
func F___isspace_4(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0 == int32(32)) | base.B2i32(base.Ui32(l0+int32(-9)) < base.Ui32(int32(5)))
}
func F_iAmPrimary(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_iAmPrimary[0]))
	if v2 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_iAmPrimary[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
		return base.B2i32(v11&int32(1) != int32(0))
	} else {
		v3 = int32(0)
		v4 = *(*int32)(unsafe.Add(mBase, _c_F_iAmPrimary[2]))
		return base.B2i32(v4 == v3)
	}
}
func F_inBioThread(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	v1 = int32(0)
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_inBioThread[0]))
	return base.B2i32(v2 != v1)
}
func F_inMainThread(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	v1 = int32(0)
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_inMainThread[0]))
	return base.B2i32(v2 == v1)
}
func F_inclinenumber(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v7 + int32(-1)
	if v7 == int32(0) {
		v18 = F_luaZ_fill(m, v6)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = v18
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v20
			switch v20 + int32(-10) {
			case 0, 3:
				if v20 == v5 {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45 + int32(1)
					if v45 < int32(2147483644) {
						return
					} else {
						v51 = m.G3
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						F_luaX_lexerror(m, l0, v51+int32(_a_F_inclinenumber_0), v54)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v27 + int32(-1)
					if v27 == int32(0) {
						v38 = F_luaZ_fill(m, v26)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v40 = v38
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v40
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45 + int32(1)
							if v45 < int32(2147483644) {
								return
							} else {
								v51 = m.G3
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								F_luaX_lexerror(m, l0, v51+int32(_a_F_inclinenumber_0), v54)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v33 + int32(1)
						v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
						v40 = v37
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v40
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45 + int32(1)
						if v45 < int32(2147483644) {
							return
						} else {
							v51 = m.G3
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							F_luaX_lexerror(m, l0, v51+int32(_a_F_inclinenumber_0), v54)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			default:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45 + int32(1)
				if v45 < int32(2147483644) {
					return
				} else {
					v51 = m.G3
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					F_luaX_lexerror(m, l0, v51+int32(_a_F_inclinenumber_0), v54)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v13 + int32(1)
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		v20 = v17
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v20
		switch v20 + int32(-10) {
		case 0, 3:
			if v20 == v5 {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45 + int32(1)
				if v45 < int32(2147483644) {
					return
				} else {
					v51 = m.G3
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					F_luaX_lexerror(m, l0, v51+int32(_a_F_inclinenumber_0), v54)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				*(*int32)(unsafe.Add(mBase, uint32(v26))) = v27 + int32(-1)
				if v27 == int32(0) {
					v38 = F_luaZ_fill(m, v26)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v40 = v38
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v40
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45 + int32(1)
						if v45 < int32(2147483644) {
							return
						} else {
							v51 = m.G3
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							F_luaX_lexerror(m, l0, v51+int32(_a_F_inclinenumber_0), v54)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v33 + int32(1)
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
					v40 = v37
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v40
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45 + int32(1)
					if v45 < int32(2147483644) {
						return
					} else {
						v51 = m.G3
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						F_luaX_lexerror(m, l0, v51+int32(_a_F_inclinenumber_0), v54)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		default:
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45 + int32(1)
			if v45 < int32(2147483644) {
				return
			} else {
				v51 = m.G3
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				F_luaX_lexerror(m, l0, v51+int32(_a_F_inclinenumber_0), v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_incrDecrCommand(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v102 int64
	_ = v102
	var v104 int32
	_ = v104
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = F_lookupKeyWrite(m, v10, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return
L2:
	;
	return
L3:
	;
	v16 = F_checkType(m, l0, v13, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v16 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v21 = F_getLongLongFromObjectOrReply(m, l0, v13, v8+int32(8), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if v21 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
	if int64(-1) < l1 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v41 = v23 + l1
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v41
	if v13 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L9:
	;
	F_addReplyError(m, l0, int32(_a_F_incrDecrCommand_2))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L17
	}
L10:
	;
	if l1 < int64(1) {
		goto L8
	} else {
		goto L14
	}
L11:
	;
	if int64(-1) < v23 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if l1 < int64(-9223372036854775807-1)-v23 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	if v23 < int64(1) {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	if l1 <= v23^int64(9223372036854775807) {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	goto L1
L18:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	F_signalModifiedKey(m, l0, v83, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L30
	}
L19:
	;
	v73 = F_createStringObjectFromLongLongForValue(m, v41)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L2
	} else {
		goto L28
	}
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v45&int32(-8) != int32(8) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v63 = F_createStringObjectFromLongLongForValue(m, v41)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L26
	}
L22:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v50&int32(240) != int32(16) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if base.Ui64(int64(4294967295)) < base.Ui64(v41+int64(2147483648)) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v13
	F_objectSetVal(m, v13, base.I32_wrap_i64(v41))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	goto L18
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v63
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	F_dbReplaceValue(m, v66, v68, v8+int32(4))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	goto L18
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v73
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	F_dbAdd(m, v76, v78, v8+int32(4))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	goto L18
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+28))
	F_notifyKeyspaceEvent(m, int32(8), int32(_a_F_incrDecrCommand_0), v91, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v96 = int32(_a_F_incrDecrCommand_1)
	v98 = *(*int64)(unsafe.Add(mBase, _c_F_incrDecrCommand[0]))
	*(*int64)(unsafe.Add(mBase, _c_F_incrDecrCommand[0])) = v98 + int64(1)
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
	F_addReplyLongLong(m, l0, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	goto L1
}
func F_incrbyCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v12 = F_getLongLongFromObjectOrReply(m, l0, v8, v5+int32(8), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 != 0 {
			m.G0 = v5 + int32(16)
			return
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
			F_incrDecrCommand(m, l0, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				m.G0 = v5 + int32(16)
				return
			}
		}
	}
}
func F_inet_ntop(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int64
	_ = v241
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v597 int32
	_ = v597
	v14 = m.G0
	v16 = v14 - int32(208)
	m.G0 = v16
	if l0 == int32(10) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v16 + int32(208)
	return v597
L2:
	;
	v597 = int32(0)
	goto L1
L3:
	;
	goto L126
L4:
	;
	goto L125
L5:
	;
	v36 = int32(_a_F_inet_ntop_0)
	v37 = int32(12)
	goto L15
L6:
	;
	if l0 != int32(2) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v22
	v31 = F_snprintf(m, l2, l3, int32(_a_F_inet_ntop_1), v16)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	if base.Ui32(v31) < base.Ui32(l3) {
		v597 = l2
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L3
L11:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	v103 = int32(8)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
	v106 = v102<<(uint(v103)%32) | v105
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
	v111 = v107<<(uint(v103)%32) | v110
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	v116 = v112<<(uint(v103)%32) | v115
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	v121 = v117<<(uint(v103)%32) | v120
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
	v126 = v122<<(uint(v103)%32) | v125
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v131 = v127<<(uint(v103)%32) | v130
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v101 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L12:
	;
	v101 = int32(0)
	goto L11
L13:
	;
	v73 = v68
	v74 = v69
	v75 = v70
	goto L23
L14:
	;
	if v58 == int32(0) {
		goto L12
	} else {
		goto L21
	}
L15:
	;
	if (v36|l1)&int32(3) != 0 {
		v68 = l1
		v69 = v36
		v70 = v37
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v45 = l1
	v46 = v36
	v47 = v37
	goto L17
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v50 != v51 {
		v68 = v45
		v69 = v46
		v70 = v47
		goto L13
	} else {
		goto L19
	}
L18:
	;
	goto L14
L19:
	;
	v53 = int32(4)
	v54 = v46 + v53
	v56 = v45 + v53
	v58 = v47 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v58) {
		v45 = v56
		v46 = v54
		v47 = v58
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v68 = v56
	v69 = v54
	v70 = v58
	goto L13
L22:
	;
	v101 = v78 - v79
	goto L11
L23:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v78 != v79 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v81 = int32(1)
	v86 = v75 + int32(-1)
	if v86 == int32(0) {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v73 = v73 + v81
	v74 = v74 + v81
	v75 = v86
	goto L23
L27:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+96)))
	if v204 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(52)))) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(48)))) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(44)))) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(40)))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(36)))) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(32)))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v131
	v199 = F_snprintf(m, v16+int32(96), int32(100), int32(_a_F_inet_ntop_2), v16+int32(16))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L31
	}
L29:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(84)))) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(80)))) = v111
	v146 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(88)))) = v137 | v132<<(uint(v146)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(92)))) = v135 | v136<<(uint(v146)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v131
	v166 = F_snprintf(m, v16+int32(96), int32(100), int32(_a_F_inet_ntop_3), v16+int32(64))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	goto L27
L32:
	;
	v502 = v16 + int32(96)
	if v502&int32(3) == int32(0) {
		v524 = v502
		goto L109
	} else {
		goto L110
	}
L33:
	;
	v208 = int32(0)
	v210 = v208
	v215 = v204
	v216 = int32(2)
	v217 = v208
	goto L34
L34:
	;
	if v210 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	if v318 < int32(4) {
		goto L32
	} else {
		goto L65
	}
L36:
	;
	v323 = v210 + int32(1)
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(96)+v323))))
	if v325 != 0 {
		v210 = v323
		v215 = v325
		v216 = v318
		v217 = v319
		goto L34
	} else {
		goto L64
	}
L37:
	;
	v231 = v16 + int32(96) + v210
	v232 = int32(_a_F_inet_ntop_4)
	v236 = m.G0
	v238 = v236 - int32(32)
	v241 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v238+int32(24)))) = v241
	*(*int64)(unsafe.Add(mBase, uint32(v238+int32(16)))) = v241
	*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v241
	*(*int64)(unsafe.Add(mBase, uint32(v238))) = v241
	v251 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_inet_ntop[0])))
	if v251 != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	if v215&int32(255) != int32(58) {
		v318 = v216
		v319 = v217
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v314 = base.B2i32(v216 < v313)
	if v216 < v313 {
		goto L58
	} else {
		goto L59
	}
L41:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_inet_ntop[1])))
	if v253 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v313 = int32(0)
	goto L40
L43:
	;
	v265 = v232
	v267 = v251
	goto L48
L44:
	;
	v255 = v231
	goto L45
L45:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	if v261 == v251 {
		v255 = v255 + int32(1)
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v313 = v255 - v231
	goto L40
L47:
	;
	goto L46
L48:
	;
	v273 = v238 + int32(base.Ui32(v267)>>(uint(int32(3))%32))&int32(28)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v275 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v273))) = v274 | v275<<(uint(v267)%32)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+1)))
	if v279 != 0 {
		v265 = v265 + v275
		v267 = v279
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v282 == int32(0) {
		v306 = v231
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	v313 = v306 - v231
	goto L40
L52:
	;
	v286 = v231
	v288 = v282
	goto L53
L53:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v238+int32(base.Ui32(v288)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v295)>>(uint(v288)%32))&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v306 = v301
	goto L51
L55:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+1)))
	v301 = v286 + int32(1)
	if v299 != 0 {
		v286 = v301
		v288 = v299
		goto L53
	} else {
		goto L57
	}
L56:
	;
	v306 = v286
	goto L51
L57:
	;
	goto L54
L58:
	;
	v315 = v313
	goto L60
L59:
	;
	v315 = v216
	goto L60
L60:
	;
	if v216 < v313 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v316 = v210
	goto L63
L62:
	;
	v316 = v217
	goto L63
L63:
	;
	v318 = v315
	v319 = v316
	goto L36
L64:
	;
	goto L35
L65:
	;
	v330 = v16 + int32(96) + v319
	v331 = int32(14906)
	*(*uint16)(unsafe.Add(mBase, uint32(v330))) = uint16(v331)
	v333 = int32(2)
	v334 = v330 + v333
	v335 = v330 + v318
	v339 = v210 - (v318 + v319) + v333
	if v334 == v335 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L32
L67:
	;
	goto L66
L68:
	;
	v343 = v339 + v334
	if base.Ui32(int32(0)-v339<<(uint(int32(1))%32)) < base.Ui32(v335-v343) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v353 = (v335 ^ v334) & int32(3)
	if base.Ui32(v335) <= base.Ui32(v334) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v350 = F___memcpy(m, v334, v335, v339)
	mBase = m.M
	goto L66
L71:
	;
	if v459 == int32(0) {
		goto L67
	} else {
		goto L103
	}
L72:
	;
	if base.Ui32(v437) <= base.Ui32(int32(3)) {
		v458 = v436
		v459 = v437
		v460 = v438
		goto L71
	} else {
		goto L99
	}
L73:
	;
	if v353 != 0 {
		v419 = v339
		goto L83
	} else {
		goto L84
	}
L74:
	;
	if v353 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if v334&int32(3) != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v458 = v335
	v459 = v339
	v460 = v334
	goto L71
L77:
	;
	v360 = v335
	v361 = v339
	v362 = v334
	goto L79
L78:
	;
	v436 = v335
	v437 = v339
	v438 = v334
	goto L72
L79:
	;
	if v361 == int32(0) {
		goto L67
	} else {
		goto L81
	}
L81:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	*(*uint8)(unsafe.Add(mBase, uint32(v362))) = uint8(v366)
	v368 = int32(1)
	v369 = v360 + v368
	v371 = v361 + int32(-1)
	v373 = v362 + v368
	if v373&int32(3) == int32(0) {
		v436 = v369
		v437 = v371
		v438 = v373
		goto L72
	} else {
		goto L82
	}
L82:
	;
	v360 = v369
	v361 = v371
	v362 = v373
	goto L79
L83:
	;
	if v419 == int32(0) {
		goto L67
	} else {
		goto L95
	}
L84:
	;
	if v343&int32(3) == int32(0) {
		v399 = v339
		goto L85
	} else {
		goto L86
	}
L85:
	;
	if base.Ui32(v399) <= base.Ui32(int32(3)) {
		v419 = v399
		goto L83
	} else {
		goto L91
	}
L86:
	;
	v384 = v339
	goto L87
L87:
	;
	if v384 == int32(0) {
		goto L67
	} else {
		goto L89
	}
L88:
	;
	v399 = v390
	goto L85
L89:
	;
	v390 = v384 + int32(-1)
	v391 = v334 + v390
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335+v390))))
	*(*uint8)(unsafe.Add(mBase, uint32(v391))) = uint8(v393)
	if v391&int32(3) != 0 {
		v384 = v390
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v406 = v399
	goto L92
L92:
	;
	v410 = v406 + int32(-4)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v335+v410)))
	*(*int32)(unsafe.Add(mBase, uint32(v334+v410))) = v413
	if base.Ui32(int32(3)) < base.Ui32(v410) {
		v406 = v410
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v419 = v410
	goto L83
L94:
	;
	goto L93
L95:
	;
	v426 = v419
	goto L96
L96:
	;
	v430 = v426 + int32(-1)
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335+v430))))
	*(*uint8)(unsafe.Add(mBase, uint32(v334+v430))) = uint8(v433)
	if v430 != 0 {
		v426 = v430
		goto L96
	} else {
		goto L98
	}
L98:
	;
	goto L67
L99:
	;
	v443 = v436
	v444 = v437
	v445 = v438
	goto L100
L100:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v443)))
	*(*int32)(unsafe.Add(mBase, uint32(v445))) = v447
	v449 = int32(4)
	v450 = v443 + v449
	v452 = v445 + v449
	v454 = v444 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v454) {
		v443 = v450
		v444 = v454
		v445 = v452
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v458 = v450
	v459 = v454
	v460 = v452
	goto L71
L102:
	;
	goto L101
L103:
	;
	v465 = v458
	v466 = v459
	v467 = v460
	goto L104
L104:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465))))
	*(*uint8)(unsafe.Add(mBase, uint32(v467))) = uint8(v469)
	v471 = int32(1)
	v476 = v466 + int32(-1)
	if v476 != 0 {
		v465 = v465 + v471
		v466 = v476
		v467 = v467 + v471
		goto L104
	} else {
		goto L106
	}
L105:
	;
	goto L67
L106:
	;
	goto L105
L107:
	;
	if base.Ui32(l3) <= base.Ui32(v557) {
		goto L3
	} else {
		goto L123
	}
L108:
	;
	v557 = v549 - v502
	goto L107
L109:
	;
	v528 = v524
	goto L117
L110:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502))))
	if v510 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v513 = v502
	goto L113
L112:
	;
	v557 = v502 - v502
	goto L107
L113:
	;
	v517 = v513 + int32(1)
	if v517&int32(3) == int32(0) {
		v524 = v517
		goto L109
	} else {
		goto L115
	}
L115:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517))))
	if v522 != 0 {
		v513 = v517
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v549 = v517
	goto L108
L117:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	v537 = int32(-2139062144)
	if (int32(16843008)-v534|v534)&v537 == v537 {
		v528 = v528 + int32(4)
		goto L117
	} else {
		goto L119
	}
L118:
	;
	v543 = v528
	goto L120
L119:
	;
	goto L118
L120:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543))))
	if v547 != 0 {
		v543 = v543 + int32(1)
		goto L120
	} else {
		goto L122
	}
L121:
	;
	v549 = v543
	goto L108
L122:
	;
	goto L121
L123:
	;
	v561 = F___stpcpy(m, l2, v16+int32(96))
	mBase = m.M
	goto L124
L124:
	;
	v597 = l2
	goto L1
L125:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inet_ntop[2])) = int32(5)
	goto L2
L126:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inet_ntop[2])) = int32(51)
	goto L2
}
func F_init_genrand64(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v64 int64
	_ = v64
	*(*int64)(unsafe.Add(mBase, _c_F_init_genrand64[0])) = l0
	v7 = l0
	v8 = int64(1)
	for {
		v11 = int32(3)
		v15 = int64(62)
		v18 = int64(6364136223846793005)
		v20 = (int64(base.Ui64(v7)>>(uint(v15)%64))^v7)*v18 + v8
		*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v8)<<(uint(v11)%32))+uint32(_c_F_init_genrand64[0]))) = v20
		v23 = v8 + int64(1)
		v34 = (int64(base.Ui64(v20)>>(uint(v15)%64))^v20)*v18 + v23
		*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v23)<<(uint(v11)%32))+uint32(_c_F_init_genrand64[0]))) = v34
		v37 = v8 + int64(2)
		v48 = (int64(base.Ui64(v34)>>(uint(v15)%64))^v34)*v18 + v37
		*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v37)<<(uint(v11)%32))+uint32(_c_F_init_genrand64[0]))) = v48
		v51 = v8 + int64(3)
		if v51 == int64(312) {
			break
		} else {
			v64 = (int64(base.Ui64(v48)>>(uint(int64(62))%64))^v48)*int64(6364136223846793005) + v51
			*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v51)<<(uint(int32(3))%32))+uint32(_c_F_init_genrand64[0]))) = v64
			v7 = v64
			v8 = v8 + int64(4)
			continue
		}
		break
	}
	*(*int32)(unsafe.Add(mBase, _c_F_init_genrand64[1])) = int32(312)
	return
}
func F_insertToBucket_VECTOR(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	switch l0 + int32(1) {
	case 0:
		goto L3
	case 1:
		goto L5
	default:
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_VECTOR_0), int32(_a_F_insertToBucket_VECTOR_1), int32(816))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L39
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_VECTOR_2), int32(_a_F_insertToBucket_VECTOR_1), int32(600))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L38
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_VECTOR_3), int32(_a_F_insertToBucket_VECTOR_1), int32(797))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L37
	}
L4:
	;
	if l0&int32(7) != int32(2) {
		goto L3
	} else {
		goto L8
	}
L5:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_VECTOR_4), int32(_a_F_insertToBucket_VECTOR_1), int32(781))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L8:
	;
	v22 = l0 & int32(-8)
	if v22 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	return int32(2) | v108
L10:
	;
	v104 = F_pvInsertAt(m, v22, l1, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L6
	} else {
		goto L35
	}
L11:
	;
	v102 = base.I32_wrap_i64(v25) & int32(1073741823)
	goto L10
L12:
	;
	v91 = F_pvInsertAt(m, v22, l1, l2)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L32
	}
L13:
	;
	if l2 < int32(0) {
		goto L11
	} else {
		goto L31
	}
L14:
	;
	if int32(-1) < l2 {
		goto L12
	} else {
		goto L30
	}
L15:
	;
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	if base.Ui64(v25&int64(1073741823)) < base.Ui64(int64(127)) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v31 = F_hashtableCreate(m, int32(_a_F_insertToBucket_VECTOR_5))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	if v33&int64(1073741823) == int64(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_valkey_free(m, v22)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L25
	}
L19:
	;
	v41 = int32(0)
	v45 = v33
	goto L20
L20:
	;
	if base.Ui32(base.I32_wrap_i64(v45)&int32(1073741823)) <= base.Ui32(v41) {
		goto L2
	} else {
		goto L22
	}
L21:
	;
	goto L18
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(8)+v41<<(uint(int32(2))%32))))
	v55 = F_hashtableAdd(m, v31, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v58 = v41 + int32(1)
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	if base.Ui32(v58) < base.Ui32(base.I32_wrap_i64(v59)&int32(1073741823)) {
		v41 = v58
		v45 = v59
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v72 = F_hashtableAdd(m, v31, l1)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	if v31 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_VECTOR_0), int32(_a_F_insertToBucket_VECTOR_1), int32(816))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L29
	}
L28:
	;
	return int32(4) | v31
L29:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v102 = int32(0)
	goto L10
L31:
	;
	goto L12
L32:
	;
	if v91 != 0 {
		v108 = v91
		goto L9
	} else {
		goto L33
	}
L33:
	;
	F__serverAssert(m, int32(_a_F_insertToBucket_VECTOR_0), int32(_a_F_insertToBucket_VECTOR_1), int32(816))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	if v104 == int32(0) {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v108 = v104
	goto L9
L37:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_integerCallback(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 < int32(0) {
		v31 = v6
	} else {
		v12 = l0 + v7<<(uint(int32(2))%32)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
		v14 = int32(1)
		v15 = v13 + v14
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v15
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1040))
		if v17 != v14 {
			v31 = v6
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v22))) = base.F64_convert_i32_u(v15)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v26 + int32(16)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v31 = v30
		}
	}
	v35 = F_lua_checkstack(m, v31, int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		return
	} else {
		if v35 != 0 {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v42))) = base.F64_convert_i64_s(l1)
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v46 + int32(16)
			F_processCollectionElementEnd(m, l0)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				return
			}
		} else {
			F__serverPanic_2(m, int32(867))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_isInsideYieldingLongCommand(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v1 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_isInsideYieldingLongCommand[0]))
	if v5 == v1 {
		v13 = v1
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
		v13 = int32(base.Ui32(v8)>>(uint(int32(3))%32)) & int32(1)
	}
	v14 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_isInsideYieldingLongCommand[1]))
	return base.B2i32(v13|v15 != v14)
}
func F_isPausedActions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_isPausedActions[0]))
	return v3 & l0
}
func F_isReplicatedClient(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v3 = int32(1)
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v4&v3 != 0 {
		v14 = v3
		return v14
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
		if v7 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v14 = base.B2i32(v10 == int32(1))
			return v14
		} else {
			return int32(0)
		}
	}
}
func F_is_lost_conn(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	v3 = int32(1)
	if base.Ui32(l0+int32(-14)) < base.Ui32(int32(2)) {
		v12 = v3
	} else {
		if l0 == int32(53) {
			v12 = v3
		} else {
			v12 = base.B2i32(l0 == int32(64))
		}
	}
	return v12
}
func F_isalnum(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0+int32(-48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(l0|int32(32)+int32(-97)) < base.Ui32(int32(26)))
}
func F_ispunct(m *base.Module, l0 int32) int32 {
	var v21 int32
	_ = v21
	if base.Ui32(int32(93)) < base.Ui32(l0+int32(-33)) {
		v21 = int32(0)
	} else {
		v21 = base.B2i32(base.B2i32(base.Ui32(l0+int32(-48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(l0|int32(32)+int32(-97)) < base.Ui32(int32(26))) == int32(0))
	}
	return v21
}
func F_iswblank(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0 == int32(32)) | base.B2i32(l0 == int32(9))
}
func F_iswspace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v4 = int32(_a_F_iswspace_0)
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return int32(0)
L3:
	;
	return base.B2i32(v30 != int32(0))
L4:
	;
	v22 = F_wcslen(m, v4)
	mBase = m.M
	v30 = v4 + v22<<(uint(int32(2))%32)
	goto L3
L5:
	;
	v9 = v4
	goto L7
L6:
	;
	if v13 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v13 == int32(0) {
		goto L6
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	if v13 != l0 {
		v9 = v9 + int32(4)
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v21 = v9
	goto L13
L12:
	;
	v21 = int32(0)
	goto L13
L13:
	;
	v30 = v21
	goto L3
}
