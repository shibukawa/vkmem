package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_unblockClientFromModule(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int64
	_ = v63
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v89 int64
	_ = v89
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v104 int64
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if v12 == int32(0) {
		v123 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
		if v123 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
			m.G0 = v8 + int32(80)
			return
		} else {
			v126 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
			if v126 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
				m.G0 = v8 + int32(80)
				return
			} else {
				v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+48))
				v131 = *(*int32)(unsafe.Add(mBase, uint32(v128)+44))
				if v131 != 0 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v128)+28)) = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v128)+48)) = int32(1)
				v137 = *(*int32)(unsafe.Add(mBase, _c_F_unblockClientFromModule[0]))
				v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
				if v138 != 0 {
					v146 = v137
				} else {
					v140 = *(*int32)(unsafe.Add(mBase, _c_F_unblockClientFromModule[1]))
					v143 = F_write(m, v140, int32(_a_F_unblockClientFromModule_0), int32(1))
					mBase = m.M
					v145 = *(*int32)(unsafe.Add(mBase, _c_F_unblockClientFromModule[0]))
					v146 = v145
				}
				v147 = F_listAddNodeTail(m, v146, v128)
				mBase = m.M
				v148 = m.ExcPending
				if v148 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
					m.G0 = v8 + int32(80)
					return
				}
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v18 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(72)))) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(64)))) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(56)))) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(48)))) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(40)))) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(32)))) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v8+int32(24)))) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(561)
		v50 = *(*int32)(unsafe.Add(mBase, _c_F_unblockClientFromModule[2]))
		v51 = int32(0)
		v52 = *(*int32)(unsafe.Add(mBase, _c_F_unblockClientFromModule[3]))
		v53 = m.T0[v52].(func(*base.Module) int64)(m)
		mBase = m.M
		if v50 == v51 {
			v63 = *(*int64)(unsafe.Add(mBase, _c_F_unblockClientFromModule[4]))
			v67 = v63*int64(1000) + v53
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, _c_F_unblockClientFromModule[5]))
			v59 = base.I32_div_s(int32(1000000), v58)
			v67 = v53 + base.I64_extend_i32_s(v59)
		}
		*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = v67
		v71 = int32(0)
		v75 = *(*int32)(unsafe.Add(mBase, _c_F_unblockClientFromModule[6]))
		*(*int32)(unsafe.Add(mBase, _c_F_unblockClientFromModule[6])) = v75 + int32(1)
		if v75 != 0 {
		} else {
			v83 = F_ustime(m)
			mBase = m.M
			v85 = int32(0)
			*(*int64)(unsafe.Add(mBase, _c_F_unblockClientFromModule[7])) = v83
			v89 = base.I64_div_s(v83, int64(1000))
			*(*int64)(unsafe.Add(mBase, _c_F_unblockClientFromModule[8])) = v89
			v93 = base.I64_div_s(v83, int64(1000000))
			*(*int64)(unsafe.Add(mBase, _c_F_unblockClientFromModule[9])) = v93
			v96 = *(*int32)(unsafe.Add(mBase, _c_F_unblockClientFromModule[10]))
			F_lrulfu_updateClockAndPolicy(m, v89, int32(base.Ui32(v96&int32(2))>>(uint(int32(1))%32)))
			mBase = m.M
			v104 = *(*int64)(unsafe.Add(mBase, _c_F_unblockClientFromModule[8]))
			*(*int64)(unsafe.Add(mBase, _c_F_unblockClientFromModule[11])) = v104
		}
		v108 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v108
		v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v110
		v114 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
		m.T0[v114].(func(*base.Module, int32, int32))(m, v8+int32(8), v11)
		mBase = m.M
		v116 = m.ExcPending
		if v116 != 0 {
			return
		} else {
			F_moduleFreeContext(m, v8+int32(8))
			mBase = m.M
			v120 = m.ExcPending
			if v120 != 0 {
				return
			} else {
				v123 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
				if v123 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
					m.G0 = v8 + int32(80)
					return
				} else {
					v126 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
					if v126 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
						m.G0 = v8 + int32(80)
						return
					} else {
						v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+48))
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v128)+44))
						if v131 != 0 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v128)+28)) = int32(0)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v128)+48)) = int32(1)
						v137 = *(*int32)(unsafe.Add(mBase, _c_F_unblockClientFromModule[0]))
						v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
						if v138 != 0 {
							v146 = v137
						} else {
							v140 = *(*int32)(unsafe.Add(mBase, _c_F_unblockClientFromModule[1]))
							v143 = F_write(m, v140, int32(_a_F_unblockClientFromModule_0), int32(1))
							mBase = m.M
							v145 = *(*int32)(unsafe.Add(mBase, _c_F_unblockClientFromModule[0]))
							v146 = v145
						}
						v147 = F_listAddNodeTail(m, v146, v128)
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
							m.G0 = v8 + int32(80)
							return
						}
					}
				}
			}
		}
	}
}
func F_unblockClientOnTimeout(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	if v4 != int32(3) {
		F_replyToBlockedClientTimedOut(m, l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
			if v14&int32(2) == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v14 & int32(-3)
			}
			F_unblockClient(m, l0, int32(1))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
		if v9 == int32(1) {
			return
		} else {
			F_replyToBlockedClientTimedOut(m, l0)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
				if v14&int32(2) == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v14 & int32(-3)
				}
				F_unblockClient(m, l0, int32(1))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_unblockClientWaitingData(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
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
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v176 int64
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	if v6 == int32(0)-v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = F_dictGetIterator(m, v5)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_dictReleaseIterator(m, v11)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L63
	}
L4:
	;
	return
L5:
	;
	v20 = v11 + int32(20)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v116 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L7:
	;
	v27 = v20
	v28 = v24
	goto L10
L8:
	;
	v24 = int32(1)
	goto L7
L9:
	;
	v24 = int32(0)
	goto L7
L10:
	;
	switch v28 {
	case 0:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v28 = int32(0)
	goto L10
L13:
	;
	goto L6
L14:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v108
	if v108 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v32 != int32(-1) {
		v71 = v32
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v72 = int32(1)
	v73 = v71 + v72
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v73
	v75 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v79+int32(26)))))
	if v83 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v36 != 0 {
		v71 = int32(-1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v38 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	if v65 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v45 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v37)+16)))
	v46 = int64(*(*int8)(unsafe.Add(mBase, uint32(v37)+27)))
	v47 = int64(*(*int32)(unsafe.Add(mBase, uint32(v37)+8)))
	v48 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v37)+12)))
	v49 = int64(*(*int8)(unsafe.Add(mBase, uint32(v37)+26)))
	v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v37)+4)))
	v51 = F_wangHash64(m, v50)
	mBase = m.M
	v53 = F_wangHash64(m, v49+v51)
	mBase = m.M
	v55 = F_wangHash64(m, v48+v53)
	mBase = m.M
	v57 = F_wangHash64(m, v47+v55)
	mBase = m.M
	v59 = F_wangHash64(m, v46+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v45+v59)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v64 = v63
	goto L19
