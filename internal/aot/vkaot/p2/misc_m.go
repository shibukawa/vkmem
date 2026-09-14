package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_MurmurHash64A(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int64
	_ = v131
	var v136 int64
	_ = v136
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v175 int64
	_ = v175
	v16 = base.I64_extend_i32_s(l1)*int64(-4132994306676758123) ^ base.I64_extend_i32_u(l2)
	v18 = l1 & int32(-8)
	if v18 == int32(0) {
		v127 = l0
		v131 = v16
	} else {
		v22 = l1 + int32(-8)
		v23 = int32(24)
		if v22&v23 != v23 {
			v27 = int32(3)
			v35 = l0
			v36 = int32(0)
			v38 = v16
			for {
				v43 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
				v44 = int64(-4132994306676758123)
				v45 = v43 * v44
				v53 = ((int64(base.Ui64(v45)>>(uint(int64(47))%64))^v45)*v44 ^ v38) * v44
				v55 = v35 + int32(8)
				v57 = v36 + int32(1)
				if v57 != (int32(base.Ui32(v22)>>(uint(v27)%32))+int32(1))&v27 {
					v35 = v55
					v36 = v57
					v38 = v53
					continue
				} else {
					break
				}
				break
			}
			v60 = v55
			v63 = v53
		} else {
			v60 = l0
			v63 = v16
		}
		v68 = l0 + v18
		if base.Ui32(v22) < base.Ui32(int32(24)) {
			v127 = v68
			v131 = v63
		} else {
			v72 = v60
			v75 = v63
			for {
				v80 = *(*int64)(unsafe.Add(mBase, uint32(v72)+24))
				v81 = int64(-4132994306676758123)
				v82 = v80 * v81
				v83 = int64(47)
				v88 = *(*int64)(unsafe.Add(mBase, uint32(v72)+16))
				v90 = v88 * v81
				v96 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
				v98 = v96 * v81
				v104 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
				v106 = v104 * v81
				v123 = ((int64(base.Ui64(v82)>>(uint(v83)%64))^v82)*v81 ^ ((int64(base.Ui64(v90)>>(uint(v83)%64))^v90)*v81^((int64(base.Ui64(v98)>>(uint(v83)%64))^v98)*v81^((int64(base.Ui64(v106)>>(uint(v83)%64))^v106)*v81^v75)*v81)*v81)*v81) * v81
				v125 = v72 + int32(32)
				if v125 != v68 {
					v72 = v125
					v75 = v123
					continue
				} else {
					break
				}
				break
			}
			v127 = v68
			v131 = v123
		}
	}
	switch l1 & int32(7) {
	default:
		v170 = v131
	case 1:
		v165 = v131
		v166 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
		v170 = (v165 ^ v166) * int64(-4132994306676758123)
	case 2:
		v160 = v131
		v161 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
		v165 = v161<<(uint(int64(8))%64) ^ v160
		v166 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
		v170 = (v165 ^ v166) * int64(-4132994306676758123)
	case 3:
		v155 = v131
		v156 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
		v160 = v156<<(uint(int64(16))%64) ^ v155
		v161 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
		v165 = v161<<(uint(int64(8))%64) ^ v160
		v166 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
		v170 = (v165 ^ v166) * int64(-4132994306676758123)
	case 4:
		v150 = v131
		v151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
		v155 = v151<<(uint(int64(24))%64) ^ v150
		v156 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
		v160 = v156<<(uint(int64(16))%64) ^ v155
		v161 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
		v165 = v161<<(uint(int64(8))%64) ^ v160
		v166 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
		v170 = (v165 ^ v166) * int64(-4132994306676758123)
	case 5:
		v145 = v131
		v146 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
		v150 = v146<<(uint(int64(32))%64) ^ v145
		v151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
		v155 = v151<<(uint(int64(24))%64) ^ v150
		v156 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
		v160 = v156<<(uint(int64(16))%64) ^ v155
		v161 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
		v165 = v161<<(uint(int64(8))%64) ^ v160
		v166 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
		v170 = (v165 ^ v166) * int64(-4132994306676758123)
	case 6:
		v140 = v131
		v141 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
		v145 = v141<<(uint(int64(40))%64) ^ v140
		v146 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
		v150 = v146<<(uint(int64(32))%64) ^ v145
		v151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
		v155 = v151<<(uint(int64(24))%64) ^ v150
		v156 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
		v160 = v156<<(uint(int64(16))%64) ^ v155
		v161 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
		v165 = v161<<(uint(int64(8))%64) ^ v160
		v166 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
		v170 = (v165 ^ v166) * int64(-4132994306676758123)
	case 7:
		v136 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
		v140 = v136<<(uint(int64(48))%64) ^ v131
		v141 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
		v145 = v141<<(uint(int64(40))%64) ^ v140
		v146 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
		v150 = v146<<(uint(int64(32))%64) ^ v145
		v151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
		v155 = v151<<(uint(int64(24))%64) ^ v150
		v156 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
		v160 = v156<<(uint(int64(16))%64) ^ v155
		v161 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
		v165 = v161<<(uint(int64(8))%64) ^ v160
		v166 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
		v170 = (v165 ^ v166) * int64(-4132994306676758123)
	}
	v171 = int64(47)
	v175 = (int64(base.Ui64(v170)>>(uint(v171)%64)) ^ v170) * int64(-4132994306676758123)
	return int64(base.Ui64(v175)>>(uint(v171)%64)) ^ v175
}
func F___memcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	if l2 == int32(0) {
	} else {
		base.MemoryCopy(m, l0, l1, l2)
	}
	return l0
}
func F___memset(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	base.MemoryFill(m, l0, base.I32_extend8_s(l1), l2)
	return l0
}
func F___mkostemps(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v243 int32
	_ = v243
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0&int32(3) == int32(0) {
		v32 = l0
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v243
L2:
	;
	v243 = int32(-1)
	goto L1
L3:
	;
	v153 = int32(100)
	goto L41
L4:
	;
	goto L40
L5:
	;
	if base.Ui32(v65) < base.Ui32(int32(6)) {
		goto L4
	} else {
		goto L21
	}
L6:
	;
	v65 = v57 - l0
	goto L5
L7:
	;
	v36 = v32
	goto L15
L8:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v21 = l0
	goto L11
L10:
	;
	v65 = l0 - l0
	goto L5
L11:
	;
	v25 = v21 + int32(1)
	if v25&int32(3) == int32(0) {
		v32 = v25
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v30 != 0 {
		v21 = v25
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v57 = v25
	goto L6
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v45 = int32(-2139062144)
	if (int32(16843008)-v42|v42)&v45 == v45 {
		v36 = v36 + int32(4)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v51 = v36
	goto L18
L17:
	;
	goto L16
L18:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 != 0 {
		v51 = v51 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v57 = v51
	goto L6
L20:
	;
	goto L19
L21:
	;
	if base.Ui32(v65+int32(-6)) < base.Ui32(l1) {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v74 = l0 + v65 - l1 + int32(-6)
	v75 = int32(_a2179)
	v76 = int32(6)
	goto L27
L23:
	;
	if v140 == int32(0) {
		goto L3
	} else {
		goto L39
	}
L24:
	;
	v140 = int32(0)
	goto L23
L25:
	;
	v112 = v107
	v113 = v108
	v114 = v109
	goto L35
L26:
	;
	if v97 == int32(0) {
		goto L24
	} else {
		goto L33
	}
L27:
	;
	if (v75|v74)&int32(3) != 0 {
		v107 = v74
		v108 = v75
		v109 = v76
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v84 = v74
	v85 = v75
	v86 = v76
	goto L29
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v89 != v90 {
		v107 = v84
		v108 = v85
		v109 = v86
		goto L25
	} else {
		goto L31
	}
L30:
	;
	goto L26
L31:
	;
	v92 = int32(4)
	v93 = v85 + v92
	v95 = v84 + v92
	v97 = v86 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v97) {
		v84 = v95
		v85 = v93
		v86 = v97
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v107 = v95
	v108 = v93
	v109 = v97
	goto L25
L34:
	;
	v140 = v117 - v118
	goto L23
L35:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v117 != v118 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v120 = int32(1)
	v125 = v114 + int32(-1)
	if v125 == int32(0) {
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v112 = v112 + v120
	v113 = v113 + v120
	v114 = v125
	goto L35
L39:
	;
	goto L4
L40:
	;
	*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
	goto L2
L41:
	;
	v158 = int32(0)
	v164 = m.G0
	v166 = v164 - int32(16)
	m.G0 = v166
	v170 = F___clock_gettime(m, v158, v166)
	mBase = m.M
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v166)+8))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	v173 = F___get_tp(m)
	mBase = m.M
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+24))
	v177 = *(*int32)(unsafe.Add(mBase, _consts[1074]))
	*(*int32)(unsafe.Add(mBase, _consts[1074])) = v177 + int32(1)
	v188 = v158
	v189 = v177 + (v171 + v172 + v174*int32(65537))
	goto L44
L42:
	;
	goto L54
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(384)
	v215 = F_open(m, l0, l2&int32(-2097348)|int32(194), v9)
	mBase = m.M
	if int32(-1) < v215 {
		v243 = v215
		goto L1
	} else {
		goto L47
	}
L44:
	;
	v196 = int32(1)
	v202 = v189&int32(15) | v189<<(uint(v196)%32)&int32(32) + int32(65)
	*(*uint8)(unsafe.Add(mBase, uint32(v74+v188))) = uint8(v202)
	v207 = v188 + v196
	if v207 != int32(6) {
		v188 = v207
		v189 = int32(base.Ui32(v189) >> (uint(int32(5)) % 32))
		goto L44
	} else {
		goto L46
	}
L45:
	;
	m.G0 = v166 + int32(16)
	goto L43
L46:
	;
	goto L45
L47:
	;
	v219 = v153 + int32(-1)
	if v219 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	goto L42
L49:
	;
	goto L50
L50:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v223 == int32(20) {
		v153 = v219
		goto L41
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	goto L2
L53:
	;
	goto L52
L54:
	;
	v230 = F__emscripten_memcpy_bulkmem(m, v74, int32(_a2179), int32(6))
	mBase = m.M
	goto L53
}
func F___mmap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64) int32 {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	if l5&int64(-17592186040321) == int64(0) {
		if base.Ui32(l1) < base.Ui32(int32(2147483647)) {
			if l3&int32(16) == int32(0) {
				v30 = int32(-48)
			} else {
				v30 = int32(-63)
			}
			v33 = F___syscall_mmap2(m, l0, l1, l2, l3, l4, int64(base.Ui64(l5)>>(uint(int64(12))%64)))
			mBase = m.M
			if l3&int32(32) != 0 {
				v37 = v30
			} else {
				v37 = int32(-63)
			}
			if v33 != int32(-63) {
				v40 = v33
			} else {
				v40 = v37
			}
			if l0 != 0 {
				v41 = v33
			} else {
				v41 = v40
			}
			if base.Ui32(v41) < base.Ui32(int32(-4095)) {
				v49 = v41
			} else {
				v44 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(0) - v41
				v49 = int32(-1)
			}
			return v49
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(48)
			return int32(-1)
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(28)
		return int32(-1)
	}
}
func F_manualFailoverCanStart(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+uint32(_consts[160])))
	if v9 != 0 {
		F__serverAssert(m, int32(_a329), int32(_a253), int32(5986))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_consts[181])))
		if v10 == int64(0) {
			v27 = v8
			*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[160]))) = int32(1)
			m.G0 = v5 + int32(16)
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_consts[181]))) = int64(0)
			v16 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(3) < v16 {
				v27 = v8
				*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[160]))) = int32(1)
				m.G0 = v5 + int32(16)
				return
			} else {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_consts[192])))
				*(*int64)(unsafe.Add(mBase, uint32(v5))) = v19
				F__serverLog(m, int32(3), int32(_a330), v5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, _consts[136]))
					v27 = v26
					*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[160]))) = int32(1)
					m.G0 = v5 + int32(16)
					return
				}
			}
		}
	}
}
func F_markRewrittenIncrAofAsHistory(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a105), int32(_a85), int32(511))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L22
	} else {
		goto L33
	}
L2:
	;
	F__serverAssert(m, int32(_a106), int32(_a85), int32(505))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L22
	} else {
		goto L32
	}
