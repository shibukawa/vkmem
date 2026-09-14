package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___clock_gettime(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v49 int32
	_ = v49
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if base.Ui32(l0) < base.Ui32(int32(4)) {
		v19 = m.Wasi_snapshot_preview1.Clock_time_get(m, l0, int64(1), v7+int32(24))
		mBase = m.M
		if v19 != 0 {
			v21 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v21))) = v19
		} else {
		}
		if v19 != 0 {
			v49 = int32(-1)
		} else {
			v26 = v7 + int32(8)
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v7)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = int32(0)
			v31 = int64(1000000000)
			v32 = base.I64_div_u_s(v27, v31)
			*(*int64)(unsafe.Add(mBase, uint32(v26))) = v32
			v36 = v27 - v32*v31
			*(*uint32)(unsafe.Add(mBase, uint32(v26)+8)) = uint32(v36)
			v44 = *(*int64)(unsafe.Add(mBase, uint32(v7+int32(16))))
			*(*int64)(unsafe.Add(mBase, uint32(l1+int32(8)))) = v44
			v46 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v46
			v49 = int32(0)
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F___clock_gettime[0])) = int32(28)
		v49 = int32(-1)
	}
	m.G0 = v7 + int32(32)
	return v49
}
func F__crc64(m *base.Module, l0 int64, l1 int32, l2 int64) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v17 int64
	_ = v17
	var v21 int64
	_ = v21
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v39 int64
	_ = v39
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v54 int64
	_ = v54
	var v60 int64
	_ = v60
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v69 int64
	_ = v69
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v84 int64
	_ = v84
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v99 int64
	_ = v99
	var v105 int64
	_ = v105
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v114 int64
	_ = v114
	var v120 int64
	_ = v120
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v129 int64
	_ = v129
	var v135 int64
	_ = v135
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v152 int64
	_ = v152
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v161 int64
	_ = v161
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v185 int64
	_ = v185
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v194 int64
	_ = v194
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v203 int64
	_ = v203
	v4 = int64(0)
	if l2 == v4 {
		v142 = l0
	} else {
		v10 = l0
		v13 = v4
		for {
			v17 = v10 << (uint(int64(1)) % 64)
			v21 = v10 & int64(-9223372036854775807-1)
			v27 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+base.I32_wrap_i64(v13)))))
			if v27&int32(1) != 0 {
				v30 = base.I64_extend_i32_u(base.B2i32(v21 == int64(0)))
			} else {
				v30 = v21
			}
			if v30 == int64(0) {
				v33 = v17
			} else {
				v33 = v17 ^ int64(-5939172356000238167)
			}
			v35 = v33 << (uint(int64(1)) % 64)
			v39 = v33 & int64(-9223372036854775807-1)
			if v27&int32(2) != 0 {
				v45 = base.I64_extend_i32_u(base.B2i32(v39 == int64(0)))
			} else {
				v45 = v39
			}
			if v45 == int64(0) {
				v48 = v35
			} else {
				v48 = v35 ^ int64(-5939172356000238167)
			}
			v50 = v48 << (uint(int64(1)) % 64)
			v54 = v48 & int64(-9223372036854775807-1)
			if v27&int32(4) != 0 {
				v60 = base.I64_extend_i32_u(base.B2i32(v54 == int64(0)))
			} else {
				v60 = v54
			}
			if v60 == int64(0) {
				v63 = v50
			} else {
				v63 = v50 ^ int64(-5939172356000238167)
			}
			v65 = v63 << (uint(int64(1)) % 64)
			v69 = v63 & int64(-9223372036854775807-1)
			if v27&int32(8) != 0 {
				v75 = base.I64_extend_i32_u(base.B2i32(v69 == int64(0)))
			} else {
				v75 = v69
			}
			if v75 == int64(0) {
				v78 = v65
			} else {
				v78 = v65 ^ int64(-5939172356000238167)
			}
			v80 = v78 << (uint(int64(1)) % 64)
			v84 = v78 & int64(-9223372036854775807-1)
			if v27&int32(16) != 0 {
				v90 = base.I64_extend_i32_u(base.B2i32(v84 == int64(0)))
			} else {
				v90 = v84
			}
			if v90 == int64(0) {
				v93 = v80
			} else {
				v93 = v80 ^ int64(-5939172356000238167)
			}
			v95 = v93 << (uint(int64(1)) % 64)
			v99 = v93 & int64(-9223372036854775807-1)
			if v27&int32(32) != 0 {
				v105 = base.I64_extend_i32_u(base.B2i32(v99 == int64(0)))
			} else {
				v105 = v99
			}
			if v105 == int64(0) {
				v108 = v95
			} else {
				v108 = v95 ^ int64(-5939172356000238167)
			}
			v110 = v108 << (uint(int64(1)) % 64)
			v114 = v108 & int64(-9223372036854775807-1)
			if v27&int32(64) != 0 {
				v120 = base.I64_extend_i32_u(base.B2i32(v114 == int64(0)))
			} else {
				v120 = v114
			}
			if v120 == int64(0) {
				v123 = v110
			} else {
				v123 = v110 ^ int64(-5939172356000238167)
			}
			v125 = v123 << (uint(int64(1)) % 64)
			v129 = v123 & int64(-9223372036854775807-1)
			if v27 < int32(0) {
				v135 = base.I64_extend_i32_u(base.B2i32(v129 == int64(0)))
			} else {
				v135 = v129
			}
			if v135 == int64(0) {
				v138 = v125
			} else {
				v138 = v125 ^ int64(-5939172356000238167)
			}
			v140 = v13 + int64(1)
			if v140 != l2 {
				v10 = v138
				v13 = v140
				continue
			} else {
				break
			}
			break
		}
		v142 = v138
	}
	v148 = int64(56)
	v150 = int64(65280)
	v152 = int64(40)
	v155 = int64(16711680)
	v157 = int64(24)
	v159 = int64(4278190080)
	v161 = int64(8)
	v182 = v142<<(uint(v148)%64) | v142&v150<<(uint(v152)%64) | (v142&v155<<(uint(v157)%64) | v142&v159<<(uint(v161)%64)) | (int64(base.Ui64(v142)>>(uint(v161)%64))&v159 | int64(base.Ui64(v142)>>(uint(v157)%64))&v155 | (int64(base.Ui64(v142)>>(uint(v152)%64))&v150 | int64(base.Ui64(v142)>>(uint(v148)%64))))
	v183 = int64(4)
	v185 = int64(1085102592571150095)
	v191 = int64(base.Ui64(v182)>>(uint(v183)%64))&v185 | v182&v185<<(uint(v183)%64)
	v192 = int64(2)
	v194 = int64(3689348814741910323)
	v200 = int64(base.Ui64(v191)>>(uint(v192)%64))&v194 | v191&v194<<(uint(v192)%64)
	v201 = int64(1)
	v203 = int64(6148914691236517205)
	return int64(base.Ui64(v200)>>(uint(v201)%64))&v203 | v200&v203<<(uint(v201)%64)
}
func F_casemap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	if base.Ui32(int32(131071)) < base.Ui32(l0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	v12 = int32(255)
	v13 = l0 & v12
	v14 = int32(3)
	v15 = base.I32_div_u_s(v13, v14)
	v21 = int32(2)
	v25 = *(*int32)(unsafe.Add(mBase, uint32((l0-v15*v14)&v12<<(uint(v21)%32))+uint32(_c_F_casemap[0])))
	v26 = int32(8)
	v27 = int32(base.Ui32(l0) >> (uint(v26) % 32))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_casemap[1]))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30*int32(86)+v15)+uint32(_c_F_casemap[1]))))
	v41 = base.I32_rem_u_s(int32(base.Ui32(v25*v36)>>(uint(int32(11))%32)), int32(6))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_casemap[2]))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32((v41+v44)<<(uint(v21)%32))+uint32(_c_F_casemap[3])))
	v52 = v50 >> (uint(v26) % 32)
	v54 = v50 & v12
	if base.Ui32(int32(1)) < base.Ui32(v54) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v64 = v52 & int32(255)
	if v64 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	return v52&(int32(0)-(v54^l1)) + l0
L5:
	;
	v72 = v64
	v73 = int32(base.Ui32(v52) >> (uint(int32(8)) % 32))
	goto L6
L6:
	;
	v78 = int32(1)
	v79 = int32(base.Ui32(v72) >> (uint(v78) % 32))
	v80 = v79 + v73
	v82 = v80 << (uint(v78) % 32)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+uint32(_c_F_casemap[4]))))
	if v13 != v85 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	v110 = base.B2i32(base.Ui32(v13) < base.Ui32(v85))
	if base.Ui32(v13) < base.Ui32(v85) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+uint32(_c_F_casemap[5]))))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87<<(uint(int32(2))%32))+uint32(_c_F_casemap[3])))
	v94 = v92 & int32(255)
	if base.Ui32(int32(1)) < base.Ui32(v94) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if l1 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	return v92>>(uint(int32(8))%32)&(int32(0)-(v94^l1)) + l0
L12:
	;
	v107 = int32(-1)
	goto L14
L13:
	;
	v107 = int32(1)
	goto L14
L14:
	;
	return v107 + l0
L15:
	;
	v111 = v73
	goto L17
L16:
	;
	v111 = v80
	goto L17
L17:
	;
	if base.Ui32(v13) < base.Ui32(v85) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v113 = v79
	goto L20
L19:
	;
	v113 = v72 - v79
	goto L20
L20:
	;
	if v113 != 0 {
		v72 = v113
		v73 = v111
		goto L6
	} else {
		goto L21
	}
L21:
	;
	goto L7
}
func F_categoryFlagsFromString(m *base.Module, l0 int32) int64 {
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
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int64
	_ = v127
	var v134 int64
	_ = v134
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int64
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int64
	_ = v166
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0&int32(3) == int32(0) {
		v32 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v70 = F_sdssplitlen(m, l0, v65, int32(_a_F_categoryFlagsFromString_0), int32(1), v9+int32(12))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v65 = v57 - l0
	goto L1
L3:
	;
	v36 = v32
	goto L11
L4:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = l0
	goto L7
L6:
	;
	v65 = l0 - l0
	goto L1
L7:
	;
	v25 = v21 + int32(1)
	if v25&int32(3) == int32(0) {
		v32 = v25
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v30 != 0 {
		v21 = v25
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v57 = v25
	goto L2
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v45 = int32(-2139062144)
	if (int32(16843008)-v42|v42)&v45 == v45 {
		v36 = v36 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v51 = v36
	goto L14
L13:
	;
	goto L12
L14:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 != 0 {
		v51 = v51 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v57 = v51
	goto L2
L16:
	;
	goto L15
L17:
	;
	return int64(0)
L18:
	;
	v74 = int64(0)
	v75 = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v76 < int32(1) {
		v151 = v75
		v154 = v74
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	F_sdsfreesplitres(m, v70, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L17
	} else {
		goto L38
	}
L20:
	;
	v79 = v75
	v82 = v74
	goto L21
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v70+v79<<(uint(int32(2))%32))))
	v89 = int32(0)
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_categoryFlagsFromString[0]))
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v95)+8))
	if base.B2i32(v96 == int64(0)) == v89 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v151 = v148
	v154 = v146
	goto L19
L23:
	;
	v146 = v134 | v82
	v148 = v79 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v148 < v149 {
		v79 = v148
		v82 = v146
		goto L21
	} else {
		goto L37
	}
L24:
	;
	if v134 != int64(0) {
		goto L23
	} else {
		goto L34
	}
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v103 = F_strcasecmp(m, v88, v102)
	mBase = m.M
	if v103 == int32(0) {
		v127 = v96
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v134 = int64(0)
	goto L24
L27:
	;
	v134 = v127
	goto L24
L28:
	;
	v107 = v89
	goto L29
L29:
	;
	v112 = v107 + int32(1)
	v115 = v95 + v112<<(uint(int32(4))%32)
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v115)+8))
	if base.B2i32(v116 == int64(0)) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v127 = v116
	goto L27
L31:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v123 = F_strcasecmp(m, v88, v122)
	mBase = m.M
	if v123 != 0 {
		v107 = v112
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v134 = int64(0)
	goto L24
L33:
	;
	goto L30
L34:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_categoryFlagsFromString[1]))
	if int32(3) < v138 {
		v151 = v79
		v154 = v82
		goto L19
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v88
	F__serverLog(m, int32(3), int32(_a_F_categoryFlagsFromString_1), v9)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L17
	} else {
		goto L36
	}
L36:
	;
	v151 = v79
	v154 = v82
	goto L19
L37:
	;
	goto L22
L38:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	m.G0 = v9 + int32(16)
	if v151 == v160 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v166 = v154
	goto L41
L40:
	;
	v166 = int64(-1)
	goto L41
