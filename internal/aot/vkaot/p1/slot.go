package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_checkSlotExportOwnership(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v227
L2:
	;
	v213 = int32(-1)
	v215 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v215 {
		v227 = v213
		goto L1
	} else {
		goto L52
	}
L3:
	;
	F__serverAssert(m, int32(_a402), int32(_a386), int32(1793))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L49
	} else {
		goto L51
	}
L4:
	;
	v14 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v14)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if base.Ui32(int32(20)) < base.Ui32(v17) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v26 = v11 + int32(24)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v27
	goto L8
L6:
	;
	if int32(1)<<(uint(v17)%32)&int32(1835040) != 0 {
		v227 = v14
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v32 = v11 + int32(24)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v34 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v34 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v34+base.B2i32(v37 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v43
	goto L10
L12:
	;
	v51 = v34
	v52 = int32(0)
	goto L13
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v58 < v57 {
		v89 = v52
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v89 == int32(0) {
		goto L2
	} else {
		goto L29
	}
L15:
	;
	v94 = v11 + int32(24)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v96 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v67 = v57
	v68 = v52
	goto L17
L17:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(52)+v67<<(uint(int32(2))%32))))
	if v75 == int32(0) {
		goto L2
	} else {
		goto L19
	}
L18:
	;
	v89 = v78
	goto L15
L19:
	;
	if v68 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v78 = v68
	goto L22
L21:
	;
	v78 = v75
	goto L22
