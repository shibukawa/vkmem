package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F___quicklistDelNode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v7) < base.Ui32(int32(268435456)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v205 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L2:
	;
	v17 = int32(0)
	goto L4
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v29
	if v29 != 0 {
		goto L1
	} else {
		goto L8
	}
L4:
	;
	v23 = l0 + int32(20) + v17<<(uint(int32(3))%32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v24 == l1 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v27 = v17 + int32(1)
	if v27 == int32(base.Ui32(v7)>>(uint(int32(28))%32)) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v17 = v27
	goto L4
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	F_valkey_free(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v35 = int32(-268435456)
	v38 = v34&v35 + v35
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v38 | v34&int32(268435455)
	v44 = v23 + int32(8)
	v49 = (int32(base.Ui32(v38)>>(uint(int32(28))%32)) - v17) << (uint(int32(3)) % 32)
	if v23 == v44 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L1
L12:
	;
	goto L11
L13:
	;
	v53 = v49 + v23
	if base.Ui32(int32(0)-v49<<(uint(int32(1))%32)) < base.Ui32(v44-v53) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v63 = (v44 ^ v23) & int32(3)
	if base.Ui32(v44) <= base.Ui32(v23) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v60 = F___memcpy(m, v23, v44, v49)
	mBase = m.M
	goto L11
L16:
	;
	if v169 == int32(0) {
		goto L12
	} else {
		goto L48
	}
L17:
	;
	if base.Ui32(v147) <= base.Ui32(int32(3)) {
		v168 = v146
		v169 = v147
		v170 = v148
		goto L16
	} else {
		goto L44
	}
L18:
	;
	if v63 != 0 {
		v129 = v49
		goto L28
	} else {
		goto L29
	}
L19:
	;
	if v63 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v23&int32(3) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v168 = v44
	v169 = v49
	v170 = v23
	goto L16
L22:
	;
	v70 = v44
	v71 = v49
	v72 = v23
	goto L24
L23:
	;
	v146 = v44
	v147 = v49
	v148 = v23
	goto L17
L24:
	;
	if v71 == int32(0) {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v76)
	v78 = int32(1)
	v79 = v70 + v78
	v81 = v71 + int32(-1)
	v83 = v72 + v78
	if v83&int32(3) == int32(0) {
		v146 = v79
		v147 = v81
		v148 = v83
		goto L17
	} else {
		goto L27
	}
L27:
	;
	v70 = v79
	v71 = v81
	v72 = v83
	goto L24
L28:
	;
	if v129 == int32(0) {
		goto L12
	} else {
		goto L40
	}
L29:
	;
	if v53&int32(3) == int32(0) {
		v109 = v49
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if base.Ui32(v109) <= base.Ui32(int32(3)) {
		v129 = v109
		goto L28
	} else {
		goto L36
	}
L31:
	;
	v94 = v49
	goto L32
L32:
	;
	if v94 == int32(0) {
		goto L12
	} else {
		goto L34
	}
L33:
	;
	v109 = v100
	goto L30
L34:
	;
	v100 = v94 + int32(-1)
	v101 = v23 + v100
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v100))))
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v103)
	if v101&int32(3) != 0 {
		v94 = v100
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v116 = v109
	goto L37
L37:
	;
	v120 = v116 + int32(-4)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v44+v120)))
	*(*int32)(unsafe.Add(mBase, uint32(v23+v120))) = v123
	if base.Ui32(int32(3)) < base.Ui32(v120) {
		v116 = v120
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v129 = v120
	goto L28
L39:
	;
	goto L38
L40:
	;
	v136 = v129
	goto L41
L41:
	;
	v140 = v136 + int32(-1)
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v23+v140))) = uint8(v143)
	if v140 != 0 {
		v136 = v140
		goto L41
	} else {
		goto L43
	}
L43:
	;
	goto L12
L44:
	;
	v153 = v146
	v154 = v147
	v155 = v148
	goto L45