L41:
	;
	return v166
}
func F_ceill(m *base.Module, l0 int32, l1 int64, l2 int64) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v101 int64
	_ = v101
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v117 int64
	_ = v117
	var v122 int64
	_ = v122
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v137 int64
	_ = v137
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int64
	_ = v152
	var v155 int64
	_ = v155
	var v163 int64
	_ = v163
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v184 int32
	_ = v184
	var v208 int32
	_ = v208
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v247 int64
	_ = v247
	var v248 int64
	_ = v248
	var v249 int64
	_ = v249
	var v250 int64
	_ = v250
	v7 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	v22 = l2 & int64(9223372036854775807)
	v23 = int64(9223090561878065152)
	if v22 == v23 {
		v27 = base.B2i32(l1 != v7)
	} else {
		v27 = base.B2i32(base.Ui64(v23) < base.Ui64(v22))
	}
	if v27 != 0 {
		v72 = int32(1)
		v76 = v72
	} else {
		if base.B2i32(v7|l1|(int64(0)|v22) == int64(0)) == int32(0) {
			if v7&l2 < int64(0) {
				if l2 == v7 {
					v63 = base.B2i32(base.Ui64(v7) < base.Ui64(l1))
				} else {
					v63 = base.B2i32(v7 < l2)
				}
				if v63 == int32(0) {
					v72 = base.B2i32(l1^v7|(l2^v7) != int64(0))
					v76 = v72
				} else {
					v76 = int32(-1)
				}
			} else {
				if l2 == v7 {
					v51 = base.B2i32(base.Ui64(l1) < base.Ui64(v7))
				} else {
					v51 = base.B2i32(l2 < v7)
				}
				if v51 == int32(0) {
					v76 = base.B2i32(l1^v7|(l2^v7) != int64(0))
				} else {
					v76 = int32(-1)
				}
			}
		} else {
			v76 = int32(0)
		}
	}
	if v76 == int32(0) {
		v249 = l1
		v250 = l2
	} else {
		v81 = base.I32_wrap_i64(int64(base.Ui64(l2) >> (uint(int64(48)) % 64)))
		v83 = v81 & int32(32767)
		if base.Ui32(int32(16494)) < base.Ui32(v83) {
			v249 = l1
			v250 = l2
		} else {
			if base.Ui32(int32(16382)) < base.Ui32(v83) {
				v96 = int64(0)
				v97 = int64(4642929740842270720)
				F___addtf3(m, v11+int32(96), l1, l2, v96, v97)
				mBase = m.M
				v101 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
				v106 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(104))))
				v108 = int64(-4580442296012505088)
				F___addtf3(m, v11+int32(80), v101, v106, v96, v108)
				mBase = m.M
				F___addtf3(m, v11+int32(64), l1, l2, v96, v108)
				mBase = m.M
				v117 = *(*int64)(unsafe.Add(mBase, uint32(v11)+64))
				v122 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(72))))
				F___addtf3(m, v11+int32(48), v117, v122, v96, v97)
				mBase = m.M
				v127 = v11 + int32(32)
				v128 = *(*int64)(unsafe.Add(mBase, uint32(v11)+80))
				v129 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
				v131 = base.B2i32(base.Ui32(v81) < base.Ui32(int32(32768)))
				if base.Ui32(v81) < base.Ui32(int32(32768)) {
					v132 = v128
				} else {
					v132 = v129
				}
				v137 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(88))))
				v142 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(56))))
				if base.Ui32(v81) < base.Ui32(int32(32768)) {
					v143 = v137
				} else {
					v143 = v142
				}
				v145 = m.G0
				v146 = int32(16)
				v147 = v145 - v146
				m.G0 = v147
				F___addtf3(m, v147, v132, v143, l1, l2^int64(-9223372036854775807-1))
				mBase = m.M
				v152 = *(*int64)(unsafe.Add(mBase, uint32(v147)))
				v155 = *(*int64)(unsafe.Add(mBase, uint32(v147+int32(8))))
				*(*int64)(unsafe.Add(mBase, uint32(v127)+8)) = v155
				*(*int64)(unsafe.Add(mBase, uint32(v127))) = v152
				m.G0 = v147 + v146
				v163 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
				v168 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(40))))
				F___addtf3(m, v11+int32(16), l1, l2, v163, v168)
				mBase = m.M
				v170 = int64(0)
				v179 = v168 & int64(9223372036854775807)
				v180 = int64(9223090561878065152)
				if v179 == v180 {
					v184 = base.B2i32(v163 != v170)
				} else {
					v184 = base.B2i32(base.Ui64(v180) < base.Ui64(v179))
				}
				if v184 != 0 {
					v229 = int32(1)
					v233 = v229
				} else {
					if base.B2i32(v170|v163|(int64(0)|v179) == int64(0)) == int32(0) {
						if v170&v168 < int64(0) {
							if v168 == v170 {
								v220 = base.B2i32(base.Ui64(v170) < base.Ui64(v163))
							} else {
								v220 = base.B2i32(v170 < v168)
							}
							if v220 == int32(0) {
								v229 = base.B2i32(v163^v170|(v168^v170) != int64(0))
								v233 = v229
							} else {
								v233 = int32(-1)
							}
						} else {
							if v168 == v170 {
								v208 = base.B2i32(base.Ui64(v163) < base.Ui64(v170))
							} else {
								v208 = base.B2i32(v168 < v170)
							}
							if v208 == int32(0) {
								v233 = base.B2i32(v163^v170|(v168^v170) != int64(0))
							} else {
								v233 = int32(-1)
							}
						}
					} else {
						v233 = int32(0)
					}
				}
				v238 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(24))))
				v239 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
				if int32(-1) < v233 {
					v249 = v239
					v250 = v238
				} else {
					F___addtf3(m, v11, v239, v238, int64(0), int64(4611404543450677248))
					mBase = m.M
					v247 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(8))))
					v248 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
					v249 = v248
					v250 = v247
				}
			} else {
				if base.Ui32(v81) < base.Ui32(int32(32768)) {
					v92 = int64(4611404543450677248)
				} else {
					v92 = int64(-9223372036854775807 - 1)
				}
				v249 = int64(0)
				v250 = v92
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v249
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v250
	m.G0 = v11 + int32(112)
	return
}
func F_changeReplicationId(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	F_getRandomHexChars(m, int32(_a_F_changeReplicationId_0), int32(40))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v8 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_changeReplicationId[0])) = uint8(v8)
		return
	}
}
func F_channelList(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v231 int32
	_ = v231
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	goto L1
L1:
	;
	v25 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_setDeferredArrayLen(m, l0, v25, v217)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L2
	} else {
		goto L43
	}
L5:
	;
	v38 = int32(0)
	v46 = v38
	v52 = v38
	goto L7
L6:
	;
	v217 = int32(0)
	goto L4
L7:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v52<<(uint(int32(2))%32))))
	if v63 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v217 = v195
	goto L4
L9:
	;
	v209 = v52 + int32(1)
	if v209 != v24 {
		v46 = v195
		v52 = v209
		goto L7
	} else {
		goto L42
	}
L10:
	;
	if v66 == int32(0) {
		v195 = v46
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v65 = F_hashtableSize(m, v63)
	mBase = m.M
	v66 = v65
	goto L10
L12:
	;
	v66 = int32(0)
	goto L10
L13:
	;
	v70 = F_kvstoreGetHashtableIterator(m, l2, v52, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L15
	}
L14:
	;
	F_kvstoreReleaseHashtableIterator(m, v70)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L2
	} else {
		goto L41
	}
L15:
	;
	v74 = F_kvstoreHashtableIteratorNext(m, v70, v22+int32(12))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if v74 == int32(0) {
		v174 = v46
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v84 = v46
	goto L18
L18:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	goto L20
L19:
	;
	v174 = v160
	goto L14
L20:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97+int32(44))))
	v101 = F_objectGetVal(m, v100)
	mBase = m.M
	if l1 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v166 = F_kvstoreHashtableIteratorNext(m, v70, v22+int32(12))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L39
	}
L22:
	;
	F_addReplyBulk(m, l0, v100)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L38
	}
L23:
	;
	v104 = int32(0)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v106 & int32(7) {
	case 0:
		goto L29
	case 1:
		goto L28
	case 2:
		goto L27
	case 3:
		goto L26
	case 4:
		goto L25
	default:
		v115 = v104
		goto L24
	}
L24:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+int32(-1)))))
	switch v118 & int32(7) {
	case 0:
		goto L35
	case 1:
		goto L34
	case 2:
		goto L33
	case 3:
		goto L32
	case 4:
		goto L31
	default:
		v135 = v104
		goto L30
	}
L25:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v115 = v114
	goto L24
L26:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v115 = v113
	goto L24
L27:
	;
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v115 = v112
	goto L24
L28:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v115 = v111
	goto L24
L29:
	;
	v115 = int32(base.Ui32(v106) >> (uint(int32(3)) % 32))
	goto L24
L30:
	;
	v136 = int32(0)
	v138 = m.G0
	v139 = int32(16)
	v140 = v138 - v139
	m.G0 = v140
	*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = v136
	v147 = F_stringmatchlen_impl(m, l1, v115, v101, v135, v136, v140+int32(12), v136)
	mBase = m.M
	m.G0 = v140 + v139
	goto L36
L31:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v101+int32(-17))))
	v135 = v134
	goto L30
L32:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v101+int32(-9))))
	v135 = v131
	goto L30
L33:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101+int32(-5)))))
	v135 = v128
	goto L30
L34:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+int32(-3)))))
	v135 = v125
	goto L30
L35:
	;
	v135 = int32(base.Ui32(v118) >> (uint(int32(3)) % 32))
	goto L30
L36:
	;
	if v147 == int32(0) {
		v160 = v84
		goto L21
	} else {
		goto L37
	}
L37:
	;
	goto L22
L38:
	;
	v160 = v84 + int32(1)
	goto L21
L39:
	;
	if v166 != 0 {
		v84 = v160
		goto L18
	} else {
		goto L40
	}
L40:
	;
	goto L19
L41:
	;
	v195 = v174
	goto L9
L42:
	;
	goto L8
L43:
	;
	m.G0 = v22 + int32(16)
	return
}
func F_chdir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v2 = m.Env.X__syscall_chdir(m, l0)
	mBase = m.M
	if base.Ui32(v2) < base.Ui32(int32(-4095)) {
		v10 = v2
	} else {
		v5 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(0) - v2
		v10 = int32(-1)
	}
	return v10
}
func F_check_match(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v12 != l1 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v18 = F_luaX_token2str(m, l0, l1)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			if l3 != v16 {
				v29 = F_luaX_token2str(m, l0, l2)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v18
					v34 = m.G3
					v39 = F_luaO_pushfstring(m, v17, v34+int32(_a_F_check_match_0), v10+int32(16))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_luaX_syntaxerror(m, l0, v39)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							m.G0 = v10 + int32(32)
							return
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
				v22 = m.G3
				v25 = F_luaO_pushfstring(m, v17, v22+int32(_a_F_check_match_1), v10)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_luaX_syntaxerror(m, l0, v25)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						m.G0 = v10 + int32(32)
						return
					}
				}
			}
		}
	} else {
		F_luaX_next(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			m.G0 = v10 + int32(32)
			return
		}
	}
}
func F_chmod(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v3 = m.Env.X__syscall_chmod(m, l0, l1)
	mBase = m.M
	if base.Ui32(v3) < base.Ui32(int32(-4095)) {
		v11 = v3
	} else {
		v6 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(0) - v3
		v11 = int32(-1)
	}
	return v11
}
func F_chown(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v6 = m.Env.X__syscall_fchownat(m, int32(-100), l0, l1, l2, int32(0))
	mBase = m.M
	if base.Ui32(v6) < base.Ui32(int32(-4095)) {
		v14 = v6
	} else {
		v9 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(0) - v6
		v14 = int32(-1)
	}
	return v14
}
func F_clearCachedClusterSlotsResponse(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[0]))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[1]))
	if v12 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	F_sdsfree(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[0])) = int32(0)
	goto L1
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[2]))
	if v21 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_sdsfree(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[1])) = int32(0)
	goto L5
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[3]))
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_sdsfree(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[2])) = int32(0)
	goto L8
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[4]))
	if v39 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	F_sdsfree(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[3])) = int32(0)
	goto L11
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[5]))
	if v48 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	F_sdsfree(m, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[4])) = int32(0)
	goto L14
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[6]))
	if v57 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	F_sdsfree(m, v48)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[5])) = int32(0)
	goto L17
L20:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[7]))
	if v66 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	F_sdsfree(m, v57)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[6])) = int32(0)
	goto L20
L23:
	;
	return
L24:
	;
	F_sdsfree(m, v66)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_clearCachedClusterSlotsResponse[7])) = int32(0)
	goto L23
}
func F_close_func(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
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
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v118 int32
	_ = v118
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
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
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+50)))
	if v15 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v131 = int32(0)
	F_luaK_ret(m, v13, v131, v131)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v19 = v13 + int32(172)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v23 = v15 & int32(3)
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if base.Ui32(v15) < base.Ui32(int32(4)) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v34 = v15
	v35 = int32(0)
	goto L6
L5:
	;
	v58 = v15
	goto L3
L6:
	;
	v37 = v34 + int32(-1)
	v38 = int32(1)
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+v37<<(uint(v38)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v20+v41*int32(12))+8)) = v21
	v47 = v35 + v38
	if v47 != v23 {
		v34 = v37
		v35 = v47
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v58 = v37
	goto L3
L8:
	;
	goto L7
L9:
	;
	v118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+50)) = uint8(v118)
	goto L1
L10:
	;
	v71 = v58
	goto L11
