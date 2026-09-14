package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_hashtableAdd(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_hashtableAddOrFind(m, l0, l1, int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_hashtableAddOrFind(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
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
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v258 int64
	_ = v258
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v14 == int32(0) {
		v22 = l1
		v23 = v13
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v22
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
		if v25 == int32(0) {
			v31 = v11 + int32(12)
			v32 = int32(4)
			v33 = int32(_a_F_hashtableAddOrFind_0)
			v41 = *(*int64)(unsafe.Add(mBase, _c_F_hashtableAddOrFind[0]))
			v43 = v41 ^ int64(8317987319222330741)
			v44 = *(*int64)(unsafe.Add(mBase, _c_F_hashtableAddOrFind[1]))
			v46 = v44 ^ int64(7237128888997146477)
			v48 = v41 ^ int64(7816392313619706465)
			v50 = v44 ^ int64(8387220255154660723)
			v55 = v11 + int32(16) - v32
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
			v201 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v201
			v206 = F_findBucket_1(m, l0, v200, v22, v11+int32(4), v201)
			mBase = m.M
			v207 = m.ExcPending
			if v207 != 0 {
				return int32(0)
			} else {
				if v206 == int32(0) {
					v220 = F_hashtableExpandIfNeeded(m, l0)
					mBase = m.M
					v221 = m.ExcPending
					if v221 != 0 {
						return int32(0)
					} else {
						v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v222 == int32(-1) {
							v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
							mBase = m.M
							v237 = m.ExcPending
							if v237 != 0 {
								return int32(0)
							} else {
								v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
								v239 = int32(2)
								v242 = int32(16)
								*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
								v245 = int32(1)
								v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
								v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
								*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
								v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
								*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
								v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
								v265 = l0 + v260<<(uint(v239)%32) + v242
								v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
								*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
								m.G0 = v11 + int32(16)
								return base.B2i32(v206 == int32(0))
							}
						} else {
							v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
							if v225 != 0 {
								v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
								mBase = m.M
								v237 = m.ExcPending
								if v237 != 0 {
									return int32(0)
								} else {
									v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
									v239 = int32(2)
									v242 = int32(16)
									*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
									v245 = int32(1)
									v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
									v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
									*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
									v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
									*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
									v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
									v265 = l0 + v260<<(uint(v239)%32) + v242
									v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
									*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
									m.G0 = v11 + int32(16)
									return base.B2i32(v206 == int32(0))
								}
							} else {
								v227 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableAddOrFind[2]))
								if v227 != int32(1) {
									v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
									mBase = m.M
									v237 = m.ExcPending
									if v237 != 0 {
										return int32(0)
									} else {
										v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
										v239 = int32(2)
										v242 = int32(16)
										*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
										v245 = int32(1)
										v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
										v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
										*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
										v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
										*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
										v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
										v265 = l0 + v260<<(uint(v239)%32) + v242
										v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
										*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
										m.G0 = v11 + int32(16)
										return base.B2i32(v206 == int32(0))
									}
								} else {
									F_rehashStep(m, l0)
									mBase = m.M
									v231 = m.ExcPending
									if v231 != 0 {
										return int32(0)
									} else {
										v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
										mBase = m.M
										v237 = m.ExcPending
										if v237 != 0 {
											return int32(0)
										} else {
											v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
											v239 = int32(2)
											v242 = int32(16)
											*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
											v245 = int32(1)
											v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
											v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
											*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
											v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
											*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
											v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
											v265 = l0 + v260<<(uint(v239)%32) + v242
											v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
											*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
											m.G0 = v11 + int32(16)
											return base.B2i32(v206 == int32(0))
										}
									}
								}
							}
						}
					}
				} else {
					if l2 == int32(0) {
					} else {
						v212 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						v218 = *(*int32)(unsafe.Add(mBase, uint32(v206+v212<<(uint(int32(2))%32)+int32(16))))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v218
					}
					m.G0 = v11 + int32(16)
					return base.B2i32(v206 == int32(0))
				}
			}
		} else {
			v28 = m.T0[v25].(func(*base.Module, int32) int64)(m, v22)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v200 = v28
				v201 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v201
				v206 = F_findBucket_1(m, l0, v200, v22, v11+int32(4), v201)
				mBase = m.M
				v207 = m.ExcPending
				if v207 != 0 {
					return int32(0)
				} else {
					if v206 == int32(0) {
						v220 = F_hashtableExpandIfNeeded(m, l0)
						mBase = m.M
						v221 = m.ExcPending
						if v221 != 0 {
							return int32(0)
						} else {
							v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v222 == int32(-1) {
								v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
								mBase = m.M
								v237 = m.ExcPending
								if v237 != 0 {
									return int32(0)
								} else {
									v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
									v239 = int32(2)
									v242 = int32(16)
									*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
									v245 = int32(1)
									v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
									v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
									*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
									v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
									*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
									v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
									v265 = l0 + v260<<(uint(v239)%32) + v242
									v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
									*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
									m.G0 = v11 + int32(16)
									return base.B2i32(v206 == int32(0))
								}
							} else {
								v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
								if v225 != 0 {
									v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
									mBase = m.M
									v237 = m.ExcPending
									if v237 != 0 {
										return int32(0)
									} else {
										v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
										v239 = int32(2)
										v242 = int32(16)
										*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
										v245 = int32(1)
										v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
										v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
										*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
										v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
										*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
										v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
										v265 = l0 + v260<<(uint(v239)%32) + v242
										v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
										*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
										m.G0 = v11 + int32(16)
										return base.B2i32(v206 == int32(0))
									}
								} else {
									v227 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableAddOrFind[2]))
									if v227 != int32(1) {
										v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
										mBase = m.M
										v237 = m.ExcPending
										if v237 != 0 {
											return int32(0)
										} else {
											v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
											v239 = int32(2)
											v242 = int32(16)
											*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
											v245 = int32(1)
											v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
											v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
											*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
											v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
											*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
											v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
											v265 = l0 + v260<<(uint(v239)%32) + v242
											v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
											*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
											m.G0 = v11 + int32(16)
											return base.B2i32(v206 == int32(0))
										}
									} else {
										F_rehashStep(m, l0)
										mBase = m.M
										v231 = m.ExcPending
										if v231 != 0 {
											return int32(0)
										} else {
											v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
												return int32(0)
											} else {
												v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
												v239 = int32(2)
												v242 = int32(16)
												*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
												v245 = int32(1)
												v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
												v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
												*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
												v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
												*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
												v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
												v265 = l0 + v260<<(uint(v239)%32) + v242
												v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
												*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
												m.G0 = v11 + int32(16)
												return base.B2i32(v206 == int32(0))
											}
										}
									}
								}
							}
						}
					} else {
						if l2 == int32(0) {
						} else {
							v212 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v206+v212<<(uint(int32(2))%32)+int32(16))))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v218
						}
						m.G0 = v11 + int32(16)
						return base.B2i32(v206 == int32(0))
					}
				}
			}
		}
	} else {
		v17 = m.T0[v14].(func(*base.Module, int32) int32)(m, l1)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v22 = v17
			v23 = v21
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v22
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
			if v25 == int32(0) {
				v31 = v11 + int32(12)
				v32 = int32(4)
				v33 = int32(_a_F_hashtableAddOrFind_0)
				v41 = *(*int64)(unsafe.Add(mBase, _c_F_hashtableAddOrFind[0]))
				v43 = v41 ^ int64(8317987319222330741)
				v44 = *(*int64)(unsafe.Add(mBase, _c_F_hashtableAddOrFind[1]))
				v46 = v44 ^ int64(7237128888997146477)
				v48 = v41 ^ int64(7816392313619706465)
				v50 = v44 ^ int64(8387220255154660723)
				v55 = v11 + int32(16) - v32
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
				v201 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v201
				v206 = F_findBucket_1(m, l0, v200, v22, v11+int32(4), v201)
				mBase = m.M
				v207 = m.ExcPending
				if v207 != 0 {
					return int32(0)
				} else {
					if v206 == int32(0) {
						v220 = F_hashtableExpandIfNeeded(m, l0)
						mBase = m.M
						v221 = m.ExcPending
						if v221 != 0 {
							return int32(0)
						} else {
							v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v222 == int32(-1) {
								v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
								mBase = m.M
								v237 = m.ExcPending
								if v237 != 0 {
									return int32(0)
								} else {
									v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
									v239 = int32(2)
									v242 = int32(16)
									*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
									v245 = int32(1)
									v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
									v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
									*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
									v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
									*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
									v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
									v265 = l0 + v260<<(uint(v239)%32) + v242
									v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
									*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
									m.G0 = v11 + int32(16)
									return base.B2i32(v206 == int32(0))
								}
							} else {
								v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
								if v225 != 0 {
									v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
									mBase = m.M
									v237 = m.ExcPending
									if v237 != 0 {
										return int32(0)
									} else {
										v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
										v239 = int32(2)
										v242 = int32(16)
										*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
										v245 = int32(1)
										v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
										v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
										*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
										v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
										*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
										v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
										v265 = l0 + v260<<(uint(v239)%32) + v242
										v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
										*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
										m.G0 = v11 + int32(16)
										return base.B2i32(v206 == int32(0))
									}
								} else {
									v227 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableAddOrFind[2]))
									if v227 != int32(1) {
										v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
										mBase = m.M
										v237 = m.ExcPending
										if v237 != 0 {
											return int32(0)
										} else {
											v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
											v239 = int32(2)
											v242 = int32(16)
											*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
											v245 = int32(1)
											v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
											v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
											*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
											v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
											*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
											v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
											v265 = l0 + v260<<(uint(v239)%32) + v242
											v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
											*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
											m.G0 = v11 + int32(16)
											return base.B2i32(v206 == int32(0))
										}
									} else {
										F_rehashStep(m, l0)
										mBase = m.M
										v231 = m.ExcPending
										if v231 != 0 {
											return int32(0)
										} else {
											v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
												return int32(0)
											} else {
												v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
												v239 = int32(2)
												v242 = int32(16)
												*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
												v245 = int32(1)
												v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
												v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
												*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
												v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
												*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
												v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
												v265 = l0 + v260<<(uint(v239)%32) + v242
												v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
												*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
												m.G0 = v11 + int32(16)
												return base.B2i32(v206 == int32(0))
											}
										}
									}
								}
							}
						}
					} else {
						if l2 == int32(0) {
						} else {
							v212 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v206+v212<<(uint(int32(2))%32)+int32(16))))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v218
						}
						m.G0 = v11 + int32(16)
						return base.B2i32(v206 == int32(0))
					}
				}
			} else {
				v28 = m.T0[v25].(func(*base.Module, int32) int64)(m, v22)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v200 = v28
					v201 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v201
					v206 = F_findBucket_1(m, l0, v200, v22, v11+int32(4), v201)
					mBase = m.M
					v207 = m.ExcPending
					if v207 != 0 {
						return int32(0)
					} else {
						if v206 == int32(0) {
							v220 = F_hashtableExpandIfNeeded(m, l0)
							mBase = m.M
							v221 = m.ExcPending
							if v221 != 0 {
								return int32(0)
							} else {
								v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v222 == int32(-1) {
									v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
									mBase = m.M
									v237 = m.ExcPending
									if v237 != 0 {
										return int32(0)
									} else {
										v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
										v239 = int32(2)
										v242 = int32(16)
										*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
										v245 = int32(1)
										v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
										v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
										*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
										v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
										*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
										v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
										v265 = l0 + v260<<(uint(v239)%32) + v242
										v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
										*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
										m.G0 = v11 + int32(16)
										return base.B2i32(v206 == int32(0))
									}
								} else {
									v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
									if v225 != 0 {
										v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
										mBase = m.M
										v237 = m.ExcPending
										if v237 != 0 {
											return int32(0)
										} else {
											v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
											v239 = int32(2)
											v242 = int32(16)
											*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
											v245 = int32(1)
											v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
											v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
											*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
											v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
											*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
											v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
											v265 = l0 + v260<<(uint(v239)%32) + v242
											v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
											*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
											m.G0 = v11 + int32(16)
											return base.B2i32(v206 == int32(0))
										}
									} else {
										v227 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableAddOrFind[2]))
										if v227 != int32(1) {
											v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
												return int32(0)
											} else {
												v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
												v239 = int32(2)
												v242 = int32(16)
												*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
												v245 = int32(1)
												v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
												v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
												*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
												v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
												*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
												v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
												v265 = l0 + v260<<(uint(v239)%32) + v242
												v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
												*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
												m.G0 = v11 + int32(16)
												return base.B2i32(v206 == int32(0))
											}
										} else {
											F_rehashStep(m, l0)
											mBase = m.M
											v231 = m.ExcPending
											if v231 != 0 {
												return int32(0)
											} else {
												v236 = F_findBucketForInsert(m, l0, v200, v11+int32(12), v11+int32(8))
												mBase = m.M
												v237 = m.ExcPending
												if v237 != 0 {
													return int32(0)
												} else {
													v238 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
													v239 = int32(2)
													v242 = int32(16)
													*(*int32)(unsafe.Add(mBase, uint32(v236+v238<<(uint(v239)%32)+v242))) = l1
													v245 = int32(1)
													v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
													v252 = v245<<(uint(v238)%32)<<(uint(v245)%32)&int32(8190) | v251
													*(*uint16)(unsafe.Add(mBase, uint32(v236))) = uint16(v252)
													v258 = int64(base.Ui64(v200) >> (uint(int64(56)) % 64))
													*(*uint8)(unsafe.Add(mBase, uint32(v236+v238+v239))) = uint8(v258)
													v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
													v265 = l0 + v260<<(uint(v239)%32) + v242
													v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
													*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 + v245
													m.G0 = v11 + int32(16)
													return base.B2i32(v206 == int32(0))
												}
											}
										}
									}
								}
							}
						} else {
							if l2 == int32(0) {
							} else {
								v212 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
								v218 = *(*int32)(unsafe.Add(mBase, uint32(v206+v212<<(uint(int32(2))%32)+int32(16))))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v218
							}
							m.G0 = v11 + int32(16)
							return base.B2i32(v206 == int32(0))
						}
					}
				}
			}
		}
	}
}
func F_hashtableBuckets(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v5 == int32(255) {
		v9 = int32(0)
	} else {
		v9 = int32(1) << (uint(v5) % 32)
	}
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v12 == int32(255) {
		v16 = int32(0)
	} else {
		v16 = int32(1) << (uint(v12) % 32)
	}
	return v9 + v16
}
func F_hashtableCleanupIterator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v66 int64
	_ = v66
	var v71 int64
	_ = v71
	var v76 int64
	_ = v76
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v87 int64
	_ = v87
	var v92 int64
	_ = v92
	var v97 int64
	_ = v97
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v108 int64
	_ = v108
	var v113 int64
	_ = v113
	var v118 int64
	_ = v118
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v129 int64
	_ = v129
	var v134 int64
	_ = v134
	var v139 int64
	_ = v139
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v150 int64
	_ = v150
	var v155 int64
	_ = v155
	var v160 int64
	_ = v160
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v171 int64
	_ = v171
	var v176 int64
	_ = v176
	var v181 int64
	_ = v181
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_hashtableCleanupIterator_0), int32(_a_F_hashtableCleanupIterator_1), int32(1240))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L22
	} else {
		goto L44
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_hashtableCleanupIterator_2), int32(_a_F_hashtableCleanupIterator_1), int32(2244))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L22
	} else {
		goto L43
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_hashtableCleanupIterator_3), int32(_a_F_hashtableCleanupIterator_1), int32(1416))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L22
	} else {
		goto L42
	}
