package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_genericGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l4+l1<<(uint(int32(2))%32))))
	v18 = F_objectGetVal(m, v17)
	mBase = m.M
	v22 = v18
	for {
		v27 = v22 + int32(1)
		v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22))))
		v29 = F___isspace_1(m, v28)
		mBase = m.M
		if v29 != 0 {
			v22 = v27
			continue
		} else {
			break
		}
		break
	}
	v30 = int32(1)
	switch v28&int32(255) + int32(-43) {
	case 0:
		v36 = v30
		v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27))))
		v38 = v27
		v39 = v37
		v40 = v36
	default:
		v38 = v22
		v39 = v28
		v40 = v30
	case 2:
		v36 = int32(0)
		v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27))))
		v38 = v27
		v39 = v37
		v40 = v36
	}
	v43 = v39 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v43) {
		v61 = int32(0)
	} else {
		v47 = int32(0)
		v48 = v38
		v49 = v43
		for {
			v51 = int32(10)
			v53 = v47*v51 - v49
			v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48)+1)))
			v58 = v54 + int32(-48)
			if base.Ui32(v58) < base.Ui32(v51) {
				v47 = v53
				v48 = v48 + int32(1)
				v49 = v58
				continue
			} else {
				break
			}
			break
		}
		v61 = v53
	}
	if v40 != 0 {
		v67 = int32(0) - v61
	} else {
		v67 = v61
	}
	if v67 < int32(1) {
		v73 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l6))) = v73
		return v73
	} else {
		v71 = base.I32_div_s(l5-l2, l3)
		if v67 <= v71 {
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
			if v77 != 0 {
				v82 = v77
				v85 = v67 + base.B2i32(l0 != int32(0))
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
				if v85 <= v86 {
					v113 = v82
					*(*int32)(unsafe.Add(mBase, uint32(l6))) = v85
					v118 = v67 & int32(3)
					v119 = int32(0)
					if base.Ui32(v67) < base.Ui32(int32(4)) {
						v189 = v119
					} else {
						v125 = int32(0)
						v132 = v125
						v133 = v125
						for {
							v140 = int32(3)
							v142 = v113 + v133<<(uint(v140)%32)
							v143 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v143
							*(*int32)(unsafe.Add(mBase, uint32(v142))) = v133*l3 + l2
							v149 = v133 | int32(1)
							v152 = v113 + v149<<(uint(v140)%32)
							*(*int32)(unsafe.Add(mBase, uint32(v152)+4)) = v143
							*(*int32)(unsafe.Add(mBase, uint32(v152))) = v149*l3 + l2
							v159 = v133 | int32(2)
							v162 = v113 + v159<<(uint(v140)%32)
							*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v143
							*(*int32)(unsafe.Add(mBase, uint32(v162))) = v159*l3 + l2
							v169 = v133 | v140
							v172 = v113 + v169<<(uint(v140)%32)
							*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v143
							*(*int32)(unsafe.Add(mBase, uint32(v172))) = v169*l3 + l2
							v178 = int32(4)
							v179 = v133 + v178
							v181 = v132 + v178
							if v181 != v67&int32(2147483644) {
								v132 = v181
								v133 = v179
								continue
							} else {
								break
							}
							break
						}
						v189 = v179
					}
					if v118 == int32(0) {
					} else {
						v202 = v119
						v204 = v189
						for {
							v213 = v113 + v204<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v213)+4)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v213))) = v204*l3 + l2
							v219 = int32(1)
							v222 = v202 + v219
							if v222 != v118 {
								v202 = v222
								v204 = v204 + v219
								continue
							} else {
								break
							}
							break
						}
					}
					if l0 == int32(0) {
					} else {
						v241 = v113 + v67<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v241)+4)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v241))) = l0
					}
					return v85
				} else {
					v89 = v85 << (uint(int32(3)) % 32)
					v91 = l6 + int32(12)
					if v82 == v91 {
						v98 = F_valkey_malloc(m, v89)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = v98
							v101 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
							if v101 == int32(0) {
								v110 = v98
							} else {
								v105 = v101 << (uint(int32(3)) % 32)
								if v105 == int32(0) {
								} else {
									v108 = F__emscripten_memcpy_bulkmem(m, v98, v91, v105)
									mBase = m.M
								}
								v110 = v98
							}
							*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v85
							v113 = v110
							*(*int32)(unsafe.Add(mBase, uint32(l6))) = v85
							v118 = v67 & int32(3)
							v119 = int32(0)
							if base.Ui32(v67) < base.Ui32(int32(4)) {
								v189 = v119
							} else {
								v125 = int32(0)
								v132 = v125
								v133 = v125
								for {
									v140 = int32(3)
									v142 = v113 + v133<<(uint(v140)%32)
									v143 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v143
									*(*int32)(unsafe.Add(mBase, uint32(v142))) = v133*l3 + l2
									v149 = v133 | int32(1)
									v152 = v113 + v149<<(uint(v140)%32)
									*(*int32)(unsafe.Add(mBase, uint32(v152)+4)) = v143
									*(*int32)(unsafe.Add(mBase, uint32(v152))) = v149*l3 + l2
									v159 = v133 | int32(2)
									v162 = v113 + v159<<(uint(v140)%32)
									*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v143
									*(*int32)(unsafe.Add(mBase, uint32(v162))) = v159*l3 + l2
									v169 = v133 | v140
									v172 = v113 + v169<<(uint(v140)%32)
									*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v143
									*(*int32)(unsafe.Add(mBase, uint32(v172))) = v169*l3 + l2
									v178 = int32(4)
									v179 = v133 + v178
									v181 = v132 + v178
									if v181 != v67&int32(2147483644) {
										v132 = v181
										v133 = v179
										continue
									} else {
										break
									}
									break
								}
								v189 = v179
							}
							if v118 == int32(0) {
							} else {
								v202 = v119
								v204 = v189
								for {
									v213 = v113 + v204<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v213)+4)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v213))) = v204*l3 + l2
									v219 = int32(1)
									v222 = v202 + v219
									if v222 != v118 {
										v202 = v222
										v204 = v204 + v219
										continue
									} else {
										break
									}
									break
								}
							}
							if l0 == int32(0) {
							} else {
								v241 = v113 + v67<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v241)+4)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v241))) = l0
							}
							return v85
						}
					} else {
						v93 = F_valkey_realloc(m, v82, v89)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = v93
							v110 = v93
							*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v85
							v113 = v110
							*(*int32)(unsafe.Add(mBase, uint32(l6))) = v85
							v118 = v67 & int32(3)
							v119 = int32(0)
							if base.Ui32(v67) < base.Ui32(int32(4)) {
								v189 = v119
							} else {
								v125 = int32(0)
								v132 = v125
								v133 = v125
								for {
									v140 = int32(3)
									v142 = v113 + v133<<(uint(v140)%32)
									v143 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v143
									*(*int32)(unsafe.Add(mBase, uint32(v142))) = v133*l3 + l2
									v149 = v133 | int32(1)
									v152 = v113 + v149<<(uint(v140)%32)
									*(*int32)(unsafe.Add(mBase, uint32(v152)+4)) = v143
									*(*int32)(unsafe.Add(mBase, uint32(v152))) = v149*l3 + l2
									v159 = v133 | int32(2)
									v162 = v113 + v159<<(uint(v140)%32)
									*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v143
									*(*int32)(unsafe.Add(mBase, uint32(v162))) = v159*l3 + l2
									v169 = v133 | v140
									v172 = v113 + v169<<(uint(v140)%32)
									*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v143
									*(*int32)(unsafe.Add(mBase, uint32(v172))) = v169*l3 + l2
									v178 = int32(4)
									v179 = v133 + v178
									v181 = v132 + v178
									if v181 != v67&int32(2147483644) {
										v132 = v181
										v133 = v179
										continue
									} else {
										break
									}
									break
								}
								v189 = v179
							}
							if v118 == int32(0) {
							} else {
								v202 = v119
								v204 = v189
								for {
									v213 = v113 + v204<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v213)+4)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v213))) = v204*l3 + l2
									v219 = int32(1)
									v222 = v202 + v219
									if v222 != v118 {
										v202 = v222
										v204 = v204 + v219
										continue
									} else {
										break
									}
									break
								}
							}
							if l0 == int32(0) {
							} else {
								v241 = v113 + v67<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v241)+4)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v241))) = l0
							}
							return v85
						}
					}
				}
			} else {
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
				if v78 != 0 {
					F__serverAssert(m, int32(_a585), int32(_a550), int32(2292))
					mBase = m.M
					v251 = m.ExcPending
					if v251 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v80 = l6 + int32(12)
					*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = v80
					v82 = v80
					v85 = v67 + base.B2i32(l0 != int32(0))
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
					if v85 <= v86 {
						v113 = v82
						*(*int32)(unsafe.Add(mBase, uint32(l6))) = v85
						v118 = v67 & int32(3)
						v119 = int32(0)
						if base.Ui32(v67) < base.Ui32(int32(4)) {
							v189 = v119
						} else {
							v125 = int32(0)
							v132 = v125
							v133 = v125
							for {
								v140 = int32(3)
								v142 = v113 + v133<<(uint(v140)%32)
								v143 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v143
								*(*int32)(unsafe.Add(mBase, uint32(v142))) = v133*l3 + l2
								v149 = v133 | int32(1)
								v152 = v113 + v149<<(uint(v140)%32)
								*(*int32)(unsafe.Add(mBase, uint32(v152)+4)) = v143
								*(*int32)(unsafe.Add(mBase, uint32(v152))) = v149*l3 + l2
								v159 = v133 | int32(2)
								v162 = v113 + v159<<(uint(v140)%32)
								*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v143
								*(*int32)(unsafe.Add(mBase, uint32(v162))) = v159*l3 + l2
								v169 = v133 | v140
								v172 = v113 + v169<<(uint(v140)%32)
								*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v143
								*(*int32)(unsafe.Add(mBase, uint32(v172))) = v169*l3 + l2
								v178 = int32(4)
								v179 = v133 + v178
								v181 = v132 + v178
								if v181 != v67&int32(2147483644) {
									v132 = v181
									v133 = v179
									continue
								} else {
									break
								}
								break
							}
							v189 = v179
						}
						if v118 == int32(0) {
						} else {
							v202 = v119
							v204 = v189
							for {
								v213 = v113 + v204<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v213)+4)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v213))) = v204*l3 + l2
								v219 = int32(1)
								v222 = v202 + v219
								if v222 != v118 {
									v202 = v222
									v204 = v204 + v219
									continue
								} else {
									break
								}
								break
							}
						}
						if l0 == int32(0) {
						} else {
							v241 = v113 + v67<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v241)+4)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v241))) = l0
						}
						return v85
					} else {
						v89 = v85 << (uint(int32(3)) % 32)
						v91 = l6 + int32(12)
						if v82 == v91 {
							v98 = F_valkey_malloc(m, v89)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = v98
								v101 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
								if v101 == int32(0) {
									v110 = v98
								} else {
									v105 = v101 << (uint(int32(3)) % 32)
									if v105 == int32(0) {
									} else {
										v108 = F__emscripten_memcpy_bulkmem(m, v98, v91, v105)
										mBase = m.M
									}
									v110 = v98
								}
								*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v85
								v113 = v110
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v85
								v118 = v67 & int32(3)
								v119 = int32(0)
								if base.Ui32(v67) < base.Ui32(int32(4)) {
									v189 = v119
								} else {
									v125 = int32(0)
									v132 = v125
									v133 = v125
									for {
										v140 = int32(3)
										v142 = v113 + v133<<(uint(v140)%32)
										v143 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v143
										*(*int32)(unsafe.Add(mBase, uint32(v142))) = v133*l3 + l2
										v149 = v133 | int32(1)
										v152 = v113 + v149<<(uint(v140)%32)
										*(*int32)(unsafe.Add(mBase, uint32(v152)+4)) = v143
										*(*int32)(unsafe.Add(mBase, uint32(v152))) = v149*l3 + l2
										v159 = v133 | int32(2)
										v162 = v113 + v159<<(uint(v140)%32)
										*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v143
										*(*int32)(unsafe.Add(mBase, uint32(v162))) = v159*l3 + l2
										v169 = v133 | v140
										v172 = v113 + v169<<(uint(v140)%32)
										*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v143
										*(*int32)(unsafe.Add(mBase, uint32(v172))) = v169*l3 + l2
										v178 = int32(4)
										v179 = v133 + v178
										v181 = v132 + v178
										if v181 != v67&int32(2147483644) {
											v132 = v181
											v133 = v179
											continue
										} else {
											break
										}
										break
									}
									v189 = v179
								}
								if v118 == int32(0) {
								} else {
									v202 = v119
									v204 = v189
									for {
										v213 = v113 + v204<<(uint(int32(3))%32)
										*(*int32)(unsafe.Add(mBase, uint32(v213)+4)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v213))) = v204*l3 + l2
										v219 = int32(1)
										v222 = v202 + v219
										if v222 != v118 {
											v202 = v222
											v204 = v204 + v219
											continue
										} else {
											break
										}
										break
									}
								}
								if l0 == int32(0) {
								} else {
									v241 = v113 + v67<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v241)+4)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v241))) = l0
								}
								return v85
							}
						} else {
							v93 = F_valkey_realloc(m, v82, v89)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = v93
								v110 = v93
								*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v85
								v113 = v110
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v85
								v118 = v67 & int32(3)
								v119 = int32(0)
								if base.Ui32(v67) < base.Ui32(int32(4)) {
									v189 = v119
								} else {
									v125 = int32(0)
									v132 = v125
									v133 = v125
									for {
										v140 = int32(3)
										v142 = v113 + v133<<(uint(v140)%32)
										v143 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v143
										*(*int32)(unsafe.Add(mBase, uint32(v142))) = v133*l3 + l2
										v149 = v133 | int32(1)
										v152 = v113 + v149<<(uint(v140)%32)
										*(*int32)(unsafe.Add(mBase, uint32(v152)+4)) = v143
										*(*int32)(unsafe.Add(mBase, uint32(v152))) = v149*l3 + l2
										v159 = v133 | int32(2)
										v162 = v113 + v159<<(uint(v140)%32)
										*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v143
										*(*int32)(unsafe.Add(mBase, uint32(v162))) = v159*l3 + l2
										v169 = v133 | v140
										v172 = v113 + v169<<(uint(v140)%32)
										*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v143
										*(*int32)(unsafe.Add(mBase, uint32(v172))) = v169*l3 + l2
										v178 = int32(4)
										v179 = v133 + v178
										v181 = v132 + v178
										if v181 != v67&int32(2147483644) {
											v132 = v181
											v133 = v179
											continue
										} else {
											break
										}
										break
									}
									v189 = v179
								}
								if v118 == int32(0) {
								} else {
									v202 = v119
									v204 = v189
									for {
										v213 = v113 + v204<<(uint(int32(3))%32)
										*(*int32)(unsafe.Add(mBase, uint32(v213)+4)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v213))) = v204*l3 + l2
										v219 = int32(1)
										v222 = v202 + v219
										if v222 != v118 {
											v202 = v222
											v204 = v204 + v219
											continue
										} else {
											break
										}
										break
									}
								}
								if l0 == int32(0) {
								} else {
									v241 = v113 + v67<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v241)+4)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v241))) = l0
								}
								return v85
							}
						}
					}
				}
			}
		} else {
			v73 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l6))) = v73
			return v73
		}
	}
}
func F_genericZpopCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int64
	_ = v168
	var v173 int64
	_ = v173
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v194 float64
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 float64
	_ = v216
	var v222 int32
	_ = v222
	var v224 int64
	_ = v224
	var v228 int32
	_ = v228
	var v229 float64
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int64
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	if l8 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = int32(0)
	if v27 < l2 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = int32(0)
	goto L1
