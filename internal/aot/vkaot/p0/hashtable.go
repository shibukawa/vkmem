package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_hashtableChannelsDestructor(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(44))))
	F_decrRefCount(m, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		F_hashtableRelease(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_hashtableClientKeyCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(v3 != v4)
}
func F_hashtableCombineStats(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v6 + v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v10 + v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if base.Ui32(v15) < base.Ui32(v14) {
		v17 = v14
	} else {
		v17 = v15
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v19 + v20
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v23 + v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v31 = int32(0)
	for {
		v35 = int32(2)
		v36 = v31 << (uint(v35) % 32)
		v37 = v27 + v36
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v28+v36)))
		*(*int32)(unsafe.Add(mBase, uint32(v37))) = v38 + v40
		v44 = v36 | int32(4)
		v45 = v27 + v44
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v28+v44)))
		*(*int32)(unsafe.Add(mBase, uint32(v45))) = v46 + v48
		v52 = v31 + v35
		if v52 != int32(50) {
			v31 = v52
			continue
		} else {
			break
		}
		break
	}
	return
}
func F_hashtableCommandGetOriginalName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	return v2
}
func F_hashtableEncObjKeyCompare(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_dictEncObjKeyCompare(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3 ^ int32(1)
	}
}
func F_hashtableEntriesPerBucket(m *base.Module) int32 {
	return int32(12)
}
func F_hashtableExpand(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 != int32(-1) {
		v17 = v3
		return v17
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(l1) < base.Ui32(v8+v9) {
			v17 = v3
			return v17
		} else {
			v13 = F_resize_1(m, l0, l1, int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = v13
				return v17
			}
		}
	}
}
func F_hashtableFindRef(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v147 int64
	_ = v147
	var v149 int64
	_ = v149
	var v150 int64
	_ = v150
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v176 int64
	_ = v176
	var v182 int64
	_ = v182
	var v186 int64
	_ = v186
	var v190 int64
	_ = v190
	var v195 int64
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v11 == v3-v13 {
		v212 = v3
		m.G0 = v8 + int32(16)
		return v212
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
		if v18 == int32(0) {
			v26 = v8 + int32(12)
			v27 = int32(4)
			v28 = int32(_a624)
			v36 = *(*int64)(unsafe.Add(mBase, _consts[292]))
			v38 = v36 ^ int64(8317987319222330741)
			v39 = *(*int64)(unsafe.Add(mBase, _consts[293]))
			v41 = v39 ^ int64(7237128888997146477)
			v43 = v36 ^ int64(7816392313619706465)
			v45 = v39 ^ int64(8387220255154660723)
			v50 = v8 + int32(16) - v27
			if v26 == v50 {
				v88 = v26
				v91 = v43
				v92 = v38
				v93 = v45
				v94 = v41
			} else {
				v52 = v26
				v55 = v43
				v56 = v38
				v57 = v45
				v58 = v41
				for {
					v62 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
					v63 = v62 ^ v57
					v64 = v63 + v55
					v65 = v56 + v58
					v68 = v65 ^ base.I64_rotl(v58, int64(13))
					v69 = v64 + v68
					v72 = v69 ^ base.I64_rotl(v68, int64(17))
					v75 = base.I64_rotl(v63, int64(16)) ^ v64
					v78 = int64(32)
					v80 = v75 + base.I64_rotl(v65, v78)
					v81 = base.I64_rotl(v75, int64(21)) ^ v80
					v83 = base.I64_rotl(v69, v78)
					v84 = v80 ^ v62
					v86 = v52 + int32(8)
					if v86 != v50 {
						v52 = v86
						v55 = v83
						v56 = v84
						v57 = v81
						v58 = v72
						continue
					} else {
						break
					}
					break
				}
				v88 = v50
				v91 = v83
				v92 = v84
				v93 = v81
				v94 = v72
			}
			v99 = base.I64_extend_i32_u(v27) << (uint(int64(56)) % 64)
			switch v27 {
			default:
				v132 = v99
			case 1:
				v129 = v99
				v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
				v132 = v129 | v130
			case 2:
				v124 = v99
				v125 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
				v129 = v125<<(uint(int64(8))%64) | v124
				v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
				v132 = v129 | v130
			case 3:
				v119 = v99
				v120 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+2)))
				v124 = v120<<(uint(int64(16))%64) | v119
				v125 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
				v129 = v125<<(uint(int64(8))%64) | v124
				v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
				v132 = v129 | v130
			case 4:
				v114 = v99
				v115 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+3)))
				v119 = v115<<(uint(int64(24))%64) | v114
				v120 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+2)))
				v124 = v120<<(uint(int64(16))%64) | v119
				v125 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
				v129 = v125<<(uint(int64(8))%64) | v124
				v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
				v132 = v129 | v130
			case 5:
				v109 = v99
				v110 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+4)))
				v114 = v110<<(uint(int64(32))%64) | v109
				v115 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+3)))
				v119 = v115<<(uint(int64(24))%64) | v114
				v120 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+2)))
				v124 = v120<<(uint(int64(16))%64) | v119
				v125 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
				v129 = v125<<(uint(int64(8))%64) | v124
				v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
				v132 = v129 | v130
			case 6:
				v104 = v99
				v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+5)))
				v109 = v105<<(uint(int64(40))%64) | v104
				v110 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+4)))
				v114 = v110<<(uint(int64(32))%64) | v109
				v115 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+3)))
				v119 = v115<<(uint(int64(24))%64) | v114
				v120 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+2)))
				v124 = v120<<(uint(int64(16))%64) | v119
				v125 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
				v129 = v125<<(uint(int64(8))%64) | v124
				v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
				v132 = v129 | v130
			case 7:
				v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+6)))
				v104 = v100<<(uint(int64(48))%64) | v99
				v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+5)))
				v109 = v105<<(uint(int64(40))%64) | v104
				v110 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+4)))
				v114 = v110<<(uint(int64(32))%64) | v109
				v115 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+3)))
				v119 = v115<<(uint(int64(24))%64) | v114
				v120 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+2)))
				v124 = v120<<(uint(int64(16))%64) | v119
				v125 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
				v129 = v125<<(uint(int64(8))%64) | v124
				v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
				v132 = v129 | v130
			}
			v133 = v132 ^ v93
			v134 = int64(16)
			v136 = v133 + v91
			v137 = base.I64_rotl(v133, v134) ^ v136
			v138 = int64(21)
			v140 = v92 + v94
			v141 = int64(32)
			v143 = v137 + base.I64_rotl(v140, v141)
			v144 = base.I64_rotl(v137, v138) ^ v143
			v147 = int64(13)
			v149 = v140 ^ base.I64_rotl(v94, v147)
			v150 = v136 + v149
			v155 = base.I64_rotl(v150, v141) ^ int64(255) + v144
			v156 = base.I64_rotl(v144, v134) ^ v155
			v160 = int64(17)
			v162 = v150 ^ base.I64_rotl(v149, v160)
			v163 = v143 ^ v132 + v162
			v166 = base.I64_rotl(v163, v141) + v156
			v167 = base.I64_rotl(v156, v138) ^ v166
			v172 = v163 ^ base.I64_rotl(v162, v147)
			v173 = v172 + v155
			v176 = base.I64_rotl(v173, v141) + v167
			v182 = base.I64_rotl(v172, v160) ^ v173
			v186 = base.I64_rotl(v182, v147) ^ (v182 + v166)
			v190 = v186 + v176
			v195 = base.I64_rotl(base.I64_rotl(v167, v134)^v176, v138) ^ base.I64_rotl(v186, v160) ^ base.I64_rotl(v190, v141) ^ v190
			v196 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v196
			v201 = F_findBucket_1(m, l0, v195, l1, v8+int32(8), v196)
			mBase = m.M
			v202 = m.ExcPending
			if v202 != 0 {
				return int32(0)
			} else {
				v203 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
				if v201 != 0 {
					v210 = v201 + v203<<(uint(int32(2))%32) + int32(16)
				} else {
					v210 = int32(0)
				}
				v212 = v210
				m.G0 = v8 + int32(16)
				return v212
			}
		} else {
			v21 = m.T0[v18].(func(*base.Module, int32) int64)(m, l1)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v195 = v21
				v196 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v196
				v201 = F_findBucket_1(m, l0, v195, l1, v8+int32(8), v196)
				mBase = m.M
				v202 = m.ExcPending
				if v202 != 0 {
					return int32(0)
				} else {
					v203 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
					if v201 != 0 {
						v210 = v201 + v203<<(uint(int32(2))%32) + int32(16)
					} else {
						v210 = int32(0)
					}
					v212 = v210
					m.G0 = v8 + int32(16)
					return v212
				}
			}
		}
	}
}
func F_hashtableGenHashFunction(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v151 int64
	_ = v151
	var v157 int64
	_ = v157
	var v161 int64
	_ = v161
	var v165 int64
	_ = v165
	v3 = int32(_a624)
	v11 = *(*int64)(unsafe.Add(mBase, _consts[292]))
	v13 = v11 ^ int64(8317987319222330741)
	v14 = *(*int64)(unsafe.Add(mBase, _consts[293]))
	v16 = v14 ^ int64(7237128888997146477)
	v18 = v11 ^ int64(7816392313619706465)
	v20 = v14 ^ int64(8387220255154660723)
	v24 = l1 & int32(7)
	v25 = l0 + l1 - v24
	if l0 == v25 {
		v63 = l0
		v66 = v18
		v67 = v13
		v68 = v20
		v69 = v16
	} else {
		v27 = l0
		v30 = v18
		v31 = v13
		v32 = v20
		v33 = v16
		for {
			v37 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
			v38 = v37 ^ v32
			v39 = v38 + v30
			v40 = v31 + v33
			v43 = v40 ^ base.I64_rotl(v33, int64(13))
			v44 = v39 + v43
			v47 = v44 ^ base.I64_rotl(v43, int64(17))
			v50 = base.I64_rotl(v38, int64(16)) ^ v39
			v53 = int64(32)
			v55 = v50 + base.I64_rotl(v40, v53)
			v56 = base.I64_rotl(v50, int64(21)) ^ v55
			v58 = base.I64_rotl(v44, v53)
			v59 = v55 ^ v37
			v61 = v27 + int32(8)
			if v61 != v25 {
				v27 = v61
				v30 = v58
				v31 = v59
				v32 = v56
				v33 = v47
				continue
			} else {
				break
			}
			break
		}
		v63 = v25
		v66 = v58
		v67 = v59
		v68 = v56
		v69 = v47
	}
	v74 = base.I64_extend_i32_u(l1) << (uint(int64(56)) % 64)
	switch v24 {
	default:
		v107 = v74
	case 1:
		v104 = v74
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 2:
		v99 = v74
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 3:
		v94 = v74
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 4:
		v89 = v74
		v90 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
		v94 = v90<<(uint(int64(24))%64) | v89
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 5:
		v84 = v74
		v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
		v89 = v85<<(uint(int64(32))%64) | v84
		v90 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
		v94 = v90<<(uint(int64(24))%64) | v89
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 6:
		v79 = v74
		v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
		v84 = v80<<(uint(int64(40))%64) | v79
		v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
		v89 = v85<<(uint(int64(32))%64) | v84
		v90 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
		v94 = v90<<(uint(int64(24))%64) | v89
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	case 7:
		v75 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+6)))
		v79 = v75<<(uint(int64(48))%64) | v74
		v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
		v84 = v80<<(uint(int64(40))%64) | v79
		v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
		v89 = v85<<(uint(int64(32))%64) | v84
		v90 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
		v94 = v90<<(uint(int64(24))%64) | v89
		v95 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
		v99 = v95<<(uint(int64(16))%64) | v94
		v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
		v107 = v104 | v105
	}
	v108 = v107 ^ v68
	v109 = int64(16)
	v111 = v108 + v66
	v112 = base.I64_rotl(v108, v109) ^ v111
	v113 = int64(21)
	v115 = v67 + v69
	v116 = int64(32)
	v118 = v112 + base.I64_rotl(v115, v116)
	v119 = base.I64_rotl(v112, v113) ^ v118
	v122 = int64(13)
	v124 = v115 ^ base.I64_rotl(v69, v122)
	v125 = v111 + v124
	v130 = base.I64_rotl(v125, v116) ^ int64(255) + v119
	v131 = base.I64_rotl(v119, v109) ^ v130
	v135 = int64(17)
	v137 = v125 ^ base.I64_rotl(v124, v135)
	v138 = v118 ^ v107 + v137
	v141 = base.I64_rotl(v138, v116) + v131
	v142 = base.I64_rotl(v131, v113) ^ v141
	v147 = v138 ^ base.I64_rotl(v137, v122)
	v148 = v147 + v130
	v151 = base.I64_rotl(v148, v116) + v142
	v157 = base.I64_rotl(v147, v135) ^ v148
	v161 = base.I64_rotl(v157, v122) ^ (v157 + v141)
	v165 = v161 + v151
	return base.I64_rotl(base.I64_rotl(v142, v109)^v151, v113) ^ base.I64_rotl(v161, v135) ^ base.I64_rotl(v165, v116) ^ v165
}
func F_hashtableIncrementalFindGetResult(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 != int32(3) {
		if v4 == int32(4) {
			return base.B2i32(v4 == int32(3))
		} else {
			F__serverAssert(m, int32(_a639), int32(_a626), int32(1987))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		if l1 == int32(0) {
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v9+v10<<(uint(int32(2))%32)+int32(16))))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v16
		}
		return base.B2i32(v4 == int32(3))
	}
}
func F_hashtableIncrementalFindInit(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v181 int64
	_ = v181
	var v187 int64
	_ = v187
	var v191 int64
	_ = v191
	var v195 int64
	_ = v195
	var v200 int64
	_ = v200
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v10 != int32(0)-v12 {
		v17 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
		if v25 == v17 {
			v31 = v8 + int32(12)
			v32 = int32(4)
			v33 = int32(_a624)
			v41 = *(*int64)(unsafe.Add(mBase, _consts[292]))
			v43 = v41 ^ int64(8317987319222330741)
			v44 = *(*int64)(unsafe.Add(mBase, _consts[293]))
			v46 = v44 ^ int64(7237128888997146477)
			v48 = v41 ^ int64(7816392313619706465)
			v50 = v44 ^ int64(8387220255154660723)
			v55 = v8 + int32(16) - v32
			if v31 == v55 {
				v93 = v31
				v96 = v48
				v97 = v43
				v98 = v50
				v99 = v46
			} else {
				v57 = v31
				v60 = v48
				v61 = v43
				v62 = v50
				v63 = v46
				for {
					v67 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
					v68 = v67 ^ v62
					v69 = v68 + v60
					v70 = v61 + v63
					v73 = v70 ^ base.I64_rotl(v63, int64(13))
					v74 = v69 + v73
					v77 = v74 ^ base.I64_rotl(v73, int64(17))
					v80 = base.I64_rotl(v68, int64(16)) ^ v69
					v83 = int64(32)
					v85 = v80 + base.I64_rotl(v70, v83)
					v86 = base.I64_rotl(v80, int64(21)) ^ v85
					v88 = base.I64_rotl(v74, v83)
					v89 = v85 ^ v67
					v91 = v57 + int32(8)
					if v91 != v55 {
						v57 = v91
						v60 = v88
						v61 = v89
						v62 = v86
						v63 = v77
						continue
					} else {
						break
					}
					break
				}
				v93 = v55
				v96 = v88
				v97 = v89
				v98 = v86
				v99 = v77
			}
			v104 = base.I64_extend_i32_u(v32) << (uint(int64(56)) % 64)
			switch v32 {
			default:
				v137 = v104
			case 1:
				v134 = v104
				v135 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
				v137 = v134 | v135
			case 2:
				v129 = v104
				v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
				v134 = v130<<(uint(int64(8))%64) | v129
				v135 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
				v137 = v134 | v135
			case 3:
				v124 = v104
				v125 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+2)))
				v129 = v125<<(uint(int64(16))%64) | v124
				v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
				v134 = v130<<(uint(int64(8))%64) | v129
				v135 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
				v137 = v134 | v135
			case 4:
				v119 = v104
				v120 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+3)))
				v124 = v120<<(uint(int64(24))%64) | v119
				v125 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+2)))
				v129 = v125<<(uint(int64(16))%64) | v124
				v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
				v134 = v130<<(uint(int64(8))%64) | v129
				v135 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
				v137 = v134 | v135
			case 5:
				v114 = v104
				v115 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+4)))
				v119 = v115<<(uint(int64(32))%64) | v114
				v120 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+3)))
				v124 = v120<<(uint(int64(24))%64) | v119
				v125 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+2)))
				v129 = v125<<(uint(int64(16))%64) | v124
				v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
				v134 = v130<<(uint(int64(8))%64) | v129
				v135 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
				v137 = v134 | v135
			case 6:
				v109 = v104
				v110 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+5)))
				v114 = v110<<(uint(int64(40))%64) | v109
				v115 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+4)))
				v119 = v115<<(uint(int64(32))%64) | v114
				v120 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+3)))
				v124 = v120<<(uint(int64(24))%64) | v119
				v125 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+2)))
				v129 = v125<<(uint(int64(16))%64) | v124
				v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
				v134 = v130<<(uint(int64(8))%64) | v129
				v135 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
				v137 = v134 | v135
			case 7:
				v105 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+6)))
				v109 = v105<<(uint(int64(48))%64) | v104
				v110 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+5)))
				v114 = v110<<(uint(int64(40))%64) | v109
				v115 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+4)))
				v119 = v115<<(uint(int64(32))%64) | v114
				v120 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+3)))
				v124 = v120<<(uint(int64(24))%64) | v119
				v125 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+2)))
				v129 = v125<<(uint(int64(16))%64) | v124
				v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
				v134 = v130<<(uint(int64(8))%64) | v129
				v135 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
				v137 = v134 | v135
			}
			v138 = v137 ^ v98
			v139 = int64(16)
			v141 = v138 + v96
			v142 = base.I64_rotl(v138, v139) ^ v141
			v143 = int64(21)
			v145 = v97 + v99
			v146 = int64(32)
			v148 = v142 + base.I64_rotl(v145, v146)
			v149 = base.I64_rotl(v142, v143) ^ v148
			v152 = int64(13)
			v154 = v145 ^ base.I64_rotl(v99, v152)
			v155 = v141 + v154
			v160 = base.I64_rotl(v155, v146) ^ int64(255) + v149
			v161 = base.I64_rotl(v149, v139) ^ v160
			v165 = int64(17)
			v167 = v155 ^ base.I64_rotl(v154, v165)
			v168 = v148 ^ v137 + v167
			v171 = base.I64_rotl(v168, v146) + v161
			v172 = base.I64_rotl(v161, v143) ^ v171
			v177 = v168 ^ base.I64_rotl(v167, v152)
			v178 = v177 + v160
			v181 = base.I64_rotl(v178, v146) + v172
			v187 = base.I64_rotl(v177, v165) ^ v178
			v191 = base.I64_rotl(v187, v152) ^ (v187 + v171)
			v195 = v191 + v181
			v200 = base.I64_rotl(base.I64_rotl(v172, v139)^v181, v143) ^ base.I64_rotl(v191, v165) ^ base.I64_rotl(v195, v146) ^ v195
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v200
			m.G0 = v8 + int32(16)
			return
		} else {
			v28 = m.T0[v25].(func(*base.Module, int32) int64)(m, l2)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v200 = v28
				*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v200
				m.G0 = v8 + int32(16)
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(4)
		m.G0 = v8 + int32(16)
		return
	}
}
func F_hashtableIncrementalFindStep(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v8 {
	case 0:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	case 3, 4:
		goto L1
	default:
		goto L3
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(4)
	goto L1
L3:
	;
	F__serverAssert(m, int32(_a107), int32(_a626), int32(1974))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L13
	} else {
		goto L47
	}
L4:
	;
	v196 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v196)
	v198 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v198
	return v198
