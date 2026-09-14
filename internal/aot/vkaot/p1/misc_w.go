package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F___wake_3(m *base.Module) {
	return
}
func F___wasi_timestamp_to_timespec(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v11 int64
	_ = v11
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v6 = int64(1000000000)
	v7 = base.I64_div_u_s(l1, v6)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v7
	v11 = l1 - v7*v6
	*(*uint32)(unsafe.Add(mBase, uint32(l0)+8)) = uint32(v11)
	return
}
func F___wasm_setjmp_test(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5 != l1 {
		v8 = int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v8 = v7
	}
	return v8
}
func F_waitaofCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
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
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
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
	var v81 int32
	_ = v81
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v16 = F_getRangeLongFromObjectOrReply(m, l0, v12, v2, int32(1), v9, v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if v16 != 0 {
			m.G0 = v9 + int32(16)
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
			v23 = F_getPositiveLongFromObjectOrReply(m, l0, v19, v9+int32(4), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				if v23 != 0 {
					m.G0 = v9 + int32(16)
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
					v30 = F_getTimeoutFromObjectOrReply(m, l0, v26, v9+int32(8), int32(1))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						if v30 != 0 {
							m.G0 = v9 + int32(16)
							return
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, _consts[166]))
							if v33 == int32(0) {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
								if v39 == int32(0) {
									v47 = int32(0)
									v48 = *(*int32)(unsafe.Add(mBase, _consts[333]))
									if base.B2i32(v48 != v47) == int32(0) {
										v55 = l0
										v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)+48))
										v57 = int32(0)
										v60 = m.G0
										v62 = v60 - int32(16)
										m.G0 = v62
										v65 = *(*int32)(unsafe.Add(mBase, _consts[78]))
										v67 = v62 + int32(8)
										F_listRewind(m, v65, v67)
										mBase = m.M
										v72 = F_listNext(m, v67)
										mBase = m.M
										if v72 == v57 {
											v93 = v57
										} else {
											v77 = v57
											v78 = v72
											for {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
												v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+104))
												v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
												if v81 != int32(9) {
													v87 = v77
												} else {
													v84 = *(*int64)(unsafe.Add(mBase, uint32(v80)+72))
													v87 = v77 + base.B2i32(v56 <= v84)
												}
												v90 = F_listNext(m, v62+int32(8))
												mBase = m.M
												if v90 != 0 {
													v77 = v87
													v78 = v90
													continue
												} else {
													break
												}
												break
											}
											v93 = v87
										}
										m.G0 = v62 + int32(16)
										v99 = *(*int64)(unsafe.Add(mBase, _consts[49]))
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
										if v93 < v100 {
											v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
											if v105&int32(16) == int32(0) {
												v120 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
												F_blockClientForReplicaAck(m, l0, v120, v56, v100, v121)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[637])) = int32(1)
													m.G0 = v9 + int32(16)
													return
												}
											} else {
												F_addReplyArrayLen(m, l0, int32(2))
												mBase = m.M
												v112 = m.ExcPending
												if v112 != 0 {
													return
												} else {
													F_addReplyLongLong(m, l0, base.I64_extend_i32_u(base.B2i32(v56 <= v99)))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v93))
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return
														} else {
															m.G0 = v9 + int32(16)
															return
														}
													}
												}
											}
										} else {
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
											if v102 <= base.B2i32(v56 <= v99) {
												F_addReplyArrayLen(m, l0, int32(2))
												mBase = m.M
												v112 = m.ExcPending
												if v112 != 0 {
													return
												} else {
													F_addReplyLongLong(m, l0, base.I64_extend_i32_u(base.B2i32(v56 <= v99)))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v93))
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return
														} else {
															m.G0 = v9 + int32(16)
															return
														}
													}
												}
											} else {
												v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
												if v105&int32(16) == int32(0) {
													v120 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
													F_blockClientForReplicaAck(m, l0, v120, v56, v100, v121)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[637])) = int32(1)
														m.G0 = v9 + int32(16)
														return
													}
												} else {
													F_addReplyArrayLen(m, l0, int32(2))
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return
													} else {
														F_addReplyLongLong(m, l0, base.I64_extend_i32_u(base.B2i32(v56 <= v99)))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v93))
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
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
										v53 = F_scriptGetCaller(m)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											v55 = v53
											v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)+48))
											v57 = int32(0)
											v60 = m.G0
											v62 = v60 - int32(16)
											m.G0 = v62
											v65 = *(*int32)(unsafe.Add(mBase, _consts[78]))
											v67 = v62 + int32(8)
											F_listRewind(m, v65, v67)
											mBase = m.M
											v72 = F_listNext(m, v67)
											mBase = m.M
											if v72 == v57 {
												v93 = v57
											} else {
												v77 = v57
												v78 = v72
												for {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
													v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+104))
													v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
													if v81 != int32(9) {
														v87 = v77
													} else {
														v84 = *(*int64)(unsafe.Add(mBase, uint32(v80)+72))
														v87 = v77 + base.B2i32(v56 <= v84)
													}
													v90 = F_listNext(m, v62+int32(8))
													mBase = m.M
													if v90 != 0 {
														v77 = v87
														v78 = v90
														continue
													} else {
														break
													}
													break
												}
												v93 = v87
											}
											m.G0 = v62 + int32(16)
											v99 = *(*int64)(unsafe.Add(mBase, _consts[49]))
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
											if v93 < v100 {
												v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
												if v105&int32(16) == int32(0) {
													v120 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
													F_blockClientForReplicaAck(m, l0, v120, v56, v100, v121)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[637])) = int32(1)
														m.G0 = v9 + int32(16)
														return
													}
												} else {
													F_addReplyArrayLen(m, l0, int32(2))
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return
													} else {
														F_addReplyLongLong(m, l0, base.I64_extend_i32_u(base.B2i32(v56 <= v99)))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v93))
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return
															} else {
																m.G0 = v9 + int32(16)
																return
															}
														}
													}
												}
											} else {
												v102 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
												if v102 <= base.B2i32(v56 <= v99) {
													F_addReplyArrayLen(m, l0, int32(2))
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return
													} else {
														F_addReplyLongLong(m, l0, base.I64_extend_i32_u(base.B2i32(v56 <= v99)))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v93))
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return
															} else {
																m.G0 = v9 + int32(16)
																return
															}
														}
													}
												} else {
													v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
													if v105&int32(16) == int32(0) {
														v120 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
														F_blockClientForReplicaAck(m, l0, v120, v56, v100, v121)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[637])) = int32(1)
															m.G0 = v9 + int32(16)
															return
														}
													} else {
														F_addReplyArrayLen(m, l0, int32(2))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															F_addReplyLongLong(m, l0, base.I64_extend_i32_u(base.B2i32(v56 <= v99)))
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return
															} else {
																F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v93))
																mBase = m.M
																v119 = m.ExcPending
																if v119 != 0 {
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
									v43 = *(*int32)(unsafe.Add(mBase, _consts[586]))
									if v43 != 0 {
										v47 = int32(0)
										v48 = *(*int32)(unsafe.Add(mBase, _consts[333]))
										if base.B2i32(v48 != v47) == int32(0) {
											v55 = l0
											v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)+48))
											v57 = int32(0)
											v60 = m.G0
											v62 = v60 - int32(16)
											m.G0 = v62
											v65 = *(*int32)(unsafe.Add(mBase, _consts[78]))
											v67 = v62 + int32(8)
											F_listRewind(m, v65, v67)
											mBase = m.M
											v72 = F_listNext(m, v67)
											mBase = m.M
											if v72 == v57 {
												v93 = v57
											} else {
												v77 = v57
												v78 = v72
												for {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
													v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+104))
													v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
													if v81 != int32(9) {
														v87 = v77
													} else {
														v84 = *(*int64)(unsafe.Add(mBase, uint32(v80)+72))
														v87 = v77 + base.B2i32(v56 <= v84)
													}
													v90 = F_listNext(m, v62+int32(8))
													mBase = m.M
													if v90 != 0 {
														v77 = v87
														v78 = v90
														continue
													} else {
														break
													}
													break
												}
												v93 = v87
											}
											m.G0 = v62 + int32(16)
											v99 = *(*int64)(unsafe.Add(mBase, _consts[49]))
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
											if v93 < v100 {
												v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
												if v105&int32(16) == int32(0) {
													v120 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
													F_blockClientForReplicaAck(m, l0, v120, v56, v100, v121)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[637])) = int32(1)
														m.G0 = v9 + int32(16)
														return
													}
												} else {
													F_addReplyArrayLen(m, l0, int32(2))
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return
													} else {
														F_addReplyLongLong(m, l0, base.I64_extend_i32_u(base.B2i32(v56 <= v99)))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v93))
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return
															} else {
																m.G0 = v9 + int32(16)
																return
															}
														}
													}
												}
											} else {
												v102 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
												if v102 <= base.B2i32(v56 <= v99) {
													F_addReplyArrayLen(m, l0, int32(2))
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return
													} else {
														F_addReplyLongLong(m, l0, base.I64_extend_i32_u(base.B2i32(v56 <= v99)))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v93))
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return
															} else {
																m.G0 = v9 + int32(16)
																return
															}
														}
													}
												} else {
													v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
													if v105&int32(16) == int32(0) {
														v120 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
														F_blockClientForReplicaAck(m, l0, v120, v56, v100, v121)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[637])) = int32(1)
															m.G0 = v9 + int32(16)
															return
														}
													} else {
														F_addReplyArrayLen(m, l0, int32(2))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															F_addReplyLongLong(m, l0, base.I64_extend_i32_u(base.B2i32(v56 <= v99)))
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return
															} else {
																F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v93))
																mBase = m.M
																v119 = m.ExcPending
																if v119 != 0 {
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
											v53 = F_scriptGetCaller(m)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return
											} else {
												v55 = v53
												v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)+48))
												v57 = int32(0)
												v60 = m.G0
												v62 = v60 - int32(16)
												m.G0 = v62
												v65 = *(*int32)(unsafe.Add(mBase, _consts[78]))
												v67 = v62 + int32(8)
												F_listRewind(m, v65, v67)
												mBase = m.M
												v72 = F_listNext(m, v67)
												mBase = m.M
												if v72 == v57 {
													v93 = v57
												} else {
													v77 = v57
													v78 = v72
													for {
														v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
														v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+104))
														v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
														if v81 != int32(9) {
															v87 = v77
														} else {
															v84 = *(*int64)(unsafe.Add(mBase, uint32(v80)+72))
															v87 = v77 + base.B2i32(v56 <= v84)
														}
														v90 = F_listNext(m, v62+int32(8))
														mBase = m.M
														if v90 != 0 {
															v77 = v87
															v78 = v90
															continue
														} else {
															break
														}
														break
													}
													v93 = v87
												}
												m.G0 = v62 + int32(16)
												v99 = *(*int64)(unsafe.Add(mBase, _consts[49]))
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
												if v93 < v100 {
													v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
													if v105&int32(16) == int32(0) {
														v120 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
														F_blockClientForReplicaAck(m, l0, v120, v56, v100, v121)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[637])) = int32(1)
															m.G0 = v9 + int32(16)
															return
														}
													} else {
														F_addReplyArrayLen(m, l0, int32(2))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															F_addReplyLongLong(m, l0, base.I64_extend_i32_u(base.B2i32(v56 <= v99)))
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return
															} else {
																F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v93))
																mBase = m.M
																v119 = m.ExcPending
																if v119 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(16)
																	return
																}
															}
														}
													}
												} else {
													v102 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
													if v102 <= base.B2i32(v56 <= v99) {
														F_addReplyArrayLen(m, l0, int32(2))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															F_addReplyLongLong(m, l0, base.I64_extend_i32_u(base.B2i32(v56 <= v99)))
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return
															} else {
																F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v93))
																mBase = m.M
																v119 = m.ExcPending
																if v119 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(16)
																	return
																}
															}
														}
													} else {
														v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
														if v105&int32(16) == int32(0) {
															v120 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
															v121 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
															F_blockClientForReplicaAck(m, l0, v120, v56, v100, v121)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _consts[637])) = int32(1)
																m.G0 = v9 + int32(16)
																return
															}
														} else {
															F_addReplyArrayLen(m, l0, int32(2))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return
															} else {
																F_addReplyLongLong(m, l0, base.I64_extend_i32_u(base.B2i32(v56 <= v99)))
																mBase = m.M
																v116 = m.ExcPending
																if v116 != 0 {
																	return
																} else {
																	F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v93))
																	mBase = m.M
																	v119 = m.ExcPending
																	if v119 != 0 {
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
										F_addReplyError(m, l0, int32(_a2007))
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											m.G0 = v9 + int32(16)
											return
										}
									}
								}
							} else {
								F_addReplyError(m, l0, int32(_a2008))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
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
func F_waitpid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	v5 = F___syscall_wait4(m, l0, l1, l2, int32(0))
	mBase = m.M
	if base.Ui32(v5) < base.Ui32(int32(-4095)) {
		v13 = v5
	} else {
		v8 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0) - v5
		v13 = int32(-1)
	}
	return v13
}
func F_wcschr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = l0
	goto L12