L45:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v157
	v159 = int32(4)
	v160 = v153 + v159
	v162 = v155 + v159
	v164 = v154 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v164) {
		v153 = v160
		v154 = v164
		v155 = v162
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v168 = v160
	v169 = v164
	v170 = v162
	goto L16
L47:
	;
	goto L46
L48:
	;
	v175 = v168
	v176 = v169
	v177 = v170
	goto L49
L49:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v179)
	v181 = int32(1)
	v186 = v176 + int32(-1)
	if v186 != 0 {
		v175 = v175 + v181
		v176 = v186
		v177 = v177 + v181
		goto L49
	} else {
		goto L51
	}
L50:
	;
	goto L12
L51:
	;
	goto L50
L52:
	;
	if v204 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v204
	goto L52
L54:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l1 != v212 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+4)) = v205
	goto L54
L56:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 != v215 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v204
	goto L56
L58:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v219 + int32(-1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v223 - v224
	F___quicklistCompress(m, l0, int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L9
	} else {
		goto L60
	}
L59:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v217
	goto L58
L60:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_valkey_free(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	F_valkey_free(m, l1)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	return
}
func F___quicklistInsertNode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	if l3 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = l1
		if l1 == int32(0) {
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v23
			if v23 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = l2
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = l2
		}
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v30 != l1 {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = l2
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = l1
		if l1 == int32(0) {
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v10
			if v10 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l2
		}
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v17 != l1 {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
		}
	}
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v34 != 0 {
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = l2
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v34 + int32(1)
	if l1 == int32(0) {
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
		if v56&int32(1048576) == int32(0) {
			F___quicklistCompress(m, l0, l2)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return
			} else {
				return
			}
		} else {
			if v56&int32(196608) != int32(65536) {
				return
			} else {
				v65 = F___quicklistCompressNode(m, l2)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		if v42&int32(1048576) == int32(0) {
			F___quicklistCompress(m, l0, l1)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				if v56&int32(1048576) == int32(0) {
					F___quicklistCompress(m, l0, l2)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						return
					}
				} else {
					if v56&int32(196608) != int32(65536) {
						return
					} else {
						v65 = F___quicklistCompressNode(m, l2)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			if v42&int32(196608) != int32(65536) {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				if v56&int32(1048576) == int32(0) {
					F___quicklistCompress(m, l0, l2)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						return
					}
				} else {
					if v56&int32(196608) != int32(65536) {
						return
					} else {
						v65 = F___quicklistCompressNode(m, l2)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v51 = F___quicklistCompressNode(m, l1)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
					if v56&int32(1048576) == int32(0) {
						F___quicklistCompress(m, l0, l2)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							return
						}
					} else {
						if v56&int32(196608) != int32(65536) {
							return
						} else {
							v65 = F___quicklistCompressNode(m, l2)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
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
func F__quicklistMergeNodes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
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
	v3 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v14 == v3 {
		v18 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v21 == int32(0) {
		v25 = v3
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v18 = v17
	goto L1
L3:
	;
	v27 = v11 << (uint(int32(18)) % 32) >> (uint(int32(18)) % 32)
	if v14 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v25 = v24
	goto L3
L5:
	;
	if v21 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L6:
	;
	if v18 == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v32&int32(786432) == int32(262144) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v37&int32(786432) == int32(262144) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v46 = v42 + v43 + int32(-7)
	if int32(-1) < v27 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v73 = F__quicklistListpackMerge(m, l0, v18, v14)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	v61 = int32(65535)
	v66 = int32(1)
	if base.Ui32(v66) < base.Ui32(v27) {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v49 = int32(-5)
	if base.Ui32(v49) < base.Ui32(v27) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v52 = v27
	goto L15
L14:
	;
	v52 = v49
	goto L15
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32((v52^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[524])))
	if base.Ui32(v46) <= base.Ui32(v59) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L5
L17:
	;
	v69 = v27
	goto L19
L18:
	;
	v69 = v66
	goto L19
L19:
	;
	if base.Ui32(v69) < base.Ui32(v37&v61+v32&v61) {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32(int32(8192)) < base.Ui32(v46) {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	goto L10
L22:
	;
	return int32(0)
L23:
	;
	goto L5
L24:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v130 == int32(0) {
		v176 = l1
		goto L42
	} else {
		goto L43
	}
L25:
	;
	if v25 == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v84&int32(786432) == int32(262144) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v89&int32(786432) == int32(262144) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v98 = v94 + v95 + int32(-7)
	if int32(-1) < v27 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v125 = F__quicklistListpackMerge(m, l0, v21, v25)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L22
	} else {
		goto L41
	}
L30:
	;
	v113 = int32(65535)
	v118 = int32(1)
	if base.Ui32(v118) < base.Ui32(v27) {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v101 = int32(-5)
	if base.Ui32(v101) < base.Ui32(v27) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v104 = v27
	goto L34
L33:
	;
	v104 = v101
	goto L34
L34:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32((v104^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[524])))
	if base.Ui32(v98) <= base.Ui32(v111) {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	goto L24
L36:
	;
	v121 = v27
	goto L38
L37:
	;
	v121 = v118
	goto L38
L38:
	;
	if base.Ui32(v121) < base.Ui32(v89&v113+v84&v113) {
		goto L24
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(int32(8192)) < base.Ui32(v98) {
		goto L24
	} else {
		goto L40
	}
L40:
	;
	goto L29
L41:
	;
	goto L24
L42:
	;
	if v176 == int32(0) {
		v228 = v176
		goto L59
	} else {
		goto L60
	}
L43:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v133&int32(786432) == int32(262144) {
		v176 = l1
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	if v138&int32(786432) == int32(262144) {
		v176 = l1
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v147 = v143 + v144 + int32(-7)
	if int32(-1) < v27 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v174 = F__quicklistListpackMerge(m, l0, v130, l1)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L22
	} else {
		goto L58
	}
L47:
	;
	v162 = int32(65535)
	v167 = int32(1)
	if base.Ui32(v167) < base.Ui32(v27) {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v150 = int32(-5)
	if base.Ui32(v150) < base.Ui32(v27) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v153 = v27
	goto L51
L50:
	;
	v153 = v150
	goto L51
L51:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32((v153^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[524])))
	if base.Ui32(v147) <= base.Ui32(v160) {
		goto L46
	} else {
		goto L52
	}
L52:
	;
	v176 = l1
	goto L42
L53:
	;
	v170 = v27
	goto L55
L54:
	;
	v170 = v167
	goto L55
L55:
	;
	if base.Ui32(v170) < base.Ui32(v138&v162+v133&v162) {
		v176 = l1
		goto L42
	} else {
		goto L56
	}
L56:
	;
	if base.Ui32(int32(8192)) < base.Ui32(v147) {
		v176 = l1
		goto L42
	} else {
		goto L57
	}
L57:
	;
	goto L46
L58:
	;
	v176 = v174
	goto L42
L59:
	;
	return v228
L60:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v182 == int32(0) {
		v228 = v176
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v176)+16))
	if v185&int32(786432) == int32(262144) {
		v228 = v176
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v182)+16))
	if v190&int32(786432) == int32(262144) {
		v228 = v176
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	v199 = v195 + v196 + int32(-7)
	if int32(-1) < v27 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v226 = F__quicklistListpackMerge(m, l0, v176, v182)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L22
	} else {
		goto L76
	}
L65:
	;
	v214 = int32(65535)
	v219 = int32(1)
	if base.Ui32(v219) < base.Ui32(v27) {
		goto L71
	} else {
		goto L72
	}
L66:
	;
	v202 = int32(-5)
	if base.Ui32(v202) < base.Ui32(v27) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v205 = v27
	goto L69
L68:
	;
	v205 = v202
	goto L69
L69:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32((v205^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[524])))
	if base.Ui32(v199) <= base.Ui32(v212) {
		goto L64
	} else {
		goto L70
	}
L70:
	;
	v228 = v176
	goto L59
L71:
	;
	v222 = v27
	goto L73
L72:
	;
	v222 = v219
	goto L73
L73:
	;
	if base.Ui32(v222) < base.Ui32(v190&v214+v185&v214) {
		v228 = v176
		goto L59
	} else {
		goto L74
	}
L74:
	;
	if base.Ui32(int32(8192)) < base.Ui32(v199) {
		v228 = v176
		goto L59
	} else {
		goto L75
	}
L75:
	;
	goto L64
L76:
	;
	v228 = v226
	goto L59
}
func F_createQuicklistObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_quicklistNew(m, l0, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
		v14 = int32(12)
		v17 = F_zmalloc_usable(m, v14, v6+v14)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v8
			*(*int64)(unsafe.Add(mBase, uint32(v17))) = int64(34359738369)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v22&int32(-241) | int32(144)
			m.G0 = v6 + int32(16)
			return v17
		}
	}
}
func F_quicklistAppendListpack(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = int64(0)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
		v13 = F_lpLength(m, l1)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v10&int32(-6291456) | v13&int32(65535) | int32(589824)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v21
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			F___quicklistInsertNode(m, l0, v23, v5, int32(1))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+16)))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27 + v28
				return
			}
		}
	}
}
func F_quicklistAppendPlainNode(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v6 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(0)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v12&int32(-6291456) | int32(327681)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F___quicklistInsertNode(m, l0, v18, v6, int32(1))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+16)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v22 + v23
			return
		}
	}
}
func F_quicklistInsertAfter(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v7 int32
	_ = v7
	F__quicklistInsert(m, l0, l1, l2, l3, int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_quicklistNew(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v5 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = v9
		*(*int64)(unsafe.Add(mBase, uint32(v5+int32(8)))) = v9
		v16 = int32(0)
		if v16 < l1 {
			v19 = l1
		} else {
			v19 = v16
		}
		if int32(16383) < l1 {
			v26 = int32(268419072)
		} else {
			v26 = v19 << (uint(int32(14)) % 32) & int32(268419072)
		}
		v28 = int32(-5)
		if v28 < l0 {
			v31 = l0
		} else {
			v31 = v28
		}
		if int32(8191) < l0 {
			v36 = int32(8191)
		} else {
			v36 = v31 & int32(16383)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v26 | v36
		return v5
	}
}
func F_quicklistNodeLimit(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	v4 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	if l0 < int32(0) {
		v15 = int32(-5)
		if base.Ui32(v15) < base.Ui32(l0) {
			v18 = l0
		} else {
			v18 = v15
		}
		v25 = *(*int32)(unsafe.Add(mBase, uint32((v18^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[524])))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25
		return
	} else {
		v10 = int32(1)
		if base.Ui32(v10) < base.Ui32(l0) {
			v13 = l0
		} else {
			v13 = v10
		}
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v13
		return
	}
}
func F_quicklistPush(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v6 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v14 == int32(0) {
			switch l3 + int32(1) {
			case 0:
				v26 = F_quicklistPushTail(m, l0, l1, l2)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					return
				}
			case 1:
				v24 = F_quicklistPushHead(m, l0, l1, l2)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					return
				}
			default:
				return
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			if v17&int32(196608) == int32(131072) {
				F__serverAssert(m, int32(_a1789), int32(_a1790), int32(1570))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				switch l3 + int32(1) {
				case 0:
					v26 = F_quicklistPushTail(m, l0, l1, l2)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						return
					}
				case 1:
					v24 = F_quicklistPushHead(m, l0, l1, l2)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						return
					}
				default:
					return
				}
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
		if v9&int32(196608) == int32(131072) {
			F__serverAssert(m, int32(_a1791), int32(_a1790), int32(1569))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v14 == int32(0) {
				switch l3 + int32(1) {
				case 0:
					v26 = F_quicklistPushTail(m, l0, l1, l2)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						return
					}
				case 1:
					v24 = F_quicklistPushHead(m, l0, l1, l2)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						return
					}
				default:
					return
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
				if v17&int32(196608) == int32(131072) {
					F__serverAssert(m, int32(_a1789), int32(_a1790), int32(1570))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					switch l3 + int32(1) {
					case 0:
						v26 = F_quicklistPushTail(m, l0, l1, l2)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							return
						}
					case 1:
						v24 = F_quicklistPushHead(m, l0, l1, l2)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							return
						}
					default:
						return
					}
				}
			}
		}
	}
}
func F_quicklistRepr(m *base.Module, l0 int32, l1 int32) {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = v12
	v17 = F_iprintf(m, int32(_a1792), v10+int32(128))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v19
	v24 = F_iprintf(m, int32(_a1793), v10+int32(112))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v27 = int32(18)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v26 << (uint(v27) % 32) >> (uint(v27) % 32)
	v35 = F_iprintf(m, int32(_a1794), v10+int32(96))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = int32(base.Ui32(v37)>>(uint(int32(14))%32)) & int32(16383)
	v46 = F_iprintf(m, int32(_a1795), v10+int32(80))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = int32(base.Ui32(v48) >> (uint(int32(28)) % 32))
	v55 = F_iprintf(m, int32(_a1796), v10+int32(64))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v57 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	m.G0 = v10 + int32(144)
	return
L8:
	;
	v63 = v57
	v66 = int32(0)
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v66
	v74 = F_iprintf(m, int32(_a1797), v10+int32(48))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L7
L11:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v84 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(36)))) = int32(base.Ui32(v81)>>(uint(int32(21))%32)) & v84
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(32)))) = int32(base.Ui32(v81)>>(uint(int32(20))%32)) & v84
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v81 & int32(65535)
	if v81&int32(196608) == int32(65536) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v102 = int32(_a1798)
	goto L14