L5:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v157&int32(1) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L6:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v151 = v149
	v153 = l0 + int32(12)
	v154 = v49
	goto L5
L7:
	;
	v102 = l0 + int32(12)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v98 != 0 {
		v151 = v103
		v153 = v102
		v154 = v98
		goto L5
	} else {
		goto L29
	}
L8:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v98 = v93
	goto L7
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49))))
	if v50&int32(8190) == int32(0) {
		goto L6
	} else {
		goto L21
	}
L10:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v9+v10<<(uint(int32(2))%32)+int32(16))))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v19 == int32(0) {
		v27 = v16
		v28 = v18
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v30 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v22 = m.T0[v19].(func(*base.Module, int32) int32)(m, v16)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v27 = v22
	v28 = v26
	goto L11
L15:
	;
	if v36 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v36 = base.B2i32(v29 != v27)
	goto L15
L17:
	;
	v33 = m.T0[v30].(func(*base.Module, int32, int32) int32)(m, v29, v27)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v36 = v33
	goto L15
L19:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	v43 = v41 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v43)
	goto L9
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(3)
	return int32(0)
L21:
	;
	v58 = int32(12) - v50&int32(1)
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v58 <= v59 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	v71 = v59
	goto L23
L23:
	;
	if int32(base.Ui32(int32(base.Ui32(v50)>>(uint(int32(1))%32))&int32(4095))>>(uint(v71)%32))&int32(1) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v91 = v71 + int32(1)
	if v91 < v58 {
		v71 = v91
		goto L23
	} else {
		goto L28
	}
L26:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+int32(2)+v71))))
	if v83 != v67&int32(255) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v71)
	return int32(1)
L28:
	;
	v98 = v49
	goto L7
L29:
	;
	v104 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v104)
	v108 = int32(-1)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+24)))
	if v109 == int32(255) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v115 = v104
	goto L32
L31:
	;
	v115 = v108<<(uint(v109)%32) ^ v108
	goto L32
L32:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v117 = v115 & v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v118 < int32(0) {
		v135 = v117
		v136 = v104
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v103+v136<<(uint(int32(2))%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v140 + v135<<(uint(int32(6))%32)
	goto L4
L34:
	;
	if base.Ui32(v118) <= base.Ui32(v117) {
		v135 = v117
		v136 = v104
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v122 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v122)
	v126 = int32(-1)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+25)))
	if v127 == int32(255) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v133 = int32(0)
	goto L38
L37:
	;
	v133 = v126<<(uint(v127)%32) ^ v126
	goto L38
L38:
	;
	v135 = v133 & v116
	v136 = v122
	goto L33
L39:
	;
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v167 != 0 {
		goto L2
	} else {
		goto L42
	}
L40:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v154)+60))
	if v162 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v162
	goto L4
L42:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if v168 < int32(0) {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v171 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v171)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	v175 = int32(-1)
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+25)))
	if v176 == int32(255) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v182 = int32(0)
	goto L46
L45:
	;
	v182 = v175<<(uint(v176)%32) ^ v175
	goto L46
L46:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v173 + v182&v183<<(uint(int32(6))%32)
	goto L4