L2:
	;
	v7 = l0
	goto L4
L3:
	;
	if v11 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v11 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L5:
	;
	goto L3
L6:
	;
	if v11 != l1 {
		v7 = v7 + int32(4)
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v19 = v7
	goto L10
L9:
	;
	v19 = int32(0)
	goto L10
L10:
	;
	return v19
L11:
	;
	return l0 + (v24-l0)>>(uint(int32(2))%32)<<(uint(int32(2))%32)
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v28 != 0 {
		v24 = v24 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	goto L13
}
func F_week_num(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	v5 = int32(53)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = int32(7)
	v11 = base.I32_rem_u_s(v7+int32(6), v10)
	v16 = base.I32_div_u_s(v6-v11+v10, v10)
	v17 = v7 - v6
	v21 = base.I32_rem_u_s(v17+int32(369), v10)
	v24 = v16 + base.B2i32(base.Ui32(v21) < base.Ui32(int32(3)))
	if v24 == v5 {
		v69 = base.I32_rem_u_s(v17+int32(371), int32(7))
		switch v69 + int32(-3) {
		case 0:
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if int32(2147481747) < v72 {
				v77 = v72 + int32(-2000)
			} else {
				v77 = v72
			}
			if v77&int32(3) == int32(0) {
				v84 = v77 + int32(1900)
				v86 = base.I32_rem_s(v84, int32(100))
				if v86 == int32(0) {
					v91 = base.I32_rem_s(v84, int32(400))
					v95 = base.B2i32(v91 == int32(0))
				} else {
					v95 = int32(1)
				}
			} else {
				v95 = int32(0)
			}
			if v95 != 0 {
				v97 = v5
			} else {
				v97 = int32(1)
			}
		case 1:
			v97 = v5
		default:
			v97 = int32(1)
		}
		return v97
	} else {
		if v24 != 0 {
			v97 = v24
			return v97
		} else {
			v27 = int32(52)
			v31 = base.I32_rem_u_s(v17+int32(6), int32(7))
			switch v31 + int32(-4) {
			case 0:
				return int32(53)
			case 1:
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v36 = base.I32_rem_s(v34, int32(400))
				v38 = v36 + int32(-1)
				if int32(2147481747) < v38 {
					v43 = v36 + int32(-2001)
				} else {
					v43 = v38
				}
				if v43&int32(3) == int32(0) {
					v50 = v43 + int32(1900)
					v52 = base.I32_rem_s(v50, int32(100))
					if v52 == int32(0) {
						v57 = base.I32_rem_s(v50, int32(400))
						v61 = base.B2i32(v57 == int32(0))
					} else {
						v61 = int32(1)
					}
				} else {
					v61 = int32(0)
				}
				if v61 == int32(0) {
					v97 = v27
					return v97
				} else {
					return int32(53)
				}
			default:
				v97 = v27
				return v97
			}
		}
	}
}
func F_wrapper_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	v16 = m.Wasi_snapshot_preview1.Fd_write(m, l0, v7+int32(8), int32(1), v7+int32(4))
	mBase = m.M
	if v16 != 0 {
		v18 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = v16
	} else {
	}
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	m.G0 = v7 + int32(16)
	if v16 != 0 {
		v27 = int32(-1)
	} else {
		v27 = v22
	}
	return v27
}
func F_writePingExtensions(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
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
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
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
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
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
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
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
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int64
	_ = v387
	var v388 int64
	_ = v388
	var v389 int64
	_ = v389
	var v390 int64
	_ = v390
	var v391 int64
	_ = v391
	var v392 int64
	_ = v392
	var v393 int64
	_ = v393
	var v395 int64
	_ = v395
	var v397 int64
	_ = v397
	var v399 int64
	_ = v399
	var v401 int64
	_ = v401
	var v403 int64
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v465 int64
	_ = v465
	var v467 int64
	_ = v467
	var v470 int64
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int64
	_ = v478
	var v482 int32
	_ = v482
	var v484 int64
	_ = v484
	var v488 int32
	_ = v488
	var v490 int64
	_ = v490
	var v494 int32
	_ = v494
	var v496 int64
	_ = v496
	var v502 int64
	_ = v502
	var v504 int64
	_ = v504
	var v505 int64
	_ = v505
	var v507 int64
	_ = v507
	var v509 int64
	_ = v509
	var v512 int64
	_ = v512
	var v514 int64
	_ = v514
	var v516 int64
	_ = v516
	var v518 int64
	_ = v518
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int64
	_ = v576
	var v582 int64
	_ = v582
	var v588 int64
	_ = v588
	var v594 int64
	_ = v594
	var v600 int64
	_ = v600
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	v3 = int32(0)
	v16 = l0 + l1*int32(104)
	v18 = v16 + int32(2256)
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = v18
	goto L3
L2:
	;
	v20 = v3
	goto L3
L3:
	;
	v21 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+2312))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(-1)))))
	switch v28 & int32(7) {
	case 0:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	case 3:
		goto L7
	case 4:
		goto L6
	default:
		v67 = v20
		v68 = v21
		v69 = v21
		goto L4
	}