L22:
	;
	if v78 != v75 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	if base.B2i32(v67 == v58) == int32(0) {
		v67 = v67 + int32(1)
		v68 = v78
		goto L17
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	if v96 != 0 {
		v51 = v96
		v52 = v89
		goto L13
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v96+base.B2i32(v99 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v105
	goto L26
L28:
	;
	goto L14
L29:
	;
	v109 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v89 == v112 {
		v227 = v109
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v115 = v89 + int32(8)
	v117 = l0 + int32(32)
	v118 = int32(40)
	goto L35
L31:
	;
	if v182 != 0 {
		goto L2
	} else {
		goto L47
	}
L32:
	;
	v182 = int32(0)
	goto L31
L33:
	;
	v154 = v149
	v155 = v150
	v156 = v151
	goto L43
L34:
	;
	if v139 == int32(0) {
		goto L32
	} else {
		goto L41
	}
L35:
	;
	if (v117|v115)&int32(3) != 0 {
		v149 = v115
		v150 = v117
		v151 = v118
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v126 = v115
	v127 = v117
	v128 = v118
	goto L37
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	if v131 != v132 {
		v149 = v126
		v150 = v127
		v151 = v128
		goto L33
	} else {
		goto L39
	}
L38:
	;
	goto L34
L39:
	;
	v134 = int32(4)
	v135 = v127 + v134
	v137 = v126 + v134
	v139 = v128 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v139) {
		v126 = v137
		v127 = v135
		v128 = v139
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v149 = v137
	v150 = v135
	v151 = v139
	goto L33
L42:
	;
	v182 = v159 - v160
	goto L31
L43:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	if v159 != v160 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v162 = int32(1)
	v167 = v156 + int32(-1)
	if v167 == int32(0) {
		goto L32
	} else {
		goto L46
	}
L46:
	;
	v154 = v154 + v162
	v155 = v155 + v162
	v156 = v167
	goto L43
L47:
	;
	v183 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v183)
	v186 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v186 {
		v227 = v109
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v189
	F__serverLog(m, int32(2), int32(_a405), v11+int32(16))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int32(0)
L50:
	;
	v227 = v109
	goto L1
L51:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v218
	F__serverLog(m, int32(3), int32(_a406), v11)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	v227 = v213
	goto L1
}
func F_doSlotRangeListsOverlap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = v8 + int32(8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v12
	goto L1
L1:
	;
	v16 = int32(0)
	v18 = v8 + int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v20 == v16 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v87
L3:
	;
	if v20 == int32(0) {
		v87 = v16
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20+base.B2i32(v23 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v29
	goto L4
L6:
	;
	v33 = v20
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v39
	goto L9
L8:
	;
	v87 = int32(1)
	goto L2
L9:
	;
	goto L10
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v49 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L8
L12:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v76 < v78 {
		goto L10
	} else {
		goto L21
	}
L13:
	;
	if v49 != 0 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v49+base.B2i32(v52 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v58
	goto L14
L16:
	;
	v61 = v8 + int32(8)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v63 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v63 == int32(0) {
		v87 = v16
		goto L2
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v63+base.B2i32(v66 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v72
	goto L18
L20:
	;
	v33 = v63
	goto L7
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v81 < v80 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	goto L11
}
func F_freeSlotMigrationJob(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v5 == int32(0) {
		if v4 == int32(0) {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
			F_sdsfree(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = int32(0)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
				F_listRelease(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
					F_sdsfree(m, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
						F_sdsfree(m, v33)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
							F_sdsfree(m, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
								F_sdsfree(m, v39)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									F_valkey_free(m, l0)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
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
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
			m.T0[v17].(func(*base.Module, int32))(m, v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = int32(0)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
				F_sdsfree(m, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = int32(0)
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
					F_listRelease(m, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						F_sdsfree(m, v30)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
							F_sdsfree(m, v33)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
								F_sdsfree(m, v36)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
									F_sdsfree(m, v39)
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return
									} else {
										F_valkey_free(m, l0)
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
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
	} else {
		if v4 != 0 {
			F__serverAssert(m, int32(_a393), int32(_a386), int32(2186))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+216)) = int32(0)
			F_freeClientAsync(m, v5)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
				F_sdsfree(m, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = int32(0)
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
					F_listRelease(m, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						F_sdsfree(m, v30)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
							F_sdsfree(m, v33)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
								F_sdsfree(m, v36)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
									F_sdsfree(m, v39)
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return
									} else {
										F_valkey_free(m, l0)
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
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
func F_slotExportConnectHandler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_proceedWithSlotMigration(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_slotExportJobBeginSnapshotToTargetSocket(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
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
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
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
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v11 = int32(-1)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	if v13 != v11 {
		v298 = v11
		m.G0 = v9 + int32(112)
		return v298
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[220]))
		if v17 != int32(-1) {
			F__serverAssert(m, int32(_a397), int32(_a386), int32(1619))
			mBase = m.M
			v310 = m.ExcPending
			if v310 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[221]))
			if v21 != int32(-1) {
				F__serverAssert(m, int32(_a397), int32(_a386), int32(1619))
				mBase = m.M
				v310 = m.ExcPending
				if v310 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v24 = int32(-1)
				v26 = v9 + int32(104)
				v31 = m.G0
				v33 = v31 - int32(64)
				m.G0 = v33
				v36 = F_pipe(m, v26)
				mBase = m.M
				if v36 != 0 {
					v93 = v24
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = int32(2048)
					v70 = F_fcntl(m, v65, int32(4), v33+int32(16))
					mBase = m.M
					if v70 != 0 {
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
						v87 = F_close(m, v86)
						mBase = m.M
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
						v89 = F_close(m, v88)
						mBase = m.M
						v93 = int32(-1)
					} else {
						v93 = int32(0)
					}
				}
				m.G0 = v33 + int32(64)
				if v93 == int32(-1) {
					v298 = v24
					m.G0 = v9 + int32(112)
					return v298
				} else {
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v9)+104))
					*(*int32)(unsafe.Add(mBase, _consts[220])) = v100
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v9)+108))
					v103 = int32(-1)
					v110 = m.G0
					v112 = v110 - int32(64)
					m.G0 = v112
					v115 = F_pipe(m, v9+int32(104))
					mBase = m.M
					if v115 != 0 {
						v172 = v103
					} else {
						v172 = int32(0)
					}
					m.G0 = v112 + int32(64)
					if v172 != int32(-1) {
						v185 = int32(_a20)
						v186 = *(*int32)(unsafe.Add(mBase, uint32(v9)+108))
						*(*int32)(unsafe.Add(mBase, _consts[221])) = v186
						v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
						v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
						*(*int32)(unsafe.Add(mBase, _consts[222])) = v190
						v192 = *(*int32)(unsafe.Add(mBase, uint32(v9)+104))
						v193 = int32(1)
						v195 = F_serverFork(m, int32(5))
						mBase = m.M
						v198 = m.ExcPending
						if v198 != 0 {
							return int32(0)
						} else {
							switch v195 + int32(1) {
							case 0:
								v243 = *(*int32)(unsafe.Add(mBase, _consts[28]))
								if int32(3) < v243 {
									v256 = F_close(m, v102)
									mBase = m.M
									v257 = int32(_a20)
									v258 = *(*int32)(unsafe.Add(mBase, _consts[220]))
									v259 = F_close(m, v258)
									mBase = m.M
									v261 = *(*int32)(unsafe.Add(mBase, _consts[221]))
									v262 = F_close(m, v261)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, _consts[222])) = int32(0)
									v298 = int32(-1)
									m.G0 = v9 + int32(112)
									return v298
								} else {
									v247 = *(*int32)(unsafe.Add(mBase, _consts[18]))
									v248 = F___strerror_l(m, v247, v247)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v248
									F__serverLog(m, int32(3), int32(_a398), v9+int32(16))
									mBase = m.M
									v255 = m.ExcPending
									if v255 != 0 {
										return int32(0)
									} else {
										v256 = F_close(m, v102)
										mBase = m.M
										v257 = int32(_a20)
										v258 = *(*int32)(unsafe.Add(mBase, _consts[220]))
										v259 = F_close(m, v258)
										mBase = m.M
										v261 = *(*int32)(unsafe.Add(mBase, _consts[221]))
										v262 = F_close(m, v261)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, _consts[222])) = int32(0)
										v298 = int32(-1)
										m.G0 = v9 + int32(112)
										return v298
									}
								}
							case 1:
								F_rioInitWithFd(m, v9+int32(24), v102)
								mBase = m.M
								v204 = m.ExcPending
								if v204 != 0 {
									return int32(0)
								} else {
									v206 = *(*int32)(unsafe.Add(mBase, _consts[220]))
									v207 = F_close(m, v206)
									mBase = m.M
									v214 = F_childSnapshotForSyncSlot(m, v9+int32(24), l0)
									mBase = m.M
									v215 = m.ExcPending
									if v215 != 0 {
										return int32(0)
									} else {
										if v214 != 0 {
											v228 = v193
											F_rioFreeFd(m, v9+int32(24))
											mBase = m.M
											v232 = m.ExcPending
											if v232 != 0 {
												return int32(0)
											} else {
												v233 = F_close(m, v102)
												mBase = m.M
												v235 = *(*int32)(unsafe.Add(mBase, _consts[221]))
												v236 = F_close(m, v235)
												mBase = m.M
												v240 = F_read(m, v192, v9+int32(104), int32(1))
												mBase = m.M
												F__exit(m, v228)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v218 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
											v219 = m.T0[v218].(func(*base.Module, int32) int32)(m, v9+int32(24))
											mBase = m.M
											v220 = m.ExcPending
											if v220 != 0 {
												return int32(0)
											} else {
												if v219 == int32(0) {
													v228 = v193
													F_rioFreeFd(m, v9+int32(24))
													mBase = m.M
													v232 = m.ExcPending
													if v232 != 0 {
														return int32(0)
													} else {
														v233 = F_close(m, v102)
														mBase = m.M
														v235 = *(*int32)(unsafe.Add(mBase, _consts[221]))
														v236 = F_close(m, v235)
														mBase = m.M
														v240 = F_read(m, v192, v9+int32(104), int32(1))
														mBase = m.M
														F__exit(m, v228)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													F_sendChildCowInfo(m, int32(4), int32(_a399))
													mBase = m.M
													v226 = m.ExcPending
													if v226 != 0 {
														return int32(0)
													} else {
														v228 = int32(0)
														F_rioFreeFd(m, v9+int32(24))
														mBase = m.M
														v232 = m.ExcPending
														if v232 != 0 {
															return int32(0)
														} else {
															v233 = F_close(m, v102)
															mBase = m.M
															v235 = *(*int32)(unsafe.Add(mBase, _consts[221]))
															v236 = F_close(m, v235)
															mBase = m.M
															v240 = F_read(m, v192, v9+int32(104), int32(1))
															mBase = m.M
															F__exit(m, v228)
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
							default:
								v268 = *(*int32)(unsafe.Add(mBase, _consts[28]))
								if int32(2) < v268 {
									v278 = F_close(m, v102)
									mBase = m.M
									v279 = int32(0)
									v280 = int32(_a20)
									v281 = *(*int32)(unsafe.Add(mBase, _consts[223]))
									v283 = *(*int32)(unsafe.Add(mBase, _consts[220]))
									v287 = F_aeCreateFileEvent(m, v281, v283, int32(1), int32(108), v279)
									mBase = m.M
									v288 = m.ExcPending
									if v288 != 0 {
										return int32(0)
									} else {
										if v287 == int32(-1) {
											F__serverPanic_1(m, int32(_a386), int32(1679), int32(_a400), int32(0))
											mBase = m.M
											v317 = m.ExcPending
											if v317 != 0 {
												return int32(0)
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v291 = F_close(m, v192)
											mBase = m.M
											v293 = *(*int32)(unsafe.Add(mBase, _consts[224]))
											if v293 == int32(0) {
												v298 = v279
												m.G0 = v9 + int32(112)
												return v298
											} else {
												F_debugPauseProcess(m)
												mBase = m.M
												v297 = m.ExcPending
												if v297 != 0 {
													return int32(0)
												} else {
													v298 = v279
													m.G0 = v9 + int32(112)
													return v298
												}
											}
										}
									}
								} else {
									v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v271
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v195
									F__serverLog(m, int32(2), int32(_a401), v9)
									mBase = m.M
									v277 = m.ExcPending
									if v277 != 0 {
										return int32(0)
									} else {
										v278 = F_close(m, v102)
										mBase = m.M
										v279 = int32(0)
										v280 = int32(_a20)
										v281 = *(*int32)(unsafe.Add(mBase, _consts[223]))
										v283 = *(*int32)(unsafe.Add(mBase, _consts[220]))
										v287 = F_aeCreateFileEvent(m, v281, v283, int32(1), int32(108), v279)
										mBase = m.M
										v288 = m.ExcPending
										if v288 != 0 {
											return int32(0)
										} else {
											if v287 == int32(-1) {
												F__serverPanic_1(m, int32(_a386), int32(1679), int32(_a400), int32(0))
												mBase = m.M
												v317 = m.ExcPending
												if v317 != 0 {
													return int32(0)
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v291 = F_close(m, v192)
												mBase = m.M
												v293 = *(*int32)(unsafe.Add(mBase, _consts[224]))
												if v293 == int32(0) {
													v298 = v279
													m.G0 = v9 + int32(112)
													return v298
												} else {
													F_debugPauseProcess(m)
													mBase = m.M
													v297 = m.ExcPending
													if v297 != 0 {
														return int32(0)
													} else {
														v298 = v279
														m.G0 = v9 + int32(112)
														return v298
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v178 = F_close(m, v102)
						mBase = m.M
						v179 = int32(_a20)
						v180 = *(*int32)(unsafe.Add(mBase, _consts[220]))
						v181 = F_close(m, v180)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, _consts[220])) = int32(-1)
						v298 = v103
						m.G0 = v9 + int32(112)
						return v298
					}
				}
			}
		}
	}
}
func F_slotExportTryDoPause(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 != 0 {
		F__serverAssert(m, int32(_a402), int32(_a386), int32(1461))
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v11 = int32(-1)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, _consts[225])))
		if v13&int32(8) != 0 {
			v55 = v11
			m.G0 = v8 + int32(16)
			return v55
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[226]))
			if v17 < int32(0) {
				v25 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				if int32(2) < v25 {
					v40 = F_mstime(m)
					mBase = m.M
					v42 = *(*int64)(unsafe.Add(mBase, _consts[227]))
					v45 = v40 + v42<<(uint(int64(1))%64)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v45
					F_pauseActions(m, int32(3), v45, int32(29))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_sendSyncSlotsMessage(m, l0, int32(_a403))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v55 = int32(0)
							m.G0 = v8 + int32(16)
							return v55
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
					v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v30
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = v29
					F__serverLog(m, int32(2), int32(_a404), v8)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v40 = F_mstime(m)
						mBase = m.M
						v42 = *(*int64)(unsafe.Add(mBase, _consts[227]))
						v45 = v40 + v42<<(uint(int64(1))%64)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v45
						F_pauseActions(m, int32(3), v45, int32(29))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_sendSyncSlotsMessage(m, l0, int32(_a403))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v55 = int32(0)
								m.G0 = v8 + int32(16)
								return v55
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)+160))
				if base.Ui64(base.I64_extend_i32_u(v17)) < base.Ui64(v21) {
					v55 = v11
					m.G0 = v8 + int32(16)
					return v55
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _consts[28]))
					if int32(2) < v25 {
						v40 = F_mstime(m)
						mBase = m.M
						v42 = *(*int64)(unsafe.Add(mBase, _consts[227]))
						v45 = v40 + v42<<(uint(int64(1))%64)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v45
						F_pauseActions(m, int32(3), v45, int32(29))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_sendSyncSlotsMessage(m, l0, int32(_a403))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v55 = int32(0)
								m.G0 = v8 + int32(16)
								return v55
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
						v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v30
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = v29
						F__serverLog(m, int32(2), int32(_a404), v8)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v40 = F_mstime(m)
							mBase = m.M
							v42 = *(*int64)(unsafe.Add(mBase, _consts[227]))
							v45 = v40 + v42<<(uint(int64(1))%64)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v45
							F_pauseActions(m, int32(3), v45, int32(29))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_sendSyncSlotsMessage(m, l0, int32(_a403))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v55 = int32(0)
									m.G0 = v8 + int32(16)
									return v55
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_slotMigrationJobReadEstablishResponse(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
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
	var v20 int32
	_ = v20
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
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
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int64
	_ = v203
	var v206 int64
	_ = v206
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+222)))
	if v16 != 0 {
		v20 = int32(1)
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	if v20 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L2
L4:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+223)))
	v20 = base.B2i32(v17 != int32(0))
	goto L3
L5:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+201)))
	if v21&int32(4) != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+216))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+156))
	if base.Ui32(int32(20)) < base.Ui32(v25) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+204))
	if v32 != 0 {
		v40 = v32
		goto L10
	} else {
		goto L11
	}
L8:
	;
	if int32(1)<<(uint(v25)%32)&int32(1835040) != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v41 = int32(0)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-1)))))
	switch v44 & int32(7) {
	case 0:
		goto L20
	case 1:
		goto L19
	case 2:
		goto L18
	case 3:
		goto L17
	case 4:
		goto L16
	default:
		v84 = v40
		v85 = v41
		goto L15
	}
L11:
	;
	v33 = F_sdsempty(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+204)) = v33
	v37 = F_sdsMakeRoomForNonGreedy(m, v33, int32(16384))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+204)) = v37
	v40 = v37
	goto L10