L3:
	;
	F__serverAssert(m, int32(_a104), int32(_a85), int32(491))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L22
	} else {
		goto L31
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v14 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	m.G0 = v9 + int32(16)
	return
L6:
	;
	v18 = v9 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v19
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	if v24 == int32(-1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v44 = v9 + int32(8)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v46 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v28 = v9 + int32(8)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v30 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v30+base.B2i32(v33 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v39
	goto L11
L13:
	;
	goto L8
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(1)
	goto L5
L15:
	;
	if v46 == int32(0) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46+base.B2i32(v49 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v55
	goto L16
L18:
	;
	v62 = v46
	goto L19
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	if v66 != int32(105) {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L14
L21:
	;
	v70 = F_valkey_calloc(m, int32(24))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v73 = F_sdsdup(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v73
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = int32(104)
	*(*int64)(unsafe.Add(mBase, uint32(v70)+8)) = v76
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = F_listAddNodeHead(m, v80, v70)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_listDelNode(m, v83, v62)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v87 = v9 + int32(8)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v89 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v89 != 0 {
		v62 = v89
		goto L19
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v89+base.B2i32(v92 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v98
	goto L28
L30:
	;
	goto L20
L31:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_match_bracket(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v484 int32
	_ = v484
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v14 == int32(94) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v33 = v28 & int32(255)
	if v33 == int32(45) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v28 = v25
	v29 = int32(0)
	v30 = l0 + int32(2)
	v31 = int32(1)
	goto L1
L3:
	;
	if v14 == int32(33) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = int32(1)
	v28 = v14
	v29 = v19
	v30 = l0 + v19
	v31 = int32(0)
	goto L1
L5:
	;
	m.G0 = v12 + int32(32)
	return v484
L6:
	;
	v484 = v29
	goto L5
L7:
	;
	v49 = int32(*(*int8)(unsafe.Add(mBase, uint32(v46+int32(-1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v49
	v57 = v46
	goto L13
L8:
	;
	if l1 == int32(45) {
		goto L6
	} else {
		goto L12
	}
L9:
	;
	if v33 != int32(93) {
		v46 = v30
		goto L7
	} else {
		goto L10
	}
L10:
	;
	if l1 == int32(93) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v46 = v30 + int32(1)
	goto L7
L12:
	;
	v46 = v30 + int32(1)
	goto L7
L13:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	switch v60 + int32(-91) {
	case 0:
		goto L20
	case 1:
		goto L19
	case 2:
		v484 = v31
		goto L5
	default:
		goto L21
	}
L15:
	;
	v57 = v464 + int32(1)
	goto L13
L16:
	;
	if v461 == l1 {
		goto L6
	} else {
		goto L145
	}
L17:
	;
	v341 = v12 + int32(28)
	if v57 != 0 {
		goto L118
	} else {
		goto L119
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v60
	v460 = v57
	v461 = v60
	goto L16
L19:
	;
	if base.I32_extend8_s(v60) < int32(0) {
		goto L17
	} else {
		goto L116
	}
L20:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	switch v195 + int32(-58) {
	case 0, 3:
		goto L56
	case 1, 2:
		goto L18
	default:
		goto L57
	}
L21:
	;
	if v60 != int32(45) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	if v65 == int32(93) {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v70 = v57 + int32(1)
	if v70 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v180 < int32(0) {
		v484 = int32(0)
		goto L5
	} else {
		goto L51
	}
L25:
	;
	goto L29
L26:
	;
	v180 = int32(0)
	goto L24
L27:
	;
	v180 = v172
	goto L24
L28:
	;
	v168 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = int32(25)
	v172 = int32(-1)
	goto L27
L29:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v78 = base.I32_extend8_s(v77)
	if v78 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v86 = F___get_tp(m)
	mBase = m.M
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+96))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v88 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	if v12 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v180 = base.B2i32(v78 != int32(0))
	goto L24
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v77
	goto L32
L34:
	;
	v97 = v77 + int32(-194)
	if base.Ui32(int32(50)) < base.Ui32(v97) {
		goto L28
	} else {
		goto L37
	}
L35:
	;
	if v12 == int32(0) {
		v172 = int32(1)
		goto L27
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v78 & int32(57343)
	v180 = int32(1)
	goto L24
L37:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v97<<(uint(int32(2))%32))+uint32(_consts[1046])))
	goto L38
L38:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	v116 = int32(base.Ui32(v114) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(7)) < base.Ui32(v116+int32(-16)|(v116+v104>>(uint(int32(26))%32))) {
		goto L28
	} else {
		goto L41
	}
L41:
	;
	v129 = v114 + int32(-128) | v104<<(uint(int32(6))%32)
	if v129 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+2)))
	v139 = v137 + int32(-128)
	if base.Ui32(int32(63)) < base.Ui32(v139) {
		goto L28
	} else {
		goto L45
	}
L43:
	;
	if v12 == int32(0) {
		v172 = int32(2)
		goto L27
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v129
	v180 = int32(2)
	goto L24
L45:
	;
	v143 = v129 << (uint(int32(6)) % 32)
	v144 = v139 | v143
	if v143 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+3)))
	v154 = v152 + int32(-128)
	if base.Ui32(int32(63)) < base.Ui32(v154) {
		goto L28
	} else {
		goto L49
	}
L47:
	;
	if v12 == int32(0) {
		v172 = int32(3)
		goto L27
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v144
	v180 = int32(3)
	goto L24
L49:
	;
	if v12 == int32(0) {
		v172 = int32(4)
		goto L27
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v154 | v144<<(uint(int32(6))%32)
	v180 = int32(4)
	goto L24
L51:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if v183 < v184 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v464 = v57 + v180 + int32(-1)
	goto L15
L53:
	;
	v187 = v183 - v184
	if base.Ui32(l1-v184) <= base.Ui32(v187) {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	if base.Ui32(l2-v184) <= base.Ui32(v187) {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	v202 = v57 + int32(3)
	goto L60
L57:
	;
	if v195 != int32(46) {
		goto L18
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	if v195 != int32(58) {
		v464 = v202
		goto L15
	} else {
		goto L65
	}
L60:
	;
	v212 = v202 + int32(-1)
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	if v213 != v195 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v202 = v202 + int32(1)
	goto L60
L63:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v215 == int32(93) {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v223 = v57 + int32(2)
	v224 = v212 - v223
	if int32(15) < v224 {
		v464 = v202
		goto L15
	} else {
		goto L66
	}
L66:
	;
	if v224 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12+v224))) = uint8(v232)
	v241 = int32(_a2171)
	v242 = int32(97)
	v243 = int32(1)
	goto L72
L68:
	;
	goto L67
L69:
	;
	v229 = F__emscripten_memcpy_bulkmem(m, v12, v223, v224)
	mBase = m.M
	goto L68
L70:
	;
	switch v261 + int32(-1) {
	case 0:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	case 3:
		goto L88
	case 4:
		goto L87
	case 5:
		goto L86
	case 6:
		goto L85
	case 7:
		goto L84
	case 8:
		goto L83
	case 9:
		goto L82
	case 10:
		goto L81
	case 11:
		goto L80
	default:
		v281 = int32(0)
		goto L79
	}
L71:
	;
	goto L70
L72:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v244 != v242&int32(255) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v261 = int32(0)
	goto L71
L74:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+6)))
	v255 = v243 + int32(1)
	if v255 != int32(13) {
		v241 = v241 + int32(6)
		v242 = v251
		v243 = v255
		goto L72
	} else {
		goto L77
	}
L75:
	;
	v248 = F_strcmp(m, v12, v241)
	mBase = m.M
	if v248 == int32(0) {
		v261 = v243
		goto L71
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	goto L73
L78:
	;
	if v283 != 0 {
		goto L6
	} else {
		goto L92
	}
L79:
	;
	v283 = v281
	goto L78
L80:
	;
	v280 = F_iswxdigit(m, l1)
	mBase = m.M
	v281 = v280
	goto L79
L81:
	;
	v279 = F_iswupper(m, l1)
	mBase = m.M
	v283 = v279
	goto L78
L82:
	;
	v278 = F_iswspace(m, l1)
	mBase = m.M
	v283 = v278
	goto L78
L83:
	;
	v277 = F_iswpunct(m, l1)
	mBase = m.M
	v283 = v277
	goto L78
L84:
	;
	v276 = F_iswprint(m, l1)
	mBase = m.M
	v283 = v276
	goto L78
L85:
	;
	v275 = F_iswlower(m, l1)
	mBase = m.M
	v283 = v275
	goto L78
L86:
	;
	v274 = F_iswgraph(m, l1)
	mBase = m.M
	v283 = v274
	goto L78
L87:
	;
	v283 = base.B2i32(base.Ui32(l1+int32(-48)) < base.Ui32(int32(10)))
	goto L78
L88:
	;
	v269 = F_iswcntrl(m, l1)
	mBase = m.M
	v283 = v269
	goto L78
L89:
	;
	v268 = F_iswblank(m, l1)
	mBase = m.M
	v283 = v268
	goto L78
L90:
	;
	v267 = F_iswalpha(m, l1)
	mBase = m.M
	v283 = v267
	goto L78
L91:
	;
	v266 = F_iswalnum(m, l1)
	mBase = m.M
	v283 = v266
	goto L78
L92:
	;
	v291 = int32(_a2171)
	v292 = int32(97)
	v293 = int32(1)
	goto L95
L93:
	;
	switch v311 + int32(-1) {
	case 0:
		goto L114
	case 1:
		goto L113
	case 2:
		goto L112
	case 3:
		goto L111
	case 4:
		goto L110
	case 5:
		goto L109
	case 6:
		goto L108
	case 7:
		goto L107
	case 8:
		goto L106
	case 9:
		goto L105
	case 10:
		goto L104
	case 11:
		goto L103
	default:
		v331 = int32(0)
		goto L102
	}
L94:
	;
	goto L93
L95:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v294 != v292&int32(255) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v311 = int32(0)
	goto L94
L97:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+6)))
	v305 = v293 + int32(1)
	if v305 != int32(13) {
		v291 = v291 + int32(6)
		v292 = v301
		v293 = v305
		goto L95
	} else {
		goto L100
	}
L98:
	;
	v298 = F_strcmp(m, v12, v291)
	mBase = m.M
	if v298 == int32(0) {
		v311 = v293
		goto L94
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	goto L96
L101:
	;
	if v333 != 0 {
		goto L6
	} else {
		goto L115
	}
L102:
	;
	v333 = v331
	goto L101
L103:
	;
	v330 = F_iswxdigit(m, l2)
	mBase = m.M
	v331 = v330
	goto L102
L104:
	;
	v329 = F_iswupper(m, l2)
	mBase = m.M
	v333 = v329
	goto L101
L105:
	;
	v328 = F_iswspace(m, l2)
	mBase = m.M
	v333 = v328
	goto L101
L106:
	;
	v327 = F_iswpunct(m, l2)
	mBase = m.M
	v333 = v327
	goto L101
L107:
	;
	v326 = F_iswprint(m, l2)
	mBase = m.M
	v333 = v326
	goto L101
L108:
	;
	v325 = F_iswlower(m, l2)
	mBase = m.M
	v333 = v325
	goto L101
L109:
	;
	v324 = F_iswgraph(m, l2)
	mBase = m.M
	v333 = v324
	goto L101
L110:
	;
	v333 = base.B2i32(base.Ui32(l2+int32(-48)) < base.Ui32(int32(10)))
	goto L101
L111:
	;
	v319 = F_iswcntrl(m, l2)
	mBase = m.M
	v333 = v319
	goto L101
L112:
	;
	v318 = F_iswblank(m, l2)
	mBase = m.M
	v333 = v318
	goto L101
L113:
	;
	v317 = F_iswalpha(m, l2)
	mBase = m.M
	v333 = v317
	goto L101
L114:
	;
	v316 = F_iswalnum(m, l2)
	mBase = m.M
	v333 = v316
	goto L101
L115:
	;
	v464 = v202
	goto L15
L116:
	;
	goto L18
L117:
	;
	if v451 < int32(0) {
		v484 = int32(0)
		goto L5
	} else {
		goto L144
	}
L118:
	;
	goto L122
L119:
	;
	v451 = int32(0)
	goto L117
L120:
	;
	v451 = v443
	goto L117
L121:
	;
	v439 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v439))) = int32(25)
	v443 = int32(-1)
	goto L120
L122:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v349 = base.I32_extend8_s(v348)
	if v349 < int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v357 = F___get_tp(m)
	mBase = m.M
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)+96))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	if v359 != 0 {
		goto L127
	} else {
		goto L128
	}
L124:
	;
	if v341 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v451 = base.B2i32(v349 != int32(0))
	goto L117
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v341))) = v348
	goto L125
L127:
	;
	v368 = v348 + int32(-194)
	if base.Ui32(int32(50)) < base.Ui32(v368) {
		goto L121
	} else {
		goto L130
	}
L128:
	;
	if v341 == int32(0) {
		v443 = int32(1)
		goto L120
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v341))) = v349 & int32(57343)
	v451 = int32(1)
	goto L117
L130:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v368<<(uint(int32(2))%32))+uint32(_consts[1046])))
	goto L131
L131:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	v387 = int32(base.Ui32(v385) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(7)) < base.Ui32(v387+int32(-16)|(v387+v375>>(uint(int32(26))%32))) {
		goto L121
	} else {
		goto L134
	}
L134:
	;
	v400 = v385 + int32(-128) | v375<<(uint(int32(6))%32)
	if v400 < int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+2)))
	v410 = v408 + int32(-128)
	if base.Ui32(int32(63)) < base.Ui32(v410) {
		goto L121
	} else {
		goto L138
	}
L136:
	;
	if v341 == int32(0) {
		v443 = int32(2)
		goto L120
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v341))) = v400
	v451 = int32(2)
	goto L117
L138:
	;
	v414 = v400 << (uint(int32(6)) % 32)
	v415 = v410 | v414
	if v414 < int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+3)))
	v425 = v423 + int32(-128)
	if base.Ui32(int32(63)) < base.Ui32(v425) {
		goto L121
	} else {
		goto L142
	}
L140:
	;
	if v341 == int32(0) {
		v443 = int32(3)
		goto L120
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v341))) = v415
	v451 = int32(3)
	goto L117
L142:
	;
	if v341 == int32(0) {
		v443 = int32(4)
		goto L120
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v341))) = v425 | v415<<(uint(int32(6))%32)
	v451 = int32(4)
	goto L117
L144:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v460 = v57 + v451 + int32(-1)
	v461 = v457
	goto L16
L145:
	;
	if v461 == l2 {
		goto L6
	} else {
		goto L146
	}
L146:
	;
	v464 = v460
	goto L15
}
func F_mbsinit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	if l0 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		return base.B2i32(v4 == int32(0))
	} else {
		return int32(1)
	}
}
func F_membersOfAllNeighbors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v153 int32
	_ = v153
	var v156 int64
	_ = v156
	var v158 int64
	_ = v158
	var v162 int64
	_ = v162
	var v164 int64
	_ = v164
	var v173 int64
	_ = v173
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v183 int64
	_ = v183
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(192)
	m.G0 = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v18
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v22
	v26 = int32(72)
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l1+v26)))
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(56)))) = v28
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v30
	v34 = int32(120)
	v36 = *(*int64)(unsafe.Add(mBase, uint32(l1+v34)))
	*(*int64)(unsafe.Add(mBase, uint32(v16+v26))) = v36
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v38
	v40 = int32(88)
	v44 = *(*int64)(unsafe.Add(mBase, uint32(l1+v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v16+v40))) = v44
	v46 = *(*int64)(unsafe.Add(mBase, uint32(l1)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v46
	v215 = int32(104)
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l1+v215)))
	*(*int64)(unsafe.Add(mBase, uint32(v16+v215))) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v56
	v62 = int32(136)
	v64 = *(*int64)(unsafe.Add(mBase, uint32(l1+v62)))
	*(*int64)(unsafe.Add(mBase, uint32(v16+v34))) = v64
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l1)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+112)) = v66
	v72 = int32(168)
	v74 = *(*int64)(unsafe.Add(mBase, uint32(l1+v72)))
	*(*int64)(unsafe.Add(mBase, uint32(v16+v62))) = v74
	v76 = *(*int64)(unsafe.Add(mBase, uint32(l1)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = v76
	v218 = int32(152)
	v84 = *(*int64)(unsafe.Add(mBase, uint32(l1+v218)))
	*(*int64)(unsafe.Add(mBase, uint32(v16+v218))) = v84
	v86 = *(*int64)(unsafe.Add(mBase, uint32(l1)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+144)) = v86
	v94 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(184))))
	*(*int64)(unsafe.Add(mBase, uint32(v16+v72))) = v94
	v96 = *(*int64)(unsafe.Add(mBase, uint32(l1)+176))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+160)) = v96
	v110 = v6
	v111 = v6
	v112 = v6
	goto L2
L1:
	;
	m.G0 = v16 + int32(192)
	return v206
L2:
	;
	v120 = v16 + int32(32) + v110<<(uint(int32(4))%32)
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v120)))
	if v121 != int64(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v206 = v196
	goto L1
L4:
	;
	v202 = v110 + int32(1)
	if v202 != int32(9) {
		v110 = v202
		v111 = v196
		v112 = v197
		goto L2
	} else {
		goto L17
	}
L5:
	;
	if v112 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+8)))
	if v124 == int32(0) {
		v196 = v111
		v197 = v112
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(l4+int32(-1)) < base.Ui32(v140) {
		v206 = v111
		goto L1
	} else {
		goto L12
	}
L9:
	;
	v133 = v16 + int32(32) + v112<<(uint(int32(4))%32)
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v133)))
	if v121 != v134 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+8)))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+8)))
	if v136 == v137 {
		v196 = v111
		v197 = v112
		goto L4
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v145 = v16 + int32(184)
	v147 = v120 + int32(8)
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v147)))
	*(*int64)(unsafe.Add(mBase, uint32(v145))) = v148
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v120)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+176)) = v150
	v153 = v16 + int32(16)
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v147)))
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(24)))) = v156
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v120)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v153)))
	v164 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v153)+8)))
	goto L13
L13:
	;
	v173 = *(*int64)(unsafe.Add(mBase, uint32(v145)))
	*(*int64)(unsafe.Add(mBase, uint32(v16+int32(8)))) = v173
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v16)+176))
	v176 = int64(1)
	v177 = v175 + v176
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v16)+176)) = v177
	v183 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)))
	goto L14
L14:
	;
	v191 = F_geoGetPointsInRange(m, l0, base.F64_convert_i64_u(v162<<(uint((int64(52)-v164<<(uint(int64(1))%64))&int64(4294967294))%64)), base.F64_convert_i64_u(v177<<(uint((int64(52)-v183<<(uint(v176)%64))&int64(4294967294))%64)), l2, l3, l4)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	v196 = v191 + v111
	v197 = v110
	goto L4