L4:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v24)+2316))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+int32(-1)))))
	switch v74 & int32(7) {
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
		v116 = v67
		v117 = v3
		v118 = v68
		goto L21
	}
L5:
	;
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(-17))))
	v45 = v44
	goto L5
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(-9))))
	v45 = v41
	goto L5
L8:
	;
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+int32(-5)))))
	v45 = v38
	goto L5
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(-3)))))
	v45 = v35
	goto L5
L10:
	;
	v45 = int32(base.Ui32(v28) >> (uint(int32(3)) % 32))
	goto L5
L11:
	;
	v51 = v45&int32(-8) + int32(16)
	v52 = int32(1)
	if l0 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v46 = int32(0)
	v67 = v20
	v68 = v46
	v69 = v46
	goto L4
L13:
	;
	v55 = F___bswap_16_1(m, int32(0))
	mBase = m.M
	goto L15
L14:
	;
	v67 = int32(0)
	v68 = v51
	v69 = v52
	goto L4
L15:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)) = uint16(v55)
	v57 = F___bswap_32_1(m, v51)
	mBase = m.M
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v57
	if v45 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v65 = F___bswap_32_2(m, v57)
	mBase = m.M
	goto L20
L18:
	;
	goto L17
L19:
	;
	v63 = F__emscripten_memcpy_bulkmem(m, v16+int32(2264), v25, v45)
	mBase = m.M
	goto L18