L47:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hashtableInsertAtPosition(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8))))
	v10 = int32(1)
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v14 = v10 << (uint(v13) % 32)
	if int32(base.Ui32(v9)>>(uint(v10)%32))&v14&int32(4095) == int32(0) {
		v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
		v27 = int32(1)
		v31 = v14<<(uint(v27)%32)&int32(8190) | v9
		*(*uint16)(unsafe.Add(mBase, uint32(v8))) = uint16(v31)
		v33 = int32(2)
		*(*int32)(unsafe.Add(mBase, uint32(v8+v13<<(uint(v33)%32))+16)) = l1
		v39 = l0 + v26<<(uint(v33)%32)
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v40 + v27
		return
	} else {
		F__serverAssert(m, int32(_a632), int32(_a626), int32(1717))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
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
func F_hashtableMetadata(m *base.Module, l0 int32) int32 {
	return l0 + int32(44)
}
func F_hashtableObjectGetKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5&int32(2) == v2 {
		v25 = v2
	} else {
		v19 = l0 + (v5&int32(4) ^ int32(12)) + v5<<(uint(int32(3))%32)&int32(8)
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
		v25 = v19 + v20 + int32(1)
	}
	return v25
}
func F_hashtablePop(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v150 int64
	_ = v150
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v179 int64
	_ = v179
	var v185 int64
	_ = v185
	var v189 int64
	_ = v189
	var v193 int64
	_ = v193
	var v198 int64
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v14 == v4-v16 {
		v285 = v4
		m.G0 = v11 + int32(16)
		return v285
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l1
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
		if v21 == int32(0) {
			v29 = v11 + int32(12)
			v30 = int32(4)
			v31 = int32(_a624)
			v39 = *(*int64)(unsafe.Add(mBase, _consts[292]))
			v41 = v39 ^ int64(8317987319222330741)
			v42 = *(*int64)(unsafe.Add(mBase, _consts[293]))
			v44 = v42 ^ int64(7237128888997146477)
			v46 = v39 ^ int64(7816392313619706465)
			v48 = v42 ^ int64(8387220255154660723)
			v53 = v11 + int32(16) - v30
			if v29 == v53 {
				v91 = v29
				v94 = v46
				v95 = v41
				v96 = v48
				v97 = v44
			} else {
				v55 = v29
				v58 = v46
				v59 = v41
				v60 = v48
				v61 = v44
				for {
					v65 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
					v66 = v65 ^ v60
					v67 = v66 + v58
					v68 = v59 + v61
					v71 = v68 ^ base.I64_rotl(v61, int64(13))
					v72 = v67 + v71
					v75 = v72 ^ base.I64_rotl(v71, int64(17))
					v78 = base.I64_rotl(v66, int64(16)) ^ v67
					v81 = int64(32)
					v83 = v78 + base.I64_rotl(v68, v81)
					v84 = base.I64_rotl(v78, int64(21)) ^ v83
					v86 = base.I64_rotl(v72, v81)
					v87 = v83 ^ v65
					v89 = v55 + int32(8)
					if v89 != v53 {
						v55 = v89
						v58 = v86
						v59 = v87
						v60 = v84
						v61 = v75
						continue
					} else {
						break
					}
					break
				}
				v91 = v53
				v94 = v86
				v95 = v87
				v96 = v84
				v97 = v75
			}
			v102 = base.I64_extend_i32_u(v30) << (uint(int64(56)) % 64)
			switch v30 {
			default:
				v135 = v102
			case 1:
				v132 = v102
				v133 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
				v135 = v132 | v133
			case 2:
				v127 = v102
				v128 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
				v132 = v128<<(uint(int64(8))%64) | v127
				v133 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
				v135 = v132 | v133
			case 3:
				v122 = v102
				v123 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+2)))
				v127 = v123<<(uint(int64(16))%64) | v122
				v128 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
				v132 = v128<<(uint(int64(8))%64) | v127
				v133 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
				v135 = v132 | v133
			case 4:
				v117 = v102
				v118 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+3)))
				v122 = v118<<(uint(int64(24))%64) | v117
				v123 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+2)))
				v127 = v123<<(uint(int64(16))%64) | v122
				v128 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
				v132 = v128<<(uint(int64(8))%64) | v127
				v133 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
				v135 = v132 | v133
			case 5:
				v112 = v102
				v113 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+4)))
				v117 = v113<<(uint(int64(32))%64) | v112
				v118 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+3)))
				v122 = v118<<(uint(int64(24))%64) | v117
				v123 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+2)))
				v127 = v123<<(uint(int64(16))%64) | v122
				v128 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
				v132 = v128<<(uint(int64(8))%64) | v127
				v133 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
				v135 = v132 | v133
			case 6:
				v107 = v102
				v108 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+5)))
				v112 = v108<<(uint(int64(40))%64) | v107
				v113 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+4)))
				v117 = v113<<(uint(int64(32))%64) | v112
				v118 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+3)))
				v122 = v118<<(uint(int64(24))%64) | v117
				v123 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+2)))
				v127 = v123<<(uint(int64(16))%64) | v122
				v128 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
				v132 = v128<<(uint(int64(8))%64) | v127
				v133 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
				v135 = v132 | v133
			case 7:
				v103 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+6)))
				v107 = v103<<(uint(int64(48))%64) | v102
				v108 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+5)))
				v112 = v108<<(uint(int64(40))%64) | v107
				v113 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+4)))
				v117 = v113<<(uint(int64(32))%64) | v112
				v118 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+3)))
				v122 = v118<<(uint(int64(24))%64) | v117
				v123 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+2)))
				v127 = v123<<(uint(int64(16))%64) | v122
				v128 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
				v132 = v128<<(uint(int64(8))%64) | v127
				v133 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
				v135 = v132 | v133
			}
			v136 = v135 ^ v96
			v137 = int64(16)
			v139 = v136 + v94
			v140 = base.I64_rotl(v136, v137) ^ v139
			v141 = int64(21)
			v143 = v95 + v97
			v144 = int64(32)
			v146 = v140 + base.I64_rotl(v143, v144)
			v147 = base.I64_rotl(v140, v141) ^ v146
			v150 = int64(13)
			v152 = v143 ^ base.I64_rotl(v97, v150)
			v153 = v139 + v152
			v158 = base.I64_rotl(v153, v144) ^ int64(255) + v147
			v159 = base.I64_rotl(v147, v137) ^ v158
			v163 = int64(17)
			v165 = v153 ^ base.I64_rotl(v152, v163)
			v166 = v146 ^ v135 + v165
			v169 = base.I64_rotl(v166, v144) + v159
			v170 = base.I64_rotl(v159, v141) ^ v169
			v175 = v166 ^ base.I64_rotl(v165, v150)
			v176 = v175 + v158
			v179 = base.I64_rotl(v176, v144) + v170
			v185 = base.I64_rotl(v175, v163) ^ v176
			v189 = base.I64_rotl(v185, v150) ^ (v185 + v169)
			v193 = v189 + v179
			v198 = base.I64_rotl(base.I64_rotl(v170, v137)^v179, v141) ^ base.I64_rotl(v189, v163) ^ base.I64_rotl(v193, v144) ^ v193
			v199 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v199
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v199
			v207 = F_findBucket_1(m, l0, v198, l1, v11+int32(8), v11+int32(4))
			mBase = m.M
			v208 = m.ExcPending
			if v208 != 0 {
				return int32(0)
			} else {
				if v207 == int32(0) {
					v285 = base.B2i32(v207 != int32(0))
					m.G0 = v11 + int32(16)
					return v285
				} else {
					v213 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					if l2 == int32(0) {
					} else {
						v221 = *(*int32)(unsafe.Add(mBase, uint32(v207+v213<<(uint(int32(2))%32)+int32(16))))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v221
					}
					v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207))))
					v226 = int32(1)
					v233 = v223&(base.I32_rotl(int32(-2), v213)<<(uint(v226)%32))&int32(8190) | v223&int32(57345)
					*(*uint16)(unsafe.Add(mBase, uint32(v207))) = uint16(v233)
					v235 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					v238 = l0 + int32(16) + v235<<(uint(int32(2))%32)
					v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
					*(*int32)(unsafe.Add(mBase, uint32(v238))) = v239 + int32(-1)
					v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
					if v243&v226 == int32(0) {
						v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v253 != int32(-1) {
							v285 = base.B2i32(v207 != int32(0))
							m.G0 = v11 + int32(16)
							return v285
						} else {
							v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
							if v256 != 0 {
								v285 = base.B2i32(v207 != int32(0))
								m.G0 = v11 + int32(16)
								return v285
							} else {
								v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v262 == int32(255) {
									v266 = int32(0)
								} else {
									v266 = int32(12) << (uint(v262) % 32)
								}
								v270 = *(*int32)(unsafe.Add(mBase, _consts[295]))
								if v270 != 0 {
									v271 = int32(3)
								} else {
									v271 = int32(13)
								}
								if base.Ui32(v266*v271) < base.Ui32(v257*int32(100)) {
									v285 = base.B2i32(v207 != int32(0))
									m.G0 = v11 + int32(16)
									return v285
								} else {
									v275 = F_resize_1(m, l0, v257, int32(0))
									mBase = m.M
									v276 = m.ExcPending
									if v276 != 0 {
										return int32(0)
									} else {
										v285 = base.B2i32(v207 != int32(0))
										m.G0 = v11 + int32(16)
										return v285
									}
								}
							}
						}
					} else {
						v248 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+26)))
						if int32(0) < v248 {
							v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v253 != int32(-1) {
								v285 = base.B2i32(v207 != int32(0))
								m.G0 = v11 + int32(16)
								return v285
							} else {
								v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
								if v256 != 0 {
									v285 = base.B2i32(v207 != int32(0))
									m.G0 = v11 + int32(16)
									return v285
								} else {
									v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
									if v262 == int32(255) {
										v266 = int32(0)
									} else {
										v266 = int32(12) << (uint(v262) % 32)
									}
									v270 = *(*int32)(unsafe.Add(mBase, _consts[295]))
									if v270 != 0 {
										v271 = int32(3)
									} else {
										v271 = int32(13)
									}
									if base.Ui32(v266*v271) < base.Ui32(v257*int32(100)) {
										v285 = base.B2i32(v207 != int32(0))
										m.G0 = v11 + int32(16)
										return v285
									} else {
										v275 = F_resize_1(m, l0, v257, int32(0))
										mBase = m.M
										v276 = m.ExcPending
										if v276 != 0 {
											return int32(0)
										} else {
											v285 = base.B2i32(v207 != int32(0))
											m.G0 = v11 + int32(16)
											return v285
										}
									}
								}
							}
						} else {
							F_fillBucketHole(m, l0, v207, v213, v235)
							mBase = m.M
							v252 = m.ExcPending
							if v252 != 0 {
								return int32(0)
							} else {
								v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v253 != int32(-1) {
									v285 = base.B2i32(v207 != int32(0))
									m.G0 = v11 + int32(16)
									return v285
								} else {
									v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
									if v256 != 0 {
										v285 = base.B2i32(v207 != int32(0))
										m.G0 = v11 + int32(16)
										return v285
									} else {
										v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v262 == int32(255) {
											v266 = int32(0)
										} else {
											v266 = int32(12) << (uint(v262) % 32)
										}
										v270 = *(*int32)(unsafe.Add(mBase, _consts[295]))
										if v270 != 0 {
											v271 = int32(3)
										} else {
											v271 = int32(13)
										}
										if base.Ui32(v266*v271) < base.Ui32(v257*int32(100)) {
											v285 = base.B2i32(v207 != int32(0))
											m.G0 = v11 + int32(16)
											return v285
										} else {
											v275 = F_resize_1(m, l0, v257, int32(0))
											mBase = m.M
											v276 = m.ExcPending
											if v276 != 0 {
												return int32(0)
											} else {
												v285 = base.B2i32(v207 != int32(0))
												m.G0 = v11 + int32(16)
												return v285
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
			v24 = m.T0[v21].(func(*base.Module, int32) int64)(m, l1)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v198 = v24
				v199 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v199
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v199
				v207 = F_findBucket_1(m, l0, v198, l1, v11+int32(8), v11+int32(4))
				mBase = m.M
				v208 = m.ExcPending
				if v208 != 0 {
					return int32(0)
				} else {
					if v207 == int32(0) {
						v285 = base.B2i32(v207 != int32(0))
						m.G0 = v11 + int32(16)
						return v285
					} else {
						v213 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
						if l2 == int32(0) {
						} else {
							v221 = *(*int32)(unsafe.Add(mBase, uint32(v207+v213<<(uint(int32(2))%32)+int32(16))))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v221
						}
						v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207))))
						v226 = int32(1)
						v233 = v223&(base.I32_rotl(int32(-2), v213)<<(uint(v226)%32))&int32(8190) | v223&int32(57345)
						*(*uint16)(unsafe.Add(mBase, uint32(v207))) = uint16(v233)
						v235 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						v238 = l0 + int32(16) + v235<<(uint(int32(2))%32)
						v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
						*(*int32)(unsafe.Add(mBase, uint32(v238))) = v239 + int32(-1)
						v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
						if v243&v226 == int32(0) {
							v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v253 != int32(-1) {
								v285 = base.B2i32(v207 != int32(0))
								m.G0 = v11 + int32(16)
								return v285
							} else {
								v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
								if v256 != 0 {
									v285 = base.B2i32(v207 != int32(0))
									m.G0 = v11 + int32(16)
									return v285
								} else {
									v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
									if v262 == int32(255) {
										v266 = int32(0)
									} else {
										v266 = int32(12) << (uint(v262) % 32)
									}
									v270 = *(*int32)(unsafe.Add(mBase, _consts[295]))
									if v270 != 0 {
										v271 = int32(3)
									} else {
										v271 = int32(13)
									}
									if base.Ui32(v266*v271) < base.Ui32(v257*int32(100)) {
										v285 = base.B2i32(v207 != int32(0))
										m.G0 = v11 + int32(16)
										return v285
									} else {
										v275 = F_resize_1(m, l0, v257, int32(0))
										mBase = m.M
										v276 = m.ExcPending
										if v276 != 0 {
											return int32(0)
										} else {
											v285 = base.B2i32(v207 != int32(0))
											m.G0 = v11 + int32(16)
											return v285
										}
									}
								}
							}
						} else {
							v248 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+26)))
							if int32(0) < v248 {
								v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v253 != int32(-1) {
									v285 = base.B2i32(v207 != int32(0))
									m.G0 = v11 + int32(16)
									return v285
								} else {
									v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
									if v256 != 0 {
										v285 = base.B2i32(v207 != int32(0))
										m.G0 = v11 + int32(16)
										return v285
									} else {
										v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v262 == int32(255) {
											v266 = int32(0)
										} else {
											v266 = int32(12) << (uint(v262) % 32)
										}
										v270 = *(*int32)(unsafe.Add(mBase, _consts[295]))
										if v270 != 0 {
											v271 = int32(3)
										} else {
											v271 = int32(13)
										}
										if base.Ui32(v266*v271) < base.Ui32(v257*int32(100)) {
											v285 = base.B2i32(v207 != int32(0))
											m.G0 = v11 + int32(16)
											return v285
										} else {
											v275 = F_resize_1(m, l0, v257, int32(0))
											mBase = m.M
											v276 = m.ExcPending
											if v276 != 0 {
												return int32(0)
											} else {
												v285 = base.B2i32(v207 != int32(0))
												m.G0 = v11 + int32(16)
												return v285
											}
										}
									}
								}
							} else {
								F_fillBucketHole(m, l0, v207, v213, v235)
								mBase = m.M
								v252 = m.ExcPending
								if v252 != 0 {
									return int32(0)
								} else {
									v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v253 != int32(-1) {
										v285 = base.B2i32(v207 != int32(0))
										m.G0 = v11 + int32(16)
										return v285
									} else {
										v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
										if v256 != 0 {
											v285 = base.B2i32(v207 != int32(0))
											m.G0 = v11 + int32(16)
											return v285
										} else {
											v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v262 == int32(255) {
												v266 = int32(0)
											} else {
												v266 = int32(12) << (uint(v262) % 32)
											}
											v270 = *(*int32)(unsafe.Add(mBase, _consts[295]))
											if v270 != 0 {
												v271 = int32(3)
											} else {
												v271 = int32(13)
											}
											if base.Ui32(v266*v271) < base.Ui32(v257*int32(100)) {
												v285 = base.B2i32(v207 != int32(0))
												m.G0 = v11 + int32(16)
												return v285
											} else {
												v275 = F_resize_1(m, l0, v257, int32(0))
												mBase = m.M
												v276 = m.ExcPending
												if v276 != 0 {
													return int32(0)
												} else {
													v285 = base.B2i32(v207 != int32(0))
													m.G0 = v11 + int32(16)
													return v285
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
func F_hashtableReleaseIterator(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	F_hashtableCleanupIterator(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		F_valkey_free(m, l0)
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			return
		}
	}
}
func F_hashtableResizeAllowed(m *base.Module, l0 int32, l1 float64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v27 int64
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	v3 = int32(0)
	v9 = *(*int64)(unsafe.Add(mBase, _consts[265]))
	if v9 == int64(0) {
		v62 = v3
	} else {
		v12 = F_zmalloc_used_memory(m)
		mBase = m.M
		v14 = *(*int64)(unsafe.Add(mBase, _consts[265]))
		if base.Ui64(base.I64_extend_i32_u(v12+l0)) <= base.Ui64(v14) {
			v62 = v3
		} else {
			v19 = int32(_a69)
			v20 = *(*int64)(unsafe.Add(mBase, _consts[266]))
			v22 = *(*int32)(unsafe.Add(mBase, _consts[267]))
			if base.I64_extend_i32_u(v22) <= v20 {
				v37 = int32(0)
			} else {
				v27 = base.I64_div_s(v20, int64(16384))
				v34 = v22 - base.I32_wrap_i64(v20+v27*int64(44)) + int32(-44)
				if base.Ui32(v22) < base.Ui32(v34) {
					v36 = int32(0)
				} else {
					v36 = v34
				}
				v37 = v36
			}
			v39 = *(*int32)(unsafe.Add(mBase, _consts[27]))
			if v39 == int32(0) {
				v46 = v37
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, _consts[268]))
				v44 = F_sdsAllocSize(m, v43)
				mBase = m.M
				v46 = v44 + v37
			}
			v47 = F_clusterIsAnySlotExporting(m)
			mBase = m.M
			if v47 == int32(0) {
				v52 = v46
			} else {
				v50 = F_clusterGetTotalSlotExportBufferMemory(m)
				mBase = m.M
				v52 = v50 + v46
			}
			v54 = *(*int64)(unsafe.Add(mBase, _consts[265]))
			v56 = v12 - v52
			if base.Ui32(v12) < base.Ui32(v56) {
				v58 = int32(0)
			} else {
				v58 = v56
			}
			v62 = base.B2i32(base.Ui64(v54) < base.Ui64(base.I64_extend_i32_u(v58+l0)))
		}
	}
	return base.B2i32(v62 == int32(0))
}
func F_hashtableRetargetIterator(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	F_hashtableCleanupIterator(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v7)
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v7
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v4)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(-1)
		if l1 == v7 {
		} else {
			if v4&int32(1) == int32(0) {
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v21
				*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = l0
			}
		}
		return
	}
}
func F_hashtableSampleEntries(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v4
	v18 = v12 + v13
	if base.Ui32(l2) < base.Ui32(v18) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = l2
	goto L3
L2:
	;
	v20 = v18
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v20
	v22 = int32(0)
	F___lock(m, int32(9116960))
	mBase = m.M
	v29 = *(*int32)(unsafe.Add(mBase, _consts[245]))
	v31 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v20 == int32(0) {
		v97 = v4
		goto L14
	} else {
		goto L15
	}
L5:
	;
	F___unlock(m, int32(9116960))
	mBase = m.M
	goto L4
L6:
	;
	v35 = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	v37 = int32(2)
	v39 = v29 + v36<<(uint(v37)%32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v42 = *(*int32)(unsafe.Add(mBase, _consts[248]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v29+v42<<(uint(v37)%32))))
	v47 = v40 + v46
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v47
	v52 = v42 + int32(1)
	if v52 == v31 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v33 = F_lcg31(m, v32)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v33
	v66 = v33
	goto L5
L8:
	;
	v54 = v35
	goto L10
L9:
	;
	v54 = v52
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[248])) = v54
	v56 = int32(0)
	v59 = v36 + int32(1)
	if v59 == v31 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v61 = v56
	goto L13
L12:
	;
	v61 = v59
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[247])) = v61
	v66 = int32(base.Ui32(v47) >> (uint(int32(1)) % 32))
	goto L5
