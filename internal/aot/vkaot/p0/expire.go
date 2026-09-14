package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_convertExpireArgumentToUnixTime(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(-1)
	v17 = F_getLongLongFromObjectOrReply(m, l0, l1, v11+int32(8), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 != 0 {
			v46 = v13
			m.G0 = v11 + int32(16)
			return v46
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
			if int64(-1) < v21 {
				if l3 != 0 {
					v33 = v21
					if base.Ui64(v33) <= base.Ui64(l2^int64(9223372036854775807)) {
						v39 = v33 + l2
						*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v39
						if l4 != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(l4))) = v39
							v46 = int32(0)
							m.G0 = v11 + int32(16)
							return v46
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, _c_F_convertExpireArgumentToUnixTime[0]))
							if v42 != 0 {
								F__serverAssert(m, int32(_a_F_convertExpireArgumentToUnixTime_0), int32(_a_F_convertExpireArgumentToUnixTime_1), int32(746))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l4))) = v39
								v46 = int32(0)
								m.G0 = v11 + int32(16)
								return v46
							}
						}
					} else {
						F_addReplyErrorExpireTime(m, l0)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v46 = v13
							m.G0 = v11 + int32(16)
							return v46
						}
					}
				} else {
					if base.Ui64(v21) < base.Ui64(int64(9223372036854776)) {
						v31 = v21 * int64(1000)
						*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v31
						v33 = v31
						if base.Ui64(v33) <= base.Ui64(l2^int64(9223372036854775807)) {
							v39 = v33 + l2
							*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v39
							if l4 != 0 {
								*(*int64)(unsafe.Add(mBase, uint32(l4))) = v39
								v46 = int32(0)
								m.G0 = v11 + int32(16)
								return v46
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, _c_F_convertExpireArgumentToUnixTime[0]))
								if v42 != 0 {
									F__serverAssert(m, int32(_a_F_convertExpireArgumentToUnixTime_0), int32(_a_F_convertExpireArgumentToUnixTime_1), int32(746))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l4))) = v39
									v46 = int32(0)
									m.G0 = v11 + int32(16)
									return v46
								}
							}
						} else {
							F_addReplyErrorExpireTime(m, l0)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v46 = v13
								m.G0 = v11 + int32(16)
								return v46
							}
						}
					} else {
						F_addReplyErrorExpireTime(m, l0)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v46 = v13
							m.G0 = v11 + int32(16)
							return v46
						}
					}
				}
			} else {
				F_addReplyErrorExpireTime(m, l0)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v46 = v13
					m.G0 = v11 + int32(16)
					return v46
				}
			}
		}
	}
}
func F_expireGenericCommand(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int64
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int64
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = F_parseExtendedExpireArgumentsOrReply(m, l0, v12+int32(4), v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return
L2:
	;
	return
L3:
	;
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = F_getLongLongFromObjectOrReply(m, l0, v15, v12+int32(8), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v27 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	if l2 != 0 {
		v39 = v29
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v39 <= l1^int64(9223372036854775807) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	if base.Ui64(int64(-18446744073709552)) < base.Ui64(v29+int64(-9223372036854776)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v37 = v29 * int64(1000)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v37
	v39 = v37
	goto L7
L10:
	;
	F_addReplyErrorExpireTime(m, l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	goto L1
L12:
	;
	v45 = v39 + l1
	v46 = int64(0)
	if v46 < v45 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	F_addReplyErrorExpireTime(m, l0)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	goto L1
L15:
	;
	v49 = v45
	goto L17
L16:
	;
	v49 = v46
	goto L17
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v52 = F_lookupKeyWrite(m, v51, v16)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L19
	}
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v58 != 0 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	if v52 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_expireGenericCommand[7]))
	F_addReply(m, l0, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	goto L1
L22:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_expireGenericCommand[0]))
	if v126 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v63&int32(1) == int32(0) {
		v74 = int64(-1)
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	v122 = v59
	goto L22
L25:
	;
	if v58&int32(1) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L25
L27:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v52+(v63&int32(4)^int32(12)))))
	v74 = v73
	goto L26