L20:
	;
	v67 = v18 + v65
	v68 = v51
	v69 = v52
	goto L4
L21:
	;
	v122 = int32(0)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v24)+2304))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+int32(-1)))))
	switch v127 & int32(7) {
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
		v168 = v116
		v169 = v118
		v170 = v122
		goto L38
	}
L22:
	;
	if v91 == int32(0) {
		v116 = v67
		v117 = v3
		v118 = v68
		goto L21
	} else {
		goto L28
	}
L23:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v71+int32(-17))))
	v91 = v90
	goto L22
L24:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v71+int32(-9))))
	v91 = v87
	goto L22
L25:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71+int32(-5)))))
	v91 = v84
	goto L22
L26:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+int32(-3)))))
	v91 = v81
	goto L22
L27:
	;
	v91 = int32(base.Ui32(v74) >> (uint(int32(3)) % 32))
	goto L22
L28:
	;
	v97 = v91&int32(-8) + int32(16)
	if v67 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v116 = v112
	v117 = int32(1)
	v118 = v97 + v68
	goto L21
L30:
	;
	v100 = F___bswap_16_1(m, int32(1))
	mBase = m.M
	goto L32
L31:
	;
	v112 = int32(0)
	goto L29
L32:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v67)+4)) = uint16(v100)
	v102 = F___bswap_32_1(m, v97)
	mBase = m.M
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v102
	if v91 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v110 = F___bswap_32_2(m, v102)
	mBase = m.M
	goto L37