L21:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+24)))
	v43 = v41 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+24)) = uint16(v43)
	v64 = v37
	goto L19
L22:
	;
	v71 = v65 + int32(-1)
	goto L16
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v71 = v68
	goto L16
L24:
	;
	v98 = int32(2)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v78+v96<<(uint(v98)%32)+int32(4))))
	v27 = v103 + v97<<(uint(v98)%32)
	v28 = int32(1)
	goto L10
L25:
	;
	v87 = v75
	goto L27
L26:
	;
	v87 = v72 << (uint(v83) % 32)
	goto L27
L27:
	;
	if v73 < v87 {
		v96 = v79
		v97 = v73
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v79 != 0 {
		v116 = v75
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	if v89 == int32(-1) {
		v116 = v75
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(4294967296)
	v96 = int32(1)
	v97 = int32(0)
	goto L24
L31:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v112
	v116 = v108
	goto L13
L32:
	;
	v123 = v116
	goto L33
L33:
	;
	F_releaseBlockedEntry(m, l0, v123, int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L35
	}
L34:
	;
	goto L3
L35:
	;
	v135 = v11 + int32(20)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v136 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	if v231 != 0 {
		v123 = v231
		goto L33
	} else {
		goto L62
	}
L37:
	;
	v142 = v135
	v143 = v139
	goto L40
L38:
	;
	v139 = int32(1)
	goto L37
L39:
	;
	v139 = int32(0)
	goto L37
L40:
	;
	switch v143 {
	case 0:
		goto L45
	default:
		goto L44
	}
L42:
	;
	v143 = int32(0)
	goto L40
L43:
	;
	goto L36
L44:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v223
	if v223 == int32(0) {
		goto L42
	} else {
		goto L61
	}
L45:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v147 != int32(-1) {
		v186 = v147
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v187 = int32(1)
	v188 = v186 + v187
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v188
	v190 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v194+int32(26)))))
	if v198 == int32(255) {
		goto L55
	} else {
		goto L56
	}
L47:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v151 != 0 {
		v186 = int32(-1)
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v153 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	if v180 != int32(-1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v160 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v152)+16)))
	v161 = int64(*(*int8)(unsafe.Add(mBase, uint32(v152)+27)))
	v162 = int64(*(*int32)(unsafe.Add(mBase, uint32(v152)+8)))
	v163 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v152)+12)))
	v164 = int64(*(*int8)(unsafe.Add(mBase, uint32(v152)+26)))
	v165 = int64(*(*int32)(unsafe.Add(mBase, uint32(v152)+4)))
	v166 = F_wangHash64(m, v165)
	mBase = m.M
	v168 = F_wangHash64(m, v164+v166)
	mBase = m.M
	v170 = F_wangHash64(m, v163+v168)
	mBase = m.M
	v172 = F_wangHash64(m, v162+v170)
	mBase = m.M
	v174 = F_wangHash64(m, v161+v172)
	mBase = m.M
	v176 = F_wangHash64(m, v160+v174)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v179 = v178
	goto L49
L51:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+24)))
	v158 = v156 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v152)+24)) = uint16(v158)
	v179 = v152
	goto L49
L52:
	;
	v186 = v180 + int32(-1)
	goto L46
L53:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v186 = v183
	goto L46
L54:
	;
	v213 = int32(2)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v193+v211<<(uint(v213)%32)+int32(4))))
	v142 = v218 + v212<<(uint(v213)%32)
	v143 = int32(1)
	goto L40
L55:
	;
	v202 = v190
	goto L57
L56:
	;
	v202 = v187 << (uint(v198) % 32)
	goto L57
L57:
	;
	if v188 < v202 {
		v211 = v194
		v212 = v188
		goto L54
	} else {
		goto L58
	}
L58:
	;
	if v194 != 0 {
		v231 = v190
		goto L43
	} else {
		goto L59
	}
L59:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v193)+20))
	if v204 == int32(-1) {
		v231 = v190
		goto L43
	} else {
		goto L60
	}
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(4294967296)
	v211 = int32(1)
	v212 = int32(0)
	goto L54
L61:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v227
	v231 = v223
	goto L43
L62:
	;
	goto L34
