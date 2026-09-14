package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F__writeToClient(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[430]))
	goto L3
L1:
	;
	if v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if base.B2i32(v7 == v2) == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v21 = l0 + int32(180)
	v22 = v15
	goto L1
L5:
	;
	v19 = int32(180)
	goto L7
L6:
	;
	v19 = int32(140)
	goto L7
L7:
	;
	v21 = l0 + v19
	v22 = v18
	goto L1
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v34 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v28 = F_writevToClient(m, l0)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
	if v23&int32(1) == int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	return int32(0)
L13:
	;
	return v28
L14:
	;
	F__serverAssert(m, int32(_a1008), int32(_a977), int32(2885))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L12
	} else {
		goto L33
	}
L15:
	;
	v40 = v33 - v34
	if int32(0) < v40 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v37 != v38 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v50 = int32(0)
	goto L22
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = int32(0)
	return int32(-1)
L20:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v83
	if v82 == v40 {
		goto L30
	} else {
		goto L31
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v50
	v82 = v50
	goto L20
L22:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+68))
	v61 = m.T0[v60].(func(*base.Module, int32, int32, int32) int32)(m, v53, v54+v55+v50, v40-v50)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L12
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v74
	if int32(1) <= v74 {
		v82 = v74
		goto L20
	} else {
		goto L29
	}
L24:
	;
	v74 = v61 + v50
	if v74 < v40 {
		v50 = v74
		goto L22
	} else {
		goto L28
	}
L25:
	;
	if int32(0) < v61 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+220)))
	v67 = v65 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+220)) = uint16(v67)
	if int32(0) < v50 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v61
	return int32(-1)
L28:
	;
	goto L23
L29:
	;
	return int32(-1)
L30:
	;
	v87 = v33
	goto L32
L31:
	;
	v87 = int32(0)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v89 + v82
	return int32(0)