L35:
	;
	goto L34
L36:
	;
	v108 = F__emscripten_memcpy_bulkmem(m, v67+int32(8), v71, v91)
	mBase = m.M
	goto L35
L37:
	;
	v112 = v67 + v110
	goto L29
L38:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v24)+2308))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+int32(-1)))))
	switch v177 & int32(7) {
	case 0:
		goto L62
	case 1:
		goto L61
	case 2:
		goto L60
	case 3:
		goto L59
	case 4:
		goto L58
	default:
		v219 = v168
		v220 = v169
		v223 = v122
		goto L56
	}
L39:
	;
	if v144 != 0 {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v124+int32(-17))))
	v144 = v143
	goto L39
L41:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v124+int32(-9))))
	v144 = v140
	goto L39
L42:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124+int32(-5)))))
	v144 = v137
	goto L39
L43:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+int32(-3)))))
	v144 = v134
	goto L39
L44:
	;
	v144 = int32(base.Ui32(v127) >> (uint(int32(3)) % 32))
	goto L39
L45:
	;
	v149 = v144&int32(-8) + int32(16)
	if v116 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v168 = v116
	v169 = v118
	v170 = int32(0)
	goto L38
L47:
	;
	v168 = v164
	v169 = v149 + v118
	v170 = int32(1)
	goto L38
L48:
	;
	v152 = F___bswap_16_1(m, int32(4))
	mBase = m.M
	goto L50
L49:
	;
	v164 = int32(0)
	goto L47
L50:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+4)) = uint16(v152)
	v154 = F___bswap_32_1(m, v149)
	mBase = m.M
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v154
	if v144 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v162 = F___bswap_32_2(m, v154)
	mBase = m.M
	goto L55
L53:
	;
	goto L52