L4:
	;
	return
L5:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 != int32(-1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v190&int32(1) == int32(0) {
		goto L4
	} else {
		goto L31
	}
L7:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v14&int32(1) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v11 == int32(0) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v56 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v5)+20)))
	v57 = int64(*(*int8)(unsafe.Add(mBase, uint32(v5)+25)))
	v58 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v5)+12)))
	v59 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v5)+16)))
	v60 = int64(*(*int8)(unsafe.Add(mBase, uint32(v5)+24)))
	v61 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v5)+8)))
	v62 = int64(21)
	v66 = v61<<(uint(v62)%64) + (v61 ^ int64(-1))
	v71 = (int64(base.Ui64(v66)>>(uint(int64(24))%64)) ^ v66) * int64(265)
	v76 = (int64(base.Ui64(v71)>>(uint(int64(14))%64)) ^ v71) * v62
	goto L24
L11:
	;
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+26)))
	v20 = int32(-1)
	v21 = v19 + v20
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+26)) = uint16(v21)
	if base.I32_extend16_s(v21) <= v20 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+28)))
	v28 = v26 + int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+28)) = uint16(v28)
	if v28&int32(65535) != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v32 != int32(-1) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+24)))
	if v40 == int32(255) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v44 = int32(0)
	goto L17
L16:
	;
	v44 = int32(12) << (uint(v40) % 32)
	goto L17
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableCleanupIterator[0]))
	if v48 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v49 = int32(3)
	goto L20
L19:
	;
	v49 = int32(13)
	goto L20
L20:
	;
	if base.Ui32(v44*v49) < base.Ui32(v35*int32(100)) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v53 = F_resize_1(m, v5, v35, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return
L23:
	;
	goto L6
L24:
	;
	v82 = v60 + (int64(base.Ui64(v76)>>(uint(int64(28))%64))^v76)*int64(2147483649)
	v83 = int64(21)
	v87 = v82<<(uint(v83)%64) + (v82 ^ int64(-1))
	v92 = (int64(base.Ui64(v87)>>(uint(int64(24))%64)) ^ v87) * int64(265)
	v97 = (int64(base.Ui64(v92)>>(uint(int64(14))%64)) ^ v92) * v83
	goto L25
L25:
	;
	v103 = v59 + (int64(base.Ui64(v97)>>(uint(int64(28))%64))^v97)*int64(2147483649)
	v104 = int64(21)
	v108 = v103<<(uint(v104)%64) + (v103 ^ int64(-1))
	v113 = (int64(base.Ui64(v108)>>(uint(int64(24))%64)) ^ v108) * int64(265)
	v118 = (int64(base.Ui64(v113)>>(uint(int64(14))%64)) ^ v113) * v104
	goto L26
L26:
	;
	v124 = v58 + (int64(base.Ui64(v118)>>(uint(int64(28))%64))^v118)*int64(2147483649)
	v125 = int64(21)
	v129 = v124<<(uint(v125)%64) + (v124 ^ int64(-1))
	v134 = (int64(base.Ui64(v129)>>(uint(int64(24))%64)) ^ v129) * int64(265)
	v139 = (int64(base.Ui64(v134)>>(uint(int64(14))%64)) ^ v134) * v125
	goto L27
L27:
	;
	v145 = v57 + (int64(base.Ui64(v139)>>(uint(int64(28))%64))^v139)*int64(2147483649)
	v146 = int64(21)
	v150 = v145<<(uint(v146)%64) + (v145 ^ int64(-1))
	v155 = (int64(base.Ui64(v150)>>(uint(int64(24))%64)) ^ v150) * int64(265)
	v160 = (int64(base.Ui64(v155)>>(uint(int64(14))%64)) ^ v155) * v146
	goto L28
L28:
	;
	v166 = v56 + (int64(base.Ui64(v160)>>(uint(int64(28))%64))^v160)*int64(2147483649)
	v167 = int64(21)
	v171 = v166<<(uint(v167)%64) + (v166 ^ int64(-1))
	v176 = (int64(base.Ui64(v171)>>(uint(int64(24))%64)) ^ v171) * int64(265)
	v181 = (int64(base.Ui64(v176)>>(uint(int64(14))%64)) ^ v176) * v167
	goto L29
L29:
	;
	if v55 != (int64(base.Ui64(v181)>>(uint(int64(28))%64))^v181)*int64(2147483649) {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	goto L6
L31:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+40))
	if v196 != l0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v220 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v220
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v220
	goto L4
L33:
	;
	if v196 == int32(0) {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+40)) = v198
	goto L32
L35:
	;
	v204 = v196
	goto L37
L36:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v204)+24)) = v214
	goto L32
L37:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v204)+24))
	if v206 == l0 {
		goto L36
	} else {
		goto L39
	}
L38:
	;
	F__serverAssert(m, int32(_a_F_hashtableCleanupIterator_0), int32(_a_F_hashtableCleanupIterator_1), int32(1243))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L22
	} else {
		goto L41
	}
L39:
	;
	if v206 != 0 {
		v204 = v206
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hashtableCreate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int64
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v5 != 0 {
		v12 = m.T0[v5].(func(*base.Module) int32)(m)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = v12 + int32(44)
			v16 = F_valkey_malloc(m, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v12 != 0 {
					v23 = F__emscripten_memset_bulkmem(m, v16+int32(44), base.I32_extend8_s(int32(0)), v12)
					mBase = m.M
					v24 = v16
					v25 = v15
				} else {
					v24 = v16
					v25 = int32(44)
				}
				v27 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v27
				*(*uint16)(unsafe.Add(mBase, uint32(v24)+28)) = uint16(v27)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v24))) = l0
				v34 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v24)+32)) = v34
				*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = int32(65535)
				*(*int64)(unsafe.Add(mBase, uint32(v24)+16)) = v34
				*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v34
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v42 == v27 {
					return v24
				} else {
					m.T0[v42].(func(*base.Module, int32, int32))(m, v24, v25)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						return v24
					}
				}
			}
		}
	} else {
		v6 = int32(44)
		v8 = F_valkey_malloc(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v24 = v8
			v25 = v6
			v27 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v27
			*(*uint16)(unsafe.Add(mBase, uint32(v24)+28)) = uint16(v27)
			*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v24))) = l0
			v34 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v24)+32)) = v34
			*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = int32(65535)
			*(*int64)(unsafe.Add(mBase, uint32(v24)+16)) = v34
			*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v34
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v42 == v27 {
				return v24
			} else {
				m.T0[v42].(func(*base.Module, int32, int32))(m, v24, v25)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					return v24
				}
			}
		}
	}
}
func F_hashtableCreateIterator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	v2 = l1
	v5 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+14)) = uint8(v9)
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v5)+24)) = v9
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+15)) = uint8(v2)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(-1)
		if l0 == v9 {
		} else {
			if v2&int32(1) == int32(0) {
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+24)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v5
			}
		}
		return v5
	}
}
func F_hashtableEmpty(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
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
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v265 int32
	_ = v265
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 == int32(-1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = l0 + int32(8)
	v32 = l0 + int32(32)
	v41 = int32(0)
	v46 = int32(1)
	goto L7
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	if v21 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(-1)
	goto L1
L4:
	;
	m.T0[v21].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	goto L3
L7:
	;
	v55 = l0 + int32(24) + v41
	v56 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55))))
	if v56 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return
L9:
	;
	v265 = int32(1)
	if v46&v265 != 0 {
		v41 = v265
		v46 = int32(0)
		goto L7
	} else {
		goto L49
	}