L3:
	;
	v31 = l2
	goto L5
L4:
	;
	v31 = v27
	goto L5
L5:
	;
	v42 = v27
	goto L13
L6:
	;
	F__serverAssertWithInfo(m, l0, v58, int32(_a1738), int32(_a1723), int32(3984))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L16
	} else {
		goto L107
	}
L7:
	;
	F__serverAssertWithInfo(m, l0, v58, int32(_a1739), int32(_a1723), int32(3977))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L16
	} else {
		goto L106
	}
L8:
	;
	F__serverAssertWithInfo(m, l0, v58, int32(_a169), int32(_a1723), int32(3965))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L16
	} else {
		goto L105
	}
L9:
	;
	F__serverAssertWithInfo(m, l0, v58, int32(_a170), int32(_a1723), int32(3956))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L16
	} else {
		goto L104
	}
L10:
	;
	m.G0 = v21 + int32(48)
	return
L11:
	;
	v80 = base.B2i32(l3 == int32(1))
	if l3 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L12:
	;
	if l7 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	if v42 == v31 {
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v63 = F_checkType(m, l0, v58, int32(3))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
	} else {
		goto L19
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1+v42<<(uint(int32(2))%32))))
	v58 = F_lookupKeyWrite(m, v55, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	if v58 == int32(0) {
		v42 = v42 + int32(1)
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	if v63 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	if l5 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[542]))
	F_addReply(m, l0, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	goto L10
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[542]))
	F_addReply(m, l0, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L16
	} else {
		goto L26
	}
L24:
	;
	F_addReplyNullArray(m, l0)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	goto L10
L26:
	;
	goto L10
L27:
	;
	v81 = int32(8)
	goto L29
L28:
	;
	v81 = int32(12)
	goto L29
L29:
	;
	if l3 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v84 = int32(-2)
	goto L32
L31:
	;
	v84 = int32(0)
	goto L32
L32:
	;
	if l5 == int32(-1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v92 = int32(1)
	goto L35
L34:
	;
	v92 = l5
	goto L35
L35:
	;
	v93 = F_zsetLength(m, v58)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L16
	} else {
		goto L36
	}
L36:
	;
	if v92 < v93 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v96 = v92
	goto L39
L38:
	;
	v96 = v93
	goto L39
L39:
	;
	v105 = v96
	v109 = int32(0)
	goto L40
L40:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	switch int32(base.Ui32(v116)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L45
	default:
		goto L44
	case 4:
		goto L46
	}
L41:
	;
	v284 = F_zsetLength(m, v58)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L16
	} else {
		goto L90
	}
L42:
	;
	v230 = F_zsetDel(m, v58, v228)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L16
	} else {
		goto L70
	}
L43:
	;
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	v228 = v140
	v229 = base.F64_convert_i64_s(v224)
	goto L42
L44:
	;
	F__serverPanic_1(m, int32(_a1723), int32(3981), int32(_a1054), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L16
	} else {
		goto L69
	}
L45:
	;
	v198 = F_objectGetVal(m, v58)
	mBase = m.M
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v199+v81)))
	if v201 == int32(0) {
		goto L7
	} else {
		goto L67
	}
L46:
	;
	v123 = F_objectGetVal(m, v58)
	mBase = m.M
	v124 = F_lpSeek(m, v123, v84)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L16
	} else {
		goto L47
	}
L47:
	;
	if v124 == int32(0) {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	v132 = F_lpGetValue(m, v124, v21+int32(28), v21+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L16
	} else {
		goto L51
	}
L49:
	;
	v141 = F_lpNext(m, v123, v124)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L16
	} else {
		goto L55
	}
L50:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v138 = F_sdsnewlen(m, v132, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L16
	} else {
		goto L54
	}
L51:
	;
	if v132 != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
	v135 = F_sdsfromlonglong(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	v140 = v135
	goto L49
L54:
	;
	v140 = v138
	goto L49
L55:
	;
	if v141 == int32(0) {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	v149 = F_lpGetValue(m, v141, v21+int32(44), v21+int32(32))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	if v149 == int32(0) {
		goto L43
	} else {
		goto L58
	}
L58:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v154 = int32(0)
	v158 = m.G0
	v160 = v158 - int32(32)
	m.G0 = v160
	v162 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v154
	v168 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	*(*int64)(unsafe.Add(mBase, uint32(v160+int32(8)))) = v168
	*(*int64)(unsafe.Add(mBase, uint32(v160)+24)) = int64(0)
	v173 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	*(*int64)(unsafe.Add(mBase, uint32(v160))) = v173
	F_ffc_from_chars_double_options(m, v160+int32(16), v149, v149+v153, v160+int32(24), v160)
	mBase = m.M
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v160)+20))
	if v181 == v154 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v228 = v140
	v229 = v194
	goto L42
L60:
	;
	goto L65
L61:
	;
	if v181 == int32(2) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v188 = int32(68)
	goto L64
L63:
	;
	v188 = int32(28)
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v188
	goto L60
L65:
	;
	v194 = *(*float64)(unsafe.Add(mBase, uint32(v160)+24))
	m.G0 = v160 + int32(32)
	goto L59
L67:
	;
	v205 = v201 + int32(16)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v209 = v205 + v206<<(uint(int32(3))%32)
	v210 = int32(*(*int8)(unsafe.Add(mBase, uint32(v209))))
	v214 = F_sdsdup(m, v209+v210+int32(1))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L16
	} else {
		goto L68
	}