L63:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+24))
	F_dictEmpty(m, v241, int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	goto L1
}
func F_unlink(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	v4 = m.Env.X__syscall_unlinkat(m, int32(-100), l0, int32(0))
	mBase = m.M
	if base.Ui32(v4) < base.Ui32(int32(-4095)) {
		v12 = v4
	} else {
		v7 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0) - v4
		v12 = int32(-1)
	}
	return v12
}
func F_unpauseActions(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	v4 = l0 << (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+uint32(_c_F_unpauseActions[0]))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4)+uint32(_c_F_unpauseActions[1]))) = int64(0)
	F_updatePausedActions(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		return
	}
}
func F_unsubscribeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v38 int32
	_ = v38
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int64
	_ = v91
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	v8 = m.G0
	v10 = v8 - int32(64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v38 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v14 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v14
	v18 = F_hashtableCreate(m, int32(_a_F_unsubscribeCommand_0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v18
	v23 = F_hashtableCreate(m, int32(_a_F_unsubscribeCommand_0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v23
	v28 = F_hashtableCreate(m, int32(_a_F_unsubscribeCommand_0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v28
	goto L1
L8:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	goto L18
L9:
	;
	v85 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_unsubscribeCommand[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v86
	v91 = *(*int64)(unsafe.Add(mBase, _c_F_unsubscribeCommand[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(16)))) = v91
	v96 = *(*int64)(unsafe.Add(mBase, _c_F_unsubscribeCommand[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(8)))) = v96
	v99 = *(*int64)(unsafe.Add(mBase, _c_F_unsubscribeCommand[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v99
	v102 = F_pubsubUnsubscribeAllChannelsInternal(m, l0, int32(1), v10)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L3
	} else {
		goto L16
	}
L10:
	;
	if v38 < int32(2) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v52 = int32(1)
	goto L12
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v52<<(uint(int32(2))%32))))
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_unsubscribeCommand[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(56)))) = v63
	v66 = *(*int64)(unsafe.Add(mBase, _c_F_unsubscribeCommand[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(48)))) = v66
	v69 = *(*int64)(unsafe.Add(mBase, _c_F_unsubscribeCommand[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v69
	v72 = *(*int64)(unsafe.Add(mBase, _c_F_unsubscribeCommand[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v72
	v77 = F_pubsubUnsubscribeChannel(m, l0, v61, int32(1), v10+int32(32))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v80 = v52 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v80 < v81 {
		v52 = v80
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L8
L16:
	;
	goto L8
L17:
	;
	m.G0 = v10 + int32(64)
	return
L18:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+20))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	goto L19
L19:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	goto L20
L20:
	;
	if v113+v114+(v118+v119) != int32(0)-(v125+v126) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v130&int32(262144) == int32(0) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v130 & int32(-262145)
	v138 = int32(_a_F_unsubscribeCommand_1)
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_unsubscribeCommand[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_unsubscribeCommand[4])) = v140 + int32(-1)
	goto L17
}
func F_updateExtendedRedisCompat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_updateExtendedRedisCompat[0]))
	v7 = v5 << (uint(int32(2)) % 32)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F_updateExtendedRedisCompat[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_updateExtendedRedisCompat[2])) = v10
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F_updateExtendedRedisCompat[3])))
	*(*int32)(unsafe.Add(mBase, _c_F_updateExtendedRedisCompat[4])) = v15
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F_updateExtendedRedisCompat[5])))
	*(*int32)(unsafe.Add(mBase, _c_F_updateExtendedRedisCompat[6])) = v20
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F_updateExtendedRedisCompat[7])))
	*(*int32)(unsafe.Add(mBase, _c_F_updateExtendedRedisCompat[8])) = v25
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F_updateExtendedRedisCompat[9])))
	*(*int32)(unsafe.Add(mBase, _c_F_updateExtendedRedisCompat[10])) = v30
	return int32(1)
}
func F_updateLoadingFileName(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_updateLoadingFileName[0])) = l0
	return
}
func F_updateLocaleCollate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_updateLocaleCollate[0]))
	v7 = F_setlocale(m, int32(3), v6)
	mBase = m.M
	if v7 != 0 {
		v11 = int32(1)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a_F_updateLocaleCollate_0)
		v11 = int32(0)
	}
	return v11
}
func F_updateMaxclients(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_updateMaxclients[0]))
	F_adjustOpenFilesLimit(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_updateMaxclients[0]))
		if v10 == v16 {
			v25 = int32(1)
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_updateMaxclients[1]))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_updateMaxclients[0]))
			v32 = v30 + int32(128)
			if base.Ui32(v32) <= base.Ui32(v28) {
				v45 = v25
				m.G0 = v7 + int32(16)
				return v45
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_updateMaxclients[1]))
				v36 = F_aeResizeSetSize(m, v35, v32)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					if v36 != int32(-1) {
						v45 = v25
					} else {
						v41 = int32(_a_F_updateMaxclients_0)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v41
						v45 = int32(0)
					}
					m.G0 = v7 + int32(16)
					return v45
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v16
			v19 = int32(_a_F_updateMaxclients_1)
			v23 = F_snprintf(m, v19, int32(128), int32(_a_F_updateMaxclients_2), v7)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v41 = v19
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v41
				v45 = int32(0)
				m.G0 = v7 + int32(16)
				return v45
			}
		}
	}
}
func F_updateMaxmemory(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v119 int64
	_ = v119
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int64)(unsafe.Add(mBase, _c_F_updateMaxmemory[0]))
	if v11 == int64(0) {
		m.G0 = v8 + int32(16)
		return int32(1)
	} else {
		v14 = int32(0)
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_updateMaxmemory[1]))
		if v23 < int32(261) {
			if v23 < int32(1) {
				v100 = v14
			} else {
				v31 = v14
				v32 = v23
				v34 = v32 & int32(3)
				if base.Ui32(int32(4)) <= base.Ui32(v32) {
					v41 = int32(0)
					v43 = v31
					v44 = v41
					v48 = v41
					for {
						v51 = v44 << (uint(int32(2)) % 32)
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+uint32(_c_F_updateMaxmemory[2])))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+uint32(_c_F_updateMaxmemory[3])))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v51)+uint32(_c_F_updateMaxmemory[4])))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v51)+uint32(_c_F_updateMaxmemory[5])))
						v67 = v54 + (v57 + (v60 + (v63 + v43)))
						v68 = int32(4)
						v69 = v44 + v68
						v71 = v48 + v68
						if v71 != v32&int32(2147483644) {
							v43 = v67
							v44 = v69
							v48 = v71
							continue
						} else {
							break
						}
						break
					}
					v73 = v67
					v74 = v69
				} else {
					v73 = v31
					v74 = int32(0)
				}
				if v34 == int32(0) {
					v100 = v73
				} else {
					v82 = v73
					v83 = v74
					v85 = int32(0)
					for {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v83<<(uint(int32(2))%32))+uint32(_c_F_updateMaxmemory[5])))
						v94 = v93 + v82
						v95 = int32(1)
						v98 = v85 + v95
						if v98 != v34 {
							v82 = v94
							v83 = v83 + v95
							v85 = v98
							continue
						} else {
							break
						}
						break
					}
					v100 = v94
				}
			}
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_updateMaxmemory[6]))
			v31 = v27
			v32 = int32(260)
			v34 = v32 & int32(3)
			if base.Ui32(int32(4)) <= base.Ui32(v32) {
				v41 = int32(0)
				v43 = v31
				v44 = v41
				v48 = v41
				for {
					v51 = v44 << (uint(int32(2)) % 32)
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+uint32(_c_F_updateMaxmemory[2])))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+uint32(_c_F_updateMaxmemory[3])))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v51)+uint32(_c_F_updateMaxmemory[4])))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v51)+uint32(_c_F_updateMaxmemory[5])))
					v67 = v54 + (v57 + (v60 + (v63 + v43)))
					v68 = int32(4)
					v69 = v44 + v68
					v71 = v48 + v68
					if v71 != v32&int32(2147483644) {
						v43 = v67
						v44 = v69
						v48 = v71
						continue
					} else {
						break
					}
					break
				}
				v73 = v67
				v74 = v69
			} else {
				v73 = v31
				v74 = int32(0)
			}
			if v34 == int32(0) {
				v100 = v73
			} else {
				v82 = v73
				v83 = v74
				v85 = int32(0)
				for {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v83<<(uint(int32(2))%32))+uint32(_c_F_updateMaxmemory[5])))
					v94 = v93 + v82
					v95 = int32(1)
					v98 = v85 + v95
					if v98 != v34 {
						v82 = v94
						v83 = v83 + v95
						v85 = v98
						continue
					} else {
						break
					}
					break
				}
				v100 = v94
			}
		}
		v111 = int32(_a_F_updateMaxmemory_0)
		v112 = *(*int64)(unsafe.Add(mBase, _c_F_updateMaxmemory[7]))
		v114 = *(*int32)(unsafe.Add(mBase, _c_F_updateMaxmemory[8]))
		if base.I64_extend_i32_u(v114) <= v112 {
			v129 = int32(0)
		} else {
			v119 = base.I64_div_s(v112, int64(16384))
			v126 = v114 - base.I32_wrap_i64(v112+v119*int64(44)) + int32(-44)
			if base.Ui32(v114) < base.Ui32(v126) {
				v128 = int32(0)
			} else {
				v128 = v126
			}
			v129 = v128
		}
		v131 = *(*int32)(unsafe.Add(mBase, _c_F_updateMaxmemory[9]))
		if v131 == int32(0) {
			v138 = v129
		} else {
			v135 = *(*int32)(unsafe.Add(mBase, _c_F_updateMaxmemory[10]))
			v136 = F_sdsAllocSize(m, v135)
			mBase = m.M
			v138 = v136 + v129
		}
		v139 = F_clusterIsAnySlotExporting(m)
		mBase = m.M
		if v139 == int32(0) {
			v144 = v138
		} else {
			v142 = F_clusterGetTotalSlotExportBufferMemory(m)
			mBase = m.M
			v144 = v142 + v138
		}
		v146 = *(*int64)(unsafe.Add(mBase, _c_F_updateMaxmemory[0]))
		v147 = v100 - v144
		if base.Ui64(base.I64_extend_i32_u(v147)) <= base.Ui64(v146) {
			F_startEvictionTimeProc(m)
			mBase = m.M
			v163 = m.ExcPending
			if v163 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return int32(1)
			}
		} else {
			v151 = *(*int32)(unsafe.Add(mBase, _c_F_updateMaxmemory[11]))
			if int32(3) < v151 {
				F_startEvictionTimeProc(m)
				mBase = m.M
				v163 = m.ExcPending
				if v163 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return int32(1)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v147
				*(*int64)(unsafe.Add(mBase, uint32(v8))) = v146
				F__serverLog(m, int32(3), int32(_a_F_updateMaxmemory_1), v8)
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
					return int32(0)
				} else {
					F_startEvictionTimeProc(m)
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return int32(1)
					}
				}
			}
		}
	}
}
func F_updateOOMScoreAdj(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	v5 = F_setOOMScoreAdj(m, int32(-1))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != int32(-1) {
			v14 = int32(1)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a_F_updateOOMScoreAdj_0)
			v14 = int32(0)
		}
		return v14
	}
}
func F_updateRdmaPort(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v5 = F_listenerByType(m, int32(3))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v10 = int32(_a_F_updateRdmaPort_0)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+68)) = int32(_a_F_updateRdmaPort_1)
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_updateRdmaPort[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+72)) = v15
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_updateRdmaPort[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+76)) = v18
			v21 = F_connectionByType(m, int32(3))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+80)) = v21
				v25 = F_changeListener(m, v5)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v25 != int32(-1) {
						v35 = int32(1)
					} else {
						v30 = int32(_a_F_updateRdmaPort_2)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v30
						v35 = int32(0)
					}
					return v35
				}
			}
		} else {
			v30 = int32(_a_F_updateRdmaPort_3)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v30
			v35 = int32(0)
			return v35
		}
	}
}
func F_updateReplBacklogSize(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	F_resizeReplicationBacklog(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_updateShardId(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int64
	_ = v217
	var v223 int64
	_ = v223
	var v229 int64
	_ = v229
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int64
	_ = v328
	var v334 int64
	_ = v334
	var v340 int64
	_ = v340
	var v346 int64
	_ = v346
	var v352 int64
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_updateShardId_0), int32(_a_F_updateShardId_1), int32(7623))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L41
	} else {
		goto L87
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_updateShardId_0), int32(_a_F_updateShardId_1), int32(7623))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L41
	} else {
		goto L86
	}