L11:
	;
	v73 = int32(1)
	v75 = v71<<(uint(v73)%32) + v19
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75+int32(-2)))))
	v79 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v20+v78*v79)+8)) = v21
	v83 = int32(-4)
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75+v83))))
	*(*int32)(unsafe.Add(mBase, uint32(v20+v85*v79)+8)) = v21
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75+int32(-6)))))
	*(*int32)(unsafe.Add(mBase, uint32(v20+v92*v79)+8)) = v21
	v98 = v71 + v83
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+v98<<(uint(v73)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v20+v102*v79)+8)) = v21
	if v98 != 0 {
		v71 = v98
		goto L11
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	goto L12
L14:
	;
	return
L15:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	if base.Ui32(int32(1073741823)) < base.Ui32(v135+int32(1)) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v152
	if base.Ui32(int32(1073741823)) < base.Ui32(v152+int32(1)) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v148 = F_luaM_toobig(m, v12)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L14
	} else {
		goto L20
	}
L18:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v142 = int32(2)
	v146 = F_luaM_realloc_(m, v12, v140, v141<<(uint(v142)%32), v135<<(uint(v142)%32))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v150 = v146
	goto L16
L20:
	;
	v150 = v148
	goto L16
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v168
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v170
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	if base.Ui32(int32(268435455)) < base.Ui32(v172+int32(1)) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v166 = F_luaM_toobig(m, v12)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L14
	} else {
		goto L25
	}
L23:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v160 = int32(2)
	v164 = F_luaM_realloc_(m, v12, v158, v159<<(uint(v160)%32), v152<<(uint(v160)%32))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	v168 = v164
	goto L21
L25:
	;
	v168 = v166
	goto L21
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if base.Ui32(int32(1073741823)) < base.Ui32(v191+int32(1)) {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	v185 = F_luaM_toobig(m, v12)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L14
	} else {
		goto L30
	}
L28:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v179 = int32(4)
	v183 = F_luaM_realloc_(m, v12, v177, v178<<(uint(v179)%32), v172<<(uint(v179)%32))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	v187 = v183
	goto L26
L30:
	;
	v187 = v185
	goto L26
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v206
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v208
	v210 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+48)))
	if base.Ui32(int32(357913941)) < base.Ui32(v210+int32(1)) {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v204 = F_luaM_toobig(m, v12)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L14
	} else {
		goto L35
	}
L33:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v198 = int32(2)
	v202 = F_luaM_realloc_(m, v12, v196, v197<<(uint(v198)%32), v191<<(uint(v198)%32))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	v206 = v202
	goto L31
L35:
	;
	v206 = v204
	goto L31
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v225
	v227 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v227
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v231 = int32(2)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+72)))
	v236 = F_luaM_realloc_(m, v12, v229, v230<<(uint(v231)%32), v233<<(uint(v231)%32))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L14
	} else {
		goto L41
	}
L37:
	;
	v223 = F_luaM_toobig(m, v12)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L14
	} else {
		goto L40
	}
L38:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v217 = int32(12)
	v221 = F_luaM_realloc_(m, v12, v215, v216*v217, v210*v217)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	v225 = v221
	goto L36
L40:
	;
	v225 = v223
	goto L36
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v236
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+72)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v239
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v241
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(int32(1)) < base.Ui32(v243+int32(-285)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v255 + int32(-32)
	return
L43:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	v252 = F_luaX_newstring(m, l0, v248+int32(16), v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	goto L42
}
func F_closedir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var __phi73 int32
	_ = __phi73
	var v74 int32
	_ = v74
	var __phi74 int32
	_ = __phi74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var __phi240 int32
	_ = __phi240
	var v241 int32
	_ = v241
	var __phi241 int32
	_ = __phi241
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v322 int32
	_ = v322
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = F_close(m, v3)
	mBase = m.M
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v4
L2:
	;
	goto L1
L3:
	;
	v15 = int32(-8)
	v16 = l0 + v15
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-4))))
	v21 = v19 & v15
	v22 = v16 + v21
	if v19&int32(1) != 0 {
		v146 = v21
		v147 = v16
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if base.Ui32(v22) <= base.Ui32(v147) {
		goto L2
	} else {
		goto L39
	}
L5:
	;
	if v19&int32(2) == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v30 = v16 - v29
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_closedir[0]))
	if base.Ui32(v30) < base.Ui32(v32) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v34 = v29 + v21
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_closedir[1]))
	if v30 == v36 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if v52 == int32(0) {
		v146 = v34
		v147 = v30
		goto L4
	} else {
		goto L27
	}
L9:
	;
	v105 = int32(0)
	goto L8
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v41
	v146 = v34
	v147 = v30
	goto L4
L11:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v87 = int32(3)
	if v86&v87 != v87 {
		v146 = v34
		v147 = v30
		goto L4
	} else {
		goto L26
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if base.Ui32(int32(255)) < base.Ui32(v29) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	if v38 == v30 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v38 != v41 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v43 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_closedir[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[2])) = v45 & base.I32_rotl(int32(-2), int32(base.Ui32(v29)>>(uint(int32(3))%32)))
	v146 = v34
	v147 = v30
	goto L4
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	if v57 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v54
	v105 = v38
	goto L8
L18:
	;
	__phi73 = v67
	__phi74 = v68
	v73 = __phi73
	v74 = __phi74
	goto L22
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	if v62 == int32(0) {
		goto L9
	} else {
		goto L21
	}
L20:
	;
	v67 = v57
	v68 = v30 + int32(20)
	goto L18
L21:
	;
	v67 = v62
	v68 = v30 + int32(16)
	goto L18
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	if v80 != 0 {
		__phi73 = v80
		__phi74 = v73 + int32(20)
		v73 = __phi73
		v74 = __phi74
		goto L22
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
	v105 = v73
	goto L8
L24:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	if v83 != 0 {
		__phi73 = v83
		__phi74 = v73 + int32(16)
		v73 = __phi73
		v74 = __phi74
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[3])) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v86 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v34 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v34
	goto L1
L27:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	v116 = v114 << (uint(int32(2)) % 32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116)+uint32(_c_F_closedir[4])))
	if v30 != v119 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v52
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	if v136 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v129 != v30 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+uint32(_c_F_closedir[4]))) = v105
	if v105 != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v122 = int32(0)
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_closedir[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[5])) = v124 & base.I32_rotl(int32(-2), v114)
	v146 = v34
	v147 = v30
	goto L4
L32:
	;
	if v105 == int32(0) {
		v146 = v34
		v147 = v30
		goto L4
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v105
	goto L32
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v105
	goto L32
L35:
	;
	goto L28
L36:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	if v141 == int32(0) {
		v146 = v34
		v147 = v30
		goto L4
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v136)+24)) = v105
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v141)+24)) = v105
	v146 = v34
	v147 = v30
	goto L4
L39:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v156&int32(1) == int32(0) {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	if v156&int32(2) != 0 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	if base.Ui32(int32(255)) < base.Ui32(v322) {
		goto L79
	} else {
		goto L80
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v202 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v147+v202))) = v202
	if v147 != v186 {
		v322 = v202
		goto L41
	} else {
		goto L78
	}
L43:
	;
	if v219 == int32(0) {
		goto L42
	} else {
		goto L66
	}
L44:
	;
	v264 = int32(0)
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v156 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v146 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v147+v146))) = v146
	v322 = v146
	goto L41
L46:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_closedir[6]))
	if v22 != v164 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_closedir[1]))
	if v22 != v186 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v166 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[6])) = v147
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_closedir[7]))
	v171 = v170 + v146
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[7])) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v171 | int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_closedir[1]))
	if v147 != v177 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v179 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[3])) = v179
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[1])) = v179
	goto L1
L50:
	;
	v202 = v156&int32(-8) + v146
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if base.Ui32(int32(255)) < base.Ui32(v156) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v188 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[1])) = v147
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_closedir[3]))
	v193 = v192 + v146
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[3])) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v193 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v147+v193))) = v193
	goto L1
L52:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v203 == v22 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v203 != v206 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+12)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v203)+8)) = v206
	goto L42
L55:
	;
	v208 = int32(0)
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_closedir[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[2])) = v210 & base.I32_rotl(int32(-2), int32(base.Ui32(v156)>>(uint(int32(3))%32)))
	goto L42
L56:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	if v224 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v221)+12)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v203)+8)) = v221
	v264 = v203
	goto L43
L58:
	;
	__phi240 = v234
	__phi241 = v235
	v240 = __phi240
	v241 = __phi241
	goto L62
L59:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v229 == int32(0) {
		goto L44
	} else {
		goto L61
	}
L60:
	;
	v234 = v224
	v235 = v22 + int32(20)
	goto L58
L61:
	;
	v234 = v229
	v235 = v22 + int32(16)
	goto L58
L62:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v240)+20))
	if v247 != 0 {
		__phi240 = v247
		__phi241 = v240 + int32(20)
		v240 = __phi240
		v241 = __phi241
		goto L62
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v241))) = int32(0)
	v264 = v240
	goto L43
L64:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v240)+16))
	if v250 != 0 {
		__phi240 = v250
		__phi241 = v240 + int32(16)
		v240 = __phi240
		v241 = __phi241
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v275 = v273 << (uint(int32(2)) % 32)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)+uint32(_c_F_closedir[4])))
	if v22 != v278 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+24)) = v219
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v295 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	if v288 != v22 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+uint32(_c_F_closedir[4]))) = v264
	if v264 != 0 {
		goto L67
	} else {
		goto L70
	}
L70:
	;
	v281 = int32(0)
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_closedir[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[5])) = v283 & base.I32_rotl(int32(-2), v273)
	goto L42
L71:
	;
	if v264 == int32(0) {
		goto L42
	} else {
		goto L74
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219)+20)) = v264
	goto L71
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219)+16)) = v264
	goto L71
L74:
	;
	goto L67
L75:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	if v300 == int32(0) {
		goto L42
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+16)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v295)+24)) = v264
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+20)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(v300)+24)) = v264
	goto L42
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[3])) = v202
	goto L1
L79:
	;
	if base.Ui32(int32(16777215)) < base.Ui32(v322) {
		v369 = int32(31)
		goto L84
	} else {
		goto L85
	}
L80:
	;
	v334 = v322 & int32(-8)
	v336 = v334 + int32(9128464)
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_closedir[2]))
	v342 = int32(1) << (uint(int32(base.Ui32(v322)>>(uint(int32(3))%32))) % 32)
	if v338&v342 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+uint32(_c_F_closedir[8]))) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v348)+12)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v147)+8)) = v348
	goto L1
L82:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v334)+uint32(_c_F_closedir[8])))
	v348 = v347
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[2])) = v338 | v342
	v348 = v336
	goto L81
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+28)) = v369
	*(*int64)(unsafe.Add(mBase, uint32(v147)+16)) = int64(0)
	v374 = v369 << (uint(int32(2)) % 32)
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_closedir[5]))
	v380 = int32(1) << (uint(v369) % 32)
	if v378&v380 != 0 {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	v359 = base.I32_clz(int32(base.Ui32(v322) >> (uint(int32(8)) % 32)))
	v362 = int32(1)
	v369 = int32(base.Ui32(v322)>>(uint(int32(38)-v359)%32))&v362 - v359<<(uint(v362)%32) + int32(62)
	goto L84
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147+v441))) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v147+v439))) = v442
	v453 = int32(0)
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_closedir[9]))
	v456 = int32(-1)
	v457 = v455 + v456
	if v457 != 0 {
		goto L98
	} else {
		goto L99
	}
L87:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v403)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v433)+12)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v403)+8)) = v147
	v439 = int32(24)
	v441 = int32(8)
	v442 = int32(0)
	v443 = v403
	v444 = v433
	goto L86
L88:
	;
	v439 = v424
	v441 = v426
	v442 = v147
	v443 = v147
	v444 = v429
	goto L86
L89:
	;
	if v369 == int32(31) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[5])) = v378 | v380
	*(*int32)(unsafe.Add(mBase, uint32(v374)+uint32(_c_F_closedir[4]))) = v147
	v424 = int32(8)
	v426 = int32(24)
	v429 = v374 + int32(9128728)
	goto L88
L91:
	;
	v395 = int32(0)
	goto L93
L92:
	;
	v395 = int32(25) - int32(base.Ui32(v369)>>(uint(int32(1))%32))
	goto L93
L93:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v374)+uint32(_c_F_closedir[4])))
	v400 = v322 << (uint(v395) % 32)
	v403 = v397
	goto L94
L94:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	if v407&int32(-8) == v322 {
		goto L87
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v417+int32(16)))) = v147
	v424 = int32(8)
	v426 = int32(24)
	v429 = v403
	goto L88
L96:
	;
	v417 = v403 + int32(base.Ui32(v400)>>(uint(int32(29))%32))&int32(4)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+16))
	if v418 != 0 {
		v400 = v400 << (uint(int32(1)) % 32)
		v403 = v418
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v459 = v457
	goto L100
L99:
	;
	v459 = v456
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_closedir[9])) = v459
	goto L2
}
func F_codearith(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v47 float64
	_ = v47
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v11 != int32(5) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v70 = F_luaK_exp2RK(m, l0, l2)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L22
	} else {
		goto L27
	}
L2:
	;
	v64 = F_luaK_exp2RK(m, l0, l3)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L22
	} else {
		goto L23
	}
L3:
	;
	switch l1 + int32(-18) {
	case 0, 2:
		v68 = int32(0)
		goto L1
	default:
		goto L2
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v14 != int32(-1) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v17 != int32(-1) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v20 != int32(5) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v23 != int32(-1) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v26 != int32(-1) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v30 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	switch l1 + int32(-12) {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L16
	case 3:
		goto L15
	case 4:
		goto L14
	case 5:
		goto L13
	case 6:
		goto L12
	default:
		v51 = float64(0)
		goto L10
	case 8:
		v68 = int32(0)
		goto L1
	}
L10:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l2)+8)) = v51
	return
