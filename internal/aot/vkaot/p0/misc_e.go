package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F__Exit(m *base.Module, l0 int32) {
	m.Wasi_snapshot_preview1.Proc_exit(m, l0)
	base.Wasm_trap_unreachable()
	for {
	}
}
func F___errno_location(m *base.Module) int32 {
	return int32(9116376)
}
func F___extenddftf2(m *base.Module, l0 int32, l1 float64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v19 int64
	_ = v19
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v72 int64
	_ = v72
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = base.I64_reinterpret_f64(l1)
	v15 = v13 & int64(4503599627370495)
	v19 = int64(base.Ui64(v13)>>(uint(int64(52))%64)) & int64(2047)
	if v19 == int64(0) {
		if base.B2i32(v15 == int64(0)) == int32(0) {
			v42 = int64(0)
			if base.Ui64(v15) < base.Ui64(int64(4294967296)) {
				v53 = base.I32_clz(base.I32_wrap_i64(v13)) | int32(32)
			} else {
				v53 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v15) >> (uint(int64(32)) % 64))))
			}
			v55 = v53 + int32(49)
			if v55&int32(64) == int32(0) {
				if v55 == int32(0) {
					v76 = v15
					v77 = v42
				} else {
					v72 = base.I64_extend_i32_u(v55)
					v76 = v15 << (uint(v72) % 64)
					v77 = int64(base.Ui64(v15)>>(uint(base.I64_extend_i32_u(int32(64)-v55))%64)) | v42<<(uint(v72)%64)
				}
			} else {
				v76 = int64(0)
				v77 = v15 << (uint(base.I64_extend_i32_u(v53+int32(-15))) % 64)
			}
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = v76
			*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v77
			v86 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(8))))
			v89 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
			v90 = v89
			v91 = base.I64_extend_i32_u(int32(15372) - v53)
			v92 = v86 ^ int64(281474976710656)
		} else {
			v39 = int64(0)
			v90 = v39
			v91 = v39
			v92 = v39
		}
	} else {
		if v19 == int64(2047) {
			v90 = v15 << (uint(int64(60)) % 64)
			v91 = int64(32767)
			v92 = int64(base.Ui64(v15) >> (uint(int64(4)) % 64))
		} else {
			v90 = v15 << (uint(int64(60)) % 64)
			v91 = v19 + int64(15360)
			v92 = int64(base.Ui64(v15) >> (uint(int64(4)) % 64))
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v90
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v91<<(uint(int64(48))%64) | v13&int64(-9223372036854775807-1) | v92
	m.G0 = v11 + int32(16)
	return
}
func F___extendsftf2(m *base.Module, l0 int32, l1 float32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v62 int64
	_ = v62
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = base.I32_reinterpret_f32(l1)
	v16 = v14 & int32(8388607)
	v18 = int32(base.Ui32(v14) >> (uint(int32(23)) % 32))
	v20 = v18 & int32(255)
	if v20 == int32(0) {
		if v16 != 0 {
			v41 = base.I64_extend_i32_u(v16)
			v42 = int64(0)
			v43 = base.I32_clz(v16)
			v45 = v43 + int32(81)
			if v45&int32(64) == int32(0) {
				if v45 == int32(0) {
					v66 = v41
					v67 = v42
				} else {
					v62 = base.I64_extend_i32_u(v45)
					v66 = v41 << (uint(v62) % 64)
					v67 = int64(base.Ui64(v41)>>(uint(base.I64_extend_i32_u(int32(64)-v45))%64)) | v42<<(uint(v62)%64)
				}
			} else {
				v66 = int64(0)
				v67 = v41 << (uint(base.I64_extend_i32_u(v43+int32(17))) % 64)
			}
			*(*int64)(unsafe.Add(mBase, uint32(v12))) = v66
			*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v67
			v75 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(8))))
			v78 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
			v79 = int32(16265) - v43
			v80 = v75 ^ int64(281474976710656)
			v81 = v78
		} else {
			v38 = int64(0)
			v79 = int32(0)
			v80 = v38
			v81 = v38
		}
	} else {
		if v20 == int32(255) {
			v79 = int32(32767)
			v80 = base.I64_extend_i32_u(v16) << (uint(int64(25)) % 64)
			v81 = int64(0)
		} else {
			v79 = v18&int32(255) + int32(16256)
			v80 = base.I64_extend_i32_u(v16) << (uint(int64(25)) % 64)
			v81 = int64(0)
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v81
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = base.I64_extend_i32_u(v79)<<(uint(int64(48))%64) | base.I64_extend_i32_u(int32(base.Ui32(v14)>>(uint(int32(31))%32)))<<(uint(int64(63))%64) | v80
	m.G0 = v12 + int32(16)
	return
}
func F_emptyData(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v222 int64
	_ = v222
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = int64(1)
	v19 = int32(1)
	v20 = l1 & v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v20 ^ v19
	if l0 < int32(-1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v222
L2:
	;
	F_moduleFireServerEvent(m, int64(2), int32(0), v14)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	goto L6
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_emptyData[0]))
	if l0 < v27 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emptyData[1])) = int32(28)
	v222 = int64(-1)
	goto L1
L7:
	;
	return int64(0)
L8:
	;
	v41 = base.B2i32(l0 == int32(-1))
	if l0 == int32(-1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	F_trackingInvalidateKeysOnFlush(m, v20)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L24
	}
L10:
	;
	v42 = int32(0)
	goto L12
L11:
	;
	v42 = l0
	goto L12
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_emptyData[0]))
	if l0 == int32(-1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = v44 + int32(-1)
	goto L15
L14:
	;
	v47 = l0
	goto L15
L15:
	;
	if v47 < v42 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_emptyData[2]))
	v58 = v42
	v60 = v50
	goto L17
L17:
	;
	v63 = v58 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60+v63)))
	if v65 == int32(0) {
		v80 = v60
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L9
L19:
	;
	if v58 != v47 {
		v58 = v58 + int32(1)
		v60 = v80
		goto L17
	} else {
		goto L23
	}
L20:
	;
	F_scanDatabaseForDeletedKeys(m, v65, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_emptyData[2]))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v72+v63)))
	F_touchAllWatchedKeysInDb(m, v74, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_emptyData[2]))
	v80 = v79
	goto L19
L23:
	;
	goto L18
L24:
	;
	v97 = int32(0)
	v100 = m.G0
	v102 = v100 - int32(16)
	m.G0 = v102
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_emptyData[3]))
	if v106 == v97 {
		v140 = v97
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v196 = l1 & int32(2)
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_emptyData[2]))
	v199 = F_emptyDbStructure(m, v198, l0, v20, l2)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L7
	} else {
		goto L52
	}
L26:
	;
	F_clusterHandleFlushDuringSlotMigration(m)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L51
	}
L27:
	;
	if v140 != 0 {
		goto L26
	} else {
		goto L38
	}
L28:
	;
	m.G0 = v102 + int32(16)
	goto L27
L29:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_emptyData[4]))
	if v110 == int32(0) {
		v140 = v97
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_emptyData[5])))
	v115 = v102 + int32(8)
	F_listRewind(m, v113, v115)
	mBase = m.M
	v117 = int32(0)
	v120 = F_listNext(m, v115)
	mBase = m.M
	if v120 == v117 {
		v140 = v117
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v125 = v120
	goto L32
L32:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v127 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v140 = v117
	goto L28
L34:
	;
	v138 = F_listNext(m, v102+int32(8))
	mBase = m.M
	if v138 != 0 {
		v125 = v138
		goto L32
	} else {
		goto L37
	}
L35:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+156))
	if base.Ui32(v130+int32(-18)) <= base.Ui32(int32(2)) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v140 = int32(1)
	goto L28
L37:
	;
	goto L33