L33:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sendToMainThread(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v4 = *(*int32)(unsafe.Add(mBase, _consts[487]))
	if v4 != 0 {
		v6 = l1 | l0
		F_flushPendingIOResponses(m, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _consts[487]))
			if v11 != 0 {
				v24 = v11
				v25 = v6
				v26 = F_listAddNodeTail(m, v24, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					return
				}
			} else {
				v13 = v6
				v16 = F_mpscEnqueue(m, int32(_a856), v13, int32(_a857))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					if v16 != 0 {
						return
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, _consts[487]))
						if v19 != 0 {
							v24 = v19
							v25 = v13
							v26 = F_listAddNodeTail(m, v24, v25)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								return
							}
						} else {
							v21 = F_listCreate(m)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[487])) = v21
								v24 = v21
								v25 = v13
								v26 = F_listAddNodeTail(m, v24, v25)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
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
		v13 = l1 | l0
		v16 = F_mpscEnqueue(m, int32(_a856), v13, int32(_a857))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 != 0 {
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _consts[487]))
				if v19 != 0 {
					v24 = v19
					v25 = v13
					v26 = F_listAddNodeTail(m, v24, v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						return
					}
				} else {
					v21 = F_listCreate(m)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[487])) = v21
						v24 = v21
						v25 = v13
						v26 = F_listAddNodeTail(m, v24, v25)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
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
func F_toClusterMsg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v3 = F___bswap_16_2(m, v2)
	mBase = m.M
	if int32(-1) < base.I32_extend16_s(v3) {
		return l0
	} else {
		F__serverAssert(m, int32(_a287), int32(_a253), int32(145))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_toClusterMsgLight(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v3 = F___bswap_16_2(m, v2)
	mBase = m.M
	if base.I32_extend16_s(v3) < int32(0) {
		return l0
	} else {
		F__serverAssert(m, int32(_a300), int32(_a253), int32(151))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_trySendReadToIOThreads(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	v7 = *(*int32)(unsafe.Add(mBase, _consts[431]))
	if int32(2) <= v7 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v12 != 0 {
			v15 = int32(0)
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
			if v16 == int32(1) {
				v221 = v15
				return v221
			} else {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
				if v19 != int32(2) {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
					if v24 == int32(1) {
						v221 = v15
						return v221
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						if v27&int32(1) != 0 {
							v47 = v27
							if v47&int32(268436624) == int32(0) {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								if v55 == int32(0) {
									v70 = int32(0)
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
									if v72 != 0 {
										v88 = v70
									} else {
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
										if v73&int32(144) != 0 {
											v88 = v70
										} else {
											v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
											if v76&int32(2) != 0 {
												v88 = v70
											} else {
												v79 = F_isInsideYieldingLongCommand(m)
												mBase = m.M
												if v79 == int32(0) {
													v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
													v88 = base.B2i32(v83&int32(1088) == int32(0))
												} else {
													v82 = F_isReplicatedClient(m, l0)
													mBase = m.M
													if v82 != 0 {
														v88 = v70
													} else {
														v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
														v88 = base.B2i32(v83&int32(1088) == int32(0))
													}
												}
											}
										}
									}
									v89 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v88 == v89) << (uint(int32(15)) % 32)
									v97 = *(*int32)(unsafe.Add(mBase, _consts[4]))
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
									if v98&int32(6) == int32(4) {
										v110 = v89
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
										v110 = int32(base.Ui32(v103^int32(-1))>>(uint(int32(23))%32)) & int32(1)
									}
									v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v110 != int32(0))<<(uint(int32(16))%32) | v115
									v119 = int32(1)
									v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
									if v120&v119 != 0 {
										v127 = v119
										v130 = v127
									} else {
										v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
										if v123 != 0 {
											v125 = F_isImportSlotMigrationJob(m, v123)
											mBase = m.M
											v127 = v125
											v130 = v127
										} else {
											v130 = int32(0)
										}
									}
									v131 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v131)
									v133 = int32(0)
									v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v130 != v133)<<(uint(int32(14))%32) | v137
									v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
									v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
									v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v142 == v133 {
										v162 = int32(_a847)
										v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
										v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
										v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
										v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
										v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
										if v175 != v170 {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
											v178 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
											v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
											*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
										}
										if v175 == v170 {
											v204 = int32(0)
											v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
											*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
											v211 = int32(_a44)
											v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
											*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
											v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
											v221 = v204
											return v221
										} else {
											v186 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
											v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v190 != 0 {
												v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
												if v193 != 0 {
													v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
													if v196 != 0 {
														m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
														mBase = m.M
														v201 = m.ExcPending
														if v201 != 0 {
															return int32(0)
														} else {
															return int32(-1)
														}
													} else {
														return int32(-1)
													}
												} else {
													return int32(-1)
												}
											} else {
												return int32(-1)
											}
										}
									} else {
										v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
										if v145 == int32(0) {
											v162 = int32(_a847)
											v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
											v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
											v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
											v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
											v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
											if v175 != v170 {
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
												v178 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
												v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
												*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
											}
											if v175 == v170 {
												v204 = int32(0)
												v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
												*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
												v211 = int32(_a44)
												v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
												*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
												v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
												v221 = v204
												return v221
											} else {
												v186 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
												v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v190 != 0 {
													v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
													if v193 != 0 {
														v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
														if v196 != 0 {
															m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
															mBase = m.M
															v201 = m.ExcPending
															if v201 != 0 {
																return int32(0)
															} else {
																return int32(-1)
															}
														} else {
															return int32(-1)
														}
													} else {
														return int32(-1)
													}
												} else {
													return int32(-1)
												}
											}
										} else {
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)+112))
											if v148 == int32(0) {
												v162 = int32(_a847)
												v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
												v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
												v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
												v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
												v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
												if v175 != v170 {
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
													v178 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
													v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
													*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
												}
												if v175 == v170 {
													v204 = int32(0)
													v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
													*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
													v211 = int32(_a44)
													v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
													*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
													v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
													v221 = v204
													return v221
												} else {
													v186 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
													v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													if v190 != 0 {
														v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
														if v193 != 0 {
															v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
															if v196 != 0 {
																m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																mBase = m.M
																v201 = m.ExcPending
																if v201 != 0 {
																	return int32(0)
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														} else {
															return int32(-1)
														}
													} else {
														return int32(-1)
													}
												}
											} else {
												v152 = base.B2i32(v140 != int32(0))
												if v141 != 0 {
													v155 = v152 | int32(2)
												} else {
													v155 = v152
												}
												m.T0[v148].(func(*base.Module, int32, int32))(m, v142, v155)
												mBase = m.M
												v159 = m.ExcPending
												if v159 != 0 {
													return int32(0)
												} else {
													v162 = int32(_a847)
													v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
													v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
													v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
													v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
													v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
													if v175 != v170 {
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
														v178 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
														v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
														*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
													}
													if v175 == v170 {
														v204 = int32(0)
														v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
														*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
														v211 = int32(_a44)
														v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
														*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
														v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
														v221 = v204
														return v221
													} else {
														v186 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
														v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														if v190 != 0 {
															v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
															if v193 != 0 {
																v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																if v196 != 0 {
																	m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int32(0)
																	} else {
																		return int32(-1)
																	}
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														} else {
															return int32(-1)
														}
													}
												}
											}
										}
									}
								} else {
									v59 = int32(1)
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
									if v60 == int32(0) {
										v67 = v59
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
										if v63 != 0 {
											v67 = v59
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+156))
											v67 = base.B2i32(v64 != int32(13))
										}
									}
									if v67 != 0 {
										v70 = int32(0)
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										if v72 != 0 {
											v88 = v70
										} else {
											v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
											if v73&int32(144) != 0 {
												v88 = v70
											} else {
												v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
												if v76&int32(2) != 0 {
													v88 = v70
												} else {
													v79 = F_isInsideYieldingLongCommand(m)
													mBase = m.M
													if v79 == int32(0) {
														v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
														v88 = base.B2i32(v83&int32(1088) == int32(0))
													} else {
														v82 = F_isReplicatedClient(m, l0)
														mBase = m.M
														if v82 != 0 {
															v88 = v70
														} else {
															v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
															v88 = base.B2i32(v83&int32(1088) == int32(0))
														}
													}
												}
											}
										}
										v89 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v88 == v89) << (uint(int32(15)) % 32)
										v97 = *(*int32)(unsafe.Add(mBase, _consts[4]))
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
										if v98&int32(6) == int32(4) {
											v110 = v89
										} else {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
											v110 = int32(base.Ui32(v103^int32(-1))>>(uint(int32(23))%32)) & int32(1)
										}
										v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v110 != int32(0))<<(uint(int32(16))%32) | v115
										v119 = int32(1)
										v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
										if v120&v119 != 0 {
											v127 = v119
											v130 = v127
										} else {
											v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
											if v123 != 0 {
												v125 = F_isImportSlotMigrationJob(m, v123)
												mBase = m.M
												v127 = v125
												v130 = v127
											} else {
												v130 = int32(0)
											}
										}
										v131 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v131)
										v133 = int32(0)
										v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v130 != v133)<<(uint(int32(14))%32) | v137
										v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
										v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
										v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v142 == v133 {
											v162 = int32(_a847)
											v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
											v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
											v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
											v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
											v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
											if v175 != v170 {
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
												v178 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
												v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
												*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
											}
											if v175 == v170 {
												v204 = int32(0)
												v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
												*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
												v211 = int32(_a44)
												v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
												*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
												v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
												v221 = v204
												return v221
											} else {
												v186 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
												v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v190 != 0 {
													v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
													if v193 != 0 {
														v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
														if v196 != 0 {
															m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
															mBase = m.M
															v201 = m.ExcPending
															if v201 != 0 {
																return int32(0)
															} else {
																return int32(-1)
															}
														} else {
															return int32(-1)
														}
													} else {
														return int32(-1)
													}
												} else {
													return int32(-1)
												}
											}
										} else {
											v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
											if v145 == int32(0) {
												v162 = int32(_a847)
												v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
												v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
												v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
												v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
												v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
												if v175 != v170 {
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
													v178 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
													v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
													*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
												}
												if v175 == v170 {
													v204 = int32(0)
													v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
													*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
													v211 = int32(_a44)
													v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
													*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
													v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
													v221 = v204
													return v221
												} else {
													v186 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
													v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													if v190 != 0 {
														v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
														if v193 != 0 {
															v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
															if v196 != 0 {
																m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																mBase = m.M
																v201 = m.ExcPending
																if v201 != 0 {
																	return int32(0)
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														} else {
															return int32(-1)
														}
													} else {
														return int32(-1)
													}
												}
											} else {
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)+112))
												if v148 == int32(0) {
													v162 = int32(_a847)
													v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
													v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
													v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
													v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
													v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
													if v175 != v170 {
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
														v178 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
														v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
														*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
													}
													if v175 == v170 {
														v204 = int32(0)
														v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
														*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
														v211 = int32(_a44)
														v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
														*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
														v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
														v221 = v204
														return v221
													} else {
														v186 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
														v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														if v190 != 0 {
															v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
															if v193 != 0 {
																v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																if v196 != 0 {
																	m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int32(0)
																	} else {
																		return int32(-1)
																	}
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														} else {
															return int32(-1)
														}
													}
												} else {
													v152 = base.B2i32(v140 != int32(0))
													if v141 != 0 {
														v155 = v152 | int32(2)
													} else {
														v155 = v152
													}
													m.T0[v148].(func(*base.Module, int32, int32))(m, v142, v155)
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
														v162 = int32(_a847)
														v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
														v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
														v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
														v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
														v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
														if v175 != v170 {
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
															v178 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
															v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
															*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
														}
														if v175 == v170 {
															v204 = int32(0)
															v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
															*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
															v211 = int32(_a44)
															v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
															*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
															v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
															v221 = v204
															return v221
														} else {
															v186 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
															*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
															v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															if v190 != 0 {
																v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
																if v193 != 0 {
																	v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																	if v196 != 0 {
																		m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return int32(0)
																		} else {
																			return int32(-1)
																		}
																	} else {
																		return int32(-1)
																	}
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														}
													}
												}
											}
										}
									} else {
										return int32(-1)
									}
								}
							} else {
								return int32(-1)
							}
						} else {
							if v27&int32(2) == int32(0) {
								if v27&int32(262144) != 0 {
									v47 = v27
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
									if v40 == int32(0) {
										v47 = v27
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
										v47 = v46
									}
								}
								if v47&int32(268436624) == int32(0) {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
									if v55 == int32(0) {
										v70 = int32(0)
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										if v72 != 0 {
											v88 = v70
										} else {
											v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
											if v73&int32(144) != 0 {
												v88 = v70
											} else {
												v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
												if v76&int32(2) != 0 {
													v88 = v70
												} else {
													v79 = F_isInsideYieldingLongCommand(m)
													mBase = m.M
													if v79 == int32(0) {
														v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
														v88 = base.B2i32(v83&int32(1088) == int32(0))
													} else {
														v82 = F_isReplicatedClient(m, l0)
														mBase = m.M
														if v82 != 0 {
															v88 = v70
														} else {
															v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
															v88 = base.B2i32(v83&int32(1088) == int32(0))
														}
													}
												}
											}
										}
										v89 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v88 == v89) << (uint(int32(15)) % 32)
										v97 = *(*int32)(unsafe.Add(mBase, _consts[4]))
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
										if v98&int32(6) == int32(4) {
											v110 = v89
										} else {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
											v110 = int32(base.Ui32(v103^int32(-1))>>(uint(int32(23))%32)) & int32(1)
										}
										v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v110 != int32(0))<<(uint(int32(16))%32) | v115
										v119 = int32(1)
										v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
										if v120&v119 != 0 {
											v127 = v119
											v130 = v127
										} else {
											v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
											if v123 != 0 {
												v125 = F_isImportSlotMigrationJob(m, v123)
												mBase = m.M
												v127 = v125
												v130 = v127
											} else {
												v130 = int32(0)
											}
										}
										v131 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v131)
										v133 = int32(0)
										v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v130 != v133)<<(uint(int32(14))%32) | v137
										v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
										v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
										v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v142 == v133 {
											v162 = int32(_a847)
											v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
											v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
											v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
											v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
											v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
											if v175 != v170 {
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
												v178 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
												v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
												*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
											}
											if v175 == v170 {
												v204 = int32(0)
												v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
												*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
												v211 = int32(_a44)
												v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
												*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
												v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
												v221 = v204
												return v221
											} else {
												v186 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
												v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v190 != 0 {
													v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
													if v193 != 0 {
														v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
														if v196 != 0 {
															m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
															mBase = m.M
															v201 = m.ExcPending
															if v201 != 0 {
																return int32(0)
															} else {
																return int32(-1)
															}
														} else {
															return int32(-1)
														}
													} else {
														return int32(-1)
													}
												} else {
													return int32(-1)
												}
											}
										} else {
											v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
											if v145 == int32(0) {
												v162 = int32(_a847)
												v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
												v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
												v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
												v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
												v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
												if v175 != v170 {
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
													v178 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
													v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
													*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
												}
												if v175 == v170 {
													v204 = int32(0)
													v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
													*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
													v211 = int32(_a44)
													v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
													*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
													v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
													v221 = v204
													return v221
												} else {
													v186 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
													v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													if v190 != 0 {
														v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
														if v193 != 0 {
															v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
															if v196 != 0 {
																m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																mBase = m.M
																v201 = m.ExcPending
																if v201 != 0 {
																	return int32(0)
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														} else {
															return int32(-1)
														}
													} else {
														return int32(-1)
													}
												}
											} else {
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)+112))
												if v148 == int32(0) {
													v162 = int32(_a847)
													v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
													v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
													v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
													v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
													v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
													if v175 != v170 {
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
														v178 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
														v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
														*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
													}
													if v175 == v170 {
														v204 = int32(0)
														v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
														*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
														v211 = int32(_a44)
														v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
														*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
														v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
														v221 = v204
														return v221
													} else {
														v186 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
														v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														if v190 != 0 {
															v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
															if v193 != 0 {
																v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																if v196 != 0 {
																	m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int32(0)
																	} else {
																		return int32(-1)
																	}
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														} else {
															return int32(-1)
														}
													}
												} else {
													v152 = base.B2i32(v140 != int32(0))
													if v141 != 0 {
														v155 = v152 | int32(2)
													} else {
														v155 = v152
													}
													m.T0[v148].(func(*base.Module, int32, int32))(m, v142, v155)
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
														v162 = int32(_a847)
														v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
														v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
														v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
														v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
														v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
														if v175 != v170 {
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
															v178 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
															v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
															*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
														}
														if v175 == v170 {
															v204 = int32(0)
															v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
															*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
															v211 = int32(_a44)
															v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
															*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
															v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
															v221 = v204
															return v221
														} else {
															v186 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
															*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
															v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															if v190 != 0 {
																v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
																if v193 != 0 {
																	v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																	if v196 != 0 {
																		m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return int32(0)
																		} else {
																			return int32(-1)
																		}
																	} else {
																		return int32(-1)
																	}
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														}
													}
												}
											}
										}
									} else {
										v59 = int32(1)
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
										if v60 == int32(0) {
											v67 = v59
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
											if v63 != 0 {
												v67 = v59
											} else {
												v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+156))
												v67 = base.B2i32(v64 != int32(13))
											}
										}
										if v67 != 0 {
											v70 = int32(0)
											v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											if v72 != 0 {
												v88 = v70
											} else {
												v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
												if v73&int32(144) != 0 {
													v88 = v70
												} else {
													v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
													if v76&int32(2) != 0 {
														v88 = v70
													} else {
														v79 = F_isInsideYieldingLongCommand(m)
														mBase = m.M
														if v79 == int32(0) {
															v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
															v88 = base.B2i32(v83&int32(1088) == int32(0))
														} else {
															v82 = F_isReplicatedClient(m, l0)
															mBase = m.M
															if v82 != 0 {
																v88 = v70
															} else {
																v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
																v88 = base.B2i32(v83&int32(1088) == int32(0))
															}
														}
													}
												}
											}
											v89 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v88 == v89) << (uint(int32(15)) % 32)
											v97 = *(*int32)(unsafe.Add(mBase, _consts[4]))
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
											if v98&int32(6) == int32(4) {
												v110 = v89
											} else {
												v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
												v110 = int32(base.Ui32(v103^int32(-1))>>(uint(int32(23))%32)) & int32(1)
											}
											v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v110 != int32(0))<<(uint(int32(16))%32) | v115
											v119 = int32(1)
											v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
											if v120&v119 != 0 {
												v127 = v119
												v130 = v127
											} else {
												v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
												if v123 != 0 {
													v125 = F_isImportSlotMigrationJob(m, v123)
													mBase = m.M
													v127 = v125
													v130 = v127
												} else {
													v130 = int32(0)
												}
											}
											v131 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v131)
											v133 = int32(0)
											v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v130 != v133)<<(uint(int32(14))%32) | v137
											v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
											v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
											v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v142 == v133 {
												v162 = int32(_a847)
												v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
												v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
												v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
												v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
												v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
												if v175 != v170 {
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
													v178 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
													v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
													*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
												}
												if v175 == v170 {
													v204 = int32(0)
													v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
													*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
													v211 = int32(_a44)
													v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
													*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
													v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
													v221 = v204
													return v221
												} else {
													v186 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
													v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													if v190 != 0 {
														v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
														if v193 != 0 {
															v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
															if v196 != 0 {
																m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																mBase = m.M
																v201 = m.ExcPending
																if v201 != 0 {
																	return int32(0)
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														} else {
															return int32(-1)
														}
													} else {
														return int32(-1)
													}
												}
											} else {
												v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
												if v145 == int32(0) {
													v162 = int32(_a847)
													v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
													v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
													v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
													v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
													v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
													if v175 != v170 {
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
														v178 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
														v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
														*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
													}
													if v175 == v170 {
														v204 = int32(0)
														v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
														*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
														v211 = int32(_a44)
														v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
														*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
														v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
														v221 = v204
														return v221
													} else {
														v186 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
														v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														if v190 != 0 {
															v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
															if v193 != 0 {
																v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																if v196 != 0 {
																	m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int32(0)
																	} else {
																		return int32(-1)
																	}
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														} else {
															return int32(-1)
														}
													}
												} else {
													v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)+112))
													if v148 == int32(0) {
														v162 = int32(_a847)
														v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
														v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
														v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
														v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
														v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
														if v175 != v170 {
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
															v178 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
															v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
															*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
														}
														if v175 == v170 {
															v204 = int32(0)
															v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
															*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
															v211 = int32(_a44)
															v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
															*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
															v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
															v221 = v204
															return v221
														} else {
															v186 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
															*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
															v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															if v190 != 0 {
																v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
																if v193 != 0 {
																	v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																	if v196 != 0 {
																		m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return int32(0)
																		} else {
																			return int32(-1)
																		}
																	} else {
																		return int32(-1)
																	}
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														}
													} else {
														v152 = base.B2i32(v140 != int32(0))
														if v141 != 0 {
															v155 = v152 | int32(2)
														} else {
															v155 = v152
														}
														m.T0[v148].(func(*base.Module, int32, int32))(m, v142, v155)
														mBase = m.M
														v159 = m.ExcPending
														if v159 != 0 {
															return int32(0)
														} else {
															v162 = int32(_a847)
															v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
															v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
															v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
															v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
															v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
															if v175 != v170 {
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
																v178 = int32(1)
																*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
																v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
																*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
															}
															if v175 == v170 {
																v204 = int32(0)
																v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
																*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
																v211 = int32(_a44)
																v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
																*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
																v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
																v221 = v204
																return v221
															} else {
																v186 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
																*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
																v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																if v190 != 0 {
																	v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
																	if v193 != 0 {
																		v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																		if v196 != 0 {
																			m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																			mBase = m.M
																			v201 = m.ExcPending
																			if v201 != 0 {
																				return int32(0)
																			} else {
																				return int32(-1)
																			}
																		} else {
																			return int32(-1)
																		}
																	} else {
																		return int32(-1)
																	}
																} else {
																	return int32(-1)
																}
															}
														}
													}
												}
											}
										} else {
											return int32(-1)
										}
									}
								} else {
									return int32(-1)
								}
							} else {
								if v27&int32(4) != 0 {
									if v27&int32(262144) != 0 {
										v47 = v27
									} else {
										v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
										if v40 == int32(0) {
											v47 = v27
										} else {
											v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
											v47 = v46
										}
									}
									if v47&int32(268436624) == int32(0) {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
										if v55 == int32(0) {
											v70 = int32(0)
											v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											if v72 != 0 {
												v88 = v70
											} else {
												v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
												if v73&int32(144) != 0 {
													v88 = v70
												} else {
													v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
													if v76&int32(2) != 0 {
														v88 = v70
													} else {
														v79 = F_isInsideYieldingLongCommand(m)
														mBase = m.M
														if v79 == int32(0) {
															v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
															v88 = base.B2i32(v83&int32(1088) == int32(0))
														} else {
															v82 = F_isReplicatedClient(m, l0)
															mBase = m.M
															if v82 != 0 {
																v88 = v70
															} else {
																v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
																v88 = base.B2i32(v83&int32(1088) == int32(0))
															}
														}
													}
												}
											}
											v89 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v88 == v89) << (uint(int32(15)) % 32)
											v97 = *(*int32)(unsafe.Add(mBase, _consts[4]))
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
											if v98&int32(6) == int32(4) {
												v110 = v89
											} else {
												v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
												v110 = int32(base.Ui32(v103^int32(-1))>>(uint(int32(23))%32)) & int32(1)
											}
											v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v110 != int32(0))<<(uint(int32(16))%32) | v115
											v119 = int32(1)
											v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
											if v120&v119 != 0 {
												v127 = v119
												v130 = v127
											} else {
												v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
												if v123 != 0 {
													v125 = F_isImportSlotMigrationJob(m, v123)
													mBase = m.M
													v127 = v125
													v130 = v127
												} else {
													v130 = int32(0)
												}
											}
											v131 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v131)
											v133 = int32(0)
											v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v130 != v133)<<(uint(int32(14))%32) | v137
											v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
											v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
											v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v142 == v133 {
												v162 = int32(_a847)
												v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
												v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
												v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
												v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
												v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
												if v175 != v170 {
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
													v178 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
													v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
													*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
												}
												if v175 == v170 {
													v204 = int32(0)
													v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
													*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
													v211 = int32(_a44)
													v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
													*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
													v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
													v221 = v204
													return v221
												} else {
													v186 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
													v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													if v190 != 0 {
														v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
														if v193 != 0 {
															v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
															if v196 != 0 {
																m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																mBase = m.M
																v201 = m.ExcPending
																if v201 != 0 {
																	return int32(0)
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														} else {
															return int32(-1)
														}
													} else {
														return int32(-1)
													}
												}
											} else {
												v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
												if v145 == int32(0) {
													v162 = int32(_a847)
													v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
													v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
													v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
													v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
													v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
													if v175 != v170 {
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
														v178 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
														v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
														*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
													}
													if v175 == v170 {
														v204 = int32(0)
														v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
														*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
														v211 = int32(_a44)
														v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
														*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
														v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
														v221 = v204
														return v221
													} else {
														v186 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
														v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														if v190 != 0 {
															v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
															if v193 != 0 {
																v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																if v196 != 0 {
																	m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int32(0)
																	} else {
																		return int32(-1)
																	}
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														} else {
															return int32(-1)
														}
													}
												} else {
													v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)+112))
													if v148 == int32(0) {
														v162 = int32(_a847)
														v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
														v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
														v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
														v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
														v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
														if v175 != v170 {
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
															v178 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
															v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
															*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
														}
														if v175 == v170 {
															v204 = int32(0)
															v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
															*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
															v211 = int32(_a44)
															v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
															*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
															v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
															v221 = v204
															return v221
														} else {
															v186 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
															*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
															v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															if v190 != 0 {
																v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
																if v193 != 0 {
																	v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																	if v196 != 0 {
																		m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return int32(0)
																		} else {
																			return int32(-1)
																		}
																	} else {
																		return int32(-1)
																	}
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														}
													} else {
														v152 = base.B2i32(v140 != int32(0))
														if v141 != 0 {
															v155 = v152 | int32(2)
														} else {
															v155 = v152
														}
														m.T0[v148].(func(*base.Module, int32, int32))(m, v142, v155)
														mBase = m.M
														v159 = m.ExcPending
														if v159 != 0 {
															return int32(0)
														} else {
															v162 = int32(_a847)
															v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
															v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
															v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
															v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
															v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
															if v175 != v170 {
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
																v178 = int32(1)
																*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
																v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
																*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
															}
															if v175 == v170 {
																v204 = int32(0)
																v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
																*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
																v211 = int32(_a44)
																v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
																*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
																v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
																v221 = v204
																return v221
															} else {
																v186 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
																*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
																v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																if v190 != 0 {
																	v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
																	if v193 != 0 {
																		v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																		if v196 != 0 {
																			m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																			mBase = m.M
																			v201 = m.ExcPending
																			if v201 != 0 {
																				return int32(0)
																			} else {
																				return int32(-1)
																			}
																		} else {
																			return int32(-1)
																		}
																	} else {
																		return int32(-1)
																	}
																} else {
																	return int32(-1)
																}
															}
														}
													}
												}
											}
										} else {
											v59 = int32(1)
											v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
											if v60 == int32(0) {
												v67 = v59
											} else {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
												if v63 != 0 {
													v67 = v59
												} else {
													v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+156))
													v67 = base.B2i32(v64 != int32(13))
												}
											}
											if v67 != 0 {
												v70 = int32(0)
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												if v72 != 0 {
													v88 = v70
												} else {
													v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
													if v73&int32(144) != 0 {
														v88 = v70
													} else {
														v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
														if v76&int32(2) != 0 {
															v88 = v70
														} else {
															v79 = F_isInsideYieldingLongCommand(m)
															mBase = m.M
															if v79 == int32(0) {
																v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
																v88 = base.B2i32(v83&int32(1088) == int32(0))
															} else {
																v82 = F_isReplicatedClient(m, l0)
																mBase = m.M
																if v82 != 0 {
																	v88 = v70
																} else {
																	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
																	v88 = base.B2i32(v83&int32(1088) == int32(0))
																}
															}
														}
													}
												}
												v89 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v88 == v89) << (uint(int32(15)) % 32)
												v97 = *(*int32)(unsafe.Add(mBase, _consts[4]))
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
												if v98&int32(6) == int32(4) {
													v110 = v89
												} else {
													v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
													v110 = int32(base.Ui32(v103^int32(-1))>>(uint(int32(23))%32)) & int32(1)
												}
												v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v110 != int32(0))<<(uint(int32(16))%32) | v115
												v119 = int32(1)
												v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
												if v120&v119 != 0 {
													v127 = v119
													v130 = v127
												} else {
													v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
													if v123 != 0 {
														v125 = F_isImportSlotMigrationJob(m, v123)
														mBase = m.M
														v127 = v125
														v130 = v127
													} else {
														v130 = int32(0)
													}
												}
												v131 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v131)
												v133 = int32(0)
												v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = base.B2i32(v130 != v133)<<(uint(int32(14))%32) | v137
												v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
												v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
												v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v142 == v133 {
													v162 = int32(_a847)
													v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
													v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
													v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
													v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
													v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
													if v175 != v170 {
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
														v178 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
														v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
														*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
													}
													if v175 == v170 {
														v204 = int32(0)
														v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
														*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
														v211 = int32(_a44)
														v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
														*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
														v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
														v221 = v204
														return v221
													} else {
														v186 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
														v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														if v190 != 0 {
															v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
															if v193 != 0 {
																v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																if v196 != 0 {
																	m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int32(0)
																	} else {
																		return int32(-1)
																	}
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														} else {
															return int32(-1)
														}
													}
												} else {
													v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
													if v145 == int32(0) {
														v162 = int32(_a847)
														v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
														v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
														v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
														v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
														v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
														if v175 != v170 {
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
															v178 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
															v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
															*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
														}
														if v175 == v170 {
															v204 = int32(0)
															v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
															*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
															v211 = int32(_a44)
															v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
															*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
															v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
															v221 = v204
															return v221
														} else {
															v186 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
															*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
															v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															if v190 != 0 {
																v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
																if v193 != 0 {
																	v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																	if v196 != 0 {
																		m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																		mBase = m.M
																		v201 = m.ExcPending
																		if v201 != 0 {
																			return int32(0)
																		} else {
																			return int32(-1)
																		}
																	} else {
																		return int32(-1)
																	}
																} else {
																	return int32(-1)
																}
															} else {
																return int32(-1)
															}
														}
													} else {
														v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)+112))
														if v148 == int32(0) {
															v162 = int32(_a847)
															v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
															v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
															v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
															v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
															v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
															if v175 != v170 {
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
																v178 = int32(1)
																*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
																v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
																*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
															}
															if v175 == v170 {
																v204 = int32(0)
																v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
																*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
																v211 = int32(_a44)
																v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
																*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
																v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
																v221 = v204
																return v221
															} else {
																v186 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
																*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
																v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																if v190 != 0 {
																	v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
																	if v193 != 0 {
																		v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																		if v196 != 0 {
																			m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																			mBase = m.M
																			v201 = m.ExcPending
																			if v201 != 0 {
																				return int32(0)
																			} else {
																				return int32(-1)
																			}
																		} else {
																			return int32(-1)
																		}
																	} else {
																		return int32(-1)
																	}
																} else {
																	return int32(-1)
																}
															}
														} else {
															v152 = base.B2i32(v140 != int32(0))
															if v141 != 0 {
																v155 = v152 | int32(2)
															} else {
																v155 = v152
															}
															m.T0[v148].(func(*base.Module, int32, int32))(m, v142, v155)
															mBase = m.M
															v159 = m.ExcPending
															if v159 != 0 {
																return int32(0)
															} else {
																v162 = int32(_a847)
																v166 = *(*int32)(unsafe.Add(mBase, _consts[484]))
																v167 = *(*int32)(unsafe.Add(mBase, _consts[485]))
																v170 = *(*int32)(unsafe.Add(mBase, _consts[439]))
																v174 = v166 + (v167+int32(-1))&v170<<(uint(int32(6))%32)
																v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
																if v175 != v170 {
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = l0
																	v178 = int32(1)
																	*(*int32)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
																	v181 = *(*int32)(unsafe.Add(mBase, _consts[439]))
																	*(*int32)(unsafe.Add(mBase, _consts[439])) = v181 + v178
																}
																if v175 == v170 {
																	v204 = int32(0)
																	v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
																	*(*int32)(unsafe.Add(mBase, _consts[434])) = v207 + int32(1)
																	v211 = int32(_a44)
																	v213 = *(*int64)(unsafe.Add(mBase, _consts[486]))
																	*(*int64)(unsafe.Add(mBase, _consts[486])) = v213 + int64(1)
																	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v217 | int32(_a0)
																	v221 = v204
																	return v221
																} else {
																	v186 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v186)
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v186
																	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	if v190 != 0 {
																		v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
																		if v193 != 0 {
																			v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+112))
																			if v196 != 0 {
																				m.T0[v196].(func(*base.Module, int32, int32))(m, v190, int32(0))
																				mBase = m.M
																				v201 = m.ExcPending
																				if v201 != 0 {
																					return int32(0)
																				} else {
																					return int32(-1)
																				}
																			} else {
																				return int32(-1)
																			}
																		} else {
																			return int32(-1)
																		}
																	} else {
																		return int32(-1)
																	}
																}
															}
														}
													}
												}
											} else {
												return int32(-1)
											}
										}
									} else {
										return int32(-1)
									}
								} else {
									return int32(-1)
								}
							}
						}
					}
				} else {
					return int32(-1)
				}
			}
		} else {
			return int32(-1)
		}
	} else {
		return int32(-1)
	}
}