L15:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+76))
	v90 = m.T0[v89].(func(*base.Module, int32, int32, int32) int32)(m, l0, v84, v85)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L12
	} else {
		goto L22
	}
L16:
	;
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v40+int32(-9))))
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v40+int32(-17))))
	v84 = v40 + base.I32_wrap_i64(v79)
	v85 = base.I32_wrap_i64(v76 - v79)
	goto L15
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-5))))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-9))))
	v84 = v40 + v71
	v85 = v68 - v71
	goto L15
L18:
	;
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40+int32(-3)))))
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40+int32(-5)))))
	v84 = v40 + v63
	v85 = v60 - v63
	goto L15
L19:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-2)))))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-3)))))
	v84 = v40 + v55
	v85 = v52 - v55
	goto L15
L20:
	;
	v84 = v40 + int32(base.Ui32(v44)>>(uint(int32(3))%32))
	v85 = v41
	goto L15
L21:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v97 == int32(3) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	if v90 < int32(1) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v24)+204))
	F_sdsIncrLen(m, v94, v90)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v24)+204))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+int32(-1)))))
	v107 = v105 & int32(7)
	switch v107 {
	case 0:
		goto L36
	case 1:
		goto L35
	case 2:
		goto L34
	case 3:
		goto L33
	case 4:
		goto L32
	default:
		goto L30
	}