L38:
	;
	v145 = int32(0)
	v148 = m.G0
	v150 = v148 - int32(16)
	m.G0 = v150
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_emptyData[3]))
	if v154 == v145 {
		v186 = v145
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v186 == int32(0) {
		goto L25
	} else {
		goto L50
	}
L40:
	;
	m.G0 = v150 + int32(16)
	goto L39
L41:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_emptyData[4]))
	if v158 == int32(0) {
		v186 = v145
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)+uint32(_c_F_emptyData[5])))
	v163 = v150 + int32(8)
	F_listRewind(m, v161, v163)
	mBase = m.M
	v165 = int32(0)
	v168 = F_listNext(m, v163)
	mBase = m.M
	if v168 == v165 {
		v186 = v165
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v173 = v168
	goto L44
L44:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if v175 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v186 = v165
	goto L40
L46:
	;
	v184 = F_listNext(m, v150+int32(8))
	mBase = m.M
	if v184 != 0 {
		v173 = v184
		goto L44
	} else {
		goto L49
	}
L47:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v174)+156))
	if base.Ui32(v176+int32(-18)) <= base.Ui32(int32(2)) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v186 = int32(1)
	goto L40
L49:
	;
	goto L45
L50:
	;
	goto L26
L51:
	;
	goto L25
L52:
	;
	if l0 != int32(-1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	F_moduleFireServerEvent(m, int64(2), int32(1), v14)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L7
	} else {
		goto L61
	}
L54:
	;
	if v196 != 0 {
		goto L53
	} else {
		goto L59
	}
L55:
	;
	F_flushReplicaKeysWithExpireList(m, v20)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	if v196 != 0 {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	F_functionReset(m, v20, l2)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	goto L53
L59:
	;
	F__serverAssert(m, int32(_a_F_emptyData_0), int32(_a_F_emptyData_1), int32(688))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	v222 = v199
	goto L1
}
func F_emptyDbStructure(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v98 int64
	_ = v98
	var v109 int64
	_ = v109
	v13 = base.B2i32(l1 == int32(-1))
	if l1 == int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v109
L2:
	;
	v14 = int32(0)
	goto L4
L3:
	;
	v14 = l1
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_emptyDbStructure[0]))
	if l1 == int32(-1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v19 = v16 + int32(-1)
	goto L7
L6:
	;
	v19 = l1
	goto L7
L7:
	;
	if v19 < v14 {
		v109 = int64(0)
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v26 = int64(0)
	v28 = v14
	goto L9
L9:
	;
	v33 = l0 + v28<<(uint(int32(2))%32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v34 == int32(0) {
		v98 = v26
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v109 = v98
	goto L1
L11:
	;
	if v28 != v19 {
		v26 = v98
		v28 = v28 + int32(1)
		goto L9
	} else {
		goto L32
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v38 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v48 == int64(0) {
		v98 = v26
		goto L11
	} else {
		goto L18
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v37)+40))
	v48 = v41
	goto L13
L16:
	;
	v45 = F_hashtableSize(m, v43)
	mBase = m.M
	v48 = base.I64_extend_i32_u(v45)
	goto L13
L17:
	;
	v48 = int64(0)
	goto L13
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	if v53 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if l2 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v58 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v52)+40))
	v63 = v56
	goto L19
L22:
	;
	v60 = F_hashtableSize(m, v58)
	mBase = m.M
	v63 = base.I64_extend_i32_u(v60)
	goto L19
L23:
	;
	v63 = int64(0)
	goto L19
L24:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v84 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v83)+32)) = v84
	*(*int64)(unsafe.Add(mBase, uint32(v83+int32(56)))) = v84
	*(*int64)(unsafe.Add(mBase, uint32(v83+int32(48)))) = v84
	*(*int64)(unsafe.Add(mBase, uint32(v83+int32(40)))) = v84
	v98 = v63 + v26
	goto L11
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	F_kvstoreEmpty(m, v71, l3)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L27
	} else {
		goto L29
	}
L26:
	;
	F_emptyDbAsync(m, v64)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int64(0)
L28:
	;
	goto L24
L29:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	F_kvstoreEmpty(m, v75, l3)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	F_kvstoreEmpty(m, v79, l3)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	goto L10
}
func F_enumConfigSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v229 int32
	_ = v229
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v380 int32
	_ = v380
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v435 int32
	_ = v435
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v398 == int32(0) {
		v406 = v17
		goto L106
	} else {
		goto L107
	}
L2:
	;
	v135 = F_sdsnew(m, int32(_a_F_enumConfigSet_0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L34
	} else {
		goto L35
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l2 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if int32(1) <= l2 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v17&int32(8) == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v27 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v391 = int32(0)
	goto L1
L9:
	;
	v30 = int32(0)
	v38 = v30
	v40 = v30
	goto L10
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1+v40<<(uint(int32(2))%32))))
	v56 = v38
	v60 = v27
	v61 = v16
	v62 = int32(1)
	goto L13
L11:
	;
	if v112 != int32(-2147483648) {
		v391 = v112
		goto L1
	} else {
		goto L33
	}
L12:
	;
	v117 = v40 + int32(1)
	if v117 != l2 {
		v38 = v112
		v40 = v117
		goto L10
	} else {
		goto L32
	}
L13:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v65 != 0 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	if v62&int32(1) != 0 {
		goto L2
	} else {
		goto L31
	}
L15:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v107 != 0 {
		v60 = v107
		v61 = v61 + int32(8)
		goto L13
	} else {
		goto L30
	}
L16:
	;
	if v97-v99 != 0 {
		goto L15
	} else {
		goto L28
	}
L17:
	;
	v97 = F_tolower(m, v93)
	mBase = m.M
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v99 = F_tolower(m, v98)
	mBase = m.M
	goto L16
L18:
	;
	v67 = v48
	v68 = v60
	v69 = v65
	goto L21
L19:
	;
	v93 = int32(0)
	v94 = v60
	goto L17
L20:
	;
	v93 = v90 & int32(255)
	v94 = v89
	goto L17
L21:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v71 == int32(0) {
		v89 = v68
		v90 = v69
		goto L20
	} else {
		goto L23
	}
L22:
	;
	v89 = v83
	v90 = int32(0)
	goto L20
L23:
	;
	v75 = v69 & int32(255)
	if v75 == v71 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v82 = int32(1)
	v83 = v68 + v82
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v84 != 0 {
		v67 = v67 + v82
		v68 = v83
		v69 = v84
		goto L21
	} else {
		goto L27
	}
L25:
	;
	v77 = F_tolower(m, v75)
	mBase = m.M
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v79 = F_tolower(m, v78)
	mBase = m.M
	if v77 == v79 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v89 = v68
	v90 = v81
	goto L20