L14:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v98 == int32(-1) {
		v107 = v97
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v75 = v66
	goto L16
L16:
	;
	v83 = int32(0)
	v85 = F_hashtableScanDefrag(m, l0, v75, int32(539), v10+int32(4), v83, v83)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v97 = v89
	goto L14
L18:
	;
	return int32(0)
L19:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if base.Ui32(v89) < base.Ui32(v20) {
		v75 = v85
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	m.G0 = v10 + int32(16)
	if base.Ui32(v107) < base.Ui32(v20) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
	if v101 != 0 {
		v107 = v97
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[295]))
	if v103 != 0 {
		v107 = v97
		goto L21
	} else {
		goto L24
	}
L24:
	;
	F_rehashStep(m, l0)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v107 = v106
	goto L21
L26:
	;
	v112 = v107
	goto L28
L27:
	;
	v112 = v20
	goto L28
L28:
	;
	return v112
}
func F_hashtableScanCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 float64
	_ = v245
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int64
	_ = v256
	var v258 int64
	_ = v258
	var v262 int64
	_ = v262
	var v282 int64
	_ = v282
	var v296 int32
	_ = v296
	var v305 int64
	_ = v305
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v310 int64
	_ = v310
	var v311 int64
	_ = v311
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int64
	_ = v330
	var v333 int64
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v415 int32
	_ = v415
	v10 = m.G0
	v12 = v10 - int32(5152)
	m.G0 = v12
	*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223]))) = int64(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v16 + int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 == int32(0) {
		F__serverAssert(m, int32(_a485), int32(_a474), int32(1060))
		mBase = m.M
		v415 = m.ExcPending
		if v415 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		switch v23&int32(15) + int32(-2) {
		case 0:
			v161 = l1
			v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v162 == int32(0) {
				v227 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				if v227&int32(15) != int32(3) {
					v363 = v161
					v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
					switch v369 & int32(7) {
					case 0:
						v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
					case 1:
						v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
						v386 = v376
					case 2:
						v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
						v386 = v379
					case 3:
						v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
						v386 = v382
					case 4:
						v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
						v386 = v385
					default:
						v386 = int32(0)
					}
					v387 = F_vectorPush(m, v365)
					mBase = m.M
					v388 = m.ExcPending
					if v388 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
						*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
						v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
						if v391 == int32(0) {
							m.G0 = v12 + int32(5152)
							return
						} else {
							v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
							v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v396 = F_vectorPush(m, v395)
							mBase = m.M
							v397 = m.ExcPending
							if v397 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
								*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
								m.G0 = v12 + int32(5152)
								return
							}
						}
					}
				} else {
					v233 = l1 + int32(16)
					v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
					v237 = v233 + v234<<(uint(int32(3))%32)
					v238 = int32(*(*int8)(unsafe.Add(mBase, uint32(v237))))
					v242 = F_sdsdup(m, v237+v238+int32(1))
					mBase = m.M
					v243 = m.ExcPending
					if v243 != 0 {
						return
					} else {
						v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						if v244 != 0 {
							v363 = v242
							v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
							switch v369 & int32(7) {
							case 0:
								v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
							case 1:
								v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
								v386 = v376
							case 2:
								v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
								v386 = v379
							case 3:
								v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
								v386 = v382
							case 4:
								v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
								v386 = v385
							default:
								v386 = int32(0)
							}
							v387 = F_vectorPush(m, v365)
							mBase = m.M
							v388 = m.ExcPending
							if v388 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
								*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
								v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
								if v391 == int32(0) {
									m.G0 = v12 + int32(5152)
									return
								} else {
									v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
									v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v396 = F_vectorPush(m, v395)
									mBase = m.M
									v397 = m.ExcPending
									if v397 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
										*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
										m.G0 = v12 + int32(5152)
										return
									}
								}
							}
						} else {
							v245 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							v252 = m.G0
							v254 = v252 - int32(16)
							m.G0 = v254
							v256 = base.I64_reinterpret_f64(v245)
							v258 = v256 & int64(4503599627370495)
							v262 = int64(base.Ui64(v256)>>(uint(int64(52))%64)) & int64(2047)
							if v262 == int64(0) {
								if base.B2i32(v258 == int64(0)) == int32(0) {
									if base.Ui64(v258) < base.Ui64(int64(4294967296)) {
										v296 = base.I32_clz(base.I32_wrap_i64(v256)) | int32(32)
									} else {
										v296 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v258) >> (uint(int64(32)) % 64))))
									}
									F___ashlti3(m, v254, v258, int64(0), v296+int32(49))
									mBase = m.M
									v305 = *(*int64)(unsafe.Add(mBase, uint32(v254+int32(8))))
									v308 = *(*int64)(unsafe.Add(mBase, uint32(v254)))
									v309 = v308
									v310 = base.I64_extend_i32_u(int32(15372) - v296)
									v311 = v305 ^ int64(281474976710656)
								} else {
									v282 = int64(0)
									v309 = v282
									v310 = v282
									v311 = v282
								}
							} else {
								if v262 == int64(2047) {
									v309 = v258 << (uint(int64(60)) % 64)
									v310 = int64(32767)
									v311 = int64(base.Ui64(v258) >> (uint(int64(4)) % 64))
								} else {
									v309 = v258 << (uint(int64(60)) % 64)
									v310 = v262 + int64(15360)
									v311 = int64(base.Ui64(v258) >> (uint(int64(4)) % 64))
								}
							}
							*(*int64)(unsafe.Add(mBase, uint32(v12))) = v309
							*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v310<<(uint(int64(48))%64) | v256&int64(-9223372036854775807-1) | v311
							m.G0 = v254 + int32(16)
							v324 = int32(0)
							v326 = v12 + int32(16)
							v330 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
							v333 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(8))))
							v335 = F_ld2string(m, v326, int32(5120), v330, v333, v324)
							mBase = m.M
							v336 = m.ExcPending
							if v336 != 0 {
								return
							} else {
								v337 = F_sdsnewlen(m, v326, v335)
								mBase = m.M
								v338 = m.ExcPending
								if v338 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223]))) = v337
									v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-1)))))
									switch v342 & int32(7) {
									case 0:
										v359 = int32(base.Ui32(v342) >> (uint(int32(3)) % 32))
									case 1:
										v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-3)))))
										v359 = v349
									case 2:
										v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337+int32(-5)))))
										v359 = v352
									case 3:
										v355 = *(*int32)(unsafe.Add(mBase, uint32(v337+int32(-9))))
										v359 = v355
									case 4:
										v358 = *(*int32)(unsafe.Add(mBase, uint32(v337+int32(-17))))
										v359 = v358
									default:
										v359 = v324
									}
									*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224]))) = v359
									v363 = v242
									v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
									switch v369 & int32(7) {
									case 0:
										v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
									case 1:
										v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
										v386 = v376
									case 2:
										v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
										v386 = v379
									case 3:
										v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
										v386 = v382
									case 4:
										v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
										v386 = v385
									default:
										v386 = int32(0)
									}
									v387 = F_vectorPush(m, v365)
									mBase = m.M
									v388 = m.ExcPending
									if v388 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
										*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
										v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
										if v391 == int32(0) {
											m.G0 = v12 + int32(5152)
											return
										} else {
											v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
											v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v396 = F_vectorPush(m, v395)
											mBase = m.M
											v397 = m.ExcPending
											if v397 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
												*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
												m.G0 = v12 + int32(5152)
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
				v165 = int32(0)
				v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+int32(-1)))))
				switch v169 & int32(7) {
				case 0:
					v186 = int32(base.Ui32(v169) >> (uint(int32(3)) % 32))
				case 1:
					v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+int32(-3)))))
					v186 = v176
				case 2:
					v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v162+int32(-5)))))
					v186 = v179
				case 3:
					v182 = *(*int32)(unsafe.Add(mBase, uint32(v162+int32(-9))))
					v186 = v182
				case 4:
					v185 = *(*int32)(unsafe.Add(mBase, uint32(v162+int32(-17))))
					v186 = v185
				default:
					v186 = v165
				}
				v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+int32(-1)))))
				switch v189 & int32(7) {
				case 0:
					v206 = int32(base.Ui32(v189) >> (uint(int32(3)) % 32))
				case 1:
					v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+int32(-3)))))
					v206 = v196
				case 2:
					v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161+int32(-5)))))
					v206 = v199
				case 3:
					v202 = *(*int32)(unsafe.Add(mBase, uint32(v161+int32(-9))))
					v206 = v202
				case 4:
					v205 = *(*int32)(unsafe.Add(mBase, uint32(v161+int32(-17))))
					v206 = v205
				default:
					v206 = v165
				}
				v207 = int32(0)
				v209 = m.G0
				v210 = int32(16)
				v211 = v209 - v210
				m.G0 = v211
				*(*int32)(unsafe.Add(mBase, uint32(v211)+12)) = v207
				v218 = F_stringmatchlen_impl(m, v162, v186, v161, v206, v207, v211+int32(12), v207)
				mBase = m.M
				m.G0 = v211 + v210
				if v218 == int32(0) {
					m.G0 = v12 + int32(5152)
					return
				} else {
					v227 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					if v227&int32(15) != int32(3) {
						v363 = v161
						v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
						switch v369 & int32(7) {
						case 0:
							v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
						case 1:
							v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
							v386 = v376
						case 2:
							v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
							v386 = v379
						case 3:
							v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
							v386 = v382
						case 4:
							v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
							v386 = v385
						default:
							v386 = int32(0)
						}
						v387 = F_vectorPush(m, v365)
						mBase = m.M
						v388 = m.ExcPending
						if v388 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
							*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
							v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
							if v391 == int32(0) {
								m.G0 = v12 + int32(5152)
								return
							} else {
								v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
								v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v396 = F_vectorPush(m, v395)
								mBase = m.M
								v397 = m.ExcPending
								if v397 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
									*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
									m.G0 = v12 + int32(5152)
									return
								}
							}
						}
					} else {
						v233 = l1 + int32(16)
						v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
						v237 = v233 + v234<<(uint(int32(3))%32)
						v238 = int32(*(*int8)(unsafe.Add(mBase, uint32(v237))))
						v242 = F_sdsdup(m, v237+v238+int32(1))
						mBase = m.M
						v243 = m.ExcPending
						if v243 != 0 {
							return
						} else {
							v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							if v244 != 0 {
								v363 = v242
								v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
								switch v369 & int32(7) {
								case 0:
									v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
								case 1:
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
									v386 = v376
								case 2:
									v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
									v386 = v379
								case 3:
									v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
									v386 = v382
								case 4:
									v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
									v386 = v385
								default:
									v386 = int32(0)
								}
								v387 = F_vectorPush(m, v365)
								mBase = m.M
								v388 = m.ExcPending
								if v388 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
									*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
									v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
									if v391 == int32(0) {
										m.G0 = v12 + int32(5152)
										return
									} else {
										v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
										v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v396 = F_vectorPush(m, v395)
										mBase = m.M
										v397 = m.ExcPending
										if v397 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
											*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
											m.G0 = v12 + int32(5152)
											return
										}
									}
								}
							} else {
								v245 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								v252 = m.G0
								v254 = v252 - int32(16)
								m.G0 = v254
								v256 = base.I64_reinterpret_f64(v245)
								v258 = v256 & int64(4503599627370495)
								v262 = int64(base.Ui64(v256)>>(uint(int64(52))%64)) & int64(2047)
								if v262 == int64(0) {
									if base.B2i32(v258 == int64(0)) == int32(0) {
										if base.Ui64(v258) < base.Ui64(int64(4294967296)) {
											v296 = base.I32_clz(base.I32_wrap_i64(v256)) | int32(32)
										} else {
											v296 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v258) >> (uint(int64(32)) % 64))))
										}
										F___ashlti3(m, v254, v258, int64(0), v296+int32(49))
										mBase = m.M
										v305 = *(*int64)(unsafe.Add(mBase, uint32(v254+int32(8))))
										v308 = *(*int64)(unsafe.Add(mBase, uint32(v254)))
										v309 = v308
										v310 = base.I64_extend_i32_u(int32(15372) - v296)
										v311 = v305 ^ int64(281474976710656)
									} else {
										v282 = int64(0)
										v309 = v282
										v310 = v282
										v311 = v282
									}
								} else {
									if v262 == int64(2047) {
										v309 = v258 << (uint(int64(60)) % 64)
										v310 = int64(32767)
										v311 = int64(base.Ui64(v258) >> (uint(int64(4)) % 64))
									} else {
										v309 = v258 << (uint(int64(60)) % 64)
										v310 = v262 + int64(15360)
										v311 = int64(base.Ui64(v258) >> (uint(int64(4)) % 64))
									}
								}
								*(*int64)(unsafe.Add(mBase, uint32(v12))) = v309
								*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v310<<(uint(int64(48))%64) | v256&int64(-9223372036854775807-1) | v311
								m.G0 = v254 + int32(16)
								v324 = int32(0)
								v326 = v12 + int32(16)
								v330 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
								v333 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(8))))
								v335 = F_ld2string(m, v326, int32(5120), v330, v333, v324)
								mBase = m.M
								v336 = m.ExcPending
								if v336 != 0 {
									return
								} else {
									v337 = F_sdsnewlen(m, v326, v335)
									mBase = m.M
									v338 = m.ExcPending
									if v338 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223]))) = v337
										v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-1)))))
										switch v342 & int32(7) {
										case 0:
											v359 = int32(base.Ui32(v342) >> (uint(int32(3)) % 32))
										case 1:
											v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-3)))))
											v359 = v349
										case 2:
											v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337+int32(-5)))))
											v359 = v352
										case 3:
											v355 = *(*int32)(unsafe.Add(mBase, uint32(v337+int32(-9))))
											v359 = v355
										case 4:
											v358 = *(*int32)(unsafe.Add(mBase, uint32(v337+int32(-17))))
											v359 = v358
										default:
											v359 = v324
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224]))) = v359
										v363 = v242
										v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
										switch v369 & int32(7) {
										case 0:
											v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
										case 1:
											v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
											v386 = v376
										case 2:
											v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
											v386 = v379
										case 3:
											v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
											v386 = v382
										case 4:
											v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
											v386 = v385
										default:
											v386 = int32(0)
										}
										v387 = F_vectorPush(m, v365)
										mBase = m.M
										v388 = m.ExcPending
										if v388 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
											*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
											v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
											if v391 == int32(0) {
												m.G0 = v12 + int32(5152)
												return
											} else {
												v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
												v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v396 = F_vectorPush(m, v395)
												mBase = m.M
												v397 = m.ExcPending
												if v397 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
													*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
													m.G0 = v12 + int32(5152)
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
		case 1:
			v152 = l1 + int32(16)
			v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
			v156 = v152 + v153<<(uint(int32(3))%32)
			v157 = int32(*(*int8)(unsafe.Add(mBase, uint32(v156))))
			v161 = v156 + v157 + int32(1)
			v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v162 == int32(0) {
				v227 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				if v227&int32(15) != int32(3) {
					v363 = v161
					v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
					switch v369 & int32(7) {
					case 0:
						v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
					case 1:
						v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
						v386 = v376
					case 2:
						v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
						v386 = v379
					case 3:
						v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
						v386 = v382
					case 4:
						v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
						v386 = v385
					default:
						v386 = int32(0)
					}
					v387 = F_vectorPush(m, v365)
					mBase = m.M
					v388 = m.ExcPending
					if v388 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
						*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
						v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
						if v391 == int32(0) {
							m.G0 = v12 + int32(5152)
							return
						} else {
							v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
							v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v396 = F_vectorPush(m, v395)
							mBase = m.M
							v397 = m.ExcPending
							if v397 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
								*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
								m.G0 = v12 + int32(5152)
								return
							}
						}
					}
				} else {
					v233 = l1 + int32(16)
					v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
					v237 = v233 + v234<<(uint(int32(3))%32)
					v238 = int32(*(*int8)(unsafe.Add(mBase, uint32(v237))))
					v242 = F_sdsdup(m, v237+v238+int32(1))
					mBase = m.M
					v243 = m.ExcPending
					if v243 != 0 {
						return
					} else {
						v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						if v244 != 0 {
							v363 = v242
							v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
							switch v369 & int32(7) {
							case 0:
								v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
							case 1:
								v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
								v386 = v376
							case 2:
								v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
								v386 = v379
							case 3:
								v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
								v386 = v382
							case 4:
								v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
								v386 = v385
							default:
								v386 = int32(0)
							}
							v387 = F_vectorPush(m, v365)
							mBase = m.M
							v388 = m.ExcPending
							if v388 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
								*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
								v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
								if v391 == int32(0) {
									m.G0 = v12 + int32(5152)
									return
								} else {
									v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
									v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v396 = F_vectorPush(m, v395)
									mBase = m.M
									v397 = m.ExcPending
									if v397 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
										*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
										m.G0 = v12 + int32(5152)
										return
									}
								}
							}
						} else {
							v245 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							v252 = m.G0
							v254 = v252 - int32(16)
							m.G0 = v254
							v256 = base.I64_reinterpret_f64(v245)
							v258 = v256 & int64(4503599627370495)
							v262 = int64(base.Ui64(v256)>>(uint(int64(52))%64)) & int64(2047)
							if v262 == int64(0) {
								if base.B2i32(v258 == int64(0)) == int32(0) {
									if base.Ui64(v258) < base.Ui64(int64(4294967296)) {
										v296 = base.I32_clz(base.I32_wrap_i64(v256)) | int32(32)
									} else {
										v296 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v258) >> (uint(int64(32)) % 64))))
									}
									F___ashlti3(m, v254, v258, int64(0), v296+int32(49))
									mBase = m.M
									v305 = *(*int64)(unsafe.Add(mBase, uint32(v254+int32(8))))
									v308 = *(*int64)(unsafe.Add(mBase, uint32(v254)))
									v309 = v308
									v310 = base.I64_extend_i32_u(int32(15372) - v296)
									v311 = v305 ^ int64(281474976710656)
								} else {
									v282 = int64(0)
									v309 = v282
									v310 = v282
									v311 = v282
								}
							} else {
								if v262 == int64(2047) {
									v309 = v258 << (uint(int64(60)) % 64)
									v310 = int64(32767)
									v311 = int64(base.Ui64(v258) >> (uint(int64(4)) % 64))
								} else {
									v309 = v258 << (uint(int64(60)) % 64)
									v310 = v262 + int64(15360)
									v311 = int64(base.Ui64(v258) >> (uint(int64(4)) % 64))
								}
							}
							*(*int64)(unsafe.Add(mBase, uint32(v12))) = v309
							*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v310<<(uint(int64(48))%64) | v256&int64(-9223372036854775807-1) | v311
							m.G0 = v254 + int32(16)
							v324 = int32(0)
							v326 = v12 + int32(16)
							v330 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
							v333 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(8))))
							v335 = F_ld2string(m, v326, int32(5120), v330, v333, v324)
							mBase = m.M
							v336 = m.ExcPending
							if v336 != 0 {
								return
							} else {
								v337 = F_sdsnewlen(m, v326, v335)
								mBase = m.M
								v338 = m.ExcPending
								if v338 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223]))) = v337
									v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-1)))))
									switch v342 & int32(7) {
									case 0:
										v359 = int32(base.Ui32(v342) >> (uint(int32(3)) % 32))
									case 1:
										v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-3)))))
										v359 = v349
									case 2:
										v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337+int32(-5)))))
										v359 = v352
									case 3:
										v355 = *(*int32)(unsafe.Add(mBase, uint32(v337+int32(-9))))
										v359 = v355
									case 4:
										v358 = *(*int32)(unsafe.Add(mBase, uint32(v337+int32(-17))))
										v359 = v358
									default:
										v359 = v324
									}
									*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224]))) = v359
									v363 = v242
									v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
									switch v369 & int32(7) {
									case 0:
										v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
									case 1:
										v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
										v386 = v376
									case 2:
										v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
										v386 = v379
									case 3:
										v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
										v386 = v382
									case 4:
										v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
										v386 = v385
									default:
										v386 = int32(0)
									}
									v387 = F_vectorPush(m, v365)
									mBase = m.M
									v388 = m.ExcPending
									if v388 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
										*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
										v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
										if v391 == int32(0) {
											m.G0 = v12 + int32(5152)
											return
										} else {
											v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
											v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v396 = F_vectorPush(m, v395)
											mBase = m.M
											v397 = m.ExcPending
											if v397 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
												*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
												m.G0 = v12 + int32(5152)
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
				v165 = int32(0)
				v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+int32(-1)))))
				switch v169 & int32(7) {
				case 0:
					v186 = int32(base.Ui32(v169) >> (uint(int32(3)) % 32))
				case 1:
					v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+int32(-3)))))
					v186 = v176
				case 2:
					v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v162+int32(-5)))))
					v186 = v179
				case 3:
					v182 = *(*int32)(unsafe.Add(mBase, uint32(v162+int32(-9))))
					v186 = v182
				case 4:
					v185 = *(*int32)(unsafe.Add(mBase, uint32(v162+int32(-17))))
					v186 = v185
				default:
					v186 = v165
				}
				v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+int32(-1)))))
				switch v189 & int32(7) {
				case 0:
					v206 = int32(base.Ui32(v189) >> (uint(int32(3)) % 32))
				case 1:
					v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+int32(-3)))))
					v206 = v196
				case 2:
					v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161+int32(-5)))))
					v206 = v199
				case 3:
					v202 = *(*int32)(unsafe.Add(mBase, uint32(v161+int32(-9))))
					v206 = v202
				case 4:
					v205 = *(*int32)(unsafe.Add(mBase, uint32(v161+int32(-17))))
					v206 = v205
				default:
					v206 = v165
				}
				v207 = int32(0)
				v209 = m.G0
				v210 = int32(16)
				v211 = v209 - v210
				m.G0 = v211
				*(*int32)(unsafe.Add(mBase, uint32(v211)+12)) = v207
				v218 = F_stringmatchlen_impl(m, v162, v186, v161, v206, v207, v211+int32(12), v207)
				mBase = m.M
				m.G0 = v211 + v210
				if v218 == int32(0) {
					m.G0 = v12 + int32(5152)
					return
				} else {
					v227 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					if v227&int32(15) != int32(3) {
						v363 = v161
						v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
						switch v369 & int32(7) {
						case 0:
							v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
						case 1:
							v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
							v386 = v376
						case 2:
							v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
							v386 = v379
						case 3:
							v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
							v386 = v382
						case 4:
							v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
							v386 = v385
						default:
							v386 = int32(0)
						}
						v387 = F_vectorPush(m, v365)
						mBase = m.M
						v388 = m.ExcPending
						if v388 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
							*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
							v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
							if v391 == int32(0) {
								m.G0 = v12 + int32(5152)
								return
							} else {
								v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
								v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v396 = F_vectorPush(m, v395)
								mBase = m.M
								v397 = m.ExcPending
								if v397 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
									*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
									m.G0 = v12 + int32(5152)
									return
								}
							}
						}
					} else {
						v233 = l1 + int32(16)
						v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
						v237 = v233 + v234<<(uint(int32(3))%32)
						v238 = int32(*(*int8)(unsafe.Add(mBase, uint32(v237))))
						v242 = F_sdsdup(m, v237+v238+int32(1))
						mBase = m.M
						v243 = m.ExcPending
						if v243 != 0 {
							return
						} else {
							v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							if v244 != 0 {
								v363 = v242
								v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
								switch v369 & int32(7) {
								case 0:
									v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
								case 1:
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
									v386 = v376
								case 2:
									v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
									v386 = v379
								case 3:
									v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
									v386 = v382
								case 4:
									v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
									v386 = v385
								default:
									v386 = int32(0)
								}
								v387 = F_vectorPush(m, v365)
								mBase = m.M
								v388 = m.ExcPending
								if v388 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
									*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
									v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
									if v391 == int32(0) {
										m.G0 = v12 + int32(5152)
										return
									} else {
										v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
										v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v396 = F_vectorPush(m, v395)
										mBase = m.M
										v397 = m.ExcPending
										if v397 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
											*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
											m.G0 = v12 + int32(5152)
											return
										}
									}
								}
							} else {
								v245 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								v252 = m.G0
								v254 = v252 - int32(16)
								m.G0 = v254
								v256 = base.I64_reinterpret_f64(v245)
								v258 = v256 & int64(4503599627370495)
								v262 = int64(base.Ui64(v256)>>(uint(int64(52))%64)) & int64(2047)
								if v262 == int64(0) {
									if base.B2i32(v258 == int64(0)) == int32(0) {
										if base.Ui64(v258) < base.Ui64(int64(4294967296)) {
											v296 = base.I32_clz(base.I32_wrap_i64(v256)) | int32(32)
										} else {
											v296 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v258) >> (uint(int64(32)) % 64))))
										}
										F___ashlti3(m, v254, v258, int64(0), v296+int32(49))
										mBase = m.M
										v305 = *(*int64)(unsafe.Add(mBase, uint32(v254+int32(8))))
										v308 = *(*int64)(unsafe.Add(mBase, uint32(v254)))
										v309 = v308
										v310 = base.I64_extend_i32_u(int32(15372) - v296)
										v311 = v305 ^ int64(281474976710656)
									} else {
										v282 = int64(0)
										v309 = v282
										v310 = v282
										v311 = v282
									}
								} else {
									if v262 == int64(2047) {
										v309 = v258 << (uint(int64(60)) % 64)
										v310 = int64(32767)
										v311 = int64(base.Ui64(v258) >> (uint(int64(4)) % 64))
									} else {
										v309 = v258 << (uint(int64(60)) % 64)
										v310 = v262 + int64(15360)
										v311 = int64(base.Ui64(v258) >> (uint(int64(4)) % 64))
									}
								}
								*(*int64)(unsafe.Add(mBase, uint32(v12))) = v309
								*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v310<<(uint(int64(48))%64) | v256&int64(-9223372036854775807-1) | v311
								m.G0 = v254 + int32(16)
								v324 = int32(0)
								v326 = v12 + int32(16)
								v330 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
								v333 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(8))))
								v335 = F_ld2string(m, v326, int32(5120), v330, v333, v324)
								mBase = m.M
								v336 = m.ExcPending
								if v336 != 0 {
									return
								} else {
									v337 = F_sdsnewlen(m, v326, v335)
									mBase = m.M
									v338 = m.ExcPending
									if v338 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223]))) = v337
										v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-1)))))
										switch v342 & int32(7) {
										case 0:
											v359 = int32(base.Ui32(v342) >> (uint(int32(3)) % 32))
										case 1:
											v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-3)))))
											v359 = v349
										case 2:
											v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337+int32(-5)))))
											v359 = v352
										case 3:
											v355 = *(*int32)(unsafe.Add(mBase, uint32(v337+int32(-9))))
											v359 = v355
										case 4:
											v358 = *(*int32)(unsafe.Add(mBase, uint32(v337+int32(-17))))
											v359 = v358
										default:
											v359 = v324
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224]))) = v359
										v363 = v242
										v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
										switch v369 & int32(7) {
										case 0:
											v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
										case 1:
											v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
											v386 = v376
										case 2:
											v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
											v386 = v379
										case 3:
											v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
											v386 = v382
										case 4:
											v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
											v386 = v385
										default:
											v386 = int32(0)
										}
										v387 = F_vectorPush(m, v365)
										mBase = m.M
										v388 = m.ExcPending
										if v388 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
											*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
											v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
											if v391 == int32(0) {
												m.G0 = v12 + int32(5152)
												return
											} else {
												v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
												v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v396 = F_vectorPush(m, v395)
												mBase = m.M
												v397 = m.ExcPending
												if v397 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
													*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
													m.G0 = v12 + int32(5152)
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
		case 2:
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if v28 != 0 {
				v161 = l1
			} else {
				v32 = v12 + int32(5144) | int32(4)
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
				v40 = v38 & int32(7)
				if v40 == int32(0) {
					switch v40 {
					case 0:
						v60 = int32(base.Ui32(v38) >> (uint(int32(3)) % 32))
					case 1:
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
						v60 = v50
					case 2:
						v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
						v60 = v53
					case 3:
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
						v60 = v56
					case 4:
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
						v60 = v59
					default:
						v60 = int32(0)
					}
					v62 = int32(1)
					v63 = F_sdsHdrSize(m, v62)
					mBase = m.M
					v64 = l1 + v60 + v63
					v66 = v64 + v62
					if v32 == int32(0) {
						v134 = v66
						v142 = v134
					} else {
						v69 = int32(0)
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v69))))
						switch v72 & int32(7) {
						case 0:
							*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(base.Ui32(v72) >> (uint(int32(3)) % 32))
							v142 = v66
						case 1:
							v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+int32(-2)))))
							*(*int32)(unsafe.Add(mBase, uint32(v32))) = v80
							v142 = v66
						case 2:
							v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64+int32(-4)))))
							*(*int32)(unsafe.Add(mBase, uint32(v32))) = v84
							v142 = v66
						case 3:
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v64+int32(-8))))
							*(*int32)(unsafe.Add(mBase, uint32(v32))) = v88
							v142 = v66
						case 4:
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v64+int32(-16))))
							v93 = v92
							*(*int32)(unsafe.Add(mBase, uint32(v32))) = v93
							v142 = v66
						default:
							v93 = v69
							*(*int32)(unsafe.Add(mBase, uint32(v32))) = v93
							v142 = v66
						}
					}
				} else {
					if v38&int32(16) != 0 {
						v95 = F_sdsAllocPtr(m, l1)
						mBase = m.M
						v97 = v95 + int32(-4)
						if v38&int32(32) == int32(0) {
							v109 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
							if v32 == int32(0) {
								v134 = v109
							} else {
								v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-1)))))
								switch v115 & int32(7) {
								case 0:
									v132 = int32(base.Ui32(v115) >> (uint(int32(3)) % 32))
								case 1:
									v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+int32(-3)))))
									v132 = v122
								case 2:
									v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109+int32(-5)))))
									v132 = v125
								case 3:
									v128 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-9))))
									v132 = v128
								case 4:
									v131 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(-17))))
									v132 = v131
								default:
									v132 = int32(0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v32))) = v132
								v134 = v109
							}
							v142 = v134
						} else {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
							if v102 != 0 {
								if v32 == int32(0) {
								} else {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v32))) = v106
								}
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
								v142 = v108
							} else {
								v142 = int32(0)
							}
						}
					} else {
						switch v40 {
						case 0:
							v60 = int32(base.Ui32(v38) >> (uint(int32(3)) % 32))
						case 1:
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
							v60 = v50
						case 2:
							v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
							v60 = v53
						case 3:
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
							v60 = v56
						case 4:
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
							v60 = v59
						default:
							v60 = int32(0)
						}
						v62 = int32(1)
						v63 = F_sdsHdrSize(m, v62)
						mBase = m.M
						v64 = l1 + v60 + v63
						v66 = v64 + v62
						if v32 == int32(0) {
							v134 = v66
							v142 = v134
						} else {
							v69 = int32(0)
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v69))))
							switch v72 & int32(7) {
							case 0:
								*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(base.Ui32(v72) >> (uint(int32(3)) % 32))
								v142 = v66
							case 1:
								v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+int32(-2)))))
								*(*int32)(unsafe.Add(mBase, uint32(v32))) = v80
								v142 = v66
							case 2:
								v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64+int32(-4)))))
								*(*int32)(unsafe.Add(mBase, uint32(v32))) = v84
								v142 = v66
							case 3:
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v64+int32(-8))))
								*(*int32)(unsafe.Add(mBase, uint32(v32))) = v88
								v142 = v66
							case 4:
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v64+int32(-16))))
								v93 = v92
								*(*int32)(unsafe.Add(mBase, uint32(v32))) = v93
								v142 = v66
							default:
								v93 = v69
								*(*int32)(unsafe.Add(mBase, uint32(v32))) = v93
								v142 = v66
							}
						}
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223]))) = v142
				v161 = l1
			}
			v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v162 == int32(0) {
				v227 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				if v227&int32(15) != int32(3) {
					v363 = v161
					v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
					switch v369 & int32(7) {
					case 0:
						v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
					case 1:
						v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
						v386 = v376
					case 2:
						v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
						v386 = v379
					case 3:
						v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
						v386 = v382
					case 4:
						v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
						v386 = v385
					default:
						v386 = int32(0)
					}
					v387 = F_vectorPush(m, v365)
					mBase = m.M
					v388 = m.ExcPending
					if v388 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
						*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
						v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
						if v391 == int32(0) {
							m.G0 = v12 + int32(5152)
							return
						} else {
							v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
							v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v396 = F_vectorPush(m, v395)
							mBase = m.M
							v397 = m.ExcPending
							if v397 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
								*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
								m.G0 = v12 + int32(5152)
								return
							}
						}
					}
				} else {
					v233 = l1 + int32(16)
					v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
					v237 = v233 + v234<<(uint(int32(3))%32)
					v238 = int32(*(*int8)(unsafe.Add(mBase, uint32(v237))))
					v242 = F_sdsdup(m, v237+v238+int32(1))
					mBase = m.M
					v243 = m.ExcPending
					if v243 != 0 {
						return
					} else {
						v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						if v244 != 0 {
							v363 = v242
							v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
							switch v369 & int32(7) {
							case 0:
								v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
							case 1:
								v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
								v386 = v376
							case 2:
								v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
								v386 = v379
							case 3:
								v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
								v386 = v382
							case 4:
								v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
								v386 = v385
							default:
								v386 = int32(0)
							}
							v387 = F_vectorPush(m, v365)
							mBase = m.M
							v388 = m.ExcPending
							if v388 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
								*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
								v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
								if v391 == int32(0) {
									m.G0 = v12 + int32(5152)
									return
								} else {
									v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
									v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v396 = F_vectorPush(m, v395)
									mBase = m.M
									v397 = m.ExcPending
									if v397 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
										*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
										m.G0 = v12 + int32(5152)
										return
									}
								}
							}
						} else {
							v245 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							v252 = m.G0
							v254 = v252 - int32(16)
							m.G0 = v254
							v256 = base.I64_reinterpret_f64(v245)
							v258 = v256 & int64(4503599627370495)
							v262 = int64(base.Ui64(v256)>>(uint(int64(52))%64)) & int64(2047)
							if v262 == int64(0) {
								if base.B2i32(v258 == int64(0)) == int32(0) {
									if base.Ui64(v258) < base.Ui64(int64(4294967296)) {
										v296 = base.I32_clz(base.I32_wrap_i64(v256)) | int32(32)
									} else {
										v296 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v258) >> (uint(int64(32)) % 64))))
									}
									F___ashlti3(m, v254, v258, int64(0), v296+int32(49))
									mBase = m.M
									v305 = *(*int64)(unsafe.Add(mBase, uint32(v254+int32(8))))
									v308 = *(*int64)(unsafe.Add(mBase, uint32(v254)))
									v309 = v308
									v310 = base.I64_extend_i32_u(int32(15372) - v296)
									v311 = v305 ^ int64(281474976710656)
								} else {
									v282 = int64(0)
									v309 = v282
									v310 = v282
									v311 = v282
								}
							} else {
								if v262 == int64(2047) {
									v309 = v258 << (uint(int64(60)) % 64)
									v310 = int64(32767)
									v311 = int64(base.Ui64(v258) >> (uint(int64(4)) % 64))
								} else {
									v309 = v258 << (uint(int64(60)) % 64)
									v310 = v262 + int64(15360)
									v311 = int64(base.Ui64(v258) >> (uint(int64(4)) % 64))
								}
							}
							*(*int64)(unsafe.Add(mBase, uint32(v12))) = v309
							*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v310<<(uint(int64(48))%64) | v256&int64(-9223372036854775807-1) | v311
							m.G0 = v254 + int32(16)
							v324 = int32(0)
							v326 = v12 + int32(16)
							v330 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
							v333 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(8))))
							v335 = F_ld2string(m, v326, int32(5120), v330, v333, v324)
							mBase = m.M
							v336 = m.ExcPending
							if v336 != 0 {
								return
							} else {
								v337 = F_sdsnewlen(m, v326, v335)
								mBase = m.M
								v338 = m.ExcPending
								if v338 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223]))) = v337
									v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-1)))))
									switch v342 & int32(7) {
									case 0:
										v359 = int32(base.Ui32(v342) >> (uint(int32(3)) % 32))
									case 1:
										v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-3)))))
										v359 = v349
									case 2:
										v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337+int32(-5)))))
										v359 = v352
									case 3:
										v355 = *(*int32)(unsafe.Add(mBase, uint32(v337+int32(-9))))
										v359 = v355
									case 4:
										v358 = *(*int32)(unsafe.Add(mBase, uint32(v337+int32(-17))))
										v359 = v358
									default:
										v359 = v324
									}
									*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224]))) = v359
									v363 = v242
									v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
									switch v369 & int32(7) {
									case 0:
										v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
									case 1:
										v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
										v386 = v376
									case 2:
										v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
										v386 = v379
									case 3:
										v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
										v386 = v382
									case 4:
										v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
										v386 = v385
									default:
										v386 = int32(0)
									}
									v387 = F_vectorPush(m, v365)
									mBase = m.M
									v388 = m.ExcPending
									if v388 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
										*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
										v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
										if v391 == int32(0) {
											m.G0 = v12 + int32(5152)
											return
										} else {
											v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
											v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v396 = F_vectorPush(m, v395)
											mBase = m.M
											v397 = m.ExcPending
											if v397 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
												*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
												m.G0 = v12 + int32(5152)
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
				v165 = int32(0)
				v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+int32(-1)))))
				switch v169 & int32(7) {
				case 0:
					v186 = int32(base.Ui32(v169) >> (uint(int32(3)) % 32))
				case 1:
					v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+int32(-3)))))
					v186 = v176
				case 2:
					v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v162+int32(-5)))))
					v186 = v179
				case 3:
					v182 = *(*int32)(unsafe.Add(mBase, uint32(v162+int32(-9))))
					v186 = v182
				case 4:
					v185 = *(*int32)(unsafe.Add(mBase, uint32(v162+int32(-17))))
					v186 = v185
				default:
					v186 = v165
				}
				v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+int32(-1)))))
				switch v189 & int32(7) {
				case 0:
					v206 = int32(base.Ui32(v189) >> (uint(int32(3)) % 32))
				case 1:
					v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+int32(-3)))))
					v206 = v196
				case 2:
					v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161+int32(-5)))))
					v206 = v199
				case 3:
					v202 = *(*int32)(unsafe.Add(mBase, uint32(v161+int32(-9))))
					v206 = v202
				case 4:
					v205 = *(*int32)(unsafe.Add(mBase, uint32(v161+int32(-17))))
					v206 = v205
				default:
					v206 = v165
				}
				v207 = int32(0)
				v209 = m.G0
				v210 = int32(16)
				v211 = v209 - v210
				m.G0 = v211
				*(*int32)(unsafe.Add(mBase, uint32(v211)+12)) = v207
				v218 = F_stringmatchlen_impl(m, v162, v186, v161, v206, v207, v211+int32(12), v207)
				mBase = m.M
				m.G0 = v211 + v210
				if v218 == int32(0) {
					m.G0 = v12 + int32(5152)
					return
				} else {
					v227 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					if v227&int32(15) != int32(3) {
						v363 = v161
						v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
						switch v369 & int32(7) {
						case 0:
							v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
						case 1:
							v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
							v386 = v376
						case 2:
							v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
							v386 = v379
						case 3:
							v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
							v386 = v382
						case 4:
							v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
							v386 = v385
						default:
							v386 = int32(0)
						}
						v387 = F_vectorPush(m, v365)
						mBase = m.M
						v388 = m.ExcPending
						if v388 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
							*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
							v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
							if v391 == int32(0) {
								m.G0 = v12 + int32(5152)
								return
							} else {
								v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
								v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v396 = F_vectorPush(m, v395)
								mBase = m.M
								v397 = m.ExcPending
								if v397 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
									*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
									m.G0 = v12 + int32(5152)
									return
								}
							}
						}
					} else {
						v233 = l1 + int32(16)
						v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
						v237 = v233 + v234<<(uint(int32(3))%32)
						v238 = int32(*(*int8)(unsafe.Add(mBase, uint32(v237))))
						v242 = F_sdsdup(m, v237+v238+int32(1))
						mBase = m.M
						v243 = m.ExcPending
						if v243 != 0 {
							return
						} else {
							v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							if v244 != 0 {
								v363 = v242
								v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
								switch v369 & int32(7) {
								case 0:
									v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
								case 1:
									v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
									v386 = v376
								case 2:
									v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
									v386 = v379
								case 3:
									v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
									v386 = v382
								case 4:
									v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
									v386 = v385
								default:
									v386 = int32(0)
								}
								v387 = F_vectorPush(m, v365)
								mBase = m.M
								v388 = m.ExcPending
								if v388 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
									*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
									v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
									if v391 == int32(0) {
										m.G0 = v12 + int32(5152)
										return
									} else {
										v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
										v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v396 = F_vectorPush(m, v395)
										mBase = m.M
										v397 = m.ExcPending
										if v397 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
											*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
											m.G0 = v12 + int32(5152)
											return
										}
									}
								}
							} else {
								v245 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								v252 = m.G0
								v254 = v252 - int32(16)
								m.G0 = v254
								v256 = base.I64_reinterpret_f64(v245)
								v258 = v256 & int64(4503599627370495)
								v262 = int64(base.Ui64(v256)>>(uint(int64(52))%64)) & int64(2047)
								if v262 == int64(0) {
									if base.B2i32(v258 == int64(0)) == int32(0) {
										if base.Ui64(v258) < base.Ui64(int64(4294967296)) {
											v296 = base.I32_clz(base.I32_wrap_i64(v256)) | int32(32)
										} else {
											v296 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v258) >> (uint(int64(32)) % 64))))
										}
										F___ashlti3(m, v254, v258, int64(0), v296+int32(49))
										mBase = m.M
										v305 = *(*int64)(unsafe.Add(mBase, uint32(v254+int32(8))))
										v308 = *(*int64)(unsafe.Add(mBase, uint32(v254)))
										v309 = v308
										v310 = base.I64_extend_i32_u(int32(15372) - v296)
										v311 = v305 ^ int64(281474976710656)
									} else {
										v282 = int64(0)
										v309 = v282
										v310 = v282
										v311 = v282
									}
								} else {
									if v262 == int64(2047) {
										v309 = v258 << (uint(int64(60)) % 64)
										v310 = int64(32767)
										v311 = int64(base.Ui64(v258) >> (uint(int64(4)) % 64))
									} else {
										v309 = v258 << (uint(int64(60)) % 64)
										v310 = v262 + int64(15360)
										v311 = int64(base.Ui64(v258) >> (uint(int64(4)) % 64))
									}
								}
								*(*int64)(unsafe.Add(mBase, uint32(v12))) = v309
								*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v310<<(uint(int64(48))%64) | v256&int64(-9223372036854775807-1) | v311
								m.G0 = v254 + int32(16)
								v324 = int32(0)
								v326 = v12 + int32(16)
								v330 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
								v333 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(8))))
								v335 = F_ld2string(m, v326, int32(5120), v330, v333, v324)
								mBase = m.M
								v336 = m.ExcPending
								if v336 != 0 {
									return
								} else {
									v337 = F_sdsnewlen(m, v326, v335)
									mBase = m.M
									v338 = m.ExcPending
									if v338 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223]))) = v337
										v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-1)))))
										switch v342 & int32(7) {
										case 0:
											v359 = int32(base.Ui32(v342) >> (uint(int32(3)) % 32))
										case 1:
											v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337+int32(-3)))))
											v359 = v349
										case 2:
											v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337+int32(-5)))))
											v359 = v352
										case 3:
											v355 = *(*int32)(unsafe.Add(mBase, uint32(v337+int32(-9))))
											v359 = v355
										case 4:
											v358 = *(*int32)(unsafe.Add(mBase, uint32(v337+int32(-17))))
											v359 = v358
										default:
											v359 = v324
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224]))) = v359
										v363 = v242
										v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-1)))))
										switch v369 & int32(7) {
										case 0:
											v386 = int32(base.Ui32(v369) >> (uint(int32(3)) % 32))
										case 1:
											v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+int32(-3)))))
											v386 = v376
										case 2:
											v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+int32(-5)))))
											v386 = v379
										case 3:
											v382 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-9))))
											v386 = v382
										case 4:
											v385 = *(*int32)(unsafe.Add(mBase, uint32(v363+int32(-17))))
											v386 = v385
										default:
											v386 = int32(0)
										}
										v387 = F_vectorPush(m, v365)
										mBase = m.M
										v388 = m.ExcPending
										if v388 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v386
											*(*int32)(unsafe.Add(mBase, uint32(v387))) = v363
											v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[223])))
											if v391 == int32(0) {
												m.G0 = v12 + int32(5152)
												return
											} else {
												v394 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[224])))
												v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v396 = F_vectorPush(m, v395)
												mBase = m.M
												v397 = m.ExcPending
												if v397 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = v394
													*(*int32)(unsafe.Add(mBase, uint32(v396))) = v391
													m.G0 = v12 + int32(5152)
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
		default:
			F__serverPanic_1(m, int32(_a474), int32(1075), int32(_a486), int32(0))
			mBase = m.M
			v149 = m.ExcPending
			if v149 != 0 {
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
}
func F_hashtableSetCanAbortShrink(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	v1 = l0
	*(*uint8)(unsafe.Add(mBase, _consts[294])) = uint8(v1)
	return
}
func F_hashtableSetType(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	return v4
}
func F_hashtableStringKeyCaseCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v37 - v39
L2:
	;
	v37 = F_tolower(m, v33)
	mBase = m.M
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v39 = F_tolower(m, v38)
	mBase = m.M
	goto L1
