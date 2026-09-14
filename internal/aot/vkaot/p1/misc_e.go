package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F__exit(m *base.Module, l0 int32) {
	F__Exit(m, l0)
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_echoCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	F_addReplyBulk(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_enableParseExactReplyTypeFlag(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3&int32(2) == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3 | int32(8)
		return
	} else {
		F__serverAssert(m, int32(_a_F_enableParseExactReplyTypeFlag_0), int32(_a_F_enableParseExactReplyTypeFlag_1), int32(626))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_enterExecutionUnit(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v41 int64
	_ = v41
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_enterExecutionUnit[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_enterExecutionUnit[0])) = v7 + int32(1)
	if l0 == v3 {
	} else {
		if v7 != 0 {
		} else {
			if l1 != int64(0) {
				v16 = l1
			} else {
				v15 = F_ustime(m)
				mBase = m.M
				v16 = v15
			}
			v17 = int32(0)
			*(*int64)(unsafe.Add(mBase, _c_F_enterExecutionUnit[1])) = v16
			v20 = int64(1000)
			v21 = base.I64_div_s(v16, v20)
			*(*int64)(unsafe.Add(mBase, _c_F_enterExecutionUnit[2])) = v21
			v25 = base.I64_div_s(v16, int64(1000000))
			*(*int64)(unsafe.Add(mBase, _c_F_enterExecutionUnit[3])) = v25
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_enterExecutionUnit[4]))
			v32 = int32(base.Ui32(v28&int32(2)) >> (uint(int32(1)) % 32))
			*(*uint8)(unsafe.Add(mBase, _c_F_enterExecutionUnit[5])) = uint8(v32)
			v37 = base.I64_div_s(v21, int64(60000))
			*(*uint16)(unsafe.Add(mBase, _c_F_enterExecutionUnit[6])) = uint16(v37)
			v41 = base.I64_div_s(v21, v20)
			*(*int32)(unsafe.Add(mBase, _c_F_enterExecutionUnit[7])) = base.I32_wrap_i64(v41) & int32(16777215)
			v46 = int32(0)
			v48 = *(*int64)(unsafe.Add(mBase, _c_F_enterExecutionUnit[2]))
			*(*int64)(unsafe.Add(mBase, _c_F_enterExecutionUnit[8])) = v48
		}
	}
	return
}
func F_enumConfigInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2))) = v3
	return
}
func F_enumConfigRewrite(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v12&int32(256) == int32(0) {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		v23 = v12
		v24 = v22
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v28 = F_configEnumGetName(m, v25, v24, v23&int32(8))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			v30 = F_sdsempty(m)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v28
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
				v35 = F_sdscatfmt(m, v30, int32(_a_F_enumConfigRewrite_0), v10)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					F_sdsfree(m, v28)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v41 = F_rewriteConfigRewriteLine(m, l2, l1, v35, base.B2i32(v24 != v39))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							m.G0 = v10 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		v18 = F_getModuleEnumConfig(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v23 = v20
			v24 = v18
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v28 = F_configEnumGetName(m, v25, v24, v23&int32(8))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v30 = F_sdsempty(m)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
					v35 = F_sdscatfmt(m, v30, int32(_a_F_enumConfigRewrite_0), v10)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_sdsfree(m, v28)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v41 = F_rewriteConfigRewriteLine(m, l2, l1, v35, base.B2i32(v24 != v39))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_evalMemory(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(0)
	F_scriptingEngineManagerForEachEngine(m, int32(521), v5+int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
		m.G0 = v5 + int32(16)
		return v16
	}
}
func F_evalRemoveScriptsFromEngine(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int64
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_evalRemoveScriptsFromEngine[0]))
	v9 = F_dictGetSafeIterator(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L4
L3:
	;
	F_dictReleaseIterator(m, v9)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L54
	}
L4:
	;
	v24 = v9 + int32(20)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v120 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L7:
	;
	v31 = v24
	v32 = v28
	goto L10
L8:
	;
	v28 = int32(1)
	goto L7
L9:
	;
	v28 = int32(0)
	goto L7
L10:
	;
	switch v32 {
	case 0:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v32 = int32(0)
	goto L10
L13:
	;
	goto L6
L14:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v112
	if v112 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v36 != int32(-1) {
		v75 = v36
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v76 = int32(1)
	v77 = v75 + v76
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v77
	v79 = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82+v83+int32(26)))))
	if v87 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v40 != 0 {
		v75 = int32(-1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v42 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	if v69 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v41)+16)))
	v50 = int64(*(*int8)(unsafe.Add(mBase, uint32(v41)+27)))
	v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v41)+8)))
	v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v41)+12)))
	v53 = int64(*(*int8)(unsafe.Add(mBase, uint32(v41)+26)))
	v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v41)+4)))
	v55 = F_wangHash64(m, v54)
	mBase = m.M
	v57 = F_wangHash64(m, v53+v55)
	mBase = m.M
	v59 = F_wangHash64(m, v52+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v51+v59)
	mBase = m.M
	v63 = F_wangHash64(m, v50+v61)
	mBase = m.M
	v65 = F_wangHash64(m, v49+v63)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v68 = v67
	goto L19