L3:
	;
	m.G0 = v10 + int32(16)
	return
L4:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v14&int32(2) == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v143 = l0 + int32(48)
	v144 = int32(40)
	goto L48
L6:
	;
	v22 = l0
	goto L8
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_updateShardId[0]))
	if v32 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+2172))
	if v26 == int32(0) {
		v30 = v22
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v30 = v26
	goto L7
L10:
	;
	if v26 != l0 {
		v22 = v26
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v37 = v30 + int32(48)
	v38 = int32(40)
	goto L19
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+2172))
	if v35 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	if v102 == int32(0) {
		goto L5
	} else {
		goto L31
	}
L16:
	;
	v102 = int32(0)
	goto L15
L17:
	;
	v74 = v69
	v75 = v70
	v76 = v71
	goto L27
L18:
	;
	if v59 == int32(0) {
		goto L16
	} else {
		goto L25
	}
L19:
	;
	if (l1|v37)&int32(3) != 0 {
		v69 = v37
		v70 = l1
		v71 = v38
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v46 = v37
	v47 = l1
	v48 = v38
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v51 != v52 {
		v69 = v46
		v70 = v47
		v71 = v48
		goto L17
	} else {
		goto L23
	}
L22:
	;
	goto L18
L23:
	;
	v54 = int32(4)
	v55 = v47 + v54
	v57 = v46 + v54
	v59 = v48 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v59) {
		v46 = v57
		v47 = v55
		v48 = v59
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v69 = v57
	v70 = v55
	v71 = v59
	goto L17
L26:
	;
	v102 = v79 - v80
	goto L15
L27:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v79 != v80 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v82 = int32(1)
	v87 = v76 + int32(-1)
	if v87 == int32(0) {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v74 = v74 + v82
	v75 = v75 + v82
	v76 = v87
	goto L27
L31:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_updateShardId[1]))
	if int32(2) < v106 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v114 = l0
	goto L34