L10:
	;
	v60 = v41 << (uint(int32(2)) % 32)
	v61 = l0 + int32(16) + v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v62 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v222 = v30 + v60
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	F_valkey_free(m, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L42
	}
L12:
	;
	v67 = v30 + v60
	v81 = int32(0)
	goto L15
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v32+v60)))
	if v64 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	if l1 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L11
L17:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v93 = v81 << (uint(int32(6)) % 32)
	v97 = v91 + v93
	goto L21
L18:
	;
	if v81&int32(65535) != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	m.T0[l1].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97))))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if v113 == int32(0) {
		v177 = v111
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v198 == int32(255) {
		goto L11
	} else {
		goto L40
	}
L23:
	;
	v178 = int32(0)
	if v177&int32(1) == v178 {
		v184 = v178
		goto L32
	} else {
		goto L33
	}
L24:
	;
	if v111&int32(8190) == int32(0) {
		v177 = v111
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v125 = int32(0)
	goto L26
L26:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97))))
	v140 = int32(1)
	if int32(base.Ui32(int32(base.Ui32(v139)>>(uint(v140)%32))&int32(4095))>>(uint(v125)%32))&v140 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97))))
	v177 = v161
	goto L23
L28:
	;
	v158 = v125 + int32(1)
	if v158 != int32(12) {
		v125 = v158
		goto L26
	} else {
		goto L31
	}
L29:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v97+int32(16)+v125<<(uint(int32(2))%32))))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
	m.T0[v154].(func(*base.Module, int32))(m, v152)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L27
L32:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v97 == v185+v93 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v97)+60))
	v184 = v183
	goto L32
L34:
	;
	if v184 != 0 {
		v97 = v184
		goto L21
	} else {
		goto L39
	}
L35:
	;
	F_valkey_free(m, v97)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+36))
	if v191 == int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	m.T0[v191].(func(*base.Module, int32, int32))(m, l0, int32(-64))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	goto L22
L40:
	;
	v202 = v81 + int32(1)
	if int32(base.Ui32(v202)>>(uint(v198)%32)) == int32(0) {
		v81 = v202
		goto L15
	} else {
		goto L41
	}
L41:
	;
	goto L16
L42:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+36))
	if v227 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v240 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v240
	v244 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v244)
	*(*int32)(unsafe.Add(mBase, uint32(v32+v60))) = v240
	goto L9
L44:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v232 == int32(255) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v236 = int32(0)
	goto L47
L46:
	;
	v236 = int32(-64) << (uint(v232) % 32)
	goto L47