L21:
	;
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+24)))
	v47 = v45 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+24)) = uint16(v47)
	v68 = v41
	goto L19
L22:
	;
	v75 = v69 + int32(-1)
	goto L16
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v75 = v72
	goto L16
L24:
	;
	v102 = int32(2)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v82+v100<<(uint(v102)%32)+int32(4))))
	v31 = v107 + v101<<(uint(v102)%32)
	v32 = int32(1)
	goto L10
L25:
	;
	v91 = v79
	goto L27
L26:
	;
	v91 = v76 << (uint(v87) % 32)
	goto L27
L27:
	;
	if v77 < v91 {
		v100 = v83
		v101 = v77
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v83 != 0 {
		v120 = v79
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v93 == int32(-1) {
		v120 = v79
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(4294967296)
	v100 = int32(1)
	v101 = int32(0)
	goto L24
L31:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v116
	v120 = v112
	goto L13
L32:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	goto L33
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v127 != l0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	goto L35
L35:
	;
	v136 = v129 + int32(-1)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v139 = v137 & int32(7)
	switch v139 {
	case 0:
		goto L42
	case 1:
		v145 = int32(4)
		goto L37
	case 2:
		goto L41
	case 3:
		goto L40
	case 4:
		goto L39
	default:
		goto L38
	}
L36:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	v171 = F_getStringObjectSdsUsedMemory(m, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L49
	}
L37:
	;
	switch v139 {
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
		v165 = int32(0)
		goto L43
	}
L38:
	;
	v145 = int32(1)
	goto L37
L39:
	;
	v145 = int32(18)
	goto L37
L40:
	;
	v145 = int32(10)
	goto L37
L41:
	;
	v145 = int32(6)
	goto L37
L42:
	;
	v140 = F_zmalloc_usable_size(m, v136)
	mBase = m.M
	v169 = v140
	goto L36
L43:
	;
	v169 = v145 + v165
	goto L36
L44:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v129+int32(-9))))
	v165 = v164
	goto L43
L45:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v129+int32(-5))))
	v169 = v145 + v160
	goto L36
L46:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129+int32(-3)))))
	v169 = v145 + v156
	goto L36
L47:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+int32(-2)))))
	v169 = v145 + v152
	goto L36
L48:
	;
	v169 = v145 + int32(base.Ui32(v137)>>(uint(int32(3))%32))
	goto L36
L49:
	;
	v173 = int32(0)
	v175 = *(*int64)(unsafe.Add(mBase, _c_F_evalRemoveScriptsFromEngine[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_evalRemoveScriptsFromEngine[1])) = v175 - base.I64_extend_i32_u(v171+v169)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v126)+24))
	if v180 == v173 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_evalRemoveScriptsFromEngine[0]))
	v189 = F_dictDelete(m, v188, v129)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_evalRemoveScriptsFromEngine[2]))
	F_listDelNode(m, v184, v180)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	goto L4