L28:
	;
	if v58&int32(2) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	if v74 == int64(-1) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_expireGenericCommand[7]))
	F_addReply(m, l0, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	goto L1
L32:
	;
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	if v58&int32(4) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	if v74 != int64(-1) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_expireGenericCommand[7]))
	F_addReply(m, l0, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	goto L1
L36:
	;
	if v58&int32(8) == int32(0) {
		v122 = v95
		goto L22
	} else {
		goto L40
	}
L37:
	;
	if base.B2i32(v74 == int64(-1))|base.B2i32(v95 <= v74) == int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_expireGenericCommand[7]))
	F_addReply(m, l0, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	goto L1
L40:
	;
	if v74 == int64(-1) {
		v122 = v95
		goto L22
	} else {
		goto L41
	}
L41:
	;
	if v95 < v74 {
		v122 = v95
		goto L22
	} else {
		goto L42
	}
L42:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_expireGenericCommand[7]))
	F_addReply(m, l0, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	goto L1
L44:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	v155 = F_setExpire(m, l0, v153, v16, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L2
	} else {
		goto L56
	}
L45:
	;
	if v144 == int32(0) {
		goto L44
	} else {
		goto L53
	}
L46:
	;
	goto L45
L47:
	;
	v132 = int32(0)
	v133 = F_commandTimeSnapshot(m)
	mBase = m.M
	if v133 < v122 {
		v144 = v132
		goto L46
	} else {
		goto L50
	}
L48:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+216))
	if v130 != 0 {
		v144 = int32(0)
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v135 = int32(_a_F_expireGenericCommand_0)
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_expireGenericCommand[1]))
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_expireGenericCommand[2]))
	if v138 != 0 {
		v144 = v132
		goto L46
	} else {
		goto L51
	}
L51:
	;
	if v136 != 0 {
		v144 = v132
		goto L46
	} else {
		goto L52
	}
L52:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_expireGenericCommand[3]))
	v144 = base.B2i32(v140 == int32(0))
	goto L46