L27:
	;
	goto L22
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v102 = v101 | v56
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v103 != 0 {
		v56 = v102
		v60 = v103
		v61 = v61 + int32(8)
		v62 = int32(0)
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v112 = v102
	goto L12
L30:
	;
	goto L14
L31:
	;
	v112 = v56
	goto L12
L32:
	;
	goto L11
L33:
	;
	goto L2
L34:
	;
	return int32(0)
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v140 == int32(0) {
		v229 = v135
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v242 = v229 + int32(-1)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	v245 = v243 & int32(7)
	switch v245 {
	case 0:
		goto L66
	case 1:
		goto L65
	case 2:
		goto L64
	case 3:
		goto L63
	case 4:
		goto L62
	default:
		goto L60
	}
L37:
	;
	v152 = v135
	v153 = v139
	v154 = v140
	goto L38
L38:
	;
	if v154&int32(3) == int32(0) {
		v177 = v154
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v229 = v215
	goto L36
L40:
	;
	v211 = F_sdscatlen(m, v152, v154, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L34
	} else {
		goto L56
	}
L41:
	;
	v210 = v202 - v154
	goto L40
L42:
	;
	v181 = v177
	goto L50
L43:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v163 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v166 = v154
	goto L46
L45:
	;
	v210 = v154 - v154
	goto L40
L46:
	;
	v170 = v166 + int32(1)
	if v170&int32(3) == int32(0) {
		v177 = v170
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v175 != 0 {
		v166 = v170
		goto L46
	} else {
		goto L49
	}
L49:
	;
	v202 = v170
	goto L41
L50:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v190 = int32(-2139062144)
	if (int32(16843008)-v187|v187)&v190 == v190 {
		v181 = v181 + int32(4)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v196 = v181
	goto L53
L52:
	;
	goto L51
L53:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	if v200 != 0 {
		v196 = v196 + int32(1)
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v202 = v196
	goto L41
L55:
	;
	goto L54
L56:
	;
	v215 = F_sdscatlen(m, v211, int32(_a_F_enumConfigSet_1), int32(2))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L34
	} else {
		goto L57
	}
L57:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	if v217 != 0 {
		v152 = v215
		v153 = v153 + int32(8)
		v154 = v217
		goto L38
	} else {
		goto L58
	}
L58:
	;
	goto L39
L59:
	;
	goto L96
L60:
	;
	goto L59
L61:
	;
	if v260 == int32(0) {
		goto L60
	} else {
		goto L67
	}
L62:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v229+int32(-17))))
	v260 = v259
	goto L61
L63:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v229+int32(-9))))
	v260 = v256
	goto L61
L64:
	;
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229+int32(-5)))))
	v260 = v253
	goto L61
L65:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+int32(-3)))))
	v260 = v250
	goto L61
L66:
	;
	v260 = int32(base.Ui32(v243) >> (uint(int32(3)) % 32))
	goto L61
L67:
	;
	v266 = int32(-1)&v260 + int32(-3)
	v270 = int32(0)&v260 + int32(0)
	v273 = v266 - v270 + int32(1)
	switch v245 {
	default:
		goto L73
	case 1:
		goto L72
	case 2:
		goto L71
	case 3:
		goto L70
	case 4:
		goto L69
	}
L68:
	;
	v289 = int32(0)
	v291 = base.B2i32(base.Ui32(v270) < base.Ui32(v288))
	if base.Ui32(v270) < base.Ui32(v288) {
		goto L75
	} else {
		goto L76
	}
L69:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v229+int32(-17))))
	v288 = v287
	goto L68
L70:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v229+int32(-9))))
	v288 = v284
	goto L68
L71:
	;
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229+int32(-5)))))
	v288 = v281
	goto L68
L72:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+int32(-3)))))
	v288 = v278
	goto L68
L73:
	;
	v288 = int32(base.Ui32(v243) >> (uint(int32(3)) % 32))
	goto L68
L74:
	;
	v305 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v229+v299))) = uint8(v305)
	switch v245 {
	default:
		goto L92
	case 1:
		goto L91
	case 2:
		goto L90
	case 3:
		goto L89
	case 4:
		goto L88
	}
L75:
	;
	v292 = v270
	goto L77
L76:
	;
	v292 = v289
	goto L77
L77:
	;
	v293 = v288 - v292
	if base.Ui32(v273) < base.Ui32(v293) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v295 = v273
	goto L80
L79:
	;
	v295 = v293
	goto L80
L80:
	;
	if v266 < v270 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v297 = v289
	goto L83
L82:
	;
	v297 = v295
	goto L83
L83:
	;
	if base.Ui32(v270) < base.Ui32(v288) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v299 = v297
	goto L86
L85:
	;
	v299 = int32(0)
	goto L86
L86:
	;
	if v299 == int32(0) {
		goto L74
	} else {
		goto L87
	}
L87:
	;
	v303 = F_memmove(m, v229, v229+v292, v299)
	mBase = m.M
	goto L74
L88:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v229+int32(-17)))) = base.I64_extend_i32_u(v299)
	goto L60
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229+int32(-9)))) = v299
	goto L59
L90:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v229+int32(-5)))) = uint16(v299)
	goto L59
L91:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v229+int32(-3)))) = uint8(v299)
	goto L59
L92:
	;
	v308 = v299 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v242))) = uint8(v308)
	goto L59
L93:
	;
	F_sdsfree(m, v229)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L34
	} else {
		goto L105
	}
L94:
	;
	goto L93
L95:
	;
	v366 = v345
	goto L102
L96:
	;
	v341 = int32(_a_F_enumConfigSet_2)
	v343 = int32(256)
	v345 = v229
	goto L98
L97:
	;
	v356 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v341))) = uint8(v356)
	goto L95
L98:
	;
	v347 = v343 + int32(-1)
	if v347 == int32(0) {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345))))
	*(*uint8)(unsafe.Add(mBase, uint32(v341))) = uint8(v350)
	v352 = int32(1)
	if v350 != 0 {
		v341 = v341 + v352
		v343 = v347
		v345 = v345 + v352
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L94
L102:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	if v368 != 0 {
		v366 = v366 + int32(1)
		goto L102
	} else {
		goto L104
	}
L103:
	;
	goto L94
L104:
	;
	goto L103
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a_F_enumConfigSet_2)
	return int32(0)
L106:
	;
	if v406&int32(256) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L107:
	;
	v401 = m.T0[v398].(func(*base.Module, int32, int32) int32)(m, v391, l3)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L34
	} else {
		goto L109
	}
L108:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v406 = v405
	goto L106
L109:
	;
	if v401 != 0 {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	return int32(0)
L111:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v416 == v391 {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v414)))
	v416 = v415
	goto L111
L113:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v412 = F_getModuleEnumConfig(m, v411)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L34
	} else {
		goto L114
	}
L114:
	;
	v416 = v412
	goto L111
L115:
	;
	if v417&int32(512) != 0 {
		goto L120
	} else {
		goto L121
	}
L116:
	;
	if v417&int32(256) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v427))) = v391
	return int32(1)
L118:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v424 = F_setModuleEnumConfig(m, v423, v391, l3)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L34
	} else {
		goto L119
	}
L119:
	;
	return v424
L120:
	;
	v435 = int32(1)
	goto L122
L121:
	;
	v435 = int32(2)
	goto L122
L122:
	;
	return v435
}
func F_evalExtractShebangFlags(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v229 int64
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
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
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v381 int32
	_ = v381
	var v383 int64
	_ = v383
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v490 int32
	_ = v490
	var v500 int64
	_ = v500
	var v502 int32
	_ = v502
	var v516 int64
	_ = v516
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v578 int32
	_ = v578
	var v590 int64
	_ = v590
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
	var v681 int32
	_ = v681
	var v682 int64
	_ = v682
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v695 int32
	_ = v695
	var v701 int64
	_ = v701
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	if l2 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v767 + int32(32)
	return v762
L2:
	;
	if l4 == int32(0) {
		v755 = v364
		goto L174
	} else {
		goto L175
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_evalExtractShebangFlags_0), int32(_a_F_evalExtractShebangFlags_1), int32(261))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L18
	} else {
		goto L173
	}
L4:
	;
	F__serverAssert(m, int32(_a_F_evalExtractShebangFlags_2), int32(_a_F_evalExtractShebangFlags_1), int32(240))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L18
	} else {
		goto L172
	}
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v24 != int32(35) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l3 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L7:
	;
	v681 = int32(0)
	v682 = int64(16)
	if l1 == v681 {
		v695 = v681
		v701 = v682
		goto L6
	} else {
		goto L168
	}
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v27 != int32(33) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v30 = int32(10)
	v31 = F___strchrnul(m, l0, v30)
	mBase = m.M
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v33 == v30 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v762 = int32(-1)
	v767 = v668
	goto L1
L11:
	;
	v47 = v37 - l0
	v50 = F_sdsnsplitargs(m, l0, v47, v20+int32(28))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L18
	} else {
		goto L20
	}