L17:
	;
	goto L3
}
func F_memcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	if base.Ui32(l2) < base.Ui32(int32(4)) {
		v27 = l0
		v28 = l1
		v29 = l2
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(0)
L2:
	;
	v39 = v34
	v40 = v35
	v41 = v36
	goto L12
L3:
	;
	if v29 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L4:
	;
	if (l1|l0)&int32(3) != 0 {
		v34 = l0
		v35 = l1
		v36 = l2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v11 = l0
	v12 = l1
	v13 = l2
	goto L6
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v16 != v17 {
		v34 = v11
		v35 = v12
		v36 = v13
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v27 = v22
	v28 = v20
	v29 = v24
	goto L3
L8:
	;
	v19 = int32(4)
	v20 = v12 + v19
	v22 = v11 + v19
	v24 = v13 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v24) {
		v11 = v22
		v12 = v20
		v13 = v24
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v34 = v27
	v35 = v28
	v36 = v29
	goto L2
L11:
	;
	return v44 - v45
L12:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v44 != v45 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v47 = int32(1)
	v52 = v41 + int32(-1)
	if v52 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v39 = v39 + v47
	v40 = v40 + v47
	v41 = v52
	goto L12
}
func F_memoryCommand(m *base.Module, l0 int32) {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
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
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v130 int64
	_ = v130
	var v135 int64
	_ = v135
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v146 int64
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v399 int64
	_ = v399
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int64
	_ = v629
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int64
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int64
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int64
	_ = v647
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int64
	_ = v653
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int64
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int64
	_ = v665
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int64
	_ = v671
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int64
	_ = v677
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int64
	_ = v683
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int64
	_ = v689
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int64
	_ = v695
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int64
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int64
	_ = v738
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int64
	_ = v744
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v761 int32
	_ = v761
	var v762 int64
	_ = v762
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v768 int64
	_ = v768
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int64
	_ = v774
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int64
	_ = v780
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int64
	_ = v786
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v792 int64
	_ = v792
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int64
	_ = v798
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 float32
	_ = v804
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v811 float32
	_ = v811
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v819 int64
	_ = v819
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v826 int64
	_ = v826
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v833 int64
	_ = v833
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v840 int64
	_ = v840
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 float32
	_ = v846
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v853 int64
	_ = v853
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v859 float32
	_ = v859
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v866 int64
	_ = v866
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 float32
	_ = v872
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v879 int64
	_ = v879
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 float32
	_ = v885
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v892 int64
	_ = v892
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1264 int32
	_ = v1264
	v10 = m.G0
	v12 = v10 - int32(64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v16&int32(4) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v80 = int32(_a621)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v83 != 0 {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v76 = v75
	goto L1
L3:
	;
	if v16&int32(1) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = int32(16)
	goto L6
L5:
	;
	v25 = int32(8)
	goto L6
L6:
	;
	v26 = v15 + v25
	if v16&int32(2) == int32(0) {
		v57 = v26
		goto L7
	} else {
		goto L8
	}
L7:
	;
	goto L17
L8:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v32 = v26 + v31
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	switch v36 & int32(7) {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	case 4:
		goto L10
	default:
		v53 = int32(0)
		goto L9
	}
L9:
	;
	v57 = v32 + int32(1) + v53 + int32(1)
	goto L7
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(-16))))
	v53 = v52
	goto L9
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(-8))))
	v53 = v49
	goto L9
L12:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+int32(-4)))))
	v53 = v46
	goto L9
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-2)))))
	v53 = v43
	goto L9
L14:
	;
	v53 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
	goto L9
L15:
	;
	v76 = v57 + v72
	goto L1
L16:
	;
	goto L15
L17:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[601]))
	goto L16
L18:
	;
	m.G0 = v12 + int32(64)
	return
L19:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v154&int32(4) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L20:
	;
	if v115-v117 != 0 {
		goto L19
	} else {
		goto L32
	}
L21:
	;
	v115 = F_tolower(m, v111)
	mBase = m.M
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v117 = F_tolower(m, v116)
	mBase = m.M
	goto L20
L22:
	;
	v85 = v76
	v86 = v80
	v87 = v83
	goto L25
L23:
	;
	v111 = int32(0)
	v112 = v80
	goto L21
L24:
	;
	v111 = v108 & int32(255)
	v112 = v107
	goto L21
L25:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v89 == int32(0) {
		v107 = v86
		v108 = v87
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v107 = v101
	v108 = int32(0)
	goto L24
L27:
	;
	v93 = v87 & int32(255)
	if v93 == v89 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v100 = int32(1)
	v101 = v86 + v100
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v102 != 0 {
		v85 = v85 + v100
		v86 = v101
		v87 = v102
		goto L25
	} else {
		goto L31
	}
L29:
	;
	v95 = F_tolower(m, v93)
	mBase = m.M
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v97 = F_tolower(m, v96)
	mBase = m.M
	if v95 == v97 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v107 = v86
	v108 = v99
	goto L24
L31:
	;
	goto L26
L32:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v119 != int32(2) {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	v124 = int32(0)
	v125 = *(*int64)(unsafe.Add(mBase, _consts[604]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(56)))) = v125
	v130 = *(*int64)(unsafe.Add(mBase, _consts[605]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(48)))) = v130
	v135 = *(*int64)(unsafe.Add(mBase, _consts[606]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(40)))) = v135
	v140 = *(*int64)(unsafe.Add(mBase, _consts[607]))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(32)))) = v140
	v143 = *(*int64)(unsafe.Add(mBase, _consts[608]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v143
	v146 = *(*int64)(unsafe.Add(mBase, _consts[609]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v146
	F_addReplyHelp(m, l0, v12+int32(16))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return
L35:
	;
	goto L18
L36:
	;
	v218 = int32(_a1064)
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	if v221 != 0 {
		goto L56
	} else {
		goto L57
	}
L37:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	v214 = v213
	goto L36
L38:
	;
	if v154&int32(1) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v163 = int32(16)
	goto L41
L40:
	;
	v163 = int32(8)
	goto L41
L41:
	;
	v164 = v153 + v163
	if v154&int32(2) == int32(0) {
		v195 = v164
		goto L42
	} else {
		goto L43
	}
L42:
	;
	goto L52
L43:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	v170 = v164 + v169
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	switch v174 & int32(7) {
	case 0:
		goto L49
	case 1:
		goto L48
	case 2:
		goto L47
	case 3:
		goto L46
	case 4:
		goto L45
	default:
		v191 = int32(0)
		goto L44
	}
L44:
	;
	v195 = v170 + int32(1) + v191 + int32(1)
	goto L42
L45:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v170+int32(-16))))
	v191 = v190
	goto L44
L46:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v170+int32(-8))))
	v191 = v187
	goto L44
L47:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170+int32(-4)))))
	v191 = v184
	goto L44
L48:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+int32(-2)))))
	v191 = v181
	goto L44
L49:
	;
	v191 = int32(base.Ui32(v174) >> (uint(int32(3)) % 32))
	goto L44
L50:
	;
	v214 = v195 + v210
	goto L36
L51:
	;
	goto L50
L52:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _consts[601]))
	goto L51
L53:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
	if v513&int32(4) == int32(0) {
		goto L140
	} else {
		goto L141
	}
L54:
	;
	if v253-v255 != 0 {
		goto L53
	} else {
		goto L66
	}
L55:
	;
	v253 = F_tolower(m, v249)
	mBase = m.M
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	v255 = F_tolower(m, v254)
	mBase = m.M
	goto L54
L56:
	;
	v223 = v214
	v224 = v218
	v225 = v221
	goto L59
L57:
	;
	v249 = int32(0)
	v250 = v218
	goto L55
L58:
	;
	v249 = v246 & int32(255)
	v250 = v245
	goto L55
L59:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	if v227 == int32(0) {
		v245 = v224
		v246 = v225
		goto L58
	} else {
		goto L61
	}
L60:
	;
	v245 = v239
	v246 = int32(0)
	goto L58
L61:
	;
	v231 = v225 & int32(255)
	if v231 == v227 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v238 = int32(1)
	v239 = v224 + v238
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
	if v240 != 0 {
		v223 = v223 + v238
		v224 = v239
		v225 = v240
		goto L59
	} else {
		goto L65
	}
L63:
	;
	v233 = F_tolower(m, v231)
	mBase = m.M
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	v235 = F_tolower(m, v234)
	mBase = m.M
	if v233 == v235 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	v245 = v224
	v246 = v237
	goto L58
L65:
	;
	goto L60
L66:
	;
	v257 = int32(3)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v258 < v257 {
		goto L53
	} else {
		goto L67
	}
L67:
	;
	if v258 == int32(3) {
		v424 = int32(5)
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+8))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if v432&int32(4) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L69:
	;
	v266 = v257
	goto L70
L70:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v273+v266<<(uint(int32(2))%32))))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	if v278&int32(4) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	if v399 == int64(0) {
		goto L113
	} else {
		goto L114
	}
L72:
	;
	v342 = int32(_a1065)
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	if v345 != 0 {
		goto L93
	} else {
		goto L94
	}
L73:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	v338 = v337
	goto L72
L74:
	;
	if v278&int32(1) != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v287 = int32(16)
	goto L77
L76:
	;
	v287 = int32(8)
	goto L77
L77:
	;
	v288 = v277 + v287
	if v278&int32(2) == int32(0) {
		v319 = v288
		goto L78
	} else {
		goto L79
	}
L78:
	;
	goto L88
L79:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	v294 = v288 + v293
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	switch v298 & int32(7) {
	case 0:
		goto L85
	case 1:
		goto L84
	case 2:
		goto L83
	case 3:
		goto L82
	case 4:
		goto L81
	default:
		v315 = int32(0)
		goto L80
	}
L80:
	;
	v319 = v294 + int32(1) + v315 + int32(1)
	goto L78
L81:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v294+int32(-16))))
	v315 = v314
	goto L80
L82:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v294+int32(-8))))
	v315 = v311
	goto L80
L83:
	;
	v308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v294+int32(-4)))))
	v315 = v308
	goto L80
L84:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+int32(-2)))))
	v315 = v305
	goto L80
L85:
	;
	v315 = int32(base.Ui32(v298) >> (uint(int32(3)) % 32))
	goto L80
L86:
	;
	v338 = v319 + v334
	goto L72
L87:
	;
	goto L86
L88:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _consts[601]))
	goto L87
L89:
	;
	v412 = v266 + int32(2)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v412 < v413 {
		v266 = v412
		goto L70
	} else {
		goto L112
	}
L90:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_addReplyErrorObject(m, l0, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L34
	} else {
		goto L111
	}
L91:
	;
	if v377-v379 != 0 {
		goto L90
	} else {
		goto L103
	}
L92:
	;
	v377 = F_tolower(m, v373)
	mBase = m.M
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374))))
	v379 = F_tolower(m, v378)
	mBase = m.M
	goto L91
L93:
	;
	v347 = v338
	v348 = v342
	v349 = v345
	goto L96
L94:
	;
	v373 = int32(0)
	v374 = v342
	goto L92
L95:
	;
	v373 = v370 & int32(255)
	v374 = v369
	goto L92
L96:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348))))
	if v351 == int32(0) {
		v369 = v348
		v370 = v349
		goto L95
	} else {
		goto L98
	}
L97:
	;
	v369 = v363
	v370 = int32(0)
	goto L95
L98:
	;
	v355 = v349 & int32(255)
	if v355 == v351 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v362 = int32(1)
	v363 = v348 + v362
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+1)))
	if v364 != 0 {
		v347 = v347 + v362
		v348 = v363
		v349 = v364
		goto L96
	} else {
		goto L102
	}
L100:
	;
	v357 = F_tolower(m, v355)
	mBase = m.M
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348))))
	v359 = F_tolower(m, v358)
	mBase = m.M
	if v357 == v359 {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347))))
	v369 = v348
	v370 = v361
	goto L95
L102:
	;
	goto L97
L103:
	;
	v382 = v266 + int32(1)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v383 <= v382 {
		goto L90
	} else {
		goto L104
	}
L104:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v385+v382<<(uint(int32(2))%32))))
	v392 = F_getLongLongFromObject(m, v389, v12+int32(16))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L34
	} else {
		goto L106
	}
L105:
	;
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	if int64(-1) < v399 {
		goto L89
	} else {
		goto L109
	}
L106:
	;
	if v392 == int32(0) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	F_addReplyError(m, l0, int32(_a1062))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L34
	} else {
		goto L108
	}
L108:
	;
	goto L18
L109:
	;
	v403 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_addReplyErrorObject(m, l0, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L34
	} else {
		goto L110
	}
L110:
	;
	goto L18
L111:
	;
	goto L18
L112:
	;
	goto L71
L113:
	;
	v419 = int32(-1)
	goto L115
L114:
	;
	v419 = base.I32_wrap_i64(v399)
	goto L115
L115:
	;
	v424 = v419
	goto L68
L116:
	;
	v496 = F_dbFind(m, v429, v492)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L34
	} else {
		goto L134
	}
L117:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v431)+8))
	v492 = v491
	goto L116
L118:
	;
	if v432&int32(1) != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v441 = int32(16)
	goto L121
L120:
	;
	v441 = int32(8)
	goto L121
L121:
	;
	v442 = v431 + v441
	if v432&int32(2) == int32(0) {
		v473 = v442
		goto L122
	} else {
		goto L123
	}
L122:
	;
	goto L132
L123:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
	v448 = v442 + v447
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
	switch v452 & int32(7) {
	case 0:
		goto L129
	case 1:
		goto L128
	case 2:
		goto L127
	case 3:
		goto L126
	case 4:
		goto L125
	default:
		v469 = int32(0)
		goto L124
	}
L124:
	;
	v473 = v448 + int32(1) + v469 + int32(1)
	goto L122
L125:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v448+int32(-16))))
	v469 = v468
	goto L124
L126:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v448+int32(-8))))
	v469 = v465
	goto L124
L127:
	;
	v462 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v448+int32(-4)))))
	v469 = v462
	goto L124
L128:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448+int32(-2)))))
	v469 = v459
	goto L124
L129:
	;
	v469 = int32(base.Ui32(v452) >> (uint(int32(3)) % 32))
	goto L124
L130:
	;
	v492 = v473 + v488
	goto L116
L131:
	;
	goto L130
L132:
	;
	v488 = *(*int32)(unsafe.Add(mBase, _consts[601]))
	goto L131
L133:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)+8))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)+28))
	v504 = F_objectComputeSize(m, v501, v496, v424, v503)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L34
	} else {
		goto L137
	}
L134:
	;
	if v496 != 0 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L34
	} else {
		goto L136
	}
L136:
	;
	goto L18
L137:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v504))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L34
	} else {
		goto L138
	}
L138:
	;
	goto L18
L139:
	;
	v577 = int32(_a795)
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573))))
	if v580 != 0 {
		goto L159
	} else {
		goto L160
	}
L140:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v512)+8))
	v573 = v572
	goto L139
L141:
	;
	if v513&int32(1) != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v522 = int32(16)
	goto L144
L143:
	;
	v522 = int32(8)
	goto L144
L144:
	;
	v523 = v512 + v522
	if v513&int32(2) == int32(0) {
		v554 = v523
		goto L145
	} else {
		goto L146
	}
L145:
	;
	goto L155
L146:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523))))
	v529 = v523 + v528
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529))))
	switch v533 & int32(7) {
	case 0:
		goto L152
	case 1:
		goto L151
	case 2:
		goto L150
	case 3:
		goto L149
	case 4:
		goto L148
	default:
		v550 = int32(0)
		goto L147
	}
L147:
	;
	v554 = v529 + int32(1) + v550 + int32(1)
	goto L145
L148:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v529+int32(-16))))
	v550 = v549
	goto L147
L149:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v529+int32(-8))))
	v550 = v546
	goto L147
L150:
	;
	v543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v529+int32(-4)))))
	v550 = v543
	goto L147
L151:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529+int32(-2)))))
	v550 = v540
	goto L147
L152:
	;
	v550 = int32(base.Ui32(v533) >> (uint(int32(3)) % 32))
	goto L147
L153:
	;
	v573 = v554 + v569
	goto L139
L154:
	;
	goto L153
L155:
	;
	v569 = *(*int32)(unsafe.Add(mBase, _consts[601]))
	goto L154
L156:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v900)+4))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v901)+4))
	if v902&int32(4) == int32(0) {
		goto L256
	} else {
		goto L257
	}
L157:
	;
	if v612-v614 != 0 {
		goto L156
	} else {
		goto L169
	}
L158:
	;
	v612 = F_tolower(m, v608)
	mBase = m.M
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
	v614 = F_tolower(m, v613)
	mBase = m.M
	goto L157
L159:
	;
	v582 = v573
	v583 = v577
	v584 = v580
	goto L162
L160:
	;
	v608 = int32(0)
	v609 = v577
	goto L158
L161:
	;
	v608 = v605 & int32(255)
	v609 = v604
	goto L158
L162:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583))))
	if v586 == int32(0) {
		v604 = v583
		v605 = v584
		goto L161
	} else {
		goto L164
	}
L163:
	;
	v604 = v598
	v605 = int32(0)
	goto L161
L164:
	;
	v590 = v584 & int32(255)
	if v590 == v586 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v597 = int32(1)
	v598 = v583 + v597
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+1)))
	if v599 != 0 {
		v582 = v582 + v597
		v583 = v598
		v584 = v599
		goto L162
	} else {
		goto L168
	}