L53:
	;
	F_deleteExpiredKeyFromOverwriteAndPropagate(m, l0, v16)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_expireGenericCommand[5]))
	F_addReply(m, l0, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	goto L1
L56:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v157, v16)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_expireGenericCommand_1), v16, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	v166 = int32(_a_F_expireGenericCommand_0)
	v168 = *(*int64)(unsafe.Add(mBase, _c_F_expireGenericCommand[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_expireGenericCommand[4])) = v168 + int64(1)
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_expireGenericCommand[5]))
	F_addReply(m, l0, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+48))
	if v177 == int32(199) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if l1 != int64(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_expireGenericCommand[6]))
	F_rewriteClientCommandArgument(m, l0, int32(0), v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	v189 = F_createStringObjectFromLongLong(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L66
	}
L64:
	;
	if l2 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	F_rewriteClientCommandArgument(m, l0, int32(2), v189)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	F_decrRefCount(m, v189)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	goto L1
}
func F_expireIfNeededWithDictIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_expireIfNeededWithDictIndex[0]))
	if v16 != 0 {
		v147 = v6
		m.G0 = v12 + int32(16)
		return v147
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_expireIfNeededWithDictIndex[1]))
		if l2 == int32(0) {
			if v18 != 0 {
				v147 = v6
				m.G0 = v12 + int32(16)
				return v147
			} else {
				v59 = F_objectGetVal(m, l1)
				mBase = m.M
				v60 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v60
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v66 = F_kvstoreHashtableFind(m, v63, l4, v59, v12+int32(12))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
					if v70 != 0 {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
						if v75&int32(1) == int32(0) {
							v86 = int64(-1)
						} else {
							v85 = *(*int64)(unsafe.Add(mBase, uint32(v70+(v75&int32(4)^int32(12)))))
							v86 = v85
						}
						v87 = v86
					} else {
						v87 = int64(-1)
					}
					if int64(0) <= v87 {
						v91 = F_commandTimeSnapshot(m)
						mBase = m.M
						v93 = base.B2i32(v87 < v91)
					} else {
						v93 = int32(0)
					}
					if v93 == int32(0) {
						v147 = v60
						m.G0 = v12 + int32(16)
						return v147
					} else {
						v99 = F_getExpirationPolicyWithFlags(m, l3)
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							if base.Ui32(v99) < base.Ui32(int32(2)) {
								v147 = v99
								m.G0 = v12 + int32(16)
								return v147
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								if v103&int32(-8) != int32(-16) {
									F_deleteExpiredKeyAndPropagateWithDictIndex(m, l0, l1, l4)
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										v147 = int32(2)
										m.G0 = v12 + int32(16)
										return v147
									}
								} else {
									v108 = F_objectGetVal(m, l1)
									mBase = m.M
									v110 = F_objectGetVal(m, l1)
									mBase = m.M
									v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+int32(-1)))))
									switch v113 & int32(7) {
									case 0:
										v132 = int32(base.Ui32(v113) >> (uint(int32(3)) % 32))
									case 1:
										v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+int32(-3)))))
										v132 = v118
									case 2:
										v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110+int32(-5)))))
										v132 = v121
									case 3:
										v124 = *(*int32)(unsafe.Add(mBase, uint32(v110+int32(-9))))
										v132 = v124
									case 4:
										v127 = *(*int32)(unsafe.Add(mBase, uint32(v110+int32(-17))))
										v132 = v127
									default:
										v132 = int32(0)
									}
									v133 = F_createStringObject_1(m, v108, v132)
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return int32(0)
									} else {
										F_deleteExpiredKeyAndPropagateWithDictIndex(m, l0, v133, l4)
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											F_decrRefCount(m, v133)
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return int32(0)
											} else {
												v147 = int32(2)
												m.G0 = v12 + int32(16)
												return v147
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
			if v18 != 0 {
				v147 = v6
				m.G0 = v12 + int32(16)
				return v147
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				if v24&int32(1) == int32(0) {
					v35 = int64(-1)
				} else {
					v34 = *(*int64)(unsafe.Add(mBase, uint32(l2+(v24&int32(4)^int32(12)))))
					v35 = v34
				}
				if int64(0) <= v35 {
					v39 = F_commandTimeSnapshot(m)
					mBase = m.M
					v41 = base.B2i32(v35 < v39)
				} else {
					v41 = int32(0)
				}
				if v41 == int32(0) {
					v147 = v6
					m.G0 = v12 + int32(16)
					return v147
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, _c_F_expireIfNeededWithDictIndex[2]))
					if v45 != 0 {
						v99 = F_getExpirationPolicyWithFlags(m, l3)
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							if base.Ui32(v99) < base.Ui32(int32(2)) {
								v147 = v99
								m.G0 = v12 + int32(16)
								return v147
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								if v103&int32(-8) != int32(-16) {
									F_deleteExpiredKeyAndPropagateWithDictIndex(m, l0, l1, l4)
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										v147 = int32(2)
										m.G0 = v12 + int32(16)
										return v147
									}
								} else {
									v108 = F_objectGetVal(m, l1)
									mBase = m.M
									v110 = F_objectGetVal(m, l1)
									mBase = m.M
									v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+int32(-1)))))
									switch v113 & int32(7) {
									case 0:
										v132 = int32(base.Ui32(v113) >> (uint(int32(3)) % 32))
									case 1:
										v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+int32(-3)))))
										v132 = v118
									case 2:
										v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110+int32(-5)))))
										v132 = v121
									case 3:
										v124 = *(*int32)(unsafe.Add(mBase, uint32(v110+int32(-9))))
										v132 = v124
									case 4:
										v127 = *(*int32)(unsafe.Add(mBase, uint32(v110+int32(-17))))
										v132 = v127
									default:
										v132 = int32(0)
									}
									v133 = F_createStringObject_1(m, v108, v132)
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return int32(0)
									} else {
										F_deleteExpiredKeyAndPropagateWithDictIndex(m, l0, v133, l4)
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											F_decrRefCount(m, v133)
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return int32(0)
											} else {
												v147 = int32(2)
												m.G0 = v12 + int32(16)
												return v147
											}
										}
									}
								}
							}
						}
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, _c_F_expireIfNeededWithDictIndex[3]))
						if v47 == int32(0) {
							v99 = F_getExpirationPolicyWithFlags(m, l3)
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								if base.Ui32(v99) < base.Ui32(int32(2)) {
									v147 = v99
									m.G0 = v12 + int32(16)
									return v147
								} else {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									if v103&int32(-8) != int32(-16) {
										F_deleteExpiredKeyAndPropagateWithDictIndex(m, l0, l1, l4)
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											v147 = int32(2)
											m.G0 = v12 + int32(16)
											return v147
										}
									} else {
										v108 = F_objectGetVal(m, l1)
										mBase = m.M
										v110 = F_objectGetVal(m, l1)
										mBase = m.M
										v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+int32(-1)))))
										switch v113 & int32(7) {
										case 0:
											v132 = int32(base.Ui32(v113) >> (uint(int32(3)) % 32))
										case 1:
											v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+int32(-3)))))
											v132 = v118
										case 2:
											v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110+int32(-5)))))
											v132 = v121
										case 3:
											v124 = *(*int32)(unsafe.Add(mBase, uint32(v110+int32(-9))))
											v132 = v124
										case 4:
											v127 = *(*int32)(unsafe.Add(mBase, uint32(v110+int32(-17))))
											v132 = v127
										default:
											v132 = int32(0)
										}
										v133 = F_createStringObject_1(m, v108, v132)
										mBase = m.M
										v134 = m.ExcPending
										if v134 != 0 {
											return int32(0)
										} else {
											F_deleteExpiredKeyAndPropagateWithDictIndex(m, l0, v133, l4)
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return int32(0)
											} else {
												F_decrRefCount(m, v133)
												mBase = m.M
												v138 = m.ExcPending
												if v138 != 0 {
													return int32(0)
												} else {
													v147 = int32(2)
													m.G0 = v12 + int32(16)
													return v147
												}
											}
										}
									}
								}
							}
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, _c_F_expireIfNeededWithDictIndex[4]))
							if v51 == int32(0) {
								v99 = F_getExpirationPolicyWithFlags(m, l3)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									if base.Ui32(v99) < base.Ui32(int32(2)) {
										v147 = v99
										m.G0 = v12 + int32(16)
										return v147
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										if v103&int32(-8) != int32(-16) {
											F_deleteExpiredKeyAndPropagateWithDictIndex(m, l0, l1, l4)
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int32(0)
											} else {
												v147 = int32(2)
												m.G0 = v12 + int32(16)
												return v147
											}
										} else {
											v108 = F_objectGetVal(m, l1)
											mBase = m.M
											v110 = F_objectGetVal(m, l1)
											mBase = m.M
											v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+int32(-1)))))
											switch v113 & int32(7) {
											case 0:
												v132 = int32(base.Ui32(v113) >> (uint(int32(3)) % 32))
											case 1:
												v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+int32(-3)))))
												v132 = v118
											case 2:
												v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110+int32(-5)))))
												v132 = v121
											case 3:
												v124 = *(*int32)(unsafe.Add(mBase, uint32(v110+int32(-9))))
												v132 = v124
											case 4:
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v110+int32(-17))))
												v132 = v127
											default:
												v132 = int32(0)
											}
											v133 = F_createStringObject_1(m, v108, v132)
											mBase = m.M
											v134 = m.ExcPending
											if v134 != 0 {
												return int32(0)
											} else {
												F_deleteExpiredKeyAndPropagateWithDictIndex(m, l0, v133, l4)
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													F_decrRefCount(m, v133)
													mBase = m.M
													v138 = m.ExcPending
													if v138 != 0 {
														return int32(0)
													} else {
														v147 = int32(2)
														m.G0 = v12 + int32(16)
														return v147
													}
												}
											}
										}
									}
								}
							} else {
								v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+207)))
								if v54&int32(32) == int32(0) {
									v99 = F_getExpirationPolicyWithFlags(m, l3)
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int32(0)
									} else {
										if base.Ui32(v99) < base.Ui32(int32(2)) {
											v147 = v99
											m.G0 = v12 + int32(16)
											return v147
										} else {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											if v103&int32(-8) != int32(-16) {
												F_deleteExpiredKeyAndPropagateWithDictIndex(m, l0, l1, l4)
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													v147 = int32(2)
													m.G0 = v12 + int32(16)
													return v147
												}
											} else {
												v108 = F_objectGetVal(m, l1)
												mBase = m.M
												v110 = F_objectGetVal(m, l1)
												mBase = m.M
												v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+int32(-1)))))
												switch v113 & int32(7) {
												case 0:
													v132 = int32(base.Ui32(v113) >> (uint(int32(3)) % 32))
												case 1:
													v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+int32(-3)))))
													v132 = v118
												case 2:
													v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110+int32(-5)))))
													v132 = v121
												case 3:
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v110+int32(-9))))
													v132 = v124
												case 4:
													v127 = *(*int32)(unsafe.Add(mBase, uint32(v110+int32(-17))))
													v132 = v127
												default:
													v132 = int32(0)
												}
												v133 = F_createStringObject_1(m, v108, v132)
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return int32(0)
												} else {
													F_deleteExpiredKeyAndPropagateWithDictIndex(m, l0, v133, l4)
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														F_decrRefCount(m, v133)
														mBase = m.M
														v138 = m.ExcPending
														if v138 != 0 {
															return int32(0)
														} else {
															v147 = int32(2)
															m.G0 = v12 + int32(16)
															return v147
														}
													}
												}
											}
										}
									}
								} else {
									v147 = v6
									m.G0 = v12 + int32(16)
									return v147
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_setExpire(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
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
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	v11 = F_objectGetVal(m, l2)
	mBase = m.M
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_setExpire[0]))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v21 = F_objectGetVal(m, l2)
	mBase = m.M
	v22 = F_kvstoreHashtableFindRef(m, v20, v19, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L10
	}