L11:
	;
	if base.F64_ne(v49, v49) != 0 {
		goto L3
	} else {
		goto L21
	}
L12:
	;
	v49 = base.F64_neg(v30)
	goto L11
L13:
	;
	v47 = F_pow(m, v30, v29)
	mBase = m.M
	v49 = v47
	goto L11
L14:
	;
	if base.F64_eq(v29, float64(0)) != 0 {
		goto L2
	} else {
		goto L20
	}
L15:
	;
	if base.F64_eq(v29, float64(0)) != 0 {
		goto L2
	} else {
		goto L19
	}
L16:
	;
	v49 = base.F64_mul(v30, v29)
	goto L11
L17:
	;
	v49 = base.F64_sub(v30, v29)
	goto L11
L18:
	;
	v49 = base.F64_add(v30, v29)
	goto L11
L19:
	;
	v49 = base.F64_div(v30, v29)
	goto L11
L20:
	;
	v49 = base.F64_sub(v30, base.F64_mul(base.F64_floor(base.F64_div(v30, v29)), v29))
	goto L11
L21:
	;
	v51 = v49
	goto L10
L22:
	;
	return
L23:
	;
	v68 = v64
	goto L1
L24:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	v131 = F_luaK_code(m, l0, v68<<(uint(int32(14))%32)|v70<<(uint(int32(23))%32)|l1, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L22
	} else {
		goto L43
	}
L25:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v117 + int32(-1)
	goto L24
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v94 != int32(12) {
		goto L36
	} else {
		goto L37
	}
L27:
	;
	if v70 <= v68 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v73 != int32(12) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v86 != int32(12) {
		goto L24
	} else {
		goto L33
	}
L30:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v76&int32(256) != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v76 < v79 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v81 + int32(-1)
	goto L29
L33:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v89&int32(256) != 0 {
		goto L24
	} else {
		goto L34
	}
L34:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v92 <= v89 {
		goto L25
	} else {
		goto L35
	}
L35:
	;
	goto L24
L36:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v107 != int32(12) {
		goto L24
	} else {
		goto L40
	}
L37:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v97&int32(256) != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v97 < v100 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v102 + int32(-1)
	goto L36
L40:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v110&int32(256) != 0 {
		goto L24
	} else {
		goto L41
	}
L41:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	if v110 < v113 {
		goto L24
	} else {
		goto L42
	}
L42:
	;
	goto L25
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v131
	return
}
func F_commandlogPushCurrentCommand(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+57)))
	if v4&int32(16) != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
		if v8 != 0 {
			v11 = v8
			v12 = int32(188)
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v11 = v9
			v12 = int32(24)
		}
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0+v12)))
		v15 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+120)))
		F_commandlogPushEntryIfNeeded(m, l0, v11, v14, v15, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+240))
			F_commandlogPushEntryIfNeeded(m, l0, v11, v14, v19, int32(1))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
				F_commandlogPushEntryIfNeeded(m, l0, v11, v14, v23, int32(2))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_commandlogPushEntryIfNeeded(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v235 int32
	_ = v235
	var v236 int64
	_ = v236
	var v241 int32
	_ = v241
	var v245 int64
	_ = v245
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	v22 = m.G0
	v24 = v22 - int32(32)
	m.G0 = v24
	v28 = l4 << (uint(int32(5)) % 32)
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_commandlogPushEntryIfNeeded[0])))
	if v30 < int64(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v24 + int32(32)
	return
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_commandlogPushEntryIfNeeded[1])))
	if v35 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l3 < v30 {
		v285 = v35
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_commandlogPushEntryIfNeeded[2])))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+20))
	if base.Ui32(v299) <= base.Ui32(v285) {
		goto L1
	} else {
		goto L65
	}
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_commandlogPushEntryIfNeeded[2])))
	v41 = F_valkey_malloc(m, int32(40))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v43 = int32(32)
	if l2 < v43 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v46 = l2
	goto L10
L9:
	;
	v46 = v43
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v46
	v50 = F_valkey_malloc(m, v46<<(uint(int32(2))%32))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v50
	if l2 < int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v235 = int32(0)
	v236 = F___time(m, v235)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v41)+16)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v41)+24)) = v236
	v241 = l4 << (uint(int32(5)) % 32)
	v245 = *(*int64)(unsafe.Add(mBase, uint32(v241)+uint32(_c_F_commandlogPushEntryIfNeeded[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v241)+uint32(_c_F_commandlogPushEntryIfNeeded[3]))) = v245 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = v245
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_commandlogPushEntryIfNeeded[4]))
	goto L54
L13:
	;
	v59 = v46 + int32(-1)
	v68 = int32(0)
	goto L14
L14:
	;
	if l2 < int32(33) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L12
L16:
	;
	v212 = v68 + int32(1)
	if v212 != v46 {
		v68 = v212
		goto L14
	} else {
		goto L52
	}
L17:
	;
	if int32(1) <= v68 {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	if v68 != v59 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v88 = F_sdsempty(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = l2 - v46 + int32(1)
	v93 = F_sdscatprintf(m, v88, int32(_a_F_commandlogPushEntryIfNeeded_0), v24)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v95 = F_createObject(m, int32(0), v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50+v59<<(uint(int32(2))%32)))) = v95
	goto L16
L23:
	;
	v120 = v68 << (uint(int32(2)) % 32)
	v121 = l1 + v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v123&int32(15) != 0 {
		v195 = v122
		goto L30
	} else {
		goto L31
	}
L24:
	;
	if v110 == int32(0) {
		goto L23
	} else {
		goto L29
	}
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if base.Ui32(v68) < base.Ui32(int32(32)) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v110 = int32(0)
	goto L24
L27:
	;
	v110 = int32(base.Ui32(v101)>>(uint(v68)%32)) & int32(1)
	goto L24
L28:
	;
	v110 = v101 & int32(1)
	goto L24
L29:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_commandlogPushEntryIfNeeded[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v50+v68<<(uint(int32(2))%32)))) = v117
	goto L16
L30:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	if base.Ui32(v197) < base.Ui32(int32(-8)) {
		goto L49
	} else {
		goto L50
	}
L31:
	;
	switch int32(base.Ui32(v123)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		goto L32
	default:
		v195 = v122
		goto L30
	}
L32:
	;
	v130 = F_objectGetVal(m, v122)
	mBase = m.M
	v131 = int32(-1)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+v131))))
	switch v133&int32(7) + v131 {
	case 0:
		goto L38
	case 1:
		goto L37
	case 2:
		goto L36
	case 3:
		goto L35
	default:
		goto L33
	}
L33:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v195 = v194
	goto L30
L34:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if base.Ui32(v150) < base.Ui32(int32(129)) {
		v195 = v151
		goto L30
	} else {
		goto L39
	}
L35:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v130+int32(-17))))
	v150 = v149
	goto L34
L36:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v130+int32(-9))))
	v150 = v146
	goto L34
L37:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130+int32(-5)))))
	v150 = v143
	goto L34
L38:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+int32(-3)))))
	v150 = v140
	goto L34
L39:
	;
	v154 = F_objectGetVal(m, v151)
	mBase = m.M
	v156 = F_sdsnewlen(m, v154, int32(128))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v160 = F_objectGetVal(m, v159)
	mBase = m.M
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+int32(-1)))))
	switch v163 & int32(7) {
	case 0:
		goto L46
	case 1:
		goto L45
	case 2:
		goto L44
	case 3:
		goto L43
	case 4:
		goto L42
	default:
		v180 = int32(0)
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v180 + int32(-128)
	v189 = F_sdscatprintf(m, v156, int32(_a_F_commandlogPushEntryIfNeeded_1), v24+int32(16))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L47
	}
L42:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v160+int32(-17))))
	v180 = v179
	goto L41
L43:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v160+int32(-9))))
	v180 = v176
	goto L41
L44:
	;
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160+int32(-5)))))
	v180 = v173
	goto L41
L45:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+int32(-3)))))
	v180 = v170
	goto L41
L46:
	;
	v180 = int32(base.Ui32(v163) >> (uint(int32(3)) % 32))
	goto L41
L47:
	;
	v191 = F_createObject(m, int32(0), v189)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50+v120))) = v191
	goto L16
L49:
	;
	v203 = F_dupStringObject(m, v195)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L6
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50+v120))) = v195
	goto L16
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50+v120))) = v203
	goto L16
L52:
	;
	goto L15
L53:
	;
	v259 = F_getClientPeerId(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L6
	} else {
		goto L57
	}
L54:
	;
	if base.B2i32(v251 != v235) == int32(0) {
		v258 = l0
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v256 = F_scriptGetCaller(m)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	v258 = v256
	goto L53
L57:
	;
	v261 = F_sdsnew(m, v259)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v261
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v258)+348))
	if v264 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v272
	v274 = F_listAddNodeHead(m, v39, v41)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L6
	} else {
		goto L64
	}
L60:
	;
	v270 = F_sdsempty(m)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L6
	} else {
		goto L63
	}
L61:
	;
	v267 = F_objectGetVal(m, v264)
	mBase = m.M
	v268 = F_sdsnew(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v272 = v268
	goto L59
L63:
	;
	v272 = v270
	goto L59
L64:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_commandlogPushEntryIfNeeded[1])))
	v285 = v276
	goto L4
L65:
	;
	v303 = v298
	goto L66
L66:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	F_listDelNode(m, v303, v322)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L6
	} else {
		goto L68
	}
L67:
	;
	goto L1
L68:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_commandlogPushEntryIfNeeded[2])))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+20))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_commandlogPushEntryIfNeeded[1])))
	if base.Ui32(v327) < base.Ui32(v326) {
		v303 = v325
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
}
func F_compactBucketChain(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	v10 = l0 + l2<<(uint(int32(2))%32)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v14 = v11 + l1<<(uint(int32(6))%32)
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
	if v15&int32(1) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = v10 + int32(32)
	v26 = v14
	v27 = v15
	goto L3
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+60))
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29))))
	if v30&int32(8191) != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L1
L5:
	;
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107))))
	if v110&int32(1) != 0 {
		v26 = v107
		v27 = v110
		goto L3
	} else {
		goto L26
	}
L6:
	;
	if v30&int32(1) != 0 {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v35
	F_valkey_free(m, v29)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	if v40 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v46 + int32(-1)
	v107 = v26
	goto L5
L11:
	;
	m.T0[v40].(func(*base.Module, int32, int32))(m, l0, int32(-64))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	if base.Ui32(int32(10)) < base.Ui32(base.I32_popcnt(int32(base.Ui32(v27)>>(uint(int32(1))%32))&int32(4095))) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v55 = int32(base.Ui32(v30)>>(uint(int32(1))%32)) & int32(4095)
	if v55&(v55+int32(-1)) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_pruneLastBucket(m, l0, v26, v29, l2)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	return
L17:
	;
	v107 = v29
	goto L5
L18:
	;
	v71 = int32(0)
	v75 = v27
	goto L19
L19:
	;
	v77 = int32(1)
	if int32(base.Ui32(int32(base.Ui32(v75)>>(uint(v77)%32))&int32(4095))>>(uint(v71)%32))&v77 != 0 {
		v91 = v75
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L17
L21:
	;
	v93 = v71 + int32(1)
	if v93 != int32(11) {
		v71 = v93
		v75 = v91
		goto L19
	} else {
		goto L25
	}
L22:
	;
	F_fillBucketHole(m, l0, v26, v71, l2)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26))))
	if v86&int32(1) == int32(0) {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v91 = v86
	goto L21
L25:
	;
	goto L20
L26:
	;
	goto L4
}
func F_computeDatasetDigest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v24 int64
	_ = v24
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
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
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v204 int32
	_ = v204
	var v207 int64
	_ = v207
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
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
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
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
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v505 int32
	_ = v505
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(160)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(16)))) = v2
	v24 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(8)))) = v24
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v24
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_computeDatasetDigest[0]))
	if v29 < int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(160)
	return
L2:
	;
	v38 = v2
	goto L3
L3:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_computeDatasetDigest[1]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v38<<(uint(int32(2))%32))))
	if v53 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v519 = v38 + int32(1)
	v521 = *(*int32)(unsafe.Add(mBase, _c_F_computeDatasetDigest[0]))
	if v519 < v521 {
		v38 = v519
		goto L3
	} else {
		goto L44
	}
L6:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v57 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v67 == int64(0) {
		goto L5
	} else {
		goto L12
	}
L8:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v62 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v56)+40))
	v67 = v60
	goto L7
L10:
	;
	v64 = F_hashtableSize(m, v62)
	mBase = m.M
	v67 = base.I64_extend_i32_u(v64)
	goto L7
L11:
	;
	v67 = int64(0)
	goto L7