L47:
	;
	m.T0[v227].(func(*base.Module, int32, int32))(m, l0, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	goto L43
L49:
	;
	goto L8
}
func F_hashtableExpandIfNeeded(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 == int32(-1) {
		v51 = int32(0)
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v54 = v52 + int32(1)
		v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v59 == int32(255) {
			v63 = v51
		} else {
			v63 = int32(12) << (uint(v59) % 32)
		}
		v67 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableExpandIfNeeded[0]))
		if v67 != 0 {
			v68 = int32(500)
		} else {
			v68 = int32(100)
		}
		if base.Ui32(v54*int32(100)) <= base.Ui32(v63*v68) {
			v76 = v51
			return v76
		} else {
			v72 = F_resize_1(m, l0, v54, int32(0))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				v76 = v72
				return v76
			}
		}
	} else {
		v11 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+25)))
		v12 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v12 <= v11 {
			v76 = int32(0)
			return v76
		} else {
			v14 = int32(0)
			v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_hashtableExpandIfNeeded[1])))
			if v16 != int32(1) {
				v76 = v14
				return v76
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v19 != 0 {
					v76 = v14
					return v76
				} else {
					v20 = int32(0)
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v24 = int32(100)
					if v11 == int32(-1) {
						v33 = v20
					} else {
						v33 = int32(6000) << (uint(v11) % 32)
					}
					if base.Ui32((v21+v22)*v24+v24) <= base.Ui32(v33) {
						v76 = v20
						return v76
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v12)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v11)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v22
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v21
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
						v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						v42 = int64(32)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = base.I64_rotl(v41, v42)
						v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = base.I64_rotl(v45, v42)
						return int32(1)
					}
				}
			}
		}
	}
}
func F_hashtableFind(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v73 int64
	_ = v73
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
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
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v156 int64
	_ = v156
	var v157 int64
	_ = v157
	var v161 int64
	_ = v161
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v177 int64
	_ = v177
	var v183 int64
	_ = v183
	var v187 int64
	_ = v187
	var v191 int64
	_ = v191
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v12 == v4-v14 {
		v219 = v4
		m.G0 = v9 + int32(16)
		return v219
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
		if v19 == int32(0) {
			v27 = v9 + int32(12)
			v28 = int32(4)
			v29 = int32(_a_F_hashtableFind_0)
			v37 = *(*int64)(unsafe.Add(mBase, _c_F_hashtableFind[0]))
			v39 = v37 ^ int64(8317987319222330741)
			v40 = *(*int64)(unsafe.Add(mBase, _c_F_hashtableFind[1]))
			v42 = v40 ^ int64(7237128888997146477)
			v44 = v37 ^ int64(7816392313619706465)
			v46 = v40 ^ int64(8387220255154660723)
			v51 = v9 + int32(16) - v28
			if v27 == v51 {
				v89 = v27
				v92 = v44
				v93 = v39
				v94 = v46
				v95 = v42
			} else {
				v53 = v27
				v56 = v44
				v57 = v39
				v58 = v46
				v59 = v42
				for {
					v63 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
					v64 = v63 ^ v58
					v65 = v64 + v56
					v66 = v57 + v59
					v69 = v66 ^ base.I64_rotl(v59, int64(13))
					v70 = v65 + v69
					v73 = v70 ^ base.I64_rotl(v69, int64(17))
					v76 = base.I64_rotl(v64, int64(16)) ^ v65
					v79 = int64(32)
					v81 = v76 + base.I64_rotl(v66, v79)
					v82 = base.I64_rotl(v76, int64(21)) ^ v81
					v84 = base.I64_rotl(v70, v79)
					v85 = v81 ^ v63
					v87 = v53 + int32(8)
					if v87 != v51 {
						v53 = v87
						v56 = v84
						v57 = v85
						v58 = v82
						v59 = v73
						continue
					} else {
						break
					}
					break
				}
				v89 = v51
				v92 = v84
				v93 = v85
				v94 = v82
				v95 = v73
			}
			v100 = base.I64_extend_i32_u(v28) << (uint(int64(56)) % 64)
			switch v28 {
			default:
				v133 = v100
			case 1:
				v130 = v100
				v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
				v133 = v130 | v131
			case 2:
				v125 = v100
				v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
				v130 = v126<<(uint(int64(8))%64) | v125
				v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
				v133 = v130 | v131
			case 3:
				v120 = v100
				v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+2)))
				v125 = v121<<(uint(int64(16))%64) | v120
				v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
				v130 = v126<<(uint(int64(8))%64) | v125
				v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
				v133 = v130 | v131
			case 4:
				v115 = v100
				v116 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+3)))
				v120 = v116<<(uint(int64(24))%64) | v115
				v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+2)))
				v125 = v121<<(uint(int64(16))%64) | v120
				v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
				v130 = v126<<(uint(int64(8))%64) | v125
				v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
				v133 = v130 | v131
			case 5:
				v110 = v100
				v111 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
				v115 = v111<<(uint(int64(32))%64) | v110
				v116 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+3)))
				v120 = v116<<(uint(int64(24))%64) | v115
				v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+2)))
				v125 = v121<<(uint(int64(16))%64) | v120
				v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
				v130 = v126<<(uint(int64(8))%64) | v125
				v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
				v133 = v130 | v131
			case 6:
				v105 = v100
				v106 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+5)))
				v110 = v106<<(uint(int64(40))%64) | v105
				v111 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
				v115 = v111<<(uint(int64(32))%64) | v110
				v116 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+3)))
				v120 = v116<<(uint(int64(24))%64) | v115
				v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+2)))
				v125 = v121<<(uint(int64(16))%64) | v120
				v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
				v130 = v126<<(uint(int64(8))%64) | v125
				v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
				v133 = v130 | v131
			case 7:
				v101 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+6)))
				v105 = v101<<(uint(int64(48))%64) | v100
				v106 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+5)))
				v110 = v106<<(uint(int64(40))%64) | v105
				v111 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
				v115 = v111<<(uint(int64(32))%64) | v110
				v116 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+3)))
				v120 = v116<<(uint(int64(24))%64) | v115
				v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+2)))
				v125 = v121<<(uint(int64(16))%64) | v120
				v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
				v130 = v126<<(uint(int64(8))%64) | v125
				v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
				v133 = v130 | v131
			}
			v134 = v133 ^ v94
			v135 = int64(16)
			v137 = v134 + v92
			v138 = base.I64_rotl(v134, v135) ^ v137
			v139 = int64(21)
			v141 = v93 + v95
			v142 = int64(32)
			v144 = v138 + base.I64_rotl(v141, v142)
			v145 = base.I64_rotl(v138, v139) ^ v144
			v148 = int64(13)
			v150 = v141 ^ base.I64_rotl(v95, v148)
			v151 = v137 + v150
			v156 = base.I64_rotl(v151, v142) ^ int64(255) + v145
			v157 = base.I64_rotl(v145, v135) ^ v156
			v161 = int64(17)
			v163 = v151 ^ base.I64_rotl(v150, v161)
			v164 = v144 ^ v133 + v163
			v167 = base.I64_rotl(v164, v142) + v157
			v168 = base.I64_rotl(v157, v139) ^ v167
			v173 = v164 ^ base.I64_rotl(v163, v148)
			v174 = v173 + v156
			v177 = base.I64_rotl(v174, v142) + v168
			v183 = base.I64_rotl(v173, v161) ^ v174
			v187 = base.I64_rotl(v183, v148) ^ (v183 + v167)
			v191 = v187 + v177
			v196 = base.I64_rotl(base.I64_rotl(v168, v135)^v177, v139) ^ base.I64_rotl(v187, v161) ^ base.I64_rotl(v191, v142) ^ v191
			v197 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v197
			v202 = F_findBucket_1(m, l0, v196, l1, v9+int32(8), v197)
			mBase = m.M
			v203 = m.ExcPending
			if v203 != 0 {
				return int32(0)
			} else {
				if l2 == int32(0) {
				} else {
					if v202 == int32(0) {
					} else {
						v208 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						v214 = *(*int32)(unsafe.Add(mBase, uint32(v202+v208<<(uint(int32(2))%32)+int32(16))))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v214
					}
				}
				v219 = base.B2i32(v202 != int32(0))
				m.G0 = v9 + int32(16)
				return v219
			}
		} else {
			v22 = m.T0[v19].(func(*base.Module, int32) int64)(m, l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v196 = v22
				v197 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v197
				v202 = F_findBucket_1(m, l0, v196, l1, v9+int32(8), v197)
				mBase = m.M
				v203 = m.ExcPending
				if v203 != 0 {
					return int32(0)
				} else {
					if l2 == int32(0) {
					} else {
						if v202 == int32(0) {
						} else {
							v208 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
							v214 = *(*int32)(unsafe.Add(mBase, uint32(v202+v208<<(uint(int32(2))%32)+int32(16))))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v214
						}
					}
					v219 = base.B2i32(v202 != int32(0))
					m.G0 = v9 + int32(16)
					return v219
				}
			}
		}
	}
}
func F_hashtableFindPositionForInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
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
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v156 int64
	_ = v156
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v172 int64
	_ = v172
	var v178 int64
	_ = v178
	var v182 int64
	_ = v182
	var v186 int64
	_ = v186
	var v191 int64
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v240 int64
	_ = v240
	var v246 int32
	_ = v246
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v14 == int32(0) {
		v22 = v10 + int32(12)
		v23 = int32(4)
		v24 = int32(_a_F_hashtableFindPositionForInsert_0)
		v32 = *(*int64)(unsafe.Add(mBase, _c_F_hashtableFindPositionForInsert[0]))
		v34 = v32 ^ int64(8317987319222330741)
		v35 = *(*int64)(unsafe.Add(mBase, _c_F_hashtableFindPositionForInsert[1]))
		v37 = v35 ^ int64(7237128888997146477)
		v39 = v32 ^ int64(7816392313619706465)
		v41 = v35 ^ int64(8387220255154660723)
		v46 = v10 + int32(16) - v23
		if v22 == v46 {
			v84 = v22
			v87 = v39
			v88 = v34
			v89 = v41
			v90 = v37
		} else {
			v48 = v22
			v51 = v39
			v52 = v34
			v53 = v41
			v54 = v37
			for {
				v58 = *(*int64)(unsafe.Add(mBase, uint32(v48)))
				v59 = v58 ^ v53
				v60 = v59 + v51
				v61 = v52 + v54
				v64 = v61 ^ base.I64_rotl(v54, int64(13))
				v65 = v60 + v64
				v68 = v65 ^ base.I64_rotl(v64, int64(17))
				v71 = base.I64_rotl(v59, int64(16)) ^ v60
				v74 = int64(32)
				v76 = v71 + base.I64_rotl(v61, v74)
				v77 = base.I64_rotl(v71, int64(21)) ^ v76
				v79 = base.I64_rotl(v65, v74)
				v80 = v76 ^ v58
				v82 = v48 + int32(8)
				if v82 != v46 {
					v48 = v82
					v51 = v79
					v52 = v80
					v53 = v77
					v54 = v68
					continue
				} else {
					break
				}
				break
			}
			v84 = v46
			v87 = v79
			v88 = v80
			v89 = v77
			v90 = v68
		}
		v95 = base.I64_extend_i32_u(v23) << (uint(int64(56)) % 64)
		switch v23 {
		default:
			v128 = v95
		case 1:
			v125 = v95
			v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
			v128 = v125 | v126
		case 2:
			v120 = v95
			v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
			v125 = v121<<(uint(int64(8))%64) | v120
			v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
			v128 = v125 | v126
		case 3:
			v115 = v95
			v116 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+2)))
			v120 = v116<<(uint(int64(16))%64) | v115
			v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
			v125 = v121<<(uint(int64(8))%64) | v120
			v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
			v128 = v125 | v126
		case 4:
			v110 = v95
			v111 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+3)))
			v115 = v111<<(uint(int64(24))%64) | v110
			v116 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+2)))
			v120 = v116<<(uint(int64(16))%64) | v115
			v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
			v125 = v121<<(uint(int64(8))%64) | v120
			v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
			v128 = v125 | v126
		case 5:
			v105 = v95
			v106 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+4)))
			v110 = v106<<(uint(int64(32))%64) | v105
			v111 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+3)))
			v115 = v111<<(uint(int64(24))%64) | v110
			v116 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+2)))
			v120 = v116<<(uint(int64(16))%64) | v115
			v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
			v125 = v121<<(uint(int64(8))%64) | v120
			v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
			v128 = v125 | v126
		case 6:
			v100 = v95
			v101 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+5)))
			v105 = v101<<(uint(int64(40))%64) | v100
			v106 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+4)))
			v110 = v106<<(uint(int64(32))%64) | v105
			v111 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+3)))
			v115 = v111<<(uint(int64(24))%64) | v110
			v116 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+2)))
			v120 = v116<<(uint(int64(16))%64) | v115
			v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
			v125 = v121<<(uint(int64(8))%64) | v120
			v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
			v128 = v125 | v126
		case 7:
			v96 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+6)))
			v100 = v96<<(uint(int64(48))%64) | v95
			v101 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+5)))
			v105 = v101<<(uint(int64(40))%64) | v100
			v106 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+4)))
			v110 = v106<<(uint(int64(32))%64) | v105
			v111 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+3)))
			v115 = v111<<(uint(int64(24))%64) | v110
			v116 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+2)))
			v120 = v116<<(uint(int64(16))%64) | v115
			v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
			v125 = v121<<(uint(int64(8))%64) | v120
			v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
			v128 = v125 | v126
		}
		v129 = v128 ^ v89
		v130 = int64(16)
		v132 = v129 + v87
		v133 = base.I64_rotl(v129, v130) ^ v132
		v134 = int64(21)
		v136 = v88 + v90
		v137 = int64(32)
		v139 = v133 + base.I64_rotl(v136, v137)
		v140 = base.I64_rotl(v133, v134) ^ v139
		v143 = int64(13)
		v145 = v136 ^ base.I64_rotl(v90, v143)
		v146 = v132 + v145
		v151 = base.I64_rotl(v146, v137) ^ int64(255) + v140
		v152 = base.I64_rotl(v140, v130) ^ v151
		v156 = int64(17)
		v158 = v146 ^ base.I64_rotl(v145, v156)
		v159 = v139 ^ v128 + v158
		v162 = base.I64_rotl(v159, v137) + v152
		v163 = base.I64_rotl(v152, v134) ^ v162
		v168 = v159 ^ base.I64_rotl(v158, v143)
		v169 = v168 + v151
		v172 = base.I64_rotl(v169, v137) + v163
		v178 = base.I64_rotl(v168, v156) ^ v169
		v182 = base.I64_rotl(v178, v143) ^ (v178 + v162)
		v186 = v182 + v172
		v191 = base.I64_rotl(base.I64_rotl(v163, v130)^v172, v134) ^ base.I64_rotl(v182, v156) ^ base.I64_rotl(v186, v137) ^ v186
		v195 = F_findBucket_1(m, l0, v191, l1, v10+int32(8), int32(0))
		mBase = m.M
		v196 = m.ExcPending
		if v196 != 0 {
			return int32(0)
		} else {
			if v195 == int32(0) {
				v209 = F_hashtableExpandIfNeeded(m, l0)
				mBase = m.M
				v210 = m.ExcPending
				if v210 != 0 {
					return int32(0)
				} else {
					v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v211 == int32(-1) {
						v225 = F_findBucketForInsert(m, l0, v191, v10+int32(8), v10+int32(4))
						mBase = m.M
						v226 = m.ExcPending
						if v226 != 0 {
							return int32(0)
						} else {
							v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225))))
							v228 = int32(1)
							v232 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
							if int32(base.Ui32(int32(base.Ui32(v227)>>(uint(v228)%32))&int32(4095))>>(uint(v232)%32))&v228 != 0 {
								F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_1), int32(_a_F_hashtableFindPositionForInsert_2), int32(1693))
								mBase = m.M
								v260 = m.ExcPending
								if v260 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v240 = int64(base.Ui64(v191) >> (uint(int64(56)) % 64))
								*(*uint8)(unsafe.Add(mBase, uint32(v225+v232+int32(2)))) = uint8(v240)
								if l2 == int32(0) {
									F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_3), int32(_a_F_hashtableFindPositionForInsert_2), int32(1700))
									mBase = m.M
									v266 = m.ExcPending
									if v266 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v232)
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v225
									v246 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
									*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v246)
									m.G0 = v10 + int32(16)
									return base.B2i32(v195 == int32(0))
								}
							}
						}
					} else {
						v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
						if v214 != 0 {
							v225 = F_findBucketForInsert(m, l0, v191, v10+int32(8), v10+int32(4))
							mBase = m.M
							v226 = m.ExcPending
							if v226 != 0 {
								return int32(0)
							} else {
								v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225))))
								v228 = int32(1)
								v232 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								if int32(base.Ui32(int32(base.Ui32(v227)>>(uint(v228)%32))&int32(4095))>>(uint(v232)%32))&v228 != 0 {
									F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_1), int32(_a_F_hashtableFindPositionForInsert_2), int32(1693))
									mBase = m.M
									v260 = m.ExcPending
									if v260 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v240 = int64(base.Ui64(v191) >> (uint(int64(56)) % 64))
									*(*uint8)(unsafe.Add(mBase, uint32(v225+v232+int32(2)))) = uint8(v240)
									if l2 == int32(0) {
										F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_3), int32(_a_F_hashtableFindPositionForInsert_2), int32(1700))
										mBase = m.M
										v266 = m.ExcPending
										if v266 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v232)
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v225
										v246 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
										*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v246)
										m.G0 = v10 + int32(16)
										return base.B2i32(v195 == int32(0))
									}
								}
							}
						} else {
							v216 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableFindPositionForInsert[2]))
							if v216 != int32(1) {
								v225 = F_findBucketForInsert(m, l0, v191, v10+int32(8), v10+int32(4))
								mBase = m.M
								v226 = m.ExcPending
								if v226 != 0 {
									return int32(0)
								} else {
									v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225))))
									v228 = int32(1)
									v232 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
									if int32(base.Ui32(int32(base.Ui32(v227)>>(uint(v228)%32))&int32(4095))>>(uint(v232)%32))&v228 != 0 {
										F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_1), int32(_a_F_hashtableFindPositionForInsert_2), int32(1693))
										mBase = m.M
										v260 = m.ExcPending
										if v260 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v240 = int64(base.Ui64(v191) >> (uint(int64(56)) % 64))
										*(*uint8)(unsafe.Add(mBase, uint32(v225+v232+int32(2)))) = uint8(v240)
										if l2 == int32(0) {
											F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_3), int32(_a_F_hashtableFindPositionForInsert_2), int32(1700))
											mBase = m.M
											v266 = m.ExcPending
											if v266 != 0 {
												return int32(0)
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v232)
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v225
											v246 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
											*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v246)
											m.G0 = v10 + int32(16)
											return base.B2i32(v195 == int32(0))
										}
									}
								}
							} else {
								F_rehashStep(m, l0)
								mBase = m.M
								v220 = m.ExcPending
								if v220 != 0 {
									return int32(0)
								} else {
									v225 = F_findBucketForInsert(m, l0, v191, v10+int32(8), v10+int32(4))
									mBase = m.M
									v226 = m.ExcPending
									if v226 != 0 {
										return int32(0)
									} else {
										v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225))))
										v228 = int32(1)
										v232 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
										if int32(base.Ui32(int32(base.Ui32(v227)>>(uint(v228)%32))&int32(4095))>>(uint(v232)%32))&v228 != 0 {
											F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_1), int32(_a_F_hashtableFindPositionForInsert_2), int32(1693))
											mBase = m.M
											v260 = m.ExcPending
											if v260 != 0 {
												return int32(0)
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v240 = int64(base.Ui64(v191) >> (uint(int64(56)) % 64))
											*(*uint8)(unsafe.Add(mBase, uint32(v225+v232+int32(2)))) = uint8(v240)
											if l2 == int32(0) {
												F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_3), int32(_a_F_hashtableFindPositionForInsert_2), int32(1700))
												mBase = m.M
												v266 = m.ExcPending
												if v266 != 0 {
													return int32(0)
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v232)
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v225
												v246 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
												*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v246)
												m.G0 = v10 + int32(16)
												return base.B2i32(v195 == int32(0))
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				if l3 == int32(0) {
				} else {
					v201 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
					v207 = *(*int32)(unsafe.Add(mBase, uint32(v195+v201<<(uint(int32(2))%32)+int32(16))))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v207
				}
				m.G0 = v10 + int32(16)
				return base.B2i32(v195 == int32(0))
			}
		}
	} else {
		v17 = m.T0[v14].(func(*base.Module, int32) int64)(m, l1)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v191 = v17
			v195 = F_findBucket_1(m, l0, v191, l1, v10+int32(8), int32(0))
			mBase = m.M
			v196 = m.ExcPending
			if v196 != 0 {
				return int32(0)
			} else {
				if v195 == int32(0) {
					v209 = F_hashtableExpandIfNeeded(m, l0)
					mBase = m.M
					v210 = m.ExcPending
					if v210 != 0 {
						return int32(0)
					} else {
						v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v211 == int32(-1) {
							v225 = F_findBucketForInsert(m, l0, v191, v10+int32(8), v10+int32(4))
							mBase = m.M
							v226 = m.ExcPending
							if v226 != 0 {
								return int32(0)
							} else {
								v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225))))
								v228 = int32(1)
								v232 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								if int32(base.Ui32(int32(base.Ui32(v227)>>(uint(v228)%32))&int32(4095))>>(uint(v232)%32))&v228 != 0 {
									F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_1), int32(_a_F_hashtableFindPositionForInsert_2), int32(1693))
									mBase = m.M
									v260 = m.ExcPending
									if v260 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v240 = int64(base.Ui64(v191) >> (uint(int64(56)) % 64))
									*(*uint8)(unsafe.Add(mBase, uint32(v225+v232+int32(2)))) = uint8(v240)
									if l2 == int32(0) {
										F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_3), int32(_a_F_hashtableFindPositionForInsert_2), int32(1700))
										mBase = m.M
										v266 = m.ExcPending
										if v266 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v232)
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v225
										v246 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
										*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v246)
										m.G0 = v10 + int32(16)
										return base.B2i32(v195 == int32(0))
									}
								}
							}
						} else {
							v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
							if v214 != 0 {
								v225 = F_findBucketForInsert(m, l0, v191, v10+int32(8), v10+int32(4))
								mBase = m.M
								v226 = m.ExcPending
								if v226 != 0 {
									return int32(0)
								} else {
									v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225))))
									v228 = int32(1)
									v232 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
									if int32(base.Ui32(int32(base.Ui32(v227)>>(uint(v228)%32))&int32(4095))>>(uint(v232)%32))&v228 != 0 {
										F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_1), int32(_a_F_hashtableFindPositionForInsert_2), int32(1693))
										mBase = m.M
										v260 = m.ExcPending
										if v260 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v240 = int64(base.Ui64(v191) >> (uint(int64(56)) % 64))
										*(*uint8)(unsafe.Add(mBase, uint32(v225+v232+int32(2)))) = uint8(v240)
										if l2 == int32(0) {
											F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_3), int32(_a_F_hashtableFindPositionForInsert_2), int32(1700))
											mBase = m.M
											v266 = m.ExcPending
											if v266 != 0 {
												return int32(0)
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v232)
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v225
											v246 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
											*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v246)
											m.G0 = v10 + int32(16)
											return base.B2i32(v195 == int32(0))
										}
									}
								}
							} else {
								v216 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableFindPositionForInsert[2]))
								if v216 != int32(1) {
									v225 = F_findBucketForInsert(m, l0, v191, v10+int32(8), v10+int32(4))
									mBase = m.M
									v226 = m.ExcPending
									if v226 != 0 {
										return int32(0)
									} else {
										v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225))))
										v228 = int32(1)
										v232 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
										if int32(base.Ui32(int32(base.Ui32(v227)>>(uint(v228)%32))&int32(4095))>>(uint(v232)%32))&v228 != 0 {
											F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_1), int32(_a_F_hashtableFindPositionForInsert_2), int32(1693))
											mBase = m.M
											v260 = m.ExcPending
											if v260 != 0 {
												return int32(0)
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v240 = int64(base.Ui64(v191) >> (uint(int64(56)) % 64))
											*(*uint8)(unsafe.Add(mBase, uint32(v225+v232+int32(2)))) = uint8(v240)
											if l2 == int32(0) {
												F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_3), int32(_a_F_hashtableFindPositionForInsert_2), int32(1700))
												mBase = m.M
												v266 = m.ExcPending
												if v266 != 0 {
													return int32(0)
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v232)
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v225
												v246 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
												*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v246)
												m.G0 = v10 + int32(16)
												return base.B2i32(v195 == int32(0))
											}
										}
									}
								} else {
									F_rehashStep(m, l0)
									mBase = m.M
									v220 = m.ExcPending
									if v220 != 0 {
										return int32(0)
									} else {
										v225 = F_findBucketForInsert(m, l0, v191, v10+int32(8), v10+int32(4))
										mBase = m.M
										v226 = m.ExcPending
										if v226 != 0 {
											return int32(0)
										} else {
											v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225))))
											v228 = int32(1)
											v232 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
											if int32(base.Ui32(int32(base.Ui32(v227)>>(uint(v228)%32))&int32(4095))>>(uint(v232)%32))&v228 != 0 {
												F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_1), int32(_a_F_hashtableFindPositionForInsert_2), int32(1693))
												mBase = m.M
												v260 = m.ExcPending
												if v260 != 0 {
													return int32(0)
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v240 = int64(base.Ui64(v191) >> (uint(int64(56)) % 64))
												*(*uint8)(unsafe.Add(mBase, uint32(v225+v232+int32(2)))) = uint8(v240)
												if l2 == int32(0) {
													F__serverAssert(m, int32(_a_F_hashtableFindPositionForInsert_3), int32(_a_F_hashtableFindPositionForInsert_2), int32(1700))
													mBase = m.M
													v266 = m.ExcPending
													if v266 != 0 {
														return int32(0)
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v232)
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = v225
													v246 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
													*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v246)
													m.G0 = v10 + int32(16)
													return base.B2i32(v195 == int32(0))
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					if l3 == int32(0) {
					} else {
						v201 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
						v207 = *(*int32)(unsafe.Add(mBase, uint32(v195+v201<<(uint(int32(2))%32)+int32(16))))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v207
					}
					m.G0 = v10 + int32(16)
					return base.B2i32(v195 == int32(0))
				}
			}
		}
	}
}
func F_hashtableGetHashFunctionSeed(m *base.Module) int32 {
	return int32(_a_F_hashtableGetHashFunctionSeed_0)
}
func F_hashtableGetStats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v8 = F_hashtableGetStatsHt(m, l2, int32(0), l3)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = F_hashtableGetStatsMsg(m, l0, l1, v8, l3)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
			F_valkey_free(m, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_valkey_free(m, v8)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					if v17 == int32(-1) {
						v37 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0+l1+int32(-1)))) = uint8(v37)
						return
					} else {
						if l1 == v10 {
							v37 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0+l1+int32(-1)))) = uint8(v37)
							return
						} else {
							v24 = F_hashtableGetStatsHt(m, l2, int32(1), l3)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								v26 = F_hashtableGetStatsMsg(m, l0+v10, l1-v10, v24, l3)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
									F_valkey_free(m, v28)
									mBase = m.M
									v30 = m.ExcPending
									if v30 != 0 {
										return
									} else {
										F_valkey_free(m, v24)
										mBase = m.M
										v32 = m.ExcPending
										if v32 != 0 {
											return
										} else {
											v37 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l0+l1+int32(-1)))) = uint8(v37)
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
func F_hashtableGetStatsHt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	v12 = F_valkey_calloc(m, int32(200))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = F_valkey_calloc(m, int32(32))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = l1
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v12
			*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v20
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+l1+int32(24)))))
			v29 = int32(1) << (uint(v28) % 32)
			v31 = base.B2i32(v28 == int32(255))
			if v28 == int32(255) {
				v32 = int32(0)
			} else {
				v32 = v29
			}
			*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v32
			v36 = l0 + l1<<(uint(int32(2))%32)
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(32))))
			*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v32 * int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v39
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(16))))
			*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v46
			if l2 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = int32(0)
				if v28 == int32(255) {
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32)+int32(8))))
					v58 = int32(0)
					v62 = v58
					v65 = v58
					for {
						v70 = int32(0)
						v73 = v57 + v65<<(uint(int32(6))%32)
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
						if v74&int32(1) == v70 {
							v96 = v70
						} else {
							v79 = v73
							v80 = v70
							for {
								v89 = int32(1)
								v90 = v80 + v89
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v79)+60))
								v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
								if v92&v89 != 0 {
									v79 = v91
									v80 = v90
									continue
								} else {
									break
								}
								break
							}
							v96 = v90
						}
						v105 = int32(49)
						if base.Ui32(v96) < base.Ui32(v105) {
							v108 = v96
						} else {
							v108 = v105
						}
						v111 = v12 + v108<<(uint(int32(2))%32)
						v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
						*(*int32)(unsafe.Add(mBase, uint32(v111))) = v112 + int32(1)
						if base.Ui32(v62) < base.Ui32(v96) {
							v117 = v96
						} else {
							v117 = v62
						}
						v119 = v65 + int32(1)
						if v119 != v29 {
							v62 = v117
							v65 = v119
							continue
						} else {
							break
						}
						break
					}
					*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v117
				}
			}
			return v17
		}
	}
}
func F_hashtableInitIterator(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	v3 = l2
	v4 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v4
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(-1)
	if l1 == v4 {
	} else {
		if v3&int32(1) == int32(0) {
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = l0
		}
	}
	return
}
func F_hashtableIsRehashingPaused(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+26)))
	return base.B2i32(int32(0) < v2)
}
func F_hashtableNext(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v73 int64
	_ = v73
	var v78 int64
	_ = v78
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v89 int64
	_ = v89
	var v94 int64
	_ = v94
	var v99 int64
	_ = v99
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v110 int64
	_ = v110
	var v115 int64
	_ = v115
	var v120 int64
	_ = v120
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v131 int64
	_ = v131
	var v136 int64
	_ = v136
	var v141 int64
	_ = v141
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v152 int64
	_ = v152
	var v157 int64
	_ = v157
	var v162 int64
	_ = v162
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v173 int64
	_ = v173
	var v178 int64
	_ = v178
	var v183 int64
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int64
	_ = v250
	var v254 int64
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int64
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v367 int32
	_ = v367
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v461 int32
	_ = v461
	var v480 int32
	_ = v480
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v599 int32
	_ = v599
	var v617 int32
	_ = v617
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L5
L2:
	;
	return int32(0)
L3:
	;
	F__serverAssert(m, int32(_a_F_hashtableNext_0), int32(_a_F_hashtableNext_1), int32(1209))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L39
	} else {
		goto L86
	}