L54:
	;
	v160 = F__emscripten_memcpy_bulkmem(m, v116+int32(8), v124, v144)
	mBase = m.M
	goto L53
L55:
	;
	v164 = v116 + v162
	goto L47
L56:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v24)+2336))
	if v226&int32(65535) != 0 {
		goto L74
	} else {
		goto L75
	}
L57:
	;
	if v194 == int32(0) {
		v219 = v168
		v220 = v169
		v223 = v122
		goto L56
	} else {
		goto L63
	}
L58:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v174+int32(-17))))
	v194 = v193
	goto L57
L59:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v174+int32(-9))))
	v194 = v190
	goto L57
L60:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v174+int32(-5)))))
	v194 = v187
	goto L57
L61:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+int32(-3)))))
	v194 = v184
	goto L57
L62:
	;
	v194 = int32(base.Ui32(v177) >> (uint(int32(3)) % 32))
	goto L57
L63:
	;
	v200 = v194&int32(-8) + int32(16)
	if v168 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v219 = v215
	v220 = v200 + v169
	v223 = int32(1)
	goto L56
L65:
	;
	v203 = F___bswap_16_1(m, int32(5))
	mBase = m.M
	goto L67
L66:
	;
	v215 = int32(0)
	goto L64
L67:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v168)+4)) = uint16(v203)
	v205 = F___bswap_32_1(m, v200)
	mBase = m.M
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v205
	if v194 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v213 = F___bswap_32_2(m, v205)
	mBase = m.M
	goto L72
L70:
	;
	goto L69
L71:
	;
	v211 = F__emscripten_memcpy_bulkmem(m, v168+int32(8), v174, v194)
	mBase = m.M
	goto L70
L72:
	;
	v215 = v168 + v213
	goto L64
L73:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v24)+2340))
	if v253&int32(65535) == int32(0) {
		v276 = v248
		v277 = v251
		v278 = int32(0)
		goto L83
	} else {
		goto L84
	}
L74:
	;
	if v219 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v248 = v219
	v249 = int32(0)
	v251 = v220
	goto L73
L76:
	;
	v248 = v243
	v249 = int32(1)
	v251 = v220 + int32(16)
	goto L73
L77:
	;
	v232 = F___bswap_16_1(m, int32(6))
	mBase = m.M
	goto L79
L78:
	;
	v243 = int32(0)
	goto L76
L79:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v219)+4)) = uint16(v232)
	v235 = F___bswap_32_1(m, int32(16))
	mBase = m.M
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v235
	v239 = F___bswap_16_1(m, v226&int32(65535))
	mBase = m.M
	goto L81
L81:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v219)+8)) = uint16(v239)
	v241 = F___bswap_32_2(m, v235)
	mBase = m.M
	goto L82
L82:
	;
	v243 = v219 + v241
	goto L76
L83:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v24)+2320))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281+int32(-1)))))
	switch v284 & int32(7) {
	case 0:
		goto L98
	case 1:
		goto L97
	case 2:
		goto L96
	case 3:
		goto L95
	case 4:
		goto L94
	default:
		v325 = v276
		v327 = v277
		v328 = int32(0)
		goto L92
	}
L84:
	;
	if v248 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v276 = v271
	v277 = v251 + int32(16)
	v278 = int32(1)
	goto L83
L86:
	;
	v260 = F___bswap_16_1(m, int32(7))
	mBase = m.M
	goto L88
L87:
	;
	v271 = int32(0)
	goto L85
L88:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v248)+4)) = uint16(v260)
	v263 = F___bswap_32_1(m, int32(16))
	mBase = m.M
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v263
	v267 = F___bswap_16_1(m, v253&int32(65535))
	mBase = m.M
	goto L90
L90:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v248)+8)) = uint16(v267)
	v269 = F___bswap_32_2(m, v263)
	mBase = m.M
	goto L91
L91:
	;
	v271 = v248 + v269
	goto L85
L92:
	;
	v330 = v117 + v69 + v170 + v223 + v249 + v278 + v328
	v332 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)+40))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v333)+16))
	if v334 == int32(0)-v336 {
		v555 = v325
		v558 = v327
		v562 = v330
		goto L110
	} else {
		goto L111
	}
L93:
	;
	if v301 != 0 {
		goto L99
	} else {
		goto L100
	}
L94:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v281+int32(-17))))
	v301 = v300
	goto L93
L95:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v281+int32(-9))))
	v301 = v297
	goto L93
L96:
	;
	v294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v281+int32(-5)))))
	v301 = v294
	goto L93
L97:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281+int32(-3)))))
	v301 = v291
	goto L93
L98:
	;
	v301 = int32(base.Ui32(v284) >> (uint(int32(3)) % 32))
	goto L93