L166:
	;
	v592 = F_tolower(m, v590)
	mBase = m.M
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583))))
	v594 = F_tolower(m, v593)
	mBase = m.M
	if v592 == v594 {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582))))
	v604 = v583
	v605 = v596
	goto L161
L168:
	;
	goto L163
L169:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v616 != int32(2) {
		goto L156
	} else {
		goto L170
	}
L170:
	;
	v619 = F_getMemoryOverheadData(m)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L34
	} else {
		goto L171
	}
L171:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v619)+108))
	F_addReplyMapLen(m, l0, v621+int32(34))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L34
	} else {
		goto L172
	}
L172:
	;
	F_addReplyBulkCString(m, l0, int32(_a1066))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L34
	} else {
		goto L173
	}
L173:
	;
	v629 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619))))
	F_addReplyLongLong(m, l0, v629)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L34
	} else {
		goto L174
	}
L174:
	;
	F_addReplyBulkCString(m, l0, int32(_a1067))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L34
	} else {
		goto L175
	}
L175:
	;
	v635 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+4)))
	F_addReplyLongLong(m, l0, v635)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L34
	} else {
		goto L176
	}
L176:
	;
	F_addReplyBulkCString(m, l0, int32(_a1068))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L34
	} else {
		goto L177
	}
L177:
	;
	v641 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+8)))
	F_addReplyLongLong(m, l0, v641)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L34
	} else {
		goto L178
	}
L178:
	;
	F_addReplyBulkCString(m, l0, int32(_a1069))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L34
	} else {
		goto L179
	}
L179:
	;
	v647 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+12)))
	F_addReplyLongLong(m, l0, v647)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L34
	} else {
		goto L180
	}
L180:
	;
	F_addReplyBulkCString(m, l0, int32(_a1070))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L34
	} else {
		goto L181
	}
L181:
	;
	v653 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+16)))
	F_addReplyLongLong(m, l0, v653)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L34
	} else {
		goto L182
	}
L182:
	;
	F_addReplyBulkCString(m, l0, int32(_a1071))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L34
	} else {
		goto L183
	}
L183:
	;
	v659 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+20)))
	F_addReplyLongLong(m, l0, v659)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L34
	} else {
		goto L184
	}
L184:
	;
	F_addReplyBulkCString(m, l0, int32(_a1072))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L34
	} else {
		goto L185
	}
L185:
	;
	v665 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+24)))
	F_addReplyLongLong(m, l0, v665)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L34
	} else {
		goto L186
	}
L186:
	;
	F_addReplyBulkCString(m, l0, int32(_a1073))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L34
	} else {
		goto L187
	}
L187:
	;
	v671 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+28)))
	F_addReplyLongLong(m, l0, v671)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L34
	} else {
		goto L188
	}
L188:
	;
	F_addReplyBulkCString(m, l0, int32(_a1074))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L34
	} else {
		goto L189
	}
L189:
	;
	v677 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+32)))
	F_addReplyLongLong(m, l0, v677)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L34
	} else {
		goto L190
	}
L190:
	;
	F_addReplyBulkCString(m, l0, int32(_a1075))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L34
	} else {
		goto L191
	}
L191:
	;
	v683 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+36)))
	F_addReplyLongLong(m, l0, v683)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L34
	} else {
		goto L192
	}
L192:
	;
	F_addReplyBulkCString(m, l0, int32(_a1076))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L34
	} else {
		goto L193
	}
L193:
	;
	v689 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+40)))
	F_addReplyLongLong(m, l0, v689)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L34
	} else {
		goto L194
	}
L194:
	;
	F_addReplyBulkCString(m, l0, int32(_a1077))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L34
	} else {
		goto L195
	}
L195:
	;
	v695 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+44)))
	F_addReplyLongLong(m, l0, v695)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L34
	} else {
		goto L196
	}
L196:
	;
	F_addReplyBulkCString(m, l0, int32(_a1078))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L34
	} else {
		goto L197
	}
L197:
	;
	v701 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+48)))
	F_addReplyLongLong(m, l0, v701)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L34
	} else {
		goto L198
	}
L198:
	;
	if v621 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	F_addReplyBulkCString(m, l0, int32(_a1079))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L34
	} else {
		goto L211
	}
L200:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v619)+124))
	v710 = int32(0)
	goto L201
L201:
	;
	v719 = v706 + v710*int32(12)
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v720
	v726 = F_snprintf(m, v12+int32(16), int32(32), int32(_a1080), v12)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L34
	} else {
		goto L203
	}
L202:
	;
	goto L199
L203:
	;
	F_addReplyBulkCString(m, l0, v12+int32(16))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L34
	} else {
		goto L204
	}
L204:
	;
	F_addReplyMapLen(m, l0, int32(2))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L34
	} else {
		goto L205
	}
L205:
	;
	F_addReplyBulkCString(m, l0, int32(_a1081))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L34
	} else {
		goto L206
	}
L206:
	;
	v738 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v719)+4)))
	F_addReplyLongLong(m, l0, v738)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L34
	} else {
		goto L207
	}
L207:
	;
	F_addReplyBulkCString(m, l0, int32(_a1082))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L34
	} else {
		goto L208
	}
L208:
	;
	v744 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v719)+8)))
	F_addReplyLongLong(m, l0, v744)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L34
	} else {
		goto L209
	}
L209:
	;
	v748 = v710 + int32(1)
	if base.Ui32(v748) < base.Ui32(v621) {
		v710 = v748
		goto L201
	} else {
		goto L210
	}
L210:
	;
	goto L202
L211:
	;
	v762 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+112)))
	F_addReplyLongLong(m, l0, v762)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L34
	} else {
		goto L212
	}
L212:
	;
	F_addReplyBulkCString(m, l0, int32(_a1083))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L34
	} else {
		goto L213
	}
L213:
	;
	v768 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+116)))
	F_addReplyLongLong(m, l0, v768)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L34
	} else {
		goto L214
	}
L214:
	;
	F_addReplyBulkCString(m, l0, int32(_a1084))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L34
	} else {
		goto L215
	}
L215:
	;
	v774 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+52)))
	F_addReplyLongLong(m, l0, v774)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L34
	} else {
		goto L216
	}
L216:
	;
	F_addReplyBulkCString(m, l0, int32(_a1085))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L34
	} else {
		goto L217
	}
L217:
	;
	v780 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+120)))
	F_addReplyLongLong(m, l0, v780)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L34
	} else {
		goto L218
	}
L218:
	;
	F_addReplyBulkCString(m, l0, int32(_a1086))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L34
	} else {
		goto L219
	}
L219:
	;
	v786 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+60)))
	F_addReplyLongLong(m, l0, v786)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L34
	} else {
		goto L220
	}
L220:
	;
	F_addReplyBulkCString(m, l0, int32(_a1087))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L34
	} else {
		goto L221
	}
L221:
	;
	v792 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+64)))
	F_addReplyLongLong(m, l0, v792)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L34
	} else {
		goto L222
	}
L222:
	;
	F_addReplyBulkCString(m, l0, int32(_a1088))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L34
	} else {
		goto L223
	}
L223:
	;
	v798 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+56)))
	F_addReplyLongLong(m, l0, v798)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L34
	} else {
		goto L224
	}
L224:
	;
	F_addReplyBulkCString(m, l0, int32(_a1089))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L34
	} else {
		goto L225
	}
L225:
	;
	v804 = *(*float32)(unsafe.Add(mBase, uint32(v619)+68))
	F_addReplyDouble(m, l0, base.F64_promote_f32(v804))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L34
	} else {
		goto L226
	}
L226:
	;
	F_addReplyBulkCString(m, l0, int32(_a1090))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L34
	} else {
		goto L227
	}
L227:
	;
	v811 = *(*float32)(unsafe.Add(mBase, uint32(v619)+72))
	F_addReplyDouble(m, l0, base.F64_promote_f32(v811))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L34
	} else {
		goto L228
	}
L228:
	;
	F_addReplyBulkCString(m, l0, int32(_a1091))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L34
	} else {
		goto L229
	}
L229:
	;
	v819 = int64(*(*uint32)(unsafe.Add(mBase, _consts[610])))
	F_addReplyLongLong(m, l0, v819)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L34
	} else {
		goto L230
	}
L230:
	;
	F_addReplyBulkCString(m, l0, int32(_a1092))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L34
	} else {
		goto L231
	}
L231:
	;
	v826 = int64(*(*uint32)(unsafe.Add(mBase, _consts[611])))
	F_addReplyLongLong(m, l0, v826)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L34
	} else {
		goto L232
	}
L232:
	;
	F_addReplyBulkCString(m, l0, int32(_a1093))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L34
	} else {
		goto L233
	}
L233:
	;
	v833 = int64(*(*uint32)(unsafe.Add(mBase, _consts[612])))
	F_addReplyLongLong(m, l0, v833)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L34
	} else {
		goto L234
	}
L234:
	;
	F_addReplyBulkCString(m, l0, int32(_a1094))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L34
	} else {
		goto L235
	}
L235:
	;
	v840 = int64(*(*uint32)(unsafe.Add(mBase, _consts[613])))
	F_addReplyLongLong(m, l0, v840)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L34
	} else {
		goto L236
	}
L236:
	;
	F_addReplyBulkCString(m, l0, int32(_a1095))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L34
	} else {
		goto L237
	}
L237:
	;
	v846 = *(*float32)(unsafe.Add(mBase, uint32(v619)+84))
	F_addReplyDouble(m, l0, base.F64_promote_f32(v846))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L34
	} else {
		goto L238
	}
L238:
	;
	F_addReplyBulkCString(m, l0, int32(_a1096))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L34
	} else {
		goto L239
	}
L239:
	;
	v853 = int64(*(*int32)(unsafe.Add(mBase, uint32(v619)+88)))
	F_addReplyLongLong(m, l0, v853)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L34
	} else {
		goto L240
	}
L240:
	;
	F_addReplyBulkCString(m, l0, int32(_a1097))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L34
	} else {
		goto L241
	}
L241:
	;
	v859 = *(*float32)(unsafe.Add(mBase, uint32(v619)+92))
	F_addReplyDouble(m, l0, base.F64_promote_f32(v859))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L34
	} else {
		goto L242
	}
L242:
	;
	F_addReplyBulkCString(m, l0, int32(_a1098))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L34
	} else {
		goto L243
	}
L243:
	;
	v866 = int64(*(*int32)(unsafe.Add(mBase, uint32(v619)+96)))
	F_addReplyLongLong(m, l0, v866)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L34
	} else {
		goto L244
	}
L244:
	;
	F_addReplyBulkCString(m, l0, int32(_a1099))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L34
	} else {
		goto L245
	}
L245:
	;
	v872 = *(*float32)(unsafe.Add(mBase, uint32(v619)+100))
	F_addReplyDouble(m, l0, base.F64_promote_f32(v872))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L34
	} else {
		goto L246
	}
L246:
	;
	F_addReplyBulkCString(m, l0, int32(_a1100))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L34
	} else {
		goto L247
	}
L247:
	;
	v879 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v619)+104)))
	F_addReplyLongLong(m, l0, v879)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L34
	} else {
		goto L248
	}
L248:
	;
	F_addReplyBulkCString(m, l0, int32(_a1101))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L34
	} else {
		goto L249
	}
L249:
	;
	v885 = *(*float32)(unsafe.Add(mBase, uint32(v619)+76))
	F_addReplyDouble(m, l0, base.F64_promote_f32(v885))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L34
	} else {
		goto L250
	}
L250:
	;
	F_addReplyBulkCString(m, l0, int32(_a1102))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L34
	} else {
		goto L251
	}
L251:
	;
	v892 = int64(*(*int32)(unsafe.Add(mBase, uint32(v619)+80)))
	F_addReplyLongLong(m, l0, v892)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L34
	} else {
		goto L252
	}
L252:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v619)+124))
	F_valkey_free(m, v895)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L34
	} else {
		goto L253
	}
L253:
	;
	F_valkey_free(m, v619)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L34
	} else {
		goto L254
	}
L254:
	;
	goto L18
L255:
	;
	v966 = int32(_a1103)
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v962))))
	if v969 != 0 {
		goto L275
	} else {
		goto L276
	}
L256:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v901)+8))
	v962 = v961
	goto L255
L257:
	;
	if v902&int32(1) != 0 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v911 = int32(16)
	goto L260
L259:
	;
	v911 = int32(8)
	goto L260
L260:
	;
	v912 = v901 + v911
	if v902&int32(2) == int32(0) {
		v943 = v912
		goto L261
	} else {
		goto L262
	}
L261:
	;
	goto L271
L262:
	;
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v912))))
	v918 = v912 + v917
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v918))))
	switch v922 & int32(7) {
	case 0:
		goto L268
	case 1:
		goto L267
	case 2:
		goto L266
	case 3:
		goto L265
	case 4:
		goto L264
	default:
		v939 = int32(0)
		goto L263
	}
L263:
	;
	v943 = v918 + int32(1) + v939 + int32(1)
	goto L261
L264:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v918+int32(-16))))
	v939 = v938
	goto L263
L265:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v918+int32(-8))))
	v939 = v935
	goto L263
L266:
	;
	v932 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v918+int32(-4)))))
	v939 = v932
	goto L263
L267:
	;
	v929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v918+int32(-2)))))
	v939 = v929
	goto L263
L268:
	;
	v939 = int32(base.Ui32(v922) >> (uint(int32(3)) % 32))
	goto L263
L269:
	;
	v962 = v943 + v958
	goto L255
L270:
	;
	goto L269
L271:
	;
	v958 = *(*int32)(unsafe.Add(mBase, _consts[601]))
	goto L270
L272:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1011)+4))
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+4))
	if v1013&int32(4) == int32(0) {
		goto L289
	} else {
		goto L290
	}
L273:
	;
	if v1001-v1003 != 0 {
		goto L272
	} else {
		goto L285
	}
L274:
	;
	v1001 = F_tolower(m, v997)
	mBase = m.M
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998))))
	v1003 = F_tolower(m, v1002)
	mBase = m.M
	goto L273
L275:
	;
	v971 = v962
	v972 = v966
	v973 = v969
	goto L278
L276:
	;
	v997 = int32(0)
	v998 = v966
	goto L274
L277:
	;
	v997 = v994 & int32(255)
	v998 = v993
	goto L274
L278:
	;
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v972))))
	if v975 == int32(0) {
		v993 = v972
		v994 = v973
		goto L277
	} else {
		goto L280
	}
L279:
	;
	v993 = v987
	v994 = int32(0)
	goto L277
L280:
	;
	v979 = v973 & int32(255)
	if v979 == v975 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v986 = int32(1)
	v987 = v972 + v986
	v988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971)+1)))
	if v988 != 0 {
		v971 = v971 + v986
		v972 = v987
		v973 = v988
		goto L278
	} else {
		goto L284
	}
L282:
	;
	v981 = F_tolower(m, v979)
	mBase = m.M
	v982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v972))))
	v983 = F_tolower(m, v982)
	mBase = m.M
	if v981 == v983 {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971))))
	v993 = v972
	v994 = v985
	goto L277
L284:
	;
	goto L279
L285:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1005 != int32(2) {
		goto L272
	} else {
		goto L286
	}
L286:
	;
	F_addReplyBulkCString(m, l0, int32(_a1104))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L34
	} else {
		goto L287
	}
L287:
	;
	goto L18
L288:
	;
	v1077 = int32(_a882)
	v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073))))
	if v1080 != 0 {
		goto L308
	} else {
		goto L309
	}
L289:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+8))
	v1073 = v1072
	goto L288
L290:
	;
	if v1013&int32(1) != 0 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1022 = int32(16)
	goto L293
L292:
	;
	v1022 = int32(8)
	goto L293
L293:
	;
	v1023 = v1012 + v1022
	if v1013&int32(2) == int32(0) {
		v1054 = v1023
		goto L294
	} else {
		goto L295
	}
L294:
	;
	goto L304