L3:
	;
	v7 = l0
	v8 = l1
	v9 = v5
	goto L6
L4:
	;
	v33 = int32(0)
	v34 = l1
	goto L2
L5:
	;
	v33 = v30 & int32(255)
	v34 = v29
	goto L2
L6:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v11 == int32(0) {
		v29 = v8
		v30 = v9
		goto L5
	} else {
		goto L8
	}
L7:
	;
	v29 = v23
	v30 = int32(0)
	goto L5
L8:
	;
	v15 = v9 & int32(255)
	if v15 == v11 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v22 = int32(1)
	v23 = v8 + v22
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	if v24 != 0 {
		v7 = v7 + v22
		v8 = v23
		v9 = v24
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v17 = F_tolower(m, v15)
	mBase = m.M
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v19 = F_tolower(m, v18)
	mBase = m.M
	if v17 == v19 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v29 = v8
	v30 = v21
	goto L5
L12:
	;
	goto L7
}
func F_hashtableSubcommandGetKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return v2
}
func F_hashtableTryExpand(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	v11 = int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 != int32(-1) {
		v29 = v11
		m.G0 = v7 + int32(16)
		return v29
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(l1) < base.Ui32(v15+v16) {
			v29 = v11
			m.G0 = v7 + int32(16)
			return v29
		} else {
			v21 = F_resize_1(m, l0, l1, v7+int32(12))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v29 = v21 | base.B2i32(v25 == int32(0))
				m.G0 = v7 + int32(16)
				return v29
			}
		}
	}
}
func F_hashtableTwoPhasePopDelete(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8))))
	v10 = int32(1)
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v14 = v10 << (uint(v13) % 32)
	if int32(base.Ui32(v9)>>(uint(v10)%32))&v14&int32(4095) == int32(0) {
		F__serverAssert(m, int32(_a637), int32(_a626), int32(1860))
		mBase = m.M
		v123 = m.ExcPending
		if v123 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
		v21 = int32(-1)
		v23 = int32(1)
		v30 = v9&((v14^v21)<<(uint(v23)%32))&int32(8190) | v9&int32(57345)
		*(*uint16)(unsafe.Add(mBase, uint32(v8))) = uint16(v30)
		v34 = l0 + v20<<(uint(int32(2))%32)
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v35 + v21
		v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
		v41 = v39 + v23
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v41)
		v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
		v45 = v43 + v21
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)) = uint16(v45)
		if base.I32_extend16_s(v45) <= v21 {
			F__serverAssert(m, int32(_a638), int32(_a626), int32(1416))
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v39)
			if v39&int32(65535) != 0 {
				v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
				if v78&int32(1) == int32(0) {
					v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
					v90 = v88 + int32(-1)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v90)
					if v90&int32(65535) != 0 {
						return
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v94 != int32(-1) {
							return
						} else {
							v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v102 == int32(255) {
								v106 = int32(0)
							} else {
								v106 = int32(12) << (uint(v102) % 32)
							}
							v110 = *(*int32)(unsafe.Add(mBase, _consts[295]))
							if v110 != 0 {
								v111 = int32(3)
							} else {
								v111 = int32(13)
							}
							if base.Ui32(v106*v111) < base.Ui32(v97*int32(100)) {
								return
							} else {
								v115 = F_resize_1(m, l0, v97, int32(0))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+26)))
					if int32(0) < v83 {
						v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
						v90 = v88 + int32(-1)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v90)
						if v90&int32(65535) != 0 {
							return
						} else {
							v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v94 != int32(-1) {
								return
							} else {
								v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v102 == int32(255) {
									v106 = int32(0)
								} else {
									v106 = int32(12) << (uint(v102) % 32)
								}
								v110 = *(*int32)(unsafe.Add(mBase, _consts[295]))
								if v110 != 0 {
									v111 = int32(3)
								} else {
									v111 = int32(13)
								}
								if base.Ui32(v106*v111) < base.Ui32(v97*int32(100)) {
									return
								} else {
									v115 = F_resize_1(m, l0, v97, int32(0))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						F_fillBucketHole(m, l0, v8, v13, v20)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
							v90 = v88 + int32(-1)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v90)
							if v90&int32(65535) != 0 {
								return
							} else {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v94 != int32(-1) {
									return
								} else {
									v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
									if v102 == int32(255) {
										v106 = int32(0)
									} else {
										v106 = int32(12) << (uint(v102) % 32)
									}
									v110 = *(*int32)(unsafe.Add(mBase, _consts[295]))
									if v110 != 0 {
										v111 = int32(3)
									} else {
										v111 = int32(13)
									}
									if base.Ui32(v106*v111) < base.Ui32(v97*int32(100)) {
										return
									} else {
										v115 = F_resize_1(m, l0, v97, int32(0))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
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
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v53 != int32(-1) {
					v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
					if v78&int32(1) == int32(0) {
						v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
						v90 = v88 + int32(-1)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v90)
						if v90&int32(65535) != 0 {
							return
						} else {
							v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v94 != int32(-1) {
								return
							} else {
								v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v102 == int32(255) {
									v106 = int32(0)
								} else {
									v106 = int32(12) << (uint(v102) % 32)
								}
								v110 = *(*int32)(unsafe.Add(mBase, _consts[295]))
								if v110 != 0 {
									v111 = int32(3)
								} else {
									v111 = int32(13)
								}
								if base.Ui32(v106*v111) < base.Ui32(v97*int32(100)) {
									return
								} else {
									v115 = F_resize_1(m, l0, v97, int32(0))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+26)))
						if int32(0) < v83 {
							v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
							v90 = v88 + int32(-1)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v90)
							if v90&int32(65535) != 0 {
								return
							} else {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v94 != int32(-1) {
									return
								} else {
									v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
									if v102 == int32(255) {
										v106 = int32(0)
									} else {
										v106 = int32(12) << (uint(v102) % 32)
									}
									v110 = *(*int32)(unsafe.Add(mBase, _consts[295]))
									if v110 != 0 {
										v111 = int32(3)
									} else {
										v111 = int32(13)
									}
									if base.Ui32(v106*v111) < base.Ui32(v97*int32(100)) {
										return
									} else {
										v115 = F_resize_1(m, l0, v97, int32(0))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						} else {
							F_fillBucketHole(m, l0, v8, v13, v20)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
								v90 = v88 + int32(-1)
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v90)
								if v90&int32(65535) != 0 {
									return
								} else {
									v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v94 != int32(-1) {
										return
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v102 == int32(255) {
											v106 = int32(0)
										} else {
											v106 = int32(12) << (uint(v102) % 32)
										}
										v110 = *(*int32)(unsafe.Add(mBase, _consts[295]))
										if v110 != 0 {
											v111 = int32(3)
										} else {
											v111 = int32(13)
										}
										if base.Ui32(v106*v111) < base.Ui32(v97*int32(100)) {
											return
										} else {
											v115 = F_resize_1(m, l0, v97, int32(0))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
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
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
					if v61 == int32(255) {
						v65 = int32(0)
					} else {
						v65 = int32(12) << (uint(v61) % 32)
					}
					v69 = *(*int32)(unsafe.Add(mBase, _consts[295]))
					if v69 != 0 {
						v70 = int32(3)
					} else {
						v70 = int32(13)
					}
					if base.Ui32(v65*v70) < base.Ui32(v56*int32(100)) {
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
						if v78&int32(1) == int32(0) {
							v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
							v90 = v88 + int32(-1)
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v90)
							if v90&int32(65535) != 0 {
								return
							} else {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v94 != int32(-1) {
									return
								} else {
									v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
									if v102 == int32(255) {
										v106 = int32(0)
									} else {
										v106 = int32(12) << (uint(v102) % 32)
									}
									v110 = *(*int32)(unsafe.Add(mBase, _consts[295]))
									if v110 != 0 {
										v111 = int32(3)
									} else {
										v111 = int32(13)
									}
									if base.Ui32(v106*v111) < base.Ui32(v97*int32(100)) {
										return
									} else {
										v115 = F_resize_1(m, l0, v97, int32(0))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						} else {
							v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+26)))
							if int32(0) < v83 {
								v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
								v90 = v88 + int32(-1)
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v90)
								if v90&int32(65535) != 0 {
									return
								} else {
									v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v94 != int32(-1) {
										return
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v102 == int32(255) {
											v106 = int32(0)
										} else {
											v106 = int32(12) << (uint(v102) % 32)
										}
										v110 = *(*int32)(unsafe.Add(mBase, _consts[295]))
										if v110 != 0 {
											v111 = int32(3)
										} else {
											v111 = int32(13)
										}
										if base.Ui32(v106*v111) < base.Ui32(v97*int32(100)) {
											return
										} else {
											v115 = F_resize_1(m, l0, v97, int32(0))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								F_fillBucketHole(m, l0, v8, v13, v20)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
									v90 = v88 + int32(-1)
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v90)
									if v90&int32(65535) != 0 {
										return
									} else {
										v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										if v94 != int32(-1) {
											return
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v102 == int32(255) {
												v106 = int32(0)
											} else {
												v106 = int32(12) << (uint(v102) % 32)
											}
											v110 = *(*int32)(unsafe.Add(mBase, _consts[295]))
											if v110 != 0 {
												v111 = int32(3)
											} else {
												v111 = int32(13)
											}
											if base.Ui32(v106*v111) < base.Ui32(v97*int32(100)) {
												return
											} else {
												v115 = F_resize_1(m, l0, v97, int32(0))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
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
						v74 = F_resize_1(m, l0, v56, int32(0))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
							if v78&int32(1) == int32(0) {
								v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
								v90 = v88 + int32(-1)
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v90)
								if v90&int32(65535) != 0 {
									return
								} else {
									v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v94 != int32(-1) {
										return
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v102 == int32(255) {
											v106 = int32(0)
										} else {
											v106 = int32(12) << (uint(v102) % 32)
										}
										v110 = *(*int32)(unsafe.Add(mBase, _consts[295]))
										if v110 != 0 {
											v111 = int32(3)
										} else {
											v111 = int32(13)
										}
										if base.Ui32(v106*v111) < base.Ui32(v97*int32(100)) {
											return
										} else {
											v115 = F_resize_1(m, l0, v97, int32(0))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+26)))
								if int32(0) < v83 {
									v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
									v90 = v88 + int32(-1)
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v90)
									if v90&int32(65535) != 0 {
										return
									} else {
										v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										if v94 != int32(-1) {
											return
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v102 == int32(255) {
												v106 = int32(0)
											} else {
												v106 = int32(12) << (uint(v102) % 32)
											}
											v110 = *(*int32)(unsafe.Add(mBase, _consts[295]))
											if v110 != 0 {
												v111 = int32(3)
											} else {
												v111 = int32(13)
											}
											if base.Ui32(v106*v111) < base.Ui32(v97*int32(100)) {
												return
											} else {
												v115 = F_resize_1(m, l0, v97, int32(0))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								} else {
									F_fillBucketHole(m, l0, v8, v13, v20)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
										v90 = v88 + int32(-1)
										*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v90)
										if v90&int32(65535) != 0 {
											return
										} else {
											v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											if v94 != int32(-1) {
												return
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v102 == int32(255) {
													v106 = int32(0)
												} else {
													v106 = int32(12) << (uint(v102) % 32)
												}
												v110 = *(*int32)(unsafe.Add(mBase, _consts[295]))
												if v110 != 0 {
													v111 = int32(3)
												} else {
													v111 = int32(13)
												}
												if base.Ui32(v106*v111) < base.Ui32(v97*int32(100)) {
													return
												} else {
													v115 = F_resize_1(m, l0, v97, int32(0))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
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
	}
}