L68:
	;
	v216 = *(*float64)(unsafe.Add(mBase, uint32(v201)))
	v228 = v214
	v229 = v216
	goto L42
L69:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	if v230 == int32(0) {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	v234 = int32(_a44)
	v236 = *(*int64)(unsafe.Add(mBase, _consts[83]))
	*(*int64)(unsafe.Add(mBase, _consts[83])) = v236 + int64(1)
	if v109 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if l6 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_consts[949])))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	F_notifyKeyspaceEvent(m, int32(128), v241, v57, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L16
	} else {
		goto L74
	}
L74:
	;
	F_addZpopInitialReply(m, l0, l4, l6, v105, v57)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228+int32(-1)))))
	switch v256 & int32(7) {
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
		v273 = int32(0)
		goto L79
	}
L77:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L16
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	F_addReplyBulkCBuffer(m, l0, v228, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L16
	} else {
		goto L85
	}
L80:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v228+int32(-17))))
	v273 = v272
	goto L79
L81:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v228+int32(-9))))
	v273 = v269
	goto L79
L82:
	;
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228+int32(-5)))))
	v273 = v266
	goto L79
L83:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228+int32(-3)))))
	v273 = v263
	goto L79
L84:
	;
	v273 = int32(base.Ui32(v256) >> (uint(int32(3)) % 32))
	goto L79
L85:
	;
	F_addReplyDouble(m, l0, v229)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L16
	} else {
		goto L86
	}
L86:
	;
	F_sdsfree(m, v228)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L16
	} else {
		goto L87
	}
L87:
	;
	v283 = v105 + int32(-1)
	if v283 != 0 {
		v105 = v283
		v109 = v109 + int32(1)
		goto L40
	} else {
		goto L88
	}
L88:
	;
	goto L41
L89:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v299, v57)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L16
	} else {
		goto L96
	}
L90:
	;
	if v284 != 0 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	if l8 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v291 = F_dbDelete(m, v290, v57)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L16
	} else {
		goto L94
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = int32(1)
	goto L92
L94:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a213), v57, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L16
	} else {
		goto L95
	}
L95:
	;
	goto L89
L96:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+48))
	if v303 != int32(352) {
		goto L10
	} else {
		goto L97
	}
L97:
	;
	v307 = F_createStringObjectFromLongLong(m, base.I64_extend_i32_s(v96))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L16
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v307
	if l3 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v316 = int32(280)
	goto L101
L100:
	;
	v316 = int32(276)
	goto L101
L101:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v316)+uint32(_consts[84])))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v318
	F_rewriteClientCommandVector(m, l0, int32(3), v21)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L16
	} else {
		goto L102
	}
L102:
	;
	F_decrRefCount(m, v307)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L16
	} else {
		goto L103
	}
L103:
	;
	goto L10
L104:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_genericZrangebyrankCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
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
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v93 float64
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int64
	_ = v135
	var v140 int64
	_ = v140
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v161 float64
	_ = v161
	var v165 int64
	_ = v165
	var v168 float64
	_ = v168
	var v169 int64
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var __phi235 int32
	_ = __phi235
	var v239 int32
	_ = v239
	var __phi239 int32
	_ = __phi239
	var v242 int32
	_ = v242
	var __phi242 int32
	_ = __phi242
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var __phi324 int32
	_ = __phi324
	var v328 int32
	_ = v328
	var __phi328 int32
	_ = __phi328
	var v331 int32
	_ = v331
	var __phi331 int32
	_ = __phi331
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 float64
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	v18 = m.G0
	v20 = v18 - int32(48)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = F_zsetLength(m, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F__serverPanic_1(m, int32(_a1723), int32(3186), int32(_a1054), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L8
	} else {
		goto L126
	}
L2:
	;
	F__serverAssertWithInfo(m, v22, l1, int32(_a106), int32(_a1723), int32(3180))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L8
	} else {
		goto L125
	}
L3:
	;
	F__serverAssertWithInfo(m, v22, l1, int32(_a1735), int32(_a1723), int32(3146))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L8
	} else {
		goto L124
	}
L4:
	;
	F__serverAssertWithInfo(m, v22, l1, int32(_a170), int32(_a1723), int32(3142))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L8
	} else {
		goto L123
	}
L5:
	;
	m.G0 = v20 + int32(48)
	return
L6:
	;
	if base.Ui32(v36) < base.Ui32(v25) {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	m.T0[v40].(func(*base.Module, int32, int32))(m, l0, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L15
	}
L8:
	;
	return
L9:
	;
	v28 = l2>>(uint(int32(31))%32)&v25 + l2
	v29 = int32(0)
	if v29 < v28 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v32 = v28
	goto L12
L11:
	;
	v32 = v29
	goto L12
L12:
	;
	v36 = l3>>(uint(int32(31))%32)&v25 + l3
	if v36 < v32 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	if v32 < v25 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	goto L7
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	m.T0[v44].(func(*base.Module, int32, int32))(m, l0, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L5
L17:
	;
	v50 = v36
	goto L19
L18:
	;
	v50 = v25 + int32(-1)
	goto L19
L19:
	;
	v51 = v50 - v32
	v53 = v51 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	m.T0[v54].(func(*base.Module, int32, int32))(m, l0, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch int32(base.Ui32(v57)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L22
	default:
		goto L1
	case 4:
		goto L23
	}
L21:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	m.T0[v469].(func(*base.Module, int32, int32))(m, l0, v53)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L8
	} else {
		goto L122
	}
L22:
	;
	v193 = F_objectGetVal(m, l1)
	mBase = m.M
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	if l5 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L23:
	;
	v64 = F_objectGetVal(m, l1)
	mBase = m.M
	v66 = v32 << (uint(int32(1)) % 32)
	if l5 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v69 = v66 ^ int32(-2)
	goto L26
L25:
	;
	v69 = v66
	goto L26
L26:
	;
	v70 = F_lpSeek(m, v64, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v70
	if v70 == int32(0) {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v75 = F_lpNext(m, v64, v70)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v75
	if v53 == int32(0) {
		goto L21
	} else {
		goto L30
	}
L30:
	;
	v84 = v51
	v93 = float64(0)
	goto L31
L31:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if v98 == int32(0) {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v101 == int32(0) {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v108 = F_lpGetValue(m, v98, v20+int32(20), v20+int32(8))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if l4 == int32(0) {
		v168 = v93
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v108 != 0 {
		goto L50
	} else {
		goto L51
	}
L37:
	;
	v116 = F_lpGetValue(m, v101, v20+int32(44), v20+int32(32))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L8
	} else {
		goto L39
	}
L38:
	;
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v20)+32))
	v168 = base.F64_convert_i64_s(v165)
	goto L36
L39:
	;
	if v116 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v121 = int32(0)
	v125 = m.G0
	v127 = v125 - int32(32)
	m.G0 = v127
	v129 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v121
	v135 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	*(*int64)(unsafe.Add(mBase, uint32(v127+int32(8)))) = v135
	*(*int64)(unsafe.Add(mBase, uint32(v127)+24)) = int64(0)
	v140 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = v140
	F_ffc_from_chars_double_options(m, v127+int32(16), v116, v116+v120, v127+int32(24), v127)
	mBase = m.M
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v127)+20))
	if v148 == v121 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v168 = v161
	goto L36
L42:
	;
	goto L47
L43:
	;
	if v148 == int32(2) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v155 = int32(68)
	goto L46
L45:
	;
	v155 = int32(28)
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v155
	goto L42
L47:
	;
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v127)+24))
	m.G0 = v127 + int32(32)
	goto L41
L49:
	;
	if l5 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	m.T0[v174].(func(*base.Module, int32, int32, int32, float64))(m, l0, v108, v173, v168)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L8
	} else {
		goto L53
	}
L51:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	m.T0[v170].(func(*base.Module, int32, int64, float64))(m, l0, v169, v168)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	goto L49
L54:
	;
	if v84 != 0 {
		v84 = v84 + int32(-1)
		v93 = v168
		goto L31
	} else {
		goto L59
	}
L55:
	;
	F_zzlNext(m, v64, v20+int32(28), v20+int32(24))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L8
	} else {
		goto L58
	}
L56:
	;
	F_zzlPrev(m, v64, v20+int32(28), v20+int32(24))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	goto L54
L59:
	;
	goto L21
L60:
	;
	if v53 == int32(0) {
		goto L21
	} else {
		goto L107
	}
L61:
	;
	if int32(1) <= v28 {
		goto L85
	} else {
		goto L86
	}
L62:
	;
	if int32(1) <= v28 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v200 = int32(0)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v194)+16))
	if v201 <= v200 {
		v376 = v200
		goto L60
	} else {
		goto L65
	}
L64:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v194)+8))
	v376 = v199
	goto L60
L65:
	;
	v204 = v25 - v32
	v210 = v194
	v217 = int32(0)
	v219 = v201
	goto L67
L66:
	;
	if v242 == v204 {
		goto L82
	} else {
		goto L83
	}
L67:
	;
	v224 = v219 + int32(-1)
	v226 = v224 << (uint(int32(3)) % 32)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v210+v226)+12))
	if v228 == int32(0) {
		v266 = v210
		v273 = v217
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v376 = v266
	goto L60
L69:
	;
	if v273 == v204 {
		goto L79
	} else {
		goto L80
	}
L70:
	;
	__phi235 = v210
	__phi239 = v228
	__phi242 = v217
	v235 = __phi235
	v239 = __phi239
	v242 = __phi242
	goto L71
L71:
	;
	if v224 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v266 = v239
	v273 = v259
	goto L69
L73:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v239+v226)+12))
	if v261 != 0 {
		__phi235 = v239
		__phi239 = v261
		__phi242 = v259
		v235 = __phi235
		v239 = __phi239
		v242 = __phi242
		goto L71
	} else {
		goto L78
	}
L74:
	;
	v257 = v242 + int32(1)
	if base.Ui32(v204) < base.Ui32(v257) {
		goto L66
	} else {
		goto L77
	}
L75:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v235+v226+int32(16))))
	v254 = v253 + v242
	if base.Ui32(v254) <= base.Ui32(v204) {
		v259 = v254
		goto L73
	} else {
		goto L76
	}
L76:
	;
	v266 = v235
	v273 = v242
	goto L69
L77:
	;
	v259 = v257
	goto L73
L78:
	;
	goto L72
L79:
	;
	goto L68
L80:
	;
	if v219 < int32(2) {
		v376 = v200
		goto L60
	} else {
		goto L81
	}