L12:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v72 = F_kvstoreIteratorInit(m, v70, int32(3))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	v74 = F___bswap_32_1(m, v38)
	mBase = m.M
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v74
	v80 = m.G0
	v81 = int32(112)
	v82 = v80 - v81
	m.G0 = v82
	v85 = v82 + int32(20)
	F_SHA1Init(m, v85)
	mBase = m.M
	F_SHA1Update(m, v85, v15+int32(12), int32(4))
	mBase = m.M
	F_SHA1Final(m, v82, v85)
	mBase = m.M
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v95 = v93 ^ v94
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v95)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	v99 = v97 ^ v98
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v99)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+2)))
	v103 = v101 ^ v102
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v103)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+3)))
	v107 = v105 ^ v106
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v107)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)))
	v111 = v109 ^ v110
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v111)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+5)))
	v115 = v113 ^ v114
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v115)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+6)))
	v119 = v117 ^ v118
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v119)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+7)))
	v123 = v121 ^ v122
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)) = uint8(v123)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+8)))
	v127 = v125 ^ v126
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v127)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+9)))
	v131 = v129 ^ v130
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v131)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+10)))
	v135 = v133 ^ v134
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v135)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+11)))
	v139 = v137 ^ v138
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v139)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+12)))
	v143 = v141 ^ v142
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v143)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+13)))
	v147 = v145 ^ v146
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v147)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+14)))
	v151 = v149 ^ v150
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v151)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+15)))
	v155 = v153 ^ v154
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v155)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+16)))
	v159 = v157 ^ v158
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v159)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+17)))
	v163 = v161 ^ v162
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v163)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+18)))
	v167 = v165 ^ v166
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v167)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+19)))
	v171 = v169 ^ v170
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v171)
	m.G0 = v82 + v81
	goto L16
L16:
	;
	v177 = v15 + int32(68)
	F_SHA1Init(m, v177)
	mBase = m.M
	F_SHA1Update(m, v177, l0, int32(20))
	mBase = m.M
	F_SHA1Final(m, l0, v177)
	mBase = m.M
	v188 = F_kvstoreIteratorNext(m, v72, v15+int32(8))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L13
	} else {
		goto L18
	}
L17:
	;
	F_kvstoreIteratorRelease(m, v72)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L13
	} else {
		goto L43
	}
L18:
	;
	if v188 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	goto L20
L20:
	;
	v204 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(32)))) = v204
	v207 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v207
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v207
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	if v216&int32(2) == v204 {
		v236 = v204
		goto L29
	} else {
		goto L30
	}
L21:
	;
	goto L17
L22:
	;
	v257 = F_createStringObject_1(m, v236, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L13
	} else {
		goto L31
	}
L23:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v236+int32(-17))))
	v256 = v255
	goto L22
L24:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v236+int32(-9))))
	v256 = v252
	goto L22
L25:
	;
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236+int32(-5)))))
	v256 = v249
	goto L22
L26:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236+int32(-3)))))
	v256 = v246
	goto L22
L27:
	;
	v256 = int32(base.Ui32(v239) >> (uint(int32(3)) % 32))
	goto L22
L28:
	;
	v238 = v236 + int32(-1)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	switch v239 & int32(7) {
	case 0:
		goto L27
	case 1:
		goto L26
	case 2:
		goto L25
	case 3:
		goto L24
	case 4:
		goto L23
	default:
		v256 = v204
		goto L22
	}
L29:
	;
	goto L28
L30:
	;
	v230 = v212 + (v216&int32(4) ^ int32(12)) + v216<<(uint(int32(3))%32)&int32(8)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	v236 = v230 + v231 + int32(1)
	goto L29
L31:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	switch v259 & int32(7) {
	case 0:
		goto L37
	case 1:
		goto L36
	case 2:
		goto L35
	case 3:
		goto L34
	case 4:
		goto L33
	default:
		v276 = v204
		goto L32
	}
L32:
	;
	v278 = v15 + int32(68)
	F_SHA1Init(m, v278)
	mBase = m.M
	F_SHA1Update(m, v278, v236, v276)
	mBase = m.M
	F_SHA1Final(m, v15+int32(48), v278)
	mBase = m.M
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+16)))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+48)))
	v290 = v288 ^ v289
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+16)) = uint8(v290)
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+17)))
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+49)))
	v294 = v292 ^ v293
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+17)) = uint8(v294)
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+18)))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+50)))
	v298 = v296 ^ v297
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+18)) = uint8(v298)
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+19)))
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+51)))
	v302 = v300 ^ v301
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+19)) = uint8(v302)
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+52)))
	v306 = v304 ^ v305
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)) = uint8(v306)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+21)))
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+53)))
	v310 = v308 ^ v309
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+21)) = uint8(v310)
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+54)))
	v314 = v312 ^ v313
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)) = uint8(v314)
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+23)))
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+55)))
	v318 = v316 ^ v317
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+23)) = uint8(v318)
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+24)))
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+56)))
	v322 = v320 ^ v321
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+24)) = uint8(v322)
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+25)))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+57)))
	v326 = v324 ^ v325
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+25)) = uint8(v326)
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+58)))
	v330 = v328 ^ v329
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)) = uint8(v330)
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)))
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+59)))
	v334 = v332 ^ v333
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)) = uint8(v334)
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+28)))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+60)))
	v338 = v336 ^ v337
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+28)) = uint8(v338)
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+29)))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+61)))
	v342 = v340 ^ v341
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+29)) = uint8(v342)
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+30)))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+62)))
	v346 = v344 ^ v345
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+30)) = uint8(v346)
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+31)))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+63)))
	v350 = v348 ^ v349
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+31)) = uint8(v350)
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+32)))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+64)))
	v354 = v352 ^ v353
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+32)) = uint8(v354)
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+33)))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+65)))
	v358 = v356 ^ v357
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+33)) = uint8(v358)
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+34)))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+66)))
	v362 = v360 ^ v361
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+34)) = uint8(v362)
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+35)))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+67)))
	v366 = v364 ^ v365
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+35)) = uint8(v366)
	F_SHA1Init(m, v278)
	mBase = m.M
	v374 = v15 + int32(16)
	F_SHA1Update(m, v278, v374, int32(20))
	mBase = m.M
	F_SHA1Final(m, v374, v278)
	mBase = m.M
	F_xorObjectDigest(m, v53, v257, v374, v212)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L13
	} else {
		goto L38
	}
L33:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v236+int32(-17))))
	v276 = v275
	goto L32
L34:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v236+int32(-9))))
	v276 = v272
	goto L32
L35:
	;
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236+int32(-5)))))
	v276 = v269
	goto L32
L36:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236+int32(-3)))))
	v276 = v266
	goto L32
L37:
	;
	v276 = int32(base.Ui32(v259) >> (uint(int32(3)) % 32))
	goto L32
L38:
	;
	v388 = int32(20)
	v390 = m.G0
	v391 = int32(112)
	v392 = v390 - v391
	m.G0 = v392
	v395 = v392 + v388
	F_SHA1Init(m, v395)
	mBase = m.M
	F_SHA1Update(m, v395, v15+int32(16), v388)
	mBase = m.M
	F_SHA1Final(m, v392, v395)
	mBase = m.M
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392))))
	v405 = v403 ^ v404
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v405)
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+1)))
	v409 = v407 ^ v408
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v409)
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+2)))
	v413 = v411 ^ v412
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v413)
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+3)))
	v417 = v415 ^ v416
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v417)
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+4)))
	v421 = v419 ^ v420
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v421)
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+5)))
	v425 = v423 ^ v424
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v425)
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+6)))
	v429 = v427 ^ v428
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v429)
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+7)))
	v433 = v431 ^ v432
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)) = uint8(v433)
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+8)))
	v437 = v435 ^ v436
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v437)
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+9)))
	v441 = v439 ^ v440
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v441)
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+10)))
	v445 = v443 ^ v444
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v445)
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)))
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+11)))
	v449 = v447 ^ v448
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v449)
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+12)))
	v453 = v451 ^ v452
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v453)
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+13)))
	v457 = v455 ^ v456
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v457)
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+14)))
	v461 = v459 ^ v460
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v461)
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+15)))
	v465 = v463 ^ v464
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v465)
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+16)))
	v469 = v467 ^ v468
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v469)
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+17)))
	v473 = v471 ^ v472
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v473)
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)))
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+18)))
	v477 = v475 ^ v476
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v477)
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+19)))
	v481 = v479 ^ v480
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v481)
	m.G0 = v392 + v391
	goto L39
L39:
	;
	F_decrRefCount(m, v257)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	v490 = F_kvstoreIteratorNext(m, v72, v15+int32(8))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L13
	} else {
		goto L41
	}
L41:
	;
	if v490 != 0 {
		goto L20
	} else {
		goto L42
	}
L42:
	;
	goto L21
L43:
	;
	goto L5
L44:
	;
	goto L4
}
func F_connectWithPrimary(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int64
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[0]))
	if v8 == int32(0) {
		v15 = F_connectionTypeTcp(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = v15
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
			v20 = m.T0[v19].(func(*base.Module) int32)(m)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[1])) = v20
				v23 = int32(_a_F_connectWithPrimary_0)
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[2]))
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[3]))
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[4]))
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[5]))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+56))
				v34 = m.T0[v33].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v20, v24, v26, v28, v30, int32(970))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					if v34 != int32(-1) {
						v64 = int32(_a_F_connectWithPrimary_0)
						v65 = int32(2)
						*(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[6])) = v65
						v69 = *(*int64)(unsafe.Add(mBase, _c_F_connectWithPrimary[7]))
						*(*int64)(unsafe.Add(mBase, _c_F_connectWithPrimary[8])) = v69
						v73 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[9]))
						if v65 < v73 {
							v82 = int32(0)
							m.G0 = v5 + int32(16)
							return v82
						} else {
							v76 = int32(0)
							F__serverLog(m, int32(2), int32(_a_F_connectWithPrimary_1), v76)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = v76
								m.G0 = v5 + int32(16)
								return v82
							}
						}
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[9]))
						if int32(3) < v39 {
							v55 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[1]))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
							m.T0[v57].(func(*base.Module, int32))(m, v55)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[1])) = int32(0)
								v82 = int32(-1)
								m.G0 = v5 + int32(16)
								return v82
							}
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[1]))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+88))
							v46 = m.T0[v45].(func(*base.Module, int32) int32)(m, v43)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = v46
								F__serverLog(m, int32(3), int32(_a_F_connectWithPrimary_2), v5)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[1]))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
									m.T0[v57].(func(*base.Module, int32))(m, v55)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[1])) = int32(0)
										v82 = int32(-1)
										m.G0 = v5 + int32(16)
										return v82
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v11 = F_connectionTypeTls(m)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v17 = v11
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
			v20 = m.T0[v19].(func(*base.Module) int32)(m)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[1])) = v20
				v23 = int32(_a_F_connectWithPrimary_0)
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[2]))
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[3]))
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[4]))
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[5]))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+56))
				v34 = m.T0[v33].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v20, v24, v26, v28, v30, int32(970))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					if v34 != int32(-1) {
						v64 = int32(_a_F_connectWithPrimary_0)
						v65 = int32(2)
						*(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[6])) = v65
						v69 = *(*int64)(unsafe.Add(mBase, _c_F_connectWithPrimary[7]))
						*(*int64)(unsafe.Add(mBase, _c_F_connectWithPrimary[8])) = v69
						v73 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[9]))
						if v65 < v73 {
							v82 = int32(0)
							m.G0 = v5 + int32(16)
							return v82
						} else {
							v76 = int32(0)
							F__serverLog(m, int32(2), int32(_a_F_connectWithPrimary_1), v76)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = v76
								m.G0 = v5 + int32(16)
								return v82
							}
						}
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[9]))
						if int32(3) < v39 {
							v55 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[1]))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
							m.T0[v57].(func(*base.Module, int32))(m, v55)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[1])) = int32(0)
								v82 = int32(-1)
								m.G0 = v5 + int32(16)
								return v82
							}
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[1]))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+88))
							v46 = m.T0[v45].(func(*base.Module, int32) int32)(m, v43)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = v46
								F__serverLog(m, int32(3), int32(_a_F_connectWithPrimary_2), v5)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[1]))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
									m.T0[v57].(func(*base.Module, int32))(m, v55)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_connectWithPrimary[1])) = int32(0)
										v82 = int32(-1)
										m.G0 = v5 + int32(16)
										return v82
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
func F_controloptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	switch l1 + int32(-32) {
	case 0:
		goto L1
	case 1:
		goto L3
	default:
		goto L2
	case 28:
		goto L4
	case 30:
		goto L5
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	v91 = m.G3
	v94 = F_lua_pushfstring(m, l0, v91+int32(_a_F_controloptions_0), v12)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L14
	} else {
		goto L22
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21))))
	if base.Ui32(int32(9)) < base.Ui32(v22+int32(-48)) {
		v71 = int32(8)
		goto L6
	} else {
		goto L7
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(1)
	goto L1
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L1
L6:
	;
	if v71 < int32(1) {
		goto L18
	} else {
		goto L19
	}
L7:
	;
	v33 = int32(0)
	v34 = v21
	v35 = v22
	goto L8
L8:
	;
	v38 = v33 * int32(10)
	if int32(214748364) < v33 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v71 = v60
	goto L6
L10:
	;
	v55 = v53 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v55
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v53))))
	v59 = int32(-48)
	v60 = v38 + v57 + v59
	v61 = int32(*(*int8)(unsafe.Add(mBase, uint32(v53)+1)))
	if base.Ui32(v61+v59) < base.Ui32(int32(10)) {
		v33 = v60
		v34 = v55
		v35 = v61
		goto L8
	} else {
		goto L16
	}