L4:
	;
	return v599
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v38 != int32(-1) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	if l1 == int32(0) {
		v599 = int32(1)
		goto L4
	} else {
		goto L85
	}
L7:
	;
	v556 = int32(1)
	v560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if int32(base.Ui32(int32(base.Ui32(v543)>>(uint(v556)%32))&int32(4095))>>(uint(v560)%32))&v556 == int32(0) {
		goto L5
	} else {
		goto L79
	}
L8:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v334&int32(2) == int32(0) {
		v397 = v328
		v399 = v330
		v402 = v333
		goto L49
	} else {
		goto L50
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v324
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v324))))
	v328 = v321
	v330 = v327
	v331 = v324
	goto L8
L10:
	;
	F_hashtableCleanupIterator(m, l0)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L39
	} else {
		goto L48
	}
L11:
	;
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v216 = int32(1)
	v217 = v215 + v216
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v217)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	if v220&v216 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L12:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v41 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v43&int32(1) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	if v192 == int32(0) {
		goto L10
	} else {
		goto L23
	}
L15:
	;
	v58 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v42)+20)))
	v59 = int64(*(*int8)(unsafe.Add(mBase, uint32(v42)+25)))
	v60 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v42)+12)))
	v61 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v42)+16)))
	v62 = int64(*(*int8)(unsafe.Add(mBase, uint32(v42)+24)))
	v63 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v42)+8)))
	v64 = int64(21)
	v68 = v63<<(uint(v64)%64) + (v63 ^ int64(-1))
	v73 = (int64(base.Ui64(v68)>>(uint(int64(24))%64)) ^ v68) * int64(265)
	v78 = (int64(base.Ui64(v73)>>(uint(int64(14))%64)) ^ v73) * v64
	goto L17