L81:
	;
	v210 = v266
	v217 = v273
	v219 = v224
	goto L67
L82:
	;
	v284 = v235
	goto L84
L83:
	;
	v284 = int32(0)
	goto L84
L84:
	;
	v376 = v284
	goto L60
L85:
	;
	v288 = int32(0)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v194)+16))
	if v289 <= v288 {
		v376 = v288
		goto L60
	} else {
		goto L87
	}
L86:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v194)+12))
	v376 = v287
	goto L60
L87:
	;
	v293 = v32 + int32(1)
	v299 = v194
	v306 = int32(0)
	v308 = v289
	goto L89
L88:
	;
	if v331 == v293 {
		goto L104
	} else {
		goto L105
	}
L89:
	;
	v313 = v308 + int32(-1)
	v315 = v313 << (uint(int32(3)) % 32)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v299+v315)+12))
	if v317 == int32(0) {
		v355 = v299
		v362 = v306
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v376 = v355
	goto L60
L91:
	;
	if v362 == v293 {
		goto L101
	} else {
		goto L102
	}
L92:
	;
	__phi324 = v299
	__phi328 = v317
	__phi331 = v306
	v324 = __phi324
	v328 = __phi328
	v331 = __phi331
	goto L93
L93:
	;
	if v313 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v355 = v328
	v362 = v348
	goto L91
L95:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v328+v315)+12))
	if v350 != 0 {
		__phi324 = v328
		__phi328 = v350
		__phi331 = v348
		v324 = __phi324
		v328 = __phi328
		v331 = __phi331
		goto L93
	} else {
		goto L100
	}
L96:
	;
	v346 = v331 + int32(1)
	if base.Ui32(v293) < base.Ui32(v346) {
		goto L88
	} else {
		goto L99
	}
L97:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v324+v315+int32(16))))
	v343 = v342 + v331
	if base.Ui32(v343) <= base.Ui32(v293) {
		v348 = v343
		goto L95
	} else {
		goto L98
	}
L98:
	;
	v355 = v324
	v362 = v331
	goto L91
L99:
	;
	v348 = v346
	goto L95
L100:
	;
	goto L94
L101:
	;
	goto L90
L102:
	;
	if v308 < int32(2) {
		v376 = v288
		goto L60
	} else {
		goto L103
	}
L103:
	;
	v299 = v355
	v306 = v362
	v308 = v313
	goto L89
L104:
	;
	v373 = v324
	goto L106
L105:
	;
	v373 = int32(0)
	goto L106
L106:
	;
	v376 = v373
	goto L60
L107:
	;
	if l5 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v395 = int32(8)
	goto L110
L109:
	;
	v395 = int32(12)
	goto L110
L110:
	;
	v398 = v376
	v399 = v51
	goto L111
L111:
	;
	if v398 == int32(0) {
		goto L2
	} else {
		goto L113
	}
L112:
	;
	goto L21
L113:
	;
	v416 = v398 + int32(16)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	v420 = v416 + v417<<(uint(int32(3))%32)
	v421 = int32(*(*int8)(unsafe.Add(mBase, uint32(v420))))
	v422 = v420 + v421
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422))))
	switch v426 & int32(7) {
	case 0:
		goto L119
	case 1:
		goto L118
	case 2:
		goto L117
	case 3:
		goto L116
	case 4:
		goto L115
	default:
		v443 = int32(0)
		goto L114
	}
L114:
	;
	v444 = *(*float64)(unsafe.Add(mBase, uint32(v398)))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	m.T0[v445].(func(*base.Module, int32, int32, int32, float64))(m, l0, v422+int32(1), v443, v444)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L8
	} else {
		goto L120
	}
L115:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v422+int32(-16))))
	v443 = v442
	goto L114
L116:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v422+int32(-8))))
	v443 = v439
	goto L114
L117:
	;
	v436 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v422+int32(-4)))))
	v443 = v436
	goto L114
L118:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422+int32(-2)))))
	v443 = v433
	goto L114
L119:
	;
	v443 = int32(base.Ui32(v426) >> (uint(int32(3)) % 32))
	goto L114
L120:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v398+v395)))
	if v399 != 0 {
		v398 = v451
		v399 = v399 + int32(-1)
		goto L111
	} else {
		goto L121
	}
L121:
	;
	goto L112
L122:
	;
	goto L5
L123:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_genericZrangebyscoreCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int64
	_ = v148
	var v153 int64
	_ = v153
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v174 float64
	_ = v174
	var v178 int64
	_ = v178
	var v180 float64
	_ = v180
	var v183 float64
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 float64
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int64
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v256 float64
	_ = v256
	var v257 float64
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 float64
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 float64
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var __phi296 int32
	_ = __phi296
	var v308 int32
	_ = v308
	var __phi308 int32
	_ = __phi308
	var v311 int32
	_ = v311
	var __phi311 int32
	_ = __phi311
	var v315 float64
	_ = v315
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v394 int32
	_ = v394
	var __phi394 int32
	_ = __phi394
	var v403 int32
	_ = v403
	var __phi403 int32
	_ = __phi403
	var v405 int32
	_ = v405
	var __phi405 int32
	_ = __phi405
	var v407 float64
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v464 int32
	_ = v464
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var __phi542 int32
	_ = __phi542
	var v548 int32
	_ = v548
	var __phi548 int32
	_ = __phi548
	var v556 int32
	_ = v556
	var __phi556 int32
	_ = __phi556
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v590 int32
	_ = v590
	var v602 int32
	_ = v602
	var v621 float64
	_ = v621
	var v624 int32
	_ = v624
	var v644 int32
	_ = v644
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v687 int32
	_ = v687
	var __phi687 int32
	_ = __phi687
	var v696 int32
	_ = v696
	var __phi696 int32
	_ = __phi696
	var v698 int32
	_ = v698
	var __phi698 int32
	_ = __phi698
	var v700 float64
	_ = v700
	var v703 int32
	_ = v703
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v731 int32
	_ = v731
	var v740 int32
	_ = v740
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var __phi908 int32
	_ = __phi908
	var v914 int32
	_ = v914
	var __phi914 int32
	_ = __phi914
	var v923 int32
	_ = v923
	var __phi923 int32
	_ = __phi923
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v957 int32
	_ = v957
	var v968 int32
	_ = v968
	var v987 float64
	_ = v987
	var v990 int32
	_ = v990
	var v1010 int32
	_ = v1010
	var v1060 int32
	_ = v1060
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1090 int32
	_ = v1090
	var v1097 float64
	_ = v1097
	var v1100 float64
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1107 float64
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1161 int32
	_ = v1161
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1193 int32
	_ = v1193
	var v1200 int32
	_ = v1200
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	m.T0[v21].(func(*base.Module, int32, int32))(m, l0, int32(-1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l3 < int32(1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F__serverPanic_1(m, int32(_a1723), int32(3316), int32(_a1054), int32(0))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L1
	} else {
		goto L246
	}
L4:
	;
	F__serverAssert(m, int32(_a169), int32(_a1723), int32(857))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L1
	} else {
		goto L245
	}
L5:
	;
	m.G0 = v18 + int32(48)
	return
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch int32(base.Ui32(v33)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L12
	default:
		goto L3
	case 4:
		goto L13
	}
L7:
	;
	v26 = F_zsetLength(m, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if l3 < v26 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	m.T0[v30].(func(*base.Module, int32, int32))(m, l0, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L5
L11:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	m.T0[v1168].(func(*base.Module, int32, int32))(m, l0, v1161)
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L1
	} else {
		goto L244
	}
L12:
	;
	v231 = int32(0)
	v232 = F_objectGetVal(m, l2)
	mBase = m.M
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	if l5 != 0 {
		goto L72
	} else {
		goto L73
	}
L13:
	;
	v40 = F_objectGetVal(m, l2)
	mBase = m.M
	if l5 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v47
	v49 = int32(0)
	if v47 == v49 {
		v1161 = v49
		goto L11
	} else {
		goto L19
	}
L15:
	;
	v45 = F_zzlFirstInRange(m, v40, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v43 = F_zzlLastInRange(m, v40, l1)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v47 = v43
	goto L14
L18:
	;
	v47 = v45
	goto L14
L19:
	;
	v52 = F_lpNext(m, v40, v47)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v52
	if l3 == int32(0) {
		v93 = v47
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v109 = v93
	v110 = l4
	v115 = int32(0)
	goto L32
L22:
	;
	v60 = l3
	goto L23
L23:
	;
	if l5 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v93 = v86
	goto L21
L25:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v86 == int32(0) {
		v1161 = v49
		goto L11
	} else {
		goto L30
	}
L26:
	;
	F_zzlNext(m, v40, v18+int32(28), v18+int32(24))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	F_zzlPrev(m, v40, v18+int32(28), v18+int32(24))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	goto L25
L30:
	;
	v90 = v60 + int32(-1)
	if v90 != 0 {
		v60 = v90
		goto L23
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	if v110 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	if v122 == int32(0) {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	v1161 = l4
	goto L11
L36:
	;
	v129 = F_lpGetValue(m, v122, v18+int32(44), v18+int32(32))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	if l5 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	v180 = base.F64_convert_i64_s(v178)
	goto L37
L39:
	;
	if v129 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v134 = int32(0)
	v138 = m.G0
	v140 = v138 - int32(32)
	m.G0 = v140
	v142 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v134
	v148 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	*(*int64)(unsafe.Add(mBase, uint32(v140+int32(8)))) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v140)+24)) = int64(0)
	v153 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	*(*int64)(unsafe.Add(mBase, uint32(v140))) = v153
	F_ffc_from_chars_double_options(m, v140+int32(16), v129, v129+v133, v140+int32(24), v140)
	mBase = m.M
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v140)+20))
	if v161 == v134 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v180 = v174
	goto L37
L42:
	;
	goto L47
L43:
	;
	if v161 == int32(2) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v168 = int32(68)
	goto L46
L45:
	;
	v168 = int32(28)
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v168
	goto L42
L47:
	;
	v174 = *(*float64)(unsafe.Add(mBase, uint32(v140)+24))
	m.G0 = v140 + int32(32)
	goto L41
L49:
	;
	v202 = F_lpGetValue(m, v109, v18+int32(20), v18+int32(8))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L62
	}
L50:
	;
	v190 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v193 != 0 {
		goto L56
	} else {
		goto L57
	}
L51:
	;
	v183 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v186 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v187 = base.F64_gt(v180, v183)
	goto L54
L53:
	;
	v187 = base.F64_ge(v180, v183)
	goto L54
L54:
	;
	if v187 == int32(1) {
		goto L49
	} else {
		goto L55
	}