L33:
	;
	if v32 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+2172))
	if v118 == int32(0) {
		v122 = v114
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v122 = v118
	goto L33
L36:
	;
	if v118 != l0 {
		v114 = v118
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v122 + int32(48)
	F__serverLog(m, int32(2), int32(_a_F_updateShardId_2), v10)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)+2172))
	if v125 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	return
L42:
	;
	goto L3
L43:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_updateShardId[2]))
	if v251 == l0 {
		goto L3
	} else {
		goto L64
	}
L44:
	;
	if v208 == int32(0) {
		goto L43
	} else {
		goto L60
	}
L45:
	;
	v208 = int32(0)
	goto L44
L46:
	;
	v180 = v175
	v181 = v176
	v182 = v177
	goto L56
L47:
	;
	if v165 == int32(0) {
		goto L45
	} else {
		goto L54
	}
L48:
	;
	if (l1|v143)&int32(3) != 0 {
		v175 = v143
		v176 = l1
		v177 = v144
		goto L46
	} else {
		goto L49
	}
L49:
	;
	v152 = v143
	v153 = l1
	v154 = v144
	goto L50
L50:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	if v157 != v158 {
		v175 = v152
		v176 = v153
		v177 = v154
		goto L46
	} else {
		goto L52
	}
L51:
	;
	goto L47
L52:
	;
	v160 = int32(4)
	v161 = v153 + v160
	v163 = v152 + v160
	v165 = v154 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v165) {
		v152 = v163
		v153 = v161
		v154 = v165
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v175 = v163
	v176 = v161
	v177 = v165
	goto L46
L55:
	;
	v208 = v185 - v186
	goto L44
L56:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	if v185 != v186 {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v188 = int32(1)
	v193 = v182 + int32(-1)
	if v193 == int32(0) {
		goto L45
	} else {
		goto L59
	}
L59:
	;
	v180 = v180 + v188
	v181 = v181 + v188
	v182 = v193
	goto L56
L60:
	;
	F_clusterRemoveNodeFromShard(m, l0)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L41
	} else {
		goto L61
	}
L61:
	;
	v217 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(32))))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(80)))) = v217
	v223 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(24))))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(72)))) = v223
	v229 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(64)))) = v229
	v235 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(56)))) = v235
	v237 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v143))) = v237
	F_clusterAddNodeToShard(m, l1, l0)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L41
	} else {
		goto L62
	}
L62:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L41
	} else {
		goto L63
	}
L63:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_updateShardId[3]))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+uint32(_c_F_updateShardId[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+uint32(_c_F_updateShardId[4]))) = v245 | int32(4)
	goto L43
L64:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v251)+2172))
	if v253 != l0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	v256 = v251 + int32(48)
	v257 = int32(40)
	goto L70
L66:
	;
	if v321 == int32(0) {
		goto L3
	} else {
		goto L82
	}
L67:
	;
	v321 = int32(0)
	goto L66
L68:
	;
	v293 = v288
	v294 = v289
	v295 = v290
	goto L78
