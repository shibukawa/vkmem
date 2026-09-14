package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_queueClientForReprocessing(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v3&int32(128) != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v3 | int32(128)
		v10 = *(*int32)(unsafe.Add(mBase, _consts[61]))
		v11 = F_listAddNodeTail(m, v10, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_queueMultiCommand(m *base.Module, l0 int32, l1 int64) {
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
	var v27 int32
	_ = v27
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v135 int64
	_ = v135
	var v154 int32
	_ = v154
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
	if v15&int32(4128) != 0 {
		m.G0 = v13 + int32(16)
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
		if v18 != 0 {
			v26 = v18
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
			if v27 == int32(0) {
				v32 = F_valkey_malloc(m, int32(40))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = int32(2)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(v35))) = v32
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
					v40 = v35
					v41 = v39
					v42 = v34
					if v41 == v42 {
						if v42 < int32(1073741823) {
							v50 = v42 << (uint(int32(1)) % 32)
						} else {
							v50 = int32(2147483647)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v50
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						v55 = F_valkey_realloc(m, v52, v50*int32(20))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
							*(*int32)(unsafe.Add(mBase, uint32(v57))) = v55
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
							v60 = v57
							v61 = v59
							v62 = v55
							v65 = v62 + v61*int32(20)
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
							*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v66
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v68
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v65))) = v70
							v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v72
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
							*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v74
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
							if v76 == int32(0) {
								v109 = v60
								v110 = v61
								v111 = v68
								*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
								v116 = base.I32_wrap_i64(l1)
								*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
								v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
								v135 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
								m.G0 = v13 + int32(16)
								return
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
								if v79 != int32(54) {
									v109 = v60
									v110 = v61
									v111 = v68
									*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
									v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
									v116 = base.I32_wrap_i64(l1)
									*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
									v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
									v135 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
									m.G0 = v13 + int32(16)
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
									v86 = m.T0[v76].(func(*base.Module, int32, int32, int32) int32)(m, v70, v68, v13+int32(12))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										if v86 == int32(0) {
											v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
											v109 = v107
											v110 = v108
											v111 = v106
											*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
											v116 = base.I32_wrap_i64(l1)
											*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
											v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
											*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
											v135 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
											*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
											m.G0 = v13 + int32(16)
											return
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
											if v90 < int32(1) {
												F_valkey_free(m, v86)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
													v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
													v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
													v109 = v107
													v110 = v108
													v111 = v106
													*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
													v116 = base.I32_wrap_i64(l1)
													*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
													v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
													v135 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
													*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
													m.G0 = v13 + int32(16)
													return
												}
											} else {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94<<(uint(int32(2))%32))))
												v99 = F_getLongLongFromObject(m, v98, v13)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													if v99 != 0 {
														F__serverAssert(m, int32(_a765), int32(_a766), int32(125))
														mBase = m.M
														v154 = m.ExcPending
														if v154 != 0 {
															return
														} else {
															F_abort(m)
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
														v102 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
														*(*uint32)(unsafe.Add(mBase, uint32(v101)+52)) = uint32(v102)
														F_valkey_free(m, v86)
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return
														} else {
															v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
															v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
															v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
															v109 = v107
															v110 = v108
															v111 = v106
															*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
															v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
															v116 = base.I32_wrap_i64(l1)
															*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
															v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
															*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
															v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
															v135 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
															*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
															m.G0 = v13 + int32(16)
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
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						v60 = v40
						v61 = v41
						v62 = v44
						v65 = v62 + v61*int32(20)
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v66
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v68
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v65))) = v70
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v72
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v74
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
						if v76 == int32(0) {
							v109 = v60
							v110 = v61
							v111 = v68
							*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
							v116 = base.I32_wrap_i64(l1)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
							v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
							v135 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
							*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
							m.G0 = v13 + int32(16)
							return
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
							if v79 != int32(54) {
								v109 = v60
								v110 = v61
								v111 = v68
								*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
								v116 = base.I32_wrap_i64(l1)
								*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
								v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
								v135 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
								m.G0 = v13 + int32(16)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
								v86 = m.T0[v76].(func(*base.Module, int32, int32, int32) int32)(m, v70, v68, v13+int32(12))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									if v86 == int32(0) {
										v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
										v109 = v107
										v110 = v108
										v111 = v106
										*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
										v116 = base.I32_wrap_i64(l1)
										*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
										v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
										v135 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
										*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
										m.G0 = v13 + int32(16)
										return
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
										if v90 < int32(1) {
											F_valkey_free(m, v86)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return
											} else {
												v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
												v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
												v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
												v109 = v107
												v110 = v108
												v111 = v106
												*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
												v116 = base.I32_wrap_i64(l1)
												*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
												v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
												v135 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
												m.G0 = v13 + int32(16)
												return
											}
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94<<(uint(int32(2))%32))))
											v99 = F_getLongLongFromObject(m, v98, v13)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												if v99 != 0 {
													F__serverAssert(m, int32(_a765), int32(_a766), int32(125))
													mBase = m.M
													v154 = m.ExcPending
													if v154 != 0 {
														return
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
													v102 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
													*(*uint32)(unsafe.Add(mBase, uint32(v101)+52)) = uint32(v102)
													F_valkey_free(m, v86)
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return
													} else {
														v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
														v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
														v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
														v109 = v107
														v110 = v108
														v111 = v106
														*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
														v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
														v116 = base.I32_wrap_i64(l1)
														*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
														*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
														v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
														*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
														v135 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
														*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
														m.G0 = v13 + int32(16)
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
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
				v40 = v26
				v41 = v27
				v42 = v30
				if v41 == v42 {
					if v42 < int32(1073741823) {
						v50 = v42 << (uint(int32(1)) % 32)
					} else {
						v50 = int32(2147483647)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v50
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					v55 = F_valkey_realloc(m, v52, v50*int32(20))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						*(*int32)(unsafe.Add(mBase, uint32(v57))) = v55
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
						v60 = v57
						v61 = v59
						v62 = v55
						v65 = v62 + v61*int32(20)
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v66
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v68
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v65))) = v70
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v72
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v74
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
						if v76 == int32(0) {
							v109 = v60
							v110 = v61
							v111 = v68
							*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
							v116 = base.I32_wrap_i64(l1)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
							v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
							v135 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
							*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
							m.G0 = v13 + int32(16)
							return
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
							if v79 != int32(54) {
								v109 = v60
								v110 = v61
								v111 = v68
								*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
								v116 = base.I32_wrap_i64(l1)
								*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
								v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
								v135 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
								m.G0 = v13 + int32(16)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
								v86 = m.T0[v76].(func(*base.Module, int32, int32, int32) int32)(m, v70, v68, v13+int32(12))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									if v86 == int32(0) {
										v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
										v109 = v107
										v110 = v108
										v111 = v106
										*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
										v116 = base.I32_wrap_i64(l1)
										*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
										v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
										v135 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
										*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
										m.G0 = v13 + int32(16)
										return
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
										if v90 < int32(1) {
											F_valkey_free(m, v86)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return
											} else {
												v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
												v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
												v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
												v109 = v107
												v110 = v108
												v111 = v106
												*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
												v116 = base.I32_wrap_i64(l1)
												*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
												v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
												v135 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
												m.G0 = v13 + int32(16)
												return
											}
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94<<(uint(int32(2))%32))))
											v99 = F_getLongLongFromObject(m, v98, v13)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												if v99 != 0 {
													F__serverAssert(m, int32(_a765), int32(_a766), int32(125))
													mBase = m.M
													v154 = m.ExcPending
													if v154 != 0 {
														return
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
													v102 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
													*(*uint32)(unsafe.Add(mBase, uint32(v101)+52)) = uint32(v102)
													F_valkey_free(m, v86)
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return
													} else {
														v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
														v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
														v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
														v109 = v107
														v110 = v108
														v111 = v106
														*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
														v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
														v116 = base.I32_wrap_i64(l1)
														*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
														*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
														v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
														*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
														v135 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
														*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
														m.G0 = v13 + int32(16)
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
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					v60 = v40
					v61 = v41
					v62 = v44
					v65 = v62 + v61*int32(20)
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v66
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v68
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v65))) = v70
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v72
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
					*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v74
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
					if v76 == int32(0) {
						v109 = v60
						v110 = v61
						v111 = v68
						*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
						v116 = base.I32_wrap_i64(l1)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
						v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
						v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
						v135 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
						*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
						m.G0 = v13 + int32(16)
						return
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
						if v79 != int32(54) {
							v109 = v60
							v110 = v61
							v111 = v68
							*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
							v116 = base.I32_wrap_i64(l1)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
							v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
							v135 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
							*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
							m.G0 = v13 + int32(16)
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
							v86 = m.T0[v76].(func(*base.Module, int32, int32, int32) int32)(m, v70, v68, v13+int32(12))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								if v86 == int32(0) {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
									v109 = v107
									v110 = v108
									v111 = v106
									*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
									v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
									v116 = base.I32_wrap_i64(l1)
									*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
									v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
									v135 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
									m.G0 = v13 + int32(16)
									return
								} else {
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
									if v90 < int32(1) {
										F_valkey_free(m, v86)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return
										} else {
											v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
											v109 = v107
											v110 = v108
											v111 = v106
											*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
											v116 = base.I32_wrap_i64(l1)
											*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
											v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
											*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
											v135 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
											*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
											m.G0 = v13 + int32(16)
											return
										}
									} else {
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94<<(uint(int32(2))%32))))
										v99 = F_getLongLongFromObject(m, v98, v13)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											if v99 != 0 {
												F__serverAssert(m, int32(_a765), int32(_a766), int32(125))
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
													return
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
												v102 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
												*(*uint32)(unsafe.Add(mBase, uint32(v101)+52)) = uint32(v102)
												F_valkey_free(m, v86)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
													v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
													v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
													v109 = v107
													v110 = v108
													v111 = v106
													*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
													v116 = base.I32_wrap_i64(l1)
													*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
													v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
													v135 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
													*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
													m.G0 = v13 + int32(16)
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
		} else {
			v20 = F_valkey_calloc(m, int32(56))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v20
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v24
				v26 = v20
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
				if v27 == int32(0) {
					v32 = F_valkey_malloc(m, int32(40))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						v34 = int32(2)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(v35))) = v32
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
						v40 = v35
						v41 = v39
						v42 = v34
						if v41 == v42 {
							if v42 < int32(1073741823) {
								v50 = v42 << (uint(int32(1)) % 32)
							} else {
								v50 = int32(2147483647)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v50
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
							v55 = F_valkey_realloc(m, v52, v50*int32(20))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
								*(*int32)(unsafe.Add(mBase, uint32(v57))) = v55
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
								v60 = v57
								v61 = v59
								v62 = v55
								v65 = v62 + v61*int32(20)
								v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
								*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v66
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v68
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v65))) = v70
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v72
								v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
								*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v74
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
								if v76 == int32(0) {
									v109 = v60
									v110 = v61
									v111 = v68
									*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
									v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
									v116 = base.I32_wrap_i64(l1)
									*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
									v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
									v135 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
									m.G0 = v13 + int32(16)
									return
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
									if v79 != int32(54) {
										v109 = v60
										v110 = v61
										v111 = v68
										*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
										v116 = base.I32_wrap_i64(l1)
										*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
										v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
										v135 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
										*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
										m.G0 = v13 + int32(16)
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
										v86 = m.T0[v76].(func(*base.Module, int32, int32, int32) int32)(m, v70, v68, v13+int32(12))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											if v86 == int32(0) {
												v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
												v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
												v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
												v109 = v107
												v110 = v108
												v111 = v106
												*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
												v116 = base.I32_wrap_i64(l1)
												*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
												v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
												v135 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
												m.G0 = v13 + int32(16)
												return
											} else {
												v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
												if v90 < int32(1) {
													F_valkey_free(m, v86)
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return
													} else {
														v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
														v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
														v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
														v109 = v107
														v110 = v108
														v111 = v106
														*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
														v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
														v116 = base.I32_wrap_i64(l1)
														*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
														*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
														v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
														*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
														v135 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
														*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
														m.G0 = v13 + int32(16)
														return
													}
												} else {
													v93 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
													v94 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94<<(uint(int32(2))%32))))
													v99 = F_getLongLongFromObject(m, v98, v13)
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return
													} else {
														if v99 != 0 {
															F__serverAssert(m, int32(_a765), int32(_a766), int32(125))
															mBase = m.M
															v154 = m.ExcPending
															if v154 != 0 {
																return
															} else {
																F_abort(m)
																mBase = m.M
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
															v102 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
															*(*uint32)(unsafe.Add(mBase, uint32(v101)+52)) = uint32(v102)
															F_valkey_free(m, v86)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return
															} else {
																v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
																v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
																v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
																v109 = v107
																v110 = v108
																v111 = v106
																*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
																v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
																v116 = base.I32_wrap_i64(l1)
																*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
																v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
																*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
																v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
																v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
																*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
																v135 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
																*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
																m.G0 = v13 + int32(16)
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
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
							v60 = v40
							v61 = v41
							v62 = v44
							v65 = v62 + v61*int32(20)
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
							*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v66
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v68
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v65))) = v70
							v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v72
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
							*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v74
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
							if v76 == int32(0) {
								v109 = v60
								v110 = v61
								v111 = v68
								*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
								v116 = base.I32_wrap_i64(l1)
								*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
								v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
								v135 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
								m.G0 = v13 + int32(16)
								return
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
								if v79 != int32(54) {
									v109 = v60
									v110 = v61
									v111 = v68
									*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
									v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
									v116 = base.I32_wrap_i64(l1)
									*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
									v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
									v135 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
									m.G0 = v13 + int32(16)
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
									v86 = m.T0[v76].(func(*base.Module, int32, int32, int32) int32)(m, v70, v68, v13+int32(12))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										if v86 == int32(0) {
											v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
											v109 = v107
											v110 = v108
											v111 = v106
											*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
											v116 = base.I32_wrap_i64(l1)
											*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
											v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
											*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
											v135 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
											*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
											m.G0 = v13 + int32(16)
											return
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
											if v90 < int32(1) {
												F_valkey_free(m, v86)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
													v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
													v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
													v109 = v107
													v110 = v108
													v111 = v106
													*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
													v116 = base.I32_wrap_i64(l1)
													*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
													v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
													v135 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
													*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
													m.G0 = v13 + int32(16)
													return
												}
											} else {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94<<(uint(int32(2))%32))))
												v99 = F_getLongLongFromObject(m, v98, v13)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													if v99 != 0 {
														F__serverAssert(m, int32(_a765), int32(_a766), int32(125))
														mBase = m.M
														v154 = m.ExcPending
														if v154 != 0 {
															return
														} else {
															F_abort(m)
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
														v102 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
														*(*uint32)(unsafe.Add(mBase, uint32(v101)+52)) = uint32(v102)
														F_valkey_free(m, v86)
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return
														} else {
															v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
															v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
															v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
															v109 = v107
															v110 = v108
															v111 = v106
															*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
															v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
															v116 = base.I32_wrap_i64(l1)
															*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
															v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
															*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
															v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
															v135 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
															*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
															m.G0 = v13 + int32(16)
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
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
					v40 = v26
					v41 = v27
					v42 = v30
					if v41 == v42 {
						if v42 < int32(1073741823) {
							v50 = v42 << (uint(int32(1)) % 32)
						} else {
							v50 = int32(2147483647)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v50
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						v55 = F_valkey_realloc(m, v52, v50*int32(20))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
							*(*int32)(unsafe.Add(mBase, uint32(v57))) = v55
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
							v60 = v57
							v61 = v59
							v62 = v55
							v65 = v62 + v61*int32(20)
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
							*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v66
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v68
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v65))) = v70
							v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v72
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
							*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v74
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
							if v76 == int32(0) {
								v109 = v60
								v110 = v61
								v111 = v68
								*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
								v116 = base.I32_wrap_i64(l1)
								*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
								v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
								v135 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
								m.G0 = v13 + int32(16)
								return
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
								if v79 != int32(54) {
									v109 = v60
									v110 = v61
									v111 = v68
									*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
									v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
									v116 = base.I32_wrap_i64(l1)
									*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
									v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
									v135 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
									*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
									m.G0 = v13 + int32(16)
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
									v86 = m.T0[v76].(func(*base.Module, int32, int32, int32) int32)(m, v70, v68, v13+int32(12))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										if v86 == int32(0) {
											v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
											v109 = v107
											v110 = v108
											v111 = v106
											*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
											v116 = base.I32_wrap_i64(l1)
											*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
											v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
											*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
											v135 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
											*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
											m.G0 = v13 + int32(16)
											return
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
											if v90 < int32(1) {
												F_valkey_free(m, v86)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
													v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
													v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
													v109 = v107
													v110 = v108
													v111 = v106
													*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
													v116 = base.I32_wrap_i64(l1)
													*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
													v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
													v135 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
													*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
													m.G0 = v13 + int32(16)
													return
												}
											} else {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94<<(uint(int32(2))%32))))
												v99 = F_getLongLongFromObject(m, v98, v13)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													if v99 != 0 {
														F__serverAssert(m, int32(_a765), int32(_a766), int32(125))
														mBase = m.M
														v154 = m.ExcPending
														if v154 != 0 {
															return
														} else {
															F_abort(m)
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
														v102 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
														*(*uint32)(unsafe.Add(mBase, uint32(v101)+52)) = uint32(v102)
														F_valkey_free(m, v86)
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return
														} else {
															v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
															v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
															v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
															v109 = v107
															v110 = v108
															v111 = v106
															*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
															v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
															v116 = base.I32_wrap_i64(l1)
															*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
															v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
															*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
															v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
															v135 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
															*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
															m.G0 = v13 + int32(16)
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
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						v60 = v40
						v61 = v41
						v62 = v44
						v65 = v62 + v61*int32(20)
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v66
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v68
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v65))) = v70
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v72
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v74
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
						if v76 == int32(0) {
							v109 = v60
							v110 = v61
							v111 = v68
							*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
							v116 = base.I32_wrap_i64(l1)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
							v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
							v135 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
							*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
							m.G0 = v13 + int32(16)
							return
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
							if v79 != int32(54) {
								v109 = v60
								v110 = v61
								v111 = v68
								*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
								v116 = base.I32_wrap_i64(l1)
								*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
								v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
								v135 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
								*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
								m.G0 = v13 + int32(16)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
								v86 = m.T0[v76].(func(*base.Module, int32, int32, int32) int32)(m, v70, v68, v13+int32(12))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									if v86 == int32(0) {
										v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
										v109 = v107
										v110 = v108
										v111 = v106
										*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
										v116 = base.I32_wrap_i64(l1)
										*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
										v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
										v135 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
										*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
										m.G0 = v13 + int32(16)
										return
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
										if v90 < int32(1) {
											F_valkey_free(m, v86)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return
											} else {
												v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
												v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
												v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
												v109 = v107
												v110 = v108
												v111 = v106
												*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
												v116 = base.I32_wrap_i64(l1)
												*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
												v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
												v135 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
												*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
												m.G0 = v13 + int32(16)
												return
											}
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94<<(uint(int32(2))%32))))
											v99 = F_getLongLongFromObject(m, v98, v13)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												if v99 != 0 {
													F__serverAssert(m, int32(_a765), int32(_a766), int32(125))
													mBase = m.M
													v154 = m.ExcPending
													if v154 != 0 {
														return
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
													v102 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
													*(*uint32)(unsafe.Add(mBase, uint32(v101)+52)) = uint32(v102)
													F_valkey_free(m, v86)
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return
													} else {
														v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
														v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
														v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
														v109 = v107
														v110 = v108
														v111 = v106
														*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110 + int32(1)
														v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
														v116 = base.I32_wrap_i64(l1)
														*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v115 | v116
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
														*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119 | (v116 ^ int32(-1))
														v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
														*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v111<<(uint(int32(2))%32) + v126 + v128
														v135 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v135
														*(*int64)(unsafe.Add(mBase, uint32(l0+int32(20)))) = v135
														m.G0 = v13 + int32(16)
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
func F_quitCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_addReply(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v6 | int32(64)
		return
	}
}