L99:
	;
	v306 = v301&int32(-8) + int32(16)
	if v276 != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v325 = v276
	v327 = v277
	v328 = int32(0)
	goto L92
L101:
	;
	v325 = v321
	v327 = v306 + v277
	v328 = int32(1)
	goto L92
L102:
	;
	v309 = F___bswap_16_1(m, int32(8))
	mBase = m.M
	goto L104
L103:
	;
	v321 = int32(0)
	goto L101
L104:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v276)+4)) = uint16(v309)
	v311 = F___bswap_32_1(m, v306)
	mBase = m.M
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = v311
	if v301 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v319 = F___bswap_32_2(m, v311)
	mBase = m.M
	goto L109
L107:
	;
	goto L106
L108:
	;
	v317 = F__emscripten_memcpy_bulkmem(m, v276+int32(8), v281, v301)
	mBase = m.M
	goto L107
L109:
	;
	v321 = v276 + v319
	goto L101
L110:
	;
	if v555 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L111:
	;
	v339 = F_dictGetIterator(m, v333)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	return int32(0)
L113:
	;
	v344 = v325
	v347 = v327
	v351 = v330
	goto L115
L114:
	;
	F_dictReleaseIterator(m, v339)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L112
	} else {
		goto L153
	}
L115:
	;
	v362 = v339 + int32(20)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v339)+16))
	if v363 != 0 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	if v458 == int32(0) {
		goto L114
	} else {
		goto L143
	}
L118:
	;
	v369 = v362
	v370 = v366
	goto L121
L119:
	;
	v366 = int32(1)
	goto L118
L120:
	;
	v366 = int32(0)
	goto L118
L121:
	;
	switch v370 {
	case 0:
		goto L126
	default:
		goto L125
	}
L123:
	;
	v370 = int32(0)
	goto L121
L124:
	;
	goto L117
L125:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	*(*int32)(unsafe.Add(mBase, uint32(v339)+16)) = v450
	if v450 == int32(0) {
		goto L123
	} else {
		goto L142
	}
L126:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	if v374 != int32(-1) {
		v413 = v374
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v414 = int32(1)
	v415 = v413 + v414
	*(*int32)(unsafe.Add(mBase, uint32(v339)+4)) = v415
	v417 = int32(0)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v339)+8))
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420+v421+int32(26)))))
	if v425 == int32(255) {
		goto L136
	} else {
		goto L137
	}
L128:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v339)+8))
	if v378 != 0 {
		v413 = int32(-1)
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v339)+12))
	if v380 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)+20))
	if v407 != int32(-1) {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v387 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v379)+16)))
	v388 = int64(*(*int8)(unsafe.Add(mBase, uint32(v379)+27)))
	v389 = int64(*(*int32)(unsafe.Add(mBase, uint32(v379)+8)))
	v390 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v379)+12)))
	v391 = int64(*(*int8)(unsafe.Add(mBase, uint32(v379)+26)))
	v392 = int64(*(*int32)(unsafe.Add(mBase, uint32(v379)+4)))
	v393 = F_wangHash64(m, v392)
	mBase = m.M
	v395 = F_wangHash64(m, v391+v393)
	mBase = m.M
	v397 = F_wangHash64(m, v390+v395)
	mBase = m.M
	v399 = F_wangHash64(m, v389+v397)
	mBase = m.M
	v401 = F_wangHash64(m, v388+v399)
	mBase = m.M
	v403 = F_wangHash64(m, v387+v401)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v339)+24)) = v403
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v406 = v405
	goto L130
L132:
	;
	v383 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v379)+24)))
	v385 = v383 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v379)+24)) = uint16(v385)
	v406 = v379
	goto L130
L133:
	;
	v413 = v407 + int32(-1)
	goto L127
L134:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	v413 = v410
	goto L127
L135:
	;
	v440 = int32(2)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v420+v438<<(uint(v440)%32)+int32(4))))
	v369 = v445 + v439<<(uint(v440)%32)
	v370 = int32(1)
	goto L121
L136:
	;
	v429 = v417
	goto L138
L137:
	;
	v429 = v414 << (uint(v425) % 32)
	goto L138
L138:
	;
	if v415 < v429 {
		v438 = v421
		v439 = v415
		goto L135
	} else {
		goto L139
	}
L139:
	;
	if v421 != 0 {
		v458 = v417
		goto L124
	} else {
		goto L140
	}
L140:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v420)+20))
	if v431 == int32(-1) {
		v458 = v417
		goto L124
	} else {
		goto L141
	}
L141:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v339)+4)) = int64(4294967296)
	v438 = int32(1)
	v439 = int32(0)
	goto L135