L2:
	;
	v15 = F_getKeySlot(m, v11)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v19 = int32(0)
	goto L1
L4:
	;
	return int32(0)
L5:
	;
	v19 = v15
	goto L1
L6:
	;
	F__serverAssert(m, int32(_a_F_setExpire_2), int32(_a_F_setExpire_1), int32(1927))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L51
	}
L7:
	;
	F__serverAssert(m, int32(_a_F_setExpire_3), int32(_a_F_setExpire_1), int32(1917))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L50
	}
L8:
	;
	F__serverAssert(m, int32(_a_F_setExpire_5), int32(_a_F_setExpire_1), int32(1913))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L49
	}
L9:
	;
	F__serverAssertWithInfo(m, int32(0), l2, int32(_a_F_setExpire_0), int32(_a_F_setExpire_1), int32(1903))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L48
	}
L10:
	;
	if v22 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v30&int32(1) == int32(0) {
		v41 = int64(-1)
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v42 = F_objectSetExpire(m, v26, l3)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L16
	}
L13:
	;
	goto L12
L14:
	;
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v26+(v30&int32(4)^int32(12)))))
	v41 = v40
	goto L13
L15:
	;
	if v41 == int64(-1) {
		goto L31
	} else {
		goto L32
	}
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v44&int32(15) != int32(4) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v49 = F_hashTypeHasVolatileFields(m, v42)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v49 == int32(0) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v53 = int32(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v56&int32(2) == v53 {
		v76 = v53
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_setExpire[0]))
	if v78 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L20