L11:
	;
	v46 = m.G3
	v50 = F_luaL_error(m, l0, v46+int32(_a_F_controloptions_1), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if v38 <= int32(-2147483601)-base.I32_extend8_s(v35) {
		v53 = v34
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	return
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v53 = v52
	goto L10
L16:
	;
	goto L9
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v71
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v71
	v81 = m.G3
	v86 = F_luaL_error(m, l0, v81+int32(_a_F_controloptions_2), v12+int32(16))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L14
	} else {
		goto L21
	}
L19:
	;
	if base.Ui32(base.I32_popcnt(v71)) < base.Ui32(int32(2)) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L17
L22:
	;
	v96 = F_luaL_argerror(m, l0, int32(1), v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	goto L1
}
func F_convert_ioctl_struct(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v301 int32
	_ = v301
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int64
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int64
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v476 int32
	_ = v476
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int64
	_ = v557
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int64
	_ = v572
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v654 int32
	_ = v654
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v683 int32
	_ = v683
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if l3&v19 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return
L2:
	;
	v23 = l0
	v24 = l1
	v25 = l2
	goto L4
L3:
	;
	v670 = v37 - v662
	if l3 != int32(1) {
		goto L92
	} else {
		goto L93
	}
L4:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)))
	if v37 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v130 = v23 + int32(20)
	v141 = m.G0
	v143 = v141 - int32(16)
	m.G0 = v143
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+9)))
	if l3&v145 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L7:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+11)))
	if v40 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v47 = int32(0)
	v56 = v47
	v57 = v47
	v60 = v47
	goto L10
L9:
	;
	v41 = int32(0)
	v662 = v41
	v663 = v41
	goto L3
L10:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(12)+v60))))
	v67 = v66 - v56
	v68 = v67 + v57
	v72 = (int32(0)-v68)&int32(7) + v68
	if l3 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v122 = v66 + int32(4)
	v124 = v72 + int32(8)
	v126 = v60 + int32(1)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+11)))
	if base.Ui32(v127) <= base.Ui32(v126) {
		v662 = v122
		v663 = v124
		goto L3
	} else {
		goto L33
	}
L13:
	;
	if v67 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	if v67 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v79 = int32(8)
	goto L20
L16:
	;
	goto L15
L17:
	;
	v77 = F__emscripten_memcpy_bulkmem(m, v24+v56, v25+v57, v67)
	mBase = m.M
	goto L16
L18:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+4)) = uint32(v87)
	v90 = int32(4)
	goto L23
L19:
	;
	goto L18
L20:
	;
	v85 = F__emscripten_memcpy_bulkmem(m, v17+v79, v25+v72, v79)
	mBase = m.M
	goto L19
L21:
	;
	goto L12
L22:
	;
	goto L21
L23:
	;
	v95 = F__emscripten_memcpy_bulkmem(m, v24+v66, v17+v90, v90)
	mBase = m.M
	goto L22
L24:
	;
	v103 = int32(4)
	goto L29
L25:
	;
	goto L24
L26:
	;
	v101 = F__emscripten_memcpy_bulkmem(m, v25+v57, v24+v56, v67)
	mBase = m.M
	goto L25
L27:
	;
	v111 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v111
	v114 = int32(8)
	goto L32
L28:
	;
	goto L27
L29:
	;
	v109 = F__emscripten_memcpy_bulkmem(m, v17+v103, v24+v66, v103)
	mBase = m.M
	goto L28
L30:
	;
	goto L12
L31:
	;
	goto L30
L32:
	;
	v119 = F__emscripten_memcpy_bulkmem(m, v25+v72, v17+v114, v114)
	mBase = m.M
	goto L31
L33:
	;
	v56 = v122
	v57 = v124
	v60 = v126
	goto L10
L34:
	;
	v301 = v23 + int32(40)
	v316 = m.G0
	v318 = v316 - int32(16)
	m.G0 = v318
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+9)))
	if l3&v320 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L35:
	;
	m.G0 = v143 + int32(16)
	goto L34
L36:
	;
	v149 = v130
	v150 = v24
	v151 = v25
	goto L38
L37:
	;
	v274 = v163 - v266
	if l3 != int32(1) {
		goto L51
	} else {
		goto L52
	}
L38:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+8)))
	if v163 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_convert_ioctl_struct(m, v149+int32(20), v150, v151, l3)
	mBase = m.M
	F_convert_ioctl_struct(m, v149+int32(40), v150+int32(4), v151+int32(8), l3)
	mBase = m.M
	v248 = v149 + int32(60)
	v251 = int32(72)
	F_convert_ioctl_struct(m, v248, v150+int32(68), v151+v251, l3)
	mBase = m.M
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+69)))
	if l3&v258 != 0 {
		v149 = v248
		v150 = v150 + v251
		v151 = v151 + int32(76)
		goto L38
	} else {
		goto L50
	}
L41:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+11)))
	if v166 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v173 = int32(0)
	v182 = v173
	v183 = v173
	v186 = v173
	goto L44
L43:
	;
	v167 = int32(0)
	v266 = v167
	v267 = v167
	goto L37
L44:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149+int32(12)+v186))))
	v193 = v192 - v182
	v194 = v193 + v183
	v198 = (int32(0)-v194)&int32(7) + v194
	if l3 != int32(1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v230 = v192 + int32(4)
	v232 = v198 + int32(8)
	v234 = v186 + int32(1)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+11)))
	if base.Ui32(v235) <= base.Ui32(v234) {
		v266 = v230
		v267 = v232
		goto L37
	} else {
		goto L49
	}
L47:
	;
	v216 = F___memcpy(m, v151+v183, v150+v182, v193)
	mBase = m.M
	v217 = int32(4)
	v221 = F___memcpy(m, v143+v217, v150+v192, v217)
	mBase = m.M
	v222 = int64(*(*int32)(unsafe.Add(mBase, uint32(v143)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v143)+8)) = v222
	v225 = int32(8)
	v228 = F___memcpy(m, v151+v198, v143+v225, v225)
	mBase = m.M
	goto L46
L48:
	;
	v201 = F___memcpy(m, v150+v182, v151+v183, v193)
	mBase = m.M
	v202 = int32(8)
	v206 = F___memcpy(m, v143+v202, v151+v198, v202)
	mBase = m.M
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v143)+4)) = uint32(v207)
	v210 = int32(4)
	v213 = F___memcpy(m, v150+v192, v143+v210, v210)
	mBase = m.M
	goto L46
L49:
	;
	v182 = v230
	v183 = v232
	v186 = v234
	goto L44
L50:
	;
	goto L35
L51:
	;
	v282 = F___memcpy(m, v151+v267, v150+v266, v274)
	mBase = m.M
	goto L35
L52:
	;
	v279 = F___memcpy(m, v150+v266, v151+v267, v274)
	mBase = m.M
	goto L35
L53:
	;
	v476 = v23 + int32(60)
	v491 = m.G0
	v493 = v491 - int32(16)
	m.G0 = v493
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+9)))
	if l3&v495 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L54:
	;
	m.G0 = v318 + int32(16)
	goto L53
L55:
	;
	v324 = v301
	v325 = v24 + int32(4)
	v326 = v25 + int32(8)
	goto L57
L56:
	;
	v449 = v338 - v441
	if l3 != int32(1) {
		goto L70
	} else {
		goto L71
	}
L57:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+8)))
	if v338 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_convert_ioctl_struct(m, v324+int32(20), v325, v326, l3)
	mBase = m.M
	F_convert_ioctl_struct(m, v324+int32(40), v325+int32(4), v326+int32(8), l3)
	mBase = m.M
	v423 = v324 + int32(60)
	v426 = int32(72)
	F_convert_ioctl_struct(m, v423, v325+int32(68), v326+v426, l3)
	mBase = m.M
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+69)))
	if l3&v433 != 0 {
		v324 = v423
		v325 = v325 + v426
		v326 = v326 + int32(76)
		goto L57
	} else {
		goto L69
	}
L60:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+11)))
	if v341 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v348 = int32(0)
	v357 = v348
	v358 = v348
	v361 = v348
	goto L63
L62:
	;
	v342 = int32(0)
	v441 = v342
	v442 = v342
	goto L56
L63:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324+int32(12)+v361))))
	v368 = v367 - v357
	v369 = v368 + v358
	v373 = (int32(0)-v369)&int32(7) + v369
	if l3 != int32(1) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v405 = v367 + int32(4)
	v407 = v373 + int32(8)
	v409 = v361 + int32(1)
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+11)))
	if base.Ui32(v410) <= base.Ui32(v409) {
		v441 = v405
		v442 = v407
		goto L56
	} else {
		goto L68
	}
L66:
	;
	v391 = F___memcpy(m, v326+v358, v325+v357, v368)
	mBase = m.M
	v392 = int32(4)
	v396 = F___memcpy(m, v318+v392, v325+v367, v392)
	mBase = m.M
	v397 = int64(*(*int32)(unsafe.Add(mBase, uint32(v318)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v318)+8)) = v397
	v400 = int32(8)
	v403 = F___memcpy(m, v326+v373, v318+v400, v400)
	mBase = m.M
	goto L65
L67:
	;
	v376 = F___memcpy(m, v325+v357, v326+v358, v368)
	mBase = m.M
	v377 = int32(8)
	v381 = F___memcpy(m, v318+v377, v326+v373, v377)
	mBase = m.M
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v318)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v318)+4)) = uint32(v382)
	v385 = int32(4)
	v388 = F___memcpy(m, v325+v367, v318+v385, v385)
	mBase = m.M
	goto L65
L68:
	;
	v357 = v405
	v358 = v407
	v361 = v409
	goto L63
L69:
	;
	goto L54
L70:
	;
	v457 = F___memcpy(m, v326+v442, v325+v441, v449)
	mBase = m.M
	goto L54
L71:
	;
	v454 = F___memcpy(m, v325+v441, v326+v442, v449)
	mBase = m.M
	goto L54
L72:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+69)))
	if l3&v654 != 0 {
		v23 = v476
		v24 = v24 + int32(72)
		v25 = v25 + int32(76)
		goto L4
	} else {
		goto L91
	}
L73:
	;
	m.G0 = v493 + int32(16)
	goto L72
L74:
	;
	v499 = v476
	v500 = v24 + int32(68)
	v501 = v25 + int32(72)
	goto L76
L75:
	;
	v624 = v513 - v616
	if l3 != int32(1) {
		goto L89
	} else {
		goto L90
	}
L76:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+8)))
	if v513 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	F_convert_ioctl_struct(m, v499+int32(20), v500, v501, l3)
	mBase = m.M
	F_convert_ioctl_struct(m, v499+int32(40), v500+int32(4), v501+int32(8), l3)
	mBase = m.M
	v598 = v499 + int32(60)
	v601 = int32(72)
	F_convert_ioctl_struct(m, v598, v500+int32(68), v501+v601, l3)
	mBase = m.M
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+69)))
	if l3&v608 != 0 {
		v499 = v598
		v500 = v500 + v601
		v501 = v501 + int32(76)
		goto L76
	} else {
		goto L88
	}
L79:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+11)))
	if v516 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v523 = int32(0)
	v532 = v523
	v533 = v523
	v536 = v523
	goto L82
L81:
	;
	v517 = int32(0)
	v616 = v517
	v617 = v517
	goto L75
L82:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499+int32(12)+v536))))
	v543 = v542 - v532
	v544 = v543 + v533
	v548 = (int32(0)-v544)&int32(7) + v544
	if l3 != int32(1) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v580 = v542 + int32(4)
	v582 = v548 + int32(8)
	v584 = v536 + int32(1)
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+11)))
	if base.Ui32(v585) <= base.Ui32(v584) {
		v616 = v580
		v617 = v582
		goto L75
	} else {
		goto L87
	}
L85:
	;
	v566 = F___memcpy(m, v501+v533, v500+v532, v543)
	mBase = m.M
	v567 = int32(4)
	v571 = F___memcpy(m, v493+v567, v500+v542, v567)
	mBase = m.M
	v572 = int64(*(*int32)(unsafe.Add(mBase, uint32(v493)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v493)+8)) = v572
	v575 = int32(8)
	v578 = F___memcpy(m, v501+v548, v493+v575, v575)
	mBase = m.M
	goto L84
L86:
	;
	v551 = F___memcpy(m, v500+v532, v501+v533, v543)
	mBase = m.M
	v552 = int32(8)
	v556 = F___memcpy(m, v493+v552, v501+v548, v552)
	mBase = m.M
	v557 = *(*int64)(unsafe.Add(mBase, uint32(v493)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v493)+4)) = uint32(v557)
	v560 = int32(4)
	v563 = F___memcpy(m, v500+v542, v493+v560, v560)
	mBase = m.M
	goto L84
L87:
	;
	v532 = v580
	v533 = v582
	v536 = v584
	goto L82
L88:
	;
	goto L73
L89:
	;
	v632 = F___memcpy(m, v501+v617, v500+v616, v624)
	mBase = m.M
	goto L73
L90:
	;
	v629 = F___memcpy(m, v500+v616, v501+v617, v624)
	mBase = m.M
	goto L73
L91:
	;
	goto L1
L92:
	;
	if v670 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L93:
	;
	if v670 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L1
L95:
	;
	goto L94
L96:
	;
	v677 = F__emscripten_memcpy_bulkmem(m, v24+v662, v25+v663, v670)
	mBase = m.M
	goto L95
L97:
	;
	goto L1
L98:
	;
	goto L97