L12:
	;
	if v37 != 0 {
		goto L11
	} else {
		goto L16
	}
L13:
	;
	v37 = v31
	goto L15
L14:
	;
	v37 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	if l4 == int32(0) {
		v668 = v20
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v41 = F_sdsnew(m, int32(_a_F_evalExtractShebangFlags_3))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v41
	v762 = int32(-1)
	v767 = v20
	goto L1
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if v50 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	F_sdsfreesplitres(m, v50, v644)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L18
	} else {
		goto L167
	}
L22:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v644 = v643
	goto L21
L23:
	;
	if l1 == int32(0) {
		v160 = v52
		goto L36
	} else {
		goto L37
	}
L24:
	;
	if l4 == int32(0) {
		v644 = v52
		goto L21
	} else {
		goto L34
	}
L25:
	;
	if v52 == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+int32(-1)))))
	v62 = v60 & int32(7)
	switch v62 {
	case 0:
		goto L32
	case 1:
		goto L31
	case 2:
		goto L30
	case 3:
		goto L29
	case 4:
		goto L28
	default:
		goto L24
	}
L27:
	;
	if base.Ui32(int32(1)) < base.Ui32(v77) {
		goto L23
	} else {
		goto L33
	}
L28:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v57+int32(-17))))
	v77 = v76
	goto L27
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v57+int32(-9))))
	v77 = v73
	goto L27
L30:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57+int32(-5)))))
	v77 = v70
	goto L27
L31:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+int32(-3)))))
	v77 = v67
	goto L27
L32:
	;
	v77 = int32(base.Ui32(v60) >> (uint(int32(3)) % 32))
	goto L27
L33:
	;
	goto L24
L34:
	;
	v87 = F_sdsnew(m, int32(_a_F_evalExtractShebangFlags_4))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L18
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v87
	goto L22
L36:
	;
	if v160 < int32(2) {
		v578 = v160
		v590 = int64(0)
		goto L59
	} else {
		goto L60
	}
L37:
	;
	switch v62 {
	default:
		goto L43
	case 1:
		goto L42
	case 2:
		goto L41
	case 3:
		goto L40
	case 4:
		goto L39
	}
L38:
	;
	if base.Ui32(v106) <= base.Ui32(int32(1)) {
		goto L3
	} else {
		goto L44
	}
L39:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v57+int32(-17))))
	v106 = v105
	goto L38
L40:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v57+int32(-9))))
	v106 = v102
	goto L38
L41:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57+int32(-5)))))
	v106 = v99
	goto L38
L42:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+int32(-3)))))
	v106 = v96
	goto L38
L43:
	;
	v106 = int32(base.Ui32(v60) >> (uint(int32(3)) % 32))
	goto L38
L44:
	;
	v110 = v106 + int32(-1)
	v111 = F_valkey_calloc(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v111
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v116 = v114 + int32(2)
	if v110 == int32(0) {
		v141 = v116
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v160 = v159
	goto L36
L47:
	;
	goto L46
L48:
	;
	v146 = v141
	goto L55
L49:
	;
	v121 = v111
	v123 = v110
	v125 = v116
	goto L51
L50:
	;
	v136 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v136)
	v141 = v125
	goto L48
L51:
	;
	v127 = v123 + int32(-1)
	if v127 == int32(0) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v130)
	v132 = int32(1)
	if v130 != 0 {
		v121 = v121 + v132
		v123 = v127
		v125 = v125 + v132
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L47
L55:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if v148 != 0 {
		v146 = v146 + int32(1)
		goto L55
	} else {
		goto L57
	}
L56:
	;
	goto L47
L57:
	;
	goto L56
L58:
	;
	if l4 == int32(0) {
		v644 = v597
		goto L21
	} else {
		goto L164
	}
L59:
	;
	F_sdsfreesplitres(m, v50, v578)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L18
	} else {
		goto L163
	}
L60:
	;
	v166 = v50 + int32(4)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v168 = int32(_a_F_evalExtractShebangFlags_5)
	goto L62
L61:
	;
	if v202-v207 != 0 {
		v597 = v160
		v598 = v166
		goto L58
	} else {
		goto L74
	}
L62:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v173 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	goto L61
L65:
	;
	v175 = v167
	v176 = v168
	v177 = int32(6)
	v178 = v173
	goto L68
L66:
	;
	v202 = int32(0)
	v203 = v168
	goto L64
L67:
	;
	v202 = v199 & int32(255)
	v203 = v197
	goto L64
L68:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	if v178&int32(255) != v182 {
		v197 = v176
		v199 = v178
		goto L67
	} else {
		goto L70
	}
L69:
	;
	v197 = v191
	v199 = int32(0)
	goto L67
L70:
	;
	if v182 == int32(0) {
		v197 = v176
		v199 = v178
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v187 = v177 + int32(-1)
	if v187 == int32(0) {
		v197 = v176
		v199 = v178
		goto L67
	} else {
		goto L72
	}
L72:
	;
	v190 = int32(1)
	v191 = v176 + v190
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	if v192 != 0 {
		v175 = v175 + v190
		v176 = v191
		v177 = v187
		v178 = v192
		goto L68
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	v218 = v166
	v225 = v167
	v229 = int64(0)
	v230 = int32(1)
	goto L75
L75:
	;
	v235 = int32(-1)
	v243 = v225 + v235
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	v246 = v244 & int32(7)
	switch v246 {
	case 0:
		goto L84
	case 1:
		goto L83
	case 2:
		goto L82
	case 3:
		goto L81
	case 4:
		goto L80
	default:
		goto L78
	}
L77:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-1)))))
	switch v340 & int32(7) {
	case 0:
		goto L116
	case 1:
		goto L115
	case 2:
		goto L114
	case 3:
		goto L113
	case 4:
		goto L112
	default:
		v357 = int32(0)
		goto L111
	}
L78:
	;
	goto L77
L79:
	;
	if v261 == int32(0) {
		goto L78
	} else {
		goto L85
	}
L80:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v225+int32(-17))))
	v261 = v260
	goto L79
L81:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v225+int32(-9))))
	v261 = v257
	goto L79
L82:
	;
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225+int32(-5)))))
	v261 = v254
	goto L79
L83:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225+int32(-3)))))
	v261 = v251
	goto L79
L84:
	;
	v261 = int32(base.Ui32(v244) >> (uint(int32(3)) % 32))
	goto L79
L85:
	;
	v267 = int32(-1)&v261 + v235
	v271 = int32(0)&v261 + int32(6)
	v274 = v267 - v271 + int32(1)
	switch v246 {
	default:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	case 3:
		goto L88
	case 4:
		goto L87
	}
L86:
	;
	v290 = int32(0)
	v292 = base.B2i32(base.Ui32(v271) < base.Ui32(v289))
	if base.Ui32(v271) < base.Ui32(v289) {
		goto L93
	} else {
		goto L94
	}
L87:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v225+int32(-17))))
	v289 = v288
	goto L86
L88:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v225+int32(-9))))
	v289 = v285
	goto L86
L89:
	;
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225+int32(-5)))))
	v289 = v282
	goto L86
L90:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225+int32(-3)))))
	v289 = v279
	goto L86
L91:
	;
	v289 = int32(base.Ui32(v244) >> (uint(int32(3)) % 32))
	goto L86
L92:
	;
	v306 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v225+v300))) = uint8(v306)
	switch v246 {
	default:
		goto L110
	case 1:
		goto L109
	case 2:
		goto L108
	case 3:
		goto L107
	case 4:
		goto L106
	}
L93:
	;
	v293 = v271
	goto L95
L94:
	;
	v293 = v290
	goto L95
L95:
	;
	v294 = v289 - v293
	if base.Ui32(v274) < base.Ui32(v294) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v296 = v274
	goto L98
L97:
	;
	v296 = v294
	goto L98