L22:
	;
	v70 = v42 + (v56&int32(4) ^ int32(12)) + v56<<(uint(int32(3))%32)&int32(8)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v76 = v70 + v71 + int32(1)
	goto L21
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84+v82<<(uint(int32(2))%32))))
	goto L27
L24:
	;
	v80 = F_getKeySlot(m, v76)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L26
	}
L25:
	;
	v82 = int32(0)
	goto L23
L26:
	;
	v82 = v80
	goto L23
L27:
	;
	v89 = F_hashtableReplaceReallocatedEntry(m, v88, v26, v42)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	if v89 == int32(0) {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	goto L15
L30:
	;
	if l0 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L31:
	;
	if v42 == v26 {
		v114 = v26
		goto L38
	} else {
		goto L39
	}
L32:
	;
	if v42 != v26 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_setExpire[3]))
	if v98 == int32(0) {
		v120 = v26
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v102 = F_kvstoreHashtableAdd(m, v101, v19, v42)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	if v102 == int32(0) {
		v120 = v26
		goto L30
	} else {
		goto L36
	}
L36:
	;
	F__serverAssert(m, int32(_a_F_setExpire_4), int32(_a_F_setExpire_1), int32(1919))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v116 = F_kvstoreHashtableAdd(m, v115, v19, v42)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v42
	v114 = v42
	goto L38
L40:
	;
	if v116 == int32(0) {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v120 = v114
	goto L30
L42:
	;
	return v120
L43:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_setExpire[1]))
	if v124 == int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_setExpire[2]))
	if v128 != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v129&int32(1) != 0 {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	F_rememberReplicaKeyWithExpire(m, l1, l2)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	goto L42
L48:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