L13:
	;
	v102 = int32(_a1799)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v102
	if v81&int32(786432) == int32(262144) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v110 = int32(_a1800)
	goto L17
L16:
	;
	v110 = int32(_a1801)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v110
	v115 = F_iprintf(m, int32(_a1802), v10+int32(16))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if l1 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v184 != 0 {
		v63 = v184
		v66 = v66 + int32(1)
		goto L9
	} else {
		goto L38
	}
L20:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	if v119&int32(196608) != int32(131072) {
		v149 = v119
		goto L21
	} else {
		goto L22
	}
L21:
	;
	switch int32(base.Ui32(v149)>>(uint(int32(18))%32))&int32(3) + int32(-1) {
	case 0:
		goto L29
	case 1:
		goto L30
	default:
		goto L28
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v119 & int32(-3211265)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v128 = F_valkey_malloc(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v135 = F_lzf_decompress(m, v130+int32(4), v133, v128, v134)
	mBase = m.M
	if v135 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_valkey_free(m, v130)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	F_valkey_free(m, v128)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v149 = v138
	goto L21
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v128
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v146 = v142&int32(-196609) | int32(65536)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v146
	v149 = v146
	goto L21
L28:
	;
	v171 = F_puts(m, int32(_a1803))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L35
	}
L29:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v165
	v168 = F_iprintf(m, int32(_a1804), v10)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L34
	}
L30:
	;
	v157 = F_puts(m, int32(_a1805))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	F_lpRepr(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v163 = F_puts(m, int32(_a1803))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L28
L34:
	;
	goto L28
L35:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	if v173&int32(1245184) != int32(1114112) {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	v178 = F___quicklistCompressNode(m, v63)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	goto L19
L38:
	;
	goto L10
}