L98:
	;
	if v267 < v271 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v298 = v290
	goto L101
L100:
	;
	v298 = v296
	goto L101
L101:
	;
	if base.Ui32(v271) < base.Ui32(v289) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v300 = v298
	goto L104
L103:
	;
	v300 = int32(0)
	goto L104
L104:
	;
	if v300 == int32(0) {
		goto L92
	} else {
		goto L105
	}
L105:
	;
	v304 = F_memmove(m, v225, v225+v293, v300)
	mBase = m.M
	goto L92
L106:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v225+int32(-17)))) = base.I64_extend_i32_u(v300)
	goto L78
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225+int32(-9)))) = v300
	goto L77
L108:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v225+int32(-5)))) = uint16(v300)
	goto L77
L109:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v225+int32(-3)))) = uint8(v300)
	goto L77
L110:
	;
	v309 = v300 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v243))) = uint8(v309)
	goto L77
L111:
	;
	v362 = F_sdssplitlen(m, v337, v357, int32(_a_F_evalExtractShebangFlags_6), int32(1), v20+int32(24))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L18
	} else {
		goto L117
	}
L112:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v337+int32(-17))))
	v357 = v356
	goto L111
L113:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v337+int32(-9))))
	v357 = v353
	goto L111
L114:
	;
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337+int32(-5)))))
	v357 = v350
	goto L111
L115:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-3)))))
	v357 = v347
	goto L111
L116:
	;
	v357 = int32(base.Ui32(v340) >> (uint(int32(3)) % 32))
	goto L111
L117:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v364 < int32(1) {
		v516 = v229
		goto L118
	} else {
		goto L119
	}
L118:
	;
	F_sdsfreesplitres(m, v362, v364)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L18
	} else {
		goto L147
	}
L119:
	;
	v367 = int32(0)
	v369 = *(*int32)(unsafe.Add(mBase, _c_F_evalExtractShebangFlags[0]))
	v381 = v367
	v383 = v229
	goto L120
L120:
	;
	if v369 == int32(0) {
		v738 = v367
		goto L2
	} else {
		goto L122
	}
L121:
	;
	v516 = v500
	goto L118
L122:
	;
	v390 = int32(_a_F_evalExtractShebangFlags_7)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v362+v381<<(uint(int32(2))%32))))
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_evalExtractShebangFlags[1]))
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396))))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	if v400 == int32(0) {
		v423 = v399
		v424 = v400
		goto L125
	} else {
		goto L126
	}
L123:
	;
	v500 = v383 | base.I64_extend_i32_u(v490)
	v502 = v381 + int32(1)
	if v502 != v364 {
		v381 = v502
		v383 = v500
		goto L120
	} else {
		goto L146
	}
L124:
	;
	if v424-v423&int32(255) == int32(0) {
		v490 = v369
		goto L123
	} else {
		goto L132
	}
L125:
	;
	goto L124
L126:
	;
	if v400 != v399&int32(255) {
		v423 = v399
		v424 = v400
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v406 = v394
	v407 = v396
	goto L128
L128:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+1)))
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+1)))
	if v411 == int32(0) {
		v423 = v410
		v424 = v411
		goto L125
	} else {
		goto L130
	}
L129:
	;
	v423 = v410
	v424 = v411
	goto L125
L130:
	;
	v414 = int32(1)
	if v411 == v410&int32(255) {
		v406 = v406 + v414
		v407 = v407 + v414
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v430 = v390
	goto L133
L133:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v430)+8))
	if v447 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v490 = v447
	goto L123
L135:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v430)+12))
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	if v454 == int32(0) {
		v477 = v453
		v478 = v454
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v738 = v381
	goto L2
L137:
	;
	if v478-v477&int32(255) != 0 {
		v430 = v430 + int32(8)
		goto L133
	} else {
		goto L145
	}
L138:
	;
	goto L137
L139:
	;
	if v454 != v453&int32(255) {
		v477 = v453
		v478 = v454
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v460 = v394
	v461 = v448
	goto L141
L141:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+1)))
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+1)))
	if v465 == int32(0) {
		v477 = v464
		v478 = v465
		goto L138
	} else {
		goto L143
	}
L142:
	;
	v477 = v464
	v478 = v465
	goto L138
L143:
	;
	v468 = int32(1)
	if v465 == v464&int32(255) {
		v460 = v460 + v468
		v461 = v461 + v468
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	goto L134
L146:
	;
	goto L121
L147:
	;
	v524 = v230 + int32(1)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if v525 <= v524 {
		v578 = v525
		v590 = v516
		goto L59
	} else {
		goto L148
	}
L148:
	;
	v529 = v50 + v524<<(uint(int32(2))%32)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	v531 = int32(_a_F_evalExtractShebangFlags_5)
	goto L150
L149:
	;
	if v565-v570 != 0 {
		v597 = v525
		v598 = v529
		goto L58
	} else {
		goto L162
	}
L150:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530))))
	if v536 != 0 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566))))
	goto L149
L153:
	;
	v538 = v530
	v539 = v531
	v540 = int32(6)
	v541 = v536
	goto L156
L154:
	;
	v565 = int32(0)
	v566 = v531
	goto L152
L155:
	;
	v565 = v562 & int32(255)
	v566 = v560
	goto L152
L156:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539))))
	if v541&int32(255) != v545 {
		v560 = v539
		v562 = v541
		goto L155
	} else {
		goto L158
	}
L157:
	;
	v560 = v554
	v562 = int32(0)
	goto L155
L158:
	;
	if v545 == int32(0) {
		v560 = v539
		v562 = v541
		goto L155
	} else {
		goto L159
	}
L159:
	;
	v550 = v540 + int32(-1)
	if v550 == int32(0) {
		v560 = v539
		v562 = v541
		goto L155
	} else {
		goto L160
	}
L160:
	;
	v553 = int32(1)
	v554 = v539 + v553
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538)+1)))
	if v555 != 0 {
		v538 = v538 + v553
		v539 = v554
		v540 = v550
		v541 = v555
		goto L156
	} else {
		goto L161
	}
L161:
	;
	goto L157
L162:
	;
	v218 = v529
	v225 = v530
	v229 = v516
	v230 = v524
	goto L75
L163:
	;
	v695 = v47
	v701 = v590
	goto L6
L164:
	;
	v616 = F_sdsempty(m)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L18
	} else {
		goto L165
	}
L165:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v598)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v618
	v623 = F_sdscatfmt(m, v616, int32(_a_F_evalExtractShebangFlags_8), v20+int32(16))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L18
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v623
	goto L22
L167:
	;
	v668 = v20
	goto L10
L168:
	;
	v686 = F_zstrdup(m, int32(_a_F_evalExtractShebangFlags_9))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L18
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v686
	v695 = v681
	v701 = v682
	goto L6
L170:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v701
	v762 = int32(0)
	v767 = v20
	goto L1
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v695
	goto L170
L172:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L174:
	;
	F_sdsfreesplitres(m, v362, v755)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L18
	} else {
		goto L178
	}
L175:
	;
	v742 = F_sdsempty(m)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L18
	} else {
		goto L176
	}
L176:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v362+v738<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v747
	v750 = F_sdscatfmt(m, v742, int32(_a_F_evalExtractShebangFlags_10), v20)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L18
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v750
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v755 = v753
	goto L174
L178:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	F_sdsfreesplitres(m, v50, v758)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L18
	} else {
		goto L179
	}