L54:
	;
	return
}
func F_evalShaRoCommand(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_evalShaCommand(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_eventLoopCbReadable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	m.T0[v8].(func(*base.Module, int32, int32, int32))(m, l1, v5, l3&int32(3))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
func F_eventLoopCbWritable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	m.T0[v8].(func(*base.Module, int32, int32, int32))(m, l1, v5, l3&int32(3))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
func F_evictionTimeProc(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	v6 = F_performEvictions(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		if v6 == int32(1) {
			v16 = int64(0)
		} else {
			v12 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _c_F_evictionTimeProc[0])) = uint8(v12)
			v16 = int64(-1)
		}
		return v16
	}
}
func F_exp2reg(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	F_discharge2reg(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v14 != int32(10) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v90 == v91 {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v17 == int32(-1) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v20 == int32(-1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v17
	goto L3
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v29 = v20
	goto L9
L8:
	;
	v54 = v17 + (v29 ^ int32(-1))
	v56 = v54 >> (uint(int32(31)) % 32)
	if base.Ui32(v54^v56-v56) < base.Ui32(int32(131072)) {
		v69 = v39
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v38 = v24 + v29<<(uint(int32(2))%32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v43 = int32(base.Ui32(v39)>>(uint(int32(14))%32)) + int32(-131071)
	if v43 == int32(-1) {
		goto L8
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v48 = v29 + v43 + int32(1)
	if v48 != int32(-1) {
		v29 = v48
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v54<<(uint(int32(14))%32) | v69&int32(16383) + int32(2147467264)
	goto L3
L14:
	;
	v61 = m.G3
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v62, v61+int32(_a_F_exp2reg_0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v69 = v67
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(12)
	return
L17:
	;
	v93 = int32(-1)
	if v90 == v93 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v392
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_patchlistaux(m, l0, v394, v392, l2, v388)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L69
	}
L19:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v219 == int32(10) {
		v298 = int32(-1)
		goto L42
	} else {
		goto L43
	}
L20:
	;
	v154 = int32(-1)
	if v91 == v154 {
		v384 = v93
		v388 = v154
		goto L18
	} else {
		goto L31
	}
L21:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v102 = v90
	goto L22
L22:
	;
	v111 = v97 + v102<<(uint(int32(2))%32)
	if v102 < int32(1) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L20
L24:
	;
	if v126&int32(63) != int32(27) {
		goto L19
	} else {
		goto L28
	}
L25:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v126 = v125
	goto L24
L26:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111+int32(-4))))
	v117 = m.G400
	v121 = int32(*(*int8)(unsafe.Add(mBase, uint32(v117+v116&int32(63)))))
	if v121 < int32(0) {
		v126 = v116
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v135 = int32(base.Ui32(v131)>>(uint(int32(14))%32)) + int32(-131071)
	if v135 == int32(-1) {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v140 = v102 + v135 + int32(1)
	if v140 != int32(-1) {
		v102 = v140
		goto L22
	} else {
		goto L30
	}
L30:
	;
	goto L23
L31:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v167 = v91
	goto L32
L32:
	;
	v172 = v158 + v167<<(uint(int32(2))%32)
	if v167 < int32(1) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v187&int32(63) != int32(27) {
		goto L19
	} else {
		goto L38
	}
L35:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v187 = v186
	goto L34
L36:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v172+int32(-4))))
	v178 = m.G400
	v182 = int32(*(*int8)(unsafe.Add(mBase, uint32(v178+v177&int32(63)))))
	if v182 < int32(0) {
		v187 = v177
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v192 = int32(-1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v197 = int32(base.Ui32(v193)>>(uint(int32(14))%32)) + int32(-131071)
	if v197 != v192 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v201 = int32(-1)
	v204 = v167 + v197 + int32(1)
	if v204 != v201 {
		v167 = v204
		goto L32
	} else {
		goto L41
	}
L40:
	;
	v384 = v192
	v388 = int32(-1)
	goto L18
L41:
	;
	v384 = v192
	v388 = v201
	goto L18
L42:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v300
	v303 = l2 << (uint(int32(6)) % 32)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+8))
	v308 = F_luaK_code(m, l0, v303|int32(16386), v307)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L56
	}
L43:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(-1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
	v228 = F_luaK_code(m, l0, int32(2147450902), v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v222 == int32(-1) {
		v298 = v228
		goto L42
	} else {
		goto L45
	}
L45:
	;
	if v228 != int32(-1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v242 = v228
	goto L49
L47:
	;
	v298 = v222
	goto L42
L48:
	;
	v265 = v222 + (v242 ^ int32(-1))
	v267 = v265 >> (uint(int32(31)) % 32)
	if base.Ui32(v265^v267-v267) < base.Ui32(int32(131072)) {
		v280 = v250
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v249 = v235 + v242<<(uint(int32(2))%32)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v254 = int32(base.Ui32(v250)>>(uint(int32(14))%32)) + int32(-131071)
	if v254 == int32(-1) {
		goto L48
	} else {
		goto L51
	}
L50:
	;
	goto L48
L51:
	;
	v259 = v242 + v254 + int32(1)
	if v259 != int32(-1) {
		v242 = v259
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = v265<<(uint(int32(14))%32) | v280&int32(16383) + int32(2147467264)
	v298 = v228
	goto L42
L54:
	;
	v272 = m.G3
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v273, v272+int32(_a_F_exp2reg_0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v280 = v278
	goto L53
L56:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v310
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+8))
	v316 = F_luaK_code(m, l0, v303|int32(_a_F_exp2reg_1), v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v318
	if v298 == int32(-1) {
		v384 = v316
		v388 = v308
		goto L18
	} else {
		goto L58
	}
L58:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v322 == int32(-1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v298
	v384 = v316
	v388 = v308
	goto L18
L60:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+12))
	v331 = v322
	goto L62
L61:
	;
	v356 = v298 + (v331 ^ int32(-1))
	v358 = v356 >> (uint(int32(31)) % 32)
	if base.Ui32(v356^v358-v358) < base.Ui32(int32(131072)) {
		v370 = v341
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v340 = v326 + v331<<(uint(int32(2))%32)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v345 = int32(base.Ui32(v341)>>(uint(int32(14))%32)) + int32(-131071)
	if v345 == int32(-1) {
		goto L61
	} else {
		goto L64
	}
L63:
	;
	goto L61
L64:
	;
	v350 = v331 + v345 + int32(1)
	if v350 != int32(-1) {
		v331 = v350
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v340))) = v356<<(uint(int32(14))%32) | v370&int32(16383) + int32(2147467264)
	v384 = v316
	v388 = v308
	goto L18
L67:
	;
	v363 = m.G3
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_luaX_syntaxerror(m, v364, v363+int32(_a_F_exp2reg_0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v370 = v369
	goto L66
L69:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_patchlistaux(m, l0, v397, v392, l2, v384)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L16
}
func F_expiretimeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v9 = F_lookupKeyReadWithFlags(m, v5, v7, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v9 != 0 {
			v14 = int64(-1)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			if v18&int32(1) == int32(0) {
				v29 = v14
			} else {
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v9+(v18&int32(4)^int32(12)))))
				v29 = v28
			}
			if v29 == int64(-1) {
				v40 = v14
			} else {
				v32 = int64(0)
				if v32 < v29 {
					v35 = v29
				} else {
					v35 = v32
				}
				v39 = base.I64_div_u_s(v35+int64(500), int64(1000))
				v40 = v39
			}
			F_addReplyLongLong(m, l0, v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				return
			}
		} else {
			F_addReplyLongLong(m, l0, int64(-2))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_extractUnitOrReply(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
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
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	v3 = F_objectGetVal(m, l1)
	mBase = m.M
	v4 = int32(_a_F_extractUnitOrReply_0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v7 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v45 = int32(_a_F_extractUnitOrReply_1)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v48 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	if v39-v41 != 0 {
		goto L1
	} else {
		goto L14
	}
L3:
	;
	v39 = F_tolower(m, v35)
	mBase = m.M
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v41 = F_tolower(m, v40)
	mBase = m.M
	goto L2
L4:
	;
	v9 = v3
	v10 = v4
	v11 = v7
	goto L7
L5:
	;
	v35 = int32(0)
	v36 = v4
	goto L3
L6:
	;
	v35 = v32 & int32(255)
	v36 = v31
	goto L3
L7:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v13 == int32(0) {
		v31 = v10
		v32 = v11
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v31 = v25
	v32 = int32(0)
	goto L6
L9:
	;
	v17 = v11 & int32(255)
	if v17 == v13 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v24 = int32(1)
	v25 = v10 + v24
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v26 != 0 {
		v9 = v9 + v24
		v10 = v25
		v11 = v26
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v19 = F_tolower(m, v17)
	mBase = m.M
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v21 = F_tolower(m, v20)
	mBase = m.M
	if v19 == v21 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v31 = v10
	v32 = v23
	goto L6
L13:
	;
	goto L8
L14:
	;
	return float64(1)
L15:
	;
	v86 = int32(_a_F_extractUnitOrReply_2)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v89 != 0 {
		goto L32
	} else {
		goto L33
	}
L16:
	;
	if v80-v82 != 0 {
		goto L15
	} else {
		goto L28
	}
L17:
	;
	v80 = F_tolower(m, v76)
	mBase = m.M
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v82 = F_tolower(m, v81)
	mBase = m.M
	goto L16
L18:
	;
	v50 = v3
	v51 = v45
	v52 = v48
	goto L21
L19:
	;
	v76 = int32(0)
	v77 = v45
	goto L17
L20:
	;
	v76 = v73 & int32(255)
	v77 = v72
	goto L17
L21:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v54 == int32(0) {
		v72 = v51
		v73 = v52
		goto L20
	} else {
		goto L23
	}
L22:
	;
	v72 = v66
	v73 = int32(0)
	goto L20
L23:
	;
	v58 = v52 & int32(255)
	if v58 == v54 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v65 = int32(1)
	v66 = v51 + v65
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v67 != 0 {
		v50 = v50 + v65
		v51 = v66
		v52 = v67
		goto L21
	} else {
		goto L27
	}
L25:
	;
	v60 = F_tolower(m, v58)
	mBase = m.M
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v62 = F_tolower(m, v61)
	mBase = m.M
	if v60 == v62 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v72 = v51
	v73 = v64
	goto L20
L27:
	;
	goto L22
L28:
	;
	return float64(1000)
L29:
	;
	v127 = int32(_a_F_extractUnitOrReply_3)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v130 != 0 {
		goto L46
	} else {
		goto L47
	}
L30:
	;
	if v121-v123 != 0 {
		goto L29
	} else {
		goto L42
	}
L31:
	;
	v121 = F_tolower(m, v117)
	mBase = m.M
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v123 = F_tolower(m, v122)
	mBase = m.M
	goto L30
L32:
	;
	v91 = v3
	v92 = v86
	v93 = v89
	goto L35
L33:
	;
	v117 = int32(0)
	v118 = v86
	goto L31
L34:
	;
	v117 = v114 & int32(255)
	v118 = v113
	goto L31
L35:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v95 == int32(0) {
		v113 = v92
		v114 = v93
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v113 = v107
	v114 = int32(0)
	goto L34
L37:
	;
	v99 = v93 & int32(255)
	if v99 == v95 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v106 = int32(1)
	v107 = v92 + v106
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	if v108 != 0 {
		v91 = v91 + v106
		v92 = v107
		v93 = v108
		goto L35
	} else {
		goto L41
	}
L39:
	;
	v101 = F_tolower(m, v99)
	mBase = m.M
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v103 = F_tolower(m, v102)
	mBase = m.M
	if v101 == v103 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v113 = v92
	v114 = v105
	goto L34
L41:
	;
	goto L36
L42:
	;
	return float64(0.3048)
L43:
	;
	F_addReplyError(m, l0, int32(_a_F_extractUnitOrReply_4))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L57
	} else {
		goto L58
	}
L44:
	;
	if v162-v164 != 0 {
		goto L43
	} else {
		goto L56
	}
L45:
	;
	v162 = F_tolower(m, v158)
	mBase = m.M
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v164 = F_tolower(m, v163)
	mBase = m.M
	goto L44
L46:
	;
	v132 = v3
	v133 = v127
	v134 = v130
	goto L49
L47:
	;
	v158 = int32(0)
	v159 = v127
	goto L45
L48:
	;
	v158 = v155 & int32(255)
	v159 = v154
	goto L45
L49:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v136 == int32(0) {
		v154 = v133
		v155 = v134
		goto L48
	} else {
		goto L51
	}
L50:
	;
	v154 = v148
	v155 = int32(0)
	goto L48
L51:
	;
	v140 = v134 & int32(255)
	if v140 == v136 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v147 = int32(1)
	v148 = v133 + v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v149 != 0 {
		v132 = v132 + v147
		v133 = v148
		v134 = v149
		goto L49
	} else {
		goto L55
	}
L53:
	;
	v142 = F_tolower(m, v140)
	mBase = m.M
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v144 = F_tolower(m, v143)
	mBase = m.M
	if v142 == v144 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v154 = v133
	v155 = v146
	goto L48
L55:
	;
	goto L50
L56:
	;
	return float64(1609.34)
L57:
	;
	return float64(0)
L58:
	;
	return float64(-1)
}