L26:
	;
	F_freeClientAsync(m, v13)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	goto L1
L28:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v221 != int32(45) {
		goto L62
	} else {
		goto L63
	}
L29:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+84))
	v219 = m.T0[v218].(func(*base.Module, int32, int32) int32)(m, l0, int32(105))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L12
	} else {
		goto L61
	}
L30:
	;
	F_finishSlotMigrationJob(m, v24, int32(18), int32(_a394))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L12
	} else {
		goto L60
	}
L31:
	;
	if base.Ui32(v122) < base.Ui32(int32(2)) {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(-17))))
	v122 = v121
	goto L31
L33:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(-9))))
	v122 = v118
	goto L31
L34:
	;
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+int32(-5)))))
	v122 = v115
	goto L31
L35:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+int32(-3)))))
	v122 = v112
	goto L31
L36:
	;
	v122 = int32(base.Ui32(v105) >> (uint(int32(3)) % 32))
	goto L31
L37:
	;
	switch v107 + int32(-1) {
	case 0:
		goto L58
	case 1:
		goto L57
	case 2:
		goto L56
	case 3:
		goto L55
	default:
		goto L30
	}
L38:
	;
	switch v107 {
	default:
		goto L44
	case 1:
		goto L43
	case 2:
		goto L42
	case 3:
		goto L41
	case 4:
		goto L40
	}