L16:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+26)))
	v49 = int32(1)
	v50 = v48 + v49
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+26)) = uint16(v50)
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+28)))
	v54 = v52 + v49
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+28)) = uint16(v54)
	v56 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v42)+16)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v56
	v191 = v42
	goto L14
L17:
	;
	v84 = v62 + (int64(base.Ui64(v78)>>(uint(int64(28))%64))^v78)*int64(2147483649)
	v85 = int64(21)
	v89 = v84<<(uint(v85)%64) + (v84 ^ int64(-1))
	v94 = (int64(base.Ui64(v89)>>(uint(int64(24))%64)) ^ v89) * int64(265)
	v99 = (int64(base.Ui64(v94)>>(uint(int64(14))%64)) ^ v94) * v85
	goto L18
L18:
	;
	v105 = v61 + (int64(base.Ui64(v99)>>(uint(int64(28))%64))^v99)*int64(2147483649)
	v106 = int64(21)
	v110 = v105<<(uint(v106)%64) + (v105 ^ int64(-1))
	v115 = (int64(base.Ui64(v110)>>(uint(int64(24))%64)) ^ v110) * int64(265)
	v120 = (int64(base.Ui64(v115)>>(uint(int64(14))%64)) ^ v115) * v106
	goto L19
L19:
	;
	v126 = v60 + (int64(base.Ui64(v120)>>(uint(int64(28))%64))^v120)*int64(2147483649)
	v127 = int64(21)
	v131 = v126<<(uint(v127)%64) + (v126 ^ int64(-1))
	v136 = (int64(base.Ui64(v131)>>(uint(int64(24))%64)) ^ v131) * int64(265)
	v141 = (int64(base.Ui64(v136)>>(uint(int64(14))%64)) ^ v136) * v127
	goto L20
L20:
	;
	v147 = v59 + (int64(base.Ui64(v141)>>(uint(int64(28))%64))^v141)*int64(2147483649)
	v148 = int64(21)
	v152 = v147<<(uint(v148)%64) + (v147 ^ int64(-1))
	v157 = (int64(base.Ui64(v152)>>(uint(int64(24))%64)) ^ v152) * int64(265)
	v162 = (int64(base.Ui64(v157)>>(uint(int64(14))%64)) ^ v157) * v148
	goto L21
L21:
	;
	v168 = v58 + (int64(base.Ui64(v162)>>(uint(int64(28))%64))^v162)*int64(2147483649)
	v169 = int64(21)
	v173 = v168<<(uint(v169)%64) + (v168 ^ int64(-1))
	v178 = (int64(base.Ui64(v173)>>(uint(int64(24))%64)) ^ v173) * int64(265)
	v183 = (int64(base.Ui64(v178)>>(uint(int64(14))%64)) ^ v178) * v169
	goto L22
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = (int64(base.Ui64(v183)>>(uint(int64(28))%64)) ^ v183) * int64(2147483649)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v191 = v190
	goto L14
L23:
	;
	v195 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v195
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if v198 == int32(-1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v201 = v195
	goto L26
L25:
	;
	v201 = v198
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v201
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v191+int32(8)+v205<<(uint(int32(2))%32))))
	v210 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v210)
	v321 = v201
	v324 = v209 + v201<<(uint(int32(6))%32)
	goto L9
L27:
	;
	v239 = v217 & int32(65535)
	if base.Ui32(v239) < base.Ui32(int32(12)) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	if base.Ui32(v217&int32(65535)) < base.Ui32(int32(11)) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v229 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v229)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	if v232&int32(1) == v229 {
		v321 = v38
		v324 = v229
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v219)+60))
	v321 = v38
	v324 = v237
	goto L9
L31:
	;
	v311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219))))
	if v239 != 0 {
		v542 = v219
		v543 = v311
		goto L7
	} else {
		goto L47
	}
L32:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v244&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v277 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v277)
	v280 = v273 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v280
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+v276)+24)))
	if v283 == int32(255) {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v243)+26)))
	if v247 != int32(1) {
		v263 = v38
		v264 = v243
		v265 = v242
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v273 = v38
	v274 = v243
	v275 = v242
	v276 = v242
	goto L33
L36:
	;
	v267 = v265 & int32(255)
	v271 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v264+v267<<(uint(int32(2))%32))+16)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v271
	v273 = v263
	v274 = v264
	v275 = v265
	v276 = v267
	goto L33
L37:
	;
	v250 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v254 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v243+v242<<(uint(int32(2))%32))+16)))
	if base.Ui64(v250) <= base.Ui64(v254) {
		v263 = v38
		v264 = v243
		v265 = v242
		goto L36
	} else {
		goto L38
	}
L38:
	;
	F_compactBucketChain(m, v243, v38, v242)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	return int32(0)
L40:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v263 = v260
	v264 = v262
	v265 = v261
	goto L36
L41:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v274+v301&int32(255)<<(uint(int32(2))%32))+8))
	v321 = v300
	v324 = v307 + v300<<(uint(int32(6))%32)
	goto L9
L42:
	;
	if v275&int32(255) != 0 {
		goto L10
	} else {
		goto L45
	}
L43:
	;
	if int32(base.Ui32(v280)>>(uint(v283)%32)) == int32(0) {
		v300 = v280
		v301 = v275
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v291 == int32(-1) {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v294 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v294)
	v297 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v297
	v300 = v297
	v301 = v294
	goto L41
L47:
	;
	v328 = v38
	v330 = v311
	v331 = v219
	goto L8
L48:
	;
	v318 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v318
	v599 = v318
	goto L4
L49:
	;
	v412 = int32(1)
	v413 = v397 + v412
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	v416 = v399 & v412
	if v416 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L50:
	;
	if v330&int32(8190) == int32(0) {
		v397 = v328
		v399 = v330
		v402 = v333
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+20))
	if v344 == int32(0) {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v353 = int32(0)
	v354 = v330
	goto L53
L53:
	;
	v367 = int32(1)
	if int32(base.Ui32(int32(base.Ui32(v354)>>(uint(v367)%32))&int32(4095))>>(uint(v353)%32))&v367 == int32(0) {
		v385 = v354
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v397 = v394
	v399 = v385
	v402 = v393
	goto L49
L55:
	;
	v386 = int32(1)
	v387 = v353 + v386
	if base.Ui32(v387) < base.Ui32(int32(12)-v385&v386) {
		v353 = v387
		v354 = v385
		goto L53
	} else {
		goto L58
	}
L56:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v331+int32(16)+v353<<(uint(int32(2))%32))))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+20))
	m.T0[v381].(func(*base.Module, int32))(m, v379)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L39
	} else {
		goto L57
	}
L57:
	;
	v384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v331))))
	v385 = v384
	goto L55
L58:
	;
	goto L54
L59:
	;
	v542 = v331
	v543 = v399
	goto L7
L60:
	;
	if v433 == int32(0) {
		goto L59
	} else {
		goto L65
	}
L61:
	;
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402+v414)+24)))
	if v421 == int32(255) {
		goto L59
	} else {
		goto L63
	}
L62:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v331)+60))
	v433 = v419
	goto L60
L63:
	;
	if int32(base.Ui32(v413)>>(uint(v421)%32)) != 0 {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v402+v414<<(uint(int32(2))%32))+8))
	v433 = v428 + v413<<(uint(int32(6))%32)
	goto L60
L65:
	;
	v436 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v433))))
	v438 = v436 & int32(1)
	if v436&int32(8190) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	if v438 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	v444 = int32(12) - v438
	v461 = int32(0)
	goto L68
L68:
	;
	v480 = v461 + int32(2)
	if v480 != v444&int32(14) {
		v461 = v480
		goto L68
	} else {
		goto L70
	}
L69:
	;
	if v444&int32(1) == int32(0) {
		goto L66
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	goto L66
L72:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402+v414)+24)))
	if v508 == int32(255) {
		goto L59
	} else {
		goto L74
	}
L73:
	;
	goto L59
L74:
	;
	if v416 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v513 = v413
	goto L77
L76:
	;
	v513 = v397 + int32(2)
	goto L77
L77:
	;
	if int32(base.Ui32(v513)>>(uint(v508)%32)) != 0 {
		goto L59
	} else {
		goto L78
	}
L78:
	;
	goto L59
L79:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v566&int32(4) != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	goto L6
L81:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v570)+12))
	if v571 == int32(0) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v542+v560<<(uint(int32(2))%32))+16))
	v578 = m.T0[v571].(func(*base.Module, int32, int32) int32)(m, v569, v577)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L39
	} else {
		goto L83
	}
L83:
	;
	if v578 == int32(0) {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	goto L80
L85:
	;
	v587 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v542+v587<<(uint(int32(2))%32))+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v591
	return int32(1)
L86:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hashtableObjKeyCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	v3 = int32(0)
	v6 = F_objectGetVal(m, l0)
	mBase = m.M
	v7 = F_objectGetVal(m, l1)
	mBase = m.M
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(-1)))))
	switch v12 & int32(7) {
	case 0:
		v29 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
	case 1:
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(-3)))))
		v29 = v19
	case 2:
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6+int32(-5)))))
		v29 = v22
	case 3:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-9))))
		v29 = v25
	case 4:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-17))))
		v29 = v28
	default:
		v29 = v3
	}
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-1)))))
	switch v32 & int32(7) {
	case 0:
		v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	case 1:
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-3)))))
		v49 = v39
	case 2:
		v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(-5)))))
		v49 = v42
	case 3:
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-9))))
		v49 = v45
	case 4:
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-17))))
		v49 = v48
	default:
		v49 = v3
	}
	if v29 != v49 {
		v105 = int32(1)
	} else {
		v52 = int32(0)
		v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(-1)))))
		switch v59 & int32(7) {
		case 0:
			v76 = int32(base.Ui32(v59) >> (uint(int32(3)) % 32))
		case 1:
			v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(-3)))))
			v76 = v66
		case 2:
			v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6+int32(-5)))))
			v76 = v69
		case 3:
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-9))))
			v76 = v72
		case 4:
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-17))))
			v76 = v75
		default:
			v76 = v52
		}
		v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-1)))))
		switch v79 & int32(7) {
		case 0:
			v96 = int32(base.Ui32(v79) >> (uint(int32(3)) % 32))
		case 1:
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-3)))))
			v96 = v86
		case 2:
			v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(-5)))))
			v96 = v89
		case 3:
			v92 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-9))))
			v96 = v92
		case 4:
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-17))))
			v96 = v95
		default:
			v96 = v52
		}
		v97 = base.B2i32(base.Ui32(v76) < base.Ui32(v96))
		if base.Ui32(v76) < base.Ui32(v96) {
			v98 = v76
		} else {
			v98 = v96
		}
		v99 = F_memcmp(m, v6, v7, v98)
		mBase = m.M
		if v99 != 0 {
			v102 = v99
		} else {
			v102 = base.B2i32(base.Ui32(v96) < base.Ui32(v76)) - v97
		}
		v105 = base.B2i32(v102 != int32(0))
	}
	return v105
}
func F_hashtableObjectDestructor(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	if l0 == int32(0) {
		return
	} else {
		F_decrRefCount(m, l0)
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			return
		}
	}
}
func F_hashtableRandomEntry(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
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
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v3
	v16 = v11 + v12
	v17 = int32(12)
	if base.Ui32(v16) < base.Ui32(v17) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = v16
	goto L3
L2:
	;
	v20 = v17
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = v9
	if v16 == int32(0) {
		v96 = v3
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v97 == int32(-1) {
		v106 = v96
		goto L21
	} else {
		goto L22
	}
L5:
	;
	goto L6
L6:
	;
	v31 = int32(0)
	F___lock(m, int32(9116960))
	mBase = m.M
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableRandomEntry[0]))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableRandomEntry[1]))
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v96 = v89
	goto L4