L179:
	;
	v762 = int32(-1)
	v767 = v20
	goto L1
}
func F_evalInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v4 = F_dictCreate(m, int32(_a_F_evalInit_0))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_evalInit[0])) = v4
		v8 = F_listCreate(m)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_evalInit[1])) = v8
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(3)
			*(*int64)(unsafe.Add(mBase, _c_F_evalInit[2])) = int64(0)
			return
		}
	}
}
func F_evalReset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	if l0 == int32(0) {
		v15 = int32(0)
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_evalReset[0]))
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_evalReset[1]))
		F_dictRelease(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_listRelease(m, v16)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_scriptingEngineManagerForEachEngine(m, int32(520), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v30 = F_dictCreate(m, int32(_a_F_evalReset_0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_evalReset[1])) = v30
						v34 = F_listCreate(m)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_evalReset[0])) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = int32(3)
							*(*int64)(unsafe.Add(mBase, _c_F_evalReset[2])) = int64(0)
							return
						}
					}
				}
			}
		}
	} else {
		v5 = F_listCreate(m)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			F_scriptingEngineManagerForEachEngine(m, int32(520), v5)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				v9 = int32(0)
				v10 = *(*int32)(unsafe.Add(mBase, _c_F_evalReset[1]))
				v12 = *(*int32)(unsafe.Add(mBase, _c_F_evalReset[0]))
				F_freeEvalScriptsAsync(m, v10, v12, v5)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v30 = F_dictCreate(m, int32(_a_F_evalReset_0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_evalReset[1])) = v30
						v34 = F_listCreate(m)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_evalReset[0])) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = int32(3)
							*(*int64)(unsafe.Add(mBase, _c_F_evalReset[2])) = int64(0)
							return
						}
					}
				}
			}
		}
	}
}
func F_evalScriptsMemory(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_evalScriptsMemory[0]))
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_evalScriptsMemory[1]))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+27)))
	if v9 == int32(255) {
		v13 = v1
	} else {
		v13 = int32(1) << (uint(v9) % 32)
	}
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+26)))
	if v16 == int32(255) {
		v20 = int32(0)
	} else {
		v20 = int32(1) << (uint(v16) % 32)
	}
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v31 = int32(0)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_evalScriptsMemory[1]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_evalScriptsMemory[2]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	return v3 + ((v13+v20)<<(uint(int32(2))%32) + (v24+v25)*int32(24)) + (v33+v34)<<(uint(int32(5))%32) + v41*int32(12)
}
func F_evalShaCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_evalShaCommand[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_replicationFeedMonitors(m, l0, v4, v6, v7, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v13 = F_objectGetVal(m, v12)
		mBase = m.M
		v14 = int32(-1)
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v14))))
		switch v16&int32(7) + v14 {
		case 0:
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
			v33 = v23
			if v33 == int32(40) {
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
				if v41&int32(16) != 0 {
					F_addReplyError(m, l0, int32(_a_F_evalShaCommand_0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						return
					}
				} else {
					F_evalGenericCommand(m, l0, int32(1))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_evalShaCommand[1]))
				F_addReplyErrorObject(m, l0, v38)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					return
				}
			}
		case 1:
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
			v33 = v26
			if v33 == int32(40) {
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
				if v41&int32(16) != 0 {
					F_addReplyError(m, l0, int32(_a_F_evalShaCommand_0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						return
					}
				} else {
					F_evalGenericCommand(m, l0, int32(1))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_evalShaCommand[1]))
				F_addReplyErrorObject(m, l0, v38)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					return
				}
			}
		case 2:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
			v33 = v29
			if v33 == int32(40) {
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
				if v41&int32(16) != 0 {
					F_addReplyError(m, l0, int32(_a_F_evalShaCommand_0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						return
					}
				} else {
					F_evalGenericCommand(m, l0, int32(1))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_evalShaCommand[1]))
				F_addReplyErrorObject(m, l0, v38)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					return
				}
			}
		case 3:
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
			v33 = v32
			if v33 == int32(40) {
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
				if v41&int32(16) != 0 {
					F_addReplyError(m, l0, int32(_a_F_evalShaCommand_0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						return
					}
				} else {
					F_evalGenericCommand(m, l0, int32(1))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_evalShaCommand[1]))
				F_addReplyErrorObject(m, l0, v38)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					return
				}
			}
		default:
			v38 = *(*int32)(unsafe.Add(mBase, _c_F_evalShaCommand[1]))
			F_addReplyErrorObject(m, l0, v38)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_evictClients(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	var v35 float64
	_ = v35
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
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
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_evictClients[0]))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+144))
	v18 = v10 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v19
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_evictClients[1]))
	if int32(-1) < v24 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v57 = int32(_a_F_evictClients_0)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_evictClients[2]))
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_evictClients[3]))
	v62 = int32(131072)
	if base.Ui32(v62) < base.Ui32(v54) {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	if v24 < int32(1) {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	v28 = *(*int64)(unsafe.Add(mBase, _c_F_evictClients[4]))
	if v28 == int64(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v35 = base.F64_div(base.F64_mul(base.F64_convert_i32_s(v24), base.F64_convert_i64_u(v28)), float64(-100))
	if base.F64_lt(v35, float64(1.8446744073709552e+19))&base.F64_ge(v35, float64(0)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if base.Ui64(v45) < base.Ui64(int64(4294967296)) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v45 = int64(0)
	goto L8
L10:
	;
	v43 = base.I64_trunc_f64_u(v35)
	v45 = v43
	goto L8
L11:
	;
	v50 = base.I32_wrap_i64(v45)
	goto L13
L12:
	;
	v50 = int32(-1)
	goto L13
L13:
	;
	v54 = v50
	goto L4
L14:
	;
	v54 = v24
	goto L4
L15:
	;
	v65 = v54
	goto L17
L16:
	;
	v65 = v62
	goto L17
L17:
	;
	if base.Ui32(v58+v60) <= base.Ui32(v65) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v73 = int32(18)
	goto L19
L19:
	;
	v76 = v10 + int32(8)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v78 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	goto L1
L21:
	;
	v147 = int32(_a_F_evictClients_0)
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_evictClients[2]))
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_evictClients[3]))
	if base.Ui32(v65) < base.Ui32(v148+v150) {
		v73 = v145
		goto L19
	} else {
		goto L42
	}
L22:
	;
	if int32(0) < v73 {
		goto L37
	} else {
		goto L38
	}
L23:
	;
	if v78 == int32(0) {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v78+base.B2i32(v81 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v87
	goto L24
L26:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	v92 = F_sdsempty(m)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_evictClients[5]))
	v96 = F_catClientInfoString(m, v92, v91, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_evictClients[6]))
	if int32(2) < v99 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v107 = F_freeClient(m, v91)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L27
	} else {
		goto L34
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v96
	F__serverLog(m, int32(2), int32(_a_F_evictClients_1), v10)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	F_sdsfree(m, v96)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L27
	} else {
		goto L36
	}
L34:
	;
	if v107 == int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v111 = int32(_a_F_evictClients_0)
	v113 = *(*int64)(unsafe.Add(mBase, _c_F_evictClients[7]))
	*(*int64)(unsafe.Add(mBase, _c_F_evictClients[7])) = v113 + int64(1)
	goto L33
L36:
	;
	v145 = v73
	goto L21
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_evictClients[0]))
	v133 = v73 + int32(-1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v131+v133<<(uint(int32(3))%32))))
	v139 = v10 + int32(8)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v140
	goto L41
L38:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_evictClients[6]))
	if int32(3) < v122 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F__serverLog(m, int32(3), int32(_a_F_evictClients_2), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L27
	} else {
		goto L40
	}
L40:
	;
	goto L1
L41:
	;
	v145 = v133
	goto L21