L55:
	;
	v1161 = v115
	goto L11
L56:
	;
	v194 = base.F64_lt(v180, v190)
	goto L58
L57:
	;
	v194 = base.F64_le(v180, v190)
	goto L58
L58:
	;
	if v194 != int32(1) {
		v1161 = v115
		goto L11
	} else {
		goto L59
	}
L59:
	;
	goto L49
L60:
	;
	if l5 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L61:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	m.T0[v209].(func(*base.Module, int32, int32, int32, float64))(m, l0, v202, v208, v180)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L65
	}
L62:
	;
	if v202 != 0 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	m.T0[v205].(func(*base.Module, int32, int64, float64))(m, l0, v204, v180)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	goto L60
L65:
	;
	goto L60
L66:
	;
	v229 = v115 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v230 != 0 {
		v109 = v230
		v110 = v110 + int32(-1)
		v115 = v229
		goto L32
	} else {
		goto L71
	}
L67:
	;
	F_zzlNext(m, v40, v18+int32(28), v18+int32(24))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	F_zzlPrev(m, v40, v18+int32(28), v18+int32(24))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	goto L66
L71:
	;
	v1161 = v229
	goto L11
L72:
	;
	v236 = int32(-1)
	goto L74
L73:
	;
	v236 = v231
	goto L74
L74:
	;
	v237 = l3 ^ v236
	v238 = int32(0)
	v256 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v257 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_gt(v256, v257) != 0 {
		v1060 = v238
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if v1060 == int32(0) {
		v1161 = v231
		goto L11
	} else {
		goto L217
	}
L76:
	;
	goto L75
L77:
	;
	if base.F64_ne(v256, v257) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v233)+8))
	if v262 == int32(0) {
		v1060 = v238
		goto L76
	} else {
		goto L82
	}
L79:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v260 != 0 {
		v1060 = v238
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v261 != 0 {
		v1060 = v238
		goto L76
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	v265 = *(*float64)(unsafe.Add(mBase, uint32(v262)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v268 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v269 = base.F64_gt(v265, v256)
	goto L85
L84:
	;
	v269 = base.F64_ge(v265, v256)
	goto L85
L85:
	;
	if v269 != int32(1) {
		v1060 = v238
		goto L76
	} else {
		goto L86
	}
L86:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	if v272 == int32(0) {
		v1060 = v238
		goto L76
	} else {
		goto L87
	}
L87:
	;
	v275 = *(*float64)(unsafe.Add(mBase, uint32(v272)))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v278 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v279 = base.F64_lt(v275, v257)
	goto L90
L89:
	;
	v279 = base.F64_le(v275, v257)
	goto L90
L90:
	;
	if v279 != int32(1) {
		v1060 = v238
		goto L76
	} else {
		goto L91
	}
L91:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v233)+16))
	v286 = v284 + int32(-1)
	v288 = v286 << (uint(int32(3)) % 32)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v233+int32(12)+v288)))
	if v290 != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	if int32(-1) < v237 {
		goto L107
	} else {
		goto L108
	}
L93:
	;
	__phi296 = v290
	__phi308 = int32(0)
	__phi311 = v233
	v296 = __phi296
	v308 = __phi308
	v311 = __phi311
	goto L95
L94:
	;
	v345 = int32(0)
	v346 = v233
	goto L92
L95:
	;
	v315 = *(*float64)(unsafe.Add(mBase, uint32(v296)))
	if v268 != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v345 = v327
	v346 = v296
	goto L92
L97:
	;
	if v284 < int32(2) {
		v326 = int32(1)
		goto L102
	} else {
		goto L103
	}
L98:
	;
	v318 = base.F64_gt(v315, v256)
	goto L100
L99:
	;
	v318 = base.F64_ge(v315, v256)
	goto L100
L100:
	;
	if v318 == int32(0) {
		goto L97
	} else {
		goto L101
	}
L101:
	;
	v345 = v308
	v346 = v311
	goto L92
L102:
	;
	v327 = v326 + v308
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v296+v288+int32(12))))
	if v331 != 0 {
		__phi296 = v331
		__phi308 = v327
		__phi311 = v296
		v296 = __phi296
		v308 = __phi308
		v311 = __phi311
		goto L95
	} else {
		goto L104
	}
L103:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v311+v288+int32(16))))
	v326 = v325
	goto L102
L104:
	;
	goto L96
L105:
	;
	v1060 = v644
	goto L76
L107:
	;
	if v284 < int32(2) {
		v740 = v346
		v755 = v345
		goto L161
	} else {
		goto L162
	}
L108:
	;
	v354 = int32(0)
	if v284 <= v354 {
		v449 = v346
		v464 = v345
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if v464 < int32(0)-v237 {
		v1060 = v354
		goto L76
	} else {
		goto L126
	}
L110:
	;
	v358 = v346
	v373 = v345
	v374 = v284
	goto L111
L111:
	;
	v378 = v374 + int32(-1)
	v380 = v378 << (uint(int32(3)) % 32)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v358+v380+int32(12))))
	if v384 == int32(0) {
		v427 = v358
		v442 = v373
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v449 = v427
	v464 = v442
	goto L109
L113:
	;
	if int32(1) < v374 {
		v358 = v427
		v373 = v442
		v374 = v378
		goto L111
	} else {
		goto L125
	}
L114:
	;
	__phi394 = v384
	__phi403 = v373
	__phi405 = v358
	v394 = __phi394
	v403 = __phi403
	v405 = __phi405
	goto L115
L115:
	;
	v407 = *(*float64)(unsafe.Add(mBase, uint32(v394)))
	if v278 != 0 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v427 = v394
	v442 = v421
	goto L113
L117:
	;
	v413 = int32(1)
	if v374 == v413 {
		v420 = v413
		goto L122
	} else {
		goto L123
	}
L118:
	;
	v410 = base.F64_lt(v407, v257)
	goto L120
L119:
	;
	v410 = base.F64_le(v407, v257)
	goto L120
L120:
	;
	if v410 == int32(1) {
		goto L117
	} else {
		goto L121
	}
L121:
	;
	v427 = v405
	v442 = v403
	goto L113
L122:
	;
	v421 = v420 + v403
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v394+v380+int32(12))))
	if v425 != 0 {
		__phi394 = v425
		__phi403 = v421
		__phi405 = v394
		v394 = __phi394
		v403 = __phi403
		v405 = __phi405
		goto L115
	} else {
		goto L124
	}
L123:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v405+v380+int32(16))))
	v420 = v419
	goto L122
L124:
	;
	goto L116
L125:
	;
	goto L112
L126:
	;
	if v237 < int32(-10) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L105
L128:
	;
	v621 = *(*float64)(unsafe.Add(mBase, uint32(v602)))
	if v268 != 0 {
		goto L156
	} else {
		goto L157
	}
L129:
	;
	v507 = int32(0)
	if v284 <= v507 {
		v644 = v507
		goto L127
	} else {
		goto L139
	}
L130:
	;
	if v237 == int32(-1) {
		v602 = v449
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v475 = int32(-2)
	if v237 < v475 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v478 = v237
	goto L134
L133:
	;
	v478 = v475
	goto L134
L134:
	;
	v481 = v449
	v487 = int32(0)
	goto L135
L135:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v481)+8))
	v502 = v487 + int32(1)
	if v478^v502 != int32(-1) {
		v481 = v500
		v487 = v502
		goto L135
	} else {
		goto L137
	}
L136:
	;
	if v500 != 0 {
		v602 = v500
		goto L128
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	v644 = int32(0)
	goto L127
L139:
	;
	v513 = v237 - v345 + v464 + int32(1)
	v516 = v346
	v527 = v286
	v530 = int32(0)
	goto L141
L140:
	;
	if v556 != v513 {
		v644 = v507
		goto L127
	} else {
		goto L155
	}
L141:
	;
	v536 = v527 << (uint(int32(3)) % 32)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v516+v536)+12))
	if v538 == int32(0) {
		v576 = v516
		v590 = v530
		goto L143
	} else {
		goto L144
	}
L143:
	;
	if v590 == v513 {
		v602 = v576
		goto L128
	} else {
		goto L153
	}
L144:
	;
	__phi542 = v516
	__phi548 = v538
	__phi556 = v530
	v542 = __phi542
	v548 = __phi548
	v556 = __phi556
	goto L145
L145:
	;
	if v527 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v576 = v548
	v590 = v572
	goto L143
L147:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v548+v536)+12))
	if v574 != 0 {
		__phi542 = v548
		__phi548 = v574
		__phi556 = v572
		v542 = __phi542
		v548 = __phi548
		v556 = __phi556
		goto L145
	} else {
		goto L152
	}
L148:
	;
	v570 = v556 + int32(1)
	if base.Ui32(v513) < base.Ui32(v570) {
		goto L140
	} else {
		goto L151
	}
L149:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v542+v536+int32(16))))
	v567 = v566 + v556
	if base.Ui32(v567) <= base.Ui32(v513) {
		v572 = v567
		goto L147
	} else {
		goto L150
	}
L150:
	;
	v576 = v542
	v590 = v556
	goto L143
L151:
	;
	v572 = v570
	goto L147
L152:
	;
	goto L146
L153:
	;
	if v527 < int32(1) {
		v644 = v507
		goto L127
	} else {
		goto L154
	}
L154:
	;
	v516 = v576
	v527 = v527 + int32(-1)
	v530 = v590
	goto L141
L155:
	;
	v602 = v542
	goto L128
L156:
	;
	v624 = base.F64_gt(v621, v256)
	goto L158
L157:
	;
	v624 = base.F64_ge(v621, v256)
	goto L158
L158:
	;
	if v624 != int32(1) {
		v1060 = v354
		goto L76
	} else {
		goto L159
	}
L159:
	;
	v644 = v602
	goto L127
L161:
	;
	v759 = int32(0)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	if base.Ui32(v761) <= base.Ui32(v755+v237) {
		v1060 = v759
		goto L76
	} else {
		goto L179
	}
L162:
	;
	v655 = v346
	v661 = v284 + int32(-2)
	v670 = v345
	goto L163