L39:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v139+int32(-2)))))
	if v143 != int32(13) {
		goto L37
	} else {
		goto L45
	}
L40:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(-17))))
	v139 = v138
	goto L39
L41:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(-9))))
	v139 = v135
	goto L39
L42:
	;
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+int32(-5)))))
	v139 = v132
	goto L39
L43:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+int32(-3)))))
	v139 = v129
	goto L39
L44:
	;
	v139 = int32(base.Ui32(v105) >> (uint(int32(3)) % 32))
	goto L39
L45:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+int32(-1)))))
	switch v151 & int32(7) {
	case 0:
		goto L52
	case 1:
		goto L51
	case 2:
		goto L50
	case 3:
		goto L49
	case 4:
		goto L48
	default:
		v168 = int32(0)
		goto L47
	}
L46:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v170+int32(-1)))))
	if v174 == int32(10) {
		goto L28
	} else {
		goto L53
	}
L47:
	;
	v170 = v168
	goto L46
L48:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(-17))))
	v168 = v167
	goto L47
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(-9))))
	v170 = v164
	goto L46
L50:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+int32(-5)))))
	v170 = v161
	goto L46
L51:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+int32(-3)))))
	v170 = v158
	goto L46
L52:
	;
	v170 = int32(base.Ui32(v151) >> (uint(int32(3)) % 32))
	goto L46