L142:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v450)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = v454
	v458 = v450
	goto L124
L143:
	;
	if v344 != 0 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v344 = v544
	v347 = v347 + int32(56)
	v351 = v351 + int32(1)
	goto L115
L145:
	;
	v465 = *(*int64)(unsafe.Add(mBase, uint32(v458)+8))
	goto L147
L146:
	;
	v544 = int32(0)
	goto L144
L147:
	;
	v467 = *(*int64)(unsafe.Add(mBase, _consts[109]))
	if v465 < v467 {
		goto L115
	} else {
		goto L148
	}
L148:
	;
	v470 = *(*int64)(unsafe.Add(mBase, _consts[109]))
	v472 = F___bswap_16_1(m, int32(2))
	mBase = m.M
	goto L149
L149:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v344)+4)) = uint16(v472)
	v475 = F___bswap_32_1(m, int32(56))
	mBase = m.M
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = v475
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v458)))
	goto L151
L151:
	;
	v478 = *(*int64)(unsafe.Add(mBase, uint32(v477)))
	*(*int64)(unsafe.Add(mBase, uint32(v344)+8)) = v478
	v482 = int32(32)
	v484 = *(*int64)(unsafe.Add(mBase, uint32(v477+v482)))
	*(*int64)(unsafe.Add(mBase, uint32(v344+int32(40)))) = v484
	v488 = int32(24)
	v490 = *(*int64)(unsafe.Add(mBase, uint32(v477+v488)))
	*(*int64)(unsafe.Add(mBase, uint32(v344+v482))) = v490
	v494 = int32(16)
	v496 = *(*int64)(unsafe.Add(mBase, uint32(v477+v494)))
	*(*int64)(unsafe.Add(mBase, uint32(v344+v488))) = v496
	v502 = *(*int64)(unsafe.Add(mBase, uint32(v477+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v344+v494))) = v502
	v504 = v465 - v470
	v505 = int64(56)
	v507 = int64(65280)
	v509 = int64(40)
	v512 = int64(16711680)
	v514 = int64(24)
	v516 = int64(4278190080)
	v518 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v344)+48)) = v504<<(uint(v505)%64) | v504&v507<<(uint(v509)%64) | (v504&v512<<(uint(v514)%64) | v504&v516<<(uint(v518)%64)) | (int64(base.Ui64(v504)>>(uint(v518)%64))&v516 | int64(base.Ui64(v504)>>(uint(v514)%64))&v512 | (int64(base.Ui64(v504)>>(uint(v509)%64))&v507 | int64(base.Ui64(v504)>>(uint(v505)%64))))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v542 = F___bswap_32_2(m, v541)
	mBase = m.M
	goto L152
L152:
	;
	v544 = v344 + v542
	goto L144
L153:
	;
	v555 = v344
	v558 = v347
	v562 = v351
	goto L110
L154:
	;
	if l0 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L155:
	;
	v569 = F___bswap_16_1(m, int32(3))
	mBase = m.M
	goto L156
L156:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v555)+4)) = uint16(v569)
	v572 = F___bswap_32_1(m, int32(48))
	mBase = m.M
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v555))) = v572
	v575 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v576 = *(*int64)(unsafe.Add(mBase, uint32(v575)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v555)+8)) = v576
	v582 = *(*int64)(unsafe.Add(mBase, uint32(v575+int32(80))))
	*(*int64)(unsafe.Add(mBase, uint32(v555+int32(40)))) = v582
	v588 = *(*int64)(unsafe.Add(mBase, uint32(v575+int32(72))))
	*(*int64)(unsafe.Add(mBase, uint32(v555+int32(32)))) = v588
	v594 = *(*int64)(unsafe.Add(mBase, uint32(v575+int32(64))))
	*(*int64)(unsafe.Add(mBase, uint32(v555+int32(24)))) = v594
	v600 = *(*int64)(unsafe.Add(mBase, uint32(v575+int32(56))))
	*(*int64)(unsafe.Add(mBase, uint32(v555+int32(16)))) = v600
	v602 = F___bswap_32_2(m, v572)
	mBase = m.M
	goto L158
L158:
	;
	goto L154
L159:
	;
	return v558 + int32(48)
L160:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2253)))
	v609 = v607 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+2253)) = uint8(v609)
	v615 = F___bswap_16_1(m, (v562+int32(1))&int32(65535))
	mBase = m.M
	goto L161
L161:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+2214)) = uint16(v615)
	goto L159
}
func F_writePointerWithPadding(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	return
}
func F_writer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v8 int32
	_ = v8
	F_luaL_addlstring(m, l3, l1, l2)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