L163:
	;
	v675 = v661 << (uint(int32(3)) % 32)
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v655+v675)+12))
	if v677 == int32(0) {
		v716 = v655
		v731 = v670
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v740 = v716
	v755 = v731
	goto L161
L165:
	;
	if int32(0) < v661 {
		v655 = v716
		v661 = v661 + int32(-1)
		v670 = v731
		goto L163
	} else {
		goto L178
	}
L166:
	;
	__phi687 = v677
	__phi696 = v670
	__phi698 = v655
	v687 = __phi687
	v696 = __phi696
	v698 = __phi698
	goto L167
L167:
	;
	v700 = *(*float64)(unsafe.Add(mBase, uint32(v687)))
	if v268 != 0 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v716 = v687
	v731 = v712
	goto L165
L169:
	;
	if v661 != 0 {
		goto L175
	} else {
		goto L176
	}
L170:
	;
	v703 = base.F64_gt(v700, v256)
	goto L172
L171:
	;
	v703 = base.F64_ge(v700, v256)
	goto L172
L172:
	;
	if v703 == int32(0) {
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v716 = v698
	v731 = v696
	goto L165
L174:
	;
	v712 = v711 + v696
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v687+v675)+12))
	if v714 != 0 {
		__phi687 = v714
		__phi696 = v712
		__phi698 = v687
		v687 = __phi687
		v696 = __phi696
		v698 = __phi698
		goto L167
	} else {
		goto L177
	}
L175:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v698+v675+int32(16))))
	v711 = v710
	goto L174
L176:
	;
	v711 = int32(1)
	goto L174
L177:
	;
	goto L168
L178:
	;
	goto L164
L179:
	;
	if int32(9) < v237 {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	v1060 = v1010
	goto L76
L181:
	;
	v987 = *(*float64)(unsafe.Add(mBase, uint32(v968)))
	if v278 != 0 {
		goto L212
	} else {
		goto L213
	}
L182:
	;
	v873 = int32(0)
	if v284 <= v873 {
		v1010 = v873
		goto L180
	} else {
		goto L195
	}
L183:
	;
	v766 = v237 + int32(1)
	v767 = int32(7)
	v768 = v766 & v767
	if base.Ui32(v237) < base.Ui32(v767) {
		v806 = v740
		goto L184
	} else {
		goto L185
	}
L184:
	;
	if v768 == int32(0) {
		v853 = v806
		goto L189
	} else {
		goto L190
	}
L185:
	;
	v775 = v740
	v781 = int32(0)
	goto L186
L186:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v775)+12))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v794)+12))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v795)+12))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v796)+12))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v797)+12))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v798)+12))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v799)+12))
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v800)+12))
	v803 = v781 + int32(8)
	if v803 != v766&int32(-8) {
		v775 = v801
		v781 = v803
		goto L186
	} else {
		goto L188
	}
L187:
	;
	v806 = v801
	goto L184
L188:
	;
	goto L187
L189:
	;
	if v853 != 0 {
		v968 = v853
		goto L181
	} else {
		goto L194
	}
L190:
	;
	v829 = v806
	v835 = int32(0)
	goto L191
L191:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v829)+12))
	v850 = v835 + int32(1)
	if v850 != v768 {
		v829 = v848
		v835 = v850
		goto L191
	} else {
		goto L193
	}
L192:
	;
	v853 = v848
	goto L189
L193:
	;
	goto L192
L194:
	;
	v1010 = int32(0)
	goto L180
L195:
	;
	v879 = v237 - v345 + v755 + int32(1)
	v882 = v346
	v893 = v286
	v897 = int32(0)
	goto L197
L196:
	;
	if v923 != v879 {
		v1010 = v873
		goto L180
	} else {
		goto L211
	}
L197:
	;
	v902 = v893 << (uint(int32(3)) % 32)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v882+v902)+12))
	if v904 == int32(0) {
		v942 = v882
		v957 = v897
		goto L199
	} else {
		goto L200
	}
L199:
	;
	if v957 == v879 {
		v968 = v942
		goto L181
	} else {
		goto L209
	}
L200:
	;
	__phi908 = v882
	__phi914 = v904
	__phi923 = v897
	v908 = __phi908
	v914 = __phi914
	v923 = __phi923
	goto L201
L201:
	;
	if v893 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	v942 = v914
	v957 = v938
	goto L199
L203:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v914+v902)+12))
	if v940 != 0 {
		__phi908 = v914
		__phi914 = v940
		__phi923 = v938
		v908 = __phi908
		v914 = __phi914
		v923 = __phi923
		goto L201
	} else {
		goto L208
	}
L204:
	;
	v936 = v923 + int32(1)
	if base.Ui32(v879) < base.Ui32(v936) {
		goto L196
	} else {
		goto L207
	}
L205:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v908+v902+int32(16))))
	v933 = v932 + v923
	if base.Ui32(v933) <= base.Ui32(v879) {
		v938 = v933
		goto L203
	} else {
		goto L206
	}
L206:
	;
	v942 = v908
	v957 = v923
	goto L199
L207:
	;
	v938 = v936
	goto L203
L208:
	;
	goto L202
L209:
	;
	if v893 < int32(1) {
		v1010 = v873
		goto L180
	} else {
		goto L210
	}
L210:
	;
	v882 = v942
	v893 = v893 + int32(-1)
	v897 = v957
	goto L197
L211:
	;
	v968 = v908
	goto L181
L212:
	;
	v990 = base.F64_lt(v987, v257)
	goto L214
L213:
	;
	v990 = base.F64_le(v987, v257)
	goto L214
L214:
	;
	if v990 != int32(1) {
		v1060 = v759
		goto L76
	} else {
		goto L215
	}
L215:
	;
	v1010 = v968
	goto L180
L217:
	;
	if l5 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1080 = int32(8)
	goto L220
L219:
	;
	v1080 = int32(12)
	goto L220
L220:
	;
	v1084 = l4
	v1085 = v1060
	v1090 = int32(0)
	goto L221
L221:
	;
	if v1084 != 0 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	v1161 = v1147
	goto L11
L223:
	;
	v1097 = *(*float64)(unsafe.Add(mBase, uint32(v1085)))
	if l5 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	v1161 = l4
	goto L11
L225:
	;
	v1116 = v1085 + int32(16)
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1116)))
	v1120 = v1116 + v1117<<(uint(int32(3))%32)
	v1121 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1120))))
	v1122 = v1120 + v1121
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1122))))
	switch v1126 & int32(7) {
	case 0:
		goto L241
	case 1:
		goto L240
	case 2:
		goto L239
	case 3:
		goto L238
	case 4:
		goto L237
	default:
		v1143 = int32(0)
		goto L236
	}
L226:
	;
	v1107 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v1110 != 0 {
		goto L232
	} else {
		goto L233
	}
L227:
	;
	v1100 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1103 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1104 = base.F64_gt(v1097, v1100)
	goto L230
L229:
	;
	v1104 = base.F64_ge(v1097, v1100)
	goto L230
L230:
	;
	if v1104 == int32(1) {
		goto L225
	} else {
		goto L231
	}
L231:
	;
	v1161 = v1090
	goto L11
L232:
	;
	v1111 = base.F64_lt(v1097, v1107)
	goto L234
L233:
	;
	v1111 = base.F64_le(v1097, v1107)
	goto L234
L234:
	;
	if v1111 != int32(1) {
		v1161 = v1090
		goto L11
	} else {
		goto L235
	}
L235:
	;
	goto L225
L236:
	;
	v1147 = v1090 + int32(1)
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	m.T0[v1148].(func(*base.Module, int32, int32, int32, float64))(m, l0, v1122+int32(1), v1143, v1097)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L1
	} else {
		goto L242
	}
L237:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1122+int32(-16))))
	v1143 = v1142
	goto L236
L238:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1122+int32(-8))))
	v1143 = v1139
	goto L236
L239:
	;
	v1136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1122+int32(-4)))))
	v1143 = v1136
	goto L236
L240:
	;
	v1133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1122+int32(-2)))))
	v1143 = v1133
	goto L236
L241:
	;
	v1143 = int32(base.Ui32(v1126) >> (uint(int32(3)) % 32))
	goto L236
L242:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1085+v1080)))
	if v1152 != 0 {
		v1084 = v1084 + int32(-1)
		v1085 = v1152
		v1090 = v1147
		goto L221
	} else {
		goto L243
	}
L243:
	;
	goto L222
L244:
	;
	goto L5