L69:
	;
	if v278 == int32(0) {
		goto L67
	} else {
		goto L76
	}
L70:
	;
	if (l1|v256)&int32(3) != 0 {
		v288 = v256
		v289 = l1
		v290 = v257
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v265 = v256
	v266 = l1
	v267 = v257
	goto L72
L72:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	if v270 != v271 {
		v288 = v265
		v289 = v266
		v290 = v267
		goto L68
	} else {
		goto L74
	}
L73:
	;
	goto L69
L74:
	;
	v273 = int32(4)
	v274 = v266 + v273
	v276 = v265 + v273
	v278 = v267 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v278) {
		v265 = v276
		v266 = v274
		v267 = v278
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v288 = v276
	v289 = v274
	v290 = v278
	goto L68
L77:
	;
	v321 = v298 - v299
	goto L66
L78:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v298 != v299 {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v301 = int32(1)
	v306 = v295 + int32(-1)
	if v306 == int32(0) {
		goto L67
	} else {
		goto L81
	}
L81:
	;
	v293 = v293 + v301
	v294 = v294 + v301
	v295 = v306
	goto L78
L82:
	;
	F_clusterRemoveNodeFromShard(m, v251)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L41
	} else {
		goto L83
	}
L83:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_updateShardId[2]))
	v328 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v327)+48)) = v328
	v334 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(32))))
	*(*int64)(unsafe.Add(mBase, uint32(v327+int32(80)))) = v334
	v340 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(24))))
	*(*int64)(unsafe.Add(mBase, uint32(v327+int32(72)))) = v340
	v346 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v327+int32(64)))) = v346
	v352 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v327+int32(56)))) = v352
	F_clusterAddNodeToShard(m, l1, v327)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L41
	} else {
		goto L84
	}
L84:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L41
	} else {
		goto L85
	}
L85:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_updateShardId[3]))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+uint32(_c_F_updateShardId[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v359)+uint32(_c_F_updateShardId[4]))) = v360 | int32(12)
	goto L3