L53:
	;
	goto L37
L54:
	;
	if v209 != 0 {
		goto L29
	} else {
		goto L59
	}
L55:
	;
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v102+int32(-9))))
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v102+int32(-17))))
	v209 = base.I32_wrap_i64(v203 - v206)
	goto L54
L56:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(-5))))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(-9))))
	v209 = v196 - v199
	goto L54
L57:
	;
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+int32(-3)))))
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+int32(-5)))))
	v209 = v189 - v192
	goto L54
L58:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+int32(-2)))))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+int32(-3)))))
	v209 = v182 - v185
	goto L54
L59:
	;
	goto L30
L60:
	;
	goto L1
L61:
	;
	goto L1
L62:
	;
	F_updateSlotMigrationJobState(m, v24, int32(12))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L12
	} else {
		goto L68
	}
L63:
	;
	v224 = F_sdsempty(m)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L12
	} else {
		goto L64
	}
L64:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v24)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v226
	v230 = F_sdscatfmt(m, v224, int32(_a396), v11)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	F_finishSlotMigrationJob(m, v24, int32(18), v230)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	F_sdsfree(m, v230)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	goto L1
L68:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+84))
	v242 = m.T0[v241].(func(*base.Module, int32, int32) int32)(m, l0, int32(107))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	F_clusterDoBeforeSleep(m, int32(64))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	F_sendSyncSlotsMessage(m, v24, int32(_a395))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v24)+204))
	F_sdsfree(m, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L12
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+204)) = int32(0)
	goto L1
}
func F_slotStatForSortAscCmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if v5 != v6 {
		return base.B2i32(base.Ui64(v6) < base.Ui64(v5)) - base.B2i32(base.Ui64(v5) < base.Ui64(v6))
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		return v8 - v9
	}
}
func F_slotStatForSortDescCmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v5 != v6 {
		return base.B2i32(base.Ui64(v6) < base.Ui64(v5)) - base.B2i32(base.Ui64(v5) < base.Ui64(v6))
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		return v8 - v9
	}
}