L245:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L246:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_popGenericCommand(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
	if v11 < int32(4) {
		if v11 == int32(3) {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
			v28 = F_getPositiveLongFromObjectOrReply(m, l0, v24, v9+int32(12), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				if v28 != 0 {
					m.G0 = v9 + int32(16)
					return
				} else {
					v33 = int32(_a1677)
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v33+v36<<(uint(int32(2))%32))))
					v41 = F_lookupKeyWriteOrReply(m, l0, v35, v40)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						if v41 == int32(0) {
							m.G0 = v9 + int32(16)
							return
						} else {
							v46 = F_checkType(m, l0, v41, int32(1))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								if v46 != 0 {
									m.G0 = v9 + int32(16)
									return
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
									if v11 != int32(3) {
										if v48 != 0 {
											v69 = F_listTypeLength(m, v41)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return
											} else {
												v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												F_initDeferredReplyBuffer(m, l0)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return
												} else {
													if v71 < v69 {
														v76 = v71
													} else {
														v76 = v69
													}
													if l1 != 0 {
														v79 = int32(0) - v76
													} else {
														v79 = int32(0)
													}
													v80 = int32(-1)
													if l1 != 0 {
														v83 = v80
													} else {
														v83 = v76 + v80
													}
													F_addListRangeReply(m, l0, v41, v79, v83, base.B2i32(l1 != int32(0)))
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return
													} else {
														F_listTypeDelRange(m, v41, v79, v76)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return
														} else {
															v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
															F_listElementsRemoved(m, l0, v91, l1, v41, v76, int32(0))
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return
															} else {
																F_commitDeferredReplyBuffer(m, l0, int32(1))
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(16)
																	return
																}
															}
														}
													}
												}
											}
										} else {
											v55 = F_listTypePop(m, v41, l1)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return
											} else {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
												F_listElementsRemoved(m, l0, v58, l1, v41, int32(1), int32(0))
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return
												} else {
													if v55 == int32(0) {
														F__serverAssert(m, int32(_a1678), int32(_a1674), int32(784))
														mBase = m.M
														v108 = m.ExcPending
														if v108 != 0 {
															return
														} else {
															F_abort(m)
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														F_addReplyBulk(m, l0, v55)
														mBase = m.M
														v66 = m.ExcPending
														if v66 != 0 {
															return
														} else {
															F_decrRefCount(m, v55)
															mBase = m.M
															v68 = m.ExcPending
															if v68 != 0 {
																return
															} else {
																m.G0 = v9 + int32(16)
																return
															}
														}
													}
												}
											}
										}
									} else {
										if v48 != 0 {
											if v48 != 0 {
												v69 = F_listTypeLength(m, v41)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return
												} else {
													v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
													F_initDeferredReplyBuffer(m, l0)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return
													} else {
														if v71 < v69 {
															v76 = v71
														} else {
															v76 = v69
														}
														if l1 != 0 {
															v79 = int32(0) - v76
														} else {
															v79 = int32(0)
														}
														v80 = int32(-1)
														if l1 != 0 {
															v83 = v80
														} else {
															v83 = v76 + v80
														}
														F_addListRangeReply(m, l0, v41, v79, v83, base.B2i32(l1 != int32(0)))
														mBase = m.M
														v87 = m.ExcPending
														if v87 != 0 {
															return
														} else {
															F_listTypeDelRange(m, v41, v79, v76)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return
															} else {
																v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
																F_listElementsRemoved(m, l0, v91, l1, v41, v76, int32(0))
																mBase = m.M
																v94 = m.ExcPending
																if v94 != 0 {
																	return
																} else {
																	F_commitDeferredReplyBuffer(m, l0, int32(1))
																	mBase = m.M
																	v97 = m.ExcPending
																	if v97 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(16)
																		return
																	}
																}
															}
														}
													}
												}
											} else {
												v55 = F_listTypePop(m, v41, l1)
												mBase = m.M
												v56 = m.ExcPending
												if v56 != 0 {
													return
												} else {
													v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
													F_listElementsRemoved(m, l0, v58, l1, v41, int32(1), int32(0))
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return
													} else {
														if v55 == int32(0) {
															F__serverAssert(m, int32(_a1678), int32(_a1674), int32(784))
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
																return
															} else {
																F_abort(m)
																mBase = m.M
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															F_addReplyBulk(m, l0, v55)
															mBase = m.M
															v66 = m.ExcPending
															if v66 != 0 {
																return
															} else {
																F_decrRefCount(m, v55)
																mBase = m.M
																v68 = m.ExcPending
																if v68 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(16)
																	return
																}
															}
														}
													}
												}
											}
										} else {
											v52 = *(*int32)(unsafe.Add(mBase, _consts[542]))
											F_addReply(m, l0, v52)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return
											} else {
												m.G0 = v9 + int32(16)
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
		} else {
			v33 = int32(_a1679)
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v33+v36<<(uint(int32(2))%32))))
			v41 = F_lookupKeyWriteOrReply(m, l0, v35, v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				if v41 == int32(0) {
					m.G0 = v9 + int32(16)
					return
				} else {
					v46 = F_checkType(m, l0, v41, int32(1))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						if v46 != 0 {
							m.G0 = v9 + int32(16)
							return
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
							if v11 != int32(3) {
								if v48 != 0 {
									v69 = F_listTypeLength(m, v41)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
										F_initDeferredReplyBuffer(m, l0)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return
										} else {
											if v71 < v69 {
												v76 = v71
											} else {
												v76 = v69
											}
											if l1 != 0 {
												v79 = int32(0) - v76
											} else {
												v79 = int32(0)
											}
											v80 = int32(-1)
											if l1 != 0 {
												v83 = v80
											} else {
												v83 = v76 + v80
											}
											F_addListRangeReply(m, l0, v41, v79, v83, base.B2i32(l1 != int32(0)))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												F_listTypeDelRange(m, v41, v79, v76)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
													F_listElementsRemoved(m, l0, v91, l1, v41, v76, int32(0))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return
													} else {
														F_commitDeferredReplyBuffer(m, l0, int32(1))
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return
														} else {
															m.G0 = v9 + int32(16)
															return
														}
													}
												}
											}
										}
									}
								} else {
									v55 = F_listTypePop(m, v41, l1)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
										F_listElementsRemoved(m, l0, v58, l1, v41, int32(1), int32(0))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return
										} else {
											if v55 == int32(0) {
												F__serverAssert(m, int32(_a1678), int32(_a1674), int32(784))
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												F_addReplyBulk(m, l0, v55)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return
												} else {
													F_decrRefCount(m, v55)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return
													} else {
														m.G0 = v9 + int32(16)
														return
													}
												}
											}
										}
									}
								}
							} else {
								if v48 != 0 {
									if v48 != 0 {
										v69 = F_listTypeLength(m, v41)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
											F_initDeferredReplyBuffer(m, l0)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												if v71 < v69 {
													v76 = v71
												} else {
													v76 = v69
												}
												if l1 != 0 {
													v79 = int32(0) - v76
												} else {
													v79 = int32(0)
												}
												v80 = int32(-1)
												if l1 != 0 {
													v83 = v80
												} else {
													v83 = v76 + v80
												}
												F_addListRangeReply(m, l0, v41, v79, v83, base.B2i32(l1 != int32(0)))
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
												} else {
													F_listTypeDelRange(m, v41, v79, v76)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return
													} else {
														v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
														F_listElementsRemoved(m, l0, v91, l1, v41, v76, int32(0))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return
														} else {
															F_commitDeferredReplyBuffer(m, l0, int32(1))
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return
															} else {
																m.G0 = v9 + int32(16)
																return
															}
														}
													}
												}
											}
										}
									} else {
										v55 = F_listTypePop(m, v41, l1)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
											F_listElementsRemoved(m, l0, v58, l1, v41, int32(1), int32(0))
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return
											} else {
												if v55 == int32(0) {
													F__serverAssert(m, int32(_a1678), int32(_a1674), int32(784))
													mBase = m.M
													v108 = m.ExcPending
													if v108 != 0 {
														return
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													F_addReplyBulk(m, l0, v55)
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return
													} else {
														F_decrRefCount(m, v55)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return
														} else {
															m.G0 = v9 + int32(16)
															return
														}
													}
												}
											}
										}
									}
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, _consts[542]))
									F_addReply(m, l0, v52)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										m.G0 = v9 + int32(16)
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
		F_addReplyErrorArity(m, l0)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			m.G0 = v9 + int32(16)
			return
		}
	}
}
func F_scanGenericCommandWithOptions(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int64
	_ = v144
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int64
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int64
	_ = v234
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int64
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v381 int32
	_ = v381
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
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
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v513 int64
	_ = v513
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v640 int32
	_ = v640
	var v649 int32
	_ = v649
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	if l4 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l1 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v24 = v23
	v25 = v22
	goto L1
L3:
	;
	v20 = int32(-1)
	v24 = v20
	v25 = v20
	goto L1
L4:
	;
	F__serverPanic_1(m, int32(_a550), int32(1361), int32(_a569), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L17
	} else {
		goto L157
	}
L5:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L17
	} else {
		goto L125
	}
L6:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v193 = v191 & int32(15)
	if base.Ui32(v193+int32(-3)) < base.Ui32(int32(2)) {
		goto L53
	} else {
		goto L54
	}
L7:
	;
	F__serverAssert(m, int32(_a570), int32(_a550), int32(1301))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L17
	} else {
		goto L51
	}
L8:
	;
	v144 = l2
	v148 = v77 * int32(10)
	goto L41
L9:
	;
	if int32(16383) < v118 {
		goto L7
	} else {
		goto L39
	}
L10:
	;
	v125 = int32(-1)
	if v122 == v125 {
		v138 = v125
		v139 = v124
		goto L8
	} else {
		goto L37
	}
L11:
	;
	if int32(0) <= v119 {
		goto L9
	} else {
		goto L36
	}
L12:
	;
	v118 = v25
	v119 = v24
	goto L11
L13:
	;
	F__serverAssert(m, int32(_a571), int32(_a550), int32(1232))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L17
	} else {
		goto L35
	}
L14:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v18 + int32(96)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v82
	v86 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v86
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v89 == v83 {
		v93 = v83
		goto L29
	} else {
		goto L30
	}
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(int32(3)) <= base.Ui32(v35&int32(15)+int32(-2)) {
		goto L13
	} else {
		goto L19
	}
L16:
	;
	F_vectorInit(m, v18+int32(96), int32(24), int32(8))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	v33 = int32(0)
	v74 = int32(1)
	v75 = v33
	v76 = v33
	goto L14
L19:
	;
	v43 = v35 & int32(255)
	switch v43 + int32(-34) {
	case 0, 2:
		goto L24
	case 1:
		goto L22
	default:
		goto L23
	}
L20:
	;
	F_vectorInit(m, v18+int32(96), int32(24), int32(8))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L17
	} else {
		goto L27
	}
L21:
	;
	v59 = F_objectGetVal(m, l1)
	mBase = m.M
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v63 = int32(0)
	v64 = v60
	v65 = int32(514)
	goto L20
L22:
	;
	F_vectorInit(m, v18+int32(96), int32(24), int32(8))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L17
	} else {
		goto L26
	}
L23:
	;
	if v43 == int32(115) {
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v48 = F_objectGetVal(m, l1)
	mBase = m.M
	v63 = int32(1)
	v64 = v48
	v65 = int32(0)
	goto L20
L25:
	;
	goto L22
L26:
	;
	v188 = int32(0)
	v190 = int32(514)
	goto L6
L27:
	;
	if v64 == int32(0) {
		v188 = v63
		v190 = v65
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v74 = v63
	v75 = v64
	v76 = v65
	goto L14
L29:
	;
	v94 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v93
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v97
	if l1 != 0 {
		goto L12
	} else {
		goto L31
	}
L30:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v93 = v92
	goto L29
L31:
	;
	v101 = int32(-1)
	if v24 != v101 {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	if v89 == int32(0) {
		v122 = v25
		v124 = v101
		goto L10
	} else {
		goto L33
	}
L33:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v107 == int32(0) {
		v122 = v25
		v124 = v101
		goto L10
	} else {
		goto L34
	}
L34:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v118 = v110
	v119 = v110
	goto L11
L35:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	v122 = v118
	v124 = v119
	goto L10
L37:
	;
	F__serverAssert(m, int32(_a572), int32(_a550), int32(1303))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L17
	} else {
		goto L38
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
	if v118 < v119 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v138 = v118
	v139 = v119
	goto L8
L41:
	;
	if l1 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v172 == int64(0) {
		v513 = v172
		v519 = v74
		v521 = v76
		v523 = v138
		goto L5
	} else {
		goto L48
	}
L44:
	;
	v169 = F_hashtableScan(m, v75, base.I32_wrap_i64(v144), int32(516), v18+int32(48))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L17
	} else {
		goto L47
	}
L45:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v163 = F_kvstoreScan(m, v158, v144, v139, v138, int32(515), int32(0), v18+int32(48))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L17
	} else {
		goto L46
	}
L46:
	;
	v172 = v163
	goto L43
L47:
	;
	v172 = base.I64_extend_i32_u(v169)
	goto L43
L48:
	;
	if v148 == int32(0) {
		v513 = v172
		v519 = v74
		v521 = v76
		v523 = v138
		goto L5
	} else {
		goto L49
	}
L49:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v179 < v180 {
		v144 = v172
		v148 = v148 + int32(-1)
		goto L41
	} else {
		goto L50
	}
L50:
	;
	v513 = v172
	v519 = v74
	v521 = v76
	v523 = v138
	goto L5
L51:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v513 = int64(0)
	v519 = v188
	v521 = v190
	v523 = v25
	goto L5
L53:
	;
	if v191&int32(240) != int32(176) {
		goto L4
	} else {
		goto L93
	}
L54:
	;
	if v193 != int32(2) {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v200 = F_setTypeInitIterator(m, l1)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L17
	} else {
		goto L57
	}
L56:
	;
	F_setTypeReleaseIterator(m, v200)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L17
	} else {
		goto L92
	}