L42:
	;
	goto L20
}
func F_evictPolicyToString(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v42 int32
	_ = v42
	v1 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_evictPolicyToString[0]))
	if v6 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverPanic_1(m, int32(_a_F_evictPolicyToString_0), int32(362), int32(_a_F_evictPolicyToString_1), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_evictPolicyToString[1]))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_evictPolicyToString[2]))
	if v10 == v12 {
		v26 = v6
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v26
L4:
	;
	v17 = int32(_a_F_evictPolicyToString_2)
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v19 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v26 = v19
	goto L3
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v10 != v22 {
		v17 = v17 + int32(8)
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_evictionPoolAlloc(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	v4 = F_valkey_malloc(m, int32(384))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v6
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(0)
		v12 = F_sdsnewlen(m, v6, int32(255))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = v14
			*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v12
			v23 = F_sdsnewlen(m, v14, int32(255))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v25 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+56)) = v25
				*(*int64)(unsafe.Add(mBase, uint32(v4)+48)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+40)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v4)+36)) = v23
				v34 = F_sdsnewlen(m, v25, int32(255))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v4)+80)) = v36
					*(*int64)(unsafe.Add(mBase, uint32(v4)+72)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v4)+64)) = v36
					*(*int32)(unsafe.Add(mBase, uint32(v4)+60)) = v34
					v45 = F_sdsnewlen(m, v36, int32(255))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						v47 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v4)+104)) = v47
						*(*int64)(unsafe.Add(mBase, uint32(v4)+96)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v4)+88)) = v47
						*(*int32)(unsafe.Add(mBase, uint32(v4)+84)) = v45
						v56 = F_sdsnewlen(m, v47, int32(255))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							v58 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v4)+128)) = v58
							*(*int64)(unsafe.Add(mBase, uint32(v4)+120)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v4)+112)) = v58
							*(*int32)(unsafe.Add(mBase, uint32(v4)+108)) = v56
							v67 = F_sdsnewlen(m, v58, int32(255))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								v69 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v4)+152)) = v69
								*(*int64)(unsafe.Add(mBase, uint32(v4)+144)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v4)+136)) = v69
								*(*int32)(unsafe.Add(mBase, uint32(v4)+132)) = v67
								v78 = F_sdsnewlen(m, v69, int32(255))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									v80 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v4)+176)) = v80
									*(*int64)(unsafe.Add(mBase, uint32(v4)+168)) = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(v4)+160)) = v80
									*(*int32)(unsafe.Add(mBase, uint32(v4)+156)) = v78
									v89 = F_sdsnewlen(m, v80, int32(255))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										v91 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v4)+200)) = v91
										*(*int64)(unsafe.Add(mBase, uint32(v4)+192)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v4)+184)) = v91
										*(*int32)(unsafe.Add(mBase, uint32(v4)+180)) = v89
										v100 = F_sdsnewlen(m, v91, int32(255))
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											v102 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v4)+224)) = v102
											*(*int64)(unsafe.Add(mBase, uint32(v4)+216)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(v4)+208)) = v102
											*(*int32)(unsafe.Add(mBase, uint32(v4)+204)) = v100
											v111 = F_sdsnewlen(m, v102, int32(255))
											mBase = m.M
											v112 = m.ExcPending
											if v112 != 0 {
												return
											} else {
												v113 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v4)+248)) = v113
												*(*int64)(unsafe.Add(mBase, uint32(v4)+240)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(v4)+232)) = v113
												*(*int32)(unsafe.Add(mBase, uint32(v4)+228)) = v111
												v122 = F_sdsnewlen(m, v113, int32(255))
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													v124 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v4)+272)) = v124
													*(*int64)(unsafe.Add(mBase, uint32(v4)+264)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(v4)+256)) = v124
													*(*int32)(unsafe.Add(mBase, uint32(v4)+252)) = v122
													v133 = F_sdsnewlen(m, v124, int32(255))
													mBase = m.M
													v134 = m.ExcPending
													if v134 != 0 {
														return
													} else {
														v135 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v4)+296)) = v135
														*(*int64)(unsafe.Add(mBase, uint32(v4)+288)) = int64(0)
														*(*int32)(unsafe.Add(mBase, uint32(v4)+280)) = v135
														*(*int32)(unsafe.Add(mBase, uint32(v4)+276)) = v133
														v144 = F_sdsnewlen(m, v135, int32(255))
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
															return
														} else {
															v146 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v4)+320)) = v146
															*(*int64)(unsafe.Add(mBase, uint32(v4)+312)) = int64(0)
															*(*int32)(unsafe.Add(mBase, uint32(v4)+304)) = v146
															*(*int32)(unsafe.Add(mBase, uint32(v4)+300)) = v144
															v155 = F_sdsnewlen(m, v146, int32(255))
															mBase = m.M
															v156 = m.ExcPending
															if v156 != 0 {
																return
															} else {
																v157 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v4)+344)) = v157
																*(*int64)(unsafe.Add(mBase, uint32(v4)+336)) = int64(0)
																*(*int32)(unsafe.Add(mBase, uint32(v4)+328)) = v157
																*(*int32)(unsafe.Add(mBase, uint32(v4)+324)) = v155
																v166 = F_sdsnewlen(m, v157, int32(255))
																mBase = m.M
																v167 = m.ExcPending
																if v167 != 0 {
																	return
																} else {
																	v168 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v4)+368)) = v168
																	*(*int64)(unsafe.Add(mBase, uint32(v4)+360)) = int64(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v4)+352)) = v168
																	*(*int32)(unsafe.Add(mBase, uint32(v4)+348)) = v166
																	v177 = F_sdsnewlen(m, v168, int32(255))
																	mBase = m.M
																	v178 = m.ExcPending
																	if v178 != 0 {
																		return
																	} else {
																		v179 = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v4)+376)) = v179
																		*(*int32)(unsafe.Add(mBase, uint32(v4)+372)) = v177
																		*(*int32)(unsafe.Add(mBase, _c_F_evictionPoolAlloc[0])) = v4
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
					}
				}
			}
		}
	}
}
func F_existsCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int64
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v5 < int32(2) {
		v31 = int64(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_addReplyLongLong(m, l0, v31)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	v11 = int64(0)
	v12 = int32(1)
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14+v12<<(uint(int32(2))%32))))
	v20 = F_lookupKey(m, v13, v18, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v31 = v25
	goto L1
L5:
	;
	return
L6:
	;
	v25 = v11 + base.I64_extend_i32_u(base.B2i32(v20 != int32(0)))
	v27 = v12 + int32(1)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v27 < v28 {
		v11 = v25
		v12 = v27
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	return
}
func F_exitExecutionUnit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_exitExecutionUnit[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_exitExecutionUnit[0])) = v3 + int32(-1)
	return
}
func F_exitFromChild(m *base.Module, l0 int32) {
	F__Exit(m, l0)
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_exp(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v39 int64
	_ = v39
	var v55 float64
	_ = v55
	var v59 float64
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v70 float64
	_ = v70
	var v73 float64
	_ = v73
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v80 float64
	_ = v80
	var v83 float64
	_ = v83
	var v87 float64
	_ = v87
	var v90 float64
	_ = v90
	var v93 int64
	_ = v93
	var v98 int32
	_ = v98
	var v101 float64
	_ = v101
	var v104 float64
	_ = v104
	var v107 int64
	_ = v107
	var v110 int64
	_ = v110
	var v120 float64
	_ = v120
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v129 float64
	_ = v129
	var v134 float64
	_ = v134
	var v138 float64
	_ = v138
	var v139 float64
	_ = v139
	var v140 float64
	_ = v140
	var v149 float64
	_ = v149
	var v152 float64
	_ = v152
	var v153 float64
	_ = v153
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v169 float64
	_ = v169
	v13 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0))>>(uint(int64(52))%64))) & int32(2047)
	v18 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
	if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v18) <= base.Ui32(v13-v18) {
		if base.Ui32(v18) <= base.Ui32(v13) {
			if base.Ui32(v13) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
				v61 = int32(0)
				v62 = int32(0)
				v63 = *(*float64)(unsafe.Add(mBase, _c_F_exp[0]))
				v66 = *(*float64)(unsafe.Add(mBase, _c_F_exp[1]))
				v67 = base.F64_add(base.F64_mul(l0, v63), v66)
				v68 = base.F64_sub(v67, v66)
				v70 = *(*float64)(unsafe.Add(mBase, _c_F_exp[2]))
				v73 = *(*float64)(unsafe.Add(mBase, _c_F_exp[3]))
				v76 = base.F64_add(base.F64_mul(v68, v70), base.F64_add(base.F64_mul(v68, v73), l0))
				v77 = base.F64_mul(v76, v76)
				v80 = *(*float64)(unsafe.Add(mBase, _c_F_exp[4]))
				v83 = *(*float64)(unsafe.Add(mBase, _c_F_exp[5]))
				v87 = *(*float64)(unsafe.Add(mBase, _c_F_exp[6]))
				v90 = *(*float64)(unsafe.Add(mBase, _c_F_exp[7]))
				v93 = base.I64_reinterpret_f64(v67)
				v98 = base.I32_wrap_i64(v93) << (uint(int32(4)) % 32) & int32(2032)
				v101 = *(*float64)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_exp[8])))
				v104 = base.F64_add(base.F64_mul(base.F64_mul(v77, v77), base.F64_add(base.F64_mul(v76, v80), v83)), base.F64_add(base.F64_mul(v77, base.F64_add(base.F64_mul(v76, v87), v90)), base.F64_add(v101, v76)))
				v107 = *(*int64)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_exp[9])))
				v110 = v107 + v93<<(uint(int64(45))%64)
				if v61 != 0 {
					v163 = base.F64_reinterpret_i64(v110)
					v169 = base.F64_add(base.F64_mul(v163, v104), v163)
					return v169
				} else {
					if v93&int64(2147483648) != int64(0) {
						v127 = base.F64_reinterpret_i64(v110 + int64(4602678819172646912))
						v128 = base.F64_mul(v127, v104)
						v129 = base.F64_add(v128, v127)
						if base.F64_lt(v129, float64(1)) == int32(0) {
							v153 = v129
						} else {
							v134 = F_fp_barrier_2(m)
							mBase = m.M
							F_fp_force_eval_1(m, base.F64_mul(v134, float64(2.2250738585072014e-308)))
							mBase = m.M
							v138 = float64(0)
							v139 = float64(1)
							v140 = base.F64_add(v129, v139)
							v149 = base.F64_add(base.F64_add(v140, base.F64_add(base.F64_add(v128, base.F64_sub(v127, v129)), base.F64_add(v129, base.F64_sub(v139, v140)))), float64(-1))
							if base.F64_eq(v149, v138) != 0 {
								v152 = v138
							} else {
								v152 = v149
							}
							v153 = v152
						}
						v161 = base.F64_mul(v153, float64(2.2250738585072014e-308))
					} else {
						v120 = base.F64_reinterpret_i64(v110 + int64(-4544132024016830464))
						v161 = base.F64_mul(base.F64_add(base.F64_mul(v120, v104), v120), float64(5.486124068793689e+303))
					}
					return v161
				}
			} else {
				v39 = base.I64_reinterpret_f64(l0)
				if v39 == int64(-4503599627370496) {
					v169 = float64(0)
					return v169
				} else {
					if base.Ui32(v13) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(math.Float64frombits(uint64(0x7ff0000000000000))))>>(uint(int64(52))%64)))) {
						if int64(-1) < v39 {
							v59 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
							mBase = m.M
							return v59
						} else {
							v55 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
							mBase = m.M
							return v55
						}
					} else {
						return base.F64_add(l0, float64(1))
					}
				}
			}
		} else {
			return base.F64_add(l0, float64(1))
		}
	} else {
		v61 = v13
		v62 = int32(0)
		v63 = *(*float64)(unsafe.Add(mBase, _c_F_exp[0]))
		v66 = *(*float64)(unsafe.Add(mBase, _c_F_exp[1]))
		v67 = base.F64_add(base.F64_mul(l0, v63), v66)
		v68 = base.F64_sub(v67, v66)
		v70 = *(*float64)(unsafe.Add(mBase, _c_F_exp[2]))
		v73 = *(*float64)(unsafe.Add(mBase, _c_F_exp[3]))
		v76 = base.F64_add(base.F64_mul(v68, v70), base.F64_add(base.F64_mul(v68, v73), l0))
		v77 = base.F64_mul(v76, v76)
		v80 = *(*float64)(unsafe.Add(mBase, _c_F_exp[4]))
		v83 = *(*float64)(unsafe.Add(mBase, _c_F_exp[5]))
		v87 = *(*float64)(unsafe.Add(mBase, _c_F_exp[6]))
		v90 = *(*float64)(unsafe.Add(mBase, _c_F_exp[7]))
		v93 = base.I64_reinterpret_f64(v67)
		v98 = base.I32_wrap_i64(v93) << (uint(int32(4)) % 32) & int32(2032)
		v101 = *(*float64)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_exp[8])))
		v104 = base.F64_add(base.F64_mul(base.F64_mul(v77, v77), base.F64_add(base.F64_mul(v76, v80), v83)), base.F64_add(base.F64_mul(v77, base.F64_add(base.F64_mul(v76, v87), v90)), base.F64_add(v101, v76)))
		v107 = *(*int64)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_exp[9])))
		v110 = v107 + v93<<(uint(int64(45))%64)
		if v61 != 0 {
			v163 = base.F64_reinterpret_i64(v110)
			v169 = base.F64_add(base.F64_mul(v163, v104), v163)
			return v169
		} else {
			if v93&int64(2147483648) != int64(0) {
				v127 = base.F64_reinterpret_i64(v110 + int64(4602678819172646912))
				v128 = base.F64_mul(v127, v104)
				v129 = base.F64_add(v128, v127)
				if base.F64_lt(v129, float64(1)) == int32(0) {
					v153 = v129
				} else {
					v134 = F_fp_barrier_2(m)
					mBase = m.M
					F_fp_force_eval_1(m, base.F64_mul(v134, float64(2.2250738585072014e-308)))
					mBase = m.M
					v138 = float64(0)
					v139 = float64(1)
					v140 = base.F64_add(v129, v139)
					v149 = base.F64_add(base.F64_add(v140, base.F64_add(base.F64_add(v128, base.F64_sub(v127, v129)), base.F64_add(v129, base.F64_sub(v139, v140)))), float64(-1))
					if base.F64_eq(v149, v138) != 0 {
						v152 = v138
					} else {
						v152 = v149
					}
					v153 = v152
				}
				v161 = base.F64_mul(v153, float64(2.2250738585072014e-308))
			} else {
				v120 = base.F64_reinterpret_i64(v110 + int64(-4544132024016830464))
				v161 = base.F64_mul(base.F64_add(base.F64_mul(v120, v104), v120), float64(5.486124068793689e+303))
			}
			return v161
		}
	}
}
func F_expireatCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_expireGenericCommand(m, l0, int64(0), int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_extractDistanceOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(-1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v17 = F_getDoubleFromObjectOrReply(m, l0, v13, v10+int32(8), int32(_a_F_extractDistanceOrReply_0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 != 0 {
			v41 = v12
			m.G0 = v10 + int32(16)
			return v41
		} else {
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
			if base.F64_lt(v21, float64(0)) == int32(0) {
				if l3 == int32(0) {
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(l3))) = v21
				}
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v33 = F_extractUnitOrReply(m, l0, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					if base.F64_lt(v33, float64(0)) != 0 {
						v41 = v12
					} else {
						v37 = int32(0)
						if l2 == v37 {
							v41 = v37
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(l2))) = v33
							v41 = v37
						}
					}
					m.G0 = v10 + int32(16)
					return v41
				}
			} else {
				F_addReplyError(m, l0, int32(_a_F_extractDistanceOrReply_1))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v41 = v12
					m.G0 = v10 + int32(16)
					return v41
				}
			}
		}
	}
}