L295:
	;
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1023))))
	v1029 = v1023 + v1028
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029))))
	switch v1033 & int32(7) {
	case 0:
		goto L301
	case 1:
		goto L300
	case 2:
		goto L299
	case 3:
		goto L298
	case 4:
		goto L297
	default:
		v1050 = int32(0)
		goto L296
	}
L296:
	;
	v1054 = v1029 + int32(1) + v1050 + int32(1)
	goto L294
L297:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1029+int32(-16))))
	v1050 = v1049
	goto L296
L298:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1029+int32(-8))))
	v1050 = v1046
	goto L296
L299:
	;
	v1043 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1029+int32(-4)))))
	v1050 = v1043
	goto L296
L300:
	;
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029+int32(-2)))))
	v1050 = v1040
	goto L296
L301:
	;
	v1050 = int32(base.Ui32(v1033) >> (uint(int32(3)) % 32))
	goto L296
L302:
	;
	v1073 = v1054 + v1069
	goto L288
L303:
	;
	goto L302
L304:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, _consts[601]))
	goto L303
L305:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1147)+4))
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+4))
	if v1149&int32(4) == int32(0) {
		goto L330
	} else {
		goto L331
	}
L306:
	;
	if v1112-v1114 != 0 {
		goto L305
	} else {
		goto L318
	}
L307:
	;
	v1112 = F_tolower(m, v1108)
	mBase = m.M
	v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1109))))
	v1114 = F_tolower(m, v1113)
	mBase = m.M
	goto L306
L308:
	;
	v1082 = v1073
	v1083 = v1077
	v1084 = v1080
	goto L311
L309:
	;
	v1108 = int32(0)
	v1109 = v1077
	goto L307
L310:
	;
	v1108 = v1105 & int32(255)
	v1109 = v1104
	goto L307
L311:
	;
	v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1083))))
	if v1086 == int32(0) {
		v1104 = v1083
		v1105 = v1084
		goto L310
	} else {
		goto L313
	}
L312:
	;
	v1104 = v1098
	v1105 = int32(0)
	goto L310
L313:
	;
	v1090 = v1084 & int32(255)
	if v1090 == v1086 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1097 = int32(1)
	v1098 = v1083 + v1097
	v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+1)))
	if v1099 != 0 {
		v1082 = v1082 + v1097
		v1083 = v1098
		v1084 = v1099
		goto L311
	} else {
		goto L317
	}
L315:
	;
	v1092 = F_tolower(m, v1090)
	mBase = m.M
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1083))))
	v1094 = F_tolower(m, v1093)
	mBase = m.M
	if v1092 == v1094 {
		goto L314
	} else {
		goto L316
	}
L316:
	;
	v1096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082))))
	v1104 = v1083
	v1105 = v1096
	goto L310
L317:
	;
	goto L312
L318:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1116 != int32(2) {
		goto L305
	} else {
		goto L319
	}
L319:
	;
	v1120 = F_getMemoryDoctorReport(m)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L34
	} else {
		goto L326
	}
L320:
	;
	F_addReplyVerbatim(m, l0, v1120, v1141, int32(_a701))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L34
	} else {
		goto L327
	}
L321:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1120+int32(-17))))
	v1141 = v1140
	goto L320
L322:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1120+int32(-9))))
	v1141 = v1137
	goto L320
L323:
	;
	v1134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1120+int32(-5)))))
	v1141 = v1134
	goto L320
L324:
	;
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1120+int32(-3)))))
	v1141 = v1131
	goto L320
L325:
	;
	v1141 = int32(base.Ui32(v1124) >> (uint(int32(3)) % 32))
	goto L320
L326:
	;
	v1124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1120+int32(-1)))))
	switch v1124 & int32(7) {
	case 0:
		goto L325
	case 1:
		goto L324
	case 2:
		goto L323
	case 3:
		goto L322
	case 4:
		goto L321
	default:
		v1141 = int32(0)
		goto L320
	}
L327:
	;
	F_sdsfree(m, v1120)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L34
	} else {
		goto L328
	}
L328:
	;
	goto L18
L329:
	;
	v1213 = int32(_a1105)
	v1216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209))))
	if v1216 != 0 {
		goto L349
	} else {
		goto L350
	}
L330:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+8))
	v1209 = v1208
	goto L329
L331:
	;
	if v1149&int32(1) != 0 {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1158 = int32(16)
	goto L334
L333:
	;
	v1158 = int32(8)
	goto L334
L334:
	;
	v1159 = v1148 + v1158
	if v1149&int32(2) == int32(0) {
		v1190 = v1159
		goto L335
	} else {
		goto L336
	}
L335:
	;
	goto L345
L336:
	;
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1159))))
	v1165 = v1159 + v1164
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1165))))
	switch v1169 & int32(7) {
	case 0:
		goto L342
	case 1:
		goto L341
	case 2:
		goto L340
	case 3:
		goto L339
	case 4:
		goto L338
	default:
		v1186 = int32(0)
		goto L337
	}
L337:
	;
	v1190 = v1165 + int32(1) + v1186 + int32(1)
	goto L335
L338:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v1165+int32(-16))))
	v1186 = v1185
	goto L337
L339:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1165+int32(-8))))
	v1186 = v1182
	goto L337
L340:
	;
	v1179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1165+int32(-4)))))
	v1186 = v1179
	goto L337
L341:
	;
	v1176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1165+int32(-2)))))
	v1186 = v1176
	goto L337
L342:
	;
	v1186 = int32(base.Ui32(v1169) >> (uint(int32(3)) % 32))
	goto L337
L343:
	;
	v1209 = v1190 + v1205
	goto L329
L344:
	;
	goto L343
L345:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, _consts[601]))
	goto L344
L346:
	;
	F_addReplySubcommandSyntaxError(m, l0)
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L34
	} else {
		goto L366
	}
L347:
	;
	if v1248-v1250 != 0 {
		goto L346
	} else {
		goto L359
	}
L348:
	;
	v1248 = F_tolower(m, v1244)
	mBase = m.M
	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1245))))
	v1250 = F_tolower(m, v1249)
	mBase = m.M
	goto L347
L349:
	;
	v1218 = v1209
	v1219 = v1213
	v1220 = v1216
	goto L352
L350:
	;
	v1244 = int32(0)
	v1245 = v1213
	goto L348
L351:
	;
	v1244 = v1241 & int32(255)
	v1245 = v1240
	goto L348
L352:
	;
	v1222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219))))
	if v1222 == int32(0) {
		v1240 = v1219
		v1241 = v1220
		goto L351
	} else {
		goto L354
	}
L353:
	;
	v1240 = v1234
	v1241 = int32(0)
	goto L351
L354:
	;
	v1226 = v1220 & int32(255)
	if v1226 == v1222 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v1233 = int32(1)
	v1234 = v1219 + v1233
	v1235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1218)+1)))
	if v1235 != 0 {
		v1218 = v1218 + v1233
		v1219 = v1234
		v1220 = v1235
		goto L352
	} else {
		goto L358
	}
L356:
	;
	v1228 = F_tolower(m, v1226)
	mBase = m.M
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219))))
	v1230 = F_tolower(m, v1229)
	mBase = m.M
	if v1228 == v1230 {
		goto L355
	} else {
		goto L357
	}
L357:
	;
	v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1218))))
	v1240 = v1219
	v1241 = v1232
	goto L351
L358:
	;
	goto L353
L359:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1252 != int32(2) {
		goto L346
	} else {
		goto L360
	}
L360:
	;
	goto L362
L362:
	;
	goto L363
L363:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	F_addReply(m, l0, v1257)
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L34
	} else {
		goto L364
	}
L364:
	;
	goto L18
L366:
	;
	goto L18
}
func F_migrateCloseSocket(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	v8 = F_sdsempty(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = F_objectGetVal(m, l0)
		mBase = m.M
		v11 = int32(0)
		v13 = F_objectGetVal(m, l0)
		mBase = m.M
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-1)))))
		switch v16 & int32(7) {
		case 0:
			v33 = int32(base.Ui32(v16) >> (uint(int32(3)) % 32))
		case 1:
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
			v33 = v23
		case 2:
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
			v33 = v26
		case 3:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
			v33 = v29
		case 4:
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
			v33 = v32
		default:
			v33 = v11
		}
		v34 = F_sdscatlen(m, v8, v10, v33)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			v38 = F_sdscatlen(m, v34, int32(_a215), int32(1))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				v40 = F_objectGetVal(m, l1)
				mBase = m.M
				v41 = F_objectGetVal(m, l1)
				mBase = m.M
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(-1)))))
				switch v44 & int32(7) {
				case 0:
					v61 = int32(base.Ui32(v44) >> (uint(int32(3)) % 32))
				case 1:
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(-3)))))
					v61 = v51
				case 2:
					v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41+int32(-5)))))
					v61 = v54
				case 3:
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-9))))
					v61 = v57
				case 4:
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-17))))
					v61 = v60
				default:
					v61 = v11
				}
				v62 = F_sdscatlen(m, v38, v40, v61)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, _consts[135]))
					v66 = F_dictFetchValue(m, v65, v62)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						if v66 == int32(0) {
							F_sdsfree(m, v62)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								return
							}
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+52))
							m.T0[v72].(func(*base.Module, int32))(m, v70)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								F_valkey_free(m, v66)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, _consts[135]))
									v79 = F_dictDelete(m, v78, v62)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										F_sdsfree(m, v62)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
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
func F_migrateGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
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
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v449 int32
	_ = v449
	v13 = int32(3)
	if int32(7) <= l2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v294 != 0 {
		v299 = v294
		goto L87
	} else {
		goto L88
	}
L2:
	;
	v24 = int32(6)
	goto L4
L3:
	;
	v286 = v13
	v287 = int32(1)
	goto L1
L4:
	;
	v32 = l1 + v24<<(uint(int32(2))%32)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = F_objectGetVal(m, v33)
	mBase = m.M
	v35 = int32(_a64)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v38 != 0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v286 = v13
	v287 = v278
	goto L1
L6:
	;
	v278 = int32(1)
	v280 = v276 + v278
	if v280 < l2 {
		v24 = v280
		goto L4
	} else {
		goto L85
	}
L7:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	v276 = v274 + v24
	goto L6
L8:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v148 = F_objectGetVal(m, v147)
	mBase = m.M
	v149 = int32(_a202)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v152 != 0 {
		goto L47
	} else {
		goto L48
	}
L9:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v119 = F_objectGetVal(m, v118)
	mBase = m.M
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119+int32(-1)))))
	switch v122 & int32(7) {
	case 0:
		goto L42
	case 1:
		goto L41
	case 2:
		goto L40
	case 3:
		goto L39
	case 4:
		goto L38
	default:
		goto L36
	}
L10:
	;
	if v70-v72 == int32(0) {
		goto L9
	} else {
		goto L22
	}
L11:
	;
	v70 = F_tolower(m, v66)
	mBase = m.M
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v72 = F_tolower(m, v71)
	mBase = m.M
	goto L10
L12:
	;
	v40 = v34
	v41 = v35
	v42 = v38
	goto L15
L13:
	;
	v66 = int32(0)
	v67 = v35
	goto L11
L14:
	;
	v66 = v63 & int32(255)
	v67 = v62
	goto L11
L15:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v44 == int32(0) {
		v62 = v41
		v63 = v42
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v62 = v56
	v63 = int32(0)
	goto L14
L17:
	;
	v48 = v42 & int32(255)
	if v48 == v44 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v55 = int32(1)
	v56 = v41 + v55
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v57 != 0 {
		v40 = v40 + v55
		v41 = v56
		v42 = v57
		goto L15
	} else {
		goto L21
	}
L19:
	;
	v50 = F_tolower(m, v48)
	mBase = m.M
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v52 = F_tolower(m, v51)
	mBase = m.M
	if v50 == v52 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v62 = v41
	v63 = v54
	goto L14
L21:
	;
	goto L16
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v77 = F_objectGetVal(m, v76)
	mBase = m.M
	v78 = int32(_a592)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v81 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	if v113-v115 != 0 {
		goto L8
	} else {
		goto L35
	}
L24:
	;
	v113 = F_tolower(m, v109)
	mBase = m.M
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v115 = F_tolower(m, v114)
	mBase = m.M
	goto L23
L25:
	;
	v83 = v77
	v84 = v78
	v85 = v81
	goto L28
L26:
	;
	v109 = int32(0)
	v110 = v78
	goto L24
L27:
	;
	v109 = v106 & int32(255)
	v110 = v105
	goto L24
L28:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v87 == int32(0) {
		v105 = v84
		v106 = v85
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v105 = v99
	v106 = int32(0)
	goto L27
L30:
	;
	v91 = v85 & int32(255)
	if v91 == v87 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v98 = int32(1)
	v99 = v84 + v98
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v100 != 0 {
		v83 = v83 + v98
		v84 = v99
		v85 = v100
		goto L28
	} else {
		goto L34
	}
L32:
	;
	v93 = F_tolower(m, v91)
	mBase = m.M
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	v95 = F_tolower(m, v94)
	mBase = m.M
	if v93 == v95 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v105 = v84
	v106 = v97
	goto L27
L34:
	;
	goto L29
L35:
	;
	v273 = int32(_a593)
	goto L7
L36:
	;
	v145 = v24 + int32(1)
	v286 = v145
	v287 = l2 - v145
	goto L1
L37:
	;
	if v139 == int32(0) {
		goto L36
	} else {
		goto L43
	}
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v119+int32(-17))))
	v139 = v138
	goto L37
L39:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v119+int32(-9))))
	v139 = v135
	goto L37
L40:
	;
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119+int32(-5)))))
	v139 = v132
	goto L37
L41:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119+int32(-3)))))
	v139 = v129
	goto L37
L42:
	;
	v139 = int32(base.Ui32(v122) >> (uint(int32(3)) % 32))
	goto L37
L43:
	;
	v286 = v13
	v287 = int32(0)
	goto L1
L44:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v190 = F_objectGetVal(m, v189)
	mBase = m.M
	v191 = int32(_a594)
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v194 != 0 {
		goto L61
	} else {
		goto L62
	}
L45:
	;
	if v184-v186 != 0 {
		goto L44
	} else {
		goto L57
	}
L46:
	;
	v184 = F_tolower(m, v180)
	mBase = m.M
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	v186 = F_tolower(m, v185)
	mBase = m.M
	goto L45
L47:
	;
	v154 = v148
	v155 = v149
	v156 = v152
	goto L50
L48:
	;
	v180 = int32(0)
	v181 = v149
	goto L46
L49:
	;
	v180 = v177 & int32(255)
	v181 = v176
	goto L46
L50:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	if v158 == int32(0) {
		v176 = v155
		v177 = v156
		goto L49
	} else {
		goto L52
	}
L51:
	;
	v176 = v170
	v177 = int32(0)
	goto L49
L52:
	;
	v162 = v156 & int32(255)
	if v162 == v158 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v169 = int32(1)
	v170 = v155 + v169
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
	if v171 != 0 {
		v154 = v154 + v169
		v155 = v170
		v156 = v171
		goto L50
	} else {
		goto L56
	}
L54:
	;
	v164 = F_tolower(m, v162)
	mBase = m.M
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	v166 = F_tolower(m, v165)
	mBase = m.M
	if v164 == v166 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	v176 = v155
	v177 = v168
	goto L49
L56:
	;
	goto L51
L57:
	;
	v273 = int32(_a595)
	goto L7
L58:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v232 = F_objectGetVal(m, v231)
	mBase = m.M
	v233 = int32(_a596)
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	if v236 != 0 {
		goto L74
	} else {
		goto L75
	}
L59:
	;
	if v226-v228 != 0 {
		goto L58
	} else {
		goto L71
	}
L60:
	;
	v226 = F_tolower(m, v222)
	mBase = m.M
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	v228 = F_tolower(m, v227)
	mBase = m.M
	goto L59
L61:
	;
	v196 = v190
	v197 = v191
	v198 = v194
	goto L64
L62:
	;
	v222 = int32(0)
	v223 = v191
	goto L60
L63:
	;
	v222 = v219 & int32(255)
	v223 = v218
	goto L60
L64:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if v200 == int32(0) {
		v218 = v197
		v219 = v198
		goto L63
	} else {
		goto L66
	}
L65:
	;
	v218 = v212
	v219 = int32(0)
	goto L63