L57:
	;
	v208 = F_setTypeNext(m, v200, v18+int32(44), v18+int32(40), v18+int32(32))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L17
	} else {
		goto L58
	}
L58:
	;
	if v208 == int32(-1) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L60
L60:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	if v227 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L56
L62:
	;
	if v279 != 0 {
		goto L74
	} else {
		goto L75
	}
L63:
	;
	v232 = v18 + int32(48)
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	if v234 <= int64(-1) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v278 = v230
	v279 = v227
	goto L62
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v275
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v278 = v275
	v279 = v277
	goto L62
L66:
	;
	v275 = int32(0)
	goto L65
L68:
	;
	v256 = F_ull2string(m, v252, v253, v254)
	mBase = m.M
	if v256 == int32(0) {
		goto L66
	} else {
		goto L72
	}
L69:
	;
	goto L71
L70:
	;
	v252 = v232
	v253 = int32(21)
	v254 = v234
	v255 = int32(0)
	goto L68
L71:
	;
	v243 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v232))) = uint8(v243)
	v252 = v18 + int32(49)
	v253 = int32(20)
	v254 = int64(0) - v234
	v255 = int32(1)
	goto L68
L72:
	;
	v275 = v256 + v255
	goto L65
L74:
	;
	v282 = v279
	goto L76
L75:
	;
	v282 = v18 + int32(48)
	goto L76
L76:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v283 == int32(0) {
		v306 = v278
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v345 = F_setTypeNext(m, v200, v18+int32(44), v18+int32(40), v18+int32(32))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L17
	} else {
		goto L90
	}
L78:
	;
	v308 = F_sdsnewlen(m, v282, v306)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L17
	} else {
		goto L88
	}
L79:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v288 = int32(0)
	v290 = m.G0
	v291 = int32(16)
	v292 = v290 - v291
	m.G0 = v292
	*(*int32)(unsafe.Add(mBase, uint32(v292)+12)) = v288
	v299 = F_stringmatchlen_impl(m, v286, v287, v282, v278, v288, v292+int32(12), v288)
	mBase = m.M
	m.G0 = v292 + v291
	goto L80
L80:
	;
	if v299 == int32(0) {
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v306 = v305
	goto L78
L82:
	;
	v332 = F_vectorPush(m, v18+int32(96))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L17
	} else {
		goto L89
	}
L83:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v308+int32(-17))))
	v329 = v328
	goto L82
L84:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v308+int32(-9))))
	v329 = v325
	goto L82
L85:
	;
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v308+int32(-5)))))
	v329 = v322
	goto L82
L86:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308+int32(-3)))))
	v329 = v319
	goto L82
L87:
	;
	v329 = int32(base.Ui32(v312) >> (uint(int32(3)) % 32))
	goto L82
L88:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308+int32(-1)))))
	switch v312 & int32(7) {
	case 0:
		goto L87
	case 1:
		goto L86
	case 2:
		goto L85
	case 3:
		goto L84
	case 4:
		goto L83
	default:
		v329 = int32(0)
		goto L82
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v332)+4)) = v329
	*(*int32)(unsafe.Add(mBase, uint32(v332))) = v308
	goto L77
L90:
	;
	if v345 != int32(-1) {
		goto L60
	} else {
		goto L91
	}
L91:
	;
	goto L61
L92:
	;
	goto L52
L93:
	;
	v370 = F_objectGetVal(m, l1)
	mBase = m.M
	v371 = F_lpFirst(m, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L17
	} else {
		goto L94
	}
L94:
	;
	if v371 == int32(0) {
		goto L52
	} else {
		goto L95
	}
L95:
	;
	v381 = v371
	goto L96
L96:
	;
	v394 = F_lpGet(m, v381, v18+int32(32), v18+int32(48))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L17
	} else {
		goto L98
	}
L97:
	;
	goto L52
L98:
	;
	v396 = F_objectGetVal(m, l1)
	mBase = m.M
	v397 = F_lpNext(m, v396, v381)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L17
	} else {
		goto L99
	}
L99:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v399 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v492 = F_objectGetVal(m, l1)
	mBase = m.M
	v493 = F_lpNext(m, v492, v397)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L17
	} else {
		goto L123
	}
L101:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v424 = F_sdsnewlen(m, v394, v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L17
	} else {
		goto L111
	}
L102:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v405 = int32(0)
	v407 = m.G0
	v408 = int32(16)
	v409 = v407 - v408
	m.G0 = v409
	*(*int32)(unsafe.Add(mBase, uint32(v409)+12)) = v405
	v416 = F_stringmatchlen_impl(m, v402, v403, v394, v404, v405, v409+int32(12), v405)
	mBase = m.M
	m.G0 = v409 + v408
	goto L103
L103:
	;
	if v416 == int32(0) {
		goto L100
	} else {
		goto L104
	}
L104:
	;
	goto L101
L105:
	;
	v448 = F_vectorPush(m, v18+int32(96))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L17
	} else {
		goto L112
	}
L106:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v424+int32(-17))))
	v445 = v444
	goto L105
L107:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v424+int32(-9))))
	v445 = v441
	goto L105
L108:
	;
	v438 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v424+int32(-5)))))
	v445 = v438
	goto L105
L109:
	;
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424+int32(-3)))))
	v445 = v435
	goto L105
L110:
	;
	v445 = int32(base.Ui32(v428) >> (uint(int32(3)) % 32))
	goto L105
L111:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424+int32(-1)))))
	switch v428 & int32(7) {
	case 0:
		goto L110
	case 1:
		goto L109
	case 2:
		goto L108
	case 3:
		goto L107
	case 4:
		goto L106
	default:
		v445 = int32(0)
		goto L105
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v448)+4)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v448))) = v424
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if v452 != 0 {
		goto L100
	} else {
		goto L113
	}
L113:
	;
	v458 = F_lpGet(m, v397, v18+int32(32), v18+int32(48))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L17
	} else {
		goto L120
	}
L114:
	;
	v485 = F_vectorPush(m, v18+int32(96))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L17
	} else {
		goto L122
	}
L115:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v461+int32(-17))))
	v482 = v481
	goto L114
L116:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v461+int32(-9))))
	v482 = v478
	goto L114
L117:
	;
	v475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v461+int32(-5)))))
	v482 = v475
	goto L114
L118:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461+int32(-3)))))
	v482 = v472
	goto L114
L119:
	;
	v482 = int32(base.Ui32(v465) >> (uint(int32(3)) % 32))
	goto L114
L120:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v461 = F_sdsnewlen(m, v458, v460)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L17
	} else {
		goto L121
	}
L121:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461+int32(-1)))))
	switch v465 & int32(7) {
	case 0:
		goto L119
	case 1:
		goto L118
	case 2:
		goto L117
	case 3:
		goto L116
	case 4:
		goto L115
	default:
		v482 = int32(0)
		goto L114
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485)+4)) = v482
	*(*int32)(unsafe.Add(mBase, uint32(v485))) = v461
	goto L100
L123:
	;
	if v493 != 0 {
		v381 = v493
		goto L96
	} else {
		goto L124
	}
L124:
	;
	goto L97
L125:
	;
	if l4 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(96))+8))
	goto L142
L127:
	;
	F_addReplyBulkLongLong(m, l0, v513)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L17
	} else {
		goto L141
	}
L128:
	;
	if l4 == int32(0) {
		goto L127
	} else {
		goto L136
	}
L129:
	;
	if v513 != int64(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	if v533 != int32(1) {
		goto L127
	} else {
		goto L131
	}
L131:
	;
	if int32(16382) < v523 {
		goto L127
	} else {
		goto L132
	}
L132:
	;
	v538 = F_sdsempty(m)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L17
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v523<<(uint(int32(2))%32) + int32(_a573)
	v548 = F_sdscatfmt(m, v538, int32(_a226), v18)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L17
	} else {
		goto L134
	}
L134:
	;
	F_addReplyBulkSds(m, l0, v548)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L17
	} else {
		goto L135
	}
L135:
	;
	goto L126
L136:
	;
	if v513 == int64(0) {
		goto L127
	} else {
		goto L137
	}
L137:
	;
	v556 = F_sdsempty(m)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L17
	} else {
		goto L138
	}
L138:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v558
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(_a225) + base.I32_wrap_i64(v513)&int32(16383)<<(uint(int32(2))%32)
	v572 = F_sdscatfmt(m, v556, int32(_a574), v18+int32(16))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L17
	} else {
		goto L139
	}
L139:
	;
	F_addReplyBulkSds(m, l0, v572)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L17
	} else {
		goto L140
	}
L140:
	;
	goto L126
L141:
	;
	goto L126
L142:
	;
	F_addReplyArrayLen(m, l0, v581)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L17
	} else {
		goto L143
	}
L143:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(96))+8))
	goto L145
L144:
	;
	F_vectorCleanup(m, v18+int32(96))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L17
	} else {
		goto L156
	}
L145:
	;
	if v586 == int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v593 = int32(0)
	goto L147
L147:
	;
	v607 = F_vectorGet(m, v18+int32(96), v593)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L17
	} else {
		goto L149
	}
L148:
	;
	goto L144
L149:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v607)))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v607)+4))
	F_addReplyBulkCBuffer(m, l0, v609, v610)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L17
	} else {
		goto L150
	}
L150:
	;
	if v519 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v617 = v593 + int32(1)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(96))+8))
	goto L154
L152:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v607)))
	m.T0[v521].(func(*base.Module, int32))(m, v613)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L17
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	if base.Ui32(v617) < base.Ui32(v620) {
		v593 = v617
		goto L147
	} else {
		goto L155
	}
L155:
	;
	goto L148
L156:
	;
	m.G0 = v18 + int32(112)
	return
L157:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