L8:
	;
	v83 = int32(0)
	v85 = F_hashtableScanDefrag(m, l0, v75, int32(539), v9+int32(52), v83, v83)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L18
	} else {
		goto L19
	}
L9:
	;
	F___unlock(m, int32(9116960))
	mBase = m.M
	goto L8
L10:
	;
	v44 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableRandomEntry[2]))
	v46 = int32(2)
	v48 = v38 + v45<<(uint(v46)%32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableRandomEntry[3]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v38+v51<<(uint(v46)%32))))
	v56 = v49 + v55
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v56
	v61 = v51 + int32(1)
	if v61 == v40 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v42 = F_lcg31(m, v41)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v42
	v75 = v42
	goto L9
L12:
	;
	v63 = v44
	goto L14
L13:
	;
	v63 = v61
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hashtableRandomEntry[3])) = v63
	v65 = int32(0)
	v68 = v45 + int32(1)
	if v68 == v40 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v70 = v65
	goto L17
L16:
	;
	v70 = v68
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hashtableRandomEntry[2])) = v70
	v75 = int32(base.Ui32(v56) >> (uint(int32(1)) % 32))
	goto L9
L18:
	;
	return int32(0)
L19:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	if base.Ui32(v89) < base.Ui32(v20) {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	goto L7
L21:
	;
	if base.Ui32(v106) < base.Ui32(v20) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
	if v100 != 0 {
		v106 = v96
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableRandomEntry[4]))
	if v102 != 0 {
		v106 = v96
		goto L21
	} else {
		goto L24
	}
L24:
	;
	F_rehashStep(m, l0)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	v106 = v105
	goto L21
L26:
	;
	m.G0 = v9 + int32(64)
	return base.B2i32(v108 != int32(0))
L27:
	;
	v108 = v106
	goto L29
L28:
	;
	v108 = v20
	goto L29
L29:
	;
	if v108 == int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v111 = int32(0)
	F___lock(m, int32(9116960))
	mBase = m.M
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableRandomEntry[0]))
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableRandomEntry[1]))
	if v120 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v160 = base.I32_rem_u_s(v155, v108)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v9+v160<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v164
	goto L26
L32:
	;
	F___unlock(m, int32(9116960))
	mBase = m.M
	goto L31
L33:
	;
	v124 = int32(0)
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableRandomEntry[2]))
	v126 = int32(2)
	v128 = v118 + v125<<(uint(v126)%32)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableRandomEntry[3]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v118+v131<<(uint(v126)%32))))
	v136 = v129 + v135
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v136
	v141 = v131 + int32(1)
	if v141 == v120 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v122 = F_lcg31(m, v121)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v122
	v155 = v122
	goto L32
L35:
	;
	v143 = v124
	goto L37
L36:
	;
	v143 = v141
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hashtableRandomEntry[3])) = v143
	v145 = int32(0)
	v148 = v125 + int32(1)
	if v148 == v120 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v150 = v145
	goto L40
L39:
	;
	v150 = v148
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hashtableRandomEntry[2])) = v150
	v155 = int32(base.Ui32(v136) >> (uint(int32(1)) % 32))
	goto L32
}
func F_hashtableRehashMicroseconds(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v3 = int32(0)
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+26)))
	if v3 < v6 {
		v40 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v40
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableRehashMicroseconds[0]))
	if v10 != 0 {
		v40 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableRehashMicroseconds[1]))
	v14 = m.T0[v13].(func(*base.Module) int64)(m)
	mBase = m.M
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == int32(-1) {
		v40 = v11
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = v11
	goto L5
L5:
	;
	F_rehashStep(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v40 = v27
	goto L1
L7:
	;
	return int32(0)
L8:
	;
	v27 = v20 + int32(1)
	if v27&int32(127) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v35 != int32(-1) {
		v20 = v27
		goto L5
	} else {
		goto L12
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableRehashMicroseconds[1]))
	v32 = m.T0[v31].(func(*base.Module) int64)(m)
	mBase = m.M
	if base.Ui64(l1) <= base.Ui64(v32-v14) {
		v40 = v27
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	goto L6
}
func F_hashtableRehashingInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 != int32(-1) {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v16 == int32(255) {
			v20 = int32(0)
		} else {
			v20 = int32(1) << (uint(v16) % 32)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v20
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
		if v24 == int32(255) {
			v28 = int32(0)
		} else {
			v28 = int32(1) << (uint(v24) % 32)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v28
		return
	} else {
		F__serverAssert(m, int32(_a_F_hashtableRehashingInfo_0), int32(_a_F_hashtableRehashingInfo_1), int32(1440))
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
func F_hashtableRelease(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v5 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_hashtableRelease_0), int32(_a_F_hashtableRelease_1), int32(1240))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L15
	} else {
		goto L27
	}
L2:
	;
	F_hashtableEmpty(m, l0, int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L15
	} else {
		goto L18
	}
L3:
	;
	v9 = v5
	goto L4
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if v13 != v9 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L2
L6:
	;
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v37
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v41 != 0 {
		v9 = v41
		goto L4
	} else {
		goto L17
	}
L7:
	;
	if v13 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v15
	goto L6
L9:
	;
	v22 = v13
	goto L11
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v31
	goto L6
L11:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v23 == v9 {
		goto L10
	} else {
		goto L13
	}
L12:
	;
	F__serverAssert(m, int32(_a_F_hashtableRelease_0), int32(_a_F_hashtableRelease_1), int32(1243))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	if v23 != 0 {
		v22 = v23
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	return
L16:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	goto L5
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+36))
	if v50 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_valkey_free(m, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L15
	} else {
		goto L26
	}
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+40))
	if v53 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	m.T0[v61].(func(*base.Module, int32, int32))(m, l0, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L15
	} else {
		goto L25
	}
L22:
	;
	v56 = m.T0[v53].(func(*base.Module) int32)(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L15
	} else {
		goto L24
	}
L23:
	;
	v61 = v50
	v62 = int32(-44)
	goto L21
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+36))
	v61 = v60
	v62 = int32(-44) - v56
	goto L21
L25:
	;
	goto L19
L26:
	;
	return
L27:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hashtableRightsizeIfNeeded(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v4 != int32(-1) {
		v33 = F_hashtableExpandIfNeeded(m, l0)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v36 = v33
			return v36
		}
	} else {
		v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
		if v7 != 0 {
			v33 = F_hashtableExpandIfNeeded(m, l0)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v36 = v33
				return v36
			}
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v13 == int32(255) {
				v17 = int32(0)
			} else {
				v17 = int32(12) << (uint(v13) % 32)
			}
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableRightsizeIfNeeded[0]))
			if v21 != 0 {
				v22 = int32(3)
			} else {
				v22 = int32(13)
			}
			if base.Ui32(v17*v22) < base.Ui32(v8*int32(100)) {
				v33 = F_hashtableExpandIfNeeded(m, l0)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v36 = v33
					return v36
				}
			} else {
				v27 = F_resize_1(m, l0, v8, int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					if v27 != 0 {
						v36 = int32(1)
						return v36
					} else {
						v33 = F_hashtableExpandIfNeeded(m, l0)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v36 = v33
							return v36
						}
					}
				}
			}
		}
	}
}
func F_hashtableScan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = int32(0)
	v7 = F_hashtableScanDefrag(m, l0, l1, l2, l3, v5, v5)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_hashtableScanDefrag(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v90 int32
	_ = v90
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v324 int32
	_ = v324
	var v340 int32
	_ = v340
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v504 int32
	_ = v504
	var v520 int32
	_ = v520
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v670 int32
	_ = v670
	var v685 int32
	_ = v685
	var v694 int32
	_ = v694
	var v703 int32
	_ = v703
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v777 int32
	_ = v777
	var v795 int32
	_ = v795
	v7 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v21 == v7-v23 {
		v777 = v7
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_hashtableScanDefrag_0), int32(_a_F_hashtableScanDefrag_1), int32(1416))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L21
	} else {
		goto L123
	}
L2:
	;
	return v777
L3:
	;
	v27 = l0 + int32(16)
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
	v29 = int32(1)
	v30 = v28 + v29
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)) = uint16(v30)
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	v34 = v32 + v29
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v34)
	v37 = l5 & v29
	v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+24)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v39 != int32(-1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v733 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)))
	v734 = int32(-1)
	v735 = v733 + v734
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)) = uint16(v735)
	if base.I32_extend16_s(v735) <= v734 {
		goto L1
	} else {
		goto L111
	}
L5:
	;
	v270 = int32(-1)
	v272 = l0 + int32(24)
	v273 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+25)))
	v274 = base.B2i32(v38 <= v273)
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272+v274))))
	if v276 == int32(255) {
		goto L39
	} else {
		goto L40
	}
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v44 = int32(-1)
	if v38 == v44 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v50 = int32(0)
	goto L9
L8:
	;
	v50 = v44<<(uint(v38)%32) ^ v44
	goto L9
L9:
	;
	v51 = v50 & l1
	v61 = v42 + v51<<(uint(int32(6))%32)
	goto L10
L10:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61))))
	if l2 == int32(0) {
		v144 = v74
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if base.Ui32(v23) <= base.Ui32(v175) {
		goto L36
	} else {
		goto L37
	}
L12:
	;
	if v144&int32(1) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	if v74&int32(8190) == int32(0) {
		v144 = v74
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v90 = int32(0)
	goto L15
L15:
	;
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61))))
	v104 = int32(1)
	if int32(base.Ui32(int32(base.Ui32(v103)>>(uint(v104)%32))&int32(4095))>>(uint(v90)%32))&v104 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61))))
	v144 = v137
	goto L12
L17:
	;
	v134 = v90 + int32(1)
	if v134 != int32(12) {
		v90 = v134
		goto L15
	} else {
		goto L27
	}