L99:
	;
	v683 = F__emscripten_memcpy_bulkmem(m, v25+v663, v24+v662, v670)
	mBase = m.M
	goto L98
}
func F_copysignl(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = l1
	v7 = int64(48)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(l4)>>(uint(v7)%64)))&int32(32768)|base.I32_wrap_i64(int64(base.Ui64(l2&int64(9223090561878065152))>>(uint(v7)%64))))<<(uint(v7)%64) | l2&int64(281474976710655)
	return
}
func F_crc64(m *base.Module, l0 int64, l1 int32, l2 int64) int64 {
	var v4 int64
	_ = v4
	v4 = m.Env.Vkmem_crc64(m, l0, l1, l2)
	return v4
}
func F_crcspeed64little_init(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v97 int64
	_ = v97
	var v100 int64
	_ = v100
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v119 int64
	_ = v119
	var v122 int64
	_ = v122
	var v130 int64
	_ = v130
	var v136 int32
	_ = v136
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v18 = v3
	for {
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v18)
		v22 = int32(3)
		v25 = int64(0)
		v27 = v10 + int32(15)
		v28 = int64(1)
		v29 = m.T0[l0].(func(*base.Module, int64, int32, int64) int64)(m, v25, v27, v28)
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, uint32(l1+v18<<(uint(v22)%32)))) = v29
		v32 = v18 | int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v32)
		v41 = m.T0[l0].(func(*base.Module, int64, int32, int64) int64)(m, v25, v27, v28)
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, uint32(l1+v32<<(uint(v22)%32)))) = v41
		v44 = v18 + int32(2)
		if v44 != int32(256) {
			v18 = v44
			continue
		} else {
			break
		}
		break
	}
	v50 = v3
	for {
		v54 = int32(3)
		v56 = l1 + v50<<(uint(v54)%32)
		v57 = *(*int64)(unsafe.Add(mBase, uint32(v56)))
		v59 = int32(255)
		v64 = *(*int64)(unsafe.Add(mBase, uint32(l1+base.I32_wrap_i64(v57)&v59<<(uint(v54)%32))))
		v65 = int64(8)
		v67 = v64 ^ int64(base.Ui64(v57)>>(uint(v65)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v56)+2048)) = v67
		v75 = *(*int64)(unsafe.Add(mBase, uint32(l1+base.I32_wrap_i64(v67)&v59<<(uint(v54)%32))))
		v78 = v75 ^ int64(base.Ui64(v67)>>(uint(v65)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v56)+uint32(_c_F_crcspeed64little_init[0]))) = v78
		v86 = *(*int64)(unsafe.Add(mBase, uint32(l1+base.I32_wrap_i64(v78)&v59<<(uint(v54)%32))))
		v89 = v86 ^ int64(base.Ui64(v78)>>(uint(v65)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v56)+uint32(_c_F_crcspeed64little_init[1]))) = v89
		v97 = *(*int64)(unsafe.Add(mBase, uint32(l1+base.I32_wrap_i64(v89)&v59<<(uint(v54)%32))))
		v100 = v97 ^ int64(base.Ui64(v89)>>(uint(v65)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v56)+uint32(_c_F_crcspeed64little_init[2]))) = v100
		v108 = *(*int64)(unsafe.Add(mBase, uint32(l1+base.I32_wrap_i64(v100)&v59<<(uint(v54)%32))))
		v111 = v108 ^ int64(base.Ui64(v100)>>(uint(v65)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v56)+uint32(_c_F_crcspeed64little_init[3]))) = v111
		v119 = *(*int64)(unsafe.Add(mBase, uint32(l1+base.I32_wrap_i64(v111)&v59<<(uint(v54)%32))))
		v122 = v119 ^ int64(base.Ui64(v111)>>(uint(v65)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v56)+uint32(_c_F_crcspeed64little_init[4]))) = v122
		v130 = *(*int64)(unsafe.Add(mBase, uint32(l1+base.I32_wrap_i64(v122)&v59<<(uint(v54)%32))))
		*(*int64)(unsafe.Add(mBase, uint32(v56)+uint32(_c_F_crcspeed64little_init[5]))) = v130 ^ int64(base.Ui64(v122)>>(uint(v65)%64))
		v136 = v50 + int32(1)
		if v136 != int32(256) {
			v50 = v136
			continue
		} else {
			break
		}
		break
	}
	m.G0 = v10 + int32(16)
	return
}
func F_createArrayObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = m.G4
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = m.T0[v11].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(48))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v61 = int32(0)
			return v61
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v6
			if l1 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l1
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v41 == int32(0) {
					v61 = v12
					return v61
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
					if base.Ui32(v45+int32(-9)) < base.Ui32(int32(4)) {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v52+v53<<(uint(int32(2))%32)))) = v12
						v61 = v12
						return v61
					} else {
						if v45 != int32(2) {
							v63 = m.G3
							m.Env.X__assert_fail(m, v63+int32(_a_F_createArrayObject_0), v63+int32(_a_F_createArrayObject_1), int32(191), v63+int32(_a_F_createArrayObject_2))
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v52+v53<<(uint(int32(2))%32)))) = v12
							v61 = v12
							return v61
						}
					}
				}
			} else {
				if base.Ui32(l1) < base.Ui32(int32(1073741824)) {
					v30 = m.G4
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					v32 = m.T0[v31].(func(*base.Module, int32, int32) int32)(m, l1, int32(4))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v32
						if v32 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l1
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v41 == int32(0) {
								v61 = v12
								return v61
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
								if base.Ui32(v45+int32(-9)) < base.Ui32(int32(4)) {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v52+v53<<(uint(int32(2))%32)))) = v12
									v61 = v12
									return v61
								} else {
									if v45 != int32(2) {
										v63 = m.G3
										m.Env.X__assert_fail(m, v63+int32(_a_F_createArrayObject_0), v63+int32(_a_F_createArrayObject_1), int32(191), v63+int32(_a_F_createArrayObject_2))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
										v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v52+v53<<(uint(int32(2))%32)))) = v12
										v61 = v12
										return v61
									}
								}
							}
						} else {
							F_freeReplyObject(m, v12)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								return int32(0)
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = int32(0)
					F_freeReplyObject(m, v12)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				}
			}
		}
	}
}
func F_createBoolObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	v7 = m.G4
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = m.T0[v8].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(48))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			return v9
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(8)
			v17 = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = base.I64_extend_i32_u(base.B2i32(l1 != v17))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v21 == v17 {
				return v9
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				if base.Ui32(v25+int32(-9)) < base.Ui32(int32(4)) {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v32+v33<<(uint(int32(2))%32)))) = v9
					return v9
				} else {
					if v25 != int32(2) {
						v41 = m.G3
						m.Env.X__assert_fail(m, v41+int32(_a_F_createBoolObject_0), v41+int32(_a_F_createBoolObject_1), int32(290), v41+int32(_a_F_createBoolObject_2))
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v32+v33<<(uint(int32(2))%32)))) = v9
						return v9
					}
				}
			}
		}
	}
}
func F_createIntegerObject(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	v8 = m.G4
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v10 = m.T0[v9].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(48))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			return v10
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(3)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v19 == int32(0) {
				return v10
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				if base.Ui32(v23+int32(-9)) < base.Ui32(int32(4)) {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(2))%32)))) = v10
					return v10
				} else {
					if v23 != int32(2) {
						v39 = m.G3
						m.Env.X__assert_fail(m, v39+int32(_a_F_createIntegerObject_0), v39+int32(_a_F_createIntegerObject_1), int32(212), v39+int32(_a_F_createIntegerObject_2))
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(2))%32)))) = v10
						return v10
					}
				}
			}
		}
	}
}
func F_createPidFile(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_createPidFile[0]))
	if v8 != 0 {
		v14 = v8
		v16 = F_fopen(m, v14, int32(_a_F_createPidFile_0))
		mBase = m.M
		if v16 == int32(0) {
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_createPidFile[1]))
			if int32(3) < v29 {
				m.G0 = v5 + int32(32)
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_createPidFile[2]))
				v34 = F___strerror_l(m, v33, v33)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v34
				F__serverLog(m, int32(3), int32(_a_F_createPidFile_1), v5)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					m.G0 = v5 + int32(32)
					return
				}
			}
		} else {
			v19 = F___syscall_getpid(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v19
			v24 = F_fiprintf(m, v16, int32(_a_F_createPidFile_2), v5+int32(16))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v26 = F_fclose(m, v16)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					m.G0 = v5 + int32(32)
					return
				}
			}
		}
	} else {
		v11 = F_zstrdup(m, int32(_a_F_createPidFile_3))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_createPidFile[0])) = v11
			v14 = v11
			v16 = F_fopen(m, v14, int32(_a_F_createPidFile_0))
			mBase = m.M
			if v16 == int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_createPidFile[1]))
				if int32(3) < v29 {
					m.G0 = v5 + int32(32)
					return
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_createPidFile[2]))
					v34 = F___strerror_l(m, v33, v33)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = v34
					F__serverLog(m, int32(3), int32(_a_F_createPidFile_1), v5)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						m.G0 = v5 + int32(32)
						return
					}
				}
			} else {
				v19 = F___syscall_getpid(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v19
				v24 = F_fiprintf(m, v16, int32(_a_F_createPidFile_2), v5+int32(16))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = F_fclose(m, v16)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						m.G0 = v5 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_createSharedObjects(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
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
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int64
	_ = v316
	var v320 int32
	_ = v320
	var v330 int32
	_ = v330
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int64
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int64
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
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
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v922 int32
	_ = v922
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	v6 = m.G0
	v8 = v6 - int32(144)
	m.G0 = v8
	v13 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_0), int32(5))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = F_makeObjectShared(m, v13)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[0])) = v15
	v21 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_1), int32(6))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v23 = F_makeObjectShared(m, v21)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[1])) = v23
	v29 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_2), int32(4))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v31 = F_makeObjectShared(m, v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[2])) = v31
	v37 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_3), int32(4))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v39 = F_makeObjectShared(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[3])) = v39
	v45 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_4), int32(4))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v47 = F_makeObjectShared(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[4])) = v47
	v53 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_5), int32(7))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v55 = F_makeObjectShared(m, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[5])) = v55
	v61 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_6), int32(9))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v63 = F_makeObjectShared(m, v61)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[6])) = v63
	v69 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_7), int32(15))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v71 = F_makeObjectShared(m, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[7])) = v71
	v77 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_8), int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v79 = F_makeObjectShared(m, v77)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[8])) = v79
	v85 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_9), int32(1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v87 = F_makeObjectShared(m, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[9])) = v87
	v93 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_10), int32(68))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v95 = F_makeObjectShared(m, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[10])) = v95
	v101 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_11), int32(6))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v103 = F_makeObjectShared(m, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[11])) = v103
	v109 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_12), int32(18))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v111 = F_makeObjectShared(m, v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[12])) = v111
	v117 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_13), int32(19))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v119 = F_makeObjectShared(m, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[13])) = v119
	v125 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_14), int32(50))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v127 = F_makeObjectShared(m, v125)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[14])) = v127
	v133 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_15), int32(25))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v135 = F_makeObjectShared(m, v133)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[15])) = v135
	v141 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_16), int32(31))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v143 = F_makeObjectShared(m, v141)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[16])) = v143
	F_createSharedObjectsForCompat(m, int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_createSharedObjectsForCompat(m, int32(1))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v152 = int32(0)
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[17]))
	v156 = v154 << (uint(int32(2)) % 32)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)+uint32(_c_F_createSharedObjects[18])))
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[19])) = v159
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v156)+uint32(_c_F_createSharedObjects[20])))
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[21])) = v164
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v156)+uint32(_c_F_createSharedObjects[22])))
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[23])) = v169
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v156)+uint32(_c_F_createSharedObjects[24])))
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[25])) = v174
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v156)+uint32(_c_F_createSharedObjects[26])))
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[27])) = v179
	v184 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_17), int32(83))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v186 = F_makeObjectShared(m, v184)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[28])) = v186
	v192 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_18), int32(56))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v194 = F_makeObjectShared(m, v192)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[29])) = v194
	v200 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_19), int32(34))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v202 = F_makeObjectShared(m, v200)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[30])) = v202
	v208 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_20), int32(58))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v210 = F_makeObjectShared(m, v208)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[31])) = v210
	v216 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_21), int32(62))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v218 = F_makeObjectShared(m, v216)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[32])) = v218
	v224 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_22), int32(48))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v226 = F_makeObjectShared(m, v224)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[33])) = v226
	v232 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_23), int32(42))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v234 = F_makeObjectShared(m, v232)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v236 = int32(0)
	*(*int64)(unsafe.Add(mBase, _c_F_createSharedObjects[34])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[35])) = v234
	v244 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_24), int32(5))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v246 = F_makeObjectShared(m, v244)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[36])) = v246
	v251 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_25), int32(3))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v253 = F_makeObjectShared(m, v251)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v255 = int32(0)
	*(*int64)(unsafe.Add(mBase, _c_F_createSharedObjects[37])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[38])) = v253
	v263 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_26), int32(5))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v265 = F_makeObjectShared(m, v263)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[39])) = v265
	v270 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_25), int32(3))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v272 = F_makeObjectShared(m, v270)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v274 = int32(0)
	*(*int64)(unsafe.Add(mBase, _c_F_createSharedObjects[40])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[41])) = v272
	v282 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_4), int32(4))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v284 = F_makeObjectShared(m, v282)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[42])) = v284
	v289 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_27), int32(4))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v291 = F_makeObjectShared(m, v289)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v293 = int32(0)
	*(*int64)(unsafe.Add(mBase, _c_F_createSharedObjects[43])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[44])) = v291
	v301 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_4), int32(4))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v303 = F_makeObjectShared(m, v301)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[45])) = v303
	v309 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_28), int32(4))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v311 = F_makeObjectShared(m, v309)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[46])) = v311
	v316 = int64(0)
	goto L68