L66:
	;
	v204 = v198 & int32(255)
	if v204 == v200 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v211 = int32(1)
	v212 = v197 + v211
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+1)))
	if v213 != 0 {
		v196 = v196 + v211
		v197 = v212
		v198 = v213
		goto L64
	} else {
		goto L70
	}
L68:
	;
	v206 = F_tolower(m, v204)
	mBase = m.M
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v208 = F_tolower(m, v207)
	mBase = m.M
	if v206 == v208 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v218 = v197
	v219 = v210
	goto L63
L70:
	;
	goto L65
L71:
	;
	v273 = int32(_a597)
	goto L7
L72:
	;
	if v268-v270 != 0 {
		v276 = v24
		goto L6
	} else {
		goto L84
	}
L73:
	;
	v268 = F_tolower(m, v264)
	mBase = m.M
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v270 = F_tolower(m, v269)
	mBase = m.M
	goto L72
L74:
	;
	v238 = v232
	v239 = v233
	v240 = v236
	goto L77
L75:
	;
	v264 = int32(0)
	v265 = v233
	goto L73
L76:
	;
	v264 = v261 & int32(255)
	v265 = v260
	goto L73
L77:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	if v242 == int32(0) {
		v260 = v239
		v261 = v240
		goto L76
	} else {
		goto L79
	}
L78:
	;
	v260 = v254
	v261 = int32(0)
	goto L76
L79:
	;
	v246 = v240 & int32(255)
	if v246 == v242 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v253 = int32(1)
	v254 = v239 + v253
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+1)))
	if v255 != 0 {
		v238 = v238 + v253
		v239 = v254
		v240 = v255
		goto L77
	} else {
		goto L83
	}
L81:
	;
	v248 = F_tolower(m, v246)
	mBase = m.M
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	v250 = F_tolower(m, v249)
	mBase = m.M
	if v248 == v250 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	v260 = v239
	v261 = v252
	goto L76
L83:
	;
	goto L78
L84:
	;
	v273 = int32(_a598)
	goto L7
L85:
	;
	goto L5
L86:
	;
	F__serverAssert(m, int32(_a585), int32(_a550), int32(2292))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L95
	} else {
		goto L113
	}
L87:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v287 <= v300 {
		v329 = v299
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v295 != 0 {
		goto L86
	} else {
		goto L89
	}
L89:
	;
	v297 = l3 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v297
	v299 = v297
	goto L87
L90:
	;
	if v287 < int32(1) {
		goto L102
	} else {
		goto L103
	}
L91:
	;
	v303 = v287 << (uint(int32(3)) % 32)
	v305 = l3 + int32(12)
	if v299 == v305 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v287
	v329 = v325
	goto L90
L93:
	;
	v312 = F_valkey_malloc(m, v303)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L95
	} else {
		goto L97
	}
L94:
	;
	v307 = F_valkey_realloc(m, v299, v303)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	return int32(0)
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v307
	v325 = v307
	goto L92
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v312
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v315 == int32(0) {
		v325 = v312
		goto L92
	} else {
		goto L98
	}
L98:
	;
	v319 = v315 << (uint(int32(3)) % 32)
	if v319 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v325 = v312
	goto L92
L100:
	;
	goto L99
L101:
	;
	v322 = F__emscripten_memcpy_bulkmem(m, v312, v305, v319)
	mBase = m.M
	goto L100
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v287
	return v287
L103:
	;
	v333 = v287 & int32(3)
	v334 = int32(0)
	if base.Ui32(v287) < base.Ui32(int32(4)) {
		v399 = v334
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if v333 == int32(0) {
		goto L102
	} else {
		goto L109
	}
L105:
	;
	v340 = int32(0)
	v343 = v340
	v348 = v340
	goto L106
L106:
	;
	v354 = int32(3)
	v356 = v329 + v348<<(uint(v354)%32)
	v357 = int32(146)
	*(*int32)(unsafe.Add(mBase, uint32(v356)+4)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v356))) = v348 + v286
	v362 = v348 | int32(1)
	v365 = v329 + v362<<(uint(v354)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v365)+4)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v362 + v286
	v371 = v348 | int32(2)
	v374 = v329 + v371<<(uint(v354)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+4)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v374))) = v371 + v286
	v380 = v348 | v354
	v383 = v329 + v380<<(uint(v354)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v383)+4)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v383))) = v380 + v286
	v388 = int32(4)
	v389 = v348 + v388
	v391 = v343 + v388
	if v391 != v287&int32(2147483644) {
		v343 = v391
		v348 = v389
		goto L106
	} else {
		goto L108
	}
L107:
	;
	v399 = v389
	goto L104
L108:
	;
	goto L107
L109:
	;
	v409 = v334
	v413 = v399
	goto L110
L110:
	;
	v421 = v329 + v413<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v421)+4)) = int32(146)
	*(*int32)(unsafe.Add(mBase, uint32(v421))) = v413 + v286
	v426 = int32(1)
	v429 = v409 + v426
	if v429 != v333 {
		v409 = v429
		v413 = v413 + v426
		goto L110
	} else {
		goto L112
	}
L111:
	;
	goto L102
L112:
	;
	goto L111
L113:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_mixDigest(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	v5 = m.G0
	v7 = v5 - int32(96)
	m.G0 = v7
	v10 = m.G0
	v11 = int32(112)
	v12 = v10 - v11
	m.G0 = v12
	v15 = v12 + int32(20)
	F_SHA1Init(m, v15)
	mBase = m.M
	F_SHA1Update(m, v15, l1, l2)
	mBase = m.M
	F_SHA1Final(m, v12, v15)
	mBase = m.M
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v25 = v23 ^ v24
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v25)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	v29 = v27 ^ v28
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v29)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+2)))
	v33 = v31 ^ v32
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v33)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+3)))
	v37 = v35 ^ v36
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v37)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)))
	v41 = v39 ^ v40
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v41)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)))
	v45 = v43 ^ v44
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v45)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+6)))
	v49 = v47 ^ v48
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v49)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)))
	v53 = v51 ^ v52
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)) = uint8(v53)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
	v57 = v55 ^ v56
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v57)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)))
	v61 = v59 ^ v60
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v61)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)))
	v65 = v63 ^ v64
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v65)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
	v69 = v67 ^ v68
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v69)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)))
	v73 = v71 ^ v72
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v73)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)))
	v77 = v75 ^ v76
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v77)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)))
	v81 = v79 ^ v80
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v81)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	v85 = v83 ^ v84
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v85)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)))
	v89 = v87 ^ v88
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+17)))
	v93 = v91 ^ v92
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v93)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+18)))
	v97 = v95 ^ v96
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v97)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+19)))
	v101 = v99 ^ v100
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v101)
	m.G0 = v12 + v11
	v107 = v7 + int32(4)
	F_SHA1Init(m, v107)
	mBase = m.M
	F_SHA1Update(m, v107, l0, int32(20))
	mBase = m.M
	F_SHA1Final(m, l0, v107)
	mBase = m.M
	m.G0 = v7 + int32(96)
	return
}
func F_modulesCollectInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
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
	var v56 int64
	_ = v56
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
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int64
	_ = v229
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v243 int64
	_ = v243
	var v245 int64
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[541]))
	v15 = F_dictGetIterator(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_dictReleaseIterator(m, v15)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L2
	} else {
		goto L76
	}
L2:
	;
	return int32(0)
L3:
	;
	v26 = v15 + int32(20)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v122 == int32(0) {
		v304 = l0
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v33 = v26
	v34 = v30
	goto L8
L6:
	;
	v30 = int32(1)
	goto L5
L7:
	;
	v30 = int32(0)
	goto L5
L8:
	;
	switch v34 {
	case 0:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v34 = int32(0)
	goto L8
L11:
	;
	goto L4
L12:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v114
	if v114 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v38 != int32(-1) {
		v77 = v38
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v78 = int32(1)
	v79 = v77 + v78
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v79
	v81 = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v85+int32(26)))))
	if v89 == int32(255) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v42 != 0 {
		v77 = int32(-1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v44 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	if v71 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v51 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43)+16)))
	v52 = int64(*(*int8)(unsafe.Add(mBase, uint32(v43)+27)))
	v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(v43)+8)))
	v54 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43)+12)))
	v55 = int64(*(*int8)(unsafe.Add(mBase, uint32(v43)+26)))
	v56 = int64(*(*int32)(unsafe.Add(mBase, uint32(v43)+4)))
	v57 = F_wangHash64(m, v56)
	mBase = m.M
	v59 = F_wangHash64(m, v55+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v54+v59)
	mBase = m.M
	v63 = F_wangHash64(m, v53+v61)
	mBase = m.M
	v65 = F_wangHash64(m, v52+v63)
	mBase = m.M
	v67 = F_wangHash64(m, v51+v65)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v70 = v69
	goto L17
L19:
	;
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+24)))
	v49 = v47 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+24)) = uint16(v49)
	v70 = v43
	goto L17
L20:
	;
	v77 = v71 + int32(-1)
	goto L14
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v77 = v74
	goto L14
L22:
	;
	v104 = int32(2)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v84+v102<<(uint(v104)%32)+int32(4))))
	v33 = v109 + v103<<(uint(v104)%32)
	v34 = int32(1)
	goto L8
L23:
	;
	v93 = v81
	goto L25
L24:
	;
	v93 = v78 << (uint(v89) % 32)
	goto L25
L25:
	;
	if v79 < v93 {
		v102 = v85
		v103 = v79
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v85 != 0 {
		v122 = v81
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	if v95 == int32(-1) {
		v122 = v81
		goto L11
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = int64(4294967296)
	v102 = int32(1)
	v103 = int32(0)
	goto L22
L29:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v118
	v122 = v114
	goto L11
L30:
	;
	v128 = l0
	v131 = l3
	v134 = v122
	goto L31
L31:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	goto L34
L32:
	;
	v304 = v193
	goto L1
L33:
	;
	v204 = v15 + int32(20)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v205 != 0 {
		goto L51
	} else {
		goto L52
	}
L34:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+56))
	if v137 == int32(0) {
		v193 = v128
		v194 = v131
		goto L33
	} else {
		goto L35
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v136
	m.T0[v137].(func(*base.Module, int32, int32))(m, v11+int32(8), l2)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v151 == int32(0) {
		v189 = v150
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v193 = v189
	v194 = v192
	goto L33
L38:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+int32(-1)))))
	switch v157 & int32(7) {
	case 0:
		goto L44
	case 1:
		goto L43
	case 2:
		goto L42
	case 3:
		goto L41
	case 4:
		goto L40
	default:
		v174 = int32(0)
		goto L39
	}
L39:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v174+int32(-1)))))
	if v178 != int32(44) {
		v185 = v150
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v150+int32(-17))))
	v174 = v173
	goto L39
L41:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v150+int32(-9))))
	v174 = v170
	goto L39
L42:
	;
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150+int32(-5)))))
	v174 = v167
	goto L39
L43:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+int32(-3)))))
	v174 = v164
	goto L39
L44:
	;
	v174 = int32(base.Ui32(v157) >> (uint(int32(3)) % 32))
	goto L39
L45:
	;
	v187 = F_sdscat(m, v185, int32(_a132))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L2
	} else {
		goto L48
	}
L46:
	;
	F_sdsIncrLen(m, v150, int32(-1))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v185 = v184
	goto L45
L48:
	;
	v189 = v187
	goto L37
L49:
	;
	if v300 != 0 {
		v128 = v193
		v131 = v194
		v134 = v300
		goto L31
	} else {
		goto L75
	}
L50:
	;
	v211 = v204
	v212 = v208
	goto L53
L51:
	;
	v208 = int32(1)
	goto L50
L52:
	;
	v208 = int32(0)
	goto L50
L53:
	;
	switch v212 {
	case 0:
		goto L58
	default:
		goto L57
	}
L55:
	;
	v212 = int32(0)
	goto L53
L56:
	;
	goto L49
L57:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v292
	if v292 == int32(0) {
		goto L55
	} else {
		goto L74
	}
L58:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v216 != int32(-1) {
		v255 = v216
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v256 = int32(1)
	v257 = v255 + v256
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v257
	v259 = int32(0)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262+v263+int32(26)))))
	if v267 == int32(255) {
		goto L68
	} else {
		goto L69
	}
L60:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v220 != 0 {
		v255 = int32(-1)
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v222 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+20))
	if v249 != int32(-1) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v229 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v221)+16)))
	v230 = int64(*(*int8)(unsafe.Add(mBase, uint32(v221)+27)))
	v231 = int64(*(*int32)(unsafe.Add(mBase, uint32(v221)+8)))
	v232 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v221)+12)))
	v233 = int64(*(*int8)(unsafe.Add(mBase, uint32(v221)+26)))
	v234 = int64(*(*int32)(unsafe.Add(mBase, uint32(v221)+4)))
	v235 = F_wangHash64(m, v234)
	mBase = m.M
	v237 = F_wangHash64(m, v233+v235)
	mBase = m.M
	v239 = F_wangHash64(m, v232+v237)
	mBase = m.M
	v241 = F_wangHash64(m, v231+v239)
	mBase = m.M
	v243 = F_wangHash64(m, v230+v241)
	mBase = m.M
	v245 = F_wangHash64(m, v229+v243)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v245
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v248 = v247
	goto L62
L64:
	;
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221)+24)))
	v227 = v225 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v221)+24)) = uint16(v227)
	v248 = v221
	goto L62
L65:
	;
	v255 = v249 + int32(-1)
	goto L59
L66:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v255 = v252
	goto L59
L67:
	;
	v282 = int32(2)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v262+v280<<(uint(v282)%32)+int32(4))))
	v211 = v287 + v281<<(uint(v282)%32)
	v212 = int32(1)
	goto L53
L68:
	;
	v271 = v259
	goto L70
L69:
	;
	v271 = v256 << (uint(v267) % 32)
	goto L70
L70:
	;
	if v257 < v271 {
		v280 = v263
		v281 = v257
		goto L67
	} else {
		goto L71
	}
L71:
	;
	if v263 != 0 {
		v300 = v259
		goto L56
	} else {
		goto L72
	}
L72:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v262)+20))
	if v273 == int32(-1) {
		v300 = v259
		goto L56
	} else {
		goto L73
	}
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = int64(4294967296)
	v280 = int32(1)
	v281 = int32(0)
	goto L67
L74:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v292)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v296
	v300 = v292
	goto L56
L75:
	;
	goto L32