L86:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_updateStatsOnUnblock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v34 int32
	_ = v34
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int64
	_ = v91
	var v96 int64
	_ = v96
	var v100 int32
	_ = v100
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v9 = l2 + l1 + v8
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+104))
	v13 = base.I64_extend_i32_s(v9)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v12 + v13
	F_clusterSlotStatsAddCpuDuration(m, l0, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)+112))
		v20 = int64(1)
		*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v19 + v20
		v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+256))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+256)) = v23 + v20
		v27 = int32(_a_F_updateStatsOnUnblock_0)
		v29 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[0]))
		*(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[0])) = v29 + v20
		v34 = *(*int32)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[1]))
		if base.Ui32(l3) < base.Ui32(int32(3)) {
			if l3 == int32(0) {
				v73 = *(*int32)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[2]))
				if v73 == int32(0) {
					v85 = v18
					F_commandlogPushCurrentCommand(m, l0, v85)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
						v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
						if v91 == int64(0) {
							return
						} else {
							v96 = base.I64_extend_i32_s(l2)
							if v96 < v91*int64(1000) {
								return
							} else {
								F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
					F_updateCommandLatencyHistogram(m, v18+int32(148), base.I64_extend_i32_s(v78*int32(1000)))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
						v85 = v84
						F_commandlogPushCurrentCommand(m, l0, v85)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
							v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
							if v91 == int64(0) {
								return
							} else {
								v96 = base.I64_extend_i32_s(l2)
								if v96 < v91*int64(1000) {
									return
								} else {
									F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			} else {
				if l3&int32(2) == int32(0) {
					if l3&int32(1) == int32(0) {
						if v34 == int32(0) {
							v73 = *(*int32)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[2]))
							if v73 == int32(0) {
								v85 = v18
								F_commandlogPushCurrentCommand(m, l0, v85)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
									v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
									if v91 == int64(0) {
										return
									} else {
										v96 = base.I64_extend_i32_s(l2)
										if v96 < v91*int64(1000) {
											return
										} else {
											F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
								F_updateCommandLatencyHistogram(m, v18+int32(148), base.I64_extend_i32_s(v78*int32(1000)))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
									v85 = v84
									F_commandlogPushCurrentCommand(m, l0, v85)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
										v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
										if v91 == int64(0) {
											return
										} else {
											v96 = base.I64_extend_i32_s(l2)
											if v96 < v91*int64(1000) {
												return
											} else {
												F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								}
							}
						} else {
							F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_updateStatsOnUnblock_2), int32(_a_F_updateStatsOnUnblock_3), int32(143))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					} else {
						v49 = int32(120)
						v50 = v18 + v49
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
						*(*int64)(unsafe.Add(mBase, uint32(v50))) = v51 + int64(1)
						v73 = *(*int32)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[2]))
						if v73 == int32(0) {
							v85 = v18
							F_commandlogPushCurrentCommand(m, l0, v85)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
								v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
								if v91 == int64(0) {
									return
								} else {
									v96 = base.I64_extend_i32_s(l2)
									if v96 < v91*int64(1000) {
										return
									} else {
										F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
							F_updateCommandLatencyHistogram(m, v18+int32(148), base.I64_extend_i32_s(v78*int32(1000)))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
								v85 = v84
								F_commandlogPushCurrentCommand(m, l0, v85)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
									v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
									if v91 == int64(0) {
										return
									} else {
										v96 = base.I64_extend_i32_s(l2)
										if v96 < v91*int64(1000) {
											return
										} else {
											F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
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
				} else {
					v49 = int32(128)
					v50 = v18 + v49
					v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
					*(*int64)(unsafe.Add(mBase, uint32(v50))) = v51 + int64(1)
					v73 = *(*int32)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[2]))
					if v73 == int32(0) {
						v85 = v18
						F_commandlogPushCurrentCommand(m, l0, v85)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
							v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
							if v91 == int64(0) {
								return
							} else {
								v96 = base.I64_extend_i32_s(l2)
								if v96 < v91*int64(1000) {
									return
								} else {
									F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						F_updateCommandLatencyHistogram(m, v18+int32(148), base.I64_extend_i32_s(v78*int32(1000)))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							v85 = v84
							F_commandlogPushCurrentCommand(m, l0, v85)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
								v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
								if v91 == int64(0) {
									return
								} else {
									v96 = base.I64_extend_i32_s(l2)
									if v96 < v91*int64(1000) {
										return
									} else {
										F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
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
		} else {
			if v34 != 0 {
				F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_updateStatsOnUnblock_4), int32(_a_F_updateStatsOnUnblock_3), int32(136))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if l3 == int32(0) {
					v73 = *(*int32)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[2]))
					if v73 == int32(0) {
						v85 = v18
						F_commandlogPushCurrentCommand(m, l0, v85)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
							v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
							if v91 == int64(0) {
								return
							} else {
								v96 = base.I64_extend_i32_s(l2)
								if v96 < v91*int64(1000) {
									return
								} else {
									F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						F_updateCommandLatencyHistogram(m, v18+int32(148), base.I64_extend_i32_s(v78*int32(1000)))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							v85 = v84
							F_commandlogPushCurrentCommand(m, l0, v85)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
								v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
								if v91 == int64(0) {
									return
								} else {
									v96 = base.I64_extend_i32_s(l2)
									if v96 < v91*int64(1000) {
										return
									} else {
										F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					}
				} else {
					if l3&int32(2) == int32(0) {
						if l3&int32(1) == int32(0) {
							if v34 == int32(0) {
								v73 = *(*int32)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[2]))
								if v73 == int32(0) {
									v85 = v18
									F_commandlogPushCurrentCommand(m, l0, v85)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
										v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
										if v91 == int64(0) {
											return
										} else {
											v96 = base.I64_extend_i32_s(l2)
											if v96 < v91*int64(1000) {
												return
											} else {
												F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
									F_updateCommandLatencyHistogram(m, v18+int32(148), base.I64_extend_i32_s(v78*int32(1000)))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
										v85 = v84
										F_commandlogPushCurrentCommand(m, l0, v85)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
											v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
											if v91 == int64(0) {
												return
											} else {
												v96 = base.I64_extend_i32_s(l2)
												if v96 < v91*int64(1000) {
													return
												} else {
													F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return
													} else {
														return
													}
												}
											}
										}
									}
								}
							} else {
								F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_updateStatsOnUnblock_2), int32(_a_F_updateStatsOnUnblock_3), int32(143))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						} else {
							v49 = int32(120)
							v50 = v18 + v49
							v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
							*(*int64)(unsafe.Add(mBase, uint32(v50))) = v51 + int64(1)
							v73 = *(*int32)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[2]))
							if v73 == int32(0) {
								v85 = v18
								F_commandlogPushCurrentCommand(m, l0, v85)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
									v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
									if v91 == int64(0) {
										return
									} else {
										v96 = base.I64_extend_i32_s(l2)
										if v96 < v91*int64(1000) {
											return
										} else {
											F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
								F_updateCommandLatencyHistogram(m, v18+int32(148), base.I64_extend_i32_s(v78*int32(1000)))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
									v85 = v84
									F_commandlogPushCurrentCommand(m, l0, v85)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
										v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
										if v91 == int64(0) {
											return
										} else {
											v96 = base.I64_extend_i32_s(l2)
											if v96 < v91*int64(1000) {
												return
											} else {
												F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
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
					} else {
						v49 = int32(128)
						v50 = v18 + v49
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
						*(*int64)(unsafe.Add(mBase, uint32(v50))) = v51 + int64(1)
						v73 = *(*int32)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[2]))
						if v73 == int32(0) {
							v85 = v18
							F_commandlogPushCurrentCommand(m, l0, v85)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
								v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
								if v91 == int64(0) {
									return
								} else {
									v96 = base.I64_extend_i32_s(l2)
									if v96 < v91*int64(1000) {
										return
									} else {
										F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
							F_updateCommandLatencyHistogram(m, v18+int32(148), base.I64_extend_i32_s(v78*int32(1000)))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
								v85 = v84
								F_commandlogPushCurrentCommand(m, l0, v85)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
									v91 = *(*int64)(unsafe.Add(mBase, _c_F_updateStatsOnUnblock[3]))
									if v91 == int64(0) {
										return
									} else {
										v96 = base.I64_extend_i32_s(l2)
										if v96 < v91*int64(1000) {
											return
										} else {
											F_latencyAddSample(m, int32(_a_F_updateStatsOnUnblock_1), v96)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
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
}
func F_updateWatchdogPeriod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_updateWatchdogPeriod[0]))
	if v10 != 0 {
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_updateWatchdogPeriod[1]))
		v25 = base.I32_div_s(int32(1000), v24)
		v27 = v25 << (uint(int32(1)) % 32)
		if v27 <= v10 {
			v31 = v10
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_updateWatchdogPeriod[0])) = v27
			v31 = v27
		}
		v32 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v32
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(0)
		v36 = int32(1000)
		v37 = base.I32_div_s(v31, v36)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = base.I64_extend_i32_s(v37)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = (v31 - v37*v36) * v36
		v48 = F_setitimer(m, v32, v7, v32)
		mBase = m.M
	} else {
		v11 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v11
		v13 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v13
		v21 = F_setitimer(m, v11, v7, v11)
		mBase = m.M
	}
	m.G0 = v7 + int32(32)
	return int32(1)
}
func F_usage(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_usage[0]))
	v7 = F_fwrite(m, int32(_a_F_usage_0), int32(60), int32(1), v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v12 = F_fwrite(m, int32(_a_F_usage_1), int32(50), int32(1), v6)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v17 = F_fwrite(m, int32(_a_F_usage_2), int32(39), int32(1), v6)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v22 = F_fwrite(m, int32(_a_F_usage_3), int32(36), int32(1), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v27 = F_fwrite(m, int32(_a_F_usage_4), int32(49), int32(1), v6)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v32 = F_fwrite(m, int32(_a_F_usage_5), int32(38), int32(1), v6)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							v35 = F_fputc(m, int32(10), v6)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								v40 = F_fwrite(m, int32(_a_F_usage_6), int32(10), int32(1), v6)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									v45 = F_fwrite(m, int32(_a_F_usage_7), int32(58), int32(1), v6)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										v50 = F_fwrite(m, int32(_a_F_usage_8), int32(50), int32(1), v6)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return
										} else {
											v55 = F_fwrite(m, int32(_a_F_usage_9), int32(45), int32(1), v6)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return
											} else {
												v60 = F_fwrite(m, int32(_a_F_usage_10), int32(35), int32(1), v6)
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													v65 = F_fwrite(m, int32(_a_F_usage_11), int32(62), int32(1), v6)
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return
													} else {
														v70 = F_fwrite(m, int32(_a_F_usage_12), int32(63), int32(1), v6)
														mBase = m.M
														v71 = m.ExcPending
														if v71 != 0 {
															return
														} else {
															v75 = F_fwrite(m, int32(_a_F_usage_13), int32(62), int32(1), v6)
															mBase = m.M
															v76 = m.ExcPending
															if v76 != 0 {
																return
															} else {
																v80 = F_fwrite(m, int32(_a_F_usage_14), int32(15), int32(1), v6)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return
																} else {
																	v85 = F_fwrite(m, int32(_a_F_usage_15), int32(53), int32(1), v6)
																	mBase = m.M
																	v86 = m.ExcPending
																	if v86 != 0 {
																		return
																	} else {
																		m.Env.Exit(m, int32(0))
																		mBase = m.M
																		base.Wasm_trap_unreachable()
																		for {
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
							}
						}
					}
				}
			}
		}
	}
}
func F_useDisklessLoad(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_useDisklessLoad[0]))
	if v4&int32(-2) == int32(2) {
		v43 = F_moduleAllDatatypesHandleErrors(m)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			if v43 != 0 {
				v54 = int32(1)
				v56 = *(*int32)(unsafe.Add(mBase, _c_F_useDisklessLoad[0]))
				if v56 != int32(2) {
					v74 = v54
					return v74
				} else {
					v59 = F_moduleAllModulesHandleReplAsyncLoad(m)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						if v59 != 0 {
							v74 = v54
							return v74
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, _c_F_useDisklessLoad[1]))
							if int32(2) < v62 {
								return int32(0)
							} else {
								v66 = int32(_a_F_useDisklessLoad_0)
								v68 = int32(0)
								F__serverLog(m, int32(2), v66, v68)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									v74 = v68
									return v74
								}
							}
						}
					}
				}
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, _c_F_useDisklessLoad[1]))
				if int32(2) < v48 {
					return int32(0)
				} else {
					v66 = int32(_a_F_useDisklessLoad_1)
					v68 = int32(0)
					F__serverLog(m, int32(2), v66, v68)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v74 = v68
						return v74
					}
				}
			}
		}
	} else {
		v9 = int32(0)
		if v4 != int32(1) {
			v74 = v9
			return v74
		} else {
			v12 = int64(0)
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_useDisklessLoad[2]))
			if v16 < int32(1) {
				v38 = v12
			} else {
				v20 = v12
				v21 = int32(0)
				for {
					v22 = F_dbHasNoKeys(m, v21)
					mBase = m.M
					if v22 != 0 {
						v32 = v20
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, _c_F_useDisklessLoad[3]))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v21<<(uint(int32(2))%32))))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						v30 = F_kvstoreSize(m, v29)
						mBase = m.M
						v32 = v30 + v20
					}
					v34 = v21 + int32(1)
					v36 = *(*int32)(unsafe.Add(mBase, _c_F_useDisklessLoad[2]))
					if v34 < v36 {
						v20 = v32
						v21 = v34
						continue
					} else {
						break
					}
					break
				}
				v38 = v32
			}
			if v38 != int64(0) {
				v74 = v9
				return v74
			} else {
				v43 = F_moduleAllDatatypesHandleErrors(m)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					if v43 != 0 {
						v54 = int32(1)
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_useDisklessLoad[0]))
						if v56 != int32(2) {
							v74 = v54
							return v74
						} else {
							v59 = F_moduleAllModulesHandleReplAsyncLoad(m)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								if v59 != 0 {
									v74 = v54
									return v74
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, _c_F_useDisklessLoad[1]))
									if int32(2) < v62 {
										return int32(0)
									} else {
										v66 = int32(_a_F_useDisklessLoad_0)
										v68 = int32(0)
										F__serverLog(m, int32(2), v66, v68)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											v74 = v68
											return v74
										}
									}
								}
							}
						}
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, _c_F_useDisklessLoad[1]))
						if int32(2) < v48 {
							return int32(0)
						} else {
							v66 = int32(_a_F_useDisklessLoad_1)
							v68 = int32(0)
							F__serverLog(m, int32(2), v66, v68)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								v74 = v68
								return v74
							}
						}
					}
				}
			}
		}
	}
}
func F_usleep(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
	v10 = int32(1000000)
	v11 = base.I32_div_u_s(l0, v10)
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = base.I64_extend_i32_u(v11)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = (l0 - v11*v10) * int32(1000)
	v20 = F_nanosleep(m, v6, v6)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v20
	}
}