L68:
	;
	v320 = v8 + int32(80)
	if v316 <= int64(-1) {
		goto L74
	} else {
		goto L75
	}
L69:
	;
	v389 = int32(0)
	v393 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_29), int32(13))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L84
	}
L70:
	;
	v363 = F_sdsempty(m)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L79
	}
L71:
	;
	v362 = int32(0)
	goto L70
L73:
	;
	v343 = F_ull2string(m, v339, v340, v341)
	mBase = m.M
	if v343 == int32(0) {
		goto L71
	} else {
		goto L77
	}
L74:
	;
	goto L76
L75:
	;
	v339 = v320
	v340 = int32(64)
	v341 = v316
	v342 = int32(0)
	goto L73
L76:
	;
	v330 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v320))) = uint8(v330)
	v339 = v8 + int32(81)
	v340 = int32(63)
	v341 = int64(0) - v316
	v342 = int32(1)
	goto L73
L77:
	;
	v362 = v343 + v342
	goto L70
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v362
	*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v8 + int32(80)
	v378 = F_sdscatprintf(m, v363, int32(_a_F_createSharedObjects_30), v8+int32(64))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v380 = F_createObject(m, int32(0), v378)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v382 = F_makeObjectShared(m, v380)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v316)<<(uint(int32(2))%32))+uint32(_c_F_createSharedObjects[47]))) = v382
	v386 = v316 + int64(1)
	if v386 != int64(10) {
		v316 = v386
		goto L68
	} else {
		goto L83
	}
L83:
	;
	goto L69
L84:
	;
	v395 = F_makeObjectShared(m, v393)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[48])) = v395
	v401 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_31), int32(14))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v403 = F_makeObjectShared(m, v401)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[49])) = v403
	v409 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_32), int32(15))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v411 = F_makeObjectShared(m, v409)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[50])) = v411
	v417 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_33), int32(18))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v419 = F_makeObjectShared(m, v417)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[51])) = v419
	v425 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_34), int32(17))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v427 = F_makeObjectShared(m, v425)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[52])) = v427
	v433 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_35), int32(19))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v435 = F_makeObjectShared(m, v433)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[53])) = v435
	v441 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_36), int32(14))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v443 = F_makeObjectShared(m, v441)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[54])) = v443
	v449 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_37), int32(17))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v451 = F_makeObjectShared(m, v449)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[55])) = v451
	v457 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_38), int32(19))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v459 = F_makeObjectShared(m, v457)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[56])) = v459
	v465 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_39), int32(3))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v467 = F_makeObjectShared(m, v465)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[57])) = v467
	v473 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_40), int32(6))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v475 = F_makeObjectShared(m, v473)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[58])) = v475
	v481 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_41), int32(4))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v483 = F_makeObjectShared(m, v481)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[59])) = v483
	v489 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_42), int32(4))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v491 = F_makeObjectShared(m, v489)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[60])) = v491
	v497 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_43), int32(5))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v499 = F_makeObjectShared(m, v497)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[61])) = v499
	v505 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_44), int32(9))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v507 = F_makeObjectShared(m, v505)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[62])) = v507
	v513 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_45), int32(5))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v515 = F_makeObjectShared(m, v513)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[63])) = v515
	v521 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_46), int32(6))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v523 = F_makeObjectShared(m, v521)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[64])) = v523
	v529 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_47), int32(7))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v531 = F_makeObjectShared(m, v529)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[65])) = v531
	v537 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_48), int32(7))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v539 = F_makeObjectShared(m, v537)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[66])) = v539
	v545 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_49), int32(5))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v547 = F_makeObjectShared(m, v545)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[67])) = v547
	v553 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_50), int32(4))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v555 = F_makeObjectShared(m, v553)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[68])) = v555
	v561 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_51), int32(4))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v563 = F_makeObjectShared(m, v561)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[69])) = v563
	v569 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_52), int32(6))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v571 = F_makeObjectShared(m, v569)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[70])) = v571
	v577 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_53), int32(4))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v579 = F_makeObjectShared(m, v577)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[71])) = v579
	v585 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_54), int32(10))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v587 = F_makeObjectShared(m, v585)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[72])) = v587
	v593 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_55), int32(8))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v595 = F_makeObjectShared(m, v593)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[73])) = v595
	v601 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_56), int32(4))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v603 = F_makeObjectShared(m, v601)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[74])) = v603
	v609 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_57), int32(6))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v611 = F_makeObjectShared(m, v609)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[75])) = v611
	v617 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_58), int32(6))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v619 = F_makeObjectShared(m, v617)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[76])) = v619
	v625 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_59), int32(6))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v627 = F_makeObjectShared(m, v625)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[77])) = v627
	v633 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_60), int32(8))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v635 = F_makeObjectShared(m, v633)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[78])) = v635
	v641 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_61), int32(9))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v643 = F_makeObjectShared(m, v641)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[79])) = v643
	v649 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_62), int32(7))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v651 = F_makeObjectShared(m, v649)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[80])) = v651
	v657 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_63), int32(7))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v659 = F_makeObjectShared(m, v657)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[81])) = v659
	v665 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_64), int32(3))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v667 = F_makeObjectShared(m, v665)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[82])) = v667
	v673 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_65), int32(4))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v675 = F_makeObjectShared(m, v673)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[83])) = v675
	v681 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_66), int32(7))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v683 = F_makeObjectShared(m, v681)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[84])) = v683
	v689 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_67), int32(9))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v691 = F_makeObjectShared(m, v689)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[85])) = v691
	v697 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_68), int32(4))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v699 = F_makeObjectShared(m, v697)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[86])) = v699
	v705 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_69), int32(4))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v707 = F_makeObjectShared(m, v705)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[87])) = v707
	v713 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_70), int32(5))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v715 = F_makeObjectShared(m, v713)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[88])) = v715
	v721 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_71), int32(4))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v723 = F_makeObjectShared(m, v721)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[89])) = v723
	v729 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_72), int32(4))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v731 = F_makeObjectShared(m, v729)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[90])) = v731
	v737 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_73), int32(10))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	v739 = F_makeObjectShared(m, v737)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[91])) = v739
	v745 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_74), int32(5))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v747 = F_makeObjectShared(m, v745)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[92])) = v747
	v753 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_75), int32(6))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v755 = F_makeObjectShared(m, v753)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[93])) = v755
	v761 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_76), int32(11))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v763 = F_makeObjectShared(m, v761)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[94])) = v763
	v769 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_77), int32(6))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	v771 = F_makeObjectShared(m, v769)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[95])) = v771
	v777 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_78), int32(7))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v779 = F_makeObjectShared(m, v777)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[96])) = v779
	v785 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_79), int32(4))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	v787 = F_makeObjectShared(m, v785)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[97])) = v787
	v793 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_80), int32(5))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v795 = F_makeObjectShared(m, v793)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[98])) = v795
	v801 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_81), int32(7))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v803 = F_makeObjectShared(m, v801)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[99])) = v803
	v809 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_82), int32(6))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	v811 = F_makeObjectShared(m, v809)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[100])) = v811
	v817 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_83), int32(4))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v819 = F_makeObjectShared(m, v817)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[101])) = v819
	v825 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_84), int32(14))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v827 = F_makeObjectShared(m, v825)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[102])) = v827
	v833 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_85), int32(6))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v835 = F_makeObjectShared(m, v833)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[103])) = v835
	v841 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_86), int32(1))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v843 = F_makeObjectShared(m, v841)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[104])) = v843
	v849 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_87), int32(1))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v851 = F_makeObjectShared(m, v849)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[105])) = v851
	v857 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_88), int32(10))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v859 = F_makeObjectShared(m, v857)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[106])) = v859
	v865 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_89), int32(6))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	v867 = F_makeObjectShared(m, v865)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[107])) = v867
	v873 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_90), int32(6))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v875 = F_makeObjectShared(m, v873)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[108])) = v875
	v881 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_91), int32(5))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v883 = F_makeObjectShared(m, v881)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[109])) = v883
	v889 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_92), int32(7))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	v891 = F_makeObjectShared(m, v889)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[110])) = v891
	v897 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_93), int32(6))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	v899 = F_makeObjectShared(m, v897)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[111])) = v899
	v905 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_94), int32(4))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	v907 = F_makeObjectShared(m, v905)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[112])) = v907
	v913 = F_createStringObject_1(m, int32(_a_F_createSharedObjects_95), int32(7))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v915 = F_makeObjectShared(m, v913)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[113])) = v915
	v922 = int32(0)
	goto L216
L216:
	;
	v929 = F_createObject(m, int32(0), v922)
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L218
	}
L217:
	;
	v945 = v389
	goto L221
L218:
	;
	v931 = F_makeObjectShared(m, v929)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v922<<(uint(int32(2))%32))+uint32(_c_F_createSharedObjects[114]))) = v931
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v931)))
	*(*int32)(unsafe.Add(mBase, uint32(v931))) = v934&int32(-241) | int32(16)
	v941 = v922 + int32(1)
	if v941 != int32(10000) {
		v922 = v941
		goto L216
	} else {
		goto L220
	}
L220:
	;
	goto L217
L221:
	;
	v949 = F_sdsempty(m)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L223
	}
L222:
	;
	v1019 = F_sdsnew(m, int32(_a_F_createSharedObjects_96))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L1
	} else {
		goto L240
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v945
	v953 = v945 << (uint(int32(2)) % 32)
	v960 = F_sdscatprintf(m, v949, int32(_a_F_createSharedObjects_97), v8+int32(48))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v962 = F_createObject(m, int32(0), v960)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v964 = F_makeObjectShared(m, v962)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v953)+uint32(_c_F_createSharedObjects[115]))) = v964
	v967 = F_sdsempty(m)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v945
	v976 = F_sdscatprintf(m, v967, int32(_a_F_createSharedObjects_98), v8+int32(32))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v978 = F_createObject(m, int32(0), v976)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	v980 = F_makeObjectShared(m, v978)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v953)+uint32(_c_F_createSharedObjects[116]))) = v980
	v983 = F_sdsempty(m)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v945
	v992 = F_sdscatprintf(m, v983, int32(_a_F_createSharedObjects_99), v8+int32(16))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v994 = F_createObject(m, int32(0), v992)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	v996 = F_makeObjectShared(m, v994)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v953)+uint32(_c_F_createSharedObjects[117]))) = v996
	v999 = F_sdsempty(m)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v945
	v1006 = F_sdscatprintf(m, v999, int32(_a_F_createSharedObjects_100), v8)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	v1008 = F_createObject(m, int32(0), v1006)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v1010 = F_makeObjectShared(m, v1008)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v953)+uint32(_c_F_createSharedObjects[118]))) = v1010
	v1014 = v945 + int32(1)
	if v1014 != int32(32) {
		v945 = v1014
		goto L221
	} else {
		goto L239
	}
L239:
	;
	goto L222
L240:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[119])) = v1019
	v1024 = F_sdsnew(m, int32(_a_F_createSharedObjects_101))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createSharedObjects[120])) = v1024
	m.G0 = v8 + int32(144)
	return
}
func F_ctime_r(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v10 = F___localtime_r(m, l0, v6+int32(4))
	if v10 != 0 {
		v12 = F___asctime_r(m, v10, l1)
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = v12
			m.G0 = v6 + int32(48)
			return v16
		}
	} else {
		v16 = int32(0)
		m.G0 = v6 + int32(48)
		return v16
	}
}
func F_cycle(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	v9 = m.G0
	v11 = v9 - int32(256)
	m.G0 = v11
	if l2 < int32(2) {
	} else {
		v17 = l1 + l2<<(uint(int32(2))%32)
		*(*int32)(unsafe.Add(mBase, uint32(v17))) = v11
		if l0 == int32(0) {
		} else {
			v21 = l0
			for {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v31 = int32(256)
				if base.Ui32(v21) < base.Ui32(v31) {
					v34 = v21
				} else {
					v34 = v31
				}
				if v34 == int32(0) {
				} else {
					v37 = F__emscripten_memcpy_bulkmem(m, v29, v30, v34)
					mBase = m.M
				}
				v46 = int32(0)
				for {
					v48 = int32(2)
					v50 = l1 + v46<<(uint(v48)%32)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					v53 = v46 + int32(1)
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1+v53<<(uint(v48)%32))))
					if v34 == int32(0) {
					} else {
						v60 = F__emscripten_memcpy_bulkmem(m, v51, v57, v34)
						mBase = m.M
					}
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					*(*int32)(unsafe.Add(mBase, uint32(v50))) = v62 + v34
					if v53 != l2 {
						v46 = v53
						continue
					} else {
						break
					}
					break
				}
				v66 = v21 - v34
				if v66 != 0 {
					v21 = v66
					continue
				} else {
					break
				}
				break
			}
		}
	}
	m.G0 = v11 + int32(256)
	return
}