L18:
	;
	v115 = v61 + int32(16) + v90<<(uint(int32(2))%32)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	if v117 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v37 != 0 {
		v128 = v115
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v121 = m.T0[v117].(func(*base.Module, int32, int32) int32)(m, l0, v120)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	if v121 == int32(0) {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	m.T0[l2].(func(*base.Module, int32, int32))(m, l3, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L21
	} else {
		goto L26
	}
L25:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v128 = v127
	goto L24
L26:
	;
	goto L17
L27:
	;
	goto L16
L28:
	;
	goto L11
L29:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v61)+60))
	if l4 == int32(0) {
		v172 = v161
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v172 != 0 {
		v61 = v172
		goto L10
	} else {
		goto L35
	}
L31:
	;
	if v161 == int32(0) {
		v172 = v161
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v166 = m.T0[l4].(func(*base.Module, int32) int32)(m, v161)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	if v166 == int32(0) {
		v172 = v161
		goto L30
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+60)) = v166
	v172 = v166
	goto L30
L35:
	;
	goto L28
L36:
	;
	v182 = l1 | (v50 ^ int32(-1))
	v183 = int32(24)
	v185 = int32(65280)
	v187 = int32(8)
	v197 = v182<<(uint(v183)%32) | v182&v185<<(uint(v187)%32) | (int32(base.Ui32(v182)>>(uint(v187)%32))&v185 | int32(base.Ui32(v182)>>(uint(v183)%32)))
	v198 = int32(4)
	v200 = int32(252645135)
	v206 = int32(base.Ui32(v197)>>(uint(v198)%32))&v200 | v197&v200<<(uint(v198)%32)
	v207 = int32(2)
	v209 = int32(858993459)
	v215 = int32(base.Ui32(v206)>>(uint(v207)%32))&v209 | v206&v209<<(uint(v207)%32)
	v216 = int32(1)
	v218 = int32(1431655765)
	v226 = int32(base.Ui32(v215)>>(uint(v216)%32))&v218 | v215&v218<<(uint(v216)%32) + v216
	v241 = v226<<(uint(v183)%32) | v226&v185<<(uint(v187)%32) | (int32(base.Ui32(v226)>>(uint(v187)%32))&v185 | int32(base.Ui32(v226)>>(uint(v183)%32)))
	v250 = int32(base.Ui32(v241)>>(uint(v198)%32))&v200 | v241&v200<<(uint(v198)%32)
	v259 = int32(base.Ui32(v250)>>(uint(v207)%32))&v209 | v250&v209<<(uint(v207)%32)
	v715 = int32(base.Ui32(v259)>>(uint(v216)%32))&v218 | v259&v218<<(uint(v216)%32)
	goto L4
L37:
	;
	F_compactBucketChain(m, l0, v51, int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L21
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v282 = int32(0)
	goto L41
L40:
	;
	v282 = v270<<(uint(v276)%32) ^ v270
	goto L41
L41:
	;
	v284 = int32(-1)
	v285 = base.B2i32(v273 < v38)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272+v285))))
	if v287 == int32(255) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v293 = int32(0)
	goto L44
L43:
	;
	v293 = v284<<(uint(v287)%32) ^ v284
	goto L44
L44:
	;
	v294 = v293 & l1
	if v273 < v38 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v450 = v274 << (uint(int32(2)) % 32)
	v451 = v27 + v450
	v456 = l1
	goto L76
L46:
	;
	v297 = v285 << (uint(int32(2)) % 32)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0+v297)+8))
	v303 = v27 + v297
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v311 = v299 + v294<<(uint(int32(6))%32)
	goto L49
L47:
	;
	if base.Ui32(v294) < base.Ui32(v39) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v324 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v311))))
	if l2 == int32(0) {
		v392 = v324
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	if base.Ui32(v304) <= base.Ui32(v423) {
		goto L45
	} else {
		goto L74
	}
L51:
	;
	if v392&int32(1) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L52:
	;
	if v324&int32(8190) == int32(0) {
		v392 = v324
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v340 = int32(0)
	goto L54
L54:
	;
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v311))))
	v354 = int32(1)
	if int32(base.Ui32(int32(base.Ui32(v353)>>(uint(v354)%32))&int32(4095))>>(uint(v340)%32))&v354 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v385 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v311))))
	v392 = v385
	goto L51
L56:
	;
	v382 = v340 + int32(1)
	if v382 != int32(12) {
		v340 = v382
		goto L54
	} else {
		goto L65
	}
L57:
	;
	v365 = v311 + int32(16) + v340<<(uint(int32(2))%32)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v366)+12))
	if v367 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v37 != 0 {
		v376 = v365
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	v371 = m.T0[v367].(func(*base.Module, int32, int32) int32)(m, l0, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L21
	} else {
		goto L60
	}
L60:
	;
	if v371 == int32(0) {
		goto L56
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	m.T0[l2].(func(*base.Module, int32, int32))(m, l3, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L21
	} else {
		goto L64
	}
L63:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	v376 = v375
	goto L62
L64:
	;
	goto L56
L65:
	;
	goto L55
L66:
	;
	goto L50
L67:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v311)+60))
	if l4 == int32(0) {
		v420 = v409
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if v420 != 0 {
		v311 = v420
		goto L49
	} else {
		goto L73
	}
L69:
	;
	if v409 == int32(0) {
		v420 = v409
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v414 = m.T0[l4].(func(*base.Module, int32) int32)(m, v409)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L21
	} else {
		goto L71
	}
L71:
	;
	if v414 == int32(0) {
		v420 = v409
		goto L68
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311)+60)) = v414
	v420 = v414
	goto L68
L73:
	;
	goto L66
L74:
	;
	F_compactBucketChain(m, l0, v294, v285)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L21
	} else {
		goto L75
	}
L75:
	;
	goto L45
L76:
	;
	v474 = v456 & v282
	if v38 <= v273 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v715 = v712
	goto L4
L78:
	;
	v626 = v456 | (v282 ^ int32(-1))
	v627 = int32(24)
	v629 = int32(65280)
	v631 = int32(8)
	v641 = v626<<(uint(v627)%32) | v626&v629<<(uint(v631)%32) | (int32(base.Ui32(v626)>>(uint(v631)%32))&v629 | int32(base.Ui32(v626)>>(uint(v627)%32)))
	v642 = int32(4)
	v644 = int32(252645135)
	v650 = int32(base.Ui32(v641)>>(uint(v642)%32))&v644 | v641&v644<<(uint(v642)%32)
	v651 = int32(2)
	v653 = int32(858993459)
	v659 = int32(base.Ui32(v650)>>(uint(v651)%32))&v653 | v650&v653<<(uint(v651)%32)
	v660 = int32(1)
	v662 = int32(1431655765)
	v670 = int32(base.Ui32(v659)>>(uint(v660)%32))&v662 | v659&v662<<(uint(v660)%32) + v660
	v685 = v670<<(uint(v627)%32) | v670&v629<<(uint(v631)%32) | (int32(base.Ui32(v670)>>(uint(v631)%32))&v629 | int32(base.Ui32(v670)>>(uint(v627)%32)))
	v694 = int32(base.Ui32(v685)>>(uint(v642)%32))&v644 | v685&v644<<(uint(v642)%32)
	v703 = int32(base.Ui32(v694)>>(uint(v651)%32))&v653 | v694&v653<<(uint(v651)%32)
	v712 = int32(base.Ui32(v703)>>(uint(v660)%32))&v662 | v703&v662<<(uint(v660)%32)
	if v712&(v282^v293) != 0 {
		v456 = v712
		goto L76
	} else {
		goto L110
	}
L79:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0+v450+int32(8))))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	v491 = v480 + v474<<(uint(int32(6))%32)
	goto L83
L80:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v475 == int32(-1) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	if base.Ui32(v474) < base.Ui32(v475) {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	v504 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v491))))
	if l2 == int32(0) {
		v572 = v504
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	if base.Ui32(v484) <= base.Ui32(v603) {
		goto L78
	} else {
		goto L108
	}
L85:
	;
	if v572&int32(1) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L86:
	;
	if v504&int32(8190) == int32(0) {
		v572 = v504
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v520 = int32(0)
	goto L88
L88:
	;
	v533 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v491))))
	v534 = int32(1)
	if int32(base.Ui32(int32(base.Ui32(v533)>>(uint(v534)%32))&int32(4095))>>(uint(v520)%32))&v534 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v565 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v491))))
	v572 = v565
	goto L85
L90:
	;
	v562 = v520 + int32(1)
	if v562 != int32(12) {
		v520 = v562
		goto L88
	} else {
		goto L99
	}
L91:
	;
	v545 = v491 + int32(16) + v520<<(uint(int32(2))%32)
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)+12))
	if v547 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if v37 != 0 {
		v556 = v545
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	v551 = m.T0[v547].(func(*base.Module, int32, int32) int32)(m, l0, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L21
	} else {
		goto L94
	}
L94:
	;
	if v551 == int32(0) {
		goto L90
	} else {
		goto L95
	}
L95:
	;
	goto L92
L96:
	;
	m.T0[l2].(func(*base.Module, int32, int32))(m, l3, v556)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L21
	} else {
		goto L98
	}
L97:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	v556 = v555
	goto L96
L98:
	;
	goto L90
L99:
	;
	goto L89
L100:
	;
	goto L84
L101:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v491)+60))
	if l4 == int32(0) {
		v600 = v589
		goto L102
	} else {
		goto L103
	}
L102:
	;
	if v600 != 0 {
		v491 = v600
		goto L83
	} else {
		goto L107
	}
L103:
	;
	if v589 == int32(0) {
		v600 = v589
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v594 = m.T0[l4].(func(*base.Module, int32) int32)(m, v589)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L21
	} else {
		goto L105
	}
L105:
	;
	if v594 == int32(0) {
		v600 = v589
		goto L102
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491)+60)) = v594
	v600 = v594
	goto L102
L107:
	;
	goto L100
L108:
	;
	F_compactBucketChain(m, l0, v474, v274)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L21
	} else {
		goto L109
	}
L109:
	;
	goto L78
L110:
	;
	goto L77
L111:
	;
	v740 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	v742 = v740 + int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v742)
	if v742&int32(65535) != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v777 = v715
	goto L2
L113:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v746 != int32(-1) {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v754 == int32(255) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v758 = int32(0)
	goto L117
L116:
	;
	v758 = int32(12) << (uint(v754) % 32)
	goto L117
L117:
	;
	v762 = *(*int32)(unsafe.Add(mBase, _c_F_hashtableScanDefrag[0]))
	if v762 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v763 = int32(3)
	goto L120
L119:
	;
	v763 = int32(13)
	goto L120
L120:
	;
	if base.Ui32(v758*v763) < base.Ui32(v749*int32(100)) {
		goto L112
	} else {
		goto L121
	}
L121:
	;
	v767 = F_resize_1(m, l0, v749, int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L21
	} else {
		goto L122
	}
L122:
	;
	goto L112
L123:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hashtableScanToKvstoreScanCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	m.T0[v5].(func(*base.Module, int32, int32, int32))(m, v3, l1, v4)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_hashtableSdsHash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v7 & int32(7) {
	case 0:
		v13 = F_siphash(m, l0, int32(base.Ui32(v7)>>(uint(int32(3))%32)), int32(_a_F_hashtableSdsHash_0))
		mBase = m.M
		return v13
	case 1:
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v19 = F_siphash(m, l0, v17, int32(_a_F_hashtableSdsHash_0))
		mBase = m.M
		return v19
	case 2:
		v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v25 = F_siphash(m, l0, v23, int32(_a_F_hashtableSdsHash_0))
		mBase = m.M
		return v25
	case 3:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v31 = F_siphash(m, l0, v29, int32(_a_F_hashtableSdsHash_0))
		mBase = m.M
		return v31
	case 4:
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v36 = v35
		v38 = F_siphash(m, l0, v36, int32(_a_F_hashtableSdsHash_0))
		mBase = m.M
		return v38
	default:
		v36 = int32(0)
		v38 = F_siphash(m, l0, v36, int32(_a_F_hashtableSdsHash_0))
		mBase = m.M
		return v38
	}
}
func F_hashtableSetHashFunctionSeed(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v8 int64
	_ = v8
	v2 = int32(0)
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, _c_F_hashtableSetHashFunctionSeed[0])) = v3
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(8))))
	*(*int64)(unsafe.Add(mBase, _c_F_hashtableSetHashFunctionSeed[1])) = v8
	return
}
func F_hashtableSize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	return v2 + v3
}