L76:
	;
	m.G0 = v11 + int32(32)
	return v304
}
func F_monotonicInfoString(m *base.Module) int32 {
	return int32(_a973)
}
func F_moveCommand(m *base.Module, l0 int32) {
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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
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
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v21 = F_getIntFromObjectOrReply(m, l0, v17, v12+int32(8), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		if v21 != 0 {
			m.G0 = v12 + int32(16)
			return
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if v23 < int32(0) {
				F_addReplyError(m, l0, int32(_a580))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					m.G0 = v12 + int32(16)
					return
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, _consts[65]))
				if v23 < v27 {
					v32 = F_createDatabaseIfNeeded(m, v23)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v32
						if v15 < int32(0) {
							v43 = v32
							if v14 != v32 {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
								v52 = F_lookupKey(m, v43, v50, int32(8))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v52
									if v52 != 0 {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
										if v62&int32(1) == int32(0) {
											v73 = int64(-1)
										} else {
											v72 = *(*int64)(unsafe.Add(mBase, uint32(v52+(v62&int32(4)^int32(12)))))
											v73 = v72
										}
										v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
										v77 = F_lookupKey(m, v32, v75, int32(8))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											if v77 == int32(0) {
												F_incrRefCount(m, v52)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													v87 = int32(_a44)
													v88 = *(*int32)(unsafe.Add(mBase, _consts[133]))
													v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
													v91 = F_objectGetVal(m, v90)
													mBase = m.M
													v92 = int32(0)
													v94 = *(*int32)(unsafe.Add(mBase, _consts[139]))
													if v94 == v92 {
														v99 = v92
														v101 = F_dbGenericDeleteWithDictIndex(m, v14, v90, v88, int32(1), v99)
														mBase = m.M
														v102 = m.ExcPending
														if v102 != 0 {
															return
														} else {
															v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
															F_dbAddInternal(m, v32, v104, v12+int32(12), int32(0))
															mBase = m.M
															v109 = m.ExcPending
															if v109 != 0 {
																return
															} else {
																if v73 == int64(-1) {
																	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
																	F_touchWatchedKey(m, v14, v117)
																	mBase = m.M
																	v119 = m.ExcPending
																	if v119 != 0 {
																		return
																	} else {
																		F_trackingInvalidateKey(m, l0, v117, int32(1))
																		mBase = m.M
																		v122 = m.ExcPending
																		if v122 != 0 {
																			return
																		} else {
																			v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
																			F_touchWatchedKey(m, v32, v124)
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return
																			} else {
																				F_trackingInvalidateKey(m, l0, v124, int32(1))
																				mBase = m.M
																				v129 = m.ExcPending
																				if v129 != 0 {
																					return
																				} else {
																					v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																					v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																					v134 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																					F_notifyKeyspaceEvent(m, int32(4), int32(_a581), v133, v134)
																					mBase = m.M
																					v136 = m.ExcPending
																					if v136 != 0 {
																						return
																					} else {
																						v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																						v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
																						v141 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
																						F_notifyKeyspaceEvent(m, int32(4), int32(_a582), v140, v141)
																						mBase = m.M
																						v143 = m.ExcPending
																						if v143 != 0 {
																							return
																						} else {
																							v144 = int32(_a44)
																							v146 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																							*(*int64)(unsafe.Add(mBase, _consts[83])) = v146 + int64(1)
																							v151 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																							F_addReply(m, l0, v151)
																							mBase = m.M
																							v153 = m.ExcPending
																							if v153 != 0 {
																								return
																							} else {
																								m.G0 = v12 + int32(16)
																								return
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
																	v114 = F_setExpire(m, l0, v32, v113, v73)
																	mBase = m.M
																	v115 = m.ExcPending
																	if v115 != 0 {
																		return
																	} else {
																		v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
																		F_touchWatchedKey(m, v14, v117)
																		mBase = m.M
																		v119 = m.ExcPending
																		if v119 != 0 {
																			return
																		} else {
																			F_trackingInvalidateKey(m, l0, v117, int32(1))
																			mBase = m.M
																			v122 = m.ExcPending
																			if v122 != 0 {
																				return
																			} else {
																				v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
																				F_touchWatchedKey(m, v32, v124)
																				mBase = m.M
																				v126 = m.ExcPending
																				if v126 != 0 {
																					return
																				} else {
																					F_trackingInvalidateKey(m, l0, v124, int32(1))
																					mBase = m.M
																					v129 = m.ExcPending
																					if v129 != 0 {
																						return
																					} else {
																						v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																						v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																						v134 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																						F_notifyKeyspaceEvent(m, int32(4), int32(_a581), v133, v134)
																						mBase = m.M
																						v136 = m.ExcPending
																						if v136 != 0 {
																							return
																						} else {
																							v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																							v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
																							v141 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
																							F_notifyKeyspaceEvent(m, int32(4), int32(_a582), v140, v141)
																							mBase = m.M
																							v143 = m.ExcPending
																							if v143 != 0 {
																								return
																							} else {
																								v144 = int32(_a44)
																								v146 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																								*(*int64)(unsafe.Add(mBase, _consts[83])) = v146 + int64(1)
																								v151 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																								F_addReply(m, l0, v151)
																								mBase = m.M
																								v153 = m.ExcPending
																								if v153 != 0 {
																									return
																								} else {
																									m.G0 = v12 + int32(16)
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
													} else {
														v97 = F_getKeySlot(m, v91)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															v99 = v97
															v101 = F_dbGenericDeleteWithDictIndex(m, v14, v90, v88, int32(1), v99)
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return
															} else {
																v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
																F_dbAddInternal(m, v32, v104, v12+int32(12), int32(0))
																mBase = m.M
																v109 = m.ExcPending
																if v109 != 0 {
																	return
																} else {
																	if v73 == int64(-1) {
																		v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
																		F_touchWatchedKey(m, v14, v117)
																		mBase = m.M
																		v119 = m.ExcPending
																		if v119 != 0 {
																			return
																		} else {
																			F_trackingInvalidateKey(m, l0, v117, int32(1))
																			mBase = m.M
																			v122 = m.ExcPending
																			if v122 != 0 {
																				return
																			} else {
																				v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
																				F_touchWatchedKey(m, v32, v124)
																				mBase = m.M
																				v126 = m.ExcPending
																				if v126 != 0 {
																					return
																				} else {
																					F_trackingInvalidateKey(m, l0, v124, int32(1))
																					mBase = m.M
																					v129 = m.ExcPending
																					if v129 != 0 {
																						return
																					} else {
																						v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																						v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																						v134 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																						F_notifyKeyspaceEvent(m, int32(4), int32(_a581), v133, v134)
																						mBase = m.M
																						v136 = m.ExcPending
																						if v136 != 0 {
																							return
																						} else {
																							v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																							v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
																							v141 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
																							F_notifyKeyspaceEvent(m, int32(4), int32(_a582), v140, v141)
																							mBase = m.M
																							v143 = m.ExcPending
																							if v143 != 0 {
																								return
																							} else {
																								v144 = int32(_a44)
																								v146 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																								*(*int64)(unsafe.Add(mBase, _consts[83])) = v146 + int64(1)
																								v151 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																								F_addReply(m, l0, v151)
																								mBase = m.M
																								v153 = m.ExcPending
																								if v153 != 0 {
																									return
																								} else {
																									m.G0 = v12 + int32(16)
																									return
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
																		v114 = F_setExpire(m, l0, v32, v113, v73)
																		mBase = m.M
																		v115 = m.ExcPending
																		if v115 != 0 {
																			return
																		} else {
																			v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
																			F_touchWatchedKey(m, v14, v117)
																			mBase = m.M
																			v119 = m.ExcPending
																			if v119 != 0 {
																				return
																			} else {
																				F_trackingInvalidateKey(m, l0, v117, int32(1))
																				mBase = m.M
																				v122 = m.ExcPending
																				if v122 != 0 {
																					return
																				} else {
																					v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																					v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
																					F_touchWatchedKey(m, v32, v124)
																					mBase = m.M
																					v126 = m.ExcPending
																					if v126 != 0 {
																						return
																					} else {
																						F_trackingInvalidateKey(m, l0, v124, int32(1))
																						mBase = m.M
																						v129 = m.ExcPending
																						if v129 != 0 {
																							return
																						} else {
																							v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																							v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																							v134 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																							F_notifyKeyspaceEvent(m, int32(4), int32(_a581), v133, v134)
																							mBase = m.M
																							v136 = m.ExcPending
																							if v136 != 0 {
																								return
																							} else {
																								v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																								v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
																								v141 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
																								F_notifyKeyspaceEvent(m, int32(4), int32(_a582), v140, v141)
																								mBase = m.M
																								v143 = m.ExcPending
																								if v143 != 0 {
																									return
																								} else {
																									v144 = int32(_a44)
																									v146 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																									*(*int64)(unsafe.Add(mBase, _consts[83])) = v146 + int64(1)
																									v151 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																									F_addReply(m, l0, v151)
																									mBase = m.M
																									v153 = m.ExcPending
																									if v153 != 0 {
																										return
																									} else {
																										m.G0 = v12 + int32(16)
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
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, _consts[85]))
												F_addReply(m, l0, v82)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return
												} else {
													m.G0 = v12 + int32(16)
													return
												}
											}
										}
									} else {
										v56 = *(*int32)(unsafe.Add(mBase, _consts[85]))
										F_addReply(m, l0, v56)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
											return
										}
									}
								}
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, _consts[348]))
								F_addReplyErrorObject(m, l0, v46)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									m.G0 = v12 + int32(16)
									return
								}
							}
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, _consts[65]))
							if v38 <= v15 {
								v43 = v32
								if v14 != v32 {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
									v52 = F_lookupKey(m, v43, v50, int32(8))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v52
										if v52 != 0 {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
											if v62&int32(1) == int32(0) {
												v73 = int64(-1)
											} else {
												v72 = *(*int64)(unsafe.Add(mBase, uint32(v52+(v62&int32(4)^int32(12)))))
												v73 = v72
											}
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
											v77 = F_lookupKey(m, v32, v75, int32(8))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												if v77 == int32(0) {
													F_incrRefCount(m, v52)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														v87 = int32(_a44)
														v88 = *(*int32)(unsafe.Add(mBase, _consts[133]))
														v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
														v91 = F_objectGetVal(m, v90)
														mBase = m.M
														v92 = int32(0)
														v94 = *(*int32)(unsafe.Add(mBase, _consts[139]))
														if v94 == v92 {
															v99 = v92
															v101 = F_dbGenericDeleteWithDictIndex(m, v14, v90, v88, int32(1), v99)
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return
															} else {
																v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
																F_dbAddInternal(m, v32, v104, v12+int32(12), int32(0))
																mBase = m.M
																v109 = m.ExcPending
																if v109 != 0 {
																	return
																} else {
																	if v73 == int64(-1) {
																		v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
																		F_touchWatchedKey(m, v14, v117)
																		mBase = m.M
																		v119 = m.ExcPending
																		if v119 != 0 {
																			return
																		} else {
																			F_trackingInvalidateKey(m, l0, v117, int32(1))
																			mBase = m.M
																			v122 = m.ExcPending
																			if v122 != 0 {
																				return
																			} else {
																				v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
																				F_touchWatchedKey(m, v32, v124)
																				mBase = m.M
																				v126 = m.ExcPending
																				if v126 != 0 {
																					return
																				} else {
																					F_trackingInvalidateKey(m, l0, v124, int32(1))
																					mBase = m.M
																					v129 = m.ExcPending
																					if v129 != 0 {
																						return
																					} else {
																						v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																						v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																						v134 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																						F_notifyKeyspaceEvent(m, int32(4), int32(_a581), v133, v134)
																						mBase = m.M
																						v136 = m.ExcPending
																						if v136 != 0 {
																							return
																						} else {
																							v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																							v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
																							v141 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
																							F_notifyKeyspaceEvent(m, int32(4), int32(_a582), v140, v141)
																							mBase = m.M
																							v143 = m.ExcPending
																							if v143 != 0 {
																								return
																							} else {
																								v144 = int32(_a44)
																								v146 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																								*(*int64)(unsafe.Add(mBase, _consts[83])) = v146 + int64(1)
																								v151 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																								F_addReply(m, l0, v151)
																								mBase = m.M
																								v153 = m.ExcPending
																								if v153 != 0 {
																									return
																								} else {
																									m.G0 = v12 + int32(16)
																									return
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
																		v114 = F_setExpire(m, l0, v32, v113, v73)
																		mBase = m.M
																		v115 = m.ExcPending
																		if v115 != 0 {
																			return
																		} else {
																			v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
																			F_touchWatchedKey(m, v14, v117)
																			mBase = m.M
																			v119 = m.ExcPending
																			if v119 != 0 {
																				return
																			} else {
																				F_trackingInvalidateKey(m, l0, v117, int32(1))
																				mBase = m.M
																				v122 = m.ExcPending
																				if v122 != 0 {
																					return
																				} else {
																					v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																					v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
																					F_touchWatchedKey(m, v32, v124)
																					mBase = m.M
																					v126 = m.ExcPending
																					if v126 != 0 {
																						return
																					} else {
																						F_trackingInvalidateKey(m, l0, v124, int32(1))
																						mBase = m.M
																						v129 = m.ExcPending
																						if v129 != 0 {
																							return
																						} else {
																							v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																							v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																							v134 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																							F_notifyKeyspaceEvent(m, int32(4), int32(_a581), v133, v134)
																							mBase = m.M
																							v136 = m.ExcPending
																							if v136 != 0 {
																								return
																							} else {
																								v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																								v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
																								v141 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
																								F_notifyKeyspaceEvent(m, int32(4), int32(_a582), v140, v141)
																								mBase = m.M
																								v143 = m.ExcPending
																								if v143 != 0 {
																									return
																								} else {
																									v144 = int32(_a44)
																									v146 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																									*(*int64)(unsafe.Add(mBase, _consts[83])) = v146 + int64(1)
																									v151 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																									F_addReply(m, l0, v151)
																									mBase = m.M
																									v153 = m.ExcPending
																									if v153 != 0 {
																										return
																									} else {
																										m.G0 = v12 + int32(16)
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
														} else {
															v97 = F_getKeySlot(m, v91)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																v99 = v97
																v101 = F_dbGenericDeleteWithDictIndex(m, v14, v90, v88, int32(1), v99)
																mBase = m.M
																v102 = m.ExcPending
																if v102 != 0 {
																	return
																} else {
																	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
																	F_dbAddInternal(m, v32, v104, v12+int32(12), int32(0))
																	mBase = m.M
																	v109 = m.ExcPending
																	if v109 != 0 {
																		return
																	} else {
																		if v73 == int64(-1) {
																			v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
																			F_touchWatchedKey(m, v14, v117)
																			mBase = m.M
																			v119 = m.ExcPending
																			if v119 != 0 {
																				return
																			} else {
																				F_trackingInvalidateKey(m, l0, v117, int32(1))
																				mBase = m.M
																				v122 = m.ExcPending
																				if v122 != 0 {
																					return
																				} else {
																					v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																					v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
																					F_touchWatchedKey(m, v32, v124)
																					mBase = m.M
																					v126 = m.ExcPending
																					if v126 != 0 {
																						return
																					} else {
																						F_trackingInvalidateKey(m, l0, v124, int32(1))
																						mBase = m.M
																						v129 = m.ExcPending
																						if v129 != 0 {
																							return
																						} else {
																							v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																							v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																							v134 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																							F_notifyKeyspaceEvent(m, int32(4), int32(_a581), v133, v134)
																							mBase = m.M
																							v136 = m.ExcPending
																							if v136 != 0 {
																								return
																							} else {
																								v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																								v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
																								v141 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
																								F_notifyKeyspaceEvent(m, int32(4), int32(_a582), v140, v141)
																								mBase = m.M
																								v143 = m.ExcPending
																								if v143 != 0 {
																									return
																								} else {
																									v144 = int32(_a44)
																									v146 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																									*(*int64)(unsafe.Add(mBase, _consts[83])) = v146 + int64(1)
																									v151 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																									F_addReply(m, l0, v151)
																									mBase = m.M
																									v153 = m.ExcPending
																									if v153 != 0 {
																										return
																									} else {
																										m.G0 = v12 + int32(16)
																										return
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
																			v114 = F_setExpire(m, l0, v32, v113, v73)
																			mBase = m.M
																			v115 = m.ExcPending
																			if v115 != 0 {
																				return
																			} else {
																				v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
																				F_touchWatchedKey(m, v14, v117)
																				mBase = m.M
																				v119 = m.ExcPending
																				if v119 != 0 {
																					return
																				} else {
																					F_trackingInvalidateKey(m, l0, v117, int32(1))
																					mBase = m.M
																					v122 = m.ExcPending
																					if v122 != 0 {
																						return
																					} else {
																						v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																						v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
																						F_touchWatchedKey(m, v32, v124)
																						mBase = m.M
																						v126 = m.ExcPending
																						if v126 != 0 {
																							return
																						} else {
																							F_trackingInvalidateKey(m, l0, v124, int32(1))
																							mBase = m.M
																							v129 = m.ExcPending
																							if v129 != 0 {
																								return
																							} else {
																								v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																								v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																								v134 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																								F_notifyKeyspaceEvent(m, int32(4), int32(_a581), v133, v134)
																								mBase = m.M
																								v136 = m.ExcPending
																								if v136 != 0 {
																									return
																								} else {
																									v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																									v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
																									v141 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
																									F_notifyKeyspaceEvent(m, int32(4), int32(_a582), v140, v141)
																									mBase = m.M
																									v143 = m.ExcPending
																									if v143 != 0 {
																										return
																									} else {
																										v144 = int32(_a44)
																										v146 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																										*(*int64)(unsafe.Add(mBase, _consts[83])) = v146 + int64(1)
																										v151 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																										F_addReply(m, l0, v151)
																										mBase = m.M
																										v153 = m.ExcPending
																										if v153 != 0 {
																											return
																										} else {
																											m.G0 = v12 + int32(16)
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
												} else {
													v82 = *(*int32)(unsafe.Add(mBase, _consts[85]))
													F_addReply(m, l0, v82)
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return
													} else {
														m.G0 = v12 + int32(16)
														return
													}
												}
											}
										} else {
											v56 = *(*int32)(unsafe.Add(mBase, _consts[85]))
											F_addReply(m, l0, v56)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return
											} else {
												m.G0 = v12 + int32(16)
												return
											}
										}
									}
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, _consts[348]))
									F_addReplyErrorObject(m, l0, v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										m.G0 = v12 + int32(16)
										return
									}
								}
							} else {
								v40 = F_createDatabaseIfNeeded(m, v15)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v40
									v43 = v40
									if v14 != v32 {
										v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
										v52 = F_lookupKey(m, v43, v50, int32(8))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v52
											if v52 != 0 {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
												if v62&int32(1) == int32(0) {
													v73 = int64(-1)
												} else {
													v72 = *(*int64)(unsafe.Add(mBase, uint32(v52+(v62&int32(4)^int32(12)))))
													v73 = v72
												}
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
												v77 = F_lookupKey(m, v32, v75, int32(8))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													if v77 == int32(0) {
														F_incrRefCount(m, v52)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return
														} else {
															v87 = int32(_a44)
															v88 = *(*int32)(unsafe.Add(mBase, _consts[133]))
															v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
															v91 = F_objectGetVal(m, v90)
															mBase = m.M
															v92 = int32(0)
															v94 = *(*int32)(unsafe.Add(mBase, _consts[139]))
															if v94 == v92 {
																v99 = v92
																v101 = F_dbGenericDeleteWithDictIndex(m, v14, v90, v88, int32(1), v99)
																mBase = m.M
																v102 = m.ExcPending
																if v102 != 0 {
																	return
																} else {
																	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
																	F_dbAddInternal(m, v32, v104, v12+int32(12), int32(0))
																	mBase = m.M
																	v109 = m.ExcPending
																	if v109 != 0 {
																		return
																	} else {
																		if v73 == int64(-1) {
																			v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
																			F_touchWatchedKey(m, v14, v117)
																			mBase = m.M
																			v119 = m.ExcPending
																			if v119 != 0 {
																				return
																			} else {
																				F_trackingInvalidateKey(m, l0, v117, int32(1))
																				mBase = m.M
																				v122 = m.ExcPending
																				if v122 != 0 {
																					return
																				} else {
																					v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																					v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
																					F_touchWatchedKey(m, v32, v124)
																					mBase = m.M
																					v126 = m.ExcPending
																					if v126 != 0 {
																						return
																					} else {
																						F_trackingInvalidateKey(m, l0, v124, int32(1))
																						mBase = m.M
																						v129 = m.ExcPending
																						if v129 != 0 {
																							return
																						} else {
																							v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																							v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																							v134 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																							F_notifyKeyspaceEvent(m, int32(4), int32(_a581), v133, v134)
																							mBase = m.M
																							v136 = m.ExcPending
																							if v136 != 0 {
																								return
																							} else {
																								v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																								v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
																								v141 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
																								F_notifyKeyspaceEvent(m, int32(4), int32(_a582), v140, v141)
																								mBase = m.M
																								v143 = m.ExcPending
																								if v143 != 0 {
																									return
																								} else {
																									v144 = int32(_a44)
																									v146 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																									*(*int64)(unsafe.Add(mBase, _consts[83])) = v146 + int64(1)
																									v151 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																									F_addReply(m, l0, v151)
																									mBase = m.M
																									v153 = m.ExcPending
																									if v153 != 0 {
																										return
																									} else {
																										m.G0 = v12 + int32(16)
																										return
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
																			v114 = F_setExpire(m, l0, v32, v113, v73)
																			mBase = m.M
																			v115 = m.ExcPending
																			if v115 != 0 {
																				return
																			} else {
																				v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
																				F_touchWatchedKey(m, v14, v117)
																				mBase = m.M
																				v119 = m.ExcPending
																				if v119 != 0 {
																					return
																				} else {
																					F_trackingInvalidateKey(m, l0, v117, int32(1))
																					mBase = m.M
																					v122 = m.ExcPending
																					if v122 != 0 {
																						return
																					} else {
																						v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																						v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
																						F_touchWatchedKey(m, v32, v124)
																						mBase = m.M
																						v126 = m.ExcPending
																						if v126 != 0 {
																							return
																						} else {
																							F_trackingInvalidateKey(m, l0, v124, int32(1))
																							mBase = m.M
																							v129 = m.ExcPending
																							if v129 != 0 {
																								return
																							} else {
																								v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																								v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																								v134 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																								F_notifyKeyspaceEvent(m, int32(4), int32(_a581), v133, v134)
																								mBase = m.M
																								v136 = m.ExcPending
																								if v136 != 0 {
																									return
																								} else {
																									v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																									v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
																									v141 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
																									F_notifyKeyspaceEvent(m, int32(4), int32(_a582), v140, v141)
																									mBase = m.M
																									v143 = m.ExcPending
																									if v143 != 0 {
																										return
																									} else {
																										v144 = int32(_a44)
																										v146 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																										*(*int64)(unsafe.Add(mBase, _consts[83])) = v146 + int64(1)
																										v151 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																										F_addReply(m, l0, v151)
																										mBase = m.M
																										v153 = m.ExcPending
																										if v153 != 0 {
																											return
																										} else {
																											m.G0 = v12 + int32(16)
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
															} else {
																v97 = F_getKeySlot(m, v91)
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return
																} else {
																	v99 = v97
																	v101 = F_dbGenericDeleteWithDictIndex(m, v14, v90, v88, int32(1), v99)
																	mBase = m.M
																	v102 = m.ExcPending
																	if v102 != 0 {
																		return
																	} else {
																		v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
																		F_dbAddInternal(m, v32, v104, v12+int32(12), int32(0))
																		mBase = m.M
																		v109 = m.ExcPending
																		if v109 != 0 {
																			return
																		} else {
																			if v73 == int64(-1) {
																				v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
																				F_touchWatchedKey(m, v14, v117)
																				mBase = m.M
																				v119 = m.ExcPending
																				if v119 != 0 {
																					return
																				} else {
																					F_trackingInvalidateKey(m, l0, v117, int32(1))
																					mBase = m.M
																					v122 = m.ExcPending
																					if v122 != 0 {
																						return
																					} else {
																						v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																						v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
																						F_touchWatchedKey(m, v32, v124)
																						mBase = m.M
																						v126 = m.ExcPending
																						if v126 != 0 {
																							return
																						} else {
																							F_trackingInvalidateKey(m, l0, v124, int32(1))
																							mBase = m.M
																							v129 = m.ExcPending
																							if v129 != 0 {
																								return
																							} else {
																								v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																								v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																								v134 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																								F_notifyKeyspaceEvent(m, int32(4), int32(_a581), v133, v134)
																								mBase = m.M
																								v136 = m.ExcPending
																								if v136 != 0 {
																									return
																								} else {
																									v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																									v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
																									v141 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
																									F_notifyKeyspaceEvent(m, int32(4), int32(_a582), v140, v141)
																									mBase = m.M
																									v143 = m.ExcPending
																									if v143 != 0 {
																										return
																									} else {
																										v144 = int32(_a44)
																										v146 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																										*(*int64)(unsafe.Add(mBase, _consts[83])) = v146 + int64(1)
																										v151 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																										F_addReply(m, l0, v151)
																										mBase = m.M
																										v153 = m.ExcPending
																										if v153 != 0 {
																											return
																										} else {
																											m.G0 = v12 + int32(16)
																											return
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			} else {
																				v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
																				v114 = F_setExpire(m, l0, v32, v113, v73)
																				mBase = m.M
																				v115 = m.ExcPending
																				if v115 != 0 {
																					return
																				} else {
																					v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																					v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
																					F_touchWatchedKey(m, v14, v117)
																					mBase = m.M
																					v119 = m.ExcPending
																					if v119 != 0 {
																						return
																					} else {
																						F_trackingInvalidateKey(m, l0, v117, int32(1))
																						mBase = m.M
																						v122 = m.ExcPending
																						if v122 != 0 {
																							return
																						} else {
																							v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																							v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
																							F_touchWatchedKey(m, v32, v124)
																							mBase = m.M
																							v126 = m.ExcPending
																							if v126 != 0 {
																								return
																							} else {
																								F_trackingInvalidateKey(m, l0, v124, int32(1))
																								mBase = m.M
																								v129 = m.ExcPending
																								if v129 != 0 {
																									return
																								} else {
																									v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																									v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
																									v134 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																									F_notifyKeyspaceEvent(m, int32(4), int32(_a581), v133, v134)
																									mBase = m.M
																									v136 = m.ExcPending
																									if v136 != 0 {
																										return
																									} else {
																										v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																										v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
																										v141 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
																										F_notifyKeyspaceEvent(m, int32(4), int32(_a582), v140, v141)
																										mBase = m.M
																										v143 = m.ExcPending
																										if v143 != 0 {
																											return
																										} else {
																											v144 = int32(_a44)
																											v146 = *(*int64)(unsafe.Add(mBase, _consts[83]))
																											*(*int64)(unsafe.Add(mBase, _consts[83])) = v146 + int64(1)
																											v151 = *(*int32)(unsafe.Add(mBase, _consts[347]))
																											F_addReply(m, l0, v151)
																											mBase = m.M
																											v153 = m.ExcPending
																											if v153 != 0 {
																												return
																											} else {
																												m.G0 = v12 + int32(16)
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
													} else {
														v82 = *(*int32)(unsafe.Add(mBase, _consts[85]))
														F_addReply(m, l0, v82)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return
														} else {
															m.G0 = v12 + int32(16)
															return
														}
													}
												}
											} else {
												v56 = *(*int32)(unsafe.Add(mBase, _consts[85]))
												F_addReply(m, l0, v56)
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return
												} else {
													m.G0 = v12 + int32(16)
													return
												}
											}
										}
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, _consts[348]))
										F_addReplyErrorObject(m, l0, v46)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
											return
										}
									}
								}
							}
						}
					}
				} else {
					F_addReplyError(m, l0, int32(_a580))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						m.G0 = v12 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_moveDbIdArgs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 < int32(3) {
		v35 = int32(0)
		m.G0 = v9 + int32(16)
		return v35
	} else {
		v14 = int32(0)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v18 = F_getLongLongFromObject(m, v15, v9+int32(8))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				v35 = v14
				m.G0 = v9 + int32(16)
				return v35
			} else {
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
				if v22 < int64(0) {
					v35 = v14
					m.G0 = v9 + int32(16)
					return v35
				} else {
					v26 = int64(*(*int32)(unsafe.Add(mBase, _consts[65])))
					if v26 <= v22 {
						v35 = v14
						m.G0 = v9 + int32(16)
						return v35
					} else {
						v29 = F_valkey_malloc(m, int32(4))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
							v35 = v29
							m.G0 = v9 + int32(16)
							return v35
						}
					}
				}
			}
		}
	}
}
func F_moveEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v8 = int32(1)
	v11 = v8 << (uint(l1) % 32)
	if int32(base.Ui32(v7)>>(uint(v8)%32))&v11&int32(4095) != 0 {
		F__serverAssert(m, int32(_a834), int32(_a827), int32(946))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
		v16 = int32(1)
		v19 = v16 << (uint(l3) % 32)
		if int32(base.Ui32(v15)>>(uint(v16)%32))&v19&int32(4095) == int32(0) {
			F__serverAssert(m, int32(_a835), int32(_a827), int32(947))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v25 = int32(2)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l2+l3<<(uint(v25)%32))+16))
			*(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(v25)%32)+int32(16)))) = v33
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+l3)+2)))
			*(*uint8)(unsafe.Add(mBase, uint32(l0+l1+v25))) = uint8(v39)
			v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
			v42 = int32(1)
			v46 = v41 | v11<<(uint(v42)%32)&int32(8190)
			*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v46)
			v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
			v55 = ((v19^int32(-1))<<(uint(v42)%32) | int32(57345)) & v54
			*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v55)
			return
		}
	}
}
func F_moveToNextTask(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v9 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v81 = m.G3
	m.Env.X__assert_fail(m, v81+int32(_a1905), v81+int32(_a1906), int32(280), v81+int32(_a1907))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	v72 = m.G3
	m.Env.X__assert_fail(m, v72+int32(_a1908), v72+int32(_a1906), int32(275), v72+int32(_a1907))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	return
L4:
	;
	if v9 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v36 <= v41 {
		goto L1
	} else {
		goto L15
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(-1)
	return
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v16 = v9
	goto L8
L8:
	;
	v25 = v14 + v16<<(uint(int32(2))%32)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(-4))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if base.Ui32(v29+int32(-9)) < base.Ui32(int32(4)) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L6
L10:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v41 = base.I64_extend_i32_s(v40)
	if v36+int64(-1) != v41 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	if v29 != int32(2) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v44 = v16 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v44
	if v44 != 0 {
		v16 = v44
		goto L8
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+8)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v40 + int32(1)
	goto L3
}
func F_mpopGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l2 < int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return
L2:
	;
	F_addReplyNullArray(m, l0)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L7
	} else {
		goto L41
	}
L3:
	;
	v24 = int32(0)
	goto L4
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1+v24<<(uint(int32(2))%32))))
	v33 = F_lookupKeyWrite(m, v28, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L2
L6:
	;
	v96 = v24 + int32(1)
	if v96 != l2 {
		v24 = v96
		goto L4
	} else {
		goto L40
	}
L7:
	;
	return
L8:
	;
	if v33 == int32(0) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v38 = F_checkType(m, l0, v33, int32(1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v40 = F_listTypeLength(m, v33)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v40 == int32(0) {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v44 = F_listTypeLength(m, v33)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	F_initDeferredReplyBuffer(m, l0)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	F_addReplyBulk(m, l0, v32)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	if l4 < v44 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v55 = l4
	goto L20
L19:
	;
	v55 = v44
	goto L20
L20:
	;
	if l3 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v58 = int32(0) - v55
	goto L23
L22:
	;
	v58 = int32(0)
	goto L23
L23:
	;
	v59 = int32(-1)
	if l3 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v62 = v59
	goto L26
L25:
	;
	v62 = v55 + v59
	goto L26
L26:
	;
	F_addListRangeReply(m, l0, v33, v58, v62, base.B2i32(l3 != int32(0)))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	F_listTypeDelRange(m, v33, v58, v55)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	F_listElementsRemoved(m, l0, v32, l3, v33, v55, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	F_commitDeferredReplyBuffer(m, l0, int32(1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	if l4 < v40 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v76 = l4
	goto L33
L32:
	;
	v76 = v40
	goto L33
L33:
	;
	v78 = F_createStringObjectFromLongLong(m, base.I64_extend_i32_s(v76))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v78
	if l3 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v85 = int32(248)
	goto L37
L36:
	;
	v85 = int32(252)
	goto L37
L37:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)+uint32(_consts[84])))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v87
	F_rewriteClientCommandVector(m, l0, int32(3), v13)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	F_decrRefCount(m, v78)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L1
L40:
	;
	goto L5
L41:
	;
	goto L1
}
func F_mpscInit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	if base.I32_popcnt(l1) != int32(1) {
		F__serverAssert(m, int32(_a1881), int32(_a1882), int32(14))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v9 = F_valkey_malloc(m, l1<<(uint(int32(2))%32))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v9
			v13 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v13
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v13
			if l1 == int32(0) {
			} else {
				v21 = int32(0)
				for {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					*(*int32)(unsafe.Add(mBase, uint32(v23+v21<<(uint(int32(2))%32)))) = int32(0)
					v30 = v21 + int32(1)
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					if base.Ui32(v30) < base.Ui32(v31) {
						v21 = v30
						continue
					} else {
						break
					}
					break
				}
			}
			return
		}
	}
}
func F_mutexQueueCreate(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = F_valkey_malloc(m, int32(80))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v15 = F_fifoCreate(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v3))) = v15
			v18 = F_fifoCreate(m)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v3)+4)) = v18
				return v3
			}
		}
	}
}
